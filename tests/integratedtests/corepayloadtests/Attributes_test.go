package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Attributes.IsEqual — Regression: logic inversion bug in IsSafeValid/HasIssuesOrEmpty
// =============================================================================

func Test_Attributes_IsEqual_Verification(t *testing.T) {
	for caseIndex, testCase := range attributesIsEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftNil := input.GetAsBoolDefault("left_nil", false)
		rightNil := input.GetAsBoolDefault("right_nil", false)
		samePointer := input.GetAsBoolDefault("same_pointer", false)

		var left, right *corepayload.Attributes

		if !leftNil {
			leftPayload, _ := input.GetAsString("left_payload")
			payload, _ := input.GetAsString("payload")

			if leftPayload == "" {
				leftPayload = payload
			}

			left = &corepayload.Attributes{
				DynamicPayloads: []byte(leftPayload),
			}
		}

		if samePointer {
			right = left
		} else if !rightNil {
			rightPayload, _ := input.GetAsString("right_payload")
			payload, _ := input.GetAsString("payload")

			if rightPayload == "" {
				rightPayload = payload
			}

			right = &corepayload.Attributes{
				DynamicPayloads: []byte(rightPayload),
			}
		}

		// Act
		result := left.IsEqual(right)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"isEqual": result,
		})
	}
}
func Test_Attributes_IsSafeValid_Verification(t *testing.T) {
	for caseIndex, testCase := range attributesIsSafeValidTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilAttr := input.GetAsBoolDefault("nil_attr", false)
		empty := input.GetAsBoolDefault("empty", false)

		var attr *corepayload.Attributes

		if !nilAttr && !empty {
			payload, _ := input.GetAsString("payload")
			attr = &corepayload.Attributes{
				DynamicPayloads: []byte(payload),
			}
		} else if !nilAttr && empty {
			attr = &corepayload.Attributes{}
		}

		// Act
		result := attr.IsSafeValid()

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"isSafeValid": result,
		})
	}
}