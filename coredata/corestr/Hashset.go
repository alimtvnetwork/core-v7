package corestr

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/internal/utilstringinternal"
)

type Hashset struct {
	hasMapUpdated bool
	isEmptySet    bool
	length        int
	items         map[string]bool
	cachedList    []string
	sync.Mutex
}

func (it *Hashset) IsEmpty() bool {
	if it == nil {
		return true
	}

	if it.hasMapUpdated {
		it.isEmptySet = len(it.items) == 0
	}

	return it.isEmptySet
}

func (it *Hashset) HasItems() bool {
	return !it.IsEmpty()
}

// AddCapacitiesLock Changing capacity creates new map and points to it.
// There is memory copy and loop is performed.
func (it *Hashset) AddCapacitiesLock(
	capacities ...int,
) *Hashset {
	length := it.LengthLock()

	if len(capacities) == 0 {
		return it
	}

	for _, capacity := range capacities {
		length += capacity
	}

	return it.ResizeLock(length)
}

// AddCapacities Changing capacity creates new map and points to it.
// There is memory copy and loop is performed.
func (it *Hashset) AddCapacities(
	capacities ...int,
) *Hashset {
	length := it.Length()

	if len(capacities) == 0 {
		return it
	}

	for _, capacity := range capacities {
		length += capacity
	}

	return it.Resize(length)
}

// Resize Changing capacity creates new map and points to it.
// There is memory copy and loop is performed.
func (it *Hashset) Resize(capacity int) *Hashset {
	length := it.Length()

	if length > capacity {
		return it
	}

	newItemsMap := make(map[string]bool, capacity)

	for val := range it.items {
		newItemsMap[val] = true
	}

	it.items = newItemsMap

	return it
}

// ResizeLock Changing capacity creates new map and points to it.
// There is memory copy and loop is performed.
func (it *Hashset) ResizeLock(capacity int) *Hashset {
	length := it.LengthLock()

	if length > capacity {
		return it
	}

	newItemsMap := make(map[string]bool, capacity)

	for val := range it.items {
		newItemsMap[val] = true
	}

	it.Lock()
	it.items = newItemsMap
	it.Unlock()

	return it
}

func (it *Hashset) Collection() *Collection {
	return NewCollectionUsingStrings(it.List(), false)
}

func (it *Hashset) IsEmptyLock() bool {
	it.Lock()
	defer it.Unlock()

	return it.IsEmpty()
}

func (it *Hashset) ConcatNewHashsets(
	isCloneCurrentOnEmpty bool,
	hashsets ...*Hashset,
) *Hashset {
	isEmpty := hashsets == nil || len(hashsets) == 0

	if isEmpty {
		return NewHashsetUsingMap(
			it.items,
			constants.Zero,
			isCloneCurrentOnEmpty)
	}

	length := it.Length() + constants.Capacity4

	for _, h := range hashsets {
		if h == nil {
			continue
		}

		length += h.Length()
	}

	newHashset := NewHashsetUsingMap(
		it.items,
		length,
		isCloneCurrentOnEmpty)

	newHashset.AddHashsetItems(it)

	for _, h := range hashsets {
		newHashset.AddHashsetItems(h)
	}

	return newHashset
}

func (it *Hashset) ConcatNewStringsPointers(
	isCloneCurrentOnEmpty bool,
	stringsOfStringsItems ...*[]string,
) *Hashset {
	isEmpty := len(stringsOfStringsItems) == 0

	if isEmpty {
		return NewHashsetUsingMap(
			it.items,
			constants.Zero,
			isCloneCurrentOnEmpty)
	}

	length := AllIndividualItemsStringsOfStringsPointerLength(&stringsOfStringsItems) +
		it.Length() +
		constants.Capacity4

	newHashset := NewHashsetUsingMap(
		it.items,
		length,
		isCloneCurrentOnEmpty)

	newHashset.AddHashsetItems(it)

	for _, stringsItems := range stringsOfStringsItems {
		newHashset.AddStringsPtr(stringsItems)
	}

	return newHashset
}

func (it *Hashset) ConcatNewStrings(
	isCloneCurrentOnEmpty bool,
	stringsOfStringsItems ...[]string,
) *Hashset {
	isEmpty := len(stringsOfStringsItems) == 0

	if isEmpty {
		return NewHashsetUsingMap(
			it.items,
			constants.Zero,
			isCloneCurrentOnEmpty)
	}

	length := AllIndividualStringsOfStringsLength(&stringsOfStringsItems) +
		it.Length() +
		constants.Capacity4
	newHashset := NewHashsetUsingMap(
		it.items,
		length,
		true)

	newHashset.AddHashsetItems(it)

	for _, stringsItems := range stringsOfStringsItems {
		newHashset.AddStrings(stringsItems)
	}

	return newHashset
}

func (it *Hashset) AddPtr(key *string) *Hashset {
	it.items[*key] = true
	it.hasMapUpdated = true

	return it
}

func (it *Hashset) AddWithWgLock(
	key string,
	group *sync.WaitGroup,
) *Hashset {
	it.Lock()
	it.items[key] = true
	it.Unlock()

	it.hasMapUpdated = true

	group.Done()

	return it
}

func (it *Hashset) AddPtrLock(key *string) *Hashset {
	it.Lock()
	it.items[*key] = true
	it.Unlock()

	it.hasMapUpdated = true

	return it
}

func (it *Hashset) Add(key string) *Hashset {
	it.items[key] = true
	it.hasMapUpdated = true

	return it
}

func (it *Hashset) AddNonEmpty(str string) *Hashset {
	if str == "" {
		return it
	}

	return it
}

func (it *Hashset) AddNonEmptyWhitespace(str string) *Hashset {
	if utilstringinternal.IsEmptyOrWhitespace(str) {
		return it
	}

	return it.Add(str)
}

func (it *Hashset) AddIf(isAdd bool, addingString string) *Hashset {
	if !isAdd {
		return it
	}

	return it.Add(addingString)
}

func (it *Hashset) AddIfMany(
	isAdd bool,
	addingStrings ...string,
) *Hashset {
	if !isAdd {
		return it
	}

	return it.Adds(addingStrings...)
}

func (it *Hashset) AddFunc(f func() string) *Hashset {
	return it.Add(f())
}

func (it *Hashset) AddFuncErr(
	funcReturnsError func() (result string, err error),
	errHandler func(errInput error),
) *Hashset {
	r, err := funcReturnsError()

	if err != nil {
		errHandler(err)

		return it
	}

	return it.Add(r)
}

func (it *Hashset) AddStringsPtrWgLock(
	keys *[]string, wg *sync.WaitGroup,
) *Hashset {
	if keys == nil {
		return it
	}

	length := len(*keys)

	if length > it.length || length > constants.ArbitraryCapacity100 {
		it.AddCapacitiesLock(length*2, constants.ArbitraryCapacity100)
	}

	it.Lock()
	for _, key := range *keys {
		it.items[key] = true
	}

	it.Unlock()
	wg.Done()

	it.hasMapUpdated = true

	return it
}

func (it *Hashset) AddHashsetItems(
	hashsetAdd *Hashset,
) *Hashset {
	if hashsetAdd == nil {
		return it
	}

	length := hashsetAdd.Length()

	if length > it.length || length > constants.ArbitraryCapacity100 {
		it.AddCapacities(length*2, constants.ArbitraryCapacity100)
	}

	for k := range hashsetAdd.items {
		it.items[k] = true
	}

	it.hasMapUpdated = true

	return it
}

// AddItemsMap only add if the value is true
func (it *Hashset) AddItemsMap(
	itemsMap map[string]bool,
) *Hashset {
	if itemsMap == nil {
		return it
	}

	length := len(itemsMap)

	if length > it.length || length > constants.ArbitraryCapacity100 {
		it.AddCapacities(length*2, constants.ArbitraryCapacity100)
	}

	for k, isEnabled := range itemsMap {
		if !isEnabled {
			continue
		}

		it.items[k] = true
	}

	it.hasMapUpdated = true

	return it
}

// AddItemsMapWgLock only add if the value is true
// Assume that wg already enqueued the job as wg.Add(...) done already.
func (it *Hashset) AddItemsMapWgLock(
	itemsMap *map[string]bool,
	wg *sync.WaitGroup,
) *Hashset {
	if itemsMap == nil {
		return it
	}

	length := len(*itemsMap)

	if length > it.length || length > constants.ArbitraryCapacity100 {
		it.AddCapacitiesLock(length*2, constants.ArbitraryCapacity100)
	}

	for k, isEnabled := range *itemsMap {
		if !isEnabled {
			continue
		}

		it.Lock()
		it.items[k] = true
		it.Unlock()
	}

	wg.Done()

	it.hasMapUpdated = true

	return it
}

func (it *Hashset) AddHashsetWgLock(
	hashsetAdd *Hashset,
	wg *sync.WaitGroup,
) *Hashset {
	if hashsetAdd == nil {
		return it
	}

	length := hashsetAdd.LengthLock()

	if length > it.length || length > constants.ArbitraryCapacity100 {
		it.AddCapacitiesLock(length*2, constants.ArbitraryCapacity100)
	}

	it.Lock()
	for k := range hashsetAdd.items {
		it.items[k] = true
	}

	it.Unlock()
	wg.Done()

	it.hasMapUpdated = true

	return it
}

func (it *Hashset) AddStringsPtr(keys *[]string) *Hashset {
	if keys == nil {
		return it
	}

	for _, key := range *keys {
		it.items[key] = true
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashset) AddStrings(keys []string) *Hashset {
	if len(keys) == 0 {
		return it
	}

	return it.AddStringsPtr(&keys)
}

func (it *Hashset) AddStringsPtrLock(keys *[]string) *Hashset {
	if keys == nil {
		return it
	}

	it.Lock()
	for _, key := range *keys {
		it.items[key] = true
	}

	it.Unlock()

	it.hasMapUpdated = true

	return it
}

func (it *Hashset) Adds(keys ...string) *Hashset {
	if keys == nil {
		return it
	}

	for _, key := range keys {
		it.items[key] = true
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashset) AddCollection(
	collection *Collection,
) *Hashset {
	if collection == nil || collection.IsEmpty() {
		return it
	}

	for _, element := range collection.items {
		it.items[element] = true
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashset) AddCollections(
	collections ...*Collection,
) *Hashset {
	if collections == nil {
		return it
	}

	for _, collection := range collections {
		if collection == nil || collection.IsEmpty() {
			continue
		}

		for _, element := range collection.items {
			it.items[element] = true
		}
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashset) AddsAnyUsingFilter(
	filter IsStringFilter,
	anys ...interface{},
) *Hashset {
	if anys == nil {
		return it
	}

	for i, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)
		result, isKeep, isBreak := filter(anyStr, i)

		if isKeep {
			it.items[result] = true
			it.hasMapUpdated = true
		}

		if isBreak {
			return it
		}
	}

	return it
}

func (it *Hashset) AddsAnyUsingFilterLock(
	filter IsStringFilter,
	anys ...interface{},
) *Hashset {
	if anys == nil {
		return it
	}

	for i, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(
			constants.SprintValueFormat,
			any)

		result, isKeep, isBreak := filter(anyStr, i)

		if isKeep {
			it.Lock()
			it.items[result] = true
			it.Unlock()

			it.hasMapUpdated = true
		}

		if isBreak {
			return it
		}
	}

	return it
}

func (it *Hashset) AddsUsingFilter(
	filter IsStringFilter,
	keys ...string,
) *Hashset {
	if keys == nil {
		return it
	}

	for i, key := range keys {
		result, isKeep, isBreak := filter(key, i)

		if isKeep {
			it.items[result] = true
			it.hasMapUpdated = true
		}

		if isBreak {
			return it
		}
	}

	return it
}

func (it *Hashset) AddLock(key string) *Hashset {
	it.Lock()
	defer it.Unlock()

	it.items[key] = true
	it.hasMapUpdated = true

	return it
}

func (it *Hashset) HasAnyItem() bool {
	return it != nil && it.Length() > 0
}

func (it *Hashset) IsMissing(key string) bool {
	_, isFound := it.items[key]

	return !isFound
}

func (it *Hashset) IsMissingLock(key string) bool {
	it.Lock()
	_, isFound := it.items[key]
	it.Unlock()

	return !isFound
}

func (it *Hashset) Has(key string) bool {
	isSet, isFound := it.items[key]

	return isFound && isSet
}

func (it *Hashset) HasLock(key string) bool {
	it.Lock()
	isSet, isFound := it.items[key]
	it.Unlock()

	return isFound && isSet
}

func (it *Hashset) HasAllStringsPtr(keys *[]string) bool {
	for _, key := range *keys {
		isSet, isFound := it.items[key]

		if !(isFound && isSet) {
			// not found
			return false
		}
	}

	// all found.
	return true
}

// HasAllCollectionItems return false on items is nil or empty.
func (it *Hashset) HasAllCollectionItems(
	collection *Collection,
) bool {
	if collection == nil || collection.IsEmpty() {
		return false
	}

	return it.HasAllStringsPtr(collection.ListPtr())
}

func (it *Hashset) HasAll(keys ...string) bool {
	for _, key := range keys {
		isSet, isFound := it.items[key]

		if !(isFound && isSet) {
			// not found
			return false
		}
	}

	// all found.
	return true
}

func (it *Hashset) IsAllMissing(keys ...string) bool {
	for _, key := range keys {
		isSet, isFound := it.items[key]

		if isFound && isSet {
			// found
			return false
		}
	}

	// all not found.
	return true
}

func (it *Hashset) HasAny(keys ...string) bool {
	for _, key := range keys {
		isSet, isFound := it.items[key]

		if isFound && isSet {
			// any found
			return true
		}
	}

	// all not found.
	return false
}

func (it *Hashset) HasWithLock(key string) bool {
	it.Lock()
	defer it.Unlock()

	isSet, isFound := it.items[key]

	return isFound && isSet
}

func (it *Hashset) OrderedList() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return it.
		Collection().
		SortedAsc().
		items
}

// GetFilteredItems must return slice.
func (it *Hashset) GetFilteredItems(
	filter IsStringFilter,
) *[]string {
	if it.IsEmpty() {
		return &([]string{})
	}

	filteredList := make(
		[]string,
		0,
		it.Length())

	i := 0
	for key := range it.items {
		result, isKeep, isBreak := filter(key, i)
		i++

		if !isKeep {
			continue
		}

		filteredList = append(
			filteredList,
			result)

		if isBreak {
			return &filteredList
		}
	}

	return &filteredList
}

// GetFilteredCollection must return items.
func (it *Hashset) GetFilteredCollection(
	filter IsStringFilter,
) *Collection {
	if it.IsEmpty() {
		return EmptyCollection()
	}

	filteredList := make(
		[]string,
		0,
		it.Length())

	i := 0
	for key := range it.items {
		result, isKeep, isBreak := filter(key, i)
		i++

		if !isKeep {
			continue
		}

		filteredList = append(
			filteredList,
			result)

		if isBreak {
			return NewCollectionUsingStrings(
				filteredList, false)
		}
	}

	return NewCollectionUsingStrings(
		filteredList, false)
}

// GetAllExceptHashset Get all hashset items except the mentioned ones in anotherHashset.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashset
// Set B = anotherHashset given in parameters.
func (it *Hashset) GetAllExceptHashset(
	anotherHashset *Hashset,
) *[]string {
	if anotherHashset == nil || anotherHashset.IsEmpty() {
		return it.ListPtr()
	}

	finalList := make(
		[]string,
		0,
		it.Length())

	for item := range it.items {
		if anotherHashset.Has(item) {
			continue
		}

		finalList = append(
			finalList,
			item)
	}

	return &finalList
}

// GetAllExcept Get all hashset items except the mentioned ones in items.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashset
// Set B = items given in parameters.
func (it *Hashset) GetAllExcept(
	items *[]string,
) *[]string {
	if items == nil {
		return it.ListPtr()
	}

	newHashset := NewHashsetUsingStrings(
		items)

	return it.GetAllExceptHashset(
		newHashset)
}

// GetAllExceptCollection Get all hashset items except the mentioned ones in collection.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashset
// Set B = collection given in parameters.
func (it *Hashset) GetAllExceptCollection(
	collection *Collection,
) *[]string {
	if collection == nil {
		return it.ListPtr()
	}

	return it.GetAllExceptHashset(
		collection.HashsetAsIs())
}

// GetAllExceptCollectionPtr Get all hashset items except the mentioned ones in collectionPtr.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashset
// Set B = collectionPtr given in parameters.
func (it *Hashset) GetAllExceptCollectionPtr(
	collectionPtr *CollectionPtr,
) *[]string {
	if collectionPtr == nil {
		return it.ListPtr()
	}

	return it.GetAllExceptHashset(
		collectionPtr.HashsetAsIs())
}

func (it *Hashset) Items() map[string]bool {
	return it.items
}

func (it *Hashset) List() []string {
	return *it.ListPtr()
}

func (it *Hashset) JoinSorted(joiner string) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	list := it.ListPtr()
	sort.Strings(*list)

	return strings.Join(*list, joiner)
}

func (it *Hashset) ListPtrSortedAsc() *[]string {
	list := it.ListPtr()
	sort.Strings(*list)

	return list
}

func (it *Hashset) ListPtrSortedDsc() *[]string {
	list := it.ListPtr()
	sort.Strings(*list)

	return stringslice.InPlaceReverse(list)
}

func (it *Hashset) ListPtr() *[]string {
	if it.hasMapUpdated || it.cachedList == nil {
		it.setCached()
	}

	return &it.cachedList
}

// ListCopyPtrLock a slice must returned
func (it *Hashset) ListCopyPtrLock() *[]string {
	it.Lock()
	defer it.Unlock()
	cloned := *it.ListPtr()

	return &cloned
}

func (it *Hashset) setCached() {
	length := it.Length()
	list := make([]string, length)

	i := 0

	for key := range it.items {
		list[i] = key
		i++
	}

	it.hasMapUpdated = false
	it.cachedList = list
}

// ToLowerSet Create a new items with all lower strings
func (it *Hashset) ToLowerSet() *Hashset {
	length := it.Length()
	newMap := make(map[string]bool, length)

	var toLower string
	for key, isEnabled := range it.items {
		toLower = strings.ToLower(key)
		newMap[toLower] = isEnabled
	}

	return NewHashsetUsingMap(
		newMap,
		length,
		false)
}

func (it *Hashset) Length() int {
	if it == nil {
		return 0
	}

	if it.hasMapUpdated {
		it.length = len(it.items)
	}

	return it.length
}

func (it *Hashset) LengthLock() int {
	it.Lock()
	defer it.Unlock()

	return it.Length()
}

//goland:noinspection GoVetCopyLock
func (it *Hashset) IsEquals(another Hashset) bool {
	return it.IsEqualsPtr(&another)
}

func (it *Hashset) IsEqualsPtrLock(another *Hashset) bool {
	it.Lock()
	defer it.Unlock()

	return it.IsEqualsPtr(another)
}

func (it *Hashset) IsEqualsPtr(another *Hashset) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it == another {
		// ptr same
		return true
	}

	if it.IsEmpty() && another.IsEmpty() {
		return true
	}

	if it.IsEmpty() || another.IsEmpty() {
		return false
	}

	leftLength := it.Length()
	rightLength := another.Length()

	if leftLength != rightLength {
		return false
	}

	for key := range it.items {
		isRes, has := another.items[key]

		if !has || !isRes {
			return false
		}
	}

	return true
}

func (it *Hashset) Remove(key string) *Hashset {
	delete(it.items, key)
	it.hasMapUpdated = true

	return it
}

func (it *Hashset) RemoveWithLock(key string) *Hashset {
	it.Lock()
	defer it.Unlock()

	it.Remove(key)

	return it
}

func (it *Hashset) String() string {
	if it.IsEmpty() {
		return commonJoiner + NoElements
	}

	return commonJoiner +
		strings.Join(
			it.List(),
			commonJoiner)
}

func (it *Hashset) StringLock() string {
	if it.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	it.Lock()
	defer it.Unlock()

	return commonJoiner +
		strings.Join(
			*it.ListPtr(),
			commonJoiner)
}

func (it *Hashset) Join(
	joiner string,
) string {
	return strings.Join(*it.ListPtr(), joiner)
}

func (it *Hashset) NonEmptyJoins(
	joiner string,
) string {
	return stringslice.NonEmptyJoinPtr(
		it.ListPtr(),
		joiner)
}

func (it *Hashset) NonWhitespaceJoins(
	joiner string,
) string {
	return stringslice.NonWhitespaceJoinPtr(
		it.ListPtr(),
		joiner)
}

//goland:noinspection GoLinterLocal
func (it *Hashset) JsonModel() *HashsetDataModel {
	return NewHashsetsDataModelUsing(it)
}

//goland:noinspection GoLinterLocal
func (it *Hashset) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *Hashset) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *Hashset) UnmarshalJSON(data []byte) error {
	var dataModel HashsetDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.items = dataModel.Items
		it.length = -1
		it.hasMapUpdated = true
		it.isEmptySet = false
	}

	return err
}

func (it Hashset) Json() corejson.Result {
	return corejson.NewFromAny(it)
}

func (it Hashset) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *Hashset) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Hashset, error) {
	err := jsonResult.Unmarshal(&it)

	if err != nil {
		return EmptyHashset(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *Hashset) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Hashset {
	hashSet, err := it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}

func (it *Hashset) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Hashset) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *Hashset) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *Hashset) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}
