package coreonce

import (
	"encoding/json"
	"sync"

	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/core/simplewrap"
)

type StringsOnce struct {
	innerData       *[]string
	mapOnce         *map[string]bool
	initializerFunc func() *[]string
	isInitialized   issetter.Value
	sync.Mutex
}

func NewStringsOnce(initializerFunc func() *[]string) StringsOnce {
	return StringsOnce{
		initializerFunc: initializerFunc,
	}
}

func NewStringsOncePtr(initializerFunc func() *[]string) *StringsOnce {
	return &StringsOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *StringsOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *StringsOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = issetter.True

	return json.Unmarshal(data, &it.innerData)
}

func (it *StringsOnce) Strings() *[]string {
	return it.Value()
}

func (it *StringsOnce) SafeStrings() []string {
	items := it.Value()

	if items == nil || len(*items) == 0 {
		return []string{}
	}

	return *it.Value()
}

func (it *StringsOnce) List() []string {
	return *it.Value()
}

func (it *StringsOnce) Values() *[]string {
	return it.Value()
}

func (it *StringsOnce) Value() *[]string {
	if it.isInitialized.IsTrue() {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = issetter.True

	return it.innerData
}

func (it *StringsOnce) Length() int {
	values := it.Value()

	if values == nil {
		return 0
	}

	return len(*values)
}

func (it *StringsOnce) HasAnyItem() bool {
	return !it.IsEmpty()
}

// IsEmpty returns true if zero
func (it *StringsOnce) IsEmpty() bool {
	values := it.Value()

	return values == nil || len(*values) == 0
}

func (it *StringsOnce) HasAll(searchTerms ...string) bool {
	for _, term := range searchTerms {
		if !it.IsContains(term) {
			return false
		}
	}

	return true
}

func (it *StringsOnce) UniqueMapLock() *map[string]bool {
	it.Lock()
	defer it.Unlock()

	return it.UniqueMap()
}

func (it *StringsOnce) UniqueMap() *map[string]bool {
	if it.mapOnce != nil {
		return it.mapOnce
	}

	values := it.Values()

	if values == nil {
		return &map[string]bool{}
	}

	hashset := make(map[string]bool, len(*values))

	for _, item := range *values {
		hashset[item] = true
	}

	it.mapOnce = &hashset

	return it.mapOnce
}

func (it *StringsOnce) Has(search string) bool {
	return it.IsContains(search)
}

func (it *StringsOnce) IsContains(search string) bool {
	for _, s := range *it.innerData {
		if s == search {
			return true
		}
	}

	return false
}

func (it *StringsOnce) CsvLines() []string {
	return simplewrap.DoubleQuoteWrapElements(
		false,
		it.List()...)
}

func (it *StringsOnce) CsvOptions(isSkipQuoteOnlyOnExistence bool) string {
	return converters.StringsToCsvPtr(isSkipQuoteOnlyOnExistence, it.Value())
}

func (it *StringsOnce) Csv() string {
	return it.CsvOptions(false)
}

func (it *StringsOnce) JsonStringMust() string {
	marshalledJsonBytes, err := it.MarshalJSON()

	if err != nil {
		errcore.MarshallingFailedType.
			HandleUsingPanic(
				"StringsOnce failed to marshall."+err.Error(), it.innerData)

	}

	return string(marshalledJsonBytes)
}

func (it *StringsOnce) String() string {
	return it.Csv()
}
