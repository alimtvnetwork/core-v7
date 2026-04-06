package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ─── Result.Map ───

func Test_C28_01_Result_Map_NilReceiver(t *testing.T) {
	var r *corejson.Result
	m := r.Map()
	actual := args.Map{"result": m == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty map", actual)
}

func Test_C28_02_Result_Map_WithBytesAndError(t *testing.T) {
	r := &corejson.Result{
		Bytes:    []byte(`"hello"`),
		Error:    errors.New("some err"),
		TypeName: "TestType",
	}
	m := r.Map()
	actual := args.Map{"result": m["Type"] != "TestType"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type", actual)
	actual := args.Map{"result": _, ok := m["Error"]; !ok}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error key", actual)
}

func Test_C28_03_Result_Map_NoBytesNoError(t *testing.T) {
	r := &corejson.Result{TypeName: "X"}
	m := r.Map()
	actual := args.Map{"result": _, ok := m["Type"]; !ok}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type key", actual)
}

// ─── Result.DeserializedFieldsToMap ───

func Test_C28_04_Result_DeserializedFieldsToMap_Nil(t *testing.T) {
	var r *corejson.Result
	m, err := r.DeserializedFieldsToMap()
	actual := args.Map{"result": err != nil || m == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty map on nil", actual)
}

func Test_C28_05_Result_DeserializedFieldsToMap_EmptyBytes(t *testing.T) {
	r := &corejson.Result{}
	m, err := r.DeserializedFieldsToMap()
	_ = m
	_ = err
}

func Test_C28_06_Result_SafeDeserializedFieldsToMap(t *testing.T) {
	r := &corejson.Result{}
	m := r.SafeDeserializedFieldsToMap()
	_ = m
}

// ─── Result.FieldsNames ───

func Test_C28_07_Result_FieldsNames_Empty(t *testing.T) {
	r := &corejson.Result{}
	names, err := r.FieldsNames()
	_ = names
	_ = err
}

func Test_C28_08_Result_SafeFieldsNames(t *testing.T) {
	r := &corejson.Result{}
	names := r.SafeFieldsNames()
	_ = names
}

// ─── Result.BytesTypeName, SafeBytesTypeName ───

func Test_C28_09_BytesTypeName_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"result": r.BytesTypeName() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C28_10_BytesTypeName_Normal(t *testing.T) {
	r := &corejson.Result{TypeName: "Foo"}
	actual := args.Map{"result": r.BytesTypeName() != "Foo"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Foo", actual)
}

func Test_C28_11_SafeBytesTypeName_Empty(t *testing.T) {
	r := &corejson.Result{}
	s := r.SafeBytesTypeName()
	_ = s
}

func Test_C28_12_SafeBytesTypeName_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	s := r.SafeBytesTypeName()
	actual := args.Map{"result": s != "T"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected T", actual)
}

// ─── Result.SafeString, JsonStringPtr nil branch ───

func Test_C28_13_SafeString(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	actual := args.Map{"result": r.SafeString() != `"hi"`}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C28_14_JsonStringPtr_Nil(t *testing.T) {
	var r *corejson.Result
	s := r.JsonStringPtr()
	actual := args.Map{"result": *s != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string", actual)
}

func Test_C28_15_JsonStringPtr_NoBytes(t *testing.T) {
	r := &corejson.Result{}
	s := r.JsonStringPtr()
	actual := args.Map{"result": *s != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ─── Result.PrettyJsonBuffer ───

func Test_C28_16_PrettyJsonBuffer_Empty(t *testing.T) {
	r := &corejson.Result{}
	buf, err := r.PrettyJsonBuffer("", "  ")
	actual := args.Map{"result": err != nil || buf == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C28_17_PrettyJsonBuffer_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	buf, err := r.PrettyJsonBuffer("", "  ")
	actual := args.Map{"result": err != nil || buf.String() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

// ─── Result.PrettyJsonString ───

func Test_C28_18_PrettyJsonString_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"result": r.PrettyJsonString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C28_19_PrettyJsonString_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	s := r.PrettyJsonString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C28_20_PrettyJsonString_InvalidJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`not-json{`)}
	s := r.PrettyJsonString()
	actual := args.Map{"result": s != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty on invalid json", actual)
}

// ─── Result.PrettyJsonStringOrErrString ───

func Test_C28_21_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	var r *corejson.Result
	s := r.PrettyJsonStringOrErrString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected message", actual)
}

func Test_C28_22_PrettyJsonStringOrErrString_HasError(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("fail")}
	s := r.PrettyJsonStringOrErrString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error string", actual)
}

func Test_C28_23_PrettyJsonStringOrErrString_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	s := r.PrettyJsonStringOrErrString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty string", actual)
}

// ─── Result.Length ───

func Test_C28_24_Length_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"result": r.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C28_25_Length_Normal(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	actual := args.Map{"result": r.Length() != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

// ─── Result.ErrorString ───

func Test_C28_26_ErrorString_NoError(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"result": r.ErrorString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C28_27_ErrorString_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("oops")}
	actual := args.Map{"result": r.ErrorString() != "oops"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected oops", actual)
}

// ─── Result.IsErrorEqual ───

func Test_C28_28_IsErrorEqual_BothNil(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"result": r.IsErrorEqual(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C28_29_IsErrorEqual_OneNil(t *testing.T) {
	r := &corejson.Result{Error: errors.New("x")}
	actual := args.Map{"result": r.IsErrorEqual(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C28_30_IsErrorEqual_LeftNilRightNotNil(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"result": r.IsErrorEqual(errors.New("x"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C28_31_IsErrorEqual_Same(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`x`), Error: errors.New("err")}
	actual := args.Map{"result": r.IsErrorEqual(errors.New("err"))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ─── Result.String ───

func Test_C28_32_String_NoError(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"hi"`), TypeName: "T"}
	s := r.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C28_33_String_WithError(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"hi"`), Error: errors.New("fail"), TypeName: "T"}
	s := r.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C28_34_String_NilBytes(t *testing.T) {
	r := corejson.Result{}
	s := r.String()
	_ = s
}

// ─── SafeNonIssueBytes, SafeBytes, Values, SafeValues, SafeValuesPtr ───

func Test_C28_35_SafeNonIssueBytes_HasIssue(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	b := r.SafeNonIssueBytes()
	actual := args.Map{"result": len(b) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C28_36_SafeNonIssueBytes_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	b := r.SafeNonIssueBytes()
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C28_37_SafeBytes_NilReceiver(t *testing.T) {
	var r *corejson.Result
	b := r.SafeBytes()
	actual := args.Map{"result": len(b) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C28_38_Values(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"result": len(r.Values()) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C28_39_SafeValues_NilBytes(t *testing.T) {
	var r *corejson.Result
	b := r.SafeValues()
	actual := args.Map{"result": len(b) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C28_40_SafeValuesPtr_HasIssue(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	b := r.SafeValuesPtr()
	actual := args.Map{"result": len(b) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C28_41_SafeValuesPtr_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	b := r.SafeValuesPtr()
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

// ─── Result.Raw, RawMust, RawString, RawStringMust, RawErrString, RawPrettyString ───

func Test_C28_42_Raw_Nil(t *testing.T) {
	var r *corejson.Result
	_, err := r.Raw()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_43_Raw_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	b, err := r.Raw()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C28_44_RawMust(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	b := r.RawMust()
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C28_45_RawString(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	s, err := r.RawString()
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C28_46_RawStringMust(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	s := r.RawStringMust()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_C28_47_RawErrString(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	b, errMsg := r.RawErrString()
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	_ = errMsg
}

func Test_C28_48_RawPrettyString(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	s, err := r.RawPrettyString()
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

// ─── Result.MeaningfulError branches ───

func Test_C28_49_MeaningfulError_Nil(t *testing.T) {
	var r *corejson.Result
	err := r.MeaningfulError()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_50_MeaningfulError_NoErrorHasBytes(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	err := r.MeaningfulError()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C28_51_MeaningfulError_EmptyBytes(t *testing.T) {
	r := &corejson.Result{}
	err := r.MeaningfulError()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_52_MeaningfulError_EmptyBytesWithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("inner")}
	err := r.MeaningfulError()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_53_MeaningfulError_HasErrorHasBytes(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("fail")}
	err := r.MeaningfulError()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_54_MeaningfulErrorMessage_NoError(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"result": r.MeaningfulErrorMessage() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C28_55_MeaningfulErrorMessage_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	actual := args.Map{"result": r.MeaningfulErrorMessage() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ─── Result.IsEmptyError, HasSafeItems, IsAnyNull, HasIssuesOrEmpty ───

func Test_C28_56_IsEmptyError(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"result": r.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C28_57_HasSafeItems_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"result": r.HasSafeItems()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C28_58_HasSafeItems_Invalid(t *testing.T) {
	r := &corejson.Result{Error: errors.New("x")}
	actual := args.Map{"result": r.HasSafeItems()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C28_59_IsAnyNull_NilBytes(t *testing.T) {
	r := &corejson.Result{}
	// Bytes is nil by default
	actual := args.Map{"result": r.IsAnyNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ─── Result.HandleErrorWithMsg ───

func Test_C28_60_HandleErrorWithMsg(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	defer func() {
		actual := args.Map{"result": recover() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	r.HandleErrorWithMsg("custom msg")
}

// ─── Result.HasBytes, HasJsonBytes, IsEmptyJsonBytes, IsEmptyJson, HasJson ───

func Test_C28_61_HasBytes(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"result": r.HasBytes()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C28_62_HasJsonBytes(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"result": r.HasJsonBytes()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C28_63_IsEmptyJsonBytes_EmptyJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{}`)}
	actual := args.Map{"result": r.IsEmptyJsonBytes()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for {}", actual)
}

func Test_C28_64_IsEmptyJsonBytes_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"result": r.IsEmptyJsonBytes()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C28_65_IsEmptyJson(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"result": r.IsEmptyJson()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C28_66_HasJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"result": r.HasJson()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C28_67_HasAnyItem(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"result": r.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C28_68_IsEmpty(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"result": r.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ─── Result.InjectInto ───

func Test_C28_69_InjectInto(t *testing.T) {
	r := corejson.NewResult.Any(map[string]string{"a": "1"})
	r2 := corejson.Result{}
	err := r.Ptr().InjectInto(&r2)
	_ = err
}

// ─── Result.Deserialize, DeserializeMust, UnmarshalMust ───

func Test_C28_70_Deserialize(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hello"`)}
	var s string
	err := r.Deserialize(&s)
	actual := args.Map{"result": err != nil || s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C28_71_DeserializeMust(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hello"`)}
	var s string
	r.DeserializeMust(&s)
	actual := args.Map{"result": s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C28_72_UnmarshalMust(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`42`)}
	var n int
	r.UnmarshalMust(&n)
	actual := args.Map{"result": n != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

// ─── Result.Unmarshal branches ───

func Test_C28_73_Unmarshal_NilResult(t *testing.T) {
	var r *corejson.Result
	err := r.Unmarshal(&struct{}{})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_74_Unmarshal_HasError(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("existing")}
	var s string
	err := r.Unmarshal(&s)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_75_Unmarshal_BadJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`not-json`), TypeName: "T"}
	var s string
	err := r.Unmarshal(&s)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ─── Result.SerializeSkipExistingIssues, serializeInternal, Serialize, SerializeMust ───

func Test_C28_76_SerializeSkipExistingIssues_HasIssues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	b, err := r.SerializeSkipExistingIssues()
	actual := args.Map{"result": b != nil || err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil nil", actual)
}

func Test_C28_77_SerializeSkipExistingIssues_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	b, err := r.SerializeSkipExistingIssues()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C28_78_Serialize_Nil(t *testing.T) {
	var r *corejson.Result
	_, err := r.Serialize()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_79_Serialize_HasError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	_, err := r.Serialize()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_80_Serialize_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	b, err := r.Serialize()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C28_81_SerializeMust(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	b := r.SerializeMust()
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

// ─── Result.UnmarshalSkipExistingIssues ───

func Test_C28_82_UnmarshalSkipExistingIssues_HasIssues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	err := r.UnmarshalSkipExistingIssues(&struct{}{})
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C28_83_UnmarshalSkipExistingIssues_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hello"`)}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	actual := args.Map{"result": err != nil || s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C28_84_UnmarshalSkipExistingIssues_BadJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`not-json`)}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ─── Result.UnmarshalResult ───

func Test_C28_85_UnmarshalResult(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	serialized := corejson.NewResult.AnyPtr(inner)
	_, _ = serialized.UnmarshalResult()
}

// ─── Result.JsonModel, JsonModelAny ───

func Test_C28_86_JsonModel_Nil(t *testing.T) {
	var r *corejson.Result
	m := r.JsonModel()
	actual := args.Map{"result": m.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_87_JsonModel_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	m := r.JsonModel()
	actual := args.Map{"result": len(m.Bytes) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C28_88_JsonModelAny(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	a := r.JsonModelAny()
	_ = a
}

// ─── Result.Json, JsonPtr ───

func Test_C28_89_Json(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	j := r.Json()
	actual := args.Map{"result": j.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

func Test_C28_90_JsonPtr(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	j := r.JsonPtr()
	actual := args.Map{"result": j.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

// ─── Result.ParseInjectUsingJson, ParseInjectUsingJsonMust ───

func Test_C28_91_ParseInjectUsingJson_Success(t *testing.T) {
	r := &corejson.Result{}
	src := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"test"`), TypeName: "T"})
	_, err := r.ParseInjectUsingJson(src)
	_ = err
}

func Test_C28_92_ParseInjectUsingJson_Failure(t *testing.T) {
	r := &corejson.Result{}
	src := &corejson.Result{Error: errors.New("fail")}
	_, err := r.ParseInjectUsingJson(src)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_93_ParseInjectUsingJsonMust_Success(t *testing.T) {
	r := &corejson.Result{}
	src := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"test"`), TypeName: "T"})
	_ = r.ParseInjectUsingJsonMust(src)
}

// ─── Result.CloneError ───

func Test_C28_94_CloneError_HasError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("orig")}
	err := r.CloneError()
	actual := args.Map{"result": err == nil || err.Error() != "orig"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C28_95_CloneError_NoError(t *testing.T) {
	r := &corejson.Result{}
	err := r.CloneError()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// ─── Result.Ptr, NonPtr, ToPtr, ToNonPtr ───

func Test_C28_96_Ptr(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	p := r.Ptr()
	actual := args.Map{"result": p == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ptr", actual)
}

func Test_C28_97_NonPtr_Nil(t *testing.T) {
	var r *corejson.Result
	v := r.NonPtr()
	actual := args.Map{"result": v.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_98_NonPtr_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	v := r.NonPtr()
	actual := args.Map{"result": len(v.Bytes) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C28_99_ToPtr(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	_ = r.ToPtr()
}

func Test_C28_100_ToNonPtr(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	_ = r.ToNonPtr()
}

// ─── Result.IsEqualPtr ───

func Test_C28_101_IsEqualPtr_BothNil(t *testing.T) {
	var a, b *corejson.Result
	actual := args.Map{"result": a.IsEqualPtr(b)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C28_102_IsEqualPtr_OneNil(t *testing.T) {
	a := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"result": a.IsEqualPtr(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C28_103_IsEqualPtr_SamePtr(t *testing.T) {
	a := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"result": a.IsEqualPtr(a)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C28_104_IsEqualPtr_DiffLength(t *testing.T) {
	a := &corejson.Result{Bytes: []byte(`"x"`)}
	b := &corejson.Result{Bytes: []byte(`"xy"`)}
	actual := args.Map{"result": a.IsEqualPtr(b)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C28_105_IsEqualPtr_DiffError(t *testing.T) {
	a := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("a")}
	b := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("b")}
	actual := args.Map{"result": a.IsEqualPtr(b)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C28_106_IsEqualPtr_DiffType(t *testing.T) {
	a := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "A"}
	b := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "B"}
	actual := args.Map{"result": a.IsEqualPtr(b)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C28_107_IsEqualPtr_Equal(t *testing.T) {
	a := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	b := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	actual := args.Map{"result": a.IsEqualPtr(b)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ─── Result.CombineErrorWithRefString, CombineErrorWithRefError ───

func Test_C28_108_CombineErrorWithRefString_NoError(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"result": r.CombineErrorWithRefString("ref") != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C28_109_CombineErrorWithRefString_HasError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	s := r.CombineErrorWithRefString("ref1", "ref2")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C28_110_CombineErrorWithRefError_NoError(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"result": r.CombineErrorWithRefError("ref") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C28_111_CombineErrorWithRefError_HasError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	err := r.CombineErrorWithRefError("ref1")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ─── Result.IsEqual ───

func Test_C28_112_IsEqual_Same(t *testing.T) {
	a := corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	b := corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	actual := args.Map{"result": a.IsEqual(b)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C28_113_IsEqual_DiffLength(t *testing.T) {
	a := corejson.Result{Bytes: []byte(`"x"`)}
	b := corejson.Result{Bytes: []byte(`"xy"`)}
	actual := args.Map{"result": a.IsEqual(b)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C28_114_IsEqual_DiffError(t *testing.T) {
	a := corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("a")}
	b := corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("b")}
	actual := args.Map{"result": a.IsEqual(b)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ─── Result.BytesError ───

func Test_C28_115_BytesError_Nil(t *testing.T) {
	var r *corejson.Result
	be := r.BytesError()
	actual := args.Map{"result": be != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C28_116_BytesError_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("err")}
	be := r.BytesError()
	actual := args.Map{"result": be == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ─── Result.Dispose ───

func Test_C28_117_Dispose_Nil(t *testing.T) {
	var r *corejson.Result
	r.Dispose() // should not panic
}

func Test_C28_118_Dispose_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e"), TypeName: "T"}
	r.Dispose()
	actual := args.Map{"result": r.Error != nil || r.Bytes != nil || r.TypeName != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected disposed", actual)
}

// ─── Result.CloneIf, ClonePtr, Clone ───

func Test_C28_119_CloneIf_NoClone(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	c := r.CloneIf(false, false)
	_ = c
}

func Test_C28_120_CloneIf_ShallowClone(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	c := r.CloneIf(true, false)
	_ = c
}

func Test_C28_121_CloneIf_DeepClone(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	c := r.CloneIf(true, true)
	_ = c
}

func Test_C28_122_ClonePtr_Nil(t *testing.T) {
	var r *corejson.Result
	c := r.ClonePtr(false)
	actual := args.Map{"result": c != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C28_123_ClonePtr_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	c := r.ClonePtr(true)
	actual := args.Map{"result": c == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C28_124_Clone_EmptyBytes(t *testing.T) {
	r := corejson.Result{}
	c := r.Clone(true)
	_ = c
}

func Test_C28_125_Clone_ShallowClone(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	c := r.Clone(false)
	_ = c
}

func Test_C28_126_Clone_DeepClone(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	c := r.Clone(true)
	_ = c
}

// ─── Result.AsJsonContractsBinder, AsJsoner, JsonParseSelfInject, AsJsonParseSelfInjector ───

func Test_C28_127_AsJsonContractsBinder(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	_ = r.AsJsonContractsBinder()
}

func Test_C28_128_AsJsoner(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	_ = r.AsJsoner()
}

func Test_C28_129_JsonParseSelfInject(t *testing.T) {
	r := corejson.Result{}
	src := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	err := r.JsonParseSelfInject(src)
	_ = err
}

func Test_C28_130_AsJsonParseSelfInjector(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	_ = r.AsJsonParseSelfInjector()
}

// ─── Result.safeJsonStringInternal (via MeaningfulError with nil) ───

func Test_C28_131_safeJsonStringInternal_NilBranch(t *testing.T) {
	// indirectly tested via MeaningfulError on nil
	var r *corejson.Result
	_ = r.MeaningfulError()
}

// ─── BytesToString, BytesToPrettyString empty branches ───

func Test_C28_132_BytesToString_Empty(t *testing.T) {
	s := corejson.BytesToString([]byte{})
	actual := args.Map{"result": s != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C28_133_BytesToPrettyString_Empty(t *testing.T) {
	s := corejson.BytesToPrettyString([]byte{})
	actual := args.Map{"result": s != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ─── JsonString ───

func Test_C28_134_JsonString_Func(t *testing.T) {
	s, err := corejson.JsonString(map[string]int{"a": 1})
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

// ─── JsonStringOrErrMsg ───

func Test_C28_135_JsonStringOrErrMsg_Valid(t *testing.T) {
	s := corejson.JsonStringOrErrMsg("hello")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C28_136_JsonStringOrErrMsg_Invalid(t *testing.T) {
	ch := make(chan int)
	s := corejson.JsonStringOrErrMsg(ch)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error message", actual)
}

// ─── BytesCloneIf ───

func Test_C28_137_BytesCloneIf_NoClone(t *testing.T) {
	b := corejson.BytesCloneIf(false, []byte("x"))
	_ = b
}

func Test_C28_138_BytesCloneIf_DeepClone(t *testing.T) {
	b := corejson.BytesCloneIf(true, []byte("hello"))
	actual := args.Map{"result": len(b) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_C28_139_BytesCloneIf_Empty(t *testing.T) {
	b := corejson.BytesCloneIf(true, []byte{})
	actual := args.Map{"result": len(b) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ─── New, NewPtr uncovered error branch ───

func Test_C28_140_New_MarshalError(t *testing.T) {
	ch := make(chan int)
	r := corejson.New(ch)
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C28_141_NewPtr_MarshalError(t *testing.T) {
	ch := make(chan int)
	r := corejson.NewPtr(ch)
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}
