package reflectinternal

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func TypeNamesReferenceString(
	isFullName bool,
	anyItems ...interface{},
) string {
	return "Reference (Types): " + strings.Join(
		TypeNames(isFullName, anyItems...),
		constants.CommaSpace)
}
