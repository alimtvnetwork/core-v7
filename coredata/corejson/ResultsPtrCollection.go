package corejson

import (
	"errors"
	"math"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/core/errcore"
)

type ResultsPtrCollection struct {
	Items []*Result `json:"JsonResultsCollection"`
}

func (it *ResultsPtrCollection) Length() int {
	if it == nil || it.Items == nil {
		return 0
	}

	return len(it.Items)
}

func (it *ResultsPtrCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *ResultsPtrCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *ResultsPtrCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *ResultsPtrCollection) FirstOrDefault() *Result {
	if it.IsEmpty() {
		return nil
	}

	return it.Items[0]
}

func (it *ResultsPtrCollection) LastOrDefault() *Result {
	if it.IsEmpty() {
		return nil
	}

	return it.Items[it.LastIndex()]
}

func (it *ResultsPtrCollection) Take(limit int) *ResultsPtrCollection {
	if it.IsEmpty() {
		return EmptyResultsPtrCollection()
	}

	return &ResultsPtrCollection{
		Items: it.Items[:limit],
	}
}

func (it *ResultsPtrCollection) Limit(limit int) *ResultsPtrCollection {
	if it.IsEmpty() {
		return EmptyResultsPtrCollection()
	}

	if limit <= constants.TakeAllMinusOne {
		return it
	}

	limit = defaultcapacity.
		MaxLimit(it.Length(), limit)

	return &ResultsPtrCollection{
		Items: it.Items[:limit],
	}
}

func (it *ResultsPtrCollection) Skip(skip int) *ResultsPtrCollection {
	if it.IsEmpty() {
		return EmptyResultsPtrCollection()
	}

	return &ResultsPtrCollection{
		Items: it.Items[skip:],
	}
}

// AddSkipOnNil skip on nil
func (it *ResultsPtrCollection) AddSkipOnNil(
	result *Result,
) *ResultsPtrCollection {
	if result == nil {
		return it
	}

	it.Items = append(
		it.Items,
		result)

	return it
}

func (it *ResultsPtrCollection) AddNonNilNonError(
	result *Result,
) *ResultsPtrCollection {
	if result == nil || result.HasError() {
		return it
	}

	it.Items = append(
		it.Items,
		result)

	return it
}

func (it *ResultsPtrCollection) GetAt(
	index int,
) *Result {
	return it.Items[index]
}

// HasError has any error
func (it *ResultsPtrCollection) HasError() bool {
	for _, result := range it.Items {
		if result != nil && result.Error != nil {
			return true
		}
	}

	return false
}

func (it *ResultsPtrCollection) AllErrors() (
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

func (it *ResultsPtrCollection) GetErrorsStrings() []string {
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

func (it *ResultsPtrCollection) GetErrorsStringsPtr() *[]string {
	errStrList := it.GetErrorsStrings()

	return &errStrList
}

func (it *ResultsPtrCollection) GetErrorsAsSingleString() string {
	errStrList := it.GetErrorsStrings()

	return strings.Join(
		errStrList,
		constants.NewLineUnix)
}

func (it *ResultsPtrCollection) GetErrorsAsSingle() error {
	errorString := it.GetErrorsAsSingleString()

	return errors.New(errorString)
}

func (it *ResultsPtrCollection) UnmarshalAt(
	index int,
	any interface{},
) error {
	result := it.Items[index]

	if result == nil || result.IsEmptyJsonBytes() {
		return nil
	}

	if result.HasError() {
		return result.MeaningfulError()
	}

	if result.IsEmptyJsonBytes() {
		return nil
	}

	return result.Unmarshal(
		any)
}

func (it *ResultsPtrCollection) InjectIntoAt(
	index int,
	injector JsonParseSelfInjector,
) error {
	return injector.JsonParseSelfInject(
		it.Items[index])
}

// InjectIntoSameIndex any nil skip
func (it *ResultsPtrCollection) InjectIntoSameIndex(
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

		if result == nil {
			continue
		}

		if result.HasError() {
			hasAnyError = true

			continue
		}

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
func (it *ResultsPtrCollection) UnmarshalIntoSameIndex(
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

		if result == nil ||
			any == nil {
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

func (it *ResultsPtrCollection) GetAtSafe(
	index int,
) *Result {
	if index > constants.InvalidNotFoundCase && index <= it.Length()-1 {
		return it.Items[index]
	}

	return nil
}

func (it *ResultsPtrCollection) GetAtSafeUsingLength(
	index, length int,
) *Result {
	if index > constants.InvalidNotFoundCase && index <= length-1 {
		return it.Items[index]
	}

	return nil
}

func (it *ResultsPtrCollection) Add(
	result *Result,
) *ResultsPtrCollection {
	it.Items = append(
		it.Items,
		result)

	return it
}

func (it *ResultsPtrCollection) AddResult(
	result Result,
) *ResultsPtrCollection {
	it.Items = append(
		it.Items,
		&result)

	return it
}

func (it *ResultsPtrCollection) Adds(
	results ...*Result,
) *ResultsPtrCollection {
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

func (it *ResultsPtrCollection) AddAny(
	any interface{},
) *ResultsPtrCollection {
	if any == nil {
		return it
	}

	it.Items = append(
		it.Items,
		NewFromAnyPtr(any))

	return it
}

// AddsAnysPtr Skip on nil
func (it *ResultsPtrCollection) AddsAnysPtr(
	anys ...interface{},
) *ResultsPtrCollection {
	if anys == nil {
		return it
	}

	for _, any := range anys {
		if any == nil {
			continue
		}

		it.Items = append(
			it.Items,
			NewFromAnyPtr(any))
	}

	return it
}

// AddResultsCollection skip on nil items
func (it *ResultsPtrCollection) AddResultsCollection(
	collection *ResultsPtrCollection,
) *ResultsPtrCollection {
	if collection == nil {
		return it
	}

	return it.AddNonNilItemsPtr(collection.Items...)
}

// AddNonNilItems skip on nil
func (it *ResultsPtrCollection) AddNonNilItems(
	results ...*Result,
) *ResultsPtrCollection {
	if results == nil {
		return it
	}

	for _, result := range results {
		if result == nil {
			continue
		}

		it.Items = append(
			it.Items,
			results...)
	}

	return it
}

// AddNonNilItemsPtr skip on nil
func (it *ResultsPtrCollection) AddNonNilItemsPtr(
	results ...*Result,
) *ResultsPtrCollection {
	if results == nil || len(results) == 0 {
		return it
	}

	for _, result := range results {
		if result == nil {
			continue
		}

		it.Items = append(
			it.Items,
			result)
	}

	return it
}

func (it *ResultsPtrCollection) Clear() *ResultsPtrCollection {
	temp := it.Items
	clearFunc := func() {
		for _, result := range temp {
			result.Dispose()
			result = nil
		}
	}

	go clearFunc()
	it.Items = []*Result{}

	return it
}

func (it *ResultsPtrCollection) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.Items = nil
}

func (it *ResultsPtrCollection) GetStrings() []string {
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

func (it *ResultsPtrCollection) GetStringsPtr() *[]string {
	list := it.GetStrings()

	return &list
}

// AddJsoners skip on nil
func (it *ResultsPtrCollection) AddJsoners(
	isIgnoreNilOrError bool,
	jsoners ...Jsoner,
) *ResultsPtrCollection {
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
			&result)
	}

	return it
}

func (it *ResultsPtrCollection) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *ResultsPtrCollection) GetPagedCollection(
	eachPageSize int,
) []*ResultsPtrCollection {
	length := it.Length()

	if length < eachPageSize {
		return []*ResultsPtrCollection{
			it,
		}
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make([]*ResultsPtrCollection, pagesPossibleCeiling)

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
func (it *ResultsPtrCollection) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *ResultsPtrCollection {
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
			CannotBeNegativeIndex.
			HandleUsingPanic(
				"pageIndex cannot be negative or zero.",
				pageIndex)
	}

	endingIndex := skipItems + eachPageSize

	if endingIndex > length {
		endingIndex = length
	}

	list := it.Items[skipItems:endingIndex]

	return NewResultsCollectionPtrUsingJsonResultsPtr(
		constants.Zero,
		list...)
}

//goland:noinspection GoLinterLocal
func (it *ResultsPtrCollection) JsonModel() *ResultsPtrCollection {
	return it
}

//goland:noinspection GoLinterLocal
func (it *ResultsPtrCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it ResultsPtrCollection) Json() Result {
	return NewFromAny(it)
}

func (it ResultsPtrCollection) JsonPtr() *Result {
	return NewFromAnyPtr(it)
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *ResultsPtrCollection) ParseInjectUsingJson(
	jsonResult *Result,
) (*ResultsPtrCollection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return EmptyResultsPtrCollection(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *ResultsPtrCollection) ParseInjectUsingJsonMust(
	jsonResult *Result,
) *ResultsPtrCollection {
	resultCollection, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return resultCollection
}

func (it *ResultsPtrCollection) AsJsoner() Jsoner {
	return it
}

func (it *ResultsPtrCollection) JsonParseSelfInject(
	jsonResult *Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *ResultsPtrCollection) AsJsonParseSelfInjector() JsonParseSelfInjector {
	return it
}

func (it *ResultsPtrCollection) Clone(
	isDeepCloneEach bool,
) *ResultsPtrCollection {
	if it == nil {
		return nil
	}

	newResults := NewResultsPtrCollection(
		it.Length())

	if newResults.Length() == 0 {
		return newResults
	}

	for _, item := range it.Items {
		newResults.Add(item.ClonePtr(isDeepCloneEach))
	}

	return newResults
}
