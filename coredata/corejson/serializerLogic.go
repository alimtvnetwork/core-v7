package corejson

import (
	"encoding/json"
	"fmt"

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

func (it serializerLogic) FromBytes(
	currentBytes []byte,
) *Result {
	return it.Apply(currentBytes)
}

func (it serializerLogic) FromStrings(
	lines []string,
) *Result {
	return it.Apply(lines)
}

func (it serializerLogic) FromStringsSpread(
	lines ...string,
) *Result {
	return it.Apply(lines)
}

func (it serializerLogic) FromString(
	line string,
) *Result {
	return it.Apply(line)
}

func (it serializerLogic) FromInteger(
	integer int,
) *Result {
	return it.Apply(integer)
}

func (it serializerLogic) FromInteger64(
	integer64 int,
) *Result {
	return it.Apply(integer64)
}

func (it serializerLogic) FromBool(
	isResult bool,
) *Result {
	return it.Apply(isResult)
}

func (it serializerLogic) FromIntegers(
	integers []int,
) *Result {
	return it.Apply(integers)
}

func (it serializerLogic) FromStringer(
	stringer fmt.Stringer,
) *Result {
	return it.Apply(stringer.String())
}

func (it serializerLogic) UsingAnyPtr(
	anyItem interface{},
) *Result {
	jsonBytes, err := json.Marshal(
		anyItem)
	typeName := reflectinternal.TypeName(
		anyItem)

	if err != nil {
		finalErr := errcore.
			MarshallingFailedType.Error(
			err.Error(),
			typeName)

		return &Result{
			Bytes:    jsonBytes,
			Error:    finalErr,
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it serializerLogic) UsingAny(
	anyItem interface{},
) Result {
	return it.Apply(anyItem).NonPtr()
}

func (it serializerLogic) Raw(
	anyItem interface{},
) ([]byte, error) {
	jsonResult := it.Apply(anyItem)

	return jsonResult.Raw()
}

func (it serializerLogic) Marshal(
	anyItem interface{},
) ([]byte, error) {
	jsonResult := it.Apply(anyItem)

	return jsonResult.Raw()
}

func (it serializerLogic) ApplyMust(
	anyItem interface{},
) *Result {
	result := it.Apply(anyItem)
	result.MustBeSafe()

	return result
}

func (it serializerLogic) ToBytesMust(
	anyItem interface{},
) []byte {
	result := it.Apply(anyItem)
	result.MustBeSafe()

	return result.Bytes
}

func (it serializerLogic) ToSafeBytesMust(
	anyItem interface{},
) []byte {
	result := it.Apply(anyItem)
	result.MustBeSafe()

	return result.SafeBytes()
}

// ToSafeBytesSwallowErr
//
// Warning or Danger:
//  - shallow err by not throwing or returning (could be dangerous as well)
//
// Notes :
//  - To inform use Err or Apply or must methods
//
// Use case (rarely):
//  - When don't care about the error just proceed with the value.
func (it serializerLogic) ToSafeBytesSwallowErr(
	anyItem interface{},
) []byte {
	result := it.Apply(anyItem)

	return result.SafeBytes()
}

// ToBytesSwallowErr
//
// Warning or Danger:
//  - shallow err by not throwing or returning (could be dangerous as well)
//
// Notes :
//  - To inform use Err or Apply or must methods
//
// Use case (rarely):
//  - When don't care about the error just proceed with the value.
func (it serializerLogic) ToBytesSwallowErr(
	anyItem interface{},
) []byte {
	result := it.Apply(anyItem)

	return result.Bytes
}

func (it serializerLogic) ToBytesErr(
	anyItem interface{},
) ([]byte, error) {
	result := it.Apply(anyItem)

	return result.Bytes, result.MeaningfulError()
}

// ToString
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
func (it serializerLogic) ToString(
	anyItem interface{},
) string {
	result := it.Apply(anyItem)

	return result.JsonString()
}

func (it serializerLogic) ToStringMust(
	anyItem interface{},
) string {
	result := it.Apply(anyItem)
	result.HandleError()

	return result.JsonString()
}

func (it serializerLogic) ToStringErr(
	anyItem interface{},
) (string, error) {
	result := it.Apply(anyItem)

	return result.RawString()
}

func (it serializerLogic) ToPrettyStringErr(
	anyItem interface{},
) (string, error) {
	result := it.Apply(anyItem)

	return result.RawPrettyString()
}

// ToPrettyStringIncludingErr
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
func (it serializerLogic) ToPrettyStringIncludingErr(
	anyItem interface{},
) string {
	result := it.Apply(anyItem)

	return result.PrettyJsonStringOrErrString()
}
