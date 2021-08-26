package chmodhelper

import (
	"io/fs"
	"path/filepath"
)

func RecursivePathsApply(
	rootPath string,
	actionFunc func(currentPath string, info fs.FileInfo, err error) error,
) error {
	return filepath.Walk(rootPath, actionFunc)
}
