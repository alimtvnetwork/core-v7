package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

func TitleSquareMetaUsingFmt(
	title,
	value,
	meta fmt.Stringer,
) string {
	return fmt.Sprintf(
		constants.SquareTitleMetaWrapFormat,
		toString(title),
		toString(value),
		meta)
}
