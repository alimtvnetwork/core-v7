package main

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/coredynamic"
)

func jsonResultUnmarshallingTest02() {
	mapAnyItems := getMapAnyItems()
	jsonResult := mapAnyItems.JsonPtr()
	var emptyMapResult *coredynamic.MapAnyItems

	err := jsonResult.Unmarshal(emptyMapResult)
	fmt.Println("err:", err)
	fmt.Println(emptyMapResult)

	var emptyMapResult2 coredynamic.MapAnyItems

	err2 := jsonResult.Unmarshal(emptyMapResult2)
	fmt.Println("err2:", err2)
	fmt.Println(emptyMapResult2)

	var emptyMapResult3 coredynamic.MapAnyItems

	err3 := jsonResult.Unmarshal(&emptyMapResult3)
	fmt.Println("err3:", err3)
	fmt.Println(emptyMapResult3)

	err4 := emptyMapResult3.JsonParseSelfInject(jsonResult)
	fmt.Println("err4:", err4)
	fmt.Println("emptyMapResult3:", emptyMapResult3)

	err5 := emptyMapResult3.JsonParseSelfInject(nil)
	fmt.Println("err5:", err5)
	fmt.Println("emptyMapResult3:", emptyMapResult3)

	err6 := emptyMapResult.JsonParseSelfInject(jsonResult)
	fmt.Println("nil emptyMapResult err6:", err6)

	jsonResult = nil
	err7 := jsonResult.Unmarshal(&emptyMapResult3)
	fmt.Println("json Result nil, err7:", err7)

	jsonResult = nil
	err8 := jsonResult.Unmarshal(emptyMapResult)
	fmt.Println("json Result, object nil, err8:", err8)
}
