package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

func CurlyWrap(
	source interface{},
) string {
	toStr := convertinteranl.
		AnyTo.
		SmartString(source)
	
	return fmt.Sprintf(
		constants.CurlyWrapFormat,
		toStr)
}
