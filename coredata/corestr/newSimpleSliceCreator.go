package corestr

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

type newSimpleSliceCreator struct{}

func (it *newSimpleSliceCreator) Cap(capacity int) *SimpleSlice {
	if capacity <= 0 {
		capacity = 0
	}

	slice := make([]string, 0, capacity)

	return it.Strings(slice)
}

func (it *newSimpleSliceCreator) ByLen(i any) *SimpleSlice {
	length := reflectinternal.SliceConverter.Length(i)

	return it.Cap(length)
}

// Default
//
//	Capacity 10
func (it *newSimpleSliceCreator) Default() *SimpleSlice {
	slice := make([]string, 0, constants.Capacity10)
	toConv := SimpleSlice(slice)

	return &toConv
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
			empty,
		)

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

	isSkipClone := !isClone

	if isSkipClone {
		return it.Strings(lines)
	}

	slice := it.Cap(len(lines))

	return slice.Adds(lines...)
}

// Lines
//
//	don't clone
func (it *newSimpleSliceCreator) Lines(
	lines ...string,
) *SimpleSlice {
	return it.Strings(lines)
}

func (it *newSimpleSliceCreator) Split(
	combined string,
	sep string,
) *SimpleSlice {
	return it.Strings(strings.Split(combined, sep))
}

func (it *newSimpleSliceCreator) SplitLines(
	combined string,
) *SimpleSlice {
	return it.Strings(
		strings.Split(
			combined, constants.NewLineUnix,
		),
	)
}

func (it *newSimpleSliceCreator) SpreadStrings(
	lines ...string,
) *SimpleSlice {
	return it.Strings(lines)
}

func (it *newSimpleSliceCreator) Hashset(
	hashset *Hashset,
) *SimpleSlice {
	if hashset.IsEmpty() {
		return it.Empty()
	}

	return it.Strings(hashset.List())
}

func (it *newSimpleSliceCreator) Map(
	i any,
) *SimpleSlice {
	keys, _ := reflectinternal.
		MapConverter.
		ToKeysStrings(i)

	if len(keys) == 0 {
		return it.Empty()
	}

	return it.Strings(keys)
}

func (it *newSimpleSliceCreator) Create(
	lines []string,
) *SimpleSlice {
	return it.Strings(lines)
}

func (it *newSimpleSliceCreator) Strings(
	lines []string,
) *SimpleSlice {
	slice := SimpleSlice(lines)

	return &slice
}

// Deprecated: Use Strings instead.
func (it *newSimpleSliceCreator) StringsPtr(
	lines []string,
) *SimpleSlice {
	if len(lines) == 0 {
		return it.Empty()
	}

	return it.Strings(lines)
}

// Deprecated: Use Strings or StringsClone instead.
func (it *newSimpleSliceCreator) StringsOptions(
	isClone bool,
	lines []string,
) *SimpleSlice {
	if len(lines) == 0 {
		return it.Empty()
	}

	isSkipClone := !isClone

	if isSkipClone {
		return it.Strings(lines)
	}

	return it.StringsClone(lines)
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

	isSkipClone := !isClone

	if isSkipClone {
		return it.Strings(lines)
	}

	slice := it.Cap(len(lines))

	return slice.Adds(lines...)
}

func (it *newSimpleSliceCreator) UsingSeparatorLine(
	sep, line string,
) *SimpleSlice {
	lines := strings.Split(line, sep)

	return it.Strings(lines)
}

func (it *newSimpleSliceCreator) UsingLine(
	combinedLine string,
) *SimpleSlice {
	lines := strings.Split(combinedLine, constants.DefaultLine)

	return it.Strings(lines)
}

func (it *newSimpleSliceCreator) Empty() *SimpleSlice {
	lines := make([]string, 0)
	slice := SimpleSlice(lines)

	return &slice
}
