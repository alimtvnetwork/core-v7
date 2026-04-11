package isanytests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/corecsv"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// Test_Extended_Defined_TypedNil verifies isany.Defined with typed-nil error and *int.
func Test_Extended_Defined_TypedNil(t *testing.T) {
	for caseIndex, testCase := range extendedDefinedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputs := input["inputs"].([]any)

		// Act
		actual := args.Map{}
		for i, v := range inputs {
			actual[fmt.Sprintf("result%d", i)] = fmt.Sprintf("%t", isany.Defined(v))
			actual[fmt.Sprintf("type%d", i)] = fmt.Sprintf("%T", v)
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Test_Extended_Null_TypedNil verifies isany.Null with typed-nil error and *int.
func Test_Extended_Null_TypedNil(t *testing.T) {
	for caseIndex, testCase := range extendedNullTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputs := input["inputs"].([]any)

		// Act
		actual := args.Map{}
		for i, v := range inputs {
			actual[fmt.Sprintf("result%d", i)] = fmt.Sprintf("%t", isany.Null(v))
			actual[fmt.Sprintf("type%d", i)] = fmt.Sprintf("%T", v)
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Test_Extended_DefinedBoth_TypedNil verifies isany.DefinedBoth with error and *int typed nils.
func Test_Extended_DefinedBoth_TypedNil(t *testing.T) {
	for caseIndex, testCase := range extendedDefinedBothTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairs := input["pairs"].([]args.TwoAny)

		// Act
		actual := args.Map{}
		for i, pair := range pairs {
			f := pair.First
			s := pair.Second
			actual[fmt.Sprintf("result%d", i)] = fmt.Sprintf("%t", isany.DefinedBoth(f, s))
			actual[fmt.Sprintf("types%d", i)] = corecsv.AnyToTypesCsvDefault(f, s)
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Test_Extended_NullBoth_TypedNil verifies isany.NullBoth with error and *int typed nils.
func Test_Extended_NullBoth_TypedNil(t *testing.T) {
	for caseIndex, testCase := range extendedNullBothTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairs := input["pairs"].([]args.TwoAny)

		// Act
		actual := args.Map{}
		for i, pair := range pairs {
			f := pair.First
			s := pair.Second
			actual[fmt.Sprintf("result%d", i)] = fmt.Sprintf("%t", isany.NullBoth(f, s))
			actual[fmt.Sprintf("types%d", i)] = corecsv.AnyToTypesCsvDefault(f, s)
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
