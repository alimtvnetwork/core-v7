package conditional

func DefOnNil(
	canBeEmpty interface{},
	onNonNil interface{},
) interface{} {
	if canBeEmpty == nil {
		return onNonNil
	}

	return canBeEmpty
}
