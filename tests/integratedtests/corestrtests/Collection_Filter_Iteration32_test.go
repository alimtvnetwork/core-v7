package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — Filter / FilterLock / FilteredCollection / FilteredCollectionLock
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_Filter_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Filter", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"aa", "b", "cc"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 2, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Filter", actual)
	})
}

func Test_Collection_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Filter_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- Filter empty", actual)
	})
}

func Test_Collection_Filter_Break(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Filter_Break", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, true, i == 0
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Filter break", actual)
	})
}

func Test_Collection_FilterLock(t *testing.T) {
	safeTest(t, "Test_I32_Collection_FilterLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 1, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- FilterLock", actual)
	})
}

func Test_Collection_FilterLock_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_FilterLock_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		result := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- FilterLock empty", actual)
	})
}

func Test_Collection_FilterLock_Break(t *testing.T) {
	safeTest(t, "Test_I32_Collection_FilterLock_Break", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, i == 1
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- FilterLock break", actual)
	})
}

func Test_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_I32_Collection_FilteredCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		fc := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 1, false
		})

		// Act
		actual := args.Map{"len": fc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- FilteredCollection", actual)
	})
}

func Test_Collection_FilteredCollectionLock(t *testing.T) {
	safeTest(t, "Test_I32_Collection_FilteredCollectionLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		fc := c.FilteredCollectionLock(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 2, false
		})

		// Act
		actual := args.Map{"len": fc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- FilteredCollectionLock", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — FilterPtr / FilterPtrLock
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_FilterPtr(t *testing.T) {
	safeTest(t, "Test_I32_Collection_FilterPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := c.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, len(*s) == 1, false
		})

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- FilterPtr", actual)
	})
}

func Test_Collection_FilterPtr_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_FilterPtr_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		result := c.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- FilterPtr empty", actual)
	})
}

func Test_Collection_FilterPtr_Break(t *testing.T) {
	safeTest(t, "Test_I32_Collection_FilterPtr_Break", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, true, i == 0
		})

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- FilterPtr break", actual)
	})
}

func Test_Collection_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_I32_Collection_FilterPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, len(*s) == 2, false
		})

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- FilterPtrLock", actual)
	})
}

func Test_Collection_FilterPtrLock_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_FilterPtrLock_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		result := c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- FilterPtrLock empty", actual)
	})
}

func Test_Collection_FilterPtrLock_Break(t *testing.T) {
	safeTest(t, "Test_I32_Collection_FilterPtrLock_Break", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, true, i == 1
		})

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- FilterPtrLock break", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — AppendAnysUsingFilter / AppendAnysUsingFilterLock
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_AppendAnysUsingFilter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AppendAnysUsingFilter", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 0, false
		}, "hello", nil, "world")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AppendAnysUsingFilter", actual)
	})
}

func Test_Collection_AppendAnysUsingFilter_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AppendAnysUsingFilter_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AppendAnysUsingFilter empty", actual)
	})
}

func Test_Collection_AppendAnysUsingFilter_Break_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AppendAnysUsingFilter_Break", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, true, true
		}, "a", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AppendAnysUsingFilter break", actual)
	})
}

func Test_Collection_AppendAnysUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AppendAnysUsingFilterLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysUsingFilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		}, "hello")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AppendAnysUsingFilterLock", actual)
	})
}

func Test_Collection_AppendAnysUsingFilterLock_Nil_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AppendAnysUsingFilterLock_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysUsingFilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		}, nil...)

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AppendAnysUsingFilterLock nil", actual)
	})
}

func Test_Collection_AppendAnysUsingFilterLock_Break(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AppendAnysUsingFilterLock_Break", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysUsingFilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, true
		}, "a", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AppendAnysUsingFilterLock break", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — NonEmptyList / NonEmptyListPtr / NonEmptyItems
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_NonEmptyList_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_NonEmptyList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})

		// Act
		actual := args.Map{
			"len": len(c.NonEmptyList()),
			"ptrLen": len(*c.NonEmptyListPtr()),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"ptrLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- NonEmptyList", actual)
	})
}

func Test_Collection_NonEmptyList_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_NonEmptyList_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"len": len(c.NonEmptyList())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- NonEmptyList empty", actual)
	})
}

func Test_Collection_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_I32_Collection_NonEmptyItems", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})

		// Act
		actual := args.Map{
			"len": len(c.NonEmptyItems()),
			"ptrLen": len(c.NonEmptyItemsPtr()),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"ptrLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- NonEmptyItems", actual)
	})
}

func Test_Collection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_I32_Collection_NonEmptyItemsOrNonWhitespace", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "  ", ""})

		// Act
		actual := args.Map{
			"len": len(c.NonEmptyItemsOrNonWhitespace()),
			"ptrLen": len(c.NonEmptyItemsOrNonWhitespacePtr()),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"ptrLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- NonEmptyItemsOrNonWhitespace", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — Unique
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_I32_Collection_UniqueBoolMap", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})

		// Act
		actual := args.Map{"len": len(c.UniqueBoolMap())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- UniqueBoolMap", actual)
	})
}

func Test_Collection_UniqueBoolMapLock(t *testing.T) {
	safeTest(t, "Test_I32_Collection_UniqueBoolMapLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "a"})

		// Act
		actual := args.Map{"len": len(c.UniqueBoolMapLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- UniqueBoolMapLock", actual)
	})
}

func Test_Collection_UniqueList_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_UniqueList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})

		// Act
		actual := args.Map{"len": len(c.UniqueList())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- UniqueList", actual)
	})
}

func Test_Collection_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_I32_Collection_UniqueListLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "a", "b"})

		// Act
		actual := args.Map{"len": len(c.UniqueListLock())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- UniqueListLock", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — Has / HasPtr / HasLock / HasAll / HasUsingSensitivity
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_Has_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Has", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"has": c.Has("a"),
			"missing": c.Has("x"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"missing": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Has", actual)
	})
}

func Test_Collection_Has_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Has_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"has": c.Has("a")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- Has empty", actual)
	})
}

func Test_Collection_HasLock(t *testing.T) {
	safeTest(t, "Test_I32_Collection_HasLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"has": c.HasLock("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HasLock", actual)
	})
}

func Test_Collection_HasPtr(t *testing.T) {
	safeTest(t, "Test_I32_Collection_HasPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"

		// Act
		actual := args.Map{
			"has": c.HasPtr(&s),
			"nil": c.HasPtr(nil),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HasPtr", actual)
	})
}

func Test_Collection_HasPtr_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_HasPtr_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		s := "a"

		// Act
		actual := args.Map{"has": c.HasPtr(&s)}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- HasPtr empty", actual)
	})
}

func Test_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_I32_Collection_HasAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"all": c.HasAll("a", "b"),
			"missing": c.HasAll("a", "x"),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"missing": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HasAll", actual)
	})
}

func Test_Collection_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_HasAll_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"has": c.HasAll("a")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- HasAll empty", actual)
	})
}

func Test_Collection_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_I32_Collection_HasUsingSensitivity", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"Hello"})

		// Act
		actual := args.Map{
			"sensitive": c.HasUsingSensitivity("hello", true),
			"insensitive": c.HasUsingSensitivity("hello", false),
		}

		// Assert
		expected := args.Map{
			"sensitive": false,
			"insensitive": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HasUsingSensitivity", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — IsContainsPtr / IsContainsAll / IsContainsAllLock / IsContainsAllSlice
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_IsContainsPtr(t *testing.T) {
	safeTest(t, "Test_I32_Collection_IsContainsPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"

		// Act
		actual := args.Map{
			"has": c.IsContainsPtr(&s),
			"nil": c.IsContainsPtr(nil),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IsContainsPtr", actual)
	})
}

func Test_Collection_IsContainsPtr_Empty_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_IsContainsPtr_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		s := "a"

		// Act
		actual := args.Map{"has": c.IsContainsPtr(&s)}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- IsContainsPtr empty", actual)
	})
}

func Test_Collection_IsContainsAll_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_IsContainsAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"all": c.IsContainsAll("a", "b"),
			"miss": c.IsContainsAll("a", "x"),
			"nil": c.IsContainsAll(nil...),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IsContainsAll", actual)
	})
}

func Test_Collection_IsContainsAllSlice_Empty_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_IsContainsAllSlice_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"empty": c.IsContainsAllSlice([]string{})}

		// Assert
		expected := args.Map{"empty": false}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- IsContainsAllSlice empty", actual)
	})
}

func Test_Collection_IsContainsAllSlice_EmptyCollection_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_IsContainsAllSlice_EmptyCollection", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"has": c.IsContainsAllSlice([]string{"a"})}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- IsContainsAllSlice empty collection", actual)
	})
}

func Test_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_I32_Collection_IsContainsAllLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"all": c.IsContainsAllLock("a", "b"),
			"nil": c.IsContainsAllLock(nil...),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IsContainsAllLock", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — GetHashsetPlusHasAll
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_GetHashsetPlusHasAll_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_GetHashsetPlusHasAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs, hasAll := c.GetHashsetPlusHasAll([]string{"a", "b"})

		// Act
		actual := args.Map{
			"hasAll": hasAll,
			"hsLen": hs.Length(),
		}

		// Assert
		expected := args.Map{
			"hasAll": true,
			"hsLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetHashsetPlusHasAll", actual)
	})
}

func Test_Collection_GetHashsetPlusHasAll_Nil(t *testing.T) {
	safeTest(t, "Test_I32_Collection_GetHashsetPlusHasAll_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		_, hasAll := c.GetHashsetPlusHasAll(nil)

		// Act
		actual := args.Map{"hasAll": hasAll}

		// Assert
		expected := args.Map{"hasAll": false}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- GetHashsetPlusHasAll nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — Sorted
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_I32_Collection_SortedListAsc", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		result := c.SortedListAsc()

		// Act
		actual := args.Map{
			"first": result[0],
			"last": result[2],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"last": "c",
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- SortedListAsc", actual)
	})
}

func Test_Collection_SortedListAsc_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_SortedListAsc_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"len": len(c.SortedListAsc())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- SortedListAsc empty", actual)
	})
}

func Test_Collection_SortedAsc(t *testing.T) {
	safeTest(t, "Test_I32_Collection_SortedAsc", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"c", "a"})
		c.SortedAsc()

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- SortedAsc", actual)
	})
}

func Test_Collection_SortedAsc_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_SortedAsc_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.SortedAsc()

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- SortedAsc empty", actual)
	})
}

func Test_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_I32_Collection_SortedAscLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"b", "a"})
		c.SortedAscLock()

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- SortedAscLock", actual)
	})
}

func Test_Collection_SortedAscLock_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_SortedAscLock_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.SortedAscLock()

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- SortedAscLock empty", actual)
	})
}

func Test_Collection_SortedListDsc_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_SortedListDsc", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "c", "b"})
		result := c.SortedListDsc()

		// Act
		actual := args.Map{
			"first": result[0],
			"last": result[2],
		}

		// Assert
		expected := args.Map{
			"first": "c",
			"last": "a",
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- SortedListDsc", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — Hashset conversions
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_HashsetAsIs_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_HashsetAsIs", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := c.HashsetAsIs()

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HashsetAsIs", actual)
	})
}

func Test_Collection_HashsetWithDoubleLength(t *testing.T) {
	safeTest(t, "Test_I32_Collection_HashsetWithDoubleLength", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetWithDoubleLength()

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- HashsetWithDoubleLength", actual)
	})
}

func Test_Collection_HashsetLock(t *testing.T) {
	safeTest(t, "Test_I32_Collection_HashsetLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetLock()

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HashsetLock", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — GetAllExcept / GetAllExceptCollection
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_GetAllExceptCollection_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_GetAllExceptCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		except := corestr.New.Collection.Strings([]string{"b"})
		result := c.GetAllExceptCollection(except)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetAllExceptCollection", actual)
	})
}

func Test_Collection_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_I32_Collection_GetAllExceptCollection_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- GetAllExceptCollection nil", actual)
	})
}

func Test_Collection_GetAllExceptCollection_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_GetAllExceptCollection_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.GetAllExceptCollection(corestr.New.Collection.Cap(5))

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- GetAllExceptCollection empty", actual)
	})
}

func Test_Collection_GetAllExcept_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_GetAllExcept", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.GetAllExcept([]string{"a"})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetAllExcept", actual)
	})
}

func Test_Collection_GetAllExcept_Nil_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_GetAllExcept_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.GetAllExcept(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- GetAllExcept nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — AddsAsync
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_AddsAsync_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AddsAsync", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddsAsync(wg, "a", "b")
		wg.Wait()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddsAsync", actual)
	})
}

func Test_Collection_AddsAsync_Nil_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AddsAsync_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddsAsync(nil, nil...)

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddsAsync nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — New / AddNonEmptyStrings / AddNonEmptyStringsSlice / AddFuncResult
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_New_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_New", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		n := c.New("a", "b")

		// Act
		actual := args.Map{"len": n.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- New", actual)
	})
}

func Test_Collection_New_Empty_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_New_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		n := c.New()

		// Act
		actual := args.Map{"empty": n.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- New empty", actual)
	})
}

func Test_Collection_AddNonEmptyStrings_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AddNonEmptyStrings", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmptyStrings("a", "", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct length -- AddNonEmptyStrings filters empty", actual)
	})
}

func Test_Collection_AddNonEmptyStrings_Empty_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AddNonEmptyStrings_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmptyStrings()

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddNonEmptyStrings empty", actual)
	})
}

func Test_Collection_AddFuncResult_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AddFuncResult", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddFuncResult(func() string { return "x" })

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "x"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddFuncResult", actual)
	})
}

func Test_Collection_AddFuncResult_Nil_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AddFuncResult_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddFuncResult(nil...)

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddFuncResult nil", actual)
	})
}

func Test_Collection_AddStringsByFuncChecking_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AddStringsByFuncChecking", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddStringsByFuncChecking([]string{"a", "bb", "c"}, func(s string) bool { return len(s) == 1 })

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddStringsByFuncChecking", actual)
	})
}

func Test_Collection_ExpandSlicePlusAdd_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_ExpandSlicePlusAdd", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.ExpandSlicePlusAdd([]string{"a,b"}, func(s string) []string {
			return []string{s + "!"}
		})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ExpandSlicePlusAdd", actual)
	})
}

func Test_Collection_MergeSlicesOfSlice_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_MergeSlicesOfSlice", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- MergeSlicesOfSlice", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — CharCollectionMap / SummaryString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_CharCollectionMap_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_CharCollectionMap", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"abc", "axy"})
		ccm := c.CharCollectionMap()

		// Act
		actual := args.Map{"notNil": ccm != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- CharCollectionMap", actual)
	})
}

func Test_Collection_SummaryString_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_SummaryString", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"notEmpty": c.SummaryString(1) != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- SummaryString", actual)
	})
}

func Test_Collection_SummaryStringWithHeader_Empty_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_SummaryStringWithHeader_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"notEmpty": c.SummaryStringWithHeader("H:") != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- SummaryStringWithHeader empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — String / StringLock / Join / JoinLine / Joins
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_String_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_String", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"notEmpty": c.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- String", actual)
	})
}

func Test_Collection_String_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_String_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"notEmpty": c.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- String empty", actual)
	})
}

func Test_Collection_StringLock(t *testing.T) {
	safeTest(t, "Test_I32_Collection_StringLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"notEmpty": c.StringLock() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- StringLock", actual)
	})
}

func Test_Collection_StringLock_Empty_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_StringLock_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"notEmpty": c.StringLock() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- StringLock empty", actual)
	})
}

func Test_Collection_Join_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Join", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"val": c.Join(",")}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Join", actual)
	})
}

func Test_Collection_Join_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Join_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"val": c.Join(",")}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- Join empty", actual)
	})
}

func Test_Collection_JoinLine(t *testing.T) {
	safeTest(t, "Test_I32_Collection_JoinLine", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"notEmpty": c.JoinLine() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JoinLine", actual)
	})
}

func Test_Collection_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_I32_Collection_JoinLine_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"val": c.JoinLine()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- JoinLine empty", actual)
	})
}

func Test_Collection_Joins_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Joins", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"val": c.Joins(",", "b")}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Joins", actual)
	})
}

func Test_Collection_Joins_NoExtra_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Joins_NoExtra", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"val": c.Joins(",")}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- Joins no extra", actual)
	})
}

func Test_Collection_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_I32_Collection_NonEmptyJoins", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})

		// Act
		actual := args.Map{"notEmpty": c.NonEmptyJoins(",") != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- NonEmptyJoins", actual)
	})
}

func Test_Collection_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_I32_Collection_NonWhitespaceJoins", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "  ", "b"})

		// Act
		actual := args.Map{"notEmpty": c.NonWhitespaceJoins(",") != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- NonWhitespaceJoins", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — CSV
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_Csv_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Csv", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"notEmpty": c.Csv() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Csv", actual)
	})
}

func Test_Collection_Csv_Empty_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Csv_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"val": c.Csv()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- Csv empty", actual)
	})
}

func Test_Collection_CsvOptions(t *testing.T) {
	safeTest(t, "Test_I32_Collection_CsvOptions", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"notEmpty": c.CsvOptions(true) != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- CsvOptions", actual)
	})
}

func Test_Collection_CsvOptions_Empty_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_CsvOptions_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"val": c.CsvOptions(false)}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- CsvOptions empty", actual)
	})
}

func Test_Collection_CsvLines_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_CsvLines", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(c.CsvLines())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- CsvLines", actual)
	})
}

func Test_Collection_CsvLinesOptions_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_CsvLinesOptions", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(c.CsvLinesOptions(true))}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- CsvLinesOptions", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection pt2 — JSON / Serialize / Deserialize
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_JsonModel(t *testing.T) {
	safeTest(t, "Test_I32_Collection_JsonModel", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"len": len(c.JsonModel()),
			"anyNotNil": c.JsonModelAny() != nil,
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"anyNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JsonModel", actual)
	})
}

func Test_Collection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_I32_Collection_MarshalJSON", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		b, err := c.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"notEmpty": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- MarshalJSON", actual)
	})
}

func Test_Collection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_I32_Collection_UnmarshalJSON", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		err := c.UnmarshalJSON([]byte(`["x","y"]`))

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": c.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- UnmarshalJSON", actual)
	})
}

func Test_Collection_UnmarshalJSON_Err(t *testing.T) {
	safeTest(t, "Test_I32_Collection_UnmarshalJSON_Err", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		err := c.UnmarshalJSON([]byte(`{bad`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- UnmarshalJSON err", actual)
	})
}

func Test_Collection_Json_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Json", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		j := c.Json()
		actual := args.Map{"notNil": j.Bytes != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Json", actual)
	})
}

func Test_Collection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_I32_Collection_JsonPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		jp := c.JsonPtr()
		actual := args.Map{"notNil": jp != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JsonPtr", actual)
	})
}

func Test_Collection_Serialize_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Serialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		b, err := c.Serialize()
		actual := args.Map{
			"noErr": err == nil,
			"notEmpty": len(b) > 0,
		}
		expected := args.Map{
			"noErr": true,
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Serialize", actual)
	})
}

func Test_Collection_Deserialize_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Deserialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		var target []string
		err := c.Deserialize(&target)
		actual := args.Map{
			"noErr": err == nil,
			"len": len(target),
		}
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Deserialize", actual)
	})
}

func Test_Collection_ParseInjectUsingJson_FromCollectionFilterIter(t *testing.T) {
	safeTest(t, "Test_I32_Collection_ParseInjectUsingJson", func() {
		src := corestr.New.Collection.Strings([]string{"a", "b"})
		jr := src.JsonPtr()
		target := corestr.New.Collection.Cap(5)
		result, err := target.ParseInjectUsingJson(jr)
		actual := args.Map{
			"noErr": err == nil,
			"len": result.Length(),
		}
		expected := args.Map{
			"noErr": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ParseInjectUsingJson", actual)
	})
}

func Test_Collection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_I32_Collection_JsonParseSelfInject", func() {
		src := corestr.New.Collection.Strings([]string{"a"})
		jr := src.JsonPtr()
		target := corestr.New.Collection.Cap(5)
		err := target.JsonParseSelfInject(jr)
		actual := args.Map{
			"noErr": err == nil,
			"len": target.Length(),
		}
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JsonParseSelfInject", actual)
	})
}

	// ══════════════════════════════════════════════════════════════════════════════
	// Collection pt2 — AsJsonMarshaller / AsJsonContractsBinder
	// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AsJsonMarshaller", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"notNil": c.AsJsonMarshaller() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AsJsonMarshaller", actual)
	})
}

func Test_Collection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_I32_Collection_AsJsonContractsBinder", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"notNil": c.AsJsonContractsBinder() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AsJsonContractsBinder", actual)
	})
}

	// ══════════════════════════════════════════════════════════════════════════════
	// Collection pt2 — Clear / Dispose
	// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_Clear(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Clear", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Clear()
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Clear", actual)
	})
}

func Test_Collection_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Clear_Nil", func() {
		var c *corestr.Collection
		result := c.Clear()
		actual := args.Map{"nil": result == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- Clear nil", actual)
	})
}

func Test_Collection_Dispose(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Dispose", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Dispose()
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Dispose", actual)
	})
}

func Test_Collection_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_I32_Collection_Dispose_Nil", func() {
		var c *corestr.Collection
		c.Dispose() // should not panic
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- Dispose nil", actual)
	})
}
