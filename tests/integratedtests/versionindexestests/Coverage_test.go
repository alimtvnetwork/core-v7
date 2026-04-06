package versionindexestests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/versionindexes"
)

func Test_Index_EnumMethods(t *testing.T) {
	// Arrange
	v := versionindexes.Major

	// Act
	actual := args.Map{
		"name":           v.Name(),
		"nameValue":      v.NameValue(),
		"typeName":       v.TypeName(),
		"isValid":        v.IsValid(),
		"isInvalid":      v.IsInvalid(),
		"isMajor":        v.IsMajor(),
		"isMinor":        v.IsMinor(),
		"isPatch":        v.IsPatch(),
		"isBuild":        v.IsBuild(),
		"valueInt":       v.ValueInt(),
		"valueByte":      int(v.ValueByte()),
		"valueString":    v.ValueString(),
		"toNumberString": v.ToNumberString(),
		"stringVal":      v.String(),
		"rangeNamesCsv":  v.RangeNamesCsv() != "",
	}
	expected := args.Map{
		"name":           "Major",
		"nameValue":      v.NameValue(),
		"typeName":       v.TypeName(),
		"isValid":        true,
		"isInvalid":      false,
		"isMajor":        true,
		"isMinor":        false,
		"isPatch":        false,
		"isBuild":        false,
		"valueInt":       0,
		"valueByte":      0,
		"valueString":    v.ToNumberString(),
		"toNumberString": v.ToNumberString(),
		"stringVal":      "Major",
		"rangeNamesCsv":  true,
	}
	expected.ShouldBeEqual(t, 0, "Index_EnumMethods returns correct value -- with args", actual)
}

func Test_Index_AllVariants(t *testing.T) {
	// Act
	actual := args.Map{
		"majorValid":   versionindexes.Major.IsValid(),
		"minorValid":   versionindexes.Minor.IsValid(),
		"patchValid":   versionindexes.Patch.IsValid(),
		"buildValid":   versionindexes.Build.IsValid(),
		"invalidValid": versionindexes.Invalid.IsValid(),
	}
	expected := args.Map{
		"majorValid":   true,
		"minorValid":   true,
		"patchValid":   true,
		"buildValid":   true,
		"invalidValid": false,
	}
	expected.ShouldBeEqual(t, 0, "Index_AllVariants returns correct value -- with args", actual)
}

func Test_Index_Comparisons(t *testing.T) {
	// Arrange
	v := versionindexes.Minor

	// Act
	actual := args.Map{
		"isNameEqual":    v.IsNameEqual("Minor"),
		"isNameNotEqual": v.IsNameEqual("Major"),
		"isValueEqual":   v.IsValueEqual(byte(versionindexes.Minor)),
		"isByteValueEq":  v.IsByteValueEqual(byte(versionindexes.Minor)),
		"isAnyNamesOf":   v.IsAnyNamesOf("Major", "Minor"),
		"isAnyValsEq":    v.IsAnyValuesEqual(byte(versionindexes.Major), byte(versionindexes.Minor)),
	}
	expected := args.Map{
		"isNameEqual":    true,
		"isNameNotEqual": false,
		"isValueEqual":   true,
		"isByteValueEq":  true,
		"isAnyNamesOf":   true,
		"isAnyValsEq":    true,
	}
	expected.ShouldBeEqual(t, 0, "Index_Comparisons returns correct value -- with args", actual)
}

func Test_Index_JSON(t *testing.T) {
	// Arrange
	v := versionindexes.Patch

	// Act
	data, err := json.Marshal(v)
	var parsed versionindexes.Index
	errUnmarshal := json.Unmarshal(data, &parsed)

	actual := args.Map{
		"marshalErr":   fmt.Sprintf("%v", err),
		"unmarshalErr": fmt.Sprintf("%v", errUnmarshal),
		"roundTrip":    parsed == v,
	}
	expected := args.Map{
		"marshalErr":   "<nil>",
		"unmarshalErr": "<nil>",
		"roundTrip":    true,
	}
	expected.ShouldBeEqual(t, 0, "Index_JSON returns correct value -- with args", actual)
}

func Test_Index_Binders(t *testing.T) {
	// Arrange
	v := versionindexes.Major

	// Act
	actual := args.Map{
		"enumContractsNil": v.AsBasicEnumContractsBinder() == nil,
		"jsonContractsNil": v.AsJsonContractsBinder() == nil,
		"byteContractsNil": v.AsBasicByteEnumContractsBinder() == nil,
		"toPtrNil":         v.ToPtr() == nil,
		"enumTypeNil":      v.EnumType() == nil,
	}
	expected := args.Map{
		"enumContractsNil": false,
		"jsonContractsNil": false,
		"byteContractsNil": false,
		"toPtrNil":         false,
		"enumTypeNil":      false,
	}
	expected.ShouldBeEqual(t, 0, "Index_Binders returns correct value -- with args", actual)
}