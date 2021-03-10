package corejson

import (
	"encoding/json"
	"errors"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type ResultsCollection struct {
	Items *[]*Result `json:"JsonResultsCollection"`
}

func (receiver *ResultsCollection) Length() int {
	return len(*receiver.Items)
}

func (receiver *ResultsCollection) IsEmpty() bool {
	return receiver.Items == nil ||
		len(*receiver.Items) == 0
}

func (receiver *ResultsCollection) HasItems() bool {
	return receiver.Items != nil &&
		len(*receiver.Items) > 0
}

// skip on nil
func (receiver *ResultsCollection) AddSkipOnNil(
	result *Result,
) *ResultsCollection {
	if result == nil {
		return receiver
	}

	*receiver.Items = append(
		*receiver.Items,
		result)

	return receiver
}

func (receiver *ResultsCollection) GetAt(
	index int,
) *Result {
	return (*receiver.Items)[index]
}

// has any error
func (receiver *ResultsCollection) HasError() bool {
	for _, result := range *receiver.Items {
		if result != nil && result.Error != nil {
			return true
		}
	}

	return false
}

func (receiver *ResultsCollection) AllErrors() (
	errListPtr *[]error,
	hasAnyError bool,
) {
	length := receiver.Length()
	errList := make(
		[]error,
		0,
		length)

	if length == 0 {
		return &errList, hasAnyError
	}

	for i := 0; i < length; i++ {
		err := (*receiver.Items)[i].Error

		if err != nil {
			hasAnyError = true
			errList = append(
				errList,
				err)
		}
	}

	return &errList, hasAnyError
}

func (receiver *ResultsCollection) GetErrorsStrings() *[]string {
	length := receiver.Length()
	errStrList := make(
		[]string,
		0,
		length)

	if length == 0 {
		return &errStrList
	}

	for _, result := range *receiver.Items {
		if result == nil || result.IsEmptyError() {
			continue
		}

		errStrList = append(
			errStrList,
			result.Error.Error())
	}

	return &errStrList
}

func (receiver *ResultsCollection) GetErrorsAsSingleString() string {
	errStrList := receiver.GetErrorsStrings()

	return strings.Join(
		*errStrList,
		constants.NewLineUnix)
}

func (receiver *ResultsCollection) GetErrorsAsSingle() error {
	errorString := receiver.GetErrorsAsSingleString()

	return errors.New(errorString)
}

func (receiver *ResultsCollection) UnmarshalAt(
	index int,
	any interface{},
) error {
	result := (*receiver.Items)[index]

	if result == nil || result.IsEmptyJsonBytes() {
		return nil
	}

	if result.HasError() {
		return result.Error
	}

	if result.IsEmptyJsonBytes() {
		return nil
	}

	err := json.Unmarshal(
		*result.Bytes,
		any)

	return err
}

func (receiver *ResultsCollection) InjectIntoAt(
	index int,
	injector ParseSelfInjector,
) error {
	return injector.JsonParseSelfInject(
		(*receiver.Items)[index])
}

// any nil skip
func (receiver *ResultsCollection) InjectIntoSameIndex(
	injectors ...ParseSelfInjector,
) (
	errListPtr *[]error,
	hasAnyError bool,
) {
	if injectors == nil {
		return &[]error{}, false
	}

	length := len(injectors)
	errList := make([]error, length)

	for i := 0; i < length; i++ {
		result := (*receiver.Items)[i]
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

	return &errList, hasAnyError
}

// any nil skip
func (receiver *ResultsCollection) UnmarshalIntoSameIndex(
	anys ...interface{},
) (
	errListPtr *[]error,
	hasAnyError bool,
) {
	if anys == nil {
		return &[]error{}, false
	}

	length := len(anys)
	errList := make([]error, length)

	for i := 0; i < length; i++ {
		result := (*receiver.Items)[i]
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

		err := json.Unmarshal(
			*result.Bytes,
			any)

		if err != nil {
			hasAnyError = true
		}

		errList[i] = err
	}

	return &errList, hasAnyError
}

func (receiver *ResultsCollection) GetAtSafe(
	index int,
) *Result {
	if index > constants.InvalidNotFoundCase && index <= receiver.Length()-1 {
		return (*receiver.Items)[index]
	}

	return nil
}

func (receiver *ResultsCollection) GetAtSafeUsingLength(
	index, length int,
) *Result {
	if index > constants.InvalidNotFoundCase && index <= length-1 {
		return (*receiver.Items)[index]
	}

	return nil
}

func (receiver *ResultsCollection) Add(
	result *Result,
) *ResultsCollection {
	*receiver.Items = append(
		*receiver.Items,
		result)

	return receiver
}

func (receiver *ResultsCollection) Adds(
	results ...*Result,
) *ResultsCollection {
	if results == nil {
		return receiver
	}

	for _, result := range results {
		*receiver.Items = append(
			*receiver.Items,
			result)
	}

	return receiver
}

// Skip on nil
func (receiver *ResultsCollection) AddsAnys(
	anys ...interface{},
) *ResultsCollection {
	if anys == nil {
		return receiver
	}

	return receiver.AddsAnysPtr(&anys)
}

// Skip on nil
func (receiver *ResultsCollection) AddsAnysPtr(
	anysPtr *[]interface{},
) *ResultsCollection {
	if anysPtr == nil {
		return receiver
	}

	for _, any := range *anysPtr {
		if any == nil {
			continue
		}

		*receiver.Items = append(
			*receiver.Items,
			NewFromAny(any))
	}

	return receiver
}

// skip on nil items
func (receiver *ResultsCollection) AddResultsCollection(
	collection *ResultsCollection,
) *ResultsCollection {
	if collection == nil {
		return receiver
	}

	return receiver.AddNonNilItemsPtr(collection.Items)
}

// skip on nil
func (receiver *ResultsCollection) AddNonNilItems(
	results ...*Result,
) *ResultsCollection {
	if results == nil {
		return receiver
	}

	for _, result := range results {
		if result == nil {
			continue
		}

		*receiver.Items = append(
			*receiver.Items,
			result)

	}

	return receiver
}

// skip on nil
func (receiver *ResultsCollection) AddNonNilItemsPtr(
	results *[]*Result,
) *ResultsCollection {
	if results == nil {
		return receiver
	}

	for _, result := range *results {
		if result == nil {
			continue
		}

		*receiver.Items = append(
			*receiver.Items,
			result)

	}

	return receiver
}

func (receiver *ResultsCollection) Clear() *ResultsCollection {
	clearedItems := (*receiver.Items)[:0]
	receiver.Items = &clearedItems

	return receiver
}

func (receiver *ResultsCollection) GetStrings() *[]string {
	length := receiver.Length()
	list := make([]string, length)

	if length == 0 {
		return &list
	}

	for i, result := range *receiver.Items {
		list[i] = *result.JsonStringPtr()
	}

	return &list
}

// skip on nil
func (receiver *ResultsCollection) AddJsoner(
	jsoners ...Jsoner,
) *ResultsCollection {
	if jsoners == nil {
		return receiver
	}

	return receiver.AddJsonerPtr(&jsoners)
}

// skip on nil
func (receiver *ResultsCollection) AddJsonerPtr(
	jsoners *[]Jsoner,
) *ResultsCollection {
	if jsoners == nil {
		return receiver
	}

	for _, jsoner := range *jsoners {
		if jsoner == nil {
			continue
		}

		result := jsoner.Json()

		if result == nil {
			continue
		}

		*receiver.Items = append(
			*receiver.Items,
			result)
	}

	return receiver
}

//goland:noinspection GoLinterLocal
func (receiver *ResultsCollection) JsonModel() *ResultsCollection {
	return receiver
}

//goland:noinspection GoLinterLocal
func (receiver *ResultsCollection) JsonModelAny() interface{} {
	return receiver.JsonModel()
}

//goland:noinspection GoLinterLocal
func (receiver *ResultsCollection) Json() *Result {
	if receiver.IsEmpty() {
		return EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(receiver)

	return NewPtr(jsonBytes, err)
}

// It will not update the self but creates a new one.
func (receiver *ResultsCollection) ParseInjectUsingJson(
	jsonResult *Result,
) (*ResultsCollection, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return EmptyResultsCollection(), nil
	}

	err := json.Unmarshal(
		*jsonResult.Bytes,
		&receiver,
	)

	if err != nil {
		return EmptyResultsCollection(), err
	}

	return receiver, nil
}

// Panic if error
func (receiver *ResultsCollection) ParseInjectUsingJsonMust(
	jsonResult *Result,
) *ResultsCollection {
	resultCollection, err := receiver.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return resultCollection
}

func (receiver *ResultsCollection) AsJsoner() *Jsoner {
	var jsoner Jsoner = receiver

	return &jsoner
}

func (receiver *ResultsCollection) JsonParseSelfInject(
	jsonResult *Result,
) error {
	_, err := receiver.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (receiver *ResultsCollection) AsJsonParseSelfInjector() *ParseSelfInjector {
	var jsonMarshaller ParseSelfInjector = receiver

	return &jsonMarshaller
}
