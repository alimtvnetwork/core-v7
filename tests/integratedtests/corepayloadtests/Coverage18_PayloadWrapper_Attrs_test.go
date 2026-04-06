package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
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
	// Arrange
	pw := newTestPW()

	// Act
	actual := args.Map{"result": pw.PayloadName() != "testName"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected testName", actual)
	actual := args.Map{"result": pw.PayloadCategory() != "category"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected category", actual)
	actual := args.Map{"result": pw.PayloadTaskType() != "taskType"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected taskType", actual)
	actual := args.Map{"result": pw.PayloadEntityType() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovPL_S1_02_Identifier_IdentifierInteger_IdentifierUnsigned(t *testing.T) {
	// Arrange
	pw := newTestPW()

	// Act
	actual := args.Map{"result": pw.IdString() != "123"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 123", actual)
	actual := args.Map{"result": pw.IdInteger() != 123}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 123", actual)
	actual := args.Map{"result": pw.IdentifierUnsignedInteger() != 123}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 123", actual)
	// empty identifier
	pw2 := corepayload.Empty.PayloadWrapper()
	actual := args.Map{"result": pw2.IdentifierInteger() >= 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": pw2.IdentifierUnsignedInteger() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S1_03_Length_Count_IsEmpty_HasItems_HasAnyItem(t *testing.T) {
	// Arrange
	pw := newTestPW()

	// Act
	actual := args.Map{"result": pw.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-zero", actual)
	actual := args.Map{"result": pw.Count() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-zero", actual)
	actual := args.Map{"result": pw.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": pw.HasItems()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": pw.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	// nil
	var nilPW *corepayload.PayloadWrapper
	actual := args.Map{"result": nilPW.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S1_04_HasSafeItems_HasIssuesOrEmpty(t *testing.T) {
	// Arrange
	pw := newTestPW()

	// Act
	actual := args.Map{"result": pw.HasSafeItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": pw.HasIssuesOrEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S1_05_HasError_IsEmptyError_HasAttributes_IsEmptyAttributes(t *testing.T) {
	// Arrange
	pw := newTestPW()

	// Act
	actual := args.Map{"result": pw.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": pw.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	// Create() with non-bytes record does NOT set Attributes (passes nil)
	actual := args.Map{"result": pw.HasAttributes()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false — Create with any record does not set Attributes", actual)
	actual := args.Map{"result": pw.IsEmptyAttributes()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	// nil
	var nilPW *corepayload.PayloadWrapper
	actual := args.Map{"result": nilPW.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": nilPW.HasAttributes()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S1_06_HasSingleRecord_IsNull_HasAnyNil(t *testing.T) {
	// Arrange
	pw := newTestPW()

	// Act
	actual := args.Map{"result": pw.HasSingleRecord()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": pw.IsNull()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": pw.HasAnyNil()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	var nilPW *corepayload.PayloadWrapper
	actual := args.Map{"result": nilPW.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovPL_S1_07_All_AllSafe(t *testing.T) {
	// Arrange
	pw := newTestPW()
	id, name, entity, cat, payload := pw.All()

	// Act
	actual := args.Map{"result": id == "" || name == "" || entity == "" || cat == "" || len(payload) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	var nilPW *corepayload.PayloadWrapper
	id2, name2, _, _, _ := nilPW.AllSafe()
	actual := args.Map{"result": id2 != "" || name2 != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovPL_S1_08_PayloadDynamic_DynamicPayloads_PayloadsString(t *testing.T) {
	// Arrange
	pw := newTestPW()

	// Act
	actual := args.Map{"result": len(pw.PayloadDynamic()) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual := args.Map{"result": len(pw.DynamicPayloads()) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual := args.Map{"result": pw.PayloadsString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// nil
	var nilPW *corepayload.PayloadWrapper
	actual := args.Map{"result": len(nilPW.DynamicPayloads()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovPL_S1_09_SetDynamicPayloads(t *testing.T) {
	// Arrange
	pw := newTestPW()
	err := pw.SetDynamicPayloads([]byte(`{"b":2}`))

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	// nil
	var nilPW *corepayload.PayloadWrapper
	err2 := nilPW.SetDynamicPayloads(nil)
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovPL_S1_10_Value_Error(t *testing.T) {
	// Arrange
	pw := newTestPW()
	_ = pw.Value()
	err := pw.Error()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S1_11_IsEqual(t *testing.T) {
	// Arrange
	pw1 := newTestPW()
	pw2 := newTestPW()

	// Act
	actual := args.Map{"result": pw1.IsEqual(pw2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	// nil both
	var nilPW *corepayload.PayloadWrapper
	actual := args.Map{"result": nilPW.IsEqual(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil==nil", actual)
	actual := args.Map{"result": nilPW.IsEqual(pw1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	// same ptr
	actual := args.Map{"result": pw1.IsEqual(pw1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovPL_S1_12_IsPayloadsEqual_IsName_IsIdentifier_IsTaskTypeName_IsEntityType_IsCategory(t *testing.T) {
	// Arrange
	pw := newTestPW()

	// Act
	actual := args.Map{"result": pw.IsPayloadsEqual(pw.PayloadDynamic())}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": pw.IsName("testName")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": pw.IsIdentifier("123")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": pw.IsTaskTypeName("taskType")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": pw.IsCategory("category")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovPL_S1_13_String_JsonString_PrettyJsonString(t *testing.T) {
	// Arrange
	pw := newTestPW()

	// Act
	actual := args.Map{"result": pw.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual := args.Map{"result": pw.JsonString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual := args.Map{"result": pw.JsonStringMust() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual := args.Map{"result": pw.PrettyJsonString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovPL_S1_14_Serialize_Json_JsonPtr(t *testing.T) {
	// Arrange
	pw := newTestPW()
	b, err := pw.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	_ = pw.SerializeMust()
	j := pw.Json()
	actual := args.Map{"result": j.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	jp := pw.JsonPtr()
	actual := args.Map{"result": jp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S1_15_Deserialize_PayloadDeserialize(t *testing.T) {
	// Arrange
	pw := newTestPW()
	var m map[string]int
	err := pw.Deserialize(&m)

	// Act
	actual := args.Map{"result": err != nil || m["a"] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a=1", actual)
	var m2 map[string]int
	err2 := pw.PayloadDeserialize(&m2)
	actual := args.Map{"result": err2 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovPL_S1_16_Clone_ClonePtr_NonPtr_ToPtr(t *testing.T) {
	// Arrange
	pw := newTestPW()
	c, err := pw.Clone(false)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = c
	c2, err2 := pw.Clone(true)
	actual := args.Map{"result": err2 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = c2
	cp, err3 := pw.ClonePtr(true)
	actual := args.Map{"result": err3 != nil || cp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// nil
	var nilPW *corepayload.PayloadWrapper
	actual := args.Map{"result": _, err4 := nilPW.ClonePtr(true); err4 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil, nil", actual)
	_ = pw.NonPtr()
	_ = pw.ToPtr()
}

func Test_CovPL_S1_17_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	pw := newTestPW()
	jr := pw.JsonPtr()
	pw2 := corepayload.Empty.PayloadWrapper()
	_, err := pw2.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
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
	// Arrange
	pw := newTestPW()
	bc := pw.BytesConverter()

	// Act
	actual := args.Map{"result": bc == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S1_21_InitializeAttributesOnNull(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}
	binder := pw.InitializeAttributesOnNull()

	// Act
	actual := args.Map{"result": binder == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S1_22_AttrAsBinder(t *testing.T) {
	pw := newTestPW()
	binder := pw.AttrAsBinder()
	_ = binder
}

func Test_CovPL_S1_23_Username(t *testing.T) {
	// Arrange
	pw := newTestPW()
	u := pw.Username()

	// Act
	actual := args.Map{"result": u != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for no user", actual)
}

func Test_CovPL_S1_24_MarshalJSON_UnmarshalJSON(t *testing.T) {
	// Arrange
	pw := newTestPW()
	b, err := pw.MarshalJSON()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	pw2 := corepayload.Empty.PayloadWrapper()
	err2 := pw2.UnmarshalJSON(b)
	actual := args.Map{"result": err2 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	// nil marshal
	var nilPW *corepayload.PayloadWrapper
	_, err3 := nilPW.MarshalJSON()
	actual := args.Map{"result": err3 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovPL_S1_25_ReCreateUsingJsonBytes_ReCreateUsingJsonResult(t *testing.T) {
	// Arrange
	pw := newTestPW()
	b, _ := pw.Serialize()
	pw2, err := pw.ReCreateUsingJsonBytes(b)

	// Act
	actual := args.Map{"result": err != nil || pw2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	jr := pw.JsonPtr()
	pw3, err2 := pw.ReCreateUsingJsonResult(jr)
	actual := args.Map{"result": err2 != nil || pw3 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// --- Attributes ---

func Test_CovPL_S1_30_Attributes_IsNull_HasSafeItems_IsEmpty_HasItems(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"result": attr.IsNull()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	// empty attributes
	actual := args.Map{"result": attr.HasSafeItems()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	actual := args.Map{"result": attr.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": attr.HasItems()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": attr.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S1_31_Attributes_HasError_IsEmptyError_Error(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"result": attr.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": attr.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": attr.Error() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	// nil
	var nilAttr *corepayload.Attributes
	actual := args.Map{"result": nilAttr.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": nilAttr.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovPL_S1_32_Attributes_Length_Count_Capacity(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"result": attr.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": attr.Count() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": attr.Capacity() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	// nil
	var nilAttr *corepayload.Attributes
	actual := args.Map{"result": nilAttr.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S1_33_Attributes_DynamicBytesLength_StringKeyValuePairsLength_AnyKeyValuePairsLength(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"result": attr.DynamicBytesLength() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": attr.StringKeyValuePairsLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": attr.AnyKeyValuePairsLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	// nil
	var nilAttr *corepayload.Attributes
	actual := args.Map{"result": nilAttr.DynamicBytesLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": nilAttr.StringKeyValuePairsLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": nilAttr.AnyKeyValuePairsLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S1_34_Attributes_HasPagingInfo_HasKeyValuePairs_HasFromTo(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"result": attr.HasPagingInfo()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": attr.HasKeyValuePairs()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": attr.HasFromTo()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	// nil
	var nilAttr *corepayload.Attributes
	actual := args.Map{"result": nilAttr.HasPagingInfo()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S1_35_Attributes_IsValid_IsInvalid_IsSafeValid_HasIssuesOrEmpty(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	// Empty() creates non-nil Attributes with no error → IsValid = (it != nil && IsEmptyError) = true
	// BUT IsInvalid = (it == nil || HasIssuesOrEmpty)
	// HasIssuesOrEmpty = IsEmpty() || !IsValid() || (BasicErr && HasError)
	// IsEmpty = len(DynamicPayloads) == 0 → true for Empty()
	// So HasIssuesOrEmpty = true, IsInvalid = true

	// Act
	actual := args.Map{"result": attr.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true — non-nil, no error", actual)
	actual := args.Map{"result": attr.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true — empty attrs are invalid (HasIssuesOrEmpty=true)", actual)
	// nil
	var nilAttr *corepayload.Attributes
	actual := args.Map{"result": nilAttr.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": nilAttr.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovPL_S1_36_Attributes_Paging_Auth_Session_Queries(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"result": attr.IsPagingInfoEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": attr.IsKeyValuePairsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": attr.IsAnyKeyValuePairsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": attr.IsUserInfoEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": attr.IsAuthInfoEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": attr.IsSessionInfoEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": attr.HasUserInfo()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": attr.HasAuthInfo()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": attr.HasSessionInfo()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": attr.SessionInfo() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": attr.AuthType() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": attr.ResourceName() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": attr.HasStringKeyValuePairs()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": attr.HasAnyKeyValuePairs()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": attr.HasDynamicPayloads()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": attr.VirtualUser() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": attr.SystemUser() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": attr.SessionUser() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S1_37_Attributes_GetStringKeyValue_GetAnyKeyValue(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	_, found := attr.GetStringKeyValue("k")

	// Act
	actual := args.Map{"result": found}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, found2 := attr.GetAnyKeyValue("k")
	actual := args.Map{"result": found2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	// nil
	var nilAttr *corepayload.Attributes
	_, found3 := nilAttr.GetStringKeyValue("k")
	actual := args.Map{"result": found3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S1_38_Attributes_HasStringKey_HasAnyKey(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"result": attr.HasStringKey("k")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": attr.HasAnyKey("k")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
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
	// Arrange
	a1 := corepayload.New.Attributes.Empty()
	a2 := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"result": a1.IsEqual(a2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	var nilAttr *corepayload.Attributes
	actual := args.Map{"result": nilAttr.IsEqual(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil==nil", actual)
	actual := args.Map{"result": nilAttr.IsEqual(a1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	// same ptr
	actual := args.Map{"result": a1.IsEqual(a1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovPL_S1_41_Attributes_IsErrorEqual_IsErrorDifferent(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"result": attr.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": attr.IsErrorDifferent(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S1_42_Attributes_Clone_ClonePtr(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	c, err := attr.Clone(false)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = c
	cp, err2 := attr.ClonePtr(true)
	actual := args.Map{"result": err2 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = cp
	// nil
	var nilAttr *corepayload.Attributes
	cp2, err3 := nilAttr.ClonePtr(true)
	actual := args.Map{"result": err3 != nil || cp2 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil, nil", actual)
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
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()

	// Act
	actual := args.Map{"result": pw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S1_51_NewPW_Deserialize(t *testing.T) {
	// Arrange
	pw := newTestPW()
	b, _ := pw.Serialize()
	pw2, err := corepayload.New.PayloadWrapper.Deserialize(b)

	// Act
	actual := args.Map{"result": err != nil || pw2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S1_52_NewPW_CastOrDeserializeFrom(t *testing.T) {
	// Arrange
	pw := newTestPW()
	pw2, err := corepayload.New.PayloadWrapper.CastOrDeserializeFrom(pw)

	// Act
	actual := args.Map{"result": err != nil || pw2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// nil
	_, err2 := corepayload.New.PayloadWrapper.CastOrDeserializeFrom(nil)
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovPL_S1_53_NewPW_DeserializeToMany(t *testing.T) {
	// Arrange
	pws := []*corepayload.PayloadWrapper{newTestPW(), newTestPW()}
	b, _ := corejson.Serialize.Raw(pws)
	many, err := corepayload.New.PayloadWrapper.DeserializeToMany(b)

	// Act
	actual := args.Map{"result": err != nil || len(many) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_CovPL_S1_54_NewPW_DeserializeToCollection(t *testing.T) {
	// Arrange
	pws := []*corepayload.PayloadWrapper{newTestPW()}
	// DeserializeToCollection calls PayloadsCollection.Deserialize which expects
	// {"Items":[...]} format, not raw array — serialize the collection struct
	pc := &corepayload.PayloadsCollection{Items: pws}
	b, _ := corejson.Serialize.Raw(pc)
	col, err := corepayload.New.PayloadWrapper.DeserializeToCollection(b)

	// Act
	actual := args.Map{"result": err != nil || col == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S1_55_NewPW_DeserializeUsingJsonResult(t *testing.T) {
	// Arrange
	pw := newTestPW()
	jr := pw.JsonPtr()
	pw2, err := corepayload.New.PayloadWrapper.DeserializeUsingJsonResult(jr)

	// Act
	actual := args.Map{"result": err != nil || pw2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S1_56_NewPW_UsingBytes(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.UsingBytes(
		"n", "1", "t", "c", "e", []byte(`{"a":1}`))

	// Act
	actual := args.Map{"result": pw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S1_57_NewPW_Record_Records_NameIdRecord(t *testing.T) {
	// Arrange
	_, err := corepayload.New.PayloadWrapper.Record("n", "1", "t", "c", map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_, err2 := corepayload.New.PayloadWrapper.Records("n", "1", "t", "c", []map[string]int{{"a": 1}})
	actual := args.Map{"result": err2 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_, err3 := corepayload.New.PayloadWrapper.NameIdRecord("n", "1", map[string]int{"a": 1})
	actual := args.Map{"result": err3 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_, err4 := corepayload.New.PayloadWrapper.NameIdTaskRecord("n", "1", "t", map[string]int{"a": 1})
	actual := args.Map{"result": err4 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_, err5 := corepayload.New.PayloadWrapper.NameIdCategory("n", "1", "c", map[string]int{"a": 1})
	actual := args.Map{"result": err5 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovPL_S1_58_NewPW_All_ManyRecords(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	pw := corepayload.New.PayloadWrapper.All("n", "1", "t", "c", "e", false, attr, []byte(`{}`))

	// Act
	actual := args.Map{"result": pw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_, err := corepayload.New.PayloadWrapper.ManyRecords("n", "1", "t", "c", []int{1, 2})
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovPL_S1_59_NewPW_NameTaskNameRecord(t *testing.T) {
	// Arrange
	_, err := corepayload.New.PayloadWrapper.NameTaskNameRecord("1", "t", map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

// --- EmptyCreator ---

func Test_CovPL_S1_60_EmptyCreator(t *testing.T) {
	_ = corepayload.Empty.Attributes()
	_ = corepayload.Empty.AttributesDefaults()
	_ = corepayload.Empty.PayloadWrapper()
	_ = corepayload.Empty.PayloadsCollection()
}
