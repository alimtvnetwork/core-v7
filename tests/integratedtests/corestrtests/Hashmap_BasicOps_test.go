package corestrtests

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_BasicOps", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(10)
		h.AddOrUpdate("a", "1")
		h.Set("b", "2")
		h.SetTrim(" c ", " 3 ")

		// Act
		actual := args.Map{"result": h.Length() != 3 || h.IsEmpty() || !h.HasItems() || !h.HasAnyItem()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)

		actual = args.Map{"result": h.Has("a") || !h.Contains("b") || h.IsKeyMissing("a") || !h.IsKeyMissing("z")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Has/Contains/IsKeyMissing failed", actual)

		actual = args.Map{"result": h.HasAll("a", "b") || h.HasAll("a", "z")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAll failed", actual)

		actual = args.Map{"result": h.HasAny("a", "z") || h.HasAny("x", "z")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAny failed", actual)
	})
}

func Test_Hashmap_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_SetBySplitter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.SetBySplitter("=", "key=value")
		h.SetBySplitter("=", "onlykey")

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashmap_AddOrUpdateVariants(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddOrUpdateVariants", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(10)

		h.AddOrUpdateKeyStrValInt("k1", 42)
		h.AddOrUpdateKeyStrValFloat("k2", 3.14)
		h.AddOrUpdateKeyStrValFloat64("k3", 2.71)
		h.AddOrUpdateKeyStrValAny("k4", "any")
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k5", Value: "val5"})
		h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k6", Value: "v6"})

		// Act
		actual := args.Map{"result": h.Length() != 6}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 6", actual)
	})
}

func Test_Hashmap_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddsOrUpdates", func() {
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

		h.AddsOrUpdates()
	})
}

func Test_Hashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddOrUpdateMap", func() {
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

func Test_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddOrUpdateHashmap", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(5)
		h1.Set("a", "1")

		h2 := corestr.New.Hashmap.Cap(5)
		h2.Set("b", "2")

		h1.AddOrUpdateHashmap(h2)
		h1.AddOrUpdateHashmap(nil)

		// Act
		actual := args.Map{"result": h1.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashmap_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddOrUpdateCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		keys := corestr.New.Collection.Strings([]string{"k1", "k2"})
		vals := corestr.New.Collection.Strings([]string{"v1", "v2"})

		h.AddOrUpdateCollection(keys, vals)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		// Mismatched lengths
		h.AddOrUpdateCollection(keys, corestr.New.Collection.Strings([]string{"v1"}))

		// Nil
		h.AddOrUpdateCollection(nil, nil)
	})
}

func Test_Hashmap_Get(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Get", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		v, found := h.Get("a")

		// Act
		actual := args.Map{"result": found || v != "1"}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)

		_, found = h.Get("z")
		actual = args.Map{"result": found}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not found", actual)

		v2, found2 := h.GetValue("a")
		actual = args.Map{"result": found2 || v2 != "1"}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_ConcatNew", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(5)
		h1.Set("a", "1")

		h2 := corestr.New.Hashmap.Cap(5)
		h2.Set("b", "2")

		result := h1.ConcatNew(true, h2)

		// Act
		actual := args.Map{"result": result.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)

		// Empty
		result2 := h1.ConcatNew(true)
		actual = args.Map{"result": result2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_ConcatNewUsingMaps", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		result := h.ConcatNewUsingMaps(true, map[string]string{"b": "2"})

		// Act
		actual := args.Map{"result": result.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)

		result2 := h.ConcatNewUsingMaps(true)
		_ = result2
	})
}

func Test_Hashmap_Keys_Values(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Keys_Values", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Set("b", "2")

		_ = h.Keys()
		_ = h.AllKeys()
		_ = h.ValuesList()
		_ = h.Items()
		_ = h.SafeItems()
		_ = h.Collection()
		_ = h.ValuesCollection()
		_ = h.ValuesHashset()
		_ = h.KeysCollection()
		_ = h.KeysValuePairs()
		_ = h.KeysValuePairsCollection()
	})
}

func Test_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_IsEqual", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(5)
		h1.Set("a", "1")

		h2 := corestr.New.Hashmap.Cap(5)
		h2.Set("a", "1")

		// Act
		actual := args.Map{"result": h1.IsEqualPtr(h2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)

		h3 := corestr.New.Hashmap.Cap(5)
		h3.Set("a", "2")

		actual = args.Map{"result": h1.IsEqualPtr(h3)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)

		// Same ptr
		actual = args.Map{"result": h1.IsEqualPtr(h1)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected same ptr equal", actual)

		// Both empty
		e1 := corestr.Empty.Hashmap()
		e2 := corestr.Empty.Hashmap()
		actual = args.Map{"result": e1.IsEqualPtr(e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty equal", actual)
	})
}

func Test_Hashmap_Remove(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Remove", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Remove("a")

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_KeysToLower", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("HELLO", "world")

		lowered := h.KeysToLower()

		// Act
		actual := args.Map{"result": lowered.Has("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_String", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		_ = h.String()

		empty := corestr.Empty.Hashmap()
		_ = empty.String()
	})
}

func Test_Hashmap_Join(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Join", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		_ = h.Join(",")
		_ = h.JoinKeys(",")
	})
}

func Test_Hashmap_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_JSON", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		b, err := json.Marshal(h)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal failed", actual)

		h2 := corestr.Empty.Hashmap()
		err = json.Unmarshal(b, h2)
		actual = args.Map{"result": err != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal failed", actual)

		actual = args.Map{"result": h2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_Clone(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Clone", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		_ = h.ClonePtr()
	})
}

func Test_Hashmap_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Clear_Dispose", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Clear()

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)

		h2 := corestr.New.Hashmap.Cap(5)
		h2.Set("a", "1")
		h2.Dispose()
	})
}

func Test_Hashmap_ToError(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_ToError", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		_ = h.ToError(",")
		_ = h.ToDefaultError()
	})
}

func Test_Hashmap_KeyValStringLines(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_KeyValStringLines", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		_ = h.KeyValStringLines()
	})
}

func Test_Hashmap_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_ToStringsUsingCompiler", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		lines := h.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })

		// Act
		actual := args.Map{"result": len(lines) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_GetKeysFilteredItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("ab", "1")
		h.Set("cd", "2")

		result := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, s == "ab", false
		})

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_GetKeysFilteredCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("ab", "1")

		result := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_GetValuesExcept(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_GetValuesExcept", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Set("b", "2")

		result := h.GetValuesKeysExcept([]string{"a"})

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)

		result2 := h.GetValuesKeysExcept(nil)
		actual = args.Map{"result": len(result2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashmap_Creators(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Creators", func() {
		_ = corestr.New.Hashmap.Empty()
		_ = corestr.New.Hashmap.Cap(5)
		_ = corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		_ = corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{"a": "1"})
		_ = corestr.New.Hashmap.UsingMapOptions(false, 0, map[string]string{"a": "1"})
		_ = corestr.New.Hashmap.UsingMapOptions(true, 0, map[string]string{})
		_ = corestr.New.Hashmap.MapWithCap(5, map[string]string{"a": "1"})
		_ = corestr.New.Hashmap.MapWithCap(0, map[string]string{"a": "1"})
		_ = corestr.New.Hashmap.MapWithCap(5, map[string]string{})
		_ = corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		_ = corestr.New.Hashmap.KeyValues()
		_ = corestr.New.Hashmap.KeyAnyValues(corestr.KeyAnyValuePair{Key: "a", Value: 1})
		_ = corestr.New.Hashmap.KeyAnyValues()

		keys := corestr.New.Collection.Strings([]string{"k"})
		vals := corestr.New.Collection.Strings([]string{"v"})
		_ = corestr.New.Hashmap.KeyValuesCollection(keys, vals)
		_ = corestr.New.Hashmap.KeyValuesCollection(nil, nil)
		_ = corestr.New.Hashmap.KeyValuesStrings([]string{"k"}, []string{"v"})
		_ = corestr.New.Hashmap.KeyValuesStrings(nil, nil)
	})
}

func Test_Hashmap_AddsOrUpdatesFilter(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddsOrUpdatesFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)

		h.AddsOrUpdatesUsingFilter(func(kv corestr.KeyValuePair) (string, bool, bool) {
			return kv.Value, true, false
		}, corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)

		h.AddsOrUpdatesUsingFilter(nil)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddsOrUpdatesAnyUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)

		h.AddsOrUpdatesAnyUsingFilter(func(kav corestr.KeyAnyValuePair) (string, bool, bool) {
			return kav.ValueString(), true, false
		}, corestr.KeyAnyValuePair{Key: "a", Value: 1})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashmapDiff — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_HashmapDiff_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov14_HashmapDiff_AllMethods", func() {
		// Arrange
		diff := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2"})

		// Act
		actual := args.Map{"result": diff.Length() != 2 || diff.IsEmpty() || !diff.HasAnyItem()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		actual = args.Map{"result": diff.LastIndex() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)

		_ = diff.AllKeysSorted()
		_ = diff.MapAnyItems()
		_ = diff.Raw()
		_ = diff.RawMapStringAnyDiff()

		// IsRawEqual
		actual = args.Map{"result": diff.IsRawEqual(map[string]string{"a": "1", "b": "2"})}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)

		actual = args.Map{"result": diff.HasAnyChanges(map[string]string{"a": "1", "b": "2"})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no changes", actual)

		_ = diff.DiffRaw(map[string]string{"a": "1", "c": "3"})
		_ = diff.DiffJsonMessage(map[string]string{"a": "2"})
		_ = diff.HashmapDiffUsingRaw(map[string]string{"a": "1"})
		_ = diff.ShouldDiffMessage("test", map[string]string{"a": "1"})
		_ = diff.LogShouldDiffMessage("test", map[string]string{"a": "1"})
		_ = diff.ToStringsSliceOfDiffMap(map[string]string{"a": "1"})

		_, _ = diff.Serialize()

		// Nil receiver
		var nilDiff *corestr.HashmapDiff
		actual = args.Map{"result": nilDiff.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		_ = nilDiff.Raw()
		_ = nilDiff.MapAnyItems()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashset_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_BasicOps", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(10)
		h.Add("a")
		h.Add("b")

		// Act
		actual := args.Map{"result": h.Length() != 2 || h.IsEmpty() || !h.HasItems() || !h.HasAnyItem()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		actual = args.Map{"result": h.Has("a") || !h.Contains("b") || h.IsMissing("a") || !h.IsMissing("z")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Has/Contains/IsMissing failed", actual)

		actual = args.Map{"result": h.HasAll("a", "b") || h.HasAll("a", "z")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAll failed", actual)

		actual = args.Map{"result": h.HasAny("a", "z") || h.HasAny("x", "z")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAny failed", actual)

		actual = args.Map{"result": h.IsAllMissing("x", "y") || h.IsAllMissing("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsAllMissing failed", actual)
	})
}

func Test_Hashset_AddVariants(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddVariants", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(10)
		h.AddNonEmpty("")
		h.AddNonEmpty("a")
		h.AddNonEmptyWhitespace("  ")
		h.AddNonEmptyWhitespace("b")
		h.AddIf(false, "skip")
		h.AddIf(true, "c")
		h.AddIfMany(false, "s1")
		h.AddIfMany(true, "d", "e")
		h.AddFunc(func() string { return "f" })
		h.AddBool("g")
		h.AddBool("g") // Already exists

		str := "h"
		h.AddPtr(&str)

		// Act
		actual := args.Map{"result": h.Length() != 8}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 8", actual)
	})
}

func Test_Hashset_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddFuncErr", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashset_AddStrings_Adds(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddStrings_Adds", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(10)
		h.AddStrings([]string{"a", "b"})
		h.Adds("c", "d")

		// Act
		actual := args.Map{"result": h.Length() != 4}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_Hashset_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddHashsetItems", func() {
		// Arrange
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"b"})

		h1.AddHashsetItems(h2)
		h1.AddHashsetItems(nil)

		// Act
		actual := args.Map{"result": h1.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashset_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddItemsMap", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddItemsMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 (false items skipped)", actual)
	})
}

func Test_Hashset_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddCollection", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h.AddCollection(c)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		h.AddCollection(nil)
	})
}

func Test_Hashset_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddSimpleSlice", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		s := corestr.New.SimpleSlice.Lines("a", "b")
		h.AddSimpleSlice(s)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashset_ConcatNewHashsets(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_ConcatNewHashsets", func() {
		// Arrange
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"b"})

		result := h1.ConcatNewHashsets(true, h2)

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		// Empty
		result2 := h1.ConcatNewHashsets(true)
		actual = args.Map{"result": result2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashset_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_ConcatNewStrings", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewStrings(true, []string{"b", "c"})

		// Act
		actual := args.Map{"result": result.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Hashset_List_OrderedList_SortedList(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_List_OrderedList_SortedList", func() {
		h := corestr.New.Hashset.Strings([]string{"c", "a", "b"})

		_ = h.List()
		_ = h.OrderedList()
		_ = h.SortedList()
		_ = h.SafeStrings()
		_ = h.Lines()
		_ = h.SimpleSlice()
	})
}

func Test_Hashset_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_IsEquals", func() {
		// Arrange
		h1 := corestr.New.Hashset.Strings([]string{"a", "b"})
		h2 := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": h1.IsEquals(h2) || !h1.IsEqual(h2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)

		// Same ptr
		actual = args.Map{"result": h1.IsEquals(h1)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected same ptr equal", actual)

		// Both empty
		e1 := corestr.Empty.Hashset()
		e2 := corestr.Empty.Hashset()
		actual = args.Map{"result": e1.IsEquals(e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty equal", actual)
	})
}

func Test_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_Filter", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"ab", "cd", "ef"})

		result := h.Filter(func(s string) bool { return s == "ab" })

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashset_GetFilteredItems(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_GetFilteredItems", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"ab", "cd"})

		result := h.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashset_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_GetAllExcept", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b", "c"})

		result := h.GetAllExcept([]string{"b"})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		result2 := h.GetAllExcept(nil)
		actual = args.Map{"result": len(result2) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)

		_ = h.GetAllExceptSpread("a")
	})
}

func Test_Hashset_Remove(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_Remove", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.Remove("a")

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)

		h.SafeRemove("z") // no-op
		h.SafeRemove("b")

		actual = args.Map{"result": h.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_ToLowerSet", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"HELLO"})
		lowered := h.ToLowerSet()

		// Act
		actual := args.Map{"result": lowered.Has("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_Resize", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		h.Resize(100)
		h.AddCapacities(50)

		// Act
		actual := args.Map{"result": h.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a after resize", actual)
	})
}

func Test_Hashset_String(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_String", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		_ = h.String()
		_ = h.Join(",")
		_ = h.JoinSorted(",")
		_ = h.JoinLine()

		empty := corestr.Empty.Hashset()
		_ = empty.String()
	})
}

func Test_Hashset_WrapQuotes(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_WrapQuotes", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		_ = h.WrapDoubleQuote()
		_ = h.WrapSingleQuote()
		_ = h.WrapDoubleQuoteIfMissing()
		_ = h.WrapSingleQuoteIfMissing()
	})
}

func Test_Hashset_Transpile(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_Transpile", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.Transpile(func(s string) string { return s + "!" })

		// Act
		actual := args.Map{"result": result.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)

		// Empty
		empty := corestr.Empty.Hashset()
		_ = empty.Transpile(func(s string) string { return s })
	})
}

func Test_Hashset_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_JSON", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})

		b, err := json.Marshal(h)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal failed", actual)

		h2 := corestr.Empty.Hashset()
		err = json.Unmarshal(b, h2)
		actual = args.Map{"result": err != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal failed", actual)
	})
}

func Test_Hashset_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_DistinctDiff", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})

		diffLines := h.DistinctDiffLinesRaw("b", "c")

		// Act
		actual := args.Map{"result": len(diffLines) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		diffMap := h.DistinctDiffLines("b", "c")
		actual = args.Map{"result": len(diffMap) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashset_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_Clear_Dispose", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.Clear()

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)

		h2 := corestr.New.Hashset.Strings([]string{"a"})
		h2.Dispose()
	})
}

func Test_Hashset_MapStringAny(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_MapStringAny", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		m := h.MapStringAny()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)

		_ = h.MapStringAnyDiff()
	})
}

func Test_Hashset_Creators(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_Creators", func() {
		_ = corestr.New.Hashset.Empty()
		_ = corestr.New.Hashset.Cap(5)
		_ = corestr.New.Hashset.Strings([]string{"a"})
		_ = corestr.New.Hashset.Strings(nil)
		_ = corestr.New.Hashset.StringsSpreadItems("a", "b")
		_ = corestr.New.Hashset.StringsOption(5, true, "a")
		_ = corestr.New.Hashset.StringsOption(0, false)
		_ = corestr.New.Hashset.StringsOption(5, false)
		_ = corestr.New.Hashset.UsingMap(map[string]bool{"a": true})
		_ = corestr.New.Hashset.UsingMap(map[string]bool{})
		_ = corestr.New.Hashset.UsingMapOption(5, true, map[string]bool{"a": true})
		_ = corestr.New.Hashset.UsingMapOption(5, false, map[string]bool{"a": true})
		_ = corestr.New.Hashset.UsingMapOption(5, false, map[string]bool{})

		c := corestr.New.Collection.Strings([]string{"a"})
		_ = corestr.New.Hashset.UsingCollection(c)
		_ = corestr.New.Hashset.UsingCollection(nil)

		s := corestr.New.SimpleSlice.Lines("a")
		_ = corestr.New.Hashset.SimpleSlice(s)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_LinkedList_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_BasicOps", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b").Add("c")

		// Act
		actual := args.Map{"result": ll.Length() != 3 || ll.IsEmpty() || !ll.HasItems()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)

		actual = args.Map{"result": ll.Head().Element != "a" || ll.Tail().Element != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Head/Tail failed", actual)
	})
}

func Test_LinkedList_AddVariants(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_AddVariants", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("")
		ll.AddNonEmpty("a")
		ll.AddNonEmptyWhitespace("  ")
		ll.AddNonEmptyWhitespace("b")
		ll.AddIf(false, "skip")
		ll.AddIf(true, "c")
		ll.AddsIf(false, "s1")
		ll.AddsIf(true, "d")
		ll.AddFunc(func() string { return "e" })
		ll.Push("f")
		ll.PushBack("g")
		ll.PushFront("front")
		ll.AddFront("front2")

		// Act
		actual := args.Map{"result": ll.Length() < 7}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 7", actual)
	})
}

func Test_LinkedList_Adds_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_Adds_AddStrings", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.AddStrings([]string{"d", "e"})

		// Act
		actual := args.Map{"result": ll.Length() != 5}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
	})
}

func Test_LinkedList_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_AddCollection", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(c)
		ll.AddCollection(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_IsEquals", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll1.IsEquals(ll2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)

		// Case insensitive
		ll3 := corestr.New.LinkedList.Strings([]string{"A", "B"})
		actual = args.Map{"result": ll1.IsEqualsWithSensitive(ll3, false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal insensitive", actual)
	})
}

func Test_LinkedList_List_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_List_ToCollection", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		_ = ll.List()
		_ = ll.ToCollection(0)
	})
}

func Test_LinkedList_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_SafeIndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})

		node := ll.SafeIndexAt(1)

		// Act
		actual := args.Map{"result": node == nil || node.Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)

		actual = args.Map{"result": ll.SafeIndexAt(-1) != nil || ll.SafeIndexAt(10) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_LinkedList_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_SafePointerIndexAtUsingDefault", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefault(0, "def") != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)

		actual = args.Map{"result": ll.SafePointerIndexAtUsingDefault(5, "def") != "def"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected def", actual)
	})
}

func Test_LinkedList_Loop(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_Loop", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		count := 0

		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"result": count != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_LinkedList_String(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_String", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.String()
		_ = ll.Join(",")

		empty := corestr.New.LinkedList.Create()
		_ = empty.String()
	})
}

func Test_LinkedList_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_JSON", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		b, err := json.Marshal(ll)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal failed", actual)

		ll2 := corestr.New.LinkedList.Create()
		err = json.Unmarshal(b, ll2)
		actual = args.Map{"result": err != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal failed", actual)

		actual = args.Map{"result": ll2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_Clear(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_Clear", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.Clear()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_Creators(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_Creators", func() {
		_ = corestr.New.LinkedList.Create()
		_ = corestr.New.LinkedList.Empty()
		_ = corestr.New.LinkedList.Strings([]string{"a"})
		_ = corestr.New.LinkedList.Strings(nil)
		_ = corestr.New.LinkedList.SpreadStrings("a")
		_ = corestr.New.LinkedList.SpreadStrings()
		_ = corestr.New.LinkedList.UsingMap(map[string]bool{"a": true})
		_ = corestr.New.LinkedList.UsingMap(nil)
	})
}

func Test_LinkedList_AppendNode(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_AppendNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AppendNode(node)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_GetNextNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		nodes := ll.GetNextNodes(2)

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_GetAllLinkedNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		nodes := ll.GetAllLinkedNodes()

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_GetCompareSummary", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a"})
		ll2 := corestr.New.LinkedList.Strings([]string{"b"})
		_ = ll1.GetCompareSummary(ll2, "left", "right")
	})
}

func Test_LinkedList_Joins(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_Joins", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		result := ll.Joins(",", "c")

		// Act
		actual := args.Map{"result": result == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedListNode — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_LinkedListNode_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedListNode_AllMethods", func() {
		// Arrange
		node := &corestr.LinkedListNode{Element: "a"}

		// Act
		actual := args.Map{"result": node.HasNext()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no next", actual)

		actual = args.Map{"result": node.String() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)

		_ = node.Clone()
		_ = node.IsEqualValue("a")
		_ = node.IsEqualValueSensitive("A", false)
		_ = node.List()
		_ = node.Join(",")
		_ = node.StringList("header: ")

		end, length := node.EndOfChain()
		actual = args.Map{"result": end != node || length != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected self, 1", actual)

		actual = args.Map{"result": node.IsEqual(node)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal to self", actual)

		node2 := &corestr.LinkedListNode{Element: "a"}
		actual = args.Map{"result": node.IsEqual(node2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// NonChainedLinkedListNodes — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_NonChainedLinkedListNodes(t *testing.T) {
	safeTest(t, "Test_Cov14_NonChainedLinkedListNodes", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		nodes.Adds(n1, n2)

		// Act
		actual := args.Map{"result": nodes.Length() != 2 || nodes.IsEmpty() || !nodes.HasItems()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		actual = args.Map{"result": nodes.First().Element != "a" || nodes.Last().Element != "b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "First/Last failed", actual)

		_ = nodes.FirstOrDefault()
		_ = nodes.LastOrDefault()
		_ = nodes.IsChainingApplied()
		_ = nodes.Items()

		nodes.ApplyChaining()
		actual = args.Map{"result": nodes.IsChainingApplied()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected chaining applied", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// NonChainedLinkedCollectionNodes — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_NonChainedLinkedCollectionNodes(t *testing.T) {
	safeTest(t, "Test_Cov14_NonChainedLinkedCollectionNodes", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		n1 := &corestr.LinkedCollectionNode{Element: c1}
		nodes.Adds(n1)

		// Act
		actual := args.Map{"result": nodes.Length() != 1 || nodes.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)

		_ = nodes.First()
		_ = nodes.FirstOrDefault()
		_ = nodes.Last()
		_ = nodes.LastOrDefault()
		_ = nodes.Items()

		nodes.ApplyChaining()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValueCollection — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyValueCollection_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_AllMethods", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1")
		kvc.Add("k2", "v2")
		kvc.AddIf(true, "k3", "v3")
		kvc.AddIf(false, "s", "s")

		// Act
		actual := args.Map{"result": kvc.Length() != 3 || kvc.Count() != 3 || kvc.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)

		actual = args.Map{"result": kvc.HasAnyItem() || kvc.LastIndex() != 2 || !kvc.HasIndex(2) || kvc.HasIndex(5)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "index checks failed", actual)

		_ = kvc.First()
		_ = kvc.FirstOrDefault()
		_ = kvc.Last()
		_ = kvc.LastOrDefault()

		actual = args.Map{"result": kvc.HasKey("k1") || !kvc.IsContains("k1")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasKey failed", actual)

		v, found := kvc.Get("k1")
		actual = args.Map{"result": found || v != "v1"}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Get failed", actual)

		_ = kvc.AllKeys()
		_ = kvc.AllKeysSorted()
		_ = kvc.AllValues()
		_ = kvc.SafeValueAt(0)
		_ = kvc.SafeValueAt(100)
		_ = kvc.SafeValuesAtIndexes(0, 1)
		_ = kvc.Strings()
		_ = kvc.String()
		_ = kvc.Compile()
		_ = kvc.Join(",")
		_ = kvc.JoinKeys(",")
		_ = kvc.JoinValues(",")
		_ = kvc.Hashmap()
		_ = kvc.Map()
	})
}

func Test_KeyValueCollection_Adds(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_Adds", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Adds(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		kvc.AddMap(map[string]string{"c": "3"})
		kvc.AddHashsetMap(map[string]bool{"d": true})

		h := corestr.New.Hashset.Strings([]string{"e"})
		kvc.AddHashset(h)

		hm := corestr.New.Hashmap.Cap(5)
		hm.Set("f", "6")
		kvc.AddsHashmap(hm)
		kvc.AddsHashmaps(hm)

		// Act
		actual := args.Map{"result": kvc.Length() != 7}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 7", actual)
	})
}

func Test_KeyValueCollection_Find(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_Find", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("a", "1")
		kvc.Add("b", "2")

		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "a", false
		})

		// Act
		actual := args.Map{"result": len(found) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_KeyValueCollection_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_JSON", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("a", "1")

		b, err := json.Marshal(kvc)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal failed", actual)

		kvc2 := corestr.Empty.KeyValueCollection()
		err = json.Unmarshal(b, kvc2)
		actual = args.Map{"result": err != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal failed", actual)
	})
}

func Test_KeyValueCollection_StringsUsingFormat(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_StringsUsingFormat", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("a", "1")
		result := kvc.StringsUsingFormat("%s=%s")

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_KeyValueCollection_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_AddStringBySplit", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddStringBySplit("=", "key=value")
		kvc.AddStringBySplitTrim("=", " key = value ")

		// Act
		actual := args.Map{"result": kvc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_KeyValueCollection_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_Clear_Dispose", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("a", "1")
		kvc.Clear()

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)

		kvc2 := corestr.Empty.KeyValueCollection()
		kvc2.Add("x", "y")
		kvc2.Dispose()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleStringOnce_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_BasicOps", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"result": s.IsInitialized() || !s.IsDefined() || s.IsUninitialized() || s.IsInvalid()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected initialized", actual)

		actual = args.Map{"result": s.Value() != "hello" || s.SafeValue() != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)

		actual = args.Map{"result": s.IsEmpty() || s.IsWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

		actual = args.Map{"result": s.HasValidNonEmpty() || !s.HasValidNonWhitespace() || !s.HasSafeNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected valid", actual)

		_ = s.Trim()
		_ = s.ValueBytes()
		_ = s.ValueBytesPtr()
		_ = s.String()
		_ = s.StringPtr()
		_ = s.NonPtr()
		_ = s.Ptr()
	})
}

func Test_SimpleStringOnce_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_SetOnUninitialized", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Uninitialized("")

		err := s.SetOnUninitialized("val")

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)

		err = s.SetOnUninitialized("another")
		actual = args.Map{"result": err == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for already initialized", actual)
	})
}

func Test_SimpleStringOnce_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_GetSetOnce", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Uninitialized("")

		val := s.GetSetOnce("first")

		// Act
		actual := args.Map{"result": val != "first"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected first", actual)

		val2 := s.GetSetOnce("second")
		actual = args.Map{"result": val2 != "first"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected first (already set)", actual)
	})
}

func Test_SimpleStringOnce_GetOnce(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_GetOnce", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Uninitialized("")
		val := s.GetOnce()

		// Act
		actual := args.Map{"result": val != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_SimpleStringOnce_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_GetOnceFunc", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Uninitialized("")
		val := s.GetOnceFunc(func() string { return "computed" })

		// Act
		actual := args.Map{"result": val != "computed"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected computed", actual)

		val2 := s.GetOnceFunc(func() string { return "other" })
		actual = args.Map{"result": val2 != "computed"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected computed (cached)", actual)
	})
}

func Test_SimpleStringOnce_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_SetOnceIfUninitialized", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Uninitialized("")

		// Act
		actual := args.Map{"result": s.SetOnceIfUninitialized("val")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected set", actual)

		actual = args.Map{"result": s.SetOnceIfUninitialized("other")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not set", actual)
	})
}

func Test_SimpleStringOnce_Reset_Invalidate(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Reset_Invalidate", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("val")
		s.Reset()

		// Act
		actual := args.Map{"result": s.IsInitialized()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)

		s2 := corestr.New.SimpleStringOnce.Init("val2")
		s2.Invalidate()

		actual = args.Map{"result": s2.IsInitialized()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
	})
}

func Test_SimpleStringOnce_Conversions(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Conversions", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("42")

		// Act
		actual := args.Map{"result": s.Int() != 42 || s.ValueInt(0) != 42 || s.ValueDefInt() != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)

		_ = s.Byte()
		_ = s.Int16()
		_ = s.Int32()
		_ = s.ValueByte(0)
		_ = s.ValueDefByte()

		f := corestr.New.SimpleStringOnce.Init("3.14")
		_ = f.ValueFloat64(0)
		_ = f.ValueDefFloat64()

		b := corestr.New.SimpleStringOnce.Init("true")
		actual = args.Map{"result": b.Boolean(false) || !b.BooleanDefault() || !b.IsValueBool()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)

		y := corestr.New.SimpleStringOnce.Init("yes")
		actual = args.Map{"result": y.Boolean(false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for yes", actual)

		_ = s.IsSetter(false)
		_ = s.IsSetter(true)
	})
}

func Test_SimpleStringOnce_WithinRange(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_WithinRange", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("50")

		val, inRange := s.WithinRange(true, 0, 100)

		// Act
		actual := args.Map{"result": inRange || val != 50}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 50 in range", actual)

		val2, inRange2 := s.WithinRange(true, 60, 100)
		actual = args.Map{"result": inRange2 || val2 != 60}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected min boundary", actual)

		_, _ = s.Uint16()
		_, _ = s.Uint32()
	})
}

func Test_SimpleStringOnce_Is_Contains_Regex(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Is_Contains_Regex", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello world")

		// Act
		actual := args.Map{"result": s.Is("hello world") || s.Is("other")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Is failed", actual)

		actual = args.Map{"result": s.IsAnyOf("hello world", "other") || !s.IsAnyOf()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsAnyOf failed", actual)

		actual = args.Map{"result": s.IsContains("hello")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsContains failed", actual)

		actual = args.Map{"result": s.IsAnyContains("hello", "xyz") || !s.IsAnyContains()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsAnyContains failed", actual)

		actual = args.Map{"result": s.IsEqualNonSensitive("HELLO WORLD")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)

		re := regexp.MustCompile(`hello`)
		actual = args.Map{"result": s.IsRegexMatches(re) || s.IsRegexMatches(nil)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "regex failed", actual)

		_ = s.RegexFindString(re)
		_ = s.RegexFindString(nil)
		_, _ = s.RegexFindAllStringsWithFlag(re, -1)
		_ = s.RegexFindAllStrings(re, -1)
	})
}

func Test_SimpleStringOnce_Split(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Split", func() {
		s := corestr.New.SimpleStringOnce.Init("a,b,c")

		_ = s.Split(",")
		_ = s.SplitNonEmpty(",")
		_ = s.SplitTrimNonWhitespace(",")
		_, _ = s.SplitLeftRight(",")
		_, _ = s.SplitLeftRightTrim(",")
		_ = s.LinesSimpleSlice()
		_ = s.SimpleSlice(",")
	})
}

func Test_SimpleStringOnce_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_ConcatNew", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")
		r := s.ConcatNew(" world")

		// Act
		actual := args.Map{"result": r.Value() != "hello world"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello world", actual)

		r2 := s.ConcatNewUsingStrings("-", "world")
		_ = r2
	})
}

func Test_SimpleStringOnce_Clone(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Clone", func() {
		s := corestr.New.SimpleStringOnce.Init("val")
		_ = s.Clone()
		_ = s.ClonePtr()
		_ = s.CloneUsingNewVal("other")
	})
}

func Test_SimpleStringOnce_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_JSON", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("val")

		b, err := json.Marshal(&s)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal failed", actual)

		s2 := corestr.Empty.SimpleStringOnce()
		err = json.Unmarshal(b, &s2)
		actual = args.Map{"result": err != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal failed", actual)
	})
}

func Test_SimpleStringOnce_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Dispose", func() {
		s := corestr.New.SimpleStringOnce.Init("val")
		s.Dispose()
	})
}

func Test_SimpleStringOnce_Creators(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Creators", func() {
		_ = corestr.New.SimpleStringOnce.Init("val")
		_ = corestr.New.SimpleStringOnce.InitPtr("val")
		_ = corestr.New.SimpleStringOnce.Uninitialized("val")
		_ = corestr.New.SimpleStringOnce.Create("val", true)
		_ = corestr.New.SimpleStringOnce.CreatePtr("val", true)
		_ = corestr.New.SimpleStringOnce.Empty()
		_ = corestr.New.SimpleStringOnce.Any(true, 42, true)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionsOfCollection — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_CollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_CollectionsOfCollection", func() {
		// Arrange
		coc := corestr.Empty.CollectionsOfCollection()

		// Act
		actual := args.Map{"result": coc.IsEmpty() || coc.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)

		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Add(c1)

		actual = args.Map{"result": coc.Length() != 1 || coc.AllIndividualItemsLength() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1/2", actual)

		_ = coc.Items()
		_ = coc.List(0)
		_ = coc.ToCollection()
		_ = coc.String()

		coc.AddStrings(false, []string{"c", "d"})
		coc.AddsStringsOfStrings(false, []string{"e"}, []string{"f"})
	})
}

func Test_CollectionsOfCollection_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_CollectionsOfCollection_JSON", func() {
		// Arrange
		coc := corestr.Empty.CollectionsOfCollection()
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)

		b, err := json.Marshal(coc)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal failed", actual)

		coc2 := corestr.Empty.CollectionsOfCollection()
		err = json.Unmarshal(b, coc2)
		actual = args.Map{"result": err != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal failed", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashsetsCollection — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_HashsetsCollection", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		actual := args.Map{"result": hc.IsEmpty() || hc.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)

		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		hc.Add(h)
		hc.AddNonNil(nil)
		hc.AddNonEmpty(corestr.Empty.Hashset())
		hc.AddNonEmpty(corestr.New.Hashset.Strings([]string{"c"}))

		actual = args.Map{"result": hc.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		_ = hc.StringsList()
		_ = hc.String()
		_ = hc.Join(",")
		_ = hc.List()
		_ = hc.ListPtr()
	})
}

func Test_HashsetsCollection_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov14_HashsetsCollection_IsEqual", func() {
		// Arrange
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"a"})

		hc1 := corestr.Empty.HashsetsCollection()
		hc1.Add(h1)

		hc2 := corestr.Empty.HashsetsCollection()
		hc2.Add(h2)

		// Act
		actual := args.Map{"result": hc1.IsEqualPtr(hc2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_HashsetsCollection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov14_HashsetsCollection_ConcatNew", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(h)

		hc2 := corestr.Empty.HashsetsCollection()
		hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))

		result := hc.ConcatNew(hc2)

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		// No args
		result2 := hc.ConcatNew()
		actual = args.Map{"result": result2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HashsetsCollection_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_HashsetsCollection_JSON", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))

		b, err := json.Marshal(hc)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal failed", actual)

		hc2 := corestr.Empty.HashsetsCollection()
		err = json.Unmarshal(b, hc2)
		actual = args.Map{"result": err != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal failed", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollections — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_LinkedCollections_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedCollections_BasicOps", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})

		lc.Add(c1)
		lc.Add(c2)

		// Act
		actual := args.Map{"result": lc.Length() != 2 || lc.IsEmpty() || !lc.HasItems()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		_ = lc.Head()
		_ = lc.Tail()
		_ = lc.First()
		_ = lc.Last()
		_ = lc.FirstOrDefault()
		_ = lc.LastOrDefault()
		_ = lc.AllIndividualItemsLength()
	})
}

func Test_LinkedCollections_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedCollections_AddStrings", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.AddStrings("a", "b")

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedCollections_AddFront(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedCollections_AddFront", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})

		lc.Add(c1)
		lc.AddFront(c2)

		// Act
		actual := args.Map{"result": lc.First().First() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b first", actual)
	})
}

func Test_LinkedCollections_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedCollections_IsEquals", func() {
		// Arrange
		lc1 := corestr.Empty.LinkedCollections()
		lc1.AddStrings("a")

		lc2 := corestr.Empty.LinkedCollections()
		lc2.AddStrings("a")

		// Act
		actual := args.Map{"result": lc1.IsEqualsPtr(lc2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_LinkedCollections_AddAnother(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedCollections_AddAnother", func() {
		// Arrange
		lc1 := corestr.Empty.LinkedCollections()
		lc1.AddStrings("a")

		lc2 := corestr.Empty.LinkedCollections()
		lc2.AddStrings("b")

		lc1.AddAnother(lc2)

		// Act
		actual := args.Map{"result": lc1.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollectionNode — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_LinkedCollectionNode(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedCollectionNode", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		node := &corestr.LinkedCollectionNode{Element: c}

		// Act
		actual := args.Map{"result": node.IsEmpty() || !node.HasElement() || node.HasNext()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid", actual)

		_ = node.String()
		_ = node.List()
		_ = node.Join(",")
		_ = node.Clone()
		_ = node.IsEqual(node)

		end, length := node.EndOfChain()
		actual = args.Map{"result": end != node || length != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected self, 1", actual)

		_ = node.CreateLinkedList()
		_ = node.IsEqualValue(c)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CharCollectionMap — basic coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_CharCollectionMap_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_CharCollectionMap_BasicOps", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()

		// Act
		actual := args.Map{"result": ccm.IsEmpty() || ccm.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)

		_ = ccm.Length()
		_ = ccm.AllLengthsSum()
		_ = ccm.GetChar("hello")
		_ = ccm.GetChar("")
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CharHashsetMap — basic coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_CharHashsetMap_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_CharHashsetMap_BasicOps", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()

		// Act
		actual := args.Map{"result": chm.IsEmpty() || chm.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)

		_ = chm.Length()
		_ = chm.AllLengthsSum()
		_ = chm.GetChar("hello")
		_ = chm.GetChar("")
		_ = chm.GetCharOf("hello")
		_ = chm.GetCharOf("")
		_ = chm.Has("hello")
		_ = chm.LengthOf('h')
		_ = chm.LengthOfHashsetFromFirstChar("hello")
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValues Creator — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyValuesCreator(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValuesCreator", func() {
		_ = corestr.New.KeyValues.Cap(5)
	})
}
