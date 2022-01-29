package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func ParenthesisWrapIf(
	isSquareWrap bool,
	source interface{},
) string {
	if !isSquareWrap {
		return fmt.Sprintf(
			constants.SprintValueFormat,
			source)
	}

	return fmt.Sprintf(
		constants.ParenthesisWrapFormat,
		source)
}
