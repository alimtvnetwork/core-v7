package corestrtests

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ===== Collection — core methods =====

func Test_C35_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JsonString", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		// Act
		result := c.JsonString()
		// Assert
		if result == "" {
			t.Error("expected non-empty JSON string")
		}
	})
}

func Test_C35_Collection_JsonStringMust(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JsonStringMust", func() {
		c := corestr.New.Collection.Strings([]string{"x"})
		result := c.JsonStringMust()
		if result == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C35_Collection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasAnyItem", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.HasAnyItem() {
			t.Error("expected HasAnyItem true")
		}
		e := corestr.New.Collection.Empty()
		if e.HasAnyItem() {
			t.Error("expected HasAnyItem false on empty")
		}
	})
}

func Test_C35_Collection_LastIndex(t *testing.T) {
	safeTest(t, "Test_C35_Collection_LastIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if c.LastIndex() != 2 {
			t.Errorf("expected 2 got %d", c.LastIndex())
		}
	})
}

func Test_C35_Collection_HasIndex(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.HasIndex(0) {
			t.Error("expected HasIndex(0) true")
		}
		if !c.HasIndex(1) {
			t.Error("expected HasIndex(1) true")
		}
		if c.HasIndex(2) {
			t.Error("expected HasIndex(2) false")
		}
		if c.HasIndex(-1) {
			t.Error("expected HasIndex(-1) false")
		}
	})
}

func Test_C35_Collection_ListStringsPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ListStringsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if len(c.ListStringsPtr()) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_ListStrings(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ListStrings", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if len(c.ListStrings()) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_StringJSON(t *testing.T) {
	safeTest(t, "Test_C35_Collection_StringJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.StringJSON() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C35_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_C35_Collection_RemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := c.RemoveAt(1)
		if !ok || c.Length() != 2 {
			t.Error("expected remove success")
		}
		if c.RemoveAt(-1) {
			t.Error("expected false for negative index")
		}
		if c.RemoveAt(99) {
			t.Error("expected false for out of range")
		}
	})
}

func Test_C35_Collection_Count(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Count", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.Count() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_Capacity(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Capacity", func() {
		c := corestr.New.Collection.Cap(10)
		if c.Capacity() < 10 {
			t.Error("expected capacity >= 10")
		}
		var nilC *corestr.Collection
		if nilC.Capacity() != 0 {
			t.Error("expected 0 for nil items")
		}
	})
}

func Test_C35_Collection_Length_NilReceiver(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Length_NilReceiver", func() {
		var c *corestr.Collection
		if c.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C35_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_LengthLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.LengthLock() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEquals", func() {
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		c := corestr.New.Collection.Strings([]string{"a", "c"})
		if !a.IsEquals(b) {
			t.Error("expected equal")
		}
		if a.IsEquals(c) {
			t.Error("expected not equal")
		}
	})
}

func Test_C35_Collection_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive", func() {
		a := corestr.New.Collection.Strings([]string{"Hello"})
		b := corestr.New.Collection.Strings([]string{"hello"})
		if a.IsEqualsWithSensitive(true, b) {
			t.Error("expected case sensitive not equal")
		}
		if !a.IsEqualsWithSensitive(false, b) {
			t.Error("expected case insensitive equal")
		}
	})
}

func Test_C35_Collection_IsEqualsWithSensitive_BothNil(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive_BothNil", func() {
		var a, b *corestr.Collection
		if !a.IsEqualsWithSensitive(true, b) {
			t.Error("both nil should be equal")
		}
	})
}

func Test_C35_Collection_IsEqualsWithSensitive_OneNil(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive_OneNil", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		var b *corestr.Collection
		if a.IsEqualsWithSensitive(true, b) {
			t.Error("one nil should not equal")
		}
	})
}

func Test_C35_Collection_IsEqualsWithSensitive_SamePtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive_SamePtr", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		if !a.IsEqualsWithSensitive(true, a) {
			t.Error("same ptr should be equal")
		}
	})
}

func Test_C35_Collection_IsEqualsWithSensitive_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive_BothEmpty", func() {
		a := corestr.New.Collection.Empty()
		b := corestr.New.Collection.Empty()
		if !a.IsEqualsWithSensitive(true, b) {
			t.Error("both empty should be equal")
		}
	})
}

func Test_C35_Collection_IsEqualsWithSensitive_OneEmpty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive_OneEmpty", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Empty()
		if a.IsEqualsWithSensitive(true, b) {
			t.Error("one empty should not equal")
		}
	})
}

func Test_C35_Collection_IsEqualsWithSensitive_DiffLength(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEqualsWithSensitive_DiffLength", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		if a.IsEqualsWithSensitive(true, b) {
			t.Error("diff length should not equal")
		}
	})
}

func Test_C35_Collection_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEmptyLock", func() {
		c := corestr.New.Collection.Empty()
		if !c.IsEmptyLock() {
			t.Error("expected empty")
		}
	})
}

func Test_C35_Collection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsEmpty", func() {
		var c *corestr.Collection
		if !c.IsEmpty() {
			t.Error("nil should be empty")
		}
	})
}

func Test_C35_Collection_HasItems(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasItems", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.HasItems() {
			t.Error("expected HasItems true")
		}
		e := corestr.New.Collection.Empty()
		if e.HasItems() {
			t.Error("expected HasItems false")
		}
	})
}

func Test_C35_Collection_AddLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddLock", func() {
		c := corestr.New.Collection.Empty()
		c.AddLock("a")
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddNonEmpty", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmpty("")
		c.AddNonEmpty("a")
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddNonEmptyWhitespace", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyWhitespace("  ")
		c.AddNonEmptyWhitespace("a")
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_Add(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Add", func() {
		c := corestr.New.Collection.Empty()
		c.Add("a")
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddError", func() {
		c := corestr.New.Collection.Empty()
		c.AddError(nil)
		c.AddError(errors.New("err1"))
		if c.Length() != 1 {
			t.Error("expected 1")
		}
		if c.First() != "err1" {
			t.Error("expected err1")
		}
	})
}

func Test_C35_Collection_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AsDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		err := c.AsDefaultError()
		if err == nil {
			t.Error("expected error")
		}
	})
}

func Test_C35_Collection_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AsError_Empty", func() {
		c := corestr.New.Collection.Empty()
		if c.AsError(",") != nil {
			t.Error("expected nil")
		}
	})
}

func Test_C35_Collection_AsError_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AsError_NonEmpty", func() {
		c := corestr.New.Collection.Strings([]string{"err"})
		if c.AsError(",") == nil {
			t.Error("expected error")
		}
	})
}

func Test_C35_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddIf", func() {
		c := corestr.New.Collection.Empty()
		c.AddIf(false, "skip")
		c.AddIf(true, "keep")
		if c.Length() != 1 || c.First() != "keep" {
			t.Error("expected only keep")
		}
	})
}

func Test_C35_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_C35_Collection_EachItemSplitBy", func() {
		c := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
		result := c.EachItemSplitBy(",")
		if len(result) != 4 {
			t.Errorf("expected 4 got %d", len(result))
		}
	})
}

func Test_C35_Collection_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ConcatNew_Empty", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		newC := c.ConcatNew(0)
		if newC.Length() != 1 {
			t.Error("expected cloned with 1 item")
		}
	})
}

func Test_C35_Collection_ConcatNew_WithItems(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ConcatNew_WithItems", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		newC := c.ConcatNew(0, "b", "c")
		if newC.Length() != 3 {
			t.Errorf("expected 3 got %d", newC.Length())
		}
	})
}

func Test_C35_Collection_ToError(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ToError", func() {
		c := corestr.New.Collection.Strings([]string{"err1"})
		if c.ToError(",") == nil {
			t.Error("expected error")
		}
	})
}

func Test_C35_Collection_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ToDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"err1"})
		if c.ToDefaultError() == nil {
			t.Error("expected error")
		}
	})
}

func Test_C35_Collection_AddIfMany(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddIfMany", func() {
		c := corestr.New.Collection.Empty()
		c.AddIfMany(false, "a", "b")
		if c.Length() != 0 {
			t.Error("expected 0")
		}
		c.AddIfMany(true, "a", "b")
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddFunc", func() {
		c := corestr.New.Collection.Empty()
		c.AddFunc(func() string { return "hello" })
		if c.Length() != 1 || c.First() != "hello" {
			t.Error("expected hello")
		}
	})
}

func Test_C35_Collection_AddFuncErr_Success(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddFuncErr_Success", func() {
		c := corestr.New.Collection.Empty()
		c.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(e error) {},
		)
		if c.Length() != 1 {
			t.Error("expected 1")
		}
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
		if c.Length() != 0 {
			t.Error("expected 0")
		}
		if caught == nil {
			t.Error("expected error caught")
		}
	})
}

func Test_C35_Collection_AddsLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddsLock", func() {
		c := corestr.New.Collection.Empty()
		c.AddsLock("a", "b")
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Adds", func() {
		c := corestr.New.Collection.Empty()
		c.Adds("a", "b", "c")
		if c.Length() != 3 {
			t.Error("expected 3")
		}
	})
}

func Test_C35_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddStrings", func() {
		c := corestr.New.Collection.Empty()
		c.AddStrings([]string{"x", "y"})
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddCollection", func() {
		c := corestr.New.Collection.Empty()
		other := corestr.New.Collection.Strings([]string{"a"})
		c.AddCollection(other)
		if c.Length() != 1 {
			t.Error("expected 1")
		}
		c.AddCollection(corestr.New.Collection.Empty())
		if c.Length() != 1 {
			t.Error("should skip empty")
		}
	})
}

func Test_C35_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddCollections", func() {
		c := corestr.New.Collection.Empty()
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"b"})
		c.AddCollections(a, b)
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddPointerCollectionsLock", func() {
		c := corestr.New.Collection.Empty()
		a := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(a)
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddHashmapsValues", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		c.AddHashmapsValues(h)
		if c.Length() != 1 {
			t.Error("expected 1")
		}
		c.AddHashmapsValues(nil)
		if c.Length() != 1 {
			t.Error("should skip nil hashmaps")
		}
	})
}

func Test_C35_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddHashmapsKeys", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k1", "v1")
		c.AddHashmapsKeys(h)
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddHashmapsKeysValues", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValues(h)
		if c.Length() != 2 {
			t.Error("expected 2 (key + value)")
		}
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
		if c.Length() != 1 {
			t.Error("expected 1")
		}
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
		if c.Length() != 1 {
			t.Error("expected 1 due to break")
		}
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
		if c.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C35_Collection_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddWithWgLock", func() {
		c := corestr.New.Collection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(wg, "a")
		wg.Wait()
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IndexAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.IndexAt(1) != "b" {
			t.Error("expected b")
		}
	})
}

func Test_C35_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SafeIndexAtUsingLength", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.SafeIndexAtUsingLength("default", 2, 1) != "b" {
			t.Error("expected b")
		}
		if c.SafeIndexAtUsingLength("default", 1, 5) != "default" {
			t.Error("expected default")
		}
	})
}

func Test_C35_Collection_First(t *testing.T) {
	safeTest(t, "Test_C35_Collection_First", func() {
		c := corestr.New.Collection.Strings([]string{"first", "second"})
		if c.First() != "first" {
			t.Error("expected first")
		}
	})
}

func Test_C35_Collection_Single(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Single", func() {
		c := corestr.New.Collection.Strings([]string{"only"})
		if c.Single() != "only" {
			t.Error("expected only")
		}
	})
}

func Test_C35_Collection_Last(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Last", func() {
		c := corestr.New.Collection.Strings([]string{"a", "last"})
		if c.Last() != "last" {
			t.Error("expected last")
		}
	})
}

func Test_C35_Collection_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_C35_Collection_LastOrDefault", func() {
		c := corestr.New.Collection.Strings([]string{"a", "last"})
		if c.LastOrDefault() != "last" {
			t.Error("expected last")
		}
		e := corestr.New.Collection.Empty()
		if e.LastOrDefault() != "" {
			t.Error("expected empty string")
		}
	})
}

func Test_C35_Collection_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FirstOrDefault", func() {
		c := corestr.New.Collection.Strings([]string{"first"})
		if c.FirstOrDefault() != "first" {
			t.Error("expected first")
		}
		e := corestr.New.Collection.Empty()
		if e.FirstOrDefault() != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C35_Collection_Take(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Take", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		taken := c.Take(2)
		if taken.Length() != 2 {
			t.Error("expected 2")
		}
		all := c.Take(10)
		if all.Length() != 3 {
			t.Error("take more than length returns all")
		}
		zero := c.Take(0)
		if zero.Length() != 0 {
			t.Error("take 0 returns empty")
		}
	})
}

func Test_C35_Collection_Skip(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Skip", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		skipped := c.Skip(1)
		if skipped.Length() != 2 {
			t.Error("expected 2")
		}
		none := c.Skip(0)
		if none != c {
			t.Error("skip 0 returns self")
		}
	})
}

func Test_C35_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Reverse", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()
		if c.First() != "c" || c.Last() != "a" {
			t.Error("expected reversed")
		}
		// len 2
		c2 := corestr.New.Collection.Strings([]string{"x", "y"})
		c2.Reverse()
		if c2.First() != "y" {
			t.Error("expected y")
		}
		// len 1
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c1.Reverse()
		if c1.First() != "a" {
			t.Error("expected a")
		}
	})
}

func Test_C35_Collection_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_C35_Collection_GetPagesSize", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		if c.GetPagesSize(2) != 3 {
			t.Error("expected 3 pages")
		}
		if c.GetPagesSize(0) != 0 {
			t.Error("expected 0 for zero page size")
		}
		if c.GetPagesSize(-1) != 0 {
			t.Error("expected 0 for negative page size")
		}
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
		if paged.Length() != 4 {
			t.Errorf("expected 4 pages got %d", paged.Length())
		}
	})
}

func Test_C35_Collection_GetPagedCollection_SmallCollection(t *testing.T) {
	safeTest(t, "Test_C35_Collection_GetPagedCollection_SmallCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		paged := c.GetPagedCollection(10)
		if paged.Length() != 1 {
			t.Error("expected 1 page")
		}
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
		if page.Length() != 5 {
			t.Error("expected 5")
		}
		// Small collection returns self
		small := corestr.New.Collection.Strings([]string{"a"})
		if small.GetSinglePageCollection(10, 1) != small {
			t.Error("expected self")
		}
	})
}

func Test_C35_Collection_InsertAt(t *testing.T) {
	safeTest(t, "Test_C35_Collection_InsertAt", func() {
		c := corestr.New.Collection.Empty()
		c.InsertAt(0, "first")
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ChainRemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(1)
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_RemoveItemsIndexes(t *testing.T) {
	safeTest(t, "Test_C35_Collection_RemoveItemsIndexes", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.RemoveItemsIndexes(true, 1)
		if c.Length() != 2 {
			t.Error("expected 2")
		}
		// nil indexes
		c.RemoveItemsIndexes(true)
		if c.Length() != 2 {
			t.Error("expected 2 (nil indexes)")
		}
	})
}

func Test_C35_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendCollectionPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		other := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollectionPtr(other)
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_AppendCollections(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendCollections", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollections(b)
		if c.Length() != 2 {
			t.Error("expected 2")
		}
		c.AppendCollections()
		if c.Length() != 2 {
			t.Error("expected 2 (empty)")
		}
	})
}

func Test_C35_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnys(42, "hello", nil)
		if c.Length() != 2 {
			t.Error("expected 2 (nil skipped)")
		}
	})
}

func Test_C35_Collection_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendAnysLock", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysLock(1, 2)
		if c.Length() != 2 {
			t.Error("expected 2")
		}
		c.AppendAnysLock()
		if c.Length() != 2 {
			t.Error("expected 2 (empty)")
		}
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
		if c.Length() != 2 {
			t.Error("expected 2")
		}
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
		if c.Length() != 1 {
			t.Error("expected 1 due to break")
		}
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
		if c.Length() != 0 {
			t.Error("expected 0")
		}
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
		if c.Length() != 1 {
			t.Error("expected 1")
		}
		c2 := corestr.New.Collection.Empty()
		c2.AppendAnysUsingFilterLock(nil)
		if c2.Length() != 0 {
			t.Error("expected 0")
		}
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
		if c.Length() != 0 {
			t.Error("expected 0 (skip + break)")
		}
	})
}

func Test_C35_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AppendNonEmptyAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendNonEmptyAnys("a", nil, "")
		if c.Length() != 1 {
			t.Errorf("expected 1 got %d", c.Length())
		}
	})
}

func Test_C35_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddsNonEmpty", func() {
		c := corestr.New.Collection.Empty()
		c.AddsNonEmpty("a", "", "b")
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_AddsNonEmptyPtrLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddsNonEmptyPtrLock", func() {
		c := corestr.New.Collection.Empty()
		a := "a"
		empty := ""
		c.AddsNonEmptyPtrLock(&a, nil, &empty)
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_C35_Collection_UniqueBoolMap", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		m := c.UniqueBoolMap()
		if len(m) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_UniqueBoolMapLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_UniqueBoolMapLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		m := c.UniqueBoolMapLock()
		if len(m) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_C35_Collection_UniqueList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		l := c.UniqueList()
		if len(l) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_UniqueListLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		l := c.UniqueListLock()
		if len(l) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Filter", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, s != "b", false
		})
		if len(result) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_Filter_Break(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Filter_Break", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, true, i == 0
		})
		if len(result) != 1 {
			t.Error("expected 1 (break at 0)")
		}
	})
}

func Test_C35_Collection_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Filter_Empty", func() {
		c := corestr.New.Collection.Empty()
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if len(result) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C35_Collection_FilterLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FilterLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if len(result) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FilteredCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		fc := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if fc.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_FilteredCollectionLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FilteredCollectionLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		fc := c.FilteredCollectionLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if fc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_FilterPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FilterPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.FilterPtr(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		if len(*result) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_FilterPtr_Empty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FilterPtr_Empty", func() {
		c := corestr.New.Collection.Empty()
		result := c.FilterPtr(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		if len(*result) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C35_Collection_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_FilterPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.FilterPtrLock(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		if len(*result) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyList", func() {
		c := corestr.New.Collection.Strings([]string{"", "a", "", "b"})
		list := c.NonEmptyList()
		if len(list) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_NonEmptyList_Empty(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyList_Empty", func() {
		c := corestr.New.Collection.Empty()
		if len(c.NonEmptyList()) != 0 {
			t.Error("expected 0")
		}
	})
}
func Test_C35_Collection_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HashsetAsIs", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		hs := c.HashsetAsIs()
		if hs.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_HashsetWithDoubleLength(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HashsetWithDoubleLength", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetWithDoubleLength()
		if hs.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_HashsetLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HashsetLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetLock()
		if hs.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyItems", func() {
		c := corestr.New.Collection.Strings([]string{"", "a"})
		if len(c.NonEmptyItems()) != 1 {
			t.Error("expected 1")
		}
	})
}
func Test_C35_Collection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyItemsOrNonWhitespace", func() {
		c := corestr.New.Collection.Strings([]string{"  ", "a"})
		if len(c.NonEmptyItemsOrNonWhitespace()) != 1 {
			t.Error("expected 1")
		}
	})
}
func Test_C35_Collection_Items(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Items", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if len(c.Items()) != 1 {
			t.Error("expected 1")
		}
	})
}
func Test_C35_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ListCopyPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		l := c.ListCopyPtrLock()
		if len(l) != 1 {
			t.Error("expected 1")
		}
		e := corestr.New.Collection.Empty()
		if len(e.ListCopyPtrLock()) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C35_Collection_Has(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Has", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.Has("a") {
			t.Error("expected has a")
		}
		if c.Has("z") {
			t.Error("expected not has z")
		}
	})
}

func Test_C35_Collection_HasLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.HasLock("a") {
			t.Error("expected has a")
		}
	})
}

func Test_C35_Collection_HasPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		a := "a"
		if !c.HasPtr(&a) {
			t.Error("expected has a")
		}
		if c.HasPtr(nil) {
			t.Error("expected false for nil")
		}
	})
}

func Test_C35_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if !c.HasAll("a", "b") {
			t.Error("expected has all")
		}
		if c.HasAll("a", "z") {
			t.Error("expected not has all")
		}
	})
}

func Test_C35_Collection_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SortedListAsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		list := c.SortedListAsc()
		if list[0] != "a" {
			t.Error("expected a first")
		}
	})
}

func Test_C35_Collection_SortedAsc(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SortedAsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a"})
		c.SortedAsc()
		if c.First() != "a" {
			t.Error("expected a")
		}
	})
}

func Test_C35_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SortedAscLock", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a"})
		c.SortedAscLock()
		if c.First() != "a" {
			t.Error("expected a")
		}
	})
}

func Test_C35_Collection_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SortedListDsc", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		list := c.SortedListDsc()
		if list[0] != "c" {
			t.Error("expected c first")
		}
	})
}

func Test_C35_Collection_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_C35_Collection_HasUsingSensitivity", func() {
		c := corestr.New.Collection.Strings([]string{"Hello"})
		if !c.HasUsingSensitivity("Hello", true) {
			t.Error("case sensitive match")
		}
		if c.HasUsingSensitivity("hello", true) {
			t.Error("case sensitive no match")
		}
		if !c.HasUsingSensitivity("hello", false) {
			t.Error("case insensitive match")
		}
	})
}

func Test_C35_Collection_IsContainsPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsContainsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		a := "a"
		if !c.IsContainsPtr(&a) {
			t.Error("expected true")
		}
		if c.IsContainsPtr(nil) {
			t.Error("nil should be false")
		}
	})
}

func Test_C35_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_C35_Collection_GetHashsetPlusHasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs, ok := c.GetHashsetPlusHasAll([]string{"a", "b"})
		if !ok || hs == nil {
			t.Error("expected true")
		}
		_, ok2 := c.GetHashsetPlusHasAll(nil)
		if ok2 {
			t.Error("expected false for nil")
		}
	})
}

func Test_C35_Collection_IsContainsAllSlice(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsContainsAllSlice", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.IsContainsAllSlice([]string{"a", "b"}) {
			t.Error("expected true")
		}
		if c.IsContainsAllSlice([]string{}) {
			t.Error("expected false for empty")
		}
	})
}

func Test_C35_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsContainsAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.IsContainsAll("a") {
			t.Error("expected true")
		}
		if c.IsContainsAll() {
			t.Error("expected false for nil")
		}
	})
}

func Test_C35_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_IsContainsAllLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.IsContainsAllLock("a") {
			t.Error("expected true")
		}
	})
}

func Test_C35_Collection_New(t *testing.T) {
	safeTest(t, "Test_C35_Collection_New", func() {
		c := corestr.New.Collection.Empty()
		n := c.New("a", "b")
		if n.Length() != 2 {
			t.Error("expected 2")
		}
		e := c.New()
		if e.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C35_Collection_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddNonEmptyStrings", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings("a", "", "b")
		if c.Length() != 2 {
			t.Error("expected 2")
		}
		c.AddNonEmptyStrings()
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddFuncResult", func() {
		c := corestr.New.Collection.Empty()
		c.AddFuncResult(func() string { return "a" }, func() string { return "b" })
		if c.Length() != 2 {
			t.Error("expected 2")
		}
		c.AddFuncResult()
		if c.Length() != 2 {
			t.Error("expected 2 (nil)")
		}
	})
}

func Test_C35_Collection_AddNonEmptyStringsSlice(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddNonEmptyStringsSlice", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStringsSlice([]string{"a"})
		if c.Length() != 1 {
			t.Error("expected 1")
		}
		c.AddNonEmptyStringsSlice(nil)
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddStringsByFuncChecking", func() {
		c := corestr.New.Collection.Empty()
		c.AddStringsByFuncChecking([]string{"a", "bb", "c"}, func(s string) bool {
			return len(s) == 1
		})
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ExpandSlicePlusAdd", func() {
		c := corestr.New.Collection.Empty()
		c.ExpandSlicePlusAdd([]string{"a,b"}, func(s string) []string {
			return []string{s + "1", s + "2"}
		})
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_C35_Collection_MergeSlicesOfSlice", func() {
		c := corestr.New.Collection.Empty()
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_C35_Collection_GetAllExceptCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		except := corestr.New.Collection.Strings([]string{"b"})
		result := c.GetAllExceptCollection(except)
		if len(result) != 2 {
			t.Error("expected 2")
		}
		// nil except
		all := c.GetAllExceptCollection(nil)
		if len(all) != 3 {
			t.Error("expected 3")
		}
	})
}

func Test_C35_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_C35_Collection_GetAllExcept", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.GetAllExcept([]string{"a"})
		if len(result) != 1 {
			t.Error("expected 1")
		}
		all := c.GetAllExcept(nil)
		if len(all) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_C35_Collection_CharCollectionMap", func() {
		c := corestr.New.Collection.Strings([]string{"apple", "avocado", "banana"})
		cm := c.CharCollectionMap()
		if cm.Length() != 2 {
			t.Error("expected 2 chars")
		}
	})
}

func Test_C35_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SummaryString", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.SummaryString(1) == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C35_Collection_SummaryStringWithHeader(t *testing.T) {
	safeTest(t, "Test_C35_Collection_SummaryStringWithHeader", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.SummaryStringWithHeader("Header") == "" {
			t.Error("expected non-empty")
		}
		e := corestr.New.Collection.Empty()
		if e.SummaryStringWithHeader("Header") == "" {
			t.Error("expected non-empty for empty collection")
		}
	})
}

func Test_C35_Collection_String(t *testing.T) {
	safeTest(t, "Test_C35_Collection_String", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.String() == "" {
			t.Error("expected non-empty")
		}
		e := corestr.New.Collection.Empty()
		if e.String() == "" {
			t.Error("expected non-empty for empty")
		}
	})
}

func Test_C35_Collection_StringLock(t *testing.T) {
	safeTest(t, "Test_C35_Collection_StringLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.StringLock() == "" {
			t.Error("expected non-empty")
		}
		e := corestr.New.Collection.Empty()
		if e.StringLock() == "" {
			t.Error("expected non-empty for empty")
		}
	})
}

func Test_C35_Collection_CsvLines(t *testing.T) {
	safeTest(t, "Test_C35_Collection_CsvLines", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		lines := c.CsvLines()
		if len(lines) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_Collection_CsvLinesOptions(t *testing.T) {
	safeTest(t, "Test_C35_Collection_CsvLinesOptions", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		lines := c.CsvLinesOptions(true)
		if len(lines) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_Csv(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Csv", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.Csv() == "" {
			t.Error("expected non-empty")
		}
		e := corestr.New.Collection.Empty()
		if e.Csv() != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C35_Collection_CsvOptions(t *testing.T) {
	safeTest(t, "Test_C35_Collection_CsvOptions", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.CsvOptions(false) == "" {
			t.Error("expected non-empty")
		}
		e := corestr.New.Collection.Empty()
		if e.CsvOptions(false) != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C35_Collection_AddCapacity(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AddCapacity", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddCapacity(10)
		if c.Capacity() < 12 {
			t.Error("expected capacity >= 12")
		}
		c.AddCapacity()
		// no-op
	})
}

func Test_C35_Collection_Resize(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Resize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Resize(100)
		if c.Capacity() < 100 {
			t.Error("expected capacity >= 100")
		}
		// smaller resize is no-op
		oldCap := c.Capacity()
		c.Resize(1)
		if c.Capacity() != oldCap {
			t.Error("should not shrink")
		}
	})
}

func Test_C35_Collection_Joins(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Joins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.Joins(",") != "a,b" {
			t.Error("expected a,b")
		}
		withExtra := c.Joins(",", "c")
		if withExtra == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C35_Collection_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonEmptyJoins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		result := c.NonEmptyJoins(",")
		if result == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C35_Collection_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_C35_Collection_NonWhitespaceJoins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "  ", "b"})
		result := c.NonWhitespaceJoins(",")
		if result == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C35_Collection_JsonModel(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JsonModel", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if len(c.JsonModel()) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JsonModelAny", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.JsonModelAny() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C35_Collection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C35_Collection_MarshalJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		data, err := c.MarshalJSON()
		if err != nil || len(data) == 0 {
			t.Error("expected marshal success")
		}
	})
}

func Test_C35_Collection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C35_Collection_UnmarshalJSON", func() {
		c := &corestr.Collection{}
		err := c.UnmarshalJSON([]byte(`["a","b"]`))
		if err != nil || c.Length() != 2 {
			t.Error("expected unmarshal success")
		}
		err2 := c.UnmarshalJSON([]byte(`invalid`))
		if err2 == nil {
			t.Error("expected error")
		}
	})
}

func Test_C35_Collection_Json(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Json", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.Json()
		if r.HasError() {
			t.Error("expected no error")
		}
	})
}

func Test_C35_Collection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JsonPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.JsonPtr()
		if r == nil || r.HasError() {
			t.Error("expected no error")
		}
	})
}

func Test_C35_Collection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ParseInjectUsingJson", func() {
		orig := corestr.New.Collection.Strings([]string{"a", "b"})
		jsonR := orig.JsonPtr()
		target := &corestr.Collection{}
		result, err := target.ParseInjectUsingJson(jsonR)
		if err != nil || result.Length() != 2 {
			t.Error("expected success")
		}
		// error case
		badJson := corejson.NewPtr("invalid{")
		_, err2 := target.ParseInjectUsingJson(badJson)
		if err2 == nil {
			t.Error("expected error")
		}
	})
}
func Test_C35_Collection_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C35_Collection_ParseInjectUsingJsonMust", func() {
		orig := corestr.New.Collection.Strings([]string{"a"})
		jsonR := orig.JsonPtr()
		target := &corestr.Collection{}
		result := target.ParseInjectUsingJsonMust(jsonR)
		if result.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_Collection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JsonParseSelfInject", func() {
		orig := corestr.New.Collection.Strings([]string{"a"})
		jsonR := orig.JsonPtr()
		target := &corestr.Collection{}
		err := target.JsonParseSelfInject(jsonR)
		if err != nil {
			t.Error("expected no error")
		}
	})
}

func Test_C35_Collection_Clear(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Clear", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Clear()
		if c.Length() != 0 {
			t.Error("expected 0")
		}
		var nilC *corestr.Collection
		if nilC.Clear() != nil {
			t.Error("nil clear should return nil")
		}
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
		if c.AsJsonMarshaller() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C35_Collection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C35_Collection_AsJsonContractsBinder", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.AsJsonContractsBinder() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C35_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Serialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		data, err := c.Serialize()
		if err != nil || len(data) == 0 {
			t.Error("expected success")
		}
	})
}

func Test_C35_Collection_Deserialize(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Deserialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		var target []string
		err := c.Deserialize(&target)
		if err != nil || len(target) != 1 {
			t.Error("expected success")
		}
	})
}

func Test_C35_Collection_Join(t *testing.T) {
	safeTest(t, "Test_C35_Collection_Join", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.Join(",") != "a,b" {
			t.Error("expected a,b")
		}
		e := corestr.New.Collection.Empty()
		if e.Join(",") != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C35_Collection_JoinLine(t *testing.T) {
	safeTest(t, "Test_C35_Collection_JoinLine", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.JoinLine()
		if result == "" {
			t.Error("expected non-empty")
		}
		e := corestr.New.Collection.Empty()
		if e.JoinLine() != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C35_Collection_List(t *testing.T) {
	safeTest(t, "Test_C35_Collection_List", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if len(c.List()) != 1 {
			t.Error("expected 1")
		}
	})
}

	// ===== newCollectionCreator =====

func Test_C35_NewCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_Empty", func() {
		c := corestr.New.Collection.Empty()
		if c.Length() != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_C35_NewCollection_Cap(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_Cap", func() {
		c := corestr.New.Collection.Cap(10)
		if c.Length() != 0 {
			t.Error("expected 0 length")
		}
	})
}

func Test_C35_NewCollection_CloneStrings(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_CloneStrings", func() {
		c := corestr.New.Collection.CloneStrings([]string{"a", "b"})
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_NewCollection_Create(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_Create", func() {
		c := corestr.New.Collection.Create([]string{"a"})
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_NewCollection_StringsOptions_Clone(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_StringsOptions_Clone", func() {
		c := corestr.New.Collection.StringsOptions(true, []string{"a"})
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_NewCollection_StringsOptions_NoClone_Empty(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_StringsOptions_NoClone_Empty", func() {
		c := corestr.New.Collection.StringsOptions(false, []string{})
		if c.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C35_NewCollection_LineUsingSep(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_LineUsingSep", func() {
		c := corestr.New.Collection.LineUsingSep(",", "a,b,c")
		if c.Length() != 3 {
			t.Error("expected 3")
		}
	})
}

func Test_C35_NewCollection_LineDefault(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_LineDefault", func() {
		c := corestr.New.Collection.LineDefault("a\nb")
		if c.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C35_NewCollection_StringsPlusCap(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_StringsPlusCap", func() {
		c := corestr.New.Collection.StringsPlusCap(5, []string{"a"})
		if c.Length() != 1 {
			t.Error("expected 1")
		}
		c2 := corestr.New.Collection.StringsPlusCap(0, []string{"a"})
		if c2.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_NewCollection_CapStrings(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_CapStrings", func() {
		c := corestr.New.Collection.CapStrings(5, []string{"a"})
		if c.Length() != 1 {
			t.Error("expected 1")
		}
		c2 := corestr.New.Collection.CapStrings(0, []string{"a"})
		if c2.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C35_NewCollection_LenCap(t *testing.T) {
	safeTest(t, "Test_C35_NewCollection_LenCap", func() {
		c := corestr.New.Collection.LenCap(3, 10)
		if c.Length() != 3 {
			t.Error("expected 3")
		}
	})
}

	// ===== Helpers =====

func Test_C35_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_C35_AllIndividualStringsOfStringsLength", func() {
		items := [][]string{{"a", "b"}, {"c"}}
		if corestr.AllIndividualStringsOfStringsLength(&items) != 3 {
			t.Error("expected 3")
		}
		if corestr.AllIndividualStringsOfStringsLength(nil) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C35_AnyToString(t *testing.T) {
	safeTest(t, "Test_C35_AnyToString", func() {
		result := corestr.AnyToString(false, 42)
		if result == "" {
			t.Error("expected non-empty")
		}
		resultWithField := corestr.AnyToString(true, 42)
		if resultWithField == "" {
			t.Error("expected non-empty")
		}
		if corestr.AnyToString(false, "") != "" {
			t.Error("expected empty for empty string")
		}
	})
}

func Test_C35_CloneSlice(t *testing.T) {
	safeTest(t, "Test_C35_CloneSlice", func() {
		orig := []string{"a", "b"}
		cloned := corestr.CloneSlice(orig)
		if len(cloned) != 2 {
			t.Error("expected 2")
		}
		if len(corestr.CloneSlice(nil)) != 0 {
			t.Error("expected 0")
		}
	})
}
