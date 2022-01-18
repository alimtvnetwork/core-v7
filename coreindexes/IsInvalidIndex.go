package coreindexes

import "gitlab.com/evatix-go/core/constants"

func IsInvalidIndex(index int) bool {
	return index <= constants.InvalidIndex
}
