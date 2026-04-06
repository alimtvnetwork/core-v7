package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// CharHashsetMap.go — Full coverage (~421 uncovered stmts, 1193 lines)
// =============================================================================

// ── GetChar / GetCharOf ──

func Test_Cov63_CharHashsetMap_GetChar_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetChar_NonEmpty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"char": chm.GetChar("hello")}

		// Assert
		expected := args.Map{"char": byte('h')}
		expected.ShouldBeEqual(t, 0, "GetChar returns first byte", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetChar_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetChar_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"char": chm.GetChar("")}

		// Assert
		expected := args.Map{"char": byte(0)}
		expected.ShouldBeEqual(t, 0, "GetChar empty returns 0", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetCharOf_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetCharOf_NonEmpty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"char": chm.GetCharOf("abc")}

		// Assert
		expected := args.Map{"char": byte('a')}
		expected.ShouldBeEqual(t, 0, "GetCharOf returns first byte", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetCharOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetCharOf_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"char": chm.GetCharOf("")}

		// Assert
		expected := args.Map{"char": byte(0)}
		expected.ShouldBeEqual(t, 0, "GetCharOf empty returns 0", actual)
	})
}

// ── IsEmpty / HasItems / Length ──

func Test_Cov63_CharHashsetMap_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEmpty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"empty": chm.IsEmpty(), "hasItems": chm.HasItems()}

		// Assert
		expected := args.Map{"empty": true, "hasItems": false}
		expected.ShouldBeEqual(t, 0, "IsEmpty on new", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEmpty_Nil", func() {
		// Act
		actual := args.Map{"empty": corestr.Empty.CharHashsetMap().IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmpty on nil items", actual)
	})
}

func Test_Cov63_CharHashsetMap_HasItems_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasItems_NonEmpty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("hello")

		// Act
		actual := args.Map{"hasItems": chm.HasItems()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "HasItems after add", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEmptyLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"empty": chm.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_Length(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Length", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("hello")
		chm.Add("world")

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Length returns char groups", actual)
	})
}

func Test_Cov63_CharHashsetMap_Length_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Length_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_LengthLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("hello")

		// Act
		actual := args.Map{"len": chm.LengthLock()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock", actual)
	})
}

// ── Add / AddStrings ──

func Test_Cov63_CharHashsetMap_Add(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Add", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Add("avocado")

		// Act
		actual := args.Map{"len": chm.Length(), "allLen": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"len": 1, "allLen": 2}
		expected.ShouldBeEqual(t, 0, "Add same char groups", actual)
	})
}

func Test_Cov63_CharHashsetMap_Add_ExistingChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Add_ExistingChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Add("apricot")

		// Act
		actual := args.Map{"allLen": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"allLen": 2}
		expected.ShouldBeEqual(t, 0, "Add to existing char bucket", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStrings", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana", "avocado")

		// Act
		actual := args.Map{"len": chm.Length(), "allLen": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"len": 2, "allLen": 3}
		expected.ShouldBeEqual(t, 0, "AddStrings", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStrings_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings()

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings nil", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddLock("apple")
		chm.AddLock("avocado")

		// Act
		actual := args.Map{"allLen": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"allLen": 2}
		expected.ShouldBeEqual(t, 0, "AddLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddLock_NewChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddLock_NewChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddLock("apple")
		chm.AddLock("banana")

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddLock new char", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStringsLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStringsLock("apple", "banana")

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddStringsLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStringsLock_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStringsLock()

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsLock nil", actual)
	})
}

// ── AddSameStartingCharItems ──

func Test_Cov63_CharHashsetMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameStartingCharItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddSameStartingCharItems('a', []string{"apple", "avocado"})

		// Act
		actual := args.Map{"allLen": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"allLen": 2}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameStartingCharItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameStartingCharItems_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddSameStartingCharItems('a', []string{})

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameStartingCharItems_Existing(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameStartingCharItems_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.AddSameStartingCharItems('a', []string{"avocado"})

		// Act
		actual := args.Map{"allLen": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"allLen": 2}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems existing", actual)
	})
}

// ── AddCollectionItems / AddCharCollectionMapItems ──

func Test_Cov63_CharHashsetMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddCollectionItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		chm.AddCollectionItems(col)

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddCollectionItems", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddCollectionItems_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddCollectionItems(nil)

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollectionItems nil", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddCharCollectionMapItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddCharCollectionMapItems_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddCharCollectionMapItems(nil)

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCharCollectionMapItems nil", actual)
	})
}

// ── AddHashsetItems ──

func Test_Cov63_CharHashsetMap_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddHashsetItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.StringsSpreadItems("apple", "banana")
		chm.AddHashsetItems(hs)

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems", actual)
	})
}

// ── Has ──

func Test_Cov63_CharHashsetMap_Has(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Has", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")

		// Act
		actual := args.Map{"has": chm.Has("apple"), "miss": chm.Has("banana")}

		// Assert
		expected := args.Map{"has": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "Has", actual)
	})
}

func Test_Cov63_CharHashsetMap_Has_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Has_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"has": chm.Has("x")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Has empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_Has_CharExistsButNotStr(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Has_CharExistsButNotStr", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")

		// Act
		actual := args.Map{"has": chm.Has("avocado")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Has char exists but not str", actual)
	})
}

// ── HasWithHashset ──

func Test_Cov63_CharHashsetMap_HasWithHashset(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasWithHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		has, hs := chm.HasWithHashset("apple")

		// Act
		actual := args.Map{"has": has, "hsNonNil": hs != nil}

		// Assert
		expected := args.Map{"has": true, "hsNonNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashset found", actual)
	})
}

func Test_Cov63_CharHashsetMap_HasWithHashset_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasWithHashset_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		has, hs := chm.HasWithHashset("x")

		// Act
		actual := args.Map{"has": has, "hsEmpty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"has": false, "hsEmpty": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashset empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_HasWithHashset_MissChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasWithHashset_MissChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		has, hs := chm.HasWithHashset("banana")

		// Act
		actual := args.Map{"has": has, "hsEmpty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"has": false, "hsEmpty": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashset miss char", actual)
	})
}

func Test_Cov63_CharHashsetMap_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasWithHashsetLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		has, hs := chm.HasWithHashsetLock("apple")

		// Act
		actual := args.Map{"has": has, "hsNonNil": hs != nil}

		// Assert
		expected := args.Map{"has": true, "hsNonNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_HasWithHashsetLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasWithHashsetLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		has, _ := chm.HasWithHashsetLock("x")

		// Act
		actual := args.Map{"has": has}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_HasWithHashsetLock_MissChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasWithHashsetLock_MissChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		has, _ := chm.HasWithHashsetLock("banana")

		// Act
		actual := args.Map{"has": has}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock miss char", actual)
	})
}

// ── LengthOf / LengthOfLock / LengthOfHashsetFromFirstChar ──

func Test_Cov63_CharHashsetMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_LengthOf", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Add("avocado")

		// Act
		actual := args.Map{"len": chm.LengthOf('a'), "miss": chm.LengthOf('z')}

		// Assert
		expected := args.Map{"len": 2, "miss": 0}
		expected.ShouldBeEqual(t, 0, "LengthOf", actual)
	})
}

func Test_Cov63_CharHashsetMap_LengthOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_LengthOf_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"len": chm.LengthOf('a')}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthOf empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_LengthOfLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")

		// Act
		actual := args.Map{"len": chm.LengthOfLock('a')}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthOfLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_LengthOfLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_LengthOfLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"len": chm.LengthOfLock('a')}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthOfLock empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_LengthOfHashsetFromFirstChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Add("avocado")

		// Act
		actual := args.Map{"len": chm.LengthOfHashsetFromFirstChar("abc"), "miss": chm.LengthOfHashsetFromFirstChar("xyz")}

		// Assert
		expected := args.Map{"len": 2, "miss": 0}
		expected.ShouldBeEqual(t, 0, "LengthOfHashsetFromFirstChar", actual)
	})
}

// ── AllLengthsSum / AllLengthsSumLock ──

func Test_Cov63_CharHashsetMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AllLengthsSum", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana", "avocado")

		// Act
		actual := args.Map{"sum": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"sum": 3}
		expected.ShouldBeEqual(t, 0, "AllLengthsSum", actual)
	})
}

func Test_Cov63_CharHashsetMap_AllLengthsSum_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AllLengthsSum_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"sum": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"sum": 0}
		expected.ShouldBeEqual(t, 0, "AllLengthsSum empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AllLengthsSumLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")

		// Act
		actual := args.Map{"sum": chm.AllLengthsSumLock()}

		// Assert
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AllLengthsSumLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_AllLengthsSumLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AllLengthsSumLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"sum": chm.AllLengthsSumLock()}

		// Assert
		expected := args.Map{"sum": 0}
		expected.ShouldBeEqual(t, 0, "AllLengthsSumLock empty", actual)
	})
}

// ── GetMap / GetCopyMapLock ──

func Test_Cov63_CharHashsetMap_GetMap(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetMap", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")

		// Act
		actual := args.Map{"nonNil": chm.GetMap() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "GetMap", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetCopyMapLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		cp := chm.GetCopyMapLock()

		// Act
		actual := args.Map{"len": len(cp)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetCopyMapLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetCopyMapLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		cp := chm.GetCopyMapLock()

		// Act
		actual := args.Map{"len": len(cp)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "GetCopyMapLock empty", actual)
	})
}

// ── List / SortedListAsc / SortedListDsc ──

func Test_Cov63_CharHashsetMap_List(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_List", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("banana", "apple")

		// Act
		actual := args.Map{"len": len(chm.List())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "List", actual)
	})
}

func Test_Cov63_CharHashsetMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_SortedListAsc", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("banana", "apple")
		list := chm.SortedListAsc()

		// Act
		actual := args.Map{"first": list[0], "second": list[1]}

		// Assert
		expected := args.Map{"first": "apple", "second": "banana"}
		expected.ShouldBeEqual(t, 0, "SortedListAsc", actual)
	})
}

func Test_Cov63_CharHashsetMap_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_SortedListDsc", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("banana", "apple")
		list := chm.SortedListDsc()

		// Act
		actual := args.Map{"first": list[0], "second": list[1]}

		// Assert
		expected := args.Map{"first": "banana", "second": "apple"}
		expected.ShouldBeEqual(t, 0, "SortedListDsc", actual)
	})
}

// ── String / SummaryString / StringLock / SummaryStringLock ──

func Test_Cov63_CharHashsetMap_String(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_String", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")

		// Act
		actual := args.Map{"nonEmpty": chm.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String", actual)
	})
}

func Test_Cov63_CharHashsetMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_SummaryString", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")

		// Act
		actual := args.Map{"nonEmpty": chm.SummaryString() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryString", actual)
	})
}

func Test_Cov63_CharHashsetMap_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_StringLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")

		// Act
		actual := args.Map{"nonEmpty": chm.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_SummaryStringLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")

		// Act
		actual := args.Map{"nonEmpty": chm.SummaryStringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringLock", actual)
	})
}

// ── Print / PrintLock ──

func Test_Cov63_CharHashsetMap_Print_True(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Print_True", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Print(true)

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Print true", actual)
	})
}

func Test_Cov63_CharHashsetMap_Print_False(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Print_False", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Print(false)

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Print false skips", actual)
	})
}

func Test_Cov63_CharHashsetMap_PrintLock_True(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_PrintLock_True", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.PrintLock(true)

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "PrintLock true", actual)
	})
}

func Test_Cov63_CharHashsetMap_PrintLock_False(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_PrintLock_False", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.PrintLock(false)

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "PrintLock false skips", actual)
	})
}

// ── IsEquals / IsEqualsLock ──

func Test_Cov63_CharHashsetMap_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_Same", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")

		// Act
		actual := args.Map{"eq": chm.IsEquals(chm)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals same ptr", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"eq": chm.IsEquals(nil)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals nil", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.New.CharHashsetMap.Cap(10, 10)
		b := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_OneEmpty", func() {
		// Arrange
		a := corestr.New.CharHashsetMap.Cap(10, 10)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals one empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_DiffLen", func() {
		// Arrange
		a := corestr.New.CharHashsetMap.Cap(10, 10)
		a.AddStrings("apple", "banana")
		b := corestr.New.CharHashsetMap.Cap(10, 10)
		b.Add("apple")

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff len", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEquals_DiffContent(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_DiffContent", func() {
		// Arrange
		a := corestr.New.CharHashsetMap.Cap(10, 10)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 10)
		b.Add("avocado")

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff content", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEquals_MissingKey(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_MissingKey", func() {
		// Arrange
		a := corestr.New.CharHashsetMap.Cap(10, 10)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 10)
		b.Add("banana")

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals missing key", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEqualsLock", func() {
		// Arrange
		a := corestr.New.CharHashsetMap.Cap(10, 10)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 10)
		b.Add("apple")

		// Act
		actual := args.Map{"eq": a.IsEqualsLock(b)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsLock", actual)
	})
}

// ── GetHashset / GetHashsetLock ──

func Test_Cov63_CharHashsetMap_GetHashset_Found(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetHashset_Found", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.GetHashset("avocado", false)

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashset found", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetHashset_MissNoAdd(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetHashset_MissNoAdd", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := chm.GetHashset("x", false)

		// Act
		actual := args.Map{"nil": hs == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "GetHashset miss no add", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetHashset_MissAddNew(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetHashset_MissAddNew", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := chm.GetHashset("x", true)

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashset miss add new", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetHashsetLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := chm.GetHashsetLock(true, "x")

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashsetLock", actual)
	})
}

// ── GetHashsetByChar / HashsetByChar / HashsetByCharLock ──

func Test_Cov63_CharHashsetMap_GetHashsetByChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetHashsetByChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.GetHashsetByChar('a')

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashsetByChar", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetByChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.HashsetByChar('a')

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByChar", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetByCharLock_Found(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetByCharLock_Found", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.HashsetByCharLock('a')

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock found", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetByCharLock_Miss(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetByCharLock_Miss", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := chm.HashsetByCharLock('z')

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock miss returns empty", actual)
	})
}

// ── HashsetByStringFirstChar / HashsetByStringFirstCharLock ──

func Test_Cov63_CharHashsetMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetByStringFirstChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.HashsetByStringFirstChar("avocado")

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstChar", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetByStringFirstCharLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.HashsetByStringFirstCharLock("avocado")

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstCharLock", actual)
	})
}

// ── HashsetsCollection / HashsetsCollectionByChars / HashsetsCollectionByStringsFirstChar ──

func Test_Cov63_CharHashsetMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollection", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		hsc := chm.HashsetsCollection()

		// Act
		actual := args.Map{"nonNil": hsc != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollection_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hsc := chm.HashsetsCollection()

		// Act
		actual := args.Map{"empty": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollectionByChars", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		hsc := chm.HashsetsCollectionByChars('a', 'b')

		// Act
		actual := args.Map{"nonNil": hsc != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetsCollectionByChars_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollectionByChars_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hsc := chm.HashsetsCollectionByChars('a')

		// Act
		actual := args.Map{"empty": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetsCollectionByChars_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollectionByChars_NilHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hsc := chm.HashsetsCollectionByChars('a', 'z')

		// Act
		actual := args.Map{"nonNil": hsc != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars skips nil", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollectionByStringsFirstChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		hsc := chm.HashsetsCollectionByStringsFirstChar("avocado", "berry")

		// Act
		actual := args.Map{"nonNil": hsc != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringsFirstChar", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetsCollectionByStringsFirstChar_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollectionByStringsFirstChar_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hsc := chm.HashsetsCollectionByStringsFirstChar("x")

		// Act
		actual := args.Map{"empty": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringsFirstChar empty", actual)
	})
}

// ── GetCharsGroups ──

func Test_Cov63_CharHashsetMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetCharsGroups", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		r := chm.GetCharsGroups("apple", "banana", "avocado")

		// Act
		actual := args.Map{"len": r.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetCharsGroups", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetCharsGroups_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		r := chm.GetCharsGroups()

		// Act
		actual := args.Map{"same": r == chm}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "GetCharsGroups empty returns self", actual)
	})
}

// ── AddSameCharsCollection ──

func Test_Cov63_CharHashsetMap_AddSameCharsCollection_New(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollection_New", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "avocado"})
		hs := chm.AddSameCharsCollection("a", col)

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection new", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsCollection_Existing(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollection_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		col := corestr.New.Collection.Strings([]string{"avocado"})
		hs := chm.AddSameCharsCollection("a", col)

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection existing", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsCollection_NilCol(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollection_NilCol", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := chm.AddSameCharsCollection("a", nil)

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection nil creates empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsCollection_ExistingNilCol(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollection_ExistingNilCol", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.AddSameCharsCollection("a", nil)

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection existing nil col", actual)
	})
}

// ── AddSameCharsHashset ──

func Test_Cov63_CharHashsetMap_AddSameCharsHashset_New(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsHashset_New", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.StringsSpreadItems("apple")
		r := chm.AddSameCharsHashset("a", hs)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset new", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsHashset_Existing(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsHashset_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := corestr.New.Hashset.StringsSpreadItems("avocado")
		r := chm.AddSameCharsHashset("a", hs)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset existing", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsHashset_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsHashset_NilHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		r := chm.AddSameCharsHashset("a", nil)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset nil creates empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsHashset_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsHashset_ExistingNil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		r := chm.AddSameCharsHashset("a", nil)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset existing nil", actual)
	})
}

// ── AddSameCharsCollectionLock ──

func Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_New(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_New", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple"})
		hs := chm.AddSameCharsCollectionLock("a", col)

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock new", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		col := corestr.New.Collection.Strings([]string{"avocado"})
		hs := chm.AddSameCharsCollectionLock("a", col)

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock existing", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_NilCol(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_NilCol", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := chm.AddSameCharsCollectionLock("a", nil)

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock nil", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_ExistingNil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.AddSameCharsCollectionLock("a", nil)

		// Act
		actual := args.Map{"nonNil": hs != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock existing nil", actual)
	})
}

// ── AddHashsetLock ──

func Test_Cov63_CharHashsetMap_AddHashsetLock_New(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddHashsetLock_New", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.StringsSpreadItems("apple")
		r := chm.AddHashsetLock("a", hs)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock new", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddHashsetLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddHashsetLock_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := corestr.New.Hashset.StringsSpreadItems("avocado")
		r := chm.AddHashsetLock("a", hs)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock existing", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddHashsetLock_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddHashsetLock_NilHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		r := chm.AddHashsetLock("a", nil)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock nil creates empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddHashsetLock_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddHashsetLock_ExistingNil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		r := chm.AddHashsetLock("a", nil)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock existing nil", actual)
	})
}

// ── JSON ──

func Test_Cov63_CharHashsetMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_JsonModel", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")

		// Act
		actual := args.Map{"nonNil": chm.JsonModel() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModel", actual)
	})
}

func Test_Cov63_CharHashsetMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_JsonModelAny", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"nonNil": chm.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny", actual)
	})
}

func Test_Cov63_CharHashsetMap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_MarshalJSON", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		b, err := chm.MarshalJSON()

		// Act
		actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}

		// Assert
		expected := args.Map{"noErr": true, "nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "MarshalJSON", actual)
	})
}

func Test_Cov63_CharHashsetMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_UnmarshalJSON", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		b, _ := chm.MarshalJSON()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := chm2.UnmarshalJSON(b)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON", actual)
	})
}

func Test_Cov63_CharHashsetMap_UnmarshalJSON_Error(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_UnmarshalJSON_Error", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		err := chm.UnmarshalJSON([]byte("invalid"))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON error", actual)
	})
}

func Test_Cov63_CharHashsetMap_Json(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Json", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		r := chm.Json()

		// Act
		actual := args.Map{"nonEmpty": r.JsonString() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Json", actual)
	})
}

func Test_Cov63_CharHashsetMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_JsonPtr", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		r := chm.JsonPtr()

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr", actual)
	})
}

func Test_Cov63_CharHashsetMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_ParseInjectUsingJson", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		r, err := chm2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil, "nonNil": r != nil}

		// Assert
		expected := args.Map{"noErr": true, "nonNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson", actual)
	})
}

func Test_Cov63_CharHashsetMap_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_ParseInjectUsingJson_Error", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		jr := &corejson.Result{Error: errors.New("fail")}
		_, err := chm.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson error", actual)
	})
}

func Test_Cov63_CharHashsetMap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_ParseInjectUsingJsonMust", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		r := chm2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust", actual)
	})
}

func Test_Cov63_CharHashsetMap_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_ParseInjectUsingJsonMust_Panics", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		jr := &corejson.Result{Error: errors.New("fail")}
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			chm.ParseInjectUsingJsonMust(jr)
		}()

		// Act
		actual := args.Map{"panicked": panicked}

		// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust panics", actual)
	})
}

func Test_Cov63_CharHashsetMap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_JsonParseSelfInject", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := chm2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject", actual)
	})
}

// ── Interface casts ──

func Test_Cov63_CharHashsetMap_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AsJsoner", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"nonNil": chm.AsJsoner() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsoner", actual)
	})
}

func Test_Cov63_CharHashsetMap_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AsJsonContractsBinder", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"nonNil": chm.AsJsonContractsBinder() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder", actual)
	})
}

func Test_Cov63_CharHashsetMap_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AsJsonMarshaller", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"nonNil": chm.AsJsonMarshaller() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonMarshaller", actual)
	})
}

func Test_Cov63_CharHashsetMap_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AsJsonParseSelfInjector", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)

		// Act
		actual := args.Map{"nonNil": chm.AsJsonParseSelfInjector() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonParseSelfInjector", actual)
	})
}

// ── RemoveAll / Clear ──

func Test_Cov63_CharHashsetMap_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_RemoveAll", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		chm.RemoveAll()

		// Act
		actual := args.Map{"empty": chm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "RemoveAll", actual)
	})
}

func Test_Cov63_CharHashsetMap_RemoveAll_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_RemoveAll_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.RemoveAll()

		// Act
		actual := args.Map{"empty": chm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "RemoveAll on empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_Clear(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Clear", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		chm.Clear()

		// Act
		actual := args.Map{"empty": chm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear", actual)
	})
}

func Test_Cov63_CharHashsetMap_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Clear_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Clear()

		// Act
		actual := args.Map{"empty": chm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear empty", actual)
	})
}

// ── Constructors ──

func Test_Cov63_CharHashsetMap_Cap_MinEnforced(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Cap_MinEnforced", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(1, 1)

		// Act
		actual := args.Map{"nonNil": chm != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "Cap enforces minimum", actual)
	})
}

func Test_Cov63_CharHashsetMap_CapItems(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_CapItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CapItems", actual)
	})
}

func Test_Cov63_CharHashsetMap_Strings(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Strings", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Strings(10, []string{"apple", "banana"})

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Strings", actual)
	})
}

func Test_Cov63_CharHashsetMap_Strings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Strings_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Strings(10, nil)

		// Act
		actual := args.Map{"nonNil": chm != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "Strings nil", actual)
	})
}

// ── AddStringsAsyncLock ──

func Test_Cov63_CharHashsetMap_AddStringsAsyncLock_Small(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStringsAsyncLock_Small", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		called := false
		chm.AddStringsAsyncLock([]string{"apple", "banana"}, func(c *corestr.CharHashsetMap) {
			called = true
		})

		// Act
		actual := args.Map{"called": called, "len": chm.Length()}

		// Assert
		expected := args.Map{"called": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsAsyncLock small", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddStringsAsyncLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStringsAsyncLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStringsAsyncLock([]string{}, nil)

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsAsyncLock empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddStringsAsyncLock_NilCallback(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStringsAsyncLock_NilCallback", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStringsAsyncLock([]string{"apple"}, nil)

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStringsAsyncLock nil callback", actual)
	})
}

// ── AddCollectionItemsAsyncLock ──

func Test_Cov63_CharHashsetMap_AddCollectionItemsAsyncLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddCollectionItemsAsyncLock_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddCollectionItemsAsyncLock(nil, nil)

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollectionItemsAsyncLock nil", actual)
	})
}

// ── AddHashsetItemsAsyncLock ──

func Test_Cov63_CharHashsetMap_AddHashsetItemsAsyncLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddHashsetItemsAsyncLock_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddHashsetItemsAsyncLock(nil, nil)

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashsetItemsAsyncLock nil", actual)
	})
}
