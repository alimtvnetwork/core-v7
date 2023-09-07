package isany

import "gitlab.com/auk-go/core/internal/strutilinternal"

func StringEqual(
	left, right interface{},
) bool {
	leftString := strutilinternal.AnyToFieldNameString(left)
	rightString := strutilinternal.AnyToFieldNameString(right)

	return leftString == rightString
}
