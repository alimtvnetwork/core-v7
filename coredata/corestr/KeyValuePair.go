package corestr

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/constants/bitsize"
	"gitlab.com/auk-go/core/coredata/corejson"
)

type KeyValuePair struct {
	Key, Value string
}

func (it KeyValuePair) KeyName() string {
	return it.Key
}

func (it KeyValuePair) VariableName() string {
	return it.Key
}

func (it KeyValuePair) ValueString() string {
	return it.Value
}

func (it KeyValuePair) IsVariableNameEqual(name string) bool {
	return it.Key == name
}

func (it KeyValuePair) IsValueEqual(valueString string) bool {
	return it.Value == valueString
}

func (it KeyValuePair) Json() corejson.Result {
	return corejson.New(it)
}

func (it KeyValuePair) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it KeyValuePair) Serialize() ([]byte, error) {
	return corejson.NewPtr(it).Raw()
}

func (it KeyValuePair) SerializeMust() (jsonBytes []byte) {
	return corejson.NewPtr(it).RawMust()
}

func (it KeyValuePair) Compile() string {
	return it.String()
}

func (it *KeyValuePair) IsKeyEmpty() bool {
	return it.Key == ""
}

func (it *KeyValuePair) IsValueEmpty() bool {
	return it.Value == ""
}

func (it *KeyValuePair) HasKey() bool {
	return it.Key != ""
}

func (it *KeyValuePair) HasValue() bool {
	return it.Value != ""
}

func (it *KeyValuePair) IsKeyValueEmpty() bool {
	return it.Key == "" && it.Value == ""
}

func (it *KeyValuePair) TrimKey() string {
	return strings.TrimSpace(it.Key)
}

func (it *KeyValuePair) TrimValue() string {
	return strings.TrimSpace(it.Value)
}

func (it *KeyValuePair) ValueBool() bool {
	if it.Value == "" {
		return false
	}

	toBool, err := strconv.ParseBool(it.Value)

	if err != nil {
		return false
	}

	return toBool
}

func (it *KeyValuePair) ValueInt(defaultInteger int) int {
	toInt, err := strconv.Atoi(it.Value)

	if err != nil {
		return defaultInteger
	}

	return toInt
}

func (it *KeyValuePair) ValueDefInt() int {
	toInt, err := strconv.Atoi(it.Value)

	if err != nil {
		return constants.Zero
	}

	return toInt
}

func (it *KeyValuePair) ValueByte(defVal byte) byte {
	toInt, err := strconv.Atoi(it.Value)

	if err != nil || toInt > constants.MaxUnit8AsInt {
		return defVal
	}

	return byte(toInt)
}

func (it *KeyValuePair) ValueDefByte() byte {
	toInt, err := strconv.Atoi(it.Value)

	if err != nil || toInt > constants.MaxUnit8AsInt {
		return constants.Zero
	}

	return byte(toInt)
}

func (it *KeyValuePair) ValueFloat64(defVal float64) float64 {
	toFloat, err := strconv.ParseFloat(it.Value, bitsize.Of64)

	if err != nil {
		return defVal
	}

	return toFloat
}

func (it *KeyValuePair) ValueDefFloat64() float64 {
	return it.ValueFloat64(constants.Zero)
}

func (it *KeyValuePair) ValueValid() ValidValue {
	return ValidValue{
		Value:   it.Value,
		IsValid: true,
		Message: constants.EmptyString,
	}
}

func (it *KeyValuePair) ValueValidOptions(
	isValid bool,
	message string,
) ValidValue {
	return ValidValue{
		Value:   it.Value,
		IsValid: isValid,
		Message: message,
	}
}

func (it *KeyValuePair) Is(key, val string) bool {
	return it != nil && it.Key == key && it.Value == val
}

func (it *KeyValuePair) IsKey(key string) bool {
	return it != nil && it.Key == key
}

func (it *KeyValuePair) IsVal(val string) bool {
	return it != nil && it.Value == val
}

func (it *KeyValuePair) IsKeyValueAnyEmpty() bool {
	return it == nil || it.Key == "" || it.Value == ""
}

// FormatString
//
//	FirstItem %v is key and next one is value
func (it *KeyValuePair) FormatString(format string) string {
	return fmt.Sprintf(
		format,
		it.Key,
		it.Value,
	)
}

func (it *KeyValuePair) String() string {
	return fmt.Sprintf(
		keyValuePrintFormat,
		it.Key,
		it.Value,
	)
}

func (it *KeyValuePair) Clear() {
	if it == nil {
		return
	}

	it.Key = ""
	it.Value = ""
}

func (it *KeyValuePair) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
}
