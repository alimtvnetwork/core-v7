package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: EmptyHashmap
// ==========================================================================

func Test_Hashmap_Empty(t *testing.T) {
	// Arrange
	tc := hashmapEmptyTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{
		"isEmpty":  hm.IsEmpty(),
		"length":   hm.Length(),
		"hasItems": hm.HasItems(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: NewHashmap
// ==========================================================================

func Test_Hashmap_New(t *testing.T) {
	// Arrange
	tc := hashmapNewTestCase
	hm := coregeneric.NewHashmap[string, int](10)

	// Act
	actual := args.Map{"isEmpty": hm.IsEmpty()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: HashmapFrom
// ==========================================================================

func Test_Hashmap_From(t *testing.T) {
	// Arrange
	tc := hashmapFromTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})

	// Act
	actual := args.Map{
		"length": hm.Length(),
		"hasKey": hm.Has("a"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: HashmapClone (function)
// ==========================================================================

func Test_Hashmap_CloneFunc(t *testing.T) {
	// Arrange
	tc := hashmapCloneFuncTestCase
	orig := coregeneric.HashmapFrom(map[string]int{"k": 1})
	cloned := coregeneric.HashmapClone(orig.Map())
	cloned.Set("k", 99)

	origVal, _ := orig.Get("k")
	clonedVal, _ := cloned.Get("k")

	// Act
	actual := args.Map{
		"origValue":   origVal,
		"clonedValue": clonedVal,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Set
// ==========================================================================

func Test_Hashmap_SetNew(t *testing.T) {
	// Arrange
	tc := hashmapSetNewTestCase
	hm := coregeneric.EmptyHashmap[string, int]()
	isNew := hm.Set("key", 42)

	// Act
	actual := args.Map{
		"isNew":  isNew,
		"length": hm.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_SetExisting(t *testing.T) {
	// Arrange
	tc := hashmapSetExistingTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"key": 1})
	isNew := hm.Set("key", 2)
	val, _ := hm.Get("key")

	// Act
	actual := args.Map{
		"isNew":        isNew,
		"updatedValue": val,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Get
// ==========================================================================

func Test_Hashmap_GetFound(t *testing.T) {
	// Arrange
	tc := hashmapGetFoundTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"k": 42})
	val, found := hm.Get("k")

	// Act
	actual := args.Map{
		"found": found,
		"value": val,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_GetNotFound(t *testing.T) {
	// Arrange
	tc := hashmapGetNotFoundTestCase
	hm := coregeneric.EmptyHashmap[string, int]()
	val, found := hm.Get("missing")

	// Act
	actual := args.Map{
		"found": found,
		"value": val,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: GetOrDefault
// ==========================================================================

func Test_Hashmap_GetOrDefaultMissing(t *testing.T) {
	// Arrange
	tc := hashmapGetOrDefaultMissingTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"value": hm.GetOrDefault("x", 99)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_GetOrDefaultFound(t *testing.T) {
	// Arrange
	tc := hashmapGetOrDefaultFoundTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"x": 5})

	// Act
	actual := args.Map{"value": hm.GetOrDefault("x", 99)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Has / Contains / IsKeyMissing
// ==========================================================================

func Test_Hashmap_Has(t *testing.T) {
	// Arrange
	tc := hashmapHasTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"has":          hm.Has("a"),
		"contains":     hm.Contains("a"),
		"isKeyMissing": hm.IsKeyMissing("a"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsKeyMissing_FromHashmap(t *testing.T) {
	// Arrange
	tc := hashmapIsKeyMissingTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"isKeyMissing": hm.IsKeyMissing("x")}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Remove
// ==========================================================================

func Test_Hashmap_RemoveExisting(t *testing.T) {
	// Arrange
	tc := hashmapRemoveExistingTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"k": 1})
	existed := hm.Remove("k")

	// Act
	actual := args.Map{
		"removed": existed,
		"isGone":  hm.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_RemoveMissing(t *testing.T) {
	// Arrange
	tc := hashmapRemoveMissingTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"removed": hm.Remove("x")}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Keys
// ==========================================================================

func Test_Hashmap_Keys_NonEmpty(t *testing.T) {
	// Arrange
	tc := hashmapKeysNonEmptyTestCase
	hm := coregeneric.HashmapFrom(map[int]string{1: "a", 2: "b"})

	// Act
	actual := args.Map{"keyCount": len(hm.Keys())}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_Keys_Empty(t *testing.T) {
	// Arrange
	tc := hashmapKeysEmptyTestCase
	hm := coregeneric.EmptyHashmap[int, string]()

	// Act
	actual := args.Map{"keyCount": len(hm.Keys())}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Values
// ==========================================================================

func Test_Hashmap_Values_NonEmpty(t *testing.T) {
	// Arrange
	tc := hashmapValuesNonEmptyTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	vals := hm.Values()

	// Act
	actual := args.Map{
		"valueCount":       len(vals),
		"containsExpected": vals[0],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_Values_Empty(t *testing.T) {
	// Arrange
	tc := hashmapValuesEmptyTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"valueCount": len(hm.Values())}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: AddOrUpdateMap
// ==========================================================================

func Test_Hashmap_AddOrUpdateMap_Merges(t *testing.T) {
	// Arrange
	tc := hashmapAddOrUpdateMapMergesTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateMap(map[string]int{"b": 2, "a": 10})
	val, _ := hm.Get("a")

	// Act
	actual := args.Map{
		"length":      hm.Length(),
		"mergedValue": val,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_AddOrUpdateMap_EmptyNoop(t *testing.T) {
	// Arrange
	tc := hashmapAddOrUpdateMapEmptyNoopTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateMap(map[string]int{})

	// Act
	actual := args.Map{"length": hm.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: AddOrUpdateHashmap
// ==========================================================================

func Test_Hashmap_AddOrUpdateHashmap_Merges(t *testing.T) {
	// Arrange
	tc := hashmapAddOrUpdateHashmapMergesTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateHashmap(coregeneric.HashmapFrom(map[string]int{"b": 2}))

	// Act
	actual := args.Map{"length": hm.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_AddOrUpdateHashmap_NilNoop(t *testing.T) {
	// Arrange
	tc := hashmapAddOrUpdateHashmapNilNoopTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateHashmap(nil)

	// Act
	actual := args.Map{"length": hm.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: ConcatNew
// ==========================================================================

func Test_Hashmap_ConcatNew_Merged(t *testing.T) {
	// Arrange
	tc := hashmapConcatNewMergedTestCase
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"b": 2})
	result := hm1.ConcatNew(hm2)

	// Act
	actual := args.Map{
		"mergedLength":   result.Length(),
		"originalLength": hm1.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_ConcatNew_Nil(t *testing.T) {
	// Arrange
	tc := hashmapConcatNewNilTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	result := hm.ConcatNew(nil)

	// Act
	actual := args.Map{"length": result.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Clone method
// ==========================================================================

func Test_Hashmap_CloneMethod(t *testing.T) {
	// Arrange
	tc := hashmapCloneMethodTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"k": 1})
	cloned := hm.Clone()
	cloned.Set("k", 99)
	origVal, _ := hm.Get("k")

	// Act
	actual := args.Map{"origValue": origVal}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: IsEquals
// ==========================================================================

func Test_Hashmap_IsEquals_SameContent(t *testing.T) {
	// Arrange
	tc := hashmapIsEqualsSameContentTestCase
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
	hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})

	// Act
	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_DifferentKeys(t *testing.T) {
	// Arrange
	tc := hashmapIsEqualsDifferentKeysTestCase
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"b": 1})

	// Act
	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_DifferentLength(t *testing.T) {
	// Arrange
	tc := hashmapIsEqualsDifferentLengthTestCase
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})

	// Act
	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// Test_Hashmap_IsEquals_BothNil is defined in CollectionBranch_test.go (line 384).
// Removed duplicate declaration here.

func Test_Hashmap_IsEquals_OneNil(t *testing.T) {
	// Arrange
	tc := hashmapIsEqualsOneNilTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"isEquals": hm.IsEquals(nil)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_SamePointer(t *testing.T) {
	// Arrange
	tc := hashmapIsEqualsSamePointerTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	// Act
	actual := args.Map{"isEquals": hm.IsEquals(hm)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: ForEach
// ==========================================================================

func Test_Hashmap_ForEach(t *testing.T) {
	// Arrange
	tc := hashmapForEachTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
	count := 0
	hm.ForEach(func(_ string, _ int) { count++ })

	// Act
	actual := args.Map{"visitCount": count}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_ForEachBreak(t *testing.T) {
	// Arrange
	tc := hashmapForEachBreakTestCase
	hm := coregeneric.HashmapFrom(map[int]int{1: 1, 2: 2, 3: 3})
	count := 0
	hm.ForEachBreak(func(_ int, _ int) bool { count++; return count >= 2 })

	// Act
	actual := args.Map{"visitCount": count}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: String
// ==========================================================================

func Test_Hashmap_String_FromHashmap(t *testing.T) {
	// Arrange
	tc := hashmapStringTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	// Act
	actual := args.Map{"isNonEmpty": hm.String() != ""}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// Note: Nil receiver tests migrated to NilReceiver_test.go using CaseNilSafe pattern.
