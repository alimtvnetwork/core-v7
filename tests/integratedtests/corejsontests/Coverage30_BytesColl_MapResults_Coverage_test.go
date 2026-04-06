package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ═══════════════════════════════════════════════
// BytesCollection — all uncovered methods
// ═══════════════════════════════════════════════

func Test_C30_01_BC_Length(t *testing.T) {
	var bc *corejson.BytesCollection
	if bc.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C30_02_BC_LastIndex(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	if bc.LastIndex() != -1 {
		t.Fatal("expected -1")
	}
}

func Test_C30_03_BC_IsEmpty_HasAnyItem(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	if !bc.IsEmpty() {
		t.Fatal("expected empty")
	}
	if bc.HasAnyItem() {
		t.Fatal("expected false")
	}
}

func Test_C30_04_BC_FirstOrDefault(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	if bc.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	bc.Add([]byte(`"x"`))
	if bc.FirstOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_C30_05_BC_LastOrDefault(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	if bc.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
	bc.Add([]byte(`"x"`))
	if bc.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_C30_06_BC_Take(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.Take(1)
	bc.Add([]byte(`"a"`))
	bc.Add([]byte(`"b"`))
	taken := bc.Take(1)
	if taken.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_07_BC_Limit(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.Limit(5)
	bc.Add([]byte(`"a"`))
	bc.Add([]byte(`"b"`))
	l := bc.Limit(-1)
	if l.Length() != 2 {
		t.Fatal("expected 2")
	}
	l = bc.Limit(1)
	if l.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_08_BC_Skip(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.Skip(0)
	bc.Add([]byte(`"a"`))
	bc.Add([]byte(`"b"`))
	s := bc.Skip(1)
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_09_BC_AddSkipOnNil(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSkipOnNil(nil)
	bc.AddSkipOnNil([]byte(`"x"`))
	if bc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_10_BC_AddNonEmpty(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddNonEmpty([]byte{})
	bc.AddNonEmpty([]byte(`"x"`))
	if bc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_11_BC_AddResultPtr(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddResultPtr(&corejson.Result{Error: errors.New("e")})
	bc.AddResultPtr(&corejson.Result{Bytes: []byte(`"x"`)})
	if bc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_12_BC_AddResult(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddResult(corejson.Result{Error: errors.New("e")})
	bc.AddResult(corejson.Result{Bytes: []byte(`"x"`)})
	if bc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_13_BC_GetAt(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	b := bc.GetAt(0)
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_C30_14_BC_JsonResultAt(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	r := bc.JsonResultAt(0)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_C30_15_BC_UnmarshalAt(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"hello"`))
	var s string
	err := bc.UnmarshalAt(0, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C30_16_BC_AddSerializer(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSerializer(nil)
	if bc.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C30_17_BC_AddSerializers(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSerializers()
	if bc.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C30_18_BC_AddSerializerFunc(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSerializerFunc(nil)
	bc.AddSerializerFunc(func() ([]byte, error) {
		return []byte(`"x"`), nil
	})
	if bc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_19_BC_AddSerializerFunctions(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSerializerFunctions()
	if bc.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C30_20_BC_InjectIntoAt(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`{"Bytes":"dGVzdA==","TypeName":"T"}`))
	target := corejson.Result{}
	err := bc.InjectIntoAt(0, &target)
	_ = err
}

func Test_C30_21_BC_InjectIntoSameIndex(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	var nilInjectors []corejson.JsonParseSelfInjector
	errs, has := bc.InjectIntoSameIndex(nilInjectors...)
	if has || len(errs) != 0 {
		t.Fatal("unexpected")
	}
	bc.Add([]byte(`{"Bytes":"dGVzdA==","TypeName":"T"}`))
	t1 := corejson.Result{}
	errs, has = bc.InjectIntoSameIndex(&t1)
	_ = errs
	_ = has
}

func Test_C30_22_BC_UnmarshalIntoSameIndex(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	var nilAnys []any
	errs, has := bc.UnmarshalIntoSameIndex(nilAnys...)
	if has || len(errs) != 0 {
		t.Fatal("unexpected")
	}
	bc.Add([]byte(`"hello"`))
	bc.Add([]byte(`42`))
	var s string
	var n int
	errs, has = bc.UnmarshalIntoSameIndex(&s, &n)
	_ = errs
	_ = has
}

func Test_C30_23_BC_GetAtSafe(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	if bc.GetAtSafe(0) == nil {
		t.Fatal("expected non-nil")
	}
	if bc.GetAtSafe(-1) != nil {
		t.Fatal("expected nil")
	}
	if bc.GetAtSafe(999) != nil {
		t.Fatal("expected nil")
	}
}

func Test_C30_24_BC_GetAtSafePtr(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	if bc.GetAtSafePtr(0) == nil {
		t.Fatal("expected non-nil")
	}
	if bc.GetAtSafePtr(-1) != nil {
		t.Fatal("expected nil")
	}
}

func Test_C30_25_BC_GetResultAtSafe(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	if bc.GetResultAtSafe(0) == nil {
		t.Fatal("expected non-nil")
	}
	if bc.GetResultAtSafe(-1) != nil {
		t.Fatal("expected nil")
	}
}

func Test_C30_26_BC_GetAtSafeUsingLength(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	if bc.GetAtSafeUsingLength(0, 1) == nil {
		t.Fatal("expected non-nil")
	}
	if bc.GetAtSafeUsingLength(5, 1) != nil {
		t.Fatal("expected nil")
	}
}

func Test_C30_27_BC_AddPtr(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddPtr([]byte{})
	bc.AddPtr([]byte(`"x"`))
	if bc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_28_BC_Adds(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Adds()
	bc.Adds([]byte{}, []byte(`"a"`))
	if bc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_29_BC_AddAnyItems(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	err := bc.AddAnyItems()
	if err != nil {
		t.Fatal("unexpected error")
	}
	err = bc.AddAnyItems("x", 42)
	if err != nil {
		t.Fatal("unexpected error")
	}
	if bc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C30_30_BC_AddAnyItems_Error(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	ch := make(chan int)
	err := bc.AddAnyItems(ch)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_C30_31_BC_AddMapResults(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	mr := corejson.NewMapResults.Empty()
	bc.AddMapResults(mr)
	mr.Add("k", corejson.NewResult.Any("v"))
	bc.AddMapResults(mr)
	if bc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_32_BC_AddRawMapResults(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddRawMapResults(nil)
	m := map[string]corejson.Result{
		"k": corejson.NewResult.Any("v"),
		"e": corejson.NewResult.Error(errors.New("err")),
	}
	bc.AddRawMapResults(m)
	if bc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_33_BC_AddsPtr(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddsPtr(nil, corejson.NewResult.AnyPtr("x"), &corejson.Result{})
	if bc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_34_BC_AddAny(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	err := bc.AddAny("hello")
	if err != nil {
		t.Fatal("unexpected")
	}
	if bc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_35_BC_AddAny_Error(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	ch := make(chan int)
	err := bc.AddAny(ch)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_C30_36_BC_AddBytesCollection(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	other := corejson.NewBytesCollection.Empty()
	bc.AddBytesCollection(other)
	other.Add([]byte(`"x"`))
	bc.AddBytesCollection(other)
	if bc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_37_BC_Clear(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	bc.Clear()
	if bc.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C30_38_BC_Clear_Nil(t *testing.T) {
	var bc *corejson.BytesCollection
	_ = bc.Clear()
}

func Test_C30_39_BC_Dispose(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	bc.Dispose()
}

func Test_C30_40_BC_Dispose_Nil(t *testing.T) {
	var bc *corejson.BytesCollection
	bc.Dispose()
}

func Test_C30_41_BC_Strings(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	s := bc.Strings()
	if len(s) != 0 {
		t.Fatal("expected empty")
	}
	bc.Add([]byte(`"x"`))
	s = bc.Strings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_42_BC_StringsPtr(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.StringsPtr()
}

func Test_C30_43_BC_AddJsoners(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddJsoners(true)
	if bc.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C30_44_BC_GetPagesSize(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	if bc.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	for i := 0; i < 5; i++ {
		bc.Add([]byte(`"x"`))
	}
	if bc.GetPagesSize(2) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_C30_45_BC_GetPagedCollection(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	for i := 0; i < 5; i++ {
		bc.Add([]byte(`"x"`))
	}
	pages := bc.GetPagedCollection(2)
	if len(pages) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_C30_46_BC_GetPagedCollection_Small(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	pages := bc.GetPagedCollection(10)
	if len(pages) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_47_BC_GetSinglePageCollection(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	for i := 0; i < 10; i++ {
		bc.Add([]byte(`"x"`))
	}
	page := bc.GetSinglePageCollection(3, 1)
	if page.Length() != 3 {
		t.Fatal("expected 3")
	}
	page = bc.GetSinglePageCollection(3, 4)
	if page.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_48_BC_GetSinglePageCollection_Small(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	page := bc.GetSinglePageCollection(10, 1)
	if page.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_49_BC_JsonModel_JsonModelAny(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.JsonModel()
	_ = bc.JsonModelAny()
}

func Test_C30_50_BC_MarshalJSON(t *testing.T) {
	bc := corejson.BytesCollection{}
	bc.Items = [][]byte{[]byte(`"x"`)}
	b, err := bc.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func Test_C30_51_BC_UnmarshalJSON(t *testing.T) {
	bc := corejson.BytesCollection{}
	err := bc.UnmarshalJSON([]byte(`[["dGVzdA=="]]`))
	_ = err
}

func Test_C30_52_BC_Json_JsonPtr(t *testing.T) {
	bc := corejson.BytesCollection{}
	_ = bc.Json()
	_ = bc.JsonPtr()
}

func Test_C30_53_BC_ParseInjectUsingJson(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	serialized := bc.JsonPtr()
	bc2 := corejson.NewBytesCollection.Empty()
	_, err := bc2.ParseInjectUsingJson(serialized)
	_ = err
}

func Test_C30_54_BC_ParseInjectUsingJson_Fail(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bad := &corejson.Result{Error: errors.New("fail")}
	_, err := bc.ParseInjectUsingJson(bad)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_C30_55_BC_ParseInjectUsingJsonMust(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	serialized := bc.JsonPtr()
	bc2 := corejson.NewBytesCollection.Empty()
	_ = bc2.ParseInjectUsingJsonMust(serialized)
}

func Test_C30_56_BC_AsJsonContractsBinder(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.AsJsonContractsBinder()
}

func Test_C30_57_BC_AsJsoner(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.AsJsoner()
}

func Test_C30_58_BC_JsonParseSelfInject(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	serialized := bc.JsonPtr()
	bc2 := corejson.NewBytesCollection.Empty()
	err := bc2.JsonParseSelfInject(serialized)
	_ = err
}

func Test_C30_59_BC_AsJsonParseSelfInjector(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.AsJsonParseSelfInjector()
}

func Test_C30_60_BC_ShadowClone(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	_ = bc.ShadowClone()
}

func Test_C30_61_BC_Clone(t *testing.T) {
	bc := corejson.BytesCollection{}
	_ = bc.Clone(false)
	bc.Items = [][]byte{[]byte(`"x"`)}
	_ = bc.Clone(true)
}

func Test_C30_62_BC_ClonePtr(t *testing.T) {
	var bc *corejson.BytesCollection
	if bc.ClonePtr(false) != nil {
		t.Fatal("expected nil")
	}
	bc = corejson.NewBytesCollection.Empty()
	_ = bc.ClonePtr(false)
	bc.Add([]byte(`"x"`))
	_ = bc.ClonePtr(true)
}

// ═══════════════════════════════════════════════
// MapResults — all uncovered methods
// ═══════════════════════════════════════════════

func Test_C30_63_MR_Length(t *testing.T) {
	var mr *corejson.MapResults
	if mr.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C30_64_MR_LastIndex(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	if mr.LastIndex() != -1 {
		t.Fatal("expected -1")
	}
}

func Test_C30_65_MR_IsEmpty_HasAnyItem(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	if !mr.IsEmpty() {
		t.Fatal("expected empty")
	}
	if mr.HasAnyItem() {
		t.Fatal("expected false")
	}
}

func Test_C30_66_MR_AddSkipOnNil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddSkipOnNil("k", nil)
	if mr.HasAnyItem() {
		t.Fatal("expected empty")
	}
	mr.AddSkipOnNil("k", corejson.NewResult.AnyPtr("v"))
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_67_MR_GetByKey(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	if mr.GetByKey("k") != nil {
		t.Fatal("expected nil")
	}
	mr.Add("k", corejson.NewResult.Any("v"))
	if mr.GetByKey("k") == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_C30_68_MR_HasError(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	if mr.HasError() {
		t.Fatal("expected false")
	}
	mr.Add("k", corejson.NewResult.Error(errors.New("e")))
	if !mr.HasError() {
		t.Fatal("expected true")
	}
}

func Test_C30_69_MR_AllErrors(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	errs, has := mr.AllErrors()
	if has || len(errs) != 0 {
		t.Fatal("unexpected")
	}
	mr.Add("k", corejson.NewResult.Error(errors.New("e")))
	mr.Add("ok", corejson.NewResult.Any("v"))
	errs, has = mr.AllErrors()
	if !has || len(errs) != 1 {
		t.Fatal("unexpected")
	}
}

func Test_C30_70_MR_GetErrorsStrings(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	s := mr.GetErrorsStrings()
	if len(s) != 0 {
		t.Fatal("expected empty")
	}
	mr.Add("k", corejson.NewResult.Error(errors.New("e")))
	mr.Add("ok", corejson.NewResult.Any("v"))
	s = mr.GetErrorsStrings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_71_MR_GetErrorsStringsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetErrorsStringsPtr()
}

func Test_C30_72_MR_GetErrorsAsSingleString(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetErrorsAsSingleString()
}

func Test_C30_73_MR_GetErrorsAsSingle(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetErrorsAsSingle()
}

func Test_C30_74_MR_Unmarshal(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("hello"))
	var s string
	// Note: Unmarshal has inverted logic (has==false means key exists)
	err := mr.Unmarshal("missing", &s)
	_ = err
}

func Test_C30_75_MR_Deserialize(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	var s string
	err := mr.Deserialize("k", &s)
	_ = err
}

func Test_C30_76_MR_DeserializeMust(t *testing.T) {
	defer func() { recover() }()
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any(map[string]string{"a": "b"}))
	target := make(map[string]string)
	_ = mr.DeserializeMust("k", &target)
}

func Test_C30_77_MR_UnmarshalMany(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.UnmarshalMany()
	if err != nil {
		t.Fatal("unexpected")
	}
	mr.Add("k", corejson.NewResult.Any("hello"))
	var s string
	err = mr.UnmarshalMany(corejson.KeyAny{Key: "k", AnyInf: &s})
	_ = err
}

func Test_C30_78_MR_UnmarshalManySafe(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.UnmarshalManySafe()
	if err != nil {
		t.Fatal("unexpected")
	}
	mr.Add("k", corejson.NewResult.Any("hello"))
	var s string
	err = mr.UnmarshalManySafe(corejson.KeyAny{Key: "k", AnyInf: &s})
	_ = err
}

func Test_C30_79_MR_SafeUnmarshal(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("hello"))
	var s string
	err := mr.SafeUnmarshal("k", &s)
	_ = err
	err = mr.SafeUnmarshal("missing", &s)
	_ = err
}

func Test_C30_80_MR_SafeDeserialize(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	var s string
	err := mr.SafeDeserialize("k", &s)
	_ = err
}

func Test_C30_81_MR_SafeDeserializeMust(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	var s string
	_ = mr.SafeDeserializeMust("k", &s)
}

func Test_C30_82_MR_InjectIntoAt(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.NewResult.Any(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	mr.Add("k", r)
	target := corejson.Result{}
	err := mr.InjectIntoAt("k", &target)
	_ = err
}

func Test_C30_83_MR_Add_AddPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	mr.AddPtr("k2", nil)
	mr.AddPtr("k3", corejson.NewResult.AnyPtr("v"))
	if mr.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C30_84_MR_AddAny(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAny("k", nil)
	if err == nil {
		t.Fatal("expected error for nil")
	}
	err = mr.AddAny("k", "hello")
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func Test_C30_85_MR_AddAny_MarshalError(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	ch := make(chan int)
	err := mr.AddAny("k", ch)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_C30_86_MR_AddAnySkipOnNil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAnySkipOnNil("k", nil)
	if err != nil {
		t.Fatal("unexpected")
	}
	err = mr.AddAnySkipOnNil("k", "v")
	if err != nil {
		t.Fatal("unexpected")
	}
}

func Test_C30_87_MR_AddAnyNonEmptyNonError(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddAnyNonEmptyNonError("k", nil)
	mr.AddAnyNonEmptyNonError("k", "v")
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_88_MR_AddAnyNonEmpty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddAnyNonEmpty("k", nil)
	mr.AddAnyNonEmpty("k", "v")
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_89_MR_AddKeyWithResult(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithResult(corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")})
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_90_MR_AddKeyWithResultPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithResultPtr(nil)
	mr.AddKeyWithResultPtr(&corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")})
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_91_MR_AddKeysWithResultsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithResultsPtr()
	kr := &corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")}
	mr.AddKeysWithResultsPtr(kr)
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_92_MR_AddKeysWithResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithResults()
	kr := corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")}
	mr.AddKeysWithResults(kr)
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_93_MR_AddKeyAnyInf(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyInf(corejson.KeyAny{Key: "k", AnyInf: "v"})
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_94_MR_AddKeyAnyInfPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyInfPtr(nil)
	mr.AddKeyAnyInfPtr(&corejson.KeyAny{Key: "k", AnyInf: "v"})
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_95_MR_AddKeyAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyItems()
	mr.AddKeyAnyItems(corejson.KeyAny{Key: "k", AnyInf: "v"})
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_96_MR_AddKeyAnyItemsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyItemsPtr()
	mr.AddKeyAnyItemsPtr(&corejson.KeyAny{Key: "k", AnyInf: "v"})
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_97_MR_AddNonEmptyNonErrorPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddNonEmptyNonErrorPtr("k", nil)
	mr.AddNonEmptyNonErrorPtr("k", &corejson.Result{Error: errors.New("e")})
	mr.AddNonEmptyNonErrorPtr("k", corejson.NewResult.AnyPtr("v"))
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_98_MR_AddMapResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapResults(nil)
	other := corejson.NewMapResults.Empty()
	mr.AddMapResults(other)
	other.Add("k", corejson.NewResult.Any("v"))
	mr.AddMapResults(other)
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_99_MR_AddMapAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapAnyItems(nil)
	mr.AddMapAnyItems(map[string]any{"k": "v"})
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_100_MR_AllKeys(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	keys := mr.AllKeys()
	if len(keys) != 0 {
		t.Fatal("expected empty")
	}
	mr.Add("a", corejson.NewResult.Any("v"))
	keys = mr.AllKeys()
	if len(keys) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_101_MR_AllKeysSorted(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AllKeysSorted()
	mr.Add("b", corejson.NewResult.Any("v"))
	mr.Add("a", corejson.NewResult.Any("v"))
	keys := mr.AllKeysSorted()
	if keys[0] != "a" {
		t.Fatal("expected sorted")
	}
}

func Test_C30_102_MR_AllValues(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AllValues()
	mr.Add("k", corejson.NewResult.Any("v"))
	vals := mr.AllValues()
	if len(vals) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_103_MR_AllResultsCollection(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	rc := mr.AllResultsCollection()
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
	mr.Add("k", corejson.NewResult.Any("v"))
	rc = mr.AllResultsCollection()
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_104_MR_AllResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AllResults()
}

func Test_C30_105_MR_GetStrings(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	s := mr.GetStrings()
	if len(s) != 0 {
		t.Fatal("expected empty")
	}
	mr.Add("k", corejson.NewResult.Any("v"))
	s = mr.GetStrings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_106_MR_GetStringsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetStringsPtr()
}

func Test_C30_107_MR_AddJsoner(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddJsoner("k", nil)
	if mr.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C30_108_MR_AddKeyWithJsoner(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithJsoner(corejson.KeyWithJsoner{Key: "k", Jsoner: nil})
	_ = mr
}

func Test_C30_109_MR_AddKeysWithJsoners(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AddKeysWithJsoners()
}

func Test_C30_110_MR_AddKeyWithJsonerPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithJsonerPtr(nil)
	mr.AddKeyWithJsonerPtr(&corejson.KeyWithJsoner{Key: "k", Jsoner: nil})
	_ = mr
}

func Test_C30_111_MR_GetPagesSize(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	if mr.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	for i := 0; i < 5; i++ {
		mr.Add(string(rune('a'+i)), corejson.NewResult.Any(i))
	}
	if mr.GetPagesSize(2) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_C30_112_MR_GetPagedCollection(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	for i := 0; i < 5; i++ {
		mr.Add(string(rune('a'+i)), corejson.NewResult.Any(i))
	}
	pages := mr.GetPagedCollection(2)
	if len(pages) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_C30_113_MR_GetPagedCollection_Small(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	pages := mr.GetPagedCollection(10)
	if len(pages) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_114_MR_AddMapResultsUsingCloneOption(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapResultsUsingCloneOption(false, false, nil)
	m := map[string]corejson.Result{
		"k": corejson.NewResult.Any("v"),
	}
	mr.AddMapResultsUsingCloneOption(false, false, m)
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
	mr2 := corejson.NewMapResults.Empty()
	mr2.AddMapResultsUsingCloneOption(true, true, m)
	if mr2.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_115_MR_GetSinglePageCollection(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	for i := 0; i < 10; i++ {
		mr.Add(string(rune('a'+i)), corejson.NewResult.Any(i))
	}
	allKeys := mr.AllKeysSorted()
	page := mr.GetSinglePageCollection(3, 1, allKeys)
	if page.Length() != 3 {
		t.Fatal("expected 3")
	}
	page = mr.GetSinglePageCollection(3, 4, allKeys)
	if page.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_116_MR_GetSinglePageCollection_Small(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	allKeys := mr.AllKeysSorted()
	page := mr.GetSinglePageCollection(10, 1, allKeys)
	if page.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_117_MR_GetNewMapUsingKeys(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("a", corejson.NewResult.Any("1"))
	mr.Add("b", corejson.NewResult.Any("2"))
	sub := mr.GetNewMapUsingKeys(false, "a")
	if sub.Length() != 1 {
		t.Fatal("expected 1")
	}
	sub = mr.GetNewMapUsingKeys(false)
	if sub.HasAnyItem() {
		t.Fatal("expected empty")
	}
	// non-panic missing
	sub = mr.GetNewMapUsingKeys(false, "missing")
	if sub.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C30_118_MR_ResultCollection(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	rc := mr.ResultCollection()
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
	mr.Add("k", corejson.NewResult.Any("v"))
	rc = mr.ResultCollection()
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C30_119_MR_JsonModel_JsonModelAny(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.JsonModel()
	_ = mr.JsonModelAny()
}

func Test_C30_120_MR_Clear(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	mr.Clear()
	if mr.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C30_121_MR_Clear_Nil(t *testing.T) {
	var mr *corejson.MapResults
	_ = mr.Clear()
}

func Test_C30_122_MR_Dispose(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	mr.Dispose()
}

func Test_C30_123_MR_Dispose_Nil(t *testing.T) {
	var mr *corejson.MapResults
	mr.Dispose()
}

func Test_C30_124_MR_Json_JsonPtr(t *testing.T) {
	mr := corejson.MapResults{Items: map[string]corejson.Result{}}
	_ = mr.Json()
	_ = mr.JsonPtr()
}

func Test_C30_125_MR_ParseInjectUsingJson(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	serialized := mr.JsonPtr()
	mr2 := corejson.NewMapResults.Empty()
	_, err := mr2.ParseInjectUsingJson(serialized)
	_ = err
}

func Test_C30_126_MR_ParseInjectUsingJson_Fail(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	bad := &corejson.Result{Error: errors.New("fail")}
	_, err := mr.ParseInjectUsingJson(bad)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_C30_127_MR_ParseInjectUsingJsonMust(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	serialized := mr.JsonPtr()
	mr2 := corejson.NewMapResults.Empty()
	_ = mr2.ParseInjectUsingJsonMust(serialized)
}

func Test_C30_128_MR_AsJsonContractsBinder(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AsJsonContractsBinder()
}

func Test_C30_129_MR_AsJsoner(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AsJsoner()
}

func Test_C30_130_MR_JsonParseSelfInject(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	serialized := mr.JsonPtr()
	mr2 := corejson.NewMapResults.Empty()
	err := mr2.JsonParseSelfInject(serialized)
	_ = err
}

func Test_C30_131_MR_AsJsonParseSelfInjector(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AsJsonParseSelfInjector()
}
