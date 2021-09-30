package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/bytetype"
	"gitlab.com/evatix-go/core/corecomparator"
)

func enumTesting() {
	names, _ := corecomparator.Equal.MarshalJSON()

	fmt.Println(string(names))

	noteq := corecomparator.NotEqual

	noteq.UnmarshalJSON([]byte("3"))

	fmt.Println(noteq.Name())
	fmt.Println(noteq.NumberString())
	fmt.Println(noteq.NumberString())
	fmt.Println(bytetype.BasicEnumImpl.AppendPrependJoinValue(".", 1, 2))
}
