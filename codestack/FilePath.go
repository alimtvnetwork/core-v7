package codestack

import (
	"runtime"

	"gitlab.com/auk-go/core/constants"
)

func FilePath(skipStack int) string {
	_, file, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if isOkay {
		return file
	}

	return constants.EmptyString
}
