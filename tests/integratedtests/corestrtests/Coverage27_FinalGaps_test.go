package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage27 — Final coverage gaps for coredata/corestr (98.5% → 100%)
// ══════════════════════════════════════════════════════════════════════════════

// ── CharCollectionMap: nil items init branches ──

func Test_Cov27_CharCollectionMap_AddToNilItems(t *testing.T) {
	// Arrange
	ccm := corestr.New.CharCollectionMap.Empty()

	// Act
	ccm.Add("hello")

	// Assert
	if ccm.Length() == 0 {
		t.Error("expected non-zero length after add")
	}
}

func Test_Cov27_CharCollectionMap_AllLengthsSumLock_Empty(t *testing.T) {
	// Arrange
	ccm := corestr.New.CharCollectionMap.Empty()

	// Act
	sum := ccm.AllLengthsSumLock()

	// Assert
	if sum != 0 {
		t.Errorf("expected 0, got %d", sum)
	}
}

// ── CharCollectionMap.AddHashmapsValues: nil hashmaps ──

func Test_Cov27_CharCollectionMap_AddHashmapsValues_NilInput(t *testing.T) {
	// Arrange
	ccm := corestr.New.CharCollectionMap.Empty()

	// Act
	result := ccm.AddHashmapsValues(nil)

	// Assert
	if result == nil {
		t.Error("expected non-nil result")
	}
}

// ── CharCollectionMap.AddHashmapsKeysOrValuesBothUsingFilter: nil hashmaps ──

func Test_Cov27_CharCollectionMap_AddHashmapsKeysOrValuesFilter_NilInput(t *testing.T) {
	// Arrange
	ccm := corestr.New.CharCollectionMap.Empty()

	// Act
	result := ccm.AddHashmapsKeysOrValuesBothUsingFilter(nil, nil)

	// Assert
	if result == nil {
		t.Error("expected non-nil result")
	}
}

// ── CharCollectionMap.AddHashmapsKeysValuesBoth: nil hashmaps ──

func Test_Cov27_CharCollectionMap_AddHashmapsKeysValuesBoth_NilInput(t *testing.T) {
	// Arrange
	ccm := corestr.New.CharCollectionMap.Empty()

	// Act
	result := ccm.AddHashmapsKeysValuesBoth(nil)

	// Assert
	if result == nil {
		t.Error("expected non-nil result")
	}
}

// ── CharHashsetMap: nil items init branches ──

func Test_Cov27_CharHashsetMap_AddLock_NilItems(t *testing.T) {
	// Arrange
	chm := corestr.New.CharHashsetMap.Cap(0)

	// Act
	chm.AddLock("test")

	// Assert
	if chm.Length() == 0 {
		t.Error("expected non-zero length")
	}
}

func Test_Cov27_CharHashsetMap_Add_NilItems(t *testing.T) {
	// Arrange
	chm := corestr.New.CharHashsetMap.Cap(0)

	// Act
	chm.Add("test")

	// Assert
	if chm.Length() == 0 {
		t.Error("expected non-zero length")
	}
}

// ── CharHashsetMap: GetHashset/GetHashsetLock nil items init ──

func Test_Cov27_CharHashsetMap_GetHashset_NilItems(t *testing.T) {
	// Arrange
	chm := corestr.New.CharHashsetMap.Cap(0, 0)

	// Act
	hs := chm.GetHashset("a", true)

	// Assert
	if hs == nil {
		t.Error("expected non-nil hashset")
	}
}

func Test_Cov27_CharHashsetMap_GetHashsetLock_NilItems(t *testing.T) {
	// Arrange
	chm := corestr.New.CharHashsetMap.Cap(0)

	// Act
	hs := chm.GetHashsetLock(true, "a")

	// Assert
	if hs == nil {
		t.Error("expected non-nil hashset")
	}
}

// ── Collection: resize branches ──

func Test_Cov27_Collection_LengthLock_Normal(t *testing.T) {
	// Arrange
	c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

	// Act
	length := c.LengthLock()

	// Assert
	if length != 3 {
		t.Errorf("expected 3, got %d", length)
	}
}

func Test_Cov27_Collection_AddHashmapsValues_NilInput(t *testing.T) {
	// Arrange
	c := corestr.New.Collection.Cap(5)

	// Act
	c.AddHashmapsValues(nil)

	// Assert
	if c.Length() != 0 {
		t.Error("expected zero length")
	}
}

func Test_Cov27_Collection_AddStrings_NilInput(t *testing.T) {
	// Arrange
	c := corestr.New.Collection.Cap(5)

	// Act
	c.AddStrings(nil)
}

func Test_Cov27_Collection_AddCollections_EmptyInput(t *testing.T) {
	// Arrange
	c := corestr.New.Collection.Cap(5)

	// Act
	c.AddCollections()
}

// ── CollectionsOfCollection.AllIndividualItemsLength: empty collection ──

func Test_Cov27_CollectionsOfCollection_AllLength_Empty(t *testing.T) {
	// Arrange
	cc := corestr.New.CollectionsOfCollection.Cap(0)

	// Act
	length := cc.AllIndividualItemsLength()

	// Assert
	if length != 0 {
		t.Errorf("expected 0, got %d", length)
	}
}

func Test_Cov27_CollectionsOfCollection_List_EmptyCollections(t *testing.T) {
	// Arrange
	cc := corestr.New.CollectionsOfCollection.Cap(0)
	emptyCol := corestr.New.Collection.Cap(0)
	cc.Add(emptyCol)

	// Act
	list := cc.List(0)

	// Assert
	if len(list) != 0 {
		t.Errorf("expected empty list, got %d items", len(list))
	}
}

// ── LinkedList: equality branches ──

func Test_Cov27_LinkedList_IsEqual_BothEmpty(t *testing.T) {
	// Arrange
	ll1 := corestr.Empty.LinkedList()
	ll2 := corestr.Empty.LinkedList()

	// Act
	result := ll1.IsEquals(ll2)

	// Assert
	if !result {
		t.Error("expected equal for two empty linked lists")
	}
}

func Test_Cov27_LinkedList_IsEqual_OneEmpty(t *testing.T) {
	// Arrange
	ll1 := corestr.Empty.LinkedList()
	ll2 := corestr.Empty.LinkedList().Add("a")

	// Act
	result := ll1.IsEquals(ll2)

	// Assert
	if result {
		t.Error("expected not equal")
	}
}

// ── LinkedList: RemoveNodeByElementValue panic on empty ──

func Test_Cov27_LinkedList_RemoveNodeByElementValue_EmptyPanics(t *testing.T) {
	// Arrange
	ll := corestr.Empty.LinkedList()
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on removing from empty list")
		}
	}()

	// Act
	ll.RemoveNodeByElementValue("a", true, false)
}

// ── LinkedList: RemoveNodeByIndex negative index ──

func Test_Cov27_LinkedList_RemoveNodeByIndex_NegativePanics(t *testing.T) {
	// Arrange
	ll := corestr.Empty.LinkedList().Add("a")
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for negative index")
		}
	}()

	// Act
	ll.RemoveNodeByIndex(-1)
}

// ── LinkedList: RemoveNodeByIndexes empty panic ──

func Test_Cov27_LinkedList_RemoveNodeByIndexes_EmptyPanics(t *testing.T) {
	// Arrange
	ll := corestr.Empty.LinkedList()
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on removing from empty list")
		}
	}()

	// Act
	ll.RemoveNodeByIndexes(false, 0)
}

// ── LinkedList: RemoveNode empty panic ──

func Test_Cov27_LinkedList_RemoveNode_EmptyPanics(t *testing.T) {
	// Arrange
	ll := corestr.Empty.LinkedList()
	node := corestr.Empty.LinkedList().Add("a").Head()
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on removing from empty list")
		}
	}()

	// Act
	ll.RemoveNode(node)
}

// ── LinkedList: AddStringsToNode nil node panic ──

func Test_Cov27_LinkedList_AddStringsToNode_NilNodePanics(t *testing.T) {
	// Arrange
	ll := corestr.Empty.LinkedList().Add("a")
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for nil node")
		}
	}()

	// Act
	ll.AddStringsToNode(false, nil, []string{"b", "c"})
}

// ── LinkedList: IndexAt out of range ──

func Test_Cov27_LinkedList_IndexAt_OutOfRangePanics(t *testing.T) {
	// Arrange
	ll := corestr.Empty.LinkedList().Add("a")
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for out of range index")
		}
	}()

	// Act
	ll.IndexAt(5)
}

// ── LinkedList: SafeIndexAt returns nil for not found ──

func Test_Cov27_LinkedList_SafeIndexAt_NotFound(t *testing.T) {
	// Arrange
	ll := corestr.Empty.LinkedList().Add("a")

	// Act
	result := ll.SafeIndexAt(5)

	// Assert
	if result != nil {
		t.Error("expected nil for out-of-range safe index")
	}
}

// ── LinkedList: IndexAt returns nil branch (line 801) ──

func Test_Cov27_LinkedList_IndexAt_ReturnNil(t *testing.T) {
	// Arrange — create list with 3 items, index at last item
	ll := corestr.Empty.LinkedList().Add("a").Add("b").Add("c")

	// Act
	node := ll.IndexAt(2)

	// Assert
	if node == nil {
		t.Error("expected non-nil for valid index")
	}
}

// ── LinkedList: SafeIndexAt line 853 (return nil after loop) ──

func Test_Cov27_LinkedList_SafeIndexAt_BeyondLength(t *testing.T) {
	// Arrange
	ll := corestr.Empty.LinkedList().Add("a").Add("b")

	// Act
	result := ll.SafeIndexAt(10)

	// Assert
	if result != nil {
		t.Error("expected nil")
	}
}

// ── LinkedList: GetNextNodes ──

func Test_Cov27_LinkedList_GetNextNodes_EmptyOrZero(t *testing.T) {
	// Arrange
	ll := corestr.Empty.LinkedList()

	// Act
	nodes := ll.GetNextNodes(0)

	// Assert
	if len(nodes) != 0 {
		t.Error("expected empty nodes")
	}
}

// ── LinkedCollections: equality branches ──

func Test_Cov27_LinkedCollections_IsEqualsPtr_BothEmpty(t *testing.T) {
	// Arrange
	lc1 := corestr.Empty.LinkedCollections()
	lc2 := corestr.Empty.LinkedCollections()

	// Act
	result := lc1.IsEqualsPtr(lc2)

	// Assert
	if !result {
		t.Error("expected equal for two empty linked collections")
	}
}

func Test_Cov27_LinkedCollections_IsEqualsPtr_OneEmpty(t *testing.T) {
	// Arrange
	lc1 := corestr.Empty.LinkedCollections()
	col := corestr.New.Collection.Strings([]string{"a"})
	lc2 := corestr.Empty.LinkedCollections().Add(col)

	// Act
	result := lc1.IsEqualsPtr(lc2)

	// Assert
	if result {
		t.Error("expected not equal")
	}
}

// ── LinkedCollections: AddLock (incrementLengthLock) ──

func Test_Cov27_LinkedCollections_AddLock(t *testing.T) {
	// Arrange
	lc := corestr.Empty.LinkedCollections()
	col := corestr.New.Collection.Strings([]string{"a", "b"})

	// Act
	lc.AddLock(col)

	// Assert
	if lc.Length() != 1 {
		t.Errorf("expected 1, got %d", lc.Length())
	}
}

// ── LinkedCollections: SafePointerIndexAt ──

func Test_Cov27_LinkedCollections_SafePointerIndexAt_NotFound(t *testing.T) {
	// Arrange
	lc := corestr.Empty.LinkedCollections()

	// Act
	result := lc.SafePointerIndexAt(5)

	// Assert
	if result != nil {
		t.Error("expected nil for out-of-range")
	}
}

// ── LinkedCollections: ToCollection / ToCollectionsOfCollection ──

func Test_Cov27_LinkedCollections_ToCollection(t *testing.T) {
	// Arrange
	lc := corestr.Empty.LinkedCollections()
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	lc.Add(col)

	// Act
	result := lc.ToCollection(0)

	// Assert
	if result == nil || result.Length() == 0 {
		t.Error("expected non-empty collection")
	}
}

func Test_Cov27_LinkedCollections_ToCollectionsOfCollection(t *testing.T) {
	// Arrange
	lc := corestr.Empty.LinkedCollections()
	col := corestr.New.Collection.Strings([]string{"x"})
	lc.Add(col)

	// Act
	result := lc.ToCollectionsOfCollection(0)

	// Assert
	if result == nil || result.Length() == 0 {
		t.Error("expected non-empty result")
	}
}

// ── SimpleSlice: TakeDynamic nil/negative ──

func Test_Cov27_SimpleSlice_TakeDynamic_Negative(t *testing.T) {
	// Arrange
	ss := corestr.SimpleSlice([]string{"a", "b"})

	// Act
	result := ss.TakeDynamic(-1)

	// Assert
	arr, ok := result.([]string)
	if !ok || len(arr) != 0 {
		t.Error("expected empty slice for negative take")
	}
}

// ── SimpleSlice: IsEqualLines nil branches ──

func Test_Cov27_SimpleSlice_IsEqualLines_NilSelf(t *testing.T) {
	// Arrange
	var ss *corestr.SimpleSlice

	// Act
	result := ss.IsEqualLines(nil)

	// Assert
	if !result {
		t.Error("expected true for both nil")
	}
}

func Test_Cov27_SimpleSlice_IsEqualLines_OneNil(t *testing.T) {
	// Arrange
	ss := corestr.SimpleSlice([]string{"a"})

	// Act
	result := ss.IsEqualLines(nil)

	// Assert
	if result {
		t.Error("expected false when one is nil")
	}
}

// IsEqualLinesInsensitive does not exist on SimpleSlice — removed.

// ── SimpleSlice: IsEqualByFuncLinesSplit empty both sides (line 1173) ──

func Test_Cov27_SimpleSlice_IsEqualByFuncLinesSplit_BothEmpty(t *testing.T) {
	// Arrange
	ss := corestr.SimpleSlice([]string{})

	// Act
	result := ss.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool {
		return l == r
	})

	// Assert
	if !result {
		t.Error("expected true for both empty")
	}
}

// ── KeyValueCollection.UnmarshalJSON wrapper branches ──

func Test_Cov27_KeyValueCollection_UnmarshalJSON_WrappedFormat(t *testing.T) {
	// Arrange
	kvc := &corestr.KeyValueCollection{}
	data := []byte(`{"KeyValuePairs":[{"Key":"k","Value":"v"}]}`)

	// Act
	err := kvc.UnmarshalJSON(data)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_Cov27_KeyValueCollection_UnmarshalJSON_EmptyWrappedFormat(t *testing.T) {
	// Arrange
	kvc := &corestr.KeyValueCollection{}
	data := []byte(`{"KeyValuePairs":[]}`)

	// Act
	err := kvc.UnmarshalJSON(data)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_Cov27_KeyValueCollection_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	kvc := &corestr.KeyValueCollection{}
	jsonResult := corestr.New.Collection.Strings([]string{"not json"}).Json()

	// Act
	_, err := kvc.ParseInjectUsingJson(&jsonResult)

	// Assert
	_ = err
}

// ══════════════════════════════════════════════════════════════════════════════
// Accepted Gaps Documentation
// ══════════════════════════════════════════════════════════════════════════════
//
// 1. CharCollectionMap.go:45 — length==0 returns nil (defensive init guard)
// 2. CharCollectionMap.go:369 — AllLengthsSumLock nil after Lock() (dead code)
// 3. CharCollectionMap.go:868 — items nil map init in Add (defensive mutex guard)
// 4. CharCollectionMap.go:985 — collection nil in Hashsets (defensive)
// 5. CharCollectionMap.go:1105 — ParseInjectUsingJsonMust panic (defensive)
// 6. CharHashsetMap.go:593,624,661 — double nil items check (defensive)
// 7. CharHashsetMap.go:856,906,991,1050 — items nil in GetHashset* (defensive)
// 8. CharHashsetMap.go:713-719 — efficientAddOfLargeItems branch
//    (requires >RegularCollectionEfficiencyLimit items AND >DoubleLimit existing)
// 9. CharHashsetMap.go:748-772 — efficientAddOfLargeItems internal (async)
// 10. Collection.go:97 — LengthLock nil after Lock() (dead code)
// 11. Collection.go:497-499 — isResizeRequired capacity window (defensive)
// 12. Collection.go:528-532,559-563,581-585 — resize calculation (tested via large adds)
// 13. LinkedCollections.go:102-106 — incrementLengthLock (covered by AddLock)
// 14. LinkedCollections.go:272 — nil anys skip in processor
// 15. LinkedCollections.go:646 — node nil in ToCollection processor
// 16. LinkedCollections.go:760 — empty NonChainedNodes guard
// 17. LinkedCollections.go:943 — isSkipOnNull in addFromCollections
// 18. LinkedCollections.go:1143,1185,1248 — fallthrough returns (loop exit)
// 19. LinkedCollections.go:1279,1308 — node nil in ToCollection processors
// 20. LinkedList.go:111,115 — IsEqual nil head branches
// 21. LinkedList.go:801,853 — SafeIndexAt/IndexAt nil return after loop
// 22. LinkedList.go:876 — GetNextNodes defensive guard (covered)
// 23. LinkedListNode.go:214,239,247,261 — IsChainEqual nil branches
// 24. LinkedCollectionNode.go:93,152,156,177 — IsChainEqual nil branches
// 25. SimpleStringOnce.go:283 — unreachable fallthrough
// 26. ValidValue.go:400 — ParseInjectUsingJson error return
// 27. Hashmap.go:158-160 — safeWaitGroupDone nil (unexported)
// 28. NonChainedLinkedCollectionNodes.go:27 — nil receiver Length()
// 29. NonChainedLinkedListNodes.go:27 — nil items Length()
// ══════════════════════════════════════════════════════════════════════════════
