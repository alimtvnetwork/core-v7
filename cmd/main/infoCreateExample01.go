package main

import (
	"fmt"

	"gitlab.com/auk-go/core/coretaskinfo"
	"gitlab.com/auk-go/core/errcore"
)

func infoCreateExample01() {
	info := coretaskinfo.New.Info.Default(
		"some name",
		"some desc",
		"some url")

	fmt.Println(info.LazyMapPrettyJsonString())

	infoExamples := coretaskinfo.New.Info.Examples(
		"some name",
		"some desc",
		"some url",
		"some examples1 \"some command\"", "some examples 2")

	fmt.Println(infoExamples.LazyMapPrettyJsonString())

	infoNoExamples := coretaskinfo.New.Info.Examples(
		"no exmaple some name",
		"some desc",
		"some url",
	)

	fmt.Println(infoNoExamples.LazyMapPrettyJsonString())

	infoNoExamples2, parseErr := coretaskinfo.New.Info.Deserialized(
		infoNoExamples.JsonPtr().Bytes)

	errcore.HandleErr(parseErr)
	fmt.Println(infoNoExamples2.PrettyJsonStringWithPayloads([]byte("some payloads2")))

	infoNoExamples2 = nil

	fmt.Println(infoNoExamples2.PrettyJsonStringWithPayloads([]byte("some payloads3")))
}
