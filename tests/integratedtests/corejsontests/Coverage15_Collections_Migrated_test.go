package corejsontests

import (
	"errors"
	"testing"
	"time"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ── Migrated from Coverage04, 08, 09, 10, 12, 17 — Collections & MapResults ──

func Test_C04_ResultsCollection_BasicOps(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	if !c.IsEmpty() || c.HasAnyItem() || c.Length() != 0 { t.Fatal("should be empty") }
	if c.LastIndex() != -1 { t.Fatal("expected -1") }
	c.Add(corejson.NewResult.Any("a"))
	c.Add(corejson.NewResult.Any("b"))
	if c.Length() != 2 { t.Fatal("expected 2") }
	if c.FirstOrDefault() == nil || c.LastOrDefault() == nil { t.Fatal("expected non-nil") }
}

func Test_C04_ResultsCollection_TakeSkipLimit(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(5)
	for i := 0; i < 5; i++ { c.Add(corejson.NewResult.Any(i)) }
	if c.Take(3).Length() != 3 { t.Fatal("expected 3") }
	if c.Skip(2).Length() != 3 { t.Fatal("expected 3") }
	if c.Limit(3).Length() != 3 { t.Fatal("expected 3") }
	if c.Limit(-2).Length() != 5 { t.Fatal("expected all") }
	empty := corejson.NewResultsCollection.Empty()
	if empty.Take(1).Length() != 0 { t.Fatal("expected 0") }
}

func Test_C04_ResultsCollection_AddMethods(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(10)
	r := corejson.NewResult.AnyPtr("x")
	c.AddSkipOnNil(r)
	c.AddSkipOnNil(nil)
	c.AddNonNilNonError(r)
	c.AddNonNilNonError(nil)
	c.AddNonNilNonError(&corejson.Result{Error: errors.New("e")})
	c.AddPtr(r)
	c.AddPtr(nil)
	c.Adds(corejson.NewResult.Any("a"), corejson.NewResult.Any("b"))
	c.AddsPtr(r, nil)
	c.AddAny("z")
	c.AddAny(nil)
	c.AddAnyItems("a", nil, "b")
}

func Test_C04_ResultsCollection_Errors(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(3)
	c.Add(corejson.NewResult.Any("ok"))
	c.Add(corejson.Result{Error: errors.New("e1")})
	if !c.HasError() { t.Fatal("expected error") }
	errs, has := c.AllErrors()
	if !has || len(errs) != 1 { t.Fatal("expected 1 error") }
	_ = c.GetErrorsStrings()
	_ = c.GetErrorsStringsPtr()
	_ = c.GetErrorsAsSingleString()
	_ = c.GetErrorsAsSingle()
}

func Test_C04_ResultsCollection_UnmarshalAt(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(1)
	c.Add(corejson.NewResult.Any("hello"))
	var s string
	err := c.UnmarshalAt(0, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func Test_C04_ResultsCollection_GetAtSafe(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(1)
	c.Add(corejson.NewResult.Any("x"))
	if c.GetAtSafe(0) == nil { t.Fatal("expected non-nil") }
	if c.GetAtSafe(-1) != nil { t.Fatal("expected nil") }
	if c.GetAtSafe(5) != nil { t.Fatal("expected nil") }
}

func Test_C04_ResultsCollection_Paging(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(10)
	for i := 0; i < 10; i++ { c.Add(corejson.NewResult.Any(i)) }
	if c.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if c.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	paged := c.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4 pages") }
	single := c.GetSinglePageCollection(3, 1)
	if single.Length() != 3 { t.Fatal("expected 3") }
}

func Test_C04_ResultsCollection_Json(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(2)
	c.Add(corejson.NewResult.Any("x"))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func Test_C04_ResultsCollection_Clone(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(2)
	c.Add(corejson.NewResult.Any("x"))
	_ = c.ShadowClone()
	_ = c.Clone(true)
	cp := c.ClonePtr(true)
	_ = cp
	var nilC *corejson.ResultsCollection
	if nilC.ClonePtr(true) != nil { t.Fatal("expected nil") }
}

func Test_C04_ResultsCollection_ClearDispose(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(2)
	c.Add(corejson.NewResult.Any("x"))
	c.Clear()
	time.Sleep(10 * time.Millisecond)
	if c.Length() != 0 { t.Fatal("expected 0") }
	c.Dispose()
	var nilC *corejson.ResultsCollection
	nilC.Clear()
	nilC.Dispose()
}

func Test_C04_ResultsCollection_GetStrings(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(2)
	c.Add(corejson.NewResult.Any("a"))
	if len(c.GetStrings()) != 1 { t.Fatal("expected 1") }
	_ = c.GetStringsPtr()
}

func Test_C04_ResultsCollection_Nil(t *testing.T) {
	var nilC *corejson.ResultsCollection
	if nilC.Length() != 0 { t.Fatal("expected 0") }
	if nilC.LastIndex() != -1 { t.Fatal("expected -1") }
	if !nilC.IsEmpty() { t.Fatal("expected empty") }
	if nilC.HasAnyItem() { t.Fatal("expected false") }
	if nilC.FirstOrDefault() != nil { t.Fatal("expected nil") }
	if nilC.LastOrDefault() != nil { t.Fatal("expected nil") }
}

// ── BytesCollection ──

func Test_C04_BytesCollection_BasicOps(t *testing.T) {
	c := corejson.NewBytesCollection.Empty()
	if !c.IsEmpty() { t.Fatal("should be empty") }
	c.Add([]byte("hello"))
	if c.Length() != 1 || !c.HasAnyItem() { t.Fatal("expected 1") }
	if c.FirstOrDefault() == nil || c.LastOrDefault() == nil { t.Fatal("expected non-nil") }
}

func Test_C04_BytesCollection_AddMethods(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	c.AddSkipOnNil(nil)
	c.AddSkipOnNil([]byte("x"))
	c.AddNonEmpty([]byte{})
	c.AddNonEmpty([]byte("y"))
	c.AddPtr([]byte{})
	c.AddPtr([]byte("z"))
	c.Adds([]byte("a"), []byte{}, []byte("b"))
}

func Test_C04_BytesCollection_TakeSkipLimit(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	c.Add([]byte("a")).Add([]byte("b")).Add([]byte("c"))
	if c.Take(2).Length() != 2 { t.Fatal("expected 2") }
	if c.Skip(1).Length() != 2 { t.Fatal("expected 2") }
	if c.Limit(2).Length() != 2 { t.Fatal("expected 2") }
	if c.Limit(-1).Length() != 3 { t.Fatal("expected 3") }
}

func Test_C04_BytesCollection_ClearDispose(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte("x"))
	c.Clear()
	time.Sleep(10 * time.Millisecond)
	if c.Length() != 0 { t.Fatal("expected 0") }
	c.Dispose()
	var nilC *corejson.BytesCollection
	nilC.Clear()
	nilC.Dispose()
}

func Test_C04_BytesCollection_Clone(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte("x"))
	_ = c.ShadowClone()
	_ = c.Clone(true)
	cp := c.ClonePtr(true)
	_ = cp
	var nilC *corejson.BytesCollection
	if nilC.ClonePtr(true) != nil { t.Fatal("expected nil") }
}

func Test_C04_BytesCollection_Json(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"x"`))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	b, err := c.MarshalJSON()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func Test_C04_BytesCollection_Paging(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ { c.Add([]byte(`"x"`)) }
	if c.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if c.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	paged := c.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4") }
}

func Test_C04_BytesCollection_Nil(t *testing.T) {
	var nilC *corejson.BytesCollection
	if nilC.Length() != 0 { t.Fatal("expected 0") }
	if nilC.LastIndex() != -1 { t.Fatal("expected -1") }
	if !nilC.IsEmpty() { t.Fatal("expected empty") }
	if nilC.HasAnyItem() { t.Fatal("expected false") }
	if nilC.FirstOrDefault() != nil { t.Fatal("expected nil") }
	if nilC.LastOrDefault() != nil { t.Fatal("expected nil") }
}

func Test_C04_BytesCollection_AddResult(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	r := corejson.NewResult.Any("hello")
	c.AddResult(r)
	c.AddResultPtr(nil)
	c.AddResultPtr(&r)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func Test_C04_BytesCollection_AddAny(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	err := c.AddAny("hello")
	if err != nil { t.Fatal(err) }
	err2 := c.AddAnyItems("a", "b")
	if err2 != nil { t.Fatal(err2) }
}

func Test_C04_BytesCollection_GetAtSafe(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"x"`))
	if c.GetAtSafe(0) == nil { t.Fatal("expected non-nil") }
	if c.GetAtSafe(-1) != nil { t.Fatal("expected nil") }
	if c.GetAtSafe(5) != nil { t.Fatal("expected nil") }
	if c.GetAtSafePtr(0) == nil { t.Fatal("expected non-nil") }
	if c.GetResultAtSafe(0) == nil { t.Fatal("expected non-nil") }
	if c.GetResultAtSafe(5) != nil { t.Fatal("expected nil") }
}

func Test_C04_BytesCollection_Strings(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	if len(c.Strings()) != 1 { t.Fatal("expected 1") }
	_ = c.StringsPtr()
}

func Test_C04_BytesCollection_Serializers(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.AddSerializer(nil)
	c.AddSerializers()
	c.AddSerializerFunc(nil)
	c.AddSerializerFunctions()
}

func Test_C04_BytesCollection_MapResults(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	mr := corejson.NewMapResults.Empty()
	c.AddMapResults(mr)
	c.AddRawMapResults(nil)
	c.AddRawMapResults(map[string]corejson.Result{"a": corejson.NewResult.Any("x")})
}

func Test_C04_BytesCollection_UnmarshalAt(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"hello"`))
	var s string
	err := c.UnmarshalAt(0, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func Test_C04_BytesCollection_AddBytesCollection(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	c2 := corejson.NewBytesCollection.UsingCap(1)
	c2.Add([]byte(`"b"`))
	c.AddBytesCollection(c2)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

// ── ResultsPtrCollection ──

func Test_C04_ResultsPtrCollection_BasicOps(t *testing.T) {
	var nilC *corejson.ResultsPtrCollection
	if nilC.Length() != 0 { t.Fatal("expected 0") }
	if !nilC.IsEmpty() { t.Fatal("expected empty") }
	if nilC.FirstOrDefault() != nil { t.Fatal("expected nil") }
	if nilC.LastOrDefault() != nil { t.Fatal("expected nil") }

	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("hello"))
	if c.Length() != 1 { t.Fatal("expected 1") }
	if c.FirstOrDefault() == nil || c.LastOrDefault() == nil { t.Fatal("expected non-nil") }
}

func Test_C04_ResultsPtrCollection_AddMethods(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(10)
	c.AddSkipOnNil(nil)
	c.AddSkipOnNil(corejson.NewResult.AnyPtr("x"))
	c.AddNonNilNonError(nil)
	c.AddNonNilNonError(&corejson.Result{Error: errors.New("e")})
	c.AddNonNilNonError(corejson.NewResult.AnyPtr("x"))
	c.AddResult(corejson.NewResult.Any("x"))
	c.Adds(nil, corejson.NewResult.AnyPtr("x"))
	c.AddAny(nil)
	c.AddAny("x")
	c.AddAnyItems(nil, "y")
	c.AddResultsCollection(nil)
	sub := corejson.NewResultsPtrCollection.UsingResults(corejson.NewResult.AnyPtr("sub"))
	c.AddResultsCollection(sub)
	c.AddNonNilItemsPtr(nil)
	c.AddNonNilItemsPtr(nil, corejson.NewResult.AnyPtr("x"))
	c.AddNonNilItems(nil, corejson.NewResult.AnyPtr("x"))
}

func Test_C04_ResultsPtrCollection_TakeSkipLimit(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(5)
	for i := 0; i < 5; i++ { c.Add(corejson.NewResult.AnyPtr(i)) }
	if c.Take(3).Length() != 3 { t.Fatal("expected 3") }
	if c.Skip(2).Length() != 3 { t.Fatal("expected 3") }
	if c.Limit(3).Length() != 3 { t.Fatal("expected 3") }
	if c.Limit(-2).Length() != 5 { t.Fatal("expected all") }
}

func Test_C04_ResultsPtrCollection_Errors(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(2)
	c.Add(corejson.NewResult.AnyPtr("ok"))
	c.Add(&corejson.Result{Error: errors.New("e")})
	if !c.HasError() { t.Fatal("expected error") }
	errs, has := c.AllErrors()
	if !has || len(errs) != 1 { t.Fatal("expected 1") }
	_ = c.GetErrorsStrings()
	_ = c.GetErrorsStringsPtr()
	_ = c.GetErrorsAsSingleString()
	_ = c.GetErrorsAsSingle()
}

func Test_C04_ResultsPtrCollection_ClearDispose(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(2)
	c.Add(corejson.NewResult.AnyPtr("x"))
	c.Clear()
	time.Sleep(10 * time.Millisecond)
	if c.Length() != 0 { t.Fatal("expected 0") }
	c.Dispose()
	var nilC *corejson.ResultsPtrCollection
	nilC.Clear()
	nilC.Dispose()
}

func Test_C04_ResultsPtrCollection_Clone(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(2)
	c.Add(corejson.NewResult.AnyPtr("x"))
	cp := c.Clone(true)
	_ = cp
	var nilC *corejson.ResultsPtrCollection
	if nilC.Clone(true) != nil { t.Fatal("expected nil") }
}

func Test_C04_ResultsPtrCollection_Json(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(1)
	c.Add(corejson.NewResult.AnyPtr("x"))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.NonPtr()
	_ = c.Ptr()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func Test_C04_ResultsPtrCollection_Paging(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(10)
	for i := 0; i < 10; i++ { c.Add(corejson.NewResult.AnyPtr(i)) }
	if c.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if c.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	paged := c.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4") }
}

func Test_C04_ResultsPtrCollection_GetStrings(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(2)
	c.Add(corejson.NewResult.AnyPtr("a"))
	if len(c.GetStrings()) != 1 { t.Fatal("expected 1") }
	_ = c.GetStringsPtr()
}

func Test_C04_ResultsPtrCollection_Serializers(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(2)
	c.AddSerializer(nil)
	c.AddSerializers()
	c.AddSerializerFunc(nil)
	c.AddSerializerFunctions()
}

func Test_C04_ResultsPtrCollection_Creators(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.AnyItems("a", "b")
	_ = corejson.NewResultsPtrCollection.AnyItemsPlusCap(5, "a")
	_ = corejson.NewResultsPtrCollection.AnyItemsPlusCap(5)
	_ = corejson.NewResultsPtrCollection.UsingResults(corejson.NewResult.AnyPtr("x"))
	_ = corejson.NewResultsPtrCollection.UsingResultsPlusCap(5, corejson.NewResult.AnyPtr("x"))
	_ = corejson.NewResultsPtrCollection.UsingResultsPlusCap(5)
	_ = corejson.NewResultsPtrCollection.Serializers()
	_, _ = corejson.NewResultsPtrCollection.UnmarshalUsingBytes([]byte(`{}`))
}

// ── MapResults ──

func Test_C04_MapResults_BasicOps(t *testing.T) {
	var nilM *corejson.MapResults
	if nilM.Length() != 0 { t.Fatal("expected 0") }
	if !nilM.IsEmpty() { t.Fatal("expected empty") }
	if nilM.HasAnyItem() { t.Fatal("expected false") }

	m := corejson.NewMapResults.Empty()
	m.Add("a", corejson.NewResult.Any("hello"))
	if m.Length() != 1 { t.Fatal("expected 1") }
	r := m.GetByKey("a")
	if r == nil { t.Fatal("expected non-nil") }
	if m.GetByKey("missing") != nil { t.Fatal("expected nil") }
}

func Test_C04_MapResults_AddMethods(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(10)
	m.AddSkipOnNil("a", nil)
	m.AddSkipOnNil("a", &corejson.Result{Bytes: []byte(`"x"`)})
	m.AddPtr("b", nil)
	m.AddPtr("b", &corejson.Result{Bytes: []byte(`"y"`)})
	_ = m.AddAny("c", "hello")
	_ = m.AddAny("d", nil)
	_ = m.AddAnySkipOnNil("e", nil)
	_ = m.AddAnySkipOnNil("e", "val")
	m.AddAnyNonEmptyNonError("f", nil)
	m.AddAnyNonEmptyNonError("f", "val")
	m.AddAnyNonEmpty("g", nil)
	m.AddAnyNonEmpty("g", "val")
	m.AddNonEmptyNonErrorPtr("h", nil)
	m.AddNonEmptyNonErrorPtr("h", &corejson.Result{Error: errors.New("e")})
	m.AddNonEmptyNonErrorPtr("h", &corejson.Result{Bytes: []byte(`"z"`)})

	m.AddKeyWithResult(corejson.KeyWithResult{Key: "i", Result: corejson.NewResult.Any("v")})
	m.AddKeyWithResultPtr(nil)
	kr := &corejson.KeyWithResult{Key: "j", Result: corejson.NewResult.Any("v")}
	m.AddKeyWithResultPtr(kr)
	m.AddKeysWithResultsPtr()
	m.AddKeysWithResultsPtr(kr)
	m.AddKeysWithResults(corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")})
	m.AddKeyAnyInf(corejson.KeyAny{Key: "l", AnyInf: "val"})
	m.AddKeyAnyInfPtr(nil)
	ka := &corejson.KeyAny{Key: "m", AnyInf: "val"}
	m.AddKeyAnyInfPtr(ka)
	m.AddKeyAnyItems(corejson.KeyAny{Key: "n", AnyInf: "val"})
	m.AddKeyAnyItemsPtr(nil)
	m.AddKeyAnyItemsPtr(ka)
	m.AddMapResults(nil)
	sub := corejson.NewMapResults.Empty()
	sub.Add("sub", corejson.NewResult.Any("v"))
	m.AddMapResults(sub)
	m.AddMapAnyItems(nil)
	m.AddMapAnyItems(map[string]any{"o": "val"})
}

func Test_C04_MapResults_Errors(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(3)
	m.Add("ok", corejson.NewResult.Any("x"))
	m.Add("err", corejson.Result{Error: errors.New("e1")})
	if !m.HasError() { t.Fatal("expected error") }
	errs, has := m.AllErrors()
	if !has || len(errs) != 1 { t.Fatal("expected 1") }
	_ = m.GetErrorsStrings()
	_ = m.GetErrorsStringsPtr()
	_ = m.GetErrorsAsSingleString()
	_ = m.GetErrorsAsSingle()
}

func Test_C04_MapResults_AllKeys(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("b", corejson.NewResult.Any("x"))
	m.Add("a", corejson.NewResult.Any("y"))
	if len(m.AllKeys()) != 2 { t.Fatal("expected 2") }
	sorted := m.AllKeysSorted()
	if sorted[0] != "a" { t.Fatal("expected a first") }
	if len(m.AllValues()) != 2 { t.Fatal("expected 2") }
	_ = m.AllResults()
	_ = m.AllResultsCollection()
}

func Test_C04_MapResults_Paging(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(10)
	for i := 0; i < 10; i++ {
		m.Add(corejson.Serialize.ToString(i), corejson.NewResult.Any(i))
	}
	if m.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if m.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	paged := m.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4") }
}

func Test_C04_MapResults_ClearDispose(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("x"))
	m.Clear()
	time.Sleep(10 * time.Millisecond)
	if m.Length() != 0 { t.Fatal("expected 0") }
	m.Dispose()
	var nilM *corejson.MapResults
	nilM.Clear()
	nilM.Dispose()
}

func Test_C04_MapResults_Json(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(1)
	m.Add("a", corejson.NewResult.Any("x"))
	_ = m.JsonModel()
	_ = m.JsonModelAny()
	_ = m.Json()
	_ = m.JsonPtr()
	_ = m.AsJsonContractsBinder()
	_ = m.AsJsoner()
	_ = m.AsJsonParseSelfInjector()
}

func Test_C04_MapResults_ResultCollection(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(1)
	m.Add("a", corejson.NewResult.Any("x"))
	rc := m.ResultCollection()
	if rc.Length() != 1 { t.Fatal("expected 1") }
}

func Test_C04_MapResults_GetStrings(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(1)
	m.Add("a", corejson.NewResult.Any("x"))
	if len(m.GetStrings()) != 1 { t.Fatal("expected 1") }
	_ = m.GetStringsPtr()
}

func Test_C04_MapResults_AddMapResultsUsingCloneOption(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	items := map[string]corejson.Result{"a": corejson.NewResult.Any("x")}
	m.AddMapResultsUsingCloneOption(false, false, items)
	m2 := corejson.NewMapResults.UsingCap(2)
	m2.AddMapResultsUsingCloneOption(true, true, items)
	m3 := corejson.NewMapResults.UsingCap(2)
	m3.AddMapResultsUsingCloneOption(false, false, nil)
}

func Test_C04_MapResults_GetNewMapUsingKeys(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("x"))
	m.Add("b", corejson.NewResult.Any("y"))
	sub := m.GetNewMapUsingKeys(false, "a")
	if sub.Length() != 1 { t.Fatal("expected 1") }
}

func Test_C04_MapResults_Creators(t *testing.T) {
	_ = corejson.NewMapResults.UsingKeyAnyItems(0, corejson.KeyAny{Key: "a", AnyInf: "x"})
	_ = corejson.NewMapResults.UsingKeyAnyItems(5)
	_ = corejson.NewMapResults.UsingMapPlusCap(5, nil)
	_ = corejson.NewMapResults.UsingMapPlusCapClone(5, nil)
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(5, nil)
	_ = corejson.NewMapResults.UsingMap(nil)
	_ = corejson.NewMapResults.UsingMapAnyItems(nil)
	_ = corejson.NewMapResults.UsingMapAnyItemsPlusCap(5, nil)
	_ = corejson.NewMapResults.UsingKeyWithResults(corejson.KeyWithResult{Key: "a", Result: corejson.NewResult.Any("x")})
	_ = corejson.NewMapResults.UsingKeyWithResultsPlusCap(5, corejson.KeyWithResult{Key: "a", Result: corejson.NewResult.Any("x")})
	_ = corejson.NewMapResults.UsingKeyWithResultsPlusCap(5)
	_ = corejson.NewMapResults.UsingMapOptions(false, false, 0, map[string]corejson.Result{"a": corejson.NewResult.Any("x")})
	_, _ = corejson.NewMapResults.UnmarshalUsingBytes([]byte(`{}`))
}

func Test_C04_MapResults_AddJsoner(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.AddJsoner("a", nil)
	_ = corejson.NewMapResults.UsingKeyJsoners()
	_ = corejson.NewMapResults.UsingKeyJsonersPlusCap(5)
}

func Test_C04_ResultsCollection_Creators(t *testing.T) {
	_ = corejson.NewResultsCollection.Serializers()
	_ = corejson.NewResultsCollection.SerializerFunctions()
	_ = corejson.NewResultsCollection.UsingJsoners()
	_ = corejson.NewResultsCollection.UsingJsonersNonNull(5)
}

func Test_C04_BytesCollection_Creators(t *testing.T) {
	_ = corejson.NewBytesCollection.JsonersPlusCap(true, 0)
	_ = corejson.NewBytesCollection.Serializers()
	_, _ = corejson.NewBytesCollection.UnmarshalUsingBytes([]byte(`{}`))
}
