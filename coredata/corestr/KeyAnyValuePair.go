package corestr

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type KeyAnyValuePair struct {
	Key         string
	valueString SimpleStringOnce
	Value       interface{}
}

func (it *KeyAnyValuePair) KeyName() string {
	return it.Key
}

func (it *KeyAnyValuePair) VariableName() string {
	return it.Key
}

func (it *KeyAnyValuePair) ValueAny() interface{} {
	return it.Value
}

func (it *KeyAnyValuePair) IsVariableNameEqual(name string) bool {
	return it.Key == name
}

func (it KeyAnyValuePair) SerializeMust() (jsonBytes []byte) {
	return corejson.NewPtr(it).RawMust()
}

func (it *KeyAnyValuePair) Compile() string {
	return it.String()
}

func (it *KeyAnyValuePair) IsValueNull() bool {
	return it == nil || reflectinternal.Is.Null(it.Value)
}

func (it *KeyAnyValuePair) HasNonNull() bool {
	return it != nil && reflectinternal.Is.Defined(it.Value)
}

func (it *KeyAnyValuePair) HasValue() bool {
	return it != nil && reflectinternal.Is.Defined(it.Value)
}

func (it *KeyAnyValuePair) IsValueEmptyString() bool {
	return it == nil || it.ValueString() == ""
}

func (it *KeyAnyValuePair) IsValueWhitespace() bool {
	return it == nil || strings.TrimSpace(it.ValueString()) == ""
}

func (it *KeyAnyValuePair) ValueString() string {
	if it.valueString.IsInitialized() {
		return it.valueString.String()
	}

	if it.HasNonNull() {
		valueString := fmt.Sprintf(constants.SprintValueFormat, it.Value)

		return it.
			valueString.
			GetSetOnce(valueString)
	}

	return it.
		valueString.
		GetOnce()
}

func (it *KeyAnyValuePair) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *KeyAnyValuePair) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*KeyAnyValuePair, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *KeyAnyValuePair) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *KeyAnyValuePair {
	parsed, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return parsed
}

func (it KeyAnyValuePair) Json() corejson.Result {
	return corejson.New(it)
}

func (it *KeyAnyValuePair) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *KeyAnyValuePair) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *KeyAnyValuePair) AsJsoner() corejson.Jsoner {
	return it
}

func (it *KeyAnyValuePair) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *KeyAnyValuePair) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it KeyAnyValuePair) String() string {
	return fmt.Sprintf(
		keyValuePrintFormat,
		it.Key,
		it.Value,
	)
}

func (it *KeyAnyValuePair) Clear() {
	if it == nil {
		return
	}

	it.Key = ""
	it.Value = nil
	it.valueString.Dispose()
}

func (it *KeyAnyValuePair) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
}
