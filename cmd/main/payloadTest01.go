package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corepayload"
	"gitlab.com/evatix-go/core/errcore"
)

func payloadTest01() {
	line := []byte("some payload")
	sysUser := corepayload.New.User.System(
		"sys-1",
		"system-user-type")

	regularUser := corepayload.New.User.All(
		false,
		"some user id",
		"regular-2",
		"system-user-type",
		"authToken",
		"passhash")

	authInfo := corepayload.AuthInfo{
		Identifier:   "authid",
		ActionType:   "actionType",
		ResourceName: "resourceIdentity",
		SessionInfo: &corepayload.SessionInfo{
			Id:          "session id",
			User:        regularUser,
			SessionPath: "sesssion path",
		},
		UserInfo: &corepayload.UserInfo{
			User:       regularUser,
			SystemUser: sysUser,
		},
	}

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
				AuthInfo:     authInfo.Ptr(),
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
	sliceOfPayloads := []*corepayload.PayloadWrapper{
		payload,
		payload2,
		payload3,
		pay4,
	}

	jsonSlice := corejson.Serialize.Apply(sliceOfPayloads)
	jsonSlice.HandleError()

	newSlice, err := corepayload.New.PayloadWrapper.DeserializeToMany(jsonSlice.Bytes)
	errcore.HandleErr(err)

	fmt.Println("new slice", newSlice)

	payloadsSlice2 := corepayload.New.PayloadsCollection.UsingWrappers(
		sliceOfPayloads...)
	fmt.Println("new slice2", payloadsSlice2.PrettyJsonString())
}
