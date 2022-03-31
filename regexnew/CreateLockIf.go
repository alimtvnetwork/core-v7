package regexnew

import "regexp"

// CreateLockIf calls Create with mutex lock and unlock if true.
func CreateLockIf(isLock bool, regularExpressionSyntax string) (*regexp.Regexp, error) {
	if isLock {
		regexMutex.Lock()

		defer regexMutex.Unlock()
	}

	return Create(regularExpressionSyntax)
}
