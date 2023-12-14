package reflectinternal

import (
	"path"
	"path/filepath"
	"runtime"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/pathinternal"
)

type reflectPath struct{}

func (it reflectPath) RepoDir() string {
	if repoDir != nil {
		return *repoDir
	}

	repoDirLocal := it.repoDirRecursive(
		it.CurDir(),
		20,
	)

	repoDir = &repoDirLocal

	return repoDirLocal
}

func (it reflectPath) repoDirRecursive(curDir string, try int) string {
	if curDir == "" {
		return ""
	}

	modFile := path.Join(curDir, "go.mod")
	gitFolder := path.Join(curDir, ".git")

	if it.isPathsExists(modFile, gitFolder) {
		return curDir
	}

	if it.isPathsExists(gitFolder) {
		return curDir
	}

	if try <= 0 {
		return curDir
	}

	parentDir := it.parentDir(curDir)

	return it.repoDirRecursive(parentDir, try-1)
}

func (it reflectPath) isPathsExists(multiPaths ...string) bool {
	if len(multiPaths) == 0 {
		return false
	}

	for _, singlePath := range multiPaths {
		if !pathinternal.IsPathExists(singlePath) {
			return false
		}
	}

	return true
}

func (it reflectPath) parentDir(curPath string) string {
	return filepath.Dir(curPath)
}

func (it reflectPath) pathName(curPath string) string {
	_, name := filepath.Split(curPath)

	return name
}

func (it reflectPath) CurDir() string {
	return it.CurDirSkipStack(defaultInternalSkip)
}

func (it reflectPath) CurDirSkipStack(skipStack int) string {
	_, filePath, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if isOkay {
		return filepath.Dir(filePath)
	}

	return constants.EmptyString
}
