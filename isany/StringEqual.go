package isany

import "gitlab.com/evatix-go/core/internal/utilstringinternal"

func StringEqual(
	left, right interface{},
) bool {
	leftString := utilstringinternal.AnyToFieldNameString(left)
	rightString := utilstringinternal.AnyToFieldNameString(right)

	return leftString == rightString
}
