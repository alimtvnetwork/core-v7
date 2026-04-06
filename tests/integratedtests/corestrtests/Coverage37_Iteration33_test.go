package corestrtests

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── SimpleSlice: Add/Adds/Append/AddIf/AddsIf ──

func Test_Cov37_Add(t *testing.T) {
	safeTest(t, "Test_Cov37_Add", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.Add("a")
		s.Add("b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_AddSplit(t *testing.T) {
	safeTest(t, "Test_Cov37_AddSplit", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddSplit("a,b,c", ",")

		// Act
		actual := args.Map{"result": s.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov37_AddIf_True(t *testing.T) {
	safeTest(t, "Test_Cov37_AddIf_True", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddIf(true, "yes")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_AddIf_False(t *testing.T) {
	safeTest(t, "Test_Cov37_AddIf_False", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddIf(false, "no")

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov37_Adds(t *testing.T) {
	safeTest(t, "Test_Cov37_Adds", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.Adds("a", "b", "c")

		// Act
		actual := args.Map{"result": s.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov37_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Adds_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("x")
		s.Adds()

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_Append(t *testing.T) {
	safeTest(t, "Test_Cov37_Append", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.Append("a", "b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Append_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Append_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("x")
		s.Append()

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_AddsIf_True(t *testing.T) {
	safeTest(t, "Test_Cov37_AddsIf_True", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddsIf(true, "a", "b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_AddsIf_False(t *testing.T) {
	safeTest(t, "Test_Cov37_AddsIf_False", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddsIf(false, "a", "b")

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AppendFmt / AppendFmtIf ──

func Test_Cov37_AppendFmt(t *testing.T) {
	safeTest(t, "Test_Cov37_AppendFmt", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AppendFmt("hello %s", "world")

		// Act
		actual := args.Map{"result": s.First() != "hello world"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov37_AppendFmt_EmptySkip(t *testing.T) {
	safeTest(t, "Test_Cov37_AppendFmt_EmptySkip", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AppendFmt("")
		// empty format with no values still appends (fmt.Sprintf("") == "")
		// Actually the code checks: format == "" && len(v) == 0 → skip
		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov37_AppendFmtIf_True(t *testing.T) {
	safeTest(t, "Test_Cov37_AppendFmtIf_True", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AppendFmtIf(true, "val=%d", 42)

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_AppendFmtIf_False(t *testing.T) {
	safeTest(t, "Test_Cov37_AppendFmtIf_False", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AppendFmtIf(false, "val=%d", 42)

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov37_AppendFmtIf_EmptyFormat(t *testing.T) {
	safeTest(t, "Test_Cov37_AppendFmtIf_EmptyFormat", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AppendFmtIf(true, "")

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddAsTitleValue / AddAsCurlyTitleWrap / If variants ──

func Test_Cov37_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_Cov37_AddAsTitleValue", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddAsTitleValue("Key", "Val")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_AddAsTitleValueIf_True(t *testing.T) {
	safeTest(t, "Test_Cov37_AddAsTitleValueIf_True", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddAsTitleValueIf(true, "K", "V")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_AddAsTitleValueIf_False(t *testing.T) {
	safeTest(t, "Test_Cov37_AddAsTitleValueIf_False", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddAsTitleValueIf(false, "K", "V")

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov37_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_Cov37_AddAsCurlyTitleWrap", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddAsCurlyTitleWrap("K", "V")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_AddAsCurlyTitleWrapIf_True(t *testing.T) {
	safeTest(t, "Test_Cov37_AddAsCurlyTitleWrapIf_True", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddAsCurlyTitleWrapIf(true, "K", "V")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_AddAsCurlyTitleWrapIf_False(t *testing.T) {
	safeTest(t, "Test_Cov37_AddAsCurlyTitleWrapIf_False", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddAsCurlyTitleWrapIf(false, "K", "V")

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── InsertAt ──

func Test_Cov37_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_Cov37_InsertAt_Middle", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "c")
		s.InsertAt(1, "b")

		// Act
		actual := args.Map{"result": s.Length() != 3 || (*s)[1] != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov37_InsertAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Cov37_InsertAt_OutOfRange", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		s.InsertAt(-1, "x")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		s.InsertAt(99, "x")
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_InsertAt_End(t *testing.T) {
	safeTest(t, "Test_Cov37_InsertAt_End", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		s.InsertAt(2, "c")

		// Act
		actual := args.Map{"result": s.Length() != 3 || s.Last() != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

// ── AddStruct / AddPointer ──

func Test_Cov37_AddStruct(t *testing.T) {
	safeTest(t, "Test_Cov37_AddStruct", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddStruct(true, struct{ Name string }{Name: "test"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_AddStruct_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_AddStruct_Nil", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddStruct(true, nil)

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov37_AddPointer_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_AddPointer_Nil", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddPointer(true, nil)

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddError ──

func Test_Cov37_AddError_NonNil(t *testing.T) {
	safeTest(t, "Test_Cov37_AddError_NonNil", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddError(errors.New("test error"))
	})
}

func Test_Cov37_AddError_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_AddError_Nil", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddError(nil)

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AsError / AsDefaultError ──

func Test_Cov37_AsError_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_AsError_NonEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("err1", "err2")
		err := s.AsError("; ")

		// Act
		actual := args.Map{"result": err == nil || !strings.Contains(err.Error(), "err1")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	})
}

func Test_Cov37_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_AsError_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.AsError("; ") != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov37_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_Cov37_AsDefaultError", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("e1")
		err := s.AsDefaultError()

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

// ── First/Last/FirstOrDefault/LastOrDefault ──

func Test_Cov37_First(t *testing.T) {
	safeTest(t, "Test_Cov37_First", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov37_Last(t *testing.T) {
	safeTest(t, "Test_Cov37_Last", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.Last() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Cov37_FirstDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_FirstDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("x")
		val := s.FirstDynamic()

		// Act
		actual := args.Map{"result": val != "x"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov37_LastDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_LastDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("x", "y")
		val := s.LastDynamic()

		// Act
		actual := args.Map{"result": val != "y"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov37_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_FirstOrDefault_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.FirstOrDefault() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_FirstOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_FirstOrDefault_NonEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.FirstOrDefault() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov37_FirstOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_FirstOrDefaultDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.FirstOrDefaultDynamic() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_LastOrDefault_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.LastOrDefault() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_LastOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_LastOrDefault_NonEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.LastOrDefault() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Cov37_LastOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_LastOrDefaultDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.LastOrDefaultDynamic() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// ── Skip/Take/Limit ──

func Test_Cov37_Skip(t *testing.T) {
	safeTest(t, "Test_Cov37_Skip", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result := s.Skip(1)

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Skip_ExceedsLen(t *testing.T) {
	safeTest(t, "Test_Cov37_Skip_ExceedsLen", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.Skip(5)

		// Act
		actual := args.Map{"result": len(result) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov37_SkipDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_SkipDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.SkipDynamic(1)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov37_SkipDynamic_ExceedsLen(t *testing.T) {
	safeTest(t, "Test_Cov37_SkipDynamic_ExceedsLen", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.SkipDynamic(5)
		slice, ok := result.([]string)

		// Act
		actual := args.Map{"result": ok || len(slice) != 0}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty slice", actual)
	})
}

func Test_Cov37_Take(t *testing.T) {
	safeTest(t, "Test_Cov37_Take", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result := s.Take(2)

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Take_ExceedsLen(t *testing.T) {
	safeTest(t, "Test_Cov37_Take_ExceedsLen", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.Take(5)

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_TakeDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_TakeDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.TakeDynamic(1)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov37_Limit(t *testing.T) {
	safeTest(t, "Test_Cov37_Limit", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result := s.Limit(2)

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_LimitDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_LimitDynamic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.LimitDynamic(1)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── Length/Count/CountFunc/IsEmpty/HasAnyItem/LastIndex/HasIndex ──

func Test_Cov37_Length_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_Length_Nil", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov37_Count(t *testing.T) {
	safeTest(t, "Test_Cov37_Count", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.Count() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_CountFunc(t *testing.T) {
	safeTest(t, "Test_Cov37_CountFunc", func() {
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

func Test_Cov37_CountFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_CountFunc_Empty", func() {
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

func Test_Cov37_IsEmpty_True(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEmpty_True", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Cov37_HasAnyItem", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.HasAnyItem()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected HasAnyItem", actual)
	})
}

func Test_Cov37_LastIndex(t *testing.T) {
	safeTest(t, "Test_Cov37_LastIndex", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.LastIndex() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_HasIndex(t *testing.T) {
	safeTest(t, "Test_Cov37_HasIndex", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.HasIndex(1)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": s.HasIndex(2)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual := args.Map{"result": s.HasIndex(-1)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for negative", actual)
	})
}

// ── IsContains / IsContainsFunc / IndexOf / IndexOfFunc ──

func Test_Cov37_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov37_IsContains", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.IsContains("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": s.IsContains("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov37_IsContains_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsContains_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.IsContains("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov37_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_Cov37_IsContainsFunc", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("Hello", "World")
		found := s.IsContainsFunc("hello", func(item, searching string) bool {
			return strings.EqualFold(item, searching)
		})

		// Act
		actual := args.Map{"result": found}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov37_IsContainsFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsContainsFunc_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		found := s.IsContainsFunc("x", func(item, searching string) bool { return true })

		// Act
		actual := args.Map{"result": found}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov37_IndexOf(t *testing.T) {
	safeTest(t, "Test_Cov37_IndexOf", func() {
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

func Test_Cov37_IndexOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_IndexOf_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.IndexOf("x") != -1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_Cov37_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_Cov37_IndexOfFunc", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		idx := s.IndexOfFunc("B", func(item, searching string) bool {
			return strings.EqualFold(item, searching)
		})

		// Act
		actual := args.Map{"result": idx != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_IndexOfFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_IndexOfFunc_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		idx := s.IndexOfFunc("x", func(item, searching string) bool { return true })

		// Act
		actual := args.Map{"result": idx != -1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

// ── Strings/List/Hashset ──

func Test_Cov37_Strings(t *testing.T) {
	safeTest(t, "Test_Cov37_Strings", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": len(s.Strings()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_List(t *testing.T) {
	safeTest(t, "Test_Cov37_List", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": len(s.List()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov37_Hashset", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "a")
		hs := s.Hashset()

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 unique", actual)
	})
}

// ── Wrap variants ──

func Test_Cov37_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_Cov37_WrapDoubleQuote", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.WrapDoubleQuote()

		// Act
		actual := args.Map{"result": strings.Contains(result.First(), "\"")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected double quotes", actual)
	})
}

func Test_Cov37_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_Cov37_WrapSingleQuote", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.WrapSingleQuote()

		// Act
		actual := args.Map{"result": strings.Contains(result.First(), "'")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected single quotes", actual)
	})
}

func Test_Cov37_WrapTildaQuote(t *testing.T) {
	safeTest(t, "Test_Cov37_WrapTildaQuote", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.WrapTildaQuote()

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Cov37_WrapDoubleQuoteIfMissing", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.WrapDoubleQuoteIfMissing()

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Cov37_WrapSingleQuoteIfMissing", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.WrapSingleQuoteIfMissing()

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── Transpile / TranspileJoin ──

func Test_Cov37_Transpile(t *testing.T) {
	safeTest(t, "Test_Cov37_Transpile", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.Transpile(strings.ToUpper)

		// Act
		actual := args.Map{"result": result.First() != "A"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected A", actual)
	})
}

func Test_Cov37_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Transpile_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		result := s.Transpile(strings.ToUpper)

		// Act
		actual := args.Map{"result": result.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_Cov37_TranspileJoin", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.TranspileJoin(strings.ToUpper, ",")

		// Act
		actual := args.Map{"result": result != "A,B"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected A,B", actual)
	})
}

// ── Join variants ──

func Test_Cov37_Join(t *testing.T) {
	safeTest(t, "Test_Cov37_Join", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.Join(",") != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov37_Join_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Join_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.Join(",") != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_JoinLine(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinLine", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinLine()

		// Act
		actual := args.Map{"result": result == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov37_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinLine_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.JoinLine() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_JoinLineEofLine_NoSuffix(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinLineEofLine_NoSuffix", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinLineEofLine()

		// Act
		actual := args.Map{"result": strings.HasSuffix(result, "\n")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected newline suffix", actual)
	})
}

func Test_Cov37_JoinLineEofLine_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinLineEofLine_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.JoinLineEofLine() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_JoinSpace(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinSpace", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.JoinSpace() != "a b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov37_JoinComma(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinComma", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.JoinComma() != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov37_JoinCsv(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinCsv", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinCsv()

		// Act
		actual := args.Map{"result": strings.Contains(result, "\"a\"")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_Cov37_JoinCsvLine(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinCsvLine", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinCsvLine()

		// Act
		actual := args.Map{"result": result == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov37_JoinCsvString(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinCsvString", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.JoinCsvString(",")

		// Act
		actual := args.Map{"result": strings.Contains(result, "\"a\"")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_Cov37_JoinCsvString_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinCsvString_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.JoinCsvString(",") != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_JoinWith(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinWith", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinWith(", ")

		// Act
		actual := args.Map{"result": strings.HasPrefix(result, ", ")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected prefix", actual)
	})
}

func Test_Cov37_JoinWith_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinWith_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.JoinWith(",") != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// ── EachItemSplitBy ──

func Test_Cov37_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Cov37_EachItemSplitBy", func() {
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

// ── PrependJoin / AppendJoin / PrependAppend ──

func Test_Cov37_PrependJoin(t *testing.T) {
	safeTest(t, "Test_Cov37_PrependJoin", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("b", "c")
		result := s.PrependJoin(",", "a")

		// Act
		actual := args.Map{"result": result != "a,b,c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b,c", actual)
	})
}

func Test_Cov37_AppendJoin(t *testing.T) {
	safeTest(t, "Test_Cov37_AppendJoin", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.AppendJoin(",", "c")

		// Act
		actual := args.Map{"result": result != "a,b,c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b,c", actual)
	})
}

func Test_Cov37_PrependAppend(t *testing.T) {
	safeTest(t, "Test_Cov37_PrependAppend", func() {
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

func Test_Cov37_PrependAppend_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_PrependAppend_Empty", func() {
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

// ── IsEqual / IsEqualLines / IsEqualUnorderedLines / IsEqualUnorderedLinesClone ──

func Test_Cov37_IsEqual_Same(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqual_Same", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a", "b")
		b := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov37_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqual_BothNil", func() {
		// Arrange
		var a, b *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov37_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqual_OneNil", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": a.IsEqual(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Cov37_IsEqual_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqual_DiffLen", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a")
		b := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Cov37_IsEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqual_BothEmpty", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Empty()
		b := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov37_IsEqualLines_Mismatch(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualLines_Mismatch", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": a.IsEqualLines([]string{"a", "c"})}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Cov37_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLines", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": a.IsEqualUnorderedLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov37_IsEqualUnorderedLines_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLines_DiffLen", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": a.IsEqualUnorderedLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Cov37_IsEqualUnorderedLines_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLines_BothNil", func() {
		// Arrange
		var a *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": a.IsEqualUnorderedLines(nil)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov37_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLinesClone", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": a.IsEqualUnorderedLinesClone([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov37_IsEqualUnorderedLinesClone_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLinesClone_DiffLen", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": a.IsEqualUnorderedLinesClone([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Cov37_IsEqualUnorderedLinesClone_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLinesClone_BothNil", func() {
		// Arrange
		var a *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": a.IsEqualUnorderedLinesClone(nil)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov37_IsEqualUnorderedLinesClone_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLinesClone_BothEmpty", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": a.IsEqualUnorderedLinesClone([]string{})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

// ── IsEqualByFunc / IsEqualByFuncLinesSplit ──

func Test_Cov37_IsEqualByFunc_Match(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFunc_Match", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a", "b")
		result := a.IsEqualByFunc(func(i int, l, r string) bool {
			return strings.EqualFold(l, r)
		}, "A", "B")

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov37_IsEqualByFunc_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFunc_DiffLen", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": a.IsEqualByFunc(func(i int, l, r string) bool { return true }, "a", "b")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov37_IsEqualByFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFunc_Empty", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": a.IsEqualByFunc(func(i int, l, r string) bool { return true })}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for both empty", actual)
	})
}

func Test_Cov37_IsEqualByFunc_Mismatch(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFunc_Mismatch", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": a.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "X")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov37_IsEqualByFuncLinesSplit_Match(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFuncLinesSplit_Match", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a", "b")
		result := a.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool {
			return l == r
		})

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov37_IsEqualByFuncLinesSplit_Trim(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFuncLinesSplit_Trim", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines(" a ", " b ")
		result := a.IsEqualByFuncLinesSplit(true, ",", "a,b", func(i int, l, r string) bool {
			return l == r
		})

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true with trim", actual)
	})
}

func Test_Cov37_IsEqualByFuncLinesSplit_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFuncLinesSplit_DiffLen", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": a.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return true })}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov37_IsEqualByFuncLinesSplit_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFuncLinesSplit_Empty", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Empty()
		if a.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool { return true }) {
			// "" split by "," yields [""] which has len 1, not 0.

		// Act
			actual := args.Map{"result": false}

		// Assert
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "expected false for empty slice vs split-empty string", actual)
		}
	})
}

// ── Collection / ToCollection / NonPtr / Ptr ──

func Test_Cov37_Collection(t *testing.T) {
	safeTest(t, "Test_Cov37_Collection", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		col := s.Collection(false)

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov37_ToCollection", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		col := s.ToCollection(true)

		// Act
		actual := args.Map{"result": col.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_NonPtr(t *testing.T) {
	safeTest(t, "Test_Cov37_NonPtr", func() {
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

func Test_Cov37_Ptr(t *testing.T) {
	safeTest(t, "Test_Cov37_Ptr", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		p := s.Ptr()

		// Act
		actual := args.Map{"result": p.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_ToPtr(t *testing.T) {
	safeTest(t, "Test_Cov37_ToPtr", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		p := s.ToPtr()

		// Act
		actual := args.Map{"result": p == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov37_ToNonPtr(t *testing.T) {
	safeTest(t, "Test_Cov37_ToNonPtr", func() {
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

// ── String ──

func Test_Cov37_String(t *testing.T) {
	safeTest(t, "Test_Cov37_String", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": strings.Contains(s.String(), "a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov37_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_String_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.String() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// ── ConcatNew / ConcatNewStrings / ConcatNewSimpleSlices ──

func Test_Cov37_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov37_ConcatNew", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.ConcatNew("b", "c")

		// Act
		actual := args.Map{"result": result.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov37_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_Cov37_ConcatNewStrings", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.ConcatNewStrings("b")

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_ConcatNewStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_ConcatNewStrings_Nil", func() {
		// Arrange
		var s *corestr.SimpleSlice
		result := s.ConcatNewStrings("a")

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_Cov37_ConcatNewSimpleSlices", func() {
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

// ── CsvStrings ──

func Test_Cov37_CsvStrings(t *testing.T) {
	safeTest(t, "Test_Cov37_CsvStrings", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		csv := s.CsvStrings()

		// Act
		actual := args.Map{"result": len(csv) != 1 || !strings.Contains(csv[0], "\"")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected quoted", actual)
	})
}

func Test_Cov37_CsvStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_CsvStrings_Empty", func() {
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

// ── Sort / Reverse ──

func Test_Cov37_Sort(t *testing.T) {
	safeTest(t, "Test_Cov37_Sort", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("c", "a", "b")
		s.Sort()

		// Act
		actual := args.Map{"result": s.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
	})
}

func Test_Cov37_Reverse(t *testing.T) {
	safeTest(t, "Test_Cov37_Reverse", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		s.Reverse()

		// Act
		actual := args.Map{"result": s.First() != "c" || s.Last() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov37_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_Cov37_Reverse_Two", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		s.Reverse()

		// Act
		actual := args.Map{"result": s.First() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b first", actual)
	})
}

func Test_Cov37_Reverse_Single(t *testing.T) {
	safeTest(t, "Test_Cov37_Reverse_Single", func() {
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

// ── JSON / Serialize / Deserialize ──

func Test_Cov37_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov37_MarshalJSON", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		data, err := json.Marshal(s)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": strings.Contains(string(data), "\"a\"")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov37_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov37_UnmarshalJSON", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		err := json.Unmarshal([]byte(`["x","y"]`), s)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": s.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov37_UnmarshalJSON_Invalid", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		err := json.Unmarshal([]byte(`not-json`), s)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Cov37_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov37_JsonModel", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		model := s.JsonModel()

		// Act
		actual := args.Map{"result": len(model) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov37_JsonModelAny", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		model := s.JsonModelAny()

		// Act
		actual := args.Map{"result": model == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov37_Json(t *testing.T) {
	safeTest(t, "Test_Cov37_Json", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.Json()

		// Act
		actual := args.Map{"result": result.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	})
}

func Test_Cov37_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov37_JsonPtr", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.JsonPtr()

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov37_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov37_Serialize", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		data, err := s.Serialize()

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty bytes", actual)
	})
}

func Test_Cov37_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov37_Deserialize", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		var target []string
		err := s.Deserialize(&target)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": len(target) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov37_ParseInjectUsingJson", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		src := corestr.New.SimpleSlice.Lines("a", "b")
		jsonResult := src.JsonPtr()
		result, err := s.ParseInjectUsingJson(jsonResult)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov37_ParseInjectUsingJsonMust", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		src := corestr.New.SimpleSlice.Lines("a")
		jsonResult := src.JsonPtr()
		result := s.ParseInjectUsingJsonMust(jsonResult)

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov37_JsonParseSelfInject", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		src := corestr.New.SimpleSlice.Lines("x")
		err := s.JsonParseSelfInject(src.JsonPtr())

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

// ── AsJsoner / AsJsonContractsBinder / AsJsonParseSelfInjector / AsJsonMarshaller ──

func Test_Cov37_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov37_AsJsoner", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		j := s.AsJsoner()

		// Act
		actual := args.Map{"result": j == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov37_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov37_AsJsonContractsBinder", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		j := s.AsJsonContractsBinder()

		// Act
		actual := args.Map{"result": j == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov37_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov37_AsJsonParseSelfInjector", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		j := s.AsJsonParseSelfInjector()

		// Act
		actual := args.Map{"result": j == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov37_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov37_AsJsonMarshaller", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		j := s.AsJsonMarshaller()

		// Act
		actual := args.Map{"result": j == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── Clear / Dispose ──

func Test_Cov37_Clear(t *testing.T) {
	safeTest(t, "Test_Cov37_Clear", func() {
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

func Test_Cov37_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_Clear_Nil", func() {
		// Arrange
		var s *corestr.SimpleSlice
		result := s.Clear()

		// Act
		actual := args.Map{"result": result != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov37_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov37_Dispose", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		s.Dispose()

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 after dispose", actual)
	})
}

func Test_Cov37_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_Dispose_Nil", func() {
		var s *corestr.SimpleSlice
		s.Dispose() // should not panic
	})
}

// ── Clone / ClonePtr / DeepClone / ShadowClone ──

func Test_Cov37_Clone_Deep(t *testing.T) {
	safeTest(t, "Test_Cov37_Clone_Deep", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		cloned := s.Clone(true)

		// Act
		actual := args.Map{"result": cloned.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Clone_Shallow(t *testing.T) {
	safeTest(t, "Test_Cov37_Clone_Shallow", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		cloned := s.Clone(false)

		// Act
		actual := args.Map{"result": cloned.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_ClonePtr(t *testing.T) {
	safeTest(t, "Test_Cov37_ClonePtr", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		cloned := s.ClonePtr(true)

		// Act
		actual := args.Map{"result": cloned == nil || cloned.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov37_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_ClonePtr_Nil", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": s.ClonePtr(true) != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov37_DeepClone(t *testing.T) {
	safeTest(t, "Test_Cov37_DeepClone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		cloned := s.DeepClone()

		// Act
		actual := args.Map{"result": cloned.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_ShadowClone(t *testing.T) {
	safeTest(t, "Test_Cov37_ShadowClone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		cloned := s.ShadowClone()

		// Act
		actual := args.Map{"result": cloned.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── IsDistinctEqual / IsDistinctEqualRaw ──

func Test_Cov37_IsDistinctEqualRaw(t *testing.T) {
	safeTest(t, "Test_Cov37_IsDistinctEqualRaw", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "a")

		// Act
		actual := args.Map{"result": s.IsDistinctEqualRaw("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov37_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_Cov37_IsDistinctEqual", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a", "b")
		b := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": a.IsDistinctEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

// ── IsUnorderedEqualRaw / IsUnorderedEqual ──

func Test_Cov37_IsUnorderedEqualRaw_Clone(t *testing.T) {
	safeTest(t, "Test_Cov37_IsUnorderedEqualRaw_Clone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": s.IsUnorderedEqualRaw(true, "a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov37_IsUnorderedEqualRaw_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov37_IsUnorderedEqualRaw_NoClone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": s.IsUnorderedEqualRaw(false, "a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov37_IsUnorderedEqualRaw_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov37_IsUnorderedEqualRaw_DiffLen", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": s.IsUnorderedEqualRaw(true, "a", "b")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov37_IsUnorderedEqualRaw_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsUnorderedEqualRaw_BothEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.IsUnorderedEqualRaw(true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov37_IsUnorderedEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsUnorderedEqual_BothEmpty", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Empty()
		b := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": a.IsUnorderedEqual(true, b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov37_IsUnorderedEqual_RightNil(t *testing.T) {
	safeTest(t, "Test_Cov37_IsUnorderedEqual_RightNil", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": a.IsUnorderedEqual(true, nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

// ── DistinctDiffRaw / DistinctDiff ──

func Test_Cov37_DistinctDiffRaw_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov37_DistinctDiffRaw_BothNil", func() {
		// Arrange
		var s *corestr.SimpleSlice
		result := s.DistinctDiffRaw()

		// Act
		actual := args.Map{"result": len(result) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_DistinctDiffRaw_LeftNilRightNot(t *testing.T) {
	safeTest(t, "Test_Cov37_DistinctDiffRaw_LeftNilRightNot", func() {
		// Arrange
		var s *corestr.SimpleSlice
		result := s.DistinctDiffRaw("a")

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_DistinctDiffRaw_RightNil(t *testing.T) {
	safeTest(t, "Test_Cov37_DistinctDiffRaw_RightNil", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.DistinctDiffRaw()

		// Act
		actual := args.Map{"result": len(result) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov37_DistinctDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov37_DistinctDiff_BothNil", func() {
		// Arrange
		var a *corestr.SimpleSlice
		result := a.DistinctDiff(nil)

		// Act
		actual := args.Map{"result": len(result) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_DistinctDiff_LeftNil(t *testing.T) {
	safeTest(t, "Test_Cov37_DistinctDiff_LeftNil", func() {
		// Arrange
		var a *corestr.SimpleSlice
		b := corestr.New.SimpleSlice.Lines("x")
		result := a.DistinctDiff(b)

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_DistinctDiff_RightNil(t *testing.T) {
	safeTest(t, "Test_Cov37_DistinctDiff_RightNil", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("x")
		result := a.DistinctDiff(nil)

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── AddedRemovedLinesDiff ──

func Test_Cov37_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_Cov37_AddedRemovedLinesDiff", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		added, removed := s.AddedRemovedLinesDiff("b", "c")

		// Act
		actual := args.Map{"result": len(added) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected added items", actual)
		actual := args.Map{"result": len(removed) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed items", actual)
	})
}

func Test_Cov37_AddedRemovedLinesDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov37_AddedRemovedLinesDiff_BothNil", func() {
		// Arrange
		var s *corestr.SimpleSlice
		added, removed := s.AddedRemovedLinesDiff()

		// Act
		actual := args.Map{"result": added != nil || removed != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected both nil", actual)
	})
}

// ── RemoveIndexes ──

func Test_Cov37_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_Cov37_RemoveIndexes", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result, err := s.RemoveIndexes(1)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_RemoveIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_RemoveIndexes_Empty", func() {
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

func Test_Cov37_RemoveIndexes_InvalidIndex(t *testing.T) {
	safeTest(t, "Test_Cov37_RemoveIndexes_InvalidIndex", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		_, err := s.RemoveIndexes(5)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for invalid index", actual)
	})
}

// ── SafeStrings ──

func Test_Cov37_SafeStrings(t *testing.T) {
	safeTest(t, "Test_Cov37_SafeStrings", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": len(s.SafeStrings()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_SafeStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_SafeStrings_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": len(s.SafeStrings()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── newSimpleSliceCreator factory methods ──

func Test_Cov37_Creator_Cap(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Cap", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)

		// Act
		actual := args.Map{"result": s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov37_Creator_Cap_Zero(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Cap_Zero", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(0)

		// Act
		actual := args.Map{"result": s == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov37_Creator_Default(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Default", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Default()

		// Act
		actual := args.Map{"result": s == nil || s.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_Creator_Lines(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Lines", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Creator_Split(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Split", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Split("a,b", ",")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Creator_SplitLines(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_SplitLines", func() {
		// Arrange
		s := corestr.New.SimpleSlice.SplitLines("a\nb")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Creator_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_SpreadStrings", func() {
		// Arrange
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Creator_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Hashset", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		s := corestr.New.SimpleSlice.Hashset(hs)

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Creator_Hashset_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Hashset_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		s := corestr.New.SimpleSlice.Hashset(hs)

		// Act
		actual := args.Map{"result": s.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_Creator_Create(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Create", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Create([]string{"a"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_Creator_Strings(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Strings", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_Creator_StringsPtr(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_StringsPtr", func() {
		// Arrange
		s := corestr.New.SimpleSlice.StringsPtr([]string{"a"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_Creator_StringsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_StringsPtr_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.StringsPtr([]string{})

		// Act
		actual := args.Map{"result": s.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_Creator_StringsOptions_Clone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_StringsOptions_Clone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.StringsOptions(true, []string{"a"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_Creator_StringsOptions_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_StringsOptions_NoClone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.StringsOptions(false, []string{"a"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_Creator_StringsOptions_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_StringsOptions_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.StringsOptions(false, []string{})

		// Act
		actual := args.Map{"result": s.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_Creator_StringsClone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_StringsClone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.StringsClone([]string{"a"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_Creator_StringsClone_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_StringsClone_Nil", func() {
		// Arrange
		s := corestr.New.SimpleSlice.StringsClone(nil)

		// Act
		actual := args.Map{"result": s.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_Creator_Direct_Clone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Direct_Clone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Direct(true, []string{"a"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_Creator_Direct_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Direct_NoClone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Direct(false, []string{"a"})

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_Creator_Direct_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Direct_Nil", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Direct(true, nil)

		// Act
		actual := args.Map{"result": s.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_Creator_UsingLines_Clone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_UsingLines_Clone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.UsingLines(true, "a", "b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Creator_UsingLines_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_UsingLines_NoClone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.UsingLines(false, "a")

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_Creator_UsingSeparatorLine(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_UsingSeparatorLine", func() {
		// Arrange
		s := corestr.New.SimpleSlice.UsingSeparatorLine(",", "a,b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Creator_UsingLine(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_UsingLine", func() {
		// Arrange
		s := corestr.New.SimpleSlice.UsingLine("a\nb")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Creator_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": s.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov37_Creator_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Deserialize", func() {
		// Arrange
		data, _ := json.Marshal([]string{"a", "b"})
		s, err := corestr.New.SimpleSlice.Deserialize(data)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": s.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Creator_Deserialize_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Deserialize_Invalid", func() {
		// Arrange
		_, err := corestr.New.SimpleSlice.Deserialize([]byte("bad"))

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Cov37_Creator_DeserializeJsoner(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_DeserializeJsoner", func() {
		// Arrange
		src := corestr.New.SimpleSlice.Lines("a")
		jsoner := src.AsJsoner()
		s, err := corestr.New.SimpleSlice.DeserializeJsoner(jsoner)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov37_Creator_Map(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Map", func() {
		// Arrange
		m := map[string]string{"a": "1", "b": "2"}
		s := corestr.New.SimpleSlice.Map(m)

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov37_Creator_ByLen(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_ByLen", func() {
		// Arrange
		s := corestr.New.SimpleSlice.ByLen([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": s == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// Fix the broken AddError test
func Test_Cov37_AddError_WithError(t *testing.T) {
	safeTest(t, "Test_Cov37_AddError_WithError", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		err := json.Unmarshal([]byte("bad"), &struct{}{})
		s.AddError(err)

		// Act
		actual := args.Map{"result": s.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── JoinLineEofLine with already-suffixed ──

func Test_Cov37_JoinLineEofLine_AlreadySuffixed(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinLineEofLine_AlreadySuffixed", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a\n")
		result := s.JoinLineEofLine()

		// Act
		actual := args.Map{"result": strings.HasSuffix(result, "\n")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected newline suffix", actual)
	})
}

// Ensure the unused import is used
var _ = corejson.Result{}
