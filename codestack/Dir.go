package codestack

import (
	"path/filepath"
	"runtime"

	"gitlab.com/evatix-go/core/constants"
)

func Dir(skipStack int) string {
	_, filePath, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if isOkay {
		return filepath.Dir(filePath)
	}

	return constants.EmptyString
}
