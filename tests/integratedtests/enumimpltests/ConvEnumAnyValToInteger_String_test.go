package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ConvEnumAnyValToInteger ──

func Test_ConvEnumAnyValToInteger_String_FromConvEnumAnyValToInte(t *testing.T) {
	// Arrange
	result := enumimpl.ConvEnumAnyValToInteger("hello")

	// Act
	actual := args.Map{"isMinInt": result < 0}

	// Assert
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvToInt_String returns correct value -- with args", actual)
}

func Test_ConvEnumAnyValToInteger_Int_FromConvEnumAnyValToInte(t *testing.T) {
	// Arrange
	result := enumimpl.ConvEnumAnyValToInteger(42)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ConvToInt_Int returns correct value -- with args", actual)
}

func Test_ConvEnumAnyValToInteger_Fallback(t *testing.T) {
	// Arrange
	result := enumimpl.ConvEnumAnyValToInteger(3.14)

	// Act
	actual := args.Map{"isInt": result != 0}

	// Assert
	expected := args.Map{"isInt": true}
	expected.ShouldBeEqual(t, 0, "ConvToInt_Fallback returns correct value -- with args", actual)
}

// ── PrependJoin / JoinPrependUsingDot ──

func Test_PrependJoin(t *testing.T) {
	// Arrange
	result := enumimpl.PrependJoin(".", "prefix", "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrependJoin returns correct value -- with args", actual)
}

func Test_JoinPrependUsingDot_FromConvEnumAnyValToInte(t *testing.T) {
	// Arrange
	result := enumimpl.JoinPrependUsingDot("prefix", "a")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinPrependUsingDot returns correct value -- with args", actual)
}

// ── KeyAnyVal ──

func Test_KeyAnyVal_Methods_FromConvEnumAnyValToInte(t *testing.T) {
	// Arrange
	kv := enumimpl.KeyAnyVal{Key: "name", AnyValue: 42}
	kvStr := enumimpl.KeyAnyVal{Key: "name", AnyValue: "hello"}

	// Act
	actual := args.Map{
		"key":         kv.KeyString(),
		"anyVal":      kv.AnyVal() != nil,
		"anyValStr":   kv.AnyValString() != "",
		"wrapKey":     kv.WrapKey() != "",
		"wrapVal":     kv.WrapValue() != "",
		"isString":    kv.IsString(),
		"valInt":      kv.ValInt(),
		"string":      kv.String() != "",
		"strIsString": kvStr.IsString(),
		"strString":   kvStr.String() != "",
	}

	// Assert
	expected := args.Map{
		"key":         "name",
		"anyVal":      true,
		"anyValStr":   true,
		"wrapKey":     true,
		"wrapVal":     true,
		"isString":    false,
		"valInt":      42,
		"string":      true,
		"strIsString": true,
		"strString":   true,
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal returns correct value -- with args", actual)
}

func Test_KeyAnyVal_KeyValInteger_FromConvEnumAnyValToInte(t *testing.T) {
	// Arrange
	kv := enumimpl.KeyAnyVal{Key: "test", AnyValue: 5}
	kvi := kv.KeyValInteger()

	// Act
	actual := args.Map{
		"key": kvi.Key,
		"val": kvi.ValueInteger,
	}

	// Assert
	expected := args.Map{
		"key": "test",
		"val": 5,
	}
	expected.ShouldBeEqual(t, 0, "KeyValInteger_conv returns correct value -- with args", actual)
}

// ── KeyValInteger ──

func Test_KeyValInteger_Methods_FromConvEnumAnyValToInte(t *testing.T) {
	// Arrange
	kvi := enumimpl.KeyValInteger{Key: "test", ValueInteger: 5}
	kviStr := enumimpl.KeyValInteger{Key: "test", ValueInteger: -9223372036854775808}

	// Act
	actual := args.Map{
		"wrapKey":   kvi.WrapKey() != "",
		"wrapVal":   kvi.WrapValue() != "",
		"isString":  kvi.IsString(),
		"string":    kvi.String() != "",
		"anyKey":    kvi.KeyAnyVal().Key,
		"strIsStr":  kviStr.IsString(),
		"strString": kviStr.String() != "",
	}

	// Assert
	expected := args.Map{
		"wrapKey":   true,
		"wrapVal":   true,
		"isString":  false,
		"string":    true,
		"anyKey":    "test",
		"strIsStr":  true,
		"strString": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValInteger returns correct value -- with args", actual)
}

// ── AllNameValues ──

func Test_AllNameValues(t *testing.T) {
	// Arrange
	names := []string{"Invalid", "Active"}
	values := []int{0, 1}
	result := enumimpl.AllNameValues(names, values)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllNameValues returns non-empty -- with args", actual)
}

// ── KeyAnyValues ──

func Test_KeyAnyValues(t *testing.T) {
	// Arrange
	result := enumimpl.KeyAnyValues([]string{"a", "b"}, []int{1, 2})
	empty := enumimpl.KeyAnyValues([]string{}, []int{})

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(empty),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns non-empty -- with args", actual)
}

// ── IntegersRangesOfAnyVal ──

func Test_IntegersRangesOfAnyVal(t *testing.T) {
	// Arrange
	result := enumimpl.IntegersRangesOfAnyVal([]int{3, 1, 2})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": 1,
		"last": 3,
	}
	expected.ShouldBeEqual(t, 0, "IntegersRangesOfAnyVal returns correct value -- with args", actual)
}

// ── DynamicMap extra coverage ──

func Test_DynamicMap_IsValueString(t *testing.T) {
	// Arrange
	dmStr := &enumimpl.DynamicMap{"a": "hello"}
	dmInt := &enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{
		"strIsStr": dmStr.IsValueString(),
		"intIsStr": dmInt.IsValueString(),
	}

	// Assert
	expected := args.Map{
		"strIsStr": true,
		"intIsStr": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_IsValueString returns non-empty -- with args", actual)
}

func Test_DynamicMap_SortedKeyValues_FromConvEnumAnyValToInte(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"b": 2, "a": 1}
	result := dm.SortedKeyValues()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SortedKeyValues returns non-empty -- with args", actual)
}

func Test_DynamicMap_SortedKeyAnyValues_FromConvEnumAnyValToInte(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"b": "y", "a": "x"}
	result := dm.SortedKeyAnyValues()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SortedKeyAnyValues returns non-empty -- with args", actual)
}

func Test_DynamicMap_JsonString_NonNil(t *testing.T) {
	// Arrange
	dm := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	result := dm.JsonString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight_JsonString returns correct value -- with args", actual)
}
