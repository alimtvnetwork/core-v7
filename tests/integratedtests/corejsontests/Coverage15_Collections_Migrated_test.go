package corejsontests

import (
	"errors"
	"testing"
	"time"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Migrated from Coverage04, 08, 09, 10, 12, 17 — Collections & MapResults ──

func Test_C04_ResultsCollection_BasicOps(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	actual := args.Map{"result": c.IsEmpty() || c.HasAnyItem() || c.Length() != 0}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
	actual := args.Map{"result": c.LastIndex() != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
	c.Add(corejson.NewResult.Any("a"))
	c.Add(corejson.NewResult.Any("b"))
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": c.FirstOrDefault() == nil || c.LastOrDefault() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C04_ResultsCollection_TakeSkipLimit(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(5)
	for i := 0; i < 5; i++ { c.Add(corejson.NewResult.Any(i)) }
	actual := args.Map{"result": c.Take(3).Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": c.Skip(2).Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": c.Limit(3).Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": c.Limit(-2).Length() != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected all", actual)
	empty := corejson.NewResultsCollection.Empty()
	actual := args.Map{"result": empty.Take(1).Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
	actual := args.Map{"result": c.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	errs, has := c.AllErrors()
	actual := args.Map{"result": has || len(errs) != 1}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 1 error", actual)
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
	actual := args.Map{"result": err != nil || s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C04_ResultsCollection_GetAtSafe(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(1)
	c.Add(corejson.NewResult.Any("x"))
	actual := args.Map{"result": c.GetAtSafe(0) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": c.GetAtSafe(-1) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": c.GetAtSafe(5) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C04_ResultsCollection_Paging(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(10)
	for i := 0; i < 10; i++ { c.Add(corejson.NewResult.Any(i)) }
	actual := args.Map{"result": c.GetPagesSize(3) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": c.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	paged := c.GetPagedCollection(3)
	actual := args.Map{"result": len(paged) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 pages", actual)
	single := c.GetSinglePageCollection(3, 1)
	actual := args.Map{"result": single.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
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
	actual := args.Map{"result": nilC.ClonePtr(true) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C04_ResultsCollection_ClearDispose(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(2)
	c.Add(corejson.NewResult.Any("x"))
	c.Clear()
	time.Sleep(10 * time.Millisecond)
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	c.Dispose()
	var nilC *corejson.ResultsCollection
	nilC.Clear()
	nilC.Dispose()
}

func Test_C04_ResultsCollection_GetStrings(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(2)
	c.Add(corejson.NewResult.Any("a"))
	actual := args.Map{"result": len(c.GetStrings()) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = c.GetStringsPtr()
}

func Test_C04_ResultsCollection_Nil(t *testing.T) {
	var nilC *corejson.ResultsCollection
	actual := args.Map{"result": nilC.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": nilC.LastIndex() != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
	actual := args.Map{"result": nilC.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": nilC.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": nilC.FirstOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": nilC.LastOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// ── BytesCollection ──

func Test_C04_BytesCollection_BasicOps(t *testing.T) {
	c := corejson.NewBytesCollection.Empty()
	actual := args.Map{"result": c.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
	c.Add([]byte("hello"))
	actual := args.Map{"result": c.Length() != 1 || !c.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual := args.Map{"result": c.FirstOrDefault() == nil || c.LastOrDefault() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
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
	actual := args.Map{"result": c.Take(2).Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": c.Skip(1).Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": c.Limit(2).Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": c.Limit(-1).Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C04_BytesCollection_ClearDispose(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte("x"))
	c.Clear()
	time.Sleep(10 * time.Millisecond)
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
	actual := args.Map{"result": nilC.ClonePtr(true) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C04_BytesCollection_Json(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"x"`))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	b, err := c.MarshalJSON()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func Test_C04_BytesCollection_Paging(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ { c.Add([]byte(`"x"`)) }
	actual := args.Map{"result": c.GetPagesSize(3) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": c.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	paged := c.GetPagedCollection(3)
	actual := args.Map{"result": len(paged) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C04_BytesCollection_Nil(t *testing.T) {
	var nilC *corejson.BytesCollection
	actual := args.Map{"result": nilC.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": nilC.LastIndex() != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
	actual := args.Map{"result": nilC.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": nilC.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": nilC.FirstOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": nilC.LastOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C04_BytesCollection_AddResult(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	r := corejson.NewResult.Any("hello")
	c.AddResult(r)
	c.AddResultPtr(nil)
	c.AddResultPtr(&r)
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C04_BytesCollection_AddAny(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	err := c.AddAny("hello")
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	err2 := c.AddAnyItems("a", "b")
	actual := args.Map{"result": err2}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err2", actual)
}

func Test_C04_BytesCollection_GetAtSafe(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"x"`))
	actual := args.Map{"result": c.GetAtSafe(0) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": c.GetAtSafe(-1) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": c.GetAtSafe(5) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": c.GetAtSafePtr(0) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": c.GetResultAtSafe(0) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": c.GetResultAtSafe(5) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C04_BytesCollection_Strings(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	actual := args.Map{"result": len(c.Strings()) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
	actual := args.Map{"result": err != nil || s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C04_BytesCollection_AddBytesCollection(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	c2 := corejson.NewBytesCollection.UsingCap(1)
	c2.Add([]byte(`"b"`))
	c.AddBytesCollection(c2)
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ── ResultsPtrCollection ──

func Test_C04_ResultsPtrCollection_BasicOps(t *testing.T) {
	var nilC *corejson.ResultsPtrCollection
	actual := args.Map{"result": nilC.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": nilC.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": nilC.FirstOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": nilC.LastOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("hello"))
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual := args.Map{"result": c.FirstOrDefault() == nil || c.LastOrDefault() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
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
	actual := args.Map{"result": c.Take(3).Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": c.Skip(2).Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": c.Limit(3).Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": c.Limit(-2).Length() != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected all", actual)
}

func Test_C04_ResultsPtrCollection_Errors(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(2)
	c.Add(corejson.NewResult.AnyPtr("ok"))
	c.Add(&corejson.Result{Error: errors.New("e")})
	actual := args.Map{"result": c.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	errs, has := c.AllErrors()
	actual := args.Map{"result": has || len(errs) != 1}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
	actual := args.Map{"result": nilC.Clone(true) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
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
	actual := args.Map{"result": c.GetPagesSize(3) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": c.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	paged := c.GetPagedCollection(3)
	actual := args.Map{"result": len(paged) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C04_ResultsPtrCollection_GetStrings(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(2)
	c.Add(corejson.NewResult.AnyPtr("a"))
	actual := args.Map{"result": len(c.GetStrings()) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
	actual := args.Map{"result": nilM.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": nilM.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": nilM.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)

	m := corejson.NewMapResults.Empty()
	m.Add("a", corejson.NewResult.Any("hello"))
	actual := args.Map{"result": m.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	r := m.GetByKey("a")
	actual := args.Map{"result": r == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": m.GetByKey("missing") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
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
	actual := args.Map{"result": m.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	errs, has := m.AllErrors()
	actual := args.Map{"result": has || len(errs) != 1}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = m.GetErrorsStrings()
	_ = m.GetErrorsStringsPtr()
	_ = m.GetErrorsAsSingleString()
	_ = m.GetErrorsAsSingle()
}

func Test_C04_MapResults_AllKeys(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("b", corejson.NewResult.Any("x"))
	m.Add("a", corejson.NewResult.Any("y"))
	actual := args.Map{"result": len(m.AllKeys()) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	sorted := m.AllKeysSorted()
	actual := args.Map{"result": sorted[0] != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a first", actual)
	actual := args.Map{"result": len(m.AllValues()) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	_ = m.AllResults()
	_ = m.AllResultsCollection()
}

func Test_C04_MapResults_Paging(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(10)
	for i := 0; i < 10; i++ {
		m.Add(corejson.Serialize.ToString(i), corejson.NewResult.Any(i))
	}
	actual := args.Map{"result": m.GetPagesSize(3) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": m.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	paged := m.GetPagedCollection(3)
	actual := args.Map{"result": len(paged) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C04_MapResults_ClearDispose(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("x"))
	m.Clear()
	time.Sleep(10 * time.Millisecond)
	actual := args.Map{"result": m.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
	actual := args.Map{"result": rc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C04_MapResults_GetStrings(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(1)
	m.Add("a", corejson.NewResult.Any("x"))
	actual := args.Map{"result": len(m.GetStrings()) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
	actual := args.Map{"result": sub.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
