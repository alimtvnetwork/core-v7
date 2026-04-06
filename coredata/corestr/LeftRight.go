package corestr

import (
	"regexp"
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/internal/strutilinternal"
)

// LeftRight is a string-specialized two-value container.
// It embeds coregeneric.Pair[string, string] for core logic
// and adds string-specific convenience methods.
type LeftRight struct {
	coregeneric.Pair[string, string]
}

func InvalidLeftRightNoMessage() *LeftRight {
	return &LeftRight{
		Pair: *coregeneric.InvalidPairNoMessage[string, string](),
	}
}

func InvalidLeftRight(message string) *LeftRight {
	return &LeftRight{
		Pair: *coregeneric.InvalidPair[string, string](message),
	}
}

func LeftRightUsingSlicePtr(slice []string) *LeftRight {
	if len(slice) == 0 {
		return LeftRightUsingSlice(nil)
	}

	return LeftRightUsingSlice(slice)
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
			Pair: coregeneric.Pair[string, string]{
				Left:    strings.TrimSpace(slice[constants.Zero]),
				Right:   constants.EmptyString,
				IsValid: length == ExpectingLengthForLeftRight,
			},
		}
	}

	return &LeftRight{
		Pair: coregeneric.Pair[string, string]{
			Left:    strings.TrimSpace(slice[constants.Zero]),
			Right:   strings.TrimSpace(slice[length-1]),
			IsValid: length == ExpectingLengthForLeftRight,
		},
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
			Pair: coregeneric.Pair[string, string]{
				Left:    slice[constants.Zero],
				Right:   constants.EmptyString,
				IsValid: length == ExpectingLengthForLeftRight,
			},
		}
	}

	return &LeftRight{
		Pair: coregeneric.Pair[string, string]{
			Left:    slice[constants.Zero],
			Right:   slice[length-1],
			IsValid: length == ExpectingLengthForLeftRight,
		},
	}
}

// NewLeftRight creates a valid LeftRight from two strings.
func NewLeftRight(left, right string) *LeftRight {
	return &LeftRight{
		Pair: *coregeneric.NewPair(left, right),
	}
}

// --- String-specific methods ---

func (it *LeftRight) LeftBytes() []byte {
	if it == nil {
		return nil
	}

	return []byte(it.Left)
}

func (it *LeftRight) RightBytes() []byte {
	if it == nil {
		return nil
	}

	return []byte(it.Right)
}

func (it *LeftRight) LeftTrim() string {
	if it == nil {
		return ""
	}

	return strings.TrimSpace(it.Left)
}

func (it *LeftRight) RightTrim() string {
	if it == nil {
		return ""
	}

	return strings.TrimSpace(it.Right)
}

func (it *LeftRight) IsLeftEmpty() bool {
	return it == nil || it.Left == ""
}

func (it *LeftRight) IsRightEmpty() bool {
	return it == nil || it.Right == ""
}

func (it *LeftRight) IsRightWhitespace() bool {
	return it == nil || strutilinternal.IsEmptyOrWhitespace(it.Right)
}

func (it *LeftRight) IsLeftWhitespace() bool {
	return it == nil || strutilinternal.IsEmptyOrWhitespace(it.Left)
}

func (it *LeftRight) HasValidNonEmptyLeft() bool {
	return it != nil && it.IsValid && !it.IsLeftEmpty()
}

func (it *LeftRight) HasValidNonEmptyRight() bool {
	return it != nil && it.IsValid && !it.IsRightEmpty()
}

func (it *LeftRight) HasValidNonWhitespaceLeft() bool {
	return it != nil && it.IsValid && !it.IsLeftWhitespace()
}

func (it *LeftRight) HasValidNonWhitespaceRight() bool {
	return it != nil && it.IsValid && !it.IsRightWhitespace()
}

// HasSafeNonEmpty receiver.IsValid &&
//
//	!receiver.IsLeftEmpty() &&
//	!receiver.IsRightEmpty()
func (it *LeftRight) HasSafeNonEmpty() bool {
	return it != nil &&
		it.IsValid &&
		!it.IsLeftEmpty() &&
		!it.IsRightEmpty()
}

func (it *LeftRight) NonPtr() LeftRight {
	if it == nil {
		return LeftRight{}
	}

	return *it
}

func (it *LeftRight) Ptr() *LeftRight {
	return it
}

func (it *LeftRight) IsLeftRegexMatch(regexp *regexp.Regexp) bool {
	if it == nil || regexp == nil {
		return false
	}

	return regexp.MatchString(it.Left)
}

func (it *LeftRight) IsRightRegexMatch(regexp *regexp.Regexp) bool {
	if it == nil || regexp == nil {
		return false
	}

	return regexp.MatchString(it.Right)
}

func (it *LeftRight) IsLeft(left string) bool {
	return it != nil && it.Left == left
}

func (it *LeftRight) IsRight(right string) bool {
	return it != nil && it.Right == right
}

func (it *LeftRight) Is(left, right string) bool {
	return it != nil && it.Left == left && it.Right == right
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
		Pair: *it.Pair.Clone(),
	}
}

func (it *LeftRight) Clear() {
	if it == nil {
		return
	}

	it.Pair.Clear()
}

func (it *LeftRight) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
}
