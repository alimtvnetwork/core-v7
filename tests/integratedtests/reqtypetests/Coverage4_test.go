package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
)

// ═══════════════════════════════════════════════
// All Is* boolean methods
// ═══════════════════════════════════════════════

func Test_Cov4_AllIsMethods(t *testing.T) {
	tests := []struct {
		name   string
		req    reqtype.Request
		method func(reqtype.Request) bool
		expect bool
	}{
		{"IsStopEnableStart", reqtype.Create, func(r reqtype.Request) bool { return r.IsStopEnableStart() }, false},
		{"IsStopDisable", reqtype.Create, func(r reqtype.Request) bool { return r.IsStopDisable() }, false},
		{"IsUndefined", reqtype.Invalid, func(r reqtype.Request) bool { return r.IsUndefined() }, true},
		{"IsUndefined_False", reqtype.Create, func(r reqtype.Request) bool { return r.IsUndefined() }, false},
		{"IsNone", reqtype.Invalid, func(r reqtype.Request) bool { return r.IsNone() }, true},
		{"IsCreateLogically", reqtype.Create, func(r reqtype.Request) bool { return r.IsCreateLogically() }, true},
		{"IsCreateOrUpdateLogically", reqtype.Create, func(r reqtype.Request) bool { return r.IsCreateOrUpdateLogically() }, true},
		{"IsDropLogically", reqtype.Drop, func(r reqtype.Request) bool { return r.IsDropLogically() }, true},
		{"IsCrudOnlyLogically", reqtype.Create, func(r reqtype.Request) bool { return r.IsCrudOnlyLogically() }, true},
		{"IsNotCrudOnlyLogically", reqtype.Append, func(r reqtype.Request) bool { return r.IsNotCrudOnlyLogically() }, true},
		{"IsReadOrEditLogically", reqtype.Read, func(r reqtype.Request) bool { return r.IsReadOrEditLogically() }, true},
		{"IsReadOrUpdateLogically", reqtype.Update, func(r reqtype.Request) bool { return r.IsReadOrUpdateLogically() }, true},
		{"IsEditOrUpdateLogically", reqtype.Update, func(r reqtype.Request) bool { return r.IsEditOrUpdateLogically() }, true},
		{"IsOnExistCheckLogically", reqtype.ExistCheck, func(r reqtype.Request) bool { return r.IsOnExistCheckLogically() }, true},
		{"IsOnExistOrSkipOnNonExistLogically", reqtype.SkipOnExist, func(r reqtype.Request) bool { return r.IsOnExistOrSkipOnNonExistLogically() }, true},
		{"IsUpdateOrRemoveLogically", reqtype.Update, func(r reqtype.Request) bool { return r.IsUpdateOrRemoveLogically() }, true},
		{"IsOverwrite", reqtype.Overwrite, func(r reqtype.Request) bool { return r.IsOverwrite() }, true},
		{"IsOverride", reqtype.Override, func(r reqtype.Request) bool { return r.IsOverride() }, true},
		{"IsEnforce", reqtype.Enforce, func(r reqtype.Request) bool { return r.IsEnforce() }, true},
		{"IsValid", reqtype.Create, func(r reqtype.Request) bool { return r.IsValid() }, true},
		{"IsInvalid", reqtype.Invalid, func(r reqtype.Request) bool { return r.IsInvalid() }, true},
		{"IsCreate", reqtype.Create, func(r reqtype.Request) bool { return r.IsCreate() }, true},
		{"IsRead", reqtype.Read, func(r reqtype.Request) bool { return r.IsRead() }, true},
		{"IsUpdate", reqtype.Update, func(r reqtype.Request) bool { return r.IsUpdate() }, true},
		{"IsDelete", reqtype.Delete, func(r reqtype.Request) bool { return r.IsDelete() }, true},
		{"IsDrop", reqtype.Drop, func(r reqtype.Request) bool { return r.IsDrop() }, true},
		{"IsCreateOrUpdate", reqtype.CreateOrUpdate, func(r reqtype.Request) bool { return r.IsCreateOrUpdate() }, true},
		{"IsExistCheck", reqtype.ExistCheck, func(r reqtype.Request) bool { return r.IsExistCheck() }, true},
		{"IsSkipOnExist", reqtype.SkipOnExist, func(r reqtype.Request) bool { return r.IsSkipOnExist() }, true},
		{"IsCreateOrSkipOnExist", reqtype.CreateOrSkipOnExist, func(r reqtype.Request) bool { return r.IsCreateOrSkipOnExist() }, true},
		{"IsUpdateOrSkipOnNonExist", reqtype.UpdateOrSkipOnNonExist, func(r reqtype.Request) bool { return r.IsUpdateOrSkipOnNonExist() }, true},
		{"IsDeleteOrSkipOnNonExist", reqtype.DeleteOrSkipOnNonExist, func(r reqtype.Request) bool { return r.IsDeleteOrSkipOnNonExist() }, true},
		{"IsDropOrSkipOnNonExist", reqtype.DropOrSkipOnNonExist, func(r reqtype.Request) bool { return r.IsDropOrSkipOnNonExist() }, true},
		{"IsUpdateOnExist", reqtype.UpdateOnExist, func(r reqtype.Request) bool { return r.IsUpdateOnExist() }, true},
		{"IsDropOnExist", reqtype.DropOnExist, func(r reqtype.Request) bool { return r.IsDropOnExist() }, true},
		{"IsDropCreate", reqtype.DropCreate, func(r reqtype.Request) bool { return r.IsDropCreate() }, true},
		{"IsAppend", reqtype.Append, func(r reqtype.Request) bool { return r.IsAppend() }, true},
		{"IsAppendByCompare", reqtype.AppendByCompare, func(r reqtype.Request) bool { return r.IsAppendByCompare() }, true},
		{"IsAppendByCompareWhereCommentFound", reqtype.AppendByCompareWhereCommentFound, func(r reqtype.Request) bool { return r.IsAppendByCompareWhereCommentFound() }, true},
		{"IsAppendLinesByCompare", reqtype.AppendLinesByCompare, func(r reqtype.Request) bool { return r.IsAppendLinesByCompare() }, true},
		{"IsAppendLines", reqtype.AppendLines, func(r reqtype.Request) bool { return r.IsAppendLines() }, true},
		{"IsCreateOrAppend", reqtype.CreateOrAppend, func(r reqtype.Request) bool { return r.IsCreateOrAppend() }, true},
		{"IsPrepend", reqtype.Prepend, func(r reqtype.Request) bool { return r.IsPrepend() }, true},
		{"IsCreateOrPrepend", reqtype.CreateOrPrepend, func(r reqtype.Request) bool { return r.IsCreateOrPrepend() }, true},
		{"IsPrependLines", reqtype.PrependLines, func(r reqtype.Request) bool { return r.IsPrependLines() }, true},
		{"IsRename", reqtype.Rename, func(r reqtype.Request) bool { return r.IsRename() }, true},
		{"IsChange", reqtype.Change, func(r reqtype.Request) bool { return r.IsChange() }, true},
		{"IsMerge", reqtype.Merge, func(r reqtype.Request) bool { return r.IsMerge() }, true},
		{"IsMergeLines", reqtype.MergeLines, func(r reqtype.Request) bool { return r.IsMergeLines() }, true},
		{"IsGetHttp", reqtype.GetHttp, func(r reqtype.Request) bool { return r.IsGetHttp() }, true},
		{"IsPutHttp", reqtype.PutHttp, func(r reqtype.Request) bool { return r.IsPutHttp() }, true},
		{"IsPostHttp", reqtype.PostHttp, func(r reqtype.Request) bool { return r.IsPostHttp() }, true},
		{"IsDeleteHttp", reqtype.DeleteHttp, func(r reqtype.Request) bool { return r.IsDeleteHttp() }, true},
		{"IsPatchHttp", reqtype.PatchHttp, func(r reqtype.Request) bool { return r.IsPatchHttp() }, true},
		{"IsTouch", reqtype.Touch, func(r reqtype.Request) bool { return r.IsTouch() }, true},
		{"IsStart", reqtype.Start, func(r reqtype.Request) bool { return r.IsStart() }, true},
		{"IsStop", reqtype.Stop, func(r reqtype.Request) bool { return r.IsStop() }, true},
		{"IsRestart", reqtype.Restart, func(r reqtype.Request) bool { return r.IsRestart() }, true},
		{"IsReload", reqtype.Reload, func(r reqtype.Request) bool { return r.IsReload() }, true},
		{"IsStopSleepStart", reqtype.StopSleepStart, func(r reqtype.Request) bool { return r.IsStopSleepStart() }, true},
		{"IsSuspend", reqtype.Suspend, func(r reqtype.Request) bool { return r.IsSuspend() }, true},
		{"IsPause", reqtype.Pause, func(r reqtype.Request) bool { return r.IsPause() }, true},
		{"IsResumed", reqtype.Resumed, func(r reqtype.Request) bool { return r.IsResumed() }, true},
		{"IsTryRestart3Times", reqtype.TryRestart3Times, func(r reqtype.Request) bool { return r.IsTryRestart3Times() }, true},
		{"IsTryRestart5Times", reqtype.TryRestart5Times, func(r reqtype.Request) bool { return r.IsTryRestart5Times() }, true},
		{"IsTryStart3Times", reqtype.TryStart3Times, func(r reqtype.Request) bool { return r.IsTryStart3Times() }, true},
		{"IsTryStart5Times", reqtype.TryStart5Times, func(r reqtype.Request) bool { return r.IsTryStart5Times() }, true},
		{"IsTryStop3Times", reqtype.TryStop3Times, func(r reqtype.Request) bool { return r.IsTryStop3Times() }, true},
		{"IsTryStop5Times", reqtype.TryStop5Times, func(r reqtype.Request) bool { return r.IsTryStop5Times() }, true},
		{"IsInheritOnly", reqtype.InheritOnly, func(r reqtype.Request) bool { return r.IsInheritOnly() }, true},
		{"IsInheritPlusOverride", reqtype.InheritPlusOverride, func(r reqtype.Request) bool { return r.IsInheritPlusOverride() }, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.method(tt.req) != tt.expect {
				t.Errorf("%s(%v) = %v, want %v", tt.name, tt.req, !tt.expect, tt.expect)
			}
		})
	}
}

// ═══════════════════════════════════════════════
// Composite Is* methods
// ═══════════════════════════════════════════════

func Test_Cov4_IsAnyApplyOnExist(t *testing.T) {
	if !reqtype.UpdateOnExist.IsAnyApplyOnExist() { t.Fatal() }
	if !reqtype.DropOnExist.IsAnyApplyOnExist() { t.Fatal() }
	if reqtype.Create.IsAnyApplyOnExist() { t.Fatal() }
}

func Test_Cov4_IsCrud(t *testing.T) {
	if !reqtype.Read.IsCrud() { t.Fatal() }
	if !reqtype.Create.IsCrud() { t.Fatal() }
	if reqtype.Append.IsCrud() { t.Fatal() }
}

func Test_Cov4_IsCrudSkip(t *testing.T) {
	if !reqtype.CreateOrSkipOnExist.IsCrudSkip() { t.Fatal() }
	if reqtype.Read.IsCrudSkip() { t.Fatal() }
}

func Test_Cov4_IsCrudOrSkip(t *testing.T) {
	if !reqtype.Read.IsCrudOrSkip() { t.Fatal() }
	if !reqtype.CreateOrSkipOnExist.IsCrudOrSkip() { t.Fatal() }
}

func Test_Cov4_IsAnyDrop(t *testing.T) {
	if !reqtype.Drop.IsAnyDrop() { t.Fatal() }
	if !reqtype.Delete.IsAnyDrop() { t.Fatal() }
	if reqtype.Read.IsAnyDrop() { t.Fatal() }
}

func Test_Cov4_IsDropSafe(t *testing.T) {
	if !reqtype.DeleteOrSkipOnNonExist.IsDropSafe() { t.Fatal() }
	if reqtype.Drop.IsDropSafe() { t.Fatal() }
}

func Test_Cov4_IsAnyCreate(t *testing.T) {
	if !reqtype.Create.IsAnyCreate() { t.Fatal() }
	if !reqtype.CreateOrAppend.IsAnyCreate() { t.Fatal() }
	if reqtype.Read.IsAnyCreate() { t.Fatal() }
}

func Test_Cov4_IsAnyHttp(t *testing.T) {
	if !reqtype.GetHttp.IsAnyHttp() { t.Fatal() }
	if reqtype.Create.IsAnyHttp() { t.Fatal() }
}

func Test_Cov4_IsAnyAction(t *testing.T) {
	if !reqtype.Start.IsAnyAction() { t.Fatal() }
	if reqtype.Create.IsAnyAction() { t.Fatal() }
}

func Test_Cov4_IsNotAnyAction(t *testing.T) {
	if !reqtype.Create.IsNotAnyAction() { t.Fatal() }
}

func Test_Cov4_IsAnyHttpMethod(t *testing.T) {
	if !reqtype.GetHttp.IsAnyHttpMethod("GetHttp") { t.Fatal() }
	if reqtype.Create.IsAnyHttpMethod("Create") { t.Fatal() }
}

func Test_Cov4_IsNotHttpMethod(t *testing.T) {
	if !reqtype.Create.IsNotHttpMethod() { t.Fatal() }
}

func Test_Cov4_IsNotOverrideOrOverwriteOrEnforce(t *testing.T) {
	if !reqtype.Create.IsNotOverrideOrOverwriteOrEnforce() { t.Fatal() }
	if reqtype.Override.IsNotOverrideOrOverwriteOrEnforce() { t.Fatal() }
}

// ═══════════════════════════════════════════════
// Enum/Name/Value methods
// ═══════════════════════════════════════════════

func Test_Cov4_Name(t *testing.T) {
	n := reqtype.Create.Name()
	if n == "" { t.Fatal("expected non-empty name") }
}

func Test_Cov4_ToNumberString(t *testing.T) {
	s := reqtype.Create.ToNumberString()
	if s == "" { t.Fatal("expected non-empty") }
}

func Test_Cov4_UnmarshallEnumToValue(t *testing.T) {
	_, err := reqtype.Create.UnmarshallEnumToValue([]byte(`"Read"`))
	if err != nil { t.Fatal(err) }
}

func Test_Cov4_IsValidRange(t *testing.T) {
	if !reqtype.Create.IsValidRange() { t.Fatal() }
}

func Test_Cov4_IsInBetween(t *testing.T) {
	if !reqtype.Update.IsInBetween(reqtype.Create, reqtype.Delete) { t.Fatal() }
	if reqtype.Append.IsInBetween(reqtype.Create, reqtype.Delete) { t.Fatal() }
}

func Test_Cov4_CurrentNotImpl(t *testing.T) {
	err := reqtype.Create.CurrentNotImpl(nil, "test")
	if err == nil { t.Fatal("expected error") }
	err2 := reqtype.Create.CurrentNotImpl("ref", "test")
	if err2 == nil { t.Fatal("expected error") }
}

func Test_Cov4_NotSupportedErr(t *testing.T) {
	err := reqtype.Create.NotSupportedErr("test msg", "ref")
	if err == nil { t.Fatal("expected error") }
}

func Test_Cov4_IsNotAnyOfReqs(t *testing.T) {
	if !reqtype.Create.IsNotAnyOfReqs(reqtype.Read, reqtype.Update) { t.Fatal() }
	if reqtype.Create.IsNotAnyOfReqs(reqtype.Create) { t.Fatal() }
	if !reqtype.Create.IsNotAnyOfReqs() { t.Fatal() }
}

func Test_Cov4_IsAnyOfReqs(t *testing.T) {
	if !reqtype.Create.IsAnyOfReqs(reqtype.Create, reqtype.Read) { t.Fatal() }
	if reqtype.Create.IsAnyOfReqs(reqtype.Read) { t.Fatal() }
	if !reqtype.Create.IsAnyOfReqs() { t.Fatal() }
}

func Test_Cov4_GetStatusAnyOf(t *testing.T) {
	s := reqtype.Create.GetStatusAnyOf(reqtype.Create, reqtype.Read)
	if !s.IsSuccess { t.Fatal() }
	s2 := reqtype.Append.GetStatusAnyOf(reqtype.Create, reqtype.Read)
	if s2.Error == nil { t.Fatal("expected error") }
	s3 := reqtype.Create.GetStatusAnyOf()
	if !s3.IsSuccess { t.Fatal() }
}

func Test_Cov4_GetInBetweenStatus(t *testing.T) {
	s := reqtype.Update.GetInBetweenStatus(reqtype.Create, reqtype.Delete)
	if !s.IsSuccess { t.Fatal() }
	s2 := reqtype.Append.GetInBetweenStatus(reqtype.Create, reqtype.Delete)
	if s2.IsSuccess { t.Fatal("expected failure") }
}

func Test_Cov4_MaxByte(t *testing.T) {
	if reqtype.Create.MaxByte() == 0 { t.Fatal() }
}

func Test_Cov4_MinByte(t *testing.T) {
	_ = reqtype.Create.MinByte()
}

func Test_Cov4_ValueByte(t *testing.T) {
	if reqtype.Create.ValueByte() == 0 { t.Fatal() }
}

func Test_Cov4_RangesByte(t *testing.T) {
	r := reqtype.Create.RangesByte()
	if len(r) == 0 { t.Fatal() }
}

func Test_Cov4_Value(t *testing.T) {
	if reqtype.Create.Value() == 0 { t.Fatal() }
}

func Test_Cov4_ValueInt(t *testing.T) {
	if reqtype.Create.ValueInt() == 0 { t.Fatal() }
}

func Test_Cov4_IsAnyOf(t *testing.T) {
	if !reqtype.Create.IsAnyOf(reqtype.Create.Value()) { t.Fatal() }
}

func Test_Cov4_String(t *testing.T) {
	s := reqtype.Create.String()
	if s == "" { t.Fatal() }
}

func Test_Cov4_UnmarshalJSON(t *testing.T) {
	r := reqtype.Invalid
	err := r.UnmarshalJSON([]byte(`"Read"`))
	if err != nil { t.Fatal(err) }
}

func Test_Cov4_ToPtr(t *testing.T) {
	p := reqtype.Create.ToPtr()
	if p == nil { t.Fatal() }
}

func Test_Cov4_ToSimple(t *testing.T) {
	p := reqtype.Create.ToPtr()
	if p.ToSimple() != reqtype.Create { t.Fatal() }
	var nilP *reqtype.Request
	if nilP.ToSimple() != reqtype.Invalid { t.Fatal() }
}

func Test_Cov4_MarshalJSON(t *testing.T) {
	data, err := reqtype.Create.MarshalJSON()
	if err != nil || len(data) == 0 { t.Fatal() }
}

func Test_Cov4_EnumType(t *testing.T) {
	if reqtype.Create.EnumType() == nil { t.Fatal() }
}

func Test_Cov4_AsBasicEnumContractsBinder(t *testing.T) {
	if reqtype.Create.AsBasicEnumContractsBinder() == nil { t.Fatal() }
}

func Test_Cov4_AsJsonMarshaller(t *testing.T) {
	r := reqtype.Create
	if r.AsJsonMarshaller() == nil { t.Fatal() }
}

func Test_Cov4_AsBasicByteEnumContractsBinder(t *testing.T) {
	if reqtype.Create.AsBasicByteEnumContractsBinder() == nil { t.Fatal() }
}

func Test_Cov4_AsCrudTyper(t *testing.T) {
	if reqtype.Create.AsCrudTyper() == nil { t.Fatal() }
}

func Test_Cov4_AsOverwriteOrRideOrEnforcer(t *testing.T) {
	if reqtype.Create.AsOverwriteOrRideOrEnforcer() == nil { t.Fatal() }
}

func Test_Cov4_AsHttpMethodTyper(t *testing.T) {
	if reqtype.Create.AsHttpMethodTyper() == nil { t.Fatal() }
}

func Test_Cov4_AsActionTyper(t *testing.T) {
	if reqtype.Create.AsActionTyper() == nil { t.Fatal() }
}

// ═══════════════════════════════════════════════
// IsEnumEqual, IsAnyEnumsEqual, IsNameEqual, etc.
// ═══════════════════════════════════════════════

func Test_Cov4_IsEnumEqual(t *testing.T) {
	r := reqtype.Create
	if !r.IsEnumEqual(&r) { t.Fatal() }
}

func Test_Cov4_IsByteValueEqual(t *testing.T) {
	if !reqtype.Create.IsByteValueEqual(reqtype.Create.Value()) { t.Fatal() }
}

func Test_Cov4_IsAnyEnumsEqual(t *testing.T) {
	r := reqtype.Create
	r2 := reqtype.Create
	if !r.IsAnyEnumsEqual(&r2) { t.Fatal() }
}

func Test_Cov4_IsNameEqual(t *testing.T) {
	if !reqtype.Read.IsNameEqual("Read") { t.Fatal() }
}

func Test_Cov4_IsAnyNamesOf(t *testing.T) {
	if !reqtype.Read.IsAnyNamesOf("Read", "Update") { t.Fatal() }
	if reqtype.Read.IsAnyNamesOf("Update") { t.Fatal() }
}

func Test_Cov4_IsValueEqual(t *testing.T) {
	if !reqtype.Read.IsValueEqual(reqtype.Read.Value()) { t.Fatal() }
}

func Test_Cov4_IsAnyValuesEqual(t *testing.T) {
	if !reqtype.Read.IsAnyValuesEqual(reqtype.Read.Value()) { t.Fatal() }
	if reqtype.Read.IsAnyValuesEqual(99) { t.Fatal() }
}

func Test_Cov4_IsUninitialized(t *testing.T) {
	if !reqtype.Invalid.IsUninitialized() { t.Fatal() }
}

// ═══════════════════════════════════════════════
// Package-level functions
// ═══════════════════════════════════════════════

func Test_Cov4_Max(t *testing.T) {
	if reqtype.Max() != reqtype.DynamicAction { t.Fatal() }
}

func Test_Cov4_Min(t *testing.T) {
	if reqtype.Min() != reqtype.Invalid { t.Fatal() }
}

func Test_Cov4_RangesInBetween(t *testing.T) {
	r := reqtype.RangesInBetween(reqtype.Create, reqtype.Delete)
	if len(r) != 4 { t.Fatalf("expected 4 got %d", len(r)) }
}

func Test_Cov4_RangesInvalidErr(t *testing.T) {
	err := reqtype.RangesInvalidErr()
	if err == nil { t.Fatal() }
}

func Test_Cov4_RangesNotMeet(t *testing.T) {
	s := reqtype.RangesNotMeet("test", reqtype.Create, reqtype.Read)
	if s == "" { t.Fatal() }
	s2 := reqtype.RangesNotMeet("test")
	if s2 != "" { t.Fatal("expected empty for no reqs") }
}

func Test_Cov4_RangesNotMeetError(t *testing.T) {
	err := reqtype.RangesNotMeetError("test", reqtype.Create)
	if err == nil { t.Fatal() }
	err2 := reqtype.RangesNotMeetError("test")
	if err2 != nil { t.Fatal("expected nil for no reqs") }
}

func Test_Cov4_RangesNotSupportedFor(t *testing.T) {
	err := reqtype.RangesNotSupportedFor("test", reqtype.Create)
	if err == nil { t.Fatal() }
	err2 := reqtype.RangesNotSupportedFor("test")
	if err2 != nil { t.Fatal() }
}

func Test_Cov4_RangesOnlySupportedFor(t *testing.T) {
	err := reqtype.RangesOnlySupportedFor("test", reqtype.Create)
	if err == nil { t.Fatal() }
	err2 := reqtype.RangesOnlySupportedFor("test")
	if err2 != nil { t.Fatal() }
}

func Test_Cov4_RangesString(t *testing.T) {
	s := reqtype.RangesString(", ", reqtype.Create, reqtype.Read)
	if s == "" { t.Fatal() }
}

func Test_Cov4_RangesStringDefaultJoiner(t *testing.T) {
	s := reqtype.RangesStringDefaultJoiner(reqtype.Create)
	if s == "" { t.Fatal() }
}

func Test_Cov4_RangesStrings(t *testing.T) {
	s := reqtype.RangesStrings(reqtype.Create, reqtype.Read)
	if len(s) != 2 { t.Fatal() }
	s2 := reqtype.RangesStrings()
	if len(s2) != 0 { t.Fatal() }
}
