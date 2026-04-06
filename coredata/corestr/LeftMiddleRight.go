package corestr

import (
	"strings"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/internal/strutilinternal"
)

// LeftMiddleRight is a string-specialized three-value container.
// It embeds coregeneric.Triple[string, string, string] for core logic
// and adds string-specific convenience methods.
type LeftMiddleRight struct {
	coregeneric.Triple[string, string, string]
}

func InvalidLeftMiddleRightNoMessage() *LeftMiddleRight {
	return &LeftMiddleRight{
		Triple: *coregeneric.InvalidTripleNoMessage[string, string, string](),
	}
}

func InvalidLeftMiddleRight(message string) *LeftMiddleRight {
	return &LeftMiddleRight{
		Triple: *coregeneric.InvalidTriple[string, string, string](message),
	}
}

// NewLeftMiddleRight creates a valid LeftMiddleRight from three strings.
func NewLeftMiddleRight(left, middle, right string) *LeftMiddleRight {
	return &LeftMiddleRight{
		Triple: *coregeneric.NewTriple(left, middle, right),
	}
}

// --- String-specific methods ---

func (it *LeftMiddleRight) LeftBytes() []byte {
	if it == nil {
		return nil
	}

	return []byte(it.Left)
}

func (it *LeftMiddleRight) RightBytes() []byte {
	if it == nil {
		return nil
	}

	return []byte(it.Right)
}

func (it *LeftMiddleRight) MiddleBytes() []byte {
	if it == nil {
		return nil
	}

	return []byte(it.Middle)
}

func (it *LeftMiddleRight) LeftTrim() string {
	if it == nil {
		return ""
	}

	return strings.TrimSpace(it.Left)
}

func (it *LeftMiddleRight) RightTrim() string {
	if it == nil {
		return ""
	}

	return strings.TrimSpace(it.Right)
}

func (it *LeftMiddleRight) MiddleTrim() string {
	if it == nil {
		return ""
	}

	return strings.TrimSpace(it.Middle)
}

func (it *LeftMiddleRight) IsLeftEmpty() bool {
	return it == nil || it.Left == ""
}

func (it *LeftMiddleRight) IsRightEmpty() bool {
	return it == nil || it.Right == ""
}

func (it *LeftMiddleRight) IsMiddleEmpty() bool {
	return it == nil || it.Middle == ""
}

func (it *LeftMiddleRight) IsMiddleWhitespace() bool {
	return it == nil || strutilinternal.IsEmptyOrWhitespace(it.Middle)
}

func (it *LeftMiddleRight) IsLeftWhitespace() bool {
	return it == nil || strutilinternal.IsEmptyOrWhitespace(it.Left)
}

func (it *LeftMiddleRight) IsRightWhitespace() bool {
	return it == nil || strutilinternal.IsEmptyOrWhitespace(it.Right)
}

func (it *LeftMiddleRight) HasValidNonEmptyLeft() bool {
	return it != nil && it.IsValid && !it.IsLeftEmpty()
}

func (it *LeftMiddleRight) HasValidNonEmptyRight() bool {
	return it != nil && it.IsValid && !it.IsRightEmpty()
}

func (it *LeftMiddleRight) HasValidNonEmptyMiddle() bool {
	return it != nil && it.IsValid && !it.IsMiddleEmpty()
}

func (it *LeftMiddleRight) HasValidNonWhitespaceLeft() bool {
	return it != nil && it.IsValid && !it.IsLeftWhitespace()
}

func (it *LeftMiddleRight) HasValidNonWhitespaceRight() bool {
	return it != nil && it.IsValid && !it.IsRightWhitespace()
}

func (it *LeftMiddleRight) HasValidNonWhitespaceMiddle() bool {
	return it != nil && it.IsValid && !it.IsMiddleWhitespace()
}

// HasSafeNonEmpty receiver.IsValid &&
//
//	!receiver.IsLeftEmpty() &&
//	!receiver.IsMiddleEmpty() &&
//	!receiver.IsRightEmpty()
func (it *LeftMiddleRight) HasSafeNonEmpty() bool {
	return it != nil &&
		it.IsValid &&
		!it.IsLeftEmpty() &&
		!it.IsMiddleEmpty() &&
		!it.IsRightEmpty()
}

func (it *LeftMiddleRight) IsAll(left, mid, right string) bool {
	return it != nil &&
		it.Left == left &&
		it.Right == right &&
		it.Middle == mid
}

func (it *LeftMiddleRight) Is(left, right string) bool {
	return it != nil && it.Left == left && it.Right == right
}

func (it *LeftMiddleRight) ToLeftRight() *LeftRight {
	if it == nil {
		return nil
	}

	return &LeftRight{
		Pair: coregeneric.Pair[string, string]{
			Left:    it.Left,
			Right:   it.Right,
			IsValid: it.IsValid,
			Message: it.Message,
		},
	}
}

func (it *LeftMiddleRight) Clear() {
	if it == nil {
		return
	}

	it.Triple.Clear()
}

func (it *LeftMiddleRight) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
}
