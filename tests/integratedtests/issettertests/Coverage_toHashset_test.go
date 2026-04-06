package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/issetter"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Cov_IsSetter_OnlySupportedErr_ExercisesToHashset(t *testing.T) {
	// OnlySupportedErr internally calls toHashset
	v := issetter.True
	err := v.OnlySupportedErr("True", "False")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for unsupported names", actual)
}

func Test_Cov_IsSetter_OnlySupportedErr_AllSupported(t *testing.T) {
	v := issetter.True
	err := v.OnlySupportedErr("Uninitialized", "True", "False", "Unset", "Set", "Wildcard")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Cov_IsSetter_OnlySupportedErr_Empty(t *testing.T) {
	v := issetter.True
	err := v.OnlySupportedErr()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}
