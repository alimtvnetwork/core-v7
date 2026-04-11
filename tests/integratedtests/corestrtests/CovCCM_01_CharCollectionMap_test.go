package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CharCollectionMap — Segment 14: Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovCCM_01_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_CovCCM_01_IsEmpty_HasItems", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()

		// Act
		actual := args.Map{"result": ccm.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ccm.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		ccm.Add("apple")
		actual = args.Map{"result": ccm.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual = args.Map{"result": ccm.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected items", actual)
	})
}

func Test_CovCCM_02_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_02_IsEmptyLock", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()

		// Act
		actual := args.Map{"result": ccm.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_CovCCM_03_GetChar(t *testing.T) {
	safeTest(t, "Test_CovCCM_03_GetChar", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		c := ccm.GetChar("hello")

		// Act
		actual := args.Map{"result": c != 'h'}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		c2 := ccm.GetChar("")
		actual = args.Map{"result": c2 != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovCCM_04_Add_AddLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_04_Add_AddLock", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		ccm.Add("ant")

		// Act
		actual := args.Map{"result": ccm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 char group", actual)
		ccm.Add("banana")
		actual = args.Map{"result": ccm.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 char groups", actual)
		// AddLock
		ccm.AddLock("cherry")
		actual = args.Map{"result": ccm.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// AddLock existing char
		ccm.AddLock("cat")
	})
}

func Test_CovCCM_05_AddStrings(t *testing.T) {
	safeTest(t, "Test_CovCCM_05_AddStrings", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.AddStrings("apple", "ant", "banana")

		// Act
		actual := args.Map{"result": ccm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ccm.AddStrings()
	})
}

func Test_CovCCM_06_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_CovCCM_06_AddSameStartingCharItems", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.AddSameStartingCharItems('a', []string{"apple", "ant"}, false)

		// Act
		actual := args.Map{"result": ccm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// add to existing
		ccm.AddSameStartingCharItems('a', []string{"axe"}, false)
		// empty
		ccm.AddSameStartingCharItems('b', nil, false)
	})
}

func Test_CovCCM_07_Length_LengthLock_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_CovCCM_07_Length_LengthLock_AllLengthsSum", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()

		// Act
		actual := args.Map{"result": ccm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": ccm.LengthLock() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": ccm.AllLengthsSum() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ccm.AddStrings("apple", "ant", "banana")
		actual = args.Map{"result": ccm.AllLengthsSum() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": ccm.AllLengthsSumLock() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovCCM_08_LengthOfCollectionFromFirstChar(t *testing.T) {
	safeTest(t, "Test_CovCCM_08_LengthOfCollectionFromFirstChar", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.AddStrings("apple", "ant")
		l := ccm.LengthOfCollectionFromFirstChar("a")

		// Act
		actual := args.Map{"result": l != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		l2 := ccm.LengthOfCollectionFromFirstChar("z")
		actual = args.Map{"result": l2 != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovCCM_09_LengthOf_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_09_LengthOf_LengthOfLock", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")

		// Act
		actual := args.Map{"result": ccm.LengthOf('a') != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": ccm.LengthOf('z') != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": ccm.LengthOfLock('a') != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		e := corestr.Empty.CharCollectionMap()
		actual = args.Map{"result": e.LengthOf('a') != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": e.LengthOfLock('a') != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovCCM_10_Has(t *testing.T) {
	safeTest(t, "Test_CovCCM_10_Has", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()

		// Act
		actual := args.Map{"result": ccm.Has("apple")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		ccm.Add("apple")
		actual = args.Map{"result": ccm.Has("apple")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ccm.Has("ant")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// missing char group
		actual = args.Map{"result": ccm.Has("banana")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovCCM_11_HasWithCollection(t *testing.T) {
	safeTest(t, "Test_CovCCM_11_HasWithCollection", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		// empty
		has, col := ccm.HasWithCollection("apple")

		// Act
		actual := args.Map{"result": has || col == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		ccm.Add("apple")
		has2, col2 := ccm.HasWithCollection("apple")
		actual = args.Map{"result": has2 || col2 == nil}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		// missing char
		has3, _ := ccm.HasWithCollection("banana")
		actual = args.Map{"result": has3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovCCM_12_HasWithCollectionLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_12_HasWithCollectionLock", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		has, _ := ccm.HasWithCollectionLock("x")

		// Act
		actual := args.Map{"result": has}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		ccm.Add("apple")
		has2, col := ccm.HasWithCollectionLock("apple")
		actual = args.Map{"result": has2 || col == nil}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		// missing char
		has3, _ := ccm.HasWithCollectionLock("banana")
		actual = args.Map{"result": has3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovCCM_13_IsEquals_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_13_IsEquals_IsEqualsLock", func() {
		// Arrange
		a := corestr.Empty.CharCollectionMap()
		a.AddStrings("apple", "banana")
		b := corestr.Empty.CharCollectionMap()
		b.AddStrings("apple", "banana")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual = args.Map{"result": a.IsEqualsLock(b)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		// nil
		actual = args.Map{"result": a.IsEquals(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// same
		actual = args.Map{"result": a.IsEquals(a)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// both empty
		e1 := corestr.Empty.CharCollectionMap()
		e2 := corestr.Empty.CharCollectionMap()
		actual = args.Map{"result": e1.IsEquals(e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// one empty
		actual = args.Map{"result": a.IsEquals(e1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// diff length
		c := corestr.Empty.CharCollectionMap()
		c.Add("apple")
		actual = args.Map{"result": a.IsEquals(c)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovCCM_14_IsEqualsCaseSensitive(t *testing.T) {
	safeTest(t, "Test_CovCCM_14_IsEqualsCaseSensitive", func() {
		// Arrange
		a := corestr.Empty.CharCollectionMap()
		a.Add("Apple")
		b := corestr.Empty.CharCollectionMap()
		b.Add("Apple")

		// Act
		actual := args.Map{"result": a.IsEqualsCaseSensitive(true, b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual = args.Map{"result": a.IsEqualsCaseSensitiveLock(true, b)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		// missing key in right
		c := corestr.Empty.CharCollectionMap()
		c.Add("Banana")
		// same length but different keys - need same number of char groups
		d := corestr.Empty.CharCollectionMap()
		d.Add("Axe")
		actual = args.Map{"result": a.IsEqualsCaseSensitive(true, d)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false - different content", actual)
	})
}

func Test_CovCCM_15_GetCollection(t *testing.T) {
	safeTest(t, "Test_CovCCM_15_GetCollection", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		col := ccm.GetCollection("a", false)

		// Act
		actual := args.Map{"result": col == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		// not found, no add
		col2 := ccm.GetCollection("z", false)
		actual = args.Map{"result": col2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		// not found, add new
		col3 := ccm.GetCollection("z", true)
		actual = args.Map{"result": col3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected new collection", actual)
	})
}

func Test_CovCCM_16_GetCollectionLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_16_GetCollectionLock", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		col := ccm.GetCollectionLock("a", false)

		// Act
		actual := args.Map{"result": col == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
	})
}

func Test_CovCCM_17_GetCollectionByChar(t *testing.T) {
	safeTest(t, "Test_CovCCM_17_GetCollectionByChar", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		col := ccm.GetCollectionByChar('a')

		// Act
		actual := args.Map{"result": col == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		col2 := ccm.GetCollectionByChar('z')
		actual = args.Map{"result": col2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_CovCCM_18_GetMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_18_GetMap_GetCopyMapLock", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		m := ccm.GetMap()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		m2 := ccm.GetCopyMapLock()
		actual = args.Map{"result": len(m2) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		e := corestr.Empty.CharCollectionMap()
		m3 := e.GetCopyMapLock()
		actual = args.Map{"result": len(m3) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovCCM_19_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_CovCCM_19_GetCharsGroups", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		r := ccm.GetCharsGroups([]string{"apple", "ant", "banana"})

		// Act
		actual := args.Map{"result": r.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		r2 := ccm.GetCharsGroups(nil)
		actual = args.Map{"result": r2 != ccm}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same", actual)
	})
}

func Test_CovCCM_20_List_ListLock_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_CovCCM_20_List_ListLock_SortedListAsc", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()

		// Act
		actual := args.Map{"result": len(ccm.List()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ccm.AddStrings("banana", "apple")
		list := ccm.List()
		actual = args.Map{"result": len(list) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ll := ccm.ListLock()
		actual = args.Map{"result": len(ll) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		sorted := ccm.SortedListAsc()
		actual = args.Map{"result": len(sorted) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty sorted
		e := corestr.Empty.CharCollectionMap()
		actual = args.Map{"result": len(e.SortedListAsc()) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovCCM_21_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_CovCCM_21_AddCollectionItems", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.AddCollectionItems(nil)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		ccm.AddCollectionItems(col)

		// Act
		actual := args.Map{"result": ccm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovCCM_22_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_CovCCM_22_AddHashmapsValues", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("k", "apple")
		ccm.AddHashmapsValues(hm, nil)

		// Act
		actual := args.Map{"result": ccm.AllLengthsSum() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ccm.AddHashmapsValues(nil)
	})
}

func Test_CovCCM_23_AddHashmapsKeysOrValuesBothUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovCCM_23_AddHashmapsKeysOrValuesBothUsingFilter", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("key", "val")
		filter := func(p corestr.KeyValuePair) (string, bool, bool) {
			return p.Value, true, false
		}
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(filter, hm, nil)

		// Act
		actual := args.Map{"result": ccm.AllLengthsSum() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// nil
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(filter, nil)
		// break
		breakFilter := func(p corestr.KeyValuePair) (string, bool, bool) {
			return p.Value, true, true
		}
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(breakFilter, hm)
	})
}

func Test_CovCCM_24_AddHashmapsKeysValuesBoth(t *testing.T) {
	safeTest(t, "Test_CovCCM_24_AddHashmapsKeysValuesBoth", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("key", "val")
		ccm.AddHashmapsKeysValuesBoth(hm)

		// Act
		actual := args.Map{"result": ccm.AllLengthsSum() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
		ccm.AddHashmapsKeysValuesBoth(nil)
	})
}

func Test_CovCCM_25_AddCharHashsetMap(t *testing.T) {
	safeTest(t, "Test_CovCCM_25_AddCharHashsetMap", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		ccm.AddCharHashsetMap(chm)

		// Act
		actual := args.Map{"result": ccm.AllLengthsSum() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		ccm.AddCharHashsetMap(corestr.Empty.CharHashsetMap())
	})
}

func Test_CovCCM_26_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_CovCCM_26_AddSameCharsCollection", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		col := corestr.New.Collection.Strings([]string{"apple", "ant"})
		r := ccm.AddSameCharsCollection("a", col)

		// Act
		actual := args.Map{"result": r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected collection", actual)
		// add more to existing
		col2 := corestr.New.Collection.Strings([]string{"axe"})
		r2 := ccm.AddSameCharsCollection("a", col2)
		actual = args.Map{"result": r2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected collection", actual)
		// nil collection, existing char
		r3 := ccm.AddSameCharsCollection("a", nil)
		actual = args.Map{"result": r3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected existing", actual)
		// nil collection, new char
		r4 := ccm.AddSameCharsCollection("z", nil)
		actual = args.Map{"result": r4 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
	})
}

func Test_CovCCM_27_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_27_AddSameCharsCollectionLock", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		col := corestr.New.Collection.Strings([]string{"apple"})
		r := ccm.AddSameCharsCollectionLock("a", col)

		// Act
		actual := args.Map{"result": r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected collection", actual)
		// existing, add more
		col2 := corestr.New.Collection.Strings([]string{"ant"})
		r2 := ccm.AddSameCharsCollectionLock("a", col2)
		actual = args.Map{"result": r2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected collection", actual)
		// existing, nil collection
		r3 := ccm.AddSameCharsCollectionLock("a", nil)
		actual = args.Map{"result": r3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected existing", actual)
		// new char, nil
		r4 := ccm.AddSameCharsCollectionLock("z", nil)
		actual = args.Map{"result": r4 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		// new char, with items
		col3 := corestr.New.Collection.Strings([]string{"banana"})
		r5 := ccm.AddSameCharsCollectionLock("b", col3)
		actual = args.Map{"result": r5 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected collection", actual)
	})
}

func Test_CovCCM_28_Resize_AddLength(t *testing.T) {
	safeTest(t, "Test_CovCCM_28_Resize_AddLength", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		ccm.Resize(10)
		ccm.Resize(0) // no-op, current >= new
		ccm.AddLength(5, 3)
		ccm.AddLength()
	})
}

func Test_CovCCM_29_HashsetByChar_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_29_HashsetByChar_HashsetByCharLock", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		hs := ccm.HashsetByChar('a')

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		hs2 := ccm.HashsetByChar('z')
		actual = args.Map{"result": hs2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		hs3 := ccm.HashsetByCharLock('a')
		actual = args.Map{"result": hs3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		// nil collection
		hs4 := ccm.HashsetByCharLock('z')
		actual = args.Map{"result": hs4 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty hashset", actual)
	})
}

func Test_CovCCM_30_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_CovCCM_30_HashsetByStringFirstChar", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		hs := ccm.HashsetByStringFirstChar("a")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		hs2 := ccm.HashsetByStringFirstCharLock("a")
		actual = args.Map{"result": hs2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_CovCCM_31_HashsetsCollectionByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_CovCCM_31_HashsetsCollectionByStringFirstChar", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.AddStrings("apple", "banana")
		hsc := ccm.HashsetsCollectionByStringFirstChar("apple", "banana", "cherry")

		// Act
		actual := args.Map{"result": hsc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		// empty
		e := corestr.Empty.CharCollectionMap()
		hsc2 := e.HashsetsCollectionByStringFirstChar("x")
		actual = args.Map{"result": hsc2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_CovCCM_32_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_CovCCM_32_HashsetsCollection", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		hsc := ccm.HashsetsCollection()

		// Act
		actual := args.Map{"result": hsc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		ccm.Add("apple")
		hsc2 := ccm.HashsetsCollection()
		actual = args.Map{"result": hsc2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_CovCCM_33_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_CovCCM_33_HashsetsCollectionByChars", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		hsc := ccm.HashsetsCollectionByChars('a', 'z')

		// Act
		actual := args.Map{"result": hsc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		// empty
		e := corestr.Empty.CharCollectionMap()
		hsc2 := e.HashsetsCollectionByChars('a')
		actual = args.Map{"result": hsc2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_CovCCM_34_String_StringLock_SummaryString(t *testing.T) {
	safeTest(t, "Test_CovCCM_34_String_StringLock_SummaryString", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		s := ccm.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		sl := ccm.StringLock()
		actual = args.Map{"result": sl == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		ss := ccm.SummaryString()
		actual = args.Map{"result": ss == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		ssl := ccm.SummaryStringLock()
		actual = args.Map{"result": ssl == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_CovCCM_35_Print_PrintLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_35_Print_PrintLock", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		ccm.Print(false) // skip
		ccm.Print(true)
		ccm.PrintLock(false) // skip
		ccm.PrintLock(true)
	})
}

func Test_CovCCM_36_JsonModel_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovCCM_36_JsonModel_JsonModelAny", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		_ = ccm.JsonModel()
		_ = ccm.JsonModelAny()
	})
}

func Test_CovCCM_37_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovCCM_37_MarshalJSON_UnmarshalJSON", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		data, err := ccm.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		ccm2 := corestr.Empty.CharCollectionMap()
		err2 := ccm2.UnmarshalJSON(data)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		// invalid
		err3 := ccm2.UnmarshalJSON([]byte("bad"))
		actual = args.Map{"result": err3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_CovCCM_38_Json_JsonPtr_ParseInject(t *testing.T) {
	safeTest(t, "Test_CovCCM_38_Json_JsonPtr_ParseInject", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		_ = ccm.Json()
		jr := ccm.JsonPtr()
		ccm2 := corestr.Empty.CharCollectionMap()
		r, err := ccm2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual = args.Map{"result": r.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_CovCCM_39_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovCCM_39_ParseInjectUsingJsonMust", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		jr := ccm.JsonPtr()
		ccm2 := corestr.Empty.CharCollectionMap()
		r := ccm2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"result": r.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_CovCCM_40_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovCCM_40_JsonParseSelfInject", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		jr := ccm.JsonPtr()
		ccm2 := corestr.Empty.CharCollectionMap()
		err := ccm2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovCCM_41_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovCCM_41_AsInterfaces", func() {
		ccm := corestr.Empty.CharCollectionMap()
		_ = ccm.AsJsonContractsBinder()
		_ = ccm.AsJsoner()
		_ = ccm.AsJsonMarshaller()
		_ = ccm.AsJsonParseSelfInjector()
	})
}

func Test_CovCCM_42_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovCCM_42_Clear_Dispose", func() {
		// Arrange
		ccm := corestr.Empty.CharCollectionMap()
		ccm.AddStrings("apple", "banana")
		ccm.Clear()

		// Act
		actual := args.Map{"result": ccm.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		// clear already empty
		ccm.Clear()
		// dispose
		ccm2 := corestr.Empty.CharCollectionMap()
		ccm2.Add("x")
		ccm2.Dispose()
	})
}
