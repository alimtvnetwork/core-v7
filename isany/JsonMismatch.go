package isany

// JsonMismatch
//
//	Inverse of JsonEqual
func JsonMismatch(
	left, right interface{},
) bool {
	return !JsonEqual(left, right)
}
