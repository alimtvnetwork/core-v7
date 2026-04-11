package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ConvEnumAnyValToInteger ──

func Test_ConvEnumAnyValToInteger_Int(t *testing.T) {
	// Arrange
	val := enumimpl.ConvEnumAnyValToInteger(42)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns 42 -- int input", actual)
}

func Test_ConvEnumAnyValToInteger_String(t *testing.T) {
	// Arrange
	val := enumimpl.ConvEnumAnyValToInteger("notAnInt")

	// Act
	actual := args.Map{"isMinInt": val < 0}

	// Assert
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns MinInt -- string input", actual)
}

// ── NameWithValue (function) ──

func Test_NameWithValue(t *testing.T) {
	// Arrange
	result := enumimpl.NameWithValue(10)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameWithValue returns formatted string -- int input", actual)
}

// ── UnsupportedNames ──

func Test_UnsupportedNames(t *testing.T) {
	// Arrange
	allNames := []string{"A", "B", "C", "D"}
	result := enumimpl.UnsupportedNames(allNames, "A", "B")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UnsupportedNames returns 2 -- two unsupported", actual)
}

func Test_UnsupportedNames_AllSupported(t *testing.T) {
	// Arrange
	allNames := []string{"A", "B"}
	result := enumimpl.UnsupportedNames(allNames, "A", "B")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "UnsupportedNames returns 0 -- all supported", actual)
}

// ── KeyAnyVal ──

func Test_KeyAnyVal(t *testing.T) {
	// Arrange
	kv := enumimpl.KeyAnyVal{Key: "test", AnyValue: 42}

	// Act
	actual := args.Map{
		"key":      kv.Key,
		"valInt":   kv.ValInt(),
		"isString": kv.IsString(),
	}

	// Assert
	expected := args.Map{
		"key": "test",
		"valInt": 42,
		"isString": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal returns correct fields -- int value", actual)
}

func Test_KeyAnyVal_StringValue(t *testing.T) {
	// Arrange
	kv := enumimpl.KeyAnyVal{Key: "strKey", AnyValue: "hello"}

	// Act
	actual := args.Map{
		"key":      kv.Key,
		"isString": kv.IsString(),
		"anyVal":   kv.AnyValString(),
	}

	// Assert
	expected := args.Map{
		"key": "strKey",
		"isString": true,
		"anyVal": "hello",
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal returns string type -- string value", actual)
}

// ── DiffLeftRight ──

func Test_DiffLeftRight_Same(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "same", Right: "same"}

	// Act
	actual := args.Map{
		"isSame":      dlr.IsSame(),
		"isNotEqual":  dlr.IsNotEqual(),
		"isEqual":     dlr.IsEqual(false),
		"diffStr":     dlr.DiffString(),
	}

	// Assert
	expected := args.Map{
		"isSame": true, "isNotEqual": false, "isEqual": true, "diffStr": "",
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight returns same -- equal values", actual)
}

func Test_DiffLeftRight_Different(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "left", Right: "right"}

	// Act
	actual := args.Map{
		"isSame":     dlr.IsSame(),
		"isNotEqual": dlr.IsNotEqual(),
		"hasMismatch": dlr.HasMismatch(false),
	}

	// Assert
	expected := args.Map{
		"isSame": false, "isNotEqual": true, "hasMismatch": true,
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight returns mismatch -- different values", actual)
}

func Test_DiffLeftRight_RegardlessOfType(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: 42, Right: 42}

	// Act
	actual := args.Map{
		"isEqualRegardless": dlr.IsEqual(true),
		"isSameTypeSame":    dlr.IsSameTypeSame(),
	}

	// Assert
	expected := args.Map{
		"isEqualRegardless": true, "isSameTypeSame": true,
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight returns equal -- regardless of type", actual)
}

func Test_DiffLeftRight_JsonString(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{"notEmpty": dlr.JsonString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight returns json -- serialized", actual)
}

func Test_DiffLeftRight_SpecificFullString(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "x", Right: "y"}
	l, r := dlr.SpecificFullString()

	// Act
	actual := args.Map{
		"leftNotEmpty":  l != "",
		"rightNotEmpty": r != "",
	}

	// Assert
	expected := args.Map{
		"leftNotEmpty": true,
		"rightNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight returns full strings -- both sides", actual)
}

// ── DefaultDiffCheckerImpl ──

func Test_DefaultDiffChecker_IsEqual(t *testing.T) {
	// Arrange
	checker := enumimpl.DefaultDiffCheckerImpl
	result := checker.IsEqual(false, 42, 42)

	// Act
	actual := args.Map{"isEqual": result}

	// Assert
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "DefaultDiffChecker returns true -- equal values", actual)
}

func Test_DefaultDiffChecker_IsEqual_Regardless(t *testing.T) {
	// Arrange
	checker := enumimpl.DefaultDiffCheckerImpl
	result := checker.IsEqual(true, 42, 42)

	// Act
	actual := args.Map{"isEqual": result}

	// Assert
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "DefaultDiffChecker returns true -- regardless mode", actual)
}

func Test_LeftRightDiffChecker_IsEqual_FromConvEnumAnyValToInte(t *testing.T) {
	// Arrange
	checker := enumimpl.LeftRightDiffCheckerImpl
	result := checker.IsEqual(false, "a", "a")

	// Act
	actual := args.Map{"isEqual": result}

	// Assert
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "LeftRightDiffChecker returns true -- equal strings", actual)
}
