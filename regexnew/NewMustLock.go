package regexnew

import "regexp"

func NewMustLock(regularExpressionSyntax string) *regexp.Regexp {
	regexMutex.Lock()
	defer regexMutex.Unlock()

	return CreateMust(regularExpressionSyntax)
}
