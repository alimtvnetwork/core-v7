package isany

// NotNull
//
// # Returns true for not nil
//
// Reference : https://stackoverflow.com/a/43896204
func NotNull(item interface{}) bool {
	return !Null(item)
}
