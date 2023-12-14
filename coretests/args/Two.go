package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Two struct {
	First    interface{}              `json:",omitempty"`
	Second   interface{}              `json:",omitempty"`
	Expect   interface{}              `json:",omitempty"`
	toSlice  *[]interface{}           `json:"-"`
	toString corestr.SimpleStringOnce `json:"-"`
}

func (it *Two) FirstItem() interface{} {
	return it.First
}

func (it *Two) SecondItem() interface{} {
	return it.Second
}

func (it *Two) Expected() interface{} {
	return it.Expect
}

func (it *Two) ArgTwo() TwoFunc {
	return TwoFunc{
		First:  it.First,
		Second: it.Second,
	}
}

func (it *Two) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *Two) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

func (it *Two) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *Two) ValidArgs() []interface{} {
	var args []interface{}

	if it.HasFirst() {
		args = append(args, it.First)
	}

	if it.HasSecond() {
		args = append(args, it.Second)
	}

	return args
}

func (it *Two) ArgsCount() int {
	return 1
}

func (it *Two) Args(upTo int) []interface{} {
	var args []interface{}

	if upTo >= 1 {
		args = append(args, it.First)
	}

	if upTo >= 2 {
		args = append(args, it.Second)
	}

	return args
}

func (it *Two) Slice() []interface{} {
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

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it *Two) GetByIndex(index int) interface{} {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it *Two) String() string {
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

func (it *Two) LeftRight() LeftRight {
	return LeftRight{
		Left:   it.First,
		Right:  it.Second,
		Expect: it.Expect,
	}
}

func (it Two) AsTwoParameter() TwoParameter {
	return &it
}

func (it Two) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
