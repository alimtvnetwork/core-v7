package main

import (
	"fmt"
	"sync"

	"gitlab.com/auk-go/core/regexnew"
)

func lazyRegExTester01() {

	patterns := []string{
		"something",
		"\\d+",
		"\\d+",
		"something",
		"something",
		"something",
		"something",
		"something",
		"something",
		"something",
		"something",
		"something",
		"something",
		"5",
		"something",
		"something",
		"something",
		"4",
		"1",
		"2",
		"3",
	}

	wg := sync.WaitGroup{}
	wg.Add(len(patterns))

	genFunc := func(index int, p string) {
		fmt.Println(index, p, " --- getting created")
		lazyRegex := regexnew.New.LazyLock(p)

		fmt.Println(index, lazyRegex.FullString())

		wg.Done()
	}

	for i, pattern := range patterns {
		go genFunc(i, pattern)
	}

	wg.Wait()
}
