package conditional

// Deprecated: Use NilVal[T] instead for type-safe nil checking.
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
