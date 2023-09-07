package defaultcapacity

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/constants/percentconst"
)

// PredictiveDefault Result must be positive possibleLen * multiplier + constants.ArbitraryCapacity50.
//
// Less than zero yields constants.ArbitraryCapacity50
func PredictiveDefault(possibleLen int) int {
	return Predictive(
		possibleLen,
		percentconst.FiftyPercentIncrement,
		constants.ArbitraryCapacity30)
}
