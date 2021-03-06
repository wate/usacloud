package funcs

import (
	"bytes"
	"fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
	"io"
	"io/ioutil"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ObjectStoragePut(ctx command.Context, params *params.PutObjectStorageParam) error {
	useStdIn := false
	srcPath := ""
	destPath := ""
	var srcInfo os.FileInfo

	switch ctx.NArgs() {
	case 1:
		useStdIn = true
		destPath = ctx.Args()[0]

		if params.Recursive {
			return fmt.Errorf("--recursie can't be used with STDIN")
		}
		// validate stdin
		fi, err := command.GlobalOption.In.Stat()
		if err != nil {
			return fmt.Errorf("STDIN Stat() is failed: %s", err)
		}
		// if using pipe with curl, fi.Size() will return zero.
		// so check file mode is os.ModeNamedPipe
		if fi.Size() == 0 && fi.Mode()&os.ModeNamedPipe == 0 {
			return fmt.Errorf("STDIN is Empty")
		}

	case 2:
		srcPath = filepath.Clean(ctx.Args()[0])
		destPath = ctx.Args()[1]
		// validate filepath
		info, err := os.Stat(srcPath)
		if err != nil {
			return fmt.Errorf("file[%s] is not exists: %s", srcPath, err)
		}
		srcInfo = info
	default:
		return fmt.Errorf("ObjectStoragePut is failed: %s", "Only two argument can be specified")
	}

	// destPath
	if destPath != "" && strings.HasPrefix(destPath, "/") {
		destPath = strings.Replace(destPath, "/", "", 1)
	}
	// if destPath is dir, set filename from srcPath
	if strings.HasSuffix(destPath, "/") {
		destPath = fmt.Sprintf("%s%s", destPath, filepath.Base(srcPath))
	}

	// on SakuraCloud, bucket name is same as AccessKey
	if params.Bucket == "" {
		params.Bucket = params.AccessKey
	}

	auth, err := aws.GetAuth(params.AccessKey, params.SecretKey)
	if err != nil {
		return fmt.Errorf("ObjectStoragePut is failed: %s", err)
	}
	client := s3.New(auth, aws.Region{
		Name:       "us-west-2",
		S3Endpoint: "https://b.sakurastorage.jp",
	})

	bucket := client.Bucket(params.Bucket)

	var putFunc func() error

	if useStdIn {
		putFunc = func() error {
			return objectStoragePutReaderMultiNonSeeker(destPath, command.GlobalOption.In, bucket, params.ContentType)
		}
	} else {
		if srcInfo.IsDir() {
			if !params.Recursive {
				return fmt.Errorf("%q is directory. Use -r or --recursive flag", srcPath)
			}
			params.ContentType = "" // when directory mode, set empty to content-type
			putFunc = func() error {
				return objectStoragePutRecursive(destPath, srcPath, srcPath, params.Recursive, bucket, params.ContentType)
			}

		} else {
			putFunc = func() error {
				return objectStoragePut(destPath, srcPath, bucket, params.ContentType)
			}
		}
	}

	return internal.ExecWithProgress(
		fmt.Sprintf("Still uploading[%q]...", destPath),
		fmt.Sprintf("Upload [%q]", destPath),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			if err := putFunc(); err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
}

func objectStoragePutRecursive(remotePath, baseDir, targetDir string, rec bool, bucket *s3.Bucket, contentType string) error {

	// if recursive is false , process only files under targetDir
	if !rec && targetDir != baseDir {
		return nil
	}

	entries, err := ioutil.ReadDir(targetDir)
	if err != nil {
		return err
	}

	for _, fi := range entries {
		src := filepath.Join(targetDir, fi.Name())
		// this is used by object storage , so use path.Join(not filepath.Join)
		dest := path.Join(remotePath, fi.Name())
		if fi.IsDir() {
			err := objectStoragePutRecursive(dest, baseDir, src, rec, bucket, contentType)
			if err != nil {
				return err
			}
		} else {
			err := objectStoragePut(dest, src, bucket, contentType)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func objectStoragePut(destPath, srcPath string, bucket *s3.Bucket, contentType string) error {

	if contentType == "" {
		// set content-type from extension
		ext := filepath.Ext(srcPath)
		contentType = mime.TypeByExtension(ext)
	}

	// target file
	file, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return objectStoragePutReader(destPath, file, bucket, contentType)
}

func objectStoragePutReader(destPath string, file *os.File, bucket *s3.Bucket, contentType string) error {
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	if fi.Size() > 8*1024*1024 { // over 8MB
		return objectStoragePutReaderMulti(destPath, file, bucket, contentType)
	}

	// put file
	err = bucket.PutReader(destPath, file, fi.Size(), contentType, s3.PublicRead)
	if err != nil {
		return err
	}

	return nil
}

func objectStoragePutReaderMulti(destPath string, file *os.File, bucket *s3.Bucket, contentType string) error {
	partSize := int64(7 * 1024 * 1024) // chunk size = 7MB
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	size := fi.Size()
	if size != 0 && size < partSize {
		partSize = size
	}

	m, err := bucket.InitMulti(destPath, contentType, s3.PublicRead)
	if err != nil {
		return err
	}

	// put file
	_, err = m.PutAll(file, partSize)
	if err != nil {
		return err
	}

	return nil
}

func objectStoragePutReaderMultiNonSeeker(destPath string, file *os.File, bucket *s3.Bucket, contentType string) error {
	m, err := bucket.Multi(destPath, contentType, s3.PublicRead)
	if err != nil {
		return err
	}
	var (
		parts      = make([]s3.Part, 0)
		partSize   = 1024 * 1024 * 5       // 5MB
		readBuffer = make([]byte, 1024*64) // 64KB
		partBuffer bytes.Buffer
		sendPart   = partSender(&partBuffer, m)
	)
	for {
		n, err := file.Read(readBuffer)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			if partBuffer.Len() > 0 {
				part, err := sendPart()
				if err != nil {
					return err
				}
				parts = append(parts, part)
			}

			break
		} else {
			combined := partBuffer.Len() + n
			if combined > partSize {
				needed := partSize - partBuffer.Len()
				partBuffer.Write(readBuffer[:needed])
				part, err := sendPart()
				if err != nil {
					return err
				}
				parts = append(parts, part)
				partBuffer.Write(readBuffer[needed:n])
			} else {
				partBuffer.Write(readBuffer[:n])
			}
		}
	}

	err = m.Complete(parts)
	if err != nil {
		return err
	}

	return nil
}

func partSender(partBuffer *bytes.Buffer, multi *s3.Multi) func() (s3.Part, error) {
	var partNr int
	return func() (s3.Part, error) {
		partNr++
		defer partBuffer.Reset()
		reader := bytes.NewReader(partBuffer.Bytes())
		return multi.PutPart(partNr, reader)
	}
}
