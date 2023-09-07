package issetter

// CombinedBooleans
//
//  Any false, final result returns as False or else True
func CombinedBooleans(
	isConditions ...bool,
) Value {
	for _, isCondition := range isConditions {
		if !isCondition {
			return False
		}
	}

	return True
}
