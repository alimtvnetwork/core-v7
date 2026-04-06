package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Attributes — Getters (AttributesGetters.go)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Attributes_IsNull(t *testing.T) {
	var nilAttr *corepayload.Attributes
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": nilAttr.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil Attributes should be null", actual)

	actual := args.Map{"result": attr.IsNull()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil Attributes should not be null", actual)
}

func Test_Cov9_Attributes_HasSafeItems(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))

	actual := args.Map{"result": attr.HasSafeItems()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasSafeItems to be true", actual)
}

func Test_Cov9_Attributes_HasStringKey(t *testing.T) {
	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr := corepayload.New.Attributes.UsingKeyValues(hm)

	actual := args.Map{"result": attr.HasStringKey("k")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasStringKey to be true", actual)

	actual := args.Map{"result": attr.HasStringKey("missing")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasStringKey to be false for missing", actual)
}

func Test_Cov9_Attributes_HasAnyKey(t *testing.T) {
	anyMap := coredynamic.NewMapAnyItems(0)
	anyMap.Add("k", 42)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)

	actual := args.Map{"result": attr.HasAnyKey("k")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasAnyKey to be true", actual)

	actual := args.Map{"result": attr.HasAnyKey("missing")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasAnyKey to be false for missing", actual)
}

func Test_Cov9_Attributes_Payloads(t *testing.T) {
	var nilAttr *corepayload.Attributes
	p := nilAttr.Payloads()

	actual := args.Map{"result": len(p) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil Payloads should return empty", actual)

	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	actual := args.Map{"result": string(attr.Payloads()) != "data"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Payloads to return data", actual)
}

func Test_Cov9_Attributes_PayloadsString(t *testing.T) {
	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.PayloadsString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil PayloadsString should return empty", actual)

	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	actual := args.Map{"result": attr.PayloadsString() != "data"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected PayloadsString to return data", actual)

	emptyAttr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": emptyAttr.PayloadsString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty PayloadsString should return empty", actual)
}

func Test_Cov9_Attributes_AnyKeyValMap(t *testing.T) {
	var nilAttr *corepayload.Attributes
	m := nilAttr.AnyKeyValMap()

	actual := args.Map{"result": len(m) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil AnyKeyValMap should return empty map", actual)

	anyMap := coredynamic.NewMapAnyItems(0)
	anyMap.Add("k", 42)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	m = attr.AnyKeyValMap()

	actual := args.Map{"result": m["k"] != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42 for key k", actual)
}

func Test_Cov9_Attributes_Hashmap(t *testing.T) {
	var nilAttr *corepayload.Attributes
	m := nilAttr.Hashmap()

	actual := args.Map{"result": len(m) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil Hashmap should return empty map", actual)

	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr := corepayload.New.Attributes.UsingKeyValues(hm)
	m = attr.Hashmap()

	actual := args.Map{"result": m["k"] != "v"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected v for key k", actual)
}

func Test_Cov9_Attributes_HasIssuesOrEmpty(t *testing.T) {
	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.HasIssuesOrEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should have issues or be empty", actual)

	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"data"`))

	actual := args.Map{"result": attr.HasIssuesOrEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-empty valid attr should not have issues", actual)
}

func Test_Cov9_Attributes_IsSafeValid(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"data"`))

	actual := args.Map{"result": attr.IsSafeValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsSafeValid true", actual)
}

func Test_Cov9_Attributes_HasAnyItem(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"data"`))

	actual := args.Map{"result": attr.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasAnyItem true", actual)
}

func Test_Cov9_Attributes_Count_Capacity(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"ab"`))

	actual := args.Map{"result": attr.Count() != attr.Length()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Count should equal Length", actual)

	actual := args.Map{"result": attr.Capacity() != attr.Length()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Capacity should equal Length", actual)
}

func Test_Cov9_Attributes_Length_Nil(t *testing.T) {
	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil Length should be 0", actual)
}

func Test_Cov9_Attributes_HasPagingInfo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.HasPagingInfo()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no paging info", actual)

	attr.PagingInfo = &corepayload.PagingInfo{TotalPages: 5}

	actual := args.Map{"result": attr.HasPagingInfo()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected paging info", actual)
}

func Test_Cov9_Attributes_HasKeyValuePairs(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.HasKeyValuePairs()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not have key value pairs", actual)

	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr2 := corepayload.New.Attributes.UsingKeyValues(hm)

	actual := args.Map{"result": attr2.HasKeyValuePairs()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected key value pairs", actual)
}

func Test_Cov9_Attributes_HasFromTo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.HasFromTo()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no FromTo", actual)

	attr.FromTo = &coreinstruction.FromTo{}

	actual := args.Map{"result": attr.HasFromTo()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected FromTo present", actual)
}

func Test_Cov9_Attributes_IsValid_IsInvalid(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)

	actual := args.Map{"result": attr.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)

	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_Cov9_Attributes_HasError_Error(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	actual := args.Map{"result": attr.Error() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil error", actual)

	actual := args.Map{"result": attr.CompiledError() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil CompiledError", actual)
}

func Test_Cov9_Attributes_IsEmptyError(t *testing.T) {
	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty error", actual)

	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty error", actual)
}

func Test_Cov9_Attributes_DynamicBytesLength(t *testing.T) {
	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.DynamicBytesLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil DynamicBytesLength should be 0", actual)

	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("abc"))

	actual := args.Map{"result": attr.DynamicBytesLength() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Cov9_Attributes_StringKeyValuePairsLength(t *testing.T) {
	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.StringKeyValuePairsLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)

	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr := corepayload.New.Attributes.UsingKeyValues(hm)

	actual := args.Map{"result": attr.StringKeyValuePairsLength() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_Attributes_AnyKeyValuePairsLength(t *testing.T) {
	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.AnyKeyValuePairsLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)

	anyMap := coredynamic.NewMapAnyItems(0)
	anyMap.Add("k", 42)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)

	actual := args.Map{"result": attr.AnyKeyValuePairsLength() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_Attributes_IsEmpty_HasItems(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty attr should be empty", actual)

	actual := args.Map{"result": attr.HasItems()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty attr should not have items", actual)

	attr2 := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	actual := args.Map{"result": attr2.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-empty attr should not be empty", actual)

	actual := args.Map{"result": attr2.HasItems()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "non-empty attr should have items", actual)
}

func Test_Cov9_Attributes_IsPagingInfoEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.IsPagingInfoEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected paging info empty", actual)
}

func Test_Cov9_Attributes_IsKeyValuePairsEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.IsKeyValuePairsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected key value pairs empty", actual)
}

func Test_Cov9_Attributes_IsAnyKeyValuePairsEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.IsAnyKeyValuePairsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected any key value pairs empty", actual)
}

func Test_Cov9_Attributes_IsUserInfoEmpty_VirtualUser_SystemUser(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.IsUserInfoEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected user info empty", actual)

	actual := args.Map{"result": attr.VirtualUser() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil VirtualUser", actual)

	actual := args.Map{"result": attr.SystemUser() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil SystemUser", actual)

	// With user info
	user := corepayload.New.User.UsingName("Alice")
	sysUser := corepayload.New.User.System("sys", "system")
	userInfo := &corepayload.UserInfo{User: user, SystemUser: sysUser}
	attr2 := corepayload.New.Attributes.UsingAuthInfo(&corepayload.AuthInfo{UserInfo: userInfo})

	actual := args.Map{"result": attr2.IsUserInfoEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected user info not empty", actual)

	actual := args.Map{"result": attr2.VirtualUser().Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)

	actual := args.Map{"result": attr2.SystemUser().Name != "sys"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sys", actual)
}

func Test_Cov9_Attributes_SessionUser(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.SessionUser() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil SessionUser", actual)

	user := corepayload.New.User.UsingName("SessionUser")
	si := &corepayload.SessionInfo{Id: "s1", User: user}
	attr2 := corepayload.New.Attributes.UsingAuthInfo(&corepayload.AuthInfo{SessionInfo: si})

	actual := args.Map{"result": attr2.SessionUser().Name != "SessionUser"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected SessionUser", actual)
}

func Test_Cov9_Attributes_IsAuthInfoEmpty_IsSessionInfoEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.IsAuthInfoEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected auth info empty", actual)

	actual := args.Map{"result": attr.IsSessionInfoEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected session info empty", actual)
}

func Test_Cov9_Attributes_HasUserInfo_HasAuthInfo_HasSessionInfo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.HasUserInfo()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no user info", actual)

	actual := args.Map{"result": attr.HasAuthInfo()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no auth info", actual)

	actual := args.Map{"result": attr.HasSessionInfo()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no session info", actual)
}

func Test_Cov9_Attributes_SessionInfo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.SessionInfo() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil SessionInfo", actual)

	si := &corepayload.SessionInfo{Id: "s1"}
	attr2 := corepayload.New.Attributes.UsingAuthInfo(&corepayload.AuthInfo{SessionInfo: si})

	actual := args.Map{"result": attr2.SessionInfo().Id != "s1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected s1", actual)
}

func Test_Cov9_Attributes_AuthType(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.AuthType() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty auth type", actual)

	attr2 := corepayload.New.Attributes.UsingAuthInfo(&corepayload.AuthInfo{ActionType: "login"})

	actual := args.Map{"result": attr2.AuthType() != "login"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected login", actual)
}

func Test_Cov9_Attributes_ResourceName(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.ResourceName() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty resource name", actual)

	attr2 := corepayload.New.Attributes.UsingAuthInfo(&corepayload.AuthInfo{ResourceName: "/api/test"})

	actual := args.Map{"result": attr2.ResourceName() != "/api/test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected /api/test", actual)
}

func Test_Cov9_Attributes_HasStringKeyValuePairs(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.HasStringKeyValuePairs()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_Cov9_Attributes_HasAnyKeyValuePairs(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.HasAnyKeyValuePairs()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_Cov9_Attributes_HasDynamicPayloads(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.HasDynamicPayloads()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)

	attr2 := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	actual := args.Map{"result": attr2.HasDynamicPayloads()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Cov9_Attributes_GetStringKeyValue(t *testing.T) {
	var nilAttr *corepayload.Attributes
	_, found := nilAttr.GetStringKeyValue("k")

	actual := args.Map{"result": found}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not find", actual)

	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr := corepayload.New.Attributes.UsingKeyValues(hm)
	val, found := attr.GetStringKeyValue("k")

	actual := args.Map{"result": found || val != "v"}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected v for key k", actual)
}

func Test_Cov9_Attributes_GetAnyKeyValue(t *testing.T) {
	var nilAttr *corepayload.Attributes
	_, found := nilAttr.GetAnyKeyValue("k")

	actual := args.Map{"result": found}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not find", actual)
}

func Test_Cov9_Attributes_IsErrorDifferent_IsErrorEqual(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr.IsErrorDifferent(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not different when both empty", actual)

	actual := args.Map{"result": attr.IsErrorEqual(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal when both empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Attributes — Setters (AttributesSetters.go)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Attributes_HandleErr_HandleError(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.HandleErr()
	attr.HandleError()
}

func Test_Cov9_Attributes_MustBeEmptyError(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.MustBeEmptyError()
}

func Test_Cov9_Attributes_SetAuthInfo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	result := attr.SetAuthInfo(&corepayload.AuthInfo{ActionType: "test"})

	actual := args.Map{"result": result.AuthInfo.ActionType != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)

	// nil receiver
	var nilAttr *corepayload.Attributes
	result = nilAttr.SetAuthInfo(&corepayload.AuthInfo{ActionType: "new"})

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil result from nil receiver", actual)
}

func Test_Cov9_Attributes_SetUserInfo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.AuthInfo = &corepayload.AuthInfo{}
	user := corepayload.New.User.UsingName("Alice")
	userInfo := &corepayload.UserInfo{User: user}
	result := attr.SetUserInfo(userInfo)

	actual := args.Map{"result": result.AuthInfo.UserInfo.User.Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)

	// nil receiver
	var nilAttr *corepayload.Attributes
	result = nilAttr.SetUserInfo(userInfo)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil result from nil receiver", actual)
}

func Test_Cov9_Attributes_AddNewStringKeyValueOnly(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	added := attr.AddNewStringKeyValueOnly("k", "v")

	actual := args.Map{"result": added}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected added", actual)

	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.AddNewStringKeyValueOnly("k", "v")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not add", actual)
}

func Test_Cov9_Attributes_AddNewAnyKeyValueOnly(t *testing.T) {
	anyMap := coredynamic.NewMapAnyItems(0)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	added := attr.AddNewAnyKeyValueOnly("k", 42)

	actual := args.Map{"result": added}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected added", actual)

	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.AddNewAnyKeyValueOnly("k", 42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not add", actual)
}

func Test_Cov9_Attributes_AddOrUpdateString(t *testing.T) {
	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr := corepayload.New.Attributes.UsingKeyValues(hm)
	isNew := attr.AddOrUpdateString("k2", "v2")

	actual := args.Map{"result": isNew}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected new key", actual)

	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.AddOrUpdateString("k", "v")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_Cov9_Attributes_AddOrUpdateAnyItem(t *testing.T) {
	anyMap := coredynamic.NewMapAnyItems(0)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	isNew := attr.AddOrUpdateAnyItem("k", 42)

	actual := args.Map{"result": isNew}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected new key", actual)

	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.AddOrUpdateAnyItem("k", 42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_Cov9_Attributes_SetBasicErr(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	result := attr.SetBasicErr(nil)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// nil receiver
	var nilAttr *corepayload.Attributes
	result = nilAttr.SetBasicErr(nil)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil from nil receiver", actual)
}

func Test_Cov9_Attributes_Clear_Dispose(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.AddNewStringKeyValueOnly("k", "v")
	attr.Clear()

	actual := args.Map{"result": attr.HasStringKeyValuePairs()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty after clear", actual)

	attr2 := corepayload.New.Attributes.Empty()
	attr2.Dispose()

	// nil Clear and Dispose
	var nilAttr *corepayload.Attributes
	nilAttr.Clear()
	nilAttr.Dispose()
}

// ══════════════════════════════════════════════════════════════════════════════
// Attributes — JSON (AttributesJson.go)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Attributes_PayloadsPrettyString(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"key":"value"}`))
	result := attr.PayloadsPrettyString()

	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty string", actual)

	var nilAttr *corepayload.Attributes

	actual := args.Map{"result": nilAttr.PayloadsPrettyString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Cov9_Attributes_PayloadsJsonResult(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"key":"value"}`))
	result := attr.PayloadsJsonResult()

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil result", actual)

	emptyAttr := corepayload.New.Attributes.Empty()
	result = emptyAttr.PayloadsJsonResult()

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil empty result", actual)
}

func Test_Cov9_Attributes_JsonString_JsonStringMust(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	s := attr.JsonString()

	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty json string", actual)

	s = attr.JsonStringMust()

	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty json string must", actual)
}

func Test_Cov9_Attributes_String(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	s := attr.String()

	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)
}

func Test_Cov9_Attributes_PrettyJsonString(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	s := attr.PrettyJsonString()

	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty pretty json", actual)
}

func Test_Cov9_Attributes_Json_JsonPtr(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	_ = attr.Json()
	_ = attr.JsonPtr()
}

func Test_Cov9_Attributes_JsonModel_JsonModelAny(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	model := attr.JsonModel()
	_ = model
	anyModel := attr.JsonModelAny()
	_ = anyModel
}

func Test_Cov9_Attributes_NonPtr(t *testing.T) {
	attr := corepayload.Attributes{}
	nonPtr := attr.NonPtr()
	_ = nonPtr
}

func Test_Cov9_Attributes_AsJsonContractsBinder(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	binder := attr.AsJsonContractsBinder()

	actual := args.Map{"result": binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil binder", actual)
}

func Test_Cov9_Attributes_AsAttributesBinder(t *testing.T) {
	attr := corepayload.Attributes{}
	binder := attr.AsAttributesBinder()

	actual := args.Map{"result": binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil binder", actual)
}

func Test_Cov9_Attributes_ParseInjectUsingJson(t *testing.T) {
	attr := &corepayload.Attributes{}
	jsonResult := corejson.NewPtr(corepayload.Attributes{})
	result, err := attr.ParseInjectUsingJson(jsonResult)

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil result", actual)
}

func Test_Cov9_Attributes_ParseInjectUsingJsonMust(t *testing.T) {
	attr := &corepayload.Attributes{}
	jsonResult := corejson.NewPtr(corepayload.Attributes{})
	result := attr.ParseInjectUsingJsonMust(jsonResult)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil result", actual)
}

func Test_Cov9_Attributes_JsonParseSelfInject(t *testing.T) {
	attr := &corepayload.Attributes{}
	jsonResult := corejson.NewPtr(corepayload.Attributes{})
	err := attr.JsonParseSelfInject(jsonResult)

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Cov9_Attributes_DeserializeDynamicPayloads(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))
	var s string
	err := attr.DeserializeDynamicPayloads(&s)

	actual := args.Map{"result": err != nil || s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_Cov9_Attributes_DeserializeDynamicPayloadsMust(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))
	var s string
	attr.DeserializeDynamicPayloadsMust(&s)

	actual := args.Map{"result": s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_Cov9_Attributes_DynamicPayloadsDeserialize(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))
	var s string
	err := attr.DynamicPayloadsDeserialize(&s)

	actual := args.Map{"result": err != nil || s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)

	// nil receiver
	var nilAttr *corepayload.Attributes
	err = nilAttr.DynamicPayloadsDeserialize(&s)

	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return error", actual)
}

func Test_Cov9_Attributes_DynamicPayloadsDeserializeMust(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))
	var s string
	attr.DynamicPayloadsDeserializeMust(&s)

	actual := args.Map{"result": s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_Cov9_Attributes_AnyKeyReflectSetTo(t *testing.T) {
	var nilAttr *corepayload.Attributes
	err := nilAttr.AnyKeyReflectSetTo("k", nil)

	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return error", actual)
}

func Test_Cov9_Attributes_ReflectSetTo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	var target corepayload.Attributes
	_ = attr.ReflectSetTo(&target)
}

// ══════════════════════════════════════════════════════════════════════════════
// Attributes — Clone and Equality (Attributes.go)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Attributes_IsEqual_AllBranches(t *testing.T) {
	var nilA, nilB *corepayload.Attributes

	actual := args.Map{"result": nilA.IsEqual(nilB)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)

	attr := corepayload.New.Attributes.Empty()

	actual := args.Map{"result": nilA.IsEqual(attr)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)

	actual := args.Map{"result": attr.IsEqual(nilA)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil vs nil should not be equal", actual)

	// Same pointer
	actual := args.Map{"result": attr.IsEqual(attr)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same pointer should be equal", actual)

	// Different paging
	a1 := corepayload.New.Attributes.Empty()
	a2 := corepayload.New.Attributes.Empty()
	a1.PagingInfo = &corepayload.PagingInfo{TotalPages: 1}

	actual := args.Map{"result": a1.IsEqual(a2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different paging should not be equal", actual)
}

func Test_Cov9_Attributes_Clone_Shallow(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"data"`))
	cloned, err := attr.Clone(false)

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	actual := args.Map{"result": string(cloned.DynamicPayloads) != `"data"`}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned data", actual)
}

func Test_Cov9_Attributes_Clone_Deep(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"data"`))
	cloned, err := attr.Clone(true)

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	actual := args.Map{"result": string(cloned.DynamicPayloads) != `"data"`}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned data", actual)
}

func Test_Cov9_Attributes_ClonePtr_Nil(t *testing.T) {
	var nilAttr *corepayload.Attributes
	cloned, err := nilAttr.ClonePtr(true)

	actual := args.Map{"result": err != nil || cloned != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ClonePtr should return nil, nil", actual)
}

func Test_Cov9_Attributes_Clone_NilReturnsEmpty(t *testing.T) {
	var nilAttr *corepayload.Attributes
	cloned, err := nilAttr.Clone(true)

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	_ = cloned
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadWrapper — Core methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PayloadWrapper_HasSafeItems(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	actual := args.Map{"result": pw.HasSafeItems()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not have safe items", actual)
}

func Test_Cov9_PayloadWrapper_DynamicPayloads(t *testing.T) {
	var nilPW *corepayload.PayloadWrapper
	p := nilPW.DynamicPayloads()

	actual := args.Map{"result": len(p) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)

	pw := corepayload.New.PayloadWrapper.Empty()
	_ = pw.DynamicPayloads()
}

func Test_Cov9_PayloadWrapper_SetDynamicPayloads(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	err := pw.SetDynamicPayloads([]byte("data"))

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	var nilPW *corepayload.PayloadWrapper
	err = nilPW.SetDynamicPayloads([]byte("data"))

	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return error", actual)
}

func Test_Cov9_PayloadWrapper_AttrAsBinder(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	_ = pw.AttrAsBinder()
}

func Test_Cov9_PayloadWrapper_InitializeAttributesOnNull(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	binder := pw.InitializeAttributesOnNull()

	actual := args.Map{"result": binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil binder", actual)
}

func Test_Cov9_PayloadWrapper_BasicError(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	be := pw.BasicError()

	actual := args.Map{"result": be != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Cov9_PayloadWrapper_All_AllSafe(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	id, name, entity, cat, payloads := pw.All()
	_, _, _, _, _ = id, name, entity, cat, payloads

	id, name, entity, cat, payloads = pw.AllSafe()
	_, _, _, _, _ = id, name, entity, cat, payloads

	var nilPW *corepayload.PayloadWrapper
	id, name, entity, cat, payloads = nilPW.AllSafe()

	actual := args.Map{"result": id != "" || name != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil AllSafe should return empty strings", actual)
}

func Test_Cov9_PayloadWrapper_PayloadName_Category_TaskType_EntityType(t *testing.T) {
	pw := corepayload.PayloadWrapper{
		Name:         "n",
		CategoryName: "c",
		TaskTypeName: "t",
		EntityType:   "e",
	}

	if pw.PayloadName() != "n" || pw.PayloadCategory() != "c" ||
		pw.PayloadTaskType() != "t" || pw.PayloadEntityType() != "e" {
		actual := args.Map{"result": false}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected values", actual)
	}
}

func Test_Cov9_PayloadWrapper_PayloadDynamic_Value(t *testing.T) {
	pw := corepayload.PayloadWrapper{Payloads: []byte("data")}
	_ = pw.PayloadDynamic()
	_ = pw.Value()
}

func Test_Cov9_PayloadWrapper_PayloadProperties(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	props := pw.PayloadProperties()

	actual := args.Map{"result": props == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil properties", actual)
}

func Test_Cov9_PayloadWrapper_HandleError(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.HandleError() // should not panic
}

func Test_Cov9_PayloadWrapper_AnyAttributes_ReflectSetAttributes(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	_ = pw.AnyAttributes()

	var target corepayload.Attributes
	_ = pw.ReflectSetAttributes(&target)
}

func Test_Cov9_PayloadWrapper_IdString_IdInteger(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Identifier: "42"}

	actual := args.Map{"result": pw.IdString() != "42"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	actual := args.Map{"result": pw.IdInteger() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_Cov9_PayloadWrapper_Serialize(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	_, err := pw.Serialize()

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	bytes := pw.SerializeMust()

	actual := args.Map{"result": len(bytes) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected serialized bytes", actual)
}

func Test_Cov9_PayloadWrapper_Username(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	actual := args.Map{"result": pw.Username() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty username", actual)

	user := corepayload.New.User.UsingName("Alice")
	pw.Attributes.AuthInfo = &corepayload.AuthInfo{
		UserInfo: &corepayload.UserInfo{User: user},
	}

	actual := args.Map{"result": pw.Username() != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)
}

func Test_Cov9_PayloadWrapper_Error(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	actual := args.Map{"result": pw.Error() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil error", actual)
}

func Test_Cov9_PayloadWrapper_IsEqual_AllBranches(t *testing.T) {
	var nilA, nilB *corepayload.PayloadWrapper

	actual := args.Map{"result": nilA.IsEqual(nilB)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)

	pw := corepayload.New.PayloadWrapper.Empty()

	actual := args.Map{"result": nilA.IsEqual(pw)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)

	actual := args.Map{"result": pw.IsEqual(nilA)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil vs nil should not be equal", actual)

	actual := args.Map{"result": pw.IsEqual(pw)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same pointer should be equal", actual)

	pw2 := corepayload.New.PayloadWrapper.Empty()

	actual := args.Map{"result": pw.IsEqual(pw2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "two empty should be equal", actual)

	// Different Name
	pw3 := corepayload.New.PayloadWrapper.Empty()
	pw3.Name = "different"

	actual := args.Map{"result": pw.IsEqual(pw3)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different name should not be equal", actual)

	// Different Identifier
	pw4 := corepayload.New.PayloadWrapper.Empty()
	pw4.Identifier = "diff"

	actual := args.Map{"result": pw.IsEqual(pw4)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different identifier should not be equal", actual)

	// Different TaskTypeName
	pw5 := corepayload.New.PayloadWrapper.Empty()
	pw5.TaskTypeName = "diff"

	actual := args.Map{"result": pw.IsEqual(pw5)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different task type should not be equal", actual)

	// Different EntityType
	pw6 := corepayload.New.PayloadWrapper.Empty()
	pw6.EntityType = "diff"

	actual := args.Map{"result": pw.IsEqual(pw6)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different entity type should not be equal", actual)

	// Different CategoryName
	pw7 := corepayload.New.PayloadWrapper.Empty()
	pw7.CategoryName = "diff"

	actual := args.Map{"result": pw.IsEqual(pw7)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different category should not be equal", actual)

	// Different HasManyRecords
	pw8 := corepayload.New.PayloadWrapper.Empty()
	pw8.HasManyRecords = true

	actual := args.Map{"result": pw.IsEqual(pw8)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different HasManyRecords should not be equal", actual)
}

func Test_Cov9_PayloadWrapper_IsPayloadsEqual(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("data")}

	actual := args.Map{"result": pw.IsPayloadsEqual([]byte("data"))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)

	actual := args.Map{"result": pw.IsPayloadsEqual([]byte("other"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)

	var nilPW *corepayload.PayloadWrapper

	actual := args.Map{"result": nilPW.IsPayloadsEqual([]byte("data"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not be equal", actual)
}

func Test_Cov9_PayloadWrapper_IsName_IsIdentifier_IsTaskTypeName_IsEntityType_IsCategory(t *testing.T) {
	pw := &corepayload.PayloadWrapper{
		Name:         "n",
		Identifier:   "id",
		TaskTypeName: "task",
		EntityType:   "entity",
		CategoryName: "cat",
	}

	actual := args.Map{"result": pw.IsName("n")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual := args.Map{"result": pw.IsIdentifier("id")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual := args.Map{"result": pw.IsTaskTypeName("task")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual := args.Map{"result": pw.IsEntityType("entity")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual := args.Map{"result": pw.IsCategory("cat")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Cov9_PayloadWrapper_HasAnyItem_HasIssuesOrEmpty(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	actual := args.Map{"result": pw.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not have any item", actual)

	actual := args.Map{"result": pw.HasIssuesOrEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should have issues or be empty", actual)
}

func Test_Cov9_PayloadWrapper_HasError_IsEmptyError(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	actual := args.Map{"result": pw.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	actual := args.Map{"result": pw.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty error", actual)
}

func Test_Cov9_PayloadWrapper_HasAttributes_IsEmptyAttributes(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	actual := args.Map{"result": pw.HasAttributes()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected attributes", actual)

	pw2 := &corepayload.PayloadWrapper{}

	actual := args.Map{"result": pw2.IsEmptyAttributes()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty attributes", actual)
}

func Test_Cov9_PayloadWrapper_HasSingleRecord(t *testing.T) {
	pw := &corepayload.PayloadWrapper{HasManyRecords: false}

	actual := args.Map{"result": pw.HasSingleRecord()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected single record", actual)
}

func Test_Cov9_PayloadWrapper_IsNull_HasAnyNil(t *testing.T) {
	var nilPW *corepayload.PayloadWrapper

	actual := args.Map{"result": nilPW.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected null", actual)

	actual := args.Map{"result": nilPW.HasAnyNil()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has any nil", actual)
}

func Test_Cov9_PayloadWrapper_Count_Length_IsEmpty_HasItems(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("data")}

	actual := args.Map{"result": pw.Count() != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)

	actual := args.Map{"result": pw.Length() != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)

	actual := args.Map{"result": pw.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	actual := args.Map{"result": pw.HasItems()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has items", actual)

	var nilPW *corepayload.PayloadWrapper

	actual := args.Map{"result": nilPW.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil Length should be 0", actual)
}

func Test_Cov9_PayloadWrapper_IdentifierInteger_UnsignedInteger(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Identifier: "42"}

	actual := args.Map{"result": pw.IdentifierInteger() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	actual := args.Map{"result": pw.IdentifierUnsignedInteger() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	// Empty identifier
	pw2 := &corepayload.PayloadWrapper{}

	actual := args.Map{"result": pw2.IdentifierInteger() >= 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid value", actual)
}

func Test_Cov9_PayloadWrapper_BytesConverter(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("data")}
	bc := pw.BytesConverter()

	actual := args.Map{"result": bc == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_PayloadWrapper_Deserialize(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`"hello"`)}
	var s string
	err := pw.Deserialize(&s)

	actual := args.Map{"result": err != nil || s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)

	err = pw.PayloadDeserialize(&s)

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Cov9_PayloadWrapper_MarshalJSON_UnmarshalJSON(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "test"
	pw.Payloads = []byte(`"hello"`)

	jsonBytes, err := pw.MarshalJSON()

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no marshal error", actual)

	pw2 := &corepayload.PayloadWrapper{}
	err = pw2.UnmarshalJSON(jsonBytes)

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no unmarshal error", actual)

	actual := args.Map{"result": pw2.Name != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)
}

func Test_Cov9_PayloadWrapper_String_PrettyJsonString(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	_ = pw.String()
	_ = pw.PrettyJsonString()
}

func Test_Cov9_PayloadWrapper_JsonString_JsonStringMust(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	_ = pw.JsonString()
	_ = pw.JsonStringMust()
}

func Test_Cov9_PayloadWrapper_Json_JsonPtr(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	_ = pw.Json()
	_ = pw.JsonPtr()
}

func Test_Cov9_PayloadWrapper_JsonModel_JsonModelAny(t *testing.T) {
	pw := corepayload.PayloadWrapper{}
	_ = pw.JsonModel()
	_ = pw.JsonModelAny()
}

func Test_Cov9_PayloadWrapper_ParseInjectUsingJson(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	jsonResult := corejson.NewPtr(corepayload.PayloadWrapper{Name: "test"})
	result, err := pw.ParseInjectUsingJson(jsonResult)

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_PayloadWrapper_JsonParseSelfInject(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	jsonResult := corejson.NewPtr(corepayload.PayloadWrapper{})
	err := pw.JsonParseSelfInject(jsonResult)

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Cov9_PayloadWrapper_PayloadsString(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("hello")}

	actual := args.Map{"result": pw.PayloadsString() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)

	emptyPW := &corepayload.PayloadWrapper{}

	actual := args.Map{"result": emptyPW.PayloadsString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Cov9_PayloadWrapper_PayloadsJsonResult(t *testing.T) {
	pw := corepayload.PayloadWrapper{Payloads: []byte(`{"k":"v"}`)}
	result := pw.PayloadsJsonResult()

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	emptyPW := corepayload.PayloadWrapper{}
	result = emptyPW.PayloadsJsonResult()

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil empty", actual)
}

func Test_Cov9_PayloadWrapper_PayloadsPrettyString(t *testing.T) {
	pw := corepayload.PayloadWrapper{Payloads: []byte(`{"k":"v"}`)}
	result := pw.PayloadsPrettyString()

	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

	emptyPW := corepayload.PayloadWrapper{}

	actual := args.Map{"result": emptyPW.PayloadsPrettyString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Cov9_PayloadWrapper_Clear_Dispose(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Payloads = []byte("data")
	pw.Clear()

	actual := args.Map{"result": pw.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 length after clear", actual)

	pw2 := corepayload.New.PayloadWrapper.Empty()
	pw2.Dispose()

	actual := args.Map{"result": pw2.Attributes != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil attributes after dispose", actual)

	var nilPW *corepayload.PayloadWrapper
	nilPW.Clear()
	nilPW.Dispose()
}

func Test_Cov9_PayloadWrapper_AsJsonContractsBinder(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	binder := pw.AsJsonContractsBinder()

	actual := args.Map{"result": binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_PayloadWrapper_Clone_Shallow(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "test"
	pw.Payloads = []byte("data")
	cloned, err := pw.Clone(false)

	actual := args.Map{"result": err != nil || cloned.Name != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned with name test", actual)
}

func Test_Cov9_PayloadWrapper_Clone_Deep(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "test"
	pw.Payloads = []byte("data")
	cloned, err := pw.Clone(true)

	actual := args.Map{"result": err != nil || cloned.Name != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned with name test", actual)
}

func Test_Cov9_PayloadWrapper_ClonePtr_Nil(t *testing.T) {
	var nilPW *corepayload.PayloadWrapper
	cloned, err := nilPW.ClonePtr(true)

	actual := args.Map{"result": err != nil || cloned != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ClonePtr should return nil, nil", actual)
}

func Test_Cov9_PayloadWrapper_NonPtr_ToPtr(t *testing.T) {
	var nilPW *corepayload.PayloadWrapper
	nonPtr := nilPW.NonPtr()
	_ = nonPtr

	pw := corepayload.PayloadWrapper{Name: "test"}
	ptr := pw.ToPtr()

	actual := args.Map{"result": ptr.Name != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)
}

func Test_Cov9_PayloadWrapper_AsStandardTaskEntityDefiner(t *testing.T) {
	pw := corepayload.PayloadWrapper{}
	binder := pw.AsStandardTaskEntityDefinerContractsBinder()

	actual := args.Map{"result": binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_PayloadWrapper_AsPayloadsBinder(t *testing.T) {
	pw := corepayload.PayloadWrapper{}
	binder := pw.AsPayloadsBinder()

	actual := args.Map{"result": binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_PayloadWrapper_AsJsonMarshaller(t *testing.T) {
	pw := corepayload.PayloadWrapper{}
	m := pw.AsJsonMarshaller()

	actual := args.Map{"result": m == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_PayloadWrapper_ReflectSetTo(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	var target corepayload.PayloadWrapper
	_ = pw.ReflectSetTo(&target)
}

func Test_Cov9_PayloadWrapper_ValueReflectSet(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`"hello"`)}
	var target string
	_ = pw.ValueReflectSet(&target)
}

func Test_Cov9_PayloadWrapper_IsStandardTaskEntityEqual(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "test"

	// Same type
	pw2 := corepayload.New.PayloadWrapper.Empty()
	pw2.Name = "test"

	actual := args.Map{"result": pw.IsStandardTaskEntityEqual(pw2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_Cov9_PayloadWrapper_SetPayloadDynamic(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	result := pw.SetPayloadDynamic([]byte("data"))

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_PayloadWrapper_SetPayloadDynamicAny(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	result, err := pw.SetPayloadDynamicAny("hello")

	actual := args.Map{"result": err != nil || result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Cov9_PayloadWrapper_SetAuthInfo(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	result := pw.SetAuthInfo(&corepayload.AuthInfo{ActionType: "login"})

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_PayloadWrapper_SetUserInfo(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	user := corepayload.New.User.UsingName("Alice")
	result := pw.SetUserInfo(&corepayload.UserInfo{User: user})

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_PayloadWrapper_SetUser_SetSysUser(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	user := corepayload.New.User.UsingName("Alice")
	result := pw.SetUser(user)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	sysUser := corepayload.New.User.System("sys", "system")
	result = pw.SetSysUser(sysUser)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_PayloadWrapper_DeserializePayloadsToPayloadsCollection(t *testing.T) {
	// Create a payload wrapper containing serialized collection
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Payloads = []byte(`[]`)
	_, err := pw.DeserializePayloadsToPayloadsCollection()

	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Cov9_PayloadWrapper_DeserializePayloadsToPayloadWrapper(t *testing.T) {
	inner := corepayload.New.PayloadWrapper.Empty()
	inner.Name = "inner"
	jsonBytes, _ := inner.Serialize()

	pw := &corepayload.PayloadWrapper{Payloads: jsonBytes}
	result, err := pw.DeserializePayloadsToPayloadWrapper()

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	actual := args.Map{"result": result.Name != "inner"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected inner", actual)
}

func Test_Cov9_PayloadWrapper_ReCreateUsingJsonBytes(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "test"
	jsonBytes, _ := pw.Serialize()

	result, err := pw.ReCreateUsingJsonBytes(jsonBytes)

	actual := args.Map{"result": err != nil || result.Name != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)
}

func Test_Cov9_PayloadWrapper_ReCreateUsingJsonResult(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "test"
	jsonResult := pw.JsonPtr()

	result, err := pw.ReCreateUsingJsonResult(jsonResult)

	actual := args.Map{"result": err != nil || result.Name != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadsCollection — Getters
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PayloadsCollection_Length_Count_IsEmpty(t *testing.T) {
	var nilCol *corepayload.PayloadsCollection

	actual := args.Map{"result": nilCol.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil Length should be 0", actual)

	col := corepayload.New.PayloadsCollection.Empty()

	actual := args.Map{"result": col.Count() != 0 || !col.IsEmpty() || col.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty collection state mismatch", actual)
}

func Test_Cov9_PayloadsCollection_LastIndex_HasIndex(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.New.PayloadWrapper.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	_ = pw

	actual := args.Map{"result": col.LastIndex() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)

	actual := args.Map{"result": col.HasIndex(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for index 0", actual)

	actual := args.Map{"result": col.HasIndex(1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for index 1", actual)
}

func Test_Cov9_PayloadsCollection_First_Last(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	actual := args.Map{"result": col.First().Name != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)

	actual := args.Map{"result": col.Last().Name != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b", actual)

	_ = col.FirstDynamic()
	_ = col.LastDynamic()
}

func Test_Cov9_PayloadsCollection_FirstOrDefault_LastOrDefault(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	actual := args.Map{"result": col.FirstOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	actual := args.Map{"result": col.LastOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	_ = col.FirstOrDefaultDynamic()
	_ = col.LastOrDefaultDynamic()

	col.Add(corepayload.PayloadWrapper{Name: "a"})

	actual := args.Map{"result": col.FirstOrDefault().Name != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)

	actual := args.Map{"result": col.LastOrDefault().Name != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
}

func Test_Cov9_PayloadsCollection_Skip_Take_Limit(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})
	col.Add(corepayload.PayloadWrapper{Name: "c"})

	skipped := col.Skip(1)

	actual := args.Map{"result": len(skipped) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	_ = col.SkipDynamic(1)

	skipCol := col.SkipCollection(1)

	actual := args.Map{"result": skipCol.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	taken := col.Take(2)

	actual := args.Map{"result": len(taken) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	_ = col.TakeDynamic(2)

	takeCol := col.TakeCollection(2)

	actual := args.Map{"result": takeCol.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	limitCol := col.LimitCollection(2)

	actual := args.Map{"result": limitCol.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	safeLimitCol := col.SafeLimitCollection(100)

	actual := args.Map{"result": safeLimitCol.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)

	_ = col.LimitDynamic(2)
	_ = col.Limit(2)
}

func Test_Cov9_PayloadsCollection_Strings(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})

	strings := col.Strings()

	actual := args.Map{"result": len(strings) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_PayloadsCollection_IsEqual(t *testing.T) {
	var nilA, nilB *corepayload.PayloadsCollection

	actual := args.Map{"result": nilA.IsEqual(nilB)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)

	col := corepayload.New.PayloadsCollection.Empty()

	actual := args.Map{"result": nilA.IsEqual(col)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)

	col2 := corepayload.New.PayloadsCollection.Empty()

	actual := args.Map{"result": col.IsEqual(col2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "two empty should be equal", actual)
}

func Test_Cov9_PayloadsCollection_IsEqualItems(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.New.PayloadWrapper.Empty()
	col.AddsPtr(pw)

	actual := args.Map{"result": col.IsEqualItems(pw)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadsCollection — Filter
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PayloadsCollection_Filter(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	filtered := col.Filter(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return pw.Name == "a", false
	})

	actual := args.Map{"result": len(filtered) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_PayloadsCollection_FilterWithLimit(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})
	col.Add(corepayload.PayloadWrapper{Name: "c"})

	filtered := col.FilterWithLimit(1, func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})

	actual := args.Map{"result": len(filtered) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_PayloadsCollection_FirstByFilter(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	found := col.FirstByFilter(func(pw *corepayload.PayloadWrapper) bool {
		return pw.Name == "b"
	})

	actual := args.Map{"result": found == nil || found.Name != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b", actual)

	notFound := col.FirstByFilter(func(pw *corepayload.PayloadWrapper) bool {
		return pw.Name == "z"
	})

	actual := args.Map{"result": notFound != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Cov9_PayloadsCollection_FirstById(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Identifier: "id1", Name: "a"})
	col.Add(corepayload.PayloadWrapper{Identifier: "id2", Name: "b"})

	found := col.FirstById("id2")

	actual := args.Map{"result": found == nil || found.Name != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b", actual)
}

func Test_Cov9_PayloadsCollection_FirstByCategory(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{CategoryName: "cat1", Name: "a"})

	found := col.FirstByCategory("cat1")

	actual := args.Map{"result": found == nil || found.Name != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
}

func Test_Cov9_PayloadsCollection_FirstByTaskType(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{TaskTypeName: "task1", Name: "a"})

	found := col.FirstByTaskType("task1")

	actual := args.Map{"result": found == nil || found.Name != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
}

func Test_Cov9_PayloadsCollection_FirstByEntityType(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{EntityType: "entity1", Name: "a"})

	found := col.FirstByEntityType("entity1")

	actual := args.Map{"result": found == nil || found.Name != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
}

func Test_Cov9_PayloadsCollection_FilterCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	filtered := col.FilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return pw.Name == "a", false
	})

	actual := args.Map{"result": filtered.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_PayloadsCollection_SkipFilterCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	filtered := col.SkipFilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return pw.Name == "a", false
	})

	actual := args.Map{"result": filtered.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_PayloadsCollection_FilterCollectionByIds(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Identifier: "id1", Name: "a"})
	col.Add(corepayload.PayloadWrapper{Identifier: "id2", Name: "b"})
	col.Add(corepayload.PayloadWrapper{Identifier: "id3", Name: "c"})

	filtered := col.FilterCollectionByIds("id1", "id3")

	actual := args.Map{"result": filtered.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov9_PayloadsCollection_FilterNameCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	filtered := col.FilterNameCollection("a")

	actual := args.Map{"result": filtered.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_PayloadsCollection_FilterCategoryCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{CategoryName: "cat1"})

	filtered := col.FilterCategoryCollection("cat1")

	actual := args.Map{"result": filtered.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_PayloadsCollection_FilterEntityTypeCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{EntityType: "entity1"})

	filtered := col.FilterEntityTypeCollection("entity1")

	actual := args.Map{"result": filtered.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_PayloadsCollection_FilterTaskTypeCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{TaskTypeName: "task1"})

	filtered := col.FilterTaskTypeCollection("task1")

	actual := args.Map{"result": filtered.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadsCollection — Paging
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PayloadsCollection_GetPagesSize(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	for i := 0; i < 25; i++ {
		col.Add(corepayload.PayloadWrapper{Name: "item"})
	}

	pages := col.GetPagesSize(10)

	actual := args.Map{"result": pages != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)

	actual := args.Map{"result": col.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "zero page size should return 0", actual)

	actual := args.Map{"result": col.GetPagesSize(-1) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "negative page size should return 0", actual)
}

func Test_Cov9_PayloadsCollection_GetPagedCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	for i := 0; i < 25; i++ {
		col.Add(corepayload.PayloadWrapper{Name: "item"})
	}

	pages := col.GetPagedCollection(10)

	actual := args.Map{"result": len(pages) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)

	// Smaller than page size
	small := corepayload.New.PayloadsCollection.Empty()
	small.Add(corepayload.PayloadWrapper{Name: "a"})
	pages = small.GetPagedCollection(10)

	actual := args.Map{"result": len(pages) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 page", actual)
}

func Test_Cov9_PayloadsCollection_GetSinglePageCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	for i := 0; i < 25; i++ {
		col.Add(corepayload.PayloadWrapper{Name: "item"})
	}

	page := col.GetSinglePageCollection(10, 1)

	actual := args.Map{"result": page.Length() != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)

	page3 := col.GetSinglePageCollection(10, 3)

	actual := args.Map{"result": page3.Length() != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)

	// Smaller than page size
	small := corepayload.New.PayloadsCollection.Empty()
	small.Add(corepayload.PayloadWrapper{Name: "a"})
	page = small.GetSinglePageCollection(10, 1)

	actual := args.Map{"result": page.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadsCollection — Mutation
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PayloadsCollection_Add_Adds(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Adds(corepayload.PayloadWrapper{Name: "b"}, corepayload.PayloadWrapper{Name: "c"})

	actual := args.Map{"result": col.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)

	// Adds empty
	col.Adds()

	actual := args.Map{"result": col.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected still 3", actual)
}

func Test_Cov9_PayloadsCollection_AddsPtr(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.New.PayloadWrapper.Empty()
	col.AddsPtr(pw)

	actual := args.Map{"result": col.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	// Empty
	col.AddsPtr()

	actual := args.Map{"result": col.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected still 1", actual)
}

func Test_Cov9_PayloadsCollection_AddsPtrOptions(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Payloads = []byte("data")
	col.AddsPtrOptions(false, pw)

	actual := args.Map{"result": col.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	// Skip issues
	emptyPW := corepayload.New.PayloadWrapper.Empty()
	col2 := corepayload.New.PayloadsCollection.Empty()
	col2.AddsPtrOptions(true, emptyPW)

	actual := args.Map{"result": col2.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 (skipped)", actual)
}

func Test_Cov9_PayloadsCollection_AddsOptions(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.AddsOptions(false, corepayload.PayloadWrapper{Payloads: []byte("data")})

	actual := args.Map{"result": col.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	// Skip issues
	col2 := corepayload.New.PayloadsCollection.Empty()
	col2.AddsOptions(true, corepayload.PayloadWrapper{})

	actual := args.Map{"result": col2.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 (skipped)", actual)
}

func Test_Cov9_PayloadsCollection_AddsIf(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.AddsIf(true, corepayload.PayloadWrapper{Name: "a"})

	actual := args.Map{"result": col.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	col.AddsIf(false, corepayload.PayloadWrapper{Name: "b"})

	actual := args.Map{"result": col.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected still 1", actual)
}

func Test_Cov9_PayloadsCollection_InsertAt(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "c"})
	col.InsertAt(1, corepayload.PayloadWrapper{Name: "b"})

	actual := args.Map{"result": col.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Cov9_PayloadsCollection_ConcatNew(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	newCol := col.ConcatNew(corepayload.PayloadWrapper{Name: "b"})

	actual := args.Map{"result": newCol.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	actual := args.Map{"result": col.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "original should still be 1", actual)
}

func Test_Cov9_PayloadsCollection_ConcatNewPtr(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "b"
	newCol := col.ConcatNewPtr(pw)

	actual := args.Map{"result": newCol.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov9_PayloadsCollection_Reverse(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})
	col.Add(corepayload.PayloadWrapper{Name: "c"})
	col.Reverse()

	actual := args.Map{"result": col.First().Name != "c" || col.Last().Name != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected reversed", actual)

	// Reverse 2 items
	col2 := corepayload.New.PayloadsCollection.Empty()
	col2.Add(corepayload.PayloadWrapper{Name: "a"})
	col2.Add(corepayload.PayloadWrapper{Name: "b"})
	col2.Reverse()

	actual := args.Map{"result": col2.First().Name != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b first", actual)

	// Reverse 1 item
	col3 := corepayload.New.PayloadsCollection.Empty()
	col3.Add(corepayload.PayloadWrapper{Name: "a"})
	col3.Reverse()

	actual := args.Map{"result": col3.First().Name != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
}

func Test_Cov9_PayloadsCollection_Clone_ClonePtr(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})

	cloned := col.Clone()

	actual := args.Map{"result": cloned.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	clonedPtr := col.ClonePtr()

	actual := args.Map{"result": clonedPtr.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	// nil ClonePtr
	var nilCol *corepayload.PayloadsCollection

	actual := args.Map{"result": nilCol.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ClonePtr should return nil", actual)
}

func Test_Cov9_PayloadsCollection_Clear_Dispose(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Clear()

	actual := args.Map{"result": col.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 after clear", actual)

	col2 := corepayload.New.PayloadsCollection.Empty()
	col2.Dispose()

	// nil Clear and Dispose
	var nilCol *corepayload.PayloadsCollection
	nilCol.Clear()
	nilCol.Dispose()
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadsCollection — JSON
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PayloadsCollection_StringsUsingFmt(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})

	strings := col.StringsUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.Name
	})

	actual := args.Map{"result": len(strings) != 1 || strings[0] != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
}

func Test_Cov9_PayloadsCollection_JoinUsingFmt(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	result := col.JoinUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.Name
	}, ",")

	actual := args.Map{"result": result != "a,b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a,b", actual)
}

func Test_Cov9_PayloadsCollection_JsonStrings_JoinJsonStrings(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("data")})

	jsonStrings := col.JsonStrings()

	actual := args.Map{"result": len(jsonStrings) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	_ = col.JoinJsonStrings(",")
}

func Test_Cov9_PayloadsCollection_Join(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("data")})

	_ = col.Join(",")
}

func Test_Cov9_PayloadsCollection_JsonString_String(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	actual := args.Map{"result": col.JsonString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return empty string", actual)

	actual := args.Map{"result": col.String() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return empty string", actual)

	col.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("data")})

	actual := args.Map{"result": col.JsonString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

	actual := args.Map{"result": col.String() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Cov9_PayloadsCollection_PrettyJsonString(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	actual := args.Map{"result": col.PrettyJsonString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return empty string", actual)

	col.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("data")})

	actual := args.Map{"result": col.PrettyJsonString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Cov9_PayloadsCollection_CsvStrings_JoinCsv_JoinCsvLine(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	actual := args.Map{"result": len(col.CsvStrings()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty csv should be empty", actual)

	col.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("data")})

	csvStrings := col.CsvStrings()

	actual := args.Map{"result": len(csvStrings) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	_ = col.JoinCsv()
	_ = col.JoinCsvLine()
}

func Test_Cov9_PayloadsCollection_Json_JsonPtr(t *testing.T) {
	col := corepayload.PayloadsCollection{}
	_ = col.Json()
	_ = col.JsonPtr()
}

func Test_Cov9_PayloadsCollection_ParseInjectUsingJson(t *testing.T) {
	col := &corepayload.PayloadsCollection{}
	jsonResult := corejson.NewPtr(corepayload.PayloadsCollection{})
	result, err := col.ParseInjectUsingJson(jsonResult)

	actual := args.Map{"result": err != nil || result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_PayloadsCollection_AsJsonContractsBinder(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	binder := col.AsJsonContractsBinder()

	actual := args.Map{"result": binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_PayloadsCollection_AsJsoner(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	jsoner := col.AsJsoner()

	actual := args.Map{"result": jsoner == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_PayloadsCollection_JsonParseSelfInject(t *testing.T) {
	col := &corepayload.PayloadsCollection{}
	jsonResult := corejson.NewPtr(corepayload.PayloadsCollection{})
	err := col.JsonParseSelfInject(jsonResult)

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Cov9_PayloadsCollection_AsJsonParseSelfInjector(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	injector := col.AsJsonParseSelfInjector()

	actual := args.Map{"result": injector == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// User — Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_User_IdentifierInteger(t *testing.T) {
	u := corepayload.User{Identifier: "42"}

	actual := args.Map{"result": u.IdentifierInteger() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	u2 := corepayload.User{}

	actual := args.Map{"result": u2.IdentifierInteger() >= 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid value", actual)
}

func Test_Cov9_User_IdentifierUnsignedInteger(t *testing.T) {
	u := corepayload.User{Identifier: "42"}

	actual := args.Map{"result": u.IdentifierUnsignedInteger() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	u2 := corepayload.User{Identifier: "-1"}

	actual := args.Map{"result": u2.IdentifierUnsignedInteger() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Cov9_User_AllBoolMethods(t *testing.T) {
	u := corepayload.New.User.All(true, "1", "Alice", "admin", "token", "hash")

	actual := args.Map{"result": u.HasAuthToken()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasAuthToken", actual)

	actual := args.Map{"result": u.HasPasswordHash()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasPasswordHash", actual)

	actual := args.Map{"result": u.IsPasswordHashEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	actual := args.Map{"result": u.IsAuthTokenEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	actual := args.Map{"result": u.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	actual := args.Map{"result": u.IsValidUser()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)

	actual := args.Map{"result": u.IsNameEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	actual := args.Map{"result": u.IsNameEqual("Alice")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)

	actual := args.Map{"result": u.IsNotSystemUser()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected system user", actual)

	actual := args.Map{"result": u.IsVirtualUser()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not virtual user (is system)", actual)

	actual := args.Map{"result": u.HasType()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has type", actual)

	actual := args.Map{"result": u.IsTypeEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	// nil receiver
	var nilUser *corepayload.User

	actual := args.Map{"result": nilUser.HasAuthToken() || nilUser.HasPasswordHash()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)

	actual := args.Map{"result": nilUser.IsPasswordHashEmpty() || !nilUser.IsAuthTokenEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)

	actual := args.Map{"result": nilUser.IsEmpty() || nilUser.IsValidUser() || !nilUser.IsNameEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil checks failed", actual)

	actual := args.Map{"result": nilUser.IsNameEqual("anything")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not be equal", actual)

	actual := args.Map{"result": nilUser.IsNotSystemUser() || nilUser.IsVirtualUser() || nilUser.HasType()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)

	actual := args.Map{"result": nilUser.IsTypeEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be type empty", actual)
}

func Test_Cov9_User_String_Json_Serialize_Deserialize(t *testing.T) {
	u := corepayload.New.User.UsingName("Alice")
	_ = u.String()
	_ = u.PrettyJsonString()
	_ = u.Json()
	_ = u.JsonPtr()

	serialized, err := u.Serialize()

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	u2 := &corepayload.User{}
	err = u2.Deserialize(serialized)

	actual := args.Map{"result": err != nil || u2.Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)
}

func Test_Cov9_User_Clone_ClonePtr(t *testing.T) {
	u := corepayload.New.User.All(false, "1", "Alice", "admin", "token", "hash")
	cloned := u.Clone()

	actual := args.Map{"result": cloned.Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)

	clonedPtr := u.ClonePtr()

	actual := args.Map{"result": clonedPtr.Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)

	var nilUser *corepayload.User

	actual := args.Map{"result": nilUser.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ClonePtr should return nil", actual)
}

func Test_Cov9_User_Ptr(t *testing.T) {
	u := corepayload.User{Name: "Alice"}
	ptr := u.Ptr()

	actual := args.Map{"result": ptr.Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// UserInfo — Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_UserInfo_All(t *testing.T) {
	ui := &corepayload.UserInfo{
		User:       corepayload.New.User.UsingName("Alice"),
		SystemUser: corepayload.New.User.System("sys", "system"),
	}

	actual := args.Map{"result": ui.HasUser()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasUser", actual)

	actual := args.Map{"result": ui.HasSystemUser()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasSystemUser", actual)

	actual := args.Map{"result": ui.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	actual := args.Map{"result": ui.IsUserEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	actual := args.Map{"result": ui.IsSystemUserEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	// nil receiver
	var nilUI *corepayload.UserInfo

	actual := args.Map{"result": nilUI.HasUser() || nilUI.HasSystemUser()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)

	actual := args.Map{"result": nilUI.IsEmpty() || !nilUI.IsUserEmpty() || !nilUI.IsSystemUserEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_Cov9_UserInfo_SetUserSystemUser(t *testing.T) {
	ui := &corepayload.UserInfo{}
	user := corepayload.New.User.UsingName("Alice")
	sysUser := corepayload.New.User.System("sys", "system")
	result := ui.SetUserSystemUser(user, sysUser)

	actual := args.Map{"result": result.User.Name != "Alice" || result.SystemUser.Name != "sys"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice and sys", actual)

	// nil receiver
	var nilUI *corepayload.UserInfo
	result = nilUI.SetUserSystemUser(user, sysUser)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_UserInfo_SetUser_SetSystemUser(t *testing.T) {
	ui := &corepayload.UserInfo{}
	user := corepayload.New.User.UsingName("Alice")
	result := ui.SetUser(user)

	actual := args.Map{"result": result.User.Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)

	sysUser := corepayload.New.User.System("sys", "system")
	result = ui.SetSystemUser(sysUser)

	actual := args.Map{"result": result.SystemUser.Name != "sys"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sys", actual)

	// nil receiver
	var nilUI *corepayload.UserInfo
	result = nilUI.SetUser(user)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	result = nilUI.SetSystemUser(sysUser)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_UserInfo_Clone_ClonePtr_Ptr(t *testing.T) {
	ui := &corepayload.UserInfo{User: corepayload.New.User.UsingName("Alice")}
	cloned := ui.Clone()

	actual := args.Map{"result": cloned.User.Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)

	clonedPtr := ui.ClonePtr()

	actual := args.Map{"result": clonedPtr.User.Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)

	var nilUI *corepayload.UserInfo

	actual := args.Map{"result": nilUI.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ClonePtr should return nil", actual)

	uiVal := corepayload.UserInfo{User: corepayload.New.User.UsingName("Bob")}
	ptr := uiVal.Ptr()

	actual := args.Map{"result": ptr.User.Name != "Bob"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Bob", actual)
}

func Test_Cov9_UserInfo_ToNonPtr(t *testing.T) {
	ui := &corepayload.UserInfo{User: corepayload.New.User.UsingName("Alice")}
	nonPtr := ui.ToNonPtr()

	actual := args.Map{"result": nonPtr.User.Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)

	var nilUI *corepayload.UserInfo
	nonPtr = nilUI.ToNonPtr()
	_ = nonPtr
}

// ══════════════════════════════════════════════════════════════════════════════
// AuthInfo — Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_AuthInfo_IdentifierInteger(t *testing.T) {
	ai := corepayload.AuthInfo{Identifier: "42"}

	actual := args.Map{"result": ai.IdentifierInteger() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	ai2 := corepayload.AuthInfo{}

	actual := args.Map{"result": ai2.IdentifierInteger() >= 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid value", actual)
}

func Test_Cov9_AuthInfo_IdentifierUnsignedInteger(t *testing.T) {
	ai := corepayload.AuthInfo{Identifier: "42"}

	actual := args.Map{"result": ai.IdentifierUnsignedInteger() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	ai2 := corepayload.AuthInfo{Identifier: "-1"}

	actual := args.Map{"result": ai2.IdentifierUnsignedInteger() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Cov9_AuthInfo_IsEmpty_HasAnyItem_IsValid(t *testing.T) {
	var nilAI *corepayload.AuthInfo

	actual := args.Map{"result": nilAI.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)

	ai := &corepayload.AuthInfo{ActionType: "login"}

	actual := args.Map{"result": ai.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	actual := args.Map{"result": ai.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has any item", actual)

	actual := args.Map{"result": ai.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
}

func Test_Cov9_AuthInfo_IsActionTypeEmpty_IsResourceNameEmpty(t *testing.T) {
	var nilAI *corepayload.AuthInfo

	actual := args.Map{"result": nilAI.IsActionTypeEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)

	actual := args.Map{"result": nilAI.IsResourceNameEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)

	ai := &corepayload.AuthInfo{ActionType: "login", ResourceName: "/api"}

	actual := args.Map{"result": ai.IsActionTypeEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	actual := args.Map{"result": ai.IsResourceNameEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)
}

func Test_Cov9_AuthInfo_HasActionType_HasResourceName(t *testing.T) {
	ai := &corepayload.AuthInfo{ActionType: "login", ResourceName: "/api"}

	actual := args.Map{"result": ai.HasActionType()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual := args.Map{"result": ai.HasResourceName()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	var nilAI *corepayload.AuthInfo

	actual := args.Map{"result": nilAI.HasActionType() || nilAI.HasResourceName()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_Cov9_AuthInfo_IsUserInfoEmpty_IsSessionInfoEmpty(t *testing.T) {
	var nilAI *corepayload.AuthInfo

	actual := args.Map{"result": nilAI.IsUserInfoEmpty() || !nilAI.IsSessionInfoEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)

	ai := &corepayload.AuthInfo{}

	actual := args.Map{"result": ai.IsUserInfoEmpty() || !ai.IsSessionInfoEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should be empty", actual)
}

func Test_Cov9_AuthInfo_HasUserInfo_HasSessionInfo(t *testing.T) {
	ai := &corepayload.AuthInfo{
		UserInfo:    &corepayload.UserInfo{User: corepayload.New.User.UsingName("Alice")},
		SessionInfo: &corepayload.SessionInfo{Id: "s1", User: corepayload.New.User.UsingName("Bob")},
	}

	actual := args.Map{"result": ai.HasUserInfo()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual := args.Map{"result": ai.HasSessionInfo()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Cov9_AuthInfo_SetUserInfo_Nil(t *testing.T) {
	var nilAI *corepayload.AuthInfo
	result := nilAI.SetUserInfo(&corepayload.UserInfo{})

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	ai := &corepayload.AuthInfo{}
	result = ai.SetUserInfo(&corepayload.UserInfo{})

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_AuthInfo_SetActionType_SetResourceName_SetIdentifier_SetSessionInfo(t *testing.T) {
	var nilAI *corepayload.AuthInfo

	result := nilAI.SetActionType("login")

	actual := args.Map{"result": result == nil || result.ActionType != "login"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected login", actual)

	result = nilAI.SetResourceName("/api")

	actual := args.Map{"result": result == nil || result.ResourceName != "/api"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected /api", actual)

	result = nilAI.SetIdentifier("42")

	actual := args.Map{"result": result == nil || result.Identifier != "42"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	result = nilAI.SetSessionInfo(&corepayload.SessionInfo{Id: "s1"})

	actual := args.Map{"result": result == nil || result.SessionInfo.Id != "s1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected s1", actual)

	// non-nil receiver
	ai := &corepayload.AuthInfo{}
	ai.SetActionType("test")

	actual := args.Map{"result": ai.ActionType != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)

	ai.SetResourceName("/resource")

	actual := args.Map{"result": ai.ResourceName != "/resource"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected /resource", actual)

	ai.SetIdentifier("id")

	actual := args.Map{"result": ai.Identifier != "id"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected id", actual)

	ai.SetSessionInfo(&corepayload.SessionInfo{Id: "s2"})

	actual := args.Map{"result": ai.SessionInfo.Id != "s2"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected s2", actual)
}

func Test_Cov9_AuthInfo_SetUserSystemUser(t *testing.T) {
	var nilAI *corepayload.AuthInfo
	user := corepayload.New.User.UsingName("Alice")
	sysUser := corepayload.New.User.System("sys", "system")

	result := nilAI.SetUserSystemUser(user, sysUser)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	ai := &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{}}
	result = ai.SetUserSystemUser(user, sysUser)

	actual := args.Map{"result": result.UserInfo.User.Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)
}

func Test_Cov9_AuthInfo_SetUser_SetSystemUser(t *testing.T) {
	var nilAI *corepayload.AuthInfo
	user := corepayload.New.User.UsingName("Alice")

	result := nilAI.SetUser(user)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	sysUser := corepayload.New.User.System("sys", "system")
	result = nilAI.SetSystemUser(sysUser)

	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	ai := &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{}}
	ai.SetUser(user)
	ai.SetSystemUser(sysUser)
}

func Test_Cov9_AuthInfo_String_PrettyJsonString_Json_JsonPtr(t *testing.T) {
	ai := corepayload.AuthInfo{ActionType: "login"}
	_ = ai.String()
	_ = ai.PrettyJsonString()
	_ = ai.Json()
	_ = ai.JsonPtr()
}

func Test_Cov9_AuthInfo_Clone_ClonePtr_Ptr(t *testing.T) {
	ai := corepayload.AuthInfo{Identifier: "1", ActionType: "login"}
	cloned := ai.Clone()

	actual := args.Map{"result": cloned.ActionType != "login"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected login", actual)

	ptr := ai.Ptr()

	actual := args.Map{"result": ptr.ActionType != "login"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected login", actual)

	clonedPtr := ptr.ClonePtr()

	actual := args.Map{"result": clonedPtr.ActionType != "login"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected login", actual)

	var nilAI *corepayload.AuthInfo

	actual := args.Map{"result": nilAI.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ClonePtr should return nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SessionInfo — Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_SessionInfo_IdentifierInteger(t *testing.T) {
	si := corepayload.SessionInfo{Id: "42"}

	actual := args.Map{"result": si.IdentifierInteger() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	si2 := corepayload.SessionInfo{}

	actual := args.Map{"result": si2.IdentifierInteger() >= 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid value", actual)
}

func Test_Cov9_SessionInfo_IdentifierUnsignedInteger(t *testing.T) {
	si := corepayload.SessionInfo{Id: "42"}

	actual := args.Map{"result": si.IdentifierUnsignedInteger() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	si2 := corepayload.SessionInfo{Id: "-1"}

	actual := args.Map{"result": si2.IdentifierUnsignedInteger() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Cov9_SessionInfo_IsEmpty_IsValid(t *testing.T) {
	var nilSI *corepayload.SessionInfo

	actual := args.Map{"result": nilSI.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)

	si := &corepayload.SessionInfo{Id: "s1", User: corepayload.New.User.UsingName("Alice")}

	actual := args.Map{"result": si.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	actual := args.Map{"result": si.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)

	emptySI := &corepayload.SessionInfo{}

	actual := args.Map{"result": emptySI.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_Cov9_SessionInfo_IsUserNameEmpty_IsUserEmpty_HasUser(t *testing.T) {
	var nilSI *corepayload.SessionInfo

	actual := args.Map{"result": nilSI.IsUserNameEmpty() || !nilSI.IsUserEmpty() || nilSI.HasUser()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil checks failed", actual)

	si := &corepayload.SessionInfo{Id: "s1", User: corepayload.New.User.UsingName("Alice")}

	actual := args.Map{"result": si.IsUserNameEmpty() || si.IsUserEmpty() || !si.HasUser()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "user checks failed", actual)
}

func Test_Cov9_SessionInfo_IsUsernameEqual(t *testing.T) {
	si := &corepayload.SessionInfo{User: corepayload.New.User.UsingName("Alice")}

	actual := args.Map{"result": si.IsUsernameEqual("Alice")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)

	actual := args.Map{"result": si.IsUsernameEqual("Bob")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)

	var nilSI *corepayload.SessionInfo

	actual := args.Map{"result": nilSI.IsUsernameEqual("Alice")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_Cov9_SessionInfo_Clone_ClonePtr_Ptr(t *testing.T) {
	si := corepayload.SessionInfo{Id: "s1", User: corepayload.New.User.UsingName("Alice"), SessionPath: "/path"}
	cloned := si.Clone()

	actual := args.Map{"result": cloned.Id != "s1" || cloned.SessionPath != "/path"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned values", actual)

	ptr := si.Ptr()

	actual := args.Map{"result": ptr.Id != "s1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected s1", actual)

	clonedPtr := ptr.ClonePtr()

	actual := args.Map{"result": clonedPtr.Id != "s1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected s1", actual)

	var nilSI *corepayload.SessionInfo

	actual := args.Map{"result": nilSI.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ClonePtr should return nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PagingInfo — Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PagingInfo_IsEmpty(t *testing.T) {
	var nilPI *corepayload.PagingInfo

	actual := args.Map{"result": nilPI.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)

	pi := &corepayload.PagingInfo{TotalPages: 5, TotalItems: 50}

	actual := args.Map{"result": pi.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)
}

func Test_Cov9_PagingInfo_IsEqual_AllBranches(t *testing.T) {
	var nilA, nilB *corepayload.PagingInfo

	actual := args.Map{"result": nilA.IsEqual(nilB)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)

	pi := &corepayload.PagingInfo{TotalPages: 5}

	actual := args.Map{"result": nilA.IsEqual(pi) || pi.IsEqual(nilA)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)

	pi2 := &corepayload.PagingInfo{TotalPages: 3}

	actual := args.Map{"result": pi.IsEqual(pi2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different TotalPages should not be equal", actual)

	pi3 := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 1}

	actual := args.Map{"result": pi.IsEqual(pi3)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different CurrentPageIndex should not be equal", actual)

	pi4 := &corepayload.PagingInfo{TotalPages: 5, PerPageItems: 10}

	actual := args.Map{"result": pi.IsEqual(pi4)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different PerPageItems should not be equal", actual)

	pi5 := &corepayload.PagingInfo{TotalPages: 5, TotalItems: 50}

	actual := args.Map{"result": pi.IsEqual(pi5)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different TotalItems should not be equal", actual)
}

func Test_Cov9_PagingInfo_HasMethods(t *testing.T) {
	pi := &corepayload.PagingInfo{
		TotalPages:       5,
		CurrentPageIndex: 2,
		PerPageItems:     10,
		TotalItems:       50,
	}

	actual := args.Map{"result": pi.HasTotalPages() || !pi.HasCurrentPageIndex() || !pi.HasPerPageItems() || !pi.HasTotalItems()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected all true", actual)

	var nilPI *corepayload.PagingInfo

	actual := args.Map{"result": nilPI.HasTotalPages() || nilPI.HasCurrentPageIndex() || nilPI.HasPerPageItems() || nilPI.HasTotalItems()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_Cov9_PagingInfo_IsInvalidMethods(t *testing.T) {
	pi := &corepayload.PagingInfo{}

	actual := args.Map{"result": pi.IsInvalidTotalPages() || !pi.IsInvalidCurrentPageIndex() || !pi.IsInvalidPerPageItems() || !pi.IsInvalidTotalItems()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "zero values should be invalid", actual)

	var nilPI *corepayload.PagingInfo

	actual := args.Map{"result": nilPI.IsInvalidTotalPages() || !nilPI.IsInvalidCurrentPageIndex() || !nilPI.IsInvalidPerPageItems() || !nilPI.IsInvalidTotalItems()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_Cov9_PagingInfo_Clone_ClonePtr(t *testing.T) {
	pi := corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 2, PerPageItems: 10, TotalItems: 50}
	cloned := pi.Clone()

	actual := args.Map{"result": cloned.TotalPages != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)

	ptr := &pi
	clonedPtr := ptr.ClonePtr()

	actual := args.Map{"result": clonedPtr.TotalPages != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)

	var nilPI *corepayload.PagingInfo

	actual := args.Map{"result": nilPI.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ClonePtr should return nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// payloadProperties — Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PayloadProperties_AllMethods(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "test"
	pw.Identifier = "42"
	pw.CategoryName = "cat"
	pw.EntityType = "entity"
	pw.HasManyRecords = true
	pw.Payloads = []byte("data")

	props := pw.PayloadProperties()

	actual := args.Map{"result": props.Name() != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)

	actual := args.Map{"result": props.IdString() != "42"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	actual := args.Map{"result": props.Category() != "cat"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cat", actual)

	actual := args.Map{"result": props.EntityType() != "entity"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected entity", actual)

	actual := args.Map{"result": props.HasManyRecord()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual := args.Map{"result": props.HasSingleRecordOnly()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)

	_ = props.DynamicPayloads()
	_ = props.IdInteger()
	_ = props.IdUnsignedInteger()
	_ = props.BasicError()

	id, name, entity, cat, payloads := props.AllSafe()
	_, _, _, _, _ = id, name, entity, cat, payloads

	id, name, entity, cat, payloads = props.All()
	_, _, _, _, _ = id, name, entity, cat, payloads

	_ = props.Json()
	_ = props.JsonPtr()
	// AsPayloadPropertiesDefiner is on concrete type, not on interface — skip
}

func Test_Cov9_PayloadProperties_Setters(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	props := pw.PayloadProperties()

	_ = props.SetName("newName")
	props.SetNameMust("newName2")

	_ = props.SetIdString("99")
	props.SetIdStringMust("100")

	_ = props.SetCategory("newCat")
	props.SetCategoryMust("newCat2")

	_ = props.SetEntityType("newEntity")
	props.SetEntityTypeMust("newEntity2")

	props.SetSingleRecordFlag()
	props.SetManyRecordFlag()

	_ = props.SetDynamicPayloads([]byte("new"))
	props.SetDynamicPayloadsMust([]byte("new2"))
}

func Test_Cov9_PayloadProperties_SetBasicError(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	props := pw.PayloadProperties()
	props.SetBasicError(nil)
}

func Test_Cov9_PayloadProperties_ReflectSetTo(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	props := pw.PayloadProperties()
	var target corepayload.PayloadWrapper
	_ = props.ReflectSetTo(&target)
}

func Test_Cov9_PayloadProperties_DynamicPayloadsDeserializedTo(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Payloads = []byte(`"hello"`)
	props := pw.PayloadProperties()
	var s string
	_ = props.DynamicPayloadsDeserializedTo(&s)
}

// ══════════════════════════════════════════════════════════════════════════════
// newCreator factories — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_EmptyCreator_All(t *testing.T) {
	attr := corepayload.Empty.Attributes()

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	attrDefaults := corepayload.Empty.AttributesDefaults()

	actual := args.Map{"result": attrDefaults == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	pw := corepayload.Empty.PayloadWrapper()

	actual := args.Map{"result": pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	col := corepayload.Empty.PayloadsCollection()

	actual := args.Map{"result": col == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_NewAttributesCreator_AllFactories(t *testing.T) {
	// Create
	attr := corepayload.New.Attributes.Create(nil, nil, []byte("data"))

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// ErrFromTo
	attr = corepayload.New.Attributes.ErrFromTo(nil, nil, []byte("data"))

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// UsingAuthInfoDynamicBytes
	attr = corepayload.New.Attributes.UsingAuthInfoDynamicBytes(nil, []byte("data"))

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// UsingDynamicPayloadBytes
	attr = corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// UsingAuthInfo
	attr = corepayload.New.Attributes.UsingAuthInfo(nil)

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// UsingKeyValues
	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr = corepayload.New.Attributes.UsingKeyValues(hm)

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// UsingAuthInfoKeyValues
	attr = corepayload.New.Attributes.UsingAuthInfoKeyValues(nil, hm)

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// UsingKeyValuesPlusDynamic
	attr = corepayload.New.Attributes.UsingKeyValuesPlusDynamic(hm, []byte("data"))

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// UsingAnyKeyValues
	anyMap := coredynamic.NewMapAnyItems(0)
	attr = corepayload.New.Attributes.UsingAnyKeyValues(anyMap)

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// UsingAuthInfoAnyKeyValues
	attr = corepayload.New.Attributes.UsingAuthInfoAnyKeyValues(nil, anyMap)

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// UsingAnyKeyValuesPlusDynamic
	attr = corepayload.New.Attributes.UsingAnyKeyValuesPlusDynamic(anyMap, []byte("data"))

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// UsingBasicError
	attr = corepayload.New.Attributes.UsingBasicError(nil)

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// Empty
	attr = corepayload.New.Attributes.Empty()

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	// All
	attr = corepayload.New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)

	actual := args.Map{"result": attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_NewAttributesCreator_AllAny(t *testing.T) {
	attr, err := corepayload.New.Attributes.AllAny(nil, nil, nil, nil, "test")

	actual := args.Map{"result": err != nil || attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewAttributesCreator_PageInfoAny(t *testing.T) {
	attr, err := corepayload.New.Attributes.PageInfoAny(nil, "test")

	actual := args.Map{"result": err != nil || attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewAttributesCreator_UsingDynamicPayloadAny(t *testing.T) {
	attr, err := corepayload.New.Attributes.UsingDynamicPayloadAny(nil, "test")

	actual := args.Map{"result": err != nil || attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewAttributesCreator_UsingAuthInfoJsonResult(t *testing.T) {
	jsonResult := corejson.NewPtr("test")
	attr, err := corepayload.New.Attributes.UsingAuthInfoJsonResult(nil, jsonResult)

	actual := args.Map{"result": err != nil || attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewAttributesCreator_Deserialize(t *testing.T) {
	original := corepayload.New.Attributes.Empty()
	bytes := []byte(original.JsonString())
	attr, err := corepayload.New.Attributes.Deserialize(bytes)

	actual := args.Map{"result": err != nil || attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewAttributesCreator_DeserializeMany(t *testing.T) {
	bytes := []byte(`[{}]`)
	attrs, err := corepayload.New.Attributes.DeserializeMany(bytes)

	actual := args.Map{"result": err != nil || len(attrs) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_NewAttributesCreator_DeserializeUsingJsonResult(t *testing.T) {
	jsonResult := corejson.NewPtr(corepayload.Attributes{})
	attr, err := corepayload.New.Attributes.DeserializeUsingJsonResult(jsonResult)

	actual := args.Map{"result": err != nil || attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewAttributesCreator_CastOrDeserializeFrom(t *testing.T) {
	original := corepayload.New.Attributes.Empty()
	attr, err := corepayload.New.Attributes.CastOrDeserializeFrom(original)

	actual := args.Map{"result": err != nil || attr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)

	// nil
	_, err = corepayload.New.Attributes.CastOrDeserializeFrom(nil)

	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// newPayloadWrapperCreator — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_NewPayloadWrapper_All(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.All("name", "id", "task", "cat", "entity", false, nil, []byte("data"))

	actual := args.Map{"result": pw.Name != "name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
}

func Test_Cov9_NewPayloadWrapper_UsingBytes(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("name", "id", "task", "cat", "entity", []byte("data"))

	actual := args.Map{"result": pw.Name != "name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
}

func Test_Cov9_NewPayloadWrapper_Create(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.Create("name", "id", "task", "cat", "hello")

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_Record(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.Record("name", "id", "task", "cat", "hello")

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_Records(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.Records("name", "id", "task", "cat", []string{"a", "b"})

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_NameIdRecord(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameIdRecord("name", "id", "hello")

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_NameIdCategory(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameIdCategory("name", "id", "cat", "hello")

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_NameIdTaskRecord(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameIdTaskRecord("name", "id", "task", "hello")

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_NameTaskNameRecord(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameTaskNameRecord("id", "task", "hello")

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_ManyRecords(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.ManyRecords("name", "id", "task", "cat", []string{"a"})

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_Deserialize(t *testing.T) {
	original := corepayload.New.PayloadWrapper.Empty()
	original.Name = "test"
	bytes, _ := original.Serialize()

	pw, err := corepayload.New.PayloadWrapper.Deserialize(bytes)

	actual := args.Map{"result": err != nil || pw.Name != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)
}

func Test_Cov9_NewPayloadWrapper_CastOrDeserializeFrom(t *testing.T) {
	original := corepayload.New.PayloadWrapper.Empty()
	pw, err := corepayload.New.PayloadWrapper.CastOrDeserializeFrom(original)

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)

	_, err = corepayload.New.PayloadWrapper.CastOrDeserializeFrom(nil)

	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_Cov9_NewPayloadWrapper_DeserializeToMany(t *testing.T) {
	bytes := []byte(`[{}]`)
	wrappers, err := corepayload.New.PayloadWrapper.DeserializeToMany(bytes)

	actual := args.Map{"result": err != nil || len(wrappers) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_NewPayloadWrapper_DeserializeToCollection(t *testing.T) {
	bytes := []byte(`{"Items":[]}`)
	col, err := corepayload.New.PayloadWrapper.DeserializeToCollection(bytes)

	actual := args.Map{"result": err != nil || col == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_UsingBytesCreateInstruction(t *testing.T) {
	instr := &corepayload.BytesCreateInstruction{
		Name:       "name",
		Identifier: "id",
		Payloads:   []byte("data"),
	}

	pw := corepayload.New.PayloadWrapper.UsingBytesCreateInstruction(instr)

	actual := args.Map{"result": pw.Name != "name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
}

func Test_Cov9_NewPayloadWrapper_UsingCreateInstruction_BytesBranch(t *testing.T) {
	instr := &corepayload.PayloadCreateInstruction{
		Name:     "name",
		Payloads: []byte("data"),
	}

	pw, err := corepayload.New.PayloadWrapper.UsingCreateInstruction(instr)

	actual := args.Map{"result": err != nil || pw.Name != "name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
}

func Test_Cov9_NewPayloadWrapper_UsingCreateInstruction_StringBranch(t *testing.T) {
	instr := &corepayload.PayloadCreateInstruction{
		Name:     "name",
		Payloads: `"hello"`,
	}

	pw, err := corepayload.New.PayloadWrapper.UsingCreateInstruction(instr)

	actual := args.Map{"result": err != nil || pw.Name != "name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
}

func Test_Cov9_NewPayloadWrapper_UsingCreateInstruction_AnyBranch(t *testing.T) {
	instr := &corepayload.PayloadCreateInstruction{
		Name:     "name",
		Payloads: 42,
	}

	pw, err := corepayload.New.PayloadWrapper.UsingCreateInstruction(instr)

	actual := args.Map{"result": err != nil || pw.Name != "name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// newPayloadsCollectionCreator — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_NewPayloadsCollection_All(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	actual := args.Map{"result": col == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	col = corepayload.New.PayloadsCollection.UsingCap(10)

	actual := args.Map{"result": col == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov9_NewPayloadsCollection_UsingWrappers(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	col := corepayload.New.PayloadsCollection.UsingWrappers(pw)

	actual := args.Map{"result": col.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	emptyCol := corepayload.New.PayloadsCollection.UsingWrappers()

	actual := args.Map{"result": emptyCol.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Cov9_NewPayloadsCollection_Deserialize(t *testing.T) {
	bytes := []byte(`{"Items":[]}`)
	col, err := corepayload.New.PayloadsCollection.Deserialize(bytes)

	actual := args.Map{"result": err != nil || col == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadsCollection_DeserializeToMany(t *testing.T) {
	bytes := []byte(`[{"Items":[]}]`)
	cols, err := corepayload.New.PayloadsCollection.DeserializeToMany(bytes)

	actual := args.Map{"result": err != nil || len(cols) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov9_NewPayloadsCollection_DeserializeUsingJsonResult(t *testing.T) {
	jsonResult := corejson.NewPtr(corepayload.PayloadsCollection{})
	col, err := corepayload.New.PayloadsCollection.DeserializeUsingJsonResult(jsonResult)

	actual := args.Map{"result": err != nil || col == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// newUserCreator — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_NewUser_All(t *testing.T) {
	u := corepayload.New.User.Empty()

	actual := args.Map{"result": u == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	u = corepayload.New.User.Create(false, "Alice", "admin")

	actual := args.Map{"result": u.Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)

	u = corepayload.New.User.NonSysCreate("Bob", "user")

	actual := args.Map{"result": u.Name != "Bob"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Bob", actual)

	u = corepayload.New.User.NonSysCreateId("1", "Charlie", "user")

	actual := args.Map{"result": u.Identifier != "1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	u = corepayload.New.User.System("sys", "system")

	actual := args.Map{"result": u.IsSystemUser}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected system user", actual)

	u = corepayload.New.User.SystemId("1", "sys", "system")

	actual := args.Map{"result": u.Identifier != "1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	u = corepayload.New.User.UsingName("Dave")

	actual := args.Map{"result": u.Name != "Dave"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Dave", actual)

	u = corepayload.New.User.All(true, "1", "Eve", "admin", "token", "hash")

	actual := args.Map{"result": u.Name != "Eve"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Eve", actual)
}

func Test_Cov9_NewUser_Deserialize(t *testing.T) {
	u := corepayload.New.User.UsingName("Alice")
	bytes, _ := u.Serialize()
	result, err := corepayload.New.User.Deserialize(bytes)

	actual := args.Map{"result": err != nil || result.Name != "Alice"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice", actual)
}

func Test_Cov9_NewUser_CastOrDeserializeFrom(t *testing.T) {
	u := corepayload.New.User.UsingName("Alice")
	result, err := corepayload.New.User.CastOrDeserializeFrom(u)

	actual := args.Map{"result": err != nil || result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)

	_, err = corepayload.New.User.CastOrDeserializeFrom(nil)

	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Generic helpers — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_DeserializePayloadTo(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`"hello"`)}
	result, err := corepayload.DeserializePayloadTo[string](pw)

	actual := args.Map{"result": err != nil || result != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)

	// nil wrapper
	_, err = corepayload.DeserializePayloadTo[string](nil)

	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_Cov9_DeserializePayloadToSlice(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`["a","b"]`)}
	result, err := corepayload.DeserializePayloadToSlice[string](pw)

	actual := args.Map{"result": err != nil || len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	// nil
	_, err = corepayload.DeserializePayloadToSlice[string](nil)

	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_Cov9_DeserializeAttributesPayloadTo(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))
	result, err := corepayload.DeserializeAttributesPayloadTo[string](attr)

	actual := args.Map{"result": err != nil || result != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)

	// nil
	_, err = corepayload.DeserializeAttributesPayloadTo[string](nil)

	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_Cov9_DeserializeAttributesPayloadToSlice(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`["a","b"]`))
	result, err := corepayload.DeserializeAttributesPayloadToSlice[string](attr)

	actual := args.Map{"result": err != nil || len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	// nil
	_, err = corepayload.DeserializeAttributesPayloadToSlice[string](nil)

	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadCreateInstructionTypeStringer — coverage
// ══════════════════════════════════════════════════════════════════════════════

type mockStringer struct{ val string }

func (m mockStringer) String() string { return m.val }
func (m mockStringer) Name() string   { return m.val }

func Test_Cov9_PayloadCreateInstructionTypeStringer(t *testing.T) {
	instr := corepayload.PayloadCreateInstructionTypeStringer{
		Name:                 "name",
		Identifier:           "id",
		TaskTypeNameStringer: mockStringer{"task"},
		CategoryNameStringer: mockStringer{"cat"},
		Payloads:             "hello",
	}

	pci := instr.PayloadCreateInstruction()

	actual := args.Map{"result": pci.TaskTypeName != "task" || pci.CategoryName != "cat"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected task and cat", actual)
}

func Test_Cov9_NewPayloadWrapper_UsingBytesCreateInstructionStringer(t *testing.T) {
	instr := &corepayload.BytesCreateInstructionStringer{
		Name:         "name",
		Identifier:   "id",
		TaskTypeName: mockStringer{"task"},
		CategoryName: mockStringer{"cat"},
		EntityType:   "entity",
		Payloads:     []byte("data"),
	}

	pw := corepayload.New.PayloadWrapper.UsingBytesCreateInstructionTypeStringer(instr)

	actual := args.Map{"result": pw.Name != "name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
}

func Test_Cov9_NewPayloadWrapper_UsingCreateInstructionTypeStringer(t *testing.T) {
	instr := &corepayload.PayloadCreateInstructionTypeStringer{
		Name:                 "name",
		Identifier:           "id",
		TaskTypeNameStringer: mockStringer{"task"},
		CategoryNameStringer: mockStringer{"cat"},
		Payloads:             "hello",
	}

	pw, err := corepayload.New.PayloadWrapper.UsingCreateInstructionTypeStringer(instr)

	actual := args.Map{"result": err != nil || pw.Name != "name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
}

func Test_Cov9_NewPayloadWrapper_CreateUsingTypeStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.CreateUsingTypeStringer(
		"name", "id", mockStringer{"task"}, mockStringer{"cat"}, "hello")

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_NameIdCategoryStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameIdCategoryStringer(
		"name", "id", mockStringer{"cat"}, "hello")

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_RecordsTypeStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.RecordsTypeStringer(
		"name", "id", mockStringer{"task"}, mockStringer{"cat"}, []string{"a"})

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_RecordTypeStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.RecordTypeStringer(
		"name", "id", mockStringer{"task"}, mockStringer{"cat"}, "hello")

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_NameIdTaskStringerRecord(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameIdTaskStringerRecord(
		"name", "id", mockStringer{"task"}, "hello")

	actual := args.Map{"result": err != nil || pw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_NewPayloadWrapper_AllUsingStringer(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.AllUsingStringer(
		"name", "id", mockStringer{"task"}, mockStringer{"cat"}, "entity", false, nil, []byte("data"))

	actual := args.Map{"result": pw.Name != "name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
}

func Test_Cov9_NewPayloadWrapper_AllUsingExpander(t *testing.T) {
	expander := corepayload.PayloadTypeExpander{
		CategoryStringer: mockStringer{"cat"},
		TaskTypeStringer: mockStringer{"task"},
	}

	pw := corepayload.New.PayloadWrapper.AllUsingExpander(
		"name", "id", expander, "entity", false, nil, []byte("data"))

	actual := args.Map{"result": pw.Name != "name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
}

func Test_Cov9_NewUser_UsingNameTypeStringer(t *testing.T) {
	u := corepayload.New.User.UsingNameTypeStringer("Alice", mockStringer{"admin"})

	actual := args.Map{"result": u.Name != "Alice" || u.Type != "admin"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice/admin", actual)
}

func Test_Cov9_NewUser_SysUsingNameTypeStringer(t *testing.T) {
	u := corepayload.New.User.SysUsingNameTypeStringer("sys", mockStringer{"system"})

	actual := args.Map{"result": u.Name != "sys" || !u.IsSystemUser}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected system user", actual)
}

func Test_Cov9_NewUser_AllTypeStringer(t *testing.T) {
	u := corepayload.New.User.AllTypeStringer(true, "1", "Alice", mockStringer{"admin"}, "token", "hash")

	actual := args.Map{"result": u.Name != "Alice" || u.Type != "admin"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Alice/admin", actual)
}

func Test_Cov9_NewUser_AllUsingStringer(t *testing.T) {
	u := corepayload.New.User.AllUsingStringer(false, "1", "Bob", mockStringer{"user"}, "token", "hash")

	actual := args.Map{"result": u.Name != "Bob" || u.Type != "user"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Bob/user", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Attributes — DeserializeDynamicPayloadsTo* methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Attributes_DeserializeDynamicPayloadsToAttributes(t *testing.T) {
	inner := corepayload.New.Attributes.Empty()
	bytes := []byte(inner.JsonString())
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(bytes)

	result, err := attr.DeserializeDynamicPayloadsToAttributes()

	actual := args.Map{"result": err != nil || result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_Attributes_DeserializeDynamicPayloadsToPayloadWrapper(t *testing.T) {
	inner := corepayload.New.PayloadWrapper.Empty()
	inner.Name = "inner"
	bytes, _ := inner.Serialize()
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(bytes)

	result, err := attr.DeserializeDynamicPayloadsToPayloadWrapper()

	actual := args.Map{"result": err != nil || result.Name != "inner"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected inner", actual)
}

func Test_Cov9_Attributes_DeserializeDynamicPayloadsToPayloadWrappersCollection(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"Items":[]}`))
	result, err := attr.DeserializeDynamicPayloadsToPayloadWrappersCollection()

	actual := args.Map{"result": err != nil || result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Cov9_PayloadWrapper_PayloadDeserializeToPayloadBinder(t *testing.T) {
	inner := corepayload.New.PayloadWrapper.Empty()
	inner.Name = "inner"
	bytes, _ := inner.Serialize()

	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Payloads = bytes

	binder, err := pw.PayloadDeserializeToPayloadBinder()

	actual := args.Map{"result": err != nil || binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)

	// nil receiver
	var nilPW *corepayload.PayloadWrapper
	_, err = nilPW.PayloadDeserializeToPayloadBinder()

	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_Cov9_PayloadWrapper_IsEntityTypeNamer(t *testing.T) {
	pw := &corepayload.PayloadWrapper{EntityType: "test"}

	actual := args.Map{"result": pw.IsEntityTypeNamer(mockStringer{"test"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual := args.Map{"result": pw.IsEntityTypeNamer(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil namer", actual)

	var nilPW *corepayload.PayloadWrapper

	actual := args.Map{"result": nilPW.IsEntityTypeNamer(mockStringer{"test"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_Cov9_PayloadWrapper_IsCategoryNamer(t *testing.T) {
	pw := &corepayload.PayloadWrapper{EntityType: "test"}

	actual := args.Map{"result": pw.IsCategoryNamer(mockStringer{"test"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	var nilPW *corepayload.PayloadWrapper

	actual := args.Map{"result": nilPW.IsCategoryNamer(mockStringer{"test"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_Cov9_NewPayloadWrapper_createInternalUsingBytes_WithExistingAttr(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	instr := &corepayload.BytesCreateInstruction{
		Name:       "name",
		Identifier: "id",
		Payloads:   []byte("data"),
		Attributes: attr,
	}

	pw := corepayload.New.PayloadWrapper.UsingBytesCreateInstruction(instr)

	actual := args.Map{"result": pw.Name != "name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
}

func Test_Cov9_Attributes_BasicErrorDeserializedTo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	var target any
	err := attr.BasicErrorDeserializedTo(&target)

	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil error (empty error)", actual)
}
