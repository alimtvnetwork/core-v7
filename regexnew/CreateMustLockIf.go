package regexnew

import "regexp"

func CreateMustLockIf(
	isLock bool,
	regularExpressionSyntax string,
) *regexp.Regexp {
	if isLock {
		regexMutex.Lock()

		defer regexMutex.Unlock()
	}

	return CreateMust(regularExpressionSyntax)
}
