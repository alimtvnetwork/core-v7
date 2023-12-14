package coreonce

import (
	"encoding/json"
	"sort"
	"sync"

	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/csvinternal"
)

type StringsOnce struct {
	innerData       []string
	mapOnce         map[string]bool
	initializerFunc func() []string
	isInitialized   bool
	sortedValues    []string
	sync.Mutex
}

func NewStringsOnce(initializerFunc func() []string) StringsOnce {
	return StringsOnce{
		initializerFunc: initializerFunc,
	}
}

func NewStringsOncePtr(initializerFunc func() []string) *StringsOnce {
	return &StringsOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *StringsOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *StringsOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = true

	return json.Unmarshal(data, &it.innerData)
}

func (it *StringsOnce) Strings() []string {
	return it.Value()
}

func (it *StringsOnce) SafeStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return it.Value()
}

func (it *StringsOnce) List() []string {
	return it.Value()
}

func (it *StringsOnce) Values() []string {
	return it.Value()
}

func (it *StringsOnce) ValuesPtr() *[]string {
	values := it.Value()

	return &values
}

func (it *StringsOnce) Value() []string {
	if it.isInitialized == true {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

func (it *StringsOnce) Execute() []string {
	return it.Value()
}

func (it *StringsOnce) Length() int {
	values := it.Value()

	if values == nil {
		return 0
	}

	return len(values)
}

func (it *StringsOnce) HasAnyItem() bool {
	return !it.IsEmpty()
}

// IsEmpty returns true if zero
func (it *StringsOnce) IsEmpty() bool {
	if it == nil || it.initializerFunc == nil {
		return true
	}

	values := it.Value()

	return values == nil || len(values) == 0
}

func (it *StringsOnce) HasAll(searchTerms ...string) bool {
	for _, term := range searchTerms {
		if !it.IsContains(term) {
			return false
		}
	}

	return true
}

func (it *StringsOnce) UniqueMapLock() map[string]bool {
	it.Lock()
	defer it.Unlock()

	return it.UniqueMap()
}

func (it *StringsOnce) UniqueMap() map[string]bool {
	if it.mapOnce != nil {
		return it.mapOnce
	}

	values := it.Values()

	if values == nil {
		return map[string]bool{}
	}

	hashset := make(map[string]bool, len(values))

	for _, item := range values {
		hashset[item] = true
	}

	it.mapOnce = hashset

	return it.mapOnce
}

func (it *StringsOnce) Has(search string) bool {
	return it.IsContains(search)
}

func (it *StringsOnce) IsContains(search string) bool {
	values := it.Values()

	for _, s := range values {
		if s == search {
			return true
		}
	}

	return false
}

func (it *StringsOnce) CsvLines() []string {
	return csvinternal.StringsToCsvStringsDefault(
		it.List()...)
}

func (it *StringsOnce) CsvOptions() string {
	return csvinternal.StringsToStringDefault(it.Value()...)
}

func (it *StringsOnce) Csv() string {
	return it.CsvOptions()
}

// Sorted
//
//	Warning : Current values will be mutated,
//	so better to make a clone of it.
func (it *StringsOnce) Sorted() []string {
	if it.sortedValues != nil {
		return it.sortedValues
	}

	it.sortedValues = it.Value()
	sort.Strings(it.sortedValues)

	return it.sortedValues
}

func (it *StringsOnce) RangesMap() map[string]int {
	values := it.Value()

	if len(values) == 0 {
		return map[string]int{}
	}

	newMap := make(map[string]int, len(values))

	for i, value := range values {
		newMap[value] = i
	}

	return newMap
}

func (it StringsOnce) Serialize() ([]byte, error) {
	values := it.Value()

	return json.Marshal(values)
}

func (it *StringsOnce) IsEqual(comparingItems ...string) bool {
	if it == nil && comparingItems == nil {
		return true
	}

	currentItems := it.Value()
	if currentItems == nil && comparingItems == nil {
		return true
	}

	if currentItems == nil || comparingItems == nil {
		return false
	}

	if len(currentItems) != len(comparingItems) {
		return false
	}

	for i, item := range currentItems {
		if item != comparingItems[i] {
			return false
		}
	}

	return true
}

func (it StringsOnce) JsonStringMust() string {
	marshalledJsonBytes, err := it.MarshalJSON()

	if err != nil {
		errcore.MarshallingFailedType.
			HandleUsingPanic(
				"StringsOnce failed to marshall."+err.Error(), it.innerData)

	}

	return string(marshalledJsonBytes)
}

func (it StringsOnce) String() string {
	return it.Csv()
}
