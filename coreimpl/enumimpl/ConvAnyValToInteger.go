package enumimpl

import (
	"fmt"
	"strconv"

	"gitlab.com/auk-go/core/constants"
)

func ConvEnumAnyValToInteger(val interface{}) int {
	_, isStr := val.(string)

	if isStr {
		// already a string represents string type enum
		return constants.MinInt
	}

	valInt, isInt := val.(int)

	if isInt {
		return valInt
	}

	switch casted := val.(type) {
	case valueByter:
		return int(casted.Value())
	case exactValueByter:
		return int(casted.ValueByte())
	case valueInter:
		return casted.Value()
	case exactValueInter:
		return casted.ValueInt()
	case valueInt8er:
		return int(casted.Value())
	case exactValueInt8er:
		return int(casted.ValueInt8())
	case valueUInt16er:
		return int(casted.Value())
	case exactValueUInt16er:
		return int(casted.ValueUInt16())
	}

	str := fmt.Sprintf(
		constants.SprintValueFormat,
		val)

	convValueInt, err := strconv.Atoi(str)

	if err != nil {
		return constants.MinInt
	}

	return convValueInt
}
