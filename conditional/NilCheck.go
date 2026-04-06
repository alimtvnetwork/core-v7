package conditional

// This function uses 'any' and loses type safety.
func NilCheck(
	canBeEmpty any,
	onNil any,
	onNonNil any,
) any {
	if canBeEmpty == nil {
		return onNil
	}

	return onNonNil
}
