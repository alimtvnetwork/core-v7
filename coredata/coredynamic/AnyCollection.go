package coredynamic

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/defaultcapacity"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
	"github.com/alimtvnetwork/core/internal/strutilinternal"
	"github.com/alimtvnetwork/core/pagingutil"
)

type AnyCollection struct {
	items []any
}

func EmptyAnyCollection() *AnyCollection {
	return NewAnyCollection(constants.Zero)
}

func NewAnyCollection(capacity int) *AnyCollection {
	slice := make([]any, 0, capacity)

	return &AnyCollection{items: slice}
}

func (it *AnyCollection) At(index int) any {
	return it.items[index]
}

func (it *AnyCollection) ReflectSetAt(
	index int,
	toPointerOrBytesSet any,
) error {
	item := it.items[index]

	return ReflectSetFromTo(item, toPointerOrBytesSet)
}

func (it *AnyCollection) AtAsDynamic(index int) Dynamic {
	return NewDynamic(it.items[index], true)
}

func (it *AnyCollection) Items() []any {
	if it.IsEmpty() {
		return []any{}
	}

	return it.items
}

func (it *AnyCollection) DynamicItems() []Dynamic {
	if it.IsEmpty() {
		return []Dynamic{}
	}

	slice := make([]Dynamic, it.Length())

	for i, item := range it.items {
		slice[i] = NewDynamic(item, true)
	}

	return slice
}

func (it *AnyCollection) DynamicCollection() *DynamicCollection {
	if it.IsEmpty() {
		return EmptyDynamicCollection()
	}

	return &DynamicCollection{
		items: it.DynamicItems(),
	}
}

func (it *AnyCollection) FirstDynamic() any {
	return it.items[0]
}

func (it *AnyCollection) First() any {
	return it.items[0]
}

func (it *AnyCollection) LastDynamic() any {
	return it.items[it.LastIndex()]
}

func (it *AnyCollection) Last() any {
	return it.items[it.LastIndex()]
}

func (it *AnyCollection) FirstOrDefaultDynamic() any {
	return it.FirstOrDefault()
}

func (it *AnyCollection) FirstOrDefault() any {
	if it.IsEmpty() {
		return nil
	}

	return it.First()
}

func (it *AnyCollection) LastOrDefaultDynamic() any {
	return it.LastOrDefault()
}

func (it *AnyCollection) LastOrDefault() any {
	if it.IsEmpty() {
		return nil
	}

	return it.Last()
}

func (it *AnyCollection) SkipDynamic(skippingItemsCount int) any {
	return it.items[skippingItemsCount:]
}

func (it *AnyCollection) Skip(skippingItemsCount int) []any {
	return it.items[skippingItemsCount:]
}

func (it *AnyCollection) SkipCollection(skippingItemsCount int) *AnyCollection {
	return &AnyCollection{
		items: it.items[skippingItemsCount:],
	}
}

func (it *AnyCollection) TakeDynamic(takeDynamicItems int) any {
	return it.items[:takeDynamicItems]
}

func (it *AnyCollection) Take(takeDynamicItems int) []any {
	return it.items[:takeDynamicItems]
}

func (it *AnyCollection) TakeCollection(takeDynamicItems int) *AnyCollection {
	return &AnyCollection{
		items: it.items[:takeDynamicItems],
	}
}

func (it *AnyCollection) LimitCollection(limit int) *AnyCollection {
	return &AnyCollection{
		items: it.items[:limit],
	}
}

func (it *AnyCollection) SafeLimitCollection(limit int) *AnyCollection {
	limit = defaultcapacity.
		MaxLimit(it.Length(), limit)

	return &AnyCollection{
		items: it.items[:limit],
	}
}

func (it *AnyCollection) LimitDynamic(limit int) any {
	return it.Take(limit)
}

func (it *AnyCollection) Limit(limit int) []any {
	return it.Take(limit)
}

func (it *AnyCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.items)
}

func (it *AnyCollection) Count() int {
	return it.Length()
}

func (it *AnyCollection) IsEmpty() bool {
	if it == nil {
		return true
	}

	return len(it.items) == 0
}

func (it *AnyCollection) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *AnyCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *AnyCollection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *AnyCollection) ListStringsPtr(isIncludeFieldName bool) []string {
	slice := make([]string, constants.Zero, it.Length()+1)

	for _, anyItem := range it.items {
		str := strutilinternal.AnyToStringUsing(
			isIncludeFieldName,
			anyItem,
		)

		slice = append(slice, str)
	}

	return slice
}

func (it *AnyCollection) ListStrings(isIncludeFieldName bool) []string {
	return it.ListStringsPtr(isIncludeFieldName)
}

func (it *AnyCollection) RemoveAt(index int) (isSuccess bool) {
	isInvalidIndex := !it.HasIndex(index)

	if isInvalidIndex {
		return false
	}

	items := it.items
	it.items = append(
		items[:index],
		items[index+constants.One:]...,
	)

	return true
}

func (it *AnyCollection) Loop(
	isRunAsync bool,
	loopProcessorFunc func(index int, item any) (isBreak bool), // break will not work on async
) *AnyCollection {
	if it.IsEmpty() {
		return it
	}

	length := it.Length()

	if isRunAsync {
		wg := sync.WaitGroup{}
		wg.Add(length)
		wrappedFunc := func(index int) {
			loopProcessorFunc(index, it.items[index])

			wg.Done()
		}

		for index := 0; index < length; index++ {
			go wrappedFunc(index)
		}

		wg.Wait()

		return it
	}

	for index := 0; index < it.Length(); index++ {
		isBreak := loopProcessorFunc(index, it.items[index])

		if isBreak {
			return it
		}
	}

	return it
}

func (it *AnyCollection) LoopDynamic(
	isRunAsync bool,
	loopProcessorFunc func(index int, item Dynamic) (isBreak bool), // break will not work on async
) *AnyCollection {
	if it.IsEmpty() {
		return it
	}

	length := it.Length()

	if isRunAsync {
		wg := sync.WaitGroup{}
		wg.Add(length)
		wrappedFunc := func(index int) {
			currentItem := it.items[index]
			dynamic := NewDynamic(
				currentItem,
				reflectinternal.Is.Defined(currentItem),
			)

			loopProcessorFunc(index, dynamic)

			wg.Done()
		}

		for index := 0; index < length; index++ {
			go wrappedFunc(index)
		}

		wg.Wait()

		return it
	}

	for index := 0; index < it.Length(); index++ {
		dynamic := NewDynamic(
			it.items[index],
			reflectinternal.Is.Defined(it.items[index]),
		)
		isBreak := loopProcessorFunc(index, dynamic)

		if isBreak {
			return it
		}
	}

	return it
}

func (it *AnyCollection) AddAny(anyItem any, isValid bool) *AnyCollection {
	it.items = append(
		it.items,
		NewDynamic(anyItem, isValid),
	)

	return it
}

func (it *AnyCollection) AddAnyItemsWithTypeValidation(
	isContinueOnError,
	isNullNotAllowed bool,
	expectedType reflect.Type,
	anyItems ...any,
) error {
	if len(anyItems) == 0 {
		return nil
	}

	if isContinueOnError {
		var sliceErr []string

		for _, anyItem := range anyItems {
			err := it.AddAnyWithTypeValidation(
				isNullNotAllowed,
				expectedType,
				anyItem,
			)

			if err != nil {
				sliceErr = append(sliceErr, err.Error())
			}
		}

		return errcore.SliceToError(sliceErr)
	}

	for _, anyItem := range anyItems {
		err := it.AddAnyWithTypeValidation(
			isNullNotAllowed,
			expectedType,
			anyItem,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *AnyCollection) AddAnyWithTypeValidation(
	isNullNotAllowed bool,
	expectedType reflect.Type,
	anyItem any,
) error {
	err := ReflectTypeValidation(
		isNullNotAllowed,
		expectedType,
		anyItem,
	)

	if err != nil {
		return err
	}

	it.items = append(
		it.items,
		anyItem,
	)

	return nil
}

func (it *AnyCollection) AddNonNull(anyItem any) *AnyCollection {
	if anyItem == nil {
		return it
	}

	it.items = append(
		it.items,
		anyItem,
	)

	return it
}

func (it *AnyCollection) AddNonNullDynamic(anyItem any, isValid bool) *AnyCollection {
	if anyItem == nil {
		return it
	}

	it.items = append(
		it.items,
		NewDynamic(anyItem, isValid),
	)

	return it
}

func (it *AnyCollection) AddAnyManyDynamic(anyItems ...any) *AnyCollection {
	if anyItems == nil {
		return it
	}

	for _, item := range anyItems {
		it.items = append(
			it.items,
			NewDynamic(item, true),
		)
	}

	return it
}

func (it *AnyCollection) Add(anyItem any) *AnyCollection {
	it.items = append(it.items, anyItem)

	return it
}

func (it *AnyCollection) AddAnySliceFromSingleItem(
	sliceList any,
) *AnyCollection {
	if sliceList == nil {
		return it
	}

	items := reflectinternal.
		SliceConverter.
		ToAnyItemsAsync(sliceList)

	return it.AddMany(items...)
}

func (it *AnyCollection) AddMany(anyItems ...any) *AnyCollection {
	if anyItems == nil {
		return it
	}

	for _, anyItem := range anyItems {
		if anyItem == nil {
			continue
		}

		it.items = append(it.items, anyItem)
	}

	return it
}

func (it *AnyCollection) JsonString() (jsonString string, err error) {
	toBytes, err := json.Marshal(it.items)

	if err != nil {
		return constants.EmptyString, nil
	}

	return string(toBytes), nil
}

func (it *AnyCollection) JsonStringMust() string {
	toString, err := it.JsonString()

	if err != nil {
		errcore.
			MarshallingFailedType.
			HandleUsingPanic(err.Error(), it.items)
	}

	return toString
}

func (it *AnyCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModelAny())
}

func (it *AnyCollection) UnmarshalJSON(data []byte) error {
	var dataModelItems []any
	err := json.Unmarshal(data, &dataModelItems)

	if err == nil {
		it.items = dataModelItems
	}

	return err
}

func (it *AnyCollection) JsonResultsCollection() *corejson.ResultsCollection {
	jsonResultsCollection := corejson.NewResultsCollection.UsingCap(it.Length())

	if it.IsEmpty() {
		return jsonResultsCollection
	}

	for _, dynamicInstance := range it.items {
		jsonResultsCollection.AddAny(
			dynamicInstance,
		)
	}

	return jsonResultsCollection
}

func (it *AnyCollection) JsonResultsPtrCollection() *corejson.ResultsPtrCollection {
	jsonResultsCollection := corejson.
		NewResultsPtrCollection.
		UsingCap(it.Length())

	if it.IsEmpty() {
		return jsonResultsCollection
	}

	for _, dynamicInstance := range it.items {
		jsonResultsCollection.AddAny(
			dynamicInstance,
		)
	}

	return jsonResultsCollection
}

func (it *AnyCollection) GetPagesSize(
	eachPageSize int,
) int {
	if eachPageSize <= 0 {
		return 0
	}

	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *AnyCollection) GetPagedCollection(
	eachPageSize int,
) []*AnyCollection {
	length := it.Length()

	if length < eachPageSize {
		return []*AnyCollection{
			it,
		}
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make(
		[]*AnyCollection,
		pagesPossibleCeiling,
	)

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

func (it *AnyCollection) GetPagingInfo(
	eachPageSize int,
	pageIndex int,
) pagingutil.PagingInfo {
	return pagingutil.GetPagingInfo(
		pagingutil.PagingRequest{
			Length:       it.Length(),
			PageIndex:    pageIndex,
			EachPageSize: eachPageSize,
		},
	)
}

// GetSinglePageCollection PageIndex is one based index. Should be above or equal 1
func (it *AnyCollection) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *AnyCollection {
	length := it.Length()

	if length < eachPageSize {
		return it
	}

	pageInfo := it.GetPagingInfo(
		eachPageSize,
		pageIndex,
	)

	list := it.items[pageInfo.SkipItems:pageInfo.EndingLength]

	return &AnyCollection{
		items: list,
	}
}

func (it *AnyCollection) JsonModel() []any {
	return it.items
}

func (it *AnyCollection) JsonModelAny() any {
	return it.JsonModel()
}

func (it AnyCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it AnyCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *AnyCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*AnyCollection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

func (it *AnyCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *AnyCollection {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *AnyCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *AnyCollection) Strings() []string {
	slice := make([]string, it.Length())

	if it.IsEmpty() {
		return slice
	}

	for i, item := range it.items {
		slice[i] = fmt.Sprintf(
			constants.SprintValueFormat,
			item,
		)
	}

	return slice
}

func (it *AnyCollection) String() string {
	return strings.Join(it.Strings(), constants.NewLineUnix)
}
