package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type One struct {
	First    interface{}              `json:",omitempty"`
	Expect   interface{}              `json:",omitempty"`
	toSlice  *[]interface{}           `json:"-"`
	toString corestr.SimpleStringOnce `json:"-"`
}

func (it *One) FirstItem() interface{} {
	return it.First
}

func (it *One) Expected() interface{} {
	return it.Expect
}

func (it *One) ArgTwo() One {
	return One{
		First:  it.First,
		Expect: it.Expect,
	}
}

func (it *One) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *One) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *One) ValidArgs() []interface{} {
	var args []interface{}

	if it.HasFirst() {
		args = append(args, it.First)
	}

	return args
}

func (it *One) Args(upTo int) []interface{} {
	var args []interface{}

	if upTo >= 1 {
		args = append(args, it.First)
	}

	return args
}

func (it *One) ArgsCount() int {
	return 1
}

func (it *One) Slice() []interface{} {
	if it.toSlice != nil {
		return *it.toSlice
	}

	var args []interface{}

	if it.HasFirst() {
		args = append(args, it.First)
	}

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it *One) GetByIndex(index int) interface{} {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it One) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"One",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it *One) LeftRight() LeftRight {
	return LeftRight{
		Left:   it.First,
		Expect: it.Expect,
	}
}

func (it One) AsOneParameter() OneParameter {
	return &it
}

func (it One) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
