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
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty expected 2 items", actual)
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
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyWhitespace expected 1", actual)
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
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddIf expected 1", actual)
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
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddIfMany expected 2", actual)
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
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddFunc expected 1", actual)
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
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddError expected 1", actual)
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
		actual := args.Map{"result": col1.IsEquals(col2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Same content should be equal", actual)
		actual = args.Map{"result": col1.IsEquals(col3)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Different case should not be equal (case-sensitive)", actual)
		actual = args.Map{"result": col1.IsEqualsWithSensitive(false, col3)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Case-insensitive should match", actual)
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
		actual := args.Map{"result": col.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "nil.IsEmpty() should be true", actual)
		actual = args.Map{"result": col.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil.Length() should be 0", actual)
		actual = args.Map{"result": col.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil.HasItems() should be false", actual)
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
		actual := args.Map{"result": ok || col.Length() != 2}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "RemoveAt(1) expected success, got ok= len=", actual)

		// Act - out of range
		ok2 := col.RemoveAt(10)

		// Assert
		actual = args.Map{"result": ok2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "RemoveAt(10) should return false", actual)

		// Act - negative
		ok3 := col.RemoveAt(-1)

		// Assert
		actual = args.Map{"result": ok3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "RemoveAt(-1) should return false", actual)
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
		actual := args.Map{"result": result.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ConcatNew expected 3", actual)

		// Act - empty
		result2 := col.ConcatNew(0)

		// Assert
		actual = args.Map{"result": result2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty expected 1", actual)
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
		actual := args.Map{"result": err == nil || err.Error() != "err1; err2"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AsError expected 'err1; err2', got ''", actual)

		// Act - empty
		emptyCol := corestr.New.Collection.Cap(0)
		err2 := emptyCol.AsDefaultError()

		// Assert
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AsDefaultError on empty should return nil", actual)
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
		actual := args.Map{"result": len(result) != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "EachItemSplitBy expected 4", actual)
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
		actual := args.Map{"result": col.HasIndex(0)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasIndex(0) should be true", actual)
		actual = args.Map{"result": col.HasIndex(2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasIndex(2) should be true", actual)
		actual = args.Map{"result": col.HasIndex(3)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasIndex(3) should be false", actual)
		actual = args.Map{"result": col.LastIndex() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LastIndex expected 2", actual)
		actual = args.Map{"result": col.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAnyItem should be true", actual)
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
		actual := args.Map{"result": col1.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddCollection expected 3", actual)
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
		actual := args.Map{"result": col1.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddCollections expected 3", actual)
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
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddIf expected 1", actual)
	})
}

func Test_ExtSimpleSlice_FirstLast_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_FirstLast_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act & Assert
		actual := args.Map{"result": ss.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "First expected 'a', got ''", actual)
		actual = args.Map{"result": ss.Last() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Last expected 'c', got ''", actual)
		actual = args.Map{"result": ss.FirstOrDefault() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault expected 'a', got ''", actual)
		actual = args.Map{"result": ss.LastOrDefault() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LastOrDefault expected 'c', got ''", actual)
	})
}

func Test_ExtSimpleSlice_EmptyDefaults_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_EmptyDefaults_Verification", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act & Assert
		actual := args.Map{"result": ss.FirstOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil.FirstOrDefault() should return empty", actual)
		actual = args.Map{"result": ss.LastOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil.LastOrDefault() should return empty", actual)
		actual = args.Map{"result": ss.IsEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "nil.IsEmpty() should be true", actual)
		actual = args.Map{"result": ss.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil.Length() should be 0", actual)
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
		actual := args.Map{"result": len(skipped) != 2 || skipped[0] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Skip(2) expected [c,d]", actual)
		actual = args.Map{"result": len(taken) != 2 || taken[0] != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Take(2) expected [a,b]", actual)
	})
}

func Test_ExtSimpleSlice_IsContains_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_IsContains_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"hello", "world"})

		// Act & Assert
		actual := args.Map{"result": ss.IsContains("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsContains should find 'hello'", actual)
		actual = args.Map{"result": ss.IsContains("missing")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsContains should not find 'missing'", actual)
	})
}

func Test_ExtSimpleSlice_IndexOf_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_IndexOf_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act & Assert
		actual := args.Map{"result": ss.IndexOf("b") != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IndexOf('b') expected 1", actual)
		actual = args.Map{"result": ss.IndexOf("z") != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IndexOf('z') expected -1", actual)
	})
}

func Test_ExtSimpleSlice_InsertAt_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_InsertAt_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "c"})

		// Act
		ss.InsertAt(1, "b")

		// Assert
		actual := args.Map{"result": ss.Length() != 3 || (*ss)[1] != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "InsertAt expected [a,b,c]", actual)
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
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddError expected 1", actual)
	})
}

func Test_ExtSimpleSlice_AsError_Verification(t *testing.T) {
	safeTest(t, "Test_ExtSimpleSlice_AsError_Verification", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"e1", "e2"})

		// Act
		err := ss.AsError(", ")

		// Assert
		actual := args.Map{"result": err == nil || err.Error() != "e1, e2"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AsError expected 'e1, e2', got ''", actual)

		// Act - nil empty
		var nilSS *corestr.SimpleSlice
		err2 := nilSS.AsError(", ")

		// Assert
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil.AsError should return nil", actual)
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
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AppendFmt expected 2 items", actual)
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
		actual := args.Map{"result": count != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CountFunc expected 2", actual)
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
		actual := args.Map{"result": lr.IsLeftEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsLeftEmpty should be false", actual)
		actual = args.Map{"result": lr.IsRightEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsRightEmpty should be false", actual)
		actual = args.Map{"result": lr.HasSafeNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasSafeNonEmpty should be true", actual)
		actual = args.Map{"result": lr.LeftTrim() != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LeftTrim expected 'hello', got ''", actual)
		actual = args.Map{"result": lr.RightTrim() != "world"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "RightTrim expected 'world', got ''", actual)
		actual = args.Map{"result": string(lr.LeftBytes()) != " hello "}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LeftBytes mismatch", actual)
		actual = args.Map{"result": string(lr.RightBytes()) != " world "}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "RightBytes mismatch", actual)
	})
}

func Test_ExtLeftRight_IsEqual_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_IsEqual_Verification", func() {
		// Arrange
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("a", "c")

		// Act & Assert
		actual := args.Map{"result": lr1.IsEqual(lr2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Same content should be equal", actual)
		actual = args.Map{"result": lr1.IsEqual(lr3)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Different content should not be equal", actual)

		// nil cases
		var nilLR *corestr.LeftRight
		actual = args.Map{"result": nilLR.IsEqual(nil)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "nil.IsEqual(nil) should be true", actual)
		actual = args.Map{"result": lr1.IsEqual(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "non-nil.IsEqual(nil) should be false", actual)
	})
}

func Test_ExtLeftRight_Clone_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_Clone_Verification", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		cloned := lr.Clone()

		// Assert
		actual := args.Map{"result": cloned.Left != "a" || cloned.Right != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Clone should copy values", actual)
	})
}

func Test_ExtLeftRight_RegexMatch_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_RegexMatch_Verification", func() {
		// Arrange
		lr := corestr.NewLeftRight("abc123", "def456")
		re := regexp.MustCompile(`\d+`)

		// Act & Assert
		actual := args.Map{"result": lr.IsLeftRegexMatch(re)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsLeftRegexMatch should be true", actual)
		actual = args.Map{"result": lr.IsRightRegexMatch(re)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsRightRegexMatch should be true", actual)
		actual = args.Map{"result": lr.IsLeftRegexMatch(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil regex should return false", actual)
	})
}

func Test_ExtLeftRight_Is_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_Is_Verification", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act & Assert
		actual := args.Map{"result": lr.IsLeft("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsLeft should be true", actual)
		actual = args.Map{"result": lr.IsRight("b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsRight should be true", actual)
		actual = args.Map{"result": lr.Is("a", "b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Is should be true", actual)
	})
}

func Test_ExtLeftRight_InvalidCreation_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_InvalidCreation_Verification", func() {
		// Arrange
		// Act
		lr1 := corestr.InvalidLeftRightNoMessage()
		lr2 := corestr.InvalidLeftRight("test error")

		// Assert
		actual := args.Map{"result": lr1.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRightNoMessage should be invalid", actual)
		actual = args.Map{"result": lr2.IsValid || lr2.Message == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRight should be invalid with message", actual)
	})
}

func Test_ExtLeftRight_Dispose_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftRight_Dispose_Verification", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		lr.Dispose()

		// Assert
		actual := args.Map{"result": lr.Left != "" || lr.Right != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Dispose should clear values", actual)
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
		actual := args.Map{"result": lmr.LeftTrim() != "L"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LeftTrim expected 'L', got ''", actual)
		actual = args.Map{"result": lmr.MiddleTrim() != "M"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "MiddleTrim expected 'M', got ''", actual)
		actual = args.Map{"result": lmr.RightTrim() != "R"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "RightTrim expected 'R', got ''", actual)
		actual = args.Map{"result": string(lmr.LeftBytes()) != " L "}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LeftBytes mismatch", actual)
		actual = args.Map{"result": string(lmr.MiddleBytes()) != " M "}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "MiddleBytes mismatch", actual)
		actual = args.Map{"result": string(lmr.RightBytes()) != " R "}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "RightBytes mismatch", actual)
		actual = args.Map{"result": lmr.HasSafeNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasSafeNonEmpty should be true", actual)
		actual = args.Map{"result": lmr.IsAll(" L ", " M ", " R ")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsAll should be true", actual)
	})
}

func Test_ExtLeftMiddleRight_ToLeftRight_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftMiddleRight_ToLeftRight_Verification", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("L", "M", "R")

		// Act
		lr := lmr.ToLeftRight()

		// Assert
		actual := args.Map{"result": lr.Left != "L" || lr.Right != "R"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ToLeftRight expected L/R, got/", actual)
	})
}

func Test_ExtLeftMiddleRight_Clone_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftMiddleRight_Clone_Verification", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("L", "M", "R")

		// Act
		cloned := lmr.Clone()

		// Assert
		actual := args.Map{"result": cloned.Left != "L" || cloned.Middle != "M" || cloned.Right != "R"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Clone should copy all values", actual)
	})
}

func Test_ExtLeftMiddleRight_Invalid_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftMiddleRight_Invalid_Verification", func() {
		// Arrange
		// Act
		lmr1 := corestr.InvalidLeftMiddleRightNoMessage()
		lmr2 := corestr.InvalidLeftMiddleRight("test")

		// Assert
		actual := args.Map{"result": lmr1.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRightNoMessage should be invalid", actual)
		actual = args.Map{"result": lmr2.IsValid || lmr2.Message == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRight should be invalid with message", actual)
	})
}

func Test_ExtLeftMiddleRight_Dispose_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftMiddleRight_Dispose_Verification", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("L", "M", "R")

		// Act
		lmr.Dispose()

		// Assert
		actual := args.Map{"result": lmr.Left != "" || lmr.Middle != "" || lmr.Right != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Dispose should clear all values", actual)
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
		actual := args.Map{"result": hs.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After adds expected 3", actual)

		// Act
		hs.Remove("b")

		// Assert
		actual = args.Map{"result": hs.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After remove expected 2", actual)
		actual = args.Map{"result": hs.Has("b")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Should not have 'b' after remove", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty expected 1", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyWhitespace expected 1", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddIf expected 1", actual)
	})
}

func Test_ExtHashset_List_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_List_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"b", "a", "c"})

		// Act
		list := hs.List()

		// Assert
		actual := args.Map{"result": len(list) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "List expected 3 items", actual)
	})
}

func Test_ExtHashset_SortedList_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_SortedList_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})

		// Act
		sorted := hs.SortedList()

		// Assert
		actual := args.Map{"result": len(sorted) != 3 || sorted[0] != "a" || sorted[1] != "b" || sorted[2] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "SortedList expected [a,b,c]", actual)
	})
}

func Test_ExtHashset_HasAll_HasAny_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_HasAll_HasAny_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})

		// Act & Assert
		actual := args.Map{"result": hs.HasAll("a", "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAll should be true", actual)
		actual = args.Map{"result": hs.HasAll("a", "z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasAll should be false with missing", actual)
		actual = args.Map{"result": hs.HasAny("z", "a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAny should be true", actual)
		actual = args.Map{"result": hs.HasAny("x", "y")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasAny should be false", actual)
	})
}

func Test_ExtHashset_NilReceiver_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_NilReceiver_Verification", func() {
		// Arrange
		var hs *corestr.Hashset

		// Act & Assert
		actual := args.Map{"result": hs.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "nil.IsEmpty() should be true", actual)
		actual = args.Map{"result": hs.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil.HasItems() should be false", actual)
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
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddCollection expected 2", actual)
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
		actual := args.Map{"result": hs1.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems expected 3", actual)
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
		actual := args.Map{"result": result.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ConcatNewHashsets expected >= 2", actual)
	})
}

func Test_ExtHashset_ConcatNewStrings_Verification(t *testing.T) {
	safeTest(t, "Test_ExtHashset_ConcatNewStrings_Verification", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.ConcatNewStrings(true, []string{"b", "c"})

		// Assert
		actual := args.Map{"result": result.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ConcatNewStrings expected >= 2", actual)
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
		actual := args.Map{"result": hs1.IsEqual(hs2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Same content should be equal", actual)
		actual = args.Map{"result": hs1.IsEqual(hs3)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Different content should not be equal", actual)
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
		actual := args.Map{"result": lr.Left != "key" || lr.Right != "value"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit expected key/value, got/", actual)
	})
}

func Test_ExtLeftMiddleRightFromSplit_Verification(t *testing.T) {
	safeTest(t, "Test_ExtLeftMiddleRightFromSplit_Verification", func() {
		// Arrange
		// Act
		lmr := corestr.LeftMiddleRightFromSplit("a/b/c", "/")

		// Assert
		actual := args.Map{"result": lmr.Left != "a" || lmr.Right != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit expected a/c, got/", actual)
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
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "JsonString should be non-empty for a populated collection", actual)
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
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice expected a/b, got/", actual)

		// Act - single element
		lr2 := corestr.LeftRightTrimmedUsingSlice([]string{"only"})

		// Assert
		actual = args.Map{"result": lr2.Left != "only" || lr2.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Single element should not be valid", actual)

		// Act - empty
		lr3 := corestr.LeftRightTrimmedUsingSlice([]string{})

		// Assert
		actual = args.Map{"result": lr3.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty slice should not be valid", actual)

		// Act - nil
		lr4 := corestr.LeftRightTrimmedUsingSlice(nil)

		// Assert
		actual = args.Map{"result": lr4.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil slice should not be valid", actual)
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
		actual := args.Map{"result": filtered.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Filter expected 2 items starting with 'a'", actual)
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
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddLock/AddsLock expected 3", actual)
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
		actual := args.Map{"result": col.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Empty.Collection should be empty", actual)
		actual = args.Map{"result": hs.IsEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Empty.Hashset should be empty", actual)
		actual = args.Map{"result": hm.IsEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Empty.Hashmap should be empty", actual)
	})
}
