package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

func ParenthesisWrap(
	source interface{},
) string {
	return fmt.Sprintf(
		constants.ParenthesisWrapFormat,
		toString(source))
}
