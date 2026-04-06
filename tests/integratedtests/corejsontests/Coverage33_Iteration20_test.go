package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ===== Result methods coverage =====

func Test_I20_Result_CloneIf_True(t *testing.T) {
	r := corejson.NewResult.Any(map[string]string{"a": "1"})
	cloned := r.CloneIf(true, true)
	if cloned.Length() == 0 {
		t.Fatal("expected cloned bytes")
	}
}

func Test_I20_Result_CloneIf_False(t *testing.T) {
	r := corejson.NewResult.Any(map[string]string{"a": "1"})
	same := r.CloneIf(false, false)
	if same.Length() != r.Length() {
		t.Fatal("expected same result")
	}
}

func Test_I20_Result_Clone_DeepClone(t *testing.T) {
	r := corejson.NewResult.Any("test")
	c := r.Clone(true)
	if c.Length() == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Result_Clone_ShallowClone(t *testing.T) {
	r := corejson.NewResult.Any("test")
	c := r.Clone(false)
	if c.Length() == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Result_Clone_Empty(t *testing.T) {
	r := corejson.NewResult.Empty()
	c := r.Clone(true)
	if c.Length() != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I20_Result_ClonePtr_Nil(t *testing.T) {
	var r *corejson.Result
	c := r.ClonePtr(true)
	if c != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Result_ClonePtr_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("test")
	c := r.ClonePtr(true)
	if c == nil || c.Length() == 0 {
		t.Fatal("expected cloned ptr")
	}
}

func Test_I20_Result_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	var r *corejson.Result
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected non-empty nil message")
	}
}

func Test_I20_Result_PrettyJsonStringOrErrString_HasError(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("test-err"))
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected error string")
	}
}

func Test_I20_Result_PrettyJsonStringOrErrString_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected pretty json")
	}
}

func Test_I20_Result_HandleErrorWithMsg_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	r.HandleErrorWithMsg("no-op") // Should not panic
}

func Test_I20_Result_HandleErrorWithMsg_Panic(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("bad"))
	defer func() {
		if rec := recover(); rec == nil {
			t.Fatal("expected panic")
		}
	}()
	r.HandleErrorWithMsg("context message")
}

func Test_I20_Result_DeserializeMust_Success(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	var s string
	r.DeserializeMust(&s)
	if s != "hello" {
		t.Fatalf("expected 'hello', got '%s'", s)
	}
}

func Test_I20_Result_DeserializeMust_Panic(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("err"))
	defer func() {
		if rec := recover(); rec == nil {
			t.Fatal("expected panic")
		}
	}()
	var s string
	r.DeserializeMust(&s)
}

func Test_I20_Result_UnmarshalMust_Success(t *testing.T) {
	r := corejson.NewResult.AnyPtr(42)
	var i int
	r.UnmarshalMust(&i)
	if i != 42 {
		t.Fatal("expected 42")
	}
}

func Test_I20_Result_SafeFieldsNames(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"key": "val"})
	names := r.SafeFieldsNames()
	// May return empty due to DeserializedFieldsToMap behavior
	_ = names
}

func Test_I20_Result_SafeDeserializedFieldsToMap(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"k": "v"})
	m := r.SafeDeserializedFieldsToMap()
	_ = m
}

func Test_I20_Result_BytesError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("test")
	be := r.BytesError()
	if be == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I20_Result_BytesError_Nil(t *testing.T) {
	var r *corejson.Result
	be := r.BytesError()
	if be != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Result_Dispose(t *testing.T) {
	r := corejson.NewResult.AnyPtr("test")
	r.Dispose()
	if r.Length() != 0 {
		t.Fatal("expected disposed")
	}
}

func Test_I20_Result_Dispose_Nil(t *testing.T) {
	var r *corejson.Result
	r.Dispose() // should not panic
}

func Test_I20_Result_NonPtr_Nil(t *testing.T) {
	var r *corejson.Result
	nr := r.NonPtr()
	if nr.Error == nil {
		t.Fatal("expected error in nonptr of nil")
	}
}

func Test_I20_Result_NonPtr_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("test")
	nr := r.NonPtr()
	if nr.Length() == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Result_CombineErrorWithRefError_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	err := r.CombineErrorWithRefError("ref")
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Result_CombineErrorWithRefError_HasError(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	err := r.CombineErrorWithRefError("ref1", "ref2")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_Result_CombineErrorWithRefString_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	s := r.CombineErrorWithRefString("ref")
	if s != "" {
		t.Fatal("expected empty")
	}
}

func Test_I20_Result_CloneError_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	if r.CloneError() != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Result_CloneError_HasError(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	if r.CloneError() == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_Result_Ptr_ToPtr(t *testing.T) {
	r := corejson.NewResult.Any("test")
	p := r.Ptr()
	if p == nil {
		t.Fatal("expected ptr")
	}
	np := r.ToPtr()
	if np == nil {
		t.Fatal("expected ptr")
	}
	np2 := r.ToNonPtr()
	if np2.Length() == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Result_IsEqualPtr_BothNil(t *testing.T) {
	var a, b *corejson.Result
	if !a.IsEqualPtr(b) {
		t.Fatal("expected equal")
	}
}

func Test_I20_Result_IsEqualPtr_OneNil(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	if r.IsEqualPtr(nil) {
		t.Fatal("expected not equal")
	}
}

func Test_I20_Result_IsEqualPtr_Same(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	if !r.IsEqualPtr(r) {
		t.Fatal("expected equal (same ptr)")
	}
}

func Test_I20_Result_IsEqualPtr_DiffLength(t *testing.T) {
	a := corejson.NewResult.AnyPtr("x")
	b := corejson.NewResult.AnyPtr("xy")
	if a.IsEqualPtr(b) {
		t.Fatal("expected not equal")
	}
}

func Test_I20_Result_IsEqualPtr_DiffError(t *testing.T) {
	a := corejson.NewResult.Ptr([]byte("x"), errors.New("e1"), "t")
	b := corejson.NewResult.Ptr([]byte("x"), errors.New("e2"), "t")
	if a.IsEqualPtr(b) {
		t.Fatal("expected not equal")
	}
}

func Test_I20_Result_IsEqualPtr_DiffType(t *testing.T) {
	a := corejson.NewResult.Ptr([]byte(`"x"`), nil, "typeA")
	b := corejson.NewResult.Ptr([]byte(`"x"`), nil, "typeB")
	if a.IsEqualPtr(b) {
		t.Fatal("expected not equal")
	}
}

func Test_I20_Result_IsEqual(t *testing.T) {
	a := corejson.NewResult.Any("hello")
	b := corejson.NewResult.Any("hello")
	if !a.IsEqual(b) {
		t.Fatal("expected equal")
	}
}

func Test_I20_Result_IsErrorEqual(t *testing.T) {
	a := corejson.NewResult.ErrorPtr(errors.New("same"))
	if !a.IsErrorEqual(errors.New("same")) {
		t.Fatal("expected equal")
	}
}

func Test_I20_Result_IsErrorEqual_BothNil(t *testing.T) {
	a := corejson.NewResult.AnyPtr("ok")
	if !a.IsErrorEqual(nil) {
		t.Fatal("expected equal (both nil)")
	}
}

func Test_I20_Result_IsErrorEqual_OnlyOneNil(t *testing.T) {
	a := corejson.NewResult.AnyPtr("ok")
	if a.IsErrorEqual(errors.New("e")) {
		t.Fatal("expected not equal")
	}
}

// ===== Result serialization coverage =====

func Test_I20_Result_Serialize_Nil(t *testing.T) {
	var r *corejson.Result
	_, err := r.Serialize()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_Result_Serialize_HasError(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	_, err := r.Serialize()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_Result_Serialize_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	b, err := r.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Result_SerializeMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	b := r.SerializeMust()
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Result_SerializeSkipExistingIssues_HasIssues(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	b, err := r.SerializeSkipExistingIssues()
	if b != nil || err != nil {
		t.Fatal("expected nil,nil for issues")
	}
}

func Test_I20_Result_SerializeSkipExistingIssues_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	b, err := r.SerializeSkipExistingIssues()
	if err != nil || len(b) == 0 {
		t.Fatal("expected success")
	}
}

func Test_I20_Result_UnmarshalSkipExistingIssues_HasIssues(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	if err != nil {
		t.Fatal("expected nil for issues")
	}
}

func Test_I20_Result_UnmarshalSkipExistingIssues_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	if err != nil {
		t.Fatal(err)
	}
	if s != "hello" {
		t.Fatalf("expected 'hello', got '%s'", s)
	}
}

func Test_I20_Result_UnmarshalSkipExistingIssues_Error(t *testing.T) {
	r := corejson.NewResult.UsingBytesTypePtr([]byte("not-json"), "test")
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	if err == nil {
		t.Fatal("expected unmarshal error")
	}
}

func Test_I20_Result_UnmarshalResult(t *testing.T) {
	inner := corejson.NewResult.Any("test")
	outerBytes, _ := inner.Serialize()
	outer := &corejson.Result{Bytes: outerBytes}
	_, _ = outer.UnmarshalResult()
}

func Test_I20_Result_JsonModel_Nil(t *testing.T) {
	var r *corejson.Result
	m := r.JsonModel()
	if m.Error == nil {
		t.Fatal("expected error in model")
	}
}

func Test_I20_Result_JsonModelAny(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	a := r.JsonModelAny()
	if a == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I20_Result_Json_JsonPtr(t *testing.T) {
	r := corejson.NewResult.Any("test")
	j := r.Json()
	if j.Length() == 0 {
		t.Fatal("expected json")
	}
	jp := r.JsonPtr()
	if jp == nil {
		t.Fatal("expected ptr")
	}
}

func Test_I20_Result_ParseInjectUsingJson_Success(t *testing.T) {
	source := corejson.NewResult.AnyPtr(map[string]string{"a": "1"})
	target := corejson.NewResult.AnyPtr(map[string]string{})
	_, err := target.ParseInjectUsingJson(source)
	// This may fail depending on types but exercises the path
	_ = err
}

func Test_I20_Result_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	source := corejson.NewResult.ErrorPtr(errors.New("err"))
	target := corejson.NewResult.AnyPtr("test")
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	target.ParseInjectUsingJsonMust(source)
}

func Test_I20_Result_AsJsonContractsBinder(t *testing.T) {
	r := corejson.NewResult.Any("test")
	b := r.AsJsonContractsBinder()
	if b == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I20_Result_AsJsoner(t *testing.T) {
	r := corejson.NewResult.Any("test")
	j := r.AsJsoner()
	if j == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I20_Result_AsJsonParseSelfInjector(t *testing.T) {
	r := corejson.NewResult.Any("test")
	inj := r.AsJsonParseSelfInjector()
	if inj == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I20_Result_JsonParseSelfInject(t *testing.T) {
	r := corejson.NewResult.Any(map[string]string{"a": "1"})
	source := corejson.NewResult.AnyPtr(map[string]string{"b": "2"})
	err := r.JsonParseSelfInject(source)
	_ = err
}

// ===== Result other methods =====

func Test_I20_Result_RawMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	b := r.RawMust()
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Result_RawStringMust_Success(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	s := r.RawStringMust()
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_Result_RawStringMust_Panic(t *testing.T) {
	var r *corejson.Result
	defer func() {
		if rec := recover(); rec == nil {
			t.Fatal("expected panic")
		}
	}()
	r.RawStringMust()
}

func Test_I20_Result_RawErrString(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	b, msg := r.RawErrString()
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
	_ = msg
}

func Test_I20_Result_RawPrettyString(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	s, err := r.RawPrettyString()
	if err != nil {
		t.Fatal(err)
	}
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_Result_HandleError_Panic(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	defer func() {
		if rec := recover(); rec == nil {
			t.Fatal("expected panic")
		}
	}()
	r.HandleError()
}

func Test_I20_Result_MustBeSafe_Panic(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	defer func() {
		if rec := recover(); rec == nil {
			t.Fatal("expected panic")
		}
	}()
	r.MustBeSafe()
}

func Test_I20_Result_SafeNonIssueBytes(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	if len(r.SafeNonIssueBytes()) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Result_SafeNonIssueBytes_HasIssues(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	if len(r.SafeNonIssueBytes()) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I20_Result_SafeValuesPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	if len(r.SafeValuesPtr()) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Result_SafeValuesPtr_HasIssues(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	if len(r.SafeValuesPtr()) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I20_Result_Values(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	if len(r.Values()) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Result_Raw_Nil(t *testing.T) {
	var r *corejson.Result
	_, err := r.Raw()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_Result_MeaningfulError_EmptyBytes(t *testing.T) {
	r := &corejson.Result{Bytes: nil, TypeName: "Test"}
	err := r.MeaningfulError()
	if err == nil {
		t.Fatal("expected error for empty bytes")
	}
}

func Test_I20_Result_MeaningfulError_HasErrorAndBytes(t *testing.T) {
	r := &corejson.Result{
		Bytes:    []byte(`"test"`),
		Error:    errors.New("some err"),
		TypeName: "Test",
	}
	err := r.MeaningfulError()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_Result_String_HasError(t *testing.T) {
	r := corejson.NewResult.Create([]byte(`"x"`), errors.New("e"), "T")
	s := r.String()
	if s == "" {
		t.Fatal("expected non-empty string with error")
	}
}

func Test_I20_Result_SafeBytesTypeName_Empty(t *testing.T) {
	r := corejson.NewResult.EmptyPtr()
	if r.SafeBytesTypeName() != "" {
		t.Fatal("expected empty")
	}
}

func Test_I20_Result_BytesTypeName_Nil(t *testing.T) {
	var r *corejson.Result
	if r.BytesTypeName() != "" {
		t.Fatal("expected empty")
	}
}

// ===== anyTo coverage =====

func Test_I20_AnyTo_SerializedRaw(t *testing.T) {
	b, err := corejson.AnyTo.SerializedRaw("hello")
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_AnyTo_SerializedString(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString("hello")
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_SerializedSafeString(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString("hello")
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_SerializedStringMust(t *testing.T) {
	s := corejson.AnyTo.SerializedStringMust("hello")
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_SafeJsonString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonString("hello")
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_PrettyStringWithError_String(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError("hello")
	if err != nil || s != "hello" {
		t.Fatal("expected passthrough")
	}
}

func Test_I20_AnyTo_PrettyStringWithError_Bytes(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":"b"}`))
	if err != nil || s == "" {
		t.Fatal("expected pretty string")
	}
}

func Test_I20_AnyTo_PrettyStringWithError_Result(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_PrettyStringWithError_ResultPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_PrettyStringWithError_ResultWithError(t *testing.T) {
	r := corejson.NewResult.Create([]byte(`"x"`), errors.New("e"), "T")
	_, err := corejson.AnyTo.PrettyStringWithError(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_AnyTo_PrettyStringWithError_ResultPtrWithError(t *testing.T) {
	r := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("e"), "T")
	_, err := corejson.AnyTo.PrettyStringWithError(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_AnyTo_PrettyStringWithError_AnyItem(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError(42)
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_SafeJsonPrettyString_String(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString("hi")
	if s != "hi" {
		t.Fatal("expected passthrough")
	}
}

func Test_I20_AnyTo_SafeJsonPrettyString_Bytes(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":"b"}`))
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_SafeJsonPrettyString_Result(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	s := corejson.AnyTo.SafeJsonPrettyString(r)
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_SafeJsonPrettyString_ResultPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s := corejson.AnyTo.SafeJsonPrettyString(r)
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_SafeJsonPrettyString_Any(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString(42)
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_JsonString_String(t *testing.T) {
	s := corejson.AnyTo.JsonString("hi")
	if s != "hi" {
		t.Fatal("expected passthrough")
	}
}

func Test_I20_AnyTo_JsonString_Bytes(t *testing.T) {
	s := corejson.AnyTo.JsonString([]byte(`{"a":"b"}`))
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_JsonString_Result(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	s := corejson.AnyTo.JsonString(r)
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_JsonString_ResultPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s := corejson.AnyTo.JsonString(r)
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_JsonString_Any(t *testing.T) {
	s := corejson.AnyTo.JsonString(42)
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_JsonStringWithErr_String(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr("hi")
	if err != nil || s != "hi" {
		t.Fatal("expected passthrough")
	}
}

func Test_I20_AnyTo_JsonStringWithErr_Bytes(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr([]byte(`"x"`))
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_JsonStringWithErr_Result_NoError(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	s, err := corejson.AnyTo.JsonStringWithErr(r)
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_JsonStringWithErr_Result_HasError(t *testing.T) {
	r := corejson.NewResult.Create([]byte(`"x"`), errors.New("e"), "T")
	_, err := corejson.AnyTo.JsonStringWithErr(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_AnyTo_JsonStringWithErr_ResultPtr_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s, err := corejson.AnyTo.JsonStringWithErr(r)
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_JsonStringWithErr_ResultPtr_HasError(t *testing.T) {
	r := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("e"), "T")
	_, err := corejson.AnyTo.JsonStringWithErr(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_AnyTo_JsonStringWithErr_Any(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr(42)
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_JsonStringMust(t *testing.T) {
	s := corejson.AnyTo.JsonStringMust("hi")
	if s != "hi" {
		t.Fatal("expected 'hi'")
	}
}

func Test_I20_AnyTo_PrettyStringMust(t *testing.T) {
	s := corejson.AnyTo.PrettyStringMust("hi")
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_AnyTo_UsingSerializer(t *testing.T) {
	r := corejson.AnyTo.UsingSerializer(nil)
	if r != nil {
		t.Fatal("expected nil for nil serializer")
	}
}

func Test_I20_AnyTo_SerializedFieldsMap(t *testing.T) {
	m, err := corejson.AnyTo.SerializedFieldsMap(map[string]string{"k": "v"})
	_ = m
	_ = err
}

func Test_I20_AnyTo_SerializedJsonResult_Nil(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(nil)
	if r == nil || r.Error == nil {
		t.Fatal("expected error result for nil")
	}
}

func Test_I20_AnyTo_SerializedJsonResult_Error_NilErr(t *testing.T) {
	var errNil error
	r := corejson.AnyTo.SerializedJsonResult(errNil)
	if r == nil {
		t.Fatal("expected result")
	}
}

// ===== Serializer coverage =====

func Test_I20_Serializer_StringsApply(t *testing.T) {
	r := corejson.Serialize.StringsApply([]string{"a", "b"})
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_Serializer_FromBytes(t *testing.T) {
	r := corejson.Serialize.FromBytes([]byte(`"test"`))
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_Serializer_FromStrings(t *testing.T) {
	r := corejson.Serialize.FromStrings([]string{"a"})
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_Serializer_FromStringsSpread(t *testing.T) {
	r := corejson.Serialize.FromStringsSpread("a", "b")
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_Serializer_FromString(t *testing.T) {
	r := corejson.Serialize.FromString("hello")
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_Serializer_FromInteger(t *testing.T) {
	r := corejson.Serialize.FromInteger(42)
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_Serializer_FromInteger64(t *testing.T) {
	r := corejson.Serialize.FromInteger64(64)
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_Serializer_FromBool(t *testing.T) {
	r := corejson.Serialize.FromBool(true)
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_Serializer_FromIntegers(t *testing.T) {
	r := corejson.Serialize.FromIntegers([]int{1, 2})
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_Serializer_UsingAnyPtr(t *testing.T) {
	r := corejson.Serialize.UsingAnyPtr("test")
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_Serializer_UsingAny(t *testing.T) {
	r := corejson.Serialize.UsingAny("test")
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_Serializer_Raw(t *testing.T) {
	b, err := corejson.Serialize.Raw("test")
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Serializer_Marshal(t *testing.T) {
	b, err := corejson.Serialize.Marshal("test")
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Serializer_ApplyMust(t *testing.T) {
	r := corejson.Serialize.ApplyMust("test")
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_Serializer_ToBytesMust(t *testing.T) {
	b := corejson.Serialize.ToBytesMust("test")
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Serializer_ToSafeBytesMust(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesMust("test")
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Serializer_ToSafeBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesSwallowErr("test")
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Serializer_ToBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToBytesSwallowErr("test")
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Serializer_ToBytesErr(t *testing.T) {
	b, err := corejson.Serialize.ToBytesErr("test")
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_Serializer_ToString(t *testing.T) {
	s := corejson.Serialize.ToString("test")
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_Serializer_ToStringMust(t *testing.T) {
	s := corejson.Serialize.ToStringMust("test")
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_Serializer_ToStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToStringErr("test")
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_Serializer_ToPrettyStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToPrettyStringErr("test")
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_Serializer_ToPrettyStringIncludingErr(t *testing.T) {
	s := corejson.Serialize.ToPrettyStringIncludingErr("test")
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_I20_Serializer_Pretty(t *testing.T) {
	s := corejson.Serialize.Pretty("test")
	if s == "" {
		t.Fatal("expected string")
	}
}

// ===== Deserializer coverage =====

func Test_I20_Deserializer_UsingStringPtr_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringPtr(nil, &s)
	if err == nil {
		t.Fatal("expected error for nil bytes")
	}
}

func Test_I20_Deserializer_UsingStringPtr_Valid(t *testing.T) {
	str := `"hello"`
	var s string
	err := corejson.Deserialize.UsingStringPtr(&str, &s)
	if err != nil {
		t.Fatal(err)
	}
	if s != "hello" {
		t.Fatalf("expected 'hello', got '%s'", s)
	}
}

func Test_I20_Deserializer_UsingError_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingError(nil, &s)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Deserializer_UsingError_Valid(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingError(errors.New(`"hello"`), &s)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I20_Deserializer_UsingErrorWhichJsonResult_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &s)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Deserializer_FromString(t *testing.T) {
	var s string
	err := corejson.Deserialize.FromString(`"hi"`, &s)
	if err != nil || s != "hi" {
		t.Fatal("expected 'hi'")
	}
}

func Test_I20_Deserializer_FromStringMust(t *testing.T) {
	var s string
	corejson.Deserialize.FromStringMust(`"hi"`, &s)
	if s != "hi" {
		t.Fatal("expected 'hi'")
	}
}

func Test_I20_Deserializer_FromStringMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	var s string
	corejson.Deserialize.FromStringMust("not-json", &s)
}

func Test_I20_Deserializer_UsingStringOption_IgnoreEmpty(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringOption(true, "", &s)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Deserializer_UsingStringIgnoreEmpty(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &s)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Deserializer_UsingBytesPointer_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointer(nil, &s)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_Deserializer_UsingBytesPointer_Valid(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointer([]byte(`"hi"`), &s)
	if err != nil || s != "hi" {
		t.Fatal("expected 'hi'")
	}
}

func Test_I20_Deserializer_UsingBytesPointerMust(t *testing.T) {
	var s string
	corejson.Deserialize.UsingBytesPointerMust([]byte(`"hi"`), &s)
	if s != "hi" {
		t.Fatal("expected 'hi'")
	}
}

func Test_I20_Deserializer_UsingBytesIf_Skip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"x"`), &s)
	if err != nil || s != "" {
		t.Fatal("expected skip")
	}
}

func Test_I20_Deserializer_UsingBytesIf_Do(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesIf(true, []byte(`"x"`), &s)
	if err != nil || s != "x" {
		t.Fatal("expected 'x'")
	}
}

func Test_I20_Deserializer_UsingBytesPointerIf_Skip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &s)
	if err != nil {
		t.Fatal("expected skip")
	}
}

func Test_I20_Deserializer_UsingBytesPointerIf_Do(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointerIf(true, []byte(`"x"`), &s)
	if err != nil || s != "x" {
		t.Fatal("expected 'x'")
	}
}

func Test_I20_Deserializer_UsingBytesMust(t *testing.T) {
	var s string
	corejson.Deserialize.UsingBytesMust([]byte(`"hi"`), &s)
	if s != "hi" {
		t.Fatal("expected 'hi'")
	}
}

func Test_I20_Deserializer_UsingSafeBytesMust_Empty(t *testing.T) {
	var s string
	corejson.Deserialize.UsingSafeBytesMust([]byte{}, &s)
	// should skip
}

func Test_I20_Deserializer_UsingSafeBytesMust_Valid(t *testing.T) {
	var s string
	corejson.Deserialize.UsingSafeBytesMust([]byte(`"hi"`), &s)
	if s != "hi" {
		t.Fatal("expected 'hi'")
	}
}

func Test_I20_Deserializer_AnyToFieldsMap(t *testing.T) {
	m, err := corejson.Deserialize.AnyToFieldsMap(map[string]string{"k": "v"})
	_ = m
	_ = err
}

func Test_I20_Deserializer_MapAnyToPointer_SkipEmpty(t *testing.T) {
	var s map[string]any
	err := corejson.Deserialize.MapAnyToPointer(true, nil, &s)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Deserializer_MapAnyToPointer_Valid(t *testing.T) {
	var s map[string]any
	err := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"k": "v"}, &s)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I20_Deserializer_UsingDeserializerToOption_SkipNil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &s)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Deserializer_UsingDeserializerToOption_NilNotSkip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerToOption(false, nil, &s)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_Deserializer_UsingDeserializerDefined_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerDefined(nil, &s)
	if err != nil {
		t.Fatal("expected nil (skip)")
	}
}

func Test_I20_Deserializer_UsingDeserializerFuncDefined_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &s)
	if err == nil {
		t.Fatal("expected error for nil func")
	}
}

func Test_I20_Deserializer_UsingDeserializerFuncDefined_Valid(t *testing.T) {
	var s string
	fn := func(toPtr any) error { return nil }
	err := corejson.Deserialize.UsingDeserializerFuncDefined(fn, &s)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Deserializer_UsingJsonerToAny_SkipNil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, &s)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Deserializer_UsingJsonerToAny_NilNotSkip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAny(false, nil, &s)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_Deserializer_UsingJsonerToAnyMust_SkipNil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, &s)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_Deserializer_UsingJsonerToAnyMust_NilNotSkip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAnyMust(false, nil, &s)
	if err == nil {
		t.Fatal("expected error")
	}
}

// ===== NewResult creator coverage =====

func Test_I20_NewResult_UsingBytesError_Nil(t *testing.T) {
	r := corejson.NewResult.UsingBytesError(nil)
	if r.Length() != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I20_NewResult_UsingErrorStringPtr_NilPtr(t *testing.T) {
	r := corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "T")
	if r.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_NewResult_UsingErrorStringPtr_Valid(t *testing.T) {
	s := `"ok"`
	r := corejson.NewResult.UsingErrorStringPtr(nil, &s, "T")
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_NewResult_UsingTypePlusStringPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingTypePlusStringPtr("T", nil)
	if r.Length() != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I20_NewResult_UsingTypePlusStringPtr_Empty(t *testing.T) {
	s := ""
	r := corejson.NewResult.UsingTypePlusStringPtr("T", &s)
	if r.Length() != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I20_NewResult_UsingTypePlusStringPtr_Valid(t *testing.T) {
	s := `"hello"`
	r := corejson.NewResult.UsingTypePlusStringPtr("T", &s)
	if r.Length() == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_NewResult_UsingStringPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingStringPtr(nil)
	if r.Length() != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I20_NewResult_UsingStringPtr_Valid(t *testing.T) {
	s := `"hello"`
	r := corejson.NewResult.UsingStringPtr(&s)
	if r.Length() == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_NewResult_Many(t *testing.T) {
	r := corejson.NewResult.Many("a", "b", "c")
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I20_NewResult_UsingJsoner_Nil(t *testing.T) {
	r := corejson.NewResult.UsingJsoner(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_NewResult_UsingSerializerFunc_Nil(t *testing.T) {
	r := corejson.NewResult.UsingSerializerFunc(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_NewResult_DeserializeUsingResult_HasIssues(t *testing.T) {
	errResult := corejson.NewResult.ErrorPtr(errors.New("e"))
	r := corejson.NewResult.DeserializeUsingResult(errResult)
	if r.Error == nil {
		t.Fatal("expected error")
	}
}

// ===== CastAny coverage =====

func Test_I20_CastAny_FromToDefault_NilFrom(t *testing.T) {
	var s string
	// FromToDefault(nil, &s) → reflectionCasting returns (err, false) for nil,
	// falls through to Serialize.Apply(nil) → "null" → Unmarshal sets zero value, no error
	err := corejson.CastAny.FromToDefault(nil, &s)
	if err != nil {
		t.Fatal("expected no error — nil serializes to null")
	}
}

func Test_I20_CastAny_FromToReflection(t *testing.T) {
	var s string
	err := corejson.CastAny.FromToReflection(`"hello"`, &s)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I20_CastAny_OrDeserializeTo(t *testing.T) {
	var s string
	err := corejson.CastAny.OrDeserializeTo(`"hello"`, &s)
	if err != nil {
		t.Fatal(err)
	}
}

// ===== Empty creator coverage =====

func Test_I20_Empty_ResultWithErr(t *testing.T) {
	r := corejson.Empty.ResultWithErr("T", errors.New("e"))
	if r.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_Empty_BytesCollection(t *testing.T) {
	bc := corejson.Empty.BytesCollection()
	if bc.Length() != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I20_Empty_BytesCollectionPtr(t *testing.T) {
	bc := corejson.Empty.BytesCollectionPtr()
	if bc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I20_Empty_MapResults(t *testing.T) {
	mr := corejson.Empty.MapResults()
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

// ===== BytesToString / BytesToPrettyString =====

func Test_I20_BytesToString_Empty(t *testing.T) {
	if corejson.BytesToString(nil) != "" {
		t.Fatal("expected empty")
	}
}

func Test_I20_BytesToString_Valid(t *testing.T) {
	if corejson.BytesToString([]byte("hi")) != "hi" {
		t.Fatal("expected 'hi'")
	}
}

func Test_I20_BytesToPrettyString_Empty(t *testing.T) {
	if corejson.BytesToPrettyString(nil) != "" {
		t.Fatal("expected empty")
	}
}

func Test_I20_BytesToPrettyString_Valid(t *testing.T) {
	s := corejson.BytesToPrettyString([]byte(`{"a":"b"}`))
	if s == "" {
		t.Fatal("expected pretty string")
	}
}

// ===== BytesDeepClone / BytesCloneIf =====

func Test_I20_BytesDeepClone_Empty(t *testing.T) {
	b := corejson.BytesDeepClone(nil)
	if len(b) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I20_BytesDeepClone_Valid(t *testing.T) {
	b := corejson.BytesDeepClone([]byte("hi"))
	if string(b) != "hi" {
		t.Fatal("expected 'hi'")
	}
}

func Test_I20_BytesCloneIf_NoClone(t *testing.T) {
	b := corejson.BytesCloneIf(false, []byte("hi"))
	if len(b) != 0 {
		t.Fatal("expected empty (no clone)")
	}
}

func Test_I20_BytesCloneIf_Clone(t *testing.T) {
	b := corejson.BytesCloneIf(true, []byte("hi"))
	if string(b) != "hi" {
		t.Fatal("expected 'hi'")
	}
}
