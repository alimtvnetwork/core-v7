package errcore

import "fmt"

func MessageVarThree(
	message string,
	var1 string,
	val1 interface{},
	var2 string,
	val2 interface{},
	var3 string,
	val3 interface{},
) string {
	return fmt.Sprintf(
		messageVar3Format,
		message,
		var1,
		var2,
		val1,
		val2,
		var3,
		val3,
	)
}
