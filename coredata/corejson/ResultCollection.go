package corejson

import (
	"errors"
	"strings"

	"gitlab.com/evatix-go/core/constants"
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
		return EmptyResultsCollection()
	}

	return &ResultsCollection{
		Items: it.Items[:limit],
	}
}

func (it *ResultsCollection) Limit(limit int) *ResultsCollection {
	if it.IsEmpty() {
		return EmptyResultsCollection()
	}

	return &ResultsCollection{
		Items: it.Items[:limit],
	}
}

func (it *ResultsCollection) Skip(skip int) *ResultsCollection {
	if it.IsEmpty() {
		return EmptyResultsCollection()
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

// AddsAnys Skip on nil
func (it *ResultsCollection) AddsAnys(
	anys ...interface{},
) *ResultsCollection {
	if anys == nil {
		return it
	}

	for _, any := range anys {
		if any == nil {
			continue
		}

		it.Items = append(
			it.Items,
			NewFromAny(any))
	}

	return it
}

// AddResultsCollection skip on nil items
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

func (it *ResultsCollection) Clear() *ResultsCollection {
	clearedItems := it.Items[:0]
	it.Items = clearedItems

	return it
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

//goland:noinspection GoLinterLocal
func (it *ResultsCollection) JsonModel() *ResultsCollection {
	return it
}

//goland:noinspection GoLinterLocal
func (it *ResultsCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it ResultsCollection) Json() Result {
	return NewFromAny(it)
}

func (it ResultsCollection) JsonPtr() *Result {
	return NewFromAnyPtr(it)
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *ResultsCollection) ParseInjectUsingJson(
	jsonResult *Result,
) (*ResultsCollection, error) {
	err := jsonResult.Unmarshal(
		&it,
	)

	if err != nil {
		return EmptyResultsCollection(), err
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

func (it *ResultsCollection) ShadowClone() *ResultsCollection {
	if it == nil {
		return nil
	}

	return it.Clone(false)
}

func (it *ResultsCollection) Clone(isDeepCloneEach bool) *ResultsCollection {
	if it == nil {
		return nil
	}

	newResults := NewResultsCollection(
		it.Length())

	if newResults.Length() == 0 {
		return newResults
	}

	for _, item := range it.Items {
		newResults.Add(*item.Clone(isDeepCloneEach))
	}

	return newResults
}
