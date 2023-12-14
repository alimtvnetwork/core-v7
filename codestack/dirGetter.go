package codestack

import (
	"path"
	"path/filepath"
	"runtime"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/pathinternal"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type dirGetter struct{}

func (it dirGetter) CurDir() string {
	_, filePath, _, isOkay := runtime.Caller(defaultInternalSkip)

	if isOkay {
		return filepath.Dir(filePath)
	}

	return constants.EmptyString
}

func (it dirGetter) CurDirJoin(relPaths ...string) string {
	curDur := it.Get(Skip2)
	relJoin := path.Join(relPaths...)

	return pathinternal.Join(curDur, relJoin)
}

func (it dirGetter) RepoDir() string {
	return reflectinternal.Path.RepoDir()
}

func (it dirGetter) RepoDirJoin(relPaths ...string) string {
	relJoin := path.Join(relPaths...)

	return pathinternal.Join(it.RepoDir(), relJoin)
}

func (it dirGetter) Get(skipStack int) string {
	_, filePath, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if isOkay {
		return filepath.Dir(filePath)
	}

	return constants.EmptyString
}
