package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func MsgWrapMsg(msg, wrappedMsg string) string {
	if msg == "" && wrappedMsg == "" {
		return ""
	}

	if msg == "" && wrappedMsg != "" {
		return wrappedMsg
	}

	if msg != "" && wrappedMsg == "" {
		return msg
	}

	return fmt.Sprintf(
		constants.ValueWrapValueFormat,
		msg,
		wrappedMsg)
}
