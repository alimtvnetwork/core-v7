package corestr

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type Hashmap struct {
	items         *map[string]string
	hasMapUpdated bool
	cachedList    *[]string
	length        int
	isEmptySet    bool
	sync.Mutex
}

func (hashmap *Hashmap) IsEmpty() bool {
	if hashmap.hasMapUpdated {
		hashmap.isEmptySet = hashmap.items == nil ||
			*hashmap.items == nil ||
			len(*hashmap.items) == 0
	}

	return hashmap.isEmptySet
}

func (hashmap *Hashmap) HasItems() bool {
	return !hashmap.IsEmpty()
}

func (hashmap *Hashmap) Collection() *Collection {
	return NewCollectionUsingStrings(hashmap.ValuesListPtr(), false)
}

func (hashmap *Hashmap) IsEmptyLock() bool {
	hashmap.Lock()
	defer hashmap.Unlock()

	return hashmap.IsEmpty()
}

func (hashmap *Hashmap) AddOrUpdatePtr(
	key, val *string,
) *Hashmap {
	(*hashmap.items)[*key] = *val
	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateWithWgLock(
	key, val string,
	group *sync.WaitGroup,
) *Hashmap {
	hashmap.Lock()

	(*hashmap.items)[key] = val
	hashmap.hasMapUpdated = true

	hashmap.Unlock()
	group.Done()

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdatePtrLock(
	key, val *string,
) *Hashmap {
	hashmap.Lock()

	(*hashmap.items)[*key] = *val
	hashmap.hasMapUpdated = true

	hashmap.Unlock()

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateKeyStrValInt(
	key string,
	val int,
) *Hashmap {
	(*hashmap.items)[key] = fmt.Sprintf("%d", val)
	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateKeyStrValFloat(
	key string,
	val float32,
) *Hashmap {
	(*hashmap.items)[key] = fmt.Sprintf("%f", val)
	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateKeyStrValFloat64(
	key string, val float64,
) *Hashmap {
	(*hashmap.items)[key] = fmt.Sprintf("%f", val)
	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateKeyStrValAny(
	key string,
	val interface{},
) *Hashmap {
	(*hashmap.items)[key] = fmt.Sprintf(constants.SprintValueFormat, val)
	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateKeyValueAny(
	pair KeyAnyValuePair,
) *Hashmap {
	(*hashmap.items)[pair.Key] = pair.ValueString()
	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateKeyVal(
	keyVal KeyValuePair,
) *Hashmap {
	(*hashmap.items)[keyVal.Key] = keyVal.Value
	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdate(key, val string) *Hashmap {
	(*hashmap.items)[key] = val
	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateStringsPtrWgLock(
	keys, values *[]string, wg *sync.WaitGroup,
) *Hashmap {
	if keys == nil || values == nil {
		return hashmap
	}

	hashmap.Lock()
	for i, key := range *keys {
		(*hashmap.items)[key] = (*values)[i]
	}

	hashmap.hasMapUpdated = true
	hashmap.Unlock()
	wg.Done()

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateStringsPtr(
	keys, values *[]string,
) *Hashmap {
	if keys == nil || values == nil {
		return hashmap
	}

	for i, key := range *keys {
		(*hashmap.items)[key] = (*values)[i]
	}

	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateStringsPtrLock(
	keys, values *[]string,
) *Hashmap {
	if keys == nil || values == nil {
		return hashmap
	}

	hashmap.Lock()
	for i, key := range *keys {
		(*hashmap.items)[key] = (*values)[i]
	}

	hashmap.hasMapUpdated = true
	hashmap.Unlock()

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateMap(
	itemsMap *map[string]string,
) *Hashmap {
	if itemsMap == nil {
		return hashmap
	}

	for key, val := range *itemsMap {
		(*hashmap.items)[key] = val
	}

	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddsOrUpdates(
	KeyValuePair ...KeyValuePair,
) *Hashmap {
	if KeyValuePair == nil {
		return hashmap
	}

	for _, keyVal := range KeyValuePair {
		(*hashmap.items)[keyVal.Key] = keyVal.Value
	}

	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateKeyAnyValsPtr(
	pairs *[]KeyAnyValuePair,
) *Hashmap {
	if pairs == nil || *pairs == nil {
		return hashmap
	}

	for _, pair := range *pairs {
		(*hashmap.items)[pair.Key] = pair.ValueString()
	}

	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateKeyValsPtr(
	pairs *[]KeyValuePair,
) *Hashmap {
	if pairs == nil || *pairs == nil {
		return hashmap
	}

	for _, pair := range *pairs {
		(*hashmap.items)[pair.Key] = pair.Value
	}

	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateCollection(
	keys, values *Collection,
) *Hashmap {
	if (keys == nil || keys.IsEmpty()) || (values == nil || values.IsEmpty()) {
		return hashmap
	}

	for i, element := range *keys.items {
		(*hashmap.items)[element] = (*values.items)[i]
	}

	hashmap.hasMapUpdated = true

	return hashmap
}

// Keep result from filter.
func (hashmap *Hashmap) AddsOrUpdatesAnyUsingFilter(
	filter IsKeyAnyValueFilter,
	pairs ...KeyAnyValuePair,
) *Hashmap {
	if pairs == nil {
		return hashmap
	}

	for _, pair := range pairs {
		result, isKeep, isBreak := filter(pair)

		if isKeep {
			(*hashmap.items)[pair.Key] = result
			hashmap.hasMapUpdated = true
		}

		if isBreak {
			return hashmap
		}
	}

	return hashmap
}

// Keep result from filter.
func (hashmap *Hashmap) AddsOrUpdatesAnyUsingFilterLock(
	filter IsKeyAnyValueFilter,
	pairs ...KeyAnyValuePair,
) *Hashmap {
	if pairs == nil {
		return hashmap
	}

	for _, pair := range pairs {
		result, isKeep, isBreak := filter(pair)

		if isKeep {
			hashmap.Lock()
			(*hashmap.items)[pair.Key] = result
			hashmap.Unlock()

			hashmap.hasMapUpdated = true
		}

		if isBreak {
			return hashmap
		}
	}

	return hashmap
}

func (hashmap *Hashmap) AddsOrUpdatesUsingFilter(
	filter IsKeyValueFilter,
	pairs ...KeyValuePair,
) *Hashmap {
	if pairs == nil {
		return hashmap
	}

	for _, pair := range pairs {
		result, isKeep, isBreak := filter(pair)

		if isKeep {
			(*hashmap.items)[pair.Key] = result
			hashmap.hasMapUpdated = true
		}

		if isBreak {
			return hashmap
		}
	}

	return hashmap
}

func (hashmap *Hashmap) AddOrUpdateLock(key, value string) *Hashmap {
	hashmap.Lock()
	defer hashmap.Unlock()

	(*hashmap.items)[key] = value
	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) Has(key string) bool {
	_, isFound := (*hashmap.items)[key]

	return isFound
}

func (hashmap *Hashmap) HasLock(key string) bool {
	hashmap.Lock()
	_, isFound := (*hashmap.items)[key]
	hashmap.Unlock()

	return isFound
}

func (hashmap *Hashmap) HasAllStringsPtr(keys *[]string) bool {
	for _, key := range *keys {
		_, isFound := (*hashmap.items)[key]

		if !isFound {
			// not found
			return false
		}
	}

	// all found.
	return true
}

// return false on items is nil or empty.
func (hashmap *Hashmap) HasAllCollectionItems(
	collection *Collection,
) bool {
	if collection == nil || collection.IsEmpty() {
		return false
	}

	return hashmap.HasAllStringsPtr(collection.items)
}

func (hashmap *Hashmap) HasAll(keys ...string) bool {
	for _, key := range keys {
		_, isFound := (*hashmap.items)[key]

		if !isFound {
			// not found
			return false
		}
	}

	// all found.
	return true
}

func (hashmap *Hashmap) HasAny(keys ...string) bool {
	for _, key := range keys {
		_, isFound := (*hashmap.items)[key]

		if isFound {
			// any found
			return true
		}
	}

	// all not found.
	return false
}

func (hashmap *Hashmap) HasWithLock(key string) bool {
	hashmap.Lock()
	defer hashmap.Unlock()

	_, isFound := (*hashmap.items)[key]

	return isFound
}

// must return slice.
func (hashmap *Hashmap) GetKeysFilteredItems(
	filter IsStringFilter,
) *[]string {
	if hashmap.IsEmpty() {
		return &([]string{})
	}

	filteredList := make(
		[]string,
		0,
		hashmap.Length())

	i := 0
	for key := range *hashmap.items {
		result, isKeep, isBreak :=
			filter(key, i)

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

// must return items.
func (hashmap *Hashmap) GetKeysFilteredCollection(
	filter IsStringFilter,
) *Collection {
	if hashmap.IsEmpty() {
		return EmptyCollection()
	}

	filteredList := make(
		[]string,
		0,
		hashmap.Length())

	i := 0
	for key := range *hashmap.items {
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
				&filteredList, false)
		}
	}

	return NewCollectionUsingStrings(
		&filteredList, false)
}

func (hashmap *Hashmap) Items() *map[string]string {
	return hashmap.items
}

//goland:noinspection GoLinterLocal
func (hashmap *Hashmap) ItemsCopyLock() *map[string]string {
	hashmap.Lock()

	copiedItemsMap := &(*hashmap.items)

	hashmap.Unlock()

	return copiedItemsMap
}

func (hashmap *Hashmap) ValuesCollection() *Collection {
	return NewCollectionUsingStrings(
		hashmap.ValuesListPtr(), false)
}

func (hashmap *Hashmap) ValuesHashset() *Hashset {
	return NewHashsetUsingStrings(
		hashmap.ValuesListPtr(),
		0,
		false)
}

func (hashmap *Hashmap) ValuesCollectionLock() *Collection {
	return NewCollectionUsingStrings(
		hashmap.ValuesListCopyPtrLock(), false)
}

func (hashmap *Hashmap) ValuesHashsetLock() *Hashset {
	return NewHashsetUsingStrings(
		hashmap.ValuesListCopyPtrLock(),
		0,
		false)
}

func (hashmap *Hashmap) ValuesList() []string {
	return *hashmap.ValuesListPtr()
}

func (hashmap *Hashmap) ValuesListPtr() *[]string {
	if hashmap.hasMapUpdated || hashmap.cachedList == nil {
		hashmap.setCached()
	}

	return hashmap.cachedList
}

func (hashmap *Hashmap) KeysValuesCollection() (
	keys, values *Collection,
) {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		keys = NewCollectionUsingStrings(hashmap.Keys(), false)

		wg.Done()
	}()

	go func() {
		values = NewCollectionUsingStrings(hashmap.ValuesListPtr(), false)

		wg.Done()
	}()

	wg.Wait()

	return keys, values
}

func (hashmap *Hashmap) KeysValuesList() (
	keys, values *[]string,
) {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		keys = hashmap.Keys()
		wg.Done()
	}()

	go func() {
		values = hashmap.ValuesListPtr()
		wg.Done()
	}()

	wg.Wait()

	return keys, values
}

func (hashmap *Hashmap) KeysValuePairs() *[]KeyValuePair {
	pairs := make([]KeyValuePair, hashmap.Length())

	i := 0
	for k, v := range *hashmap.items {
		pairs[i] = KeyValuePair{
			Key:   k,
			Value: v,
		}

		i++
	}

	return &pairs
}

func (hashmap *Hashmap) KeysValuesListLock() (
	keys, values *[]string,
) {
	hashmap.Lock()
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		keys = hashmap.Keys()
		wg.Done()
	}()
	go func() {
		values = hashmap.ValuesListPtr()
		wg.Done()
	}()

	wg.Wait()
	hashmap.Unlock()

	return keys, values
}

func (hashmap *Hashmap) Keys() *[]string {
	length := len(*hashmap.items)
	keys := make([]string, length)

	if length == 0 {
		return &keys
	}

	i := 0
	for k := range *hashmap.items {
		keys[i] = k
		i++
	}

	return &keys
}

func (hashmap *Hashmap) KeysCollection() *Collection {
	return NewCollectionUsingStrings(
		hashmap.Keys(),
		false)
}

func (hashmap *Hashmap) KeysLock() *[]string {
	length := hashmap.LengthLock()
	keys := make([]string, length)

	if length == 0 {
		return &keys
	}

	i := 0
	hashmap.Lock()
	for k := range *hashmap.items {
		keys[i] = k
		i++
	}

	hashmap.Unlock()

	return &keys
}

// a slice must returned
func (hashmap *Hashmap) ValuesListCopyPtrLock() *[]string {
	hashmap.Lock()
	defer hashmap.Unlock()

	return &(*hashmap.ValuesListPtr())
}

func (hashmap *Hashmap) setCached() {
	length := hashmap.Length()
	list := make([]string, length)

	if length == 0 {
		hashmap.cachedList = &list
		hashmap.hasMapUpdated = false

		return
	}

	i := 0

	for _, val := range *hashmap.items {
		list[i] = val
		i++
	}

	hashmap.hasMapUpdated = false
	hashmap.cachedList = &list
}

// Create a new items with all lower strings
func (hashmap *Hashmap) ValuesToLower() *Hashmap {
	newMap := make(map[string]string, hashmap.Length())

	var toLower string
	for key, value := range *hashmap.items {
		toLower = strings.ToLower(key)
		newMap[toLower] = value
	}

	return NewHashmapUsingMap(
		&newMap,
		0,
		false)
}

func (hashmap *Hashmap) Length() int {
	if hashmap.hasMapUpdated || hashmap.length < 0 {
		if hashmap.items == nil || *hashmap.items == nil {
			hashmap.length = 0

			return hashmap.length
		}

		hashmap.length = len(*hashmap.items)
	}

	return hashmap.length
}

func (hashmap *Hashmap) LengthLock() int {
	hashmap.Lock()
	defer hashmap.Unlock()

	return hashmap.Length()
}

//goland:noinspection GoLinterLocal,GoVetCopyLock
func (hashmap *Hashmap) IsEquals(another Hashmap) bool { //nolint:govet
	return hashmap.IsEqualsPtr(&another)
}

func (hashmap *Hashmap) IsEqualsPtrLock(another *Hashmap) bool {
	hashmap.Lock()
	defer hashmap.Unlock()

	return hashmap.IsEqualsPtr(another)
}

func (hashmap *Hashmap) IsEqualsPtr(another *Hashmap) bool {
	if hashmap == nil && another == nil {
		return true
	}

	if hashmap == nil || another == nil {
		return false
	}

	if hashmap == another {
		// ptr same
		return true
	}

	if hashmap.IsEmpty() && another.IsEmpty() {
		return true
	}

	if hashmap.IsEmpty() || another.IsEmpty() {
		return false
	}

	leftLength := hashmap.Length()
	rightLength := another.Length()

	if leftLength != rightLength {
		return false
	}

	for key, value := range *hashmap.items {
		result, has := (*another.items)[key]

		if !has || !(result != value) {
			return false
		}
	}

	return true
}

func (hashmap *Hashmap) Remove(key string) *Hashmap {
	delete(*hashmap.items, key)
	hashmap.hasMapUpdated = true

	return hashmap
}

func (hashmap *Hashmap) RemoveWithLock(key string) *Hashmap {
	hashmap.Lock()
	defer hashmap.Unlock()

	hashmap.Remove(key)

	return hashmap
}

func (hashmap *Hashmap) String() string {
	if hashmap.IsEmpty() {
		return commonJoiner + NoElements
	}

	return commonJoiner +
		strings.Join(
			hashmap.ValuesList(),
			commonJoiner)
}

func (hashmap *Hashmap) StringLock() string {
	if hashmap.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	hashmap.Lock()
	defer hashmap.Unlock()

	return commonJoiner +
		strings.Join(
			*hashmap.ValuesListPtr(),
			commonJoiner)
}

// Get all Collection except the mentioned ones.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashmap
// Set B = anotherHashset given in parameters.
func (hashmap *Hashmap) GetValuesExceptKeysInHashset(
	anotherHashset *Hashset,
) *[]string {
	if anotherHashset == nil || anotherHashset.IsEmpty() {
		return hashmap.ValuesListPtr()
	}

	finalList := make(
		[]string,
		0,
		hashmap.Length())

	for key, value := range *hashmap.items {
		if anotherHashset.Has(key) {
			continue
		}

		finalList = append(
			finalList,
			value)
	}

	return &finalList
}

// Get all items except the mentioned ones.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashmap
// Set B = items given in parameters.
func (hashmap *Hashmap) GetValuesKeysExcept(
	items *[]string,
) *[]string {
	if items == nil {
		return hashmap.ValuesListPtr()
	}

	newCollection := NewHashsetUsingStrings(
		items,
		0,
		false)

	return hashmap.GetValuesExceptKeysInHashset(
		newCollection)
}

// Get all Hashmap items except the mentioned ones in collection.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashmap
// Set B = collection given in parameters.
func (hashmap *Hashmap) GetAllExceptCollection(
	collection *Collection,
) *[]string {
	if collection == nil {
		return hashmap.ValuesListPtr()
	}

	return hashmap.GetValuesExceptKeysInHashset(
		collection.HashsetAsIs())
}

// Get all items except the mentioned ones in collectionPtr.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashmap
// Set B = collectionPtr given in parameters.
func (hashmap *Hashmap) GetAllExceptCollectionPtr(
	collectionPtr *CollectionPtr,
) *[]string {
	if collectionPtr == nil {
		return hashmap.ValuesListPtr()
	}

	return hashmap.GetValuesExceptKeysInHashset(
		collectionPtr.HashsetAsIs())
}

// Joins values
func (hashmap *Hashmap) Join(
	separator string,
) string {
	return strings.Join(*hashmap.ValuesListPtr(), separator)
}

// Joins Keys
func (hashmap *Hashmap) JoinKeys(
	separator string,
) string {
	return strings.Join(*hashmap.Keys(), separator)
}

func (hashmap *Hashmap) JsonModel() *HashmapDataModel {
	return NewHashmapsDataModelUsing(hashmap)
}

func (hashmap *Hashmap) JsonModelAny() interface{} {
	return hashmap.JsonModel()
}

func (hashmap *Hashmap) MarshalJSON() ([]byte, error) {
	return json.Marshal(hashmap.JsonModel())
}

func (hashmap *Hashmap) UnmarshalJSON(data []byte) error {
	var dataModel HashmapDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		hashmap.items = dataModel.Items
		hashmap.length = len(*hashmap.items)
		hashmap.hasMapUpdated = false
		hashmap.isEmptySet = hashmap.length == 0
		hashmap.cachedList = nil
	}

	return err
}

func (hashmap *Hashmap) Json() *corejson.Result {
	if hashmap.IsEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(hashmap)

	return corejson.NewPtr(jsonBytes, err)
}

// It will not update the self but creates a new one.
func (hashmap *Hashmap) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Hashmap, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return EmptyHashmap(), nil
	}

	err := json.Unmarshal(*jsonResult.Bytes, &hashmap)

	if err != nil {
		return EmptyHashmap(), err
	}

	return hashmap, nil
}

// Panic if error
func (hashmap *Hashmap) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Hashmap {
	hashSet, err := hashmap.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}

func (hashmap *Hashmap) AsJsoner() *corejson.Jsoner {
	var jsoner corejson.Jsoner = hashmap

	return &jsoner
}

func (hashmap *Hashmap) JsonParseSelfInject(jsonResult *corejson.Result) {
	hashmap.ParseInjectUsingJsonMust(jsonResult)
}

func (hashmap *Hashmap) AsJsonParseSelfInjector() *corejson.ParseSelfInjector {
	var jsonMarshaller corejson.ParseSelfInjector = hashmap

	return &jsonMarshaller
}

func (hashmap *Hashmap) AsJsonMarshaller() *corejson.JsonMarshaller {
	var jsonMarshaller corejson.JsonMarshaller = hashmap

	return &jsonMarshaller
}
