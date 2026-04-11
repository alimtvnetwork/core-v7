package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// SimpleSlice — Core operations
// =============================================================================

func Test_SS01_Add(t *testing.T) {
	safeTest(t, "Test_I8_SS01_Add", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		ss.Add("b")

		// Act
		actual := args.Map{"result": ss.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SS02_AddSplit(t *testing.T) {
	safeTest(t, "Test_I8_SS02_AddSplit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddSplit("a.b.c", ".")

		// Act
		actual := args.Map{"result": ss.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_SS03_AddIf(t *testing.T) {
	safeTest(t, "Test_I8_SS03_AddIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddIf(false, "skip")

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss.AddIf(true, "keep")
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SS04_Adds(t *testing.T) {
	safeTest(t, "Test_I8_SS04_Adds", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.Adds("a", "b")

		// Act
		actual := args.Map{"result": ss.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SS05_Append(t *testing.T) {
	safeTest(t, "Test_I8_SS05_Append", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.Append("a", "b")

		// Act
		actual := args.Map{"result": ss.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SS06_AppendFmt(t *testing.T) {
	safeTest(t, "Test_I8_SS06_AppendFmt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AppendFmt("hello %s", "world")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty format
		ss.AppendFmt("")
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 1", actual)
	})
}

func Test_SS07_AppendFmtIf(t *testing.T) {
	safeTest(t, "Test_I8_SS07_AppendFmtIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AppendFmtIf(false, "skip %s", "x")

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss.AppendFmtIf(true, "keep %s", "x")
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SS08_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_I8_SS08_AddAsTitleValue", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddAsTitleValue("Title", "Val")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SS09_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_I8_SS09_AddAsCurlyTitleWrap", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddAsCurlyTitleWrap("Title", "Val")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SS10_AddAsCurlyTitleWrapIf(t *testing.T) {
	safeTest(t, "Test_I8_SS10_AddAsCurlyTitleWrapIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddAsCurlyTitleWrapIf(false, "T", "V")

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss.AddAsCurlyTitleWrapIf(true, "T", "V")
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SS11_AddAsTitleValueIf(t *testing.T) {
	safeTest(t, "Test_I8_SS11_AddAsTitleValueIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddAsTitleValueIf(false, "T", "V")

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss.AddAsTitleValueIf(true, "T", "V")
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SS12_InsertAt(t *testing.T) {
	safeTest(t, "Test_I8_SS12_InsertAt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "c")
		ss.InsertAt(1, "b")

		// Act
		actual := args.Map{"result": ss.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// out of range
		ss.InsertAt(-1, "x")
		actual = args.Map{"result": ss.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 3", actual)
	})
}

func Test_SS13_AddStruct(t *testing.T) {
	safeTest(t, "Test_I8_SS13_AddStruct", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		type testS struct{ Name string }
		ss.AddStruct(true, testS{Name: "hello"})

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// nil
		ss.AddStruct(true, nil)
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 1", actual)
	})
}

func Test_SS14_AddPointer(t *testing.T) {
	safeTest(t, "Test_I8_SS14_AddPointer", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		s := "hello"
		ss.AddPointer(true, &s)

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ss.AddPointer(true, nil)
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 1", actual)
	})
}

func Test_SS15_AddsIf(t *testing.T) {
	safeTest(t, "Test_I8_SS15_AddsIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddsIf(false, "a", "b")

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss.AddsIf(true, "a", "b")
		actual = args.Map{"result": ss.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SS16_AddError(t *testing.T) {
	safeTest(t, "Test_I8_SS16_AddError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddError(errors.New("err"))

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ss.AddError(nil)
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 1", actual)
	})
}

func Test_SS17_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_I8_SS17_AsDefaultError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("err1")
		err := ss.AsDefaultError()

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_SS18_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_I8_SS18_AsError_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(0)
		err := ss.AsError(",")

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

// =============================================================================
// SimpleSlice — Access, Skip, Take
// =============================================================================

func Test_SS19_First(t *testing.T) {
	safeTest(t, "Test_I8_SS19_First", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_SS20_Last(t *testing.T) {
	safeTest(t, "Test_I8_SS20_Last", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.Last() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'b'", actual)
	})
}

func Test_SS21_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_I8_SS21_FirstOrDefault", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.FirstOrDefault() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
		empty := corestr.New.SimpleSlice.Cap(0)
		actual = args.Map{"result": empty.FirstOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_SS22_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_I8_SS22_LastOrDefault", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.LastOrDefault() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
		empty := corestr.New.SimpleSlice.Cap(0)
		actual = args.Map{"result": empty.LastOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_SS23_FirstDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS23_FirstDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.FirstDynamic() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_SS24_LastDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS24_LastDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.LastDynamic() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_SS25_FirstOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS25_FirstOrDefaultDynamic", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.FirstOrDefaultDynamic()
	})
}

func Test_SS26_LastOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS26_LastOrDefaultDynamic", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.LastOrDefaultDynamic()
	})
}

func Test_SS27_Skip(t *testing.T) {
	safeTest(t, "Test_I8_SS27_Skip", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result := ss.Skip(1)

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// skip more than length
		result2 := ss.Skip(10)
		actual = args.Map{"result": len(result2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SS28_SkipDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS28_SkipDynamic", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		_ = ss.SkipDynamic(1)
		_ = ss.SkipDynamic(10)
	})
}

func Test_SS29_Take(t *testing.T) {
	safeTest(t, "Test_I8_SS29_Take", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result := ss.Take(2)

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// take more
		result2 := ss.Take(10)
		actual = args.Map{"result": len(result2) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_SS30_TakeDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS30_TakeDynamic", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		_ = ss.TakeDynamic(1)
		_ = ss.TakeDynamic(10)
	})
}

func Test_SS31_Limit(t *testing.T) {
	safeTest(t, "Test_I8_SS31_Limit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.Limit(1)

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SS32_LimitDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS32_LimitDynamic", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.LimitDynamic(1)
	})
}

func Test_SS33_Count(t *testing.T) {
	safeTest(t, "Test_I8_SS33_Count", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.Count() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SS34_CountFunc(t *testing.T) {
	safeTest(t, "Test_I8_SS34_CountFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "bb", "ccc")
		c := ss.CountFunc(func(i int, s string) bool { return len(s) > 1 })

		// Act
		actual := args.Map{"result": c != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SS35_IsContains(t *testing.T) {
	safeTest(t, "Test_I8_SS35_IsContains", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.IsContains("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.IsContains("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

// =============================================================================
// SimpleSlice — JSON, String, Sort
// =============================================================================

func Test_SS36_Json(t *testing.T) {
	safeTest(t, "Test_I8_SS36_Json", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		j := ss.Json()

		// Act
		actual := args.Map{"result": j.JsonString() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_SS37_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_I8_SS37_ParseInjectUsingJson", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		jr := ss.JsonPtr()
		ss2 := corestr.New.SimpleSlice.Cap(1)
		_, err := ss2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_SS38_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_I8_SS38_ParseInjectUsingJson_Error", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(1)
		bad := corejson.NewResult.UsingString(`invalid`)
		_, err := ss.ParseInjectUsingJson(bad)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_SS39_String(t *testing.T) {
	safeTest(t, "Test_I8_SS39_String", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_SS40_Join(t *testing.T) {
	safeTest(t, "Test_I8_SS40_Join", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.Join(",") != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b'", actual)
	})
}

func Test_SS41_List(t *testing.T) {
	safeTest(t, "Test_I8_SS41_List", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": len(ss.List()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SS42_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_I8_SS42_RemoveIndexes", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result, err := ss.RemoveIndexes(1)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual = args.Map{"result": result.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SS43_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_I8_SS43_MarshalJSON", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		data, err := ss.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected marshal", actual)
	})
}

func Test_SS44_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_I8_SS44_UnmarshalJSON", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(0)
		err := ss.UnmarshalJSON([]byte(`["a","b"]`))

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual = args.Map{"result": ss.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SS45_SafeStrings(t *testing.T) {
	safeTest(t, "Test_I8_SS45_SafeStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "", "b")
		result := ss.SafeStrings()

		// Act
		actual := args.Map{"result": len(result) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_SS46_Serialize(t *testing.T) {
	safeTest(t, "Test_I8_SS46_Serialize", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		bytes, err := ss.Serialize()

		// Act
		actual := args.Map{"result": err != nil || len(bytes) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected serialization", actual)
	})
}
