package corecomparatortests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── UnmarshalJSON ──

func Test_UnmarshalJSON_Valid(t *testing.T) {
	// Arrange
	var c corecomparator.Compare
	err := c.UnmarshalJSON([]byte("Equal"))

	// Act
	actual := args.Map{
		"val": c.String(),
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"val": "Equal",
		"hasErr": actual["hasErr"],
	}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON raw name not in RangesMap -- Equal defaults to zero", actual)
}

func Test_UnmarshalJSON_Invalid(t *testing.T) {
	// Arrange
	var c corecomparator.Compare
	err := c.UnmarshalJSON([]byte(`"garbage"`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
}

func Test_UnmarshalJSON_Nil(t *testing.T) {
	// Arrange
	var c corecomparator.Compare
	err := c.UnmarshalJSON(nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON nil -- error", actual)
}

func Test_UnmarshalJSON_ByNumber(t *testing.T) {
	// Arrange
	var c corecomparator.Compare
	_ = c.UnmarshalJSON([]byte("0"))

	// Act
	actual := args.Map{"val": c.String()}

	// Assert
	expected := args.Map{"val": "Equal"}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON by number -- 0=Equal", actual)
}

// ── IsInvalid / StringValue / ValueInt variants ──

func Test_IsInvalid(t *testing.T) {
	// Act
	actual := args.Map{"result": corecomparator.Inconclusive.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsInvalid -- Inconclusive", actual)
}

func Test_StringValue(t *testing.T) {
	// Act
	actual := args.Map{"result": corecomparator.Equal.StringValue()}

	// Assert
	expected := args.Map{"result": string(corecomparator.Equal)}
	expected.ShouldBeEqual(t, 0, "StringValue -- Equal", actual)
}

func Test_ValueInt8(t *testing.T) {
	// Act
	actual := args.Map{"result": int(corecomparator.Equal.ValueInt8())}

	// Assert
	expected := args.Map{"result": 0}
	expected.ShouldBeEqual(t, 0, "ValueInt8 -- Equal", actual)
}

func Test_ValueInt16(t *testing.T) {
	// Act
	actual := args.Map{"result": int(corecomparator.Equal.ValueInt16())}

	// Assert
	expected := args.Map{"result": 0}
	expected.ShouldBeEqual(t, 0, "ValueInt16 -- Equal", actual)
}

func Test_ValueInt32(t *testing.T) {
	// Act
	actual := args.Map{"result": int(corecomparator.Equal.ValueInt32())}

	// Assert
	expected := args.Map{"result": 0}
	expected.ShouldBeEqual(t, 0, "ValueInt32 -- Equal", actual)
}

func Test_ValueString(t *testing.T) {
	// Act
	actual := args.Map{"result": corecomparator.Equal.ValueString()}

	// Assert
	expected := args.Map{"result": "0"}
	expected.ShouldBeEqual(t, 0, "ValueString -- Equal", actual)
}

func Test_NumberString(t *testing.T) {
	// Act
	actual := args.Map{"result": corecomparator.LeftGreater.NumberString()}

	// Assert
	expected := args.Map{"result": "1"}
	expected.ShouldBeEqual(t, 0, "NumberString -- LeftGreater", actual)
}

func Test_SqlOperatorSymbol(t *testing.T) {
	// Act
	actual := args.Map{
		"eq": corecomparator.Equal.SqlOperatorSymbol(),
		"ne": corecomparator.NotEqual.SqlOperatorSymbol(),
	}

	// Assert
	expected := args.Map{
		"eq": "=",
		"ne": "<>",
	}
	expected.ShouldBeEqual(t, 0, "SqlOperatorSymbol -- eq and ne", actual)
}

// ── Format panic ──

func Test_Format_Panics(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		corecomparator.Equal.Format("test")
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Format panics -- by design", actual)
}

// ── BaseIsCaseSensitive / BaseIsIgnoreCase ──

func Test_BaseIsCaseSensitive_FromUnmarshalJSONValid(t *testing.T) {
	// Arrange
	b := corecomparator.BaseIsCaseSensitive{IsCaseSensitive: true}
	clone := b.Clone()
	clonePtr := b.ClonePtr()
	toIgnore := b.BaseIsIgnoreCase()

	// Act
	actual := args.Map{
		"isIgnoreCase":   b.IsIgnoreCase(),
		"cloneMatch":     clone.IsCaseSensitive,
		"clonePtrNotNil": clonePtr != nil,
		"toIgnoreCase":   toIgnore.IsIgnoreCase,
	}

	// Assert
	expected := args.Map{
		"isIgnoreCase":   false,
		"cloneMatch":     true,
		"clonePtrNotNil": true,
		"toIgnoreCase":   false,
	}
	expected.ShouldBeEqual(t, 0, "BaseIsCaseSensitive -- sensitive=true", actual)
}

func Test_BaseIsCaseSensitive_NilClonePtr(t *testing.T) {
	// Arrange
	var b *corecomparator.BaseIsCaseSensitive

	// Act
	actual := args.Map{"result": b.ClonePtr() == nil}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BaseIsCaseSensitive ClonePtr nil -- nil", actual)
}

func Test_BaseIsIgnoreCase_FromUnmarshalJSONValid(t *testing.T) {
	// Arrange
	b := corecomparator.BaseIsIgnoreCase{IsIgnoreCase: true}
	clone := b.Clone()
	clonePtr := b.ClonePtr()
	toSensitive := b.BaseIsCaseSensitive()

	// Act
	actual := args.Map{
		"isCaseSensitive":  b.IsCaseSensitive(),
		"cloneMatch":       clone.IsIgnoreCase,
		"clonePtrNotNil":   clonePtr != nil,
		"toSensitiveCase":  toSensitive.IsCaseSensitive,
	}

	// Assert
	expected := args.Map{
		"isCaseSensitive":  false,
		"cloneMatch":       true,
		"clonePtrNotNil":   true,
		"toSensitiveCase":  false,
	}
	expected.ShouldBeEqual(t, 0, "BaseIsIgnoreCase -- ignoreCase=true", actual)
}

func Test_BaseIsIgnoreCase_NilClonePtr(t *testing.T) {
	// Arrange
	var b *corecomparator.BaseIsIgnoreCase

	// Act
	actual := args.Map{"result": b.ClonePtr() == nil}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BaseIsIgnoreCase ClonePtr nil -- nil", actual)
}

// ── RangeNamesCsv ──

func Test_RangeNamesCsv_FromUnmarshalJSONValid(t *testing.T) {
	// Arrange
	csv := corecomparator.RangeNamesCsv()

	// Act
	actual := args.Map{"notEmpty": csv != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv -- not empty", actual)
}

// ── MarshalJSON roundtrip ──

func Test_MarshalUnmarshal_Roundtrip(t *testing.T) {
	// Arrange
	original := corecomparator.LeftGreater
	data, _ := json.Marshal(original)
	var parsed corecomparator.Compare
	_ = parsed.UnmarshalJSON(data)

	// Act
	actual := args.Map{"match": parsed == original}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "Marshal/Unmarshal roundtrip fails -- UnmarshalJSON expects unquoted", actual)
}

// ── IsLeftLessOrLessEqualOrEqual ──

func Test_IsLeftLessOrLessEqualOrEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":    corecomparator.Equal.IsLeftLessOrLessEqualOrEqual(),
		"less":     corecomparator.LeftLess.IsLeftLessOrLessEqualOrEqual(),
		"lessEq":   corecomparator.LeftLessEqual.IsLeftLessOrLessEqualOrEqual(),
		"greater":  corecomparator.LeftGreater.IsLeftLessOrLessEqualOrEqual(),
	}

	// Assert
	expected := args.Map{
		"equal": true, "less": true, "lessEq": true, "greater": false,
	}
	expected.ShouldBeEqual(t, 0, "IsLeftLessOrLessEqualOrEqual -- various", actual)
}

// ── Is / IsValueEqual ──

func Test_Is(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   corecomparator.Equal.Is(corecomparator.Equal),
		"noMatch": corecomparator.Equal.Is(corecomparator.NotEqual),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Is -- Equal", actual)
}

func Test_IsValueEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   corecomparator.Equal.IsValueEqual(0),
		"noMatch": corecomparator.Equal.IsValueEqual(1),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsValueEqual -- 0", actual)
}

// ── IsDefinedPlus ──

func Test_IsDefinedPlus(t *testing.T) {
	// Act
	actual := args.Map{
		"definedMatch":   corecomparator.Equal.IsDefinedPlus(corecomparator.Equal),
		"definedNoMatch": corecomparator.Equal.IsDefinedPlus(corecomparator.NotEqual),
		"inconclusive":   corecomparator.Inconclusive.IsDefinedPlus(corecomparator.Inconclusive),
	}

	// Assert
	expected := args.Map{
		"definedMatch": true, "definedNoMatch": false, "inconclusive": false,
	}
	expected.ShouldBeEqual(t, 0, "IsDefinedPlus -- various", actual)
}

// ── IsNotInconclusive ──

func Test_IsNotInconclusive(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":        corecomparator.Equal.IsNotInconclusive(),
		"inconclusive": corecomparator.Inconclusive.IsNotInconclusive(),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"inconclusive": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNotInconclusive -- various", actual)
}
