package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
)

// ==========================================
// Additional method coverage (unique tests only)
// ==========================================

func Test_Request_IsNone_Ext(t *testing.T) {
	if !reqtype.Invalid.IsNone() {
		t.Error("Invalid should be None")
	}
	if reqtype.Create.IsNone() {
		t.Error("Create should not be None")
	}
}

func Test_Request_IsStopEnableStart_Ext(t *testing.T) {
	if reqtype.Create.IsStopEnableStart() {
		t.Error("should return false")
	}
}

func Test_Request_IsStopDisable_Ext(t *testing.T) {
	if reqtype.Create.IsStopDisable() {
		t.Error("should return false")
	}
}

func Test_Request_IsUndefined_Ext(t *testing.T) {
	if !reqtype.Invalid.IsUndefined() {
		t.Error("Invalid should be undefined")
	}
}

func Test_Request_ValueUInt16_Ext(t *testing.T) {
	r := reqtype.Create.ValueUInt16()
	if r != 1 {
		t.Errorf("expected 1, got %d", r)
	}
}

func Test_Request_IntegerEnumRanges_Ext(t *testing.T) {
	r := reqtype.Create.IntegerEnumRanges()
	if len(r) == 0 {
		t.Error("should return non-empty")
	}
}

func Test_Request_MinMaxAny_Ext(t *testing.T) {
	min, max := reqtype.Create.MinMaxAny()
	if min == nil || max == nil {
		t.Error("should return non-nil")
	}
}

func Test_Request_MinMaxValueString_Ext(t *testing.T) {
	if reqtype.Create.MinValueString() == "" {
		t.Error("should return non-empty")
	}
	if reqtype.Create.MaxValueString() == "" {
		t.Error("should return non-empty")
	}
}

func Test_Request_MinMaxInt_Ext(t *testing.T) {
	_ = reqtype.Create.MinInt()
	_ = reqtype.Create.MaxInt()
}

func Test_Request_RangesDynamicMap_Ext(t *testing.T) {
	m := reqtype.Create.RangesDynamicMap()
	if len(m) == 0 {
		t.Error("should return non-empty map")
	}
}

func Test_Request_IsNotOverrideOrOverwriteOrEnforce_Ext(t *testing.T) {
	if !reqtype.Create.IsNotOverrideOrOverwriteOrEnforce() {
		t.Error("Create should not match override group")
	}
	if reqtype.Override.IsNotOverrideOrOverwriteOrEnforce() {
		t.Error("Override should match override group")
	}
}

func Test_Request_IsOverwrite_Ext(t *testing.T) {
	if !reqtype.Overwrite.IsOverwrite() {
		t.Error("Overwrite should match")
	}
}

func Test_Request_IsOverride_Ext(t *testing.T) {
	if !reqtype.Override.IsOverride() {
		t.Error("Override should match")
	}
}

func Test_Request_IsEnforce_Ext(t *testing.T) {
	if !reqtype.Enforce.IsEnforce() {
		t.Error("Enforce should match")
	}
}

func Test_Request_IsValueEqual_Ext(t *testing.T) {
	if !reqtype.Create.IsValueEqual(byte(reqtype.Create)) {
		t.Error("should be equal")
	}
}

func Test_Request_IsOnExistOrSkipOnNonExistLogically_Ext(t *testing.T) {
	if !reqtype.ExistCheck.IsOnExistOrSkipOnNonExistLogically() {
		t.Error("ExistCheck should match")
	}
}

func Test_Request_IsReadOrUpdateLogically_Ext(t *testing.T) {
	if !reqtype.Read.IsReadOrUpdateLogically() {
		t.Error("Read should match")
	}
}

func Test_Request_IsRestartOrReload_Ext(t *testing.T) {
	if !reqtype.Restart.IsRestartOrReload() {
		t.Error("Restart should match")
	}
	if !reqtype.Reload.IsRestartOrReload() {
		t.Error("Reload should match")
	}
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
	if err != nil {
		t.Errorf("all names supported should not error: %v", err)
	}
}

// ==========================================
// RangesOnlySupportedFor
// ==========================================

func Test_RangesOnlySupportedFor_Empty(t *testing.T) {
	err := reqtype.RangesOnlySupportedFor("msg")
	if err != nil {
		t.Error("empty should return nil")
	}
}

func Test_RangesOnlySupportedFor_NonEmpty(t *testing.T) {
	err := reqtype.RangesOnlySupportedFor("msg", reqtype.Create, reqtype.Read)
	if err == nil {
		t.Error("should return error")
	}
}

// ==========================================
// RangesString / RangesStrings
// ==========================================

func Test_RangesString_Ext(t *testing.T) {
	r := reqtype.RangesString(", ", reqtype.Create, reqtype.Read)
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RangesStrings_Ext(t *testing.T) {
	r := reqtype.RangesStrings(reqtype.Create, reqtype.Read)
	if len(r) != 2 {
		t.Errorf("expected 2, got %d", len(r))
	}
}

func Test_RangesStringDefaultJoiner_Ext(t *testing.T) {
	r := reqtype.RangesStringDefaultJoiner(reqtype.Create, reqtype.Read)
	if r == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// Min / Max
// ==========================================

func Test_Request_Min_Ext(t *testing.T) {
	m := reqtype.Min()
	if m != reqtype.Invalid {
		t.Errorf("expected Invalid")
	}
}

func Test_Request_Max_Ext(t *testing.T) {
	m := reqtype.Max()
	if m == reqtype.Invalid {
		t.Error("max should not be Invalid")
	}
}

// ==========================================
// RangesNotMeet / RangesNotMeetError / RangesNotSupportedFor
// ==========================================

func Test_RangesNotMeet_Ext(t *testing.T) {
	r := reqtype.RangesNotMeet("msg", reqtype.Create, reqtype.Read)
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RangesNotMeetError_Ext(t *testing.T) {
	err := reqtype.RangesNotMeetError("msg", reqtype.Create, reqtype.Read)
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RangesNotSupportedFor_Empty_Ext(t *testing.T) {
	err := reqtype.RangesNotSupportedFor("msg")
	if err != nil {
		t.Error("empty should return nil")
	}
}

func Test_RangesNotSupportedFor_NonEmpty_Ext(t *testing.T) {
	err := reqtype.RangesNotSupportedFor("msg", reqtype.Create)
	if err == nil {
		t.Error("should return error")
	}
}

// ==========================================
// RangesInvalidErr
// ==========================================

func Test_RangesInvalidErr_Ext(t *testing.T) {
	err := reqtype.RangesInvalidErr()
	if err == nil {
		t.Error("should return error")
	}
}

// ==========================================
// RangesInBetween
// ==========================================

func Test_RangesInBetween_Ext(t *testing.T) {
	r := reqtype.RangesInBetween(reqtype.Create, reqtype.Read)
	if len(r) == 0 {
		t.Error("should return non-empty")
	}
}

// ==========================================
// ResultStatus
// ==========================================

func Test_ResultStatus_Ext(t *testing.T) {
	rs := reqtype.ResultStatus{}
	if rs.Error != nil {
		t.Error("default should not have error")
	}
	if rs.IsSuccess {
		t.Error("default should not be success")
	}
}

// ==========================================
// OnlySupportedMsgErr
// ==========================================

func Test_Request_OnlySupportedMsgErr_Ext(t *testing.T) {
	err := reqtype.Create.OnlySupportedMsgErr("test", "NonExistent")
	if err == nil {
		t.Error("should return error for unsupported name")
	}
}

// ==========================================
// Additional coverage: RangesNotMeet empty
// ==========================================

func Test_RangesNotMeet_Empty_Ext(t *testing.T) {
	r := reqtype.RangesNotMeet("msg")
	if r != "" {
		t.Error("empty should return empty string")
	}
}

func Test_RangesNotMeetError_Empty_Ext(t *testing.T) {
	err := reqtype.RangesNotMeetError("msg")
	if err != nil {
		t.Error("empty should return nil")
	}
}

// ==========================================
// Additional coverage: RangesStrings empty
// ==========================================

func Test_RangesStrings_Empty_Ext(t *testing.T) {
	r := reqtype.RangesStrings()
	if len(r) != 0 {
		t.Errorf("expected 0, got %d", len(r))
	}
}

// ==========================================
// Additional coverage: IsAnyHttpMethod
// ==========================================

func Test_Request_IsAnyHttpMethod_Ext(t *testing.T) {
	name := reqtype.GetHttp.Name()
	if !reqtype.GetHttp.IsAnyHttpMethod(name) {
		t.Error("GetHttp should match its own name")
	}
	if reqtype.Create.IsAnyHttpMethod("Create") {
		t.Error("Create should not be HTTP method")
	}
}

// ==========================================
// Additional coverage: IsEnumEqual / IsAnyEnumsEqual
// ==========================================

func Test_Request_IsEnumEqual_Ext(t *testing.T) {
	if !reqtype.Create.IsEnumEqual(reqtype.Create.AsBasicEnumContractsBinder()) {
		t.Error("should be equal to self")
	}
}

func Test_Request_IsAnyEnumsEqual_Ext(t *testing.T) {
	r := reqtype.Create
	if !r.IsAnyEnumsEqual(reqtype.Read.AsBasicEnumContractsBinder(), reqtype.Create.AsBasicEnumContractsBinder()) {
		t.Error("should match Create")
	}
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
	if reqtype.Create.NameValue() == "" {
		t.Error("NameValue should not be empty")
	}
}

func Test_Request_RangeNamesCsv_Ext(t *testing.T) {
	if reqtype.Create.RangeNamesCsv() == "" {
		t.Error("RangeNamesCsv should not be empty")
	}
}

func Test_Request_IsUninitialized_Ext(t *testing.T) {
	if !reqtype.Invalid.IsUninitialized() {
		t.Error("Invalid should be uninitialized")
	}
}

// ==========================================
// Additional coverage: CurrentNotImpl
// ==========================================

func Test_Request_CurrentNotImpl_Ext(t *testing.T) {
	err := reqtype.Create.CurrentNotImpl(nil, "test")
	if err == nil {
		t.Error("CurrentNotImpl should return error")
	}
	err = reqtype.Create.CurrentNotImpl("ref", "test")
	if err == nil {
		t.Error("CurrentNotImpl with ref should return error")
	}
}

func Test_Request_NotSupportedErr_Ext(t *testing.T) {
	err := reqtype.Create.NotSupportedErr("not supported", "ref")
	if err == nil {
		t.Error("NotSupportedErr should return error")
	}
}

// ==========================================
// Additional: IsEditOrUpdateLogically
// ==========================================

func Test_Request_IsEditOrUpdateLogically_Ext(t *testing.T) {
	if !reqtype.Update.IsEditOrUpdateLogically() {
		t.Error("Update should match")
	}
}

// Additional: IsCreateOrUpdateLogically
func Test_Request_IsCreateOrUpdateLogically_Ext(t *testing.T) {
	if !reqtype.Create.IsCreateOrUpdateLogically() {
		t.Error("Create should match")
	}
}

// Additional: IsNotCrudOnlyLogically
func Test_Request_IsNotCrudOnlyLogically_Ext(t *testing.T) {
	if reqtype.Read.IsNotCrudOnlyLogically() {
		t.Error("Read is CRUD, should not return true")
	}
	if !reqtype.Touch.IsNotCrudOnlyLogically() {
		t.Error("Touch is not CRUD, should return true")
	}
}

// Additional: IsOnExistCheckLogically
func Test_Request_IsOnExistCheckLogically_Ext(t *testing.T) {
	if !reqtype.ExistCheck.IsOnExistCheckLogically() {
		t.Error("ExistCheck should match")
	}
}

// Additional: DynamicAction
func Test_Request_IsDynamicAction_Ext(t *testing.T) {
	if reqtype.DynamicAction.IsValid() == false {
		t.Error("DynamicAction should be valid")
	}
}

// Additional: IsInBetween edges
func Test_Request_IsInBetween_NotInRange_Ext(t *testing.T) {
	if reqtype.Enforce.IsInBetween(reqtype.Create, reqtype.Delete) {
		t.Error("Enforce should not be between Create and Delete")
	}
}

// Additional: IsAnyOfReqs empty
func Test_Request_IsAnyOfReqs_Empty_Ext(t *testing.T) {
	if !reqtype.Create.IsAnyOfReqs() {
		t.Error("empty should return true")
	}
}

// Additional: IsNotAnyOfReqs empty
func Test_Request_IsNotAnyOfReqs_Empty_Ext(t *testing.T) {
	if !reqtype.Create.IsNotAnyOfReqs() {
		t.Error("empty should return true")
	}
}

// Additional: GetStatusAnyOf empty
func Test_Request_GetStatusAnyOf_Empty_Ext(t *testing.T) {
	status := reqtype.Create.GetStatusAnyOf()
	if !status.IsSuccess {
		t.Error("empty should be success")
	}
}

// Additional: UnmarshalJSON
func Test_Request_UnmarshalJSON_Invalid_Ext(t *testing.T) {
	var r reqtype.Request
	err := r.UnmarshalJSON([]byte(`"NonExistent"`))
	if err == nil {
		t.Error("should error on invalid name")
	}
}

// Additional: Format
func Test_Request_Format_Ext(t *testing.T) {
	result := reqtype.Create.Format("{name}")
	if result == "" {
		t.Error("Format should not be empty")
	}
}

// Additional: ToNumberString
func Test_Request_ToNumberString_Ext(t *testing.T) {
	if reqtype.Create.ToNumberString() == "" {
		t.Error("should not be empty")
	}
}
