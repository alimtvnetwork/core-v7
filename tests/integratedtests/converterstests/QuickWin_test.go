package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_ToStringsUsingProcessor_NilInput(t *testing.T) {
	// Arrange
	result := converters.AnyTo.ToStringsUsingProcessor(
		false,
		func(index int, in any) (string, bool, bool) { return "", true, false },
		nil,
	)

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty slice for nil input", actual)
}
