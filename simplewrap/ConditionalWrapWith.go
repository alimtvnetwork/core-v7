package simplewrap

import "fmt"

// ConditionalWrapWith
//
// Combine start and end char only if missing.
//
// If start missing then add start, if end missing then add end.
func ConditionalWrapWith(
	startChar byte,
	input string,
	endChar byte,
) string {
	length := len(input)
	
	if length == 0 {
		return fmt.Sprintf(
			"%c%c",
			startChar,
			endChar)
	}
	
	// more than 2 char
	firstChar := input[0]
	lastChar := input[length-1]
	isPresent :=
		firstChar == startChar &&
			lastChar == endChar
	
	if isPresent && length == 1 {
		// first one present and single char
		
		return fmt.Sprintf(
			"%s%c",
			input,
			endChar)
	}
	
	if isPresent {
		return input
	}
	
	isRightMissing :=
		firstChar == startChar &&
			lastChar != endChar
	
	if isRightMissing {
		return fmt.Sprintf(
			"%s%c",
			input,
			endChar)
	}
	
	isLeftMissing :=
		firstChar != startChar &&
			lastChar == endChar
	if isLeftMissing {
		return fmt.Sprintf(
			"%c%s",
			startChar,
			input,
		)
	}
	
	// both missing
	
	return fmt.Sprintf(
		"%c%s%c",
		startChar,
		input,
		endChar,
	)
}
