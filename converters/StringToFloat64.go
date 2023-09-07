package converters

import (
	"strconv"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/constants/bitsize"
	"gitlab.com/auk-go/core/errcore"
)

func StringToFloat64(input string) (value float64, err error) {
	value, err2 := strconv.ParseFloat(input, bitsize.Of64)

	if err2 != nil {
		reference := input +
			constants.NewLineUnix +
			err2.Error()

		return constants.Zero, errcore.
			ParsingFailedType.Error(
			errcore.FailedToConvertType.String(),
			reference)
	}

	return value, err
}
