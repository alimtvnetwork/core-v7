package chmodhelper

import (
	"os"
	"path"

	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/errcore"
)

type DirFilesWithRwxPermission struct {
	DirWithFiles
	ApplyRwx chmodins.RwxOwnerGroupOther
}

func (it *DirFilesWithRwxPermission) GetPaths() []string {
	collection := corestr.New.Collection.Cap(constants.ArbitraryCapacity50)

	for _, file := range it.Files {
		compiledPath := path.Join(it.Dir, file)
		collection.Add(compiledPath)
	}

	return collection.List()
}

func (it *DirFilesWithRwxPermission) GetFilesChmodMap() *corestr.Hashmap {
	files := it.GetPaths()

	hashmap, err := GetFilesChmodRwxFullMap(files)

	errcore.SimpleHandleErr(
		err,
		"GetFilesChmodMap() failed to retrieve hashmap from file paths",
	)

	return hashmap
}

func (it *DirFilesWithRwxPermission) CreatePaths(
	isRemoveBeforeCreate bool,
) error {
	return CreateDirFilesWithRwxPermission(
		isRemoveBeforeCreate,
		it,
	)
}

func (it *DirFilesWithRwxPermission) CreateUsingFileMode(
	isRemoveBeforeCreate bool,
	fileMode os.FileMode,
) error {
	return it.DirWithFiles.CreatePaths(
		isRemoveBeforeCreate,
		fileMode,
	)
}
