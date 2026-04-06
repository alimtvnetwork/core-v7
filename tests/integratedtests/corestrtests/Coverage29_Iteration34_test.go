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
		s := corestr.New.SimpleSlice.Default()
		s.Add("a")
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_02_SimpleSlice_AddSplit(t *testing.T) {
	safeTest(t, "Test_C29_02_SimpleSlice_AddSplit", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddSplit("a,b,c", ",")
		actual := args.Map{"result": s.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C29_03_SimpleSlice_AddIf_True(t *testing.T) {
	safeTest(t, "Test_C29_03_SimpleSlice_AddIf_True", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddIf(true, "a")
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_04_SimpleSlice_AddIf_False(t *testing.T) {
	safeTest(t, "Test_C29_04_SimpleSlice_AddIf_False", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddIf(false, "a")
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_05_SimpleSlice_Adds(t *testing.T) {
	safeTest(t, "Test_C29_05_SimpleSlice_Adds", func() {
		s := corestr.New.SimpleSlice.Default()
		s.Adds("a", "b", "c")
		actual := args.Map{"result": s.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C29_06_SimpleSlice_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_C29_06_SimpleSlice_Adds_Empty", func() {
		s := corestr.New.SimpleSlice.Default()
		s.Adds()
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_07_SimpleSlice_Append(t *testing.T) {
	safeTest(t, "Test_C29_07_SimpleSlice_Append", func() {
		s := corestr.New.SimpleSlice.Default()
		s.Append("a", "b")
		actual := args.Map{"result": s.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_08_SimpleSlice_Append_Empty(t *testing.T) {
	safeTest(t, "Test_C29_08_SimpleSlice_Append_Empty", func() {
		s := corestr.New.SimpleSlice.Default()
		s.Append()
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_09_SimpleSlice_AppendFmt(t *testing.T) {
	safeTest(t, "Test_C29_09_SimpleSlice_AppendFmt", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AppendFmt("hello %s", "world")
		actual := args.Map{"result": s.Length() != 1 || s.First() != "hello world"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello world", actual)
	})
}

func Test_C29_10_SimpleSlice_AppendFmt_Empty(t *testing.T) {
	safeTest(t, "Test_C29_10_SimpleSlice_AppendFmt_Empty", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AppendFmt("")
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_11_SimpleSlice_AppendFmtIf_True(t *testing.T) {
	safeTest(t, "Test_C29_11_SimpleSlice_AppendFmtIf_True", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AppendFmtIf(true, "val=%d", 42)
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_12_SimpleSlice_AppendFmtIf_False(t *testing.T) {
	safeTest(t, "Test_C29_12_SimpleSlice_AppendFmtIf_False", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AppendFmtIf(false, "val=%d", 42)
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_13_SimpleSlice_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_C29_13_SimpleSlice_AddAsTitleValue", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddAsTitleValue("key", "val")
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_14_SimpleSlice_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_C29_14_SimpleSlice_AddAsCurlyTitleWrap", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddAsCurlyTitleWrap("key", "val")
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_15_SimpleSlice_AddAsCurlyTitleWrapIf_True(t *testing.T) {
	safeTest(t, "Test_C29_15_SimpleSlice_AddAsCurlyTitleWrapIf_True", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddAsCurlyTitleWrapIf(true, "key", "val")
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_16_SimpleSlice_AddAsCurlyTitleWrapIf_False(t *testing.T) {
	safeTest(t, "Test_C29_16_SimpleSlice_AddAsCurlyTitleWrapIf_False", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddAsCurlyTitleWrapIf(false, "key", "val")
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_17_SimpleSlice_AddAsTitleValueIf_True(t *testing.T) {
	safeTest(t, "Test_C29_17_SimpleSlice_AddAsTitleValueIf_True", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddAsTitleValueIf(true, "key", "val")
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_18_SimpleSlice_AddAsTitleValueIf_False(t *testing.T) {
	safeTest(t, "Test_C29_18_SimpleSlice_AddAsTitleValueIf_False", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddAsTitleValueIf(false, "key", "val")
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_19_SimpleSlice_InsertAt(t *testing.T) {
	safeTest(t, "Test_C29_19_SimpleSlice_InsertAt", func() {
		s := corestr.New.SimpleSlice.Lines("a", "c")
		s.InsertAt(1, "b")
		actual := args.Map{"result": s.Length() != 3 || (*s)[1] != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b at index 1", actual)
	})
}

func Test_C29_20_SimpleSlice_InsertAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_C29_20_SimpleSlice_InsertAt_OutOfRange", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		s.InsertAt(-1, "b")
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_21_SimpleSlice_InsertAt_End(t *testing.T) {
	safeTest(t, "Test_C29_21_SimpleSlice_InsertAt_End", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		s.InsertAt(1, "b")
		actual := args.Map{"result": s.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_22_SimpleSlice_AddStruct(t *testing.T) {
	safeTest(t, "Test_C29_22_SimpleSlice_AddStruct", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddStruct(true, struct{ Name string }{"test"})
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_23_SimpleSlice_AddStruct_Nil(t *testing.T) {
	safeTest(t, "Test_C29_23_SimpleSlice_AddStruct_Nil", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddStruct(true, nil)
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_24_SimpleSlice_AddPointer(t *testing.T) {
	safeTest(t, "Test_C29_24_SimpleSlice_AddPointer", func() {
		val := "test"
		s := corestr.New.SimpleSlice.Default()
		s.AddPointer(false, &val)
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_25_SimpleSlice_AddPointer_Nil(t *testing.T) {
	safeTest(t, "Test_C29_25_SimpleSlice_AddPointer_Nil", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddPointer(false, nil)
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_26_SimpleSlice_AddsIf_True(t *testing.T) {
	safeTest(t, "Test_C29_26_SimpleSlice_AddsIf_True", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddsIf(true, "a", "b")
		actual := args.Map{"result": s.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_27_SimpleSlice_AddsIf_False(t *testing.T) {
	safeTest(t, "Test_C29_27_SimpleSlice_AddsIf_False", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddsIf(false, "a", "b")
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_28_SimpleSlice_AddError(t *testing.T) {
	safeTest(t, "Test_C29_28_SimpleSlice_AddError", func() {
		s := corestr.New.SimpleSlice.Default()
		s.AddError(nil)
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_29_SimpleSlice_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_C29_29_SimpleSlice_AsDefaultError", func() {
		s := corestr.New.SimpleSlice.Lines("err1", "err2")
		err := s.AsDefaultError()
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C29_30_SimpleSlice_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_C29_30_SimpleSlice_AsError_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.AsError(",") != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C29_31_SimpleSlice_AsError_Nil(t *testing.T) {
	safeTest(t, "Test_C29_31_SimpleSlice_AsError_Nil", func() {
		var s *corestr.SimpleSlice
		actual := args.Map{"result": s.AsError(",") != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C29_32_SimpleSlice_First(t *testing.T) {
	safeTest(t, "Test_C29_32_SimpleSlice_First", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_33_SimpleSlice_FirstDynamic(t *testing.T) {
	safeTest(t, "Test_C29_33_SimpleSlice_FirstDynamic", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.FirstDynamic().(string) != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_34_SimpleSlice_Last(t *testing.T) {
	safeTest(t, "Test_C29_34_SimpleSlice_Last", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.Last() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C29_35_SimpleSlice_LastDynamic(t *testing.T) {
	safeTest(t, "Test_C29_35_SimpleSlice_LastDynamic", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.LastDynamic().(string) != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C29_36_SimpleSlice_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_C29_36_SimpleSlice_FirstOrDefault", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.FirstOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_37_SimpleSlice_FirstOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C29_37_SimpleSlice_FirstOrDefault_NonEmpty", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.FirstOrDefault() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_38_SimpleSlice_FirstOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_C29_38_SimpleSlice_FirstOrDefaultDynamic", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.FirstOrDefaultDynamic().(string) != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_39_SimpleSlice_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_C29_39_SimpleSlice_LastOrDefault", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.LastOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_40_SimpleSlice_LastOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C29_40_SimpleSlice_LastOrDefault_NonEmpty", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.LastOrDefault() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C29_41_SimpleSlice_LastOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_C29_41_SimpleSlice_LastOrDefaultDynamic", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.LastOrDefaultDynamic().(string) != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_42_SimpleSlice_Skip(t *testing.T) {
	safeTest(t, "Test_C29_42_SimpleSlice_Skip", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		skipped := s.Skip(1)
		actual := args.Map{"result": len(skipped) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_43_SimpleSlice_Skip_All(t *testing.T) {
	safeTest(t, "Test_C29_43_SimpleSlice_Skip_All", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		skipped := s.Skip(5)
		actual := args.Map{"result": len(skipped) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_44_SimpleSlice_SkipDynamic(t *testing.T) {
	safeTest(t, "Test_C29_44_SimpleSlice_SkipDynamic", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		skipped := s.SkipDynamic(1)
		actual := args.Map{"result": len(skipped.([]string)) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_45_SimpleSlice_Take(t *testing.T) {
	safeTest(t, "Test_C29_45_SimpleSlice_Take", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		taken := s.Take(2)
		actual := args.Map{"result": len(taken) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_46_SimpleSlice_Take_All(t *testing.T) {
	safeTest(t, "Test_C29_46_SimpleSlice_Take_All", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		taken := s.Take(5)
		actual := args.Map{"result": len(taken) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_47_SimpleSlice_TakeDynamic(t *testing.T) {
	safeTest(t, "Test_C29_47_SimpleSlice_TakeDynamic", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		taken := s.TakeDynamic(1)
		actual := args.Map{"result": len(taken.([]string)) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_48_SimpleSlice_Limit(t *testing.T) {
	safeTest(t, "Test_C29_48_SimpleSlice_Limit", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		actual := args.Map{"result": len(s.Limit(2)) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_49_SimpleSlice_LimitDynamic(t *testing.T) {
	safeTest(t, "Test_C29_49_SimpleSlice_LimitDynamic", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		limited := s.LimitDynamic(1)
		actual := args.Map{"result": len(limited.([]string)) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_50_SimpleSlice_Length(t *testing.T) {
	safeTest(t, "Test_C29_50_SimpleSlice_Length", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_51_SimpleSlice_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C29_51_SimpleSlice_Length_Nil", func() {
		var s *corestr.SimpleSlice
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_52_SimpleSlice_Count(t *testing.T) {
	safeTest(t, "Test_C29_52_SimpleSlice_Count", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.Count() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_53_SimpleSlice_CountFunc(t *testing.T) {
	safeTest(t, "Test_C29_53_SimpleSlice_CountFunc", func() {
		s := corestr.New.SimpleSlice.Lines("a", "bb", "ccc")
		count := s.CountFunc(func(i int, item string) bool {
			return len(item) > 1
		})
		actual := args.Map{"result": count != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_54_SimpleSlice_CountFunc_Empty(t *testing.T) {
	safeTest(t, "Test_C29_54_SimpleSlice_CountFunc_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		count := s.CountFunc(func(i int, item string) bool { return true })
		actual := args.Map{"result": count != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_55_SimpleSlice_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C29_55_SimpleSlice_IsEmpty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_56_SimpleSlice_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_C29_56_SimpleSlice_IsEmpty_Nil", func() {
		var s *corestr.SimpleSlice
		actual := args.Map{"result": s.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_57_SimpleSlice_IsContains(t *testing.T) {
	safeTest(t, "Test_C29_57_SimpleSlice_IsContains", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.IsContains("b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": s.IsContains("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_58_SimpleSlice_IsContains_Empty(t *testing.T) {
	safeTest(t, "Test_C29_58_SimpleSlice_IsContains_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.IsContains("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_59_SimpleSlice_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_C29_59_SimpleSlice_IsContainsFunc", func() {
		s := corestr.New.SimpleSlice.Lines("abc", "def")
		found := s.IsContainsFunc("de", func(item, searching string) bool {
			return len(item) > 2 && item[:2] == searching
		})
		actual := args.Map{"result": found}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_60_SimpleSlice_IsContainsFunc_Empty(t *testing.T) {
	safeTest(t, "Test_C29_60_SimpleSlice_IsContainsFunc_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.IsContainsFunc("a", func(item, searching string) bool { return true })}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_61_SimpleSlice_IndexOf(t *testing.T) {
	safeTest(t, "Test_C29_61_SimpleSlice_IndexOf", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		actual := args.Map{"result": s.IndexOf("b") != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": s.IndexOf("z") != -1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_C29_62_SimpleSlice_IndexOf_Empty(t *testing.T) {
	safeTest(t, "Test_C29_62_SimpleSlice_IndexOf_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.IndexOf("a") != -1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_C29_63_SimpleSlice_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_C29_63_SimpleSlice_IndexOfFunc", func() {
		s := corestr.New.SimpleSlice.Lines("aa", "bb")
		idx := s.IndexOfFunc("bb", func(item, searching string) bool {
			return item == searching
		})
		actual := args.Map{"result": idx != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_64_SimpleSlice_IndexOfFunc_Empty(t *testing.T) {
	safeTest(t, "Test_C29_64_SimpleSlice_IndexOfFunc_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		idx := s.IndexOfFunc("a", func(item, searching string) bool { return true })
		actual := args.Map{"result": idx != -1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_C29_65_SimpleSlice_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C29_65_SimpleSlice_HasAnyItem", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_66_SimpleSlice_LastIndex(t *testing.T) {
	safeTest(t, "Test_C29_66_SimpleSlice_LastIndex", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.LastIndex() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_67_SimpleSlice_HasIndex(t *testing.T) {
	safeTest(t, "Test_C29_67_SimpleSlice_HasIndex", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.HasIndex(0)}
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
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": len(s.Strings()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_69_SimpleSlice_List(t *testing.T) {
	safeTest(t, "Test_C29_69_SimpleSlice_List", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": len(s.List()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_70_SimpleSlice_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_C29_70_SimpleSlice_WrapDoubleQuote", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		w := s.WrapDoubleQuote()
		actual := args.Map{"result": w.First() != `"a"`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_C29_71_SimpleSlice_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_C29_71_SimpleSlice_WrapSingleQuote", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		w := s.WrapSingleQuote()
		actual := args.Map{"result": w.First() != "'a'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_C29_72_SimpleSlice_WrapTildaQuote(t *testing.T) {
	safeTest(t, "Test_C29_72_SimpleSlice_WrapTildaQuote", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		w := s.WrapTildaQuote()
		actual := args.Map{"result": w.First() != "`a`"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_C29_73_SimpleSlice_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_C29_73_SimpleSlice_WrapDoubleQuoteIfMissing", func() {
		s := corestr.New.SimpleSlice.Lines("a", `"b"`)
		w := s.WrapDoubleQuoteIfMissing()
		actual := args.Map{"result": w.First() != `"a"`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_C29_74_SimpleSlice_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_C29_74_SimpleSlice_WrapSingleQuoteIfMissing", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		w := s.WrapSingleQuoteIfMissing()
		actual := args.Map{"result": w.First() != "'a'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_C29_75_SimpleSlice_Transpile(t *testing.T) {
	safeTest(t, "Test_C29_75_SimpleSlice_Transpile", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.Transpile(func(s string) string { return s + "!" })
		actual := args.Map{"result": result.First() != "a!"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a!", actual)
	})
}

func Test_C29_76_SimpleSlice_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_C29_76_SimpleSlice_Transpile_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		result := s.Transpile(func(s string) string { return s })
		actual := args.Map{"result": result.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_77_SimpleSlice_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_C29_77_SimpleSlice_TranspileJoin", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.TranspileJoin(func(s string) string { return s + "!" }, ",")
		actual := args.Map{"result": result != "a!,b!"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a!,b!", actual)
	})
}

func Test_C29_78_SimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_C29_78_SimpleSlice_Hashset", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "a")
		hs := s.Hashset()
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_79_SimpleSlice_Join(t *testing.T) {
	safeTest(t, "Test_C29_79_SimpleSlice_Join", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.Join(",") != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
	})
}

func Test_C29_80_SimpleSlice_Join_Empty(t *testing.T) {
	safeTest(t, "Test_C29_80_SimpleSlice_Join_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.Join(",") != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_81_SimpleSlice_JoinLine(t *testing.T) {
	safeTest(t, "Test_C29_81_SimpleSlice_JoinLine", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.JoinLine() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_82_SimpleSlice_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_C29_82_SimpleSlice_JoinLine_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.JoinLine() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_83_SimpleSlice_JoinLineEofLine(t *testing.T) {
	safeTest(t, "Test_C29_83_SimpleSlice_JoinLineEofLine", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinLineEofLine()
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_84_SimpleSlice_JoinLineEofLine_Empty(t *testing.T) {
	safeTest(t, "Test_C29_84_SimpleSlice_JoinLineEofLine_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.JoinLineEofLine() != ""}
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
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.JoinSpace() != "a b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a b", actual)
	})
}

func Test_C29_87_SimpleSlice_JoinComma(t *testing.T) {
	safeTest(t, "Test_C29_87_SimpleSlice_JoinComma", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.JoinComma() != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
	})
}

func Test_C29_88_SimpleSlice_JoinCsv(t *testing.T) {
	safeTest(t, "Test_C29_88_SimpleSlice_JoinCsv", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		csv := s.JoinCsv()
		actual := args.Map{"result": csv == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_89_SimpleSlice_JoinCsvLine(t *testing.T) {
	safeTest(t, "Test_C29_89_SimpleSlice_JoinCsvLine", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		csv := s.JoinCsvLine()
		actual := args.Map{"result": csv == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_90_SimpleSlice_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_C29_90_SimpleSlice_EachItemSplitBy", func() {
		s := corestr.New.SimpleSlice.Lines("a,b", "c,d")
		result := s.EachItemSplitBy(",")
		actual := args.Map{"result": result.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_C29_91_SimpleSlice_PrependJoin(t *testing.T) {
	safeTest(t, "Test_C29_91_SimpleSlice_PrependJoin", func() {
		s := corestr.New.SimpleSlice.Lines("c")
		result := s.PrependJoin(",", "a", "b")
		actual := args.Map{"result": result != "a,b,c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b,c", actual)
	})
}

func Test_C29_92_SimpleSlice_AppendJoin(t *testing.T) {
	safeTest(t, "Test_C29_92_SimpleSlice_AppendJoin", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.AppendJoin(",", "b", "c")
		actual := args.Map{"result": result != "a,b,c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b,c", actual)
	})
}

func Test_C29_93_SimpleSlice_PrependAppend(t *testing.T) {
	safeTest(t, "Test_C29_93_SimpleSlice_PrependAppend", func() {
		s := corestr.New.SimpleSlice.Lines("b")
		s.PrependAppend([]string{"a"}, []string{"c"})
		actual := args.Map{"result": s.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C29_94_SimpleSlice_PrependAppend_Empty(t *testing.T) {
	safeTest(t, "Test_C29_94_SimpleSlice_PrependAppend_Empty", func() {
		s := corestr.New.SimpleSlice.Lines("b")
		s.PrependAppend(nil, nil)
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_95_SimpleSlice_IsEqual(t *testing.T) {
	safeTest(t, "Test_C29_95_SimpleSlice_IsEqual", func() {
		s1 := corestr.New.SimpleSlice.Lines("a", "b")
		s2 := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s1.IsEqual(s2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_96_SimpleSlice_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_C29_96_SimpleSlice_IsEqual_BothNil", func() {
		var s1, s2 *corestr.SimpleSlice
		actual := args.Map{"result": s1.IsEqual(s2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_97_SimpleSlice_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_C29_97_SimpleSlice_IsEqual_OneNil", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.IsEqual(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_98_SimpleSlice_IsEqual_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_98_SimpleSlice_IsEqual_DiffLength", func() {
		s1 := corestr.New.SimpleSlice.Lines("a")
		s2 := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s1.IsEqual(s2)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_99_SimpleSlice_IsEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C29_99_SimpleSlice_IsEqual_BothEmpty", func() {
		s1 := corestr.New.SimpleSlice.Empty()
		s2 := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s1.IsEqual(s2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_100_SimpleSlice_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_C29_100_SimpleSlice_IsEqualLines", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.IsEqualLines([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_101_SimpleSlice_IsEqualLines_BothNil(t *testing.T) {
	safeTest(t, "Test_C29_101_SimpleSlice_IsEqualLines_BothNil", func() {
		var s *corestr.SimpleSlice
		actual := args.Map{"result": s.IsEqualLines(nil)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_102_SimpleSlice_IsEqualLines_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_102_SimpleSlice_IsEqualLines_DiffLength", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.IsEqualLines([]string{"a", "b"})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_103_SimpleSlice_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_C29_103_SimpleSlice_IsEqualUnorderedLines", func() {
		s := corestr.New.SimpleSlice.Lines("b", "a")
		actual := args.Map{"result": s.IsEqualUnorderedLines([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_104_SimpleSlice_IsEqualUnorderedLines_BothNil(t *testing.T) {
	safeTest(t, "Test_C29_104_SimpleSlice_IsEqualUnorderedLines_BothNil", func() {
		var s *corestr.SimpleSlice
		actual := args.Map{"result": s.IsEqualUnorderedLines(nil)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_105_SimpleSlice_IsEqualUnorderedLines_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_105_SimpleSlice_IsEqualUnorderedLines_DiffLength", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.IsEqualUnorderedLines([]string{"a", "b"})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_106_SimpleSlice_IsEqualUnorderedLines_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C29_106_SimpleSlice_IsEqualUnorderedLines_BothEmpty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.IsEqualUnorderedLines([]string{})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_107_SimpleSlice_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_C29_107_SimpleSlice_IsEqualUnorderedLinesClone", func() {
		s := corestr.New.SimpleSlice.Lines("b", "a")
		actual := args.Map{"result": s.IsEqualUnorderedLinesClone([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_108_SimpleSlice_IsEqualUnorderedLinesClone_BothNil(t *testing.T) {
	safeTest(t, "Test_C29_108_SimpleSlice_IsEqualUnorderedLinesClone_BothNil", func() {
		var s *corestr.SimpleSlice
		actual := args.Map{"result": s.IsEqualUnorderedLinesClone(nil)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_109_SimpleSlice_IsEqualUnorderedLinesClone_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_109_SimpleSlice_IsEqualUnorderedLinesClone_DiffLength", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.IsEqualUnorderedLinesClone([]string{"a", "b"})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_110_SimpleSlice_IsEqualUnorderedLinesClone_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C29_110_SimpleSlice_IsEqualUnorderedLinesClone_BothEmpty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.IsEqualUnorderedLinesClone([]string{})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_111_SimpleSlice_Collection(t *testing.T) {
	safeTest(t, "Test_C29_111_SimpleSlice_Collection", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		c := s.Collection(true)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_112_SimpleSlice_NonPtr(t *testing.T) {
	safeTest(t, "Test_C29_112_SimpleSlice_NonPtr", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		np := s.NonPtr()
		actual := args.Map{"result": np.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_113_SimpleSlice_Ptr(t *testing.T) {
	safeTest(t, "Test_C29_113_SimpleSlice_Ptr", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.Ptr() != s}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same", actual)
	})
}

func Test_C29_114_SimpleSlice_String(t *testing.T) {
	safeTest(t, "Test_C29_114_SimpleSlice_String", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_115_SimpleSlice_String_Empty(t *testing.T) {
	safeTest(t, "Test_C29_115_SimpleSlice_String_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.String() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_116_SimpleSlice_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_C29_116_SimpleSlice_ConcatNewSimpleSlices", func() {
		s1 := corestr.New.SimpleSlice.Lines("a")
		s2 := corestr.New.SimpleSlice.Lines("b")
		result := s1.ConcatNewSimpleSlices(s2)
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_117_SimpleSlice_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_C29_117_SimpleSlice_ConcatNewStrings", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.ConcatNewStrings("b", "c")
		actual := args.Map{"result": len(result) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C29_118_SimpleSlice_ConcatNewStrings_Nil(t *testing.T) {
	safeTest(t, "Test_C29_118_SimpleSlice_ConcatNewStrings_Nil", func() {
		var s *corestr.SimpleSlice
		result := s.ConcatNewStrings("b")
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_119_SimpleSlice_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C29_119_SimpleSlice_ConcatNew", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.ConcatNew("b")
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_120_SimpleSlice_ToCollection(t *testing.T) {
	safeTest(t, "Test_C29_120_SimpleSlice_ToCollection", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		c := s.ToCollection(false)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_121_SimpleSlice_CsvStrings(t *testing.T) {
	safeTest(t, "Test_C29_121_SimpleSlice_CsvStrings", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		csv := s.CsvStrings()
		actual := args.Map{"result": len(csv) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_122_SimpleSlice_CsvStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C29_122_SimpleSlice_CsvStrings_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		csv := s.CsvStrings()
		actual := args.Map{"result": len(csv) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_123_SimpleSlice_JoinCsvString(t *testing.T) {
	safeTest(t, "Test_C29_123_SimpleSlice_JoinCsvString", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.JoinCsvString(",") == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_124_SimpleSlice_JoinCsvString_Empty(t *testing.T) {
	safeTest(t, "Test_C29_124_SimpleSlice_JoinCsvString_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.JoinCsvString(",") != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_125_SimpleSlice_JoinWith(t *testing.T) {
	safeTest(t, "Test_C29_125_SimpleSlice_JoinWith", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinWith(",")
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C29_126_SimpleSlice_JoinWith_Empty(t *testing.T) {
	safeTest(t, "Test_C29_126_SimpleSlice_JoinWith_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.JoinWith(",") != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C29_127_SimpleSlice_JsonModel(t *testing.T) {
	safeTest(t, "Test_C29_127_SimpleSlice_JsonModel", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": len(s.JsonModel()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_128_SimpleSlice_Sort(t *testing.T) {
	safeTest(t, "Test_C29_128_SimpleSlice_Sort", func() {
		s := corestr.New.SimpleSlice.Lines("c", "a", "b")
		s.Sort()
		actual := args.Map{"result": s.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_129_SimpleSlice_Reverse(t *testing.T) {
	safeTest(t, "Test_C29_129_SimpleSlice_Reverse", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		s.Reverse()
		actual := args.Map{"result": s.First() != "c" || s.Last() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected reversed", actual)
	})
}

func Test_C29_130_SimpleSlice_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_C29_130_SimpleSlice_Reverse_Two", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		s.Reverse()
		actual := args.Map{"result": s.First() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C29_131_SimpleSlice_Reverse_One(t *testing.T) {
	safeTest(t, "Test_C29_131_SimpleSlice_Reverse_One", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		s.Reverse()
		actual := args.Map{"result": s.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C29_132_SimpleSlice_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C29_132_SimpleSlice_JsonModelAny", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C29_133_SimpleSlice_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C29_133_SimpleSlice_MarshalJSON", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		b, err := s.MarshalJSON()
		actual := args.Map{"result": err != nil || len(b) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_C29_134_SimpleSlice_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C29_134_SimpleSlice_UnmarshalJSON", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		b, _ := s.MarshalJSON()
		s2 := corestr.New.SimpleSlice.Empty()
		err := s2.UnmarshalJSON(b)
		actual := args.Map{"result": err != nil || s2.Length() != 1}
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
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.JsonPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C29_137_SimpleSlice_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C29_137_SimpleSlice_ParseInjectUsingJson", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		jp := s.JsonPtr()
		s2 := corestr.New.SimpleSlice.Empty()
		_, err := s2.ParseInjectUsingJson(jp)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_C29_138_SimpleSlice_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C29_138_SimpleSlice_ParseInjectUsingJsonMust", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		jp := s.JsonPtr()
		s2 := corestr.New.SimpleSlice.Empty()
		result := s2.ParseInjectUsingJsonMust(jp)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_139_SimpleSlice_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C29_139_SimpleSlice_AsJsonContractsBinder", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C29_140_SimpleSlice_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C29_140_SimpleSlice_AsJsoner", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C29_141_SimpleSlice_ToPtr(t *testing.T) {
	safeTest(t, "Test_C29_141_SimpleSlice_ToPtr", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		p := s.ToPtr()
		actual := args.Map{"result": p.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_142_SimpleSlice_ToNonPtr(t *testing.T) {
	safeTest(t, "Test_C29_142_SimpleSlice_ToNonPtr", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		np := s.ToNonPtr()
		actual := args.Map{"result": np.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_143_SimpleSlice_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C29_143_SimpleSlice_JsonParseSelfInject", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		jp := s.JsonPtr()
		s2 := corestr.New.SimpleSlice.Empty()
		err := s2.JsonParseSelfInject(jp)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_C29_144_SimpleSlice_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C29_144_SimpleSlice_AsJsonParseSelfInjector", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C29_145_SimpleSlice_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C29_145_SimpleSlice_AsJsonMarshaller", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C29_146_SimpleSlice_Clear(t *testing.T) {
	safeTest(t, "Test_C29_146_SimpleSlice_Clear", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		s.Clear()
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_147_SimpleSlice_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_C29_147_SimpleSlice_Clear_Nil", func() {
		var s *corestr.SimpleSlice
		actual := args.Map{"result": s.Clear() != nil}
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
		s := corestr.New.SimpleSlice.Lines("a", "b")
		c := s.Clone(true)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_151_SimpleSlice_ClonePtr(t *testing.T) {
	safeTest(t, "Test_C29_151_SimpleSlice_ClonePtr", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		c := s.ClonePtr(true)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_152_SimpleSlice_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_C29_152_SimpleSlice_ClonePtr_Nil", func() {
		var s *corestr.SimpleSlice
		actual := args.Map{"result": s.ClonePtr(true) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C29_153_SimpleSlice_DeepClone(t *testing.T) {
	safeTest(t, "Test_C29_153_SimpleSlice_DeepClone", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		c := s.DeepClone()
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_154_SimpleSlice_ShadowClone(t *testing.T) {
	safeTest(t, "Test_C29_154_SimpleSlice_ShadowClone", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		c := s.ShadowClone()
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_155_SimpleSlice_IsDistinctEqualRaw(t *testing.T) {
	safeTest(t, "Test_C29_155_SimpleSlice_IsDistinctEqualRaw", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.IsDistinctEqualRaw("b", "a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_156_SimpleSlice_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_C29_156_SimpleSlice_IsDistinctEqual", func() {
		s1 := corestr.New.SimpleSlice.Lines("a", "b")
		s2 := corestr.New.SimpleSlice.Lines("b", "a")
		actual := args.Map{"result": s1.IsDistinctEqual(s2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_157_SimpleSlice_IsUnorderedEqualRaw_Clone(t *testing.T) {
	safeTest(t, "Test_C29_157_SimpleSlice_IsUnorderedEqualRaw_Clone", func() {
		s := corestr.New.SimpleSlice.Lines("b", "a")
		actual := args.Map{"result": s.IsUnorderedEqualRaw(true, "a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_158_SimpleSlice_IsUnorderedEqualRaw_NoClone(t *testing.T) {
	safeTest(t, "Test_C29_158_SimpleSlice_IsUnorderedEqualRaw_NoClone", func() {
		s := corestr.New.SimpleSlice.Lines("b", "a")
		actual := args.Map{"result": s.IsUnorderedEqualRaw(false, "a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_159_SimpleSlice_IsUnorderedEqualRaw_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_159_SimpleSlice_IsUnorderedEqualRaw_DiffLength", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.IsUnorderedEqualRaw(true, "a", "b")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_160_SimpleSlice_IsUnorderedEqualRaw_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C29_160_SimpleSlice_IsUnorderedEqualRaw_BothEmpty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.IsUnorderedEqualRaw(true)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_161_SimpleSlice_IsUnorderedEqual(t *testing.T) {
	safeTest(t, "Test_C29_161_SimpleSlice_IsUnorderedEqual", func() {
		s1 := corestr.New.SimpleSlice.Lines("b", "a")
		s2 := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s1.IsUnorderedEqual(true, s2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_162_SimpleSlice_IsUnorderedEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C29_162_SimpleSlice_IsUnorderedEqual_BothEmpty", func() {
		s1 := corestr.New.SimpleSlice.Empty()
		s2 := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s1.IsUnorderedEqual(true, s2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_163_SimpleSlice_IsUnorderedEqual_RightNil(t *testing.T) {
	safeTest(t, "Test_C29_163_SimpleSlice_IsUnorderedEqual_RightNil", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": s.IsUnorderedEqual(true, nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_164_SimpleSlice_IsEqualByFunc(t *testing.T) {
	safeTest(t, "Test_C29_164_SimpleSlice_IsEqualByFunc", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b")
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_165_SimpleSlice_IsEqualByFunc_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_165_SimpleSlice_IsEqualByFunc_DiffLength", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.IsEqualByFunc(func(i int, l, r string) bool { return true }, "a", "b")
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_166_SimpleSlice_IsEqualByFunc_Empty(t *testing.T) {
	safeTest(t, "Test_C29_166_SimpleSlice_IsEqualByFunc_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		result := s.IsEqualByFunc(func(i int, l, r string) bool { return true })
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_167_SimpleSlice_IsEqualByFunc_Mismatch(t *testing.T) {
	safeTest(t, "Test_C29_167_SimpleSlice_IsEqualByFunc_Mismatch", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "x")
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_168_SimpleSlice_IsEqualByFuncLinesSplit(t *testing.T) {
	safeTest(t, "Test_C29_168_SimpleSlice_IsEqualByFuncLinesSplit", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return l == r })
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_169_SimpleSlice_IsEqualByFuncLinesSplit_Trim(t *testing.T) {
	safeTest(t, "Test_C29_169_SimpleSlice_IsEqualByFuncLinesSplit_Trim", func() {
		s := corestr.New.SimpleSlice.Lines(" a ", " b ")
		result := s.IsEqualByFuncLinesSplit(true, ",", "a,b", func(i int, l, r string) bool { return l == r })
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C29_170_SimpleSlice_IsEqualByFuncLinesSplit_DiffLength(t *testing.T) {
	safeTest(t, "Test_C29_170_SimpleSlice_IsEqualByFuncLinesSplit_DiffLength", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return true })
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C29_171_SimpleSlice_IsEqualByFuncLinesSplit_Empty(t *testing.T) {
	safeTest(t, "Test_C29_171_SimpleSlice_IsEqualByFuncLinesSplit_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		result := s.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool { return true })
		actual := args.Map{"result": result}
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
		var s *corestr.SimpleSlice
		diff := s.DistinctDiffRaw("a")
		actual := args.Map{"result": len(diff) != 1}
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
		var s *corestr.SimpleSlice
		diff := s.DistinctDiffRaw()
		actual := args.Map{"result": len(diff) != 0}
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
		var s *corestr.SimpleSlice
		s2 := corestr.New.SimpleSlice.Lines("a")
		diff := s.DistinctDiff(s2)
		actual := args.Map{"result": len(diff) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_178_SimpleSlice_DistinctDiff_RightNil(t *testing.T) {
	safeTest(t, "Test_C29_178_SimpleSlice_DistinctDiff_RightNil", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		diff := s.DistinctDiff(nil)
		actual := args.Map{"result": len(diff) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_179_SimpleSlice_DistinctDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_C29_179_SimpleSlice_DistinctDiff_BothNil", func() {
		var s *corestr.SimpleSlice
		diff := s.DistinctDiff(nil)
		actual := args.Map{"result": len(diff) != 0}
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
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result, err := s.RemoveIndexes(1)
		actual := args.Map{"result": err != nil || result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_183_SimpleSlice_RemoveIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_C29_183_SimpleSlice_RemoveIndexes_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		_, err := s.RemoveIndexes(0)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C29_184_SimpleSlice_RemoveIndexes_InvalidIndex(t *testing.T) {
	safeTest(t, "Test_C29_184_SimpleSlice_RemoveIndexes_InvalidIndex", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		_, err := s.RemoveIndexes(5)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C29_185_SimpleSlice_Serialize(t *testing.T) {
	safeTest(t, "Test_C29_185_SimpleSlice_Serialize", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		b, err := s.Serialize()
		actual := args.Map{"result": err != nil || len(b) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_C29_186_SimpleSlice_Deserialize(t *testing.T) {
	safeTest(t, "Test_C29_186_SimpleSlice_Deserialize", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		var target corestr.SimpleSlice
		err := s.Deserialize(&target)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_C29_187_SimpleSlice_SafeStrings(t *testing.T) {
	safeTest(t, "Test_C29_187_SimpleSlice_SafeStrings", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": len(s.SafeStrings()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_188_SimpleSlice_SafeStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C29_188_SimpleSlice_SafeStrings_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": len(s.SafeStrings()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// newSimpleSliceCreator — factory coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C29_189_Creator_Cap(t *testing.T) {
	safeTest(t, "Test_C29_189_Creator_Cap", func() {
		s := corestr.New.SimpleSlice.Cap(10)
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_190_Creator_Cap_Negative(t *testing.T) {
	safeTest(t, "Test_C29_190_Creator_Cap_Negative", func() {
		s := corestr.New.SimpleSlice.Cap(-1)
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_191_Creator_Default(t *testing.T) {
	safeTest(t, "Test_C29_191_Creator_Default", func() {
		s := corestr.New.SimpleSlice.Default()
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_192_Creator_Lines(t *testing.T) {
	safeTest(t, "Test_C29_192_Creator_Lines", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{"result": s.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_193_Creator_Create(t *testing.T) {
	safeTest(t, "Test_C29_193_Creator_Create", func() {
		s := corestr.New.SimpleSlice.Create([]string{"a"})
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_194_Creator_Strings(t *testing.T) {
	safeTest(t, "Test_C29_194_Creator_Strings", func() {
		s := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		actual := args.Map{"result": s.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_195_Creator_Empty(t *testing.T) {
	safeTest(t, "Test_C29_195_Creator_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C29_196_Creator_Split(t *testing.T) {
	safeTest(t, "Test_C29_196_Creator_Split", func() {
		s := corestr.New.SimpleSlice.Split("a,b,c", ",")
		actual := args.Map{"result": s.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C29_197_Creator_SplitLines(t *testing.T) {
	safeTest(t, "Test_C29_197_Creator_SplitLines", func() {
		s := corestr.New.SimpleSlice.SplitLines("a\nb")
		actual := args.Map{"result": s.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C29_198_Creator_Direct_Clone(t *testing.T) {
	safeTest(t, "Test_C29_198_Creator_Direct_Clone", func() {
		s := corestr.New.SimpleSlice.Direct(true, []string{"a"})
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_199_Creator_Direct_NoClone(t *testing.T) {
	safeTest(t, "Test_C29_199_Creator_Direct_NoClone", func() {
		s := corestr.New.SimpleSlice.Direct(false, []string{"a"})
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C29_200_Creator_Direct_Nil(t *testing.T) {
	safeTest(t, "Test_C29_200_Creator_Direct_Nil", func() {
		s := corestr.New.SimpleSlice.Direct(true, nil)
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}
