package regexnew

import "regexp"

// CreateApplicableLock
//
//	calls Create with mutex lock and unlock.
func CreateApplicableLock(regularExpressionPattern string) (
	regEx *regexp.Regexp,
	err error,
	isApplicable bool,
) {
	regexMutex.Lock()
	defer regexMutex.Unlock()

	regex, err := Create(regularExpressionPattern)

	return regex, err, err != nil && regEx != nil
}
