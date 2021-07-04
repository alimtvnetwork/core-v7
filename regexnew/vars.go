package regexnew

import (
	"regexp"
	"sync"

	"gitlab.com/evatix-go/core/constants"
)

var (
	regexMutex = sync.Mutex{}
	regexMaps  = make(
		map[string]*regexp.Regexp,
		constants.ArbitraryCapacity50)
)
