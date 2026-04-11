package coregenerictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Test: Pair — NewPair valid
// ==========================================

func Test_Pair_NewPair_Valid(t *testing.T) {
	for caseIndex, testCase := range pairNewValidTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		// Act
		pair := coregeneric.NewPair(left, right)
		actual := args.Map{
			"left":         pair.Left,
			"right":        pair.Right,
			"isValid":      pair.IsValid,
			"errorMessage": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// ==========================================
// Test: Pair — InvalidPair
// ==========================================

func Test_Pair_InvalidPair(t *testing.T) {
	for caseIndex, testCase := range pairInvalidTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		message, _ := input.GetAsString("message")

		// Act
		var pair *coregeneric.Pair[string, string]
		if message == "" {
			pair = coregeneric.InvalidPairNoMessage[string, string]()
		} else {
			pair = coregeneric.InvalidPair[string, string](message)
		}

		actual := args.Map{
			"left":         pair.Left,
			"right":        pair.Right,
			"isValid":      pair.IsValid,
			"errorMessage": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// ==========================================
// Test: Pair — Clone independence
// ==========================================

func Test_Pair_Clone_Independence(t *testing.T) {
	for caseIndex, testCase := range pairCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		// Act
		original := coregeneric.NewPair(left, right)
		cloned := original.Clone()
		cloned.Left = "mutated-left"

		actual := args.Map{
			"clonedLeft":            original.Left,
			"clonedRight":           original.Right,
			"isValid":               original.IsValid,
			"originalAfterMutation": cloned.Left,
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// ==========================================
// Test: Pair — nil Clone
// ==========================================

func Test_Pair_Nil_Clone(t *testing.T) {
	for caseIndex, testCase := range pairNilCloneTestCases {
		// Act
		var pair *coregeneric.Pair[string, string]
		cloned := pair.Clone()

		actLines := []string{
			fmt.Sprintf("%v", cloned == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Pair — IsEqual
// ==========================================

func Test_Pair_IsEqual_Same(t *testing.T) {
	tc := pairIsEqualSameTestCase
	input := tc.ArrangeInput.(args.Map)
	left, _ := input.GetAsString("left")
	right, _ := input.GetAsString("right")

	a := coregeneric.NewPair(left, right)
	b := coregeneric.NewPair(left, right)

	// Assert
	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

func Test_Pair_IsEqual_DiffLeft(t *testing.T) {
	tc := pairIsEqualDiffLeftTestCase
	a := coregeneric.NewPair("a", "b")
	b := coregeneric.NewPair("x", "b")

	// Assert
	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

func Test_Pair_IsEqual_NilVsNonNil(t *testing.T) {
	tc := pairIsEqualNilVsNonNilTestCase
	a := coregeneric.NewPair("a", "b")
	var b *coregeneric.Pair[string, string]

	// Assert
	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

func Test_Pair_IsEqual_BothNil(t *testing.T) {
	tc := pairIsEqualBothNilTestCase
	var a *coregeneric.Pair[string, string]
	var b *coregeneric.Pair[string, string]

	// Assert
	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

// ==========================================
// Test: Pair — Values()
// ==========================================

func Test_Pair_Values(t *testing.T) {
	for caseIndex, testCase := range pairValuesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		// Act
		pair := coregeneric.NewPair(left, right)
		l, r := pair.Values()
		actual := args.Map{
			"left":  l,
			"right": r,
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// ==========================================
// Test: Pair — Clear
// ==========================================

func Test_Pair_Clear(t *testing.T) {
	for caseIndex, testCase := range pairClearTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		// Act
		pair := coregeneric.NewPair(left, right)
		pair.Clear()
		actual := args.Map{
			"clearedLeft":  pair.Left,
			"clearedRight": pair.Right,
			"isValid":      pair.IsValid,
			"errorMessage": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// ==========================================
// Test: Triple — NewTriple valid
// ==========================================

func Test_Triple_NewTriple_Valid(t *testing.T) {
	for caseIndex, testCase := range tripleNewValidTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		middle, _ := input.GetAsString("middle")
		right, _ := input.GetAsString("right")

		// Act
		triple := coregeneric.NewTriple(left, middle, right)
		actual := args.Map{
			"left":         triple.Left,
			"middle":       triple.Middle,
			"right":        triple.Right,
			"isValid":      triple.IsValid,
			"errorMessage": triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// ==========================================
// Test: Triple — InvalidTriple
// ==========================================

func Test_Triple_InvalidTriple(t *testing.T) {
	for caseIndex, testCase := range tripleInvalidTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		message, _ := input.GetAsString("message")

		// Act
		var triple *coregeneric.Triple[string, string, string]
		if message == "" {
			triple = coregeneric.InvalidTripleNoMessage[string, string, string]()
		} else {
			triple = coregeneric.InvalidTriple[string, string, string](message)
		}

		actual := args.Map{
			"left":         triple.Left,
			"middle":       triple.Middle,
			"right":        triple.Right,
			"isValid":      triple.IsValid,
			"errorMessage": triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// ==========================================
// Test: Triple — Clone independence
// ==========================================

func Test_Triple_Clone_Independence(t *testing.T) {
	for caseIndex, testCase := range tripleCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		middle, _ := input.GetAsString("middle")
		right, _ := input.GetAsString("right")

		// Act
		original := coregeneric.NewTriple(left, middle, right)
		cloned := original.Clone()
		cloned.Left = "mutated"

		actual := args.Map{
			"clonedLeft":            original.Left,
			"clonedMiddle":          original.Middle,
			"clonedRight":           original.Right,
			"isValid":               original.IsValid,
			"originalAfterMutation": cloned.Left,
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// ==========================================
// Test: Triple — nil Clone
// ==========================================

func Test_Triple_Nil_Clone(t *testing.T) {
	for caseIndex, testCase := range tripleNilCloneTestCases {
		// Act
		var triple *coregeneric.Triple[string, string, string]
		cloned := triple.Clone()

		actLines := []string{
			fmt.Sprintf("%v", cloned == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Triple — Values()
// ==========================================

func Test_Triple_Values(t *testing.T) {
	for caseIndex, testCase := range tripleValuesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		middle, _ := input.GetAsString("middle")
		right, _ := input.GetAsString("right")

		// Act
		triple := coregeneric.NewTriple(left, middle, right)
		a, b, c := triple.Values()
		actual := args.Map{
			"left":   a,
			"middle": b,
			"right":  c,
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// ==========================================
// Test: Triple — Clear
// ==========================================

func Test_Triple_Clear(t *testing.T) {
	for caseIndex, testCase := range tripleClearTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		middle, _ := input.GetAsString("middle")
		right, _ := input.GetAsString("right")

		// Act
		triple := coregeneric.NewTriple(left, middle, right)
		triple.Clear()
		actual := args.Map{
			"clearedLeft":   triple.Left,
			"clearedMiddle": triple.Middle,
			"clearedRight":  triple.Right,
			"isValid":       triple.IsValid,
			"errorMessage":  triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// ==========================================
// Test: New.Pair Creator shortcuts
// ==========================================

func Test_New_Pair_Creator(t *testing.T) {
	// Arrange
	// StringString
	p := coregeneric.New.Pair.StringString("k", "v")

	// Act
	actual := args.Map{"result": p.Left != "k" || p.Right != "v" || !p.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "New.Pair.StringString failed:", actual)

	// StringInt
	pi := coregeneric.New.Pair.StringInt("age", 30)
	actual = args.Map{"result": pi.Left != "age" || pi.Right != 30 || !pi.IsValid}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "New.Pair.StringInt failed:", actual)

	// Any
	pa := coregeneric.New.Pair.Any("x", 42)
	actual = args.Map{"result": pa.Left != "x" || pa.Right != 42 || !pa.IsValid}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "New.Pair.Any failed:", actual)

	// InvalidStringString
	inv := coregeneric.New.Pair.InvalidStringString("err")
	actual = args.Map{"result": inv.IsValid || inv.Message != "err"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "New.Pair.InvalidStringString failed:", actual)
}

// ==========================================
// Test: New.Triple Creator shortcuts
// ==========================================

func Test_New_Triple_Creator(t *testing.T) {
	// Arrange
	// StringStringString
	tr := coregeneric.New.Triple.StringStringString("a", "b", "c")

	// Act
	actual := args.Map{"result": tr.Left != "a" || tr.Middle != "b" || tr.Right != "c" || !tr.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "New.Triple.StringStringString failed:", actual)

	// Any
	ta := coregeneric.New.Triple.Any("x", 1, true)
	actual = args.Map{"result": ta.Left != "x" || ta.Middle != 1 || ta.Right != true || !ta.IsValid}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "New.Triple.Any failed:", actual)

	// InvalidStringStringString
	inv := coregeneric.New.Triple.InvalidStringStringString("bad")
	actual = args.Map{"result": inv.IsValid || inv.Message != "bad"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "New.Triple.InvalidStringStringString failed:", actual)
}
