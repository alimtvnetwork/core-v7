package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — IsEmpty / HasItems / Length
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_IsEmpty_New(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_IsEmpty_New", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{
			"empty": hm.IsEmpty(),
			"items": hm.HasItems(),
			"len": hm.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"items": false,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- empty", actual)
	})
}

func Test_Hashmap_HasAnyItem_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_HasAnyItem", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{
			"hasAny": hm.HasAnyItem(),
			"len": hm.Length(),
		}

		// Assert
		expected := args.Map{
			"hasAny": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- HasAnyItem", actual)
	})
}

func Test_Hashmap_Length_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Length_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- nil length", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — AddOrUpdate variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_AddOrUpdate_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdate", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		isNew1 := hm.AddOrUpdate("k", "v1")
		isNew2 := hm.AddOrUpdate("k", "v2")
		v, _ := hm.Get("k")

		// Act
		actual := args.Map{
			"isNew1": isNew1,
			"isNew2": isNew2,
			"val": v,
		}

		// Assert
		expected := args.Map{
			"isNew1": true,
			"isNew2": false,
			"val": "v2",
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdate", actual)
	})
}

func Test_Hashmap_Set_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Set", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		isNew := hm.Set("k", "v")

		// Act
		actual := args.Map{"isNew": isNew}

		// Assert
		expected := args.Map{"isNew": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Set", actual)
	})
}

func Test_Hashmap_SetTrim_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_SetTrim", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
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
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- SetTrim", actual)
	})
}

func Test_Hashmap_SetBySplitter_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_SetBySplitter", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		isNew := hm.SetBySplitter("=", "key=value")
		v, _ := hm.Get("key")

		// Act
		actual := args.Map{
			"isNew": isNew,
			"val": v,
		}

		// Assert
		expected := args.Map{
			"isNew": true,
			"val": "value",
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- SetBySplitter", actual)
	})
}

func Test_Hashmap_SetBySplitter_NoSplit_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_SetBySplitter_NoSplit", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.SetBySplitter("=", "keyonly")
		v, found := hm.Get("keyonly")

		// Act
		actual := args.Map{
			"found": found,
			"val": v,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"val": "",
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- SetBySplitter no split", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValInt_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateKeyStrValInt", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdateKeyStrValInt("k", 42)
		v, _ := hm.Get("k")

		// Act
		actual := args.Map{"val": v}

		// Assert
		expected := args.Map{"val": "42"}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateKeyStrValInt", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValFloat_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateKeyStrValFloat", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdateKeyStrValFloat("k", 3.14)
		v, found := hm.Get("k")

		// Act
		actual := args.Map{
			"found": found,
			"notEmpty": v != "",
		}

		// Assert
		expected := args.Map{
			"found": true,
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateKeyStrValFloat", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValFloat64_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateKeyStrValFloat64", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdateKeyStrValFloat64("k", 2.718)
		v, found := hm.Get("k")

		// Act
		actual := args.Map{
			"found": found,
			"notEmpty": v != "",
		}

		// Assert
		expected := args.Map{
			"found": true,
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateKeyStrValFloat64", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValAny_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateKeyStrValAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdateKeyStrValAny("k", 99)
		_, found := hm.Get("k")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateKeyStrValAny", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyVal_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateKeyVal", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		kv := corestr.KeyValuePair{Key: "a", Value: "1"}
		isNew := hm.AddOrUpdateKeyVal(kv)

		// Act
		actual := args.Map{"isNew": isNew}

		// Assert
		expected := args.Map{"isNew": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateKeyVal", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyValueAny_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateKeyValueAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		kav := corestr.KeyAnyValuePair{Key: "x", Value: "hello"}
		hm.AddOrUpdateKeyValueAny(kav)
		v, _ := hm.Get("x")

		// Act
		actual := args.Map{"val": v}

		// Assert
		expected := args.Map{"val": "hello"}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateKeyValueAny", actual)
	})
}

func Test_Hashmap_AddOrUpdateHashmap_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateHashmap_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		result := hm.AddOrUpdateHashmap(nil)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- AddOrUpdateHashmap nil", actual)
	})
}

func Test_Hashmap_AddOrUpdateHashmap_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateHashmap", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("b", "2")
		hm.AddOrUpdateHashmap(other)

		// Act
		actual := args.Map{
			"len": hm.Length(),
			"hasB": hm.Has("b"),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"hasB": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateHashmap", actual)
	})
}

func Test_Hashmap_AddOrUpdateMap_Empty_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateMap_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdateMap(map[string]string{})

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- AddOrUpdateMap empty", actual)
	})
}

func Test_Hashmap_AddOrUpdateMap_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateMap", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdateMap(map[string]string{"x": "y"})

		// Act
		actual := args.Map{"has": hm.Has("x")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateMap", actual)
	})
}

func Test_Hashmap_AddsOrUpdates_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddsOrUpdates_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		result := hm.AddsOrUpdates(nil...)

		// Act
		actual := args.Map{"empty": result.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- AddsOrUpdates nil", actual)
	})
}

func Test_Hashmap_AddsOrUpdates_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddsOrUpdates", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddsOrUpdates(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddsOrUpdates", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyAnyValues_Empty_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateKeyAnyValues_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdateKeyAnyValues()

		// Act
		actual := args.Map{"empty": hm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- AddOrUpdateKeyAnyValues empty", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyValues_Empty_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateKeyValues_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdateKeyValues()

		// Act
		actual := args.Map{"empty": hm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- AddOrUpdateKeyValues empty", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyValues_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateKeyValues", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdateKeyValues(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Act
		actual := args.Map{"has": hm.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- AddOrUpdateKeyValues", actual)
	})
}

func Test_Hashmap_AddOrUpdateLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdateLock("k", "v")
		v, _ := hm.Get("k")

		// Act
		actual := args.Map{"val": v}

		// Assert
		expected := args.Map{"val": "v"}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateLock", actual)
	})
}

func Test_Hashmap_AddOrUpdateWithWgLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateWithWgLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hm.AddOrUpdateWithWgLock("k", "v", wg)
		wg.Wait()

		// Act
		actual := args.Map{"has": hm.Has("k")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- AddOrUpdateWithWgLock", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Has / Contains / IsKeyMissing
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_Has_Contains(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Has_Contains", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{
			"has": hm.Has("a"),
			"contains": hm.Contains("a"),
			"missing": hm.IsKeyMissing("b"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"contains": true,
			"missing": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Has/Contains", actual)
	})
}

func Test_Hashmap_ContainsLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ContainsLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{
			"cl": hm.ContainsLock("a"),
			"hl": hm.HasLock("a"),
			"hwl": hm.HasWithLock("a"),
			"mkl": hm.IsKeyMissingLock("z"),
		}

		// Assert
		expected := args.Map{
			"cl": true,
			"hl": true,
			"hwl": true,
			"mkl": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- lock variants", actual)
	})
}

func Test_Hashmap_HasAllStrings_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_HasAllStrings", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")

		// Act
		actual := args.Map{
			"all": hm.HasAllStrings("a", "b"),
			"missing": hm.HasAllStrings("a", "c"),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"missing": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- HasAllStrings", actual)
	})
}

func Test_Hashmap_HasAll_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_HasAll", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{
			"all": hm.HasAll("a"),
			"miss": hm.HasAll("a", "z"),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- HasAll", actual)
	})
}

func Test_Hashmap_HasAny_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_HasAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{
			"any": hm.HasAny("z", "a"),
			"none": hm.HasAny("x", "y"),
		}

		// Assert
		expected := args.Map{
			"any": true,
			"none": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- HasAny", actual)
	})
}

func Test_Hashmap_HasAllCollectionItems_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_HasAllCollectionItems", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		coll := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"all": hm.HasAllCollectionItems(coll),
			"nil": hm.HasAllCollectionItems(nil),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- HasAllCollectionItems", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Get / GetValue
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_Get_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Get", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		val, found := hm.Get("k")
		val2, found2 := hm.GetValue("k")
		_, notFound := hm.Get("missing")

		// Act
		actual := args.Map{
			"val": val,
			"found": found,
			"val2": val2,
			"found2": found2,
			"notFound": notFound,
		}

		// Assert
		expected := args.Map{
			"val": "v",
			"found": true,
			"val2": "v",
			"found2": true,
			"notFound": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Get/GetValue", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Items / SafeItems / Keys / Values
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_Items_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Items", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"len": len(hm.Items())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Items", actual)
	})
}

func Test_Hashmap_SafeItems_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_SafeItems_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{"nil": hm.SafeItems() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- SafeItems nil", actual)
	})
}

func Test_Hashmap_Keys_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Keys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		keys := hm.Keys()
		allKeys := hm.AllKeys()

		// Act
		actual := args.Map{
			"keysLen": len(keys),
			"allKeysLen": len(allKeys),
		}

		// Assert
		expected := args.Map{
			"keysLen": 2,
			"allKeysLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Keys", actual)
	})
}

func Test_Hashmap_KeysCollection_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_KeysCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		kc := hm.KeysCollection()

		// Act
		actual := args.Map{"len": kc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- KeysCollection", actual)
	})
}

func Test_Hashmap_ValuesList_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ValuesList", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		vl := hm.ValuesList()

		// Act
		actual := args.Map{
			"len": len(vl),
			"val": vl[0],
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"val": "1",
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- ValuesList", actual)
	})
}

func Test_Hashmap_ValuesCollection_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ValuesCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		vc := hm.ValuesCollection()

		// Act
		actual := args.Map{"len": vc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- ValuesCollection", actual)
	})
}

func Test_Hashmap_ValuesHashset_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ValuesHashset", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		vh := hm.ValuesHashset()

		// Act
		actual := args.Map{"has": vh.Has("1")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- ValuesHashset", actual)
	})
}

func Test_Hashmap_Collection_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Collection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		coll := hm.Collection()

		// Act
		actual := args.Map{"len": coll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Collection", actual)
	})
}

func Test_Hashmap_KeysValuesList_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_KeysValuesList", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		keys, vals := hm.KeysValuesList()

		// Act
		actual := args.Map{
			"kLen": len(keys),
			"vLen": len(vals),
		}

		// Assert
		expected := args.Map{
			"kLen": 1,
			"vLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- KeysValuesList", actual)
	})
}

func Test_Hashmap_KeysValuesCollection_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_KeysValuesCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		keys, vals := hm.KeysValuesCollection()

		// Act
		actual := args.Map{
			"kLen": keys.Length(),
			"vLen": vals.Length(),
		}

		// Assert
		expected := args.Map{
			"kLen": 1,
			"vLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- KeysValuesCollection", actual)
	})
}

func Test_Hashmap_KeysValuePairs_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_KeysValuePairs", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		pairs := hm.KeysValuePairs()

		// Act
		actual := args.Map{"len": len(pairs)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- KeysValuePairs", actual)
	})
}

func Test_Hashmap_KeysValuePairsCollection_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_KeysValuePairsCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		kvColl := hm.KeysValuePairsCollection()

		// Act
		actual := args.Map{"len": kvColl.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- KeysValuePairsCollection", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Lock variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_IsEmptyLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_IsEmptyLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"empty": hm.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- IsEmptyLock", actual)
	})
}

func Test_Hashmap_LengthLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_LengthLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"len": hm.LengthLock()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- LengthLock", actual)
	})
}

func Test_Hashmap_KeysLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_KeysLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		kl := hm.KeysLock()

		// Act
		actual := args.Map{"len": len(kl)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- KeysLock", actual)
	})
}

func Test_Hashmap_ValuesListCopyLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ValuesListCopyLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		vl := hm.ValuesListCopyLock()

		// Act
		actual := args.Map{"len": len(vl)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- ValuesListCopyLock", actual)
	})
}

func Test_Hashmap_ValuesCollectionLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ValuesCollectionLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		vc := hm.ValuesCollectionLock()

		// Act
		actual := args.Map{"len": vc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- ValuesCollectionLock", actual)
	})
}

func Test_Hashmap_ValuesHashsetLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ValuesHashsetLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		vh := hm.ValuesHashsetLock()

		// Act
		actual := args.Map{"has": vh.Has("1")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- ValuesHashsetLock", actual)
	})
}

func Test_Hashmap_ItemsCopyLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ItemsCopyLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		cp := hm.ItemsCopyLock()

		// Act
		actual := args.Map{"len": len(*cp)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- ItemsCopyLock", actual)
	})
}

func Test_Hashmap_KeysValuesListLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_KeysValuesListLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		keys, vals := hm.KeysValuesListLock()

		// Act
		actual := args.Map{
			"kLen": len(keys),
			"vLen": len(vals),
		}

		// Assert
		expected := args.Map{
			"kLen": 1,
			"vLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- KeysValuesListLock", actual)
	})
}

func Test_Hashmap_StringLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_StringLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		s := hm.StringLock()

		// Act
		actual := args.Map{"notEmpty": s != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- StringLock", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Diff / ConcatNew
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_ConcatNew_NoArgs_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ConcatNew_NoArgs", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		c := hm.ConcatNew(true)

		// Act
		actual := args.Map{
			"len": c.Length(),
			"has": c.Has("a"),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"has": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- ConcatNew no args", actual)
	})
}

func Test_Hashmap_ConcatNew_WithArgs_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ConcatNew_WithArgs", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("b", "2")
		c := hm.ConcatNew(true, other)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- ConcatNew with args", actual)
	})
}

func Test_Hashmap_ConcatNew_NilInArgs(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ConcatNew_NilInArgs", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		c := hm.ConcatNew(true, nil)

		// Act
		actual := args.Map{"has": c.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- ConcatNew nil in args", actual)
	})
}

func Test_Hashmap_ConcatNewUsingMaps_NoArgs_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ConcatNewUsingMaps_NoArgs", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		c := hm.ConcatNewUsingMaps(true)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- ConcatNewUsingMaps no args", actual)
	})
}

func Test_Hashmap_ConcatNewUsingMaps_WithArgs(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ConcatNewUsingMaps_WithArgs", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		c := hm.ConcatNewUsingMaps(true, map[string]string{"b": "2"}, nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- ConcatNewUsingMaps with args", actual)
	})
}

func Test_Hashmap_Diff_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Diff", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("a", "1")
		other.AddOrUpdate("b", "99")
		diff := hm.Diff(other)

		// Act
		actual := args.Map{"hasDiff": diff.HasAnyItem()}

		// Assert
		expected := args.Map{"hasDiff": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Diff", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Remove / Clear / Dispose
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_Remove_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Remove", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.Remove("a")

		// Act
		actual := args.Map{"has": hm.Has("a")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Remove", actual)
	})
}

func Test_Hashmap_RemoveWithLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_RemoveWithLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.RemoveWithLock("a")

		// Act
		actual := args.Map{"has": hm.Has("a")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- RemoveWithLock", actual)
	})
}

func Test_Hashmap_Clear_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Clear", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.Clear()

		// Act
		actual := args.Map{"empty": hm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Clear", actual)
	})
}

func Test_Hashmap_Clear_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Clear_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap
		result := hm.Clear()

		// Act
		actual := args.Map{"nil": result == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- Clear nil", actual)
	})
}

func Test_Hashmap_Dispose_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Dispose", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.Dispose()

		// Act
		actual := args.Map{"empty": hm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Dispose", actual)
	})
}

func Test_Hashmap_Dispose_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Dispose_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap
		hm.Dispose() // should not panic

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- Dispose nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — IsEqual / Clone
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_IsEqualPtr_Same(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_IsEqualPtr_Same", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"same": hm.IsEqualPtr(hm)}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- IsEqualPtr same ptr", actual)
	})
}

func Test_Hashmap_IsEqualPtr_BothNil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_IsEqualPtr_BothNil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{"eq": hm.IsEqualPtr(nil)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- IsEqualPtr both nil", actual)
	})
}

func Test_Hashmap_IsEqualPtr_OneNil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_IsEqualPtr_OneNil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"eq": hm.IsEqualPtr(nil)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- IsEqualPtr one nil", actual)
	})
}

func Test_Hashmap_IsEqualPtr_BothEmpty_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_IsEqualPtr_BothEmpty", func() {
		// Arrange
		hm1 := corestr.New.Hashmap.Cap(5)
		hm2 := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"eq": hm1.IsEqualPtr(hm2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- IsEqualPtr both empty", actual)
	})
}

func Test_Hashmap_IsEqualPtr_DiffLength(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_IsEqualPtr_DiffLength", func() {
		// Arrange
		hm1 := corestr.New.Hashmap.Cap(5)
		hm1.AddOrUpdate("a", "1")
		hm2 := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"eq": hm1.IsEqualPtr(hm2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- IsEqualPtr diff length", actual)
	})
}

func Test_Hashmap_IsEqualPtr_DiffValues(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_IsEqualPtr_DiffValues", func() {
		// Arrange
		hm1 := corestr.New.Hashmap.Cap(5)
		hm1.AddOrUpdate("a", "1")
		hm2 := corestr.New.Hashmap.Cap(5)
		hm2.AddOrUpdate("a", "2")

		// Act
		actual := args.Map{"eq": hm1.IsEqualPtr(hm2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- IsEqualPtr diff values", actual)
	})
}

func Test_Hashmap_IsEqual_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_IsEqual", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		cloned := hm.Clone()

		// Act
		actual := args.Map{"eq": hm.IsEqual(cloned)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- IsEqual", actual)
	})
}

func Test_Hashmap_IsEqualPtrLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_IsEqualPtrLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"eq": hm.IsEqualPtrLock(hm)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- IsEqualPtrLock", actual)
	})
}

func Test_Hashmap_Clone_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Clone", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		c := hm.Clone()

		// Act
		actual := args.Map{
			"len": c.Length(),
			"has": c.Has("a"),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"has": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Clone", actual)
	})
}

func Test_Hashmap_ClonePtr_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ClonePtr", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		cp := hm.ClonePtr()

		// Act
		actual := args.Map{
			"notNil": cp != nil,
			"has": cp.Has("a"),
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"has": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- ClonePtr", actual)
	})
}

func Test_Hashmap_ClonePtr_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ClonePtr_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap
		cp := hm.ClonePtr()

		// Act
		actual := args.Map{"nil": cp == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- ClonePtr nil", actual)
	})
}

func Test_Hashmap_Clone_Empty_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Clone_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		c := hm.Clone()

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- Clone empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — String / Join / KeysToLower
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_String_Empty_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_String_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"notEmpty": hm.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- String empty", actual)
	})
}

func Test_Hashmap_String_WithItems(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_String_WithItems", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"notEmpty": hm.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- String with items", actual)
	})
}

func Test_Hashmap_Join_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Join", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"val": hm.Join(",")}

		// Assert
		expected := args.Map{"val": "1"}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Join", actual)
	})
}

func Test_Hashmap_JoinKeys_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_JoinKeys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"val": hm.JoinKeys(",")}

		// Assert
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- JoinKeys", actual)
	})
}

func Test_Hashmap_KeysToLower_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_KeysToLower", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("ABC", "1")
		lower := hm.KeysToLower()

		// Act
		actual := args.Map{"has": lower.Has("abc")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- KeysToLower", actual)
	})
}

func Test_Hashmap_ValuesToLower_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ValuesToLower", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("ABC", "1")
		lower := hm.ValuesToLower()

		// Act
		actual := args.Map{"has": lower.Has("abc")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- ValuesToLower deprecated", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — JSON / Serialize / Deserialize
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_JsonModel_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_JsonModel", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		jm := hm.JsonModel()

		// Act
		actual := args.Map{"len": len(jm)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- JsonModel", actual)
	})
}

func Test_Hashmap_JsonModelAny_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_JsonModelAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"notNil": hm.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- JsonModelAny", actual)
	})
}

func Test_Hashmap_MarshalJSON_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_MarshalJSON", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		b, err := hm.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- MarshalJSON", actual)
	})
}

func Test_Hashmap_UnmarshalJSON_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_UnmarshalJSON", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		err := hm.UnmarshalJSON([]byte(`{"x":"y"}`))
		v, _ := hm.Get("x")

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"val": v,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"val": "y",
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- UnmarshalJSON", actual)
	})
}

func Test_Hashmap_UnmarshalJSON_Err(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_UnmarshalJSON_Err", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		err := hm.UnmarshalJSON([]byte(`{invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns error -- UnmarshalJSON err", actual)
	})
}


func Test_Hashmap_Json_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Json", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		j := hm.Json()
		actual := args.Map{"hasBytes": j.HasBytes()}
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Json", actual)
	})
}

func Test_Hashmap_JsonPtr_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_JsonPtr", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		jp := hm.JsonPtr()
		actual := args.Map{"notNil": jp != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- JsonPtr", actual)
	})
}

func Test_Hashmap_Serialize_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Serialize", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		b, err := hm.Serialize()
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Serialize", actual)
	})
}

func Test_Hashmap_Deserialize_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_Deserialize", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		target := map[string]string{}
		err := hm.Deserialize(&target)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Deserialize", actual)
	})
}

func Test_Hashmap_ParseInjectUsingJson_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ParseInjectUsingJson", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		jr := hm.JsonPtr()
		hm2 := corestr.New.Hashmap.Cap(5)
		result, err := hm2.ParseInjectUsingJson(jr)
		actual := args.Map{
			"noErr": err == nil,
			"has": result.Has("a"),
		}
		expected := args.Map{
			"noErr": true,
			"has": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- ParseInjectUsingJson", actual)
	})
}

func Test_Hashmap_ParseInjectUsingJson_Err(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ParseInjectUsingJson_Err", func() {
		hm := corestr.New.Hashmap.Cap(5)
		badJson := corejson.NewPtr(42) // not a map
		_, err := hm.ParseInjectUsingJson(badJson)
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns error -- ParseInjectUsingJson err", actual)
	})
}

func Test_Hashmap_JsonParseSelfInject_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_JsonParseSelfInject", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		jr := hm.JsonPtr()
		hm2 := corestr.New.Hashmap.Cap(5)
		err := hm2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- JsonParseSelfInject", actual)
	})
}

func Test_Hashmap_AsJsoner_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AsJsoner", func() {
		hm := corestr.New.Hashmap.Cap(5)
		actual := args.Map{"notNil": hm.AsJsoner() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AsJsoner", actual)
	})
}

func Test_Hashmap_AsJsonContractsBinder_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AsJsonContractsBinder", func() {
		hm := corestr.New.Hashmap.Cap(5)
		actual := args.Map{"notNil": hm.AsJsonContractsBinder() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AsJsonContractsBinder", actual)
	})
}

func Test_Hashmap_AsJsonParseSelfInjector_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AsJsonParseSelfInjector", func() {
		hm := corestr.New.Hashmap.Cap(5)
		actual := args.Map{"notNil": hm.AsJsonParseSelfInjector() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AsJsonParseSelfInjector", actual)
	})
}

func Test_Hashmap_AsJsonMarshaller_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AsJsonMarshaller", func() {
		hm := corestr.New.Hashmap.Cap(5)
		actual := args.Map{"notNil": hm.AsJsonMarshaller() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AsJsonMarshaller", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — ToError / KeyValStringLines / ToStringsUsingCompiler
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_ToError_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ToError", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		err := hm.ToError(", ")
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns error -- ToError", actual)
	})
}

func Test_Hashmap_ToDefaultError_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ToDefaultError", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		err := hm.ToDefaultError()
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns error -- ToDefaultError", actual)
	})
}

func Test_Hashmap_KeyValStringLines_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_KeyValStringLines", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		lines := hm.KeyValStringLines()
		actual := args.Map{"len": len(lines)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- KeyValStringLines", actual)
	})
}

func Test_Hashmap_ToStringsUsingCompiler_Empty_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_ToStringsUsingCompiler_Empty", func() {
		hm := corestr.New.Hashmap.Cap(5)
		result := hm.ToStringsUsingCompiler(func(k, v string) string { return k + v })
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- ToStringsUsingCompiler empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — GetValuesExcept / GetAllExcept / Filter
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_GetValuesExceptKeysInHashset_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_GetValuesExceptKeysInHashset", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		exclude := corestr.New.Hashset.Strings([]string{"a"})
		result := hm.GetValuesExceptKeysInHashset(exclude)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- GetValuesExceptKeysInHashset", actual)
	})
}

func Test_Hashmap_GetValuesExceptKeysInHashset_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_GetValuesExceptKeysInHashset_Nil", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		result := hm.GetValuesExceptKeysInHashset(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- GetValuesExceptKeysInHashset nil", actual)
	})
}

func Test_Hashmap_GetValuesKeysExcept_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_GetValuesKeysExcept", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		result := hm.GetValuesKeysExcept([]string{"a"})
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- GetValuesKeysExcept", actual)
	})
}

func Test_Hashmap_GetValuesKeysExcept_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_GetValuesKeysExcept_Nil", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		result := hm.GetValuesKeysExcept(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- GetValuesKeysExcept nil", actual)
	})
}

func Test_Hashmap_GetAllExceptCollection_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_GetAllExceptCollection", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		coll := corestr.New.Collection.Strings([]string{"a"})
		result := hm.GetAllExceptCollection(coll)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- GetAllExceptCollection", actual)
	})
}

func Test_Hashmap_GetAllExceptCollection_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_GetAllExceptCollection_Nil", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		result := hm.GetAllExceptCollection(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- GetAllExceptCollection nil", actual)
	})
}

func Test_Hashmap_GetKeysFilteredItems_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_GetKeysFilteredItems", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("abc", "1")
		hm.AddOrUpdate("xyz", "2")
		result := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 3, false
		})
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- GetKeysFilteredItems", actual)
	})
}

func Test_Hashmap_GetKeysFilteredItems_Empty_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_GetKeysFilteredItems_Empty", func() {
		hm := corestr.New.Hashmap.Cap(5)
		result := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- GetKeysFilteredItems empty", actual)
	})
}

func Test_Hashmap_GetKeysFilteredCollection_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_GetKeysFilteredCollection", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		coll := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"len": coll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- GetKeysFilteredCollection", actual)
	})
}

func Test_Hashmap_GetKeysFilteredCollection_Empty_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_GetKeysFilteredCollection_Empty", func() {
		hm := corestr.New.Hashmap.Cap(5)
		coll := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"empty": coll.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- GetKeysFilteredCollection empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — AddOrUpdateCollection / AddOrUpdateStringsPtrWgLock
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_AddOrUpdateCollection_Nil(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateCollection_Nil", func() {
		hm := corestr.New.Hashmap.Cap(5)
		result := hm.AddOrUpdateCollection(nil, nil)
		actual := args.Map{"empty": result.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- AddOrUpdateCollection nil", actual)
	})
}

func Test_Hashmap_AddOrUpdateCollection_Mismatch_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateCollection_Mismatch", func() {
		hm := corestr.New.Hashmap.Cap(5)
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1"})
		result := hm.AddOrUpdateCollection(keys, vals)
		actual := args.Map{"empty": result.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateCollection mismatch", actual)
	})
}

func Test_Hashmap_AddOrUpdateCollection_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateCollection", func() {
		hm := corestr.New.Hashmap.Cap(5)
		keys := corestr.New.Collection.Strings([]string{"a"})
		vals := corestr.New.Collection.Strings([]string{"1"})
		hm.AddOrUpdateCollection(keys, vals)
		v, _ := hm.Get("a")
		actual := args.Map{"val": v}
		expected := args.Map{"val": "1"}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateCollection", actual)
	})
}

func Test_Hashmap_AddOrUpdateStringsPtrWgLock_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateStringsPtrWgLock", func() {
		hm := corestr.New.Hashmap.Cap(10)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{"a", "b"}, []string{"1", "2"})
		wg.Wait()
		actual := args.Map{"len": hm.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateStringsPtrWgLock", actual)
	})
}

func Test_Hashmap_AddOrUpdateStringsPtrWgLock_Empty_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddOrUpdateStringsPtrWgLock_Empty", func() {
		hm := corestr.New.Hashmap.Cap(5)
		result := hm.AddOrUpdateStringsPtrWgLock(&sync.WaitGroup{}, []string{}, []string{})
		actual := args.Map{"empty": result.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- AddOrUpdateStringsPtrWgLock empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Filter variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_AddsOrUpdatesAnyUsingFilter_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddsOrUpdatesAnyUsingFilter_Nil", func() {
		hm := corestr.New.Hashmap.Cap(5)
		result := hm.AddsOrUpdatesAnyUsingFilter(nil, nil...)
		actual := args.Map{"empty": result.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- AddsOrUpdatesAnyUsingFilter nil", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilter_Break_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddsOrUpdatesAnyUsingFilter_Break", func() {
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(p corestr.KeyAnyValuePair) (string, bool, bool) {
			return p.ValueString(), true, true
		}
		hm.AddsOrUpdatesAnyUsingFilter(filter,
			corestr.KeyAnyValuePair{Key: "a", Value: "1"},
			corestr.KeyAnyValuePair{Key: "b", Value: "2"},
		)
		actual := args.Map{"len": hm.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddsOrUpdatesAnyUsingFilter break", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Nil", func() {
		hm := corestr.New.Hashmap.Cap(5)
		result := hm.AddsOrUpdatesAnyUsingFilterLock(nil, nil...)
		actual := args.Map{"empty": result.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- AddsOrUpdatesAnyUsingFilterLock nil", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Break_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Break", func() {
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(p corestr.KeyAnyValuePair) (string, bool, bool) {
			return p.ValueString(), true, true
		}
		hm.AddsOrUpdatesAnyUsingFilterLock(filter,
			corestr.KeyAnyValuePair{Key: "a", Value: "1"},
			corestr.KeyAnyValuePair{Key: "b", Value: "2"},
		)
		actual := args.Map{"len": hm.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddsOrUpdatesAnyUsingFilterLock break", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesUsingFilter_Nil_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddsOrUpdatesUsingFilter_Nil", func() {
		hm := corestr.New.Hashmap.Cap(5)
		result := hm.AddsOrUpdatesUsingFilter(nil, nil...)
		actual := args.Map{"empty": result.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- AddsOrUpdatesUsingFilter nil", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_AddsOrUpdatesUsingFilter_Break", func() {
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(p corestr.KeyValuePair) (string, bool, bool) {
			return p.Value, true, true
		}
		hm.AddsOrUpdatesUsingFilter(filter,
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)
		actual := args.Map{"len": hm.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddsOrUpdatesUsingFilter break", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — GetKeysFilteredItems/Collection with break
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_GetKeysFilteredItems_Break_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_GetKeysFilteredItems_Break", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		result := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, true // break on first
		})
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- GetKeysFilteredItems break", actual)
	})
}

func Test_Hashmap_GetKeysFilteredCollection_Break_FromHashmapIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I29_Hashmap_GetKeysFilteredCollection_Break", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		coll := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, true // break on first
		})
		actual := args.Map{"len": coll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- GetKeysFilteredCollection break", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashsetDataModel
// ══════════════════════════════════════════════════════════════════════════════

func Test_HashsetDataModel_NewUsing(t *testing.T) {
	safeTest(t, "Test_I29_HashsetDataModel_NewUsing", func() {
		dm := &corestr.HashsetDataModel{Items: map[string]bool{"a": true}}
		hs := corestr.NewHashsetUsingDataModel(dm)
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HashsetDataModel returns correct value -- NewUsing", actual)
	})
}

func Test_HashsetDataModel_NewFromCollection(t *testing.T) {
	safeTest(t, "Test_I29_HashsetDataModel_NewFromCollection", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		dm := corestr.NewHashsetsDataModelUsing(hs)
		actual := args.Map{
			"notNil": dm != nil,
			"len": len(dm.Items),
		}
		expected := args.Map{
			"notNil": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "HashsetDataModel returns correct value -- NewFromCollection", actual)
	})
}
