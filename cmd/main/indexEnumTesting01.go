package main

import (
	"fmt"

	"gitlab.com/auk-go/core/enums/versionindexes"
)

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
