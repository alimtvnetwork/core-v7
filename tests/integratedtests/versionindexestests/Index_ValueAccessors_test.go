package versionindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/versionindexes"
)

// ── Extended enum methods ──

func Test_Index_ValueAccessors(t *testing.T) {
	// Arrange
	v := versionindexes.Minor

	// Act
	actual := args.Map{
		"valueInt8":   v.ValueInt8(),
		"valueInt16":  v.ValueInt16(),
		"valueInt32":  v.ValueInt32(),
		"valueUInt16": v.ValueUInt16(),
		"valueByte":   v.ValueByte(),
		"valueInt":    v.ValueInt(),
		"valueStr":    v.ValueString(),
	}

	// Assert
	expected := args.Map{
		"valueInt8": int8(1), "valueInt16": int16(1), "valueInt32": int32(1),
		"valueUInt16": uint16(1), "valueByte": byte(1), "valueInt": 1,
		"valueStr": v.ToNumberString(),
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- ValueAccessors", actual)
}

func Test_Index_MinMaxAny(t *testing.T) {
	// Arrange
	v := versionindexes.Major
	min, max := v.MinMaxAny()

	// Act
	actual := args.Map{
		"minNotNil": min != nil,
		"maxNotNil": max != nil,
	}

	// Assert
	expected := args.Map{
		"minNotNil": true,
		"maxNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- MinMaxAny", actual)
}

func Test_Index_MinMaxValueString(t *testing.T) {
	// Arrange
	v := versionindexes.Major

	// Act
	actual := args.Map{
		"minStr": v.MinValueString() != "",
		"maxStr": v.MaxValueString() != "",
		"minInt": v.MinInt() >= 0,
		"maxInt": v.MaxInt() > 0,
	}

	// Assert
	expected := args.Map{
		"minStr": true,
		"maxStr": true,
		"minInt": true,
		"maxInt": true,
	}
	expected.ShouldBeEqual(t, 0, "Index returns non-empty -- MinMaxValueString", actual)
}

func Test_Index_RangesDynamic(t *testing.T) {
	// Arrange
	v := versionindexes.Major

	// Act
	actual := args.Map{
		"rangesMap": len(v.RangesDynamicMap()) > 0,
		"intRanges": len(v.IntegerEnumRanges()) > 0,
		"rangesByte": len(v.RangesByte()) > 0,
		"maxByte":   v.MaxByte() > 0,
		"minByte":   v.MinByte(),
	}

	// Assert
	expected := args.Map{
		"rangesMap": true, "intRanges": true, "rangesByte": true,
		"maxByte": true, "minByte": byte(0),
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- RangesDynamic", actual)
}

func Test_Index_OnlySupportedErr(t *testing.T) {
	// Arrange
	noErr := versionindexes.Major.OnlySupportedErr("Major")
	hasErr := versionindexes.Invalid.OnlySupportedMsgErr("msg", "Major")
	noErrResult := noErr == nil

	// Act
	actual := args.Map{
		"noErr": noErrResult,
		"hasErr": hasErr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": noErrResult,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Index returns error -- OnlySupportedErr", actual)
}

func Test_Index_Format(t *testing.T) {
	// Arrange
	result := versionindexes.Major.Format("%s")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- Format", actual)
}

func Test_Index_IsEnumEqual(t *testing.T) {
	// Arrange
	maj := versionindexes.Major
	mp := &maj
	min := versionindexes.Minor
	minp := &min

	// Act
	actual := args.Map{
		"equal":    versionindexes.Major.IsEnumEqual(mp),
		"notEqual": versionindexes.Major.IsEnumEqual(minp),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- IsEnumEqual", actual)
}

func Test_Index_IsAnyEnumsEqual(t *testing.T) {
	// Arrange
	v := versionindexes.Minor
	maj := versionindexes.Major
	mp := &maj
	minv := versionindexes.Minor
	minp := &minv
	patch := versionindexes.Patch
	pp := &patch
	build := versionindexes.Build
	bp := &build

	// Act
	actual := args.Map{
		"found":    v.IsAnyEnumsEqual(mp, minp),
		"notFound": v.IsAnyEnumsEqual(pp, bp),
	}

	// Assert
	expected := args.Map{
		"found": true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- IsAnyEnumsEqual", actual)
}

func Test_Index_IsByteValueEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":    versionindexes.Major.IsByteValueEqual(0),
		"notEqual": versionindexes.Major.IsByteValueEqual(1),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- IsByteValueEqual", actual)
}

func Test_Index_IsAnyValuesEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"found":    versionindexes.Minor.IsAnyValuesEqual(0, 1),
		"notFound": versionindexes.Minor.IsAnyValuesEqual(2, 3),
	}

	// Assert
	expected := args.Map{
		"found": true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "Index returns non-empty -- IsAnyValuesEqual", actual)
}

func Test_Index_IsAnyNamesOf(t *testing.T) {
	// Act
	actual := args.Map{
		"found":    versionindexes.Minor.IsAnyNamesOf("Major", "Minor"),
		"notFound": versionindexes.Minor.IsAnyNamesOf("Patch", "Build"),
	}

	// Assert
	expected := args.Map{
		"found": true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- IsAnyNamesOf", actual)
}

func Test_Index_EnumType(t *testing.T) {
	// Act
	actual := args.Map{"notNil": versionindexes.Major.EnumType() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- EnumType", actual)
}

func Test_Index_Contracts(t *testing.T) {
	// Arrange
	v := versionindexes.Major

	// Act
	actual := args.Map{
		"basicBinder":     v.AsBasicEnumContractsBinder() != nil,
		"jsonBinder":      v.AsJsonContractsBinder() != nil,
		"basicByteBinder": v.AsBasicByteEnumContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"basicBinder": true,
		"jsonBinder": true,
		"basicByteBinder": true,
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- Contracts", actual)
}

func Test_Index_ToPtr(t *testing.T) {
	// Arrange
	v := versionindexes.Major
	p := v.ToPtr()

	// Act
	actual := args.Map{
		"notNil": p != nil,
		"isMajor": p.IsMajor(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"isMajor": true,
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- ToPtr", actual)
}

func Test_Index_JsonParseSelfInject(t *testing.T) {
	// Arrange
	v := versionindexes.Minor
	jsonResult := v.JsonPtr()
	var v2 versionindexes.Index
	err := v2.JsonParseSelfInject(jsonResult)
	errNil := v2.JsonParseSelfInject(nil)

	// Act
	actual := args.Map{
		"noErr": err == nil, "isMinor": v2.IsMinor(),
		"nilErr": errNil != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"isMinor": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- JsonParseSelfInject", actual)
}

func Test_Index_Json(t *testing.T) {
	// Arrange
	v := versionindexes.Major
	json := v.Json()
	jsonPtr := v.JsonPtr()

	// Act
	actual := args.Map{
		"noErr":      json.HasError() == false,
		"ptrNotNil":  jsonPtr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"ptrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- Json", actual)
}

func Test_Index_UnmarshalJSON_Invalid(t *testing.T) {
	// Arrange
	var v versionindexes.Index
	err := v.UnmarshalJSON([]byte("invalid"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Index returns error -- UnmarshalJSON invalid", actual)
}
