package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

func main() {
	items := &[]string{
		"00",
		"01",
		"02",
		"03",
		"04",
		"05",
		"06",
		"07",
		"08",
		"09",
		"10",
		"11",
		"12",
	}

	collection := corestr.NewCollectionPtrUsingStrings(items, 0)

	fmt.Println(collection.GetPagedCollection(3).String())
}
