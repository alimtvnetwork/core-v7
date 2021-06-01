package coreonce

import (
	"encoding/json"
	"sync"

	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/core/msgtype"
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

func (receiver *StringsOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(receiver.Value())
}

func (receiver *StringsOnce) UnmarshalJSON(data []byte) error {
	receiver.isInitialized = issetter.True

	return json.Unmarshal(data, &receiver.innerData)
}

func (receiver *StringsOnce) Strings() *[]string {
	return receiver.Value()
}

func (receiver *StringsOnce) Values() *[]string {
	return receiver.Value()
}

func (receiver *StringsOnce) Value() *[]string {
	if receiver.isInitialized.IsTrue() {
		return receiver.innerData
	}

	receiver.innerData = receiver.initializerFunc()
	receiver.isInitialized = issetter.True

	return receiver.innerData
}

func (receiver *StringsOnce) Length() int {
	values := receiver.Value()

	if values == nil {
		return 0
	}

	return len(*values)
}

func (receiver *StringsOnce) HasAnyItem() bool {
	return !receiver.IsEmpty()
}

// IsEmpty returns true if zero
func (receiver *StringsOnce) IsEmpty() bool {
	values := receiver.Value()

	return values == nil || len(*values) == 0
}

func (receiver *StringsOnce) HasAll(searchTerms ...string) bool {
	for _, term := range searchTerms {
		if !receiver.IsContains(term) {
			return false
		}
	}

	return true
}

func (receiver *StringsOnce) UniqueMapLock() *map[string]bool {
	receiver.Lock()
	defer receiver.Unlock()

	return receiver.UniqueMap()
}

func (receiver *StringsOnce) UniqueMap() *map[string]bool {
	if receiver.mapOnce != nil {
		return receiver.mapOnce
	}

	values := receiver.Values()

	if values == nil {
		return &map[string]bool{}
	}

	hashset := make(map[string]bool, len(*values))

	for _, item := range *values {
		hashset[item] = true
	}

	receiver.mapOnce = &hashset

	return receiver.mapOnce
}

func (receiver *StringsOnce) Has(search string) bool {
	return receiver.IsContains(search)
}

func (receiver *StringsOnce) IsContains(search string) bool {
	for _, s := range *receiver.innerData {
		if s == search {
			return true
		}
	}

	return false
}

func (receiver *StringsOnce) CsvLines() *[]string {
	return simplewrap.DoubleQuoteWrapElements(
		receiver.Value(),
		false)
}

func (receiver *StringsOnce) CsvOptions(isSkipQuoteOnlyOnExistence bool) string {
	return converters.StringsToCsv(receiver.Value(), isSkipQuoteOnlyOnExistence)
}

func (receiver *StringsOnce) Csv() string {
	return receiver.CsvOptions(false)
}

func (receiver *StringsOnce) StringJsonMust() string {
	marshalledJsonBytes, err := receiver.MarshalJSON()

	if err != nil {
		msgtype.MarshallingFailed.
			HandleUsingPanic(
				"StringsOnce failed to marshall."+err.Error(), receiver.innerData)

	}

	return string(marshalledJsonBytes)
}

func (receiver *StringsOnce) String() string {
	return receiver.Csv()
}
