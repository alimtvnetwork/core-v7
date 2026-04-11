package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
)

// === anyItemConverter (converters.AnyTo) uncovered branches ===

func Test_AnyTo_ToStringsUsingProcessor_Break(t *testing.T) {
	// Arrange
	result := converters.AnyTo.ToStringsUsingProcessor(
		false,
		func(index int, in any) (string, bool, bool) {
			return "x", true, true // take + break
		},
		[]string{"a", "b"},
	)

	// Act
	actual := args.Map{"result": len(result) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_AnyTo_ToStringsUsingSimpleProcessor_Empty(t *testing.T) {
	// Arrange
	result := converters.AnyTo.ToStringsUsingSimpleProcessor(
		false,
		func(index int, in any) string { return "x" },
		[]string{},
	)

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_AnyTo_ToPrettyJson_Error(t *testing.T) {
	// Arrange
	// channels can't be marshaled
	ch := make(chan int)
	result := converters.AnyTo.ToPrettyJson(ch)

	// Act
	actual := args.Map{"result": result != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for unmarshalable", actual)
}

func Test_AnyTo_Bytes_Error(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic for unmarshalable", actual)
	}()
	ch := make(chan int)
	converters.AnyTo.Bytes(ch)
}

// === stringTo uncovered branches ===

func Test_StringTo_Float64Must_Panic(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	converters.StringTo.Float64Must("notanumber")
}

// === stringsTo uncovered branches ===

func Test_StringsTo_IntegersOptionPanic_Panic(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	converters.StringsTo.IntegersOptionPanic(true, "not_int")
}

func Test_StringsTo_BytesConditional_Break(t *testing.T) {
	// Arrange
	result := converters.StringsTo.BytesConditional(
		func(in string) (byte, bool, bool) {
			return 0, true, true
		},
		[]string{"1", "2"},
	)

	// Act
	actual := args.Map{"result": len(result) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_StringsTo_BytesMust_Panic(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	converters.StringsTo.BytesMust("not_byte")
}

func Test_StringsTo_Float64sMust_Panic(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	converters.StringsTo.Float64sMust("not_float")
}

func Test_StringsTo_Float64sConditional_Break(t *testing.T) {
	// Arrange
	result := converters.StringsTo.Float64sConditional(
		func(in string) (float64, bool, bool) {
			return 0, true, true
		},
		[]string{"1.0", "2.0"},
	)

	// Act
	actual := args.Map{"result": len(result) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}
