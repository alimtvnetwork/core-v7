package coretests

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

func LogOnFail(
	isPass bool,
	expected, actual interface{},
) {
	if isPass {
		return
	}

	logMessage := fmt.Sprintf(msgformats.LogFormat, expected, actual)
	fmt.Println(logMessage)
}
