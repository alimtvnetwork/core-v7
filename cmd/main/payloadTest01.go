package main

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corepayload"
	"gitlab.com/evatix-go/core/errcore"
)

func payloadTest01() {
	line := []byte("some payload")

	payload := corepayload.New.PayloadWrapper.UsingCreateInstruction(
		&corepayload.PayloadCreateInstruction{
			Name:           "name -- as payload",
			Identifier:     "id",
			TaskTypeName:   "task type",
			EntityType:     "entity",
			CategoryName:   "category",
			HasManyRecords: false,
			Payloads:       &line,
			Attributes: &corepayload.Attributes{
				ErrorMessage: "some err",
			},
		})

	jsResult := payload.JsonPtr()
	println(jsResult.PrettyJsonString())

	payload2, err := corepayload.New.PayloadWrapper.DeserializeUsingJsonResult(jsResult)
	errcore.HandleErr(err)

	println(payload2.JsonPtr().PrettyJsonString())
	println(payload2.BytesConverter().SafeCastString())

	payload3 := corepayload.New.PayloadWrapper.Create(
		"name3",
		"id3", "taskname3", "category3", jsResult.Bytes)

	println(payload3.JsonPtr().PrettyJsonString())
	println(payload3.DeserializePayloadsToPayloadWrapperMust().JsonPtr().PrettyJsonString())
	pay4, err := payload3.ClonePtr(true)
	errcore.HandleErr(err)
	pay4.Name = "pay 4"
	pay4.Attributes.AddOrUpdateAnyItem(
		"some key",
		payload3.DeserializePayloadsToPayloadWrapperMust().JsonPtr())
	pay4Json := pay4.JsonPtr()
	pay5 := corepayload.PayloadWrapper{}
	pay4Json.DeserializeMust(&pay5)
	println("Pay 5", pay5.PrettyJsonString())

	newJson := corejson.Result{}
	pay5.Attributes.AnyKeyValuePairs.DeserializeMust(
		"some key",
		&newJson)

	println("conv JSON", newJson.PrettyJsonString())
	println(payload3.PrettyJsonString())
}
