package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
)

// ══════════════════════════════════════════════════════════════════════════════
// corepayload Coverage — Segment 1: PayloadWrapper + Attributes comprehensive
// ══════════════════════════════════════════════════════════════════════════════

func newTestPW() *corepayload.PayloadWrapper {
	pw, _ := corepayload.New.PayloadWrapper.Create(
		"testName", "123", "taskType", "category",
		map[string]int{"a": 1},
	)
	return pw
}

// --- PayloadWrapper basic getters ---

func Test_CovPL_S1_01_PayloadName_Category_TaskType_EntityType(t *testing.T) {
	pw := newTestPW()
	if pw.PayloadName() != "testName" {
		t.Fatal("expected testName")
	}
	if pw.PayloadCategory() != "category" {
		t.Fatal("expected category")
	}
	if pw.PayloadTaskType() != "taskType" {
		t.Fatal("expected taskType")
	}
	if pw.PayloadEntityType() == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovPL_S1_02_Identifier_IdentifierInteger_IdentifierUnsigned(t *testing.T) {
	pw := newTestPW()
	if pw.IdString() != "123" {
		t.Fatal("expected 123")
	}
	if pw.IdInteger() != 123 {
		t.Fatal("expected 123")
	}
	if pw.IdentifierUnsignedInteger() != 123 {
		t.Fatal("expected 123")
	}
	// empty identifier
	pw2 := corepayload.Empty.PayloadWrapper()
	if pw2.IdentifierInteger() >= 0 {
		t.Fatal("expected invalid")
	}
	if pw2.IdentifierUnsignedInteger() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S1_03_Length_Count_IsEmpty_HasItems_HasAnyItem(t *testing.T) {
	pw := newTestPW()
	if pw.Length() == 0 {
		t.Fatal("expected non-zero")
	}
	if pw.Count() == 0 {
		t.Fatal("expected non-zero")
	}
	if pw.IsEmpty() {
		t.Fatal("expected false")
	}
	if !pw.HasItems() {
		t.Fatal("expected true")
	}
	if !pw.HasAnyItem() {
		t.Fatal("expected true")
	}
	// nil
	var nilPW *corepayload.PayloadWrapper
	if nilPW.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S1_04_HasSafeItems_HasIssuesOrEmpty(t *testing.T) {
	pw := newTestPW()
	if !pw.HasSafeItems() {
		t.Fatal("expected true")
	}
	if pw.HasIssuesOrEmpty() {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S1_05_HasError_IsEmptyError_HasAttributes_IsEmptyAttributes(t *testing.T) {
	pw := newTestPW()
	if pw.HasError() {
		t.Fatal("expected false")
	}
	if !pw.IsEmptyError() {
		t.Fatal("expected true")
	}
	// Create() with non-bytes record does NOT set Attributes (passes nil)
	if pw.HasAttributes() {
		t.Fatal("expected false — Create with any record does not set Attributes")
	}
	if !pw.IsEmptyAttributes() {
		t.Fatal("expected true")
	}
	// nil
	var nilPW *corepayload.PayloadWrapper
	if !nilPW.IsEmptyError() {
		t.Fatal("expected true")
	}
	if nilPW.HasAttributes() {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S1_06_HasSingleRecord_IsNull_HasAnyNil(t *testing.T) {
	pw := newTestPW()
	if !pw.HasSingleRecord() {
		t.Fatal("expected true")
	}
	if pw.IsNull() {
		t.Fatal("expected false")
	}
	if pw.HasAnyNil() {
		t.Fatal("expected false")
	}
	var nilPW *corepayload.PayloadWrapper
	if !nilPW.IsNull() {
		t.Fatal("expected true")
	}
}

func Test_CovPL_S1_07_All_AllSafe(t *testing.T) {
	pw := newTestPW()
	id, name, entity, cat, payload := pw.All()
	if id == "" || name == "" || entity == "" || cat == "" || len(payload) == 0 {
		t.Fatal("expected non-empty")
	}
	var nilPW *corepayload.PayloadWrapper
	id2, name2, _, _, _ := nilPW.AllSafe()
	if id2 != "" || name2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovPL_S1_08_PayloadDynamic_DynamicPayloads_PayloadsString(t *testing.T) {
	pw := newTestPW()
	if len(pw.PayloadDynamic()) == 0 {
		t.Fatal("expected non-empty")
	}
	if len(pw.DynamicPayloads()) == 0 {
		t.Fatal("expected non-empty")
	}
	if pw.PayloadsString() == "" {
		t.Fatal("expected non-empty")
	}
	// nil
	var nilPW *corepayload.PayloadWrapper
	if len(nilPW.DynamicPayloads()) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_CovPL_S1_09_SetDynamicPayloads(t *testing.T) {
	pw := newTestPW()
	err := pw.SetDynamicPayloads([]byte(`{"b":2}`))
	if err != nil {
		t.Fatal("expected no error")
	}
	// nil
	var nilPW *corepayload.PayloadWrapper
	err2 := nilPW.SetDynamicPayloads(nil)
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S1_10_Value_Error(t *testing.T) {
	pw := newTestPW()
	_ = pw.Value()
	err := pw.Error()
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S1_11_IsEqual(t *testing.T) {
	pw1 := newTestPW()
	pw2 := newTestPW()
	if !pw1.IsEqual(pw2) {
		t.Fatal("expected equal")
	}
	// nil both
	var nilPW *corepayload.PayloadWrapper
	if !nilPW.IsEqual(nil) {
		t.Fatal("expected nil==nil")
	}
	if nilPW.IsEqual(pw1) {
		t.Fatal("expected false")
	}
	// same ptr
	if !pw1.IsEqual(pw1) {
		t.Fatal("expected true")
	}
}

func Test_CovPL_S1_12_IsPayloadsEqual_IsName_IsIdentifier_IsTaskTypeName_IsEntityType_IsCategory(t *testing.T) {
	pw := newTestPW()
	if !pw.IsPayloadsEqual(pw.PayloadDynamic()) {
		t.Fatal("expected true")
	}
	if !pw.IsName("testName") {
		t.Fatal("expected true")
	}
	if !pw.IsIdentifier("123") {
		t.Fatal("expected true")
	}
	if !pw.IsTaskTypeName("taskType") {
		t.Fatal("expected true")
	}
	if !pw.IsCategory("category") {
		t.Fatal("expected true")
	}
}

func Test_CovPL_S1_13_String_JsonString_PrettyJsonString(t *testing.T) {
	pw := newTestPW()
	if pw.String() == "" {
		t.Fatal("expected non-empty")
	}
	if pw.JsonString() == "" {
		t.Fatal("expected non-empty")
	}
	if pw.JsonStringMust() == "" {
		t.Fatal("expected non-empty")
	}
	if pw.PrettyJsonString() == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovPL_S1_14_Serialize_Json_JsonPtr(t *testing.T) {
	pw := newTestPW()
	b, err := pw.Serialize()
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
	_ = pw.SerializeMust()
	j := pw.Json()
	if j.HasError() {
		t.Fatal("expected no error")
	}
	jp := pw.JsonPtr()
	if jp == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S1_15_Deserialize_PayloadDeserialize(t *testing.T) {
	pw := newTestPW()
	var m map[string]int
	err := pw.Deserialize(&m)
	if err != nil || m["a"] != 1 {
		t.Fatal("expected a=1")
	}
	var m2 map[string]int
	err2 := pw.PayloadDeserialize(&m2)
	if err2 != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovPL_S1_16_Clone_ClonePtr_NonPtr_ToPtr(t *testing.T) {
	pw := newTestPW()
	c, err := pw.Clone(false)
	if err != nil {
		t.Fatal("expected no error")
	}
	_ = c
	c2, err2 := pw.Clone(true)
	if err2 != nil {
		t.Fatal("expected no error")
	}
	_ = c2
	cp, err3 := pw.ClonePtr(true)
	if err3 != nil || cp == nil {
		t.Fatal("expected non-nil")
	}
	// nil
	var nilPW *corepayload.PayloadWrapper
	if _, err4 := nilPW.ClonePtr(true); err4 != nil {
		t.Fatal("expected nil, nil")
	}
	_ = pw.NonPtr()
	_ = pw.ToPtr()
}

func Test_CovPL_S1_17_ParseInjectUsingJson(t *testing.T) {
	pw := newTestPW()
	jr := pw.JsonPtr()
	pw2 := corepayload.Empty.PayloadWrapper()
	_, err := pw2.ParseInjectUsingJson(jr)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovPL_S1_18_Clear_Dispose(t *testing.T) {
	pw := newTestPW()
	pw.Clear()
	pw2 := newTestPW()
	pw2.Dispose()
	// nil
	var nilPW *corepayload.PayloadWrapper
	nilPW.Clear()
	nilPW.Dispose()
}

func Test_CovPL_S1_19_Interfaces(t *testing.T) {
	pw := newTestPW()
	_ = pw.AsJsonContractsBinder()
	_ = pw.AsStandardTaskEntityDefinerContractsBinder()
	_ = pw.AsPayloadsBinder()
	_ = pw.AsJsonMarshaller()
	_ = pw.JsonModel()
	_ = pw.JsonModelAny()
	_ = pw.PayloadProperties()
}

func Test_CovPL_S1_20_BytesConverter(t *testing.T) {
	pw := newTestPW()
	bc := pw.BytesConverter()
	if bc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S1_21_InitializeAttributesOnNull(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	binder := pw.InitializeAttributesOnNull()
	if binder == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S1_22_AttrAsBinder(t *testing.T) {
	pw := newTestPW()
	binder := pw.AttrAsBinder()
	_ = binder
}

func Test_CovPL_S1_23_Username(t *testing.T) {
	pw := newTestPW()
	u := pw.Username()
	if u != "" {
		t.Fatal("expected empty for no user")
	}
}

func Test_CovPL_S1_24_MarshalJSON_UnmarshalJSON(t *testing.T) {
	pw := newTestPW()
	b, err := pw.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
	pw2 := corepayload.Empty.PayloadWrapper()
	err2 := pw2.UnmarshalJSON(b)
	if err2 != nil {
		t.Fatal("expected no error")
	}
	// nil marshal
	var nilPW *corepayload.PayloadWrapper
	_, err3 := nilPW.MarshalJSON()
	if err3 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S1_25_ReCreateUsingJsonBytes_ReCreateUsingJsonResult(t *testing.T) {
	pw := newTestPW()
	b, _ := pw.Serialize()
	pw2, err := pw.ReCreateUsingJsonBytes(b)
	if err != nil || pw2 == nil {
		t.Fatal("expected non-nil")
	}
	jr := pw.JsonPtr()
	pw3, err2 := pw.ReCreateUsingJsonResult(jr)
	if err2 != nil || pw3 == nil {
		t.Fatal("expected non-nil")
	}
}

// --- Attributes ---

func Test_CovPL_S1_30_Attributes_IsNull_HasSafeItems_IsEmpty_HasItems(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	if attr.IsNull() {
		t.Fatal("expected false")
	}
	// empty attributes
	if attr.HasSafeItems() {
		t.Fatal("expected false for empty")
	}
	if !attr.IsEmpty() {
		t.Fatal("expected true")
	}
	if attr.HasItems() {
		t.Fatal("expected false")
	}
	if attr.HasAnyItem() {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S1_31_Attributes_HasError_IsEmptyError_Error(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	if attr.HasError() {
		t.Fatal("expected false")
	}
	if !attr.IsEmptyError() {
		t.Fatal("expected true")
	}
	if attr.Error() != nil {
		t.Fatal("expected nil")
	}
	// nil
	var nilAttr *corepayload.Attributes
	if nilAttr.HasError() {
		t.Fatal("expected false")
	}
	if !nilAttr.IsEmptyError() {
		t.Fatal("expected true")
	}
}

func Test_CovPL_S1_32_Attributes_Length_Count_Capacity(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	if attr.Length() != 0 {
		t.Fatal("expected 0")
	}
	if attr.Count() != 0 {
		t.Fatal("expected 0")
	}
	if attr.Capacity() != 0 {
		t.Fatal("expected 0")
	}
	// nil
	var nilAttr *corepayload.Attributes
	if nilAttr.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S1_33_Attributes_DynamicBytesLength_StringKeyValuePairsLength_AnyKeyValuePairsLength(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	if attr.DynamicBytesLength() != 0 {
		t.Fatal("expected 0")
	}
	if attr.StringKeyValuePairsLength() != 0 {
		t.Fatal("expected 0")
	}
	if attr.AnyKeyValuePairsLength() != 0 {
		t.Fatal("expected 0")
	}
	// nil
	var nilAttr *corepayload.Attributes
	if nilAttr.DynamicBytesLength() != 0 {
		t.Fatal("expected 0")
	}
	if nilAttr.StringKeyValuePairsLength() != 0 {
		t.Fatal("expected 0")
	}
	if nilAttr.AnyKeyValuePairsLength() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S1_34_Attributes_HasPagingInfo_HasKeyValuePairs_HasFromTo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	if attr.HasPagingInfo() {
		t.Fatal("expected false")
	}
	if attr.HasKeyValuePairs() {
		t.Fatal("expected false")
	}
	if attr.HasFromTo() {
		t.Fatal("expected false")
	}
	// nil
	var nilAttr *corepayload.Attributes
	if nilAttr.HasPagingInfo() {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S1_35_Attributes_IsValid_IsInvalid_IsSafeValid_HasIssuesOrEmpty(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	// Empty() creates non-nil Attributes with no error → IsValid = (it != nil && IsEmptyError) = true
	// BUT IsInvalid = (it == nil || HasIssuesOrEmpty)
	// HasIssuesOrEmpty = IsEmpty() || !IsValid() || (BasicErr && HasError)
	// IsEmpty = len(DynamicPayloads) == 0 → true for Empty()
	// So HasIssuesOrEmpty = true, IsInvalid = true
	if !attr.IsValid() {
		t.Fatal("expected true — non-nil, no error")
	}
	if !attr.IsInvalid() {
		t.Fatal("expected true — empty attrs are invalid (HasIssuesOrEmpty=true)")
	}
	// nil
	var nilAttr *corepayload.Attributes
	if nilAttr.IsValid() {
		t.Fatal("expected false")
	}
	if !nilAttr.IsInvalid() {
		t.Fatal("expected true")
	}
}

func Test_CovPL_S1_36_Attributes_Paging_Auth_Session_Queries(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	if !attr.IsPagingInfoEmpty() {
		t.Fatal("expected true")
	}
	if !attr.IsKeyValuePairsEmpty() {
		t.Fatal("expected true")
	}
	if !attr.IsAnyKeyValuePairsEmpty() {
		t.Fatal("expected true")
	}
	if !attr.IsUserInfoEmpty() {
		t.Fatal("expected true")
	}
	if !attr.IsAuthInfoEmpty() {
		t.Fatal("expected true")
	}
	if !attr.IsSessionInfoEmpty() {
		t.Fatal("expected true")
	}
	if attr.HasUserInfo() {
		t.Fatal("expected false")
	}
	if attr.HasAuthInfo() {
		t.Fatal("expected false")
	}
	if attr.HasSessionInfo() {
		t.Fatal("expected false")
	}
	if attr.SessionInfo() != nil {
		t.Fatal("expected nil")
	}
	if attr.AuthType() != "" {
		t.Fatal("expected empty")
	}
	if attr.ResourceName() != "" {
		t.Fatal("expected empty")
	}
	if attr.HasStringKeyValuePairs() {
		t.Fatal("expected false")
	}
	if attr.HasAnyKeyValuePairs() {
		t.Fatal("expected false")
	}
	if attr.HasDynamicPayloads() {
		t.Fatal("expected false")
	}
	if attr.VirtualUser() != nil {
		t.Fatal("expected nil")
	}
	if attr.SystemUser() != nil {
		t.Fatal("expected nil")
	}
	if attr.SessionUser() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S1_37_Attributes_GetStringKeyValue_GetAnyKeyValue(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	_, found := attr.GetStringKeyValue("k")
	if found {
		t.Fatal("expected false")
	}
	_, found2 := attr.GetAnyKeyValue("k")
	if found2 {
		t.Fatal("expected false")
	}
	// nil
	var nilAttr *corepayload.Attributes
	_, found3 := nilAttr.GetStringKeyValue("k")
	if found3 {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S1_38_Attributes_HasStringKey_HasAnyKey(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	if attr.HasStringKey("k") {
		t.Fatal("expected false")
	}
	if attr.HasAnyKey("k") {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S1_39_Attributes_Payloads_PayloadsString_Hashmap_AnyKeyValMap(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	_ = attr.Payloads()
	_ = attr.PayloadsString()
	_ = attr.Hashmap()
	_ = attr.AnyKeyValMap()
	_ = attr.CompiledError()
}

func Test_CovPL_S1_40_Attributes_IsEqual(t *testing.T) {
	a1 := corepayload.New.Attributes.Empty()
	a2 := corepayload.New.Attributes.Empty()
	if !a1.IsEqual(a2) {
		t.Fatal("expected equal")
	}
	var nilAttr *corepayload.Attributes
	if !nilAttr.IsEqual(nil) {
		t.Fatal("expected nil==nil")
	}
	if nilAttr.IsEqual(a1) {
		t.Fatal("expected false")
	}
	// same ptr
	if !a1.IsEqual(a1) {
		t.Fatal("expected true")
	}
}

func Test_CovPL_S1_41_Attributes_IsErrorEqual_IsErrorDifferent(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	if !attr.IsErrorEqual(nil) {
		t.Fatal("expected true")
	}
	if attr.IsErrorDifferent(nil) {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S1_42_Attributes_Clone_ClonePtr(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	c, err := attr.Clone(false)
	if err != nil {
		t.Fatal("expected no error")
	}
	_ = c
	cp, err2 := attr.ClonePtr(true)
	if err2 != nil {
		t.Fatal("expected no error")
	}
	_ = cp
	// nil
	var nilAttr *corepayload.Attributes
	cp2, err3 := nilAttr.ClonePtr(true)
	if err3 != nil || cp2 != nil {
		t.Fatal("expected nil, nil")
	}
}

func Test_CovPL_S1_43_Attributes_SetBasicErr_SetAuthInfo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	_ = attr.SetBasicErr(nil)
	_ = attr.SetAuthInfo(nil)
	// nil attr
	var nilAttr *corepayload.Attributes
	_ = nilAttr.SetBasicErr(nil)
	_ = nilAttr.SetAuthInfo(nil)
}

func Test_CovPL_S1_44_Attributes_Clear_Dispose(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.Clear()
	attr.Dispose()
	var nilAttr *corepayload.Attributes
	nilAttr.Clear()
	nilAttr.Dispose()
}

func Test_CovPL_S1_45_Attributes_HandleErr_HandleError(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.HandleErr()
	attr.HandleError()
}

// --- NewPayloadWrapperCreator ---

func Test_CovPL_S1_50_NewPW_Empty(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	if pw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S1_51_NewPW_Deserialize(t *testing.T) {
	pw := newTestPW()
	b, _ := pw.Serialize()
	pw2, err := corepayload.New.PayloadWrapper.Deserialize(b)
	if err != nil || pw2 == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S1_52_NewPW_CastOrDeserializeFrom(t *testing.T) {
	pw := newTestPW()
	pw2, err := corepayload.New.PayloadWrapper.CastOrDeserializeFrom(pw)
	if err != nil || pw2 == nil {
		t.Fatal("expected non-nil")
	}
	// nil
	_, err2 := corepayload.New.PayloadWrapper.CastOrDeserializeFrom(nil)
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S1_53_NewPW_DeserializeToMany(t *testing.T) {
	pws := []*corepayload.PayloadWrapper{newTestPW(), newTestPW()}
	b, _ := corejson.Serialize.Raw(pws)
	many, err := corepayload.New.PayloadWrapper.DeserializeToMany(b)
	if err != nil || len(many) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CovPL_S1_54_NewPW_DeserializeToCollection(t *testing.T) {
	pws := []*corepayload.PayloadWrapper{newTestPW()}
	// DeserializeToCollection calls PayloadsCollection.Deserialize which expects
	// {"Items":[...]} format, not raw array — serialize the collection struct
	pc := &corepayload.PayloadsCollection{Items: pws}
	b, _ := corejson.Serialize.Raw(pc)
	col, err := corepayload.New.PayloadWrapper.DeserializeToCollection(b)
	if err != nil || col == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S1_55_NewPW_DeserializeUsingJsonResult(t *testing.T) {
	pw := newTestPW()
	jr := pw.JsonPtr()
	pw2, err := corepayload.New.PayloadWrapper.DeserializeUsingJsonResult(jr)
	if err != nil || pw2 == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S1_56_NewPW_UsingBytes(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes(
		"n", "1", "t", "c", "e", []byte(`{"a":1}`))
	if pw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S1_57_NewPW_Record_Records_NameIdRecord(t *testing.T) {
	_, err := corepayload.New.PayloadWrapper.Record("n", "1", "t", "c", map[string]int{"a": 1})
	if err != nil {
		t.Fatal("expected no error")
	}
	_, err2 := corepayload.New.PayloadWrapper.Records("n", "1", "t", "c", []map[string]int{{"a": 1}})
	if err2 != nil {
		t.Fatal("expected no error")
	}
	_, err3 := corepayload.New.PayloadWrapper.NameIdRecord("n", "1", map[string]int{"a": 1})
	if err3 != nil {
		t.Fatal("expected no error")
	}
	_, err4 := corepayload.New.PayloadWrapper.NameIdTaskRecord("n", "1", "t", map[string]int{"a": 1})
	if err4 != nil {
		t.Fatal("expected no error")
	}
	_, err5 := corepayload.New.PayloadWrapper.NameIdCategory("n", "1", "c", map[string]int{"a": 1})
	if err5 != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovPL_S1_58_NewPW_All_ManyRecords(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	pw := corepayload.New.PayloadWrapper.All("n", "1", "t", "c", "e", false, attr, []byte(`{}`))
	if pw == nil {
		t.Fatal("expected non-nil")
	}
	_, err := corepayload.New.PayloadWrapper.ManyRecords("n", "1", "t", "c", []int{1, 2})
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovPL_S1_59_NewPW_NameTaskNameRecord(t *testing.T) {
	_, err := corepayload.New.PayloadWrapper.NameTaskNameRecord("1", "t", map[string]int{"a": 1})
	if err != nil {
		t.Fatal("expected no error")
	}
}

// --- EmptyCreator ---

func Test_CovPL_S1_60_EmptyCreator(t *testing.T) {
	_ = corepayload.Empty.Attributes()
	_ = corepayload.Empty.AttributesDefaults()
	_ = corepayload.Empty.PayloadWrapper()
	_ = corepayload.Empty.PayloadsCollection()
}
