package msgtype

func Var2NoType(
	var1 string,
	val1 interface{},
	var2 string,
	val2 interface{},
) string {
	return Var2(
		false,
		var1, val1,
		var2, val2)
}
