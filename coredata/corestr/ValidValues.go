package corestr

import (
	"gitlab.com/auk-go/core/conditional"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/defaultcapacity"
	"gitlab.com/auk-go/core/internal/strutilinternal"
)

type ValidValues struct {
	ValidValues []*ValidValue `json:"ValidValues,omitempty"`
}

func NewValidValuesUsingValues(values ...ValidValue) *ValidValues {
	if len(values) == 0 {
		return EmptyValidValues()
	}

	slice := make(
		[]*ValidValue,
		len(values))

	for i, value := range values {
		slice[i] = &value
	}

	return &ValidValues{
		ValidValues: slice,
	}
}

func NewValidValues(capacity int) *ValidValues {
	slice := make([]*ValidValue, 0, capacity)

	return &ValidValues{ValidValues: slice}
}

func EmptyValidValues() *ValidValues {
	return NewValidValues(
		constants.Zero)
}

func (it *ValidValues) Count() int {
	return it.Length()
}

func (it *ValidValues) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *ValidValues) LastIndex() int {
	return it.Length() - 1
}

func (it *ValidValues) HasIndex(
	index int,
) bool {
	return index != constants.InvalidNotFoundCase && it.LastIndex() >= index
}

func (it *ValidValues) Find(
	finder func(index int, valueValid *ValidValue) (foundItem *ValidValue, isFound, isBreak bool),
) []*ValidValue {
	length := it.Length()

	if length == 0 {
		return []*ValidValue{}
	}

	slice := make(
		[]*ValidValue,
		0,
		defaultcapacity.OfSearch(length))

	for i, item := range it.ValidValues {
		foundItem, isFound, isBreak := finder(i, item)

		if isFound && foundItem != nil {
			slice = append(slice, foundItem)
		}

		if isBreak {
			return slice
		}
	}

	return slice
}

func (it *ValidValues) SafeValueAt(index int) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	if it.HasIndex(index) {
		return it.ValidValues[index].Value
	}

	return constants.EmptyString
}

func (it *ValidValues) SafeValidValueAt(index int) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	if it.HasIndex(index) {
		validVal := it.ValidValues[index]

		return conditional.String(
			validVal.IsValid,
			validVal.Value,
			constants.EmptyString)
	}

	return constants.EmptyString
}

func (it *ValidValues) SafeValuesAtIndexes(indexes ...int) []string {
	requestLength := len(indexes)
	slice := make([]string, requestLength)

	if requestLength == 0 {
		return slice
	}

	for i, index := range indexes {
		slice[i] = it.SafeValueAt(index)
	}

	return slice
}

func (it *ValidValues) SafeValidValuesAtIndexes(indexes ...int) []string {
	requestLength := len(indexes)
	slice := make([]string, requestLength)

	if requestLength == 0 {
		return slice
	}

	for i, index := range indexes {
		slice[i] = it.SafeValidValueAt(index)
	}

	return slice
}

func (it *ValidValues) Strings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, val := range it.ValidValues {
		slice[i] = val.String()
	}

	return slice
}

func (it *ValidValues) FullStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, val := range it.ValidValues {
		slice[i] = val.FullString()
	}

	return slice
}

func (it *ValidValues) String() string {
	return strutilinternal.AnyToString(it.Strings())
}

func (it *ValidValues) Length() int {
	if it == nil {
		return 0
	}

	return len(it.ValidValues)
}

func (it *ValidValues) IsEmpty() bool {
	return it.Length() == 0
}

func (it *ValidValues) Add(val string) *ValidValues {
	it.ValidValues = append(it.ValidValues, &ValidValue{
		Value:   val,
		IsValid: true,
		Message: constants.EmptyString,
	})

	return it
}

func (it *ValidValues) AddFull(isValid bool, val, message string) *ValidValues {
	it.ValidValues = append(it.ValidValues, &ValidValue{
		Value:   val,
		IsValid: isValid,
		Message: message,
	})

	return it
}

func (it *ValidValues) ConcatNew(
	isCloneOnEmpty bool,
	validValuesCollection ...*ValidValues,
) *ValidValues {
	isEmpty := validValuesCollection == nil || len(validValuesCollection) == 0

	if isEmpty && isCloneOnEmpty {
		newValues := NewValidValues(it.Length())

		return newValues.AddsPtr(it.ValidValues...)
	}

	if isEmpty && !isCloneOnEmpty {
		return it
	}

	newValues := NewValidValues(it.Length())
	newValues.AddsPtr(it.ValidValues...)

	for _, validValues := range validValuesCollection {
		newValues.AddValidValues(validValues)
	}

	return newValues
}

func (it *ValidValues) AddValidValues(validValues *ValidValues) *ValidValues {
	if validValues == nil || validValues.IsEmpty() {
		return it
	}

	return it.AddsPtr(validValues.ValidValues...)
}

func (it *ValidValues) Adds(validValues ...ValidValue) *ValidValues {
	if len(validValues) == 0 {
		return it
	}

	for _, validValue := range validValues {
		it.ValidValues = append(
			it.ValidValues, &validValue)
	}

	return it
}

func (it *ValidValues) AddsPtr(validValues ...*ValidValue) *ValidValues {
	if len(validValues) == 0 {
		return it
	}

	for _, validValue := range validValues {
		it.ValidValues = append(
			it.ValidValues, validValue)
	}

	return it
}

func (it *ValidValues) AddHashsetMap(
	inputMap map[string]bool,
) *ValidValues {
	if inputMap == nil || len(inputMap) == 0 {
		return it
	}

	for val, isValid := range inputMap {
		it.ValidValues = append(it.ValidValues, &ValidValue{
			Value:   val,
			IsValid: isValid,
			Message: constants.EmptyString,
		})
	}

	return it
}

func (it *ValidValues) AddHashset(
	inputHashset *Hashset,
) *ValidValues {
	if inputHashset == nil || inputHashset.IsEmpty() {
		return it
	}

	return it.AddHashsetMap(inputHashset.items)
}

func (it *ValidValues) Hashmap() *Hashmap {
	length := it.Length()
	hashmap := New.Hashmap.Cap(length)

	if length == 0 {
		return hashmap
	}

	for _, keyVal := range it.ValidValues {
		hashmap.AddOrUpdate(keyVal.Value, keyVal.Message)
	}

	return hashmap
}

func (it *ValidValues) Map() map[string]string {
	hashmap := it.Hashmap()

	return hashmap.items
}
