package main

import (
	"errors"
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corejson"
)

func jsonResultPrettyTest01() {
	mapAnyItems := getMapAnyItems()
	jsonResult := mapAnyItems.JsonPtr()

	fmt.Println(jsonResult.PrettyJsonStringWithErr())

	var rs2 *corejson.Result

	fmt.Println(rs2.PrettyJsonStringWithErr())

	rs2 = corejson.NewPtr([]byte{}, errors.New("something wrong"), "t1")

	fmt.Println(rs2.PrettyJsonStringWithErr())

	rs2 = corejson.NewPtr(nil, errors.New("something wrong"), "t1")

	fmt.Println(rs2.PrettyJsonStringWithErr())
}
