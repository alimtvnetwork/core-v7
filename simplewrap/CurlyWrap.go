package simplewrap

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func CurlyWrap(
	source interface{},
) string {
	return fmt.Sprintf(
		constants.CurlyWrapFormat,
		source)
}
