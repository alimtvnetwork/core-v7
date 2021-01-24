package corestr

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

var (
	StaticEmptyCollection           = *EmptyCollection()
	StaticEmptyCollectionPtr        = &StaticEmptyCollection
	StaticEmptyCharCollectionMap    = *EmptyCharCollectionMap()
	StaticEmptyCharCollectionMapPtr = &StaticEmptyCharCollectionMap
	StaticEmptyCharHashsetMap       = *EmptyCharHashsetMap()
	StaticEmptyCharHashsetMapPtr    = &StaticEmptyCharHashsetMap
	StaticEmptyHashset              = *EmptyHashset()
	StaticEmptyHashsetPtr           = &StaticEmptyHashset
	StaticJsonError                 = msgtype.EmptyResultCannotMakeJson.
		Error(constants.EmptyString, constants.EmptyString)
)
