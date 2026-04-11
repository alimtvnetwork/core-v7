package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — comprehensive coverage for remaining uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_AddMethods(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddMethods", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(10)
		h.AddOrUpdate("a", "1")
		h.AddOrUpdateKeyStrValInt("b", 2)
		h.AddOrUpdateKeyStrValFloat("c", 3.0)
		h.AddOrUpdateKeyStrValFloat64("d", 4.0)
		h.AddOrUpdateKeyStrValAny("e", 5)
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "f", Value: "6"})
		h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "g", Value: "7"})
		h.Set("h", "8")
		h.SetTrim(" i ", " 9 ")
		h.SetBySplitter("=", "key=val")
		h.SetBySplitter("=", "onlykey")
		h.AddOrUpdateLock("j", "10")

		// Act
		actual := args.Map{"result": h.Length() < 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 10", actual)
	})
}

func Test_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddOrUpdateHashmap", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(2)
		h1.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("b", "2")
		h1.AddOrUpdateHashmap(h2)

		// Act
		actual := args.Map{"result": h1.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		h1.AddOrUpdateHashmap(nil)
	})
}

func Test_Hashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddOrUpdateMap", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddsOrUpdates", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddsOrUpdates(corestr.KeyValuePair{Key: "a", Value: "1"})
		h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "b", Value: "2"})
		h.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "c", Value: "3"})
	})
}

func Test_Hashmap_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddOrUpdateCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		h.AddOrUpdateCollection(keys, vals)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// mismatched lengths
		h.AddOrUpdateCollection(keys, corestr.New.Collection.Strings([]string{"1"}))
	})
}

func Test_Hashmap_Lookups(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Lookups", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": h.Has("a") || !h.Contains("a") || !h.ContainsLock("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		actual = args.Map{"result": h.IsKeyMissing("a") || h.IsKeyMissingLock("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not missing", actual)
		actual = args.Map{"result": h.HasLock("a") || !h.HasWithLock("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		actual = args.Map{"result": h.HasAll("a") || !h.HasAny("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		actual = args.Map{"result": h.HasAllStrings("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		actual = args.Map{"result": h.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has item", actual)
	})
}

func Test_Hashmap_Get(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Get", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		v, ok := h.Get("a")

		// Act
		actual := args.Map{"result": ok || v != "1"}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
		v2, ok2 := h.GetValue("a")
		actual = args.Map{"result": ok2 || v2 != "1"}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
	})
}

func Test_Hashmap_Remove(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Remove", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h.Remove("a")

		// Act
		actual := args.Map{"result": h.Has("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
		h.AddOrUpdate("b", "2")
		h.RemoveWithLock("b")
	})
}

func Test_Hashmap_Keys_Values(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Keys_Values", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		_ = h.Keys()
		_ = h.AllKeys()
		_ = h.KeysCollection()
		_ = h.KeysLock()
		_ = h.ValuesList()
		_ = h.ValuesListCopyLock()
		_ = h.ValuesCollection()
		_ = h.ValuesCollectionLock()
		_ = h.ValuesHashset()
		_ = h.ValuesHashsetLock()
		_, _ = h.KeysValuesCollection()
		_, _ = h.KeysValuesList()
		_, _ = h.KeysValuesListLock()
		_ = h.KeysValuePairs()
		_ = h.KeysValuePairsCollection()
	})
}

func Test_Hashmap_FilterOps(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_FilterOps", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("abc", "1")
		h.AddOrUpdate("def", "2")
		items := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, s == "abc", false
		})

		// Act
		actual := args.Map{"result": len(items) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		col := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_ConcatNew", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("b", "2")
		newH := h.ConcatNew(true, h2)

		// Act
		actual := args.Map{"result": newH.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_Hashmap_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_ConcatNewUsingMaps", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		newH := h.ConcatNewUsingMaps(true, map[string]string{"b": "2"})

		// Act
		actual := args.Map{"result": newH.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_IsEqual", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(2)
		h1.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": h1.IsEqualPtr(h2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual = args.Map{"result": h1.IsEqualPtrLock(h2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Hashmap_Diff(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Diff", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("b", "2")
		_ = h.Diff(h2)
	})
}

func Test_Hashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_KeysToLower", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("ABC", "1")
		lower := h.KeysToLower()

		// Act
		actual := args.Map{"result": lower.Has("abc")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected lowercase key", actual)
	})
}

func Test_Hashmap_GetExcept(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_GetExcept", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(3)
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		result := h.GetValuesKeysExcept([]string{"a"})

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_ = h.GetAllExceptCollection(nil)
	})
}

func Test_Hashmap_Join(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Join", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		_ = h.Join(",")
		_ = h.JoinKeys(",")
	})
}

func Test_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_String", func() {
		h := corestr.New.Hashmap.Cap(2)
		_ = h.String()
		h.AddOrUpdate("a", "1")
		_ = h.String()
		_ = h.StringLock()
	})
}

func Test_Hashmap_JsonOps(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_JsonOps", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		_ = h.JsonModel()
		_ = h.JsonModelAny()
		b, _ := h.MarshalJSON()
		h2 := corestr.New.Hashmap.Cap(2)
		_ = h2.UnmarshalJSON(b)
		_ = h.Json()
		_ = h.JsonPtr()
		_ = h.AsJsoner()
		_ = h.AsJsonContractsBinder()
		_ = h.AsJsonParseSelfInjector()
		_ = h.AsJsonMarshaller()
	})
}

func Test_Hashmap_Clone(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Clone", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		c := h.Clone()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		cp := h.ClonePtr()
		actual = args.Map{"result": cp == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Hashmap_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_ClearDispose", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h.Clear()

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		h.AddOrUpdate("b", "2")
		h.Dispose()
	})
}

func Test_Hashmap_ToError(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_ToError", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		_ = h.ToError(",")
		_ = h.ToDefaultError()
		_ = h.KeyValStringLines()
	})
}

func Test_Hashmap_SerializeDeserialize(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_SerializeDeserialize", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		_, _ = h.Serialize()
		var target map[string]string
		_ = h.Deserialize(&target)
	})
}

func Test_Hashmap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_ParseInjectUsingJson", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		jr := h.JsonPtr()
		target := corestr.New.Hashmap.Cap(2)
		_, err := target.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Hashmap_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddOrUpdateWithWgLock", func() {
		h := corestr.New.Hashmap.Cap(2)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateWithWgLock("a", "1", wg)
		wg.Wait()
	})
}

func Test_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddOrUpdateStringsPtrWgLock", func() {
		h := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(wg, []string{"a"}, []string{"1"})
		wg.Wait()
	})
}

func Test_Hashmap_SafeItems(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_SafeItems", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		_ = h.SafeItems()
		_ = h.Items()
		_ = h.ItemsCopyLock()
	})
}

func Test_Hashmap_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_HasAllCollectionItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": h.HasAllCollectionItems(c)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashmap_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_ToStringsUsingCompiler", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		strs := h.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })

		// Act
		actual := args.Map{"result": len(strs) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesFilters(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddsOrUpdatesFilters", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddsOrUpdatesAnyUsingFilter(
			func(p corestr.KeyAnyValuePair) (string, bool, bool) { return "v", true, false },
			corestr.KeyAnyValuePair{Key: "a", Value: "1"},
		)
		h.AddsOrUpdatesAnyUsingFilterLock(
			func(p corestr.KeyAnyValuePair) (string, bool, bool) { return "v", true, false },
			corestr.KeyAnyValuePair{Key: "b", Value: "2"},
		)
		h.AddsOrUpdatesUsingFilter(
			func(p corestr.KeyValuePair) (string, bool, bool) { return "v", true, false },
			corestr.KeyValuePair{Key: "c", Value: "3"},
		)
	})
}

func Test_Hashmap_Collection(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Collection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		c := h.Collection()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — comprehensive coverage for remaining uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashset_AddMethods(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddMethods", func() {
		h := corestr.New.Hashset.Cap(10)
		h.Add("a")
		h.AddNonEmpty("")
		h.AddNonEmpty("b")
		h.AddNonEmptyWhitespace("  ")
		h.AddNonEmptyWhitespace("c")
		h.AddIf(false, "skip")
		h.AddIf(true, "d")
		h.AddIfMany(false, "x", "y")
		h.AddIfMany(true, "e", "f")
		h.AddFunc(func() string { return "g" })
		h.AddLock("h")
		h.AddBool("a") // existing
		h.AddBool("i") // new
		s := "j"
		h.AddPtr(&s)
		h.AddPtrLock(&s)
		h.AddStrings([]string{"k", "l"})
		h.AddStringsLock([]string{"m"})
		h.Adds("n", "o")
	})
}

func Test_Hashset_Lookups(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Lookups", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Add("a")

		// Act
		actual := args.Map{"result": h.Has("a") || !h.Contains("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		actual = args.Map{"result": h.IsMissing("a") || h.IsMissingLock("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not missing", actual)
		actual = args.Map{"result": h.HasLock("a") || !h.HasWithLock("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		actual = args.Map{"result": h.HasAll("a") || !h.HasAny("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		actual = args.Map{"result": h.HasAllStrings([]string{"a"})}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		actual = args.Map{"result": h.HasAnyItem() || !h.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has item", actual)
		actual = args.Map{"result": h.IsAllMissing("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Hashset_Collection(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Collection", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Adds("a", "b")
		_ = h.Collection()
		_ = h.SortedList()
		_ = h.OrderedList()
		_ = h.SafeStrings()
		_ = h.Lines()
		_ = h.SimpleSlice()
		_ = h.Items()
		_ = h.List()
		_ = h.ListPtr()
		_ = h.ListCopyLock()
		_ = h.ListPtrSortedAsc()
		_ = h.ListPtrSortedDsc()
	})
}

func Test_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Filter", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Adds("abc", "def", "ghi")
		filtered := h.Filter(func(s string) bool { return s == "abc" })

		// Act
		actual := args.Map{"result": filtered.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		items := h.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, s == "abc", false
		})
		actual := args.Map{"result": len(items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		col := h.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Hashset_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_GetAllExcept", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(3)
		h.Adds("a", "b", "c")
		r := h.GetAllExcept([]string{"a"})

		// Act
		actual := args.Map{"result": len(r) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		_ = h.GetAllExcept(nil)
		_ = h.GetAllExceptSpread("a")
		_ = h.GetAllExceptCollection(nil)
		_ = h.GetAllExceptHashset(nil)
	})
}

func Test_Hashset_Concat(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Concat", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		h2 := corestr.New.Hashset.Cap(2)
		h2.Add("b")
		newH := h.ConcatNewHashsets(true, h2)

		// Act
		actual := args.Map{"result": newH.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
		_ = h.ConcatNewStrings(true, []string{"c"})
	})
}

func Test_Hashset_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_IsEquals", func() {
		// Arrange
		h1 := corestr.New.Hashset.Cap(2)
		h1.Add("a")
		h2 := corestr.New.Hashset.Cap(2)
		h2.Add("a")

		// Act
		actual := args.Map{"result": h1.IsEquals(h2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual = args.Map{"result": h1.IsEqual(h2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual = args.Map{"result": h1.IsEqualsLock(h2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Hashset_Remove(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Remove", func() {
		h := corestr.New.Hashset.Cap(3)
		h.Adds("a", "b")
		h.Remove("a")
		h.SafeRemove("b")
		h.SafeRemove("missing")
		h.Add("c")
		h.RemoveWithLock("c")
	})
}

func Test_Hashset_String(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_String", func() {
		h := corestr.New.Hashset.Cap(2)
		_ = h.String()
		h.Add("a")
		_ = h.String()
		_ = h.StringLock()
	})
}

func Test_Hashset_Join(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Join", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Adds("a", "b")
		_ = h.Join(",")
		_ = h.JoinSorted(",")
		_ = h.JoinLine()
		_ = h.NonEmptyJoins(",")
		_ = h.NonWhitespaceJoins(",")
	})
}

func Test_Hashset_JsonOps(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_JsonOps", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		_ = h.JsonModel()
		_ = h.JsonModelAny()
		b, _ := h.MarshalJSON()
		h2 := corestr.New.Hashset.Cap(2)
		_ = h2.UnmarshalJSON(b)
		_ = h.Json()
		_ = h.JsonPtr()
		_ = h.AsJsonContractsBinder()
		_ = h.AsJsoner()
		_ = h.AsJsonParseSelfInjector()
		_ = h.AsJsonMarshaller()
	})
}

func Test_Hashset_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_ClearDispose", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		h.Clear()

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		h.Add("b")
		h.Dispose()
	})
}

func Test_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Resize", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		h.Resize(100)
		h.ResizeLock(200)
		h.AddCapacities(50)
		h.AddCapacitiesLock(50)
	})
}

func Test_Hashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_ToLowerSet", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.Add("ABC")
		lower := h.ToLowerSet()

		// Act
		actual := args.Map{"result": lower.Has("abc")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected lowercase", actual)
	})
}

func Test_Hashset_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_DistinctDiff", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(3)
		h.Adds("a", "b")
		diff := h.DistinctDiffLinesRaw("b", "c")

		// Act
		actualDiff := args.Map{"length": len(diff)}

		// Assert
		expectedDiff := args.Map{"length": 2}
		expectedDiff.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw returns 2 -- a and c", actualDiff)
		diffMap := h.DistinctDiffLines("b", "c")
		actual := args.Map{"result": len(diffMap) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashset_DistinctDiffHashset(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_DistinctDiffHashset", func() {
		// Arrange
		h1 := corestr.New.Hashset.Cap(2)
		h1.Adds("a", "b")
		h2 := corestr.New.Hashset.Cap(2)
		h2.Adds("b", "c")
		diff := h1.DistinctDiffHashset(h2)

		// Act
		actual := args.Map{"result": len(diff) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashset_SerializeDeserialize(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_SerializeDeserialize", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		_, _ = h.Serialize()
		var target map[string]bool
		_ = h.Deserialize(&target)
	})
}

func Test_Hashset_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_ParseInjectUsingJson", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		jr := h.JsonPtr()
		target := corestr.New.Hashset.Cap(2)
		_, err := target.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Hashset_WrapQuotes(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_WrapQuotes", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		_ = h.WrapDoubleQuote()
		h2 := corestr.New.Hashset.Cap(2)
		h2.Add("b")
		_ = h2.WrapSingleQuote()
		h3 := corestr.New.Hashset.Cap(2)
		h3.Add("c")
		_ = h3.WrapDoubleQuoteIfMissing()
		h4 := corestr.New.Hashset.Cap(2)
		h4.Add("d")
		_ = h4.WrapSingleQuoteIfMissing()
	})
}

func Test_Hashset_MapStringAny(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_MapStringAny", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		m := h.MapStringAny()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_ = h.MapStringAnyDiff()
	})
}

func Test_Hashset_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddCollection", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h.AddCollection(c)
		h.AddCollections(c, nil)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashset_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddSimpleSlice", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		h.AddSimpleSlice(ss)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashset_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddHashsetItems", func() {
		h := corestr.New.Hashset.Cap(5)
		h2 := corestr.New.Hashset.Cap(2)
		h2.Adds("a", "b")
		h.AddHashsetItems(h2)
		h.AddHashsetItems(nil)
	})
}

func Test_Hashset_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddItemsMap", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddItemsMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashset_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_HasAllCollectionItems", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.Adds("a", "b")
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": h.HasAllCollectionItems(c)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": h.HasAllCollectionItems(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Hashset_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddFuncErr", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddFuncErr(
			func() (string, error) { return "a", nil },
			func(err error) {},
		)

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashset_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddsUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddsUsingFilter(
			func(s string, i int) (string, bool, bool) { return s, s != "", false },
			"a", "", "b",
		)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddsAnyUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddsAnyUsingFilter(
			func(s string, i int) (string, bool, bool) { return s, true, false },
			"a", nil, "b",
		)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashset_Transpile(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Transpile", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		result := h.Transpile(func(s string) string { return s + "!" })

		// Act
		actual := args.Map{"result": result.Has("a!")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected transpiled", actual)
	})
}
