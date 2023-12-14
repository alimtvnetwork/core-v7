package main

import (
	"fmt"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/errcore"
)

func readWriteTest01() {
	fmt.Println(chmodhelper.FileModeFriendlyString(0777))
	thisFileRw := chmodhelper.New.SimpleFileReaderWriter.DefaultCleanPath(
		false,
		"cmd/main/main.go",
	)

	fmt.Println(thisFileRw)
	jsonResult := thisFileRw.Json()

	fmt.Println(jsonResult.JsonString())
	anotherRw := chmodhelper.New.SimpleFileReaderWriter.Path(
		false,
		0111, 0111,
		"dwdddw",
	)
	err := anotherRw.JsonParseSelfInject(jsonResult.Ptr())
	errcore.HandleErr(err)
	fmt.Println("unmarshalled", anotherRw.JsonPtr().JsonString())

	fmt.Println(thisFileRw.JsonPtr().JsonString())
	fmt.Println(thisFileRw.ReadStringMust())
}
