// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package cli

import (
	"encoding/json"
	"fmt"
	"github.com/imdario/mergo"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/completion"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
	"strings"
)

func init() {
	listParam := params.NewListPriceParam()

	cliCommand := &cli.Command{
		Name:    "price",
		Aliases: []string{"public-price"},
		Usage:   "A manage commands of Price",
		Action: func(c *cli.Context) error {
			comm := c.App.Command("list")
			if comm != nil {
				return comm.Run(c)
			}
			return cli.ShowSubcommandHelp(c)
		},
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls", "find"},
				Usage:   "List Price",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:  "name",
						Usage: "set filter by name(s)",
					},
					&cli.Int64SliceFlag{
						Name:  "id",
						Usage: "set filter by id(s)",
					},
					&cli.IntFlag{
						Name:    "from",
						Aliases: []string{"offset"},
						Usage:   "set offset",
					},
					&cli.IntFlag{
						Name:    "max",
						Aliases: []string{"limit"},
						Usage:   "set limit",
					},
					&cli.StringSliceFlag{
						Name:  "sort",
						Usage: "set field(s) for sort",
					},
					&cli.StringFlag{
						Name:  "param-template",
						Usage: "Set input parameter from string(JSON)",
					},
					&cli.StringFlag{
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
					&cli.StringFlag{
						Name:    "output-type",
						Aliases: []string{"out"},
						Usage:   "Output type [json/csv/tsv]",
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:    "quiet",
						Aliases: []string{"q"},
						Usage:   "Only display IDs",
					},
					&cli.StringFlag{
						Name:    "format",
						Aliases: []string{"fmt"},
						Usage:   "Output format(see text/template package document for detail)",
					},
					&cli.StringFlag{
						Name:  "format-file",
						Usage: "Output format from file(see text/template package document for detail)",
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					command.GlobalOption.Validate(false)

					// build command context
					ctx := command.NewContext(c, realArgs, listParam)

					// Set option values
					if c.IsSet("name") {
						listParam.Name = c.StringSlice("name")
					}
					if c.IsSet("id") {
						listParam.Id = c.Int64Slice("id")
					}
					if c.IsSet("from") {
						listParam.From = c.Int("from")
					}
					if c.IsSet("max") {
						listParam.Max = c.Int("max")
					}
					if c.IsSet("sort") {
						listParam.Sort = c.StringSlice("sort")
					}
					if c.IsSet("param-template") {
						listParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						listParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						listParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						listParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						listParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						listParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						listParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						listParam.FormatFile = c.String("format-file")
					}

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.PriceListCompleteArgs(ctx, listParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completion.FlagNames(c, commandName)
										return
									} else {
										completion.PriceListCompleteArgs(ctx, listParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.PriceListCompleteFlags(ctx, listParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completion.FlagNames(c, commandName)
							return
						} else {
							completion.PriceListCompleteArgs(ctx, listParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					listParam.ParamTemplate = c.String("param-template")
					listParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(listParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewListPriceParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.MergeWithOverwrite(listParam, p)
					}

					// Set option values
					if c.IsSet("name") {
						listParam.Name = c.StringSlice("name")
					}
					if c.IsSet("id") {
						listParam.Id = c.Int64Slice("id")
					}
					if c.IsSet("from") {
						listParam.From = c.Int("from")
					}
					if c.IsSet("max") {
						listParam.Max = c.Int("max")
					}
					if c.IsSet("sort") {
						listParam.Sort = c.StringSlice("sort")
					}
					if c.IsSet("param-template") {
						listParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						listParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						listParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						listParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						listParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						listParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						listParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						listParam.FormatFile = c.String("format-file")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(false); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// Generate skeleton
					if listParam.GenerateSkeleton {
						listParam.GenerateSkeleton = false
						listParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(listParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := listParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), listParam)

					// Run command with params
					return funcs.PriceList(ctx, listParam)

				},
			},
		},
	}

	// build Category-Resource mapping
	AppendResourceCategoryMap("price", &schema.Category{
		Key:         "information",
		DisplayName: "Service/Product informations",
		Order:       90,
	})

	// build Category-Command mapping

	AppendCommandCategoryMap("price", "list", &schema.Category{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       1,
	})

	// build Category-Param mapping

	AppendFlagCategoryMap("price", "list", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("price", "list", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("price", "list", "format-file", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("price", "list", "from", &schema.Category{
		Key:         "limit-offset",
		DisplayName: "Limit/Offset options",
		Order:       2147483597,
	})
	AppendFlagCategoryMap("price", "list", "generate-skeleton", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("price", "list", "id", &schema.Category{
		Key:         "filter",
		DisplayName: "Filter options",
		Order:       2147483587,
	})
	AppendFlagCategoryMap("price", "list", "max", &schema.Category{
		Key:         "limit-offset",
		DisplayName: "Limit/Offset options",
		Order:       2147483597,
	})
	AppendFlagCategoryMap("price", "list", "name", &schema.Category{
		Key:         "filter",
		DisplayName: "Filter options",
		Order:       2147483587,
	})
	AppendFlagCategoryMap("price", "list", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("price", "list", "param-template", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("price", "list", "param-template-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("price", "list", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("price", "list", "sort", &schema.Category{
		Key:         "sort",
		DisplayName: "Sort options",
		Order:       2147483607,
	})

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}
