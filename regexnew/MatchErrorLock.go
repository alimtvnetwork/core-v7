package regexnew

import "fmt"

// MatchErrorLock creates new regex using lock
// and then calls match
// if doesn't match then returns error
func MatchErrorLock(regex, comparing string) error {
	regEx, err := NewLock(regex)

	if regEx != nil && regEx.MatchString(comparing) {
		return nil
	}

	if err != nil {
		return fmt.Errorf(
			"[%q], regex compile failed / invalid cannot match with [%q]",
			err.Error(),
			comparing)
	}

	if regEx == nil {
		return fmt.Errorf(
			"given regex [%q] invalid cannot match with [%q]",
			regex,
			comparing)
	}

	return fmt.Errorf(
		"given regex [%q] doesn't match with [%q]",
		regex,
		comparing)
}
