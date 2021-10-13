package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/enums/versionindexes"
)

func compareEnumTesting02() {
	eq := corecomparator.Equal

	err := eq.
		OnlySupportedErr(
			"dining doesn't support more",
			corecomparator.Inconclusive,
			corecomparator.NotEqual)

	fmt.Println(err)
}

func indexEnumTesting01() {
	build := versionindexes.Patch

	toJson := build.Json()

	fmt.Println(toJson.PrettyJsonString())
	minor := versionindexes.Minor
	fmt.Println(minor.Name(), minor)
	minor.JsonParseSelfInject(&toJson)
	fmt.Println(minor)

	fmt.Println(minor.NameValue())
}
