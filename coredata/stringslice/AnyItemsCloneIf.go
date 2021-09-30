package stringslice

func AnyItemsCloneIf(
	isClone bool,
	additionalCap int,
	slice []interface{},
) (newSlice []interface{}) {
	if slice == nil && !isClone {
		return []interface{}{}
	}

	if !isClone {
		return slice
	}

	return AnyItemsCloneUsingCap(additionalCap, slice)
}
