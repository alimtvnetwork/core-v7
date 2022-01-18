package isany

import (
	"reflect"
)

func TypeSame(
	left, right interface{},
) bool {
	leftRt := reflect.TypeOf(left)
	rightRt := reflect.TypeOf(right)

	return leftRt == rightRt
}
