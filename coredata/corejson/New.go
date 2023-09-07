package corejson

import (
	"encoding/json"

	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func New(anyItem interface{}) Result {
	jsonBytes, err := json.Marshal(anyItem)
	typeName := reflectinternal.TypeName(anyItem)

	if err != nil {
		nextErr := errcore.
			MarshallingFailedType.
			Error(
				err.Error(),
				typeName)

		return Result{
			Bytes:    jsonBytes,
			Error:    nextErr,
			TypeName: typeName,
		}
	}

	return Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}
