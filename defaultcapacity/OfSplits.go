package defaultcapacity

import (
	"gitlab.com/evatix-go/core/constants"
)

// OfSplits returns max as 100
func OfSplits(wholeTextLength int, limits int) int {
	if limits > constants.MinusOne {
		return limits
	}

	return OfSearch(wholeTextLength)
}
