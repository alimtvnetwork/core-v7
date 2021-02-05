package corestr

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/coresort/strsort"
	"gitlab.com/evatix-go/core/msgtype"
)

type CollectionPtr struct {
	items *[]*string
	sync.Mutex
}

func (collectionPtr *CollectionPtr) Capacity() int {
	if collectionPtr.items == nil {
		return 0
	}

	return cap(*collectionPtr.items)
}

func (collectionPtr *CollectionPtr) Length() int {
	if collectionPtr.items == nil {
		return 0
	}

	return len(*collectionPtr.items)
}

func (collectionPtr *CollectionPtr) LengthLock() int {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	if collectionPtr.items == nil {
		return 0
	}

	return len(*collectionPtr.items)
}

//goland:noinspection GoVetCopyLock
func (collectionPtr *CollectionPtr) IsEquals(
	anotherCollectionPtr CollectionPtr,
) bool {
	return collectionPtr.IsEqualsWithSensitivePtr(
		&anotherCollectionPtr,
		true)
}

func (collectionPtr *CollectionPtr) IsEqualsPtr(
	anotherCollectionPtr *CollectionPtr,
) bool {
	return collectionPtr.IsEqualsWithSensitivePtr(
		anotherCollectionPtr,
		true)
}

func (collectionPtr *CollectionPtr) IsEqualsWithSensitivePtr(
	anotherCollectionPtr *CollectionPtr,
	isCaseSensitive bool,
) bool {
	if anotherCollectionPtr == nil && collectionPtr == nil {
		return true
	}

	if anotherCollectionPtr == nil || collectionPtr == nil {
		return false
	}

	if collectionPtr == anotherCollectionPtr {
		return true
	}

	if collectionPtr.IsEmpty() && anotherCollectionPtr.IsEmpty() {
		return true
	}

	if collectionPtr.IsEmpty() || anotherCollectionPtr.IsEmpty() {
		return false
	}

	if collectionPtr.Length() != anotherCollectionPtr.Length() {
		return false
	}

	leftItems := collectionPtr.items
	rightItems := anotherCollectionPtr.items

	if isCaseSensitive {
		for i, leftVal := range *leftItems {
			if leftVal != (*rightItems)[i] {
				return false
			}
		}

		return true
	}

	for i, leftVal := range *leftItems {
		if !strings.EqualFold(*leftVal, *(*rightItems)[i]) {
			return false
		}
	}

	return true
}

func (collectionPtr *CollectionPtr) IsEmptyLock() bool {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	return collectionPtr.items == nil ||
		*collectionPtr.items == nil ||
		len(*collectionPtr.items) == 0
}

func (collectionPtr *CollectionPtr) IsEmpty() bool {
	return collectionPtr.items == nil ||
		*collectionPtr.items == nil ||
		len(*collectionPtr.items) == 0
}

func (collectionPtr *CollectionPtr) AddLock(str string) *CollectionPtr {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	*collectionPtr.items = append(
		*collectionPtr.items,
		&str)

	return collectionPtr
}

func (collectionPtr *CollectionPtr) Add(str string) *CollectionPtr {
	*collectionPtr.items = append(
		*collectionPtr.items,
		&str)

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddsLock(items ...string) *CollectionPtr {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	for i := range items {
		*collectionPtr.items = append(
			*collectionPtr.items,
			&(items[i]))
	}

	return collectionPtr
}

func (collectionPtr *CollectionPtr) Adds(items ...string) *CollectionPtr {
	for i := range items {
		*collectionPtr.items = append(
			*collectionPtr.items,
			&(items[i]))
	}

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddHashmapsValues(
	hashmaps ...*Hashmap,
) *CollectionPtr {
	if hashmaps == nil {
		return collectionPtr
	}

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for _, v := range *hashmap.items {
			newV := v
			*collectionPtr.items = append(
				*collectionPtr.items,
				&newV)
		}
	}

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddHashmapsKeys(
	hashmaps ...*Hashmap,
) *CollectionPtr {
	if hashmaps == nil {
		return collectionPtr
	}

	collectionPtr.resizeForHashmaps(
		&hashmaps,
		constants.One)

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k := range *hashmap.items {
			keyCopied := k
			*collectionPtr.items = append(
				*collectionPtr.items,
				&keyCopied)
		}
	}

	return collectionPtr
}

func (collectionPtr *CollectionPtr) resizeForHashmaps(
	hashmaps *[]*Hashmap,
	multiplier int,
) *CollectionPtr {
	if hashmaps == nil {
		return collectionPtr
	}

	length := 0

	for _, hashmap := range *hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		length += hashmap.Length()
	}

	if length < constants.ArbitraryCapacity100 {
		return collectionPtr
	}

	finalLength :=
		length*multiplier +
			length/2

	return collectionPtr.AddCapacity(finalLength)
}

func (collectionPtr *CollectionPtr) resizeForCollectionPtrs(
	collections *[]*CollectionPtr,
	multiplier int,
) *CollectionPtr {
	if collections == nil {
		return collectionPtr
	}

	length := 0

	for _, collection := range *collections {
		if collection == nil || collection.IsEmpty() {
			continue
		}

		length += collection.Length()
	}

	if length < constants.ArbitraryCapacity100 {
		return collectionPtr
	}

	finalLength :=
		length*multiplier +
			length/2

	return collectionPtr.AddCapacity(finalLength)
}

func (collectionPtr *CollectionPtr) resizeForItems(
	items *[]*string,
	multiplier int,
) *CollectionPtr {
	if items == nil {
		return collectionPtr
	}

	length := len(*items)
	if length < constants.ArbitraryCapacity100 {
		return collectionPtr
	}

	finalLength :=
		length*multiplier +
			length/2

	return collectionPtr.AddCapacity(finalLength)
}

func (collectionPtr *CollectionPtr) resizeForAnys(
	items *[]interface{},
	multiplier int,
) *CollectionPtr {
	if items == nil {
		return collectionPtr
	}

	length := len(*items)
	if length < constants.ArbitraryCapacity100 {
		return collectionPtr
	}

	finalLength :=
		length*multiplier +
			length/2

	return collectionPtr.AddCapacity(finalLength)
}

func (collectionPtr *CollectionPtr) AddHashmapsKeysValues(
	hashmaps ...*Hashmap,
) *CollectionPtr {
	if hashmaps == nil {
		return collectionPtr
	}

	collectionPtr.resizeForHashmaps(
		&hashmaps,
		constants.ArbitraryCapacity2)

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k, v := range *hashmap.items {
			kc := k
			vc := v
			*collectionPtr.items = append(
				*collectionPtr.items,
				&kc)
			*collectionPtr.items = append(
				*collectionPtr.items,
				&vc)
		}
	}

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddHashmapsKeysValuesUsingFilter(
	filter IsKeyValueFilter,
	hashmaps ...*Hashmap,
) *CollectionPtr {
	if hashmaps == nil {
		return collectionPtr
	}

	collectionPtr.resizeForHashmaps(
		&hashmaps,
		constants.One)

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k, v := range *hashmap.items {
			result, isAcceptable := filter(KeyValuePair{
				Key:   k,
				Value: v,
			})

			if isAcceptable {
				*collectionPtr.items = append(
					*collectionPtr.items,
					&result)
			}
		}
	}

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddPtr(str *string) *CollectionPtr {
	*collectionPtr.items = append(
		*collectionPtr.items,
		str)

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddPtrLock(str *string) *CollectionPtr {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	*collectionPtr.items = append(
		*collectionPtr.items,
		str)

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddWithWgLock(
	str string,
	group *sync.WaitGroup,
) *CollectionPtr {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	*collectionPtr.items = append(
		*collectionPtr.items,
		&str)

	group.Done()

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddsPtrLock(itemsPtr ...*string) *CollectionPtr {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	for _, str := range itemsPtr {
		*collectionPtr.items = append(
			*collectionPtr.items,
			str)
	}

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddStringsPtrWgLock(
	stringItems *[]string,
	group *sync.WaitGroup,
) *CollectionPtr {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	collectionPtr.AddStringsPtr(stringItems)
	group.Done()

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddStringsPtr(stringItems *[]string) *CollectionPtr {
	for i := range *stringItems {
		*collectionPtr.items = append(
			*collectionPtr.items,
			&(*stringItems)[i])
	}

	return collectionPtr
}

func (collectionPtr *CollectionPtr) InsertItemsAt(index int, stringItems *[]string) *CollectionPtr {
	length := collectionPtr.Length()
	isAtFirst := length == 0
	isAtLast := length-1 == index
	isAppendItems := isAtFirst || isAtLast

	if isAppendItems {
		return collectionPtr.AddStringsPtr(stringItems)
	}

	pointerStrings := converters.StringsToPointerStrings(stringItems)

	// https://bit.ly/3pIDfRY
	*collectionPtr.items =
		append(
			(*collectionPtr.items)[:index],
			*pointerStrings...)

	*collectionPtr.items = append(
		*collectionPtr.items,
		(*collectionPtr.items)[index:]...)

	return collectionPtr
}

func (collectionPtr *CollectionPtr) RemoveAt(index int) *CollectionPtr {
	*collectionPtr.items = append(
		(*collectionPtr.items)[:index],
		(*collectionPtr.items)[index+1:]...)

	return collectionPtr
}

// creates a new collection without the indexes mentioned.
//
// it is better to filter out than remove.
func (collectionPtr *CollectionPtr) RemoveItemsIndexes(
	isIgnoreRemoveError bool,
	indexes ...int,
) *CollectionPtr {
	if isIgnoreRemoveError && indexes == nil {
		return collectionPtr
	}

	return collectionPtr.
		RemoveItemsIndexesPtr(isIgnoreRemoveError, &indexes)
}

// creates a new collection without the indexes mentioned.
//
// it is better to filter out than remove.
func (collectionPtr *CollectionPtr) RemoveItemsIndexesPtr(
	isIgnoreRemoveError bool,
	indexes *[]int,
) *CollectionPtr {
	length := collectionPtr.Length()
	indexesLength := len(*indexes)
	hasPossibleError := length == 0 && indexesLength > 0

	if hasPossibleError && !isIgnoreRemoveError {
		panic(msgtype.CannotRemoveIndexesFromEmptyCollection)
	}

	if !isIgnoreRemoveError {
		msgtype.PanicOnIndexOutOfRange(length, indexes)
	}

	if hasPossibleError {
		return collectionPtr
	}

	newList := make([]*string, 0, collectionPtr.Capacity())
	for i, s := range *collectionPtr.items {
		if coreindexes.HasIndex(indexes, i) {
			continue
		}

		newList = append(newList, s)
	}

	collectionPtr.items = &newList

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddPointerStringsPtrLock(pointerStringItems *[]*string) *CollectionPtr {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	return collectionPtr.
		AddPointerStringsPtr(pointerStringItems)
}

func (collectionPtr *CollectionPtr) AddPointerStringsPtr(pointerStringItems *[]*string) *CollectionPtr {
	for i := range *pointerStringItems {
		*collectionPtr.items = append(
			*collectionPtr.items,
			(*pointerStringItems)[i])
	}

	return collectionPtr
}

//goland:noinspection GoVetCopyLock
func (collectionPtr *CollectionPtr) AppendCollectionPtr(
	anotherCollectionPtr CollectionPtr,
) *CollectionPtr {
	collectionPtr.resizeForItems(
		anotherCollectionPtr.items,
		constants.One)

	*collectionPtr.items = append(
		*collectionPtr.items,
		*anotherCollectionPtr.items...)

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AppendCollectionPtrPtr(
	anotherCollectionPtr *CollectionPtr,
) *CollectionPtr {
	collectionPtr.resizeForItems(
		anotherCollectionPtr.items,
		constants.One)

	*collectionPtr.items = append(
		*collectionPtr.items,
		*anotherCollectionPtr.items...)

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AppendCollectionPtrsPtr(
	anotherCollectionPtrsPtr ...*CollectionPtr,
) *CollectionPtr {
	if anotherCollectionPtrsPtr == nil {
		return collectionPtr
	}

	collectionPtr.resizeForCollectionPtrs(
		&anotherCollectionPtrsPtr,
		constants.One)

	capacitiesIncrease := 0
	for _, currentCollectionPtr := range anotherCollectionPtrsPtr {
		if currentCollectionPtr == nil || currentCollectionPtr.IsEmpty() {
			continue
		}

		capacitiesIncrease += currentCollectionPtr.Length()
	}

	collectionPtr.AddCapacity(capacitiesIncrease)

	for _, currentCollectionPtr := range anotherCollectionPtrsPtr {
		if currentCollectionPtr == nil || currentCollectionPtr.IsEmpty() {
			continue
		}

		*collectionPtr.items = append(
			*collectionPtr.items,
			*currentCollectionPtr.items...)
	}

	return collectionPtr
}

// Continue on nil
func (collectionPtr *CollectionPtr) AppendAnysLock(anys ...interface{}) *CollectionPtr {
	if anys == nil {
		return collectionPtr
	}

	collectionPtr.resizeForAnys(
		&anys,
		constants.One)

	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)

		collectionPtr.Lock()
		*collectionPtr.items = append(
			*collectionPtr.items,
			&anyStr)
		collectionPtr.Unlock()
	}

	return collectionPtr
}

// Continue on nil
func (collectionPtr *CollectionPtr) AppendAnys(anys ...interface{}) *CollectionPtr {
	if anys == nil {
		return collectionPtr
	}

	collectionPtr.resizeForAnys(
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

		*collectionPtr.items = append(
			*collectionPtr.items,
			&anyStr)
	}

	return collectionPtr
}

// Skip on nil
func (collectionPtr *CollectionPtr) AppendAnysUsingFilter(
	filter IsStringFilter,
	anys ...interface{},
) *CollectionPtr {
	if anys == nil {
		return collectionPtr
	}

	collectionPtr.resizeForAnys(
		&anys,
		constants.One)

	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(
			constants.SprintValueFormat,
			any)

		result, isKeep := filter(anyStr)

		if !isKeep {
			continue
		}

		*collectionPtr.items = append(
			*collectionPtr.items,
			&result)
	}

	return collectionPtr
}

// Skip on nil
func (collectionPtr *CollectionPtr) AppendAnysUsingFilterLock(
	filter IsStringFilter,
	anys ...interface{},
) *CollectionPtr {
	if anys == nil {
		return collectionPtr
	}

	collectionPtr.resizeForAnys(
		&anys,
		constants.One)

	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)
		result, isKeep := filter(anyStr)

		if !isKeep {
			continue
		}

		collectionPtr.Lock()
		*collectionPtr.items = append(
			*collectionPtr.items,
			&result)
		collectionPtr.Unlock()
	}

	return collectionPtr
}

// Continue on nil
func (collectionPtr *CollectionPtr) AppendNonEmptyAnys(
	anys ...interface{},
) *CollectionPtr {
	if anys == nil {
		return collectionPtr
	}

	collectionPtr.resizeForAnys(
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

		*collectionPtr.items = append(
			*collectionPtr.items,
			&anyStr)
	}

	return collectionPtr
}

// adds nil
func (collectionPtr *CollectionPtr) AddsPtr(
	itemsPtr ...*string,
) *CollectionPtr {
	if itemsPtr == nil {
		return collectionPtr
	}

	for _, str := range itemsPtr {
		*collectionPtr.items = append(
			*collectionPtr.items,
			str)
	}

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddsNonEmptyPtr(
	itemsPtr ...*string,
) *CollectionPtr {
	if itemsPtr == nil {
		return collectionPtr
	}

	for _, str := range itemsPtr {
		if str == nil || *str == "" {
			continue
		}

		*collectionPtr.items = append(
			*collectionPtr.items,
			str)
	}

	return collectionPtr
}

func (collectionPtr *CollectionPtr) AddsNonEmptyPtrLock(
	itemsPtr ...*string,
) *CollectionPtr {
	if itemsPtr == nil {
		return collectionPtr
	}

	for _, str := range itemsPtr {
		if str == nil || *str == "" {
			continue
		}

		collectionPtr.Lock()
		*collectionPtr.items = append(
			*collectionPtr.items,
			str)
		collectionPtr.Unlock()
	}

	return collectionPtr
}

func (collectionPtr *CollectionPtr) UniqueBoolMapLock(
	isAddEmptyStringOnNil bool,
) *map[string]bool {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	return collectionPtr.UniqueBoolMap(isAddEmptyStringOnNil)
}

// continue on nil
func (collectionPtr *CollectionPtr) UniqueBoolMap(
	isAddEmptyStringOnNil bool,
) *map[string]bool {
	respectiveMap := make(
		map[string]bool,
		collectionPtr.Length())

	for _, item := range *collectionPtr.items {
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

func (collectionPtr *CollectionPtr) UniqueListPtr(
	isAddEmptyStringOnNil bool,
) *[]string {
	boolMap := collectionPtr.UniqueBoolMap(isAddEmptyStringOnNil)
	list := make([]string, len(*boolMap))

	i := 0
	for str := range *boolMap {
		list[i] = str
		i++
	}

	return &list
}

func (collectionPtr *CollectionPtr) UniqueListPtrLock(
	isAddEmptyStringOnNil bool,
) *[]string {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	return collectionPtr.UniqueListPtr(isAddEmptyStringOnNil)
}

func (collectionPtr *CollectionPtr) UniqueListLock(
	isAddEmptyStringOnNil bool,
) []string {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	return collectionPtr.UniqueList(isAddEmptyStringOnNil)
}

func (collectionPtr *CollectionPtr) UniqueList(
	isAddEmptyStringOnNil bool,
) []string {
	return *collectionPtr.UniqueListPtr(isAddEmptyStringOnNil)
}

func (collectionPtr *CollectionPtr) List() []*string {
	return *collectionPtr.items
}

// must return a slice
func (collectionPtr *CollectionPtr) FilterSimpleArray(
	filter IsStringPointerFilter,
) *[]string {
	if collectionPtr.IsEmpty() {
		return &([]string{})
	}

	list := make([]string, 0, collectionPtr.Length())

	for _, element := range *collectionPtr.items {
		result, isKeep := filter(element)

		if isKeep && result != nil {
			list = append(list, *result)
		}
	}

	return &list
}

// assumed wg is added to be done here.
//
// must return a slice
func (collectionPtr *CollectionPtr) FilterSimpleArrayLock(
	filter IsStringPointerFilter,
	wg *sync.WaitGroup,
) *[]string {
	copyList := collectionPtr.ListCopyPtrLock()
	length := len(*copyList)
	if length == 0 {
		return &([]string{})
	}

	list := make([]string, 0, length)

	for _, element := range *collectionPtr.items {
		result, isKeep := filter(element)

		if isKeep && result != nil {
			list = append(list, *result)
		}
	}

	wg.Done()

	return &list
}

// must return a slice
func (collectionPtr *CollectionPtr) Filter(
	filter IsStringPointerFilter,
) *[]*string {
	if collectionPtr.IsEmpty() {
		return &([]*string{})
	}

	list := make([]*string, 0, collectionPtr.Length())

	for _, element := range *collectionPtr.items {
		result, isKeep := filter(element)

		if isKeep {
			list = append(list, result)
		}
	}

	return &list
}

// must return a slice
func (collectionPtr *CollectionPtr) FilterLock(
	filter IsStringPointerFilter,
) *[]*string {
	elements := collectionPtr.ListCopyPtrLock()
	length := len(*elements)

	if length == 0 {
		return elements
	}

	list := make([]*string, 0, length)

	for _, element := range *elements {
		result, isKeep := filter(element)

		if isKeep {
			list = append(list, result)
		}
	}

	return &list
}

// must return a items
func (collectionPtr *CollectionPtr) FilteredCollection(
	filter IsStringPointerFilter,
) *CollectionPtr {
	return NewCollectionPtrUsingStrings(collectionPtr.FilterSimpleArray(filter), 0)
}

// must return a items
func (collectionPtr *CollectionPtr) FilteredCollectionLock(
	filter IsStringPointerFilter,
) *CollectionPtr {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	return NewCollectionPtrUsingStrings(collectionPtr.FilterSimpleArray(filter), 0)
}

// must return a slice
func (collectionPtr *CollectionPtr) FilterPtrLock(
	filterPtr IsStringPointerFilter,
) *[]*string {
	elements := collectionPtr.ListCopyPtrLock()
	length := len(*elements)

	if length == 0 {
		return &([]*string{})
	}

	list := make([]*string, 0, length)

	for _, element := range *elements {
		result, isKeep := filterPtr(element)

		if isKeep {
			list = append(list, result)
		}
	}

	return &list
}

// must return a slice
func (collectionPtr *CollectionPtr) FilterPtr(
	filterPtr IsStringPointerFilter,
) *[]*string {
	if collectionPtr.IsEmpty() {
		return &([]*string{})
	}

	list := make([]*string, 0, collectionPtr.Length())

	for _, element := range *collectionPtr.items {
		result, isKeep := filterPtr(element)

		if isKeep {
			list = append(list, result)
		}
	}

	return &list
}

// must return a slice
func (collectionPtr *CollectionPtr) NonEmptySimpleListPtr() *[]string {
	if collectionPtr.IsEmpty() {
		return &([]string{})
	}

	list := make([]string, 0, collectionPtr.Length())

	for _, element := range *collectionPtr.items {
		if element == nil || *element == "" {
			continue
		}

		list = append(list, *element)
	}

	return &list
}

func (collectionPtr *CollectionPtr) Hashset() *Hashset {
	return NewHashsetUsingStringPointersArray(
		collectionPtr.items,
		collectionPtr.Length()*2,
		true)
}

func (collectionPtr *CollectionPtr) HashsetLock() *Hashset {
	return NewHashsetUsingStringPointersArray(
		collectionPtr.ListCopyPtrLock(),
		constants.ArbitraryCapacity100,
		false)
}

func (collectionPtr *CollectionPtr) SimpleList() []string {
	return *converters.PointerStringsToStrings(collectionPtr.items)
}

func (collectionPtr *CollectionPtr) SimpleListPtr() *[]string {
	return converters.PointerStringsToStrings(collectionPtr.items)
}

func (collectionPtr *CollectionPtr) SimpleListPtrLock() *[]string {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	return converters.PointerStringsToStrings(collectionPtr.items)
}

// direct return pointer
func (collectionPtr *CollectionPtr) ListPtr() *[]*string {
	return collectionPtr.items
}

// returns a copy of the items
//
// must return a slice
func (collectionPtr *CollectionPtr) ListCopyPtrLock() *[]*string {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	if collectionPtr.items == nil ||
		*collectionPtr.items == nil {
		return &([]*string{})
	}

	return &(*collectionPtr.items)
}

func (collectionPtr *CollectionPtr) HasLock(str string) bool {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	return collectionPtr.Has(str)
}

func (collectionPtr *CollectionPtr) Has(str string) bool {
	if collectionPtr.IsEmpty() {
		return false
	}

	for _, element := range *collectionPtr.items {
		if element == nil {
			continue
		}

		if *element == str {
			return true
		}
	}

	return false
}

func (collectionPtr *CollectionPtr) HasUsingSensitivity(
	str string, isCaseSensitive bool,
) bool {
	if collectionPtr.IsEmpty() {
		return false
	}

	if isCaseSensitive {
		return collectionPtr.Has(str)
	}

	for _, element := range *collectionPtr.items {
		if element == nil {
			continue
		}

		if strings.EqualFold(*element, str) {
			return true
		}
	}

	return false
}

func (collectionPtr *CollectionPtr) HasAll(items ...string) bool {
	if collectionPtr.IsEmpty() {
		return false
	}

	for _, element := range items {
		if !collectionPtr.IsContainsPtr(&element) {
			return false
		}
	}

	return true
}

func (collectionPtr *CollectionPtr) First() string {
	return *(*collectionPtr.items)[0]
}

func (collectionPtr *CollectionPtr) Single() string {
	length := collectionPtr.Length()
	if length != 1 {
		msgtype.LengthShouldBeEqualToMessage.HandleUsingPanic("1", length)
	}

	return *(*collectionPtr.items)[0]
}

func (collectionPtr *CollectionPtr) Last() string {
	length := collectionPtr.Length()

	return *(*collectionPtr.items)[length-1]
}

func (collectionPtr *CollectionPtr) LastOrDefault() string {
	length := collectionPtr.Length()

	if length == 0 {
		return constants.EmptyString
	}

	return *(*collectionPtr.items)[length-1]
}

func (collectionPtr *CollectionPtr) FirstOrDefault() string {
	if collectionPtr.IsEmpty() {
		return constants.EmptyString
	}

	return *(*collectionPtr.items)[0]
}

func (collectionPtr *CollectionPtr) IndexAt(
	index int,
) string {
	return *(*collectionPtr.items)[index]
}

func (collectionPtr *CollectionPtr) SafePointerIndexAt(
	index int,
) *string {
	length := collectionPtr.Length()
	if length-1 < index {
		return nil
	}

	return (*collectionPtr.items)[index]
}

func (collectionPtr *CollectionPtr) SafePointerIndexAtUsingLength(
	length, index int,
) *string {
	if length-1 < index {
		return nil
	}

	return (*collectionPtr.items)[index]
}

func (collectionPtr *CollectionPtr) SafeIndexAtUsingLength(
	defaultString string, length, index int,
) string {
	if length-1 < index {
		return defaultString
	}

	pointer := (*collectionPtr.items)[index]

	if pointer == nil {
		return defaultString
	}

	return *pointer
}

// Creates new doesn't modify current collection
func (collectionPtr *CollectionPtr) SortedListAsc() *[]string {
	if collectionPtr.IsEmpty() {
		return &[]string{}
	}

	list := collectionPtr.SimpleListPtr()
	sort.Strings(*list)

	return list
}

// mutates current collection
func (collectionPtr *CollectionPtr) SortedAsc() *CollectionPtr {
	if collectionPtr.IsEmpty() {
		return collectionPtr
	}

	strsort.QuickPtr(collectionPtr.items)

	return collectionPtr
}

// mutates current collection
func (collectionPtr *CollectionPtr) SortedAscLock() *CollectionPtr {
	if collectionPtr.IsEmptyLock() {
		return collectionPtr
	}

	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	strsort.QuickPtr(collectionPtr.items)

	return collectionPtr
}

// Creates new one.
func (collectionPtr *CollectionPtr) SortedListDsc() *[]string {
	list := collectionPtr.SortedListAsc()
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
func (collectionPtr *CollectionPtr) SortedDsc() *CollectionPtr {
	list := collectionPtr.items
	length := len(*list)
	mid := length / 2

	for i := 0; i < mid; i++ {
		temp := (*list)[i]
		(*list)[i] = (*list)[length-1-i]
		(*list)[length-1-i] = temp
	}

	return collectionPtr
}

func (collectionPtr *CollectionPtr) IsContainsPtr(
	item *string,
) bool {
	if item == nil || collectionPtr.IsEmpty() {
		return false
	}

	for _, element := range *collectionPtr.items {
		if element == nil {
			continue
		}

		if *element == *item {
			return true
		}
	}

	return false
}

// nil will return false.
func (collectionPtr *CollectionPtr) GetHashsetPlusHasAll(
	items *[]string,
) (*Hashset, bool) {
	hashset := collectionPtr.Hashset()

	if items == nil || collectionPtr.IsEmpty() {
		return hashset, false
	}

	return hashset, hashset.HasAllStringsPtr(items)
}

// nil will return false.
func (collectionPtr *CollectionPtr) IsContainsAllPtr(
	items *[]string,
) bool {
	if items == nil {
		return false
	}

	if collectionPtr.IsEmpty() {
		return false
	}

	for _, item := range *items {
		if !collectionPtr.IsContainsPtr(&item) {
			return false
		}
	}

	return true
}

// nil will return false.
func (collectionPtr *CollectionPtr) IsContainsAll(
	items ...string,
) bool {
	if items == nil {
		return false
	}

	return collectionPtr.IsContainsAllPtr(&items)
}

// nil will return false.
func (collectionPtr *CollectionPtr) IsContainsAllLock(
	items ...string,
) bool {
	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	if items == nil {
		return false
	}

	return collectionPtr.IsContainsAllPtr(&items)
}

func (collectionPtr *CollectionPtr) CharCollectionPtrMap() *CharCollectionMap {
	length := collectionPtr.Length()
	lengthByFourBestGuess := length / 4
	runeMap := NewCharCollectionMap(
		length,
		lengthByFourBestGuess)

	for _, item := range *collectionPtr.items {
		runeMap.AddStringPtr(item)
	}

	return runeMap
}

func (collectionPtr *CollectionPtr) String() string {
	if collectionPtr.IsEmpty() {
		return commonJoiner + NoElements
	}

	return commonJoiner +
		strings.Join(
			*collectionPtr.SimpleListPtr(),
			commonJoiner)
}

func (collectionPtr *CollectionPtr) StringLock() string {
	if collectionPtr.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	collectionPtr.Lock()
	defer collectionPtr.Unlock()

	return commonJoiner +
		strings.Join(
			*collectionPtr.SimpleListPtr(),
			commonJoiner)
}

func (collectionPtr *CollectionPtr) Join(
	separator string,
) string {
	return strings.Join(
		*collectionPtr.SimpleListPtr(),
		separator)
}

func (collectionPtr *CollectionPtr) AddCapacity(
	capacities ...int,
) *CollectionPtr {
	if capacities == nil || len(capacities) == 0 {
		return collectionPtr
	}

	currentCapacity := collectionPtr.Capacity()

	for _, capacity := range capacities {
		currentCapacity += capacity
	}

	return collectionPtr.Resize(currentCapacity)
}

// Only resize if capacity is bigger than the current one
// Warning changes current pointer with new one.
func (collectionPtr *CollectionPtr) Resize(
	newCapacity int,
) *CollectionPtr {
	capacity := collectionPtr.Capacity()
	if capacity >= newCapacity {
		return collectionPtr
	}

	newItems := make([]*string, collectionPtr.Length(), newCapacity)
	copy(newItems, *collectionPtr.items)

	collectionPtr.items = &newItems

	return collectionPtr
}

func (collectionPtr *CollectionPtr) Joins(
	separator string,
	items ...string,
) string {
	newItems := make([]string, 0, collectionPtr.Length()+len(items))

	for _, item := range *collectionPtr.items {
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

func (collectionPtr *CollectionPtr) JsonModel() *CollectionPtrDataModel {
	return &CollectionPtrDataModel{
		Items: collectionPtr.items,
	}
}

func (collectionPtr *CollectionPtr) JsonModelAny() interface{} {
	return collectionPtr.JsonModel()
}

func (collectionPtr *CollectionPtr) MarshalJSON() ([]byte, error) {
	return json.Marshal(*collectionPtr.JsonModel())
}

func (collectionPtr *CollectionPtr) UnmarshalJSON(
	data []byte,
) error {
	var dataModel CollectionPtrDataModel

	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		collectionPtr.items = dataModel.Items
	}

	return err
}

func (collectionPtr *CollectionPtr) Json() *corejson.Result {
	if collectionPtr.IsEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(collectionPtr)

	return corejson.NewPtr(jsonBytes, err)
}

func (collectionPtr *CollectionPtr) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*CollectionPtr, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return EmptyCollectionPtr(), nil
	}

	err := json.Unmarshal(*jsonResult.Bytes, &collectionPtr)

	if err != nil {
		return EmptyCollectionPtr(), err
	}

	return collectionPtr, nil
}

// Panic if error
func (collectionPtr *CollectionPtr) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *CollectionPtr {
	newUsingJson, err :=
		collectionPtr.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

// Panic if error
func (collectionPtr *CollectionPtr) JsonParseSelfInject(
	jsonResult *corejson.Result,
) {
	collectionPtr.ParseInjectUsingJsonMust(jsonResult)
}

func (collectionPtr *CollectionPtr) AsJsoner() *corejson.Jsoner {
	var jsoner corejson.Jsoner = collectionPtr

	return &jsoner
}

func (collectionPtr *CollectionPtr) AsJsonParseSelfInjector() *corejson.ParseSelfInjector {
	var jsonInjector corejson.ParseSelfInjector = collectionPtr

	return &jsonInjector
}

func (collectionPtr *CollectionPtr) AsJsonMarshaller() *corejson.JsonMarshaller {
	var jsonMarshaller corejson.JsonMarshaller = collectionPtr

	return &jsonMarshaller
}
