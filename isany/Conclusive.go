package isany

import (
	"reflect"
)

// Conclusive
//
//   - left AND right is null equal.
//   - either left OR right is null not equal.
//   - if both are defined and same pointer equal.
//   - if both are defined not pointer inconclusive.
func Conclusive(left, right interface{}) (isEqual, isConclusive bool) {
	if left == right {
		return true, true
	}

	if left == nil && right == nil {
		return true, true
	}

	if left == nil || right == nil {
		return false, true
	}

	leftRv := reflect.ValueOf(left)
	rightRv := reflect.ValueOf(right)
	isLeftNull := Null(leftRv)
	isRightNull := Null(rightRv)
	isBothEqual := isLeftNull == isRightNull

	if isLeftNull && isBothEqual {
		// both null
		return true, true
	} else if !isBothEqual && isLeftNull || isRightNull {
		// any null but the other is not
		return false, true
	}

	if leftRv.Type() != rightRv.Type() {
		return false, true
	}

	return false, false
}
