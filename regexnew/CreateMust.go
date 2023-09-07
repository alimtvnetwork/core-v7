package regexnew

import "regexp"

// CreateMust creates regex if not already exist in dictionary.
//
// if any error then panics
func CreateMust(regularExpressionSyntax string) *regexp.Regexp {
	regex, has := regexMaps[regularExpressionSyntax]

	if has {
		return regex
	}

	newRegex := regexp.MustCompile(regularExpressionSyntax)
	regexMaps[regularExpressionSyntax] = newRegex

	return newRegex
}
