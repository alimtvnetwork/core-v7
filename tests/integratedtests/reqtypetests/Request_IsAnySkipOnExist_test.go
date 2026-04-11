package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Request_IsAnySkipOnExist(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.SkipOnExist.IsAnySkipOnExist()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Request_IsOverrideOrOverwriteOrEnforce(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Overwrite.IsOverrideOrOverwriteOrEnforce()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": reqtype.Override.IsOverrideOrOverwriteOrEnforce()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": reqtype.Enforce.IsOverrideOrOverwriteOrEnforce()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Request_IsRestartOrReload(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Restart.IsRestartOrReload()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": reqtype.Reload.IsRestartOrReload()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}
