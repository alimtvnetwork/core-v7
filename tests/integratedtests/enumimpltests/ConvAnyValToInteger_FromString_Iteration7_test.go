package enumimpltests

import (
	"fmt"
	"math"
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// ConvAnyValToInteger — type switch branches
// Covers ConvAnyValToInteger.go L25-40
// ══════════════════════════════════════════════════════════════════════════════

func Test_ConvAnyValToInteger_FromString(t *testing.T) {
	// Arrange
	result := enumimpl.ConvEnumAnyValToInteger("hello")

	// Act
	actual := args.Map{"isMinInt": result < -999999}

	// Assert
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger returns correct value -- string", actual)
}

func Test_ConvAnyValToInteger_FromInt(t *testing.T) {
	// Arrange
	result := enumimpl.ConvEnumAnyValToInteger(42)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": 42}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger returns correct value -- int", actual)
}

func Test_ConvAnyValToInteger_FromFloat(t *testing.T) {
	// Arrange
	// float64 won't match any switch case, falls through to Atoi("3.14") → fail
	result := enumimpl.ConvEnumAnyValToInteger(3.14)

	// Act
	actual := args.Map{"isMinInt": result < -999999}

	// Assert
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger returns correct value -- float", actual)
}

func Test_ConvAnyValToInteger_FromByte(t *testing.T) {
	// Arrange
	// byte is uint8 — passes through Atoi path
	result := enumimpl.ConvEnumAnyValToInteger(byte(7))

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": 7}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger returns correct value -- byte", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicMap — isEqualSingle regardless, ConvMap overflow branches
// Covers DynamicMap.go L867-881, L1263, L1288, L1313, L1338, L1363
// ══════════════════════════════════════════════════════════════════════════════

func Test_DynamicMap_DiffRaw_RegardlessType(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": "1"}
	diff := dm.DiffRaw(true, enumimpl.DynamicMap{"a": "1", "b": 1})

	// Act
	actual := args.Map{"diffLen": len(diff)}

	// Assert
	expected := args.Map{"diffLen": 0}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns correct value -- regardless type", actual)
}

func Test_DynamicMap_ConvMapInt8String_Overflow(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"overflow": int(math.MaxInt8 + 1)}
	result := dm.ConvMapInt8String()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapInt8String returns correct value -- overflow", actual)
}

func Test_DynamicMap_ConvMapInt16String_Overflow(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"overflow": int(math.MaxInt16 + 1)}
	result := dm.ConvMapInt16String()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapInt16String returns correct value -- overflow", actual)
}

func Test_DynamicMap_ConvMapInt32String_Overflow(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"overflow": int(math.MaxInt32 + 1)}
	result := dm.ConvMapInt32String()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapInt32String returns correct value -- overflow", actual)
}

func Test_DynamicMap_ConvMapUInt16String_Negative(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"negative": int(-1)}
	result := dm.ConvMapUInt16String()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapUInt16String returns correct value -- negative", actual)
}

func Test_DynamicMap_ConvMapStringString_NotFound(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"key": 123}
	result := dm.ConvMapStringString()
	// KeyValueString uses fmt.Sprintf so int 123 → "123" is found; len is 1

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ConvMapStringString returns correct value -- int value converts to string", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicMap — KeyValueByte and KeyValueInt edge cases
// Covers DynamicMap.go L958-970, L987-989, L1023-1048
// ══════════════════════════════════════════════════════════════════════════════

func Test_DynamicMap_KeyValueByte_NotANumber(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"key": "not-a-number"}
	val, isFound, isFailed := dm.KeyValueByte("key")

	// Act
	actual := args.Map{
		"val": fmt.Sprintf("%d", val),
		"isFound": isFound,
		"isFailed": isFailed,
	}

	// Assert
	expected := args.Map{
		"val": "0",
		"isFound": true,
		"isFailed": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyValueByte returns correct value -- not-a-number", actual)
}

func Test_DynamicMap_KeyValueByte_OutOfRange(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"key": 999}
	val, isFound, isFailed := dm.KeyValueByte("key")

	// Act
	actual := args.Map{
		"val": fmt.Sprintf("%d", val),
		"isFound": isFound,
		"isFailed": isFailed,
	}

	// Assert
	expected := args.Map{
		"val": "0",
		"isFound": true,
		"isFailed": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValueByte returns correct value -- out of range", actual)
}

func Test_DynamicMap_KeyValueInt_NotANumber(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"key": "not-int"}
	_, isFound, isFailed := dm.KeyValueInt("key")

	// Act
	actual := args.Map{
		"isFound": isFound,
		"isFailed": isFailed,
	}

	// Assert
	expected := args.Map{
		"isFound": true,
		"isFailed": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValueInt returns correct value -- not-a-number", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicByte/BasicString — GetValueByName wrapped-quote branch
// Covers BasicByte.go L81, BasicString.go L139
// ══════════════════════════════════════════════════════════════════════════════

func Test_BasicByte_GetValueByName_WrappedQuote(t *testing.T) {
	dm := enumimpl.DynamicMap{"Alpha": byte(0), "Beta": byte(1)}
	bb := dm.BasicByte("TestByte")

	// Act — pass unwrapped name; the map stores "Alpha" as double-quoted key
	// The method first tries exact match, then wraps with quotes
	val, err := bb.GetValueByName("Alpha")

	actual := args.Map{
		"val": fmt.Sprintf("%d", val),
		"hasErr": err != nil,
	}
	expected := args.Map{
		"val": "0",
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte returns correct value -- GetValueByName", actual)
}

func Test_BasicString_GetValueByName_WrappedQuote(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Alpha": "Alpha", "Beta": "Beta"}
	bs := dm.BasicString("TestString")

	val, err := bs.GetValueByName("Alpha")

	// Act
	actual := args.Map{
		"val": val,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"val": "Alpha",
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BasicString returns correct value -- GetValueByName", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// newBasicStringCreator — CreateUsingStringersSpread
// Covers newBasicStringCreator.go L139-164
// ══════════════════════════════════════════════════════════════════════════════

type testStringer struct{ name string }

func (s testStringer) String() string { return s.name }

func Test_NewBasicStringCreator_CreateUsingStringersSpread(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.CreateUsingStringersSpread(
		"TestStringerEnum",
		testStringer{"Alpha"},
		testStringer{"Beta"},
		testStringer{"Gamma"},
	)

	// Act
	actual := args.Map{
		"typeName": bs.TypeName(),
		"length": bs.Length(),
	}

	// Assert
	expected := args.Map{
		"typeName": "TestStringerEnum",
		"length": 3,
	}
	expected.ShouldBeEqual(t, 0, "CreateUsingStringersSpread returns correct value -- with args", actual)
}

func Test_NewBasicStringCreator_CreateUsingStringersSpread_MinMax(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.CreateUsingStringersSpread(
		"MinMaxEnum",
		testStringer{"C"},
		testStringer{"A"},
		testStringer{"B"},
	)

	// Act
	actual := args.Map{"length": bs.Length()}

	// Assert
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "CreateUsingStringersSpread returns correct value -- min/max", actual)
}
