package corejson

import (
	"errors"
	"math"
	"strings"
	"sync"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/defaultcapacity"
	"gitlab.com/auk-go/core/errcore"
)

type ResultsCollection struct {
	Items []Result `json:"JsonResultsCollection"`
}

func (it *ResultsCollection) Length() int {
	if it == nil || it.Items == nil {
		return 0
	}

	return len(it.Items)
}

func (it *ResultsCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *ResultsCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *ResultsCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *ResultsCollection) FirstOrDefault() *Result {
	if it.IsEmpty() {
		return nil
	}

	return &it.Items[0]
}

func (it *ResultsCollection) LastOrDefault() *Result {
	if it.IsEmpty() {
		return nil
	}

	return &it.Items[it.LastIndex()]
}

func (it *ResultsCollection) Take(limit int) *ResultsCollection {
	if it.IsEmpty() {
		return Empty.ResultsCollection()
	}

	return &ResultsCollection{
		Items: it.Items[:limit],
	}
}

func (it *ResultsCollection) Limit(limit int) *ResultsCollection {
	if it.IsEmpty() {
		return Empty.ResultsCollection()
	}

	if limit <= constants.TakeAllMinusOne {
		return it
	}

	limit = defaultcapacity.
		MaxLimit(it.Length(), limit)

	return &ResultsCollection{
		Items: it.Items[:limit],
	}
}

func (it *ResultsCollection) Skip(skip int) *ResultsCollection {
	if it.IsEmpty() {
		return Empty.ResultsCollection()
	}

	return &ResultsCollection{
		Items: it.Items[skip:],
	}
}

// AddSkipOnNil skip on nil
func (it *ResultsCollection) AddSkipOnNil(
	result *Result,
) *ResultsCollection {
	if result == nil {
		return it
	}

	it.Items = append(
		it.Items,
		*result)

	return it
}

func (it *ResultsCollection) AddNonNilNonError(
	result *Result,
) *ResultsCollection {
	if result == nil || result.HasError() {
		return it
	}

	it.Items = append(
		it.Items,
		*result)

	return it
}

func (it *ResultsCollection) GetAt(
	index int,
) *Result {
	return &it.Items[index]
}

// HasError has any error
func (it *ResultsCollection) HasError() bool {
	for _, result := range it.Items {
		if result.HasError() {
			return true
		}
	}

	return false
}

func (it *ResultsCollection) AllErrors() (
	errListPtr []error,
	hasAnyError bool,
) {
	length := it.Length()
	errList := make(
		[]error,
		0,
		length)

	if length == 0 {
		return errList, hasAnyError
	}

	for i := 0; i < length; i++ {
		err := it.Items[i].Error

		if err != nil {
			hasAnyError = true
			errList = append(
				errList,
				err)
		}
	}

	return errList, hasAnyError
}

func (it *ResultsCollection) GetErrorsStrings() []string {
	length := it.Length()
	errStrList := make(
		[]string,
		0,
		length)

	if length == 0 {
		return errStrList
	}

	for _, result := range it.Items {
		if result.IsEmptyError() {
			continue
		}

		errStrList = append(
			errStrList,
			result.Error.Error())
	}

	return errStrList
}

func (it *ResultsCollection) GetErrorsStringsPtr() *[]string {
	errStrList := it.GetErrorsStrings()

	return &errStrList
}

func (it *ResultsCollection) GetErrorsAsSingleString() string {
	errStrList := it.GetErrorsStrings()

	return strings.Join(
		errStrList,
		constants.NewLineUnix)
}

func (it *ResultsCollection) GetErrorsAsSingle() error {
	errorString := it.GetErrorsAsSingleString()

	return errors.New(errorString)
}

func (it *ResultsCollection) UnmarshalAt(
	index int,
	any interface{},
) error {
	result := it.Items[index]

	return result.Unmarshal(
		any)
}

func (it *ResultsCollection) InjectIntoAt(
	index int,
	injector JsonParseSelfInjector,
) error {
	return injector.JsonParseSelfInject(
		&it.Items[index])
}

// InjectIntoSameIndex any nil skip
func (it *ResultsCollection) InjectIntoSameIndex(
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
		result := it.Items[i]
		injector := injectors[i]

		if result.HasError() {
			hasAnyError = true

			continue
		}

		if injector == nil {
			continue
		}

		err := injector.
			JsonParseSelfInject(
				&result)

		if err != nil {
			hasAnyError = true
		}

		errList[i] = err
	}

	return errList, hasAnyError
}

// UnmarshalIntoSameIndex any nil skip
func (it *ResultsCollection) UnmarshalIntoSameIndex(
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
		result := it.Items[i]
		any := anys[i]

		if any == nil {
			continue
		}

		if result.HasError() {
			hasAnyError = true
			errList[i] = result.Error

			continue
		}

		if result.IsEmptyJsonBytes() {
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

func (it *ResultsCollection) GetAtSafe(
	index int,
) *Result {
	if index > constants.InvalidNotFoundCase && index <= it.Length()-1 {
		return &it.Items[index]
	}

	return nil
}

func (it *ResultsCollection) GetAtSafeUsingLength(
	index, length int,
) *Result {
	if index > constants.InvalidNotFoundCase && index <= length-1 {
		return &it.Items[index]
	}

	return nil
}

func (it *ResultsCollection) AddPtr(
	result *Result,
) *ResultsCollection {
	if result == nil {
		return it
	}

	it.Items = append(
		it.Items,
		*result)

	return it
}

func (it *ResultsCollection) Add(
	result Result,
) *ResultsCollection {
	it.Items = append(
		it.Items,
		result)

	return it
}

func (it *ResultsCollection) Adds(
	results ...Result,
) *ResultsCollection {
	if results == nil {
		return it
	}

	for _, result := range results {
		it.Items = append(
			it.Items,
			result)
	}

	return it
}

func (it *ResultsCollection) AddSerializer(
	serializer bytesSerializer,
) *ResultsCollection {
	if serializer == nil {
		return it
	}

	result := NewResult.UsingSerializer(
		serializer)

	return it.AddSkipOnNil(result)
}

func (it *ResultsCollection) AddSerializers(
	serializers ...bytesSerializer,
) *ResultsCollection {
	if len(serializers) == 0 {
		return it
	}

	for _, serializer := range serializers {
		it.AddSerializer(serializer)
	}

	return it
}

func (it *ResultsCollection) AddSerializerFunc(
	serializerFunc func() ([]byte, error),
) *ResultsCollection {
	if serializerFunc == nil {
		return it
	}

	result := NewResult.UsingSerializerFunc(
		serializerFunc)

	return it.AddSkipOnNil(result)
}

func (it *ResultsCollection) AddSerializerFunctions(
	serializerFunctions ...func() ([]byte, error),
) *ResultsCollection {
	if len(serializerFunctions) == 0 {
		return it
	}

	for _, serializer := range serializerFunctions {
		it.AddSerializerFunc(serializer)
	}

	return it
}

func (it *ResultsCollection) AddMapResults(
	mapResults *MapResults,
) *ResultsCollection {
	if mapResults.IsEmpty() {
		return it
	}

	return it.AddRawMapResults(mapResults.Items)
}

func (it *ResultsCollection) AddRawMapResults(
	mapResults map[string]Result,
) *ResultsCollection {
	if len(mapResults) == 0 {
		return it
	}

	for _, result := range mapResults {
		it.Items = append(
			it.Items,
			result)
	}

	return it
}

func (it *ResultsCollection) AddsPtr(
	results ...*Result,
) *ResultsCollection {
	if results == nil {
		return it
	}

	for _, result := range results {
		if result == nil {
			continue
		}

		it.Items = append(
			it.Items,
			*result)
	}

	return it
}

func (it *ResultsCollection) AddAny(
	any interface{},
) *ResultsCollection {
	if any == nil {
		return it
	}

	it.Items = append(
		it.Items,
		New(any))

	return it
}

// AddAnyItems Skip on nil
func (it *ResultsCollection) AddAnyItems(
	anyItems ...interface{},
) *ResultsCollection {
	if anyItems == nil {
		return it
	}

	for _, any := range anyItems {
		if any == nil {
			continue
		}

		it.Items = append(
			it.Items,
			New(any))
	}

	return it
}

// AddAnyItemsSlice
//
//	Skip on nil
func (it *ResultsCollection) AddAnyItemsSlice(
	anyItems []interface{},
) *ResultsCollection {
	if anyItems == nil {
		return it
	}

	for _, any := range anyItems {
		if any == nil {
			continue
		}

		it.Items = append(
			it.Items,
			New(any))
	}

	return it
}

// AddResultsCollection
//
//	skip on nil items
func (it *ResultsCollection) AddResultsCollection(
	collection *ResultsCollection,
) *ResultsCollection {
	if collection == nil {
		return it
	}

	return it.Adds(collection.Items...)
}

// AddNonNilItemsPtr skip on nil
func (it *ResultsCollection) AddNonNilItemsPtr(
	results ...*Result,
) *ResultsCollection {
	if results == nil || len(results) == 0 {
		return it
	}

	for _, result := range results {
		if result == nil {
			continue
		}

		it.Items = append(
			it.Items,
			*result)
	}

	return it
}

func (it ResultsCollection) NonPtr() ResultsCollection {
	return it
}

func (it *ResultsCollection) Ptr() *ResultsCollection {
	return it
}

func (it *ResultsCollection) Clear() *ResultsCollection {
	if it == nil {
		return it
	}

	temp := it.Items
	clearFunc := func() {
		for _, result := range temp {
			result.Dispose()
		}
	}

	go clearFunc()
	it.Items = []Result{}

	return it
}

func (it *ResultsCollection) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.Items = nil
}

func (it *ResultsCollection) GetStrings() []string {
	length := it.Length()
	list := make([]string, length)

	if length == 0 {
		return list
	}

	for i, result := range it.Items {
		list[i] = *result.JsonStringPtr()
	}

	return list
}

func (it *ResultsCollection) GetStringsPtr() *[]string {
	list := it.GetStrings()

	return &list
}

// AddJsoners skip on nil
func (it *ResultsCollection) AddJsoners(
	isIgnoreNilOrError bool,
	jsoners ...Jsoner,
) *ResultsCollection {
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
			result)
	}

	return it
}

func (it *ResultsCollection) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *ResultsCollection) GetPagedCollection(
	eachPageSize int,
) []*ResultsCollection {
	length := it.Length()

	if length < eachPageSize {
		return []*ResultsCollection{
			it,
		}
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make([]*ResultsCollection, pagesPossibleCeiling)

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
func (it *ResultsCollection) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *ResultsCollection {
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

	return NewResultsCollection.UsingResults(
		list...)
}

//goland:noinspection GoLinterLocal
func (it *ResultsCollection) JsonModel() *ResultsCollection {
	return it
}

//goland:noinspection GoLinterLocal
func (it *ResultsCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it ResultsCollection) Json() Result {
	return New(it)
}

func (it ResultsCollection) JsonPtr() *Result {
	return NewPtr(it)
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *ResultsCollection) ParseInjectUsingJson(
	jsonResult *Result,
) (*ResultsCollection, error) {
	err := jsonResult.Unmarshal(
		&it,
	)

	if err != nil {
		return Empty.ResultsCollection(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *ResultsCollection) ParseInjectUsingJsonMust(
	jsonResult *Result,
) *ResultsCollection {
	resultCollection, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return resultCollection
}

func (it *ResultsCollection) AsJsonContractsBinder() JsonContractsBinder {
	return it
}

func (it *ResultsCollection) AsJsoner() Jsoner {
	return it
}

func (it *ResultsCollection) JsonParseSelfInject(
	jsonResult *Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *ResultsCollection) AsJsonParseSelfInjector() JsonParseSelfInjector {
	return it
}

func (it ResultsCollection) ShadowClone() ResultsCollection {
	return it.Clone(false)
}

func (it ResultsCollection) Clone(isDeepCloneEach bool) ResultsCollection {
	newResults := NewResultsCollection.
		UsingCap(it.Length())

	if newResults.Length() == 0 {
		return *newResults
	}

	for _, item := range it.Items {
		newResults.Add(*item.ClonePtr(isDeepCloneEach))
	}

	return *newResults
}

func (it *ResultsCollection) ClonePtr(isDeepCloneEach bool) *ResultsCollection {
	if it == nil {
		return nil
	}

	newResults := NewResultsCollection.UsingCap(
		it.Length())

	if newResults.Length() == 0 {
		return newResults
	}

	for _, item := range it.Items {
		newResults.Add(*item.ClonePtr(isDeepCloneEach))
	}

	return newResults
}
