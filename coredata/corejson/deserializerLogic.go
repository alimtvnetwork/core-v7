package corejson

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type deserializerLogic struct {
	BytesTo  deserializeFromBytesTo
	ResultTo deserializeFromResultTo
}

func (it deserializerLogic) Apply(
	jsonResult *Result,
	unmarshalToPointer interface{},
) error {
	return jsonResult.Unmarshal(
		unmarshalToPointer)
}

func (it deserializerLogic) ApplyMust(
	jsonResult *Result,
	unmarshalToPointer interface{},
) {
	err := jsonResult.Unmarshal(
		unmarshalToPointer)

	if err != nil {
		panic(err)
	}
}

func (it deserializerLogic) UsingString(
	jsonString string,
	unmarshalToPointer interface{},
) error {
	return it.UsingBytes(
		[]byte(jsonString),
		unmarshalToPointer)
}

func (it deserializerLogic) UsingStringOption(
	isIgnoreEmptyString bool,
	jsonString string,
	unmarshalToPointer interface{},
) error {
	if isIgnoreEmptyString && jsonString == "" {
		return nil
	}

	return it.UsingBytes(
		[]byte(jsonString),
		unmarshalToPointer)
}

func (it deserializerLogic) UsingStringIgnoreEmpty(
	jsonString string,
	unmarshalToPointer interface{},
) error {
	if jsonString == "" {
		return nil
	}

	return it.UsingBytes(
		[]byte(jsonString),
		unmarshalToPointer)
}

// UsingBytes
//
// json.Unmarshal bytes to object
func (it deserializerLogic) UsingBytes(
	rawBytes []byte,
	unmarshalToPointer interface{},
) error {
	err := json.Unmarshal(
		rawBytes,
		unmarshalToPointer)

	if err != nil {
		reference := errcore.VarTwoNoType(
			"JsonResult Error", err.Error(),
			"To Reference Type", reflectinternal.TypeName(unmarshalToPointer))

		return errcore.
			UnMarshallingFailedType.
			Error(
				"failed to unmarshal bytes.",
				reference)
	}

	return nil
}

func (it deserializerLogic) UsingBytesPointerMust(
	rawBytesPointer *[]byte,
	unmarshalToPointer interface{},
) {
	err := it.UsingBytesPointer(
		rawBytesPointer,
		unmarshalToPointer)

	if err != nil {
		panic(err)
	}
}

func (it deserializerLogic) UsingBytesIf(
	isDeserialize bool,
	rawBytes []byte,
	unmarshalToPointer interface{},
) error {
	if !isDeserialize {
		return nil
	}

	return it.UsingBytes(
		rawBytes,
		unmarshalToPointer)
}

func (it deserializerLogic) UsingBytesPointerIf(
	isDeserialize bool,
	rawBytesPointer *[]byte,
	unmarshalToPointer interface{},
) error {
	if !isDeserialize {
		return nil
	}

	return it.UsingBytesPointer(
		rawBytesPointer,
		unmarshalToPointer)
}

func (it deserializerLogic) UsingBytesPointer(
	rawBytesPointer *[]byte,
	unmarshalToPointer interface{},
) error {
	if rawBytesPointer == nil || *rawBytesPointer == nil {
		reference := errcore.VarTwoNoType(
			"rawBytesPointer", constants.NilAngelBracket,
			"To Reference Type", reflectinternal.TypeName(unmarshalToPointer))

		return errcore.
			UnMarshallingFailedType.
			Error(
				"failed to unmarshal nil bytes pointer.",
				reference)
	}

	return it.UsingBytes(
		*rawBytesPointer,
		unmarshalToPointer)
}

func (it deserializerLogic) UsingBytesMust(
	rawBytes []byte,
	unmarshalToPointer interface{},
) {
	err := it.UsingBytes(
		rawBytes,
		unmarshalToPointer)

	if err != nil {
		panic(err)
	}
}

func (it deserializerLogic) UsingSafeBytesMust(
	rawBytes []byte,
	unmarshalToPointer interface{},
) {
	if len(rawBytes) == 0 {
		return
	}

	err := it.UsingBytes(rawBytes, unmarshalToPointer)

	if err != nil {
		panic(err)
	}
}
