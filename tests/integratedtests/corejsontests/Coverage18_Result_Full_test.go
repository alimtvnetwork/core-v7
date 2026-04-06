package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ══════════════════════════════════════════════════════════════════════════════
// Result.go — comprehensive coverage for ALL uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov18_Result_Map_WithAllFields(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hello"`), Error: errors.New("e"), TypeName: "T"}
	m := r.Map()
	if len(m) != 3 {
		t.Fatal("expected 3 fields in map")
	}
}

func Test_Cov18_Result_Map_Nil(t *testing.T) {
	var r *corejson.Result
	m := r.Map()
	if len(m) != 0 {
		t.Fatal("expected empty map for nil")
	}
}

func Test_Cov18_Result_Map_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	m := r.Map()
	if _, ok := m["Error"]; ok {
		t.Fatal("should not have error key")
	}
}

func Test_Cov18_Result_SafeDeserializedFieldsToMap(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"k": "v"})
	fm := r.SafeDeserializedFieldsToMap()
	_ = fm
}

func Test_Cov18_Result_SafeDeserializedFieldsToMap_Nil(t *testing.T) {
	var r *corejson.Result
	fm := r.SafeDeserializedFieldsToMap()
	if len(fm) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_Cov18_Result_FieldsNames(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"a": "1", "b": "2"})
	names, err := r.FieldsNames()
	_ = names
	_ = err
}

func Test_Cov18_Result_SafeFieldsNames(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"a": "1"})
	names := r.SafeFieldsNames()
	_ = names
}

func Test_Cov18_Result_BytesTypeName(t *testing.T) {
	r := &corejson.Result{TypeName: "MyType"}
	if r.BytesTypeName() != "MyType" {
		t.Fatal("wrong type name")
	}
}

func Test_Cov18_Result_BytesTypeName_Nil(t *testing.T) {
	var r *corejson.Result
	if r.BytesTypeName() != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_Cov18_Result_SafeBytesTypeName(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	name := r.SafeBytesTypeName()
	if name == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov18_Result_SafeBytesTypeName_Empty(t *testing.T) {
	r := &corejson.Result{}
	name := r.SafeBytesTypeName()
	if name != "" {
		t.Fatal("expected empty for empty result")
	}
}

func Test_Cov18_Result_SafeString(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s := r.SafeString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov18_Result_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	var r *corejson.Result
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected message for nil")
	}
}

func Test_Cov18_Result_PrettyJsonStringOrErrString_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected error message")
	}
}

func Test_Cov18_Result_PrettyJsonStringOrErrString_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]int{"a": 1})
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected pretty json")
	}
}

func Test_Cov18_Result_String_WithError(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e"), TypeName: "T"}
	s := r.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov18_Result_String_NoError(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	s := r.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov18_Result_SafeNonIssueBytes(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	b := r.SafeNonIssueBytes()
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_Cov18_Result_SafeNonIssueBytes_Empty(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	b := r.SafeNonIssueBytes()
	if len(b) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_Cov18_Result_Values(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	if len(r.Values()) == 0 {
		t.Fatal("expected values")
	}
}

func Test_Cov18_Result_SafeValues(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	if len(r.SafeValues()) == 0 {
		t.Fatal("expected values")
	}
}

func Test_Cov18_Result_SafeValues_Nil(t *testing.T) {
	var r *corejson.Result
	if len(r.SafeValues()) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_Cov18_Result_SafeValuesPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	if len(r.SafeValuesPtr()) == 0 {
		t.Fatal("expected values")
	}
}

func Test_Cov18_Result_SafeValuesPtr_Issues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	if len(r.SafeValuesPtr()) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_Cov18_Result_RawMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	b := r.RawMust()
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_Cov18_Result_RawString(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s, err := r.RawString()
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func Test_Cov18_Result_RawStringMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s := r.RawStringMust()
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_Cov18_Result_RawErrString(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	b, errMsg := r.RawErrString()
	if len(b) == 0 || errMsg != "" {
		t.Fatal("unexpected")
	}
}

func Test_Cov18_Result_RawPrettyString(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]int{"a": 1})
	s, err := r.RawPrettyString()
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func Test_Cov18_Result_MeaningfulErrorMessage_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	msg := r.MeaningfulErrorMessage()
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_Cov18_Result_MeaningfulErrorMessage_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	msg := r.MeaningfulErrorMessage()
	if msg == "" {
		t.Fatal("expected message")
	}
}

func Test_Cov18_Result_HasSafeItems(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	if !r.HasSafeItems() {
		t.Fatal("expected true")
	}
}

func Test_Cov18_Result_HasSafeItems_Empty(t *testing.T) {
	r := &corejson.Result{}
	if r.HasSafeItems() {
		t.Fatal("expected false")
	}
}

func Test_Cov18_Result_HasJsonBytes(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	if !r.HasJsonBytes() {
		t.Fatal("expected true")
	}
}

func Test_Cov18_Result_HasAnyItem(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	if !r.HasAnyItem() {
		t.Fatal("expected true")
	}
}

func Test_Cov18_Result_HasJson(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	if !r.HasJson() {
		t.Fatal("expected true")
	}
}

func Test_Cov18_Result_DeserializeMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	var s string
	r.DeserializeMust(&s)
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_Cov18_Result_UnmarshalMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(42)
	var n int
	r.UnmarshalMust(&n)
	if n != 42 {
		t.Fatal("unexpected")
	}
}

func Test_Cov18_Result_SerializeSkipExistingIssues_HasIssues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	b, err := r.SerializeSkipExistingIssues()
	if b != nil || err != nil {
		t.Fatal("expected nil,nil for issues")
	}
}

func Test_Cov18_Result_SerializeSkipExistingIssues_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	b, err := r.SerializeSkipExistingIssues()
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func Test_Cov18_Result_Serialize(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	b, err := r.Serialize()
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func Test_Cov18_Result_Serialize_Nil(t *testing.T) {
	var r *corejson.Result
	_, err := r.Serialize()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov18_Result_Serialize_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	_, err := r.Serialize()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov18_Result_SerializeMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	b := r.SerializeMust()
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_Cov18_Result_UnmarshalSkipExistingIssues_HasIssues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	if err != nil {
		t.Fatal("expected nil for issues")
	}
}

func Test_Cov18_Result_UnmarshalResult(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	r, err := jr.UnmarshalResult()
	_ = r
	_ = err
}

func Test_Cov18_Result_JsonModel(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	m := r.JsonModel()
	if m.TypeName == "" {
		t.Fatal("expected type name")
	}
}

func Test_Cov18_Result_JsonModel_Nil(t *testing.T) {
	var r *corejson.Result
	m := r.JsonModel()
	if m.Error == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_Cov18_Result_JsonModelAny(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	a := r.JsonModelAny()
	if a == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov18_Result_Json(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	j := r.Json()
	if j.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_Cov18_Result_JsonPtr(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	j := r.JsonPtr()
	if j.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_Cov18_Result_ParseInjectUsingJson(t *testing.T) {
	inner := corejson.NewResult.Any(corejson.Result{Bytes: []byte(`"hi"`), TypeName: "t"})
	target := &corejson.Result{}
	_, err := target.ParseInjectUsingJson(inner.Ptr())
	_ = err
}

func Test_Cov18_Result_CloneError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	ce := r.CloneError()
	if ce == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov18_Result_CloneError_Nil(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	ce := r.CloneError()
	if ce != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov18_Result_Ptr(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	p := r.Ptr()
	if p == nil {
		t.Fatal("expected ptr")
	}
}

func Test_Cov18_Result_NonPtr_Nil(t *testing.T) {
	var r *corejson.Result
	np := r.NonPtr()
	if np.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov18_Result_ToPtr(t *testing.T) {
	r := corejson.NewResult.Any("x")
	p := r.ToPtr()
	if p == nil {
		t.Fatal("expected ptr")
	}
}

func Test_Cov18_Result_ToNonPtr(t *testing.T) {
	r := corejson.NewResult.Any("x")
	np := r.ToNonPtr()
	_ = np
}

func Test_Cov18_Result_IsEqualPtr(t *testing.T) {
	a := corejson.NewResult.AnyPtr("hello")
	b := corejson.NewResult.AnyPtr("hello")
	if !a.IsEqualPtr(b) {
		t.Fatal("expected equal")
	}
}

func Test_Cov18_Result_IsEqualPtr_BothNil(t *testing.T) {
	var a, b *corejson.Result
	if !a.IsEqualPtr(b) {
		t.Fatal("expected equal for nil")
	}
}

func Test_Cov18_Result_IsEqualPtr_OneNil(t *testing.T) {
	a := corejson.NewResult.AnyPtr("x")
	if a.IsEqualPtr(nil) {
		t.Fatal("expected not equal")
	}
}

func Test_Cov18_Result_IsEqualPtr_Same(t *testing.T) {
	a := corejson.NewResult.AnyPtr("x")
	if !a.IsEqualPtr(a) {
		t.Fatal("expected equal for same ptr")
	}
}

func Test_Cov18_Result_IsEqualPtr_DiffLen(t *testing.T) {
	a := corejson.NewResult.AnyPtr("hello")
	b := corejson.NewResult.AnyPtr("hi")
	if a.IsEqualPtr(b) {
		t.Fatal("expected not equal")
	}
}

func Test_Cov18_Result_IsEqualPtr_DiffError(t *testing.T) {
	a := &corejson.Result{Bytes: []byte("x"), Error: errors.New("a")}
	b := &corejson.Result{Bytes: []byte("x"), Error: errors.New("b")}
	if a.IsEqualPtr(b) {
		t.Fatal("expected not equal")
	}
}

func Test_Cov18_Result_IsEqualPtr_DiffType(t *testing.T) {
	a := &corejson.Result{Bytes: []byte("x"), TypeName: "A"}
	b := &corejson.Result{Bytes: []byte("x"), TypeName: "B"}
	if a.IsEqualPtr(b) {
		t.Fatal("expected not equal")
	}
}

func Test_Cov18_Result_CombineErrorWithRefString(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	s := r.CombineErrorWithRefString("ref1", "ref2")
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_Cov18_Result_CombineErrorWithRefString_NoErr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	s := r.CombineErrorWithRefString("ref")
	if s != "" {
		t.Fatal("expected empty")
	}
}

func Test_Cov18_Result_CombineErrorWithRefError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	err := r.CombineErrorWithRefError("ref")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov18_Result_CombineErrorWithRefError_NoErr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	err := r.CombineErrorWithRefError("ref")
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov18_Result_IsEqual(t *testing.T) {
	a := corejson.NewResult.Any("hello")
	b := corejson.NewResult.Any("hello")
	if !a.IsEqual(b) {
		t.Fatal("expected equal")
	}
}

func Test_Cov18_Result_IsEqual_DiffLen(t *testing.T) {
	a := corejson.NewResult.Any("hello")
	b := corejson.NewResult.Any("hi")
	if a.IsEqual(b) {
		t.Fatal("expected not equal")
	}
}

func Test_Cov18_Result_BytesError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	be := r.BytesError()
	if be == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov18_Result_BytesError_Nil(t *testing.T) {
	var r *corejson.Result
	be := r.BytesError()
	if be != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov18_Result_Dispose(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	r.Dispose()
	if r.Bytes != nil || r.Error != nil {
		t.Fatal("expected disposed")
	}
}

func Test_Cov18_Result_Dispose_Nil(t *testing.T) {
	var r *corejson.Result
	r.Dispose() // should not panic
}

func Test_Cov18_Result_CloneIf(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	c1 := r.CloneIf(true, true)
	c2 := r.CloneIf(false, false)
	if c1.HasError() || c2.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_Cov18_Result_ClonePtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	c := r.ClonePtr(true)
	if c == nil {
		t.Fatal("expected clone")
	}
}

func Test_Cov18_Result_ClonePtr_Nil(t *testing.T) {
	var r *corejson.Result
	c := r.ClonePtr(true)
	if c != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov18_Result_Clone_DeepAndShallow(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	deep := r.Clone(true)
	shallow := r.Clone(false)
	if deep.HasError() || shallow.HasError() {
		t.Fatal("unexpected")
	}
}

func Test_Cov18_Result_Clone_Empty(t *testing.T) {
	r := corejson.NewResult.Any("")
	c := r.Clone(true)
	_ = c
}

func Test_Cov18_Result_AsJsonContractsBinder(t *testing.T) {
	r := corejson.NewResult.Any("x")
	b := r.AsJsonContractsBinder()
	if b == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov18_Result_AsJsoner(t *testing.T) {
	r := corejson.NewResult.Any("x")
	j := r.AsJsoner()
	if j == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov18_Result_JsonParseSelfInject(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	inner := corejson.NewResult.AnyPtr(corejson.NewResult.Any("world"))
	err := r.JsonParseSelfInject(inner)
	_ = err
}

func Test_Cov18_Result_AsJsonParseSelfInjector(t *testing.T) {
	r := corejson.NewResult.Any("x")
	inj := r.AsJsonParseSelfInjector()
	if inj == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov18_Result_InjectInto(t *testing.T) {
	r := corejson.NewResult.AnyPtr([]string{"a", "b"})
	target := corejson.NewResult.Any("x")
	err := r.InjectInto(&target)
	_ = err
}

// ── IsErrorEqual branches ──

func Test_Cov18_Result_IsErrorEqual_BothNil(t *testing.T) {
	r := &corejson.Result{}
	if !r.IsErrorEqual(nil) {
		t.Fatal("expected true both nil")
	}
}

func Test_Cov18_Result_IsErrorEqual_OneNil(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	if r.IsErrorEqual(nil) {
		t.Fatal("expected false one nil")
	}
}

func Test_Cov18_Result_IsErrorEqual_LeftNil(t *testing.T) {
	r := &corejson.Result{}
	if r.IsErrorEqual(errors.New("e")) {
		t.Fatal("expected false left nil")
	}
}

func Test_Cov18_Result_IsErrorEqual_Same(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	if !r.IsErrorEqual(errors.New("e")) {
		t.Fatal("expected true same msg")
	}
}

func Test_Cov18_Result_IsErrorEqual_Different(t *testing.T) {
	r := &corejson.Result{Error: errors.New("a")}
	if r.IsErrorEqual(errors.New("b")) {
		t.Fatal("expected false different")
	}
}
