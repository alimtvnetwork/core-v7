package simplewrap

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func SquareWrap(
	source interface{},
) string {
	return fmt.Sprintf(
		constants.SquareWrapFormat,
		source)
}
