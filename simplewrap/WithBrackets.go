package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// WithBrackets
//
// [%v]
func WithBrackets(
	source interface{},
) string {
	return fmt.Sprintf(constants.BracketWrapFormat, source)
}
