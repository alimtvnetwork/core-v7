package coreonce

import (
	"encoding/json"
	"fmt"
	"sort"

	"gitlab.com/auk-go/core/constants"
)

type IntegersOnce struct {
	innerData       []int
	initializerFunc func() []int
	isInitialized   bool
	sortedValues    []int
}

func NewIntegersOnce(initializerFunc func() []int) IntegersOnce {
	return IntegersOnce{
		initializerFunc: initializerFunc,
	}
}

func NewIntegersOncePtr(initializerFunc func() []int) *IntegersOnce {
	return &IntegersOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *IntegersOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *IntegersOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = true

	return json.Unmarshal(data, &it.innerData)
}

func (it *IntegersOnce) Value() []int {
	if it.isInitialized {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

// IsEmpty returns true if zero
func (it *IntegersOnce) IsEmpty() bool {
	return len(it.Value()) == 0
}

func (it *IntegersOnce) IsZero() bool {
	return len(it.Value()) == 0
}

// Sorted
//
//  Warning : Current values will be mutated,
//  so better to make a clone of it.
func (it *IntegersOnce) Sorted() []int {
	if it.sortedValues != nil {
		return it.sortedValues
	}

	it.sortedValues = it.Value()
	sort.Ints(it.sortedValues)

	return it.sortedValues
}

func (it *IntegersOnce) RangesMap() map[int]int {
	values := it.Value()

	if len(values) == 0 {
		return map[int]int{}
	}

	newMap := make(map[int]int, len(values))

	for i, value := range values {
		newMap[value] = i
	}

	return newMap
}

func (it *IntegersOnce) RangesBoolMap() map[int]bool {
	values := it.Value()

	if len(values) == 0 {
		return map[int]bool{}
	}

	newMap := make(map[int]bool, len(values))

	for _, value := range values {
		newMap[value] = true
	}

	return newMap
}

func (it *IntegersOnce) UniqueMap() map[int]bool {
	values := it.Value()

	if len(values) == 0 {
		return map[int]bool{}
	}

	newMap := make(map[int]bool, len(values))

	for _, value := range values {
		newMap[value] = true
	}

	return newMap
}

func (it *IntegersOnce) Serialize() ([]byte, error) {
	values := it.Value()

	return json.Marshal(values)
}

func (it *IntegersOnce) IsEqual(integerItems ...int) bool {
	if it == nil && integerItems == nil {
		return true
	}

	currentItems := it.Value()
	if currentItems == nil && integerItems == nil {
		return true
	}

	if currentItems == nil || integerItems == nil {
		return false
	}

	if len(currentItems) != len(integerItems) {
		return false
	}

	for i, item := range currentItems {
		if item != integerItems[i] {
			return false
		}
	}

	return true
}

func (it *IntegersOnce) String() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.Value())
}

func (it *IntegersOnce) Values() []int {
	return it.Value()
}

func (it *IntegersOnce) Execute() []int {
	return it.Value()
}

func (it *IntegersOnce) Integers() []int {
	return it.Value()
}

func (it *IntegersOnce) Slice() []int {
	return it.Value()
}

func (it *IntegersOnce) List() []int {
	return it.Value()
}
