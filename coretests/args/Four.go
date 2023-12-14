package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Four struct {
	First    interface{} `json:",omitempty"`
	Second   interface{} `json:",omitempty"`
	Third    interface{} `json:",omitempty"`
	Fourth   interface{} `json:",omitempty"`
	Expect   interface{} `json:",omitempty"`
	toSlice  *[]interface{}
	toString corestr.SimpleStringOnce
}

func (it *Four) ArgsCount() int {
	return 4
}

func (it *Four) FirstItem() interface{} {
	return it.First
}

func (it *Four) SecondItem() interface{} {
	return it.Second
}

func (it *Four) ThirdItem() interface{} {
	return it.Third
}

func (it *Four) FourthItem() interface{} {
	return it.Fourth
}

func (it *Four) Expected() interface{} {
	return it.Expect
}

func (it *Four) ArgTwo() Two {
	return Two{
		First:  it.First,
		Second: it.Second,
	}
}

func (it *Four) ArgThree() Three {
	return Three{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it *Four) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *Four) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

func (it *Four) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

func (it *Four) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

func (it *Four) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *Four) ValidArgs() []interface{} {
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

	return args
}

func (it *Four) Args(upTo int) []interface{} {
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

	return args
}

func (it *Four) Slice() []interface{} {
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

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it *Four) GetByIndex(index int) interface{} {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it *Four) String() string {
	var args []string

	if it.HasFirst() {
		args = append(args, toString(it.First))
	}

	if it.HasSecond() {
		args = append(args, toString(it.Second))
	}

	if it.HasThird() {
		args = append(args, toString(it.Third))
	}

	if it.HasFourth() {
		args = append(args, toString(it.Fourth))
	}

	if it.HasExpect() {
		args = append(args, toString(it.Expect))
	}

	return fmt.Sprintf(
		selfToStringFmt,
		"Four",
		strings.Join(args, constants.CommaSpace),
	)
}

func (it Four) AsFourParameter() FourParameter {
	return &it
}

func (it Four) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
