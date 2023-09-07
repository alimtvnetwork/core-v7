package codestack

import (
	"path/filepath"
	"runtime"

	"gitlab.com/auk-go/core/constants"
)

func FileName(skipStack int) string {
	_, file, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if !isOkay && file == "" {
		return constants.EmptyString
	}

	_, fileName := filepath.Split(file)

	return fileName
}
