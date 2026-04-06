package ostypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/ostype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Cov_Group_IsWindows(t *testing.T) {
	actual := args.Map{"result": ostype.WindowsGroup.IsWindows()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected windows", actual)
}

func Test_Cov_Group_IsUnix(t *testing.T) {
	actual := args.Map{"result": ostype.UnixGroup.IsUnix()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected unix", actual)
}

func Test_Cov_Group_IsAndroid(t *testing.T) {
	actual := args.Map{"result": ostype.AndroidGroup.IsAndroid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected android", actual)
}

func Test_Cov_Group_IsInvalidGroup(t *testing.T) {
	actual := args.Map{"result": ostype.InvalidGroup.IsInvalidGroup()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_Cov_Variation_Group_Android(t *testing.T) {
	g := ostype.Android.Group()
	actual := args.Map{"result": g.IsAndroid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected android group", actual)
}

func Test_Cov_Variation_Group_Unix(t *testing.T) {
	g := ostype.Linux.Group()
	actual := args.Map{"result": g.IsUnix()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected unix group", actual)
}

func Test_Cov_Variation_IsActualGroupUnix(t *testing.T) {
	actual := args.Map{"result": ostype.Linux.IsActualGroupUnix()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected actual group unix", actual)
}

func Test_Cov_Variation_IsPossibleUnixGroup(t *testing.T) {
	actual := args.Map{"result": ostype.Linux.IsPossibleUnixGroup()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected possible unix", actual)
	actual := args.Map{"result": ostype.Windows.IsPossibleUnixGroup()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "windows should not be unix", actual)
}
