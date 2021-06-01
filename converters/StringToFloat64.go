package converters

import (
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/bitsize"
	"gitlab.com/evatix-go/core/msgtype"
)

func StringToFloat64(input string) (value float64, err error) {
	value, err2 := strconv.ParseFloat(input, bitsize.Of64)

	if err2 != nil {
		reference := input +
			constants.NewLineUnix +
			err2.Error()

		return constants.Zero, msgtype.
			ParsingFailed.Error(
			msgtype.FailedToConvert.String(),
			reference)
	}

	return value, err
}
