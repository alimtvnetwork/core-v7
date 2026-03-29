package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

type caseV1Compat struct {
	Name     string
	Expected any
	Actual   any
	Args     args.Map
}

func (it caseV1Compat) ShouldBeEqual(t *testing.T) {
	expected := args.Map{"value": it.Expected}
	actual := args.Map{"value": it.Actual}
	expected.ShouldBeEqual(t, 0, it.Name, actual)
}

// ─── CharCollectionMap: IsEmpty / HasItems / Length ──────────────

func Test_Cov41_CharCollectionMap_IsEmpty_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_IsEmpty_Empty", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "Empty CharCollectionMap IsEmpty",
			Expected: true,
			Actual:   m.IsEmpty(),
			Args:     args.Map{},
		}
		// Act & Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_IsEmpty_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_IsEmpty_NonEmpty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		tc := caseV1Compat{
			Name:     "NonEmpty CharCollectionMap IsEmpty",
			Expected: false,
			Actual:   m.IsEmpty(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HasItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HasItems_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "Empty CharCollectionMap HasItems",
			Expected: false,
			Actual:   m.HasItems(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HasItems_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HasItems_NonEmpty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("world")
		tc := caseV1Compat{
			Name:     "NonEmpty CharCollectionMap HasItems",
			Expected: true,
			Actual:   m.HasItems(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_Length_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Length_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "Empty CharCollectionMap Length",
			Expected: 0,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_Length_WithItems(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Length_WithItems", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("apple")
		m.Add("banana")
		tc := caseV1Compat{
			Name:     "CharCollectionMap Length with 2 chars",
			Expected: 2,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_IsEmptyLock", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "IsEmptyLock on empty",
			Expected: true,
			Actual:   m.IsEmptyLock(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_LengthLock", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("cat")
		tc := caseV1Compat{
			Name:     "LengthLock with 1 item",
			Expected: 1,
			Actual:   m.LengthLock(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── Add / AddStrings / AddLock ──────────────

func Test_Cov41_CharCollectionMap_Add_SameChar(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Add_SameChar", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("apple")
		m.Add("avocado")
		tc := caseV1Compat{
			Name:     "Add same starting char groups together",
			Expected: 1,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_Add_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Add_AllLengthsSum", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("apple")
		m.Add("avocado")
		m.Add("banana")
		tc := caseV1Compat{
			Name:     "AllLengthsSum after 3 adds",
			Expected: 3,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddStrings_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddStrings()
		tc := caseV1Compat{
			Name:     "AddStrings with no args",
			Expected: 0,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddStrings_Multiple(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddStrings_Multiple", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddStrings("apple", "banana", "cherry")
		tc := caseV1Compat{
			Name:     "AddStrings adds 3 different chars",
			Expected: 3,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddLock", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddLock("hello")
		tc := caseV1Compat{
			Name:     "AddLock adds item",
			Expected: 1,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddLock_ExistingChar(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddLock_ExistingChar", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddLock("hello")
		m.AddLock("happy")
		tc := caseV1Compat{
			Name:     "AddLock with existing char",
			Expected: 2,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── GetChar / Has / HasWithCollection ──────────────

func Test_Cov41_CharCollectionMap_GetChar_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_GetChar_NonEmpty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "GetChar returns first byte",
			Expected: byte('h'),
			Actual:   m.GetChar("hello"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_GetChar_EmptyString(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_GetChar_EmptyString", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "GetChar on empty string returns 0",
			Expected: byte(0),
			Actual:   m.GetChar(""),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_Has_Found(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Has_Found", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		tc := caseV1Compat{
			Name:     "Has finds existing item",
			Expected: true,
			Actual:   m.Has("hello"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_Has_NotFound(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Has_NotFound", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		tc := caseV1Compat{
			Name:     "Has returns false for missing",
			Expected: false,
			Actual:   m.Has("world"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_Has_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Has_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "Has on empty returns false",
			Expected: false,
			Actual:   m.Has("anything"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HasWithCollection_Found(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HasWithCollection_Found", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		found, col := m.HasWithCollection("hello")
		tc := caseV1Compat{
			Name:     "HasWithCollection found",
			Expected: true,
			Actual:   found && col.HasItems(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HasWithCollection_NotFound(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HasWithCollection_NotFound", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		found, _ := m.HasWithCollection("world")
		tc := caseV1Compat{
			Name:     "HasWithCollection not found",
			Expected: false,
			Actual:   found,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HasWithCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HasWithCollection_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		found, _ := m.HasWithCollection("anything")
		tc := caseV1Compat{
			Name:     "HasWithCollection on empty",
			Expected: false,
			Actual:   found,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HasWithCollectionLock_Found(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HasWithCollectionLock_Found", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		found, col := m.HasWithCollectionLock("hello")
		tc := caseV1Compat{
			Name:     "HasWithCollectionLock found",
			Expected: true,
			Actual:   found && col.HasItems(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HasWithCollectionLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HasWithCollectionLock_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		found, _ := m.HasWithCollectionLock("anything")
		tc := caseV1Compat{
			Name:     "HasWithCollectionLock on empty",
			Expected: false,
			Actual:   found,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HasWithCollectionLock_MissingChar(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HasWithCollectionLock_MissingChar", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		found, _ := m.HasWithCollectionLock("world")
		tc := caseV1Compat{
			Name:     "HasWithCollectionLock missing char",
			Expected: false,
			Actual:   found,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── LengthOf / LengthOfLock / LengthOfCollectionFromFirstChar ──────────────

func Test_Cov41_CharCollectionMap_LengthOf_Exists(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_LengthOf_Exists", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		m.Add("happy")
		tc := caseV1Compat{
			Name:     "LengthOf existing char",
			Expected: 2,
			Actual:   m.LengthOf(byte('h')),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_LengthOf_Missing(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_LengthOf_Missing", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "LengthOf missing char",
			Expected: 0,
			Actual:   m.LengthOf(byte('x')),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_LengthOfLock_Exists(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_LengthOfLock_Exists", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		tc := caseV1Compat{
			Name:     "LengthOfLock existing char",
			Expected: 1,
			Actual:   m.LengthOfLock(byte('h')),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_LengthOfLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_LengthOfLock_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "LengthOfLock on empty",
			Expected: 0,
			Actual:   m.LengthOfLock(byte('x')),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_LengthOfCollectionFromFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_LengthOfCollectionFromFirstChar", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		m.Add("happy")
		tc := caseV1Compat{
			Name:     "LengthOfCollectionFromFirstChar",
			Expected: 2,
			Actual:   m.LengthOfCollectionFromFirstChar("hi"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_LengthOfCollectionFromFirstChar_Missing(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_LengthOfCollectionFromFirstChar_Missing", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "LengthOfCollectionFromFirstChar missing",
			Expected: 0,
			Actual:   m.LengthOfCollectionFromFirstChar("zzz"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── AllLengthsSum / AllLengthsSumLock ──────────────

func Test_Cov41_CharCollectionMap_AllLengthsSum_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AllLengthsSum_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "AllLengthsSum empty",
			Expected: 0,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AllLengthsSumLock", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("apple")
		m.Add("banana")
		tc := caseV1Compat{
			Name:     "AllLengthsSumLock",
			Expected: 2,
			Actual:   m.AllLengthsSumLock(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── List / ListLock / SortedListAsc ──────────────

func Test_Cov41_CharCollectionMap_List_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_List_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "List on empty",
			Expected: 0,
			Actual:   len(m.List()),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_List_WithItems(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_List_WithItems", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("apple")
		m.Add("banana")
		tc := caseV1Compat{
			Name:     "List returns all items",
			Expected: 2,
			Actual:   len(m.List()),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_ListLock(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_ListLock", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("x")
		tc := caseV1Compat{
			Name:     "ListLock returns items",
			Expected: 1,
			Actual:   len(m.ListLock()),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_SortedListAsc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_SortedListAsc_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "SortedListAsc empty",
			Expected: 0,
			Actual:   len(m.SortedListAsc()),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_SortedListAsc_Sorted(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_SortedListAsc_Sorted", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("banana")
		m.Add("apple")
		sorted := m.SortedListAsc()
		tc := caseV1Compat{
			Name:     "SortedListAsc sorts items",
			Expected: "apple",
			Actual:   sorted[0],
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── GetMap / GetCopyMapLock / GetCollection / GetCollectionLock ──────────────

func Test_Cov41_CharCollectionMap_GetMap(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_GetMap", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("x")
		tc := caseV1Compat{
			Name:     "GetMap returns underlying map",
			Expected: 1,
			Actual:   len(m.GetMap()),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_GetCopyMapLock_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "GetCopyMapLock on empty",
			Expected: 0,
			Actual:   len(m.GetCopyMapLock()),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_GetCopyMapLock_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_GetCopyMapLock_NonEmpty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("abc")
		tc := caseV1Compat{
			Name:     "GetCopyMapLock non-empty",
			Expected: 1,
			Actual:   len(m.GetCopyMapLock()),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_GetCollection_Exists(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_GetCollection_Exists", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		col := m.GetCollection("hi", false)
		tc := caseV1Compat{
			Name:     "GetCollection existing",
			Expected: true,
			Actual:   col != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_GetCollection_Missing_NoAdd(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_GetCollection_Missing_NoAdd", func() {
		m := corestr.New.CharCollectionMap.Empty()
		col := m.GetCollection("z", false)
		tc := caseV1Compat{
			Name:     "GetCollection missing no add",
			Expected: true,
			Actual:   col == nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_GetCollection_Missing_AddNew(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_GetCollection_Missing_AddNew", func() {
		m := corestr.New.CharCollectionMap.Empty()
		col := m.GetCollection("z", true)
		tc := caseV1Compat{
			Name:     "GetCollection missing with add new",
			Expected: true,
			Actual:   col != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_GetCollectionLock(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_GetCollectionLock", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("test")
		col := m.GetCollectionLock("testing", false)
		tc := caseV1Compat{
			Name:     "GetCollectionLock existing",
			Expected: true,
			Actual:   col != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_GetCollectionByChar(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_GetCollectionByChar", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		col := m.GetCollectionByChar(byte('h'))
		tc := caseV1Compat{
			Name:     "GetCollectionByChar existing",
			Expected: true,
			Actual:   col != nil && col.HasItems(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── AddSameStartingCharItems ──────────────

func Test_Cov41_CharCollectionMap_AddSameStartingCharItems_New(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddSameStartingCharItems_New", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddSameStartingCharItems(byte('a'), []string{"apple", "avocado"}, false)
		tc := caseV1Compat{
			Name:     "AddSameStartingCharItems new char",
			Expected: 2,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddSameStartingCharItems_Existing(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddSameStartingCharItems_Existing", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("apple")
		m.AddSameStartingCharItems(byte('a'), []string{"avocado"}, false)
		tc := caseV1Compat{
			Name:     "AddSameStartingCharItems existing char",
			Expected: 2,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddSameStartingCharItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddSameStartingCharItems_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddSameStartingCharItems(byte('a'), []string{}, false)
		tc := caseV1Compat{
			Name:     "AddSameStartingCharItems empty slice",
			Expected: 0,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── AddSameCharsCollection / AddSameCharsCollectionLock ──────────────

func Test_Cov41_CharCollectionMap_AddSameCharsCollection_New(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddSameCharsCollection_New", func() {
		m := corestr.New.CharCollectionMap.Empty()
		col := corestr.New.Collection.Strings([]string{"apple", "avocado"})
		m.AddSameCharsCollection("apple", col)
		tc := caseV1Compat{
			Name:     "AddSameCharsCollection new",
			Expected: 2,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddSameCharsCollection_Existing(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddSameCharsCollection_Existing", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("apple")
		col := corestr.New.Collection.Strings([]string{"avocado"})
		m.AddSameCharsCollection("abc", col)
		tc := caseV1Compat{
			Name:     "AddSameCharsCollection existing char",
			Expected: 2,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddSameCharsCollection_NilCollection(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddSameCharsCollection_NilCollection", func() {
		m := corestr.New.CharCollectionMap.Empty()
		result := m.AddSameCharsCollection("abc", nil)
		tc := caseV1Compat{
			Name:     "AddSameCharsCollection nil creates empty collection",
			Expected: true,
			Actual:   result != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddSameCharsCollection_ExistingNilCol(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddSameCharsCollection_ExistingNilCol", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("apple")
		result := m.AddSameCharsCollection("abc", nil)
		tc := caseV1Compat{
			Name:     "AddSameCharsCollection existing char nil col returns existing",
			Expected: true,
			Actual:   result != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddSameCharsCollectionLock_New(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddSameCharsCollectionLock_New", func() {
		m := corestr.New.CharCollectionMap.Empty()
		col := corestr.New.Collection.Strings([]string{"banana"})
		m.AddSameCharsCollectionLock("bbb", col)
		tc := caseV1Compat{
			Name:     "AddSameCharsCollectionLock new",
			Expected: 1,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddSameCharsCollectionLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddSameCharsCollectionLock_Existing", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("banana")
		col := corestr.New.Collection.Strings([]string{"berry"})
		m.AddSameCharsCollectionLock("bbb", col)
		tc := caseV1Compat{
			Name:     "AddSameCharsCollectionLock existing",
			Expected: 2,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddSameCharsCollectionLock_NilCol(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddSameCharsCollectionLock_NilCol", func() {
		m := corestr.New.CharCollectionMap.Empty()
		result := m.AddSameCharsCollectionLock("bbb", nil)
		tc := caseV1Compat{
			Name:     "AddSameCharsCollectionLock nil col",
			Expected: true,
			Actual:   result != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddSameCharsCollectionLock_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddSameCharsCollectionLock_ExistingNil", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("banana")
		result := m.AddSameCharsCollectionLock("bbb", nil)
		tc := caseV1Compat{
			Name:     "AddSameCharsCollectionLock existing nil",
			Expected: true,
			Actual:   result != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── AddCollectionItems / AddHashmapsValues / AddHashmapsKeysValuesBoth ──────

func Test_Cov41_CharCollectionMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddCollectionItems_Nil", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddCollectionItems(nil)
		tc := caseV1Compat{
			Name:     "AddCollectionItems nil",
			Expected: 0,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddCollectionItems_Valid(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddCollectionItems_Valid", func() {
		m := corestr.New.CharCollectionMap.Empty()
		col := corestr.New.Collection.Strings([]string{"alpha", "beta"})
		m.AddCollectionItems(col)
		tc := caseV1Compat{
			Name:     "AddCollectionItems valid",
			Expected: 2,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddHashmapsValues_Nil(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddHashmapsValues_Nil", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddHashmapsValues(nil)
		tc := caseV1Compat{
			Name:     "AddHashmapsValues nil",
			Expected: 0,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddHashmapsValues_Valid(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddHashmapsValues_Valid", func() {
		m := corestr.New.CharCollectionMap.Empty()
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "val1"})
		m.AddHashmapsValues(hm)
		tc := caseV1Compat{
			Name:     "AddHashmapsValues valid",
			Expected: true,
			Actual:   m.Has("val1"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddHashmapsKeysValuesBoth_Nil(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddHashmapsKeysValuesBoth_Nil", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddHashmapsKeysValuesBoth(nil)
		tc := caseV1Compat{
			Name:     "AddHashmapsKeysValuesBoth nil",
			Expected: 0,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddHashmapsKeysValuesBoth_Valid(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddHashmapsKeysValuesBoth_Valid", func() {
		m := corestr.New.CharCollectionMap.Empty()
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "key", Value: "val"})
		m.AddHashmapsKeysValuesBoth(hm)
		tc := caseV1Compat{
			Name:     "AddHashmapsKeysValuesBoth adds both key and value",
			Expected: 2,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── AddHashmapsKeysOrValuesBothUsingFilter ──────

func Test_Cov41_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Nil", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddHashmapsKeysOrValuesBothUsingFilter(nil, nil)
		tc := caseV1Compat{
			Name:     "AddHashmapsKeysOrValuesBothUsingFilter nil",
			Expected: 0,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Accept(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Accept", func() {
		m := corestr.New.CharCollectionMap.Empty()
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Value, true, false
		}
		m.AddHashmapsKeysOrValuesBothUsingFilter(filter, hm)
		tc := caseV1Compat{
			Name:     "AddHashmapsKeysOrValuesBothUsingFilter accept",
			Expected: true,
			Actual:   m.Has("v1"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Break", func() {
		m := corestr.New.CharCollectionMap.Empty()
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Value, false, true
		}
		m.AddHashmapsKeysOrValuesBothUsingFilter(filter, hm)
		tc := caseV1Compat{
			Name:     "AddHashmapsKeysOrValuesBothUsingFilter break",
			Expected: 0,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── AddCharHashsetMap ──────

func Test_Cov41_CharCollectionMap_AddCharHashsetMap(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddCharHashsetMap", func() {
		m := corestr.New.CharCollectionMap.Empty()
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		m.AddCharHashsetMap(chm)
		tc := caseV1Compat{
			Name:     "AddCharHashsetMap",
			Expected: true,
			Actual:   m.Has("hello"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── IsEquals / IsEqualsCaseSensitive / IsEqualsLock ──────────────

func Test_Cov41_CharCollectionMap_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_IsEquals_Same", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		tc := caseV1Compat{
			Name:     "IsEquals same pointer",
			Expected: true,
			Actual:   m.IsEquals(m),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_IsEquals_Nil", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "IsEquals nil",
			Expected: false,
			Actual:   m.IsEquals(nil),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_IsEquals_BothEmpty", func() {
		m1 := corestr.New.CharCollectionMap.Empty()
		m2 := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "IsEquals both empty",
			Expected: true,
			Actual:   m1.IsEquals(m2),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_IsEquals_DiffLen", func() {
		m1 := corestr.New.CharCollectionMap.Empty()
		m1.Add("a")
		m2 := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "IsEquals diff length",
			Expected: false,
			Actual:   m1.IsEquals(m2),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_IsEquals_DiffContent(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_IsEquals_DiffContent", func() {
		m1 := corestr.New.CharCollectionMap.Empty()
		m1.Add("apple")
		m2 := corestr.New.CharCollectionMap.Empty()
		m2.Add("avocado")
		tc := caseV1Compat{
			Name:     "IsEquals diff content same char",
			Expected: false,
			Actual:   m1.IsEquals(m2),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_IsEquals_MissingKey(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_IsEquals_MissingKey", func() {
		m1 := corestr.New.CharCollectionMap.Empty()
		m1.Add("apple")
		m2 := corestr.New.CharCollectionMap.Empty()
		m2.Add("banana")
		tc := caseV1Compat{
			Name:     "IsEquals missing key",
			Expected: false,
			Actual:   m1.IsEquals(m2),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_IsEqualsLock", func() {
		m1 := corestr.New.CharCollectionMap.Empty()
		m1.Add("x")
		m2 := corestr.New.CharCollectionMap.Empty()
		m2.Add("x")
		tc := caseV1Compat{
			Name:     "IsEqualsLock equal",
			Expected: true,
			Actual:   m1.IsEqualsLock(m2),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_IsEqualsCaseSensitive_Insensitive(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_IsEqualsCaseSensitive_Insensitive", func() {
		m1 := corestr.New.CharCollectionMap.Empty()
		m1.Add("Hello")
		m2 := corestr.New.CharCollectionMap.Empty()
		m2.Add("Hello")
		tc := caseV1Compat{
			Name:     "IsEqualsCaseSensitive false",
			Expected: true,
			Actual:   m1.IsEqualsCaseSensitive(false, m2),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_IsEqualsCaseSensitiveLock(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_IsEqualsCaseSensitiveLock", func() {
		m1 := corestr.New.CharCollectionMap.Empty()
		m1.Add("test")
		m2 := corestr.New.CharCollectionMap.Empty()
		m2.Add("test")
		tc := caseV1Compat{
			Name:     "IsEqualsCaseSensitiveLock",
			Expected: true,
			Actual:   m1.IsEqualsCaseSensitiveLock(true, m2),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── Hashset methods ──────────────

func Test_Cov41_CharCollectionMap_HashsetByChar_Found(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HashsetByChar_Found", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		hs := m.HashsetByChar(byte('h'))
		tc := caseV1Compat{
			Name:     "HashsetByChar found",
			Expected: true,
			Actual:   hs != nil && hs.Has("hello"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HashsetByChar_Missing(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HashsetByChar_Missing", func() {
		m := corestr.New.CharCollectionMap.Empty()
		hs := m.HashsetByChar(byte('z'))
		tc := caseV1Compat{
			Name:     "HashsetByChar missing",
			Expected: true,
			Actual:   hs == nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HashsetByCharLock_Found(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HashsetByCharLock_Found", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		hs := m.HashsetByCharLock(byte('h'))
		tc := caseV1Compat{
			Name:     "HashsetByCharLock found",
			Expected: true,
			Actual:   hs != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HashsetByCharLock_Missing(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HashsetByCharLock_Missing", func() {
		m := corestr.New.CharCollectionMap.Empty()
		hs := m.HashsetByCharLock(byte('z'))
		tc := caseV1Compat{
			Name:     "HashsetByCharLock missing returns empty",
			Expected: true,
			Actual:   hs != nil && hs.IsEmpty(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HashsetByStringFirstChar", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		hs := m.HashsetByStringFirstChar("hi")
		tc := caseV1Compat{
			Name:     "HashsetByStringFirstChar",
			Expected: true,
			Actual:   hs != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HashsetByStringFirstCharLock", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		hs := m.HashsetByStringFirstCharLock("hi")
		tc := caseV1Compat{
			Name:     "HashsetByStringFirstCharLock",
			Expected: true,
			Actual:   hs != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HashsetsCollection_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		hsc := m.HashsetsCollection()
		tc := caseV1Compat{
			Name:     "HashsetsCollection empty",
			Expected: true,
			Actual:   hsc.IsEmpty(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HashsetsCollection_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HashsetsCollection_NonEmpty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		m.Add("banana")
		hsc := m.HashsetsCollection()
		tc := caseV1Compat{
			Name:     "HashsetsCollection non-empty",
			Expected: 2,
			Actual:   hsc.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HashsetsCollectionByChars_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HashsetsCollectionByChars_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		hsc := m.HashsetsCollectionByChars(byte('a'))
		tc := caseV1Compat{
			Name:     "HashsetsCollectionByChars empty",
			Expected: true,
			Actual:   hsc.IsEmpty(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HashsetsCollectionByChars_Valid(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HashsetsCollectionByChars_Valid", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("apple")
		hsc := m.HashsetsCollectionByChars(byte('a'))
		tc := caseV1Compat{
			Name:     "HashsetsCollectionByChars valid",
			Expected: 1,
			Actual:   hsc.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HashsetsCollectionByStringFirstChar_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HashsetsCollectionByStringFirstChar_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		hsc := m.HashsetsCollectionByStringFirstChar("apple")
		tc := caseV1Compat{
			Name:     "HashsetsCollectionByStringFirstChar empty",
			Expected: true,
			Actual:   hsc.IsEmpty(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_HashsetsCollectionByStringFirstChar_Valid(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_HashsetsCollectionByStringFirstChar_Valid", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("apple")
		hsc := m.HashsetsCollectionByStringFirstChar("abc")
		tc := caseV1Compat{
			Name:     "HashsetsCollectionByStringFirstChar valid",
			Expected: 1,
			Actual:   hsc.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── Resize / AddLength ──────────────

func Test_Cov41_CharCollectionMap_Resize_Larger(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Resize_Larger", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("x")
		m.Resize(100)
		tc := caseV1Compat{
			Name:     "Resize larger",
			Expected: true,
			Actual:   m.Has("x"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_Resize_AlreadyLarger(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Resize_AlreadyLarger", func() {
		m := corestr.New.CharCollectionMap.CapSelfCap(100, 10)
		m.Resize(5)
		tc := caseV1Compat{
			Name:     "Resize already larger",
			Expected: 0,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddLength(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddLength", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddLength(50)
		tc := caseV1Compat{
			Name:     "AddLength",
			Expected: 0,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AddLength_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AddLength_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.AddLength()
		tc := caseV1Compat{
			Name:     "AddLength no args",
			Expected: 0,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── GetCharsGroups ──────────────

func Test_Cov41_CharCollectionMap_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_GetCharsGroups_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		result := m.GetCharsGroups([]string{})
		tc := caseV1Compat{
			Name:     "GetCharsGroups empty",
			Expected: 0,
			Actual:   result.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_GetCharsGroups_Valid(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_GetCharsGroups_Valid", func() {
		m := corestr.New.CharCollectionMap.Empty()
		result := m.GetCharsGroups([]string{"apple", "banana", "avocado"})
		tc := caseV1Compat{
			Name:     "GetCharsGroups valid",
			Expected: 3,
			Actual:   result.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── String / SummaryString / Print / PrintLock ──────────────

func Test_Cov41_CharCollectionMap_String(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_String", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		s := m.String()
		tc := caseV1Compat{
			Name:     "String non-empty",
			Expected: true,
			Actual:   len(s) > 0,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_StringLock", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		s := m.StringLock()
		tc := caseV1Compat{
			Name:     "StringLock non-empty",
			Expected: true,
			Actual:   len(s) > 0,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_SummaryString", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("test")
		s := m.SummaryString()
		tc := caseV1Compat{
			Name:     "SummaryString",
			Expected: true,
			Actual:   len(s) > 0,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_SummaryStringLock", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("test")
		s := m.SummaryStringLock()
		tc := caseV1Compat{
			Name:     "SummaryStringLock",
			Expected: true,
			Actual:   len(s) > 0,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_Print_Skip(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Print_Skip", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Print(false) // should not panic
		tc := caseV1Compat{
			Name:     "Print skip",
			Expected: true,
			Actual:   true,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_PrintLock_Skip(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_PrintLock_Skip", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.PrintLock(false) // should not panic
		tc := caseV1Compat{
			Name:     "PrintLock skip",
			Expected: true,
			Actual:   true,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── JSON methods ──────────────

func Test_Cov41_CharCollectionMap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_MarshalJSON", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("test")
		bytes, err := m.MarshalJSON()
		tc := caseV1Compat{
			Name:     "MarshalJSON success",
			Expected: true,
			Actual:   err == nil && len(bytes) > 0,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_UnmarshalJSON", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("test")
		bytes, _ := m.MarshalJSON()
		m2 := corestr.New.CharCollectionMap.Empty()
		err := m2.UnmarshalJSON(bytes)
		tc := caseV1Compat{
			Name:     "UnmarshalJSON success",
			Expected: true,
			Actual:   err == nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_Json(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Json", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("x")
		j := m.Json()
		tc := caseV1Compat{
			Name:     "Json returns result",
			Expected: true,
			Actual:   j.Error == nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_JsonPtr", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("x")
		j := m.JsonPtr()
		tc := caseV1Compat{
			Name:     "JsonPtr returns non-nil",
			Expected: true,
			Actual:   j != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_JsonModel", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("x")
		model := m.JsonModel()
		tc := caseV1Compat{
			Name:     "JsonModel non-nil",
			Expected: true,
			Actual:   model != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_JsonModelAny", func() {
		m := corestr.New.CharCollectionMap.Empty()
		result := m.JsonModelAny()
		tc := caseV1Compat{
			Name:     "JsonModelAny non-nil",
			Expected: true,
			Actual:   result != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_ParseInjectUsingJson", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		j := m.JsonPtr()
		m2 := corestr.New.CharCollectionMap.Empty()
		result, err := m2.ParseInjectUsingJson(j)
		tc := caseV1Compat{
			Name:     "ParseInjectUsingJson success",
			Expected: true,
			Actual:   err == nil && result != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_ParseInjectUsingJson_Error", func() {
		m := corestr.New.CharCollectionMap.Empty()
		badJson := corejson.NewPtr("invalid")
		_, err := m.ParseInjectUsingJson(badJson)
		tc := caseV1Compat{
			Name:     "ParseInjectUsingJson error",
			Expected: true,
			Actual:   err != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_ParseInjectUsingJsonMust", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		j := m.JsonPtr()
		m2 := corestr.New.CharCollectionMap.Empty()
		result := m2.ParseInjectUsingJsonMust(j)
		tc := caseV1Compat{
			Name:     "ParseInjectUsingJsonMust success",
			Expected: true,
			Actual:   result != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_JsonParseSelfInject", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		j := m.JsonPtr()
		m2 := corestr.New.CharCollectionMap.Empty()
		err := m2.JsonParseSelfInject(j)
		tc := caseV1Compat{
			Name:     "JsonParseSelfInject success",
			Expected: true,
			Actual:   err == nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── Interface casts ──────────────

func Test_Cov41_CharCollectionMap_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AsJsoner", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "AsJsoner non-nil",
			Expected: true,
			Actual:   m.AsJsoner() != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AsJsonContractsBinder", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "AsJsonContractsBinder non-nil",
			Expected: true,
			Actual:   m.AsJsonContractsBinder() != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AsJsonMarshaller", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "AsJsonMarshaller non-nil",
			Expected: true,
			Actual:   m.AsJsonMarshaller() != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_AsJsonParseSelfInjector", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "AsJsonParseSelfInjector non-nil",
			Expected: true,
			Actual:   m.AsJsonParseSelfInjector() != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── Clear / Dispose ──────────────

func Test_Cov41_CharCollectionMap_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Clear_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Clear()
		tc := caseV1Compat{
			Name:     "Clear on empty",
			Expected: 0,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_Clear_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Clear_NonEmpty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		m.Clear()
		tc := caseV1Compat{
			Name:     "Clear on non-empty",
			Expected: 0,
			Actual:   m.Length(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_CharCollectionMap_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov41_CharCollectionMap_Dispose", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		m.Dispose()
		tc := caseV1Compat{
			Name:     "Dispose",
			Expected: true,
			Actual:   true, // no panic
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── newCharCollectionMapCreator ──────────────

func Test_Cov41_newCharCollectionMapCreator_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_newCharCollectionMapCreator_Empty", func() {
		m := corestr.New.CharCollectionMap.Empty()
		tc := caseV1Compat{
			Name:     "Creator Empty",
			Expected: true,
			Actual:   m != nil && m.IsEmpty(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_newCharCollectionMapCreator_CapSelfCap(t *testing.T) {
	safeTest(t, "Test_Cov41_newCharCollectionMapCreator_CapSelfCap", func() {
		m := corestr.New.CharCollectionMap.CapSelfCap(20, 5)
		tc := caseV1Compat{
			Name:     "Creator CapSelfCap",
			Expected: true,
			Actual:   m != nil && m.IsEmpty(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_newCharCollectionMapCreator_CapSelfCap_BelowMin(t *testing.T) {
	safeTest(t, "Test_Cov41_newCharCollectionMapCreator_CapSelfCap_BelowMin", func() {
		m := corestr.New.CharCollectionMap.CapSelfCap(1, 1)
		tc := caseV1Compat{
			Name:     "Creator CapSelfCap below min",
			Expected: true,
			Actual:   m != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_newCharCollectionMapCreator_Items(t *testing.T) {
	safeTest(t, "Test_Cov41_newCharCollectionMapCreator_Items", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		tc := caseV1Compat{
			Name:     "Creator Items",
			Expected: 2,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_newCharCollectionMapCreator_Items_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_newCharCollectionMapCreator_Items_Empty", func() {
		m := corestr.New.CharCollectionMap.Items([]string{})
		tc := caseV1Compat{
			Name:     "Creator Items empty",
			Expected: true,
			Actual:   m.IsEmpty(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_newCharCollectionMapCreator_ItemsPtrWithCap(t *testing.T) {
	safeTest(t, "Test_Cov41_newCharCollectionMapCreator_ItemsPtrWithCap", func() {
		m := corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 10, []string{"apple"})
		tc := caseV1Compat{
			Name:     "Creator ItemsPtrWithCap",
			Expected: 1,
			Actual:   m.AllLengthsSum(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_newCharCollectionMapCreator_ItemsPtrWithCap_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_newCharCollectionMapCreator_ItemsPtrWithCap_Empty", func() {
		m := corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 10, []string{})
		tc := caseV1Compat{
			Name:     "Creator ItemsPtrWithCap empty",
			Expected: true,
			Actual:   m.IsEmpty(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── CharCollectionDataModel ──────────────

func Test_Cov41_NewCharCollectionMapUsingDataModel(t *testing.T) {
	safeTest(t, "Test_Cov41_NewCharCollectionMapUsingDataModel", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		model := m.JsonModel()
		result := corestr.NewCharCollectionMapUsingDataModel(model)
		tc := caseV1Compat{
			Name:     "NewCharCollectionMapUsingDataModel",
			Expected: true,
			Actual:   result.Has("hello"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_NewCharCollectionMapDataModelUsing(t *testing.T) {
	safeTest(t, "Test_Cov41_NewCharCollectionMapDataModelUsing", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("hello")
		model := corestr.NewCharCollectionMapDataModelUsing(m)
		tc := caseV1Compat{
			Name:     "NewCharCollectionMapDataModelUsing",
			Expected: true,
			Actual:   model != nil && model.Items != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── emptyCreator ──────────────

func Test_Cov41_EmptyCreator_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_Cov41_EmptyCreator_CharCollectionMap", func() {
		m := corestr.Empty.CharCollectionMap()
		tc := caseV1Compat{
			Name:     "Empty.CharCollectionMap",
			Expected: true,
			Actual:   m != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_EmptyCreator_CharHashsetMap(t *testing.T) {
	safeTest(t, "Test_Cov41_EmptyCreator_CharHashsetMap", func() {
		m := corestr.Empty.CharHashsetMap()
		tc := caseV1Compat{
			Name:     "Empty.CharHashsetMap",
			Expected: true,
			Actual:   m != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_EmptyCreator_CollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_Cov41_EmptyCreator_CollectionsOfCollection", func() {
		c := corestr.Empty.CollectionsOfCollection()
		tc := caseV1Compat{
			Name:     "Empty.CollectionsOfCollection",
			Expected: true,
			Actual:   c != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_EmptyCreator_KeyValuesCollection(t *testing.T) {
	safeTest(t, "Test_Cov41_EmptyCreator_KeyValuesCollection", func() {
		c := corestr.Empty.KeyValuesCollection()
		tc := caseV1Compat{
			Name:     "Empty.KeyValuesCollection",
			Expected: true,
			Actual:   c != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_EmptyCreator_SimpleStringOnce(t *testing.T) {
	safeTest(t, "Test_Cov41_EmptyCreator_SimpleStringOnce", func() {
		s := corestr.Empty.SimpleStringOnce()
		tc := caseV1Compat{
			Name:     "Empty.SimpleStringOnce",
			Expected: true,
			Actual:   s.IsUninitialized(),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_EmptyCreator_SimpleStringOncePtr(t *testing.T) {
	safeTest(t, "Test_Cov41_EmptyCreator_SimpleStringOncePtr", func() {
		s := corestr.Empty.SimpleStringOncePtr()
		tc := caseV1Compat{
			Name:     "Empty.SimpleStringOncePtr",
			Expected: true,
			Actual:   s != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_EmptyCreator_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov41_EmptyCreator_KeyAnyValuePair", func() {
		p := corestr.Empty.KeyAnyValuePair()
		tc := caseV1Compat{
			Name:     "Empty.KeyAnyValuePair",
			Expected: true,
			Actual:   p != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_EmptyCreator_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov41_EmptyCreator_KeyValuePair", func() {
		p := corestr.Empty.KeyValuePair()
		tc := caseV1Compat{
			Name:     "Empty.KeyValuePair",
			Expected: true,
			Actual:   p != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_EmptyCreator_KeyValueCollection(t *testing.T) {
	safeTest(t, "Test_Cov41_EmptyCreator_KeyValueCollection", func() {
		c := corestr.Empty.KeyValueCollection()
		tc := caseV1Compat{
			Name:     "Empty.KeyValueCollection",
			Expected: true,
			Actual:   c != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_EmptyCreator_LeftRight(t *testing.T) {
	safeTest(t, "Test_Cov41_EmptyCreator_LeftRight", func() {
		lr := corestr.Empty.LeftRight()
		tc := caseV1Compat{
			Name:     "Empty.LeftRight",
			Expected: true,
			Actual:   lr != nil,
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

// ─── StringUtils ──────────────

func Test_Cov41_StringUtils_WrapDouble(t *testing.T) {
	safeTest(t, "Test_Cov41_StringUtils_WrapDouble", func() {
		tc := caseV1Compat{
			Name:     "WrapDouble",
			Expected: "\"hello\"",
			Actual:   corestr.StringUtils.WrapDouble("hello"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_StringUtils_WrapSingle(t *testing.T) {
	safeTest(t, "Test_Cov41_StringUtils_WrapSingle", func() {
		tc := caseV1Compat{
			Name:     "WrapSingle",
			Expected: "'hello'",
			Actual:   corestr.StringUtils.WrapSingle("hello"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_StringUtils_WrapTilda(t *testing.T) {
	safeTest(t, "Test_Cov41_StringUtils_WrapTilda", func() {
		tc := caseV1Compat{
			Name:     "WrapTilda",
			Expected: "`hello`",
			Actual:   corestr.StringUtils.WrapTilda("hello"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_StringUtils_WrapDoubleIfMissing_Already(t *testing.T) {
	safeTest(t, "Test_Cov41_StringUtils_WrapDoubleIfMissing_Already", func() {
		tc := caseV1Compat{
			Name:     "WrapDoubleIfMissing already wrapped",
			Expected: "\"hello\"",
			Actual:   corestr.StringUtils.WrapDoubleIfMissing("\"hello\""),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_StringUtils_WrapDoubleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_Cov41_StringUtils_WrapDoubleIfMissing_NotWrapped", func() {
		tc := caseV1Compat{
			Name:     "WrapDoubleIfMissing not wrapped",
			Expected: "\"hello\"",
			Actual:   corestr.StringUtils.WrapDoubleIfMissing("hello"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_StringUtils_WrapDoubleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_StringUtils_WrapDoubleIfMissing_Empty", func() {
		tc := caseV1Compat{
			Name:     "WrapDoubleIfMissing empty",
			Expected: "\"\"",
			Actual:   corestr.StringUtils.WrapDoubleIfMissing(""),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_StringUtils_WrapSingleIfMissing_Already(t *testing.T) {
	safeTest(t, "Test_Cov41_StringUtils_WrapSingleIfMissing_Already", func() {
		tc := caseV1Compat{
			Name:     "WrapSingleIfMissing already wrapped",
			Expected: "'hello'",
			Actual:   corestr.StringUtils.WrapSingleIfMissing("'hello'"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_StringUtils_WrapSingleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_Cov41_StringUtils_WrapSingleIfMissing_NotWrapped", func() {
		tc := caseV1Compat{
			Name:     "WrapSingleIfMissing not wrapped",
			Expected: "'hello'",
			Actual:   corestr.StringUtils.WrapSingleIfMissing("hello"),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov41_StringUtils_WrapSingleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_Cov41_StringUtils_WrapSingleIfMissing_Empty", func() {
		tc := caseV1Compat{
			Name:     "WrapSingleIfMissing empty",
			Expected: "''",
			Actual:   corestr.StringUtils.WrapSingleIfMissing(""),
			Args:     args.Map{},
		}
		tc.ShouldBeEqual(t)
	})
}
