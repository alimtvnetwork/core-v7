package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BoolOnce ──

func Test_BoolOnce_Value(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOnce(func() bool { return true })

	// Act
	actual := args.Map{
		"value":   bo.Value(),
		"execute": bo.Execute(),
		"string":  bo.String(),
	}

	// Assert
	expected := args.Map{
		"value":   true,
		"execute": true,
		"string":  "true",
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce returns true -- true func", actual)
}

func Test_BoolOnce_False(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOnce(func() bool { return false })

	// Act
	actual := args.Map{
		"value":  bo.Value(),
		"string": bo.String(),
	}

	// Assert
	expected := args.Map{
		"value":  false,
		"string": "false",
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce returns false -- false func", actual)
}

func Test_BoolOnce_Serialize(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOnce(func() bool { return true })
	b, err := bo.Serialize()

	// Act
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce.Serialize succeeds -- true", actual)
}

func Test_BoolOnce_MarshalJSON(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOnce(func() bool { return true })
	b, err := bo.MarshalJSON()

	// Act
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce.MarshalJSON succeeds -- true", actual)
}

func Test_BoolOnce_UnmarshalJSON(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOncePtr(func() bool { return false })
	err := bo.UnmarshalJSON([]byte("true"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "BoolOnce.UnmarshalJSON succeeds -- true bytes", actual)
}

func Test_BoolOnce_Ptr(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOncePtr(func() bool { return true })

	// Act
	actual := args.Map{
		"notNil": bo != nil,
		"value":  bo.Value(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"value":  true,
	}
	expected.ShouldBeEqual(t, 0, "NewBoolOncePtr returns non-nil -- true", actual)
}
