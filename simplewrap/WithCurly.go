package simplewrap

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// WithCurly
//
// {%v}
func WithCurly(
	source interface{},
) string {
	return fmt.Sprintf(
		constants.CurlyWrapFormat,
		source)
}
