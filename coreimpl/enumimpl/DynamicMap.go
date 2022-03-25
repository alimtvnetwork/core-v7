package enumimpl

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type DynamicMap map[string]interface{}

func (it DynamicMap) AddOrUpdate(key string, val interface{}) (isAddNewly bool) {
	_, isAlreadyExist := it[key]
	it[key] = val

	return !isAlreadyExist
}

func (it *DynamicMap) Set(key string, val interface{}) (isAddNewly bool) {
	if it == nil {
		// mutating because it is part of
		*it = make(map[string]interface{}, constants.Capacity5)
	}

	_, isAlreadyExist := (*it)[key]
	(*it)[key] = val

	return !isAlreadyExist
}

// AddNewOnly
//
//  Don't update existing
func (it *DynamicMap) AddNewOnly(key string, val interface{}) (isAdded bool) {
	if it == nil {
		// mutating because it is part of
		*it = make(map[string]interface{}, constants.Capacity5)
	}

	_, isAlreadyExist := (*it)[key]
	if isAlreadyExist {
		return false
	}

	(*it)[key] = val

	return true
}

func (it DynamicMap) AllKeys() []string {
	if it.IsEmpty() {
		return []string{}
	}

	allKeys := make(
		[]string,
		it.Length())

	index := 0
	for key := range it {
		allKeys[index] = key
		index++
	}

	return allKeys
}

func (it DynamicMap) AllKeysSorted() []string {
	if it.IsEmpty() {
		return []string{}
	}

	allKeys := it.AllKeys()
	sort.Strings(allKeys)

	return allKeys
}

func (it DynamicMap) AllValuesStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	allValues := make(
		[]string,
		it.Length())

	index := 0
	for _, value := range it {
		allValues[index] = fmt.Sprintf(
			constants.SprintValueFormat,
			value)
		index++
	}

	return allValues
}

func (it DynamicMap) AllValuesStringsSorted() []string {
	if it.IsEmpty() {
		return []string{}
	}

	allValues := it.AllValuesStrings()
	sort.Strings(allValues)

	return allValues
}

func (it DynamicMap) AllValuesIntegers() []int {
	if it.IsEmpty() {
		return []int{}
	}

	allValues := make(
		[]int,
		it.Length())

	index := 0
	for _, value := range it {
		allValues[index] = ConvEnumAnyValToInteger(value)

		index++
	}

	return allValues
}

func (it DynamicMap) MapIntegerString() (
	rangeMap map[int]string,
	allKeysSorted []int,
) {
	if it.IsEmpty() {
		return map[int]string{}, []int{}
	}

	rangeMap = make(
		map[int]string,
		it.Length()+2)

	allKeysSorted = make(
		[]int,
		it.Length())

	if it.IsValueString() {
		return it.stringValueMapIntegerString(rangeMap, allKeysSorted)
	}

	index := 0
	for key, value := range it {
		valInt := ConvEnumAnyValToInteger(value)
		rangeMap[valInt] = key
		allKeysSorted[index] = valInt

		index++
	}

	sort.Ints(allKeysSorted)

	return rangeMap, allKeysSorted
}

func (it DynamicMap) SortedKeyValues() (
	keyValues []KeyValInteger,
) {
	if it.IsEmpty() {
		return keyValues
	}

	keyValues = make(
		[]KeyValInteger,
		it.Length())

	rangesMap, AllKeysSorted := it.MapIntegerString()

	for i, keyInt := range AllKeysSorted {
		name := rangesMap[keyInt]
		keyValues[i] = KeyValInteger{
			Key:          name,
			ValueInteger: keyInt,
		}
	}

	return keyValues
}

func (it DynamicMap) SortedKeyAnyValues() (
	keyAnyValues []KeyAnyVal,
) {
	if it.IsEmpty() {
		return keyAnyValues
	}

	keyAnyValues = make(
		[]KeyAnyVal,
		it.Length())

	if it.IsValueString() {
		return it.sortedKeyAnyValuesString()
	}

	rangesMap, AllKeysSorted := it.MapIntegerString()

	for i, keyInt := range AllKeysSorted {
		name := rangesMap[keyInt]
		keyAnyValues[i] = KeyAnyVal{
			Key:      name,
			AnyValue: keyInt,
		}
	}

	return keyAnyValues
}

func (it DynamicMap) First() (key string, valInf interface{}) {
	for key, valInf = range it {
		return key, valInf
	}

	return "", nil
}

func (it DynamicMap) IsValueTypeOf(rfType reflect.Type) bool {
	_, v := it.First()

	return reflect.TypeOf(v) == rfType
}

func (it DynamicMap) IsValueString() bool {
	_, v := it.First()
	_, isString := v.(string)

	return isString
}

func (it *DynamicMap) Length() int {
	if it == nil {
		return 0
	}

	return len(*it)
}

func (it DynamicMap) Count() int {
	return it.Length()
}

func (it DynamicMap) IsEmpty() bool {
	return it.Length() == 0
}

func (it DynamicMap) HasAnyItem() bool {
	return it.Length() > 0
}

func (it DynamicMap) LastIndex() int {
	return it.Length() - 1
}

func (it DynamicMap) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it DynamicMap) HasKey(key string) bool {
	_, has := it[key]

	return has
}

func (it DynamicMap) HasAllKeys(keys ...string) bool {
	for _, key := range keys {
		if it.IsMissingKey(key) {
			return false
		}
	}

	return true
}

func (it DynamicMap) HasAnyKeys(keys ...string) bool {
	for _, key := range keys {
		if it.HasKey(key) {
			return true
		}
	}

	return false
}

func (it DynamicMap) IsMissingKey(key string) bool {
	_, has := it[key]

	return !has
}

func (it *DynamicMap) IsMismatch(
	isRegardlessType bool,
	rightMap *DynamicMap,
) bool {
	return !it.IsEqual(isRegardlessType, rightMap)
}

func (it *DynamicMap) IsRawMismatch(
	isRegardlessType bool,
	rightMap map[string]interface{},
) bool {
	return !it.IsRawEqual(isRegardlessType, rightMap)
}

func (it *DynamicMap) IsEqual(
	isRegardlessType bool,
	rightMap *DynamicMap,
) bool {
	if it == nil && rightMap == nil {
		return true
	}

	if it == nil || rightMap == nil {
		return false
	}

	if it == rightMap {
		return true
	}

	return it.IsRawEqual(
		isRegardlessType,
		*rightMap)
}

func (it *DynamicMap) IsRawEqual(
	isRegardlessType bool,
	rightMap map[string]interface{},
) bool {
	if it == nil && rightMap == nil {
		return true
	}

	if it == nil || rightMap == nil {
		return false
	}

	if it.Length() != len(rightMap) {
		return false
	}

	for key, leftValInf := range *it {
		rightValInf, has := rightMap[key]

		if !has {
			return false
		}

		if it.isNotEqual(
			isRegardlessType,
			leftValInf,
			rightValInf) {
			return false
		}
	}

	return true
}

func (it DynamicMap) Raw() map[string]interface{} {
	return it
}

func (it *DynamicMap) DiffRaw(
	isRegardlessType bool,
	rightMap map[string]interface{},
) DynamicMap {
	if it == nil && rightMap == nil {
		return map[string]interface{}{}
	}

	if it == nil && rightMap != nil {
		return rightMap
	}

	if it != nil && rightMap == nil {
		return *it
	}

	length := it.Length() / 3
	diffMap := make(
		map[string]interface{},
		length)

	for key, leftValInf := range *it {
		rightValInf, has := rightMap[key]

		if !has {
			diffMap[key] = leftValInf

			continue
		}

		if it.isNotEqual(
			isRegardlessType,
			leftValInf,
			rightValInf) {
			diffMap[key] = leftValInf
		}
	}

	if len(diffMap) == 0 && it.Length() == len(rightMap) {
		return diffMap
	}

	leftMap := *it
	for rightKey, rightAnyVal := range rightMap {
		_, hasDiff := diffMap[rightKey]

		if hasDiff {
			// already added

			continue
		}

		leftVal, has := leftMap[rightKey]

		if !has {
			diffMap[rightKey] = rightAnyVal

			continue
		}

		if it.isNotEqual(
			isRegardlessType,
			rightAnyVal,
			leftVal) {
			diffMap[rightKey] = rightAnyVal
		}
	}

	return diffMap
}

func (it *DynamicMap) DiffJsonMessage(
	isRegardlessType bool,
	rightMap map[string]interface{},
) string {
	diffMap := it.DiffRaw(isRegardlessType, rightMap)

	if diffMap.Length() == 0 {
		return ""
	}

	slice := toStringsSliceOfDiffMap(diffMap)
	compiledString := strings.Join(
		slice,
		constants.CommaUnixNewLine)

	return fmt.Sprintf(
		curlyWrapFormat,
		compiledString)
}

func (it *DynamicMap) ShouldDiffMessage(
	isRegardlessType bool,
	title string,
	rightMap map[string]interface{},
) string {
	diffMessage := it.DiffJsonMessage(
		isRegardlessType,
		rightMap)

	if diffMessage == "" {
		return ""
	}

	return fmt.Sprintf(
		diffBetweenMapShouldBeMessageFormat,
		title,
		diffMessage)
}

func (it *DynamicMap) LogShouldDiffMessage(
	isRegardlessType bool,
	title string,
	rightMap map[string]interface{},
) (diffMessage string) {
	diffMessage = it.ShouldDiffMessage(
		isRegardlessType,
		title,
		rightMap)

	if diffMessage == "" {
		return
	}

	fmt.Println(diffMessage)

	return diffMessage
}

func (it *DynamicMap) ExpectingMessage(
	title string,
	expected map[string]interface{},
) string {
	expectedMap := DynamicMap(expected)
	actualMapString := it.String()
	expectedMapString := expectedMap.String()

	isMapEqual := actualMapString == expectedMapString

	if isMapEqual {
		return ""
	}

	return fmt.Sprintf(
		actualVsExpectingMessageFormat,
		title,
		actualMapString,
		expectedMapString,
	)
}

func (it *DynamicMap) LogExpectingMessage(
	title string,
	expected map[string]interface{},
) {
	expectingMessage := it.ExpectingMessage(title, expected)

	if expectingMessage == "" {
		return
	}

	fmt.Println(expectingMessage)
}

func (it *DynamicMap) isNotEqual(
	isRegardlessType bool,
	left,
	right interface{},
) bool {
	if isRegardlessType {
		leftString := fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			left)
		rightString := fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			right)

		return leftString != rightString
	}

	return !reflect.DeepEqual(left, right)
}

func (it *DynamicMap) IsKeysEqualOnly(
	rightMap map[string]interface{},
) bool {
	if it == nil && rightMap == nil {
		return true
	}

	if it == nil || rightMap == nil {
		return false
	}

	if it.Length() != len(rightMap) {
		return false
	}

	for key := range *it {
		_, has := rightMap[key]

		if !has {
			return false
		}
	}

	return true
}

func (it DynamicMap) KeyValue(
	key string,
) (val interface{}, isFound bool) {
	val, isFound = it[key]

	return val, isFound
}

func (it DynamicMap) KeyValueString(
	key string,
) (val string, isFound bool) {
	valInf, isFound := it[key]

	if isFound {
		convString := fmt.Sprintf(
			constants.SprintValueFormat,
			valInf)

		return convString, isFound
	}

	return "", isFound
}

func (it DynamicMap) KeyValueIntDefault(
	key string,
) (val int) {
	valInt, isFound, isConvFailed := it.KeyValueInt(key)

	if !isFound || isConvFailed {
		return constants.InvalidValue
	}

	return valInt
}

func (it DynamicMap) KeyValueByte(
	key string,
) (val byte, isFound, isConvFailed bool) {
	valInf, isFound := it[key]

	if !isFound {
		return constants.Zero, isFound, true
	}

	valueByterCasted, isSuccess := valInf.(valueByter)

	if isSuccess {
		return valueByterCasted.Value(),
			true,
			false
	}

	exactValueByterCasted, isSuccess := valInf.(exactValueByter)

	if isSuccess {
		return exactValueByterCasted.ValueByte(),
			true,
			false
	}

	toByteCasted, isSuccess := valInf.(byte)

	if isSuccess {
		return toByteCasted,
			true,
			false
	}

	toString := fmt.Sprintf(
		constants.SprintValueFormat,
		valInf)

	toInt, err := strconv.Atoi(toString)

	if err != nil {
		return constants.Zero, true, false
	}

	if toInt >= 0 && toInt <= 255 {
		return byte(toInt), true, false
	}

	return constants.Zero, true, true
}

func (it *DynamicMap) Add(
	key string,
	valInf interface{},
) *DynamicMap {
	(*it)[key] = valInf

	return it
}

func (it DynamicMap) KeyValueInt(
	key string,
) (val int, isFound, isConvFailed bool) {
	valInf, isFound := it[key]

	if !isFound {
		return constants.InvalidValue, isFound, true
	}

	valInt, isInt := valInf.(int)
	if isInt {
		return valInt, isFound, false
	}

	valueByterCasted, isByter := valInf.(valueByter)

	if isByter {
		return int(valueByterCasted.Value()), isFound, false
	}

	exactValueByterCasted, isExactByter := valInf.(exactValueByter)

	if isExactByter {
		return int(exactValueByterCasted.ValueByte()), isFound, false
	}

	valByte, isByte := valInf.(byte)
	if isByte {
		return int(valByte), isFound, false
	}

	toString := fmt.Sprintf(
		constants.SprintValueFormat,
		valInf)

	toInt, err := strconv.Atoi(toString)

	if err != nil {
		// failed
		return constants.InvalidValue, true, true
	}

	return toInt, true, false
}

func (it DynamicMap) BasicByte(typeName string) *BasicByte {
	return New.BasicByte.CreateUsingMap(
		typeName,
		it.ConvMapByteString())
}

func (it DynamicMap) BasicByteUsingAliasMap(
	typeName string,
	aliasingMap map[string]byte,
) *BasicByte {
	return New.BasicByte.CreateUsingMapPlusAliasMap(
		typeName,
		it.ConvMapByteString(),
		aliasingMap)
}

func (it DynamicMap) BasicInt8(typeName string) *BasicInt8 {
	return New.
		BasicInt8.
		CreateUsingMap(
			typeName,
			it.ConvMapInt8String())
}

func (it DynamicMap) BasicInt8UsingAliasMap(
	typeName string,
	aliasingMap map[string]int8,
) *BasicInt8 {
	return New.
		BasicInt8.
		CreateUsingMapPlusAliasMap(
			typeName,
			it.ConvMapInt8String(),
			aliasingMap)
}

func (it DynamicMap) BasicInt16(
	typeName string,
) *BasicInt16 {
	return New.
		BasicInt16.
		CreateUsingMap(
			typeName,
			it.ConvMapInt16String())
}

func (it DynamicMap) BasicInt16UsingAliasMap(
	typeName string,
	aliasingMap map[string]int16,
) *BasicInt16 {
	return New.BasicInt16.CreateUsingMapPlusAliasMap(
		typeName,
		it.ConvMapInt16String(),
		aliasingMap)
}

func (it DynamicMap) BasicInt32(
	typeName string,
) *BasicInt32 {
	return New.
		BasicInt32.
		CreateUsingMap(
			typeName,
			it.ConvMapInt32String())
}

func (it DynamicMap) BasicInt32UsingAliasMap(
	typeName string,
	aliasingMap map[string]int32,
) *BasicInt32 {
	return New.
		BasicInt32.
		CreateUsingMapPlusAliasMap(
			typeName,
			it.ConvMapInt32String(),
			aliasingMap)
}

func (it DynamicMap) BasicString(
	typeName string,
) *BasicString {
	return New.
		BasicString.
		Create(
			typeName,
			it.AllKeysSorted())
}

func (it DynamicMap) BasicStringUsingAliasMap(
	typeName string,
	aliasingMap map[string]string,
) *BasicString {
	return New.
		BasicString.
		CreateAliasMapOnly(
			typeName,
			it.AllKeysSorted(),
			aliasingMap)
}

func (it DynamicMap) BasicUInt16(
	typeName string,
) *BasicUInt16 {
	return New.
		BasicUInt16.
		CreateUsingMap(
			typeName,
			it.ConvMapUInt16String())
}

func (it DynamicMap) BasicUInt16UsingAliasMap(
	typeName string,
	aliasingMap map[string]uint16,
) *BasicUInt16 {
	return New.
		BasicUInt16.
		CreateUsingMapPlusAliasMap(
			typeName,
			it.ConvMapUInt16String(),
			aliasingMap)
}

// ConvMapStringInteger
//
//  Conv value to key and key to value.
func (it DynamicMap) ConvMapStringInteger() map[string]int {
	if it.IsEmpty() {
		return map[string]int{}
	}

	newMap := make(map[string]int, it.Length())

	for key := range it {
		valInt := it.KeyValueIntDefault(key)
		newMap[key] = valInt
	}

	return newMap
}

// ConvMapIntegerString
//
//  Conv value to key and key to value.
func (it DynamicMap) ConvMapIntegerString() map[int]string {
	if it.IsEmpty() {
		return map[int]string{}
	}

	newMap := make(map[int]string, it.Length())

	for key := range it {
		valInt := it.KeyValueIntDefault(key)
		newMap[valInt] = key
	}

	return newMap
}

// ConvMapByteString
//
//  Conv value to key and key to value.
func (it DynamicMap) ConvMapByteString() map[byte]string {
	if it.IsEmpty() {
		return map[byte]string{}
	}

	newMap := make(map[byte]string, it.Length())

	for key := range it {
		valByte, isFound, isFailed := it.KeyValueByte(
			key)

		if !isFound || isFailed {
			continue
		}

		newMap[valByte] = key
	}

	return newMap
}

// ConvMapInt8String
//
//  Conv value to key and key to value.
func (it DynamicMap) ConvMapInt8String() map[int8]string {
	if it.IsEmpty() {
		return map[int8]string{}
	}

	newMap := make(map[int8]string, it.Length())

	for key := range it {
		valInt := it.KeyValueIntDefault(
			key)

		if valInt < math.MinInt8 || valInt > math.MaxInt8 {
			continue
		}

		newMap[int8(valInt)] = key
	}

	return newMap
}

// ConvMapInt16String
//
//  Conv value to key and key to value.
func (it DynamicMap) ConvMapInt16String() map[int16]string {
	if it.IsEmpty() {
		return map[int16]string{}
	}

	newMap := make(map[int16]string, it.Length())

	for key := range it {
		valInt := it.KeyValueIntDefault(
			key)

		if valInt < math.MinInt16 || valInt > math.MaxInt16 {
			continue
		}

		newMap[int16(valInt)] = key
	}

	return newMap
}

// ConvMapInt32String
//
//  Conv value to key and key to value.
func (it DynamicMap) ConvMapInt32String() map[int32]string {
	if it.IsEmpty() {
		return map[int32]string{}
	}

	newMap := make(map[int32]string, it.Length())

	for key := range it {
		valInt := it.KeyValueIntDefault(
			key)

		if valInt < math.MinInt32 || valInt > math.MaxInt32 {
			continue
		}

		newMap[int32(valInt)] = key
	}

	return newMap
}

// ConvMapUInt16String
//
//  Conv value to key and key to value.
func (it DynamicMap) ConvMapUInt16String() map[uint16]string {
	if it.IsEmpty() {
		return map[uint16]string{}
	}

	newMap := make(map[uint16]string, it.Length())

	for key := range it {
		valInt := it.KeyValueIntDefault(
			key)

		if valInt < 0 || valInt > math.MaxUint16 {
			continue
		}

		newMap[uint16(valInt)] = key
	}

	return newMap
}

// ConvMapStringString
//
//  Conv value to key and key to value.
func (it DynamicMap) ConvMapStringString() map[string]string {
	if it.IsEmpty() {
		return map[string]string{}
	}

	newMap := make(map[string]string, it.Length())

	for key := range it {
		valString, isFound := it.KeyValueString(
			key)

		if !isFound {
			continue
		}

		newMap[valString] = key
	}

	return newMap
}

// ConvMapInt64String
//
//  Conv value to key and key to value.
func (it DynamicMap) ConvMapInt64String() map[int64]string {
	if it.IsEmpty() {
		return map[int64]string{}
	}

	newMap := make(map[int64]string, it.Length())

	for key := range it {
		valInt := it.KeyValueIntDefault(
			key)

		newMap[int64(valInt)] = key
	}

	return newMap
}

func (it DynamicMap) ConcatNew(
	isOverrideExisting bool,
	another DynamicMap,
) DynamicMap {
	if it.IsEmpty() && another.IsEmpty() {
		return map[string]interface{}{}
	}

	var newMap DynamicMap = make(
		map[string]interface{},
		it.Length()+another.Length()+1)

	if it.HasAnyItem() {
		for key, val := range it {
			newMap[key] = val
		}
	}

	hasAnother := another.HasAnyItem()
	if hasAnother && isOverrideExisting {
		for key, val := range another {
			newMap[key] = val
		}
	} else if hasAnother && !isOverrideExisting {
		for key, val := range another {
			newMap.AddNewOnly(key, val)
		}
	}

	return newMap
}

func (it DynamicMap) Strings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())
	allKeysSorted := it.AllKeysSorted()

	index := 0
	for _, key := range allKeysSorted {
		val := it[key]

		slice[index] = fmt.Sprintf(
			constants.KeyValJsonFormat,
			key,
			val)

		index++
	}

	return slice
}

func (it DynamicMap) StringsUsingFmt(
	formatter func(index int, key string, val interface{}) string,
) []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())
	allKeysSorted := it.AllKeysSorted()

	for i, key := range allKeysSorted {
		val := it[key]
		slice[i] = formatter(
			i,
			key,
			val)
	}

	return slice
}

func (it DynamicMap) String() string {
	return strings.Join(
		it.Strings(),
		constants.DefaultLine)
}

func (it DynamicMap) IsStringEqual(anotherMapString string) bool {
	return it.String() == anotherMapString
}

func (it DynamicMap) Serialize() ([]byte, error) {
	return json.Marshal(it)
}

func (it DynamicMap) sortedKeyAnyValuesString() []KeyAnyVal {
	allStringsSorted := it.AllKeysSorted()
	newSlice := make([]KeyAnyVal, len(allStringsSorted))

	for i, keyName := range allStringsSorted {
		newSlice[i] = KeyAnyVal{
			Key:      keyName,
			AnyValue: keyName,
		}
	}

	return newSlice
}

func (it DynamicMap) stringValueMapIntegerString(
	rangeMap map[int]string, allNumberSorted []int,
) (integerToStringMap map[int]string, sortedIntegers []int) {
	allNames := it.AllKeysSorted()

	for i, name := range allNames {
		rangeMap[constants.MinInt] = name
		allNumberSorted[i] = constants.MinInt
	}

	return rangeMap, allNumberSorted
}
