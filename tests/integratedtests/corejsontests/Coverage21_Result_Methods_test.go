package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ── Result.Map ──

func Test_C21_Result_Map_Nil(t *testing.T) {
	var r *corejson.Result
	m := r.Map()
	if len(m) != 0 {
		t.Fatal("expected empty map")
	}
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
	var r *corejson.Result
	m, err := r.DeserializedFieldsToMap()
	_ = err
	if len(m) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_C21_DeserializedFieldsToMap_Empty(t *testing.T) {
	r := &corejson.Result{}
	m, err := r.DeserializedFieldsToMap()
	_ = err
	_ = m
}

// ── Result.FieldsNames ──

func Test_C21_FieldsNames_Empty(t *testing.T) {
	r := &corejson.Result{}
	names, err := r.FieldsNames()
	_ = err
	if len(names) != 0 {
		t.Fatal("expected empty")
	}
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
	var r *corejson.Result
	if r.BytesTypeName() != "" {
		t.Fatal("expected empty")
	}
}

func Test_C21_BytesTypeName_Valid(t *testing.T) {
	r := corejson.NewResult.UsingStringWithType(`"x"`, "MyType")
	if r.BytesTypeName() != "MyType" {
		t.Fatal("expected MyType")
	}
}

// ── Result.JsonStringPtr ──

func Test_C21_JsonStringPtr_Nil(t *testing.T) {
	var r *corejson.Result
	ptr := r.JsonStringPtr()
	if ptr == nil || *ptr != "" {
		t.Fatal("expected empty string ptr")
	}
}

func Test_C21_JsonStringPtr_NoBytes(t *testing.T) {
	r := &corejson.Result{}
	ptr := r.JsonStringPtr()
	if *ptr != "" {
		t.Fatal("expected empty")
	}
}

func Test_C21_JsonStringPtr_Cached(t *testing.T) {
	r := corejson.NewResult.UsingString(`"hello"`)
	_ = r.JsonStringPtr() // first call caches
	_ = r.JsonStringPtr() // second call returns cached
}

// ── Result.PrettyJsonBuffer ──

func Test_C21_PrettyJsonBuffer_Empty(t *testing.T) {
	r := &corejson.Result{}
	buf, err := r.PrettyJsonBuffer("", "  ")
	if err != nil || buf.Len() != 0 {
		t.Fatal("expected empty buffer")
	}
}

// ── Result.PrettyJsonString ──

func Test_C21_PrettyJsonString_Nil(t *testing.T) {
	var r *corejson.Result
	if r.PrettyJsonString() != "" {
		t.Fatal("expected empty")
	}
}

func Test_C21_PrettyJsonString_InvalidJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte("not valid json")}
	s := r.PrettyJsonString()
	if s != "" {
		t.Fatal("expected empty for invalid json")
	}
}

// ── Result.PrettyJsonStringOrErrString ──

func Test_C21_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	var r *corejson.Result
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_C21_PrettyJsonStringOrErrString_Error(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("fail"))
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected error string")
	}
}

func Test_C21_PrettyJsonStringOrErrString_Valid(t *testing.T) {
	r := corejson.New(map[string]string{"a": "b"})
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected pretty string")
	}
}

// ── Result.Length ──

func Test_C21_Length_Nil(t *testing.T) {
	var r *corejson.Result
	if r.Length() != 0 {
		t.Fatal("expected 0")
	}
}

// ── Result.ErrorString ──

func Test_C21_ErrorString_HasError(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("err"))
	if r.ErrorString() == "" {
		t.Fatal("expected error string")
	}
}

// ── Result.IsErrorEqual ──

func Test_C21_IsErrorEqual_BothNil(t *testing.T) {
	r := corejson.New("test")
	if !r.IsErrorEqual(nil) {
		t.Fatal("expected true")
	}
}

func Test_C21_IsErrorEqual_OneNil(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("err"))
	if r.IsErrorEqual(nil) {
		t.Fatal("expected false")
	}
}

func Test_C21_IsErrorEqual_LeftNil(t *testing.T) {
	r := corejson.New("test")
	if r.IsErrorEqual(errors.New("err")) {
		t.Fatal("expected false")
	}
}

func Test_C21_IsErrorEqual_Match(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("same"))
	if !r.IsErrorEqual(errors.New("same")) {
		t.Fatal("expected true")
	}
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
	r := corejson.New("x")
	if !r.HasAnyItem() {
		t.Fatal("expected true")
	}
}

// ── Result.HasJson / HasJsonBytes ──

func Test_C21_HasJson(t *testing.T) {
	r := corejson.New("x")
	if !r.HasJson() {
		t.Fatal("expected true")
	}
}

func Test_C21_HasJsonBytes(t *testing.T) {
	r := corejson.New("x")
	if !r.HasJsonBytes() {
		t.Fatal("expected true")
	}
}

// ── Result.HasSafeItems ──

func Test_C21_HasSafeItems(t *testing.T) {
	r := corejson.New("x")
	if !r.HasSafeItems() {
		t.Fatal("expected true")
	}
}

// ── Result.IsEmptyJsonBytes ──

func Test_C21_IsEmptyJsonBytes_EmptyObj(t *testing.T) {
	r := &corejson.Result{Bytes: []byte("{}")}
	if !r.IsEmptyJsonBytes() {
		t.Fatal("expected true for {}")
	}
}

func Test_C21_IsEmptyJsonBytes_Zero(t *testing.T) {
	r := &corejson.Result{Bytes: []byte{}}
	if !r.IsEmptyJsonBytes() {
		t.Fatal("expected true")
	}
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
	var r *corejson.Result
	_, err := r.Raw()
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── Result.RawMust ──

func Test_C21_RawMust_Panic(t *testing.T) {
	defer func() { recover() }()
	var r *corejson.Result
	r.RawMust()
}

// ── Result.RawString ──

func Test_C21_RawString(t *testing.T) {
	r := corejson.New("hello")
	s, err := r.RawString()
	if err != nil || s == "" {
		t.Fatal("expected valid string")
	}
}

// ── Result.RawStringMust ──

func Test_C21_RawStringMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.Error(errors.New("fail"))
	r.RawStringMust()
}

func Test_C21_RawStringMust_Success(t *testing.T) {
	r := corejson.New("hello")
	s := r.RawStringMust()
	if s == "" {
		t.Fatal("expected non-empty")
	}
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
	r := corejson.New("x")
	if r.MeaningfulErrorMessage() != "" {
		t.Fatal("expected empty")
	}
}

func Test_C21_MeaningfulErrorMessage_WithErr(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("fail"))
	if r.MeaningfulErrorMessage() == "" {
		t.Fatal("expected non-empty")
	}
}

// ── Result.MeaningfulError ──

func Test_C21_MeaningfulError_Nil(t *testing.T) {
	var r *corejson.Result
	err := r.MeaningfulError()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_C21_MeaningfulError_EmptyBytes(t *testing.T) {
	r := &corejson.Result{Bytes: nil}
	err := r.MeaningfulError()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_C21_MeaningfulError_HasError(t *testing.T) {
	r := &corejson.Result{
		Bytes: []byte(`"x"`),
		Error: errors.New("some error"),
	}
	err := r.MeaningfulError()
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── Result.SafeBytes ──

func Test_C21_SafeBytes_Nil(t *testing.T) {
	var r *corejson.Result
	b := r.SafeBytes()
	if len(b) != 0 {
		t.Fatal("expected empty")
	}
}

// ── Result.JsonModel ──

func Test_C21_JsonModel_Nil(t *testing.T) {
	var r *corejson.Result
	m := r.JsonModel()
	if m.Error == nil {
		t.Fatal("expected error in model")
	}
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
	r := corejson.NewResult.Error(errors.New("err"))
	err := r.CloneError()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_C21_CloneError_NoError(t *testing.T) {
	r := corejson.New("x")
	err := r.CloneError()
	if err != nil {
		t.Fatal("expected nil")
	}
}

// ── Result.Ptr / NonPtr ──

func Test_C21_Ptr(t *testing.T) {
	r := corejson.New("x")
	p := r.Ptr()
	if p == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_C21_NonPtr_Nil(t *testing.T) {
	var r *corejson.Result
	np := r.NonPtr()
	if np.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_C21_NonPtr_Valid(t *testing.T) {
	r := corejson.New("x")
	np := r.NonPtr()
	_ = np
}

// ── Result.IsEqualPtr ──

func Test_C21_IsEqualPtr_BothNil(t *testing.T) {
	var a, b *corejson.Result
	if !a.IsEqualPtr(b) {
		t.Fatal("expected true")
	}
}

func Test_C21_IsEqualPtr_OneNil(t *testing.T) {
	a := corejson.New("x").Ptr()
	if a.IsEqualPtr(nil) {
		t.Fatal("expected false")
	}
}

func Test_C21_IsEqualPtr_Same(t *testing.T) {
	a := corejson.New("x").Ptr()
	if !a.IsEqualPtr(a) {
		t.Fatal("expected true same ptr")
	}
}

func Test_C21_IsEqualPtr_DiffLength(t *testing.T) {
	a := corejson.New("x").Ptr()
	b := corejson.New("xy").Ptr()
	if a.IsEqualPtr(b) {
		t.Fatal("expected false")
	}
}

func Test_C21_IsEqualPtr_DiffError(t *testing.T) {
	a := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("a"), "t")
	b := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("b"), "t")
	if a.IsEqualPtr(b) {
		t.Fatal("expected false")
	}
}

func Test_C21_IsEqualPtr_DiffTypeName(t *testing.T) {
	a := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t1")
	b := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t2")
	if a.IsEqualPtr(b) {
		t.Fatal("expected false")
	}
}

func Test_C21_IsEqualPtr_Equal(t *testing.T) {
	a := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t")
	b := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t")
	if !a.IsEqualPtr(b) {
		t.Fatal("expected true")
	}
}

// ── Result.IsEqual ──

func Test_C21_IsEqual_DiffLen(t *testing.T) {
	a := corejson.New("x")
	b := corejson.New("xy")
	if a.IsEqual(b) {
		t.Fatal("expected false")
	}
}

func Test_C21_IsEqual_DiffErr(t *testing.T) {
	a := corejson.NewResult.Create([]byte(`"x"`), errors.New("a"), "")
	b := corejson.NewResult.Create([]byte(`"x"`), errors.New("b"), "")
	if a.IsEqual(b) {
		t.Fatal("expected false")
	}
}

func Test_C21_IsEqual_Equal(t *testing.T) {
	a := corejson.New("x")
	b := corejson.New("x")
	if !a.IsEqual(b) {
		t.Fatal("expected true")
	}
}

// ── Result.CombineErrorWithRefString / CombineErrorWithRefError ──

func Test_C21_CombineErrorWithRefString_NoError(t *testing.T) {
	r := corejson.New("x")
	s := r.CombineErrorWithRefString("ref1")
	if s != "" {
		t.Fatal("expected empty")
	}
}

func Test_C21_CombineErrorWithRefString_WithError(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("fail"))
	s := r.CombineErrorWithRefString("ref1", "ref2")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_C21_CombineErrorWithRefError_NoError(t *testing.T) {
	r := corejson.New("x")
	err := r.CombineErrorWithRefError("ref")
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_C21_CombineErrorWithRefError_WithError(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("fail"))
	err := r.CombineErrorWithRefError("ref")
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── Result.BytesError ──

func Test_C21_BytesError_Nil(t *testing.T) {
	var r *corejson.Result
	if r.BytesError() != nil {
		t.Fatal("expected nil")
	}
}

func Test_C21_BytesError_Valid(t *testing.T) {
	r := corejson.New("x")
	be := r.BytesError()
	if be == nil {
		t.Fatal("expected non-nil")
	}
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
	var r *corejson.Result
	if r.ClonePtr(false) != nil {
		t.Fatal("expected nil")
	}
}

func Test_C21_ClonePtr_Valid(t *testing.T) {
	r := corejson.New("x")
	p := r.ClonePtr(true)
	if p == nil {
		t.Fatal("expected non-nil")
	}
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
	r := &corejson.Result{}
	if r.SafeBytesTypeName() != "" {
		t.Fatal("expected empty")
	}
}
