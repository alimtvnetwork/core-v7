package corestrtests

import (
	"errors"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Hashmap.go — Full coverage (~458 uncovered stmts, 1300 lines)
// =============================================================================

// ── IsEmpty / HasItems ──

func Test_Hashmap_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEmpty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{
			"empty": hm.IsEmpty(),
			"hasItems": hm.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty on empty", actual)
	})
}

func Test_Hashmap_IsEmpty_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEmpty_NonEmpty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{
			"empty": hm.IsEmpty(),
			"hasItems": hm.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": false,
			"hasItems": true,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty on non-empty", actual)
	})
}

func Test_Hashmap_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEmpty_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{"empty": hm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmpty on nil", actual)
	})
}

func Test_Hashmap_Collection(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Collection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		col := hm.Collection()

		// Act
		actual := args.Map{"len": col.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns values", actual)
	})
}

func Test_Hashmap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEmptyLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"empty": hm.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock", actual)
	})
}

// ── AddOrUpdate variants ──

func Test_Hashmap_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateWithWgLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hm.AddOrUpdateWithWgLock("k", "v", wg)
		wg.Wait()

		// Act
		actual := args.Map{"has": hm.Has("k")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateWithWgLock", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValInt(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateKeyStrValInt", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdateKeyStrValInt("k", 42)
		v, _ := hm.Get("k")

		// Act
		actual := args.Map{"val": v}

		// Assert
		expected := args.Map{"val": "42"}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyStrValInt", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValFloat(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateKeyStrValFloat", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdateKeyStrValFloat("k", 1.5)
		_, found := hm.Get("k")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyStrValFloat", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateKeyStrValFloat64", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdateKeyStrValFloat64("k", 2.5)
		_, found := hm.Get("k")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyStrValFloat64", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValAny(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateKeyStrValAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdateKeyStrValAny("k", 123)
		_, found := hm.Get("k")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyStrValAny", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyValueAny(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateKeyValueAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		pair := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		hm.AddOrUpdateKeyValueAny(pair)

		// Act
		actual := args.Map{"has": hm.Has("k")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyValueAny", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyVal_New(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateKeyVal_New", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		isNew := hm.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v"})

		// Act
		actual := args.Map{"isNew": isNew}

		// Assert
		expected := args.Map{"isNew": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyVal new", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyVal_Existing(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateKeyVal_Existing", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v1")
		isNew := hm.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v2"})

		// Act
		actual := args.Map{"isNew": isNew}

		// Assert
		expected := args.Map{"isNew": false}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyVal existing", actual)
	})
}

func Test_Hashmap_AddOrUpdate_New(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdate_New", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		isNew := hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"isNew": isNew}

		// Assert
		expected := args.Map{"isNew": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdate new", actual)
	})
}

func Test_Hashmap_AddOrUpdate_Existing(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdate_Existing", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v1")
		isNew := hm.AddOrUpdate("k", "v2")

		// Act
		actual := args.Map{"isNew": isNew}

		// Assert
		expected := args.Map{"isNew": false}
		expected.ShouldBeEqual(t, 0, "AddOrUpdate existing", actual)
	})
}

func Test_Hashmap_Set(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Set", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		isNew := hm.Set("k", "v")

		// Act
		actual := args.Map{"isNew": isNew}

		// Assert
		expected := args.Map{"isNew": true}
		expected.ShouldBeEqual(t, 0, "Set", actual)
	})
}

func Test_Hashmap_SetTrim(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_SetTrim", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.SetTrim("  k  ", "  v  ")
		v, found := hm.Get("k")

		// Act
		actual := args.Map{
			"found": found,
			"val": v,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"val": "v",
		}
		expected.ShouldBeEqual(t, 0, "SetTrim trims", actual)
	})
}

func Test_Hashmap_SetBySplitter_TwoParts(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_SetBySplitter_TwoParts", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.SetBySplitter("=", "key=value")
		v, _ := hm.Get("key")

		// Act
		actual := args.Map{"val": v}

		// Assert
		expected := args.Map{"val": "value"}
		expected.ShouldBeEqual(t, 0, "SetBySplitter two parts", actual)
	})
}

func Test_Hashmap_SetBySplitter_OnePart(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_SetBySplitter_OnePart", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.SetBySplitter("=", "key")
		v, _ := hm.Get("key")

		// Act
		actual := args.Map{"val": v}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "SetBySplitter one part", actual)
	})
}

func Test_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateStringsPtrWgLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{"a", "b"}, []string{"1", "2"})
		wg.Wait()

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateStringsPtrWgLock", actual)
	})
}

func Test_Hashmap_AddOrUpdateStringsPtrWgLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateStringsPtrWgLock_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{}, []string{})
		wg.Wait()

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateStringsPtrWgLock empty", actual)
	})
}

func Test_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateHashmap", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("a", "1")
		b := corestr.New.Hashmap.Empty()
		b.AddOrUpdate("b", "2")
		a.AddOrUpdateHashmap(b)

		// Act
		actual := args.Map{"len": a.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateHashmap merges", actual)
	})
}

func Test_Hashmap_AddOrUpdateHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateHashmap_Nil", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("a", "1")
		a.AddOrUpdateHashmap(nil)

		// Act
		actual := args.Map{"len": a.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateHashmap nil", actual)
	})
}

func Test_Hashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateMap", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdateMap(map[string]string{"a": "1", "b": "2"})

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateMap", actual)
	})
}

func Test_Hashmap_AddOrUpdateMap_Empty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateMap_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdateMap(map[string]string{})

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateMap empty", actual)
	})
}

func Test_Hashmap_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddsOrUpdates", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddsOrUpdates(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdates", actual)
	})
}

func Test_Hashmap_AddsOrUpdates_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddsOrUpdates_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		var kvs []corestr.KeyValuePair
		hm.AddsOrUpdates(kvs...)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdates nil", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyAnyValues(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateKeyAnyValues", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdateKeyAnyValues(
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
		)

		// Act
		actual := args.Map{"has": hm.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyAnyValues", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyAnyValues_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateKeyAnyValues_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		var pairs []corestr.KeyAnyValuePair
		hm.AddOrUpdateKeyAnyValues(pairs...)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyAnyValues nil", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyValues(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateKeyValues", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdateKeyValues(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Act
		actual := args.Map{"has": hm.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyValues", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyValues_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateKeyValues_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		var pairs []corestr.KeyValuePair
		hm.AddOrUpdateKeyValues(pairs...)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyValues nil", actual)
	})
}

func Test_Hashmap_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		hm.AddOrUpdateCollection(keys, vals)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateCollection", actual)
	})
}

func Test_Hashmap_AddOrUpdateCollection_NilKeys(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateCollection_NilKeys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		vals := corestr.New.Collection.Strings([]string{"1"})
		hm.AddOrUpdateCollection(nil, vals)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateCollection nil keys", actual)
	})
}

func Test_Hashmap_AddOrUpdateCollection_MismatchLen(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateCollection_MismatchLen", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		keys := corestr.New.Collection.Strings([]string{"a"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		hm.AddOrUpdateCollection(keys, vals)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateCollection mismatch", actual)
	})
}

// ── Filter methods ──

func Test_Hashmap_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddsOrUpdatesAnyUsingFilter", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return "filtered", true, false
		}
		hm.AddsOrUpdatesAnyUsingFilter(filter, corestr.KeyAnyValuePair{Key: "a", Value: 1})
		v, _ := hm.Get("a")

		// Act
		actual := args.Map{"val": v}

		// Assert
		expected := args.Map{"val": "filtered"}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilter", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddsOrUpdatesAnyUsingFilter_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return "v", true, true
		}
		hm.AddsOrUpdatesAnyUsingFilter(filter,
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
			corestr.KeyAnyValuePair{Key: "b", Value: 2},
		)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilter break", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilter_Skip(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddsOrUpdatesAnyUsingFilter_Skip", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return "", false, false
		}
		hm.AddsOrUpdatesAnyUsingFilter(filter, corestr.KeyAnyValuePair{Key: "a", Value: 1})

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilter skip", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddsOrUpdatesAnyUsingFilter_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return "", true, false
		}
		var pairs []corestr.KeyAnyValuePair
		hm.AddsOrUpdatesAnyUsingFilter(filter, pairs...)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilter nil", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddsOrUpdatesAnyUsingFilterLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return "v", true, false
		}
		hm.AddsOrUpdatesAnyUsingFilterLock(filter, corestr.KeyAnyValuePair{Key: "a", Value: 1})

		// Act
		actual := args.Map{"has": hm.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilterLock", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Break(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return "v", true, true
		}
		hm.AddsOrUpdatesAnyUsingFilterLock(filter,
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
			corestr.KeyAnyValuePair{Key: "b", Value: 2},
		)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilterLock break", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return "", true, false
		}
		var pairs []corestr.KeyAnyValuePair
		hm.AddsOrUpdatesAnyUsingFilterLock(filter, pairs...)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilterLock nil", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddsOrUpdatesUsingFilter", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Value + "!", true, false
		}
		hm.AddsOrUpdatesUsingFilter(filter, corestr.KeyValuePair{Key: "a", Value: "v"})
		v, _ := hm.Get("a")

		// Act
		actual := args.Map{"val": v}

		// Assert
		expected := args.Map{"val": "v!"}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesUsingFilter", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddsOrUpdatesUsingFilter_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return "", true, false
		}
		var pairs []corestr.KeyValuePair
		hm.AddsOrUpdatesUsingFilter(filter, pairs...)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesUsingFilter nil", actual)
	})
}

// ── ConcatNew ──

func Test_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ConcatNew", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("a", "1")
		b := corestr.New.Hashmap.Empty()
		b.AddOrUpdate("b", "2")
		r := a.ConcatNew(false, b)

		// Act
		actual := args.Map{
			"hasA": r.Has("a"),
			"hasB": r.Has("b"),
		}

		// Assert
		expected := args.Map{
			"hasA": true,
			"hasB": true,
		}
		expected.ShouldBeEqual(t, 0, "ConcatNew merges", actual)
	})
}

func Test_Hashmap_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ConcatNew_Empty", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("a", "1")
		r := a.ConcatNew(true)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty clones", actual)
	})
}

func Test_Hashmap_ConcatNew_NilHashmap(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ConcatNew_NilHashmap", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("a", "1")
		r := a.ConcatNew(false, nil)

		// Act
		actual := args.Map{"hasA": r.Has("a")}

		// Assert
		expected := args.Map{"hasA": true}
		expected.ShouldBeEqual(t, 0, "ConcatNew skips nil", actual)
	})
}

func Test_Hashmap_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ConcatNewUsingMaps", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("a", "1")
		r := a.ConcatNewUsingMaps(false, map[string]string{"b": "2"})

		// Act
		actual := args.Map{"hasB": r.Has("b")}

		// Assert
		expected := args.Map{"hasB": true}
		expected.ShouldBeEqual(t, 0, "ConcatNewUsingMaps", actual)
	})
}

func Test_Hashmap_ConcatNewUsingMaps_Empty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ConcatNewUsingMaps_Empty", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("a", "1")
		r := a.ConcatNewUsingMaps(true)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ConcatNewUsingMaps empty clones", actual)
	})
}

func Test_Hashmap_ConcatNewUsingMaps_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ConcatNewUsingMaps_Nil", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("a", "1")
		r := a.ConcatNewUsingMaps(false, nil)

		// Act
		actual := args.Map{"hasA": r.Has("a")}

		// Assert
		expected := args.Map{"hasA": true}
		expected.ShouldBeEqual(t, 0, "ConcatNewUsingMaps nil skipped", actual)
	})
}

func Test_Hashmap_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AddOrUpdateLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdateLock("k", "v")

		// Act
		actual := args.Map{"has": hm.Has("k")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateLock", actual)
	})
}

// ── Has / Contains / Missing ──

func Test_Hashmap_Has(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Has", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{
			"has": hm.Has("k"),
			"miss": hm.Has("z"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "Has", actual)
	})
}

func Test_Hashmap_Contains(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Contains", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"has": hm.Contains("k")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Contains", actual)
	})
}

func Test_Hashmap_ContainsLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ContainsLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"has": hm.ContainsLock("k")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "ContainsLock", actual)
	})
}

func Test_Hashmap_IsKeyMissing(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsKeyMissing", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{
			"missing": hm.IsKeyMissing("z"),
			"found": hm.IsKeyMissing("k"),
		}

		// Assert
		expected := args.Map{
			"missing": true,
			"found": false,
		}
		expected.ShouldBeEqual(t, 0, "IsKeyMissing", actual)
	})
}

func Test_Hashmap_IsKeyMissingLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsKeyMissingLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"missing": hm.IsKeyMissingLock("z")}

		// Assert
		expected := args.Map{"missing": true}
		expected.ShouldBeEqual(t, 0, "IsKeyMissingLock", actual)
	})
}

func Test_Hashmap_HasLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_HasLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"has": hm.HasLock("k")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasLock", actual)
	})
}

func Test_Hashmap_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_HasAllStrings", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")

		// Act
		actual := args.Map{
			"all": hm.HasAllStrings("a", "b"),
			"miss": hm.HasAllStrings("a", "z"),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAllStrings", actual)
	})
}

func Test_Hashmap_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_HasAllCollectionItems", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"has": hm.HasAllCollectionItems(col),
			"nil": hm.HasAllCollectionItems(nil),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAllCollectionItems", actual)
	})
}

func Test_Hashmap_HasAll(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_HasAll", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")

		// Act
		actual := args.Map{
			"all": hm.HasAll("a", "b"),
			"miss": hm.HasAll("a", "z"),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAll", actual)
	})
}

func Test_Hashmap_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_HasAnyItem", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"empty": hm.HasAnyItem()}

		// Assert
		expected := args.Map{"empty": false}
		expected.ShouldBeEqual(t, 0, "HasAnyItem empty", actual)
	})
}

func Test_Hashmap_HasAny(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_HasAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{
			"any": hm.HasAny("a", "z"),
			"none": hm.HasAny("x", "y"),
		}

		// Assert
		expected := args.Map{
			"any": true,
			"none": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAny", actual)
	})
}

func Test_Hashmap_HasWithLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_HasWithLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"has": hm.HasWithLock("k")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasWithLock", actual)
	})
}

// ── Diff ──

func Test_Hashmap_DiffRaw(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_DiffRaw", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		r := hm.DiffRaw(map[string]string{"b": "2", "c": "3"})

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "DiffRaw", actual)
	})
}

func Test_Hashmap_Diff(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Diff", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("a", "1")
		b := corestr.New.Hashmap.Empty()
		b.AddOrUpdate("b", "2")
		r := a.Diff(b)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "Diff", actual)
	})
}

// ── Filter methods ──

func Test_Hashmap_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetKeysFilteredItems", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("aa", "1")
		hm.AddOrUpdate("b", "2")
		r := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredItems", actual)
	})
}

func Test_Hashmap_GetKeysFilteredItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetKeysFilteredItems_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		r := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredItems empty", actual)
	})
}

func Test_Hashmap_GetKeysFilteredItems_Break(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetKeysFilteredItems_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		r := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredItems break", actual)
	})
}

func Test_Hashmap_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetKeysFilteredCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("aa", "1")
		hm.AddOrUpdate("b", "2")
		r := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"len": r.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredCollection", actual)
	})
}

func Test_Hashmap_GetKeysFilteredCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetKeysFilteredCollection_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		r := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"empty": r.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredCollection empty", actual)
	})
}

func Test_Hashmap_GetKeysFilteredCollection_Break(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetKeysFilteredCollection_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		r := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})

		// Act
		actual := args.Map{"len": r.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredCollection break", actual)
	})
}

// ── Items / Keys / Values ──

func Test_Hashmap_Items(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Items", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": len(hm.Items())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Items", actual)
	})
}

func Test_Hashmap_SafeItems(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_SafeItems", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"nonNil": hm.SafeItems() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SafeItems", actual)
	})
}

func Test_Hashmap_SafeItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_SafeItems_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{"nil": hm.SafeItems() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "SafeItems nil", actual)
	})
}

func Test_Hashmap_ItemsCopyLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ItemsCopyLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		cp := hm.ItemsCopyLock()

		// Act
		actual := args.Map{"len": len(*cp)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ItemsCopyLock", actual)
	})
}

func Test_Hashmap_ValuesCollection(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ValuesCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": hm.ValuesCollection().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesCollection", actual)
	})
}

func Test_Hashmap_ValuesHashset(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ValuesHashset", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": hm.ValuesHashset().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesHashset", actual)
	})
}

func Test_Hashmap_ValuesCollectionLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ValuesCollectionLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": hm.ValuesCollectionLock().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesCollectionLock", actual)
	})
}

func Test_Hashmap_ValuesHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ValuesHashsetLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": hm.ValuesHashsetLock().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesHashsetLock", actual)
	})
}

func Test_Hashmap_ValuesList(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ValuesList", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": len(hm.ValuesList())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesList", actual)
	})
}

func Test_Hashmap_KeysValuesCollection(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_KeysValuesCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		keys, values := hm.KeysValuesCollection()

		// Act
		actual := args.Map{
			"keysLen": keys.Length(),
			"valsLen": values.Length(),
		}

		// Assert
		expected := args.Map{
			"keysLen": 1,
			"valsLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "KeysValuesCollection", actual)
	})
}

func Test_Hashmap_KeysValuesList(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_KeysValuesList", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		keys, values := hm.KeysValuesList()

		// Act
		actual := args.Map{
			"keysLen": len(keys),
			"valsLen": len(values),
		}

		// Assert
		expected := args.Map{
			"keysLen": 1,
			"valsLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "KeysValuesList", actual)
	})
}

func Test_Hashmap_KeysValuePairs(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_KeysValuePairs", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		pairs := hm.KeysValuePairs()

		// Act
		actual := args.Map{"len": len(pairs)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeysValuePairs", actual)
	})
}

func Test_Hashmap_KeysValuePairsCollection(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_KeysValuePairsCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		kvc := hm.KeysValuePairsCollection()

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeysValuePairsCollection", actual)
	})
}

func Test_Hashmap_KeysValuesListLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_KeysValuesListLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		keys, values := hm.KeysValuesListLock()

		// Act
		actual := args.Map{
			"keysLen": len(keys),
			"valsLen": len(values),
		}

		// Assert
		expected := args.Map{
			"keysLen": 1,
			"valsLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "KeysValuesListLock", actual)
	})
}

func Test_Hashmap_AllKeys(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AllKeys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": len(hm.AllKeys())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AllKeys", actual)
	})
}

func Test_Hashmap_AllKeys_Empty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AllKeys_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"len": len(hm.AllKeys())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllKeys empty", actual)
	})
}

func Test_Hashmap_Keys(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Keys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": len(hm.Keys())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Keys", actual)
	})
}

func Test_Hashmap_KeysCollection(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_KeysCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": hm.KeysCollection().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeysCollection", actual)
	})
}

func Test_Hashmap_KeysLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_KeysLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": len(hm.KeysLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeysLock", actual)
	})
}

func Test_Hashmap_KeysLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_KeysLock_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"len": len(hm.KeysLock())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "KeysLock empty", actual)
	})
}

func Test_Hashmap_ValuesListCopyLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ValuesListCopyLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": len(hm.ValuesListCopyLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesListCopyLock", actual)
	})
}

// ── KeysToLower / ValuesToLower ──

func Test_Hashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_KeysToLower", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("KEY", "val")
		r := hm.KeysToLower()

		// Act
		actual := args.Map{"has": r.Has("key")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "KeysToLower", actual)
	})
}

func Test_Hashmap_ValuesToLower(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ValuesToLower", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("KEY", "val")
		r := hm.ValuesToLower()

		// Act
		actual := args.Map{"has": r.Has("key")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "ValuesToLower (deprecated alias)", actual)
	})
}

// ── Length ──

func Test_Hashmap_Length(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Length", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Length", actual)
	})
}

func Test_Hashmap_Length_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Length_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length nil", actual)
	})
}

func Test_Hashmap_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_LengthLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"len": hm.LengthLock()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthLock", actual)
	})
}

// ── IsEqual ──

func Test_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEqual", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("k", "v")
		b := *corestr.New.Hashmap.Empty()
		b.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"eq": a.IsEqual(b)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqual same content", actual)
	})
}

func Test_Hashmap_IsEqualPtr_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEqualPtr_SamePtr", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"eq": a.IsEqualPtr(a)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr same ptr", actual)
	})
}

func Test_Hashmap_IsEqualPtr_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEqualPtr_BothNil", func() {
		// Arrange
		var a *corestr.Hashmap

		// Act
		actual := args.Map{"eq": a.IsEqualPtr(nil)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr both nil", actual)
	})
}

func Test_Hashmap_IsEqualPtr_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEqualPtr_OneNil", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"eq": a.IsEqualPtr(nil)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr one nil", actual)
	})
}

func Test_Hashmap_IsEqualPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEqualPtr_BothEmpty", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		b := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"eq": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr both empty", actual)
	})
}

func Test_Hashmap_IsEqualPtr_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEqualPtr_OneEmpty", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("k", "v")
		b := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"eq": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr one empty", actual)
	})
}

func Test_Hashmap_IsEqualPtr_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEqualPtr_DiffLen", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("a", "1")
		b := corestr.New.Hashmap.Empty()
		b.AddOrUpdate("a", "1")
		b.AddOrUpdate("b", "2")

		// Act
		actual := args.Map{"eq": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr diff len", actual)
	})
}

func Test_Hashmap_IsEqualPtr_DiffContent(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEqualPtr_DiffContent", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("a", "1")
		b := corestr.New.Hashmap.Empty()
		b.AddOrUpdate("a", "2")

		// Act
		actual := args.Map{"eq": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr diff content", actual)
	})
}

func Test_Hashmap_IsEqualPtr_MissingKey(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEqualPtr_MissingKey", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("a", "1")
		b := corestr.New.Hashmap.Empty()
		b.AddOrUpdate("b", "1")

		// Act
		actual := args.Map{"eq": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr missing key", actual)
	})
}

func Test_Hashmap_IsEqualPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_IsEqualPtrLock", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("k", "v")
		b := corestr.New.Hashmap.Empty()
		b.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"eq": a.IsEqualPtrLock(b)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualPtrLock", actual)
	})
}

// ── Remove ──

func Test_Hashmap_Remove(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Remove", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		hm.Remove("k")

		// Act
		actual := args.Map{"has": hm.Has("k")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Remove", actual)
	})
}

func Test_Hashmap_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_RemoveWithLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		hm.RemoveWithLock("k")

		// Act
		actual := args.Map{"has": hm.Has("k")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "RemoveWithLock", actual)
	})
}

// ── String ──

func Test_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_String", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"nonEmpty": hm.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String", actual)
	})
}

func Test_Hashmap_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_String_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"nonEmpty": hm.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty", actual)
	})
}

func Test_Hashmap_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_StringLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"nonEmpty": hm.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock", actual)
	})
}

func Test_Hashmap_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_StringLock_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"nonEmpty": hm.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty", actual)
	})
}

// ── GetValues Except ──

func Test_Hashmap_GetValuesExceptKeysInHashset(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetValuesExceptKeysInHashset", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hm.GetValuesExceptKeysInHashset(hs)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetValuesExceptKeysInHashset", actual)
	})
}

func Test_Hashmap_GetValuesExceptKeysInHashset_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetValuesExceptKeysInHashset_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")
		r := hm.GetValuesExceptKeysInHashset(nil)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetValuesExceptKeysInHashset nil", actual)
	})
}

func Test_Hashmap_GetValuesKeysExcept(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetValuesKeysExcept", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		r := hm.GetValuesKeysExcept([]string{"a"})

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetValuesKeysExcept", actual)
	})
}

func Test_Hashmap_GetValuesKeysExcept_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetValuesKeysExcept_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")
		r := hm.GetValuesKeysExcept(nil)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetValuesKeysExcept nil", actual)
	})
}

func Test_Hashmap_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetAllExceptCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		col := corestr.New.Collection.Strings([]string{"a"})
		r := hm.GetAllExceptCollection(col)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection", actual)
	})
}

func Test_Hashmap_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetAllExceptCollection_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("a", "1")
		r := hm.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection nil", actual)
	})
}

// ── Join ──

func Test_Hashmap_Join(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Join", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"val": hm.Join(",")}

		// Assert
		expected := args.Map{"val": "v"}
		expected.ShouldBeEqual(t, 0, "Join", actual)
	})
}

func Test_Hashmap_JoinKeys(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_JoinKeys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"val": hm.JoinKeys(",")}

		// Assert
		expected := args.Map{"val": "k"}
		expected.ShouldBeEqual(t, 0, "JoinKeys", actual)
	})
}

// ── JSON ──

func Test_Hashmap_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_JsonModel", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": len(hm.JsonModel())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel", actual)
	})
}

func Test_Hashmap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_JsonModelAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"nonNil": hm.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny", actual)
	})
}

func Test_Hashmap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_MarshalJSON", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		b, err := hm.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonEmpty": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON", actual)
	})
}

func Test_Hashmap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_UnmarshalJSON", func() {
		// Arrange
		hm := &corestr.Hashmap{}
		err := hm.UnmarshalJSON([]byte(`{"a":"1"}`))

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": hm.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON", actual)
	})
}

func Test_Hashmap_UnmarshalJSON_Error(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_UnmarshalJSON_Error", func() {
		// Arrange
		hm := &corestr.Hashmap{}
		err := hm.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON error", actual)
	})
}

func Test_Hashmap_Json(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Json", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		r := hm.Json()

		// Act
		actual := args.Map{"nonEmpty": r.JsonString() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Json", actual)
	})
}

func Test_Hashmap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_JsonPtr", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		r := hm.JsonPtr()

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr", actual)
	})
}

func Test_Hashmap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ParseInjectUsingJson", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		jr := corejson.NewPtr(map[string]string{"a": "1"})
		r, err := hm.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonNil": r != nil,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonNil": true,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson", actual)
	})
}

func Test_Hashmap_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ParseInjectUsingJson_Error", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		jr := &corejson.Result{Error: errors.New("fail")}
		_, err := hm.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson error", actual)
	})
}

func Test_Hashmap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ParseInjectUsingJsonMust", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		jr := corejson.NewPtr(map[string]string{"a": "1"})
		r := hm.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust", actual)
	})
}

func Test_Hashmap_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ParseInjectUsingJsonMust_Panics", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		jr := &corejson.Result{Error: errors.New("fail")}
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			hm.ParseInjectUsingJsonMust(jr)
		}()

		// Act
		actual := args.Map{"panicked": panicked}

		// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust panics", actual)
	})
}

// ── Error ──

func Test_Hashmap_ToError(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ToError", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		err := hm.ToError(",")

		// Act
		actual := args.Map{"nonNil": err != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ToError", actual)
	})
}

func Test_Hashmap_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ToDefaultError", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		err := hm.ToDefaultError()

		// Act
		actual := args.Map{"nonNil": err != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ToDefaultError", actual)
	})
}

func Test_Hashmap_KeyValStringLines(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_KeyValStringLines", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		lines := hm.KeyValStringLines()

		// Act
		actual := args.Map{"len": len(lines)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeyValStringLines", actual)
	})
}

// ── Clear / Dispose ──

func Test_Hashmap_Clear(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Clear", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		hm.Clear()

		// Act
		actual := args.Map{"empty": hm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear", actual)
	})
}

func Test_Hashmap_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Clear_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap
		r := hm.Clear()

		// Act
		actual := args.Map{"nil": r == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clear nil", actual)
	})
}

func Test_Hashmap_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Dispose", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		hm.Dispose()

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Dispose", actual)
	})
}

func Test_Hashmap_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Dispose_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap
		hm.Dispose()

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Dispose nil", actual)
	})
}

// ── ToStringsUsingCompiler ──

func Test_Hashmap_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ToStringsUsingCompiler", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		r := hm.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ToStringsUsingCompiler", actual)
	})
}

func Test_Hashmap_ToStringsUsingCompiler_Empty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ToStringsUsingCompiler_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		r := hm.ToStringsUsingCompiler(func(k, v string) string { return k })

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ToStringsUsingCompiler empty", actual)
	})
}

// ── Interface casts ──

func Test_Hashmap_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AsJsoner", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"nonNil": hm.AsJsoner() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsoner", actual)
	})
}

func Test_Hashmap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_JsonParseSelfInject", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		jr := corejson.NewPtr(map[string]string{"a": "1"})
		err := hm.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject", actual)
	})
}

func Test_Hashmap_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AsJsonContractsBinder", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"nonNil": hm.AsJsonContractsBinder() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder", actual)
	})
}

func Test_Hashmap_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AsJsonParseSelfInjector", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"nonNil": hm.AsJsonParseSelfInjector() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonParseSelfInjector", actual)
	})
}

func Test_Hashmap_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_AsJsonMarshaller", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"nonNil": hm.AsJsonMarshaller() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonMarshaller", actual)
	})
}

// ── Clone ──

func Test_Hashmap_Clone(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Clone", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		cloned := hm.Clone()

		// Act
		actual := args.Map{"len": cloned.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Clone", actual)
	})
}

func Test_Hashmap_Clone_Empty(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Clone_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		cloned := hm.Clone()

		// Act
		actual := args.Map{"empty": cloned.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clone empty", actual)
	})
}

func Test_Hashmap_ClonePtr(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ClonePtr", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		r := hm.ClonePtr()

		// Act
		actual := args.Map{
			"nonNil": r != nil,
			"len": r.Length(),
		}

		// Assert
		expected := args.Map{
			"nonNil": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "ClonePtr", actual)
	})
}

func Test_Hashmap_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_ClonePtr_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap
		r := hm.ClonePtr()

		// Act
		actual := args.Map{"nil": r == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "ClonePtr nil", actual)
	})
}

// ── Get / GetValue ──

func Test_Hashmap_Get(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Get", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		val, found := hm.Get("k")

		// Act
		actual := args.Map{
			"val": val,
			"found": found,
		}

		// Assert
		expected := args.Map{
			"val": "v",
			"found": true,
		}
		expected.ShouldBeEqual(t, 0, "Get", actual)
	})
}

func Test_Hashmap_Get_Missing(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Get_Missing", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		_, found := hm.Get("z")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": false}
		expected.ShouldBeEqual(t, 0, "Get missing", actual)
	})
}

func Test_Hashmap_GetValue(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_GetValue", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		val, found := hm.GetValue("k")

		// Act
		actual := args.Map{
			"val": val,
			"found": found,
		}

		// Assert
		expected := args.Map{
			"val": "v",
			"found": true,
		}
		expected.ShouldBeEqual(t, 0, "GetValue", actual)
	})
}

// ── Serialize / Deserialize ──

func Test_Hashmap_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Serialize", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		b, err := hm.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonEmpty": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize", actual)
	})
}

func Test_Hashmap_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov62_Hashmap_Deserialize", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		var target map[string]string
		err := hm.Deserialize(&target)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": len(target),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "Deserialize", actual)
	})
}
