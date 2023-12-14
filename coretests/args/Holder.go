package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// Holder
//
// it is used to hold more dynamic parameters for
// the Act function in the unit or integration test
//
// If parameters are not enough use the Hashmap
type Holder struct {
	First    interface{}              `json:",omitempty"`
	Second   interface{}              `json:",omitempty"`
	Third    interface{}              `json:",omitempty"`
	Fourth   interface{}              `json:",omitempty"`
	Fifth    interface{}              `json:",omitempty"`
	Sixth    interface{}              `json:",omitempty"`
	WorkFunc interface{}              `json:"-"`
	Expect   interface{}              `json:",omitempty"`
	Hashmap  Map                      `json:",omitempty"`
	toSlice  *[]interface{}           `json:"-"`
	toString corestr.SimpleStringOnce `json:"-"`
}

func (it *Holder) GetWorkFunc() interface{} {
	return it.WorkFunc
}

func (it *Holder) ArgsCount() int {
	return 7
}

func (it *Holder) FirstItem() interface{} {
	return it.First
}

func (it *Holder) SecondItem() interface{} {
	return it.Second
}

func (it *Holder) ThirdItem() interface{} {
	return it.Third
}

func (it *Holder) FourthItem() interface{} {
	return it.Fourth
}

func (it *Holder) FifthItem() interface{} {
	return it.Fifth
}

func (it *Holder) SixthItem() interface{} {
	return it.Sixth
}

func (it *Holder) Expected() interface{} {
	return it.Expect
}

func (it *Holder) ArgTwo() TwoFunc {
	return TwoFunc{
		First:  it.First,
		Second: it.Second,
	}
}

func (it *Holder) ArgThree() ThreeFunc {
	return ThreeFunc{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it *Holder) ArgFour() FourFunc {
	return FourFunc{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it *Holder) ArgFive() FiveFunc {
	return FiveFunc{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it *Holder) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *Holder) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

func (it *Holder) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

func (it *Holder) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

func (it *Holder) HasFifth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fifth)
}

func (it *Holder) HasSixth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Sixth)
}

func (it *Holder) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

func (it *Holder) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *Holder) GetFuncName() string {
	return reflectinternal.GetFunc.Name(it.WorkFunc)
}

func (it *Holder) FuncWrap() *FuncWrap {
	return NewFuncWrap.Default(it.WorkFunc)
}

func (it *Holder) Invoke(args ...interface{}) (
	results []interface{}, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

func (it *Holder) InvokeMust(args ...interface{}) (results []interface{}) {
	results, err := it.FuncWrap().Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it *Holder) InvokeWithValidArgs() (
	results []interface{}, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.ValidArgs()

	return funcWrap.Invoke(validArgs...)
}

func (it *Holder) InvokeArgs(upTo int) (
	results []interface{}, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.Args(upTo)

	return funcWrap.Invoke(validArgs...)
}

func (it *Holder) ValidArgs() []interface{} {
	var args []interface{}

	if it.HasFirst() {
		args = append(args, it.First)
	}

	if it.HasSecond() {
		args = append(args, it.Second)
	}

	if it.HasThird() {
		args = append(args, it.Third)
	}

	if it.HasFourth() {
		args = append(args, it.Fourth)
	}

	if it.HasFifth() {
		args = append(args, it.Fifth)
	}

	if it.HasSixth() {
		args = append(args, it.Sixth)
	}

	return args
}

func (it *Holder) Args(upTo int) []interface{} {
	var args []interface{}

	if upTo >= 1 {
		args = append(args, it.First)
	}

	if upTo >= 2 {
		args = append(args, it.Second)
	}

	if upTo >= 3 {
		args = append(args, it.Third)
	}

	if upTo >= 4 {
		args = append(args, it.Fourth)
	}

	if upTo >= 5 {
		args = append(args, it.Fifth)
	}

	if upTo >= 6 {
		args = append(args, it.Sixth)
	}

	return args
}

func (it *Holder) Slice() []interface{} {
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

	if it.HasThird() {
		args = append(args, it.Third)
	}

	if it.HasFourth() {
		args = append(args, it.Fourth)
	}

	if it.HasFifth() {
		args = append(args, it.Fifth)
	}

	if it.HasSixth() {
		args = append(args, it.Sixth)
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

func (it *Holder) GetByIndex(index int) interface{} {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it *Holder) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"Holder",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it Holder) AsSixthParameter() coreinterface.SixthParameter {
	return &it
}

func (it Holder) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}
