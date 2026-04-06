package versionindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/versionindexes"
)

// ── Extended enum methods ──

func Test_Cov2_Index_ValueAccessors(t *testing.T) {
	v := versionindexes.Minor
	actual := args.Map{
		"valueInt8":   v.ValueInt8(),
		"valueInt16":  v.ValueInt16(),
		"valueInt32":  v.ValueInt32(),
		"valueUInt16": v.ValueUInt16(),
		"valueByte":   v.ValueByte(),
		"valueInt":    v.ValueInt(),
		"valueStr":    v.ValueString(),
	}
	expected := args.Map{
		"valueInt8": int8(1), "valueInt16": int16(1), "valueInt32": int32(1),
		"valueUInt16": uint16(1), "valueByte": byte(1), "valueInt": 1,
		"valueStr": v.ToNumberString(),
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- ValueAccessors", actual)
}

func Test_Cov2_Index_MinMaxAny(t *testing.T) {
	v := versionindexes.Major
	min, max := v.MinMaxAny()
	actual := args.Map{"minNotNil": min != nil, "maxNotNil": max != nil}
	expected := args.Map{"minNotNil": true, "maxNotNil": true}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- MinMaxAny", actual)
}

func Test_Cov2_Index_MinMaxValueString(t *testing.T) {
	v := versionindexes.Major
	actual := args.Map{
		"minStr": v.MinValueString() != "",
		"maxStr": v.MaxValueString() != "",
		"minInt": v.MinInt() >= 0,
		"maxInt": v.MaxInt() > 0,
	}
	expected := args.Map{"minStr": true, "maxStr": true, "minInt": true, "maxInt": true}
	expected.ShouldBeEqual(t, 0, "Index returns non-empty -- MinMaxValueString", actual)
}

func Test_Cov2_Index_RangesDynamic(t *testing.T) {
	v := versionindexes.Major
	actual := args.Map{
		"rangesMap": len(v.RangesDynamicMap()) > 0,
		"intRanges": len(v.IntegerEnumRanges()) > 0,
		"rangesByte": len(v.RangesByte()) > 0,
		"maxByte":   v.MaxByte() > 0,
		"minByte":   v.MinByte(),
	}
	expected := args.Map{
		"rangesMap": true, "intRanges": true, "rangesByte": true,
		"maxByte": true, "minByte": byte(0),
	}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- RangesDynamic", actual)
}

func Test_Cov2_Index_OnlySupportedErr(t *testing.T) {
	noErr := versionindexes.Major.OnlySupportedErr("Major")
	hasErr := versionindexes.Invalid.OnlySupportedMsgErr("msg", "Major")
	noErrResult := noErr == nil
	actual := args.Map{"noErr": noErrResult, "hasErr": hasErr != nil}
	expected := args.Map{"noErr": noErrResult, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "Index returns error -- OnlySupportedErr", actual)
}

func Test_Cov2_Index_Format(t *testing.T) {
	result := versionindexes.Major.Format("%s")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- Format", actual)
}

func Test_Cov2_Index_IsEnumEqual(t *testing.T) {
	maj := versionindexes.Major
	mp := &maj
	min := versionindexes.Minor
	minp := &min
	actual := args.Map{
		"equal":    versionindexes.Major.IsEnumEqual(mp),
		"notEqual": versionindexes.Major.IsEnumEqual(minp),
	}
	expected := args.Map{"equal": true, "notEqual": false}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- IsEnumEqual", actual)
}

func Test_Cov2_Index_IsAnyEnumsEqual(t *testing.T) {
	v := versionindexes.Minor
	maj := versionindexes.Major
	mp := &maj
	minv := versionindexes.Minor
	minp := &minv
	patch := versionindexes.Patch
	pp := &patch
	build := versionindexes.Build
	bp := &build
	actual := args.Map{
		"found":    v.IsAnyEnumsEqual(mp, minp),
		"notFound": v.IsAnyEnumsEqual(pp, bp),
	}
	expected := args.Map{"found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- IsAnyEnumsEqual", actual)
}

func Test_Cov2_Index_IsByteValueEqual(t *testing.T) {
	actual := args.Map{
		"equal":    versionindexes.Major.IsByteValueEqual(0),
		"notEqual": versionindexes.Major.IsByteValueEqual(1),
	}
	expected := args.Map{"equal": true, "notEqual": false}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- IsByteValueEqual", actual)
}

func Test_Cov2_Index_IsAnyValuesEqual(t *testing.T) {
	actual := args.Map{
		"found":    versionindexes.Minor.IsAnyValuesEqual(0, 1),
		"notFound": versionindexes.Minor.IsAnyValuesEqual(2, 3),
	}
	expected := args.Map{"found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "Index returns non-empty -- IsAnyValuesEqual", actual)
}

func Test_Cov2_Index_IsAnyNamesOf(t *testing.T) {
	actual := args.Map{
		"found":    versionindexes.Minor.IsAnyNamesOf("Major", "Minor"),
		"notFound": versionindexes.Minor.IsAnyNamesOf("Patch", "Build"),
	}
	expected := args.Map{"found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- IsAnyNamesOf", actual)
}

func Test_Cov2_Index_EnumType(t *testing.T) {
	actual := args.Map{"notNil": versionindexes.Major.EnumType() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- EnumType", actual)
}

func Test_Cov2_Index_Contracts(t *testing.T) {
	v := versionindexes.Major
	actual := args.Map{
		"basicBinder":     v.AsBasicEnumContractsBinder() != nil,
		"jsonBinder":      v.AsJsonContractsBinder() != nil,
		"basicByteBinder": v.AsBasicByteEnumContractsBinder() != nil,
	}
	expected := args.Map{"basicBinder": true, "jsonBinder": true, "basicByteBinder": true}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- Contracts", actual)
}

func Test_Cov2_Index_ToPtr(t *testing.T) {
	v := versionindexes.Major
	p := v.ToPtr()
	actual := args.Map{"notNil": p != nil, "isMajor": p.IsMajor()}
	expected := args.Map{"notNil": true, "isMajor": true}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- ToPtr", actual)
}

func Test_Cov2_Index_JsonParseSelfInject(t *testing.T) {
	v := versionindexes.Minor
	jsonResult := v.JsonPtr()
	var v2 versionindexes.Index
	err := v2.JsonParseSelfInject(jsonResult)
	errNil := v2.JsonParseSelfInject(nil)
	actual := args.Map{
		"noErr": err == nil, "isMinor": v2.IsMinor(),
		"nilErr": errNil != nil,
	}
	expected := args.Map{"noErr": true, "isMinor": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- JsonParseSelfInject", actual)
}

func Test_Cov2_Index_Json(t *testing.T) {
	v := versionindexes.Major
	json := v.Json()
	jsonPtr := v.JsonPtr()
	actual := args.Map{
		"noErr":      json.HasError() == false,
		"ptrNotNil":  jsonPtr != nil,
	}
	expected := args.Map{"noErr": true, "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "Index returns correct value -- Json", actual)
}

func Test_Cov2_Index_UnmarshalJSON_Invalid(t *testing.T) {
	var v versionindexes.Index
	err := v.UnmarshalJSON([]byte("invalid"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Index returns error -- UnmarshalJSON invalid", actual)
}
