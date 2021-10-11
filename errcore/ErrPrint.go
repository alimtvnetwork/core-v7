package errcore

import "fmt"

func ErrPrint(
	err error,
) {
	if err != nil {
		fmt.Print(err)
	}
}
