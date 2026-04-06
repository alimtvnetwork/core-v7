package corestrtests

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===== Collection — core methods =====

func Test_C35_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JsonString", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		// Act
		result := c.JsonString()
		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty JSON string", actual)
	})
}

func Test_C35_Collection_JsonStringMust(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JsonStringMust", func() {
		c := corestr.New.Collection.Strings([]string{"x"})
		result := c.JsonStringMust()
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C35_Collection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasAnyItem", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected HasAnyItem true", actual)
		e := corestr.New.Collection.Empty()
		actual := args.Map{"result": e.HasAnyItem()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected HasAnyItem false on empty", actual)
	})
}

func Test_C35_Collection_LastIndex(t *testing.T) {
	safeTest(t, "Test_C35_Collection_LastIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		actual := args.Map{"result": c.LastIndex() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_HasIndex(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.HasIndex(0)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected HasIndex(0) true", actual)
		actual := args.Map{"result": c.HasIndex(1)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected HasIndex(1) true", actual)
		actual := args.Map{"result": c.HasIndex(2)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected HasIndex(2) false", actual)
		actual := args.Map{"result": c.HasIndex(-1)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected HasIndex(-1) false", actual)
	})
}

func Test_C35_Collection_ListStringsPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ListStringsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": len(c.ListStringsPtr()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_ListStrings(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ListStrings", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": len(c.ListStrings()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_StringJSON(t *testing.T) {
	safeTest(t, "Test_C35_Collection_StringJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.StringJSON() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C35_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_C35_Collection_RemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := c.RemoveAt(1)
		actual := args.Map{"result": ok || c.Length() != 2}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected remove success", actual)
		actual := args.Map{"result": c.RemoveAt(-1)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for negative index", actual)
		actual := args.Map{"result": c.RemoveAt(99)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for out of range", actual)
	})
}

func Test_C35_Collection_Count(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Count", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.Count() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_Capacity(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Capacity", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"result": c.Capacity() < 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected capacity >= 10", actual)
		var nilC *corestr.Collection
		actual := args.Map{"result": nilC.Capacity() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for nil items", actual)
	})
}

func Test_C35_Collection_Length_NilReceiver(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Length_NilReceiver", func() {
		var c *corestr.Collection
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C35_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_LengthLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.LengthLock() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEquals", func() {
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		c := corestr.New.Collection.Strings([]string{"a", "c"})
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual := args.Map{"result": a.IsEquals(c)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_C35_Collection_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive", func() {
		a := corestr.New.Collection.Strings([]string{"Hello"})
		b := corestr.New.Collection.Strings([]string{"hello"})
		actual := args.Map{"result": a.IsEqualsWithSensitive(true, b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected case sensitive not equal", actual)
		actual := args.Map{"result": a.IsEqualsWithSensitive(false, b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected case insensitive equal", actual)
	})
}

func Test_C35_Collection_IsEqualsWithSensitive_BothNil(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive_BothNil", func() {
		var a, b *corestr.Collection
		actual := args.Map{"result": a.IsEqualsWithSensitive(true, b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	})
}

func Test_C35_Collection_IsEqualsWithSensitive_OneNil(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive_OneNil", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		var b *corestr.Collection
		actual := args.Map{"result": a.IsEqualsWithSensitive(true, b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "one nil should not equal", actual)
	})
}

func Test_C35_Collection_IsEqualsWithSensitive_SamePtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive_SamePtr", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": a.IsEqualsWithSensitive(true, a)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same ptr should be equal", actual)
	})
}

func Test_C35_Collection_IsEqualsWithSensitive_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive_BothEmpty", func() {
		a := corestr.New.Collection.Empty()
		b := corestr.New.Collection.Empty()
		actual := args.Map{"result": a.IsEqualsWithSensitive(true, b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "both empty should be equal", actual)
	})
}

func Test_C35_Collection_IsEqualsWithSensitive_OneEmpty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive_OneEmpty", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Empty()
		actual := args.Map{"result": a.IsEqualsWithSensitive(true, b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "one empty should not equal", actual)
	})
}

func Test_C35_Collection_IsEqualsWithSensitive_DiffLength(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive_DiffLength", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": a.IsEqualsWithSensitive(true, b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "diff length should not equal", actual)
	})
}

func Test_C35_Collection_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEmptyLock", func() {
		c := corestr.New.Collection.Empty()
		actual := args.Map{"result": c.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C35_Collection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEmpty", func() {
		var c *corestr.Collection
		actual := args.Map{"result": c.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
	})
}

func Test_C35_Collection_HasItems(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasItems", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected HasItems true", actual)
		e := corestr.New.Collection.Empty()
		actual := args.Map{"result": e.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected HasItems false", actual)
	})
}

func Test_C35_Collection_AddLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddLock", func() {
		c := corestr.New.Collection.Empty()
		c.AddLock("a")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddNonEmpty", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmpty("")
		c.AddNonEmpty("a")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddNonEmptyWhitespace", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyWhitespace("  ")
		c.AddNonEmptyWhitespace("a")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_Add(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Add", func() {
		c := corestr.New.Collection.Empty()
		c.Add("a")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddError", func() {
		c := corestr.New.Collection.Empty()
		c.AddError(nil)
		c.AddError(errors.New("err1"))
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": c.First() != "err1"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected err1", actual)
	})
}

func Test_C35_Collection_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AsDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		err := c.AsDefaultError()
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C35_Collection_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AsError_Empty", func() {
		c := corestr.New.Collection.Empty()
		actual := args.Map{"result": c.AsError(",") != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C35_Collection_AsError_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AsError_NonEmpty", func() {
		c := corestr.New.Collection.Strings([]string{"err"})
		actual := args.Map{"result": c.AsError(",") == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C35_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddIf", func() {
		c := corestr.New.Collection.Empty()
		c.AddIf(false, "skip")
		c.AddIf(true, "keep")
		actual := args.Map{"result": c.Length() != 1 || c.First() != "keep"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected only keep", actual)
	})
}

func Test_C35_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_C35_Collection_EachItemSplitBy", func() {
		c := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
		result := c.EachItemSplitBy(",")
		actual := args.Map{"result": len(result) != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_C35_Collection_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ConcatNew_Empty", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		newC := c.ConcatNew(0)
		actual := args.Map{"result": newC.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cloned with 1 item", actual)
	})
}

func Test_C35_Collection_ConcatNew_WithItems(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ConcatNew_WithItems", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		newC := c.ConcatNew(0, "b", "c")
		actual := args.Map{"result": newC.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C35_Collection_ToError(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ToError", func() {
		c := corestr.New.Collection.Strings([]string{"err1"})
		actual := args.Map{"result": c.ToError(",") == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C35_Collection_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ToDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"err1"})
		actual := args.Map{"result": c.ToDefaultError() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C35_Collection_AddIfMany(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddIfMany", func() {
		c := corestr.New.Collection.Empty()
		c.AddIfMany(false, "a", "b")
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		c.AddIfMany(true, "a", "b")
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddFunc", func() {
		c := corestr.New.Collection.Empty()
		c.AddFunc(func() string { return "hello" })
		actual := args.Map{"result": c.Length() != 1 || c.First() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_C35_Collection_AddFuncErr_Success(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddFuncErr_Success", func() {
		c := corestr.New.Collection.Empty()
		c.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(e error) {},
		)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_AddFuncErr_Error(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddFuncErr_Error", func() {
		c := corestr.New.Collection.Empty()
		var caught error
		c.AddFuncErr(
			func() (string, error) { return "", errors.New("fail") },
			func(e error) { caught = e },
		)
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual := args.Map{"result": caught == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error caught", actual)
	})
}

func Test_C35_Collection_AddsLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddsLock", func() {
		c := corestr.New.Collection.Empty()
		c.AddsLock("a", "b")
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Adds", func() {
		c := corestr.New.Collection.Empty()
		c.Adds("a", "b", "c")
		actual := args.Map{"result": c.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C35_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddStrings", func() {
		c := corestr.New.Collection.Empty()
		c.AddStrings([]string{"x", "y"})
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddCollection", func() {
		c := corestr.New.Collection.Empty()
		other := corestr.New.Collection.Strings([]string{"a"})
		c.AddCollection(other)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AddCollection(corestr.New.Collection.Empty())
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should skip empty", actual)
	})
}

func Test_C35_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddCollections", func() {
		c := corestr.New.Collection.Empty()
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"b"})
		c.AddCollections(a, b)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddPointerCollectionsLock", func() {
		c := corestr.New.Collection.Empty()
		a := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(a)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddHashmapsValues", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		c.AddHashmapsValues(h)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AddHashmapsValues(nil)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should skip nil hashmaps", actual)
	})
}

func Test_C35_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddHashmapsKeys", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k1", "v1")
		c.AddHashmapsKeys(h)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddHashmapsKeysValues", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValues(h)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (key + value)", actual)
	})
}

func Test_C35_Collection_AddHashmapsKeysValuesUsingFilter(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddHashmapsKeysValuesUsingFilter", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValuesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Key + "=" + pair.Value, true, false
			},
			h,
		)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_AddHashmapsKeysValuesUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddHashmapsKeysValuesUsingFilter_Break", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k1", "v1")
		h.AddOrUpdate("k2", "v2")
		c.AddHashmapsKeysValuesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Key, true, true
			},
			h,
		)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 due to break", actual)
	})
}

func Test_C35_Collection_AddHashmapsKeysValuesUsingFilter_NilHashmaps(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddHashmapsKeysValuesUsingFilter_NilHashmaps", func() {
		c := corestr.New.Collection.Empty()
		c.AddHashmapsKeysValuesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return "", false, false
			},
		)
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C35_Collection_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddWithWgLock", func() {
		c := corestr.New.Collection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(wg, "a")
		wg.Wait()
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IndexAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.IndexAt(1) != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C35_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SafeIndexAtUsingLength", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.SafeIndexAtUsingLength("default", 2, 1) != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		actual := args.Map{"result": c.SafeIndexAtUsingLength("default", 1, 5) != "default"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default", actual)
	})
}

func Test_C35_Collection_First(t *testing.T) {
	safeTest(t, "Test_C35_Collection_First", func() {
		c := corestr.New.Collection.Strings([]string{"first", "second"})
		actual := args.Map{"result": c.First() != "first"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected first", actual)
	})
}

func Test_C35_Collection_Single(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Single", func() {
		c := corestr.New.Collection.Strings([]string{"only"})
		actual := args.Map{"result": c.Single() != "only"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected only", actual)
	})
}

func Test_C35_Collection_Last(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Last", func() {
		c := corestr.New.Collection.Strings([]string{"a", "last"})
		actual := args.Map{"result": c.Last() != "last"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected last", actual)
	})
}

func Test_C35_Collection_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_C35_Collection_LastOrDefault", func() {
		c := corestr.New.Collection.Strings([]string{"a", "last"})
		actual := args.Map{"result": c.LastOrDefault() != "last"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected last", actual)
		e := corestr.New.Collection.Empty()
		actual := args.Map{"result": e.LastOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty string", actual)
	})
}

func Test_C35_Collection_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FirstOrDefault", func() {
		c := corestr.New.Collection.Strings([]string{"first"})
		actual := args.Map{"result": c.FirstOrDefault() != "first"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected first", actual)
		e := corestr.New.Collection.Empty()
		actual := args.Map{"result": e.FirstOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C35_Collection_Take(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Take", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		taken := c.Take(2)
		actual := args.Map{"result": taken.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		all := c.Take(10)
		actual := args.Map{"result": all.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "take more than length returns all", actual)
		zero := c.Take(0)
		actual := args.Map{"result": zero.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "take 0 returns empty", actual)
	})
}

func Test_C35_Collection_Skip(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Skip", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		skipped := c.Skip(1)
		actual := args.Map{"result": skipped.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		none := c.Skip(0)
		actual := args.Map{"result": none != c}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "skip 0 returns self", actual)
	})
}

func Test_C35_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Reverse", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()
		actual := args.Map{"result": c.First() != "c" || c.Last() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected reversed", actual)
		// len 2
		c2 := corestr.New.Collection.Strings([]string{"x", "y"})
		c2.Reverse()
		actual := args.Map{"result": c2.First() != "y"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected y", actual)
		// len 1
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c1.Reverse()
		actual := args.Map{"result": c1.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C35_Collection_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_C35_Collection_GetPagesSize", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		actual := args.Map{"result": c.GetPagesSize(2) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)
		actual := args.Map{"result": c.GetPagesSize(0) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for zero page size", actual)
		actual := args.Map{"result": c.GetPagesSize(-1) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for negative page size", actual)
	})
}

func Test_C35_Collection_GetPagedCollection(t *testing.T) {
	safeTest(t, "Test_C35_Collection_GetPagedCollection", func() {
		items := make([]string, 10)
		for i := range items {
			items[i] = fmt.Sprintf("item%d", i)
		}
		c := corestr.New.Collection.Strings(items)
		paged := c.GetPagedCollection(3)
		actual := args.Map{"result": paged.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4 pages", actual)
	})
}

func Test_C35_Collection_GetPagedCollection_SmallCollection(t *testing.T) {
	safeTest(t, "Test_C35_Collection_GetPagedCollection_SmallCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		paged := c.GetPagedCollection(10)
		actual := args.Map{"result": paged.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 page", actual)
	})
}

func Test_C35_Collection_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_C35_Collection_GetSinglePageCollection", func() {
		items := make([]string, 20)
		for i := range items {
			items[i] = fmt.Sprintf("i%d", i)
		}
		c := corestr.New.Collection.Strings(items)
		page := c.GetSinglePageCollection(5, 2)
		actual := args.Map{"result": page.Length() != 5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
		// Small collection returns self
		small := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": small.GetSinglePageCollection(10, 1) != small}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected self", actual)
	})
}

func Test_C35_Collection_InsertAt(t *testing.T) {
	safeTest(t, "Test_C35_Collection_InsertAt", func() {
		c := corestr.New.Collection.Empty()
		c.InsertAt(0, "first")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ChainRemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(1)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_RemoveItemsIndexes(t *testing.T) {
	safeTest(t, "Test_C35_Collection_RemoveItemsIndexes", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.RemoveItemsIndexes(true, 1)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// nil indexes
		c.RemoveItemsIndexes(true)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (nil indexes)", actual)
	})
}

func Test_C35_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendCollectionPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		other := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollectionPtr(other)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_AppendCollections(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendCollections", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollections(b)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		c.AppendCollections()
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (empty)", actual)
	})
}

func Test_C35_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnys(42, "hello", nil)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (nil skipped)", actual)
	})
}

func Test_C35_Collection_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendAnysLock", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysLock(1, 2)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		c.AppendAnysLock()
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (empty)", actual)
	})
}

func Test_C35_Collection_AppendAnysUsingFilter(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendAnysUsingFilter", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilter(
			func(s string, i int) (string, bool, bool) {
				return s, true, false
			},
			"a", nil, "b",
		)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_AppendAnysUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendAnysUsingFilter_Break", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilter(
			func(s string, i int) (string, bool, bool) {
				return s, true, true
			},
			"a", "b",
		)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 due to break", actual)
	})
}

func Test_C35_Collection_AppendAnysUsingFilter_Skip(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendAnysUsingFilter_Skip", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilter(
			func(s string, i int) (string, bool, bool) {
				return s, false, false
			},
			"a",
		)
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C35_Collection_AppendAnysUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendAnysUsingFilterLock", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilterLock(
			func(s string, i int) (string, bool, bool) {
				return s, true, false
			},
			"a",
		)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c2 := corestr.New.Collection.Empty()
		c2.AppendAnysUsingFilterLock(nil)
		actual := args.Map{"result": c2.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C35_Collection_AppendAnysUsingFilterLock_BreakAndSkip(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendAnysUsingFilterLock_BreakAndSkip", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilterLock(
			func(s string, i int) (string, bool, bool) {
				return s, false, true
			},
			"a",
		)
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 (skip + break)", actual)
	})
}

func Test_C35_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendNonEmptyAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendNonEmptyAnys("a", nil, "")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddsNonEmpty", func() {
		c := corestr.New.Collection.Empty()
		c.AddsNonEmpty("a", "", "b")
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_AddsNonEmptyPtrLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddsNonEmptyPtrLock", func() {
		c := corestr.New.Collection.Empty()
		a := "a"
		empty := ""
		c.AddsNonEmptyPtrLock(&a, nil, &empty)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_C35_Collection_UniqueBoolMap", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		m := c.UniqueBoolMap()
		actual := args.Map{"result": len(m) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_UniqueBoolMapLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_UniqueBoolMapLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		m := c.UniqueBoolMapLock()
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_C35_Collection_UniqueList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		l := c.UniqueList()
		actual := args.Map{"result": len(l) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_UniqueListLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		l := c.UniqueListLock()
		actual := args.Map{"result": len(l) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Filter", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, s != "b", false
		})
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_Filter_Break(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Filter_Break", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, true, i == 0
		})
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 (break at 0)", actual)
	})
}

func Test_C35_Collection_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Filter_Empty", func() {
		c := corestr.New.Collection.Empty()
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C35_Collection_FilterLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FilterLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FilteredCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		fc := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"result": fc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_FilteredCollectionLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FilteredCollectionLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		fc := c.FilteredCollectionLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"result": fc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_FilterPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FilterPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.FilterPtr(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		actual := args.Map{"result": len(*result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_FilterPtr_Empty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FilterPtr_Empty", func() {
		c := corestr.New.Collection.Empty()
		result := c.FilterPtr(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		actual := args.Map{"result": len(*result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C35_Collection_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FilterPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.FilterPtrLock(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		actual := args.Map{"result": len(*result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyList", func() {
		c := corestr.New.Collection.Strings([]string{"", "a", "", "b"})
		list := c.NonEmptyList()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_NonEmptyList_Empty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyList_Empty", func() {
		c := corestr.New.Collection.Empty()
		actual := args.Map{"result": len(c.NonEmptyList()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C35_Collection_NonEmptyListPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyListPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": len(*c.NonEmptyListPtr()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HashsetAsIs", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		hs := c.HashsetAsIs()
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_HashsetWithDoubleLength(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HashsetWithDoubleLength", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetWithDoubleLength()
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_HashsetLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HashsetLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetLock()
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyItems", func() {
		c := corestr.New.Collection.Strings([]string{"", "a"})
		actual := args.Map{"result": len(c.NonEmptyItems()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_NonEmptyItemsPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyItemsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"", "a"})
		actual := args.Map{"result": len(c.NonEmptyItemsPtr()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyItemsOrNonWhitespace", func() {
		c := corestr.New.Collection.Strings([]string{"  ", "a"})
		actual := args.Map{"result": len(c.NonEmptyItemsOrNonWhitespace()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_NonEmptyItemsOrNonWhitespacePtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyItemsOrNonWhitespacePtr", func() {
		c := corestr.New.Collection.Strings([]string{"  ", "a"})
		actual := args.Map{"result": len(c.NonEmptyItemsOrNonWhitespacePtr()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_Items(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Items", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": len(c.Items()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_ListPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ListPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": len(c.ListPtr()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ListCopyPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		l := c.ListCopyPtrLock()
		actual := args.Map{"result": len(l) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		e := corestr.New.Collection.Empty()
		actual := args.Map{"result": len(e.ListCopyPtrLock()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C35_Collection_Has(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Has", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has a", actual)
		actual := args.Map{"result": c.Has("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not has z", actual)
	})
}

func Test_C35_Collection_HasLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.HasLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has a", actual)
	})
}

func Test_C35_Collection_HasPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		a := "a"
		actual := args.Map{"result": c.HasPtr(&a)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has a", actual)
		actual := args.Map{"result": c.HasPtr(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_C35_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		actual := args.Map{"result": c.HasAll("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has all", actual)
		actual := args.Map{"result": c.HasAll("a", "z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not has all", actual)
	})
}

func Test_C35_Collection_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SortedListAsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		list := c.SortedListAsc()
		actual := args.Map{"result": list[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
	})
}

func Test_C35_Collection_SortedAsc(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SortedAsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a"})
		c.SortedAsc()
		actual := args.Map{"result": c.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C35_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SortedAscLock", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a"})
		c.SortedAscLock()
		actual := args.Map{"result": c.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C35_Collection_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SortedListDsc", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		list := c.SortedListDsc()
		actual := args.Map{"result": list[0] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c first", actual)
	})
}

func Test_C35_Collection_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasUsingSensitivity", func() {
		c := corestr.New.Collection.Strings([]string{"Hello"})
		actual := args.Map{"result": c.HasUsingSensitivity("Hello", true)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "case sensitive match", actual)
		actual := args.Map{"result": c.HasUsingSensitivity("hello", true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "case sensitive no match", actual)
		actual := args.Map{"result": c.HasUsingSensitivity("hello", false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "case insensitive match", actual)
	})
}

func Test_C35_Collection_IsContainsPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsContainsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		a := "a"
		actual := args.Map{"result": c.IsContainsPtr(&a)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.IsContainsPtr(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil should be false", actual)
	})
}

func Test_C35_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_C35_Collection_GetHashsetPlusHasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs, ok := c.GetHashsetPlusHasAll([]string{"a", "b"})
		actual := args.Map{"result": ok || hs == nil}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		_, ok2 := c.GetHashsetPlusHasAll(nil)
		actual := args.Map{"result": ok2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_C35_Collection_IsContainsAllSlice(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsContainsAllSlice", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.IsContainsAllSlice([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.IsContainsAllSlice([]string{})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_C35_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsContainsAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.IsContainsAll("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.IsContainsAll()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_C35_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsContainsAllLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.IsContainsAllLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C35_Collection_New(t *testing.T) {
	safeTest(t, "Test_C35_Collection_New", func() {
		c := corestr.New.Collection.Empty()
		n := c.New("a", "b")
		actual := args.Map{"result": n.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		e := c.New()
		actual := args.Map{"result": e.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C35_Collection_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddNonEmptyStrings", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings("a", "", "b")
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		c.AddNonEmptyStrings()
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddFuncResult", func() {
		c := corestr.New.Collection.Empty()
		c.AddFuncResult(func() string { return "a" }, func() string { return "b" })
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		c.AddFuncResult()
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (nil)", actual)
	})
}

func Test_C35_Collection_AddNonEmptyStringsSlice(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddNonEmptyStringsSlice", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStringsSlice([]string{"a"})
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AddNonEmptyStringsSlice(nil)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddStringsByFuncChecking", func() {
		c := corestr.New.Collection.Empty()
		c.AddStringsByFuncChecking([]string{"a", "bb", "c"}, func(s string) bool {
			return len(s) == 1
		})
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ExpandSlicePlusAdd", func() {
		c := corestr.New.Collection.Empty()
		c.ExpandSlicePlusAdd([]string{"a,b"}, func(s string) []string {
			return []string{s + "1", s + "2"}
		})
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_C35_Collection_MergeSlicesOfSlice", func() {
		c := corestr.New.Collection.Empty()
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_C35_Collection_GetAllExceptCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		except := corestr.New.Collection.Strings([]string{"b"})
		result := c.GetAllExceptCollection(except)
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// nil except
		all := c.GetAllExceptCollection(nil)
		actual := args.Map{"result": len(all) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C35_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_C35_Collection_GetAllExcept", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.GetAllExcept([]string{"a"})
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		all := c.GetAllExcept(nil)
		actual := args.Map{"result": len(all) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_C35_Collection_CharCollectionMap", func() {
		c := corestr.New.Collection.Strings([]string{"apple", "avocado", "banana"})
		cm := c.CharCollectionMap()
		actual := args.Map{"result": cm.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 chars", actual)
	})
}

func Test_C35_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SummaryString", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.SummaryString(1) == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C35_Collection_SummaryStringWithHeader(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SummaryStringWithHeader", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.SummaryStringWithHeader("Header") == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		e := corestr.New.Collection.Empty()
		actual := args.Map{"result": e.SummaryStringWithHeader("Header") == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty for empty collection", actual)
	})
}

func Test_C35_Collection_String(t *testing.T) {
	safeTest(t, "Test_C35_Collection_String", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		e := corestr.New.Collection.Empty()
		actual := args.Map{"result": e.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty for empty", actual)
	})
}

func Test_C35_Collection_StringLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_StringLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.StringLock() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		e := corestr.New.Collection.Empty()
		actual := args.Map{"result": e.StringLock() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty for empty", actual)
	})
}

func Test_C35_Collection_CsvLines(t *testing.T) {
	safeTest(t, "Test_C35_Collection_CsvLines", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		lines := c.CsvLines()
		actual := args.Map{"result": len(lines) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_Collection_CsvLinesOptions(t *testing.T) {
	safeTest(t, "Test_C35_Collection_CsvLinesOptions", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		lines := c.CsvLinesOptions(true)
		actual := args.Map{"result": len(lines) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_Csv(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Csv", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.Csv() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		e := corestr.New.Collection.Empty()
		actual := args.Map{"result": e.Csv() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C35_Collection_CsvOptions(t *testing.T) {
	safeTest(t, "Test_C35_Collection_CsvOptions", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.CsvOptions(false) == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		e := corestr.New.Collection.Empty()
		actual := args.Map{"result": e.CsvOptions(false) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C35_Collection_AddCapacity(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddCapacity", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddCapacity(10)
		actual := args.Map{"result": c.Capacity() < 12}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected capacity >= 12", actual)
		c.AddCapacity()
		// no-op
	})
}

func Test_C35_Collection_Resize(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Resize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Resize(100)
		actual := args.Map{"result": c.Capacity() < 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected capacity >= 100", actual)
		// smaller resize is no-op
		oldCap := c.Capacity()
		c.Resize(1)
		actual := args.Map{"result": c.Capacity() != oldCap}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not shrink", actual)
	})
}

func Test_C35_Collection_Joins(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Joins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.Joins(",") != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
		withExtra := c.Joins(",", "c")
		actual := args.Map{"result": withExtra == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C35_Collection_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyJoins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		result := c.NonEmptyJoins(",")
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C35_Collection_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonWhitespaceJoins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "  ", "b"})
		result := c.NonWhitespaceJoins(",")
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C35_Collection_JsonModel(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JsonModel", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": len(c.JsonModel()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JsonModelAny", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C35_Collection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C35_Collection_MarshalJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		data, err := c.MarshalJSON()
		actual := args.Map{"result": err != nil || len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected marshal success", actual)
	})
}

func Test_C35_Collection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C35_Collection_UnmarshalJSON", func() {
		c := &corestr.Collection{}
		err := c.UnmarshalJSON([]byte(`["a","b"]`))
		actual := args.Map{"result": err != nil || c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected unmarshal success", actual)
		err2 := c.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"result": err2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C35_Collection_Json(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Json", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.Json()
		actual := args.Map{"result": r.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_C35_Collection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JsonPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.JsonPtr()
		actual := args.Map{"result": r == nil || r.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_C35_Collection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ParseInjectUsingJson", func() {
		orig := corestr.New.Collection.Strings([]string{"a", "b"})
		jsonR := orig.JsonPtr()
		target := &corestr.Collection{}
		result, err := target.ParseInjectUsingJson(jsonR)
		actual := args.Map{"result": err != nil || result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected success", actual)
		// error case
		badJson := corejson.NewPtr("invalid{")
		_, err2 := target.ParseInjectUsingJson(badJson)
		actual := args.Map{"result": err2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}
func Test_C35_Collection_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ParseInjectUsingJsonMust", func() {
		orig := corestr.New.Collection.Strings([]string{"a"})
		jsonR := orig.JsonPtr()
		target := &corestr.Collection{}
		result := target.ParseInjectUsingJsonMust(jsonR)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_Collection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JsonParseSelfInject", func() {
		orig := corestr.New.Collection.Strings([]string{"a"})
		jsonR := orig.JsonPtr()
		target := &corestr.Collection{}
		err := target.JsonParseSelfInject(jsonR)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_C35_Collection_Clear(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Clear", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Clear()
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		var nilC *corestr.Collection
		actual := args.Map{"result": nilC.Clear() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil clear should return nil", actual)
	})
}

func Test_C35_Collection_Dispose(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Dispose", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Dispose()
		var nilC *corestr.Collection
		nilC.Dispose() // should not panic
	})
}

func Test_C35_Collection_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AsJsonMarshaller", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C35_Collection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AsJsonContractsBinder", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C35_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Serialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		data, err := c.Serialize()
		actual := args.Map{"result": err != nil || len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected success", actual)
	})
}

func Test_C35_Collection_Deserialize(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Deserialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		var target []string
		err := c.Deserialize(&target)
		actual := args.Map{"result": err != nil || len(target) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected success", actual)
	})
}

func Test_C35_Collection_Join(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Join", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.Join(",") != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
		e := corestr.New.Collection.Empty()
		actual := args.Map{"result": e.Join(",") != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C35_Collection_JoinLine(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JoinLine", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.JoinLine()
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		e := corestr.New.Collection.Empty()
		actual := args.Map{"result": e.JoinLine() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C35_Collection_List(t *testing.T) {
	safeTest(t, "Test_C35_Collection_List", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": len(c.List()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

	// ===== newCollectionCreator =====

func Test_C35_NewCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_Empty", func() {
		c := corestr.New.Collection.Empty()
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C35_NewCollection_Cap(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_Cap", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 length", actual)
	})
}

func Test_C35_NewCollection_CloneStrings(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_CloneStrings", func() {
		c := corestr.New.Collection.CloneStrings([]string{"a", "b"})
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_NewCollection_Create(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_Create", func() {
		c := corestr.New.Collection.Create([]string{"a"})
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_NewCollection_StringsOptions_Clone(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_StringsOptions_Clone", func() {
		c := corestr.New.Collection.StringsOptions(true, []string{"a"})
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_NewCollection_StringsOptions_NoClone_Empty(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_StringsOptions_NoClone_Empty", func() {
		c := corestr.New.Collection.StringsOptions(false, []string{})
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C35_NewCollection_LineUsingSep(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_LineUsingSep", func() {
		c := corestr.New.Collection.LineUsingSep(",", "a,b,c")
		actual := args.Map{"result": c.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C35_NewCollection_LineDefault(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_LineDefault", func() {
		c := corestr.New.Collection.LineDefault("a\nb")
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C35_NewCollection_StringsPlusCap(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_StringsPlusCap", func() {
		c := corestr.New.Collection.StringsPlusCap(5, []string{"a"})
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c2 := corestr.New.Collection.StringsPlusCap(0, []string{"a"})
		actual := args.Map{"result": c2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_NewCollection_CapStrings(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_CapStrings", func() {
		c := corestr.New.Collection.CapStrings(5, []string{"a"})
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c2 := corestr.New.Collection.CapStrings(0, []string{"a"})
		actual := args.Map{"result": c2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C35_NewCollection_LenCap(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_LenCap", func() {
		c := corestr.New.Collection.LenCap(3, 10)
		actual := args.Map{"result": c.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

	// ===== Helpers =====

func Test_C35_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_C35_AllIndividualStringsOfStringsLength", func() {
		items := [][]string{{"a", "b"}, {"c"}}
		actual := args.Map{"result": corestr.AllIndividualStringsOfStringsLength(&items) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual := args.Map{"result": corestr.AllIndividualStringsOfStringsLength(nil) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C35_AnyToString(t *testing.T) {
	safeTest(t, "Test_C35_AnyToString", func() {
		result := corestr.AnyToString(false, 42)
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		resultWithField := corestr.AnyToString(true, 42)
		actual := args.Map{"result": resultWithField == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual := args.Map{"result": corestr.AnyToString(false, "") != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for empty string", actual)
	})
}

func Test_C35_CloneSlice(t *testing.T) {
	safeTest(t, "Test_C35_CloneSlice", func() {
		orig := []string{"a", "b"}
		cloned := corestr.CloneSlice(orig)
		actual := args.Map{"result": len(cloned) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual := args.Map{"result": len(corestr.CloneSlice(nil)) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}
