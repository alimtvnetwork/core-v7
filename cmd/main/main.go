package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

func main() {
	rwx := chmodins.RwxOwnerGroupOther{
		Owner: "rwx",
		Group: "r--",
		Other: "-wx",
	}

	fmt.Println(rwx.String())
	wrapper, _ := chmodhelper.NewUsingRwxOwnerGroupOther(&rwx)

	fmt.Println(wrapper.ToRwxOwnerGroupOther().String())

	// items := &[]string{
	// 	"00",
	// 	"01",
	// 	"02",
	// 	"03",
	// 	"04",
	// 	"05",
	// 	"06",
	// 	"07",
	// 	"08",
	// 	"09",
	// 	"10",
	// 	"11",
	// 	"12",
	// }
	//
	// // collectionPtr := corestr.NewCollectionPtrUsingStrings(items, 0)
	// collection := corestr.NewCollectionUsingStrings(items, false)
	// jsonResults := corejson.NewResultsCollectionUsingJsoners(1, collection)
	// jsonResultFromResults := jsonResults.Json()
	//
	// fmt.Println(jsonResultFromResults.JsonString())
	//
	// res2 := corejson.EmptyResultsCollection()
	//
	// res2.ParseInjectUsingJson(jsonResultFromResults)
	//
	// fmt.Println(res2.Json().JsonString())
	// collect2 := corestr.EmptyCollection()
	//
	// // res2.InjectIntoAt(0, collect2)
	// // res2.UnmarshalAt(0, collect2)
	// res2.UnmarshalIntoSameIndex(collect2)
	//
	// fmt.Println(collect2)

	// PrintCollection(collection)

	// moredata3 := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}
	// dynamicCollection := coredynamic.NewDynamicCollection(100)
	//
	// dynamicCollection.AddAny(1, true)
	// dynamicCollection.AddAny(2, true)
	// dynamicCollection.AddAny(3, true)
	// dynamicCollection.AddAny(4, true)
	// dynamicCollection.AddAny(5, true)
	// dynamicCollection.AddAny(moredata3, true)
	// maps, _ := result.
	// 	MapToKeyVal()

	// fmt.Printf(constants.SprintPropertyNameValueFormat, maps)
	// fmt.Println(dynamicCollection.RemoveAt(1))
	// fmt.Println(dynamicCollection.Length())
	// fmt.Println(dynamicCollection.ListStrings())
	// fmt.Println(dynamicCollection.StringJson())
	// fmt.Println(dynamicCollection.At(2).IsPrimitive())
	// fmt.Println(dynamicCollection.Items())
	//
	// _ = bytetype.Variant(1).RangesInvalidMessage()
	// bt := bytetype.Variant(1)
	// fmt.Println(bt.StringJsonMust())
	// fmt.Println(bt.Value())
	// fmt.Println(bt.IsValidRange())
	//
	// fmt.Println(chmodhelper.MergeRwxWildcardWithFixedRwx("-w*", "r-x"))
	// fmt.Println(constants.MinInt)
	// t3 := bytetype.T3
	// b, e := t3.MarshalJSON()
	// fmt.Println(string(b), e)
	//
	// fmt.Println(t3.RangesInvalidMessage())
	// fmt.Println(t3.String())
	//
	// fmt.Println(bytetype.BasicEnum2Impl.StringJson(bytetype.Ab4))

	// fmt.Println(result.IsSliceOrArray())
	// fmt.Println(result.IsMap())
	// fmt.Println(result.InvalidError())
	// fmt.Println(result.GetErrorOnTypeMismatch(reflect.TypeOf(map[int]string{}), true))
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
