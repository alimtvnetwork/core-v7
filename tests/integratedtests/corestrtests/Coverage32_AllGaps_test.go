package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage32 — corestr remaining ~42 lines (mostly nil-receiver dead code)
// ══════════════════════════════════════════════════════════════════════════════

// ── CharCollectionMap nil-receiver guards ──
// Lines 45, 369, 567, 589, 622, 868, 985, 1105 — documented as dead code

// ── CharHashsetMap nil-receiver guards ──
// Lines 593, 624, 657, 661, 713, 856, 906, 991, 1050 — documented as dead code

// ── CharHashsetMap.efficientAddOfLargeItems (line 748-772) ──

func Test_Cov32_CharHashsetMap_EfficientAdd_I29(t *testing.T) {
	// Arrange
	chm := corestr.New.CharHashsetMap.Cap(3, 3)

	// Act — add large strings to trigger efficient add path
	chm.AddStrings("alpha", "beta", "gamma", "delta", "epsilon")

	// Assert
	actual := args.Map{
		"hasItems": chm.Length() > 0,
	}
	expected := args.Map{
		"hasItems": true,
	}
	actual.ShouldBeEqual(t, 1, "CharHashsetMap efficient add", expected)
}

// ── Collection.JsonString error path (line 97) ──
// json.Marshal on []string won't fail — dead code

// ── Collection JSON operations (lines 497, 508, 528, 539, 559, 570, 581, 592) ──
// Various nil/error return paths in JSON — mostly dead code on valid data

func Test_Cov32_Collection_JsonString_I29(t *testing.T) {
	// Arrange
	coll := corestr.New.Collection.Strings([]string{"a", "b", "c"})

	// Act
	jsonStr := coll.JsonString()

	// Assert
	actual := args.Map{
		"hasContent": len(jsonStr) > 0,
	}
	expected := args.Map{
		"hasContent": true,
	}
	actual.ShouldBeEqual(t, 1, "Collection JsonString valid", expected)
}

// ── CollectionsOfCollection nil guards (lines 45, 68) ──
// Nil receiver guard — dead code

// ── Hashmap nil guard (line 158) — dead code ──

// ── LinkedCollectionNode nil guards (lines 93, 152, 156, 177, 256) — dead code ──

// ── LinkedCollections various paths ──

func Test_Cov32_LinkedCollections_SafePointerIndexAt_OutOfRange_I29(t *testing.T) {
	// Arrange
	lc := corestr.New.LinkedCollection.Create()
	lc.Add("a", corestr.New.Collection.Strings([]string{"x"}))

	// Act
	result := lc.SafePointerIndexAt(99)

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	actual.ShouldBeEqual(t, 1, "LinkedCollections SafePointerIndexAt out of range", expected)
}

// ── LinkedCollections.FlattenToSingleCollection (line 646) ──

func Test_Cov32_LinkedCollections_FlattenToSingleCollection_I29(t *testing.T) {
	// Arrange
	lc := corestr.New.LinkedCollection.Create()
	lc.Add("key1", corestr.New.Collection.Strings([]string{"v1", "v2"}))
	lc.Add("key2", corestr.New.Collection.Strings([]string{"v3"}))

	// Act
	flattened := lc.FlattenToSingleCollection(0)

	// Assert
	actual := args.Map{"length": flattened.Length()}
	expected := args.Map{"length": 3}
	actual.ShouldBeEqual(t, 1, "LinkedCollections FlattenToSingleCollection", expected)
}

// ── LinkedCollections: remaining dead-code paths ──
// Lines 102, 147, 151, 760, 943, 1143, 1182, 1185, 1248 are nil-receiver
// guards or unreachable fallback returns. Documented as accepted dead code.
