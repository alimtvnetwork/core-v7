package corestr

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreindexes"
)

type CharHashsetMap struct {
	items               map[byte]*Hashset
	eachHashsetCapacity int
	sync.Mutex
}

func (it *CharHashsetMap) GetChar(
	str string,
) byte {
	if str != "" {
		return str[coreindexes.First]
	}

	return emptyChar
}

func (it *CharHashsetMap) HashsetsCollectionByChars(
	chars ...byte,
) *HashsetsCollection {
	if it.IsEmpty() {
		return Empty.HashsetsCollection()
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

	return New.HashsetsCollection.UsingHashsetsPointers(hashsets...)
}

func (it *CharHashsetMap) HashsetsCollectionByStringsFirstChar(
	stringItems ...string,
) *HashsetsCollection {
	if it.IsEmpty() {
		return Empty.HashsetsCollection()
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

	return New.HashsetsCollection.UsingHashsetsPointers(hashsets...)
}

func (it *CharHashsetMap) HashsetsCollection() *HashsetsCollection {
	if it.IsEmpty() {
		return Empty.HashsetsCollection()
	}

	hashsets := make(
		[]Hashset,
		0,
		it.Length())

	for _, hashset := range it.items {
		//goland:noinspection ALL
		hashsets = append(hashsets, *hashset)
	}

	return New.HashsetsCollection.UsingHashsets(hashsets...)
}

func (it *CharHashsetMap) GetCharOfPtr(
	str *string,
) byte {
	if str == nil || *str == "" {
		return emptyChar
	}

	return (*str)[coreindexes.First]
}

func (it *CharHashsetMap) GetCharsPtrGroups(
	items *[]string,
) *CharHashsetMap {
	if items == nil || *items == nil {
		return it
	}

	length := len(*items)

	if length == 0 {
		return nil
	}

	hashsetMap := New.CharHashsetMap.Cap(
		length,
		length/3)

	return hashsetMap.AddStringsPtr(items)
}

func (it *CharHashsetMap) GetMap() map[byte]*Hashset {
	return it.items
}

// GetCopyMapLock Sends a copy of items
func (it *CharHashsetMap) GetCopyMapLock() map[byte]*Hashset {
	it.Lock()
	defer it.Unlock()

	if it.IsEmpty() {
		return map[byte]*Hashset{}
	}

	// todo fix copying
	return it.items
}

func (it *CharHashsetMap) SummaryStringLock() string {
	length := it.LengthLock()
	hashsetOfHashset := make(
		[]string,
		length+1)

	hashsetOfHashset[coreindexes.First] = fmt.Sprintf(
		summaryOfCharHashsetMapLengthFormat,
		it,
		length,
		coreindexes.First)

	i := 1
	for key, hashset := range it.GetCopyMapLock() {
		hashsetOfHashset[i] = fmt.Sprintf(
			charHashsetMapSingleItemFormat,
			i,
			string(key),
			hashset.LengthLock())

		i++
	}

	return strings.Join(
		hashsetOfHashset,
		constants.EmptyString)
}

func (it *CharHashsetMap) SummaryString() string {
	hashsetOfHashset := make(
		[]string,
		it.Length()+1)

	hashsetOfHashset[coreindexes.First] = fmt.Sprintf(
		summaryOfCharHashsetMapLengthFormat,
		it,
		it.Length(),
		coreindexes.First)

	i := 1
	for key, hashset := range it.items {
		hashsetOfHashset[i] = fmt.Sprintf(
			charHashsetMapSingleItemFormat,
			i,
			string(key),
			hashset.Length())

		i++
	}

	return strings.Join(
		hashsetOfHashset,
		constants.EmptyString)
}

func (it *CharHashsetMap) String() string {
	hashsetOfHashset := make(
		[]string,
		it.Length()*2+1)

	hashsetOfHashset[coreindexes.First] =
		it.SummaryString()

	i := 1
	for key, hashset := range it.items {
		hashsetOfHashset[i] = fmt.Sprintf(
			charHashsetMapLengthFormat,
			string(key))

		i++
		hashsetOfHashset[i] = hashset.String()
		i++
	}

	return strings.Join(
		hashsetOfHashset,
		constants.EmptyString)
}

func (it *CharHashsetMap) StringLock() string {
	hashsetOfHashset := make(
		[]string,
		it.LengthLock()*2+1)

	hashsetOfHashset[coreindexes.First] =
		it.SummaryStringLock()

	i := 1
	for key, hashset := range it.GetCopyMapLock() {

		hashsetOfHashset[i] = fmt.Sprintf(
			charHashsetMapLengthFormat,
			string(key))

		i++

		hashsetOfHashset[i] = hashset.StringLock()
		i++
	}

	return strings.Join(
		hashsetOfHashset,
		constants.EmptyString)
}

func (it *CharHashsetMap) List() *[]string {
	list := make([]string, it.AllLengthsSum())

	i := 0
	for _, hashset := range it.items {
		for s := range hashset.items {
			list[i] = s
			i++
		}
	}

	return &list
}

func (it *CharHashsetMap) SortedListAsc() *[]string {
	list := it.List()
	sort.Strings(*list)

	return list
}

func (it *CharHashsetMap) SortedListDsc() *[]string {
	list := it.SortedListAsc()
	length := len(*list)
	mid := length / 2

	for i := 0; i < mid; i++ {
		temp := (*list)[i]
		(*list)[i] = (*list)[length-1-i]
		(*list)[length-1-i] = temp
	}

	return list
}

func (it *CharHashsetMap) Print(isPrint bool) {
	if !isPrint {
		return
	}

	fmt.Println(
		it.String(),
	)
}

func (it *CharHashsetMap) PrintLock(isPrint bool) {
	if !isPrint {
		return
	}

	fmt.Println(
		it.StringLock(),
	)
}

func (it *CharHashsetMap) IsEmpty() bool {
	return it == nil ||
		len(it.items) == 0
}

func (it *CharHashsetMap) HasItems() bool {
	return it != nil && len(it.items) > 0
}

func (it *CharHashsetMap) IsEmptyLock() bool {
	it.Lock()
	defer it.Unlock()

	return it.IsEmpty()
}

// LengthOfHashsetFromFirstChar Get the char of the string given and get the length of how much is there.
func (it *CharHashsetMap) LengthOfHashsetFromFirstChar(
	str string,
) int {
	char := it.GetChar(str)

	hashset, has := it.items[char]

	if has {
		return hashset.Length()
	}

	return 0
}

func (it *CharHashsetMap) Has(
	str string,
) bool {
	if it.IsEmpty() {
		return false
	}

	char := it.
		GetChar(str)

	hashset, has := it.items[char]

	if has {
		return hashset.Has(str)
	}

	return false
}

func (it *CharHashsetMap) HasWithHashset(
	str string,
) (bool, *Hashset) {
	if it.IsEmpty() {
		return false, New.Hashset.Empty()
	}

	char := it.
		GetChar(str)

	hashset, has := it.items[char]

	if has {
		return hashset.Has(str), hashset
	}

	return false, New.Hashset.Empty()
}

func (it *CharHashsetMap) HasWithHashsetLock(
	str string,
) (bool, *Hashset) {
	it.Lock()
	defer it.Unlock()

	if it.IsEmpty() {
		return false, New.Hashset.Empty()
	}

	char := it.
		GetChar(str)

	hashset, has := it.items[char]

	if has {
		return hashset.HasLock(str), hashset
	}

	return false, New.Hashset.Empty()
}

func (it *CharHashsetMap) LengthOf(char byte) int {
	if it.IsEmpty() {
		return 0
	}

	hashset, has := it.items[char]

	if has {
		return hashset.Length()
	}

	return 0
}

func (it *CharHashsetMap) LengthOfLock(char byte) int {
	it.Lock()
	defer it.Unlock()

	if it.IsEmpty() {
		return 0
	}

	hashset, has := it.items[char]

	if has {
		return hashset.Length()
	}

	return 0
}

// AllLengthsSum All lengths sum.
func (it *CharHashsetMap) AllLengthsSum() int {
	if it.IsEmpty() {
		return 0
	}

	allLengthsSum := 0

	for _, hashset := range it.items {
		allLengthsSum += hashset.Length()
	}

	return allLengthsSum
}

// AllLengthsSumLock All lengths sum.
func (it *CharHashsetMap) AllLengthsSumLock() int {
	it.Lock()
	defer it.Unlock()

	if it.IsEmpty() {
		return 0
	}

	allLengthsSum := 0

	for _, hashset := range it.items {
		allLengthsSum += hashset.LengthLock()
	}

	return allLengthsSum
}

func (it *CharHashsetMap) AddCharCollectionMapItems(
	charCollectionMap *CharCollectionMap,
) *CharHashsetMap {
	if charCollectionMap == nil ||
		charCollectionMap.IsEmpty() {
		return it
	}

	it.AddStringsPtr(
		charCollectionMap.List())

	return it
}

func (it *CharHashsetMap) AddCollectionItems(
	collectionWithDiffStarts *Collection,
) *CharHashsetMap {
	if collectionWithDiffStarts == nil ||
		collectionWithDiffStarts.IsEmpty() {
		return it
	}

	it.AddStrings(
		collectionWithDiffStarts.items...)

	return it
}

func (it *CharHashsetMap) AddCollectionItemsAsyncLock(
	collectionWithDiffStarts *Collection,
	onComplete OnCompleteCharHashsetMap,
) *CharHashsetMap {
	if collectionWithDiffStarts == nil ||
		collectionWithDiffStarts.IsEmpty() {
		return it
	}

	go it.AddStringsPtrAsyncLock(
		&collectionWithDiffStarts.items,
		onComplete)

	return it
}

func (it *CharHashsetMap) Length() int {
	if it.IsEmpty() {
		return 0
	}

	return len(it.items)
}

func (it *CharHashsetMap) LengthLock() int {
	it.Lock()
	defer it.Unlock()

	if it.IsEmpty() {
		return 0
	}

	return len(it.items)
}

func (it *CharHashsetMap) IsEqualsPtrLock(
	another *CharHashsetMap,
) bool {
	it.Lock()
	defer it.Unlock()

	return it.IsEqualsPtr(
		another)
}

func (it *CharHashsetMap) IsEqualsPtr(
	another *CharHashsetMap,
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

	for key, hashset := range leftMap {
		rHashset, has := rightMap[key]

		if !has {
			return false
		}

		if !rHashset.IsEqualsPtr(hashset) {
			return false
		}
	}

	return true
}

func (it *CharHashsetMap) AddLock(
	str string,
) *CharHashsetMap {
	char := it.GetChar(str)

	it.Lock()
	hashset, has := it.items[char]
	it.Unlock()

	if has {
		hashset.AddLock(str)

		return it
	}

	newHashset := New.Hashset.Cap(it.eachHashsetCapacity)
	newHashset.Add(str)

	it.Lock()
	it.items[char] = newHashset
	it.Unlock()

	return it
}

func (it *CharHashsetMap) Add(
	str string,
) *CharHashsetMap {
	char := it.GetChar(str)

	hashset, has := it.items[char]

	if has {
		hashset.Add(str)

		return it
	}

	newHashset := New.Hashset.Cap(it.eachHashsetCapacity)
	newHashset.Add(str)
	it.items[char] = newHashset

	return it
}

func (it *CharHashsetMap) AddStringPtr(
	str *string,
) *CharHashsetMap {
	char := it.GetCharOfPtr(str)

	hashset, has := it.items[char]

	if has {
		hashset.AddPtr(str)

		return it
	}

	newHashset := New.Hashset.Cap(it.eachHashsetCapacity)
	newHashset.AddPtr(str)
	it.items[char] = newHashset

	return it
}

func (it *CharHashsetMap) AddStringPtrLock(
	str *string,
) *CharHashsetMap {
	defer it.Unlock()
	char := it.GetCharOfPtr(str)

	it.Lock()
	hashset, has := it.items[char]
	it.Unlock()

	if has {
		hashset.AddPtrLock(str)

		return it
	}

	newHashset := New.Hashset.Cap(it.eachHashsetCapacity)
	newHashset.AddPtr(str)

	it.Lock()
	it.items[char] = newHashset
	it.Unlock()

	return it
}

// AddSameStartingCharItems Assuming all items starts with same chars
func (it *CharHashsetMap) AddSameStartingCharItems(
	char byte,
	allItemsWithSameChar *[]string,
) *CharHashsetMap {
	if allItemsWithSameChar == nil ||
		*allItemsWithSameChar == nil {
		return it
	}

	length := len(*allItemsWithSameChar)

	if length == 0 {
		return it
	}

	values, has := it.items[char]

	if has {
		values.AddStringsPtr(allItemsWithSameChar)

		return it
	}

	it.items[char] =
		New.Hashset.StringsPtr(
			allItemsWithSameChar)

	return it
}

func (it *CharHashsetMap) AddPtrStringsLock(
	simpleStrings *[]*string,
) *CharHashsetMap {
	if simpleStrings == nil ||
		*simpleStrings == nil ||
		len(*simpleStrings) == 0 {
		return it
	}

	for _, item := range *simpleStrings {
		foundHashset := it.GetHashsetLock(
			*item, true)

		foundHashset.AddPtrLock(item)
	}

	return it
}

func (it *CharHashsetMap) AddStringsPtrAsyncLock(
	largeStringsHashset *[]string,
	onComplete OnCompleteCharHashsetMap,
) *CharHashsetMap {
	if largeStringsHashset == nil ||
		*largeStringsHashset == nil {
		return it
	}

	length := len(*largeStringsHashset)

	if length == 0 {
		return it
	}

	isListIsTooLargeAndHasExistingData :=
		length > RegularCollectionEfficiencyLimit &&
			it.Length() > DoubleLimit

	if isListIsTooLargeAndHasExistingData {
		return it.
			efficientAddOfLargeItems(
				largeStringsHashset,
				onComplete)
	}

	wg := &sync.WaitGroup{}
	wg.Add(length)

	for _, item := range *largeStringsHashset {
		foundHashset := it.GetHashsetLock(
			item,
			true)

		go foundHashset.AddWithWgLock(
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

func (it *CharHashsetMap) efficientAddOfLargeItems(
	largeStringsHashset *[]string,
	onComplete OnCompleteCharHashsetMap,
) *CharHashsetMap {
	allCharsMap := it.
		GetCharsPtrGroups(largeStringsHashset)

	wg := &sync.WaitGroup{}
	wg.Add(allCharsMap.Length())

	for key, hashset := range allCharsMap.items {
		foundHashset := it.GetHashsetLock(
			string(key),
			true)

		go foundHashset.AddHashsetWgLock(
			hashset,
			wg,
		)
	}

	wg.Wait()

	if onComplete != nil {
		onComplete(it)
	}

	return it
}

func (it *CharHashsetMap) AddStringsPtr(
	items *[]string,
) *CharHashsetMap {
	if items == nil ||
		*items == nil ||
		len(*items) == 0 {
		return it
	}

	for _, item := range *items {
		it.AddStringPtr(&item)
	}

	return it
}

func (it *CharHashsetMap) AddStrings(
	items ...string,
) *CharHashsetMap {
	if items == nil ||
		len(items) == 0 {
		return it
	}

	for _, item := range items {
		it.AddStringPtr(&item)
	}

	return it
}

func (it *CharHashsetMap) GetHashset(
	strFirstChar string,
	isAddNewOnEmpty bool,
) *Hashset {
	char := it.GetChar(strFirstChar)

	hashset, has := it.items[char]

	if has {
		return hashset
	}

	if isAddNewOnEmpty {
		newHashset := New.Hashset.Cap(it.eachHashsetCapacity)
		it.items[char] = newHashset

		return newHashset
	}

	return nil
}

func (it *CharHashsetMap) GetHashsetLock(
	strFirstChar string,
	isAddNewOnEmpty bool,
) *Hashset {
	it.Lock()
	defer it.Unlock()

	return it.GetHashset(
		strFirstChar,
		isAddNewOnEmpty)
}

func (it *CharHashsetMap) AddSameCharsCollection(
	str string,
	stringsWithSameStartChar *Collection,
) *Hashset {
	isNilOrEmptyHashsetGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundHashset := it.GetHashset(
		str,
		false)

	has := foundHashset != nil
	isAddToHashset := has && !isNilOrEmptyHashsetGiven
	hasHashsetHoweverNothingToAdd := has && isNilOrEmptyHashsetGiven

	if isAddToHashset {
		foundHashset.AddCollection(stringsWithSameStartChar)

		return foundHashset
	} else if hasHashsetHoweverNothingToAdd {
		return foundHashset
	}

	char := it.GetChar(str)

	if isNilOrEmptyHashsetGiven {
		// create new
		newHashset := New.Hashset.Cap(
			it.eachHashsetCapacity)
		it.items[char] = newHashset

		return newHashset
	}

	// items exist or stringsWithSameStartChar exists
	//goland:noinspection GoNilness
	toHashset := stringsWithSameStartChar.HashsetAsIs()
	it.items[char] = toHashset

	return toHashset
}

func (it *CharHashsetMap) AddSameCharsHashset(
	str string,
	stringsWithSameStartChar *Hashset,
) *Hashset {
	isNilOrEmptyHashsetGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundHashset := it.GetHashset(
		str,
		false)

	has := foundHashset != nil
	isAddToHashset := has && !isNilOrEmptyHashsetGiven
	hasHashsetHoweverNothingToAdd := has && isNilOrEmptyHashsetGiven

	if isAddToHashset {
		foundHashset.AddHashsetItems(stringsWithSameStartChar)

		return foundHashset
	} else if hasHashsetHoweverNothingToAdd {
		return foundHashset
	}

	char := it.GetChar(str)

	if isNilOrEmptyHashsetGiven {
		// create new
		newHashset := New.Hashset.Cap(
			it.eachHashsetCapacity)
		it.items[char] = newHashset

		return newHashset
	}

	// items exist or stringsWithSameStartChar exists
	it.items[char] =
		stringsWithSameStartChar

	return stringsWithSameStartChar
}

func (it *CharHashsetMap) AddHashsetItems(
	hashsetWithDiffStarts *Hashset,
) *CharHashsetMap {
	if hashsetWithDiffStarts == nil ||
		hashsetWithDiffStarts.IsEmpty() {
		return it
	}

	it.AddStringsPtr(
		hashsetWithDiffStarts.ListPtr())

	return it
}

func (it *CharHashsetMap) AddHashsetItemsAsyncLock(
	hashsetWithDiffStarts *Hashset,
	onComplete OnCompleteCharHashsetMap,
) *CharHashsetMap {
	if hashsetWithDiffStarts == nil ||
		hashsetWithDiffStarts.IsEmpty() {
		return it
	}

	go it.AddStringsPtrAsyncLock(
		hashsetWithDiffStarts.ListCopyPtrLock(),
		onComplete)

	return it
}

func (it *CharHashsetMap) AddSameCharsCollectionLock(
	str string,
	stringsWithSameStartChar *Collection,
) *Hashset {
	isNilOrEmptyHashsetGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundHashset := it.GetHashsetLock(
		str,
		false)
	has := foundHashset != nil
	isAddToHashset := has &&
		!isNilOrEmptyHashsetGiven
	hasHashsetHoweverNothingToAdd := has &&
		isNilOrEmptyHashsetGiven

	if isAddToHashset {
		list := stringsWithSameStartChar.
			ListCopyPtrLock()

		foundHashset.
			AddStringsPtrLock(&list)

		return foundHashset
	} else if hasHashsetHoweverNothingToAdd {
		return foundHashset
	}

	char := it.GetChar(str)

	if isNilOrEmptyHashsetGiven {
		// create new
		newHashset := New.Hashset.Cap(
			it.eachHashsetCapacity)
		it.Lock()
		it.items[char] = newHashset
		it.Unlock()

		return newHashset
	}

	// items exist or stringsWithSameStartChar exists
	//goland:noinspection GoNilness
	hashset := stringsWithSameStartChar.HashsetAsIs()
	//goland:noinspection GoLinterLocal
	it.Lock()
	it.items[char] =
		hashset
	it.Unlock()

	return hashset
}

func (it *CharHashsetMap) AddHashsetLock(
	str string,
	stringsWithSameStartChar *Hashset,
) *Hashset {
	isNilOrEmptyHashsetGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundHashset := it.GetHashsetLock(
		str,
		false)
	has := foundHashset != nil
	isAddToHashset := has && !isNilOrEmptyHashsetGiven
	hasHashsetHoweverNothingToAdd := has && isNilOrEmptyHashsetGiven

	if isAddToHashset {
		//goland:noinspection GoNilness
		foundHashset.AddStringsPtrLock(
			stringsWithSameStartChar.ListPtr())

		return foundHashset
	} else if hasHashsetHoweverNothingToAdd {
		return foundHashset
	}

	// current str char, no lock required
	char := it.GetChar(str)

	if isNilOrEmptyHashsetGiven {
		// create new
		newHashset := New.Hashset.Cap(
			it.eachHashsetCapacity)
		it.Lock()
		it.items[char] = newHashset
		it.Unlock()

		return newHashset
	}

	// items exist or stringsWithSameStartChar exists
	it.Lock()
	it.items[char] =
		stringsWithSameStartChar
	it.Unlock()

	return stringsWithSameStartChar
}

func (it *CharHashsetMap) GetHashsetByChar(
	char byte,
) *Hashset {
	return it.items[char]
}

func (it *CharHashsetMap) HashsetByChar(
	char byte,
) *Hashset {
	hashset := it.items[char]

	return hashset
}

func (it *CharHashsetMap) HashsetByCharLock(
	char byte,
) *Hashset {
	it.Lock()
	hashset := it.items[char]
	it.Unlock()

	if hashset == nil {
		return New.Hashset.Empty()
	}

	return hashset
}

func (it *CharHashsetMap) HashsetByStringFirstChar(
	str string,
) *Hashset {
	char := it.GetChar(str)

	return it.HashsetByChar(char)
}

func (it *CharHashsetMap) HashsetByStringFirstCharLock(
	str string,
) *Hashset {
	char := it.GetChar(str)

	return it.HashsetByCharLock(char)
}

func (it *CharHashsetMap) JsonModel() *CharHashsetDataModel {
	return &CharHashsetDataModel{
		Items: it.items,
		EachHashsetCapacity: it.
			eachHashsetCapacity,
	}
}

func (it *CharHashsetMap) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *CharHashsetMap) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *CharHashsetMap) AsJsoner() corejson.Jsoner {
	return it
}

func (it *CharHashsetMap) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it *CharHashsetMap) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *CharHashsetMap) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *CharHashsetMap) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*CharHashsetMap, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return Empty.CharHashsetMap(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *CharHashsetMap) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *CharHashsetMap {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *CharHashsetMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(*it.JsonModel())
}

func (it *CharHashsetMap) UnmarshalJSON(data []byte) error {
	var dataModel CharHashsetDataModel

	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.items = dataModel.Items
		it.eachHashsetCapacity =
			dataModel.EachHashsetCapacity
	}

	return err
}

func (it CharHashsetMap) Json() corejson.Result {
	return corejson.New(it)
}

func (it CharHashsetMap) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

// RemoveAll remove all existing items, deletes items using delete(*charCollectionMap.items, char), expensive operation
func (it *CharHashsetMap) RemoveAll() *CharHashsetMap {
	if it.IsEmpty() {
		return it
	}

	return it.Clear()
}

// Clear points to a new map and collects old pointer and remove all elements from pointer in separate goroutine.
func (it *CharHashsetMap) Clear() *CharHashsetMap {
	if it.IsEmpty() {
		return it
	}

	tempCollection := it.items
	it.items = nil
	it.items = make(map[byte]*Hashset, 0)

	go func() {
		for char, values := range tempCollection {
			values.Dispose()
			values = nil

			delete(tempCollection, char)
		}
	}()

	return it
}
