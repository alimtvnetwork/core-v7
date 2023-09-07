package main

import (
	"fmt"

	"gitlab.com/auk-go/core/isany"
)

func NullPrinter(index int, item interface{}) {
	fmt.Println(index, "- {(defined, isnull, ==nil, %T} == ",
		isany.Defined(item),
		isany.Null(item),
		item == nil,
		fmt.Sprintf("%T", item))
}
