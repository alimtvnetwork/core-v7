package regexnew

// MatchErrorLock
//
// creates new regex using lock
// and then calls match.
// On condition mismatch returns error
// or else nil
func MatchErrorLock(regex, comparing string) error {
	regEx, err := CreateLock(regex)

	if regEx != nil && regEx.MatchString(comparing) {
		return nil
	}

	return regExMatchValidationError(
		regex,
		comparing,
		err,
		regEx)
}
