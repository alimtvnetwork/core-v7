package reflectinternal

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func TypeNamesReferenceString(
	isFullName bool,
	anyItems ...interface{},
) string {
	return "Reference (Types): " + strings.Join(
		TypeNames(isFullName, anyItems...),
		constants.CommaSpace)
}
