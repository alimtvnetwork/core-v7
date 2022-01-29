package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func ParenthesisWrap(
	source interface{},
) string {
	return fmt.Sprintf(
		constants.ParenthesisWrapFormat,
		source)
}
