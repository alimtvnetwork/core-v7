package corestr

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
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
