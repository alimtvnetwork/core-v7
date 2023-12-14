package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

// WithCurly
//
// {%v}
func WithCurly(
	source interface{},
) string {
	toStr := convertinteranl.
		AnyTo.
		SmartString(source)
	
	return fmt.Sprintf(
		constants.CurlyWrapFormat,
		toStr)
}
