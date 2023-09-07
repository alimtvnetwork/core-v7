package typesconv

import (
	"strconv"

	"gitlab.com/auk-go/core/constants"
)

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

func StringToBool(s string) bool {
	if s == "" {
		return false
	}

	switch s {
	case "yes", "Yes", "YES":
		return true
	case "no", "NO", "No":
		return false
	}

	isBool, err := strconv.ParseBool(s)

	if err != nil {
		return false
	}

	return isBool
}

func StringPointerToBool(s *string) bool {
	if s == nil || *s == "" {
		return false
	}

	return StringToBool(*s)
}

func StringPointerToBoolPtr(s *string) *bool {
	if s == nil || *s == "" {
		toFalse := false

		return &toFalse
	}

	result := StringToBool(*s)

	return &result
}

func StringToBoolPtr(s string) *bool {
	if s == "" {
		toFalse := false

		return &toFalse
	}

	result := StringToBool(s)

	return &result
}
