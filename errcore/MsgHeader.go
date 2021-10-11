package errcore

import (
	"fmt"

	"gitlab.com/evatix-go/core/internal/msgformats"
)

func MsgHeader(
	items ...interface{},
) string {
	return fmt.Sprintf(
		msgformats.MsgHeaderFormat,
		items...)
}
