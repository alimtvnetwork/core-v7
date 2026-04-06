package corestrtests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =======================================================
// CharCollectionMap
// =======================================================

func Test_C26_CharCollectionMap_Empty(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Empty()
		if !ccm.IsEmpty() {
			// Empty() creates an empty map — IsEmpty() should return true

		// Act
			actual := args.Map{"result": false}

		// Assert
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "expected empty internal state", actual)
		}
		actual := args.Map{"result": ccm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharCollectionMap_CapSelfCap(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_CapSelfCap", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(20, 15)

		// Act
		actual := args.Map{"result": ccm == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_Items(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Items", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana", "avocado"})

		// Act
		actual := args.Map{"length": ccm.Length()}

		// Assert
		expected := args.Map{"length": 2}
		expected.ShouldBeEqual(t, 0, "CharCollectionMap.Items returns 2 -- a and b groups", actual)
	})
}

func Test_C26_CharCollectionMap_Items_Empty(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Items_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items(nil)

		// Act
		actual := args.Map{"result": ccm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharCollectionMap_ItemsPtrWithCap(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ItemsPtrWithCap", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 10, []string{"abc", "aef"})

		// Act
		actual := args.Map{"result": ccm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharCollectionMap_ItemsPtrWithCap_Empty(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ItemsPtrWithCap_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 10, nil)

		// Act
		actual := args.Map{"result": ccm == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_GetChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Empty()
		ch := ccm.GetChar("hello")

		// Act
		actual := args.Map{"result": ch != 'h'}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		ch = ccm.GetChar("")
		actual := args.Map{"result": ch != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharCollectionMap_Add(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Add", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.Add("apple")
		ccm.Add("avocado")
		ccm.Add("banana")

		// Act
		actual := args.Map{"result": ccm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharCollectionMap_AddLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.AddLock("apple")

		// Act
		actual := args.Map{"result": ccm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharCollectionMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddStrings", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.AddStrings("apple", "banana", "cherry")

		// Act
		actual := args.Map{"result": ccm.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C26_CharCollectionMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameStartingCharItems", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.AddSameStartingCharItems('a', []string{"apple", "avocado"}, false)

		// Act
		actual := args.Map{"result": ccm.LengthOf('a') != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// add more to existing
		ccm.AddSameStartingCharItems('a', []string{"apricot"}, false)
		actual := args.Map{"result": ccm.LengthOf('a') != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C26_CharCollectionMap_Has(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Has", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})

		// Act
		actual := args.Map{"result": ccm.Has("apple")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have apple", actual)
		actual := args.Map{"result": ccm.Has("cherry")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have cherry", actual)
	})
}

func Test_C26_CharCollectionMap_HasWithCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HasWithCollection", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		has, col := ccm.HasWithCollection("apple")

		// Act
		actual := args.Map{"result": has || col == nil}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should find apple", actual)
		has, _ = ccm.HasWithCollection("xyz")
		actual := args.Map{"result": has}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not find xyz", actual)
	})
}

func Test_C26_CharCollectionMap_HasWithCollectionLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HasWithCollectionLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		has, col := ccm.HasWithCollectionLock("apple")

		// Act
		actual := args.Map{"result": has || col == nil}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should find apple", actual)
	})
}

func Test_C26_CharCollectionMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_LengthOf", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "avocado"})

		// Act
		actual := args.Map{"result": ccm.LengthOf('a') != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual := args.Map{"result": ccm.LengthOf('z') != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharCollectionMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_LengthOfLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": ccm.LengthOfLock('a') != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharCollectionMap_LengthOfCollectionFromFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_LengthOfCollectionFromFirstChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "avocado"})

		// Act
		actual := args.Map{"result": ccm.LengthOfCollectionFromFirstChar("apple") != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual := args.Map{"result": ccm.LengthOfCollectionFromFirstChar("xyz") != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharCollectionMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AllLengthsSum", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "avocado", "banana"})

		// Act
		actual := args.Map{"result": ccm.AllLengthsSum() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C26_CharCollectionMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AllLengthsSumLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": ccm.AllLengthsSumLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharCollectionMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_LengthLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": ccm.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharCollectionMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEmptyLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Empty()
		if !ccm.IsEmptyLock() {
			// items map is non-nil but empty, so IsEmpty checks len

		// Act
			actual := args.Map{"result": false}

		// Assert
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "expected empty", actual)
		}
	})
}

func Test_C26_CharCollectionMap_HasItems(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HasItems", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": ccm.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C26_CharCollectionMap_IsEquals(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEquals", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		ccm2 := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})

		// Act
		actual := args.Map{"result": ccm1.IsEquals(ccm2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C26_CharCollectionMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEqualsLock", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm2 := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": ccm1.IsEqualsLock(ccm2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C26_CharCollectionMap_IsEqualsCaseSensitive(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEqualsCaseSensitive", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.Items([]string{"Apple"})
		ccm2 := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": ccm1.IsEqualsCaseSensitive(true, ccm2)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be equal case-sensitive (different first chars)", actual)
	})
}

func Test_C26_CharCollectionMap_IsEqualsCaseSensitiveLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEqualsCaseSensitiveLock", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm2 := corestr.New.CharCollectionMap.Items([]string{"apple"})

		// Act
		actual := args.Map{"result": ccm1.IsEqualsCaseSensitiveLock(true, ccm2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C26_CharCollectionMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEquals_Nil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"a"})

		// Act
		actual := args.Map{"result": ccm.IsEquals(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not equal nil", actual)
	})
}

func Test_C26_CharCollectionMap_GetMap(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetMap", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m := ccm.GetMap()

		// Act
		actual := args.Map{"result": m == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCopyMapLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m := ccm.GetCopyMapLock()

		// Act
		actual := args.Map{"result": m == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_GetCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCollection", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := ccm.GetCollection("abc", false)

		// Act
		actual := args.Map{"result": col == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find 'a' bucket", actual)
		col2 := ccm.GetCollection("xyz", true)
		actual := args.Map{"result": col2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should create on empty", actual)
		col3 := ccm.GetCollection("zzz", false)
		actual := args.Map{"result": col3 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not create", actual)
	})
}

func Test_C26_CharCollectionMap_GetCollectionLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCollectionLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := ccm.GetCollectionLock("abc", false)

		// Act
		actual := args.Map{"result": col == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find", actual)
	})
}

func Test_C26_CharCollectionMap_GetCollectionByChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCollectionByChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := ccm.GetCollectionByChar('a')

		// Act
		actual := args.Map{"result": col == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find", actual)
	})
}

func Test_C26_CharCollectionMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddCollectionItems", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		ccm.AddCollectionItems(col)

		// Act
		actual := args.Map{"result": ccm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharCollectionMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddCollectionItems_Nil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.AddCollectionItems(nil)

		// Act
		actual := args.Map{"result": ccm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
	})
}

func Test_C26_CharCollectionMap_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddHashmapsValues", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "apple")
		hm.AddOrUpdate("k2", "banana")
		ccm.AddHashmapsValues(hm)

		// Act
		actual := args.Map{"result": ccm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharCollectionMap_AddHashmapsKeysValuesBoth(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddHashmapsKeysValuesBoth", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("key", "val")
		ccm.AddHashmapsKeysValuesBoth(hm)

		// Act
		actual := args.Map{"result": ccm.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C26_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("key", "val")
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Value, true, false
			},
			hm,
		)

		// Act
		actual := args.Map{"result": ccm.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C26_CharCollectionMap_AddCharHashsetMap(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddCharHashsetMap", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		ccm.AddCharHashsetMap(chm)

		// Act
		actual := args.Map{"result": ccm.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharCollectionMap_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameCharsCollection", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "avocado"})
		result := ccm.AddSameCharsCollection("abc", col)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_AddSameCharsCollection_Existing(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameCharsCollection_Existing", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := corestr.New.Collection.Strings([]string{"avocado"})
		result := ccm.AddSameCharsCollection("abc", col)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_AddSameCharsCollection_NilCol(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameCharsCollection_NilCol", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		result := ccm.AddSameCharsCollection("abc", nil)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should create new empty collection", actual)
	})
}

func Test_C26_CharCollectionMap_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameCharsCollectionLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple"})
		result := ccm.AddSameCharsCollectionLock("abc", col)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_Resize(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Resize", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.Resize(100)

		// Act
		actual := args.Map{"result": ccm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should keep items", actual)
	})
}

func Test_C26_CharCollectionMap_AddLength(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddLength", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.AddLength(10, 20)
		// just verifying no panic
	})
}

func Test_C26_CharCollectionMap_List(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_List", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		list := ccm.List()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharCollectionMap_ListLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ListLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		list := ccm.ListLock()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharCollectionMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_SortedListAsc", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"cherry", "apple", "banana"})
		list := ccm.SortedListAsc()

		// Act
		actual := args.Map{"result": len(list) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C26_CharCollectionMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCharsGroups", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		result := ccm.GetCharsGroups([]string{"apple", "banana"})

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetByChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := ccm.HashsetByChar('a')

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetByCharLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := ccm.HashsetByCharLock('a')

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetByStringFirstChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := ccm.HashsetByStringFirstChar("abc")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetByStringFirstCharLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := ccm.HashsetByStringFirstCharLock("abc")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetsCollection", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := ccm.HashsetsCollection()

		// Act
		actual := args.Map{"result": hsc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetsCollectionByChars", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := ccm.HashsetsCollectionByChars('a', 'b')

		// Act
		actual := args.Map{"result": hsc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetsCollectionByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetsCollectionByStringFirstChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := ccm.HashsetsCollectionByStringFirstChar("abc", "bcd")

		// Act
		actual := args.Map{"result": hsc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_SummaryString", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := ccm.SummaryString()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharCollectionMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_SummaryStringLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := ccm.SummaryStringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharCollectionMap_String(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_String", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := ccm.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharCollectionMap_StringLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_StringLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := ccm.StringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharCollectionMap_Print(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Print", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.Print(false) // no output
		ccm.Print(true)  // prints
	})
}

func Test_C26_CharCollectionMap_PrintLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_PrintLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.PrintLock(false)
		ccm.PrintLock(true)
	})
}

func Test_C26_CharCollectionMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_JsonModel", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		model := ccm.JsonModel()

		// Act
		actual := args.Map{"result": model == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_JsonModelAny", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"a"})

		// Act
		actual := args.Map{"result": ccm.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_MarshalJSON", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		data, err := json.Marshal(ccm)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have data", actual)
	})
}

func Test_C26_CharCollectionMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_UnmarshalJSON", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		data, _ := json.Marshal(ccm)
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		err := json.Unmarshal(data, ccm2)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C26_CharCollectionMap_Json(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Json", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		result := ccm.Json()

		// Act
		actual := args.Map{"result": result.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not error", actual)
	})
}

func Test_C26_CharCollectionMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_JsonPtr", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ptr := ccm.JsonPtr()

		// Act
		actual := args.Map{"result": ptr == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ParseInjectUsingJson", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		jsonResult := ccm.Json()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		_, err := ccm2.ParseInjectUsingJson(&jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C26_CharCollectionMap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ParseInjectUsingJsonMust", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		jsonResult := ccm.Json()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		result := ccm2.ParseInjectUsingJsonMust(&jsonResult)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_JsonParseSelfInject", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		jsonResult := ccm.Json()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		err := ccm2.JsonParseSelfInject(&jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C26_CharCollectionMap_AsJsonInterfaces(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AsJsonInterfaces", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"a"})

		// Act
		actual := args.Map{"result": ccm.AsJsonContractsBinder() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": ccm.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": ccm.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": ccm.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_Clear(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Clear", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.Clear()

		// Act
		actual := args.Map{"result": ccm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
	})
}

func Test_C26_CharCollectionMap_Dispose(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Dispose", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.Dispose()
	})
}

func Test_C26_CharCollectionMap_DataModel(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_DataModel", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		dm := corestr.NewCharCollectionMapDataModelUsing(ccm)

		// Act
		actual := args.Map{"result": dm == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		ccm2 := corestr.NewCharCollectionMapUsingDataModel(dm)
		actual := args.Map{"result": ccm2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

// =======================================================
// CharHashsetMap
// =======================================================

func Test_C26_CharHashsetMap_Cap(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Cap", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(20, 10)

		// Act
		actual := args.Map{"result": chm == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_CapItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_CapItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")

		// Act
		actual := args.Map{"result": chm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_Strings(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Strings", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Strings(10, []string{"apple", "banana", "avocado"})

		// Act
		actual := args.Map{"result": chm.AllLengthsSum() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C26_CharHashsetMap_Strings_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Strings_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Strings(10, nil)

		// Act
		actual := args.Map{"result": chm == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_GetChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"result": chm.GetChar("hello") != 'h'}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		actual := args.Map{"result": chm.GetChar("") != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharHashsetMap_GetCharOf(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetCharOf", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"result": chm.GetCharOf("hello") != 'h'}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		actual := args.Map{"result": chm.GetCharOf("") != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharHashsetMap_Add(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Add", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Add("avocado")
		chm.Add("banana")

		// Act
		actual := args.Map{"result": chm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AddLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddLock("apple")

		// Act
		actual := args.Map{"result": chm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharHashsetMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddStrings", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")

		// Act
		actual := args.Map{"result": chm.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddStringsLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStringsLock("apple", "banana")

		// Act
		actual := args.Map{"result": chm.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddSameStartingCharItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddSameStartingCharItems('a', []string{"apple", "avocado"})

		// Act
		actual := args.Map{"result": chm.LengthOf('a') != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// Add to existing
		chm.AddSameStartingCharItems('a', []string{"apricot"})
		actual := args.Map{"result": chm.LengthOf('a') != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C26_CharHashsetMap_Has(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Has", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")

		// Act
		actual := args.Map{"result": chm.Has("apple")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have apple", actual)
		actual := args.Map{"result": chm.Has("cherry")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have cherry", actual)
	})
}

func Test_C26_CharHashsetMap_HasWithHashset(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HasWithHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		has, hs := chm.HasWithHashset("apple")

		// Act
		actual := args.Map{"result": has || hs == nil}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should find apple", actual)
		has, _ = chm.HasWithHashset("xyz")
		actual := args.Map{"result": has}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not find xyz", actual)
	})
}

func Test_C26_CharHashsetMap_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HasWithHashsetLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		has, hs := chm.HasWithHashsetLock("apple")

		// Act
		actual := args.Map{"result": has || hs == nil}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should find apple", actual)
	})
}

func Test_C26_CharHashsetMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_LengthOf", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "avocado")

		// Act
		actual := args.Map{"result": chm.LengthOf('a') != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_LengthOfLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": chm.LengthOfLock('a') != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharHashsetMap_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_LengthOfHashsetFromFirstChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "avocado")

		// Act
		actual := args.Map{"result": chm.LengthOfHashsetFromFirstChar("abc") != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AllLengthsSum", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")

		// Act
		actual := args.Map{"result": chm.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AllLengthsSumLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": chm.AllLengthsSumLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharHashsetMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_LengthLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": chm.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharHashsetMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_IsEmptyLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"result": chm.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C26_CharHashsetMap_HasItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HasItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": chm.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C26_CharHashsetMap_IsEquals(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_IsEquals", func() {
		// Arrange
		chm1 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		chm2 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")

		// Act
		actual := args.Map{"result": chm1.IsEquals(chm2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C26_CharHashsetMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_IsEqualsLock", func() {
		// Arrange
		chm1 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm2 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": chm1.IsEqualsLock(chm2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C26_CharHashsetMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_IsEquals_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "a")

		// Act
		actual := args.Map{"result": chm.IsEquals(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not equal nil", actual)
	})
}

func Test_C26_CharHashsetMap_GetMap(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetMap", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": chm.GetMap() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetCopyMapLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")

		// Act
		actual := args.Map{"result": chm.GetCopyMapLock() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_GetHashset(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.GetHashset("abc", false)

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find 'a' bucket", actual)
		hs2 := chm.GetHashset("xyz", true)
		actual := args.Map{"result": hs2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should create on empty", actual)
	})
}

func Test_C26_CharHashsetMap_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetHashsetLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.GetHashsetLock(false, "abc")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find", actual)
	})
}

func Test_C26_CharHashsetMap_GetHashsetByChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetHashsetByChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.GetHashsetByChar('a')

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetByChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.HashsetByChar('a')

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetByCharLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.HashsetByCharLock('a')

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		hs2 := chm.HashsetByCharLock('z')
		actual := args.Map{"result": hs2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should return empty, not nil", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetByStringFirstChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.HashsetByStringFirstChar("abc")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetByStringFirstCharLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.HashsetByStringFirstCharLock("abc")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddCollectionItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		chm.AddCollectionItems(col)

		// Act
		actual := args.Map{"result": chm.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddCollectionItems_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddCollectionItems(nil)

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
	})
}

func Test_C26_CharHashsetMap_AddCharCollectionMapItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddCharCollectionMapItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		chm.AddCharCollectionMapItems(ccm)

		// Act
		actual := args.Map{"result": chm.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddHashsetItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple", "banana"})
		chm.AddHashsetItems(hs)

		// Act
		actual := args.Map{"result": chm.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddSameCharsCollection", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "avocado"})
		result := chm.AddSameCharsCollection("abc", col)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_AddSameCharsHashset(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddSameCharsHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		result := chm.AddSameCharsHashset("abc", hs)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_AddHashsetLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddHashsetLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		result := chm.AddHashsetLock("abc", hs)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddSameCharsCollectionLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple"})
		result := chm.AddSameCharsCollectionLock("abc", col)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetsCollection", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		hsc := chm.HashsetsCollection()

		// Act
		actual := args.Map{"result": hsc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetsCollectionByChars", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		hsc := chm.HashsetsCollectionByChars('a', 'b')

		// Act
		actual := args.Map{"result": hsc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetsCollectionByStringsFirstChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		hsc := chm.HashsetsCollectionByStringsFirstChar("abc", "bcd")

		// Act
		actual := args.Map{"result": hsc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetCharsGroups", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		result := chm.GetCharsGroups("apple", "banana")

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_List(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_List", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		list := chm.List()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_SortedListAsc", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "cherry", "apple", "banana")
		list := chm.SortedListAsc()

		// Act
		actual := args.Map{"result": len(list) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual := args.Map{"result": list[0] != "apple"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first should be apple", actual)
	})
}

func Test_C26_CharHashsetMap_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_SortedListDsc", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "cherry")
		list := chm.SortedListDsc()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual := args.Map{"result": list[0] != "cherry"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first should be cherry", actual)
	})
}

func Test_C26_CharHashsetMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_SummaryString", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := chm.SummaryString()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharHashsetMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_SummaryStringLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := chm.SummaryStringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharHashsetMap_String(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_String", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := chm.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharHashsetMap_StringLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_StringLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := chm.StringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharHashsetMap_Print(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Print", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm.Print(false)
		chm.Print(true)
	})
}

func Test_C26_CharHashsetMap_PrintLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_PrintLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm.PrintLock(false)
		chm.PrintLock(true)
	})
}

func Test_C26_CharHashsetMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_JsonModel", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		model := chm.JsonModel()

		// Act
		actual := args.Map{"result": model == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_JsonModelAny", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "a")

		// Act
		actual := args.Map{"result": chm.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_MarshalJSON", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		data, err := json.Marshal(chm)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have data", actual)
	})
}

func Test_C26_CharHashsetMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_UnmarshalJSON", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		data, _ := json.Marshal(chm)
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := json.Unmarshal(data, chm2)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C26_CharHashsetMap_Json(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Json", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		result := chm.Json()

		// Act
		actual := args.Map{"result": result.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not error", actual)
	})
}

func Test_C26_CharHashsetMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_JsonPtr", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		ptr := chm.JsonPtr()

		// Act
		actual := args.Map{"result": ptr == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_ParseInjectUsingJson", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		jsonResult := chm.Json()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		_, err := chm2.ParseInjectUsingJson(&jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C26_CharHashsetMap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_ParseInjectUsingJsonMust", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		jsonResult := chm.Json()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		result := chm2.ParseInjectUsingJsonMust(&jsonResult)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_JsonParseSelfInject", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		jsonResult := chm.Json()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := chm2.JsonParseSelfInject(&jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C26_CharHashsetMap_AsJsonInterfaces(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AsJsonInterfaces", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "a")

		// Act
		actual := args.Map{"result": chm.AsJsonContractsBinder() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": chm.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": chm.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": chm.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_RemoveAll(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_RemoveAll", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm.RemoveAll()
		time.Sleep(10 * time.Millisecond) // allow goroutine cleanup

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
	})
}

func Test_C26_CharHashsetMap_Clear(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Clear", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm.Clear()
		time.Sleep(10 * time.Millisecond)

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
	})
}

func Test_C26_CharHashsetMap_DataModel(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_DataModel", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		dm := corestr.NewCharHashsetMapDataModelUsing(chm)

		// Act
		actual := args.Map{"result": dm == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		chm2 := corestr.NewCharHashsetMapUsingDataModel(dm)
		actual := args.Map{"result": chm2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_AddCollectionItemsAsyncLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddCollectionItemsAsyncLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		done := make(chan bool, 1)
		chm.AddCollectionItemsAsyncLock(col, func(charHashset *corestr.CharHashsetMap) {
			done <- true
		})
		select {
		case <-done:
		case <-time.After(2 * time.Second):

		// Act
			actual := args.Map{"result": false}

		// Assert
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "timeout", actual)
		}
	})
}

func Test_C26_CharHashsetMap_AddHashsetItemsAsyncLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddHashsetItemsAsyncLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		done := make(chan bool, 1)
		chm.AddHashsetItemsAsyncLock(hs, func(charHashset *corestr.CharHashsetMap) {
			done <- true
		})
		select {
		case <-done:
		case <-time.After(2 * time.Second):

		// Act
			actual := args.Map{"result": false}

		// Assert
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "timeout", actual)
		}
	})
}
