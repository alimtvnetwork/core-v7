package defaultcapacity

import "gitlab.com/evatix-go/core/constants"

// MaxLimit
//
// returns
//  - limit:   -1, returns predictive length based on wholeLength (max 100)
//  - limit: >= 0, returns limit if limit < wholeLength or
//          else returns wholeLength
func MaxLimit(wholeLength int, limit int) int {
	hasLimit := limit > constants.MinusOne

	if hasLimit && limit >= wholeLength {
		return wholeLength
	} else if hasLimit && limit < wholeLength {
		return limit
	}

	// no limit
	return OfSearch(wholeLength)
}
