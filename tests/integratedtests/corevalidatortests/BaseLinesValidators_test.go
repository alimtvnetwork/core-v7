package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// BaseLinesValidators
// ==========================================

func Test_BaseLinesValidators_Empty(t *testing.T) {
	// Arrange
	b := corevalidator.BaseLinesValidators{}

	// Act
	actual := args.Map{"result": b.IsEmptyLinesValidators()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should be empty", actual)
	actual = args.Map{"result": b.HasLinesValidators()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not have validators", actual)
	actual = args.Map{"result": b.LinesValidatorsLength() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_BaseLinesValidators_WithItems_FromBaseLinesValidators(t *testing.T) {
	// Arrange
	b := corevalidator.BaseLinesValidators{
		LinesValidators: []corevalidator.LineValidator{
			{
				LineNumber: corevalidator.LineNumber{LineNumber: -1},
				TextValidator: corevalidator.TextValidator{
					Search:    "a",
					SearchAs:  stringcompareas.Equal,
					Condition: corevalidator.DefaultDisabledCoreCondition,
				},
			},
		},
	}

	// Act
	actual := args.Map{"result": b.IsEmptyLinesValidators()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	actual = args.Map{"result": b.HasLinesValidators()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have validators", actual)
	actual = args.Map{"result": b.LinesValidatorsLength() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_BaseLinesValidators_ToLinesValidators_Empty_FromBaseLinesValidators(t *testing.T) {
	// Arrange
	b := corevalidator.BaseLinesValidators{}
	lv := b.ToLinesValidators()

	// Act
	actual := args.Map{"result": lv == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	actual = args.Map{"result": lv.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
}

func Test_BaseLinesValidators_ToLinesValidators_NonEmpty(t *testing.T) {
	// Arrange
	b := corevalidator.BaseLinesValidators{
		LinesValidators: []corevalidator.LineValidator{
			{
				LineNumber: corevalidator.LineNumber{LineNumber: 0},
				TextValidator: corevalidator.TextValidator{
					Search:    "test",
					SearchAs:  stringcompareas.Equal,
					Condition: corevalidator.DefaultDisabledCoreCondition,
				},
			},
			{
				LineNumber: corevalidator.LineNumber{LineNumber: 1},
				TextValidator: corevalidator.TextValidator{
					Search:    "test2",
					SearchAs:  stringcompareas.Equal,
					Condition: corevalidator.DefaultDisabledCoreCondition,
				},
			},
		},
	}
	lv := b.ToLinesValidators()

	// Act
	actual := args.Map{"result": lv.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// (nil receiver tests migrated to BaseLinesValidators_NilReceiver_testcases.go)

// ==========================================
// LinesValidators — collection
// ==========================================

func Test_LinesValidators_New(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)

	// Act
	actual := args.Map{"result": lv == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	actual = args.Map{"result": lv.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "new should be empty", actual)
}

func Test_LinesValidators_Add_FromBaseLinesValidators(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})

	// Act
	actual := args.Map{"result": lv.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": lv.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
}

func Test_LinesValidators_AddPtr_Nil_FromBaseLinesValidators(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.AddPtr(nil)

	// Act
	actual := args.Map{"result": lv.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil add should not increase length", actual)
}

func Test_LinesValidators_HasIndex(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{})

	// Act
	actual := args.Map{"result": lv.HasIndex(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have index 0", actual)
	actual = args.Map{"result": lv.HasIndex(1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have index 1", actual)
}

// (nil receiver tests migrated to BaseLinesValidators_NilReceiver_testcases.go)

// ==========================================
// LinesValidators.IsMatchText
// ==========================================

func Test_LinesValidators_IsMatchText_Empty_FromBaseLinesValidators(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(0)

	// Act
	actual := args.Map{"result": lv.IsMatchText("anything", true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty validators should match any text", actual)
}

func Test_LinesValidators_IsMatchText_Match_FromBaseLinesValidators(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Contains,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})

	// Act
	actual := args.Map{"result": lv.IsMatchText("hello world", true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "contains should match", actual)
}

func Test_LinesValidators_IsMatchText_NoMatch(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "xyz",
			SearchAs:  stringcompareas.Contains,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})

	// Act
	actual := args.Map{"result": lv.IsMatchText("hello world", true)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing substring should not match", actual)
}

// ==========================================
// BaseValidatorCoreCondition
// ==========================================

func Test_BaseValidatorCoreCondition_Default_NilPtr(t *testing.T) {
	// Arrange
	b := corevalidator.BaseValidatorCoreCondition{}
	c := b.ValidatorCoreConditionDefault()

	// Act
	actual := args.Map{"result": c.IsTrimCompare || c.IsUniqueWordOnly}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "default condition should have all false", actual)
	// should set the ptr
	actual = args.Map{"result": b.ValidatorCoreCondition == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have set the pointer", actual)
}

func Test_BaseValidatorCoreCondition_Default_ExistingPtr(t *testing.T) {
	// Arrange
	cond := corevalidator.Condition{IsTrimCompare: true}
	b := corevalidator.BaseValidatorCoreCondition{
		ValidatorCoreCondition: &cond,
	}
	c := b.ValidatorCoreConditionDefault()

	// Act
	actual := args.Map{"result": c.IsTrimCompare}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should return existing condition", actual)
}
