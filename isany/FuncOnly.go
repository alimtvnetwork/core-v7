package isany

// FuncOnly
//
// Returns false for nil, struct, anything else
// Returns true for Func only
func FuncOnly(item interface{}) (isFunc bool) {
	isFunc, _ = Function(item)

	return isFunc
}
