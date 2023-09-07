package typesconv

import "gitlab.com/auk-go/core/constants"

func IntPtr(val int) *int {
	return &val
}

// IntPtrToSimple if nil then 0
func IntPtrToSimple(val *int) int {
	if val == nil {
		return constants.Zero
	}

	return *val
}

// IntPtrToSimpleDef if nil then 0
func IntPtrToSimpleDef(val *int, defVal int) int {
	if val == nil {
		return defVal
	}

	return *val
}

// IntPtrToDefPtr if nil then 0
func IntPtrToDefPtr(val *int, defVal int) *int {
	if val == nil {
		return &defVal
	}

	return val
}

// IntPtrDefValFunc if nil then executes returns defValFunc result as pointer
func IntPtrDefValFunc(val *int, defValFunc func() int) *int {
	if val == nil {
		result := defValFunc()

		return &result
	}

	return val
}
