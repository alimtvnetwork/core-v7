package main

import (
	"errors"
	"fmt"

	"gitlab.com/auk-go/core/coredata/corejson"
)

func jsonResultPrettyTest01() {
	mapAnyItems := getMapAnyItems()
	jsonResult := mapAnyItems.JsonPtr()

	fmt.Println(jsonResult.PrettyJsonStringOrErrString())

	var rs2 *corejson.Result

	fmt.Println(rs2.PrettyJsonStringOrErrString())

	rs2 = corejson.NewResult.Ptr([]byte{}, errors.New("something wrong"), "t1")

	fmt.Println(rs2.PrettyJsonStringOrErrString())

	rs2 = corejson.NewResult.Ptr(nil, errors.New("something wrong"), "t1")

	fmt.Println(rs2.PrettyJsonStringOrErrString())
}
