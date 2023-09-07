package main

import (
	"errors"
	"fmt"

	"gitlab.com/auk-go/core/isany"
)

func nullTesting02() {
	var ex error
	var inx, inx2 *int
	items := []interface{}{
		nil,
		nil,
		errors.New("x"),
		ex,
		inx,
		inx2,
	}

	fmt.Println("DefinedBoth(nil, errors.New(\"x\")) = ", isany.DefinedBoth(nil, errors.New("x")))
	fmt.Println("NullBoth(nil,errors.New(\"x\")) = ", isany.NullBoth(nil, errors.New("x")))
	fmt.Println("NullBoth(nil,nil) = ", isany.NullBoth(nil, nil))
	fmt.Println("DefinedBoth(nil,nil) = ", isany.DefinedBoth(nil, nil))
	fmt.Println("NullBoth(inx,inx2) = ", isany.NullBoth(inx, inx2))

	for i, item := range items {
		fmt.Println(i, "- {(DefinedBoth, NullBoth, ==nil, %T} == ",
			isany.DefinedBoth(item, item),
			isany.NullBoth(item, item),
			item == nil,
			fmt.Sprintf("%T", item))
	}

	fmt.Println("----------------------")

	for i, item := range items {
		NullPrinter(i, item)
	}
}
