package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
)

// ==========================================
// Condition.IsSplitByWhitespace
// ==========================================

func Test_Condition_IsSplitByWhitespace_AllFalse_FromCondition(t *testing.T) {
	// Arrange
	tc := conditionAllFalseTestCase
	c := corevalidator.Condition{}

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Condition_IsSplitByWhitespace_UniqueWordOnly(t *testing.T) {
	// Arrange
	tc := conditionUniqueWordOnlyTestCase
	c := corevalidator.Condition{IsUniqueWordOnly: true}

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Condition_IsSplitByWhitespace_NonEmptyWhitespace(t *testing.T) {
	// Arrange
	tc := conditionNonEmptyWhitespaceTestCase
	c := corevalidator.Condition{IsNonEmptyWhitespace: true}

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Condition_IsSplitByWhitespace_SortBySpace(t *testing.T) {
	// Arrange
	tc := conditionSortBySpaceTestCase
	c := corevalidator.Condition{IsSortStringsBySpace: true}

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Condition_IsSplitByWhitespace_TrimOnlyNotEnough(t *testing.T) {
	// Arrange
	tc := conditionTrimOnlyTestCase
	c := corevalidator.Condition{IsTrimCompare: true}

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Preset Conditions
// ==========================================

func Test_DefaultDisabledCondition_NoSplit(t *testing.T) {
	// Arrange
	tc := conditionDisabledTestCase
	c := corevalidator.DefaultDisabledCoreCondition

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_DefaultTrimCondition_NoSplit(t *testing.T) {
	// Arrange
	tc := conditionTrimTestCase
	c := corevalidator.DefaultTrimCoreCondition

	// Act
	actual := args.Map{
		"isSplit":       c.IsSplitByWhitespace(),
		"isTrimCompare": c.IsTrimCompare,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_DefaultSortTrimCondition_Split(t *testing.T) {
	// Arrange
	tc := conditionSortTrimTestCase
	c := corevalidator.DefaultSortTrimCoreCondition

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_DefaultUniqueWordsCondition_Split(t *testing.T) {
	// Arrange
	tc := conditionUniqueWordsTestCase
	c := corevalidator.DefaultUniqueWordsCoreCondition

	// Act
	actual := args.Map{
		"isSplit":          c.IsSplitByWhitespace(),
		"isUniqueWordOnly": c.IsUniqueWordOnly,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
