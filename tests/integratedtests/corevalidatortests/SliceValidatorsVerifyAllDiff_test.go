package corevalidatortests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

var sliceValidatorsVerifyAllDiffTestCases = []coretestcases.CaseV1{
	{
		Title: "VerifyAll error embeds diff output for single mismatch",
		ArrangeInput: args.Map{
			"when":     "given actual != expected with one validator",
			"actual":   []string{"hello", "world"},
			"expected": []string{"hello", "planet"},
			"header":   "diff-test-header",
		},
		ExpectedInput: []string{
			"true", // error is not nil
			"true", // error contains header
			"true", // error contains actual line "world"
			"true", // error contains expected line "planet"
		},
	},
	{
		Title: "VerifyAll error embeds diff for multiple validators with one failing",
		ArrangeInput: args.Map{
			"when":      "given two validators, second mismatches",
			"actual1":   []string{"a"},
			"expected1": []string{"a"},
			"actual2":   []string{"x"},
			"expected2": []string{"y"},
			"header":    "multi-validator-diff",
		},
		ExpectedInput: []string{
			"true",  // error is not nil
			"true",  // error contains header
			"false", // error does NOT contain "a" mismatch (it matched)
			"true",  // error contains "x" (actual)
			"true",  // error contains "y" (expected)
		},
	},
	{
		Title: "VerifyAll returns nil when all validators match",
		ArrangeInput: args.Map{
			"when":     "given actual == expected",
			"actual":   []string{"match"},
			"expected": []string{"match"},
			"header":   "no-diff-header",
		},
		ExpectedInput: []string{
			"false", // error is nil
		},
	},
}

func Test_SliceValidators_VerifyAll_DiffOutput_Verification(t *testing.T) {
	for caseIndex, testCase := range sliceValidatorsVerifyAllDiffTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		header, _ := input.GetAsString("header")

		var actLines []string

		// Check if multi-validator case
		_, hasActual2 := input.Get("actual2")

		if hasActual2 {
			actual1Val, _ := input.Get("actual1")
			expected1Val, _ := input.Get("expected1")
			actual2Val, _ := input.Get("actual2")
			expected2Val, _ := input.Get("expected2")

			v := &corevalidator.SliceValidators{
				Validators: []corevalidator.SliceValidator{
					{
						Condition:     corevalidator.DefaultDisabledCoreCondition,
						CompareAs:     stringcompareas.Equal,
						ActualLines:   actual1Val.([]string),
						ExpectedLines: expected1Val.([]string),
					},
					{
						Condition:     corevalidator.DefaultDisabledCoreCondition,
						CompareAs:     stringcompareas.Equal,
						ActualLines:   actual2Val.([]string),
						ExpectedLines: expected2Val.([]string),
					},
				},
			}

			params := &corevalidator.Parameter{
				CaseIndex:       caseIndex,
				Header:          header,
				IsCaseSensitive: true,
			}

			// Act
			err := v.VerifyAll(header, params, false)

			errStr := ""
			if err != nil {
				errStr = err.Error()
			}

			actLines = []string{
				fmt.Sprintf("%v", err != nil),
				fmt.Sprintf("%v", strings.Contains(errStr, header)),
				fmt.Sprintf("%v", strings.Contains(errStr, "Actual-Processed: `\"a\"`")),
				fmt.Sprintf("%v", strings.Contains(errStr, "x")),
				fmt.Sprintf("%v", strings.Contains(errStr, "y")),
			}
		} else {
			actualVal, _ := input.Get("actual")
			expectedVal, _ := input.Get("expected")

			v := &corevalidator.SliceValidators{
				Validators: []corevalidator.SliceValidator{
					{
						Condition:     corevalidator.DefaultDisabledCoreCondition,
						CompareAs:     stringcompareas.Equal,
						ActualLines:   actualVal.([]string),
						ExpectedLines: expectedVal.([]string),
					},
				},
			}

			params := &corevalidator.Parameter{
				CaseIndex:       caseIndex,
				Header:          header,
				IsCaseSensitive: true,
			}

			// Act
			err = v.VerifyAll(header, params, false)

			if err == nil {
				actLines = []string{"false"}
			} else {
				errStr := err.Error()
				actLines = []string{
					fmt.Sprintf("%v", err != nil),
					fmt.Sprintf("%v", strings.Contains(errStr, header)),
					fmt.Sprintf("%v", strings.Contains(errStr, "world")),
					fmt.Sprintf("%v", strings.Contains(errStr, "planet")),
				}
			}
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
