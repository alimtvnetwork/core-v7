package typesconv

import "gitlab.com/auk-go/core/constants"

func FloatPtr(val float32) *float32 {
	return &val
}

// FloatPtrToSimple if nil then 0
func FloatPtrToSimple(val *float32) float32 {
	if val == nil {
		return constants.Zero
	}

	return *val
}

// FloatPtrToSimpleDef if nil then 0
func FloatPtrToSimpleDef(val *float32, defVal float32) float32 {
	if val == nil {
		return defVal
	}

	return *val
}

// FloatPtrToDefPtr if nil then 0
func FloatPtrToDefPtr(val *float32, defVal float32) *float32 {
	if val == nil {
		return &defVal
	}

	return val
}

// FloatPtrDefValFunc if nil then executes returns defValFunc result as pointer
func FloatPtrDefValFunc(val *float32, defValFunc func() float32) *float32 {
	if val == nil {
		result := defValFunc()

		return &result
	}

	return val
}
