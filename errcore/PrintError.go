package errcore

import "fmt"

func PrintError(
	err error,
) {
	if err != nil {
		fmt.Print(err)
	}
}
