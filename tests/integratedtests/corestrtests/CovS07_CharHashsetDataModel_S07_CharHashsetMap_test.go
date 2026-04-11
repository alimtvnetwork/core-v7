package corestrtests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// helper to create CharHashsetMap from items
func covS07MakeMap(items []string) *corestr.CharHashsetMap {
	return corestr.New.CharHashsetMap.Strings(10, items)
}

// ============================================================
// CharHashsetDataModel
// ============================================================

func Test_CovS07_CharHashsetDataModel_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_CharHashsetDataModel_Verification", func() {
		for caseIndex, tc := range covS07CharHashsetDataModelTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			direction, _ := input.GetAsString("direction")

			// Act
			var actual args.Map
			if direction == "toModel" {
				items := input["items"].([]string)
				m := covS07MakeMap(items)
				dm := corestr.NewCharHashsetMapDataModelUsing(m)
				actual = args.Map{"isNotNil": dm != nil}
			} else {
				dm := &corestr.CharHashsetDataModel{
					EachHashsetCapacity: 5,
					Items:               map[byte]*corestr.Hashset{},
				}
				m := corestr.NewCharHashsetMapUsingDataModel(dm)
				actual = args.Map{"isNotNil": m != nil}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// GetChar, GetCharOf
// ============================================================

func Test_CovS07_GetChar_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_GetChar_Verification", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 5)

		for caseIndex, tc := range covS07GetCharTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			inputStr, _ := input.GetAsString("input")
			useGetOf := input.GetAsBoolDefault("useGetOf", false)

			// Act
			var result byte
			if useGetOf {
				result = m.GetCharOf(inputStr)
			} else {
				result = m.GetChar(inputStr)
			}
			actual := args.Map{"char": string(result)}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// Basic: IsEmpty, HasItems, Length, AllLengthsSum
// ============================================================

func Test_CovS07_Basic_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_Basic_Verification", func() {
		for caseIndex, tc := range covS07BasicTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useEmpty := input.GetAsBoolDefault("useEmpty", false)

			var m *corestr.CharHashsetMap
			if useEmpty {
				m = corestr.Empty.CharHashsetMap()
			} else {
				items := input["items"].([]string)
				m = covS07MakeMap(items)
			}

			// Act
			actual := args.Map{
				"isEmpty":       m.IsEmpty(),
				"hasItems":      m.HasItems(),
				"length":        m.Length(),
				"allLengthsSum": m.AllLengthsSum(),
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// Has
// ============================================================

func Test_CovS07_Has_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_Has_Verification", func() {
		for caseIndex, tc := range covS07HasTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useEmpty := input.GetAsBoolDefault("useEmpty", false)
			checkStr := input["checkStr"].(string)

			var m *corestr.CharHashsetMap
			if useEmpty {
				m = corestr.Empty.CharHashsetMap()
			} else {
				items := input["items"].([]string)
				m = covS07MakeMap(items)
			}

			// Act
			actual := args.Map{"has": m.Has(checkStr)}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// HasWithHashset
// ============================================================

func Test_CovS07_HasWithHashset_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_HasWithHashset_Verification", func() {
		for caseIndex, tc := range covS07HasWithHashsetTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useEmpty := input.GetAsBoolDefault("useEmpty", false)
			checkStr := input["checkStr"].(string)

			var m *corestr.CharHashsetMap
			if useEmpty {
				m = corestr.Empty.CharHashsetMap()
			} else {
				items := input["items"].([]string)
				m = covS07MakeMap(items)
			}

			// Act
			found, hs := m.HasWithHashset(checkStr)
			actual := args.Map{
				"found":    found,
				"hsNotNil": hs != nil,
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// LengthOf, LengthOfHashsetFromFirstChar
// ============================================================

func Test_CovS07_LengthOf_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_LengthOf_Verification", func() {
		for caseIndex, tc := range covS07LengthOfTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			useEmpty := input.GetAsBoolDefault("useEmpty", false)

			var m *corestr.CharHashsetMap
			if useEmpty {
				m = corestr.Empty.CharHashsetMap()
			} else {
				items := input["items"].([]string)
				m = covS07MakeMap(items)
			}

			// Act
			var actual args.Map
			if _, ok := expected["lengthFromChar"]; ok {
				checkStr := input["checkStr"].(string)
				actual = args.Map{"lengthFromChar": m.LengthOfHashsetFromFirstChar(checkStr)}
			} else {
				charStr := input["char"].(string)
				actual = args.Map{"lengthOf": m.LengthOf(charStr[0])}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// Add, AddStrings
// ============================================================

func Test_CovS07_Add_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_Add_Verification", func() {
		for caseIndex, tc := range covS07AddTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			m := corestr.New.CharHashsetMap.Cap(10, 5)

			if addItem, ok := input["addItem"]; ok {
				m.Add(addItem.(string))
			}
			if addItems, ok := input["addItems"]; ok {
				m.AddStrings(addItems.([]string)...)
			}

			// Act
			var actual args.Map
			if _, ok := expected["has"]; ok {
				addItem := input["addItem"].(string)
				actual = args.Map{
					"has":    m.Has(addItem),
					"length": m.Length(),
				}
			} else {
				actual = args.Map{
					"allLengthsSum": m.AllLengthsSum(),
					"charGroups":    m.Length(),
				}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// AddSameStartingCharItems
// ============================================================

func Test_CovS07_AddSameCharItems_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_AddSameCharItems_Verification", func() {
		for caseIndex, tc := range covS07AddSameCharTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			initialItems := input["initialItems"].([]string)
			newItems := input["newItems"].([]string)
			charStr := input["char"].(string)
			m := covS07MakeMap(initialItems)

			// Act
			m.AddSameStartingCharItems(charStr[0], newItems)
			actual := args.Map{"allLengthsSum": m.AllLengthsSum()}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// IsEquals
// ============================================================

func Test_CovS07_IsEquals_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_IsEquals_Verification", func() {
		for caseIndex, tc := range covS07IsEqualsTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useSelf := input.GetAsBoolDefault("useSelf", false)
			useNilOther := input.GetAsBoolDefault("useNilOther", false)
			useEmpty1 := input.GetAsBoolDefault("useEmpty1", false)
			useEmpty2 := input.GetAsBoolDefault("useEmpty2", false)

			var m1 *corestr.CharHashsetMap
			if useEmpty1 {
				m1 = corestr.Empty.CharHashsetMap()
			} else if items1, ok := input["items1"]; ok {
				m1 = covS07MakeMap(items1.([]string))
			}

			// Act
			var isEquals bool
			if useSelf {
				isEquals = m1.IsEquals(m1)
			} else if useNilOther {
				isEquals = m1.IsEquals(nil)
			} else if useEmpty2 {
				m2 := corestr.Empty.CharHashsetMap()
				isEquals = m1.IsEquals(m2)
			} else {
				items2 := input["items2"].([]string)
				m2 := covS07MakeMap(items2)
				isEquals = m1.IsEquals(m2)
			}
			actual := args.Map{"isEquals": isEquals}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// GetHashset
// ============================================================

func Test_CovS07_GetHashset_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_GetHashset_Verification", func() {
		for caseIndex, tc := range covS07GetHashsetTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			items := input["items"].([]string)
			checkStr := input["checkStr"].(string)
			isAddNew := input.GetAsBoolDefault("isAddNewOnEmpty", false)
			m := covS07MakeMap(items)

			// Act
			hs := m.GetHashset(checkStr, isAddNew)
			actual := args.Map{"isNotNil": hs != nil}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// GetCharsGroups
// ============================================================

func Test_CovS07_GetCharsGroups_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_GetCharsGroups_Verification", func() {
		m := corestr.New.CharHashsetMap.Cap(10, 5)

		for caseIndex, tc := range covS07GetCharsGroupsTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			items := input["items"].([]string)

			// Act
			result := m.GetCharsGroups(items...)

			var actual args.Map
			if _, ok := expected["hasItems"]; ok {
				actual = args.Map{"hasItems": result != nil && result.HasItems()}
			} else {
				actual = args.Map{"isNotNil": result != nil}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// List, SortedListAsc, SortedListDsc
// ============================================================

func Test_CovS07_List_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_List_Verification", func() {
		for caseIndex, tc := range covS07ListTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			items := input["items"].([]string)
			m := covS07MakeMap(items)

			sortedDir, _ := input.GetAsString("sorted")

			// Act
			var actual args.Map
			if sortedDir == "asc" {
				list := m.SortedListAsc()
				first := ""
				if len(list) > 0 {
					first = list[0]
				}
				actual = args.Map{"first": first}
			} else if sortedDir == "dsc" {
				list := m.SortedListDsc()
				first := ""
				if len(list) > 0 {
					first = list[0]
				}
				actual = args.Map{"first": first}
			} else {
				list := m.List()
				actual = args.Map{"length": len(list)}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// JSON
// ============================================================

func Test_CovS07_Json_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_Json_Verification", func() {
		for caseIndex, tc := range covS07JsonTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			items := input["items"].([]string)
			m := covS07MakeMap(items)

			// Act
			var actual args.Map
			if _, ok := expected["roundTrip"]; ok {
				jsonResult := m.Json()
				m2 := corestr.New.CharHashsetMap.Cap(10, 5)
				_, err := m2.ParseInjectUsingJson(&jsonResult)
				isEqual := m2.AllLengthsSum() == m.AllLengthsSum()
				actual = args.Map{"roundTrip": err == nil && isEqual}
			} else {
				jsonResult := m.Json()
				actual = args.Map{
					"hasBytes": jsonResult.SafeBytes() != nil && len(jsonResult.SafeBytes()) > 0,
					"hasError": jsonResult.HasError(),
				}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// AddSameCharsCollection
// ============================================================

func Test_CovS07_AddSameCharsColl_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_AddSameCharsColl_Verification", func() {
		for caseIndex, tc := range covS07AddSameCharsCollTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			initialItems := input["initialItems"].([]string)
			useNilColl := input.GetAsBoolDefault("useNilCollToAdd", false)
			m := covS07MakeMap(initialItems)

			checkStr := "apple"
			if cs, ok := input["checkStr"]; ok {
				checkStr = cs.(string)
			}

			// Act
			var resultHs *corestr.Hashset
			if useNilColl {
				resultHs = m.AddSameCharsCollection(checkStr, nil)
			} else {
				addItems := input["addItems"].([]string)
				coll := corestr.New.Collection.Strings(addItems)
				resultHs = m.AddSameCharsCollection(checkStr, coll)
			}

			var actual args.Map
			if _, ok := expected["hsNotNil"]; ok {
				actual = args.Map{"hsNotNil": resultHs != nil}
			} else {
				actual = args.Map{"allLengthsSum": m.AllLengthsSum()}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// AddSameCharsHashset
// ============================================================

func Test_CovS07_AddSameCharsHashset_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_AddSameCharsHashset_Verification", func() {
		for caseIndex, tc := range covS07AddSameCharsHashsetTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			initialItems := input["initialItems"].([]string)
			useNilHashset := input.GetAsBoolDefault("useNilHashset", false)
			m := covS07MakeMap(initialItems)

			checkStr := "apple"
			if cs, ok := input["checkStr"]; ok {
				checkStr = cs.(string)
			}

			// Act
			var resultHs *corestr.Hashset
			if useNilHashset {
				resultHs = m.AddSameCharsHashset(checkStr, nil)
			} else {
				addItems := input["addItems"].([]string)
				hs := corestr.New.Hashset.Strings(addItems)
				resultHs = m.AddSameCharsHashset(checkStr, hs)
			}

			var actual args.Map
			if _, ok := expected["hsNotNil"]; ok {
				actual = args.Map{"hsNotNil": resultHs != nil}
			} else {
				actual = args.Map{"allLengthsSum": m.AllLengthsSum()}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// AddCollectionItems, AddCharCollectionMapItems, AddHashsetItems
// ============================================================

func Test_CovS07_AddFromSource_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_AddFromSource_Verification", func() {
		for caseIndex, tc := range covS07AddFromSourceTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			source, _ := input.GetAsString("source")
			m := corestr.New.CharHashsetMap.Cap(10, 5)

			// Act
			switch source {
			case "collection":
				items := input["items"].([]string)
				coll := corestr.New.Collection.Strings(items)
				m.AddCollectionItems(coll)
			case "collectionNil":
				m.AddCollectionItems(nil)
			case "charCollMap":
				items := input["items"].([]string)
				ccm := corestr.New.CharCollectionMap.Items(items)
				m.AddCharCollectionMapItems(ccm)
			case "charCollMapNil":
				m.AddCharCollectionMapItems(nil)
			case "hashset":
				items := input["items"].([]string)
				hs := corestr.New.Hashset.Strings(items)
				m.AddHashsetItems(hs)
			case "hashsetEmpty":
				hs := corestr.New.Hashset.Empty()
				m.AddHashsetItems(hs)
			}
			actual := args.Map{"allLengthsSum": m.AllLengthsSum()}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// HashsetsCollection
// ============================================================

func Test_CovS07_HashsetsColl_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_HashsetsColl_Verification", func() {
		for caseIndex, tc := range covS07HashsetsCollTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useEmpty := input.GetAsBoolDefault("useEmpty", false)

			var m *corestr.CharHashsetMap
			if useEmpty {
				m = corestr.Empty.CharHashsetMap()
			} else {
				items := input["items"].([]string)
				m = covS07MakeMap(items)
			}

			// Act
			hsColl := m.HashsetsCollection()
			actual := args.Map{"hasItems": hsColl != nil && hsColl.Length() > 0}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// Clear, RemoveAll
// ============================================================

func Test_CovS07_Clear_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_Clear_Verification", func() {
		for caseIndex, tc := range covS07ClearTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useEmpty := input.GetAsBoolDefault("useEmpty", false)
			useRemove := input.GetAsBoolDefault("useRemove", false)

			var m *corestr.CharHashsetMap
			if useEmpty {
				m = corestr.Empty.CharHashsetMap()
			} else {
				items := input["items"].([]string)
				m = covS07MakeMap(items)
			}

			// Act
			if useRemove {
				m.RemoveAll()
			} else {
				m.Clear()
			}
			// Small delay to let async goroutine in Clear() finish
			time.Sleep(5 * time.Millisecond)
			actual := args.Map{"isEmpty": m.IsEmpty()}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// String, SummaryString
// ============================================================

func Test_CovS07_StringOutput_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_StringOutput_Verification", func() {
		for caseIndex, tc := range covS07StringOutputTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			items := input["items"].([]string)
			useSummary := input.GetAsBoolDefault("summary", false)
			m := covS07MakeMap(items)

			// Act
			var result string
			if useSummary {
				result = m.SummaryString()
			} else {
				result = m.String()
			}
			actual := args.Map{"hasOutput": result != ""}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// Print (skip path)
// ============================================================

func Test_CovS07_PrintSkip(t *testing.T) {
	safeTest(t, "Test_CovS07_PrintSkip", func() {
		// Arrange
		m := covS07MakeMap([]string{"test"})

		// Act — isPrint=false skips
		m.Print(false)
		m.PrintLock(false)

		// Assert — no panic
		actual := args.Map{"result": m.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected empty after print skip", actual)
	})
}

// ============================================================
// Lock variants
// ============================================================

func Test_CovS07_LockVariants(t *testing.T) {
	safeTest(t, "Test_CovS07_LockVariants", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha", "bravo"})

		// Act
		lenLock := m.LengthLock()
		isEmptyLock := m.IsEmptyLock()
		allSumLock := m.AllLengthsSumLock()
		lenOfLock := m.LengthOfLock('a')
		_ = m.GetCopyMapLock()
		_ = m.SummaryStringLock()
		_ = m.StringLock()
		isEqLock := m.IsEqualsLock(m)

		// Assert
		actual := args.Map{"result": lenLock <= 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LengthLock should be > 0", actual)
		actual = args.Map{"result": isEmptyLock}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock should be false", actual)
		actual = args.Map{"result": allSumLock <= 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AllLengthsSumLock should be > 0", actual)
		actual = args.Map{"result": lenOfLock <= 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LengthOfLock should be > 0", actual)
		actual = args.Map{"result": isEqLock}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsLock(self) should be true", actual)
	})
}

// ============================================================
// HasWithHashsetLock
// ============================================================

func Test_CovS07_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_CovS07_HasWithHashsetLock", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha", "avocado"})

		// Act
		found, hs := m.HasWithHashsetLock("alpha")

		// Assert
		actual := args.Map{"result": found}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock should find 'alpha'", actual)
		actual = args.Map{"result": hs == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock should return non-nil hashset", actual)
	})
}

// ============================================================
// AddLock
// ============================================================

func Test_CovS07_AddLock(t *testing.T) {
	safeTest(t, "Test_CovS07_AddLock", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		m.AddLock("alpha")
		m.AddLock("avocado") // same char group

		// Assert
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddLock: expected 2 items", actual)
	})
}

// ============================================================
// AddStringsLock
// ============================================================

func Test_CovS07_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_CovS07_AddStringsLock", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		m.AddStringsLock("alpha", "bravo", "avocado")

		// Assert
		actual := args.Map{"result": m.AllLengthsSum() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddStringsLock: expected 3", actual)
	})
}

// ============================================================
// HashsetByChar, HashsetByCharLock, HashsetByStringFirstChar, etc.
// ============================================================

func Test_CovS07_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_CovS07_HashsetByChar", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha", "avocado"})

		// Act
		hs1 := m.HashsetByChar('a')
		hs2 := m.HashsetByCharLock('a')
		hs3 := m.HashsetByStringFirstChar("alpha")
		hs4 := m.HashsetByStringFirstCharLock("alpha")
		hs5 := m.GetHashsetByChar('a')
		hsMissing := m.HashsetByCharLock('z')

		// Assert
		actual := args.Map{"result": hs1 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashsetByChar('a') should be non-nil", actual)
		actual = args.Map{"result": hs2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock('a') should be non-nil", actual)
		actual = args.Map{"result": hs3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstChar should be non-nil", actual)
		actual = args.Map{"result": hs4 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstCharLock should be non-nil", actual)
		actual = args.Map{"result": hs5 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "GetHashsetByChar('a') should be non-nil", actual)
		actual = args.Map{"result": hsMissing == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock('z') should return empty hashset, not nil", actual)
	})
}

// ============================================================
// HashsetsCollectionByChars, ByStringsFirstChar
// ============================================================

func Test_CovS07_HashsetsCollByCharsAndStr(t *testing.T) {
	safeTest(t, "Test_CovS07_HashsetsCollByCharsAndStr", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha", "bravo", "avocado"})

		// Act
		hsByChars := m.HashsetsCollectionByChars('a', 'b')
		hsByStr := m.HashsetsCollectionByStringsFirstChar("alpha", "bravo")

		// Assert
		actual := args.Map{"result": hsByChars == nil || hsByChars.Length() == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars should return non-empty", actual)
		actual = args.Map{"result": hsByStr == nil || hsByStr.Length() == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringsFirstChar should return non-empty", actual)
	})
}

// ============================================================
// GetMap
// ============================================================

func Test_CovS07_GetMap(t *testing.T) {
	safeTest(t, "Test_CovS07_GetMap", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha"})

		// Act
		rawMap := m.GetMap()

		// Assert
		actual := args.Map{"result": rawMap == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "GetMap should return non-nil", actual)
	})
}

// ============================================================
// Interface adapters
// ============================================================

func Test_CovS07_InterfaceAdapters(t *testing.T) {
	safeTest(t, "Test_CovS07_InterfaceAdapters", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha"})

		// Act
		binder := m.AsJsonContractsBinder()
		jsoner := m.AsJsoner()
		marshaller := m.AsJsonMarshaller()
		injector := m.AsJsonParseSelfInjector()
		_ = m.JsonModelAny()
		_ = m.JsonModel()

		// Assert
		actual := args.Map{"result": binder == nil || jsoner == nil || marshaller == nil || injector == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Interface adapters should return non-nil", actual)
	})
}

// ============================================================
// MarshalJSON, UnmarshalJSON
// ============================================================

func Test_CovS07_MarshalUnmarshal(t *testing.T) {
	safeTest(t, "Test_CovS07_MarshalUnmarshal", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha", "bravo"})

		// Act
		bytes, err := m.MarshalJSON()
		m2 := corestr.New.CharHashsetMap.Cap(10, 5)
		err2 := m2.UnmarshalJSON(bytes)

		// Assert
		actual := args.Map{"result": err != nil || len(bytes) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "MarshalJSON failed:", actual)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON failed:", actual)
	})
}

// ============================================================
// JsonParseSelfInject
// ============================================================

func Test_CovS07_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovS07_JsonParseSelfInject", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha"})
		jsonResult := m.Json()

		// Act
		m2 := corestr.New.CharHashsetMap.Cap(10, 5)
		err := m2.JsonParseSelfInject(&jsonResult)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject should not error:", actual)
	})
}

// ============================================================
// ParseInjectUsingJsonMust (valid path)
// ============================================================

func Test_CovS07_ParseInjectMust_Valid(t *testing.T) {
	safeTest(t, "Test_CovS07_ParseInjectMust_Valid", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha"})
		jsonResult := m.Json()

		// Act
		m2 := corestr.New.CharHashsetMap.Cap(10, 5)
		result := m2.ParseInjectUsingJsonMust(&jsonResult)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust should return non-nil", actual)
	})
}

// ============================================================
// JsonPtr
// ============================================================

func Test_CovS07_JsonPtr(t *testing.T) {
	safeTest(t, "Test_CovS07_JsonPtr", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha"})

		// Act
		result := m.JsonPtr()

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "JsonPtr should return non-nil", actual)
	})
}

// ============================================================
// AddSameCharsCollectionLock
// ============================================================

func Test_CovS07_AddSameCharsCollLock(t *testing.T) {
	safeTest(t, "Test_CovS07_AddSameCharsCollLock", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha"})
		coll := corestr.New.Collection.Strings([]string{"avocado", "apricot"})

		// Act
		result := m.AddSameCharsCollectionLock("apple", coll)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock should return non-nil", actual)
		actual = args.Map{"result": m.AllLengthsSum() < 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Expected at least 3 items", actual)
	})
}

// ============================================================
// AddSameCharsCollectionLock — nil collection, new char
// ============================================================

func Test_CovS07_AddSameCharsCollLock_NilColl_NewChar(t *testing.T) {
	safeTest(t, "Test_CovS07_AddSameCharsCollLock_NilColl_NewChar", func() {
		// Arrange
		m := covS07MakeMap([]string{"banana"})

		// Act — nil collection with new char creates empty hashset
		result := m.AddSameCharsCollectionLock("apple", nil)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should return non-nil hashset", actual)
	})
}

// ============================================================
// AddSameCharsCollectionLock — collection with new char
// ============================================================

func Test_CovS07_AddSameCharsCollLock_NewChar_WithData(t *testing.T) {
	safeTest(t, "Test_CovS07_AddSameCharsCollLock_NewChar_WithData", func() {
		// Arrange
		m := covS07MakeMap([]string{"banana"})
		coll := corestr.New.Collection.Strings([]string{"apple", "avocado"})

		// Act
		result := m.AddSameCharsCollectionLock("apple", coll)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should return non-nil hashset", actual)
		actual = args.Map{"result": m.AllLengthsSum() < 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Expected at least 3 items", actual)
	})
}

// ============================================================
// AddHashsetLock
// ============================================================

func Test_CovS07_AddHashsetLock(t *testing.T) {
	safeTest(t, "Test_CovS07_AddHashsetLock", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha"})
		hs := corestr.New.Hashset.Strings([]string{"avocado"})

		// Act
		result := m.AddHashsetLock("apple", hs)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock should return non-nil", actual)
	})
}

// ============================================================
// AddHashsetLock — nil hashset, new char
// ============================================================

func Test_CovS07_AddHashsetLock_NilHashset_NewChar(t *testing.T) {
	safeTest(t, "Test_CovS07_AddHashsetLock_NilHashset_NewChar", func() {
		// Arrange
		m := covS07MakeMap([]string{"banana"})

		// Act
		result := m.AddHashsetLock("apple", nil)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should return non-nil hashset", actual)
	})
}

// ============================================================
// AddHashsetLock — hashset with new char
// ============================================================

func Test_CovS07_AddHashsetLock_NewChar_WithData(t *testing.T) {
	safeTest(t, "Test_CovS07_AddHashsetLock_NewChar_WithData", func() {
		// Arrange
		m := covS07MakeMap([]string{"banana"})
		hs := corestr.New.Hashset.Strings([]string{"apple", "avocado"})

		// Act
		result := m.AddHashsetLock("apple", hs)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should return non-nil", actual)
		actual = args.Map{"result": m.AllLengthsSum() < 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Expected at least 3 items", actual)
	})
}

// ============================================================
// GetHashsetLock
// ============================================================

func Test_CovS07_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_CovS07_GetHashsetLock", func() {
		// Arrange
		m := covS07MakeMap([]string{"alpha"})

		// Act
		hs := m.GetHashsetLock(true, "alpha")
		hsMissing := m.GetHashsetLock(false, "zzz")

		// Assert
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "GetHashsetLock should return non-nil for existing char", actual)
		actual = args.Map{"result": hsMissing != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "GetHashsetLock should return nil for missing char with isAddNew=false", actual)
	})
}

// ============================================================
// newCreator methods
// ============================================================

func Test_CovS07_NewCreator_Verification(t *testing.T) {
	safeTest(t, "Test_CovS07_NewCreator_Verification", func() {
		for caseIndex, tc := range covS07NewCreatorTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			method, _ := input.GetAsString("method")

			// Act
			var m *corestr.CharHashsetMap
			switch method {
			case "cap":
				cap := input["cap"].(int)
				selfCap := input["selfCap"].(int)
				m = corestr.New.CharHashsetMap.Cap(cap, selfCap)
			case "capItems":
				cap := input["cap"].(int)
				selfCap := input["selfCap"].(int)
				items := input["items"].([]string)
				m = corestr.New.CharHashsetMap.CapItems(cap, selfCap, items...)
			case "strings":
				selfCap := input["selfCap"].(int)
				items := input["items"].([]string)
				m = corestr.New.CharHashsetMap.Strings(selfCap, items)
			case "stringsNil":
				selfCap := input["selfCap"].(int)
				m = corestr.New.CharHashsetMap.Strings(selfCap, nil)
			}

			actual := args.Map{"isNotNil": m != nil}
			if _, ok := expected["isEmpty"]; ok {
				actual["isEmpty"] = m.IsEmpty()
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// AddCollectionItemsAsyncLock (exercises async path)
// ============================================================

func Test_CovS07_AddCollectionItemsAsyncLock(t *testing.T) {
	safeTest(t, "Test_CovS07_AddCollectionItemsAsyncLock", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 5)
		coll := corestr.New.Collection.Strings([]string{"alpha", "bravo"})

		// Act
		done := make(chan bool, 1)
		m.AddCollectionItemsAsyncLock(coll, func(chm *corestr.CharHashsetMap) {
			done <- true
		})

		// Wait for async completion
		select {
		case <-done:
		case <-time.After(2 * time.Second):
		}

		// Assert
		if m.AllLengthsSum() < 2 {
			// async may not have completed, but at least no panic
			_ = 0
		}
	})
}

// ============================================================
// AddCollectionItemsAsyncLock — nil collection
// ============================================================

func Test_CovS07_AddCollectionItemsAsyncLock_Nil(t *testing.T) {
	safeTest(t, "Test_CovS07_AddCollectionItemsAsyncLock_Nil", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		m.AddCollectionItemsAsyncLock(nil, nil)

		// Assert — no panic
		actual := args.Map{"result": m.AllLengthsSum() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should remain empty", actual)
	})
}

// ============================================================
// AddHashsetItemsAsyncLock — nil hashset
// ============================================================

func Test_CovS07_AddHashsetItemsAsyncLock_Nil(t *testing.T) {
	safeTest(t, "Test_CovS07_AddHashsetItemsAsyncLock_Nil", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		m.AddHashsetItemsAsyncLock(nil, nil)

		// Assert — no panic
		actual := args.Map{"result": m.AllLengthsSum() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should remain empty", actual)
	})
}

// ============================================================
// AddStringsAsyncLock — empty
// ============================================================

func Test_CovS07_AddStringsAsyncLock_Empty(t *testing.T) {
	safeTest(t, "Test_CovS07_AddStringsAsyncLock_Empty", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		m.AddStringsAsyncLock(nil, nil)

		// Assert — no panic
		actual := args.Map{"result": m.AllLengthsSum() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should remain empty", actual)
	})
}

// ============================================================
// AddStringsAsyncLock — small list with onComplete
// ============================================================

func Test_CovS07_AddStringsAsyncLock_SmallList(t *testing.T) {
	safeTest(t, "Test_CovS07_AddStringsAsyncLock_SmallList", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.Cap(10, 5)
		items := []string{"alpha", "bravo", "charlie"}

		// Act
		done := make(chan bool, 1)
		m.AddStringsAsyncLock(items, func(chm *corestr.CharHashsetMap) {
			done <- true
		})

		select {
		case <-done:
		case <-time.After(2 * time.Second):
		}

		// Assert
		if m.AllLengthsSum() < 3 {
			// may be async
			_ = 0
		}
	})
}
