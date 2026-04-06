package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CharHashsetMap — Segment 6b
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg6_CHM_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEmpty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)

		// Act
		actual := args.Map{
			"empty": chm.IsEmpty(),
			"hasItems": chm.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_Seg6_CHM_Add(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Add", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("avocado").Add("banana")

		// Act
		actual := args.Map{
			"len": chm.Length(),
			"allLen": chm.AllLengthsSum(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"allLen": 3,
		}
		expected.ShouldBeEqual(t, 0, "Add -- 2 groups 3 total", actual)
	})
}

func Test_Seg6_CHM_AddLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddLock("apple").AddLock("banana")

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddLock -- 2 groups", actual)
	})
}

func Test_Seg6_CHM_AddStrings(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddStrings", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddStrings("apple", "avocado", "banana")

		// Act
		actual := args.Map{
			"len": chm.Length(),
			"sum": chm.AllLengthsSum(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"sum": 3,
		}
		expected.ShouldBeEqual(t, 0, "AddStrings -- 2 groups 3 total", actual)
	})
}

func Test_Seg6_CHM_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddStrings_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddStrings()

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings empty -- no change", actual)
	})
}

func Test_Seg6_CHM_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddStringsLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddStringsLock("apple", "banana")

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsLock -- 2 groups", actual)
	})
}

func Test_Seg6_CHM_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddStringsLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddStringsLock()

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsLock empty -- no change", actual)
	})
}

func Test_Seg6_CHM_GetChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)

		// Act
		actual := args.Map{
			"char": chm.GetChar("abc"),
			"empty": chm.GetChar(""),
		}

		// Assert
		expected := args.Map{
			"char": byte('a'),
			"empty": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "GetChar -- first char or empty", actual)
	})
}

func Test_Seg6_CHM_GetCharOf(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetCharOf", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)

		// Act
		actual := args.Map{
			"char": chm.GetCharOf("abc"),
			"empty": chm.GetCharOf(""),
		}

		// Assert
		expected := args.Map{
			"char": byte('a'),
			"empty": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "GetCharOf -- first char or empty", actual)
	})
}

func Test_Seg6_CHM_Has(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Has", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")

		// Act
		actual := args.Map{
			"has": chm.Has("apple"),
			"miss": chm.Has("banana"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "Has -- found and missing", actual)
	})
}

func Test_Seg6_CHM_Has_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Has_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)

		// Act
		actual := args.Map{"has": chm.Has("apple")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Has empty -- false", actual)
	})
}

func Test_Seg6_CHM_HasWithHashset(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HasWithHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		has, hs := chm.HasWithHashset("apple")

		// Act
		actual := args.Map{
			"has": has,
			"notNil": hs != nil,
		}

		// Assert
		expected := args.Map{
			"has": true,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HasWithHashset -- found", actual)
	})
}

func Test_Seg6_CHM_HasWithHashset_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HasWithHashset_Miss", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		has, hs := chm.HasWithHashset("banana")

		// Act
		actual := args.Map{
			"has": has,
			"notNil": hs != nil,
		}

		// Assert
		expected := args.Map{
			"has": false,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HasWithHashset miss -- not found", actual)
	})
}

func Test_Seg6_CHM_HasWithHashset_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HasWithHashset_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		has, hs := chm.HasWithHashset("apple")

		// Act
		actual := args.Map{
			"has": has,
			"notNil": hs != nil,
		}

		// Assert
		expected := args.Map{
			"has": false,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HasWithHashset empty -- not found", actual)
	})
}

func Test_Seg6_CHM_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HasWithHashsetLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		has, hs := chm.HasWithHashsetLock("apple")

		// Act
		actual := args.Map{
			"has": has,
			"notNil": hs != nil,
		}

		// Assert
		expected := args.Map{
			"has": true,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock -- found", actual)
	})
}

func Test_Seg6_CHM_HasWithHashsetLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HasWithHashsetLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		has, hs := chm.HasWithHashsetLock("apple")

		// Act
		actual := args.Map{
			"has": has,
			"notNil": hs != nil,
		}

		// Assert
		expected := args.Map{
			"has": false,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock empty -- not found", actual)
	})
}

func Test_Seg6_CHM_HasWithHashsetLock_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HasWithHashsetLock_Miss", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		has, hs := chm.HasWithHashsetLock("banana")

		// Act
		actual := args.Map{
			"has": has,
			"notNil": hs != nil,
		}

		// Assert
		expected := args.Map{
			"has": false,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock miss -- not found", actual)
	})
}

func Test_Seg6_CHM_LengthOf(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_LengthOf", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("avocado")

		// Act
		actual := args.Map{
			"lenA": chm.LengthOf(byte('a')),
			"lenZ": chm.LengthOf(byte('z')),
		}

		// Assert
		expected := args.Map{
			"lenA": 2,
			"lenZ": 0,
		}
		expected.ShouldBeEqual(t, 0, "LengthOf -- 2 for a, 0 for z", actual)
	})
}

func Test_Seg6_CHM_LengthOf_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_LengthOf_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)

		// Act
		actual := args.Map{"len": chm.LengthOf(byte('a'))}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthOf empty -- 0", actual)
	})
}

func Test_Seg6_CHM_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_LengthOfLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")

		// Act
		actual := args.Map{
			"len": chm.LengthOfLock(byte('a')),
			"miss": chm.LengthOfLock(byte('z')),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"miss": 0,
		}
		expected.ShouldBeEqual(t, 0, "LengthOfLock -- found and missing", actual)
	})
}

func Test_Seg6_CHM_LengthOfLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_LengthOfLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)

		// Act
		actual := args.Map{"len": chm.LengthOfLock(byte('a'))}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthOfLock empty -- 0", actual)
	})
}

func Test_Seg6_CHM_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_LengthOfHashsetFromFirstChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("avocado")

		// Act
		actual := args.Map{
			"len": chm.LengthOfHashsetFromFirstChar("abc"),
			"miss": chm.LengthOfHashsetFromFirstChar("xyz"),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"miss": 0,
		}
		expected.ShouldBeEqual(t, 0, "LengthOfHashsetFromFirstChar -- 2 and 0", actual)
	})
}

func Test_Seg6_CHM_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AllLengthsSum", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("banana")

		// Act
		actual := args.Map{"sum": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AllLengthsSum -- 2", actual)
	})
}

func Test_Seg6_CHM_AllLengthsSum_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AllLengthsSum_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)

		// Act
		actual := args.Map{"sum": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"sum": 0}
		expected.ShouldBeEqual(t, 0, "AllLengthsSum empty -- 0", actual)
	})
}

func Test_Seg6_CHM_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AllLengthsSumLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")

		// Act
		actual := args.Map{"sum": chm.AllLengthsSumLock()}

		// Assert
		expected := args.Map{"sum": 1}
		expected.ShouldBeEqual(t, 0, "AllLengthsSumLock -- 1", actual)
	})
}

func Test_Seg6_CHM_AllLengthsSumLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AllLengthsSumLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)

		// Act
		actual := args.Map{"sum": chm.AllLengthsSumLock()}

		// Assert
		expected := args.Map{"sum": 0}
		expected.ShouldBeEqual(t, 0, "AllLengthsSumLock empty -- 0", actual)
	})
}

func Test_Seg6_CHM_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEmptyLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)

		// Act
		actual := args.Map{"empty": chm.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock -- true", actual)
	})
}

func Test_Seg6_CHM_LengthLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_LengthLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")

		// Act
		actual := args.Map{"len": chm.LengthLock()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock -- 1", actual)
	})
}

func Test_Seg6_CHM_GetMap(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetMap", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")

		// Act
		actual := args.Map{"notNil": chm.GetMap() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetMap -- not nil", actual)
	})
}

func Test_Seg6_CHM_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetCopyMapLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")

		// Act
		actual := args.Map{"len": len(chm.GetCopyMapLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetCopyMapLock -- 1", actual)
	})
}

func Test_Seg6_CHM_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetCopyMapLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)

		// Act
		actual := args.Map{"len": len(chm.GetCopyMapLock())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "GetCopyMapLock empty -- 0", actual)
	})
}

func Test_Seg6_CHM_List(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_List", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("banana")

		// Act
		actual := args.Map{"len": len(chm.List())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "List -- 2 items", actual)
	})
}

func Test_Seg6_CHM_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_SortedListAsc", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("banana").Add("apple")
		sorted := chm.SortedListAsc()

		// Act
		actual := args.Map{"first": sorted[0]}

		// Assert
		expected := args.Map{"first": "apple"}
		expected.ShouldBeEqual(t, 0, "SortedListAsc -- sorted", actual)
	})
}

func Test_Seg6_CHM_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_SortedListDsc", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("banana")
		sorted := chm.SortedListDsc()

		// Act
		actual := args.Map{"first": sorted[0]}

		// Assert
		expected := args.Map{"first": "banana"}
		expected.ShouldBeEqual(t, 0, "SortedListDsc -- descending", actual)
	})
}

func Test_Seg6_CHM_String(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_String", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")

		// Act
		actual := args.Map{"nonEmpty": chm.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg6_CHM_StringLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_StringLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")

		// Act
		actual := args.Map{"nonEmpty": chm.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock -- non-empty", actual)
	})
}

func Test_Seg6_CHM_SummaryString(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_SummaryString", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")

		// Act
		actual := args.Map{"nonEmpty": chm.SummaryString() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryString -- non-empty", actual)
	})
}

func Test_Seg6_CHM_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_SummaryStringLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")

		// Act
		actual := args.Map{"nonEmpty": chm.SummaryStringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringLock -- non-empty", actual)
	})
}

func Test_Seg6_CHM_Print(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Print", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		chm.Print(false)
		chm.Print(true)
	})
}

func Test_Seg6_CHM_PrintLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_PrintLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		chm.PrintLock(false)
		chm.PrintLock(true)
	})
}

// ── IsEquals ────────────────────────────────────────────────────────────────

func Test_Seg6_CHM_IsEquals(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEquals", func() {
		// Arrange
		chm1 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm2.Add("apple")

		// Act
		actual := args.Map{
			"eq":   chm1.IsEquals(chm2),
			"same": chm1.IsEquals(chm1),
			"nil":  chm1.IsEquals(nil),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"same": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEquals -- various", actual)
	})
}

func Test_Seg6_CHM_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEquals_BothEmpty", func() {
		// Arrange
		chm1 := corestr.New.CharHashsetMap.Cap(0, 4)
		chm2 := corestr.New.CharHashsetMap.Cap(0, 4)

		// Act
		actual := args.Map{"eq": chm1.IsEquals(chm2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty -- true", actual)
	})
}

func Test_Seg6_CHM_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEquals_OneEmpty", func() {
		// Arrange
		chm1 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(0, 4)

		// Act
		actual := args.Map{"eq": chm1.IsEquals(chm2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals one empty -- false", actual)
	})
}

func Test_Seg6_CHM_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEquals_DiffLen", func() {
		// Arrange
		chm1 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm2.Add("apple").Add("banana")

		// Act
		actual := args.Map{"eq": chm1.IsEquals(chm2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff len -- false", actual)
	})
}

func Test_Seg6_CHM_IsEquals_DiffItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEquals_DiffItems", func() {
		// Arrange
		chm1 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm2.Add("avocado")

		// Act
		actual := args.Map{"eq": chm1.IsEquals(chm2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff items -- false", actual)
	})
}

func Test_Seg6_CHM_IsEquals_MissingKey(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEquals_MissingKey", func() {
		// Arrange
		chm1 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm2.Add("banana")

		// Act
		actual := args.Map{"eq": chm1.IsEquals(chm2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals missing key -- false", actual)
	})
}

func Test_Seg6_CHM_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEqualsLock", func() {
		// Arrange
		chm1 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm2.Add("apple")

		// Act
		actual := args.Map{"eq": chm1.IsEqualsLock(chm2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsLock -- true", actual)
	})
}

// ── AddSameStartingCharItems ────────────────────────────────────────────────

func Test_Seg6_CHM_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameStartingCharItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddSameStartingCharItems(byte('a'), []string{"apple", "avocado"})

		// Act
		actual := args.Map{"sum": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems -- 2 items", actual)
	})
}

func Test_Seg6_CHM_AddSameStartingCharItems_Existing(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameStartingCharItems_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		chm.AddSameStartingCharItems(byte('a'), []string{"avocado"})

		// Act
		actual := args.Map{"sum": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems existing -- 2 items", actual)
	})
}

func Test_Seg6_CHM_AddSameStartingCharItems_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameStartingCharItems_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddSameStartingCharItems(byte('a'), []string{})

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems empty -- no change", actual)
	})
}

// ── AddCollectionItems / AddCharCollectionMapItems ──────────────────────────

func Test_Seg6_CHM_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddCollectionItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		c := corestr.New.Collection.Strings([]string{"apple", "banana"})
		chm.AddCollectionItems(c)

		// Act
		actual := args.Map{"sum": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AddCollectionItems -- 2 items", actual)
	})
}

func Test_Seg6_CHM_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddCollectionItems_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddCollectionItems(nil)

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollectionItems nil -- no change", actual)
	})
}

func Test_Seg6_CHM_AddCharCollectionMapItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddCharCollectionMapItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		chm.AddCharCollectionMapItems(ccm)

		// Act
		actual := args.Map{"has": chm.Has("apple")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddCharCollectionMapItems -- added", actual)
	})
}

func Test_Seg6_CHM_AddCharCollectionMapItems_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddCharCollectionMapItems_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddCharCollectionMapItems(nil)

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCharCollectionMapItems nil -- no change", actual)
	})
}

func Test_Seg6_CHM_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddHashsetItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := corestr.New.Hashset.Strings([]string{"apple", "banana"})
		chm.AddHashsetItems(hs)

		// Act
		actual := args.Map{"sum": chm.AllLengthsSum()}

		// Assert
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems -- 2 items", actual)
	})
}

func Test_Seg6_CHM_AddHashsetItems_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddHashsetItems_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := corestr.New.Hashset.Empty()
		chm.AddHashsetItems(hs)

		// Act
		actual := args.Map{"len": chm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems empty -- no change", actual)
	})
}

// ── GetHashset / GetHashsetLock ─────────────────────────────────────────────

func Test_Seg6_CHM_GetHashset(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := chm.GetHashset("a", false)

		// Act
		actual := args.Map{"notNil": hs != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashset -- found", actual)
	})
}

func Test_Seg6_CHM_GetHashset_AddNew(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetHashset_AddNew", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := chm.GetHashset("z", true)

		// Act
		actual := args.Map{"notNil": hs != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashset addNew -- created", actual)
	})
}

func Test_Seg6_CHM_GetHashset_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetHashset_Miss", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := chm.GetHashset("z", false)

		// Act
		actual := args.Map{"nil": hs == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "GetHashset miss -- nil", actual)
	})
}

func Test_Seg6_CHM_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetHashsetLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := chm.GetHashsetLock(false, "a")

		// Act
		actual := args.Map{"notNil": hs != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashsetLock -- found", actual)
	})
}

func Test_Seg6_CHM_GetHashsetByChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetHashsetByChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")

		// Act
		actual := args.Map{"notNil": chm.GetHashsetByChar(byte('a')) != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashsetByChar -- found", actual)
	})
}

// ── HashsetByChar / HashsetByStringFirstChar ────────────────────────────────

func Test_Seg6_CHM_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetByChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := chm.HashsetByChar(byte('a'))

		// Act
		actual := args.Map{"notNil": hs != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByChar -- found", actual)
	})
}

func Test_Seg6_CHM_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetByCharLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := chm.HashsetByCharLock(byte('a'))

		// Act
		actual := args.Map{"notNil": hs != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock -- found", actual)
	})
}

func Test_Seg6_CHM_HashsetByCharLock_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetByCharLock_Miss", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := chm.HashsetByCharLock(byte('z'))

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock miss -- empty", actual)
	})
}

func Test_Seg6_CHM_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetByStringFirstChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := chm.HashsetByStringFirstChar("avocado")

		// Act
		actual := args.Map{"notNil": hs != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstChar -- found", actual)
	})
}

func Test_Seg6_CHM_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetByStringFirstCharLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := chm.HashsetByStringFirstCharLock("avocado")

		// Act
		actual := args.Map{"notNil": hs != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstCharLock -- found", actual)
	})
}

// ── HashsetsCollection ──────────────────────────────────────────────────────

func Test_Seg6_CHM_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetsCollection", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("banana")
		hsc := chm.HashsetsCollection()

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection -- 2 hashsets", actual)
	})
}

func Test_Seg6_CHM_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetsCollection_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		hsc := chm.HashsetsCollection()

		// Act
		actual := args.Map{"empty": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection empty -- empty", actual)
	})
}

func Test_Seg6_CHM_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetsCollectionByChars", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("banana")
		hsc := chm.HashsetsCollectionByChars(byte('a'))

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars -- 1 hashset", actual)
	})
}

func Test_Seg6_CHM_HashsetsCollectionByChars_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetsCollectionByChars_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		hsc := chm.HashsetsCollectionByChars(byte('a'))

		// Act
		actual := args.Map{"empty": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars empty -- empty", actual)
	})
}

func Test_Seg6_CHM_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetsCollectionByStringsFirstChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hsc := chm.HashsetsCollectionByStringsFirstChar("avocado")

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringsFirstChar -- 1", actual)
	})
}

func Test_Seg6_CHM_HashsetsCollectionByStringsFirstChar_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetsCollectionByStringsFirstChar_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		hsc := chm.HashsetsCollectionByStringsFirstChar("a")

		// Act
		actual := args.Map{"empty": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringsFirstChar empty -- empty", actual)
	})
}

// ── GetCharsGroups ──────────────────────────────────────────────────────────

func Test_Seg6_CHM_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetCharsGroups", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		result := chm.GetCharsGroups("apple", "banana")

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetCharsGroups -- 2 groups", actual)
	})
}

func Test_Seg6_CHM_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetCharsGroups_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		result := chm.GetCharsGroups()

		// Act
		actual := args.Map{"same": result == chm}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "GetCharsGroups empty -- returns self", actual)
	})
}

// ── AddSameCharsCollection / AddSameCharsHashset ────────────────────────────

func Test_Seg6_CHM_AddSameCharsCollection_Existing(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollection_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		c := corestr.New.Collection.Strings([]string{"avocado"})
		result := chm.AddSameCharsCollection("a", c)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection existing -- merged", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollection_NilColl(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollection_NilColl", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		result := chm.AddSameCharsCollection("a", nil)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection nil -- returns existing", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollection_NewChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollection_NewChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		c := corestr.New.Collection.Strings([]string{"banana"})
		result := chm.AddSameCharsCollection("b", c)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection new -- added", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollection_NewNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollection_NewNil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		result := chm.AddSameCharsCollection("z", nil)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection new nil -- created empty", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsHashset_Existing(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsHashset_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := corestr.New.Hashset.Strings([]string{"avocado"})
		result := chm.AddSameCharsHashset("a", hs)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset existing -- merged", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsHashset_NewChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsHashset_NewChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := corestr.New.Hashset.Strings([]string{"banana"})
		result := chm.AddSameCharsHashset("b", hs)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset new -- added", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsHashset_NewNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsHashset_NewNil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		result := chm.AddSameCharsHashset("z", nil)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset new nil -- created empty", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsHashset_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsHashset_ExistingNil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		result := chm.AddSameCharsHashset("a", nil)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset existing nil -- returns existing", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollectionLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		c := corestr.New.Collection.Strings([]string{"avocado"})
		result := chm.AddSameCharsCollectionLock("a", c)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock -- merged", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollectionLock_NewNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollectionLock_NewNil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		result := chm.AddSameCharsCollectionLock("z", nil)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock new nil -- created", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollectionLock_NewChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollectionLock_NewChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		c := corestr.New.Collection.Strings([]string{"banana"})
		result := chm.AddSameCharsCollectionLock("b", c)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock new char -- added", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollectionLock_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollectionLock_ExistingNil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		result := chm.AddSameCharsCollectionLock("a", nil)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock existing nil -- returned", actual)
	})
}

func Test_Seg6_CHM_AddHashsetLock_Existing(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddHashsetLock_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := corestr.New.Hashset.Strings([]string{"avocado"})
		result := chm.AddHashsetLock("a", hs)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock existing -- merged", actual)
	})
}

func Test_Seg6_CHM_AddHashsetLock_NewChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddHashsetLock_NewChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := corestr.New.Hashset.Strings([]string{"banana"})
		result := chm.AddHashsetLock("b", hs)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock new char -- added", actual)
	})
}

func Test_Seg6_CHM_AddHashsetLock_NewNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddHashsetLock_NewNil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		result := chm.AddHashsetLock("z", nil)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock new nil -- created empty", actual)
	})
}

func Test_Seg6_CHM_AddHashsetLock_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddHashsetLock_ExistingNil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		result := chm.AddHashsetLock("a", nil)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock existing nil -- returned", actual)
	})
}

// ── JSON ────────────────────────────────────────────────────────────────────

func Test_Seg6_CHM_Json(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Json", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		j := chm.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg6_CHM_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_MarshalJSON", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		b, err := chm.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg6_CHM_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_UnmarshalJSON", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		b, _ := chm.MarshalJSON()
		chm2 := corestr.New.CharHashsetMap.Cap(0, 4)
		err := chm2.UnmarshalJSON(b)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Seg6_CHM_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_UnmarshalJSON_Invalid", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		err := chm.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Seg6_CHM_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_JsonModel", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")

		// Act
		actual := args.Map{"notNil": chm.JsonModel() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModel -- non-nil", actual)
	})
}

func Test_Seg6_CHM_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_JsonModelAny", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)

		// Act
		actual := args.Map{"notNil": chm.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Seg6_CHM_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_InterfaceCasts", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)

		// Act
		actual := args.Map{
			"jsoner":   chm.AsJsoner() != nil,
			"binder":   chm.AsJsonContractsBinder() != nil,
			"injector": chm.AsJsonParseSelfInjector() != nil,
			"marsh":    chm.AsJsonMarshaller() != nil,
		}

		// Assert
		expected := args.Map{
			"jsoner": true,
			"binder": true,
			"injector": true,
			"marsh": true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_Seg6_CHM_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_ParseInjectUsingJson", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.New.CharHashsetMap.Cap(0, 4)
		_, err := chm2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_Seg6_CHM_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_ParseInjectUsingJsonMust", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.New.CharHashsetMap.Cap(0, 4)
		result := chm2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg6_CHM_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_JsonParseSelfInject", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.New.CharHashsetMap.Cap(0, 4)
		err := chm2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ── Clear / RemoveAll ───────────────────────────────────────────────────────

func Test_Seg6_CHM_Clear(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Clear", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		chm.Clear()

		// Act
		actual := args.Map{"empty": chm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg6_CHM_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Clear_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		result := chm.Clear()

		// Act
		actual := args.Map{"same": result == chm}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Clear empty -- returns self", actual)
	})
}

func Test_Seg6_CHM_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_RemoveAll", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		chm.RemoveAll()

		// Act
		actual := args.Map{"empty": chm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "RemoveAll -- emptied", actual)
	})
}

func Test_Seg6_CHM_RemoveAll_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_RemoveAll_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		result := chm.RemoveAll()

		// Act
		actual := args.Map{"same": result == chm}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "RemoveAll empty -- returns self", actual)
	})
}
