package errcoretests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/errcore"
)

// ── Formatters (Coverage06 exported) ──

var varTwoTestCases = []coretestcases.CaseV1{
	{Title: "VarTwo returns non-empty -- with type", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "VarTwo returns non-empty -- without type", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "VarTwoNoType returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "VarThree returns non-empty -- with type", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "VarThree returns non-empty -- without type", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "VarThreeNoType returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
}

var varMapTestCases = []coretestcases.CaseV1{
	{Title: "VarMap returns empty -- nil map", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "VarMap returns non-empty -- with items", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "VarMapStrings returns empty -- nil map", ExpectedInput: args.Map{"len": 0}},
	{Title: "VarMapStrings returns 1 -- with items", ExpectedInput: args.Map{"len": 1}},
}

var messageVarTestCases = []coretestcases.CaseV1{
	{Title: "MessageVarTwo returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "MessageVarThree returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "MessageVarMap returns just msg -- nil map", ExpectedInput: args.Map{"result": "msg"}},
	{Title: "MessageVarMap returns non-empty -- with items", ExpectedInput: args.Map{"nonEmpty": true}},
}

var refTestCases = []coretestcases.CaseV1{
	{Title: "Ref returns empty -- nil ref", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "Ref returns non-empty -- with ref", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "RefToError returns nil -- nil ref", ExpectedInput: args.Map{"isNil": true}},
	{Title: "RefToError returns non-nil -- with ref", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "ToError returns nil -- empty msg", ExpectedInput: args.Map{"isNil": true}},
	{Title: "ToError returns non-nil -- with msg", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "ToString returns empty -- nil error", ExpectedInput: args.Map{"result": ""}},
	{Title: "ToString returns error text -- with error", ExpectedInput: args.Map{"result": "e"}},
}

var sourceDestTestCases = []coretestcases.CaseV1{
	{Title: "SourceDestination returns non-empty -- with type", ExpectedInput: args.Map{"nonEmpty1": true, "nonEmpty2": true}},
	{Title: "SourceDestinationNoType returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "SourceDestinationErr returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
}

var combineTestCases = []coretestcases.CaseV1{
	{Title: "Combine returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "CombineWithMsgTypeNoStack returns non-empty -- various combos", ExpectedInput: args.Map{"nonEmpty1": true, "nonEmpty2": true}},
	{Title: "CombineWithMsgTypeStackTrace returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "StackTracesCompiled returns non-empty -- with traces", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "GherkinsString returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "GherkinsStringWithExpectation returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
}

var rangeTestCases = []coretestcases.CaseV1{
	{Title: "RangeNotMeet returns non-empty -- with range", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "RangeNotMeet returns non-empty -- without range", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "PanicRangeNotMeet returns non-empty -- with range", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "PanicRangeNotMeet returns non-empty -- without range", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "EnumRangeNotMeet returns non-empty -- with range", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "EnumRangeNotMeet returns non-empty -- without range", ExpectedInput: args.Map{"nonEmpty": true}},
}

var msgHeaderTestCases = []coretestcases.CaseV1{
	{Title: "MsgHeader returns non-empty -- valid", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "MsgHeaderIf returns non-empty -- true", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "MsgHeaderIf returns non-empty -- false", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "MsgHeaderPlusEnding returns non-empty -- valid", ExpectedInput: args.Map{"nonEmpty": true}},
}

var sliceErrorTestCases = []coretestcases.CaseV1{
	{Title: "SliceError returns nil -- nil items", ExpectedInput: args.Map{"isNil": true}},
	{Title: "SliceError returns non-nil -- with items", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "SliceErrorDefault returns non-nil -- with items", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "SliceToError returns nil -- nil items", ExpectedInput: args.Map{"isNil": true}},
	{Title: "SliceToError returns non-nil -- with items", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "SliceToErrorPtr returns nil -- nil items", ExpectedInput: args.Map{"isNil": true}},
	{Title: "SliceToErrorPtr returns non-nil -- with items", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "SliceErrorsToStrings returns empty -- nil", ExpectedInput: args.Map{"len": 0}},
	{Title: "SliceErrorsToStrings returns 2 -- mixed nil and errors", ExpectedInput: args.Map{"len": 2}},
	{Title: "ManyErrorToSingle returns non-nil -- with errors", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "ManyErrorToSingleDirect returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "MergeErrors returns non-nil -- two errors", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "MergeErrorsToString returns empty -- nil", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "MergeErrorsToString returns non-empty -- with error", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "MergeErrorsToStringDefault returns empty -- nil", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "MergeErrorsToStringDefault returns non-empty -- with error", ExpectedInput: args.Map{"nonEmpty": true}},
}

var stringLinesTestCases = []coretestcases.CaseV1{
	{Title: "StringLinesToQuoteLines returns empty -- nil", ExpectedInput: args.Map{"len": 0}},
	{Title: "StringLinesToQuoteLines returns 1 -- with items", ExpectedInput: args.Map{"len": 1}},
	{Title: "StringLinesToQuoteLinesToSingle returns non-empty -- with items", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "LinesToDoubleQuoteLinesWithTabs returns empty -- nil", ExpectedInput: args.Map{"len": 0}},
	{Title: "LinesToDoubleQuoteLinesWithTabs returns 1 -- with items", ExpectedInput: args.Map{"len": 1}},
}

var debugPrintTestCases = []coretestcases.CaseV1{
	{Title: "FmtDebug does not panic -- valid args", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "FmtDebugIf does not panic -- false", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "FmtDebugIf does not panic -- true", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "ValidPrint does not panic -- both combos", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "FailedPrint does not panic -- both combos", ExpectedInput: args.Map{"noPanic": true}},
}

var expectedButTestCases = []coretestcases.CaseV1{
	{Title: "Expected.But returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "Expected.ButFoundAsMsg returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "Expected.ButFoundWithTypeAsMsg returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "Expected.ButUsingType returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
}

var shouldBeTestCases = []coretestcases.CaseV1{
	{Title: "ShouldBe.StrEqMsg returns non-empty -- diff strings", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ShouldBe.StrEqErr returns non-nil -- diff strings", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "ShouldBe.AnyEqMsg returns non-empty -- diff values", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ShouldBe.AnyEqErr returns non-nil -- diff values", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "ShouldBe.JsonEqMsg returns non-empty -- diff values", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ShouldBe.JsonEqErr returns non-nil -- diff values", ExpectedInput: args.Map{"nonNil": true}},
}

// ── RawErrorType (Coverage07) ──

var rawErrorTypeTestCases = []coretestcases.CaseV1{
	{Title: "RawErrorType.String returns non-empty -- InvalidType", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "RawErrorType.Combine returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "RawErrorType.CombineWithAnother returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "RawErrorType.TypesAttach returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "RawErrorType.TypesAttachErr returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.SrcDestination returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "RawErrorType.SrcDestinationErr returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.Error returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.ErrorSkip returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.Fmt returns non-nil -- empty format", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.Fmt returns non-nil -- with format", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.FmtIf returns nil -- false condition", ExpectedInput: args.Map{"isNil": true}},
	{Title: "RawErrorType.FmtIf returns non-nil -- true condition", ExpectedInput: args.Map{"nonNil": true}},
}

var rawErrorTypeMergeTestCases = []coretestcases.CaseV1{
	{Title: "RawErrorType.MergeError returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "RawErrorType.MergeError returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.MergeErrorWithMessage returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "RawErrorType.MergeErrorWithMessage returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.MergeErrorWithMessageRef returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "RawErrorType.MergeErrorWithMessageRef returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.MergeErrorWithRef returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "RawErrorType.MergeErrorWithRef returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
}

var rawErrorTypeMsgTestCases = []coretestcases.CaseV1{
	{Title: "RawErrorType.MsgCsvRef returns non-empty -- msg only", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "RawErrorType.MsgCsvRef returns non-empty -- empty msg with ref", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "RawErrorType.MsgCsvRef returns non-empty -- msg with refs", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "RawErrorType.MsgCsvRefError returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.ErrorRefOnly returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.Expecting returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.NoRef returns non-empty -- empty msg", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "RawErrorType.NoRef returns non-empty -- with msg", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "RawErrorType.ErrorNoRefs returns non-nil -- with msg", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.ErrorNoRefs returns non-nil -- empty msg", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.ErrorNoRefsSkip returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrorType.HandleUsingPanic panics -- valid args", ExpectedInput: args.Map{"panics": true}},
}

var getSetTestCases = []coretestcases.CaseV1{
	{Title: "GetSet returns first -- true condition", ExpectedInput: args.Map{"result": errcore.InvalidType}},
	{Title: "GetSet returns second -- false condition", ExpectedInput: args.Map{"result": errcore.NotFound}},
	{Title: "GetSetVariant returns first -- true condition", ExpectedInput: args.Map{"result": "a"}},
	{Title: "GetSetVariant returns second -- false condition", ExpectedInput: args.Map{"result": "b"}},
}

var meaningfulErrorTestCases = []coretestcases.CaseV1{
	{Title: "MeaningfulError returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "MeaningfulError returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "MeaningfulErrorHandle does not panic -- nil error", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "MeaningfulErrorHandle panics -- with error", ExpectedInput: args.Map{"panics": true}},
	{Title: "MeaningfulErrorWithData returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "MeaningfulErrorWithData returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "MeaningfulMessageError returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "MeaningfulMessageError returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "PathMeaningfulMessage returns nil -- no messages", ExpectedInput: args.Map{"isNil": true}},
	{Title: "PathMeaningfulMessage returns non-nil -- with messages", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "PathMeaningfulError returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "PathMeaningfulError returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
}

// ── StackEnhance (Coverage08) ──

var stackEnhanceTestCases = []coretestcases.CaseV1{
	{Title: "StackEnhance.Error returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "StackEnhance.Error returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "StackEnhance.ErrorSkip returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "StackEnhance.ErrorSkip returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "StackEnhance.Msg returns empty -- empty msg", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "StackEnhance.Msg returns non-empty -- with msg", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "StackEnhance.MsgSkip returns empty -- empty msg", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "StackEnhance.MsgSkip returns non-empty -- with msg", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "StackEnhance.MsgSkip returns non-empty -- already has stack trace", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "StackEnhance.MsgToErrSkip returns nil -- empty msg", ExpectedInput: args.Map{"isNil": true}},
	{Title: "StackEnhance.MsgToErrSkip returns non-nil -- with msg", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "StackEnhance.FmtSkip returns nil -- empty format", ExpectedInput: args.Map{"isNil": true}},
	{Title: "StackEnhance.FmtSkip returns non-nil -- with format", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "StackEnhance.MsgErrorSkip returns empty -- nil error", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "StackEnhance.MsgErrorSkip returns non-empty -- with error", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "StackEnhance.MsgErrorSkip returns non-empty -- already has stack", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "StackEnhance.MsgErrorToErrSkip returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "StackEnhance.MsgErrorToErrSkip returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
}

var countStateChangeTrackerTestCases = []coretestcases.CaseV1{
	{Title: "CountStateChangeTracker returns correct state -- initial and after change", ExpectedInput: args.Map{
		"isSame": true, "isValid": true, "isSuccess": true,
		"noChanges": true, "notFailed": true, "sameWithZero": true,
		"changedAfterAdd": true, "hasChanges": true, "isFailed": true,
	}},
}

// ── VarNameValues / MessageNameValues ──

var varNameValuesTestCases = []coretestcases.CaseV1{
	{Title: "VarNameValues returns empty -- no items", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "VarNameValues returns non-empty -- with items", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "VarNameValuesJoiner returns empty -- no items", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "VarNameValuesJoiner returns non-empty -- with items", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "VarNameValuesStrings returns empty -- no items", ExpectedInput: args.Map{"len": 0}},
	{Title: "VarNameValuesStrings returns 1 -- with items", ExpectedInput: args.Map{"len": 1}},
	{Title: "MessageNameValues returns just msg -- no items", ExpectedInput: args.Map{"result": "msg"}},
	{Title: "MessageNameValues returns non-empty -- with items", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "MessageWithRef returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "MessageWithRefToError returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
}

var toStringPtrTestCases = []coretestcases.CaseV1{
	{Title: "ToStringPtr returns empty ptr -- nil error", ExpectedInput: args.Map{"ptrEmpty": true}},
	{Title: "ToStringPtr returns error ptr -- with error", ExpectedInput: args.Map{"ptrValue": "e"}},
	{Title: "ToValueString returns non-empty -- valid input", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ToExitError returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "ToExitError returns nil -- non-exit error", ExpectedInput: args.Map{"isNil": true}},
}

var getActualExpectTestCases = []coretestcases.CaseV1{
	{Title: "GetActualAndExpectProcessedMessage returns non-empty -- valid", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "GetSearchLineNumberExpectationMessage returns non-empty -- valid", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "GetSearchTermExpectationMessage returns non-empty -- with info", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "GetSearchTermExpectationMessage returns non-empty -- nil info", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "GetSearchTermExpectationSimpleMessage returns non-empty -- valid", ExpectedInput: args.Map{"nonEmpty": true}},
}

var expectedReflectTestCases = []coretestcases.CaseV1{
	{Title: "Expected.ReflectButFound returns non-nil -- valid kinds", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "Expected.PrimitiveButFound returns non-nil -- map kind", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "Expected.ValueHasNoElements returns non-nil -- slice kind", ExpectedInput: args.Map{"nonNil": true}},
}
