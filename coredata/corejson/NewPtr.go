package corejson

import (
	"encoding/json"

	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func NewPtr(anyItem interface{}) *Result {
	jsonBytes, err := json.Marshal(anyItem)
	typeName := reflectinternal.TypeName(anyItem)

	if err != nil {
		return &Result{
			Bytes: jsonBytes,
			Error: errcore.MarshallingFailedType.Error(
				err.Error(),
				typeName),
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}
