package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type DynamicFunc struct {
	Params   Map         `json:",omitempty"`
	WorkFunc interface{} `json:",omitempty"`
	Expect   interface{} `json:",omitempty"`
	toSlice  *[]interface{}
	toString corestr.SimpleStringOnce
}

func (it *DynamicFunc) ArgsCount() int {
	if it == nil {
		return 0
	}

	return it.Params.ArgsCount()
}

func (it *DynamicFunc) GetWorkFunc() interface{} {
	return it.WorkFunc
}

func (it *DynamicFunc) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Params)
}

func (it *DynamicFunc) HasFirst() bool {
	return reflectinternal.Is.Defined(it.FirstItem())
}

func (it *DynamicFunc) GetByIndex(index int) interface{} {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it *DynamicFunc) FirstItem() interface{} {
	return it.Params.FirstItem()
}

func (it *DynamicFunc) SecondItem() interface{} {
	return it.Params.SecondItem()
}

func (it *DynamicFunc) ThirdItem() interface{} {
	return it.Params.ThirdItem()
}

func (it *DynamicFunc) FourthItem() interface{} {
	return it.Params.FourthItem()
}

func (it *DynamicFunc) FifthItem() interface{} {
	return it.Params.FifthItem()
}

func (it *DynamicFunc) SixthItem() interface{} {
	return it.Params.SixthItem()
}

func (it *DynamicFunc) Expected() interface{} {
	return it.Expect
}

// HasDefined
//
// Confirms that key is present and defined.
func (it *DynamicFunc) HasDefined(name string) bool {
	if it == nil {
		return false
	}

	item, has := it.Params[name]

	return has &&
		reflectinternal.Is.Defined(item)
}

// Has
//
//	Confirms that key is present only.
//
//	Don't confirm not null.
//
// Use HasDefined to check not null.
func (it *DynamicFunc) Has(name string) bool {
	if it == nil {
		return false
	}

	_, has := it.Params[name]

	return has
}

// HasDefinedAll
//
// Confirms that key is present and defined.
func (it *DynamicFunc) HasDefinedAll(names ...string) bool {
	if it == nil || len(names) == 0 {
		return false
	}

	for _, name := range names {
		if it.IsKeyInvalid(name) {
			return false
		}
	}

	// all defined

	return true
}

// IsKeyInvalid
//
// confirms yes if key is missing or null
func (it *DynamicFunc) IsKeyInvalid(name string) bool {
	if it == nil {
		return false
	}

	item, has := it.Params[name]

	return !has ||
		reflectinternal.Is.Null(item)
}

// IsKeyMissing
//
// confirms yes if key is missing  only.
// To check either missing or null use IsKeyInvalid.
func (it *DynamicFunc) IsKeyMissing(name string) bool {
	if it == nil {
		return false
	}

	_, has := it.Params[name]

	return !has
}

func (it DynamicFunc) When() (item interface{}) {
	return it.Params["when"]
}

func (it DynamicFunc) Title() (item interface{}) {
	return it.Params["title"]
}

func (it DynamicFunc) GetLowerCase(name string) (item interface{}, isValid bool) {
	lower := strings.ToLower(name)

	return it.Get(lower)
}

func (it DynamicFunc) GetDirectLower(name string) interface{} {
	x, has := it.Params[strings.ToLower(name)]

	if has {
		return x
	}

	return nil
}

func (it DynamicFunc) Actual() interface{} {
	return it.GetDirectLower("actual")
}

func (it DynamicFunc) Arrange() interface{} {
	return it.GetDirectLower("arrange")
}

func (it *DynamicFunc) Get(name string) (item interface{}, isValid bool) {
	if it == nil {
		return nil, false
	}

	item, has := it.Params[name]

	if has {
		return item, reflectinternal.Is.Defined(item)
	}

	return nil, false
}

func (it *DynamicFunc) GetAsInt(name string) (item int, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return 0, false
	}

	conv, isValid := i.(int)

	return conv, isValid
}

func (it *DynamicFunc) GetAsString(name string) (item string, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return "", false
	}

	conv, isValid := i.(string)

	return conv, isValid
}

func (it *DynamicFunc) GetAsStrings(name string) (items []string, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return []string{}, false
	}

	conv, isValid := i.([]string)

	return conv, isValid
}

func (it *DynamicFunc) GetAsAnyItems(name string) (items []interface{}, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return []interface{}{}, false
	}

	conv, isValid := i.([]interface{})

	return conv, isValid
}

func (it *DynamicFunc) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

func (it *DynamicFunc) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *DynamicFunc) GetFuncName() string {
	return reflectinternal.GetFunc.Name(it.WorkFunc)
}

func (it *DynamicFunc) FuncWrap() *FuncWrap {
	return NewFuncWrap.Default(it.WorkFunc)
}

func (it *DynamicFunc) Invoke(args ...interface{}) (
	results []interface{}, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

func (it *DynamicFunc) InvokeMust(args ...interface{}) (results []interface{}) {
	results, err := it.FuncWrap().Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it *DynamicFunc) InvokeWithValidArgs() (
	results []interface{}, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.ValidArgs()

	return funcWrap.Invoke(validArgs...)
}

func (it *DynamicFunc) InvokeArgs(names ...string) (
	results []interface{}, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.Args(names...)

	return funcWrap.Invoke(validArgs...)
}

func (it *DynamicFunc) ValidArgs() []interface{} {
	return it.Params.ValidArgs()
}

func (it *DynamicFunc) Args(names ...string) []interface{} {
	return it.Params.Args(names...)
}

func (it DynamicFunc) Slice() []interface{} {
	if it.toSlice != nil {
		return *it.toSlice
	}

	var args []interface{}

	keys, err := converters.Map.SortedKeys(it.Params)

	if err != nil {
		panic(err)
	}

	for i, key := range keys {
		value := it.Params[key]
		args = append(
			args, fmt.Sprintf(
				"%d. %s : %s",
				i,
				key,
				value,
			),
		)
	}

	if it.HasFunc() {
		args = append(args, it.GetFuncName())
	}

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it DynamicFunc) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"DynamicFunc",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it DynamicFunc) AsArgsMapper() ArgsMapper {
	return &it
}

func (it DynamicFunc) AsArgFuncNameContractsBinder() ArgFuncNameContractsBinder {
	return &it
}

func (it DynamicFunc) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
