package corestr

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
)

// --------- Hashset starts ----------

func EmptyHashset() *Hashset {
	return NewHashset(constants.Zero)
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

// NewHashsetWithValues addCapacity will not work if it is not a clone.
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
		&items)
}

// NewHashsetUsingStringPointersArray addCapacity will not work if it is not a clone.
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

// NewHashsetUsingCollection addCapacity will not work if it is not a clone.
func NewHashsetUsingCollection(
	collection *Collection,
) *Hashset {
	if collection == nil || collection.IsEmpty() {
		return EmptyHashset()
	}

	return NewHashsetUsingStrings(
		collection.items)
}

// NewHashsetUsingStrings addCapacity will not work if it is not a clone.
func NewHashsetUsingStrings(
	inputArray *[]string,
) *Hashset {
	if inputArray == nil || *inputArray == nil {
		return EmptyHashset()
	}

	maps := converters.StringsToMap(inputArray)

	return NewHashsetUsingMap(
		maps,
		constants.Zero,
		false)
}

// NewHashsetUsingMap addCapacity will not work if it is not a clone.
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
		isEmptySet:    length == constants.Zero,
		Mutex:         sync.Mutex{},
	}
}

// --------- ToCollection starts ----------

func NewCollection(capacity int) *Collection {
	collection := make([]string, constants.Zero, capacity)

	return &Collection{
		items: &collection,
	}
}

func EmptyCollection() *Collection {
	collection := make([]string, constants.Zero)

	return &Collection{
		items: &collection,
	}
}

func NewCollectionUsingStrings(stringItems *[]string, isMakeClone bool) *Collection {
	if isMakeClone {
		length := LengthOfStrings(stringItems)
		slice := make([]string, 0, length+constants.Capacity4)

		collection := &Collection{
			items: &slice,
		}

		return collection.AddStringsPtr(stringItems)
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

func NewCollectionUsingLength(length, capacity int) *Collection {
	collection := make([]string, length, capacity)

	return &Collection{
		items: &collection,
	}
}

// --------- CollectionPtr starts ----------

func NewCollectionPtr(capacity int) *CollectionPtr {
	collection := make([]*string, constants.Zero, capacity)

	return &CollectionPtr{
		items: &collection,
	}
}

func EmptyCollectionPtr() *CollectionPtr {
	collection := make([]*string, constants.Zero)

	return &CollectionPtr{
		items: &collection,
	}
}

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

func NewCollectionPtrUsingLength(length, capacity int) *CollectionPtr {
	collection := make([]*string, length, capacity)

	return &CollectionPtr{
		items: &collection,
	}
}

// --------- CharCollectionMap starts ----------

// NewCharCollectionMap CharCollectionMap.eachCollectionCapacity,
// capacity minimum 10 will be set if lower than 10 is given.
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

// EmptyCharCollectionMap eachCollectionCapacity = 0
func EmptyCharCollectionMap() *CharCollectionMap {
	mapElements := make(map[byte]*Collection, constants.Zero)

	return &CharCollectionMap{
		items:                  &mapElements,
		eachCollectionCapacity: defaultEachCollectionCapacity,
	}
}

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
		eachCollectionCapacity: constants.Zero,
	}

	charCollectionMap.AddStringsPtr(items)

	return charCollectionMap
}

func NewCharCollectionMapUsingItemsPlusCap(
	items *[]string,
	additionalCapacityOrLength int,
	eachCollectionCapacity int,
) *CharCollectionMap {
	isDefined := items != nil && *items != nil
	length := 0
	if items != nil && *items != nil {
		length = len(*items)
		additionalCapacityOrLength += length
	}

	mapElements := make(
		map[byte]*Collection,
		additionalCapacityOrLength)

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
	collection := make([]*Hashset, constants.Zero, constants.Zero)

	return &HashsetsCollection{
		items: &collection,
	}
}

func NewHashsetsCollection(
	hashsets *[]Hashset,
) *HashsetsCollection {
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

func NewHashsetsCollectionUsingPointerHashsets(
	hashsets *[]*Hashset,
) *HashsetsCollection {
	if hashsets == nil ||
		*hashsets == nil {
		return EmptyHashsetsCollection()
	}

	return &HashsetsCollection{
		items: hashsets,
	}
}

func NewHashsetsCollectionUsingLength(
	length, capacity int,
) *HashsetsCollection {
	collection := make([]*Hashset, length, capacity)

	return &HashsetsCollection{
		items: &collection,
	}
}

// --------- Hashmap starts ----------

func EmptyHashmap() *Hashmap {
	return NewHashmap(constants.Zero)
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

func NewHashmapUsingKeyAnyValues(
	keyAnyValues *[]KeyAnyValuePair,
) *Hashmap {
	if keyAnyValues == nil || *keyAnyValues == nil {
		return NewHashmap(defaultHashsetItems)
	}

	length := len(*keyAnyValues)
	hashMap := NewHashmap(length + constants.ArbitraryCapacity10)
	hashMap.AddOrUpdateKeyAnyValsPtr(keyAnyValues)

	return hashMap
}

func NewHashmapUsingKeyValues(
	keyValues *[]KeyValuePair,
) *Hashmap {
	if keyValues == nil || *keyValues == nil {
		return NewHashmap(defaultHashsetItems)
	}

	length := len(*keyValues)
	hashMap := NewHashmap(length + constants.ArbitraryCapacity10)
	hashMap.AddOrUpdateKeyValsPtr(keyValues)

	return hashMap
}

func NewHashmapUsingCollection(
	keys, values *Collection,
) *Hashmap {
	if keys == nil || keys.IsEmpty() {
		return EmptyHashmap()
	}

	itemsMap := converters.KeysValuesStringsToMap(
		keys.items,
		values.items)

	return NewHashmapUsingMap(
		itemsMap,
		constants.Zero,
		false)
}

func NewHashmapUsingStrings(
	keys, values *[]string,
) *Hashmap {
	if keys == nil || *keys == nil {
		return EmptyHashmap()
	}

	itemsMap := converters.KeysValuesStringsToMap(
		keys,
		values)

	return NewHashmapUsingMap(
		itemsMap,
		constants.Zero,
		false)
}

// NewHashmapUsingMap
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
		isEmptySet:    length == constants.Zero,
	}
}

// NewHashmapUsingMapUsingAddCapacity always returns the clone of the items.
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

// NewCharHashsetMap CharHashsetMap.eachHashsetCapacity, capacity minimum 10 will be set if lower than 10 is given.
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

func NewCharHashsetMapUsingItemsPlusCap(
	items *[]string,
	capacity, selfHashsetCapacity int,
) *CharHashsetMap {
	charHashsetMap := NewCharHashsetMap(
		capacity, selfHashsetCapacity)

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

// EmptyCharHashsetMap eachHashsetCapacity = 0
func EmptyCharHashsetMap() *CharHashsetMap {
	mapElements := make(
		map[byte]*Hashset,
		constants.Zero)

	return &CharHashsetMap{
		items:               &mapElements,
		eachHashsetCapacity: constants.Zero,
	}
}

// --------- LinkedCollections starts ----------

func NewLinkedCollections() *LinkedCollections {
	return &LinkedCollections{}
}

func EmptyLinkedCollections() *LinkedCollections {
	return &LinkedCollections{}
}

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
	collection := make([]*Collection, constants.Zero, capacity)

	return &CollectionsOfCollection{
		items: &collection,
	}
}

func EmptyCollectionsOfCollection() *CollectionsOfCollection {
	collection := make([]*Collection, constants.Zero)

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
		constants.Zero,
		length,
	).AddsStringsOfPointerStrings(isMakeClone, &stringItems)
}

func NewCollectionsOfCollectionUsingStringsOfPointerStrings(
	isMakeClone bool,
	stringItems *[]*[]string,
) *CollectionsOfCollection {
	length := LengthOfStringsOfPointerStrings(
		stringItems)

	return NewCollectionsOfCollectionUsingLength(
		constants.Zero,
		length,
	).AddsStringsOfPointerStrings(
		isMakeClone,
		stringItems)
}

func NewCollectionsOfCollectionUsingStrings(
	stringItems *[]string,
	isMakeClone bool,
) *CollectionsOfCollection {
	length := LengthOfStrings(
		stringItems)

	return NewCollectionsOfCollectionUsingLength(
		constants.Zero,
		length,
	).AddStringsPtr(
		stringItems,
		isMakeClone)
}

func NewCollectionsOfCollectionUsingStringsPlusCap(
	stringItems *[]string,
	capacity int,
	isMakeClone bool,
) *CollectionsOfCollection {
	length := LengthOfStrings(
		stringItems)
	collection := NewCollectionsOfCollection(
		length + capacity)

	return collection.AddStringsPtr(
		stringItems, isMakeClone)
}

func NewCollectionsOfCollectionUsingPointerStringsPlusCap(
	stringItems *[]*string,
	capacity int,
) *CollectionsOfCollection {
	length := LengthOfPointerStrings(
		stringItems)
	collection := NewCollectionsOfCollection(
		length + capacity)

	return collection.AddPointerStringsPtr(
		stringItems)
}

func NewCollectionsOfCollectionUsingLength(
	length,
	capacity int,
) *CollectionsOfCollection {
	collection := make(
		[]*Collection,
		length,
		capacity)

	return &CollectionsOfCollection{
		items: &collection,
	}
}

// --------- CollectionsOfCollectionPtr starts ----------

func NewCollectionsOfCollectionPtr(
	capacity int,
) *CollectionsOfCollectionPtr {
	collection := make(
		[]*CollectionPtr,
		constants.Zero,
		capacity)

	return &CollectionsOfCollectionPtr{
		items: &collection,
	}
}

func EmptyCollectionsOfCollectionPtr() *CollectionsOfCollectionPtr {
	collection := make([]*CollectionPtr, constants.Zero)

	return &CollectionsOfCollectionPtr{
		items: &collection,
	}
}

func NewCollectionsOfCollectionPtrUsingStringsOfStrings(
	stringItems ...*[]string,
) *CollectionsOfCollectionPtr {
	length := LengthOfStringsOfPointerStrings(
		&stringItems)

	return NewCollectionsOfCollectionPtrUsingLength(
		constants.Zero,
		length,
	).AddsStringsOfPointerStrings(
		constants.Zero,
		&stringItems)
}

func NewCollectionsOfCollectionPtrUsingStringsOfPointerStrings(
	stringItems *[]*[]string,
) *CollectionsOfCollectionPtr {
	length := LengthOfStringsOfPointerStrings(
		stringItems)

	return NewCollectionsOfCollectionPtrUsingLength(
		constants.Zero,
		length,
	).AddsStringsOfPointerStrings(
		constants.Zero,
		stringItems)
}

func NewCollectionsOfCollectionPtrUsingStrings(
	stringItems *[]string,
) *CollectionsOfCollectionPtr {
	length := LengthOfStrings(stringItems)

	return NewCollectionsOfCollectionPtrUsingLength(
		constants.Zero,
		length,
	).AddStringsPtr(stringItems, constants.Zero)
}

func NewCollectionsOfCollectionPtrUsingStringsPlusCap(
	stringItems *[]string,
	addCapacity int,
) *CollectionsOfCollectionPtr {
	length := LengthOfStrings(
		stringItems)
	collection := NewCollectionsOfCollectionPtr(
		length + addCapacity)

	return collection.AddStringsPtr(
		stringItems,
		constants.Zero)
}

func NewCollectionsOfCollectionPtrUsingPointerStringsPlusCap(
	stringItems *[]*string,
	capacity int,
) *CollectionsOfCollectionPtr {
	length := LengthOfPointerStrings(
		stringItems)
	collection := NewCollectionsOfCollectionPtr(
		length + capacity)

	return collection.AddPointerStringsPtr(
		stringItems)
}

func NewCollectionsOfCollectionPtrUsingLength(
	length,
	capacity int,
) *CollectionsOfCollectionPtr {
	collection := make(
		[]*CollectionPtr,
		length,
		capacity)

	return &CollectionsOfCollectionPtr{
		items: &collection,
	}
}
