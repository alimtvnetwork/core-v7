package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type TwoFunc struct {
	First    interface{}              `json:",omitempty"`
	Second   interface{}              `json:",omitempty"`
	WorkFunc interface{}              `json:"-"`
	Expect   interface{}              `json:",omitempty"`
	toSlice  *[]interface{}           `json:"-"`
	toString corestr.SimpleStringOnce `json:"-"`
}

func (it *TwoFunc) GetWorkFunc() interface{} {
	return it.WorkFunc
}

func (it *TwoFunc) ArgsCount() int {
	return 2
}

func (it *TwoFunc) FirstItem() interface{} {
	return it.First
}

func (it *TwoFunc) SecondItem() interface{} {
	return it.Second
}

func (it *TwoFunc) Expected() interface{} {
	return it.Expect
}

func (it *TwoFunc) ArgTwo() TwoFunc {
	return TwoFunc{
		First:  it.First,
		Second: it.Second,
	}
}

func (it *TwoFunc) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *TwoFunc) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

func (it *TwoFunc) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

func (it *TwoFunc) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *TwoFunc) GetFuncName() string {
	return reflectinternal.GetFunc.Name(it.WorkFunc)
}

func (it *TwoFunc) FuncWrap() *FuncWrap {
	return NewFuncWrap.Default(it.WorkFunc)
}

func (it *TwoFunc) Invoke(args ...interface{}) (
	results []interface{}, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

func (it *TwoFunc) InvokeMust(args ...interface{}) (results []interface{}) {
	results, err := it.FuncWrap().Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it *TwoFunc) InvokeWithValidArgs() (
	results []interface{}, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.ValidArgs()

	return funcWrap.Invoke(validArgs...)
}

func (it *TwoFunc) InvokeArgs(upTo int) (
	results []interface{}, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.Args(upTo)

	return funcWrap.Invoke(validArgs...)
}

func (it *TwoFunc) ValidArgs() []interface{} {
	var args []interface{}

	if it.HasFirst() {
		args = append(args, it.First)
	}

	if it.HasSecond() {
		args = append(args, it.Second)
	}

	return args
}

func (it *TwoFunc) Args(upTo int) []interface{} {
	var args []interface{}

	if upTo >= 1 {
		args = append(args, it.First)
	}

	if upTo >= 2 {
		args = append(args, it.Second)
	}

	return args
}

func (it *TwoFunc) Slice() []interface{} {
	if it.toSlice != nil {
		return *it.toSlice
	}

	var args []interface{}

	if it.HasFirst() {
		args = append(args, it.First)
	}

	if it.HasSecond() {
		args = append(args, it.Second)
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

func (it *TwoFunc) GetByIndex(index int) interface{} {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it TwoFunc) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"TwoFunc",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it *TwoFunc) LeftRight() LeftRight {
	return LeftRight{
		Left:   it.First,
		Right:  it.Second,
		Expect: it.Expect,
	}
}

func (it TwoFunc) AsTwoFuncParameter() TwoFuncParameter {
	return &it
}

func (it TwoFunc) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}

func (it TwoFunc) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
