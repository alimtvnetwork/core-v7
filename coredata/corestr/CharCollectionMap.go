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
	items                  map[byte]*Collection
	eachCollectionCapacity int
	sync.Mutex
}

func (it *CharCollectionMap) GetChar(
	str string,
) byte {
	if str != "" {
		return str[coreindexes.First]
	}

	return emptyChar
}

func (it *CharCollectionMap) GetCharOfPtr(
	str *string,
) byte {
	if str == nil || *str == "" {
		return emptyChar
	}

	return (*str)[coreindexes.First]
}

func (it *CharCollectionMap) GetCharsPtrGroups(
	items *[]string,
) *CharCollectionMap {
	if items == nil || *items == nil {
		return it
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

func (it *CharCollectionMap) GetMap() map[byte]*Collection {
	return it.items
}

// GetCopyMapLock Sends a copy of items
func (it *CharCollectionMap) GetCopyMapLock() map[byte]*Collection {
	it.Lock()
	defer it.Unlock()

	if it.IsEmpty() {
		return map[byte]*Collection{}
	}

	// TODO Fix copy logic
	return it.items
}

func (it *CharCollectionMap) SummaryStringLock() string {
	length := it.LengthLock()
	collectionOfCollection := make(
		[]string,
		length+1)

	collectionOfCollection[coreindexes.First] = fmt.Sprintf(
		summaryOfCharCollectionMapLengthFormat,
		it,
		length,
		coreindexes.First)

	i := 1
	for key, collection := range it.GetCopyMapLock() {
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

func (it *CharCollectionMap) SummaryString() string {
	collectionOfCollection := make(
		[]string,
		it.Length()+1)

	collectionOfCollection[coreindexes.First] = fmt.Sprintf(
		summaryOfCharCollectionMapLengthFormat,
		it,
		it.Length(),
		coreindexes.First+1)

	i := 1
	for key, collection := range it.items {
		collectionOfCollection[i] = fmt.Sprintf(
			charCollectionMapSingleItemFormat,
			i,
			string(key),
			collection.Length())

		i++
	}

	return strings.Join(
		collectionOfCollection,
		constants.EmptyString)
}

func (it *CharCollectionMap) String() string {
	collectionOfCollection := make(
		[]string,
		it.Length()*2+1)

	collectionOfCollection[coreindexes.First] =
		it.SummaryString()

	i := 1
	for key, collection := range it.items {
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

func (it *CharCollectionMap) SortedListAsc() *[]string {
	list := it.List()
	sort.Strings(*list)

	return list
}

func (it *CharCollectionMap) StringLock() string {
	collectionOfCollection := make(
		[]string,
		it.LengthLock()*2+1)

	collectionOfCollection[coreindexes.First] =
		it.SummaryStringLock()

	i := 1
	for key, collection := range it.GetCopyMapLock() {
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

func (it *CharCollectionMap) Print(isPrint bool) {
	if !isPrint {
		return
	}

	fmt.Println(
		it.String(),
	)
}

func (it *CharCollectionMap) PrintLock(isPrint bool) {
	if !isPrint {
		return
	}

	fmt.Println(
		it.StringLock(),
	)
}

func (it *CharCollectionMap) IsEmpty() bool {
	return it == nil ||
		len(it.items) == 0
}

func (it *CharCollectionMap) HasItems() bool {
	return !it.IsEmpty()
}

func (it *CharCollectionMap) IsEmptyLock() bool {
	it.Lock()
	defer it.Unlock()

	return it.IsEmpty()
}

// LengthOfCollectionFromFirstChar Get the char of the string given and get the length of how much is there.
func (it *CharCollectionMap) LengthOfCollectionFromFirstChar(
	str string,
) int {
	char := it.GetChar(str)

	collection, has := it.items[char]

	if has {
		return collection.Length()
	}

	return 0
}

func (it *CharCollectionMap) Has(
	str string,
) bool {
	if it.IsEmpty() {
		return false
	}

	char := it.
		GetChar(str)

	collection, has := it.items[char]

	if has {
		return collection.Has(str)
	}

	return false
}

func (it *CharCollectionMap) HasWithCollection(
	str string,
) (bool, *Collection) {
	if it.IsEmpty() {
		return false, EmptyCollection()
	}

	char := it.
		GetChar(str)

	collection, has := it.items[char]

	if has {
		return collection.Has(str), collection
	}

	return false, EmptyCollection()
}

func (it *CharCollectionMap) HasWithCollectionLock(
	str string,
) (bool, *Collection) {
	it.Lock()
	defer it.Unlock()

	if it.IsEmpty() {
		return false, EmptyCollection()
	}

	char := it.
		GetChar(str)

	collection, has := it.items[char]

	if has {
		return collection.HasLock(str), collection
	}

	return false, EmptyCollection()
}

func (it *CharCollectionMap) LengthOf(char byte) int {
	if it.IsEmpty() {
		return 0
	}

	collection, has := it.items[char]

	if has {
		return collection.Length()
	}

	return 0
}

func (it *CharCollectionMap) LengthOfLock(char byte) int {
	it.Lock()
	defer it.Unlock()

	if it.IsEmpty() {
		return 0
	}

	collection, has := it.items[char]

	if has {
		return collection.Length()
	}

	return 0
}

// AllLengthsSum All lengths sum.
func (it *CharCollectionMap) AllLengthsSum() int {
	if it == nil || it.items == nil {
		return 0
	}

	allLengthsSum := 0

	for _, collection := range it.items {
		allLengthsSum += collection.Length()
	}

	return allLengthsSum
}

// AllLengthsSumLock All lengths sum.
func (it *CharCollectionMap) AllLengthsSumLock() int {
	it.Lock()
	defer it.Unlock()

	if it == nil || it.items == nil {
		return 0
	}

	allLengthsSum := 0

	for _, collection := range it.items {
		allLengthsSum += collection.LengthLock()
	}

	return allLengthsSum
}

// Length Returns the length of chars which is the map length.
func (it *CharCollectionMap) Length() int {
	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

func (it *CharCollectionMap) LengthLock() int {
	it.Lock()
	defer it.Unlock()

	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

func (it *CharCollectionMap) IsEqualsPtrLock(
	another *CharCollectionMap,
) bool {
	it.Lock()
	defer it.Unlock()

	return it.IsEqualsWithCaseSensitivityPtr(
		another,
		true)
}

func (it *CharCollectionMap) IsEqualsPtr(
	another *CharCollectionMap,
) bool {
	return it.IsEqualsWithCaseSensitivityPtr(
		another,
		true)
}

func (it *CharCollectionMap) IsEqualsWithCaseSensitivityPtrLock(
	another *CharCollectionMap,
	isCaseSensitive bool,
) bool {
	it.Lock()
	defer it.Unlock()

	return it.IsEqualsWithCaseSensitivityPtr(
		another,
		isCaseSensitive)
}

func (it *CharCollectionMap) IsEqualsWithCaseSensitivityPtr(
	another *CharCollectionMap,
	isCaseSensitive bool,
) bool {
	if another == nil {
		return false
	}

	if another == it {
		return true
	}

	if another.IsEmpty() && it.IsEmpty() {
		return true
	}

	if another.IsEmpty() || it.IsEmpty() {
		return false
	}

	if another.Length() != it.Length() {
		return false
	}

	leftMap := it.items
	rightMap := another.items

	for key, collection := range leftMap {
		rCollection, has := rightMap[key]

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

func (it *CharCollectionMap) AddLock(
	str string,
) *CharCollectionMap {
	char := it.GetChar(str)

	it.Lock()
	collection, has := it.items[char]
	it.Unlock()

	if has {
		collection.AddLock(str)

		return it
	}

	newCollection := NewCollection(it.eachCollectionCapacity)
	newCollection.Add(str)

	it.Lock()
	it.items[char] = newCollection
	it.Unlock()

	return it
}

func (it *CharCollectionMap) Add(
	str string,
) *CharCollectionMap {
	char := it.GetChar(str)

	collection, has := it.items[char]

	if has {
		collection.Add(str)

		return it
	}

	newCollection := NewCollection(it.eachCollectionCapacity)
	newCollection.Add(str)
	it.items[char] = newCollection

	return it
}

func (it *CharCollectionMap) AddStringPtr(
	str *string,
) *CharCollectionMap {
	char := it.GetCharOfPtr(str)

	collection, has := it.items[char]

	if has {
		collection.AddPtr(str)

		return it
	}

	newCollection := NewCollection(it.eachCollectionCapacity)
	newCollection.AddPtr(str)
	it.items[char] = newCollection

	return it
}

func (it *CharCollectionMap) AddStringPtrLock(
	str *string,
) *CharCollectionMap {
	defer it.Unlock()
	char := it.GetCharOfPtr(str)

	it.Lock()
	collection, has := it.items[char]
	it.Unlock()

	if has {
		collection.AddPtrLock(str)

		return it
	}

	newCollection := NewCollection(it.eachCollectionCapacity)
	newCollection.AddPtr(str)

	it.Lock()
	it.items[char] = newCollection
	it.Unlock()

	return it
}

// AddSameStartingCharItems Assuming all items starts with same chars
func (it *CharCollectionMap) AddSameStartingCharItems(
	char byte,
	allItemsWithSameChar []string,
	isCloneAdd bool,
) *CharCollectionMap {
	if len(allItemsWithSameChar) == 0 {
		return it
	}

	values, has := it.items[char]

	if has {
		values.Adds(allItemsWithSameChar...)

		return it
	}

	it.items[char] =
		NewCollectionUsingStrings(
			isCloneAdd,
			allItemsWithSameChar,
		)

	return it
}

func (it *CharCollectionMap) AddPtrStringsLock(
	simpleStrings *[]*string,
) *CharCollectionMap {
	if simpleStrings == nil ||
		*simpleStrings == nil ||
		len(*simpleStrings) == 0 {
		return it
	}

	for _, item := range *simpleStrings {
		foundCollection := it.GetCollectionLock(
			*item, true)

		foundCollection.AddPtrLock(item)
	}

	return it
}

func (it *CharCollectionMap) AddHashmapsValues(
	hashmaps ...*Hashmap,
) *CharCollectionMap {
	if hashmaps == nil {
		return it
	}

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for _, v := range hashmap.items {
			vc := v
			it.AddStringPtr(&vc)
		}
	}

	return it
}

func (it *CharCollectionMap) AddHashmapsKeysOrValuesBothUsingFilter(
	filter IsKeyValueFilter,
	hashmaps ...*Hashmap,
) *CharCollectionMap {
	if hashmaps == nil {
		return it
	}

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k, v := range hashmap.items {
			result, isAccept, isBreak := filter(KeyValuePair{
				Key:   k,
				Value: v,
			})

			if isAccept {
				it.AddStringPtr(&result)
			}

			if isBreak {
				return it
			}
		}
	}

	return it
}

func (it *CharCollectionMap) AddHashmapsKeysValuesBoth(
	hashmaps ...*Hashmap,
) *CharCollectionMap {
	if hashmaps == nil {
		return it
	}

	for _, hashmap := range hashmaps {
		if hashmap.IsEmpty() {
			continue
		}

		for k, v := range hashmap.items {
			vc := v
			kc := k
			it.AddStringPtr(&vc)
			it.AddStringPtr(&kc)
		}
	}

	return it
}

func (it *CharCollectionMap) AddStringsPtrAsyncLock(
	largeStringsCollection *[]string,
	onComplete OnCompleteCharCollectionMap,
) *CharCollectionMap {
	if largeStringsCollection == nil ||
		*largeStringsCollection == nil {
		return it
	}

	length := len(*largeStringsCollection)

	if length == 0 {
		return it
	}

	isListIsTooLargeAndHasExistingData :=
		length > RegularCollectionEfficiencyLimit &&
			it.Length() > DoubleLimit

	if isListIsTooLargeAndHasExistingData {
		return it.
			efficientAddOfLargeItems(
				largeStringsCollection,
				onComplete)
	}

	wg := &sync.WaitGroup{}
	wg.Add(length)

	for _, item := range *largeStringsCollection {
		foundCollection := it.GetCollectionLock(
			item,
			true)

		go foundCollection.AddWithWgLock(
			item,
			wg,
		)
	}

	wg.Wait()

	if onComplete != nil {
		onComplete(it)
	}

	return it
}

func (it *CharCollectionMap) efficientAddOfLargeItems(
	largeStringsCollection *[]string,
	onComplete OnCompleteCharCollectionMap,
) *CharCollectionMap {
	allCharsMap := it.
		GetCharsPtrGroups(largeStringsCollection)

	wg := &sync.WaitGroup{}
	wg.Add(allCharsMap.Length())

	for key, collection := range allCharsMap.items {
		foundCollection := it.GetCollectionLock(
			string(key),
			true)

		go foundCollection.AddStringsPtrWgLock(
			&collection.items,
			wg,
		)
	}

	wg.Wait()

	if onComplete != nil {
		onComplete(it)
	}

	return it
}

func (it *CharCollectionMap) AddStringsPtr(
	items *[]string,
) *CharCollectionMap {
	if items == nil ||
		*items == nil ||
		len(*items) == 0 {
		return it
	}

	for _, item := range *items {
		itemC := item
		it.AddStringPtr(&itemC)
	}

	return it
}

func (it *CharCollectionMap) AddStrings(
	items ...string,
) *CharCollectionMap {
	if len(items) == 0 {
		return it
	}

	for i := range items {
		it.AddStringPtr(&(items)[i])
	}

	return it
}

func (it *CharCollectionMap) GetCollection(
	strFirstChar string,
	isAddNewOnEmpty bool,
) *Collection {
	char := it.GetChar(strFirstChar)

	collection, has := it.items[char]

	if has {
		return collection
	}

	if isAddNewOnEmpty {
		newCollection := NewCollection(it.eachCollectionCapacity)
		it.items[char] = newCollection

		return newCollection
	}

	return nil
}

func (it *CharCollectionMap) GetCollectionLock(
	strFirstChar string,
	isAddNewOnEmpty bool,
) *Collection {
	it.Lock()
	defer it.Unlock()

	return it.GetCollection(
		strFirstChar,
		isAddNewOnEmpty)
}

func (it *CharCollectionMap) AddSameCharsCollection(
	str string,
	stringsWithSameStartChar *Collection,
) *Collection {
	isNilOrEmptyCollectionGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundCollection := it.GetCollection(
		str,
		false)

	has := foundCollection != nil
	isAddToCollection := has && !isNilOrEmptyCollectionGiven
	hasCollectionHoweverNothingToAdd := has && isNilOrEmptyCollectionGiven

	if isAddToCollection {
		//goland:noinspection GoNilness
		foundCollection.AddStringsPtr(
			&stringsWithSameStartChar.items)

		return foundCollection
	} else if hasCollectionHoweverNothingToAdd {
		return foundCollection
	}

	char := it.GetChar(str)

	if isNilOrEmptyCollectionGiven {
		// create new
		newCollection := NewCollection(
			it.eachCollectionCapacity)
		it.items[char] = newCollection

		return newCollection
	}

	// items exist or stringsWithSameStartChar exists
	it.items[char] =
		stringsWithSameStartChar

	return stringsWithSameStartChar
}

func (it *CharCollectionMap) AddCollectionItems(
	collectionWithDiffStarts *Collection,
) *CharCollectionMap {
	if collectionWithDiffStarts == nil ||
		collectionWithDiffStarts.IsEmpty() {
		return it
	}

	it.AddStringsPtr(
		&collectionWithDiffStarts.items)

	return it
}

func (it *CharCollectionMap) AddCharHashsetMap(
	charHashsetMap *CharHashsetMap,
) *CharCollectionMap {
	if charHashsetMap.IsEmpty() {
		return it
	}

	for _, hashset := range charHashsetMap.items {
		for item := range hashset.items {
			it.Add(item)
		}
	}

	return it
}

func (it *CharCollectionMap) Resize(
	newLength int,
) *CharCollectionMap {
	currentLength := it.Length()

	if currentLength >= newLength {
		return it
	}

	newCollection := make(map[byte]*Collection, newLength)

	for key, element := range it.items {
		newCollection[key] = element
	}

	it.items = nil
	it.items = newCollection

	return it
}

func (it *CharCollectionMap) AddLength(
	lengths ...int,
) *CharCollectionMap {
	if len(lengths) == 0 {
		return it
	}

	currentLength := it.Length()

	for _, capacity := range lengths {
		currentLength += capacity
	}

	return it.Resize(currentLength)
}

func (it *CharCollectionMap) AddCollectionItemsAsyncLock(
	collectionWithDiffStarts *Collection,
	onComplete OnCompleteCharCollectionMap,
) *CharCollectionMap {
	if collectionWithDiffStarts == nil ||
		collectionWithDiffStarts.IsEmpty() {
		return it
	}

	go it.AddStringsPtrAsyncLock(
		&collectionWithDiffStarts.items,
		onComplete)

	return it
}

func (it *CharCollectionMap) List() *[]string {
	if it == nil ||
		it.IsEmpty() {
		return constants.EmptyStringsPtr
	}

	list := make([]string, it.AllLengthsSum())

	i := 0
	for _, collection := range it.items {
		for _, itemInList := range collection.items {
			list[i] = itemInList
			i++
		}
	}

	return &list
}

func (it *CharCollectionMap) ListLock() *[]string {
	it.Lock()
	defer it.Unlock()

	return it.List()
}

func (it *CharCollectionMap) AddSameCharsCollectionLock(
	str string,
	stringsWithSameStartChar *Collection,
) *Collection {
	isNilOrEmptyCollectionGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundCollection := it.GetCollectionLock(
		str,
		false)
	has := foundCollection != nil
	isAddToCollection := has && !isNilOrEmptyCollectionGiven
	hasCollectionHoweverNothingToAdd := has && isNilOrEmptyCollectionGiven

	if isAddToCollection {
		//goland:noinspection GoNilness
		foundCollection.AddStringsPtrLock(&stringsWithSameStartChar.items)

		return foundCollection
	} else if hasCollectionHoweverNothingToAdd {
		return foundCollection
	}

	char := it.GetChar(str)

	if isNilOrEmptyCollectionGiven {
		// create new
		newCollection := NewCollection(
			it.eachCollectionCapacity)

		it.Lock()

		it.items[char] = newCollection

		it.Unlock()

		return newCollection
	}

	// items exist or stringsWithSameStartChar exists
	it.Lock()
	it.items[char] =
		stringsWithSameStartChar
	it.Unlock()

	return stringsWithSameStartChar
}

func (it *CharCollectionMap) GetCollectionByChar(
	char byte,
) *Collection {
	return it.items[char]
}

func (it *CharCollectionMap) HashsetByChar(
	char byte,
) *Hashset {
	collection, has := it.items[char]

	if !has {
		return nil
	}

	return NewHashsetUsingCollection(
		collection)
}

func (it *CharCollectionMap) HashsetByCharLock(
	char byte,
) *Hashset {
	it.Lock()
	collection := it.items[char]
	it.Unlock()

	if collection == nil {
		return EmptyHashset()
	}

	items := collection.ListCopyPtrLock()

	return NewHashsetUsingStrings(
		items,
	)
}

func (it *CharCollectionMap) HashsetByStringFirstChar(
	str string,
) *Hashset {
	char := it.GetChar(str)

	return it.HashsetByChar(char)
}

func (it *CharCollectionMap) HashsetByStringFirstCharLock(
	str string,
) *Hashset {
	char := it.GetChar(str)

	return it.HashsetByCharLock(char)
}

func (it *CharCollectionMap) HashsetsCollectionByStringFirstChar(
	stringItems ...string,
) *HashsetsCollection {
	if it.IsEmpty() {
		return EmptyHashsetsCollection()
	}

	hashsets := make(
		[]*Hashset,
		0,
		it.Length())

	for _, item := range stringItems {
		char := it.GetChar(item)
		hashset := it.HashsetByChar(char)
		if hashset == nil || hashset.IsEmpty() {
			continue
		}

		hashsets = append(hashsets, hashset)
	}

	return NewHashsetsCollectionUsingPointerHashsets(hashsets...)
}

func (it *CharCollectionMap) HashsetsCollection() *HashsetsCollection {
	if it.IsEmpty() {
		return EmptyHashsetsCollection()
	}

	hashsets := make(
		[]*Hashset,
		0,
		it.Length())

	for _, collection := range it.items {
		if collection == nil ||
			collection.IsEmpty() {
			continue
		}

		hashset := collection.HashsetAsIs()
		hashsets = append(hashsets, hashset)
	}

	return NewHashsetsCollectionUsingPointerHashsets(hashsets...)
}

func (it *CharCollectionMap) HashsetsCollectionByChars(
	chars ...byte,
) *HashsetsCollection {
	if it.IsEmpty() {
		return EmptyHashsetsCollection()
	}

	hashsets := make(
		[]*Hashset,
		0,
		it.Length())

	for _, char := range chars {
		hashset := it.HashsetByChar(char)
		if hashset == nil ||
			hashset.IsEmpty() {
			continue
		}

		hashsets = append(hashsets, hashset)
	}

	return NewHashsetsCollectionUsingPointerHashsets(hashsets...)
}

func (it *CharCollectionMap) JsonModel() *CharCollectionDataModel {
	return &CharCollectionDataModel{
		Items: it.items,
		EachCollectionCapacity: it.
			eachCollectionCapacity,
	}
}

func (it *CharCollectionMap) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *CharCollectionMap) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *CharCollectionMap) AsJsoner() corejson.Jsoner {
	return it
}

func (it *CharCollectionMap) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it *CharCollectionMap) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *CharCollectionMap) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *CharCollectionMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(*it.JsonModel())
}

func (it *CharCollectionMap) UnmarshalJSON(data []byte) error {
	var dataModel CharCollectionDataModel

	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.items = dataModel.Items
		it.eachCollectionCapacity =
			dataModel.EachCollectionCapacity
	}

	return err
}

func (it CharCollectionMap) Json() corejson.Result {
	return corejson.NewFromAny(it)
}

func (it CharCollectionMap) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
}

func (it *CharCollectionMap) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*CharCollectionMap, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return EmptyCharCollectionMap(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *CharCollectionMap) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *CharCollectionMap {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

// Clear clears existing items, deletes items using delete(*charCollectionMap.items, char)
func (it *CharCollectionMap) Clear() *CharCollectionMap {
	if it.IsEmpty() {
		return it
	}

	for char, values := range it.items {
		values.Dispose()
		values = nil

		delete(it.items, char)
	}

	return it
}

func (it *CharCollectionMap) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.items = nil
}
