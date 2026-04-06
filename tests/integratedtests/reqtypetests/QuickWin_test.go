package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_RangesNotMeet_EmptyReqs(t *testing.T) {
	// Arrange
	// Covers start() and end() with empty slice
	result := reqtype.RangesNotMeet("msg")

	// Act
	actual := args.Map{"result": result != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for no requests", actual)
}

func Test_QW_RangesString_EmptyReqs(t *testing.T) {
	// Arrange
	result := reqtype.RangesString(",")

	// Act
	actual := args.Map{"result": result != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}
