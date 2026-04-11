package iserrortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/iserror"
)

// ============================================================================
// ExitError
// ============================================================================

func Test_ExitError_Ext_Verification(t *testing.T) {
	for caseIndex, tc := range extExitErrorTestCases {
		// Arrange
		err := errors.New("not an exec exit error")

		// Act
		actual := args.Map{
			"isExitError": iserror.ExitError(err),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
