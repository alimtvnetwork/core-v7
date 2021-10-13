package regexnew

import "regexp"

type (
	RegexValidationFunc func(regex *regexp.Regexp, lookingTerm string) bool
	CustomizeErr        func(
		regexPattern,
		matchLookingTerm string,
		err error,
		regexp *regexp.Regexp,
	) error
)
