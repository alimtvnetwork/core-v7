package isany

// Defined
//
// # Returns true for not nil
//
// Reference : https://stackoverflow.com/a/43896204
func Defined(item interface{}) bool {
	return !Null(item)
}
