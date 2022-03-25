package converters

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreappend"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type anyItem struct{}

func (it anyItem) ToString(
	isIncludeFullName bool,
	any interface{},
) string {
	if any == nil {
		return ""
	}

	if isIncludeFullName {
		return fmt.Sprintf(
			constants.SprintFullPropertyNameValueFormat,
			any)
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		any)
}

// ToSafeSerializedString
//
//  warning : on error swallows it
func (it anyItem) ToSafeSerializedString(
	any interface{},
) string {
	if any == nil {
		return ""
	}

	switch casted := any.(type) {
	case []byte:
		return BytesToString(casted)
	case *[]byte:
		return BytesPtrToString(casted)
	}

	allBytes, _ := json.Marshal(any)

	return BytesToString(allBytes)
}

// ToSafeSerializedStringSprintValue
//
//  return value using %v
//
//  warning : on error swallows it
func (it anyItem) ToSafeSerializedStringSprintValue(
	any interface{},
) string {
	value := it.ToSafeSerializedString(
		any)

	return fmt.Sprintf(
		constants.SprintValueFormat,
		value)
}

func (it anyItem) ToStrings(
	isSkipOnNil bool,
	anyItem interface{},
) []string {
	if isSkipOnNil && anyItem == nil {
		return []string{}
	}

	reflectVal := reflect.ValueOf(anyItem)

	anyItems := reflectinternal.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal)

	return it.ManyToStringsSkipOnNil(anyItems)
}

func (it anyItem) ToStringsUsingProcessor(
	isSkipOnNil bool,
	processor func(index int, in interface{}) (out string, isTake, isBreak bool),
	any interface{},
) []string {
	if any == nil {
		return []string{}
	}

	anyItems := it.ToAnyItems(isSkipOnNil, any)
	slice := make([]string, 0, len(anyItems))

	if len(anyItems) == 0 {
		return slice
	}

	for i, item := range anyItems {
		out, isTake, isBreak := processor(i, item)

		if isTake {
			slice = append(slice, out)
		}

		if isBreak {
			return slice
		}
	}

	return slice
}

func (it anyItem) ToStringsUsingSimpleProcessor(
	isSkipOnNil bool,
	simpleProcessor func(index int, in interface{}) (out string),
	any interface{},
) []string {
	if any == nil {
		return []string{}
	}

	anyItems := it.ToAnyItems(isSkipOnNil, any)
	slice := make([]string, len(anyItems))

	if len(anyItems) == 0 {
		return slice
	}

	for i, item := range anyItems {
		out := simpleProcessor(i, item)

		slice[i] = out
	}

	return slice
}

func (it anyItem) ToValueString(
	any interface{},
) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		any)
}

func (it anyItem) ToValueStringWithType(
	any interface{},
) string {
	if any == nil {
		return fmt.Sprintf(
			constants.SprintNilValueTypeInParenthesisFormat,
			any)
	}

	return fmt.Sprintf(
		constants.SprintValueWithTypeFormat,
		any,
		any)
}

func (it anyItem) ToAnyItems(
	isSkipOnNil bool,
	anyItem interface{},
) []interface{} {
	if isSkipOnNil && anyItem == nil {
		return []interface{}{}
	}

	reflectVal := reflect.ValueOf(anyItem)

	return reflectinternal.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal)
}

func (it anyItem) ToNonNullItems(
	isSkipOnNil bool,
	anyItem interface{},
) []interface{} {
	if isSkipOnNil && anyItem == nil || reflectinternal.IsNull(anyItem) {
		return []interface{}{}
	}

	reflectVal := reflect.ValueOf(anyItem)

	return reflectinternal.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal)
}

func (it anyItem) ManyToStringsSkipOnNil(
	anyItems ...interface{},
) []string {
	return coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		nil,
		nil,
		anyItems...)
}

func (it anyItem) ManyJoin(
	joiner string,
	anyItems ...interface{},
) string {
	if anyItems == nil {
		return constants.EmptyString
	}

	anyStrings := it.ManyToStringsSkipOnNil(anyItems...)

	return strings.Join(anyStrings, joiner)
}

func (it anyItem) ToItemsThenJoin(
	isSkipOnNil bool,
	joiner string,
	anySlice interface{},
) string {
	if anySlice == nil {
		return constants.EmptyString
	}

	anyStrings := it.ToStrings(
		isSkipOnNil,
		anySlice)

	return strings.Join(
		anyStrings,
		joiner)
}

func (it anyItem) ToFullNameValueString(
	any interface{},
) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintFullPropertyNameValueFormat,
		any)
}
