package core

// Deprecated: Use EmptySlicePtr[any]() instead.
func EmptyAnysPtr() *[]any {
	return EmptySlicePtr[any]()
}

// Deprecated: Use EmptySlicePtr[float32]() instead.
func EmptyFloat32Ptr() *[]float32 {
	return EmptySlicePtr[float32]()
}

// Deprecated: Use EmptySlicePtr[float64]() instead.
func EmptyFloat64Ptr() *[]float64 {
	return EmptySlicePtr[float64]()
}

// Deprecated: Use EmptySlicePtr[bool]() instead.
func EmptyBoolsPtr() *[]bool {
	return EmptySlicePtr[bool]()
}

// Deprecated: Use EmptySlicePtr[int]() instead.
func EmptyIntsPtr() *[]int {
	return EmptySlicePtr[int]()
}

// Deprecated: Use []byte{} instead.
func EmptyBytePtr() []byte {
	return []byte{}
}

// Deprecated: Use EmptyMapPtr[string, string]() instead.
func EmptyStringsMapPtr() *map[string]string {
	return EmptyMapPtr[string, string]()
}

// Deprecated: Use EmptyMapPtr[string, int]() instead.
func EmptyStringToIntMapPtr() *map[string]int {
	return EmptyMapPtr[string, int]()
}

// Deprecated: Use EmptySlicePtr[string]() instead.
func EmptyStringsPtr() *[]string {
	return EmptySlicePtr[string]()
}

// Deprecated: Use EmptySlicePtr[*string]() instead.
func EmptyPointerStringsPtr() *[]*string {
	return EmptySlicePtr[*string]()
}

// Deprecated: Use SlicePtrByLength[string](length) instead.
func StringsPtrByLength(length int) *[]string {
	return SlicePtrByLength[string](length)
}

// Deprecated: Use SlicePtrByCapacity[string](length, cap) instead.
func StringsPtrByCapacity(length, cap int) *[]string {
	return SlicePtrByCapacity[string](length, cap)
}

// Deprecated: Use SlicePtrByCapacity[*string](length, cap) instead.
func PointerStringsPtrByCapacity(length, cap int) *[]*string {
	return SlicePtrByCapacity[*string](length, cap)
}
