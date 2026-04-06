package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══ Hashmap ═══

func Test_C38_Hashmap_CRUD(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_CRUD", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("a", "1")
		h.Set("b", "2")
		h.SetTrim("  c  ", "  3  ")

		// Act
		actual := args.Map{"result": h.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": h.Has("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": h.Contains("b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": h.IsKeyMissing("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": h.IsKeyMissing("z")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_SetBySplitter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.SetBySplitter("=", "key=val")

		// Act
		actual := args.Map{
			"result": v,
			ok := h.Get("key"); !ok || v != "val",
		}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h.SetBySplitter("=", "noequals")
		actual := args.Map{
			"result": v,
			ok := h.Get("noequals"); !ok || v,
		}
		expected := args.Map{"result": ""}
		expected.ShouldBeEqual(t, 0, "v", actual)
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyStrValInt(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyStrValInt", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValInt("k", 42)
		v, _ := h.Get("k")

		// Act
		actual := args.Map{"result": v != "42"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyStrValFloat(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyStrValFloat", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValFloat("k", 3.14)
		v, _ := h.Get("k")

		// Act
		actual := args.Map{"result": v == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyStrValFloat64", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValFloat64("k", 2.71)
		v, _ := h.Get("k")

		// Act
		actual := args.Map{"result": v == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyStrValAny(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyStrValAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValAny("k", []int{1, 2})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyValueAny(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyValueAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k", Value: 42})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyVal(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyVal", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		isNew := h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v"})

		// Act
		actual := args.Map{"result": isNew}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		isNew2 := h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v2"})
		actual := args.Map{"result": isNew2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not new", actual)
	})
}

func Test_C38_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateHashmap", func() {
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
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h1.AddOrUpdateHashmap(nil)
	})
}

func Test_C38_Hashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateMap", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h.AddOrUpdateMap(nil)
	})
}

func Test_C38_Hashmap_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddsOrUpdates", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdates(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h.AddsOrUpdates(nil...)
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyAnyValues(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyAnyValues", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: 1})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h.AddOrUpdateKeyAnyValues()
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyValues(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyValues", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "k", Value: "v"})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h.AddOrUpdateKeyValues()
	})
}

func Test_C38_Hashmap_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		h.AddOrUpdateCollection(keys, vals)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// nil
		h.AddOrUpdateCollection(nil, nil)
		// mismatch
		h.AddOrUpdateCollection(keys, corestr.New.Collection.Strings([]string{"1"}))
	})
}

func Test_C38_Hashmap_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateWithWgLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateWithWgLock("a", "1", &wg)
		wg.Wait()

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateStringsPtrWgLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(&wg, []string{"a"}, []string{"1"})
		wg.Wait()

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateLock("a", "1")

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_Lock_Methods(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Lock_Methods", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": h.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.ContainsLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.HasLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.IsKeyMissingLock("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.HasWithLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_HasAllStrings", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(3)
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")

		// Act
		actual := args.Map{"result": h.HasAllStrings("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.HasAllStrings("a", "z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_HasAll(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_HasAll", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": h.HasAll("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.HasAll("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_HasAnyItem", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": h.HasAnyItem()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_HasAny(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_HasAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": h.HasAny("z", "a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.HasAny("x", "y")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_HasAllCollectionItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": h.HasAllCollectionItems(c)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.HasAllCollectionItems(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_Diff(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Diff", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		_ = h.DiffRaw(map[string]string{"b": "2"})
	})
}

func Test_C38_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_ConcatNew", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("b", "2")
		cn := h.ConcatNew(true, h2)

		// Act
		actual := args.Map{"result": cn.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		cn2 := h.ConcatNew(true)
		actual := args.Map{"result": cn2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_ConcatNewUsingMaps", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		cn := h.ConcatNewUsingMaps(true, map[string]string{"b": "2"})

		// Act
		actual := args.Map{"result": cn.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		cn2 := h.ConcatNewUsingMaps(true)
		actual := args.Map{"result": cn2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_GetKeysFilteredItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		r := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })

		// Act
		actual := args.Map{"result": len(r) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r2 := corestr.Empty.Hashmap().GetKeysFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })
		actual := args.Map{"result": len(r2) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_GetKeysFilteredCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		fc := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })

		// Act
		actual := args.Map{"result": fc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		fc2 := corestr.Empty.Hashmap().GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })
		actual := args.Map{"result": fc2.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddsOrUpdatesUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdatesUsingFilter(
			func(p corestr.KeyValuePair) (string, bool, bool) { return p.Value, true, false },
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h.AddsOrUpdatesUsingFilter(func(p corestr.KeyValuePair) (string, bool, bool) { return "", false, false }, nil...)
	})
}

func Test_C38_Hashmap_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddsOrUpdatesAnyUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdatesAnyUsingFilter(
			func(p corestr.KeyAnyValuePair) (string, bool, bool) { return "v", true, false },
			corestr.KeyAnyValuePair{Key: "k", Value: 42},
		)

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_AddsOrUpdatesAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddsOrUpdatesAnyUsingFilterLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdatesAnyUsingFilterLock(
			func(p corestr.KeyAnyValuePair) (string, bool, bool) { return "v", true, false },
			corestr.KeyAnyValuePair{Key: "k", Value: 42},
		)

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_Values_Keys(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Values_Keys", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": len(h.ValuesList()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(h.Keys()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(h.AllKeys()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.ValuesCollection().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.ValuesHashset().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.KeysCollection().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(h.KeysLock()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(h.ValuesListCopyLock()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.ValuesCollectionLock().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.ValuesHashsetLock().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.Collection().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_KeysValuesList(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeysValuesList", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		k, v := h.KeysValuesList()

		// Act
		actual := args.Map{"result": len(k) != 1 || len(v) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_KeysValuesCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeysValuesCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		k, v := h.KeysValuesCollection()

		// Act
		actual := args.Map{"result": k.Length() != 1 || v.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_KeysValuesListLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeysValuesListLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		k, v := h.KeysValuesListLock()

		// Act
		actual := args.Map{"result": len(k) != 1 || len(v) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_KeysValuePairs(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeysValuePairs", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		pairs := h.KeysValuePairs()

		// Act
		actual := args.Map{"result": len(pairs) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_KeysValuePairsCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeysValuePairsCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		pc := h.KeysValuePairsCollection()

		// Act
		actual := args.Map{"result": pc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_Items_SafeItems_ItemsCopyLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Items_SafeItems_ItemsCopyLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": len(h.Items()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(h.SafeItems()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		cm := h.ItemsCopyLock()
		actual := args.Map{"result": len(*cm) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		var nilH *corestr.Hashmap
		actual := args.Map{"result": nilH.SafeItems() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_IsEqual", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(2)
		h1.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": h1.IsEqualPtr(h2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h1.IsEqualPtrLock(h2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h1.IsEqual(*h2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// nil checks
		var nilH *corestr.Hashmap
		actual := args.Map{"result": nilH.IsEqualPtr(nil)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": nilH.IsEqualPtr(h1)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// same ptr
		actual := args.Map{"result": h1.IsEqualPtr(h1)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// both empty
		e1, e2 := corestr.Empty.Hashmap(), corestr.Empty.Hashmap()
		actual := args.Map{"result": e1.IsEqualPtr(e2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// one empty
		actual := args.Map{"result": h1.IsEqualPtr(corestr.Empty.Hashmap())}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// diff lengths
		h3 := corestr.New.Hashmap.Cap(2)
		h3.AddOrUpdate("a", "1")
		h3.AddOrUpdate("b", "2")
		actual := args.Map{"result": h1.IsEqualPtr(h3)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// diff values
		h4 := corestr.New.Hashmap.Cap(2)
		h4.AddOrUpdate("a", "999")
		actual := args.Map{"result": h1.IsEqualPtr(h4)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_Remove(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Remove", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h.Remove("a")

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_RemoveWithLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h.RemoveWithLock("a")

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_String", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": h.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": corestr.Empty.Hashmap().String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_StringLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_StringLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": h.StringLock() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_GetValuesExceptKeysInHashset(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_GetValuesExceptKeysInHashset", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(3)
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := h.GetValuesExceptKeysInHashset(hs)

		// Act
		actual := args.Map{"result": len(r) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r2 := h.GetValuesExceptKeysInHashset(nil)
		actual := args.Map{"result": len(r2) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_GetValuesKeysExcept(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_GetValuesKeysExcept", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		r := h.GetValuesKeysExcept(nil)

		// Act
		actual := args.Map{"result": len(r) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_GetAllExceptCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		r := h.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{"result": len(r) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_Join(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Join", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": h.Join(",") == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.JoinKeys(",") == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeysToLower", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("UPPER", "val")
		lower := h.KeysToLower()

		// Act
		actual := args.Map{"result": lower.Has("upper")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_ValuesToLower(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_ValuesToLower", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("UPPER", "val")
		_ = h.ValuesToLower()
	})
}

func Test_C38_Hashmap_ToError(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_ToError", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": h.ToError(",") == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.ToDefaultError() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_KeyValStringLines(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeyValStringLines", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": len(h.KeyValStringLines()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_ToStringsUsingCompiler", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		r := h.ToStringsUsingCompiler(func(k, v string) string { return k + v })

		// Act
		actual := args.Map{"result": len(r) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r2 := corestr.Empty.Hashmap().ToStringsUsingCompiler(func(k, v string) string { return "" })
		actual := args.Map{"result": len(r2) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_Clone(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Clone", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		c := h.Clone()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		cp := h.ClonePtr()
		actual := args.Map{"result": cp.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		var nilH *corestr.Hashmap
		actual := args.Map{"result": nilH.ClonePtr() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// clone empty
		ec := corestr.Empty.Hashmap().Clone()
		actual := args.Map{"result": ec.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Clear_Dispose", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h.Clear()

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		var nilH *corestr.Hashmap
		actual := args.Map{"result": nilH.Clear() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Dispose()
		var nilH2 *corestr.Hashmap
		nilH2.Dispose()
	})
}

func Test_C38_Hashmap_JSON(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_JSON", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		j := h.Json()

		// Act
		actual := args.Map{"hasError": j.HasError()}

		// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "Json returns no error", actual)
		jp := h.JsonPtr()
		actual := args.Map{"hasError": jp.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns no error", actual)
		actual := args.Map{"result": h.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		b, err := h.MarshalJSON()
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": len(b) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h2 := &corestr.Hashmap{}
		err2 := h2.UnmarshalJSON(b)
		actual := args.Map{"result": err2}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err2", actual)
		// invalid
		err3 := h2.UnmarshalJSON([]byte(`{invalid`))
		actual := args.Map{"result": err3 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_ParseInjectUsingJson", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		jr := h.JsonPtr()
		h2 := &corestr.Hashmap{}
		result, err := h2.ParseInjectUsingJson(jr)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": result.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Serialize_Deserialize", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		b, err := h.Serialize()
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": len(b) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		var target map[string]string
		err2 := h.Deserialize(&target)
		actual := args.Map{"result": err2}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err2", actual)
	})
}

func Test_C38_Hashmap_Get_GetValue(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Get_GetValue", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		v, ok := h.Get("a")
		actual := args.Map{"result": ok || v != "1"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v2, ok2 := h.GetValue("a")
		actual := args.Map{"result": ok2 || v2 != "1"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_InterfaceCasts", func() {
		h := corestr.New.Hashmap.Cap(2)
		actual := args.Map{"result": h.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": h.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashmap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_JsonParseSelfInject", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		jr := h.JsonPtr()
		h2 := &corestr.Hashmap{}
		err := h2.JsonParseSelfInject(jr)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

	// ── newHashmapCreator ──

func Test_C38_NewHashmapCreator_Methods(t *testing.T) {
	safeTest(t, "Test_C38_NewHashmapCreator_Methods", func() {
		h1 := corestr.New.Hashmap.Empty()
		actual := args.Map{"result": h1.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h2 := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		actual := args.Map{"result": h2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h3 := corestr.New.Hashmap.KeyValues()
		actual := args.Map{"result": h3.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h4 := corestr.New.Hashmap.KeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: 1})
		actual := args.Map{"result": h4.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h5 := corestr.New.Hashmap.KeyAnyValues()
		actual := args.Map{"result": h5.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h6 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		actual := args.Map{"result": h6.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h7 := corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{"a": "1"})
		actual := args.Map{"result": h7.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h8 := corestr.New.Hashmap.UsingMapOptions(false, 0, nil)
		actual := args.Map{"result": h8.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h9 := corestr.New.Hashmap.UsingMapOptions(false, 0, map[string]string{"a": "1"})
		actual := args.Map{"result": h9.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h10 := corestr.New.Hashmap.KeyValuesStrings([]string{"a"}, []string{"1"})
		actual := args.Map{"result": h10.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h11 := corestr.New.Hashmap.KeyValuesStrings(nil, nil)
		actual := args.Map{"result": h11.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		k := corestr.New.Collection.Strings([]string{"a"})
		v := corestr.New.Collection.Strings([]string{"1"})
		h12 := corestr.New.Hashmap.KeyValuesCollection(k, v)
		actual := args.Map{"result": h12.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h13 := corestr.New.Hashmap.KeyValuesCollection(nil, nil)
		actual := args.Map{"result": h13.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h14 := corestr.New.Hashmap.MapWithCap(0, map[string]string{"a": "1"})
		actual := args.Map{"result": h14.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h15 := corestr.New.Hashmap.MapWithCap(5, map[string]string{"a": "1"})
		actual := args.Map{"result": h15.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h16 := corestr.New.Hashmap.MapWithCap(5, nil)
		actual := args.Map{"result": h16.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

	// ═══ Hashset (comprehensive) ═══

func Test_C38_Hashset_CRUD(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_CRUD", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.Add("a")
		hs.AddNonEmpty("")
		hs.AddNonEmpty("b")
		hs.AddNonEmptyWhitespace("  ")
		hs.AddNonEmptyWhitespace("c")
		actual := args.Map{"result": hs.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual := args.Map{"result": hs.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.Contains("b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.IsMissing("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.IsMissing("z")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_AddIf_AddIfMany(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddIf_AddIfMany", func() {
		hs := corestr.New.Hashset.Cap(3)
		hs.AddIf(false, "skip")
		hs.AddIf(true, "keep")
		hs.AddIfMany(false, "x", "y")
		hs.AddIfMany(true, "m", "n")
		actual := args.Map{"result": hs.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_AddFunc_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddFunc_AddFuncErr", func() {
		hs := corestr.New.Hashset.Cap(2)
		hs.AddFunc(func() string { return "fn" })
		actual := args.Map{"result": hs.Has("fn")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		hs.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		actual := args.Map{"result": hs.Has("ok")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		hs.AddFuncErr(func() (string, error) { return "", errForTest }, func(e error) {})
	})
}

func Test_C38_Hashset_AddBool(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddBool", func() {
		hs := corestr.New.Hashset.Cap(2)
		existed := hs.AddBool("a")
		actual := args.Map{"result": existed}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		existed2 := hs.AddBool("a")
		actual := args.Map{"result": existed2}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C38_Hashset_Adds_AddStrings(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Adds_AddStrings", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.Adds("a", "b")
		hs.AddStrings([]string{"c"})
		actual := args.Map{"result": hs.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		hs.Adds(nil...)
		hs.AddStrings(nil)
	})
}

func Test_C38_Hashset_AddCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddCollection", func() {
		hs := corestr.New.Hashset.Cap(5)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs.AddCollection(c)
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		hs.AddCollection(nil)
	})
}

func Test_C38_Hashset_AddCollections(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddCollections", func() {
		hs := corestr.New.Hashset.Cap(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		hs.AddCollections(c1, c2, nil)
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddHashsetItems", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs2 := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs.AddHashsetItems(hs2)
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		hs.AddHashsetItems(nil)
	})
}

func Test_C38_Hashset_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddItemsMap", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddItemsMap(map[string]bool{"a": true, "b": false})
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1, false items skipped", actual)
		hs.AddItemsMap(nil)
	})
}

func Test_C38_Hashset_Lock_Methods(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Lock_Methods", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.AddLock("b")
		actual := args.Map{"result": hs.HasLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.IsMissingLock("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.IsEmptyLock()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.LengthLock() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.HasWithLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_AddPtr_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddPtr_AddPtrLock", func() {
		hs := corestr.New.Hashset.Cap(2)
		s := "ptr"
		hs.AddPtr(&s)
		actual := args.Map{"result": hs.Has("ptr")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s2 := "lock"
		hs.AddPtrLock(&s2)
		actual := args.Map{"result": hs.Has("lock")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddStringsLock", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddStringsLock([]string{"a", "b"})
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		hs.AddStringsLock(nil)
	})
}

func Test_C38_Hashset_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddSimpleSlice", func() {
		hs := corestr.New.Hashset.Cap(5)
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		hs.AddSimpleSlice(ss)
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_HasAll_HasAny_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_HasAll_HasAny_IsAllMissing", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		actual := args.Map{"result": hs.HasAll("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.HasAll("a", "z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.HasAny("z", "a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.HasAny("x", "y")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.IsAllMissing("x", "y")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.IsAllMissing("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_HasAllStrings", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		actual := args.Map{"result": hs.HasAllStrings([]string{"a"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_HasAllCollectionItems", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": hs.HasAllCollectionItems(c)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.HasAllCollectionItems(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_IsEquals(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_IsEquals", func() {
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		b := corestr.New.Hashset.StringsSpreadItems("a", "b")
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": a.IsEqualsLock(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": a.IsEqual(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_ConcatNew", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs2 := corestr.New.Hashset.StringsSpreadItems("b")
		cn := hs.ConcatNewHashsets(true, hs2)
		actual := args.Map{"result": cn.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		cn2 := hs.ConcatNewHashsets(true)
		actual := args.Map{"result": cn2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_ConcatNewStrings", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		cn := hs.ConcatNewStrings(true, []string{"b"})
		actual := args.Map{"result": cn.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		cn2 := hs.ConcatNewStrings(true)
		actual := args.Map{"result": cn2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Resize", func() {
		hs := corestr.New.Hashset.Cap(2)
		hs.Add("a")
		hs.Resize(100)
		actual := args.Map{"result": hs.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		hs.Resize(1) // no-op
	})
}

func Test_C38_Hashset_ResizeLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_ResizeLock", func() {
		hs := corestr.New.Hashset.Cap(2)
		hs.Add("a")
		hs.ResizeLock(100)
		actual := args.Map{"result": hs.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_AddCapacities(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddCapacities", func() {
		hs := corestr.New.Hashset.Cap(2)
		hs.AddCapacities(10, 20)
		hs.AddCapacities()
		hs.AddCapacitiesLock(10)
		hs.AddCapacitiesLock()
	})
}

func Test_C38_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Filter", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("abc", "xyz")
		filtered := hs.Filter(func(s string) bool { return s == "abc" })
		actual := args.Map{"result": filtered.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_GetFilteredItems(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_GetFilteredItems", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "bb")
		r := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, len(s) > 1, false })
		actual := args.Map{"result": len(r) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r2 := corestr.Empty.Hashset().GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })
		actual := args.Map{"result": len(r2) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_GetFilteredCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_GetFilteredCollection", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		fc := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })
		actual := args.Map{"result": fc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddsUsingFilter", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, false }, "a", "b")
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_AddsAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddsAnyUsingFilter", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, false }, 42, nil)
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_Remove_SafeRemove(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Remove_SafeRemove", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs.Remove("a")
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		hs.SafeRemove("b")
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		hs.SafeRemove("z") // no-op
	})
}

func Test_C38_Hashset_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_RemoveWithLock", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.RemoveWithLock("a")
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_List_OrderedList_SortedList(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_List_OrderedList_SortedList", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("b", "a")
		actual := args.Map{"result": len(hs.List()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(hs.OrderedList()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(hs.SortedList()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(hs.ListPtrSortedAsc()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(hs.ListPtrSortedDsc()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(hs.SafeStrings()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(hs.Lines()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(hs.ListPtr()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(hs.ListCopyLock()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_Collection_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Collection_SimpleSlice", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		actual := args.Map{"result": hs.Collection().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.SimpleSlice().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": corestr.Empty.Hashset().SimpleSlice().Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_String_StringLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_String_StringLock", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		actual := args.Map{"result": hs.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.StringLock() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": corestr.Empty.Hashset().String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_Join_JoinSorted_JoinLine(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Join_JoinSorted_JoinLine", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		_ = hs.Join(",")
		_ = hs.JoinSorted(",")
		_ = hs.JoinLine()
		_ = hs.NonEmptyJoins(",")
		_ = hs.NonWhitespaceJoins(",")
		_ = corestr.Empty.Hashset().JoinSorted(",")
	})
}

func Test_C38_Hashset_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_GetAllExcept", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.GetAllExcept([]string{"a"})
		actual := args.Map{"result": len(r) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r2 := hs.GetAllExcept(nil)
		actual := args.Map{"result": len(r2) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_GetAllExceptSpread(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_GetAllExceptSpread", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.GetAllExceptSpread("a")
		actual := args.Map{"result": len(r) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_GetAllExceptCollection", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.GetAllExceptCollection(nil)
		actual := args.Map{"result": len(r) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_GetAllExceptHashset(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_GetAllExceptHashset", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.GetAllExceptHashset(nil)
		actual := args.Map{"result": len(r) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_MapStringAny(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_MapStringAny", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		m := hs.MapStringAny()
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": len(corestr.Empty.Hashset().MapStringAny()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = hs.MapStringAnyDiff()
	})
}

func Test_C38_Hashset_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_DistinctDiff", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.DistinctDiffLinesRaw("b", "c")
		actual := args.Map{"result": len(r) < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r2 := hs.DistinctDiffLinesRaw()
		actual := args.Map{"result": len(r2) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		e := corestr.Empty.Hashset()
		r3 := e.DistinctDiffLinesRaw("x")
		actual := args.Map{"result": len(r3) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r4 := e.DistinctDiffLinesRaw()
		actual := args.Map{"result": len(r4) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_DistinctDiffLines(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_DistinctDiffLines", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.DistinctDiffLines("b")
		actual := args.Map{"result": len(r) < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = hs.DistinctDiffHashset(corestr.New.Hashset.StringsSpreadItems("b"))
		r2 := hs.DistinctDiffLines()
		actual := args.Map{"result": len(r2) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		e := corestr.Empty.Hashset()
		r3 := e.DistinctDiffLines("x")
		actual := args.Map{"result": len(r3) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_ToLowerSet", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("ABC")
		lower := hs.ToLowerSet()
		actual := args.Map{"result": lower.Has("abc")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_WrapQuotes(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_WrapQuotes", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		_ = hs.WrapDoubleQuote()
		hs2 := corestr.New.Hashset.StringsSpreadItems("a")
		_ = hs2.WrapSingleQuote()
		hs3 := corestr.New.Hashset.StringsSpreadItems("a")
		_ = hs3.WrapDoubleQuoteIfMissing()
		hs4 := corestr.New.Hashset.StringsSpreadItems("a")
		_ = hs4.WrapSingleQuoteIfMissing()
		// empty transpile
		_ = corestr.Empty.Hashset().WrapDoubleQuote()
	})
}

func Test_C38_Hashset_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Clear_Dispose", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.Clear()
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		hs.Dispose()
		var nilH *corestr.Hashset
		nilH.Dispose()
	})
}

func Test_C38_Hashset_JSON(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_JSON", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		j := hs.Json()
		actual := args.Map{"hasError": j.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "Json returns no error", actual)
		jp := hs.JsonPtr()
		actual := args.Map{"hasError": jp.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns no error", actual)
		actual := args.Map{"result": hs.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		b, err := hs.MarshalJSON()
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		hs2 := &corestr.Hashset{}
		err2 := hs2.UnmarshalJSON(b)
		actual := args.Map{"result": err2}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err2", actual)
		err3 := hs2.UnmarshalJSON([]byte(`{bad`))
		actual := args.Map{"result": err3 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// empty json model
		_ = corestr.Empty.Hashset().JsonModel()
	})
}
func Test_C38_Hashset_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_ParseInjectUsingJson", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		jr := hs.JsonPtr()
		hs2 := &corestr.Hashset{}
		result, err := hs2.ParseInjectUsingJson(jr)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": result.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C38_Hashset_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Serialize_Deserialize", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		b, err := hs.Serialize()
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": len(b) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		var target map[string]bool
		err2 := hs.Deserialize(&target)
		actual := args.Map{"result": err2}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err2", actual)
	})
}

func Test_C38_Hashset_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_InterfaceCasts", func() {
		hs := corestr.New.Hashset.Cap(1)
		actual := args.Map{"result": hs.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual := args.Map{"result": hs.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		err := hs.JsonParseSelfInject(hs.JsonPtr())
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C38_Hashset_Wg_Methods(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Wg_Methods", func() {
		hs := corestr.New.Hashset.Cap(5)
		wg := sync.WaitGroup{}
		wg.Add(1)
		hs.AddWithWgLock("a", &wg)
		wg.Wait()
		wg2 := sync.WaitGroup{}
		wg2.Add(1)
		hs.AddStringsPtrWgLock([]string{"b"}, &wg2)
		wg2.Wait()
		wg3 := sync.WaitGroup{}
		wg3.Add(1)
		hs.AddHashsetWgLock(corestr.New.Hashset.StringsSpreadItems("c"), &wg3)
		wg3.Wait()
		wg4 := sync.WaitGroup{}
		wg4.Add(1)
		m := map[string]bool{"d": true}
		hs.AddItemsMapWgLock(&m, &wg4)
		wg4.Wait()
		actual := args.Map{"result": hs.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

	// ── newHashsetCreator ──

func Test_C38_NewHashsetCreator_Methods(t *testing.T) {
	safeTest(t, "Test_C38_NewHashsetCreator_Methods", func() {
		h1 := corestr.New.Hashset.Empty()
		actual := args.Map{"result": h1.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h2 := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"result": h2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h3 := corestr.New.Hashset.Strings(nil)
		actual := args.Map{"result": h3.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h4 := corestr.New.Hashset.StringsSpreadItems("a", "b")
		actual := args.Map{"result": h4.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h5 := corestr.New.Hashset.StringsOption(5, true, "a")
		actual := args.Map{"result": h5.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h6 := corestr.New.Hashset.StringsOption(0, false)
		actual := args.Map{"result": h6.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h7 := corestr.New.Hashset.StringsOption(5, false)
		actual := args.Map{"result": h7.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h8 := corestr.New.Hashset.UsingMap(map[string]bool{"a": true})
		actual := args.Map{"result": h8.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h9 := corestr.New.Hashset.UsingMap(nil)
		actual := args.Map{"result": h9.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h10 := corestr.New.Hashset.UsingMapOption(5, true, map[string]bool{"a": true})
		actual := args.Map{"result": h10.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h11 := corestr.New.Hashset.UsingMapOption(0, false, nil)
		actual := args.Map{"result": h11.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h12 := corestr.New.Hashset.UsingCollection(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": h12.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h13 := corestr.New.Hashset.UsingCollection(nil)
		actual := args.Map{"result": h13.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		ss := corestr.New.SimpleSlice.Lines("a")
		h14 := corestr.New.Hashset.SimpleSlice(ss)
		actual := args.Map{"result": h14.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		h15 := corestr.New.Hashset.SimpleSlice(corestr.Empty.SimpleSlice())
		actual := args.Map{"result": h15.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}
