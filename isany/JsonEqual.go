package isany

import (
	"bytes"
	"encoding/json"
)

// JsonEqual
//
//	first checks if string is passed, if yes then only check string.
//	Or, else, marshal and check with error equal if both equal then true.
func JsonEqual(
	left, right interface{},
) bool {
	leftString, isLeftString := left.(string)
	rightString, isRightString := right.(string)

	if isLeftString && isRightString {
		return leftString == rightString
	}

	leftInteger, isLeftInteger := left.(int)
	rightInteger, isRightInteger := right.(int)

	if isLeftInteger && isRightInteger {
		return leftInteger == rightInteger
	}

	leftBytes, leftErr := json.Marshal(left)
	rightBytes, rightErr := json.Marshal(right)

	if leftErr != nil && rightErr != nil && rightErr.Error() != leftErr.Error() {
		return false
	}

	if leftErr != nil || rightErr != nil {
		return false
	}

	return bytes.Equal(leftBytes, rightBytes)
}
