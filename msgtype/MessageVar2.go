package msgtype

import "fmt"

func MessageVar2(
	message string,
	var1 string,
	val1 interface{},
	var2 string,
	val2 interface{},
) string {
	return fmt.Sprintf(
		messageVar2Format,
		message,
		var1,
		var2,
		val1,
		val2)
}
