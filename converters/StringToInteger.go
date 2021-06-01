package converters

import (
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

func StringToInteger(
	input string,
) (value int, err error) {
	value, err2 := strconv.Atoi(input)

	if err2 != nil {
		reference := input +
			constants.NewLineUnix +
			err2.Error()

		return constants.Zero, msgtype.ParsingFailed.Error(
			msgtype.FailedToConvert.String(),
			reference)
	}

	return value, err
}
