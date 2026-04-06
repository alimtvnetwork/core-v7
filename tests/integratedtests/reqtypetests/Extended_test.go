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
	actual := args.Map{"result": reqtype.Invalid.IsNone()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be None", actual)
	actual := args.Map{"result": reqtype.Create.IsNone()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Create should not be None", actual)
}

func Test_Request_IsStopEnableStart_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsStopEnableStart()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return false", actual)
}

func Test_Request_IsStopDisable_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsStopDisable()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return false", actual)
}

func Test_Request_IsUndefined_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Invalid.IsUndefined()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be undefined", actual)
}

func Test_Request_ValueUInt16_Ext(t *testing.T) {
	r := reqtype.Create.ValueUInt16()
	actual := args.Map{"result": r != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Request_IntegerEnumRanges_Ext(t *testing.T) {
	r := reqtype.Create.IntegerEnumRanges()
	actual := args.Map{"result": len(r) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Request_MinMaxAny_Ext(t *testing.T) {
	min, max := reqtype.Create.MinMaxAny()
	actual := args.Map{"result": min == nil || max == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil", actual)
}

func Test_Request_MinMaxValueString_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.MinValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
	actual := args.Map{"result": reqtype.Create.MaxValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Request_MinMaxInt_Ext(t *testing.T) {
	_ = reqtype.Create.MinInt()
	_ = reqtype.Create.MaxInt()
}

func Test_Request_RangesDynamicMap_Ext(t *testing.T) {
	m := reqtype.Create.RangesDynamicMap()
	actual := args.Map{"result": len(m) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty map", actual)
}

func Test_Request_IsNotOverrideOrOverwriteOrEnforce_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsNotOverrideOrOverwriteOrEnforce()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should not match override group", actual)
	actual := args.Map{"result": reqtype.Override.IsNotOverrideOrOverwriteOrEnforce()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Override should match override group", actual)
}

func Test_Request_IsOverwrite_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Overwrite.IsOverwrite()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Overwrite should match", actual)
}

func Test_Request_IsOverride_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Override.IsOverride()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Override should match", actual)
}

func Test_Request_IsEnforce_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Enforce.IsEnforce()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Enforce should match", actual)
}

func Test_Request_IsValueEqual_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsValueEqual(byte(reqtype.Create))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Request_IsOnExistOrSkipOnNonExistLogically_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.ExistCheck.IsOnExistOrSkipOnNonExistLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ExistCheck should match", actual)
}

func Test_Request_IsReadOrUpdateLogically_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Read.IsReadOrUpdateLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Read should match", actual)
}

func Test_Request_IsRestartOrReload_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Restart.IsRestartOrReload()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Restart should match", actual)
	actual := args.Map{"result": reqtype.Reload.IsRestartOrReload()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Reload should match", actual)
}

func Test_Request_OnlySupportedErr_Ext(t *testing.T) {
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
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all names supported should not error:", actual)
}

// ==========================================
// RangesOnlySupportedFor
// ==========================================

func Test_RangesOnlySupportedFor_Empty(t *testing.T) {
	err := reqtype.RangesOnlySupportedFor("msg")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_RangesOnlySupportedFor_NonEmpty(t *testing.T) {
	err := reqtype.RangesOnlySupportedFor("msg", reqtype.Create, reqtype.Read)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// ==========================================
// RangesString / RangesStrings
// ==========================================

func Test_RangesString_Ext(t *testing.T) {
	r := reqtype.RangesString(", ", reqtype.Create, reqtype.Read)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_RangesStrings_Ext(t *testing.T) {
	r := reqtype.RangesStrings(reqtype.Create, reqtype.Read)
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_RangesStringDefaultJoiner_Ext(t *testing.T) {
	r := reqtype.RangesStringDefaultJoiner(reqtype.Create, reqtype.Read)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// ==========================================
// Min / Max
// ==========================================

func Test_Request_Min_Ext(t *testing.T) {
	m := reqtype.Min()
	actual := args.Map{"result": m != reqtype.Invalid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Invalid", actual)
}

func Test_Request_Max_Ext(t *testing.T) {
	m := reqtype.Max()
	actual := args.Map{"result": m == reqtype.Invalid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "max should not be Invalid", actual)
}

// ==========================================
// RangesNotMeet / RangesNotMeetError / RangesNotSupportedFor
// ==========================================

func Test_RangesNotMeet_Ext(t *testing.T) {
	r := reqtype.RangesNotMeet("msg", reqtype.Create, reqtype.Read)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_RangesNotMeetError_Ext(t *testing.T) {
	err := reqtype.RangesNotMeetError("msg", reqtype.Create, reqtype.Read)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

func Test_RangesNotSupportedFor_Empty_Ext(t *testing.T) {
	err := reqtype.RangesNotSupportedFor("msg")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_RangesNotSupportedFor_NonEmpty_Ext(t *testing.T) {
	err := reqtype.RangesNotSupportedFor("msg", reqtype.Create)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// ==========================================
// RangesInvalidErr
// ==========================================

func Test_RangesInvalidErr_Ext(t *testing.T) {
	err := reqtype.RangesInvalidErr()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// ==========================================
// RangesInBetween
// ==========================================

func Test_RangesInBetween_Ext(t *testing.T) {
	r := reqtype.RangesInBetween(reqtype.Create, reqtype.Read)
	actual := args.Map{"result": len(r) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// ==========================================
// ResultStatus
// ==========================================

func Test_ResultStatus_Ext(t *testing.T) {
	rs := reqtype.ResultStatus{}
	actual := args.Map{"result": rs.Error != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "default should not have error", actual)
	actual := args.Map{"result": rs.IsSuccess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "default should not be success", actual)
}

// ==========================================
// OnlySupportedMsgErr
// ==========================================

func Test_Request_OnlySupportedMsgErr_Ext(t *testing.T) {
	err := reqtype.Create.OnlySupportedMsgErr("test", "NonExistent")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error for unsupported name", actual)
}

// ==========================================
// Additional coverage: RangesNotMeet empty
// ==========================================

func Test_RangesNotMeet_Empty_Ext(t *testing.T) {
	r := reqtype.RangesNotMeet("msg")
	actual := args.Map{"result": r != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return empty string", actual)
}

func Test_RangesNotMeetError_Empty_Ext(t *testing.T) {
	err := reqtype.RangesNotMeetError("msg")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

// ==========================================
// Additional coverage: RangesStrings empty
// ==========================================

func Test_RangesStrings_Empty_Ext(t *testing.T) {
	r := reqtype.RangesStrings()
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// Additional coverage: IsAnyHttpMethod
// ==========================================

func Test_Request_IsAnyHttpMethod_Ext(t *testing.T) {
	name := reqtype.GetHttp.Name()
	actual := args.Map{"result": reqtype.GetHttp.IsAnyHttpMethod(name)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetHttp should match its own name", actual)
	actual := args.Map{"result": reqtype.Create.IsAnyHttpMethod("Create")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Create should not be HTTP method", actual)
}

// ==========================================
// Additional coverage: IsEnumEqual / IsAnyEnumsEqual
// ==========================================

func Test_Request_IsEnumEqual_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsEnumEqual(reqtype.Create.AsBasicEnumContractsBinder())}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal to self", actual)
}

func Test_Request_IsAnyEnumsEqual_Ext(t *testing.T) {
	r := reqtype.Create
	actual := args.Map{"result": r.IsAnyEnumsEqual(reqtype.Read.AsBasicEnumContractsBinder(), reqtype.Create.AsBasicEnumContractsBinder())}
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
	actual := args.Map{"result": reqtype.Create.NameValue() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NameValue should not be empty", actual)
}

func Test_Request_RangeNamesCsv_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.RangeNamesCsv() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv should not be empty", actual)
}

func Test_Request_IsUninitialized_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Invalid.IsUninitialized()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be uninitialized", actual)
}

// ==========================================
// Additional coverage: CurrentNotImpl
// ==========================================

func Test_Request_CurrentNotImpl_Ext(t *testing.T) {
	err := reqtype.Create.CurrentNotImpl(nil, "test")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CurrentNotImpl should return error", actual)
	err = reqtype.Create.CurrentNotImpl("ref", "test")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CurrentNotImpl with ref should return error", actual)
}

func Test_Request_NotSupportedErr_Ext(t *testing.T) {
	err := reqtype.Create.NotSupportedErr("not supported", "ref")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotSupportedErr should return error", actual)
}

// ==========================================
// Additional: IsEditOrUpdateLogically
// ==========================================

func Test_Request_IsEditOrUpdateLogically_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Update.IsEditOrUpdateLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Update should match", actual)
}

// Additional: IsCreateOrUpdateLogically
func Test_Request_IsCreateOrUpdateLogically_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsCreateOrUpdateLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should match", actual)
}

// Additional: IsNotCrudOnlyLogically
func Test_Request_IsNotCrudOnlyLogically_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Read.IsNotCrudOnlyLogically()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Read is CRUD, should not return true", actual)
	actual := args.Map{"result": reqtype.Touch.IsNotCrudOnlyLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Touch is not CRUD, should return true", actual)
}

// Additional: IsOnExistCheckLogically
func Test_Request_IsOnExistCheckLogically_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.ExistCheck.IsOnExistCheckLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ExistCheck should match", actual)
}

// Additional: DynamicAction
func Test_Request_IsDynamicAction_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.DynamicAction.IsValid() == false}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DynamicAction should be valid", actual)
}

// Additional: IsInBetween edges
func Test_Request_IsInBetween_NotInRange_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Enforce.IsInBetween(reqtype.Create, reqtype.Delete)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Enforce should not be between Create and Delete", actual)
}

// Additional: IsAnyOfReqs empty
func Test_Request_IsAnyOfReqs_Empty_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsAnyOfReqs()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)
}

// Additional: IsNotAnyOfReqs empty
func Test_Request_IsNotAnyOfReqs_Empty_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsNotAnyOfReqs()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)
}

// Additional: GetStatusAnyOf empty
func Test_Request_GetStatusAnyOf_Empty_Ext(t *testing.T) {
	status := reqtype.Create.GetStatusAnyOf()
	actual := args.Map{"result": status.IsSuccess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should be success", actual)
}

// Additional: UnmarshalJSON
func Test_Request_UnmarshalJSON_Invalid_Ext(t *testing.T) {
	var r reqtype.Request
	err := r.UnmarshalJSON([]byte(`"NonExistent"`))
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should error on invalid name", actual)
}

// Additional: Format
func Test_Request_Format_Ext(t *testing.T) {
	result := reqtype.Create.Format("{name}")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Format should not be empty", actual)
}

// Additional: ToNumberString
func Test_Request_ToNumberString_Ext(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.ToNumberString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}
