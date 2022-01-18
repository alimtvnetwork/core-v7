package errcore

func VarThreeNoType(
	var1 string,
	val1 interface{},
	var2 string,
	val2 interface{},
	var3 string,
	val3 interface{},
) string {
	return VarThree(
		false,
		var1, val1,
		var2, val2,
		var3, val3)
}
