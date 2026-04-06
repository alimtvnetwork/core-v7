package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: Hashset — Add / AddBool edge cases
// ==========================================================================

func Test_Hashset_AddDuplicate(t *testing.T) {
	// Arrange
	tc := hashsetAddDuplicateTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.Add(1).Add(2).Add(3).Add(1).Add(2)

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddBool(t *testing.T) {
	// Arrange
	tc := hashsetAddBoolTestCase
	hs := coregeneric.EmptyHashset[string]()
	first := hs.AddBool("a")
	second := hs.AddBool("a")

	// Act
	actual := args.Map{
		"firstExisted":  first,
		"secondExisted": second,
		"length":        hs.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddsVariadic(t *testing.T) {
	// Arrange
	tc := hashsetAddsVariadicTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.Adds(10, 20, 30, 10)

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddSlice(t *testing.T) {
	// Arrange
	tc := hashsetAddSliceTestCase
	hs := coregeneric.EmptyHashset[string]()
	hs.AddSlice([]string{"x", "y", "z"})

	// Act
	actual := args.Map{
		"length": hs.Length(),
		"hasX":   hs.Has("x"),
		"hasY":   hs.Has("y"),
		"hasZ":   hs.Has("z"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — AddIf / AddIfMany
// ==========================================================================

func Test_Hashset_AddIfTrue(t *testing.T) {
	// Arrange
	tc := hashsetAddIfTrueTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIf(true, 42)

	// Act
	actual := args.Map{
		"length":  hs.Length(),
		"hasItem": hs.Has(42),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddIfFalse(t *testing.T) {
	// Arrange
	tc := hashsetAddIfFalseTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIf(false, 42)

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddIfManyTrue(t *testing.T) {
	// Arrange
	tc := hashsetAddIfManyTrueTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIfMany(true, 1, 2, 3)

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddIfManyFalse(t *testing.T) {
	// Arrange
	tc := hashsetAddIfManyFalseTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIfMany(false, 1, 2, 3)

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — AddHashsetItems / AddItemsMap
// ==========================================================================

func Test_Hashset_MergeOtherSet(t *testing.T) {
	// Arrange
	tc := hashsetMergeOtherSetTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2})
	other := coregeneric.HashsetFrom([]int{3, 4})
	hs.AddHashsetItems(other)

	// Act
	actual := args.Map{
		"length": hs.Length(),
		"has3":   hs.Has(3),
		"has4":   hs.Has(4),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_MergeNilOther(t *testing.T) {
	// Arrange
	tc := hashsetMergeNilOtherTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2})
	hs.AddHashsetItems(nil)

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_MergeEmptyOther(t *testing.T) {
	// Arrange
	tc := hashsetMergeEmptyOtherTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2})
	hs.AddHashsetItems(coregeneric.EmptyHashset[int]())

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddItemsMap(t *testing.T) {
	// Arrange
	tc := hashsetAddItemsMapTestCase
	hs := coregeneric.EmptyHashset[string]()
	hs.AddItemsMap(map[string]bool{
		"yes":  true,
		"also": true,
		"nope": false,
	})

	// Act
	actual := args.Map{
		"length":  hs.Length(),
		"hasYes":  hs.Has("yes"),
		"hasNope": hs.Has("nope"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — Remove edge cases
// ==========================================================================

func Test_Hashset_RemoveExisting(t *testing.T) {
	// Arrange
	tc := hashsetRemoveExistingTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	existed := hs.Remove(2)

	// Act
	actual := args.Map{
		"existed":  existed,
		"length":   hs.Length(),
		"stillHas": hs.Has(2),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_RemoveNonExisting(t *testing.T) {
	// Arrange
	tc := hashsetRemoveNonExistingTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	existed := hs.Remove(99)

	// Act
	actual := args.Map{
		"existed": existed,
		"length":  hs.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — Has / Contains
// ==========================================================================

func Test_Hashset_Has(t *testing.T) {
	// Arrange
	tc := hashsetHasTestCase
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})

	// Act
	actual := args.Map{
		"hasExisting": hs.Has("a"),
		"hasMissing":  hs.Has("z"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_ContainsAlias(t *testing.T) {
	// Arrange
	tc := hashsetContainsAliasTestCase
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})

	// Act
	actual := args.Map{
		"containsExisting": hs.Contains("b"),
		"containsMissing":  hs.Contains("z"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — HasAll / HasAny
// ==========================================================================

func Test_Hashset_HasAllTrue(t *testing.T) {
	// Arrange
	tc := hashsetHasAllTrueTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	// Act
	actual := args.Map{"hasAll": hs.HasAll(1, 3, 5)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_HasAllFalse(t *testing.T) {
	// Arrange
	tc := hashsetHasAllFalseTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	// Act
	actual := args.Map{"hasAll": hs.HasAll(1, 99)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_HasAnyTrue(t *testing.T) {
	// Arrange
	tc := hashsetHasAnyTrueTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	// Act
	actual := args.Map{"hasAny": hs.HasAny(99, 3)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_HasAnyFalse(t *testing.T) {
	// Arrange
	tc := hashsetHasAnyFalseTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	// Act
	actual := args.Map{"hasAny": hs.HasAny(99, 100)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_HasAllEmptyArgs(t *testing.T) {
	// Arrange
	tc := hashsetHasAllEmptyArgsTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	// Act
	actual := args.Map{"hasAll": hs.HasAll()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_HasAnyEmptyArgs(t *testing.T) {
	// Arrange
	tc := hashsetHasAnyEmptyArgsTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	// Act
	actual := args.Map{"hasAny": hs.HasAny()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — IsEquals
// ==========================================================================

func Test_Hashset_IsEquals_SameItems(t *testing.T) {
	// Arrange
	tc := hashsetIsEqualsSameItemsTestCase
	a := coregeneric.HashsetFrom([]int{1, 2, 3})
	b := coregeneric.HashsetFrom([]int{3, 2, 1})

	// Act
	actual := args.Map{"isEquals": a.IsEquals(b)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_IsEquals_DifferentItems(t *testing.T) {
	// Arrange
	tc := hashsetIsEqualsDifferentItemsTestCase
	a := coregeneric.HashsetFrom([]int{1, 2, 3})
	b := coregeneric.HashsetFrom([]int{1, 2, 4})

	// Act
	actual := args.Map{"isEquals": a.IsEquals(b)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_IsEquals_DifferentLength(t *testing.T) {
	// Arrange
	tc := hashsetIsEqualsDifferentLengthTestCase
	a := coregeneric.HashsetFrom([]int{1, 2})
	b := coregeneric.HashsetFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"isEquals": a.IsEquals(b)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_IsEquals_BothNil(t *testing.T) {
	// Arrange
	tc := hashsetIsEqualsBothNilTestCase
	var a *coregeneric.Hashset[int]
	var b *coregeneric.Hashset[int]

	// Act
	actual := args.Map{"isEquals": a.IsEquals(b)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_IsEquals_NilVsNonNil(t *testing.T) {
	// Arrange
	tc := hashsetIsEqualsNilVsNonNilTestCase
	var a *coregeneric.Hashset[int]
	b := coregeneric.EmptyHashset[int]()

	// Act
	actual := args.Map{"isEquals": a.IsEquals(b)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_IsEquals_SamePointer(t *testing.T) {
	// Arrange
	tc := hashsetIsEqualsSamePointerTestCase
	a := coregeneric.HashsetFrom([]int{1, 2})

	// Act
	actual := args.Map{"isEquals": a.IsEquals(a)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_IsEquals_BothEmpty(t *testing.T) {
	// Arrange
	tc := hashsetIsEqualsBothEmptyTestCase
	a := coregeneric.EmptyHashset[int]()
	b := coregeneric.EmptyHashset[int]()

	// Act
	actual := args.Map{"isEquals": a.IsEquals(b)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — Resize
// ==========================================================================

func Test_Hashset_ResizeLarger(t *testing.T) {
	// Arrange
	tc := hashsetResizeLargerTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	hs.Resize(100)

	// Act
	actual := args.Map{
		"length": hs.Length(),
		"has1":   hs.Has(1),
		"has2":   hs.Has(2),
		"has3":   hs.Has(3),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_ResizeSmaller(t *testing.T) {
	// Arrange
	tc := hashsetResizeSmallerTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	hs.Resize(1)

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — List / ListPtr / Map / Collection / String
// ==========================================================================

func Test_Hashset_OutputList(t *testing.T) {
	// Arrange
	tc := hashsetOutputListTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"listLen": len(hs.List())}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_OutputListEmpty(t *testing.T) {
	// Arrange
	tc := hashsetOutputListEmptyTestCase
	hs := coregeneric.EmptyHashset[int]()

	// Act
	actual := args.Map{"listLen": len(hs.List())}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_OutputListPtr(t *testing.T) {
	// Arrange
	tc := hashsetOutputListPtrTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"isNotNil": hs.ListPtr() != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_OutputMap(t *testing.T) {
	// Arrange
	tc := hashsetOutputMapTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"mapLen": len(hs.Map())}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_OutputCollection(t *testing.T) {
	// Arrange
	tc := hashsetOutputCollectionTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	col := hs.Collection()

	// Act
	actual := args.Map{"collectionLen": col.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — Lock variants
// ==========================================================================

func Test_Hashset_LockAddContains(t *testing.T) {
	// Arrange
	tc := hashsetLockAddContainsTestCase
	hs := coregeneric.EmptyHashset[string]()
	hs.AddLock("a")
	hs.AddLock("b")

	// Act
	actual := args.Map{
		"length":    hs.Length(),
		"containsA": hs.ContainsLock("a"),
		"containsB": hs.ContainsLock("b"),
		"containsZ": hs.ContainsLock("z"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_LockAddSlice(t *testing.T) {
	// Arrange
	tc := hashsetLockAddSliceTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddSliceLock([]int{10, 20, 30})

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_LockRemove(t *testing.T) {
	// Arrange
	tc := hashsetLockRemoveTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	existed := hs.RemoveLock(2)

	// Act
	actual := args.Map{
		"existed":  existed,
		"length":   hs.Length(),
		"stillHas": hs.Has(2),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_LockIsEmptyLength(t *testing.T) {
	// Arrange
	tc := hashsetLockIsEmptyLengthTestCase
	hs := coregeneric.EmptyHashset[int]()

	emptyBefore := hs.IsEmptyLock()
	lengthBefore := hs.LengthLock()

	hs.Adds(1, 2)

	emptyAfter := hs.IsEmptyLock()
	lengthAfter := hs.LengthLock()

	// Act
	actual := args.Map{
		"emptyBefore":  emptyBefore,
		"lengthBefore": lengthBefore,
		"emptyAfter":   emptyAfter,
		"lengthAfter":  lengthAfter,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — Constructors
// ==========================================================================

func Test_Hashset_ConstructorEmpty(t *testing.T) {
	// Arrange
	tc := hashsetConstructorEmptyTestCase
	hs := coregeneric.EmptyHashset[int]()

	// Act
	actual := args.Map{
		"length":  hs.Length(),
		"isEmpty": hs.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_ConstructorNewCap(t *testing.T) {
	// Arrange
	tc := hashsetConstructorNewCapTestCase
	hs := coregeneric.NewHashset[string](10)

	// Act
	actual := args.Map{
		"length":  hs.Length(),
		"isEmpty": hs.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_ConstructorFrom(t *testing.T) {
	// Arrange
	tc := hashsetConstructorFromTestCase
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})

	// Act
	actual := args.Map{
		"length": hs.Length(),
		"hasA":   hs.Has("a"),
		"hasB":   hs.Has("b"),
		"hasC":   hs.Has("c"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_ConstructorFromMap(t *testing.T) {
	// Arrange
	tc := hashsetConstructorFromMapTestCase
	hs := coregeneric.HashsetFromMap(map[int]bool{10: true, 20: true})

	// Act
	actual := args.Map{
		"length": hs.Length(),
		"has10":  hs.Has(10),
		"has20":  hs.Has(20),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_ConstructorHasItems(t *testing.T) {
	// Arrange
	tc := hashsetConstructorHasItemsTestCase
	pop := coregeneric.HashsetFrom([]int{1})
	empty := coregeneric.EmptyHashset[int]()

	// Act
	actual := args.Map{
		"populatedHasItems": pop.HasItems(),
		"emptyHasItems":     empty.HasItems(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — String output
// ==========================================================================

func Test_Hashset_StringNotEmpty(t *testing.T) {
	// Arrange
	tc := hashsetStringNotEmptyTestCase
	hs := coregeneric.HashsetFrom([]int{1})

	// Act
	actual := args.Map{"isNonEmpty": hs.String() != ""}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — Creator pattern (New.Hashset.X)
// ==========================================================================

func Test_Hashset_CreatorStringItems(t *testing.T) {
	// Arrange
	tc := hashsetCreatorStringItemsTestCase
	hs := coregeneric.New.Hashset.String.Items("a", "b", "c")

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_CreatorIntFrom(t *testing.T) {
	// Arrange
	tc := hashsetCreatorIntFromTestCase
	hs := coregeneric.New.Hashset.Int.From([]int{1, 2, 3, 1})

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_CreatorEmpty(t *testing.T) {
	// Arrange
	tc := hashsetCreatorEmptyTestCase
	hs := coregeneric.New.Hashset.Float64.Empty()

	// Act
	actual := args.Map{"isEmpty": hs.IsEmpty()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_CreatorCap(t *testing.T) {
	// Arrange
	tc := hashsetCreatorCapTestCase
	hs := coregeneric.New.Hashset.Bool.Cap(10)

	// Act
	actual := args.Map{"isEmpty": hs.IsEmpty()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_CreatorUsingMap(t *testing.T) {
	// Arrange
	tc := hashsetCreatorUsingMapTestCase
	m := map[uint]bool{1: true, 2: true}
	hs := coregeneric.New.Hashset.Uint.UsingMap(m)

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
