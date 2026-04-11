package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

func Test_SimpleSliceValidator_SetActual_FromSimpleSliceValidator(t *testing.T) {
	tc := simpleSliceValidatorSetActualTestCase

	// Arrange
	expected := corestr.New.SimpleSlice.Direct(false, []string{"a", "b"})
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}

	// Act
	result := v.SetActual([]string{"a", "b"})

	actual := args.Map{
		"sameInstance": result == v,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleSliceValidator_SliceValidator_FromSimpleSliceValidator(t *testing.T) {
	tc := simpleSliceValidatorSliceValidatorTestCase

	// Arrange
	expected := corestr.New.SimpleSlice.Direct(false, []string{"a"})
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}
	v.SetActual([]string{"a"})

	// Act
	sv := v.SliceValidator()

	actual := args.Map{
		"isNotNil": sv != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleSliceValidator_VerifyAll_FromSimpleSliceValidator(t *testing.T) {
	for caseIndex, tc := range simpleSliceValidatorVerifyAllTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		expectedLines := input["expected"].([]string)
		actualLines := input["actual"].([]string)

		expected := corestr.New.SimpleSlice.Direct(false, expectedLines)
		v := &corevalidator.SimpleSliceValidator{
			Expected:  expected,
			Condition: corevalidator.DefaultDisabledCoreCondition,
			CompareAs: stringcompareas.Equal,
		}
		v.SetActual(actualLines)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyAll(actualLines, params)

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_SimpleSliceValidator_VerifyFirst_FromSimpleSliceValidator(t *testing.T) {
	tc := simpleSliceValidatorVerifyFirstTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	expectedLines := input["expected"].([]string)
	actualLines := input["actual"].([]string)

	expected := corestr.New.SimpleSlice.Direct(false, expectedLines)
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}
	v.SetActual(actualLines)
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := v.VerifyFirst(actualLines, params)

	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleSliceValidator_VerifyUpto_FromSimpleSliceValidator(t *testing.T) {
	tc := simpleSliceValidatorVerifyUptoTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	expectedLines := input["expected"].([]string)
	actualLines := input["actual"].([]string)
	length, _ := input.GetAsInt("length")

	expected := corestr.New.SimpleSlice.Direct(false, expectedLines)
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}
	v.SetActual(actualLines)
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := v.VerifyUpto(actualLines, params, length)

	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
