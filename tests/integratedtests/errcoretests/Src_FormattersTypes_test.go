package errcoretests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/namevalue"
)

// ══════════════════════════════════════════════════════════════════════════════
// Formatters (Coverage06 - exported only)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Src06_VarTwo_WithType(t *testing.T) {
	tc := varTwoTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarTwo(true, "a", 1, "b", 2) != ""})
}

func Test_Src06_VarTwo_WithoutType(t *testing.T) {
	tc := varTwoTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarTwo(false, "a", 1, "b", 2) != ""})
}

func Test_Src06_VarTwoNoType(t *testing.T) {
	tc := varTwoTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarTwoNoType("a", 1, "b", 2) != ""})
}

func Test_Src06_VarThree_WithType(t *testing.T) {
	tc := varTwoTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarThree(true, "a", 1, "b", 2, "c", 3) != ""})
}

func Test_Src06_VarThree_WithoutType(t *testing.T) {
	tc := varTwoTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarThree(false, "a", 1, "b", 2, "c", 3) != ""})
}

func Test_Src06_VarThreeNoType(t *testing.T) {
	tc := varTwoTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarThreeNoType("a", 1, "b", 2, "c", 3) != ""})
}

func Test_Src06_VarMap_Empty(t *testing.T) {
	tc := varMapTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.VarMap(nil) == ""})
}

func Test_Src06_VarMap_WithItems(t *testing.T) {
	tc := varMapTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarMap(map[string]any{"a": 1}) != ""})
}

func Test_Src06_VarMapStrings_Empty(t *testing.T) {
	tc := varMapTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.VarMapStrings(nil))})
}

func Test_Src06_VarMapStrings_WithItems(t *testing.T) {
	tc := varMapTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.VarMapStrings(map[string]any{"a": 1}))})
}

func Test_Src06_VarNameValues_Empty(t *testing.T) {
	tc := varNameValuesTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.VarNameValues() == ""})
}

func Test_Src06_VarNameValues_WithItems(t *testing.T) {
	tc := varNameValuesTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarNameValues(namevalue.StringAny{Name: "a", Value: 1}) != ""})
}

func Test_Src06_VarNameValuesJoiner_Empty(t *testing.T) {
	tc := varNameValuesTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.VarNameValuesJoiner(",") == ""})
}

func Test_Src06_VarNameValuesJoiner_WithItems(t *testing.T) {
	tc := varNameValuesTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarNameValuesJoiner(",", namevalue.StringAny{Name: "a", Value: 1}) != ""})
}

func Test_Src06_VarNameValuesStrings_Empty(t *testing.T) {
	tc := varNameValuesTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.VarNameValuesStrings())})
}

func Test_Src06_VarNameValuesStrings_WithItems(t *testing.T) {
	tc := varNameValuesTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.VarNameValuesStrings(namevalue.StringAny{Name: "a", Value: 1}))})
}

func Test_Src06_MessageVarTwo(t *testing.T) {
	tc := messageVarTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MessageVarTwo("msg", "a", 1, "b", 2) != ""})
}

func Test_Src06_MessageVarThree(t *testing.T) {
	tc := messageVarTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3) != ""})
}

func Test_Src06_MessageVarMap_Empty(t *testing.T) {
	tc := messageVarTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.MessageVarMap("msg", nil)})
}

func Test_Src06_MessageVarMap_WithItems(t *testing.T) {
	tc := messageVarTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MessageVarMap("msg", map[string]any{"a": 1}) != ""})
}

func Test_Src06_MessageNameValues_Empty(t *testing.T) {
	tc := varNameValuesTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.MessageNameValues("msg")})
}

func Test_Src06_MessageNameValues_WithItems(t *testing.T) {
	tc := varNameValuesTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MessageNameValues("msg", namevalue.StringAny{Name: "a", Value: 1}) != ""})
}

func Test_Src06_MessageWithRef(t *testing.T) {
	tc := varNameValuesTestCases[8]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MessageWithRef("msg", "ref") != ""})
}

func Test_Src06_MessageWithRefToError(t *testing.T) {
	tc := varNameValuesTestCases[9]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.MessageWithRefToError("msg", "ref") != nil})
}

func Test_Src06_Ref_Nil(t *testing.T) {
	tc := refTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.Ref(nil) == ""})
}

func Test_Src06_Ref_WithRef(t *testing.T) {
	tc := refTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.Ref("ref") != ""})
}

func Test_Src06_RefToError_Nil(t *testing.T) {
	tc := refTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.RefToError(nil) == nil})
}

func Test_Src06_RefToError_WithRef(t *testing.T) {
	tc := refTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.RefToError("ref") != nil})
}

func Test_Src06_ToError_Empty(t *testing.T) {
	tc := refTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.ToError("") == nil})
}

func Test_Src06_ToError_WithMsg(t *testing.T) {
	tc := refTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.ToError("msg") != nil})
}

func Test_Src06_ToString_Nil(t *testing.T) {
	tc := refTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.ToString(nil)})
}

func Test_Src06_ToString_WithErr(t *testing.T) {
	tc := refTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.ToString(errors.New("e"))})
}

func Test_Src06_ToStringPtr_Nil(t *testing.T) {
	tc := toStringPtrTestCases[0]
	p := errcore.ToStringPtr(nil)

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"ptrEmpty": p != nil && *p == ""})
}

func Test_Src06_ToStringPtr_WithErr(t *testing.T) {
	tc := toStringPtrTestCases[1]
	p := errcore.ToStringPtr(errors.New("e"))

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"ptrValue": *p})
}

func Test_Src06_ToValueString(t *testing.T) {
	tc := toStringPtrTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.ToValueString("hello") != ""})
}

func Test_Src06_ToExitError_Nil(t *testing.T) {
	tc := toStringPtrTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.ToExitError(nil) == nil})
}

func Test_Src06_ToExitError_NonExit(t *testing.T) {
	tc := toStringPtrTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.ToExitError(errors.New("e")) == nil})
}

func Test_Src06_SourceDestination(t *testing.T) {
	tc := sourceDestTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"nonEmpty1": errcore.SourceDestination(true, "src", "dst") != "",
		"nonEmpty2": errcore.SourceDestination(false, "src", "dst") != "",
	})
}

func Test_Src06_SourceDestinationNoType(t *testing.T) {
	tc := sourceDestTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.SourceDestinationNoType("src", "dst") != ""})
}

func Test_Src06_SourceDestinationErr(t *testing.T) {
	tc := sourceDestTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.SourceDestinationErr(true, "src", "dst") != nil})
}

func Test_Src06_Combine(t *testing.T) {
	tc := combineTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.Combine("gen", "other", "ref") != ""})
}

func Test_Src06_CombineWithMsgTypeNoStack(t *testing.T) {
	tc := combineTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"nonEmpty1": errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "", nil) != "",
		"nonEmpty2": errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "msg", "ref") != "",
	})
}

func Test_Src06_CombineWithMsgTypeStackTrace(t *testing.T) {
	tc := combineTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.CombineWithMsgTypeStackTrace(errcore.InvalidType, "msg", nil) != ""})
}

func Test_Src06_StackTracesCompiled(t *testing.T) {
	tc := combineTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StackTracesCompiled([]string{"a", "b"}) != ""})
}

func Test_Src06_GherkinsString(t *testing.T) {
	tc := combineTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GherkinsString(0, "f", "g", "w", "t") != ""})
}

func Test_Src06_GherkinsStringWithExpectation(t *testing.T) {
	tc := combineTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GherkinsStringWithExpectation(0, "f", "g", "w", "t", "a", "e") != ""})
}

func Test_Src06_RangeNotMeet_WithRange(t *testing.T) {
	tc := rangeTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.RangeNotMeet("msg", 0, 10, []int{1, 2}) != ""})
}

func Test_Src06_RangeNotMeet_WithoutRange(t *testing.T) {
	tc := rangeTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.RangeNotMeet("msg", 0, 10, nil) != ""})
}

func Test_Src06_PanicRangeNotMeet_WithRange(t *testing.T) {
	tc := rangeTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.PanicRangeNotMeet("msg", 0, 10, []int{1}) != ""})
}

func Test_Src06_PanicRangeNotMeet_WithoutRange(t *testing.T) {
	tc := rangeTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.PanicRangeNotMeet("msg", 0, 10, nil) != ""})
}

func Test_Src06_EnumRangeNotMeet_WithRange(t *testing.T) {
	tc := rangeTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.EnumRangeNotMeet(0, 10, "range") != ""})
}

func Test_Src06_EnumRangeNotMeet_WithoutRange(t *testing.T) {
	tc := rangeTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.EnumRangeNotMeet(0, 10, nil) != ""})
}

func Test_Src06_MsgHeader(t *testing.T) {
	tc := msgHeaderTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MsgHeader("test") != ""})
}

func Test_Src06_MsgHeaderIf_True(t *testing.T) {
	tc := msgHeaderTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MsgHeaderIf(true, "test") != ""})
}

func Test_Src06_MsgHeaderIf_False(t *testing.T) {
	tc := msgHeaderTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MsgHeaderIf(false, "test") != ""})
}

func Test_Src06_MsgHeaderPlusEnding(t *testing.T) {
	tc := msgHeaderTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MsgHeaderPlusEnding("header", "msg") != ""})
}

func Test_Src06_SliceError_Nil(t *testing.T) {
	tc := sliceErrorTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.SliceError(",", nil) == nil})
}

func Test_Src06_SliceError_WithItems(t *testing.T) {
	tc := sliceErrorTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.SliceError(",", []string{"a", "b"}) != nil})
}

func Test_Src06_SliceErrorDefault(t *testing.T) {
	tc := sliceErrorTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.SliceErrorDefault([]string{"a"}) != nil})
}

func Test_Src06_SliceToError_Nil(t *testing.T) {
	tc := sliceErrorTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.SliceToError(nil) == nil})
}

func Test_Src06_SliceToError_WithItems(t *testing.T) {
	tc := sliceErrorTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.SliceToError([]string{"a"}) != nil})
}

func Test_Src06_SliceToErrorPtr_Nil(t *testing.T) {
	tc := sliceErrorTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.SliceToErrorPtr(nil) == nil})
}

func Test_Src06_SliceToErrorPtr_WithItems(t *testing.T) {
	tc := sliceErrorTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.SliceToErrorPtr([]string{"a"}) != nil})
}

func Test_Src06_SliceErrorsToStrings_Nil(t *testing.T) {
	tc := sliceErrorTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.SliceErrorsToStrings(nil...))})
}

func Test_Src06_SliceErrorsToStrings_WithItems(t *testing.T) {
	tc := sliceErrorTestCases[8]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.SliceErrorsToStrings(errors.New("a"), nil, errors.New("b")))})
}

func Test_Src06_ManyErrorToSingle(t *testing.T) {
	tc := sliceErrorTestCases[9]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.ManyErrorToSingle([]error{errors.New("a"), nil}) != nil})
}

func Test_Src06_ManyErrorToSingleDirect(t *testing.T) {
	tc := sliceErrorTestCases[10]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.ManyErrorToSingleDirect(errors.New("a")) != nil})
}

func Test_Src06_MergeErrors(t *testing.T) {
	tc := sliceErrorTestCases[11]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.MergeErrors(errors.New("a"), errors.New("b")) != nil})
}

func Test_Src06_MergeErrorsToString_Nil(t *testing.T) {
	tc := sliceErrorTestCases[12]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.MergeErrorsToString(",", nil...) == ""})
}

func Test_Src06_MergeErrorsToString_WithItems(t *testing.T) {
	tc := sliceErrorTestCases[13]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MergeErrorsToString(",", errors.New("a")) != ""})
}

func Test_Src06_MergeErrorsToStringDefault_Nil(t *testing.T) {
	tc := sliceErrorTestCases[14]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.MergeErrorsToStringDefault(nil...) == ""})
}

func Test_Src06_MergeErrorsToStringDefault_WithItems(t *testing.T) {
	tc := sliceErrorTestCases[15]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MergeErrorsToStringDefault(errors.New("a")) != ""})
}

func Test_Src06_StringLinesToQuoteLines_Empty(t *testing.T) {
	tc := stringLinesTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.StringLinesToQuoteLines(nil))})
}

func Test_Src06_StringLinesToQuoteLines_WithItems(t *testing.T) {
	tc := stringLinesTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.StringLinesToQuoteLines([]string{"a"}))})
}

func Test_Src06_StringLinesToQuoteLinesToSingle(t *testing.T) {
	tc := stringLinesTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"}) != ""})
}

func Test_Src06_LinesToDoubleQuoteLinesWithTabs_Empty(t *testing.T) {
	tc := stringLinesTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.LinesToDoubleQuoteLinesWithTabs(2, nil))})
}

func Test_Src06_LinesToDoubleQuoteLinesWithTabs_WithItems(t *testing.T) {
	tc := stringLinesTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.LinesToDoubleQuoteLinesWithTabs(4, []string{"a"}))})
}

func Test_Src06_FmtDebug(t *testing.T) {
	tc := debugPrintTestCases[0]
	noPanic := !callPanicsErrcore(func() { errcore.FmtDebug("test %s", "v") })

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src06_FmtDebugIf_False(t *testing.T) {
	tc := debugPrintTestCases[1]
	noPanic := !callPanicsErrcore(func() { errcore.FmtDebugIf(false, "skip") })

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src06_FmtDebugIf_True(t *testing.T) {
	tc := debugPrintTestCases[2]
	noPanic := !callPanicsErrcore(func() { errcore.FmtDebugIf(true, "test %s", "v") })

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src06_ValidPrint(t *testing.T) {
	tc := debugPrintTestCases[3]
	noPanic := !callPanicsErrcore(func() {
		errcore.ValidPrint(false, "skip")
		errcore.ValidPrint(true, "show")
	})

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src06_FailedPrint(t *testing.T) {
	tc := debugPrintTestCases[4]
	noPanic := !callPanicsErrcore(func() {
		errcore.FailedPrint(false, "skip")
		errcore.FailedPrint(true, "show")
	})

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src06_GetActualAndExpectProcessedMessage(t *testing.T) {
	tc := getActualExpectTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GetActualAndExpectProcessedMessage(0, "a", "e", "ap", "ep") != ""})
}

func Test_Src06_GetSearchLineNumberExpectationMessage(t *testing.T) {
	tc := getActualExpectTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GetSearchLineNumberExpectationMessage(0, 1, 2, "content", "search", "info") != ""})
}

func Test_Src06_GetSearchTermExpectationMessage_WithInfo(t *testing.T) {
	tc := getActualExpectTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GetSearchTermExpectationMessage(0, "h", "e", 1, "a", "e", "info") != ""})
}

func Test_Src06_GetSearchTermExpectationMessage_NilInfo(t *testing.T) {
	tc := getActualExpectTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GetSearchTermExpectationMessage(0, "h", "e", 1, "a", "e", nil) != ""})
}

func Test_Src06_GetSearchTermExpectationSimpleMessage(t *testing.T) {
	tc := getActualExpectTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GetSearchTermExpectationSimpleMessage(0, "e", 1, "c", "s") != ""})
}

func Test_Src06_Expected_But(t *testing.T) {
	tc := expectedButTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.Expected.But("t", "e", "a") != nil})
}

func Test_Src06_Expected_ButFoundAsMsg(t *testing.T) {
	tc := expectedButTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.Expected.ButFoundAsMsg("t", "e", "a") != ""})
}

func Test_Src06_Expected_ButFoundWithTypeAsMsg(t *testing.T) {
	tc := expectedButTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.Expected.ButFoundWithTypeAsMsg("t", "e", "a") != ""})
}

func Test_Src06_Expected_ButUsingType(t *testing.T) {
	tc := expectedButTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.Expected.ButUsingType("t", "e", "a") != nil})
}

func Test_Src06_Expected_ReflectButFound(t *testing.T) {
	tc := expectedReflectTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.Expected.ReflectButFound(reflect.Int, reflect.String) != nil})
}

func Test_Src06_Expected_PrimitiveButFound(t *testing.T) {
	tc := expectedReflectTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.Expected.PrimitiveButFound(reflect.Map) != nil})
}

func Test_Src06_Expected_ValueHasNoElements(t *testing.T) {
	tc := expectedReflectTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.Expected.ValueHasNoElements(reflect.Slice) != nil})
}

func Test_Src06_ShouldBe_StrEqMsg(t *testing.T) {
	tc := shouldBeTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.ShouldBe.StrEqMsg("a", "b") != ""})
}

func Test_Src06_ShouldBe_StrEqErr(t *testing.T) {
	tc := shouldBeTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.ShouldBe.StrEqErr("a", "b") != nil})
}

func Test_Src06_ShouldBe_AnyEqMsg(t *testing.T) {
	tc := shouldBeTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.ShouldBe.AnyEqMsg(1, 2) != ""})
}

func Test_Src06_ShouldBe_AnyEqErr(t *testing.T) {
	tc := shouldBeTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.ShouldBe.AnyEqErr(1, 2) != nil})
}

func Test_Src06_ShouldBe_JsonEqMsg(t *testing.T) {
	tc := shouldBeTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.ShouldBe.JsonEqMsg("a", "b") != ""})
}

func Test_Src06_ShouldBe_JsonEqErr(t *testing.T) {
	tc := shouldBeTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.ShouldBe.JsonEqErr("a", "b") != nil})
}

// ══════════════════════════════════════════════════════════════════════════════
// RawErrorType (Coverage07)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Src07_RawErrorType_String(t *testing.T) {
	tc := rawErrorTypeTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.String() != ""})
}

func Test_Src07_RawErrorType_Combine(t *testing.T) {
	tc := rawErrorTypeTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.Combine("msg", "ref") != ""})
}

func Test_Src07_RawErrorType_CombineWithAnother(t *testing.T) {
	tc := rawErrorTypeTestCases[2]
	r := errcore.InvalidType.CombineWithAnother(errcore.NotFound, "msg", "ref")

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": r.String() != ""})
}

func Test_Src07_RawErrorType_TypesAttach(t *testing.T) {
	tc := rawErrorTypeTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.TypesAttach("msg", "hello", 42) != ""})
}

func Test_Src07_RawErrorType_TypesAttachErr(t *testing.T) {
	tc := rawErrorTypeTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.TypesAttachErr("msg", "hello") != nil})
}

func Test_Src07_RawErrorType_SrcDestination(t *testing.T) {
	tc := rawErrorTypeTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.SrcDestination("msg", "src", "sv", "dst", "dv") != ""})
}

func Test_Src07_RawErrorType_SrcDestinationErr(t *testing.T) {
	tc := rawErrorTypeTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.SrcDestinationErr("msg", "src", "sv", "dst", "dv") != nil})
}

func Test_Src07_RawErrorType_Error(t *testing.T) {
	tc := rawErrorTypeTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.Error("msg", "ref") != nil})
}

func Test_Src07_RawErrorType_ErrorSkip(t *testing.T) {
	tc := rawErrorTypeTestCases[8]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.ErrorSkip(0, "msg", "ref") != nil})
}

func Test_Src07_RawErrorType_Fmt_Empty(t *testing.T) {
	tc := rawErrorTypeTestCases[9]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.Fmt("") != nil})
}

func Test_Src07_RawErrorType_Fmt_WithFormat(t *testing.T) {
	tc := rawErrorTypeTestCases[10]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.Fmt("val=%d", 42) != nil})
}

func Test_Src07_RawErrorType_FmtIf_False(t *testing.T) {
	tc := rawErrorTypeTestCases[11]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.InvalidType.FmtIf(false, "val=%d", 42) == nil})
}

func Test_Src07_RawErrorType_FmtIf_True(t *testing.T) {
	tc := rawErrorTypeTestCases[12]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.FmtIf(true, "val=%d", 42) != nil})
}

func Test_Src07_RawErrorType_MergeError_Nil(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.InvalidType.MergeError(nil) == nil})
}

func Test_Src07_RawErrorType_MergeError_WithErr(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.MergeError(errors.New("e")) != nil})
}

func Test_Src07_RawErrorType_MergeErrorWithMessage_Nil(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.InvalidType.MergeErrorWithMessage(nil, "msg") == nil})
}

func Test_Src07_RawErrorType_MergeErrorWithMessage_WithErr(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.MergeErrorWithMessage(errors.New("e"), "msg") != nil})
}

func Test_Src07_RawErrorType_MergeErrorWithMessageRef_Nil(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.InvalidType.MergeErrorWithMessageRef(nil, "msg", "ref") == nil})
}

func Test_Src07_RawErrorType_MergeErrorWithMessageRef_WithErr(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.MergeErrorWithMessageRef(errors.New("e"), "msg", "ref") != nil})
}

func Test_Src07_RawErrorType_MergeErrorWithRef_Nil(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.InvalidType.MergeErrorWithRef(nil, "ref") == nil})
}

func Test_Src07_RawErrorType_MergeErrorWithRef_WithErr(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.MergeErrorWithRef(errors.New("e"), "ref") != nil})
}

func Test_Src07_RawErrorType_MsgCsvRef_MsgOnly(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.MsgCsvRef("msg") != ""})
}

func Test_Src07_RawErrorType_MsgCsvRef_EmptyMsgWithRef(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.MsgCsvRef("", "ref") != ""})
}

func Test_Src07_RawErrorType_MsgCsvRef_MsgWithRefs(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.MsgCsvRef("msg", "r1", "r2") != ""})
}

func Test_Src07_RawErrorType_MsgCsvRefError(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.MsgCsvRefError("msg", "r1") != nil})
}

func Test_Src07_RawErrorType_ErrorRefOnly(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.ErrorRefOnly("ref") != nil})
}

func Test_Src07_RawErrorType_Expecting(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.Expecting("exp", "act") != nil})
}

func Test_Src07_RawErrorType_NoRef_EmptyMsg(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.NoRef("") != ""})
}

func Test_Src07_RawErrorType_NoRef_WithMsg(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.NoRef("msg") != ""})
}

func Test_Src07_RawErrorType_ErrorNoRefs_WithMsg(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[8]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.ErrorNoRefs("msg") != nil})
}

func Test_Src07_RawErrorType_ErrorNoRefs_EmptyMsg(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[9]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.ErrorNoRefs("") != nil})
}

func Test_Src07_RawErrorType_ErrorNoRefsSkip(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[10]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.ErrorNoRefsSkip(0, "msg") != nil})
}

func Test_Src07_RawErrorType_HandleUsingPanic(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[11]
	panics := callPanicsErrcore(func() { errcore.InvalidType.HandleUsingPanic("msg", "ref") })

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"panics": panics})
}

func Test_Src07_GetSet_True(t *testing.T) {
	tc := getSetTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.GetSet(true, errcore.InvalidType, errcore.NotFound)})
}

func Test_Src07_GetSet_False(t *testing.T) {
	tc := getSetTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.GetSet(false, errcore.InvalidType, errcore.NotFound)})
}

func Test_Src07_GetSetVariant_True(t *testing.T) {
	tc := getSetTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.GetSetVariant(true, "a", "b")})
}

func Test_Src07_GetSetVariant_False(t *testing.T) {
	tc := getSetTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.GetSetVariant(false, "a", "b")})
}

func Test_Src07_MeaningfulError_Nil(t *testing.T) {
	tc := meaningfulErrorTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.MeaningfulError(errcore.InvalidType, "fn", nil) == nil})
}

func Test_Src07_MeaningfulError_WithErr(t *testing.T) {
	tc := meaningfulErrorTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.MeaningfulError(errcore.InvalidType, "fn", errors.New("e")) != nil})
}

func Test_Src07_MeaningfulErrorHandle_Nil(t *testing.T) {
	tc := meaningfulErrorTestCases[2]
	noPanic := !callPanicsErrcore(func() { errcore.MeaningfulErrorHandle(errcore.InvalidType, "fn", nil) })

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src07_MeaningfulErrorHandle_WithErr(t *testing.T) {
	tc := meaningfulErrorTestCases[3]
	panics := callPanicsErrcore(func() { errcore.MeaningfulErrorHandle(errcore.InvalidType, "fn", errors.New("e")) })

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"panics": panics})
}

func Test_Src07_MeaningfulErrorWithData_Nil(t *testing.T) {
	tc := meaningfulErrorTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.MeaningfulErrorWithData(errcore.InvalidType, "fn", nil, "data") == nil})
}

func Test_Src07_MeaningfulErrorWithData_WithErr(t *testing.T) {
	tc := meaningfulErrorTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.MeaningfulErrorWithData(errcore.InvalidType, "fn", errors.New("e"), "data") != nil})
}

func Test_Src07_MeaningfulMessageError_Nil(t *testing.T) {
	tc := meaningfulErrorTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.MeaningfulMessageError(errcore.InvalidType, "fn", nil, "msg") == nil})
}

func Test_Src07_MeaningfulMessageError_WithErr(t *testing.T) {
	tc := meaningfulErrorTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.MeaningfulMessageError(errcore.InvalidType, "fn", errors.New("e"), "msg") != nil})
}

func Test_Src07_PathMeaningfulMessage_Empty(t *testing.T) {
	tc := meaningfulErrorTestCases[8]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.PathMeaningfulMessage(errcore.InvalidType, "fn", "/path") == nil})
}

func Test_Src07_PathMeaningfulMessage_WithMsgs(t *testing.T) {
	tc := meaningfulErrorTestCases[9]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.PathMeaningfulMessage(errcore.InvalidType, "fn", "/path", "a", "b") != nil})
}

func Test_Src07_PathMeaningfulError_Nil(t *testing.T) {
	tc := meaningfulErrorTestCases[10]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.PathMeaningfulError(errcore.InvalidType, nil, "/path") == nil})
}

func Test_Src07_PathMeaningfulError_WithErr(t *testing.T) {
	tc := meaningfulErrorTestCases[11]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.PathMeaningfulError(errcore.InvalidType, errors.New("e"), "/path") != nil})
}

// ══════════════════════════════════════════════════════════════════════════════
// StackEnhance (Coverage08)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Src08_StackEnhance_Error_Nil(t *testing.T) {
	tc := stackEnhanceTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.StackEnhance.Error(nil) == nil})
}

func Test_Src08_StackEnhance_Error_WithErr(t *testing.T) {
	tc := stackEnhanceTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.StackEnhance.Error(errors.New("e")) != nil})
}

func Test_Src08_StackEnhance_ErrorSkip_Nil(t *testing.T) {
	tc := stackEnhanceTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.StackEnhance.ErrorSkip(0, nil) == nil})
}

func Test_Src08_StackEnhance_ErrorSkip_WithErr(t *testing.T) {
	tc := stackEnhanceTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.StackEnhance.ErrorSkip(0, errors.New("e")) != nil})
}

func Test_Src08_StackEnhance_Msg_Empty(t *testing.T) {
	tc := stackEnhanceTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.StackEnhance.Msg("") == ""})
}

func Test_Src08_StackEnhance_Msg_WithMsg(t *testing.T) {
	tc := stackEnhanceTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StackEnhance.Msg("hello") != ""})
}

func Test_Src08_StackEnhance_MsgSkip_Empty(t *testing.T) {
	tc := stackEnhanceTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.StackEnhance.MsgSkip(0, "") == ""})
}

func Test_Src08_StackEnhance_MsgSkip_WithMsg(t *testing.T) {
	tc := stackEnhanceTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StackEnhance.MsgSkip(0, "hello") != ""})
}

func Test_Src08_StackEnhance_MsgSkip_AlreadyHasStackTrace(t *testing.T) {
	tc := stackEnhanceTestCases[8]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StackEnhance.MsgSkip(0, "hello Stack Trace: existing") != ""})
}

func Test_Src08_StackEnhance_MsgToErrSkip_Empty(t *testing.T) {
	tc := stackEnhanceTestCases[9]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.StackEnhance.MsgToErrSkip(0, "") == nil})
}

func Test_Src08_StackEnhance_MsgToErrSkip_WithMsg(t *testing.T) {
	tc := stackEnhanceTestCases[10]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.StackEnhance.MsgToErrSkip(0, "hello") != nil})
}

func Test_Src08_StackEnhance_FmtSkip_Empty(t *testing.T) {
	tc := stackEnhanceTestCases[11]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.StackEnhance.FmtSkip(0, "") == nil})
}

func Test_Src08_StackEnhance_FmtSkip_WithFmt(t *testing.T) {
	tc := stackEnhanceTestCases[12]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.StackEnhance.FmtSkip(0, "hello %s", "world") != nil})
}

func Test_Src08_StackEnhance_MsgErrorSkip_NilErr(t *testing.T) {
	tc := stackEnhanceTestCases[13]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.StackEnhance.MsgErrorSkip(0, "msg", nil) == ""})
}

func Test_Src08_StackEnhance_MsgErrorSkip_WithErr(t *testing.T) {
	tc := stackEnhanceTestCases[14]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StackEnhance.MsgErrorSkip(0, "msg", errors.New("e")) != ""})
}

func Test_Src08_StackEnhance_MsgErrorSkip_AlreadyHasStack(t *testing.T) {
	tc := stackEnhanceTestCases[15]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StackEnhance.MsgErrorSkip(0, "msg Stack Trace: existing", errors.New("e")) != ""})
}

func Test_Src08_StackEnhance_MsgErrorToErrSkip_NilErr(t *testing.T) {
	tc := stackEnhanceTestCases[16]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", nil) == nil})
}

func Test_Src08_StackEnhance_MsgErrorToErrSkip_WithErr(t *testing.T) {
	tc := stackEnhanceTestCases[17]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", errors.New("e")) != nil})
}

func Test_Src08_CountStateChangeTracker(t *testing.T) {
	tc := countStateChangeTrackerTestCases[0]
	c := &errcore.RawErrCollection{}
	tracker := errcore.NewCountStateChangeTracker(c)

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"isSame":           tracker.IsSameState(),
		"isValid":          tracker.IsValid(),
		"isSuccess":        tracker.IsSuccess(),
		"noChanges":        !tracker.HasChanges(),
		"notFailed":        !tracker.IsFailed(),
		"sameWithZero":     tracker.IsSameStateUsingCount(0),
		"changedAfterAdd":  func() bool { c.Add(errors.New("e")); return !tracker.IsSameState() }(),
		"hasChanges":       tracker.HasChanges(),
		"isFailed":         tracker.IsFailed(),
	})
}
