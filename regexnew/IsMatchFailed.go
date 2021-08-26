package regexnew

// IsMatchFailed creates new regex using lock
// and then calls match
// if doesn't match returns true
func IsMatchFailed(regex, comparing string) bool {
	return !IsMatchLock(regex, comparing)
}
