package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func TitleSquare(
	title, value interface{},
) string {
	return fmt.Sprintf(
		constants.SquareTitleWrapFormat,
		title,
		value)
}
