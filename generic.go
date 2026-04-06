package core

// EmptySlice returns an empty slice of type T.
//
// Usage:
//
//	ints := core.EmptySlice[int]()       // returns []int
//	strs := core.EmptySlice[string]()    // returns []string
func EmptySlice[T any]() []T {
	return make([]T, 0)
}

// SliceByLength returns a zero-valued slice of type T with the given length.
func SliceByLength[T any](length int) []T {
	return make([]T, length)
}

// SliceByCapacity returns a slice of type T with the given length and capacity.
func SliceByCapacity[T any](length, cap int) []T {
	return make([]T, length, cap)
}

// Deprecated: Use EmptySlice instead.
func EmptySlicePtr[T any]() *[]T {
	s := make([]T, 0)
	return &s
}

// Deprecated: Use SliceByLength instead.
func SlicePtrByLength[T any](length int) *[]T {
	s := make([]T, length)
	return &s
}

// Deprecated: Use SliceByCapacity instead.
func SlicePtrByCapacity[T any](length, cap int) *[]T {
	s := make([]T, length, cap)
	return &s
}

// EmptyMapPtr returns a pointer to an empty map of type map[K]V.
// It replaces EmptyStringsMapPtr, EmptyStringToIntMapPtr and similar per-type functions.
func EmptyMapPtr[K comparable, V any]() *map[K]V {
	m := make(map[K]V)
	return &m
}
