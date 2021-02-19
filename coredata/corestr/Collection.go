package corestr

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/msgtype"
)

type Collection struct {
	items *[]string
	sync.Mutex
}

func (collection *Collection) Capacity() int {
	if collection.items == nil {
		return 0
	}

	return cap(*collection.items)
}

func (collection *Collection) Length() int {
	if collection.items == nil {
		return 0
	}

	return len(*collection.items)
}

func (collection *Collection) LengthLock() int {
	collection.Lock()
	defer collection.Unlock()

	if collection.items == nil {
		return 0
	}

	return len(*collection.items)
}

func (collection *Collection) IsEqualsPtr(
	anotherCollection *Collection,
) bool {
	return collection.IsEqualsWithSensitivePtr(
		anotherCollection,
		true)
}

func (collection *Collection) IsEqualsWithSensitivePtr(
	anotherCollection *Collection,
	isCaseSensitive bool,
) bool {
	if anotherCollection == nil && collection == nil {
		return true
	}

	if anotherCollection == nil || collection == nil {
		return false
	}

	if collection == anotherCollection {
		return true
	}

	if collection.IsEmpty() && anotherCollection.IsEmpty() {
		return true
	}

	if collection.IsEmpty() || anotherCollection.IsEmpty() {
		return false
	}

	if collection.Length() != anotherCollection.Length() {
		return false
	}

	leftItems := collection.items
	rightItems := anotherCollection.items

	if isCaseSensitive {
		for i, leftVal := range *leftItems {
			if leftVal != (*rightItems)[i] {
				return false
			}
		}

		return true
	}

	for i, leftVal := range *leftItems {
		if !strings.EqualFold(leftVal, (*rightItems)[i]) {
			return false
		}
	}

	return true
}

func (collection *Collection) IsEmptyLock() bool {
	collection.Lock()
	defer collection.Unlock()

	return collection.items == nil ||
		*collection.items == nil ||
		len(*collection.items) == 0
}

func (collection *Collection) IsEmpty() bool {
	return collection.items == nil ||
		*collection.items == nil ||
		len(*collection.items) == 0
}

func (collection *Collection) AddLock(str string) *Collection {
	collection.Lock()
	defer collection.Unlock()

	*collection.items = append(
		*collection.items,
		str)

	return collection
}

func (collection *Collection) Add(str string) *Collection {
	*collection.items = append(
		*collection.items,
		str)

	return collection
}

func (collection *Collection) AddsLock(items ...string) *Collection {
	collection.Lock()
	defer collection.Unlock()

	*collection.items = append(
		*collection.items,
		items...)

	return collection
}

func (collection *Collection) Adds(items ...string) *Collection {
	*collection.items = append(
		*collection.items,
		items...)

	return collection
}

func (collection *Collection) AddCollection(collectionIn *Collection) *Collection {
	return collection.AddStringsPtr(collectionIn.items)
}

// skip on nil
func (collection *Collection) AddCollections(collectionsIn ...*Collection) *Collection {
	for _, collectionIn := range collectionsIn {
		if collectionIn == nil || collectionIn.items == nil {
			continue
		}

		collection.AddStringsPtr(collectionIn.items)
	}

	return collection
}

// skip on nil
func (collection *Collection) AddPointerCollections(collectionsIn *[]*Collection) *Collection {
	for _, collectionIn := range *collectionsIn {
		if collectionIn == nil || collectionIn.items == nil {
			continue
		}

		collection.AddStringsPtr(collectionIn.items)
	}

	return collection
}

func (collection *Collection) AddPointerCollectionsLock(collectionsIn *[]*Collection) *Collection {
	collection.Lock()
	defer collection.Unlock()

	return collection.AddPointerCollections(collectionsIn)
}

func (collection *Collection) AddHashmapsValues(
	hashmaps ...*Hashmap,
) *Collection {
	if hashmaps == nil {
		return collection
	}

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for _, v := range *hashmap.items {
			*collection.items = append(
				*collection.items,
				v)
		}
	}

	return collection
}

func (collection *Collection) AddHashmapsKeys(
	hashmaps ...*Hashmap,
) *Collection {
	if hashmaps == nil {
		return collection
	}

	collection.resizeForHashmaps(
		&hashmaps,
		constants.One)

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k := range *hashmap.items {
			*collection.items = append(
				*collection.items,
				k)
		}
	}

	return collection
}

func (collection *Collection) isResizeRequired(
	length int,
) bool {
	if length < constants.ArbitraryCapacity200 {
		return false
	}

	windowLength := collection.Capacity() - collection.Length()
	if windowLength >= length {
		return false
	}

	return true
}

func (collection *Collection) resizeForHashmaps(
	hashmaps *[]*Hashmap,
	multiplier int,
) *Collection {
	if hashmaps == nil {
		return collection
	}

	length := 0

	for _, hashmap := range *hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		length += hashmap.Length()
	}

	if !collection.isResizeRequired(length) {
		return collection
	}

	finalLength :=
		length*multiplier +
			length/2

	return collection.AddCapacity(finalLength)
}

func (collection *Collection) resizeForCollections(
	collections *[]*Collection,
	multiplier int,
) *Collection {
	if collections == nil {
		return collection
	}

	length := 0

	for _, hashmap := range *collections {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		length += hashmap.Length()
	}

	if !collection.isResizeRequired(length) {
		return collection
	}

	finalLength :=
		length*multiplier +
			length/2

	return collection.AddCapacity(finalLength)
}

func (collection *Collection) resizeForItems(
	items *[]string,
	multiplier int,
) *Collection {
	if items == nil {
		return collection
	}

	length := len(*items)
	if !collection.isResizeRequired(length) {
		return collection
	}

	finalLength :=
		length*multiplier +
			length/2

	return collection.AddCapacity(finalLength)
}

func (collection *Collection) resizeForPointerItems(
	items *[]*string,
	multiplier int,
) *Collection {
	if items == nil {
		return collection
	}

	length := len(*items)
	if !collection.isResizeRequired(length) {
		return collection
	}

	finalLength :=
		length*multiplier +
			length/2

	return collection.AddCapacity(finalLength)
}

func (collection *Collection) resizeForAnys(
	items *[]interface{},
	multiplier int,
) *Collection {
	if items == nil {
		return collection
	}

	length := len(*items)
	if !collection.isResizeRequired(length) {
		return collection
	}

	finalLength :=
		length*multiplier +
			length/2

	return collection.AddCapacity(finalLength)
}

func (collection *Collection) AddHashmapsKeysValues(
	hashmaps ...*Hashmap,
) *Collection {
	if hashmaps == nil {
		return collection
	}

	collection.resizeForHashmaps(
		&hashmaps,
		constants.ArbitraryCapacity2)

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k, v := range *hashmap.items {
			*collection.items = append(
				*collection.items,
				k)
			*collection.items = append(
				*collection.items,
				v)
		}
	}

	return collection
}

func (collection *Collection) AddHashmapsKeysValuesUsingFilter(
	filter IsKeyValueFilter,
	hashmaps ...*Hashmap,
) *Collection {
	if hashmaps == nil {
		return collection
	}

	collection.resizeForHashmaps(
		&hashmaps,
		constants.One)

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k, v := range *hashmap.items {
			result, isAcceptable, isBreak := filter(KeyValuePair{
				Key:   k,
				Value: v,
			})

			if isAcceptable {
				*collection.items = append(
					*collection.items,
					result)
			}

			if isBreak {
				return collection
			}
		}
	}

	return collection
}

func (collection *Collection) AddPtr(str *string) *Collection {
	*collection.items = append(
		*collection.items,
		*str)

	return collection
}

func (collection *Collection) AddPtrLock(str *string) *Collection {
	collection.Lock()
	defer collection.Unlock()

	*collection.items = append(
		*collection.items,
		*str)

	return collection
}

func (collection *Collection) AddWithWgLock(
	str string,
	group *sync.WaitGroup,
) *Collection {
	collection.Lock()
	defer collection.Unlock()

	*collection.items = append(
		*collection.items,
		str)

	group.Done()

	return collection
}

func (collection *Collection) AddsPtrLock(itemsPtr ...*string) *Collection {
	collection.Lock()
	defer collection.Unlock()

	for _, str := range itemsPtr {
		*collection.items = append(
			*collection.items,
			*str)
	}

	return collection
}

func (collection *Collection) AddStringsPtrWgLock(
	str *[]string,
	group *sync.WaitGroup,
) *Collection {
	collection.Lock()
	defer collection.Unlock()

	*collection.items = append(
		*collection.items,
		*str...)

	group.Done()

	return collection
}

// skip on nil
func (collection *Collection) AddPointerStringsPtrLock(
	pointerStringItems *[]*string,
) *Collection {
	collection.Lock()
	defer collection.Unlock()

	return collection.
		AddPointerStringsPtr(pointerStringItems)
}

// skip on nil
func (collection *Collection) AddPointerStringsPtr(
	pointerStringItems *[]*string,
) *Collection {
	for i := range *pointerStringItems {
		newPtr := (*pointerStringItems)[i]

		if newPtr == nil {
			continue
		}

		*collection.items = append(
			*collection.items,
			*(*pointerStringItems)[i])
	}

	return collection
}

func (collection *Collection) IndexAt(index int) string {
	return (*collection.items)[index]
}

func (collection *Collection) SafePointerIndexAt(index int) *string {
	length := collection.Length()
	if length-1 < index {
		return nil
	}

	return &(*collection.items)[index]
}

func (collection *Collection) SafePointerIndexAtUsingLength(length, index int) *string {
	if length-1 < index {
		return nil
	}

	return &(*collection.items)[index]
}

func (collection *Collection) SafeIndexAtUsingLength(defaultString string, length, index int) string {
	if length-1 < index {
		return defaultString
	}

	return (*collection.items)[index]
}

func (collection *Collection) First() string {
	return (*collection.items)[0]
}

func (collection *Collection) Single() string {
	length := collection.Length()
	if length != 1 {
		msgtype.LengthShouldBeEqualToMessage.HandleUsingPanic("1", length)
	}

	return (*collection.items)[0]
}

func (collection *Collection) Last() string {
	length := collection.Length()

	return (*collection.items)[length-1]
}

func (collection *Collection) LastOrDefault() string {
	length := collection.Length()

	if length == 0 {
		return constants.EmptyString
	}

	return (*collection.items)[length-1]
}

func (collection *Collection) FirstOrDefault() string {
	if collection.IsEmpty() {
		return constants.EmptyString
	}

	return (*collection.items)[0]
}

func (collection *Collection) AddStringsPtr(stringItems *[]string) *Collection {
	if stringItems == nil {
		return collection
	}

	collection.resizeForItems(
		stringItems,
		constants.One)

	*collection.items = append(
		*collection.items,
		*stringItems...)

	return collection
}

func (collection *Collection) AddStringsPtrAsync(
	wg *sync.WaitGroup,
	stringItems *[]string,
) *Collection {
	if stringItems == nil {
		return collection
	}

	go func() {
		collection.Lock()

		collection.AddStringsPtr(stringItems)

		collection.Unlock()

		wg.Done()
	}()

	return collection
}

func (collection *Collection) InsertItemsAt(index int, stringItems *[]string) *Collection {
	length := collection.Length()
	isAtFirst := length == 0
	isAtLast := length-1 == index
	isAppendItems := isAtFirst || isAtLast

	if isAppendItems {
		return collection.AddStringsPtr(stringItems)
	}

	// https://bit.ly/3pIDfRY
	*collection.items =
		append(
			(*collection.items)[:index],
			*stringItems...)

	*collection.items = append(
		*collection.items,
		(*collection.items)[index:]...)

	return collection
}

func (collection *Collection) RemoveAt(index int) *Collection {
	*collection.items = append(
		(*collection.items)[:index],
		(*collection.items)[index+1:]...)

	return collection
}

// creates a new collection without the indexes mentioned.
//
// it is better to filter out than remove.
func (collection *Collection) RemoveItemsIndexes(
	isIgnoreRemoveError bool,
	indexes ...int,
) *Collection {
	if isIgnoreRemoveError && indexes == nil {
		return collection
	}

	return collection.
		RemoveItemsIndexesPtr(isIgnoreRemoveError, &indexes)
}

// creates a new collection without the indexes mentioned.
//
// it is better to filter out than remove.
func (collection *Collection) RemoveItemsIndexesPtr(
	isIgnoreRemoveError bool,
	indexes *[]int,
) *Collection {
	if indexes == nil {
		return collection
	}

	length := collection.Length()
	indexesLength := len(*indexes)
	hasPossibleError := length == 0 && indexesLength > 0

	if hasPossibleError && !isIgnoreRemoveError {
		panic(msgtype.CannotRemoveIndexesFromEmptyCollection)
	}

	if !isIgnoreRemoveError {
		msgtype.PanicOnIndexOutOfRange(length, indexes)
	}

	if hasPossibleError {
		return collection
	}

	newList := make([]string, 0, collection.Capacity())
	for i, s := range *collection.items { //nolint:wsl
		if coreindexes.HasIndex(indexes, i) {
			continue
		}

		newList = append(newList, s)
	}

	collection.items = &newList

	return collection
}

func (collection *Collection) AppendCollectionPtr(
	anotherCollection *Collection,
) *Collection {
	collection.resizeForItems(
		anotherCollection.items,
		constants.One)

	*collection.items = append(
		*collection.items,
		*anotherCollection.items...)

	return collection
}

func (collection *Collection) AppendCollectionsPtr(
	anotherCollectionsPtr ...*Collection,
) *Collection {
	if anotherCollectionsPtr == nil {
		return collection
	}

	return collection.AppendPointersCollectionsPtr(
		&anotherCollectionsPtr)
}

func (collection *Collection) AppendPointersCollectionsPtr(
	anotherCollectionsPtr *[]*Collection,
) *Collection {
	if anotherCollectionsPtr == nil {
		return collection
	}

	collection.resizeForCollections(
		anotherCollectionsPtr,
		constants.One)

	capacitiesIncrease := 0
	for _, currentCollection := range *anotherCollectionsPtr {

		if currentCollection == nil || currentCollection.IsEmpty() {
			continue
		}

		capacitiesIncrease += currentCollection.Length()
	}

	collection.AddCapacity(capacitiesIncrease)

	for _, currentCollection := range *anotherCollectionsPtr {
		if currentCollection == nil || currentCollection.IsEmpty() {
			continue
		}

		*collection.items = append(
			*collection.items,
			*currentCollection.items...)
	}

	return collection
}

func (collection *Collection) AppendCollectionsPtrAsync(
	wg *sync.WaitGroup,
	anotherCollectionsPtr ...*Collection,
) *Collection {
	if anotherCollectionsPtr == nil {
		return collection
	}

	go func() {
		collection.AppendPointersCollectionsPtr(
			&anotherCollectionsPtr)

		wg.Done()
	}()

	return collection
}

// Continue on nil
func (collection *Collection) AppendAnysAsync(
	wg *sync.WaitGroup,
	anys ...interface{},
) *Collection {
	if anys == nil {
		return collection
	}

	go func() {
		collection.Lock()
		collection.resizeForAnys(
			&anys,
			constants.One)
		collection.Unlock()

		collection.AppendAnysLock(&anys)

		wg.Done()
	}()

	return collection
}

// Continue on nil
func (collection *Collection) AppendAnysLock(
	anys *[]interface{},
) *Collection {
	if anys == nil {
		return collection
	}

	collection.resizeForAnys(
		anys,
		constants.One)

	for _, any := range *anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)

		collection.Lock()
		*collection.items = append(
			*collection.items,
			anyStr)
		collection.Unlock()
	}

	return collection
}

// Continue on nil
func (collection *Collection) AppendAnys(
	anys ...interface{},
) *Collection {
	if anys == nil {
		return collection
	}

	collection.resizeForAnys(
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

		*collection.items = append(
			*collection.items,
			anyStr)
	}

	return collection
}

// Skip on nil
func (collection *Collection) AppendAnysUsingFilter(
	filter IsStringFilter,
	anys ...interface{},
) *Collection {
	if anys == nil {
		return collection
	}

	collection.resizeForAnys(
		&anys,
		constants.One)

	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(
			constants.SprintValueFormat,
			any)

		result, isKeep, isBreak := filter(anyStr)

		if !isKeep {
			continue
		}

		*collection.items = append(
			*collection.items,
			result)

		if isBreak {
			return collection
		}
	}

	return collection
}

// Skip on nil
func (collection *Collection) AppendAnysUsingFilterLock(
	filter IsStringFilter,
	anys ...interface{},
) *Collection {
	if anys == nil {
		return collection
	}

	collection.resizeForAnys(
		&anys,
		constants.One)

	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)
		result, isKeep, isBreak := filter(anyStr)

		if !isKeep {
			continue
		}

		collection.Lock()
		*collection.items = append(
			*collection.items,
			result)
		collection.Unlock()

		if isBreak {
			return collection
		}
	}

	return collection
}

// Continue on nil
func (collection *Collection) AppendNonEmptyAnys(
	anys ...interface{},
) *Collection {
	if anys == nil {
		return collection
	}

	collection.resizeForAnys(
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

		*collection.items = append(
			*collection.items,
			anyStr)
	}

	return collection
}

// Skip on nil
func (collection *Collection) AddsPtr(itemsPtr ...*string) *Collection {
	if itemsPtr == nil {
		return collection
	}

	for _, str := range itemsPtr {
		if str == nil {
			continue
		}

		*collection.items = append(
			*collection.items,
			*str)
	}

	return collection
}

// Skip on nil
func (collection *Collection) AddsPtrAsync(
	wg *sync.WaitGroup,
	itemsPtr ...*string,
) *Collection {
	if itemsPtr == nil {
		return collection
	}

	go func() {
		collection.Lock()
		collection.resizeForPointerItems(
			&itemsPtr,
			constants.One)

		collection.Unlock()

		for _, str := range itemsPtr {
			if str == nil {
				continue
			}

			collection.Lock()

			*collection.items = append(
				*collection.items,
				*str)

			collection.Unlock()
		}

		wg.Done()
	}()

	return collection
}

func (collection *Collection) AddsNonEmptyPtr(itemsPtr ...*string) *Collection {
	if itemsPtr == nil {
		return collection
	}

	for _, str := range itemsPtr {
		if str == nil || *str == "" {
			continue
		}

		*collection.items = append(
			*collection.items,
			*str)
	}

	return collection
}

func (collection *Collection) AddsNonEmptyPtrLock(itemsPtr ...*string) *Collection {
	if itemsPtr == nil {
		return collection
	}

	for _, str := range itemsPtr {
		if str == nil || *str == "" {
			continue
		}

		collection.Lock()
		*collection.items = append(
			*collection.items,
			*str)
		collection.Unlock()
	}

	return collection
}

func (collection *Collection) UniqueBoolMapLock() *map[string]bool {
	collection.Lock()
	defer collection.Unlock()

	return collection.UniqueBoolMap()
}

func (collection *Collection) UniqueBoolMap() *map[string]bool {
	respectiveMap := make(
		map[string]bool,
		collection.Length())

	for _, item := range *collection.items {
		respectiveMap[item] = true
	}

	return &respectiveMap
}

func (collection *Collection) UniqueListPtr() *[]string {
	boolMap := collection.UniqueBoolMap()
	list := make([]string, len(*boolMap))

	i := 0
	for str := range *boolMap {
		list[i] = str
		i++
	}

	return &list
}

func (collection *Collection) UniqueListPtrLock() *[]string {
	collection.Lock()
	defer collection.Unlock()

	return collection.UniqueListPtr()
}

func (collection *Collection) UniqueListLock() []string {
	collection.Lock()
	defer collection.Unlock()

	return collection.UniqueList()
}

func (collection *Collection) UniqueList() []string {
	return *collection.UniqueListPtr()
}

func (collection *Collection) List() []string {
	return *collection.items
}

// must return a slice
func (collection *Collection) Filter(filter IsStringFilter) *[]string {
	if collection.IsEmpty() {
		return &([]string{})
	}

	list := make([]string, 0, collection.Length())

	for _, element := range *collection.items {
		result, isKeep, isBreak := filter(element)

		if isKeep {
			list = append(list, result)
		}

		if isBreak {
			return &list
		}
	}

	return &list
}

// must return a slice
func (collection *Collection) FilterLock(filter IsStringFilter) *[]string {
	elements := collection.ListCopyPtrLock()
	length := len(*elements)

	if length == 0 {
		return elements
	}

	list := make([]string, 0, length)

	for _, element := range *elements {
		result, isKeep, isBreak := filter(element)

		if isKeep {
			list = append(list, result)
		}

		if isBreak {
			return &list
		}
	}

	return &list
}

// must return a items
func (collection *Collection) FilteredCollection(filter IsStringFilter) *Collection {
	return NewCollectionUsingStrings(collection.Filter(filter), false)
}

// must return a items
func (collection *Collection) FilteredCollectionLock(filter IsStringFilter) *Collection {
	return NewCollectionUsingStrings(collection.FilterLock(filter), false)
}

// must return a slice
func (collection *Collection) FilterPtrLock(filterPtr IsStringPointerFilter) *[]*string {
	elements := collection.ListCopyPtrLock()
	length := len(*elements)

	if length == 0 {
		return &([]*string{})
	}

	list := make([]*string, 0, length)

	for i := range *elements {
		result, isKeep, isBreak := filterPtr(&(*elements)[i])

		if isKeep {
			list = append(list, result)
		}

		if isBreak {
			return &list
		}
	}

	return &list
}

// must return a slice
func (collection *Collection) FilterPtr(filterPtr IsStringPointerFilter) *[]*string {
	if collection.IsEmpty() {
		return &([]*string{})
	}

	list := make([]*string, 0, collection.Length())

	for _, element := range *collection.items {
		result, isKeep, isBreak := filterPtr(&element)

		if isKeep {
			list = append(list, result)
		}

		if isBreak {
			return &list
		}
	}

	return &list
}

// must return a slice
func (collection *Collection) NonEmptyListPtr() *[]string {
	if collection.IsEmpty() {
		return &([]string{})
	}

	list := make([]string, 0, collection.Length())

	for _, element := range *collection.items {
		if element == "" {
			continue
		}

		list = append(list, element)
	}

	return &list
}

func (collection *Collection) Hashset() *Hashset {
	return NewHashsetUsingStrings(
		collection.items,
		collection.Length()*2,
		true)
}

func (collection *Collection) HashsetLock() *Hashset {
	return NewHashsetUsingStrings(
		collection.ListCopyPtrLock(),
		0,
		false)
}

// direct return pointer
func (collection *Collection) Items() *[]string {
	return collection.items
}

// direct return pointer
func (collection *Collection) ListPtr() *[]string {
	return collection.items
}

// returns a copy of the items
//
// must return a slice
func (collection *Collection) ListCopyPtrLock() *[]string {
	collection.Lock()
	defer collection.Unlock()

	if collection.items == nil ||
		*collection.items == nil {
		return &([]string{})
	}

	return &(*collection.items)
}

func (collection *Collection) HasLock(str string) bool {
	collection.Lock()
	defer collection.Unlock()

	return collection.Has(str)
}

func (collection *Collection) Has(str string) bool {
	if collection.IsEmpty() {
		return false
	}

	for _, element := range *collection.items {
		if element == str {
			return true
		}
	}

	return false
}

func (collection *Collection) HasAll(items ...string) bool {
	if collection.IsEmpty() {
		return false
	}

	for _, element := range items {
		if !collection.IsContainsPtr(&element) {
			return false
		}
	}

	return true
}

// Creates new doesn't modify current collection
func (collection *Collection) SortedListAsc() *[]string {
	if collection.IsEmpty() {
		return &[]string{}
	}

	list := &(*collection.items)
	sort.Strings(*list)

	return list
}

// mutates current collection
func (collection *Collection) SortedAsc() *Collection {
	if collection.IsEmpty() {
		return collection
	}

	sort.Strings(*collection.items)

	return collection
}

// mutates current collection
func (collection *Collection) SortedAscLock() *Collection {
	if collection.IsEmptyLock() {
		return collection
	}

	collection.Lock()
	defer collection.Unlock()

	sort.Strings(*collection.items)

	return collection
}

// Creates new one.
func (collection *Collection) SortedListDsc() *[]string {
	list := collection.SortedListAsc()
	length := len(*list)
	mid := length / 2

	for i := 0; i < mid; i++ {
		temp := (*list)[i]
		(*list)[i] = (*list)[length-1-i]
		(*list)[length-1-i] = temp
	}

	return list
}

// mutates itself.
func (collection *Collection) SortedDsc() *Collection {
	list := collection.items
	length := len(*list)
	mid := length / 2

	for i := 0; i < mid; i++ {
		temp := (*list)[i]
		(*list)[i] = (*list)[length-1-i]
		(*list)[length-1-i] = temp
	}

	return collection
}

func (collection *Collection) HasUsingSensitivity(str string, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return collection.Has(str)
	}

	for _, element := range *collection.items {
		if strings.EqualFold(element, str) {
			return true
		}
	}

	return false
}

func (collection *Collection) IsContainsPtr(item *string) bool {
	if item == nil || collection.IsEmpty() {
		return false
	}

	for _, element := range *collection.items {
		if element == *item {
			return true
		}
	}

	return false
}

// nil will return false.
func (collection *Collection) GetHashsetPlusHasAll(items *[]string) (*Hashset, bool) {
	hashset := collection.Hashset()

	if items == nil || collection.IsEmpty() {
		return hashset, false
	}

	return hashset, hashset.HasAllStringsPtr(items)
}

// nil will return false.
func (collection *Collection) IsContainsAllPtr(items *[]string) bool {
	if items == nil {
		return false
	}

	if collection.IsEmpty() {
		return false
	}

	for _, item := range *items {
		if !collection.IsContainsPtr(&item) {
			return false
		}
	}

	return true
}

// nil will return false.
func (collection *Collection) IsContainsAll(items ...string) bool {
	if items == nil {
		return false
	}

	return collection.IsContainsAllPtr(&items)
}

// nil will return false.
func (collection *Collection) IsContainsAllLock(items ...string) bool {
	collection.Lock()
	defer collection.Unlock()

	if items == nil {
		return false
	}

	return collection.IsContainsAllPtr(&items)
}

func (collection *Collection) CharCollectionMap() *CharCollectionMap {
	length := collection.Length()
	lengthByFourBestGuess := length / 4
	runeMap := NewCharCollectionMap(
		length,
		lengthByFourBestGuess)

	for _, item := range *collection.items {
		runeMap.AddStringPtr(&item)
	}

	return runeMap
}

func (collection *Collection) SummaryString(sequence int) string {
	header := fmt.Sprintf(
		summaryOfCharCollectionMapLengthFormat,
		collection,
		collection.Length(),
		sequence)

	return collection.SummaryStringWithHeader(header)
}

func (collection *Collection) SummaryStringWithHeader(header string) string {
	if collection.IsEmpty() {
		return header + commonJoiner + NoElements
	}

	return header + collection.String()
}

func (collection *Collection) String() string {
	if collection.IsEmpty() {
		return commonJoiner + NoElements
	}

	return commonJoiner +
		strings.Join(
			*collection.items,
			commonJoiner)
}

func (collection *Collection) StringLock() string {
	if collection.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	collection.Lock()
	defer collection.Unlock()

	return commonJoiner +
		strings.Join(
			*collection.items,
			commonJoiner)
}

func (collection *Collection) Join(
	separator string,
) string {
	return strings.Join(*collection.items, separator)
}

func (collection *Collection) AddCapacity(
	capacities ...int,
) *Collection {
	if capacities == nil || len(capacities) == 0 {
		return collection
	}

	currentCapacity := collection.Capacity()

	for _, capacity := range capacities {
		currentCapacity += capacity
	}

	return collection.Resize(currentCapacity)
}

// Only resize if capacity is bigger than the current one
func (collection *Collection) Resize(
	newCapacity int,
) *Collection {
	capacity := collection.Capacity()
	if capacity >= newCapacity {
		return collection
	}

	newItems := make([]string, collection.Length(), newCapacity)
	copy(newItems, *collection.items)

	collection.items = &newItems

	return collection
}

func (collection *Collection) Joins(
	separator string,
	items ...string,
) string {
	if len(items) == 0 {
		return strings.Join(*collection.items, separator)
	}

	newItems := make([]string, 0, collection.Length()+len(items))
	copy(newItems, *collection.items)

	newItems = append(newItems, items...)

	return strings.Join(newItems, separator)
}

// clears existing items.
func (collection *Collection) Clear() *Collection {
	if collection.IsEmpty() {
		return collection
	}

	clearedItems := (*collection.items)[:0]
	collection.items = &clearedItems

	return collection
}

func (collection *Collection) JsonModel() *CollectionDataModel {
	return &CollectionDataModel{
		Items: collection.items,
	}
}

func (collection *Collection) JsonModelAny() interface{} {
	return collection.JsonModel()
}

func (collection *Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(*collection.JsonModel())
}

func (collection *Collection) UnmarshalJSON(data []byte) error {
	var dataModel CollectionDataModel

	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		collection.items = dataModel.Items
	}

	return err
}

//goland:noinspection GoLinterLocal
func (collection *Collection) Json() *corejson.Result {
	if collection.IsEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(collection)

	return corejson.NewPtr(jsonBytes, err)
}

//goland:noinspection GoLinterLocal
func (collection *Collection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Collection, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return EmptyCollection(), nil
	}

	err := json.Unmarshal(*jsonResult.Bytes, &collection)

	if err != nil {
		return EmptyCollection(), err
	}

	return collection, nil
}

// Panic if error
//goland:noinspection GoLinterLocal
func (collection *Collection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Collection {
	newUsingJson, err :=
		collection.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

// Panic if error
func (collection *Collection) JsonParseSelfInject(jsonResult *corejson.Result) {
	collection.ParseInjectUsingJsonMust(jsonResult)
}

func (collection *Collection) AsJsoner() *corejson.Jsoner {
	var jsoner corejson.Jsoner = collection

	return &jsoner
}

func (collection *Collection) AsJsonParseSelfInjector() *corejson.ParseSelfInjector {
	var jsonInjector corejson.ParseSelfInjector = collection

	return &jsonInjector
}

func (collection *Collection) AsJsonMarshaller() *corejson.JsonMarshaller {
	var jsonMarshaller corejson.JsonMarshaller = collection

	return &jsonMarshaller
}
