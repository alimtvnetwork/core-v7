package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type LeftRight struct {
	Left     interface{} `json:",omitempty"`
	Right    interface{} `json:",omitempty"`
	Expect   interface{} `json:",omitempty"`
	toSlice  *[]interface{}
	toString corestr.SimpleStringOnce
}

func (it *LeftRight) ArgsCount() int {
	return 2
}

func (it *LeftRight) FirstItem() interface{} {
	return it.Left
}

func (it *LeftRight) SecondItem() interface{} {
	return it.Right
}

func (it *LeftRight) Expected() interface{} {
	return it.Expect
}

func (it *LeftRight) ArgTwo() TwoFunc {
	return TwoFunc{
		First:  it.Left,
		Second: it.Right,
	}
}

func (it *LeftRight) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.Left)
}

func (it *LeftRight) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Right)
}

func (it *LeftRight) HasLeft() bool {
	return it != nil && reflectinternal.Is.Defined(it.Left)
}

func (it *LeftRight) HasRight() bool {
	return it != nil && reflectinternal.Is.Defined(it.Right)
}

func (it *LeftRight) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *LeftRight) ValidArgs() []interface{} {
	var args []interface{}

	if it.HasFirst() {
		args = append(args, it.Left)
	}

	if it.HasSecond() {
		args = append(args, it.Right)
	}

	return args
}

func (it *LeftRight) Args(upTo int) []interface{} {
	var args []interface{}

	if upTo >= 1 {
		args = append(args, it.Left)
	}

	if upTo >= 2 {
		args = append(args, it.Right)
	}

	return args
}

func (it *LeftRight) Slice() []interface{} {
	if it.toSlice != nil {
		return *it.toSlice
	}

	var args []interface{}

	if it.HasFirst() {
		args = append(args, it.Left)
	}

	if it.HasSecond() {
		args = append(args, it.Right)
	}

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it *LeftRight) GetByIndex(index int) interface{} {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it *LeftRight) String() string {
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

func (it *LeftRight) Clone() LeftRight {
	return LeftRight{
		Left:   it.Left,
		Right:  it.Right,
		Expect: it.Expect,
	}
}

func (it LeftRight) AsTwoParameter() TwoParameter {
	return &it
}

func (it LeftRight) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
