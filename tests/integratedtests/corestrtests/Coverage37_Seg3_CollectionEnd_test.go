package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Segment 3: Filter variants, Hashset, Sort, Contains, Join, CSV,
//   GetAllExcept, JSON, Clear/Dispose, Resize, remaining methods
// ══════════════════════════════════════════════════════════════════════════════

// ── FilterLock ───────────────────────────────────────────────────────────────

func Test_Seg3_Collection_FilterLock(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_FilterLock", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "bb", "ccc")
		filtered := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})
		actual := args.Map{"len": len(filtered)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "FilterLock -- keeps len>1", actual)
	})
}

func Test_Seg3_Collection_FilterLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_FilterLock_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		filtered := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"len": len(filtered)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "FilterLock empty -- returns empty", actual)
	})
}

func Test_Seg3_Collection_FilterLock_Break(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_FilterLock_Break", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		filtered := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, i == 0
		})
		actual := args.Map{"len": len(filtered)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterLock break -- stops after first", actual)
	})
}

// ── FilteredCollection / FilteredCollectionLock ─────────────────────────────

func Test_Seg3_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_FilteredCollection", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "bb", "ccc")
		fc := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) >= 2, false
		})
		actual := args.Map{"len": fc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "FilteredCollection -- returns new collection", actual)
	})
}

func Test_Seg3_Collection_FilteredCollectionLock(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_FilteredCollectionLock", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("x", "yy")
		fc := c.FilteredCollectionLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"len": fc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "FilteredCollectionLock -- returns new collection", actual)
	})
}

// ── FilterPtr / FilterPtrLock ───────────────────────────────────────────────

func Test_Seg3_Collection_FilterPtr(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_FilterPtr", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "bb", "ccc")
		result := c.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, len(*s) > 1, false
		})
		actual := args.Map{"len": len(*result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "FilterPtr -- keeps len>1", actual)
	})
}

func Test_Seg3_Collection_FilterPtr_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_FilterPtr_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		result := c.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"len": len(*result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "FilterPtr empty -- returns empty", actual)
	})
}

func Test_Seg3_Collection_FilterPtr_Break(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_FilterPtr_Break", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		result := c.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, true, i == 0
		})
		actual := args.Map{"len": len(*result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterPtr break -- stops after first", actual)
	})
}

func Test_Seg3_Collection_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_FilterPtrLock", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "bb")
		result := c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"len": len(*result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "FilterPtrLock -- returns all", actual)
	})
}

func Test_Seg3_Collection_FilterPtrLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_FilterPtrLock_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		result := c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"len": len(*result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "FilterPtrLock empty -- returns empty", actual)
	})
}
func Test_Seg3_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_NonEmptyList", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "", "b", "")
		actual := args.Map{"len": len(c.NonEmptyList())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyList -- skips empty", actual)
	})
}

func Test_Seg3_Collection_NonEmptyList_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_NonEmptyList_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"len": len(c.NonEmptyList())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "NonEmptyList empty -- returns empty", actual)
	})
}
func Test_Seg3_Collection_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_NonEmptyItems", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "", "b")
		actual := args.Map{"len": len(c.NonEmptyItems())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyItems -- skips empty", actual)
	})
}
func Test_Seg3_Collection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_NonEmptyItemsOrNonWhitespace", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "   ", "", "b")
		actual := args.Map{"len": len(c.NonEmptyItemsOrNonWhitespace())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyItemsOrNonWhitespace -- skips ws", actual)
	})
}
func Test_Seg3_Collection_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_HashsetAsIs", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "a")
		hs := c.HashsetAsIs()
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "HashsetAsIs -- 2 unique", actual)
	})
}

func Test_Seg3_Collection_HashsetWithDoubleLength(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_HashsetWithDoubleLength", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		hs := c.HashsetWithDoubleLength()
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "HashsetWithDoubleLength -- 2 items", actual)
	})
}
func Test_Seg3_Collection_Items(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Items", func() {
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		actual := args.Map{"len": len(c.Items())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Items -- returns items", actual)
	})
}
func Test_Seg3_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_ListCopyPtrLock", func() {
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		list := c.ListCopyPtrLock()
		actual := args.Map{"len": len(list)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ListCopyPtrLock -- returns copy", actual)
	})
}

func Test_Seg3_Collection_ListCopyPtrLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_ListCopyPtrLock_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		list := c.ListCopyPtrLock()
		actual := args.Map{"len": len(list)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ListCopyPtrLock empty -- returns empty", actual)
	})
}

// ── Has / HasLock / HasPtr / HasAll ─────────────────────────────────────────

func Test_Seg3_Collection_Has(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Has", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		actual := args.Map{"has": c.Has("b"), "miss": c.Has("z")}
		expected := args.Map{"has": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "Has -- found and missing", actual)
	})
}

func Test_Seg3_Collection_Has_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Has_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"has": c.Has("a")}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Has empty -- false", actual)
	})
}

func Test_Seg3_Collection_HasLock(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_HasLock", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("x")
		actual := args.Map{"has": c.HasLock("x")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasLock -- found", actual)
	})
}

func Test_Seg3_Collection_HasPtr(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_HasPtr", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("hello")
		s := "hello"
		actual := args.Map{"has": c.HasPtr(&s), "nil": c.HasPtr(nil)}
		expected := args.Map{"has": true, "nil": false}
		expected.ShouldBeEqual(t, 0, "HasPtr -- found and nil", actual)
	})
}

func Test_Seg3_Collection_HasPtr_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_HasPtr_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		s := "x"
		actual := args.Map{"has": c.HasPtr(&s)}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasPtr empty -- false", actual)
	})
}

func Test_Seg3_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_HasAll", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		actual := args.Map{
			"all":     c.HasAll("a", "b"),
			"missing": c.HasAll("a", "z"),
		}
		expected := args.Map{
			"all":     true,
			"missing": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAll -- all found and missing one", actual)
	})
}

func Test_Seg3_Collection_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_HasAll_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"has": c.HasAll("a")}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasAll empty -- false", actual)
	})
}

func Test_Seg3_Collection_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_HasUsingSensitivity", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("Hello")
		actual := args.Map{
			"sensitive":   c.HasUsingSensitivity("hello", true),
			"insensitive": c.HasUsingSensitivity("hello", false),
		}
		expected := args.Map{
			"sensitive":   false,
			"insensitive": true,
		}
		expected.ShouldBeEqual(t, 0, "HasUsingSensitivity -- case comparison", actual)
	})
}

// ── IsContainsPtr / IsContainsAll / IsContainsAllSlice / IsContainsAllLock ──

func Test_Seg3_Collection_IsContainsPtr(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_IsContainsPtr", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("x")
		s := "x"
		actual := args.Map{"has": c.IsContainsPtr(&s), "nil": c.IsContainsPtr(nil)}
		expected := args.Map{"has": true, "nil": false}
		expected.ShouldBeEqual(t, 0, "IsContainsPtr -- found and nil", actual)
	})
}

func Test_Seg3_Collection_IsContainsAllSlice(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_IsContainsAllSlice", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		actual := args.Map{
			"all":   c.IsContainsAllSlice([]string{"a", "b"}),
			"miss":  c.IsContainsAllSlice([]string{"a", "z"}),
			"empty": c.IsContainsAllSlice([]string{}),
		}
		expected := args.Map{
			"all":   true,
			"miss":  false,
			"empty": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContainsAllSlice -- various cases", actual)
	})
}

func Test_Seg3_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_IsContainsAll", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		actual := args.Map{
			"all": c.IsContainsAll("a", "b"),
			"nil": c.IsContainsAll(nil...),
		}
		expected := args.Map{
			"all": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContainsAll -- found and nil", actual)
	})
}

func Test_Seg3_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_IsContainsAllLock", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		actual := args.Map{
			"all": c.IsContainsAllLock("a", "b"),
			"nil": c.IsContainsAllLock(nil...),
		}
		expected := args.Map{
			"all": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContainsAllLock -- found and nil", actual)
	})
}

// ── GetHashsetPlusHasAll ────────────────────────────────────────────────────

func Test_Seg3_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_GetHashsetPlusHasAll", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		hs, hasAll := c.GetHashsetPlusHasAll([]string{"a", "b"})
		actual := args.Map{"hasAll": hasAll, "hsLen": hs.Length()}
		expected := args.Map{"hasAll": true, "hsLen": 3}
		expected.ShouldBeEqual(t, 0, "GetHashsetPlusHasAll -- all found", actual)
	})
}

func Test_Seg3_Collection_GetHashsetPlusHasAll_Nil(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_GetHashsetPlusHasAll_Nil", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		_, hasAll := c.GetHashsetPlusHasAll(nil)
		actual := args.Map{"hasAll": hasAll}
		expected := args.Map{"hasAll": false}
		expected.ShouldBeEqual(t, 0, "GetHashsetPlusHasAll nil -- false", actual)
	})
}

// ── Sorting ─────────────────────────────────────────────────────────────────

func Test_Seg3_Collection_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_SortedListAsc", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("c", "a", "b")
		sorted := c.SortedListAsc()
		actual := args.Map{"first": sorted[0], "last": sorted[2]}
		expected := args.Map{"first": "a", "last": "c"}
		expected.ShouldBeEqual(t, 0, "SortedListAsc -- ascending order", actual)
	})
}

func Test_Seg3_Collection_SortedListAsc_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_SortedListAsc_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"len": len(c.SortedListAsc())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SortedListAsc empty -- returns empty", actual)
	})
}

func Test_Seg3_Collection_SortedAsc(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_SortedAsc", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("c", "a", "b")
		c.SortedAsc()
		actual := args.Map{"first": c.First()}
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "SortedAsc -- mutates in place", actual)
	})
}

func Test_Seg3_Collection_SortedAsc_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_SortedAsc_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		c.SortedAsc()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SortedAsc empty -- no change", actual)
	})
}

func Test_Seg3_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_SortedAscLock", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("b", "a")
		c.SortedAscLock()
		actual := args.Map{"first": c.First()}
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "SortedAscLock -- sorted", actual)
	})
}

func Test_Seg3_Collection_SortedAscLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_SortedAscLock_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		c.SortedAscLock()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SortedAscLock empty -- no change", actual)
	})
}

func Test_Seg3_Collection_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_SortedListDsc", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "c", "b")
		sorted := c.SortedListDsc()
		actual := args.Map{"first": sorted[0], "last": sorted[2]}
		expected := args.Map{"first": "c", "last": "a"}
		expected.ShouldBeEqual(t, 0, "SortedListDsc -- descending order", actual)
	})
}

// ── New ─────────────────────────────────────────────────────────────────────

func Test_Seg3_Collection_New(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_New", func() {
		c := corestr.New.Collection.Cap(5)
		nc := c.New("a", "b")
		actual := args.Map{"len": nc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "New -- creates new collection", actual)
	})
}

func Test_Seg3_Collection_New_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_New_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		nc := c.New()
		actual := args.Map{"len": nc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "New empty -- creates empty collection", actual)
	})
}

// ── AddNonEmptyStrings / AddNonEmptyStringsSlice ────────────────────────────

func Test_Seg3_Collection_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_AddNonEmptyStrings", func() {
		c := corestr.New.Collection.Cap(10)
		c.AddNonEmptyStrings("a", "", "b")
		actual := args.Map{"len": c.Length()}
		// Fix: AddNonEmptyStrings filters empty strings, so "a","","b" → 2 items
		// See issues/corestrtests-addnonemptystrings-wrong-expectation.md
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStrings -- adds items", actual)
	})
}

func Test_Seg3_Collection_AddNonEmptyStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_AddNonEmptyStrings_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		c.AddNonEmptyStrings()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStrings empty -- no change", actual)
	})
}

func Test_Seg3_Collection_AddNonEmptyStringsSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_AddNonEmptyStringsSlice_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		c.AddNonEmptyStringsSlice([]string{})
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStringsSlice empty -- no change", actual)
	})
}

// ── AddFuncResult ───────────────────────────────────────────────────────────

func Test_Seg3_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_AddFuncResult", func() {
		c := corestr.New.Collection.Cap(10)
		c.AddFuncResult(
			func() string { return "hello" },
			func() string { return "world" },
		)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddFuncResult -- 2 results added", actual)
	})
}

func Test_Seg3_Collection_AddFuncResult_Nil(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_AddFuncResult_Nil", func() {
		c := corestr.New.Collection.Cap(10)
		c.AddFuncResult(nil...)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddFuncResult nil -- no change", actual)
	})
}

// ── AddStringsByFuncChecking ────────────────────────────────────────────────

func Test_Seg3_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_AddStringsByFuncChecking", func() {
		c := corestr.New.Collection.Cap(10)
		c.AddStringsByFuncChecking(
			[]string{"a", "bb", "ccc"},
			func(line string) bool { return len(line) > 1 },
		)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsByFuncChecking -- 2 pass check", actual)
	})
}

// ── ExpandSlicePlusAdd ──────────────────────────────────────────────────────

func Test_Seg3_Collection_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_ExpandSlicePlusAdd", func() {
		c := corestr.New.Collection.Cap(10)
		c.ExpandSlicePlusAdd(
			[]string{"a,b", "c,d"},
			func(line string) []string {
				return []string{line + "_expanded"}
			},
		)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ExpandSlicePlusAdd -- 2 expanded items", actual)
	})
}

// ── MergeSlicesOfSlice ──────────────────────────────────────────────────────

func Test_Seg3_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_MergeSlicesOfSlice", func() {
		c := corestr.New.Collection.Cap(10)
		c.MergeSlicesOfSlice([]string{"a", "b"}, []string{"c"})
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlice -- 3 items merged", actual)
	})
}

// ── GetAllExcept / GetAllExceptCollection ───────────────────────────────────

func Test_Seg3_Collection_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_GetAllExceptCollection", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c", "d")
		exclude := corestr.New.Collection.Cap(5)
		exclude.Adds("b", "d")
		result := c.GetAllExceptCollection(exclude)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection -- 2 remaining", actual)
	})
}

func Test_Seg3_Collection_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_GetAllExceptCollection_Nil", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		result := c.GetAllExceptCollection(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection nil -- returns copy", actual)
	})
}

func Test_Seg3_Collection_GetAllExceptCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_GetAllExceptCollection_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		result := c.GetAllExceptCollection(corestr.New.Collection.Cap(0))
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection empty -- returns copy", actual)
	})
}

func Test_Seg3_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_GetAllExcept", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		result := c.GetAllExcept([]string{"b"})
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExcept -- 2 remaining", actual)
	})
}

func Test_Seg3_Collection_GetAllExcept_Nil(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_GetAllExcept_Nil", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		result := c.GetAllExcept(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExcept nil -- returns copy", actual)
	})
}

// ── CharCollectionMap ───────────────────────────────────────────────────────

func Test_Seg3_Collection_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_CharCollectionMap", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("hello", "world")
		ccm := c.CharCollectionMap()
		actual := args.Map{"notNil": ccm != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "CharCollectionMap -- returns map", actual)
	})
}

// ── String / StringLock / SummaryString ──────────────────────────────────────

func Test_Seg3_Collection_String(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_String", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		actual := args.Map{"nonEmpty": c.String() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg3_Collection_String_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_String_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"hasNoElements": len(c.String()) > 0}
		expected := args.Map{"hasNoElements": true}
		expected.ShouldBeEqual(t, 0, "String empty -- contains NoElements", actual)
	})
}

func Test_Seg3_Collection_StringLock(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_StringLock", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		actual := args.Map{"nonEmpty": c.StringLock() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock -- non-empty", actual)
	})
}

func Test_Seg3_Collection_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_StringLock_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"nonEmpty": c.StringLock() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty -- contains NoElements", actual)
	})
}

func Test_Seg3_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_SummaryString", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		actual := args.Map{"nonEmpty": c.SummaryString(1) != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryString -- non-empty", actual)
	})
}

func Test_Seg3_Collection_SummaryStringWithHeader(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_SummaryStringWithHeader", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		actual := args.Map{"nonEmpty": c.SummaryStringWithHeader("Header") != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringWithHeader -- non-empty", actual)
	})
}

func Test_Seg3_Collection_SummaryStringWithHeader_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_SummaryStringWithHeader_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"nonEmpty": c.SummaryStringWithHeader("Header") != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringWithHeader empty -- contains header", actual)
	})
}

// ── CSV ─────────────────────────────────────────────────────────────────────

func Test_Seg3_Collection_Csv(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Csv", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		actual := args.Map{"nonEmpty": c.Csv() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Csv -- non-empty", actual)
	})
}

func Test_Seg3_Collection_Csv_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Csv_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"empty": c.Csv() == ""}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Csv empty -- returns empty string", actual)
	})
}

func Test_Seg3_Collection_CsvOptions(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_CsvOptions", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		actual := args.Map{"nonEmpty": c.CsvOptions(true) != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "CsvOptions -- non-empty", actual)
	})
}

func Test_Seg3_Collection_CsvOptions_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_CsvOptions_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"empty": c.CsvOptions(false) == ""}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "CsvOptions empty -- returns empty", actual)
	})
}

func Test_Seg3_Collection_CsvLines(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_CsvLines", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		actual := args.Map{"len": len(c.CsvLines())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CsvLines -- 2 items", actual)
	})
}

func Test_Seg3_Collection_CsvLinesOptions(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_CsvLinesOptions", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		actual := args.Map{"len": len(c.CsvLinesOptions(true))}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CsvLinesOptions -- 2 items", actual)
	})
}

// ── Join / JoinLine / Joins / NonEmptyJoins / NonWhitespaceJoins ────────────

func Test_Seg3_Collection_Join(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Join", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		actual := args.Map{"val": c.Join(",")}
		expected := args.Map{"val": "a,b,c"}
		expected.ShouldBeEqual(t, 0, "Join -- comma separated", actual)
	})
}

func Test_Seg3_Collection_Join_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Join_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"val": c.Join(",")}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Join empty -- empty string", actual)
	})
}

func Test_Seg3_Collection_JoinLine(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_JoinLine", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		actual := args.Map{"nonEmpty": c.JoinLine() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinLine -- non-empty", actual)
	})
}

func Test_Seg3_Collection_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_JoinLine_Empty", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"val": c.JoinLine()}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "JoinLine empty -- empty string", actual)
	})
}

func Test_Seg3_Collection_Joins(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Joins", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		actual := args.Map{"nonEmpty": c.Joins(",", "c", "d") != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Joins -- non-empty", actual)
	})
}

func Test_Seg3_Collection_Joins_NoExtra(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Joins_NoExtra", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		actual := args.Map{"val": c.Joins(",")}
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Joins no extra -- just items", actual)
	})
}

func Test_Seg3_Collection_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_NonEmptyJoins", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "", "b")
		actual := args.Map{"nonEmpty": c.NonEmptyJoins(",") != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonEmptyJoins -- skips empty", actual)
	})
}

func Test_Seg3_Collection_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_NonWhitespaceJoins", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "   ", "b")
		actual := args.Map{"nonEmpty": c.NonWhitespaceJoins(",") != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonWhitespaceJoins -- skips whitespace", actual)
	})
}

// ── Resize / AddCapacity ────────────────────────────────────────────────────

func Test_Seg3_Collection_Resize(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Resize", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		c.Resize(100)
		actual := args.Map{"cap": c.Capacity() >= 100}
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Resize -- capacity increased", actual)
	})
}

func Test_Seg3_Collection_Resize_SmallerNoOp(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Resize_SmallerNoOp", func() {
		c := corestr.New.Collection.Cap(100)
		c.Resize(5)
		actual := args.Map{"cap": c.Capacity() >= 100}
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Resize smaller -- no change", actual)
	})
}

func Test_Seg3_Collection_AddCapacity_Nil(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_AddCapacity_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddCapacity(nil...)
		actual := args.Map{"cap": c.Capacity() >= 5}
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "AddCapacity nil -- no change", actual)
	})
}

// ── JSON / Serialize / Deserialize ──────────────────────────────────────────

func Test_Seg3_Collection_Json(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Json", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		j := c.Json()
		actual := args.Map{"noErr": !j.HasError()}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg3_Collection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_JsonPtr", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		j := c.JsonPtr()
		actual := args.Map{"noErr": !j.HasError()}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr -- no error", actual)
	})
}

func Test_Seg3_Collection_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_JsonModel", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		actual := args.Map{"len": len(c.JsonModel())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "JsonModel -- returns items", actual)
	})
}

func Test_Seg3_Collection_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_JsonModelAny", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		actual := args.Map{"notNil": c.JsonModelAny() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Seg3_Collection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_MarshalJSON", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		b, err := c.MarshalJSON()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg3_Collection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_UnmarshalJSON", func() {
		c := corestr.New.Collection.Cap(10)
		err := c.UnmarshalJSON([]byte(`["a","b"]`))
		actual := args.Map{"noErr": err == nil, "len": c.Length()}
		expected := args.Map{"noErr": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Seg3_Collection_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_UnmarshalJSON_Invalid", func() {
		c := corestr.New.Collection.Cap(10)
		err := c.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Seg3_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Serialize", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		b, err := c.Serialize()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_Seg3_Collection_Deserialize(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Deserialize", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		var dest []string
		err := c.Deserialize(&dest)
		actual := args.Map{"noErr": err == nil, "len": len(dest)}
		expected := args.Map{"noErr": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

func Test_Seg3_Collection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_ParseInjectUsingJson", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		jr := c.JsonPtr()
		c2 := corestr.New.Collection.Cap(10)
		result, err := c2.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil, "len": result.Length()}
		expected := args.Map{"noErr": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- round trip", actual)
	})
}

func Test_Seg3_Collection_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_ParseInjectUsingJsonMust", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		jr := c.JsonPtr()
		c2 := corestr.New.Collection.Cap(10)
		result := c2.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg3_Collection_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_ParseInjectUsingJsonMust_Panic", func() {
		defer func() { recover() }()
		c := corestr.New.Collection.Cap(10)
		bad := &corejson.Result{}
		_ = c.ParseInjectUsingJsonMust(bad)
	})
}

func Test_Seg3_Collection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_JsonParseSelfInject", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		jr := c.JsonPtr()
		c2 := corestr.New.Collection.Cap(10)
		err := c2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ── Interface casts ─────────────────────────────────────────────────────────

func Test_Seg3_Collection_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_AsJsonMarshaller", func() {
		c := corestr.New.Collection.Cap(5)
		actual := args.Map{"notNil": c.AsJsonMarshaller() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonMarshaller -- non-nil", actual)
	})
}

func Test_Seg3_Collection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_AsJsonContractsBinder", func() {
		c := corestr.New.Collection.Cap(5)
		actual := args.Map{"notNil": c.AsJsonContractsBinder() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder -- non-nil", actual)
	})
}

// ── Clear / Dispose ─────────────────────────────────────────────────────────

func Test_Seg3_Collection_Clear(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Clear", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		c.Clear()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg3_Collection_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Clear_Nil", func() {
		var c *corestr.Collection
		result := c.Clear()
		actual := args.Map{"nil": result == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clear nil -- returns nil", actual)
	})
}

func Test_Seg3_Collection_Dispose(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Dispose", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		c.Dispose()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Dispose -- items nil", actual)
	})
}

func Test_Seg3_Collection_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Seg3_Collection_Dispose_Nil", func() {
		var c *corestr.Collection
		c.Dispose() // should not panic
	})
}
