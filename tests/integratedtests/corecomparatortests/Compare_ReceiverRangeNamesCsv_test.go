package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Cover receiver method RangeNamesCsv (vs package-level func) ──

func Test_Compare_ReceiverRangeNamesCsv(t *testing.T) {
	// Arrange
	csv := corecomparator.Equal.RangeNamesCsv()

	// Act
	actual := args.Map{"notEmpty": csv != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Compare.RangeNamesCsv receiver -- not empty", actual)
}

// ── IsCompareEqualLogically: branch where expectedCompare is Equal but it is not ──

func Test_IsCompareEqualLogically_ExpectedEqual_ItNotEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"result": corecomparator.LeftGreater.IsCompareEqualLogically(corecomparator.Equal),
	}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsCompareEqualLogically expected=Equal it=LeftGreater -- false", actual)
}

// ── OnlySupportedDirectErr error path ──

func Test_OnlySupportedDirectErr_NotMatching(t *testing.T) {
	// Arrange
	err := corecomparator.Inconclusive.OnlySupportedDirectErr(corecomparator.Equal, corecomparator.LeftGreater)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedDirectErr no match -- error", actual)
}

// ── IsLeftLessEqualLogically on non-less values ──

func Test_IsLeftLessEqualLogically_Greater(t *testing.T) {
	// Act
	actual := args.Map{
		"greater": corecomparator.LeftGreater.IsLeftLessEqualLogically(),
		"notEq":   corecomparator.NotEqual.IsLeftLessEqualLogically(),
	}

	// Assert
	expected := args.Map{
		"greater": false,
		"notEq": false,
	}
	expected.ShouldBeEqual(t, 0, "IsLeftLessEqualLogically non-less values -- false", actual)
}

// ── IsLeftGreaterEqualLogically on non-greater values ──

func Test_IsLeftGreaterEqualLogically_Less(t *testing.T) {
	// Act
	actual := args.Map{
		"less":  corecomparator.LeftLess.IsLeftGreaterEqualLogically(),
		"notEq": corecomparator.NotEqual.IsLeftGreaterEqualLogically(),
	}

	// Assert
	expected := args.Map{
		"less": false,
		"notEq": false,
	}
	expected.ShouldBeEqual(t, 0, "IsLeftGreaterEqualLogically non-greater values -- false", actual)
}

// ── MinLength equal values ──

func Test_MinLength_Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecomparator.MinLength(3, 3)}

	// Assert
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "MinLength equal -- returns same", actual)
}
