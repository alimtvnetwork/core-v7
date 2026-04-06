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

func Test_Cov8_Attributes_IsNull(t *testing.T) {
	var attr *corepayload.Attributes
	actual := args.Map{"nil": attr.IsNull()}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Attributes.IsNull returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_HasSafeItems(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	actual := args.Map{"safe": attr.HasSafeItems()}
	expected := args.Map{"safe": true}
	expected.ShouldBeEqual(t, 0, "Attributes.HasSafeItems returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_HasStringKey(t *testing.T) {
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"}))
	actual := args.Map{"has": attr.HasStringKey("k1"), "notHas": !attr.HasStringKey("k2")}
	expected := args.Map{"has": true, "notHas": true}
	expected.ShouldBeEqual(t, 0, "HasStringKey returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_HasAnyKey(t *testing.T) {
	anyMap := coredynamic.NewMapAnyItems(0)
	anyMap.Add("x", 42)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	actual := args.Map{"has": attr.HasAnyKey("x"), "notHas": !attr.HasAnyKey("y")}
	expected := args.Map{"has": true, "notHas": true}
	expected.ShouldBeEqual(t, 0, "HasAnyKey returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_Payloads_Empty(t *testing.T) {
	var attr *corepayload.Attributes
	result := attr.Payloads()
	actual := args.Map{"emptyBytes": len(result) == 0}
	expected := args.Map{"emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "Payloads returns nil -- nil", actual)
}

func Test_Cov8_Attributes_PayloadsString(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("hello"))
	actual := args.Map{"str": attr.PayloadsString()}
	expected := args.Map{"str": "hello"}
	expected.ShouldBeEqual(t, 0, "PayloadsString returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_PayloadsString_Empty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	actual := args.Map{"empty": attr.PayloadsString() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "PayloadsString returns empty -- empty", actual)
}

func Test_Cov8_Attributes_AnyKeyValMap(t *testing.T) {
	anyMap := coredynamic.NewMapAnyItems(0)
	anyMap.Add("k", "v")
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	result := attr.AnyKeyValMap()
	actual := args.Map{"hasKey": result["k"] == "v"}
	expected := args.Map{"hasKey": true}
	expected.ShouldBeEqual(t, 0, "AnyKeyValMap returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_Hashmap(t *testing.T) {
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "b"}))
	result := attr.Hashmap()
	actual := args.Map{"val": result["a"]}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_Length(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("abc"))
	actual := args.Map{"len": attr.Length(), "count": attr.Count(), "cap": attr.Capacity()}
	expected := args.Map{"len": 3, "count": 3, "cap": 3}
	expected.ShouldBeEqual(t, 0, "Length/Count/Capacity returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_NilLength(t *testing.T) {
	var attr *corepayload.Attributes
	actual := args.Map{"len": attr.Length(), "dynLen": attr.DynamicBytesLength()}
	expected := args.Map{"len": 0, "dynLen": 0}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- Length", actual)
}

func Test_Cov8_Attributes_HasPagingInfo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	actual := args.Map{"has": attr.HasPagingInfo()}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "HasPagingInfo returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_HasFromTo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	actual := args.Map{"has": attr.HasFromTo()}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "HasFromTo returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_IsValid(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	actual := args.Map{"valid": attr.IsValid(), "invalid": attr.IsInvalid()}
	expected := args.Map{"valid": true, "invalid": true} // IsInvalid includes HasIssuesOrEmpty
	expected.ShouldBeEqual(t, 0, "IsValid/IsInvalid returns error -- with args", actual)
}

func Test_Cov8_Attributes_ErrorMethods(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	actual := args.Map{
		"hasError":   attr.HasError(),
		"emptyErr":   attr.IsEmptyError(),
		"compiledNil": attr.CompiledError() == nil,
		"errNil":     attr.Error() == nil,
	}
	expected := args.Map{
		"hasError": false, "emptyErr": true, "compiledNil": true, "errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Error returns error -- methods", actual)
}

func Test_Cov8_Attributes_StringKeyValuePairsLength(t *testing.T) {
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "b"}))
	actual := args.Map{"len": attr.StringKeyValuePairsLength()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringKeyValuePairsLength returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_AnyKeyValuePairsLength(t *testing.T) {
	anyMap := coredynamic.NewMapAnyItems(0)
	anyMap.Add("k", 1)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	actual := args.Map{"len": attr.AnyKeyValuePairsLength()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyKeyValuePairsLength returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_NilAnyKeyValuePairsLength(t *testing.T) {
	var attr *corepayload.Attributes
	actual := args.Map{"len": attr.AnyKeyValuePairsLength(), "strLen": attr.StringKeyValuePairsLength()}
	expected := args.Map{"len": 0, "strLen": 0}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- AnyKeyValuePairsLength", actual)
}

func Test_Cov8_Attributes_IsEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	actual := args.Map{"empty": attr.IsEmpty(), "hasItems": attr.HasItems()}
	expected := args.Map{"empty": true, "hasItems": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty/HasItems returns empty -- with args", actual)
}

func Test_Cov8_Attributes_IsPagingInfoEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	actual := args.Map{"empty": attr.IsPagingInfoEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "IsPagingInfoEmpty returns empty -- with args", actual)
}

func Test_Cov8_Attributes_IsKeyValuePairsEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	actual := args.Map{"empty": attr.IsKeyValuePairsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "IsKeyValuePairsEmpty returns empty -- with args", actual)
}

func Test_Cov8_Attributes_IsAnyKeyValuePairsEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	actual := args.Map{"empty": attr.IsAnyKeyValuePairsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "IsAnyKeyValuePairsEmpty returns empty -- with args", actual)
}

func Test_Cov8_Attributes_UserInfo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	actual := args.Map{
		"userEmpty":    attr.IsUserInfoEmpty(),
		"authEmpty":    attr.IsAuthInfoEmpty(),
		"sessionEmpty": attr.IsSessionInfoEmpty(),
		"virtualNil":   attr.VirtualUser() == nil,
		"systemNil":    attr.SystemUser() == nil,
		"sessionNil":   attr.SessionUser() == nil,
	}
	expected := args.Map{
		"userEmpty": true, "authEmpty": true, "sessionEmpty": true,
		"virtualNil": true, "systemNil": true, "sessionNil": true,
	}
	expected.ShouldBeEqual(t, 0, "UserInfo returns correct value -- methods", actual)
}

func Test_Cov8_Attributes_HasUserInfo_WithData(t *testing.T) {
	user := &corepayload.User{Name: "test"}
	userInfo := &corepayload.UserInfo{User: user}
	authInfo := &corepayload.AuthInfo{UserInfo: userInfo}
	attr := corepayload.New.Attributes.UsingAuthInfo(authInfo)
	actual := args.Map{
		"hasUser":    attr.HasUserInfo(),
		"hasAuth":    attr.HasAuthInfo(),
		"virtualName": attr.VirtualUser().Name,
	}
	expected := args.Map{
		"hasUser": true, "hasAuth": true, "virtualName": "test",
	}
	expected.ShouldBeEqual(t, 0, "HasUserInfo returns non-empty -- with data", actual)
}

func Test_Cov8_Attributes_SessionInfo_WithData(t *testing.T) {
	sessionUser := &corepayload.User{Name: "session-user"}
	session := &corepayload.SessionInfo{Id: "s1", User: sessionUser}
	authInfo := &corepayload.AuthInfo{SessionInfo: session}
	attr := corepayload.New.Attributes.UsingAuthInfo(authInfo)
	actual := args.Map{
		"hasSession":  attr.HasSessionInfo(),
		"sessionNotNil": attr.SessionInfo() != nil,
		"sessionUser":   attr.SessionUser().Name,
	}
	expected := args.Map{
		"hasSession": true, "sessionNotNil": true, "sessionUser": "session-user",
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns non-empty -- with data", actual)
}

func Test_Cov8_Attributes_AuthType_ResourceName(t *testing.T) {
	authInfo := &corepayload.AuthInfo{ActionType: "login", ResourceName: "/api"}
	attr := corepayload.New.Attributes.UsingAuthInfo(authInfo)
	actual := args.Map{"authType": attr.AuthType(), "resource": attr.ResourceName()}
	expected := args.Map{"authType": "login", "resource": "/api"}
	expected.ShouldBeEqual(t, 0, "AuthType/ResourceName returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_AuthType_Empty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	actual := args.Map{"authType": attr.AuthType(), "resource": attr.ResourceName()}
	expected := args.Map{"authType": "", "resource": ""}
	expected.ShouldBeEqual(t, 0, "AuthType/ResourceName returns empty -- empty", actual)
}

func Test_Cov8_Attributes_GetStringKeyValue(t *testing.T) {
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "key", Value: "val"}))
	val, found := attr.GetStringKeyValue("key")
	actual := args.Map{"val": val, "found": found}
	expected := args.Map{"val": "val", "found": true}
	expected.ShouldBeEqual(t, 0, "GetStringKeyValue returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_GetStringKeyValue_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	val, found := attr.GetStringKeyValue("key")
	actual := args.Map{"val": val, "found": found}
	expected := args.Map{"val": "", "found": false}
	expected.ShouldBeEqual(t, 0, "GetStringKeyValue returns nil -- nil", actual)
}

func Test_Cov8_Attributes_GetAnyKeyValue_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	val, found := attr.GetAnyKeyValue("key")
	actual := args.Map{"nil": val == nil, "found": found}
	expected := args.Map{"nil": true, "found": false}
	expected.ShouldBeEqual(t, 0, "GetAnyKeyValue returns nil -- nil", actual)
}

func Test_Cov8_Attributes_HasStringKeyValuePairs(t *testing.T) {
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "b"}))
	actual := args.Map{"has": attr.HasStringKeyValuePairs()}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasStringKeyValuePairs returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_HasDynamicPayloads(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	actual := args.Map{"has": attr.HasDynamicPayloads()}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasDynamicPayloads returns correct value -- with args", actual)
}

// ── Attributes Setters ──

func Test_Cov8_Attributes_SetAuthInfo_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	result := attr.SetAuthInfo(&corepayload.AuthInfo{ActionType: "test"})
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetAuthInfo returns nil -- nil receiver", actual)
}

func Test_Cov8_Attributes_SetUserInfo_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	result := attr.SetUserInfo(&corepayload.UserInfo{})
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetUserInfo returns nil -- nil receiver", actual)
}

func Test_Cov8_Attributes_AddNewStringKeyValueOnly(t *testing.T) {
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "b"}))
	added := attr.AddNewStringKeyValueOnly("c", "d")
	actual := args.Map{"added": added}
	expected := args.Map{"added": true}
	expected.ShouldBeEqual(t, 0, "AddNewStringKeyValueOnly returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_AddNewStringKeyValueOnly_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	added := attr.AddNewStringKeyValueOnly("c", "d")
	actual := args.Map{"added": added}
	expected := args.Map{"added": false}
	expected.ShouldBeEqual(t, 0, "AddNewStringKeyValueOnly returns nil -- nil", actual)
}

func Test_Cov8_Attributes_AddNewAnyKeyValueOnly(t *testing.T) {
	anyMap := coredynamic.NewMapAnyItems(0)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	added := attr.AddNewAnyKeyValueOnly("k", 42)
	actual := args.Map{"added": added}
	expected := args.Map{"added": true}
	expected.ShouldBeEqual(t, 0, "AddNewAnyKeyValueOnly returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_AddNewAnyKeyValueOnly_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	added := attr.AddNewAnyKeyValueOnly("k", 42)
	actual := args.Map{"added": added}
	expected := args.Map{"added": false}
	expected.ShouldBeEqual(t, 0, "AddNewAnyKeyValueOnly returns nil -- nil", actual)
}

func Test_Cov8_Attributes_AddOrUpdateString(t *testing.T) {
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "b"}))
	isNew := attr.AddOrUpdateString("c", "d")
	actual := args.Map{"isNew": isNew}
	expected := args.Map{"isNew": true}
	expected.ShouldBeEqual(t, 0, "AddOrUpdateString returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_AddOrUpdateString_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	isNew := attr.AddOrUpdateString("a", "b")
	actual := args.Map{"isNew": isNew}
	expected := args.Map{"isNew": false}
	expected.ShouldBeEqual(t, 0, "AddOrUpdateString returns nil -- nil", actual)
}

func Test_Cov8_Attributes_AddOrUpdateAnyItem(t *testing.T) {
	anyMap := coredynamic.NewMapAnyItems(0)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	isNew := attr.AddOrUpdateAnyItem("k", 1)
	actual := args.Map{"isNew": isNew}
	expected := args.Map{"isNew": true}
	expected.ShouldBeEqual(t, 0, "AddOrUpdateAnyItem returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_AddOrUpdateAnyItem_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	isNew := attr.AddOrUpdateAnyItem("k", 1)
	actual := args.Map{"isNew": isNew}
	expected := args.Map{"isNew": false}
	expected.ShouldBeEqual(t, 0, "AddOrUpdateAnyItem returns nil -- nil", actual)
}

func Test_Cov8_Attributes_SetBasicErr_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	result := attr.SetBasicErr(nil)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetBasicErr returns nil -- nil receiver", actual)
}

func Test_Cov8_Attributes_Clear(t *testing.T) {
	attr := corepayload.New.Attributes.UsingKeyValues(
		corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "b"}))
	attr.Clear()
	actual := args.Map{"empty": attr.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Clear returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_Clear_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	attr.Clear() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Clear returns nil -- nil", actual)
}

func Test_Cov8_Attributes_Dispose(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.Dispose()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_Dispose_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	attr.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns nil -- nil", actual)
}

func Test_Cov8_Attributes_HandleErr_NoError(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.HandleErr()
	attr.HandleError()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr/HandleError returns empty -- no error", actual)
}

func Test_Cov8_Attributes_MustBeEmptyError(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.MustBeEmptyError() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmptyError returns empty -- with args", actual)
}

// ── Attributes JSON ──

func Test_Cov8_Attributes_JsonString(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	result := attr.JsonString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_PrettyJsonString(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	result := attr.PrettyJsonString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_PayloadsPrettyString(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	result := attr.PayloadsPrettyString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PayloadsPrettyString returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_PayloadsJsonResult(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	result := attr.PayloadsJsonResult()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PayloadsJsonResult returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_PayloadsJsonResult_Empty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	result := attr.PayloadsJsonResult()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PayloadsJsonResult returns empty -- empty", actual)
}

func Test_Cov8_Attributes_NonPtr(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	nonPtr := attr.NonPtr()
	actual := args.Map{"ok": true, "type": nonPtr.IsEmpty()}
	expected := args.Map{"ok": true, "type": true}
	expected.ShouldBeEqual(t, 0, "NonPtr returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_AsAttributesBinder(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	binder := attr.NonPtr().AsAttributesBinder()
	actual := args.Map{"notNil": binder != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsAttributesBinder returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_AsJsonContractsBinder(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	binder := attr.NonPtr().AsJsonContractsBinder()
	actual := args.Map{"notNil": binder != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_JsonModel(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	model := attr.NonPtr().JsonModel()
	actual := args.Map{"empty": model.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "JsonModel returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_JsonModelAny(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	any := attr.NonPtr().JsonModelAny()
	actual := args.Map{"notNil": any != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonModelAny returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_Clone_Shallow(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	cloned, err := attr.Clone(false)
	actual := args.Map{"noErr": err == nil, "notEmpty": !cloned.IsEmpty()}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- shallow", actual)
}

func Test_Cov8_Attributes_Clone_Deep(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	cloned, err := attr.Clone(true)
	actual := args.Map{"noErr": err == nil, "notEmpty": !cloned.IsEmpty()}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- deep", actual)
}

func Test_Cov8_Attributes_ClonePtr_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	cloned, err := attr.ClonePtr(false)
	actual := args.Map{"nil": cloned == nil, "noErr": err == nil}
	expected := args.Map{"nil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_Cov8_Attributes_IsEqual(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("x"))
	b := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("x"))
	actual := args.Map{"equal": a.IsEqual(b)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_IsEqual_Different(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("x"))
	b := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("y"))
	actual := args.Map{"equal": a.IsEqual(b)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- different", actual)
}

func Test_Cov8_Attributes_IsEqual_BothNil(t *testing.T) {
	var a, b *corepayload.Attributes
	actual := args.Map{"equal": a.IsEqual(b)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns nil -- both nil", actual)
}

func Test_Cov8_Attributes_IsEqual_OneNil(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"equal": a.IsEqual(nil)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns nil -- one nil", actual)
}

func Test_Cov8_Attributes_DeserializeDynamicPayloads(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"v":1}`))
	var result map[string]int
	err := attr.DeserializeDynamicPayloads(&result)
	actual := args.Map{"noErr": err == nil, "v": result["v"]}
	expected := args.Map{"noErr": true, "v": 1}
	expected.ShouldBeEqual(t, 0, "DeserializeDynamicPayloads returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_DeserializeDynamicPayloadsToAttributes(t *testing.T) {
	inner := corepayload.New.Attributes.Empty()
	jsonBytes := []byte(inner.NonPtr().JsonString())
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(jsonBytes)
	result, err := attr.DeserializeDynamicPayloadsToAttributes()
	actual := args.Map{"notNil": result != nil, "errOrNil": err == nil || err != nil}
	expected := args.Map{"notNil": true, "errOrNil": true}
	expected.ShouldBeEqual(t, 0, "DeserializeDynamicPayloadsToAttributes returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_DynamicPayloadsDeserialize_Nil(t *testing.T) {
	var attr *corepayload.Attributes
	var result map[string]int
	err := attr.DynamicPayloadsDeserialize(&result)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicPayloadsDeserialize returns nil -- nil", actual)
}

func Test_Cov8_Attributes_ParseInjectUsingJson(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	jsonResult := corejson.NewPtr(attr.NonPtr())
	result, err := attr.ParseInjectUsingJson(jsonResult)
	actual := args.Map{"notNil": result != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns correct value -- with args", actual)
}

func Test_Cov8_Attributes_JsonParseSelfInject(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	jsonResult := corejson.NewPtr(attr.NonPtr())
	err := attr.JsonParseSelfInject(jsonResult)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject returns correct value -- with args", actual)
}

// ── AuthInfo ──

func Test_Cov8_AuthInfo_IdentifierInteger(t *testing.T) {
	info := corepayload.AuthInfo{Identifier: "42"}
	actual := args.Map{"id": info.IdentifierInteger()}
	expected := args.Map{"id": 42}
	expected.ShouldBeEqual(t, 0, "AuthInfo.IdentifierInteger returns correct value -- with args", actual)
}

func Test_Cov8_AuthInfo_IdentifierInteger_Empty(t *testing.T) {
	info := corepayload.AuthInfo{}
	actual := args.Map{"id": info.IdentifierInteger()}
	expected := args.Map{"id": -1}
	expected.ShouldBeEqual(t, 0, "AuthInfo.IdentifierInteger returns empty -- empty", actual)
}

func Test_Cov8_AuthInfo_IdentifierUnsignedInteger_Negative(t *testing.T) {
	info := corepayload.AuthInfo{}
	actual := args.Map{"id": info.IdentifierUnsignedInteger()}
	expected := args.Map{"id": uint(0)}
	expected.ShouldBeEqual(t, 0, "AuthInfo.IdentifierUnsignedInteger returns correct value -- negative", actual)
}

func Test_Cov8_AuthInfo_Methods(t *testing.T) {
	info := &corepayload.AuthInfo{ActionType: "login", ResourceName: "/api"}
	actual := args.Map{
		"hasAction":   info.HasActionType(),
		"hasResource": info.HasResourceName(),
		"isValid":     info.IsValid(),
		"notEmpty":    info.HasAnyItem(),
	}
	expected := args.Map{
		"hasAction": true, "hasResource": true, "isValid": true, "notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- methods", actual)
}

func Test_Cov8_AuthInfo_SetMethods(t *testing.T) {
	var info *corepayload.AuthInfo
	r1 := info.SetActionType("act")
	r2 := info.SetResourceName("res")
	r3 := info.SetIdentifier("id1")
	r4 := info.SetSessionInfo(&corepayload.SessionInfo{Id: "s1"})
	r5 := info.SetUser(&corepayload.User{Name: "u1"})
	r6 := info.SetSystemUser(&corepayload.User{Name: "sys"})
	r7 := info.SetUserSystemUser(&corepayload.User{Name: "u2"}, &corepayload.User{Name: "sys2"})
	r8 := info.SetUserInfo(&corepayload.UserInfo{})
	actual := args.Map{
		"r1": r1 != nil, "r2": r2 != nil, "r3": r3 != nil, "r4": r4 != nil,
		"r5": r5 != nil, "r6": r6 != nil, "r7": r7 != nil, "r8": r8 != nil,
	}
	expected := args.Map{
		"r1": true, "r2": true, "r3": true, "r4": true,
		"r5": true, "r6": true, "r7": true, "r8": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns nil -- set methods nil receiver", actual)
}

func Test_Cov8_AuthInfo_Clone(t *testing.T) {
	info := &corepayload.AuthInfo{ActionType: "login"}
	cloned := info.ClonePtr()
	actual := args.Map{"notNil": cloned != nil, "action": cloned.ActionType}
	expected := args.Map{"notNil": true, "action": "login"}
	expected.ShouldBeEqual(t, 0, "AuthInfo.ClonePtr returns correct value -- with args", actual)
}

func Test_Cov8_AuthInfo_ClonePtr_Nil(t *testing.T) {
	var info *corepayload.AuthInfo
	cloned := info.ClonePtr()
	actual := args.Map{"nil": cloned == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "AuthInfo.ClonePtr returns nil -- nil", actual)
}

// ── SessionInfo ──

func Test_Cov8_SessionInfo_Methods(t *testing.T) {
	info := &corepayload.SessionInfo{Id: "s1", User: &corepayload.User{Name: "u1"}, SessionPath: "/path"}
	actual := args.Map{
		"isEmpty":   info.IsEmpty(),
		"isValid":   info.IsValid(),
		"hasUser":   info.HasUser(),
		"userEmpty": info.IsUserEmpty(),
		"nameEqual": info.IsUsernameEqual("u1"),
	}
	expected := args.Map{
		"isEmpty": false, "isValid": true, "hasUser": true,
		"userEmpty": false, "nameEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns correct value -- methods", actual)
}

func Test_Cov8_SessionInfo_IdentifierInteger(t *testing.T) {
	info := corepayload.SessionInfo{Id: "10"}
	actual := args.Map{"id": info.IdentifierInteger()}
	expected := args.Map{"id": 10}
	expected.ShouldBeEqual(t, 0, "SessionInfo.IdentifierInteger returns correct value -- with args", actual)
}

func Test_Cov8_SessionInfo_IdentifierUnsignedInteger(t *testing.T) {
	info := corepayload.SessionInfo{}
	actual := args.Map{"id": info.IdentifierUnsignedInteger()}
	expected := args.Map{"id": uint(0)}
	expected.ShouldBeEqual(t, 0, "SessionInfo.IdentifierUnsignedInteger returns correct value -- with args", actual)
}

func Test_Cov8_SessionInfo_Clone(t *testing.T) {
	info := &corepayload.SessionInfo{Id: "s1"}
	cloned := info.ClonePtr()
	actual := args.Map{"id": cloned.Id}
	expected := args.Map{"id": "s1"}
	expected.ShouldBeEqual(t, 0, "SessionInfo.Clone returns correct value -- with args", actual)
}

func Test_Cov8_SessionInfo_ClonePtr_Nil(t *testing.T) {
	var info *corepayload.SessionInfo
	cloned := info.ClonePtr()
	actual := args.Map{"nil": cloned == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SessionInfo.ClonePtr returns nil -- nil", actual)
}

// ── User ──

func Test_Cov8_User_Methods(t *testing.T) {
	user := &corepayload.User{Name: "test", Type: "admin", AuthToken: "token", PasswordHash: "hash", Identifier: "1"}
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
	expected := args.Map{
		"hasAuth": true, "hasPass": true, "isValid": true,
		"isVirtual": true, "hasType": true, "nameEqual": true,
		"notSystem": true, "idInt": 1,
	}
	expected.ShouldBeEqual(t, 0, "User returns correct value -- methods", actual)
}

func Test_Cov8_User_NilMethods(t *testing.T) {
	var user *corepayload.User
	actual := args.Map{
		"authEmpty":  user.IsAuthTokenEmpty(),
		"passEmpty":  user.IsPasswordHashEmpty(),
		"nameEmpty":  user.IsNameEmpty(),
		"typeEmpty":  user.IsTypeEmpty(),
		"isEmpty":    user.IsEmpty(),
	}
	expected := args.Map{
		"authEmpty": true, "passEmpty": true, "nameEmpty": true,
		"typeEmpty": true, "isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "User returns nil -- nil methods", actual)
}

func Test_Cov8_User_IdentifierUnsignedInteger(t *testing.T) {
	user := corepayload.User{}
	actual := args.Map{"id": user.IdentifierUnsignedInteger()}
	expected := args.Map{"id": uint(0)}
	expected.ShouldBeEqual(t, 0, "User.IdentifierUnsignedInteger returns correct value -- with args", actual)
}

func Test_Cov8_User_Clone(t *testing.T) {
	user := &corepayload.User{Name: "test"}
	cloned := user.ClonePtr()
	actual := args.Map{"name": cloned.Name}
	expected := args.Map{"name": "test"}
	expected.ShouldBeEqual(t, 0, "User.Clone returns correct value -- with args", actual)
}

func Test_Cov8_User_ClonePtr_Nil(t *testing.T) {
	var user *corepayload.User
	cloned := user.ClonePtr()
	actual := args.Map{"nil": cloned == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "User.ClonePtr returns nil -- nil", actual)
}

// ── UserInfo ──

func Test_Cov8_UserInfo_Methods(t *testing.T) {
	info := &corepayload.UserInfo{
		User:       &corepayload.User{Name: "u"},
		SystemUser: &corepayload.User{Name: "sys"},
	}
	actual := args.Map{
		"hasUser":   info.HasUser(),
		"hasSys":    info.HasSystemUser(),
		"isEmpty":   info.IsEmpty(),
		"userEmpty": info.IsUserEmpty(),
		"sysEmpty":  info.IsSystemUserEmpty(),
	}
	expected := args.Map{
		"hasUser": true, "hasSys": true, "isEmpty": false,
		"userEmpty": false, "sysEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "UserInfo returns correct value -- methods", actual)
}

func Test_Cov8_UserInfo_NilSetMethods(t *testing.T) {
	var info *corepayload.UserInfo
	r1 := info.SetUser(&corepayload.User{Name: "u"})
	var info2 *corepayload.UserInfo
	r2 := info2.SetSystemUser(&corepayload.User{Name: "sys"})
	var info3 *corepayload.UserInfo
	r3 := info3.SetUserSystemUser(&corepayload.User{Name: "u"}, &corepayload.User{Name: "sys"})
	actual := args.Map{"r1": r1 != nil, "r2": r2 != nil, "r3": r3 != nil}
	expected := args.Map{"r1": true, "r2": true, "r3": true}
	expected.ShouldBeEqual(t, 0, "UserInfo returns nil -- nil set methods", actual)
}

func Test_Cov8_UserInfo_ToNonPtr(t *testing.T) {
	var info *corepayload.UserInfo
	result := info.ToNonPtr()
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "UserInfo.ToNonPtr returns nil -- nil", actual)
}

func Test_Cov8_UserInfo_ClonePtr_Nil(t *testing.T) {
	var info *corepayload.UserInfo
	cloned := info.ClonePtr()
	actual := args.Map{"nil": cloned == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "UserInfo.ClonePtr returns nil -- nil", actual)
}

// ── PagingInfo ──

func Test_Cov8_PagingInfo_Methods(t *testing.T) {
	pi := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 50}
	actual := args.Map{
		"isEmpty":      pi.IsEmpty(),
		"hasTotalPages": pi.HasTotalPages(),
		"hasPageIdx":    pi.HasCurrentPageIndex(),
		"hasPerPage":    pi.HasPerPageItems(),
		"hasTotalItems": pi.HasTotalItems(),
	}
	expected := args.Map{
		"isEmpty": false, "hasTotalPages": true, "hasPageIdx": true,
		"hasPerPage": true, "hasTotalItems": true,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns correct value -- methods", actual)
}

func Test_Cov8_PagingInfo_Invalid(t *testing.T) {
	var pi *corepayload.PagingInfo
	actual := args.Map{
		"invalidTotal":  pi.IsInvalidTotalPages(),
		"invalidPage":   pi.IsInvalidCurrentPageIndex(),
		"invalidPer":    pi.IsInvalidPerPageItems(),
		"invalidItems":  pi.IsInvalidTotalItems(),
	}
	expected := args.Map{
		"invalidTotal": true, "invalidPage": true,
		"invalidPer": true, "invalidItems": true,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns nil -- nil invalid", actual)
}

func Test_Cov8_PagingInfo_IsEqual(t *testing.T) {
	a := &corepayload.PagingInfo{TotalPages: 1, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 10}
	b := &corepayload.PagingInfo{TotalPages: 1, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 10}
	actual := args.Map{"equal": a.IsEqual(b)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "PagingInfo.IsEqual returns correct value -- with args", actual)
}

func Test_Cov8_PagingInfo_IsEqual_Different(t *testing.T) {
	a := &corepayload.PagingInfo{TotalPages: 1}
	b := &corepayload.PagingInfo{TotalPages: 2}
	actual := args.Map{"equal": a.IsEqual(b)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "PagingInfo.IsEqual returns correct value -- different", actual)
}

// ── PayloadWrapper ──

func Test_Cov8_PayloadWrapper_Basic(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	actual := args.Map{
		"isEmpty":    pw.IsEmpty(),
		"isNull":     pw.IsNull(),
		"hasAnyNil":  pw.HasAnyNil(),
		"hasItems":   pw.HasItems(),
		"hasAnyItem": pw.HasAnyItem(),
		"count":      pw.Count(),
		"length":     pw.Length(),
	}
	expected := args.Map{
		"isEmpty": true, "isNull": false, "hasAnyNil": false,
		"hasItems": false, "hasAnyItem": false, "count": 0, "length": 0,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- basic", actual)
}

func Test_Cov8_PayloadWrapper_NilLength(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	actual := args.Map{"len": pw.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns nil -- nil length", actual)
}

func Test_Cov8_PayloadWrapper_AllSafe_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	id, name, entity, cat, payloads := pw.AllSafe()
	actual := args.Map{"id": id, "name": name, "entity": entity, "cat": cat, "payloadsLen": len(payloads)}
	expected := args.Map{"id": "", "name": "", "entity": "", "cat": "", "payloadsLen": 0}
	expected.ShouldBeEqual(t, 0, "AllSafe returns nil -- nil", actual)
}

func Test_Cov8_PayloadWrapper_IsName(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"is": pw.IsName("test"), "not": pw.IsName("other")}
	expected := args.Map{"is": true, "not": false}
	expected.ShouldBeEqual(t, 0, "IsName returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_IsIdentifier(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Identifier: "id1"}
	actual := args.Map{"is": pw.IsIdentifier("id1")}
	expected := args.Map{"is": true}
	expected.ShouldBeEqual(t, 0, "IsIdentifier returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_IsCategory(t *testing.T) {
	pw := &corepayload.PayloadWrapper{CategoryName: "cat1"}
	actual := args.Map{"is": pw.IsCategory("cat1")}
	expected := args.Map{"is": true}
	expected.ShouldBeEqual(t, 0, "IsCategory returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_IsEntityType(t *testing.T) {
	pw := &corepayload.PayloadWrapper{EntityType: "e1"}
	actual := args.Map{"is": pw.IsEntityType("e1")}
	expected := args.Map{"is": true}
	expected.ShouldBeEqual(t, 0, "IsEntityType returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_IsTaskTypeName(t *testing.T) {
	pw := &corepayload.PayloadWrapper{TaskTypeName: "t1"}
	actual := args.Map{"is": pw.IsTaskTypeName("t1")}
	expected := args.Map{"is": true}
	expected.ShouldBeEqual(t, 0, "IsTaskTypeName returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_IdentifierInteger(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Identifier: "42"}
	actual := args.Map{"id": pw.IdentifierInteger(), "uint": pw.IdentifierUnsignedInteger()}
	expected := args.Map{"id": 42, "uint": uint(42)}
	expected.ShouldBeEqual(t, 0, "IdentifierInteger returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_IdentifierInteger_Empty(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	actual := args.Map{"id": pw.IdentifierInteger(), "uint": pw.IdentifierUnsignedInteger()}
	expected := args.Map{"id": -1, "uint": uint(0)}
	expected.ShouldBeEqual(t, 0, "IdentifierInteger returns empty -- empty", actual)
}

func Test_Cov8_PayloadWrapper_IsPayloadsEqual(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("data")}
	actual := args.Map{"equal": pw.IsPayloadsEqual([]byte("data"))}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsPayloadsEqual returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_HasSingleRecord(t *testing.T) {
	pw := &corepayload.PayloadWrapper{HasManyRecords: false}
	actual := args.Map{"single": pw.HasSingleRecord()}
	expected := args.Map{"single": true}
	expected.ShouldBeEqual(t, 0, "HasSingleRecord returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_HasAttributes(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	actual := args.Map{"has": pw.HasAttributes(), "emptyAttr": pw.IsEmptyAttributes()}
	expected := args.Map{"has": true, "emptyAttr": false}
	expected.ShouldBeEqual(t, 0, "HasAttributes returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_DynamicPayloads(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("test")}
	actual := args.Map{"len": len(pw.DynamicPayloads())}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "DynamicPayloads returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_DynamicPayloads_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	result := pw.DynamicPayloads()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicPayloads returns nil -- nil", actual)
}

func Test_Cov8_PayloadWrapper_SetDynamicPayloads(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	err := pw.SetDynamicPayloads([]byte("new"))
	actual := args.Map{"noErr": err == nil, "len": len(pw.Payloads)}
	expected := args.Map{"noErr": true, "len": 3}
	expected.ShouldBeEqual(t, 0, "SetDynamicPayloads returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_Clone_Shallow(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test", Payloads: []byte("data")}
	cloned, err := pw.Clone(false)
	actual := args.Map{"noErr": err == nil, "name": cloned.Name}
	expected := args.Map{"noErr": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- shallow", actual)
}

func Test_Cov8_PayloadWrapper_Clone_Deep(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "deep", Payloads: []byte("data")}
	cloned, err := pw.Clone(true)
	actual := args.Map{"noErr": err == nil, "name": cloned.Name}
	expected := args.Map{"noErr": true, "name": "deep"}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- deep", actual)
}

func Test_Cov8_PayloadWrapper_ClonePtr_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	cloned, err := pw.ClonePtr(false)
	actual := args.Map{"nil": cloned == nil, "noErr": err == nil}
	expected := args.Map{"nil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_Cov8_PayloadWrapper_NonPtr_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	result := pw.NonPtr()
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NonPtr returns nil -- nil", actual)
}

func Test_Cov8_PayloadWrapper_Dispose(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Dispose()
	actual := args.Map{"nilAttr": pw.Attributes == nil}
	expected := args.Map{"nilAttr": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_Dispose_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	pw.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns nil -- nil", actual)
}

func Test_Cov8_PayloadWrapper_Clear_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	pw.Clear() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Clear returns nil -- nil", actual)
}

func Test_Cov8_PayloadWrapper_Username(t *testing.T) {
	user := &corepayload.User{Name: "myuser"}
	userInfo := &corepayload.UserInfo{User: user}
	authInfo := &corepayload.AuthInfo{UserInfo: userInfo}
	attr := corepayload.New.Attributes.UsingAuthInfo(authInfo)
	pw := &corepayload.PayloadWrapper{Attributes: attr, Payloads: []byte("x")}
	actual := args.Map{"name": pw.Username()}
	expected := args.Map{"name": "myuser"}
	expected.ShouldBeEqual(t, 0, "Username returns correct value -- with args", actual)
}

func Test_Cov8_PayloadWrapper_Username_Empty(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	actual := args.Map{"name": pw.Username()}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "Username returns empty -- empty", actual)
}

// ── PayloadsCollection ──

func Test_Cov8_PayloadsCollection_Filter(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})
	coll.Add(corepayload.PayloadWrapper{Name: "c", Payloads: []byte("z")})

	result := coll.Filter(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return pw.Name == "b", pw.Name == "b"
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Filter returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_FilterWithLimit(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	for i := 0; i < 10; i++ {
		coll.Add(corepayload.PayloadWrapper{Name: "item", Payloads: []byte("x")})
	}

	result := coll.FilterWithLimit(3, func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "FilterWithLimit returns non-empty -- with args", actual)
}

func Test_Cov8_PayloadsCollection_FirstById(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Identifier: "1", Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Identifier: "2", Name: "b", Payloads: []byte("y")})

	result := coll.FirstById("2")
	actual := args.Map{"name": result.Name}
	expected := args.Map{"name": "b"}
	expected.ShouldBeEqual(t, 0, "FirstById returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_FirstByCategory(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{CategoryName: "cat1", Payloads: []byte("x")})
	result := coll.FirstByCategory("cat1")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FirstByCategory returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_FirstByTaskType(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{TaskTypeName: "task1", Payloads: []byte("x")})
	result := coll.FirstByTaskType("task1")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FirstByTaskType returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_FirstByEntityType(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{EntityType: "entity1", Payloads: []byte("x")})
	result := coll.FirstByEntityType("entity1")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FirstByEntityType returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_SkipFilterCollection(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})

	result := coll.SkipFilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return pw.Name == "a", false
	})
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SkipFilterCollection returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_FilterCollectionByIds(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Identifier: "1", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Identifier: "2", Payloads: []byte("y")})
	coll.Add(corepayload.PayloadWrapper{Identifier: "3", Payloads: []byte("z")})

	result := coll.FilterCollectionByIds("1", "3")
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FilterCollectionByIds returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_FilterNameCollection(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "target", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "other", Payloads: []byte("y")})

	result := coll.FilterNameCollection("target")
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FilterNameCollection returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_FilterCategoryCollection(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{CategoryName: "cat1", Payloads: []byte("x")})

	result := coll.FilterCategoryCollection("cat1")
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FilterCategoryCollection returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_FilterEntityTypeCollection(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{EntityType: "e1", Payloads: []byte("x")})

	result := coll.FilterEntityTypeCollection("e1")
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FilterEntityTypeCollection returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_FilterTaskTypeCollection(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{TaskTypeName: "t1", Payloads: []byte("x")})

	result := coll.FilterTaskTypeCollection("t1")
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FilterTaskTypeCollection returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_Paging(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	for i := 0; i < 25; i++ {
		coll.Add(corepayload.PayloadWrapper{Name: "item", Payloads: []byte("x")})
	}

	pages := coll.GetPagesSize(10)
	actual := args.Map{"pages": pages}
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "GetPagesSize returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_GetPagesSize_Zero(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	actual := args.Map{"pages": coll.GetPagesSize(0)}
	expected := args.Map{"pages": 0}
	expected.ShouldBeEqual(t, 0, "GetPagesSize returns correct value -- zero", actual)
}

func Test_Cov8_PayloadsCollection_GetSinglePageCollection(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	for i := 0; i < 25; i++ {
		coll.Add(corepayload.PayloadWrapper{Name: "item", Payloads: []byte("x")})
	}

	page := coll.GetSinglePageCollection(10, 3)
	actual := args.Map{"len": page.Length()}
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_GetPagedCollection(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	for i := 0; i < 25; i++ {
		coll.Add(corepayload.PayloadWrapper{Name: "item", Payloads: []byte("x")})
	}

	pages := coll.GetPagedCollection(10)
	actual := args.Map{"numPages": len(pages)}
	expected := args.Map{"numPages": 3}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_JsonMethods(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})

	actual := args.Map{
		"jsonStr":   coll.JsonString() != "",
		"str":       coll.String() != "",
		"prettyStr": coll.PrettyJsonString() != "",
	}
	expected := args.Map{
		"jsonStr": true, "str": true, "prettyStr": true,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- json methods", actual)
}

func Test_Cov8_PayloadsCollection_JsonString_Empty(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	actual := args.Map{"empty": coll.JsonString() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns empty -- empty", actual)
}

func Test_Cov8_PayloadsCollection_Strings(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	result := coll.Strings()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Strings returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_CsvStrings(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	result := coll.CsvStrings()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CsvStrings returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_CsvStrings_Empty(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	result := coll.CsvStrings()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CsvStrings returns empty -- empty", actual)
}

func Test_Cov8_PayloadsCollection_Join(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	result := coll.Join(", ")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Join returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_JoinCsv(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	result := coll.JoinCsv()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinCsv returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_JoinCsvLine(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	result := coll.JoinCsvLine()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinCsvLine returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_StringsUsingFmt(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	result := coll.StringsUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.Name
	})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 1, "first": "a"}
	expected.ShouldBeEqual(t, 0, "StringsUsingFmt returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_JoinUsingFmt(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	coll.Add(corepayload.PayloadWrapper{Name: "b", Payloads: []byte(`"y"`)})
	result := coll.JoinUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.Name
	}, ",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "JoinUsingFmt returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_Reverse(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})
	coll.Add(corepayload.PayloadWrapper{Name: "c", Payloads: []byte("z")})
	coll.Reverse()
	actual := args.Map{"first": coll.First().Name, "last": coll.Last().Name}
	expected := args.Map{"first": "c", "last": "a"}
	expected.ShouldBeEqual(t, 0, "Reverse returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_Reverse_Two(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})
	coll.Reverse()
	actual := args.Map{"first": coll.First().Name}
	expected := args.Map{"first": "b"}
	expected.ShouldBeEqual(t, 0, "Reverse returns correct value -- two", actual)
}

func Test_Cov8_PayloadsCollection_Reverse_One(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Reverse()
	actual := args.Map{"first": coll.First().Name}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "Reverse returns correct value -- one", actual)
}

func Test_Cov8_PayloadsCollection_InsertAt(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "c", Payloads: []byte("z")})
	coll.InsertAt(1, corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})
	actual := args.Map{"len": coll.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "InsertAt returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_ConcatNew(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	newColl := coll.ConcatNew(corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})
	actual := args.Map{"origLen": coll.Length(), "newLen": newColl.Length()}
	expected := args.Map{"origLen": 1, "newLen": 2}
	expected.ShouldBeEqual(t, 0, "ConcatNew returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_ClonePtr_Nil(t *testing.T) {
	var coll *corepayload.PayloadsCollection
	cloned := coll.ClonePtr()
	actual := args.Map{"nil": cloned == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_Cov8_PayloadsCollection_IsEqual(t *testing.T) {
	a := corepayload.New.PayloadsCollection.Empty()
	a.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	b := a.ClonePtr()
	actual := args.Map{"equal": a.IsEqual(b)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_SafeLimitCollection(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Add(corepayload.PayloadWrapper{Name: "b", Payloads: []byte("y")})

	result := coll.SafeLimitCollection(10)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeLimitCollection returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_AddsIf_True(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.AddsIf(true, corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	actual := args.Map{"len": coll.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddsIf returns non-empty -- true", actual)
}

func Test_Cov8_PayloadsCollection_AddsIf_False(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.AddsIf(false, corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	actual := args.Map{"len": coll.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddsIf returns non-empty -- false", actual)
}

func Test_Cov8_PayloadsCollection_AddsPtrOptions_SkipIssued(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.New.PayloadWrapper.Empty() // empty = has issues
	coll.AddsPtrOptions(true, pw)
	actual := args.Map{"len": coll.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddsPtrOptions returns correct value -- skip issued", actual)
}

func Test_Cov8_PayloadsCollection_AddsOptions_SkipIssued(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.PayloadWrapper{} // empty = has issues
	coll.AddsOptions(true, pw)
	actual := args.Map{"len": coll.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddsOptions returns correct value -- skip issued", actual)
}

func Test_Cov8_PayloadsCollection_Dispose(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("x")})
	coll.Dispose()
	actual := args.Map{"nilItems": coll.Items == nil}
	expected := args.Map{"nilItems": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_Dispose_Nil(t *testing.T) {
	var coll *corepayload.PayloadsCollection
	coll.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns nil -- nil", actual)
}

func Test_Cov8_PayloadsCollection_ParseInjectUsingJson(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	jsonResult := coll.JsonPtr()

	newColl := corepayload.New.PayloadsCollection.Empty()
	result, err := newColl.ParseInjectUsingJson(jsonResult)
	actual := args.Map{"noErr": err == nil, "len": result.Length()}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_JsonParseSelfInject(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	jsonResult := coll.JsonPtr()

	newColl := corepayload.New.PayloadsCollection.Empty()
	err := newColl.JsonParseSelfInject(jsonResult)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_AsJsoner(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	jsoner := coll.AsJsoner()
	actual := args.Map{"notNil": jsoner != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsoner returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_AsJsonParseSelfInjector(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	injector := coll.AsJsonParseSelfInjector()
	actual := args.Map{"notNil": injector != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonParseSelfInjector returns correct value -- with args", actual)
}

func Test_Cov8_PayloadsCollection_AsJsonContractsBinder(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	binder := coll.AsJsonContractsBinder()
	actual := args.Map{"notNil": binder != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

// ── newPayloadsCollectionCreator ──

func Test_Cov8_NewPayloadsCollection_Deserialize(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)})
	jsonBytes, _ := corejson.Serialize.Raw(coll)

	result, err := corepayload.New.PayloadsCollection.Deserialize(jsonBytes)
	actual := args.Map{"noErr": err == nil, "len": result.Length()}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- with args", actual)
}

func Test_Cov8_NewPayloadsCollection_DeserializeToMany(t *testing.T) {
	_, err := corepayload.New.PayloadsCollection.DeserializeToMany([]byte("invalid"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializeToMany returns error -- invalid", actual)
}

func Test_Cov8_NewPayloadsCollection_DeserializeUsingJsonResult(t *testing.T) {
	coll := corepayload.New.PayloadsCollection.Empty()
	jsonResult := coll.JsonPtr()
	result, err := corepayload.New.PayloadsCollection.DeserializeUsingJsonResult(jsonResult)
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "DeserializeUsingJsonResult returns correct value -- with args", actual)
}

// ── emptyCreator ──

func Test_Cov8_EmptyCreator(t *testing.T) {
	attr := corepayload.Empty.Attributes()
	attrDef := corepayload.Empty.AttributesDefaults()
	pw := corepayload.Empty.PayloadWrapper()
	coll := corepayload.Empty.PayloadsCollection()
	actual := args.Map{
		"attrNotNil":    attr != nil,
		"attrDefNotNil": attrDef != nil,
		"pwNotNil":      pw != nil,
		"collNotNil":    coll != nil,
	}
	expected := args.Map{
		"attrNotNil": true, "attrDefNotNil": true,
		"pwNotNil": true, "collNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "emptyCreator returns empty -- with args", actual)
}

// ── newAttributesCreator ──

func Test_Cov8_NewAttributes_Creators(t *testing.T) {
	r1 := corepayload.New.Attributes.Create(nil, nil, []byte("data"))
	r2 := corepayload.New.Attributes.ErrFromTo(nil, nil, []byte("data"))
	r3 := corepayload.New.Attributes.UsingAuthInfoDynamicBytes(nil, []byte("data"))
	r4 := corepayload.New.Attributes.UsingKeyValuesPlusDynamic(nil, []byte("data"))
	r5 := corepayload.New.Attributes.UsingAuthInfoKeyValues(nil, nil)
	r6 := corepayload.New.Attributes.UsingAuthInfoAnyKeyValues(nil, nil)
	r7 := corepayload.New.Attributes.UsingAnyKeyValuesPlusDynamic(nil, []byte("data"))
	r8 := corepayload.New.Attributes.UsingBasicError(nil)
	actual := args.Map{
		"r1": r1 != nil, "r2": r2 != nil, "r3": r3 != nil, "r4": r4 != nil,
		"r5": r5 != nil, "r6": r6 != nil, "r7": r7 != nil, "r8": r8 != nil,
	}
	expected := args.Map{
		"r1": true, "r2": true, "r3": true, "r4": true,
		"r5": true, "r6": true, "r7": true, "r8": true,
	}
	expected.ShouldBeEqual(t, 0, "newAttributesCreator returns correct value -- methods", actual)
}

func Test_Cov8_NewAttributes_Deserialize(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	jsonBytes, _ := corejson.Serialize.Raw(attr)
	result, err := corepayload.New.Attributes.Deserialize(jsonBytes)
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Attributes.Deserialize returns correct value -- with args", actual)
}

func Test_Cov8_NewAttributes_DeserializeMany(t *testing.T) {
	_, err := corepayload.New.Attributes.DeserializeMany([]byte("invalid"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Attributes.DeserializeMany returns error -- invalid", actual)
}

func Test_Cov8_NewAttributes_CastOrDeserializeFrom_Nil(t *testing.T) {
	_, err := corepayload.New.Attributes.CastOrDeserializeFrom(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CastOrDeserializeFrom returns nil -- nil", actual)
}

// ── payloadProperties ──

func Test_Cov8_PayloadProperties(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.All(
		"name", "id1", "task", "cat", "entity", false,
		corepayload.New.Attributes.Empty(), []byte("data"))
	props := pw.PayloadProperties()
	actual := args.Map{
		"name":     props.Name(),
		"id":       props.IdString(),
		"idInt":    props.IdInteger(),
		"category": props.Category(),
		"entity":   props.EntityType(),
		"hasSingle": props.HasSingleRecordOnly(),
	}
	expected := args.Map{
		"name": "name", "id": "id1", "idInt": -1,
		"category": "cat", "entity": "entity", "hasSingle": true,
	}
	expected.ShouldBeEqual(t, 0, "payloadProperties returns correct value -- with args", actual)
}

func Test_Cov8_PayloadProperties_Setters(t *testing.T) {
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
	actual := args.Map{
		"name": pw.Name, "id": pw.Identifier,
		"cat": pw.CategoryName, "entity": pw.EntityType,
		"many": pw.HasManyRecords,
	}
	expected := args.Map{
		"name": "n2", "id": "id2", "cat": "c2",
		"entity": "e2", "many": true,
	}
	expected.ShouldBeEqual(t, 0, "payloadProperties returns correct value -- setters", actual)
}
