package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
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
		actual := args.Map{"result": lc.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_CovLC2_02_AddCollectionsToNode(t *testing.T) {
	safeTest(t, "Test_CovLC2_02_AddCollectionsToNode", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"base"}))
		lc.AddCollectionsToNode(true, lc.Head(), corestr.New.Collection.Strings([]string{"x"}))
		actual := args.Map{"result": lc.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
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
		actual := args.Map{"result": lc.Length() < 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
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
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLC2_06_ConcatNew(t *testing.T) {
	safeTest(t, "Test_CovLC2_06_ConcatNew", func() {
		a := corestr.Empty.LinkedCollections()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))

		// empty with clone
		cloned := a.ConcatNew(true)
		actual := args.Map{"result": cloned.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty without clone
		same := a.ConcatNew(false)
		actual := args.Map{"result": same.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// with others
		b := corestr.Empty.LinkedCollections()
		b.Add(corestr.New.Collection.Strings([]string{"b"}))
		merged := a.ConcatNew(false, b)
		actual := args.Map{"result": merged.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		lc2 := corestr.Empty.LinkedCollections()
		lc2.AddAsyncFuncItemsPointer(wg, false)
	})
}

func Test_CovLC2_09_AddStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_CovLC2_09_AddStringsOfStrings", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.AddStringsOfStrings(false, []string{"a"}, nil, []string{"b"})
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": n == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		// index 1
		n1 := lc.IndexAt(1)
		actual := args.Map{"result": n1 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		// index 2
		n2 := lc.IndexAt(2)
		actual := args.Map{"result": n2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		// negative
		nn := lc.IndexAt(-1)
		actual := args.Map{"result": nn != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil for negative", actual)
	})
}

func Test_CovLC2_11_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_CovLC2_11_SafePointerIndexAt", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		col := lc.SafePointerIndexAt(0)
		actual := args.Map{"result": col == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		nilCol := lc.SafePointerIndexAt(99)
		actual := args.Map{"result": nilCol != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_CovLC2_12_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_CovLC2_12_SafeIndexAt", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		n := lc.SafeIndexAt(0)
		actual := args.Map{"result": n == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		n1 := lc.SafeIndexAt(1)
		actual := args.Map{"result": n1 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		// out of range
		actual := args.Map{"result": lc.SafeIndexAt(-1) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		actual := args.Map{"result": lc.SafeIndexAt(99) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_CovLC2_13_AddStringsAsync(t *testing.T) {
	safeTest(t, "Test_CovLC2_13_AddStringsAsync", func() {
		lc := corestr.Empty.LinkedCollections()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddStringsAsync(wg, []string{"a", "b"})
		wg.Wait()
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// nil
		lc.AddStringsAsync(wg, nil)
	})
}

func Test_CovLC2_14_AddCollection(t *testing.T) {
	safeTest(t, "Test_CovLC2_14_AddCollection", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		lc.AddCollection(nil)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 1", actual)
	})
}

func Test_CovLC2_15_AddCollectionsPtr_AddCollections(t *testing.T) {
	safeTest(t, "Test_CovLC2_15_AddCollectionsPtr_AddCollections", func() {
		lc := corestr.Empty.LinkedCollections()
		cols := []*corestr.Collection{corestr.New.Collection.Strings([]string{"a"})}
		lc.AddCollectionsPtr(cols)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		actual := args.Map{"result": len(*ptr) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		strs := lc.ToStrings()
		actual := args.Map{"result": len(strs) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLC2_17_ToCollectionSimple_ToCollection(t *testing.T) {
	safeTest(t, "Test_CovLC2_17_ToCollectionSimple_ToCollection", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		col := lc.ToCollectionSimple()
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		col2 := lc.ToCollection(5)
		actual := args.Map{"result": col2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		e := corestr.Empty.LinkedCollections()
		actual := args.Map{"result": e.ToCollection(0).Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLC2_18_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_CovLC2_18_ToCollectionsOfCollection", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		coc := lc.ToCollectionsOfCollection(0)
		actual := args.Map{"result": coc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		e := corestr.Empty.LinkedCollections()
		actual := args.Map{"result": e.ToCollectionsOfCollection(0).Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLC2_19_ItemsOfItems(t *testing.T) {
	safeTest(t, "Test_CovLC2_19_ItemsOfItems", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b", "c"}))
		ii := lc.ItemsOfItems()
		actual := args.Map{"result": len(ii) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		e := corestr.Empty.LinkedCollections()
		actual := args.Map{"result": len(e.ItemsOfItems()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLC2_20_ItemsOfItemsCollection(t *testing.T) {
	safeTest(t, "Test_CovLC2_20_ItemsOfItemsCollection", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		cols := lc.ItemsOfItemsCollection()
		actual := args.Map{"result": len(cols) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		e := corestr.Empty.LinkedCollections()
		actual := args.Map{"result": len(e.ItemsOfItemsCollection()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLC2_21_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_CovLC2_21_SimpleSlice", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		ss := lc.SimpleSlice()
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLC2_22_ListPtr_List(t *testing.T) {
	safeTest(t, "Test_CovLC2_22_ListPtr_List", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		ptr := lc.ListPtr()
		actual := args.Map{"result": len(*ptr) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		list := lc.List()
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		e := corestr.Empty.LinkedCollections()
		actual := args.Map{"result": len(e.List()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLC2_23_String(t *testing.T) {
	safeTest(t, "Test_CovLC2_23_String", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		s := lc.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
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
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		e := corestr.Empty.LinkedCollections()
		_ = e.StringLock()
	})
}

func Test_CovLC2_25_Join(t *testing.T) {
	safeTest(t, "Test_CovLC2_25_Join", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		s := lc.Join(",")
		actual := args.Map{"result": s != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a', got ''", actual)
	})
}

func Test_CovLC2_26_Joins(t *testing.T) {
	safeTest(t, "Test_CovLC2_26_Joins", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		// with items
		s := lc.Joins(",", "b")
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
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
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_ = lc.JsonModelAny()
	})
}

func Test_CovLC2_28_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovLC2_28_MarshalJSON", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		data, err := lc.MarshalJSON()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected data", actual)
	})
}

func Test_CovLC2_29_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovLC2_29_UnmarshalJSON", func() {
		lc := corestr.Empty.LinkedCollections()
		err := lc.UnmarshalJSON([]byte(`["a","b"]`))
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 collection with 2 items", actual)
		// invalid
		err2 := lc.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"result": err2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_CovLC2_30_RemoveAll_Clear(t *testing.T) {
	safeTest(t, "Test_CovLC2_30_RemoveAll_Clear", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.RemoveAll()
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
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
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_CovLC2_35_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovLC2_35_JsonParseSelfInject", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jr := lc.JsonPtr()
		lc2 := corestr.Empty.LinkedCollections()
		err := lc2.JsonParseSelfInject(jr)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
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
		actual := args.Map{"result": lc.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_CovLC2_38_AttachWithNode(t *testing.T) {
	safeTest(t, "Test_CovLC2_38_AttachWithNode", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		// err: node nil
		err := lc.AttachWithNode(nil, &corestr.LinkedCollectionNode{})
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for nil node", actual)
		// node.next not nil -> error
		head := lc.Head()
		// head.next is nil, so this should succeed
		addingNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		err2 := lc.AttachWithNode(head, addingNode)
		actual := args.Map{"result": err2}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	})
}
