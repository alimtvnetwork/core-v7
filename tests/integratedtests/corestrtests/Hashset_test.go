package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Constructors
// ==========================================

func Test_StrHashset_Empty(t *testing.T) {
	safeTest(t, "Test_StrHashset_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"result": hs.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Empty hashset should be empty", actual)
		actual = args.Map{"result": hs.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty hashset length: expected 0", actual)
	})
}

func Test_StrHashset_Cap(t *testing.T) {
	safeTest(t, "Test_StrHashset_Cap", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(50)

		// Act
		actual := args.Map{"result": hs.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Cap hashset should be empty initially", actual)
	})
}

func Test_StrHashset_Strings(t *testing.T) {
	safeTest(t, "Test_StrHashset_Strings", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c", "a"})

		// Act
		actual := args.Map{"result": hs.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Strings with duplicates: expected 3", actual)
	})
}

func Test_StrHashset_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_StrHashset_Strings_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{})

		// Act
		actual := args.Map{"result": hs.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Strings from empty slice should be empty", actual)
	})
}

func Test_StrHashset_StringsSpreadItems(t *testing.T) {
	safeTest(t, "Test_StrHashset_StringsSpreadItems", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("x", "y", "z")

		// Act
		actual := args.Map{"result": hs.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "StringsSpreadItems: expected 3", actual)
	})
}

func Test_StrHashset_UsingMap(t *testing.T) {
	safeTest(t, "Test_StrHashset_UsingMap", func() {
		// Arrange
		m := map[string]bool{"a": true, "b": true}
		hs := corestr.New.Hashset.UsingMap(m)

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UsingMap: expected 2", actual)
	})
}

// ==========================================
// Add / AddBool — caching behavior (bug 42)
// ==========================================

func Test_StrHashset_Add_SetsHasMapUpdated(t *testing.T) {
	safeTest(t, "Test_StrHashset_Add_SetsHasMapUpdated", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("hello")

		// Act
		actual := args.Map{"result": hs.Has("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Add should insert item", actual)
		actual = args.Map{"result": hs.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After Add: expected 1", actual)
	})
}

func Test_StrHashset_Add_Duplicate_NoIncrease(t *testing.T) {
	safeTest(t, "Test_StrHashset_Add_Duplicate_NoIncrease", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a").Add("a").Add("a")

		// Act
		actual := args.Map{"result": hs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Add duplicate: expected 1", actual)
	})
}

func Test_StrHashset_AddBool_FirstAdd_ReturnsFalse(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddBool_FirstAdd_ReturnsFalse", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		existed := hs.AddBool("x")

		// Act
		actual := args.Map{"result": existed}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddBool first add should return false (did not exist)", actual)
		actual = args.Map{"result": hs.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After AddBool: expected 1", actual)
	})
}

func Test_StrHashset_AddBool_SecondAdd_ReturnsTrue(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddBool_SecondAdd_ReturnsTrue", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddBool("x")
		existed := hs.AddBool("x")

		// Act
		actual := args.Map{"result": existed}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "AddBool second add should return true (already existed)", actual)
		actual = args.Map{"result": hs.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddBool duplicate: expected 1", actual)
	})
}

func Test_StrHashset_AddBool_CacheInvalidation(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddBool_CacheInvalidation", func() {
		// Arrange
		// Bug 42 context: hasMapUpdated must be set on AddBool for new items
		hs := corestr.New.Hashset.Empty()
		// Force cache by calling List
		_ = hs.List()
		hs.AddBool("new-item")
		list := hs.List()
		found := false
		for _, v := range list {
			if v == "new-item" {
				found = true
			}
		}

		// Act
		actual := args.Map{"result": found}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "AddBool should invalidate cache so List reflects new item", actual)
	})
}

func Test_StrHashset_AddNonEmpty_SkipsEmpty(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddNonEmpty_SkipsEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmpty("")

		// Act
		actual := args.Map{"result": hs.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty should skip empty string", actual)
		hs.AddNonEmpty("valid")
		actual = args.Map{"result": hs.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty: expected 1", actual)
	})
}

func Test_StrHashset_AddNonEmptyWhitespace_SkipsWhitespace(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddNonEmptyWhitespace_SkipsWhitespace", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmptyWhitespace("   ")

		// Act
		actual := args.Map{"result": hs.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyWhitespace should skip whitespace-only", actual)
		hs.AddNonEmptyWhitespace("valid")
		actual = args.Map{"result": hs.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ==========================================
// Adds / AddStrings / AddIf / AddIfMany
// ==========================================

func Test_StrHashset_Adds_Variadic(t *testing.T) {
	safeTest(t, "Test_StrHashset_Adds_Variadic", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b", "c")

		// Act
		actual := args.Map{"result": hs.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Adds: expected 3", actual)
	})
}

func Test_StrHashset_AddStrings(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(10)
		hs.AddStrings([]string{"x", "y", "x"})

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddStrings with dup: expected 2", actual)
	})
}

func Test_StrHashset_AddIf_True(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddIf_True", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddIf(true, "yes")

		// Act
		actual := args.Map{"result": hs.Has("yes")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "AddIf(true) should add item", actual)
	})
}

func Test_StrHashset_AddIf_False(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddIf_False", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddIf(false, "no")

		// Act
		actual := args.Map{"result": hs.Has("no")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddIf(false) should not add item", actual)
	})
}

func Test_StrHashset_AddIfMany_True(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddIfMany_True", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddIfMany(true, "a", "b")

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddIfMany(true): expected 2", actual)
	})
}

func Test_StrHashset_AddIfMany_False(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddIfMany_False", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddIfMany(false, "a", "b")

		// Act
		actual := args.Map{"result": hs.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddIfMany(false): expected 0", actual)
	})
}

// ==========================================
// AddHashsetItems / AddItemsMap
// ==========================================

func Test_StrHashset_AddHashsetItems_Merge(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddHashsetItems_Merge", func() {
		// Arrange
		hs1 := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs2 := corestr.New.Hashset.StringsSpreadItems("b", "c")
		hs1.AddHashsetItems(hs2)

		// Act
		actual := args.Map{"result": hs1.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Merge: expected 3", actual)
	})
}

func Test_StrHashset_AddHashsetItems_Nil(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddHashsetItems_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		result := hs.AddHashsetItems(nil)

		// Act
		actual := args.Map{"result": result != hs || hs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems(nil) should be no-op", actual)
	})
}

func Test_StrHashset_AddItemsMap_OnlyTrueValues(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddItemsMap_OnlyTrueValues", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddItemsMap(map[string]bool{"yes": true, "no": false, "also": true})

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddItemsMap: expected 2 (only true)", actual)
		actual = args.Map{"result": hs.Has("no")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddItemsMap should skip false-valued entries", actual)
	})
}

// ==========================================
// Has / HasAll / HasAny
// ==========================================

func Test_StrHashset_Has_Existing(t *testing.T) {
	safeTest(t, "Test_StrHashset_Has_Existing", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("alpha", "beta")

		// Act
		actual := args.Map{"result": hs.Has("alpha")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Has should find existing item", actual)
	})
}

func Test_StrHashset_Has_Missing(t *testing.T) {
	safeTest(t, "Test_StrHashset_Has_Missing", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("alpha", "beta")

		// Act
		actual := args.Map{"result": hs.Has("gamma")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Has should not find missing item", actual)
	})
}

func Test_StrHashset_HasAll_AllPresent(t *testing.T) {
	safeTest(t, "Test_StrHashset_HasAll_AllPresent", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")

		// Act
		actual := args.Map{"result": hs.HasAll("a", "c")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAll should return true when all present", actual)
	})
}

func Test_StrHashset_HasAll_OneMissing(t *testing.T) {
	safeTest(t, "Test_StrHashset_HasAll_OneMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{"result": hs.HasAll("a", "z")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasAll should return false when one missing", actual)
	})
}

func Test_StrHashset_HasAll_EmptyArgs(t *testing.T) {
	safeTest(t, "Test_StrHashset_HasAll_EmptyArgs", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"result": hs.HasAll()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAll with no args should return true", actual)
	})
}

func Test_StrHashset_HasAny_OnePresent(t *testing.T) {
	safeTest(t, "Test_StrHashset_HasAny_OnePresent", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")

		// Act
		actual := args.Map{"result": hs.HasAny("z", "b", "y")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAny should return true when at least one present", actual)
	})
}

func Test_StrHashset_HasAny_NonePresent(t *testing.T) {
	safeTest(t, "Test_StrHashset_HasAny_NonePresent", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{"result": hs.HasAny("x", "y")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasAny should return false when none present", actual)
	})
}

func Test_StrHashset_HasAny_EmptyArgs(t *testing.T) {
	safeTest(t, "Test_StrHashset_HasAny_EmptyArgs", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"result": hs.HasAny()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasAny with no args should return false", actual)
	})
}

func Test_StrHashset_IsMissing(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"result": hs.IsMissing("z")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsMissing should return true for absent key", actual)
		actual = args.Map{"result": hs.IsMissing("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsMissing should return false for present key", actual)
	})
}

func Test_StrHashset_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsAllMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{"result": hs.IsAllMissing("x", "y")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsAllMissing should return true when all missing", actual)
		actual = args.Map{"result": hs.IsAllMissing("x", "a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsAllMissing should return false when any present", actual)
	})
}

// ==========================================
// Remove / SafeRemove
// ==========================================

func Test_StrHashset_Remove(t *testing.T) {
	safeTest(t, "Test_StrHashset_Remove", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		hs.Remove("b")

		// Act
		actual := args.Map{"result": hs.Has("b")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Remove should delete item", actual)
		actual = args.Map{"result": hs.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After remove: expected 2", actual)
	})
}

func Test_StrHashset_SafeRemove_Existing(t *testing.T) {
	safeTest(t, "Test_StrHashset_SafeRemove_Existing", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs.SafeRemove("a")

		// Act
		actual := args.Map{"result": hs.Has("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "SafeRemove should delete existing item", actual)
	})
}

func Test_StrHashset_SafeRemove_Missing(t *testing.T) {
	safeTest(t, "Test_StrHashset_SafeRemove_Missing", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.SafeRemove("z") // should not panic

		// Act
		actual := args.Map{"result": hs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "SafeRemove missing: expected 1", actual)
	})
}

// ==========================================
// Resize
// ==========================================

func Test_StrHashset_Resize_LargerPreservesItems(t *testing.T) {
	safeTest(t, "Test_StrHashset_Resize_LargerPreservesItems", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs.Resize(100)

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Resize should preserve items: expected 2", actual)
		actual = args.Map{"result": hs.Has("a") || !hs.Has("b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Resize should preserve all items", actual)
	})
}

func Test_StrHashset_Resize_SmallerIsNoOp(t *testing.T) {
	safeTest(t, "Test_StrHashset_Resize_SmallerIsNoOp", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		hs.Resize(1)

		// Act
		actual := args.Map{"result": hs.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Resize smaller: expected 3", actual)
	})
}

func Test_StrHashset_AddCapacities(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddCapacities", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.AddCapacities(10, 20)

		// Act
		actual := args.Map{"result": hs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddCapacities should preserve items: expected 1", actual)
	})
}

// ==========================================
// IsEquals
// ==========================================

func Test_StrHashset_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_BothNil", func() {
		// Arrange
		var a, b *corestr.Hashset

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Two nil hashsets should be equal", actual)
	})
}

func Test_StrHashset_IsEquals_OneNil(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_OneNil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		var nilHs *corestr.Hashset

		// Act
		actual := args.Map{"result": hs.IsEquals(nilHs)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Non-nil vs nil should not be equal", actual)
	})
}

func Test_StrHashset_IsEquals_SamePointer(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_SamePointer", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{"result": hs.IsEquals(hs)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Same pointer should be equal", actual)
	})
}

func Test_StrHashset_IsEquals_SameContent(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_SameContent", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("x", "y")
		b := corestr.New.Hashset.StringsSpreadItems("y", "x")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Same content should be equal", actual)
	})
}

func Test_StrHashset_IsEquals_DifferentContent(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_DifferentContent", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		b := corestr.New.Hashset.StringsSpreadItems("a", "c")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Different content should not be equal", actual)
	})
}

func Test_StrHashset_IsEquals_DifferentLength(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_DifferentLength", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		b := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Different length should not be equal", actual)
	})
}

func Test_StrHashset_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.New.Hashset.Empty()
		b := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Two empty hashsets should be equal", actual)
	})
}

// ==========================================
// List caching (bug 42 context)
// ==========================================

func Test_StrHashset_List_CacheInvalidatedAfterAdd(t *testing.T) {
	safeTest(t, "Test_StrHashset_List_CacheInvalidatedAfterAdd", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		list1 := hs.List()

		// Act
		actual := args.Map{"result": len(list1) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Initial list: expected 1", actual)
		hs.Add("b")
		list2 := hs.List()
		actual = args.Map{"result": len(list2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After Add, list should reflect new item: expected 2", actual)
	})
}

func Test_StrHashset_List_CacheInvalidatedAfterRemove(t *testing.T) {
	safeTest(t, "Test_StrHashset_List_CacheInvalidatedAfterRemove", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		_ = hs.List() // populate cache
		hs.Remove("a")
		list := hs.List()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After Remove, list should reflect removal: expected 1", actual)
	})
}

func Test_StrHashset_List_CacheInvalidatedAfterAdds(t *testing.T) {
	safeTest(t, "Test_StrHashset_List_CacheInvalidatedAfterAdds", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		_ = hs.List() // populate cache
		hs.Adds("x", "y")
		list := hs.List()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After Adds, list should reflect new items: expected 2", actual)
	})
}

// ==========================================
// Clear / Dispose
// ==========================================

func Test_StrHashset_Clear(t *testing.T) {
	safeTest(t, "Test_StrHashset_Clear", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs.Clear()

		// Act
		actual := args.Map{"result": hs.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Clear should make hashset empty", actual)
	})
}

// ==========================================
// Nil receiver guards (CaseNilSafe pattern)
// ==========================================

func Test_StrHashset_NilReceiver(t *testing.T) {
	safeTest(t, "Test_StrHashset_NilReceiver", func() {
		for caseIndex, tc := range hashsetNilReceiverTestCases {
			// Arrange (implicit — nil receiver)

			// Act & Assert
			tc.ShouldBeSafe(t, caseIndex)
		}
	})
}

// ==========================================
// String-specific methods
// ==========================================

func Test_StrHashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_StrHashset_ToLowerSet", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("ABC", "Def")
		lower := hs.ToLowerSet()

		// Act
		actual := args.Map{"result": lower.Has("abc") || !lower.Has("def")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "ToLowerSet should lowercase all keys", actual)
		actual = args.Map{"result": lower.Has("ABC")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ToLowerSet should not retain original case", actual)
	})
}

func Test_StrHashset_GetAllExceptHashset(t *testing.T) {
	safeTest(t, "Test_StrHashset_GetAllExceptHashset", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		except := corestr.New.Hashset.StringsSpreadItems("b")
		result := hs.GetAllExceptHashset(except)

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "GetAllExceptHashset: expected 2", actual)
	})
}

func Test_StrHashset_GetAllExceptHashset_NilExcept(t *testing.T) {
	safeTest(t, "Test_StrHashset_GetAllExceptHashset_NilExcept", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		result := hs.GetAllExceptHashset(nil)

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "GetAllExceptHashset(nil): expected 2", actual)
	})
}

func Test_StrHashset_Collection_Conversion(t *testing.T) {
	safeTest(t, "Test_StrHashset_Collection_Conversion", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		col := hs.Collection()

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Collection: expected 2", actual)
	})
}

func Test_StrHashset_OrderedList(t *testing.T) {
	safeTest(t, "Test_StrHashset_OrderedList", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("c", "a", "b")
		list := hs.OrderedList()

		// Act
		actual := args.Map{"result": len(list) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "OrderedList: expected 3", actual)
		actual = args.Map{"result": list[0] != "a" || list[1] != "b" || list[2] != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "OrderedList should be sorted asc", actual)
	})
}

func Test_StrHashset_JoinSorted(t *testing.T) {
	safeTest(t, "Test_StrHashset_JoinSorted", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("c", "a", "b")
		result := hs.JoinSorted(",")

		// Act
		actual := args.Map{"result": result != "a,b,c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "JoinSorted: expected 'a,b,c', got ''", actual)
	})
}
