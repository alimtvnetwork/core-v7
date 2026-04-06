package errcoretests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── LineDiff ──

var lineDiffTestCases = []coretestcases.CaseV1{
	{Title: "LineDiff returns match status -- equal slices", ExpectedInput: args.Map{
		"len": 2,
		"status0": "  ",
	}},
	{Title: "LineDiff returns mismatch status -- different content", ExpectedInput: args.Map{"status0": "!!"}},
	{Title: "LineDiff returns extra actual -- longer actual", ExpectedInput: args.Map{"status1": "+"}},
	{Title: "LineDiff returns missing expected -- longer expected", ExpectedInput: args.Map{"status1": "-"}},
}

var lineDiffToStringTestCases = []coretestcases.CaseV1{
	{Title: "LineDiffToString returns empty -- empty slices", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "LineDiffToString returns non-empty -- with diffs", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "LineDiffToString returns non-empty -- all match", ExpectedInput: args.Map{"nonEmpty": true}},
}

var printLineDiffTestCases = []coretestcases.CaseV1{
	{Title: "PrintLineDiff does not panic -- empty", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "PrintLineDiff does not panic -- with output", ExpectedInput: args.Map{"noPanic": true}},
}

var hasAnyMismatchTestCases = []coretestcases.CaseV1{
	{Title: "HasAnyMismatchOnLines returns false -- match", ExpectedInput: args.Map{"hasMismatch": false}},
	{Title: "HasAnyMismatchOnLines returns true -- diff len", ExpectedInput: args.Map{"hasMismatch": true}},
	{Title: "HasAnyMismatchOnLines returns true -- diff content", ExpectedInput: args.Map{"hasMismatch": true}},
}

var printLineDiffOnFailTestCases = []coretestcases.CaseV1{
	{Title: "PrintLineDiffOnFail does not panic -- no mismatch", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "PrintLineDiffOnFail does not panic -- with mismatch", ExpectedInput: args.Map{"noPanic": true}},
}

var errorToLinesLineDiffTestCases = []coretestcases.CaseV1{
	{Title: "ErrorToLinesLineDiff returns non-empty -- nil error", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ErrorToLinesLineDiff does not panic -- with error", ExpectedInput: args.Map{"noPanic": true}},
}

var sliceDiffSummaryTestCases = []coretestcases.CaseV1{
	{Title: "SliceDiffSummary returns all match -- matching slices", ExpectedInput: args.Map{"result": "all lines match"}},
	{Title: "SliceDiffSummary returns mismatch -- different slices", ExpectedInput: args.Map{"notAllMatch": true}},
}

var mapMismatchErrorTestCases = []coretestcases.CaseV1{
	{Title: "MapMismatchError returns non-empty -- with entries", ExpectedInput: args.Map{"nonEmpty": true}},
}

// ── Expecting ──

var expectingTestCases = []coretestcases.CaseV1{
	{Title: "Expecting returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ExpectingSimple returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ExpectingSimpleNoType returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ExpectingNotEqualSimpleNoType returns non-empty -- valid args", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ExpectingSimpleNoTypeError returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "ExpectingErrorSimpleNoType returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "ExpectingErrorSimpleNoTypeNewLineEnds returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "WasExpectingErrorF returns non-nil -- valid args", ExpectedInput: args.Map{"nonNil": true}},
}

var expectingFutureTestCases = []coretestcases.CaseV1{
	{Title: "ExpectingFuture returns valid record -- valid args", ExpectedInput: args.Map{
		"nonNil":            true,
		"msgNonEmpty":       true,
		"msgSimpleNonEmpty": true,
		"msgNoTypeNonEmpty": true,
		"errNonNil":         true,
		"errSimpleNonNil":   true,
		"errNoTypeNonNil":   true,
	}},
}

var expectationMessageDefTestCases = []coretestcases.CaseV1{
	{Title: "ExpectationMessageDef.ExpectedSafeString returns non-empty -- valid", ExpectedInput: args.Map{
		"nonEmpty": true,
		"cached": true,
	}},
	{Title: "ExpectationMessageDef.ExpectedSafeString returns empty -- nil expected", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "ExpectationMessageDef.ExpectedStringTrim returns trimmed -- whitespace input", ExpectedInput: args.Map{"result": "hello"}},
	{Title: "ExpectationMessageDef.ToString returns non-empty -- valid def", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ExpectationMessageDef.PrintIf does not panic -- false and true", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "ExpectationMessageDef.PrintIfFailed does not panic -- various combos", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "ExpectationMessageDef.Print does not panic -- valid def", ExpectedInput: args.Map{"noPanic": true}},
}
