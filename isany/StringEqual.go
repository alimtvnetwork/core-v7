package isany

import "gitlab.com/evatix-go/core/internal/strutilinternal"

func StringEqual(
	left, right interface{},
) bool {
	leftString := strutilinternal.AnyToFieldNameString(left)
	rightString := strutilinternal.AnyToFieldNameString(right)

	return leftString == rightString
}
