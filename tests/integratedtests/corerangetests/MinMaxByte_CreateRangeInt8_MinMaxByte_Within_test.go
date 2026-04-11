package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_MinMaxByte_CreateRangeInt8_FromMinMaxByteCreateRang(t *testing.T) {
	// Arrange
	mmb := &corerange.MinMaxByte{Min: 0, Max: 10}
	r := mmb.CreateRangeInt8("0-10", "-")

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_MinMaxByte_CreateRangeInt16_FromMinMaxByteCreateRang(t *testing.T) {
	// Arrange
	mmb := &corerange.MinMaxByte{Min: 0, Max: 10}
	r := mmb.CreateRangeInt16("0-10", "-")

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Within_StringRangeInt32_FromMinMaxByteCreateRang(t *testing.T) {
	// Arrange
	val, ok := corerange.Within.StringRangeInt32("100")

	// Act
	actual := args.Map{"result": ok || val != 100}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)
}

func Test_Within_StringRangeInt16_FromMinMaxByteCreateRang(t *testing.T) {
	// Arrange
	val, ok := corerange.Within.StringRangeInt16("100")

	// Act
	actual := args.Map{"result": ok || val != 100}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)
}

func Test_Within_StringRangeInt8_FromMinMaxByteCreateRang(t *testing.T) {
	// Arrange
	val, ok := corerange.Within.StringRangeInt8("50")

	// Act
	actual := args.Map{"result": ok || val != 50}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 50", actual)
}

func Test_Within_StringRangeByte_FromMinMaxByteCreateRang(t *testing.T) {
	// Arrange
	val, ok := corerange.Within.StringRangeByte("200")

	// Act
	actual := args.Map{"result": ok || val != 200}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 200", actual)
}

func Test_Within_StringRangeUint16(t *testing.T) {
	// Arrange
	val, ok := corerange.Within.StringRangeUint16("1000")

	// Act
	actual := args.Map{"result": ok || val != 1000}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 1000", actual)
}

func Test_Within_StringRangeUint32(t *testing.T) {
	// Arrange
	val, ok := corerange.Within.StringRangeUint32("1000")

	// Act
	actual := args.Map{"result": ok || val != 1000}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 1000", actual)
}

func Test_Within_StringRangeIntegerDefault(t *testing.T) {
	// Arrange
	val, ok := corerange.Within.StringRangeIntegerDefault(0, 100, "50")

	// Act
	actual := args.Map{"result": ok || val != 50}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 50", actual)
	// below min
	val2, ok2 := corerange.Within.StringRangeIntegerDefault(0, 100, "-5")
	actual = args.Map{"result": ok2 || val2 != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for below min", actual)
	// above max
	val3, ok3 := corerange.Within.StringRangeIntegerDefault(0, 100, "200")
	actual = args.Map{"result": ok3 || val3 != 100}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 100 for above max", actual)
}

func Test_Within_StringRangeFloat(t *testing.T) {
	// Arrange
	val, ok := corerange.Within.StringRangeFloat(true, 0, 100, "50.5")

	// Act
	actual := args.Map{"result": ok || val != 50.5}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 50.5", actual)
}

func Test_Within_StringRangeFloatDefault(t *testing.T) {
	// Arrange
	_, ok := corerange.Within.StringRangeFloatDefault("50.5")

	// Act
	actual := args.Map{"result": ok}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected in range", actual)
}

func Test_Within_StringRangeFloat64(t *testing.T) {
	// Arrange
	val, ok := corerange.Within.StringRangeFloat64(true, 0, 100, "50.5")

	// Act
	actual := args.Map{"result": ok || val != 50.5}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 50.5", actual)
}

func Test_Within_StringRangeFloat64Default(t *testing.T) {
	// Arrange
	_, ok := corerange.Within.StringRangeFloat64Default("50.5")

	// Act
	actual := args.Map{"result": ok}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected in range", actual)
}

func Test_Within_RangeByteDefault(t *testing.T) {
	// Arrange
	val, ok := corerange.Within.RangeByteDefault(100)

	// Act
	actual := args.Map{"result": ok || val != 100}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)
}

func Test_Within_RangeUint16Default(t *testing.T) {
	// Arrange
	val, ok := corerange.Within.RangeUint16Default(1000)

	// Act
	actual := args.Map{"result": ok || val != 1000}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 1000", actual)
}

func Test_Within_RangeFloat(t *testing.T) {
	// Arrange
	val, ok := corerange.Within.RangeFloat(true, 0, 100, 50)

	// Act
	actual := args.Map{"result": ok || val != 50}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 50", actual)
	// below min with boundary
	val2, ok2 := corerange.Within.RangeFloat(true, 10, 100, 5)
	actual = args.Map{"result": ok2 || val2 != 10}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)
	// above max with boundary
	val3, ok3 := corerange.Within.RangeFloat(true, 0, 100, 200)
	actual = args.Map{"result": ok3 || val3 != 100}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)
	// no boundary
	val4, ok4 := corerange.Within.RangeFloat(false, 0, 100, 200)
	actual = args.Map{"result": ok4 || val4 != 200}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 200", actual)
}
