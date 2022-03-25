package corejson

import (
	"errors"
	"reflect"

	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type castingAny struct{}

func (it castingAny) FromToDefault(
	fromAny,
	castedToPtr interface{},
) (failedOrDeserialized error) {
	return it.FromToOption(
		true,
		fromAny,
		castedToPtr)
}

func (it castingAny) FromToReflection(
	fromAny,
	castedToPtr interface{},
) (failedOrDeserialized error) {
	return it.FromToOption(
		true,
		fromAny,
		castedToPtr)
}

// FromToOption
//
// Giving nil is not support from to.
//
// Warning: must check nil before for from, to both.
//
// Casting from to steps:
//  - reflection first if equal type + right ptr and not nil.
//  - []byte
//  - string
//  - Jsoner
//  - Result
//  - *Result
//  - bytesSerializer
//  - serializerFunc
//  - error to string then cast from json string then to actual unmarshal
func (it castingAny) FromToOption(
	isUseReflection bool,
	fromAny,
	castedToPtr interface{},
) (failedOrDeserialized error) {
	err, isApplicable := it.reflectionCasting(
		isUseReflection,
		fromAny,
		castedToPtr)
	if isApplicable {
		return err
	}

	switch castedFrom := fromAny.(type) {
	case []byte:
		return Deserialize.UsingBytes(
			castedFrom,
			castedToPtr)
	case string:
		return Deserialize.UsingBytes(
			[]byte(castedFrom),
			castedToPtr)
	case Jsoner:
		jsonResult := castedFrom.Json()

		return jsonResult.Deserialize(castedToPtr)
	case Result:
		return castedFrom.Deserialize(castedToPtr)
	case *Result:
		return castedFrom.Deserialize(castedToPtr)
	case bytesSerializer:
		allBytes, parsingErr := castedFrom.Serialize()

		if parsingErr != nil {
			// usually this error
			// contains all info
			return parsingErr
		}

		return Deserialize.UsingBytes(
			allBytes,
			castedToPtr)
	case func() ([]byte, error):
		jsonResult := NewResult.UsingSerializerFunc(
			castedFrom)

		return jsonResult.Deserialize(castedToPtr)
	case error:
		if castedFrom == nil {
			return nil
		}

		parsingErr := Deserialize.UsingBytes(
			[]byte(castedFrom.Error()),
			castedToPtr)

		if parsingErr != nil {
			return errors.New(
				castedFrom.Error() +
					parsingErr.Error())
		}

		return nil
	}

	// from
	serializeJsonResult := Serialize.Apply(
		fromAny)

	// to
	return serializeJsonResult.Deserialize(
		castedToPtr)
}

// reflectionCasting
//
//  todo refactor return err
func (it castingAny) reflectionCasting(
	isUseReflection bool,
	fromAny interface{},
	castedToPtr interface{},
) (err error, isApplicable bool) {
	if !isUseReflection {
		return nil, false
	}

	if fromAny == nil || castedToPtr == nil {
		// represents interface nil
		// having type to nil will not be captured here.
		// intentionally not taking it -- not a mistake
		return errors.New(
			"cannot cast from to if any from or to is null"), false
	}

	leftType := reflect.TypeOf(fromAny)
	rightType := reflect.TypeOf(castedToPtr)

	if leftType != rightType {
		return nil, false
	}

	isRightPtr := rightType.Kind() == reflect.Ptr

	if !isRightPtr {
		return nil, false
	}

	isLeftDefined := reflectinternal.IsNotNull(fromAny)

	if !isLeftDefined {
		return nil, false
	}

	isRightDefined := reflectinternal.IsNotNull(castedToPtr)

	if !isRightDefined {
		return nil, false
	}

	// ptr, same
	toVal := reflect.
		ValueOf(castedToPtr).
		Elem()
	reflect.
		ValueOf(fromAny).Elem().
		Set(toVal)

	return nil, true
}

func (it castingAny) OrDeserializeTo(
	fromAny,
	castedToPtr interface{},
) (failedOrDeserialized error) {
	return it.FromToDefault(fromAny, castedToPtr)
}
