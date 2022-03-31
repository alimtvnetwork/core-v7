package regexnew

// IsMatchLock creates new regex using lock
// and then calls match
// if doesn't match or invalid regex then returns false.
func IsMatchLock(regex, comparing string) bool {
	regEx, _ := CreateLock(regex)

	if regEx != nil && regEx.MatchString(comparing) {
		return true
	}

	return false
}
