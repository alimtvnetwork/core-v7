package corestr

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
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

func (hashset *Hashset) Collection() *Collection {
	return NewCollectionUsingStrings(hashset.ListPtr())
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

	for k := range *hashsetAdd.items {
		(*hashset.items)[k] = true
	}

	hashset.hasMapUpdated = true

	return hashset
}

func (hashset *Hashset) AddHashsetWgLock(
	hashsetAdd *Hashset, wg *sync.WaitGroup,
) *Hashset {
	if hashsetAdd == nil {
		return hashset
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
		result, isKeep := filter(anyStr)

		if isKeep {
			(*hashset.items)[result] = true
			hashset.hasMapUpdated = true
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

		result, isKeep := filter(anyStr)

		if isKeep {
			hashset.Lock()
			(*hashset.items)[result] = true
			hashset.Unlock()

			hashset.hasMapUpdated = true
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
		result, isKeep := filter(key)

		if isKeep {
			(*hashset.items)[result] = true
			hashset.hasMapUpdated = true
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

	for key, _ := range *hashset.items {
		result, isKeep := filter(key)

		if !isKeep {
			continue
		}

		filteredList = append(
			filteredList,
			result)
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

	for key, _ := range *hashset.items {
		result, isKeep := filter(key)

		if !isKeep {
			continue
		}

		filteredList = append(
			filteredList,
			result)
	}

	return NewCollectionUsingStrings(
		&filteredList)
}

func (hashset *Hashset) Items() *map[string]bool {
	return hashset.items
}

func (hashset *Hashset) List() []string {
	return *hashset.ListPtr()
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

	return &(*hashset.ListPtr())
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
	newMap := make(map[string]bool, hashset.Length())

	var toLower string
	for key, isEnabled := range *hashset.items {
		toLower = strings.ToLower(key)
		newMap[toLower] = isEnabled
	}

	return NewHashsetUsingMap(&newMap)
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

func (hashset *Hashset) IsEquals(another Hashset) bool {
	return hashset.IsEqualsPtr(&another)
}

func (hashset *Hashset) IsEqualsPtrLock(another *Hashset) bool {
	if hashset == nil {
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

	for key, _ := range *hashset.items {
		isRes, has := (*another.items)[key]

		if !has || !isRes {
			return false
		}
	}

	return true
}

func (hashset *Hashset) IsEqualsPtr(another *Hashset) bool {
	if hashset == nil {
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

	for key, _ := range *hashset.items {
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

func (hashset *Hashset) JsonModel() *HashsetDataModel {
	return NewHashsetsDataModelUsing(hashset)
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

func (hashset *Hashset) Json() *JsonResult {
	if hashset.IsEmpty() {
		return EmptyJsonResultWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(hashset)

	return NewJsonResultPtr(jsonBytes, err)
}

// It will not update the self but creates a new one.
func (hashset *Hashset) ParseInjectUsingJson(
	jsonResult *JsonResult,
) (*Hashset, error) {
	if jsonResult == nil || jsonResult.IsBytesEmpty() {
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
	jsonResult *JsonResult,
) *Hashset {
	hashSet, err := hashset.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}
