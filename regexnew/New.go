package regexnew

import "regexp"

// New creates regex if not already exist in dictionary.
//
// if any error then doesn't save to map and returns the error
func New(regularExpressionPattern string) (*regexp.Regexp, error) {
	regex, has := regexMaps[regularExpressionPattern]

	if has {
		return regex, nil
	}

	newRegex, err := regexp.Compile(regularExpressionPattern)

	if err == nil {
		regexMaps[regularExpressionPattern] = newRegex
	}

	return newRegex, err
}
