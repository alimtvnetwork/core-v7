package regexnew

import (
	"regexp"
)

// NewLock calls New with mutex lock and unlock.
func NewLock(regularExpressionPattern string) (*regexp.Regexp, error) {
	regexMutex.Lock()
	defer regexMutex.Unlock()

	return New(regularExpressionPattern)
}
