package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/codestack"
)

func stackTrackerTest7() {
	cases := []string{
		"gitlab.com/evatix-go/errorwrapper/trydo.WrapPanic.func1",
		"trydo.WrapPanic.func1",
		"evatix-go/errorwrapper/trydo.WrapPanic.func1",
	}

	for i, s := range cases {
		fmt.Println("index :", i)
		fmt.Println(s, "= {")
		packageName, methodName := codestack.MethodNamePackageName(s)
		fmt.Println("   pkg=", packageName)
		fmt.Println("   methodName=", methodName)
		fmt.Println("}\n-------")
	}
}
