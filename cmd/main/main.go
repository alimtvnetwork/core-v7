package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

func main() {
	// fmt.Println(osconsts.IsWindows)
	// fmt.Println(osconsts.IsUnixGroup)

	items := []string{" Alim NewHashset first,", "alim 1,", "alim next, ", "0 alim ,"}
	collection := corestr.NewCollectionUsingStrings(&items)

	collection.Resize(100)
	fmt.Println("Capacity :", collection.Capacity())

	length := 10
	newItems := make([]string, 0, length+len(items))
	newItems = append(newItems, items...)

	for i := 0; i < length; i++ {
		str := fmt.Sprintf("%d . %s", i, "alim ")
		newItems = append(newItems, str)
	}

	// fmt.Println(newItems)
	// fmt.Println("Sorted : ", collection.Add("0. Next Alim").Sorted().String())

	cmap := corestr.NewCharCollectionMap(
		10,
		5)

	var onComplete corestr.OnCompleteCharCollectionMap = func(stringsMap *corestr.CharCollectionMap) {
		stringsMap.PrintLock(true)

		json := cmap.Json()

		fmt.Println("json1:\n", json.JsonString())

		h, e := stringsMap.ParseInjectUsingJson(json)

		fmt.Println("Data from JSON1:\n", h, e)

		hashset1 := cmap.HashsetByStringFirstCharLock("a")
		json2 := hashset1.Json()

		fmt.Println("json2:\n", json2.JsonString())

		h2, e2 := hashset1.ParseInjectUsingJson(json2)

		fmt.Println("Data from JSON2:\n", h2.List(), h2, e2)

		hCollections := stringsMap.HashsetsCollection()

		json3 := hCollections.Json()

		fmt.Println("json3:\n", json3.JsonString())

		h3, _ := hCollections.ParseInjectUsingJson(json3)

		fmt.Println("Data from JSON3:\n", h3.String())
	}

	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), onComplete)
}
