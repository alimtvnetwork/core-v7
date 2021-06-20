package corestr

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/strutilinternal"
)

type LeftRight struct {
	Left, Right string
	IsValid     bool
	Message     string
}

func InvalidLeftRightNoMessage() *LeftRight {
	return &LeftRight{
		Left:    constants.EmptyString,
		Right:   constants.EmptyString,
		IsValid: false,
		Message: constants.EmptyString,
	}
}

func InvalidLeftRight(message string) *LeftRight {
	return &LeftRight{
		Left:    constants.EmptyString,
		Right:   constants.EmptyString,
		IsValid: false,
		Message: message,
	}
}

func LeftRightUsingSlicePtr(slice *[]string) *LeftRight {
	if slice == nil || *slice == nil {
		return LeftRightUsingSlice(nil)
	}

	return LeftRightUsingSlice(*slice)
}

func LeftRightUsingSlice(slice []string) *LeftRight {
	length := len(slice)

	if length == 0 {
		return InvalidLeftRight(
			LeftRightExpectingLengthMessager.Message(
				length))
	}

	return &LeftRight{
		Left:    slice[constants.Zero],
		Right:   slice[length-1],
		IsValid: length == ExpectingLengthForLeftRight,
	}
}

func (receiver *LeftRight) LeftBytes() []byte {
	return []byte(receiver.Left)
}

func (receiver *LeftRight) RightBytes() []byte {
	return []byte(receiver.Right)
}

func (receiver *LeftRight) IsLeftEmpty() bool {
	return receiver.Left == ""
}

func (receiver *LeftRight) IsRightEmpty() bool {
	return receiver.Right == ""
}

func (receiver *LeftRight) IsRightWhitespace() bool {
	return strutilinternal.IsEmptyOrWhitespace(receiver.Right)
}

func (receiver *LeftRight) IsLeftWhitespace() bool {
	return strutilinternal.IsEmptyOrWhitespace(receiver.Left)
}

func (receiver *LeftRight) HasValidNonEmptyLeft() bool {
	return receiver.IsValid && !receiver.IsLeftEmpty()
}

func (receiver *LeftRight) HasValidNonEmptyRight() bool {
	return receiver.IsValid && !receiver.IsRightEmpty()
}

func (receiver *LeftRight) HasValidNonWhitespaceLeft() bool {
	return receiver.IsValid && !receiver.IsLeftWhitespace()
}

func (receiver *LeftRight) HasValidNonWhitespaceRight() bool {
	return receiver.IsValid && !receiver.IsRightWhitespace()
}

// HasSafeNonEmpty receiver.IsValid &&
//		!receiver.IsLeftEmpty() &&
//		!receiver.IsRightEmpty()
func (receiver *LeftRight) HasSafeNonEmpty() bool {
	return receiver.IsValid &&
		!receiver.IsLeftEmpty() &&
		!receiver.IsRightEmpty()
}

func (receiver *LeftRight) IsLeft(left string) bool {
	return receiver.Left == left
}

func (receiver *LeftRight) IsRight(right string) bool {
	return receiver.Right == right
}

func (receiver *LeftRight) Is(left, right string) bool {
	return receiver.Left == left && receiver.Right == right
}
