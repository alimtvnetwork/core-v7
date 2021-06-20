package conditional

func NilCheck(
	canBeEmpty interface{},
	onNil interface{},
	onNonNil interface{},
) interface{} {
	if canBeEmpty == nil {
		return onNil
	}

	return onNonNil
}
