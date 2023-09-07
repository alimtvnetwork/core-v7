package main

import (
	"fmt"

	"gitlab.com/auk-go/core/codestack"
)

func stackTrackerTest7() {
	cases := []string{
		"gitlab.com/evatix-go/errorwrapper/trydo.WrapPanic.func1",
		"trydo.WrapPanic.func1",
		"evatix-go/errorwrapper/trydo.WrapPanic.func1",
		"",
		"something",
		"something.new",
		".....new",
		"/",
		" ",
		" /",
		" . ",
	}

	for i, s := range cases {
		fmt.Println("index :", i)
		fmt.Println(s, "= {")
		signature, packageName, methodName := codestack.MethodNamePackageName(s)
		fmt.Println("   signature   =", signature)
		fmt.Println("   pkg         =", packageName)
		fmt.Println("   methodName  =", methodName)
		fmt.Println("}\n-------")
	}
}
