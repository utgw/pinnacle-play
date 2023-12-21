package render

import (
	"pinnacle-play/usecase/output"
	"pinnacle-play/view/model/json"
)

// JSONRender api render interface
type JSONRender interface {
	ConvertError(err error, code int) *json.Error
	ConvertPostGroupResult(output output.PostGroupOutput) *json.PostGroupResult
}
