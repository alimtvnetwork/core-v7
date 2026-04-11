package corejsontests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Result_IsEqual_FromResultIsEqual(t *testing.T) {
	for caseIndex, tc := range resultIsEqualTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		a := input["a"].(corejson.Result)
		b := input["b"].(corejson.Result)

		// Act
		result := fmt.Sprintf("%v", a.IsEqual(b))

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_Result_IsEqualPtr_FromResultIsEqual(t *testing.T) {
	for caseIndex, tc := range resultIsEqualPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		a := input["aPtr"].(*corejson.Result)
		b := input["bPtr"].(*corejson.Result)

		// Act
		result := fmt.Sprintf("%v", a.IsEqualPtr(b))

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}
