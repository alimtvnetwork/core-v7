package codestack

import (
	"path/filepath"
	"runtime"

	"gitlab.com/auk-go/core/constants"
)

func CurDir() string {
	_, filePath, _, isOkay := runtime.Caller(defaultInternalSkip)

	if isOkay {
		return filepath.Dir(filePath)
	}

	return constants.EmptyString
}
