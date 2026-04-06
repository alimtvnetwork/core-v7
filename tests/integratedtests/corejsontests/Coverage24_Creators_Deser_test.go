package corejsontests

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ── New / NewPtr ──

func Test_C24_New_Valid(t *testing.T) {
	r := corejson.New("hello")
	if r.HasError() { t.Fatal("expected no error") }
}

func Test_C24_New_Error(t *testing.T) {
	r := corejson.New(make(chan int))
	if !r.HasError() { t.Fatal("expected error") }
}

func Test_C24_NewPtr_Valid(t *testing.T) {
	r := corejson.NewPtr("hello")
	if r.HasError() { t.Fatal("expected no error") }
}

func Test_C24_NewPtr_Error(t *testing.T) {
	r := corejson.NewPtr(make(chan int))
	if !r.HasError() { t.Fatal("expected error") }
}

// ── newResultCreator methods ──

func Test_C24_NRC_UnmarshalUsingBytes(t *testing.T) {
	original := corejson.New("test")
	b, _ := original.Serialize()
	r := corejson.NewResult.UnmarshalUsingBytes(b)
	_ = r
}

func Test_C24_NRC_DeserializeUsingBytes(t *testing.T) {
	r := corejson.NewResult.DeserializeUsingBytes([]byte(`{"Bytes":"dGVzdA==","TypeName":"t"}`))
	_ = r
}

func Test_C24_NRC_DeserializeUsingBytes_Error(t *testing.T) {
	r := corejson.NewResult.DeserializeUsingBytes([]byte(`invalid`))
	if r.Error == nil { t.Fatal("expected error") }
}

func Test_C24_NRC_DeserializeUsingResult_HasIssue(t *testing.T) {
	bad := &corejson.Result{}
	r := corejson.NewResult.DeserializeUsingResult(bad)
	_ = r
}

func Test_C24_NRC_DeserializeUsingResult_Valid(t *testing.T) {
	original := corejson.New("x")
	serialized := corejson.New(original)
	r := corejson.NewResult.DeserializeUsingResult(serialized.Ptr())
	_ = r
}

func Test_C24_NRC_UsingBytesType(t *testing.T) {
	r := corejson.NewResult.UsingBytesType([]byte(`"x"`), "TestType")
	if r.TypeName != "TestType" { t.Fatal("wrong type") }
}

func Test_C24_NRC_UsingBytesPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtr(nil)
	if r == nil { t.Fatal("expected non-nil") }
}

func Test_C24_NRC_UsingBytesPtr_Valid(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtr([]byte(`"x"`))
	if r == nil { t.Fatal("expected non-nil") }
}

func Test_C24_NRC_UsingBytesPtrErrPtr(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "t")
	if r.Error == nil { t.Fatal("expected error") }
	r2 := corejson.NewResult.UsingBytesPtrErrPtr([]byte(`"x"`), nil, "t")
	_ = r2
}

func Test_C24_NRC_UsingBytesErrPtr(t *testing.T) {
	r := corejson.NewResult.UsingBytesErrPtr(nil, errors.New("e"), "t")
	_ = r
	r2 := corejson.NewResult.UsingBytesErrPtr([]byte(`"x"`), nil, "t")
	_ = r2
}

func Test_C24_NRC_PtrUsingStringPtr_Nil(t *testing.T) {
	r := corejson.NewResult.PtrUsingStringPtr(nil, "t")
	if r.Error == nil { t.Fatal("expected error") }
}

func Test_C24_NRC_PtrUsingStringPtr_Valid(t *testing.T) {
	s := `"x"`
	r := corejson.NewResult.PtrUsingStringPtr(&s, "t")
	_ = r
}

func Test_C24_NRC_UsingErrorStringPtr(t *testing.T) {
	r := corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "t")
	_ = r
	s := `"x"`
	r2 := corejson.NewResult.UsingErrorStringPtr(nil, &s, "t")
	_ = r2
	r3 := corejson.NewResult.UsingErrorStringPtr(errors.New("e"), &s, "t")
	_ = r3
}

func Test_C24_NRC_Ptr(t *testing.T) {
	r := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t")
	_ = r
}

func Test_C24_NRC_UsingJsonBytesTypeError(t *testing.T) {
	r := corejson.NewResult.UsingJsonBytesTypeError([]byte(`"x"`), nil, "t")
	_ = r
}

func Test_C24_NRC_UsingJsonBytesError(t *testing.T) {
	r := corejson.NewResult.UsingJsonBytesError([]byte(`"x"`), nil)
	_ = r
}

func Test_C24_NRC_UsingTypePlusString(t *testing.T) {
	r := corejson.NewResult.UsingTypePlusString("t", `"x"`)
	_ = r
}

func Test_C24_NRC_UsingTypePlusStringPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingTypePlusStringPtr("t", nil)
	_ = r
}

func Test_C24_NRC_UsingTypePlusStringPtr_Empty(t *testing.T) {
	s := ""
	r := corejson.NewResult.UsingTypePlusStringPtr("t", &s)
	_ = r
}

func Test_C24_NRC_UsingTypePlusStringPtr_Valid(t *testing.T) {
	s := `"x"`
	r := corejson.NewResult.UsingTypePlusStringPtr("t", &s)
	_ = r
}

func Test_C24_NRC_UsingStringWithType(t *testing.T) {
	_ = corejson.NewResult.UsingStringWithType(`"x"`, "t")
}

func Test_C24_NRC_UsingString(t *testing.T) {
	_ = corejson.NewResult.UsingString(`"x"`)
}

func Test_C24_NRC_UsingStringPtr_Nil(t *testing.T) {
	_ = corejson.NewResult.UsingStringPtr(nil)
}

func Test_C24_NRC_UsingStringPtr_Empty(t *testing.T) {
	s := ""
	_ = corejson.NewResult.UsingStringPtr(&s)
}

func Test_C24_NRC_UsingStringPtr_Valid(t *testing.T) {
	s := `"x"`
	_ = corejson.NewResult.UsingStringPtr(&s)
}

func Test_C24_NRC_CreatePtr(t *testing.T) {
	_ = corejson.NewResult.CreatePtr([]byte(`"x"`), nil, "t")
}

func Test_C24_NRC_NonPtr(t *testing.T) {
	_ = corejson.NewResult.NonPtr([]byte(`"x"`), nil, "t")
}

func Test_C24_NRC_Create(t *testing.T) {
	_ = corejson.NewResult.Create([]byte(`"x"`), nil, "t")
}

func Test_C24_NRC_PtrUsingBytesPtr(t *testing.T) {
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "t")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, nil, "t")
	_ = corejson.NewResult.PtrUsingBytesPtr([]byte(`"x"`), nil, "t")
}

func Test_C24_NRC_CastingAny(t *testing.T) {
	_ = corejson.NewResult.CastingAny("hello")
}

func Test_C24_NRC_Any(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	_ = r
}

func Test_C24_NRC_Any_Error(t *testing.T) {
	r := corejson.NewResult.Any(make(chan int))
	if !r.HasError() { t.Fatal("expected error") }
}

func Test_C24_NRC_AnyPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	_ = r
}

func Test_C24_NRC_AnyPtr_Error(t *testing.T) {
	r := corejson.NewResult.AnyPtr(make(chan int))
	if !r.HasError() { t.Fatal("expected error") }
}

func Test_C24_NRC_UsingBytesError_Nil(t *testing.T) {
	r := corejson.NewResult.UsingBytesError(nil)
	_ = r
}

func Test_C24_NRC_Error(t *testing.T) {
	_ = corejson.NewResult.Error(errors.New("e"))
}

func Test_C24_NRC_ErrorPtr(t *testing.T) {
	_ = corejson.NewResult.ErrorPtr(errors.New("e"))
}

func Test_C24_NRC_Empty(t *testing.T) {
	_ = corejson.NewResult.Empty()
}

func Test_C24_NRC_EmptyPtr(t *testing.T) {
	_ = corejson.NewResult.EmptyPtr()
}

func Test_C24_NRC_TypeName(t *testing.T) {
	_ = corejson.NewResult.TypeName("t")
}

func Test_C24_NRC_TypeNameBytes(t *testing.T) {
	_ = corejson.NewResult.TypeNameBytes("t")
}

func Test_C24_NRC_Many(t *testing.T) {
	_ = corejson.NewResult.Many("a", "b")
}

func Test_C24_NRC_Serialize(t *testing.T) {
	r := corejson.NewResult.Serialize("hello")
	_ = r
}

func Test_C24_NRC_Serialize_Error(t *testing.T) {
	r := corejson.NewResult.Serialize(make(chan int))
	if !r.HasError() { t.Fatal("expected error") }
}

func Test_C24_NRC_Marshal(t *testing.T) {
	r := corejson.NewResult.Marshal("hello")
	_ = r
}

func Test_C24_NRC_Marshal_Error(t *testing.T) {
	r := corejson.NewResult.Marshal(make(chan int))
	if !r.HasError() { t.Fatal("expected error") }
}

func Test_C24_NRC_UsingSerializer_Nil(t *testing.T) {
	r := corejson.NewResult.UsingSerializer(nil)
	if r != nil { t.Fatal("expected nil") }
}

func Test_C24_NRC_UsingSerializerFunc_Nil(t *testing.T) {
	r := corejson.NewResult.UsingSerializerFunc(nil)
	if r != nil { t.Fatal("expected nil") }
}

func Test_C24_NRC_UsingSerializerFunc_Valid(t *testing.T) {
	r := corejson.NewResult.UsingSerializerFunc(func() ([]byte, error) {
		return json.Marshal("test")
	})
	if r == nil || r.HasError() { t.Fatal("unexpected") }
}

func Test_C24_NRC_UsingJsoner_Nil(t *testing.T) {
	r := corejson.NewResult.UsingJsoner(nil)
	if r != nil { t.Fatal("expected nil") }
}

func Test_C24_NRC_AnyToCastingResult(t *testing.T) {
	_ = corejson.NewResult.AnyToCastingResult("hello")
}

// ── emptyCreator ──

func Test_C24_Empty_Result(t *testing.T) { _ = corejson.Empty.Result() }
func Test_C24_Empty_ResultWithErr(t *testing.T) { _ = corejson.Empty.ResultWithErr("t", errors.New("e")) }
func Test_C24_Empty_BytesCollection(t *testing.T) { _ = corejson.Empty.BytesCollection() }
func Test_C24_Empty_BytesCollectionPtr(t *testing.T) { _ = corejson.Empty.BytesCollectionPtr() }
func Test_C24_Empty_ResultsCollection(t *testing.T) { _ = corejson.Empty.ResultsCollection() }
func Test_C24_Empty_ResultsPtrCollection(t *testing.T) { _ = corejson.Empty.ResultsPtrCollection() }
func Test_C24_Empty_MapResults(t *testing.T) { _ = corejson.Empty.MapResults() }

// ── BytesCloneIf ──

func Test_C24_BytesCloneIf_NoClone(t *testing.T) {
	b := corejson.BytesCloneIf(false, []byte("hello"))
	if len(b) != 0 { t.Fatal("expected empty for no clone") }
}

func Test_C24_BytesCloneIf_DeepClone(t *testing.T) {
	b := corejson.BytesCloneIf(true, []byte("hello"))
	if len(b) != 5 { t.Fatal("expected 5") }
}

func Test_C24_BytesCloneIf_Empty(t *testing.T) {
	b := corejson.BytesCloneIf(true, []byte{})
	if len(b) != 0 { t.Fatal("expected 0") }
}

// ── BytesToString / BytesToPrettyString ──

func Test_C24_BytesToString_Empty(t *testing.T) {
	if corejson.BytesToString(nil) != "" { t.Fatal("expected empty") }
}

func Test_C24_BytesToString_Valid(t *testing.T) {
	if corejson.BytesToString([]byte(`"x"`)) != `"x"` { t.Fatal("unexpected") }
}

func Test_C24_BytesToPrettyString_Empty(t *testing.T) {
	if corejson.BytesToPrettyString(nil) != "" { t.Fatal("expected empty") }
}

func Test_C24_BytesToPrettyString_Valid(t *testing.T) {
	s := corejson.BytesToPrettyString([]byte(`{"a":"b"}`))
	if s == "" { t.Fatal("expected non-empty") }
}

// ── JsonString / JsonStringOrErrMsg ──

func Test_C24_JsonString_Valid(t *testing.T) {
	s, err := corejson.JsonString("hello")
	if err != nil || s == "" { t.Fatal("unexpected") }
}

func Test_C24_JsonStringOrErrMsg_Valid(t *testing.T) {
	s := corejson.JsonStringOrErrMsg("hello")
	if s == "" { t.Fatal("expected non-empty") }
}

func Test_C24_JsonStringOrErrMsg_Error(t *testing.T) {
	s := corejson.JsonStringOrErrMsg(make(chan int))
	if s == "" { t.Fatal("expected error message") }
}
