package main

import (
	"fmt"

	"gitlab.com/auk-go/core/keymk"
)

func testMakerTesting() {
	key := keymk.NewKey.All(keymk.BracketJoinerOption, "alim")
	key.AppendChainStrings("2", "alim-4")

	fmt.Println(key.String())

	fmt.Println(key.Compile("hello-world"))
	fmt.Println(key.Compile("hello-5-world"))
	fmt.Println(key.Compile("hello-5-world"))
	fmt.Println(key.AppendChainStrings("alim3"))
	fmt.Println(key.AppendChainStrings("alim4"))
	fmt.Println(key.String())
	fmt.Println(key.Finalized("alim{complete}"))
	fmt.Println(key.String())
	fmt.Println(key.CompileStrings("alim4new"))

	key2 := key.ClonePtr("alim4new").AppendChain("alim5new")

	fmt.Println(key2.Finalized("|BreakPoint|"))

	fmt.Println(key2.CompileKeys(key))
	// fmt.Println(key.AppendChainStrings("alim4new"))
}
