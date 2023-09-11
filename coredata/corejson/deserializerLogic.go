package corejson

import (
	"encoding/json"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type deserializerLogic struct {
	BytesTo  deserializeFromBytesTo
	ResultTo deserializeFromResultTo
}

func (it deserializerLogic) Apply(
	jsonResult *Result,
	toPtr interface{},
) error {
	return jsonResult.Unmarshal(
		toPtr)
}

func (it deserializerLogic) UsingStringPtr(
	jsonString *string,
	toPtr interface{},
) error {
	if jsonString == nil {
		return it.UsingBytes(
			nil,
			toPtr)
	}

	return it.UsingString(
		*jsonString,
		toPtr)
}

func (it deserializerLogic) UsingError(
	errInJsonFormat error,
	toPtr interface{},
) error {
	if errInJsonFormat == nil {
		return nil
	}

	return it.UsingString(
		errInJsonFormat.Error(),
		toPtr)
}

// UsingErrorWhichJsonResult
//
//	given error is in json format for json result
func (it deserializerLogic) UsingErrorWhichJsonResult(
	errInJsonResultJson error,
	toPtr interface{},
) error {
	if errInJsonResultJson == nil {
		return nil
	}

	jsonResult := NewResult.UsingStringWithType(
		errInJsonResultJson.Error(),
		"ErrorAsJsonResult")

	return jsonResult.Deserialize(toPtr)
}

func (it deserializerLogic) UsingResult(
	jsonResult *Result,
	toPtr interface{},
) error {
	return jsonResult.Unmarshal(
		toPtr)
}

func (it deserializerLogic) ApplyMust(
	jsonResult *Result,
	toPtr interface{},
) {
	err := jsonResult.Unmarshal(
		toPtr)

	if err != nil {
		panic(err)
	}
}

func (it deserializerLogic) UsingString(
	jsonString string,
	toPtr interface{},
) error {
	return it.UsingBytes(
		[]byte(jsonString),
		toPtr)
}

func (it deserializerLogic) FromString(
	jsonString string,
	toPtr interface{},
) error {
	return it.UsingBytes(
		[]byte(jsonString),
		toPtr)
}

// FromTo
//
// Giving nil is not support from to.
//
// Warning: must check nil before for from, to both.
//
// Casting from to steps:
//   - reflection first if equal type + right ptr and not nil.
//   - []byte
//   - string
//   - Jsoner
//   - Result
//   - *Result
//   - bytesSerializer
//   - serializerFunc
//   - error to string then cast from json string then to actual unmarshal
func (it deserializerLogic) FromTo(
	fromAny interface{},
	toPtr interface{},
) error {
	return CastAny.FromToDefault(
		fromAny,
		toPtr)
}

func (it deserializerLogic) MapAnyToPointer(
	isSkipOnEmpty bool,
	currentItemMap map[string]interface{},
	toPtr interface{},
) error {
	if isSkipOnEmpty && len(currentItemMap) == 0 {
		return nil
	}

	jsonResult := New(currentItemMap)

	if jsonResult.HasIssuesOrEmpty() {
		return jsonResult.MeaningfulError()
	}

	return jsonResult.Deserialize(toPtr)
}

func (it deserializerLogic) UsingStringOption(
	isIgnoreEmptyString bool,
	jsonString string,
	toPtr interface{},
) error {
	if isIgnoreEmptyString && jsonString == "" {
		return nil
	}

	return it.UsingBytes(
		[]byte(jsonString),
		toPtr)
}

func (it deserializerLogic) UsingStringIgnoreEmpty(
	jsonString string,
	toPtr interface{},
) error {
	if jsonString == "" {
		return nil
	}

	return it.UsingBytes(
		[]byte(jsonString),
		toPtr)
}

// UsingBytes
//
// json.Unmarshal bytes to object
func (it deserializerLogic) UsingBytes(
	rawBytes []byte,
	toPtr interface{},
) error {
	err := json.Unmarshal(
		rawBytes,
		toPtr)

	if err == nil {
		return nil
	}

	var payloadString string
	if len(rawBytes) > 0 {
		payloadString = string(rawBytes)
	}

	// has error
	compiledMessage := errcore.MessageVarMap(
		"json unmarshal failed",
		map[string]interface{}{
			"err":     err,
			"dst":     reflectinternal.TypeName(toPtr),
			"payload": payloadString,
		})

	return errcore.
		UnMarshallingFailedType.
		ErrorNoRefs(compiledMessage)

}

func (it deserializerLogic) UsingBytesPointerMust(
	rawBytesPointer *[]byte,
	toPtr interface{},
) {
	err := it.UsingBytesPointer(
		rawBytesPointer,
		toPtr)

	if err != nil {
		panic(err)
	}
}

func (it deserializerLogic) UsingBytesIf(
	isDeserialize bool,
	rawBytes []byte,
	toPtr interface{},
) error {
	if !isDeserialize {
		return nil
	}

	return it.UsingBytes(
		rawBytes,
		toPtr)
}

func (it deserializerLogic) UsingBytesPointerIf(
	isDeserialize bool,
	rawBytesPointer *[]byte,
	toPtr interface{},
) error {
	if !isDeserialize {
		return nil
	}

	return it.UsingBytesPointer(
		rawBytesPointer,
		toPtr)
}

func (it deserializerLogic) UsingBytesPointer(
	rawBytesPointer *[]byte,
	toPtr interface{},
) error {
	if rawBytesPointer == nil || *rawBytesPointer == nil {
		reference := errcore.VarTwoNoType(
			"rawBytesPointer", constants.NilAngelBracket,
			"To Reference Type", reflectinternal.TypeName(toPtr))

		return errcore.
			UnMarshallingFailedType.
			Error(
				"failed to unmarshal nil bytes pointer.",
				reference)
	}

	return it.UsingBytes(
		*rawBytesPointer,
		toPtr)
}

func (it deserializerLogic) UsingBytesMust(
	rawBytes []byte,
	toPtr interface{},
) {
	err := it.UsingBytes(
		rawBytes,
		toPtr)

	if err != nil {
		panic(err)
	}
}

func (it deserializerLogic) UsingSafeBytesMust(
	rawBytes []byte,
	toPtr interface{},
) {
	if len(rawBytes) == 0 {
		return
	}

	err := it.UsingBytes(rawBytes, toPtr)

	if err != nil {
		panic(err)
	}
}

func (it deserializerLogic) AnyToFieldsMap(
	anyItem interface{},
) (map[string]interface{}, error) {
	jsonResult := New(anyItem)

	return jsonResult.DeserializedFieldsToMap()
}

func (it deserializerLogic) UsingSerializerTo(
	serializer bytesSerializer,
	toPtr interface{},
) (parsingErr error) {
	jsonResult := NewResult.UsingSerializer(
		serializer)

	return jsonResult.Deserialize(toPtr)
}

func (it deserializerLogic) UsingSerializerFuncTo(
	serializerFunc func() ([]byte, error),
	toPtr interface{},
) (parsingErr error) {
	jsonResult := NewResult.UsingSerializerFunc(
		serializerFunc)

	return jsonResult.Deserialize(toPtr)
}

func (it deserializerLogic) UsingDeserializerToOption(
	isSkipOnDeserializerNull bool,
	deserializer bytesDeserializer,
	toPtr interface{},
) (parsingErr error) {
	if isSkipOnDeserializerNull && deserializer == nil {
		return nil
	}

	if deserializer == nil {
		return errcore.CannotBeNilType.ErrorNoRefs(
			"deserializer is nil",
		)
	}

	return deserializer.Deserialize(toPtr)
}

// UsingDeserializerDefined
//
//	on deserializer null it will not do anything but return nil error
//
// only deserialize if deserializer is not null.
func (it deserializerLogic) UsingDeserializerDefined(
	deserializer bytesDeserializer,
	toPtr interface{},
) (parsingErr error) {
	return it.UsingDeserializerToOption(
		true,
		deserializer,
		toPtr)
}

// UsingDeserializerFuncDefined
//
//	on deserializer null it will not do anything but return nil error
//
// only deserialize if deserializer is not null.
func (it deserializerLogic) UsingDeserializerFuncDefined(
	deserializerFunc func(toPtr interface{}) error,
	toPtr interface{},
) (parsingErr error) {
	if deserializerFunc == nil {
		return errcore.CannotBeNilType.ErrorNoRefs(
			"deserializer function is nil",
		)
	}

	return deserializerFunc(toPtr)
}

func (it deserializerLogic) UsingJsonerToAny(
	isSkipOnNullJsoner bool,
	jsoner Jsoner,
	toPtr interface{},
) error {
	if isSkipOnNullJsoner && jsoner == nil {
		return nil
	}

	if jsoner == nil {
		return errcore.
			CannotBeNilType.
			ErrorNoRefs("jsoner given as nil cannot deserialize to")
	}

	jsonResult := jsoner.JsonPtr()

	return jsonResult.Deserialize(toPtr)
}

func (it deserializerLogic) UsingJsonerToAnyMust(
	isSkipOnNullJsoner bool,
	jsoner Jsoner,
	toPtr interface{},
) error {
	if isSkipOnNullJsoner && jsoner == nil {
		return nil
	}

	if jsoner == nil {
		return errcore.
			CannotBeNilType.
			ErrorNoRefs("jsoner given as nil cannot deserialize to")
	}

	jsonResult := jsoner.JsonPtr()

	return jsonResult.Deserialize(toPtr)
}
