// Code generated by "stringer -type=HandlerType handler_type.go"; DO NOT EDIT

package schema

import "fmt"

const _HandlerType_name = "HandlerPathThroughHandlerPathThroughEachHandlerSortHandlerFilterByHandlerAndParamsHandlerOrParamsHandlerCustomFuncHandlerNoop"

var _HandlerType_index = [...]uint8{0, 18, 40, 51, 66, 82, 97, 114, 125}

func (i HandlerType) String() string {
	if i < 0 || i >= HandlerType(len(_HandlerType_index)-1) {
		return fmt.Sprintf("HandlerType(%d)", i)
	}
	return _HandlerType_name[_HandlerType_index[i]:_HandlerType_index[i+1]]
}