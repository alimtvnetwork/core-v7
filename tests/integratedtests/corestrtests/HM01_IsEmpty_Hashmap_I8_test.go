package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Hashmap — Core operations
// =============================================================================

func Test_HM01_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I8_HM01_IsEmpty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"result": h.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_HM02_HasItems(t *testing.T) {
	safeTest(t, "Test_I8_HM02_HasItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"result": h.HasItems()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		h.AddOrUpdate("k", "v")
		actual = args.Map{"result": h.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HM03_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_I8_HM03_IsEmptyLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"result": h.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_HM04_AddOrUpdate(t *testing.T) {
	safeTest(t, "Test_I8_HM04_AddOrUpdate", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		isNew := h.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"result": isNew}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		isNew2 := h.AddOrUpdate("k", "v2")
		actual = args.Map{"result": isNew2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not new", actual)
	})
}

func Test_HM05_Set(t *testing.T) {
	safeTest(t, "Test_I8_HM05_Set", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		isNew := h.Set("k", "v")

		// Act
		actual := args.Map{"result": isNew}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
	})
}

func Test_HM06_SetTrim(t *testing.T) {
	safeTest(t, "Test_I8_HM06_SetTrim", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.SetTrim(" k ", " v ")

		// Act
		actual := args.Map{"result": h.Has("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected trimmed key", actual)
	})
}

func Test_HM07_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_I8_HM07_SetBySplitter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.SetBySplitter("=", "key=value")

		// Act
		actual := args.Map{"result": h.Has("key")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected key", actual)
		// no splitter
		h.SetBySplitter("=", "nosplit")
		actual = args.Map{"result": h.Has("nosplit")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected nosplit", actual)
	})
}

func Test_HM08_AddOrUpdateKeyStrValInt(t *testing.T) {
	safeTest(t, "Test_I8_HM08_AddOrUpdateKeyStrValInt", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyStrValInt("k", 42)

		// Act
		actual := args.Map{"result": h.Has("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected key", actual)
	})
}

func Test_HM09_AddOrUpdateKeyStrValFloat(t *testing.T) {
	safeTest(t, "Test_I8_HM09_AddOrUpdateKeyStrValFloat", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyStrValFloat("k", 3.14)

		// Act
		actual := args.Map{"result": h.Has("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected key", actual)
	})
}

func Test_HM10_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	safeTest(t, "Test_I8_HM10_AddOrUpdateKeyStrValFloat64", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyStrValFloat64("k", 3.14)

		// Act
		actual := args.Map{"result": h.Has("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected key", actual)
	})
}

func Test_HM11_AddOrUpdateKeyStrValAny(t *testing.T) {
	safeTest(t, "Test_I8_HM11_AddOrUpdateKeyStrValAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyStrValAny("k", "val")

		// Act
		actual := args.Map{"result": h.Has("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected key", actual)
	})
}

func Test_HM12_AddOrUpdateKeyVal(t *testing.T) {
	safeTest(t, "Test_I8_HM12_AddOrUpdateKeyVal", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		isNew := h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v"})

		// Act
		actual := args.Map{"result": isNew}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
	})
}

func Test_HM13_AddOrUpdateKeyValueAny(t *testing.T) {
	safeTest(t, "Test_I8_HM13_AddOrUpdateKeyValueAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k", Value: 42})

		// Act
		actual := args.Map{"result": h.Has("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected key", actual)
	})
}

func Test_HM14_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_I8_HM14_AddOrUpdateLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateLock("k", "v")

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM15_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_I8_HM15_AddOrUpdateWithWgLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateWithWgLock("k", "v", wg)
		wg.Wait()

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM16_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_I8_HM16_AddOrUpdateHashmap", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		other := corestr.New.Hashmap.Cap(2)
		other.AddOrUpdate("a", "1")
		h.AddOrUpdateHashmap(other)

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		h.AddOrUpdateHashmap(nil)
		actual = args.Map{"result": h.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 1", actual)
	})
}

func Test_HM17_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_I8_HM17_AddOrUpdateMap", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateMap(map[string]string{"a": "1", "b": "2"})

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		h.AddOrUpdateMap(nil)
	})
}

func Test_HM18_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_I8_HM18_AddsOrUpdates", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddsOrUpdates(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_HM19_AddOrUpdateKeyAnyValues(t *testing.T) {
	safeTest(t, "Test_I8_HM19_AddOrUpdateKeyAnyValues", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyAnyValues(
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
		)

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM20_AddOrUpdateKeyValues(t *testing.T) {
	safeTest(t, "Test_I8_HM20_AddOrUpdateKeyValues", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyValues(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM21_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_I8_HM21_AddOrUpdateCollection", func() {
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
		// length mismatch
		vals2 := corestr.New.Collection.Strings([]string{"1"})
		h.AddOrUpdateCollection(keys, vals2)
		// nil keys
		h.AddOrUpdateCollection(nil, vals)
	})
}

// =============================================================================
// Hashmap — Query operations
// =============================================================================

func Test_HM22_Has(t *testing.T) {
	safeTest(t, "Test_I8_HM22_Has", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k", "v")

		// Act
		actual := args.Map{"result": h.Has("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": h.Has("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_HM23_Contains(t *testing.T) {
	safeTest(t, "Test_I8_HM23_Contains", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k", "v")

		// Act
		actual := args.Map{"result": h.Contains("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HM24_ContainsLock(t *testing.T) {
	safeTest(t, "Test_I8_HM24_ContainsLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k", "v")

		// Act
		actual := args.Map{"result": h.ContainsLock("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HM25_IsKeyMissing(t *testing.T) {
	safeTest(t, "Test_I8_HM25_IsKeyMissing", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k", "v")

		// Act
		actual := args.Map{"result": h.IsKeyMissing("k")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": h.IsKeyMissing("z")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HM26_IsKeyMissingLock(t *testing.T) {
	safeTest(t, "Test_I8_HM26_IsKeyMissingLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k", "v")

		// Act
		actual := args.Map{"result": h.IsKeyMissingLock("k")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_HM27_HasLock(t *testing.T) {
	safeTest(t, "Test_I8_HM27_HasLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k", "v")

		// Act
		actual := args.Map{"result": h.HasLock("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HM28_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_I8_HM28_HasAllStrings", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Set("b", "2")

		// Act
		actual := args.Map{"result": h.HasAllStrings("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": h.HasAllStrings("a", "z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_HM29_HasAll(t *testing.T) {
	safeTest(t, "Test_I8_HM29_HasAll", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		// Act
		actual := args.Map{"result": h.HasAll("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HM30_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_I8_HM30_HasAnyItem", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"result": h.HasAnyItem()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		h.Set("a", "1")
		actual = args.Map{"result": h.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HM31_HasAny(t *testing.T) {
	safeTest(t, "Test_I8_HM31_HasAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		// Act
		actual := args.Map{"result": h.HasAny("a", "z")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": h.HasAny("x", "y")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_HM32_HasWithLock(t *testing.T) {
	safeTest(t, "Test_I8_HM32_HasWithLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		// Act
		actual := args.Map{"result": h.HasWithLock("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HM33_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_I8_HM33_HasAllCollectionItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Set("b", "2")
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": h.HasAllCollectionItems(c)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": h.HasAllCollectionItems(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

// =============================================================================
// Hashmap — Get, Values, Keys, Diff, Filter
// =============================================================================

func Test_HM34_Items(t *testing.T) {
	safeTest(t, "Test_I8_HM34_Items", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		// Act
		actual := args.Map{"result": len(h.Items()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM35_SafeItems(t *testing.T) {
	safeTest(t, "Test_I8_HM35_SafeItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		// Act
		actual := args.Map{"result": len(h.SafeItems()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM36_ValuesList(t *testing.T) {
	safeTest(t, "Test_I8_HM36_ValuesList", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		// Act
		actual := args.Map{"result": len(h.ValuesList()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM37_ValuesCollection(t *testing.T) {
	safeTest(t, "Test_I8_HM37_ValuesCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		c := h.ValuesCollection()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM38_ValuesHashset(t *testing.T) {
	safeTest(t, "Test_I8_HM38_ValuesHashset", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		hs := h.ValuesHashset()

		// Act
		actual := args.Map{"result": hs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM39_KeysValuesCollection(t *testing.T) {
	safeTest(t, "Test_I8_HM39_KeysValuesCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		keys, values := h.KeysValuesCollection()

		// Act
		actual := args.Map{"result": keys.Length() != 1 || values.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 each", actual)
	})
}

func Test_HM40_KeysValuesList(t *testing.T) {
	safeTest(t, "Test_I8_HM40_KeysValuesList", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		keys, values := h.KeysValuesList()

		// Act
		actual := args.Map{"result": len(keys) != 1 || len(values) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 each", actual)
	})
}

func Test_HM41_KeysValuePairs(t *testing.T) {
	safeTest(t, "Test_I8_HM41_KeysValuePairs", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		pairs := h.KeysValuePairs()

		// Act
		actual := args.Map{"result": len(pairs) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM42_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_I8_HM42_GetKeysFilteredItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("abc", "1")
		h.Set("de", "2")
		result := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 2, false
		})

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM43_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_I8_HM43_GetKeysFilteredCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("abc", "1")
		c := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM44_ConcatNew(t *testing.T) {
	safeTest(t, "Test_I8_HM44_ConcatNew", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		other := corestr.New.Hashmap.Cap(2)
		other.Set("b", "2")
		result := h.ConcatNew(true, other)

		// Act
		actual := args.Map{"result": result.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_HM45_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_I8_HM45_ConcatNew_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		result := h.ConcatNew(true)

		// Act
		actual := args.Map{"result": result.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_HM46_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_I8_HM46_ConcatNewUsingMaps", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		m := map[string]string{"b": "2"}
		result := h.ConcatNewUsingMaps(true, m)

		// Act
		actual := args.Map{"result": result.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_HM47_Diff(t *testing.T) {
	safeTest(t, "Test_I8_HM47_Diff", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Set("b", "2")
		other := corestr.New.Hashmap.Cap(5)
		other.Set("a", "1")
		diff := h.Diff(other)
		_ = diff
	})
}

func Test_HM48_DiffRaw(t *testing.T) {
	safeTest(t, "Test_I8_HM48_DiffRaw", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		rawDiff := h.DiffRaw(map[string]string{"a": "1"})
		_ = rawDiff
	})
}

func Test_HM49_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_HM49_AddsOrUpdatesUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddsOrUpdatesUsingFilter(func(p corestr.KeyValuePair) (string, bool, bool) {
			return p.Value, true, false
		}, corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM50_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_HM50_AddsOrUpdatesAnyUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddsOrUpdatesAnyUsingFilter(func(p corestr.KeyAnyValuePair) (string, bool, bool) {
			return p.ValueString(), true, false
		}, corestr.KeyAnyValuePair{Key: "a", Value: 42})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM51_Json(t *testing.T) {
	safeTest(t, "Test_I8_HM51_Json", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		j := h.Json()

		// Act
		actual := args.Map{"result": j.JsonString() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_HM52_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_I8_HM52_ParseInjectUsingJson", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		jr := h.JsonPtr()
		h2 := corestr.New.Hashmap.Cap(1)
		_, err := h2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_HM53_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_I8_HM53_ParseInjectUsingJson_Error", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(1)
		bad := corejson.NewResult.UsingString(`invalid`)
		_, err := h.ParseInjectUsingJson(bad)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_HM54_String(t *testing.T) {
	safeTest(t, "Test_I8_HM54_String", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		s := h.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_HM55_Clear(t *testing.T) {
	safeTest(t, "Test_I8_HM55_Clear", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Clear()

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_HM56_Dispose(t *testing.T) {
	safeTest(t, "Test_I8_HM56_Dispose", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Dispose()

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_HM57_ItemsCopyLock(t *testing.T) {
	safeTest(t, "Test_I8_HM57_ItemsCopyLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		m := h.ItemsCopyLock()

		// Act
		actual := args.Map{"result": len(*m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM58_Collection(t *testing.T) {
	safeTest(t, "Test_I8_HM58_Collection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		c := h.Collection()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM59_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_I8_HM59_AddOrUpdateStringsPtrWgLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(wg, []string{"a"}, []string{"1"})
		wg.Wait()

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HM60_KeysValuePairsCollection(t *testing.T) {
	safeTest(t, "Test_I8_HM60_KeysValuePairsCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		kvc := h.KeysValuePairsCollection()

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
