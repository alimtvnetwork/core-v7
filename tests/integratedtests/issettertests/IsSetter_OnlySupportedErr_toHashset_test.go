package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/issetter"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_IsSetter_OnlySupportedErr_ExercisesToHashset(t *testing.T) {
	// Arrange
	// OnlySupportedErr internally calls toHashset
	v := issetter.True
	err := v.OnlySupportedErr("True", "False")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for unsupported names", actual)
}

func Test_IsSetter_OnlySupportedErr_AllSupported(t *testing.T) {
	// Arrange
	v := issetter.True
	err := v.OnlySupportedErr("Uninitialized", "True", "False", "Unset", "Set", "Wildcard")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_IsSetter_OnlySupportedErr_Empty(t *testing.T) {
	// Arrange
	v := issetter.True
	err := v.OnlySupportedErr()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}
