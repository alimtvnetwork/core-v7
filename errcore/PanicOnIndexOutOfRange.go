package errcore

func PanicOnIndexOutOfRange(length int, indexes []int) {
	lastIndex := length - 1

	//goland:noinspection GoNilness
	for _, indexValue := range indexes {
		if indexValue > lastIndex {
			OutOfRangeType.
				HandleUsingPanic(
					"Index cannot be out of range, error is not ignored",
					indexValue)
		}
	}
}
