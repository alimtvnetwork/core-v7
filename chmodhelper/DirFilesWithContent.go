package chmodhelper

import (
	"errors"
	"os"
	"path"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/pathinternal"
)

type DirFilesWithContent struct {
	Dir         string
	Files       []FileWithContent
	DirFileMode os.FileMode
}

func (it *DirFilesWithContent) GetPaths() []string {
	collection := corestr.New.Collection.Cap(
		constants.ArbitraryCapacity50,
	)

	for _, file := range it.Files {
		compiledPath := path.Join(it.Dir, file.RelativePath)
		collection.Add(compiledPath)
	}

	return collection.List()
}

func (it *DirFilesWithContent) GetFilesChmodMap() *corestr.Hashmap {
	files := it.GetPaths()

	hashmap, err := GetFilesChmodRwxFullMap(files)

	errcore.SimpleHandleErr(
		err,
		"GetFilesChmodMap() failed to retrieve hashmap from file paths",
	)

	return hashmap
}

func (it *DirFilesWithContent) Create(
	isRemoveBeforeCreate bool,
) error {
	var sliceErr []string
	fileCreatorFunc := fileWriter{}.Bytes.Chmod

	removeDirErr := removeDirIf(
		isRemoveBeforeCreate,
		it.Dir,
		"Create",
	)

	if removeDirErr != nil {
		sliceErr = append(
			sliceErr,
			removeDirErr.Error(),
		)
	}

	for _, file := range it.Files {
		compiledPath := pathinternal.Join(
			it.Dir,
			file.RelativePath,
		)

		err := fileCreatorFunc(
			isRemoveBeforeCreate,
			it.DirFileMode,
			file.FileMode,
			compiledPath,
			file.ContentToBytes(),
		)

		if err != nil {
			sliceErr = append(
				sliceErr,
				err.Error(),
			)
		}
	}

	if len(sliceErr) == 0 {
		return nil
	}

	// has error
	toString := strings.Join(sliceErr, constants.NewLineUnix)

	return errors.New(toString)
}
