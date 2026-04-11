package corestr

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/internal/mapdiffinternal"
	"github.com/alimtvnetwork/core/internal/strutilinternal"
)

type Hashset struct {
	hasMapUpdated bool
	items         map[string]bool
	cachedList    []string
	sync.RWMutex
}

func (it *Hashset) IsEmpty() bool {
	return it == nil || len(it.items) == 0
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
	return New.Collection.StringsOptions(false, it.List())
}

func (it *Hashset) IsEmptyLock() bool {
	it.RLock()
	defer it.RUnlock()

	return it.IsEmpty()
}

func (it *Hashset) ConcatNewHashsets(
	isCloneCurrentOnEmpty bool,
	hashsets ...*Hashset,
) *Hashset {
	isEmpty := hashsets == nil || len(hashsets) == 0

	if isEmpty {
		return New.Hashset.UsingMapOption(
			constants.Zero,
			isCloneCurrentOnEmpty,
			it.items,
		)
	}

	length := it.Length() + constants.Capacity4

	for _, h := range hashsets {
		if h == nil {
			continue
		}

		length += h.Length()
	}

	newHashset := New.Hashset.UsingMapOption(
		length,
		isCloneCurrentOnEmpty,
		it.items,
	)

	newHashset.AddHashsetItems(it)

	for _, h := range hashsets {
		newHashset.AddHashsetItems(h)
	}

	return newHashset
}

func (it *Hashset) ConcatNewStrings(
	isCloneCurrentOnEmpty bool,
	stringsOfStringsItems ...[]string,
) *Hashset {
	isEmpty := len(stringsOfStringsItems) == 0

	if isEmpty {
		return New.Hashset.UsingMapOption(
			constants.Zero,
			isCloneCurrentOnEmpty,
			it.items,
		)
	}

	length := AllIndividualStringsOfStringsLength(&stringsOfStringsItems) +
		it.Length() +
		constants.Capacity4
	newHashset := New.Hashset.UsingMapOption(
		length,
		true,
		it.items,
	)

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
	it.hasMapUpdated = true
	it.Unlock()

	group.Done()

	return it
}

func (it *Hashset) AddPtrLock(key *string) *Hashset {
	it.Lock()
	it.items[*key] = true
	it.hasMapUpdated = true
	it.Unlock()

	return it
}

func (it *Hashset) Add(key string) *Hashset {
	it.items[key] = true
	it.hasMapUpdated = true

	return it
}

func (it *Hashset) AddBool(key string) (isExist bool) {
	_, has := it.items[key]
	isNew := !has

	if isNew {
		it.items[key] = true
		it.hasMapUpdated = true
	}

	return has
}

func (it *Hashset) AddNonEmpty(str string) *Hashset {
	if str == "" {
		return it
	}

	return it.Add(str)
}

func (it *Hashset) AddNonEmptyWhitespace(str string) *Hashset {
	if strutilinternal.IsEmptyOrWhitespace(str) {
		return it
	}

	return it.Add(str)
}

func (it *Hashset) AddIf(isAdd bool, addingString string) *Hashset {
	isSkip := !isAdd

	if isSkip {
		return it
	}

	return it.Add(addingString)
}

func (it *Hashset) AddIfMany(
	isAdd bool,
	addingStrings ...string,
) *Hashset {
	isSkip := !isAdd

	if isSkip {
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
	keys []string, wg *sync.WaitGroup,
) *Hashset {
	length := len(keys)

	if length > len(it.items) || length > constants.ArbitraryCapacity100 {
		it.AddCapacitiesLock(length*2, constants.ArbitraryCapacity100)
	}

	it.RLock()
	for _, key := range keys {
		it.items[key] = true
	}

	it.hasMapUpdated = true
	it.RUnlock()

	wg.Done()

	return it
}

func (it *Hashset) AddHashsetItems(
	hashsetAdd *Hashset,
) *Hashset {
	if hashsetAdd == nil {
		return it
	}

	length := hashsetAdd.Length()

	if length > len(it.items) || length > constants.ArbitraryCapacity100 {
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

	if length > len(it.items) || length > constants.ArbitraryCapacity100 {
		it.AddCapacities(length*2, constants.ArbitraryCapacity100)
	}

	for k, isEnabled := range itemsMap {
		isDisabled := !isEnabled

		if isDisabled {
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

	if length > len(it.items) || length > constants.ArbitraryCapacity100 {
		it.AddCapacitiesLock(length*2, constants.ArbitraryCapacity100)
	}

	it.RLock()
	for k, isEnabled := range *itemsMap {
		isDisabled := !isEnabled

		if isDisabled {
			continue
		}

		it.items[k] = true
	}

	it.hasMapUpdated = true
	it.RUnlock()

	wg.Done()

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

	if length > len(it.items) || length > constants.ArbitraryCapacity100 {
		it.AddCapacitiesLock(length*2, constants.ArbitraryCapacity100)
	}

	it.RLock()
	for k := range hashsetAdd.items {
		it.items[k] = true
	}

	it.hasMapUpdated = true
	it.RUnlock()

	wg.Done()

	return it
}

func (it *Hashset) AddStrings(keys []string) *Hashset {
	if keys == nil {
		return it
	}

	for _, key := range keys {
		it.items[key] = true
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashset) AddSimpleSlice(simpleSlice *SimpleSlice) *Hashset {
	if simpleSlice.IsEmpty() {
		return it
	}

	return it.Adds(*simpleSlice...)
}

func (it *Hashset) AddStringsLock(keys []string) *Hashset {
	if keys == nil {
		return it
	}

	it.RLock()
	for _, key := range keys {
		it.items[key] = true
	}

	it.hasMapUpdated = true
	it.RUnlock()

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
	anys ...any,
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
	anys ...any,
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
			any,
		)

		result, isKeep, isBreak := filter(anyStr, i)

		if isKeep {
			it.RLock()
			it.items[result] = true
			it.hasMapUpdated = true
			it.RUnlock()
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
	it.RLock()
	_, isFound := it.items[key]
	it.RUnlock()

	return !isFound
}

func (it *Hashset) Has(key string) bool {
	isSet, isFound := it.items[key]

	return isFound && isSet
}

// Contains is an alias for Has.
func (it *Hashset) Contains(key string) bool {
	return it.Has(key)
}

// IsEqual is an alias for IsEquals.
func (it *Hashset) IsEqual(another *Hashset) bool {
	return it.IsEquals(another)
}

// SortedList returns the list of keys sorted in ascending order.
func (it *Hashset) SortedList() []string {
	list := it.List()
	sorted := make([]string, len(list))
	copy(sorted, list)
	sort.Strings(sorted)

	return sorted
}

// Filter returns a new Hashset containing only keys for which the predicate returns true.
func (it *Hashset) Filter(predicate func(string) bool) *Hashset {
	result := New.Hashset.Cap(it.Length())

	for key, isSet := range it.items {
		if isSet && predicate(key) {
			result.Add(key)
		}
	}

	return result
}

func (it *Hashset) HasLock(key string) bool {
	it.RLock()
	isSet, isFound := it.items[key]
	it.RUnlock()

	return isFound && isSet
}

func (it *Hashset) HasAllStrings(keys []string) bool {
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

// HasAllCollectionItems return false on items is nil or Empty.
func (it *Hashset) HasAllCollectionItems(
	collection *Collection,
) bool {
	if collection == nil || collection.IsEmpty() {
		return false
	}

	return it.HasAllStrings(collection.List())
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
	it.RLock()
	defer it.RUnlock()

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

func (it *Hashset) SafeStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return it.List()
}

func (it *Hashset) Lines() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return it.List()
}

func (it *Hashset) SimpleSlice() *SimpleSlice {
	if it.IsEmpty() {
		return Empty.SimpleSlice()
	}

	var list SimpleSlice = it.List()

	return &list
}

// GetFilteredItems must return slice.
func (it *Hashset) GetFilteredItems(
	filter IsStringFilter,
) []string {
	if it.IsEmpty() {
		return []string{}
	}

	filteredList := make(
		[]string,
		0,
		it.Length(),
	)

	i := 0
	for key := range it.items {
		result, isKeep, isBreak := filter(key, i)
		i++
		isSkip := !isKeep

		if isSkip {
			continue
		}

		filteredList = append(
			filteredList,
			result,
		)

		if isBreak {
			return filteredList
		}
	}

	return filteredList
}

// GetFilteredCollection must return items.
func (it *Hashset) GetFilteredCollection(
	filter IsStringFilter,
) *Collection {
	if it.IsEmpty() {
		return Empty.Collection()
	}

	filteredList := make(
		[]string,
		0,
		it.Length(),
	)

	i := 0
	for key := range it.items {
		result, isKeep, isBreak := filter(key, i)
		i++
		isSkip := !isKeep

		if isSkip {
			continue
		}

		filteredList = append(
			filteredList,
			result,
		)

		if isBreak {
			return New.Collection.StringsOptions(
				false,
				filteredList,
			)
		}
	}

	return New.Collection.StringsOptions(
		false,
		filteredList,
	)
}

// GetAllExceptHashset Get all hashset items except the mentioned ones in anotherHashset.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashset
// Set B = anotherHashset given in parameters.
func (it *Hashset) GetAllExceptHashset(
	anotherHashset *Hashset,
) []string {
	if anotherHashset == nil || anotherHashset.IsEmpty() {
		return it.List()
	}

	finalList := make(
		[]string,
		0,
		it.Length(),
	)

	for item := range it.items {
		if anotherHashset.Has(item) {
			continue
		}

		finalList = append(
			finalList,
			item,
		)
	}

	return finalList
}

// GetAllExcept Get all hashset items except the mentioned ones in items.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashset
// Set B = items given in parameters.
func (it *Hashset) GetAllExcept(
	items []string,
) []string {
	if items == nil {
		return it.List()
	}

	newHashset := New.Hashset.Strings(
		items,
	)

	return it.GetAllExceptHashset(
		newHashset,
	)
}

func (it *Hashset) GetAllExceptSpread(
	items ...string,
) []string {
	if items == nil {
		return it.List()
	}

	newHashset := New.Hashset.Strings(
		items,
	)

	return it.GetAllExceptHashset(
		newHashset,
	)
}

// GetAllExceptCollection Get all hashset items except the mentioned ones in collection.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashset
// Set B = collection given in parameters.
func (it *Hashset) GetAllExceptCollection(
	collection *Collection,
) []string {
	if collection == nil {
		return it.List()
	}

	return it.GetAllExceptHashset(
		collection.HashsetAsIs(),
	)
}

func (it *Hashset) Items() map[string]bool {
	return it.items
}

func (it *Hashset) List() []string {
	if it.hasMapUpdated || it.cachedList == nil {
		it.setCached()
	}

	return it.cachedList
}

func (it *Hashset) MapStringAny() map[string]any {
	if it.IsEmpty() {
		return map[string]any{}
	}

	newMap := make(
		map[string]any,
		it.Length()+1,
	)

	for name, isSet := range it.items {
		newMap[name] = isSet
	}

	return newMap
}

func (it *Hashset) MapStringAnyDiff() mapdiffinternal.MapStringAnyDiff {
	return it.MapStringAny()
}

func (it *Hashset) JoinSorted(joiner string) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	list := it.List()
	sort.Strings(list)

	return strings.Join(list, joiner)
}

func (it *Hashset) ListPtrSortedAsc() []string {
	list := it.List()
	sort.Strings(list)

	return list
}

func (it *Hashset) ListPtrSortedDsc() []string {
	list := it.List()
	sort.Strings(list)

	return *stringslice.InPlaceReverse(&list)
}

func (it *Hashset) ListPtr() []string {
	return it.List()
}

func (it *Hashset) Clear() *Hashset {
	if it == nil {
		return it
	}

	it.items = nil
	it.items = make(map[string]bool)
	it.cachedList = []string{}
	it.hasMapUpdated = true

	return it
}

func (it *Hashset) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.items = nil
	it.cachedList = nil
}

// ListCopyLock a slice must returned
func (it *Hashset) ListCopyLock() []string {
	it.RLock()
	defer it.RUnlock()
	cloned := it.List()

	return cloned
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

// ToLowerSet CreateUsingAliasMap a new items with all lower strings
func (it *Hashset) ToLowerSet() *Hashset {
	length := it.Length()
	newMap := make(map[string]bool, length)

	var toLower string
	for key, isEnabled := range it.items {
		toLower = strings.ToLower(key)
		newMap[toLower] = isEnabled
	}

	return New.Hashset.UsingMapOption(
		length,
		false,
		newMap,
	)
}

func (it *Hashset) Length() int {
	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

func (it *Hashset) LengthLock() int {
	it.RLock()
	defer it.RUnlock()

	return it.Length()
}

func (it *Hashset) IsEqualsLock(another *Hashset) bool {
	it.RLock()
	defer it.RUnlock()

	return it.IsEquals(another)
}

func (it *Hashset) IsEquals(another *Hashset) bool {
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

func (it *Hashset) SafeRemove(key string) *Hashset {
	if it.Has(key) {
		delete(it.items, key)
		it.hasMapUpdated = true
	}

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
			commonJoiner,
		)
}

func (it *Hashset) StringLock() string {
	if it.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	it.RLock()
	defer it.RUnlock()

	return commonJoiner +
		strings.Join(
			it.List(),
			commonJoiner,
		)
}

func (it Hashset) Join(
	joiner string,
) string {
	return strings.Join(it.List(), joiner)
}

func (it Hashset) NonEmptyJoins(
	joiner string,
) string {
	return stringslice.NonEmptyJoinPtr(
		it.List(),
		joiner,
	)
}

func (it Hashset) NonWhitespaceJoins(
	joiner string,
) string {
	return stringslice.NonWhitespaceJoinPtr(
		it.List(),
		joiner,
	)
}

//goland:noinspection GoLinterLocal
func (it Hashset) JsonModel() map[string]bool {
	if it.IsEmpty() {
		return map[string]bool{}
	}

	return it.items
}

//goland:noinspection GoLinterLocal
func (it Hashset) JsonModelAny() any {
	return it.JsonModel()
}

func (it Hashset) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *Hashset) UnmarshalJSON(data []byte) error {
	var dataModelItems map[string]bool
	err := json.Unmarshal(data, &dataModelItems)

	if err == nil {
		it.items = dataModelItems
		it.hasMapUpdated = true
		it.cachedList = nil
	}

	return err
}

func (it Hashset) Json() corejson.Result {
	return corejson.New(it)
}

func (it Hashset) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *Hashset) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Hashset, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return New.Hashset.Empty(), err
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

func (it *Hashset) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
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

func (it *Hashset) DistinctDiffLinesRaw(
	rightLines ...string,
) []string {
	isLeftEmpty := it.IsEmpty()

	if isLeftEmpty && len(rightLines) == 0 {
		return []string{}
	}

	if !isLeftEmpty && len(rightLines) == 0 {
		return it.List()
	}

	if isLeftEmpty && len(rightLines) > 0 {
		return rightLines
	}

	diffLines := make(
		[]string,
		0,
		it.Length()+len(rightLines),
	)

	for _, rightItem := range rightLines {
		_, has := it.items[rightItem]
		isMissing := !has

		if isMissing {
			diffLines = append(diffLines, rightItem)
		}
	}

	rightHashset := converters.StringsTo.Hashset(
		rightLines,
	)

	for leftItem := range it.items {
		_, has := rightHashset[leftItem]
		isMissing := !has

		if isMissing {
			diffLines = append(diffLines, leftItem)
		}
	}

	return diffLines
}

func (it *Hashset) DistinctDiffHashset(
	rightHashset *Hashset,
) map[string]bool {
	return it.DistinctDiffLines(
		rightHashset.Lines()...,
	)
}

func (it *Hashset) DistinctDiffLines(
	rightLines ...string,
) map[string]bool {
	isLeftEmpty := it.IsEmpty()

	if isLeftEmpty && len(rightLines) == 0 {
		return map[string]bool{}
	}

	isLeftNotEmpty := !isLeftEmpty

	if isLeftNotEmpty && len(rightLines) == 0 {
		return it.Items()
	}

	if isLeftEmpty && len(rightLines) > 0 {
		return converters.StringsTo.Hashset(rightLines)
	}

	diffMap := make(
		map[string]bool,
		it.Length()+len(rightLines),
	)

	for _, rightItem := range rightLines {
		_, has := it.items[rightItem]
		isMissing := !has

		if isMissing {
			diffMap[rightItem] = true
		}
	}

	rightHashset := converters.StringsTo.Hashset(
		rightLines,
	)

	for leftItem := range it.items {
		_, has := rightHashset[leftItem]
		isMissing := !has

		if isMissing {
			diffMap[leftItem] = true
		}
	}

	return diffMap
}

func (it *Hashset) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *Hashset) Deserialize(toPtr any) (parsingErr error) {
	return it.JsonPtr().Deserialize(toPtr)
}

func (it *Hashset) WrapDoubleQuote() *Hashset {
	return it.Transpile(StringUtils.WrapDouble)
}

func (it *Hashset) WrapDoubleQuoteIfMissing() *Hashset {
	return it.Transpile(StringUtils.WrapDoubleIfMissing)
}

func (it *Hashset) WrapSingleQuote() *Hashset {
	return it.Transpile(StringUtils.WrapSingle)
}

func (it *Hashset) WrapSingleQuoteIfMissing() *Hashset {
	return it.Transpile(StringUtils.WrapSingleIfMissing)
}

// Transpile applies fmtFunc to each key and returns a new Hashset.
// Fix: build new map to avoid adding keys while iterating over old map.
// See issues/corestrtests-hashset-transpile-mutation.md
func (it *Hashset) Transpile(
	fmtFunc func(s string) string,
) *Hashset {
	if it.IsEmpty() {
		return Empty.Hashset()
	}

	newItems := make(map[string]bool, len(it.items))
	for k, v := range it.items {
		newItems[fmtFunc(k)] = v
	}

	it.items = newItems

	return it
}

func (it *Hashset) JoinLine() string {
	return it.Join("\n")
}
