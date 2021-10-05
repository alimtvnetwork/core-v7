package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/coredynamic"
)

func jsonResultUnmarshallingTest01() {
	mapAnyItems := getMapAnyItems()
	jsonResult := mapAnyItems.JsonPtr()
	emptyMapResult := coredynamic.EmptyMapAnyItems()

	err := jsonResult.Unmarshal(emptyMapResult)
	fmt.Println("err:", err)
	fmt.Println(emptyMapResult)
}
