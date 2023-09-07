package chmodhelper

import (
	"io/fs"
	"path/filepath"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

func GetRecursivePathsContinueOnError(
	rootPath string,
) ([]string, error) {
	stat := GetPathExistStat(rootPath)

	if !stat.IsExist {
		return []string{}, errcore.PathsMissingOrHavingIssuesType.
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
					err.Error()+constants.HyphenAngelRight+path)

				return nil
			}

			allPaths = append(allPaths, path)

			return nil
		})

	if finalErr != nil {
		sliceErr = append(
			sliceErr,
			finalErr.Error()+constants.HyphenAngelRight+rootPath)
	}

	return allPaths, errcore.SliceToError(sliceErr)
}
