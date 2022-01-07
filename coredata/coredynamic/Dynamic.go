package coredynamic

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/bitsize"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/coreonce"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/messages"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
	"gitlab.com/evatix-go/core/internal/utilstringinternal"
	"gitlab.com/evatix-go/core/issetter"
)

type Dynamic struct {
	innerData       interface{}
	isValid         bool
	reflectType     reflect.Type
	reflectVal      *reflect.Value
	innerDataString *string
	typeName        coreonce.StringOnce
	length          coreonce.IntegerOnce
	isPointer       issetter.Value
}

func InvalidDynamic() Dynamic {
	return *InvalidDynamicPtr()
}

func InvalidDynamicPtr() *Dynamic {
	return NewDynamicPtr(
		nil,
		false)
}

func NewDynamicValid(
	data interface{},
) Dynamic {
	return *NewDynamicPtr(data, true)
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
		typeName: coreonce.NewStringOnce(func() string {
			return fmt.Sprintf(constants.SprintTypeFormat, data)
		}),
		length: coreonce.NewIntegerOnce(func() int {
			if data == nil {
				return 0
			}

			return LengthOfReflect(reflect.ValueOf(data))
		}),
	}
}

func (it *Dynamic) Data() interface{} {
	return it.innerData
}

func (it *Dynamic) Value() interface{} {
	return it.innerData
}

// Length Returns length of a slice, map, array
//
// It will also reduce from pointer
//
// Reference : https://cutt.ly/PnaWAFn | https://cutt.ly/jnaEig8 | https://play.golang.org/p/UCORoShXlv1
func (it *Dynamic) Length() int {
	return it.length.Value()
}

func (it *Dynamic) StructStringPtr() *string {
	if it.innerDataString != nil {
		return it.innerDataString
	}

	toString := utilstringinternal.AnyToString(it.innerData)
	it.innerDataString = &toString

	return it.innerDataString
}

func (it *Dynamic) ReflectValue() *reflect.Value {
	if it.reflectVal != nil {
		return it.reflectVal
	}

	reflectValueOfAny := reflect.ValueOf(it.innerData)
	it.reflectVal = &reflectValueOfAny

	return it.reflectVal
}

func (it *Dynamic) MapToKeyVal() (*KeyValCollection, error) {
	return MapAsKeyValSlice(*it.ReflectValue())
}

func (it *Dynamic) ReflectKind() reflect.Kind {
	return it.ReflectValue().Kind()
}

func (it *Dynamic) ReflectTypeName() string {
	return it.typeName.Value()
}

func (it *Dynamic) ReflectType() reflect.Type {
	if it.reflectType != nil {
		return it.reflectType
	}

	reflectType := reflect.TypeOf(it.innerData)
	it.reflectType = reflectType

	return it.reflectType
}

func (it *Dynamic) IsReflectTypeOf(
	typeRequest reflect.Type,
) bool {
	return it.ReflectType() == typeRequest
}

func (it *Dynamic) String() string {
	return *it.StructStringPtr()
}

func (it *Dynamic) StructString() string {
	return *it.StructStringPtr()
}

func (it *Dynamic) IsReflectKind(checkingKind reflect.Kind) bool {
	return it.ReflectKind() == checkingKind
}

func (it *Dynamic) IsPointer() bool {
	if it.isPointer.IsUninitialized() {
		it.isPointer = issetter.GetBool(
			it.IsReflectKind(reflect.Ptr))
	}

	return it.isPointer.IsTrue()
}

func (it *Dynamic) IsValueType() bool {
	return !it.IsPointer()
}

func (it *Dynamic) IsStructStringNullOrEmpty() bool {
	return it.IsNull() || utilstringinternal.IsNullOrEmpty(
		it.StructStringPtr())
}

func (it *Dynamic) IsStructStringNullOrEmptyOrWhitespace() bool {
	return it.IsNull() || utilstringinternal.IsNullOrEmptyOrWhitespace(
		it.StructStringPtr())
}

func (it *Dynamic) IsPrimitive() bool {
	return reflectinternal.IsPrimitive(it.ReflectKind())
}

// IsNumber true if float (any), byte, int (any), uint(any)
func (it *Dynamic) IsNumber() bool {
	return reflectinternal.IsNumber(it.ReflectKind())
}

func (it *Dynamic) IntDefault(defaultInt int) (val int, isSuccess bool) {
	if it.IsNull() {
		return defaultInt, false
	}

	stringVal := it.StructString()

	return converters.StringToIntegerWithDefault(stringVal, defaultInt)
}

func (it *Dynamic) Float64() (val float64, err error) {
	if it.IsNull() {
		return constants.Zero, errcore.
			ParsingFailedType.Error(
			messages.DynamicFailedToParseToFloat64BecauseNull,
			it.String())
	}

	stringVal := it.StructString()
	valFloat, err2 := strconv.ParseFloat(stringVal, bitsize.Of64)

	if err2 != nil {
		reference := stringVal +
			constants.NewLineUnix +
			err2.Error()

		return constants.Zero, errcore.
			ParsingFailedType.Error(
			errcore.FailedToConvertType.String(),
			reference)
	}

	return valFloat, err
}

func (it *Dynamic) IsStruct() bool {
	return it.ReflectKind() == reflect.Struct
}

func (it *Dynamic) IsFunc() bool {
	return it.ReflectKind() == reflect.Func
}

func (it *Dynamic) IsSliceOrArray() bool {
	k := it.ReflectKind()

	return k == reflect.Slice || k == reflect.Array
}

func (it *Dynamic) IsSliceOrArrayOrMap() bool {
	k := it.ReflectKind()

	return k == reflect.Slice ||
		k == reflect.Array ||
		k == reflect.Map
}

func (it *Dynamic) IsMap() bool {
	return it.ReflectKind() == reflect.Map
}

func (it *Dynamic) IsNull() bool {
	return it.innerData == nil
}

func (it *Dynamic) IsValid() bool {
	return it.isValid
}

func (it *Dynamic) IsInvalid() bool {
	return !it.isValid
}

func (it *Dynamic) ConvertUsingFunc(
	converter SimpleInOutConverter,
	expectedType reflect.Type,
) *SimpleResult {
	return converter(it.innerData, expectedType)
}

func (it Dynamic) NonPtr() Dynamic {
	return it
}

func (it *Dynamic) Ptr() *Dynamic {
	return it
}

// JsonBytesPtr returns empty string on nil.
// no error on nil.
func (it *Dynamic) JsonBytesPtr() (jsonBytesPtr *[]byte, err error) {
	if it.IsNull() {
		return &[]byte{}, nil
	}

	marshalledBytes, e := json.Marshal(it.innerData)

	if e != nil {
		return &[]byte{}, e
	}

	return &marshalledBytes, nil
}

func (it *Dynamic) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.innerData)
}

func (it *Dynamic) UnmarshalJSON(data []byte) error {
	return errcore.
		NotImplementedType.
		Error(errcore.UnMarshallingFailedType.String(), data)
}

func (it *Dynamic) JsonModel() interface{} {
	return it.innerData
}

func (it *Dynamic) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it Dynamic) Json() corejson.Result {
	return corejson.New(it)
}

func (it Dynamic) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *Dynamic) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Dynamic, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//goland:noinspection GoLinterLocal
func (it *Dynamic) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Dynamic {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *Dynamic) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *Dynamic) JsonBytes() (jsonBytesPtr []byte, err error) {
	allBytes, err := it.JsonBytesPtr()

	if err != nil {
		return []byte{}, err
	}

	return *allBytes, err
}

func (it *Dynamic) JsonString() (jsonString string, err error) {
	marshalledBytes, err := it.JsonBytes()

	if err != nil {
		return constants.EmptyString, err
	}

	return string(marshalledBytes), err
}

func (it *Dynamic) JsonStringMust() string {
	marshalledBytes, err := it.JsonBytes()

	if err != nil {
		errcore.
			MarshallingFailedType.
			HandleUsingPanic(err.Error(), it.innerDataString)
	}

	return string(marshalledBytes)
}

func (it Dynamic) Clone() Dynamic {
	return NewDynamic(
		it.innerData,
		it.isValid)
}

func (it *Dynamic) ClonePtr() *Dynamic {
	if it == nil {
		return nil
	}

	return NewDynamicPtr(
		it.innerData,
		it.isValid)
}
