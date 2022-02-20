package enumimpl

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
)

type DynamicMap map[string]interface{}

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
	valInt, _, _ := it.KeyValueInt(key)

	return valInt
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
		return valInt, isFound, true
	}

	toString := fmt.Sprintf(
		constants.SprintValueFormat,
		valInf)

	toInt, err := strconv.Atoi(toString)

	if err != nil {
		return constants.InvalidValue, true, false
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
		CreateUsingMap(
			typeName,
			it.ConvMapStringString())
}

func (it DynamicMap) BasicStringUsingAliasMap(
	typeName string,
	aliasingMap map[string]string,
) *BasicString {
	return New.
		BasicString.
		CreateUsingMapPlusAliasMap(
			typeName,
			it.ConvMapStringString(),
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
		valInt := it.KeyValueIntDefault(
			key)

		if valInt < 0 || valInt > constants.MaxUnit8AsInt {
			continue
		}

		newMap[byte(valInt)] = key
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

		if valInt < math.MinInt64 || valInt > math.MaxInt64 {
			continue
		}

		newMap[int64(valInt)] = key
	}

	return newMap
}

func (it DynamicMap) Strings() []string {
	slice := make([]string, it.Length())

	index := 0
	for key, value := range it {
		slice[index] = fmt.Sprintf(
			constants.KeyValShortFormat,
			key,
			value)

		index++
	}

	return slice
}

func (it DynamicMap) String() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.Strings())
}
