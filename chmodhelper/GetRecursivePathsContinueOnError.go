package chmodhelper

import (
	"io/fs"
	"path/filepath"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
)

func GetRecursivePathsContinueOnError(
	rootPath string,
) ([]string, error) {
	stat := GetPathExistStat(rootPath)

	if !stat.IsExist {
		return []string{}, errcore.PathsMissingOrHavingIssues.
			ErrorRefOnly(rootPath)
	}

	if stat.IsFile() {
		return []string{rootPath}, nil
	}

	allPaths := make(
		[]string,
		0,
		constants.Capacity128)
	var sliceErr []string

	finalErr := filepath.Walk(
		rootPath,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				sliceErr = append(
					sliceErr,
					err.Error()+constants.HypenAngelRight+path)

				return nil
			}

			allPaths = append(allPaths, path)

			return nil
		})

	if finalErr != nil {
		sliceErr = append(
			sliceErr,
			finalErr.Error()+constants.HypenAngelRight+rootPath)
	}

	return allPaths, errcore.SliceToError(sliceErr)
}
