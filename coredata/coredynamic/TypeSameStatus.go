package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func TypeSameStatus(
	left, right interface{},
) TypeStatus {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	isLeftUnknownNull := reflectinternal.IsNull(leftType)
	isRightUnknownNull := reflectinternal.IsNull(rightType)

	return TypeStatus{
		IsSame:             leftType == rightType,
		IsLeftUnknownNull:  isLeftUnknownNull,
		IsRightUnknownNull: isRightUnknownNull,
		IsRightPointer:     !isRightUnknownNull && rightType.Kind() == reflect.Ptr,
		IsLeftPointer:      !isLeftUnknownNull && leftType.Kind() == reflect.Ptr,
		Left:               leftType,
		Right:              rightType,
	}
}
