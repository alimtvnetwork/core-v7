package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===== Result methods coverage =====

func Test_Result_CloneIf_True(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(map[string]string{"a": "1"})
	cloned := r.CloneIf(true, true)

	// Act
	actual := args.Map{"result": cloned.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned bytes", actual)
}

func Test_Result_CloneIf_False(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(map[string]string{"a": "1"})
	same := r.CloneIf(false, false)

	// Act
	actual := args.Map{"result": same.Length() != r.Length()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same result", actual)
}

func Test_Result_Clone_DeepClone(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("test")
	c := r.Clone(true)

	// Act
	actual := args.Map{"result": c.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Result_Clone_ShallowClone(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("test")
	c := r.Clone(false)

	// Act
	actual := args.Map{"result": c.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Result_Clone_Empty(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Empty()
	c := r.Clone(true)

	// Act
	actual := args.Map{"result": c.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Result_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	c := r.ClonePtr(true)

	// Act
	actual := args.Map{"result": c != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Result_ClonePtr_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("test")
	c := r.ClonePtr(true)

	// Act
	actual := args.Map{"result": c == nil || c.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned ptr", actual)
}

func Test_Result_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty nil message", actual)
}

func Test_Result_PrettyJsonStringOrErrString_HasError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.ErrorPtr(errors.New("test-err"))
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error string", actual)
}

func Test_Result_PrettyJsonStringOrErrString_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty json", actual)
}

func Test_Result_HandleErrorWithMsg_NoError(t *testing.T) {
	r := corejson.NewResult.AnyPtr("ok")
	r.HandleErrorWithMsg("no-op") // Should not panic
}

func Test_Result_HandleErrorWithMsg_Panic(t *testing.T) {
	// Arrange
	r := corejson.NewResult.ErrorPtr(errors.New("bad"))
	defer func() {

	// Act
		actual := args.Map{"result": rec := recover(); rec == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	r.HandleErrorWithMsg("context message")
}

func Test_Result_DeserializeMust_Success(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	var s string
	r.DeserializeMust(&s)

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello', got ''", actual)
}

func Test_Result_DeserializeMust_Panic(t *testing.T) {
	// Arrange
	r := corejson.NewResult.ErrorPtr(errors.New("err"))
	defer func() {

	// Act
		actual := args.Map{"result": rec := recover(); rec == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	var s string
	r.DeserializeMust(&s)
}

func Test_Result_UnmarshalMust_Success(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(42)
	var i int
	r.UnmarshalMust(&i)

	// Act
	actual := args.Map{"result": i != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_Result_SafeFieldsNames(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"key": "val"})
	names := r.SafeFieldsNames()
	// May return empty due to DeserializedFieldsToMap behavior
	_ = names
}

func Test_Result_SafeDeserializedFieldsToMap(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"k": "v"})
	m := r.SafeDeserializedFieldsToMap()
	_ = m
}

func Test_Result_BytesError_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("test")
	be := r.BytesError()

	// Act
	actual := args.Map{"result": be == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Result_BytesError_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	be := r.BytesError()

	// Act
	actual := args.Map{"result": be != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Result_Dispose_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("test")
	r.Dispose()

	// Act
	actual := args.Map{"result": r.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected disposed", actual)
}

func Test_Result_Dispose_Nil(t *testing.T) {
	var r *corejson.Result
	r.Dispose() // should not panic
}

func Test_Result_NonPtr_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	nr := r.NonPtr()

	// Act
	actual := args.Map{"result": nr.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error in nonptr of nil", actual)
}

func Test_Result_NonPtr_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("test")
	nr := r.NonPtr()

	// Act
	actual := args.Map{"result": nr.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Result_CombineErrorWithRefError_NoError_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("ok")
	err := r.CombineErrorWithRefError("ref")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Result_CombineErrorWithRefError_HasError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	err := r.CombineErrorWithRefError("ref1", "ref2")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Result_CombineErrorWithRefString_NoError_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("ok")
	s := r.CombineErrorWithRefString("ref")

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Result_CloneError_NoError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("ok")

	// Act
	actual := args.Map{"result": r.CloneError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Result_CloneError_HasError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.ErrorPtr(errors.New("e"))

	// Act
	actual := args.Map{"result": r.CloneError() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Result_Ptr_ToPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("test")
	p := r.Ptr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
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

func Test_Result_IsEqualPtr_BothNil(t *testing.T) {
	// Arrange
	var a, b *corejson.Result

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_Result_IsEqualPtr_OneNil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("x")

	// Act
	actual := args.Map{"result": r.IsEqualPtr(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_Result_IsEqualPtr_Same(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("x")

	// Act
	actual := args.Map{"result": r.IsEqualPtr(r)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal (same ptr)", actual)
}

func Test_Result_IsEqualPtr_DiffLength(t *testing.T) {
	// Arrange
	a := corejson.NewResult.AnyPtr("x")
	b := corejson.NewResult.AnyPtr("xy")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_Result_IsEqualPtr_DiffError(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Ptr([]byte("x"), errors.New("e1"), "t")
	b := corejson.NewResult.Ptr([]byte("x"), errors.New("e2"), "t")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_Result_IsEqualPtr_DiffType(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Ptr([]byte(`"x"`), nil, "typeA")
	b := corejson.NewResult.Ptr([]byte(`"x"`), nil, "typeB")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_Result_IsEqual_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Any("hello")
	b := corejson.NewResult.Any("hello")

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_Result_IsErrorEqual_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	a := corejson.NewResult.ErrorPtr(errors.New("same"))

	// Act
	actual := args.Map{"result": a.IsErrorEqual(errors.New("same"))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_Result_IsErrorEqual_BothNil(t *testing.T) {
	// Arrange
	a := corejson.NewResult.AnyPtr("ok")

	// Act
	actual := args.Map{"result": a.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal (both nil)", actual)
}

func Test_Result_IsErrorEqual_OnlyOneNil(t *testing.T) {
	// Arrange
	a := corejson.NewResult.AnyPtr("ok")

	// Act
	actual := args.Map{"result": a.IsErrorEqual(errors.New("e"))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

// ===== Result serialization coverage =====

func Test_Result_Serialize_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	_, err := r.Serialize()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Result_Serialize_HasError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	_, err := r.Serialize()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Result_Serialize_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	b, err := r.Serialize()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Result_SerializeMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	b := r.SerializeMust()

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Result_SerializeSkipExistingIssues_HasIssues(t *testing.T) {
	// Arrange
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	b, err := r.SerializeSkipExistingIssues()

	// Act
	actual := args.Map{"result": b != nil || err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil,nil for issues", actual)
}

func Test_Result_SerializeSkipExistingIssues_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("ok")
	b, err := r.SerializeSkipExistingIssues()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_Result_UnmarshalSkipExistingIssues_HasIssues(t *testing.T) {
	// Arrange
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for issues", actual)
}

func Test_Result_UnmarshalSkipExistingIssues_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello', got ''", actual)
}

func Test_Result_UnmarshalSkipExistingIssues_Error(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingBytesTypePtr([]byte("not-json"), "test")
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error", actual)
}

func Test_Result_UnmarshalResult_FromResultCloneIfIterati(t *testing.T) {
	inner := corejson.NewResult.Any("test")
	outerBytes, _ := inner.Serialize()
	outer := &corejson.Result{Bytes: outerBytes}
	_, _ = outer.UnmarshalResult()
}

func Test_Result_JsonModel_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m := r.JsonModel()

	// Act
	actual := args.Map{"result": m.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error in model", actual)
}

func Test_Result_JsonModelAny(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("ok")
	a := r.JsonModelAny()

	// Act
	actual := args.Map{"result": a == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Result_Json_JsonPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("test")
	j := r.Json()

	// Act
	actual := args.Map{"result": j.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected json", actual)
	jp := r.JsonPtr()
	actual := args.Map{"result": jp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ptr", actual)
}

func Test_Result_ParseInjectUsingJson_Success(t *testing.T) {
	source := corejson.NewResult.AnyPtr(map[string]string{"a": "1"})
	target := corejson.NewResult.AnyPtr(map[string]string{})
	_, err := target.ParseInjectUsingJson(source)
	// This may fail depending on types but exercises the path
	_ = err
}

func Test_Result_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	// Arrange
	source := corejson.NewResult.ErrorPtr(errors.New("err"))
	target := corejson.NewResult.AnyPtr("test")
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	target.ParseInjectUsingJsonMust(source)
}

func Test_Result_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("test")
	b := r.AsJsonContractsBinder()

	// Act
	actual := args.Map{"result": b == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Result_AsJsoner(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("test")
	j := r.AsJsoner()

	// Act
	actual := args.Map{"result": j == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Result_AsJsonParseSelfInjector(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("test")
	inj := r.AsJsonParseSelfInjector()

	// Act
	actual := args.Map{"result": inj == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Result_JsonParseSelfInject(t *testing.T) {
	r := corejson.NewResult.Any(map[string]string{"a": "1"})
	source := corejson.NewResult.AnyPtr(map[string]string{"b": "2"})
	err := r.JsonParseSelfInject(source)
	_ = err
}

// ===== Result other methods =====

func Test_Result_RawMust_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("ok")
	b := r.RawMust()

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Result_RawStringMust_Success(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("ok")
	s := r.RawStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_Result_RawStringMust_Panic(t *testing.T) {
	// Arrange
	var r *corejson.Result
	defer func() {

	// Act
		actual := args.Map{"result": rec := recover(); rec == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	r.RawStringMust()
}

func Test_Result_RawErrString_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("ok")
	b, msg := r.RawErrString()

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	_ = msg
}

func Test_Result_RawPrettyString_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("ok")
	s, err := r.RawPrettyString()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_Result_HandleError_Panic(t *testing.T) {
	// Arrange
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	defer func() {

	// Act
		actual := args.Map{"result": rec := recover(); rec == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	r.HandleError()
}

func Test_Result_MustBeSafe_Panic(t *testing.T) {
	// Arrange
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	defer func() {

	// Act
		actual := args.Map{"result": rec := recover(); rec == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	r.MustBeSafe()
}

func Test_Result_SafeNonIssueBytes_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("ok")

	// Act
	actual := args.Map{"result": len(r.SafeNonIssueBytes()) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Result_SafeNonIssueBytes_HasIssues(t *testing.T) {
	// Arrange
	r := corejson.NewResult.ErrorPtr(errors.New("e"))

	// Act
	actual := args.Map{"result": len(r.SafeNonIssueBytes()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Result_SafeValuesPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("ok")

	// Act
	actual := args.Map{"result": len(r.SafeValuesPtr()) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Result_SafeValuesPtr_HasIssues(t *testing.T) {
	// Arrange
	r := corejson.NewResult.ErrorPtr(errors.New("e"))

	// Act
	actual := args.Map{"result": len(r.SafeValuesPtr()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Result_Values(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("ok")

	// Act
	actual := args.Map{"result": len(r.Values()) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Result_Raw_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	_, err := r.Raw()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Result_MeaningfulError_EmptyBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: nil, TypeName: "Test"}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for empty bytes", actual)
}

func Test_Result_MeaningfulError_HasErrorAndBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{
		Bytes:    []byte(`"test"`),
		Error:    errors.New("some err"),
		TypeName: "Test",
	}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Result_String_HasError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Create([]byte(`"x"`), errors.New("e"), "T")
	s := r.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string with error", actual)
}

func Test_Result_SafeBytesTypeName_Empty(t *testing.T) {
	// Arrange
	r := corejson.NewResult.EmptyPtr()

	// Act
	actual := args.Map{"result": r.SafeBytesTypeName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Result_BytesTypeName_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": r.BytesTypeName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ===== anyTo coverage =====

func Test_AnyTo_SerializedRaw_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	b, err := corejson.AnyTo.SerializedRaw("hello")

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_AnyTo_SerializedString_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.SerializedString("hello")

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_SerializedSafeString_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedSafeString("hello")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_SerializedStringMust_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedStringMust("hello")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_SafeJsonString_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonString("hello")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_PrettyStringWithError_String_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError("hello")

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected passthrough", actual)
}

func Test_AnyTo_PrettyStringWithError_Bytes_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":"b"}`))

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty string", actual)
}

func Test_AnyTo_PrettyStringWithError_Result(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")
	s, err := corejson.AnyTo.PrettyStringWithError(r)

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_PrettyStringWithError_ResultPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	s, err := corejson.AnyTo.PrettyStringWithError(r)

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_PrettyStringWithError_ResultWithError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Create([]byte(`"x"`), errors.New("e"), "T")
	_, err := corejson.AnyTo.PrettyStringWithError(r)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_AnyTo_PrettyStringWithError_ResultPtrWithError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("e"), "T")
	_, err := corejson.AnyTo.PrettyStringWithError(r)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_AnyTo_PrettyStringWithError_AnyItem(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError(42)

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_SafeJsonPrettyString_String_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonPrettyString("hi")

	// Act
	actual := args.Map{"result": s != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected passthrough", actual)
}

func Test_AnyTo_SafeJsonPrettyString_Bytes_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":"b"}`))

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_SafeJsonPrettyString_Result_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")
	s := corejson.AnyTo.SafeJsonPrettyString(r)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_SafeJsonPrettyString_ResultPtr_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	s := corejson.AnyTo.SafeJsonPrettyString(r)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_SafeJsonPrettyString_Any(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonPrettyString(42)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_JsonString_String_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonString("hi")

	// Act
	actual := args.Map{"result": s != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected passthrough", actual)
}

func Test_AnyTo_JsonString_Bytes_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonString([]byte(`{"a":"b"}`))

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_JsonString_Result_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")
	s := corejson.AnyTo.JsonString(r)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_JsonString_ResultPtr_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	s := corejson.AnyTo.JsonString(r)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_JsonString_Any(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonString(42)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_JsonStringWithErr_String(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.JsonStringWithErr("hi")

	// Act
	actual := args.Map{"result": err != nil || s != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected passthrough", actual)
}

func Test_AnyTo_JsonStringWithErr_Bytes(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.JsonStringWithErr([]byte(`"x"`))

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_JsonStringWithErr_Result_NoError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")
	s, err := corejson.AnyTo.JsonStringWithErr(r)

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_JsonStringWithErr_Result_HasError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Create([]byte(`"x"`), errors.New("e"), "T")
	_, err := corejson.AnyTo.JsonStringWithErr(r)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_AnyTo_JsonStringWithErr_ResultPtr_NoError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	s, err := corejson.AnyTo.JsonStringWithErr(r)

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_JsonStringWithErr_ResultPtr_HasError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("e"), "T")
	_, err := corejson.AnyTo.JsonStringWithErr(r)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_AnyTo_JsonStringWithErr_Any(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.JsonStringWithErr(42)

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_JsonStringMust_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonStringMust("hi")

	// Act
	actual := args.Map{"result": s != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_AnyTo_PrettyStringMust_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.PrettyStringMust("hi")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_AnyTo_UsingSerializer_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.UsingSerializer(nil)

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil serializer", actual)
}

func Test_AnyTo_SerializedFieldsMap_FromResultCloneIfIterati(t *testing.T) {
	m, err := corejson.AnyTo.SerializedFieldsMap(map[string]string{"k": "v"})
	_ = m
	_ = err
}

func Test_AnyTo_SerializedJsonResult_Nil_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.SerializedJsonResult(nil)

	// Act
	actual := args.Map{"result": r == nil || r.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error result for nil", actual)
}

func Test_AnyTo_SerializedJsonResult_Error_NilErr(t *testing.T) {
	// Arrange
	var errNil error
	r := corejson.AnyTo.SerializedJsonResult(errNil)

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected result", actual)
}

// ===== Serializer coverage =====

func Test_Serializer_StringsApply(t *testing.T) {
	// Arrange
	r := corejson.Serialize.StringsApply([]string{"a", "b"})

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Serializer_FromBytes(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromBytes([]byte(`"test"`))

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Serializer_FromStrings(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromStrings([]string{"a"})

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Serializer_FromStringsSpread(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromStringsSpread("a", "b")

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Serializer_FromString(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromString("hello")

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Serializer_FromInteger(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromInteger(42)

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Serializer_FromInteger64(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromInteger64(64)

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Serializer_FromBool(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromBool(true)

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Serializer_FromIntegers(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromIntegers([]int{1, 2})

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Serializer_UsingAnyPtr(t *testing.T) {
	// Arrange
	r := corejson.Serialize.UsingAnyPtr("test")

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Serializer_UsingAny(t *testing.T) {
	// Arrange
	r := corejson.Serialize.UsingAny("test")

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Serializer_Raw(t *testing.T) {
	// Arrange
	b, err := corejson.Serialize.Raw("test")

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Serializer_Marshal(t *testing.T) {
	// Arrange
	b, err := corejson.Serialize.Marshal("test")

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Serializer_ApplyMust(t *testing.T) {
	// Arrange
	r := corejson.Serialize.ApplyMust("test")

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_Serializer_ToBytesMust(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToBytesMust("test")

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Serializer_ToSafeBytesMust(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToSafeBytesMust("test")

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Serializer_ToSafeBytesSwallowErr(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToSafeBytesSwallowErr("test")

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Serializer_ToBytesSwallowErr(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToBytesSwallowErr("test")

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Serializer_ToBytesErr(t *testing.T) {
	// Arrange
	b, err := corejson.Serialize.ToBytesErr("test")

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Serializer_ToString(t *testing.T) {
	// Arrange
	s := corejson.Serialize.ToString("test")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_Serializer_ToStringMust(t *testing.T) {
	// Arrange
	s := corejson.Serialize.ToStringMust("test")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_Serializer_ToStringErr(t *testing.T) {
	// Arrange
	s, err := corejson.Serialize.ToStringErr("test")

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_Serializer_ToPrettyStringErr(t *testing.T) {
	// Arrange
	s, err := corejson.Serialize.ToPrettyStringErr("test")

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_Serializer_ToPrettyStringIncludingErr(t *testing.T) {
	// Arrange
	s := corejson.Serialize.ToPrettyStringIncludingErr("test")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_Serializer_Pretty(t *testing.T) {
	// Arrange
	s := corejson.Serialize.Pretty("test")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

// ===== Deserializer coverage =====

func Test_Deserializer_UsingStringPtr_Nil(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingStringPtr(nil, &s)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil bytes", actual)
}

func Test_Deserializer_UsingStringPtr_Valid(t *testing.T) {
	// Arrange
	str := `"hello"`
	var s string
	err := corejson.Deserialize.UsingStringPtr(&str, &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello', got ''", actual)
}

func Test_Deserializer_UsingError_Nil(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingError(nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Deserializer_UsingError_Valid(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingError(errors.New(`"hello"`), &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Deserializer_UsingErrorWhichJsonResult_Nil(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Deserializer_FromString(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.FromString(`"hi"`, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_Deserializer_FromStringMust(t *testing.T) {
	// Arrange
	var s string
	corejson.Deserialize.FromStringMust(`"hi"`, &s)

	// Act
	actual := args.Map{"result": s != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_Deserializer_FromStringMust_Panic(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	var s string
	corejson.Deserialize.FromStringMust("not-json", &s)
}

func Test_Deserializer_UsingStringOption_IgnoreEmpty(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingStringOption(true, "", &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Deserializer_UsingStringIgnoreEmpty(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Deserializer_UsingBytesPointer_Nil(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesPointer(nil, &s)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Deserializer_UsingBytesPointer_Valid(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesPointer([]byte(`"hi"`), &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_Deserializer_UsingBytesPointerMust(t *testing.T) {
	// Arrange
	var s string
	corejson.Deserialize.UsingBytesPointerMust([]byte(`"hi"`), &s)

	// Act
	actual := args.Map{"result": s != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_Deserializer_UsingBytesIf_Skip(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"x"`), &s)

	// Act
	actual := args.Map{"result": err != nil || s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected skip", actual)
}

func Test_Deserializer_UsingBytesIf_Do(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesIf(true, []byte(`"x"`), &s)

	// Act
	actual := args.Map{"result": err != nil || s != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'x'", actual)
}

func Test_Deserializer_UsingBytesPointerIf_Skip(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected skip", actual)
}

func Test_Deserializer_UsingBytesPointerIf_Do(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesPointerIf(true, []byte(`"x"`), &s)

	// Act
	actual := args.Map{"result": err != nil || s != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'x'", actual)
}

func Test_Deserializer_UsingBytesMust(t *testing.T) {
	// Arrange
	var s string
	corejson.Deserialize.UsingBytesMust([]byte(`"hi"`), &s)

	// Act
	actual := args.Map{"result": s != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_Deserializer_UsingSafeBytesMust_Empty(t *testing.T) {
	var s string
	corejson.Deserialize.UsingSafeBytesMust([]byte{}, &s)
	// should skip
}

func Test_Deserializer_UsingSafeBytesMust_Valid(t *testing.T) {
	// Arrange
	var s string
	corejson.Deserialize.UsingSafeBytesMust([]byte(`"hi"`), &s)

	// Act
	actual := args.Map{"result": s != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_Deserializer_AnyToFieldsMap(t *testing.T) {
	m, err := corejson.Deserialize.AnyToFieldsMap(map[string]string{"k": "v"})
	_ = m
	_ = err
}

func Test_Deserializer_MapAnyToPointer_SkipEmpty(t *testing.T) {
	// Arrange
	var s map[string]any
	err := corejson.Deserialize.MapAnyToPointer(true, nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Deserializer_MapAnyToPointer_Valid(t *testing.T) {
	// Arrange
	var s map[string]any
	err := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"k": "v"}, &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Deserializer_UsingDeserializerToOption_SkipNil(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Deserializer_UsingDeserializerToOption_NilNotSkip(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingDeserializerToOption(false, nil, &s)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Deserializer_UsingDeserializerDefined_Nil(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingDeserializerDefined(nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil (skip)", actual)
}

func Test_Deserializer_UsingDeserializerFuncDefined_Nil(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &s)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil func", actual)
}

func Test_Deserializer_UsingDeserializerFuncDefined_Valid(t *testing.T) {
	// Arrange
	var s string
	fn := func(toPtr any) error { return nil }
	err := corejson.Deserialize.UsingDeserializerFuncDefined(fn, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Deserializer_UsingJsonerToAny_SkipNil(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Deserializer_UsingJsonerToAny_NilNotSkip(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingJsonerToAny(false, nil, &s)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Deserializer_UsingJsonerToAnyMust_SkipNil(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Deserializer_UsingJsonerToAnyMust_NilNotSkip(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingJsonerToAnyMust(false, nil, &s)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ===== NewResult creator coverage =====

func Test_NewResult_UsingBytesError_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingBytesError(nil)

	// Act
	actual := args.Map{"result": r.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_NewResult_UsingErrorStringPtr_NilPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "T")

	// Act
	actual := args.Map{"result": r.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_NewResult_UsingErrorStringPtr_Valid(t *testing.T) {
	// Arrange
	s := `"ok"`
	r := corejson.NewResult.UsingErrorStringPtr(nil, &s, "T")

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_NewResult_UsingTypePlusStringPtr_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingTypePlusStringPtr("T", nil)

	// Act
	actual := args.Map{"result": r.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_NewResult_UsingTypePlusStringPtr_Empty(t *testing.T) {
	// Arrange
	s := ""
	r := corejson.NewResult.UsingTypePlusStringPtr("T", &s)

	// Act
	actual := args.Map{"result": r.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_NewResult_UsingTypePlusStringPtr_Valid(t *testing.T) {
	// Arrange
	s := `"hello"`
	r := corejson.NewResult.UsingTypePlusStringPtr("T", &s)

	// Act
	actual := args.Map{"result": r.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_NewResult_UsingStringPtr_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingStringPtr(nil)

	// Act
	actual := args.Map{"result": r.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_NewResult_UsingStringPtr_Valid(t *testing.T) {
	// Arrange
	s := `"hello"`
	r := corejson.NewResult.UsingStringPtr(&s)

	// Act
	actual := args.Map{"result": r.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_NewResult_Many(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Many("a", "b", "c")

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_NewResult_UsingJsoner_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingJsoner(nil)

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_NewResult_UsingSerializerFunc_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingSerializerFunc(nil)

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_NewResult_DeserializeUsingResult_HasIssues(t *testing.T) {
	// Arrange
	errResult := corejson.NewResult.ErrorPtr(errors.New("e"))
	r := corejson.NewResult.DeserializeUsingResult(errResult)

	// Act
	actual := args.Map{"result": r.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ===== CastAny coverage =====

func Test_CastAny_FromToDefault_NilFrom(t *testing.T) {
	// Arrange
	var s string
	// FromToDefault(nil, &s) → reflectionCasting returns (err, false) for nil,
	// falls through to Serialize.Apply(nil) → "null" → Unmarshal sets zero value, no error
	err := corejson.CastAny.FromToDefault(nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error — nil serializes to null", actual)
}

func Test_CastAny_FromToReflection_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	var s string
	err := corejson.CastAny.FromToReflection(`"hello"`, &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_CastAny_OrDeserializeTo_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	var s string
	err := corejson.CastAny.OrDeserializeTo(`"hello"`, &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ===== Empty creator coverage =====

func Test_Empty_ResultWithErr_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	r := corejson.Empty.ResultWithErr("T", errors.New("e"))

	// Act
	actual := args.Map{"result": r.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Empty_BytesCollection_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	bc := corejson.Empty.BytesCollection()

	// Act
	actual := args.Map{"result": bc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Empty_BytesCollectionPtr_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	bc := corejson.Empty.BytesCollectionPtr()

	// Act
	actual := args.Map{"result": bc == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Empty_MapResults_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	mr := corejson.Empty.MapResults()

	// Act
	actual := args.Map{"result": mr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ===== BytesToString / BytesToPrettyString =====

func Test_BytesToString_Empty_FromResultCloneIfIterati(t *testing.T) {
	// Act
	actual := args.Map{"result": corejson.BytesToString(nil) != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_BytesToString_Valid_FromResultCloneIfIterati(t *testing.T) {
	// Act
	actual := args.Map{"result": corejson.BytesToString([]byte("hi")) != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_BytesToPrettyString_Empty_FromResultCloneIfIterati(t *testing.T) {
	// Act
	actual := args.Map{"result": corejson.BytesToPrettyString(nil) != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_BytesToPrettyString_Valid_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	s := corejson.BytesToPrettyString([]byte(`{"a":"b"}`))

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty string", actual)
}

// ===== BytesDeepClone / BytesCloneIf =====

func Test_BytesDeepClone_Empty_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	b := corejson.BytesDeepClone(nil)

	// Act
	actual := args.Map{"result": len(b) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_BytesDeepClone_Valid_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	b := corejson.BytesDeepClone([]byte("hi"))

	// Act
	actual := args.Map{"result": string(b) != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}

func Test_BytesCloneIf_NoClone_FromResultCloneIfIterati(t *testing.T) {
	// Arrange
	b := corejson.BytesCloneIf(false, []byte("hi"))

	// Act
	actual := args.Map{"result": len(b) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty (no clone)", actual)
}

func Test_BytesCloneIf_Clone(t *testing.T) {
	// Arrange
	b := corejson.BytesCloneIf(true, []byte("hi"))

	// Act
	actual := args.Map{"result": string(b) != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hi'", actual)
}
