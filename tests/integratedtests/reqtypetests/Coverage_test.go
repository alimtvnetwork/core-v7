package reqtypetests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
)

// ── Constants ──

func Test_Request_Constants(t *testing.T) {
	if reqtype.Invalid != 0 {
		t.Error("Invalid should be 0")
	}
	if reqtype.Create != 1 {
		t.Error("Create should be 1")
	}
}

// ── Identity checks ──

func Test_Request_IdentityChecks(t *testing.T) {
	checks := map[reqtype.Request]string{
		reqtype.Create:                         "IsCreate",
		reqtype.Read:                           "IsRead",
		reqtype.Update:                         "IsUpdate",
		reqtype.Delete:                         "IsDelete",
		reqtype.Drop:                           "IsDrop",
		reqtype.CreateOrUpdate:                 "IsCreateOrUpdate",
		reqtype.ExistCheck:                     "IsExistCheck",
		reqtype.SkipOnExist:                    "IsSkipOnExist",
		reqtype.CreateOrSkipOnExist:            "IsCreateOrSkipOnExist",
		reqtype.UpdateOrSkipOnNonExist:         "IsUpdateOrSkipOnNonExist",
		reqtype.DeleteOrSkipOnNonExist:         "IsDeleteOrSkipOnNonExist",
		reqtype.DropOrSkipOnNonExist:           "IsDropOrSkipOnNonExist",
		reqtype.UpdateOnExist:                  "IsUpdateOnExist",
		reqtype.DropOnExist:                    "IsDropOnExist",
		reqtype.DropCreate:                     "IsDropCreate",
		reqtype.Append:                         "IsAppend",
		reqtype.AppendByCompare:                "IsAppendByCompare",
		reqtype.AppendByCompareWhereCommentFound: "IsAppendByCompareWhereCommentFound",
		reqtype.AppendLinesByCompare:           "IsAppendLinesByCompare",
		reqtype.AppendLines:                    "IsAppendLines",
		reqtype.CreateOrAppend:                 "IsCreateOrAppend",
		reqtype.Prepend:                        "IsPrepend",
		reqtype.CreateOrPrepend:                "IsCreateOrPrepend",
		reqtype.PrependLines:                   "IsPrependLines",
		reqtype.Rename:                         "IsRename",
		reqtype.Change:                         "IsChange",
		reqtype.Merge:                          "IsMerge",
		reqtype.MergeLines:                     "IsMergeLines",
		reqtype.GetHttp:                        "IsGetHttp",
		reqtype.PutHttp:                        "IsPutHttp",
		reqtype.PostHttp:                       "IsPostHttp",
		reqtype.DeleteHttp:                     "IsDeleteHttp",
		reqtype.PatchHttp:                      "IsPatchHttp",
		reqtype.Touch:                          "IsTouch",
		reqtype.Start:                          "IsStart",
		reqtype.Stop:                           "IsStop",
		reqtype.Restart:                        "IsRestart",
		reqtype.Reload:                         "IsReload",
		reqtype.StopSleepStart:                 "IsStopSleepStart",
		reqtype.Suspend:                        "IsSuspend",
		reqtype.Pause:                          "IsPause",
		reqtype.Resumed:                        "IsResumed",
		reqtype.TryRestart3Times:               "IsTryRestart3Times",
		reqtype.TryRestart5Times:               "IsTryRestart5Times",
		reqtype.TryStart3Times:                 "IsTryStart3Times",
		reqtype.TryStart5Times:                 "IsTryStart5Times",
		reqtype.TryStop3Times:                  "IsTryStop3Times",
		reqtype.TryStop5Times:                  "IsTryStop5Times",
		reqtype.InheritOnly:                    "IsInheritOnly",
		reqtype.InheritPlusOverride:            "IsInheritPlusOverride",
		reqtype.Overwrite:                      "IsOverwrite",
		reqtype.Override:                       "IsOverride",
		reqtype.Enforce:                        "IsEnforce",
	}

	for req, name := range checks {
		if req.String() == "" {
			t.Errorf("%s: String should not be empty", name)
		}
	}

	// Verify individual Is methods
	if !reqtype.Create.IsCreate() { t.Error("IsCreate") }
	if !reqtype.Read.IsRead() { t.Error("IsRead") }
	if !reqtype.Update.IsUpdate() { t.Error("IsUpdate") }
	if !reqtype.Delete.IsDelete() { t.Error("IsDelete") }
	if !reqtype.Drop.IsDrop() { t.Error("IsDrop") }
	if !reqtype.CreateOrUpdate.IsCreateOrUpdate() { t.Error("IsCreateOrUpdate") }
	if !reqtype.ExistCheck.IsExistCheck() { t.Error("IsExistCheck") }
	if !reqtype.SkipOnExist.IsSkipOnExist() { t.Error("IsSkipOnExist") }
	if !reqtype.Overwrite.IsOverwrite() { t.Error("IsOverwrite") }
	if !reqtype.Override.IsOverride() { t.Error("IsOverride") }
	if !reqtype.Enforce.IsEnforce() { t.Error("IsEnforce") }
	if !reqtype.GetHttp.IsGetHttp() { t.Error("IsGetHttp") }
	if !reqtype.PostHttp.IsPostHttp() { t.Error("IsPostHttp") }
	if !reqtype.PutHttp.IsPutHttp() { t.Error("IsPutHttp") }
	if !reqtype.DeleteHttp.IsDeleteHttp() { t.Error("IsDeleteHttp") }
	if !reqtype.PatchHttp.IsPatchHttp() { t.Error("IsPatchHttp") }
	if !reqtype.Touch.IsTouch() { t.Error("IsTouch") }
	if !reqtype.Start.IsStart() { t.Error("IsStart") }
	if !reqtype.Stop.IsStop() { t.Error("IsStop") }
	if !reqtype.Restart.IsRestart() { t.Error("IsRestart") }
	if !reqtype.Reload.IsReload() { t.Error("IsReload") }
}

// ── Logical groupings ──

func Test_Request_LogicalGroups(t *testing.T) {
	if !reqtype.Create.IsCreateLogically() {
		t.Error("Create should be create logically")
	}
	if !reqtype.Create.IsCreateOrUpdateLogically() {
		t.Error("Create should be create/update logically")
	}
	if !reqtype.Drop.IsDropLogically() {
		t.Error("Drop should be drop logically")
	}
	if !reqtype.Read.IsCrudOnlyLogically() {
		t.Error("Read should be CRUD only")
	}
	if reqtype.Read.IsNotCrudOnlyLogically() {
		t.Error("Read should not be NOT CRUD")
	}
	if !reqtype.Read.IsReadOrEditLogically() {
		t.Error("Read should be read/edit")
	}
	if !reqtype.Update.IsEditOrUpdateLogically() {
		t.Error("Update should be edit/update")
	}
	if !reqtype.ExistCheck.IsOnExistCheckLogically() {
		t.Error("ExistCheck should be on exist check")
	}
	if !reqtype.Delete.IsUpdateOrRemoveLogically() {
		t.Error("Delete should be update/remove")
	}
}

func Test_Request_OverrideGroup(t *testing.T) {
	if !reqtype.Override.IsOverrideOrOverwriteOrEnforce() {
		t.Error("Override should match group")
	}
	if !reqtype.Overwrite.IsOverrideOrOverwriteOrEnforce() {
		t.Error("Overwrite should match group")
	}
	if reqtype.Create.IsOverrideOrOverwriteOrEnforce() {
		t.Error("Create should not match override group")
	}
	if !reqtype.Create.IsNotOverrideOrOverwriteOrEnforce() {
		t.Error("Create should not match override")
	}
}

func Test_Request_RestartReload(t *testing.T) {
	if !reqtype.Restart.IsRestartOrReload() {
		t.Error("Restart should be restart/reload")
	}
	if !reqtype.Reload.IsRestartOrReload() {
		t.Error("Reload should be restart/reload")
	}
}

func Test_Request_AnySkipOnExist(t *testing.T) {
	if !reqtype.SkipOnExist.IsAnySkipOnExist() {
		t.Error("SkipOnExist should be any skip on exist")
	}
}

func Test_Request_AnyApplyOnExist(t *testing.T) {
	if !reqtype.UpdateOnExist.IsAnyApplyOnExist() {
		t.Error("UpdateOnExist should be any apply on exist")
	}
}

func Test_Request_IsCrud(t *testing.T) {
	if !reqtype.Create.IsCrud() {
		t.Error("Create should be CRUD")
	}
}

func Test_Request_IsCrudSkip(t *testing.T) {
	if !reqtype.CreateOrSkipOnExist.IsCrudSkip() {
		t.Error("CreateOrSkipOnExist should be CRUD skip")
	}
}

func Test_Request_IsCrudOrSkip(t *testing.T) {
	if !reqtype.Create.IsCrudOrSkip() {
		t.Error("Create should be CRUD or skip")
	}
}

func Test_Request_IsAnyDrop(t *testing.T) {
	if !reqtype.Drop.IsAnyDrop() {
		t.Error("Drop should be any drop")
	}
}

func Test_Request_IsDropSafe(t *testing.T) {
	if !reqtype.DropOnExist.IsDropSafe() {
		t.Error("DropOnExist should be drop safe")
	}
}

func Test_Request_IsAnyCreate(t *testing.T) {
	if !reqtype.Create.IsAnyCreate() {
		t.Error("Create should be any create")
	}
}

func Test_Request_IsAnyHttp(t *testing.T) {
	if !reqtype.GetHttp.IsAnyHttp() {
		t.Error("GetHttp should be any HTTP")
	}
	if reqtype.Create.IsAnyHttp() {
		t.Error("Create should not be any HTTP")
	}
	if !reqtype.Create.IsNotHttpMethod() {
		t.Error("Create should not be HTTP method")
	}
}

func Test_Request_IsAnyAction(t *testing.T) {
	if !reqtype.Start.IsAnyAction() {
		t.Error("Start should be any action")
	}
	if reqtype.Create.IsAnyAction() {
		t.Error("Create should not be any action")
	}
	if !reqtype.Create.IsNotAnyAction() {
		t.Error("Create should be not any action")
	}
}

// ── Value conversions ──

func Test_Request_ValueConversions(t *testing.T) {
	r := reqtype.Create

	if r.Value() != 1 {
		t.Error("Value mismatch")
	}
	if r.ValueByte() != 1 {
		t.Error("ValueByte mismatch")
	}
	if r.ValueInt() != 1 {
		t.Error("ValueInt mismatch")
	}
	if r.ValueInt8() != 1 {
		t.Error("ValueInt8 mismatch")
	}
	if r.ValueInt16() != 1 {
		t.Error("ValueInt16 mismatch")
	}
	if r.ValueInt32() != 1 {
		t.Error("ValueInt32 mismatch")
	}
	if r.ValueUInt16() != 1 {
		t.Error("ValueUInt16 mismatch")
	}
	if r.ValueString() == "" {
		t.Error("ValueString should not be empty")
	}
}

func Test_Request_ValidInvalid(t *testing.T) {
	if reqtype.Invalid.IsValid() {
		t.Error("Invalid should not be valid")
	}
	if !reqtype.Invalid.IsInvalid() {
		t.Error("Invalid should be invalid")
	}
	if !reqtype.Invalid.IsUndefined() {
		t.Error("Invalid should be undefined")
	}
	if !reqtype.Invalid.IsNone() {
		t.Error("Invalid should be none")
	}
	if !reqtype.Invalid.IsUninitialized() {
		t.Error("Invalid should be uninitialized")
	}
}

// ── Name / String / NameValue ──

func Test_Request_Name(t *testing.T) {
	if reqtype.Create.Name() == "" {
		t.Error("Name should not be empty")
	}
	if reqtype.Create.String() == "" {
		t.Error("String should not be empty")
	}
	if reqtype.Create.NameValue() == "" {
		t.Error("NameValue should not be empty")
	}
	if reqtype.Create.ToNumberString() == "" {
		t.Error("ToNumberString should not be empty")
	}
}

func Test_Request_IsNameEqual(t *testing.T) {
	name := reqtype.Create.Name()
	if !reqtype.Create.IsNameEqual(name) {
		t.Error("IsNameEqual should be true")
	}
}

func Test_Request_IsAnyNamesOf(t *testing.T) {
	name := reqtype.Create.Name()
	if !reqtype.Create.IsAnyNamesOf("NonExist", name) {
		t.Error("IsAnyNamesOf should find match")
	}
}

// ── Enum info ──

func Test_Request_EnumInfo(t *testing.T) {
	if reqtype.Create.TypeName() == "" {
		t.Error("TypeName should not be empty")
	}
	if reqtype.Create.RangeNamesCsv() == "" {
		t.Error("RangeNamesCsv should not be empty")
	}
	if !reqtype.Create.IsValidRange() {
		t.Error("Create should be valid range")
	}
}

func Test_Request_MinMax(t *testing.T) {
	min, max := reqtype.Create.MinMaxAny()
	if min == nil || max == nil {
		t.Error("MinMaxAny should not return nil")
	}
	if reqtype.Create.MinValueString() == "" {
		t.Error("MinValueString should not be empty")
	}
	if reqtype.Create.MaxValueString() == "" {
		t.Error("MaxValueString should not be empty")
	}
	if reqtype.Create.MaxByte() == 0 {
		t.Error("MaxByte should not be 0")
	}
}

func Test_Request_Ranges(t *testing.T) {
	if len(reqtype.Create.IntegerEnumRanges()) == 0 {
		t.Error("IntegerEnumRanges should not be empty")
	}
	if len(reqtype.Create.RangesDynamicMap()) == 0 {
		t.Error("RangesDynamicMap should not be empty")
	}
	if len(reqtype.Create.RangesByte()) == 0 {
		t.Error("RangesByte should not be empty")
	}
}

// ── JSON ──

func Test_Request_JSON(t *testing.T) {
	data, err := json.Marshal(reqtype.Create)
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}
	if len(data) == 0 {
		t.Error("MarshalJSON should not be empty")
	}

	var r reqtype.Request
	err = json.Unmarshal(data, &r)
	if err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
	}
	if r != reqtype.Create {
		t.Error("should unmarshal to Create")
	}
}

// ── IsInBetween / IsNotAnyOfReqs / IsAnyOfReqs ──

func Test_Request_IsInBetween(t *testing.T) {
	if !reqtype.Read.IsInBetween(reqtype.Create, reqtype.Delete) {
		t.Error("Read should be between Create and Delete")
	}
}

func Test_Request_IsNotAnyOfReqs(t *testing.T) {
	if !reqtype.Create.IsNotAnyOfReqs(reqtype.Read, reqtype.Update) {
		t.Error("Create should not be any of Read,Update")
	}
	if reqtype.Create.IsNotAnyOfReqs(reqtype.Create) {
		t.Error("Create should be found")
	}
}

func Test_Request_IsAnyOfReqs(t *testing.T) {
	if !reqtype.Create.IsAnyOfReqs(reqtype.Read, reqtype.Create) {
		t.Error("Create should be any of Read,Create")
	}
}

func Test_Request_IsAnyOf(t *testing.T) {
	if !reqtype.Create.IsAnyOf(byte(reqtype.Create)) {
		t.Error("IsAnyOf should match")
	}
}

func Test_Request_IsAnyValuesEqual(t *testing.T) {
	if !reqtype.Create.IsAnyValuesEqual(1) {
		t.Error("should match value 1")
	}
}

func Test_Request_IsByteValueEqual(t *testing.T) {
	if !reqtype.Create.IsByteValueEqual(1) {
		t.Error("should match byte 1")
	}
}

// ── GetStatusAnyOf / GetInBetweenStatus ──

func Test_Request_GetStatusAnyOf(t *testing.T) {
	status := reqtype.Create.GetStatusAnyOf(reqtype.Create, reqtype.Read)
	if !status.IsSuccess {
		t.Error("should be success")
	}

	status = reqtype.Create.GetStatusAnyOf(reqtype.Read, reqtype.Update)
	if status.Error == nil {
		t.Error("should have error for no match")
	}
}

func Test_Request_GetInBetweenStatus(t *testing.T) {
	status := reqtype.Read.GetInBetweenStatus(reqtype.Create, reqtype.Delete)
	if !status.IsSuccess {
		t.Error("Read should be in between")
	}

	status = reqtype.Overwrite.GetInBetweenStatus(reqtype.Create, reqtype.Delete)
	if status.IsSuccess {
		t.Error("Overwrite should not be in between Create-Delete")
	}
}

// ── ToPtr / ToSimple ──

func Test_Request_ToPtr(t *testing.T) {
	ptr := reqtype.Create.ToPtr()
	if ptr == nil || *ptr != reqtype.Create {
		t.Error("ToPtr mismatch")
	}
}

func Test_Request_ToSimple(t *testing.T) {
	r := reqtype.Create
	if r.ToSimple() != reqtype.Create {
		t.Error("ToSimple mismatch")
	}

	var nilPtr *reqtype.Request
	if nilPtr.ToSimple() != reqtype.Invalid {
		t.Error("nil ToSimple should return Invalid")
	}
}

// ── Interface bindings ──

func Test_Request_InterfaceBindings(t *testing.T) {
	r := reqtype.Create

	if r.AsBasicEnumContractsBinder() == nil {
		t.Error("AsBasicEnumContractsBinder should not be nil")
	}
	if r.AsBasicByteEnumContractsBinder() == nil {
		t.Error("AsBasicByteEnumContractsBinder should not be nil")
	}
	if r.AsCrudTyper() == nil {
		t.Error("AsCrudTyper should not be nil")
	}
	if r.AsOverwriteOrRideOrEnforcer() == nil {
		t.Error("AsOverwriteOrRideOrEnforcer should not be nil")
	}
	if r.AsHttpMethodTyper() == nil {
		t.Error("AsHttpMethodTyper should not be nil")
	}
	if r.AsActionTyper() == nil {
		t.Error("AsActionTyper should not be nil")
	}
	if r.EnumType() == nil {
		t.Error("EnumType should not be nil")
	}
	if r.AsJsonMarshaller() == nil {
		t.Error("AsJsonMarshaller should not be nil")
	}
}

// ── Format ──

func Test_Request_Format(t *testing.T) {
	result := reqtype.Create.Format("{name}")
	if result == "" {
		t.Error("Format should not be empty")
	}
}

// ── CurrentNotImpl ──

func Test_Request_CurrentNotImpl(t *testing.T) {
	err := reqtype.Create.CurrentNotImpl(nil, "test")
	if err == nil {
		t.Error("CurrentNotImpl should return error")
	}

	err = reqtype.Create.CurrentNotImpl("ref", "test")
	if err == nil {
		t.Error("CurrentNotImpl with ref should return error")
	}
}

func Test_Request_NotSupportedErr(t *testing.T) {
	err := reqtype.Create.NotSupportedErr("not supported", "ref")
	if err == nil {
		t.Error("NotSupportedErr should return error")
	}
}

// ── StopEnableStart / StopDisable ──

func Test_Request_StopEnableStartDisable(t *testing.T) {
	if reqtype.Create.IsStopEnableStart() {
		t.Error("should be false")
	}
	if reqtype.Create.IsStopDisable() {
		t.Error("should be false")
	}
}

// ── AllNameValues ──

func Test_Request_AllNameValues(t *testing.T) {
	names := reqtype.Create.AllNameValues()
	if len(names) == 0 {
		t.Error("AllNameValues should not be empty")
	}
}
