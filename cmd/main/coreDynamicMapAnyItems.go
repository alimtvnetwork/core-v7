package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/msgtype"
)

func coreDynamicMapAnyItems() {
	fmt.Println(msgtype.Var2NoType("s1", "ss", "s2", 2))
	fmt.Println(msgtype.MessageVar2("current message", "s1", "ss", "s2", 2))
	fmt.Println(msgtype.MessageVarMap(
		"current message",
		map[string]interface{}{
			"key1": 1,
			"key2": 1,
			"key3": "",
		}))

	fmt.Println(msgtype.MessageNameValues(
		"current message",
		msgtype.NameVal{
			Name:  "name1",
			Value: "nil",
		},
		msgtype.NameVal{
			Name:  "name2",
			Value: 2,
		}))

	fmt.Println("MapAnyItems")
	mapAnyItems := coredynamic.NewMapAnyItems(200)
	collection := corestr.NewCollection(100)
	collection.Adds("alim-1", "alim-2", "alim-3", "alim-4")
	mapAnyItems.Add("alim-something", collection)
	mapAnyItems.Add("alim-something2", collection)
	mapAnyItems.Add("alim-something3", collection.ConcatNew(1, "alim 5"))
	jsonResult := mapAnyItems.Json()
	emptyCollection4 := corestr.EmptyCollection()
	mapAnyItems.GetItemRef("alim-something3", emptyCollection4)
	fmt.Println("4", emptyCollection4)

	emptyMapAnyItems := coredynamic.EmptyMapAnyItems()
	emptyCollection3 := corestr.EmptyCollection()
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
	newLinkedList := corestr.EmptyLinkedList()

	newLinkedList.JsonParseSelfInject(collectionJsonResult)
	fmt.Println(newLinkedList)

	fmt.Println(mapAnyItems)
}
