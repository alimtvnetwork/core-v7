package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollections — Segment 5: Remaining methods (L800-1551)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovLC2_01_AddCollectionsToNodeAsync(t *testing.T) {
	safeTest(t, "Test_CovLC2_01_AddCollectionsToNodeAsync", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"base"}))
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddCollectionsToNodeAsync(
			true, wg, lc.Head(),
			corestr.New.Collection.Strings([]string{"added"}),
		)
		wg.Wait()
		if lc.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_CovLC2_02_AddCollectionsToNode(t *testing.T) {
	safeTest(t, "Test_CovLC2_02_AddCollectionsToNode", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"base"}))
		lc.AddCollectionsToNode(true, lc.Head(), corestr.New.Collection.Strings([]string{"x"}))
		if lc.Length() < 2 {
			t.Fatal("expected at least 2")
		}
		// nil skip
		lc.AddCollectionsToNode(true, lc.Head())
	})
}

func Test_CovLC2_03_AddCollectionsPointerToNode(t *testing.T) {
	safeTest(t, "Test_CovLC2_03_AddCollectionsPointerToNode", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"base"}))
		cols := []*corestr.Collection{
			corestr.New.Collection.Strings([]string{"a"}),
			corestr.New.Collection.Strings([]string{"b"}),
		}
		lc.AddCollectionsPointerToNode(true, lc.Head(), &cols)
		if lc.Length() < 3 {
			t.Fatalf("expected at least 3, got %d", lc.Length())
		}
		// nil items
		lc.AddCollectionsPointerToNode(true, lc.Head(), nil)
		// nil node with skip
		lc.AddCollectionsPointerToNode(true, nil, &cols)
		// single item
		single := []*corestr.Collection{corestr.New.Collection.Strings([]string{"s"})}
		lc.AddCollectionsPointerToNode(true, lc.Head(), &single)
		// empty items
		empty := []*corestr.Collection{}
		lc.AddCollectionsPointerToNode(true, lc.Head(), &empty)
	})
}

func Test_CovLC2_04_AddAfterNode(t *testing.T) {
	safeTest(t, "Test_CovLC2_04_AddAfterNode", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AddAfterNode(lc.Head(), corestr.New.Collection.Strings([]string{"b"}))
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLC2_05_AddAfterNodeAsync(t *testing.T) {
	safeTest(t, "Test_CovLC2_05_AddAfterNodeAsync", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAfterNodeAsync(wg, lc.Head(), corestr.New.Collection.Strings([]string{"b"}))
		wg.Wait()
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLC2_06_ConcatNew(t *testing.T) {
	safeTest(t, "Test_CovLC2_06_ConcatNew", func() {
		a := corestr.Empty.LinkedCollections()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))

		// empty with clone
		cloned := a.ConcatNew(true)
		if cloned.Length() != 1 {
			t.Fatal("expected 1")
		}
		// empty without clone
		same := a.ConcatNew(false)
		if same.Length() != 1 {
			t.Fatal("expected 1")
		}
		// with others
		b := corestr.Empty.LinkedCollections()
		b.Add(corestr.New.Collection.Strings([]string{"b"}))
		merged := a.ConcatNew(false, b)
		if merged.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLC2_07_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_CovLC2_07_AddAsyncFuncItems", func() {
		lc := corestr.Empty.LinkedCollections()
		wg := &sync.WaitGroup{}
		wg.Add(2)
		lc.AddAsyncFuncItems(wg, false,
			func() []string { return []string{"a"} },
			func() []string { return []string{} }, // empty
		)
		if lc.Length() != 1 {
			t.Fatalf("expected 1, got %d", lc.Length())
		}
		// nil funcs
		lc2 := corestr.Empty.LinkedCollections()
		lc2.AddAsyncFuncItems(wg, false)
	})
}

func Test_CovLC2_08_AddAsyncFuncItemsPointer(t *testing.T) {
	safeTest(t, "Test_CovLC2_08_AddAsyncFuncItemsPointer", func() {
		lc := corestr.Empty.LinkedCollections()
		wg := &sync.WaitGroup{}
		wg.Add(2)
		lc.AddAsyncFuncItemsPointer(wg, false,
			func() []string { return []string{"a"} },
			func() []string { return []string{} },
		)
		if lc.Length() != 1 {
			t.Fatalf("expected 1, got %d", lc.Length())
		}
		lc2 := corestr.Empty.LinkedCollections()
		lc2.AddAsyncFuncItemsPointer(wg, false)
	})
}

func Test_CovLC2_09_AddStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_CovLC2_09_AddStringsOfStrings", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.AddStringsOfStrings(false, []string{"a"}, nil, []string{"b"})
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
		lc.AddStringsOfStrings(false)
	})
}

func Test_CovLC2_10_IndexAt(t *testing.T) {
	safeTest(t, "Test_CovLC2_10_IndexAt", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		// index 0
		n := lc.IndexAt(0)
		if n == nil {
			t.Fatal("expected non-nil")
		}
		// index 1
		n1 := lc.IndexAt(1)
		if n1 == nil {
			t.Fatal("expected non-nil")
		}
		// index 2
		n2 := lc.IndexAt(2)
		if n2 == nil {
			t.Fatal("expected non-nil")
		}
		// negative
		nn := lc.IndexAt(-1)
		if nn != nil {
			t.Fatal("expected nil for negative")
		}
	})
}

func Test_CovLC2_11_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_CovLC2_11_SafePointerIndexAt", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		col := lc.SafePointerIndexAt(0)
		if col == nil {
			t.Fatal("expected non-nil")
		}
		nilCol := lc.SafePointerIndexAt(99)
		if nilCol != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_CovLC2_12_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_CovLC2_12_SafeIndexAt", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		n := lc.SafeIndexAt(0)
		if n == nil {
			t.Fatal("expected non-nil")
		}
		n1 := lc.SafeIndexAt(1)
		if n1 == nil {
			t.Fatal("expected non-nil")
		}
		// out of range
		if lc.SafeIndexAt(-1) != nil {
			t.Fatal("expected nil")
		}
		if lc.SafeIndexAt(99) != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_CovLC2_13_AddStringsAsync(t *testing.T) {
	safeTest(t, "Test_CovLC2_13_AddStringsAsync", func() {
		lc := corestr.Empty.LinkedCollections()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddStringsAsync(wg, []string{"a", "b"})
		wg.Wait()
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		// nil
		lc.AddStringsAsync(wg, nil)
	})
}

func Test_CovLC2_14_AddCollection(t *testing.T) {
	safeTest(t, "Test_CovLC2_14_AddCollection", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		lc.AddCollection(nil)
		if lc.Length() != 1 {
			t.Fatal("expected still 1")
		}
	})
}

func Test_CovLC2_15_AddCollectionsPtr_AddCollections(t *testing.T) {
	safeTest(t, "Test_CovLC2_15_AddCollectionsPtr_AddCollections", func() {
		lc := corestr.Empty.LinkedCollections()
		cols := []*corestr.Collection{corestr.New.Collection.Strings([]string{"a"})}
		lc.AddCollectionsPtr(cols)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		lc.AddCollectionsPtr(nil)
		lc.AddCollections(nil)
		// with nil in slice
		cols2 := []*corestr.Collection{nil, corestr.New.Collection.Strings([]string{"b"})}
		lc2 := corestr.Empty.LinkedCollections()
		lc2.AddCollections(cols2)
	})
}

func Test_CovLC2_16_ToStringsPtr_ToStrings(t *testing.T) {
	safeTest(t, "Test_CovLC2_16_ToStringsPtr_ToStrings", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		ptr := lc.ToStringsPtr()
		if len(*ptr) != 2 {
			t.Fatal("expected 2")
		}
		strs := lc.ToStrings()
		if len(strs) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLC2_17_ToCollectionSimple_ToCollection(t *testing.T) {
	safeTest(t, "Test_CovLC2_17_ToCollectionSimple_ToCollection", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		col := lc.ToCollectionSimple()
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
		col2 := lc.ToCollection(5)
		if col2.Length() != 1 {
			t.Fatal("expected 1")
		}
		// empty
		e := corestr.Empty.LinkedCollections()
		if e.ToCollection(0).Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovLC2_18_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_CovLC2_18_ToCollectionsOfCollection", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		coc := lc.ToCollectionsOfCollection(0)
		if coc.Length() != 2 {
			t.Fatal("expected 2")
		}
		// empty
		e := corestr.Empty.LinkedCollections()
		if e.ToCollectionsOfCollection(0).Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovLC2_19_ItemsOfItems(t *testing.T) {
	safeTest(t, "Test_CovLC2_19_ItemsOfItems", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b", "c"}))
		ii := lc.ItemsOfItems()
		if len(ii) != 2 {
			t.Fatal("expected 2")
		}
		// empty
		e := corestr.Empty.LinkedCollections()
		if len(e.ItemsOfItems()) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovLC2_20_ItemsOfItemsCollection(t *testing.T) {
	safeTest(t, "Test_CovLC2_20_ItemsOfItemsCollection", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		cols := lc.ItemsOfItemsCollection()
		if len(cols) != 1 {
			t.Fatal("expected 1")
		}
		// empty
		e := corestr.Empty.LinkedCollections()
		if len(e.ItemsOfItemsCollection()) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovLC2_21_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_CovLC2_21_SimpleSlice", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		ss := lc.SimpleSlice()
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLC2_22_ListPtr_List(t *testing.T) {
	safeTest(t, "Test_CovLC2_22_ListPtr_List", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		ptr := lc.ListPtr()
		if len(*ptr) != 1 {
			t.Fatal("expected 1")
		}
		list := lc.List()
		if len(list) != 1 {
			t.Fatal("expected 1")
		}
		// empty
		e := corestr.Empty.LinkedCollections()
		if len(e.List()) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovLC2_23_String(t *testing.T) {
	safeTest(t, "Test_CovLC2_23_String", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		s := lc.String()
		if s == "" {
			t.Fatal("expected non-empty")
		}
		// empty
		e := corestr.Empty.LinkedCollections()
		_ = e.String()
	})
}

func Test_CovLC2_24_StringLock(t *testing.T) {
	safeTest(t, "Test_CovLC2_24_StringLock", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		s := lc.StringLock()
		if s == "" {
			t.Fatal("expected non-empty")
		}
		e := corestr.Empty.LinkedCollections()
		_ = e.StringLock()
	})
}

func Test_CovLC2_25_Join(t *testing.T) {
	safeTest(t, "Test_CovLC2_25_Join", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		s := lc.Join(",")
		if s != "a" {
			t.Fatalf("expected 'a', got '%s'", s)
		}
	})
}

func Test_CovLC2_26_Joins(t *testing.T) {
	safeTest(t, "Test_CovLC2_26_Joins", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		// with items
		s := lc.Joins(",", "b")
		if s == "" {
			t.Fatal("expected non-empty")
		}
		// nil items or empty LC
		e := corestr.Empty.LinkedCollections()
		_ = e.Joins(",")
	})
}

func Test_CovLC2_27_JsonModel_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovLC2_27_JsonModel_JsonModelAny", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		m := lc.JsonModel()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
		_ = lc.JsonModelAny()
	})
}

func Test_CovLC2_28_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovLC2_28_MarshalJSON", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		data, err := lc.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected error")
		}
		if len(data) == 0 {
			t.Fatal("expected data")
		}
	})
}

func Test_CovLC2_29_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovLC2_29_UnmarshalJSON", func() {
		lc := corestr.Empty.LinkedCollections()
		err := lc.UnmarshalJSON([]byte(`["a","b"]`))
		if err != nil {
			t.Fatal("unexpected error")
		}
		if lc.Length() != 1 {
			t.Fatal("expected 1 collection with 2 items")
		}
		// invalid
		err2 := lc.UnmarshalJSON([]byte(`invalid`))
		if err2 == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovLC2_30_RemoveAll_Clear(t *testing.T) {
	safeTest(t, "Test_CovLC2_30_RemoveAll_Clear", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.RemoveAll()
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
		// clear empty
		e := corestr.Empty.LinkedCollections()
		e.Clear()
	})
}

func Test_CovLC2_31_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_CovLC2_31_Json_JsonPtr", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		_ = lc.Json()
		_ = lc.JsonPtr()
	})
}

func Test_CovLC2_32_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_CovLC2_32_ParseInjectUsingJson", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jr := lc.JsonPtr()
		lc2 := corestr.Empty.LinkedCollections()
		result, err := lc2.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
		_ = result
	})
}

func Test_CovLC2_33_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovLC2_33_ParseInjectUsingJsonMust", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jr := lc.JsonPtr()
		lc2 := corestr.Empty.LinkedCollections()
		_ = lc2.ParseInjectUsingJsonMust(jr)
	})
}

func Test_CovLC2_34_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_CovLC2_34_GetCompareSummary", func() {
		a := corestr.Empty.LinkedCollections()
		a.Add(corestr.New.Collection.Strings([]string{"x"}))
		b := corestr.Empty.LinkedCollections()
		b.Add(corestr.New.Collection.Strings([]string{"y"}))
		s := a.GetCompareSummary(b, "left", "right")
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_CovLC2_35_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovLC2_35_JsonParseSelfInject", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jr := lc.JsonPtr()
		lc2 := corestr.Empty.LinkedCollections()
		err := lc2.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovLC2_36_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovLC2_36_AsInterfaces", func() {
		lc := corestr.Empty.LinkedCollections()
		_ = lc.AsJsonContractsBinder()
		_ = lc.AsJsoner()
		_ = lc.AsJsonParseSelfInjector()
		_ = lc.AsJsonMarshaller()
	})
}

func Test_CovLC2_37_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_CovLC2_37_AddCollectionToNode", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"base"}))
		lc.AddCollectionToNode(true, lc.Head(), corestr.New.Collection.Strings([]string{"x"}))
		if lc.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_CovLC2_38_AttachWithNode(t *testing.T) {
	safeTest(t, "Test_CovLC2_38_AttachWithNode", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		// err: node nil
		err := lc.AttachWithNode(nil, &corestr.LinkedCollectionNode{})
		if err == nil {
			t.Fatal("expected error for nil node")
		}
		// node.next not nil -> error
		head := lc.Head()
		// head.next is nil, so this should succeed
		addingNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		err2 := lc.AttachWithNode(head, addingNode)
		if err2 != nil {
			t.Fatal("unexpected error:", err2)
		}
	})
}
