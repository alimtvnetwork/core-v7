package errcore

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func ErrorToSplitLines(err error) []string {
	if err == nil {
		return []string{}
	}

	return strings.Split(
		err.Error(),
		constants.NewLineUnix)
}
