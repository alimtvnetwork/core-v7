package regexnew

// MatchUsingFuncErrorLock
//
// creates new regex using lock
// and then calls match.
// On condition mismatch returns error
// or else nil
func MatchUsingFuncErrorLock(
	regexPattern, comparing string,
	matchFunc RegexValidationFunc,
) error {
	regEx, err := CreateLock(regexPattern)

	if regEx != nil && matchFunc(regEx, comparing) {
		return nil
	}

	return regExMatchValidationError(
		regexPattern,
		comparing,
		err,
		regEx)
}
