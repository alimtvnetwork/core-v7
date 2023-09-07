package converters

import (
	"strconv"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

func StringToByte(input string) (byte, error) {
	if input == "" {
		return 0, errcore.FailedToConvertType.
			Error(errcore.CannotConvertStringToByte, input)
	}

	if input == "0" {
		return 0, nil
	}

	if input == "1" {
		return 1, nil
	}

	vInt, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}

	if vInt < 0 {
		return 0, errcore.FailedToConvertType.
			Error(errcore.CannotConvertStringToByteForLessThanZero, input)
	}

	if vInt > constants.MaxUnit8AsInt {
		return 0, errcore.FailedToConvertType.
			Error(errcore.CannotConvertStringToByteForMoreThan255, input)
	}

	return byte(vInt), nil
}
