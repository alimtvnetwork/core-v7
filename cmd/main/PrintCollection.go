package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

func PrintCollection(collection *corestr.Collection) {
	fmt.Println(collection.GetPagedCollection(3).String())

	fmt.Print("\n\nTake 5:\n\n")
	fmt.Println(collection.Take(5))
	fmt.Print("\n\n Skip 2:\n\n")
	fmt.Println(collection.Skip(2))
	fmt.Print("\n\n Skip 0:\n\n")
	fmt.Println(collection.Skip(0))
	fmt.Print("\n\n Take 0:\n\n")
	fmt.Println(collection.Take(0))
	fmt.Print("\n\n Skip(5).Take(2):\n\n")
	fmt.Println(collection.Skip(5).Take(2))

}
