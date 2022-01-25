package corepayload

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/core/errcore"
)

type PayloadsCollection struct {
	Items []*PayloadWrapper
}

func (it *PayloadsCollection) Add(
	payloadWrapper PayloadWrapper,
) *PayloadsCollection {
	it.Items = append(
		it.Items,
		&payloadWrapper)

	return it
}

func (it *PayloadsCollection) Adds(
	payloadWrappers ...PayloadWrapper,
) *PayloadsCollection {
	if len(payloadWrappers) == 0 {
		return it
	}

	for i := 0; i < len(payloadWrappers); i++ {
		it.Items = append(
			it.Items,
			&payloadWrappers[i])
	}

	return it
}

func (it *PayloadsCollection) AddsPtr(
	payloadWrappers ...*PayloadWrapper,
) *PayloadsCollection {
	if len(payloadWrappers) == 0 {
		return it
	}

	it.Items = append(
		it.Items,
		payloadWrappers...)

	return it
}

func (it *PayloadsCollection) AddsPtrOptions(
	isSkipHasIssuedPayloads bool,
	payloadWrappers ...*PayloadWrapper,
) *PayloadsCollection {
	if len(payloadWrappers) == 0 {
		return it
	}

	for i := 0; i < len(payloadWrappers); i++ {
		item := payloadWrappers[i]

		if isSkipHasIssuedPayloads && item.HasIssuesOrEmpty() {
			continue
		}

		it.Items = append(
			it.Items,
			item)
	}

	return it
}

func (it *PayloadsCollection) AddsOptions(
	isSkipHasIssuedPayloads bool,
	payloadWrappers ...PayloadWrapper,
) *PayloadsCollection {
	if len(payloadWrappers) == 0 {
		return it
	}

	for i := 0; i < len(payloadWrappers); i++ {
		item := payloadWrappers[i]

		if isSkipHasIssuedPayloads && item.HasIssuesOrEmpty() {
			continue
		}

		it.Items = append(
			it.Items,
			&item)
	}

	return it
}

func (it *PayloadsCollection) ConcatNew(
	additionalItems ...PayloadWrapper,
) *PayloadsCollection {
	cloned := it.Clone()

	return cloned.Adds(additionalItems...)
}

func (it *PayloadsCollection) ConcatNewPtr(
	additionalItemsPtr ...*PayloadWrapper,
) *PayloadsCollection {
	cloned := it.Clone()

	return cloned.AddsPtr(
		additionalItemsPtr...)
}

func (it *PayloadsCollection) AddsIf(
	isAdd bool,
	payloadWrappers ...PayloadWrapper,
) *PayloadsCollection {
	if !isAdd {
		return it
	}

	return it.Adds(payloadWrappers...)
}

func (it *PayloadsCollection) InsertAt(
	index int,
	item PayloadWrapper,
) *PayloadsCollection {
	it.Items = append(it.Items[:index+1], it.Items[index:]...)
	it.Items[index] = &item

	return it
}

func (it *PayloadsCollection) FirstDynamic() interface{} {
	return it.Items[0]
}

func (it *PayloadsCollection) First() *PayloadWrapper {
	return it.Items[0]
}

func (it *PayloadsCollection) LastDynamic() interface{} {
	return it.Items[it.LastIndex()]
}

func (it *PayloadsCollection) Last() *PayloadWrapper {
	return it.Items[it.LastIndex()]
}

func (it *PayloadsCollection) FirstOrDefaultDynamic() interface{} {
	return it.FirstOrDefault()
}

func (it *PayloadsCollection) FirstOrDefault() *PayloadWrapper {
	if it.IsEmpty() {
		return nil
	}

	return it.First()
}

func (it *PayloadsCollection) LastOrDefaultDynamic() interface{} {
	return it.LastOrDefault()
}

func (it *PayloadsCollection) LastOrDefault() *PayloadWrapper {
	if it.IsEmpty() {
		return nil
	}

	return it.Last()
}

func (it *PayloadsCollection) SkipDynamic(skippingItemsCount int) interface{} {
	return it.Items[skippingItemsCount:]
}

func (it *PayloadsCollection) Skip(skippingItemsCount int) []*PayloadWrapper {
	return it.Items[skippingItemsCount:]
}

func (it *PayloadsCollection) SkipCollection(skippingItemsCount int) *PayloadsCollection {
	return &PayloadsCollection{
		Items: it.Items[skippingItemsCount:],
	}
}

func (it *PayloadsCollection) TakeDynamic(takeDynamicItems int) interface{} {
	return it.Items[:takeDynamicItems]
}

func (it *PayloadsCollection) Take(takeDynamicItems int) []*PayloadWrapper {
	return it.Items[:takeDynamicItems]
}

func (it *PayloadsCollection) TakeCollection(takeDynamicItems int) *PayloadsCollection {
	return &PayloadsCollection{
		Items: it.Items[:takeDynamicItems],
	}
}

func (it *PayloadsCollection) LimitCollection(limit int) *PayloadsCollection {
	return &PayloadsCollection{
		Items: it.Items[:limit],
	}
}

func (it *PayloadsCollection) SafeLimitCollection(limit int) *PayloadsCollection {
	limit = defaultcapacity.
		MaxLimit(it.Length(), limit)

	return &PayloadsCollection{
		Items: it.Items[:limit],
	}
}

func (it *PayloadsCollection) LimitDynamic(limit int) interface{} {
	return it.Take(limit)
}

func (it *PayloadsCollection) Limit(limit int) []*PayloadWrapper {
	return it.Take(limit)
}

func (it *PayloadsCollection) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *PayloadsCollection) GetPagedCollection(
	eachPageSize int,
) []*PayloadsCollection {
	length := it.Length()

	if length < eachPageSize {
		return []*PayloadsCollection{
			it,
		}
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make([]*PayloadsCollection, pagesPossibleCeiling)

	wg := sync.WaitGroup{}
	addPagedItemsFunc := func(oneBasedPageIndex int) {
		pagedCollection := it.GetSinglePageCollection(
			eachPageSize,
			oneBasedPageIndex,
		)

		collectionOfCollection[oneBasedPageIndex-1] = pagedCollection

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
func (it *PayloadsCollection) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *PayloadsCollection {
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

	list := it.Items[skipItems:endingIndex]

	return New.
		PayloadsCollection.
		UsingWrappers(list...)
}

func (it *PayloadsCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *PayloadsCollection) Count() int {
	return it.Length()
}

func (it *PayloadsCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *PayloadsCollection) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *PayloadsCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *PayloadsCollection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *PayloadsCollection) Strings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.String()
	}

	return list
}

func (it *PayloadsCollection) Filter(
	filterFunc FilterFunc,
) []*PayloadWrapper {
	list := make(
		[]*PayloadWrapper, 0, it.Length())

	for _, item := range it.Items {
		isTake, isBreak := filterFunc(item)

		if isTake {
			list = append(list, item)
		}

		if isBreak {
			return list
		}
	}

	return list
}

func (it *PayloadsCollection) FilterWithLimit(
	limit int,
	filterFunc FilterFunc,
) []*PayloadWrapper {
	length := defaultcapacity.MaxLimit(
		it.Length(),
		limit)
	list := make(
		[]*PayloadWrapper,
		0,
		length)

	collectedItems := 0
	for _, item := range it.Items {
		isTake, isBreak := filterFunc(item)

		if isTake {
			list = append(list, item)
			collectedItems++
		}

		if isBreak {
			return list
		}

		if collectedItems >= length {
			return list
		}
	}

	return list
}

func (it *PayloadsCollection) FirstByFilter(
	findByFunc func(payloadWrapper *PayloadWrapper) (isFound bool),
) *PayloadWrapper {
	items := it.Filter(func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool) {
		isTake = findByFunc(payloadWrapper)

		return isTake, isTake
	})

	if len(items) > 0 {
		return items[0]
	}

	return nil
}

func (it *PayloadsCollection) FirstById(
	id string,
) *PayloadWrapper {
	return it.FirstByFilter(func(payloadWrapper *PayloadWrapper) (isFound bool) {
		return payloadWrapper.IsIdentifier(id)
	})
}

func (it *PayloadsCollection) FirstByCategory(
	category string,
) *PayloadWrapper {
	return it.FirstByFilter(func(payloadWrapper *PayloadWrapper) (isFound bool) {
		return payloadWrapper.IsCategory(category)
	})
}

func (it *PayloadsCollection) FirstByTaskType(
	taskType string,
) *PayloadWrapper {
	return it.FirstByFilter(func(payloadWrapper *PayloadWrapper) (isFound bool) {
		return payloadWrapper.IsTaskTypeName(taskType)
	})
}

func (it *PayloadsCollection) FirstByEntityType(
	entityType string,
) *PayloadWrapper {
	return it.FirstByFilter(func(payloadWrapper *PayloadWrapper) (isFound bool) {
		return payloadWrapper.IsEntityType(entityType)
	})
}

func (it *PayloadsCollection) FilterCollection(
	filterFunc FilterFunc,
) *PayloadsCollection {
	list := it.Filter(filterFunc)

	collection := New.PayloadsCollection.UsingWrappers(
		list...)

	return collection
}

func (it *PayloadsCollection) SkipFilterCollection(
	skipFilterFunc SkipFilterFunc,
) *PayloadsCollection {
	list := make(
		[]*PayloadWrapper,
		0,
		it.Length())

	for _, item := range it.Items {
		isSkip, isBreak := skipFilterFunc(item)

		if !isSkip {
			list = append(list, item)
		}

		if isBreak {
			break
		}
	}

	return New.
		PayloadsCollection.
		UsingWrappers(list...)
}

func (it *PayloadsCollection) FilterCollectionByIds(
	ids ...string,
) *PayloadsCollection {
	idsHashmap := corestr.
		New.
		Hashset.
		Strings(ids)

	return it.FilterCollection(func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool) {
		return idsHashmap.Has(payloadWrapper.Identifier), false
	})
}

func (it *PayloadsCollection) FilterNameCollection(
	name string,
) *PayloadsCollection {
	return it.FilterCollection(func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool) {
		return payloadWrapper.Name == name, false
	})
}

func (it *PayloadsCollection) FilterCategoryCollection(
	categoryName string,
) *PayloadsCollection {
	return it.FilterCollection(func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool) {
		return payloadWrapper.CategoryName == categoryName, false
	})
}

func (it *PayloadsCollection) FilterEntityTypeCollection(
	entityTypeName string,
) *PayloadsCollection {
	return it.FilterCollection(func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool) {
		return payloadWrapper.EntityType == entityTypeName, false
	})
}

func (it *PayloadsCollection) FilterTaskTypeCollection(
	taskType string,
) *PayloadsCollection {
	return it.FilterCollection(func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool) {
		return payloadWrapper.TaskTypeName == taskType, false
	})
}

func (it *PayloadsCollection) StringsUsingFmt(formatter Formatter) []string {
	list := make([]string, it.Length())

	for i := range it.Items {
		list[i] = formatter(it.Items[i])
	}

	return list
}

func (it *PayloadsCollection) JoinUsingFmt(formatter Formatter, joiner string) string {
	lines := it.StringsUsingFmt(formatter)

	return strings.Join(lines, joiner)
}

func (it *PayloadsCollection) Reverse() *PayloadsCollection {
	length := it.Length()

	if length <= 1 {
		return it
	}

	if length == 2 {
		it.Items[0], it.Items[1] = it.Items[1], it.Items[0]

		return it
	}

	mid := length / 2
	lastIndex := length - 1

	for i := 0; i < mid; i++ {
		it.Items[i], it.Items[lastIndex-i] =
			it.Items[lastIndex-i], it.Items[i]
	}

	return it
}

func (it *PayloadsCollection) JsonStrings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.JsonString()
	}

	return list
}

func (it *PayloadsCollection) JoinJsonStrings(joiner string) string {
	return strings.Join(it.JsonStrings(), joiner)
}

func (it *PayloadsCollection) Join(joiner string) string {
	return strings.Join(it.Strings(), joiner)
}

func (it *PayloadsCollection) JoinCsv() string {
	return strings.Join(it.CsvStrings(), constants.Comma)
}

func (it *PayloadsCollection) JoinCsvLine() string {
	return strings.Join(it.CsvStrings(), constants.CommaUnixNewLine)
}

func (it *PayloadsCollection) IsEqual(another *PayloadsCollection) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it.Length() != another.Length() {
		return false
	}

	return it.IsEqualItems(another.Items...)
}

func (it *PayloadsCollection) IsEqualItems(lines ...*PayloadWrapper) bool {
	if it == nil && lines == nil {
		return true
	}

	if it == nil || lines == nil {
		return false
	}

	if it.Length() != len(lines) {
		return false
	}

	for i, item := range it.Items {
		anotherItem := lines[i]

		if !item.IsEqual(anotherItem) {
			return false
		}
	}

	return true
}

func (it *PayloadsCollection) JsonString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.Json().JsonString()
}

func (it *PayloadsCollection) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.Json().JsonString()
}

func (it *PayloadsCollection) PrettyJsonString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JsonPtr().PrettyJsonString()
}

func (it *PayloadsCollection) CsvStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	newSlice := make([]string, it.Length())

	for i, item := range it.Items {
		newSlice[i] = fmt.Sprintf(
			constants.SprintDoubleQuoteFormat,
			item.String())
	}

	return newSlice
}

func (it PayloadsCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it PayloadsCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *PayloadsCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*PayloadsCollection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return Empty.PayloadsCollection(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *PayloadsCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *PayloadsCollection {
	hashSet, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}

func (it *PayloadsCollection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *PayloadsCollection) AsJsoner() corejson.Jsoner {
	return it
}

func (it *PayloadsCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *PayloadsCollection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *PayloadsCollection) Clear() *PayloadsCollection {
	if it == nil {
		return it
	}

	tempItems := it.Items
	clearFunc := func() {
		for _, item := range tempItems {
			item.Dispose()
		}
	}

	go clearFunc()

	it.Items = []*PayloadWrapper{}

	return it
}

func (it *PayloadsCollection) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.Items = nil
}

func (it PayloadsCollection) Clone() PayloadsCollection {
	list := New.PayloadsCollection.UsingCap(it.Length())

	return *list.AddsPtr(it.Items...)
}

func (it *PayloadsCollection) ClonePtr() *PayloadsCollection {
	if it == nil {
		return nil
	}

	list := New.PayloadsCollection.UsingCap(it.Length())

	return list.AddsPtr(it.Items...)
}
