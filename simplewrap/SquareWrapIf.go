package simplewrap

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func SquareWrapIf(
	isSquareWrap bool,
	source interface{},
) string {
	if !isSquareWrap {
		return fmt.Sprintf(
			constants.SprintValueFormat,
			source)
	}

	return fmt.Sprintf(
		constants.SquareWrapFormat,
		source)
}
