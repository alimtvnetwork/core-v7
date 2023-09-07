package main

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/namevalue"
)

func coreDynamicMapAnyItems() {
	fmt.Println(errcore.VarTwoNoType("s1", "ss", "s2", 2))
	fmt.Println(errcore.MessageVarTwo("current message", "s1", "ss", "s2", 2))
	fmt.Println(errcore.MessageVarMap(
		"current message",
		map[string]interface{}{
			"key1": 1,
			"key2": 1,
			"key3": "",
		}))

	fmt.Println(errcore.MessageNameValues(
		"current message",
		namevalue.Instance{
			Name:  "name1",
			Value: "nil",
		},
		namevalue.Instance{
			Name:  "name2",
			Value: 2,
		}))

	fmt.Println("MapAnyItems")
	mapAnyItems := coredynamic.NewMapAnyItems(200)
	collection := corestr.New.Collection.Cap(100)
	collection.Adds("alim-1", "alim-2", "alim-3", "alim-4")
	mapAnyItems.Add("alim-something", collection)
	mapAnyItems.Add("alim-something2", collection)
	mapAnyItems.Add("alim-something3", collection.ConcatNew(1, "alim 5"))
	mapAnyItems.Add("alim-something4", collection)
	mapAnyItems.Add("alim-something5", collection)
	mapAnyItems.Add("alim-something6", collection)
	mapAnyItems.Add("alim-something7", collection)
	mapAnyItems.Add("alim-something8", collection)
	mapAnyItems.Add("alim-something9", collection)

	splittedItems := mapAnyItems.GetPagedCollection(2)

	for _, splittedItem := range splittedItems {
		fmt.Println(splittedItem.AllKeys())
	}

	jsonResult := mapAnyItems.JsonPtr()
	emptyCollection4 := corestr.Empty.Collection()
	mapAnyItems.GetItemRef("alim-something3", emptyCollection4)
	fmt.Println("4", emptyCollection4)

	emptyMapAnyItems := coredynamic.EmptyMapAnyItems()
	emptyCollection3 := corestr.Empty.Collection()
	req := corejson.KeyAny{
		Key:    "alim-something3",
		AnyInf: emptyCollection3,
	}

	err := emptyMapAnyItems.JsonParseSelfInject(jsonResult)
	newJsonResult := emptyMapAnyItems.Json()
	fmt.Println(err)
	collectionJsonResult := emptyMapAnyItems.JsonResultOfKey("alim-something")

	err2 := emptyMapAnyItems.GetManyItemsRefs(req)
	fmt.Println("alim-something3, err:", err2)
	fmt.Println("\"alim-something3\":", req.AnyInf)
	fmt.Println("\"alim-something3\":", emptyCollection3)

	err3 := emptyMapAnyItems.GetUsingUnmarshallManyAt(req)
	fmt.Println("alim-something3, err:", err3)
	fmt.Println("\"alim-something3\":", req.AnyInf)
	fmt.Println("\"alim-something3\":", emptyCollection3)

	fmt.Println(jsonResult.JsonString())
	fmt.Println(newJsonResult.JsonString())
	fmt.Println("jsonResult == newJsonResult :", jsonResult.IsEqual(newJsonResult))
	fmt.Println(collectionJsonResult.JsonString())
	newLinkedList := corestr.Empty.LinkedList()

	newLinkedList.JsonParseSelfInject(collectionJsonResult)
	fmt.Println(newLinkedList)
	fmt.Println(mapAnyItems)

	anyCollection := coredynamic.NewAnyCollection(10)
	anyCollection.AddAnySliceFromSingleItem(splittedItems[0].AllKeys())
	fmt.Println(anyCollection)
}
