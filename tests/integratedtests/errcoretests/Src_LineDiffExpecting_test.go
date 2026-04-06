package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ── LineDiff ──

func Test_Src02_LineDiff_Equal(t *testing.T) {
	// Arrange
	tc := lineDiffTestCases[0]
	// Act
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a", "b"})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"len":     len(diffs),
		"status0": diffs[0].Status,
	})
}

func Test_Src02_LineDiff_Mismatch(t *testing.T) {
	// Arrange
	tc := lineDiffTestCases[1]
	// Act
	diffs := errcore.LineDiff([]string{"a"}, []string{"b"})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"status0": diffs[0].Status})
}

func Test_Src02_LineDiff_ExtraActual(t *testing.T) {
	// Arrange
	tc := lineDiffTestCases[2]
	// Act
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a"})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"status1": diffs[1].Status})
}

func Test_Src02_LineDiff_MissingExpected(t *testing.T) {
	// Arrange
	tc := lineDiffTestCases[3]
	// Act
	diffs := errcore.LineDiff([]string{"a"}, []string{"a", "b"})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"status1": diffs[1].Status})
}

func Test_Src02_LineDiffToString_Empty(t *testing.T) {
	// Arrange
	tc := lineDiffToStringTestCases[0]
	// Act
	s := errcore.LineDiffToString(0, "test", []string{}, []string{})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": s == ""})
}

func Test_Src02_LineDiffToString_WithDiffs(t *testing.T) {
	// Arrange
	tc := lineDiffToStringTestCases[1]
	// Act
	s := errcore.LineDiffToString(0, "test", []string{"a"}, []string{"b"})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": s != ""})
}

func Test_Src02_LineDiffToString_AllMatch(t *testing.T) {
	// Arrange
	tc := lineDiffToStringTestCases[2]
	// Act
	s := errcore.LineDiffToString(0, "test", []string{"a"}, []string{"a"})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": s != ""})
}

func Test_Src02_PrintLineDiff_Empty(t *testing.T) {
	// Arrange
	tc := printLineDiffTestCases[0]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.PrintLineDiff(0, "test", []string{}, []string{}) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src02_PrintLineDiff_WithOutput(t *testing.T) {
	// Arrange
	tc := printLineDiffTestCases[1]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.PrintLineDiff(0, "test", []string{"a"}, []string{"b"}) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src02_HasAnyMismatch_Match(t *testing.T) {
	// Arrange
	tc := hasAnyMismatchTestCases[0]
	// Act
	result := errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a"})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"hasMismatch": result})
}

func Test_Src02_HasAnyMismatch_DiffLen(t *testing.T) {
	// Arrange
	tc := hasAnyMismatchTestCases[1]
	// Act
	result := errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a", "b"})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"hasMismatch": result})
}

func Test_Src02_HasAnyMismatch_DiffContent(t *testing.T) {
	// Arrange
	tc := hasAnyMismatchTestCases[2]
	// Act
	result := errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"b"})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"hasMismatch": result})
}

func Test_Src02_PrintLineDiffOnFail_NoMismatch(t *testing.T) {
	// Arrange
	tc := printLineDiffOnFailTestCases[0]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.PrintLineDiffOnFail(0, "test", []string{"a"}, []string{"a"}) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src02_PrintLineDiffOnFail_WithMismatch(t *testing.T) {
	// Arrange
	tc := printLineDiffOnFailTestCases[1]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.PrintLineDiffOnFail(0, "test", []string{"a"}, []string{"b"}) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src02_ErrorToLinesLineDiff_NilErr(t *testing.T) {
	// Arrange
	tc := errorToLinesLineDiffTestCases[0]
	// Act
	s := errcore.ErrorToLinesLineDiff(0, "test", nil, []string{"a"})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": s != ""})
}

func Test_Src02_ErrorToLinesLineDiff_WithErr(t *testing.T) {
	// Arrange
	tc := errorToLinesLineDiffTestCases[1]
	// Act
	noPanic := !callPanicsErrcore(func() {
		_ = errcore.ErrorToLinesLineDiff(0, "test", errors.New("a"), []string{"a"})
	})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src02_PrintErrorLineDiff(t *testing.T) {
	// Arrange & Act
	noPanic := !callPanicsErrcore(func() {
		errcore.PrintErrorLineDiff(0, "test", errors.New("a"), []string{"b"})
	})
	// Assert
	actual := args.Map{"result": noPanic}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected panic", actual)
}

func Test_Src02_SliceDiffSummary_Match(t *testing.T) {
	// Arrange
	tc := sliceDiffSummaryTestCases[0]
	// Act
	s := errcore.SliceDiffSummary([]string{"a"}, []string{"a"})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": s})
}

func Test_Src02_SliceDiffSummary_Mismatch(t *testing.T) {
	// Arrange
	tc := sliceDiffSummaryTestCases[1]
	// Act
	s := errcore.SliceDiffSummary([]string{"a"}, []string{"b"})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"notAllMatch": s != "all lines match"})
}

func Test_Src02_PrintDiffOnMismatch_NoMismatch(t *testing.T) {
	// Act
	noPanic := !callPanicsErrcore(func() {
		errcore.PrintDiffOnMismatch(0, "test", []string{"a"}, []string{"a"})
	})
	actual := args.Map{"result": noPanic}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected panic", actual)
}

func Test_Src02_PrintDiffOnMismatch_WithMismatch(t *testing.T) {
	// Act
	noPanic := !callPanicsErrcore(func() {
		errcore.PrintDiffOnMismatch(0, "test", []string{"a"}, []string{"b"}, "ctx1")
	})
	actual := args.Map{"result": noPanic}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected panic", actual)
}

func Test_Src02_MapMismatchError(t *testing.T) {
	// Arrange
	tc := mapMismatchErrorTestCases[0]
	// Act
	s := errcore.MapMismatchError("TestFunc", 1, "title", []string{`"a": 1,`}, []string{`"a": 2,`})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": s != ""})
}

func Test_Src02_AssertDiffOnMismatch_NoMismatch(t *testing.T) {
	// Act — no failure when matching
	errcore.AssertDiffOnMismatch(t, 0, "test", []string{"a"}, []string{"a"})
}

func Test_Src02_AssertErrorDiffOnMismatch_NilErr(t *testing.T) {
	// Act — no failure when nil error
	errcore.AssertErrorDiffOnMismatch(t, 0, "test", nil, []string{})
}

func Test_Src02_AssertErrorDiffOnMismatch_NoMismatch(t *testing.T) {
	// Act — no failure when matching
	errcore.AssertErrorDiffOnMismatch(t, 0, "test", errors.New("a"), []string{"a"})
}

// ── Expecting ──

func Test_Src03_Expecting(t *testing.T) {
	// Arrange
	tc := expectingTestCases[0]
	// Act
	s := errcore.Expecting("title", "exp", "act")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": s != ""})
}

func Test_Src03_ExpectingSimple(t *testing.T) {
	// Arrange
	tc := expectingTestCases[1]
	// Act
	s := errcore.ExpectingSimple("title", "exp", "act")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": s != ""})
}

func Test_Src03_ExpectingSimpleNoType(t *testing.T) {
	// Arrange
	tc := expectingTestCases[2]
	// Act
	s := errcore.ExpectingSimpleNoType("title", "exp", "act")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": s != ""})
}

func Test_Src03_ExpectingNotEqualSimpleNoType(t *testing.T) {
	// Arrange
	tc := expectingTestCases[3]
	// Act
	s := errcore.ExpectingNotEqualSimpleNoType("title", "exp", "act")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": s != ""})
}

func Test_Src03_ExpectingSimpleNoTypeError(t *testing.T) {
	// Arrange
	tc := expectingTestCases[4]
	// Act
	err := errcore.ExpectingSimpleNoTypeError("title", "exp", "act")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": err != nil})
}

func Test_Src03_ExpectingErrorSimpleNoType(t *testing.T) {
	// Arrange
	tc := expectingTestCases[5]
	// Act
	err := errcore.ExpectingErrorSimpleNoType("title", "exp", "act")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": err != nil})
}

func Test_Src03_ExpectingErrorSimpleNoTypeNewLineEnds(t *testing.T) {
	// Arrange
	tc := expectingTestCases[6]
	// Act
	err := errcore.ExpectingErrorSimpleNoTypeNewLineEnds("title", "exp", "act")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": err != nil})
}

func Test_Src03_WasExpectingErrorF(t *testing.T) {
	// Arrange
	tc := expectingTestCases[7]
	// Act
	err := errcore.WasExpectingErrorF("exp", "act", "title %s", "val")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": err != nil})
}

func Test_Src03_ExpectingFuture(t *testing.T) {
	// Arrange
	tc := expectingFutureTestCases[0]
	// Act
	rec := errcore.ExpectingFuture("title", "exp")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"nonNil":            rec != nil,
		"msgNonEmpty":       rec.Message("act") != "",
		"msgSimpleNonEmpty": rec.MessageSimple("act") != "",
		"msgNoTypeNonEmpty": rec.MessageSimpleNoType("act") != "",
		"errNonNil":         rec.Error("act") != nil,
		"errSimpleNonNil":   rec.ErrorSimple("act") != nil,
		"errNoTypeNonNil":   rec.ErrorSimpleNoType("act") != nil,
	})
}

func Test_Src03_ExpectationMessageDef_SafeString(t *testing.T) {
	// Arrange
	tc := expectationMessageDefTestCases[0]
	def := errcore.ExpectationMessageDef{Expected: "hello"}
	// Act
	s := def.ExpectedSafeString()
	s2 := def.ExpectedSafeString()
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"nonEmpty": s != "",
		"cached":   s == s2,
	})
}

func Test_Src03_ExpectationMessageDef_SafeString_Nil(t *testing.T) {
	// Arrange
	tc := expectationMessageDefTestCases[1]
	def := errcore.ExpectationMessageDef{}
	// Act
	s := def.ExpectedSafeString()
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": s == ""})
}

func Test_Src03_ExpectationMessageDef_StringTrim(t *testing.T) {
	// Arrange
	tc := expectationMessageDefTestCases[2]
	def := errcore.ExpectationMessageDef{Expected: " hello "}
	// Act
	s := def.ExpectedStringTrim()
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": s})
}

func Test_Src03_ExpectationMessageDef_ToString(t *testing.T) {
	// Arrange
	tc := expectationMessageDefTestCases[3]
	def := errcore.ExpectationMessageDef{Expected: "e", When: "w", FuncName: "f"}
	// Act
	s := def.ToString("act")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": s != ""})
}

func Test_Src03_ExpectationMessageDef_PrintIf(t *testing.T) {
	// Arrange
	tc := expectationMessageDefTestCases[4]
	def := errcore.ExpectationMessageDef{Expected: "e", When: "w"}
	// Act
	noPanic := !callPanicsErrcore(func() {
		def.PrintIf(false, "act")
		def.PrintIf(true, "act")
	})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src03_ExpectationMessageDef_PrintIfFailed(t *testing.T) {
	// Arrange
	tc := expectationMessageDefTestCases[5]
	def := errcore.ExpectationMessageDef{Expected: "e", When: "w"}
	// Act
	noPanic := !callPanicsErrcore(func() {
		def.PrintIfFailed(true, false, "act")
		def.PrintIfFailed(false, true, "act")
		def.PrintIfFailed(true, true, "act")
	})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src03_ExpectationMessageDef_Print(t *testing.T) {
	// Arrange
	tc := expectationMessageDefTestCases[6]
	def := errcore.ExpectationMessageDef{Expected: "e", When: "w"}
	// Act
	noPanic := !callPanicsErrcore(func() { def.Print("act") })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}
