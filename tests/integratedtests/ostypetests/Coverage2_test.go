package ostypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/ostype"
)

// ── Group methods ──

func Test_Cov2_Group_BasicMethods(t *testing.T) {
	g := ostype.UnixGroup

	actual := args.Map{
		"isWindows":   g.IsWindows(),
		"isUnix":      g.IsUnix(),
		"isAndroid":   g.IsAndroid(),
		"isInvalid":   g.IsInvalidGroup(),
		"isValid":     g.IsValid(),
		"byte":        g.Byte(),
		"value":       g.Value(),
		"valueByte":   g.ValueByte(),
		"valueInt":    g.ValueInt(),
		"valueInt8":   g.ValueInt8(),
		"valueInt16":  g.ValueInt16(),
		"valueInt32":  g.ValueInt32(),
		"valueUInt16": g.ValueUInt16(),
		"is":          g.Is(ostype.UnixGroup),
		"isNot":       g.Is(ostype.WindowsGroup),
		"nameNotEmp":  g.Name() != "",
		"strNotEmp":   g.String() != "",
		"nameValNE":   g.NameValue() != "",
		"numStrNE":    g.ToNumberString() != "",
		"valStrNE":    g.ValueString() != "",
		"typeNameNE":  g.TypeName() != "",
		"rangesCsvNE": g.RangeNamesCsv() != "",
	}
	expected := args.Map{
		"isWindows":   false,
		"isUnix":      true,
		"isAndroid":   false,
		"isInvalid":   false,
		"isValid":     true,
		"byte":        byte(ostype.UnixGroup),
		"value":       byte(ostype.UnixGroup),
		"valueByte":   byte(ostype.UnixGroup),
		"valueInt":    int(ostype.UnixGroup),
		"valueInt8":   int8(ostype.UnixGroup),
		"valueInt16":  int16(ostype.UnixGroup),
		"valueInt32":  int32(ostype.UnixGroup),
		"valueUInt16": uint16(ostype.UnixGroup),
		"is":          true,
		"isNot":       false,
		"nameNotEmp":  true,
		"strNotEmp":   true,
		"nameValNE":   true,
		"numStrNE":    true,
		"valStrNE":    true,
		"typeNameNE":  true,
		"rangesCsvNE": true,
	}
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- basic methods", actual)
}

func Test_Cov2_Group_EnumMethods(t *testing.T) {
	g := ostype.UnixGroup

	actual := args.Map{
		"rangesLen":    len(g.RangesByte()) > 0,
		"minByte":      g.MinByte() == 0,
		"maxByte":      g.MaxByte() > 0,
		"allNamesLen":  len(g.AllNameValues()) > 0,
		"enumType":     g.EnumType() != nil,
		"isByteEq":     g.IsByteValueEqual(byte(ostype.UnixGroup)),
		"isNameEq":     g.IsNameEqual(g.Name()),
		"isValEq":      g.IsValueEqual(byte(ostype.UnixGroup)),
		"formatNE":     g.Format("%s") != "",
		"toPtr":        g.ToPtr() != nil,
	}
	expected := args.Map{
		"rangesLen":    true,
		"minByte":      true,
		"maxByte":      true,
		"allNamesLen":  true,
		"enumType":     true,
		"isByteEq":     true,
		"isNameEq":     true,
		"isValEq":      true,
		"formatNE":     true,
		"toPtr":        true,
	}
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- enum methods", actual)
}

func Test_Cov2_Group_AnyMethods(t *testing.T) {
	g := ostype.UnixGroup

	actual := args.Map{
		"isAnyNames":     g.IsAnyNamesOf(g.Name(), "Invalid"),
		"isAnyNamesNo":   g.IsAnyNamesOf("Invalid"),
		"isAnyValues":    g.IsAnyValuesEqual(byte(ostype.UnixGroup)),
		"isAnyValuesNo":  g.IsAnyValuesEqual(99),
	}
	expected := args.Map{
		"isAnyNames":     true,
		"isAnyNamesNo":   false,
		"isAnyValues":    true,
		"isAnyValuesNo":  false,
	}
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- any methods", actual)
}

func Test_Cov2_Group_JSON(t *testing.T) {
	g := ostype.UnixGroup
	bytes, err := g.MarshalJSON()
	var g2 ostype.Group
	err2 := g2.UnmarshalJSON(bytes)

	actual := args.Map{
		"marshalErr":   err == nil,
		"unmarshalErr": err2 == nil,
		"same":         g2 == ostype.UnixGroup,
	}
	expected := args.Map{
		"marshalErr":   true,
		"unmarshalErr": true,
		"same":         true,
	}
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- JSON", actual)
}

func Test_Cov2_Group_Binders(t *testing.T) {
	g := ostype.UnixGroup
	actual := args.Map{
		"basicBinder":    g.AsBasicEnumContractsBinder() != nil,
		"jsonBinder":     g.AsJsonContractsBinder() != nil,
		"byteBinder":     g.AsBasicByteEnumContractsBinder() != nil,
	}
	expected := args.Map{
		"basicBinder":    true,
		"jsonBinder":     true,
		"byteBinder":     true,
	}
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- binders", actual)
}

// ── Variation methods ──

func Test_Cov2_Variation_BasicMethods(t *testing.T) {
	v := ostype.Linux

	actual := args.Map{
		"isWindows":   v.IsWindows(),
		"isLinux":     v.IsLinux(),
		"isDarwin":    v.IsDarwinOrMacOs(),
		"isJs":        v.IsJavaScript(),
		"isFreeBsd":   v.IsFreeBsd(),
		"isNetBsd":    v.IsNetBsd(),
		"isOpenBsd":   v.IsOpenBsd(),
		"isDragonFly": v.IsDragonFly(),
		"isValid":     v.IsValid(),
		"isInvalid":   v.IsInvalid(),
		"isAnyOs":     v.IsAnyOperatingSystem(),
		"isLinuxMac":  v.IsLinuxOrMac(),
		"isPossUnix":  v.IsPossibleUnixGroup(),
		"isActualUnix": v.IsActualGroupUnix(),
		"group":       v.Group() == ostype.UnixGroup,
	}
	expected := args.Map{
		"isWindows":   false,
		"isLinux":     true,
		"isDarwin":    false,
		"isJs":        false,
		"isFreeBsd":   false,
		"isNetBsd":    false,
		"isOpenBsd":   false,
		"isDragonFly": false,
		"isValid":     true,
		"isInvalid":   false,
		"isAnyOs":     false,
		"isLinuxMac":  true,
		"isPossUnix":  true,
		"isActualUnix": true,
		"group":       true,
	}
	expected.ShouldBeEqual(t, 0, "Variation returns correct value -- basic methods", actual)
}

func Test_Cov2_Variation_MatchMethods(t *testing.T) {
	v := ostype.Linux

	actual := args.Map{
		"is":           v.Is(ostype.Linux),
		"isNot":        v.Is(ostype.Windows),
		"isByte":       v.IsByte(byte(ostype.Linux)),
		"isAnyMatch":   v.IsAnyMatch(ostype.Windows, ostype.Linux),
		"isAnyMatchNo": v.IsAnyMatch(ostype.Windows),
		"isStrMatch":   v.IsStringsMatchAny("linux"),
		"isStrMatchNo": v.IsStringsMatchAny("windows"),
	}
	expected := args.Map{
		"is":           true,
		"isNot":        false,
		"isByte":       true,
		"isAnyMatch":   true,
		"isAnyMatchNo": false,
		"isStrMatch":   true,
		"isStrMatchNo": false,
	}
	expected.ShouldBeEqual(t, 0, "Variation returns correct value -- match methods", actual)
}

func Test_Cov2_Variation_GroupForWindows(t *testing.T) {
	actual := args.Map{"group": ostype.Windows.Group() == ostype.WindowsGroup}
	expected := args.Map{"group": true}
	expected.ShouldBeEqual(t, 0, "Variation returns correct value -- group windows", actual)
}

func Test_Cov2_Variation_GroupForAndroid(t *testing.T) {
	actual := args.Map{"group": ostype.Android.Group() == ostype.AndroidGroup}
	expected := args.Map{"group": true}
	expected.ShouldBeEqual(t, 0, "Variation returns correct value -- group android", actual)
}

// ── GetGroup / GetVariant / GetGroupVariant ──

func Test_Cov2_GetGroup_Windows(t *testing.T) {
	actual := args.Map{"result": ostype.GetGroup("windows") == ostype.WindowsGroup}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetGroup returns correct value -- windows", actual)
}

func Test_Cov2_GetGroup_Android(t *testing.T) {
	actual := args.Map{"result": ostype.GetGroup("android") == ostype.AndroidGroup}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetGroup returns correct value -- android", actual)
}

func Test_Cov2_GetGroup_Linux(t *testing.T) {
	actual := args.Map{"result": ostype.GetGroup("linux") == ostype.UnixGroup}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetGroup returns correct value -- linux", actual)
}

func Test_Cov2_GetGroup_Invalid(t *testing.T) {
	actual := args.Map{"result": ostype.GetGroup("invalid-os") == ostype.InvalidGroup}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetGroup returns error -- invalid", actual)
}

func Test_Cov2_GetVariant_Linux(t *testing.T) {
	actual := args.Map{"result": ostype.GetVariant("linux") == ostype.Linux}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetVariant returns correct value -- linux", actual)
}

func Test_Cov2_GetVariant_Unknown(t *testing.T) {
	actual := args.Map{"result": ostype.GetVariant("nonexistent") == ostype.Unknown}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetVariant returns correct value -- unknown", actual)
}

func Test_Cov2_GetGroupVariant(t *testing.T) {
	gv := ostype.GetGroupVariant()
	actual := args.Map{"groupValid": gv.Group.IsValid() || gv.Group == ostype.InvalidGroup}
	expected := args.Map{"groupValid": true}
	expected.ShouldBeEqual(t, 0, "GetGroupVariant returns correct value -- with args", actual)
}

func Test_Cov2_GetGroupVariantPtr(t *testing.T) {
	gv := ostype.GetGroupVariantPtr()
	actual := args.Map{"notNil": gv != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetGroupVariantPtr returns correct value -- with args", actual)
}

func Test_Cov2_GetCurrentGroup(t *testing.T) {
	g := ostype.GetCurrentGroup()
	actual := args.Map{"valid": g != ostype.InvalidGroup}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "GetCurrentGroup returns correct value -- with args", actual)
}

func Test_Cov2_GetCurrentVariant(t *testing.T) {
	v := ostype.GetCurrentVariant()
	actual := args.Map{"valid": v != ostype.Unknown}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "GetCurrentVariant returns correct value -- with args", actual)
}

// ── Variation JSON ──

func Test_Cov2_Variation_JSON(t *testing.T) {
	v := ostype.Linux
	bytes, err := v.MarshalJSON()
	var v2 ostype.Variation
	err2 := v2.UnmarshalJSON(bytes)

	actual := args.Map{
		"marshalErr":   err == nil,
		"unmarshalErr": err2 == nil,
		"same":         v2 == ostype.Linux,
	}
	expected := args.Map{
		"marshalErr":   true,
		"unmarshalErr": true,
		"same":         true,
	}
	expected.ShouldBeEqual(t, 0, "Variation returns correct value -- JSON", actual)
}

func Test_Cov2_Variation_Binders(t *testing.T) {
	v := ostype.Linux
	actual := args.Map{
		"basic": v.AsBasicEnumContractsBinder() != nil,
		"json":  v.AsJsonContractsBinder() != nil,
		"byte":  v.AsBasicByteEnumContractsBinder() != nil,
		"toPtr": v.ToPtr() != nil,
	}
	expected := args.Map{
		"basic": true,
		"json":  true,
		"byte":  true,
		"toPtr": true,
	}
	expected.ShouldBeEqual(t, 0, "Variation returns correct value -- binders", actual)
}
