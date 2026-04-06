package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════════════
// SimpleSlice — comprehensive coverage (200 tests)
// ═══════════════════════════════════════════════════════════════════════

func Test_C29_01_SimpleSlice_Add(t *testing.T) {
	safeTest(t, "Test_C29_01_SimpleSlice_Add", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.Add("a")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_02_SimpleSlice_AddSplit(t *testing.T) {
	safeTest(t, "Test_C29_02_SimpleSlice_AddSplit", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddSplit("a,b,c", ",")

		// Act
		actual := args.Map{"result": s.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C29_03_SimpleSlice_AddIf_True(t *testing.T) {
	safeTest(t, "Test_C29_03_SimpleSlice_AddIf_True", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddIf(true, "a")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_04_SimpleSlice_AddIf_False(t *testing.T) {
	safeTest(t, "Test_C29_04_SimpleSlice_AddIf_False", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddIf(false, "a")

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_05_SimpleSlice_Adds(t *testing.T) {
	safeTest(t, "Test_C29_05_SimpleSlice_Adds", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.Adds("a", "b", "c")

		// Act
		actual := args.Map{"result": s.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C29_06_SimpleSlice_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_C29_06_SimpleSlice_Adds_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.Adds()

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_07_SimpleSlice_Append(t *testing.T) {
	safeTest(t, "Test_C29_07_SimpleSlice_Append", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.Append("a", "b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_08_SimpleSlice_Append_Empty(t *testing.T) {
	safeTest(t, "Test_C29_08_SimpleSlice_Append_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.Append()

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_09_SimpleSlice_AppendFmt(t *testing.T) {
	safeTest(t, "Test_C29_09_SimpleSlice_AppendFmt", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AppendFmt("hello %s", "world")

		// Act
		actual := args.Map{"result": s.Length() != 1 || s.First() != "hello world"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello world", actual)
	})
}

func Test_C29_10_SimpleSlice_AppendFmt_Empty(t *testing.T) {
	safeTest(t, "Test_C29_10_SimpleSlice_AppendFmt_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AppendFmt("")

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_11_SimpleSlice_AppendFmtIf_True(t *testing.T) {
	safeTest(t, "Test_C29_11_SimpleSlice_AppendFmtIf_True", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AppendFmtIf(true, "val=%d", 42)

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_12_SimpleSlice_AppendFmtIf_False(t *testing.T) {
	safeTest(t, "Test_C29_12_SimpleSlice_AppendFmtIf_False", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AppendFmtIf(false, "val=%d", 42)

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_13_SimpleSlice_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_C29_13_SimpleSlice_AddAsTitleValue", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddAsTitleValue("key", "val")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_14_SimpleSlice_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_C29_14_SimpleSlice_AddAsCurlyTitleWrap", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddAsCurlyTitleWrap("key", "val")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_15_SimpleSlice_AddAsCurlyTitleWrapIf_True(t *testing.T) {
	safeTest(t, "Test_C29_15_SimpleSlice_AddAsCurlyTitleWrapIf_True", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddAsCurlyTitleWrapIf(true, "key", "val")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_16_SimpleSlice_AddAsCurlyTitleWrapIf_False(t *testing.T) {
	safeTest(t, "Test_C29_16_SimpleSlice_AddAsCurlyTitleWrapIf_False", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddAsCurlyTitleWrapIf(false, "key", "val")

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_17_SimpleSlice_AddAsTitleValueIf_True(t *testing.T) {
	safeTest(t, "Test_C29_17_SimpleSlice_AddAsTitleValueIf_True", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddAsTitleValueIf(true, "key", "val")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_18_SimpleSlice_AddAsTitleValueIf_False(t *testing.T) {
	safeTest(t, "Test_C29_18_SimpleSlice_AddAsTitleValueIf_False", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddAsTitleValueIf(false, "key", "val")

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_19_SimpleSlice_InsertAt(t *testing.T) {
	safeTest(t, "Test_C29_19_SimpleSlice_InsertAt", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "c")
		s.InsertAt(1, "b")

		// Act
		actual := args.Map{"result": s.Length() != 3 || (*s)[1] != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b at index 1", actual)
	})
}

func Test_C29_20_SimpleSlice_InsertAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_C29_20_SimpleSlice_InsertAt_OutOfRange", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		s.InsertAt(-1, "b")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_21_SimpleSlice_InsertAt_End(t *testing.T) {
	safeTest(t, "Test_C29_21_SimpleSlice_InsertAt_End", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		s.InsertAt(1, "b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_22_SimpleSlice_AddStruct(t *testing.T) {
	safeTest(t, "Test_C29_22_SimpleSlice_AddStruct", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddStruct(true, struct{ Name string }{"test"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_23_SimpleSlice_AddStruct_Nil(t *testing.T) {
	safeTest(t, "Test_C29_23_SimpleSlice_AddStruct_Nil", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddStruct(true, nil)

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_24_SimpleSlice_AddPointer(t *testing.T) {
	safeTest(t, "Test_C29_24_SimpleSlice_AddPointer", func() {
		// Arrange
		val := "test"
		s := corestr.New.SimpleSlice.Default()
		s.AddPointer(false, &val)

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_25_SimpleSlice_AddPointer_Nil(t *testing.T) {
	safeTest(t, "Test_C29_25_SimpleSlice_AddPointer_Nil", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddPointer(false, nil)

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_26_SimpleSlice_AddsIf_True(t *testing.T) {
	safeTest(t, "Test_C29_26_SimpleSlice_AddsIf_True", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddsIf(true, "a", "b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_27_SimpleSlice_AddsIf_False(t *testing.T) {
	safeTest(t, "Test_C29_27_SimpleSlice_AddsIf_False", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddsIf(false, "a", "b")

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_28_SimpleSlice_AddError(t *testing.T) {
	safeTest(t, "Test_C29_28_SimpleSlice_AddError", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()
		s.AddError(nil)

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_29_SimpleSlice_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_C29_29_SimpleSlice_AsDefaultError", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("err1", "err2")
		err := s.AsDefaultError()

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C29_30_SimpleSlice_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_C29_30_SimpleSlice_AsError_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.AsError(",") != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C29_31_SimpleSlice_AsError_Nil(t *testing.T) {
	safeTest(t, "Test_C29_31_SimpleSlice_AsError_Nil", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": s.AsError(",") != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C29_32_SimpleSlice_First(t *testing.T) {
	safeTest(t, "Test_C29_32_SimpleSlice_First", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_33_SimpleSlice_FirstDynamic(t *testing.T) {
	safeTest(t, "Test_C29_33_SimpleSlice_FirstDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.FirstDynamic().(string) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_34_SimpleSlice_Last(t *testing.T) {
	safeTest(t, "Test_C29_34_SimpleSlice_Last", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.Last() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C29_35_SimpleSlice_LastDynamic(t *testing.T) {
	safeTest(t, "Test_C29_35_SimpleSlice_LastDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.LastDynamic().(string) != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C29_36_SimpleSlice_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_C29_36_SimpleSlice_FirstOrDefault", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.FirstOrDefault() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_37_SimpleSlice_FirstOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C29_37_SimpleSlice_FirstOrDefault_NonEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.FirstOrDefault() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_38_SimpleSlice_FirstOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_C29_38_SimpleSlice_FirstOrDefaultDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.FirstOrDefaultDynamic().(string) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_39_SimpleSlice_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_C29_39_SimpleSlice_LastOrDefault", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.LastOrDefault() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_40_SimpleSlice_LastOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C29_40_SimpleSlice_LastOrDefault_NonEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.LastOrDefault() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C29_41_SimpleSlice_LastOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_C29_41_SimpleSlice_LastOrDefaultDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.LastOrDefaultDynamic().(string) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_42_SimpleSlice_Skip(t *testing.T) {
	safeTest(t, "Test_C29_42_SimpleSlice_Skip", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		skipped := s.Skip(1)

		// Act
		actual := args.Map{"result": len(skipped) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_43_SimpleSlice_Skip_All(t *testing.T) {
	safeTest(t, "Test_C29_43_SimpleSlice_Skip_All", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		skipped := s.Skip(5)

		// Act
		actual := args.Map{"result": len(skipped) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_44_SimpleSlice_SkipDynamic(t *testing.T) {
	safeTest(t, "Test_C29_44_SimpleSlice_SkipDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		skipped := s.SkipDynamic(1)

		// Act
		actual := args.Map{"result": len(skipped.([]string)) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_45_SimpleSlice_Take(t *testing.T) {
	safeTest(t, "Test_C29_45_SimpleSlice_Take", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		taken := s.Take(2)

		// Act
		actual := args.Map{"result": len(taken) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_46_SimpleSlice_Take_All(t *testing.T) {
	safeTest(t, "Test_C29_46_SimpleSlice_Take_All", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		taken := s.Take(5)

		// Act
		actual := args.Map{"result": len(taken) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_47_SimpleSlice_TakeDynamic(t *testing.T) {
	safeTest(t, "Test_C29_47_SimpleSlice_TakeDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		taken := s.TakeDynamic(1)

		// Act
		actual := args.Map{"result": len(taken.([]string)) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_48_SimpleSlice_Limit(t *testing.T) {
	safeTest(t, "Test_C29_48_SimpleSlice_Limit", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		actual := args.Map{"result": len(s.Limit(2)) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_49_SimpleSlice_LimitDynamic(t *testing.T) {
	safeTest(t, "Test_C29_49_SimpleSlice_LimitDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		limited := s.LimitDynamic(1)

		// Act
		actual := args.Map{"result": len(limited.([]string)) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_50_SimpleSlice_Length(t *testing.T) {
	safeTest(t, "Test_C29_50_SimpleSlice_Length", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_51_SimpleSlice_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C29_51_SimpleSlice_Length_Nil", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_52_SimpleSlice_Count(t *testing.T) {
	safeTest(t, "Test_C29_52_SimpleSlice_Count", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.Count() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_53_SimpleSlice_CountFunc(t *testing.T) {
	safeTest(t, "Test_C29_53_SimpleSlice_CountFunc", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "bb", "ccc")
		count := s.CountFunc(func(i int, item string) bool {
			return len(item) > 1
		})

		// Act
		actual := args.Map{"result": count != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_54_SimpleSlice_CountFunc_Empty(t *testing.T) {
	safeTest(t, "Test_C29_54_SimpleSlice_CountFunc_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		count := s.CountFunc(func(i int, item string) bool { return true })

		// Act
		actual := args.Map{"result": count != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_55_SimpleSlice_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C29_55_SimpleSlice_IsEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_56_SimpleSlice_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_C29_56_SimpleSlice_IsEmpty_Nil", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": s.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_57_SimpleSlice_IsContains(t *testing.T) {
	safeTest(t, "Test_C29_57_SimpleSlice_IsContains", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.IsContains("b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": s.IsContains("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_58_SimpleSlice_IsContains_Empty(t *testing.T) {
	safeTest(t, "Test_C29_58_SimpleSlice_IsContains_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.IsContains("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_59_SimpleSlice_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_C29_59_SimpleSlice_IsContainsFunc", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("abc", "def")
		found := s.IsContainsFunc("de", func(item, searching string) bool {
			return len(item) > 2 && item[:2] == searching
		})

		// Act
		actual := args.Map{"result": found}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_60_SimpleSlice_IsContainsFunc_Empty(t *testing.T) {
	safeTest(t, "Test_C29_60_SimpleSlice_IsContainsFunc_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.IsContainsFunc("a", func(item, searching string) bool { return true })}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_61_SimpleSlice_IndexOf(t *testing.T) {
	safeTest(t, "Test_C29_61_SimpleSlice_IndexOf", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		actual := args.Map{"result": s.IndexOf("b") != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": s.IndexOf("z") != -1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_C29_62_SimpleSlice_IndexOf_Empty(t *testing.T) {
	safeTest(t, "Test_C29_62_SimpleSlice_IndexOf_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.IndexOf("a") != -1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_C29_63_SimpleSlice_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_C29_63_SimpleSlice_IndexOfFunc", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("aa", "bb")
		idx := s.IndexOfFunc("bb", func(item, searching string) bool {
			return item == searching
		})

		// Act
		actual := args.Map{"result": idx != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_64_SimpleSlice_IndexOfFunc_Empty(t *testing.T) {
	safeTest(t, "Test_C29_64_SimpleSlice_IndexOfFunc_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		idx := s.IndexOfFunc("a", func(item, searching string) bool { return true })

		// Act
		actual := args.Map{"result": idx != -1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_C29_65_SimpleSlice_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C29_65_SimpleSlice_HasAnyItem", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.HasAnyItem()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_66_SimpleSlice_LastIndex(t *testing.T) {
	safeTest(t, "Test_C29_66_SimpleSlice_LastIndex", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.LastIndex() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_67_SimpleSlice_HasIndex(t *testing.T) {
	safeTest(t, "Test_C29_67_SimpleSlice_HasIndex", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.HasIndex(0)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": s.HasIndex(5)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual := args.Map{"result": s.HasIndex(-1)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_68_SimpleSlice_Strings(t *testing.T) {
	safeTest(t, "Test_C29_68_SimpleSlice_Strings", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": len(s.Strings()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_69_SimpleSlice_List(t *testing.T) {
	safeTest(t, "Test_C29_69_SimpleSlice_List", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": len(s.List()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_70_SimpleSlice_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_C29_70_SimpleSlice_WrapDoubleQuote", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		w := s.WrapDoubleQuote()

		// Act
		actual := args.Map{"result": w.First() != `"a"`}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_C29_71_SimpleSlice_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_C29_71_SimpleSlice_WrapSingleQuote", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		w := s.WrapSingleQuote()

		// Act
		actual := args.Map{"result": w.First() != "'a'"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_C29_72_SimpleSlice_WrapTildaQuote(t *testing.T) {
	safeTest(t, "Test_C29_72_SimpleSlice_WrapTildaQuote", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		w := s.WrapTildaQuote()

		// Act
		actual := args.Map{"result": w.First() != "`a`"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_C29_73_SimpleSlice_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_C29_73_SimpleSlice_WrapDoubleQuoteIfMissing", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", `"b"`)
		w := s.WrapDoubleQuoteIfMissing()

		// Act
		actual := args.Map{"result": w.First() != `"a"`}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_C29_74_SimpleSlice_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_C29_74_SimpleSlice_WrapSingleQuoteIfMissing", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		w := s.WrapSingleQuoteIfMissing()

		// Act
		actual := args.Map{"result": w.First() != "'a'"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_C29_75_SimpleSlice_Transpile(t *testing.T) {
	safeTest(t, "Test_C29_75_SimpleSlice_Transpile", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.Transpile(func(s string) string { return s + "!" })

		// Act
		actual := args.Map{"result": result.First() != "a!"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a!", actual)
	})
}

func Test_C29_76_SimpleSlice_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_C29_76_SimpleSlice_Transpile_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		result := s.Transpile(func(s string) string { return s })

		// Act
		actual := args.Map{"result": result.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_77_SimpleSlice_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_C29_77_SimpleSlice_TranspileJoin", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.TranspileJoin(func(s string) string { return s + "!" }, ",")

		// Act
		actual := args.Map{"result": result != "a!,b!"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a!,b!", actual)
	})
}

func Test_C29_78_SimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_C29_78_SimpleSlice_Hashset", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "a")
		hs := s.Hashset()

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_79_SimpleSlice_Join(t *testing.T) {
	safeTest(t, "Test_C29_79_SimpleSlice_Join", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.Join(",") != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
	})
}

func Test_C29_80_SimpleSlice_Join_Empty(t *testing.T) {
	safeTest(t, "Test_C29_80_SimpleSlice_Join_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.Join(",") != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_81_SimpleSlice_JoinLine(t *testing.T) {
	safeTest(t, "Test_C29_81_SimpleSlice_JoinLine", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.JoinLine() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_82_SimpleSlice_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_C29_82_SimpleSlice_JoinLine_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.JoinLine() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_83_SimpleSlice_JoinLineEofLine(t *testing.T) {
	safeTest(t, "Test_C29_83_SimpleSlice_JoinLineEofLine", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinLineEofLine()

		// Act
		actual := args.Map{"result": result == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_84_SimpleSlice_JoinLineEofLine_Empty(t *testing.T) {
	safeTest(t, "Test_C29_84_SimpleSlice_JoinLineEofLine_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.JoinLineEofLine() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_85_SimpleSlice_JoinLineEofLine_AlreadyHasNewline(t *testing.T) {
	safeTest(t, "Test_C29_85_SimpleSlice_JoinLineEofLine_AlreadyHasNewline", func() {
		s := corestr.New.SimpleSlice.Lines("a\n")
		result := s.JoinLineEofLine()
		_ = result
	})
}

func Test_C29_86_SimpleSlice_JoinSpace(t *testing.T) {
	safeTest(t, "Test_C29_86_SimpleSlice_JoinSpace", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.JoinSpace() != "a b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a b", actual)
	})
}

func Test_C29_87_SimpleSlice_JoinComma(t *testing.T) {
	safeTest(t, "Test_C29_87_SimpleSlice_JoinComma", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.JoinComma() != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
	})
}

func Test_C29_88_SimpleSlice_JoinCsv(t *testing.T) {
	safeTest(t, "Test_C29_88_SimpleSlice_JoinCsv", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		csv := s.JoinCsv()

		// Act
		actual := args.Map{"result": csv == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_89_SimpleSlice_JoinCsvLine(t *testing.T) {
	safeTest(t, "Test_C29_89_SimpleSlice_JoinCsvLine", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		csv := s.JoinCsvLine()

		// Act
		actual := args.Map{"result": csv == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_90_SimpleSlice_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_C29_90_SimpleSlice_EachItemSplitBy", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a,b", "c,d")
		result := s.EachItemSplitBy(",")

		// Act
		actual := args.Map{"result": result.Length() != 4}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_C29_91_SimpleSlice_PrependJoin(t *testing.T) {
	safeTest(t, "Test_C29_91_SimpleSlice_PrependJoin", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("c")
		result := s.PrependJoin(",", "a", "b")

		// Act
		actual := args.Map{"result": result != "a,b,c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b,c", actual)
	})
}

func Test_C29_92_SimpleSlice_AppendJoin(t *testing.T) {
	safeTest(t, "Test_C29_92_SimpleSlice_AppendJoin", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.AppendJoin(",", "b", "c")

		// Act
		actual := args.Map{"result": result != "a,b,c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b,c", actual)
	})
}

func Test_C29_93_SimpleSlice_PrependAppend(t *testing.T) {
	safeTest(t, "Test_C29_93_SimpleSlice_PrependAppend", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("b")
		s.PrependAppend([]string{"a"}, []string{"c"})

		// Act
		actual := args.Map{"result": s.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C29_94_SimpleSlice_PrependAppend_Empty(t *testing.T) {
	safeTest(t, "Test_C29_94_SimpleSlice_PrependAppend_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("b")
		s.PrependAppend(nil, nil)

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_95_SimpleSlice_IsEqual(t *testing.T) {
	safeTest(t, "Test_C29_95_SimpleSlice_IsEqual", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.Lines("a", "b")
		s2 := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s1.IsEqual(s2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_96_SimpleSlice_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_C29_96_SimpleSlice_IsEqual_BothNil", func() {
		// Arrange
		var s1, s2 *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": s1.IsEqual(s2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_97_SimpleSlice_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_C29_97_SimpleSlice_IsEqual_OneNil", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.IsEqual(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_98_SimpleSlice_IsEqual_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_98_SimpleSlice_IsEqual_DiffLength", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.Lines("a")
		s2 := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s1.IsEqual(s2)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_99_SimpleSlice_IsEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C29_99_SimpleSlice_IsEqual_BothEmpty", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.Empty()
		s2 := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s1.IsEqual(s2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_100_SimpleSlice_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_C29_100_SimpleSlice_IsEqualLines", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.IsEqualLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_101_SimpleSlice_IsEqualLines_BothNil(t *testing.T) {
	safeTest(t, "Test_C29_101_SimpleSlice_IsEqualLines_BothNil", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": s.IsEqualLines(nil)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_102_SimpleSlice_IsEqualLines_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_102_SimpleSlice_IsEqualLines_DiffLength", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.IsEqualLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_103_SimpleSlice_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_C29_103_SimpleSlice_IsEqualUnorderedLines", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": s.IsEqualUnorderedLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_104_SimpleSlice_IsEqualUnorderedLines_BothNil(t *testing.T) {
	safeTest(t, "Test_C29_104_SimpleSlice_IsEqualUnorderedLines_BothNil", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": s.IsEqualUnorderedLines(nil)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_105_SimpleSlice_IsEqualUnorderedLines_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_105_SimpleSlice_IsEqualUnorderedLines_DiffLength", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.IsEqualUnorderedLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_106_SimpleSlice_IsEqualUnorderedLines_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C29_106_SimpleSlice_IsEqualUnorderedLines_BothEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.IsEqualUnorderedLines([]string{})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_107_SimpleSlice_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_C29_107_SimpleSlice_IsEqualUnorderedLinesClone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": s.IsEqualUnorderedLinesClone([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_108_SimpleSlice_IsEqualUnorderedLinesClone_BothNil(t *testing.T) {
	safeTest(t, "Test_C29_108_SimpleSlice_IsEqualUnorderedLinesClone_BothNil", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": s.IsEqualUnorderedLinesClone(nil)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_109_SimpleSlice_IsEqualUnorderedLinesClone_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_109_SimpleSlice_IsEqualUnorderedLinesClone_DiffLength", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.IsEqualUnorderedLinesClone([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_110_SimpleSlice_IsEqualUnorderedLinesClone_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C29_110_SimpleSlice_IsEqualUnorderedLinesClone_BothEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.IsEqualUnorderedLinesClone([]string{})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_111_SimpleSlice_Collection(t *testing.T) {
	safeTest(t, "Test_C29_111_SimpleSlice_Collection", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		c := s.Collection(true)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_112_SimpleSlice_NonPtr(t *testing.T) {
	safeTest(t, "Test_C29_112_SimpleSlice_NonPtr", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		np := s.NonPtr()

		// Act
		actual := args.Map{"result": np.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_113_SimpleSlice_Ptr(t *testing.T) {
	safeTest(t, "Test_C29_113_SimpleSlice_Ptr", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.Ptr() != s}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same", actual)
	})
}

func Test_C29_114_SimpleSlice_String(t *testing.T) {
	safeTest(t, "Test_C29_114_SimpleSlice_String", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_115_SimpleSlice_String_Empty(t *testing.T) {
	safeTest(t, "Test_C29_115_SimpleSlice_String_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.String() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_116_SimpleSlice_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_C29_116_SimpleSlice_ConcatNewSimpleSlices", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.Lines("a")
		s2 := corestr.New.SimpleSlice.Lines("b")
		result := s1.ConcatNewSimpleSlices(s2)

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_117_SimpleSlice_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_C29_117_SimpleSlice_ConcatNewStrings", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.ConcatNewStrings("b", "c")

		// Act
		actual := args.Map{"result": len(result) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C29_118_SimpleSlice_ConcatNewStrings_Nil(t *testing.T) {
	safeTest(t, "Test_C29_118_SimpleSlice_ConcatNewStrings_Nil", func() {
		// Arrange
		var s *corestr.SimpleSlice
		result := s.ConcatNewStrings("b")

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_119_SimpleSlice_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C29_119_SimpleSlice_ConcatNew", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.ConcatNew("b")

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_120_SimpleSlice_ToCollection(t *testing.T) {
	safeTest(t, "Test_C29_120_SimpleSlice_ToCollection", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		c := s.ToCollection(false)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_121_SimpleSlice_CsvStrings(t *testing.T) {
	safeTest(t, "Test_C29_121_SimpleSlice_CsvStrings", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		csv := s.CsvStrings()

		// Act
		actual := args.Map{"result": len(csv) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_122_SimpleSlice_CsvStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C29_122_SimpleSlice_CsvStrings_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		csv := s.CsvStrings()

		// Act
		actual := args.Map{"result": len(csv) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_123_SimpleSlice_JoinCsvString(t *testing.T) {
	safeTest(t, "Test_C29_123_SimpleSlice_JoinCsvString", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.JoinCsvString(",") == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_124_SimpleSlice_JoinCsvString_Empty(t *testing.T) {
	safeTest(t, "Test_C29_124_SimpleSlice_JoinCsvString_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.JoinCsvString(",") != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_125_SimpleSlice_JoinWith(t *testing.T) {
	safeTest(t, "Test_C29_125_SimpleSlice_JoinWith", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinWith(",")

		// Act
		actual := args.Map{"result": result == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_126_SimpleSlice_JoinWith_Empty(t *testing.T) {
	safeTest(t, "Test_C29_126_SimpleSlice_JoinWith_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.JoinWith(",") != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_127_SimpleSlice_JsonModel(t *testing.T) {
	safeTest(t, "Test_C29_127_SimpleSlice_JsonModel", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": len(s.JsonModel()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_128_SimpleSlice_Sort(t *testing.T) {
	safeTest(t, "Test_C29_128_SimpleSlice_Sort", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("c", "a", "b")
		s.Sort()

		// Act
		actual := args.Map{"result": s.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_129_SimpleSlice_Reverse(t *testing.T) {
	safeTest(t, "Test_C29_129_SimpleSlice_Reverse", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		s.Reverse()

		// Act
		actual := args.Map{"result": s.First() != "c" || s.Last() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected reversed", actual)
	})
}

func Test_C29_130_SimpleSlice_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_C29_130_SimpleSlice_Reverse_Two", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		s.Reverse()

		// Act
		actual := args.Map{"result": s.First() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C29_131_SimpleSlice_Reverse_One(t *testing.T) {
	safeTest(t, "Test_C29_131_SimpleSlice_Reverse_One", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		s.Reverse()

		// Act
		actual := args.Map{"result": s.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_132_SimpleSlice_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C29_132_SimpleSlice_JsonModelAny", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C29_133_SimpleSlice_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C29_133_SimpleSlice_MarshalJSON", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		b, err := s.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil || len(b) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_C29_134_SimpleSlice_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C29_134_SimpleSlice_UnmarshalJSON", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		b, _ := s.MarshalJSON()
		s2 := corestr.New.SimpleSlice.Empty()
		err := s2.UnmarshalJSON(b)

		// Act
		actual := args.Map{"result": err != nil || s2.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_135_SimpleSlice_Json(t *testing.T) {
	safeTest(t, "Test_C29_135_SimpleSlice_Json", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		j := s.Json()
		_ = j
	})
}

func Test_C29_136_SimpleSlice_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C29_136_SimpleSlice_JsonPtr", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C29_137_SimpleSlice_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C29_137_SimpleSlice_ParseInjectUsingJson", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		jp := s.JsonPtr()
		s2 := corestr.New.SimpleSlice.Empty()
		_, err := s2.ParseInjectUsingJson(jp)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_C29_138_SimpleSlice_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C29_138_SimpleSlice_ParseInjectUsingJsonMust", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		jp := s.JsonPtr()
		s2 := corestr.New.SimpleSlice.Empty()
		result := s2.ParseInjectUsingJsonMust(jp)

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_139_SimpleSlice_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C29_139_SimpleSlice_AsJsonContractsBinder", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.AsJsonContractsBinder() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C29_140_SimpleSlice_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C29_140_SimpleSlice_AsJsoner", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.AsJsoner() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C29_141_SimpleSlice_ToPtr(t *testing.T) {
	safeTest(t, "Test_C29_141_SimpleSlice_ToPtr", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		p := s.ToPtr()

		// Act
		actual := args.Map{"result": p.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_142_SimpleSlice_ToNonPtr(t *testing.T) {
	safeTest(t, "Test_C29_142_SimpleSlice_ToNonPtr", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		np := s.ToNonPtr()

		// Act
		actual := args.Map{"result": np.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_143_SimpleSlice_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C29_143_SimpleSlice_JsonParseSelfInject", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		jp := s.JsonPtr()
		s2 := corestr.New.SimpleSlice.Empty()
		err := s2.JsonParseSelfInject(jp)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_C29_144_SimpleSlice_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C29_144_SimpleSlice_AsJsonParseSelfInjector", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.AsJsonParseSelfInjector() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C29_145_SimpleSlice_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C29_145_SimpleSlice_AsJsonMarshaller", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.AsJsonMarshaller() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C29_146_SimpleSlice_Clear(t *testing.T) {
	safeTest(t, "Test_C29_146_SimpleSlice_Clear", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		s.Clear()

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_147_SimpleSlice_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_C29_147_SimpleSlice_Clear_Nil", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": s.Clear() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C29_148_SimpleSlice_Dispose(t *testing.T) {
	safeTest(t, "Test_C29_148_SimpleSlice_Dispose", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		s.Dispose()
	})
}

func Test_C29_149_SimpleSlice_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C29_149_SimpleSlice_Dispose_Nil", func() {
		var s *corestr.SimpleSlice
		s.Dispose()
	})
}

func Test_C29_150_SimpleSlice_Clone(t *testing.T) {
	safeTest(t, "Test_C29_150_SimpleSlice_Clone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		c := s.Clone(true)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_151_SimpleSlice_ClonePtr(t *testing.T) {
	safeTest(t, "Test_C29_151_SimpleSlice_ClonePtr", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		c := s.ClonePtr(true)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_152_SimpleSlice_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_C29_152_SimpleSlice_ClonePtr_Nil", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": s.ClonePtr(true) != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C29_153_SimpleSlice_DeepClone(t *testing.T) {
	safeTest(t, "Test_C29_153_SimpleSlice_DeepClone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		c := s.DeepClone()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_154_SimpleSlice_ShadowClone(t *testing.T) {
	safeTest(t, "Test_C29_154_SimpleSlice_ShadowClone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		c := s.ShadowClone()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_155_SimpleSlice_IsDistinctEqualRaw(t *testing.T) {
	safeTest(t, "Test_C29_155_SimpleSlice_IsDistinctEqualRaw", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.IsDistinctEqualRaw("b", "a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_156_SimpleSlice_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_C29_156_SimpleSlice_IsDistinctEqual", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.Lines("a", "b")
		s2 := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": s1.IsDistinctEqual(s2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_157_SimpleSlice_IsUnorderedEqualRaw_Clone(t *testing.T) {
	safeTest(t, "Test_C29_157_SimpleSlice_IsUnorderedEqualRaw_Clone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": s.IsUnorderedEqualRaw(true, "a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_158_SimpleSlice_IsUnorderedEqualRaw_NoClone(t *testing.T) {
	safeTest(t, "Test_C29_158_SimpleSlice_IsUnorderedEqualRaw_NoClone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": s.IsUnorderedEqualRaw(false, "a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_159_SimpleSlice_IsUnorderedEqualRaw_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_159_SimpleSlice_IsUnorderedEqualRaw_DiffLength", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.IsUnorderedEqualRaw(true, "a", "b")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_160_SimpleSlice_IsUnorderedEqualRaw_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C29_160_SimpleSlice_IsUnorderedEqualRaw_BothEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.IsUnorderedEqualRaw(true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_161_SimpleSlice_IsUnorderedEqual(t *testing.T) {
	safeTest(t, "Test_C29_161_SimpleSlice_IsUnorderedEqual", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.Lines("b", "a")
		s2 := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s1.IsUnorderedEqual(true, s2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_162_SimpleSlice_IsUnorderedEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C29_162_SimpleSlice_IsUnorderedEqual_BothEmpty", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.Empty()
		s2 := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s1.IsUnorderedEqual(true, s2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_163_SimpleSlice_IsUnorderedEqual_RightNil(t *testing.T) {
	safeTest(t, "Test_C29_163_SimpleSlice_IsUnorderedEqual_RightNil", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.IsUnorderedEqual(true, nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_164_SimpleSlice_IsEqualByFunc(t *testing.T) {
	safeTest(t, "Test_C29_164_SimpleSlice_IsEqualByFunc", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b")

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_165_SimpleSlice_IsEqualByFunc_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_165_SimpleSlice_IsEqualByFunc_DiffLength", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.IsEqualByFunc(func(i int, l, r string) bool { return true }, "a", "b")

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_166_SimpleSlice_IsEqualByFunc_Empty(t *testing.T) {
	safeTest(t, "Test_C29_166_SimpleSlice_IsEqualByFunc_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		result := s.IsEqualByFunc(func(i int, l, r string) bool { return true })

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_167_SimpleSlice_IsEqualByFunc_Mismatch(t *testing.T) {
	safeTest(t, "Test_C29_167_SimpleSlice_IsEqualByFunc_Mismatch", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "x")

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_168_SimpleSlice_IsEqualByFuncLinesSplit(t *testing.T) {
	safeTest(t, "Test_C29_168_SimpleSlice_IsEqualByFuncLinesSplit", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return l == r })

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_169_SimpleSlice_IsEqualByFuncLinesSplit_Trim(t *testing.T) {
	safeTest(t, "Test_C29_169_SimpleSlice_IsEqualByFuncLinesSplit_Trim", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines(" a ", " b ")
		result := s.IsEqualByFuncLinesSplit(true, ",", "a,b", func(i int, l, r string) bool { return l == r })

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_170_SimpleSlice_IsEqualByFuncLinesSplit_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_170_SimpleSlice_IsEqualByFuncLinesSplit_DiffLength", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return true })

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_171_SimpleSlice_IsEqualByFuncLinesSplit_Empty(t *testing.T) {
	safeTest(t, "Test_C29_171_SimpleSlice_IsEqualByFuncLinesSplit_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		result := s.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool { return true })

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false - empty slice vs split of empty string", actual)
	})
}

func Test_C29_172_SimpleSlice_DistinctDiffRaw(t *testing.T) {
	safeTest(t, "Test_C29_172_SimpleSlice_DistinctDiffRaw", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		diff := s.DistinctDiffRaw("b", "c")
		_ = diff
	})
}

func Test_C29_173_SimpleSlice_DistinctDiffRaw_LeftNil(t *testing.T) {
	safeTest(t, "Test_C29_173_SimpleSlice_DistinctDiffRaw_LeftNil", func() {
		// Arrange
		var s *corestr.SimpleSlice
		diff := s.DistinctDiffRaw("a")

		// Act
		actual := args.Map{"result": len(diff) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_174_SimpleSlice_DistinctDiffRaw_RightNil(t *testing.T) {
	safeTest(t, "Test_C29_174_SimpleSlice_DistinctDiffRaw_RightNil", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		diff := s.DistinctDiffRaw()
		_ = diff
	})
}

func Test_C29_175_SimpleSlice_DistinctDiffRaw_BothNil(t *testing.T) {
	safeTest(t, "Test_C29_175_SimpleSlice_DistinctDiffRaw_BothNil", func() {
		// Arrange
		var s *corestr.SimpleSlice
		diff := s.DistinctDiffRaw()

		// Act
		actual := args.Map{"result": len(diff) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_176_SimpleSlice_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_C29_176_SimpleSlice_DistinctDiff", func() {
		s1 := corestr.New.SimpleSlice.Lines("a", "b")
		s2 := corestr.New.SimpleSlice.Lines("b", "c")
		diff := s1.DistinctDiff(s2)
		_ = diff
	})
}

func Test_C29_177_SimpleSlice_DistinctDiff_LeftNil(t *testing.T) {
	safeTest(t, "Test_C29_177_SimpleSlice_DistinctDiff_LeftNil", func() {
		// Arrange
		var s *corestr.SimpleSlice
		s2 := corestr.New.SimpleSlice.Lines("a")
		diff := s.DistinctDiff(s2)

		// Act
		actual := args.Map{"result": len(diff) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_178_SimpleSlice_DistinctDiff_RightNil(t *testing.T) {
	safeTest(t, "Test_C29_178_SimpleSlice_DistinctDiff_RightNil", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		diff := s.DistinctDiff(nil)

		// Act
		actual := args.Map{"result": len(diff) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_179_SimpleSlice_DistinctDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_C29_179_SimpleSlice_DistinctDiff_BothNil", func() {
		// Arrange
		var s *corestr.SimpleSlice
		diff := s.DistinctDiff(nil)

		// Act
		actual := args.Map{"result": len(diff) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_180_SimpleSlice_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_C29_180_SimpleSlice_AddedRemovedLinesDiff", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		added, removed := s.AddedRemovedLinesDiff("b", "c")
		_ = added
		_ = removed
	})
}

func Test_C29_181_SimpleSlice_AddedRemovedLinesDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_C29_181_SimpleSlice_AddedRemovedLinesDiff_BothNil", func() {
		var s *corestr.SimpleSlice
		added, removed := s.AddedRemovedLinesDiff()
		_ = added
		_ = removed
	})
}

func Test_C29_182_SimpleSlice_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_C29_182_SimpleSlice_RemoveIndexes", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result, err := s.RemoveIndexes(1)

		// Act
		actual := args.Map{"result": err != nil || result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_183_SimpleSlice_RemoveIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_C29_183_SimpleSlice_RemoveIndexes_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		_, err := s.RemoveIndexes(0)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C29_184_SimpleSlice_RemoveIndexes_InvalidIndex(t *testing.T) {
	safeTest(t, "Test_C29_184_SimpleSlice_RemoveIndexes_InvalidIndex", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		_, err := s.RemoveIndexes(5)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C29_185_SimpleSlice_Serialize(t *testing.T) {
	safeTest(t, "Test_C29_185_SimpleSlice_Serialize", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		b, err := s.Serialize()

		// Act
		actual := args.Map{"result": err != nil || len(b) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_C29_186_SimpleSlice_Deserialize(t *testing.T) {
	safeTest(t, "Test_C29_186_SimpleSlice_Deserialize", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		var target corestr.SimpleSlice
		err := s.Deserialize(&target)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_C29_187_SimpleSlice_SafeStrings(t *testing.T) {
	safeTest(t, "Test_C29_187_SimpleSlice_SafeStrings", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": len(s.SafeStrings()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_188_SimpleSlice_SafeStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C29_188_SimpleSlice_SafeStrings_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": len(s.SafeStrings()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// newSimpleSliceCreator — factory coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C29_189_Creator_Cap(t *testing.T) {
	safeTest(t, "Test_C29_189_Creator_Cap", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(10)

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_190_Creator_Cap_Negative(t *testing.T) {
	safeTest(t, "Test_C29_190_Creator_Cap_Negative", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(-1)

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_191_Creator_Default(t *testing.T) {
	safeTest(t, "Test_C29_191_Creator_Default", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_192_Creator_Lines(t *testing.T) {
	safeTest(t, "Test_C29_192_Creator_Lines", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_193_Creator_Create(t *testing.T) {
	safeTest(t, "Test_C29_193_Creator_Create", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Create([]string{"a"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_194_Creator_Strings(t *testing.T) {
	safeTest(t, "Test_C29_194_Creator_Strings", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_195_Creator_Empty(t *testing.T) {
	safeTest(t, "Test_C29_195_Creator_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_196_Creator_Split(t *testing.T) {
	safeTest(t, "Test_C29_196_Creator_Split", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Split("a,b,c", ",")

		// Act
		actual := args.Map{"result": s.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C29_197_Creator_SplitLines(t *testing.T) {
	safeTest(t, "Test_C29_197_Creator_SplitLines", func() {
		// Arrange
		s := corestr.New.SimpleSlice.SplitLines("a\nb")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_198_Creator_Direct_Clone(t *testing.T) {
	safeTest(t, "Test_C29_198_Creator_Direct_Clone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Direct(true, []string{"a"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_199_Creator_Direct_NoClone(t *testing.T) {
	safeTest(t, "Test_C29_199_Creator_Direct_NoClone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Direct(false, []string{"a"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_200_Creator_Direct_Nil(t *testing.T) {
	safeTest(t, "Test_C29_200_Creator_Direct_Nil", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Direct(true, nil)

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}
