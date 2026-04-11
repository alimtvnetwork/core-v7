package isanytests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── ReflectNull with reflect.Value input (line 15-16 branch) ──

func Test_ReflectNull_ReflectValueInput_Invalid(t *testing.T) {
	// Arrange
	rv := reflect.Value{} // invalid reflect.Value

	// Act
	actual := args.Map{"result": isany.ReflectNull(rv)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns error -- reflect.Value invalid", actual)
}

func Test_ReflectNull_ReflectValueInput_Valid(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)

	// Act
	actual := args.Map{"result": isany.ReflectNull(rv)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns non-empty -- reflect.Value valid int", actual)
}

func Test_ReflectNull_ReflectValueInput_NilPtr(t *testing.T) {
	// Arrange
	var ptr *int
	rv := reflect.ValueOf(ptr)

	// Act
	actual := args.Map{"result": isany.ReflectNull(rv)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns nil -- reflect.Value nil ptr", actual)
}

// ── JsonEqual error paths ──

func Test_JsonEqual_BothUnmarshalable(t *testing.T) {
	// Arrange
	// channels cannot be marshaled
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Act
	actual := args.Map{"result": isany.JsonEqual(ch1, ch2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonEqual returns correct value -- both unmarshalable", actual)
}

func Test_JsonEqual_OneUnmarshalable(t *testing.T) {
	// Arrange
	ch := make(chan int)

	// Act
	actual := args.Map{"result": isany.JsonEqual(ch, 42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonEqual returns correct value -- one unmarshalable", actual)
}

func Test_JsonEqual_OtherUnmarshalable(t *testing.T) {
	// Arrange
	ch := make(chan int)

	// Act
	actual := args.Map{"result": isany.JsonEqual(42, ch)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonEqual returns correct value -- other unmarshalable", actual)
}

// ── JsonMismatch with channels ──

func Test_JsonMismatch_BothUnmarshalable(t *testing.T) {
	// Arrange
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Act
	actual := args.Map{"result": isany.JsonMismatch(ch1, ch2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "JsonMismatch returns correct value -- both unmarshalable", actual)
}

// ── Null with non-nil chan ──

func Test_Null_NonNilChan(t *testing.T) {
	// Arrange
	ch := make(chan int)

	// Act
	actual := args.Map{"result": isany.Null(ch)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- non-nil chan", actual)
}

// ── ReflectValueNull invalid ──

func Test_ReflectValueNull_Invalid(t *testing.T) {
	// Arrange
	rv := reflect.Value{}

	// Act
	actual := args.Map{"result": isany.ReflectValueNull(rv)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull returns error -- invalid", actual)
}

// ── FuncOnly with func ──

func Test_FuncOnly_ValidFunc(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FuncOnly(func() {})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FuncOnly returns non-empty -- valid func", actual)
}

// ── Conclusive same nil values but non-interface ──

func Test_Conclusive_BothTypedNilDiffType(t *testing.T) {
	// Arrange
	var a *int
	var b *string
	isEq, isConcl := isany.Conclusive(a, b)

	// Act
	actual := args.Map{
		"isEqual": isEq,
		"isConclusive": isConcl,
	}

	// Assert
	expected := args.Map{
		"isEqual": true,
		"isConclusive": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns nil -- both typed nil diff type", actual)
}

// ── ReflectNull with nil directly ──

func Test_ReflectNull_NilDirect(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.ReflectNull(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns nil -- nil direct", actual)
}

// ── ReflectNotNull with value ──

func Test_ReflectNotNull_Value_FromReflectNullReflectVa(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.ReflectNotNull(42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNotNull returns correct value -- value", actual)
}

// ── PositiveIntegerType uint ──

func Test_PositiveIntegerType_Uint_FromReflectNullReflectVa(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerType(uint(42))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerType returns correct value -- uint", actual)
}

// ── FloatingPointType float64 ──

func Test_FloatingPointType_Float64_FromReflectNullReflectVa(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FloatingPointType(float64(3.14))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FloatingPointType returns correct value -- float64", actual)
}

func Test_FloatingPointType_Int_FromReflectNullReflectVa(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FloatingPointType(42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FloatingPointType returns correct value -- int", actual)
}
