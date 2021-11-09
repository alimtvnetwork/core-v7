package corestr

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
)

//goland:noinspection ALL
var (
	New = &newCreator{
		Collection:                 &newCollectionCreator{},
		CollectionPtr:              &newCollectionPtrCreator{},
		CharHashsetMap:             &newCharHashsetMapCreator{},
		CharCollectionMap:          &newCharCollectionMapCreator{},
		SimpleStringOnce:           &newSimpleStringOnceCreator{},
		SimpleSlice:                &newSimpleSliceCreator{},
		Hashset:                    &newHashsetCreator{},
		HashsetsCollection:         &newHashsetsCollectionCreator{},
		Hashmap:                    &newHashmapCreator{},
		LinkedList:                 &newLinkedListCreator{},
		LinkedCollection:           &newLinkedListCollectionsCreator{},
		CollectionsOfCollection:    &newCollectionsOfCollectionCreator{},
		CollectionsOfCollectionPtr: &newCollectionsOfCollectionPtrCreator{},
	}

	Empty           = &emptyCreator{}
	StaticJsonError = errcore.EmptyResultCannotMakeJsonType.
			Error(constants.EmptyString, constants.EmptyString)
	ExpectingLengthForLeftRight      = constants.Two
	LeftRightExpectingLengthMessager = errcore.ExpectingFuture(
		"Expecting length at least",
		ExpectingLengthForLeftRight)
)
