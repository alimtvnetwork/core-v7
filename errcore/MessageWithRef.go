package errcore

import "fmt"

func MessageWithRef(msg string, reference interface{}) string {
	return fmt.Sprintf(
		messageMapFormat,
		msg,
		reference)
}
