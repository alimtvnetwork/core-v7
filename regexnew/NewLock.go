package regexnew

import "regexp"

// NewLock calls New with mutex lock and unlock.
func NewLock(regularExpressionSyntax string) (*regexp.Regexp, error) {
	regexMutex.Lock()
	defer regexMutex.Unlock()

	return New(regularExpressionSyntax)
}
