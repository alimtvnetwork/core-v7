package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

func main() {
	x := corestr.SimpleSlice{Items: []string{"a", "b"}}
	result := x.Json()

	emptyX := corestr.EmptySimpleSlice()
	// result.Error = errors.New("alim errs")

	resultOfResult := result.Json()
	coreJson2, err := resultOfResult.UnmarshalResult()
	fmt.Println(err, coreJson2.Error)
	emptyX.JsonParseSelfInject(coreJson2)
	fmt.Println(result.Json().JsonString())
	fmt.Println(resultOfResult.JsonString())
	fmt.Println(coreJson2.JsonString())
	fmt.Println(emptyX.Json().String())
	fmt.Println(string(resultOfResult.ValuesNonPtr()))
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
