package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reqtype"
)

// ── Package-level functions ──

func Test_Min_FromMin(t *testing.T) {
	// Act
	actual := args.Map{"value": reqtype.Min()}

	// Assert
	expected := args.Map{"value": reqtype.Invalid}
	expected.ShouldBeEqual(t, 0, "Min returns error -- returns Invalid", actual)
}

func Test_Max_FromMin(t *testing.T) {
	// Act
	actual := args.Map{"notInvalid": reqtype.Max() != reqtype.Invalid}

	// Assert
	expected := args.Map{"notInvalid": true}
	expected.ShouldBeEqual(t, 0, "Max returns error -- returns non-Invalid", actual)
}

func Test_RangesInBetween_FromMin(t *testing.T) {
	// Arrange
	result := reqtype.RangesInBetween(reqtype.Create, reqtype.Delete)

	// Act
	actual := args.Map{
		"len": len(result),
		"firstIsCreate": result[0] == reqtype.Create,
	}

	// Assert
	expected := args.Map{
		"len": 4,
		"firstIsCreate": true,
	}
	expected.ShouldBeEqual(t, 0, "RangesInBetween returns correct value -- Create-Delete", actual)
}

func Test_RangesStrings_FromMin(t *testing.T) {
	// Arrange
	result := reqtype.RangesStrings(reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{
		"len": len(result),
		"firstNotEmpty": result[0] != "",
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"firstNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "RangesStrings returns correct value -- with args", actual)
}

func Test_RangesStrings_Empty(t *testing.T) {
	// Arrange
	result := reqtype.RangesStrings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RangesStrings returns empty -- empty", actual)
}

func Test_RangesString_FromMin(t *testing.T) {
	// Arrange
	result := reqtype.RangesString(", ", reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangesString returns correct value -- with args", actual)
}

func Test_RangesStringDefaultJoiner_FromMin(t *testing.T) {
	// Arrange
	result := reqtype.RangesStringDefaultJoiner(reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangesStringDefaultJoiner returns correct value -- with args", actual)
}

func Test_RangesNotMeet_FromMin(t *testing.T) {
	// Arrange
	result := reqtype.RangesNotMeet("test msg", reqtype.Create, reqtype.Read)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangesNotMeet returns correct value -- with args", actual)
}

func Test_RangesNotMeet_Empty(t *testing.T) {
	// Arrange
	result := reqtype.RangesNotMeet("test msg")

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "RangesNotMeet returns empty -- empty", actual)
}

func Test_RangesNotMeetError_FromMin(t *testing.T) {
	// Arrange
	err := reqtype.RangesNotMeetError("test msg", reqtype.Create)

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "RangesNotMeetError returns error -- with args", actual)
}

func Test_RangesNotMeetError_Empty(t *testing.T) {
	// Arrange
	err := reqtype.RangesNotMeetError("test msg")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RangesNotMeetError returns empty -- empty", actual)
}

func Test_RangesInvalidErr_FromMin(t *testing.T) {
	// Arrange
	err := reqtype.RangesInvalidErr()

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "RangesInvalidErr returns error -- with args", actual)
}

func Test_RangesNotSupportedFor_FromMin(t *testing.T) {
	// Arrange
	err := reqtype.RangesNotSupportedFor("msg", reqtype.Create)

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "RangesNotSupportedFor returns correct value -- with args", actual)
}

func Test_RangesNotSupportedFor_Empty(t *testing.T) {
	// Arrange
	err := reqtype.RangesNotSupportedFor("msg")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RangesNotSupportedFor returns empty -- empty", actual)
}

func Test_RangesOnlySupportedFor_FromMin(t *testing.T) {
	// Arrange
	err := reqtype.RangesOnlySupportedFor("msg", reqtype.Create)

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "RangesOnlySupportedFor returns correct value -- with args", actual)
}

func Test_RangesOnlySupportedFor_Empty(t *testing.T) {
	// Arrange
	err := reqtype.RangesOnlySupportedFor("msg")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RangesOnlySupportedFor returns empty -- empty", actual)
}

// ── Request methods not yet covered ──

func Test_Request_IsStopEnableStart(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsStopEnableStart()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStopEnableStart returns non-empty -- always false", actual)
}

func Test_Request_IsStopDisable(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsStopDisable()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStopDisable returns non-empty -- always false", actual)
}

func Test_Request_AllNameValues(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": len(reqtype.Create.AllNameValues()) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AllNameValues returns non-empty -- with args", actual)
}

func Test_Request_OnlySupportedErr(t *testing.T) {
	// Arrange
	err := reqtype.Create.OnlySupportedErr("Create")

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns error -- with args", actual)
}

func Test_Request_OnlySupportedMsgErr(t *testing.T) {
	// Arrange
	err := reqtype.Create.OnlySupportedMsgErr("test", "Create")

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedMsgErr returns error -- with args", actual)
}

func Test_Request_IsReadOrUpdateLogically(t *testing.T) {
	// Act
	actual := args.Map{"read": reqtype.Read.IsReadOrUpdateLogically()}

	// Assert
	expected := args.Map{"read": true}
	expected.ShouldBeEqual(t, 0, "IsReadOrUpdateLogically returns correct value -- with args", actual)
}

func Test_Request_IsOnExistOrSkipOnNonExistLogically(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.ExistCheck.IsOnExistOrSkipOnNonExistLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsOnExistOrSkipOnNonExistLogically returns correct value -- with args", actual)
}

func Test_Request_CurrentNotImpl(t *testing.T) {
	// Arrange
	err := reqtype.Create.CurrentNotImpl(nil)

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "CurrentNotImpl returns nil -- nil ref", actual)

	err2 := reqtype.Create.CurrentNotImpl("ref", "extra msg")
	actual2 := args.Map{"hasError": err2 != nil}
	expected2 := args.Map{"hasError": true}
	expected2.ShouldBeEqual(t, 1, "CurrentNotImpl returns non-empty -- with ref", actual2)
}

func Test_Request_NotSupportedErr(t *testing.T) {
	// Arrange
	err := reqtype.Create.NotSupportedErr("msg", "ref")

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "NotSupportedErr returns error -- with args", actual)
}

func Test_Request_IsNotAnyOfReqs_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsNotAnyOfReqs()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsNotAnyOfReqs returns empty -- empty", actual)
}

func Test_Request_IsAnyOfReqs_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsAnyOfReqs()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsAnyOfReqs returns empty -- empty", actual)
}

func Test_Request_GetStatusAnyOf_Empty(t *testing.T) {
	// Arrange
	status := reqtype.Create.GetStatusAnyOf()

	// Act
	actual := args.Map{"isSuccess": status.IsSuccess}

	// Assert
	expected := args.Map{"isSuccess": true}
	expected.ShouldBeEqual(t, 0, "GetStatusAnyOf returns empty -- empty", actual)
}

func Test_Request_IsAnyHttpMethod(t *testing.T) {
	// Arrange
	name := reqtype.GetHttp.Name()

	// Act
	actual := args.Map{
		"match":   reqtype.GetHttp.IsAnyHttpMethod(name),
		"noMatch": reqtype.Create.IsAnyHttpMethod(name),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyHttpMethod returns correct value -- with args", actual)
}

func Test_Request_IsEnumEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": reqtype.Create.IsEnumEqual(reqtype.Create.AsBasicEnumContractsBinder())}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEnumEqual returns correct value -- with args", actual)
}

func Test_Request_IsAnyEnumsEqual(t *testing.T) {
	// Arrange
	r := reqtype.Create

	// Act
	actual := args.Map{"result": r.IsAnyEnumsEqual(reqtype.Read.AsBasicEnumContractsBinder(), reqtype.Create.AsBasicEnumContractsBinder())}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsAnyEnumsEqual returns correct value -- with args", actual)
}

func Test_Request_MinInt_MaxInt(t *testing.T) {
	// Act
	actual := args.Map{
		"minValid": reqtype.Create.MinInt() >= 0,
		"maxValid": reqtype.Create.MaxInt() > 0,
		"minByte":  reqtype.Create.MinByte() == 0,
	}

	// Assert
	expected := args.Map{
		"minValid": true,
		"maxValid": true,
		"minByte":  true,
	}
	expected.ShouldBeEqual(t, 0, "MinInt returns correct value -- MaxInt MinByte", actual)
}

func Test_Request_Format(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": reqtype.Create.Format("%s") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Format returns correct value -- with args", actual)
}

func Test_Request_DynamicAction(t *testing.T) {
	// Act
	actual := args.Map{"isDynamic": reqtype.DynamicAction.String() != ""}

	// Assert
	expected := args.Map{"isDynamic": true}
	expected.ShouldBeEqual(t, 0, "DynamicAction returns correct value -- string", actual)
}

// ── Identity checks for remaining values ──

func Test_Request_RemainingIdentity(t *testing.T) {
	// Act
	actual := args.Map{
		"stopSleepStart":    reqtype.StopSleepStart.IsStopSleepStart(),
		"suspend":           reqtype.Suspend.IsSuspend(),
		"pause":             reqtype.Pause.IsPause(),
		"resumed":           reqtype.Resumed.IsResumed(),
		"tryRestart3":       reqtype.TryRestart3Times.IsTryRestart3Times(),
		"tryRestart5":       reqtype.TryRestart5Times.IsTryRestart5Times(),
		"tryStart3":         reqtype.TryStart3Times.IsTryStart3Times(),
		"tryStart5":         reqtype.TryStart5Times.IsTryStart5Times(),
		"tryStop3":          reqtype.TryStop3Times.IsTryStop3Times(),
		"tryStop5":          reqtype.TryStop5Times.IsTryStop5Times(),
		"inheritOnly":       reqtype.InheritOnly.IsInheritOnly(),
		"inheritPlusOvrd":   reqtype.InheritPlusOverride.IsInheritPlusOverride(),
		"createOrSkipExist": reqtype.CreateOrSkipOnExist.IsCreateOrSkipOnExist(),
		"updateSkipNon":     reqtype.UpdateOrSkipOnNonExist.IsUpdateOrSkipOnNonExist(),
		"deleteSkipNon":     reqtype.DeleteOrSkipOnNonExist.IsDeleteOrSkipOnNonExist(),
		"dropSkipNon":       reqtype.DropOrSkipOnNonExist.IsDropOrSkipOnNonExist(),
		"updateOnExist":     reqtype.UpdateOnExist.IsUpdateOnExist(),
		"dropOnExist":       reqtype.DropOnExist.IsDropOnExist(),
		"dropCreate":        reqtype.DropCreate.IsDropCreate(),
		"appendBC":          reqtype.AppendByCompare.IsAppendByCompare(),
		"appendBCW":         reqtype.AppendByCompareWhereCommentFound.IsAppendByCompareWhereCommentFound(),
		"appendLBC":         reqtype.AppendLinesByCompare.IsAppendLinesByCompare(),
		"appendLines":       reqtype.AppendLines.IsAppendLines(),
		"createOrAppend":    reqtype.CreateOrAppend.IsCreateOrAppend(),
		"prepend":           reqtype.Prepend.IsPrepend(),
		"createOrPrepend":   reqtype.CreateOrPrepend.IsCreateOrPrepend(),
		"prependLines":      reqtype.PrependLines.IsPrependLines(),
		"rename":            reqtype.Rename.IsRename(),
		"change":            reqtype.Change.IsChange(),
		"merge":             reqtype.Merge.IsMerge(),
		"mergeLines":        reqtype.MergeLines.IsMergeLines(),
	}

	// Assert
	expected := args.Map{
		"stopSleepStart":    true,
		"suspend":           true,
		"pause":             true,
		"resumed":           true,
		"tryRestart3":       true,
		"tryRestart5":       true,
		"tryStart3":         true,
		"tryStart5":         true,
		"tryStop3":          true,
		"tryStop5":          true,
		"inheritOnly":       true,
		"inheritPlusOvrd":   true,
		"createOrSkipExist": true,
		"updateSkipNon":     true,
		"deleteSkipNon":     true,
		"dropSkipNon":       true,
		"updateOnExist":     true,
		"dropOnExist":       true,
		"dropCreate":        true,
		"appendBC":          true,
		"appendBCW":         true,
		"appendLBC":         true,
		"appendLines":       true,
		"createOrAppend":    true,
		"prepend":           true,
		"createOrPrepend":   true,
		"prependLines":      true,
		"rename":            true,
		"change":            true,
		"merge":             true,
		"mergeLines":        true,
	}
	expected.ShouldBeEqual(t, 0, "Remaining returns correct value -- identity checks", actual)
}

// ── Logical group negatives ──

func Test_Request_LogicalGroupNegatives(t *testing.T) {
	// Act
	actual := args.Map{
		"readNotCreate":    reqtype.Read.IsCreateLogically(),
		"readNotDrop":      reqtype.Read.IsDropLogically(),
		"appendNotCrud":    reqtype.Append.IsCrudOnlyLogically(),
		"appendIsNotCrud":  reqtype.Append.IsNotCrudOnlyLogically(),
		"touchNotExist":    reqtype.Touch.IsOnExistCheckLogically(),
		"touchNotUpdate":   reqtype.Touch.IsUpdateOrRemoveLogically(),
		"touchNotOverride": reqtype.Touch.IsOverrideOrOverwriteOrEnforce(),
	}

	// Assert
	expected := args.Map{
		"readNotCreate":    false,
		"readNotDrop":      false,
		"appendNotCrud":    false,
		"appendIsNotCrud":  true,
		"touchNotExist":    false,
		"touchNotUpdate":   false,
		"touchNotOverride": false,
	}
	expected.ShouldBeEqual(t, 0, "Logical returns correct value -- group negatives", actual)
}

// ── Composite logical groups ──

func Test_Request_IsAnyAction_Values(t *testing.T) {
	// Act
	actual := args.Map{
		"stop":    reqtype.Stop.IsAnyAction(),
		"restart": reqtype.Restart.IsAnyAction(),
		"reload":  reqtype.Reload.IsAnyAction(),
		"suspend": reqtype.Suspend.IsAnyAction(),
		"pause":   reqtype.Pause.IsAnyAction(),
		"resumed": reqtype.Resumed.IsAnyAction(),
	}

	// Assert
	expected := args.Map{
		"stop":    true,
		"restart": true,
		"reload":  true,
		"suspend": true,
		"pause":   true,
		"resumed": true,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyAction returns non-empty -- values", actual)
}

func Test_Request_IsAnyHttp_Values(t *testing.T) {
	// Act
	actual := args.Map{
		"put":    reqtype.PutHttp.IsAnyHttp(),
		"post":   reqtype.PostHttp.IsAnyHttp(),
		"delete": reqtype.DeleteHttp.IsAnyHttp(),
		"patch":  reqtype.PatchHttp.IsAnyHttp(),
	}

	// Assert
	expected := args.Map{
		"put":    true,
		"post":   true,
		"delete": true,
		"patch":  true,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyHttp returns non-empty -- values", actual)
}

func Test_Request_IsAnyCreate_Values(t *testing.T) {
	// Act
	actual := args.Map{
		"createOrUpdate":  reqtype.CreateOrUpdate.IsAnyCreate(),
		"createOrAppend":  reqtype.CreateOrAppend.IsAnyCreate(),
		"createOrPrepend": reqtype.CreateOrPrepend.IsAnyCreate(),
		"createOrSkip":    reqtype.CreateOrSkipOnExist.IsAnyCreate(),
		"dropCreate":      reqtype.DropCreate.IsAnyCreate(),
	}

	// Assert
	expected := args.Map{
		"createOrUpdate":  true,
		"createOrAppend":  true,
		"createOrPrepend": true,
		"createOrSkip":    true,
		"dropCreate":      true,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyCreate returns non-empty -- values", actual)
}

func Test_Request_IsAnyDrop_Values(t *testing.T) {
	// Act
	actual := args.Map{
		"delete":      reqtype.Delete.IsAnyDrop(),
		"deleteSkip":  reqtype.DeleteOrSkipOnNonExist.IsAnyDrop(),
		"dropOnExist": reqtype.DropOnExist.IsAnyDrop(),
		"dropCreate":  reqtype.DropCreate.IsAnyDrop(),
		"dropSkip":    reqtype.DropOrSkipOnNonExist.IsAnyDrop(),
	}

	// Assert
	expected := args.Map{
		"delete":      true,
		"deleteSkip":  true,
		"dropOnExist": true,
		"dropCreate":  true,
		"dropSkip":    true,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyDrop returns non-empty -- values", actual)
}

func Test_Request_IsDropSafe_Values(t *testing.T) {
	// Act
	actual := args.Map{
		"deleteSkip": reqtype.DeleteOrSkipOnNonExist.IsDropSafe(),
		"dropSkip":   reqtype.DropOrSkipOnNonExist.IsDropSafe(),
	}

	// Assert
	expected := args.Map{
		"deleteSkip": true,
		"dropSkip":   true,
	}
	expected.ShouldBeEqual(t, 0, "IsDropSafe returns non-empty -- values", actual)
}

func Test_Request_IsAnySkipOnExist_Values(t *testing.T) {
	// Act
	actual := args.Map{
		"createSkip":   reqtype.CreateOrSkipOnExist.IsAnySkipOnExist(),
		"updateSkip":   reqtype.UpdateOrSkipOnNonExist.IsAnySkipOnExist(),
		"deleteSkip":   reqtype.DeleteOrSkipOnNonExist.IsAnySkipOnExist(),
		"dropSkip":     reqtype.DropOrSkipOnNonExist.IsAnySkipOnExist(),
	}

	// Assert
	expected := args.Map{
		"createSkip":   true,
		"updateSkip":   true,
		"deleteSkip":   true,
		"dropSkip":     true,
	}
	expected.ShouldBeEqual(t, 0, "IsAnySkipOnExist returns non-empty -- values", actual)
}
