package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

func MsgHeader(
	items ...interface{},
) string {
	return fmt.Sprintf(
		msgformats.MsgHeaderFormat,
		items...)
}
