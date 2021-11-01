package corerange

import (
	"math"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/bitsize"
)

type within struct{}

func (it *within) StringRangeInt64(
	input string,
) (val int64, isInRange bool) {
	finalInt, isInRange := it.StringRangeInteger(
		true,
		math.MinInt64,
		math.MaxInt64,
		input)

	return int64(finalInt), isInRange
}

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
		math.MaxUint32,
		input)

	return uint32(finalInt), isInRange
}

func (it *within) StringRangeUint64(
	input string,
) (val uint64, isInRange bool) {
	finalInt, isInRange := it.StringRangeInteger(
		true,
		constants.Zero,
		math.MaxUint64,
		input)

	return uint64(finalInt), isInRange
}

func (it *within) StringRangeIntegerDefault(
	min, max int,
	input string,
) (val int, isInRange bool) {
	toInt, err := strconv.Atoi(input)

	if err != nil {
		return constants.Zero, false
	}

	return it.RangeInteger(
		true,
		min,
		max,
		toInt)

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

	if input > max {
		return max, false
	}

	return constants.Zero, false
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

	if input > max {
		return max, false
	}

	return constants.Zero, false
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

	if input > max {
		return max, false
	}

	return constants.Zero, false
}
