package defaultcapacity

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/constants/percentconst"
)

func PredictiveDefaultSmall(possibleLen int) int {
	return Predictive(
		possibleLen,
		percentconst.OnePointTwoPercentIncrement,
		constants.Capacity4)
}
