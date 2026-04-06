package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Async / Concurrency methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_Collection_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_AddWithWgLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		wg := sync.WaitGroup{}

		// Act
		wg.Add(3)
		go c.AddWithWgLock(&wg, "a")
		go c.AddWithWgLock(&wg, "b")
		go c.AddWithWgLock(&wg, "c")
		wg.Wait()

		// Assert
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddWithWgLock", actual)
	})
}

func Test_Cov20_Collection_AddStringsAsync(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_AddStringsAsync", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		wg := sync.WaitGroup{}

		// Act
		c.AddStringsAsync(&wg, []string{"a", "b", "c"})
		wg.Wait()

		// Assert
		actual := args.Map{"hasItems": c.HasItems()}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddStringsAsync", actual)
	})
}

func Test_Cov20_Collection_AddStringsAsync_Empty(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_AddStringsAsync_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		wg := sync.WaitGroup{}

		// Act
		c.AddStringsAsync(&wg, []string{})

		// Assert
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddStringsAsync empty", actual)
	})
}

func Test_Cov20_Collection_AddsAsync(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_AddsAsync", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		wg := sync.WaitGroup{}
		wg.Add(1)

		// Act
		c.AddsAsync(&wg, "x", "y")
		wg.Wait()

		// Assert
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddsAsync", actual)
	})
}

func Test_Cov20_Collection_AddsAsync_Nil(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_AddsAsync_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		c.AddsAsync(nil)

		// Assert
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddsAsync nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Paging methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_Collection_GetPagedCollection(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_GetPagedCollection", func() {
		// Arrange
		items := make([]string, 10)
		for i := range items {
			items[i] = "x"
		}
		c := corestr.New.Collection.Strings(items)

		// Act
		paged := c.GetPagedCollection(3)

		// Assert
		actual := args.Map{"pages": paged.Length()}
		expected := args.Map{"pages": 4}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetPagedCollection", actual)
	})
}

func Test_Cov20_Collection_GetPagedCollection_Small(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_GetPagedCollection_Small", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		paged := c.GetPagedCollection(5)

		// Assert
		actual := args.Map{"pages": paged.Length()}
		expected := args.Map{"pages": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetPagedCollection small", actual)
	})
}

func Test_Cov20_Collection_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_GetSinglePageCollection", func() {
		// Arrange
		items := make([]string, 10)
		for i := range items {
			items[i] = "x"
		}
		c := corestr.New.Collection.Strings(items)

		// Act
		page := c.GetSinglePageCollection(3, 2)

		// Assert
		actual := args.Map{"len": page.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetSinglePageCollection", actual)
	})
}

func Test_Cov20_Collection_GetSinglePageCollection_LastPage(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_GetSinglePageCollection_LastPage", func() {
		// Arrange
		items := make([]string, 10)
		for i := range items {
			items[i] = "x"
		}
		c := corestr.New.Collection.Strings(items)

		// Act — last page has only 1 item
		page := c.GetSinglePageCollection(3, 4)

		// Assert
		actual := args.Map{"len": page.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetSinglePageCollection last", actual)
	})
}

func Test_Cov20_Collection_GetSinglePageCollection_Small(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_GetSinglePageCollection_Small", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		page := c.GetSinglePageCollection(5, 1)

		// Assert — returns self when length < eachPageSize
		actual := args.Map{"len": page.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetSinglePageCollection small", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Index/Remove methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_SafeIndexAtUsingLength", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		found := c.SafeIndexAtUsingLength("def", 2, 0)
		oob := c.SafeIndexAtUsingLength("def", 2, 5)

		// Assert
		actual := args.Map{"found": found, "oob": oob}
		expected := args.Map{"found": "a", "oob": "def"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- SafeIndexAtUsingLength", actual)
	})
}

func Test_Cov20_Collection_Single(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_Single", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"only"})

		// Act
		val := c.Single()

		// Assert
		actual := args.Map{"val": val}
		expected := args.Map{"val": "only"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Single", actual)
	})
}

func Test_Cov20_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_ChainRemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		c.ChainRemoveAt(1)

		// Assert
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ChainRemoveAt", actual)
	})
}

func Test_Cov20_Collection_RemoveItemsIndexes(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_RemoveItemsIndexes", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})

		// Act
		c.RemoveItemsIndexes(false, 1, 3)

		// Assert
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveItemsIndexes", actual)
	})
}

func Test_Cov20_Collection_RemoveItemsIndexes_NilIndexes(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_RemoveItemsIndexes_NilIndexes", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		c.RemoveItemsIndexesPtr(false, nil)

		// Assert
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- RemoveItemsIndexes nil", actual)
	})
}

func Test_Cov20_Collection_RemoveItemsIndexes_IgnoreError(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_RemoveItemsIndexes_IgnoreError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act — ignore error with nil indexes
		c.RemoveItemsIndexes(true)

		// Assert
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- RemoveItemsIndexes ignoreErr nil", actual)
	})
}

func Test_Cov20_Collection_AddHashmapsKeysValuesUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_AddHashmapsKeysValuesUsingFilter", func() {
		// Arrange
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		c := corestr.New.Collection.Cap(5)

		// Act
		c.AddHashmapsKeysValuesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Key + "=" + pair.Value, true, false
			},
			hm,
		)

		// Assert
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddHashmapsKeysValuesUsingFilter", actual)
	})
}

func Test_Cov20_Collection_AddHashmapsKeysValuesUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_AddHashmapsKeysValuesUsingFilter_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		c.AddHashmapsKeysValuesUsingFilter(nil, nil)

		// Assert
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddHashmapsKeysValuesUsingFilter nil", actual)
	})
}

func Test_Cov20_Collection_AddHashmapsKeysValuesUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_AddHashmapsKeysValuesUsingFilter_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.KeyValues(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)
		c := corestr.New.Collection.Cap(5)

		// Act
		c.AddHashmapsKeysValuesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Key, true, true // break after first
			},
			hm,
		)

		// Assert
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddHashmapsKeysValuesUsingFilter break", actual)
	})
}

func Test_Cov20_Collection_InsertAt(t *testing.T) {
	safeTest(t, "Test_Cov20_Collection_InsertAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "c"})

		// Act
		c.InsertAt(0, "z") // at last (index == length-1)

		// Assert
		actual := args.Map{"hasItems": c.HasItems()}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- InsertAt", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionsOfCollection — Async
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_CollectionsOfCollection_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_Cov20_CollectionsOfCollection_AddAsyncFuncItems", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(2)

		// Act
		coc.AddAsyncFuncItems(
			wg,
			false,
			func() []string { return []string{"a", "b"} },
			func() []string { return []string{"c"} },
		)

		// Assert
		actual := args.Map{"len": coc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection returns correct value -- AddAsyncFuncItems", actual)
	})
}

func Test_Cov20_CollectionsOfCollection_AddAsyncFuncItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov20_CollectionsOfCollection_AddAsyncFuncItems_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		coc.AddAsyncFuncItems(
			wg,
			false,
			func() []string { return []string{} },
		)

		// Assert
		actual := args.Map{"len": coc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection returns empty -- AddAsyncFuncItems empty", actual)
	})
}

func Test_Cov20_CollectionsOfCollection_AddAsyncFuncItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov20_CollectionsOfCollection_AddAsyncFuncItems_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		coc.AddAsyncFuncItems(nil, false)

		// Assert
		actual := args.Map{"len": coc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection returns nil -- AddAsyncFuncItems nil", actual)
	})
}

func Test_Cov20_CollectionsOfCollection_AddsStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Cov20_CollectionsOfCollection_AddsStringsOfStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		coc.AddsStringsOfStrings(false, []string{"a"}, []string{"b"})
		coc.AddsStringsOfStrings(false)

		// Assert
		actual := args.Map{"len": coc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection returns correct value -- AddsStringsOfStrings", actual)
	})
}

func Test_Cov20_CollectionsOfCollection_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov20_CollectionsOfCollection_AddCollections", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c := *corestr.New.Collection.Strings([]string{"a"})

		// Act
		coc.AddCollections(c)
		coc.Adds(c)

		// Assert
		actual := args.Map{"len": coc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection returns correct value -- AddCollections", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollections — Async methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_LinkedCollections_AddAsync(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddAsync", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		lc.AddAsync(corestr.New.Collection.Strings([]string{"a"}), wg)
		wg.Wait()

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddAsync", actual)
	})
}

func Test_Cov20_LinkedCollections_AddsAsyncOnComplete(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddsAsyncOnComplete", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		done := make(chan bool, 1)

		// Act
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.AddsAsyncOnComplete(
			func(lcs *corestr.LinkedCollections) {
				done <- true
			},
			false,
			c1, c2,
		)
		<-done

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddsAsyncOnComplete", actual)
	})
}

func Test_Cov20_LinkedCollections_AddsUsingProcessorAsyncOnComplete(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddsUsingProcessorAsyncOnComplete", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		done := make(chan bool, 1)

		// Act
		lc.AddsUsingProcessorAsyncOnComplete(
			func(lcs *corestr.LinkedCollections) {
				done <- true
			},
			func(a any, i int) *corestr.Collection {
				return corestr.New.Collection.Strings([]string{a.(string)})
			},
			false,
			"hello", "world",
		)
		<-done

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddsUsingProcessorAsyncOnComplete", actual)
	})
}

func Test_Cov20_LinkedCollections_AddsUsingProcessorAsyncOnComplete_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddsUsingProcessorAsyncOnComplete_NilSkip", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		done := make(chan bool, 1)

		// Act — nil items with isSkipOnNil true
		lc.AddsUsingProcessorAsyncOnComplete(
			func(lcs *corestr.LinkedCollections) {
				done <- true
			},
			func(a any, i int) *corestr.Collection {
				return corestr.New.Collection.Strings([]string{a.(string)})
			},
			true,
		)
		<-done

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns nil -- AddsUsingProcessorAsyncOnComplete nilSkip", actual)
	})
}

func Test_Cov20_LinkedCollections_AddsUsingProcessorAsync(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddsUsingProcessorAsync", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		lc.AddsUsingProcessorAsync(
			wg,
			func(a any, i int) *corestr.Collection {
				return corestr.New.Collection.Strings([]string{a.(string)})
			},
			false,
			"x", "y",
		)
		wg.Wait()

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddsUsingProcessorAsync", actual)
	})
}

func Test_Cov20_LinkedCollections_AddsUsingProcessorAsync_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddsUsingProcessorAsync_NilSkip", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		lc.AddsUsingProcessorAsync(wg, nil, true)
		wg.Wait()

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns nil -- AddsUsingProcessorAsync nilSkip", actual)
	})
}

func Test_Cov20_LinkedCollections_AppendChainOfNodesAsync(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AppendChainOfNodesAsync", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		node := &corestr.LinkedCollectionNode{
			Element: corestr.New.Collection.Strings([]string{"a"}),
		}

		// Act
		lc.AppendChainOfNodesAsync(node, wg)
		wg.Wait()

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AppendChainOfNodesAsync", actual)
	})
}

func Test_Cov20_LinkedCollections_AddCollectionsToNodeAsync(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddCollectionsToNodeAsync", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		node := lc.Head()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		lc.AddCollectionsToNodeAsync(
			false,
			wg,
			node,
			corestr.New.Collection.Strings([]string{"inserted"}),
		)
		wg.Wait()

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddCollectionsToNodeAsync", actual)
	})
}

func Test_Cov20_LinkedCollections_AddCollectionsToNodeAsync_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddCollectionsToNodeAsync_NilSkip", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act — nil collections with isSkipOnNull true
		lc.AddCollectionsToNodeAsync(true, nil, nil)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns nil -- AddCollectionsToNodeAsync nil", actual)
	})
}

func Test_Cov20_LinkedCollections_AddAfterNodeAsync(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddAfterNodeAsync", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		node := lc.Head()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		lc.AddAfterNodeAsync(wg, node, corestr.New.Collection.Strings([]string{"b"}))
		wg.Wait()

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddAfterNodeAsync", actual)
	})
}

func Test_Cov20_LinkedCollections_AddStringsAsync(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddStringsAsync", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		lc.AddStringsAsync(wg, []string{"a", "b"})
		wg.Wait()

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddStringsAsync", actual)
	})
}

func Test_Cov20_LinkedCollections_AddStringsAsync_Nil(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddStringsAsync_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		lc.AddStringsAsync(nil, nil)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns nil -- AddStringsAsync nil", actual)
	})
}

func Test_Cov20_LinkedCollections_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddAsyncFuncItems", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(2)

		// Act
		lc.AddAsyncFuncItems(
			wg,
			false,
			func() []string { return []string{"a"} },
			func() []string { return []string{"b"} },
		)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddAsyncFuncItems", actual)
	})
}

func Test_Cov20_LinkedCollections_AddAsyncFuncItems_EmptyReturn(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddAsyncFuncItems_EmptyReturn", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		lc.AddAsyncFuncItems(wg, false, func() []string { return []string{} })

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns empty -- AddAsyncFuncItems empty", actual)
	})
}

func Test_Cov20_LinkedCollections_AddAsyncFuncItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddAsyncFuncItems_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		lc.AddAsyncFuncItems(nil, false)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns nil -- AddAsyncFuncItems nil", actual)
	})
}

func Test_Cov20_LinkedCollections_AddAsyncFuncItemsPointer(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddAsyncFuncItemsPointer", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(2)

		// Act
		lc.AddAsyncFuncItemsPointer(
			wg,
			true,
			func() []string { return []string{"x"} },
			func() []string { return []string{"y"} },
		)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddAsyncFuncItemsPointer", actual)
	})
}

func Test_Cov20_LinkedCollections_AddAsyncFuncItemsPointer_Nil(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddAsyncFuncItemsPointer_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		lc.AddAsyncFuncItemsPointer(nil, false)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns nil -- AddAsyncFuncItemsPointer nil", actual)
	})
}

func Test_Cov20_LinkedCollections_AddAsyncFuncItemsPointer_EmptyReturn(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddAsyncFuncItemsPointer_EmptyReturn", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		lc.AddAsyncFuncItemsPointer(wg, false, func() []string { return nil })

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns empty -- AddAsyncFuncItemsPointer empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollections — Non-async uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_LinkedCollections_AttachWithNode(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AttachWithNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c1)
		current := lc.Head()

		adding := &corestr.LinkedCollectionNode{
			Element: corestr.New.Collection.Strings([]string{"b"}),
		}

		// Act
		err := lc.AttachWithNode(current, adding)

		// Assert
		actual := args.Map{"err": err == nil, "len": lc.Length()}
		expected := args.Map{"err": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns non-empty -- AttachWithNode", actual)
	})
}

func Test_Cov20_LinkedCollections_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AttachWithNode_NilCurrent", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		adding := &corestr.LinkedCollectionNode{
			Element: corestr.New.Collection.Strings([]string{"b"}),
		}

		// Act
		err := lc.AttachWithNode(nil, adding)

		// Assert
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns nil -- AttachWithNode nil current", actual)
	})
}

func Test_Cov20_LinkedCollections_FilterAsCollection(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_FilterAsCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		result := lc.FilterAsCollection(
			func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
				return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
			},
			0,
		)

		// Assert
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- FilterAsCollection", actual)
	})
}

func Test_Cov20_LinkedCollections_FilterAsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_FilterAsCollection_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		result := lc.FilterAsCollection(
			func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
				return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
			},
			0,
		)

		// Assert
		actual := args.Map{"empty": result.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns empty -- FilterAsCollection empty", actual)
	})
}

func Test_Cov20_LinkedCollections_FilterAsCollections(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_FilterAsCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))

		// Act
		result := lc.FilterAsCollections(
			func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
				return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
			},
		)

		// Assert
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- FilterAsCollections", actual)
	})
}

func Test_Cov20_LinkedCollections_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_RemoveNodeByIndex", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act — remove middle
		lc.RemoveNodeByIndex(1)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- RemoveNodeByIndex middle", actual)
	})
}

func Test_Cov20_LinkedCollections_RemoveNodeByIndex_Last(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_RemoveNodeByIndex_Last", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))

		// Act — remove last
		lc.RemoveNodeByIndex(1)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- RemoveNodeByIndex last", actual)
	})
}

func Test_Cov20_LinkedCollections_RemoveNode(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_RemoveNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		node := lc.Head()

		// Act — remove head by reference
		lc.RemoveNode(node)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- RemoveNode head", actual)
	})
}

func Test_Cov20_LinkedCollections_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_RemoveNodeByIndexes", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		lc.RemoveNodeByIndexes(true, 0, 2)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- RemoveNodeByIndexes", actual)
	})
}

func Test_Cov20_LinkedCollections_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_ConcatNew", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))

		// Act
		result := lc1.ConcatNew(false, lc2)

		// Assert
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- ConcatNew", actual)
	})
}

func Test_Cov20_LinkedCollections_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_ConcatNew_EmptyClone", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		result := lc.ConcatNew(true)

		// Assert
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns empty -- ConcatNew emptyClone", actual)
	})
}

func Test_Cov20_LinkedCollections_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_ConcatNew_EmptyNoClone", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		result := lc.ConcatNew(false)

		// Assert — returns self
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns empty -- ConcatNew emptyNoClone", actual)
	})
}

func Test_Cov20_LinkedCollections_AppendCollections(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AppendCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		lc.AppendCollections(false,
			corestr.New.Collection.Strings([]string{"a"}),
			corestr.New.Collection.Strings([]string{"b"}),
		)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AppendCollections", actual)
	})
}

func Test_Cov20_LinkedCollections_AppendCollectionsPointersLock(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AppendCollectionsPointersLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		cols := []*corestr.Collection{
			corestr.New.Collection.Strings([]string{"a"}),
			corestr.New.Collection.Strings([]string{"b"}),
		}

		// Act
		lc.AppendCollectionsPointersLock(false, &cols)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AppendCollectionsPointersLock", actual)
	})
}

func Test_Cov20_LinkedCollections_AddStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddStringsOfStrings", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		lc.AddStringsOfStrings(false, []string{"a"}, []string{"b"})
		lc.AddStringsOfStrings(false)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddStringsOfStrings", actual)
	})
}

func Test_Cov20_LinkedCollections_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		lc.AddCollections([]*corestr.Collection{
			corestr.New.Collection.Strings([]string{"a"}),
		})
		lc.AddCollections(nil)
		lc.AddCollectionsPtr([]*corestr.Collection{
			corestr.New.Collection.Strings([]string{"b"}),
		})
		lc.AddCollectionsPtr(nil)

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddCollections", actual)
	})
}

func Test_Cov20_LinkedCollections_InsertAt(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_InsertAt", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		lc.InsertAt(1, corestr.New.Collection.Strings([]string{"b"}))

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- InsertAt", actual)
	})
}

func Test_Cov20_LinkedCollections_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_InsertAt_Front", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))

		// Act
		lc.InsertAt(0, corestr.New.Collection.Strings([]string{"a"}))

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- InsertAt front", actual)
	})
}

func Test_Cov20_LinkedCollections_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_GetNextNodes", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		nodes := lc.GetNextNodes(2)
		all := lc.GetAllLinkedNodes()

		// Assert
		actual := args.Map{"next": len(nodes), "all": len(all)}
		expected := args.Map{"next": 2, "all": 3}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- GetNextNodes", actual)
	})
}

func Test_Cov20_LinkedCollections_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_SafeIndexAt", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))

		// Act
		node := lc.SafeIndexAt(1)
		nilNode := lc.SafeIndexAt(99)
		negNode := lc.SafeIndexAt(-1)
		ptr := lc.SafePointerIndexAt(0)
		nilPtr := lc.SafePointerIndexAt(-1)

		// Assert
		actual := args.Map{
			"found":   node != nil,
			"nilNode": nilNode == nil,
			"negNode": negNode == nil,
			"ptr":     ptr != nil,
			"nilPtr":  nilPtr == nil,
		}
		expected := args.Map{
			"found": true, "nilNode": true, "negNode": true,
			"ptr": true, "nilPtr": true,
		}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- SafeIndexAt", actual)
	})
}

func Test_Cov20_LinkedCollections_JsonRoundTrip(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_JsonRoundTrip", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))

		// Act
		_ = lc.String()
		_ = lc.StringLock()
		_ = lc.Join(",")
		_ = lc.Joins(",", "extra")
		_ = lc.JsonModel()
		_ = lc.JsonModelAny()
		_, _ = lc.MarshalJSON()
		_ = lc.AsJsoner()
		_ = lc.AsJsonContractsBinder()
		_ = lc.AsJsonParseSelfInjector()
		_ = lc.AsJsonMarshaller()
		_ = lc.GetCompareSummary(lc, "left", "right")

		// Assert
		actual := args.Map{"done": true}
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- JsonRoundTrip", actual)
	})
}

func Test_Cov20_LinkedCollections_Clear(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_Clear", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		lc.Clear()
		lc.RemoveAll()

		// Assert
		actual := args.Map{"empty": lc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- Clear", actual)
	})
}

func Test_Cov20_LinkedCollections_AppendNode(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AppendNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{
			Element: corestr.New.Collection.Strings([]string{"a"}),
		}

		// Act
		lc.AppendNode(node)
		lc.AddBackNode(&corestr.LinkedCollectionNode{
			Element: corestr.New.Collection.Strings([]string{"b"}),
		})

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AppendNode", actual)
	})
}

func Test_Cov20_LinkedCollections_AddAnother(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddAnother", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		lc1.AddAnother(lc2)
		lc1.AddAnother(nil)

		// Assert
		actual := args.Map{"len": lc1.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddAnother", actual)
	})
}

func Test_Cov20_LinkedCollections_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_AddCollectionToNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		node := lc.Head()

		// Act
		lc.AddCollectionToNode(false, node, corestr.New.Collection.Strings([]string{"b"}))

		// Assert
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- AddCollectionToNode", actual)
	})
}
func Test_Cov20_LinkedCollections_IsEqualsPtr(t *testing.T) {
	safeTest(t, "Test_Cov20_LinkedCollections_IsEqualsPtr", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Strings("a")
		lc2 := corestr.New.LinkedCollection.Strings("a")
		empty1 := corestr.New.LinkedCollection.Create()
		empty2 := corestr.New.LinkedCollection.Create()

		// Act & Assert
		actual := args.Map{
			"equal":      lc1.IsEqualsPtr(lc2),
			"self":       lc1.IsEqualsPtr(lc1),
			"nil":        lc1.IsEqualsPtr(nil),
			"bothEmpty":  empty1.IsEqualsPtr(empty2),
			"oneEmpty":   lc1.IsEqualsPtr(empty1),
		}
		expected := args.Map{
			"equal": true, "self": true, "nil": false,
			"bothEmpty": true, "oneEmpty": false,
		}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- IsEqualsPtr", actual)
	})
}
