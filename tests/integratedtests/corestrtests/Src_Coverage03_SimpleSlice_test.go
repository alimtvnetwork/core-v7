package corestrtests

import (
	"errors"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC03_SimpleSlice_BasicState_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_BasicState_Verification", func() {
		for caseIndex, tc := range srcC03BasicStateTestCases {
			// Arrange
			s := corestr.New.SimpleSlice.Empty()

			// Act
			actual := args.Map{
				"isEmpty":   s.IsEmpty(),
				"hasAny":    s.HasAnyItem(),
				"length":    s.Length(),
				"count":     s.Count(),
				"lastIndex": s.LastIndex(),
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_SrcC03_SimpleSlice_NilReceiver_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_NilReceiver_Verification", func() {
		for caseIndex, tc := range srcC03NilReceiverTestCases {
			// Arrange
			var s *corestr.SimpleSlice

			// Act
			actual := args.Map{
				"length":  s.Length(),
				"isEmpty": s.IsEmpty(),
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_SrcC03_SimpleSlice_AddMethods_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_AddMethods_Verification", func() {
		for caseIndex, tc := range srcC03AddMethodsTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			method, _ := input.GetAsString("method")
			s := corestr.New.SimpleSlice.Empty()

			// Act
			switch method {
			case "Add":
				s.Add("a").Add("b")
			case "AddSplit":
				s.AddSplit("a,b,c", ",")
			case "AddIf":
				s.AddIf(false, "skip")
				s.AddIf(true, "add")
			case "Adds":
				s.Adds()
				s.Adds("a", "b")
			case "Append":
				s.Append()
				s.Append("a")
			case "AppendFmt":
				s.AppendFmt("")
				s.AppendFmt("hello %s", "world")
			case "AppendFmtIf":
				s.AppendFmtIf(false, "skip")
				s.AppendFmtIf(true, "val=%d", 42)
			}
			actual := args.Map{
				"length": s.Length(),
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_SrcC03_SimpleSlice_TitleValue_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_TitleValue_Verification", func() {
		for caseIndex, tc := range srcC03TitleValueTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			method, _ := input.GetAsString("method")
			s := corestr.New.SimpleSlice.Empty()

			// Act
			switch method {
			case "AddAsTitleValue":
				s.AddAsTitleValue("key", "val")
			case "AddAsCurlyTitleWrap":
				s.AddAsCurlyTitleWrap("key", "val")
			case "AddAsCurlyTitleWrapIf":
				s.AddAsCurlyTitleWrapIf(false, "k", "v")
				s.AddAsCurlyTitleWrapIf(true, "k", "v")
			case "AddAsTitleValueIf":
				s.AddAsTitleValueIf(false, "k", "v")
				s.AddAsTitleValueIf(true, "k", "v")
			}
			actual := args.Map{
				"length": s.Length(),
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_SrcC03_SimpleSlice_InsertAt_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_InsertAt_Verification", func() {
		// Arrange
		tc := srcC03InsertAtTestCase
		s := corestr.New.SimpleSlice.Lines("a", "c")

		// Act
		s.InsertAt(1, "b")
		s.InsertAt(-1, "x")  // out of range
		s.InsertAt(100, "x") // out of range
		actual := args.Map{
			"length":  s.Length(),
			"atIndex": (*s)[1],
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC03_SimpleSlice_ConditionalAdds_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_ConditionalAdds_Verification", func() {
		for caseIndex, tc := range srcC03ConditionalAddsTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			method, _ := input.GetAsString("method")
			s := corestr.New.SimpleSlice.Empty()

			// Act
			switch method {
			case "AddsIf":
				s.AddsIf(false, "skip")
				s.AddsIf(true, "a", "b")
			case "AddError":
				s.AddError(nil)
				s.AddError(errors.New("e"))
			case "AddStruct":
				s.AddStruct(false, nil)
				s.AddStruct(false, "hello")
			case "AddPointer":
				s.AddPointer(false, nil)
				val := "hello"
				s.AddPointer(false, &val)
			}
			actual := args.Map{
				"length": s.Length(),
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_SrcC03_SimpleSlice_AsError_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_AsError_Verification", func() {
		for caseIndex, tc := range srcC03AsErrorTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			hasItems := input.GetAsBoolDefault("hasItems", false)
			s := corestr.New.SimpleSlice.Empty()
			if hasItems {
				s.Add("e")
			}

			// Act
			actual := args.Map{}
			if !hasItems {
				actual["defaultErrorNil"] = s.AsDefaultError() == nil
				actual["asErrorNil"] = s.AsError(",") == nil
			} else {
				actual["defaultErrorNil"] = s.AsDefaultError() == nil
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_SrcC03_SimpleSlice_FirstLast_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_FirstLast_Verification", func() {
		// Arrange
		tc := srcC03FirstLastTestCase
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		actual := args.Map{
			"first":                  s.First(),
			"last":                   s.Last(),
			"firstDynamic":           s.FirstDynamic(),
			"lastDynamic":            s.LastDynamic(),
			"firstOrDefault":         s.FirstOrDefault(),
			"lastOrDefault":          s.LastOrDefault(),
			"firstOrDefaultDynamic":  s.FirstOrDefaultDynamic(),
			"lastOrDefaultDynamic":   s.LastOrDefaultDynamic(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC03_SimpleSlice_FirstLastEmpty_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_FirstLastEmpty_Verification", func() {
		// Arrange
		tc := srcC03FirstLastEmptyTestCase
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{
			"firstOrDefault": s.FirstOrDefault(),
			"lastOrDefault":  s.LastOrDefault(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC03_SimpleSlice_SkipTake_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_SkipTake_Verification", func() {
		// Arrange
		tc := srcC03SkipTakeTestCase
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		_ = s.SkipDynamic(1)
		_ = s.TakeDynamic(1)
		_ = s.LimitDynamic(1)
		_ = s.Limit(1)
		actual := args.Map{
			"skip1":   len(s.Skip(1)),
			"skip100": len(s.Skip(100)),
			"take2":   len(s.Take(2)),
			"take100": len(s.Take(100)),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC03_SimpleSlice_CountFunc_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_CountFunc_Verification", func() {
		for caseIndex, tc := range srcC03CountFuncTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			itemsStr, _ := input.GetAsString("items")
			var s *corestr.SimpleSlice
			if itemsStr == "" {
				s = corestr.New.SimpleSlice.Empty()
			} else {
				s = corestr.New.SimpleSlice.Lines(strings.Split(itemsStr, ",")...)
			}

			// Act
			count := s.CountFunc(func(i int, item string) bool { return len(item) > 1 })
			actual := args.Map{
				"count": count,
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_SrcC03_SimpleSlice_Contains_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_Contains_Verification", func() {
		for caseIndex, tc := range srcC03ContainsTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			search, _ := input.GetAsString("search")
			itemsStr, _ := input.GetAsString("items")
			var s *corestr.SimpleSlice
			if itemsStr == "" {
				s = corestr.New.SimpleSlice.Empty()
			} else {
				s = corestr.New.SimpleSlice.Lines(strings.Split(itemsStr, ",")...)
			}

			// Act
			actual := args.Map{
				"found": s.IsContains(search),
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_SrcC03_SimpleSlice_ContainsFunc_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_ContainsFunc_Verification", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("hello", "world")
		empty := corestr.New.SimpleSlice.Empty()

		// Act
		found := s.IsContainsFunc("hello", func(item, searching string) bool { return item == searching })
		notFound := empty.IsContainsFunc("a", func(item, searching string) bool { return true })

		// Assert
		actual := args.Map{"result": found}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		actual := args.Map{"result": notFound}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not found", actual)
	})
}

func Test_SrcC03_SimpleSlice_IndexOf_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_IndexOf_Verification", func() {
		for caseIndex, tc := range srcC03IndexOfTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			search, _ := input.GetAsString("search")
			itemsStr, _ := input.GetAsString("items")
			var s *corestr.SimpleSlice
			if itemsStr == "" {
				s = corestr.New.SimpleSlice.Empty()
			} else {
				s = corestr.New.SimpleSlice.Lines(strings.Split(itemsStr, ",")...)
			}

			// Act
			actual := args.Map{
				"index": s.IndexOf(search),
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_SrcC03_SimpleSlice_IndexOfFunc_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_IndexOfFunc_Verification", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		empty := corestr.New.SimpleSlice.Empty()

		// Act
		idx := s.IndexOfFunc("b", func(item, searching string) bool { return item == searching })
		emptyIdx := empty.IndexOfFunc("a", func(item, searching string) bool { return true })

		// Assert
		actual := args.Map{"result": idx != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": emptyIdx != -1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_SrcC03_SimpleSlice_HasIndex_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_HasIndex_Verification", func() {
		// Arrange
		tc := srcC03HasIndexTestCase
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{
			"has0":    s.HasIndex(0),
			"has1":    s.HasIndex(1),
			"has2":    s.HasIndex(2),
			"hasNeg1": s.HasIndex(-1),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC03_SimpleSlice_StringsList_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_StringsList_Verification", func() {
		// Arrange
		tc := srcC03StringsListTestCase
		s := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{
			"stringsLen": len(s.Strings()),
			"listLen":    len(s.List()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC03_SimpleSlice_WrapQuotes_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC03_SimpleSlice_WrapQuotes_Verification", func() {
		// Arrange
		tc := srcC03WrapQuotesTestCase

		// Act
		noPanic := !callPanicsSrcC03(func() {
			s := corestr.New.SimpleSlice.Lines("a")
			_ = s.WrapDoubleQuote()
			s2 := corestr.New.SimpleSlice.Lines("a")
			_ = s2.WrapSingleQuote()
			s3 := corestr.New.SimpleSlice.Lines("a")
			_ = s3.WrapTildaQuote()
			s4 := corestr.New.SimpleSlice.Lines("a")
			_ = s4.WrapDoubleQuoteIfMissing()
			s5 := corestr.New.SimpleSlice.Lines("a")
			_ = s5.WrapSingleQuoteIfMissing()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC03(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
