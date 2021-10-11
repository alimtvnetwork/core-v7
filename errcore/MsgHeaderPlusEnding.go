package errcore

import (
	"fmt"

	"gitlab.com/evatix-go/core/internal/msgformats"
)

func MsgHeaderPlusEnding(
	header, message interface{},
) string {
	return fmt.Sprintf(
		msgformats.MsgHeaderPlusEndingFormat,
		header,
		message)
}
