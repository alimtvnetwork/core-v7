package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/constants"
)

func stackTrackerTest2() {
	one := codestack.New(1)
	two := codestack.New(2)
	three := codestack.New(3)

	fmt.Println(one.JsonString())
	fmt.Println(two.JsonString())
	fmt.Println(three.JsonString())
}

func stackTrackerTest3() {
	collection := codestack.NewStacksDefaultCount(
		codestack.SkipNone)

	fmt.Println(collection.Length())
	// fmt.Println(collection.JoinJsonStrings(constants.NewLineUnix))
	// fmt.Println(collection.JoinFileWithLinesStrings(constants.NewLineUnix))

	st2 := collection.ConcatNewUsingSkip(0)

	fmt.Println(st2.JoinJsonStrings(constants.NewLineUnix))

	trace3 := codestack.
		NewStacksCollection().
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

func stackTrackerTest4() {
	stackTrackerTest3()
}

func stackTrackerTest5() {
	stackTrackerTest4()
}
