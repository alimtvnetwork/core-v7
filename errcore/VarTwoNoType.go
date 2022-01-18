package errcore

func VarTwoNoType(
	var1 string,
	val1 interface{},
	var2 string,
	val2 interface{},
) string {
	return VarTwo(
		false,
		var1, val1,
		var2, val2)
}
