package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Six struct {
	First    interface{} `json:",omitempty"`
	Second   interface{} `json:",omitempty"`
	Third    interface{} `json:",omitempty"`
	Fourth   interface{} `json:",omitempty"`
	Fifth    interface{} `json:",omitempty"`
	Sixth    interface{} `json:",omitempty"`
	Expect   interface{} `json:",omitempty"`
	toSlice  *[]interface{}
	toString corestr.SimpleStringOnce
}

func (it *Six) ArgsCount() int {
	return 6
}

func (it *Six) FirstItem() interface{} {
	return it.First
}

func (it *Six) SecondItem() interface{} {
	return it.Second
}

func (it *Six) ThirdItem() interface{} {
	return it.Third
}

func (it *Six) FourthItem() interface{} {
	return it.Fourth
}

func (it *Six) FifthItem() interface{} {
	return it.Fifth
}

func (it *Six) SixthItem() interface{} {
	return it.Sixth
}

func (it *Six) Expected() interface{} {
	return it.Expect
}

func (it Six) ArgTwo() Two {
	return Two{
		First:  it.First,
		Second: it.Second,
	}
}

func (it Six) ArgThree() Three {
	return Three{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it Six) ArgFour() Four {
	return Four{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it Six) ArgFive() Five {
	return Five{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it *Six) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *Six) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

func (it *Six) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

func (it *Six) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

func (it *Six) HasFifth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fifth)
}

func (it *Six) HasSixth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Sixth)
}

func (it *Six) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it Six) Slice() []interface{} {
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

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it Six) GetByIndex(index int) interface{} {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it *Six) ValidArgs() []interface{} {
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

func (it *Six) Args(upTo int) []interface{} {
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

func (it Six) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"Six",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it Six) AsSixthParameter() SixthParameter {
	return &it
}

func (it Six) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
