package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

// WithBrackets
//
// [%v]
func WithBrackets(
	source interface{},
) string {
	toStr := toString(source)
	
	return fmt.Sprintf(
		constants.BracketWrapFormat,
		toStr)
}
