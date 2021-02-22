package corestr

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type Hashset struct {
	items         *map[string]bool
	hasMapUpdated bool
	cachedList    *[]string
	length        int
	isEmptySet    bool
	sync.Mutex
}

func (hashset *Hashset) IsEmpty() bool {
	if hashset.hasMapUpdated {
		hashset.isEmptySet = hashset.items == nil ||
			*hashset.items == nil ||
			len(*hashset.items) == 0
	}

	return hashset.isEmptySet
}

// Changing capacity creates new map and points to it.
// There is memory copy and loop is performed.
func (hashset *Hashset) AddCapacitiesLock(
	capacities ...int,
) *Hashset {
	length := hashset.LengthLock()

	if len(capacities) == 0 {
		return hashset
	}

	for _, capacity := range capacities {
		length += capacity
	}

	return hashset.ResizeLock(length)
}

// Changing capacity creates new map and points to it.
// There is memory copy and loop is performed.
func (hashset *Hashset) AddCapacities(
	capacities ...int,
) *Hashset {
	length := hashset.Length()

	if len(capacities) == 0 {
		return hashset
	}

	for _, capacity := range capacities {
		length += capacity
	}

	return hashset.Resize(length)
}

// Changing capacity creates new map and points to it.
// There is memory copy and loop is performed.
func (hashset *Hashset) Resize(capacity int) *Hashset {
	length := hashset.Length()

	if length > capacity {
		return hashset
	}

	newItemsMap := make(map[string]bool, capacity)

	for val := range *hashset.items {
		newItemsMap[val] = true
	}

	hashset.items = &newItemsMap

	return hashset
}

// Changing capacity creates new map and points to it.
// There is memory copy and loop is performed.
func (hashset *Hashset) ResizeLock(capacity int) *Hashset {
	length := hashset.LengthLock()

	if length > capacity {
		return hashset
	}

	newItemsMap := make(map[string]bool, capacity)

	for val := range *hashset.items {
		newItemsMap[val] = true
	}

	hashset.Lock()
	hashset.items = &newItemsMap
	hashset.Unlock()

	return hashset
}

func (hashset *Hashset) Collection() *Collection {
	return NewCollectionUsingStrings(hashset.ListPtr(), false)
}

func (hashset *Hashset) IsEmptyLock() bool {
	hashset.Lock()
	defer hashset.Unlock()

	return hashset.IsEmpty()
}

func (hashset *Hashset) AddPtr(key *string) *Hashset {
	(*hashset.items)[*key] = true
	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) AddWithWgLock(key string, group *sync.WaitGroup) *Hashset {
	hashset.Lock()
	(*hashset.items)[key] = true
	hashset.Unlock()

	hashset.hasMapUpdated = true

	group.Done()

	return hashset
}

func (hashset *Hashset) AddPtrLock(key *string) *Hashset {
	hashset.Lock()

	(*hashset.items)[*key] = true
	hashset.Unlock()

	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) Add(key string) *Hashset {
	(*hashset.items)[key] = true
	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) AddStringsPtrWgLock(keys *[]string, wg *sync.WaitGroup) *Hashset {
	if keys == nil {
		return hashset
	}

	length := len(*keys)

	if length > hashset.length || length > constants.ArbitraryCapacity100 {
		hashset.AddCapacitiesLock(length*2, constants.ArbitraryCapacity100)
	}

	hashset.Lock()
	for _, key := range *keys {
		(*hashset.items)[key] = true
	}

	hashset.Unlock()
	wg.Done()

	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) AddHashsetItems(
	hashsetAdd *Hashset,
) *Hashset {
	if hashsetAdd == nil {
		return hashset
	}

	length := hashsetAdd.Length()

	if length > hashset.length || length > constants.ArbitraryCapacity100 {
		hashset.AddCapacities(length*2, constants.ArbitraryCapacity100)
	}

	for k := range *hashsetAdd.items {
		(*hashset.items)[k] = true
	}

	hashset.hasMapUpdated = true

	return hashset
}

// only add if the value is true
func (hashset *Hashset) AddItemsMap(
	itemsMap *map[string]bool,
) *Hashset {
	if itemsMap == nil {
		return hashset
	}

	length := len(*itemsMap)

	if length > hashset.length || length > constants.ArbitraryCapacity100 {
		hashset.AddCapacities(length*2, constants.ArbitraryCapacity100)
	}

	for k, isEnabled := range *itemsMap {
		if !isEnabled {
			continue
		}

		(*hashset.items)[k] = true
	}

	hashset.hasMapUpdated = true

	return hashset
}

// only add if the value is true
// Assume that wg already enqueued the job as wg.Add(...) done already.
func (hashset *Hashset) AddItemsMapWgLock(
	itemsMap *map[string]bool,
	wg *sync.WaitGroup,
) *Hashset {
	if itemsMap == nil {
		return hashset
	}

	length := len(*itemsMap)

	if length > hashset.length || length > constants.ArbitraryCapacity100 {
		hashset.AddCapacitiesLock(length*2, constants.ArbitraryCapacity100)
	}

	for k, isEnabled := range *itemsMap {
		if !isEnabled {
			continue
		}

		hashset.Lock()
		(*hashset.items)[k] = true
		hashset.Unlock()
	}

	wg.Done()

	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) AddHashsetWgLock(
	hashsetAdd *Hashset,
	wg *sync.WaitGroup,
) *Hashset {
	if hashsetAdd == nil {
		return hashset
	}

	length := hashsetAdd.LengthLock()

	if length > hashset.length || length > constants.ArbitraryCapacity100 {
		hashset.AddCapacitiesLock(length*2, constants.ArbitraryCapacity100)
	}

	hashset.Lock()
	for k := range *hashsetAdd.items {
		(*hashset.items)[k] = true
	}

	hashset.Unlock()
	wg.Done()

	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) AddStringsPtr(keys *[]string) *Hashset {
	if keys == nil {
		return hashset
	}

	for _, key := range *keys {
		(*hashset.items)[key] = true
	}

	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) AddStringsPtrLock(keys *[]string) *Hashset {
	if keys == nil {
		return hashset
	}

	hashset.Lock()
	for _, key := range *keys {
		(*hashset.items)[key] = true
	}

	hashset.Unlock()

	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) Adds(keys ...string) *Hashset {
	if keys == nil {
		return hashset
	}

	for _, key := range keys {
		(*hashset.items)[key] = true
	}

	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) AddCollection(
	collection *Collection,
) *Hashset {
	if collection == nil || collection.IsEmpty() {
		return hashset
	}

	for _, element := range *collection.items {
		(*hashset.items)[element] = true
	}

	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) AddCollections(
	collections ...*Collection,
) *Hashset {
	if collections == nil {
		return hashset
	}

	for _, collection := range collections {
		if collection == nil || collection.IsEmpty() {
			continue
		}

		for _, element := range *collection.items {
			(*hashset.items)[element] = true
		}
	}

	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) AddsAnyUsingFilter(
	filter IsStringFilter,
	anys ...interface{},
) *Hashset {
	if anys == nil {
		return hashset
	}

	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)
		result, isKeep, isBreak := filter(anyStr)

		if isKeep {
			(*hashset.items)[result] = true
			hashset.hasMapUpdated = true
		}

		if isBreak {
			return hashset
		}
	}

	return hashset
}

func (hashset *Hashset) AddsAnyUsingFilterLock(
	filter IsStringFilter,
	anys ...interface{},
) *Hashset {
	if anys == nil {
		return hashset
	}

	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(
			constants.SprintValueFormat,
			any)

		result, isKeep, isBreak := filter(anyStr)

		if isKeep {
			hashset.Lock()
			(*hashset.items)[result] = true
			hashset.Unlock()

			hashset.hasMapUpdated = true
		}

		if isBreak {
			return hashset
		}
	}

	return hashset
}

func (hashset *Hashset) AddsUsingFilter(
	filter IsStringFilter,
	keys ...string,
) *Hashset {
	if keys == nil {
		return hashset
	}

	for _, key := range keys {
		result, isKeep, isBreak := filter(key)

		if isKeep {
			(*hashset.items)[result] = true
			hashset.hasMapUpdated = true
		}

		if isBreak {
			return hashset
		}
	}

	return hashset
}

func (hashset *Hashset) AddLock(key string) *Hashset {
	hashset.Lock()
	defer hashset.Unlock()

	(*hashset.items)[key] = true
	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) Has(key string) bool {
	isSet, isFound := (*hashset.items)[key]

	return isFound && isSet
}

func (hashset *Hashset) HasLock(key string) bool {
	hashset.Lock()
	isSet, isFound := (*hashset.items)[key]
	hashset.Unlock()

	return isFound && isSet
}

func (hashset *Hashset) HasAllStringsPtr(keys *[]string) bool {
	for _, key := range *keys {
		isSet, isFound := (*hashset.items)[key]

		if !(isFound && isSet) {
			// not found
			return false
		}
	}

	// all found.
	return true
}

// return false on items is nil or empty.
func (hashset *Hashset) HasAllCollectionItems(collection *Collection) bool {
	if collection == nil || collection.IsEmpty() {
		return false
	}

	return hashset.HasAllStringsPtr(collection.items)
}

func (hashset *Hashset) HasAll(keys ...string) bool {
	for _, key := range keys {
		isSet, isFound := (*hashset.items)[key]

		if !(isFound && isSet) {
			// not found
			return false
		}
	}

	// all found.
	return true
}

func (hashset *Hashset) HasAny(keys ...string) bool {
	for _, key := range keys {
		isSet, isFound := (*hashset.items)[key]

		if isFound && isSet {
			// any found
			return true
		}
	}

	// all not found.
	return false
}

func (hashset *Hashset) HasWithLock(key string) bool {
	hashset.Lock()
	defer hashset.Unlock()

	isSet, isFound := (*hashset.items)[key]

	return isFound && isSet
}

// must return slice.
func (hashset *Hashset) GetFilteredItems(
	filter IsStringFilter,
) *[]string {
	if hashset.IsEmpty() {
		return &([]string{})
	}

	filteredList := make(
		[]string,
		0,
		hashset.Length())

	for key := range *hashset.items {
		result, isKeep, isBreak := filter(key)

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

// must return items.
func (hashset *Hashset) GetFilteredCollection(
	filter IsStringFilter,
) *Collection {
	if hashset.IsEmpty() {
		return EmptyCollection()
	}

	filteredList := make(
		[]string,
		0,
		hashset.Length())

	for key := range *hashset.items {
		result, isKeep, isBreak := filter(key)

		if !isKeep {
			continue
		}

		filteredList = append(
			filteredList,
			result)

		if isBreak {
			return NewCollectionUsingStrings(
				&filteredList, false)
		}
	}

	return NewCollectionUsingStrings(
		&filteredList, false)
}

// Get all hashset items except the mentioned ones in anotherHashset.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashset
// Set B = anotherHashset given in parameters.
func (hashset *Hashset) GetAllExceptHashset(anotherHashset *Hashset) *[]string {
	if anotherHashset == nil || anotherHashset.IsEmpty() {
		return hashset.ListPtr()
	}

	finalList := make(
		[]string,
		0,
		hashset.Length())

	for item := range *hashset.items {
		if anotherHashset.Has(item) {
			continue
		}

		finalList = append(
			finalList,
			item)
	}

	return &finalList
}

// Get all hashset items except the mentioned ones in items.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashset
// Set B = items given in parameters.
func (hashset *Hashset) GetAllExcept(items *[]string) *[]string {
	if items == nil {
		return hashset.ListPtr()
	}

	newHashset := NewHashsetUsingStrings(
		items,
		0,
		false)

	return hashset.GetAllExceptHashset(
		newHashset)
}

// Get all hashset items except the mentioned ones in collection.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashset
// Set B = collection given in parameters.
func (hashset *Hashset) GetAllExceptCollection(collection *Collection) *[]string {
	if collection == nil {
		return hashset.ListPtr()
	}

	return hashset.GetAllExceptHashset(
		collection.HashsetAsIs())
}

// Get all hashset items except the mentioned ones in collectionPtr.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashset
// Set B = collectionPtr given in parameters.
func (hashset *Hashset) GetAllExceptCollectionPtr(
	collectionPtr *CollectionPtr,
) *[]string {
	if collectionPtr == nil {
		return hashset.ListPtr()
	}

	return hashset.GetAllExceptHashset(
		collectionPtr.HashsetAsIs())
}

func (hashset *Hashset) Items() *map[string]bool {
	return hashset.items
}

func (hashset *Hashset) List() []string {
	return *hashset.ListPtr()
}

func (hashset *Hashset) ListPtrSortedAsc() *[]string {
	list := hashset.ListPtr()
	sort.Strings(*list)

	return list
}

func (hashset *Hashset) ListPtr() *[]string {
	if hashset.hasMapUpdated || hashset.cachedList == nil {
		hashset.setCached()
	}

	return hashset.cachedList
}

// a slice must returned
func (hashset *Hashset) ListCopyPtrLock() *[]string {
	hashset.Lock()
	defer hashset.Unlock()
	cloned := *hashset.ListPtr()

	return &cloned
}

func (hashset *Hashset) setCached() {
	length := hashset.Length()
	list := make([]string, length)

	i := 0

	for key := range *hashset.items {
		list[i] = key
		i++
	}

	hashset.hasMapUpdated = false
	hashset.cachedList = &list
}

// Create a new items with all lower strings
func (hashset *Hashset) ToLowerSet() *Hashset {
	length := hashset.Length()
	newMap := make(map[string]bool, length)

	var toLower string
	for key, isEnabled := range *hashset.items {
		toLower = strings.ToLower(key)
		newMap[toLower] = isEnabled
	}

	return NewHashsetUsingMap(
		&newMap,
		length,
		false)
}

func (hashset *Hashset) Length() int {
	if hashset.hasMapUpdated {
		if hashset.items == nil || *hashset.items == nil {
			hashset.length = 0

			return hashset.length
		}

		hashset.length = len(*hashset.items)
	}

	return hashset.length
}

func (hashset *Hashset) LengthLock() int {
	hashset.Lock()
	defer hashset.Unlock()

	return hashset.Length()
}

//goland:noinspection GoVetCopyLock
func (hashset *Hashset) IsEquals(another Hashset) bool {
	return hashset.IsEqualsPtr(&another)
}

func (hashset *Hashset) IsEqualsPtrLock(another *Hashset) bool {
	hashset.Lock()
	defer hashset.Unlock()

	return hashset.IsEqualsPtr(another)
}

func (hashset *Hashset) IsEqualsPtr(another *Hashset) bool {
	if hashset == nil && another == nil {
		return true
	}

	if hashset == nil || another == nil {
		return false
	}

	if hashset == another {
		// ptr same
		return true
	}

	if hashset.IsEmpty() && another.IsEmpty() {
		return true
	}

	if hashset.IsEmpty() || another.IsEmpty() {
		return false
	}

	leftLength := hashset.Length()
	rightLength := another.Length()

	if leftLength != rightLength {
		return false
	}

	for key := range *hashset.items {
		isRes, has := (*another.items)[key]

		if !has || !isRes {
			return false
		}
	}

	return true
}

func (hashset *Hashset) Remove(key string) *Hashset {
	delete(*hashset.items, key)
	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) RemoveWithLock(key string) *Hashset {
	hashset.Lock()
	defer hashset.Unlock()

	hashset.Remove(key)

	return hashset
}

func (hashset *Hashset) String() string {
	if hashset.IsEmpty() {
		return commonJoiner + NoElements
	}

	return commonJoiner +
		strings.Join(
			hashset.List(),
			commonJoiner)
}

func (hashset *Hashset) StringLock() string {
	if hashset.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	hashset.Lock()
	defer hashset.Unlock()

	return commonJoiner +
		strings.Join(
			*hashset.ListPtr(),
			commonJoiner)
}

func (hashset *Hashset) Join(
	separator string,
) string {
	return strings.Join(*hashset.ListPtr(), separator)
}

//goland:noinspection GoLinterLocal
func (hashset *Hashset) JsonModel() *HashsetDataModel {
	return NewHashsetsDataModelUsing(hashset)
}

//goland:noinspection GoLinterLocal
func (hashset *Hashset) JsonModelAny() interface{} {
	return hashset.JsonModel()
}

func (hashset *Hashset) MarshalJSON() ([]byte, error) {
	return json.Marshal(hashset.JsonModel())
}

func (hashset *Hashset) UnmarshalJSON(data []byte) error {
	var dataModel HashsetDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		hashset.items = dataModel.Items
		hashset.length = -1
		hashset.hasMapUpdated = true
		hashset.isEmptySet = false
	}

	return err
}

//goland:noinspection GoLinterLocal
func (hashset *Hashset) Json() *corejson.Result {
	if hashset.IsEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(hashset)

	return corejson.NewPtr(jsonBytes, err)
}

// It will not update the self but creates a new one.
func (hashset *Hashset) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Hashset, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return EmptyHashset(), nil
	}

	err := json.Unmarshal(*jsonResult.Bytes, &hashset)

	if err != nil {
		return EmptyHashset(), err
	}

	return hashset, nil
}

// Panic if error
func (hashset *Hashset) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Hashset {
	hashSet, err := hashset.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}
