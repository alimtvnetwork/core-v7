package converters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreappend"
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type anyItemConverter struct{}

func (it anyItemConverter) ToString(
	isIncludeFullName bool,
	any interface{},
) string {
	if any == nil {
		return ""
	}

	if isIncludeFullName {
		return fmt.Sprintf(
			constants.SprintFullPropertyNameValueFormat,
			any,
		)
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		any,
	)
}

func (it anyItemConverter) String(
	any interface{},
) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		any,
	)
}

func (it anyItemConverter) FullString(
	any interface{},
) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		any,
	)
}

func (it anyItemConverter) StringWithType(
	any interface{},
) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintPropertyValueWithTypeFormat,
		any,
		any,
	)
}

// ToSafeSerializedString
//
//	warning : on error swallows it
func (it anyItemConverter) ToSafeSerializedString(
	any interface{},
) string {
	if any == nil {
		return ""
	}

	switch casted := any.(type) {
	case []byte:
		return BytesTo.String(casted)
	case *[]byte:
		return BytesTo.PtrString(casted)
	}

	allBytes, _ := json.Marshal(any)

	return BytesTo.String(allBytes)
}

// ToSafeSerializedStringSprintValue
//
//	return value using %v
//
//	warning : on error swallows it
func (it anyItemConverter) ToSafeSerializedStringSprintValue(
	any interface{},
) string {
	value := it.ToSafeSerializedString(
		any,
	)

	return fmt.Sprintf(
		constants.SprintValueFormat,
		value,
	)
}

func (it anyItemConverter) ToStrings(
	isSkipOnNil bool,
	anyItem interface{},
) []string {
	if isSkipOnNil && anyItem == nil {
		return []string{}
	}

	reflectVal := reflect.ValueOf(anyItem)

	anyItems := reflectinternal.Converter.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal,
	)

	return it.ItemsToStringsSkipOnNil(anyItems)
}

func (it anyItemConverter) ToStringsUsingProcessor(
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

func (it anyItemConverter) ToStringsUsingSimpleProcessor(
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

func (it anyItemConverter) ToValueString(
	any interface{},
) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		any,
	)
}

func (it anyItemConverter) ToValueStringWithType(
	any interface{},
) string {
	if any == nil {
		return fmt.Sprintf(
			constants.SprintNilValueTypeInParenthesisFormat,
			any,
		)
	}

	return fmt.Sprintf(
		constants.SprintValueWithTypeFormat,
		any,
		any,
	)
}

func (it anyItemConverter) ToAnyItems(
	isSkipOnNil bool,
	anyItem interface{},
) []interface{} {
	if isSkipOnNil && anyItem == nil {
		return []interface{}{}
	}

	reflectVal := reflect.ValueOf(anyItem)

	return reflectinternal.Converter.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal,
	)
}

func (it anyItemConverter) ToNonNullItems(
	isSkipOnNil bool,
	anyItem interface{},
) []interface{} {
	if isSkipOnNil && anyItem == nil || reflectinternal.Is.Null(anyItem) {
		return []interface{}{}
	}

	reflectVal := reflect.ValueOf(anyItem)

	return reflectinternal.Converter.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal,
	)
}

func (it anyItemConverter) ItemsToStringsSkipOnNil(
	anyItems ...interface{},
) []string {
	return coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		nil,
		nil,
		anyItems...,
	)
}

func (it anyItemConverter) ItemsJoin(
	joiner string,
	anyItems ...interface{},
) string {
	if anyItems == nil {
		return constants.EmptyString
	}

	anyStrings := it.ItemsToStringsSkipOnNil(anyItems...)

	return strings.Join(anyStrings, joiner)
}

func (it anyItemConverter) ToItemsThenJoin(
	isSkipOnNil bool,
	joiner string,
	anySlice interface{},
) string {
	if anySlice == nil {
		return constants.EmptyString
	}

	anyStrings := it.ToStrings(
		isSkipOnNil,
		anySlice,
	)

	return strings.Join(
		anyStrings,
		joiner,
	)
}

func (it anyItemConverter) ToFullNameValueString(
	any interface{},
) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintFullPropertyNameValueFormat,
		any,
	)
}

// ToPrettyJson
//
// Warning:
//
//	swallows error
func (it anyItemConverter) ToPrettyJson(
	anyItem interface{},
) string {
	if anyItem == nil {
		return ""
	}

	allBytes, err := json.Marshal(anyItem)

	if err != nil || len(allBytes) == 0 {
		return ""
	}

	var prettyJSON bytes.Buffer

	json.Indent(
		&prettyJSON,
		allBytes,
		constants.EmptyString,
		constants.Tab,
	)

	return prettyJSON.String()
}

// Bytes
//
// ## Steps:
//   - If already in  []byte then return as is.
//   - If already in *[]byte then return as []byte without pointer by checking if not null.
//   - If already in  string then return as []byte(string).
//   - For rest of the cases, convert to json using Marshal and then returns the bytes
//
// Panic if json marshal has error.
func (it anyItemConverter) Bytes(anyItem interface{}) []byte {
	switch expectedAs := anyItem.(type) {
	case []byte:
		if expectedAs == nil {
			return []byte{}
		}

		return expectedAs
	case *[]byte:
		if expectedAs == nil || *expectedAs == nil {
			return []byte{}
		}

		return *expectedAs
	case string:
		return []byte(expectedAs)
	default:
		toBytes, err := json.Marshal(expectedAs)

		if err != nil {
			panic(err)
		}

		return toBytes
	}
}

// ValueString
//
// If nil then returns ""
// Or, returns %v of the value given.
func (it anyItemConverter) ValueString(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem,
	)
}

// SmartString
//
//   - If nil return ""
//   - If string return just returns
//   - Or, else return %v of value
func (it anyItemConverter) SmartString(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	return convertinteranl.AnyTo.SmartString(anyItem)
}
