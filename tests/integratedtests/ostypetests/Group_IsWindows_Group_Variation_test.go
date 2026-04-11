package ostypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/ostype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Group_IsWindows(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.WindowsGroup.IsWindows()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected windows", actual)
}

func Test_Group_IsUnix(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.UnixGroup.IsUnix()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected unix", actual)
}

func Test_Group_IsAndroid(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.AndroidGroup.IsAndroid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected android", actual)
}

func Test_Group_IsInvalidGroup(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.InvalidGroup.IsInvalidGroup()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_Variation_Group_Android(t *testing.T) {
	// Arrange
	g := ostype.Android.Group()

	// Act
	actual := args.Map{"result": g.IsAndroid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected android group", actual)
}

func Test_Variation_Group_Unix(t *testing.T) {
	// Arrange
	g := ostype.Linux.Group()

	// Act
	actual := args.Map{"result": g.IsUnix()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected unix group", actual)
}

func Test_Variation_IsActualGroupUnix(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.Linux.IsActualGroupUnix()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected actual group unix", actual)
}

func Test_Variation_IsPossibleUnixGroup(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.Linux.IsPossibleUnixGroup()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected possible unix", actual)
	actual = args.Map{"result": ostype.Windows.IsPossibleUnixGroup()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "windows should not be unix", actual)
}
