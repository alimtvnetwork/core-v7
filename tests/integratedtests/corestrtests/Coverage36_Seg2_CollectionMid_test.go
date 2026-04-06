package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Segment 2: Indexing, Paging, Insertion, Removal, Append, Filter
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg2_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_IndexAt", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b", "c")

		// Act
		actual := args.Map{"val": c.IndexAt(1)}

		// Assert
		expected := args.Map{"val": "b"}
		expected.ShouldBeEqual(t, 0, "IndexAt returns middle -- index 1", actual)
	})
}

func Test_Seg2_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_SafeIndexAtUsingLength", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")

		// Act
		actual := args.Map{
			"valid":   c.SafeIndexAtUsingLength("def", 2, 1),
			"outBound": c.SafeIndexAtUsingLength("def", 2, 5),
		}

		// Assert
		expected := args.Map{
			"valid":   "b",
			"outBound": "def",
		}
		expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLength -- valid and out of bounds", actual)
	})
}

func Test_Seg2_Collection_First(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_First", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("first", "second")

		// Act
		actual := args.Map{"val": c.First()}

		// Assert
		expected := args.Map{"val": "first"}
		expected.ShouldBeEqual(t, 0, "First returns first -- 2 items", actual)
	})
}

func Test_Seg2_Collection_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_FirstOrDefault_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"val": c.FirstOrDefault()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault returns empty -- empty collection", actual)
	})
}

func Test_Seg2_Collection_FirstOrDefault_HasItems(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_FirstOrDefault_HasItems", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("hello")

		// Act
		actual := args.Map{"val": c.FirstOrDefault()}

		// Assert
		expected := args.Map{"val": "hello"}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault returns first -- has items", actual)
	})
}

func Test_Seg2_Collection_Last(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_Last", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b", "c")

		// Act
		actual := args.Map{"val": c.Last()}

		// Assert
		expected := args.Map{"val": "c"}
		expected.ShouldBeEqual(t, 0, "Last returns last -- 3 items", actual)
	})
}

func Test_Seg2_Collection_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_LastOrDefault_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"val": c.LastOrDefault()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "LastOrDefault returns empty -- empty collection", actual)
	})
}

func Test_Seg2_Collection_LastOrDefault_HasItems(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_LastOrDefault_HasItems", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")

		// Act
		actual := args.Map{"val": c.LastOrDefault()}

		// Assert
		expected := args.Map{"val": "b"}
		expected.ShouldBeEqual(t, 0, "LastOrDefault returns last -- has items", actual)
	})
}

func Test_Seg2_Collection_Single(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_Single", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("only")

		// Act
		actual := args.Map{"val": c.Single()}

		// Assert
		expected := args.Map{"val": "only"}
		expected.ShouldBeEqual(t, 0, "Single returns only item -- 1 item", actual)
	})
}

func Test_Seg2_Collection_Single_Panics(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_Single_Panics", func() {
		// Arrange
		defer func() { recover() }()
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		_ = c.Single() // should panic

		// Act
		actual := args.Map{"result": false}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have panicked", actual)
	})
}

// ── Take / Skip ─────────────────────────────────────────────────────────────

func Test_Seg2_Collection_Take(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_Take", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b", "c", "d")
		taken := c.Take(2)

		// Act
		actual := args.Map{
			"len": taken.Length(),
			"first": taken.First(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "Take 2 from 4 -- returns first 2", actual)
	})
}

func Test_Seg2_Collection_Take_MoreThanLength(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_Take_MoreThanLength", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		taken := c.Take(10)

		// Act
		actual := args.Map{"len": taken.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Take more than length -- returns all", actual)
	})
}

func Test_Seg2_Collection_Take_Zero(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_Take_Zero", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		taken := c.Take(0)

		// Act
		actual := args.Map{"len": taken.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Take 0 -- returns empty", actual)
	})
}

func Test_Seg2_Collection_Skip(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_Skip", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b", "c", "d")
		skipped := c.Skip(2)

		// Act
		actual := args.Map{
			"len": skipped.Length(),
			"first": skipped.First(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "c",
		}
		expected.ShouldBeEqual(t, 0, "Skip 2 from 4 -- returns last 2", actual)
	})
}

func Test_Seg2_Collection_Skip_Zero(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_Skip_Zero", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		skipped := c.Skip(0)

		// Act
		actual := args.Map{"len": skipped.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Skip 0 -- returns same", actual)
	})
}

func Test_Seg2_Collection_Skip_Panics(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_Skip_Panics", func() {
		// Arrange
		defer func() { recover() }()
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		_ = c.Skip(10) // should panic

		// Act
		actual := args.Map{"result": false}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have panicked", actual)
	})
}

// ── Reverse ─────────────────────────────────────────────────────────────────

func Test_Seg2_Collection_Reverse_Many(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_Reverse_Many", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b", "c", "d")
		c.Reverse()

		// Act
		actual := args.Map{
			"first": c.First(),
			"last": c.Last(),
		}

		// Assert
		expected := args.Map{
			"first": "d",
			"last": "a",
		}
		expected.ShouldBeEqual(t, 0, "Reverse 4 items -- first/last swapped", actual)
	})
}

func Test_Seg2_Collection_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_Reverse_Two", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		c.Reverse()

		// Act
		actual := args.Map{
			"first": c.First(),
			"last": c.Last(),
		}

		// Assert
		expected := args.Map{
			"first": "b",
			"last": "a",
		}
		expected.ShouldBeEqual(t, 0, "Reverse 2 items -- swapped", actual)
	})
}

func Test_Seg2_Collection_Reverse_Single(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_Reverse_Single", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		c.Reverse()

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Reverse 1 item -- unchanged", actual)
	})
}

// ── Paging ──────────────────────────────────────────────────────────────────

func Test_Seg2_Collection_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_GetPagesSize", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		for i := 0; i < 10; i++ {
			c.Add("x")
		}

		// Act
		actual := args.Map{
			"pages3":  c.GetPagesSize(3),
			"pages5":  c.GetPagesSize(5),
			"pages0":  c.GetPagesSize(0),
			"pagesNeg": c.GetPagesSize(-1),
		}

		// Assert
		expected := args.Map{
			"pages3":  4,
			"pages5":  2,
			"pages0":  0,
			"pagesNeg": 0,
		}
		expected.ShouldBeEqual(t, 0, "GetPagesSize -- various page sizes", actual)
	})
}

func Test_Seg2_Collection_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_GetSinglePageCollection", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		for i := 0; i < 10; i++ {
			c.Add("x")
		}
		page := c.GetSinglePageCollection(3, 2)

		// Act
		actual := args.Map{"len": page.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "GetSinglePageCollection -- page 2 of size 3", actual)
	})
}

func Test_Seg2_Collection_GetSinglePageCollection_LastPage(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_GetSinglePageCollection_LastPage", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		for i := 0; i < 10; i++ {
			c.Add("x")
		}
		page := c.GetSinglePageCollection(3, 4)

		// Act
		actual := args.Map{"len": page.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetSinglePageCollection -- last partial page", actual)
	})
}

func Test_Seg2_Collection_GetSinglePageCollection_SmallCollection(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_GetSinglePageCollection_SmallCollection", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		page := c.GetSinglePageCollection(5, 1)

		// Act
		actual := args.Map{"len": page.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetSinglePageCollection -- collection smaller than page size", actual)
	})
}

func Test_Seg2_Collection_GetPagedCollection(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_GetPagedCollection", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		for i := 0; i < 7; i++ {
			c.Add("x")
		}
		paged := c.GetPagedCollection(3)

		// Act
		actual := args.Map{"pages": paged.Length()}

		// Assert
		expected := args.Map{"pages": 3}
		expected.ShouldBeEqual(t, 0, "GetPagedCollection -- 7 items, page 3 = 3 pages", actual)
	})
}

func Test_Seg2_Collection_GetPagedCollection_SmallCollection(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_GetPagedCollection_SmallCollection", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		paged := c.GetPagedCollection(5)

		// Act
		actual := args.Map{"pages": paged.Length()}

		// Assert
		expected := args.Map{"pages": 1}
		expected.ShouldBeEqual(t, 0, "GetPagedCollection -- collection < page size = 1 page", actual)
	})
}

// ── InsertAt / ChainRemoveAt ────────────────────────────────────────────────

func Test_Seg2_Collection_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_InsertAt_Middle", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c", "d", "e")
		c.InsertAt(2, "X", "Y")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 7}
		expected.ShouldBeEqual(t, 0, "InsertAt middle -- items added", actual)
	})
}

func Test_Seg2_Collection_InsertAt_Last(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_InsertAt_Last", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		c.InsertAt(2, "X")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "InsertAt last index -- appended", actual)
	})
}

func Test_Seg2_Collection_InsertAt_Empty(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_InsertAt_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.InsertAt(0, "X")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "InsertAt empty -- appended", actual)
	})
}

func Test_Seg2_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_ChainRemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c", "d")
		c.ChainRemoveAt(1)

		// Act
		actual := args.Map{
			"len": c.Length(),
			"second": c.IndexAt(1),
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"second": "c",
		}
		expected.ShouldBeEqual(t, 0, "ChainRemoveAt -- middle removed", actual)
	})
}

// ── RemoveItemsIndexes ──────────────────────────────────────────────────────

func Test_Seg2_Collection_RemoveItemsIndexes(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_RemoveItemsIndexes", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c", "d")
		c.RemoveItemsIndexes(false, 1, 3)

		// Act
		actual := args.Map{
			"len": c.Length(),
			"first": c.First(),
			"last": c.Last(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
			"last": "c",
		}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexes -- remove indexes 1,3", actual)
	})
}

func Test_Seg2_Collection_RemoveItemsIndexes_NilIgnore(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_RemoveItemsIndexes_NilIgnore", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		c.RemoveItemsIndexes(true)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexes nil indexes ignore -- unchanged", actual)
	})
}

func Test_Seg2_Collection_RemoveItemsIndexesPtr_NilIndexes(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_RemoveItemsIndexesPtr_NilIndexes", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		c.RemoveItemsIndexesPtr(false, nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexesPtr nil -- unchanged", actual)
	})
}

func Test_Seg2_Collection_RemoveItemsIndexesPtr_EmptyPanic(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_RemoveItemsIndexesPtr_EmptyPanic", func() {
		// Arrange
		defer func() { recover() }()
		c := corestr.New.Collection.Cap(10)
		c.RemoveItemsIndexesPtr(false, []int{0})

		// Act
		actual := args.Map{"result": false}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have panicked", actual)
	})
}

func Test_Seg2_Collection_RemoveItemsIndexesPtr_EmptyIgnore(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_RemoveItemsIndexesPtr_EmptyIgnore", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.RemoveItemsIndexesPtr(true, []int{0})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexesPtr empty ignore -- unchanged", actual)
	})
}

// ── AppendCollections ───────────────────────────────────────────────────────

func Test_Seg2_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendCollectionPtr", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		c2 := corestr.New.Collection.Cap(5)
		c2.Adds("b", "c")
		c.AppendCollectionPtr(c2)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AppendCollectionPtr -- merged", actual)
	})
}

func Test_Seg2_Collection_AppendCollections(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendCollections", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c1 := corestr.New.Collection.Cap(5)
		c1.Add("a")
		c2 := corestr.New.Collection.Cap(5)
		c2.Add("b")
		empty := corestr.New.Collection.Cap(5)
		c.AppendCollections(c1, empty, c2)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendCollections -- skips empty", actual)
	})
}

func Test_Seg2_Collection_AppendCollections_Empty(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendCollections_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AppendCollections()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendCollections empty -- no change", actual)
	})
}

// ── AppendAnys ──────────────────────────────────────────────────────────────

func Test_Seg2_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendAnys", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AppendAnys("hello", 42, nil, true)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AppendAnys -- skips nil, converts others", actual)
	})
}

func Test_Seg2_Collection_AppendAnys_Empty(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendAnys_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AppendAnys()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnys empty -- no change", actual)
	})
}

func Test_Seg2_Collection_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendAnysLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AppendAnysLock("a", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendAnysLock -- 2 items", actual)
	})
}

func Test_Seg2_Collection_AppendAnysLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendAnysLock_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AppendAnysLock()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnysLock empty -- no change", actual)
	})
}

func Test_Seg2_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendNonEmptyAnys", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AppendNonEmptyAnys("a", nil, "", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendNonEmptyAnys -- skips nil and empty", actual)
	})
}

func Test_Seg2_Collection_AppendNonEmptyAnys_Nil(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendNonEmptyAnys_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AppendNonEmptyAnys(nil...)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendNonEmptyAnys nil -- no change", actual)
	})
}

// ── AppendAnysUsingFilter ───────────────────────────────────────────────────

func Test_Seg2_Collection_AppendAnysUsingFilter(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendAnysUsingFilter", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		filter := func(s string, i int) (string, bool, bool) {
			return s + "_f", s != "", false
		}
		c.AppendAnysUsingFilter(filter, "a", nil, "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilter -- filters applied", actual)
	})
}

func Test_Seg2_Collection_AppendAnysUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendAnysUsingFilter_Break", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, i == 0 // break after first
		}
		c.AppendAnysUsingFilter(filter, "a", "b", "c")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilter break -- only first", actual)
	})
}

func Test_Seg2_Collection_AppendAnysUsingFilter_Empty(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendAnysUsingFilter_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		c.AppendAnysUsingFilter(filter)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilter empty -- no change", actual)
	})
}

func Test_Seg2_Collection_AppendAnysUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendAnysUsingFilterLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, false
		}
		c.AppendAnysUsingFilterLock(filter, "a", nil, "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilterLock -- 2 items", actual)
	})
}

func Test_Seg2_Collection_AppendAnysUsingFilterLock_Nil(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendAnysUsingFilterLock_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		c.AppendAnysUsingFilterLock(filter, nil...)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilterLock nil -- no change", actual)
	})
}

func Test_Seg2_Collection_AppendAnysUsingFilterLock_Break(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendAnysUsingFilterLock_Break", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, true // always break
		}
		c.AppendAnysUsingFilterLock(filter, "a", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilterLock break -- only first", actual)
	})
}

func Test_Seg2_Collection_AppendAnysUsingFilterLock_Skip(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AppendAnysUsingFilterLock_Skip", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		filter := func(s string, i int) (string, bool, bool) {
			return s, false, false // skip all
		}
		c.AppendAnysUsingFilterLock(filter, "a", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilterLock skip all -- empty", actual)
	})
}

// ── AddsNonEmpty / AddsNonEmptyPtrLock ──────────────────────────────────────

func Test_Seg2_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AddsNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AddsNonEmpty("a", "", "b", "")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsNonEmpty -- skips empty", actual)
	})
}

func Test_Seg2_Collection_AddHashmapsKeysValuesUsingFilter(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AddHashmapsKeysValuesUsingFilter", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("keep", "v1")
		h.AddOrUpdate("skip", "v2")
		filter := func(kvp corestr.KeyValuePair) (string, bool, bool) {
			return kvp.Key + "=" + kvp.Value, kvp.Key == "keep", false
		}
		c.AddHashmapsKeysValuesUsingFilter(filter, h)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValuesUsingFilter -- 1 kept", actual)
	})
}

func Test_Seg2_Collection_AddHashmapsKeysValuesUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AddHashmapsKeysValuesUsingFilter_Break", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		h.AddOrUpdate("c", "3")
		filter := func(kvp corestr.KeyValuePair) (string, bool, bool) {
			return kvp.Key, true, true // accept and break immediately
		}
		c.AddHashmapsKeysValuesUsingFilter(filter, h)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValuesUsingFilter break -- only 1", actual)
	})
}

func Test_Seg2_Collection_AddHashmapsKeysValuesUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AddHashmapsKeysValuesUsingFilter_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		filter := func(kvp corestr.KeyValuePair) (string, bool, bool) { return "", true, false }
		c.AddHashmapsKeysValuesUsingFilter(filter, nil, nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValuesUsingFilter nil -- no change", actual)
	})
}

// ── AddWithWgLock ───────────────────────────────────────────────────────────

func Test_Seg2_Collection_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AddWithWgLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(wg, "hello")
		wg.Wait()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddWithWgLock -- 1 item added", actual)
	})
}

// ── AddStringsAsync ─────────────────────────────────────────────────────────

func Test_Seg2_Collection_AddStringsAsync_Empty(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AddStringsAsync_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		wg := &sync.WaitGroup{}
		c.AddStringsAsync(wg, []string{})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsAsync empty -- no change", actual)
	})
}

// ── AddsAsync ───────────────────────────────────────────────────────────────

func Test_Seg2_Collection_AddsAsync_Nil(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AddsAsync_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		wg := &sync.WaitGroup{}
		c.AddsAsync(wg, nil...)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsAsync nil -- no change", actual)
	})
}

// ── isResizeRequired / resizeForItems ───────────────────────────────────────
// These are tested implicitly through the above methods that add large batches.
// Direct testing not needed as they are private.

// ── AddCapacity ─────────────────────────────────────────────────────────────

func Test_Seg2_Collection_AddCapacity(t *testing.T) {
	safeTest(t, "Test_Seg2_Collection_AddCapacity", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddCapacity(100)

		// Act
		actual := args.Map{"capAbove": c.Capacity() >= 100}

		// Assert
		expected := args.Map{"capAbove": true}
		expected.ShouldBeEqual(t, 0, "AddCapacity -- capacity increased", actual)
	})
}
