package corejson

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
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
	CastAny                 = castingAny{}
	AnyTo                   = anyTo{}
	StaticJsonError         = errcore.
				EmptyResultCannotMakeJsonType.
				ErrorNoRefs(constants.EmptyString)
)
