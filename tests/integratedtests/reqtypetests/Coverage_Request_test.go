package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Cov_Request_IsAnySkipOnExist(t *testing.T) {
	actual := args.Map{"result": reqtype.SkipOnExist.IsAnySkipOnExist()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Cov_Request_IsOverrideOrOverwriteOrEnforce(t *testing.T) {
	actual := args.Map{"result": reqtype.Overwrite.IsOverrideOrOverwriteOrEnforce()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": reqtype.Override.IsOverrideOrOverwriteOrEnforce()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": reqtype.Enforce.IsOverrideOrOverwriteOrEnforce()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Cov_Request_IsRestartOrReload(t *testing.T) {
	actual := args.Map{"result": reqtype.Restart.IsRestartOrReload()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": reqtype.Reload.IsRestartOrReload()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}
