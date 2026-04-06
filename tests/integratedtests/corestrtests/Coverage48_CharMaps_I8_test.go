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
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()

		// Act
		actual := args.Map{"result": m.IsEmpty() || m.HasItems() || m.Length() != 0}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C48_CharCollectionMap_Add(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Add", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("apple")
		m.Add("apricot")
		m.Add("banana")

		// Act
		actual := args.Map{"result": m.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 char groups", actual)
	})
}

func Test_C48_CharCollectionMap_AddLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddLock", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		m.AddLock("abc")

		// Act
		actual := args.Map{"result": m.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddStrings", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		m.AddStrings("apple", "banana", "avocado")

		// Act
		actual := args.Map{"result": m.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharCollectionMap_Has(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Has", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})

		// Act
		actual := args.Map{"result": m.Has("apple")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have apple", actual)
		actual := args.Map{"result": m.Has("cherry")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have cherry", actual)
	})
}

func Test_C48_CharCollectionMap_Has_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Has_Empty", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()

		// Act
		actual := args.Map{"result": m.Has("x")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty map should not have anything", actual)
	})
}

func Test_C48_CharCollectionMap_HasWithCollection(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HasWithCollection", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		has, col := m.HasWithCollection("apple")

		// Act
		actual := args.Map{"result": has || col == nil}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		has, _ := m.HasWithCollection("x")

		// Act
		actual := args.Map{"result": has}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty should not have", actual)
	})
}

func Test_C48_CharCollectionMap_HasWithCollectionLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HasWithCollectionLock", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		has, _ := m.HasWithCollectionLock("apple")

		// Act
		actual := args.Map{"result": has}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have apple", actual)
	})
}

func Test_C48_CharCollectionMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_LengthOf", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "apricot"})
		l := m.LengthOf('a')

		// Act
		actual := args.Map{"result": l != 2}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		l := m.LengthOfLock('a')

		// Act
		actual := args.Map{"result": l != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_LengthOfCollectionFromFirstChar(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_LengthOfCollectionFromFirstChar", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "apricot"})
		l := m.LengthOfCollectionFromFirstChar("any")

		// Act
		actual := args.Map{"result": l != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharCollectionMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AllLengthsSum", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "apricot", "banana"})
		sum := m.AllLengthsSum()

		// Act
		actual := args.Map{"result": sum != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C48_CharCollectionMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AllLengthsSumLock", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		sum := m.AllLengthsSumLock()

		// Act
		actual := args.Map{"result": sum != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_GetChar(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetChar", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		c := m.GetChar("hello")

		// Act
		actual := args.Map{"result": c != 'h'}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := m.GetCollection("any", false)

		// Act
		actual := args.Map{"result": col == nil}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := m.GetCollectionLock("apple", false)

		// Act
		actual := args.Map{"result": col == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected collection", actual)
	})
}

func Test_C48_CharCollectionMap_GetCollectionByChar(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetCollectionByChar", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := m.GetCollectionByChar('a')

		// Act
		actual := args.Map{"result": col == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected collection for 'a'", actual)
	})
}

func Test_C48_CharCollectionMap_IsEquals(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEquals", func() {
		// Arrange
		m1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m2 := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": m1.IsEquals(m2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C48_CharCollectionMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEquals_Nil", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": m.IsEquals(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil should not be equal", actual)
	})
}

func Test_C48_CharCollectionMap_IsEquals_SameRef(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEquals_SameRef", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": m.IsEquals(m)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same ref should be equal", actual)
	})
}

func Test_C48_CharCollectionMap_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEquals_DiffLength", func() {
		// Arrange
		m1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m2 := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})

		// Act
		actual := args.Map{"result": m1.IsEquals(m2)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "diff length should not be equal", actual)
	})
}

func Test_C48_CharCollectionMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEqualsLock", func() {
		// Arrange
		m1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m2 := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": m1.IsEqualsLock(m2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C48_CharCollectionMap_IsEqualsCaseSensitive(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEqualsCaseSensitive", func() {
		// Arrange
		m1 := corestr.New.CharCollectionMap.Items([]string{"Apple"})
		m2 := corestr.New.CharCollectionMap.Items([]string{"Apple"})

		// Act
		actual := args.Map{"result": m1.IsEqualsCaseSensitive(true, m2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C48_CharCollectionMap_IsEqualsCaseSensitiveLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_IsEqualsCaseSensitiveLock", func() {
		// Arrange
		m1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m2 := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": m1.IsEqualsCaseSensitiveLock(true, m2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C48_CharCollectionMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddSameStartingCharItems", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		m.AddSameStartingCharItems('a', []string{"apple", "avocado"}, false)

		// Act
		actual := args.Map{"result": m.LengthOf('a') != 2}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "apple")
		hm.AddOrUpdate("k2", "banana")
		m.AddHashmapsValues(hm)

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 2}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("apple", "avocado")
		m.AddHashmapsKeysValuesBoth(hm)

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharCollectionMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddCollectionItems", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		m.AddCollectionItems(col)

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 2}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		m.AddCharHashsetMap(chm)

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharCollectionMap_Resize(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Resize", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"a1"})
		m.Resize(100)

		// Act
		actual := args.Map{"result": m.Length() != 1}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		list := m.List()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharCollectionMap_ListLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_ListLock", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		list := m.ListLock()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_SortedListAsc", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"banana", "apple"})
		list := m.SortedListAsc()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual := args.Map{"result": list[0] != "apple"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected apple first", actual)
	})
}

func Test_C48_CharCollectionMap_SortedListAsc_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_SortedListAsc_Empty", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		list := m.SortedListAsc()

		// Act
		actual := args.Map{"result": len(list) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharCollectionMap_String(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_String", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := m.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_SummaryString", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := m.SummaryString()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_StringLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_StringLock", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := m.StringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_SummaryStringLock", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := m.SummaryStringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()

		// Act
		actual := args.Map{"result": m.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C48_CharCollectionMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_LengthLock", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": m.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetByChar", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := m.HashsetByChar('a')

		// Act
		actual := args.Map{"result": hs == nil || hs.IsEmpty()}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := m.HashsetByCharLock('a')

		// Act
		actual := args.Map{"result": hs == nil || hs.IsEmpty()}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := m.HashsetByStringFirstChar("anything")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset for 'a'", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetByStringFirstCharLock", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := m.HashsetByStringFirstCharLock("anything")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetsCollection", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := m.HashsetsCollection()

		// Act
		actual := args.Map{"result": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetsCollection_Empty", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		hsc := m.HashsetsCollection()

		// Act
		actual := args.Map{"result": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetsCollectionByChars", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := m.HashsetsCollectionByChars('a')

		// Act
		actual := args.Map{"result": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_HashsetsCollectionByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_HashsetsCollectionByStringFirstChar", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hsc := m.HashsetsCollectionByStringFirstChar("anything")

		// Act
		actual := args.Map{"result": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharCollectionMap_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_AddSameCharsCollection", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		col := corestr.New.Collection.Strings([]string{"apple", "avocado"})
		result := m.AddSameCharsCollection("abc", col)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		result := m.GetCharsGroups([]string{"apple", "banana"})

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected result", actual)
	})
}

func Test_C48_CharCollectionMap_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetCharsGroups_Empty", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		result := m.GetCharsGroups([]string{})

		// Act
		actual := args.Map{"result": result != m}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected self", actual)
	})
}

func Test_C48_CharCollectionMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetCopyMapLock", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		cm := m.GetCopyMapLock()

		// Act
		actual := args.Map{"result": len(cm) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetCopyMapLock_Empty", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		cm := m.GetCopyMapLock()

		// Act
		actual := args.Map{"result": len(cm) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharCollectionMap_GetMap(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_GetMap", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		gm := m.GetMap()

		// Act
		actual := args.Map{"result": len(gm) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMap_Clear(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Clear", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m.Clear()

		// Act
		actual := args.Map{"result": m.IsEmpty()}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		data, err := json.Marshal(m)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty json", actual)
	})
}

func Test_C48_CharCollectionMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_UnmarshalJSON", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		data, _ := json.Marshal(m)
		m2 := corestr.New.CharCollectionMap.Empty()
		err := json.Unmarshal(data, m2)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C48_CharCollectionMap_Json(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_Json", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		j := m.Json()

		// Act
		actual := args.Map{"result": j.Error}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "j.Error", actual)
	})
}

func Test_C48_CharCollectionMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_JsonPtr", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		j := m.JsonPtr()

		// Act
		actual := args.Map{"result": j == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharCollectionMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_JsonModel", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		jm := m.JsonModel()

		// Act
		actual := args.Map{"result": jm == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharCollectionMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_JsonModelAny", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		a := m.JsonModelAny()

		// Act
		actual := args.Map{"result": a == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharCollectionMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMap_ParseInjectUsingJson", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		j := m.Json()
		m2 := corestr.New.CharCollectionMap.Empty()
		_, err := m2.ParseInjectUsingJson(&j)

		// Act
		actual := args.Map{"result": err}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		j := m.Json()
		m2 := corestr.New.CharCollectionMap.Empty()
		err := m2.JsonParseSelfInject(&j)

		// Act
		actual := args.Map{"result": err}

		// Assert
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
		// Arrange
		m := corestr.New.CharCollectionMap.CapSelfCap(20, 15)

		// Act
		actual := args.Map{"result": m == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharCollectionMapCreator_ItemsPtrWithCap(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMapCreator_ItemsPtrWithCap", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 10, []string{"apple"})

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharCollectionMapCreator_ItemsPtrWithCap_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMapCreator_ItemsPtrWithCap_Empty", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 10, []string{})

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharCollectionMapCreator_Items_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionMapCreator_Items_Empty", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{})

		// Act
		actual := args.Map{"result": m.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// DataModel
func Test_C48_CharCollectionDataModel(t *testing.T) {
	safeTest(t, "Test_C48_CharCollectionDataModel", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple"})
		dm := corestr.NewCharCollectionMapDataModelUsing(m)
		m2 := corestr.NewCharCollectionMapUsingDataModel(dm)

		// Act
		actual := args.Map{"result": m2.AllLengthsSum() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// EmptyCreator
func Test_C48_EmptyCreator_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_C48_EmptyCreator_CharCollectionMap", func() {
		// Arrange
		m := corestr.Empty.CharCollectionMap()

		// Act
		actual := args.Map{"result": m == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ===================== CharHashsetMap =====================

func Test_C48_CharHashsetMap_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Empty", func() {
		// Arrange
		m := corestr.Empty.CharHashsetMap()

		// Act
		actual := args.Map{"result": m == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharHashsetMap_Cap(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Cap", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(20, 15)

		// Act
		actual := args.Map{"result": m == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharHashsetMap_CapItems(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_CapItems", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(20, 15, "apple", "banana")

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_Strings(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Strings", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Strings(10, []string{"apple", "banana"})

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_Strings_Nil(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Strings_Nil", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Strings(10, nil)

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharHashsetMap_Add(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Add", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.Add("apple")
		m.Add("apricot")
		m.Add("banana")

		// Act
		actual := args.Map{"result": m.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_AddLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddLock", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddLock("abc")

		// Act
		actual := args.Map{"result": m.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddStrings", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddStrings("apple", "banana")

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 2}

		// Assert
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
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddStringsLock("apple")

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 1}

		// Assert
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
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": m.Has("apple")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have apple", actual)
		actual := args.Map{"result": m.Has("cherry")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have cherry", actual)
	})
}

func Test_C48_CharHashsetMap_Has_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Has_Empty", func() {
		// Arrange
		m := corestr.Empty.CharHashsetMap()

		// Act
		actual := args.Map{"result": m.Has("x")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty should not have", actual)
	})
}

func Test_C48_CharHashsetMap_HasWithHashset(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HasWithHashset", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		has, hs := m.HasWithHashset("apple")

		// Act
		actual := args.Map{"result": has || hs == nil}

		// Assert
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
		// Arrange
		m := corestr.Empty.CharHashsetMap()
		has, _ := m.HasWithHashset("x")

		// Act
		actual := args.Map{"result": has}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty should not have", actual)
	})
}

func Test_C48_CharHashsetMap_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HasWithHashsetLock", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		has, _ := m.HasWithHashsetLock("apple")

		// Act
		actual := args.Map{"result": has}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have", actual)
	})
}

func Test_C48_CharHashsetMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_LengthOf", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": m.LengthOf('a') != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": m.LengthOf('z') != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharHashsetMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_LengthOfLock", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": m.LengthOfLock('a') != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_LengthOfHashsetFromFirstChar", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": m.LengthOfHashsetFromFirstChar("any") != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AllLengthsSum", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AllLengthsSumLock", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": m.AllLengthsSumLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_GetChar(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetChar", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"result": m.GetChar("hello") != 'h'}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		actual := args.Map{"result": m.GetChar("") != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharHashsetMap_GetCharOf(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetCharOf", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"result": m.GetCharOf("hello") != 'h'}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		actual := args.Map{"result": m.GetCharOf("") != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharHashsetMap_GetHashset(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetHashset", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.GetHashset("any", false)

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
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
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.GetHashsetLock(false, "apple")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_C48_CharHashsetMap_GetHashsetByChar(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetHashsetByChar", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.GetHashsetByChar('a')

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetByChar", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.HashsetByChar('a')

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetByCharLock", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.HashsetByCharLock('a')

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
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
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.HashsetByStringFirstChar("anything")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetByStringFirstCharLock", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := m.HashsetByStringFirstCharLock("anything")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_C48_CharHashsetMap_IsEquals(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_IsEquals", func() {
		// Arrange
		m1 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		m2 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": m1.IsEquals(m2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C48_CharHashsetMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_IsEquals_Nil", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": m.IsEquals(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil should not be equal", actual)
	})
}

func Test_C48_CharHashsetMap_IsEquals_SameRef(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_IsEquals_SameRef", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": m.IsEquals(m)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same ref", actual)
	})
}

func Test_C48_CharHashsetMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_IsEqualsLock", func() {
		// Arrange
		m1 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		m2 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": m1.IsEqualsLock(m2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C48_CharHashsetMap_List(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_List", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		list := m.List()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_SortedListAsc", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "banana", "apple")
		list := m.SortedListAsc()

		// Act
		actual := args.Map{"result": len(list) != 2 || list[0] != "apple"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected sorted", actual)
	})
}

func Test_C48_CharHashsetMap_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_SortedListDsc", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		list := m.SortedListDsc()

		// Act
		actual := args.Map{"result": len(list) != 2 || list[0] != "banana"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected descending", actual)
	})
}

func Test_C48_CharHashsetMap_String(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_String", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := m.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharHashsetMap_StringLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_StringLock", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := m.StringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
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
		// Arrange
		m := corestr.Empty.CharHashsetMap()

		// Act
		actual := args.Map{"result": m.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C48_CharHashsetMap_HasItems(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HasItems", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": m.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetsCollection", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		hsc := m.HashsetsCollection()

		// Act
		actual := args.Map{"result": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetsCollection_Empty", func() {
		// Arrange
		m := corestr.Empty.CharHashsetMap()
		hsc := m.HashsetsCollection()

		// Act
		actual := args.Map{"result": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetsCollectionByChars", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hsc := m.HashsetsCollectionByChars('a')

		// Act
		actual := args.Map{"result": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharHashsetMap_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_HashsetsCollectionByStringsFirstChar", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hsc := m.HashsetsCollectionByStringsFirstChar("anything")

		// Act
		actual := args.Map{"result": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharHashsetMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetCharsGroups", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		result := m.GetCharsGroups("apple", "banana")

		// Act
		actual := args.Map{"result": result.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetCharsGroups_Empty", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		result := m.GetCharsGroups()

		// Act
		actual := args.Map{"result": result != m}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected self", actual)
	})
}

func Test_C48_CharHashsetMap_GetMap(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetMap", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		gm := m.GetMap()

		// Act
		actual := args.Map{"result": len(gm) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetCopyMapLock", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		cm := m.GetCopyMapLock()

		// Act
		actual := args.Map{"result": len(cm) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C48_CharHashsetMap_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_GetCopyMapLock_Empty", func() {
		// Arrange
		m := corestr.Empty.CharHashsetMap()
		cm := m.GetCopyMapLock()

		// Act
		actual := args.Map{"result": len(cm) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C48_CharHashsetMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddSameStartingCharItems", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		m.AddSameStartingCharItems('a', []string{"apple"})

		// Act
		actual := args.Map{"result": m.LengthOf('a') != 1}

		// Assert
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
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		m.AddCollectionItems(col)

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 2}

		// Assert
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
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m.AddCharCollectionMapItems(ccm)

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 1}

		// Assert
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
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple", "banana"})
		m.AddHashsetItems(hs)

		// Act
		actual := args.Map{"result": m.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C48_CharHashsetMap_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_AddSameCharsCollection", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple"})
		result := m.AddSameCharsCollection("abc", col)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
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
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		m.RemoveAll()

		// Act
		actual := args.Map{"result": m.IsEmpty()}

		// Assert
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
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		m.Clear()

		// Act
		actual := args.Map{"result": m.IsEmpty()}

		// Assert
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
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		data, err := json.Marshal(m)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C48_CharHashsetMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_UnmarshalJSON", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		data, _ := json.Marshal(m)
		m2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := json.Unmarshal(data, m2)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C48_CharHashsetMap_Json(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_Json", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		j := m.Json()

		// Act
		actual := args.Map{"result": j.Error}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "j.Error", actual)
	})
}

func Test_C48_CharHashsetMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_JsonPtr", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		j := m.JsonPtr()

		// Act
		actual := args.Map{"result": j == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharHashsetMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_JsonModel", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		jm := m.JsonModel()

		// Act
		actual := args.Map{"result": jm == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharHashsetMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_JsonModelAny", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		a := m.JsonModelAny()

		// Act
		actual := args.Map{"result": a == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C48_CharHashsetMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C48_CharHashsetMap_ParseInjectUsingJson", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		j := m.Json()
		m2 := corestr.New.CharHashsetMap.Cap(10, 10)
		_, err := m2.ParseInjectUsingJson(&j)

		// Act
		actual := args.Map{"result": err}

		// Assert
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
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		j := m.Json()
		m2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := m2.JsonParseSelfInject(&j)

		// Act
		actual := args.Map{"result": err}

		// Assert
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
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		dm := corestr.NewCharHashsetMapDataModelUsing(m)
		m2 := corestr.NewCharHashsetMapUsingDataModel(dm)

		// Act
		actual := args.Map{"result": m2.AllLengthsSum() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
