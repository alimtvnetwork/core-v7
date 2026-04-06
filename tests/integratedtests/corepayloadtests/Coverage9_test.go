package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coreinstruction"
)

// ══════════════════════════════════════════════════════════════════════════════
// Attributes — Getters (AttributesGetters.go)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Attributes_IsNull(t *testing.T) {
	var nilAttr *corepayload.Attributes
	attr := corepayload.New.Attributes.Empty()

	if !nilAttr.IsNull() {
		t.Fatal("nil Attributes should be null")
	}

	if attr.IsNull() {
		t.Fatal("non-nil Attributes should not be null")
	}
}

func Test_Cov9_Attributes_HasSafeItems(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))

	if !attr.HasSafeItems() {
		t.Fatal("expected HasSafeItems to be true")
	}
}

func Test_Cov9_Attributes_HasStringKey(t *testing.T) {
	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr := corepayload.New.Attributes.UsingKeyValues(hm)

	if !attr.HasStringKey("k") {
		t.Fatal("expected HasStringKey to be true")
	}

	if attr.HasStringKey("missing") {
		t.Fatal("expected HasStringKey to be false for missing")
	}
}

func Test_Cov9_Attributes_HasAnyKey(t *testing.T) {
	anyMap := coredynamic.NewMapAnyItems(0)
	anyMap.Add("k", 42)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)

	if !attr.HasAnyKey("k") {
		t.Fatal("expected HasAnyKey to be true")
	}

	if attr.HasAnyKey("missing") {
		t.Fatal("expected HasAnyKey to be false for missing")
	}
}

func Test_Cov9_Attributes_Payloads(t *testing.T) {
	var nilAttr *corepayload.Attributes
	p := nilAttr.Payloads()

	if len(p) != 0 {
		t.Fatal("nil Payloads should return empty")
	}

	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	if string(attr.Payloads()) != "data" {
		t.Fatal("expected Payloads to return data")
	}
}

func Test_Cov9_Attributes_PayloadsString(t *testing.T) {
	var nilAttr *corepayload.Attributes

	if nilAttr.PayloadsString() != "" {
		t.Fatal("nil PayloadsString should return empty")
	}

	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	if attr.PayloadsString() != "data" {
		t.Fatal("expected PayloadsString to return data")
	}

	emptyAttr := corepayload.New.Attributes.Empty()

	if emptyAttr.PayloadsString() != "" {
		t.Fatal("empty PayloadsString should return empty")
	}
}

func Test_Cov9_Attributes_AnyKeyValMap(t *testing.T) {
	var nilAttr *corepayload.Attributes
	m := nilAttr.AnyKeyValMap()

	if len(m) != 0 {
		t.Fatal("nil AnyKeyValMap should return empty map")
	}

	anyMap := coredynamic.NewMapAnyItems(0)
	anyMap.Add("k", 42)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	m = attr.AnyKeyValMap()

	if m["k"] != 42 {
		t.Fatal("expected 42 for key k")
	}
}

func Test_Cov9_Attributes_Hashmap(t *testing.T) {
	var nilAttr *corepayload.Attributes
	m := nilAttr.Hashmap()

	if len(m) != 0 {
		t.Fatal("nil Hashmap should return empty map")
	}

	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr := corepayload.New.Attributes.UsingKeyValues(hm)
	m = attr.Hashmap()

	if m["k"] != "v" {
		t.Fatal("expected v for key k")
	}
}

func Test_Cov9_Attributes_HasIssuesOrEmpty(t *testing.T) {
	var nilAttr *corepayload.Attributes

	if !nilAttr.HasIssuesOrEmpty() {
		t.Fatal("nil should have issues or be empty")
	}

	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"data"`))

	if attr.HasIssuesOrEmpty() {
		t.Fatal("non-empty valid attr should not have issues")
	}
}

func Test_Cov9_Attributes_IsSafeValid(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"data"`))

	if !attr.IsSafeValid() {
		t.Fatal("expected IsSafeValid true")
	}
}

func Test_Cov9_Attributes_HasAnyItem(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"data"`))

	if !attr.HasAnyItem() {
		t.Fatal("expected HasAnyItem true")
	}
}

func Test_Cov9_Attributes_Count_Capacity(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"ab"`))

	if attr.Count() != attr.Length() {
		t.Fatal("Count should equal Length")
	}

	if attr.Capacity() != attr.Length() {
		t.Fatal("Capacity should equal Length")
	}
}

func Test_Cov9_Attributes_Length_Nil(t *testing.T) {
	var nilAttr *corepayload.Attributes

	if nilAttr.Length() != 0 {
		t.Fatal("nil Length should be 0")
	}
}

func Test_Cov9_Attributes_HasPagingInfo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.HasPagingInfo() {
		t.Fatal("expected no paging info")
	}

	attr.PagingInfo = &corepayload.PagingInfo{TotalPages: 5}

	if !attr.HasPagingInfo() {
		t.Fatal("expected paging info")
	}
}

func Test_Cov9_Attributes_HasKeyValuePairs(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.HasKeyValuePairs() {
		t.Fatal("empty should not have key value pairs")
	}

	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr2 := corepayload.New.Attributes.UsingKeyValues(hm)

	if !attr2.HasKeyValuePairs() {
		t.Fatal("expected key value pairs")
	}
}

func Test_Cov9_Attributes_HasFromTo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.HasFromTo() {
		t.Fatal("expected no FromTo")
	}

	attr.FromTo = &coreinstruction.FromTo{}

	if !attr.HasFromTo() {
		t.Fatal("expected FromTo present")
	}
}

func Test_Cov9_Attributes_IsValid_IsInvalid(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if !attr.IsValid() {
		t.Fatal("expected valid")
	}

	if !attr.IsInvalid() {
		t.Fatal("expected invalid")
	}

	var nilAttr *corepayload.Attributes

	if !nilAttr.IsInvalid() {
		t.Fatal("nil should be invalid")
	}
}

func Test_Cov9_Attributes_HasError_Error(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.HasError() {
		t.Fatal("expected no error")
	}

	if attr.Error() != nil {
		t.Fatal("expected nil error")
	}

	if attr.CompiledError() != nil {
		t.Fatal("expected nil CompiledError")
	}
}

func Test_Cov9_Attributes_IsEmptyError(t *testing.T) {
	var nilAttr *corepayload.Attributes

	if !nilAttr.IsEmptyError() {
		t.Fatal("nil should be empty error")
	}

	attr := corepayload.New.Attributes.Empty()

	if !attr.IsEmptyError() {
		t.Fatal("expected empty error")
	}
}

func Test_Cov9_Attributes_DynamicBytesLength(t *testing.T) {
	var nilAttr *corepayload.Attributes

	if nilAttr.DynamicBytesLength() != 0 {
		t.Fatal("nil DynamicBytesLength should be 0")
	}

	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("abc"))

	if attr.DynamicBytesLength() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_Cov9_Attributes_StringKeyValuePairsLength(t *testing.T) {
	var nilAttr *corepayload.Attributes

	if nilAttr.StringKeyValuePairsLength() != 0 {
		t.Fatal("nil should return 0")
	}

	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr := corepayload.New.Attributes.UsingKeyValues(hm)

	if attr.StringKeyValuePairsLength() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_Attributes_AnyKeyValuePairsLength(t *testing.T) {
	var nilAttr *corepayload.Attributes

	if nilAttr.AnyKeyValuePairsLength() != 0 {
		t.Fatal("nil should return 0")
	}

	anyMap := coredynamic.NewMapAnyItems(0)
	anyMap.Add("k", 42)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)

	if attr.AnyKeyValuePairsLength() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_Attributes_IsEmpty_HasItems(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if !attr.IsEmpty() {
		t.Fatal("empty attr should be empty")
	}

	if attr.HasItems() {
		t.Fatal("empty attr should not have items")
	}

	attr2 := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	if attr2.IsEmpty() {
		t.Fatal("non-empty attr should not be empty")
	}

	if !attr2.HasItems() {
		t.Fatal("non-empty attr should have items")
	}
}

func Test_Cov9_Attributes_IsPagingInfoEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if !attr.IsPagingInfoEmpty() {
		t.Fatal("expected paging info empty")
	}
}

func Test_Cov9_Attributes_IsKeyValuePairsEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if !attr.IsKeyValuePairsEmpty() {
		t.Fatal("expected key value pairs empty")
	}
}

func Test_Cov9_Attributes_IsAnyKeyValuePairsEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if !attr.IsAnyKeyValuePairsEmpty() {
		t.Fatal("expected any key value pairs empty")
	}
}

func Test_Cov9_Attributes_IsUserInfoEmpty_VirtualUser_SystemUser(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if !attr.IsUserInfoEmpty() {
		t.Fatal("expected user info empty")
	}

	if attr.VirtualUser() != nil {
		t.Fatal("expected nil VirtualUser")
	}

	if attr.SystemUser() != nil {
		t.Fatal("expected nil SystemUser")
	}

	// With user info
	user := corepayload.New.User.UsingName("Alice")
	sysUser := corepayload.New.User.System("sys", "system")
	userInfo := &corepayload.UserInfo{User: user, SystemUser: sysUser}
	attr2 := corepayload.New.Attributes.UsingAuthInfo(&corepayload.AuthInfo{UserInfo: userInfo})

	if attr2.IsUserInfoEmpty() {
		t.Fatal("expected user info not empty")
	}

	if attr2.VirtualUser().Name != "Alice" {
		t.Fatal("expected Alice")
	}

	if attr2.SystemUser().Name != "sys" {
		t.Fatal("expected sys")
	}
}

func Test_Cov9_Attributes_SessionUser(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.SessionUser() != nil {
		t.Fatal("expected nil SessionUser")
	}

	user := corepayload.New.User.UsingName("SessionUser")
	si := &corepayload.SessionInfo{Id: "s1", User: user}
	attr2 := corepayload.New.Attributes.UsingAuthInfo(&corepayload.AuthInfo{SessionInfo: si})

	if attr2.SessionUser().Name != "SessionUser" {
		t.Fatal("expected SessionUser")
	}
}

func Test_Cov9_Attributes_IsAuthInfoEmpty_IsSessionInfoEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if !attr.IsAuthInfoEmpty() {
		t.Fatal("expected auth info empty")
	}

	if !attr.IsSessionInfoEmpty() {
		t.Fatal("expected session info empty")
	}
}

func Test_Cov9_Attributes_HasUserInfo_HasAuthInfo_HasSessionInfo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.HasUserInfo() {
		t.Fatal("expected no user info")
	}

	if attr.HasAuthInfo() {
		t.Fatal("expected no auth info")
	}

	if attr.HasSessionInfo() {
		t.Fatal("expected no session info")
	}
}

func Test_Cov9_Attributes_SessionInfo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.SessionInfo() != nil {
		t.Fatal("expected nil SessionInfo")
	}

	si := &corepayload.SessionInfo{Id: "s1"}
	attr2 := corepayload.New.Attributes.UsingAuthInfo(&corepayload.AuthInfo{SessionInfo: si})

	if attr2.SessionInfo().Id != "s1" {
		t.Fatal("expected s1")
	}
}

func Test_Cov9_Attributes_AuthType(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.AuthType() != "" {
		t.Fatal("expected empty auth type")
	}

	attr2 := corepayload.New.Attributes.UsingAuthInfo(&corepayload.AuthInfo{ActionType: "login"})

	if attr2.AuthType() != "login" {
		t.Fatal("expected login")
	}
}

func Test_Cov9_Attributes_ResourceName(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.ResourceName() != "" {
		t.Fatal("expected empty resource name")
	}

	attr2 := corepayload.New.Attributes.UsingAuthInfo(&corepayload.AuthInfo{ResourceName: "/api/test"})

	if attr2.ResourceName() != "/api/test" {
		t.Fatal("expected /api/test")
	}
}

func Test_Cov9_Attributes_HasStringKeyValuePairs(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.HasStringKeyValuePairs() {
		t.Fatal("expected false")
	}
}

func Test_Cov9_Attributes_HasAnyKeyValuePairs(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.HasAnyKeyValuePairs() {
		t.Fatal("expected false")
	}
}

func Test_Cov9_Attributes_HasDynamicPayloads(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.HasDynamicPayloads() {
		t.Fatal("expected false")
	}

	attr2 := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	if !attr2.HasDynamicPayloads() {
		t.Fatal("expected true")
	}
}

func Test_Cov9_Attributes_GetStringKeyValue(t *testing.T) {
	var nilAttr *corepayload.Attributes
	_, found := nilAttr.GetStringKeyValue("k")

	if found {
		t.Fatal("nil should not find")
	}

	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr := corepayload.New.Attributes.UsingKeyValues(hm)
	val, found := attr.GetStringKeyValue("k")

	if !found || val != "v" {
		t.Fatal("expected v for key k")
	}
}

func Test_Cov9_Attributes_GetAnyKeyValue(t *testing.T) {
	var nilAttr *corepayload.Attributes
	_, found := nilAttr.GetAnyKeyValue("k")

	if found {
		t.Fatal("nil should not find")
	}
}

func Test_Cov9_Attributes_IsErrorDifferent_IsErrorEqual(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()

	if attr.IsErrorDifferent(nil) {
		t.Fatal("expected not different when both empty")
	}

	if !attr.IsErrorEqual(nil) {
		t.Fatal("expected equal when both empty")
	}
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

	if result.AuthInfo.ActionType != "test" {
		t.Fatal("expected test")
	}

	// nil receiver
	var nilAttr *corepayload.Attributes
	result = nilAttr.SetAuthInfo(&corepayload.AuthInfo{ActionType: "new"})

	if result == nil {
		t.Fatal("expected non-nil result from nil receiver")
	}
}

func Test_Cov9_Attributes_SetUserInfo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.AuthInfo = &corepayload.AuthInfo{}
	user := corepayload.New.User.UsingName("Alice")
	userInfo := &corepayload.UserInfo{User: user}
	result := attr.SetUserInfo(userInfo)

	if result.AuthInfo.UserInfo.User.Name != "Alice" {
		t.Fatal("expected Alice")
	}

	// nil receiver
	var nilAttr *corepayload.Attributes
	result = nilAttr.SetUserInfo(userInfo)

	if result == nil {
		t.Fatal("expected non-nil result from nil receiver")
	}
}

func Test_Cov9_Attributes_AddNewStringKeyValueOnly(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	added := attr.AddNewStringKeyValueOnly("k", "v")

	if !added {
		t.Fatal("expected added")
	}

	var nilAttr *corepayload.Attributes

	if nilAttr.AddNewStringKeyValueOnly("k", "v") {
		t.Fatal("nil should not add")
	}
}

func Test_Cov9_Attributes_AddNewAnyKeyValueOnly(t *testing.T) {
	anyMap := coredynamic.NewMapAnyItems(0)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	added := attr.AddNewAnyKeyValueOnly("k", 42)

	if !added {
		t.Fatal("expected added")
	}

	var nilAttr *corepayload.Attributes

	if nilAttr.AddNewAnyKeyValueOnly("k", 42) {
		t.Fatal("nil should not add")
	}
}

func Test_Cov9_Attributes_AddOrUpdateString(t *testing.T) {
	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr := corepayload.New.Attributes.UsingKeyValues(hm)
	isNew := attr.AddOrUpdateString("k2", "v2")

	if !isNew {
		t.Fatal("expected new key")
	}

	var nilAttr *corepayload.Attributes

	if nilAttr.AddOrUpdateString("k", "v") {
		t.Fatal("nil should return false")
	}
}

func Test_Cov9_Attributes_AddOrUpdateAnyItem(t *testing.T) {
	anyMap := coredynamic.NewMapAnyItems(0)
	attr := corepayload.New.Attributes.UsingAnyKeyValues(anyMap)
	isNew := attr.AddOrUpdateAnyItem("k", 42)

	if !isNew {
		t.Fatal("expected new key")
	}

	var nilAttr *corepayload.Attributes

	if nilAttr.AddOrUpdateAnyItem("k", 42) {
		t.Fatal("nil should return false")
	}
}

func Test_Cov9_Attributes_SetBasicErr(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	result := attr.SetBasicErr(nil)

	if result == nil {
		t.Fatal("expected non-nil")
	}

	// nil receiver
	var nilAttr *corepayload.Attributes
	result = nilAttr.SetBasicErr(nil)

	if result == nil {
		t.Fatal("expected non-nil from nil receiver")
	}
}

func Test_Cov9_Attributes_Clear_Dispose(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.AddNewStringKeyValueOnly("k", "v")
	attr.Clear()

	if attr.HasStringKeyValuePairs() {
		t.Fatal("expected empty after clear")
	}

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

	if result == "" {
		t.Fatal("expected pretty string")
	}

	var nilAttr *corepayload.Attributes

	if nilAttr.PayloadsPrettyString() != "" {
		t.Fatal("nil should return empty")
	}
}

func Test_Cov9_Attributes_PayloadsJsonResult(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"key":"value"}`))
	result := attr.PayloadsJsonResult()

	if result == nil {
		t.Fatal("expected non-nil result")
	}

	emptyAttr := corepayload.New.Attributes.Empty()
	result = emptyAttr.PayloadsJsonResult()

	if result == nil {
		t.Fatal("expected non-nil empty result")
	}
}

func Test_Cov9_Attributes_JsonString_JsonStringMust(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	s := attr.JsonString()

	if s == "" {
		t.Fatal("expected non-empty json string")
	}

	s = attr.JsonStringMust()

	if s == "" {
		t.Fatal("expected non-empty json string must")
	}
}

func Test_Cov9_Attributes_String(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	s := attr.String()

	if s == "" {
		t.Fatal("expected non-empty string")
	}
}

func Test_Cov9_Attributes_PrettyJsonString(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	s := attr.PrettyJsonString()

	if s == "" {
		t.Fatal("expected non-empty pretty json")
	}
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

	if binder == nil {
		t.Fatal("expected non-nil binder")
	}
}

func Test_Cov9_Attributes_AsAttributesBinder(t *testing.T) {
	attr := corepayload.Attributes{}
	binder := attr.AsAttributesBinder()

	if binder == nil {
		t.Fatal("expected non-nil binder")
	}
}

func Test_Cov9_Attributes_ParseInjectUsingJson(t *testing.T) {
	attr := &corepayload.Attributes{}
	jsonResult := corejson.NewPtr(corepayload.Attributes{})
	result, err := attr.ParseInjectUsingJson(jsonResult)

	if err != nil {
		t.Fatal("expected no error")
	}

	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func Test_Cov9_Attributes_ParseInjectUsingJsonMust(t *testing.T) {
	attr := &corepayload.Attributes{}
	jsonResult := corejson.NewPtr(corepayload.Attributes{})
	result := attr.ParseInjectUsingJsonMust(jsonResult)

	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func Test_Cov9_Attributes_JsonParseSelfInject(t *testing.T) {
	attr := &corepayload.Attributes{}
	jsonResult := corejson.NewPtr(corepayload.Attributes{})
	err := attr.JsonParseSelfInject(jsonResult)

	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_Cov9_Attributes_DeserializeDynamicPayloads(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))
	var s string
	err := attr.DeserializeDynamicPayloads(&s)

	if err != nil || s != "hello" {
		t.Fatal("expected hello")
	}
}

func Test_Cov9_Attributes_DeserializeDynamicPayloadsMust(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))
	var s string
	attr.DeserializeDynamicPayloadsMust(&s)

	if s != "hello" {
		t.Fatal("expected hello")
	}
}

func Test_Cov9_Attributes_DynamicPayloadsDeserialize(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))
	var s string
	err := attr.DynamicPayloadsDeserialize(&s)

	if err != nil || s != "hello" {
		t.Fatal("expected hello")
	}

	// nil receiver
	var nilAttr *corepayload.Attributes
	err = nilAttr.DynamicPayloadsDeserialize(&s)

	if err == nil {
		t.Fatal("nil should return error")
	}
}

func Test_Cov9_Attributes_DynamicPayloadsDeserializeMust(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))
	var s string
	attr.DynamicPayloadsDeserializeMust(&s)

	if s != "hello" {
		t.Fatal("expected hello")
	}
}

func Test_Cov9_Attributes_AnyKeyReflectSetTo(t *testing.T) {
	var nilAttr *corepayload.Attributes
	err := nilAttr.AnyKeyReflectSetTo("k", nil)

	if err == nil {
		t.Fatal("nil should return error")
	}
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

	if !nilA.IsEqual(nilB) {
		t.Fatal("both nil should be equal")
	}

	attr := corepayload.New.Attributes.Empty()

	if nilA.IsEqual(attr) {
		t.Fatal("nil vs non-nil should not be equal")
	}

	if attr.IsEqual(nilA) {
		t.Fatal("non-nil vs nil should not be equal")
	}

	// Same pointer
	if !attr.IsEqual(attr) {
		t.Fatal("same pointer should be equal")
	}

	// Different paging
	a1 := corepayload.New.Attributes.Empty()
	a2 := corepayload.New.Attributes.Empty()
	a1.PagingInfo = &corepayload.PagingInfo{TotalPages: 1}

	if a1.IsEqual(a2) {
		t.Fatal("different paging should not be equal")
	}
}

func Test_Cov9_Attributes_Clone_Shallow(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"data"`))
	cloned, err := attr.Clone(false)

	if err != nil {
		t.Fatal("expected no error")
	}

	if string(cloned.DynamicPayloads) != `"data"` {
		t.Fatal("expected cloned data")
	}
}

func Test_Cov9_Attributes_Clone_Deep(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"data"`))
	cloned, err := attr.Clone(true)

	if err != nil {
		t.Fatal("expected no error")
	}

	if string(cloned.DynamicPayloads) != `"data"` {
		t.Fatal("expected cloned data")
	}
}

func Test_Cov9_Attributes_ClonePtr_Nil(t *testing.T) {
	var nilAttr *corepayload.Attributes
	cloned, err := nilAttr.ClonePtr(true)

	if err != nil || cloned != nil {
		t.Fatal("nil ClonePtr should return nil, nil")
	}
}

func Test_Cov9_Attributes_Clone_NilReturnsEmpty(t *testing.T) {
	var nilAttr *corepayload.Attributes
	cloned, err := nilAttr.Clone(true)

	if err != nil {
		t.Fatal("expected no error")
	}

	_ = cloned
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadWrapper — Core methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PayloadWrapper_HasSafeItems(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	if pw.HasSafeItems() {
		t.Fatal("empty should not have safe items")
	}
}

func Test_Cov9_PayloadWrapper_DynamicPayloads(t *testing.T) {
	var nilPW *corepayload.PayloadWrapper
	p := nilPW.DynamicPayloads()

	if len(p) != 0 {
		t.Fatal("nil should return empty")
	}

	pw := corepayload.New.PayloadWrapper.Empty()
	_ = pw.DynamicPayloads()
}

func Test_Cov9_PayloadWrapper_SetDynamicPayloads(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	err := pw.SetDynamicPayloads([]byte("data"))

	if err != nil {
		t.Fatal("expected no error")
	}

	var nilPW *corepayload.PayloadWrapper
	err = nilPW.SetDynamicPayloads([]byte("data"))

	if err == nil {
		t.Fatal("nil should return error")
	}
}

func Test_Cov9_PayloadWrapper_AttrAsBinder(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	_ = pw.AttrAsBinder()
}

func Test_Cov9_PayloadWrapper_InitializeAttributesOnNull(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	binder := pw.InitializeAttributesOnNull()

	if binder == nil {
		t.Fatal("expected non-nil binder")
	}
}

func Test_Cov9_PayloadWrapper_BasicError(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	be := pw.BasicError()

	if be != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov9_PayloadWrapper_All_AllSafe(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	id, name, entity, cat, payloads := pw.All()
	_, _, _, _, _ = id, name, entity, cat, payloads

	id, name, entity, cat, payloads = pw.AllSafe()
	_, _, _, _, _ = id, name, entity, cat, payloads

	var nilPW *corepayload.PayloadWrapper
	id, name, entity, cat, payloads = nilPW.AllSafe()

	if id != "" || name != "" {
		t.Fatal("nil AllSafe should return empty strings")
	}
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
		t.Fatal("unexpected values")
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

	if props == nil {
		t.Fatal("expected non-nil properties")
	}
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

	if pw.IdString() != "42" {
		t.Fatal("expected 42")
	}

	if pw.IdInteger() != 42 {
		t.Fatal("expected 42")
	}
}

func Test_Cov9_PayloadWrapper_Serialize(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	_, err := pw.Serialize()

	if err != nil {
		t.Fatal("expected no error")
	}

	bytes := pw.SerializeMust()

	if len(bytes) == 0 {
		t.Fatal("expected serialized bytes")
	}
}

func Test_Cov9_PayloadWrapper_Username(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	if pw.Username() != "" {
		t.Fatal("expected empty username")
	}

	user := corepayload.New.User.UsingName("Alice")
	pw.Attributes.AuthInfo = &corepayload.AuthInfo{
		UserInfo: &corepayload.UserInfo{User: user},
	}

	if pw.Username() != "Alice" {
		t.Fatal("expected Alice")
	}
}

func Test_Cov9_PayloadWrapper_Error(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	if pw.Error() != nil {
		t.Fatal("expected nil error")
	}
}

func Test_Cov9_PayloadWrapper_IsEqual_AllBranches(t *testing.T) {
	var nilA, nilB *corepayload.PayloadWrapper

	if !nilA.IsEqual(nilB) {
		t.Fatal("both nil should be equal")
	}

	pw := corepayload.New.PayloadWrapper.Empty()

	if nilA.IsEqual(pw) {
		t.Fatal("nil vs non-nil should not be equal")
	}

	if pw.IsEqual(nilA) {
		t.Fatal("non-nil vs nil should not be equal")
	}

	if !pw.IsEqual(pw) {
		t.Fatal("same pointer should be equal")
	}

	pw2 := corepayload.New.PayloadWrapper.Empty()

	if !pw.IsEqual(pw2) {
		t.Fatal("two empty should be equal")
	}

	// Different Name
	pw3 := corepayload.New.PayloadWrapper.Empty()
	pw3.Name = "different"

	if pw.IsEqual(pw3) {
		t.Fatal("different name should not be equal")
	}

	// Different Identifier
	pw4 := corepayload.New.PayloadWrapper.Empty()
	pw4.Identifier = "diff"

	if pw.IsEqual(pw4) {
		t.Fatal("different identifier should not be equal")
	}

	// Different TaskTypeName
	pw5 := corepayload.New.PayloadWrapper.Empty()
	pw5.TaskTypeName = "diff"

	if pw.IsEqual(pw5) {
		t.Fatal("different task type should not be equal")
	}

	// Different EntityType
	pw6 := corepayload.New.PayloadWrapper.Empty()
	pw6.EntityType = "diff"

	if pw.IsEqual(pw6) {
		t.Fatal("different entity type should not be equal")
	}

	// Different CategoryName
	pw7 := corepayload.New.PayloadWrapper.Empty()
	pw7.CategoryName = "diff"

	if pw.IsEqual(pw7) {
		t.Fatal("different category should not be equal")
	}

	// Different HasManyRecords
	pw8 := corepayload.New.PayloadWrapper.Empty()
	pw8.HasManyRecords = true

	if pw.IsEqual(pw8) {
		t.Fatal("different HasManyRecords should not be equal")
	}
}

func Test_Cov9_PayloadWrapper_IsPayloadsEqual(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("data")}

	if !pw.IsPayloadsEqual([]byte("data")) {
		t.Fatal("expected equal")
	}

	if pw.IsPayloadsEqual([]byte("other")) {
		t.Fatal("expected not equal")
	}

	var nilPW *corepayload.PayloadWrapper

	if nilPW.IsPayloadsEqual([]byte("data")) {
		t.Fatal("nil should not be equal")
	}
}

func Test_Cov9_PayloadWrapper_IsName_IsIdentifier_IsTaskTypeName_IsEntityType_IsCategory(t *testing.T) {
	pw := &corepayload.PayloadWrapper{
		Name:         "n",
		Identifier:   "id",
		TaskTypeName: "task",
		EntityType:   "entity",
		CategoryName: "cat",
	}

	if !pw.IsName("n") {
		t.Fatal("expected true")
	}

	if !pw.IsIdentifier("id") {
		t.Fatal("expected true")
	}

	if !pw.IsTaskTypeName("task") {
		t.Fatal("expected true")
	}

	if !pw.IsEntityType("entity") {
		t.Fatal("expected true")
	}

	if !pw.IsCategory("cat") {
		t.Fatal("expected true")
	}
}

func Test_Cov9_PayloadWrapper_HasAnyItem_HasIssuesOrEmpty(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	if pw.HasAnyItem() {
		t.Fatal("empty should not have any item")
	}

	if !pw.HasIssuesOrEmpty() {
		t.Fatal("empty should have issues or be empty")
	}
}

func Test_Cov9_PayloadWrapper_HasError_IsEmptyError(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	if pw.HasError() {
		t.Fatal("expected no error")
	}

	if !pw.IsEmptyError() {
		t.Fatal("expected empty error")
	}
}

func Test_Cov9_PayloadWrapper_HasAttributes_IsEmptyAttributes(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	if !pw.HasAttributes() {
		t.Fatal("expected attributes")
	}

	pw2 := &corepayload.PayloadWrapper{}

	if !pw2.IsEmptyAttributes() {
		t.Fatal("expected empty attributes")
	}
}

func Test_Cov9_PayloadWrapper_HasSingleRecord(t *testing.T) {
	pw := &corepayload.PayloadWrapper{HasManyRecords: false}

	if !pw.HasSingleRecord() {
		t.Fatal("expected single record")
	}
}

func Test_Cov9_PayloadWrapper_IsNull_HasAnyNil(t *testing.T) {
	var nilPW *corepayload.PayloadWrapper

	if !nilPW.IsNull() {
		t.Fatal("expected null")
	}

	if !nilPW.HasAnyNil() {
		t.Fatal("expected has any nil")
	}
}

func Test_Cov9_PayloadWrapper_Count_Length_IsEmpty_HasItems(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("data")}

	if pw.Count() != 4 {
		t.Fatal("expected 4")
	}

	if pw.Length() != 4 {
		t.Fatal("expected 4")
	}

	if pw.IsEmpty() {
		t.Fatal("expected not empty")
	}

	if !pw.HasItems() {
		t.Fatal("expected has items")
	}

	var nilPW *corepayload.PayloadWrapper

	if nilPW.Length() != 0 {
		t.Fatal("nil Length should be 0")
	}
}

func Test_Cov9_PayloadWrapper_IdentifierInteger_UnsignedInteger(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Identifier: "42"}

	if pw.IdentifierInteger() != 42 {
		t.Fatal("expected 42")
	}

	if pw.IdentifierUnsignedInteger() != 42 {
		t.Fatal("expected 42")
	}

	// Empty identifier
	pw2 := &corepayload.PayloadWrapper{}

	if pw2.IdentifierInteger() >= 0 {
		t.Fatal("expected invalid value")
	}
}

func Test_Cov9_PayloadWrapper_BytesConverter(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("data")}
	bc := pw.BytesConverter()

	if bc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_PayloadWrapper_Deserialize(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`"hello"`)}
	var s string
	err := pw.Deserialize(&s)

	if err != nil || s != "hello" {
		t.Fatal("expected hello")
	}

	err = pw.PayloadDeserialize(&s)

	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_Cov9_PayloadWrapper_MarshalJSON_UnmarshalJSON(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "test"
	pw.Payloads = []byte(`"hello"`)

	jsonBytes, err := pw.MarshalJSON()

	if err != nil {
		t.Fatal("expected no marshal error")
	}

	pw2 := &corepayload.PayloadWrapper{}
	err = pw2.UnmarshalJSON(jsonBytes)

	if err != nil {
		t.Fatal("expected no unmarshal error")
	}

	if pw2.Name != "test" {
		t.Fatal("expected test")
	}
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

	if err != nil {
		t.Fatal("expected no error")
	}

	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_PayloadWrapper_JsonParseSelfInject(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	jsonResult := corejson.NewPtr(corepayload.PayloadWrapper{})
	err := pw.JsonParseSelfInject(jsonResult)

	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_Cov9_PayloadWrapper_PayloadsString(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("hello")}

	if pw.PayloadsString() != "hello" {
		t.Fatal("expected hello")
	}

	emptyPW := &corepayload.PayloadWrapper{}

	if emptyPW.PayloadsString() != "" {
		t.Fatal("expected empty")
	}
}

func Test_Cov9_PayloadWrapper_PayloadsJsonResult(t *testing.T) {
	pw := corepayload.PayloadWrapper{Payloads: []byte(`{"k":"v"}`)}
	result := pw.PayloadsJsonResult()

	if result == nil {
		t.Fatal("expected non-nil")
	}

	emptyPW := corepayload.PayloadWrapper{}
	result = emptyPW.PayloadsJsonResult()

	if result == nil {
		t.Fatal("expected non-nil empty")
	}
}

func Test_Cov9_PayloadWrapper_PayloadsPrettyString(t *testing.T) {
	pw := corepayload.PayloadWrapper{Payloads: []byte(`{"k":"v"}`)}
	result := pw.PayloadsPrettyString()

	if result == "" {
		t.Fatal("expected non-empty")
	}

	emptyPW := corepayload.PayloadWrapper{}

	if emptyPW.PayloadsPrettyString() != "" {
		t.Fatal("expected empty")
	}
}

func Test_Cov9_PayloadWrapper_Clear_Dispose(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Payloads = []byte("data")
	pw.Clear()

	if pw.Length() != 0 {
		t.Fatal("expected 0 length after clear")
	}

	pw2 := corepayload.New.PayloadWrapper.Empty()
	pw2.Dispose()

	if pw2.Attributes != nil {
		t.Fatal("expected nil attributes after dispose")
	}

	var nilPW *corepayload.PayloadWrapper
	nilPW.Clear()
	nilPW.Dispose()
}

func Test_Cov9_PayloadWrapper_AsJsonContractsBinder(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	binder := pw.AsJsonContractsBinder()

	if binder == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_PayloadWrapper_Clone_Shallow(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "test"
	pw.Payloads = []byte("data")
	cloned, err := pw.Clone(false)

	if err != nil || cloned.Name != "test" {
		t.Fatal("expected cloned with name test")
	}
}

func Test_Cov9_PayloadWrapper_Clone_Deep(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "test"
	pw.Payloads = []byte("data")
	cloned, err := pw.Clone(true)

	if err != nil || cloned.Name != "test" {
		t.Fatal("expected cloned with name test")
	}
}

func Test_Cov9_PayloadWrapper_ClonePtr_Nil(t *testing.T) {
	var nilPW *corepayload.PayloadWrapper
	cloned, err := nilPW.ClonePtr(true)

	if err != nil || cloned != nil {
		t.Fatal("nil ClonePtr should return nil, nil")
	}
}

func Test_Cov9_PayloadWrapper_NonPtr_ToPtr(t *testing.T) {
	var nilPW *corepayload.PayloadWrapper
	nonPtr := nilPW.NonPtr()
	_ = nonPtr

	pw := corepayload.PayloadWrapper{Name: "test"}
	ptr := pw.ToPtr()

	if ptr.Name != "test" {
		t.Fatal("expected test")
	}
}

func Test_Cov9_PayloadWrapper_AsStandardTaskEntityDefiner(t *testing.T) {
	pw := corepayload.PayloadWrapper{}
	binder := pw.AsStandardTaskEntityDefinerContractsBinder()

	if binder == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_PayloadWrapper_AsPayloadsBinder(t *testing.T) {
	pw := corepayload.PayloadWrapper{}
	binder := pw.AsPayloadsBinder()

	if binder == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_PayloadWrapper_AsJsonMarshaller(t *testing.T) {
	pw := corepayload.PayloadWrapper{}
	m := pw.AsJsonMarshaller()

	if m == nil {
		t.Fatal("expected non-nil")
	}
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

	if !pw.IsStandardTaskEntityEqual(pw2) {
		t.Fatal("expected equal")
	}
}

func Test_Cov9_PayloadWrapper_SetPayloadDynamic(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	result := pw.SetPayloadDynamic([]byte("data"))

	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_PayloadWrapper_SetPayloadDynamicAny(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	result, err := pw.SetPayloadDynamicAny("hello")

	if err != nil || result == nil {
		t.Fatal("expected no error")
	}
}

func Test_Cov9_PayloadWrapper_SetAuthInfo(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	result := pw.SetAuthInfo(&corepayload.AuthInfo{ActionType: "login"})

	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_PayloadWrapper_SetUserInfo(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	user := corepayload.New.User.UsingName("Alice")
	result := pw.SetUserInfo(&corepayload.UserInfo{User: user})

	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_PayloadWrapper_SetUser_SetSysUser(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	user := corepayload.New.User.UsingName("Alice")
	result := pw.SetUser(user)

	if result == nil {
		t.Fatal("expected non-nil")
	}

	sysUser := corepayload.New.User.System("sys", "system")
	result = pw.SetSysUser(sysUser)

	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_PayloadWrapper_DeserializePayloadsToPayloadsCollection(t *testing.T) {
	// Create a payload wrapper containing serialized collection
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Payloads = []byte(`[]`)
	_, err := pw.DeserializePayloadsToPayloadsCollection()

	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov9_PayloadWrapper_DeserializePayloadsToPayloadWrapper(t *testing.T) {
	inner := corepayload.New.PayloadWrapper.Empty()
	inner.Name = "inner"
	jsonBytes, _ := inner.Serialize()

	pw := &corepayload.PayloadWrapper{Payloads: jsonBytes}
	result, err := pw.DeserializePayloadsToPayloadWrapper()

	if err != nil {
		t.Fatal("expected no error")
	}

	if result.Name != "inner" {
		t.Fatal("expected inner")
	}
}

func Test_Cov9_PayloadWrapper_ReCreateUsingJsonBytes(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "test"
	jsonBytes, _ := pw.Serialize()

	result, err := pw.ReCreateUsingJsonBytes(jsonBytes)

	if err != nil || result.Name != "test" {
		t.Fatal("expected test")
	}
}

func Test_Cov9_PayloadWrapper_ReCreateUsingJsonResult(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "test"
	jsonResult := pw.JsonPtr()

	result, err := pw.ReCreateUsingJsonResult(jsonResult)

	if err != nil || result.Name != "test" {
		t.Fatal("expected test")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadsCollection — Getters
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PayloadsCollection_Length_Count_IsEmpty(t *testing.T) {
	var nilCol *corepayload.PayloadsCollection

	if nilCol.Length() != 0 {
		t.Fatal("nil Length should be 0")
	}

	col := corepayload.New.PayloadsCollection.Empty()

	if col.Count() != 0 || !col.IsEmpty() || col.HasAnyItem() {
		t.Fatal("empty collection state mismatch")
	}
}

func Test_Cov9_PayloadsCollection_LastIndex_HasIndex(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.New.PayloadWrapper.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	_ = pw

	if col.LastIndex() != 0 {
		t.Fatal("expected 0")
	}

	if !col.HasIndex(0) {
		t.Fatal("expected true for index 0")
	}

	if col.HasIndex(1) {
		t.Fatal("expected false for index 1")
	}
}

func Test_Cov9_PayloadsCollection_First_Last(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	if col.First().Name != "a" {
		t.Fatal("expected a")
	}

	if col.Last().Name != "b" {
		t.Fatal("expected b")
	}

	_ = col.FirstDynamic()
	_ = col.LastDynamic()
}

func Test_Cov9_PayloadsCollection_FirstOrDefault_LastOrDefault(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	if col.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}

	if col.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}

	_ = col.FirstOrDefaultDynamic()
	_ = col.LastOrDefaultDynamic()

	col.Add(corepayload.PayloadWrapper{Name: "a"})

	if col.FirstOrDefault().Name != "a" {
		t.Fatal("expected a")
	}

	if col.LastOrDefault().Name != "a" {
		t.Fatal("expected a")
	}
}

func Test_Cov9_PayloadsCollection_Skip_Take_Limit(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})
	col.Add(corepayload.PayloadWrapper{Name: "c"})

	skipped := col.Skip(1)

	if len(skipped) != 2 {
		t.Fatal("expected 2")
	}

	_ = col.SkipDynamic(1)

	skipCol := col.SkipCollection(1)

	if skipCol.Length() != 2 {
		t.Fatal("expected 2")
	}

	taken := col.Take(2)

	if len(taken) != 2 {
		t.Fatal("expected 2")
	}

	_ = col.TakeDynamic(2)

	takeCol := col.TakeCollection(2)

	if takeCol.Length() != 2 {
		t.Fatal("expected 2")
	}

	limitCol := col.LimitCollection(2)

	if limitCol.Length() != 2 {
		t.Fatal("expected 2")
	}

	safeLimitCol := col.SafeLimitCollection(100)

	if safeLimitCol.Length() != 3 {
		t.Fatal("expected 3")
	}

	_ = col.LimitDynamic(2)
	_ = col.Limit(2)
}

func Test_Cov9_PayloadsCollection_Strings(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})

	strings := col.Strings()

	if len(strings) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_PayloadsCollection_IsEqual(t *testing.T) {
	var nilA, nilB *corepayload.PayloadsCollection

	if !nilA.IsEqual(nilB) {
		t.Fatal("both nil should be equal")
	}

	col := corepayload.New.PayloadsCollection.Empty()

	if nilA.IsEqual(col) {
		t.Fatal("nil vs non-nil should not be equal")
	}

	col2 := corepayload.New.PayloadsCollection.Empty()

	if !col.IsEqual(col2) {
		t.Fatal("two empty should be equal")
	}
}

func Test_Cov9_PayloadsCollection_IsEqualItems(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.New.PayloadWrapper.Empty()
	col.AddsPtr(pw)

	if !col.IsEqualItems(pw) {
		t.Fatal("expected equal")
	}
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

	if len(filtered) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_PayloadsCollection_FilterWithLimit(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})
	col.Add(corepayload.PayloadWrapper{Name: "c"})

	filtered := col.FilterWithLimit(1, func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})

	if len(filtered) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_PayloadsCollection_FirstByFilter(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	found := col.FirstByFilter(func(pw *corepayload.PayloadWrapper) bool {
		return pw.Name == "b"
	})

	if found == nil || found.Name != "b" {
		t.Fatal("expected b")
	}

	notFound := col.FirstByFilter(func(pw *corepayload.PayloadWrapper) bool {
		return pw.Name == "z"
	})

	if notFound != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov9_PayloadsCollection_FirstById(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Identifier: "id1", Name: "a"})
	col.Add(corepayload.PayloadWrapper{Identifier: "id2", Name: "b"})

	found := col.FirstById("id2")

	if found == nil || found.Name != "b" {
		t.Fatal("expected b")
	}
}

func Test_Cov9_PayloadsCollection_FirstByCategory(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{CategoryName: "cat1", Name: "a"})

	found := col.FirstByCategory("cat1")

	if found == nil || found.Name != "a" {
		t.Fatal("expected a")
	}
}

func Test_Cov9_PayloadsCollection_FirstByTaskType(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{TaskTypeName: "task1", Name: "a"})

	found := col.FirstByTaskType("task1")

	if found == nil || found.Name != "a" {
		t.Fatal("expected a")
	}
}

func Test_Cov9_PayloadsCollection_FirstByEntityType(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{EntityType: "entity1", Name: "a"})

	found := col.FirstByEntityType("entity1")

	if found == nil || found.Name != "a" {
		t.Fatal("expected a")
	}
}

func Test_Cov9_PayloadsCollection_FilterCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	filtered := col.FilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return pw.Name == "a", false
	})

	if filtered.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_PayloadsCollection_SkipFilterCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	filtered := col.SkipFilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return pw.Name == "a", false
	})

	if filtered.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_PayloadsCollection_FilterCollectionByIds(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Identifier: "id1", Name: "a"})
	col.Add(corepayload.PayloadWrapper{Identifier: "id2", Name: "b"})
	col.Add(corepayload.PayloadWrapper{Identifier: "id3", Name: "c"})

	filtered := col.FilterCollectionByIds("id1", "id3")

	if filtered.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_Cov9_PayloadsCollection_FilterNameCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	filtered := col.FilterNameCollection("a")

	if filtered.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_PayloadsCollection_FilterCategoryCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{CategoryName: "cat1"})

	filtered := col.FilterCategoryCollection("cat1")

	if filtered.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_PayloadsCollection_FilterEntityTypeCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{EntityType: "entity1"})

	filtered := col.FilterEntityTypeCollection("entity1")

	if filtered.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_PayloadsCollection_FilterTaskTypeCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{TaskTypeName: "task1"})

	filtered := col.FilterTaskTypeCollection("task1")

	if filtered.Length() != 1 {
		t.Fatal("expected 1")
	}
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

	if pages != 3 {
		t.Fatalf("expected 3 pages, got %d", pages)
	}

	if col.GetPagesSize(0) != 0 {
		t.Fatal("zero page size should return 0")
	}

	if col.GetPagesSize(-1) != 0 {
		t.Fatal("negative page size should return 0")
	}
}

func Test_Cov9_PayloadsCollection_GetPagedCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	for i := 0; i < 25; i++ {
		col.Add(corepayload.PayloadWrapper{Name: "item"})
	}

	pages := col.GetPagedCollection(10)

	if len(pages) != 3 {
		t.Fatalf("expected 3 pages, got %d", len(pages))
	}

	// Smaller than page size
	small := corepayload.New.PayloadsCollection.Empty()
	small.Add(corepayload.PayloadWrapper{Name: "a"})
	pages = small.GetPagedCollection(10)

	if len(pages) != 1 {
		t.Fatal("expected 1 page")
	}
}

func Test_Cov9_PayloadsCollection_GetSinglePageCollection(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	for i := 0; i < 25; i++ {
		col.Add(corepayload.PayloadWrapper{Name: "item"})
	}

	page := col.GetSinglePageCollection(10, 1)

	if page.Length() != 10 {
		t.Fatalf("expected 10, got %d", page.Length())
	}

	page3 := col.GetSinglePageCollection(10, 3)

	if page3.Length() != 5 {
		t.Fatalf("expected 5, got %d", page3.Length())
	}

	// Smaller than page size
	small := corepayload.New.PayloadsCollection.Empty()
	small.Add(corepayload.PayloadWrapper{Name: "a"})
	page = small.GetSinglePageCollection(10, 1)

	if page.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadsCollection — Mutation
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PayloadsCollection_Add_Adds(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Adds(corepayload.PayloadWrapper{Name: "b"}, corepayload.PayloadWrapper{Name: "c"})

	if col.Length() != 3 {
		t.Fatal("expected 3")
	}

	// Adds empty
	col.Adds()

	if col.Length() != 3 {
		t.Fatal("expected still 3")
	}
}

func Test_Cov9_PayloadsCollection_AddsPtr(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.New.PayloadWrapper.Empty()
	col.AddsPtr(pw)

	if col.Length() != 1 {
		t.Fatal("expected 1")
	}

	// Empty
	col.AddsPtr()

	if col.Length() != 1 {
		t.Fatal("expected still 1")
	}
}

func Test_Cov9_PayloadsCollection_AddsPtrOptions(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Payloads = []byte("data")
	col.AddsPtrOptions(false, pw)

	if col.Length() != 1 {
		t.Fatal("expected 1")
	}

	// Skip issues
	emptyPW := corepayload.New.PayloadWrapper.Empty()
	col2 := corepayload.New.PayloadsCollection.Empty()
	col2.AddsPtrOptions(true, emptyPW)

	if col2.Length() != 0 {
		t.Fatal("expected 0 (skipped)")
	}
}

func Test_Cov9_PayloadsCollection_AddsOptions(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.AddsOptions(false, corepayload.PayloadWrapper{Payloads: []byte("data")})

	if col.Length() != 1 {
		t.Fatal("expected 1")
	}

	// Skip issues
	col2 := corepayload.New.PayloadsCollection.Empty()
	col2.AddsOptions(true, corepayload.PayloadWrapper{})

	if col2.Length() != 0 {
		t.Fatal("expected 0 (skipped)")
	}
}

func Test_Cov9_PayloadsCollection_AddsIf(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.AddsIf(true, corepayload.PayloadWrapper{Name: "a"})

	if col.Length() != 1 {
		t.Fatal("expected 1")
	}

	col.AddsIf(false, corepayload.PayloadWrapper{Name: "b"})

	if col.Length() != 1 {
		t.Fatal("expected still 1")
	}
}

func Test_Cov9_PayloadsCollection_InsertAt(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "c"})
	col.InsertAt(1, corepayload.PayloadWrapper{Name: "b"})

	if col.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_Cov9_PayloadsCollection_ConcatNew(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	newCol := col.ConcatNew(corepayload.PayloadWrapper{Name: "b"})

	if newCol.Length() != 2 {
		t.Fatal("expected 2")
	}

	if col.Length() != 1 {
		t.Fatal("original should still be 1")
	}
}

func Test_Cov9_PayloadsCollection_ConcatNewPtr(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Name = "b"
	newCol := col.ConcatNewPtr(pw)

	if newCol.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_Cov9_PayloadsCollection_Reverse(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})
	col.Add(corepayload.PayloadWrapper{Name: "c"})
	col.Reverse()

	if col.First().Name != "c" || col.Last().Name != "a" {
		t.Fatal("expected reversed")
	}

	// Reverse 2 items
	col2 := corepayload.New.PayloadsCollection.Empty()
	col2.Add(corepayload.PayloadWrapper{Name: "a"})
	col2.Add(corepayload.PayloadWrapper{Name: "b"})
	col2.Reverse()

	if col2.First().Name != "b" {
		t.Fatal("expected b first")
	}

	// Reverse 1 item
	col3 := corepayload.New.PayloadsCollection.Empty()
	col3.Add(corepayload.PayloadWrapper{Name: "a"})
	col3.Reverse()

	if col3.First().Name != "a" {
		t.Fatal("expected a")
	}
}

func Test_Cov9_PayloadsCollection_Clone_ClonePtr(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})

	cloned := col.Clone()

	if cloned.Length() != 1 {
		t.Fatal("expected 1")
	}

	clonedPtr := col.ClonePtr()

	if clonedPtr.Length() != 1 {
		t.Fatal("expected 1")
	}

	// nil ClonePtr
	var nilCol *corepayload.PayloadsCollection

	if nilCol.ClonePtr() != nil {
		t.Fatal("nil ClonePtr should return nil")
	}
}

func Test_Cov9_PayloadsCollection_Clear_Dispose(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Clear()

	if col.Length() != 0 {
		t.Fatal("expected 0 after clear")
	}

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

	if len(strings) != 1 || strings[0] != "a" {
		t.Fatal("expected a")
	}
}

func Test_Cov9_PayloadsCollection_JoinUsingFmt(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a"})
	col.Add(corepayload.PayloadWrapper{Name: "b"})

	result := col.JoinUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.Name
	}, ",")

	if result != "a,b" {
		t.Fatal("expected a,b")
	}
}

func Test_Cov9_PayloadsCollection_JsonStrings_JoinJsonStrings(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("data")})

	jsonStrings := col.JsonStrings()

	if len(jsonStrings) != 1 {
		t.Fatal("expected 1")
	}

	_ = col.JoinJsonStrings(",")
}

func Test_Cov9_PayloadsCollection_Join(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	col.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("data")})

	_ = col.Join(",")
}

func Test_Cov9_PayloadsCollection_JsonString_String(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	if col.JsonString() != "" {
		t.Fatal("empty should return empty string")
	}

	if col.String() != "" {
		t.Fatal("empty should return empty string")
	}

	col.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("data")})

	if col.JsonString() == "" {
		t.Fatal("expected non-empty")
	}

	if col.String() == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov9_PayloadsCollection_PrettyJsonString(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	if col.PrettyJsonString() != "" {
		t.Fatal("empty should return empty string")
	}

	col.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("data")})

	if col.PrettyJsonString() == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov9_PayloadsCollection_CsvStrings_JoinCsv_JoinCsvLine(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	if len(col.CsvStrings()) != 0 {
		t.Fatal("empty csv should be empty")
	}

	col.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte("data")})

	csvStrings := col.CsvStrings()

	if len(csvStrings) != 1 {
		t.Fatal("expected 1")
	}

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

	if err != nil || result == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_PayloadsCollection_AsJsonContractsBinder(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	binder := col.AsJsonContractsBinder()

	if binder == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_PayloadsCollection_AsJsoner(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	jsoner := col.AsJsoner()

	if jsoner == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_PayloadsCollection_JsonParseSelfInject(t *testing.T) {
	col := &corepayload.PayloadsCollection{}
	jsonResult := corejson.NewPtr(corepayload.PayloadsCollection{})
	err := col.JsonParseSelfInject(jsonResult)

	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_Cov9_PayloadsCollection_AsJsonParseSelfInjector(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()
	injector := col.AsJsonParseSelfInjector()

	if injector == nil {
		t.Fatal("expected non-nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// User — Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_User_IdentifierInteger(t *testing.T) {
	u := corepayload.User{Identifier: "42"}

	if u.IdentifierInteger() != 42 {
		t.Fatal("expected 42")
	}

	u2 := corepayload.User{}

	if u2.IdentifierInteger() >= 0 {
		t.Fatal("expected invalid value")
	}
}

func Test_Cov9_User_IdentifierUnsignedInteger(t *testing.T) {
	u := corepayload.User{Identifier: "42"}

	if u.IdentifierUnsignedInteger() != 42 {
		t.Fatal("expected 42")
	}

	u2 := corepayload.User{Identifier: "-1"}

	if u2.IdentifierUnsignedInteger() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_Cov9_User_AllBoolMethods(t *testing.T) {
	u := corepayload.New.User.All(true, "1", "Alice", "admin", "token", "hash")

	if !u.HasAuthToken() {
		t.Fatal("expected HasAuthToken")
	}

	if !u.HasPasswordHash() {
		t.Fatal("expected HasPasswordHash")
	}

	if u.IsPasswordHashEmpty() {
		t.Fatal("expected not empty")
	}

	if u.IsAuthTokenEmpty() {
		t.Fatal("expected not empty")
	}

	if u.IsEmpty() {
		t.Fatal("expected not empty")
	}

	if !u.IsValidUser() {
		t.Fatal("expected valid")
	}

	if u.IsNameEmpty() {
		t.Fatal("expected not empty")
	}

	if !u.IsNameEqual("Alice") {
		t.Fatal("expected equal")
	}

	if u.IsNotSystemUser() {
		t.Fatal("expected system user")
	}

	if u.IsVirtualUser() {
		t.Fatal("expected not virtual user (is system)")
	}

	if !u.HasType() {
		t.Fatal("expected has type")
	}

	if u.IsTypeEmpty() {
		t.Fatal("expected not empty")
	}

	// nil receiver
	var nilUser *corepayload.User

	if nilUser.HasAuthToken() || nilUser.HasPasswordHash() {
		t.Fatal("nil should return false")
	}

	if !nilUser.IsPasswordHashEmpty() || !nilUser.IsAuthTokenEmpty() {
		t.Fatal("nil should be empty")
	}

	if !nilUser.IsEmpty() || nilUser.IsValidUser() || !nilUser.IsNameEmpty() {
		t.Fatal("nil checks failed")
	}

	if nilUser.IsNameEqual("anything") {
		t.Fatal("nil should not be equal")
	}

	if nilUser.IsNotSystemUser() || nilUser.IsVirtualUser() || nilUser.HasType() {
		t.Fatal("nil should return false")
	}

	if !nilUser.IsTypeEmpty() {
		t.Fatal("nil should be type empty")
	}
}

func Test_Cov9_User_String_Json_Serialize_Deserialize(t *testing.T) {
	u := corepayload.New.User.UsingName("Alice")
	_ = u.String()
	_ = u.PrettyJsonString()
	_ = u.Json()
	_ = u.JsonPtr()

	serialized, err := u.Serialize()

	if err != nil {
		t.Fatal("expected no error")
	}

	u2 := &corepayload.User{}
	err = u2.Deserialize(serialized)

	if err != nil || u2.Name != "Alice" {
		t.Fatal("expected Alice")
	}
}

func Test_Cov9_User_Clone_ClonePtr(t *testing.T) {
	u := corepayload.New.User.All(false, "1", "Alice", "admin", "token", "hash")
	cloned := u.Clone()

	if cloned.Name != "Alice" {
		t.Fatal("expected Alice")
	}

	clonedPtr := u.ClonePtr()

	if clonedPtr.Name != "Alice" {
		t.Fatal("expected Alice")
	}

	var nilUser *corepayload.User

	if nilUser.ClonePtr() != nil {
		t.Fatal("nil ClonePtr should return nil")
	}
}

func Test_Cov9_User_Ptr(t *testing.T) {
	u := corepayload.User{Name: "Alice"}
	ptr := u.Ptr()

	if ptr.Name != "Alice" {
		t.Fatal("expected Alice")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// UserInfo — Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_UserInfo_All(t *testing.T) {
	ui := &corepayload.UserInfo{
		User:       corepayload.New.User.UsingName("Alice"),
		SystemUser: corepayload.New.User.System("sys", "system"),
	}

	if !ui.HasUser() {
		t.Fatal("expected HasUser")
	}

	if !ui.HasSystemUser() {
		t.Fatal("expected HasSystemUser")
	}

	if ui.IsEmpty() {
		t.Fatal("expected not empty")
	}

	if ui.IsUserEmpty() {
		t.Fatal("expected not empty")
	}

	if ui.IsSystemUserEmpty() {
		t.Fatal("expected not empty")
	}

	// nil receiver
	var nilUI *corepayload.UserInfo

	if nilUI.HasUser() || nilUI.HasSystemUser() {
		t.Fatal("nil should return false")
	}

	if !nilUI.IsEmpty() || !nilUI.IsUserEmpty() || !nilUI.IsSystemUserEmpty() {
		t.Fatal("nil should be empty")
	}
}

func Test_Cov9_UserInfo_SetUserSystemUser(t *testing.T) {
	ui := &corepayload.UserInfo{}
	user := corepayload.New.User.UsingName("Alice")
	sysUser := corepayload.New.User.System("sys", "system")
	result := ui.SetUserSystemUser(user, sysUser)

	if result.User.Name != "Alice" || result.SystemUser.Name != "sys" {
		t.Fatal("expected Alice and sys")
	}

	// nil receiver
	var nilUI *corepayload.UserInfo
	result = nilUI.SetUserSystemUser(user, sysUser)

	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_UserInfo_SetUser_SetSystemUser(t *testing.T) {
	ui := &corepayload.UserInfo{}
	user := corepayload.New.User.UsingName("Alice")
	result := ui.SetUser(user)

	if result.User.Name != "Alice" {
		t.Fatal("expected Alice")
	}

	sysUser := corepayload.New.User.System("sys", "system")
	result = ui.SetSystemUser(sysUser)

	if result.SystemUser.Name != "sys" {
		t.Fatal("expected sys")
	}

	// nil receiver
	var nilUI *corepayload.UserInfo
	result = nilUI.SetUser(user)

	if result == nil {
		t.Fatal("expected non-nil")
	}

	result = nilUI.SetSystemUser(sysUser)

	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_UserInfo_Clone_ClonePtr_Ptr(t *testing.T) {
	ui := &corepayload.UserInfo{User: corepayload.New.User.UsingName("Alice")}
	cloned := ui.Clone()

	if cloned.User.Name != "Alice" {
		t.Fatal("expected Alice")
	}

	clonedPtr := ui.ClonePtr()

	if clonedPtr.User.Name != "Alice" {
		t.Fatal("expected Alice")
	}

	var nilUI *corepayload.UserInfo

	if nilUI.ClonePtr() != nil {
		t.Fatal("nil ClonePtr should return nil")
	}

	uiVal := corepayload.UserInfo{User: corepayload.New.User.UsingName("Bob")}
	ptr := uiVal.Ptr()

	if ptr.User.Name != "Bob" {
		t.Fatal("expected Bob")
	}
}

func Test_Cov9_UserInfo_ToNonPtr(t *testing.T) {
	ui := &corepayload.UserInfo{User: corepayload.New.User.UsingName("Alice")}
	nonPtr := ui.ToNonPtr()

	if nonPtr.User.Name != "Alice" {
		t.Fatal("expected Alice")
	}

	var nilUI *corepayload.UserInfo
	nonPtr = nilUI.ToNonPtr()
	_ = nonPtr
}

// ══════════════════════════════════════════════════════════════════════════════
// AuthInfo — Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_AuthInfo_IdentifierInteger(t *testing.T) {
	ai := corepayload.AuthInfo{Identifier: "42"}

	if ai.IdentifierInteger() != 42 {
		t.Fatal("expected 42")
	}

	ai2 := corepayload.AuthInfo{}

	if ai2.IdentifierInteger() >= 0 {
		t.Fatal("expected invalid value")
	}
}

func Test_Cov9_AuthInfo_IdentifierUnsignedInteger(t *testing.T) {
	ai := corepayload.AuthInfo{Identifier: "42"}

	if ai.IdentifierUnsignedInteger() != 42 {
		t.Fatal("expected 42")
	}

	ai2 := corepayload.AuthInfo{Identifier: "-1"}

	if ai2.IdentifierUnsignedInteger() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_Cov9_AuthInfo_IsEmpty_HasAnyItem_IsValid(t *testing.T) {
	var nilAI *corepayload.AuthInfo

	if !nilAI.IsEmpty() {
		t.Fatal("nil should be empty")
	}

	ai := &corepayload.AuthInfo{ActionType: "login"}

	if ai.IsEmpty() {
		t.Fatal("expected not empty")
	}

	if !ai.HasAnyItem() {
		t.Fatal("expected has any item")
	}

	if !ai.IsValid() {
		t.Fatal("expected valid")
	}
}

func Test_Cov9_AuthInfo_IsActionTypeEmpty_IsResourceNameEmpty(t *testing.T) {
	var nilAI *corepayload.AuthInfo

	if !nilAI.IsActionTypeEmpty() {
		t.Fatal("nil should be empty")
	}

	if !nilAI.IsResourceNameEmpty() {
		t.Fatal("nil should be empty")
	}

	ai := &corepayload.AuthInfo{ActionType: "login", ResourceName: "/api"}

	if ai.IsActionTypeEmpty() {
		t.Fatal("expected not empty")
	}

	if ai.IsResourceNameEmpty() {
		t.Fatal("expected not empty")
	}
}

func Test_Cov9_AuthInfo_HasActionType_HasResourceName(t *testing.T) {
	ai := &corepayload.AuthInfo{ActionType: "login", ResourceName: "/api"}

	if !ai.HasActionType() {
		t.Fatal("expected true")
	}

	if !ai.HasResourceName() {
		t.Fatal("expected true")
	}

	var nilAI *corepayload.AuthInfo

	if nilAI.HasActionType() || nilAI.HasResourceName() {
		t.Fatal("nil should return false")
	}
}

func Test_Cov9_AuthInfo_IsUserInfoEmpty_IsSessionInfoEmpty(t *testing.T) {
	var nilAI *corepayload.AuthInfo

	if !nilAI.IsUserInfoEmpty() || !nilAI.IsSessionInfoEmpty() {
		t.Fatal("nil should be empty")
	}

	ai := &corepayload.AuthInfo{}

	if !ai.IsUserInfoEmpty() || !ai.IsSessionInfoEmpty() {
		t.Fatal("empty should be empty")
	}
}

func Test_Cov9_AuthInfo_HasUserInfo_HasSessionInfo(t *testing.T) {
	ai := &corepayload.AuthInfo{
		UserInfo:    &corepayload.UserInfo{User: corepayload.New.User.UsingName("Alice")},
		SessionInfo: &corepayload.SessionInfo{Id: "s1", User: corepayload.New.User.UsingName("Bob")},
	}

	if !ai.HasUserInfo() {
		t.Fatal("expected true")
	}

	if !ai.HasSessionInfo() {
		t.Fatal("expected true")
	}
}

func Test_Cov9_AuthInfo_SetUserInfo_Nil(t *testing.T) {
	var nilAI *corepayload.AuthInfo
	result := nilAI.SetUserInfo(&corepayload.UserInfo{})

	if result == nil {
		t.Fatal("expected non-nil")
	}

	ai := &corepayload.AuthInfo{}
	result = ai.SetUserInfo(&corepayload.UserInfo{})

	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_AuthInfo_SetActionType_SetResourceName_SetIdentifier_SetSessionInfo(t *testing.T) {
	var nilAI *corepayload.AuthInfo

	result := nilAI.SetActionType("login")

	if result == nil || result.ActionType != "login" {
		t.Fatal("expected login")
	}

	result = nilAI.SetResourceName("/api")

	if result == nil || result.ResourceName != "/api" {
		t.Fatal("expected /api")
	}

	result = nilAI.SetIdentifier("42")

	if result == nil || result.Identifier != "42" {
		t.Fatal("expected 42")
	}

	result = nilAI.SetSessionInfo(&corepayload.SessionInfo{Id: "s1"})

	if result == nil || result.SessionInfo.Id != "s1" {
		t.Fatal("expected s1")
	}

	// non-nil receiver
	ai := &corepayload.AuthInfo{}
	ai.SetActionType("test")

	if ai.ActionType != "test" {
		t.Fatal("expected test")
	}

	ai.SetResourceName("/resource")

	if ai.ResourceName != "/resource" {
		t.Fatal("expected /resource")
	}

	ai.SetIdentifier("id")

	if ai.Identifier != "id" {
		t.Fatal("expected id")
	}

	ai.SetSessionInfo(&corepayload.SessionInfo{Id: "s2"})

	if ai.SessionInfo.Id != "s2" {
		t.Fatal("expected s2")
	}
}

func Test_Cov9_AuthInfo_SetUserSystemUser(t *testing.T) {
	var nilAI *corepayload.AuthInfo
	user := corepayload.New.User.UsingName("Alice")
	sysUser := corepayload.New.User.System("sys", "system")

	result := nilAI.SetUserSystemUser(user, sysUser)

	if result == nil {
		t.Fatal("expected non-nil")
	}

	ai := &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{}}
	result = ai.SetUserSystemUser(user, sysUser)

	if result.UserInfo.User.Name != "Alice" {
		t.Fatal("expected Alice")
	}
}

func Test_Cov9_AuthInfo_SetUser_SetSystemUser(t *testing.T) {
	var nilAI *corepayload.AuthInfo
	user := corepayload.New.User.UsingName("Alice")

	result := nilAI.SetUser(user)

	if result == nil {
		t.Fatal("expected non-nil")
	}

	sysUser := corepayload.New.User.System("sys", "system")
	result = nilAI.SetSystemUser(sysUser)

	if result == nil {
		t.Fatal("expected non-nil")
	}

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

	if cloned.ActionType != "login" {
		t.Fatal("expected login")
	}

	ptr := ai.Ptr()

	if ptr.ActionType != "login" {
		t.Fatal("expected login")
	}

	clonedPtr := ptr.ClonePtr()

	if clonedPtr.ActionType != "login" {
		t.Fatal("expected login")
	}

	var nilAI *corepayload.AuthInfo

	if nilAI.ClonePtr() != nil {
		t.Fatal("nil ClonePtr should return nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SessionInfo — Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_SessionInfo_IdentifierInteger(t *testing.T) {
	si := corepayload.SessionInfo{Id: "42"}

	if si.IdentifierInteger() != 42 {
		t.Fatal("expected 42")
	}

	si2 := corepayload.SessionInfo{}

	if si2.IdentifierInteger() >= 0 {
		t.Fatal("expected invalid value")
	}
}

func Test_Cov9_SessionInfo_IdentifierUnsignedInteger(t *testing.T) {
	si := corepayload.SessionInfo{Id: "42"}

	if si.IdentifierUnsignedInteger() != 42 {
		t.Fatal("expected 42")
	}

	si2 := corepayload.SessionInfo{Id: "-1"}

	if si2.IdentifierUnsignedInteger() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_Cov9_SessionInfo_IsEmpty_IsValid(t *testing.T) {
	var nilSI *corepayload.SessionInfo

	if !nilSI.IsEmpty() {
		t.Fatal("nil should be empty")
	}

	si := &corepayload.SessionInfo{Id: "s1", User: corepayload.New.User.UsingName("Alice")}

	if si.IsEmpty() {
		t.Fatal("expected not empty")
	}

	if !si.IsValid() {
		t.Fatal("expected valid")
	}

	emptySI := &corepayload.SessionInfo{}

	if emptySI.IsValid() {
		t.Fatal("expected invalid")
	}
}

func Test_Cov9_SessionInfo_IsUserNameEmpty_IsUserEmpty_HasUser(t *testing.T) {
	var nilSI *corepayload.SessionInfo

	if !nilSI.IsUserNameEmpty() || !nilSI.IsUserEmpty() || nilSI.HasUser() {
		t.Fatal("nil checks failed")
	}

	si := &corepayload.SessionInfo{Id: "s1", User: corepayload.New.User.UsingName("Alice")}

	if si.IsUserNameEmpty() || si.IsUserEmpty() || !si.HasUser() {
		t.Fatal("user checks failed")
	}
}

func Test_Cov9_SessionInfo_IsUsernameEqual(t *testing.T) {
	si := &corepayload.SessionInfo{User: corepayload.New.User.UsingName("Alice")}

	if !si.IsUsernameEqual("Alice") {
		t.Fatal("expected equal")
	}

	if si.IsUsernameEqual("Bob") {
		t.Fatal("expected not equal")
	}

	var nilSI *corepayload.SessionInfo

	if nilSI.IsUsernameEqual("Alice") {
		t.Fatal("nil should return false")
	}
}

func Test_Cov9_SessionInfo_Clone_ClonePtr_Ptr(t *testing.T) {
	si := corepayload.SessionInfo{Id: "s1", User: corepayload.New.User.UsingName("Alice"), SessionPath: "/path"}
	cloned := si.Clone()

	if cloned.Id != "s1" || cloned.SessionPath != "/path" {
		t.Fatal("expected cloned values")
	}

	ptr := si.Ptr()

	if ptr.Id != "s1" {
		t.Fatal("expected s1")
	}

	clonedPtr := ptr.ClonePtr()

	if clonedPtr.Id != "s1" {
		t.Fatal("expected s1")
	}

	var nilSI *corepayload.SessionInfo

	if nilSI.ClonePtr() != nil {
		t.Fatal("nil ClonePtr should return nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// PagingInfo — Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PagingInfo_IsEmpty(t *testing.T) {
	var nilPI *corepayload.PagingInfo

	if !nilPI.IsEmpty() {
		t.Fatal("nil should be empty")
	}

	pi := &corepayload.PagingInfo{TotalPages: 5, TotalItems: 50}

	if pi.IsEmpty() {
		t.Fatal("expected not empty")
	}
}

func Test_Cov9_PagingInfo_IsEqual_AllBranches(t *testing.T) {
	var nilA, nilB *corepayload.PagingInfo

	if !nilA.IsEqual(nilB) {
		t.Fatal("both nil should be equal")
	}

	pi := &corepayload.PagingInfo{TotalPages: 5}

	if nilA.IsEqual(pi) || pi.IsEqual(nilA) {
		t.Fatal("nil vs non-nil should not be equal")
	}

	pi2 := &corepayload.PagingInfo{TotalPages: 3}

	if pi.IsEqual(pi2) {
		t.Fatal("different TotalPages should not be equal")
	}

	pi3 := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 1}

	if pi.IsEqual(pi3) {
		t.Fatal("different CurrentPageIndex should not be equal")
	}

	pi4 := &corepayload.PagingInfo{TotalPages: 5, PerPageItems: 10}

	if pi.IsEqual(pi4) {
		t.Fatal("different PerPageItems should not be equal")
	}

	pi5 := &corepayload.PagingInfo{TotalPages: 5, TotalItems: 50}

	if pi.IsEqual(pi5) {
		t.Fatal("different TotalItems should not be equal")
	}
}

func Test_Cov9_PagingInfo_HasMethods(t *testing.T) {
	pi := &corepayload.PagingInfo{
		TotalPages:       5,
		CurrentPageIndex: 2,
		PerPageItems:     10,
		TotalItems:       50,
	}

	if !pi.HasTotalPages() || !pi.HasCurrentPageIndex() || !pi.HasPerPageItems() || !pi.HasTotalItems() {
		t.Fatal("expected all true")
	}

	var nilPI *corepayload.PagingInfo

	if nilPI.HasTotalPages() || nilPI.HasCurrentPageIndex() || nilPI.HasPerPageItems() || nilPI.HasTotalItems() {
		t.Fatal("nil should return false")
	}
}

func Test_Cov9_PagingInfo_IsInvalidMethods(t *testing.T) {
	pi := &corepayload.PagingInfo{}

	if !pi.IsInvalidTotalPages() || !pi.IsInvalidCurrentPageIndex() || !pi.IsInvalidPerPageItems() || !pi.IsInvalidTotalItems() {
		t.Fatal("zero values should be invalid")
	}

	var nilPI *corepayload.PagingInfo

	if !nilPI.IsInvalidTotalPages() || !nilPI.IsInvalidCurrentPageIndex() || !nilPI.IsInvalidPerPageItems() || !nilPI.IsInvalidTotalItems() {
		t.Fatal("nil should be invalid")
	}
}

func Test_Cov9_PagingInfo_Clone_ClonePtr(t *testing.T) {
	pi := corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 2, PerPageItems: 10, TotalItems: 50}
	cloned := pi.Clone()

	if cloned.TotalPages != 5 {
		t.Fatal("expected 5")
	}

	ptr := &pi
	clonedPtr := ptr.ClonePtr()

	if clonedPtr.TotalPages != 5 {
		t.Fatal("expected 5")
	}

	var nilPI *corepayload.PagingInfo

	if nilPI.ClonePtr() != nil {
		t.Fatal("nil ClonePtr should return nil")
	}
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

	if props.Name() != "test" {
		t.Fatal("expected test")
	}

	if props.IdString() != "42" {
		t.Fatal("expected 42")
	}

	if props.Category() != "cat" {
		t.Fatal("expected cat")
	}

	if props.EntityType() != "entity" {
		t.Fatal("expected entity")
	}

	if !props.HasManyRecord() {
		t.Fatal("expected true")
	}

	if props.HasSingleRecordOnly() {
		t.Fatal("expected false")
	}

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

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	attrDefaults := corepayload.Empty.AttributesDefaults()

	if attrDefaults == nil {
		t.Fatal("expected non-nil")
	}

	pw := corepayload.Empty.PayloadWrapper()

	if pw == nil {
		t.Fatal("expected non-nil")
	}

	col := corepayload.Empty.PayloadsCollection()

	if col == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_NewAttributesCreator_AllFactories(t *testing.T) {
	// Create
	attr := corepayload.New.Attributes.Create(nil, nil, []byte("data"))

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// ErrFromTo
	attr = corepayload.New.Attributes.ErrFromTo(nil, nil, []byte("data"))

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// UsingAuthInfoDynamicBytes
	attr = corepayload.New.Attributes.UsingAuthInfoDynamicBytes(nil, []byte("data"))

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// UsingDynamicPayloadBytes
	attr = corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// UsingAuthInfo
	attr = corepayload.New.Attributes.UsingAuthInfo(nil)

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// UsingKeyValues
	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	attr = corepayload.New.Attributes.UsingKeyValues(hm)

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// UsingAuthInfoKeyValues
	attr = corepayload.New.Attributes.UsingAuthInfoKeyValues(nil, hm)

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// UsingKeyValuesPlusDynamic
	attr = corepayload.New.Attributes.UsingKeyValuesPlusDynamic(hm, []byte("data"))

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// UsingAnyKeyValues
	anyMap := coredynamic.NewMapAnyItems(0)
	attr = corepayload.New.Attributes.UsingAnyKeyValues(anyMap)

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// UsingAuthInfoAnyKeyValues
	attr = corepayload.New.Attributes.UsingAuthInfoAnyKeyValues(nil, anyMap)

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// UsingAnyKeyValuesPlusDynamic
	attr = corepayload.New.Attributes.UsingAnyKeyValuesPlusDynamic(anyMap, []byte("data"))

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// UsingBasicError
	attr = corepayload.New.Attributes.UsingBasicError(nil)

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// Empty
	attr = corepayload.New.Attributes.Empty()

	if attr == nil {
		t.Fatal("expected non-nil")
	}

	// All
	attr = corepayload.New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)

	if attr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_NewAttributesCreator_AllAny(t *testing.T) {
	attr, err := corepayload.New.Attributes.AllAny(nil, nil, nil, nil, "test")

	if err != nil || attr == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewAttributesCreator_PageInfoAny(t *testing.T) {
	attr, err := corepayload.New.Attributes.PageInfoAny(nil, "test")

	if err != nil || attr == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewAttributesCreator_UsingDynamicPayloadAny(t *testing.T) {
	attr, err := corepayload.New.Attributes.UsingDynamicPayloadAny(nil, "test")

	if err != nil || attr == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewAttributesCreator_UsingAuthInfoJsonResult(t *testing.T) {
	jsonResult := corejson.NewPtr("test")
	attr, err := corepayload.New.Attributes.UsingAuthInfoJsonResult(nil, jsonResult)

	if err != nil || attr == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewAttributesCreator_Deserialize(t *testing.T) {
	original := corepayload.New.Attributes.Empty()
	bytes := []byte(original.JsonString())
	attr, err := corepayload.New.Attributes.Deserialize(bytes)

	if err != nil || attr == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewAttributesCreator_DeserializeMany(t *testing.T) {
	bytes := []byte(`[{}]`)
	attrs, err := corepayload.New.Attributes.DeserializeMany(bytes)

	if err != nil || len(attrs) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_NewAttributesCreator_DeserializeUsingJsonResult(t *testing.T) {
	jsonResult := corejson.NewPtr(corepayload.Attributes{})
	attr, err := corepayload.New.Attributes.DeserializeUsingJsonResult(jsonResult)

	if err != nil || attr == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewAttributesCreator_CastOrDeserializeFrom(t *testing.T) {
	original := corepayload.New.Attributes.Empty()
	attr, err := corepayload.New.Attributes.CastOrDeserializeFrom(original)

	if err != nil || attr == nil {
		t.Fatal("expected success")
	}

	// nil
	_, err = corepayload.New.Attributes.CastOrDeserializeFrom(nil)

	if err == nil {
		t.Fatal("expected error for nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// newPayloadWrapperCreator — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_NewPayloadWrapper_All(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.All("name", "id", "task", "cat", "entity", false, nil, []byte("data"))

	if pw.Name != "name" {
		t.Fatal("expected name")
	}
}

func Test_Cov9_NewPayloadWrapper_UsingBytes(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("name", "id", "task", "cat", "entity", []byte("data"))

	if pw.Name != "name" {
		t.Fatal("expected name")
	}
}

func Test_Cov9_NewPayloadWrapper_Create(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.Create("name", "id", "task", "cat", "hello")

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_Record(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.Record("name", "id", "task", "cat", "hello")

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_Records(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.Records("name", "id", "task", "cat", []string{"a", "b"})

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_NameIdRecord(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameIdRecord("name", "id", "hello")

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_NameIdCategory(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameIdCategory("name", "id", "cat", "hello")

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_NameIdTaskRecord(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameIdTaskRecord("name", "id", "task", "hello")

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_NameTaskNameRecord(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameTaskNameRecord("id", "task", "hello")

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_ManyRecords(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.ManyRecords("name", "id", "task", "cat", []string{"a"})

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_Deserialize(t *testing.T) {
	original := corepayload.New.PayloadWrapper.Empty()
	original.Name = "test"
	bytes, _ := original.Serialize()

	pw, err := corepayload.New.PayloadWrapper.Deserialize(bytes)

	if err != nil || pw.Name != "test" {
		t.Fatal("expected test")
	}
}

func Test_Cov9_NewPayloadWrapper_CastOrDeserializeFrom(t *testing.T) {
	original := corepayload.New.PayloadWrapper.Empty()
	pw, err := corepayload.New.PayloadWrapper.CastOrDeserializeFrom(original)

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}

	_, err = corepayload.New.PayloadWrapper.CastOrDeserializeFrom(nil)

	if err == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_Cov9_NewPayloadWrapper_DeserializeToMany(t *testing.T) {
	bytes := []byte(`[{}]`)
	wrappers, err := corepayload.New.PayloadWrapper.DeserializeToMany(bytes)

	if err != nil || len(wrappers) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_NewPayloadWrapper_DeserializeToCollection(t *testing.T) {
	bytes := []byte(`{"Items":[]}`)
	col, err := corepayload.New.PayloadWrapper.DeserializeToCollection(bytes)

	if err != nil || col == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_UsingBytesCreateInstruction(t *testing.T) {
	instr := &corepayload.BytesCreateInstruction{
		Name:       "name",
		Identifier: "id",
		Payloads:   []byte("data"),
	}

	pw := corepayload.New.PayloadWrapper.UsingBytesCreateInstruction(instr)

	if pw.Name != "name" {
		t.Fatal("expected name")
	}
}

func Test_Cov9_NewPayloadWrapper_UsingCreateInstruction_BytesBranch(t *testing.T) {
	instr := &corepayload.PayloadCreateInstruction{
		Name:     "name",
		Payloads: []byte("data"),
	}

	pw, err := corepayload.New.PayloadWrapper.UsingCreateInstruction(instr)

	if err != nil || pw.Name != "name" {
		t.Fatal("expected name")
	}
}

func Test_Cov9_NewPayloadWrapper_UsingCreateInstruction_StringBranch(t *testing.T) {
	instr := &corepayload.PayloadCreateInstruction{
		Name:     "name",
		Payloads: `"hello"`,
	}

	pw, err := corepayload.New.PayloadWrapper.UsingCreateInstruction(instr)

	if err != nil || pw.Name != "name" {
		t.Fatal("expected name")
	}
}

func Test_Cov9_NewPayloadWrapper_UsingCreateInstruction_AnyBranch(t *testing.T) {
	instr := &corepayload.PayloadCreateInstruction{
		Name:     "name",
		Payloads: 42,
	}

	pw, err := corepayload.New.PayloadWrapper.UsingCreateInstruction(instr)

	if err != nil || pw.Name != "name" {
		t.Fatal("expected name")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// newPayloadsCollectionCreator — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_NewPayloadsCollection_All(t *testing.T) {
	col := corepayload.New.PayloadsCollection.Empty()

	if col == nil {
		t.Fatal("expected non-nil")
	}

	col = corepayload.New.PayloadsCollection.UsingCap(10)

	if col == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov9_NewPayloadsCollection_UsingWrappers(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	col := corepayload.New.PayloadsCollection.UsingWrappers(pw)

	if col.Length() != 1 {
		t.Fatal("expected 1")
	}

	emptyCol := corepayload.New.PayloadsCollection.UsingWrappers()

	if emptyCol.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_Cov9_NewPayloadsCollection_Deserialize(t *testing.T) {
	bytes := []byte(`{"Items":[]}`)
	col, err := corepayload.New.PayloadsCollection.Deserialize(bytes)

	if err != nil || col == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadsCollection_DeserializeToMany(t *testing.T) {
	bytes := []byte(`[{"Items":[]}]`)
	cols, err := corepayload.New.PayloadsCollection.DeserializeToMany(bytes)

	if err != nil || len(cols) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov9_NewPayloadsCollection_DeserializeUsingJsonResult(t *testing.T) {
	jsonResult := corejson.NewPtr(corepayload.PayloadsCollection{})
	col, err := corepayload.New.PayloadsCollection.DeserializeUsingJsonResult(jsonResult)

	if err != nil || col == nil {
		t.Fatal("expected success")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// newUserCreator — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_NewUser_All(t *testing.T) {
	u := corepayload.New.User.Empty()

	if u == nil {
		t.Fatal("expected non-nil")
	}

	u = corepayload.New.User.Create(false, "Alice", "admin")

	if u.Name != "Alice" {
		t.Fatal("expected Alice")
	}

	u = corepayload.New.User.NonSysCreate("Bob", "user")

	if u.Name != "Bob" {
		t.Fatal("expected Bob")
	}

	u = corepayload.New.User.NonSysCreateId("1", "Charlie", "user")

	if u.Identifier != "1" {
		t.Fatal("expected 1")
	}

	u = corepayload.New.User.System("sys", "system")

	if !u.IsSystemUser {
		t.Fatal("expected system user")
	}

	u = corepayload.New.User.SystemId("1", "sys", "system")

	if u.Identifier != "1" {
		t.Fatal("expected 1")
	}

	u = corepayload.New.User.UsingName("Dave")

	if u.Name != "Dave" {
		t.Fatal("expected Dave")
	}

	u = corepayload.New.User.All(true, "1", "Eve", "admin", "token", "hash")

	if u.Name != "Eve" {
		t.Fatal("expected Eve")
	}
}

func Test_Cov9_NewUser_Deserialize(t *testing.T) {
	u := corepayload.New.User.UsingName("Alice")
	bytes, _ := u.Serialize()
	result, err := corepayload.New.User.Deserialize(bytes)

	if err != nil || result.Name != "Alice" {
		t.Fatal("expected Alice")
	}
}

func Test_Cov9_NewUser_CastOrDeserializeFrom(t *testing.T) {
	u := corepayload.New.User.UsingName("Alice")
	result, err := corepayload.New.User.CastOrDeserializeFrom(u)

	if err != nil || result == nil {
		t.Fatal("expected success")
	}

	_, err = corepayload.New.User.CastOrDeserializeFrom(nil)

	if err == nil {
		t.Fatal("expected error for nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Generic helpers — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_DeserializePayloadTo(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`"hello"`)}
	result, err := corepayload.DeserializePayloadTo[string](pw)

	if err != nil || result != "hello" {
		t.Fatal("expected hello")
	}

	// nil wrapper
	_, err = corepayload.DeserializePayloadTo[string](nil)

	if err == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_Cov9_DeserializePayloadToSlice(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`["a","b"]`)}
	result, err := corepayload.DeserializePayloadToSlice[string](pw)

	if err != nil || len(result) != 2 {
		t.Fatal("expected 2")
	}

	// nil
	_, err = corepayload.DeserializePayloadToSlice[string](nil)

	if err == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_Cov9_DeserializeAttributesPayloadTo(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))
	result, err := corepayload.DeserializeAttributesPayloadTo[string](attr)

	if err != nil || result != "hello" {
		t.Fatal("expected hello")
	}

	// nil
	_, err = corepayload.DeserializeAttributesPayloadTo[string](nil)

	if err == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_Cov9_DeserializeAttributesPayloadToSlice(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`["a","b"]`))
	result, err := corepayload.DeserializeAttributesPayloadToSlice[string](attr)

	if err != nil || len(result) != 2 {
		t.Fatal("expected 2")
	}

	// nil
	_, err = corepayload.DeserializeAttributesPayloadToSlice[string](nil)

	if err == nil {
		t.Fatal("expected error for nil")
	}
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

	if pci.TaskTypeName != "task" || pci.CategoryName != "cat" {
		t.Fatal("expected task and cat")
	}
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

	if pw.Name != "name" {
		t.Fatal("expected name")
	}
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

	if err != nil || pw.Name != "name" {
		t.Fatal("expected name")
	}
}

func Test_Cov9_NewPayloadWrapper_CreateUsingTypeStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.CreateUsingTypeStringer(
		"name", "id", mockStringer{"task"}, mockStringer{"cat"}, "hello")

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_NameIdCategoryStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameIdCategoryStringer(
		"name", "id", mockStringer{"cat"}, "hello")

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_RecordsTypeStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.RecordsTypeStringer(
		"name", "id", mockStringer{"task"}, mockStringer{"cat"}, []string{"a"})

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_RecordTypeStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.RecordTypeStringer(
		"name", "id", mockStringer{"task"}, mockStringer{"cat"}, "hello")

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_NameIdTaskStringerRecord(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameIdTaskStringerRecord(
		"name", "id", mockStringer{"task"}, "hello")

	if err != nil || pw == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_NewPayloadWrapper_AllUsingStringer(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.AllUsingStringer(
		"name", "id", mockStringer{"task"}, mockStringer{"cat"}, "entity", false, nil, []byte("data"))

	if pw.Name != "name" {
		t.Fatal("expected name")
	}
}

func Test_Cov9_NewPayloadWrapper_AllUsingExpander(t *testing.T) {
	expander := corepayload.PayloadTypeExpander{
		CategoryStringer: mockStringer{"cat"},
		TaskTypeStringer: mockStringer{"task"},
	}

	pw := corepayload.New.PayloadWrapper.AllUsingExpander(
		"name", "id", expander, "entity", false, nil, []byte("data"))

	if pw.Name != "name" {
		t.Fatal("expected name")
	}
}

func Test_Cov9_NewUser_UsingNameTypeStringer(t *testing.T) {
	u := corepayload.New.User.UsingNameTypeStringer("Alice", mockStringer{"admin"})

	if u.Name != "Alice" || u.Type != "admin" {
		t.Fatal("expected Alice/admin")
	}
}

func Test_Cov9_NewUser_SysUsingNameTypeStringer(t *testing.T) {
	u := corepayload.New.User.SysUsingNameTypeStringer("sys", mockStringer{"system"})

	if u.Name != "sys" || !u.IsSystemUser {
		t.Fatal("expected system user")
	}
}

func Test_Cov9_NewUser_AllTypeStringer(t *testing.T) {
	u := corepayload.New.User.AllTypeStringer(true, "1", "Alice", mockStringer{"admin"}, "token", "hash")

	if u.Name != "Alice" || u.Type != "admin" {
		t.Fatal("expected Alice/admin")
	}
}

func Test_Cov9_NewUser_AllUsingStringer(t *testing.T) {
	u := corepayload.New.User.AllUsingStringer(false, "1", "Bob", mockStringer{"user"}, "token", "hash")

	if u.Name != "Bob" || u.Type != "user" {
		t.Fatal("expected Bob/user")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Attributes — DeserializeDynamicPayloadsTo* methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Attributes_DeserializeDynamicPayloadsToAttributes(t *testing.T) {
	inner := corepayload.New.Attributes.Empty()
	bytes := []byte(inner.JsonString())
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(bytes)

	result, err := attr.DeserializeDynamicPayloadsToAttributes()

	if err != nil || result == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_Attributes_DeserializeDynamicPayloadsToPayloadWrapper(t *testing.T) {
	inner := corepayload.New.PayloadWrapper.Empty()
	inner.Name = "inner"
	bytes, _ := inner.Serialize()
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(bytes)

	result, err := attr.DeserializeDynamicPayloadsToPayloadWrapper()

	if err != nil || result.Name != "inner" {
		t.Fatal("expected inner")
	}
}

func Test_Cov9_Attributes_DeserializeDynamicPayloadsToPayloadWrappersCollection(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"Items":[]}`))
	result, err := attr.DeserializeDynamicPayloadsToPayloadWrappersCollection()

	if err != nil || result == nil {
		t.Fatal("expected success")
	}
}

func Test_Cov9_PayloadWrapper_PayloadDeserializeToPayloadBinder(t *testing.T) {
	inner := corepayload.New.PayloadWrapper.Empty()
	inner.Name = "inner"
	bytes, _ := inner.Serialize()

	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Payloads = bytes

	binder, err := pw.PayloadDeserializeToPayloadBinder()

	if err != nil || binder == nil {
		t.Fatal("expected success")
	}

	// nil receiver
	var nilPW *corepayload.PayloadWrapper
	_, err = nilPW.PayloadDeserializeToPayloadBinder()

	if err == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_Cov9_PayloadWrapper_IsEntityTypeNamer(t *testing.T) {
	pw := &corepayload.PayloadWrapper{EntityType: "test"}

	if !pw.IsEntityTypeNamer(mockStringer{"test"}) {
		t.Fatal("expected true")
	}

	if pw.IsEntityTypeNamer(nil) {
		t.Fatal("expected false for nil namer")
	}

	var nilPW *corepayload.PayloadWrapper

	if nilPW.IsEntityTypeNamer(mockStringer{"test"}) {
		t.Fatal("nil should return false")
	}
}

func Test_Cov9_PayloadWrapper_IsCategoryNamer(t *testing.T) {
	pw := &corepayload.PayloadWrapper{EntityType: "test"}

	if !pw.IsCategoryNamer(mockStringer{"test"}) {
		t.Fatal("expected true")
	}

	var nilPW *corepayload.PayloadWrapper

	if nilPW.IsCategoryNamer(mockStringer{"test"}) {
		t.Fatal("nil should return false")
	}
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

	if pw.Name != "name" {
		t.Fatal("expected name")
	}
}

func Test_Cov9_Attributes_BasicErrorDeserializedTo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	var target any
	err := attr.BasicErrorDeserializedTo(&target)

	if err != nil {
		t.Fatal("expected nil error (empty error)")
	}
}
