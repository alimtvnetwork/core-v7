package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ── Migrated from Coverage01_Result_test.go ──

func Test_C01_New(t *testing.T) {
	r := corejson.New("hello")
	if r.HasError() { t.Fatal("expected no error") }
	if r.JsonString() != `"hello"` { t.Fatal("unexpected json") }
	if r.TypeName == "" { t.Fatal("expected type name") }
}

func Test_C01_New_MarshalError(t *testing.T) {
	ch := make(chan int)
	r := corejson.New(ch)
	if !r.HasError() { t.Fatal("expected error for channel") }
}

func Test_C01_NewPtr(t *testing.T) {
	r := corejson.NewPtr(42)
	if r == nil { t.Fatal("expected non-nil") }
	if r.HasError() { t.Fatal("expected no error") }
	if r.JsonString() != "42" { t.Fatal("unexpected json") }
}

func Test_C01_NewPtr_MarshalError(t *testing.T) {
	ch := make(chan int)
	r := corejson.NewPtr(ch)
	if !r.HasError() { t.Fatal("expected error for channel") }
}

func Test_C01_Result_Map(t *testing.T) {
	r := corejson.NewResult.Any("test")
	m := r.Map()
	if _, ok := m["Bytes"]; !ok { t.Fatal("expected bytes key") }

	rErr := corejson.NewResult.Error(errors.New("fail"))
	m2 := rErr.Map()
	if _, ok := m2["Error"]; !ok { t.Fatal("expected error key") }

	var nilR *corejson.Result
	m3 := nilR.Map()
	if len(m3) != 0 { t.Fatal("expected empty map for nil") }
}

func Test_C01_Result_JsonStringPtr(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	s1 := r.JsonStringPtr()
	if s1 == nil || *s1 == "" { t.Fatal("expected non-empty string") }
	s2 := r.JsonStringPtr()
	if *s1 != *s2 { t.Fatal("cache miss") }

	var nilR *corejson.Result
	s3 := nilR.JsonStringPtr()
	if s3 == nil || *s3 != "" { t.Fatal("expected empty string for nil") }

	emptyR := corejson.Result{}
	s4 := emptyR.JsonStringPtr()
	if s4 == nil { t.Fatal("expected non-nil") }
}

func Test_C01_Result_SafeString(t *testing.T) {
	r := corejson.NewResult.Any(123)
	s := r.SafeString()
	if s != "123" { t.Fatal("unexpected safe string") }
}

func Test_C01_Result_PrettyJsonBuffer(t *testing.T) {
	r := corejson.NewResult.Any(map[string]int{"a": 1})
	buf, err := r.PrettyJsonBuffer("", "  ")
	if err != nil { t.Fatal(err) }
	if buf.Len() == 0 { t.Fatal("expected non-empty buffer") }

	emptyR := corejson.Result{}
	buf2, _ := emptyR.PrettyJsonBuffer("", "  ")
	if buf2.Len() != 0 { t.Fatal("expected empty buffer for empty result") }
}

func Test_C01_Result_PrettyJsonString(t *testing.T) {
	r := corejson.NewResult.Any(map[string]int{"a": 1})
	s := r.PrettyJsonString()
	if s == "" { t.Fatal("expected non-empty pretty string") }

	var nilR *corejson.Result
	s2 := nilR.PrettyJsonString()
	if s2 != "" { t.Fatal("expected empty for nil") }

	emptyR := &corejson.Result{}
	s3 := emptyR.PrettyJsonString()
	if s3 != "" { t.Fatal("expected empty for empty") }
}

func Test_C01_Result_PrettyJsonStringOrErrString(t *testing.T) {
	r := corejson.NewResult.Any(42)
	s := r.PrettyJsonStringOrErrString()
	if s == "" { t.Fatal("expected non-empty") }

	rErr := corejson.NewResult.Error(errors.New("boom"))
	s2 := rErr.PrettyJsonStringOrErrString()
	if s2 == "" { t.Fatal("expected error message") }

	var nilR *corejson.Result
	s3 := nilR.PrettyJsonStringOrErrString()
	if s3 == "" { t.Fatal("expected nil message") }
}

func Test_C01_Result_Length(t *testing.T) {
	r := corejson.NewResult.Any("hi")
	if r.Length() == 0 { t.Fatal("expected non-zero length") }
	var nilR *corejson.Result
	if nilR.Length() != 0 { t.Fatal("expected zero for nil") }
}

func Test_C01_Result_HasError_ErrorString(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if r.HasError() { t.Fatal("should not have error") }
	if r.ErrorString() != "" { t.Fatal("expected empty error string") }

	rErr := corejson.NewResult.Error(errors.New("fail"))
	if !rErr.HasError() { t.Fatal("should have error") }
	if rErr.ErrorString() != "fail" { t.Fatal("unexpected error string") }
}

func Test_C01_Result_IsErrorEqual(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if !r.IsErrorEqual(nil) { t.Fatal("both nil should be equal") }

	rErr := corejson.NewResult.Error(errors.New("boom"))
	if rErr.IsErrorEqual(nil) { t.Fatal("error vs nil should not be equal") }
	if !rErr.IsErrorEqual(errors.New("boom")) { t.Fatal("same error message should be equal") }
	if rErr.IsErrorEqual(errors.New("other")) { t.Fatal("different errors should not be equal") }
}

func Test_C01_Result_String(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	_ = r.String() // may return empty depending on IsAnyNull
	rErr := corejson.NewResult.Error(errors.New("err"))
	_ = rErr.String()
}

func Test_C01_Result_SafeNonIssueBytes(t *testing.T) {
	r := corejson.NewResult.Any(42)
	if len(r.SafeNonIssueBytes()) == 0 { t.Fatal("expected non-empty bytes") }
	rErr := corejson.NewResult.Error(errors.New("e"))
	if len(rErr.SafeNonIssueBytes()) != 0 { t.Fatal("expected empty bytes for error result") }
}

func Test_C01_Result_SafeBytes_Values_SafeValues(t *testing.T) {
	r := corejson.NewResult.Any(1)
	if len(r.SafeBytes()) == 0 { t.Fatal("expected bytes") }
	if len(r.Values()) == 0 { t.Fatal("expected values") }
	if len(r.SafeValues()) == 0 { t.Fatal("expected safe values") }
	if len(r.SafeValuesPtr()) == 0 { t.Fatal("expected safe values ptr") }

	var nilR *corejson.Result
	if len(nilR.SafeBytes()) != 0 { t.Fatal("expected empty for nil") }
	if len(nilR.SafeValues()) != 0 { t.Fatal("expected empty for nil") }
}

func Test_C01_Result_Raw(t *testing.T) {
	r := corejson.NewResult.Any("x")
	b, err := r.Raw()
	if err != nil || len(b) == 0 { t.Fatal("expected raw bytes") }
	var nilR *corejson.Result
	_, err2 := nilR.Raw()
	if err2 == nil { t.Fatal("expected error for nil") }
}

func Test_C01_Result_RawMust(t *testing.T) {
	r := corejson.NewResult.Any("y")
	b := r.RawMust()
	if len(b) == 0 { t.Fatal("expected bytes") }
}

func Test_C01_Result_RawString(t *testing.T) {
	r := corejson.NewResult.Any("z")
	s, err := r.RawString()
	if err != nil || s == "" { t.Fatal("expected raw string") }
}

func Test_C01_Result_RawStringMust(t *testing.T) {
	r := corejson.NewResult.Any("a")
	if r.RawStringMust() == "" { t.Fatal("expected string") }
}

func Test_C01_Result_RawErrString(t *testing.T) {
	r := corejson.NewResult.Any("b")
	b, errStr := r.RawErrString()
	if len(b) == 0 { t.Fatal("expected bytes") }
	_ = errStr
}

func Test_C01_Result_RawPrettyString(t *testing.T) {
	r := corejson.NewResult.Any(map[string]int{"k": 1})
	s, err := r.RawPrettyString()
	if err != nil || s == "" { t.Fatal("expected pretty string") }
}

func Test_C01_Result_MeaningfulError(t *testing.T) {
	var nilR *corejson.Result
	if nilR.MeaningfulError() == nil { t.Fatal("expected error for nil") }
	r := corejson.NewResult.Any("good")
	if r.MeaningfulError() != nil { t.Fatal("expected nil error") }
	emptyR := &corejson.Result{}
	if emptyR.MeaningfulError() == nil { t.Fatal("expected error for empty bytes") }
	rErr := corejson.NewResult.Error(errors.New("boom"))
	if rErr.MeaningfulError() == nil { t.Fatal("expected error") }
}

func Test_C01_Result_MeaningfulErrorMessage(t *testing.T) {
	r := corejson.NewResult.Any("ok")
	if r.MeaningfulErrorMessage() != "" { t.Fatal("expected empty") }
	rErr := corejson.NewResult.Error(errors.New("x"))
	if rErr.MeaningfulErrorMessage() == "" { t.Fatal("expected message") }
}

func Test_C01_Result_IsEmptyError(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if !r.IsEmptyError() { t.Fatal("expected empty error") }
	var nilR *corejson.Result
	if !nilR.IsEmptyError() { t.Fatal("expected empty error for nil") }
}

func Test_C01_Result_HasSafeItems(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if !r.HasSafeItems() { t.Fatal("expected safe items") }
	rErr := corejson.NewResult.Error(errors.New("e"))
	if rErr.HasSafeItems() { t.Fatal("should not have safe items") }
}

func Test_C01_Result_IsAnyNull(t *testing.T) {
	var nilR *corejson.Result
	if !nilR.IsAnyNull() { t.Fatal("expected null for nil") }
	r := corejson.Result{}
	if !r.IsAnyNull() { t.Fatal("expected null for empty bytes") }
	r2 := corejson.NewResult.Any(1)
	if r2.IsAnyNull() { t.Fatal("should not be null") }
}

func Test_C01_Result_HasIssuesOrEmpty(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if r.HasIssuesOrEmpty() { t.Fatal("should not have issues") }
	r2 := corejson.NewResult.Error(errors.New("e"))
	if !r2.HasIssuesOrEmpty() { t.Fatal("should have issues") }
}

func Test_C01_Result_IsEmpty_HasAnyItem(t *testing.T) {
	r := corejson.Result{}
	if !r.IsEmpty() { t.Fatal("expected empty") }
	if r.HasAnyItem() { t.Fatal("should not have items") }
	r2 := corejson.NewResult.Any("x")
	if r2.IsEmpty() { t.Fatal("should not be empty") }
	if !r2.HasAnyItem() { t.Fatal("should have items") }
}

func Test_C01_Result_IsEmptyJson_HasJson_HasBytes_HasJsonBytes(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if r.IsEmptyJson() { t.Fatal("should not be empty json") }
	if !r.HasJson() { t.Fatal("should have json") }
	if !r.HasBytes() { t.Fatal("should have bytes") }
	if !r.HasJsonBytes() { t.Fatal("should have json bytes") }
	empty := corejson.Result{Bytes: []byte("{}")}
	if !empty.IsEmptyJsonBytes() { t.Fatal("should be empty json for {}") }
}

func Test_C01_Result_Deserialize_Unmarshal(t *testing.T) {
	r := corejson.NewResult.Any(map[string]string{"k": "v"})
	var out map[string]string
	err := r.Deserialize(&out)
	if err != nil { t.Fatal(err) }
	if out["k"] != "v" { t.Fatal("unexpected value") }
	r.DeserializeMust(&out)

	var nilR *corejson.Result
	err2 := nilR.Unmarshal(&out)
	if err2 == nil { t.Fatal("expected error for nil") }

	rErr := corejson.NewResult.Error(errors.New("e"))
	err3 := rErr.Unmarshal(&out)
	if err3 == nil { t.Fatal("expected error") }
}

func Test_C01_Result_UnmarshalSkipExistingIssues(t *testing.T) {
	rErr := corejson.NewResult.Error(errors.New("e"))
	var out string
	err := rErr.UnmarshalSkipExistingIssues(&out)
	if err != nil { t.Fatal("should skip and return nil") }
	r := corejson.NewResult.Any("hello")
	err2 := r.UnmarshalSkipExistingIssues(&out)
	if err2 != nil { t.Fatal(err2) }
	if out != "hello" { t.Fatal("unexpected value") }
}

func Test_C01_Result_Serialize(t *testing.T) {
	r := corejson.NewResult.Any(42)
	b, err := r.Serialize()
	if err != nil || len(b) == 0 { t.Fatal("expected serialized bytes") }
	var nilR *corejson.Result
	_, err2 := nilR.Serialize()
	if err2 == nil { t.Fatal("expected error for nil") }
	rErr := corejson.NewResult.Error(errors.New("e"))
	_, err3 := rErr.Serialize()
	if err3 == nil { t.Fatal("expected error") }
}

func Test_C01_Result_SerializeSkipExistingIssues(t *testing.T) {
	rErr := corejson.NewResult.Error(errors.New("e"))
	b, err := rErr.SerializeSkipExistingIssues()
	if b != nil || err != nil { t.Fatal("should return nil,nil for issues") }
	r := corejson.NewResult.Any(42)
	b2, err2 := r.SerializeSkipExistingIssues()
	if err2 != nil || len(b2) == 0 { t.Fatal("expected bytes") }
}

func Test_C01_Result_SerializeMust(t *testing.T) {
	r := corejson.NewResult.Any(42)
	if len(r.SerializeMust()) == 0 { t.Fatal("expected bytes") }
}

func Test_C01_Result_UnmarshalResult(t *testing.T) {
	r := corejson.NewResult.Any(corejson.Result{Bytes: []byte(`"x"`), TypeName: "test"})
	_, _ = r.UnmarshalResult()
}

func Test_C01_Result_JsonModel_JsonModelAny(t *testing.T) {
	r := corejson.NewResult.Any("x")
	_ = r.JsonModel()
	var nilR *corejson.Result
	m2 := nilR.JsonModel()
	if m2.Error == nil { t.Fatal("expected error for nil") }
	r2 := corejson.NewResult.Any("y")
	_ = r2.JsonModelAny()
}

func Test_C01_Result_Json_JsonPtr(t *testing.T) {
	r := corejson.NewResult.Any("x")
	j := r.Json()
	if j.HasError() { t.Fatal("expected no error") }
	if r.JsonPtr() == nil { t.Fatal("expected non-nil") }
}

func Test_C01_Result_ParseInjectUsingJson(t *testing.T) {
	r := corejson.NewResult.Any("x")
	target := &corejson.Result{}
	_, _ = target.ParseInjectUsingJson(r.Ptr())
}

func Test_C01_Result_CloneError(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if r.CloneError() != nil { t.Fatal("should be nil") }
	rErr := corejson.NewResult.Error(errors.New("e"))
	if rErr.CloneError() == nil { t.Fatal("should have error") }
}

func Test_C01_Result_Ptr_NonPtr_ToPtr_ToNonPtr(t *testing.T) {
	r := corejson.NewResult.Any("x")
	p := r.Ptr()
	if p == nil { t.Fatal("expected ptr") }
	_ = p.NonPtr()
	var nilR *corejson.Result
	np2 := nilR.NonPtr()
	if np2.Error == nil { t.Fatal("expected error") }
	r2 := corejson.NewResult.Any("y")
	_ = r2.ToPtr()
	_ = r2.ToNonPtr()
}

func Test_C01_Result_IsEqualPtr(t *testing.T) {
	r1 := corejson.NewResult.AnyPtr("x")
	r2 := corejson.NewResult.AnyPtr("x")
	if !r1.IsEqualPtr(r2) { t.Fatal("should be equal") }
	var nilR *corejson.Result
	if !nilR.IsEqualPtr(nil) { t.Fatal("both nil should be equal") }
	if nilR.IsEqualPtr(r1) { t.Fatal("nil vs non-nil should not be equal") }
	if r1.IsEqualPtr(nil) { t.Fatal("non-nil vs nil should not be equal") }
	r3 := corejson.NewResult.AnyPtr("y")
	if r1.IsEqualPtr(r3) { t.Fatal("different should not be equal") }
}

func Test_C01_Result_IsEqual(t *testing.T) {
	r1 := corejson.NewResult.Any("x")
	r2 := corejson.NewResult.Any("x")
	if !r1.IsEqual(r2) { t.Fatal("should be equal") }
	r3 := corejson.NewResult.Any("y")
	if r1.IsEqual(r3) { t.Fatal("should not be equal") }
}

func Test_C01_Result_CombineErrorWithRefString(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if r.CombineErrorWithRefString("ref1") != "" { t.Fatal("expected empty for no error") }
	rErr := corejson.NewResult.Error(errors.New("e"))
	if rErr.CombineErrorWithRefString("ref1") == "" { t.Fatal("expected combined string") }
}

func Test_C01_Result_CombineErrorWithRefError(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if r.CombineErrorWithRefError("ref") != nil { t.Fatal("expected nil") }
	rErr := corejson.NewResult.Error(errors.New("e"))
	if rErr.CombineErrorWithRefError("ref") == nil { t.Fatal("expected error") }
}

func Test_C01_Result_BytesError(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if r.BytesError() == nil { t.Fatal("expected non-nil") }
	var nilR *corejson.Result
	if nilR.BytesError() != nil { t.Fatal("expected nil") }
}

func Test_C01_Result_Dispose(t *testing.T) {
	r := corejson.NewResult.Any("x")
	r.Dispose()
	if r.Bytes != nil { t.Fatal("expected nil bytes after dispose") }
	var nilR *corejson.Result
	nilR.Dispose()
}

func Test_C01_Result_Clone(t *testing.T) {
	r := corejson.NewResult.Any("test")
	c := r.Clone(false)
	if c.JsonString() != r.JsonString() { t.Fatal("shallow clone mismatch") }
	c2 := r.Clone(true)
	if c2.JsonString() != r.JsonString() { t.Fatal("deep clone mismatch") }
	empty := corejson.Result{}
	_ = empty.Clone(true)
}

func Test_C01_Result_CloneIf(t *testing.T) {
	r := corejson.NewResult.Any("x")
	_ = r.CloneIf(true, false)
	_ = r.CloneIf(false, false)
}

func Test_C01_Result_ClonePtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	c := r.ClonePtr(true)
	if c == nil { t.Fatal("expected non-nil clone") }
	var nilR *corejson.Result
	if nilR.ClonePtr(true) != nil { t.Fatal("expected nil for nil") }
}

func Test_C01_Result_InjectInto(t *testing.T) {
	r := corejson.NewResult.Any("x")
	target := &corejson.Result{}
	_ = r.InjectInto(target)
}

func Test_C01_Result_AsJsonContractsBinder(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if r.AsJsonContractsBinder() == nil { t.Fatal("expected non-nil") }
}

func Test_C01_Result_AsJsoner(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if r.AsJsoner() == nil { t.Fatal("expected non-nil") }
}

func Test_C01_Result_JsonParseSelfInject(t *testing.T) {
	r := corejson.NewResult.Any("x")
	target := corejson.NewResult.Any("y")
	_ = target.JsonParseSelfInject(r.Ptr())
}

func Test_C01_Result_AsJsonParseSelfInjector(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if r.AsJsonParseSelfInjector() == nil { t.Fatal("expected non-nil") }
}

func Test_C01_Result_DeserializedFieldsToMap(t *testing.T) {
	r := corejson.NewResult.Any(map[string]int{"a": 1})
	_, _ = r.DeserializedFieldsToMap()
	var nilR *corejson.Result
	fm, err := nilR.DeserializedFieldsToMap()
	if err != nil || len(fm) != 0 { t.Fatal("expected empty for nil") }
}

func Test_C01_Result_SafeDeserializedFieldsToMap(t *testing.T) {
	r := corejson.NewResult.Any(map[string]int{"a": 1})
	_ = r.SafeDeserializedFieldsToMap()
}

func Test_C01_Result_FieldsNames(t *testing.T) {
	r := corejson.NewResult.Any(map[string]int{"a": 1})
	_, _ = r.FieldsNames()
}

func Test_C01_Result_SafeFieldsNames(t *testing.T) {
	r := corejson.NewResult.Any(map[string]int{"a": 1})
	_ = r.SafeFieldsNames()
}

func Test_C01_Result_BytesTypeName(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if r.BytesTypeName() == "" { t.Fatal("expected type name") }
	var nilR *corejson.Result
	if nilR.BytesTypeName() != "" { t.Fatal("expected empty for nil") }
}

func Test_C01_Result_SafeBytesTypeName(t *testing.T) {
	r := corejson.NewResult.Any("x")
	if r.SafeBytesTypeName() == "" { t.Fatal("expected type name") }
	emptyR := &corejson.Result{}
	if emptyR.SafeBytesTypeName() != "" { t.Fatal("expected empty for empty result") }
}

func Test_C01_Result_HandleError_NoPanic(t *testing.T) {
	r := corejson.NewResult.Any("x")
	r.HandleError()
}

func Test_C01_Result_MustBeSafe_NoPanic(t *testing.T) {
	r := corejson.NewResult.Any("x")
	r.MustBeSafe()
}

func Test_C01_Result_UnmarshalMust(t *testing.T) {
	r := corejson.NewResult.Any(42)
	var i int
	r.UnmarshalMust(&i)
	if i != 42 { t.Fatal("unexpected") }
}
