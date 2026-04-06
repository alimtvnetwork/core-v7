package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_RangesNotMeet_EmptyReqs(t *testing.T) {
	// Covers start() and end() with empty slice
	result := reqtype.RangesNotMeet("msg")
	actual := args.Map{"result": result != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for no requests", actual)
}

func Test_QW_RangesString_EmptyReqs(t *testing.T) {
	result := reqtype.RangesString(",")
	actual := args.Map{"result": result != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}
