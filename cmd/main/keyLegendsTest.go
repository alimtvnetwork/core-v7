package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/keymk"
)

func keyLegendsTest() {
	k := keymk.NewKeyWithLegend(
		keymk.JoinerOption,
		keymk.ShortLegends,
		false,
		"cimux",
		"main",
		"myg")
	// fmt.Println(k.GroupItem("g", ""))

	// fmt.Println(k.Group("g"))
	// fmt.Println(k.GroupString("g"))

	// fmt.Println(k.Item("itemId1"))
	fmt.Println(k.ItemIntRange(5, 10))
	fmt.Println(k.GroupIntRange(5, 10))
	fmt.Println(k.GroupItem("mygroup", "item"))
	//
	// fmt.Println(k.GroupItem("g", "id"))
	// fmt.Println(k.Item("id"))
	// fmt.Println(k.Item("idx"))
	// fmt.Println(k.GroupItemIntRange("newg", 5, 10))
	// fmt.Println(k.GroupItemIntRange("newG#s", 5, 10))
	//
	// fmt.Println(k.GroupIntRange(5, 10))
	// fmt.Println(k.GroupName())
	// fmt.Println(k.UptoGroup(""))
	// fmt.Println(k.ItemIntRange( 10, 20))
	// fmt.Println(k.UptoGroup( "mynewuser1"))
}
