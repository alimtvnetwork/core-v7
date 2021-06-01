package defaultcapacity

import (
	"gitlab.com/evatix-go/core/constants"
)

// OfSplits returns max as 100
func OfSplits(wholeTextLength int, limits int) int {
	if limits > constants.MinusOne {
		return limits
	}

	defaultCapacity := wholeTextLength

	if wholeTextLength > constants.ArbitraryCapacity1000 {
		defaultCapacity = constants.ArbitraryCapacity100
	} else if wholeTextLength > constants.ArbitraryCapacity250 {
		defaultCapacity = wholeTextLength / constants.N20
	} else if wholeTextLength >= constants.ArbitraryCapacity100 {
		defaultCapacity = wholeTextLength / constants.N10
	} else {
		defaultCapacity = wholeTextLength / constants.N5
	}

	return defaultCapacity
}
