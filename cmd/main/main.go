package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corejson"
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

	// collectionPtr := corestr.NewCollectionPtrUsingStrings(items, 0)
	collection := corestr.NewCollectionUsingStrings(items, false)
	jsonResults := corejson.NewResultsCollectionUsingJsoners(1, collection)
	jsonResultFromResults := jsonResults.Json()

	fmt.Println(jsonResultFromResults.JsonString())

	res2 := corejson.EmptyResultsCollection()

	res2.ParseInjectUsingJson(jsonResultFromResults)

	fmt.Println(res2.Json().JsonString())
	collect2 := corestr.EmptyCollection()

	// res2.InjectIntoAt(0, collect2)
	// res2.UnmarshalAt(0, collect2)
	res2.UnmarshalIntoSameIndex(collect2)

	fmt.Println(collect2)

	// PrintCollection(collection)
}

func PrintCollection(collection *corestr.Collection) {
	fmt.Println(collection.GetPagedCollection(3).String())

	fmt.Print("\n\nTake 5:\n\n")
	fmt.Println(collection.Take(5))
	fmt.Print("\n\n Skip 2:\n\n")
	fmt.Println(collection.Skip(2))
	fmt.Print("\n\n Skip 0:\n\n")
	fmt.Println(collection.Skip(0))
	fmt.Print("\n\n Take 0:\n\n")
	fmt.Println(collection.Take(0))
	fmt.Print("\n\n Skip(5).Take(2):\n\n")
	fmt.Println(collection.Skip(5).Take(2))

}

func PrintCollectionPtr(collectionPtr *corestr.CollectionPtr) {
	fmt.Println(collectionPtr.GetPagedCollection(3).String())

	fmt.Print("\n\nTake 5:\n\n")
	fmt.Println(collectionPtr.Take(5))
	fmt.Print("\n\n Skip 2:\n\n")
	fmt.Println(collectionPtr.Skip(2))
	fmt.Print("\n\n Skip 0:\n\n")
	fmt.Println(collectionPtr.Skip(0))
	fmt.Print("\n\n Take 0:\n\n")
	fmt.Println(collectionPtr.Take(0))
	fmt.Print("\n\n Skip(5).Take(2):\n\n")
	fmt.Println(collectionPtr.Skip(5).Take(2))
}
