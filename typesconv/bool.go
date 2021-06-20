package typesconv

func BoolPtr(val bool) *bool {
	return &val
}

// BoolPtrToSimple if nil then 0
func BoolPtrToSimple(val *bool) bool {
	if val == nil {
		return false
	}

	return *val
}

// BoolPtrToSimpleDef if nil then 0
func BoolPtrToSimpleDef(val *bool, defVal bool) bool {
	if val == nil {
		return defVal
	}

	return *val
}

// BoolPtrToDefPtr if nil then 0
func BoolPtrToDefPtr(val *bool, defVal bool) *bool {
	if val == nil {
		return &defVal
	}

	return val
}

// BoolPtrDefValFunc if nil then executes returns defValFunc result as pointer
func BoolPtrDefValFunc(val *bool, defValFunc func() bool) *bool {
	if val == nil {
		result := defValFunc()

		return &result
	}

	return val
}
