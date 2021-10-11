package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// WithParenthesis
//
// (%v)
func WithParenthesis(
	source interface{},
) string {
	return fmt.Sprintf(constants.ParenthesisWrap, source)
}
