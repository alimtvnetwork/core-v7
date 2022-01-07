package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func TitleSquareMetaUsingFmt(
	title,
	value,
	meta fmt.Stringer,
) string {
	return fmt.Sprintf(
		constants.SquareTitleMetaWrapFormat,
		title,
		value,
		meta)
}
