package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// TitleCurlyMeta
//
// Example :
//  - constants.CurlyTitleMetaWrapFormat
//  -  "%v: {%v} (%v)"
func TitleCurlyMeta(
	title,
	value,
	meta interface{},
) string {
	return fmt.Sprintf(
		constants.CurlyTitleMetaWrapFormat,
		title,
		value,
		meta)
}
