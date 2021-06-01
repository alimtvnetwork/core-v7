package defaultcapacity

import (
	"gitlab.com/evatix-go/core/constants/percentconst"
)

// Result must be positive possibleLen * multiplier + additionalCap.
//
// Less than zero yields additionalCap
func PredictiveFiftyPercentIncrement(possibleLen int, additionalCap int) int {
	return Predictive(
		possibleLen,
		percentconst.FiftyPercentIncrement,
		additionalCap)
}
