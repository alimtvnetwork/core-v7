package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CharHashsetMap — Segment 15: Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovCHM_01_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_CovCHM_01_IsEmpty_HasItems", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()

		// Act
		actual := args.Map{"result": chm.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": chm.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		chm.Add("apple")
		actual = args.Map{"result": chm.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual = args.Map{"result": chm.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected items", actual)
	})
}

func Test_CovCHM_02_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_02_IsEmptyLock", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()

		// Act
		actual := args.Map{"result": chm.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_CovCHM_03_GetChar_GetCharOf(t *testing.T) {
	safeTest(t, "Test_CovCHM_03_GetChar_GetCharOf", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()

		// Act
		actual := args.Map{"result": chm.GetChar("hello") != 'h'}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		actual = args.Map{"result": chm.GetChar("") != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": chm.GetCharOf("hello") != 'h'}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		actual = args.Map{"result": chm.GetCharOf("") != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovCHM_04_Add_AddLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_04_Add_AddLock", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		chm.Add("ant") // same char group

		// Act
		actual := args.Map{"result": chm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		chm.Add("banana")
		actual = args.Map{"result": chm.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// AddLock
		chm.AddLock("cherry")
		actual = args.Map{"result": chm.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		chm.AddLock("cat") // existing
	})
}

func Test_CovCHM_05_AddStrings_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_05_AddStrings_AddStringsLock", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.AddStrings("apple", "banana")

		// Act
		actual := args.Map{"result": chm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		chm.AddStrings()
		chm.AddStringsLock("cherry", "cat")
		chm.AddStringsLock()
	})
}

func Test_CovCHM_06_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_CovCHM_06_AddSameStartingCharItems", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.AddSameStartingCharItems('a', []string{"apple", "ant"})

		// Act
		actual := args.Map{"result": chm.LengthOf('a') != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// existing
		chm.AddSameStartingCharItems('a', []string{"axe"})
		// empty
		chm.AddSameStartingCharItems('b', nil)
	})
}

func Test_CovCHM_07_Length_LengthLock_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_CovCHM_07_Length_LengthLock_AllLengthsSum", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": chm.LengthLock() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": chm.AllLengthsSum() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		chm.AddStrings("apple", "ant", "banana")
		actual = args.Map{"result": chm.AllLengthsSum() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": chm.AllLengthsSumLock() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovCHM_08_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_CovCHM_08_LengthOfHashsetFromFirstChar", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.AddStrings("apple", "ant")

		// Act
		actual := args.Map{"result": chm.LengthOfHashsetFromFirstChar("a") != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": chm.LengthOfHashsetFromFirstChar("z") != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovCHM_09_LengthOf_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_09_LengthOf_LengthOfLock", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.LengthOf('a') != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": chm.LengthOf('z') != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": chm.LengthOfLock('a') != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		e := corestr.Empty.CharHashsetMap()
		actual = args.Map{"result": e.LengthOf('a') != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": e.LengthOfLock('a') != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovCHM_10_Has(t *testing.T) {
	safeTest(t, "Test_CovCHM_10_Has", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()

		// Act
		actual := args.Map{"result": chm.Has("apple")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		chm.Add("apple")
		actual = args.Map{"result": chm.Has("apple")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": chm.Has("banana")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovCHM_11_HasWithHashset_Lock(t *testing.T) {
	safeTest(t, "Test_CovCHM_11_HasWithHashset_Lock", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		has, _ := chm.HasWithHashset("x")

		// Act
		actual := args.Map{"result": has}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		chm.Add("apple")
		has2, hs := chm.HasWithHashset("apple")
		actual = args.Map{"result": has2 || hs == nil}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		// missing char
		has3, _ := chm.HasWithHashset("banana")
		actual = args.Map{"result": has3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// Lock variant
		has4, _ := chm.HasWithHashsetLock("apple")
		actual = args.Map{"result": has4}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		has5, _ := chm.HasWithHashsetLock("banana")
		actual = args.Map{"result": has5}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// empty lock
		e := corestr.Empty.CharHashsetMap()
		has6, _ := e.HasWithHashsetLock("x")
		actual = args.Map{"result": has6}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovCHM_12_IsEquals_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_12_IsEquals_IsEqualsLock", func() {
		// Arrange
		a := corestr.Empty.CharHashsetMap()
		a.AddStrings("apple", "banana")
		b := corestr.Empty.CharHashsetMap()
		b.AddStrings("apple", "banana")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual = args.Map{"result": a.IsEqualsLock(b)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual = args.Map{"result": a.IsEquals(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": a.IsEquals(a)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		e1 := corestr.Empty.CharHashsetMap()
		e2 := corestr.Empty.CharHashsetMap()
		actual = args.Map{"result": e1.IsEquals(e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": a.IsEquals(e1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// diff length
		c := corestr.Empty.CharHashsetMap()
		c.Add("apple")
		actual = args.Map{"result": a.IsEquals(c)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// same length, diff content
		d := corestr.Empty.CharHashsetMap()
		d.AddStrings("axe", "cherry")
		actual = args.Map{"result": a.IsEquals(d)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovCHM_13_GetMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_13_GetMap_GetCopyMapLock", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")

		// Act
		actual := args.Map{"result": len(chm.GetMap()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(chm.GetCopyMapLock()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		e := corestr.Empty.CharHashsetMap()
		actual = args.Map{"result": len(e.GetCopyMapLock()) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovCHM_14_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_CovCHM_14_GetCharsGroups", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		r := chm.GetCharsGroups("apple", "ant", "banana")

		// Act
		actual := args.Map{"result": r.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		r2 := chm.GetCharsGroups()
		actual = args.Map{"result": r2 != chm}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same", actual)
	})
}

func Test_CovCHM_15_GetHashset_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_15_GetHashset_GetHashsetLock", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		hs := chm.GetHashset("a", false)

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		hs2 := chm.GetHashset("z", false)
		actual = args.Map{"result": hs2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		hs3 := chm.GetHashset("z", true)
		actual = args.Map{"result": hs3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		hs4 := chm.GetHashsetLock(false, "a")
		actual = args.Map{"result": hs4 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
	})
}

func Test_CovCHM_16_GetHashsetByChar_HashsetByChar_Lock(t *testing.T) {
	safeTest(t, "Test_CovCHM_16_GetHashsetByChar_HashsetByChar_Lock", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.GetHashsetByChar('a') == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		actual = args.Map{"result": chm.HashsetByChar('a') == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		hl := chm.HashsetByCharLock('a')
		actual = args.Map{"result": hl == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		// missing
		hl2 := chm.HashsetByCharLock('z')
		actual = args.Map{"result": hl2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty hashset", actual)
	})
}

func Test_CovCHM_17_HashsetByStringFirstChar_Lock(t *testing.T) {
	safeTest(t, "Test_CovCHM_17_HashsetByStringFirstChar_Lock", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		hs := chm.HashsetByStringFirstChar("a")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		hs2 := chm.HashsetByStringFirstCharLock("a")
		actual = args.Map{"result": hs2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
	})
}

func Test_CovCHM_18_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_CovCHM_18_AddCollectionItems", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.AddCollectionItems(nil)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		chm.AddCollectionItems(col)

		// Act
		actual := args.Map{"result": chm.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovCHM_19_AddCharCollectionMapItems(t *testing.T) {
	safeTest(t, "Test_CovCHM_19_AddCharCollectionMapItems", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		chm.AddCharCollectionMapItems(ccm)

		// Act
		actual := args.Map{"result": chm.AllLengthsSum() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		chm.AddCharCollectionMapItems(nil)
	})
}

func Test_CovCHM_20_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_CovCHM_20_AddHashsetItems", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		chm.AddHashsetItems(hs)

		// Act
		actual := args.Map{"result": chm.AllLengthsSum() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		chm.AddHashsetItems(corestr.New.Hashset.Empty())
	})
}

func Test_CovCHM_21_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_CovCHM_21_AddSameCharsCollection", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		col := corestr.New.Collection.Strings([]string{"apple", "ant"})
		r := chm.AddSameCharsCollection("a", col)

		// Act
		actual := args.Map{"result": r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		// existing, add more
		col2 := corestr.New.Collection.Strings([]string{"axe"})
		r2 := chm.AddSameCharsCollection("a", col2)
		actual = args.Map{"result": r2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		// existing, nil
		r3 := chm.AddSameCharsCollection("a", nil)
		actual = args.Map{"result": r3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected existing", actual)
		// new, nil
		r4 := chm.AddSameCharsCollection("z", nil)
		actual = args.Map{"result": r4 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		// new, with items
		col3 := corestr.New.Collection.Strings([]string{"banana"})
		r5 := chm.AddSameCharsCollection("b", col3)
		actual = args.Map{"result": r5 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_CovCHM_22_AddSameCharsHashset(t *testing.T) {
	safeTest(t, "Test_CovCHM_22_AddSameCharsHashset", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		r := chm.AddSameCharsHashset("a", hs)

		// Act
		actual := args.Map{"result": r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		// existing, add
		hs2 := corestr.New.Hashset.Strings([]string{"ant"})
		r2 := chm.AddSameCharsHashset("a", hs2)
		actual = args.Map{"result": r2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		// existing, nil
		r3 := chm.AddSameCharsHashset("a", nil)
		actual = args.Map{"result": r3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected existing", actual)
		// new, nil
		r4 := chm.AddSameCharsHashset("z", nil)
		actual = args.Map{"result": r4 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		// new, with items
		hs3 := corestr.New.Hashset.Strings([]string{"banana"})
		r5 := chm.AddSameCharsHashset("b", hs3)
		actual = args.Map{"result": r5 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_CovCHM_23_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_23_AddSameCharsCollectionLock", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		col := corestr.New.Collection.Strings([]string{"apple"})
		r := chm.AddSameCharsCollectionLock("a", col)

		// Act
		actual := args.Map{"result": r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		// existing, add
		col2 := corestr.New.Collection.Strings([]string{"ant"})
		r2 := chm.AddSameCharsCollectionLock("a", col2)
		actual = args.Map{"result": r2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		// existing, nil
		r3 := chm.AddSameCharsCollectionLock("a", nil)
		actual = args.Map{"result": r3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected existing", actual)
		// new, nil
		r4 := chm.AddSameCharsCollectionLock("z", nil)
		actual = args.Map{"result": r4 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		// new, with items
		col3 := corestr.New.Collection.Strings([]string{"banana"})
		r5 := chm.AddSameCharsCollectionLock("b", col3)
		actual = args.Map{"result": r5 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_CovCHM_24_AddHashsetLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_24_AddHashsetLock", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		r := chm.AddHashsetLock("a", hs)

		// Act
		actual := args.Map{"result": r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		// existing, add
		hs2 := corestr.New.Hashset.Strings([]string{"ant"})
		r2 := chm.AddHashsetLock("a", hs2)
		actual = args.Map{"result": r2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
		// existing, nil
		r3 := chm.AddHashsetLock("a", nil)
		actual = args.Map{"result": r3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected existing", actual)
		// new, nil
		r4 := chm.AddHashsetLock("z", nil)
		actual = args.Map{"result": r4 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		// new, with items
		hs3 := corestr.New.Hashset.Strings([]string{"banana"})
		r5 := chm.AddHashsetLock("b", hs3)
		actual = args.Map{"result": r5 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_CovCHM_25_List_SortedListAsc_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_CovCHM_25_List_SortedListAsc_SortedListDsc", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.AddStrings("banana", "apple")
		list := chm.List()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		asc := chm.SortedListAsc()
		actual = args.Map{"result": len(asc) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		dsc := chm.SortedListDsc()
		actual = args.Map{"result": len(dsc) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovCHM_26_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_CovCHM_26_HashsetsCollectionByChars", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.AddStrings("apple", "banana")
		hsc := chm.HashsetsCollectionByChars('a', 'z')

		// Act
		actual := args.Map{"result": hsc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		// empty
		e := corestr.Empty.CharHashsetMap()
		hsc2 := e.HashsetsCollectionByChars('a')
		actual = args.Map{"result": hsc2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_CovCHM_27_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_CovCHM_27_HashsetsCollectionByStringsFirstChar", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.AddStrings("apple", "banana")
		hsc := chm.HashsetsCollectionByStringsFirstChar("apple", "cherry")

		// Act
		actual := args.Map{"result": hsc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		e := corestr.Empty.CharHashsetMap()
		hsc2 := e.HashsetsCollectionByStringsFirstChar("x")
		actual = args.Map{"result": hsc2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_CovCHM_28_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_CovCHM_28_HashsetsCollection", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		hsc := chm.HashsetsCollection()

		// Act
		actual := args.Map{"result": hsc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		chm.Add("apple")
		hsc2 := chm.HashsetsCollection()
		actual = args.Map{"result": hsc2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_CovCHM_29_String_StringLock_SummaryString(t *testing.T) {
	safeTest(t, "Test_CovCHM_29_String_StringLock_SummaryString", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual = args.Map{"result": chm.StringLock() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual = args.Map{"result": chm.SummaryString() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual = args.Map{"result": chm.SummaryStringLock() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_CovCHM_30_Print_PrintLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_30_Print_PrintLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		chm.Print(false)
		chm.Print(true)
		chm.PrintLock(false)
		chm.PrintLock(true)
	})
}

func Test_CovCHM_31_JsonModel_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovCHM_31_JsonModel_JsonModelAny", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		_ = chm.JsonModel()
		_ = chm.JsonModelAny()
	})
}

func Test_CovCHM_32_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovCHM_32_MarshalJSON_UnmarshalJSON", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		data, err := chm.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		chm2 := corestr.Empty.CharHashsetMap()
		err2 := chm2.UnmarshalJSON(data)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		err3 := chm2.UnmarshalJSON([]byte("bad"))
		actual = args.Map{"result": err3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_CovCHM_33_Json_JsonPtr_ParseInject(t *testing.T) {
	safeTest(t, "Test_CovCHM_33_Json_JsonPtr_ParseInject", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		_ = chm.Json()
		jr := chm.JsonPtr()
		chm2 := corestr.Empty.CharHashsetMap()
		r, err := chm2.ParseInjectUsingJson(jr)

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

func Test_CovCHM_34_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovCHM_34_ParseInjectUsingJsonMust", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.Empty.CharHashsetMap()
		r := chm2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"result": r.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_CovCHM_35_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovCHM_35_JsonParseSelfInject", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.Empty.CharHashsetMap()
		err := chm2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovCHM_36_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovCHM_36_AsInterfaces", func() {
		chm := corestr.Empty.CharHashsetMap()
		_ = chm.AsJsonContractsBinder()
		_ = chm.AsJsoner()
		_ = chm.AsJsonMarshaller()
		_ = chm.AsJsonParseSelfInjector()
	})
}

func Test_CovCHM_37_Clear_RemoveAll(t *testing.T) {
	safeTest(t, "Test_CovCHM_37_Clear_RemoveAll", func() {
		// Arrange
		chm := corestr.Empty.CharHashsetMap()
		chm.AddStrings("apple", "banana")
		chm.Clear()

		// Act
		actual := args.Map{"result": chm.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		chm.Clear() // already empty
		chm.Add("x")
		chm.RemoveAll()
		actual = args.Map{"result": chm.IsEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		// RemoveAll on empty
		e := corestr.Empty.CharHashsetMap()
		e.RemoveAll()
	})
}

func Test_CovCHM_38_AddCollectionItemsAsyncLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_38_AddCollectionItemsAsyncLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		// nil
		chm.AddCollectionItemsAsyncLock(nil, nil)
		// empty
		chm.AddCollectionItemsAsyncLock(corestr.Empty.Collection(), nil)
	})
}

func Test_CovCHM_39_AddHashsetItemsAsyncLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_39_AddHashsetItemsAsyncLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.AddHashsetItemsAsyncLock(nil, nil)
		chm.AddHashsetItemsAsyncLock(corestr.New.Hashset.Empty(), nil)
	})
}
