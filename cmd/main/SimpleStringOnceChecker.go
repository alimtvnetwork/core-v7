package main

import (
	"fmt"
	"strconv"

	"gitlab.com/auk-go/core/coredata/corestr"
)

type simpleStringTester struct {
	count  int
	simple corestr.SimpleStringOnce
}

func (it *simpleStringTester) Something() (val string, count int) {
	if it.simple.IsInitialized() {
		return it.simple.Value(), it.count
	}

	it.count++

	return it.
			simple.
			GetSetOnce("some value " + strconv.Itoa(it.count)),
		it.count
}

var (
	simpleStringTester2 = &simpleStringTester{}
)

func SimpleStringOnceChecker(testCount int) {
	for i := 0; i < testCount; i++ {
		v, count := simpleStringTester2.Something()

		fmt.Println(i, v, count)
	}
}
