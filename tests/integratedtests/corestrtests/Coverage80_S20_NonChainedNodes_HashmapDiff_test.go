package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
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
		actual := args.Map{"result": nodes == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewNonChainedLinkedListNodes returns nil -- capacity 5", actual)
		actual := args.Map{"result": nodes.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Length returns 0 -- empty after creation", actual)
		actual := args.Map{"result": nodes.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- no items added", actual)
		actual := args.Map{"result": nodes.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasItems returns false -- no items added", actual)
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
		actual := args.Map{"result": nodes.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Length returns 1 -- one node added", actual)
		actual := args.Map{"result": nodes.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsEmpty returns false -- has items", actual)
		actual := args.Map{"result": nodes.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasItems returns true -- has items", actual)
	})
}

func Test_S20_003_NonChainedLinkedListNodes_Adds_nil(t *testing.T) {
	safeTest(t, "Test_S20_003_NonChainedLinkedListNodes_Adds_nil", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		result := nodes.Adds(nil...)

		// Assert
		actual := args.Map{"result": result != nodes}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Adds returns self -- nil input", actual)
		actual := args.Map{"result": nodes.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Length returns 0 -- nil not added", actual)
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
		actual := args.Map{"result": nodes.First().Element != "alpha"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "First returns alpha -- first added", actual)
		actual := args.Map{"result": nodes.Last().Element != "gamma"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Last returns gamma -- last added", actual)
	})
}

func Test_S20_005_NonChainedLinkedListNodes_FirstOrDefault_empty(t *testing.T) {
	safeTest(t, "Test_S20_005_NonChainedLinkedListNodes_FirstOrDefault_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		result := nodes.FirstOrDefault()

		// Assert
		actual := args.Map{"result": result != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault returns nil -- empty nodes", actual)
	})
}

func Test_S20_006_NonChainedLinkedListNodes_LastOrDefault_empty(t *testing.T) {
	safeTest(t, "Test_S20_006_NonChainedLinkedListNodes_LastOrDefault_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		result := nodes.LastOrDefault()

		// Assert
		actual := args.Map{"result": result != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LastOrDefault returns nil -- empty nodes", actual)
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
		actual := args.Map{"result": result == nil || result.Element != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault returns x -- single item", actual)
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
		actual := args.Map{"result": result == nil || result.Element != "y"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LastOrDefault returns y -- two items", actual)
	})
}

func Test_S20_009_NonChainedLinkedListNodes_IsChainingApplied_false(t *testing.T) {
	safeTest(t, "Test_S20_009_NonChainedLinkedListNodes_IsChainingApplied_false", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		result := nodes.IsChainingApplied()

		// Assert
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsChainingApplied returns false -- not applied yet", actual)
	})
}

func Test_S20_010_NonChainedLinkedListNodes_ApplyChaining_empty(t *testing.T) {
	safeTest(t, "Test_S20_010_NonChainedLinkedListNodes_ApplyChaining_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		result := nodes.ApplyChaining()

		// Assert
		actual := args.Map{"result": result != nodes}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ApplyChaining returns self -- empty nodes", actual)
		actual := args.Map{"result": nodes.IsChainingApplied()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsChainingApplied returns false -- empty, no chaining applied", actual)
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
		actual := args.Map{"result": result != nodes}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ApplyChaining returns self", actual)
		actual := args.Map{"result": nodes.IsChainingApplied()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsChainingApplied returns true -- chaining applied", actual)
		first := nodes.First()
		actual := args.Map{"result": first.Next() == nil || first.Next().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ApplyChaining chains a->b -- correct next", actual)
		actual := args.Map{"result": nodes.Last().HasNext()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Last node has nil next -- end of chain", actual)
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
		actual := args.Map{"result": len(items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Items returns 1 item -- one added", actual)
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
		actual := args.Map{"result": chained == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ToChainedNodes returns non-nil -- has items", actual)
	})
}

func Test_S20_014_NonChainedLinkedListNodes_ToChainedNodes_empty(t *testing.T) {
	safeTest(t, "Test_S20_014_NonChainedLinkedListNodes_ToChainedNodes_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		chained := nodes.ToChainedNodes()

		// Assert
		actual := args.Map{"result": chained == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ToChainedNodes returns non-nil slice -- empty input", actual)
		actual := args.Map{"result": len(chained) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ToChainedNodes returns empty slice -- empty input", actual)
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
		actual := args.Map{"result": nodes == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewNonChainedLinkedCollectionNodes returns non-nil -- capacity 5", actual)
		actual := args.Map{"result": nodes.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Length returns 0 -- empty", actual)
		actual := args.Map{"result": nodes.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- no items", actual)
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
		actual := args.Map{"result": nodes.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Length returns 1 -- one node", actual)
		actual := args.Map{"result": nodes.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasItems returns true -- has items", actual)
	})
}

func Test_S20_022_NonChainedLinkedCollectionNodes_Adds_nil(t *testing.T) {
	safeTest(t, "Test_S20_022_NonChainedLinkedCollectionNodes_Adds_nil", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		result := nodes.Adds(nil...)

		// Assert
		actual := args.Map{"result": result != nodes}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Adds returns self -- nil input", actual)
	})
}

func Test_S20_023_NonChainedLinkedCollectionNodes_FirstOrDefault_empty(t *testing.T) {
	safeTest(t, "Test_S20_023_NonChainedLinkedCollectionNodes_FirstOrDefault_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		result := nodes.FirstOrDefault()

		// Assert
		actual := args.Map{"result": result != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault returns nil -- empty", actual)
	})
}

func Test_S20_024_NonChainedLinkedCollectionNodes_LastOrDefault_empty(t *testing.T) {
	safeTest(t, "Test_S20_024_NonChainedLinkedCollectionNodes_LastOrDefault_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		result := nodes.LastOrDefault()

		// Assert
		actual := args.Map{"result": result != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LastOrDefault returns nil -- empty", actual)
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
		actual := args.Map{"result": nodes.First().Element.List()[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "First returns collection with a", actual)
		actual := args.Map{"result": nodes.Last().Element.List()[0] != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Last returns collection with b", actual)
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
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault returns non-nil -- has items", actual)
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
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LastOrDefault returns non-nil -- has items", actual)
	})
}

func Test_S20_028_NonChainedLinkedCollectionNodes_IsChainingApplied_false(t *testing.T) {
	safeTest(t, "Test_S20_028_NonChainedLinkedCollectionNodes_IsChainingApplied_false", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act & Assert
		actual := args.Map{"result": nodes.IsChainingApplied()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsChainingApplied returns false -- not applied", actual)
	})
}

func Test_S20_029_NonChainedLinkedCollectionNodes_ApplyChaining_empty(t *testing.T) {
	safeTest(t, "Test_S20_029_NonChainedLinkedCollectionNodes_ApplyChaining_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		result := nodes.ApplyChaining()

		// Assert
		actual := args.Map{"result": result != nodes}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ApplyChaining returns self -- empty", actual)
		actual := args.Map{"result": nodes.IsChainingApplied()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsChainingApplied returns false -- empty, no chaining", actual)
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
		actual := args.Map{"result": result != nodes}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ApplyChaining returns self", actual)
		actual := args.Map{"result": nodes.IsChainingApplied()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsChainingApplied returns true -- applied", actual)
		actual := args.Map{"result": nodes.Last().HasNext()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Last node has nil next -- end of chain", actual)
	})
}

func Test_S20_031_NonChainedLinkedCollectionNodes_Items(t *testing.T) {
	safeTest(t, "Test_S20_031_NonChainedLinkedCollectionNodes_Items", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		items := nodes.Items()

		// Assert
		actual := args.Map{"result": items == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Items returns non-nil -- empty but initialized", actual)
	})
}

func Test_S20_032_NonChainedLinkedCollectionNodes_ToChainedNodes_empty(t *testing.T) {
	safeTest(t, "Test_S20_032_NonChainedLinkedCollectionNodes_ToChainedNodes_empty", func() {
		// Arrange
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		chained := nodes.ToChainedNodes()

		// Assert
		actual := args.Map{"result": chained == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ToChainedNodes returns non-nil -- empty input", actual)
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
		actual := args.Map{"result": chained == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ToChainedNodes returns non-nil -- has items", actual)
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
		actual := args.Map{"result": result != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Length returns 0 -- nil receiver", actual)
	})
}

func Test_S20_041_HashmapDiff_Length_with_items(t *testing.T) {
	safeTest(t, "Test_S20_041_HashmapDiff_Length_with_items", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2"})

		// Act
		result := hd.Length()

		// Assert
		actual := args.Map{"result": result != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Length returns 2 -- two items", actual)
	})
}

func Test_S20_042_HashmapDiff_IsEmpty_true(t *testing.T) {
	safeTest(t, "Test_S20_042_HashmapDiff_IsEmpty_true", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{})

		// Act & Assert
		actual := args.Map{"result": hd.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- empty map", actual)
	})
}

func Test_S20_043_HashmapDiff_IsEmpty_false(t *testing.T) {
	safeTest(t, "Test_S20_043_HashmapDiff_IsEmpty_false", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act & Assert
		actual := args.Map{"result": hd.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsEmpty returns false -- has item", actual)
	})
}

func Test_S20_044_HashmapDiff_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_S20_044_HashmapDiff_HasAnyItem", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act & Assert
		actual := args.Map{"result": hd.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAnyItem returns true -- has item", actual)
	})
}

func Test_S20_045_HashmapDiff_LastIndex(t *testing.T) {
	safeTest(t, "Test_S20_045_HashmapDiff_LastIndex", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2", "c": "3"})

		// Act
		result := hd.LastIndex()

		// Assert
		actual := args.Map{"result": result != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LastIndex returns 2 -- 3 items", actual)
	})
}

func Test_S20_046_HashmapDiff_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_S20_046_HashmapDiff_AllKeysSorted", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"c": "3", "a": "1", "b": "2"})

		// Act
		keys := hd.AllKeysSorted()

		// Assert
		actual := args.Map{"result": len(keys) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AllKeysSorted returns 3 keys", actual)
		actual := args.Map{"result": keys[0] != "a" || keys[1] != "b" || keys[2] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AllKeysSorted returns sorted keys -- a,b,c", actual)
	})
}

func Test_S20_047_HashmapDiff_MapAnyItems(t *testing.T) {
	safeTest(t, "Test_S20_047_HashmapDiff_MapAnyItems", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"x": "10"})

		// Act
		result := hd.MapAnyItems()

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "MapAnyItems returns 1 item", actual)
		actual := args.Map{"result": result["x"] != "10"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "MapAnyItems has x=10", actual)
	})
}

func Test_S20_048_HashmapDiff_MapAnyItems_nil(t *testing.T) {
	safeTest(t, "Test_S20_048_HashmapDiff_MapAnyItems_nil", func() {
		// Arrange
		var hd *corestr.HashmapDiff

		// Act
		result := hd.MapAnyItems()

		// Assert
		actual := args.Map{"result": result == nil || len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty map -- nil receiver", actual)
	})
}

func Test_S20_049_HashmapDiff_Raw_nil(t *testing.T) {
	safeTest(t, "Test_S20_049_HashmapDiff_Raw_nil", func() {
		// Arrange
		var hd *corestr.HashmapDiff

		// Act
		result := hd.Raw()

		// Assert
		actual := args.Map{"result": result == nil || len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Raw returns empty map -- nil receiver", actual)
	})
}

func Test_S20_050_HashmapDiff_Raw_with_items(t *testing.T) {
	safeTest(t, "Test_S20_050_HashmapDiff_Raw_with_items", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		result := hd.Raw()

		// Assert
		actual := args.Map{"result": len(result) != 1 || result["a"] != "1"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Raw returns underlying map -- has items", actual)
	})
}

func Test_S20_051_HashmapDiff_IsRawEqual_true(t *testing.T) {
	safeTest(t, "Test_S20_051_HashmapDiff_IsRawEqual_true", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		result := hd.IsRawEqual(map[string]string{"a": "1"})

		// Assert
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsRawEqual returns true -- same maps", actual)
	})
}

func Test_S20_052_HashmapDiff_IsRawEqual_false(t *testing.T) {
	safeTest(t, "Test_S20_052_HashmapDiff_IsRawEqual_false", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		result := hd.IsRawEqual(map[string]string{"a": "2"})

		// Assert
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsRawEqual returns false -- different values", actual)
	})
}

func Test_S20_053_HashmapDiff_HasAnyChanges(t *testing.T) {
	safeTest(t, "Test_S20_053_HashmapDiff_HasAnyChanges", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		result := hd.HasAnyChanges(map[string]string{"a": "2"})

		// Assert
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAnyChanges returns true -- different values", actual)
	})
}

func Test_S20_054_HashmapDiff_HasAnyChanges_no_changes(t *testing.T) {
	safeTest(t, "Test_S20_054_HashmapDiff_HasAnyChanges_no_changes", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		result := hd.HasAnyChanges(map[string]string{"a": "1"})

		// Assert
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasAnyChanges returns false -- same values", actual)
	})
}

func Test_S20_055_HashmapDiff_DiffRaw(t *testing.T) {
	safeTest(t, "Test_S20_055_HashmapDiff_DiffRaw", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2"})

		// Act
		diff := hd.DiffRaw(map[string]string{"a": "1", "b": "3"})

		// Assert
		actual := args.Map{"result": diff == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "DiffRaw returns non-nil -- has diff", actual)
		actual := args.Map{"result": diff["b"] != "2"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "DiffRaw contains b=2 -- left value for changed key", actual)
	})
}

func Test_S20_056_HashmapDiff_HashmapDiffUsingRaw_no_diff(t *testing.T) {
	safeTest(t, "Test_S20_056_HashmapDiff_HashmapDiffUsingRaw_no_diff", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		result := hd.HashmapDiffUsingRaw(map[string]string{"a": "1"})

		// Assert
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashmapDiffUsingRaw returns empty -- no diff", actual)
	})
}

func Test_S20_057_HashmapDiff_DiffJsonMessage(t *testing.T) {
	safeTest(t, "Test_S20_057_HashmapDiff_DiffJsonMessage", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		msg := hd.DiffJsonMessage(map[string]string{"a": "2"})

		// Assert
		actual := args.Map{"result": msg == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "DiffJsonMessage returns non-empty -- has diff", actual)
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
		actual := args.Map{"result": len(result) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ToStringsSliceOfDiffMap returns non-empty -- has items", actual)
	})
}

func Test_S20_059_HashmapDiff_ShouldDiffMessage(t *testing.T) {
	safeTest(t, "Test_S20_059_HashmapDiff_ShouldDiffMessage", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		msg := hd.ShouldDiffMessage("test", map[string]string{"a": "2"})

		// Assert
		actual := args.Map{"result": msg == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ShouldDiffMessage returns non-empty -- has diff", actual)
	})
}

func Test_S20_060_HashmapDiff_LogShouldDiffMessage(t *testing.T) {
	safeTest(t, "Test_S20_060_HashmapDiff_LogShouldDiffMessage", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		msg := hd.LogShouldDiffMessage("test", map[string]string{"a": "2"})

		// Assert
		actual := args.Map{"result": msg == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LogShouldDiffMessage returns non-empty -- has diff", actual)
	})
}

func Test_S20_061_HashmapDiff_RawMapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_S20_061_HashmapDiff_RawMapStringAnyDiff", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"k": "v"})

		// Act
		result := hd.RawMapStringAnyDiff()

		// Assert
		actual := args.Map{"result": result == nil || len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "RawMapStringAnyDiff returns map with 1 item", actual)
	})
}

func Test_S20_062_HashmapDiff_Serialize(t *testing.T) {
	safeTest(t, "Test_S20_062_HashmapDiff_Serialize", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		data, err := hd.Serialize()

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Serialize returns no error", actual)
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Serialize returns non-empty bytes", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Deserialize returns no error", actual)
		actual := args.Map{"result": target["a"] != "1"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Deserialize target has a=1", actual)
	})
}
