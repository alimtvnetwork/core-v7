package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// WithCurly
//
// {%v}
func WithCurly(
	source interface{},
) string {
	return fmt.Sprintf(constants.CurlyWrap, source)
}
