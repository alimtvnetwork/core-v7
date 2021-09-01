package corestr

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/core/internal/utilstringinternal"
)

type KeyValueCollection struct {
	KeyValuePairs []*KeyValuePair `json:"KeyValuePairs,omitempty"`
}

func (it *KeyValueCollection) Count() int {
	return it.Length()
}

func (it *KeyValueCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *KeyValueCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *KeyValueCollection) HasIndex(
	index int,
) bool {
	return index != constants.InvalidNotFoundCase && it.LastIndex() >= index
}

func (it *KeyValueCollection) Find(
	finder func(index int, currentKeyVal *KeyValuePair) (foundItem *KeyValuePair, isFound, isBreak bool),
) []*KeyValuePair {
	length := it.Length()

	if length == 0 {
		return []*KeyValuePair{}
	}

	slice := make(
		[]*KeyValuePair,
		0,
		defaultcapacity.OfSearch(length))

	for i, item := range it.KeyValuePairs {
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

func (it *KeyValueCollection) SafeValueAt(index int) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	if it.HasIndex(index) {
		return it.KeyValuePairs[index].Value
	}

	return constants.EmptyString
}

func (it *KeyValueCollection) SafeValuesAtIndexes(indexes ...int) []string {
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

func (it *KeyValueCollection) Strings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, keyVal := range it.KeyValuePairs {
		slice[i] = keyVal.String()
	}

	return slice
}

func (it *KeyValueCollection) StringsUsingFormat(
	format string,
) []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, keyVal := range it.KeyValuePairs {
		slice[i] = keyVal.FormatString(format)
	}

	return slice
}

func (it *KeyValueCollection) String() string {
	return utilstringinternal.AnyToString(it.Strings())
}

func (it *KeyValueCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.KeyValuePairs)
}

func (it *KeyValueCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *KeyValueCollection) Add(key, val string) *KeyValueCollection {
	it.KeyValuePairs = append(it.KeyValuePairs, &KeyValuePair{
		Key:   key,
		Value: val,
	})

	return it
}

func (it *KeyValueCollection) Adds(keyValues ...KeyValuePair) *KeyValueCollection {
	if len(keyValues) == 0 {
		return it
	}

	for _, keyVal := range keyValues {
		it.KeyValuePairs = append(it.KeyValuePairs, &KeyValuePair{
			Key:   keyVal.Key,
			Value: keyVal.Value,
		})
	}

	return it
}

func (it *KeyValueCollection) AddMap(
	inputMap map[string]string,
) *KeyValueCollection {
	if inputMap == nil || len(inputMap) == 0 {
		return it
	}

	for key, val := range inputMap {
		it.KeyValuePairs = append(it.KeyValuePairs, &KeyValuePair{
			Key:   key,
			Value: val,
		})
	}

	return it
}

func (it *KeyValueCollection) AddHashsetMap(
	inputMap map[string]bool,
) *KeyValueCollection {
	if inputMap == nil || len(inputMap) == 0 {
		return it
	}

	for key := range inputMap {
		it.KeyValuePairs = append(it.KeyValuePairs, &KeyValuePair{
			Key:   key,
			Value: key,
		})
	}

	return it
}

func (it *KeyValueCollection) AddHashset(
	inputHashset *Hashset,
) *KeyValueCollection {
	if inputHashset == nil || inputHashset.IsEmpty() {
		return it
	}

	for key := range inputHashset.items {
		it.KeyValuePairs = append(it.KeyValuePairs, &KeyValuePair{
			Key:   key,
			Value: key,
		})
	}

	return it
}

func (it *KeyValueCollection) AddsHashmap(
	hashmap *Hashmap,
) *KeyValueCollection {
	if hashmap == nil || hashmap.IsEmpty() {
		return it
	}

	for key, val := range hashmap.items {
		it.KeyValuePairs = append(it.KeyValuePairs, &KeyValuePair{
			Key:   key,
			Value: val,
		})
	}

	return it
}

func (it *KeyValueCollection) Hashmap() *Hashmap {
	length := it.Length()
	hashmap := NewHashmap(length)

	if length == 0 {
		return hashmap
	}

	for _, keyVal := range it.KeyValuePairs {
		hashmap.AddOrUpdate(keyVal.Key, keyVal.Value)
	}

	return hashmap
}

func (it *KeyValueCollection) Map() map[string]string {
	hashmap := it.Hashmap()

	return hashmap.items
}

func (it *KeyValueCollection) AddsHashmaps(
	hashmaps ...*Hashmap,
) *KeyValueCollection {
	if hashmaps == nil || len(hashmaps) == 0 {
		return it
	}

	for _, hashmap := range hashmaps {
		it.AddsHashmap(hashmap)
	}

	return it
}
