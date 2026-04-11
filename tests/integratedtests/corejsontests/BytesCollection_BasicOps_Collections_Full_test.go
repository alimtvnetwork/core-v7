package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// BytesCollection — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_BytesCollection_BasicOps(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(5)

	// Act
	actual := args.Map{"result": c.Length() != 0 || !c.IsEmpty() || c.HasAnyItem() || c.LastIndex() != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "basic empty checks failed", actual)
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`)).Add([]byte(`"c"`))
	actual = args.Map{"result": c.Length() != 3 || c.IsEmpty() || !c.HasAnyItem() || c.LastIndex() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "basic filled checks failed", actual)
}

func Test_BytesCollection_FirstLastOrDefault(t *testing.T) {
	// Arrange
	empty := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{"result": empty.FirstOrDefault() != nil || empty.LastOrDefault() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`))
	actual = args.Map{"result": string(c.FirstOrDefault()) != `"a"` || string(c.LastOrDefault()) != `"b"`}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "first/last wrong", actual)
}

func Test_BytesCollection_TakeLimitSkip(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(5)
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`)).Add([]byte(`"c"`)).Add([]byte(`"d"`))
	tk := c.Take(2)

	// Act
	actual := args.Map{"result": tk.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "take wrong", actual)
	lm := c.Limit(2)
	actual = args.Map{"result": lm.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "limit wrong", actual)
	lmAll := c.Limit(-1)
	actual = args.Map{"result": lmAll.Length() != 4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "limit all wrong", actual)
	sk := c.Skip(2)
	actual = args.Map{"result": sk.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "skip wrong", actual)
}

func Test_BytesCollection_AddMethods(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": c.Length() != 7}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 7", actual)
}

func Test_BytesCollection_GetAt(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`))

	// Act
	actual := args.Map{"result": string(c.GetAt(0)) != `"a"`}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetAt wrong", actual)
}

func Test_BytesCollection_JsonResultAt(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"hello"`))
	r := c.JsonResultAt(0)

	// Act
	actual := args.Map{"result": r == nil || r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid result", actual)
}

func Test_BytesCollection_UnmarshalAt(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"hello"`))
	var s string
	err := c.UnmarshalAt(0, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_BytesCollection_AddSerializerFunc(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	c.AddSerializerFunc(nil)
	c.AddSerializerFunc(func() ([]byte, error) { return []byte(`"x"`), nil })

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_BytesCollection_AddSerializerFunctions(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	c.AddSerializerFunctions(
		func() ([]byte, error) { return []byte(`"a"`), nil },
		func() ([]byte, error) { return []byte(`"b"`), nil },
	)

	// Act
	actual := args.Map{"result": c.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_BytesCollection_AddAnyItems(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(5)
	err := c.AddAnyItems("a", "b")

	// Act
	actual := args.Map{"result": err != nil || c.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_BytesCollection_AddAny(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	err := c.AddAny("hello")

	// Act
	actual := args.Map{"result": err != nil || c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_BytesCollection_AddsPtr(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(5)
	r1 := corejson.NewResult.AnyPtr("a")
	c.AddsPtr(r1, nil)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_BytesCollection_AddBytesCollection(t *testing.T) {
	// Arrange
	c1 := corejson.NewBytesCollection.UsingCap(2)
	c1.Add([]byte(`"a"`))
	c2 := corejson.NewBytesCollection.UsingCap(2)
	c2.Add([]byte(`"b"`))
	c1.AddBytesCollection(c2)

	// Act
	actual := args.Map{"result": c1.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_BytesCollection_GetAtSafe(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))

	// Act
	actual := args.Map{"result": c.GetAtSafe(0) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": c.GetAtSafe(5) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": c.GetAtSafe(-1) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for neg", actual)
}

func Test_BytesCollection_GetAtSafePtr(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))

	// Act
	actual := args.Map{"result": c.GetAtSafePtr(0) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_BytesCollection_GetResultAtSafe(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))

	// Act
	actual := args.Map{"result": c.GetResultAtSafe(0) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": c.GetResultAtSafe(5) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_BytesCollection_GetAtSafeUsingLength(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))

	// Act
	actual := args.Map{"result": c.GetAtSafeUsingLength(0, 1) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": c.GetAtSafeUsingLength(5, 1) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_BytesCollection_Strings(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`))
	strs := c.Strings()

	// Act
	actual := args.Map{"result": len(strs) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_BytesCollection_StringsPtr(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))
	strs := c.StringsPtr()

	// Act
	actual := args.Map{"result": len(strs) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_BytesCollection_ClearDispose(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	c.Clear()

	// Act
	actual := args.Map{"result": c.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 after clear", actual)
	c.Add([]byte(`"b"`))
	c.Dispose()
}

func Test_BytesCollection_PagesSize(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		c.Add([]byte(`"x"`))
	}

	// Act
	actual := args.Map{"result": c.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 pages", actual)
	actual = args.Map{"result": c.GetPagesSize(0) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for 0 size", actual)
}

func Test_BytesCollection_GetPagedCollection(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		c.Add([]byte(`"x"`))
	}
	pages := c.GetPagedCollection(2)

	// Act
	actual := args.Map{"result": len(pages) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)
}

func Test_BytesCollection_GetPagedCollection_SmallPage(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	pages := c.GetPagedCollection(5)

	// Act
	actual := args.Map{"result": len(pages) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 page", actual)
}

func Test_BytesCollection_GetSinglePageCollection(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		c.Add([]byte(`"x"`))
	}
	page := c.GetSinglePageCollection(3, 2)

	// Act
	actual := args.Map{"result": page.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	lastPage := c.GetSinglePageCollection(3, 4)
	actual = args.Map{"result": lastPage.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 for last page", actual)
}

func Test_BytesCollection_JsonOps(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	j := c.Json()

	// Act
	actual := args.Map{"result": j.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	jp := c.JsonPtr()
	actual = args.Map{"result": jp.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

func Test_BytesCollection_CloneOps(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`))
	sc := c.ShadowClone()
	_ = sc
	dc := c.Clone(true)
	_ = dc
	cp := c.ClonePtr(true)

	// Act
	actual := args.Map{"result": cp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected clone ptr", actual)
}

func Test_BytesCollection_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var c *corejson.BytesCollection
	cp := c.ClonePtr(true)

	// Act
	actual := args.Map{"result": cp != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_BytesCollection_InterfaceMethods(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func Test_BytesCollection_ParseInjectUsingJson(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"a"`))
	jr := c.JsonPtr()
	target := corejson.NewBytesCollection.Empty()
	_, err := target.ParseInjectUsingJson(jr)
	_ = err
}

func Test_BytesCollection_InjectIntoAt(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`["a","b"]`))
	target := corejson.NewResult.Any("x")
	err := c.InjectIntoAt(0, &target)
	_ = err
}

func Test_BytesCollection_InjectIntoSameIndex(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"hello"`))
	r := corejson.NewResult.Any("x")
	errs, hasErr := c.InjectIntoSameIndex(&r)
	_ = errs
	_ = hasErr
}

func Test_BytesCollection_InjectIntoSameIndex_Nil(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	// Pass true nil variadic to hit the nil early return
	var nilSlice []corejson.JsonParseSelfInjector
	errs, hasErr := c.InjectIntoSameIndex(nilSlice...)
	_ = errs
	_ = hasErr
}

func Test_BytesCollection_UnmarshalIntoSameIndex(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"hello"`))
	var s string
	errs, hasErr := c.UnmarshalIntoSameIndex(&s)
	_ = errs
	_ = hasErr
}

func Test_BytesCollection_UnmarshalIntoSameIndex_Nil(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(1)

	// Test true nil variadic (no args) - hits the anys==nil early return
	var nilSlice []any
	errs, hasErr := c.UnmarshalIntoSameIndex(nilSlice...)

	// Act
	actual := args.Map{"result": hasErr || len(errs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error for nil variadic", actual)

	// Test with a nil element but collection has an item
	c.AddAnyItems(true, "hello")
	errs2, hasErr2 := c.UnmarshalIntoSameIndex(nil)
	actual = args.Map{"result": hasErr2 || len(errs2) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 error slot with nil skip", actual)
}

func Test_BytesCollection_AddMapResults(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	c := corejson.NewBytesCollection.UsingCap(2)
	c.AddMapResults(mr)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_BytesCollection_AddRawMapResults(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	c.AddRawMapResults(map[string]corejson.Result{
		"k": corejson.NewResult.Any("v"),
	})

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_BytesCollection_AddJsoners(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	r := corejson.NewResult.Any("x")
	c.AddJsoners(true, &r)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ResultsCollection — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_ResultsCollection_BasicOps(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{"result": c.IsEmpty() || c.HasAnyItem() || c.Length() != 0 || c.LastIndex() != -1}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "basic checks failed", actual)
	c.Add(corejson.NewResult.Any("a")).Add(corejson.NewResult.Any("b"))
	actual = args.Map{"result": c.Length() != 2 || c.IsEmpty() || !c.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "filled checks failed", actual)
}

func Test_ResultsCollection_FirstLast(t *testing.T) {
	// Arrange
	e := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{"result": e.FirstOrDefault() != nil || e.LastOrDefault() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("a")).Add(corejson.NewResult.Any("b"))
	actual = args.Map{"result": c.FirstOrDefault() == nil || c.LastOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_ResultsCollection_TakeLimitSkip(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	for i := 0; i < 5; i++ {
		c.Add(corejson.NewResult.Any("x"))
	}

	// Act
	actual := args.Map{"result": c.Take(2).Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "take wrong", actual)
	actual = args.Map{"result": c.Limit(2).Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "limit wrong", actual)
	actual = args.Map{"result": c.Limit(-1).Length() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "limit all wrong", actual)
	actual = args.Map{"result": c.Skip(3).Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "skip wrong", actual)
}

func Test_ResultsCollection_AddMethods(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": c.Length() < 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 5", actual)
}

func Test_ResultsCollection_Errors(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("ok"))
	c.Add(corejson.NewResult.Error(errors.New("e")))

	// Act
	actual := args.Map{"result": c.HasError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	errs, hasErr := c.AllErrors()
	actual = args.Map{"result": hasErr || len(errs) == 0}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected errors", actual)
	strs := c.GetErrorsStrings()
	actual = args.Map{"result": len(strs) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error strings", actual)
	_ = c.GetErrorsStringsPtr()
	_ = c.GetErrorsAsSingleString()
	_ = c.GetErrorsAsSingle()
}

func Test_ResultsCollection_GetAtSafe(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))

	// Act
	actual := args.Map{"result": c.GetAtSafe(0) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": c.GetAtSafe(5) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": c.GetAtSafeUsingLength(0, 1) == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_ResultsCollection_PagingOps(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	for i := 0; i < 10; i++ {
		c.Add(corejson.NewResult.Any("x"))
	}

	// Act
	actual := args.Map{"result": c.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 pages", actual)
	pages := c.GetPagedCollection(3)
	actual = args.Map{"result": len(pages) != 4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 pages", actual)
}

func Test_ResultsCollection_ClearDispose(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	c.Clear()
	c.Add(corejson.NewResult.Any("y"))
	c.Dispose()
}

func Test_ResultsCollection_GetStrings(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	strs := c.GetStrings()

	// Act
	actual := args.Map{"result": len(strs) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = c.GetStringsPtr()
}

func Test_ResultsCollection_JsonOps(t *testing.T) {
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

func Test_ResultsCollection_CloneOps(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	sc := c.ShadowClone()
	_ = sc
	dc := c.Clone(true)
	_ = dc
	cp := c.ClonePtr(true)

	// Act
	actual := args.Map{"result": cp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected clone ptr", actual)
}

func Test_ResultsCollection_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var c *corejson.ResultsCollection
	cp := c.ClonePtr(true)

	// Act
	actual := args.Map{"result": cp != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_ResultsCollection_SerializerMethods(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.AddSerializerFunc(nil)
	c.AddSerializerFunc(func() ([]byte, error) { return []byte(`"x"`), nil })
	c.AddSerializerFunctions(
		func() ([]byte, error) { return []byte(`"a"`), nil },
	)

	// Act
	actual := args.Map{"result": c.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_ResultsCollection_AddMapResults(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	c := corejson.NewResultsCollection.Empty()
	c.AddMapResults(mr)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_ResultsCollection_AddResultsCollection(t *testing.T) {
	// Arrange
	c1 := corejson.NewResultsCollection.Empty()
	c1.Add(corejson.NewResult.Any("a"))
	c2 := corejson.NewResultsCollection.Empty()
	c2.AddResultsCollection(c1)

	// Act
	actual := args.Map{"result": c2.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_ResultsCollection_AddAnyItemsSlice(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.AddAnyItemsSlice([]any{"a", nil, "b"})

	// Act
	actual := args.Map{"result": c.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_ResultsCollection_AddJsoners(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	r := corejson.NewResult.Any("x")
	c.AddJsoners(true, &r)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_ResultsCollection_UnmarshalAt(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("hello"))
	var s string
	err := c.UnmarshalAt(0, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_ResultsCollection_InjectIntoSameIndex(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	r := corejson.NewResult.Any("y")
	errs, hasErr := c.InjectIntoSameIndex(&r)
	_ = errs
	_ = hasErr
}

func Test_ResultsCollection_UnmarshalIntoSameIndex(t *testing.T) {
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

func Test_ResultsPtrCollection_BasicOps(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()

	// Act
	actual := args.Map{"result": c.IsEmpty() || c.HasAnyItem() || c.Length() != 0}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "basic checks failed", actual)
	c.Add(corejson.NewResult.AnyPtr("a"))
	actual = args.Map{"result": c.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_ResultsPtrCollection_FirstLast(t *testing.T) {
	// Arrange
	e := corejson.NewResultsPtrCollection.Default()

	// Act
	actual := args.Map{"result": e.FirstOrDefault() != nil || e.LastOrDefault() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("a")).Add(corejson.NewResult.AnyPtr("b"))
	actual = args.Map{"result": c.FirstOrDefault() == nil || c.LastOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_ResultsPtrCollection_TakeLimitSkip(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	for i := 0; i < 5; i++ {
		c.Add(corejson.NewResult.AnyPtr("x"))
	}

	// Act
	actual := args.Map{"result": c.Take(2).Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "take wrong", actual)
	actual = args.Map{"result": c.Limit(2).Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "limit wrong", actual)
	actual = args.Map{"result": c.Limit(-1).Length() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "limit all wrong", actual)
	actual = args.Map{"result": c.Skip(3).Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "skip wrong", actual)
}

func Test_ResultsPtrCollection_AddMethods(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": c.Length() < 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 5", actual)
}

func Test_ResultsPtrCollection_Errors(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("ok"))
	c.Add(corejson.NewResult.ErrorPtr(errors.New("e")))

	// Act
	actual := args.Map{"result": c.HasError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	errs, hasErr := c.AllErrors()
	actual = args.Map{"result": hasErr || len(errs) == 0}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected errors", actual)
	_ = c.GetErrorsStrings()
	_ = c.GetErrorsStringsPtr()
	_ = c.GetErrorsAsSingleString()
	_ = c.GetErrorsAsSingle()
}

func Test_ResultsPtrCollection_PagingOps(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	for i := 0; i < 10; i++ {
		c.Add(corejson.NewResult.AnyPtr("x"))
	}

	// Act
	actual := args.Map{"result": c.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 pages", actual)
	pages := c.GetPagedCollection(3)
	actual = args.Map{"result": len(pages) != 4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_ResultsPtrCollection_ClearDispose(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	c.Clear()
	c.Add(corejson.NewResult.AnyPtr("y"))
	c.Dispose()
}

func Test_ResultsPtrCollection_GetStrings(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	strs := c.GetStrings()

	// Act
	actual := args.Map{"result": len(strs) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = c.GetStringsPtr()
}

func Test_ResultsPtrCollection_JsonOps(t *testing.T) {
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

func Test_ResultsPtrCollection_Clone(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	cp := c.Clone(true)

	// Act
	actual := args.Map{"result": cp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected clone", actual)
}

func Test_ResultsPtrCollection_Clone_Nil(t *testing.T) {
	// Arrange
	var c *corejson.ResultsPtrCollection
	cp := c.Clone(true)

	// Act
	actual := args.Map{"result": cp != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_ResultsPtrCollection_UnmarshalAt(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("hello"))
	var s string
	err := c.UnmarshalAt(0, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_ResultsPtrCollection_SerializerMethods(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	c.AddSerializerFunc(nil)
	c.AddSerializerFunc(func() ([]byte, error) { return []byte(`"x"`), nil })
	c.AddSerializerFunctions(
		func() ([]byte, error) { return []byte(`"a"`), nil },
	)

	// Act
	actual := args.Map{"result": c.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_ResultsPtrCollection_AddResultsCollection(t *testing.T) {
	// Arrange
	c1 := corejson.NewResultsPtrCollection.Default()
	c1.Add(corejson.NewResult.AnyPtr("a"))
	c2 := corejson.NewResultsPtrCollection.Default()
	c2.AddResultsCollection(c1)

	// Act
	actual := args.Map{"result": c2.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_ResultsPtrCollection_AddJsoners(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	r := corejson.NewResult.Any("x")
	c.AddJsoners(true, &r)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_ResultsPtrCollection_GetAtSafe(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))

	// Act
	actual := args.Map{"result": c.GetAtSafe(0) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": c.GetAtSafe(5) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": c.GetAtSafeUsingLength(0, 1) == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}
