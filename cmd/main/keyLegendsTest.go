package main

import (
	"fmt"

	"gitlab.com/auk-go/core/keymk"
)

func keyLegendsTest() {
	k := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.ShortLegends,
		false,
		"cimux",
		"main",
		"myg",
		"stateName")
	// fmt.Println(k.GroupItem("g", ""))

	// fmt.Println(k.Group("g"))
	// fmt.Println(k.GroupString("g"))

	// fmt.Println(k.ItemWithoutUser("itemId1"))
	// fmt.Println(k.ItemIntRange(5, 10))
	fmt.Println(k.GroupIntRange(5, 10))
	// fmt.Println(k.GroupItem("mygroup", "item"))
	//
	// fmt.Println(k.GroupItem("g", "id"))
	// fmt.Println(k.ItemWithoutUser("id"))
	// fmt.Println(k.ItemWithoutUser("idx"))
	// fmt.Println(k.GroupItemIntRange("newg", 5, 10))
	// fmt.Println(k.GroupItemIntRange("newG#s", 5, 10))
	//
	// fmt.Println(k.GroupIntRange(5, 10))
	// fmt.Println(k.GroupName())
	// fmt.Println(k.UserStringWithoutState(""))
	// fmt.Println(k.ItemIntRange( 10, 20))
	fmt.Println(k.UserStringWithoutState("mynewuser1"))
	fmt.Println(k.UpToState("my-user"))
}
