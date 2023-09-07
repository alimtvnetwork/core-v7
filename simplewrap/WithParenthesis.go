package simplewrap

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// WithParenthesis
//
// (%v)
func WithParenthesis(
	source interface{},
) string {
	return fmt.Sprintf(constants.ParenthesisWrapFormat, source)
}
