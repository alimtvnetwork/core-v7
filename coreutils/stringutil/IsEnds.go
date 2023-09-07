package stringutil

// IsEnds
//
// searches for case-sensitive terms
func IsEnds(
	content,
	endsWith string,
) bool {
	return IsEndsWith(
		content,
		endsWith,
		false)
}
