package corestr

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/coresort/strsort"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/utilstringinternal"
)

type CollectionPtr struct {
	items []*string
	sync.Mutex
}

func (it *CollectionPtr) Count() int {
	return it.Length()
}

func (it *CollectionPtr) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *CollectionPtr) LastIndex() int {
	return it.Length() - 1
}

func (it *CollectionPtr) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *CollectionPtr) ListStringsPtr() *[]string {
	return it.SimpleListPtr()
}

func (it *CollectionPtr) ListStrings() []string {
	return *it.SimpleListPtr()
}

func (it *CollectionPtr) StringJSON() string {
	return it.Json().JsonString()
}

func (it *CollectionPtr) RemoveAt(index int) (isSuccess bool) {
	if !it.HasIndex(index) {
		return false
	}

	it.ChainRemoveAt(index)

	return true
}

func (it *CollectionPtr) Capacity() int {
	if it.items == nil {
		return 0
	}

	return cap(it.items)
}

func (it *CollectionPtr) Length() int {
	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

func (it *CollectionPtr) LengthLock() int {
	it.Lock()
	defer it.Unlock()

	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

//goland:noinspection GoVetCopyLock
func (it *CollectionPtr) IsEquals(
	anotherCollectionPtr CollectionPtr,
) bool {
	return it.IsEqualsWithSensitivePtr(
		&anotherCollectionPtr,
		true)
}

func (it *CollectionPtr) IsEqualsPtr(
	anotherCollectionPtr *CollectionPtr,
) bool {
	return it.IsEqualsWithSensitivePtr(
		anotherCollectionPtr,
		true)
}

func (it *CollectionPtr) IsEqualsWithSensitivePtr(
	anotherCollectionPtr *CollectionPtr,
	isCaseSensitive bool,
) bool {
	if anotherCollectionPtr == nil && it == nil {
		return true
	}

	if anotherCollectionPtr == nil || it == nil {
		return false
	}

	if it == anotherCollectionPtr {
		return true
	}

	if it.IsEmpty() && anotherCollectionPtr.IsEmpty() {
		return true
	}

	if it.IsEmpty() || anotherCollectionPtr.IsEmpty() {
		return false
	}

	if it.Length() != anotherCollectionPtr.Length() {
		return false
	}

	leftItems := it.items
	rightItems := anotherCollectionPtr.items

	if isCaseSensitive {
		for i, leftVal := range leftItems {
			if leftVal != rightItems[i] {
				return false
			}
		}

		return true
	}

	for i, leftVal := range leftItems {
		rightVal := rightItems[i]

		if leftVal == nil && rightVal == nil {
			continue
		}

		if leftVal == nil || rightVal == nil {
			return false
		}

		if !strings.EqualFold(*leftVal, *rightVal) {
			return false
		}
	}

	return true
}

func (it *CollectionPtr) IsEmptyLock() bool {
	it.Lock()
	defer it.Unlock()

	return it.Length() == 0
}

func (it *CollectionPtr) IsEmpty() bool {
	return it.Length() == 0
}

func (it *CollectionPtr) HasItems() bool {
	return it.Length() > 0
}

func (it *CollectionPtr) AddLock(str string) *CollectionPtr {
	it.Lock()
	defer it.Unlock()

	it.items = append(
		it.items,
		&str)

	return it
}

func (it *CollectionPtr) Add(str string) *CollectionPtr {
	it.items = append(
		it.items,
		&str)

	return it
}

func (it *CollectionPtr) AddNonEmptyWhitespace(str string) *CollectionPtr {
	if str == "" {
		return it
	}

	if utilstringinternal.IsEmptyOrWhitespace(str) {
		return it
	}

	it.items = append(
		it.items,
		&str)

	return it
}

func (it *CollectionPtr) AddIf(isAdd bool, addingStringPtr *string) *CollectionPtr {
	if !isAdd {
		return it
	}

	it.items = append(
		it.items,
		addingStringPtr)

	return it
}

func (it *CollectionPtr) AddIfMany(
	isAdd bool,
	addingStrings ...*string,
) *CollectionPtr {
	if !isAdd {
		return it
	}

	it.items = append(
		it.items,
		addingStrings...)

	return it
}

func (it *CollectionPtr) AddFunc(f func() *string) *CollectionPtr {
	it.items = append(
		it.items,
		f())

	return it
}

func (it *CollectionPtr) AddFuncErr(
	funcReturnsError func() (result *string, err error),
	errHandler func(errInput error),
) *CollectionPtr {
	r, err := funcReturnsError()

	if err != nil {
		errHandler(err)

		return it
	}

	it.items = append(
		it.items,
		r)

	return it
}

func (it *CollectionPtr) AddsLock(items ...string) *CollectionPtr {
	it.Lock()
	defer it.Unlock()

	for i := range items {
		it.items = append(
			it.items,
			&(items[i]))
	}

	return it
}

func (it *CollectionPtr) Adds(items ...string) *CollectionPtr {
	for i := range items {
		it.items = append(
			it.items,
			&(items[i]))
	}

	return it
}

func (it *CollectionPtr) AddHashmapsValues(
	hashmaps ...*Hashmap,
) *CollectionPtr {
	if hashmaps == nil {
		return it
	}

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for _, v := range hashmap.items {
			newV := v
			it.items = append(
				it.items,
				&newV)
		}
	}

	return it
}

func (it *CollectionPtr) AddHashmapsKeys(
	hashmaps ...*Hashmap,
) *CollectionPtr {
	if hashmaps == nil {
		return it
	}

	it.resizeForHashmaps(
		&hashmaps,
		constants.One)

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k := range hashmap.items {
			keyCopied := k
			it.items = append(
				it.items,
				&keyCopied)
		}
	}

	return it
}

func (it *CollectionPtr) resizeForHashmaps(
	hashmaps *[]*Hashmap,
	multiplier int,
) *CollectionPtr {
	if hashmaps == nil {
		return it
	}

	length := 0

	for _, hashmap := range *hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		length += hashmap.Length()
	}

	if !it.isResizeRequired(length) {
		return it
	}

	finalLength :=
		length*multiplier +
			length/2

	return it.AddCapacity(finalLength)
}

func (it *CollectionPtr) resizeForCollectionPtrs(
	collections *[]*CollectionPtr,
	multiplier int,
) *CollectionPtr {
	if collections == nil {
		return it
	}

	length := 0

	for _, collection := range *collections {
		if collection == nil || collection.IsEmpty() {
			continue
		}

		length += collection.Length()
	}

	if !it.isResizeRequired(length) {
		return it
	}

	finalLength :=
		length*multiplier +
			length/2

	return it.AddCapacity(finalLength)
}

func (it *CollectionPtr) resizeForItems(
	items []*string,
	multiplier int,
) *CollectionPtr {
	if items == nil {
		return it
	}

	length := len(items)
	if !it.isResizeRequired(length) {
		return it
	}

	finalLength :=
		length*multiplier +
			length/2

	return it.AddCapacity(finalLength)
}

func (it *CollectionPtr) isResizeRequired(
	length int,
) bool {
	if length < constants.ArbitraryCapacity200 {
		return false
	}

	windowLength := it.Capacity() - it.Length()
	if windowLength >= length {
		return false
	}

	return true
}

func (it *CollectionPtr) resizeForAnys(
	items *[]interface{},
	multiplier int,
) *CollectionPtr {
	if items == nil {
		return it
	}

	length := len(*items)
	if !it.isResizeRequired(length) {
		return it
	}

	finalLength :=
		length*multiplier +
			length/2

	return it.AddCapacity(finalLength)
}

func (it *CollectionPtr) AddHashmapsKeysValues(
	hashmaps ...*Hashmap,
) *CollectionPtr {
	if hashmaps == nil {
		return it
	}

	it.resizeForHashmaps(
		&hashmaps,
		constants.ArbitraryCapacity2)

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k, v := range hashmap.items {
			kc := k
			vc := v
			it.items = append(
				it.items,
				&kc)
			it.items = append(
				it.items,
				&vc)
		}
	}

	return it
}

func (it *CollectionPtr) AddHashmapsKeysValuesUsingFilter(
	filter IsKeyValueFilter,
	hashmaps ...*Hashmap,
) *CollectionPtr {
	if hashmaps == nil {
		return it
	}

	it.resizeForHashmaps(
		&hashmaps,
		constants.One)

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k, v := range hashmap.items {
			result, isAcceptable, isBreak := filter(KeyValuePair{
				Key:   k,
				Value: v,
			})

			if isAcceptable {
				it.items = append(
					it.items,
					&result)
			}

			if isBreak {
				return it
			}
		}
	}

	return it
}

func (it *CollectionPtr) AddPtr(str *string) *CollectionPtr {
	it.items = append(
		it.items,
		str)

	return it
}

func (it *CollectionPtr) AddPtrLock(str *string) *CollectionPtr {
	it.Lock()
	defer it.Unlock()

	it.items = append(
		it.items,
		str)

	return it
}

func (it *CollectionPtr) AddWithWgLock(
	str string,
	group *sync.WaitGroup,
) *CollectionPtr {
	it.Lock()
	defer it.Unlock()

	it.items = append(
		it.items,
		&str)

	group.Done()

	return it
}

func (it *CollectionPtr) AddsPtrLock(itemsPtr ...*string) *CollectionPtr {
	it.Lock()
	defer it.Unlock()

	for _, str := range itemsPtr {
		it.items = append(
			it.items,
			str)
	}

	return it
}

func (it *CollectionPtr) AddStringsPtrWgLock(
	stringItems *[]string,
	group *sync.WaitGroup,
) *CollectionPtr {
	it.Lock()
	defer it.Unlock()

	it.AddStringsPtr(stringItems)
	group.Done()

	return it
}

func (it *CollectionPtr) AddStringsPtr(stringItems *[]string) *CollectionPtr {
	for i := range *stringItems {
		it.items = append(
			it.items,
			&(*stringItems)[i])
	}

	return it
}

func (it *CollectionPtr) InsertItemsAt(index int, stringItems *[]string) *CollectionPtr {
	length := it.Length()
	isAtFirst := length == 0
	isAtLast := length-1 == index
	isAppendItems := isAtFirst || isAtLast

	if isAppendItems {
		return it.AddStringsPtr(stringItems)
	}

	pointerStrings := converters.StringsToPointerStrings(stringItems)

	// https://bit.ly/3pIDfRY
	it.items =
		append(
			it.items[:index],
			*pointerStrings...)

	it.items = append(
		it.items,
		it.items[index:]...)

	return it
}

func (it *CollectionPtr) ChainRemoveAt(index int) *CollectionPtr {
	it.items = append(
		it.items[:index],
		it.items[index+1:]...)

	return it
}

// RemoveItemsIndexes creates a new collection without the indexes mentioned.
//
// it is better to filter out than remove.
func (it *CollectionPtr) RemoveItemsIndexes(
	isIgnoreRemoveError bool,
	indexes ...int,
) *CollectionPtr {
	if isIgnoreRemoveError && indexes == nil {
		return it
	}

	return it.
		RemoveItemsIndexesPtr(isIgnoreRemoveError, indexes)
}

// RemoveItemsIndexesPtr creates a new collection without the indexes mentioned.
//
// it is better to filter out than remove.
func (it *CollectionPtr) RemoveItemsIndexesPtr(
	isIgnoreRemoveError bool,
	indexes []int,
) *CollectionPtr {
	length := it.Length()
	indexesLength := len(indexes)
	hasPossibleError := length == 0 && indexesLength > 0

	if hasPossibleError && !isIgnoreRemoveError {
		panic(errcore.CannotRemoveIndexesFromEmptyCollectionType)
	}

	if !isIgnoreRemoveError {
		errcore.PanicOnIndexOutOfRange(length, indexes)
	}

	if hasPossibleError {
		return it
	}

	newList := make([]*string, constants.Zero, it.Capacity())
	for i, s := range it.items {
		if coreindexes.HasIndex(indexes, i) {
			continue
		}

		newList = append(newList, s)
	}

	it.items = newList

	return it
}

func (it *CollectionPtr) AddPointerStringsPtrLock(
	pointerStringItems []*string,
) *CollectionPtr {
	it.Lock()
	defer it.Unlock()

	return it.
		AddPointerStrings(pointerStringItems...)
}

func (it *CollectionPtr) AddPointerStrings(
	pointerStringItems ...*string,
) *CollectionPtr {
	for i := range pointerStringItems {
		it.items = append(
			it.items,
			pointerStringItems[i])
	}

	return it
}

//goland:noinspection GoVetCopyLock
func (it *CollectionPtr) AppendCollectionPtr(
	anotherCollectionPtr CollectionPtr,
) *CollectionPtr {
	it.resizeForItems(
		anotherCollectionPtr.items,
		constants.One)

	it.items = append(
		it.items,
		anotherCollectionPtr.items...)

	return it
}

func (it *CollectionPtr) AppendCollectionPtrPtr(
	anotherCollectionPtr *CollectionPtr,
) *CollectionPtr {
	it.resizeForItems(
		anotherCollectionPtr.items,
		constants.One)

	it.items = append(
		it.items,
		anotherCollectionPtr.items...)

	return it
}

func (it *CollectionPtr) AppendCollectionPtrsPtr(
	anotherCollectionPtrsPtr ...*CollectionPtr,
) *CollectionPtr {
	if anotherCollectionPtrsPtr == nil {
		return it
	}

	it.resizeForCollectionPtrs(
		&anotherCollectionPtrsPtr,
		constants.One)

	capacitiesIncrease := 0
	for _, currentCollectionPtr := range anotherCollectionPtrsPtr {
		if currentCollectionPtr == nil || currentCollectionPtr.IsEmpty() {
			continue
		}

		capacitiesIncrease += currentCollectionPtr.Length()
	}

	it.AddCapacity(capacitiesIncrease)

	for _, currentCollectionPtr := range anotherCollectionPtrsPtr {
		if currentCollectionPtr == nil || currentCollectionPtr.IsEmpty() {
			continue
		}

		it.items = append(
			it.items,
			currentCollectionPtr.items...)
	}

	return it
}

// AppendAnysLock Continue on nil
func (it *CollectionPtr) AppendAnysLock(
	anys ...interface{},
) *CollectionPtr {
	if anys == nil {
		return it
	}

	it.resizeForAnys(
		&anys,
		constants.One)

	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)

		it.Lock()
		it.items = append(
			it.items,
			&anyStr)
		it.Unlock()
	}

	return it
}

// AppendAnys Continue on nil
func (it *CollectionPtr) AppendAnys(
	anys ...interface{},
) *CollectionPtr {
	if anys == nil {
		return it
	}

	it.resizeForAnys(
		&anys,
		constants.One)

	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(
			constants.SprintValueFormat,
			any,
		)

		it.items = append(
			it.items,
			&anyStr)
	}

	return it
}

// AppendAnysUsingFilter Skip on nil
func (it *CollectionPtr) AppendAnysUsingFilter(
	filter IsStringFilter,
	anys ...interface{},
) *CollectionPtr {
	if anys == nil {
		return it
	}

	it.resizeForAnys(
		&anys,
		constants.One)

	for i, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(
			constants.SprintValueFormat,
			any)

		result, isKeep, isBreak := filter(anyStr, i)

		if !isKeep {
			continue
		}

		it.items = append(
			it.items,
			&result)

		if isBreak {
			return it
		}
	}

	return it
}

// AppendAnysUsingFilterLock Skip on nil
func (it *CollectionPtr) AppendAnysUsingFilterLock(
	filter IsStringFilter,
	anys ...interface{},
) *CollectionPtr {
	if anys == nil {
		return it
	}

	it.resizeForAnys(
		&anys,
		constants.One)

	for i, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)
		result, isKeep, isBreak := filter(anyStr, i)

		if !isKeep {
			continue
		}

		it.Lock()
		it.items = append(
			it.items,
			&result)
		it.Unlock()

		if isBreak {
			return it
		}
	}

	return it
}

// AppendNonEmptyAnys Continue on nil
func (it *CollectionPtr) AppendNonEmptyAnys(
	anys ...interface{},
) *CollectionPtr {
	if anys == nil {
		return it
	}

	it.resizeForAnys(
		&anys,
		constants.One)

	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)
		if anyStr == "" {
			continue
		}

		it.items = append(
			it.items,
			&anyStr)
	}

	return it
}

// AddsPtr adds nil
func (it *CollectionPtr) AddsPtr(
	itemsPtr ...*string,
) *CollectionPtr {
	if itemsPtr == nil {
		return it
	}

	for _, str := range itemsPtr {
		it.items = append(
			it.items,
			str)
	}

	return it
}

func (it *CollectionPtr) AddsNonEmptyPtr(
	itemsPtr ...*string,
) *CollectionPtr {
	if itemsPtr == nil {
		return it
	}

	for _, str := range itemsPtr {
		if str == nil || *str == "" {
			continue
		}

		it.items = append(
			it.items,
			str)
	}

	return it
}

func (it *CollectionPtr) AddsNonEmptyPtrLock(
	itemsPtr ...*string,
) *CollectionPtr {
	if itemsPtr == nil {
		return it
	}

	for _, str := range itemsPtr {
		if str == nil || *str == "" {
			continue
		}

		it.Lock()
		it.items = append(
			it.items,
			str)
		it.Unlock()
	}

	return it
}

func (it *CollectionPtr) UniqueBoolMapLock(
	isAddEmptyStringOnNil bool,
) *map[string]bool {
	it.Lock()
	defer it.Unlock()

	return it.UniqueBoolMap(isAddEmptyStringOnNil)
}

// UniqueBoolMap continue on nil
func (it *CollectionPtr) UniqueBoolMap(
	isAddEmptyStringOnNil bool,
) *map[string]bool {
	respectiveMap := make(
		map[string]bool,
		it.Length())

	for _, item := range it.items {
		if item == nil && !isAddEmptyStringOnNil {
			continue
		} else if item == nil && isAddEmptyStringOnNil {
			respectiveMap[""] = true

			continue
		}

		//goland:noinspection GoNilness
		respectiveMap[*item] = true
	}

	return &respectiveMap
}

func (it *CollectionPtr) UniqueListPtr(
	isAddEmptyStringOnNil bool,
) *[]string {
	boolMap := it.UniqueBoolMap(isAddEmptyStringOnNil)
	list := make([]string, len(*boolMap))

	i := 0
	for str := range *boolMap {
		list[i] = str
		i++
	}

	return &list
}

func (it *CollectionPtr) UniqueListPtrLock(
	isAddEmptyStringOnNil bool,
) *[]string {
	it.Lock()
	defer it.Unlock()

	return it.UniqueListPtr(isAddEmptyStringOnNil)
}

func (it *CollectionPtr) UniqueListLock(
	isAddEmptyStringOnNil bool,
) []string {
	it.Lock()
	defer it.Unlock()

	return it.UniqueList(isAddEmptyStringOnNil)
}

func (it *CollectionPtr) UniqueList(
	isAddEmptyStringOnNil bool,
) []string {
	return *it.UniqueListPtr(isAddEmptyStringOnNil)
}

func (it *CollectionPtr) List() []*string {
	return it.items
}

// ListCopyPtrLock returns a copy of the items
//
// must return a slice
func (it *CollectionPtr) ListCopyPtrLock() []*string {
	it.Lock()
	defer it.Unlock()

	if it.IsEmpty() {
		return []*string{}
	}

	newSlice := make([]*string, it.Length())
	copy(newSlice, it.items)

	return newSlice
}

// FilterSimpleArray must return a slice
func (it *CollectionPtr) FilterSimpleArray(
	filter IsStringPointerFilter,
) *[]string {
	if it.IsEmpty() {
		return &([]string{})
	}

	list := make([]string, constants.Zero, it.Length())

	for i, element := range it.items {
		result, isKeep, isBreak := filter(element, i)

		if isKeep && result != nil {
			list = append(list, *result)
		}

		if isBreak {
			return &list
		}
	}

	return &list
}

// FilterSimpleArrayLock assumed wg is added to be done here.
//
// must return a slice
func (it *CollectionPtr) FilterSimpleArrayLock(
	filter IsStringPointerFilter,
	wg *sync.WaitGroup,
) *[]string {
	copyList := it.ListCopyPtrLock()
	length := len(copyList)
	if length == 0 {
		return &([]string{})
	}

	list := make([]string, constants.Zero, length)

	for i, element := range it.items {
		result, isKeep, isBreak := filter(element, i)

		if isKeep && result != nil {
			list = append(list, *result)
		}

		if isBreak {
			return &list
		}
	}

	wg.Done()

	return &list
}

// Filter must return a slice
func (it *CollectionPtr) Filter(
	filter IsStringPointerFilter,
) *[]*string {
	if it.IsEmpty() {
		return &([]*string{})
	}

	list := make([]*string, constants.Zero, it.Length())

	for i, element := range it.items {
		result, isKeep, isBreak := filter(element, i)

		if isKeep {
			list = append(list, result)
		}

		if isBreak {
			return &list
		}
	}

	return &list
}

// FilterLock must return a slice
func (it *CollectionPtr) FilterLock(
	filter IsStringPointerFilter,
) []*string {
	elements := it.ListCopyPtrLock()
	length := len(elements)

	if length == 0 {
		return elements
	}

	list := make([]*string, constants.Zero, length)

	for i, element := range elements {
		result, isKeep, isBreak := filter(element, i)

		if isKeep {
			list = append(list, result)
		}

		if isBreak {
			return list
		}
	}

	return list
}

// FilteredCollection must return a items
func (it *CollectionPtr) FilteredCollection(
	filter IsStringPointerFilter,
) *CollectionPtr {
	return New.CollectionPtr.StringsPtr(
		it.FilterSimpleArray(
			filter))
}

// FilteredCollectionLock must return a items
func (it *CollectionPtr) FilteredCollectionLock(
	filter IsStringPointerFilter,
) *CollectionPtr {
	it.Lock()
	defer it.Unlock()

	return New.CollectionPtr.StringsPtr(
		it.FilterSimpleArray(filter),
	)
}

// FilterPtrLock must return a slice
func (it *CollectionPtr) FilterPtrLock(
	filterPtr IsStringPointerFilter,
) []*string {
	elements := it.ListCopyPtrLock()
	length := len(elements)

	if length == 0 {
		return []*string{}
	}

	list := make(
		[]*string,
		constants.Zero,
		length)

	for i, element := range elements {
		result, isKeep, isBreak :=
			filterPtr(element, i)

		if isKeep {
			list = append(
				list,
				result)
		}

		if isBreak {
			return list
		}
	}

	return list
}

// FilterPtr must return a slice
func (it *CollectionPtr) FilterPtr(
	filterPtr IsStringPointerFilter,
) []*string {
	if it.IsEmpty() {
		return []*string{}
	}

	list := make([]*string, constants.Zero, it.Length())

	for i, element := range it.items {
		result, isKeep, isBreak := filterPtr(element, i)

		if isKeep {
			list = append(list, result)
		}

		if isBreak {
			return list
		}
	}

	return list
}

// GetAllExceptCollection
//
// Get all items except the mentioned ones in itemsCollection.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this collection
// Set B = itemsCollection given in parameters.
func (it *CollectionPtr) GetAllExceptCollection(
	itemsCollection *CollectionPtr,
) *[]*string {
	if itemsCollection == nil || itemsCollection.IsEmpty() {
		newItems := it.items

		return &newItems
	}

	finalList := make(
		[]*string,
		constants.Zero,
		it.Length())

	for _, item := range it.items {
		if itemsCollection.HasPtr(item) {
			continue
		}

		finalList = append(
			finalList,
			item)
	}

	return &finalList
}

// GetAllExcept Get all items except the mentioned ones.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this collection
// Set B = items given in parameters.
func (it *CollectionPtr) GetAllExcept(
	items []*string,
) *[]*string {
	if items == nil {
		newItems := it.items

		return &newItems
	}

	newCollection := New.CollectionPtr.Default(
		items)

	return it.GetAllExceptCollection(
		newCollection)
}

// NonEmptySimpleListPtr must return a slice
func (it *CollectionPtr) NonEmptySimpleListPtr() *[]string {
	if it.IsEmpty() {
		return &([]string{})
	}

	list := make([]string, constants.Zero, it.Length())

	for _, element := range it.items {
		if element == nil || *element == "" {
			continue
		}

		list = append(list, *element)
	}

	return &list
}

func (it *CollectionPtr) HashsetAsIs() *Hashset {
	if it.IsEmpty() {
		return New.Hashset.Empty()
	}

	return New.Hashset.PointerStrings(
		it.items,
	)
}

func (it *CollectionPtr) HashsetDoubleLength() *Hashset {
	if it.IsEmpty() {
		return New.Hashset.Empty()
	}

	return New.Hashset.PointerStringsPtrOption(
		it.Length()*2,
		false,
		&it.items,
	)
}

func (it *CollectionPtr) HashsetLock() *Hashset {
	items := it.ListCopyPtrLock()

	return New.Hashset.PointerStrings(
		items,
	)
}

func (it *CollectionPtr) SimpleList() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return *converters.PointerStringsToStrings(
		&it.items)
}

func (it *CollectionPtr) SimpleListPtr() *[]string {
	if it.IsEmpty() {
		return &[]string{}
	}

	return converters.PointerStringsToStrings(
		&it.items)
}

func (it *CollectionPtr) SimpleListPtrLock() *[]string {
	it.Lock()
	defer it.Unlock()

	return it.SimpleListPtr()
}

// ListPtr direct return pointer
func (it *CollectionPtr) ListPtr() []*string {
	return it.items
}

func (it *CollectionPtr) HasLock(str string) bool {
	it.Lock()
	defer it.Unlock()

	return it.Has(str)
}

func (it *CollectionPtr) Has(str string) bool {
	if it.IsEmpty() {
		return false
	}

	for _, element := range it.items {
		if element == nil {
			continue
		}

		if *element == str {
			return true
		}
	}

	return false
}

func (it *CollectionPtr) HasPtr(str *string) bool {
	if it.IsEmpty() {
		return false
	}

	for _, element := range it.items {
		if element == nil && str == nil {
			return true
		}

		if element == nil {
			continue
		}

		if *element == *str {
			return true
		}
	}

	return false
}

func (it *CollectionPtr) HasUsingSensitivity(
	str string, isCaseSensitive bool,
) bool {
	if it.IsEmpty() {
		return false
	}

	if isCaseSensitive {
		return it.Has(str)
	}

	for _, element := range it.items {
		if element == nil {
			continue
		}

		if strings.EqualFold(*element, str) {
			return true
		}
	}

	return false
}

func (it *CollectionPtr) HasAll(
	items ...string,
) bool {
	if it.IsEmpty() {
		return false
	}

	for _, element := range items {
		if !it.IsContainsPtr(&element) {
			return false
		}
	}

	return true
}

func (it *CollectionPtr) First() string {
	return *it.items[0]
}

func (it *CollectionPtr) Single() string {
	length := it.Length()
	if length != 1 {
		errcore.LengthShouldBeEqualToType.HandleUsingPanic("1", length)
	}

	return *it.items[0]
}

func (it *CollectionPtr) Last() string {
	length := it.Length()

	return *it.items[length-1]
}

func (it *CollectionPtr) LastOrDefault() string {
	length := it.Length()

	if length == 0 {
		return constants.EmptyString
	}

	return *it.items[length-1]
}

func (it *CollectionPtr) FirstOrDefault() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return *it.items[0]
}

// Take use One based index
func (it *CollectionPtr) Take(
	take int,
) *CollectionPtr {
	length := it.Length()

	if length <= take {
		return it
	}

	if take == 0 {
		return Empty.CollectionPtr()
	}

	list := it.items[:take]

	return New.CollectionPtr.Default(
		list,
	)
}

// Skip use One based index
func (it *CollectionPtr) Skip(
	skip int,
) *CollectionPtr {
	length := it.Length()

	if length < skip {
		errcore.
			LengthShouldBeEqualToType.
			HandleUsingPanic(
				"Length is lower than skip value. Skip:",
				skip)
	}

	if skip == 0 {
		return it
	}

	list := it.items[skip:]

	return New.CollectionPtr.Default(
		list,
	)
}

func (it *CollectionPtr) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *CollectionPtr) GetPagedCollection(
	eachPageSize int,
) *CollectionsOfCollectionPtr {
	length := it.Length()

	if length < eachPageSize {
		return New.CollectionsOfCollectionPtr.PointerStrings(
			it.items,
		)
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := New.CollectionsOfCollectionPtr.Cap(
		pagesPossibleCeiling)

	for i := 1; i <= pagesPossibleCeiling; i++ {
		pagedCollection := it.GetSinglePageCollection(
			eachPageSize, i)

		collectionOfCollection.Adds(
			pagedCollection)
	}

	return collectionOfCollection
}

// GetSinglePageCollection PageIndex is one based index. Should be above or equal 1
func (it *CollectionPtr) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *CollectionPtr {
	length := it.Length()

	if length < eachPageSize {
		return it
	}

	/**
	 * eachPageItems = 10
	 * pageIndex = 4
	 * skipItems = 10 * (4 - 1) = 30
	 */
	skipItems := eachPageSize * (pageIndex - 1)
	if skipItems < 0 {
		errcore.
			CannotBeNegativeIndexType.
			HandleUsingPanic(
				"pageIndex cannot be negative or zero.",
				pageIndex)
	}

	endingIndex := skipItems + eachPageSize

	if endingIndex > length {
		endingIndex = length
	}

	list := it.items[skipItems:endingIndex]

	return New.CollectionPtr.PointerStrings(
		list,
	)
}

func (it *CollectionPtr) IndexAt(
	index int,
) string {
	return *it.items[index]
}

func (it *CollectionPtr) SafePointerIndexAt(
	index int,
) *string {
	length := it.Length()
	if length-1 < index {
		return nil
	}

	return it.items[index]
}

func (it *CollectionPtr) SafePointerIndexAtUsingLength(
	length, index int,
) *string {
	if length-1 < index {
		return nil
	}

	return it.items[index]
}

func (it *CollectionPtr) SafeIndexAtUsingLength(
	defaultString string, length, index int,
) string {
	if length-1 < index {
		return defaultString
	}

	pointer := it.items[index]

	if pointer == nil {
		return defaultString
	}

	return *pointer
}

// SortedListAsc Creates new doesn't modify current collection
func (it *CollectionPtr) SortedListAsc() *[]string {
	if it.IsEmpty() {
		return &[]string{}
	}

	list := it.SimpleListPtr()
	sort.Strings(*list)

	return list
}

// SortedAsc mutates current collection
func (it *CollectionPtr) SortedAsc() *CollectionPtr {
	if it.IsEmpty() {
		return it
	}

	strsort.QuickPtr(&it.items)

	return it
}

// SortedAscLock mutates current collection
func (it *CollectionPtr) SortedAscLock() *CollectionPtr {
	if it.IsEmptyLock() {
		return it
	}

	it.Lock()
	defer it.Unlock()

	strsort.QuickPtr(&it.items)

	return it
}

// SortedListDsc Creates new one.
func (it *CollectionPtr) SortedListDsc() *[]string {
	list := it.SortedListAsc()
	length := len(*list)
	mid := length / 2

	for i := 0; i < mid; i++ {
		temp := (*list)[i]
		(*list)[i] = (*list)[length-1-i]
		(*list)[length-1-i] = temp
	}

	return list
}

// SortedDsc mutates itself.
func (it *CollectionPtr) SortedDsc() *CollectionPtr {
	list := it.items
	length := len(list)
	mid := length / 2

	for i := 0; i < mid; i++ {
		temp := list[i]
		list[i] = list[length-1-i]
		list[length-1-i] = temp
	}

	return it
}

func (it *CollectionPtr) IsContainsPtr(
	item *string,
) bool {
	if item == nil || it.IsEmpty() {
		return false
	}

	for _, element := range it.items {
		if element == nil {
			continue
		}

		if *element == *item {
			return true
		}
	}

	return false
}

// GetHashsetPlusHasAll nil will return false.
func (it *CollectionPtr) GetHashsetPlusHasAll(
	items []string,
) (*Hashset, bool) {
	hashset := it.HashsetAsIs()

	if items == nil || it.IsEmpty() {
		return hashset, false
	}

	return hashset, hashset.HasAllStrings(items)
}

// IsContainsAllPtr nil will return false.
func (it *CollectionPtr) IsContainsAllPtr(
	items *[]string,
) bool {
	if items == nil {
		return false
	}

	if it.IsEmpty() {
		return false
	}

	for _, item := range *items {
		if !it.IsContainsPtr(&item) {
			return false
		}
	}

	return true
}

// IsContainsAll nil will return false.
func (it *CollectionPtr) IsContainsAll(
	items ...string,
) bool {
	if items == nil {
		return false
	}

	return it.IsContainsAllPtr(&items)
}

// IsContainsAllLock
//
// nil will return false.
func (it *CollectionPtr) IsContainsAllLock(
	items ...string,
) bool {
	it.Lock()
	defer it.Unlock()

	if items == nil {
		return false
	}

	return it.IsContainsAllPtr(&items)
}

func (it *CollectionPtr) CharCollectionPtrMap() *CharCollectionMap {
	length := it.Length()
	lengthByFourBestGuess := length / 4
	runeMap := New.CharCollectionMap.CapSelfCap(
		length,
		lengthByFourBestGuess)

	for _, item := range it.items {
		runeMap.AddStringPtr(item)
	}

	return runeMap
}

func (it *CollectionPtr) SummaryString(
	sequence int,
) string {
	header := fmt.Sprintf(
		summaryOfCharCollectionMapLengthFormat,
		it,
		it.Length(),
		sequence)

	return it.SummaryStringWithHeader(header)
}

func (it *CollectionPtr) SummaryStringWithHeader(
	header string,
) string {
	if it.IsEmpty() {
		return header + commonJoiner + NoElements
	}

	return header + it.String()
}

func (it *CollectionPtr) String() string {
	if it.IsEmpty() {
		return commonJoiner + NoElements
	}

	return commonJoiner +
		strings.Join(
			*it.SimpleListPtr(),
			commonJoiner)
}

func (it *CollectionPtr) StringLock() string {
	if it.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	it.Lock()
	defer it.Unlock()

	return commonJoiner +
		strings.Join(
			*it.SimpleListPtr(),
			commonJoiner)
}

func (it *Collection) Join(
	joiner string,
) string {
	if it.IsEmpty() {
		return ""
	}

	return strings.Join(it.items, joiner)
}

func (it *Collection) JoinWith(
	joiner string,
) string {
	if it.IsEmpty() {
		return ""
	}

	return joiner + strings.Join(it.items, joiner)
}

func (it *CollectionPtr) AddCapacity(
	capacities ...int,
) *CollectionPtr {
	if capacities == nil || len(capacities) == 0 {
		return it
	}

	currentCapacity := it.Capacity()

	for _, capacity := range capacities {
		currentCapacity += capacity
	}

	return it.Resize(currentCapacity)
}

// Resize Only resize if capacity is bigger than the current one
// Warning changes current pointer with new one.
func (it *CollectionPtr) Resize(
	newCapacity int,
) *CollectionPtr {
	capacity := it.Capacity()
	if capacity >= newCapacity {
		return it
	}

	newItems := make([]*string, it.Length(), newCapacity)
	copy(newItems, it.items)

	it.items = newItems

	return it
}

func (it *CollectionPtr) Joins(
	separator string,
	items ...string,
) string {
	newItems := make([]string, constants.Zero, it.Length()+len(items))

	for _, item := range it.items {
		if item == nil {
			continue
		}

		newItems = append(newItems, *item)
	}

	if len(items) == 0 {
		return strings.Join(newItems, separator)
	}

	for _, item := range items {
		newItems = append(newItems, item)
	}

	return strings.Join(newItems, separator)
}

func (it *CollectionPtr) JsonModel() *CollectionPtrDataModel {
	return &CollectionPtrDataModel{
		Items: it.items,
	}
}

func (it *CollectionPtr) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *CollectionPtr) MarshalJSON() ([]byte, error) {
	return json.Marshal(*it.JsonModel())
}

func (it *CollectionPtr) UnmarshalJSON(
	data []byte,
) error {
	var dataModel CollectionPtrDataModel

	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.items = dataModel.Items
	}

	return err
}

func (it CollectionPtr) Json() corejson.Result {
	return corejson.New(it)
}

func (it CollectionPtr) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *CollectionPtr) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*CollectionPtr, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *CollectionPtr) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *CollectionPtr {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *CollectionPtr) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *CollectionPtr) AsJsoner() corejson.Jsoner {
	return it
}

func (it *CollectionPtr) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *CollectionPtr) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it *CollectionPtr) JsonString() (jsonString string, err error) {
	return it.Json().JsonString(), nil
}

func (it *CollectionPtr) JsonStringMust() string {
	return it.Json().JsonString()
}

func (it *CollectionPtr) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}
