package coreonce

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
)

type MapStringStringOnce struct {
	innerData                                          map[string]string
	initializerFunc                                    func() map[string]string
	isInitialized                                      bool
	compiledStrings                                    []string
	allKeys, allValues, allKeysSorted, allValuesSorted []string
	sync.RWMutex
}

func NewMapStringStringOnce(initializerFunc func() map[string]string) MapStringStringOnce {
	return MapStringStringOnce{
		initializerFunc: initializerFunc,
	}
}

func NewMapStringStringOncePtr(initializerFunc func() map[string]string) *MapStringStringOnce {
	return &MapStringStringOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *MapStringStringOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *MapStringStringOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = true

	return json.Unmarshal(data, &it.innerData)
}

func (it *MapStringStringOnce) Strings() []string {
	if it.compiledStrings != nil {
		return it.compiledStrings
	}

	listMap := it.ItemsMap()

	if len(listMap) == 0 {
		return []string{}
	}

	allKeyValues := make([]string, len(listMap))
	index := 0
	for key, value := range listMap {
		allKeyValues[index] = fmt.Sprintf(
			constants.KeyValJsonFormat,
			key,
			value)

		index++
	}

	it.compiledStrings = allKeyValues

	return it.compiledStrings
}

func (it *MapStringStringOnce) List() map[string]string {
	return it.Value()
}

func (it *MapStringStringOnce) ItemsMap() map[string]string {
	return it.Value()
}

func (it *MapStringStringOnce) Values() map[string]string {
	return it.Value()
}

func (it *MapStringStringOnce) ValuesPtr() *map[string]string {
	values := it.Value()

	return &values
}

func (it *MapStringStringOnce) Value() map[string]string {
	if it.isInitialized == true {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

func (it *MapStringStringOnce) Execute() map[string]string {
	return it.Value()
}

func (it *MapStringStringOnce) Length() int {
	values := it.Value()

	if values == nil {
		return 0
	}

	return len(values)
}

func (it *MapStringStringOnce) HasAnyItem() bool {
	return !it.IsEmpty()
}

// IsEmpty returns true if zero
func (it *MapStringStringOnce) IsEmpty() bool {
	if it == nil || it.initializerFunc == nil {
		return true
	}

	values := it.Value()

	return values == nil || len(values) == 0
}

func (it *MapStringStringOnce) HasAll(searchTerms ...string) bool {
	for _, term := range searchTerms {
		isMissing := !it.IsContains(term)

		if isMissing {
			return false
		}
	}

	return true
}

func (it *MapStringStringOnce) Has(search string) bool {
	return it.IsContains(search)
}

func (it *MapStringStringOnce) IsContains(search string) bool {
	itemsMap := it.Values()
	_, has := itemsMap[search]

	return has
}

func (it *MapStringStringOnce) IsMissing(search string) bool {
	itemsMap := it.Values()
	_, has := itemsMap[search]

	return !has
}

func (it *MapStringStringOnce) GetValue(key string) (val string) {
	itemsMap := it.Values()

	return itemsMap[key]
}

func (it *MapStringStringOnce) GetValueWithStatus(key string) (val string, hasItem bool) {
	itemsMap := it.Values()
	val, hasItem = itemsMap[key]

	return val, hasItem
}

func (it *MapStringStringOnce) AllKeys() []string {
	if it.allKeys != nil {
		return it.allKeys
	}

	listMap := it.ItemsMap()

	if len(listMap) == 0 {
		return []string{}
	}

	allKeys := make([]string, len(listMap))
	index := 0
	for key := range listMap {
		allKeys[index] = key
		index++
	}

	it.allKeys = allKeys

	return it.allKeys
}

func (it *MapStringStringOnce) AllValues() []string {
	if it.allValues != nil {
		return it.allValues
	}

	listMap := it.ItemsMap()

	if len(listMap) == 0 {
		return []string{}
	}

	allValues := make([]string, len(listMap))
	index := 0
	for _, value := range listMap {
		allValues[index] = value
		index++
	}

	it.allValues = allValues

	return it.allValues
}

func (it *MapStringStringOnce) AllKeysSorted() []string {
	if it.allKeysSorted != nil {
		return it.allKeysSorted
	}

	listMap := it.ItemsMap()
	if len(listMap) == 0 {
		return []string{}
	}

	allKeys := make([]string, len(listMap))
	index := 0
	for key := range listMap {
		allKeys[index] = key
		index++
	}

	sort.Strings(allKeys)
	it.allKeysSorted = allKeys

	return it.allKeysSorted
}

func (it *MapStringStringOnce) AllValuesSorted() []string {
	if it.allValuesSorted != nil {
		return it.allValuesSorted
	}

	listMap := it.ItemsMap()

	if len(listMap) == 0 {
		return []string{}
	}

	allValues := make([]string, len(listMap))
	index := 0
	for _, value := range listMap {
		allValues[index] = value
		index++
	}

	sort.Strings(allValues)
	it.allValuesSorted = allValues

	return it.allValuesSorted
}

func (it MapStringStringOnce) Serialize() ([]byte, error) {
	values := it.Value()

	return json.Marshal(values)
}

func (it *MapStringStringOnce) IsEqual(rightMap map[string]string) bool {
	if it == nil && rightMap == nil {
		return true
	}

	currentItems := it.Value()
	if currentItems == nil && rightMap == nil {
		return true
	}

	if currentItems == nil || rightMap == nil {
		return false
	}

	if len(currentItems) != len(rightMap) {
		return false
	}

	for leftKey, leftVal := range currentItems {
		rightVal, hasRight := rightMap[leftKey]
		isMissing := !hasRight

		if isMissing {
			return false
		}

		if leftVal != rightVal {
			return false
		}
	}

	return true
}

func (it MapStringStringOnce) JsonStringMust() string {
	marshalledJsonBytes, err := it.MarshalJSON()

	if err != nil {
		errcore.MarshallingFailedType.
			HandleUsingPanic(
				"MapStringStringOnce failed to marshall."+err.Error(), it.innerData)

	}

	return string(marshalledJsonBytes)
}

func (it MapStringStringOnce) String() string {
	return strings.Join(it.Strings(), constants.CommaUnixNewLine)
}
