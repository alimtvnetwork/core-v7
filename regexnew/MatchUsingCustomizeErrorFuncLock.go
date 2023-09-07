package regexnew

// MatchUsingCustomizeErrorFuncLock
//
// creates new regex using lock
// and then calls match.
// On condition mismatch returns error
// or else nil
func MatchUsingCustomizeErrorFuncLock(
	regexPattern, comparing string,
	matchFunc RegexValidationFunc,
	customizeErrFunc CustomizeErr, // can be nil, if nil then use default one
) error {
	regEx, err := CreateLock(regexPattern)

	if regEx != nil && matchFunc(regEx, comparing) {
		return nil
	}

	if customizeErrFunc == nil {
		return regExMatchValidationError(
			regexPattern,
			comparing,
			err,
			regEx)
	}

	return customizeErrFunc(
		regexPattern,
		comparing,
		err,
		regEx)
}
