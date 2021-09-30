package corecomparator

import (
	"encoding/json"
	"strconv"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/msgtype"
)

type Compare byte

const (
	Equal Compare = iota
	LeftGreater
	LeftGreaterEqual
	LeftLess
	LeftLessEqual
	NotEqual
	Inconclusive
)

func (it Compare) Is(other Compare) bool {
	return it == other
}

func (it Compare) IsEqual() bool {
	return it == Equal
}

func (it Compare) IsLeftGreater() bool {
	return it == LeftGreater
}

func (it Compare) IsLeftGreaterEqual() bool {
	return it == LeftGreaterEqual
}

func (it Compare) IsLeftLess() bool {
	return it == LeftLess
}

func (it Compare) IsLeftLessEqual() bool {
	return it == LeftLessEqual
}

func (it Compare) IsLeftLessOrLessEqualOrEqual() bool {
	return it == Equal || it == LeftLess || it == LeftLessEqual
}

func (it Compare) IsLeftGreaterOrGreaterEqualOrEqual() bool {
	return it == Equal || it == LeftGreater || it == LeftGreaterEqual
}

func (it Compare) IsNotEqual() bool {
	return it == NotEqual
}

func (it Compare) IsNotEqualLogically() bool {
	return it != Equal
}

func (it Compare) IsInconclusive() bool {
	return it == Inconclusive
}

func (it Compare) IsNotInconclusive() bool {
	return it != Inconclusive
}

func (it Compare) IsDefinedProperly() bool {
	return it != Inconclusive
}

func (it Compare) IsInconclusiveOrNotEqual() bool {
	return it == Inconclusive || it == NotEqual
}

func (it Compare) IsAnyOf(values ...Compare) bool {
	if len(values) == 0 {
		return true
	}

	for _, value := range values {
		if it == value {
			return true
		}
	}

	return false
}

func (it Compare) Name() string {
	return it.String()
}

func (it Compare) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.String())
}

func (it *Compare) UnmarshalJSON(data []byte) error {
	if data == nil {
		return defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	name := string(data)

	compare, has := RangesMap[name]

	if has {
		*it = compare

		return nil
	}

	return msgtype.
		FailedToConvert.
		Error(string(data)+" failed to convert to core-compare. Must be any of the values.",
			strings.Join(Ranges(),
				constants.Comma))
}

func (it Compare) Value() byte {
	return byte(it)
}

func (it Compare) OperatorSymbol() string {
	return CompareOperatorsSymbols[it]
}

func (it Compare) OperatorShortForm() string {
	return CompareOperatorsShotNames[it]
}

func (it Compare) SqlOperatorSymbol() string {
	return SqlCompareOperators[it]
}

func (it Compare) NumberString() string {
	return strconv.Itoa(int(it))
}

func (it Compare) NumberJsonString() string {
	return "\"" + strconv.Itoa(int(it)) + "\""
}

func (it Compare) StringValue() string {
	return string(it)
}

func (it Compare) String() string {
	return CompareNames[it]
}
