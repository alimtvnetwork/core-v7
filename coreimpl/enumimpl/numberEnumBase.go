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
	stringRanges         []string
	rangesCsvString      *coreonce.StringOnce
	rangesInvalidMessage *coreonce.StringOnce
	invalidError         *coreonce.ErrorOnce
	typeName             string
}

func newNumberEnumBase(
	typeName string,
	actualRangesAnyType interface{},
	stringRanges []string,
	min, max interface{},
) *numberEnumBase {
	if stringRanges == nil {
		msgtype.MeaningfulErrorHandle(
			msgtype.CannotBeNilMessage,
			"newNumberEnumBase",
			errors.New("StringRanges cannot be nil"))
	}

	rangesToCsvOnce := coreonce.NewStringOncePtr(func() string {
		return converters.StringsToCsvWithIndexes(
			stringRanges,
		)
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
		typeName:        typeName,
	}
}

func (it *numberEnumBase) TypeName() string {
	return it.typeName
}

func (it *numberEnumBase) RangeNamesCsv() string {
	return it.rangesCsvString.Value()
}

func (it *numberEnumBase) RangesInvalidMessage() string {
	return it.rangesInvalidMessage.Value()
}

func (it *numberEnumBase) RangesInvalidErr() error {
	return it.invalidError.Value()
}

func (it *numberEnumBase) StringRangesPtr() *[]string {
	return &it.stringRanges
}

func (it *numberEnumBase) StringRanges() []string {
	return it.stringRanges
}

func (it *numberEnumBase) StringJson(input interface{}) (jsonString string, err error) {
	return it.ToEnumString(input), nil
}

func (it *numberEnumBase) StringJsonMust(input interface{}) string {
	return it.ToEnumString(input)
}

func (it *numberEnumBase) ToEnumString(input interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, input)
}
