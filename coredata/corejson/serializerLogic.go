package corejson

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type serializerLogic struct{}

func (it serializerLogic) StringsApply(
	slice []string,
) *Result {
	return it.Apply(slice)
}

func (it serializerLogic) Apply(
	anyItem interface{},
) *Result {
	jsonBytes, err := json.Marshal(
		anyItem)
	typeName := reflectinternal.TypeName(
		anyItem)

	if err != nil {
		return &Result{
			Bytes: jsonBytes,
			Error: errcore.
				MarshallingFailedType.Error(
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

func (it serializerLogic) ApplyMust(
	anyItem interface{},
) *Result {
	result := it.Apply(anyItem)
	result.MustSafe()

	return result
}

func (it serializerLogic) BytesMust(
	anyItem interface{},
) []byte {
	result := it.Apply(anyItem)
	result.MustSafe()

	return result.Bytes
}

func (it serializerLogic) SafeBytesMust(
	anyItem interface{},
) []byte {
	result := it.Apply(anyItem)
	result.MustSafe()

	return result.SafeBytes()
}

// SafeBytes
//
// Warning or Danger:
//  - shallow err by not throwing or returning (could be dangerous as well)
//
// Notes :
//  - To inform use Err or Apply or must methods
//
// Use case (rarely):
//  - When don't care about the error just proceed with the value.
func (it serializerLogic) SafeBytes(
	anyItem interface{},
) []byte {
	result := it.Apply(anyItem)

	return result.SafeBytes()
}

// Bytes
//
// Warning or Danger:
//  - shallow err by not throwing or returning (could be dangerous as well)
//
// Notes :
//  - To inform use Err or Apply or must methods
//
// Use case (rarely):
//  - When don't care about the error just proceed with the value.
func (it serializerLogic) Bytes(
	anyItem interface{},
) []byte {
	result := it.Apply(anyItem)

	return result.Bytes
}

func (it serializerLogic) BytesErr(
	anyItem interface{},
) ([]byte, error) {
	result := it.Apply(anyItem)

	return result.Bytes, result.MeaningfulError()
}

// String
//
// Warning:
//  - Shallow err by not throwing or
//      returning (could be dangerous as well)
//  - However, with this version
//      if error occurred then error will be returned as string.
//
// Notes :
//  - To inform use Err or Apply or must methods
//
// Use case (rarely):
//  - When don't care about the error just proceed with the value.
func (it serializerLogic) String(
	anyItem interface{},
) string {
	result := it.Apply(anyItem)

	return result.JsonString()
}

func (it serializerLogic) StringMust(
	anyItem interface{},
) string {
	result := it.Apply(anyItem)
	result.HandleError()

	return result.JsonString()
}

func (it serializerLogic) StringErr(
	anyItem interface{},
) (string, error) {
	result := it.Apply(anyItem)

	return result.RawString()
}

func (it serializerLogic) PrettyStringErr(
	anyItem interface{},
) (string, error) {
	result := it.Apply(anyItem)

	return result.RawPrettyString()
}

// PrettyStringIncludingErr
//
// Warning:
//  - Shallow err by not throwing or
//      returning (could be dangerous as well)
//  - However, with this version
//      if error occurred then error will be returned as string.
//
// Notes :
//  - To inform use Err or Apply or must methods
//
// Use case (rarely):
//  - When don't care about the error just proceed with the value.
func (it serializerLogic) PrettyStringIncludingErr(
	anyItem interface{},
) string {
	result := it.Apply(anyItem)

	return result.PrettyJsonStringWithErr()
}
