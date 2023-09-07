package main

import (
	"errors"
	"fmt"

	"gitlab.com/auk-go/core/isany"
)

func nullTesting01() {
	var ex error
	var inx *int
	items := []interface{}{
		nil,
		errors.New(""),
		ex,
		inx,
	}

	for i, item := range items {
		fmt.Println(i, "- {(defined, isnull, ==nil, %T} == ",
			isany.Defined(item),
			isany.Null(item),
			item == nil,
			fmt.Sprintf("%T", item))
	}

	fmt.Println("----------------------")

	for i, item := range items {
		NullPrinter(i, item)
	}
}
