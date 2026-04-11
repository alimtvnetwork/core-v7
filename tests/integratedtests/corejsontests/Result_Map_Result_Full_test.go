package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Result.go — comprehensive coverage for ALL uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Result_Map_WithAllFields(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"hello"`), Error: errors.New("e"), TypeName: "T"}
	m := r.Map()

	// Act
	actual := args.Map{"result": len(m) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 fields in map", actual)
}

func Test_Result_Map_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m := r.Map()

	// Act
	actual := args.Map{"result": len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty map for nil", actual)
}

func Test_Result_Map_NoError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	m := r.Map()

	// Act
	actual := args.Map{
		"result": _,
		ok := m["Error"]; ok,
	}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have error key", actual)
}

func Test_Result_SafeDeserializedFieldsToMap(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"k": "v"})
	fm := r.SafeDeserializedFieldsToMap()
	_ = fm
}

func Test_Result_SafeDeserializedFieldsToMap_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	fm := r.SafeDeserializedFieldsToMap()

	// Act
	actual := args.Map{"result": len(fm) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Result_FieldsNames(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"a": "1", "b": "2"})
	names, err := r.FieldsNames()
	_ = names
	_ = err
}

func Test_Result_SafeFieldsNames(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"a": "1"})
	names := r.SafeFieldsNames()
	_ = names
}

func Test_Result_BytesTypeName(t *testing.T) {
	// Arrange
	r := &corejson.Result{TypeName: "MyType"}

	// Act
	actual := args.Map{"result": r.BytesTypeName() != "MyType"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "wrong type name", actual)
}

func Test_Result_BytesTypeName_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": r.BytesTypeName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_Result_SafeBytesTypeName(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	name := r.SafeBytesTypeName()

	// Act
	actual := args.Map{"result": name == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Result_SafeBytesTypeName_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	name := r.SafeBytesTypeName()

	// Act
	actual := args.Map{"result": name != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for empty result", actual)
}

func Test_Result_SafeString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	s := r.SafeString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Result_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected message for nil", actual)
}

func Test_Result_PrettyJsonStringOrErrString_WithError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error message", actual)
}

func Test_Result_PrettyJsonStringOrErrString_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(map[string]int{"a": 1})
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty json", actual)
}

func Test_Result_String_WithError(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e"), TypeName: "T"}
	s := r.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Result_String_NoError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")
	s := r.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Result_SafeNonIssueBytes(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	b := r.SafeNonIssueBytes()

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Result_SafeNonIssueBytes_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	b := r.SafeNonIssueBytes()

	// Act
	actual := args.Map{"result": len(b) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Result_Values(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("x")

	// Act
	actual := args.Map{"result": len(r.Values()) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected values", actual)
}

func Test_Result_SafeValues(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("x")

	// Act
	actual := args.Map{"result": len(r.SafeValues()) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected values", actual)
}

func Test_Result_SafeValues_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": len(r.SafeValues()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Result_SafeValuesPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("x")

	// Act
	actual := args.Map{"result": len(r.SafeValuesPtr()) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected values", actual)
}

func Test_Result_SafeValuesPtr_Issues(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}

	// Act
	actual := args.Map{"result": len(r.SafeValuesPtr()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Result_RawMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	b := r.RawMust()

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_Result_RawString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	s, err := r.RawString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Result_RawStringMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	s := r.RawStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_Result_RawErrString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	b, errMsg := r.RawErrString()

	// Act
	actual := args.Map{"result": len(b) == 0 || errMsg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Result_RawPrettyString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(map[string]int{"a": 1})
	s, err := r.RawPrettyString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Result_MeaningfulErrorMessage_NoError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	msg := r.MeaningfulErrorMessage()

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Result_MeaningfulErrorMessage_WithError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	msg := r.MeaningfulErrorMessage()

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected message", actual)
}

func Test_Result_HasSafeItems(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")

	// Act
	actual := args.Map{"result": r.HasSafeItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Result_HasSafeItems_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}

	// Act
	actual := args.Map{"result": r.HasSafeItems()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_Result_HasJsonBytes(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")

	// Act
	actual := args.Map{"result": r.HasJsonBytes()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Result_HasAnyItem(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")

	// Act
	actual := args.Map{"result": r.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Result_HasJson(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")

	// Act
	actual := args.Map{"result": r.HasJson()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Result_DeserializeMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	var s string
	r.DeserializeMust(&s)

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Result_UnmarshalMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(42)
	var n int
	r.UnmarshalMust(&n)

	// Act
	actual := args.Map{"result": n != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Result_SerializeSkipExistingIssues_HasIssues(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	b, err := r.SerializeSkipExistingIssues()

	// Act
	actual := args.Map{"result": b != nil || err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil,nil for issues", actual)
}

func Test_Result_SerializeSkipExistingIssues_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	b, err := r.SerializeSkipExistingIssues()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Result_Serialize(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	b, err := r.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

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

func Test_Result_Serialize_WithError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	_, err := r.Serialize()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
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

func Test_Result_UnmarshalSkipExistingIssues_HasIssues(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for issues", actual)
}

func Test_Result_UnmarshalResult(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	r, err := jr.UnmarshalResult()
	_ = r
	_ = err
}

func Test_Result_JsonModel(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	m := r.JsonModel()

	// Act
	actual := args.Map{"result": m.TypeName == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type name", actual)
}

func Test_Result_JsonModel_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m := r.JsonModel()

	// Act
	actual := args.Map{"result": m.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_Result_JsonModelAny(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	a := r.JsonModelAny()

	// Act
	actual := args.Map{"result": a == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Result_Json(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")
	j := r.Json()

	// Act
	actual := args.Map{"result": j.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

func Test_Result_JsonPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")
	j := r.JsonPtr()

	// Act
	actual := args.Map{"result": j.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

func Test_Result_ParseInjectUsingJson(t *testing.T) {
	inner := corejson.NewResult.Any(corejson.Result{Bytes: []byte(`"hi"`), TypeName: "t"})
	target := &corejson.Result{}
	_, err := target.ParseInjectUsingJson(inner.Ptr())
	_ = err
}

func Test_Result_CloneError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	ce := r.CloneError()

	// Act
	actual := args.Map{"result": ce == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Result_CloneError_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("x")
	ce := r.CloneError()

	// Act
	actual := args.Map{"result": ce != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Result_Ptr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")
	p := r.Ptr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ptr", actual)
}

func Test_Result_NonPtr_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	np := r.NonPtr()

	// Act
	actual := args.Map{"result": np.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Result_ToPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")
	p := r.ToPtr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ptr", actual)
}

func Test_Result_ToNonPtr(t *testing.T) {
	r := corejson.NewResult.Any("x")
	np := r.ToNonPtr()
	_ = np
}

func Test_Result_IsEqualPtr(t *testing.T) {
	// Arrange
	a := corejson.NewResult.AnyPtr("hello")
	b := corejson.NewResult.AnyPtr("hello")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_Result_IsEqualPtr_BothNil(t *testing.T) {
	// Arrange
	var a, b *corejson.Result

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal for nil", actual)
}

func Test_Result_IsEqualPtr_OneNil(t *testing.T) {
	// Arrange
	a := corejson.NewResult.AnyPtr("x")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_Result_IsEqualPtr_Same(t *testing.T) {
	// Arrange
	a := corejson.NewResult.AnyPtr("x")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(a)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal for same ptr", actual)
}

func Test_Result_IsEqualPtr_DiffLen(t *testing.T) {
	// Arrange
	a := corejson.NewResult.AnyPtr("hello")
	b := corejson.NewResult.AnyPtr("hi")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_Result_IsEqualPtr_DiffError(t *testing.T) {
	// Arrange
	a := &corejson.Result{Bytes: []byte("x"), Error: errors.New("a")}
	b := &corejson.Result{Bytes: []byte("x"), Error: errors.New("b")}

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_Result_IsEqualPtr_DiffType(t *testing.T) {
	// Arrange
	a := &corejson.Result{Bytes: []byte("x"), TypeName: "A"}
	b := &corejson.Result{Bytes: []byte("x"), TypeName: "B"}

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_Result_CombineErrorWithRefString(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	s := r.CombineErrorWithRefString("ref1", "ref2")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_Result_CombineErrorWithRefString_NoErr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("x")
	s := r.CombineErrorWithRefString("ref")

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Result_CombineErrorWithRefError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	err := r.CombineErrorWithRefError("ref")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Result_CombineErrorWithRefError_NoErr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("x")
	err := r.CombineErrorWithRefError("ref")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Result_IsEqual(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Any("hello")
	b := corejson.NewResult.Any("hello")

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_Result_IsEqual_DiffLen(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Any("hello")
	b := corejson.NewResult.Any("hi")

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_Result_BytesError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("x")
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

func Test_Result_Dispose(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("x")
	r.Dispose()

	// Act
	actual := args.Map{"result": r.Bytes != nil || r.Error != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected disposed", actual)
}

func Test_Result_Dispose_Nil(t *testing.T) {
	var r *corejson.Result
	r.Dispose() // should not panic
}

func Test_Result_CloneIf(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")
	c1 := r.CloneIf(true, true)
	c2 := r.CloneIf(false, false)

	// Act
	actual := args.Map{"result": c1.HasError() || c2.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

func Test_Result_ClonePtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	c := r.ClonePtr(true)

	// Act
	actual := args.Map{"result": c == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected clone", actual)
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

func Test_Result_Clone_DeepAndShallow(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")
	deep := r.Clone(true)
	shallow := r.Clone(false)

	// Act
	actual := args.Map{"result": deep.HasError() || shallow.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Result_Clone_Empty(t *testing.T) {
	r := corejson.NewResult.Any("")
	c := r.Clone(true)
	_ = c
}

func Test_Result_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")
	b := r.AsJsonContractsBinder()

	// Act
	actual := args.Map{"result": b == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Result_AsJsoner(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")
	j := r.AsJsoner()

	// Act
	actual := args.Map{"result": j == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Result_JsonParseSelfInject(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	inner := corejson.NewResult.AnyPtr(corejson.NewResult.Any("world"))
	err := r.JsonParseSelfInject(inner)
	_ = err
}

func Test_Result_AsJsonParseSelfInjector(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")
	inj := r.AsJsonParseSelfInjector()

	// Act
	actual := args.Map{"result": inj == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Result_InjectInto(t *testing.T) {
	r := corejson.NewResult.AnyPtr([]string{"a", "b"})
	target := corejson.NewResult.Any("x")
	err := r.InjectInto(&target)
	_ = err
}

// ── IsErrorEqual branches ──

func Test_Result_IsErrorEqual_BothNil(t *testing.T) {
	// Arrange
	r := &corejson.Result{}

	// Act
	actual := args.Map{"result": r.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true both nil", actual)
}

func Test_Result_IsErrorEqual_OneNil(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}

	// Act
	actual := args.Map{"result": r.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false one nil", actual)
}

func Test_Result_IsErrorEqual_LeftNil(t *testing.T) {
	// Arrange
	r := &corejson.Result{}

	// Act
	actual := args.Map{"result": r.IsErrorEqual(errors.New("e"))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false left nil", actual)
}

func Test_Result_IsErrorEqual_Same(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}

	// Act
	actual := args.Map{"result": r.IsErrorEqual(errors.New("e"))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true same msg", actual)
}

func Test_Result_IsErrorEqual_Different(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("a")}

	// Act
	actual := args.Map{"result": r.IsErrorEqual(errors.New("b"))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false different", actual)
}
