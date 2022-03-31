package corestr

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type newSimpleSliceCreator struct{}

func (it *newSimpleSliceCreator) Cap(capacity int) *SimpleSlice {
	slice := make([]string, 0, capacity)

	return &SimpleSlice{
		slice,
	}
}

// Default
//
//  Capacity 10
func (it *newSimpleSliceCreator) Default() *SimpleSlice {
	slice := make([]string, 0, constants.Capacity10)

	return &SimpleSlice{
		slice,
	}
}

func (it *newSimpleSliceCreator) DefaultSlice() SimpleSlice {
	slice := make([]string, 0, constants.Capacity5)

	return SimpleSlice{
		slice,
	}
}

func (it *newSimpleSliceCreator) Deserialize(
	jsonBytes []byte,
) (*SimpleSlice, error) {
	lines, err := corejson.Deserialize.BytesTo.Strings(jsonBytes)

	if err == nil {
		return it.Strings(lines), nil
	}

	return it.Empty(), err
}

func (it *newSimpleSliceCreator) DeserializeJsoner(
	jsoner corejson.Jsoner,
) (*SimpleSlice, error) {
	empty := it.Empty()

	err := corejson.
		Deserialize.
		UsingJsonerToAny(
			true,
			jsoner,
			empty)

	if err == nil {
		return empty, nil
	}

	return empty, err
}

func (it *newSimpleSliceCreator) UsingLines(
	isClone bool,
	lines ...string,
) *SimpleSlice {
	if lines == nil {
		return it.Empty()
	}

	if !isClone {
		return &SimpleSlice{
			lines,
		}
	}

	slice := it.Cap(len(lines))

	return slice.Adds(lines...)
}

// Lines
//
//  don't clone
func (it *newSimpleSliceCreator) Lines(
	lines ...string,
) *SimpleSlice {
	return &SimpleSlice{
		lines,
	}
}

func (it *newSimpleSliceCreator) SpreadStrings(
	lines ...string,
) *SimpleSlice {
	return &SimpleSlice{
		lines,
	}
}

func (it *newSimpleSliceCreator) Create(
	lines []string,
) *SimpleSlice {
	return &SimpleSlice{
		lines,
	}
}

func (it *newSimpleSliceCreator) Strings(
	lines []string,
) *SimpleSlice {
	return &SimpleSlice{
		lines,
	}
}

func (it *newSimpleSliceCreator) StringsPtr(
	lines *[]string,
) *SimpleSlice {
	if lines == nil || len(*lines) == 0 {
		return it.Empty()
	}

	return &SimpleSlice{
		*lines,
	}
}

func (it *newSimpleSliceCreator) StringsPtrOption(
	isClone bool,
	lines *[]string,
) *SimpleSlice {
	if lines == nil || len(*lines) == 0 {
		return it.Empty()
	}

	if !isClone {
		return &SimpleSlice{
			*lines,
		}
	}

	return it.StringsClone(*lines)
}

func (it *newSimpleSliceCreator) StringsClone(
	lines []string,
) *SimpleSlice {
	if lines == nil {
		return it.Empty()
	}

	slice := it.Cap(len(lines))

	return slice.Adds(lines...)
}

func (it *newSimpleSliceCreator) Direct(
	isClone bool,
	lines []string,
) *SimpleSlice {
	if lines == nil {
		return it.Empty()
	}

	if !isClone {
		return &SimpleSlice{
			lines,
		}
	}

	slice := it.Cap(len(lines))

	return slice.Adds(lines...)
}

func (it *newSimpleSliceCreator) UsingSeparatorLine(
	sep, line string,
) *SimpleSlice {
	lines := strings.Split(line, sep)

	return &SimpleSlice{
		Items: lines,
	}
}

func (it *newSimpleSliceCreator) UsingLine(
	combinedLine string,
) *SimpleSlice {
	lines := strings.Split(combinedLine, constants.DefaultLine)

	return &SimpleSlice{
		Items: lines,
	}
}

func (it *newSimpleSliceCreator) Empty() *SimpleSlice {
	return &SimpleSlice{
		nil,
	}
}
