package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ═══════════════════════════════════════════════════════════════
// Collection — remaining: string, join, csv, json, resize, dispose
// ═══════════════════════════════════════════════════════════════

func Test_Cov49_Collection_String(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_String", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "String", Expected: true, Actual: len(c.String()) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_String_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "String empty", Expected: true, Actual: len(c.String()) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_StringLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "StringLock", Expected: true, Actual: len(c.StringLock()) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_StringLock_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "StringLock empty", Expected: true, Actual: len(c.StringLock()) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_SummaryString", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "SummaryString", Expected: true, Actual: len(c.SummaryString(1)) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_SummaryStringWithHeader_Empty(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_SummaryStringWithHeader_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "SummaryStringWithHeader empty", Expected: true, Actual: len(c.SummaryStringWithHeader("H")) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_SummaryStringWithHeader_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_SummaryStringWithHeader_HasItems", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "SummaryStringWithHeader has", Expected: true, Actual: len(c.SummaryStringWithHeader("H")) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Join(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Join", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "Join", Expected: "a,b", Actual: c.Join(",")}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Join_Empty(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Join_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "Join empty", Expected: "", Actual: c.Join(",")}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_JoinLine(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_JoinLine", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "JoinLine", Expected: true, Actual: len(c.JoinLine()) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_JoinLine_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "JoinLine empty", Expected: "", Actual: c.JoinLine()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Joins_NoExtra(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Joins_NoExtra", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "Joins no extra", Expected: "a,b", Actual: c.Joins(",")}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Joins_WithExtra(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Joins_WithExtra", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.Joins(",", "b", "c")
		tc := caseV1Compat{Name: "Joins with extra", Expected: true, Actual: len(result) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_NonEmptyJoins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		result := c.NonEmptyJoins(",")
		tc := caseV1Compat{Name: "NonEmptyJoins", Expected: "a,b", Actual: result}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_NonWhitespaceJoins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "  ", "b"})
		result := c.NonWhitespaceJoins(",")
		tc := caseV1Compat{Name: "NonWhitespaceJoins", Expected: "a,b", Actual: result}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Csv(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Csv", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.Csv()
		tc := caseV1Compat{Name: "Csv", Expected: true, Actual: len(result) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Csv_Empty(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Csv_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "Csv empty", Expected: "", Actual: c.Csv()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_CsvOptions(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_CsvOptions", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.CsvOptions(true)
		tc := caseV1Compat{Name: "CsvOptions", Expected: true, Actual: len(result) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_CsvOptions_Empty(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_CsvOptions_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "CsvOptions empty", Expected: "", Actual: c.CsvOptions(false)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_CsvLines(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_CsvLines", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.CsvLines()
		tc := caseV1Compat{Name: "CsvLines", Expected: 2, Actual: len(result)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_CsvLinesOptions(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_CsvLinesOptions", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.CsvLinesOptions(true)
		tc := caseV1Compat{Name: "CsvLinesOptions", Expected: 1, Actual: len(result)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_GetAllExcept", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.GetAllExcept([]string{"b"})
		tc := caseV1Compat{Name: "GetAllExcept", Expected: 2, Actual: len(result)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_GetAllExcept_Nil(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_GetAllExcept_Nil", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.GetAllExcept(nil)
		tc := caseV1Compat{Name: "GetAllExcept nil", Expected: 2, Actual: len(result)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_CharCollectionMap", func() {
		c := corestr.New.Collection.Strings([]string{"hello", "hi", "abc"})
		ccm := c.CharCollectionMap()
		tc := caseV1Compat{Name: "CharCollectionMap", Expected: true, Actual: ccm != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_AddCapacity(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_AddCapacity", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddCapacity(10)
		tc := caseV1Compat{Name: "AddCapacity", Expected: true, Actual: c.Capacity() >= 10}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_AddCapacity_Empty(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_AddCapacity_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddCapacity()
		tc := caseV1Compat{Name: "AddCapacity empty", Expected: true, Actual: c != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Resize_Bigger(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Resize_Bigger", func() {
		c := corestr.New.Collection.Cap(2)
		c.Resize(20)
		tc := caseV1Compat{Name: "Resize bigger", Expected: true, Actual: c.Capacity() >= 20}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Resize_Smaller(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Resize_Smaller", func() {
		c := corestr.New.Collection.Cap(20)
		c.Resize(5)
		tc := caseV1Compat{Name: "Resize smaller noop", Expected: true, Actual: c.Capacity() >= 20}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Clear(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Clear", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Clear()
		tc := caseV1Compat{Name: "Clear", Expected: 0, Actual: c.Length()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Clear_Nil", func() {
		var c *corestr.Collection
		result := c.Clear()
		tc := caseV1Compat{Name: "Clear nil", Expected: true, Actual: result == nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Dispose", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Dispose()
		tc := caseV1Compat{Name: "Dispose", Expected: true, Actual: c.IsEmpty()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Dispose_Nil", func() {
		var c *corestr.Collection
		c.Dispose() // should not panic
		tc := caseV1Compat{Name: "Dispose nil", Expected: true, Actual: true}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_JsonModel", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "JsonModel", Expected: 1, Actual: len(c.JsonModel())}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_JsonModelAny", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "JsonModelAny", Expected: true, Actual: c.JsonModelAny() != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_MarshalJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		data, err := c.MarshalJSON()
		tc := caseV1Compat{Name: "MarshalJSON", Expected: true, Actual: err == nil && len(data) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_UnmarshalJSON", func() {
		c := &corestr.Collection{}
		err := c.UnmarshalJSON([]byte(`["a","b"]`))
		tc := caseV1Compat{Name: "UnmarshalJSON", Expected: true, Actual: err == nil}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "UnmarshalJSON length", Expected: 2, Actual: c.Length()}
		tc2.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_Serialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		data, err := c.Serialize()
		tc := caseV1Compat{Name: "Serialize", Expected: true, Actual: err == nil && len(data) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_AsJsonContractsBinder", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "AsJsonContractsBinder", Expected: true, Actual: c.AsJsonContractsBinder() != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_AsJsonMarshaller", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "AsJsonMarshaller", Expected: true, Actual: c.AsJsonMarshaller() != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_Collection_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_Cov49_Collection_ExpandSlicePlusAdd", func() {
		c := corestr.New.Collection.Cap(5)
		c.ExpandSlicePlusAdd([]string{"a,b"}, func(s string) []string {
			return []string{s + "_expanded"}
		})
		tc := caseV1Compat{Name: "ExpandSlicePlusAdd", Expected: 1, Actual: c.Length()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// CharHashsetMap — remaining: json, serialize, clear, remove
// ═══════════════════════════════════════════════════════════════

func Test_Cov49_CharHashsetMap_GetHashsetByChar(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_GetHashsetByChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		hs := chm.GetHashsetByChar('h')
		tc := caseV1Compat{Name: "CHM GetHashsetByChar", Expected: true, Actual: hs != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_HashsetByCharLock_Found(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_HashsetByCharLock_Found", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		hs := chm.HashsetByCharLock('h')
		tc := caseV1Compat{Name: "CHM HashsetByCharLock found", Expected: true, Actual: hs.Has("hello")}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_HashsetByCharLock_NotFound(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_HashsetByCharLock_NotFound", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := chm.HashsetByCharLock('z')
		tc := caseV1Compat{Name: "CHM HashsetByCharLock not found", Expected: true, Actual: hs.IsEmpty()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_HashsetByStringFirstChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		hs := chm.HashsetByStringFirstChar("hello")
		tc := caseV1Compat{Name: "CHM HashsetByStringFirstChar", Expected: true, Actual: hs != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_HashsetByStringFirstCharLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		hs := chm.HashsetByStringFirstCharLock("hello")
		tc := caseV1Compat{Name: "CHM HashsetByStringFirstCharLock", Expected: true, Actual: hs.Has("hello")}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_JsonModel", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		tc := caseV1Compat{Name: "CHM JsonModel", Expected: true, Actual: chm.JsonModel() != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_JsonModelAny", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM JsonModelAny", Expected: true, Actual: chm.JsonModelAny() != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_AsJsonContractsBinder", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM AsJsonContractsBinder", Expected: true, Actual: chm.AsJsonContractsBinder() != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_AsJsoner", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM AsJsoner", Expected: true, Actual: chm.AsJsoner() != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_AsJsonMarshaller", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM AsJsonMarshaller", Expected: true, Actual: chm.AsJsonMarshaller() != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_AsJsonParseSelfInjector", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM AsJsonParseSelfInjector", Expected: true, Actual: chm.AsJsonParseSelfInjector() != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_RemoveAll", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		chm.RemoveAll()
		tc := caseV1Compat{Name: "CHM RemoveAll", Expected: true, Actual: chm.IsEmpty()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_RemoveAll_Empty(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_RemoveAll_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.RemoveAll()
		tc := caseV1Compat{Name: "CHM RemoveAll empty", Expected: true, Actual: chm.IsEmpty()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_Clear(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_Clear", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		chm.Clear()
		tc := caseV1Compat{Name: "CHM Clear", Expected: true, Actual: chm.IsEmpty()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_StringLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		tc := caseV1Compat{Name: "CHM StringLock", Expected: true, Actual: len(chm.StringLock()) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_SummaryStringLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		tc := caseV1Compat{Name: "CHM SummaryStringLock", Expected: true, Actual: len(chm.SummaryStringLock()) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_Print_True(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_Print_True", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		chm.Print(true) // should not panic
		tc := caseV1Compat{Name: "CHM Print true", Expected: true, Actual: true}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_PrintLock_False(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_PrintLock_False", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.PrintLock(false)
		tc := caseV1Compat{Name: "CHM PrintLock false", Expected: true, Actual: true}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_PrintLock_True(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_PrintLock_True", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		chm.PrintLock(true)
		tc := caseV1Compat{Name: "CHM PrintLock true", Expected: true, Actual: true}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_AddHashsetLock_New(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_AddHashsetLock_New", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := corestr.New.Hashset.StringsSpreadItems("abc")
		result := chm.AddHashsetLock("a", hs)
		tc := caseV1Compat{Name: "CHM AddHashsetLock new", Expected: true, Actual: result != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_AddHashsetLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_AddHashsetLock_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("abc")
		hs := corestr.New.Hashset.StringsSpreadItems("axy")
		result := chm.AddHashsetLock("a", hs)
		tc := caseV1Compat{Name: "CHM AddHashsetLock existing", Expected: true, Actual: result != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_AddHashsetLock_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_AddHashsetLock_NilHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.AddHashsetLock("a", nil)
		tc := caseV1Compat{Name: "CHM AddHashsetLock nil", Expected: true, Actual: result != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_AddSameCharsCollectionLock_New(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_AddSameCharsCollectionLock_New", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		col := corestr.New.Collection.Strings([]string{"abc"})
		result := chm.AddSameCharsCollectionLock("a", col)
		tc := caseV1Compat{Name: "CHM AddSameCharsCollectionLock new", Expected: true, Actual: result != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_AddSameCharsCollectionLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_AddSameCharsCollectionLock_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("abc")
		col := corestr.New.Collection.Strings([]string{"axy"})
		result := chm.AddSameCharsCollectionLock("a", col)
		tc := caseV1Compat{Name: "CHM AddSameCharsCollectionLock existing", Expected: true, Actual: result != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_AddSameCharsCollectionLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_AddSameCharsCollectionLock_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.AddSameCharsCollectionLock("a", nil)
		tc := caseV1Compat{Name: "CHM AddSameCharsCollectionLock nil", Expected: true, Actual: result != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov49_CharHashsetMap_AddSameCharsCollectionLock_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Cov49_CharHashsetMap_AddSameCharsCollectionLock_ExistingNil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("abc")
		result := chm.AddSameCharsCollectionLock("a", nil)
		tc := caseV1Compat{Name: "CHM AddSameCharsCollectionLock existing nil", Expected: true, Actual: result != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}
