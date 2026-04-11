package corestr

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/mapdiffinternal"
)

type Hashmap struct {
	hasMapUpdated bool
	items         map[string]string
	cachedList    []string
	sync.RWMutex
}

func (it *Hashmap) IsEmpty() bool {
	return it == nil || len(it.items) == 0
}

func (it *Hashmap) HasItems() bool {
	return it != nil && !it.IsEmpty()
}

func (it *Hashmap) Collection() *Collection {
	return New.Collection.StringsOptions(false, it.ValuesList())
}

func (it *Hashmap) IsEmptyLock() bool {
	it.RLock()
	defer it.RUnlock()

	return it.IsEmpty()
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
	val any,
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
) (isAddedNewly bool) {
	_, isAlreadyExist := it.items[keyVal.Key]

	it.items[keyVal.Key] = keyVal.Value
	it.hasMapUpdated = true

	return !isAlreadyExist
}

func (it *Hashmap) AddOrUpdate(key, val string) (isAddedNewly bool) {
	_, isAlreadyExist := it.items[key]

	it.items[key] = val
	it.hasMapUpdated = true

	return !isAlreadyExist
}

func (it *Hashmap) Set(key, val string) (isAddedNewly bool) {
	_, isAlreadyExist := it.items[key]

	it.items[key] = val
	it.hasMapUpdated = true

	return !isAlreadyExist
}

func (it *Hashmap) SetTrim(key, val string) (isAddedNewly bool) {
	key = strings.TrimSpace(key)
	val = strings.TrimSpace(val)

	return it.Set(key, val)
}

func (it *Hashmap) SetBySplitter(
	splitter, line string,
) (isAddedNewly bool) {
	splits := strings.SplitN(
		line, splitter, constants.Two,
	)

	if len(splits) >= 2 {
		// all okay

		return it.Set(splits[0], splits[len(splits)-1])
	}

	return it.Set(splits[0], "")
}

func safeWaitGroupDone(wg *sync.WaitGroup) {
	if wg == nil {
		return
	}

	defer func() {
		_ = recover()
	}()

	wg.Done()
}

func (it *Hashmap) AddOrUpdateStringsPtrWgLock(
	wg *sync.WaitGroup,
	keys, values []string,
) *Hashmap {
	if len(keys) != len(values) {
		panic(
			fmt.Sprintf(
				"cannot add keys (%d) and values (%d) with different lengths",
				len(keys),
				len(values),
			),
		)
	}

	if len(keys) == 0 {
		// See issues/corestrtests-waitgroup-deadlock-empty-keys.md
		// and issues/corestrtests-wg-negative-counter-panic.md
		safeWaitGroupDone(wg)
		return it
	}

	it.RLock()
	for i, key := range keys {
		it.items[key] = values[i]
	}

	it.hasMapUpdated = true
	it.RUnlock()
	safeWaitGroupDone(wg)

	return it
}

func (it *Hashmap) AddOrUpdateHashmap(
	nextHashmap *Hashmap,
) *Hashmap {
	if nextHashmap == nil {
		return it
	}

	for key, val := range nextHashmap.items {
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

func (it *Hashmap) AddOrUpdateKeyAnyValues(
	pairs ...KeyAnyValuePair,
) *Hashmap {
	if len(pairs) == 0 {
		return it
	}

	for _, pair := range pairs {
		it.items[pair.Key] = pair.ValueString()
	}

	it.hasMapUpdated = true

	return it
}

func (it *Hashmap) AddOrUpdateKeyValues(
	pairs ...KeyValuePair,
) *Hashmap {
	if len(pairs) == 0 {
		return it
	}

	for _, pair := range pairs {
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

	isLengthMismatch := keys.Length() != values.Length()

	if isLengthMismatch {
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
			it.RLock()
			it.items[pair.Key] = result
			it.RUnlock()

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
		return New.Hashmap.UsingMapOptions(
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

		length += h.Length()
	}

	newHashmap := New.Hashmap.UsingMapOptions(
		true,
		length,
		it.items,
	)

	newHashmap.AddOrUpdateHashmap(it)

	for _, hashmap2 := range hashmaps {
		newHashmap.AddOrUpdateHashmap(
			hashmap2,
		)
	}

	return newHashmap
}

func (it *Hashmap) ConcatNewUsingMaps(
	isCloneOnEmptyAsWell bool,
	hashmaps ...map[string]string,
) *Hashmap {
	if len(hashmaps) == 0 {
		return New.Hashmap.UsingMapOptions(
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

		length += len(h)
	}

	newHashmap := New.Hashmap.UsingMapOptions(
		true,
		length,
		it.items,
	)

	newHashmap.AddOrUpdateHashmap(it)

	for _, nextMap := range hashmaps {
		newHashmap.AddOrUpdateMap(
			nextMap,
		)
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

func (it *Hashmap) Contains(key string) bool {
	_, isFound := it.items[key]

	return isFound
}

func (it *Hashmap) ContainsLock(key string) bool {
	it.RLock()
	_, isFound := it.items[key]
	it.RUnlock()

	return isFound
}

func (it *Hashmap) IsKeyMissing(key string) bool {
	_, isFound := it.items[key]

	return !isFound
}

func (it *Hashmap) IsKeyMissingLock(key string) bool {
	it.RLock()
	_, isFound := it.items[key]
	it.RUnlock()

	return !isFound
}

func (it *Hashmap) HasLock(key string) bool {
	it.RLock()
	_, isFound := it.items[key]
	it.RUnlock()

	return isFound
}

func (it *Hashmap) HasAllStrings(keys ...string) bool {
	for _, key := range keys {
		_, isFound := it.items[key]
		isMissing := !isFound

		if isMissing {
			// not found
			return false
		}
	}

	// all found.
	return true
}

func (it *Hashmap) DiffRaw(
	rightMap map[string]string,
) map[string]string {
	mapDiffer := mapdiffinternal.HashmapDiff(it.items)

	return mapDiffer.DiffRaw(rightMap)
}

func (it *Hashmap) Diff(
	rightMap *Hashmap,
) *Hashmap {
	rawMap := it.DiffRaw(rightMap.Items())

	return New.Hashmap.UsingMap(rawMap)
}

// HasAllCollectionItems return false on items is nil or Empty.
func (it *Hashmap) HasAllCollectionItems(
	collection *Collection,
) bool {
	if collection == nil || collection.IsEmpty() {
		return false
	}

	return it.HasAllStrings(collection.List()...)
}

func (it *Hashmap) HasAll(keys ...string) bool {
	for _, key := range keys {
		_, isFound := it.items[key]
		isMissing := !isFound

		if isMissing {
			// not found
			return false
		}
	}

	// all found.
	return true
}

func (it *Hashmap) HasAnyItem() bool {
	return it != nil && it.Length() > 0
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
	it.RLock()
	defer it.RUnlock()

	_, isFound := it.items[key]

	return isFound
}

// GetKeysFilteredItems must return slice.
func (it *Hashmap) GetKeysFilteredItems(
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
		result, isKeep, isBreak :=
			filter(key, i)

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

// GetKeysFilteredCollection must return items.
func (it *Hashmap) GetKeysFilteredCollection(
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
				false, filteredList,
			)
		}
	}

	return New.Collection.StringsOptions(
		false, filteredList,
	)
}

func (it *Hashmap) Items() map[string]string {
	return it.items
}

func (it *Hashmap) SafeItems() map[string]string {
	if it == nil {
		return nil
	}

	return it.items
}

//goland:noinspection GoLinterLocal
func (it *Hashmap) ItemsCopyLock() *map[string]string {
	it.RLock()
	defer it.RUnlock()

	copiedMap := make(map[string]string, len(it.items))
	for k, v := range it.items {
		copiedMap[k] = v
	}

	return &copiedMap
}

func (it *Hashmap) ValuesCollection() *Collection {
	return New.Collection.StringsOptions(
		false, it.ValuesList(),
	)
}

func (it *Hashmap) ValuesHashset() *Hashset {
	return New.Hashset.Strings(
		it.ValuesList(),
	)
}

func (it *Hashmap) ValuesCollectionLock() *Collection {
	return New.Collection.StringsOptions(
		false, it.ValuesListCopyLock(),
	)
}

func (it *Hashmap) ValuesHashsetLock() *Hashset {
	return New.Hashset.Strings(
		it.ValuesListCopyLock(),
	)
}

func (it *Hashmap) ValuesList() []string {
	if it.hasMapUpdated || it.cachedList == nil {
		it.setCached()
	}

	return it.cachedList
}

func (it *Hashmap) KeysValuesCollection() (
	keys, values *Collection,
) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		keys = New.Collection.Strings(
			it.Keys(),
		)

		wg.Done()
	}()

	go func() {
		values = New.Collection.Strings(
			it.ValuesList(),
		)

		wg.Done()
	}()

	wg.Wait()

	return keys, values
}

func (it *Hashmap) KeysValuesList() (
	keys, values []string,
) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		keys = it.Keys()
		wg.Done()
	}()

	go func() {
		values = it.ValuesList()
		wg.Done()
	}()

	wg.Wait()

	return keys, values
}

func (it *Hashmap) KeysValuePairs() []*KeyValuePair {
	pairs := make([]*KeyValuePair, it.Length())

	i := 0
	for k, v := range it.items {
		pairs[i] = &KeyValuePair{
			Key:   k,
			Value: v,
		}

		i++
	}

	return pairs
}
func (it *Hashmap) KeysValuePairsCollection() *KeyValueCollection {
	pairs := New.KeyValues.Cap(it.Length())

	for k, v := range it.items {
		pairs.Add(k, v)
	}

	return pairs
}

func (it *Hashmap) KeysValuesListLock() (
	keys, values []string,
) {
	it.RLock()
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		keys = it.Keys()
		wg.Done()
	}()
	go func() {
		values = it.ValuesList()
		wg.Done()
	}()

	wg.Wait()
	it.RUnlock()

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

func (it *Hashmap) Keys() []string {
	return it.AllKeys()
}

func (it *Hashmap) KeysCollection() *Collection {
	return New.Collection.Strings(
		it.Keys(),
	)
}

func (it *Hashmap) KeysLock() []string {
	length := it.LengthLock()
	keys := make([]string, length)

	if length == 0 {
		return keys
	}

	i := 0
	it.RLock()
	for k := range it.items {
		keys[i] = k
		i++
	}

	it.RUnlock()

	return keys
}

// ValuesListCopyPtrLock
//
//	a slice must be returned
func (it *Hashmap) ValuesListCopyLock() []string {
	it.RLock()
	defer it.RUnlock()

	return it.ValuesList()
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

func (it *Hashmap) ValuesToLower() *Hashmap {
	return it.KeysToLower()
}

// KeysToLower creates a new hashmap with all keys lowercased.
func (it *Hashmap) KeysToLower() *Hashmap {
	newMap := make(map[string]string, it.Length())

	for key, value := range it.items {
		toLower := strings.ToLower(key)
		newMap[toLower] = value
	}

	return New.Hashmap.UsingMapOptions(
		false,
		0,
		newMap,
	)
}

func (it *Hashmap) Length() int {
	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

func (it *Hashmap) LengthLock() int {
	it.RLock()
	defer it.RUnlock()

	return it.Length()
}

//goland:noinspection GoLinterLocal,GoVetCopyLock
func (it *Hashmap) IsEqual(another Hashmap) bool { //nolint:govet
	return it.IsEqualPtr(&another)
}

func (it *Hashmap) IsEqualPtrLock(another *Hashmap) bool {
	it.RLock()
	defer it.RUnlock()

	return it.IsEqualPtr(another)
}

func (it *Hashmap) IsEqualPtr(another *Hashmap) bool {
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

		if !has || result != value {
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
			commonJoiner,
		)
}

func (it *Hashmap) StringLock() string {
	if it.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	it.RLock()
	defer it.RUnlock()

	return commonJoiner +
		strings.Join(
			it.ValuesList(),
			commonJoiner,
		)
}

// GetValuesExceptKeysInHashset Get all Collection except the mentioned ones.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashmap
// Set B = anotherHashset given in parameters.
func (it *Hashmap) GetValuesExceptKeysInHashset(
	anotherHashset *Hashset,
) []string {
	if anotherHashset == nil || anotherHashset.IsEmpty() {
		return it.ValuesList()
	}

	finalList := make(
		[]string,
		0,
		it.Length(),
	)

	for key, value := range it.items {
		if anotherHashset.Has(key) {
			continue
		}

		finalList = append(
			finalList,
			value,
		)
	}

	return finalList
}

// GetValuesKeysExcept Get all items except the mentioned ones.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashmap
// Set B = items given in parameters.
func (it *Hashmap) GetValuesKeysExcept(
	items []string,
) []string {
	if items == nil {
		return it.ValuesList()
	}

	newCollection := New.Hashset.Strings(
		items,
	)

	return it.GetValuesExceptKeysInHashset(
		newCollection,
	)
}

// GetAllExceptCollection Get all Hashmap items except the mentioned ones in collection.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this Hashmap
// Set B = collection given in parameters.
func (it *Hashmap) GetAllExceptCollection(
	collection *Collection,
) []string {
	if collection == nil {
		return it.ValuesList()
	}

	return it.GetValuesExceptKeysInHashset(
		collection.HashsetAsIs(),
	)
}

// Join values
func (it *Hashmap) Join(
	separator string,
) string {
	return strings.Join(it.ValuesList(), separator)
}

func (it *Hashmap) JoinKeys(
	separator string,
) string {
	return strings.Join(it.Keys(), separator)
}

func (it *Hashmap) JsonModel() map[string]string {
	return it.items
}

func (it *Hashmap) JsonModelAny() any {
	return it.JsonModel()
}

func (it *Hashmap) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *Hashmap) UnmarshalJSON(data []byte) error {
	var dataModelItems map[string]string
	err := json.Unmarshal(data, &dataModelItems)

	if err == nil {
		it.items = dataModelItems
		it.hasMapUpdated = true
		it.cachedList = nil
	}

	return err
}

func (it Hashmap) Json() corejson.Result {
	return corejson.New(&it)
}

func (it Hashmap) JsonPtr() *corejson.Result {
	return corejson.NewPtr(&it)
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *Hashmap) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Hashmap, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return Empty.Hashmap(), err
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
		constants.NewLineUnix, it.KeyValStringLines(),
	)
}

func (it *Hashmap) KeyValStringLines() []string {
	return it.ToStringsUsingCompiler(
		func(key, val string) string {
			return key + constants.HyphenAngelRight + val
		},
	)
}

func (it *Hashmap) Clear() *Hashmap {
	if it == nil {
		return nil
	}

	it.items = map[string]string{}

	if it.cachedList != nil {
		it.cachedList = it.cachedList[:0]
	}

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
) []string {
	length := it.Length()
	slice := make([]string, length)

	if length == 0 {
		return slice
	}

	index := 0
	for key, val := range it.items {
		line := compilerFunc(key, val)
		slice[index] = line

		index++
	}

	return slice
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
	cloned := *Empty.Hashmap()

	if len(it.items) == 0 {
		return cloned
	}

	cloned.items = make(map[string]string, len(it.items))
	for key, val := range it.items {
		cloned.items[key] = val
	}
	cloned.hasMapUpdated = it.hasMapUpdated

	return cloned
}

func (it *Hashmap) Get(key string) (val string, isFound bool) {
	val, isFound = it.items[key]

	return val, isFound
}

// GetValue is an alias for Get.
func (it *Hashmap) GetValue(key string) (val string, isFound bool) {
	return it.Get(key)
}

func (it *Hashmap) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *Hashmap) Deserialize(toPtr any) (parsingErr error) {
	return it.JsonPtr().Deserialize(toPtr)
}
