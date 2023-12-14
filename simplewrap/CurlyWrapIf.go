package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

func CurlyWrapIf(
	isCurly bool,
	source interface{},
) string {
	if !isCurly {
		return toString(source)
	}
	
	return fmt.Sprintf(
		constants.CurlyWrapFormat,
		toString(source))
}
