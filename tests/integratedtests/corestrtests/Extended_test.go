package corestrtests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ==========================================
// Collection - Add / Adds / Length
// ==========================================

func Test_ExtCollection_Add_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_Add_Verification", func() {
		for caseIndex, testCase := range extCollectionAddTestCases {
			// Arrange
			input := testCase.ArrangeInput.(args.Map)
			items := input["items"].([]string)

			// Act
			col := corestr.New.Collection.Cap(10)
			col.Adds(items...)

			// Assert
			testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", col.Length()))
		}
	})
}

// ==========================================
// Collection - Join
// ==========================================

func Test_ExtCollection_Join_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_Join_Verification", func() {
		for caseIndex, testCase := range extCollectionJoinTestCases {
			// Arrange
			input := testCase.ArrangeInput.(args.Map)
			items := input["items"].([]string)
			joiner, _ := input.GetAsString("joiner")

			// Act
			col := corestr.New.Collection.Cap(10)
			col.Adds(items...)
			result := col.Join(joiner)

			// Assert
			testCase.ShouldBeEqual(t, caseIndex, result)
		}
	})
}

// ==========================================
// Collection - AddNonEmpty / AddNonEmptyWhitespace
// ==========================================

func Test_ExtCollection_AddNonEmpty_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_AddNonEmpty_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)

		// Act
		col.AddNonEmpty("")
		col.AddNonEmpty("hello")
		col.AddNonEmpty("")
		col.AddNonEmpty("world")

		// Assert
		if col.Length() != 2 {
			t.Errorf("AddNonEmpty expected 2 items, got %d", col.Length())
		}
	})
}

func Test_ExtCollection_AddNonEmptyWhitespace_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_AddNonEmptyWhitespace_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)

		// Act
		col.AddNonEmptyWhitespace("  ")
		col.AddNonEmptyWhitespace("hello")
		col.AddNonEmptyWhitespace("\t\n")

		// Assert
		if col.Length() != 1 {
			t.Errorf("AddNonEmptyWhitespace expected 1, got %d", col.Length())
		}
	})
}

// ==========================================
// Collection - AddIf / AddIfMany
// ==========================================

func Test_ExtCollection_AddIf_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_AddIf_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)

		// Act
		col.AddIf(true, "yes")
		col.AddIf(false, "no")

		// Assert
		if col.Length() != 1 {
			t.Errorf("AddIf expected 1, got %d", col.Length())
		}
	})
}

func Test_ExtCollection_AddIfMany_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_AddIfMany_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)

		// Act
		col.AddIfMany(true, "a", "b")
		col.AddIfMany(false, "c", "d")

		// Assert
		if col.Length() != 2 {
			t.Errorf("AddIfMany expected 2, got %d", col.Length())
		}
	})
}

// ==========================================
// Collection - AddFunc / AddError
// ==========================================

func Test_ExtCollection_AddFunc_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_AddFunc_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)

		// Act
		col.AddFunc(func() string { return "from func" })

		// Assert
		if col.Length() != 1 {
			t.Errorf("AddFunc expected 1, got %d", col.Length())
		}
	})
}

func Test_ExtCollection_AddError_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_AddError_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)

		// Act
		col.AddError(nil) // should skip
		col.AddError(fmt.Errorf("test error"))

		// Assert
		if col.Length() != 1 {
			t.Errorf("AddError expected 1, got %d", col.Length())
		}
	})
}

// ==========================================
// Collection - IsEquals / IsEqualsWithSensitive
// ==========================================

func Test_ExtCollection_IsEquals_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_IsEquals_Verification", func() {
		// Arrange
		col1 := corestr.New.Collection.Strings([]string{"a", "b"})
		col2 := corestr.New.Collection.Strings([]string{"a", "b"})
		col3 := corestr.New.Collection.Strings([]string{"A", "B"})

		// Act & Assert
		if !col1.IsEquals(col2) {
			t.Error("Same content should be equal")
		}
		if col1.IsEquals(col3) {
			t.Error("Different case should not be equal (case-sensitive)")
		}
		if !col1.IsEqualsWithSensitive(false, col3) {
			t.Error("Case-insensitive should match")
		}
	})
}

// ==========================================
// Collection - nil receiver
// ==========================================

func Test_ExtCollection_NilReceiver_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_NilReceiver_Verification", func() {
		// Arrange
		var col *corestr.Collection

		// Act & Assert
		if !col.IsEmpty() {
			t.Error("nil.IsEmpty() should be true")
		}
		if col.Length() != 0 {
			t.Error("nil.Length() should be 0")
		}
		if col.HasItems() {
			t.Error("nil.HasItems() should be false")
		}
	})
}

// ==========================================
// Collection - RemoveAt
// ==========================================

func Test_ExtCollection_RemoveAt_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_RemoveAt_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		ok := col.RemoveAt(1)

		// Assert
		if !ok || col.Length() != 2 {
			t.Errorf("RemoveAt(1) expected success, got ok=%v len=%d", ok, col.Length())
		}

		// Act - out of range
		ok2 := col.RemoveAt(10)

		// Assert
		if ok2 {
			t.Error("RemoveAt(10) should return false")
		}

		// Act - negative
		ok3 := col.RemoveAt(-1)

		// Assert
		if ok3 {
			t.Error("RemoveAt(-1) should return false")
		}
	})
}

// ==========================================
// Collection - ConcatNew
// ==========================================

func Test_ExtCollection_ConcatNew_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_ConcatNew_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		result := col.ConcatNew(0, "b", "c")

		// Assert
		if result.Length() != 3 {
			t.Errorf("ConcatNew expected 3, got %d", result.Length())
		}

		// Act - empty
		result2 := col.ConcatNew(0)

		// Assert
		if result2.Length() != 1 {
			t.Errorf("ConcatNew empty expected 1, got %d", result2.Length())
		}
	})
}

// ==========================================
// Collection - AsDefaultError / AsError
// ==========================================

func Test_ExtCollection_AsError_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_AsError_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"err1", "err2"})

		// Act
		err := col.AsError("; ")

		// Assert
		if err == nil || err.Error() != "err1; err2" {
			t.Errorf("AsError expected 'err1; err2', got '%v'", err)
		}

		// Act - empty
		emptyCol := corestr.New.Collection.Cap(0)
		err2 := emptyCol.AsDefaultError()

		// Assert
		if err2 != nil {
			t.Error("AsDefaultError on empty should return nil")
		}
	})
}

// ==========================================
// Collection - EachItemSplitBy
// ==========================================

func Test_ExtCollection_EachItemSplitBy_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_EachItemSplitBy_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a,b", "c,d"})

		// Act
		result := col.EachItemSplitBy(",")

		// Assert
		if len(result) != 4 {
			t.Errorf("EachItemSplitBy expected 4, got %d", len(result))
		}
	})
}

// ==========================================
// Collection - HasIndex / LastIndex / HasAnyItem
// ==========================================

func Test_ExtCollection_HasIndex_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_HasIndex_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act & Assert
		if !col.HasIndex(0) {
			t.Error("HasIndex(0) should be true")
		}
		if !col.HasIndex(2) {
			t.Error("HasIndex(2) should be true")
		}
		if col.HasIndex(3) {
			t.Error("HasIndex(3) should be false")
		}
		if col.LastIndex() != 2 {
			t.Errorf("LastIndex expected 2, got %d", col.LastIndex())
		}
		if !col.HasAnyItem() {
			t.Error("HasAnyItem should be true")
		}
	})
}

// ==========================================
// Collection - AddCollection / AddCollections
// ==========================================

func Test_ExtCollection_AddCollection_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_AddCollection_Verification", func() {
		// Arrange
		col1 := corestr.New.Collection.Strings([]string{"a"})
		col2 := corestr.New.Collection.Strings([]string{"b", "c"})

		// Act
		col1.AddCollection(col2)

		// Assert
		if col1.Length() != 3 {
			t.Errorf("AddCollection expected 3, got %d", col1.Length())
		}
	})
}

func Test_ExtCollection_AddCollections_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_AddCollections_Verification", func() {
		// Arrange
		col1 := corestr.New.Collection.Strings([]string{"a"})
		col2 := corestr.New.Collection.Strings([]string{"b"})
		col3 := corestr.New.Collection.Strings([]string{"c"})

		// Act
		col1.AddCollections(col2, col3)

		// Assert
		if col1.Length() != 3 {
			t.Errorf("AddCollections expected 3, got %d", col1.Length())
		}
	})
}

// ==========================================
// SimpleSlice
// ==========================================

func Test_ExtSimpleSlice_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_Verification", func() {
		for caseIndex, testCase := range extSimpleSliceTestCases {
			// Arrange
			input := testCase.ArrangeInput.(args.Map)
			items := input["items"].([]string)

			// Act
			ss := corestr.New.SimpleSlice.Cap(10)
			ss.Adds(items...)

			// Assert
			testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", ss.Length()))
		}
	})
}

func Test_ExtSimpleSlice_AddIf_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_AddIf_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(10)

		// Act
		ss.AddIf(true, "yes")
		ss.AddIf(false, "no")

		// Assert
		if ss.Length() != 1 {
			t.Errorf("AddIf expected 1, got %d", ss.Length())
		}
	})
}

func Test_ExtSimpleSlice_FirstLast_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_FirstLast_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act & Assert
		if ss.First() != "a" {
			t.Errorf("First expected 'a', got '%s'", ss.First())
		}
		if ss.Last() != "c" {
			t.Errorf("Last expected 'c', got '%s'", ss.Last())
		}
		if ss.FirstOrDefault() != "a" {
			t.Errorf("FirstOrDefault expected 'a', got '%s'", ss.FirstOrDefault())
		}
		if ss.LastOrDefault() != "c" {
			t.Errorf("LastOrDefault expected 'c', got '%s'", ss.LastOrDefault())
		}
	})
}

func Test_ExtSimpleSlice_EmptyDefaults_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_EmptyDefaults_Verification", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act & Assert
		if ss.FirstOrDefault() != "" {
			t.Error("nil.FirstOrDefault() should return empty")
		}
		if ss.LastOrDefault() != "" {
			t.Error("nil.LastOrDefault() should return empty")
		}
		if !ss.IsEmpty() {
			t.Error("nil.IsEmpty() should be true")
		}
		if ss.Length() != 0 {
			t.Error("nil.Length() should be 0")
		}
	})
}

func Test_ExtSimpleSlice_SkipTake_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_SkipTake_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c", "d"})

		// Act
		skipped := ss.Skip(2)
		taken := ss.Take(2)

		// Assert
		if len(skipped) != 2 || skipped[0] != "c" {
			t.Errorf("Skip(2) expected [c,d], got %v", skipped)
		}
		if len(taken) != 2 || taken[0] != "a" {
			t.Errorf("Take(2) expected [a,b], got %v", taken)
		}
	})
}

func Test_ExtSimpleSlice_IsContains_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_IsContains_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"hello", "world"})

		// Act & Assert
		if !ss.IsContains("hello") {
			t.Error("IsContains should find 'hello'")
		}
		if ss.IsContains("missing") {
			t.Error("IsContains should not find 'missing'")
		}
	})
}

func Test_ExtSimpleSlice_IndexOf_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_IndexOf_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act & Assert
		if ss.IndexOf("b") != 1 {
			t.Errorf("IndexOf('b') expected 1, got %d", ss.IndexOf("b"))
		}
		if ss.IndexOf("z") != -1 {
			t.Errorf("IndexOf('z') expected -1, got %d", ss.IndexOf("z"))
		}
	})
}

func Test_ExtSimpleSlice_InsertAt_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_InsertAt_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "c"})

		// Act
		ss.InsertAt(1, "b")

		// Assert
		if ss.Length() != 3 || (*ss)[1] != "b" {
			t.Errorf("InsertAt expected [a,b,c], got %v", *ss)
		}
	})
}

func Test_ExtSimpleSlice_AddError_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_AddError_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddError(nil)
		ss.AddError(fmt.Errorf("oops"))

		// Assert
		if ss.Length() != 1 {
			t.Errorf("AddError expected 1, got %d", ss.Length())
		}
	})
}

func Test_ExtSimpleSlice_AsError_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_AsError_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"e1", "e2"})

		// Act
		err := ss.AsError(", ")

		// Assert
		if err == nil || err.Error() != "e1, e2" {
			t.Errorf("AsError expected 'e1, e2', got '%v'", err)
		}

		// Act - nil empty
		var nilSS *corestr.SimpleSlice
		err2 := nilSS.AsError(", ")

		// Assert
		if err2 != nil {
			t.Error("nil.AsError should return nil")
		}
	})
}

func Test_ExtSimpleSlice_AppendFmt_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_AppendFmt_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AppendFmt("Hello %s", "World")
		ss.AppendFmt("") // empty format+values should skip
		ss.AppendFmtIf(true, "Yes %d", 1)
		ss.AppendFmtIf(false, "No %d", 2)

		// Assert
		if ss.Length() != 2 {
			t.Errorf("AppendFmt expected 2 items, got %d", ss.Length())
		}
	})
}

func Test_ExtSimpleSlice_CountFunc_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_CountFunc_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "bb", "ccc"})

		// Act
		count := ss.CountFunc(func(index int, item string) bool {
			return len(item) > 1
		})

		// Assert
		if count != 2 {
			t.Errorf("CountFunc expected 2, got %d", count)
		}
	})
}

// ==========================================
// LeftRight
// ==========================================

func Test_ExtLeftRight_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_Verification", func() {
		for caseIndex, testCase := range extLeftRightTestCases {
			// Arrange
			input := testCase.ArrangeInput.(args.Map)

			var lr *corestr.LeftRight

			if sliceRaw, hasSlice := input["slice"]; hasSlice {
				// Act - from slice
				lr = corestr.LeftRightUsingSlice(sliceRaw.([]string))
			} else {
				left, _ := input.GetAsString("left")
				right, _ := input.GetAsString("right")
				// Act - from constructor
				lr = corestr.NewLeftRight(left, right)
			}

			actual := args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}

			// Assert
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_ExtLeftRight_Methods_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_Methods_Verification", func() {
		// Arrange
		lr := corestr.NewLeftRight(" hello ", " world ")

		// Act & Assert
		if lr.IsLeftEmpty() {
			t.Error("IsLeftEmpty should be false")
		}
		if lr.IsRightEmpty() {
			t.Error("IsRightEmpty should be false")
		}
		if !lr.HasSafeNonEmpty() {
			t.Error("HasSafeNonEmpty should be true")
		}
		if lr.LeftTrim() != "hello" {
			t.Errorf("LeftTrim expected 'hello', got '%s'", lr.LeftTrim())
		}
		if lr.RightTrim() != "world" {
			t.Errorf("RightTrim expected 'world', got '%s'", lr.RightTrim())
		}
		if string(lr.LeftBytes()) != " hello " {
			t.Errorf("LeftBytes mismatch")
		}
		if string(lr.RightBytes()) != " world " {
			t.Errorf("RightBytes mismatch")
		}
	})
}

func Test_ExtLeftRight_IsEqual_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_IsEqual_Verification", func() {
		// Arrange
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("a", "c")

		// Act & Assert
		if !lr1.IsEqual(lr2) {
			t.Error("Same content should be equal")
		}
		if lr1.IsEqual(lr3) {
			t.Error("Different content should not be equal")
		}

		// nil cases
		var nilLR *corestr.LeftRight
		if !nilLR.IsEqual(nil) {
			t.Error("nil.IsEqual(nil) should be true")
		}
		if lr1.IsEqual(nil) {
			t.Error("non-nil.IsEqual(nil) should be false")
		}
	})
}

func Test_ExtLeftRight_Clone_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_Clone_Verification", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		cloned := lr.Clone()

		// Assert
		if cloned.Left != "a" || cloned.Right != "b" {
			t.Error("Clone should copy values")
		}
	})
}

func Test_ExtLeftRight_RegexMatch_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_RegexMatch_Verification", func() {
		// Arrange
		lr := corestr.NewLeftRight("abc123", "def456")
		re := regexp.MustCompile(`\d+`)

		// Act & Assert
		if !lr.IsLeftRegexMatch(re) {
			t.Error("IsLeftRegexMatch should be true")
		}
		if !lr.IsRightRegexMatch(re) {
			t.Error("IsRightRegexMatch should be true")
		}
		if lr.IsLeftRegexMatch(nil) {
			t.Error("nil regex should return false")
		}
	})
}

func Test_ExtLeftRight_Is_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_Is_Verification", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act & Assert
		if !lr.IsLeft("a") {
			t.Error("IsLeft should be true")
		}
		if !lr.IsRight("b") {
			t.Error("IsRight should be true")
		}
		if !lr.Is("a", "b") {
			t.Error("Is should be true")
		}
	})
}

func Test_ExtLeftRight_InvalidCreation_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_InvalidCreation_Verification", func() {
		// Arrange
		// Act
		lr1 := corestr.InvalidLeftRightNoMessage()
		lr2 := corestr.InvalidLeftRight("test error")

		// Assert
		if lr1.IsValid {
			t.Error("InvalidLeftRightNoMessage should be invalid")
		}
		if lr2.IsValid || lr2.Message == "" {
			t.Error("InvalidLeftRight should be invalid with message")
		}
	})
}

func Test_ExtLeftRight_Dispose_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_Dispose_Verification", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		lr.Dispose()

		// Assert
		if lr.Left != "" || lr.Right != "" {
			t.Error("Dispose should clear values")
		}
	})
}

// ==========================================
// LeftMiddleRight
// ==========================================

func Test_ExtLeftMiddleRight_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftMiddleRight_Verification", func() {
		for caseIndex, testCase := range extLeftMiddleRightTestCases {
			// Arrange
			input := testCase.ArrangeInput.(args.Map)
			left, _ := input.GetAsString("left")
			middle, _ := input.GetAsString("middle")
			right, _ := input.GetAsString("right")

			// Act
			lmr := corestr.NewLeftMiddleRight(left, middle, right)
			actual := args.Map{
				"left":          lmr.Left,
				"middle":        lmr.Middle,
				"right":         lmr.Right,
				"isLeftEmpty":   fmt.Sprintf("%v", lmr.IsLeftEmpty()),
				"isMiddleEmpty": fmt.Sprintf("%v", lmr.IsMiddleEmpty()),
				"isRightEmpty":  fmt.Sprintf("%v", lmr.IsRightEmpty()),
			}

			// Assert
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_ExtLeftMiddleRight_Methods_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftMiddleRight_Methods_Verification", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight(" L ", " M ", " R ")

		// Act & Assert
		if lmr.LeftTrim() != "L" {
			t.Errorf("LeftTrim expected 'L', got '%s'", lmr.LeftTrim())
		}
		if lmr.MiddleTrim() != "M" {
			t.Errorf("MiddleTrim expected 'M', got '%s'", lmr.MiddleTrim())
		}
		if lmr.RightTrim() != "R" {
			t.Errorf("RightTrim expected 'R', got '%s'", lmr.RightTrim())
		}
		if string(lmr.LeftBytes()) != " L " {
			t.Error("LeftBytes mismatch")
		}
		if string(lmr.MiddleBytes()) != " M " {
			t.Error("MiddleBytes mismatch")
		}
		if string(lmr.RightBytes()) != " R " {
			t.Error("RightBytes mismatch")
		}
		if !lmr.HasSafeNonEmpty() {
			t.Error("HasSafeNonEmpty should be true")
		}
		if !lmr.IsAll(" L ", " M ", " R ") {
			t.Error("IsAll should be true")
		}
	})
}

func Test_ExtLeftMiddleRight_ToLeftRight_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftMiddleRight_ToLeftRight_Verification", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("L", "M", "R")

		// Act
		lr := lmr.ToLeftRight()

		// Assert
		if lr.Left != "L" || lr.Right != "R" {
			t.Errorf("ToLeftRight expected L/R, got %s/%s", lr.Left, lr.Right)
		}
	})
}

func Test_ExtLeftMiddleRight_Clone_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftMiddleRight_Clone_Verification", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("L", "M", "R")

		// Act
		cloned := lmr.Clone()

		// Assert
		if cloned.Left != "L" || cloned.Middle != "M" || cloned.Right != "R" {
			t.Error("Clone should copy all values")
		}
	})
}

func Test_ExtLeftMiddleRight_Invalid_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftMiddleRight_Invalid_Verification", func() {
		// Arrange
		// Act
		lmr1 := corestr.InvalidLeftMiddleRightNoMessage()
		lmr2 := corestr.InvalidLeftMiddleRight("test")

		// Assert
		if lmr1.IsValid {
			t.Error("InvalidLeftMiddleRightNoMessage should be invalid")
		}
		if lmr2.IsValid || lmr2.Message == "" {
			t.Error("InvalidLeftMiddleRight should be invalid with message")
		}
	})
}

func Test_ExtLeftMiddleRight_Dispose_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftMiddleRight_Dispose_Verification", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("L", "M", "R")

		// Act
		lmr.Dispose()

		// Assert
		if lmr.Left != "" || lmr.Middle != "" || lmr.Right != "" {
			t.Error("Dispose should clear all values")
		}
	})
}

// ==========================================
// Hashset
// ==========================================

func Test_ExtHashset_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_Verification", func() {
		for caseIndex, testCase := range extHashsetTestCases {
			// Arrange
			input := testCase.ArrangeInput.(args.Map)
			items := input["items"].([]string)

			// Act
			hs := corestr.New.Hashset.Strings(items)
			actual := args.Map{
				"length":  fmt.Sprintf("%d", hs.Length()),
				"hasA":    fmt.Sprintf("%v", hs.Has("a")),
				"hasB":    fmt.Sprintf("%v", hs.Has("b")),
				"hasMiss": fmt.Sprintf("%v", hs.Has("missing")),
			}

			// Assert
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_ExtHashset_AddRemove_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_AddRemove_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(10)

		// Act
		hs.Add("a")
		hs.Add("b")
		hs.Add("c")

		// Assert
		if hs.Length() != 3 {
			t.Errorf("After adds expected 3, got %d", hs.Length())
		}

		// Act
		hs.Remove("b")

		// Assert
		if hs.Length() != 2 {
			t.Errorf("After remove expected 2, got %d", hs.Length())
		}
		if hs.Has("b") {
			t.Error("Should not have 'b' after remove")
		}
	})
}

func Test_ExtHashset_AddNonEmpty_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_AddNonEmpty_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(10)

		// Act
		hs.AddNonEmpty("")
		hs.AddNonEmpty("hello")

		// Assert
		if hs.Length() != 1 {
			t.Errorf("AddNonEmpty expected 1, got %d", hs.Length())
		}
	})
}

func Test_ExtHashset_AddNonEmptyWhitespace_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_AddNonEmptyWhitespace_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(10)

		// Act
		hs.AddNonEmptyWhitespace("  ")
		hs.AddNonEmptyWhitespace("hello")

		// Assert
		if hs.Length() != 1 {
			t.Errorf("AddNonEmptyWhitespace expected 1, got %d", hs.Length())
		}
	})
}

func Test_ExtHashset_AddIf_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_AddIf_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(10)

		// Act
		hs.AddIf(true, "yes")
		hs.AddIf(false, "no")

		// Assert
		if hs.Length() != 1 {
			t.Errorf("AddIf expected 1, got %d", hs.Length())
		}
	})
}

func Test_ExtHashset_List_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_List_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"b", "a", "c"})

		// Act
		list := hs.List()

		// Assert
		if len(list) != 3 {
			t.Errorf("List expected 3 items, got %d", len(list))
		}
	})
}

func Test_ExtHashset_SortedList_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_SortedList_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})

		// Act
		sorted := hs.SortedList()

		// Assert
		if len(sorted) != 3 || sorted[0] != "a" || sorted[1] != "b" || sorted[2] != "c" {
			t.Errorf("SortedList expected [a,b,c], got %v", sorted)
		}
	})
}

func Test_ExtHashset_HasAll_HasAny_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_HasAll_HasAny_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})

		// Act & Assert
		if !hs.HasAll("a", "c") {
			t.Error("HasAll should be true")
		}
		if hs.HasAll("a", "z") {
			t.Error("HasAll should be false with missing")
		}
		if !hs.HasAny("z", "a") {
			t.Error("HasAny should be true")
		}
		if hs.HasAny("x", "y") {
			t.Error("HasAny should be false")
		}
	})
}

func Test_ExtHashset_NilReceiver_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_NilReceiver_Verification", func() {
		// Arrange
		var hs *corestr.Hashset

		// Act & Assert
		if !hs.IsEmpty() {
			t.Error("nil.IsEmpty() should be true")
		}
		if hs.HasItems() {
			t.Error("nil.HasItems() should be false")
		}
	})
}

func Test_ExtHashset_AddCollection_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_AddCollection_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(10)
		col := corestr.New.Collection.Strings([]string{"x", "y"})

		// Act
		hs.AddCollection(col)

		// Assert
		if hs.Length() != 2 {
			t.Errorf("AddCollection expected 2, got %d", hs.Length())
		}
	})
}

func Test_ExtHashset_AddHashsetItems_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_AddHashsetItems_Verification", func() {
		// Arrange
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hs2 := corestr.New.Hashset.Strings([]string{"b", "c"})

		// Act
		hs1.AddHashsetItems(hs2)

		// Assert
		if hs1.Length() != 3 {
			t.Errorf("AddHashsetItems expected 3, got %d", hs1.Length())
		}
	})
}

// ==========================================
// Hashset - ConcatNew
// ==========================================

func Test_ExtHashset_ConcatNewHashsets_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_ConcatNewHashsets_Verification", func() {
		// Arrange
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hs2 := corestr.New.Hashset.Strings([]string{"b", "c"})

		// Act
		result := hs1.ConcatNewHashsets(true, hs2)

		// Assert
		if result.Length() < 2 {
			t.Errorf("ConcatNewHashsets expected >= 2, got %d", result.Length())
		}
	})
}

func Test_ExtHashset_ConcatNewStrings_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_ConcatNewStrings_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.ConcatNewStrings(true, []string{"b", "c"})

		// Assert
		if result.Length() < 2 {
			t.Errorf("ConcatNewStrings expected >= 2, got %d", result.Length())
		}
	})
}

// ==========================================
// Hashset - IsEqual
// ==========================================

func Test_ExtHashset_IsEqual_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_IsEqual_Verification", func() {
		// Arrange
		hs1 := corestr.New.Hashset.Strings([]string{"a", "b"})
		hs2 := corestr.New.Hashset.Strings([]string{"b", "a"})
		hs3 := corestr.New.Hashset.Strings([]string{"a", "c"})

		// Act & Assert
		if !hs1.IsEqual(hs2) {
			t.Error("Same content should be equal")
		}
		if hs1.IsEqual(hs3) {
			t.Error("Different content should not be equal")
		}
	})
}

// ==========================================
// ValidValue
// ==========================================

func Test_ExtValidValue_Verification(t *testing.T) {
	safeTest(t, "Test_ExtValidValue_Verification", func() {
		for caseIndex, testCase := range extValidValueTestCases {
			// Arrange
			input := testCase.ArrangeInput.(args.Map)
			value, _ := input.GetAsString("value")
			isValidRaw, _ := input.Get("isValid")
			isValid := isValidRaw == true

			// Act
			var vv corestr.ValidValue
			if isValid {
				vv = corestr.ValidValue{Value: value, IsValid: true}
			} else {
				vv = corestr.ValidValue{Value: value, IsValid: false}
			}

			actual := args.Map{
				"value":   vv.Value,
				"isValid": fmt.Sprintf("%v", vv.IsValid),
			}

			// Assert
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

// ==========================================
// LeftRightFromSplit / LeftMiddleRightFromSplit (cover the corestr funcs)
// ==========================================

func Test_ExtLeftRightFromSplit_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRightFromSplit_Verification", func() {
		// Arrange
		// Act
		lr := corestr.LeftRightFromSplit("key=value", "=")

		// Assert
		if lr.Left != "key" || lr.Right != "value" {
			t.Errorf("LeftRightFromSplit expected key/value, got %s/%s", lr.Left, lr.Right)
		}
	})
}

func Test_ExtLeftMiddleRightFromSplit_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftMiddleRightFromSplit_Verification", func() {
		// Arrange
		// Act
		lmr := corestr.LeftMiddleRightFromSplit("a/b/c", "/")

		// Assert
		if lmr.Left != "a" || lmr.Right != "c" {
			t.Errorf("LeftMiddleRightFromSplit expected a/c, got %s/%s", lmr.Left, lmr.Right)
		}
	})
}

// ==========================================
// Collection - JsonString
// ==========================================

func Test_ExtCollection_JsonString_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_JsonString_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		result := col.JsonString()

		// Assert — JsonPtr uses value receiver but slice is a reference type,
		// so the copy retains the underlying items and serialization works.
		if result == "" {
			t.Error("JsonString should be non-empty for a populated collection")
		}
	})
}

// ==========================================
// LeftRightTrimmedUsingSlice
// ==========================================

func Test_ExtLeftRightTrimmedUsingSlice_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRightTrimmedUsingSlice_Verification", func() {
		// Arrange
		// Act
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})

		// Assert
		if lr.Left != "a" || lr.Right != "b" {
			t.Errorf("LeftRightTrimmedUsingSlice expected a/b, got %s/%s", lr.Left, lr.Right)
		}

		// Act - single element
		lr2 := corestr.LeftRightTrimmedUsingSlice([]string{"only"})

		// Assert
		if lr2.Left != "only" || lr2.IsValid {
			t.Error("Single element should not be valid")
		}

		// Act - empty
		lr3 := corestr.LeftRightTrimmedUsingSlice([]string{})

		// Assert
		if lr3.IsValid {
			t.Error("Empty slice should not be valid")
		}

		// Act - nil
		lr4 := corestr.LeftRightTrimmedUsingSlice(nil)

		// Assert
		if lr4.IsValid {
			t.Error("nil slice should not be valid")
		}
	})
}

// ==========================================
// Hashset - Filter / Diff
// ==========================================

func Test_ExtHashset_Filter_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_Filter_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"apple", "banana", "avocado"})

		// Act
		filtered := hs.Filter(func(s string) bool {
			return strings.HasPrefix(s, "a")
		})

		// Assert
		if filtered.Length() != 2 {
			t.Errorf("Filter expected 2 items starting with 'a', got %d", filtered.Length())
		}
	})
}

// ==========================================
// Collection - AddLock / AddsLock
// ==========================================

func Test_ExtCollection_AddLock_Verification(t *testing.T) {
	safeTest(t, "Test_ExtCollection_AddLock_Verification", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)

		// Act
		col.AddLock("a")
		col.AddsLock("b", "c")

		// Assert
		if col.Length() != 3 {
			t.Errorf("AddLock/AddsLock expected 3, got %d", col.Length())
		}
	})
}

// ==========================================
// emptyCreator
// ==========================================

func Test_ExtEmptyCreator_Verification(t *testing.T) {
	safeTest(t, "Test_ExtEmptyCreator_Verification", func() {
		// Arrange
		// Act
		col := corestr.Empty.Collection()
		hs := corestr.Empty.Hashset()
		hm := corestr.Empty.Hashmap()

		// Assert
		if !col.IsEmpty() {
			t.Error("Empty.Collection should be empty")
		}
		if !hs.IsEmpty() {
			t.Error("Empty.Hashset should be empty")
		}
		if !hm.IsEmpty() {
			t.Error("Empty.Hashmap should be empty")
		}
	})
}
