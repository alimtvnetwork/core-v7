package corestr

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/strutilinternal"
)

type LeftMiddleRight struct {
	LeftRight
	Middle string
}

func InvalidLeftMiddleRightNoMessage() *LeftMiddleRight {
	return &LeftMiddleRight{
		LeftRight: LeftRight{
			Left:    constants.EmptyString,
			Right:   constants.EmptyString,
			IsValid: false,
			Message: constants.EmptyString,
		},
		Middle: constants.EmptyString,
	}
}

func InvalidLeftMiddleRight(message string) *LeftMiddleRight {
	return &LeftMiddleRight{
		LeftRight: LeftRight{
			Left:    constants.EmptyString,
			Right:   constants.EmptyString,
			IsValid: false,
			Message: message,
		},
		Middle: constants.EmptyString,
	}
}

func (it *LeftMiddleRight) MiddleTrim() string {
	return strings.TrimSpace(it.Middle)
}

func (it *LeftMiddleRight) MiddleBytes() []byte {
	return []byte(it.Middle)
}

func (it *LeftMiddleRight) IsMiddleEmpty() bool {
	return it.Middle == ""
}

func (it *LeftMiddleRight) IsMiddleWhitespace() bool {
	return strutilinternal.IsEmptyOrWhitespace(it.Middle)
}

func (it *LeftMiddleRight) HasValidNonEmptyMiddle() bool {
	return it.IsValid && !it.IsMiddleEmpty()
}

func (it *LeftMiddleRight) HasValidNonWhitespaceMiddle() bool {
	return it.IsValid && !it.IsMiddleWhitespace()
}

// HasSafeNonEmpty receiver.IsValid &&
//
//	!receiver.IsLeftEmpty() &&
//	!receiver.IsMiddleEmpty() &&
//	!receiver.IsRightEmpty()
func (it *LeftMiddleRight) HasSafeNonEmpty() bool {
	return it.IsValid &&
		!it.IsLeftEmpty() &&
		!it.IsMiddleEmpty() &&
		!it.IsRightEmpty()
}

func (it *LeftMiddleRight) IsAll(left, mid, right string) bool {
	return it.Left == left &&
		it.Right == right &&
		it.Middle == mid
}

func (it *LeftMiddleRight) Clone() *LeftMiddleRight {
	return &LeftMiddleRight{
		LeftRight: *it.LeftRight.Clone(),
		Middle:    it.Middle,
	}
}

func (it *LeftMiddleRight) Clear() {
	if it == nil {
		return
	}

	it.LeftRight.Clear()
	it.Middle = ""
}

func (it *LeftMiddleRight) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
}
