package corestr

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/strutilinternal"
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

func (receiver *LeftMiddleRight) MiddleTrim() string {
	return strings.TrimSpace(receiver.Middle)
}

func (receiver *LeftMiddleRight) MiddleBytes() []byte {
	return []byte(receiver.Middle)
}

func (receiver *LeftMiddleRight) IsMiddleEmpty() bool {
	return receiver.Middle == ""
}

func (receiver *LeftMiddleRight) IsMiddleWhitespace() bool {
	return strutilinternal.IsEmptyOrWhitespace(receiver.Middle)
}

func (receiver *LeftMiddleRight) HasValidNonEmptyMiddle() bool {
	return receiver.IsValid && !receiver.IsMiddleEmpty()
}

func (receiver *LeftMiddleRight) HasValidNonWhitespaceMiddle() bool {
	return receiver.IsValid && !receiver.IsMiddleWhitespace()
}

// HasSafeNonEmpty receiver.IsValid &&
//		!receiver.IsLeftEmpty() &&
//		!receiver.IsMiddleEmpty() &&
//		!receiver.IsRightEmpty()
func (receiver *LeftMiddleRight) HasSafeNonEmpty() bool {
	return receiver.IsValid &&
		!receiver.IsLeftEmpty() &&
		!receiver.IsMiddleEmpty() &&
		!receiver.IsRightEmpty()
}

func (receiver *LeftMiddleRight) IsAll(left, mid, right string) bool {
	return receiver.Left == left &&
		receiver.Right == right &&
		receiver.Middle == mid
}

func (receiver *LeftMiddleRight) Clone() *LeftMiddleRight {
	return &LeftMiddleRight{
		LeftRight: *receiver.LeftRight.Clone(),
		Middle:    receiver.Middle,
	}
}
