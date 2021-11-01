package codestack

import (
	"runtime"

	"gitlab.com/evatix-go/core/constants"
)

func CurrentFilePath() string {
	_, filePath, _, isOkay := runtime.Caller(defaultInternalSkip)

	if isOkay {
		return filePath
	}

	return constants.EmptyString
}
