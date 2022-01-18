package coredynamic

import (
	"errors"
	"reflect"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
	"gitlab.com/evatix-go/core/issetter"
)

type TypeStatus struct {
	IsSame                                bool
	IsLeftUnknownNull, IsRightUnknownNull bool
	IsRightPointer, IsLeftPointer         bool
	Left, Right                           reflect.Type
	isValid                               issetter.Value
}

func (it *TypeStatus) IsValid() bool {
	if it == nil {
		return false
	}

	if it.isValid.IsInitBoolean() {
		return it.isValid.IsTrue()
	}

	it.isValid = issetter.GetBool(
		it != nil &&
			!reflectinternal.IsNull(it.Left) &&
			!reflectinternal.IsNull(it.Right))

	return it.isValid.IsTrue()
}

func (it *TypeStatus) IsInvalid() bool {
	if it == nil {
		return true
	}

	return !it.IsValid()
}

func (it TypeStatus) IsNotSame() bool {
	return !it.IsSame
}

func (it TypeStatus) IsNotEqualTypes() bool {
	return !it.IsSame
}

func (it TypeStatus) IsAnyPointer() bool {
	return it.IsLeftPointer || it.IsRightPointer
}

func (it TypeStatus) IsBothPointer() bool {
	return it.IsLeftPointer && it.IsRightPointer
}

func (it TypeStatus) NonPointerLeft() reflect.Type {
	if it.IsLeftPointer || it.Left.Kind() == reflect.Interface {
		return it.Left.Elem()
	}

	return it.Left
}

func (it TypeStatus) NonPointerRight() reflect.Type {
	if it.IsRightPointer || it.Right.Kind() == reflect.Interface {
		return it.Right.Elem()
	}

	return it.Right
}

func (it TypeStatus) IsSameRegardlessPointer() bool {
	if it.IsSame {
		return true
	}

	return it.NonPointerLeft() == it.NonPointerRight()
}

func (it TypeStatus) NotEqualSrcDestinationMessage() string {
	return it.NotMatchMessage(
		constants.SourceLower,
		constants.DestinationLower)
}

func (it TypeStatus) LeftName() string {
	if reflectinternal.IsNull(it.Left) {
		return constants.NilAngelBracket
	}

	return it.Left.Name()
}

func (it TypeStatus) RightName() string {
	if reflectinternal.IsNull(it.Right) {
		return constants.NilAngelBracket
	}

	return it.Right.Name()
}

func (it TypeStatus) LeftFullName() string {
	if reflectinternal.IsNull(it.Left) {
		return constants.NilAngelBracket
	}

	return it.Left.String()
}

func (it TypeStatus) RightFullName() string {
	if reflectinternal.IsNull(it.Right) {
		return constants.NilAngelBracket
	}

	return it.Right.String()
}

func (it TypeStatus) NotMatchMessage(
	leftName,
	rightName string,
) string {
	if it.IsSame {
		return ""
	}

	return errcore.
		TypeMismatchType.
		SrcDestination(
			"type validation failed!",
			leftName, it.LeftFullName(),
			rightName, it.RightFullName())
}

func (it TypeStatus) NotMatchErr(
	leftName,
	rightName string,
) error {
	if it.IsSame {
		return nil
	}

	return errors.New(it.NotMatchMessage(leftName, rightName))
}

func (it TypeStatus) MustBeSame() {
	if it.IsSame {
		return
	}

	panic(it.NotMatchMessage(constants.LeftLower, constants.RightLower))
}

func (it TypeStatus) SrcDestinationMustBeSame() {
	if it.IsSame {
		return
	}

	panic(it.NotMatchMessage(constants.SourceLower, constants.DestinationLower))
}

func (it TypeStatus) NotEqualSrcDestinationErr() error {
	return it.NotMatchErr(
		constants.SourceLower,
		constants.DestinationLower)
}

func (it *TypeStatus) IsEqual(next *TypeStatus) bool {
	if it == nil && next == nil {
		return true
	}

	if it == nil || next == nil {
		return false
	}

	if it.IsSame != next.IsSame {
		return false
	}

	if it.Left != next.Left {
		return false
	}

	if it.Right != next.Right {
		return false
	}

	return true
}
