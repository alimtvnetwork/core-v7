package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===== Result methods coverage =====

func Test_I20_Result_CloneIf_True(t *testing.T) {
	r := corejson.NewResult.Any(map[string]string{"a": "1"})
	cloned := r.CloneIf(true, true)
	actual := args.Map{"result": cloned.Length() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned bytes", actual)
}

func Test_I20_Result_CloneIf_False(t *testing.T) {
	r := corejson.NewResult.Any(map[string]string{"a": "1"})
	same := r.CloneIf(false, false)
	actual := args.Map{"result": same.Length() != r.Length()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same result", actual)
}

func Test_I20_Result_Clone_DeepClone(t *testing.T) {
	r := corejson.NewResult.Any("test")
	c := r.Clone(true)
	actual := args.Map{"result": c.Length() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Result_Clone_ShallowClone(t *testing.T) {
	r := corejson.NewResult.Any("test")
	c := r.Clone(false)
	actual := args.Map{"result": c.Length() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Result_Clone_Empty(t *testing.T) {
	r := corejson.NewResult.Empty()
	c := r.Clone(true)
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_Result_ClonePtr_Nil(t *testing.T) {
	var r *corejson.Result
	c := r.ClonePtr(true)
	actual := args.Map{"result": c != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Result_ClonePtr_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("test")
	c := r.ClonePtr(true)
	actual := args.Map{"result": c == nil || c.Length() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned ptr", actual)
}

func Test_I20_Result_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	var r *corejson.Result
	s := r.PrettyJsonStringOrErrString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty nil message", actual)
}

func Test_I20_Result_PrettyJsonStringOrErrString_HasError(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("test-err"))
	s := r.PrettyJsonStringOrErrString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error string", actual)
}

func Test_I20_Result_PrettyJsonStringOrErrString_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s := r.PrettyJsonStringOrErrString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty json", actual)
}

func Test_I20_Result_HandleErrorWithMsg_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	r.HandleErrorWithMsg("no-op") // Should not panic
}

func Test_I20_Result_HandleErrorWithMsg_Panic(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("bad"))
	defer func() {
		actual := args.Map{"result": rec := recover(); rec == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	r.HandleErrorWithMsg("context message")
}

func Test_I20_Result_DeserializeMust_Success(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	var s string
	r.DeserializeMust(&s)
	actual := args.Map{"result": s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello', got ''", actual)
}

func Test_I20_Result_DeserializeMust_Panic(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("err"))
	defer func() {
		actual := args.Map{"result": rec := recover(); rec == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	var s string
	r.DeserializeMust(&s)
}

func Test_I20_Result_UnmarshalMust_Success(t *testing.T) {
	r := corejson.NewResult.AnyPtr(42)
	var i int
	r.UnmarshalMust(&i)
	actual := args.Map{"result": i != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
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
	actual := args.Map{"result": be == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_I20_Result_BytesError_Nil(t *testing.T) {
	var r *corejson.Result
	be := r.BytesError()
	actual := args.Map{"result": be != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Result_Dispose(t *testing.T) {
	r := corejson.NewResult.AnyPtr("test")
	r.Dispose()
	actual := args.Map{"result": r.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected disposed", actual)
}

func Test_I20_Result_Dispose_Nil(t *testing.T) {
	var r *corejson.Result
	r.Dispose() // should not panic
}

func Test_I20_Result_NonPtr_Nil(t *testing.T) {
	var r *corejson.Result
	nr := r.NonPtr()
	actual := args.Map{"result": nr.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error in nonptr of nil", actual)
}

func Test_I20_Result_NonPtr_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("test")
	nr := r.NonPtr()
	actual := args.Map{"result": nr.Length() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Result_CombineErrorWithRefError_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	err := r.CombineErrorWithRefError("ref")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Result_CombineErrorWithRefError_HasError(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	err := r.CombineErrorWithRefError("ref1", "ref2")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_Result_CombineErrorWithRefString_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	s := r.CombineErrorWithRefString("ref")
	actual := args.Map{"result": s != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_Result_CloneError_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	actual := args.Map{"result": r.CloneError() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Result_CloneError_HasError(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	actual := args.Map{"result": r.CloneError() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_Result_Ptr_ToPtr(t *testing.T) {
	r := corejson.NewResult.Any("test")
	p := r.Ptr()
	actual := args.Map{"result": p == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ptr", actual)
	np := r.ToPtr()
	actual := args.Map{"result": np == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ptr", actual)
	np2 := r.ToNonPtr()
	actual := args.Map{"result": np2.Length() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Result_IsEqualPtr_BothNil(t *testing.T) {
	var a, b *corejson.Result
	actual := args.Map{"result": a.IsEqualPtr(b)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_I20_Result_IsEqualPtr_OneNil(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	actual := args.Map{"result": r.IsEqualPtr(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_I20_Result_IsEqualPtr_Same(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	actual := args.Map{"result": r.IsEqualPtr(r)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal (same ptr)", actual)
}

func Test_I20_Result_IsEqualPtr_DiffLength(t *testing.T) {
	a := corejson.NewResult.AnyPtr("x")
	b := corejson.NewResult.AnyPtr("xy")
	actual := args.Map{"result": a.IsEqualPtr(b)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_I20_Result_IsEqualPtr_DiffError(t *testing.T) {
	a := corejson.NewResult.Ptr([]byte("x"), errors.New("e1"), "t")
	b := corejson.NewResult.Ptr([]byte("x"), errors.New("e2"), "t")
	actual := args.Map{"result": a.IsEqualPtr(b)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_I20_Result_IsEqualPtr_DiffType(t *testing.T) {
	a := corejson.NewResult.Ptr([]byte(`"x"`), nil, "typeA")
	b := corejson.NewResult.Ptr([]byte(`"x"`), nil, "typeB")
	actual := args.Map{"result": a.IsEqualPtr(b)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_I20_Result_IsEqual(t *testing.T) {
	a := corejson.NewResult.Any("hello")
	b := corejson.NewResult.Any("hello")
	actual := args.Map{"result": a.IsEqual(b)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_I20_Result_IsErrorEqual(t *testing.T) {
	a := corejson.NewResult.ErrorPtr(errors.New("same"))
	actual := args.Map{"result": a.IsErrorEqual(errors.New("same"))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_I20_Result_IsErrorEqual_BothNil(t *testing.T) {
	a := corejson.NewResult.AnyPtr("ok")
	actual := args.Map{"result": a.IsErrorEqual(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal (both nil)", actual)
}

func Test_I20_Result_IsErrorEqual_OnlyOneNil(t *testing.T) {
	a := corejson.NewResult.AnyPtr("ok")
	actual := args.Map{"result": a.IsErrorEqual(errors.New("e"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

// ===== Result serialization coverage =====

func Test_I20_Result_Serialize_Nil(t *testing.T) {
	var r *corejson.Result
	_, err := r.Serialize()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_Result_Serialize_HasError(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	_, err := r.Serialize()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_Result_Serialize_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	b, err := r.Serialize()
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Result_SerializeMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	b := r.SerializeMust()
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Result_SerializeSkipExistingIssues_HasIssues(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	b, err := r.SerializeSkipExistingIssues()
	actual := args.Map{"result": b != nil || err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil,nil for issues", actual)
}

func Test_I20_Result_SerializeSkipExistingIssues_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	b, err := r.SerializeSkipExistingIssues()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_I20_Result_UnmarshalSkipExistingIssues_HasIssues(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for issues", actual)
}

func Test_I20_Result_UnmarshalSkipExistingIssues_Valid(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello', got ''", actual)
}

func Test_I20_Result_UnmarshalSkipExistingIssues_Error(t *testing.T) {
	r := corejson.NewResult.UsingBytesTypePtr([]byte("not-json"), "test")
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error", actual)
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
	actual := args.Map{"result": m.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error in model", actual)
}

func Test_I20_Result_JsonModelAny(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	a := r.JsonModelAny()
	actual := args.Map{"result": a == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_I20_Result_Json_JsonPtr(t *testing.T) {
	r := corejson.NewResult.Any("test")
	j := r.Json()
	actual := args.Map{"result": j.Length() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected json", actual)
	jp := r.JsonPtr()
	actual := args.Map{"result": jp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ptr", actual)
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
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	target.ParseInjectUsingJsonMust(source)
}

func Test_I20_Result_AsJsonContractsBinder(t *testing.T) {
	r := corejson.NewResult.Any("test")
	b := r.AsJsonContractsBinder()
	actual := args.Map{"result": b == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_I20_Result_AsJsoner(t *testing.T) {
	r := corejson.NewResult.Any("test")
	j := r.AsJsoner()
	actual := args.Map{"result": j == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_I20_Result_AsJsonParseSelfInjector(t *testing.T) {
	r := corejson.NewResult.Any("test")
	inj := r.AsJsonParseSelfInjector()
	actual := args.Map{"result": inj == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
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
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Result_RawStringMust_Success(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	s := r.RawStringMust()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_Result_RawStringMust_Panic(t *testing.T) {
	var r *corejson.Result
	defer func() {
		actual := args.Map{"result": rec := recover(); rec == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	r.RawStringMust()
}

func Test_I20_Result_RawErrString(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	b, msg := r.RawErrString()
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	_ = msg
}

func Test_I20_Result_RawPrettyString(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	s, err := r.RawPrettyString()
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_Result_HandleError_Panic(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	defer func() {
		actual := args.Map{"result": rec := recover(); rec == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	r.HandleError()
}

func Test_I20_Result_MustBeSafe_Panic(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	defer func() {
		actual := args.Map{"result": rec := recover(); rec == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	r.MustBeSafe()
}

func Test_I20_Result_SafeNonIssueBytes(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	actual := args.Map{"result": len(r.SafeNonIssueBytes()) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Result_SafeNonIssueBytes_HasIssues(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	actual := args.Map{"result": len(r.SafeNonIssueBytes()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_Result_SafeValuesPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	actual := args.Map{"result": len(r.SafeValuesPtr()) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Result_SafeValuesPtr_HasIssues(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	actual := args.Map{"result": len(r.SafeValuesPtr()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_Result_Values(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	actual := args.Map{"result": len(r.Values()) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Result_Raw_Nil(t *testing.T) {
	var r *corejson.Result
	_, err := r.Raw()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_Result_MeaningfulError_EmptyBytes(t *testing.T) {
	r := &corejson.Result{Bytes: nil, TypeName: "Test"}
	err := r.MeaningfulError()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for empty bytes", actual)
}

func Test_I20_Result_MeaningfulError_HasErrorAndBytes(t *testing.T) {
	r := &corejson.Result{
		Bytes:    []byte(`"test"`),
		Error:    errors.New("some err"),
		TypeName: "Test",
	}
	err := r.MeaningfulError()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_Result_String_HasError(t *testing.T) {
	r := corejson.NewResult.Create([]byte(`"x"`), errors.New("e"), "T")
	s := r.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string with error", actual)
}

func Test_I20_Result_SafeBytesTypeName_Empty(t *testing.T) {
	r := corejson.NewResult.EmptyPtr()
	actual := args.Map{"result": r.SafeBytesTypeName() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_Result_BytesTypeName_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"result": r.BytesTypeName() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ===== anyTo coverage =====

func Test_I20_AnyTo_SerializedRaw(t *testing.T) {
	b, err := corejson.AnyTo.SerializedRaw("hello")
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_AnyTo_SerializedString(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString("hello")
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_SerializedSafeString(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString("hello")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_SerializedStringMust(t *testing.T) {
	s := corejson.AnyTo.SerializedStringMust("hello")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_SafeJsonString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonString("hello")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_PrettyStringWithError_String(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError("hello")
	actual := args.Map{"result": err != nil || s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected passthrough", actual)
}

func Test_I20_AnyTo_PrettyStringWithError_Bytes(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":"b"}`))
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty string", actual)
}

func Test_I20_AnyTo_PrettyStringWithError_Result(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_PrettyStringWithError_ResultPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_PrettyStringWithError_ResultWithError(t *testing.T) {
	r := corejson.NewResult.Create([]byte(`"x"`), errors.New("e"), "T")
	_, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_AnyTo_PrettyStringWithError_ResultPtrWithError(t *testing.T) {
	r := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("e"), "T")
	_, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_AnyTo_PrettyStringWithError_AnyItem(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError(42)
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_SafeJsonPrettyString_String(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString("hi")
	actual := args.Map{"result": s != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected passthrough", actual)
}

func Test_I20_AnyTo_SafeJsonPrettyString_Bytes(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":"b"}`))
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_SafeJsonPrettyString_Result(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	s := corejson.AnyTo.SafeJsonPrettyString(r)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_SafeJsonPrettyString_ResultPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s := corejson.AnyTo.SafeJsonPrettyString(r)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_SafeJsonPrettyString_Any(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString(42)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_JsonString_String(t *testing.T) {
	s := corejson.AnyTo.JsonString("hi")
	actual := args.Map{"result": s != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected passthrough", actual)
}

func Test_I20_AnyTo_JsonString_Bytes(t *testing.T) {
	s := corejson.AnyTo.JsonString([]byte(`{"a":"b"}`))
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_JsonString_Result(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	s := corejson.AnyTo.JsonString(r)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_JsonString_ResultPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s := corejson.AnyTo.JsonString(r)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_JsonString_Any(t *testing.T) {
	s := corejson.AnyTo.JsonString(42)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_JsonStringWithErr_String(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr("hi")
	actual := args.Map{"result": err != nil || s != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected passthrough", actual)
}

func Test_I20_AnyTo_JsonStringWithErr_Bytes(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr([]byte(`"x"`))
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_JsonStringWithErr_Result_NoError(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	s, err := corejson.AnyTo.JsonStringWithErr(r)
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_JsonStringWithErr_Result_HasError(t *testing.T) {
	r := corejson.NewResult.Create([]byte(`"x"`), errors.New("e"), "T")
	_, err := corejson.AnyTo.JsonStringWithErr(r)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_AnyTo_JsonStringWithErr_ResultPtr_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s, err := corejson.AnyTo.JsonStringWithErr(r)
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_JsonStringWithErr_ResultPtr_HasError(t *testing.T) {
	r := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("e"), "T")
	_, err := corejson.AnyTo.JsonStringWithErr(r)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_AnyTo_JsonStringWithErr_Any(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr(42)
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_JsonStringMust(t *testing.T) {
	s := corejson.AnyTo.JsonStringMust("hi")
	actual := args.Map{"result": s != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_I20_AnyTo_PrettyStringMust(t *testing.T) {
	s := corejson.AnyTo.PrettyStringMust("hi")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_AnyTo_UsingSerializer(t *testing.T) {
	r := corejson.AnyTo.UsingSerializer(nil)
	actual := args.Map{"result": r != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil serializer", actual)
}

func Test_I20_AnyTo_SerializedFieldsMap(t *testing.T) {
	m, err := corejson.AnyTo.SerializedFieldsMap(map[string]string{"k": "v"})
	_ = m
	_ = err
}

func Test_I20_AnyTo_SerializedJsonResult_Nil(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(nil)
	actual := args.Map{"result": r == nil || r.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error result for nil", actual)
}

func Test_I20_AnyTo_SerializedJsonResult_Error_NilErr(t *testing.T) {
	var errNil error
	r := corejson.AnyTo.SerializedJsonResult(errNil)
	actual := args.Map{"result": r == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected result", actual)
}

// ===== Serializer coverage =====

func Test_I20_Serializer_StringsApply(t *testing.T) {
	r := corejson.Serialize.StringsApply([]string{"a", "b"})
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_Serializer_FromBytes(t *testing.T) {
	r := corejson.Serialize.FromBytes([]byte(`"test"`))
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_Serializer_FromStrings(t *testing.T) {
	r := corejson.Serialize.FromStrings([]string{"a"})
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_Serializer_FromStringsSpread(t *testing.T) {
	r := corejson.Serialize.FromStringsSpread("a", "b")
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_Serializer_FromString(t *testing.T) {
	r := corejson.Serialize.FromString("hello")
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_Serializer_FromInteger(t *testing.T) {
	r := corejson.Serialize.FromInteger(42)
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_Serializer_FromInteger64(t *testing.T) {
	r := corejson.Serialize.FromInteger64(64)
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_Serializer_FromBool(t *testing.T) {
	r := corejson.Serialize.FromBool(true)
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_Serializer_FromIntegers(t *testing.T) {
	r := corejson.Serialize.FromIntegers([]int{1, 2})
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_Serializer_UsingAnyPtr(t *testing.T) {
	r := corejson.Serialize.UsingAnyPtr("test")
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_Serializer_UsingAny(t *testing.T) {
	r := corejson.Serialize.UsingAny("test")
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_Serializer_Raw(t *testing.T) {
	b, err := corejson.Serialize.Raw("test")
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Serializer_Marshal(t *testing.T) {
	b, err := corejson.Serialize.Marshal("test")
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Serializer_ApplyMust(t *testing.T) {
	r := corejson.Serialize.ApplyMust("test")
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_Serializer_ToBytesMust(t *testing.T) {
	b := corejson.Serialize.ToBytesMust("test")
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Serializer_ToSafeBytesMust(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesMust("test")
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Serializer_ToSafeBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesSwallowErr("test")
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Serializer_ToBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToBytesSwallowErr("test")
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Serializer_ToBytesErr(t *testing.T) {
	b, err := corejson.Serialize.ToBytesErr("test")
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_Serializer_ToString(t *testing.T) {
	s := corejson.Serialize.ToString("test")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_Serializer_ToStringMust(t *testing.T) {
	s := corejson.Serialize.ToStringMust("test")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_Serializer_ToStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToStringErr("test")
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_Serializer_ToPrettyStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToPrettyStringErr("test")
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_Serializer_ToPrettyStringIncludingErr(t *testing.T) {
	s := corejson.Serialize.ToPrettyStringIncludingErr("test")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_I20_Serializer_Pretty(t *testing.T) {
	s := corejson.Serialize.Pretty("test")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

// ===== Deserializer coverage =====

func Test_I20_Deserializer_UsingStringPtr_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringPtr(nil, &s)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil bytes", actual)
}

func Test_I20_Deserializer_UsingStringPtr_Valid(t *testing.T) {
	str := `"hello"`
	var s string
	err := corejson.Deserialize.UsingStringPtr(&str, &s)
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello', got ''", actual)
}

func Test_I20_Deserializer_UsingError_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingError(nil, &s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Deserializer_UsingError_Valid(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingError(errors.New(`"hello"`), &s)
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_I20_Deserializer_UsingErrorWhichJsonResult_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Deserializer_FromString(t *testing.T) {
	var s string
	err := corejson.Deserialize.FromString(`"hi"`, &s)
	actual := args.Map{"result": err != nil || s != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_I20_Deserializer_FromStringMust(t *testing.T) {
	var s string
	corejson.Deserialize.FromStringMust(`"hi"`, &s)
	actual := args.Map{"result": s != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_I20_Deserializer_FromStringMust_Panic(t *testing.T) {
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	var s string
	corejson.Deserialize.FromStringMust("not-json", &s)
}

func Test_I20_Deserializer_UsingStringOption_IgnoreEmpty(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringOption(true, "", &s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Deserializer_UsingStringIgnoreEmpty(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Deserializer_UsingBytesPointer_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointer(nil, &s)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_Deserializer_UsingBytesPointer_Valid(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointer([]byte(`"hi"`), &s)
	actual := args.Map{"result": err != nil || s != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_I20_Deserializer_UsingBytesPointerMust(t *testing.T) {
	var s string
	corejson.Deserialize.UsingBytesPointerMust([]byte(`"hi"`), &s)
	actual := args.Map{"result": s != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_I20_Deserializer_UsingBytesIf_Skip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"x"`), &s)
	actual := args.Map{"result": err != nil || s != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected skip", actual)
}

func Test_I20_Deserializer_UsingBytesIf_Do(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesIf(true, []byte(`"x"`), &s)
	actual := args.Map{"result": err != nil || s != "x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'x'", actual)
}

func Test_I20_Deserializer_UsingBytesPointerIf_Skip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected skip", actual)
}

func Test_I20_Deserializer_UsingBytesPointerIf_Do(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointerIf(true, []byte(`"x"`), &s)
	actual := args.Map{"result": err != nil || s != "x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'x'", actual)
}

func Test_I20_Deserializer_UsingBytesMust(t *testing.T) {
	var s string
	corejson.Deserialize.UsingBytesMust([]byte(`"hi"`), &s)
	actual := args.Map{"result": s != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_I20_Deserializer_UsingSafeBytesMust_Empty(t *testing.T) {
	var s string
	corejson.Deserialize.UsingSafeBytesMust([]byte{}, &s)
	// should skip
}

func Test_I20_Deserializer_UsingSafeBytesMust_Valid(t *testing.T) {
	var s string
	corejson.Deserialize.UsingSafeBytesMust([]byte(`"hi"`), &s)
	actual := args.Map{"result": s != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_I20_Deserializer_AnyToFieldsMap(t *testing.T) {
	m, err := corejson.Deserialize.AnyToFieldsMap(map[string]string{"k": "v"})
	_ = m
	_ = err
}

func Test_I20_Deserializer_MapAnyToPointer_SkipEmpty(t *testing.T) {
	var s map[string]any
	err := corejson.Deserialize.MapAnyToPointer(true, nil, &s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Deserializer_MapAnyToPointer_Valid(t *testing.T) {
	var s map[string]any
	err := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"k": "v"}, &s)
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_I20_Deserializer_UsingDeserializerToOption_SkipNil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Deserializer_UsingDeserializerToOption_NilNotSkip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerToOption(false, nil, &s)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_Deserializer_UsingDeserializerDefined_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerDefined(nil, &s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil (skip)", actual)
}

func Test_I20_Deserializer_UsingDeserializerFuncDefined_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &s)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil func", actual)
}

func Test_I20_Deserializer_UsingDeserializerFuncDefined_Valid(t *testing.T) {
	var s string
	fn := func(toPtr any) error { return nil }
	err := corejson.Deserialize.UsingDeserializerFuncDefined(fn, &s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Deserializer_UsingJsonerToAny_SkipNil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, &s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Deserializer_UsingJsonerToAny_NilNotSkip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAny(false, nil, &s)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_Deserializer_UsingJsonerToAnyMust_SkipNil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, &s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_Deserializer_UsingJsonerToAnyMust_NilNotSkip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAnyMust(false, nil, &s)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ===== NewResult creator coverage =====

func Test_I20_NewResult_UsingBytesError_Nil(t *testing.T) {
	r := corejson.NewResult.UsingBytesError(nil)
	actual := args.Map{"result": r.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_NewResult_UsingErrorStringPtr_NilPtr(t *testing.T) {
	r := corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "T")
	actual := args.Map{"result": r.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_NewResult_UsingErrorStringPtr_Valid(t *testing.T) {
	s := `"ok"`
	r := corejson.NewResult.UsingErrorStringPtr(nil, &s, "T")
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_NewResult_UsingTypePlusStringPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingTypePlusStringPtr("T", nil)
	actual := args.Map{"result": r.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_NewResult_UsingTypePlusStringPtr_Empty(t *testing.T) {
	s := ""
	r := corejson.NewResult.UsingTypePlusStringPtr("T", &s)
	actual := args.Map{"result": r.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_NewResult_UsingTypePlusStringPtr_Valid(t *testing.T) {
	s := `"hello"`
	r := corejson.NewResult.UsingTypePlusStringPtr("T", &s)
	actual := args.Map{"result": r.Length() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_NewResult_UsingStringPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingStringPtr(nil)
	actual := args.Map{"result": r.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_NewResult_UsingStringPtr_Valid(t *testing.T) {
	s := `"hello"`
	r := corejson.NewResult.UsingStringPtr(&s)
	actual := args.Map{"result": r.Length() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_NewResult_Many(t *testing.T) {
	r := corejson.NewResult.Many("a", "b", "c")
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_I20_NewResult_UsingJsoner_Nil(t *testing.T) {
	r := corejson.NewResult.UsingJsoner(nil)
	actual := args.Map{"result": r != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_NewResult_UsingSerializerFunc_Nil(t *testing.T) {
	r := corejson.NewResult.UsingSerializerFunc(nil)
	actual := args.Map{"result": r != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_NewResult_DeserializeUsingResult_HasIssues(t *testing.T) {
	errResult := corejson.NewResult.ErrorPtr(errors.New("e"))
	r := corejson.NewResult.DeserializeUsingResult(errResult)
	actual := args.Map{"result": r.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ===== CastAny coverage =====

func Test_I20_CastAny_FromToDefault_NilFrom(t *testing.T) {
	var s string
	// FromToDefault(nil, &s) → reflectionCasting returns (err, false) for nil,
	// falls through to Serialize.Apply(nil) → "null" → Unmarshal sets zero value, no error
	err := corejson.CastAny.FromToDefault(nil, &s)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error — nil serializes to null", actual)
}

func Test_I20_CastAny_FromToReflection(t *testing.T) {
	var s string
	err := corejson.CastAny.FromToReflection(`"hello"`, &s)
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_I20_CastAny_OrDeserializeTo(t *testing.T) {
	var s string
	err := corejson.CastAny.OrDeserializeTo(`"hello"`, &s)
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ===== Empty creator coverage =====

func Test_I20_Empty_ResultWithErr(t *testing.T) {
	r := corejson.Empty.ResultWithErr("T", errors.New("e"))
	actual := args.Map{"result": r.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_Empty_BytesCollection(t *testing.T) {
	bc := corejson.Empty.BytesCollection()
	actual := args.Map{"result": bc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_Empty_BytesCollectionPtr(t *testing.T) {
	bc := corejson.Empty.BytesCollectionPtr()
	actual := args.Map{"result": bc == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_I20_Empty_MapResults(t *testing.T) {
	mr := corejson.Empty.MapResults()
	actual := args.Map{"result": mr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ===== BytesToString / BytesToPrettyString =====

func Test_I20_BytesToString_Empty(t *testing.T) {
	actual := args.Map{"result": corejson.BytesToString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_BytesToString_Valid(t *testing.T) {
	actual := args.Map{"result": corejson.BytesToString([]byte("hi")) != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_I20_BytesToPrettyString_Empty(t *testing.T) {
	actual := args.Map{"result": corejson.BytesToPrettyString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_BytesToPrettyString_Valid(t *testing.T) {
	s := corejson.BytesToPrettyString([]byte(`{"a":"b"}`))
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty string", actual)
}

// ===== BytesDeepClone / BytesCloneIf =====

func Test_I20_BytesDeepClone_Empty(t *testing.T) {
	b := corejson.BytesDeepClone(nil)
	actual := args.Map{"result": len(b) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_BytesDeepClone_Valid(t *testing.T) {
	b := corejson.BytesDeepClone([]byte("hi"))
	actual := args.Map{"result": string(b) != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_I20_BytesCloneIf_NoClone(t *testing.T) {
	b := corejson.BytesCloneIf(false, []byte("hi"))
	actual := args.Map{"result": len(b) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty (no clone)", actual)
}

func Test_I20_BytesCloneIf_Clone(t *testing.T) {
	b := corejson.BytesCloneIf(true, []byte("hi"))
	actual := args.Map{"result": string(b) != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}
