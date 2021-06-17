package chmodhelpertestwrappers

import (
	"path"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/msgtype"
)

type CreatePathsInstruction struct {
	Dir      string
	Files    []string
	ApplyRwx chmodins.RwxOwnerGroupOther
}

func (receiver *CreatePathsInstruction) GetPaths() []string {
	return *receiver.GetPathsPtr()
}

func (receiver *CreatePathsInstruction) GetPathsPtr() *[]string {
	collection := corestr.NewCollection(constants.ArbitraryCapacity50)

	for _, file := range receiver.Files {
		compiledPath := path.Join(receiver.Dir, file)
		collection.Add(compiledPath)
	}

	return collection.ListPtr()
}

func (receiver *CreatePathsInstruction) GetFilesChmodMap() *corestr.Hashmap {
	files := receiver.GetPathsPtr()

	hashmap, err := chmodhelper.GetFilesChmodRwxFullMap(*files)

	msgtype.SimpleHandleErr(
		err,
		"GetFilesChmodMap() failed to retrive hashmap from file paths")

	return hashmap
}
