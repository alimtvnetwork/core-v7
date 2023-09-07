package issetter

// GetSetterByComparing
//
// returns true value if any of ranges value matches
func GetSetterByComparing(
	trueVal, falseVal Value,
	expectedVal interface{},
	trueRanges ...interface{},
) Value {
	for _, s := range trueRanges {
		if s == expectedVal {
			return trueVal
		}
	}

	return falseVal
}
