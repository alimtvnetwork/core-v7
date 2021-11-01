package corejson

import (
	"encoding/json"
	"math"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/core/errcore"
)

type BytesCollection struct {
	Items [][]byte `json:"JsonBytesCollection"`
}

func (it *BytesCollection) Length() int {
	if it == nil || it.Items == nil {
		return 0
	}

	return len(it.Items)
}

func (it *BytesCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *BytesCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *BytesCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *BytesCollection) FirstOrDefault() []byte {
	if it.IsEmpty() {
		return nil
	}

	return it.Items[0]
}

func (it *BytesCollection) LastOrDefault() []byte {
	if it.IsEmpty() {
		return nil
	}

	return it.Items[it.LastIndex()]
}

func (it *BytesCollection) Take(limit int) *BytesCollection {
	if it.IsEmpty() {
		return EmptyBytesCollection()
	}

	return &BytesCollection{
		Items: it.Items[:limit],
	}
}

func (it *BytesCollection) Limit(limit int) *BytesCollection {
	if it.IsEmpty() {
		return EmptyBytesCollection()
	}

	if limit <= constants.TakeAllMinusOne {
		return it
	}

	limit = defaultcapacity.
		MaxLimit(it.Length(), limit)

	return &BytesCollection{
		Items: it.Items[:limit],
	}
}

func (it *BytesCollection) Skip(skip int) *BytesCollection {
	if it.IsEmpty() {
		return EmptyBytesCollection()
	}

	return &BytesCollection{
		Items: it.Items[skip:],
	}
}

// AddSkipOnNil skip on nil
func (it *BytesCollection) AddSkipOnNil(
	rawBytes []byte,
) *BytesCollection {
	if rawBytes == nil {
		return it
	}

	it.Items = append(
		it.Items,
		rawBytes)

	return it
}

// AddNonEmpty
//
// skip on empty
func (it *BytesCollection) AddNonEmpty(
	rawBytes []byte,
) *BytesCollection {
	if len(rawBytes) == 0 {
		return it
	}

	it.Items = append(
		it.Items,
		rawBytes)

	return it
}

// AddResultPtr
//
// skip on empty or has issue
func (it *BytesCollection) AddResultPtr(
	result *Result,
) *BytesCollection {
	if result.HasIssuesOrEmpty() {
		return it
	}

	it.Items = append(
		it.Items,
		result.Bytes)

	return it
}

// AddResult
//
// skip on empty or has issue
func (it *BytesCollection) AddResult(
	result Result,
) *BytesCollection {
	if result.HasIssuesOrEmpty() {
		return it
	}

	it.Items = append(
		it.Items,
		result.Bytes)

	return it
}

func (it *BytesCollection) GetAt(
	index int,
) []byte {
	return it.Items[index]
}

func (it *BytesCollection) JsonResultAt(
	index int,
) *Result {
	return &Result{
		Bytes: it.Items[index],
	}
}

func (it *BytesCollection) UnmarshalAt(
	index int,
	any interface{},
) error {
	rawBytes := it.Items[index]

	return json.Unmarshal(
		rawBytes,
		any)
}

func (it *BytesCollection) InjectIntoAt(
	index int,
	injector JsonParseSelfInjector,
) error {
	return injector.JsonParseSelfInject(
		it.JsonResultAt(index))
}

// InjectIntoSameIndex any nil skip
func (it *BytesCollection) InjectIntoSameIndex(
	injectors ...JsonParseSelfInjector,
) (
	errListPtr []error,
	hasAnyError bool,
) {
	if injectors == nil {
		return []error{}, false
	}

	length := len(injectors)
	errList := make([]error, length)

	for i := 0; i < length; i++ {
		result := it.JsonResultAt(i)
		injector := injectors[i]

		if injector == nil {
			continue
		}

		err := injector.
			JsonParseSelfInject(
				result)

		if err != nil {
			hasAnyError = true
		}

		errList[i] = err
	}

	return errList, hasAnyError
}

// UnmarshalIntoSameIndex any nil skip
func (it *BytesCollection) UnmarshalIntoSameIndex(
	anys ...interface{},
) (
	errListPtr []error,
	hasAnyError bool,
) {
	if anys == nil {
		return []error{}, false
	}

	length := len(anys)
	errList := make([]error, length)

	for i := 0; i < length; i++ {
		result := it.JsonResultAt(i)
		any := anys[i]

		if any == nil {
			continue
		}

		err := result.Unmarshal(
			any)

		if err != nil {
			hasAnyError = true
		}

		errList[i] = err
	}

	return errList, hasAnyError
}

func (it *BytesCollection) GetAtSafe(
	index int,
) []byte {
	if index > constants.InvalidNotFoundCase && index <= it.Length()-1 {
		return it.Items[index]
	}

	return nil
}

func (it *BytesCollection) GetAtSafePtr(
	index int,
) *[]byte {
	if index > constants.InvalidNotFoundCase && index <= it.Length()-1 {
		return &it.Items[index]
	}

	return nil
}

func (it *BytesCollection) GetResultAtSafe(
	index int,
) *Result {
	if index > constants.InvalidNotFoundCase && index <= it.Length()-1 {
		return it.JsonResultAt(index)
	}

	return nil
}

func (it *BytesCollection) GetAtSafeUsingLength(
	index, length int,
) *Result {
	if index > constants.InvalidNotFoundCase && index <= length-1 {
		return it.JsonResultAt(index)
	}

	return nil
}

func (it *BytesCollection) AddPtr(
	rawBytes *[]byte,
) *BytesCollection {
	if rawBytes == nil || len(*rawBytes) == 0 {
		return it
	}

	it.Items = append(
		it.Items,
		*rawBytes)

	return it
}

func (it *BytesCollection) Add(
	result []byte,
) *BytesCollection {
	it.Items = append(
		it.Items,
		result)

	return it
}

func (it *BytesCollection) Adds(
	rawBytesCollection ...[]byte,
) *BytesCollection {
	if len(rawBytesCollection) == 0 {
		return it
	}

	for _, rawBytes := range rawBytesCollection {
		if len(rawBytes) == 0 {
			continue
		}

		it.Items = append(
			it.Items,
			rawBytes)
	}

	return it
}

func (it *BytesCollection) AddAnyItems(
	anyItems ...interface{},
) error {
	if len(anyItems) == 0 {
		return nil
	}

	for _, anyItem := range anyItems {
		jsonResult := NewFromAnyPtr(anyItem)
		if jsonResult.HasError() {
			return jsonResult.MeaningfulError()
		}

		it.Items = append(
			it.Items,
			jsonResult.Bytes)
	}

	return nil
}

func (it *BytesCollection) AddMapResults(
	mapResults *MapResults,
) *BytesCollection {
	if mapResults.IsEmpty() {
		return it
	}

	return it.AddRawMapResults(mapResults.Items)
}

func (it *BytesCollection) AddRawMapResults(
	mapResults map[string]Result,
) *BytesCollection {
	if len(mapResults) == 0 {
		return it
	}

	for _, result := range mapResults {
		if result.HasError() {
			continue
		}

		it.Items = append(
			it.Items,
			result.Bytes)
	}

	return it
}

func (it *BytesCollection) AddsPtr(
	results ...*Result,
) *BytesCollection {
	if results == nil {
		return it
	}

	for _, result := range results {
		if result.IsAnyNull() {
			continue
		}

		it.Items = append(
			it.Items,
			result.Bytes)
	}

	return it
}

func (it *BytesCollection) AddAny(
	any interface{},
) error {
	result := NewFromAny(any)

	if result.HasError() {
		return result.MeaningfulError()
	}

	it.Items = append(
		it.Items,
		result.Bytes)

	return nil
}

// AddBytesCollection skip on nil items
func (it *BytesCollection) AddBytesCollection(
	collection *BytesCollection,
) *BytesCollection {
	if collection.IsEmpty() {
		return it
	}

	return it.Adds(collection.Items...)
}

func (it *BytesCollection) Clear() *BytesCollection {
	if it == nil {
		return it
	}

	tempItems := it.Items
	clearFunc := func() {
		for i := range tempItems {
			tempItems[i] = nil
		}
	}

	go clearFunc()
	it.Items = [][]byte{}

	return it
}

func (it *BytesCollection) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.Items = nil
}

func (it *BytesCollection) Strings() []string {
	length := it.Length()
	list := make([]string, length)

	if length == 0 {
		return list
	}

	for i, rawBytes := range it.Items {
		list[i] = string(rawBytes)
	}

	return list
}

func (it *BytesCollection) StringsPtr() *[]string {
	list := it.Strings()

	return &list
}

// AddJsoners skip on nil
func (it *BytesCollection) AddJsoners(
	isIgnoreNilOrError bool,
	jsoners ...Jsoner,
) *BytesCollection {
	if jsoners == nil {
		return it
	}

	for _, jsoner := range jsoners {
		if jsoner == nil {
			continue
		}

		result := jsoner.Json()

		if isIgnoreNilOrError && result.HasError() {
			continue
		}

		it.Items = append(
			it.Items,
			result.Bytes)
	}

	return it
}

func (it *BytesCollection) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *BytesCollection) GetPagedCollection(
	eachPageSize int,
) []*BytesCollection {
	length := it.Length()

	if length < eachPageSize {
		return []*BytesCollection{
			it,
		}
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make([]*BytesCollection, pagesPossibleCeiling)

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
func (it *BytesCollection) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *BytesCollection {
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

	return &BytesCollection{
		Items: list,
	}
}

//goland:noinspection GoLinterLocal
func (it *BytesCollection) JsonModel() *BytesCollection {
	return it
}

//goland:noinspection GoLinterLocal
func (it *BytesCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it BytesCollection) Json() Result {
	return NewFromAny(it)
}

func (it BytesCollection) JsonPtr() *Result {
	return NewFromAnyPtr(it)
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *BytesCollection) ParseInjectUsingJson(
	jsonResult *Result,
) (*BytesCollection, error) {
	err := jsonResult.Unmarshal(
		&it,
	)

	if err != nil {
		return EmptyBytesCollection(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *BytesCollection) ParseInjectUsingJsonMust(
	jsonResult *Result,
) *BytesCollection {
	resultCollection, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return resultCollection
}

func (it *BytesCollection) AsJsonContractsBinder() JsonContractsBinder {
	return it
}

func (it *BytesCollection) AsJsoner() Jsoner {
	return it
}

func (it *BytesCollection) JsonParseSelfInject(
	jsonResult *Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *BytesCollection) AsJsonParseSelfInjector() JsonParseSelfInjector {
	return it
}

func (it *BytesCollection) ShadowClone() BytesCollection {
	return it.Clone(false)
}

func (it BytesCollection) Clone(isDeepCloneEach bool) BytesCollection {
	newResults := NewBytesCollection(
		it.Length())

	if newResults.Length() == 0 {
		return *newResults
	}

	for _, item := range it.Items {
		newResults.Add(BytesCloneIf(isDeepCloneEach, item))
	}

	return *newResults
}

func (it *BytesCollection) ClonePtr(isDeepCloneEach bool) *BytesCollection {
	if it == nil {
		return nil
	}

	newResults := NewBytesCollection(
		it.Length())

	if newResults.Length() == 0 {
		return newResults
	}

	for _, item := range it.Items {
		newResults.Add(BytesCloneIf(isDeepCloneEach, item))
	}

	return newResults
}
