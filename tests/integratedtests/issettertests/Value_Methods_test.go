package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/issetter"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Value_Methods_Ext2(t *testing.T) {
	// Arrange
	v, err := issetter.New("Set")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)

	// Assert
	actual = args.Map{"result": v.IsUnset()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be unset", actual)
	actual = args.Map{"result": v.IsSet()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be set", actual)
}

func Test_NewBool_Ext2(t *testing.T) {
	// Arrange
	v := issetter.NewBool(true)

	// Assert
	actual := args.Map{"result": v.IsUnset()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be unset", actual)
	actual = args.Map{"result": v.Boolean()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be true", actual)
}

func Test_NewMust_Ext2(t *testing.T) {
	// Act
	v := issetter.NewMust("True")

	// Assert
	actual := args.Map{"result": v.IsUnset()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be unset", actual)
	actual = args.Map{"result": v.IsTrue()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be True", actual)
}

func Test_Max_Ext2(t *testing.T) {
	// Act
	result := issetter.Max()

	// Assert
	actual := args.Map{"result": result != issetter.Wildcard}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Wildcard to be max", actual)
}

func Test_Min_Ext2(t *testing.T) {
	// Act
	result := issetter.Min()

	// Assert
	actual := args.Map{"result": result != issetter.Uninitialized}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Uninitialized to be min", actual)
}
