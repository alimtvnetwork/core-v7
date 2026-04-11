package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════
// All Is* boolean methods
// ═══════════════════════════════════════════════

func Test_AllIsMethods(t *testing.T) {
	// Arrange
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

	// Act
			actual := args.Map{"result": tt.method(tt.req) != tt.expect}

	// Assert
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "() =, want", actual)
		})
	}
}

// ═══════════════════════════════════════════════
// Composite Is* methods
// ═══════════════════════════════════════════════

func Test_IsAnyApplyOnExist(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.UpdateOnExist.IsAnyApplyOnExist()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.DropOnExist.IsAnyApplyOnExist()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Create.IsAnyApplyOnExist()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsCrud(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Read.IsCrud()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Create.IsCrud()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Append.IsCrud()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsCrudSkip(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.CreateOrSkipOnExist.IsCrudSkip()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Read.IsCrudSkip()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsCrudOrSkip(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Read.IsCrudOrSkip()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.CreateOrSkipOnExist.IsCrudOrSkip()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsAnyDrop(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Drop.IsAnyDrop()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Delete.IsAnyDrop()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Read.IsAnyDrop()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsDropSafe(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.DeleteOrSkipOnNonExist.IsDropSafe()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Drop.IsDropSafe()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsAnyCreate(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsAnyCreate()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.CreateOrAppend.IsAnyCreate()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Read.IsAnyCreate()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsAnyHttp(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.GetHttp.IsAnyHttp()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Create.IsAnyHttp()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsAnyAction(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Start.IsAnyAction()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Create.IsAnyAction()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsNotAnyAction(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsNotAnyAction()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsAnyHttpMethod(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.GetHttp.IsAnyHttpMethod("GetHttp")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Create.IsAnyHttpMethod("Create")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsNotHttpMethod(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsNotHttpMethod()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsNotOverrideOrOverwriteOrEnforce(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsNotOverrideOrOverwriteOrEnforce()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Override.IsNotOverrideOrOverwriteOrEnforce()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// Enum/Name/Value methods
// ═══════════════════════════════════════════════

func Test_Name(t *testing.T) {
	// Arrange
	n := reqtype.Create.Name()

	// Act
	actual := args.Map{"result": n == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty name", actual)
}

func Test_ToNumberString(t *testing.T) {
	// Arrange
	s := reqtype.Create.ToNumberString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_UnmarshallEnumToValue(t *testing.T) {
	// Arrange
	_, err := reqtype.Create.UnmarshallEnumToValue([]byte(`"Read"`))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_IsValidRange(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsValidRange()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsInBetween(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Update.IsInBetween(reqtype.Create, reqtype.Delete)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Append.IsInBetween(reqtype.Create, reqtype.Delete)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_CurrentNotImpl(t *testing.T) {
	// Arrange
	err := reqtype.Create.CurrentNotImpl(nil, "test")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := reqtype.Create.CurrentNotImpl("ref", "test")
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_NotSupportedErr(t *testing.T) {
	// Arrange
	err := reqtype.Create.NotSupportedErr("test msg", "ref")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_IsNotAnyOfReqs(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsNotAnyOfReqs(reqtype.Read, reqtype.Update)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Create.IsNotAnyOfReqs(reqtype.Create)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Create.IsNotAnyOfReqs()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsAnyOfReqs(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsAnyOfReqs(reqtype.Create, reqtype.Read)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Create.IsAnyOfReqs(reqtype.Read)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Create.IsAnyOfReqs()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_GetStatusAnyOf(t *testing.T) {
	// Arrange
	s := reqtype.Create.GetStatusAnyOf(reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"result": s.IsSuccess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	s2 := reqtype.Append.GetStatusAnyOf(reqtype.Create, reqtype.Read)
	actual = args.Map{"result": s2.Error == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	s3 := reqtype.Create.GetStatusAnyOf()
	actual = args.Map{"result": s3.IsSuccess}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_GetInBetweenStatus(t *testing.T) {
	// Arrange
	s := reqtype.Update.GetInBetweenStatus(reqtype.Create, reqtype.Delete)

	// Act
	actual := args.Map{"result": s.IsSuccess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	s2 := reqtype.Append.GetInBetweenStatus(reqtype.Create, reqtype.Delete)
	actual = args.Map{"result": s2.IsSuccess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected failure", actual)
}

func Test_MaxByte(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.MaxByte() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_MinByte(t *testing.T) {
	_ = reqtype.Create.MinByte()
}

func Test_ValueByte(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.ValueByte() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_RangesByte(t *testing.T) {
	// Arrange
	r := reqtype.Create.RangesByte()

	// Act
	actual := args.Map{"result": len(r) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_Value(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.Value() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_ValueInt(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.ValueInt() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsAnyOf(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsAnyOf(reqtype.Create.Value())}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_String(t *testing.T) {
	// Arrange
	s := reqtype.Create.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_UnmarshalJSON(t *testing.T) {
	// Arrange
	r := reqtype.Invalid
	err := r.UnmarshalJSON([]byte(`"Read"`))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_ToPtr(t *testing.T) {
	// Arrange
	p := reqtype.Create.ToPtr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_ToSimple(t *testing.T) {
	// Arrange
	p := reqtype.Create.ToPtr()

	// Act
	actual := args.Map{"result": p.ToSimple() != reqtype.Create}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	var nilP *reqtype.Request
	actual = args.Map{"result": nilP.ToSimple() != reqtype.Invalid}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_MarshalJSON(t *testing.T) {
	// Arrange
	data, err := reqtype.Create.MarshalJSON()

	// Act
	actual := args.Map{"result": err != nil || len(data) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_EnumType(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.EnumType() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_AsBasicEnumContractsBinder(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.AsBasicEnumContractsBinder() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_AsJsonMarshaller(t *testing.T) {
	// Arrange
	r := reqtype.Create

	// Act
	actual := args.Map{"result": r.AsJsonMarshaller() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_AsBasicByteEnumContractsBinder(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.AsBasicByteEnumContractsBinder() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_AsCrudTyper(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.AsCrudTyper() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_AsOverwriteOrRideOrEnforcer(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.AsOverwriteOrRideOrEnforcer() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_AsHttpMethodTyper(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.AsHttpMethodTyper() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_AsActionTyper(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.AsActionTyper() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// IsEnumEqual, IsAnyEnumsEqual, IsNameEqual, etc.
// ═══════════════════════════════════════════════

func Test_IsEnumEqual(t *testing.T) {
	// Arrange
	r := reqtype.Create

	// Act
	actual := args.Map{"result": r.IsEnumEqual(&r)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsByteValueEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsByteValueEqual(reqtype.Create.Value())}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsAnyEnumsEqual(t *testing.T) {
	// Arrange
	r := reqtype.Create
	r2 := reqtype.Create

	// Act
	actual := args.Map{"result": r.IsAnyEnumsEqual(&r2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsNameEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Read.IsNameEqual("Read")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsAnyNamesOf(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Read.IsAnyNamesOf("Read", "Update")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Read.IsAnyNamesOf("Update")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsValueEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Read.IsValueEqual(reqtype.Read.Value())}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsAnyValuesEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Read.IsAnyValuesEqual(reqtype.Read.Value())}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": reqtype.Read.IsAnyValuesEqual(99)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_IsUninitialized(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Invalid.IsUninitialized()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// Package-level functions
// ═══════════════════════════════════════════════

func Test_Max(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Max() != reqtype.DynamicAction}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_Min(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Min() != reqtype.Invalid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_RangesInBetween(t *testing.T) {
	// Arrange
	r := reqtype.RangesInBetween(reqtype.Create, reqtype.Delete)

	// Act
	actual := args.Map{"result": len(r)}

	// Assert
	expected := args.Map{"result": 4}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_RangesInvalidErr(t *testing.T) {
	// Arrange
	err := reqtype.RangesInvalidErr()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_RangesNotMeet(t *testing.T) {
	// Arrange
	s := reqtype.RangesNotMeet("test", reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	s2 := reqtype.RangesNotMeet("test")
	actual = args.Map{"result": s2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for no reqs", actual)
}

func Test_RangesNotMeetError(t *testing.T) {
	// Arrange
	err := reqtype.RangesNotMeetError("test", reqtype.Create)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	err2 := reqtype.RangesNotMeetError("test")
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for no reqs", actual)
}

func Test_RangesNotSupportedFor(t *testing.T) {
	// Arrange
	err := reqtype.RangesNotSupportedFor("test", reqtype.Create)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	err2 := reqtype.RangesNotSupportedFor("test")
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_RangesOnlySupportedFor(t *testing.T) {
	// Arrange
	err := reqtype.RangesOnlySupportedFor("test", reqtype.Create)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	err2 := reqtype.RangesOnlySupportedFor("test")
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_RangesString(t *testing.T) {
	// Arrange
	s := reqtype.RangesString(", ", reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_RangesStringDefaultJoiner(t *testing.T) {
	// Arrange
	s := reqtype.RangesStringDefaultJoiner(reqtype.Create)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_RangesStrings(t *testing.T) {
	// Arrange
	s := reqtype.RangesStrings(reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"result": len(s) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	s2 := reqtype.RangesStrings()
	actual = args.Map{"result": len(s2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}
