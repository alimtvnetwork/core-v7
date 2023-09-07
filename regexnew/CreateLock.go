package regexnew

import (
	"regexp"
)

// CreateLock calls Create with mutex lock and unlock.
func CreateLock(regularExpressionPattern string) (*regexp.Regexp, error) {
	regexMutex.Lock()
	defer regexMutex.Unlock()

	return Create(regularExpressionPattern)
}
