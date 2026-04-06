package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ══════════════════════════════════════════════════════════════════════════════
// BytesCollection — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov19_BytesCollection_BasicOps(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	if c.Length() != 0 || !c.IsEmpty() || c.HasAnyItem() || c.LastIndex() != -1 {
		t.Fatal("basic empty checks failed")
	}
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`)).Add([]byte(`"c"`))
	if c.Length() != 3 || c.IsEmpty() || !c.HasAnyItem() || c.LastIndex() != 2 {
		t.Fatal("basic filled checks failed")
	}
}

func Test_Cov19_BytesCollection_FirstLastOrDefault(t *testing.T) {
	empty := corejson.NewBytesCollection.Empty()
	if empty.FirstOrDefault() != nil || empty.LastOrDefault() != nil {
		t.Fatal("expected nil for empty")
	}
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`))
	if string(c.FirstOrDefault()) != `"a"` || string(c.LastOrDefault()) != `"b"` {
		t.Fatal("first/last wrong")
	}
}

func Test_Cov19_BytesCollection_TakeLimitSkip(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`)).Add([]byte(`"c"`)).Add([]byte(`"d"`))
	tk := c.Take(2)
	if tk.Length() != 2 {
		t.Fatal("take wrong")
	}
	lm := c.Limit(2)
	if lm.Length() != 2 {
		t.Fatal("limit wrong")
	}
	lmAll := c.Limit(-1)
	if lmAll.Length() != 4 {
		t.Fatal("limit all wrong")
	}
	sk := c.Skip(2)
	if sk.Length() != 2 {
		t.Fatal("skip wrong")
	}
}

func Test_Cov19_BytesCollection_AddMethods(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(10)
	c.AddSkipOnNil(nil)
	c.AddSkipOnNil([]byte(`"x"`))
	c.AddNonEmpty([]byte{})
	c.AddNonEmpty([]byte(`"y"`))
	r := corejson.NewResult.AnyPtr("z")
	c.AddResultPtr(r)
	c.AddResult(corejson.NewResult.Any("w"))
	c.AddPtr([]byte{})
	c.AddPtr([]byte(`"q"`))
	c.Adds([]byte(`"a"`), []byte(`"b"`))
	if c.Length() != 7 {
		t.Fatalf("expected 7, got %d", c.Length())
	}
}

func Test_Cov19_BytesCollection_GetAt(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`))
	if string(c.GetAt(0)) != `"a"` {
		t.Fatal("GetAt wrong")
	}
}

func Test_Cov19_BytesCollection_JsonResultAt(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"hello"`))
	r := c.JsonResultAt(0)
	if r == nil || r.HasError() {
		t.Fatal("expected valid result")
	}
}

func Test_Cov19_BytesCollection_UnmarshalAt(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"hello"`))
	var s string
	err := c.UnmarshalAt(0, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_Cov19_BytesCollection_AddSerializerFunc(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.AddSerializerFunc(nil)
	c.AddSerializerFunc(func() ([]byte, error) { return []byte(`"x"`), nil })
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov19_BytesCollection_AddSerializerFunctions(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.AddSerializerFunctions(
		func() ([]byte, error) { return []byte(`"a"`), nil },
		func() ([]byte, error) { return []byte(`"b"`), nil },
	)
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_Cov19_BytesCollection_AddAnyItems(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	err := c.AddAnyItems("a", "b")
	if err != nil || c.Length() != 2 {
		t.Fatal("unexpected")
	}
}

func Test_Cov19_BytesCollection_AddAny(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	err := c.AddAny("hello")
	if err != nil || c.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_Cov19_BytesCollection_AddsPtr(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	r1 := corejson.NewResult.AnyPtr("a")
	c.AddsPtr(r1, nil)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov19_BytesCollection_AddBytesCollection(t *testing.T) {
	c1 := corejson.NewBytesCollection.UsingCap(2)
	c1.Add([]byte(`"a"`))
	c2 := corejson.NewBytesCollection.UsingCap(2)
	c2.Add([]byte(`"b"`))
	c1.AddBytesCollection(c2)
	if c1.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_Cov19_BytesCollection_GetAtSafe(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	if c.GetAtSafe(0) == nil {
		t.Fatal("expected non-nil")
	}
	if c.GetAtSafe(5) != nil {
		t.Fatal("expected nil")
	}
	if c.GetAtSafe(-1) != nil {
		t.Fatal("expected nil for neg")
	}
}

func Test_Cov19_BytesCollection_GetAtSafePtr(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))
	if c.GetAtSafePtr(0) == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov19_BytesCollection_GetResultAtSafe(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))
	if c.GetResultAtSafe(0) == nil {
		t.Fatal("expected non-nil")
	}
	if c.GetResultAtSafe(5) != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov19_BytesCollection_GetAtSafeUsingLength(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))
	if c.GetAtSafeUsingLength(0, 1) == nil {
		t.Fatal("expected non-nil")
	}
	if c.GetAtSafeUsingLength(5, 1) != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov19_BytesCollection_Strings(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`))
	strs := c.Strings()
	if len(strs) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_Cov19_BytesCollection_StringsPtr(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))
	strs := c.StringsPtr()
	if len(strs) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov19_BytesCollection_ClearDispose(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	c.Clear()
	if c.Length() != 0 {
		t.Fatal("expected 0 after clear")
	}
	c.Add([]byte(`"b"`))
	c.Dispose()
}

func Test_Cov19_BytesCollection_PagesSize(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		c.Add([]byte(`"x"`))
	}
	if c.GetPagesSize(3) != 4 {
		t.Fatal("expected 4 pages")
	}
	if c.GetPagesSize(0) != 0 {
		t.Fatal("expected 0 for 0 size")
	}
}

func Test_Cov19_BytesCollection_GetPagedCollection(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		c.Add([]byte(`"x"`))
	}
	pages := c.GetPagedCollection(2)
	if len(pages) != 3 {
		t.Fatalf("expected 3 pages, got %d", len(pages))
	}
}

func Test_Cov19_BytesCollection_GetPagedCollection_SmallPage(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	pages := c.GetPagedCollection(5)
	if len(pages) != 1 {
		t.Fatal("expected 1 page")
	}
}

func Test_Cov19_BytesCollection_GetSinglePageCollection(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		c.Add([]byte(`"x"`))
	}
	page := c.GetSinglePageCollection(3, 2)
	if page.Length() != 3 {
		t.Fatalf("expected 3, got %d", page.Length())
	}
	lastPage := c.GetSinglePageCollection(3, 4)
	if lastPage.Length() != 1 {
		t.Fatalf("expected 1 for last page, got %d", lastPage.Length())
	}
}

func Test_Cov19_BytesCollection_JsonOps(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	j := c.Json()
	if j.HasError() {
		t.Fatal("unexpected error")
	}
	jp := c.JsonPtr()
	if jp.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_Cov19_BytesCollection_CloneOps(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`))
	sc := c.ShadowClone()
	_ = sc
	dc := c.Clone(true)
	_ = dc
	cp := c.ClonePtr(true)
	if cp == nil {
		t.Fatal("expected clone ptr")
	}
}

func Test_Cov19_BytesCollection_ClonePtr_Nil(t *testing.T) {
	var c *corejson.BytesCollection
	cp := c.ClonePtr(true)
	if cp != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov19_BytesCollection_InterfaceMethods(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func Test_Cov19_BytesCollection_ParseInjectUsingJson(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))
	jr := c.JsonPtr()
	target := corejson.NewBytesCollection.Empty()
	_, err := target.ParseInjectUsingJson(jr)
	_ = err
}

func Test_Cov19_BytesCollection_InjectIntoAt(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`["a","b"]`))
	target := corejson.NewResult.Any("x")
	err := c.InjectIntoAt(0, &target)
	_ = err
}

func Test_Cov19_BytesCollection_InjectIntoSameIndex(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"hello"`))
	r := corejson.NewResult.Any("x")
	errs, hasErr := c.InjectIntoSameIndex(&r)
	_ = errs
	_ = hasErr
}

func Test_Cov19_BytesCollection_InjectIntoSameIndex_Nil(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	// Pass true nil variadic to hit the nil early return
	var nilSlice []corejson.JsonParseSelfInjector
	errs, hasErr := c.InjectIntoSameIndex(nilSlice...)
	_ = errs
	_ = hasErr
}

func Test_Cov19_BytesCollection_UnmarshalIntoSameIndex(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"hello"`))
	var s string
	errs, hasErr := c.UnmarshalIntoSameIndex(&s)
	_ = errs
	_ = hasErr
}

func Test_Cov19_BytesCollection_UnmarshalIntoSameIndex_Nil(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)

	// Test true nil variadic (no args) - hits the anys==nil early return
	var nilSlice []any
	errs, hasErr := c.UnmarshalIntoSameIndex(nilSlice...)
	if hasErr || len(errs) != 0 {
		t.Fatal("expected no error for nil variadic")
	}

	// Test with a nil element but collection has an item
	c.AddAnyItems(true, "hello")
	errs2, hasErr2 := c.UnmarshalIntoSameIndex(nil)
	if hasErr2 || len(errs2) != 1 {
		t.Fatal("expected 1 error slot with nil skip")
	}
}

func Test_Cov19_BytesCollection_AddMapResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	c := corejson.NewBytesCollection.UsingCap(2)
	c.AddMapResults(mr)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov19_BytesCollection_AddRawMapResults(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.AddRawMapResults(map[string]corejson.Result{
		"k": corejson.NewResult.Any("v"),
	})
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov19_BytesCollection_AddJsoners(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	r := corejson.NewResult.Any("x")
	c.AddJsoners(true, &r)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// ResultsCollection — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov19_ResultsCollection_BasicOps(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	if !c.IsEmpty() || c.HasAnyItem() || c.Length() != 0 || c.LastIndex() != -1 {
		t.Fatal("basic checks failed")
	}
	c.Add(corejson.NewResult.Any("a")).Add(corejson.NewResult.Any("b"))
	if c.Length() != 2 || c.IsEmpty() || !c.HasAnyItem() {
		t.Fatal("filled checks failed")
	}
}

func Test_Cov19_ResultsCollection_FirstLast(t *testing.T) {
	e := corejson.NewResultsCollection.Empty()
	if e.FirstOrDefault() != nil || e.LastOrDefault() != nil {
		t.Fatal("expected nil for empty")
	}
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("a")).Add(corejson.NewResult.Any("b"))
	if c.FirstOrDefault() == nil || c.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov19_ResultsCollection_TakeLimitSkip(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	for i := 0; i < 5; i++ {
		c.Add(corejson.NewResult.Any("x"))
	}
	if c.Take(2).Length() != 2 {
		t.Fatal("take wrong")
	}
	if c.Limit(2).Length() != 2 {
		t.Fatal("limit wrong")
	}
	if c.Limit(-1).Length() != 5 {
		t.Fatal("limit all wrong")
	}
	if c.Skip(3).Length() != 2 {
		t.Fatal("skip wrong")
	}
}

func Test_Cov19_ResultsCollection_AddMethods(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddSkipOnNil(nil)
	r := corejson.NewResult.AnyPtr("x")
	c.AddSkipOnNil(r)
	c.AddNonNilNonError(nil)
	c.AddNonNilNonError(r)
	c.AddPtr(nil)
	c.AddPtr(r)
	c.Adds(corejson.NewResult.Any("y"))
	c.AddAny("z")
	c.AddAny(nil) // should skip
	c.AddAnyItems("a", nil, "b")
	c.AddsPtr(r, nil)
	c.AddNonNilItemsPtr(r, nil)
	if c.Length() < 5 {
		t.Fatal("expected at least 5")
	}
}

func Test_Cov19_ResultsCollection_Errors(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("ok"))
	c.Add(corejson.NewResult.Error(errors.New("e")))
	if !c.HasError() {
		t.Fatal("expected error")
	}
	errs, hasErr := c.AllErrors()
	if !hasErr || len(errs) == 0 {
		t.Fatal("expected errors")
	}
	strs := c.GetErrorsStrings()
	if len(strs) == 0 {
		t.Fatal("expected error strings")
	}
	_ = c.GetErrorsStringsPtr()
	_ = c.GetErrorsAsSingleString()
	_ = c.GetErrorsAsSingle()
}

func Test_Cov19_ResultsCollection_GetAtSafe(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	if c.GetAtSafe(0) == nil {
		t.Fatal("expected non-nil")
	}
	if c.GetAtSafe(5) != nil {
		t.Fatal("expected nil")
	}
	if c.GetAtSafeUsingLength(0, 1) == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov19_ResultsCollection_PagingOps(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	for i := 0; i < 10; i++ {
		c.Add(corejson.NewResult.Any("x"))
	}
	if c.GetPagesSize(3) != 4 {
		t.Fatal("expected 4 pages")
	}
	pages := c.GetPagedCollection(3)
	if len(pages) != 4 {
		t.Fatalf("expected 4 pages, got %d", len(pages))
	}
}

func Test_Cov19_ResultsCollection_ClearDispose(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	c.Clear()
	c.Add(corejson.NewResult.Any("y"))
	c.Dispose()
}

func Test_Cov19_ResultsCollection_GetStrings(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	strs := c.GetStrings()
	if len(strs) != 1 {
		t.Fatal("expected 1")
	}
	_ = c.GetStringsPtr()
}

func Test_Cov19_ResultsCollection_JsonOps(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
	_ = c.NonPtr()
	_ = c.Ptr()
}

func Test_Cov19_ResultsCollection_CloneOps(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	sc := c.ShadowClone()
	_ = sc
	dc := c.Clone(true)
	_ = dc
	cp := c.ClonePtr(true)
	if cp == nil {
		t.Fatal("expected clone ptr")
	}
}

func Test_Cov19_ResultsCollection_ClonePtr_Nil(t *testing.T) {
	var c *corejson.ResultsCollection
	cp := c.ClonePtr(true)
	if cp != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov19_ResultsCollection_SerializerMethods(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddSerializerFunc(nil)
	c.AddSerializerFunc(func() ([]byte, error) { return []byte(`"x"`), nil })
	c.AddSerializerFunctions(
		func() ([]byte, error) { return []byte(`"a"`), nil },
	)
	if c.Length() != 2 {
		t.Fatalf("expected 2, got %d", c.Length())
	}
}

func Test_Cov19_ResultsCollection_AddMapResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	c := corejson.NewResultsCollection.Empty()
	c.AddMapResults(mr)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov19_ResultsCollection_AddResultsCollection(t *testing.T) {
	c1 := corejson.NewResultsCollection.Empty()
	c1.Add(corejson.NewResult.Any("a"))
	c2 := corejson.NewResultsCollection.Empty()
	c2.AddResultsCollection(c1)
	if c2.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov19_ResultsCollection_AddAnyItemsSlice(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddAnyItemsSlice([]any{"a", nil, "b"})
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_Cov19_ResultsCollection_AddJsoners(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	r := corejson.NewResult.Any("x")
	c.AddJsoners(true, &r)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov19_ResultsCollection_UnmarshalAt(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("hello"))
	var s string
	err := c.UnmarshalAt(0, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_Cov19_ResultsCollection_InjectIntoSameIndex(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	r := corejson.NewResult.Any("y")
	errs, hasErr := c.InjectIntoSameIndex(&r)
	_ = errs
	_ = hasErr
}

func Test_Cov19_ResultsCollection_UnmarshalIntoSameIndex(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("hello"))
	var s string
	errs, hasErr := c.UnmarshalIntoSameIndex(&s)
	_ = errs
	_ = hasErr
}

// ══════════════════════════════════════════════════════════════════════════════
// ResultsPtrCollection — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov19_ResultsPtrCollection_BasicOps(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	if !c.IsEmpty() || c.HasAnyItem() || c.Length() != 0 {
		t.Fatal("basic checks failed")
	}
	c.Add(corejson.NewResult.AnyPtr("a"))
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov19_ResultsPtrCollection_FirstLast(t *testing.T) {
	e := corejson.NewResultsPtrCollection.Default()
	if e.FirstOrDefault() != nil || e.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("a")).Add(corejson.NewResult.AnyPtr("b"))
	if c.FirstOrDefault() == nil || c.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov19_ResultsPtrCollection_TakeLimitSkip(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	for i := 0; i < 5; i++ {
		c.Add(corejson.NewResult.AnyPtr("x"))
	}
	if c.Take(2).Length() != 2 {
		t.Fatal("take wrong")
	}
	if c.Limit(2).Length() != 2 {
		t.Fatal("limit wrong")
	}
	if c.Limit(-1).Length() != 5 {
		t.Fatal("limit all wrong")
	}
	if c.Skip(3).Length() != 2 {
		t.Fatal("skip wrong")
	}
}

func Test_Cov19_ResultsPtrCollection_AddMethods(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.AddSkipOnNil(nil)
	r := corejson.NewResult.AnyPtr("x")
	c.AddSkipOnNil(r)
	c.AddNonNilNonError(nil)
	c.AddNonNilNonError(r)
	c.Adds(r, nil)
	c.AddAny("z")
	c.AddAny(nil)
	c.AddAnyItems("a", nil, "b")
	c.AddResult(corejson.NewResult.Any("w"))
	c.AddNonNilItems(r, nil)
	c.AddNonNilItemsPtr(r, nil)
	if c.Length() < 5 {
		t.Fatal("expected at least 5")
	}
}

func Test_Cov19_ResultsPtrCollection_Errors(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("ok"))
	c.Add(corejson.NewResult.ErrorPtr(errors.New("e")))
	if !c.HasError() {
		t.Fatal("expected error")
	}
	errs, hasErr := c.AllErrors()
	if !hasErr || len(errs) == 0 {
		t.Fatal("expected errors")
	}
	_ = c.GetErrorsStrings()
	_ = c.GetErrorsStringsPtr()
	_ = c.GetErrorsAsSingleString()
	_ = c.GetErrorsAsSingle()
}

func Test_Cov19_ResultsPtrCollection_PagingOps(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	for i := 0; i < 10; i++ {
		c.Add(corejson.NewResult.AnyPtr("x"))
	}
	if c.GetPagesSize(3) != 4 {
		t.Fatal("expected 4 pages")
	}
	pages := c.GetPagedCollection(3)
	if len(pages) != 4 {
		t.Fatalf("expected 4, got %d", len(pages))
	}
}

func Test_Cov19_ResultsPtrCollection_ClearDispose(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	c.Clear()
	c.Add(corejson.NewResult.AnyPtr("y"))
	c.Dispose()
}

func Test_Cov19_ResultsPtrCollection_GetStrings(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	strs := c.GetStrings()
	if len(strs) != 1 {
		t.Fatal("expected 1")
	}
	_ = c.GetStringsPtr()
}

func Test_Cov19_ResultsPtrCollection_JsonOps(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
	_ = c.NonPtr()
	_ = c.Ptr()
}

func Test_Cov19_ResultsPtrCollection_Clone(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	cp := c.Clone(true)
	if cp == nil {
		t.Fatal("expected clone")
	}
}

func Test_Cov19_ResultsPtrCollection_Clone_Nil(t *testing.T) {
	var c *corejson.ResultsPtrCollection
	cp := c.Clone(true)
	if cp != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov19_ResultsPtrCollection_UnmarshalAt(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("hello"))
	var s string
	err := c.UnmarshalAt(0, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_Cov19_ResultsPtrCollection_SerializerMethods(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.AddSerializerFunc(nil)
	c.AddSerializerFunc(func() ([]byte, error) { return []byte(`"x"`), nil })
	c.AddSerializerFunctions(
		func() ([]byte, error) { return []byte(`"a"`), nil },
	)
	if c.Length() != 2 {
		t.Fatalf("expected 2, got %d", c.Length())
	}
}

func Test_Cov19_ResultsPtrCollection_AddResultsCollection(t *testing.T) {
	c1 := corejson.NewResultsPtrCollection.Default()
	c1.Add(corejson.NewResult.AnyPtr("a"))
	c2 := corejson.NewResultsPtrCollection.Default()
	c2.AddResultsCollection(c1)
	if c2.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov19_ResultsPtrCollection_AddJsoners(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	r := corejson.NewResult.Any("x")
	c.AddJsoners(true, &r)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov19_ResultsPtrCollection_GetAtSafe(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	if c.GetAtSafe(0) == nil {
		t.Fatal("expected non-nil")
	}
	if c.GetAtSafe(5) != nil {
		t.Fatal("expected nil")
	}
	if c.GetAtSafeUsingLength(0, 1) == nil {
		t.Fatal("expected non-nil")
	}
}
