package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Segment 10: Add/Update variants, Has, Filter (L1-700)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovHM1_01_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_CovHM1_01_IsEmpty_HasItems", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		actual := args.Map{"result": hm.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": hm.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		hm.AddOrUpdate("a", "1")
		actual = args.Map{"result": hm.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual = args.Map{"result": hm.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected items", actual)
	})
}

func Test_CovHM1_02_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_CovHM1_02_IsEmptyLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		actual := args.Map{"result": hm.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_CovHM1_03_Collection(t *testing.T) {
	safeTest(t, "Test_CovHM1_03_Collection", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		col := hm.Collection()

		// Act
		actual := args.Map{"result": col.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM1_04_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_CovHM1_04_AddOrUpdateWithWgLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hm.AddOrUpdateWithWgLock("a", "1", wg)
		wg.Wait()

		// Act
		actual := args.Map{"result": hm.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_CovHM1_05_AddOrUpdateKeyStrValInt(t *testing.T) {
	safeTest(t, "Test_CovHM1_05_AddOrUpdateKeyStrValInt", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyStrValInt("age", 25)
		v, _ := hm.Get("age")

		// Act
		actual := args.Map{"result": v != "25"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '25', got ''", actual)
	})
}

func Test_CovHM1_06_AddOrUpdateKeyStrValFloat(t *testing.T) {
	safeTest(t, "Test_CovHM1_06_AddOrUpdateKeyStrValFloat", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyStrValFloat("f", 1.5)
		_, ok := hm.Get("f")

		// Act
		actual := args.Map{"result": ok}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
	})
}

func Test_CovHM1_07_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	safeTest(t, "Test_CovHM1_07_AddOrUpdateKeyStrValFloat64", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyStrValFloat64("f", 2.5)
		_, ok := hm.Get("f")

		// Act
		actual := args.Map{"result": ok}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
	})
}

func Test_CovHM1_08_AddOrUpdateKeyStrValAny(t *testing.T) {
	safeTest(t, "Test_CovHM1_08_AddOrUpdateKeyStrValAny", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyStrValAny("k", 42)

		// Act
		actual := args.Map{"result": hm.Has("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected k", actual)
	})
}

func Test_CovHM1_09_AddOrUpdateKeyValueAny(t *testing.T) {
	safeTest(t, "Test_CovHM1_09_AddOrUpdateKeyValueAny", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k", Value: 42})

		// Act
		actual := args.Map{"result": hm.Has("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected k", actual)
	})
}

func Test_CovHM1_10_AddOrUpdateKeyVal(t *testing.T) {
	safeTest(t, "Test_CovHM1_10_AddOrUpdateKeyVal", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		isNew := hm.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v"})

		// Act
		actual := args.Map{"result": isNew}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		isNew2 := hm.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v2"})
		actual = args.Map{"result": isNew2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not new", actual)
	})
}

func Test_CovHM1_11_AddOrUpdate_Set(t *testing.T) {
	safeTest(t, "Test_CovHM1_11_AddOrUpdate_Set", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		isNew := hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": isNew}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		isNew2 := hm.Set("a", "2")
		actual = args.Map{"result": isNew2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not new", actual)
	})
}

func Test_CovHM1_12_SetTrim(t *testing.T) {
	safeTest(t, "Test_CovHM1_12_SetTrim", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.SetTrim(" key ", " val ")

		// Act
		actual := args.Map{"result": hm.Has("key")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected trimmed key", actual)
	})
}

func Test_CovHM1_13_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_CovHM1_13_SetBySplitter", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.SetBySplitter("=", "key=value")
		v, _ := hm.Get("key")

		// Act
		actual := args.Map{"result": v != "value"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'value', got ''", actual)
		// no value
		hm.SetBySplitter("=", "onlykey")
		v2, _ := hm.Get("onlykey")
		actual = args.Map{"result": v2 != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty value", actual)
	})
}

func Test_CovHM1_14_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_CovHM1_14_AddOrUpdateStringsPtrWgLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{"a", "b"}, []string{"1", "2"})
		wg.Wait()

		// Act
		actual := args.Map{"result": hm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		wg.Add(1)
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{}, []string{})
		wg.Wait()
	})
}

func Test_CovHM1_15_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_CovHM1_15_AddOrUpdateHashmap", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		other := corestr.Empty.Hashmap()
		other.AddOrUpdate("x", "1")
		hm.AddOrUpdateHashmap(other)

		// Act
		actual := args.Map{"result": hm.Has("x")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
		hm.AddOrUpdateHashmap(nil)
	})
}

func Test_CovHM1_16_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_CovHM1_16_AddOrUpdateMap", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"result": hm.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		hm.AddOrUpdateMap(nil)
	})
}

func Test_CovHM1_17_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_CovHM1_17_AddsOrUpdates", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddsOrUpdates(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Act
		actual := args.Map{"result": hm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		hm.AddsOrUpdates()
	})
}

func Test_CovHM1_18_AddOrUpdateKeyAnyValues(t *testing.T) {
	safeTest(t, "Test_CovHM1_18_AddOrUpdateKeyAnyValues", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyAnyValues(
			corestr.KeyAnyValuePair{Key: "k", Value: 1},
		)

		// Act
		actual := args.Map{"result": hm.Has("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected k", actual)
		hm.AddOrUpdateKeyAnyValues()
	})
}

func Test_CovHM1_19_AddOrUpdateKeyValues(t *testing.T) {
	safeTest(t, "Test_CovHM1_19_AddOrUpdateKeyValues", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyValues(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Act
		actual := args.Map{"result": hm.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		hm.AddOrUpdateKeyValues()
	})
}

func Test_CovHM1_20_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_CovHM1_20_AddOrUpdateCollection", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		hm.AddOrUpdateCollection(keys, vals)

		// Act
		actual := args.Map{"result": hm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// nil
		hm.AddOrUpdateCollection(nil, nil)
		// length mismatch
		hm.AddOrUpdateCollection(
			corestr.New.Collection.Strings([]string{"a"}),
			corestr.New.Collection.Strings([]string{"1", "2"}),
		)
	})
}

func Test_CovHM1_21_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovHM1_21_AddsOrUpdatesAnyUsingFilter", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		filter := func(p corestr.KeyAnyValuePair) (string, bool, bool) {
			return p.ValueString(), true, false
		}
		hm.AddsOrUpdatesAnyUsingFilter(filter, corestr.KeyAnyValuePair{Key: "a", Value: 1})

		// Act
		actual := args.Map{"result": hm.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		// nil
		hm.AddsOrUpdatesAnyUsingFilter(filter)
		// break
		breakFilter := func(p corestr.KeyAnyValuePair) (string, bool, bool) {
			return "", true, true
		}
		hm.AddsOrUpdatesAnyUsingFilter(breakFilter, corestr.KeyAnyValuePair{Key: "b", Value: 2})
	})
}

func Test_CovHM1_22_AddsOrUpdatesAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_CovHM1_22_AddsOrUpdatesAnyUsingFilterLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		filter := func(p corestr.KeyAnyValuePair) (string, bool, bool) {
			return p.ValueString(), true, false
		}
		hm.AddsOrUpdatesAnyUsingFilterLock(filter, corestr.KeyAnyValuePair{Key: "a", Value: 1})

		// Act
		actual := args.Map{"result": hm.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		hm.AddsOrUpdatesAnyUsingFilterLock(filter)
		// break
		breakFilter := func(p corestr.KeyAnyValuePair) (string, bool, bool) {
			return "", true, true
		}
		hm.AddsOrUpdatesAnyUsingFilterLock(breakFilter, corestr.KeyAnyValuePair{Key: "x", Value: 1})
	})
}

func Test_CovHM1_23_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovHM1_23_AddsOrUpdatesUsingFilter", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		filter := func(p corestr.KeyValuePair) (string, bool, bool) {
			return p.Value, true, false
		}
		hm.AddsOrUpdatesUsingFilter(filter, corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		actual := args.Map{"result": hm.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		hm.AddsOrUpdatesUsingFilter(filter)
		// break
		breakFilter := func(p corestr.KeyValuePair) (string, bool, bool) {
			return "", true, true
		}
		hm.AddsOrUpdatesUsingFilter(breakFilter, corestr.KeyValuePair{Key: "x", Value: "1"})
	})
}

func Test_CovHM1_24_ConcatNew(t *testing.T) {
	safeTest(t, "Test_CovHM1_24_ConcatNew", func() {
		// Arrange
		a := corestr.Empty.Hashmap()
		a.AddOrUpdate("a", "1")
		// empty
		r := a.ConcatNew(true)

		// Act
		actual := args.Map{"result": r.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// with hashmap
		b := corestr.Empty.Hashmap()
		b.AddOrUpdate("b", "2")
		r2 := a.ConcatNew(false, b, nil)
		actual = args.Map{"result": r2.Length() < 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_CovHM1_25_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_CovHM1_25_ConcatNewUsingMaps", func() {
		// Arrange
		a := corestr.Empty.Hashmap()
		a.AddOrUpdate("a", "1")
		// empty
		r := a.ConcatNewUsingMaps(true)

		// Act
		actual := args.Map{"result": r.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// with map
		r2 := a.ConcatNewUsingMaps(false, map[string]string{"b": "2"}, nil)
		actual = args.Map{"result": r2.Length() < 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_CovHM1_26_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_CovHM1_26_AddOrUpdateLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateLock("a", "1")

		// Act
		actual := args.Map{"result": hm.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_CovHM1_27_Has_Contains_HasLock_HasWithLock(t *testing.T) {
	safeTest(t, "Test_CovHM1_27_Has_Contains_HasLock_HasWithLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": hm.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hm.Contains("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hm.ContainsLock("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hm.HasLock("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hm.HasWithLock("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovHM1_28_IsKeyMissing_Lock(t *testing.T) {
	safeTest(t, "Test_CovHM1_28_IsKeyMissing_Lock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": hm.IsKeyMissing("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		actual = args.Map{"result": hm.IsKeyMissing("z")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected missing", actual)
		actual = args.Map{"result": hm.IsKeyMissingLock("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
	})
}

func Test_CovHM1_29_HasAllStrings_HasAll(t *testing.T) {
	safeTest(t, "Test_CovHM1_29_HasAllStrings_HasAll", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")

		// Act
		actual := args.Map{"result": hm.HasAllStrings("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hm.HasAllStrings("a", "z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": hm.HasAll("a", "b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovHM1_30_HasAnyItem_HasAny(t *testing.T) {
	safeTest(t, "Test_CovHM1_30_HasAnyItem_HasAny", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		actual := args.Map{"result": hm.HasAnyItem()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		hm.AddOrUpdate("a", "1")
		actual = args.Map{"result": hm.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hm.HasAny("z", "a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hm.HasAny("x", "y")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovHM1_31_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_CovHM1_31_HasAllCollectionItems", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": hm.HasAllCollectionItems(corestr.New.Collection.Strings([]string{"a"}))}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hm.HasAllCollectionItems(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovHM1_32_DiffRaw_Diff(t *testing.T) {
	safeTest(t, "Test_CovHM1_32_DiffRaw_Diff", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		other := corestr.Empty.Hashmap()
		other.AddOrUpdate("b", "2")
		other.AddOrUpdate("c", "3")
		_ = hm.DiffRaw(other.Items())
		_ = hm.Diff(other)
	})
}

func Test_CovHM1_33_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_CovHM1_33_GetKeysFilteredItems", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		// empty
		r := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"result": len(r) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		// keep all
		r2 := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual = args.Map{"result": len(r2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// skip
		r3 := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, false, false
		})
		actual = args.Map{"result": len(r3) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// break
		r4 := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})
		actual = args.Map{"result": len(r4) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM1_34_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_CovHM1_34_GetKeysFilteredCollection", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		// empty
		col := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hm.AddOrUpdate("a", "1")
		// keep
		col2 := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual = args.Map{"result": col2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// break
		hm.AddOrUpdate("b", "2")
		col3 := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})
		actual = args.Map{"result": col3.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
