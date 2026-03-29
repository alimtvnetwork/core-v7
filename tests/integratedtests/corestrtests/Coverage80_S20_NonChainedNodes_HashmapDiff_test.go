package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// =============================================
// S20: NonChainedLinkedListNodes
// =============================================

func Test_S20_001_NewNonChainedLinkedListNodes_creates_with_capacity(t *testing.T) {
	safeTest(t, "Test_S20_001_NewNonChainedLinkedListNodes_creates_with_capacity", func() {
		// Arrange
		capacity := 5

		// Act
		nodes := corestr.NewNonChainedLinkedListNodes(capacity)

		// Assert
		if nodes == nil {
			t.Fatal("NewNonChainedLinkedListNodes returns nil -- capacity 5")
		}
		if nodes.Length() != 0 {
			t.Errorf("Length returns 0 -- empty after creation, got %d", nodes.Length())
		}
		if !nodes.IsEmpty() {
			t.Error("IsEmpty returns true -- no items added")
		}
		if nodes.HasItems() {
			t.Error("HasItems returns false -- no items added")
		}
	})
}

func Test_S20_002_NonChainedLinkedListNodes_Adds_single(t *testing.T) {
	safeTest(t, "Test_S20_002_NonChainedLinkedListNodes_Adds_single", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		head := ll.Head()

		// Act
		nodes.Adds(head)

		// Assert
		if nodes.Length() != 1 {
			t.Errorf("Length returns 1 -- one node added, got %d", nodes.Length())
		}
		if nodes.IsEmpty() {
			t.Error("IsEmpty returns false -- has items")
		}
		if !nodes.HasItems() {
			t.Error("HasItems returns true -- has items")
		}
	})
}

func Test_S20_003_NonChainedLinkedListNodes_Adds_nil(t *testing.T) {
	safeTest(t, "Test_S20_003_NonChainedLinkedListNodes_Adds_nil", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		result := nodes.Adds(nil...)

		// Assert
		if result != nodes {
			t.Error("Adds returns self -- nil input")
		}
		if nodes.Length() != 0 {
			t.Errorf("Length returns 0 -- nil not added, got %d", nodes.Length())
		}
	})
}

func Test_S20_004_NonChainedLinkedListNodes_First_Last(t *testing.T) {
	safeTest(t, "Test_S20_004_NonChainedLinkedListNodes_First_Last", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)
		ll := corestr.New.LinkedList.SpreadStrings("alpha", "beta", "gamma")
		head := ll.Head()
		node2 := head.Next()
		node3 := node2.Next()

		// Act
		nodes.Adds(head, node2, node3)

		// Assert
		if nodes.First().Element != "alpha" {
			t.Errorf("First returns alpha -- first added, got %s", nodes.First().Element)
		}
		if nodes.Last().Element != "gamma" {
			t.Errorf("Last returns gamma -- last added, got %s", nodes.Last().Element)
		}
	})
}

func Test_S20_005_NonChainedLinkedListNodes_FirstOrDefault_empty(t *testing.T) {
	safeTest(t, "Test_S20_005_NonChainedLinkedListNodes_FirstOrDefault_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		result := nodes.FirstOrDefault()

		// Assert
		if result != nil {
			t.Error("FirstOrDefault returns nil -- empty nodes")
		}
	})
}

func Test_S20_006_NonChainedLinkedListNodes_LastOrDefault_empty(t *testing.T) {
	safeTest(t, "Test_S20_006_NonChainedLinkedListNodes_LastOrDefault_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		result := nodes.LastOrDefault()

		// Assert
		if result != nil {
			t.Error("LastOrDefault returns nil -- empty nodes")
		}
	})
}

func Test_S20_007_NonChainedLinkedListNodes_FirstOrDefault_has_items(t *testing.T) {
	safeTest(t, "Test_S20_007_NonChainedLinkedListNodes_FirstOrDefault_has_items", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)
		ll := corestr.New.LinkedList.SpreadStrings("x")
		nodes.Adds(ll.Head())

		// Act
		result := nodes.FirstOrDefault()

		// Assert
		if result == nil || result.Element != "x" {
			t.Error("FirstOrDefault returns x -- single item")
		}
	})
}

func Test_S20_008_NonChainedLinkedListNodes_LastOrDefault_has_items(t *testing.T) {
	safeTest(t, "Test_S20_008_NonChainedLinkedListNodes_LastOrDefault_has_items", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)
		ll := corestr.New.LinkedList.SpreadStrings("x", "y")
		nodes.Adds(ll.Head(), ll.Head().Next())

		// Act
		result := nodes.LastOrDefault()

		// Assert
		if result == nil || result.Element != "y" {
			t.Error("LastOrDefault returns y -- two items")
		}
	})
}

func Test_S20_009_NonChainedLinkedListNodes_IsChainingApplied_false(t *testing.T) {
	safeTest(t, "Test_S20_009_NonChainedLinkedListNodes_IsChainingApplied_false", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		result := nodes.IsChainingApplied()

		// Assert
		if result {
			t.Error("IsChainingApplied returns false -- not applied yet")
		}
	})
}

func Test_S20_010_NonChainedLinkedListNodes_ApplyChaining_empty(t *testing.T) {
	safeTest(t, "Test_S20_010_NonChainedLinkedListNodes_ApplyChaining_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		result := nodes.ApplyChaining()

		// Assert
		if result != nodes {
			t.Error("ApplyChaining returns self -- empty nodes")
		}
		if nodes.IsChainingApplied() {
			t.Error("IsChainingApplied returns false -- empty, no chaining applied")
		}
	})
}

func Test_S20_011_NonChainedLinkedListNodes_ApplyChaining_multi(t *testing.T) {
	safeTest(t, "Test_S20_011_NonChainedLinkedListNodes_ApplyChaining_multi", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)
		// Create individual nodes (not chained)
		ll1 := corestr.New.LinkedList.SpreadStrings("a")
		ll2 := corestr.New.LinkedList.SpreadStrings("b")
		ll3 := corestr.New.LinkedList.SpreadStrings("c")
		nodes.Adds(ll1.Head(), ll2.Head(), ll3.Head())

		// Act
		result := nodes.ApplyChaining()

		// Assert
		if result != nodes {
			t.Error("ApplyChaining returns self")
		}
		if !nodes.IsChainingApplied() {
			t.Error("IsChainingApplied returns true -- chaining applied")
		}
		first := nodes.First()
		if first.Next() == nil || first.Next().Element != "b" {
			t.Error("ApplyChaining chains a->b -- correct next")
		}
		if nodes.Last().HasNext() {
			t.Error("Last node has nil next -- end of chain")
		}
	})
}

func Test_S20_012_NonChainedLinkedListNodes_Items(t *testing.T) {
	safeTest(t, "Test_S20_012_NonChainedLinkedListNodes_Items", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)
		ll := corestr.New.LinkedList.SpreadStrings("x")
		nodes.Adds(ll.Head())

		// Act
		items := nodes.Items()

		// Assert
		if len(items) != 1 {
			t.Errorf("Items returns 1 item -- one added, got %d", len(items))
		}
	})
}

func Test_S20_013_NonChainedLinkedListNodes_ToChainedNodes(t *testing.T) {
	safeTest(t, "Test_S20_013_NonChainedLinkedListNodes_ToChainedNodes", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		nodes.Adds(ll.Head(), ll.Head().Next())

		// Act
		chained := nodes.ToChainedNodes()

		// Assert
		if chained == nil {
			t.Error("ToChainedNodes returns non-nil -- has items")
		}
	})
}

func Test_S20_014_NonChainedLinkedListNodes_ToChainedNodes_empty(t *testing.T) {
	safeTest(t, "Test_S20_014_NonChainedLinkedListNodes_ToChainedNodes_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		chained := nodes.ToChainedNodes()

		// Assert
		if chained == nil {
			t.Error("ToChainedNodes returns non-nil slice -- empty input")
		}
		if len(chained) != 0 {
			t.Errorf("ToChainedNodes returns empty slice -- empty input, got %d", len(chained))
		}
	})
}

// =============================================
// S20: NonChainedLinkedCollectionNodes
// =============================================

func Test_S20_020_NewNonChainedLinkedCollectionNodes_creates(t *testing.T) {
	safeTest(t, "Test_S20_020_NewNonChainedLinkedCollectionNodes_creates", func() {
		// Arrange & Act
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Assert
		if nodes == nil {
			t.Fatal("NewNonChainedLinkedCollectionNodes returns non-nil -- capacity 5")
		}
		if nodes.Length() != 0 {
			t.Errorf("Length returns 0 -- empty, got %d", nodes.Length())
		}
		if !nodes.IsEmpty() {
			t.Error("IsEmpty returns true -- no items")
		}
	})
}

func Test_S20_021_NonChainedLinkedCollectionNodes_Adds(t *testing.T) {
	safeTest(t, "Test_S20_021_NonChainedLinkedCollectionNodes_Adds", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		lc := corestr.New.LinkedCollection.UsingCollections(col)
		head := lc.Head()

		// Act
		nodes.Adds(head)

		// Assert
		if nodes.Length() != 1 {
			t.Errorf("Length returns 1 -- one node, got %d", nodes.Length())
		}
		if !nodes.HasItems() {
			t.Error("HasItems returns true -- has items")
		}
	})
}

func Test_S20_022_NonChainedLinkedCollectionNodes_Adds_nil(t *testing.T) {
	safeTest(t, "Test_S20_022_NonChainedLinkedCollectionNodes_Adds_nil", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		result := nodes.Adds(nil...)

		// Assert
		if result != nodes {
			t.Error("Adds returns self -- nil input")
		}
	})
}

func Test_S20_023_NonChainedLinkedCollectionNodes_FirstOrDefault_empty(t *testing.T) {
	safeTest(t, "Test_S20_023_NonChainedLinkedCollectionNodes_FirstOrDefault_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		result := nodes.FirstOrDefault()

		// Assert
		if result != nil {
			t.Error("FirstOrDefault returns nil -- empty")
		}
	})
}

func Test_S20_024_NonChainedLinkedCollectionNodes_LastOrDefault_empty(t *testing.T) {
	safeTest(t, "Test_S20_024_NonChainedLinkedCollectionNodes_LastOrDefault_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		result := nodes.LastOrDefault()

		// Assert
		if result != nil {
			t.Error("LastOrDefault returns nil -- empty")
		}
	})
}

func Test_S20_025_NonChainedLinkedCollectionNodes_First_Last(t *testing.T) {
	safeTest(t, "Test_S20_025_NonChainedLinkedCollectionNodes_First_Last", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)
		col1 := corestr.New.Collection.Strings([]string{"a"})
		col2 := corestr.New.Collection.Strings([]string{"b"})
		lc1 := corestr.New.LinkedCollection.UsingCollections(col1)
		lc2 := corestr.New.LinkedCollection.UsingCollections(col2)

		// Act
		nodes.Adds(lc1.Head(), lc2.Head())

		// Assert
		if nodes.First().Element.List()[0] != "a" {
			t.Error("First returns collection with a")
		}
		if nodes.Last().Element.List()[0] != "b" {
			t.Error("Last returns collection with b")
		}
	})
}

func Test_S20_026_NonChainedLinkedCollectionNodes_FirstOrDefault_has_items(t *testing.T) {
	safeTest(t, "Test_S20_026_NonChainedLinkedCollectionNodes_FirstOrDefault_has_items", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)
		col := corestr.New.Collection.Strings([]string{"x"})
		lc := corestr.New.LinkedCollection.UsingCollections(col)
		nodes.Adds(lc.Head())

		// Act
		result := nodes.FirstOrDefault()

		// Assert
		if result == nil {
			t.Error("FirstOrDefault returns non-nil -- has items")
		}
	})
}

func Test_S20_027_NonChainedLinkedCollectionNodes_LastOrDefault_has_items(t *testing.T) {
	safeTest(t, "Test_S20_027_NonChainedLinkedCollectionNodes_LastOrDefault_has_items", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)
		col := corestr.New.Collection.Strings([]string{"x"})
		lc := corestr.New.LinkedCollection.UsingCollections(col)
		nodes.Adds(lc.Head())

		// Act
		result := nodes.LastOrDefault()

		// Assert
		if result == nil {
			t.Error("LastOrDefault returns non-nil -- has items")
		}
	})
}

func Test_S20_028_NonChainedLinkedCollectionNodes_IsChainingApplied_false(t *testing.T) {
	safeTest(t, "Test_S20_028_NonChainedLinkedCollectionNodes_IsChainingApplied_false", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act & Assert
		if nodes.IsChainingApplied() {
			t.Error("IsChainingApplied returns false -- not applied")
		}
	})
}

func Test_S20_029_NonChainedLinkedCollectionNodes_ApplyChaining_empty(t *testing.T) {
	safeTest(t, "Test_S20_029_NonChainedLinkedCollectionNodes_ApplyChaining_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		result := nodes.ApplyChaining()

		// Assert
		if result != nodes {
			t.Error("ApplyChaining returns self -- empty")
		}
		if nodes.IsChainingApplied() {
			t.Error("IsChainingApplied returns false -- empty, no chaining")
		}
	})
}

func Test_S20_030_NonChainedLinkedCollectionNodes_ApplyChaining_multi(t *testing.T) {
	safeTest(t, "Test_S20_030_NonChainedLinkedCollectionNodes_ApplyChaining_multi", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)
		col1 := corestr.New.Collection.Strings([]string{"a"})
		col2 := corestr.New.Collection.Strings([]string{"b"})
		col3 := corestr.New.Collection.Strings([]string{"c"})
		lc1 := corestr.New.LinkedCollection.UsingCollections(col1)
		lc2 := corestr.New.LinkedCollection.UsingCollections(col2)
		lc3 := corestr.New.LinkedCollection.UsingCollections(col3)
		nodes.Adds(lc1.Head(), lc2.Head(), lc3.Head())

		// Act
		result := nodes.ApplyChaining()

		// Assert
		if result != nodes {
			t.Error("ApplyChaining returns self")
		}
		if !nodes.IsChainingApplied() {
			t.Error("IsChainingApplied returns true -- applied")
		}
		if nodes.Last().HasNext() {
			t.Error("Last node has nil next -- end of chain")
		}
	})
}

func Test_S20_031_NonChainedLinkedCollectionNodes_Items(t *testing.T) {
	safeTest(t, "Test_S20_031_NonChainedLinkedCollectionNodes_Items", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		items := nodes.Items()

		// Assert
		if items == nil {
			t.Error("Items returns non-nil -- empty but initialized")
		}
	})
}

func Test_S20_032_NonChainedLinkedCollectionNodes_ToChainedNodes_empty(t *testing.T) {
	safeTest(t, "Test_S20_032_NonChainedLinkedCollectionNodes_ToChainedNodes_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		chained := nodes.ToChainedNodes()

		// Assert
		if chained == nil {
			t.Error("ToChainedNodes returns non-nil -- empty input")
		}
	})
}

func Test_S20_033_NonChainedLinkedCollectionNodes_ToChainedNodes_multi(t *testing.T) {
	safeTest(t, "Test_S20_033_NonChainedLinkedCollectionNodes_ToChainedNodes_multi", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)
		col1 := corestr.New.Collection.Strings([]string{"a"})
		col2 := corestr.New.Collection.Strings([]string{"b"})
		lc1 := corestr.New.LinkedCollection.UsingCollections(col1)
		lc2 := corestr.New.LinkedCollection.UsingCollections(col2)
		nodes.Adds(lc1.Head(), lc2.Head())

		// Act
		chained := nodes.ToChainedNodes()

		// Assert
		if chained == nil {
			t.Error("ToChainedNodes returns non-nil -- has items")
		}
	})
}

// =============================================
// S20: HashmapDiff
// =============================================

func Test_S20_040_HashmapDiff_Length_nil(t *testing.T) {
	safeTest(t, "Test_S20_040_HashmapDiff_Length_nil", func() {
		// Arrange
		var hd *corestr.HashmapDiff

		// Act
		result := hd.Length()

		// Assert
		if result != 0 {
			t.Errorf("Length returns 0 -- nil receiver, got %d", result)
		}
	})
}

func Test_S20_041_HashmapDiff_Length_with_items(t *testing.T) {
	safeTest(t, "Test_S20_041_HashmapDiff_Length_with_items", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2"})

		// Act
		result := hd.Length()

		// Assert
		if result != 2 {
			t.Errorf("Length returns 2 -- two items, got %d", result)
		}
	})
}

func Test_S20_042_HashmapDiff_IsEmpty_true(t *testing.T) {
	safeTest(t, "Test_S20_042_HashmapDiff_IsEmpty_true", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{})

		// Act & Assert
		if !hd.IsEmpty() {
			t.Error("IsEmpty returns true -- empty map")
		}
	})
}

func Test_S20_043_HashmapDiff_IsEmpty_false(t *testing.T) {
	safeTest(t, "Test_S20_043_HashmapDiff_IsEmpty_false", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act & Assert
		if hd.IsEmpty() {
			t.Error("IsEmpty returns false -- has item")
		}
	})
}

func Test_S20_044_HashmapDiff_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_S20_044_HashmapDiff_HasAnyItem", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act & Assert
		if !hd.HasAnyItem() {
			t.Error("HasAnyItem returns true -- has item")
		}
	})
}

func Test_S20_045_HashmapDiff_LastIndex(t *testing.T) {
	safeTest(t, "Test_S20_045_HashmapDiff_LastIndex", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2", "c": "3"})

		// Act
		result := hd.LastIndex()

		// Assert
		if result != 2 {
			t.Errorf("LastIndex returns 2 -- 3 items, got %d", result)
		}
	})
}

func Test_S20_046_HashmapDiff_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_S20_046_HashmapDiff_AllKeysSorted", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"c": "3", "a": "1", "b": "2"})

		// Act
		keys := hd.AllKeysSorted()

		// Assert
		if len(keys) != 3 {
			t.Fatalf("AllKeysSorted returns 3 keys, got %d", len(keys))
		}
		if keys[0] != "a" || keys[1] != "b" || keys[2] != "c" {
			t.Errorf("AllKeysSorted returns sorted keys -- a,b,c, got %v", keys)
		}
	})
}

func Test_S20_047_HashmapDiff_MapAnyItems(t *testing.T) {
	safeTest(t, "Test_S20_047_HashmapDiff_MapAnyItems", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"x": "10"})

		// Act
		result := hd.MapAnyItems()

		// Assert
		if len(result) != 1 {
			t.Errorf("MapAnyItems returns 1 item, got %d", len(result))
		}
		if result["x"] != "10" {
			t.Errorf("MapAnyItems has x=10, got %v", result["x"])
		}
	})
}

func Test_S20_048_HashmapDiff_MapAnyItems_nil(t *testing.T) {
	safeTest(t, "Test_S20_048_HashmapDiff_MapAnyItems_nil", func() {
		// Arrange
		var hd *corestr.HashmapDiff

		// Act
		result := hd.MapAnyItems()

		// Assert
		if result == nil || len(result) != 0 {
			t.Error("MapAnyItems returns empty map -- nil receiver")
		}
	})
}

func Test_S20_049_HashmapDiff_Raw_nil(t *testing.T) {
	safeTest(t, "Test_S20_049_HashmapDiff_Raw_nil", func() {
		// Arrange
		var hd *corestr.HashmapDiff

		// Act
		result := hd.Raw()

		// Assert
		if result == nil || len(result) != 0 {
			t.Error("Raw returns empty map -- nil receiver")
		}
	})
}

func Test_S20_050_HashmapDiff_Raw_with_items(t *testing.T) {
	safeTest(t, "Test_S20_050_HashmapDiff_Raw_with_items", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		result := hd.Raw()

		// Assert
		if len(result) != 1 || result["a"] != "1" {
			t.Error("Raw returns underlying map -- has items")
		}
	})
}

func Test_S20_051_HashmapDiff_IsRawEqual_true(t *testing.T) {
	safeTest(t, "Test_S20_051_HashmapDiff_IsRawEqual_true", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		result := hd.IsRawEqual(map[string]string{"a": "1"})

		// Assert
		if !result {
			t.Error("IsRawEqual returns true -- same maps")
		}
	})
}

func Test_S20_052_HashmapDiff_IsRawEqual_false(t *testing.T) {
	safeTest(t, "Test_S20_052_HashmapDiff_IsRawEqual_false", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		result := hd.IsRawEqual(map[string]string{"a": "2"})

		// Assert
		if result {
			t.Error("IsRawEqual returns false -- different values")
		}
	})
}

func Test_S20_053_HashmapDiff_HasAnyChanges(t *testing.T) {
	safeTest(t, "Test_S20_053_HashmapDiff_HasAnyChanges", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		result := hd.HasAnyChanges(map[string]string{"a": "2"})

		// Assert
		if !result {
			t.Error("HasAnyChanges returns true -- different values")
		}
	})
}

func Test_S20_054_HashmapDiff_HasAnyChanges_no_changes(t *testing.T) {
	safeTest(t, "Test_S20_054_HashmapDiff_HasAnyChanges_no_changes", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		result := hd.HasAnyChanges(map[string]string{"a": "1"})

		// Assert
		if result {
			t.Error("HasAnyChanges returns false -- same values")
		}
	})
}

func Test_S20_055_HashmapDiff_DiffRaw(t *testing.T) {
	safeTest(t, "Test_S20_055_HashmapDiff_DiffRaw", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2"})

		// Act
		diff := hd.DiffRaw(map[string]string{"a": "1", "b": "3"})

		// Assert
		if diff == nil {
			t.Fatal("DiffRaw returns non-nil -- has diff")
		}
		if diff["b"] != "2" {
			t.Errorf("DiffRaw contains b=2 -- left value for changed key, got %s", diff["b"])
		}
	})
}

func Test_S20_056_HashmapDiff_HashmapDiffUsingRaw_no_diff(t *testing.T) {
	safeTest(t, "Test_S20_056_HashmapDiff_HashmapDiffUsingRaw_no_diff", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		result := hd.HashmapDiffUsingRaw(map[string]string{"a": "1"})

		// Assert
		if len(result) != 0 {
			t.Error("HashmapDiffUsingRaw returns empty -- no diff")
		}
	})
}

func Test_S20_057_HashmapDiff_DiffJsonMessage(t *testing.T) {
	safeTest(t, "Test_S20_057_HashmapDiff_DiffJsonMessage", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		msg := hd.DiffJsonMessage(map[string]string{"a": "2"})

		// Assert
		if msg == "" {
			t.Error("DiffJsonMessage returns non-empty -- has diff")
		}
	})
}

func Test_S20_058_HashmapDiff_ToStringsSliceOfDiffMap(t *testing.T) {
	safeTest(t, "Test_S20_058_HashmapDiff_ToStringsSliceOfDiffMap", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		diffMap := map[string]string{"a": "2"}

		// Act
		result := hd.ToStringsSliceOfDiffMap(diffMap)

		// Assert
		if len(result) == 0 {
			t.Error("ToStringsSliceOfDiffMap returns non-empty -- has items")
		}
	})
}

func Test_S20_059_HashmapDiff_ShouldDiffMessage(t *testing.T) {
	safeTest(t, "Test_S20_059_HashmapDiff_ShouldDiffMessage", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		msg := hd.ShouldDiffMessage("test", map[string]string{"a": "2"})

		// Assert
		if msg == "" {
			t.Error("ShouldDiffMessage returns non-empty -- has diff")
		}
	})
}

func Test_S20_060_HashmapDiff_LogShouldDiffMessage(t *testing.T) {
	safeTest(t, "Test_S20_060_HashmapDiff_LogShouldDiffMessage", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		msg := hd.LogShouldDiffMessage("test", map[string]string{"a": "2"})

		// Assert
		if msg == "" {
			t.Error("LogShouldDiffMessage returns non-empty -- has diff")
		}
	})
}

func Test_S20_061_HashmapDiff_RawMapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_S20_061_HashmapDiff_RawMapStringAnyDiff", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"k": "v"})

		// Act
		result := hd.RawMapStringAnyDiff()

		// Assert
		if result == nil || len(result) != 1 {
			t.Error("RawMapStringAnyDiff returns map with 1 item")
		}
	})
}

func Test_S20_062_HashmapDiff_Serialize(t *testing.T) {
	safeTest(t, "Test_S20_062_HashmapDiff_Serialize", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		data, err := hd.Serialize()

		// Assert
		if err != nil {
			t.Errorf("Serialize returns no error, got %v", err)
		}
		if len(data) == 0 {
			t.Error("Serialize returns non-empty bytes")
		}
	})
}

func Test_S20_063_HashmapDiff_Deserialize(t *testing.T) {
	safeTest(t, "Test_S20_063_HashmapDiff_Deserialize", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		var target map[string]string
		err := hd.Deserialize(&target)

		// Assert
		if err != nil {
			t.Errorf("Deserialize returns no error, got %v", err)
		}
		if target["a"] != "1" {
			t.Errorf("Deserialize target has a=1, got %s", target["a"])
		}
	})
}
