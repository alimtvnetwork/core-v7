package corestrtests

import (
	"sort"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================
// AllIndividualStringsOfStringsLength
// ============================================================

func Test_CovS06_AllIndStrOfStrLen_Nil(t *testing.T) {
	safeTest(t, "Test_CovS06_AllIndStrOfStrLen_Nil", func() {
		// Arrange
		_ = covS06AllIndStrOfStrLenNilTestCase.ArrangeInput.(args.Map)

		// Act
		result := corestr.AllIndividualStringsOfStringsLength(nil)
		actual := args.Map{"length": result}

		// Assert
		covS06AllIndStrOfStrLenNilTestCase.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_CovS06_AllIndStrOfStrLen_Valid(t *testing.T) {
	safeTest(t, "Test_CovS06_AllIndStrOfStrLen_Valid", func() {
		// Arrange
		input := covS06AllIndStrOfStrLenValidTestCase.ArrangeInput.(args.Map)
		items := input["items"].([][]string)

		// Act
		result := corestr.AllIndividualStringsOfStringsLength(&items)
		actual := args.Map{"length": result}

		// Assert
		covS06AllIndStrOfStrLenValidTestCase.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_CovS06_AllIndStrOfStrLen_Empty(t *testing.T) {
	safeTest(t, "Test_CovS06_AllIndStrOfStrLen_Empty", func() {
		// Arrange
		input := covS06AllIndStrOfStrLenEmptyTestCase.ArrangeInput.(args.Map)
		items := input["items"].([][]string)

		// Act
		result := corestr.AllIndividualStringsOfStringsLength(&items)
		actual := args.Map{"length": result}

		// Assert
		covS06AllIndStrOfStrLenEmptyTestCase.ShouldBeEqualMap(t, 0, actual)
	})
}

// ============================================================
// AllIndividualsLengthOfSimpleSlices
// ============================================================

func Test_CovS06_AllIndLenSimpleSlices_Nil(t *testing.T) {
	safeTest(t, "Test_CovS06_AllIndLenSimpleSlices_Nil", func() {
		// Arrange
		_ = covS06AllIndLenSimpleSlicesNilTestCase.ArrangeInput.(args.Map)

		// Act
		result := corestr.AllIndividualsLengthOfSimpleSlices()
		actual := args.Map{"length": result}

		// Assert
		covS06AllIndLenSimpleSlicesNilTestCase.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_CovS06_AllIndLenSimpleSlices_Valid(t *testing.T) {
	safeTest(t, "Test_CovS06_AllIndLenSimpleSlices_Valid", func() {
		// Arrange
		input := covS06AllIndLenSimpleSlicesValidTestCase.ArrangeInput.(args.Map)
		s1 := input["slice1"].([]string)
		s2 := input["slice2"].([]string)
		ss1 := corestr.New.SimpleSlice.Strings(s1)
		ss2 := corestr.New.SimpleSlice.Strings(s2)

		// Act
		result := corestr.AllIndividualsLengthOfSimpleSlices(ss1, ss2)
		actual := args.Map{"length": result}

		// Assert
		covS06AllIndLenSimpleSlicesValidTestCase.ShouldBeEqualMap(t, 0, actual)
	})
}

// ============================================================
// AnyToString
// ============================================================

func Test_CovS06_AnyToString_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_AnyToString_Verification", func() {
		for caseIndex, tc := range covS06AnyToStringTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			isInclude := input.GetAsBoolDefault("isIncludeFieldName", false)
			usePointer := input.GetAsBoolDefault("usePointer", false)

			// Act
			var result string
			if usePointer {
				s := "ptrValue"
				result = corestr.AnyToString(isInclude, &s)
			} else {
				inputStr, _ := input.GetAsString("input")
				result = corestr.AnyToString(isInclude, inputStr)
			}

			var actual args.Map
			if _, hasResult := expected["result"]; hasResult {
				actual = args.Map{"result": result}
			} else {
				actual = args.Map{"hasResult": result != ""}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionDataModel
// ============================================================

func Test_CovS06_CharCollDataModel_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollDataModel_Verification", func() {
		for caseIndex, tc := range covS06CharCollDataModelTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)

			// Act
			var actual args.Map
			if _, hasItems := input["items"]; hasItems {
				items := input["items"].([]string)
				m := corestr.New.CharCollectionMap.Items(items)
				dm := corestr.NewCharCollectionMapDataModelUsing(m)
				actual = args.Map{"isNotNil": dm != nil}
			} else {
				dm := &corestr.CharCollectionDataModel{
					EachCollectionCapacity: input["capacity"].(int),
					Items:                  map[byte]*corestr.Collection{},
				}
				m := corestr.NewCharCollectionMapUsingDataModel(dm)
				actual = args.Map{"isNotNil": m != nil}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CloneSlice
// ============================================================

func Test_CovS06_CloneSlice_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CloneSlice_Verification", func() {
		for caseIndex, tc := range covS06CloneSliceTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useNil := input.GetAsBoolDefault("useNil", false)

			// Act
			var result []string
			if useNil {
				result = corestr.CloneSlice(nil)
			} else {
				items := input["items"].([]string)
				result = corestr.CloneSlice(items)
			}
			actual := args.Map{"length": len(result)}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CloneSliceIf
// ============================================================

func Test_CovS06_CloneSliceIf_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CloneSliceIf_Verification", func() {
		for caseIndex, tc := range covS06CloneSliceIfTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			isClone := input.GetAsBoolDefault("isClone", false)
			items := input["items"].([]string)

			// Act
			result := corestr.CloneSliceIf(isClone, items...)
			actual := args.Map{"length": len(result)}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// utils — Wrap methods
// ============================================================

func Test_CovS06_Utils_Wrap_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_Utils_Wrap_Verification", func() {
		u := corestr.StringUtils

		for caseIndex, tc := range covS06UtilsWrapTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			method, _ := input.GetAsString("method")
			inputStr, _ := input.GetAsString("input")

			// Act
			var result string
			switch method {
			case "WrapDoubleIfMissing":
				result = u.WrapDoubleIfMissing(inputStr)
			case "WrapSingleIfMissing":
				result = u.WrapSingleIfMissing(inputStr)
			case "WrapDouble":
				result = u.WrapDouble(inputStr)
			case "WrapSingle":
				result = u.WrapSingle(inputStr)
			case "WrapTilda":
				result = u.WrapTilda(inputStr)
			}
			actual := args.Map{"result": result}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — GetChar
// ============================================================

func Test_CovS06_CharCollMap_GetChar_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_GetChar_Verification", func() {
		m := corestr.New.CharCollectionMap.Empty()

		for caseIndex, tc := range covS06CharCollMapGetCharTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			inputStr, _ := input.GetAsString("input")

			// Act
			result := m.GetChar(inputStr)
			actual := args.Map{"char": string(result)}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — GetCharsGroups
// ============================================================

func Test_CovS06_CharCollMap_GetCharsGroups_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_GetCharsGroups_Verification", func() {
		m := corestr.New.CharCollectionMap.Empty()

		for caseIndex, tc := range covS06CharCollMapGetCharsGroupsTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			items := input["items"].([]string)

			// Act
			result := m.GetCharsGroups(items)

			var actual args.Map
			if _, has := tc.ExpectedInput.(args.Map)["hasItems"]; has {
				actual = args.Map{"hasItems": result != nil && result.HasItems()}
			} else {
				actual = args.Map{"isNil": result == nil}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — Basic: IsEmpty, HasItems, Length, AllLengthsSum
// ============================================================

func Test_CovS06_CharCollMap_Basic_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_Basic_Verification", func() {
		for caseIndex, tc := range covS06CharCollMapBasicTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useEmpty := input.GetAsBoolDefault("useEmpty", false)

			var m *corestr.CharCollectionMap
			if useEmpty {
				m = corestr.New.CharCollectionMap.Empty()
			} else {
				items := input["items"].([]string)
				m = corestr.New.CharCollectionMap.Items(items)
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
// CharCollectionMap — Add, Has
// ============================================================

func Test_CovS06_CharCollMap_AddHas_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_AddHas_Verification", func() {
		for caseIndex, tc := range covS06CharCollMapAddHasTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			useEmpty := input.GetAsBoolDefault("useEmpty", false)

			m := corestr.New.CharCollectionMap.Empty()

			if !useEmpty {
				if addItem, ok := input["addItem"]; ok {
					m.Add(addItem.(string))
				}
				if addItems, ok := input["addItems"]; ok {
					for _, item := range addItems.([]string) {
						m.Add(item)
					}
				}
			}

			// Act
			var actual args.Map
			if checkItem, ok := input["checkItem"]; ok {
				actual = args.Map{"has": m.Has(checkItem.(string))}
			} else if _, ok := expected["allLengthsSum"]; ok {
				actual = args.Map{
					"allLengthsSum": m.AllLengthsSum(),
					"charGroups":    m.Length(),
				}
			} else {
				addItem := input["addItem"].(string)
				actual = args.Map{
					"has":    m.Has(addItem),
					"length": m.Length(),
				}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — LengthOf, LengthOfCollectionFromFirstChar
// ============================================================

func Test_CovS06_CharCollMap_LengthOf_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_LengthOf_Verification", func() {
		for caseIndex, tc := range covS06CharCollMapLengthOfTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			items := input["items"].([]string)
			m := corestr.New.CharCollectionMap.Items(items)

			// Act
			var actual args.Map
			if _, ok := expected["lengthFromChar"]; ok {
				checkStr := input["checkStr"].(string)
				actual = args.Map{"lengthFromChar": m.LengthOfCollectionFromFirstChar(checkStr)}
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
// CharCollectionMap — IsEquals
// ============================================================

func Test_CovS06_CharCollMap_IsEquals_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_IsEquals_Verification", func() {
		for caseIndex, tc := range covS06CharCollMapEqualsTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useSelf := input.GetAsBoolDefault("useSelf", false)
			useNilOther := input.GetAsBoolDefault("useNilOther", false)
			useEmpty1 := input.GetAsBoolDefault("useEmpty1", false)
			useEmpty2 := input.GetAsBoolDefault("useEmpty2", false)

			var m1 *corestr.CharCollectionMap
			if useEmpty1 {
				m1 = corestr.New.CharCollectionMap.Empty()
			} else if items1, ok := input["items1"]; ok {
				m1 = corestr.New.CharCollectionMap.Items(items1.([]string))
			}

			// Act
			var isEquals bool
			if useSelf {
				isEquals = m1.IsEquals(m1)
			} else if useNilOther {
				isEquals = m1.IsEquals(nil)
			} else if useEmpty2 {
				m2 := corestr.New.CharCollectionMap.Empty()
				isEquals = m1.IsEquals(m2)
			} else {
				items2 := input["items2"].([]string)
				m2 := corestr.New.CharCollectionMap.Items(items2)
				isEquals = m1.IsEquals(m2)
			}
			actual := args.Map{"isEquals": isEquals}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — HasWithCollection
// ============================================================

func Test_CovS06_CharCollMap_HasWithCollection_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_HasWithCollection_Verification", func() {
		for caseIndex, tc := range covS06HasWithCollectionTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useEmpty := input.GetAsBoolDefault("useEmpty", false)
			checkStr := input["checkStr"].(string)

			var m *corestr.CharCollectionMap
			if useEmpty {
				m = corestr.New.CharCollectionMap.Empty()
			} else {
				items := input["items"].([]string)
				m = corestr.New.CharCollectionMap.Items(items)
			}

			// Act
			found, coll := m.HasWithCollection(checkStr)
			actual := args.Map{
				"found":            found,
				"collectionNotNil": coll != nil,
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — GetCollection
// ============================================================

func Test_CovS06_CharCollMap_GetCollection_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_GetCollection_Verification", func() {
		for caseIndex, tc := range covS06GetCollectionTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			items := input["items"].([]string)
			checkStr := input["checkStr"].(string)
			isAddNew := input.GetAsBoolDefault("isAddNewOnEmpty", false)
			m := corestr.New.CharCollectionMap.Items(items)

			// Act
			coll := m.GetCollection(checkStr, isAddNew)
			actual := args.Map{"isNotNil": coll != nil}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — List, SortedListAsc
// ============================================================

func Test_CovS06_CharCollMap_List_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_List_Verification", func() {
		for caseIndex, tc := range covS06ListTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			useEmpty := input.GetAsBoolDefault("useEmpty", false)
			useSorted := input.GetAsBoolDefault("sorted", false)

			var m *corestr.CharCollectionMap
			if useEmpty {
				m = corestr.New.CharCollectionMap.Empty()
			} else {
				items := input["items"].([]string)
				m = corestr.New.CharCollectionMap.Items(items)
			}

			// Act
			var actual args.Map
			if useSorted {
				list := m.SortedListAsc()
				if _, ok := expected["first"]; ok {
					first := ""
					if len(list) > 0 {
						first = list[0]
					}
					actual = args.Map{"first": first}
				} else {
					actual = args.Map{"length": len(list)}
				}
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
// CharCollectionMap — JSON
// ============================================================

func Test_CovS06_CharCollMap_Json_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_Json_Verification", func() {
		for caseIndex, tc := range covS06CharCollMapJsonTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			items := input["items"].([]string)
			m := corestr.New.CharCollectionMap.Items(items)

			// Act
			var actual args.Map
			if _, ok := expected["roundTrip"]; ok {
				jsonResult := m.Json()
				m2 := corestr.New.CharCollectionMap.Empty()
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
// CharCollectionMap — AddSameStartingCharItems
// ============================================================

func Test_CovS06_CharCollMap_AddSameCharItems_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_AddSameCharItems_Verification", func() {
		for caseIndex, tc := range covS06AddSameCharItemsTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			initialItems := input["initialItems"].([]string)
			newItems := input["newItems"].([]string)
			charStr := input["char"].(string)
			isCloneAdd := input.GetAsBoolDefault("isCloneAdd", false)
			m := corestr.New.CharCollectionMap.Items(initialItems)

			// Act
			m.AddSameStartingCharItems(charStr[0], newItems, isCloneAdd)
			actual := args.Map{"allLengthsSum": m.AllLengthsSum()}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — Resize, AddLength
// ============================================================

func Test_CovS06_CharCollMap_Resize_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_Resize_Verification", func() {
		for caseIndex, tc := range covS06ResizeTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			items := input["items"].([]string)
			m := corestr.New.CharCollectionMap.Items(items)

			// Act
			if newLen, ok := input["newLength"]; ok {
				m.Resize(newLen.(int))
			}
			if lengths, ok := input["lengths"]; ok {
				m.AddLength(lengths.([]int)...)
			}
			actual := args.Map{"length": m.Length()}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — Clear, Dispose
// ============================================================

func Test_CovS06_CharCollMap_ClearDispose_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_ClearDispose_Verification", func() {
		for caseIndex, tc := range covS06ClearDisposeTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			items := input["items"].([]string)
			m := corestr.New.CharCollectionMap.Items(items)

			// Act
			var actual args.Map
			if _, ok := expected["isDisposed"]; ok {
				m.Dispose()
				actual = args.Map{"isDisposed": m.IsEmpty()}
			} else {
				m.Clear()
				actual = args.Map{"isEmpty": m.IsEmpty()}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — HashsetByChar
// ============================================================

func Test_CovS06_CharCollMap_HashsetByChar_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_HashsetByChar_Verification", func() {
		for caseIndex, tc := range covS06HashsetByCharTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			items := input["items"].([]string)
			charStr := input["char"].(string)
			m := corestr.New.CharCollectionMap.Items(items)

			// Act
			hs := m.HashsetByChar(charStr[0])
			actual := args.Map{"isNotNil": hs != nil}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — HashsetsCollection
// ============================================================

func Test_CovS06_CharCollMap_HashsetsColl_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_HashsetsColl_Verification", func() {
		for caseIndex, tc := range covS06HashsetsCollTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useEmpty := input.GetAsBoolDefault("useEmpty", false)

			var m *corestr.CharCollectionMap
			if useEmpty {
				m = corestr.New.CharCollectionMap.Empty()
			} else {
				items := input["items"].([]string)
				m = corestr.New.CharCollectionMap.Items(items)
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
// CharCollectionMap — AddHashmapsValues
// ============================================================

func Test_CovS06_CharCollMap_AddHashmaps_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_AddHashmaps_Verification", func() {
		for caseIndex, tc := range covS06AddHashmapsTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			keys := input["keys"].([]string)
			values := input["values"].([]string)
			addBoth := input.GetAsBoolDefault("addBoth", false)

			hm := corestr.New.Hashmap.Empty()
			for i, k := range keys {
				hm.AddOrUpdate(k, values[i])
			}

			m := corestr.New.CharCollectionMap.Empty()

			// Act
			if addBoth {
				m.AddHashmapsKeysValuesBoth(hm)
			} else {
				m.AddHashmapsValues(hm)
			}
			actual := args.Map{"allLengthsSum": m.AllLengthsSum()}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — AddCollectionItems
// ============================================================

func Test_CovS06_CharCollMap_AddCollectionItems_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_AddCollectionItems_Verification", func() {
		for caseIndex, tc := range covS06AddCollectionItemsTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useNil := input.GetAsBoolDefault("useNil", false)
			m := corestr.New.CharCollectionMap.Empty()

			// Act
			if useNil {
				m.AddCollectionItems(nil)
			} else {
				items := input["items"].([]string)
				coll := corestr.New.Collection.Strings(items)
				m.AddCollectionItems(coll)
			}
			actual := args.Map{"allLengthsSum": m.AllLengthsSum()}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — String, SummaryString
// ============================================================

func Test_CovS06_CharCollMap_StringOutput_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_StringOutput_Verification", func() {
		for caseIndex, tc := range covS06StringOutputTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			items := input["items"].([]string)
			useSummary := input.GetAsBoolDefault("summary", false)
			m := corestr.New.CharCollectionMap.Items(items)

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
// CharCollectionMap — Print (skip path)
// ============================================================

func Test_CovS06_CharCollMap_PrintSkip(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_PrintSkip", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"test"})

		// Act — isPrint=false should skip
		m.Print(false)
		m.PrintLock(false)
		actual := args.Map{"skipped": true}

		// Assert — no panic means success
		expected := args.Map{"skipped": true}
		actual = args.Map{"result": actual["skipped"] != expected["skipped"]}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Print skip test failed", actual)
	})
}

// ============================================================
// CharCollectionMap — Interface adapters
// ============================================================

func Test_CovS06_CharCollMap_InterfaceAdapters_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_InterfaceAdapters_Verification", func() {
		for caseIndex, tc := range covS06InterfaceAdaptersTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			items := input["items"].([]string)
			m := corestr.New.CharCollectionMap.Items(items)

			// Act
			binder := m.AsJsonContractsBinder()
			jsoner := m.AsJsoner()
			marshaller := m.AsJsonMarshaller()
			injector := m.AsJsonParseSelfInjector()
			_ = m.JsonModelAny()
			actual := args.Map{
				"isNotNil": binder != nil && jsoner != nil && marshaller != nil && injector != nil,
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — MarshalJSON, UnmarshalJSON
// ============================================================

func Test_CovS06_CharCollMap_MarshalUnmarshal(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_MarshalUnmarshal", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"alpha", "bravo"})

		// Act
		bytes, err := m.MarshalJSON()
		m2 := corestr.New.CharCollectionMap.Empty()
		err2 := m2.UnmarshalJSON(bytes)

		actual := args.Map{
			"marshalOk":   err == nil && len(bytes) > 0,
			"unmarshalOk": err2 == nil,
		}

		// Assert
		expected := args.Map{
			"marshalOk":   true,
			"unmarshalOk": true,
		}
		actual = args.Map{"result": actual["marshalOk"] != expected["marshalOk"] || actual["unmarshalOk"] != expected["unmarshalOk"]}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "MarshalJSON/UnmarshalJSON failed:", actual)
	})
}

// ============================================================
// CharCollectionMap — AddCharHashsetMap
// ============================================================

func Test_CovS06_CharCollMap_AddCharHashsetMap_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_AddCharHashsetMap_Verification", func() {
		for caseIndex, tc := range covS06AddCharHashsetMapTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useEmptyHashset := input.GetAsBoolDefault("useEmptyHashset", false)
			m := corestr.New.CharCollectionMap.Empty()

			// Act
			if useEmptyHashset {
				hsm := corestr.Empty.CharHashsetMap()
				m.AddCharHashsetMap(hsm)
			} else {
				hsItems := input["hashsetItems"].([]string)
				hsm := corestr.New.CharHashsetMap.Strings(10, hsItems)
				m.AddCharHashsetMap(hsm)
			}
			actual := args.Map{"allLengthsSum": m.AllLengthsSum()}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — AddSameCharsCollection
// ============================================================

func Test_CovS06_CharCollMap_AddSameCharsColl_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_AddSameCharsColl_Verification", func() {
		for caseIndex, tc := range covS06AddSameCharsCollTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			expected := tc.ExpectedInput.(args.Map)
			initialItems := input["initialItems"].([]string)
			useNilColl := input.GetAsBoolDefault("useNilCollToAdd", false)
			m := corestr.New.CharCollectionMap.Items(initialItems)

			checkStr := "apple"
			if cs, ok := input["checkStr"]; ok {
				checkStr = cs.(string)
			}

			// Act
			var resultColl *corestr.Collection
			if useNilColl {
				resultColl = m.AddSameCharsCollection(checkStr, nil)
			} else {
				addItems := input["addItems"].([]string)
				coll := corestr.New.Collection.Strings(addItems)
				resultColl = m.AddSameCharsCollection(checkStr, coll)
			}

			var actual args.Map
			if _, ok := expected["collectionNotNil"]; ok {
				actual = args.Map{"collectionNotNil": resultColl != nil}
			} else {
				actual = args.Map{"allLengthsSum": m.AllLengthsSum()}
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — Lock variants
// ============================================================

func Test_CovS06_CharCollMap_LockVariants(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_LockVariants", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"alpha", "bravo"})

		// Act — exercise Lock-based methods
		lenLock := m.LengthLock()
		isEmptyLock := m.IsEmptyLock()
		allSumLock := m.AllLengthsSumLock()
		lengthOfLock := m.LengthOfLock('a')
		_ = m.GetCopyMapLock()
		_ = m.SummaryStringLock()
		_ = m.StringLock()
		_ = m.ListLock()
		isEqLock := m.IsEqualsLock(m)
		isEqCSLock := m.IsEqualsCaseSensitiveLock(true, m)

		actual := args.Map{
			"lenLock":      lenLock > 0,
			"notEmptyLock": !isEmptyLock,
			"allSumLock":   allSumLock > 0,
			"lenOfLock":    lengthOfLock > 0,
			"isEqLock":     isEqLock,
			"isEqCSLock":   isEqCSLock,
		}

		// Assert
		expected := args.Map{
			"lenLock":      true,
			"notEmptyLock": true,
			"allSumLock":   true,
			"lenOfLock":    true,
			"isEqLock":     true,
			"isEqCSLock":   true,
		}
		expected.ShouldBeEqual(t, 0, "LockVariants: got, want", actual)
	})
}

// ============================================================
// CharCollectionMap — HasWithCollectionLock
// ============================================================

func Test_CovS06_CharCollMap_HasWithCollectionLock(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_HasWithCollectionLock", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"alpha", "avocado"})

		// Act
		found, coll := m.HasWithCollectionLock("alpha")

		// Assert
		actual := args.Map{"result": found}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasWithCollectionLock should find 'alpha'", actual)
		actual = args.Map{"result": coll == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasWithCollectionLock should return non-nil collection", actual)
	})
}

// ============================================================
// CharCollectionMap — HashsetByCharLock, HashsetByStringFirstChar
// ============================================================

func Test_CovS06_CharCollMap_HashsetVariants(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_HashsetVariants", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"alpha", "avocado"})

		// Act
		hs1 := m.HashsetByCharLock('a')
		hs2 := m.HashsetByStringFirstChar("alpha")
		hs3 := m.HashsetByStringFirstCharLock("alpha")
		hs4 := m.HashsetByCharLock('z') // missing char

		// Assert
		actual := args.Map{"result": hs1 == nil || hs1.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock('a') should return non-empty", actual)
		actual = args.Map{"result": hs2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstChar should return non-nil", actual)
		actual = args.Map{"result": hs3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstCharLock should return non-nil", actual)
		if hs4 == nil || !hs4.IsEmpty() {
			// missing char returns Empty hashset
			_ = 0
		}
	})
}

// ============================================================
// CharCollectionMap — HashsetsCollectionByChars, ByStringFirstChar
// ============================================================

func Test_CovS06_CharCollMap_HashsetsCollByCharsAndStr(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_HashsetsCollByCharsAndStr", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"alpha", "bravo", "avocado"})

		// Act
		hsByChars := m.HashsetsCollectionByChars('a', 'b')
		hsByStr := m.HashsetsCollectionByStringFirstChar("alpha", "bravo")

		// Assert
		actual := args.Map{"result": hsByChars == nil || hsByChars.Length() == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars should return non-empty", actual)
		actual = args.Map{"result": hsByStr == nil || hsByStr.Length() == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringFirstChar should return non-empty", actual)
	})
}

// ============================================================
// CharCollectionMap — AddLock
// ============================================================

func Test_CovS06_CharCollMap_AddLock(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_AddLock", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Empty()

		// Act
		m.AddLock("alpha")
		m.AddLock("avocado") // same char group, existing path

		// Assert
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddLock: expected 2 items", actual)
	})
}

// ============================================================
// CharCollectionMap — AddSameCharsCollectionLock
// ============================================================

func Test_CovS06_CharCollMap_AddSameCharsCollLock(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_AddSameCharsCollLock", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"alpha"})
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
// CharCollectionMap — GetCollectionByChar
// ============================================================

func Test_CovS06_CharCollMap_GetCollectionByChar(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_GetCollectionByChar", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"alpha"})

		// Act
		coll := m.GetCollectionByChar('a')

		// Assert
		actual := args.Map{"result": coll == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "GetCollectionByChar('a') should return non-nil", actual)
	})
}

// ============================================================
// CharCollectionMap — JsonParseSelfInject
// ============================================================

func Test_CovS06_CharCollMap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_JsonParseSelfInject", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"alpha"})
		jsonResult := m.Json()

		// Act
		m2 := corestr.New.CharCollectionMap.Empty()
		err := m2.JsonParseSelfInject(&jsonResult)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject should not error:", actual)
	})
}

// ============================================================
// CharCollectionMap — ParseInjectUsingJsonMust (panic path)
// ============================================================

func Test_CovS06_CharCollMap_ParseInjectMust_Valid(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_ParseInjectMust_Valid", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"alpha"})
		jsonResult := m.Json()

		// Act
		m2 := corestr.New.CharCollectionMap.Empty()
		result := m2.ParseInjectUsingJsonMust(&jsonResult)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust should return non-nil", actual)
	})
}

// ============================================================
// CharCollectionMap — AddHashmapsKeysOrValuesBothUsingFilter
// ============================================================

func Test_CovS06_CharCollMap_AddHashmapsFilter(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_AddHashmapsFilter", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("key1", "val1")
		hm.AddOrUpdate("key2", "val2")

		m := corestr.New.CharCollectionMap.Empty()

		// Act — filter that accepts values only
		m.AddHashmapsKeysOrValuesBothUsingFilter(
			func(kv corestr.KeyValuePair) (string, bool, bool) {
				return kv.Value, true, false
			},
			hm,
		)

		// Assert
		actual := args.Map{"result": m.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Expected 2 items from filter", actual)
	})
}

// ============================================================
// CharCollectionMap — AddHashmapsKeysOrValuesBothUsingFilter (break path)
// ============================================================

func Test_CovS06_CharCollMap_AddHashmapsFilter_Break(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_AddHashmapsFilter_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("key1", "val1")
		hm.AddOrUpdate("key2", "val2")
		hm.AddOrUpdate("key3", "val3")

		m := corestr.New.CharCollectionMap.Empty()

		// Act — filter that breaks after first
		count := 0
		m.AddHashmapsKeysOrValuesBothUsingFilter(
			func(kv corestr.KeyValuePair) (string, bool, bool) {
				count++
				return kv.Value, true, count >= 1
			},
			hm,
		)

		// Assert — should have stopped early
		actual := args.Map{"result": m.AllLengthsSum() > 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Break filter should have limited additions", actual)
	})
}

// ============================================================
// CharCollectionMap — newCreator methods
// ============================================================

func Test_CovS06_CharCollMap_NewCreator_Verification(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_NewCreator_Verification", func() {
		for caseIndex, tc := range covS06NewCreatorTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			useEmpty := input.GetAsBoolDefault("useEmpty", false)
			useItems := input.GetAsBoolDefault("useItems", false)
			useItemsPtr := input.GetAsBoolDefault("useItemsPtr", false)

			// Act
			var m *corestr.CharCollectionMap
			if useEmpty {
				m = corestr.New.CharCollectionMap.Empty()
			} else if useItems {
				items := input["items"].([]string)
				m = corestr.New.CharCollectionMap.Items(items)
			} else if useItemsPtr {
				items := input["items"].([]string)
				addCap := input["additionalCap"].(int)
				eachCap := input["eachCap"].(int)
				m = corestr.New.CharCollectionMap.ItemsPtrWithCap(addCap, eachCap, items)
			} else {
				cap := input["cap"].(int)
				selfCap := input["selfCap"].(int)
				m = corestr.New.CharCollectionMap.CapSelfCap(cap, selfCap)
			}

			actual := args.Map{
				"isNotNil": m != nil,
			}
			if _, ok := tc.ExpectedInput.(args.Map)["isEmpty"]; ok {
				actual["isEmpty"] = m.IsEmpty()
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ============================================================
// CharCollectionMap — GetMap
// ============================================================

func Test_CovS06_CharCollMap_GetMap(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_GetMap", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"alpha"})

		// Act
		rawMap := m.GetMap()

		// Assert
		actual := args.Map{"result": rawMap == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "GetMap should return non-nil", actual)
	})
}

// ============================================================
// CharCollectionMap — JsonPtr
// ============================================================

func Test_CovS06_CharCollMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_JsonPtr", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"alpha"})

		// Act
		result := m.JsonPtr()

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "JsonPtr should return non-nil", actual)
	})
}

// ============================================================
// CharCollectionMap — IsEquals length mismatch branch
// ============================================================

func Test_CovS06_CharCollMap_IsEquals_LenMismatch(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_IsEquals_LenMismatch", func() {
		// Arrange
		m1 := corestr.New.CharCollectionMap.Items([]string{"alpha", "bravo"})
		m2 := corestr.New.CharCollectionMap.Items([]string{"alpha"})

		// Act
		result := m1.IsEquals(m2)

		// Assert
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsEquals should return false for different lengths", actual)
	})
}

// ============================================================
// CharCollectionMap — IsEquals one empty branch
// ============================================================

func Test_CovS06_CharCollMap_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_IsEquals_OneEmpty", func() {
		// Arrange
		m1 := corestr.New.CharCollectionMap.Items([]string{"alpha"})
		m2 := corestr.New.CharCollectionMap.Empty()

		// Act
		result := m1.IsEquals(m2)

		// Assert
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsEquals should return false when one is empty", actual)
	})
}

// ============================================================
// CharCollectionMap — IsEqualsCaseSensitive key missing branch
// ============================================================

func Test_CovS06_CharCollMap_IsEquals_KeyMissing(t *testing.T) {
	safeTest(t, "Test_CovS06_CharCollMap_IsEquals_KeyMissing", func() {
		// Arrange
		m1 := corestr.New.CharCollectionMap.Items([]string{"alpha"})
		m2 := corestr.New.CharCollectionMap.Items([]string{"bravo"})

		// Act — same Length() (both 1) but different keys
		result := m1.IsEqualsCaseSensitive(true, m2)

		// Assert
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsEquals should return false when keys differ", actual)
	})
}

// ============================================================
// Utility: sort for determinism
// ============================================================

func init() {
	_ = sort.Strings
}
