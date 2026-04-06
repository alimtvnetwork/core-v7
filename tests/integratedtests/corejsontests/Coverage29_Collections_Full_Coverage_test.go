package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ═══════════════════════════════════════════════
// ResultsCollection — all uncovered methods
// ═══════════════════════════════════════════════

func Test_C29_01_RC_Length(t *testing.T) {
	var rc *corejson.ResultsCollection
	if rc.Length() != 0 {
		t.Fatal("expected 0")
	}
	rc2 := &corejson.ResultsCollection{}
	if rc2.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C29_02_RC_LastIndex(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	if rc.LastIndex() != -1 {
		t.Fatal("expected -1")
	}
}

func Test_C29_03_RC_IsEmpty_HasAnyItem(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	if !rc.IsEmpty() {
		t.Fatal("expected empty")
	}
	if rc.HasAnyItem() {
		t.Fatal("expected false")
	}
}

func Test_C29_04_RC_FirstOrDefault(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	if rc.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	rc.Add(corejson.NewResult.Any("x"))
	if rc.FirstOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_C29_05_RC_LastOrDefault(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	if rc.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
	rc.Add(corejson.NewResult.Any("x"))
	if rc.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_C29_06_RC_Take(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	taken := rc.Take(1)
	if taken.HasAnyItem() {
		t.Fatal("expected empty")
	}
	rc.Add(corejson.NewResult.Any("a"))
	rc.Add(corejson.NewResult.Any("b"))
	taken = rc.Take(1)
	if taken.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_07_RC_Limit(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	l := rc.Limit(5)
	if l.HasAnyItem() {
		t.Fatal("expected empty")
	}
	rc.Add(corejson.NewResult.Any("a"))
	rc.Add(corejson.NewResult.Any("b"))
	// TakeAllMinusOne is -1
	l = rc.Limit(-1)
	if l.Length() != 2 {
		t.Fatal("expected 2")
	}
	l = rc.Limit(1)
	if l.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_08_RC_Skip(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	s := rc.Skip(0)
	if s.HasAnyItem() {
		t.Fatal("expected empty")
	}
	rc.Add(corejson.NewResult.Any("a"))
	rc.Add(corejson.NewResult.Any("b"))
	s = rc.Skip(1)
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_09_RC_AddSkipOnNil(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSkipOnNil(nil)
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
	r := corejson.NewResult.AnyPtr("x")
	rc.AddSkipOnNil(r)
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_10_RC_AddNonNilNonError(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddNonNilNonError(nil)
	rc.AddNonNilNonError(&corejson.Result{Error: errors.New("e")})
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
	rc.AddNonNilNonError(corejson.NewResult.AnyPtr("x"))
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_11_RC_GetAt(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	r := rc.GetAt(0)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_C29_12_RC_HasError(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	if rc.HasError() {
		t.Fatal("expected false")
	}
	rc.Add(corejson.NewResult.Error(errors.New("e")))
	if !rc.HasError() {
		t.Fatal("expected true")
	}
}

func Test_C29_13_RC_AllErrors(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	errs, has := rc.AllErrors()
	if has || len(errs) != 0 {
		t.Fatal("unexpected")
	}
	rc.Add(corejson.NewResult.Any("x"))
	rc.Add(corejson.NewResult.Error(errors.New("e")))
	errs, has = rc.AllErrors()
	if !has || len(errs) != 1 {
		t.Fatal("unexpected")
	}
}

func Test_C29_14_RC_GetErrorsStrings(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	s := rc.GetErrorsStrings()
	if len(s) != 0 {
		t.Fatal("expected empty")
	}
	rc.Add(corejson.NewResult.Any("x"))
	rc.Add(corejson.NewResult.Error(errors.New("e")))
	s = rc.GetErrorsStrings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_15_RC_GetErrorsStringsPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	s := rc.GetErrorsStringsPtr()
	_ = s
}

func Test_C29_16_RC_GetErrorsAsSingleString(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	s := rc.GetErrorsAsSingleString()
	_ = s
}

func Test_C29_17_RC_GetErrorsAsSingle(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	err := rc.GetErrorsAsSingle()
	_ = err
}

func Test_C29_18_RC_UnmarshalAt(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("hello"))
	var s string
	err := rc.UnmarshalAt(0, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_C29_19_RC_InjectIntoAt(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	r := corejson.NewResult.Any(corejson.Result{Bytes: []byte(`"test"`), TypeName: "T"})
	rc.Add(r)
	target := corejson.Result{}
	err := rc.InjectIntoAt(0, &target)
	_ = err
}

func Test_C29_20_RC_InjectIntoSameIndex(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	// Pass true nil variadic slice
	var nilSlice []corejson.JsonParseSelfInjector
	errs, has := rc.InjectIntoSameIndex(nilSlice...)
	if has || len(errs) != 0 {
		t.Fatal("unexpected")
	}
}

func Test_C29_21_RC_UnmarshalIntoSameIndex(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	var nilSlice []any
	errs, has := rc.UnmarshalIntoSameIndex(nilSlice...)
	if has || len(errs) != 0 {
		t.Fatal("unexpected")
	}

	rc.Add(corejson.NewResult.Any("hello"))
	rc.Add(corejson.NewResult.Error(errors.New("e")))
	rc.Add(corejson.NewResult.Any("world"))
	var s1 string
	var s3 string
	errs, has = rc.UnmarshalIntoSameIndex(&s1, nil, &s3)
	_ = errs
	_ = has
}

func Test_C29_22_RC_GetAtSafe(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	r := rc.GetAtSafe(0)
	if r == nil {
		t.Fatal("expected non-nil")
	}
	r = rc.GetAtSafe(-1)
	if r != nil {
		t.Fatal("expected nil")
	}
	r = rc.GetAtSafe(999)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func Test_C29_23_RC_GetAtSafeUsingLength(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	r := rc.GetAtSafeUsingLength(0, 1)
	if r == nil {
		t.Fatal("expected non-nil")
	}
	r = rc.GetAtSafeUsingLength(5, 1)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func Test_C29_24_RC_AddPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddPtr(nil)
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
	rc.AddPtr(corejson.NewResult.AnyPtr("x"))
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_25_RC_Add(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_26_RC_Adds(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Adds()
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
	rc.Adds(corejson.NewResult.Any("a"), corejson.NewResult.Any("b"))
	if rc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C29_27_RC_AddSerializer(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializer(nil)
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C29_28_RC_AddSerializers(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializers()
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C29_29_RC_AddSerializerFunc(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializerFunc(nil)
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
	rc.AddSerializerFunc(func() ([]byte, error) {
		return []byte(`"x"`), nil
	})
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_30_RC_AddSerializerFunctions(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializerFunctions()
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C29_31_RC_AddMapResults(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	mr := corejson.NewMapResults.Empty()
	rc.AddMapResults(mr)
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
	mr.Add("k", corejson.NewResult.Any("v"))
	rc.AddMapResults(mr)
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_32_RC_AddRawMapResults(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddRawMapResults(nil)
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C29_33_RC_AddsPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddsPtr(nil, corejson.NewResult.AnyPtr("x"))
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_34_RC_AddAny(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddAny(nil)
	rc.AddAny("hello")
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_35_RC_AddAnyItems(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddAnyItems(nil, "a", nil, "b")
	if rc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C29_36_RC_AddAnyItemsSlice(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddAnyItemsSlice(nil)
	rc.AddAnyItemsSlice([]any{nil, "a"})
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_37_RC_AddResultsCollection(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddResultsCollection(nil)
	other := corejson.NewResultsCollection.Empty()
	other.Add(corejson.NewResult.Any("x"))
	rc.AddResultsCollection(other)
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_38_RC_AddNonNilItemsPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddNonNilItemsPtr()
	rc.AddNonNilItemsPtr(nil, corejson.NewResult.AnyPtr("x"))
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_39_RC_NonPtr_Ptr(t *testing.T) {
	rc := corejson.ResultsCollection{}
	_ = rc.NonPtr()
	_ = rc.Ptr()
}

func Test_C29_40_RC_Clear(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	rc.Clear()
	if rc.HasAnyItem() {
		t.Fatal("expected empty after clear")
	}
}

func Test_C29_41_RC_Clear_Nil(t *testing.T) {
	var rc *corejson.ResultsCollection
	result := rc.Clear()
	_ = result
}

func Test_C29_42_RC_Dispose(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	rc.Dispose()
}

func Test_C29_43_RC_Dispose_Nil(t *testing.T) {
	var rc *corejson.ResultsCollection
	rc.Dispose()
}

func Test_C29_44_RC_GetStrings(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	s := rc.GetStrings()
	if len(s) != 0 {
		t.Fatal("expected empty")
	}
	rc.Add(corejson.NewResult.Any("hello"))
	s = rc.GetStrings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_45_RC_GetStringsPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	_ = rc.GetStringsPtr()
}

func Test_C29_46_RC_AddJsoners(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddJsoners(true)
	if rc.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_C29_47_RC_GetPagesSize(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	if rc.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	if rc.GetPagesSize(-1) != 0 {
		t.Fatal("expected 0")
	}
	for i := 0; i < 5; i++ {
		rc.Add(corejson.NewResult.Any(i))
	}
	if rc.GetPagesSize(2) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_C29_48_RC_GetPagedCollection(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	for i := 0; i < 5; i++ {
		rc.Add(corejson.NewResult.Any(i))
	}
	pages := rc.GetPagedCollection(2)
	if len(pages) != 3 {
		t.Fatal("expected 3 pages")
	}
}

func Test_C29_49_RC_GetPagedCollection_SmallSize(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	pages := rc.GetPagedCollection(10)
	if len(pages) != 1 {
		t.Fatal("expected 1 page")
	}
}

func Test_C29_50_RC_GetSinglePageCollection(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	for i := 0; i < 10; i++ {
		rc.Add(corejson.NewResult.Any(i))
	}
	page := rc.GetSinglePageCollection(3, 1)
	if page.Length() != 3 {
		t.Fatal("expected 3")
	}
	page = rc.GetSinglePageCollection(3, 4)
	if page.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_51_RC_GetSinglePageCollection_Small(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	page := rc.GetSinglePageCollection(10, 1)
	if page.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C29_52_RC_JsonModel_JsonModelAny(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	_ = rc.JsonModel()
	_ = rc.JsonModelAny()
}

func Test_C29_53_RC_Json_JsonPtr(t *testing.T) {
	rc := corejson.ResultsCollection{}
	_ = rc.Json()
	_ = rc.JsonPtr()
}

func Test_C29_54_RC_ParseInjectUsingJson(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	serialized := rc.JsonPtr()
	rc2 := corejson.NewResultsCollection.Empty()
	_, err := rc2.ParseInjectUsingJson(serialized)
	_ = err
}

func Test_C29_55_RC_ParseInjectUsingJson_Fail(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	bad := &corejson.Result{Error: errors.New("fail")}
	_, err := rc.ParseInjectUsingJson(bad)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_C29_56_RC_ParseInjectUsingJsonMust(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	serialized := rc.JsonPtr()
	rc2 := corejson.NewResultsCollection.Empty()
	_ = rc2.ParseInjectUsingJsonMust(serialized)
}

func Test_C29_57_RC_AsJsonContractsBinder(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	_ = rc.AsJsonContractsBinder()
}

func Test_C29_58_RC_AsJsoner(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	_ = rc.AsJsoner()
}

func Test_C29_59_RC_JsonParseSelfInject(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	serialized := rc.JsonPtr()
	rc2 := corejson.NewResultsCollection.Empty()
	err := rc2.JsonParseSelfInject(serialized)
	_ = err
}

func Test_C29_60_RC_AsJsonParseSelfInjector(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	_ = rc.AsJsonParseSelfInjector()
}

func Test_C29_61_RC_ShadowClone(t *testing.T) {
	rc := corejson.ResultsCollection{}
	_ = rc.ShadowClone()
}

func Test_C29_62_RC_Clone(t *testing.T) {
	rc := corejson.ResultsCollection{}
	_ = rc.Clone(false)
	rc.Items = []corejson.Result{corejson.NewResult.Any("x")}
	_ = rc.Clone(true)
}

func Test_C29_63_RC_ClonePtr(t *testing.T) {
	var rc *corejson.ResultsCollection
	if rc.ClonePtr(false) != nil {
		t.Fatal("expected nil")
	}
	rc = corejson.NewResultsCollection.Empty()
	_ = rc.ClonePtr(false)
	rc.Add(corejson.NewResult.Any("x"))
	_ = rc.ClonePtr(true)
}

// ─── UnmarshalIntoSameIndex edge: empty json bytes item ───

func Test_C29_64_RC_UnmarshalIntoSameIndex_EmptyJsonBytes(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Create([]byte(`{}`), nil, "T"))
	var m map[string]any
	errs, has := rc.UnmarshalIntoSameIndex(&m)
	_ = errs
	_ = has
}

// ─── InjectIntoSameIndex with error result and valid injector ───

func Test_C29_65_RC_InjectIntoSameIndex_ErrorResult(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Error(errors.New("e")))
	r := corejson.Result{}
	errs, has := rc.InjectIntoSameIndex(&r)
	if !has {
		t.Fatal("expected hasAnyError true")
	}
	_ = errs
}
