package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Attributes Getters ──

func Test_Attributes_IsNull(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes

	// Act
	actual := args.Map{"nil": attr.IsNull()}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Attributes.IsNull returns correct value -- with args", actual)
}

func Test_Attributes_HasSafeItems(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))

	// Act
	actual := args.Map{"safe": attr.HasSafeItems()}

	// Assert
	expected := args.Map{"safe": true}
	expected.ShouldBeEqual(t, 0, "Attributes.HasSafeItems returns correct value -- with args", actual)
}

func Test_Attributes_HasStringKey(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"}))

	// Act
	actual := args.Map{
		"has": attr.HasStringKey("k1"),
		"notHas": !attr.HasStringKey("k2"),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"notHas": true,
	}
	expected.ShouldBeEqual(t, 0, "HasStringKey returns correct value -- with args", actual)
}

func Test_Attributes_HasAnyKey(t *testing.T) {
	// Arrange
	anyMap := coredynamic.NewMapAnyItems(0)
	anyMap.Add("x", 42)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)

	// Act
	actual := args.Map{
		"has": attr.HasAnyKey("x"),
		"notHas": !attr.HasAnyKey("y"),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"notHas": true,
	}
	expected.ShouldBeEqual(t, 0, "HasAnyKey returns correct value -- with args", actual)
}

func Test_Attributes_Payloads_Empty(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	result := attr.Payloads()

	// Act
	actual := args.Map{"emptyBytes": len(result) == 0}

	// Assert
	expected := args.Map{"emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "Payloads returns nil -- nil", actual)
}

func Test_Attributes_PayloadsString(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("hello"))

	// Act
	actual := args.Map{"str": attr.PayloadsString()}

	// Assert
	expected := args.Map{"str": "hello"}
	expected.ShouldBeEqual(t, 0, "PayloadsString returns correct value -- with args", actual)
}

func Test_Attributes_PayloadsString_Empty(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"empty": attr.PayloadsString() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "PayloadsString returns empty -- empty", actual)
}

func Test_Attributes_AnyKeyValMap(t *testing.T) {
	// Arrange
	anyMap := coredynamic.NewMapAnyItems(0)
	anyMap.Add("k", "v")
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	result := attr.AnyKeyValMap()

	// Act
	actual := args.Map{"hasKey": result["k"] == "v"}

	// Assert
	expected := args.Map{"hasKey": true}
	expected.ShouldBeEqual(t, 0, "AnyKeyValMap returns correct value -- with args", actual)
}

func Test_Attributes_Hashmap(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "b"}))
	result := attr.Hashmap()

	// Act
	actual := args.Map{"val": result["a"]}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- with args", actual)
}

func Test_Attributes_Length(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("abc"))

	// Act
	actual := args.Map{
		"len": attr.Length(),
		"count": attr.Count(),
		"cap": attr.Capacity(),
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"count": 3,
		"cap": 3,
	}
	expected.ShouldBeEqual(t, 0, "Length/Count/Capacity returns correct value -- with args", actual)
}

func Test_Attributes_NilLength(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes

	// Act
	actual := args.Map{
		"len": attr.Length(),
		"dynLen": attr.DynamicBytesLength(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"dynLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- Length", actual)
}

func Test_Attributes_HasPagingInfo(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"has": attr.HasPagingInfo()}

	// Assert
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "HasPagingInfo returns correct value -- with args", actual)
}

func Test_Attributes_HasFromTo(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"has": attr.HasFromTo()}

	// Assert
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "HasFromTo returns correct value -- with args", actual)
}

func Test_Attributes_IsValid(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{
		"valid": attr.IsValid(),
		"invalid": attr.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"invalid": true,
	} // IsInvalid includes HasIssuesOrEmpty
	expected.ShouldBeEqual(t, 0, "IsValid/IsInvalid returns error -- with args", actual)
}

func Test_Attributes_ErrorMethods(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{
		"hasError":   attr.HasError(),
		"emptyErr":   attr.IsEmptyError(),
		"compiledNil": attr.CompiledError() == nil,
		"errNil":     attr.Error() == nil,
	}

	// Assert
	expected := args.Map{
		"hasError": false, "emptyErr": true, "compiledNil": true, "errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Error returns error -- methods", actual)
}

func Test_Attributes_StringKeyValuePairsLength(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "b"}))

	// Act
	actual := args.Map{"len": attr.StringKeyValuePairsLength()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringKeyValuePairsLength returns correct value -- with args", actual)
}

func Test_Attributes_AnyKeyValuePairsLength(t *testing.T) {
	// Arrange
	anyMap := coredynamic.NewMapAnyItems(0)
	anyMap.Add("k", 1)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)

	// Act
	actual := args.Map{"len": attr.AnyKeyValuePairsLength()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyKeyValuePairsLength returns correct value -- with args", actual)
}

func Test_Attributes_NilAnyKeyValuePairsLength(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes

	// Act
	actual := args.Map{
		"len": attr.AnyKeyValuePairsLength(),
		"strLen": attr.StringKeyValuePairsLength(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"strLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- AnyKeyValuePairsLength", actual)
}

func Test_Attributes_IsEmpty(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{
		"empty": attr.IsEmpty(),
		"hasItems": attr.HasItems(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"hasItems": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmpty/HasItems returns empty -- with args", actual)
}

func Test_Attributes_IsPagingInfoEmpty(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"empty": attr.IsPagingInfoEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "IsPagingInfoEmpty returns empty -- with args", actual)
}

func Test_Attributes_IsKeyValuePairsEmpty(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"empty": attr.IsKeyValuePairsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "IsKeyValuePairsEmpty returns empty -- with args", actual)
}

func Test_Attributes_IsAnyKeyValuePairsEmpty(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"empty": attr.IsAnyKeyValuePairsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "IsAnyKeyValuePairsEmpty returns empty -- with args", actual)
}

func Test_Attributes_UserInfo(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{
		"userEmpty":    attr.IsUserInfoEmpty(),
		"authEmpty":    attr.IsAuthInfoEmpty(),
		"sessionEmpty": attr.IsSessionInfoEmpty(),
		"virtualNil":   attr.VirtualUser() == nil,
		"systemNil":    attr.SystemUser() == nil,
		"sessionNil":   attr.SessionUser() == nil,
	}

	// Assert
	expected := args.Map{
		"userEmpty": true, "authEmpty": true, "sessionEmpty": true,
		"virtualNil": true, "systemNil": true, "sessionNil": true,
	}
	expected.ShouldBeEqual(t, 0, "UserInfo returns correct value -- methods", actual)
}

func Test_Attributes_HasUserInfo_WithData(t *testing.T) {
	// Arrange
	user := &corepayload.User{Name: "test"}
	userInfo := &corepayload.UserInfo{User: user}
	authInfo := &corepayload.AuthInfo{UserInfo: userInfo}
	attr := corepayload.New.Attributes.UsingAuthInfo(authInfo)

	// Act
	actual := args.Map{
		"hasUser":    attr.HasUserInfo(),
		"hasAuth":    attr.HasAuthInfo(),
		"virtualName": attr.VirtualUser().Name,
	}

	// Assert
	expected := args.Map{
		"hasUser": true, "hasAuth": true, "virtualName": "test",
	}
	expected.ShouldBeEqual(t, 0, "HasUserInfo returns non-empty -- with data", actual)
}

func Test_Attributes_SessionInfo_WithData(t *testing.T) {
	// Arrange
	sessionUser := &corepayload.User{Name: "session-user"}
	session := &corepayload.SessionInfo{Id: "s1", User: sessionUser}
	authInfo := &corepayload.AuthInfo{SessionInfo: session}
	attr := corepayload.New.Attributes.UsingAuthInfo(authInfo)

	// Act
	actual := args.Map{
		"hasSession":  attr.HasSessionInfo(),
		"sessionNotNil": attr.SessionInfo() != nil,
		"sessionUser":   attr.SessionUser().Name,
	}

	// Assert
	expected := args.Map{
		"hasSession": true, "sessionNotNil": true, "sessionUser": "session-user",
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns non-empty -- with data", actual)
}

func Test_Attributes_AuthType_ResourceName(t *testing.T) {
	// Arrange
	authInfo := &corepayload.AuthInfo{ActionType: "login", ResourceName: "/api"}
	attr := corepayload.New.Attributes.UsingAuthInfo(authInfo)

	// Act
	actual := args.Map{
		"authType": attr.AuthType(),
		"resource": attr.ResourceName(),
	}

	// Assert
	expected := args.Map{
		"authType": "login",
		"resource": "/api",
	}
	expected.ShouldBeEqual(t, 0, "AuthType/ResourceName returns correct value -- with args", actual)
}

func Test_Attributes_AuthType_Empty(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{
		"authType": attr.AuthType(),
		"resource": attr.ResourceName(),
	}

	// Assert
	expected := args.Map{
		"authType": "",
		"resource": "",
	}
	expected.ShouldBeEqual(t, 0, "AuthType/ResourceName returns empty -- empty", actual)
}

func Test_Attributes_GetStringKeyValue_FromAttributesIsNullV2(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "key", Value: "val"}))
	val, found := attr.GetStringKeyValue("key")

	// Act
	actual := args.Map{
		"val": val,
		"found": found,
	}

	// Assert
	expected := args.Map{
		"val": "val",
		"found": true,
	}
	expected.ShouldBeEqual(t, 0, "GetStringKeyValue returns correct value -- with args", actual)
}

func Test_Attributes_GetStringKeyValue_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	val, found := attr.GetStringKeyValue("key")

	// Act
	actual := args.Map{
		"val": val,
		"found": found,
	}

	// Assert
	expected := args.Map{
		"val": "",
		"found": false,
	}
	expected.ShouldBeEqual(t, 0, "GetStringKeyValue returns nil -- nil", actual)
}

func Test_Attributes_GetAnyKeyValue_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	val, found := attr.GetAnyKeyValue("key")

	// Act
	actual := args.Map{
		"nil": val == nil,
		"found": found,
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"found": false,
	}
	expected.ShouldBeEqual(t, 0, "GetAnyKeyValue returns nil -- nil", actual)
}

func Test_Attributes_HasStringKeyValuePairs(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "b"}))

	// Act
	actual := args.Map{"has": attr.HasStringKeyValuePairs()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasStringKeyValuePairs returns correct value -- with args", actual)
}

func Test_Attributes_HasDynamicPayloads(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	// Act
	actual := args.Map{"has": attr.HasDynamicPayloads()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasDynamicPayloads returns correct value -- with args", actual)
}

// ── Attributes Setters ──

func Test_Attributes_SetAuthInfo_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	result := attr.SetAuthInfo(&corepayload.AuthInfo{ActionType: "test"})

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetAuthInfo returns nil -- nil receiver", actual)
}

func Test_Attributes_SetUserInfo_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	result := attr.SetUserInfo(&corepayload.UserInfo{})

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetUserInfo returns nil -- nil receiver", actual)
}

func Test_Attributes_AddNewStringKeyValueOnly(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "b"}))
	added := attr.AddNewStringKeyValueOnly("c", "d")

	// Act
	actual := args.Map{"added": added}

	// Assert
	expected := args.Map{"added": true}
	expected.ShouldBeEqual(t, 0, "AddNewStringKeyValueOnly returns correct value -- with args", actual)
}

func Test_Attributes_AddNewStringKeyValueOnly_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	added := attr.AddNewStringKeyValueOnly("c", "d")

	// Act
	actual := args.Map{"added": added}

	// Assert
	expected := args.Map{"added": false}
	expected.ShouldBeEqual(t, 0, "AddNewStringKeyValueOnly returns nil -- nil", actual)
}

func Test_Attributes_AddNewAnyKeyValueOnly(t *testing.T) {
	// Arrange
	anyMap := coredynamic.NewMapAnyItems(0)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	added := attr.AddNewAnyKeyValueOnly("k", 42)

	// Act
	actual := args.Map{"added": added}

	// Assert
	expected := args.Map{"added": true}
	expected.ShouldBeEqual(t, 0, "AddNewAnyKeyValueOnly returns correct value -- with args", actual)
}

func Test_Attributes_AddNewAnyKeyValueOnly_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	added := attr.AddNewAnyKeyValueOnly("k", 42)

	// Act
	actual := args.Map{"added": added}

	// Assert
	expected := args.Map{"added": false}
	expected.ShouldBeEqual(t, 0, "AddNewAnyKeyValueOnly returns nil -- nil", actual)
}

func Test_Attributes_AddOrUpdateString(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "b"}))
	isNew := attr.AddOrUpdateString("c", "d")

	// Act
	actual := args.Map{"isNew": isNew}

	// Assert
	expected := args.Map{"isNew": true}
	expected.ShouldBeEqual(t, 0, "AddOrUpdateString returns correct value -- with args", actual)
}

func Test_Attributes_AddOrUpdateString_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	isNew := attr.AddOrUpdateString("a", "b")

	// Act
	actual := args.Map{"isNew": isNew}

	// Assert
	expected := args.Map{"isNew": false}
	expected.ShouldBeEqual(t, 0, "AddOrUpdateString returns nil -- nil", actual)
}

func Test_Attributes_AddOrUpdateAnyItem(t *testing.T) {
	// Arrange
	anyMap := coredynamic.NewMapAnyItems(0)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	isNew := attr.AddOrUpdateAnyItem("k", 1)

	// Act
	actual := args.Map{"isNew": isNew}

	// Assert
	expected := args.Map{"isNew": true}
	expected.ShouldBeEqual(t, 0, "AddOrUpdateAnyItem returns correct value -- with args", actual)
}

func Test_Attributes_AddOrUpdateAnyItem_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	isNew := attr.AddOrUpdateAnyItem("k", 1)

	// Act
	actual := args.Map{"isNew": isNew}

	// Assert
	expected := args.Map{"isNew": false}
	expected.ShouldBeEqual(t, 0, "AddOrUpdateAnyItem returns nil -- nil", actual)
}

func Test_Attributes_SetBasicErr_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	result := attr.SetBasicErr(nil)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetBasicErr returns nil -- nil receiver", actual)
}

func Test_Attributes_Clear(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "b"}))
	attr.Clear()

	// Act
	actual := args.Map{"empty": attr.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Clear returns correct value -- with args", actual)
}

func Test_Attributes_Clear_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	attr.Clear() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Clear returns nil -- nil", actual)
}

func Test_Attributes_Dispose(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	attr.Dispose()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual)
}

func Test_Attributes_Dispose_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	attr.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns nil -- nil", actual)
}

func Test_Attributes_HandleErr_NoError_FromAttributesIsNullV2(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	attr.HandleErr()
	attr.HandleError()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr/HandleError returns empty -- no error", actual)
}

func Test_Attributes_MustBeEmptyError_FromAttributesIsNullV2(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	attr.MustBeEmptyError() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmptyError returns empty -- with args", actual)
}

// ── Attributes JSON ──

func Test_Attributes_JsonString(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	result := attr.JsonString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct value -- with args", actual)
}

func Test_Attributes_PrettyJsonString(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	result := attr.PrettyJsonString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString returns correct value -- with args", actual)
}

func Test_Attributes_PayloadsPrettyString(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	result := attr.PayloadsPrettyString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PayloadsPrettyString returns correct value -- with args", actual)
}

func Test_Attributes_PayloadsJsonResult(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	result := attr.PayloadsJsonResult()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PayloadsJsonResult returns correct value -- with args", actual)
}

func Test_Attributes_PayloadsJsonResult_Empty(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	result := attr.PayloadsJsonResult()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PayloadsJsonResult returns empty -- empty", actual)
}

func Test_Attributes_NonPtr(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	nonPtr := attr.NonPtr()

	// Act
	actual := args.Map{
		"ok": true,
		"type": nonPtr.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"type": true,
	}
	expected.ShouldBeEqual(t, 0, "NonPtr returns correct value -- with args", actual)
}

func Test_Attributes_AsAttributesBinder(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	binder := attr.NonPtr().AsAttributesBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsAttributesBinder returns correct value -- with args", actual)
}

func Test_Attributes_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	binder := attr.NonPtr().AsJsonContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

func Test_Attributes_JsonModel(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	model := attr.NonPtr().JsonModel()

	// Act
	actual := args.Map{"empty": model.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "JsonModel returns correct value -- with args", actual)
}

func Test_Attributes_JsonModelAny(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	any := attr.NonPtr().JsonModelAny()

	// Act
	actual := args.Map{"notNil": any != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonModelAny returns correct value -- with args", actual)
}

func Test_Attributes_Clone_Shallow_FromAttributesIsNullV2(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	cloned, err := attr.Clone(false)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": !cloned.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- shallow", actual)
}

func Test_Attributes_Clone_Deep(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	cloned, err := attr.Clone(true)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": !cloned.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- deep", actual)
}

func Test_Attributes_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	cloned, err := attr.ClonePtr(false)

	// Act
	actual := args.Map{
		"nil": cloned == nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_Attributes_IsEqual_FromAttributesIsNullV2(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("x"))
	b := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("x"))

	// Act
	actual := args.Map{"equal": a.IsEqual(b)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- with args", actual)
}

func Test_Attributes_IsEqual_Different(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("x"))
	b := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("y"))

	// Act
	actual := args.Map{"equal": a.IsEqual(b)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- different", actual)
}

func Test_Attributes_IsEqual_BothNil_FromAttributesIsNullV2(t *testing.T) {
	// Arrange
	var a, b *corepayload.Attributes

	// Act
	actual := args.Map{"equal": a.IsEqual(b)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns nil -- both nil", actual)
}

func Test_Attributes_IsEqual_OneNil_FromAttributesIsNullV2(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"equal": a.IsEqual(nil)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns nil -- one nil", actual)
}

func Test_Attributes_DeserializeDynamicPayloads(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"v":1}`))
	var result map[string]int
	err := attr.DeserializeDynamicPayloads(&result)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"v": result["v"],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"v": 1,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeDynamicPayloads returns correct value -- with args", actual)
}

func Test_Attributes_DeserializeDynamicPayloadsToAttributes(t *testing.T) {
	// Arrange
	inner := corepayload.New.Attributes.Empty()
	jsonBytes := []byte(inner.NonPtr().JsonString())
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(jsonBytes)
	result, err := attr.DeserializeDynamicPayloadsToAttributes()

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"errOrNil": err == nil || err != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"errOrNil": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeDynamicPayloadsToAttributes returns correct value -- with args", actual)
}

func Test_Attributes_DynamicPayloadsDeserialize_Nil(t *testing.T) {
	// Arrange
	var attr *corepayload.Attributes
	var result map[string]int
	err := attr.DynamicPayloadsDeserialize(&result)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicPayloadsDeserialize returns nil -- nil", actual)
}

func Test_Attributes_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	jsonResult := corejson.NewPtr(attr.NonPtr())
	result, err := attr.ParseInjectUsingJson(jsonResult)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns correct value -- with args", actual)
}

func Test_Attributes_JsonParseSelfInject(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	jsonResult := corejson.NewPtr(attr.NonPtr())
	err := attr.JsonParseSelfInject(jsonResult)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject returns correct value -- with args", actual)
}

// ── AuthInfo ──

func Test_AuthInfo_IdentifierInteger(t *testing.T) {
	// Arrange
	info := corepayload.AuthInfo{Identifier: "42"}

	// Act
	actual := args.Map{"id": info.IdentifierInteger()}

	// Assert
	expected := args.Map{"id": 42}
	expected.ShouldBeEqual(t, 0, "AuthInfo.IdentifierInteger returns correct value -- with args", actual)
}

func Test_AuthInfo_IdentifierInteger_Empty(t *testing.T) {
	// Arrange
	info := corepayload.AuthInfo{}

	// Act
	actual := args.Map{"id": info.IdentifierInteger()}

	// Assert
	expected := args.Map{"id": -1}
	expected.ShouldBeEqual(t, 0, "AuthInfo.IdentifierInteger returns empty -- empty", actual)
}

func Test_AuthInfo_IdentifierUnsignedInteger_Negative(t *testing.T) {
	// Arrange
	info := corepayload.AuthInfo{}

	// Act
	actual := args.Map{"id": info.IdentifierUnsignedInteger()}

	// Assert
	expected := args.Map{"id": uint(0)}
	expected.ShouldBeEqual(t, 0, "AuthInfo.IdentifierUnsignedInteger returns correct value -- negative", actual)
}

func Test_AuthInfo_Methods(t *testing.T) {
	// Arrange
	info := &corepayload.AuthInfo{ActionType: "login", ResourceName: "/api"}

	// Act
	actual := args.Map{
		"hasAction":   info.HasActionType(),
		"hasResource": info.HasResourceName(),
		"isValid":     info.IsValid(),
		"notEmpty":    info.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"hasAction": true, "hasResource": true, "isValid": true, "notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- methods", actual)
}

func Test_AuthInfo_SetMethods(t *testing.T) {
	// Arrange
	var info *corepayload.AuthInfo
	r1 := info.SetActionType("act")
	r2 := info.SetResourceName("res")
	r3 := info.SetIdentifier("id1")
	r4 := info.SetSessionInfo(&corepayload.SessionInfo{Id: "s1"})
	r5 := info.SetUser(&corepayload.User{Name: "u1"})
	r6 := info.SetSystemUser(&corepayload.User{Name: "sys"})
	r7 := info.SetUserSystemUser(&corepayload.User{Name: "u2"}, &corepayload.User{Name: "sys2"})
	r8 := info.SetUserInfo(&corepayload.UserInfo{})

	// Act
	actual := args.Map{
		"r1": r1 != nil, "r2": r2 != nil, "r3": r3 != nil, "r4": r4 != nil,
		"r5": r5 != nil, "r6": r6 != nil, "r7": r7 != nil, "r8": r8 != nil,
	}

	// Assert
	expected := args.Map{
		"r1": true, "r2": true, "r3": true, "r4": true,
		"r5": true, "r6": true, "r7": true, "r8": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns nil -- set methods nil receiver", actual)
}

func Test_AuthInfo_Clone_FromAttributesIsNullV2(t *testing.T) {
	// Arrange
	info := &corepayload.AuthInfo{ActionType: "login"}
	cloned := info.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": cloned != nil,
		"action": cloned.ActionType,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"action": "login",
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo.ClonePtr returns correct value -- with args", actual)
}

func Test_AuthInfo_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var info *corepayload.AuthInfo
	cloned := info.ClonePtr()

	// Act
	actual := args.Map{"nil": cloned == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "AuthInfo.ClonePtr returns nil -- nil", actual)
}

// ── SessionInfo ──

func Test_SessionInfo_Methods(t *testing.T) {
	// Arrange
	info := &corepayload.SessionInfo{Id: "s1", User: &corepayload.User{Name: "u1"}, SessionPath: "/path"}

	// Act
	actual := args.Map{
		"isEmpty":   info.IsEmpty(),
		"isValid":   info.IsValid(),
		"hasUser":   info.HasUser(),
		"userEmpty": info.IsUserEmpty(),
		"nameEqual": info.IsUsernameEqual("u1"),
	}

	// Assert
	expected := args.Map{
		"isEmpty": false, "isValid": true, "hasUser": true,
		"userEmpty": false, "nameEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns correct value -- methods", actual)
}

func Test_SessionInfo_IdentifierInteger(t *testing.T) {
	// Arrange
	info := corepayload.SessionInfo{Id: "10"}

	// Act
	actual := args.Map{"id": info.IdentifierInteger()}

	// Assert
	expected := args.Map{"id": 10}
	expected.ShouldBeEqual(t, 0, "SessionInfo.IdentifierInteger returns correct value -- with args", actual)
}

func Test_SessionInfo_IdentifierUnsignedInteger(t *testing.T) {
	// Arrange
	info := corepayload.SessionInfo{}

	// Act
	actual := args.Map{"id": info.IdentifierUnsignedInteger()}

	// Assert
	expected := args.Map{"id": uint(0)}
	expected.ShouldBeEqual(t, 0, "SessionInfo.IdentifierUnsignedInteger returns correct value -- with args", actual)
}

func Test_SessionInfo_Clone(t *testing.T) {
	// Arrange
	info := &corepayload.SessionInfo{Id: "s1"}
	cloned := info.ClonePtr()

	// Act
	actual := args.Map{"id": cloned.Id}

	// Assert
	expected := args.Map{"id": "s1"}
	expected.ShouldBeEqual(t, 0, "SessionInfo.Clone returns correct value -- with args", actual)
}

func Test_SessionInfo_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var info *corepayload.SessionInfo
	cloned := info.ClonePtr()

	// Act
	actual := args.Map{"nil": cloned == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SessionInfo.ClonePtr returns nil -- nil", actual)
}

// ── User ──

func Test_User_Methods(t *testing.T) {
	// Arrange
	user := &corepayload.User{Name: "test", Type: "admin", AuthToken: "token", PasswordHash: "hash", Identifier: "1"}

	// Act
	actual := args.Map{
		"hasAuth":     user.HasAuthToken(),
		"hasPass":     user.HasPasswordHash(),
		"isValid":     user.IsValidUser(),
		"isVirtual":   user.IsVirtualUser(),
		"hasType":     user.HasType(),
		"nameEqual":   user.IsNameEqual("test"),
		"notSystem":   user.IsNotSystemUser(),
		"idInt":       user.IdentifierInteger(),
	}

	// Assert
	expected := args.Map{
		"hasAuth": true, "hasPass": true, "isValid": true,
		"isVirtual": true, "hasType": true, "nameEqual": true,
		"notSystem": true, "idInt": 1,
	}
	expected.ShouldBeEqual(t, 0, "User returns correct value -- methods", actual)
}

func Test_User_NilMethods(t *testing.T) {
	// Arrange
	var user *corepayload.User

	// Act
	actual := args.Map{
		"authEmpty":  user.IsAuthTokenEmpty(),
		"passEmpty":  user.IsPasswordHashEmpty(),
		"nameEmpty":  user.IsNameEmpty(),
		"typeEmpty":  user.IsTypeEmpty(),
		"isEmpty":    user.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"authEmpty": true, "passEmpty": true, "nameEmpty": true,
		"typeEmpty": true, "isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "User returns nil -- nil methods", actual)
}

func Test_User_IdentifierUnsignedInteger(t *testing.T) {
	// Arrange
	user := corepayload.User{}

	// Act
	actual := args.Map{"id": user.IdentifierUnsignedInteger()}

	// Assert
	expected := args.Map{"id": uint(0)}
	expected.ShouldBeEqual(t, 0, "User.IdentifierUnsignedInteger returns correct value -- with args", actual)
}

func Test_User_Clone(t *testing.T) {
	// Arrange
	user := &corepayload.User{Name: "test"}
	cloned := user.ClonePtr()

	// Act
	actual := args.Map{"name": cloned.Name}

	// Assert
	expected := args.Map{"name": "test"}
	expected.ShouldBeEqual(t, 0, "User.Clone returns correct value -- with args", actual)
}

func Test_User_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var user *corepayload.User
	cloned := user.ClonePtr()

	// Act
	actual := args.Map{"nil": cloned == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "User.ClonePtr returns nil -- nil", actual)
}

// ── UserInfo ──

func Test_UserInfo_Methods(t *testing.T) {
	// Arrange
	info := &corepayload.UserInfo{
		User:       &corepayload.User{Name: "u"},
		SystemUser: &corepayload.User{Name: "sys"},
	}

	// Act
	actual := args.Map{
		"hasUser":   info.HasUser(),
		"hasSys":    info.HasSystemUser(),
		"isEmpty":   info.IsEmpty(),
		"userEmpty": info.IsUserEmpty(),
		"sysEmpty":  info.IsSystemUserEmpty(),
	}

	// Assert
	expected := args.Map{
		"hasUser": true, "hasSys": true, "isEmpty": false,
		"userEmpty": false, "sysEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "UserInfo returns correct value -- methods", actual)
}

func Test_UserInfo_NilSetMethods(t *testing.T) {
	// Arrange
	var info *corepayload.UserInfo
	r1 := info.SetUser(&corepayload.User{Name: "u"})
	var info2 *corepayload.UserInfo
	r2 := info2.SetSystemUser(&corepayload.User{Name: "sys"})
	var info3 *corepayload.UserInfo
	r3 := info3.SetUserSystemUser(&corepayload.User{Name: "u"}, &corepayload.User{Name: "sys"})

	// Act
	actual := args.Map{
		"r1": r1 != nil,
		"r2": r2 != nil,
		"r3": r3 != nil,
	}

	// Assert
	expected := args.Map{
		"r1": true,
		"r2": true,
		"r3": true,
	}
	expected.ShouldBeEqual(t, 0, "UserInfo returns nil -- nil set methods", actual)
}

func Test_UserInfo_ToNonPtr(t *testing.T) {
	// Arrange
	var info *corepayload.UserInfo
	result := info.ToNonPtr()

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "UserInfo.ToNonPtr returns nil -- nil", actual)
}

func Test_UserInfo_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var info *corepayload.UserInfo
	cloned := info.ClonePtr()

	// Act
	actual := args.Map{"nil": cloned == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "UserInfo.ClonePtr returns nil -- nil", actual)
}

// ── PagingInfo ──

func Test_PagingInfo_Methods(t *testing.T) {
	// Arrange
	pi := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 50}

	// Act
	actual := args.Map{
		"isEmpty":      pi.IsEmpty(),
		"hasTotalPages": pi.HasTotalPages(),
		"hasPageIdx":    pi.HasCurrentPageIndex(),
		"hasPerPage":    pi.HasPerPageItems(),
		"hasTotalItems": pi.HasTotalItems(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": false, "hasTotalPages": true, "hasPageIdx": true,
		"hasPerPage": true, "hasTotalItems": true,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns correct value -- methods", actual)
}

func Test_PagingInfo_Invalid(t *testing.T) {
	// Arrange
	var pi *corepayload.PagingInfo

	// Act
	actual := args.Map{
		"invalidTotal":  pi.IsInvalidTotalPages(),
		"invalidPage":   pi.IsInvalidCurrentPageIndex(),
		"invalidPer":    pi.IsInvalidPerPageItems(),
		"invalidItems":  pi.IsInvalidTotalItems(),
	}

	// Assert
	expected := args.Map{
		"invalidTotal": true, "invalidPage": true,
		"invalidPer": true, "invalidItems": true,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns nil -- nil invalid", actual)
}

func Test_PagingInfo_IsEqual(t *testing.T) {
	// Arrange
	a := &corepayload.PagingInfo{TotalPages: 1, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 10}
	b := &corepayload.PagingInfo{TotalPages: 1, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 10}

	// Act
	actual := args.Map{"equal": a.IsEqual(b)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "PagingInfo.IsEqual returns correct value -- with args", actual)
}

func Test_PagingInfo_IsEqual_Different(t *testing.T) {
	// Arrange
	a := &corepayload.PagingInfo{TotalPages: 1}
	b := &corepayload.PagingInfo{TotalPages: 2}

	// Act
	actual := args.Map{"equal": a.IsEqual(b)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "PagingInfo.IsEqual returns correct value -- different", actual)
}

// ── PayloadWrapper ──

func Test_PayloadWrapper_Basic(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()

	// Act
	actual := args.Map{
		"isEmpty":    pw.IsEmpty(),
		"isNull":     pw.IsNull(),
		"hasAnyNil":  pw.HasAnyNil(),
		"hasItems":   pw.HasItems(),
		"hasAnyItem": pw.HasAnyItem(),
		"count":      pw.Count(),
		"length":     pw.Length(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true, "isNull": false, "hasAnyNil": false,
		"hasItems": false, "hasAnyItem": false, "count": 0, "length": 0,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- basic", actual)
}

func Test_PayloadWrapper_NilLength(t *testing.T) {
	// Arrange
	var pw *corepayload.PayloadWrapper

	// Act
	actual := args.Map{"len": pw.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns nil -- nil length", actual)
}

func Test_PayloadWrapper_AllSafe_Nil(t *testing.T) {
	// Arrange
	var pw *corepayload.PayloadWrapper
	id, name, entity, cat, payloads := pw.AllSafe()

	// Act
	actual := args.Map{
		"id": id,
		"name": name,
		"entity": entity,
		"cat": cat,
		"payloadsLen": len(payloads),
	}

	// Assert
	expected := args.Map{
		"id": "",
		"name": "",
		"entity": "",
		"cat": "",
		"payloadsLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "AllSafe returns nil -- nil", actual)
}

func Test_PayloadWrapper_IsName(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test"}

	// Act
	actual := args.Map{
		"is": pw.IsName("test"),
		"not": pw.IsName("other"),
	}

	// Assert
	expected := args.Map{
		"is": true,
		"not": false,
	}
	expected.ShouldBeEqual(t, 0, "IsName returns correct value -- with args", actual)
}

func Test_PayloadWrapper_IsIdentifier(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Identifier: "id1"}

	// Act
	actual := args.Map{"is": pw.IsIdentifier("id1")}

	// Assert
	expected := args.Map{"is": true}
	expected.ShouldBeEqual(t, 0, "IsIdentifier returns correct value -- with args", actual)
}

func Test_PayloadWrapper_IsCategory(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{CategoryName: "cat1"}

	// Act
	actual := args.Map{"is": pw.IsCategory("cat1")}

	// Assert
	expected := args.Map{"is": true}
	expected.ShouldBeEqual(t, 0, "IsCategory returns correct value -- with args", actual)
}

func Test_PayloadWrapper_IsEntityType(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{EntityType: "e1"}

	// Act
	actual := args.Map{"is": pw.IsEntityType("e1")}

	// Assert
	expected := args.Map{"is": true}
	expected.ShouldBeEqual(t, 0, "IsEntityType returns correct value -- with args", actual)
}

func Test_PayloadWrapper_IsTaskTypeName(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{TaskTypeName: "t1"}

	// Act
	actual := args.Map{"is": pw.IsTaskTypeName("t1")}

	// Assert
	expected := args.Map{"is": true}
	expected.ShouldBeEqual(t, 0, "IsTaskTypeName returns correct value -- with args", actual)
}

func Test_PayloadWrapper_IdentifierInteger(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Identifier: "42"}

	// Act
	actual := args.Map{
		"id": pw.IdentifierInteger(),
		"uint": pw.IdentifierUnsignedInteger(),
	}

	// Assert
	expected := args.Map{
		"id": 42,
		"uint": uint(42),
	}
	expected.ShouldBeEqual(t, 0, "IdentifierInteger returns correct value -- with args", actual)
}

func Test_PayloadWrapper_IdentifierInteger_Empty(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}

	// Act
	actual := args.Map{
		"id": pw.IdentifierInteger(),
		"uint": pw.IdentifierUnsignedInteger(),
	}

	// Assert
	expected := args.Map{
		"id": -1,
		"uint": uint(0),
	}
	expected.ShouldBeEqual(t, 0, "IdentifierInteger returns empty -- empty", actual)
}

func Test_PayloadWrapper_IsPayloadsEqual(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Payloads: []byte("data")}

	// Act
	actual := args.Map{"equal": pw.IsPayloadsEqual([]byte("data"))}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsPayloadsEqual returns correct value -- with args", actual)
}

func Test_PayloadWrapper_HasSingleRecord(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{HasManyRecords: false}

	// Act
	actual := args.Map{"single": pw.HasSingleRecord()}

	// Assert
	expected := args.Map{"single": true}
	expected.ShouldBeEqual(t, 0, "HasSingleRecord returns correct value -- with args", actual)
}

func Test_PayloadWrapper_HasAttributes(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()

	// Act
	actual := args.Map{
		"has": pw.HasAttributes(),
		"emptyAttr": pw.IsEmptyAttributes(),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"emptyAttr": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAttributes returns correct value -- with args", actual)
}

func Test_PayloadWrapper_DynamicPayloads(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Payloads: []byte("test")}

	// Act
	actual := args.Map{"len": len(pw.DynamicPayloads())}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "DynamicPayloads returns correct value -- with args", actual)
}

func Test_PayloadWrapper_DynamicPayloads_Nil(t *testing.T) {
	// Arrange
	var pw *corepayload.PayloadWrapper
	result := pw.DynamicPayloads()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicPayloads returns nil -- nil", actual)
}

func Test_PayloadWrapper_SetDynamicPayloads(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()
	err := pw.SetDynamicPayloads([]byte("new"))

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(pw.Payloads),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "SetDynamicPayloads returns correct value -- with args", actual)
}

func Test_PayloadWrapper_Clone_Shallow(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test", Payloads: []byte("data")}
	cloned, err := pw.Clone(false)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": cloned.Name,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "test",
	}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- shallow", actual)
}

func Test_PayloadWrapper_Clone_Deep(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "deep", Payloads: []byte("data")}
	cloned, err := pw.Clone(true)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": cloned.Name,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "deep",
	}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- deep", actual)
}

func Test_PayloadWrapper_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var pw *corepayload.PayloadWrapper
	cloned, err := pw.ClonePtr(false)

	// Act
	actual := args.Map{
		"nil": cloned == nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_PayloadWrapper_NonPtr_Nil(t *testing.T) {
	// Arrange
	var pw *corepayload.PayloadWrapper
	result := pw.NonPtr()

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NonPtr returns nil -- nil", actual)
}

func Test_PayloadWrapper_Dispose_FromAttributesIsNullV2(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Dispose()

	// Act
	actual := args.Map{"nilAttr": pw.Attributes == nil}

	// Assert
	expected := args.Map{"nilAttr": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual)
}

func Test_PayloadWrapper_Dispose_Nil(t *testing.T) {
	// Arrange
	var pw *corepayload.PayloadWrapper
	pw.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns nil -- nil", actual)
}

func Test_PayloadWrapper_Clear_Nil(t *testing.T) {
	// Arrange
	var pw *corepayload.PayloadWrapper
	pw.Clear() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Clear returns nil -- nil", actual)
}

func Test_PayloadWrapper_Username(t *testing.T) {
	// Arrange
	user := &corepayload.User{Name: "myuser"}
	userInfo := &corepayload.UserInfo{User: user}
	authInfo := &corepayload.AuthInfo{UserInfo: userInfo}
	attr := corepayload.New.Attributes.UsingAuthInfo(authInfo)
	pw := &corepayload.PayloadWrapper{Attributes: attr, Payloads: []byte("x")}

	// Act
	actual := args.Map{"name": pw.Username()}

	// Assert
	expected := args.Map{"name": "myuser"}
	expected.ShouldBeEqual(t, 0, "Username returns correct value -- with args", actual)
}

func Test_PayloadWrapper_Username_Empty(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()

	// Act
	actual := args.Map{"name": pw.Username()}

	// Assert
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "Username returns empty -- empty", actual)
}

// ── PayloadsCollection ──

func Test_PayloadsCollection_Filter(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})
	coll.Add(corepayload.PayloadWrapper{Name: "c", Payloads: []byte("z")})

	result := coll.Filter(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return pw.Name == "b", pw.Name == "b"
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Filter returns correct value -- with args", actual)
}

func Test_PayloadsCollection_FilterWithLimit_FromAttributesIsNullV2(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	for i := 0; i < 10; i++ {
		coll.Add(corepayload.PayloadWrapper{Name: "item", Payloads: []byte("x")})
	}

	result := coll.FilterWithLimit(3, func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "FilterWithLimit returns non-empty -- with args", actual)
}

func Test_PayloadsCollection_FirstById(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Identifier: "1", Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Identifier: "2", Name: "b", Payloads: []byte("y")})

	result := coll.FirstById("2")

	// Act
	actual := args.Map{"name": result.Name}

	// Assert
	expected := args.Map{"name": "b"}
	expected.ShouldBeEqual(t, 0, "FirstById returns correct value -- with args", actual)
}

func Test_PayloadsCollection_FirstByCategory(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{CategoryName: "cat1", Payloads: []byte("x")})
	result := coll.FirstByCategory("cat1")

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FirstByCategory returns correct value -- with args", actual)
}

func Test_PayloadsCollection_FirstByTaskType(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{TaskTypeName: "task1", Payloads: []byte("x")})
	result := coll.FirstByTaskType("task1")

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FirstByTaskType returns correct value -- with args", actual)
}

func Test_PayloadsCollection_FirstByEntityType(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{EntityType: "entity1", Payloads: []byte("x")})
	result := coll.FirstByEntityType("entity1")

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FirstByEntityType returns correct value -- with args", actual)
}

func Test_PayloadsCollection_SkipFilterCollection(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})

	result := coll.SkipFilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return pw.Name == "a", false
	})

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SkipFilterCollection returns correct value -- with args", actual)
}

func Test_PayloadsCollection_FilterCollectionByIds(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Identifier: "1", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Identifier: "2", Payloads: []byte("y")})
	coll.Add(corepayload.PayloadWrapper{Identifier: "3", Payloads: []byte("z")})

	result := coll.FilterCollectionByIds("1", "3")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FilterCollectionByIds returns correct value -- with args", actual)
}

func Test_PayloadsCollection_FilterNameCollection(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "target", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "other", Payloads: []byte("y")})

	result := coll.FilterNameCollection("target")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FilterNameCollection returns correct value -- with args", actual)
}

func Test_PayloadsCollection_FilterCategoryCollection(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{CategoryName: "cat1", Payloads: []byte("x")})

	result := coll.FilterCategoryCollection("cat1")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FilterCategoryCollection returns correct value -- with args", actual)
}

func Test_PayloadsCollection_FilterEntityTypeCollection(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{EntityType: "e1", Payloads: []byte("x")})

	result := coll.FilterEntityTypeCollection("e1")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FilterEntityTypeCollection returns correct value -- with args", actual)
}

func Test_PayloadsCollection_FilterTaskTypeCollection(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{TaskTypeName: "t1", Payloads: []byte("x")})

	result := coll.FilterTaskTypeCollection("t1")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FilterTaskTypeCollection returns correct value -- with args", actual)
}

func Test_PayloadsCollection_Paging(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	for i := 0; i < 25; i++ {
		coll.Add(corepayload.PayloadWrapper{Name: "item", Payloads: []byte("x")})
	}

	pages := coll.GetPagesSize(10)

	// Act
	actual := args.Map{"pages": pages}

	// Assert
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "GetPagesSize returns correct value -- with args", actual)
}

func Test_PayloadsCollection_GetPagesSize_Zero(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()

	// Act
	actual := args.Map{"pages": coll.GetPagesSize(0)}

	// Assert
	expected := args.Map{"pages": 0}
	expected.ShouldBeEqual(t, 0, "GetPagesSize returns correct value -- zero", actual)
}

func Test_PayloadsCollection_GetSinglePageCollection(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	for i := 0; i < 25; i++ {
		coll.Add(corepayload.PayloadWrapper{Name: "item", Payloads: []byte("x")})
	}

	page := coll.GetSinglePageCollection(10, 3)

	// Act
	actual := args.Map{"len": page.Length()}

	// Assert
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection returns correct value -- with args", actual)
}

func Test_PayloadsCollection_GetPagedCollection(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	for i := 0; i < 25; i++ {
		coll.Add(corepayload.PayloadWrapper{Name: "item", Payloads: []byte("x")})
	}

	pages := coll.GetPagedCollection(10)

	// Act
	actual := args.Map{"numPages": len(pages)}

	// Assert
	expected := args.Map{"numPages": 3}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection returns correct value -- with args", actual)
}

func Test_PayloadsCollection_JsonMethods(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})

	// Act
	actual := args.Map{
		"jsonStr":   coll.JsonString() != "",
		"str":       coll.String() != "",
		"prettyStr": coll.PrettyJsonString() != "",
	}

	// Assert
	expected := args.Map{
		"jsonStr": true, "str": true, "prettyStr": true,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- json methods", actual)
}

func Test_PayloadsCollection_JsonString_Empty(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()

	// Act
	actual := args.Map{"empty": coll.JsonString() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns empty -- empty", actual)
}

func Test_PayloadsCollection_Strings(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	result := coll.Strings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Strings returns correct value -- with args", actual)
}

func Test_PayloadsCollection_CsvStrings(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	result := coll.CsvStrings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CsvStrings returns correct value -- with args", actual)
}

func Test_PayloadsCollection_CsvStrings_Empty(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	result := coll.CsvStrings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CsvStrings returns empty -- empty", actual)
}

func Test_PayloadsCollection_Join(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	result := coll.Join(", ")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Join returns correct value -- with args", actual)
}

func Test_PayloadsCollection_JoinCsv(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	result := coll.JoinCsv()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinCsv returns correct value -- with args", actual)
}

func Test_PayloadsCollection_JoinCsvLine(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	result := coll.JoinCsvLine()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinCsvLine returns correct value -- with args", actual)
}

func Test_PayloadsCollection_StringsUsingFmt(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	result := coll.StringsUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.Name
	})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "StringsUsingFmt returns correct value -- with args", actual)
}

func Test_PayloadsCollection_JoinUsingFmt(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	coll.Add(corepayload.PayloadWrapper{Name: "b", Payloads: []byte(`"y"`)})
	result := coll.JoinUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.Name
	}, ",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "JoinUsingFmt returns correct value -- with args", actual)
}

func Test_PayloadsCollection_Reverse(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})
	coll.Add(corepayload.PayloadWrapper{Name: "c", Payloads: []byte("z")})
	coll.Reverse()

	// Act
	actual := args.Map{
		"first": coll.First().Name,
		"last": coll.Last().Name,
	}

	// Assert
	expected := args.Map{
		"first": "c",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "Reverse returns correct value -- with args", actual)
}

func Test_PayloadsCollection_Reverse_Two(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})
	coll.Reverse()

	// Act
	actual := args.Map{"first": coll.First().Name}

	// Assert
	expected := args.Map{"first": "b"}
	expected.ShouldBeEqual(t, 0, "Reverse returns correct value -- two", actual)
}

func Test_PayloadsCollection_Reverse_One(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Reverse()

	// Act
	actual := args.Map{"first": coll.First().Name}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "Reverse returns correct value -- one", actual)
}

func Test_PayloadsCollection_InsertAt(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "c", Payloads: []byte("z")})
	coll.InsertAt(1, corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})

	// Act
	actual := args.Map{"len": coll.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "InsertAt returns correct value -- with args", actual)
}

func Test_PayloadsCollection_ConcatNew(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	newColl := coll.ConcatNew(corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})

	// Act
	actual := args.Map{
		"origLen": coll.Length(),
		"newLen": newColl.Length(),
	}

	// Assert
	expected := args.Map{
		"origLen": 1,
		"newLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "ConcatNew returns correct value -- with args", actual)
}

func Test_PayloadsCollection_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var coll *corepayload.PayloadsCollection
	cloned := coll.ClonePtr()

	// Act
	actual := args.Map{"nil": cloned == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_PayloadsCollection_IsEqual(t *testing.T) {
	// Arrange
	a := corepayload.New.PayloadsCollection.Empty()
	a.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	b := a.ClonePtr()

	// Act
	actual := args.Map{"equal": a.IsEqual(b)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- with args", actual)
}

func Test_PayloadsCollection_SafeLimitCollection(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})

	result := coll.SafeLimitCollection(10)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeLimitCollection returns correct value -- with args", actual)
}

func Test_PayloadsCollection_AddsIf_True(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.AddsIf(true, corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})

	// Act
	actual := args.Map{"len": coll.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddsIf returns non-empty -- true", actual)
}

func Test_PayloadsCollection_AddsIf_False(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.AddsIf(false, corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})

	// Act
	actual := args.Map{"len": coll.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddsIf returns non-empty -- false", actual)
}

func Test_PayloadsCollection_AddsPtrOptions_SkipIssued(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.New.PayloadWrapper.Empty() // empty = has issues
	coll.AddsPtrOptions(true, pw)

	// Act
	actual := args.Map{"len": coll.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddsPtrOptions returns correct value -- skip issued", actual)
}

func Test_PayloadsCollection_AddsOptions_SkipIssued(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.PayloadWrapper{} // empty = has issues
	coll.AddsOptions(true, pw)

	// Act
	actual := args.Map{"len": coll.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddsOptions returns correct value -- skip issued", actual)
}

func Test_PayloadsCollection_Dispose(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Dispose()

	// Act
	actual := args.Map{"nilItems": coll.Items == nil}

	// Assert
	expected := args.Map{"nilItems": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual)
}

func Test_PayloadsCollection_Dispose_Nil(t *testing.T) {
	// Arrange
	var coll *corepayload.PayloadsCollection
	coll.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns nil -- nil", actual)
}

func Test_PayloadsCollection_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	jsonResult := coll.JsonPtr()

	newColl := corepayload.New.PayloadsCollection.Empty()
	result, err := newColl.ParseInjectUsingJson(jsonResult)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": result.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns correct value -- with args", actual)
}

func Test_PayloadsCollection_JsonParseSelfInject(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	jsonResult := coll.JsonPtr()

	newColl := corepayload.New.PayloadsCollection.Empty()
	err := newColl.JsonParseSelfInject(jsonResult)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject returns correct value -- with args", actual)
}

func Test_PayloadsCollection_AsJsoner(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	jsoner := coll.AsJsoner()

	// Act
	actual := args.Map{"notNil": jsoner != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsoner returns correct value -- with args", actual)
}

func Test_PayloadsCollection_AsJsonParseSelfInjector(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	injector := coll.AsJsonParseSelfInjector()

	// Act
	actual := args.Map{"notNil": injector != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonParseSelfInjector returns correct value -- with args", actual)
}

func Test_PayloadsCollection_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	binder := coll.AsJsonContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

// ── newPayloadsCollectionCreator ──

func Test_NewPayloadsCollection_Deserialize(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	jsonBytes, _ := corejson.Serialize.Raw(coll)

	result, err := corepayload.New.PayloadsCollection.Deserialize(jsonBytes)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": result.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- with args", actual)
}

func Test_NewPayloadsCollection_DeserializeToMany(t *testing.T) {
	// Arrange
	_, err := corepayload.New.PayloadsCollection.DeserializeToMany([]byte("invalid"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializeToMany returns error -- invalid", actual)
}

func Test_NewPayloadsCollection_DeserializeUsingJsonResult(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	jsonResult := coll.JsonPtr()
	result, err := corepayload.New.PayloadsCollection.DeserializeUsingJsonResult(jsonResult)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeUsingJsonResult returns correct value -- with args", actual)
}

// ── emptyCreator ──

func Test_EmptyCreator_FromAttributesIsNullV2(t *testing.T) {
	// Arrange
	attr := corepayload.Empty.Attributes()
	attrDef := corepayload.Empty.AttributesDefaults()
	pw := corepayload.Empty.PayloadWrapper()
	coll := corepayload.Empty.PayloadsCollection()

	// Act
	actual := args.Map{
		"attrNotNil":    attr != nil,
		"attrDefNotNil": attrDef != nil,
		"pwNotNil":      pw != nil,
		"collNotNil":    coll != nil,
	}

	// Assert
	expected := args.Map{
		"attrNotNil": true, "attrDefNotNil": true,
		"pwNotNil": true, "collNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "emptyCreator returns empty -- with args", actual)
}

// ── newAttributesCreator ──

func Test_NewAttributes_Creators(t *testing.T) {
	// Arrange
	r1 := corepayload.New.Attributes.Create(nil, nil, []byte("data"))
	r2 := corepayload.New.Attributes.ErrFromTo(nil, nil, []byte("data"))
	r3 := corepayload.New.Attributes.UsingAuthInfoDynamicBytes(nil, []byte("data"))
	r4 := corepayload.New.Attributes.UsingKeyValuesPlusDynamic(nil, []byte("data"))
	r5 := corepayload.New.Attributes.UsingAuthInfoKeyValues(nil, nil)
	r6 := corepayload.New.Attributes.UsingAuthInfoAnyKeyValues(nil, nil)
	r7 := corepayload.New.Attributes.UsingAnyKeyValuesPlusDynamic(nil, []byte("data"))
	r8 := corepayload.New.Attributes.UsingBasicError(nil)

	// Act
	actual := args.Map{
		"r1": r1 != nil, "r2": r2 != nil, "r3": r3 != nil, "r4": r4 != nil,
		"r5": r5 != nil, "r6": r6 != nil, "r7": r7 != nil, "r8": r8 != nil,
	}

	// Assert
	expected := args.Map{
		"r1": true, "r2": true, "r3": true, "r4": true,
		"r5": true, "r6": true, "r7": true, "r8": true,
	}
	expected.ShouldBeEqual(t, 0, "newAttributesCreator returns correct value -- methods", actual)
}

func Test_NewAttributes_Deserialize(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	jsonBytes, _ := corejson.Serialize.Raw(attr)
	result, err := corepayload.New.Attributes.Deserialize(jsonBytes)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Attributes.Deserialize returns correct value -- with args", actual)
}

func Test_NewAttributes_DeserializeMany(t *testing.T) {
	// Arrange
	_, err := corepayload.New.Attributes.DeserializeMany([]byte("invalid"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Attributes.DeserializeMany returns error -- invalid", actual)
}

func Test_NewAttributes_CastOrDeserializeFrom_Nil(t *testing.T) {
	// Arrange
	_, err := corepayload.New.Attributes.CastOrDeserializeFrom(nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CastOrDeserializeFrom returns nil -- nil", actual)
}

// ── payloadProperties ──

func Test_PayloadProperties(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.All(
		"name", "id1", "task", "cat", "entity", false,
		corepayload.New.Attributes.Empty(), []byte("data"))
	props := pw.PayloadProperties()

	// Act
	actual := args.Map{
		"name":     props.Name(),
		"id":       props.IdString(),
		"idInt":    props.IdInteger(),
		"category": props.Category(),
		"entity":   props.EntityType(),
		"hasSingle": props.HasSingleRecordOnly(),
	}

	// Assert
	expected := args.Map{
		"name": "name", "id": "id1", "idInt": -1,
		"category": "cat", "entity": "entity", "hasSingle": true,
	}
	expected.ShouldBeEqual(t, 0, "payloadProperties returns correct value -- with args", actual)
}

func Test_PayloadProperties_Setters(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()
	props := pw.PayloadProperties()
	_ = props.SetName("n")
	props.SetNameMust("n2")
	_ = props.SetIdString("id")
	props.SetIdStringMust("id2")
	_ = props.SetCategory("c")
	props.SetCategoryMust("c2")
	_ = props.SetEntityType("e")
	props.SetEntityTypeMust("e2")
	props.SetManyRecordFlag()
	_ = props.SetDynamicPayloads([]byte("d"))
	props.SetDynamicPayloadsMust([]byte("d2"))

	// Act
	actual := args.Map{
		"name": pw.Name, "id": pw.Identifier,
		"cat": pw.CategoryName, "entity": pw.EntityType,
		"many": pw.HasManyRecords,
	}

	// Assert
	expected := args.Map{
		"name": "n2", "id": "id2", "cat": "c2",
		"entity": "e2", "many": true,
	}
	expected.ShouldBeEqual(t, 0, "payloadProperties returns correct value -- setters", actual)
}
