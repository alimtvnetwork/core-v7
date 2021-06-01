package enumimpl

import (
	"errors"
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/coreonce"
	"gitlab.com/evatix-go/core/msgtype"
)

type numberEnumBase struct {
	actualValueRanges    interface{}
	stringRanges         *[]string
	rangesCsvString      *coreonce.StringOnce
	rangesInvalidMessage *coreonce.StringOnce
	invalidError         *coreonce.ErrorOnce
}

func newNumberEnumBase(
	actualRangesAnyType interface{},
	stringRanges *[]string,
	min, max interface{},
) *numberEnumBase {
	rangesToCsvOnce := coreonce.NewStringOncePtr(func() string {
		return converters.StringsToCsv(
			stringRanges,
			false)
	})

	invalidMessageOnce := coreonce.NewStringOncePtr(func() string {
		msg := msgtype.EnumRangeNotMeet(
			min,
			max,
			rangesToCsvOnce.Value())

		return msg
	})

	return &numberEnumBase{
		actualValueRanges:    actualRangesAnyType,
		stringRanges:         stringRanges,
		rangesInvalidMessage: invalidMessageOnce,
		invalidError: coreonce.NewErrorOncePtr(func() error {
			return errors.New(invalidMessageOnce.Value())
		}),
		rangesCsvString: rangesToCsvOnce,
	}
}

func (receiver *numberEnumBase) RangeNamesCsv() string {
	return receiver.rangesCsvString.Value()
}

func (receiver *numberEnumBase) RangesInvalidMessage() string {
	return receiver.rangesInvalidMessage.Value()
}

func (receiver *numberEnumBase) RangesInvalidErr() error {
	return receiver.invalidError.Value()
}

func (receiver *numberEnumBase) StringRangesPtr() *[]string {
	return receiver.stringRanges
}

func (receiver *numberEnumBase) StringRanges() []string {
	return *receiver.stringRanges
}

func (receiver *numberEnumBase) StringJson(input interface{}) (jsonString string, err error) {
	return receiver.ToEnumString(input), nil
}

func (receiver *numberEnumBase) StringJsonMust(input interface{}) string {
	return receiver.ToEnumString(input)
}

func (receiver *numberEnumBase) ToEnumString(input interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, input)
}
