package corerange

import (
	"math"
	"strconv"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/constants/bitsize"
)

type within struct{}

func (it *within) StringRangeInt32(
	input string,
) (val int32, isInRange bool) {
	finalInt, isInRange := it.StringRangeInteger(
		true,
		math.MinInt32,
		math.MaxInt32,
		input)

	return int32(finalInt), isInRange
}

func (it *within) StringRangeInt16(
	input string,
) (val int16, isInRange bool) {
	finalInt, isInRange := it.StringRangeInteger(
		true,
		math.MinInt16,
		math.MaxInt16,
		input)

	return int16(finalInt), isInRange
}

func (it *within) StringRangeInt8(
	input string,
) (val int8, isInRange bool) {
	finalInt, isInRange := it.StringRangeInteger(
		true,
		math.MinInt8,
		math.MaxInt8,
		input)

	return int8(finalInt), isInRange
}

func (it *within) StringRangeByte(
	input string,
) (val byte, isInRange bool) {
	finalInt, isInRange := it.StringRangeInteger(
		true,
		constants.Zero,
		math.MaxUint8,
		input)

	return byte(finalInt), isInRange
}

func (it *within) StringRangeUint16(
	input string,
) (val uint16, isInRange bool) {
	finalInt, isInRange := it.StringRangeInteger(
		true,
		constants.Zero,
		math.MaxUint16,
		input)

	return uint16(finalInt), isInRange
}

func (it *within) StringRangeUint32(
	input string,
) (val uint32, isInRange bool) {
	finalInt, isInRange := it.StringRangeInteger(
		true,
		constants.Zero,
		math.MaxInt32,
		input)

	// fix https://t.ly/6aoW,
	// https://gitlab.com/auk-go/core/-/issues/81
	// use MaxInt32 instead of uint32Max
	if finalInt <= math.MaxInt32 {
		return uint32(finalInt), isInRange
	}

	return 0, isInRange
}

func (it *within) StringRangeIntegerDefault(
	min, max int,
	input string,
) (val int, isInRange bool) {
	toInt, err := strconv.ParseInt(
		input,
		10,
		64)

	if err != nil {
		return constants.Zero, false
	}

	isInRange = toInt >= int64(min) &&
		toInt <= int64(max)

	if isInRange {
		return int(toInt), isInRange
	}

	isLessMin := toInt < int64(min)
	if isLessMin {
		return min, false
	}

	// above
	return max, false
}

func (it *within) StringRangeInteger(
	isUsageMinMaxBoundary bool,
	min, max int,
	input string,
) (val int, isInRange bool) {
	toInt, err := strconv.Atoi(input)

	if err != nil {
		return constants.Zero, false
	}

	return it.RangeInteger(
		isUsageMinMaxBoundary,
		min,
		max,
		toInt)
}

func (it *within) StringRangeFloat(
	isUsageMinMaxBoundary bool,
	min, max float32,
	input string,
) (val float32, isInRange bool) {
	toFloat64, err := strconv.ParseFloat(input, bitsize.Of32)

	if err != nil {
		return constants.Zero, false
	}

	rangedValue, isInRange := it.RangeFloat64(
		isUsageMinMaxBoundary,
		float64(min),
		float64(max),
		toFloat64)

	if isInRange || isUsageMinMaxBoundary {
		return float32(rangedValue), isInRange
	}

	return constants.Zero, isInRange
}

func (it *within) StringRangeFloatDefault(
	input string,
) (val float32, isInRange bool) {
	toFloat64, err := strconv.ParseFloat(input, bitsize.Of32)

	if err != nil {
		return constants.Zero, false
	}

	rangedValue, isInRange := it.RangeFloat64(
		true,
		math.SmallestNonzeroFloat32,
		math.MaxFloat32,
		toFloat64)

	return float32(rangedValue), isInRange
}

func (it *within) StringRangeFloat64(
	isUsageMinMaxBoundary bool,
	min, max float64,
	input string,
) (val float64, isInRange bool) {
	toFloat, err := strconv.ParseFloat(input, bitsize.Of64)

	if err != nil {
		return constants.Zero, false
	}

	return it.RangeFloat64(
		isUsageMinMaxBoundary,
		min,
		max,
		toFloat)
}

func (it *within) StringRangeFloat64Default(
	input string,
) (val float64, isInRange bool) {
	toFloat, err := strconv.ParseFloat(input, bitsize.Of64)

	if err != nil {
		return constants.Zero, false
	}

	return it.RangeFloat64(
		true,
		math.SmallestNonzeroFloat32,
		math.MaxFloat32,
		toFloat)
}

func (it *within) RangeDefaultInteger(
	min, max, input int,
) (val int, isInRange bool) {
	return it.RangeInteger(
		true,
		min,
		max,
		input)
}

func (it *within) RangeInteger(
	isUsageMinMaxBoundary bool,
	min, max,
	input int,
) (val int, isInRange bool) {
	if input >= min && input <= max {
		return input, true
	}

	if !isUsageMinMaxBoundary {
		return input, false
	}

	if input < min {
		return min, false
	}

	return max, false
}

func (it *within) RangeByteDefault(
	input int,
) (val byte, isInRange bool) {
	return it.RangeByte(
		true,
		input)
}

func (it *within) RangeByte(
	isUsageMinMaxBoundary bool,
	input int,
) (val byte, isInRange bool) {
	if input >= constants.Zero && input <= math.MaxUint8 {
		return byte(input), true
	}

	if !isUsageMinMaxBoundary {
		return constants.Zero, false
	}

	if input < constants.Zero {
		return constants.Zero, false
	}

	return math.MaxUint8, false
}

func (it *within) RangeUint16Default(
	input int,
) (val uint16, isInRange bool) {
	return it.RangeUint16(
		true,
		input)
}

func (it *within) RangeUint16(
	isUsageMinMaxBoundary bool,
	input int,
) (val uint16, isInRange bool) {
	toInt, isInRange := it.RangeInteger(
		isUsageMinMaxBoundary,
		constants.Zero,
		math.MaxUint16,
		input)

	if isInRange || isUsageMinMaxBoundary {
		return uint16(toInt), isInRange
	}

	return constants.Zero, isInRange
}

func (it *within) RangeFloat(
	isUsageMinMaxBoundary bool,
	min, max,
	input float32,
) (val float32, isInRange bool) {
	if input >= min && input <= max {
		return input, true
	}

	if !isUsageMinMaxBoundary {
		return input, false
	}

	if input < min {
		return min, false
	}

	return max, false
}

func (it *within) RangeFloat64(
	isUsageMinMaxBoundary bool,
	min, max,
	input float64,
) (val float64, isInRange bool) {
	if input >= min && input <= max {
		return input, true
	}

	if !isUsageMinMaxBoundary {
		return input, false
	}

	if input < min {
		return min, false
	}

	return max, false
}
