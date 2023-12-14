package corestr

import (
	"regexp"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/strutilinternal"
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

func LeftRightTrimmedUsingSlice(slice []string) *LeftRight {
	if slice == nil {
		return LeftRightUsingSlice(nil)
	}

	length := len(slice)

	if length == 0 {
		return InvalidLeftRight(
			LeftRightExpectingLengthMessager.Message(
				length))
	}

	if length == 1 {
		return &LeftRight{
			Left:    slice[constants.Zero],
			Right:   constants.EmptyString,
			IsValid: length == ExpectingLengthForLeftRight,
		}
	}

	return &LeftRight{
		Left:    strings.TrimSpace(slice[constants.Zero]),
		Right:   strings.TrimSpace(slice[length-1]),
		IsValid: length == ExpectingLengthForLeftRight,
	}
}

func LeftRightUsingSlice(slice []string) *LeftRight {
	length := len(slice)

	if length == 0 {
		return InvalidLeftRight(
			LeftRightExpectingLengthMessager.Message(
				length))
	}

	if length == 1 {
		return &LeftRight{
			Left:    slice[constants.Zero],
			Right:   constants.EmptyString,
			IsValid: length == ExpectingLengthForLeftRight,
		}
	}

	return &LeftRight{
		Left:    slice[constants.Zero],
		Right:   slice[length-1],
		IsValid: length == ExpectingLengthForLeftRight,
	}
}

func (it *LeftRight) LeftBytes() []byte {
	return []byte(it.Left)
}

func (it *LeftRight) RightBytes() []byte {
	return []byte(it.Right)
}

func (it *LeftRight) LeftTrim() string {
	return strings.TrimSpace(it.Left)
}

func (it *LeftRight) RightTrim() string {
	return strings.TrimSpace(it.Right)
}

func (it *LeftRight) IsLeftEmpty() bool {
	return it.Left == ""
}

func (it *LeftRight) IsRightEmpty() bool {
	return it.Right == ""
}

func (it *LeftRight) IsRightWhitespace() bool {
	return strutilinternal.IsEmptyOrWhitespace(it.Right)
}

func (it *LeftRight) IsLeftWhitespace() bool {
	return strutilinternal.IsEmptyOrWhitespace(it.Left)
}

func (it *LeftRight) HasValidNonEmptyLeft() bool {
	return it.IsValid && !it.IsLeftEmpty()
}

func (it *LeftRight) HasValidNonEmptyRight() bool {
	return it.IsValid && !it.IsRightEmpty()
}

func (it *LeftRight) HasValidNonWhitespaceLeft() bool {
	return it.IsValid && !it.IsLeftWhitespace()
}

func (it *LeftRight) HasValidNonWhitespaceRight() bool {
	return it.IsValid && !it.IsRightWhitespace()
}

// HasSafeNonEmpty receiver.IsValid &&
//
//	!receiver.IsLeftEmpty() &&
//	!receiver.IsRightEmpty()
func (it *LeftRight) HasSafeNonEmpty() bool {
	return it.IsValid &&
		!it.IsLeftEmpty() &&
		!it.IsRightEmpty()
}

func (it LeftRight) NonPtr() LeftRight {
	return it
}

func (it *LeftRight) Ptr() *LeftRight {
	return it
}

func (it *LeftRight) IsLeftRegexMatch(regexp *regexp.Regexp) bool {
	if regexp == nil {
		return false
	}

	return regexp.MatchString(it.Left)
}

func (it *LeftRight) IsRightRegexMatch(regexp *regexp.Regexp) bool {
	if regexp == nil {
		return false
	}

	return regexp.MatchString(it.Right)
}

func (it *LeftRight) IsLeft(left string) bool {
	return it.Left == left
}

func (it *LeftRight) IsRight(right string) bool {
	return it.Right == right
}

func (it *LeftRight) Is(left, right string) bool {
	return it.Left == left && it.Right == right
}

func (it *LeftRight) IsEqual(another *LeftRight) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	return it.IsValid == another.IsValid &&
		it.Is(another.Left, another.Right)
}

func (it *LeftRight) Clone() *LeftRight {
	return &LeftRight{
		Left:    it.Left,
		Right:   it.Right,
		IsValid: it.IsValid,
		Message: it.Message,
	}
}

func (it *LeftRight) Clear() {
	if it == nil {
		return
	}

	it.Left = ""
	it.Right = ""
	it.IsValid = false
	it.Message = ""
}

func (it *LeftRight) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
}
