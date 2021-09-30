package corejson

import (
	"errors"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

type MapResults struct {
	Items map[string]Result `json:"JsonResultsMap"`
}

func (it *MapResults) Length() int {
	if it == nil || it.Items == nil {
		return 0
	}

	return len(it.Items)
}

func (it *MapResults) LastIndex() int {
	return it.Length() - 1
}

func (it *MapResults) IsEmpty() bool {
	return it.Length() == 0
}

func (it *MapResults) HasAnyItem() bool {
	return it.Length() > 0
}

// AddSkipOnNil skip on nil
func (it *MapResults) AddSkipOnNil(
	key string,
	result *Result,
) *MapResults {
	if result == nil {
		return it
	}

	it.Items[key] = *result

	return it
}

func (it *MapResults) GetAt(
	key string,
) *Result {
	r, has := it.Items[key]

	if has {
		return &r
	}

	return nil
}

// HasError has any error
func (it *MapResults) HasError() bool {
	for _, result := range it.Items {
		if result.HasError() {
			return true
		}
	}

	return false
}

func (it *MapResults) AllErrors() (
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

	for key, val := range it.Items {
		err := val.Error

		if err != nil {
			hasAnyError = true
			errList = append(
				errList,
				errors.New(key+constants.HypenAngelRight+err.Error()))
		}
	}

	return errList, hasAnyError
}

func (it *MapResults) GetErrorsStrings() []string {
	length := it.Length()
	errStrList := make(
		[]string,
		0,
		length)

	if length == 0 {
		return errStrList
	}

	for key, result := range it.Items {
		if result.IsEmptyError() {
			continue
		}

		errStrList = append(
			errStrList,
			key+constants.HypenAngelRight+result.Error.Error())
	}

	return errStrList
}

func (it *MapResults) GetErrorsStringsPtr() *[]string {
	errStrList := it.GetErrorsStrings()

	return &errStrList
}

func (it *MapResults) GetErrorsAsSingleString() string {
	errStrList := it.GetErrorsStrings()

	return strings.Join(
		errStrList,
		constants.NewLineUnix)
}

func (it *MapResults) GetErrorsAsSingle() error {
	errorString := it.GetErrorsAsSingleString()

	return errors.New(errorString)
}

func (it *MapResults) UnmarshalAt(
	key string,
	any interface{},
) error {
	result, has := it.Items[key]

	if has {
		return msgtype.
			KeyNotExistInMap.
			Error("Given key not found!", key)
	}

	if result.IsEmptyJsonBytes() {
		return msgtype.
			EmptyResultCannotMakeJson.
			Error("Cannot make json of empty bytes!", key)
	}

	return result.Unmarshal(
		any)
}

func (it *MapResults) UnmarshalManys(
	keyAnys ...KeyAny,
) error {
	if len(keyAnys) == 0 {
		return nil
	}

	for _, keyAny := range keyAnys {
		err := it.UnmarshalAt(
			keyAny.Key,
			keyAny.AnyInf)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *MapResults) UnmarshalManysSafe(
	keyAnys ...KeyAny,
) error {
	if len(keyAnys) == 0 {
		return nil
	}

	for _, keyAny := range keyAnys {
		err := it.SafeUnmarshalAt(
			keyAny.Key,
			keyAny.AnyInf)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *MapResults) SafeUnmarshalAt(
	key string,
	any interface{},
) error {
	result, has := it.Items[key]

	if has || result.IsEmptyJsonBytes() {
		return nil
	}

	return result.Unmarshal(
		any)
}

func (it *MapResults) InjectIntoAt(
	key string,
	injector JsonParseSelfInjector,
) error {
	return injector.JsonParseSelfInject(
		it.GetAt(key))
}

func (it *MapResults) GetAtSafe(
	key string,
) *Result {
	return it.GetAt(key)
}

func (it *MapResults) Add(
	key string,
	result Result,
) *MapResults {
	it.Items[key] = result

	return it
}

func (it *MapResults) AddPtr(
	key string,
	result *Result,
) *MapResults {
	if result == nil {
		return it
	}

	it.Items[key] = *result

	return it
}

// AddAny returns error if any during marshalling it.
func (it *MapResults) AddAny(
	key string,
	item interface{},
) error {
	if item == nil {
		return msgtype.MarshallingFailed.Error(
			msgtype.CannotBeNilMessage.String(),
			key)
	}

	jsonResult := NewFromAny(item)

	if jsonResult.HasError() {
		return jsonResult.MeaningfulError()
	}

	it.Add(key, jsonResult)

	return nil
}

// AddAnySkipOnNil returns error if any during marshalling it.
func (it *MapResults) AddAnySkipOnNil(
	key string,
	item interface{},
) error {
	if item == nil {
		return nil
	}

	jsonResult := NewFromAny(item)

	if jsonResult.HasError() {
		return jsonResult.MeaningfulError()
	}

	it.Add(key, jsonResult)

	return nil
}

func (it *MapResults) AddAnyNonEmptyNonError(
	key string,
	item interface{},
) *MapResults {
	if item == nil {
		return it
	}

	return it.AddNonEmptyNonErrorPtr(
		key,
		NewFromAnyPtr(item))
}

func (it *MapResults) AddAnyNonEmpty(
	key string,
	item interface{},
) *MapResults {
	if item == nil {
		return it
	}

	return it.AddPtr(
		key,
		NewFromAnyPtr(item))
}

func (it *MapResults) AddKeyWithResult(
	result KeyWithResult,
) *MapResults {
	return it.AddPtr(result.Key, &result.Result)
}

func (it *MapResults) AddKeyWithResultPtr(
	result *KeyWithResult,
) *MapResults {
	if result == nil {
		return it
	}

	return it.AddPtr(result.Key, &result.Result)
}

func (it *MapResults) AddKeysWithResultsPtr(
	results ...*KeyWithResult,
) *MapResults {
	if len(results) == 0 {
		return it
	}

	for _, result := range results {
		it.AddKeyWithResultPtr(result)
	}

	return it
}

func (it *MapResults) AddKeysWithResults(
	results ...KeyWithResult,
) *MapResults {
	if len(results) == 0 {
		return it
	}

	for _, result := range results {
		it.AddKeyWithResult(result)
	}

	return it
}

func (it *MapResults) AddKeyAnyInf(
	result KeyAny,
) *MapResults {
	return it.AddAnyNonEmpty(
		result.Key,
		result.AnyInf)
}

func (it *MapResults) AddKeyAnyInfPtr(
	result *KeyAny,
) *MapResults {
	if result == nil {
		return it
	}

	return it.AddAnyNonEmpty(
		result.Key,
		result.AnyInf)
}

func (it *MapResults) AddKeyAnys(
	results ...KeyAny,
) *MapResults {
	if results == nil {
		return it
	}

	for _, result := range results {
		it.AddKeyAnyInf(result)
	}

	return it
}

func (it *MapResults) AddKeyAnysPtr(
	results ...*KeyAny,
) *MapResults {
	if results == nil {
		return it
	}

	for _, result := range results {
		it.AddKeyAnyInfPtr(result)
	}

	return it
}

func (it *MapResults) AddNonEmptyNonErrorPtr(
	key string,
	result *Result,
) *MapResults {
	if result == nil || result.HasError() {
		return it
	}

	it.Items[key] = *result

	return it
}

func (it *MapResults) AddMapResults(
	mapResults *MapResults,
) *MapResults {
	if mapResults == nil || mapResults.IsEmpty() {
		return it
	}

	for key, r := range mapResults.Items {
		it.Items[key] = r
	}

	return it
}

func (it *MapResults) AllKeys() []string {
	if it.IsEmpty() {
		return []string{}
	}

	keys := make([]string, it.Length())

	index := 0
	for key := range it.Items {
		keys[index] = key
		index++
	}

	return keys
}

func (it *MapResults) AllValues() []Result {
	if it.IsEmpty() {
		return []Result{}
	}

	values := make([]Result, it.Length())

	index := 0
	for _, result := range it.Items {
		values[index] = result
		index++
	}

	return values
}

func (it *MapResults) AllResultsCollection() *ResultsCollection {
	if it.IsEmpty() {
		return EmptyResultsCollection()
	}

	resultsCollection := NewResultsCollection(it.Length())

	index := 0
	for _, result := range it.Items {
		resultsCollection.Add(result)
		index++
	}

	return resultsCollection
}

func (it *MapResults) AllResults() []Result {
	return it.AllValues()
}

func (it *MapResults) Clear() *MapResults {
	it.Items = make(
		map[string]Result,
		constants.Zero)

	return it
}

func (it *MapResults) GetStrings() []string {
	length := it.Length()
	list := make([]string, length)

	if length == 0 {
		return list
	}

	index := 0
	for _, result := range it.Items {
		list[index] = *result.JsonStringPtr()
		index++
	}

	return list
}

func (it *MapResults) GetStringsPtr() *[]string {
	stringsItems := it.GetStrings()

	return &stringsItems
}

// AddJsoner skip on nil
func (it *MapResults) AddJsoner(
	key string,
	jsoner Jsoner,
) *MapResults {
	if jsoner == nil {
		return it
	}

	return it.AddPtr(key, jsoner.JsonPtr())
}

func (it *MapResults) AddKeyWithJsoner(
	keyWithJsoner KeyWithJsoner,
) *MapResults {
	return it.AddJsoner(
		keyWithJsoner.Key,
		keyWithJsoner.Jsoner)
}

func (it *MapResults) AddKeysWithJsoners(
	keysWithJsoners ...KeyWithJsoner,
) *MapResults {
	if keysWithJsoners == nil {
		return nil
	}

	for _, jsoner := range keysWithJsoners {
		it.AddKeyWithJsoner(jsoner)
	}

	return it
}

func (it *MapResults) AddKeyWithJsonerPtr(
	keyWithJsoner *KeyWithJsoner,
) *MapResults {
	if keyWithJsoner == nil || keyWithJsoner.Jsoner == nil {
		return it
	}

	return it.AddJsoner(
		keyWithJsoner.Key,
		keyWithJsoner.Jsoner)
}

//goland:noinspection GoLinterLocal
func (it *MapResults) JsonModel() *MapResults {
	return it
}

//goland:noinspection GoLinterLocal
func (it *MapResults) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it MapResults) Json() Result {
	return NewFromAny(it)
}

func (it MapResults) JsonPtr() *Result {
	return NewFromAnyPtr(it)
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *MapResults) ParseInjectUsingJson(
	jsonResult *Result,
) (*MapResults, error) {
	err := jsonResult.Unmarshal(
		&it,
	)

	if err != nil {
		return EmptyMapResultsUsingCap(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *MapResults) ParseInjectUsingJsonMust(
	jsonResult *Result,
) *MapResults {
	resultCollection, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return resultCollection
}

func (it *MapResults) AsJsoner() Jsoner {
	return it
}

func (it *MapResults) JsonParseSelfInject(
	jsonResult *Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *MapResults) AsJsonParseSelfInjector() JsonParseSelfInjector {
	return it
}
