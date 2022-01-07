package corejson

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
)

var (
	NewResult               = newResultCreator{}
	NewResultsCollection    = newResultsCollectionCreator{}
	NewBytesCollection      = newBytesCollectionCreator{}
	NewResultsPtrCollection = newResultsPtrCollectionCreator{}
	NewMapResults           = newMapResultsCreator{}
	Empty                   = emptyCreator{}
	StaticJsonError         = errcore.EmptyResultCannotMakeJsonType.
				Error(constants.EmptyString, constants.EmptyString)
)
