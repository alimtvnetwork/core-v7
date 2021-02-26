package corestr

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreindexes"
)

type CharCollectionMap struct {
	items                  *map[byte]*Collection
	eachCollectionCapacity int
	sync.Mutex
}

func (charCollectionMap *CharCollectionMap) GetChar(
	str string,
) byte {
	if str != "" {
		return str[coreindexes.First]
	}

	return emptyChar
}

func (charCollectionMap *CharCollectionMap) GetCharOfPtr(
	str *string,
) byte {
	if str == nil || *str == "" {
		return emptyChar
	}

	return (*str)[coreindexes.First]
}

func (charCollectionMap *CharCollectionMap) GetCharsPtrGroups(
	items *[]string,
) *CharCollectionMap {
	if items == nil || *items == nil {
		return charCollectionMap
	}

	length := len(*items)
	lenBy4 := length / 3

	if lenBy4 < defaultEachCollectionCapacity {
		lenBy4 = defaultEachCollectionCapacity
	}

	if length == 0 {
		return nil
	}

	collectionMap := NewCharCollectionMap(
		length,
		length/3)

	return collectionMap.AddStringsPtr(items)
}

func (charCollectionMap *CharCollectionMap) GetMap() *map[byte]*Collection {
	return charCollectionMap.items
}

// Sends a copy of items
func (charCollectionMap *CharCollectionMap) GetCopyMapLock() *map[byte]*Collection {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	if charCollectionMap.IsEmpty() {
		return &(map[byte]*Collection{})
	}

	return &(*charCollectionMap.items)
}

func (charCollectionMap *CharCollectionMap) SummaryStringLock() string {
	length := charCollectionMap.LengthLock()
	collectionOfCollection := make(
		[]string,
		length+1)

	collectionOfCollection[coreindexes.First] = fmt.Sprintf(
		summaryOfCharCollectionMapLengthFormat,
		charCollectionMap,
		length,
		coreindexes.First)

	i := 1
	for key, collection := range *charCollectionMap.GetCopyMapLock() {
		collectionOfCollection[i] = fmt.Sprintf(
			charCollectionMapSingleItemFormat,
			i+1,
			string(key),
			collection.LengthLock())

		i++
	}

	return strings.Join(
		collectionOfCollection,
		constants.EmptyString)
}

func (charCollectionMap *CharCollectionMap) SummaryString() string {
	collectionOfCollection := make(
		[]string,
		charCollectionMap.Length()+1)

	collectionOfCollection[coreindexes.First] = fmt.Sprintf(
		summaryOfCharCollectionMapLengthFormat,
		charCollectionMap,
		charCollectionMap.Length(),
		coreindexes.First+1)

	i := 1
	for key, collection := range *charCollectionMap.items {
		collectionOfCollection[i] = fmt.Sprintf(
			charCollectionMapSingleItemFormat,
			i+1,
			string(key),
			collection.Length())

		i++
	}

	return strings.Join(
		collectionOfCollection,
		constants.EmptyString)
}

func (charCollectionMap *CharCollectionMap) String() string {
	collectionOfCollection := make(
		[]string,
		charCollectionMap.Length()*2+1)

	collectionOfCollection[coreindexes.First] =
		charCollectionMap.SummaryString()

	i := 1
	for key, collection := range *charCollectionMap.items {
		collectionOfCollection[i] = fmt.Sprintf(
			charCollectionMapLengthFormat,
			string(key))

		i++
		collectionOfCollection[i] = collection.String()
		i++
	}

	return strings.Join(
		collectionOfCollection,
		constants.EmptyString)
}

func (charCollectionMap *CharCollectionMap) SortedListAsc() *[]string {
	list := charCollectionMap.List()
	sort.Strings(*list)

	return list
}

func (charCollectionMap *CharCollectionMap) StringLock() string {
	collectionOfCollection := make(
		[]string,
		charCollectionMap.LengthLock()*2+1)

	collectionOfCollection[coreindexes.First] =
		charCollectionMap.SummaryStringLock()

	i := 1
	for key, collection := range *charCollectionMap.GetCopyMapLock() {
		collectionOfCollection[i] = fmt.Sprintf(
			charCollectionMapLengthFormat,
			string(key))

		i++
		collectionOfCollection[i] =
			collection.StringLock()
		i++
	}

	return strings.Join(
		collectionOfCollection,
		constants.EmptyString)
}

func (charCollectionMap *CharCollectionMap) Print(isPrint bool) {
	if !isPrint {
		return
	}

	fmt.Println(
		charCollectionMap.String(),
	)
}

func (charCollectionMap *CharCollectionMap) PrintLock(isPrint bool) {
	if !isPrint {
		return
	}

	fmt.Println(
		charCollectionMap.StringLock(),
	)
}

func (charCollectionMap *CharCollectionMap) IsEmpty() bool {
	return charCollectionMap.items == nil ||
		*charCollectionMap.items == nil ||
		len(*charCollectionMap.items) == 0
}

func (charCollectionMap *CharCollectionMap) HasItems() bool {
	return charCollectionMap.items != nil &&
		*charCollectionMap.items != nil &&
		len(*charCollectionMap.items) > 0
}

func (charCollectionMap *CharCollectionMap) IsEmptyLock() bool {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	return charCollectionMap.
		items == nil ||
		*charCollectionMap.items == nil ||
		len(*charCollectionMap.items) == 0
}

// Get the char of the string given and get the length of how much is there.
func (charCollectionMap *CharCollectionMap) LengthOfCollectionFromFirstChar(
	str string,
) int {
	char := charCollectionMap.GetChar(str)

	collection, has := (*charCollectionMap.items)[char]

	if has {
		return collection.Length()
	}

	return 0
}

func (charCollectionMap *CharCollectionMap) Has(
	str string,
) bool {
	if charCollectionMap.IsEmpty() {
		return false
	}

	char := charCollectionMap.
		GetChar(str)

	collection, has := (*charCollectionMap.items)[char]

	if has {
		return collection.Has(str)
	}

	return false
}

func (charCollectionMap *CharCollectionMap) HasWithCollection(
	str string,
) (bool, *Collection) {
	if charCollectionMap.IsEmpty() {
		return false, EmptyCollection()
	}

	char := charCollectionMap.
		GetChar(str)

	collection, has := (*charCollectionMap.items)[char]

	if has {
		return collection.Has(str), collection
	}

	return false, EmptyCollection()
}

func (charCollectionMap *CharCollectionMap) HasWithCollectionLock(
	str string,
) (bool, *Collection) {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	if charCollectionMap.IsEmpty() {
		return false, EmptyCollection()
	}

	char := charCollectionMap.
		GetChar(str)

	collection, has := (*charCollectionMap.items)[char]

	if has {
		return collection.HasLock(str), collection
	}

	return false, EmptyCollection()
}

func (charCollectionMap *CharCollectionMap) LengthOf(char byte) int {
	if charCollectionMap.IsEmpty() {
		return 0
	}

	collection, has := (*charCollectionMap.items)[char]

	if has {
		return collection.Length()
	}

	return 0
}

func (charCollectionMap *CharCollectionMap) LengthOfLock(char byte) int {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	if charCollectionMap.IsEmpty() {
		return 0
	}

	collection, has := (*charCollectionMap.items)[char]

	if has {
		return collection.Length()
	}

	return 0
}

// All lengths sum.
func (charCollectionMap *CharCollectionMap) AllLengthsSum() int {
	if charCollectionMap.
		items == nil ||
		*charCollectionMap.items == nil {
		return 0
	}

	allLengthsSum := 0

	for _, collection := range *charCollectionMap.items {
		allLengthsSum += collection.Length()
	}

	return allLengthsSum
}

// All lengths sum.
func (charCollectionMap *CharCollectionMap) AllLengthsSumLock() int {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	if charCollectionMap.
		items == nil ||
		*charCollectionMap.items == nil {
		return 0
	}

	allLengthsSum := 0

	for _, collection := range *charCollectionMap.items {
		allLengthsSum += collection.LengthLock()
	}

	return allLengthsSum
}

// Returns the length of chars which is the map length.
func (charCollectionMap *CharCollectionMap) Length() int {
	if charCollectionMap.
		items == nil ||
		*charCollectionMap.items == nil {
		return 0
	}

	return len(*charCollectionMap.items)
}

func (charCollectionMap *CharCollectionMap) LengthLock() int {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	if charCollectionMap.
		items == nil ||
		*charCollectionMap.items == nil {
		return 0
	}

	return len(*charCollectionMap.items)
}

func (charCollectionMap *CharCollectionMap) IsEqualsPtrLock(
	another *CharCollectionMap,
) bool {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	return charCollectionMap.IsEqualsWithCaseSensitivityPtr(
		another,
		true)
}

func (charCollectionMap *CharCollectionMap) IsEqualsPtr(
	another *CharCollectionMap,
) bool {
	return charCollectionMap.IsEqualsWithCaseSensitivityPtr(
		another,
		true)
}

func (charCollectionMap *CharCollectionMap) IsEqualsWithCaseSensitivityPtrLock(
	another *CharCollectionMap,
	isCaseSensitive bool,
) bool {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	return charCollectionMap.IsEqualsWithCaseSensitivityPtr(
		another,
		isCaseSensitive)
}

func (charCollectionMap *CharCollectionMap) IsEqualsWithCaseSensitivityPtr(
	another *CharCollectionMap,
	isCaseSensitive bool,
) bool {
	if another == nil {
		return false
	}

	if another == charCollectionMap {
		return true
	}

	if another.IsEmpty() && charCollectionMap.IsEmpty() {
		return true
	}

	if another.IsEmpty() || charCollectionMap.IsEmpty() {
		return false
	}

	if another.Length() != charCollectionMap.Length() {
		return false
	}

	leftMap := charCollectionMap.items
	rightMap := another.items

	for key, collection := range *leftMap {
		rCollection, has := (*rightMap)[key]

		if !has {
			return false
		}

		if !rCollection.IsEqualsWithSensitivePtr(
			collection,
			isCaseSensitive) {
			return false
		}
	}

	return true
}

func (charCollectionMap *CharCollectionMap) AddLock(
	str string,
) *CharCollectionMap {
	char := charCollectionMap.GetChar(str)

	charCollectionMap.Lock()
	collection, has := (*charCollectionMap.
		items)[char]
	charCollectionMap.Unlock()

	if has {
		collection.AddLock(str)

		return charCollectionMap
	}

	newCollection := NewCollection(charCollectionMap.eachCollectionCapacity)
	newCollection.Add(str)

	charCollectionMap.Lock()
	(*charCollectionMap.items)[char] = newCollection
	charCollectionMap.Unlock()

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) Add(
	str string,
) *CharCollectionMap {
	char := charCollectionMap.GetChar(str)

	collection, has := (*charCollectionMap.
		items)[char]

	if has {
		collection.Add(str)
	}

	newCollection := NewCollection(charCollectionMap.eachCollectionCapacity)
	newCollection.Add(str)
	(*charCollectionMap.
		items)[char] = newCollection

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddStringPtr(
	str *string,
) *CharCollectionMap {
	char := charCollectionMap.GetCharOfPtr(str)

	collection, has := (*charCollectionMap.
		items)[char]

	if has {
		collection.AddPtr(str)
	}

	newCollection := NewCollection(charCollectionMap.eachCollectionCapacity)
	newCollection.AddPtr(str)
	(*charCollectionMap.
		items)[char] = newCollection

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddStringPtrLock(
	str *string,
) *CharCollectionMap {
	defer charCollectionMap.Unlock()
	char := charCollectionMap.GetCharOfPtr(str)

	charCollectionMap.Lock()
	collection, has := (*charCollectionMap.
		items)[char]
	charCollectionMap.Unlock()

	if has {
		collection.AddPtrLock(str)

		return charCollectionMap
	}

	newCollection := NewCollection(charCollectionMap.eachCollectionCapacity)
	newCollection.AddPtr(str)

	charCollectionMap.Lock()
	(*charCollectionMap.
		items)[char] = newCollection
	charCollectionMap.Unlock()

	return charCollectionMap
}

// Assuming all items starts with same chars
func (charCollectionMap *CharCollectionMap) AddSameStartingCharItems(
	char byte,
	allItemsWithSameChar *[]string,
	isCloneAdd bool,
) *CharCollectionMap {
	if allItemsWithSameChar == nil ||
		*allItemsWithSameChar == nil ||
		len(*allItemsWithSameChar) == 0 {
		return charCollectionMap
	}

	values, has := (*charCollectionMap.
		items)[char]

	if has {
		values.AddStringsPtr(allItemsWithSameChar)

		return charCollectionMap
	}

	(*charCollectionMap.
		items)[char] =
		NewCollectionUsingStrings(allItemsWithSameChar, isCloneAdd)

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddPtrStringsLock(
	simpleStrings *[]*string,
) *CharCollectionMap {
	if simpleStrings == nil ||
		*simpleStrings == nil ||
		len(*simpleStrings) == 0 {
		return charCollectionMap
	}

	for _, item := range *simpleStrings {
		foundCollection := charCollectionMap.GetCollectionLock(
			*item, true)

		foundCollection.AddPtrLock(item)
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddHashmapsValues(
	hashmaps ...*Hashmap,
) *CharCollectionMap {
	if hashmaps == nil {
		return charCollectionMap
	}

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for _, v := range *hashmap.items {
			vc := v
			charCollectionMap.AddStringPtr(&vc)
		}
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddHashmapsKeysOrValuesBothUsingFilter(
	filter IsKeyValueFilter,
	hashmaps ...*Hashmap,
) *CharCollectionMap {
	if hashmaps == nil {
		return charCollectionMap
	}

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k, v := range *hashmap.items {
			result, isAccept, isBreak := filter(KeyValuePair{
				Key:   k,
				Value: v,
			})

			if isAccept {
				charCollectionMap.AddStringPtr(&result)
			}

			if isBreak {
				return charCollectionMap
			}
		}
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddHashmapsKeysValuesBoth(
	hashmaps ...*Hashmap,
) *CharCollectionMap {
	if hashmaps == nil {
		return charCollectionMap
	}

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k, v := range *hashmap.items {
			vc := v
			kc := k
			charCollectionMap.AddStringPtr(&vc)
			charCollectionMap.AddStringPtr(&kc)
		}
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddStringsPtrAsyncLock(
	largeStringsCollection *[]string,
	onComplete OnCompleteCharCollectionMap,
) *CharCollectionMap {
	if largeStringsCollection == nil ||
		*largeStringsCollection == nil {
		return charCollectionMap
	}

	length := len(*largeStringsCollection)

	if length == 0 {
		return charCollectionMap
	}

	isListIsTooLargeAndHasExistingData := length > RegularCollectionEfficiencyLimit &&
		charCollectionMap.Length() > DoubleLimit

	if isListIsTooLargeAndHasExistingData {
		return charCollectionMap.
			efficientAddOfLargeItems(
				largeStringsCollection,
				onComplete)
	}

	wg := &sync.WaitGroup{}
	wg.Add(length)

	for _, item := range *largeStringsCollection {
		foundCollection := charCollectionMap.GetCollectionLock(
			item,
			true)

		go foundCollection.AddWithWgLock(
			item,
			wg,
		)
	}

	wg.Wait()

	if onComplete != nil {
		onComplete(charCollectionMap)
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) efficientAddOfLargeItems(
	largeStringsCollection *[]string, onComplete OnCompleteCharCollectionMap,
) *CharCollectionMap {
	allCharsMap := charCollectionMap.
		GetCharsPtrGroups(largeStringsCollection)

	wg := &sync.WaitGroup{}
	wg.Add(allCharsMap.Length())

	for key, collection := range *allCharsMap.items {
		foundCollection := charCollectionMap.GetCollectionLock(
			string(key),
			true)

		go foundCollection.AddStringsPtrWgLock(
			collection.items,
			wg,
		)
	}

	wg.Wait()

	if onComplete != nil {
		onComplete(charCollectionMap)
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddStringsPtr(
	items *[]string,
) *CharCollectionMap {
	if items == nil ||
		*items == nil ||
		len(*items) == 0 {
		return charCollectionMap
	}

	for _, item := range *items {
		itemC := item
		charCollectionMap.AddStringPtr(&itemC)
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddStrings(
	items ...string,
) *CharCollectionMap {
	if len(items) == 0 {
		return charCollectionMap
	}

	for i := range items {
		charCollectionMap.AddStringPtr(&(items)[i])
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) GetCollection(
	strFirstChar string,
	isAddNewOnEmpty bool,
) *Collection {
	char := charCollectionMap.GetChar(strFirstChar)

	collection, has := (*charCollectionMap.items)[char]

	if has {
		return collection
	}

	if isAddNewOnEmpty {
		newCollection := NewCollection(charCollectionMap.eachCollectionCapacity)
		(*charCollectionMap.items)[char] = newCollection

		return newCollection
	}

	return nil
}

func (charCollectionMap *CharCollectionMap) GetCollectionLock(
	strFirstChar string,
	isAddNewOnEmpty bool,
) *Collection {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	return charCollectionMap.GetCollection(
		strFirstChar,
		isAddNewOnEmpty)
}

func (charCollectionMap *CharCollectionMap) AddSameCharsCollection(
	str string,
	stringsWithSameStartChar *Collection,
) *Collection {
	isNilOrEmptyCollectionGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundCollection := charCollectionMap.GetCollection(
		str,
		false)

	has := foundCollection != nil
	isAddToCollection := has && !isNilOrEmptyCollectionGiven
	hasCollectionHoweverNothingToAdd := has && isNilOrEmptyCollectionGiven

	if isAddToCollection {
		//goland:noinspection GoNilness
		foundCollection.AddStringsPtr(stringsWithSameStartChar.items)

		return foundCollection
	} else if hasCollectionHoweverNothingToAdd {
		return foundCollection
	}

	char := charCollectionMap.GetChar(str)

	if isNilOrEmptyCollectionGiven {
		// create new
		newCollection := NewCollection(
			charCollectionMap.eachCollectionCapacity)
		(*charCollectionMap.items)[char] = newCollection

		return newCollection
	}

	// items exist or stringsWithSameStartChar exists
	(*charCollectionMap.items)[char] =
		stringsWithSameStartChar

	return stringsWithSameStartChar
}

func (charCollectionMap *CharCollectionMap) AddCollectionItems(
	collectionWithDiffStarts *Collection,
) *CharCollectionMap {
	if collectionWithDiffStarts == nil ||
		collectionWithDiffStarts.IsEmpty() {
		return charCollectionMap
	}

	charCollectionMap.AddStringsPtr(
		collectionWithDiffStarts.items)

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddCharHashsetMap(
	charHashsetMap *CharHashsetMap,
) *CharCollectionMap {
	if charHashsetMap == nil ||
		charHashsetMap.IsEmpty() {
		return charCollectionMap
	}

	for _, hashset := range *charHashsetMap.items {
		for item := range *hashset.items {
			charCollectionMap.Add(item)
		}
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) Resize(
	newLength int,
) *CharCollectionMap {
	currentLength := charCollectionMap.Length()

	if currentLength >= newLength {
		return charCollectionMap
	}

	newCollection := make(map[byte]*Collection, newLength)

	for key, element := range *charCollectionMap.items {
		newCollection[key] = element
	}

	charCollectionMap.items = nil
	charCollectionMap.items = &newCollection

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddLength(
	lengths ...int,
) *CharCollectionMap {
	if len(lengths) == 0 {
		return charCollectionMap
	}

	currentLength := charCollectionMap.Length()

	for _, capacity := range lengths {
		currentLength += capacity
	}

	return charCollectionMap.Resize(currentLength)
}

func (charCollectionMap *CharCollectionMap) AddCollectionItemsAsyncLock(
	collectionWithDiffStarts *Collection,
	onComplete OnCompleteCharCollectionMap,
) *CharCollectionMap {
	if collectionWithDiffStarts == nil ||
		collectionWithDiffStarts.IsEmpty() {
		return charCollectionMap
	}

	go charCollectionMap.AddStringsPtrAsyncLock(
		collectionWithDiffStarts.items,
		onComplete)

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) List() *[]string {
	if charCollectionMap == nil ||
		charCollectionMap.IsEmpty() {
		return constants.EmptyStringsPtr
	}

	list := make([]string, charCollectionMap.AllLengthsSum())

	i := 0
	for _, collection := range *charCollectionMap.items {

		for _, itemInList := range *collection.items {
			list[i] = itemInList
			i++
		}
	}

	return &list
}

func (charCollectionMap *CharCollectionMap) ListLock() *[]string {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	return charCollectionMap.List()
}

func (charCollectionMap *CharCollectionMap) AddSameCharsCollectionLock(
	str string,
	stringsWithSameStartChar *Collection,
) *Collection {
	isNilOrEmptyCollectionGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundCollection := charCollectionMap.GetCollectionLock(
		str,
		false)
	has := foundCollection != nil
	isAddToCollection := has && !isNilOrEmptyCollectionGiven
	hasCollectionHoweverNothingToAdd := has && isNilOrEmptyCollectionGiven

	if isAddToCollection {
		//goland:noinspection GoNilness
		foundCollection.AddStringsPtr(stringsWithSameStartChar.items)

		return foundCollection
	} else if hasCollectionHoweverNothingToAdd {
		return foundCollection
	}

	char := charCollectionMap.GetChar(str)

	if isNilOrEmptyCollectionGiven {
		// create new
		newCollection := NewCollection(
			charCollectionMap.eachCollectionCapacity)

		charCollectionMap.Lock()

		(*charCollectionMap.items)[char] = newCollection

		charCollectionMap.Unlock()

		return newCollection
	}

	// items exist or stringsWithSameStartChar exists
	charCollectionMap.Lock()
	(*charCollectionMap.items)[char] =
		stringsWithSameStartChar
	charCollectionMap.Unlock()

	return stringsWithSameStartChar
}

func (charCollectionMap *CharCollectionMap) GetCollectionByChar(
	char byte,
) *Collection {
	return (*charCollectionMap.items)[char]
}

func (charCollectionMap *CharCollectionMap) HashsetByChar(
	char byte,
) *Hashset {
	collection, has := (*charCollectionMap.items)[char]

	if !has {
		return nil
	}

	return NewHashsetUsingCollection(
		collection,
		0,
		false)
}

func (charCollectionMap *CharCollectionMap) HashsetByCharLock(
	char byte,
) *Hashset {
	charCollectionMap.Lock()
	collection := (*charCollectionMap.items)[char]
	charCollectionMap.Unlock()

	if collection == nil {
		return EmptyHashset()
	}

	items := collection.ListCopyPtrLock()

	return NewHashsetUsingStrings(
		items,
		0,
		false)
}

func (charCollectionMap *CharCollectionMap) HashsetByStringFirstChar(
	str string,
) *Hashset {
	char := charCollectionMap.GetChar(str)

	return charCollectionMap.HashsetByChar(char)
}

func (charCollectionMap *CharCollectionMap) HashsetByStringFirstCharLock(
	str string,
) *Hashset {
	char := charCollectionMap.GetChar(str)

	return charCollectionMap.HashsetByCharLock(char)
}

func (charCollectionMap *CharCollectionMap) HashsetsCollectionByStringFirstChar(
	stringItems ...string,
) *HashsetsCollection {
	if charCollectionMap.IsEmpty() {
		return EmptyHashsetsCollection()
	}

	hashsets := make(
		[]*Hashset,
		0,
		charCollectionMap.Length())

	for _, item := range stringItems {
		char := charCollectionMap.GetChar(item)
		hashset := charCollectionMap.HashsetByChar(char)
		if hashset == nil || hashset.IsEmpty() {
			continue
		}

		hashsets = append(hashsets, hashset)
	}

	return NewHashsetsCollectionUsingPointerHashsets(&hashsets)
}

func (charCollectionMap *CharCollectionMap) HashsetsCollection() *HashsetsCollection {
	if charCollectionMap.IsEmpty() {
		return EmptyHashsetsCollection()
	}

	hashsets := make(
		[]*Hashset,
		0,
		charCollectionMap.Length())

	for _, collection := range *charCollectionMap.items {
		if collection == nil ||
			collection.IsEmpty() {
			continue
		}

		hashset := collection.HashsetAsIs()
		hashsets = append(hashsets, hashset)
	}

	return NewHashsetsCollectionUsingPointerHashsets(&hashsets)
}

func (charCollectionMap *CharCollectionMap) HashsetsCollectionByChars(
	chars ...byte,
) *HashsetsCollection {
	if charCollectionMap.IsEmpty() {
		return EmptyHashsetsCollection()
	}

	hashsets := make(
		[]*Hashset,
		0,
		charCollectionMap.Length())

	for _, char := range chars {
		hashset := charCollectionMap.HashsetByChar(char)
		if hashset == nil ||
			hashset.IsEmpty() {
			continue
		}

		hashsets = append(hashsets, hashset)
	}

	return NewHashsetsCollectionUsingPointerHashsets(&hashsets)
}

func (charCollectionMap *CharCollectionMap) JsonModel() *CharCollectionDataModel {
	return &CharCollectionDataModel{
		Items: charCollectionMap.items,
		EachCollectionCapacity: charCollectionMap.
			eachCollectionCapacity,
	}
}

func (charCollectionMap *CharCollectionMap) JsonModelAny() interface{} {
	return charCollectionMap.JsonModel()
}

func (charCollectionMap *CharCollectionMap) AsJsoner() *corejson.Jsoner {
	var jsoner corejson.Jsoner = charCollectionMap

	return &jsoner
}

func (charCollectionMap *CharCollectionMap) AsJsonMarshaller() *corejson.JsonMarshaller {
	var jsonMarshaller corejson.JsonMarshaller = charCollectionMap

	return &jsonMarshaller
}

func (charCollectionMap *CharCollectionMap) AsJsonParseSelfInjector() *corejson.ParseSelfInjector {
	var jsonMarshaller corejson.ParseSelfInjector = charCollectionMap

	return &jsonMarshaller
}

func (charCollectionMap *CharCollectionMap) JsonParseSelfInject(jsonResult *corejson.Result) {
	charCollectionMap.ParseInjectUsingJsonMust(jsonResult)
}

func (charCollectionMap *CharCollectionMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(*charCollectionMap.JsonModel())
}

func (charCollectionMap *CharCollectionMap) UnmarshalJSON(data []byte) error {
	var dataModel CharCollectionDataModel

	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		charCollectionMap.items = dataModel.Items
		charCollectionMap.eachCollectionCapacity =
			dataModel.EachCollectionCapacity
	}

	return err
}

func (charCollectionMap *CharCollectionMap) Json() *corejson.Result {
	if charCollectionMap.IsEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(charCollectionMap.JsonModel())

	return corejson.NewPtr(jsonBytes, err)
}

func (charCollectionMap *CharCollectionMap) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*CharCollectionMap, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return EmptyCharCollectionMap(), nil
	}

	err := json.Unmarshal(*jsonResult.Bytes, &charCollectionMap)

	if err != nil {
		return EmptyCharCollectionMap(), err
	}

	return charCollectionMap, nil
}

// Panic if error
func (charCollectionMap *CharCollectionMap) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *CharCollectionMap {
	newUsingJson, err :=
		charCollectionMap.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

// clears existing items, deletes items using delete(*charCollectionMap.items, char)
func (charCollectionMap *CharCollectionMap) Clear() *CharCollectionMap {
	if charCollectionMap.IsEmpty() {
		return charCollectionMap
	}

	for char := range *charCollectionMap.items {
		delete(*charCollectionMap.items, char)
	}

	return charCollectionMap
}
