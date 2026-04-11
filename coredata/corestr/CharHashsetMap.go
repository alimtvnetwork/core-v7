package corestr

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"sort"
	"strings"
	"sync"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coreindexes"
)

type CharHashsetMap struct {
	items               map[byte]*Hashset
	eachHashsetCapacity int
	sync.RWMutex
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
		it.Length(),
	)

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
		it.Length(),
	)

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
		it.Length(),
	)

	for _, hashset := range it.items {
		//goland:noinspection ALL
		hashsets = append(hashsets, *hashset)
	}

	return New.HashsetsCollection.UsingHashsets(hashsets...)
}

func (it *CharHashsetMap) GetCharOf(
	str string,
) byte {
	if len(str) == 0 {
		return emptyChar
	}

	return str[coreindexes.First]
}

func (it *CharHashsetMap) GetCharsGroups(
	items ...string,
) *CharHashsetMap {
	if len(items) == 0 {
		return it
	}

	hashsetMap := New.CharHashsetMap.Cap(
		len(items),
		len(items)/3,
	)

	return hashsetMap.AddStrings(items...)
}

func (it *CharHashsetMap) GetMap() map[byte]*Hashset {
	return it.items
}

// GetCopyMapLock Sends a copy of items
func (it *CharHashsetMap) GetCopyMapLock() map[byte]*Hashset {
	it.RLock()
	defer it.RUnlock()

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
		length+1,
	)

	hashsetOfHashset[coreindexes.First] = fmt.Sprintf(
		summaryOfCharHashsetMapLengthFormat,
		it,
		length,
		coreindexes.First,
	)

	i := 1
	for key, hashset := range it.GetCopyMapLock() {
		hashsetOfHashset[i] = fmt.Sprintf(
			charHashsetMapSingleItemFormat,
			i,
			string(key),
			hashset.LengthLock(),
		)

		i++
	}

	return strings.Join(
		hashsetOfHashset,
		constants.EmptyString,
	)
}

func (it *CharHashsetMap) SummaryString() string {
	hashsetOfHashset := make(
		[]string,
		it.Length()+1,
	)

	hashsetOfHashset[coreindexes.First] = fmt.Sprintf(
		summaryOfCharHashsetMapLengthFormat,
		it,
		it.Length(),
		coreindexes.First,
	)

	i := 1
	for key, hashset := range it.items {
		hashsetOfHashset[i] = fmt.Sprintf(
			charHashsetMapSingleItemFormat,
			i,
			string(key),
			hashset.Length(),
		)

		i++
	}

	return strings.Join(
		hashsetOfHashset,
		constants.EmptyString,
	)
}

func (it *CharHashsetMap) String() string {
	hashsetOfHashset := make(
		[]string,
		it.Length()*2+1,
	)

	hashsetOfHashset[coreindexes.First] =
		it.SummaryString()

	i := 1
	for key, hashset := range it.items {
		hashsetOfHashset[i] = fmt.Sprintf(
			charHashsetMapLengthFormat,
			string(key),
		)

		i++
		hashsetOfHashset[i] = hashset.String()
		i++
	}

	return strings.Join(
		hashsetOfHashset,
		constants.EmptyString,
	)
}

func (it *CharHashsetMap) StringLock() string {
	hashsetOfHashset := make(
		[]string,
		it.LengthLock()*2+1,
	)

	hashsetOfHashset[coreindexes.First] =
		it.SummaryStringLock()

	i := 1
	for key, hashset := range it.GetCopyMapLock() {

		hashsetOfHashset[i] = fmt.Sprintf(
			charHashsetMapLengthFormat,
			string(key),
		)

		i++

		hashsetOfHashset[i] = hashset.StringLock()
		i++
	}

	return strings.Join(
		hashsetOfHashset,
		constants.EmptyString,
	)
}

func (it *CharHashsetMap) List() []string {
	list := make([]string, it.AllLengthsSum())

	i := 0
	for _, hashset := range it.items {
		for s := range hashset.items {
			list[i] = s
			i++
		}
	}

	return list
}

func (it *CharHashsetMap) SortedListAsc() []string {
	list := it.List()
	sort.Strings(list)

	return list
}

func (it *CharHashsetMap) SortedListDsc() []string {
	list := it.SortedListAsc()
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

	return list
}

func (it *CharHashsetMap) Print(isPrint bool) {
	isSkipPrint := !isPrint

	if isSkipPrint {
		return
	}

	slog.Info("char hashset map", "content", it.String())
}

func (it *CharHashsetMap) PrintLock(isPrint bool) {
	isSkipPrint := !isPrint

	if isSkipPrint {
		return
	}

	slog.Info("char hashset map (locked)", "content", it.StringLock())
}

func (it *CharHashsetMap) IsEmpty() bool {
	return it == nil ||
		len(it.items) == 0
}

func (it *CharHashsetMap) HasItems() bool {
	return it != nil && len(it.items) > 0
}

func (it *CharHashsetMap) IsEmptyLock() bool {
	it.RLock()
	defer it.RUnlock()

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
	it.RLock()
	defer it.RUnlock()

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
	it.RLock()
	defer it.RUnlock()

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
	it.RLock()
	defer it.RUnlock()

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

	it.AddStrings(charCollectionMap.List()...)

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
		collectionWithDiffStarts.items...,
	)

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

	go it.AddStringsAsyncLock(
		collectionWithDiffStarts.items,
		onComplete,
	)

	return it
}

func (it *CharHashsetMap) Length() int {
	if it.IsEmpty() {
		return 0
	}

	return len(it.items)
}

func (it *CharHashsetMap) LengthLock() int {
	it.RLock()
	defer it.RUnlock()

	if it.IsEmpty() {
		return 0
	}

	return len(it.items)
}

func (it *CharHashsetMap) IsEqualsLock(
	another *CharHashsetMap,
) bool {
	it.RLock()
	defer it.RUnlock()

	return it.IsEquals(another)
}

func (it *CharHashsetMap) IsEquals(
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
		isMissing := !has

		if isMissing {
			return false
		}

		isDifferent := !rHashset.IsEquals(hashset)

		if isDifferent {
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
	if it.items == nil {
		it.items = make(map[byte]*Hashset, defaultHashsetItems)
	}
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
	if it.items == nil {
		it.items = make(map[byte]*Hashset, defaultHashsetItems)
	}

	char := it.GetChar(str)

	if it.items == nil {
		it.items = make(map[byte]*Hashset, defaultHashsetItems)
	}
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

// AddSameStartingCharItems Assuming all items starts with same chars
func (it *CharHashsetMap) AddSameStartingCharItems(
	char byte,
	allItemsWithSameChar []string,
) *CharHashsetMap {
	if len(allItemsWithSameChar) == 0 {
		return it
	}

	if it.items == nil {
		it.items = make(map[byte]*Hashset, defaultHashsetItems)
	}

	length := len(allItemsWithSameChar)

	if length == 0 {
		return it
	}

	if it.items == nil {
		it.items = make(map[byte]*Hashset, defaultHashsetItems)
	}
	values, has := it.items[char]

	if has {
		values.AddStrings(allItemsWithSameChar)

		return it
	}

	it.items[char] =
		New.Hashset.Strings(
			allItemsWithSameChar,
		)

	return it
}

func (it *CharHashsetMap) AddStringsLock(
	simpleStrings ...string,
) *CharHashsetMap {
	if len(simpleStrings) == 0 {
		return it
	}

	for _, item := range simpleStrings {
		foundHashset := it.GetHashsetLock(
			true,
			item,
		)

		foundHashset.AddLock(item)
	}

	return it
}

func (it *CharHashsetMap) AddStringsAsyncLock(
	largeStringsHashset []string,
	onComplete OnCompleteCharHashsetMap,
) *CharHashsetMap {
	length := len(largeStringsHashset)

	if length == 0 {
		return it
	}

	isListIsTooLargeAndHasExistingData :=
		length > RegularCollectionEfficiencyLimit &&
			it.Length() > DoubleLimit

	if isListIsTooLargeAndHasExistingData {
		return it.
			efficientAddOfLargeItems(
				onComplete,
				largeStringsHashset...,
			)
	}

	wg := &sync.WaitGroup{}
	wg.Add(length)

	for _, item := range largeStringsHashset {
		foundHashset := it.GetHashsetLock(
			true,
			item,
		)

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
	onComplete OnCompleteCharHashsetMap,
	largeStringsHashset ...string,
) *CharHashsetMap {
	allCharsMap := it.GetCharsGroups(largeStringsHashset...)

	wg := &sync.WaitGroup{}
	wg.Add(allCharsMap.Length())

	for key, hashset := range allCharsMap.items {
		foundHashset := it.GetHashsetLock(
			true,
			string(key),
		)

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

func (it *CharHashsetMap) AddStrings(
	items ...string,
) *CharHashsetMap {
	if items == nil ||
		len(items) == 0 {
		return it
	}

	for _, item := range items {
		it.Add(item)
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
	isAddNewOnEmpty bool,
	strFirstChar string,
) *Hashset {
	it.RLock()
	defer it.RUnlock()

	return it.GetHashset(
		strFirstChar,
		isAddNewOnEmpty,
	)
}

func (it *CharHashsetMap) AddSameCharsCollection(
	str string,
	stringsWithSameStartChar *Collection,
) *Hashset {
	isNilOrEmptyHashsetGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundHashset := it.GetHashset(
		str,
		false,
	)

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
			it.eachHashsetCapacity,
		)
		if it.items == nil {
			it.items = make(map[byte]*Hashset, 4)
		}
		it.items[char] = newHashset

		return newHashset
	}

	// items exist or stringsWithSameStartChar exists
	//goland:noinspection GoNilness
	toHashset := stringsWithSameStartChar.HashsetAsIs()
	if it.items == nil {
		it.items = make(map[byte]*Hashset, 4)
	}
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
		false,
	)

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
			it.eachHashsetCapacity,
		)
		if it.items == nil {
			it.items = make(map[byte]*Hashset, 4)
		}
		it.items[char] = newHashset

		return newHashset
	}

	// items exist or stringsWithSameStartChar exists
	if it.items == nil {
		it.items = make(map[byte]*Hashset, 4)
	}
	it.items[char] =
		stringsWithSameStartChar

	return stringsWithSameStartChar
}

func (it *CharHashsetMap) AddHashsetItems(
	hashsetWithDiffStarts *Hashset,
) *CharHashsetMap {
	if hashsetWithDiffStarts.IsEmpty() {
		return it
	}

	it.AddStrings(
		hashsetWithDiffStarts.List()...,
	)

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

	go it.AddStringsAsyncLock(
		hashsetWithDiffStarts.ListCopyLock(),
		onComplete,
	)

	return it
}

func (it *CharHashsetMap) AddSameCharsCollectionLock(
	str string,
	stringsWithSameStartChar *Collection,
) *Hashset {
	isNilOrEmptyHashsetGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundHashset := it.GetHashsetLock(
		false,
		str,
	)
	has := foundHashset != nil
	isAddToHashset := has &&
		!isNilOrEmptyHashsetGiven
	hasHashsetHoweverNothingToAdd := has &&
		isNilOrEmptyHashsetGiven

	if isAddToHashset {
		list := stringsWithSameStartChar.
			ListCopyPtrLock()

		foundHashset.AddStringsLock(list)

		return foundHashset
	} else if hasHashsetHoweverNothingToAdd {
		return foundHashset
	}

	char := it.GetChar(str)

	if isNilOrEmptyHashsetGiven {
		// create new
		newHashset := New.Hashset.Cap(
			it.eachHashsetCapacity,
		)
		it.Lock()
		if it.items == nil {
			it.items = make(map[byte]*Hashset, 4)
		}
		it.items[char] = newHashset
		it.Unlock()

		return newHashset
	}

	// items exist or stringsWithSameStartChar exists
	//goland:noinspection GoNilness
	hashset := stringsWithSameStartChar.HashsetAsIs()
	//goland:noinspection GoLinterLocal
	it.Lock()
	if it.items == nil {
		it.items = make(map[byte]*Hashset, 4)
	}
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
		false,
		str,
	)
	has := foundHashset != nil
	isAddToHashset := has && !isNilOrEmptyHashsetGiven
	hasHashsetHoweverNothingToAdd := has && isNilOrEmptyHashsetGiven

	if isAddToHashset {
		//goland:noinspection GoNilness
		foundHashset.AddStringsLock(
			stringsWithSameStartChar.List(),
		)

		return foundHashset
	} else if hasHashsetHoweverNothingToAdd {
		return foundHashset
	}

	// current str char, no lock required
	char := it.GetChar(str)

	if isNilOrEmptyHashsetGiven {
		// create new
		newHashset := New.Hashset.Cap(
			it.eachHashsetCapacity,
		)
		it.RLock()
		if it.items == nil {
			it.items = make(map[byte]*Hashset, 4)
		}
		it.items[char] = newHashset
		it.RUnlock()

		return newHashset
	}

	// items exist or stringsWithSameStartChar exists
	it.RLock()
	if it.items == nil {
		it.items = make(map[byte]*Hashset, 4)
	}
	it.items[char] =
		stringsWithSameStartChar
	it.RUnlock()

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
	it.RLock()
	hashset := it.items[char]
	it.RUnlock()

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

func (it *CharHashsetMap) JsonModelAny() any {
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
	return corejson.New(&it)
}

func (it CharHashsetMap) JsonPtr() *corejson.Result {
	return corejson.NewPtr(&it)
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
