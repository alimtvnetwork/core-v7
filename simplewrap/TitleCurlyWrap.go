package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func TitleCurlyWrap(
	title, value interface{},
) string {
	return fmt.Sprintf(
		constants.CurlyTitleWrapFormat,
		title,
		value)
}
