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
	tc := hashsetAddDuplicateTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.Add(1).Add(2).Add(3).Add(1).Add(2)

	actual := args.Map{"length": hs.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddBool(t *testing.T) {
	tc := hashsetAddBoolTestCase
	hs := coregeneric.EmptyHashset[string]()
	first := hs.AddBool("a")
	second := hs.AddBool("a")

	actual := args.Map{
		"firstExisted":  first,
		"secondExisted": second,
		"length":        hs.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddsVariadic(t *testing.T) {
	tc := hashsetAddsVariadicTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.Adds(10, 20, 30, 10)

	actual := args.Map{"length": hs.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddSlice(t *testing.T) {
	tc := hashsetAddSliceTestCase
	hs := coregeneric.EmptyHashset[string]()
	hs.AddSlice([]string{"x", "y", "z"})

	actual := args.Map{
		"length": hs.Length(),
		"hasX":   hs.Has("x"),
		"hasY":   hs.Has("y"),
		"hasZ":   hs.Has("z"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — AddIf / AddIfMany
// ==========================================================================

func Test_Hashset_AddIfTrue(t *testing.T) {
	tc := hashsetAddIfTrueTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIf(true, 42)

	actual := args.Map{
		"length":  hs.Length(),
		"hasItem": hs.Has(42),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddIfFalse(t *testing.T) {
	tc := hashsetAddIfFalseTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIf(false, 42)

	actual := args.Map{"length": hs.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddIfManyTrue(t *testing.T) {
	tc := hashsetAddIfManyTrueTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIfMany(true, 1, 2, 3)

	actual := args.Map{"length": hs.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddIfManyFalse(t *testing.T) {
	tc := hashsetAddIfManyFalseTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIfMany(false, 1, 2, 3)

	actual := args.Map{"length": hs.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — AddHashsetItems / AddItemsMap
// ==========================================================================

func Test_Hashset_MergeOtherSet(t *testing.T) {
	tc := hashsetMergeOtherSetTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2})
	other := coregeneric.HashsetFrom([]int{3, 4})
	hs.AddHashsetItems(other)

	actual := args.Map{
		"length": hs.Length(),
		"has3":   hs.Has(3),
		"has4":   hs.Has(4),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_MergeNilOther(t *testing.T) {
	tc := hashsetMergeNilOtherTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2})
	hs.AddHashsetItems(nil)

	actual := args.Map{"length": hs.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_MergeEmptyOther(t *testing.T) {
	tc := hashsetMergeEmptyOtherTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2})
	hs.AddHashsetItems(coregeneric.EmptyHashset[int]())

	actual := args.Map{"length": hs.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_AddItemsMap(t *testing.T) {
	tc := hashsetAddItemsMapTestCase
	hs := coregeneric.EmptyHashset[string]()
	hs.AddItemsMap(map[string]bool{
		"yes":  true,
		"also": true,
		"nope": false,
	})

	actual := args.Map{
		"length":  hs.Length(),
		"hasYes":  hs.Has("yes"),
		"hasNope": hs.Has("nope"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — Remove edge cases
// ==========================================================================

func Test_Hashset_RemoveExisting(t *testing.T) {
	tc := hashsetRemoveExistingTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	existed := hs.Remove(2)

	actual := args.Map{
		"existed":  existed,
		"length":   hs.Length(),
		"stillHas": hs.Has(2),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_RemoveNonExisting(t *testing.T) {
	tc := hashsetRemoveNonExistingTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	existed := hs.Remove(99)

	actual := args.Map{
		"existed": existed,
		"length":  hs.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — Has / Contains
// ==========================================================================

func Test_Hashset_Has(t *testing.T) {
	tc := hashsetHasTestCase
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})

	actual := args.Map{
		"hasExisting": hs.Has("a"),
		"hasMissing":  hs.Has("z"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_ContainsAlias(t *testing.T) {
	tc := hashsetContainsAliasTestCase
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})

	actual := args.Map{
		"containsExisting": hs.Contains("b"),
		"containsMissing":  hs.Contains("z"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — HasAll / HasAny
// ==========================================================================

func Test_Hashset_HasAllTrue(t *testing.T) {
	tc := hashsetHasAllTrueTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	actual := args.Map{"hasAll": hs.HasAll(1, 3, 5)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_HasAllFalse(t *testing.T) {
	tc := hashsetHasAllFalseTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	actual := args.Map{"hasAll": hs.HasAll(1, 99)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_HasAnyTrue(t *testing.T) {
	tc := hashsetHasAnyTrueTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	actual := args.Map{"hasAny": hs.HasAny(99, 3)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_HasAnyFalse(t *testing.T) {
	tc := hashsetHasAnyFalseTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	actual := args.Map{"hasAny": hs.HasAny(99, 100)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_HasAllEmptyArgs(t *testing.T) {
	tc := hashsetHasAllEmptyArgsTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	actual := args.Map{"hasAll": hs.HasAll()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_HasAnyEmptyArgs(t *testing.T) {
	tc := hashsetHasAnyEmptyArgsTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	actual := args.Map{"hasAny": hs.HasAny()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — IsEquals
// ==========================================================================

func Test_Hashset_IsEquals_SameItems(t *testing.T) {
	tc := hashsetIsEqualsSameItemsTestCase
	a := coregeneric.HashsetFrom([]int{1, 2, 3})
	b := coregeneric.HashsetFrom([]int{3, 2, 1})

	actual := args.Map{"isEquals": a.IsEquals(b)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_IsEquals_DifferentItems(t *testing.T) {
	tc := hashsetIsEqualsDifferentItemsTestCase
	a := coregeneric.HashsetFrom([]int{1, 2, 3})
	b := coregeneric.HashsetFrom([]int{1, 2, 4})

	actual := args.Map{"isEquals": a.IsEquals(b)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_IsEquals_DifferentLength(t *testing.T) {
	tc := hashsetIsEqualsDifferentLengthTestCase
	a := coregeneric.HashsetFrom([]int{1, 2})
	b := coregeneric.HashsetFrom([]int{1, 2, 3})

	actual := args.Map{"isEquals": a.IsEquals(b)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_IsEquals_BothNil(t *testing.T) {
	tc := hashsetIsEqualsBothNilTestCase
	var a *coregeneric.Hashset[int]
	var b *coregeneric.Hashset[int]

	actual := args.Map{"isEquals": a.IsEquals(b)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_IsEquals_NilVsNonNil(t *testing.T) {
	tc := hashsetIsEqualsNilVsNonNilTestCase
	var a *coregeneric.Hashset[int]
	b := coregeneric.EmptyHashset[int]()

	actual := args.Map{"isEquals": a.IsEquals(b)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_IsEquals_SamePointer(t *testing.T) {
	tc := hashsetIsEqualsSamePointerTestCase
	a := coregeneric.HashsetFrom([]int{1, 2})

	actual := args.Map{"isEquals": a.IsEquals(a)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_IsEquals_BothEmpty(t *testing.T) {
	tc := hashsetIsEqualsBothEmptyTestCase
	a := coregeneric.EmptyHashset[int]()
	b := coregeneric.EmptyHashset[int]()

	actual := args.Map{"isEquals": a.IsEquals(b)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — Resize
// ==========================================================================

func Test_Hashset_ResizeLarger(t *testing.T) {
	tc := hashsetResizeLargerTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	hs.Resize(100)

	actual := args.Map{
		"length": hs.Length(),
		"has1":   hs.Has(1),
		"has2":   hs.Has(2),
		"has3":   hs.Has(3),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_ResizeSmaller(t *testing.T) {
	tc := hashsetResizeSmallerTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	hs.Resize(1)

	actual := args.Map{"length": hs.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — List / ListPtr / Map / Collection / String
// ==========================================================================

func Test_Hashset_OutputList(t *testing.T) {
	tc := hashsetOutputListTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})

	actual := args.Map{"listLen": len(hs.List())}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_OutputListEmpty(t *testing.T) {
	tc := hashsetOutputListEmptyTestCase
	hs := coregeneric.EmptyHashset[int]()

	actual := args.Map{"listLen": len(hs.List())}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_OutputListPtr(t *testing.T) {
	tc := hashsetOutputListPtrTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})

	actual := args.Map{"isNotNil": hs.ListPtr() != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_OutputMap(t *testing.T) {
	tc := hashsetOutputMapTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})

	actual := args.Map{"mapLen": len(hs.Map())}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_OutputCollection(t *testing.T) {
	tc := hashsetOutputCollectionTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	col := hs.Collection()

	actual := args.Map{"collectionLen": col.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — Lock variants
// ==========================================================================

func Test_Hashset_LockAddContains(t *testing.T) {
	tc := hashsetLockAddContainsTestCase
	hs := coregeneric.EmptyHashset[string]()
	hs.AddLock("a")
	hs.AddLock("b")

	actual := args.Map{
		"length":    hs.Length(),
		"containsA": hs.ContainsLock("a"),
		"containsB": hs.ContainsLock("b"),
		"containsZ": hs.ContainsLock("z"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_LockAddSlice(t *testing.T) {
	tc := hashsetLockAddSliceTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddSliceLock([]int{10, 20, 30})

	actual := args.Map{"length": hs.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_LockRemove(t *testing.T) {
	tc := hashsetLockRemoveTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	existed := hs.RemoveLock(2)

	actual := args.Map{
		"existed":  existed,
		"length":   hs.Length(),
		"stillHas": hs.Has(2),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_LockIsEmptyLength(t *testing.T) {
	tc := hashsetLockIsEmptyLengthTestCase
	hs := coregeneric.EmptyHashset[int]()

	emptyBefore := hs.IsEmptyLock()
	lengthBefore := hs.LengthLock()

	hs.Adds(1, 2)

	emptyAfter := hs.IsEmptyLock()
	lengthAfter := hs.LengthLock()

	actual := args.Map{
		"emptyBefore":  emptyBefore,
		"lengthBefore": lengthBefore,
		"emptyAfter":   emptyAfter,
		"lengthAfter":  lengthAfter,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — Constructors
// ==========================================================================

func Test_Hashset_ConstructorEmpty(t *testing.T) {
	tc := hashsetConstructorEmptyTestCase
	hs := coregeneric.EmptyHashset[int]()

	actual := args.Map{
		"length":  hs.Length(),
		"isEmpty": hs.IsEmpty(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_ConstructorNewCap(t *testing.T) {
	tc := hashsetConstructorNewCapTestCase
	hs := coregeneric.NewHashset[string](10)

	actual := args.Map{
		"length":  hs.Length(),
		"isEmpty": hs.IsEmpty(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_ConstructorFrom(t *testing.T) {
	tc := hashsetConstructorFromTestCase
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})

	actual := args.Map{
		"length": hs.Length(),
		"hasA":   hs.Has("a"),
		"hasB":   hs.Has("b"),
		"hasC":   hs.Has("c"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_ConstructorFromMap(t *testing.T) {
	tc := hashsetConstructorFromMapTestCase
	hs := coregeneric.HashsetFromMap(map[int]bool{10: true, 20: true})

	actual := args.Map{
		"length": hs.Length(),
		"has10":  hs.Has(10),
		"has20":  hs.Has(20),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_ConstructorHasItems(t *testing.T) {
	tc := hashsetConstructorHasItemsTestCase
	pop := coregeneric.HashsetFrom([]int{1})
	empty := coregeneric.EmptyHashset[int]()

	actual := args.Map{
		"populatedHasItems": pop.HasItems(),
		"emptyHasItems":     empty.HasItems(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — String output
// ==========================================================================

func Test_Hashset_StringNotEmpty(t *testing.T) {
	tc := hashsetStringNotEmptyTestCase
	hs := coregeneric.HashsetFrom([]int{1})

	actual := args.Map{"isNonEmpty": hs.String() != ""}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashset — Creator pattern (New.Hashset.X)
// ==========================================================================

func Test_Hashset_CreatorStringItems(t *testing.T) {
	tc := hashsetCreatorStringItemsTestCase
	hs := coregeneric.New.Hashset.String.Items("a", "b", "c")

	actual := args.Map{"length": hs.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_CreatorIntFrom(t *testing.T) {
	tc := hashsetCreatorIntFromTestCase
	hs := coregeneric.New.Hashset.Int.From([]int{1, 2, 3, 1})

	actual := args.Map{"length": hs.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_CreatorEmpty(t *testing.T) {
	tc := hashsetCreatorEmptyTestCase
	hs := coregeneric.New.Hashset.Float64.Empty()

	actual := args.Map{"isEmpty": hs.IsEmpty()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_CreatorCap(t *testing.T) {
	tc := hashsetCreatorCapTestCase
	hs := coregeneric.New.Hashset.Bool.Cap(10)

	actual := args.Map{"isEmpty": hs.IsEmpty()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashset_CreatorUsingMap(t *testing.T) {
	tc := hashsetCreatorUsingMapTestCase
	m := map[uint]bool{1: true, 2: true}
	hs := coregeneric.New.Hashset.Uint.UsingMap(m)

	actual := args.Map{"length": hs.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}
