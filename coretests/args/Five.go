package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Five struct {
	First    interface{} `json:",omitempty"`
	Second   interface{} `json:",omitempty"`
	Third    interface{} `json:",omitempty"`
	Fourth   interface{} `json:",omitempty"`
	Fifth    interface{} `json:",omitempty"`
	Expect   interface{} `json:",omitempty"`
	toSlice  *[]interface{}
	toString corestr.SimpleStringOnce
}

func (it *Five) ArgsCount() int {
	return 5
}

func (it *Five) FirstItem() interface{} {
	return it.First
}

func (it *Five) SecondItem() interface{} {
	return it.Second
}

func (it *Five) ThirdItem() interface{} {
	return it.Third
}

func (it *Five) FourthItem() interface{} {
	return it.Fourth
}

func (it *Five) FifthItem() interface{} {
	return it.Fifth
}

func (it *Five) Expected() interface{} {
	return it.Expect
}

func (it *Five) ArgTwo() Two {
	return Two{
		First:  it.First,
		Second: it.Second,
	}
}

func (it *Five) ArgThree() Three {
	return Three{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it *Five) ArgFour() Four {
	return Four{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it *Five) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *Five) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

func (it *Five) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

func (it *Five) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

func (it *Five) HasFifth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fifth)
}

func (it *Five) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *Five) ValidArgs() []interface{} {
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

	return args
}

func (it *Five) Args(upTo int) []interface{} {
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

	return args
}

func (it *Five) Slice() []interface{} {
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

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it *Five) GetByIndex(index int) interface{} {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it *Five) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"Five",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it Five) AsFifthParameter() FifthParameter {
	return &it
}

func (it Five) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
