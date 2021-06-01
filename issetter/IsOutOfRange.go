package issetter

// IsOutOfRange n < Uninitialized.Value() || n > Set.Value()
func IsOutOfRange(n byte) bool {
	return n < Uninitialized.Value() || n > MaxByte()
}
