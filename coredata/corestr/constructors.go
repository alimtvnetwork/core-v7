package corestr

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
)

// --------- Hashset starts ----------

func EmptyHashset() *Hashset {
	return NewHashset(0)
}

func NewHashset(length int) *Hashset {
	hashset := make(map[string]bool, length)

	return &Hashset{
		items:         &hashset,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    true,
		Mutex:         sync.Mutex{},
	}
}

// addCapacity will not work if it is not a clone.
//goland:noinspection ALL
func NewHashsetWithValues(
	addCapacity int,
	isMakeClone bool,
	items ...string,
) *Hashset {
	if items == nil {
		return EmptyHashset()
	}

	return NewHashsetUsingStrings(
		&items,
		addCapacity,
		isMakeClone)
}

// addCapacity will not work if it is not a clone.
func NewHashsetUsingStringPointersArray(
	inputArray *[]*string,
	addCapacity int,
	isMakeClone bool,
) *Hashset {
	if inputArray == nil || *inputArray == nil {
		return NewHashset(defaultHashsetItems)
	}

	maps := converters.StringsPointersToStringBoolMap(inputArray)

	return NewHashsetUsingMap(
		maps,
		addCapacity,
		isMakeClone)
}

// addCapacity will not work if it is not a clone.
func NewHashsetUsingCollection(
	collection *Collection,
	addCapacity int,
	isMakeClone bool,
) *Hashset {
	if collection == nil || collection.IsEmpty() {
		return EmptyHashset()
	}

	return NewHashsetUsingStrings(
		collection.items,
		addCapacity,
		isMakeClone)
}

// addCapacity will not work if it is not a clone.
func NewHashsetUsingStrings(
	inputArray *[]string,
	addCapacity int,
	isMakeClone bool,
) *Hashset {
	if inputArray == nil || *inputArray == nil {
		return EmptyHashset()
	}

	maps := converters.StringsToMap(inputArray)

	return NewHashsetUsingMap(
		maps,
		addCapacity,
		isMakeClone)
}

// addCapacity will not work if it is not a clone.
func NewHashsetUsingMap(
	itemsMap *map[string]bool,
	addCapacity int,
	isMakeClone bool,
) *Hashset {
	if itemsMap == nil || *itemsMap == nil {
		return NewHashset(defaultHashsetItems)
	}

	length := len(*itemsMap)

	if isMakeClone {
		hashset := NewHashset(length + addCapacity)

		hashset.AddItemsMap(itemsMap)

		return hashset
	}

	return &Hashset{
		items:         itemsMap,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    length == 0,
		Mutex:         sync.Mutex{},
	}
}

// --------- ToCollection starts ----------

func NewCollection(capacity int) *Collection {
	collection := make([]string, 0, capacity)

	return &Collection{
		items: &collection,
	}
}

func EmptyCollection() *Collection {
	collection := make([]string, 0)

	return &Collection{
		items: &collection,
	}
}

func NewCollectionUsingStrings(stringItems *[]string, isMakeClone bool) *Collection {
	if isMakeClone {
		cloned := *stringItems

		return &Collection{
			items: &cloned,
		}
	}

	return &Collection{
		items: stringItems,
	}
}

func NewCollectionUsingStringsPlusCap(stringItems *[]string, capacity int) *Collection {
	length := LengthOfStrings(stringItems)
	collection := NewCollection(length + capacity)

	return collection.AddStringsPtr(stringItems)
}

func NewCollectionUsingPointerStringsPlusCap(stringItems *[]*string, capacity int) *Collection {
	length := LengthOfPointerStrings(stringItems)
	collection := NewCollection(length + capacity)

	return collection.AddPointerStringsPtr(stringItems)
}

//goland:noinspection ALL
func NewCollectionUsingLength(len, capacity int) *Collection {
	collection := make([]string, len, capacity)

	return &Collection{
		items: &collection,
	}
}

// --------- CollectionPtr starts ----------

func NewCollectionPtr(capacity int) *CollectionPtr {
	collection := make([]*string, 0, capacity)

	return &CollectionPtr{
		items: &collection,
	}
}

func EmptyCollectionPtr() *CollectionPtr {
	collection := make([]*string, 0)

	return &CollectionPtr{
		items: &collection,
	}
}

//goland:noinspection ALL
func NewCollectionPtrUsingPointerStrings(
	stringItems *[]*string,
	addCapacity int,
) *CollectionPtr {
	if addCapacity == 0 {
		return &CollectionPtr{
			items: stringItems,
		}
	}

	if stringItems == nil {
		return NewCollectionPtr(addCapacity)
	}

	length := len(*stringItems)
	collection := NewCollectionPtr(length + addCapacity)

	return collection.
		AddPointerStringsPtr(stringItems)
}

func NewCollectionPtrUsingStrings(
	stringItems *[]string,
	addCapacity int,
) *CollectionPtr {
	if addCapacity == 0 {
		return &CollectionPtr{
			items: converters.StringsToPointerStrings(stringItems),
		}
	}

	if stringItems == nil {
		return NewCollectionPtr(addCapacity)
	}

	length := len(*stringItems)
	collection := NewCollectionPtr(length + addCapacity)

	return collection.
		AddStringsPtr(stringItems)
}

//goland:noinspection ALL
func NewCollectionPtrUsingLength(len, capacity int) *CollectionPtr {
	collection := make([]*string, len, capacity)

	return &CollectionPtr{
		items: &collection,
	}
}

// --------- CharCollectionMap starts ----------

// CharCollectionMap.eachCollectionCapacity, capacity minimum 10 will be set if lower than 10 is given.
//
// For lower than 5 use the EmptyCharCollectionMap items definition.
func NewCharCollectionMap(
	capacity, selfCollectionCapacity int,
) *CharCollectionMap {
	if capacity < charCollectionDefaultCapacity {
		capacity = charCollectionDefaultCapacity
	}

	mapElements := make(map[byte]*Collection, capacity)

	if selfCollectionCapacity < charCollectionDefaultCapacity {
		selfCollectionCapacity = charCollectionDefaultCapacity
	}

	return &CharCollectionMap{
		items:                  &mapElements,
		eachCollectionCapacity: selfCollectionCapacity,
	}
}

// eachCollectionCapacity = 0
func EmptyCharCollectionMap() *CharCollectionMap {
	mapElements := make(map[byte]*Collection, 0)

	return &CharCollectionMap{
		items:                  &mapElements,
		eachCollectionCapacity: defaultEachCollectionCapacity,
	}
}

//goland:noinspection ALL
func NewCharCollectionMapUsingItems(
	items []string,
) *CharCollectionMap {
	if items == nil {
		return EmptyCharCollectionMap()
	}

	return NewCharCollectionMapUsingItemsPtr(
		&items)
}

func NewCharCollectionMapUsingItemsPtr(
	items *[]string,
) *CharCollectionMap {
	if items == nil {
		return EmptyCharCollectionMap()
	}

	length := len(*items)
	if length == 0 {
		return EmptyCharCollectionMap()
	}

	mapElements := make(map[byte]*Collection, length)
	charCollectionMap := &CharCollectionMap{
		items:                  &mapElements,
		eachCollectionCapacity: 0,
	}

	charCollectionMap.AddStringsPtr(items)

	return charCollectionMap
}

//goland:noinspection ALL
func NewCharCollectionMapUsingItemsPlusCap(
	items *[]string,
	additionalCapacityOrLength int,
	eachCollectionCapacity int,
) *CharCollectionMap {
	isDefined := items != nil && *items != nil
	length := 0
	if isDefined {
		length = len(*items)
		additionalCapacityOrLength += length
	}

	mapElements := make(map[byte]*Collection, additionalCapacityOrLength)

	charCollectionMap := &CharCollectionMap{
		items:                  &mapElements,
		eachCollectionCapacity: eachCollectionCapacity,
	}

	if !isDefined || length == 0 {
		return charCollectionMap
	}

	return charCollectionMap.
		AddStringsPtr(items)
}

// --------- HashsetsCollection starts ----------

func EmptyHashsetsCollection() *HashsetsCollection {
	collection := make([]*Hashset, 0, 0)

	return &HashsetsCollection{
		items: &collection,
	}
}

func NewHashsetsCollection(hashsets *[]Hashset) *HashsetsCollection {
	if hashsets == nil ||
		*hashsets == nil {
		return EmptyHashsetsCollection()
	}

	length := len(*hashsets)
	collection := make(
		[]*Hashset,
		length,
		length+constants.ArbitraryCapacity10)

	//goland:noinspection GoLinterLocal,GoVetCopyLock
	for i, hashset := range *hashsets { //nolint:govet
		//goland:noinspection GoLinterLocal
		collection[i] = &hashset
	}

	return &HashsetsCollection{
		items: &collection,
	}
}

func NewHashsetsCollectionUsingPointerHashsets(hashsets *[]*Hashset) *HashsetsCollection {
	if hashsets == nil ||
		*hashsets == nil {
		return EmptyHashsetsCollection()
	}

	return &HashsetsCollection{
		items: hashsets,
	}
}

//goland:noinspection ALL
func NewHashsetsCollectionUsingLength(len, capacity int) *HashsetsCollection {
	collection := make([]*Hashset, len, capacity)

	return &HashsetsCollection{
		items: &collection,
	}
}

// --------- Hashmap starts ----------

func EmptyHashmap() *Hashmap {
	return NewHashmap(0)
}

func NewHashmap(length int) *Hashmap {
	hashset := make(map[string]string, length)

	return &Hashmap{
		items:         &hashset,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    true,
		Mutex:         sync.Mutex{},
	}
}

//goland:noinspection ALL
func NewHashmapUsingKeyAnyValues(keyAnyValues *[]KeyAnyValuePair) *Hashmap {
	if keyAnyValues == nil || *keyAnyValues == nil {
		return NewHashmap(defaultHashsetItems)
	}

	length := len(*keyAnyValues)
	hashMap := NewHashmap(length + constants.ArbitraryCapacity10)
	hashMap.AddOrUpdateKeyAnyValsPtr(keyAnyValues)

	return hashMap
}

//goland:noinspection ALL
func NewHashmapUsingKeyValues(keyValues *[]KeyValuePair) *Hashmap {
	if keyValues == nil || *keyValues == nil {
		return NewHashmap(defaultHashsetItems)
	}

	length := len(*keyValues)
	hashMap := NewHashmap(length + constants.ArbitraryCapacity10)
	hashMap.AddOrUpdateKeyValsPtr(keyValues)

	return hashMap
}

//goland:noinspection ALL
func NewHashmapUsingCollection(keys, values *Collection) *Hashmap {
	if keys == nil || keys.IsEmpty() {
		return EmptyHashmap()
	}

	itemsMap := converters.KeysValuesStringsToMap(
		keys.items,
		values.items)

	return NewHashmapUsingMap(
		itemsMap,
		0,
		false)
}

//goland:noinspection ALL
func NewHashmapUsingStrings(keys, values *[]string) *Hashmap {
	if keys == nil || *keys == nil {
		return EmptyHashmap()
	}

	itemsMap := converters.KeysValuesStringsToMap(
		keys,
		values)

	return NewHashmapUsingMap(
		itemsMap,
		0,
		false)
}

// isMakeClone : copies itemsMap or else use the same one as pointer assign.
func NewHashmapUsingMap(
	itemsMap *map[string]string,
	addCapacity int,
	isMakeClone bool,
) *Hashmap {
	if itemsMap == nil || *itemsMap == nil {
		return NewHashmap(defaultHashsetItems)
	}

	length := len(*itemsMap)

	if isMakeClone {
		hashMap := NewHashmap(length + addCapacity)
		hashMap.AddOrUpdateMap(itemsMap)

		return hashMap
	}

	return &Hashmap{
		items:         itemsMap,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    length == 0,
	}
}

// always returns the clone of the items.
//goland:noinspection ALL
func NewHashmapUsingMapUsingAddCapacity(
	itemsMap *map[string]string,
	addCapacity int,
) *Hashmap {
	if itemsMap == nil || *itemsMap == nil {
		return NewHashmap(defaultHashsetItems)
	}

	length := len(*itemsMap)
	hashMap := NewHashmap(length + addCapacity)
	hashMap.AddOrUpdateMap(itemsMap)

	return hashMap
}

// --------- LinkedList starts ----------

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func EmptyLinkedList() *LinkedList {
	return &LinkedList{}
}

//goland:noinspection ALL
func NewLinkedListUsingPointerStrings(
	stringItems *[]*string,
) *LinkedList {
	if stringItems == nil {
		return &LinkedList{}
	}

	linkedList := NewLinkedList()

	return linkedList.
		AddPointerStringsPtr(stringItems)
}

//goland:noinspection ALL
func NewLinkedListUsingStringsPtr(
	stringItems *[]string,
) *LinkedList {
	if stringItems == nil {
		return &LinkedList{}
	}

	linkedList := NewLinkedList()

	return linkedList.
		AddStringsPtr(stringItems)
}

// --------- CharHashset starts ----------

// CharHashsetMap.eachHashsetCapacity, capacity minimum 10 will be set if lower than 10 is given.
//
// For lower than 5 use the EmptyCharHashsetMap hashset definition.
func NewCharHashsetMap(
	capacity, selfHashsetCapacity int,
) *CharHashsetMap {
	const limit = constants.ArbitraryCapacity10

	if capacity < limit {
		capacity = limit
	}

	mapElements := make(
		map[byte]*Hashset,
		capacity)

	if selfHashsetCapacity < limit {
		selfHashsetCapacity = limit
	}

	return &CharHashsetMap{
		items:               &mapElements,
		eachHashsetCapacity: selfHashsetCapacity,
	}
}

//goland:noinspection ALL
func NewCharHashsetMapUsingItemsPlusCap(
	items *[]string,
	capacity, selfHashsetCapacity int,
) *CharHashsetMap {
	charHashsetMap := NewCharHashsetMap(capacity, selfHashsetCapacity)

	charHashsetMap.AddStringsPtr(items)

	return charHashsetMap
}

func NewCharHashsetMapUsingItems(
	items []string,
	selfHashsetCapacity int,
) *CharHashsetMap {
	if items == nil {
		return NewCharHashsetMap(
			constants.ArbitraryCapacity5,
			selfHashsetCapacity)
	}

	length := len(items)
	charHashsetMap := NewCharHashsetMap(
		length,
		selfHashsetCapacity)

	charHashsetMap.AddStrings(items...)

	return charHashsetMap
}

func NewCharHashsetMapUsingItemsPtr(
	items *[]string,
	selfHashsetCapacity int,
) *CharHashsetMap {
	if items == nil {
		return NewCharHashsetMap(
			constants.ArbitraryCapacity5,
			selfHashsetCapacity)
	}

	length := len(*items)
	charHashsetMap := NewCharHashsetMap(
		length,
		selfHashsetCapacity)

	charHashsetMap.AddStringsPtr(items)

	return charHashsetMap
}

// eachHashsetCapacity = 0
func EmptyCharHashsetMap() *CharHashsetMap {
	mapElements := make(
		map[byte]*Hashset,
		0)

	return &CharHashsetMap{
		items:               &mapElements,
		eachHashsetCapacity: 0,
	}
}

// --------- LinkedCollections starts ----------

func NewLinkedCollections() *LinkedCollections {
	return &LinkedCollections{}
}

func EmptyLinkedCollections() *LinkedCollections {
	return &LinkedCollections{}
}

//goland:noinspection ALL
func NewLinkedCollectionsUsingPointerStrings(
	stringItems *[]*string,
) *LinkedCollections {
	if stringItems == nil {
		return &LinkedCollections{}
	}

	linkedList := NewLinkedCollections()

	return linkedList.
		AddPointerStringsPtr(stringItems)
}

func NewLinkedCollectionsUsingCollections(
	collections ...*Collection,
) *LinkedCollections {
	if collections == nil {
		return &LinkedCollections{}
	}

	linkedList := NewLinkedCollections()

	return linkedList.
		AppendCollectionsPointers(true, &collections)
}

//goland:noinspection ALL
func NewLinkedCollectionsUsingStringsPtr(
	stringItems *[]string,
	isMakeClone bool,
) *LinkedCollections {
	if stringItems == nil {
		return &LinkedCollections{}
	}

	linkedList := NewLinkedCollections()

	return linkedList.
		AddStringsPtr(stringItems, isMakeClone)
}

// --------- CollectionsOfCollection starts ----------

func NewCollectionsOfCollection(
	capacity int,
) *CollectionsOfCollection {
	collection := make([]*Collection, 0, capacity)

	return &CollectionsOfCollection{
		items: &collection,
	}
}

func EmptyCollectionsOfCollection() *CollectionsOfCollection {
	collection := make([]*Collection, 0)

	return &CollectionsOfCollection{
		items: &collection,
	}
}

func NewCollectionsOfCollectionUsingStringsOfStrings(
	isMakeClone bool,
	stringItems ...*[]string,
) *CollectionsOfCollection {
	length := LengthOfStringsOfPointerStrings(&stringItems)

	return NewCollectionsOfCollectionUsingLength(
		0,
		length,
	).AddsStringsOfPointerStrings(isMakeClone, &stringItems)
}

func NewCollectionsOfCollectionUsingStringsOfPointerStrings(
	isMakeClone bool,
	stringItems *[]*[]string,
) *CollectionsOfCollection {
	length := LengthOfStringsOfPointerStrings(stringItems)

	return NewCollectionsOfCollectionUsingLength(
		0,
		length,
	).AddsStringsOfPointerStrings(isMakeClone, stringItems)
}

func NewCollectionsOfCollectionUsingStrings(
	stringItems *[]string,
	isMakeClone bool,
) *CollectionsOfCollection {
	length := LengthOfStrings(stringItems)

	return NewCollectionsOfCollectionUsingLength(
		0,
		length,
	).AddStringsPtr(stringItems, isMakeClone)
}

func NewCollectionsOfCollectionUsingStringsPlusCap(
	stringItems *[]string,
	capacity int,
	isMakeClone bool,
) *CollectionsOfCollection {
	length := LengthOfStrings(stringItems)
	collection := NewCollectionsOfCollection(length + capacity)

	return collection.AddStringsPtr(stringItems, isMakeClone)
}

func NewCollectionsOfCollectionUsingPointerStringsPlusCap(
	stringItems *[]*string, capacity int,
) *CollectionsOfCollection {
	length := LengthOfPointerStrings(stringItems)
	collection := NewCollectionsOfCollection(length + capacity)

	return collection.AddPointerStringsPtr(stringItems)
}

//goland:noinspection ALL
func NewCollectionsOfCollectionUsingLength(
	len,
	capacity int,
) *CollectionsOfCollection {
	collection := make([]*Collection, len, capacity)

	return &CollectionsOfCollection{
		items: &collection,
	}
}
