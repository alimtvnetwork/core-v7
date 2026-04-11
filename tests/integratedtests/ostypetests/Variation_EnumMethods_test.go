package ostypetests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/ostype"
)

// ── Variation enum methods ──

func Test_Variation_EnumMethods(t *testing.T) {
	// Arrange
	v := ostype.Windows

	// Act
	actual := args.Map{
		"name":           v.Name(),
		"nameValue":      v.NameValue(),
		"typeName":       v.TypeName(),
		"rangeNamesCsv":  v.RangeNamesCsv() != "",
		"valueInt":       v.ValueInt(),
		"valueByte":      int(v.ValueByte()),
		"valueString":    v.ValueString(),
		"toNumberString": v.ToNumberString(),
		"valueUInt16":    int(v.ValueUInt16()),
		"valueInt8":      int(v.ValueInt8()),
		"valueInt16":     int(v.ValueInt16()),
		"valueInt32":     int(v.ValueInt32()),
		"goosName":       v.GoosName(),
		"isValid":        v.IsValid(),
		"isInvalid":      v.IsInvalid(),
		"stringVal":      v.String(),
	}
	expected := args.Map{
		"name":           "windows",
		"nameValue":      v.NameValue(),
		"typeName":       v.TypeName(),
		"rangeNamesCsv":  true,
		"valueInt":       int(ostype.Windows),
		"valueByte":      int(ostype.Windows),
		"valueString":    v.ToNumberString(),
		"toNumberString": v.ToNumberString(),
		"valueUInt16":    int(ostype.Windows),
		"valueInt8":      int(ostype.Windows),
		"valueInt16":     int(ostype.Windows),
		"valueInt32":     int(ostype.Windows),
		"goosName":       "windows",
		"isValid":        true,
		"isInvalid":      false,
		"stringVal":      "windows",
	}
	expected.ShouldBeEqual(t, 0, "Variation_EnumMethods returns correct value -- with args", actual)
}

func Test_Variation_Identity_Extended(t *testing.T) {
	// Act
	actual := args.Map{
		"windowsIsWindows":   ostype.Windows.IsWindows(),
		"linuxIsLinux":       ostype.Linux.IsLinux(),
		"darwinIsDarwin":     ostype.DarwinOrMacOs.IsDarwinOrMacOs(),
		"jsIsJavaScript":     ostype.JavaScript.IsJavaScript(),
		"freebsdIsFreeBsd":   ostype.FreeBsd.IsFreeBsd(),
		"netbsdIsNetBsd":     ostype.NetBsd.IsNetBsd(),
		"openbsdIsOpenBsd":   ostype.OpenBsd.IsOpenBsd(),
		"dragonIsDF":         ostype.DragonFly.IsDragonFly(),
		"linuxIsLinuxOrMac":  ostype.Linux.IsLinuxOrMac(),
		"darwinIsLinuxOrMac": ostype.DarwinOrMacOs.IsLinuxOrMac(),
		"anyIsAny":           ostype.Any.IsAnyOperatingSystem(),
	}
	expected := args.Map{
		"windowsIsWindows":   true,
		"linuxIsLinux":       true,
		"darwinIsDarwin":     true,
		"jsIsJavaScript":     true,
		"freebsdIsFreeBsd":   true,
		"netbsdIsNetBsd":     true,
		"openbsdIsOpenBsd":   true,
		"dragonIsDF":         true,
		"linuxIsLinuxOrMac":  true,
		"darwinIsLinuxOrMac": true,
		"anyIsAny":           true,
	}
	expected.ShouldBeEqual(t, 0, "Variation_Identity_Extended returns correct value -- with args", actual)
}

func Test_Variation_Is(t *testing.T) {
	// Act
	actual := args.Map{
		"isSame":  ostype.Windows.Is(ostype.Windows),
		"isDiff":  ostype.Windows.Is(ostype.Linux),
		"isByte":  ostype.Windows.IsByte(byte(ostype.Windows)),
		"isByteF": ostype.Windows.IsByte(byte(ostype.Linux)),
	}
	expected := args.Map{
		"isSame":  true,
		"isDiff":  false,
		"isByte":  true,
		"isByteF": false,
	}
	expected.ShouldBeEqual(t, 0, "Variation_Is returns correct value -- with args", actual)
}

func Test_Variation_IsAnyMatch(t *testing.T) {
	// Act
	actual := args.Map{
		"matchFound":    ostype.Windows.IsAnyMatch(ostype.Linux, ostype.Windows),
		"matchNotFound": ostype.Windows.IsAnyMatch(ostype.Linux, ostype.DarwinOrMacOs),
	}
	expected := args.Map{
		"matchFound":    true,
		"matchNotFound": false,
	}
	expected.ShouldBeEqual(t, 0, "Variation_IsAnyMatch returns correct value -- with args", actual)
}

func Test_Variation_IsStringsMatchAny(t *testing.T) {
	// Act
	actual := args.Map{
		"matchFound":    ostype.Windows.IsStringsMatchAny("linux", "windows"),
		"matchNotFound": ostype.Windows.IsStringsMatchAny("linux", "darwin"),
	}
	expected := args.Map{
		"matchFound":    true,
		"matchNotFound": false,
	}
	expected.ShouldBeEqual(t, 0, "Variation_IsStringsMatchAny returns correct value -- with args", actual)
}

func Test_Variation_IsNameEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"nameMatch":   ostype.Windows.IsNameEqual("windows"),
		"nameNoMatch": ostype.Windows.IsNameEqual("linux"),
		"anyNames":    ostype.Windows.IsAnyNamesOf("linux", "windows"),
		"anyNamesNo":  ostype.Windows.IsAnyNamesOf("linux", "darwin"),
	}
	expected := args.Map{
		"nameMatch":   true,
		"nameNoMatch": false,
		"anyNames":    true,
		"anyNamesNo":  false,
	}
	expected.ShouldBeEqual(t, 0, "Variation_IsNameEqual returns correct value -- with args", actual)
}

func Test_Variation_IsValueEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"valMatch":    ostype.Windows.IsValueEqual(byte(ostype.Windows)),
		"valNoMatch":  ostype.Windows.IsValueEqual(byte(ostype.Linux)),
		"anyValsT":    ostype.Windows.IsAnyValuesEqual(byte(ostype.Linux), byte(ostype.Windows)),
		"anyValsF":    ostype.Windows.IsAnyValuesEqual(byte(ostype.Linux), byte(ostype.DarwinOrMacOs)),
		"byteValEq":   ostype.Windows.IsByteValueEqual(byte(ostype.Windows)),
	}
	expected := args.Map{
		"valMatch":    true,
		"valNoMatch":  false,
		"anyValsT":    true,
		"anyValsF":    false,
		"byteValEq":   true,
	}
	expected.ShouldBeEqual(t, 0, "Variation_IsValueEqual returns correct value -- with args", actual)
}

func Test_Variation_Group_Extended(t *testing.T) {
	// Act
	actual := args.Map{
		"windowsGroup":    ostype.Windows.Group().Name(),
		"linuxGroup":      ostype.Linux.Group().Name(),
		"androidGroup":    ostype.Android.Group().Name(),
		"isUnixGroup":     ostype.Linux.IsActualGroupUnix(),
		"isPossibleUnix":  ostype.Linux.IsPossibleUnixGroup(),
		"windowsNotUnix":  ostype.Windows.IsPossibleUnixGroup(),
	}
	expected := args.Map{
		"windowsGroup":    "WindowsGroup",
		"linuxGroup":      "UnixGroup",
		"androidGroup":    "AndroidGroup",
		"isUnixGroup":     true,
		"isPossibleUnix":  true,
		"windowsNotUnix":  false,
	}
	expected.ShouldBeEqual(t, 0, "Variation_Group_Extended returns correct value -- with args", actual)
}

func Test_Variation_JSON_FromVariationEnumMethods(t *testing.T) {
	// Arrange
	v := ostype.Windows

	// Act
	data, err := json.Marshal(v)
	var parsed ostype.Variation
	errUnmarshal := json.Unmarshal(data, &parsed)

	// Assert
	actual := args.Map{
		"marshalErr":   fmt.Sprintf("%v", err),
		"unmarshalErr": fmt.Sprintf("%v", errUnmarshal),
		"roundTrip":    parsed == v,
		"dataNotEmpty": len(data) > 0,
	}
	expected := args.Map{
		"marshalErr":   "<nil>",
		"unmarshalErr": "<nil>",
		"roundTrip":    true,
		"dataNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Variation_JSON returns correct value -- with args", actual)
}

func Test_Variation_Binders_FromVariationEnumMethods(t *testing.T) {
	// Arrange
	v := ostype.Windows

	// Act
	actual := args.Map{
		"enumContractsNil": v.AsBasicEnumContractsBinder() == nil,
		"jsonContractsNil": v.AsJsonContractsBinder() == nil,
		"byteContractsNil": v.AsBasicByteEnumContractsBinder() == nil,
		"toPtrNil":         v.ToPtr() == nil,
	}
	expected := args.Map{
		"enumContractsNil": false,
		"jsonContractsNil": false,
		"byteContractsNil": false,
		"toPtrNil":         false,
	}
	expected.ShouldBeEqual(t, 0, "Variation_Binders returns correct value -- with args", actual)
}

func Test_Variation_EnumMetadata(t *testing.T) {
	// Arrange
	v := ostype.Windows

	// Act
	actual := args.Map{
		"enumTypeNil":       v.EnumType() == nil,
		"allNameValuesLen":  len(v.AllNameValues()) > 0,
		"integerRangesLen":  len(v.IntegerEnumRanges()) > 0,
		"rangesDynMapLen":   len(v.RangesDynamicMap()) > 0,
		"formatNotEmpty":    v.Format("{name}") != "",
		"minMaxAnyNotNil":   true,
		"minValueStr":       v.MinValueString() != "",
		"maxValueStr":       v.MaxValueString() != "",
		"maxByte":           int(v.MaxByte()) > 0,
		"rangesByteLen":     len(v.RangesByte()) > 0,
	}
	expected := args.Map{
		"enumTypeNil":       false,
		"allNameValuesLen":  true,
		"integerRangesLen":  true,
		"rangesDynMapLen":   true,
		"formatNotEmpty":    true,
		"minMaxAnyNotNil":   true,
		"minValueStr":       true,
		"maxValueStr":       true,
		"maxByte":           true,
		"rangesByteLen":     true,
	}
	expected.ShouldBeEqual(t, 0, "Variation_EnumMetadata returns correct value -- with args", actual)
}

func Test_Variation_UnmarshallEnumToValue(t *testing.T) {
	// Arrange
	data, _ := json.Marshal(ostype.Windows)

	// Act
	val, err := ostype.Windows.UnmarshallEnumToValue(data)

	// Assert
	actual := args.Map{
		"noError":  err == nil,
		"valMatch": val == byte(ostype.Windows),
	}
	expected := args.Map{
		"noError":  true,
		"valMatch": true,
	}
	expected.ShouldBeEqual(t, 0, "Variation_UnmarshallEnumToValue returns correct value -- with args", actual)
}

// ── Group enum methods ──

func Test_Group_EnumMethods_FromVariationEnumMethods(t *testing.T) {
	// Arrange
	g := ostype.UnixGroup

	// Act
	actual := args.Map{
		"name":          g.Name(),
		"nameValue":     g.NameValue(),
		"typeName":      g.TypeName(),
		"isValid":       g.IsValid(),
		"isInvalid":     g.IsInvalid(),
		"isUnix":        g.IsUnix(),
		"isWindows":     g.IsWindows(),
		"isAndroid":     g.IsAndroid(),
		"isInvalidGrp":  g.IsInvalidGroup(),
		"is":            g.Is(ostype.UnixGroup),
		"byte":          int(g.Byte()),
		"stringVal":     g.String(),
	}
	expected := args.Map{
		"name":          "UnixGroup",
		"nameValue":     g.NameValue(),
		"typeName":      g.TypeName(),
		"isValid":       true,
		"isInvalid":     false,
		"isUnix":        true,
		"isWindows":     false,
		"isAndroid":     false,
		"isInvalidGrp":  false,
		"is":            true,
		"byte":          int(ostype.UnixGroup),
		"stringVal":     "UnixGroup",
	}
	expected.ShouldBeEqual(t, 0, "Group_EnumMethods returns correct value -- with args", actual)
}

func Test_Group_JSON_FromVariationEnumMethods(t *testing.T) {
	// Arrange
	g := ostype.WindowsGroup

	// Act
	data, _ := json.Marshal(g)
	var parsed ostype.Group
	err := json.Unmarshal(data, &parsed)

	// Assert
	actual := args.Map{
		"noError":    err == nil,
		"roundTrip":  parsed == g,
	}
	expected := args.Map{
		"noError":    true,
		"roundTrip":  true,
	}
	expected.ShouldBeEqual(t, 0, "Group_JSON returns correct value -- with args", actual)
}

func Test_Group_Binders_FromVariationEnumMethods(t *testing.T) {
	// Arrange
	g := ostype.UnixGroup

	// Act
	actual := args.Map{
		"enumContractsNil": g.AsBasicEnumContractsBinder() == nil,
		"jsonContractsNil": g.AsJsonContractsBinder() == nil,
		"byteContractsNil": g.AsBasicByteEnumContractsBinder() == nil,
		"toPtrNil":         g.ToPtr() == nil,
	}
	expected := args.Map{
		"enumContractsNil": false,
		"jsonContractsNil": false,
		"byteContractsNil": false,
		"toPtrNil":         false,
	}
	expected.ShouldBeEqual(t, 0, "Group_Binders returns correct value -- with args", actual)
}

// ── GetCurrentVariant / GetCurrentGroup / GetGroupVariant ──

func Test_GetCurrentVariant_FromVariationEnumMethods(t *testing.T) {
	// Act
	v := ostype.GetCurrentVariant()

	// Assert
	actual := args.Map{
		"isValid": v.IsValid(),
	}
	expected := args.Map{
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "GetCurrentVariant returns correct value -- with args", actual)
}

func Test_GetCurrentGroup_FromVariationEnumMethods(t *testing.T) {
	// Act
	g := ostype.GetCurrentGroup()

	// Assert
	actual := args.Map{
		"isValid": g.IsValid(),
	}
	expected := args.Map{
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "GetCurrentGroup returns correct value -- with args", actual)
}

func Test_GetGroupVariant_FromVariationEnumMethods(t *testing.T) {
	// Act
	gv := ostype.GetGroupVariant()

	// Assert
	actual := args.Map{
		"groupValid":     gv.Group.IsValid(),
		"variationValid": gv.Variation.IsValid(),
	}
	expected := args.Map{
		"groupValid":     true,
		"variationValid": true,
	}
	expected.ShouldBeEqual(t, 0, "GetGroupVariant returns correct value -- with args", actual)
}

func Test_GetGroupVariantPtr_FromVariationEnumMethods(t *testing.T) {
	// Act
	gv := ostype.GetGroupVariantPtr()

	// Assert
	actual := args.Map{
		"notNil": gv != nil,
	}
	expected := args.Map{
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "GetGroupVariantPtr returns correct value -- with args", actual)
}
