package main

import (
	"fmt"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/constants"
)

func stackTrackerTest3() {
	collection := codestack.NewStacksDefaultCount(
		codestack.SkipNone)

	fmt.Println(collection.Length())
	fmt.Println(collection.CodeStacksStringLimit(2))
	// fmt.Println(collection.JoinJsonStrings(constants.NewLineUnix))
	// fmt.Println(collection.JoinFileWithLinesStrings(constants.NewLineUnix))

	st2 := collection.ConcatNewUsingSkip(0)

	fmt.Println(st2.JoinJsonStrings(constants.NewLineUnix))

	trace3 := codestack.NewStacksCollection().
		AddsUsingSkipDefault(0)

	fmt.Println(trace3.
		JoinShortStrings(constants.NewLineUnix))
	fmt.Println("-------------")
	fmt.Println(collection.
		JoinFileWithLinesStrings(constants.NewLineUnix))

	fmt.Println("-------------")

	fmt.Println(collection.
		Reverse().JoinFileWithLinesStrings(constants.NewLineUnix))
	fmt.Println("-------------")

	fmt.Println(collection.Add(collection.First()).
		Reverse().JoinFileWithLinesStrings(constants.NewLineUnix))
}
