package coredynamic

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type LeftRight struct {
	Left, Right interface{}
}

func (it *LeftRight) IsEmpty() bool {
	return it == nil ||
		reflectinternal.Is.Null(it.Left) &&
			reflectinternal.Is.Null(it.Right)
}

func (it *LeftRight) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *LeftRight) HasLeft() bool {
	return it != nil &&
		reflectinternal.Is.Null(it.Left)
}

func (it *LeftRight) HasRight() bool {
	return it != nil &&
		reflectinternal.Is.Null(it.Right)
}

func (it *LeftRight) IsLeftEmpty() bool {
	return it == nil ||
		reflectinternal.Is.Null(it.Left)
}

func (it *LeftRight) IsRightEmpty() bool {
	return it == nil ||
		reflectinternal.Is.Null(it.Right)
}

func (it *LeftRight) LeftReflectSet(
	toPointerOrBytesPointer interface{},
) error {
	return ReflectSetFromTo(it.Left, toPointerOrBytesPointer)
}

func (it *LeftRight) RightReflectSet(
	toPointerOrBytesPointer interface{},
) error {
	return ReflectSetFromTo(it.Right, toPointerOrBytesPointer)
}

func (it *LeftRight) DeserializeLeft() *corejson.Result {
	return corejson.NewPtr(it.Left)
}

func (it *LeftRight) DeserializeRight() *corejson.Result {
	return corejson.NewPtr(it.Right)
}

func (it *LeftRight) LeftToDynamic() *Dynamic {
	return NewDynamicPtr(it.Left, true)
}

func (it *LeftRight) RightToDynamic() *Dynamic {
	return NewDynamicPtr(it.Right, true)
}

func (it *LeftRight) TypeStatus() TypeStatus {
	return TypeSameStatus(it.Left, it.Right)
}
