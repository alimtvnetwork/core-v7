package typesconv

import "gitlab.com/evatix-go/core/constants"

func StringPtr(val string) *string {
	return &val
}

// StringPtrToSimple if nil then 0
func StringPtrToSimple(val *string) string {
	if val == nil {
		return constants.EmptyString
	}

	return *val
}

// StringPtrToSimpleDef if nil then 0
func StringPtrToSimpleDef(val *string, defVal string) string {
	if val == nil {
		return defVal
	}

	return *val
}

// StringPtrToDefPtr if nil then 0
func StringPtrToDefPtr(val *string, defVal string) *string {
	if val == nil {
		return &defVal
	}

	return val
}

// StringPtrDefValFunc if nil then executes returns defValFunc result as pointer
func StringPtrDefValFunc(val *string, defValFunc func() string) *string {
	if val == nil {
		result := defValFunc()

		return &result
	}

	return val
}
