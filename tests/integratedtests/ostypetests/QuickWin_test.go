package ostypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/ostype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_Group_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	g := ostype.UnixGroup
	actual := args.Map{"result": g.IsAnyEnumsEqual()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for empty enums", actual)
}

func Test_QW_Group_MinByte(t *testing.T) {
	_ = ostype.UnixGroup.MinByte()
}

func Test_QW_Variation_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	v := ostype.Linux
	actual := args.Map{"result": v.IsAnyEnumsEqual()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for empty enums", actual)
}

func Test_QW_Variation_MinByte(t *testing.T) {
	_ = ostype.Linux.MinByte()
}
