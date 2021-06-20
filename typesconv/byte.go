package typesconv

import "gitlab.com/evatix-go/core/constants"

func BytePtr(val byte) *byte {
	return &val
}

// BytePtrToSimple if nil then 0
func BytePtrToSimple(val *byte) byte {
	if val == nil {
		return constants.Zero
	}

	return *val
}

// BytePtrToSimpleDef if nil then 0
func BytePtrToSimpleDef(val *byte, defVal byte) byte {
	if val == nil {
		return defVal
	}

	return *val
}

// BytePtrToDefPtr if nil then 0
func BytePtrToDefPtr(val *byte, defVal byte) *byte {
	if val == nil {
		return &defVal
	}

	return val
}

// BytePtrDefValFunc if nil then executes returns defValFunc result as pointer
func BytePtrDefValFunc(val *byte, defValFunc func() byte) *byte {
	if val == nil {
		result := defValFunc()

		return &result
	}

	return val
}
