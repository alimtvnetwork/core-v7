package converters

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/constants/bitsize"
	"gitlab.com/auk-go/core/converters/coreconverted"
	"gitlab.com/auk-go/core/errcore"
)

type stringTo struct{}

func (it stringTo) IntegerWithDefault(
	input string,
	defaultInt int,
) (value int, isSuccess bool) {
	if input == constants.EmptyString {
		return defaultInt, false
	}

	convertedVal, err := strconv.Atoi(input)

	if err != nil {
		return defaultInt, false
	}

	return convertedVal, true
}

func (it stringTo) IntegersWithDefaults(
	stringInput,
	separator string,
	defaultInt int,
) *coreconverted.Integers {
	if stringInput == "" {
		return &coreconverted.Integers{
			Values:        []int{},
			CombinedError: nil,
		}
	}

	splits := strings.Split(stringInput, separator)
	results := make([]int, len(splits))
	var errMessages []string

	for i, v := range splits {
		vInt, err := strconv.Atoi(v)

		if err != nil {
			results[i] = defaultInt
			errMessage := constants.IndexColonSpace +
				strconv.Itoa(i) +
				err.Error()
			errMessages = append(
				errMessages,
				errMessage,
			)

			continue
		}

		results[i] = vInt
	}

	var combinedError error
	if len(errMessages) > 0 {
		errCompiledMessage := strings.Join(errMessages, constants.NewLineUnix)
		combinedError = errors.New(errCompiledMessage)
	}

	return &coreconverted.Integers{
		Values:        results,
		CombinedError: combinedError,
	}
}

func (it stringTo) IntegersConditional(
	stringInput,
	separator string,
	processor func(in string) (out int, isTake, isBreak bool),
) *[]int {
	if stringInput == "" {
		return &[]int{}
	}

	splits := strings.Split(stringInput, separator)
	results := make([]int, 0, len(splits))

	for _, v := range splits {
		out, isTake, isBreak := processor(v)

		if isTake {
			results = append(results, out)
		}

		if isBreak {
			break
		}
	}

	return &results
}

func (it stringTo) IntegerMust(
	input string,
) (value int) {
	value, err := it.Integer(input)

	if err != nil {
		panic(err)
	}

	return value
}

func (it stringTo) IntegerDefault(
	input string,
) int {
	value, err2 := strconv.Atoi(input)

	if err2 != nil {
		return constants.Zero
	}

	return value
}

func (it stringTo) Integer(
	input string,
) (value int, err error) {
	value, err2 := strconv.Atoi(input)

	if err2 != nil {
		reference := input +
			constants.NewLineUnix +
			err2.Error()

		return constants.Zero, errcore.ParsingFailedType.Error(
			errcore.FailedToConvertType.String(),
			reference,
		)
	}

	return value, err
}

func (it stringTo) Float64Must(input string) float64 {
	value, err2 := it.Float64(input)

	if err2 != nil {
		panic(err2)
	}

	return value
}

func (it stringTo) Float64Default(
	input string, defaultFloat64 float64,
) (value float64, isSuccess bool) {
	value, err2 := strconv.ParseFloat(input, bitsize.Of64)

	if err2 != nil {
		return defaultFloat64, false
	}

	return value, true
}

func (it stringTo) Float64Conditional(
	input string, defaultFloat64 float64,
) (value float64, isSuccess bool) {
	value, err2 := strconv.ParseFloat(input, bitsize.Of64)

	if err2 != nil {
		return defaultFloat64, false
	}

	return value, true
}

func (it stringTo) Float64(input string) (value float64, err error) {
	value, err2 := strconv.ParseFloat(input, bitsize.Of64)

	if err2 != nil {
		reference := input +
			constants.NewLineUnix +
			err2.Error()

		return constants.Zero, errcore.
			ParsingFailedType.Error(
			errcore.FailedToConvertType.String(),
			reference,
		)
	}

	return value, err
}

func (it stringTo) ByteWithDefault(
	input string, defaultByte byte,
) (value byte, isSuccess bool) {
	vByte, err := it.Byte(input)

	if err != nil {
		return defaultByte, false
	}

	return vByte, true
}

func (it stringTo) BytesConditional(
	stringInput,
	separator string,
	processor func(in string) (out byte, isTake, isBreak bool),
) *[]byte {
	if stringInput == "" {
		return &[]byte{}
	}

	splits := strings.Split(stringInput, separator)
	results := make([]byte, 0, len(splits))

	for _, v := range splits {
		out, isTake, isBreak := processor(v)

		if isTake {
			results = append(results, out)
		}

		if isBreak {
			break
		}
	}

	return &results
}

func (it stringTo) Byte(input string) (byte, error) {
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

// JsonBytes
//
// Wraps with double quote and then returns as bytes
func (it stringTo) JsonBytes(name string) []byte {
	doubleQuoted := fmt.Sprintf(
		constants.SprintDoubleQuoteFormat,
		name,
	)

	return []byte(doubleQuoted)
}
