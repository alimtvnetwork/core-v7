package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reqtype"
)

// ── Request — remaining uncovered methods ──

func Test_Request_ValueTypes(t *testing.T) {
	// Act
	actual := args.Map{
		"valueUInt16":  reqtype.Create.ValueUInt16() > 0,
		"valueInt8":    reqtype.Create.ValueInt8() > 0,
		"valueInt16":   reqtype.Create.ValueInt16() > 0,
		"valueInt32":   reqtype.Create.ValueInt32() > 0,
		"valueString":  reqtype.Create.ValueString() != "",
		"valueInt":     reqtype.Create.ValueInt() > 0,
		"valueByte":    reqtype.Create.ValueByte() > 0,
		"value":        reqtype.Create.Value() > 0,
	}

	// Assert
	expected := args.Map{
		"valueUInt16":  true,
		"valueInt8":    true,
		"valueInt16":   true,
		"valueInt32":   true,
		"valueString":  true,
		"valueInt":     true,
		"valueByte":    true,
		"value":        true,
	}
	expected.ShouldBeEqual(t, 0, "Value returns correct value -- type accessors", actual)
}

func Test_Request_EnumMethods(t *testing.T) {
	// Act
	actual := args.Map{
		"integerRanges": len(reqtype.Create.IntegerEnumRanges()) > 0,
		"rangeNamesCsv": reqtype.Create.RangeNamesCsv() != "",
		"typeName":      reqtype.Create.TypeName() != "",
		"nameValue":     reqtype.Create.NameValue() != "",
		"isValid":       reqtype.Create.IsValid(),
		"isInvalid":     reqtype.Invalid.IsInvalid(),
		"isUndefined":   reqtype.Invalid.IsUndefined(),
		"isNone":        reqtype.Invalid.IsNone(),
		"isUninitialized": reqtype.Invalid.IsUninitialized(),
		"minValueString": reqtype.Create.MinValueString() != "",
		"maxValueString": reqtype.Create.MaxValueString() != "",
		"rangesByte":    len(reqtype.Create.RangesByte()) > 0,
		"maxByte":       reqtype.Create.MaxByte() > 0,
		"isValidRange":  reqtype.Create.IsValidRange(),
		"rangeDynMap":   len(reqtype.Create.RangesDynamicMap()) > 0,
	}

	// Assert
	expected := args.Map{
		"integerRanges": true,
		"rangeNamesCsv": true,
		"typeName":      true,
		"nameValue":     true,
		"isValid":       true,
		"isInvalid":     true,
		"isUndefined":   true,
		"isNone":        true,
		"isUninitialized": true,
		"minValueString": true,
		"maxValueString": true,
		"rangesByte":    true,
		"maxByte":       true,
		"isValidRange":  true,
		"rangeDynMap":   true,
	}
	expected.ShouldBeEqual(t, 0, "Enum returns correct value -- methods", actual)
}

func Test_Request_MinMaxAny(t *testing.T) {
	// Act
	min, max := reqtype.Create.MinMaxAny()

	// Assert
	actual := args.Map{
		"minNotNil": min != nil,
		"maxNotNil": max != nil,
	}
	expected := args.Map{
		"minNotNil": true,
		"maxNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxAny returns correct value -- with args", actual)
}

func Test_Request_IsInBetween_FromRequestValueTypes(t *testing.T) {
	// Act
	actual := args.Map{
		"inBetween":  reqtype.Read.IsInBetween(reqtype.Create, reqtype.Delete),
		"outOfRange": reqtype.Touch.IsInBetween(reqtype.Create, reqtype.Delete),
	}

	// Assert
	expected := args.Map{
		"inBetween": true,
		"outOfRange": false,
	}
	expected.ShouldBeEqual(t, 0, "IsInBetween returns correct value -- with args", actual)
}

func Test_Request_GetInBetweenStatus_FromRequestValueTypes(t *testing.T) {
	// Act
	inStatus := reqtype.Read.GetInBetweenStatus(reqtype.Create, reqtype.Delete)
	outStatus := reqtype.Touch.GetInBetweenStatus(reqtype.Create, reqtype.Delete)

	// Assert
	actual := args.Map{
		"inSuccess":  inStatus.IsSuccess,
		"outSuccess": outStatus.IsSuccess,
		"outHasErr":  outStatus.Error != nil,
	}
	expected := args.Map{
		"inSuccess":  true,
		"outSuccess": false,
		"outHasErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "GetInBetweenStatus returns correct value -- with args", actual)
}

func Test_Request_GetStatusAnyOf_Match(t *testing.T) {
	// Act
	status := reqtype.Create.GetStatusAnyOf(reqtype.Read, reqtype.Create)

	// Assert
	actual := args.Map{
		"isSuccess": status.IsSuccess,
		"indexMatch": status.IndexMatch,
	}
	expected := args.Map{
		"isSuccess": true,
		"indexMatch": 1,
	}
	expected.ShouldBeEqual(t, 0, "GetStatusAnyOf returns correct value -- match", actual)
}

func Test_Request_GetStatusAnyOf_NoMatch(t *testing.T) {
	// Act
	status := reqtype.Touch.GetStatusAnyOf(reqtype.Read, reqtype.Create)

	// Assert
	actual := args.Map{"hasError": status.Error != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "GetStatusAnyOf returns empty -- no match", actual)
}

func Test_Request_CompositeLogical(t *testing.T) {
	// Act
	actual := args.Map{
		"crud":          reqtype.Create.IsCrud(),
		"crudSkip":      reqtype.CreateOrSkipOnExist.IsCrudSkip(),
		"crudOrSkip":    reqtype.Create.IsCrudOrSkip(),
		"anyApplyExist": reqtype.UpdateOnExist.IsAnyApplyOnExist(),
		"restartReload": reqtype.Restart.IsRestartOrReload(),
		"notAnyAction":  reqtype.Create.IsNotAnyAction(),
		"notHttp":       reqtype.Create.IsNotHttpMethod(),
		"notOverride":   reqtype.Create.IsNotOverrideOrOverwriteOrEnforce(),
		"overwrite":     reqtype.Overwrite.IsOverwrite(),
		"override":      reqtype.Override.IsOverride(),
		"enforce":       reqtype.Enforce.IsEnforce(),
		"overrideGroup": reqtype.Override.IsOverrideOrOverwriteOrEnforce(),
		"dropSafe":      reqtype.DropOnExist.IsDropSafe(),
		"isNameEqual":   reqtype.Create.IsNameEqual(reqtype.Create.Name()),
		"isByteEqual":   reqtype.Create.IsByteValueEqual(reqtype.Create.Value()),
		"isAnyOf":       reqtype.Create.IsAnyOf(reqtype.Create.Value()),
		"isAnyValues":   reqtype.Create.IsAnyValuesEqual(reqtype.Create.Value()),
	}

	// Assert
	expected := args.Map{
		"crud":          true,
		"crudSkip":      true,
		"crudOrSkip":    true,
		"anyApplyExist": true,
		"restartReload": true,
		"notAnyAction":  true,
		"notHttp":       true,
		"notOverride":   true,
		"overwrite":     true,
		"override":      true,
		"enforce":       true,
		"overrideGroup": true,
		"dropSafe":      true,
		"isNameEqual":   true,
		"isByteEqual":   true,
		"isAnyOf":       true,
		"isAnyValues":   true,
	}
	expected.ShouldBeEqual(t, 0, "Composite returns correct value -- logical", actual)
}

func Test_Request_JsonMarshalUnmarshal(t *testing.T) {
	// Arrange
	r := reqtype.Create

	// Act
	bytes, marshalErr := r.MarshalJSON()
	var r2 reqtype.Request
	unmarshalErr := r2.UnmarshalJSON(bytes)

	// Assert
	actual := args.Map{
		"marshalNoErr":   marshalErr == nil,
		"unmarshalNoErr": unmarshalErr == nil,
		"equal":          r2 == reqtype.Create,
	}
	expected := args.Map{
		"marshalNoErr":   true,
		"unmarshalNoErr": true,
		"equal":          true,
	}
	expected.ShouldBeEqual(t, 0, "JSON returns correct value -- marshal/unmarshal", actual)
}

func Test_Request_ToPtrToSimple(t *testing.T) {
	// Arrange
	r := reqtype.Create
	var nilR *reqtype.Request

	// Act
	actual := args.Map{
		"ptrNotNil":   r.ToPtr() != nil,
		"simpleValue": r.ToPtr().ToSimple() == reqtype.Create,
		"nilSimple":   nilR.ToSimple() == reqtype.Invalid,
	}

	// Assert
	expected := args.Map{
		"ptrNotNil":   true,
		"simpleValue": true,
		"nilSimple":   true,
	}
	expected.ShouldBeEqual(t, 0, "ToPtr returns correct value -- ToSimple", actual)
}

func Test_Request_EnumType(t *testing.T) {
	// Act
	actual := args.Map{"notNil": reqtype.Create.EnumType() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "EnumType returns correct value -- with args", actual)
}

func Test_Request_AsBinders(t *testing.T) {
	// Act
	r := reqtype.Create
	actual := args.Map{
		"basicEnum":     r.AsBasicEnumContractsBinder() != nil,
		"basicByteEnum": r.AsBasicByteEnumContractsBinder() != nil,
		"crudTyper":     r.AsCrudTyper() != nil,
		"overwriter":    r.AsOverwriteOrRideOrEnforcer() != nil,
		"httpTyper":     r.AsHttpMethodTyper() != nil,
		"actionTyper":   r.AsActionTyper() != nil,
		"jsonMarshaller": r.ToPtr().AsJsonMarshaller() != nil,
	}

	// Assert
	expected := args.Map{
		"basicEnum":     true,
		"basicByteEnum": true,
		"crudTyper":     true,
		"overwriter":    true,
		"httpTyper":     true,
		"actionTyper":   true,
		"jsonMarshaller": true,
	}
	expected.ShouldBeEqual(t, 0, "AsBinder returns correct value -- interfaces", actual)
}

func Test_Request_IsNotAnyOfReqs_Match(t *testing.T) {
	// Act
	actual := args.Map{
		"notAny": reqtype.Create.IsNotAnyOfReqs(reqtype.Read, reqtype.Update),
		"match":  reqtype.Create.IsNotAnyOfReqs(reqtype.Read, reqtype.Create),
	}

	// Assert
	expected := args.Map{
		"notAny": true,
		"match": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNotAnyOfReqs returns correct value -- match", actual)
}

func Test_Request_IsAnyOfReqs_Match(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   reqtype.Create.IsAnyOfReqs(reqtype.Read, reqtype.Create),
		"noMatch": reqtype.Create.IsAnyOfReqs(reqtype.Read, reqtype.Update),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyOfReqs returns correct value -- match", actual)
}

func Test_Request_IsAnyNamesOf_FromRequestValueTypes(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   reqtype.Create.IsAnyNamesOf("ReadX", reqtype.Create.Name()),
		"noMatch": reqtype.Create.IsAnyNamesOf("foo", "bar"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf returns correct value -- with args", actual)
}

func Test_Request_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	// Act
	r := reqtype.Create
	actual := args.Map{
		"noMatch": r.IsAnyEnumsEqual(reqtype.Read.AsBasicEnumContractsBinder()),
	}

	// Assert
	expected := args.Map{"noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsAnyEnumsEqual returns empty -- no match", actual)
}
