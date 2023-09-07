package regexnew

// MatchError
//
// creates new regex using without lock (use in vars)
// and then calls match.
// On condition mismatch returns error
// or else nil
func MatchError(regex, comparing string) error {
	regEx, err := Create(regex)

	if regEx != nil && regEx.MatchString(comparing) {
		return nil
	}

	return regExMatchValidationError(
		regex,
		comparing,
		err,
		regEx)
}
