package reflectinternal

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func TypeNamesString(
	isFullName bool,
	anyItems ...interface{},
) string {
	return strings.Join(
		TypeNames(isFullName, anyItems...),
		constants.CommaSpace)
}
