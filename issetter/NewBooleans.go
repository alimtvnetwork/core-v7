package issetter

// NewBooleans
//
//	Any false, final result returns as False or else True
func NewBooleans(isResults ...bool) Value {
	return CombinedBooleans(isResults...)
}
