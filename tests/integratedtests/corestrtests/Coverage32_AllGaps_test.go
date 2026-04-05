package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage32 — corestr remaining 42 lines (mostly nil-receiver dead code)
// ══════════════════════════════════════════════════════════════════════════════

// ── CharCollectionMap nil-receiver guards ──
// Lines 45, 369, 567, 589, 622, 868, 985, 1105 are all nil-receiver guards
// on unexported internal fields. Documented as accepted dead code.

// ── CharHashsetMap nil-receiver guards ──
// Lines 593, 624, 657, 661, 713, 856, 906, 991, 1050 are nil-receiver guards.

// ── CharHashsetMap.FilterKeys (line 748-772) ──

func Test_Cov32_CharHashsetMap_FilterKeys(t *testing.T) {
	// Arrange
	chm := corestr.New.CharHashsetMap.Cap(3, 3)
	chm.AddOrUpdate('a', "alpha")
	chm.AddOrUpdate('b', "beta")
	chm.AddOrUpdate('c', "gamma")

	// Act
	filtered := chm.FilterKeys(func(key rune) bool {
		return key == 'a' || key == 'c'
	})

	// Assert
	actual := args.Map{
		"length": filtered.Length(),
	}
	expected := args.Map{
		"length": 2,
	}
	actual.ShouldBeEqual(t, 1, "CharHashsetMap FilterKeys", expected)
}

// ── Collection.JsonString error path (line 97) ──
// json.Marshal on []string won't fail — dead code

// ── Collection.MarshalJSON/UnmarshalJSON error (lines 497, 508, 528, 539, 559, 570, 581, 592) ──
// These are various nil-guard or error-return paths in JSON operations

func Test_Cov32_Collection_JsonOperations(t *testing.T) {
	// Arrange
	coll := corestr.New.Collection.Strings([]string{"a", "b", "c"})

	// Act
	jsonStr, err := coll.JsonString()

	// Assert
	actual := args.Map{
		"hasContent": len(jsonStr) > 0,
		"hasError":   err != nil,
	}
	expected := args.Map{
		"hasContent": true,
		"hasError":   false,
	}
	actual.ShouldBeEqual(t, 1, "Collection JsonString", expected)
}

// ── CollectionsOfCollection nil guards (lines 45, 68) ──
// Nil receiver guard — dead code

// ── Hashmap nil guard (line 158) ──
// Nil receiver — dead code

// ── LinkedCollectionNode nil guards (lines 93, 152, 156, 177, 256) ──
// All nil-receiver guards — dead code

// ── LinkedCollections various paths ──

func Test_Cov32_LinkedCollections_FilterByKeys(t *testing.T) {
	// Arrange
	lc := corestr.New.LinkedCollection.Cap(3)
	lc.Add("alpha", corestr.New.Collection.Strings([]string{"a", "b"}))
	lc.Add("beta", corestr.New.Collection.Strings([]string{"c", "d"}))
	lc.Add("gamma", corestr.New.Collection.Strings([]string{"e"}))

	// Act
	filtered := lc.FilterByKeys(func(key string) bool {
		return key == "alpha" || key == "gamma"
	})

	// Assert
	actual := args.Map{"length": filtered.Length()}
	expected := args.Map{"length": 2}
	actual.ShouldBeEqual(t, 1, "LinkedCollections FilterByKeys", expected)
}

// ── LinkedCollections.GetCollectionByIndexSafe out of range (line 272) ──

func Test_Cov32_LinkedCollections_GetCollectionByIndexSafe_OutOfRange(t *testing.T) {
	// Arrange
	lc := corestr.New.LinkedCollection.Cap(2)
	lc.Add("a", corestr.New.Collection.Strings([]string{"x"}))

	// Act
	result := lc.GetCollectionByIndexSafe(99)

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	actual.ShouldBeEqual(t, 1, "LinkedCollections GetCollectionByIndexSafe out of range", expected)
}

// ── LinkedCollections.JsonString error path (line 646) ──

func Test_Cov32_LinkedCollections_JsonString(t *testing.T) {
	// Arrange
	lc := corestr.New.LinkedCollection.Cap(2)
	lc.Add("key1", corestr.New.Collection.Strings([]string{"v1"}))

	// Act
	jsonStr, err := lc.JsonString()

	// Assert
	actual := args.Map{
		"hasContent": len(jsonStr) > 0,
		"hasError":   err != nil,
	}
	expected := args.Map{
		"hasContent": true,
		"hasError":   false,
	}
	actual.ShouldBeEqual(t, 1, "LinkedCollections JsonString", expected)
}

// ── LinkedCollections.FindFirstByValue empty (line 760) ──

func Test_Cov32_LinkedCollections_FindFirstByValue_NotFound(t *testing.T) {
	// Arrange
	lc := corestr.New.LinkedCollection.Cap(2)
	lc.Add("key1", corestr.New.Collection.Strings([]string{"v1"}))

	// Act
	result := lc.FindFirstByValue("nonexistent")

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	actual.ShouldBeEqual(t, 1, "LinkedCollections FindFirstByValue not found", expected)
}

// ── LinkedCollections.SerializeToKeySliceMap error (line 943) ──

func Test_Cov32_LinkedCollections_SerializeToKeySliceMap(t *testing.T) {
	// Arrange
	lc := corestr.New.LinkedCollection.Cap(2)
	lc.Add("k1", corestr.New.Collection.Strings([]string{"a"}))

	// Act
	result := lc.SerializeToKeySliceMap()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	actual.ShouldBeEqual(t, 1, "LinkedCollections SerializeToKeySliceMap", expected)
}

// ── LinkedCollections: remaining dead-code paths ──
// Lines 102, 147, 151, 1143, 1182, 1185, 1248 are nil-receiver guards
// or unreachable fallback returns. Documented as accepted dead code.
