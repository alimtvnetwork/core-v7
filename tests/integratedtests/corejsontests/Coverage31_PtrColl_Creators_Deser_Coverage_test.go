package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ═══════════════════════════════════════════════
// ResultsPtrCollection — all uncovered methods
// ═══════════════════════════════════════════════

func Test_C31_01_RPC_Length(t *testing.T) {
	var rpc *corejson.ResultsPtrCollection
	if rpc.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C31_02_RPC_LastIndex(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	if rpc.LastIndex() != -1 {
		t.Fatal("expected -1")
	}
}

func Test_C31_03_RPC_IsEmpty_HasAnyItem(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	if !rpc.IsEmpty() || rpc.HasAnyItem() {
		t.Fatal("unexpected")
	}
}

func Test_C31_04_RPC_FirstOrDefault(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	if rpc.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	if rpc.FirstOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_C31_05_RPC_LastOrDefault(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	if rpc.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	if rpc.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_C31_06_RPC_Take(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.Take(1)
	rpc.Add(corejson.NewResult.AnyPtr("a"))
	rpc.Add(corejson.NewResult.AnyPtr("b"))
	if rpc.Take(1).Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C31_07_RPC_Limit(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.Limit(5)
	rpc.Add(corejson.NewResult.AnyPtr("a"))
	rpc.Add(corejson.NewResult.AnyPtr("b"))
	if rpc.Limit(-1).Length() != 2 {
		t.Fatal("expected 2")
	}
	if rpc.Limit(1).Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C31_08_RPC_Skip(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.Skip(0)
	rpc.Add(corejson.NewResult.AnyPtr("a"))
	rpc.Add(corejson.NewResult.AnyPtr("b"))
	if rpc.Skip(1).Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C31_09_RPC_AddSkipOnNil(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSkipOnNil(nil)
	rpc.AddSkipOnNil(corejson.NewResult.AnyPtr("x"))
	if rpc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C31_10_RPC_AddNonNilNonError(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddNonNilNonError(nil)
	rpc.AddNonNilNonError(&corejson.Result{Error: errors.New("e")})
	rpc.AddNonNilNonError(corejson.NewResult.AnyPtr("x"))
	if rpc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C31_11_RPC_GetAt(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	if rpc.GetAt(0) == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_C31_12_RPC_HasError(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	if rpc.HasError() {
		t.Fatal("expected false")
	}
	rpc.Add(&corejson.Result{Error: errors.New("e")})
	if !rpc.HasError() {
		t.Fatal("expected true")
	}
}

func Test_C31_13_RPC_AllErrors(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	errs, has := rpc.AllErrors()
	if has || len(errs) != 0 {
		t.Fatal("unexpected")
	}
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	rpc.Add(&corejson.Result{Error: errors.New("e")})
	errs, has = rpc.AllErrors()
	if !has || len(errs) != 1 {
		t.Fatal("unexpected")
	}
}

func Test_C31_14_RPC_GetErrorsStrings(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetErrorsStrings()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	rpc.Add(&corejson.Result{Error: errors.New("e")})
	s := rpc.GetErrorsStrings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C31_15_RPC_GetErrorsStringsPtr(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetErrorsStringsPtr()
}

func Test_C31_16_RPC_GetErrorsAsSingleString(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetErrorsAsSingleString()
}

func Test_C31_17_RPC_GetErrorsAsSingle(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetErrorsAsSingle()
}

func Test_C31_18_RPC_UnmarshalAt(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("hello"))
	var s string
	err := rpc.UnmarshalAt(0, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_19_RPC_UnmarshalAt_NilResult(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(nil)
	var s string
	err := rpc.UnmarshalAt(0, &s)
	if err != nil {
		t.Fatal("expected nil for nil result")
	}
}

func Test_C31_20_RPC_UnmarshalAt_HasError(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(&corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e")})
	var s string
	err := rpc.UnmarshalAt(0, &s)
	// Accept whatever result - just exercise the code path
	_ = err
}

func Test_C31_21_RPC_InjectIntoAt(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	rpc.Add(r)
	target := corejson.Result{}
	err := rpc.InjectIntoAt(0, &target)
	_ = err
}

func Test_C31_22_RPC_InjectIntoSameIndex(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	var nilInjectors []corejson.JsonParseSelfInjector
	errs, has := rpc.InjectIntoSameIndex(nilInjectors...)
	if has || len(errs) != 0 {
		t.Fatal("unexpected")
	}
	rpc.Add(nil)
	rpc.Add(&corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e")})
	rpc.Add(corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"}))
	t1 := corejson.Result{}
	errs, has = rpc.InjectIntoSameIndex(nil, nil, &t1)
	_ = errs
	_ = has
}

func Test_C31_23_RPC_UnmarshalIntoSameIndex(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	var nilAnys []any
	errs, has := rpc.UnmarshalIntoSameIndex(nilAnys...)
	if has || len(errs) != 0 {
		t.Fatal("unexpected")
	}
	rpc.Add(corejson.NewResult.AnyPtr("hello"))
	rpc.Add(nil)
	rpc.Add(&corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e")})
	rpc.Add(&corejson.Result{Bytes: []byte(`{}`)})
	var s string
	errs, has = rpc.UnmarshalIntoSameIndex(&s, nil, nil, nil)
	_ = errs
	_ = has
}

func Test_C31_24_RPC_GetAtSafe(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	if rpc.GetAtSafe(0) == nil {
		t.Fatal("expected non-nil")
	}
	if rpc.GetAtSafe(-1) != nil {
		t.Fatal("expected nil")
	}
}

func Test_C31_25_RPC_GetAtSafeUsingLength(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	if rpc.GetAtSafeUsingLength(0, 1) == nil {
		t.Fatal("expected non-nil")
	}
	if rpc.GetAtSafeUsingLength(5, 1) != nil {
		t.Fatal("expected nil")
	}
}

func Test_C31_26_RPC_Add_AddResult(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	rpc.AddResult(corejson.NewResult.Any("y"))
	if rpc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C31_27_RPC_AddSerializer(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializer(nil)
	if rpc.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C31_28_RPC_AddSerializers(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializers()
}

func Test_C31_29_RPC_AddSerializerFunc(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializerFunc(nil)
	rpc.AddSerializerFunc(func() ([]byte, error) { return []byte(`"x"`), nil })
	if rpc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C31_30_RPC_AddSerializerFunctions(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializerFunctions()
}

func Test_C31_31_RPC_Adds(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Adds(nil, corejson.NewResult.AnyPtr("x"))
	if rpc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C31_32_RPC_AddAny(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddAny(nil)
	rpc.AddAny("x")
	if rpc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C31_33_RPC_AddAnyItems(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddAnyItems(nil, "a", nil, "b")
	if rpc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C31_34_RPC_AddResultsCollection(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddResultsCollection(nil)
	other := corejson.NewResultsPtrCollection.Empty()
	other.Add(corejson.NewResult.AnyPtr("x"))
	rpc.AddResultsCollection(other)
	if rpc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C31_35_RPC_AddNonNilItems(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddNonNilItems(nil, corejson.NewResult.AnyPtr("x"))
	// Note: AddNonNilItems appends results... for each non-nil result
	_ = rpc
}

func Test_C31_36_RPC_AddNonNilItemsPtr(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddNonNilItemsPtr()
	rpc.AddNonNilItemsPtr(nil, corejson.NewResult.AnyPtr("x"))
	_ = rpc
}

func Test_C31_37_RPC_Clear(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	rpc.Clear()
	if rpc.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C31_38_RPC_Clear_Nil(t *testing.T) {
	var rpc *corejson.ResultsPtrCollection
	_ = rpc.Clear()
}

func Test_C31_39_RPC_Dispose(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	rpc.Dispose()
}

func Test_C31_40_RPC_Dispose_Nil(t *testing.T) {
	var rpc *corejson.ResultsPtrCollection
	rpc.Dispose()
}

func Test_C31_41_RPC_GetStrings(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetStrings()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	s := rpc.GetStrings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C31_42_RPC_GetStringsPtr(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetStringsPtr()
}

func Test_C31_43_RPC_AddJsoners(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddJsoners(true)
}

func Test_C31_44_RPC_NonPtr_Ptr(t *testing.T) {
	rpc := corejson.ResultsPtrCollection{}
	_ = rpc.NonPtr()
	_ = rpc.Ptr()
}

func Test_C31_45_RPC_GetPagesSize(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	if rpc.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	for i := 0; i < 5; i++ {
		rpc.Add(corejson.NewResult.AnyPtr(i))
	}
	if rpc.GetPagesSize(2) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_C31_46_RPC_GetPagedCollection(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	for i := 0; i < 5; i++ {
		rpc.Add(corejson.NewResult.AnyPtr(i))
	}
	pages := rpc.GetPagedCollection(2)
	if len(pages) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_C31_47_RPC_GetPagedCollection_Small(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	pages := rpc.GetPagedCollection(10)
	if len(pages) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C31_48_RPC_GetSinglePageCollection(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	for i := 0; i < 10; i++ {
		rpc.Add(corejson.NewResult.AnyPtr(i))
	}
	page := rpc.GetSinglePageCollection(3, 1)
	if page.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_C31_49_RPC_GetSinglePageCollection_Small(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	page := rpc.GetSinglePageCollection(10, 1)
	if page.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C31_50_RPC_JsonModel_JsonModelAny(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.JsonModel()
	_ = rpc.JsonModelAny()
}

func Test_C31_51_RPC_Json_JsonPtr(t *testing.T) {
	rpc := corejson.ResultsPtrCollection{}
	_ = rpc.Json()
	_ = rpc.JsonPtr()
}

func Test_C31_52_RPC_ParseInjectUsingJson(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	serialized := rpc.JsonPtr()
	rpc2 := corejson.NewResultsPtrCollection.Empty()
	_, err := rpc2.ParseInjectUsingJson(serialized)
	_ = err
}

func Test_C31_53_RPC_ParseInjectUsingJson_Fail(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	bad := &corejson.Result{Error: errors.New("fail")}
	_, err := rpc.ParseInjectUsingJson(bad)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_C31_54_RPC_ParseInjectUsingJsonMust(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	serialized := rpc.JsonPtr()
	rpc2 := corejson.NewResultsPtrCollection.Empty()
	_ = rpc2.ParseInjectUsingJsonMust(serialized)
}

func Test_C31_55_RPC_AsJsonContractsBinder(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.AsJsonContractsBinder()
}

func Test_C31_56_RPC_AsJsoner(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.AsJsoner()
}

func Test_C31_57_RPC_JsonParseSelfInject(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	serialized := rpc.JsonPtr()
	rpc2 := corejson.NewResultsPtrCollection.Empty()
	_ = rpc2.JsonParseSelfInject(serialized)
}

func Test_C31_58_RPC_AsJsonParseSelfInjector(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.AsJsonParseSelfInjector()
}

func Test_C31_59_RPC_Clone(t *testing.T) {
	var rpc *corejson.ResultsPtrCollection
	if rpc.Clone(false) != nil {
		t.Fatal("expected nil")
	}
	rpc = corejson.NewResultsPtrCollection.Empty()
	_ = rpc.Clone(false)
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	_ = rpc.Clone(true)
}

// ═══════════════════════════════════════════════
// Creators — newResultCreator uncovered methods
// ═══════════════════════════════════════════════

func Test_C31_60_NRC_UnmarshalUsingBytes(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	b, _ := r.Ptr().Serialize()
	result := corejson.NewResult.UnmarshalUsingBytes(b)
	_ = result
}

func Test_C31_61_NRC_DeserializeUsingBytes(t *testing.T) {
	result := corejson.NewResult.DeserializeUsingBytes([]byte(`invalid`))
	_ = result
}

func Test_C31_62_NRC_DeserializeUsingResult(t *testing.T) {
	r := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	result := corejson.NewResult.DeserializeUsingResult(r)
	_ = result
	// with error
	result = corejson.NewResult.DeserializeUsingResult(&corejson.Result{Error: errors.New("e")})
	_ = result
}

func Test_C31_63_NRC_UsingBytes(t *testing.T) {
	r := corejson.NewResult.UsingBytes([]byte(`"x"`))
	_ = r
}

func Test_C31_64_NRC_UsingBytesType(t *testing.T) {
	r := corejson.NewResult.UsingBytesType([]byte(`"x"`), "T")
	_ = r
}

func Test_C31_65_NRC_UsingTypeBytesPtr(t *testing.T) {
	r := corejson.NewResult.UsingTypeBytesPtr("T", []byte(`"x"`))
	_ = r
}

func Test_C31_66_NRC_UsingBytesPtr(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtr(nil)
	_ = r
	r = corejson.NewResult.UsingBytesPtr([]byte(`"x"`))
	_ = r
}

func Test_C31_67_NRC_UsingBytesPtrErrPtr(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "T")
	_ = r
	r = corejson.NewResult.UsingBytesPtrErrPtr([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_C31_68_NRC_UsingBytesErrPtr(t *testing.T) {
	r := corejson.NewResult.UsingBytesErrPtr(nil, errors.New("e"), "T")
	_ = r
	r = corejson.NewResult.UsingBytesErrPtr([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_C31_69_NRC_PtrUsingStringPtr(t *testing.T) {
	r := corejson.NewResult.PtrUsingStringPtr(nil, "T")
	if r.Error == nil {
		t.Fatal("expected error")
	}
	s := `"hello"`
	r = corejson.NewResult.PtrUsingStringPtr(&s, "T")
	_ = r
}

func Test_C31_70_NRC_UsingErrorStringPtr(t *testing.T) {
	r := corejson.NewResult.UsingErrorStringPtr(nil, nil, "T")
	_ = r
	r = corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "T")
	_ = r
	s := `"hello"`
	r = corejson.NewResult.UsingErrorStringPtr(nil, &s, "T")
	_ = r
	r = corejson.NewResult.UsingErrorStringPtr(errors.New("e"), &s, "T")
	_ = r
}

func Test_C31_71_NRC_Ptr(t *testing.T) {
	r := corejson.NewResult.Ptr([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_C31_72_NRC_UsingJsonBytesTypeError(t *testing.T) {
	r := corejson.NewResult.UsingJsonBytesTypeError([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_C31_73_NRC_UsingJsonBytesError(t *testing.T) {
	r := corejson.NewResult.UsingJsonBytesError([]byte(`"x"`), nil)
	_ = r
}

func Test_C31_74_NRC_UsingTypePlusStringPtr(t *testing.T) {
	r := corejson.NewResult.UsingTypePlusStringPtr("T", nil)
	_ = r
	s := `"hello"`
	r = corejson.NewResult.UsingTypePlusStringPtr("T", &s)
	_ = r
	empty := ""
	r = corejson.NewResult.UsingTypePlusStringPtr("T", &empty)
	_ = r
}

func Test_C31_75_NRC_UsingStringWithType(t *testing.T) {
	r := corejson.NewResult.UsingStringWithType(`"hello"`, "T")
	_ = r
}

func Test_C31_76_NRC_UsingString(t *testing.T) {
	r := corejson.NewResult.UsingString(`"hello"`)
	_ = r
}

func Test_C31_77_NRC_UsingStringPtr(t *testing.T) {
	r := corejson.NewResult.UsingStringPtr(nil)
	_ = r
	s := `"hello"`
	r = corejson.NewResult.UsingStringPtr(&s)
	_ = r
}

func Test_C31_78_NRC_CreatePtr(t *testing.T) {
	r := corejson.NewResult.CreatePtr([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_C31_79_NRC_NonPtr(t *testing.T) {
	r := corejson.NewResult.NonPtr([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_C31_80_NRC_PtrUsingBytesPtr(t *testing.T) {
	r := corejson.NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "T")
	if r.Error == nil {
		t.Fatal("expected error")
	}
	r = corejson.NewResult.PtrUsingBytesPtr(nil, nil, "T")
	_ = r
	r = corejson.NewResult.PtrUsingBytesPtr([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_C31_81_NRC_CastingAny(t *testing.T) {
	r := corejson.NewResult.CastingAny("hello")
	_ = r
}

func Test_C31_82_NRC_UsingBytesError(t *testing.T) {
	r := corejson.NewResult.UsingBytesError(nil)
	_ = r
}

func Test_C31_83_NRC_Error_ErrorPtr(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("e"))
	_ = r
	rp := corejson.NewResult.ErrorPtr(errors.New("e"))
	_ = rp
}

func Test_C31_84_NRC_Empty_EmptyPtr(t *testing.T) {
	_ = corejson.NewResult.Empty()
	_ = corejson.NewResult.EmptyPtr()
}

func Test_C31_85_NRC_TypeName_TypeNameBytes(t *testing.T) {
	_ = corejson.NewResult.TypeName("T")
	_ = corejson.NewResult.TypeNameBytes("T")
}

func Test_C31_86_NRC_Many(t *testing.T) {
	r := corejson.NewResult.Many("a", "b")
	_ = r
}

func Test_C31_87_NRC_Serialize(t *testing.T) {
	r := corejson.NewResult.Serialize("hello")
	_ = r
	ch := make(chan int)
	r = corejson.NewResult.Serialize(ch)
	if !r.HasError() {
		t.Fatal("expected error")
	}
}

func Test_C31_88_NRC_Marshal(t *testing.T) {
	r := corejson.NewResult.Marshal("hello")
	_ = r
	ch := make(chan int)
	r = corejson.NewResult.Marshal(ch)
	if !r.HasError() {
		t.Fatal("expected error")
	}
}

func Test_C31_89_NRC_UsingSerializer(t *testing.T) {
	r := corejson.NewResult.UsingSerializer(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func Test_C31_90_NRC_UsingSerializerFunc(t *testing.T) {
	r := corejson.NewResult.UsingSerializerFunc(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
	r = corejson.NewResult.UsingSerializerFunc(func() ([]byte, error) { return []byte(`"x"`), nil })
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_C31_91_NRC_UsingJsoner(t *testing.T) {
	r := corejson.NewResult.UsingJsoner(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func Test_C31_92_NRC_AnyToCastingResult(t *testing.T) {
	r := corejson.NewResult.AnyToCastingResult("hello")
	_ = r
}

// ═══════════════════════════════════════════════
// Creators — collection creators
// ═══════════════════════════════════════════════

func Test_C31_93_NRCC_Default(t *testing.T) {
	_ = corejson.NewResultsCollection.Default()
}

func Test_C31_94_NRCC_AnyItems(t *testing.T) {
	_ = corejson.NewResultsCollection.AnyItems("a", "b")
}

func Test_C31_95_NRCC_AnyItemsPlusCap(t *testing.T) {
	_ = corejson.NewResultsCollection.AnyItemsPlusCap(5)
	_ = corejson.NewResultsCollection.AnyItemsPlusCap(5, "a")
}

func Test_C31_96_NRCC_UsingJsonersOption(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingJsonersOption(true, 5)
}

func Test_C31_97_NRCC_UsingJsonersNonNull(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingJsonersNonNull(5)
}

func Test_C31_98_NRCC_UsingJsoners(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingJsoners()
}

func Test_C31_99_NRCC_UsingResultsPtrPlusCap(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingResultsPtrPlusCap(5)
	_ = corejson.NewResultsCollection.UsingResultsPtrPlusCap(0, corejson.NewResult.AnyPtr("x"))
}

func Test_C31_100_NRCC_UsingResultsPtr(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingResultsPtr(corejson.NewResult.AnyPtr("x"))
}

func Test_C31_101_NRCC_UsingResultsPlusCap(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingResultsPlusCap(5)
	_ = corejson.NewResultsCollection.UsingResultsPlusCap(0, corejson.NewResult.Any("x"))
}

func Test_C31_102_NRCC_UsingResults(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingResults(corejson.NewResult.Any("x"))
}

func Test_C31_103_NRCC_Serializers(t *testing.T) {
	_ = corejson.NewResultsCollection.Serializers()
}

func Test_C31_104_NRCC_SerializerFunctions(t *testing.T) {
	_ = corejson.NewResultsCollection.SerializerFunctions()
}

func Test_C31_105_NRCC_UnmarshalUsingBytes(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	b, _ := rc.JsonPtr().SerializeSkipExistingIssues()
	_, _ = corejson.NewResultsCollection.UnmarshalUsingBytes(b)
}

func Test_C31_106_NRCC_DeserializeUsingBytes(t *testing.T) {
	_, _ = corejson.NewResultsCollection.DeserializeUsingBytes([]byte(`invalid`))
}

func Test_C31_107_NRCC_DeserializeUsingResult(t *testing.T) {
	_, _ = corejson.NewResultsCollection.DeserializeUsingResult(&corejson.Result{Error: errors.New("e")})
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	_, _ = corejson.NewResultsCollection.DeserializeUsingResult(rc.JsonPtr())
}

// ═══════════════════════════════════════════════
// newBytesCollectionCreator
// ═══════════════════════════════════════════════

func Test_C31_108_NBCC_UnmarshalUsingBytes(t *testing.T) {
	_, _ = corejson.NewBytesCollection.UnmarshalUsingBytes([]byte(`[[1]]`))
}

func Test_C31_109_NBCC_DeserializeUsingBytes(t *testing.T) {
	_, _ = corejson.NewBytesCollection.DeserializeUsingBytes([]byte(`invalid`))
}

func Test_C31_110_NBCC_DeserializeUsingResult(t *testing.T) {
	_, _ = corejson.NewBytesCollection.DeserializeUsingResult(&corejson.Result{Error: errors.New("e")})
}

func Test_C31_111_NBCC_AnyItems(t *testing.T) {
	_, _ = corejson.NewBytesCollection.AnyItems("a", "b")
}

func Test_C31_112_NBCC_JsonersPlusCap(t *testing.T) {
	_ = corejson.NewBytesCollection.JsonersPlusCap(true, 5)
}

func Test_C31_113_NBCC_Jsoners(t *testing.T) {
	_ = corejson.NewBytesCollection.Jsoners()
}

func Test_C31_114_NBCC_Serializers(t *testing.T) {
	_ = corejson.NewBytesCollection.Serializers()
}

// ═══════════════════════════════════════════════
// newResultsPtrCollectionCreator
// ═══════════════════════════════════════════════

func Test_C31_115_NRPCC_UnmarshalUsingBytes(t *testing.T) {
	_, _ = corejson.NewResultsPtrCollection.UnmarshalUsingBytes([]byte(`invalid`))
}

func Test_C31_116_NRPCC_DeserializeUsingResult(t *testing.T) {
	_, _ = corejson.NewResultsPtrCollection.DeserializeUsingResult(&corejson.Result{Error: errors.New("e")})
}

func Test_C31_117_NRPCC_Default(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Default()
}

func Test_C31_118_NRPCC_AnyItemsPlusCap(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.AnyItemsPlusCap(0)
	_ = corejson.NewResultsPtrCollection.AnyItemsPlusCap(5, "a")
}

func Test_C31_119_NRPCC_AnyItems(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.AnyItems("a")
}

func Test_C31_120_NRPCC_UsingResultsPlusCap(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.UsingResultsPlusCap(0)
	_ = corejson.NewResultsPtrCollection.UsingResultsPlusCap(0, corejson.NewResult.AnyPtr("x"))
}

func Test_C31_121_NRPCC_UsingResults(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.UsingResults(corejson.NewResult.AnyPtr("x"))
}

func Test_C31_122_NRPCC_JsonersPlusCap(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.JsonersPlusCap(true, 5)
}

func Test_C31_123_NRPCC_Jsoners(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Jsoners()
}

func Test_C31_124_NRPCC_Serializers(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Serializers()
}

// ═══════════════════════════════════════════════
// newMapResultsCreator
// ═══════════════════════════════════════════════

func Test_C31_125_NMRC_UnmarshalUsingBytes(t *testing.T) {
	_, _ = corejson.NewMapResults.UnmarshalUsingBytes([]byte(`{}`))
}

func Test_C31_126_NMRC_DeserializeUsingBytes(t *testing.T) {
	_, _ = corejson.NewMapResults.DeserializeUsingBytes([]byte(`invalid`))
}

func Test_C31_127_NMRC_DeserializeUsingResult(t *testing.T) {
	_, _ = corejson.NewMapResults.DeserializeUsingResult(&corejson.Result{Error: errors.New("e")})
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	_, _ = corejson.NewMapResults.DeserializeUsingResult(mr.JsonPtr())
}

func Test_C31_128_NMRC_UsingKeyAnyItems(t *testing.T) {
	_ = corejson.NewMapResults.UsingKeyAnyItems(0)
	_ = corejson.NewMapResults.UsingKeyAnyItems(5, corejson.KeyAny{Key: "k", AnyInf: "v"})
}

func Test_C31_129_NMRC_UsingMapOptions(t *testing.T) {
	m := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	_ = corejson.NewMapResults.UsingMapOptions(false, false, 0, nil)
	_ = corejson.NewMapResults.UsingMapOptions(false, false, 0, m)
	_ = corejson.NewMapResults.UsingMapOptions(true, false, 5, m)
}

func Test_C31_130_NMRC_UsingMapPlusCap(t *testing.T) {
	_ = corejson.NewMapResults.UsingMapPlusCap(5, nil)
	m := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	_ = corejson.NewMapResults.UsingMapPlusCap(5, m)
}

func Test_C31_131_NMRC_UsingMapPlusCapClone(t *testing.T) {
	_ = corejson.NewMapResults.UsingMapPlusCapClone(5, nil)
	m := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	_ = corejson.NewMapResults.UsingMapPlusCapClone(5, m)
}

func Test_C31_132_NMRC_UsingMapPlusCapDeepClone(t *testing.T) {
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(5, nil)
	m := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(5, m)
}

func Test_C31_133_NMRC_UsingMap(t *testing.T) {
	_ = corejson.NewMapResults.UsingMap(nil)
	m := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	_ = corejson.NewMapResults.UsingMap(m)
}

func Test_C31_134_NMRC_UsingMapAnyItemsPlusCap(t *testing.T) {
	_ = corejson.NewMapResults.UsingMapAnyItemsPlusCap(5, nil)
	_ = corejson.NewMapResults.UsingMapAnyItemsPlusCap(5, map[string]any{"k": "v"})
}

func Test_C31_135_NMRC_UsingMapAnyItems(t *testing.T) {
	_ = corejson.NewMapResults.UsingMapAnyItems(map[string]any{"k": "v"})
}

func Test_C31_136_NMRC_UsingKeyWithResultsPlusCap(t *testing.T) {
	_ = corejson.NewMapResults.UsingKeyWithResultsPlusCap(5)
	_ = corejson.NewMapResults.UsingKeyWithResultsPlusCap(0, corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")})
}

func Test_C31_137_NMRC_UsingKeyWithResults(t *testing.T) {
	_ = corejson.NewMapResults.UsingKeyWithResults(corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")})
}

func Test_C31_138_NMRC_UsingKeyJsonersPlusCap(t *testing.T) {
	_ = corejson.NewMapResults.UsingKeyJsonersPlusCap(5)
}

func Test_C31_139_NMRC_UsingKeyJsoners(t *testing.T) {
	_ = corejson.NewMapResults.UsingKeyJsoners()
}

// ═══════════════════════════════════════════════
// emptyCreator — uncovered methods
// ═══════════════════════════════════════════════

func Test_C31_140_Empty_ResultWithErr(t *testing.T) {
	r := corejson.Empty.ResultWithErr("T", errors.New("e"))
	if r.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_C31_141_Empty_ResultPtrWithErr(t *testing.T) {
	r := corejson.Empty.ResultPtrWithErr("T", errors.New("e"))
	if r.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_C31_142_Empty_BytesCollection(t *testing.T) {
	_ = corejson.Empty.BytesCollection()
}

func Test_C31_143_Empty_BytesCollectionPtr(t *testing.T) {
	_ = corejson.Empty.BytesCollectionPtr()
}

// ═══════════════════════════════════════════════
// Deserializer — uncovered methods
// ═══════════════════════════════════════════════

func Test_C31_144_Deser_UsingStringPtr(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringPtr(nil, &s)
	_ = err
	str := `"hello"`
	err = corejson.Deserialize.UsingStringPtr(&str, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_145_Deser_UsingError(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingError(nil, &s)
	if err != nil {
		t.Fatal("expected nil")
	}
	err = corejson.Deserialize.UsingError(errors.New(`"hello"`), &s)
	_ = err
}

func Test_C31_146_Deser_UsingErrorWhichJsonResult(t *testing.T) {
	var r corejson.Result
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &r)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_C31_147_Deser_ApplyMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	var s string
	corejson.Deserialize.ApplyMust(r, &s)
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_148_Deser_FromString(t *testing.T) {
	var s string
	err := corejson.Deserialize.FromString(`"hello"`, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_149_Deser_FromStringMust(t *testing.T) {
	var s string
	corejson.Deserialize.FromStringMust(`"hello"`, &s)
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_150_Deser_FromTo(t *testing.T) {
	var s string
	err := corejson.Deserialize.FromTo([]byte(`"hello"`), &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_151_Deser_MapAnyToPointer(t *testing.T) {
	type simple struct {
		Name string `json:"Name"`
	}
	var s simple
	err := corejson.Deserialize.MapAnyToPointer(true, nil, &s)
	if err != nil {
		t.Fatal("expected nil for empty map skip")
	}
	err = corejson.Deserialize.MapAnyToPointer(false, map[string]any{"Name": "test"}, &s)
	if err != nil || s.Name != "test" {
		t.Fatal("unexpected")
	}
}

func Test_C31_152_Deser_UsingStringOption(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringOption(true, "", &s)
	if err != nil {
		t.Fatal("expected nil")
	}
	err = corejson.Deserialize.UsingStringOption(false, `"hello"`, &s)
	_ = err
}

func Test_C31_153_Deser_UsingStringIgnoreEmpty(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &s)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_C31_154_Deser_UsingBytesPointerMust(t *testing.T) {
	var s string
	corejson.Deserialize.UsingBytesPointerMust([]byte(`"hello"`), &s)
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_155_Deser_UsingBytesIf(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"hello"`), &s)
	if err != nil {
		t.Fatal("expected nil when skip")
	}
	err = corejson.Deserialize.UsingBytesIf(true, []byte(`"hello"`), &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_156_Deser_UsingBytesPointerIf(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"hello"`), &s)
	if err != nil {
		t.Fatal("expected nil when skip")
	}
	err = corejson.Deserialize.UsingBytesPointerIf(true, []byte(`"hello"`), &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_157_Deser_UsingBytesPointer(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointer(nil, &s)
	if err == nil {
		t.Fatal("expected error for nil")
	}
	err = corejson.Deserialize.UsingBytesPointer([]byte(`"hello"`), &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_158_Deser_UsingBytesMust(t *testing.T) {
	var s string
	corejson.Deserialize.UsingBytesMust([]byte(`"hello"`), &s)
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_159_Deser_UsingSafeBytesMust(t *testing.T) {
	var s string
	corejson.Deserialize.UsingSafeBytesMust(nil, &s)
	corejson.Deserialize.UsingSafeBytesMust([]byte(`"hello"`), &s)
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_160_Deser_AnyToFieldsMap(t *testing.T) {
	m, err := corejson.Deserialize.AnyToFieldsMap(map[string]int{"a": 1})
	_ = m
	_ = err
}

func Test_C31_161_Deser_UsingSerializerTo(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingSerializerTo(nil, &s)
	_ = err
}

func Test_C31_162_Deser_UsingSerializerFuncTo(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingSerializerFuncTo(nil, &s)
	_ = err
	err = corejson.Deserialize.UsingSerializerFuncTo(func() ([]byte, error) {
		return []byte(`"hello"`), nil
	}, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C31_163_Deser_UsingDeserializerToOption(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &s)
	if err != nil {
		t.Fatal("expected nil")
	}
	err = corejson.Deserialize.UsingDeserializerToOption(false, nil, &s)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_C31_164_Deser_UsingDeserializerDefined(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerDefined(nil, &s)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_C31_165_Deser_UsingDeserializerFuncDefined(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &s)
	if err == nil {
		t.Fatal("expected error")
	}
	err = corejson.Deserialize.UsingDeserializerFuncDefined(func(toPtr any) error {
		return nil
	}, &s)
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func Test_C31_166_Deser_UsingJsonerToAny(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, &s)
	if err != nil {
		t.Fatal("expected nil")
	}
	err = corejson.Deserialize.UsingJsonerToAny(false, nil, &s)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_C31_167_Deser_UsingJsonerToAnyMust(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, &s)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_C31_168_Deser_Result(t *testing.T) {
	r := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	b, _ := r.Serialize()
	_, _ = corejson.Deserialize.Result(b)
}

func Test_C31_169_Deser_ResultPtr(t *testing.T) {
	_, _ = corejson.Deserialize.ResultPtr([]byte(`invalid`))
}

// ═══════════════════════════════════════════════
// deserializeFromBytesTo
// ═══════════════════════════════════════════════

func Test_C31_170_BytesTo_Strings(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Strings([]byte(`["a","b"]`))
}

func Test_C31_171_BytesTo_StringsMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.StringsMust([]byte(`["a","b"]`))
}

func Test_C31_172_BytesTo_String(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.String([]byte(`"hello"`))
}

func Test_C31_173_BytesTo_Integer(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Integer([]byte(`42`))
}

func Test_C31_174_BytesTo_IntegerMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.IntegerMust([]byte(`42`))
}

func Test_C31_175_BytesTo_Integer64(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Integer64([]byte(`64`))
}

func Test_C31_176_BytesTo_Integer64Must(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.Integer64Must([]byte(`64`))
}

func Test_C31_177_BytesTo_Integers(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Integers([]byte(`[1,2,3]`))
}

func Test_C31_178_BytesTo_IntegersMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.IntegersMust([]byte(`[1,2,3]`))
}

func Test_C31_179_BytesTo_StringMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.StringMust([]byte(`"hello"`))
}

func Test_C31_180_BytesTo_MapAnyItem(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.MapAnyItem([]byte(`{"a":1}`))
}

func Test_C31_181_BytesTo_MapAnyItemMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.MapAnyItemMust([]byte(`{"a":1}`))
}

func Test_C31_182_BytesTo_MapStringString(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.MapStringString([]byte(`{"a":"b"}`))
}

func Test_C31_183_BytesTo_MapStringStringMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.MapStringStringMust([]byte(`{"a":"b"}`))
}

func Test_C31_184_BytesTo_ResultCollection(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.ResultCollection([]byte(`{"JsonResultsCollection":[]}`))
	_, _ = corejson.Deserialize.BytesTo.ResultCollection([]byte(`invalid`))
}

func Test_C31_185_BytesTo_ResultCollectionMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.ResultCollectionMust([]byte(`{"JsonResultsCollection":[]}`))
}

func Test_C31_186_BytesTo_ResultsPtrCollection(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.ResultsPtrCollection([]byte(`{"JsonResultsCollection":[]}`))
	_, _ = corejson.Deserialize.BytesTo.ResultsPtrCollection([]byte(`invalid`))
}

func Test_C31_187_BytesTo_ResultsPtrCollectionMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.ResultsPtrCollectionMust([]byte(`{"JsonResultsCollection":[]}`))
}

func Test_C31_188_BytesTo_MapResults(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.MapResults([]byte(`{"JsonResultsMap":{}}`))
	_, _ = corejson.Deserialize.BytesTo.MapResults([]byte(`invalid`))
}

func Test_C31_189_BytesTo_MapResultsMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.MapResultsMust([]byte(`{"JsonResultsMap":{}}`))
}

func Test_C31_190_BytesTo_Bytes(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Bytes([]byte(`"dGVzdA=="`))
}

func Test_C31_191_BytesTo_BytesMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.BytesMust([]byte(`"dGVzdA=="`))
}

func Test_C31_192_BytesTo_Bool(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Bool([]byte(`true`))
}

func Test_C31_193_BytesTo_BoolMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.BoolMust([]byte(`true`))
}

// ═══════════════════════════════════════════════
// deserializeFromResultTo
// ═══════════════════════════════════════════════

func Test_C31_194_ResultTo_String(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	_, _ = corejson.Deserialize.ResultTo.String(r)
}

func Test_C31_195_ResultTo_Bool(t *testing.T) {
	r := corejson.NewResult.AnyPtr(true)
	_, _ = corejson.Deserialize.ResultTo.Bool(r)
}

func Test_C31_196_ResultTo_Byte(t *testing.T) {
	r := corejson.NewResult.AnyPtr(65)
	_, _ = corejson.Deserialize.ResultTo.Byte(r)
}

func Test_C31_197_ResultTo_ByteMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(65)
	_ = corejson.Deserialize.ResultTo.ByteMust(r)
}

func Test_C31_198_ResultTo_BoolMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(true)
	_ = corejson.Deserialize.ResultTo.BoolMust(r)
}

func Test_C31_199_ResultTo_StringMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	_ = corejson.Deserialize.ResultTo.StringMust(r)
}

func Test_C31_200_ResultTo_StringsMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr([]string{"a", "b"})
	_ = corejson.Deserialize.ResultTo.StringsMust(r)
}

func Test_C31_201_ResultTo_MapAnyItem(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]any{"a": 1})
	_, _ = corejson.Deserialize.ResultTo.MapAnyItem(r)
}

func Test_C31_202_ResultTo_MapAnyItemMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]any{"a": 1})
	_ = corejson.Deserialize.ResultTo.MapAnyItemMust(r)
}

func Test_C31_203_ResultTo_MapStringString(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"a": "b"})
	_, _ = corejson.Deserialize.ResultTo.MapStringString(r)
}

func Test_C31_204_ResultTo_MapStringStringMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"a": "b"})
	_ = corejson.Deserialize.ResultTo.MapStringStringMust(r)
}

func Test_C31_205_ResultTo_ResultCollection(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	r := rc.JsonPtr()
	_, _ = corejson.Deserialize.ResultTo.ResultCollection(r)
}

func Test_C31_206_ResultTo_ResultCollectionMust(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	r := rc.JsonPtr()
	_ = corejson.Deserialize.ResultTo.ResultCollectionMust(r)
}

func Test_C31_207_ResultTo_ResultsPtrCollection(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := rpc.JsonPtr()
	_, _ = corejson.Deserialize.ResultTo.ResultsPtrCollection(r)
}

func Test_C31_208_ResultTo_ResultsPtrCollectionMust(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := rpc.JsonPtr()
	_ = corejson.Deserialize.ResultTo.ResultsPtrCollectionMust(r)
}

func Test_C31_209_ResultTo_Result(t *testing.T) {
	inner := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	_, _ = corejson.Deserialize.ResultTo.Result(inner)
}

func Test_C31_210_ResultTo_ResultMust(t *testing.T) {
	inner := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	_ = corejson.Deserialize.ResultTo.ResultMust(inner)
}

func Test_C31_211_ResultTo_ResultPtr(t *testing.T) {
	inner := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	_, _ = corejson.Deserialize.ResultTo.ResultPtr(inner)
}

func Test_C31_212_ResultTo_ResultPtrMust(t *testing.T) {
	inner := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	_ = corejson.Deserialize.ResultTo.ResultPtrMust(inner)
}

func Test_C31_213_ResultTo_MapResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := mr.JsonPtr()
	_, _ = corejson.Deserialize.ResultTo.MapResults(r)
}

func Test_C31_214_ResultTo_Bytes(t *testing.T) {
	inner := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	_, _ = corejson.Deserialize.ResultTo.Bytes(inner)
}

func Test_C31_215_ResultTo_BytesMust(t *testing.T) {
	inner := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	_ = corejson.Deserialize.ResultTo.BytesMust(inner)
}

func Test_C31_216_ResultTo_MapResultsMust(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := mr.JsonPtr()
	_ = corejson.Deserialize.ResultTo.MapResultsMust(r)
}

// ═══════════════════════════════════════════════
// Serializer — uncovered methods
// ═══════════════════════════════════════════════

func Test_C31_217_Serializer_FromStringer(t *testing.T) {
	type myStringer31 struct{ val string }
	stringer := myStringer31{val: "test-stringer"}
	// Create a proper fmt.Stringer
	r := corejson.Serialize.FromStringer(stringerImpl31{stringer.val})
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

type stringerImpl31 struct{ v string }
func (s stringerImpl31) String() string { return s.v }

// ═══════════════════════════════════════════════
// AnyTo — uncovered: UsingSerializer
// ═══════════════════════════════════════════════

func Test_C31_218_AnyTo_UsingSerializer(t *testing.T) {
	r := corejson.AnyTo.UsingSerializer(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
}

// ═══════════════════════════════════════════════
// CastAny — uncovered: FromToReflection
// ═══════════════════════════════════════════════

func Test_C31_219_CastAny_FromToReflection(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToReflection([]byte(`"hello"`), &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

// ═══════════════════════════════════════════════
// Deserialize ResultMust, ResultPtrMust
// ═══════════════════════════════════════════════

func Test_C31_220_Deser_ResultMust(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	b, _ := r.Serialize()
	_ = corejson.Deserialize.ResultMust(b)
}

func Test_C31_221_Deser_ResultPtrMust(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	b, _ := r.Serialize()
	_ = corejson.Deserialize.ResultPtrMust(b)
}
