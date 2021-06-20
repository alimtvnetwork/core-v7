package coredynamic

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/bitsize"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/coreonce"
	"gitlab.com/evatix-go/core/internal/messages"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
	"gitlab.com/evatix-go/core/internal/strutilinternal"
	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/core/msgtype"
)

type Dynamic struct {
	innerData       interface{}
	isValid         bool
	reflectType     reflect.Type
	reflectVal      *reflect.Value
	innerDataString *string
	typeName        *coreonce.StringOnce
	length          *coreonce.IntegerOnce
	isPointer       issetter.Value
}

func InvalidDynamic(
) Dynamic {
	return *InvalidDynamicPtr()
}

func InvalidDynamicPtr(
) *Dynamic {
	return NewDynamicPtr(
		nil,
		false)
}

func NewDynamic(
	data interface{},
	isValid bool,
) Dynamic {
	return *NewDynamicPtr(data, isValid)
}

func NewDynamicPtr(
	data interface{},
	isValid bool,
) *Dynamic {
	return &Dynamic{
		innerData: data,
		isValid:   isValid,
		typeName: coreonce.NewStringOncePtr(func() string {
			return fmt.Sprintf(constants.SprintTypeFormat, data)
		}),
		length: coreonce.NewIntegerOncePtr(func() int {
			if data == nil {
				return 0
			}

			return LengthOfReflect(reflect.ValueOf(data))
		}),
	}
}

func (receiver *Dynamic) Data() interface{} {
	return receiver.innerData
}

func (receiver *Dynamic) Value() interface{} {
	return receiver.innerData
}

// Length Returns length of a slice, map, array
//
// It will also reduce from pointer
//
// Reference : https://cutt.ly/PnaWAFn | https://cutt.ly/jnaEig8 | https://play.golang.org/p/UCORoShXlv1
func (receiver *Dynamic) Length() int {
	return receiver.length.Value()
}

func (receiver *Dynamic) StructStringPtr() *string {
	if receiver.innerDataString != nil {
		return receiver.innerDataString
	}

	toString := strutilinternal.AnyToString(receiver.innerData)
	receiver.innerDataString = &toString

	return receiver.innerDataString
}

func (receiver *Dynamic) ReflectValue() *reflect.Value {
	if receiver.reflectVal != nil {
		return receiver.reflectVal
	}

	reflectValueOfAny := reflect.ValueOf(receiver.innerData)
	receiver.reflectVal = &reflectValueOfAny

	return receiver.reflectVal
}

func (receiver *Dynamic) MapToKeyVal() (*KeyValCollection, error) {
	return MapAsKeyValSlice(*receiver.ReflectValue())
}

func (receiver *Dynamic) ReflectKind() reflect.Kind {
	return receiver.ReflectValue().Kind()
}

func (receiver *Dynamic) ReflectTypeName() string {
	return receiver.typeName.Value()
}

func (receiver *Dynamic) ReflectType() reflect.Type {
	if receiver.reflectType != nil {
		return receiver.reflectType
	}

	reflectType := reflect.TypeOf(receiver.innerData)
	receiver.reflectType = reflectType

	return receiver.reflectType
}

func (receiver *Dynamic) IsReflectTypeOf(
	typeRequest reflect.Type,
) bool {
	return receiver.ReflectType() == typeRequest
}

func (receiver *Dynamic) String() string {
	return *receiver.StructStringPtr()
}

func (receiver *Dynamic) StructString() string {
	return *receiver.StructStringPtr()
}

func (receiver *Dynamic) IsReflectKind(checkingKind reflect.Kind) bool {
	return receiver.ReflectKind() == checkingKind
}

func (receiver *Dynamic) IsPointer() bool {
	if receiver.isPointer.IsUninitialized() {
		receiver.isPointer = issetter.GetBool(
			receiver.IsReflectKind(reflect.Ptr))
	}

	return receiver.isPointer.IsTrue()
}

func (receiver *Dynamic) IsValueType() bool {
	return !receiver.IsPointer()
}

func (receiver *Dynamic) IsStructStringNullOrEmpty() bool {
	return receiver.IsNull() || strutilinternal.IsNullOrEmpty(
		receiver.StructStringPtr())
}

func (receiver *Dynamic) IsStructStringNullOrEmptyOrWhitespace() bool {
	return receiver.IsNull() || strutilinternal.IsNullOrEmptyOrWhitespace(
		receiver.StructStringPtr())
}

func (receiver *Dynamic) IsPrimitive() bool {
	return reflectinternal.IsPrimitive(receiver.ReflectKind())
}

// IsNumber true if float (any), byte, int (any), uint(any)
func (receiver *Dynamic) IsNumber() bool {
	return reflectinternal.IsNumber(receiver.ReflectKind())
}

func (receiver *Dynamic) IntDefault(defaultInt int) (val int, isSuccess bool) {
	if receiver.IsNull() {
		return defaultInt, false
	}

	stringVal := receiver.StructString()

	return converters.StringToIntegerWithDefault(stringVal, defaultInt)
}

func (receiver *Dynamic) Float64() (val float64, err error) {
	if receiver.IsNull() {
		return constants.Zero, msgtype.
			ParsingFailed.Error(
			messages.DynamicFailedToParseToFloat64BecauseNull,
			receiver.String())
	}

	stringVal := receiver.StructString()
	valFloat, err2 := strconv.ParseFloat(stringVal, bitsize.Of64)

	if err2 != nil {
		reference := stringVal +
			constants.NewLineUnix +
			err2.Error()

		return constants.Zero, msgtype.
			ParsingFailed.Error(
			msgtype.FailedToConvert.String(),
			reference)
	}

	return valFloat, err
}

func (receiver *Dynamic) IsStruct() bool {
	return receiver.ReflectKind() == reflect.Struct
}

func (receiver *Dynamic) IsFunc() bool {
	return receiver.ReflectKind() == reflect.Func
}

func (receiver *Dynamic) IsSliceOrArray() bool {
	k := receiver.ReflectKind()

	return k == reflect.Slice || k == reflect.Array
}

func (receiver *Dynamic) IsSliceOrArrayOrMap() bool {
	k := receiver.ReflectKind()

	return k == reflect.Slice ||
		k == reflect.Array ||
		k == reflect.Map
}

func (receiver *Dynamic) IsMap() bool {
	return receiver.ReflectKind() == reflect.Map
}

func (receiver *Dynamic) IsNull() bool {
	return receiver.innerData == nil
}

func (receiver *Dynamic) IsValid() bool {
	return receiver.isValid
}

func (receiver *Dynamic) IsInvalid() bool {
	return !receiver.isValid
}

func (receiver *Dynamic) ConvertUsingFunc(
	converter SimpleInOutConverter,
	expectedType reflect.Type,
) *SimpleResult {
	return converter(receiver.innerData, expectedType)
}

// JsonBytesPtr returns empty string on nil.
// no error on nil.
func (receiver *Dynamic) JsonBytesPtr() (jsonBytesPtr *[]byte, err error) {
	if receiver.IsNull() {
		return &[]byte{}, nil
	}

	marshalledBytes, e := json.Marshal(receiver.innerData)

	if e != nil {
		return &[]byte{}, e
	}

	return &marshalledBytes, nil
}

func (receiver *Dynamic) MarshalJSON() ([]byte, error) {
	return json.Marshal(receiver.innerData)
}

func (receiver *Dynamic) UnmarshalJSON(data []byte) error {
	return msgtype.
		NotImplemented.
		Error(msgtype.UnMarshallingFailed.String(), data)
}

func (receiver *Dynamic) JsonBytes() (jsonBytesPtr []byte, err error) {
	allBytes, err := receiver.JsonBytesPtr()

	if err != nil {
		return []byte{}, err
	}

	return *allBytes, err
}

func (receiver *Dynamic) StringJson() (jsonString string, err error) {
	marshalledBytes, err := receiver.JsonBytes()

	if err != nil {
		return constants.EmptyString, err
	}

	return string(marshalledBytes), err
}

func (receiver *Dynamic) StringJsonMust() string {
	marshalledBytes, err := receiver.JsonBytes()

	if err != nil {
		msgtype.
			MarshallingFailed.
			HandleUsingPanic(err.Error(), receiver.innerDataString)
	}

	return string(marshalledBytes)
}

func (receiver *Dynamic) Clone() *Dynamic {
	if receiver == nil {
		return nil
	}

	return NewDynamicPtr(
		receiver.innerData,
		receiver.isValid)
}
