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
	"gitlab.com/evatix-go/core/defaulterr"
)

type CharHashsetMap struct {
	items               *map[byte]*Hashset
	eachHashsetCapacity int
	sync.Mutex
}

func (charHashsetMap *CharHashsetMap) GetChar(
	str string,
) byte {
	if str != "" {
		return str[coreindexes.First]
	}

	return emptyChar
}

func (charHashsetMap *CharHashsetMap) HashsetsCollectionByChars(
	chars ...byte,
) *HashsetsCollection {
	if charHashsetMap.IsEmpty() {
		return EmptyHashsetsCollection()
	}

	hashsets := make(
		[]*Hashset,
		0,
		charHashsetMap.Length())

	for _, char := range chars {
		hashset := charHashsetMap.HashsetByChar(char)
		if hashset == nil ||
			hashset.IsEmpty() {
			continue
		}

		hashsets = append(hashsets, hashset)
	}

	return NewHashsetsCollectionUsingPointerHashsets(&hashsets)
}

func (charHashsetMap *CharHashsetMap) HashsetsCollectionByStringsFirstChar(
	stringItems ...string,
) *HashsetsCollection {
	if charHashsetMap.IsEmpty() {
		return EmptyHashsetsCollection()
	}

	hashsets := make(
		[]*Hashset,
		0,
		charHashsetMap.Length())

	for _, item := range stringItems {
		char := charHashsetMap.GetChar(item)
		hashset := charHashsetMap.HashsetByChar(char)

		if hashset == nil || hashset.IsEmpty() {
			continue
		}

		hashsets = append(hashsets, hashset)
	}

	return NewHashsetsCollectionUsingPointerHashsets(&hashsets)
}

func (charHashsetMap *CharHashsetMap) HashsetsCollection() *HashsetsCollection {
	if charHashsetMap.IsEmpty() {
		return EmptyHashsetsCollection()
	}

	hashsets := make(
		[]Hashset,
		0,
		charHashsetMap.Length())

	for _, hashset := range *charHashsetMap.items {
		//goland:noinspection ALL
		hashsets = append(hashsets, *hashset)
	}

	return NewHashsetsCollection(&hashsets)
}

func (charHashsetMap *CharHashsetMap) GetCharOfPtr(
	str *string,
) byte {
	if str == nil || *str == "" {
		return emptyChar
	}

	return (*str)[coreindexes.First]
}

func (charHashsetMap *CharHashsetMap) GetCharsPtrGroups(
	items *[]string,
) *CharHashsetMap {
	if items == nil || *items == nil {
		return charHashsetMap
	}

	length := len(*items)

	if length == 0 {
		return nil
	}

	hashsetMap := NewCharHashsetMap(
		length,
		length/3)

	return hashsetMap.AddStringsPtr(items)
}

func (charHashsetMap *CharHashsetMap) GetMap() *map[byte]*Hashset {
	return charHashsetMap.items
}

// GetCopyMapLock Sends a copy of items
func (charHashsetMap *CharHashsetMap) GetCopyMapLock() *map[byte]*Hashset {
	charHashsetMap.Lock()
	defer charHashsetMap.Unlock()

	if charHashsetMap.IsEmpty() {
		return &(map[byte]*Hashset{})
	}

	return &(*charHashsetMap.items)
}

func (charHashsetMap *CharHashsetMap) SummaryStringLock() string {
	length := charHashsetMap.LengthLock()
	hashsetOfHashset := make(
		[]string,
		length+1)

	hashsetOfHashset[coreindexes.First] = fmt.Sprintf(
		summaryOfCharHashsetMapLengthFormat,
		charHashsetMap,
		length,
		coreindexes.First)

	i := 1
	for key, hashset := range *charHashsetMap.GetCopyMapLock() {
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

func (charHashsetMap *CharHashsetMap) SummaryString() string {
	hashsetOfHashset := make(
		[]string,
		charHashsetMap.Length()+1)

	hashsetOfHashset[coreindexes.First] = fmt.Sprintf(
		summaryOfCharHashsetMapLengthFormat,
		charHashsetMap,
		charHashsetMap.Length(),
		coreindexes.First)

	i := 1
	for key, hashset := range *charHashsetMap.items {
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

func (charHashsetMap *CharHashsetMap) String() string {
	hashsetOfHashset := make(
		[]string,
		charHashsetMap.Length()*2+1)

	hashsetOfHashset[coreindexes.First] =
		charHashsetMap.SummaryString()

	i := 1
	for key, hashset := range *charHashsetMap.items {
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

func (charHashsetMap *CharHashsetMap) StringLock() string {
	hashsetOfHashset := make(
		[]string,
		charHashsetMap.LengthLock()*2+1)

	hashsetOfHashset[coreindexes.First] =
		charHashsetMap.SummaryStringLock()

	i := 1
	for key, hashset := range *charHashsetMap.GetCopyMapLock() {

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

func (charHashsetMap *CharHashsetMap) List() *[]string {
	list := make([]string, charHashsetMap.AllLengthsSum())

	i := 0
	for _, hashset := range *charHashsetMap.items {
		for s := range *hashset.items {
			list[i] = s
			i++
		}
	}

	return &list
}

func (charHashsetMap *CharHashsetMap) SortedListAsc() *[]string {
	list := charHashsetMap.List()
	sort.Strings(*list)

	return list
}

func (charHashsetMap *CharHashsetMap) SortedListDsc() *[]string {
	list := charHashsetMap.SortedListAsc()
	length := len(*list)
	mid := length / 2

	for i := 0; i < mid; i++ {
		temp := (*list)[i]
		(*list)[i] = (*list)[length-1-i]
		(*list)[length-1-i] = temp
	}

	return list
}

func (charHashsetMap *CharHashsetMap) Print(isPrint bool) {
	if !isPrint {
		return
	}

	fmt.Println(
		charHashsetMap.String(),
	)
}

func (charHashsetMap *CharHashsetMap) PrintLock(isPrint bool) {
	if !isPrint {
		return
	}

	fmt.Println(
		charHashsetMap.StringLock(),
	)
}

func (charHashsetMap *CharHashsetMap) IsEmpty() bool {
	return charHashsetMap.items == nil ||
		*charHashsetMap.items == nil ||
		len(*charHashsetMap.items) == 0
}

func (charHashsetMap *CharHashsetMap) HasItems() bool {
	return charHashsetMap.items != nil &&
		*charHashsetMap.items != nil &&
		len(*charHashsetMap.items) > 0
}

func (charHashsetMap *CharHashsetMap) IsEmptyLock() bool {
	charHashsetMap.Lock()
	defer charHashsetMap.Unlock()

	return charHashsetMap.
		items == nil ||
		*charHashsetMap.items == nil ||
		len(*charHashsetMap.items) == 0
}

// LengthOfHashsetFromFirstChar Get the char of the string given and get the length of how much is there.
func (charHashsetMap *CharHashsetMap) LengthOfHashsetFromFirstChar(
	str string,
) int {
	char := charHashsetMap.GetChar(str)

	hashset, has := (*charHashsetMap.items)[char]

	if has {
		return hashset.Length()
	}

	return 0
}

func (charHashsetMap *CharHashsetMap) Has(
	str string,
) bool {
	if charHashsetMap.IsEmpty() {
		return false
	}

	char := charHashsetMap.
		GetChar(str)

	hashset, has := (*charHashsetMap.items)[char]

	if has {
		return hashset.Has(str)
	}

	return false
}

func (charHashsetMap *CharHashsetMap) HasWithHashset(
	str string,
) (bool, *Hashset) {
	if charHashsetMap.IsEmpty() {
		return false, EmptyHashset()
	}

	char := charHashsetMap.
		GetChar(str)

	hashset, has := (*charHashsetMap.items)[char]

	if has {
		return hashset.Has(str), hashset
	}

	return false, EmptyHashset()
}

func (charHashsetMap *CharHashsetMap) HasWithHashsetLock(
	str string,
) (bool, *Hashset) {
	charHashsetMap.Lock()
	defer charHashsetMap.Unlock()

	if charHashsetMap.IsEmpty() {
		return false, EmptyHashset()
	}

	char := charHashsetMap.
		GetChar(str)

	hashset, has := (*charHashsetMap.items)[char]

	if has {
		return hashset.HasLock(str), hashset
	}

	return false, EmptyHashset()
}

func (charHashsetMap *CharHashsetMap) LengthOf(char byte) int {
	if charHashsetMap.IsEmpty() {
		return 0
	}

	hashset, has := (*charHashsetMap.items)[char]

	if has {
		return hashset.Length()
	}

	return 0
}

func (charHashsetMap *CharHashsetMap) LengthOfLock(char byte) int {
	charHashsetMap.Lock()
	defer charHashsetMap.Unlock()

	if charHashsetMap.IsEmpty() {
		return 0
	}

	hashset, has := (*charHashsetMap.items)[char]

	if has {
		return hashset.Length()
	}

	return 0
}

// AllLengthsSum All lengths sum.
func (charHashsetMap *CharHashsetMap) AllLengthsSum() int {
	if charHashsetMap.
		items == nil ||
		*charHashsetMap.items == nil {
		return 0
	}

	allLengthsSum := 0

	for _, hashset := range *charHashsetMap.items {
		allLengthsSum += hashset.Length()
	}

	return allLengthsSum
}

// AllLengthsSumLock All lengths sum.
func (charHashsetMap *CharHashsetMap) AllLengthsSumLock() int {
	charHashsetMap.Lock()
	defer charHashsetMap.Unlock()

	if charHashsetMap.
		items == nil ||
		*charHashsetMap.items == nil {
		return 0
	}

	allLengthsSum := 0

	for _, hashset := range *charHashsetMap.items {
		allLengthsSum += hashset.LengthLock()
	}

	return allLengthsSum
}

func (charHashsetMap *CharHashsetMap) AddCharCollectionMapItems(
	charCollectionMap *CharCollectionMap,
) *CharHashsetMap {
	if charCollectionMap == nil ||
		charCollectionMap.IsEmpty() {
		return charHashsetMap
	}

	charHashsetMap.AddStringsPtr(
		charCollectionMap.List())

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) AddCollectionItems(
	collectionWithDiffStarts *Collection,
) *CharHashsetMap {
	if collectionWithDiffStarts == nil ||
		collectionWithDiffStarts.IsEmpty() {
		return charHashsetMap
	}

	charHashsetMap.AddStringsPtr(
		collectionWithDiffStarts.items)

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) AddCollectionItemsAsyncLock(
	collectionWithDiffStarts *Collection,
	onComplete OnCompleteCharHashsetMap,
) *CharHashsetMap {
	if collectionWithDiffStarts == nil ||
		collectionWithDiffStarts.IsEmpty() {
		return charHashsetMap
	}

	go charHashsetMap.AddStringsPtrAsyncLock(
		collectionWithDiffStarts.items,
		onComplete)

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) Length() int {
	if charHashsetMap.
		items == nil ||
		*charHashsetMap.items == nil {
		return 0
	}

	return len(*charHashsetMap.items)
}

func (charHashsetMap *CharHashsetMap) LengthLock() int {
	charHashsetMap.Lock()
	defer charHashsetMap.Unlock()

	if charHashsetMap.
		items == nil ||
		*charHashsetMap.items == nil {
		return 0
	}

	return len(*charHashsetMap.items)
}

func (charHashsetMap *CharHashsetMap) IsEqualsPtrLock(
	another *CharHashsetMap,
) bool {
	charHashsetMap.Lock()
	defer charHashsetMap.Unlock()

	return charHashsetMap.IsEqualsPtr(
		another)
}

func (charHashsetMap *CharHashsetMap) IsEqualsPtr(
	another *CharHashsetMap,
) bool {
	if another == nil {
		return false
	}

	if another == charHashsetMap {
		return true
	}

	if another.IsEmpty() && charHashsetMap.IsEmpty() {
		return true
	}

	if another.IsEmpty() || charHashsetMap.IsEmpty() {
		return false
	}

	if another.Length() != charHashsetMap.Length() {
		return false
	}

	leftMap := charHashsetMap.items
	rightMap := another.items

	for key, hashset := range *leftMap {
		rHashset, has := (*rightMap)[key]

		if !has {
			return false
		}

		if !rHashset.IsEqualsPtr(hashset) {
			return false
		}
	}

	return true
}

func (charHashsetMap *CharHashsetMap) AddLock(
	str string,
) *CharHashsetMap {
	char := charHashsetMap.GetChar(str)

	charHashsetMap.Lock()
	hashset, has := (*charHashsetMap.items)[char]
	charHashsetMap.Unlock()

	if has {
		hashset.AddLock(str)

		return charHashsetMap
	}

	newHashset := NewHashset(charHashsetMap.eachHashsetCapacity)
	newHashset.Add(str)

	charHashsetMap.Lock()
	(*charHashsetMap.items)[char] = newHashset
	charHashsetMap.Unlock()

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) Add(
	str string,
) *CharHashsetMap {
	char := charHashsetMap.GetChar(str)

	hashset, has := (*charHashsetMap.
		items)[char]

	if has {
		hashset.Add(str)

		return charHashsetMap
	}

	newHashset := NewHashset(charHashsetMap.eachHashsetCapacity)
	newHashset.Add(str)
	(*charHashsetMap.
		items)[char] = newHashset

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) AddStringPtr(
	str *string,
) *CharHashsetMap {
	char := charHashsetMap.GetCharOfPtr(str)

	hashset, has := (*charHashsetMap.
		items)[char]

	if has {
		hashset.AddPtr(str)

		return charHashsetMap
	}

	newHashset := NewHashset(charHashsetMap.eachHashsetCapacity)
	newHashset.AddPtr(str)
	(*charHashsetMap.
		items)[char] = newHashset

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) AddStringPtrLock(
	str *string,
) *CharHashsetMap {
	defer charHashsetMap.Unlock()
	char := charHashsetMap.GetCharOfPtr(str)

	charHashsetMap.Lock()
	hashset, has := (*charHashsetMap.
		items)[char]
	charHashsetMap.Unlock()

	if has {
		hashset.AddPtrLock(str)

		return charHashsetMap
	}

	newHashset := NewHashset(charHashsetMap.eachHashsetCapacity)
	newHashset.AddPtr(str)

	charHashsetMap.Lock()
	(*charHashsetMap.
		items)[char] = newHashset
	charHashsetMap.Unlock()

	return charHashsetMap
}

// AddSameStartingCharItems Assuming all items starts with same chars
func (charHashsetMap *CharHashsetMap) AddSameStartingCharItems(
	char byte,
	allItemsWithSameChar *[]string,
) *CharHashsetMap {
	if allItemsWithSameChar == nil ||
		*allItemsWithSameChar == nil {
		return charHashsetMap
	}

	length := len(*allItemsWithSameChar)

	if length == 0 {
		return charHashsetMap
	}

	values, has := (*charHashsetMap.
		items)[char]

	if has {
		values.AddStringsPtr(allItemsWithSameChar)

		return charHashsetMap
	}

	(*charHashsetMap.items)[char] =
		NewHashsetUsingStrings(
			allItemsWithSameChar,
			length*2,
			true)

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) AddPtrStringsLock(
	simpleStrings *[]*string,
) *CharHashsetMap {
	if simpleStrings == nil ||
		*simpleStrings == nil ||
		len(*simpleStrings) == 0 {
		return charHashsetMap
	}

	for _, item := range *simpleStrings {
		foundHashset := charHashsetMap.GetHashsetLock(
			*item, true)

		foundHashset.AddPtrLock(item)
	}

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) AddStringsPtrAsyncLock(
	largeStringsHashset *[]string,
	onComplete OnCompleteCharHashsetMap,
) *CharHashsetMap {
	if largeStringsHashset == nil ||
		*largeStringsHashset == nil {
		return charHashsetMap
	}

	length := len(*largeStringsHashset)

	if length == 0 {
		return charHashsetMap
	}

	isListIsTooLargeAndHasExistingData :=
		length > RegularCollectionEfficiencyLimit &&
			charHashsetMap.Length() > DoubleLimit

	if isListIsTooLargeAndHasExistingData {
		return charHashsetMap.
			efficientAddOfLargeItems(
				largeStringsHashset,
				onComplete)
	}

	wg := &sync.WaitGroup{}
	wg.Add(length)

	for _, item := range *largeStringsHashset {
		foundHashset := charHashsetMap.GetHashsetLock(
			item,
			true)

		go foundHashset.AddWithWgLock(
			item,
			wg,
		)
	}

	wg.Wait()

	if onComplete != nil {
		onComplete(charHashsetMap)
	}

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) efficientAddOfLargeItems(
	largeStringsHashset *[]string,
	onComplete OnCompleteCharHashsetMap,
) *CharHashsetMap {
	allCharsMap := charHashsetMap.
		GetCharsPtrGroups(largeStringsHashset)

	wg := &sync.WaitGroup{}
	wg.Add(allCharsMap.Length())

	for key, hashset := range *allCharsMap.items {
		foundHashset := charHashsetMap.GetHashsetLock(
			string(key),
			true)

		go foundHashset.AddHashsetWgLock(
			hashset,
			wg,
		)
	}

	wg.Wait()

	if onComplete != nil {
		onComplete(charHashsetMap)
	}

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) AddStringsPtr(
	items *[]string,
) *CharHashsetMap {
	if items == nil ||
		*items == nil ||
		len(*items) == 0 {
		return charHashsetMap
	}

	for _, item := range *items {
		charHashsetMap.AddStringPtr(&item)
	}

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) AddStrings(
	items ...string,
) *CharHashsetMap {
	if items == nil ||
		len(items) == 0 {
		return charHashsetMap
	}

	for _, item := range items {
		charHashsetMap.AddStringPtr(&item)
	}

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) GetHashset(
	strFirstChar string,
	isAddNewOnEmpty bool,
) *Hashset {
	char := charHashsetMap.GetChar(strFirstChar)

	hashset, has := (*charHashsetMap.items)[char]

	if has {
		return hashset
	}

	if isAddNewOnEmpty {
		newHashset := NewHashset(charHashsetMap.eachHashsetCapacity)
		(*charHashsetMap.items)[char] = newHashset

		return newHashset
	}

	return nil
}

func (charHashsetMap *CharHashsetMap) GetHashsetLock(
	strFirstChar string,
	isAddNewOnEmpty bool,
) *Hashset {
	charHashsetMap.Lock()
	defer charHashsetMap.Unlock()

	return charHashsetMap.GetHashset(
		strFirstChar,
		isAddNewOnEmpty)
}

func (charHashsetMap *CharHashsetMap) AddSameCharsCollection(
	str string,
	stringsWithSameStartChar *Collection,
) *Hashset {
	isNilOrEmptyHashsetGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundHashset := charHashsetMap.GetHashset(
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

	char := charHashsetMap.GetChar(str)

	if isNilOrEmptyHashsetGiven {
		// create new
		newHashset := NewHashset(
			charHashsetMap.eachHashsetCapacity)
		(*charHashsetMap.items)[char] = newHashset

		return newHashset
	}

	// items exist or stringsWithSameStartChar exists
	//goland:noinspection GoNilness
	toHashset := stringsWithSameStartChar.HashsetAsIs()
	(*charHashsetMap.items)[char] = toHashset

	return toHashset
}

func (charHashsetMap *CharHashsetMap) AddSameCharsHashset(
	str string,
	stringsWithSameStartChar *Hashset,
) *Hashset {
	isNilOrEmptyHashsetGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundHashset := charHashsetMap.GetHashset(
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

	char := charHashsetMap.GetChar(str)

	if isNilOrEmptyHashsetGiven {
		// create new
		newHashset := NewHashset(
			charHashsetMap.eachHashsetCapacity)
		(*charHashsetMap.items)[char] = newHashset

		return newHashset
	}

	// items exist or stringsWithSameStartChar exists
	(*charHashsetMap.items)[char] =
		stringsWithSameStartChar

	return stringsWithSameStartChar
}

func (charHashsetMap *CharHashsetMap) AddHashsetItems(
	hashsetWithDiffStarts *Hashset,
) *CharHashsetMap {
	if hashsetWithDiffStarts == nil ||
		hashsetWithDiffStarts.IsEmpty() {
		return charHashsetMap
	}

	charHashsetMap.AddStringsPtr(
		hashsetWithDiffStarts.ListPtr())

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) AddHashsetItemsAsyncLock(
	hashsetWithDiffStarts *Hashset,
	onComplete OnCompleteCharHashsetMap,
) *CharHashsetMap {
	if hashsetWithDiffStarts == nil ||
		hashsetWithDiffStarts.IsEmpty() {
		return charHashsetMap
	}

	go charHashsetMap.AddStringsPtrAsyncLock(
		hashsetWithDiffStarts.ListCopyPtrLock(),
		onComplete)

	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) AddSameCharsCollectionLock(
	str string,
	stringsWithSameStartChar *Collection,
) *Hashset {
	isNilOrEmptyHashsetGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundHashset := charHashsetMap.GetHashsetLock(
		str,
		false)
	has := foundHashset != nil
	isAddToHashset := has &&
		!isNilOrEmptyHashsetGiven
	hasHashsetHoweverNothingToAdd := has &&
		isNilOrEmptyHashsetGiven

	if isAddToHashset {
		//goland:noinspection GoNilness
		list := stringsWithSameStartChar.
			ListCopyPtrLock()

		foundHashset.
			AddStringsPtrLock(list)

		return foundHashset
	} else if hasHashsetHoweverNothingToAdd {
		return foundHashset
	}

	char := charHashsetMap.GetChar(str)

	if isNilOrEmptyHashsetGiven {
		// create new
		newHashset := NewHashset(
			charHashsetMap.eachHashsetCapacity)
		charHashsetMap.Lock()
		(*charHashsetMap.items)[char] = newHashset
		charHashsetMap.Unlock()

		return newHashset
	}

	// items exist or stringsWithSameStartChar exists
	//goland:noinspection GoNilness
	hashset := stringsWithSameStartChar.HashsetAsIs()
	//goland:noinspection GoLinterLocal
	charHashsetMap.Lock()
	(*charHashsetMap.items)[char] =
		hashset
	charHashsetMap.Unlock()

	return hashset
}

func (charHashsetMap *CharHashsetMap) AddHashsetLock(
	str string,
	stringsWithSameStartChar *Hashset,
) *Hashset {
	isNilOrEmptyHashsetGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundHashset := charHashsetMap.GetHashsetLock(
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
	char := charHashsetMap.GetChar(str)

	if isNilOrEmptyHashsetGiven {
		// create new
		newHashset := NewHashset(
			charHashsetMap.eachHashsetCapacity)
		charHashsetMap.Lock()
		(*charHashsetMap.items)[char] = newHashset
		charHashsetMap.Unlock()

		return newHashset
	}

	// items exist or stringsWithSameStartChar exists
	charHashsetMap.Lock()
	(*charHashsetMap.items)[char] =
		stringsWithSameStartChar
	charHashsetMap.Unlock()

	return stringsWithSameStartChar
}

func (charHashsetMap *CharHashsetMap) GetHashsetByChar(
	char byte,
) *Hashset {
	return (*charHashsetMap.items)[char]
}

func (charHashsetMap *CharHashsetMap) HashsetByChar(
	char byte,
) *Hashset {
	hashset := (*charHashsetMap.items)[char]

	return hashset
}

func (charHashsetMap *CharHashsetMap) HashsetByCharLock(
	char byte,
) *Hashset {
	charHashsetMap.Lock()
	hashset := (*charHashsetMap.items)[char]
	charHashsetMap.Unlock()

	if hashset == nil {
		return EmptyHashset()
	}

	return hashset
}

func (charHashsetMap *CharHashsetMap) HashsetByStringFirstChar(
	str string,
) *Hashset {
	char := charHashsetMap.GetChar(str)

	return charHashsetMap.HashsetByChar(char)
}

func (charHashsetMap *CharHashsetMap) HashsetByStringFirstCharLock(
	str string,
) *Hashset {
	char := charHashsetMap.GetChar(str)

	return charHashsetMap.HashsetByCharLock(char)
}

func (charHashsetMap *CharHashsetMap) JsonModel() *CharHashsetDataModel {
	return &CharHashsetDataModel{
		Items: charHashsetMap.items,
		EachHashsetCapacity: charHashsetMap.
			eachHashsetCapacity,
	}
}

func (charHashsetMap *CharHashsetMap) JsonModelAny() interface{} {
	return charHashsetMap.JsonModel()
}

func (charHashsetMap *CharHashsetMap) AsJsoner() corejson.Jsoner {
	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) AsJsonMarshaller() corejson.JsonMarshaller {
	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return charHashsetMap
}

func (charHashsetMap *CharHashsetMap) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := charHashsetMap.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (charHashsetMap *CharHashsetMap) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*CharHashsetMap, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return EmptyCharHashsetMap(), defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(*jsonResult.Bytes, &charHashsetMap)

	if err != nil {
		return EmptyCharHashsetMap(), err
	}

	return charHashsetMap, nil
}

// ParseInjectUsingJsonMust Panic if error
func (charHashsetMap *CharHashsetMap) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *CharHashsetMap {
	newUsingJson, err :=
		charHashsetMap.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (charHashsetMap *CharHashsetMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(*charHashsetMap.JsonModel())
}

func (charHashsetMap *CharHashsetMap) UnmarshalJSON(data []byte) error {
	var dataModel CharHashsetDataModel

	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		charHashsetMap.items = dataModel.Items
		charHashsetMap.eachHashsetCapacity =
			dataModel.EachHashsetCapacity
	}

	return err
}

func (charHashsetMap *CharHashsetMap) Json() *corejson.Result {
	if charHashsetMap.IsEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(charHashsetMap.JsonModel())

	return corejson.NewPtr(jsonBytes, err)
}

// RemoveAll remove all existing items, deletes items using delete(*charCollectionMap.items, char), expensive operation
func (charHashsetMap *CharHashsetMap) RemoveAll() *CharHashsetMap {
	if charHashsetMap.IsEmpty() {
		return charHashsetMap
	}

	for char := range *charHashsetMap.items {
		delete(*charHashsetMap.items, char)
	}

	return charHashsetMap
}

// Clear points to a new map and collects old pointer and remove all elements from pointer in separate goroutine.
func (charHashsetMap *CharHashsetMap) Clear() *CharHashsetMap {
	if charHashsetMap.IsEmpty() {
		return charHashsetMap
	}

	newMap := make(map[byte]*Hashset, 0)
	tempCollection := charHashsetMap.items
	charHashsetMap.items = nil
	charHashsetMap.items = &newMap

	go func() {
		for char := range *tempCollection {
			delete(*tempCollection, char)
		}
	}()

	return charHashsetMap
}
