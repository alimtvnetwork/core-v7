package coreindexes

import "gitlab.com/auk-go/core/constants"

func IsInvalidIndex(index int) bool {
	return index <= constants.InvalidIndex
}
