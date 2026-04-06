package reqtypetests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Constants ──

func Test_Request_Constants(t *testing.T) {
	actual := args.Map{"result": reqtype.Invalid != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Invalid should be 0", actual)
	actual := args.Map{"result": reqtype.Create != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Create should be 1", actual)
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
		actual := args.Map{"result": req.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, ": String should not be empty", actual)
	}

	// Verify individual Is methods
	actual := args.Map{"result": reqtype.Create.IsCreate()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsCreate", actual)
	actual := args.Map{"result": reqtype.Read.IsRead()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsRead", actual)
	actual := args.Map{"result": reqtype.Update.IsUpdate()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsUpdate", actual)
	actual := args.Map{"result": reqtype.Delete.IsDelete()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsDelete", actual)
	actual := args.Map{"result": reqtype.Drop.IsDrop()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsDrop", actual)
	actual := args.Map{"result": reqtype.CreateOrUpdate.IsCreateOrUpdate()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsCreateOrUpdate", actual)
	actual := args.Map{"result": reqtype.ExistCheck.IsExistCheck()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsExistCheck", actual)
	actual := args.Map{"result": reqtype.SkipOnExist.IsSkipOnExist()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsSkipOnExist", actual)
	actual := args.Map{"result": reqtype.Overwrite.IsOverwrite()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsOverwrite", actual)
	actual := args.Map{"result": reqtype.Override.IsOverride()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsOverride", actual)
	actual := args.Map{"result": reqtype.Enforce.IsEnforce()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEnforce", actual)
	actual := args.Map{"result": reqtype.GetHttp.IsGetHttp()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsGetHttp", actual)
	actual := args.Map{"result": reqtype.PostHttp.IsPostHttp()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsPostHttp", actual)
	actual := args.Map{"result": reqtype.PutHttp.IsPutHttp()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsPutHttp", actual)
	actual := args.Map{"result": reqtype.DeleteHttp.IsDeleteHttp()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsDeleteHttp", actual)
	actual := args.Map{"result": reqtype.PatchHttp.IsPatchHttp()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsPatchHttp", actual)
	actual := args.Map{"result": reqtype.Touch.IsTouch()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsTouch", actual)
	actual := args.Map{"result": reqtype.Start.IsStart()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStart", actual)
	actual := args.Map{"result": reqtype.Stop.IsStop()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStop", actual)
	actual := args.Map{"result": reqtype.Restart.IsRestart()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsRestart", actual)
	actual := args.Map{"result": reqtype.Reload.IsReload()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsReload", actual)
}

// ── Logical groupings ──

func Test_Request_LogicalGroups(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsCreateLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should be create logically", actual)
	actual := args.Map{"result": reqtype.Create.IsCreateOrUpdateLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should be create/update logically", actual)
	actual := args.Map{"result": reqtype.Drop.IsDropLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Drop should be drop logically", actual)
	actual := args.Map{"result": reqtype.Read.IsCrudOnlyLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Read should be CRUD only", actual)
	actual := args.Map{"result": reqtype.Read.IsNotCrudOnlyLogically()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Read should not be NOT CRUD", actual)
	actual := args.Map{"result": reqtype.Read.IsReadOrEditLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Read should be read/edit", actual)
	actual := args.Map{"result": reqtype.Update.IsEditOrUpdateLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Update should be edit/update", actual)
	actual := args.Map{"result": reqtype.ExistCheck.IsOnExistCheckLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ExistCheck should be on exist check", actual)
	actual := args.Map{"result": reqtype.Delete.IsUpdateOrRemoveLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Delete should be update/remove", actual)
}

func Test_Request_OverrideGroup(t *testing.T) {
	actual := args.Map{"result": reqtype.Override.IsOverrideOrOverwriteOrEnforce()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Override should match group", actual)
	actual := args.Map{"result": reqtype.Overwrite.IsOverrideOrOverwriteOrEnforce()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Overwrite should match group", actual)
	actual := args.Map{"result": reqtype.Create.IsOverrideOrOverwriteOrEnforce()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Create should not match override group", actual)
	actual := args.Map{"result": reqtype.Create.IsNotOverrideOrOverwriteOrEnforce()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should not match override", actual)
}

func Test_Request_RestartReload(t *testing.T) {
	actual := args.Map{"result": reqtype.Restart.IsRestartOrReload()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Restart should be restart/reload", actual)
	actual := args.Map{"result": reqtype.Reload.IsRestartOrReload()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Reload should be restart/reload", actual)
}

func Test_Request_AnySkipOnExist(t *testing.T) {
	actual := args.Map{"result": reqtype.SkipOnExist.IsAnySkipOnExist()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "SkipOnExist should be any skip on exist", actual)
}

func Test_Request_AnyApplyOnExist(t *testing.T) {
	actual := args.Map{"result": reqtype.UpdateOnExist.IsAnyApplyOnExist()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "UpdateOnExist should be any apply on exist", actual)
}

func Test_Request_IsCrud(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsCrud()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should be CRUD", actual)
}

func Test_Request_IsCrudSkip(t *testing.T) {
	actual := args.Map{"result": reqtype.CreateOrSkipOnExist.IsCrudSkip()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "CreateOrSkipOnExist should be CRUD skip", actual)
}

func Test_Request_IsCrudOrSkip(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsCrudOrSkip()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should be CRUD or skip", actual)
}

func Test_Request_IsAnyDrop(t *testing.T) {
	actual := args.Map{"result": reqtype.Drop.IsAnyDrop()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Drop should be any drop", actual)
}

func Test_Request_IsDropSafe(t *testing.T) {
	actual := args.Map{"result": reqtype.DropOnExist.IsDropSafe()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DropOnExist should be drop safe", actual)
}

func Test_Request_IsAnyCreate(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsAnyCreate()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should be any create", actual)
}

func Test_Request_IsAnyHttp(t *testing.T) {
	actual := args.Map{"result": reqtype.GetHttp.IsAnyHttp()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetHttp should be any HTTP", actual)
	actual := args.Map{"result": reqtype.Create.IsAnyHttp()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Create should not be any HTTP", actual)
	actual := args.Map{"result": reqtype.Create.IsNotHttpMethod()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should not be HTTP method", actual)
}

func Test_Request_IsAnyAction(t *testing.T) {
	actual := args.Map{"result": reqtype.Start.IsAnyAction()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Start should be any action", actual)
	actual := args.Map{"result": reqtype.Create.IsAnyAction()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Create should not be any action", actual)
	actual := args.Map{"result": reqtype.Create.IsNotAnyAction()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should be not any action", actual)
}

// ── Value conversions ──

func Test_Request_ValueConversions(t *testing.T) {
	r := reqtype.Create

	actual := args.Map{"result": r.Value() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Value mismatch", actual)
	actual := args.Map{"result": r.ValueByte() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueByte mismatch", actual)
	actual := args.Map{"result": r.ValueInt() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueInt mismatch", actual)
	actual := args.Map{"result": r.ValueInt8() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueInt8 mismatch", actual)
	actual := args.Map{"result": r.ValueInt16() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueInt16 mismatch", actual)
	actual := args.Map{"result": r.ValueInt32() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueInt32 mismatch", actual)
	actual := args.Map{"result": r.ValueUInt16() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueUInt16 mismatch", actual)
	actual := args.Map{"result": r.ValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueString should not be empty", actual)
}

func Test_Request_ValidInvalid(t *testing.T) {
	actual := args.Map{"result": reqtype.Invalid.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Invalid should not be valid", actual)
	actual := args.Map{"result": reqtype.Invalid.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be invalid", actual)
	actual := args.Map{"result": reqtype.Invalid.IsUndefined()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be undefined", actual)
	actual := args.Map{"result": reqtype.Invalid.IsNone()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be none", actual)
	actual := args.Map{"result": reqtype.Invalid.IsUninitialized()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be uninitialized", actual)
}

// ── Name / String / NameValue ──

func Test_Request_Name(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.Name() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Name should not be empty", actual)
	actual := args.Map{"result": reqtype.Create.String() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should not be empty", actual)
	actual := args.Map{"result": reqtype.Create.NameValue() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NameValue should not be empty", actual)
	actual := args.Map{"result": reqtype.Create.ToNumberString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToNumberString should not be empty", actual)
}

func Test_Request_IsNameEqual(t *testing.T) {
	name := reqtype.Create.Name()
	actual := args.Map{"result": reqtype.Create.IsNameEqual(name)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsNameEqual should be true", actual)
}

func Test_Request_IsAnyNamesOf(t *testing.T) {
	name := reqtype.Create.Name()
	actual := args.Map{"result": reqtype.Create.IsAnyNamesOf("NonExist", name)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf should find match", actual)
}

// ── Enum info ──

func Test_Request_EnumInfo(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.TypeName() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TypeName should not be empty", actual)
	actual := args.Map{"result": reqtype.Create.RangeNamesCsv() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv should not be empty", actual)
	actual := args.Map{"result": reqtype.Create.IsValidRange()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should be valid range", actual)
}

func Test_Request_MinMax(t *testing.T) {
	min, max := reqtype.Create.MinMaxAny()
	actual := args.Map{"result": min == nil || max == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinMaxAny should not return nil", actual)
	actual := args.Map{"result": reqtype.Create.MinValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinValueString should not be empty", actual)
	actual := args.Map{"result": reqtype.Create.MaxValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxValueString should not be empty", actual)
	actual := args.Map{"result": reqtype.Create.MaxByte() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxByte should not be 0", actual)
}

func Test_Request_Ranges(t *testing.T) {
	actual := args.Map{"result": len(reqtype.Create.IntegerEnumRanges()) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IntegerEnumRanges should not be empty", actual)
	actual := args.Map{"result": len(reqtype.Create.RangesDynamicMap()) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangesDynamicMap should not be empty", actual)
	actual := args.Map{"result": len(reqtype.Create.RangesByte()) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangesByte should not be empty", actual)
}

// ── JSON ──

func Test_Request_JSON(t *testing.T) {
	data, err := json.Marshal(reqtype.Create)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON error:", actual)
	actual := args.Map{"result": len(data) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON should not be empty", actual)

	var r reqtype.Request
	err = json.Unmarshal(data, &r)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON error:", actual)
	actual := args.Map{"result": r != reqtype.Create}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should unmarshal to Create", actual)
}

// ── IsInBetween / IsNotAnyOfReqs / IsAnyOfReqs ──

func Test_Request_IsInBetween(t *testing.T) {
	actual := args.Map{"result": reqtype.Read.IsInBetween(reqtype.Create, reqtype.Delete)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Read should be between Create and Delete", actual)
}

func Test_Request_IsNotAnyOfReqs(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsNotAnyOfReqs(reqtype.Read, reqtype.Update)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should not be any of Read,Update", actual)
	actual := args.Map{"result": reqtype.Create.IsNotAnyOfReqs(reqtype.Create)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Create should be found", actual)
}

func Test_Request_IsAnyOfReqs(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsAnyOfReqs(reqtype.Read, reqtype.Create)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Create should be any of Read,Create", actual)
}

func Test_Request_IsAnyOf(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsAnyOf(byte(reqtype.Create))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsAnyOf should match", actual)
}

func Test_Request_IsAnyValuesEqual(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsAnyValuesEqual(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match value 1", actual)
}

func Test_Request_IsByteValueEqual(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsByteValueEqual(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match byte 1", actual)
}

// ── GetStatusAnyOf / GetInBetweenStatus ──

func Test_Request_GetStatusAnyOf(t *testing.T) {
	status := reqtype.Create.GetStatusAnyOf(reqtype.Create, reqtype.Read)
	actual := args.Map{"result": status.IsSuccess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be success", actual)

	status = reqtype.Create.GetStatusAnyOf(reqtype.Read, reqtype.Update)
	actual := args.Map{"result": status.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have error for no match", actual)
}

func Test_Request_GetInBetweenStatus(t *testing.T) {
	status := reqtype.Read.GetInBetweenStatus(reqtype.Create, reqtype.Delete)
	actual := args.Map{"result": status.IsSuccess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Read should be in between", actual)

	status = reqtype.Overwrite.GetInBetweenStatus(reqtype.Create, reqtype.Delete)
	actual := args.Map{"result": status.IsSuccess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Overwrite should not be in between Create-Delete", actual)
}

// ── ToPtr / ToSimple ──

func Test_Request_ToPtr(t *testing.T) {
	ptr := reqtype.Create.ToPtr()
	actual := args.Map{"result": ptr == nil || *ptr != reqtype.Create}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToPtr mismatch", actual)
}

func Test_Request_ToSimple(t *testing.T) {
	r := reqtype.Create
	actual := args.Map{"result": r.ToSimple() != reqtype.Create}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToSimple mismatch", actual)

	var nilPtr *reqtype.Request
	actual := args.Map{"result": nilPtr.ToSimple() != reqtype.Invalid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ToSimple should return Invalid", actual)
}

// ── Interface bindings ──

func Test_Request_InterfaceBindings(t *testing.T) {
	r := reqtype.Create

	actual := args.Map{"result": r.AsBasicEnumContractsBinder() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AsBasicEnumContractsBinder should not be nil", actual)
	actual := args.Map{"result": r.AsBasicByteEnumContractsBinder() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AsBasicByteEnumContractsBinder should not be nil", actual)
	actual := args.Map{"result": r.AsCrudTyper() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AsCrudTyper should not be nil", actual)
	actual := args.Map{"result": r.AsOverwriteOrRideOrEnforcer() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AsOverwriteOrRideOrEnforcer should not be nil", actual)
	actual := args.Map{"result": r.AsHttpMethodTyper() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AsHttpMethodTyper should not be nil", actual)
	actual := args.Map{"result": r.AsActionTyper() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AsActionTyper should not be nil", actual)
	actual := args.Map{"result": r.EnumType() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "EnumType should not be nil", actual)
	actual := args.Map{"result": r.AsJsonMarshaller() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AsJsonMarshaller should not be nil", actual)
}

// ── Format ──

func Test_Request_Format(t *testing.T) {
	result := reqtype.Create.Format("{name}")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Format should not be empty", actual)
}

// ── CurrentNotImpl ──

func Test_Request_CurrentNotImpl(t *testing.T) {
	err := reqtype.Create.CurrentNotImpl(nil, "test")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CurrentNotImpl should return error", actual)

	err = reqtype.Create.CurrentNotImpl("ref", "test")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CurrentNotImpl with ref should return error", actual)
}

func Test_Request_NotSupportedErr(t *testing.T) {
	err := reqtype.Create.NotSupportedErr("not supported", "ref")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotSupportedErr should return error", actual)
}

// ── StopEnableStart / StopDisable ──

func Test_Request_StopEnableStartDisable(t *testing.T) {
	actual := args.Map{"result": reqtype.Create.IsStopEnableStart()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be false", actual)
	actual := args.Map{"result": reqtype.Create.IsStopDisable()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be false", actual)
}

// ── AllNameValues ──

func Test_Request_AllNameValues(t *testing.T) {
	names := reqtype.Create.AllNameValues()
	actual := args.Map{"result": len(names) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNameValues should not be empty", actual)
}
