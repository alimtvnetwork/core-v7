package corestrtests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===================== CharCollectionMap =====================

func Test_C48_CharCollectionMap_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		actual := args.Map{"result": m.IsEmpty() || m.HasItems() || m.Length() != 0}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C48_CharCollectionMap_Add(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Add", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("apple")
		m.Add("apricot")
		m.Add("banana")
		actual := args.Map{"result": m.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 char groups", actual)
	})
}

func Test_C48_CharCollectionMap_AddLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddLock", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddLock("abc")
		actual := args.Map{"result": m.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddStrings", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddStrings("apple", "banana", "avocado")
		actual := args.Map{"result": m.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharCollectionMap_Has(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Has", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		actual := args.Map{"result": m.Has("apple")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have apple", actual)
		actual := args.Map{"result": m.Has("cherry")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have cherry", actual)
	})
}

func Test_C48_CharCollectionMap_Has_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Has_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		actual := args.Map{"result": m.Has("x")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty map should not have anything", actual)
	})
}

func Test_C48_CharCollectionMap_HasWithCollection(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HasWithCollection", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		has, col := m.HasWithCollection("apple")
		actual := args.Map{"result": has || col == nil}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have apple with collection", actual)
		has2, _ := m.HasWithCollection("missing")
		actual := args.Map{"result": has2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not find missing", actual)
	})
}

func Test_C48_CharCollectionMap_HasWithCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HasWithCollection_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		has, _ := m.HasWithCollection("x")
		actual := args.Map{"result": has}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty should not have", actual)
	})
}

func Test_C48_CharCollectionMap_HasWithCollectionLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HasWithCollectionLock", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		has, _ := m.HasWithCollectionLock("apple")
		actual := args.Map{"result": has}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have apple", actual)
	})
}

func Test_C48_CharCollectionMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_LengthOf", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "apricot"})
		l := m.LengthOf('a')
		actual := args.Map{"result": l != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		l2 := m.LengthOf('z')
		actual := args.Map{"result": l2 != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharCollectionMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_LengthOfLock", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		l := m.LengthOfLock('a')
		actual := args.Map{"result": l != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_LengthOfCollectionFromFirstChar(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_LengthOfCollectionFromFirstChar", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "apricot"})
		l := m.LengthOfCollectionFromFirstChar("any")
		actual := args.Map{"result": l != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharCollectionMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AllLengthsSum", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "apricot", "banana"})
		sum := m.AllLengthsSum()
		actual := args.Map{"result": sum != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C48_CharCollectionMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AllLengthsSumLock", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		sum := m.AllLengthsSumLock()
		actual := args.Map{"result": sum != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_GetChar(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetChar", func() {
		m := corestr.New.CharCollectionMap.Empty()
		c := m.GetChar("hello")
		actual := args.Map{"result": c != 'h'}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		c2 := m.GetChar("")
		actual := args.Map{"result": c2 != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharCollectionMap_GetCollection(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetCollection", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := m.GetCollection("any", false)
		actual := args.Map{"result": col == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected collection for 'a'", actual)
		col2 := m.GetCollection("zzz", false)
		actual := args.Map{"result": col2 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil for 'z'", actual)
		col3 := m.GetCollection("zzz", true)
		actual := args.Map{"result": col3 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected new collection for 'z' with addNew=true", actual)
	})
}

func Test_C48_CharCollectionMap_GetCollectionLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetCollectionLock", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := m.GetCollectionLock("apple", false)
		actual := args.Map{"result": col == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected collection", actual)
	})
}

func Test_C48_CharCollectionMap_GetCollectionByChar(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetCollectionByChar", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := m.GetCollectionByChar('a')
		actual := args.Map{"result": col == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected collection for 'a'", actual)
	})
}

func Test_C48_CharCollectionMap_IsEquals(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEquals", func() {
		m1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m2 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": m1.IsEquals(m2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C48_CharCollectionMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEquals_Nil", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": m.IsEquals(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil should not be equal", actual)
	})
}

func Test_C48_CharCollectionMap_IsEquals_SameRef(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEquals_SameRef", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": m.IsEquals(m)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same ref should be equal", actual)
	})
}

func Test_C48_CharCollectionMap_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEquals_DiffLength", func() {
		m1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m2 := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		actual := args.Map{"result": m1.IsEquals(m2)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "diff length should not be equal", actual)
	})
}

func Test_C48_CharCollectionMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEqualsLock", func() {
		m1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m2 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": m1.IsEqualsLock(m2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C48_CharCollectionMap_IsEqualsCaseSensitive(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEqualsCaseSensitive", func() {
		m1 := corestr.New.CharCollectionMap.Items([]string{"Apple"})
		m2 := corestr.New.CharCollectionMap.Items([]string{"Apple"})
		actual := args.Map{"result": m1.IsEqualsCaseSensitive(true, m2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C48_CharCollectionMap_IsEqualsCaseSensitiveLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEqualsCaseSensitiveLock", func() {
		m1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m2 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": m1.IsEqualsCaseSensitiveLock(true, m2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C48_CharCollectionMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddSameStartingCharItems", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddSameStartingCharItems('a', []string{"apple", "avocado"}, false)
		actual := args.Map{"result": m.LengthOf('a') != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		m.AddSameStartingCharItems('a', []string{"apricot"}, false)
		actual := args.Map{"result": m.LengthOf('a') != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C48_CharCollectionMap_AddSameStartingCharItems_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddSameStartingCharItems_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddSameStartingCharItems('a', []string{}, false)
	})
}

func Test_C48_CharCollectionMap_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddHashmapsValues", func() {
		m := corestr.New.CharCollectionMap.Empty()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "apple")
		hm.AddOrUpdate("k2", "banana")
		m.AddHashmapsValues(hm)
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharCollectionMap_AddHashmapsValues_Nil(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddHashmapsValues_Nil", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddHashmapsValues(nil)
	})
}

func Test_C48_CharCollectionMap_AddHashmapsKeysValuesBoth(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddHashmapsKeysValuesBoth", func() {
		m := corestr.New.CharCollectionMap.Empty()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("apple", "avocado")
		m.AddHashmapsKeysValuesBoth(hm)
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharCollectionMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddCollectionItems", func() {
		m := corestr.New.CharCollectionMap.Empty()
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		m.AddCollectionItems(col)
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharCollectionMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddCollectionItems_Nil", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddCollectionItems(nil)
	})
}

func Test_C48_CharCollectionMap_AddCharHashsetMap(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddCharHashsetMap", func() {
		m := corestr.New.CharCollectionMap.Empty()
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		m.AddCharHashsetMap(chm)
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharCollectionMap_Resize(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Resize", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"a1"})
		m.Resize(100)
		actual := args.Map{"result": m.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "resize should preserve items", actual)
		// resize with smaller - no change
		m.Resize(0)
	})
}

func Test_C48_CharCollectionMap_AddLength(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddLength", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"a1"})
		m.AddLength(10)
	})
}

func Test_C48_CharCollectionMap_AddLength_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddLength_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddLength()
	})
}

func Test_C48_CharCollectionMap_List(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_List", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		list := m.List()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharCollectionMap_ListLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_ListLock", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		list := m.ListLock()
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_SortedListAsc", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"banana", "apple"})
		list := m.SortedListAsc()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual := args.Map{"result": list[0] != "apple"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected apple first", actual)
	})
}

func Test_C48_CharCollectionMap_SortedListAsc_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_SortedListAsc_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		list := m.SortedListAsc()
		actual := args.Map{"result": len(list) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharCollectionMap_String(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_String", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := m.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_SummaryString", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := m.SummaryString()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_StringLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_StringLock", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := m.StringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_SummaryStringLock", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := m.SummaryStringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_Print(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Print", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m.Print(false) // skip
		m.Print(true)
	})
}

func Test_C48_CharCollectionMap_PrintLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_PrintLock", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m.PrintLock(false)
		m.PrintLock(true)
	})
}

func Test_C48_CharCollectionMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEmptyLock", func() {
		m := corestr.New.CharCollectionMap.Empty()
		actual := args.Map{"result": m.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C48_CharCollectionMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_LengthLock", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": m.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetByChar", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := m.HashsetByChar('a')
		actual := args.Map{"result": hs == nil || hs.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset with items", actual)
		hs2 := m.HashsetByChar('z')
		actual := args.Map{"result": hs2 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil for missing char", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetByCharLock", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := m.HashsetByCharLock('a')
		actual := args.Map{"result": hs == nil || hs.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		hs2 := m.HashsetByCharLock('z')
		actual := args.Map{"result": hs2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty hashset, not nil", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetByStringFirstChar", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := m.HashsetByStringFirstChar("anything")
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset for 'a'", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetByStringFirstCharLock", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := m.HashsetByStringFirstCharLock("anything")
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetsCollection", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := m.HashsetsCollection()
		actual := args.Map{"result": hsc.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetsCollection_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		hsc := m.HashsetsCollection()
		actual := args.Map{"result": hsc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetsCollectionByChars", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := m.HashsetsCollectionByChars('a')
		actual := args.Map{"result": hsc.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetsCollectionByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetsCollectionByStringFirstChar", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hsc := m.HashsetsCollectionByStringFirstChar("anything")
		actual := args.Map{"result": hsc.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddSameCharsCollection", func() {
		m := corestr.New.CharCollectionMap.Empty()
		col := corestr.New.Collection.Strings([]string{"apple", "avocado"})
		result := m.AddSameCharsCollection("abc", col)
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected collection", actual)
		// Add to existing
		col2 := corestr.New.Collection.Strings([]string{"apricot"})
		m.AddSameCharsCollection("abc", col2)
		// Nil collection
		m.AddSameCharsCollection("xyz", nil)
	})
}

func Test_C48_CharCollectionMap_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddSameCharsCollectionLock", func() {
		m := corestr.New.CharCollectionMap.Empty()
		col := corestr.New.Collection.Strings([]string{"apple"})
		m.AddSameCharsCollectionLock("abc", col)
		// nil
		m.AddSameCharsCollectionLock("xyz", nil)
	})
}

func Test_C48_CharCollectionMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetCharsGroups", func() {
		m := corestr.New.CharCollectionMap.Empty()
		result := m.GetCharsGroups([]string{"apple", "banana"})
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected result", actual)
	})
}

func Test_C48_CharCollectionMap_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetCharsGroups_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		result := m.GetCharsGroups([]string{})
		actual := args.Map{"result": result != m}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected self", actual)
	})
}

func Test_C48_CharCollectionMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetCopyMapLock", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		cm := m.GetCopyMapLock()
		actual := args.Map{"result": len(cm) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetCopyMapLock_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		cm := m.GetCopyMapLock()
		actual := args.Map{"result": len(cm) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharCollectionMap_GetMap(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetMap", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		gm := m.GetMap()
		actual := args.Map{"result": len(gm) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_Clear(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Clear", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m.Clear()
		actual := args.Map{"result": m.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty after clear", actual)
	})
}

func Test_C48_CharCollectionMap_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Clear_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Clear()
	})
}

func Test_C48_CharCollectionMap_Dispose(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Dispose", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m.Dispose()
	})
}

func Test_C48_CharCollectionMap_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Dispose_Nil", func() {
		var m *corestr.CharCollectionMap
		m.Dispose()
	})
}

// JSON
func Test_C48_CharCollectionMap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_MarshalJSON", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		data, err := json.Marshal(m)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty json", actual)
	})
}

func Test_C48_CharCollectionMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_UnmarshalJSON", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		data, _ := json.Marshal(m)
		m2 := corestr.New.CharCollectionMap.Empty()
		err := json.Unmarshal(data, m2)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C48_CharCollectionMap_Json(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Json", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		j := m.Json()
		actual := args.Map{"result": j.Error}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "j.Error", actual)
	})
}

func Test_C48_CharCollectionMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_JsonPtr", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		j := m.JsonPtr()
		actual := args.Map{"result": j == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharCollectionMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_JsonModel", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		jm := m.JsonModel()
		actual := args.Map{"result": jm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharCollectionMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_JsonModelAny", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		a := m.JsonModelAny()
		actual := args.Map{"result": a == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharCollectionMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_ParseInjectUsingJson", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		j := m.Json()
		m2 := corestr.New.CharCollectionMap.Empty()
		_, err := m2.ParseInjectUsingJson(&j)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C48_CharCollectionMap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_ParseInjectUsingJsonMust", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		j := m.Json()
		m2 := corestr.New.CharCollectionMap.Empty()
		m2.ParseInjectUsingJsonMust(&j)
	})
}

func Test_C48_CharCollectionMap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_JsonParseSelfInject", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		j := m.Json()
		m2 := corestr.New.CharCollectionMap.Empty()
		err := m2.JsonParseSelfInject(&j)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C48_CharCollectionMap_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AsInterfaces", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		_ = m.AsJsonContractsBinder()
		_ = m.AsJsoner()
		_ = m.AsJsonMarshaller()
		_ = m.AsJsonParseSelfInjector()
	})
}

// Creator tests
func Test_C48_CharCollectionMapCreator_CapSelfCap(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMapCreator_CapSelfCap", func() {
		m := corestr.New.CharCollectionMap.CapSelfCap(20, 15)
		actual := args.Map{"result": m == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharCollectionMapCreator_ItemsPtrWithCap(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMapCreator_ItemsPtrWithCap", func() {
		m := corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 10, []string{"apple"})
		actual := args.Map{"result": m.AllLengthsSum() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMapCreator_ItemsPtrWithCap_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMapCreator_ItemsPtrWithCap_Empty", func() {
		m := corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 10, []string{})
		actual := args.Map{"result": m.AllLengthsSum() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharCollectionMapCreator_Items_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMapCreator_Items_Empty", func() {
		m := corestr.New.CharCollectionMap.Items([]string{})
		actual := args.Map{"result": m.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// DataModel
func Test_C48_CharCollectionDataModel(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionDataModel", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		dm := corestr.NewCharCollectionMapDataModelUsing(m)
		m2 := corestr.NewCharCollectionMapUsingDataModel(dm)
		actual := args.Map{"result": m2.AllLengthsSum() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// EmptyCreator
func Test_C48_EmptyCreator_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_C48_EmptyCreator_CharCollectionMap", func() {
		m := corestr.Empty.CharCollectionMap()
		actual := args.Map{"result": m == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ===================== CharHashsetMap =====================

func Test_C48_CharHashsetMap_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Empty", func() {
		m := corestr.Empty.CharHashsetMap()
		actual := args.Map{"result": m == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharHashsetMap_Cap(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Cap", func() {
		m := corestr.New.CharHashsetMap.Cap(20, 15)
		actual := args.Map{"result": m == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharHashsetMap_CapItems(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_CapItems", func() {
		m := corestr.New.CharHashsetMap.CapItems(20, 15, "apple", "banana")
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_Strings(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Strings", func() {
		m := corestr.New.CharHashsetMap.Strings(10, []string{"apple", "banana"})
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_Strings_Nil(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Strings_Nil", func() {
		m := corestr.New.CharHashsetMap.Strings(10, nil)
		actual := args.Map{"result": m.AllLengthsSum() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharHashsetMap_Add(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Add", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.Add("apple")
		m.Add("apricot")
		m.Add("banana")
		actual := args.Map{"result": m.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_AddLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddLock", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddLock("abc")
		actual := args.Map{"result": m.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddStrings", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddStrings("apple", "banana")
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_AddStrings_Nil(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddStrings_Nil", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddStrings()
	})
}

func Test_C48_CharHashsetMap_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddStringsLock", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddStringsLock("apple")
		actual := args.Map{"result": m.AllLengthsSum() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddStringsLock_Empty", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddStringsLock()
	})
}

func Test_C48_CharHashsetMap_Has(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Has", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": m.Has("apple")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have apple", actual)
		actual := args.Map{"result": m.Has("cherry")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have cherry", actual)
	})
}

func Test_C48_CharHashsetMap_Has_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Has_Empty", func() {
		m := corestr.Empty.CharHashsetMap()
		actual := args.Map{"result": m.Has("x")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty should not have", actual)
	})
}

func Test_C48_CharHashsetMap_HasWithHashset(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HasWithHashset", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		has, hs := m.HasWithHashset("apple")
		actual := args.Map{"result": has || hs == nil}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have", actual)
		has2, _ := m.HasWithHashset("missing")
		actual := args.Map{"result": has2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have", actual)
	})
}

func Test_C48_CharHashsetMap_HasWithHashset_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HasWithHashset_Empty", func() {
		m := corestr.Empty.CharHashsetMap()
		has, _ := m.HasWithHashset("x")
		actual := args.Map{"result": has}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty should not have", actual)
	})
}

func Test_C48_CharHashsetMap_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HasWithHashsetLock", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		has, _ := m.HasWithHashsetLock("apple")
		actual := args.Map{"result": has}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have", actual)
	})
}

func Test_C48_CharHashsetMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_LengthOf", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": m.LengthOf('a') != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": m.LengthOf('z') != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharHashsetMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_LengthOfLock", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": m.LengthOfLock('a') != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_LengthOfHashsetFromFirstChar", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": m.LengthOfHashsetFromFirstChar("any") != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AllLengthsSum", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AllLengthsSumLock", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": m.AllLengthsSumLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_GetChar(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetChar", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"result": m.GetChar("hello") != 'h'}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		actual := args.Map{"result": m.GetChar("") != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharHashsetMap_GetCharOf(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetCharOf", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"result": m.GetCharOf("hello") != 'h'}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		actual := args.Map{"result": m.GetCharOf("") != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharHashsetMap_GetHashset(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetHashset", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.GetHashset("any", false)
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset for 'a'", actual)
		hs2 := m.GetHashset("zzz", false)
		actual := args.Map{"result": hs2 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		hs3 := m.GetHashset("zzz", true)
		actual := args.Map{"result": hs3 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected new hashset", actual)
	})
}

func Test_C48_CharHashsetMap_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetHashsetLock", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.GetHashsetLock(false, "apple")
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_C48_CharHashsetMap_GetHashsetByChar(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetHashsetByChar", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.GetHashsetByChar('a')
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetByChar", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.HashsetByChar('a')
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetByCharLock", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.HashsetByCharLock('a')
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		hs2 := m.HashsetByCharLock('z')
		actual := args.Map{"result": hs2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty hashset", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetByStringFirstChar", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.HashsetByStringFirstChar("anything")
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetByStringFirstCharLock", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.HashsetByStringFirstCharLock("anything")
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_C48_CharHashsetMap_IsEquals(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_IsEquals", func() {
		m1 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		m2 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": m1.IsEquals(m2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C48_CharHashsetMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_IsEquals_Nil", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": m.IsEquals(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil should not be equal", actual)
	})
}

func Test_C48_CharHashsetMap_IsEquals_SameRef(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_IsEquals_SameRef", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": m.IsEquals(m)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same ref", actual)
	})
}

func Test_C48_CharHashsetMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_IsEqualsLock", func() {
		m1 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		m2 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": m1.IsEqualsLock(m2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C48_CharHashsetMap_List(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_List", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		list := m.List()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_SortedListAsc", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "banana", "apple")
		list := m.SortedListAsc()
		actual := args.Map{"result": len(list) != 2 || list[0] != "apple"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected sorted", actual)
	})
}

func Test_C48_CharHashsetMap_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_SortedListDsc", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		list := m.SortedListDsc()
		actual := args.Map{"result": len(list) != 2 || list[0] != "banana"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected descending", actual)
	})
}

func Test_C48_CharHashsetMap_String(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_String", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := m.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharHashsetMap_StringLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_StringLock", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := m.StringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharHashsetMap_Print(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Print", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		m.Print(false)
		m.Print(true)
	})
}

func Test_C48_CharHashsetMap_PrintLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_PrintLock", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		m.PrintLock(false)
		m.PrintLock(true)
	})
}

func Test_C48_CharHashsetMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_IsEmptyLock", func() {
		m := corestr.Empty.CharHashsetMap()
		actual := args.Map{"result": m.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C48_CharHashsetMap_HasItems(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HasItems", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": m.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetsCollection", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		hsc := m.HashsetsCollection()
		actual := args.Map{"result": hsc.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetsCollection_Empty", func() {
		m := corestr.Empty.CharHashsetMap()
		hsc := m.HashsetsCollection()
		actual := args.Map{"result": hsc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetsCollectionByChars", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hsc := m.HashsetsCollectionByChars('a')
		actual := args.Map{"result": hsc.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetsCollectionByStringsFirstChar", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hsc := m.HashsetsCollectionByStringsFirstChar("anything")
		actual := args.Map{"result": hsc.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharHashsetMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetCharsGroups", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		result := m.GetCharsGroups("apple", "banana")
		actual := args.Map{"result": result.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetCharsGroups_Empty", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		result := m.GetCharsGroups()
		actual := args.Map{"result": result != m}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected self", actual)
	})
}

func Test_C48_CharHashsetMap_GetMap(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetMap", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		gm := m.GetMap()
		actual := args.Map{"result": len(gm) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetCopyMapLock", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		cm := m.GetCopyMapLock()
		actual := args.Map{"result": len(cm) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetCopyMapLock_Empty", func() {
		m := corestr.Empty.CharHashsetMap()
		cm := m.GetCopyMapLock()
		actual := args.Map{"result": len(cm) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharHashsetMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddSameStartingCharItems", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddSameStartingCharItems('a', []string{"apple"})
		actual := args.Map{"result": m.LengthOf('a') != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		m.AddSameStartingCharItems('a', []string{"avocado"})
		actual := args.Map{"result": m.LengthOf('a') != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_AddSameStartingCharItems_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddSameStartingCharItems_Empty", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddSameStartingCharItems('a', []string{})
	})
}

func Test_C48_CharHashsetMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddCollectionItems", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		m.AddCollectionItems(col)
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddCollectionItems_Nil", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddCollectionItems(nil)
	})
}

func Test_C48_CharHashsetMap_AddCharCollectionMapItems(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddCharCollectionMapItems", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m.AddCharCollectionMapItems(ccm)
		actual := args.Map{"result": m.AllLengthsSum() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_AddCharCollectionMapItems_Nil(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddCharCollectionMapItems_Nil", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddCharCollectionMapItems(nil)
	})
}

func Test_C48_CharHashsetMap_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddHashsetItems", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple", "banana"})
		m.AddHashsetItems(hs)
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddSameCharsCollection", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple"})
		result := m.AddSameCharsCollection("abc", col)
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		// Add to existing
		col2 := corestr.New.Collection.Strings([]string{"avocado"})
		m.AddSameCharsCollection("abc", col2)
		// Nil collection
		m.AddSameCharsCollection("xyz", nil)
	})
}

func Test_C48_CharHashsetMap_AddSameCharsHashset(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddSameCharsHashset", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		m.AddSameCharsHashset("abc", hs)
		// Add to existing
		hs2 := corestr.New.Hashset.Strings([]string{"avocado"})
		m.AddSameCharsHashset("abc", hs2)
		// Nil
		m.AddSameCharsHashset("xyz", nil)
	})
}

func Test_C48_CharHashsetMap_AddHashsetLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddHashsetLock", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		m.AddHashsetLock("abc", hs)
		// Nil
		m.AddHashsetLock("xyz", nil)
	})
}

func Test_C48_CharHashsetMap_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddSameCharsCollectionLock", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple"})
		m.AddSameCharsCollectionLock("abc", col)
		// nil
		m.AddSameCharsCollectionLock("xyz", nil)
	})
}

func Test_C48_CharHashsetMap_RemoveAll(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_RemoveAll", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		m.RemoveAll()
		actual := args.Map{"result": m.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C48_CharHashsetMap_RemoveAll_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_RemoveAll_Empty", func() {
		m := corestr.Empty.CharHashsetMap()
		m.RemoveAll()
	})
}

func Test_C48_CharHashsetMap_Clear(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Clear", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		m.Clear()
		actual := args.Map{"result": m.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C48_CharHashsetMap_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Clear_Empty", func() {
		m := corestr.Empty.CharHashsetMap()
		m.Clear()
	})
}

// JSON
func Test_C48_CharHashsetMap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_MarshalJSON", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		data, err := json.Marshal(m)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharHashsetMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_UnmarshalJSON", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		data, _ := json.Marshal(m)
		m2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := json.Unmarshal(data, m2)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C48_CharHashsetMap_Json(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Json", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		j := m.Json()
		actual := args.Map{"result": j.Error}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "j.Error", actual)
	})
}

func Test_C48_CharHashsetMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_JsonPtr", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		j := m.JsonPtr()
		actual := args.Map{"result": j == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharHashsetMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_JsonModel", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		jm := m.JsonModel()
		actual := args.Map{"result": jm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharHashsetMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_JsonModelAny", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		a := m.JsonModelAny()
		actual := args.Map{"result": a == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharHashsetMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_ParseInjectUsingJson", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		j := m.Json()
		m2 := corestr.New.CharHashsetMap.Cap(10, 10)
		_, err := m2.ParseInjectUsingJson(&j)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C48_CharHashsetMap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_ParseInjectUsingJsonMust", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		j := m.Json()
		m2 := corestr.New.CharHashsetMap.Cap(10, 10)
		m2.ParseInjectUsingJsonMust(&j)
	})
}

func Test_C48_CharHashsetMap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_JsonParseSelfInject", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		j := m.Json()
		m2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := m2.JsonParseSelfInject(&j)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C48_CharHashsetMap_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AsInterfaces", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		_ = m.AsJsonContractsBinder()
		_ = m.AsJsoner()
		_ = m.AsJsonMarshaller()
		_ = m.AsJsonParseSelfInjector()
	})
}

// DataModel
func Test_C48_CharHashsetDataModel(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetDataModel", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		dm := corestr.NewCharHashsetMapDataModelUsing(m)
		m2 := corestr.NewCharHashsetMapUsingDataModel(dm)
		actual := args.Map{"result": m2.AllLengthsSum() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
