package mutexbykeytests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/mutexbykey"
)

// ============================================================================
// Get same key twice returns same mutex
// ============================================================================

func Test_Get_SameKeyTwice_Ext_Verification(t *testing.T) {
	for caseIndex, tc := range extGetSameKeyTwiceTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")

		// Act
		mutex1 := mutexbykey.Get(key)
		mutex2 := mutexbykey.Get(key)

		actual := args.Map{
			"isSame": mutex1 == mutex2,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)

		// Cleanup
		mutexbykey.Delete(key)
	}
}
