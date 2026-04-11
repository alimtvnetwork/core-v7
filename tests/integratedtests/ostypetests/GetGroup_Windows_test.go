package ostypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/osconsts"
	"github.com/alimtvnetwork/core/ostype"
)

// ── GetGroup ──

func Test_GetGroup_Windows(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.GetGroup(osconsts.Windows).Name()}

	// Assert
	expected := args.Map{"result": "WindowsGroup"}
	expected.ShouldBeEqual(t, 0, "GetGroup returns correct value -- Windows", actual)
}

func Test_GetGroup_Android(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.GetGroup("android").Name()}

	// Assert
	expected := args.Map{"result": "AndroidGroup"}
	expected.ShouldBeEqual(t, 0, "GetGroup returns correct value -- Android", actual)
}

func Test_GetGroup_Linux(t *testing.T) {
	// Act
	actual := args.Map{"isUnix": ostype.GetGroup("linux").IsUnix()}

	// Assert
	expected := args.Map{"isUnix": true}
	expected.ShouldBeEqual(t, 0, "GetGroup returns correct value -- Linux", actual)
}

func Test_GetGroup_Invalid(t *testing.T) {
	// Act
	actual := args.Map{"isInvalid": ostype.GetGroup("fakeos").IsInvalidGroup()}

	// Assert
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "GetGroup returns error -- invalid", actual)
}

// ── GetVariant ──

func Test_GetVariant_Windows(t *testing.T) {
	// Act
	actual := args.Map{"isWindows": ostype.GetVariant("windows").IsWindows()}

	// Assert
	expected := args.Map{"isWindows": true}
	expected.ShouldBeEqual(t, 0, "GetVariant returns correct value -- Windows", actual)
}

func Test_GetVariant_Unknown(t *testing.T) {
	// Act
	actual := args.Map{"isUnknown": ostype.GetVariant("fakeos").Is(ostype.Unknown)}

	// Assert
	expected := args.Map{"isUnknown": true}
	expected.ShouldBeEqual(t, 0, "GetVariant returns correct value -- unknown", actual)
}

// ── GetGroupVariant / GetGroupVariantPtr ──

func Test_GetGroupVariant(t *testing.T) {
	// Arrange
	gv := ostype.GetGroupVariant()

	// Act
	actual := args.Map{"groupDefined": gv.Group.IsValid() || gv.Group.IsInvalidGroup()}

	// Assert
	expected := args.Map{"groupDefined": true}
	expected.ShouldBeEqual(t, 0, "GetGroupVariant returns correct value -- with args", actual)
}

func Test_GetGroupVariantPtr(t *testing.T) {
	// Arrange
	gv := ostype.GetGroupVariantPtr()

	// Act
	actual := args.Map{"notNil": gv != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetGroupVariantPtr returns correct value -- with args", actual)
}

// ── GetCurrentGroup / GetCurrentVariant ──

func Test_GetCurrentGroup(t *testing.T) {
	// Arrange
	g := ostype.GetCurrentGroup()

	// Act
	actual := args.Map{"nameNotEmpty": g.Name() != ""}

	// Assert
	expected := args.Map{"nameNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetCurrentGroup returns correct value -- with args", actual)
}

func Test_GetCurrentVariant(t *testing.T) {
	// Arrange
	v := ostype.GetCurrentVariant()

	// Act
	actual := args.Map{"nameNotEmpty": v.Name() != ""}

	// Assert
	expected := args.Map{"nameNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetCurrentVariant returns correct value -- with args", actual)
}

// ── Group enum methods ──

func Test_Group_AllMethods(t *testing.T) {
	// Arrange
	g := ostype.WindowsGroup

	// Act
	actual := args.Map{
		"name":            g.Name(),
		"nameValue":       g.NameValue() != "",
		"typeName":        g.TypeName() != "",
		"rangeNamesCsv":   g.RangeNamesCsv() != "",
		"toNumberString":  g.ToNumberString(),
		"allNameValues":   len(g.AllNameValues()) > 0,
		"integerRanges":   len(g.IntegerEnumRanges()) > 0,
		"rangesDynamic":   len(g.RangesDynamicMap()) > 0,
		"isNameEqual":     g.IsNameEqual("WindowsGroup"),
		"isAnyNamesOf":    g.IsAnyNamesOf("WindowsGroup", "UnixGroup"),
		"isValueEqual":    g.IsValueEqual(byte(ostype.WindowsGroup)),
		"isAnyValsEqual":  g.IsAnyValuesEqual(byte(ostype.WindowsGroup)),
		"isByteValEqual":  g.IsByteValueEqual(byte(ostype.WindowsGroup)),
		"format":          g.Format("%s") != "",
		"valueInt":        g.ValueInt(),
		"valueInt8":       g.ValueInt8(),
		"valueInt16":      g.ValueInt16(),
		"valueInt32":      g.ValueInt32(),
		"valueString":     g.ValueString(),
		"valueUInt16":     g.ValueUInt16(),
		"value":           g.Value(),
		"valueByte":       g.ValueByte(),
		"maxByte":         g.MaxByte() > 0,
		"minByte":         g.MinByte(),
		"rangesByte":      len(g.RangesByte()) > 0,
		"isValid":         g.IsValid(),
		"isInvalid":       g.IsInvalid(),
		"string":          g.String(),
		"isWindows":       g.IsWindows(),
		"isUnix":          g.IsUnix(),
		"isAndroid":       g.IsAndroid(),
		"isInvalidGroup":  g.IsInvalidGroup(),
		"is":              g.Is(ostype.WindowsGroup),
		"toPtr":           g.ToPtr() != nil,
	}

	// Assert
	expected := args.Map{
		"name": "WindowsGroup", "nameValue": true, "typeName": true,
		"rangeNamesCsv": true, "toNumberString": "0", "allNameValues": true,
		"integerRanges": true, "rangesDynamic": true,
		"isNameEqual": true, "isAnyNamesOf": true,
		"isValueEqual": true, "isAnyValsEqual": true, "isByteValEqual": true,
		"format": true, "valueInt": 0, "valueInt8": int8(0),
		"valueInt16": int16(0), "valueInt32": int32(0),
		"valueString": "0", "valueUInt16": uint16(0),
		"value": byte(0), "valueByte": byte(0),
		"maxByte": true, "minByte": byte(0), "rangesByte": true,
		"isValid": true, "isInvalid": false, "string": "WindowsGroup",
		"isWindows": true, "isUnix": false, "isAndroid": false,
		"isInvalidGroup": false, "is": true, "toPtr": true,
	}
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- all methods", actual)
}

func Test_Group_EnumEqual(t *testing.T) {
	// Arrange
	g := ostype.WindowsGroup
	u := ostype.UnixGroup
	up := &u
	wp := &g

	// Act
	actual := args.Map{
		"isEnumEqual":     g.IsEnumEqual(wp),
		"isAnyEnumsEqual": g.IsAnyEnumsEqual(up, wp),
	}

	// Assert
	expected := args.Map{
		"isEnumEqual": true,
		"isAnyEnumsEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- EnumEqual", actual)
}

func Test_Group_MarshalUnmarshalJSON(t *testing.T) {
	// Arrange
	g := ostype.WindowsGroup
	data, _ := g.MarshalJSON()
	var g2 ostype.Group
	_ = g2.UnmarshalJSON(data)

	// Act
	actual := args.Map{
		"notEmpty": len(data) > 0,
		"name": g2.Name(),
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"name": "WindowsGroup",
	}
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- MarshalUnmarshalJSON", actual)
}

func Test_Group_MinMaxAny(t *testing.T) {
	// Arrange
	g := ostype.WindowsGroup
	min, max := g.MinMaxAny()

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
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- MinMaxAny", actual)
}

func Test_Group_MinMaxValueString(t *testing.T) {
	// Arrange
	g := ostype.WindowsGroup

	// Act
	actual := args.Map{
		"minStr": g.MinValueString() != "",
		"maxStr": g.MaxValueString() != "",
		"minInt": g.MinInt() >= 0,
		"maxInt": g.MaxInt() > 0,
	}

	// Assert
	expected := args.Map{
		"minStr": true,
		"maxStr": true,
		"minInt": true,
		"maxInt": true,
	}
	expected.ShouldBeEqual(t, 0, "Group returns non-empty -- MinMaxValueString", actual)
}

func Test_Group_OnlySupportedErr(t *testing.T) {
	// Arrange
	err := ostype.WindowsGroup.OnlySupportedErr("WindowsGroup")
	noErr := err == nil

	// Act
	actual := args.Map{"noErr": noErr}

	// Assert
	expected := args.Map{"noErr": noErr}
	expected.ShouldBeEqual(t, 0, "Group returns error -- OnlySupportedErr", actual)
}

func Test_Group_OnlySupportedMsgErr(t *testing.T) {
	// Arrange
	err := ostype.InvalidGroup.OnlySupportedMsgErr("msg", "WindowsGroup")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Group returns error -- OnlySupportedMsgErr", actual)
}

func Test_Group_EnumType(t *testing.T) {
	// Act
	actual := args.Map{"notNil": ostype.WindowsGroup.EnumType() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- EnumType", actual)
}

func Test_Group_Contracts(t *testing.T) {
	// Arrange
	g := ostype.WindowsGroup

	// Act
	actual := args.Map{
		"basicEnumBinder":     g.AsBasicEnumContractsBinder() != nil,
		"jsonBinder":          g.AsJsonContractsBinder() != nil,
		"basicByteEnumBinder": g.AsBasicByteEnumContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"basicEnumBinder": true,
		"jsonBinder": true,
		"basicByteEnumBinder": true,
	}
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- contract binders", actual)
}

// ── Variation extended ──

func Test_Variation_Extended(t *testing.T) {
	// Arrange
	v := ostype.Linux
	lp := &v

	// Act
	actual := args.Map{
		"isAnyMatch":       v.IsAnyMatch(ostype.Windows, ostype.Linux),
		"isAnyMatchFalse":  v.IsAnyMatch(ostype.Windows),
		"isPossibleUnix":   v.IsPossibleUnixGroup(),
		"group":            v.Group().Name(),
		"isStringsMatchAny": v.IsStringsMatchAny("linux", "windows"),
		"allNameValues":    len(v.AllNameValues()) > 0,
		"integerRanges":    len(v.IntegerEnumRanges()) > 0,
		"rangesDynamic":    len(v.RangesDynamicMap()) > 0,
		"format":           v.Format("%s") != "",
		"isEnumEqual":      v.IsEnumEqual(lp),
		"isAnyEnumsEqual":  v.IsAnyEnumsEqual(lp),
		"minMaxStr":        v.MinValueString() != "",
		"maxStr":           v.MaxValueString() != "",
		"minInt":           v.MinInt() >= 0,
		"maxInt":           v.MaxInt() > 0,
		"enumType":         v.EnumType() != nil,
	}

	// Assert
	expected := args.Map{
		"isAnyMatch": true, "isAnyMatchFalse": false,
		"isPossibleUnix": true, "group": "UnixGroup",
		"isStringsMatchAny": true, "allNameValues": true,
		"integerRanges": true, "rangesDynamic": true, "format": true,
		"isEnumEqual": true, "isAnyEnumsEqual": true,
		"minMaxStr": true, "maxStr": true, "minInt": true, "maxInt": true,
		"enumType": true,
	}
	expected.ShouldBeEqual(t, 0, "Variation returns correct value -- extended", actual)
}

func Test_Variation_WindowsGroup(t *testing.T) {
	// Act
	actual := args.Map{
		"group":       ostype.Windows.Group().Name(),
		"isPossUnix":  ostype.Windows.IsPossibleUnixGroup(),
	}

	// Assert
	expected := args.Map{
		"group": "WindowsGroup",
		"isPossUnix": false,
	}
	expected.ShouldBeEqual(t, 0, "Variation returns correct value -- Windows group", actual)
}

func Test_Variation_AndroidGroup(t *testing.T) {
	// Act
	actual := args.Map{"group": ostype.Android.Group().Name()}

	// Assert
	expected := args.Map{"group": "AndroidGroup"}
	expected.ShouldBeEqual(t, 0, "Variation returns correct value -- Android group", actual)
}

func Test_Variation_OnlySupportedErr(t *testing.T) {
	// Arrange
	err := ostype.Linux.OnlySupportedErr("linux")
	noErr := err == nil

	// Act
	actual := args.Map{"noErr": noErr}

	// Assert
	expected := args.Map{"noErr": noErr}
	expected.ShouldBeEqual(t, 0, "Variation returns error -- OnlySupportedErr", actual)
}

func Test_Variation_OnlySupportedMsgErr(t *testing.T) {
	// Arrange
	err := ostype.Unknown.OnlySupportedMsgErr("msg", "linux")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Variation returns error -- OnlySupportedMsgErr", actual)
}

func Test_Variation_MinMaxAny(t *testing.T) {
	// Arrange
	min, max := ostype.Linux.MinMaxAny()

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
	expected.ShouldBeEqual(t, 0, "Variation returns correct value -- MinMaxAny", actual)
}

func Test_Variation_Contracts(t *testing.T) {
	// Arrange
	v := ostype.Linux

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
	expected.ShouldBeEqual(t, 0, "Variation returns correct value -- contract binders", actual)
}

func Test_Group_UnmarshalJSON_Invalid(t *testing.T) {
	// Arrange
	var g ostype.Group
	err := g.UnmarshalJSON([]byte("invalid"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Group returns error -- UnmarshalJSON invalid", actual)
}

func Test_Variation_UnmarshalJSON_Invalid(t *testing.T) {
	// Arrange
	var v ostype.Variation
	err := v.UnmarshalJSON([]byte("invalid"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Variation returns error -- UnmarshalJSON invalid", actual)
}
