package regexnew

import "regexp"

// New creates regex if not already exist in dictionary.
//
// if any error then doesn't save to map and returns the error
func New(regularExpressionSyntax string) (*regexp.Regexp, error) {
	regex, has := regexMaps[regularExpressionSyntax]

	if has {
		return regex, nil
	}

	newRegex, err := regexp.Compile(regularExpressionSyntax)

	if err == nil {
		regexMaps[regularExpressionSyntax] = newRegex
	}

	return newRegex, err
}
