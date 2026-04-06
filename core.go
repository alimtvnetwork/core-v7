package core

func EmptyAnysPtr() *[]any {
	return EmptySlicePtr[any]()
}

func EmptyFloat32Ptr() *[]float32 {
	return EmptySlicePtr[float32]()
}

func EmptyFloat64Ptr() *[]float64 {
	return EmptySlicePtr[float64]()
}

func EmptyBoolsPtr() *[]bool {
	return EmptySlicePtr[bool]()
}

func EmptyIntsPtr() *[]int {
	return EmptySlicePtr[int]()
}

func EmptyBytePtr() []byte {
	return []byte{}
}

func EmptyStringsMapPtr() *map[string]string {
	return EmptyMapPtr[string, string]()
}

func EmptyStringToIntMapPtr() *map[string]int {
	return EmptyMapPtr[string, int]()
}

func EmptyStringsPtr() *[]string {
	return EmptySlicePtr[string]()
}

func EmptyPointerStringsPtr() *[]*string {
	return EmptySlicePtr[*string]()
}

func StringsPtrByLength(length int) *[]string {
	return SlicePtrByLength[string](length)
}

func StringsPtrByCapacity(length, cap int) *[]string {
	return SlicePtrByCapacity[string](length, cap)
}

func PointerStringsPtrByCapacity(length, cap int) *[]*string {
	return SlicePtrByCapacity[*string](length, cap)
}
