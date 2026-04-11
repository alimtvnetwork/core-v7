package coreconvertedtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/converters/coreconverted"
)

func Test_Bytes_HandleWithPanic_NoError(t *testing.T) {
	// Arrange
	b := &coreconverted.Bytes{Values: []byte{1}, CombinedError: nil}
	// Should not panic
	b.HandleWithPanic()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleWithPanic_NoError panics -- with args", actual)
}

func Test_Bytes_HandleWithPanic_WithError(t *testing.T) {
	// Arrange
	b := &coreconverted.Bytes{Values: nil, CombinedError: errors.New("err")}
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		b.HandleWithPanic()
	}()

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleWithPanic_WithError panics -- with args", actual)
}

func Test_Integers_HandleWithPanic_NoError(t *testing.T) {
	// Arrange
	i := &coreconverted.Integers{Values: []int{1}, CombinedError: nil}
	i.HandleWithPanic()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Integers_HandleWithPanic_NoError panics -- with args", actual)
}

func Test_Integers_HandleWithPanic_WithError(t *testing.T) {
	// Arrange
	i := &coreconverted.Integers{Values: nil, CombinedError: errors.New("err")}
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		i.HandleWithPanic()
	}()

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Integers_HandleWithPanic_WithError panics -- with args", actual)
}

func Test_Bytes_HasAnyItem_Empty(t *testing.T) {
	// Arrange
	b := &coreconverted.Bytes{Values: []byte{}}

	// Act
	actual := args.Map{
		"hasAny": b.HasAnyItem(),
		"isEmpty": b.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"hasAny": false,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Bytes_HasAnyItem_Empty returns empty -- with args", actual)
}

func Test_Integers_HasAnyItem_Empty(t *testing.T) {
	// Arrange
	i := &coreconverted.Integers{Values: []int{}}

	// Act
	actual := args.Map{
		"hasAny": i.HasAnyItem(),
		"isEmpty": i.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"hasAny": false,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Integers_HasAnyItem_Empty returns empty -- with args", actual)
}
