package corestr

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/strutilinternal"
	"gitlab.com/evatix-go/core/simplewrap"
)

type Collection struct {
	items []string
	sync.Mutex
}

func (it *Collection) JsonString() string {
	return it.JsonPtr().JsonString()
}

func (it *Collection) JsonStringMust() string {
	return it.JsonPtr().JsonString()
}

func (it *Collection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *Collection) LastIndex() int {
	return it.Length() - 1
}

func (it *Collection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *Collection) ListStringsPtr() []string {
	return it.items
}

func (it *Collection) ListStrings() []string {
	return it.items
}

func (it *Collection) StringJSON() string {
	return it.JsonPtr().JsonString()
}

func (it *Collection) RemoveAt(index int) (isSuccess bool) {
	length := it.Length()
	if length-1 > index {
		return false
	}

	// a = append(a[:i], a[i+1:]...)
	// https://github.com/golang/go/wiki/SliceTricks
	items := it.items
	it.items = append(items[:index], items[index+1:]...)

	return true
}

func (it *Collection) Count() int {
	return it.Length()
}

func (it *Collection) Capacity() int {
	if it.items == nil {
		return 0
	}

	return cap(it.items)
}

func (it *Collection) Length() int {
	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

func (it *Collection) LengthLock() int {
	it.Lock()
	defer it.Unlock()

	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

func (it *Collection) IsEqualsPtr(
	anotherCollection *Collection,
) bool {
	return it.IsEqualsWithSensitivePtr(
		anotherCollection,
		true)
}

func (it *Collection) IsEqualsWithSensitivePtr(
	anotherCollection *Collection,
	isCaseSensitive bool,
) bool {
	if anotherCollection == nil && it == nil {
		return true
	}

	if anotherCollection == nil || it == nil {
		return false
	}

	if it == anotherCollection {
		return true
	}

	if it.IsEmpty() && anotherCollection.IsEmpty() {
		return true
	}

	if it.IsEmpty() || anotherCollection.IsEmpty() {
		return false
	}

	if it.Length() != anotherCollection.Length() {
		return false
	}

	leftItems := it.items
	rightItems := anotherCollection.items

	if isCaseSensitive {
		for i, leftVal := range leftItems {
			if leftVal != rightItems[i] {
				return false
			}
		}

		return true
	}

	for i, leftVal := range leftItems {
		if !strings.EqualFold(leftVal, rightItems[i]) {
			return false
		}
	}

	return true
}

func (it *Collection) IsEmptyLock() bool {
	it.Lock()
	defer it.Unlock()

	return it == nil ||
		len(it.items) == 0
}

func (it *Collection) IsEmpty() bool {
	return it == nil ||
		len(it.items) == 0
}

func (it *Collection) HasItems() bool {
	return it != nil &&
		len(it.items) > 0
}

func (it *Collection) AddLock(str string) *Collection {
	it.Lock()
	defer it.Unlock()

	it.items = append(
		it.items,
		str)

	return it
}

func (it *Collection) AddNonEmpty(str string) *Collection {
	if str == "" {
		return it
	}

	it.items = append(
		it.items,
		str)

	return it
}

func (it *Collection) AddNonEmptyWhitespace(str string) *Collection {
	if strutilinternal.IsEmptyOrWhitespace(str) {
		return it
	}

	it.items = append(
		it.items,
		str)

	return it
}

func (it *Collection) Add(str string) *Collection {
	it.items = append(
		it.items,
		str)

	return it
}

func (it *Collection) AddError(err error) *Collection {
	if err == nil {
		return it
	}

	it.items = append(
		it.items,
		err.Error())

	return it
}

func (it *Collection) AsDefaultError() error {
	return it.AsError(constants.NewLineUnix)
}

func (it *Collection) AsError(sep string) error {
	if it.Length() == 0 {
		return nil
	}

	toStr := it.Join(sep)

	return errors.New(toStr)
}

func (it *Collection) AddIf(
	isAdd bool,
	addingString string,
) *Collection {
	if !isAdd {
		return it
	}

	it.items = append(
		it.items,
		addingString)

	return it
}

func (it *Collection) EachItemSplitBy(splitBy string) []string {
	slice := make([]string, 0, it.Length()*constants.Capacity3)

	for _, item := range it.items {
		splitItems := strings.Split(item, splitBy)
		slice = append(slice, splitItems...)
	}

	return slice
}

func (it *Collection) ConcatNew(
	predictiveLengthAdd int,
	addingStrings ...string,
) *Collection {
	length := len(addingStrings)

	if length == 0 {
		return New.Collection.StringsOptions(
			true,
			it.items,
		)
	}

	finalLength := it.Length() + length
	capacity := defaultcapacity.PredictiveFiftyPercentIncrement(
		finalLength,
		predictiveLengthAdd)

	return New.Collection.Cap(capacity).
		Adds(it.items...).
		AddStringsPtr(&addingStrings)
}

func (it *Collection) ToError(sep string) error {
	return errcore.SliceError(sep, &it.items)
}

func (it *Collection) ToDefaultError() error {
	return errcore.SliceError(
		constants.NewLineUnix, &it.items)
}

func (it *Collection) AddIfMany(
	isAdd bool,
	addingStrings ...string,
) *Collection {
	if !isAdd {
		return it
	}

	it.items = append(
		it.items,
		addingStrings...)

	return it
}

func (it *Collection) AddFunc(f func() string) *Collection {
	it.items = append(
		it.items,
		f())

	return it
}

func (it *Collection) AddFuncErr(
	funcReturnsStringError func() (result string, err error),
	errHandler func(errInput error),
) *Collection {
	r, err := funcReturnsStringError()

	if err != nil {
		errHandler(err)

		return it
	}

	it.items = append(
		it.items,
		r)

	return it
}

func (it *Collection) AddsLock(items ...string) *Collection {
	it.Lock()
	defer it.Unlock()

	it.items = append(
		it.items,
		items...)

	return it
}

func (it *Collection) Adds(items ...string) *Collection {
	it.items = append(
		it.items,
		items...)

	return it
}

func (it *Collection) AddCollection(collectionIn *Collection) *Collection {
	return it.AddStringsPtr(&collectionIn.items)
}

// AddCollections skip on nil
func (it *Collection) AddCollections(collectionsIn ...*Collection) *Collection {
	for _, collectionIn := range collectionsIn {
		if collectionIn == nil || collectionIn.items == nil {
			continue
		}

		it.AddStringsPtr(&collectionIn.items)
	}

	return it
}

// AddPointerCollections skip on nil
func (it *Collection) AddPointerCollections(collectionsIn *[]*Collection) *Collection {
	for _, collectionIn := range *collectionsIn {
		if collectionIn == nil || collectionIn.items == nil {
			continue
		}

		it.AddStringsPtr(&collectionIn.items)
	}

	return it
}

func (it *Collection) AddPointerCollectionsLock(collectionsIn *[]*Collection) *Collection {
	it.Lock()
	defer it.Unlock()

	return it.AddPointerCollections(collectionsIn)
}

func (it *Collection) AddHashmapsValues(
	hashmaps ...*Hashmap,
) *Collection {
	if hashmaps == nil {
		return it
	}

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for _, v := range hashmap.items {
			it.items = append(
				it.items,
				v)
		}
	}

	return it
}

func (it *Collection) AddHashmapsKeys(
	hashmaps ...*Hashmap,
) *Collection {
	if hashmaps == nil {
		return it
	}

	it.resizeForHashmaps(
		&hashmaps,
		constants.One)

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k := range hashmap.items {
			it.items = append(
				it.items,
				k)
		}
	}

	return it
}

func (it *Collection) isResizeRequired(
	length int,
) bool {
	if length < constants.ArbitraryCapacity200 {
		return false
	}

	windowLength := it.Capacity() - it.Length()
	if windowLength >= length {
		return false
	}

	return true
}

func (it *Collection) resizeForHashmaps(
	hashmaps *[]*Hashmap,
	multiplier int,
) *Collection {
	if hashmaps == nil {
		return it
	}

	length := 0

	for _, hashmap := range *hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		length += hashmap.Length()
	}

	if !it.isResizeRequired(length) {
		return it
	}

	finalLength :=
		length*multiplier +
			length/2

	return it.AddCapacity(finalLength)
}

func (it *Collection) resizeForCollections(
	collections *[]*Collection,
	multiplier int,
) *Collection {
	if collections == nil {
		return it
	}

	length := 0

	for _, hashmap := range *collections {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		length += hashmap.Length()
	}

	if !it.isResizeRequired(length) {
		return it
	}

	finalLength :=
		length*multiplier +
			length/2

	return it.AddCapacity(finalLength)
}

func (it *Collection) resizeForItems(
	items *[]string,
	multiplier int,
) *Collection {
	if items == nil {
		return it
	}

	length := len(*items)
	if !it.isResizeRequired(length) {
		return it
	}

	finalLength :=
		length*multiplier +
			length/2

	return it.AddCapacity(finalLength)
}

func (it *Collection) resizeForPointerItems(
	items *[]*string,
	multiplier int,
) *Collection {
	if items == nil {
		return it
	}

	length := len(*items)
	if !it.isResizeRequired(length) {
		return it
	}

	finalLength :=
		length*multiplier +
			length/2

	return it.AddCapacity(finalLength)
}

func (it *Collection) resizeForAnys(
	items *[]interface{},
	multiplier int,
) *Collection {
	if items == nil {
		return it
	}

	length := len(*items)
	if !it.isResizeRequired(length) {
		return it
	}

	finalLength :=
		length*multiplier +
			length/2

	return it.AddCapacity(finalLength)
}

func (it *Collection) AddHashmapsKeysValues(
	hashmaps ...*Hashmap,
) *Collection {
	if hashmaps == nil {
		return it
	}

	it.resizeForHashmaps(
		&hashmaps,
		constants.ArbitraryCapacity2)

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k, v := range hashmap.items {
			it.items = append(
				it.items,
				k)
			it.items = append(
				it.items,
				v)
		}
	}

	return it
}

func (it *Collection) AddHashmapsKeysValuesUsingFilter(
	filter IsKeyValueFilter,
	hashmaps ...*Hashmap,
) *Collection {
	if hashmaps == nil {
		return it
	}

	it.resizeForHashmaps(
		&hashmaps,
		constants.One)

	for _, hashmap := range hashmaps {
		if hashmap == nil || hashmap.IsEmpty() {
			continue
		}

		for k, v := range hashmap.items {
			result, isAcceptable, isBreak := filter(KeyValuePair{
				Key:   k,
				Value: v,
			})

			if isAcceptable {
				it.items = append(
					it.items,
					result)
			}

			if isBreak {
				return it
			}
		}
	}

	return it
}

func (it *Collection) AddPtr(str *string) *Collection {
	it.items = append(
		it.items,
		*str)

	return it
}

func (it *Collection) AddPtrLock(str *string) *Collection {
	it.Lock()
	defer it.Unlock()

	it.items = append(
		it.items,
		*str)

	return it
}

func (it *Collection) AddWithWgLock(
	str string,
	group *sync.WaitGroup,
) *Collection {
	it.Lock()
	defer it.Unlock()

	it.items = append(
		it.items,
		str)

	group.Done()

	return it
}

func (it *Collection) AddsPtrLock(itemsPtr ...*string) *Collection {
	it.Lock()
	defer it.Unlock()

	for _, str := range itemsPtr {
		it.items = append(
			it.items,
			*str)
	}

	return it
}

func (it *Collection) AddStringsPtrWgLock(
	str *[]string,
	group *sync.WaitGroup,
) *Collection {
	it.Lock()
	defer it.Unlock()

	it.items = append(
		it.items,
		*str...)

	group.Done()

	return it
}

// AddPointerStringsPtrLock skip on nil
func (it *Collection) AddPointerStringsPtrLock(
	pointerStringItems *[]*string,
) *Collection {
	it.Lock()
	defer it.Unlock()

	return it.
		AddPointerStringsPtr(pointerStringItems)
}

// AddPointerStringsPtr skip on nil
func (it *Collection) AddPointerStringsPtr(
	pointerStringItems *[]*string,
) *Collection {
	for i := range *pointerStringItems {
		newPtr := (*pointerStringItems)[i]

		if newPtr == nil {
			continue
		}

		it.items = append(
			it.items,
			*(*pointerStringItems)[i])
	}

	return it
}

func (it *Collection) IndexAt(
	index int,
) string {
	return (it.items)[index]
}

func (it *Collection) SafePointerIndexAt(
	index int,
) *string {
	length := it.Length()
	if length-1 < index {
		return nil
	}

	return &(it.items)[index]
}

func (it *Collection) SafePointerIndexAtUsingLength(
	length, index int,
) *string {
	if length-1 < index {
		return nil
	}

	return &(it.items)[index]
}

func (it *Collection) SafeIndexAtUsingLength(
	defaultString string, length, index int,
) string {
	if length-1 < index {
		return defaultString
	}

	return (it.items)[index]
}

func (it *Collection) First() string {
	return (it.items)[0]
}

func (it *Collection) Single() string {
	length := it.Length()
	if length != 1 {
		errcore.LengthShouldBeEqualToType.HandleUsingPanic(
			"1",
			length)
	}

	return (it.items)[0]
}

func (it *Collection) Last() string {
	length := it.Length()

	return (it.items)[length-1]
}

func (it *Collection) LastOrDefault() string {
	length := it.Length()

	if length == 0 {
		return constants.EmptyString
	}

	return (it.items)[length-1]
}

func (it *Collection) FirstOrDefault() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return (it.items)[0]
}

// Take use One based index
func (it *Collection) Take(
	take int,
) *Collection {
	length := it.Length()

	if length <= take {
		return it
	}

	if take == 0 {
		return Empty.Collection()
	}

	list := (it.items)[:take]

	return New.Collection.StringsOptions(
		false,
		list,
	)
}

// Skip use One based index
func (it *Collection) Skip(
	skip int,
) *Collection {
	length := it.Length()

	if length < skip {
		errcore.
			LengthShouldBeEqualToType.
			HandleUsingPanic(
				"Length is lower than skip value. Skip:",
				skip)
	}

	if skip == 0 {
		return it
	}

	list := (it.items)[skip:]

	return New.Collection.StringsOptions(
		false,
		list,
	)
}

func (it *Collection) Reverse() *Collection {
	length := it.Length()

	if length <= constants.Capacity1 {
		return it
	}

	if length == constants.Capacity2 {
		it.items[0], it.items[1] = it.items[1], it.items[0]

		return it
	}

	mid := length / 2
	lastIndex := length - 1

	for i := 0; i < mid; i++ {
		it.items[i], it.items[lastIndex-i] =
			it.items[lastIndex-i], it.items[i]
	}

	return it
}

func (it *Collection) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *Collection) GetPagedCollection(
	eachPageSize int,
) *CollectionsOfCollection {
	length := it.Length()

	if length < eachPageSize {
		return New.CollectionsOfCollection.Strings(
			it.items)
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := New.CollectionsOfCollection.Cap(
		pagesPossibleCeiling)

	wg := sync.WaitGroup{}
	addPagedItemsFunc := func(oneBasedPageIndex int) {
		pagedCollection := it.GetSinglePageCollection(
			eachPageSize,
			oneBasedPageIndex,
		)

		collectionOfCollection.items[oneBasedPageIndex-1] = pagedCollection

		wg.Done()
	}

	wg.Add(pagesPossibleCeiling)
	for i := 1; i <= pagesPossibleCeiling; i++ {
		go addPagedItemsFunc(i)
	}

	wg.Wait()

	return collectionOfCollection
}

// GetSinglePageCollection PageIndex is one based index. Should be above or equal 1
func (it *Collection) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *Collection {
	length := it.Length()

	if length < eachPageSize {
		return it
	}

	/**
	 * eachPageItems = 10
	 * pageIndex = 4
	 * skipItems = 10 * (4 - 1) = 30
	 */
	skipItems := eachPageSize * (pageIndex - 1)
	if skipItems < 0 {
		errcore.
			CannotBeNegativeIndexType.
			HandleUsingPanic(
				"pageIndex cannot be negative or zero.",
				pageIndex)
	}

	endingIndex := skipItems + eachPageSize

	if endingIndex > length {
		endingIndex = length
	}

	list := (it.items)[skipItems:endingIndex]

	return New.Collection.StringsOptions(
		false,
		list,
	)
}

func (it *Collection) AddStringsPtr(
	stringItems *[]string,
) *Collection {
	if stringItems == nil {
		return it
	}

	it.resizeForItems(
		stringItems,
		constants.One)

	it.items = append(
		it.items,
		*stringItems...)

	return it
}

func (it *Collection) AddStringsPtrLock(
	stringItems *[]string,
) *Collection {
	it.Lock()
	defer it.Unlock()

	it.AddStringsPtr(stringItems)

	return it
}

func (it *Collection) AddStringsPtrAsync(
	wg *sync.WaitGroup,
	stringItems *[]string,
) *Collection {
	if stringItems == nil {
		return it
	}

	go func() {
		it.Lock()

		it.AddStringsPtr(stringItems)

		it.Unlock()

		wg.Done()
	}()

	return it
}

func (it *Collection) InsertItemsAt(
	index int, stringItems *[]string,
) *Collection {
	length := it.Length()
	isAtFirst := length == 0
	isAtLast := length-1 == index
	isAppendItems := isAtFirst || isAtLast

	if isAppendItems {
		return it.AddStringsPtr(stringItems)
	}

	// https://bit.ly/3pIDfRY
	it.items =
		append(
			(it.items)[:index],
			*stringItems...)

	it.items = append(
		it.items,
		(it.items)[index:]...)

	return it
}

func (it *Collection) ChainRemoveAt(
	index int,
) *Collection {
	it.items = append(
		(it.items)[:index],
		(it.items)[index+1:]...)

	return it
}

// RemoveItemsIndexes creates a new collection without the indexes mentioned.
//
// it is better to filter out than remove.
func (it *Collection) RemoveItemsIndexes(
	isIgnoreRemoveError bool,
	indexes ...int,
) *Collection {
	if isIgnoreRemoveError && indexes == nil {
		return it
	}

	return it.
		RemoveItemsIndexesPtr(isIgnoreRemoveError, indexes)
}

// RemoveItemsIndexesPtr creates a new collection without the indexes mentioned.
//
// it is better to filter out than remove.
func (it *Collection) RemoveItemsIndexesPtr(
	isIgnoreRemoveError bool,
	indexes []int,
) *Collection {
	if indexes == nil {
		return it
	}

	length := it.Length()
	indexesLength := len(indexes)
	hasPossibleError := length == 0 && indexesLength > 0

	if hasPossibleError && !isIgnoreRemoveError {
		panic(errcore.CannotRemoveIndexesFromEmptyCollectionType)
	}

	if !isIgnoreRemoveError {
		errcore.PanicOnIndexOutOfRange(length, indexes)
	}

	if hasPossibleError {
		return it
	}

	newList := make([]string, 0, it.Capacity())
	for i, s := range it.items { //nolint:wsl
		if coreindexes.HasIndex(indexes, i) {
			continue
		}

		newList = append(newList, s)
	}

	it.items = newList

	return it
}

func (it *Collection) AppendCollectionPtr(
	anotherCollection *Collection,
) *Collection {
	it.resizeForItems(
		&anotherCollection.items,
		constants.One)

	it.items = append(
		it.items,
		anotherCollection.items...)

	return it
}

func (it *Collection) AppendCollectionsPtr(
	anotherCollectionsPtr ...*Collection,
) *Collection {
	if anotherCollectionsPtr == nil {
		return it
	}

	return it.AppendPointersCollectionsPtr(
		&anotherCollectionsPtr)
}

func (it *Collection) AppendPointersCollectionsPtr(
	anotherCollectionsPtr *[]*Collection,
) *Collection {
	if anotherCollectionsPtr == nil {
		return it
	}

	it.resizeForCollections(
		anotherCollectionsPtr,
		constants.One)

	capacitiesIncrease := 0
	for _, currentCollection := range *anotherCollectionsPtr {

		if currentCollection == nil || currentCollection.IsEmpty() {
			continue
		}

		capacitiesIncrease += currentCollection.Length()
	}

	it.AddCapacity(capacitiesIncrease)

	for _, currentCollection := range *anotherCollectionsPtr {
		if currentCollection == nil || currentCollection.IsEmpty() {
			continue
		}

		it.items = append(
			it.items,
			currentCollection.items...)
	}

	return it
}

func (it *Collection) AppendCollectionsPtrAsync(
	wg *sync.WaitGroup,
	anotherCollectionsPtr ...*Collection,
) *Collection {
	if anotherCollectionsPtr == nil {
		return it
	}

	go func() {
		it.AppendPointersCollectionsPtr(
			&anotherCollectionsPtr)

		wg.Done()
	}()

	return it
}

// AppendAnysAsync Continue on nil
func (it *Collection) AppendAnysAsync(
	wg *sync.WaitGroup,
	anys ...interface{},
) *Collection {
	if anys == nil {
		return it
	}

	go func() {
		it.Lock()
		it.resizeForAnys(
			&anys,
			constants.One)
		it.Unlock()

		it.AppendAnysLock(&anys)

		wg.Done()
	}()

	return it
}

// AppendAnysLock Continue on nil
func (it *Collection) AppendAnysLock(
	anys *[]interface{},
) *Collection {
	if anys == nil {
		return it
	}

	it.resizeForAnys(
		anys,
		constants.One)

	for _, any := range *anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)

		it.Lock()
		it.items = append(
			it.items,
			anyStr)
		it.Unlock()
	}

	return it
}

// AppendAnys Continue on nil
func (it *Collection) AppendAnys(
	anys ...interface{},
) *Collection {
	if anys == nil {
		return it
	}

	it.resizeForAnys(
		&anys,
		constants.One)

	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(
			constants.SprintValueFormat,
			any,
		)

		it.items = append(
			it.items,
			anyStr)
	}

	return it
}

// AppendAnysUsingFilter Skip on nil
func (it *Collection) AppendAnysUsingFilter(
	filter IsStringFilter,
	anys ...interface{},
) *Collection {
	if anys == nil {
		return it
	}

	it.resizeForAnys(
		&anys,
		constants.One)

	for i, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(
			constants.SprintValueFormat,
			any)

		result, isKeep, isBreak := filter(anyStr, i)

		if !isKeep {
			continue
		}

		it.items = append(
			it.items,
			result)

		if isBreak {
			return it
		}
	}

	return it
}

// AppendAnysUsingFilterLock Skip on nil
func (it *Collection) AppendAnysUsingFilterLock(
	filter IsStringFilter,
	anys ...interface{},
) *Collection {
	if anys == nil {
		return it
	}

	it.resizeForAnys(
		&anys,
		constants.One)

	for i, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)
		result, isKeep, isBreak := filter(anyStr, i)

		if !isKeep {
			continue
		}

		it.Lock()
		it.items = append(
			it.items,
			result)
		it.Unlock()

		if isBreak {
			return it
		}
	}

	return it
}

// AppendNonEmptyAnys Continue on nil
func (it *Collection) AppendNonEmptyAnys(
	anys ...interface{},
) *Collection {
	if anys == nil {
		return it
	}

	it.resizeForAnys(
		&anys,
		constants.One)

	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)
		if anyStr == "" {
			continue
		}

		it.items = append(
			it.items,
			anyStr)
	}

	return it
}

// AddsPtr Skip on nil
func (it *Collection) AddsPtr(itemsPtr ...*string) *Collection {
	if itemsPtr == nil {
		return it
	}

	for _, str := range itemsPtr {
		if str == nil {
			continue
		}

		it.items = append(
			it.items,
			*str)
	}

	return it
}

// AddsPtrAsync Skip on nil
func (it *Collection) AddsPtrAsync(
	wg *sync.WaitGroup,
	itemsPtr ...*string,
) *Collection {
	if itemsPtr == nil {
		return it
	}

	go func() {
		it.Lock()
		it.resizeForPointerItems(
			&itemsPtr,
			constants.One)

		it.Unlock()

		for _, str := range itemsPtr {
			if str == nil {
				continue
			}

			it.Lock()

			it.items = append(
				it.items,
				*str)

			it.Unlock()
		}

		wg.Done()
	}()

	return it
}

func (it *Collection) AddsNonEmptyPtr(itemsPtr ...*string) *Collection {
	if itemsPtr == nil {
		return it
	}

	for _, str := range itemsPtr {
		if str == nil || *str == "" {
			continue
		}

		it.items = append(
			it.items,
			*str)
	}

	return it
}

func (it *Collection) AddsNonEmptyPtrLock(
	itemsPtr ...*string,
) *Collection {
	if itemsPtr == nil {
		return it
	}

	for _, str := range itemsPtr {
		if str == nil || *str == "" {
			continue
		}

		it.Lock()
		it.items = append(
			it.items,
			*str)
		it.Unlock()
	}

	return it
}

func (it *Collection) UniqueBoolMapLock() *map[string]bool {
	it.Lock()
	defer it.Unlock()

	return it.UniqueBoolMap()
}

func (it *Collection) UniqueBoolMap() *map[string]bool {
	respectiveMap := make(
		map[string]bool,
		it.Length())

	for _, item := range it.items {
		respectiveMap[item] = true
	}

	return &respectiveMap
}

func (it *Collection) UniqueListPtr() *[]string {
	boolMap := it.UniqueBoolMap()
	list := make([]string, len(*boolMap))

	i := 0
	for str := range *boolMap {
		list[i] = str
		i++
	}

	return &list
}

func (it *Collection) UniqueListPtrLock() *[]string {
	it.Lock()
	defer it.Unlock()

	return it.UniqueListPtr()
}

func (it *Collection) UniqueListLock() []string {
	it.Lock()
	defer it.Unlock()

	return it.UniqueList()
}

func (it *Collection) UniqueList() []string {
	return *it.UniqueListPtr()
}

func (it *Collection) List() []string {
	return it.items
}

// Filter must return a slice
func (it *Collection) Filter(filter IsStringFilter) []string {
	if it.IsEmpty() {
		return []string{}
	}

	list := make([]string, 0, it.Length())

	for i, element := range it.items {
		result, isKeep, isBreak := filter(element, i)

		if isKeep {
			list = append(list, result)
		}

		if isBreak {
			return list
		}
	}

	return list
}

// FilterLock must return a slice
func (it *Collection) FilterLock(filter IsStringFilter) []string {
	elements := it.ListCopyPtrLock()
	length := len(elements)

	if length == 0 {
		return elements
	}

	list := make([]string, 0, length)

	for i, element := range elements {
		result, isKeep, isBreak := filter(element, i)

		if isKeep {
			list = append(list, result)
		}

		if isBreak {
			return list
		}
	}

	return list
}

// FilteredCollection must return a items
func (it *Collection) FilteredCollection(filter IsStringFilter) *Collection {
	return New.Collection.Strings(it.Filter(filter))
}

// FilteredCollectionLock must return a items
func (it *Collection) FilteredCollectionLock(filter IsStringFilter) *Collection {
	return New.Collection.Strings(it.FilterLock(filter))
}

// FilterPtrLock must return a slice
func (it *Collection) FilterPtrLock(
	filterPtr IsStringPointerFilter,
) *[]*string {
	elements := it.ListCopyPtrLock()
	length := len(elements)

	if length == 0 {
		return &([]*string{})
	}

	list := make([]*string, 0, length)

	for i := range elements {
		copyTo := elements[i]
		result, isKeep, isBreak :=
			filterPtr(&copyTo, i)

		if isKeep {
			list = append(list, result)
		}

		if isBreak {
			return &list
		}
	}

	return &list
}

// FilterPtr must return a slice
func (it *Collection) FilterPtr(filterPtr IsStringPointerFilter) *[]*string {
	if it.IsEmpty() {
		return &([]*string{})
	}

	list := make([]*string, 0, it.Length())

	for i := range it.items {
		result, isKeep, isBreak := filterPtr(
			&it.items[i], i)

		if isKeep {
			list = append(list, result)
		}

		if isBreak {
			return &list
		}
	}

	return &list
}

// NonEmptyListPtr must return a slice
func (it *Collection) NonEmptyListPtr() *[]string {
	if it.IsEmpty() {
		return &([]string{})
	}

	list := make([]string, 0, it.Length())

	for _, element := range it.items {
		if element == "" {
			continue
		}

		list = append(list, element)
	}

	return &list
}

func (it *Collection) HashsetAsIs() *Hashset {
	return New.Hashset.Strings(
		it.items)
}

func (it *Collection) HashsetWithDoubleLength() *Hashset {
	return New.Hashset.StringsOption(
		it.Length()*2,
		false,
		it.items...)
}

func (it *Collection) HashsetLock() *Hashset {
	return New.Hashset.Strings(
		it.ListCopyPtrLock())
}

func (it *Collection) NonEmptyItems() []string {
	return stringslice.NonEmptySlice(it.items)
}

func (it *Collection) NonEmptyItemsPtr() *[]string {
	return stringslice.NonEmptySlicePtr(&it.items)
}

func (it *Collection) NonEmptyItemsOrNonWhitespacePtr() *[]string {
	return stringslice.NonWhitespaceSlicePtr(&it.items)
}

// Items direct return pointer
func (it *Collection) Items() []string {
	return it.items
}

// ListPtr direct return pointer
func (it *Collection) ListPtr() *[]string {
	return &it.items
}

// ListCopyPtrLock returns a copy of the items
//
// must return a slice
func (it *Collection) ListCopyPtrLock() []string {
	it.Lock()
	defer it.Unlock()

	if it.IsEmpty() {
		return []string{}
	}

	return it.items
}

func (it *Collection) HasLock(str string) bool {
	it.Lock()
	defer it.Unlock()

	return it.Has(str)
}

func (it *Collection) Has(str string) bool {
	if it.IsEmpty() {
		return false
	}

	for _, element := range it.items {
		if element == str {
			return true
		}
	}

	return false
}

func (it *Collection) HasPtr(str *string) bool {
	if str == nil || it.IsEmpty() {
		return false
	}

	for _, element := range it.items {
		if element == *str {
			return true
		}
	}

	return false
}

func (it *Collection) HasAll(items ...string) bool {
	if it.IsEmpty() {
		return false
	}

	for _, element := range items {
		if !it.IsContainsPtr(&element) {
			return false
		}
	}

	return true
}

// SortedListAsc Creates new doesn't modify current collection
func (it *Collection) SortedListAsc() *[]string {
	if it.IsEmpty() {
		return &[]string{}
	}

	list := &(it.items)
	sort.Strings(*list)

	return list
}

// SortedAsc mutates current collection
func (it *Collection) SortedAsc() *Collection {
	if it.IsEmpty() {
		return it
	}

	sort.Strings(it.items)

	return it
}

// SortedAscLock mutates current collection
func (it *Collection) SortedAscLock() *Collection {
	if it.IsEmptyLock() {
		return it
	}

	it.Lock()
	defer it.Unlock()

	sort.Strings(it.items)

	return it
}

// SortedListDsc Creates new one.
func (it *Collection) SortedListDsc() *[]string {
	list := it.SortedListAsc()
	stringslice.InPlaceReverse(
		&it.items)

	return list
}

func (it *Collection) HasUsingSensitivity(str string, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return it.Has(str)
	}

	for _, element := range it.items {
		if strings.EqualFold(element, str) {
			return true
		}
	}

	return false
}

func (it *Collection) IsContainsPtr(item *string) bool {
	if item == nil || it.IsEmpty() {
		return false
	}

	for _, element := range it.items {
		if element == *item {
			return true
		}
	}

	return false
}

// GetHashsetPlusHasAll nil will return false.
func (it *Collection) GetHashsetPlusHasAll(items []string) (*Hashset, bool) {
	hashset := it.HashsetAsIs()

	if items == nil || it.IsEmpty() {
		return hashset, false
	}

	return hashset, hashset.HasAllStrings(items)
}

// IsContainsAllPtr nil will return false.
func (it *Collection) IsContainsAllPtr(items *[]string) bool {
	if items == nil {
		return false
	}

	if it.IsEmpty() {
		return false
	}

	for _, item := range *items {
		if !it.IsContainsPtr(&item) {
			return false
		}
	}

	return true
}

// IsContainsAll nil will return false.
func (it *Collection) IsContainsAll(items ...string) bool {
	if items == nil {
		return false
	}

	return it.IsContainsAllPtr(&items)
}

// IsContainsAllLock nil will return false.
func (it *Collection) IsContainsAllLock(items ...string) bool {
	it.Lock()
	defer it.Unlock()

	if items == nil {
		return false
	}

	return it.IsContainsAllPtr(&items)
}

func (it *Collection) New(
	slice ...string,
) *Collection {
	length := len(slice)

	if length == 0 {
		return New.Collection.Cap(constants.Zero)
	}

	newCollection := New.Collection.Cap(constants.Zero)

	return newCollection.AddStringsPtr(&slice)
}

func (it *Collection) AddNonEmptyStrings(
	slice ...string,
) *Collection {
	if len(slice) == 0 {
		return it
	}

	return it.
		AddNonEmptyStringsPtr(&slice)
}

func (it *Collection) AddFuncResult(
	getterFunctions ...func() string,
) *Collection {
	if getterFunctions == nil {
		return it
	}

	items := it.items

	for _, getterFunc := range getterFunctions {
		item := getterFunc()

		items = append(items, item)
	}

	it.items = items

	return it
}

func (it *Collection) AddNonEmptyStringsPtr(
	slice *[]string,
) *Collection {
	if slice == nil || len(*slice) == 0 {
		return it
	}

	items := it.items

	for _, addingItem := range *slice {
		items = append(items, addingItem)
	}

	it.items = items

	return it
}

func (it *Collection) AddStringsByFuncChecking(
	slice *[]string,
	isIntegrityOkay func(line string) bool,
) *Collection {

	for _, item := range *slice {
		if isIntegrityOkay(item) {
			it.Add(item)
		}
	}

	return it
}

func (it *Collection) ExpandSlicePlusAdd(
	slice *[]string,
	expandFunc func(line string) *[]string,
) *Collection {
	items := stringslice.ExpandByFunc(slice, expandFunc)

	return it.AddStringsPtr(items)
}

func (it *Collection) MergeSlicesOfSlicePtr(slices *[]*[]string) *Collection {
	it.items = *stringslice.MergeNewSlicesPtrOfSlicesPtr(slices)

	return it
}

func (it *Collection) MergeSlicesOfSlice(slices ...*[]string) *Collection {
	it.items = *stringslice.MergeNewSlicesPtrOfSlicesPtr(&slices)

	return it
}

// GetAllExceptCollection Get all items except the mentioned ones.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this collection
// Set B = itemsCollection given in parameters.
func (it *Collection) GetAllExceptCollection(itemsCollection *Collection) *[]string {
	if itemsCollection == nil || itemsCollection.IsEmpty() {
		newItems := it.items

		return &newItems
	}

	finalList := make(
		[]string,
		0,
		it.Length())

	for _, item := range it.items {
		if itemsCollection.Has(item) {
			continue
		}

		finalList = append(
			finalList,
			item)
	}

	return &finalList
}

// GetAllExcept Get all items except the mentioned ones.
// Always returns a copy of new strings.
// It is like set A - B
// Set A = this collection
// Set B = items given in parameters.
func (it *Collection) GetAllExcept(items []string) *[]string {
	if items == nil {
		newItems := it.items

		return &newItems
	}

	newCollection := New.Collection.StringsOptions(
		false,
		items,
	)

	return it.GetAllExceptCollection(
		newCollection)
}

func (it *Collection) CharCollectionMap() *CharCollectionMap {
	length := it.Length()
	lengthByFourBestGuess := length / 4
	runeMap := New.CharCollectionMap.CapSelfCap(
		length,
		lengthByFourBestGuess)

	return runeMap.AddStringsPtr(&it.items)
}

func (it *Collection) SummaryString(sequence int) string {
	header := fmt.Sprintf(
		summaryOfCharCollectionMapLengthFormat,
		it,
		it.Length(),
		sequence)

	return it.SummaryStringWithHeader(header)
}

func (it *Collection) SummaryStringWithHeader(header string) string {
	if it.IsEmpty() {
		return header + commonJoiner + NoElements
	}

	return header + it.String()
}

func (it *Collection) String() string {
	if it.IsEmpty() {
		return commonJoiner + NoElements
	}

	return commonJoiner +
		strings.Join(
			it.items,
			commonJoiner)
}

func (it *Collection) CsvLines() []string {
	return simplewrap.DoubleQuoteWrapElements(
		false,
		it.items...,
	)
}

func (it *Collection) CsvLinesOptions(
	isSkipQuoteOnlyOnExistence bool,
) []string {
	return simplewrap.DoubleQuoteWrapElements(
		isSkipQuoteOnlyOnExistence,
		it.items...)
}

func (it *Collection) Csv() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.CsvOptions(false)
}

func (it *Collection) CsvOptions(isSkipQuoteOnlyOnExistence bool) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return strings.Join(
		it.CsvLinesOptions(isSkipQuoteOnlyOnExistence),
		constants.Comma)
}

func (it *Collection) StringLock() string {
	if it.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	it.Lock()
	defer it.Unlock()

	return commonJoiner +
		strings.Join(
			it.items,
			commonJoiner)
}

func (it *Collection) AddCapacity(
	capacities ...int,
) *Collection {
	if capacities == nil || len(capacities) == 0 {
		return it
	}

	currentCapacity := it.Capacity()

	for _, capacity := range capacities {
		currentCapacity += capacity
	}

	return it.Resize(currentCapacity)
}

// Resize Only resize if capacity is bigger than the current one
func (it *Collection) Resize(
	newCapacity int,
) *Collection {
	capacity := it.Capacity()
	if capacity >= newCapacity {
		return it
	}

	newItems := make([]string, it.Length(), newCapacity)
	copy(newItems, it.items)

	it.items = newItems

	return it
}

func (it *Collection) Joins(
	separator string,
	items ...string,
) string {
	if len(items) == 0 {
		return strings.Join(it.items, separator)
	}

	newItems := make([]string, 0, it.Length()+len(items))
	copy(newItems, it.items)

	newItems = append(newItems, items...)

	return strings.Join(newItems, separator)
}

func (it *Collection) NonEmptyJoins(
	joiner string,
) string {
	return stringslice.NonEmptyJoinPtr(
		&it.items,
		joiner)
}

func (it *Collection) NonWhitespaceJoins(
	joiner string,
) string {
	return stringslice.NonWhitespaceJoinPtr(
		&it.items,
		joiner)
}

func (it *Collection) JsonModel() []string {
	return it.items
}

func (it *Collection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *Collection) UnmarshalJSON(data []byte) error {
	var dataModel []string

	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.items = dataModel
	}

	return err
}

func (it Collection) Json() corejson.Result {
	return corejson.New(it)
}

func (it Collection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *Collection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Collection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return Empty.Collection(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//goland:noinspection GoLinterLocal
func (it *Collection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Collection {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *Collection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *Collection) Clear() *Collection {
	if it == nil {
		return nil
	}

	it.items = it.items[:0]

	return it
}

func (it *Collection) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.items = nil
}

func (it *Collection) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it *Collection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Collection) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *Collection) Deserialize(toPtr interface{}) (parsingErr error) {
	return it.JsonPtr().Deserialize(toPtr)
}
