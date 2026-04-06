package chmodclasstypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Variant identity checks ──

func Test_Cov_Variant_Identity(t *testing.T) {
	actual := args.Map{
		"invalidIsInvalid": chmodclasstype.Invalid.IsInvalid(),
		"invalidIsValid":   chmodclasstype.Invalid.IsValid(),
		"allIsAll":         chmodclasstype.All.IsAll(),
		"ownerIsOwner":     chmodclasstype.Owner.IsOwner(),
		"groupIsGroup":     chmodclasstype.Group.IsGroup(),
		"otherIsOther":     chmodclasstype.Other.IsOther(),
		"ogIsOwnerGroup":   chmodclasstype.OwnerGroup.IsOwnerGroup(),
		"goIsGroupOther":   chmodclasstype.GroupOther.IsGroupOther(),
		"ooIsOwnerOther":   chmodclasstype.OwnerOther.IsOwnerOther(),
		"uninit":           chmodclasstype.Invalid.IsUnInitialized(),
	}
	expected := args.Map{
		"invalidIsInvalid": true,
		"invalidIsValid":   false,
		"allIsAll":         true,
		"ownerIsOwner":     true,
		"groupIsGroup":     true,
		"otherIsOther":     true,
		"ogIsOwnerGroup":   true,
		"goIsGroupOther":   true,
		"ooIsOwnerOther":   true,
		"uninit":           true,
	}
	expected.ShouldBeEqual(t, 0, "Variant identity checks -- all variants", actual)
}

// ── Variant value methods ──

func Test_Cov_Variant_Values(t *testing.T) {
	v := chmodclasstype.Owner
	actual := args.Map{
		"valueByte":   int(v.Value()),
		"valueInt":    v.ValueInt(),
		"valueInt8":   int(v.ValueInt8()),
		"valueInt16":  int(v.ValueInt16()),
		"valueInt32":  int(v.ValueInt32()),
		"valueUInt16": int(v.ValueUInt16()),
		"valueString": v.ValueString() != "",
		"valueByteFn": int(v.ValueByte()),
	}
	expected := args.Map{
		"valueByte":   int(chmodclasstype.Owner),
		"valueInt":    int(chmodclasstype.Owner),
		"valueInt8":   int(chmodclasstype.Owner),
		"valueInt16":  int(chmodclasstype.Owner),
		"valueInt32":  int(chmodclasstype.Owner),
		"valueUInt16": int(chmodclasstype.Owner),
		"valueString": true,
		"valueByteFn": int(chmodclasstype.Owner),
	}
	expected.ShouldBeEqual(t, 0, "Variant value methods -- Owner", actual)
}

// ── Variant name methods ──

func Test_Cov_Variant_Names(t *testing.T) {
	owner := chmodclasstype.Owner
	actual := args.Map{
		"ownerName":   (&owner).Name(),
		"ownerString": owner.String(),
		"nameValue":   owner.NameValue() != "",
		"typeName":    owner.TypeName() != "",
	}
	expected := args.Map{
		"ownerName":   "Owner",
		"ownerString": "Owner",
		"nameValue":   true,
		"typeName":    true,
	}
	expected.ShouldBeEqual(t, 0, "Variant name methods -- Owner", actual)
}

// ── Variant comparison methods ──

func Test_Cov_Variant_Comparison(t *testing.T) {
	v := chmodclasstype.Owner
	ownerEnum := chmodclasstype.Owner
	actual := args.Map{
		"isNameEqual":    v.IsNameEqual("Owner"),
		"isNameNotEqual": v.IsNameEqual("Group"),
		"isByteEqual":    v.IsByteValueEqual(byte(chmodclasstype.Owner)),
		"isValueEqual":   v.IsValueEqual(byte(chmodclasstype.Owner)),
		"isEnumEqual":    v.IsEnumEqual(&ownerEnum),
		"isAnyNames":     v.IsAnyNamesOf("Owner", "Group"),
		"isAnyNamesFail": v.IsAnyNamesOf("Group", "Other"),
		"isAnyValues":    v.IsAnyValuesEqual(byte(chmodclasstype.Owner), byte(chmodclasstype.Group)),
		"isAnyValFail":   v.IsAnyValuesEqual(byte(chmodclasstype.Group)),
	}
	expected := args.Map{
		"isNameEqual":    true,
		"isNameNotEqual": false,
		"isByteEqual":    true,
		"isValueEqual":   true,
		"isEnumEqual":    true,
		"isAnyNames":     true,
		"isAnyNamesFail": false,
		"isAnyValues":    true,
		"isAnyValFail":   false,
	}
	expected.ShouldBeEqual(t, 0, "Variant comparison methods -- Owner", actual)
}

// ── Variant IsAnyEnumsEqual ──

func Test_Cov_Variant_IsAnyEnumsEqual(t *testing.T) {
	v := chmodclasstype.Owner
	group := chmodclasstype.Group
	owner := chmodclasstype.Owner
	other := chmodclasstype.Other
	actual := args.Map{
		"match":   v.IsAnyEnumsEqual(&group, &owner),
		"noMatch": v.IsAnyEnumsEqual(&group, &other),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Variant IsAnyEnumsEqual -- match and no match", actual)
}

// ── Variant enum metadata ──

func Test_Cov_Variant_Metadata(t *testing.T) {
	v := chmodclasstype.Owner
	actual := args.Map{
		"allNamesLen": len(v.AllNameValues()) > 0,
		"rangesLen":   len(v.IntegerEnumRanges()) > 0,
		"rangesCsv":   v.RangeNamesCsv() != "",
		"rangesMap":   len(v.RangesDynamicMap()) > 0,
		"rangesByte":  len(v.RangesByte()) > 0,
		"maxByte":     int(v.MaxByte()) > 0,
		"minByte":     int(v.MinByte()) == 0,
		"maxInt":      v.MaxInt() > 0,
		"minInt":      v.MinInt() == 0,
		"maxStr":      v.MaxValueString() != "",
		"minStr":      v.MinValueString() != "",
		"format":      v.Format("%s") != "",
		"enumType":    v.EnumType() != nil,
	}
	expected := args.Map{
		"allNamesLen": true,
		"rangesLen":   true,
		"rangesCsv":   true,
		"rangesMap":   true,
		"rangesByte":  true,
		"maxByte":     true,
		"minByte":     true,
		"maxInt":      true,
		"minInt":      true,
		"maxStr":      true,
		"minStr":      true,
		"format":      true,
		"enumType":    true,
	}
	expected.ShouldBeEqual(t, 0, "Variant metadata methods -- Owner", actual)
}

// ── Variant MinMaxAny ──

func Test_Cov_Variant_MinMaxAny(t *testing.T) {
	v := chmodclasstype.Owner
	min, max := v.MinMaxAny()
	actual := args.Map{
		"minNotNil": min != nil,
		"maxNotNil": max != nil,
	}
	expected := args.Map{
		"minNotNil": true,
		"maxNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant MinMaxAny returns non-nil -- Owner", actual)
}

// ── Variant JSON ──

func Test_Cov_Variant_MarshalJSON(t *testing.T) {
	v := chmodclasstype.Owner
	bytes, err := v.MarshalJSON()
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"noErr":    err == nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"noErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "Variant MarshalJSON returns bytes -- Owner", actual)
}

func Test_Cov_Variant_UnmarshalJSON(t *testing.T) {
	v := chmodclasstype.Owner
	bytes, _ := v.MarshalJSON()

	var target chmodclasstype.Variant
	err := target.UnmarshalJSON(bytes)
	actual := args.Map{
		"noErr":   err == nil,
		"isOwner": target.IsOwner(),
	}
	expected := args.Map{
		"noErr":   true,
		"isOwner": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant UnmarshalJSON roundtrip -- Owner", actual)
}

func Test_Cov_Variant_UnmarshalJSON_Invalid(t *testing.T) {
	var target chmodclasstype.Variant
	err := target.UnmarshalJSON([]byte(`"invalid_enum_name"`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Variant UnmarshalJSON returns error -- invalid name", actual)
}

// ── Variant OnlySupportedErr ──

func Test_Cov_Variant_OnlySupportedErr(t *testing.T) {
	v := chmodclasstype.Owner
	// OnlySupportedErr checks ALL enum names against supported list.
	// Passing a subset means unsupported names exist → error returned.
	err := v.OnlySupportedErr("Owner", "Group")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Variant OnlySupportedErr returns error -- subset supported", actual)
}

func Test_Cov_Variant_OnlySupportedMsgErr(t *testing.T) {
	v := chmodclasstype.Owner
	// Same: passing a subset means error is returned.
	err := v.OnlySupportedMsgErr("test message", "Owner")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Variant OnlySupportedMsgErr returns error -- subset supported", actual)
}

// ── Variant AsContractsBinder ──

func Test_Cov_Variant_AsContractsBinder(t *testing.T) {
	v := chmodclasstype.Owner
	actual := args.Map{
		"basicBinder":     v.AsBasicEnumContractsBinder() != nil,
		"basicByteBinder": v.AsBasicByteEnumContractsBinder() != nil,
	}
	expected := args.Map{
		"basicBinder":     true,
		"basicByteBinder": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant AsContractsBinder returns non-nil -- Owner", actual)
}

// ── Variant UnmarshallEnumToValue ──

func Test_Cov_Variant_UnmarshallEnumToValue(t *testing.T) {
	v := chmodclasstype.Owner
	val, err := v.UnmarshallEnumToValue([]byte(`"Owner"`))
	actual := args.Map{
		"noErr": err == nil,
		"val":   int(val),
	}
	expected := args.Map{
		"noErr": true,
		"val":   int(chmodclasstype.Owner),
	}
	expected.ShouldBeEqual(t, 0, "Variant UnmarshallEnumToValue returns correct -- Owner", actual)
}
