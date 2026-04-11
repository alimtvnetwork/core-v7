package ostypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/ostype"
)

// Test_Cov4_Group_UnmarshallEnumToValue covers
// ostype/Group.go L191-193: UnmarshallEnumToValue.
func Test_Group_UnmarshallEnumToValue(t *testing.T) {
	// Arrange
	g := ostype.UnixGroup
	validBytes := []byte("1")

	// Act
	val, err := g.UnmarshallEnumToValue(validBytes)

	// Assert
	actual := args.Map{
		"value": val,
		"hasError": err != nil,
	}
	expected := args.Map{
		"value": byte(ostype.UnixGroup),
		"hasError": false,
	}
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- UnmarshallEnumToValue", actual)
}
