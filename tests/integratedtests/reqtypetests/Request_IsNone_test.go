package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Additional method coverage (unique tests only)
// ==========================================

func Test_Request_IsNone_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Invalid.IsNone()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be None", actual)
	actual = args.Map{"result": reqtype.Create.IsNone()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Create should not be None", actual)
}

func Test_Request_IsStopEnableStart_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsStopEnableStart()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return false", actual)
}

func Test_Request_IsStopDisable_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsStopDisable()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return false", actual)
}

func Test_Request_IsUndefined_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Invalid.IsUndefined()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be undefined", actual)
}

func Test_Request_ValueUInt16_Ext(t *testing.T) {
	// Arrange
	r := reqtype.Create.ValueUInt16()

	// Act
	actual := args.Map{"result": r != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Request_IntegerEnumRanges_Ext(t *testing.T) {
	// Arrange
	r := reqtype.Create.IntegerEnumRanges()

	// Act
	actual := args.Map{"result": len(r) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Request_MinMaxAny_Ext(t *testing.T) {
	// Arrange
	min, max := reqtype.Create.MinMaxAny()

	// Act
	actual := args.Map{"result": min == nil || max == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil", actual)
}

func Test_Request_MinMaxValueString_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.MinValueString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
	actual = args.Map{"result": reqtype.Create.MaxValueString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Request_MinMaxInt_Ext(t *testing.T) {
	_ = reqtype.Create.MinInt()
	_ = reqtype.Create.MaxInt()
}

func Test_Request_RangesDynamicMap_Ext(t *testing.T) {
	// Arrange
	m := reqtype.Create.RangesDynamicMap()

	// Act
	actual := args.Map{"result": len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty map", actual)
}

func Test_Request_IsNotOverrideOrOverwriteOrEnforce_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsNotOverrideOrOverwriteOrEnforce()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should not match override group", actual)
	actual = args.Map{"result": reqtype.Override.IsNotOverrideOrOverwriteOrEnforce()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Override should match override group", actual)
}

func Test_Request_IsOverwrite_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Overwrite.IsOverwrite()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Overwrite should match", actual)
}

func Test_Request_IsOverride_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Override.IsOverride()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Override should match", actual)
}

func Test_Request_IsEnforce_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Enforce.IsEnforce()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Enforce should match", actual)
}

func Test_Request_IsValueEqual_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsValueEqual(byte(reqtype.Create))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Request_IsOnExistOrSkipOnNonExistLogically_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.ExistCheck.IsOnExistOrSkipOnNonExistLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ExistCheck should match", actual)
}

func Test_Request_IsReadOrUpdateLogically_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Read.IsReadOrUpdateLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Read should match", actual)
}

func Test_Request_IsRestartOrReload_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Restart.IsRestartOrReload()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Restart should match", actual)
	actual = args.Map{"result": reqtype.Reload.IsRestartOrReload()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Reload should match", actual)
}

func Test_Request_OnlySupportedErr_Ext(t *testing.T) {
	// Arrange
	allNames := []string{
		"Invalid",
		"CreateUsingAliasMap",
		"Read",
		"Update",
		"Delete",
		"Drop",
		"CreateOrUpdate",
		"ExistCheck",
		"SkipOnExist",
		"CreateOrSkipOnExist",
		"UpdateOrSkipOnNonExist",
		"DeleteOrSkipOnNonExist",
		"DropOrSkipOnNonExist",
		"UpdateOnExist",
		"DropOnExist",
		"DropCreate",
		"Append",
		"AppendByCompare",
		"AppendByCompareWhereCommentFound",
		"AppendLinesByCompare",
		"AppendLines",
		"CreateOrAppend",
		"Prepend",
		"CreateOrPrepend",
		"PrependLines",
		"Rename",
		"Change",
		"Merge",
		"MergeLines",
		"GetHttp",
		"PutHttp",
		"PostHttp",
		"DeleteHttp",
		"PatchHttp",
		"Touch",
		"Start",
		"Stop",
		"Restart",
		"Reload",
		"StopSleepStart",
		"Suspend",
		"Pause",
		"Resumed",
		"TryRestart3Times",
		"TryRestart5Times",
		"TryStart3Times",
		"TryStart5Times",
		"TryStop3Times",
		"TryStop5Times",
		"InheritOnly",
		"InheritPlusOverride",
		"DynamicAction",
		"Override",
		"Overwrite",
		"Enforce",
	}
	err := reqtype.Create.OnlySupportedErr(allNames...)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all names supported should not error:", actual)
}

// ==========================================
// RangesOnlySupportedFor
// ==========================================

func Test_RangesOnlySupportedFor_Empty_FromRequestIsNone(t *testing.T) {
	// Arrange
	err := reqtype.RangesOnlySupportedFor("msg")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_RangesOnlySupportedFor_NonEmpty(t *testing.T) {
	// Arrange
	err := reqtype.RangesOnlySupportedFor("msg", reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// ==========================================
// RangesString / RangesStrings
// ==========================================

func Test_RangesString_Ext(t *testing.T) {
	// Arrange
	r := reqtype.RangesString(", ", reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_RangesStrings_Ext(t *testing.T) {
	// Arrange
	r := reqtype.RangesStrings(reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"result": len(r) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_RangesStringDefaultJoiner_Ext(t *testing.T) {
	// Arrange
	r := reqtype.RangesStringDefaultJoiner(reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// ==========================================
// Min / Max
// ==========================================

func Test_Request_Min_Ext(t *testing.T) {
	// Arrange
	m := reqtype.Min()

	// Act
	actual := args.Map{"result": m != reqtype.Invalid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Invalid", actual)
}

func Test_Request_Max_Ext(t *testing.T) {
	// Arrange
	m := reqtype.Max()

	// Act
	actual := args.Map{"result": m == reqtype.Invalid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "max should not be Invalid", actual)
}

// ==========================================
// RangesNotMeet / RangesNotMeetError / RangesNotSupportedFor
// ==========================================

func Test_RangesNotMeet_Ext(t *testing.T) {
	// Arrange
	r := reqtype.RangesNotMeet("msg", reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_RangesNotMeetError_Ext(t *testing.T) {
	// Arrange
	err := reqtype.RangesNotMeetError("msg", reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

func Test_RangesNotSupportedFor_Empty_Ext(t *testing.T) {
	// Arrange
	err := reqtype.RangesNotSupportedFor("msg")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_RangesNotSupportedFor_NonEmpty_Ext(t *testing.T) {
	// Arrange
	err := reqtype.RangesNotSupportedFor("msg", reqtype.Create)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// ==========================================
// RangesInvalidErr
// ==========================================

func Test_RangesInvalidErr_Ext(t *testing.T) {
	// Arrange
	err := reqtype.RangesInvalidErr()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// ==========================================
// RangesInBetween
// ==========================================

func Test_RangesInBetween_Ext(t *testing.T) {
	// Arrange
	r := reqtype.RangesInBetween(reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"result": len(r) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// ==========================================
// ResultStatus
// ==========================================

func Test_ResultStatus_Ext(t *testing.T) {
	// Arrange
	rs := reqtype.ResultStatus{}

	// Act
	actual := args.Map{"result": rs.Error != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "default should not have error", actual)
	actual = args.Map{"result": rs.IsSuccess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "default should not be success", actual)
}

// ==========================================
// OnlySupportedMsgErr
// ==========================================

func Test_Request_OnlySupportedMsgErr_Ext(t *testing.T) {
	// Arrange
	err := reqtype.Create.OnlySupportedMsgErr("test", "NonExistent")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error for unsupported name", actual)
}

// ==========================================
// Additional coverage: RangesNotMeet empty
// ==========================================

func Test_RangesNotMeet_Empty_Ext(t *testing.T) {
	// Arrange
	r := reqtype.RangesNotMeet("msg")

	// Act
	actual := args.Map{"result": r != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return empty string", actual)
}

func Test_RangesNotMeetError_Empty_Ext(t *testing.T) {
	// Arrange
	err := reqtype.RangesNotMeetError("msg")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

// ==========================================
// Additional coverage: RangesStrings empty
// ==========================================

func Test_RangesStrings_Empty_Ext(t *testing.T) {
	// Arrange
	r := reqtype.RangesStrings()

	// Act
	actual := args.Map{"result": len(r) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// Additional coverage: IsAnyHttpMethod
// ==========================================

func Test_Request_IsAnyHttpMethod_Ext(t *testing.T) {
	// Arrange
	name := reqtype.GetHttp.Name()

	// Act
	actual := args.Map{"result": reqtype.GetHttp.IsAnyHttpMethod(name)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetHttp should match its own name", actual)
	actual = args.Map{"result": reqtype.Create.IsAnyHttpMethod("Create")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Create should not be HTTP method", actual)
}

// ==========================================
// Additional coverage: IsEnumEqual / IsAnyEnumsEqual
// ==========================================

func Test_Request_IsEnumEqual_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsEnumEqual(reqtype.Create.AsBasicEnumContractsBinder())}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal to self", actual)
}

func Test_Request_IsAnyEnumsEqual_Ext(t *testing.T) {
	// Arrange
	r := reqtype.Create

	// Act
	actual := args.Map{"result": r.IsAnyEnumsEqual(reqtype.Read.AsBasicEnumContractsBinder(), reqtype.Create.AsBasicEnumContractsBinder())}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match Create", actual)
}

// ==========================================
// Additional coverage: MinByte
// ==========================================

func Test_Request_MinByte_Ext(t *testing.T) {
	_ = reqtype.Create.MinByte()
}

// ==========================================
// Additional coverage: NameValue / RangeNamesCsv / TypeName
// ==========================================

func Test_Request_NameValue_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.NameValue() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NameValue should not be empty", actual)
}

func Test_Request_RangeNamesCsv_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.RangeNamesCsv() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv should not be empty", actual)
}

func Test_Request_IsUninitialized_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Invalid.IsUninitialized()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be uninitialized", actual)
}

// ==========================================
// Additional coverage: CurrentNotImpl
// ==========================================

func Test_Request_CurrentNotImpl_Ext(t *testing.T) {
	// Arrange
	err := reqtype.Create.CurrentNotImpl(nil, "test")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CurrentNotImpl should return error", actual)
	err = reqtype.Create.CurrentNotImpl("ref", "test")
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CurrentNotImpl with ref should return error", actual)
}

func Test_Request_NotSupportedErr_Ext(t *testing.T) {
	// Arrange
	err := reqtype.Create.NotSupportedErr("not supported", "ref")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotSupportedErr should return error", actual)
}

// ==========================================
// Additional: IsEditOrUpdateLogically
// ==========================================

func Test_Request_IsEditOrUpdateLogically_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Update.IsEditOrUpdateLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Update should match", actual)
}

// Additional: IsCreateOrUpdateLogically
func Test_Request_IsCreateOrUpdateLogically_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsCreateOrUpdateLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should match", actual)
}

// Additional: IsNotCrudOnlyLogically
func Test_Request_IsNotCrudOnlyLogically_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Read.IsNotCrudOnlyLogically()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Read is CRUD, should not return true", actual)
	actual = args.Map{"result": reqtype.Touch.IsNotCrudOnlyLogically()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Touch is not CRUD, should return true", actual)
}

// Additional: IsOnExistCheckLogically
func Test_Request_IsOnExistCheckLogically_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.ExistCheck.IsOnExistCheckLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ExistCheck should match", actual)
}

// Additional: DynamicAction
func Test_Request_IsDynamicAction_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.DynamicAction.IsValid() == false}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DynamicAction should be valid", actual)
}

// Additional: IsInBetween edges
func Test_Request_IsInBetween_NotInRange_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Enforce.IsInBetween(reqtype.Create, reqtype.Delete)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Enforce should not be between Create and Delete", actual)
}

// Additional: IsAnyOfReqs empty
func Test_Request_IsAnyOfReqs_Empty_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsAnyOfReqs()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)
}

// Additional: IsNotAnyOfReqs empty
func Test_Request_IsNotAnyOfReqs_Empty_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsNotAnyOfReqs()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)
}

// Additional: GetStatusAnyOf empty
func Test_Request_GetStatusAnyOf_Empty_Ext(t *testing.T) {
	// Arrange
	status := reqtype.Create.GetStatusAnyOf()

	// Act
	actual := args.Map{"result": status.IsSuccess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should be success", actual)
}

// Additional: UnmarshalJSON
func Test_Request_UnmarshalJSON_Invalid_Ext(t *testing.T) {
	// Arrange
	var r reqtype.Request
	err := r.UnmarshalJSON([]byte(`"NonExistent"`))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should error on invalid name", actual)
}

// Additional: Format
func Test_Request_Format_Ext(t *testing.T) {
	// Arrange
	result := reqtype.Create.Format("{name}")

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Format should not be empty", actual)
}

// Additional: ToNumberString
func Test_Request_ToNumberString_Ext(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.ToNumberString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}
