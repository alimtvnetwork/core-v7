package enumimpl

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

type KeyAnyVal struct {
	Key      string
	AnyValue interface{}
}

func (it KeyAnyVal) KeyString() string {
	return it.Key
}

func (it KeyAnyVal) AnyVal() interface{} {
	return it.AnyValue
}

func (it KeyAnyVal) AnyValString() string {
	return convAnyValToString(it.AnyValue)
}

func (it KeyAnyVal) WrapKey() string {
	return toJsonName(it.Key)
}

func (it KeyAnyVal) WrapValue() string {
	return toJsonName(it.AnyValue)
}

func (it KeyAnyVal) IsString() bool {
	_, isString := it.AnyValue.(string)

	if isString {
		return true
	}

	castInt, isInt := it.AnyValue.(int)

	if isInt {
		return castInt == constants.MinInt
	}

	return false
}

func (it KeyAnyVal) ValInt() int {
	return ConvEnumAnyValToInteger(it.AnyValue)
}

func (it KeyAnyVal) KeyValInteger() KeyValInteger {
	return KeyValInteger{
		Key:          it.Key,
		ValueInteger: it.ValInt(),
	}
}

func (it KeyAnyVal) String() string {
	if it.IsString() {
		// stringer
		return fmt.Sprintf(
			constants.StringEnumNameValueFormat,
			it.Key,
		)
	}

	return fmt.Sprintf(
		constants.EnumNameValueFormat,
		it.Key,
		it.AnyValue)
}
