package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Collection.go — Seg-01: Lines 27–700 (~200 uncovered stmts)
// Covers: JsonString, JsonStringMust, HasAnyItem, LastIndex, HasIndex,
//         ListStringsPtr, ListStrings, StringJSON, RemoveAt, Count, Capacity,
//         Length, LengthLock, IsEquals, isCollectionPrecheckEqual,
//         IsEqualsWithSensitive, IsEmptyLock, IsEmpty, HasItems,
//         AddLock, AddNonEmpty, AddNonEmptyWhitespace, Add, AddError,
//         AsDefaultError, AsError, AddIf, EachItemSplitBy, ConcatNew,
//         ToError, ToDefaultError, AddIfMany, AddFunc, AddFuncErr,
//         AddsLock, Adds, AddStrings, AddCollection, AddCollections,
//         AddPointerCollectionsLock, AddHashmapsValues, AddHashmapsKeys,
//         isResizeRequired, resizeForHashmaps, resizeForCollections,
//         resizeForItems, resizeForAnys, AddHashmapsKeysValues,
//         AddHashmapsKeysValuesUsingFilter, AddWithWgLock, IndexAt,
//         SafeIndexAtUsingLength, First, Single, Last, LastOrDefault,
//         FirstOrDefault, Take, Skip, Reverse, GetPagesSize,
//         GetPagedCollection, GetSinglePageCollection
// =============================================================================

func Test_Cov58_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_JsonString", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.JsonString()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JsonString returns non-empty", actual)
	})
}

func Test_Cov58_Collection_JsonStringMust(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_JsonStringMust", func() {
		c := corestr.New.Collection.Strings([]string{"x"})
		s := c.JsonStringMust()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JsonStringMust returns non-empty", actual)
	})
}

func Test_Cov58_Collection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_HasAnyItem", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		e := corestr.New.Collection.Empty()
		actual := args.Map{"has": c.HasAnyItem(), "empty": e.HasAnyItem()}
		expected := args.Map{"has": true, "empty": false}
		expected.ShouldBeEqual(t, 0, "HasAnyItem returns correct bool", actual)
	})
}

func Test_Cov58_Collection_LastIndex(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_LastIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		actual := args.Map{"lastIndex": c.LastIndex()}
		expected := args.Map{"lastIndex": 2}
		expected.ShouldBeEqual(t, 0, "LastIndex returns len-1", actual)
	})
}

func Test_Cov58_Collection_HasIndex(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_HasIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"valid": c.HasIndex(1), "invalid": c.HasIndex(5), "neg": c.HasIndex(-1)}
		expected := args.Map{"valid": true, "invalid": false, "neg": false}
		expected.ShouldBeEqual(t, 0, "HasIndex checks bounds correctly", actual)
	})
}

func Test_Cov58_Collection_ListStringsPtr(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_ListStringsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"len": len(c.ListStringsPtr())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListStringsPtr returns items", actual)
	})
}

func Test_Cov58_Collection_ListStrings(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_ListStrings", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"len": len(c.ListStrings())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListStrings returns items", actual)
	})
}

func Test_Cov58_Collection_StringJSON(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_StringJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.StringJSON()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringJSON returns non-empty", actual)
	})
}

func Test_Cov58_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_RemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok1 := c.RemoveAt(1)
		ok2 := c.RemoveAt(-1)
		ok3 := c.RemoveAt(100)
		actual := args.Map{"ok": ok1, "neg": ok2, "oob": ok3, "len": c.Length()}
		expected := args.Map{"ok": true, "neg": false, "oob": false, "len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveAt removes item at index", actual)
	})
}

func Test_Cov58_Collection_Count(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Count", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"count": c.Count()}
		expected := args.Map{"count": 2}
		expected.ShouldBeEqual(t, 0, "Count returns length", actual)
	})
}

func Test_Cov58_Collection_Capacity(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Capacity", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		e := corestr.New.Collection.Empty()
		actual := args.Map{"hasCapNonEmpty": c.Capacity() > 0, "emptyCapGte0": e.Capacity() >= 0}
		expected := args.Map{"hasCapNonEmpty": true, "emptyCapGte0": true}
		expected.ShouldBeEqual(t, 0, "Capacity returns cap", actual)
	})
}

func Test_Cov58_Collection_Length_NilItems(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Length_NilItems", func() {
		var c *corestr.Collection
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length on nil returns 0", actual)
	})
}

func Test_Cov58_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_LengthLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"len": c.LengthLock()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LengthLock returns length", actual)
	})
}

func Test_Cov58_Collection_IsEquals_SameContent(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_SameContent", func() {
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals returns true for same content", actual)
	})
}

func Test_Cov58_Collection_IsEquals_DiffContent(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_DiffContent", func() {
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "c"})
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals returns false for different content", actual)
	})
}

func Test_Cov58_Collection_IsEquals_NilBoth(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_NilBoth", func() {
		var a, b *corestr.Collection
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals nil==nil is true", actual)
	})
}

func Test_Cov58_Collection_IsEquals_NilOne(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_NilOne", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		var b *corestr.Collection
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals nil vs non-nil is false", actual)
	})
}

func Test_Cov58_Collection_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_DiffLength", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff length is false", actual)
	})
}

func Test_Cov58_Collection_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_BothEmpty", func() {
		a := corestr.New.Collection.Empty()
		b := corestr.New.Collection.Empty()
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty is true", actual)
	})
}

func Test_Cov58_Collection_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_SamePtr", func() {
		a := corestr.New.Collection.Strings([]string{"x"})
		actual := args.Map{"eq": a.IsEquals(a)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals same pointer is true", actual)
	})
}

func Test_Cov58_Collection_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEqualsWithSensitive_CaseInsensitive", func() {
		a := corestr.New.Collection.Strings([]string{"Hello", "World"})
		b := corestr.New.Collection.Strings([]string{"hello", "world"})
		actual := args.Map{"caseSensitive": a.IsEqualsWithSensitive(true, b), "caseInsensitive": a.IsEqualsWithSensitive(false, b)}
		expected := args.Map{"caseSensitive": false, "caseInsensitive": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsWithSensitive handles case", actual)
	})
}

func Test_Cov58_Collection_IsEqualsWithSensitive_CaseInsensitiveMismatch(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEqualsWithSensitive_CaseInsensitiveMismatch", func() {
		a := corestr.New.Collection.Strings([]string{"hello"})
		b := corestr.New.Collection.Strings([]string{"xyz"})
		actual := args.Map{"eq": a.IsEqualsWithSensitive(false, b)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualsWithSensitive case insensitive mismatch", actual)
	})
}

func Test_Cov58_Collection_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEmptyLock", func() {
		c := corestr.New.Collection.Empty()
		d := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"empty": c.IsEmptyLock(), "notEmpty": d.IsEmptyLock()}
		expected := args.Map{"empty": true, "notEmpty": false}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock returns correct bool", actual)
	})
}

func Test_Cov58_Collection_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_HasItems", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		e := corestr.New.Collection.Empty()
		actual := args.Map{"has": c.HasItems(), "empty": e.HasItems()}
		expected := args.Map{"has": true, "empty": false}
		expected.ShouldBeEqual(t, 0, "HasItems returns correct bool", actual)
	})
}

func Test_Cov58_Collection_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddLock", func() {
		c := corestr.New.Collection.Empty()
		c.AddLock("x")
		actual := args.Map{"len": c.Length(), "first": c.First()}
		expected := args.Map{"len": 1, "first": "x"}
		expected.ShouldBeEqual(t, 0, "AddLock adds item", actual)
	})
}

func Test_Cov58_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddNonEmpty", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmpty("a")
		c.AddNonEmpty("")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty skips empty strings", actual)
	})
}

func Test_Cov58_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddNonEmptyWhitespace", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyWhitespace("a")
		c.AddNonEmptyWhitespace("   ")
		c.AddNonEmptyWhitespace("")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyWhitespace skips whitespace", actual)
	})
}

func Test_Cov58_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddError", func() {
		c := corestr.New.Collection.Empty()
		c.AddError(nil)
		c.AddError(errForTest)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddError skips nil error", actual)
	})
}

func Test_Cov58_Collection_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AsDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"err1", "err2"})
		e := c.AsDefaultError()
		actual := args.Map{"nonNil": e != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsDefaultError returns error", actual)
	})
}

func Test_Cov58_Collection_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AsError_Empty", func() {
		c := corestr.New.Collection.Empty()
		e := c.AsError(",")
		actual := args.Map{"nil": e == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "AsError on empty returns nil", actual)
	})
}

func Test_Cov58_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddIf", func() {
		c := corestr.New.Collection.Empty()
		c.AddIf(true, "yes")
		c.AddIf(false, "no")
		actual := args.Map{"len": c.Length(), "first": c.First()}
		expected := args.Map{"len": 1, "first": "yes"}
		expected.ShouldBeEqual(t, 0, "AddIf conditionally adds", actual)
	})
}

func Test_Cov58_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_EachItemSplitBy", func() {
		c := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
		result := c.EachItemSplitBy(",")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "EachItemSplitBy splits items", actual)
	})
}

func Test_Cov58_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_ConcatNew", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		n := c.ConcatNew(0, "b", "c")
		actual := args.Map{"len": n.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ConcatNew creates new collection", actual)
	})
}

func Test_Cov58_Collection_ConcatNew_NoAdding(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_ConcatNew_NoAdding", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		n := c.ConcatNew(0)
		actual := args.Map{"len": n.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ConcatNew with no additions clones", actual)
	})
}

func Test_Cov58_Collection_ToError(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_ToError", func() {
		c := corestr.New.Collection.Strings([]string{"e1", "e2"})
		e := c.ToError(",")
		actual := args.Map{"nonNil": e != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ToError returns error", actual)
	})
}

func Test_Cov58_Collection_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_ToDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"e1"})
		e := c.ToDefaultError()
		actual := args.Map{"nonNil": e != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ToDefaultError returns error", actual)
	})
}

func Test_Cov58_Collection_AddIfMany(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddIfMany", func() {
		c := corestr.New.Collection.Empty()
		c.AddIfMany(true, "a", "b")
		c.AddIfMany(false, "c")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddIfMany adds conditionally", actual)
	})
}

func Test_Cov58_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddFunc", func() {
		c := corestr.New.Collection.Empty()
		c.AddFunc(func() string { return "val" })
		actual := args.Map{"len": c.Length(), "first": c.First()}
		expected := args.Map{"len": 1, "first": "val"}
		expected.ShouldBeEqual(t, 0, "AddFunc adds func result", actual)
	})
}

func Test_Cov58_Collection_AddFuncErr_Success(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddFuncErr_Success", func() {
		c := corestr.New.Collection.Empty()
		c.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)
		actual := args.Map{"len": c.Length(), "first": c.First()}
		expected := args.Map{"len": 1, "first": "ok"}
		expected.ShouldBeEqual(t, 0, "AddFuncErr adds on success", actual)
	})
}

func Test_Cov58_Collection_AddFuncErr_Error(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddFuncErr_Error", func() {
		c := corestr.New.Collection.Empty()
		handlerCalled := false
		c.AddFuncErr(
			func() (string, error) { return "", errForTest },
			func(err error) { handlerCalled = true },
		)
		actual := args.Map{"len": c.Length(), "handled": handlerCalled}
		expected := args.Map{"len": 0, "handled": true}
		expected.ShouldBeEqual(t, 0, "AddFuncErr calls handler on error", actual)
	})
}

func Test_Cov58_Collection_AddsLock(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddsLock", func() {
		c := corestr.New.Collection.Empty()
		c.AddsLock("a", "b")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsLock adds items", actual)
	})
}

func Test_Cov58_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		other := corestr.New.Collection.Strings([]string{"b", "c"})
		c.AddCollection(other)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AddCollection appends items", actual)
	})
}

func Test_Cov58_Collection_AddCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddCollection_Empty", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		empty := corestr.New.Collection.Empty()
		c.AddCollection(empty)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollection skips empty", actual)
	})
}

func Test_Cov58_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddCollections", func() {
		c := corestr.New.Collection.Empty()
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"b"})
		e := corestr.New.Collection.Empty()
		c.AddCollections(a, e, b)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddCollections adds non-empty collections", actual)
	})
}

func Test_Cov58_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddPointerCollectionsLock", func() {
		c := corestr.New.Collection.Empty()
		a := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(a)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddPointerCollectionsLock adds items", actual)
	})
}

func Test_Cov58_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsValues", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		c.AddHashmapsValues(hm)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashmapsValues adds values", actual)
	})
}

func Test_Cov58_Collection_AddHashmapsValues_NilAndEmpty(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsValues_NilAndEmpty", func() {
		c := corestr.New.Collection.Empty()
		c.AddHashmapsValues(nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsValues nil returns same", actual)
	})
}

func Test_Cov58_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeys", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		c.AddHashmapsKeys(hm)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeys adds keys", actual)
	})
}

func Test_Cov58_Collection_AddHashmapsKeys_Nil(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeys_Nil", func() {
		c := corestr.New.Collection.Empty()
		c.AddHashmapsKeys(nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeys nil returns same", actual)
	})
}

func Test_Cov58_Collection_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeysValues", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		c.AddHashmapsKeysValues(hm)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValues adds key+value", actual)
	})
}

func Test_Cov58_Collection_AddHashmapsKeysValues_Nil(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeysValues_Nil", func() {
		c := corestr.New.Collection.Empty()
		c.AddHashmapsKeysValues(nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValues nil returns same", actual)
	})
}

func Test_Cov58_Collection_AddHashmapsKeysValuesUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeysValuesUsingFilter", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v", "k2": "v2"})
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Key + "=" + pair.Value, true, false
		}
		c.AddHashmapsKeysValuesUsingFilter(filter, hm)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValuesUsingFilter adds filtered", actual)
	})
}

func Test_Cov58_Collection_AddHashmapsKeysValuesUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeysValuesUsingFilter_Break", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2", "c": "3"})
		count := 0
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			count++
			return pair.Key, true, count >= 1
		}
		c.AddHashmapsKeysValuesUsingFilter(filter, hm)
		actual := args.Map{"stopped": c.Length() <= 1}
		expected := args.Map{"stopped": true}
		expected.ShouldBeEqual(t, 0, "Filter with break stops early", actual)
	})
}

func Test_Cov58_Collection_AddHashmapsKeysValuesUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeysValuesUsingFilter_Nil", func() {
		c := corestr.New.Collection.Empty()
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) { return "", false, false }
		c.AddHashmapsKeysValuesUsingFilter(filter, nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Filter nil hashmaps returns same", actual)
	})
}

func Test_Cov58_Collection_AddHashmapsKeysValuesUsingFilter_Skip(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeysValuesUsingFilter_Skip", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return "", false, false
		}
		c.AddHashmapsKeysValuesUsingFilter(filter, hm)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Filter skip returns no items", actual)
	})
}

func Test_Cov58_Collection_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddWithWgLock", func() {
		c := corestr.New.Collection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(wg, "val")
		wg.Wait()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddWithWgLock adds and signals wg", actual)
	})
}

func Test_Cov58_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IndexAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		actual := args.Map{"val": c.IndexAt(1)}
		expected := args.Map{"val": "b"}
		expected.ShouldBeEqual(t, 0, "IndexAt returns item at index", actual)
	})
}

func Test_Cov58_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_SafeIndexAtUsingLength", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{
			"valid":   c.SafeIndexAtUsingLength("def", 2, 1),
			"invalid": c.SafeIndexAtUsingLength("def", 2, 5),
		}
		expected := args.Map{"valid": "b", "invalid": "def"}
		expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLength returns default on oob", actual)
	})
}

func Test_Cov58_Collection_First(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_First", func() {
		c := corestr.New.Collection.Strings([]string{"first", "second"})
		actual := args.Map{"first": c.First()}
		expected := args.Map{"first": "first"}
		expected.ShouldBeEqual(t, 0, "First returns first item", actual)
	})
}

func Test_Cov58_Collection_Single(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Single", func() {
		c := corestr.New.Collection.Strings([]string{"only"})
		actual := args.Map{"val": c.Single()}
		expected := args.Map{"val": "only"}
		expected.ShouldBeEqual(t, 0, "Single returns only item", actual)
	})
}

func Test_Cov58_Collection_Single_Panics(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Single_Panics", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			c.Single()
		}()
		actual := args.Map{"panicked": panicked}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "Single panics on multiple items", actual)
	})
}

func Test_Cov58_Collection_Last(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Last", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "last"})
		actual := args.Map{"last": c.Last()}
		expected := args.Map{"last": "last"}
		expected.ShouldBeEqual(t, 0, "Last returns last item", actual)
	})
}

func Test_Cov58_Collection_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_LastOrDefault", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		e := corestr.New.Collection.Empty()
		actual := args.Map{"last": c.LastOrDefault(), "empty": e.LastOrDefault()}
		expected := args.Map{"last": "b", "empty": ""}
		expected.ShouldBeEqual(t, 0, "LastOrDefault returns default on empty", actual)
	})
}

func Test_Cov58_Collection_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_FirstOrDefault", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		e := corestr.New.Collection.Empty()
		actual := args.Map{"first": c.FirstOrDefault(), "empty": e.FirstOrDefault()}
		expected := args.Map{"first": "a", "empty": ""}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault returns default on empty", actual)
	})
}

func Test_Cov58_Collection_Take(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Take", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		r := c.Take(2)
		actual := args.Map{"len": r.Length(), "first": r.First()}
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "Take returns first N items", actual)
	})
}

func Test_Cov58_Collection_Take_MoreThanLength(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Take_MoreThanLength", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.Take(10)
		actual := args.Map{"len": r.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Take returns all when N > length", actual)
	})
}

func Test_Cov58_Collection_Take_Zero(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Take_Zero", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.Take(0)
		actual := args.Map{"empty": r.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Take 0 returns empty", actual)
	})
}

func Test_Cov58_Collection_Skip(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Skip", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r := c.Skip(1)
		actual := args.Map{"len": r.Length(), "first": r.First()}
		expected := args.Map{"len": 2, "first": "b"}
		expected.ShouldBeEqual(t, 0, "Skip returns items after N", actual)
	})
}

func Test_Cov58_Collection_Skip_Zero(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Skip_Zero", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.Skip(0)
		actual := args.Map{"len": r.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Skip 0 returns same", actual)
	})
}

func Test_Cov58_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Reverse", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()
		actual := args.Map{"first": c.First(), "last": c.Last()}
		expected := args.Map{"first": "c", "last": "a"}
		expected.ShouldBeEqual(t, 0, "Reverse reverses items", actual)
	})
}

func Test_Cov58_Collection_Reverse_TwoItems(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Reverse_TwoItems", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Reverse()
		actual := args.Map{"first": c.First(), "last": c.Last()}
		expected := args.Map{"first": "b", "last": "a"}
		expected.ShouldBeEqual(t, 0, "Reverse two items swaps", actual)
	})
}

func Test_Cov58_Collection_Reverse_Single(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Reverse_Single", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Reverse()
		actual := args.Map{"first": c.First()}
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Reverse single item unchanged", actual)
	})
}

func Test_Cov58_Collection_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_GetPagesSize", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		actual := args.Map{"pages": c.GetPagesSize(2), "zero": c.GetPagesSize(0), "neg": c.GetPagesSize(-1)}
		expected := args.Map{"pages": 3, "zero": 0, "neg": 0}
		expected.ShouldBeEqual(t, 0, "GetPagesSize calculates pages", actual)
	})
}

func Test_Cov58_Collection_GetPagedCollection(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_GetPagedCollection", func() {
		items := make([]string, 10)
		for i := range items {
			items[i] = "item"
		}
		c := corestr.New.Collection.Strings(items)
		paged := c.GetPagedCollection(3)
		actual := args.Map{"len": paged.Length()}
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "GetPagedCollection returns correct pages", actual)
	})
}

func Test_Cov58_Collection_GetPagedCollection_SmallSet(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_GetPagedCollection_SmallSet", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		paged := c.GetPagedCollection(5)
		actual := args.Map{"len": paged.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetPagedCollection small set returns 1 page", actual)
	})
}

func Test_Cov58_Collection_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_GetSinglePageCollection", func() {
		items := []string{"a", "b", "c", "d", "e", "f"}
		c := corestr.New.Collection.Strings(items)
		p := c.GetSinglePageCollection(2, 2)
		actual := args.Map{"len": p.Length(), "first": p.First()}
		expected := args.Map{"len": 2, "first": "c"}
		expected.ShouldBeEqual(t, 0, "GetSinglePageCollection returns correct page", actual)
	})
}

func Test_Cov58_Collection_GetSinglePageCollection_SmallSet(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_GetSinglePageCollection_SmallSet", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		p := c.GetSinglePageCollection(5, 1)
		actual := args.Map{"len": p.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetSinglePageCollection small set returns all", actual)
	})
}

func Test_Cov58_Collection_GetSinglePageCollection_LastPage(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_GetSinglePageCollection_LastPage", func() {
		items := []string{"a", "b", "c", "d", "e"}
		c := corestr.New.Collection.Strings(items)
		p := c.GetSinglePageCollection(2, 3)
		actual := args.Map{"len": p.Length(), "first": p.First()}
		expected := args.Map{"len": 1, "first": "e"}
		expected.ShouldBeEqual(t, 0, "GetSinglePageCollection last page has remainder", actual)
	})
}
