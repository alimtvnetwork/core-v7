package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection iterators — All, Values, Backward
// Covers CollectionIter.go L9-16, L22-29, L35-42
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_All_Iterator(t *testing.T) {
	// Arrange
	col := coregeneric.New.Collection.String.Items("a", "b", "c")
	count := 0
	for i, item := range col.All() {
		_ = i
		_ = item
		count++
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- All iterator", actual)
}

func Test_Collection_All_BreakEarly(t *testing.T) {
	// Arrange
	col := coregeneric.New.Collection.String.Items("a", "b", "c")
	count := 0
	for _, _ = range col.All() {
		count++
		if count == 1 {
			break
		}
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- All break early", actual)
}

func Test_Collection_Values_Iterator(t *testing.T) {
	// Arrange
	col := coregeneric.New.Collection.String.Items("x", "y")
	count := 0
	for item := range col.Values() {
		_ = item
		count++
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- Values iterator", actual)
}

func Test_Collection_Values_BreakEarly(t *testing.T) {
	// Arrange
	col := coregeneric.New.Collection.String.Items("x", "y", "z")
	count := 0
	for _ = range col.Values() {
		count++
		if count == 1 {
			break
		}
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- Values break", actual)
}

func Test_Collection_Backward_Iterator(t *testing.T) {
	// Arrange
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)
	var items []int
	for _, item := range col.Backward() {
		items = append(items, item)
	}

	// Act
	actual := args.Map{
		"first": items[0],
		"last": items[2],
	}

	// Assert
	expected := args.Map{
		"first": 30,
		"last": 10,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Backward", actual)
}

func Test_Collection_Backward_BreakEarly(t *testing.T) {
	// Arrange
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)
	count := 0
	for _, _ = range col.Backward() {
		count++
		if count == 1 {
			break
		}
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Backward break", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap iterators — All, IterKeys, IterValues
// Covers HashmapIter.go L8-16, L21-29, L34-42
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_All_Iterator(t *testing.T) {
	// Arrange
	hm := coregeneric.New.Hashmap.StringString.Cap(2)
	hm.Set("a", "1")
	hm.Set("b", "2")
	count := 0
	for _, _ = range hm.All() {
		count++
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- All", actual)
}

func Test_Hashmap_All_BreakEarly(t *testing.T) {
	// Arrange
	hm := coregeneric.New.Hashmap.StringString.Cap(2)
	hm.Set("a", "1")
	hm.Set("b", "2")
	count := 0
	for _, _ = range hm.All() {
		count++
		break
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- All break", actual)
}

func Test_Hashmap_IterKeys_BreakEarly(t *testing.T) {
	// Arrange
	hm := coregeneric.New.Hashmap.StringString.Cap(2)
	hm.Set("a", "1")
	hm.Set("b", "2")
	count := 0
	for _ = range hm.IterKeys() {
		count++
		break
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- IterKeys break", actual)
}

func Test_Hashmap_IterValues_BreakEarly(t *testing.T) {
	// Arrange
	hm := coregeneric.New.Hashmap.StringString.Cap(2)
	hm.Set("a", "1")
	hm.Set("b", "2")
	count := 0
	for _ = range hm.IterValues() {
		count++
		break
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- IterValues break", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset iterators — All, Values
// Covers HashsetIter.go L9-16, L22-29
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashset_All_BreakEarly(t *testing.T) {
	// Arrange
	hs := coregeneric.New.Hashset.String.Items("a", "b")
	count := 0
	for _, _ = range hs.All() {
		count++
		break
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- All break", actual)
}

func Test_Hashset_Values_BreakEarly(t *testing.T) {
	// Arrange
	hs := coregeneric.New.Hashset.String.Items("a", "b")
	count := 0
	for _ = range hs.Values() {
		count++
		break
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- Values break", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice iterators — All, Values
// Covers SimpleSliceIter.go L8-15, L21-28
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleSlice_All_BreakEarly(t *testing.T) {
	// Arrange
	ss := coregeneric.New.SimpleSlice.String.Items("x", "y", "z")
	count := 0
	for _, _ = range ss.All() {
		count++
		break
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- All break", actual)
}

func Test_SimpleSlice_Values_BreakEarly(t *testing.T) {
	// Arrange
	ss := coregeneric.New.SimpleSlice.String.Items("x", "y", "z")
	count := 0
	for _ = range ss.Values() {
		count++
		break
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns non-empty -- Values break", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList iterators — All, Values
// Covers LinkedListIter.go L8-21, L27-38
// ══════════════════════════════════════════════════════════════════════════════

func Test_LinkedList_All_Iterator(t *testing.T) {
	// Arrange
	ll := coregeneric.New.LinkedList.String.Empty()
	ll.Add("a")
	ll.Add("b")
	ll.Add("c")
	count := 0
	for _, _ = range ll.All() {
		count++
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- All", actual)
}

func Test_LinkedList_All_BreakEarly(t *testing.T) {
	// Arrange
	ll := coregeneric.New.LinkedList.String.Empty()
	ll.Add("a")
	ll.Add("b")
	count := 0
	for _, _ = range ll.All() {
		count++
		break
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- All break", actual)
}

func Test_LinkedList_Values_BreakEarly(t *testing.T) {
	// Arrange
	ll := coregeneric.New.LinkedList.String.Empty()
	ll.Add("a")
	ll.Add("b")
	count := 0
	for _ = range ll.Values() {
		count++
		break
	}

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "LinkedList returns non-empty -- Values break", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList — nodeAt end-of-list branch
// Covers LinkedList.go L319-321
// ══════════════════════════════════════════════════════════════════════════════

// This branch is practically dead — nodeAt checks index < length first,
// so the inner isEndOfList can't trigger. Noted as dead code.

// ══════════════════════════════════════════════════════════════════════════════
// numericfuncs — MaxOf, MaxOfSlice
// Covers numericfuncs.go L151-153, L178-180
// ══════════════════════════════════════════════════════════════════════════════

func Test_MaxOf_FirstIsGreater(t *testing.T) {
	// Arrange
	result := coregeneric.MaxOf(10, 5)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "MaxOf returns correct value -- first greater", actual)
}

func Test_MaxOfSlice_MultipleValues(t *testing.T) {
	// Arrange
	result := coregeneric.MaxOfSlice([]int{3, 7, 1, 9, 2})

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": 9}
	expected.ShouldBeEqual(t, 0, "MaxOfSlice returns correct value -- with args", actual)
}
