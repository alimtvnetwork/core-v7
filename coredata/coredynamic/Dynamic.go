package coredynamic

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/constants/bitsize"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/defaulterr"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/messages"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/internal/strutilinternal"
	"gitlab.com/auk-go/core/issetter"
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

	toString := strutilinternal.AnyToString(it.innerData)
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

func (it *Dynamic) ItemReflectValueUsingIndex(index int) reflect.Value {
	return it.ReflectValue().Index(index)
}

func (it *Dynamic) ItemReflectValueUsingKey(key interface{}) reflect.Value {
	return it.ReflectValue().MapIndex(reflect.ValueOf(key))
}

func (it *Dynamic) ItemUsingIndex(index int) interface{} {
	return it.ReflectValue().Index(index).Interface()
}

func (it *Dynamic) ItemUsingKey(key interface{}) interface{} {
	return it.ReflectValue().MapIndex(reflect.ValueOf(key)).Interface()
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
	return it.IsNull() || strutilinternal.IsNullOrEmpty(
		it.StructStringPtr())
}

func (it *Dynamic) IsStructStringNullOrEmptyOrWhitespace() bool {
	return it.IsNull() || strutilinternal.IsNullOrEmptyOrWhitespace(
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
	toInt, err := strconv.Atoi(stringVal)

	if err == nil {
		return toInt, true
	}

	return defaultInt, false
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

func (it *Dynamic) Loop(
	loopProcessorFunc func(index int, item interface{}) (isBreak bool),
) (isCalled bool) {
	if it.IsInvalid() || it.IsNull() || it.Length() <= 0 {
		return false
	}

	length := it.Length()
	rv := *it.ReflectValue()

	for i := 0; i < length; i++ {
		isBreak := loopProcessorFunc(
			i,
			rv.Index(i).Interface())

		if isBreak {
			return true
		}
	}

	return true
}

func (it *Dynamic) FilterAsDynamicCollection(
	filterFunc func(index int, itemAsDynamic Dynamic) (isTake, isBreak bool),
) *DynamicCollection {
	if it.IsInvalid() || it.IsNull() || it.Length() <= 0 {
		return EmptyDynamicCollection()
	}

	length := it.Length()
	rv := *it.ReflectValue()
	dynamicCollection := NewDynamicCollection(length / 2)

	for i := 0; i < length; i++ {
		currentRv := rv.Index(i)
		valInf := currentRv.Interface()
		currentDynamic := NewDynamic(valInf, currentRv.IsValid())

		isTake, isBreak := filterFunc(
			i,
			currentDynamic)

		if isTake {
			dynamicCollection.Add(currentDynamic)
		}

		if isBreak {
			return dynamicCollection
		}
	}

	return dynamicCollection
}

func (it *Dynamic) LoopMap(
	mapLoopProcessorFunc func(index int, key, value interface{}) (isBreak bool),
) (isCalled bool) {
	if it.IsInvalid() || it.IsNull() || it.Length() <= 0 {
		return false
	}

	rv := *it.ReflectValue()
	mapIterator := rv.MapRange()
	index := 0
	for mapIterator.Next() {
		k := mapIterator.Key()
		v := mapIterator.Value()
		isBreak := mapLoopProcessorFunc(index, k.Interface(), v.Interface())

		if isBreak {
			return true
		}

		index++
	}

	return true
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

func (it *Dynamic) Bytes() (rawBytes []byte, isSuccess bool) {
	if it == nil {
		return nil, false
	}

	rawBytes, isSuccess = it.innerData.([]byte)

	if isSuccess {
		return rawBytes, isSuccess
	}

	rawBytes, err := json.Marshal(it.innerData)

	return rawBytes, err != nil
}

func (it *Dynamic) ReflectSetTo(toPointer interface{}) error {
	if it == nil {
		return defaulterr.NilResult
	}

	return ReflectSetFromTo(
		it.innerData,
		toPointer)
}

func (it *Dynamic) Deserialize(jsonBytes []byte) (deserialized *Dynamic, err error) {
	if it == nil {
		return InvalidDynamicPtr(), defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	err = corejson.
		Deserialize.
		UsingBytes(jsonBytes, it.innerData)

	it.isValid = err == nil

	return it, err
}

func (it *Dynamic) ValueMarshal() (jsonBytes []byte, err error) {
	if it == nil {
		return nil, defaulterr.NilResult
	}

	return corejson.
		Serialize.
		ToBytesErr(it.innerData)
}

func (it *Dynamic) JsonPayloadMust() (jsonBytes []byte) {
	return corejson.
		Serialize.
		ToSafeBytesMust(it.innerData)
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
	return corejson.
		Serialize.
		ToBytesErr(it.innerData)
}

func (it *Dynamic) UnmarshalJSON(data []byte) error {
	if it == nil {
		return defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	err := corejson.
		Deserialize.
		UsingBytes(data, it.innerData)

	it.isValid = err == nil

	return err
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

func (it *Dynamic) IsStringType() bool {
	_, isString := it.innerData.(string)

	return isString
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

func (it *Dynamic) ValueInt() int {
	casted, isSuccess := it.innerData.(int)

	if isSuccess {
		return casted
	}

	return constants.InvalidValue
}

func (it *Dynamic) ValueUInt() uint {
	casted, isSuccess := it.innerData.(uint)

	if isSuccess {
		return casted
	}

	return constants.Zero
}

func (it *Dynamic) ValueStrings() []string {
	casted, isSuccess := it.innerData.([]string)

	if isSuccess {
		return casted
	}

	return nil
}

func (it *Dynamic) ValueBool() bool {
	casted, isSuccess := it.innerData.(bool)

	if isSuccess {
		return casted
	}

	return false
}

func (it *Dynamic) ValueInt64() int64 {
	casted, isSuccess := it.innerData.(int64)

	if isSuccess {
		return casted
	}

	return constants.InvalidValue
}

func (it *Dynamic) ValueNullErr() error {
	if it == nil {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("Dynamic is nil or null")
	}

	if reflectinternal.IsNull(it.innerData) {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("Dynamic internal data is nil.")
	}

	return nil
}

func (it *Dynamic) ValueString() string {
	if it == nil || it.innerData == nil {
		return constants.EmptyString
	}

	currentString, isString := it.innerData.(string)

	if isString {
		return currentString
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.innerData,
	)
}
