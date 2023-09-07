package regexnew

import (
	"regexp"
	"sync"

	"gitlab.com/auk-go/core/constants"
)

var (
	regexMutex    = sync.Mutex{}
	lazyRegexLock = sync.Mutex{}
	regexMaps     = make(
		map[string]*regexp.Regexp,
		constants.ArbitraryCapacity30)
	lazyRegexOnceMap = lazyRegexMap{
		items: make(
			map[string]*LazyRegex,
			constants.ArbitraryCapacity30),
	}

	New = newCreator{}
)
