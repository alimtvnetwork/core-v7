package codestack

import (
	"path/filepath"
	"runtime"

	"gitlab.com/auk-go/core/constants"
)

type fileGetter struct{}

func (it fileGetter) Name(skipStack int) string {
	_, file, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if !isOkay && file == "" {
		return constants.EmptyString
	}

	_, fileName := filepath.Split(file)

	return fileName
}

func (it fileGetter) Path(skipStack int) string {
	_, file, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if isOkay {
		return file
	}

	return constants.EmptyString
}

func (it fileGetter) PathLineSep(skipStack int) (
	filePath string, lineNumber int,
) {
	stack := New.Create(Skip1 + skipStack)
	fileWithLine := stack.FileWithLine()
	filePath = fileWithLine.FullFilePath()
	lineNumber = fileWithLine.LineNumber()

	stack.Dispose()

	return filePath, lineNumber
}

func (it fileGetter) PathLineSepDefault() (filePath string, lineNumber int) {
	return it.PathLineSep(defaultInternalSkip)
}

func (it fileGetter) FilePathWithLineString(skipStack int) string {
	stack := New.Create(Skip1 + skipStack)
	fileWithLine := stack.FileWithLineString()
	stack.Dispose()

	return fileWithLine
}

func (it fileGetter) PathLineStringDefault() string {
	stack := New.Create(Skip1)
	fileWithLine := stack.FileWithLineString()
	stack.Dispose()

	return fileWithLine
}

func (it fileGetter) CurrentFilePath() string {
	_, filePath, _, isOkay := runtime.Caller(defaultInternalSkip)

	if isOkay {
		return filePath
	}

	return constants.EmptyString
}
