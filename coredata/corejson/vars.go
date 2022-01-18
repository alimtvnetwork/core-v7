package corejson

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

var (
	resultTypeName          = reflectinternal.TypeName(Result{})
	Empty                   = emptyCreator{}
	Serialize               = serializerLogic{}   // deals with json.Marshal
	Deserialize             = deserializerLogic{} // deals with json.Unmarshal
	NewResult               = newResultCreator{}
	NewResultsCollection    = newResultsCollectionCreator{}
	NewBytesCollection      = newBytesCollectionCreator{}
	NewResultsPtrCollection = newResultsPtrCollectionCreator{}
	NewMapResults           = newMapResultsCreator{}
	StaticJsonError         = errcore.
				EmptyResultCannotMakeJsonType.
				ErrorNoRefs(constants.EmptyString)
)
