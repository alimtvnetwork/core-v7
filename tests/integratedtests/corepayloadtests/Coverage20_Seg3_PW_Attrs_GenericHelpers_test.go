package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
)

// ══════════════════════════════════════════════════════════════════════════════
// corepayload Coverage — Segment 3: PayloadWrapper deep methods,
//                         Attributes (getters, setters, json, clone),
//                         generic_helpers, typed_collection_funcs,
//                         typed_collection_paging, payloadProperties
// ══════════════════════════════════════════════════════════════════════════════

func newPWForSeg3() *corepayload.PayloadWrapper {
	pw, _ := corepayload.New.PayloadWrapper.Create(
		"seg3", "5", "taskType", "category",
		map[string]int{"a": 1},
	)
	return pw
}

// --- PayloadWrapper deep methods ---

func Test_CovPL_S3_01_MarshalJSON_UnmarshalJSON(t *testing.T) {
	pw := newPWForSeg3()
	b, err := pw.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
	pw2 := &corepayload.PayloadWrapper{}
	err2 := pw2.UnmarshalJSON(b)
	if err2 != nil {
		t.Fatal("expected no error")
	}
	// nil
	var nilPW *corepayload.PayloadWrapper
	_, err3 := nilPW.MarshalJSON()
	if err3 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S3_02_ReCreateUsingJsonBytes_ReCreateUsingJsonResult(t *testing.T) {
	pw := newPWForSeg3()
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

func Test_CovPL_S3_03_HasSafeItems_DynamicPayloads_SetDynamicPayloads(t *testing.T) {
	pw := newPWForSeg3()
	if !pw.HasSafeItems() {
		t.Fatal("expected true")
	}
	dp := pw.DynamicPayloads()
	if len(dp) == 0 {
		t.Fatal("expected bytes")
	}
	var nilPW *corepayload.PayloadWrapper
	if len(nilPW.DynamicPayloads()) != 0 {
		t.Fatal("expected empty")
	}
	err := pw.SetDynamicPayloads([]byte("test"))
	if err != nil {
		t.Fatal("expected no error")
	}
	var nilPW2 *corepayload.PayloadWrapper
	err2 := nilPW2.SetDynamicPayloads([]byte("test"))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S3_04_AttrAsBinder_InitializeAttributesOnNull(t *testing.T) {
	pw := newPWForSeg3()
	_ = pw.AttrAsBinder()
	_ = pw.InitializeAttributesOnNull()
}

func Test_CovPL_S3_05_BasicError_PayloadDeserializeToPayloadBinder(t *testing.T) {
	pw := newPWForSeg3()
	if pw.BasicError() != nil {
		t.Fatal("expected nil")
	}
	// PayloadDeserializeToPayloadBinder requires payloads to be PW json
	pw2, _ := corepayload.New.PayloadWrapper.Create("inner", "1", "t", "c", map[string]int{"b": 2})
	b2, _ := pw2.Serialize()
	pw3 := corepayload.New.PayloadWrapper.All("outer", "2", "t", "c", "", false, nil, b2)
	_, _ = pw3.PayloadDeserializeToPayloadBinder()
}

func Test_CovPL_S3_06_All_AllSafe(t *testing.T) {
	pw := newPWForSeg3()
	id, name, entity, category, dp := pw.All()
	if id != "5" || name != "seg3" || entity == "" || category != "category" || len(dp) == 0 {
		t.Fatal("unexpected values")
	}
	var nilPW *corepayload.PayloadWrapper
	id2, _, _, _, _ := nilPW.AllSafe()
	if id2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovPL_S3_07_PayloadName_PayloadCategory_PayloadTaskType_PayloadEntityType_PayloadDynamic(t *testing.T) {
	pw := newPWForSeg3()
	_ = pw.PayloadName()
	_ = pw.PayloadCategory()
	_ = pw.PayloadTaskType()
	_ = pw.PayloadEntityType()
	_ = pw.PayloadDynamic()
}

func Test_CovPL_S3_08_SetPayloadDynamic_SetPayloadDynamicAny(t *testing.T) {
	pw := newPWForSeg3()
	pw.SetPayloadDynamic([]byte("test"))
	_, _ = pw.SetPayloadDynamicAny(map[string]int{"x": 1})
}

func Test_CovPL_S3_09_SetAuthInfo_SetUserInfo_SetUser_SetSysUser(t *testing.T) {
	pw := newPWForSeg3()
	pw.SetAuthInfo(&corepayload.AuthInfo{})
	pw.SetUserInfo(&corepayload.UserInfo{})
	u := &corepayload.User{Name: "u"}
	pw.SetUser(u)
	pw.SetSysUser(u)
}

func Test_CovPL_S3_10_PayloadProperties(t *testing.T) {
	pw := newPWForSeg3()
	pp := pw.PayloadProperties()
	if pp == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S3_11_HandleError_ReflectSetTo(t *testing.T) {
	pw := newPWForSeg3()
	pw.HandleError()
	pw2 := &corepayload.PayloadWrapper{}
	_ = pw.ReflectSetTo(pw2)
}

func Test_CovPL_S3_12_AnyAttributes_ReflectSetAttributes(t *testing.T) {
	pw := newPWForSeg3()
	_ = pw.AnyAttributes()
	a := &corepayload.Attributes{}
	_ = pw.ReflectSetAttributes(a)
}

func Test_CovPL_S3_13_IdString_IdInteger_IdentifierInteger_IdentifierUnsignedInteger(t *testing.T) {
	pw := newPWForSeg3()
	if pw.IdString() != "5" {
		t.Fatal("expected 5")
	}
	if pw.IdInteger() != 5 {
		t.Fatal("expected 5")
	}
	if pw.IdentifierInteger() != 5 {
		t.Fatal("expected 5")
	}
	if pw.IdentifierUnsignedInteger() != 5 {
		t.Fatal("expected 5")
	}
	// empty id
	pw2 := corepayload.New.PayloadWrapper.All("n", "", "t", "c", "", false, nil, []byte("x"))
	if pw2.IdentifierInteger() >= 0 {
		t.Fatal("expected invalid")
	}
}

func Test_CovPL_S3_14_IsEqual_IsPayloadsEqual_IsName_IsIdentifier_IsTaskType_IsEntity_IsCategory(t *testing.T) {
	pw := newPWForSeg3()
	pw2 := newPWForSeg3()
	if !pw.IsEqual(pw2) {
		t.Fatal("expected true")
	}
	if !pw.IsPayloadsEqual(pw2.Payloads) {
		t.Fatal("expected true")
	}
	if !pw.IsName("seg3") {
		t.Fatal("expected true")
	}
	if !pw.IsIdentifier("5") {
		t.Fatal("expected true")
	}
	if !pw.IsTaskTypeName("taskType") {
		t.Fatal("expected true")
	}
	if !pw.IsEntityType(pw.EntityType) {
		t.Fatal("expected true")
	}
	if !pw.IsCategory("category") {
		t.Fatal("expected true")
	}
	var nilPW *corepayload.PayloadWrapper
	if !nilPW.IsEqual(nil) {
		t.Fatal("expected true")
	}
	if nilPW.IsEqual(pw) {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S3_15_HasIssuesOrEmpty_HasError_IsEmptyError_HasAttributes_IsEmptyAttributes(t *testing.T) {
	pw := newPWForSeg3()
	_ = pw.HasIssuesOrEmpty()
	_ = pw.HasError()
	_ = pw.IsEmptyError()
	_ = pw.HasAttributes()
	_ = pw.IsEmptyAttributes()
}

func Test_CovPL_S3_16_HasSingleRecord_IsNull_HasAnyNil_Count_Length_IsEmpty_HasItems_HasAnyItem(t *testing.T) {
	pw := newPWForSeg3()
	_ = pw.HasSingleRecord()
	_ = pw.IsNull()
	_ = pw.HasAnyNil()
	_ = pw.Count()
	_ = pw.Length()
	_ = pw.IsEmpty()
	_ = pw.HasItems()
	_ = pw.HasAnyItem()
}

func Test_CovPL_S3_17_BytesConverter(t *testing.T) {
	pw := newPWForSeg3()
	_ = pw.BytesConverter()
}

func Test_CovPL_S3_18_Deserialize_DeserializeMust_PayloadDeserialize_PayloadDeserializeMust(t *testing.T) {
	pw := newPWForSeg3()
	var m map[string]int
	_ = pw.Deserialize(&m)
	pw2 := newPWForSeg3()
	var m2 map[string]int
	pw2.DeserializeMust(&m2)
	var m3 map[string]int
	_ = pw.PayloadDeserialize(&m3)
	pw3 := newPWForSeg3()
	var m4 map[string]int
	pw3.PayloadDeserializeMust(&m4)
}

func Test_CovPL_S3_19_DeserializePayloads_Nested(t *testing.T) {
	pw := newPWForSeg3()
	// DeserializePayloadsToPayloadsCollection needs payload to be a collection
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWForSeg3())
	b, _ := corejson.Serialize.Raw(pc)
	pw2 := corepayload.New.PayloadWrapper.All("outer", "1", "t", "c", "", false, nil, b)
	_, _ = pw2.DeserializePayloadsToPayloadsCollection()
	_, _ = pw.DeserializePayloadsToPayloadWrapper()
}

func Test_CovPL_S3_20_PW_JsonModel_JsonModelAny_Json_JsonPtr_AsJsonContractsBinder(t *testing.T) {
	pw := newPWForSeg3()
	_ = pw.JsonModel()
	_ = pw.JsonModelAny()
	_ = pw.Json()
	_ = pw.JsonPtr()
	_ = pw.AsJsonContractsBinder()
}

func Test_CovPL_S3_21_PW_ParseInjectUsingJson_ParseInjectUsingJsonMust(t *testing.T) {
	pw := newPWForSeg3()
	jr := pw.JsonPtr()
	pw2 := &corepayload.PayloadWrapper{}
	_, _ = pw2.ParseInjectUsingJson(jr)
	pw3 := &corepayload.PayloadWrapper{}
	_ = pw3.ParseInjectUsingJsonMust(jr)
}

func Test_CovPL_S3_22_PW_JsonParseSelfInject(t *testing.T) {
	pw := newPWForSeg3()
	jr := pw.JsonPtr()
	pw2 := &corepayload.PayloadWrapper{}
	_ = pw2.JsonParseSelfInject(jr)
}

func Test_CovPL_S3_23_PW_String_PrettyJsonString_JsonString_JsonStringMust(t *testing.T) {
	pw := newPWForSeg3()
	if pw.String() == "" {
		t.Fatal("expected non-empty")
	}
	if pw.PrettyJsonString() == "" {
		t.Fatal("expected non-empty")
	}
	if pw.JsonString() == "" {
		t.Fatal("expected non-empty")
	}
	if pw.JsonStringMust() == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovPL_S3_24_PW_PayloadsString_PayloadsPrettyString_PayloadsJsonResult(t *testing.T) {
	pw := newPWForSeg3()
	if pw.PayloadsString() == "" {
		t.Fatal("expected non-empty")
	}
	if pw.PayloadsPrettyString() == "" {
		t.Fatal("expected non-empty")
	}
	jr := pw.PayloadsJsonResult()
	if jr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S3_25_PW_Clear_Dispose(t *testing.T) {
	pw := newPWForSeg3()
	pw.Clear()
	pw2 := newPWForSeg3()
	pw2.Dispose()
	var nilPW *corepayload.PayloadWrapper
	nilPW.Clear()
	nilPW.Dispose()
}

func Test_CovPL_S3_26_PW_Clone_ClonePtr_NonPtr_ToPtr(t *testing.T) {
	pw := newPWForSeg3()
	_, _ = pw.Clone(true)
	_, _ = pw.Clone(false)
	_, _ = pw.ClonePtr(true)
	_, _ = pw.ClonePtr(false)
	_ = pw.NonPtr()
	_ = pw.ToPtr()
	var nilPW *corepayload.PayloadWrapper
	if nilPW.NonPtr().Name != "" {
		t.Fatal("expected empty")
	}
	c, _ := nilPW.ClonePtr(true)
	if c != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S3_27_PW_AsInterfaces(t *testing.T) {
	pw := newPWForSeg3()
	_ = pw.AsStandardTaskEntityDefinerContractsBinder()
	_ = pw.AsPayloadsBinder()
	_ = pw.AsJsonMarshaller()
}

func Test_CovPL_S3_28_PW_Serialize_SerializeMust_Username_Value_Error(t *testing.T) {
	pw := newPWForSeg3()
	_, _ = pw.Serialize()
	_ = pw.SerializeMust()
	_ = pw.Username()
	_ = pw.Value()
	_ = pw.Error()
}

func Test_CovPL_S3_29_PW_ValueReflectSet(t *testing.T) {
	pw := newPWForSeg3()
	var b []byte
	_ = pw.ValueReflectSet(&b)
}

func Test_CovPL_S3_30_PW_IsStandardTaskEntityEqual(t *testing.T) {
	pw := newPWForSeg3()
	pw2 := newPWForSeg3()
	if !pw.IsStandardTaskEntityEqual(pw2) {
		t.Fatal("expected true")
	}
}

// --- Attributes Getters ---

func Test_CovPL_S3_40_Attr_Getters(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	_ = attr.IsNull()
	_ = attr.HasSafeItems()
	_ = attr.Payloads()
	_ = attr.PayloadsString()
	_ = attr.AnyKeyValMap()
	_ = attr.Hashmap()
	_ = attr.CompiledError()
	_ = attr.HasIssuesOrEmpty()
	_ = attr.IsSafeValid()
	_ = attr.HasAnyItem()
	_ = attr.Count()
	_ = attr.Capacity()
	_ = attr.Length()
	_ = attr.HasPagingInfo()
	_ = attr.HasKeyValuePairs()
	_ = attr.HasFromTo()
	_ = attr.IsValid()
	_ = attr.IsInvalid()
	_ = attr.HasError()
	_ = attr.Error()
	_ = attr.IsEmptyError()
	_ = attr.DynamicBytesLength()
	_ = attr.StringKeyValuePairsLength()
	_ = attr.AnyKeyValuePairsLength()
	_ = attr.IsEmpty()
	_ = attr.HasItems()
	_ = attr.IsPagingInfoEmpty()
	_ = attr.IsKeyValuePairsEmpty()
	_ = attr.IsAnyKeyValuePairsEmpty()
	_ = attr.IsUserInfoEmpty()
	_ = attr.VirtualUser()
	_ = attr.SystemUser()
	_ = attr.SessionUser()
	_ = attr.IsAuthInfoEmpty()
	_ = attr.IsSessionInfoEmpty()
	_ = attr.HasUserInfo()
	_ = attr.HasAuthInfo()
	_ = attr.HasSessionInfo()
	_ = attr.SessionInfo()
	_ = attr.AuthType()
	_ = attr.ResourceName()
	_ = attr.HasStringKeyValuePairs()
	_ = attr.HasAnyKeyValuePairs()
	_ = attr.HasDynamicPayloads()
}

func Test_CovPL_S3_41_Attr_GetStringKeyValue_GetAnyKeyValue(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.AddNewStringKeyValueOnly("k", "v")
	v, found := attr.GetStringKeyValue("k")
	if !found || v != "v" {
		t.Fatal("expected v")
	}
	_, found2 := attr.GetStringKeyValue("missing")
	if found2 {
		t.Fatal("expected false")
	}
	_, _ = attr.GetAnyKeyValue("missing")
	var nilAttr *corepayload.Attributes
	_, f := nilAttr.GetStringKeyValue("k")
	if f {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S3_42_Attr_HasStringKey_HasAnyKey(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.AddNewStringKeyValueOnly("k", "v")
	if !attr.HasStringKey("k") {
		t.Fatal("expected true")
	}
	if attr.HasAnyKey("k") {
		t.Fatal("expected false") // AnyKeyValuePairs is empty
	}
}

func Test_CovPL_S3_43_Attr_IsErrorEqual_IsErrorDifferent(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	if !attr.IsErrorEqual(nil) {
		t.Fatal("expected true")
	}
	if attr.IsErrorDifferent(nil) {
		t.Fatal("expected false")
	}
}

// --- Attributes Setters ---

func Test_CovPL_S3_44_Attr_HandleErr_HandleError_MustBeEmptyError(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.HandleErr()
	attr.HandleError()
	attr.MustBeEmptyError()
}

func Test_CovPL_S3_45_Attr_SetAuthInfo_SetUserInfo_NilReceiver(t *testing.T) {
	var nilAttr *corepayload.Attributes
	r := nilAttr.SetAuthInfo(&corepayload.AuthInfo{})
	if r == nil {
		t.Fatal("expected non-nil")
	}
	var nilAttr2 *corepayload.Attributes
	r2 := nilAttr2.SetUserInfo(&corepayload.UserInfo{})
	if r2 == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S3_46_Attr_AddNewStringKeyValueOnly_AddNewAnyKeyValueOnly(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	if !attr.AddNewStringKeyValueOnly("k", "v") {
		t.Fatal("expected true")
	}
	if !attr.AddNewAnyKeyValueOnly("k", "v") {
		t.Fatal("expected true")
	}
	var nilAttr *corepayload.Attributes
	if nilAttr.AddNewStringKeyValueOnly("k", "v") {
		t.Fatal("expected false")
	}
	if nilAttr.AddNewAnyKeyValueOnly("k", "v") {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S3_47_Attr_AddOrUpdateString_AddOrUpdateAnyItem(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	_ = attr.AddOrUpdateString("k", "v")
	_ = attr.AddOrUpdateAnyItem("k", "v")
	var nilAttr *corepayload.Attributes
	_ = nilAttr.AddOrUpdateString("k", "v")
	_ = nilAttr.AddOrUpdateAnyItem("k", "v")
}

func Test_CovPL_S3_48_Attr_SetBasicErr_NilReceiver(t *testing.T) {
	var nilAttr *corepayload.Attributes
	r := nilAttr.SetBasicErr(nil)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S3_49_Attr_ReflectSetTo_AnyKeyReflectSetTo(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr2 := &corepayload.Attributes{}
	_ = attr.ReflectSetTo(attr2)
}

func Test_CovPL_S3_50_Attr_Clear_Dispose(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr.Clear()
	attr2 := corepayload.New.Attributes.Empty()
	attr2.Dispose()
	var nilAttr *corepayload.Attributes
	nilAttr.Clear()
	nilAttr.Dispose()
}

// --- Attributes JSON ---

func Test_CovPL_S3_51_Attr_Json_Methods(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	if attr.PayloadsPrettyString() == "" {
		t.Fatal("expected non-empty")
	}
	jr := attr.PayloadsJsonResult()
	if jr == nil {
		t.Fatal("expected non-nil")
	}
	if attr.JsonString() == "" {
		t.Fatal("expected non-empty")
	}
	_ = attr.JsonStringMust()
	_ = attr.String()
	_ = attr.PrettyJsonString()
	_ = attr.Json()
	_ = attr.JsonPtr()
	_ = attr.JsonModel()
	_ = attr.JsonModelAny()
	_ = attr.NonPtr()
	_ = attr.AsAttributesBinder()
	_ = attr.AsJsonContractsBinder()
}

func Test_CovPL_S3_52_Attr_ParseInjectUsingJson(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	jr := attr.JsonPtr()
	attr2 := &corepayload.Attributes{}
	_, _ = attr2.ParseInjectUsingJson(jr)
	attr3 := &corepayload.Attributes{}
	_ = attr3.ParseInjectUsingJsonMust(jr)
	attr4 := &corepayload.Attributes{}
	_ = attr4.JsonParseSelfInject(jr)
}

func Test_CovPL_S3_53_Attr_DeserializeDynamicPayloads(t *testing.T) {
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	var m map[string]int
	_ = attr.DeserializeDynamicPayloads(&m)
	attr.DeserializeDynamicPayloadsMust(&m)
	_ = attr.DynamicPayloadsDeserialize(&m)
	attr.DynamicPayloadsDeserializeMust(&m)
}

func Test_CovPL_S3_54_Attr_DeserializeDynamicPayloadsToAttributes(t *testing.T) {
	inner := corepayload.New.Attributes.Empty()
	b, _ := corejson.Serialize.Raw(inner)
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(b)
	_, _ = attr.DeserializeDynamicPayloadsToAttributes()
}

func Test_CovPL_S3_55_Attr_DeserializeDynamicPayloadsToPayloadWrapper(t *testing.T) {
	pw := newPWForSeg3()
	b, _ := pw.Serialize()
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(b)
	_, _ = attr.DeserializeDynamicPayloadsToPayloadWrapper()
}

func Test_CovPL_S3_56_Attr_DeserializeDynamicPayloadsToPayloadWrappersCollection(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWForSeg3())
	b, _ := corejson.Serialize.Raw(pc)
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(b)
	_, _ = attr.DeserializeDynamicPayloadsToPayloadWrappersCollection()
}

func Test_CovPL_S3_57_Attr_DynamicPayloadsDeserialize_Nil(t *testing.T) {
	var nilAttr *corepayload.Attributes
	err := nilAttr.DynamicPayloadsDeserialize(nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S3_58_Attr_Clone_ClonePtr(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	_, _ = attr.Clone(true)
	_, _ = attr.Clone(false)
	_, _ = attr.ClonePtr(true)
	_, _ = attr.ClonePtr(false)
	var nilAttr *corepayload.Attributes
	c, _ := nilAttr.ClonePtr(true)
	if c != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S3_59_Attr_IsEqual(t *testing.T) {
	attr := corepayload.New.Attributes.Empty()
	attr2 := corepayload.New.Attributes.Empty()
	if !attr.IsEqual(attr2) {
		t.Fatal("expected true")
	}
	var nilAttr *corepayload.Attributes
	if !nilAttr.IsEqual(nil) {
		t.Fatal("expected true")
	}
	if nilAttr.IsEqual(attr) {
		t.Fatal("expected false")
	}
}

// --- Attributes Creator ---

func Test_CovPL_S3_60_AttrCreator(t *testing.T) {
	_ = corepayload.New.Attributes.Empty()
	_ = corepayload.New.Attributes.Create(nil, nil, nil)
	_ = corepayload.New.Attributes.UsingAuthInfo(nil)
	_ = corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("x"))
	_ = corepayload.New.Attributes.UsingBasicError(nil)
	_ = corepayload.New.Attributes.UsingKeyValues(nil)
	_ = corepayload.New.Attributes.UsingAnyKeyValues(nil)
	_ = corepayload.New.Attributes.UsingAuthInfoKeyValues(nil, nil)
	_ = corepayload.New.Attributes.UsingAuthInfoAnyKeyValues(nil, nil)
	_ = corepayload.New.Attributes.UsingAuthInfoDynamicBytes(nil, nil)
	_ = corepayload.New.Attributes.UsingKeyValuesPlusDynamic(nil, nil)
	_ = corepayload.New.Attributes.UsingAnyKeyValuesPlusDynamic(nil, nil)
	_ = corepayload.New.Attributes.ErrFromTo(nil, nil, nil)
	_ = corepayload.New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)
	_, _ = corepayload.New.Attributes.AllAny(nil, nil, nil, nil, 1)
	_, _ = corepayload.New.Attributes.PageInfoAny(nil, 1)
	_, _ = corepayload.New.Attributes.UsingDynamicPayloadAny(nil, 1)
	_, _ = corepayload.New.Attributes.UsingAuthInfoJsonResult(nil, corejson.NewPtr(1))
	_, _ = corepayload.New.Attributes.Deserialize([]byte(`{}`))
	_, _ = corepayload.New.Attributes.DeserializeMany([]byte(`[{}]`))
	_, _ = corepayload.New.Attributes.DeserializeUsingJsonResult(corejson.NewPtr(corepayload.Attributes{}))
	_, _ = corepayload.New.Attributes.CastOrDeserializeFrom(corepayload.New.Attributes.Empty())
	_, _ = corepayload.New.Attributes.CastOrDeserializeFrom(nil)
}

// --- generic_helpers ---

func Test_CovPL_S3_70_DeserializePayloadTo(t *testing.T) {
	type D struct{ A int }
	pw, _ := corepayload.New.PayloadWrapper.Create("n", "1", "t", "c", D{A: 1})
	d, err := corepayload.DeserializePayloadTo[D](pw)
	if err != nil || d.A != 1 {
		t.Fatal("expected A=1")
	}
	_ = corepayload.DeserializePayloadToMust[D](pw)
	// nil
	_, err2 := corepayload.DeserializePayloadTo[D](nil)
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S3_71_DeserializePayloadToSlice(t *testing.T) {
	type D struct{ A int }
	pw, _ := corepayload.New.PayloadWrapper.Create("n", "1", "t", "c", []D{{A: 1}, {A: 2}})
	ds, err := corepayload.DeserializePayloadToSlice[D](pw)
	if err != nil || len(ds) != 2 {
		t.Fatal("expected 2")
	}
	_ = corepayload.DeserializePayloadToSliceMust[D](pw)
	// nil
	_, err2 := corepayload.DeserializePayloadToSlice[D](nil)
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S3_72_DeserializeAttributesPayloadTo(t *testing.T) {
	type D struct{ A int }
	b, _ := corejson.Serialize.Raw(D{A: 1})
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(b)
	d, err := corepayload.DeserializeAttributesPayloadTo[D](attr)
	if err != nil || d.A != 1 {
		t.Fatal("expected A=1")
	}
	_ = corepayload.DeserializeAttributesPayloadToMust[D](attr)
	// nil
	_, err2 := corepayload.DeserializeAttributesPayloadTo[D](nil)
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S3_73_DeserializeAttributesPayloadToSlice(t *testing.T) {
	type D struct{ A int }
	b, _ := corejson.Serialize.Raw([]D{{A: 1}, {A: 2}})
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(b)
	ds, err := corepayload.DeserializeAttributesPayloadToSlice[D](attr)
	if err != nil || len(ds) != 2 {
		t.Fatal("expected 2")
	}
	// nil
	_, err2 := corepayload.DeserializeAttributesPayloadToSlice[D](nil)
	if err2 == nil {
		t.Fatal("expected error")
	}
}

// --- typed_collection_funcs ---

func Test_CovPL_S3_80_MapTypedPayloads(t *testing.T) {
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](2)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col.Add(tw)
	names := corepayload.MapTypedPayloads[D, string](col, func(item *corepayload.TypedPayloadWrapper[D]) string {
		return item.Name()
	})
	if len(names) != 1 {
		t.Fatal("expected 1")
	}
	// empty
	empty := corepayload.EmptyTypedPayloadCollection[D]()
	r := corepayload.MapTypedPayloads[D, string](empty, func(item *corepayload.TypedPayloadWrapper[D]) string { return "" })
	if len(r) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S3_81_MapTypedPayloadData(t *testing.T) {
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 5})
	col.Add(tw)
	vals := corepayload.MapTypedPayloadData[D, int](col, func(d D) int { return d.A })
	if vals[0] != 5 {
		t.Fatal("expected 5")
	}
}

func Test_CovPL_S3_82_FlatMapTypedPayloads(t *testing.T) {
	type D struct{ Tags []string }
	col := corepayload.NewTypedPayloadCollection[D](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{Tags: []string{"a", "b"}})
	col.Add(tw)
	tags := corepayload.FlatMapTypedPayloads[D, string](col, func(item *corepayload.TypedPayloadWrapper[D]) []string {
		return item.Data().Tags
	})
	if len(tags) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CovPL_S3_83_FlatMapTypedPayloadData(t *testing.T) {
	type D struct{ Tags []string }
	col := corepayload.NewTypedPayloadCollection[D](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{Tags: []string{"a"}})
	col.Add(tw)
	tags := corepayload.FlatMapTypedPayloadData[D, string](col, func(d D) []string { return d.Tags })
	if len(tags) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovPL_S3_84_ReduceTypedPayloads(t *testing.T) {
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](2)
	tw1, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 3})
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "2", "e", D{A: 7})
	col.Add(tw1)
	col.Add(tw2)
	sum := corepayload.ReduceTypedPayloads[D, int](col, 0, func(acc int, item *corepayload.TypedPayloadWrapper[D]) int {
		return acc + item.Data().A
	})
	if sum != 10 {
		t.Fatal("expected 10")
	}
}

func Test_CovPL_S3_85_ReduceTypedPayloadData(t *testing.T) {
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 5})
	col.Add(tw)
	sum := corepayload.ReduceTypedPayloadData[D, int](col, 0, func(acc int, d D) int { return acc + d.A })
	if sum != 5 {
		t.Fatal("expected 5")
	}
}

func Test_CovPL_S3_86_GroupTypedPayloads_GroupTypedPayloadData(t *testing.T) {
	type D struct{ Cat string }
	col := corepayload.NewTypedPayloadCollection[D](2)
	tw1, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{Cat: "a"})
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "2", "e", D{Cat: "b"})
	col.Add(tw1)
	col.Add(tw2)
	groups := corepayload.GroupTypedPayloads[D, string](col, func(item *corepayload.TypedPayloadWrapper[D]) string {
		return item.Data().Cat
	})
	if len(groups) != 2 {
		t.Fatal("expected 2")
	}
	groups2 := corepayload.GroupTypedPayloadData[D, string](col, func(d D) string { return d.Cat })
	if len(groups2) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CovPL_S3_87_PartitionTypedPayloads(t *testing.T) {
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](2)
	tw1, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "2", "e", D{A: 2})
	col.Add(tw1)
	col.Add(tw2)
	m, nm := corepayload.PartitionTypedPayloads[D](col, func(item *corepayload.TypedPayloadWrapper[D]) bool {
		return item.Data().A == 1
	})
	if m.Length() != 1 || nm.Length() != 1 {
		t.Fatal("expected 1,1")
	}
}

func Test_CovPL_S3_88_AnyTypedPayload_AllTypedPayloads(t *testing.T) {
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col.Add(tw)
	if !corepayload.AnyTypedPayload[D](col, func(item *corepayload.TypedPayloadWrapper[D]) bool { return true }) {
		t.Fatal("expected true")
	}
	if !corepayload.AllTypedPayloads[D](col, func(item *corepayload.TypedPayloadWrapper[D]) bool { return true }) {
		t.Fatal("expected true")
	}
	// empty
	empty := corepayload.EmptyTypedPayloadCollection[D]()
	if corepayload.AnyTypedPayload[D](empty, func(item *corepayload.TypedPayloadWrapper[D]) bool { return true }) {
		t.Fatal("expected false")
	}
	if !corepayload.AllTypedPayloads[D](empty, func(item *corepayload.TypedPayloadWrapper[D]) bool { return false }) {
		t.Fatal("expected true (vacuous)")
	}
}

func Test_CovPL_S3_89_ConvertTypedPayloads(t *testing.T) {
	type D struct{ A int }
	type D2 struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col.Add(tw)
	converted, err := corepayload.ConvertTypedPayloads[D, D2](col)
	if err != nil || converted.Length() != 1 {
		t.Fatal("expected 1")
	}
	// empty
	empty := corepayload.EmptyTypedPayloadCollection[D]()
	c2, _ := corepayload.ConvertTypedPayloads[D, D2](empty)
	if c2.Length() != 0 {
		t.Fatal("expected 0")
	}
}

// --- typed_collection_paging ---

func Test_CovPL_S3_90_TPC_Paging(t *testing.T) {
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](10)
	for i := 0; i < 10; i++ {
		tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: i})
		col.Add(tw)
	}
	if col.GetPagesSize(3) != 4 {
		t.Fatal("expected 4")
	}
	if col.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	_ = col.GetPagingInfo(3, 1)
	_ = col.GetSinglePageCollection(3, 2)
	pages := col.GetPagedCollection(3)
	if len(pages) < 3 {
		t.Fatal("expected at least 3")
	}
	withInfo := col.GetPagedCollectionWithInfo(3)
	if len(withInfo) < 3 {
		t.Fatal("expected at least 3")
	}
	// small
	small := corepayload.NewTypedPayloadCollection[D](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	small.Add(tw)
	_ = small.GetPagedCollection(10)
	_ = small.GetSinglePageCollection(10, 1)
}

// --- PayloadCreateInstruction / BytesCreateInstruction ---

func Test_CovPL_S3_95_PayloadCreateInstruction(t *testing.T) {
	inst := corepayload.PayloadCreateInstruction{
		Name: "n", Identifier: "1", TaskTypeName: "t",
		EntityType: "e", CategoryName: "c",
		HasManyRecords: false, Payloads: map[string]int{"a": 1},
	}
	if inst.Name != "n" {
		t.Fatal("expected n")
	}
}

func Test_CovPL_S3_96_BytesCreateInstruction(t *testing.T) {
	inst := corepayload.BytesCreateInstruction{
		Name: "n", Identifier: "1", TaskTypeName: "t",
		EntityType: "e", CategoryName: "c",
		HasManyRecords: false, Payloads: []byte("x"),
	}
	if inst.Name != "n" {
		t.Fatal("expected n")
	}
}

type seg3Stringer struct{ v string }

func (s seg3Stringer) String() string { return s.v }

func Test_CovPL_S3_97_PayloadTypeExpander(t *testing.T) {
	pe := corepayload.PayloadTypeExpander{
		CategoryStringer: seg3Stringer{"cat"},
		TaskTypeStringer: seg3Stringer{"task"},
	}
	if pe.CategoryStringer.String() != "cat" {
		t.Fatal("expected cat")
	}
}
