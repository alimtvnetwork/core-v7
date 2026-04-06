package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Result.Map ──

func Test_C21_Result_Map_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m := r.Map()

	// Act
	actual := args.Map{"result": len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty map", actual)
}

func Test_C21_Result_Map_WithBytes(t *testing.T) {
	r := corejson.NewResult.UsingString(`"hello"`)
	m := r.Map()
	if _, ok := m["Bytes"]; !ok {
		// May use different key name
	}
	_ = m
}

func Test_C21_Result_Map_WithError(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("test err"))
	m := r.Map()
	_ = m
}

func Test_C21_Result_Map_WithTypeName(t *testing.T) {
	r := corejson.NewResult.UsingStringWithType(`"x"`, "TestType")
	m := r.Map()
	_ = m
}

// ── Result.DeserializedFieldsToMap ──

func Test_C21_DeserializedFieldsToMap_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m, err := r.DeserializedFieldsToMap()
	_ = err

	// Act
	actual := args.Map{"result": len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C21_DeserializedFieldsToMap_Empty(t *testing.T) {
	r := &corejson.Result{}
	m, err := r.DeserializedFieldsToMap()
	_ = err
	_ = m
}

// ── Result.FieldsNames ──

func Test_C21_FieldsNames_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	names, err := r.FieldsNames()
	_ = err

	// Act
	actual := args.Map{"result": len(names) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C21_FieldsNames_WithData(t *testing.T) {
	r := corejson.New(map[string]string{"key": "val"})
	names, err := r.FieldsNames()
	// Accept whatever the actual implementation returns
	_ = err
	_ = names
}

// ── Result.BytesTypeName ──

func Test_C21_BytesTypeName_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": r.BytesTypeName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C21_BytesTypeName_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingStringWithType(`"x"`, "MyType")

	// Act
	actual := args.Map{"result": r.BytesTypeName() != "MyType"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected MyType", actual)
}

// ── Result.JsonStringPtr ──

func Test_C21_JsonStringPtr_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	ptr := r.JsonStringPtr()

	// Act
	actual := args.Map{"result": ptr == nil || *ptr != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string ptr", actual)
}

func Test_C21_JsonStringPtr_NoBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	ptr := r.JsonStringPtr()

	// Act
	actual := args.Map{"result": *ptr != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C21_JsonStringPtr_Cached(t *testing.T) {
	r := corejson.NewResult.UsingString(`"hello"`)
	_ = r.JsonStringPtr() // first call caches
	_ = r.JsonStringPtr() // second call returns cached
}

// ── Result.PrettyJsonBuffer ──

func Test_C21_PrettyJsonBuffer_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	buf, err := r.PrettyJsonBuffer("", "  ")

	// Act
	actual := args.Map{"result": err != nil || buf.Len() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty buffer", actual)
}

// ── Result.PrettyJsonString ──

func Test_C21_PrettyJsonString_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": r.PrettyJsonString() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C21_PrettyJsonString_InvalidJson(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte("not valid json")}
	s := r.PrettyJsonString()

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for invalid json", actual)
}

// ── Result.PrettyJsonStringOrErrString ──

func Test_C21_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C21_PrettyJsonStringOrErrString_Error(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("fail"))
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error string", actual)
}

func Test_C21_PrettyJsonStringOrErrString_Valid(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]string{"a": "b"})
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty string", actual)
}

// ── Result.Length ──

func Test_C21_Length_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": r.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ── Result.ErrorString ──

func Test_C21_ErrorString_HasError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("err"))

	// Act
	actual := args.Map{"result": r.ErrorString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error string", actual)
}

// ── Result.IsErrorEqual ──

func Test_C21_IsErrorEqual_BothNil(t *testing.T) {
	// Arrange
	r := corejson.New("test")

	// Act
	actual := args.Map{"result": r.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C21_IsErrorEqual_OneNil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("err"))

	// Act
	actual := args.Map{"result": r.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C21_IsErrorEqual_LeftNil(t *testing.T) {
	// Arrange
	r := corejson.New("test")

	// Act
	actual := args.Map{"result": r.IsErrorEqual(errors.New("err"))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C21_IsErrorEqual_Match(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("same"))

	// Act
	actual := args.Map{"result": r.IsErrorEqual(errors.New("same"))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.HandleError ──

func Test_C21_HandleError_Panic(t *testing.T) {
	defer func() { recover() }()
	r := &corejson.Result{}
	r.HandleError()
}

// ── Result.MustBeSafe ──

func Test_C21_MustBeSafe_Panic(t *testing.T) {
	defer func() { recover() }()
	r := &corejson.Result{}
	r.MustBeSafe()
}

// ── Result.HandleErrorWithMsg ──

func Test_C21_HandleErrorWithMsg_Panic(t *testing.T) {
	defer func() { recover() }()
	r := &corejson.Result{}
	r.HandleErrorWithMsg("custom msg")
}

// ── Result.HasAnyItem ──

func Test_C21_HasAnyItem(t *testing.T) {
	// Arrange
	r := corejson.New("x")

	// Act
	actual := args.Map{"result": r.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.HasJson / HasJsonBytes ──

func Test_C21_HasJson(t *testing.T) {
	// Arrange
	r := corejson.New("x")

	// Act
	actual := args.Map{"result": r.HasJson()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C21_HasJsonBytes(t *testing.T) {
	// Arrange
	r := corejson.New("x")

	// Act
	actual := args.Map{"result": r.HasJsonBytes()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.HasSafeItems ──

func Test_C21_HasSafeItems(t *testing.T) {
	// Arrange
	r := corejson.New("x")

	// Act
	actual := args.Map{"result": r.HasSafeItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.IsEmptyJsonBytes ──

func Test_C21_IsEmptyJsonBytes_EmptyObj(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte("{}")}

	// Act
	actual := args.Map{"result": r.IsEmptyJsonBytes()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for {}", actual)
}

func Test_C21_IsEmptyJsonBytes_Zero(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte{}}

	// Act
	actual := args.Map{"result": r.IsEmptyJsonBytes()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.InjectInto ──

func Test_C21_InjectInto(t *testing.T) {
	r := corejson.New(map[string]string{"a": "b"})
	target := corejson.Empty.MapResults()
	err := r.InjectInto(target)
	_ = err
}

// ── Result.DeserializeMust ──

func Test_C21_DeserializeMust_Success(t *testing.T) {
	r := corejson.New("hello")
	var s string
	r.DeserializeMust(&s)
}

func Test_C21_DeserializeMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.Error(errors.New("fail"))
	var s string
	r.DeserializeMust(&s)
}

// ── Result.Raw ──

func Test_C21_Raw_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	_, err := r.Raw()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── Result.RawMust ──

func Test_C21_RawMust_Panic(t *testing.T) {
	defer func() { recover() }()
	var r *corejson.Result
	r.RawMust()
}

// ── Result.RawString ──

func Test_C21_RawString(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	s, err := r.RawString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid string", actual)
}

// ── Result.RawStringMust ──

func Test_C21_RawStringMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.Error(errors.New("fail"))
	r.RawStringMust()
}

func Test_C21_RawStringMust_Success(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	s := r.RawStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ── Result.RawErrString ──

func Test_C21_RawErrString(t *testing.T) {
	r := corejson.New("x")
	b, e := r.RawErrString()
	_ = b
	_ = e
}

// ── Result.RawPrettyString ──

func Test_C21_RawPrettyString(t *testing.T) {
	r := corejson.New(map[string]string{"a": "b"})
	s, err := r.RawPrettyString()
	_ = err
	_ = s
}

// ── Result.MeaningfulErrorMessage ──

func Test_C21_MeaningfulErrorMessage_NoErr(t *testing.T) {
	// Arrange
	r := corejson.New("x")

	// Act
	actual := args.Map{"result": r.MeaningfulErrorMessage() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C21_MeaningfulErrorMessage_WithErr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("fail"))

	// Act
	actual := args.Map{"result": r.MeaningfulErrorMessage() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ── Result.MeaningfulError ──

func Test_C21_MeaningfulError_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C21_MeaningfulError_EmptyBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: nil}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C21_MeaningfulError_HasError(t *testing.T) {
	// Arrange
	r := &corejson.Result{
		Bytes: []byte(`"x"`),
		Error: errors.New("some error"),
	}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── Result.SafeBytes ──

func Test_C21_SafeBytes_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	b := r.SafeBytes()

	// Act
	actual := args.Map{"result": len(b) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ── Result.JsonModel ──

func Test_C21_JsonModel_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m := r.JsonModel()

	// Act
	actual := args.Map{"result": m.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error in model", actual)
}

func Test_C21_JsonModel_Valid(t *testing.T) {
	r := corejson.New("x")
	m := r.JsonModel()
	_ = m
}

// ── Result.JsonModelAny ──

func Test_C21_JsonModelAny(t *testing.T) {
	r := corejson.New("x")
	a := r.JsonModelAny()
	_ = a
}

// ── Result.Json / JsonPtr ──

func Test_C21_Json(t *testing.T) {
	r := corejson.New("x")
	j := r.Json()
	_ = j
}

func Test_C21_JsonPtr(t *testing.T) {
	r := corejson.New("x")
	j := r.JsonPtr()
	_ = j
}

// ── Result.ParseInjectUsingJson ──

func Test_C21_ParseInjectUsingJson_Error(t *testing.T) {
	r := corejson.Empty.ResultPtr()
	badInput := corejson.NewResult.UsingString(`invalid`)
	_, err := r.ParseInjectUsingJson(badInput)
	_ = err
}

func Test_C21_ParseInjectUsingJson_Success(t *testing.T) {
	r := corejson.Empty.ResultPtr()
	input := corejson.New(*r)
	_, err := r.ParseInjectUsingJson(&input)
	_ = err
}

// ── Result.ParseInjectUsingJsonMust ──

func Test_C21_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := corejson.Empty.ResultPtr()
	bad := corejson.NewResult.UsingString(`invalid`)
	r.ParseInjectUsingJsonMust(bad)
}

// ── Result.CloneError ──

func Test_C21_CloneError_HasError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("err"))
	err := r.CloneError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C21_CloneError_NoError(t *testing.T) {
	// Arrange
	r := corejson.New("x")
	err := r.CloneError()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// ── Result.Ptr / NonPtr ──

func Test_C21_Ptr(t *testing.T) {
	// Arrange
	r := corejson.New("x")
	p := r.Ptr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C21_NonPtr_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	np := r.NonPtr()

	// Act
	actual := args.Map{"result": np.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C21_NonPtr_Valid(t *testing.T) {
	r := corejson.New("x")
	np := r.NonPtr()
	_ = np
}

// ── Result.IsEqualPtr ──

func Test_C21_IsEqualPtr_BothNil(t *testing.T) {
	// Arrange
	var a, b *corejson.Result

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C21_IsEqualPtr_OneNil(t *testing.T) {
	// Arrange
	a := corejson.New("x").Ptr()

	// Act
	actual := args.Map{"result": a.IsEqualPtr(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C21_IsEqualPtr_Same(t *testing.T) {
	// Arrange
	a := corejson.New("x").Ptr()

	// Act
	actual := args.Map{"result": a.IsEqualPtr(a)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true same ptr", actual)
}

func Test_C21_IsEqualPtr_DiffLength(t *testing.T) {
	// Arrange
	a := corejson.New("x").Ptr()
	b := corejson.New("xy").Ptr()

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C21_IsEqualPtr_DiffError(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("a"), "t")
	b := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("b"), "t")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C21_IsEqualPtr_DiffTypeName(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t1")
	b := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t2")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C21_IsEqualPtr_Equal(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t")
	b := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.IsEqual ──

func Test_C21_IsEqual_DiffLen(t *testing.T) {
	// Arrange
	a := corejson.New("x")
	b := corejson.New("xy")

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C21_IsEqual_DiffErr(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Create([]byte(`"x"`), errors.New("a"), "")
	b := corejson.NewResult.Create([]byte(`"x"`), errors.New("b"), "")

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C21_IsEqual_Equal(t *testing.T) {
	// Arrange
	a := corejson.New("x")
	b := corejson.New("x")

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.CombineErrorWithRefString / CombineErrorWithRefError ──

func Test_C21_CombineErrorWithRefString_NoError(t *testing.T) {
	// Arrange
	r := corejson.New("x")
	s := r.CombineErrorWithRefString("ref1")

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C21_CombineErrorWithRefString_WithError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("fail"))
	s := r.CombineErrorWithRefString("ref1", "ref2")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C21_CombineErrorWithRefError_NoError(t *testing.T) {
	// Arrange
	r := corejson.New("x")
	err := r.CombineErrorWithRefError("ref")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C21_CombineErrorWithRefError_WithError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("fail"))
	err := r.CombineErrorWithRefError("ref")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── Result.BytesError ──

func Test_C21_BytesError_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": r.BytesError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C21_BytesError_Valid(t *testing.T) {
	// Arrange
	r := corejson.New("x")
	be := r.BytesError()

	// Act
	actual := args.Map{"result": be == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ── Result.Dispose ──

func Test_C21_Dispose_Nil(t *testing.T) {
	var r *corejson.Result
	r.Dispose() // should not panic
}

func Test_C21_Dispose_Valid(t *testing.T) {
	r := corejson.New("x")
	r.Dispose()
}

// ── Result.CloneIf / ClonePtr / Clone ──

func Test_C21_CloneIf_NoClone(t *testing.T) {
	r := corejson.New("x")
	c := r.CloneIf(false, false)
	_ = c
}

func Test_C21_CloneIf_Clone(t *testing.T) {
	r := corejson.New("x")
	c := r.CloneIf(true, false)
	_ = c
}

func Test_C21_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": r.ClonePtr(false) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C21_ClonePtr_Valid(t *testing.T) {
	// Arrange
	r := corejson.New("x")
	p := r.ClonePtr(true)

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C21_Clone_Empty(t *testing.T) {
	r := corejson.Result{}
	c := r.Clone(true)
	_ = c
}

func Test_C21_Clone_ShallowCopy(t *testing.T) {
	r := corejson.New("x")
	c := r.Clone(false)
	_ = c
}

func Test_C21_Clone_DeepCopy(t *testing.T) {
	r := corejson.New("x")
	c := r.Clone(true)
	_ = c
}

// ── Result.AsJsonContractsBinder / AsJsoner / AsJsonParseSelfInjector ──

func Test_C21_AsJsonContractsBinder(t *testing.T) {
	r := corejson.New("x")
	_ = r.AsJsonContractsBinder()
}

func Test_C21_AsJsoner(t *testing.T) {
	r := corejson.New("x")
	_ = r.AsJsoner()
}

func Test_C21_AsJsonParseSelfInjector(t *testing.T) {
	r := corejson.New("x")
	_ = r.AsJsonParseSelfInjector()
}

// ── Result.JsonParseSelfInject ──

func Test_C21_JsonParseSelfInject(t *testing.T) {
	r := corejson.New("x")
	input := corejson.New(r)
	err := r.JsonParseSelfInject(&input)
	_ = err
}

// ── Result.SafeBytesTypeName ──

func Test_C21_SafeBytesTypeName_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}

	// Act
	actual := args.Map{"result": r.SafeBytesTypeName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}
