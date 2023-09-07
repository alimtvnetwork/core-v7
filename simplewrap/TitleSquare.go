package simplewrap

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func TitleSquare(
	title, value interface{},
) string {
	return fmt.Sprintf(
		constants.SquareTitleWrapFormat,
		title,
		value)
}
