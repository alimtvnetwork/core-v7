package corestr

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/errcore"
)

type Hashmap struct {
	hasMapUpdated bool
	isEmptySet    bool
	length        int
	items         map[string]string
	cachedList    []string
	sync.Mutex
}

func (it *Hashmap) IsEmpty() bool {
	if it == nil {
		return true
	}

	if it.hasMapUpdated {
		it.isEmptySet = len(it.items) == 0
	}

	return it.isEmptySet
}

func (it *Hashmap) HasItems() bool {
	return it != nil && !it.IsEmpty()
}

func (it *Hashmap) Collection() *Collection {
	return NewCollectionUsingStrings(false, it.ValuesList())
}

func (it *Hashmap) IsEmptyLock() bool {
	it.Lock()
	defer it.Unlock()

	return it.IsEmpty()
}

func (it *Hashmap) AddOrUpdatePtr(
	key, val *string,
) *Hashmap {
	it.items[*key] = *val
	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateWithWgLock(
	key, val string,
	group *sync.WaitGroup,
) *Hashmap {
	it.Lock()

	it.items[key] = val
	it.hasMapUpdated = true

	it.Unlock()
	group.Done()

	return it
}

func (it *Hashmap) AddOrUpdatePtrLock(
	key, val *string,
) *Hashmap {
	it.Lock()

	it.items[*key] = *val
	it.hasMapUpdated = true

	it.Unlock()

	return it
}

func (it *Hashmap) AddOrUpdateKeyStrValInt(
	key string,
	val int,
) *Hashmap {
	it.items[key] = strconv.Itoa(val)
	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateKeyStrValFloat(
	key string,
	val float32,
) *Hashmap {
	it.items[key] = fmt.Sprintf("%f", val)
	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateKeyStrValFloat64(
	key string, val float64,
) *Hashmap {
	it.items[key] = fmt.Sprintf("%f", val)
	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateKeyStrValAny(
	key string,
	val interface{},
) *Hashmap {
	it.items[key] = fmt.Sprintf(constants.SprintValueFormat, val)
	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateKeyValueAny(
	pair KeyAnyValuePair,
) *Hashmap {
	it.items[pair.Key] = pair.ValueString()
	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateKeyVal(
	keyVal KeyValuePair,
) *Hashmap {
	it.items[keyVal.Key] = keyVal.Value
	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdate(key, val string) *Hashmap {
	it.items[key] = val
	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateStringsPtrWgLock(
	keys, values *[]string, wg *sync.WaitGroup,
) *Hashmap {
	if keys == nil || values == nil {
		return it
	}

	it.Lock()
	for i, key := range *keys {
		it.items[key] = (*values)[i]
	}

	it.hasMapUpdated = true
	it.Unlock()
	wg.Done()

	return it
}

func (it *Hashmap) AddOrUpdateStringsPtr(
	keys, values *[]string,
) *Hashmap {
	if keys == nil || values == nil {
		return it
	}

	for i, key := range *keys {
		it.items[key] = (*values)[i]
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateStringsPtrLock(
	keys, values *[]string,
) *Hashmap {
	if keys == nil || values == nil {
		return it
	}

	it.Lock()
	for i, key := range *keys {
		it.items[key] = (*values)[i]
	}

	it.hasMapUpdated = true
	it.Unlock()

	return it
}

func (it *Hashmap) AddOrUpdateHashmap(
	hashmap2 *Hashmap,
) *Hashmap {
	if hashmap2 == nil {
		return it
	}

	for key, val := range hashmap2.items {
		it.items[key] = val
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateMap(
	itemsMap map[string]string,
) *Hashmap {
	if len(itemsMap) == 0 {
		return it
	}

	for key, val := range itemsMap {
		it.items[key] = val
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateMapPtr(
	itemsMap *map[string]string,
) *Hashmap {
	if itemsMap == nil || len(*itemsMap) == 0 {
		return it
	}

	for key, val := range *itemsMap {
		it.items[key] = val
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddsOrUpdates(
	KeyValuePair ...KeyValuePair,
) *Hashmap {
	if KeyValuePair == nil {
		return it
	}

	for _, keyVal := range KeyValuePair {
		it.items[keyVal.Key] = keyVal.Value
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateKeyAnyValsPtr(
	pairs *[]KeyAnyValuePair,
) *Hashmap {
	if pairs == nil || *pairs == nil {
		return it
	}

	for _, pair := range *pairs {
		it.items[pair.Key] = pair.ValueString()
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateKeyValsPtr(
	pairs *[]KeyValuePair,
) *Hashmap {
	if pairs == nil || *pairs == nil {
		return it
	}

	for _, pair := range *pairs {
		it.items[pair.Key] = pair.Value
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateCollection(
	keys, values *Collection,
) *Hashmap {
	if (keys == nil || keys.IsEmpty()) || (values == nil || values.IsEmpty()) {
		return it
	}

	for i, element := range keys.items {
		it.items[element] = values.items[i]
	}

	it.hasMapUpdated = true

	return it
}

// AddsOrUpdatesAnyUsingFilter Keep result from filter.
func (it *Hashmap) AddsOrUpdatesAnyUsingFilter(
	filter IsKeyAnyValueFilter,
	pairs ...KeyAnyValuePair,
) *Hashmap {
	if pairs == nil {
		return it
	}

	for _, pair := range pairs {
		result, isKeep, isBreak := filter(pair)

		if isKeep {
			it.items[pair.Key] = result
			it.hasMapUpdated = true
		}

		if isBreak {
			return it
		}
	}

	return it
}

// AddsOrUpdatesAnyUsingFilterLock Keep result from filter.
func (it *Hashmap) AddsOrUpdatesAnyUsingFilterLock(
	filter IsKeyAnyValueFilter,
	pairs ...KeyAnyValuePair,
) *Hashmap {
	if pairs == nil {
		return it
	}

	for _, pair := range pairs {
		result, isKeep, isBreak := filter(pair)

		if isKeep {
			it.Lock()
			it.items[pair.Key] = result
			it.Unlock()

			it.hasMapUpdated = true
		}

		if isBreak {
			return it
		}
	}

	return it
}

func (it *Hashmap) AddsOrUpdatesUsingFilter(
	filter IsKeyValueFilter,
	pairs ...KeyValuePair,
) *Hashmap {
	if pairs == nil {
		return it
	}

	for _, pair := range pairs {
		result, isKeep, isBreak := filter(pair)

		if isKeep {
			it.items[pair.Key] = result
			it.hasMapUpdated = true
		}

		if isBreak {
			return it
		}
	}

	return it
}

func (it *Hashmap) ConcatNew(
	isCloneOnEmptyAsWell bool,
	hashmaps ...*Hashmap,
) *Hashmap {
	if len(hashmaps) == 0 {
		return NewHashmapUsingMap(
			isCloneOnEmptyAsWell,
			constants.Zero,
			it.items,
		)
	}

	length := it.Length() + constants.Capacity2

	for _, h := range hashmaps {
		if h == nil {
			continue
		}

		length += h.length
	}

	newHashmap := NewHashmapUsingMap(
		true,
		length,
		it.items,
	)

	newHashmap.AddOrUpdateHashmap(it)

	for _, hashmap2 := range hashmaps {
		newHashmap.AddOrUpdateHashmap(
			hashmap2)
	}

	return newHashmap
}

func (it *Hashmap) ConcatNewUsingMaps(
	isCloneOnEmptyAsWell bool,
	hashmaps ...*map[string]string,
) *Hashmap {
	if len(hashmaps) == 0 {
		return NewHashmapUsingMap(
			isCloneOnEmptyAsWell,
			constants.Zero,
			it.items,
		)
	}

	length := it.Length() +
		constants.Capacity5
	for _, h := range hashmaps {
		if h == nil {
			continue
		}

		length += len(*h)
	}

	newHashmap := NewHashmapUsingMap(
		true,
		length,
		it.items,
	)

	newHashmap.AddOrUpdateHashmap(it)

	for _, hashmap2 := range hashmaps {
		newHashmap.AddOrUpdateMapPtr(
			hashmap2)
	}

	return newHashmap
}

func (it *Hashmap) AddOrUpdateLock(key, value string) *Hashmap {
	it.Lock()
	defer it.Unlock()

	it.items[key] = value
	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) Has(key string) bool {
	_, isFound := it.items[key]

	return isFound
}

func (it *Hashmap) IsKeyMissing(key string) bool {
	_, isFound := it.items[key]

	return !isFound
}

func (it *Hashmap) IsKeyMissingLock(key string) bool {
	it.Lock()
	_, isFound := it.items[key]
	it.Unlock()

	return !isFound
}

func (it *Hashmap) HasLock(key string) bool {
	it.Lock()
	_, isFound := it.items[key]
	it.Unlock()

	return isFound
}

func (it *Hashmap) HasAllStringsPtr(keys *[]string) bool {
	for _, key := range *keys {
		_, isFound := it.items[key]

		if !isFound {
			// not found
			return false
		}
	}

	// all found.
	return true
}

// HasAllCollectionItems return false on items is nil or empty.
func (it *Hashmap) HasAllCollectionItems(
	collection *Collection,
) bool {
	if collection == nil || collection.IsEmpty() {
		return false
	}

	return it.HasAllStringsPtr(collection.ListPtr())
}

func (it *Hashmap) HasAll(keys ...string) bool {
	for _, key := range keys {
		_, isFound := it.items[key]

		if !isFound {
			// not found
			return false
		}
	}

	// all found.
	return true
}

func (it *Hashmap) HasAny(keys ...string) bool {
	for _, key := range keys {
		_, isFound := it.items[key]

		if isFound {
			// any found
			return true
		}
	}

	// all not found.
	return false
}

func (it *Hashmap) HasWithLock(key string) bool {
	it.Lock()
	defer it.Unlock()

	_, isFound := it.items[key]

	return isFound
}

// GetKeysFilteredItems must return slice.
func (it *Hashmap) GetKeysFilteredItems(
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

// GetKeysFilteredCollection must return items.
func (it *Hashmap) GetKeysFilteredCollection(
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
				false, filteredList)
		}
	}

	return NewCollectionUsingStrings(
		false, filteredList)
}

func (it *Hashmap) Items() map[string]string {
	return it.items
}

//goland:noinspection GoLinterLocal
func (it *Hashmap) ItemsCopyLock() *map[string]string {
	it.Lock()

	copiedItemsMap := &it.items

	it.Unlock()

	return copiedItemsMap
}

func (it *Hashmap) ValuesCollection() *Collection {
	return NewCollectionUsingStrings(
		false, it.ValuesList())
}

func (it *Hashmap) ValuesHashset() *Hashset {
	return NewHashsetUsingStrings(
		it.ValuesListPtr())
}

func (it *Hashmap) ValuesCollectionLock() *Collection {
	return NewCollectionUsingStrings(
		false, *it.ValuesListCopyPtrLock())
}

func (it *Hashmap) ValuesHashsetLock() *Hashset {
	return NewHashsetUsingStrings(
		it.ValuesListCopyPtrLock())
}

func (it *Hashmap) ValuesList() []string {
	return *it.ValuesListPtr()
}

func (it *Hashmap) ValuesListPtr() *[]string {
	if it.hasMapUpdated || it.cachedList == nil {
		it.setCached()
	}

	return &it.cachedList
}

func (it *Hashmap) KeysValuesCollection() (
	keys, values *Collection,
) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		keys = NewCollectionUsingStringsPtr(
			false,
			it.Keys(),
		)

		wg.Done()
	}()

	go func() {
		values = NewCollectionUsingStringsPtr(
			false,
			it.ValuesListPtr(),
		)

		wg.Done()
	}()

	wg.Wait()

	return keys, values
}

func (it *Hashmap) KeysValuesList() (
	keys, values *[]string,
) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		keys = it.Keys()
		wg.Done()
	}()

	go func() {
		values = it.ValuesListPtr()
		wg.Done()
	}()

	wg.Wait()

	return keys, values
}

func (it *Hashmap) KeysValuePairs() *[]KeyValuePair {
	pairs := make([]KeyValuePair, it.Length())

	i := 0
	for k, v := range it.items {
		pairs[i] = KeyValuePair{
			Key:   k,
			Value: v,
		}

		i++
	}

	return &pairs
}

func (it *Hashmap) KeysValuesListLock() (
	keys, values *[]string,
) {
	it.Lock()
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		keys = it.Keys()
		wg.Done()
	}()
	go func() {
		values = it.ValuesListPtr()
		wg.Done()
	}()

	wg.Wait()
	it.Unlock()

	return keys, values
}

func (it *Hashmap) AllKeys() []string {
	length := len(it.items)
	keys := make([]string, length)

	if length == 0 {
		return keys
	}

	i := 0
	for k := range it.items {
		keys[i] = k
		i++
	}

	return keys
}

func (it *Hashmap) Keys() *[]string {
	keys := it.AllKeys()

	return &keys
}

func (it *Hashmap) KeysCollection() *Collection {
	return NewCollectionUsingStringsPtr(
		false,
		it.Keys(),
	)
}

func (it *Hashmap) KeysLock() *[]string {
	length := it.LengthLock()
	keys := make([]string, length)

	if length == 0 {
		return &keys
	}

	i := 0
	it.Lock()
	for k := range it.items {
		keys[i] = k
		i++
	}

	it.Unlock()

	return &keys
}

// ValuesListCopyPtrLock  a slice must returned
func (it *Hashmap) ValuesListCopyPtrLock() *[]string {
	it.Lock()
	defer it.Unlock()

	return &(*it.ValuesListPtr())
}

func (it *Hashmap) setCached() {
	length := it.Length()
	list := make([]string, length)

	if length == 0 {
		it.cachedList = list
		it.hasMapUpdated = false

		return
	}

	i := 0

	for _, val := range it.items {
		list[i] = val
		i++
	}

	it.hasMapUpdated = false
	it.cachedList = list
}

// ValuesToLower Create a new items with all lower strings
func (it *Hashmap) ValuesToLower() *Hashmap {
	newMap := make(map[string]string, it.Length())

	var toLower string
	for key, value := range it.items {
		toLower = strings.ToLower(key)
		newMap[toLower] = value
	}

	return NewHashmapUsingMap(
		false,
		0,
		newMap,
	)
}

func (it *Hashmap) Length() int {
	if it == nil {
		return 0
	}

	if it.hasMapUpdated || it.length < 0 {
		it.length = len(it.items)
	}

	return it.length
}

func (it *Hashmap) LengthLock() int {
	it.Lock()
	defer it.Unlock()

	return it.Length()
}

//goland:noinspection GoLinterLocal,GoVetCopyLock
func (it *Hashmap) IsEquals(another Hashmap) bool { //nolint:govet
	return it.IsEqualsPtr(&another)
}

func (it *Hashmap) IsEqualsPtrLock(another *Hashmap) bool {
	it.Lock()
	defer it.Unlock()

	return it.IsEqualsPtr(another)
}

func (it *Hashmap) IsEqualsPtr(another *Hashmap) bool {
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

	for key, value := range it.items {
		result, has := another.items[key]

		if !has || !(result != value) {
			return false
		}
	}

	return true
}

func (it *Hashmap) Remove(key string) *Hashmap {
	delete(it.items, key)
	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) RemoveWithLock(key string) *Hashmap {
	it.Lock()
	defer it.Unlock()

	it.Remove(key)

	return it
}

func (it *Hashmap) String() string {
	if it.IsEmpty() {
		return commonJoiner + NoElements
	}

	return commonJoiner +
		strings.Join(
			it.ValuesList(),
			commonJoiner)
}

func (it *Hashmap) StringLock() string {
	if it.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	it.Lock()
	defer it.Unlock()

	return commonJoiner +
		strings.Join(
			*it.ValuesListPtr(),
			commonJoiner)
}

// GetValuesExceptKeysInHashset Get all Collection except the mentioned ones.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashmap
// Set B = anotherHashset given in parameters.
func (it *Hashmap) GetValuesExceptKeysInHashset(
	anotherHashset *Hashset,
) *[]string {
	if anotherHashset == nil || anotherHashset.IsEmpty() {
		return it.ValuesListPtr()
	}

	finalList := make(
		[]string,
		0,
		it.Length())

	for key, value := range it.items {
		if anotherHashset.Has(key) {
			continue
		}

		finalList = append(
			finalList,
			value)
	}

	return &finalList
}

// GetValuesKeysExcept Get all items except the mentioned ones.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashmap
// Set B = items given in parameters.
func (it *Hashmap) GetValuesKeysExcept(
	items *[]string,
) *[]string {
	if items == nil {
		return it.ValuesListPtr()
	}

	newCollection := NewHashsetUsingStrings(
		items)

	return it.GetValuesExceptKeysInHashset(
		newCollection)
}

// GetAllExceptCollection Get all Hashmap items except the mentioned ones in collection.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashmap
// Set B = collection given in parameters.
func (it *Hashmap) GetAllExceptCollection(
	collection *Collection,
) *[]string {
	if collection == nil {
		return it.ValuesListPtr()
	}

	return it.GetValuesExceptKeysInHashset(
		collection.HashsetAsIs())
}

// GetAllExceptCollectionPtr Get all items except the mentioned ones in collectionPtr.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashmap
// Set B = collectionPtr given in parameters.
func (it *Hashmap) GetAllExceptCollectionPtr(
	collectionPtr *CollectionPtr,
) *[]string {
	if collectionPtr == nil {
		return it.ValuesListPtr()
	}

	return it.GetValuesExceptKeysInHashset(
		collectionPtr.HashsetAsIs())
}

// Join values
func (it *Hashmap) Join(
	separator string,
) string {
	return strings.Join(*it.ValuesListPtr(), separator)
}

func (it *Hashmap) JoinKeys(
	separator string,
) string {
	return strings.Join(*it.Keys(), separator)
}

func (it *Hashmap) JsonModel() *HashmapDataModel {
	return NewHashmapsDataModelUsing(it)
}

func (it *Hashmap) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *Hashmap) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *Hashmap) UnmarshalJSON(data []byte) error {
	var dataModel HashmapDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.items = dataModel.Items
		it.length = len(it.items)
		it.hasMapUpdated = false
		it.isEmptySet = it.length == 0
		it.cachedList = nil
	}

	return err
}

func (it Hashmap) Json() corejson.Result {
	return corejson.NewFromAny(it)
}

func (it Hashmap) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *Hashmap) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Hashmap, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return EmptyHashmap(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *Hashmap) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Hashmap {
	hashSet, err := it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}

func (it *Hashmap) ToError(sep string) error {
	return errcore.SliceError(sep, it.KeyValStringLines())
}

func (it *Hashmap) ToDefaultError() error {
	return errcore.SliceError(
		constants.NewLineUnix, it.KeyValStringLines())
}

func (it *Hashmap) KeyValStringLines() *[]string {
	return it.ToStringsUsingCompiler(func(key, val string) string {
		return key + constants.HyphenAngelRight + val
	})
}

func (it *Hashmap) Clear() *Hashmap {
	if it == nil {
		return it
	}

	it.items = nil
	it.items = map[string]string{}
	it.cachedList = it.cachedList[:0]
	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) Dispose() {
	if it == nil {
		return
	}

	it.items = nil
	it.cachedList = nil
}

func (it *Hashmap) ToStringsUsingCompiler(
	compilerFunc func(
		key,
		val string,
	) string,
) *[]string {
	length := it.Length()
	slice := make([]string, length)

	if length == 0 {
		return &slice
	}

	index := 0
	for key, val := range it.items {
		line := compilerFunc(key, val)
		slice[index] = line

		index++
	}

	return &slice
}

func (it *Hashmap) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Hashmap) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *Hashmap) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Hashmap) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *Hashmap) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it *Hashmap) ClonePtr() *Hashmap {
	if it == nil {
		return nil
	}

	cloned := it.Clone()

	return &cloned
}

func (it Hashmap) Clone() Hashmap {
	empty := EmptyHashmap()
	jsonResult := it.JsonPtr()

	return *empty.ParseInjectUsingJsonMust(jsonResult)
}
