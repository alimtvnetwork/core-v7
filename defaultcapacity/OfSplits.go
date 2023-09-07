package defaultcapacity

import (
	"gitlab.com/auk-go/core/constants"
)

// OfSplits
//
// returns
//  - limit:   -1, returns predictive length based on wholeLength (max 100)
//  - limit: >= 0, returns limit if limit < wholeLength or
//          else returns predictive length based on wholeLength (max 100)
func OfSplits(wholeLength int, limit int) int {
	hasLimit := limit > constants.MinusOne

	if hasLimit && limit >= wholeLength {
		return OfSearch(wholeLength)
	} else if hasLimit && limit < wholeLength {
		return limit
	}

	// no limit
	return OfSearch(wholeLength)
}
