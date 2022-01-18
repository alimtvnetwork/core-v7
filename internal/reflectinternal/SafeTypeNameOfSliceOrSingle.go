package reflectinternal

func SafeTypeNameOfSliceOrSingle(
	isSingle bool,
	any interface{},
) string {
	if isSingle {
		return SafeTypeName(any)
	}

	return SafeSliceToTypeName(any)
}
