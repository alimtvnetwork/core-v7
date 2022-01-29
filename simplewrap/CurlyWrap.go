package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func CurlyWrap(
	source interface{},
) string {
	return fmt.Sprintf(
		constants.CurlyWrapFormat,
		source)
}
