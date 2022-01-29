package keymk

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func curlyWrap(
	source interface{},
) string {
	return fmt.Sprintf(
		constants.CurlyWrapFormat,
		source)
}
