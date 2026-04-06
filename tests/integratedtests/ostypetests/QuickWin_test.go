package ostypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/ostype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_Group_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	// Arrange
	g := ostype.UnixGroup

	// Act
	actual := args.Map{"result": g.IsAnyEnumsEqual()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for empty enums", actual)
}

func Test_QW_Group_MinByte(t *testing.T) {
	_ = ostype.UnixGroup.MinByte()
}

func Test_QW_Variation_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	// Arrange
	v := ostype.Linux

	// Act
	actual := args.Map{"result": v.IsAnyEnumsEqual()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for empty enums", actual)
}

func Test_QW_Variation_MinByte(t *testing.T) {
	_ = ostype.Linux.MinByte()
}
