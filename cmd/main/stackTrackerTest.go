package main

import (
	"fmt"

	"gitlab.com/auk-go/core/codestack"
)

func stackTrackerTest() {
	stackTrackerTest2()
	one := codestack.New(1)
	two := codestack.New(2)
	three := codestack.New(3)

	fmt.Println(one.String())
	fmt.Println(two.String())
	fmt.Println(three.String())
}
