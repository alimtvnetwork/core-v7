package stringslice

func AnyItemsCloneUsingCap(
	newAdditionalCap int,
	slice []interface{},
) []interface{} {
	newSlice := make(
		[]interface{},
		0,
		newAdditionalCap+len(slice))

	if len(slice) == 0 {
		return newSlice
	}

	newSlice = append(newSlice, slice...)

	return newSlice
}
