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

func Test_Collection_JsonString_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_JsonString", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.JsonString()

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JsonString returns non-empty", actual)
	})
}

func Test_Collection_JsonStringMust(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_JsonStringMust", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"x"})
		s := c.JsonStringMust()

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JsonStringMust returns non-empty", actual)
	})
}

func Test_Collection_HasAnyItem_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_HasAnyItem", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		e := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{
			"has": c.HasAnyItem(),
			"empty": e.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"empty": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAnyItem returns correct bool", actual)
	})
}

func Test_Collection_LastIndex_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_LastIndex", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"lastIndex": c.LastIndex()}

		// Assert
		expected := args.Map{"lastIndex": 2}
		expected.ShouldBeEqual(t, 0, "LastIndex returns len-1", actual)
	})
}

func Test_Collection_HasIndex_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_HasIndex", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"valid": c.HasIndex(1),
			"invalid": c.HasIndex(5),
			"neg": c.HasIndex(-1),
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"invalid": false,
			"neg": false,
		}
		expected.ShouldBeEqual(t, 0, "HasIndex checks bounds correctly", actual)
	})
}

func Test_Collection_ListStringsPtr(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_ListStringsPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(c.ListStringsPtr())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListStringsPtr returns items", actual)
	})
}

func Test_Collection_ListStrings_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_ListStrings", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(c.ListStrings())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListStrings returns items", actual)
	})
}

func Test_Collection_StringJSON(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_StringJSON", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.StringJSON()

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringJSON returns non-empty", actual)
	})
}

func Test_Collection_RemoveAt_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_RemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok1 := c.RemoveAt(1)
		ok2 := c.RemoveAt(-1)
		ok3 := c.RemoveAt(100)

		// Act
		actual := args.Map{
			"ok": ok1,
			"neg": ok2,
			"oob": ok3,
			"len": c.Length(),
		}

		// Assert
		expected := args.Map{
			"ok": true,
			"neg": false,
			"oob": false,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "RemoveAt removes item at index", actual)
	})
}

func Test_Collection_Count(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Count", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"count": c.Count()}

		// Assert
		expected := args.Map{"count": 2}
		expected.ShouldBeEqual(t, 0, "Count returns length", actual)
	})
}

func Test_Collection_Capacity_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Capacity", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		e := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{
			"hasCapNonEmpty": c.Capacity() > 0,
			"emptyCapGte0": e.Capacity() >= 0,
		}

		// Assert
		expected := args.Map{
			"hasCapNonEmpty": true,
			"emptyCapGte0": true,
		}
		expected.ShouldBeEqual(t, 0, "Capacity returns cap", actual)
	})
}

func Test_Collection_Length_NilItems(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Length_NilItems", func() {
		// Arrange
		var c *corestr.Collection

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length on nil returns 0", actual)
	})
}

func Test_Collection_LengthLock_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_LengthLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"len": c.LengthLock()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LengthLock returns length", actual)
	})
}

func Test_Collection_IsEquals_SameContent(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_SameContent", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals returns true for same content", actual)
	})
}

func Test_Collection_IsEquals_DiffContent(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_DiffContent", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "c"})

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals returns false for different content", actual)
	})
}

func Test_Collection_IsEquals_NilBoth(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_NilBoth", func() {
		// Arrange
		var a, b *corestr.Collection

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals nil==nil is true", actual)
	})
}

func Test_Collection_IsEquals_NilOne(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_NilOne", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a"})
		var b *corestr.Collection

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals nil vs non-nil is false", actual)
	})
}

func Test_Collection_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_DiffLength", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff length is false", actual)
	})
}

func Test_Collection_IsEquals_BothEmpty_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.New.Collection.Empty()
		b := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty is true", actual)
	})
}

func Test_Collection_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEquals_SamePtr", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"x"})

		// Act
		actual := args.Map{"eq": a.IsEquals(a)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals same pointer is true", actual)
	})
}

func Test_Collection_IsEqualsWithSensitive_CaseInsensitive_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEqualsWithSensitive_CaseInsensitive", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"Hello", "World"})
		b := corestr.New.Collection.Strings([]string{"hello", "world"})

		// Act
		actual := args.Map{
			"caseSensitive": a.IsEqualsWithSensitive(true, b),
			"caseInsensitive": a.IsEqualsWithSensitive(false, b),
		}

		// Assert
		expected := args.Map{
			"caseSensitive": false,
			"caseInsensitive": true,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualsWithSensitive handles case", actual)
	})
}

func Test_Collection_IsEqualsWithSensitive_CaseInsensitiveMismatch(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEqualsWithSensitive_CaseInsensitiveMismatch", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"hello"})
		b := corestr.New.Collection.Strings([]string{"xyz"})

		// Act
		actual := args.Map{"eq": a.IsEqualsWithSensitive(false, b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualsWithSensitive case insensitive mismatch", actual)
	})
}

func Test_Collection_IsEmptyLock_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IsEmptyLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		d := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"empty": c.IsEmptyLock(),
			"notEmpty": d.IsEmptyLock(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"notEmpty": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock returns correct bool", actual)
	})
}

func Test_Collection_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_HasItems", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		e := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{
			"has": c.HasItems(),
			"empty": e.HasItems(),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"empty": false,
		}
		expected.ShouldBeEqual(t, 0, "HasItems returns correct bool", actual)
	})
}

func Test_Collection_AddLock_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddLock("x")

		// Act
		actual := args.Map{
			"len": c.Length(),
			"first": c.First(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"first": "x",
		}
		expected.ShouldBeEqual(t, 0, "AddLock adds item", actual)
	})
}

func Test_Collection_AddNonEmpty_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmpty("a")
		c.AddNonEmpty("")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty skips empty strings", actual)
	})
}

func Test_Collection_AddNonEmptyWhitespace_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddNonEmptyWhitespace", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyWhitespace("a")
		c.AddNonEmptyWhitespace("   ")
		c.AddNonEmptyWhitespace("")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyWhitespace skips whitespace", actual)
	})
}

func Test_Collection_AddError_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddError", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddError(nil)
		c.AddError(errForTest)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddError skips nil error", actual)
	})
}

func Test_Collection_AsDefaultError_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AsDefaultError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"err1", "err2"})
		e := c.AsDefaultError()

		// Act
		actual := args.Map{"nonNil": e != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsDefaultError returns error", actual)
	})
}

func Test_Collection_AsError_Empty_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AsError_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		e := c.AsError(",")

		// Act
		actual := args.Map{"nil": e == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "AsError on empty returns nil", actual)
	})
}

func Test_Collection_AddIf_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddIf", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddIf(true, "yes")
		c.AddIf(false, "no")

		// Act
		actual := args.Map{
			"len": c.Length(),
			"first": c.First(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"first": "yes",
		}
		expected.ShouldBeEqual(t, 0, "AddIf conditionally adds", actual)
	})
}

func Test_Collection_EachItemSplitBy_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_EachItemSplitBy", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
		result := c.EachItemSplitBy(",")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "EachItemSplitBy splits items", actual)
	})
}

func Test_Collection_ConcatNew_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_ConcatNew", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		n := c.ConcatNew(0, "b", "c")

		// Act
		actual := args.Map{"len": n.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ConcatNew creates new collection", actual)
	})
}

func Test_Collection_ConcatNew_NoAdding(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_ConcatNew_NoAdding", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		n := c.ConcatNew(0)

		// Act
		actual := args.Map{"len": n.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ConcatNew with no additions clones", actual)
	})
}

func Test_Collection_ToError_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_ToError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"e1", "e2"})
		e := c.ToError(",")

		// Act
		actual := args.Map{"nonNil": e != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ToError returns error", actual)
	})
}

func Test_Collection_ToDefaultError_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_ToDefaultError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"e1"})
		e := c.ToDefaultError()

		// Act
		actual := args.Map{"nonNil": e != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ToDefaultError returns error", actual)
	})
}

func Test_Collection_AddIfMany_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddIfMany", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddIfMany(true, "a", "b")
		c.AddIfMany(false, "c")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddIfMany adds conditionally", actual)
	})
}

func Test_Collection_AddFunc_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddFunc", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddFunc(func() string { return "val" })

		// Act
		actual := args.Map{
			"len": c.Length(),
			"first": c.First(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"first": "val",
		}
		expected.ShouldBeEqual(t, 0, "AddFunc adds func result", actual)
	})
}

func Test_Collection_AddFuncErr_Success(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddFuncErr_Success", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddFuncErr(
			func() (string, error) { return "ok", nil },

		// Assert
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)

		// Act
		actual = args.Map{
			"len": c.Length(),
			"first": c.First(),
		}
		expected = args.Map{
			"len": 1,
			"first": "ok",
		}
		expected.ShouldBeEqual(t, 0, "AddFuncErr adds on success", actual)
	})
}

func Test_Collection_AddFuncErr_Error(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddFuncErr_Error", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		handlerCalled := false
		c.AddFuncErr(
			func() (string, error) { return "", errForTest },
			func(err error) { handlerCalled = true },
		)

		// Act
		actual := args.Map{
			"len": c.Length(),
			"handled": handlerCalled,
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"handled": true,
		}
		expected.ShouldBeEqual(t, 0, "AddFuncErr calls handler on error", actual)
	})
}

func Test_Collection_AddsLock_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddsLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddsLock("a", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsLock adds items", actual)
	})
}

func Test_Collection_AddCollection_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		other := corestr.New.Collection.Strings([]string{"b", "c"})
		c.AddCollection(other)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AddCollection appends items", actual)
	})
}

func Test_Collection_AddCollection_Empty_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddCollection_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		empty := corestr.New.Collection.Empty()
		c.AddCollection(empty)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollection skips empty", actual)
	})
}

func Test_Collection_AddCollections_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddCollections", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"b"})
		e := corestr.New.Collection.Empty()
		c.AddCollections(a, e, b)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddCollections adds non-empty collections", actual)
	})
}

func Test_Collection_AddPointerCollectionsLock_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddPointerCollectionsLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		a := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(a)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddPointerCollectionsLock adds items", actual)
	})
}

func Test_Collection_AddHashmapsValues_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsValues", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		c.AddHashmapsValues(hm)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashmapsValues adds values", actual)
	})
}

func Test_Collection_AddHashmapsValues_NilAndEmpty(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsValues_NilAndEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddHashmapsValues(nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsValues nil returns same", actual)
	})
}

func Test_Collection_AddHashmapsKeys_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeys", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		c.AddHashmapsKeys(hm)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeys adds keys", actual)
	})
}

func Test_Collection_AddHashmapsKeys_Nil_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeys_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddHashmapsKeys(nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeys nil returns same", actual)
	})
}

func Test_Collection_AddHashmapsKeysValues_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeysValues", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		c.AddHashmapsKeysValues(hm)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValues adds key+value", actual)
	})
}

func Test_Collection_AddHashmapsKeysValues_Nil_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeysValues_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddHashmapsKeysValues(nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValues nil returns same", actual)
	})
}

func Test_Collection_AddHashmapsKeysValuesUsingFilter_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeysValuesUsingFilter", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v", "k2": "v2"})
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Key + "=" + pair.Value, true, false
		}
		c.AddHashmapsKeysValuesUsingFilter(filter, hm)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValuesUsingFilter adds filtered", actual)
	})
}

func Test_Collection_AddHashmapsKeysValuesUsingFilter_Break_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeysValuesUsingFilter_Break", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2", "c": "3"})
		count := 0
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			count++
			return pair.Key, true, count >= 1
		}
		c.AddHashmapsKeysValuesUsingFilter(filter, hm)

		// Act
		actual := args.Map{"stopped": c.Length() <= 1}

		// Assert
		expected := args.Map{"stopped": true}
		expected.ShouldBeEqual(t, 0, "Filter with break stops early", actual)
	})
}

func Test_Collection_AddHashmapsKeysValuesUsingFilter_Nil_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeysValuesUsingFilter_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) { return "", false, false }
		c.AddHashmapsKeysValuesUsingFilter(filter, nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Filter nil hashmaps returns same", actual)
	})
}

func Test_Collection_AddHashmapsKeysValuesUsingFilter_Skip(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddHashmapsKeysValuesUsingFilter_Skip", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return "", false, false
		}
		c.AddHashmapsKeysValuesUsingFilter(filter, hm)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Filter skip returns no items", actual)
	})
}

func Test_Collection_AddWithWgLock_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_AddWithWgLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(wg, "val")
		wg.Wait()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddWithWgLock adds and signals wg", actual)
	})
}

func Test_Collection_IndexAt_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_IndexAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"val": c.IndexAt(1)}

		// Assert
		expected := args.Map{"val": "b"}
		expected.ShouldBeEqual(t, 0, "IndexAt returns item at index", actual)
	})
}

func Test_Collection_SafeIndexAtUsingLength_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_SafeIndexAtUsingLength", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"valid":   c.SafeIndexAtUsingLength("def", 2, 1),
			"invalid": c.SafeIndexAtUsingLength("def", 2, 5),
		}

		// Assert
		expected := args.Map{
			"valid": "b",
			"invalid": "def",
		}
		expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLength returns default on oob", actual)
	})
}

func Test_Collection_First_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_First", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"first", "second"})

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "first"}
		expected.ShouldBeEqual(t, 0, "First returns first item", actual)
	})
}

func Test_Collection_Single_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Single", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"only"})

		// Act
		actual := args.Map{"val": c.Single()}

		// Assert
		expected := args.Map{"val": "only"}
		expected.ShouldBeEqual(t, 0, "Single returns only item", actual)
	})
}

func Test_Collection_Single_Panics(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Single_Panics", func() {
		// Arrange
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

		// Act
		actual := args.Map{"panicked": panicked}

		// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "Single panics on multiple items", actual)
	})
}

func Test_Collection_Last_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Last", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "last"})

		// Act
		actual := args.Map{"last": c.Last()}

		// Assert
		expected := args.Map{"last": "last"}
		expected.ShouldBeEqual(t, 0, "Last returns last item", actual)
	})
}

func Test_Collection_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_LastOrDefault", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		e := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{
			"last": c.LastOrDefault(),
			"empty": e.LastOrDefault(),
		}

		// Assert
		expected := args.Map{
			"last": "b",
			"empty": "",
		}
		expected.ShouldBeEqual(t, 0, "LastOrDefault returns default on empty", actual)
	})
}

func Test_Collection_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_FirstOrDefault", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		e := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{
			"first": c.FirstOrDefault(),
			"empty": e.FirstOrDefault(),
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"empty": "",
		}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault returns default on empty", actual)
	})
}

func Test_Collection_Take_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Take", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		r := c.Take(2)

		// Act
		actual := args.Map{
			"len": r.Length(),
			"first": r.First(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "Take returns first N items", actual)
	})
}

func Test_Collection_Take_MoreThanLength_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Take_MoreThanLength", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.Take(10)

		// Act
		actual := args.Map{"len": r.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Take returns all when N > length", actual)
	})
}

func Test_Collection_Take_Zero_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Take_Zero", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.Take(0)

		// Act
		actual := args.Map{"empty": r.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Take 0 returns empty", actual)
	})
}

func Test_Collection_Skip_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Skip", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r := c.Skip(1)

		// Act
		actual := args.Map{
			"len": r.Length(),
			"first": r.First(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "b",
		}
		expected.ShouldBeEqual(t, 0, "Skip returns items after N", actual)
	})
}

func Test_Collection_Skip_Zero_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Skip_Zero", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.Skip(0)

		// Act
		actual := args.Map{"len": r.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Skip 0 returns same", actual)
	})
}

func Test_Collection_Reverse_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Reverse", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()

		// Act
		actual := args.Map{
			"first": c.First(),
			"last": c.Last(),
		}

		// Assert
		expected := args.Map{
			"first": "c",
			"last": "a",
		}
		expected.ShouldBeEqual(t, 0, "Reverse reverses items", actual)
	})
}

func Test_Collection_Reverse_TwoItems(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Reverse_TwoItems", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Reverse()

		// Act
		actual := args.Map{
			"first": c.First(),
			"last": c.Last(),
		}

		// Assert
		expected := args.Map{
			"first": "b",
			"last": "a",
		}
		expected.ShouldBeEqual(t, 0, "Reverse two items swaps", actual)
	})
}

func Test_Collection_Reverse_Single_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_Reverse_Single", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Reverse()

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Reverse single item unchanged", actual)
	})
}

func Test_Collection_GetPagesSize_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_GetPagesSize", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})

		// Act
		actual := args.Map{
			"pages": c.GetPagesSize(2),
			"zero": c.GetPagesSize(0),
			"neg": c.GetPagesSize(-1),
		}

		// Assert
		expected := args.Map{
			"pages": 3,
			"zero": 0,
			"neg": 0,
		}
		expected.ShouldBeEqual(t, 0, "GetPagesSize calculates pages", actual)
	})
}

func Test_Collection_GetPagedCollection_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_GetPagedCollection", func() {
		// Arrange
		items := make([]string, 10)
		for i := range items {
			items[i] = "item"
		}
		c := corestr.New.Collection.Strings(items)
		paged := c.GetPagedCollection(3)

		// Act
		actual := args.Map{"len": paged.Length()}

		// Assert
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "GetPagedCollection returns correct pages", actual)
	})
}

func Test_Collection_GetPagedCollection_SmallSet(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_GetPagedCollection_SmallSet", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		paged := c.GetPagedCollection(5)

		// Act
		actual := args.Map{"len": paged.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetPagedCollection small set returns 1 page", actual)
	})
}

func Test_Collection_GetSinglePageCollection_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_GetSinglePageCollection", func() {
		// Arrange
		items := []string{"a", "b", "c", "d", "e", "f"}
		c := corestr.New.Collection.Strings(items)
		p := c.GetSinglePageCollection(2, 2)

		// Act
		actual := args.Map{
			"len": p.Length(),
			"first": p.First(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "c",
		}
		expected.ShouldBeEqual(t, 0, "GetSinglePageCollection returns correct page", actual)
	})
}

func Test_Collection_GetSinglePageCollection_SmallSet(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_GetSinglePageCollection_SmallSet", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		p := c.GetSinglePageCollection(5, 1)

		// Act
		actual := args.Map{"len": p.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetSinglePageCollection small set returns all", actual)
	})
}

func Test_Collection_GetSinglePageCollection_LastPage_FromCollectionJsonString(t *testing.T) {
	safeTest(t, "Test_Cov58_Collection_GetSinglePageCollection_LastPage", func() {
		// Arrange
		items := []string{"a", "b", "c", "d", "e"}
		c := corestr.New.Collection.Strings(items)
		p := c.GetSinglePageCollection(2, 3)

		// Act
		actual := args.Map{
			"len": p.Length(),
			"first": p.First(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"first": "e",
		}
		expected.ShouldBeEqual(t, 0, "GetSinglePageCollection last page has remainder", actual)
	})
}
