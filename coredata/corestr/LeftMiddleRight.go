package corestr

import "gitlab.com/evatix-go/core/internal/stringutil"

type LeftMiddleRight struct {
	LeftRight
	Middle string
}

func (receiver *LeftMiddleRight) MiddleBytes() []byte {
	return []byte(receiver.Middle)
}

func (receiver *LeftMiddleRight) IsMiddleEmpty() bool {
	return receiver.Middle == ""
}

func (receiver *LeftMiddleRight) IsMiddleWhitespace() bool {
	return stringutil.IsEmptyOrWhitespace(receiver.Middle)
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
