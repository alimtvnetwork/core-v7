package main

import (
	"fmt"

	"gitlab.com/auk-go/core/errcore"
)

func errTypePrintTest01() {
	errType := errcore.BytesAreNilOrEmptyType

	msg := errType.
		MsgCsvRef("some 1", "alim-1")

	fmt.Println(msg)

	msg = errType.
		Combine("", "alim-2 no msg")

	fmt.Println(msg)

	msg = errType.
		Combine("", "")

	fmt.Println(msg)

	msg = errType.
		Combine("some 2", "alim-1")

	fmt.Println(msg)

	msg = errType.
		Combine("", "alim-5 final")

	fmt.Println(msg)
}
