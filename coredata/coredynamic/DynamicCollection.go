package coredynamic

import (
	"encoding/json"
	"math"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/msgtype"
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

func (it *DynamicCollection) AddAny(anyItem interface{}, isValid bool) *DynamicCollection {
	it.items = append(
		it.items,
		NewDynamic(anyItem, isValid))

	return it
}

func (it *DynamicCollection) AddAnyNonNull(anyItem interface{}, isValid bool) *DynamicCollection {
	if anyItem == nil {
		return it
	}

	it.items = append(
		it.items,
		NewDynamic(anyItem, isValid))

	return it
}

func (it *DynamicCollection) AddAnyMany(anyItems ...interface{}) *DynamicCollection {
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

func (it *DynamicCollection) Add(dynamic Dynamic) *DynamicCollection {
	it.items = append(it.items, dynamic)

	return it
}

func (it *DynamicCollection) AddPtr(dynamic *Dynamic) *DynamicCollection {
	if dynamic == nil {
		return it
	}

	it.items = append(it.items, *dynamic)

	return it
}

func (it *DynamicCollection) AddManyPtr(dynamicItems ...*Dynamic) *DynamicCollection {
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
		msgtype.
			MarshallingFailed.
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
	jsonResultsCollection := corejson.NewResultsCollection(it.Length())

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
	jsonResultsCollection := corejson.NewResultsPtrCollection(it.Length())

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
	return corejson.NewFromAny(it)
}

func (it DynamicCollection) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
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
