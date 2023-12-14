package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

// TitleSquareMeta
//
// Example :
//   - constants.SquareTitleMetaWrapFormat
//   - "%v: [%v] (%v)"
func TitleSquareMeta(
	title,
	value,
	meta interface{},
) string {
	return fmt.Sprintf(
		constants.SquareTitleMetaWrapFormat,
		toString(title),
		toString(value),
		meta)
}
