package coredynamic

import (
	"encoding/json"
	"math"
	"reflect"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/pagingutil"
)

type DynamicCollection struct {
	items []Dynamic
}

func EmptyDynamicCollection() *DynamicCollection {
	return NewDynamicCollection(constants.Zero)
}

func NewDynamicCollection(capacity int) *DynamicCollection {
	slice := make([]Dynamic, 0, capacity)

	return &DynamicCollection{items: slice}
}

func (it *DynamicCollection) At(index int) Dynamic {
	return it.items[index]
}

func (it *DynamicCollection) Items() []Dynamic {
	if it == nil || it.items == nil {
		return []Dynamic{}
	}

	return it.items
}

func (it *DynamicCollection) FirstDynamic() interface{} {
	return it.items[0]
}

func (it *DynamicCollection) First() Dynamic {
	return it.items[0]
}

func (it *DynamicCollection) LastDynamic() interface{} {
	return it.items[it.LastIndex()]
}

func (it *DynamicCollection) Last() Dynamic {
	return it.items[it.LastIndex()]
}

func (it *DynamicCollection) FirstOrDefaultDynamic() interface{} {
	return it.FirstOrDefault()
}

func (it *DynamicCollection) FirstOrDefault() *Dynamic {
	if it.IsEmpty() {
		return nil
	}

	first := it.First()

	return &first
}

func (it *DynamicCollection) LastOrDefaultDynamic() interface{} {
	return it.LastOrDefault()
}

func (it *DynamicCollection) LastOrDefault() *Dynamic {
	if it.IsEmpty() {
		return nil
	}

	last := it.Last()

	return &last
}

func (it *DynamicCollection) SkipDynamic(skippingItemsCount int) interface{} {
	return it.items[skippingItemsCount:]
}

func (it *DynamicCollection) Skip(skippingItemsCount int) []Dynamic {
	return it.items[skippingItemsCount:]
}

func (it *DynamicCollection) SkipCollection(skippingItemsCount int) *DynamicCollection {
	return &DynamicCollection{
		items: it.items[skippingItemsCount:],
	}
}

func (it *DynamicCollection) TakeDynamic(takeDynamicItems int) interface{} {
	return it.items[:takeDynamicItems]
}

func (it *DynamicCollection) Take(takeDynamicItems int) []Dynamic {
	return it.items[:takeDynamicItems]
}

func (it *DynamicCollection) TakeCollection(takeDynamicItems int) *DynamicCollection {
	return &DynamicCollection{
		items: it.items[:takeDynamicItems],
	}
}

func (it *DynamicCollection) LimitCollection(limit int) *DynamicCollection {
	return &DynamicCollection{
		items: it.items[:limit],
	}
}

func (it *DynamicCollection) SafeLimitCollection(limit int) *DynamicCollection {
	limit = defaultcapacity.
		MaxLimit(it.Length(), limit)

	return &DynamicCollection{
		items: it.items[:limit],
	}
}

func (it *DynamicCollection) LimitDynamic(limit int) interface{} {
	return it.Take(limit)
}

func (it *DynamicCollection) Limit(limit int) []Dynamic {
	return it.Take(limit)
}

func (it *DynamicCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.items)
}

func (it *DynamicCollection) Count() int {
	return it.Length()
}

func (it *DynamicCollection) IsEmpty() bool {
	if it == nil {
		return true
	}

	return len(it.items) == 0
}

func (it *DynamicCollection) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *DynamicCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *DynamicCollection) Loop(
	loopProcessorFunc func(index int, dynamicItem *Dynamic) (isBreak bool),
) {
	if it.IsEmpty() {
		return
	}

	for i := range it.items {
		isBreak := loopProcessorFunc(i, &it.items[i])

		if isBreak {
			return
		}
	}
}

func (it *DynamicCollection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *DynamicCollection) ListStringsPtr() *[]string {
	slice := make([]string, constants.Zero, it.Length()+1)

	for _, dynamic := range it.items {
		str, _ := dynamic.JsonString()

		slice = append(slice, str)
	}

	return &slice
}

func (it *DynamicCollection) ListStrings() []string {
	return *it.ListStringsPtr()
}

func (it *DynamicCollection) RemoveAt(index int) (isSuccess bool) {
	if !it.HasIndex(index) {
		return false
	}

	items := it.items
	it.items = append(
		items[:index],
		items[index+constants.One:]...)

	return true
}

func (it *DynamicCollection) AddAnyItemsWithTypeValidation(
	isContinueOnError,
	isNullNotAllowed bool,
	expectedType reflect.Type,
	anyItems ...interface{},
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
				anyItem)

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
			anyItem)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *DynamicCollection) AddAnyWithTypeValidation(
	isNullNotAllowed bool,
	expectedType reflect.Type,
	anyItem interface{},
) error {
	err := ReflectTypeValidation(
		isNullNotAllowed,
		expectedType,
		anyItem)

	if err != nil {
		return err
	}

	it.items = append(
		it.items,
		NewDynamic(anyItem, true))

	return nil
}

func (it *DynamicCollection) AddAny(
	anyItem interface{}, isValid bool,
) *DynamicCollection {
	it.items = append(
		it.items,
		NewDynamic(anyItem, isValid))

	return it
}

func (it *DynamicCollection) AddAnyNonNull(
	anyItem interface{}, isValid bool,
) *DynamicCollection {
	if anyItem == nil {
		return it
	}

	it.items = append(
		it.items,
		NewDynamic(anyItem, isValid))

	return it
}

func (it *DynamicCollection) AddAnyMany(
	anyItems ...interface{},
) *DynamicCollection {
	if anyItems == nil {
		return it
	}

	for _, item := range anyItems {
		it.items = append(
			it.items,
			NewDynamic(item, true))
	}

	return it
}

func (it *DynamicCollection) Add(
	dynamic Dynamic,
) *DynamicCollection {
	it.items = append(it.items, dynamic)

	return it
}

func (it *DynamicCollection) AddPtr(
	dynamic *Dynamic,
) *DynamicCollection {
	if dynamic == nil {
		return it
	}

	it.items = append(it.items, *dynamic)

	return it
}

func (it *DynamicCollection) AddManyPtr(
	dynamicItems ...*Dynamic,
) *DynamicCollection {
	if dynamicItems == nil {
		return it
	}

	for _, item := range dynamicItems {
		if item == nil {
			continue
		}

		it.items = append(it.items, *item)
	}

	return it
}

func (it *DynamicCollection) AnyItems() []interface{} {
	if it.IsEmpty() {
		return []interface{}{}
	}

	slice := make([]interface{}, it.Length())

	for i, dynamicInstance := range it.items {
		slice[i] = dynamicInstance.Value()
	}

	return slice
}

func (it *DynamicCollection) AddAnySliceFromSingleItem(
	isValid bool,
	sliceList interface{},
) *DynamicCollection {
	if sliceList == nil {
		return it
	}

	items := AnySliceValToInterfacesAsync(sliceList)
	for _, item := range items {
		it.items = append(
			it.items,
			NewDynamic(item, isValid))
	}

	return it
}

func (it *DynamicCollection) AnyItemsCollection() *AnyCollection {
	if it.IsEmpty() {
		return EmptyAnyCollection()
	}

	slice := it.AnyItems()

	return &AnyCollection{items: slice}
}

func (it *DynamicCollection) JsonString() (jsonString string, err error) {
	toBytes, err := json.Marshal(it.items)

	if err != nil {
		return constants.EmptyString, nil
	}

	return string(toBytes), nil
}

func (it *DynamicCollection) JsonStringMust() string {
	toString, err := it.JsonString()

	if err != nil {
		errcore.
			MarshallingFailedType.
			HandleUsingPanic(err.Error(), it.items)
	}

	return toString
}

func (it *DynamicCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModelAny())
}

func (it *DynamicCollection) UnmarshalJSON(data []byte) error {
	var dataModel DynamicCollectionModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.items = dataModel.Items
	}

	return err
}

func (it *DynamicCollection) JsonResultsCollection() *corejson.ResultsCollection {
	jsonResultsCollection := corejson.NewResultsCollection.UsingCap(it.Length())

	if it.IsEmpty() {
		return jsonResultsCollection
	}

	for _, dynamicInstance := range it.items {
		jsonResultsCollection.AddAny(
			dynamicInstance.Value())
	}

	return jsonResultsCollection
}

func (it *DynamicCollection) JsonResultsPtrCollection() *corejson.ResultsPtrCollection {
	jsonResultsCollection := corejson.
		NewResultsPtrCollection.
		UsingCap(it.Length())

	if it.IsEmpty() {
		return jsonResultsCollection
	}

	for _, dynamicInstance := range it.items {
		jsonResultsCollection.AddAny(
			dynamicInstance.Value())
	}

	return jsonResultsCollection
}

func (it *DynamicCollection) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *DynamicCollection) GetPagedCollection(
	eachPageSize int,
) []*DynamicCollection {
	length := it.Length()

	if length < eachPageSize {
		return []*DynamicCollection{
			it,
		}
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make(
		[]*DynamicCollection,
		pagesPossibleCeiling)

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

func (it *DynamicCollection) GetPagingInfo(
	eachPageSize int,
	pageIndex int,
) pagingutil.PagingInfo {
	return pagingutil.GetPagingInfo(pagingutil.PagingRequest{
		Length:       it.Length(),
		PageIndex:    pageIndex,
		EachPageSize: eachPageSize,
	})
}

// GetSinglePageCollection PageIndex is one based index. Should be above or equal 1
func (it *DynamicCollection) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *DynamicCollection {
	length := it.Length()

	if length < eachPageSize {
		return it
	}

	pageInfo := it.GetPagingInfo(
		eachPageSize,
		pageIndex)

	list := it.items[pageInfo.SkipItems:pageInfo.EndingLength]

	return &DynamicCollection{
		items: list,
	}
}

func (it *DynamicCollection) JsonModel() DynamicCollectionModel {
	return DynamicCollectionModel{
		Items: it.items,
	}
}

func (it *DynamicCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it DynamicCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it DynamicCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *DynamicCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*DynamicCollection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

func (it *DynamicCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *DynamicCollection {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *DynamicCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *DynamicCollection) Strings() []string {
	slice := make([]string, it.Length())

	if it.IsEmpty() {
		return slice
	}

	for i, item := range it.items {
		slice[i] = item.String()
	}

	return slice
}

func (it *DynamicCollection) String() string {
	return strings.Join(it.Strings(), constants.NewLineUnix)
}
