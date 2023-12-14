package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

func TitleCurlyWrap(
	title, value interface{},
) string {
	return fmt.Sprintf(
		constants.CurlyTitleWrapFormat,
		toString(title),
		toString(value))
}
