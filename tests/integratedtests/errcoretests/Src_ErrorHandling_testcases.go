package errcoretests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── CompiledError ──

var compiledErrorTestCases = []coretestcases.CaseV1{
	{Title: "CompiledError returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "CompiledError returns same error -- empty msg", ExpectedInput: args.Map{"isSame": true}},
	{Title: "CompiledError returns non-empty -- with msg", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "CompiledErrorString returns empty -- nil error", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "CompiledErrorString returns non-empty -- with error", ExpectedInput: args.Map{"nonEmpty": true}},
}

// ── JoinErrors ──

var joinErrorsTestCases = []coretestcases.CaseV1{
	{Title: "JoinErrors returns non-nil -- two errors", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "JoinErrors returns nil -- no errors", ExpectedInput: args.Map{"isNil": true}},
}

// ── ErrorWithRef ──

var errorWithRefTestCases = []coretestcases.CaseV1{
	{Title: "ErrorWithRef returns empty -- nil error", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "ErrorWithRef returns just error -- nil ref", ExpectedInput: args.Map{"result": "e"}},
	{Title: "ErrorWithRef returns just error -- empty ref", ExpectedInput: args.Map{"result": "e"}},
	{Title: "ErrorWithRef returns non-empty -- with ref", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ErrorWithRefToError returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "ErrorWithRefToError returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
}

// ── ErrorWithCompiledTraceRef ──

var errorWithCompiledTraceRefTestCases = []coretestcases.CaseV1{
	{Title: "ErrorWithCompiledTraceRef returns empty -- nil error", ExpectedInput: args.Map{"isEmpty": true}},
	{Title: "ErrorWithCompiledTraceRef returns non-empty -- empty traces", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ErrorWithCompiledTraceRef returns non-empty -- nil ref", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ErrorWithCompiledTraceRef returns non-empty -- all filled", ExpectedInput: args.Map{"nonEmpty": true}},
	{Title: "ErrorWithCompiledTraceRefToError returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "ErrorWithCompiledTraceRefToError returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
}

// ── ErrorWithTracesRefToError ──

var errorWithTracesRefToErrorTestCases = []coretestcases.CaseV1{
	{Title: "ErrorWithTracesRefToError returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "ErrorWithTracesRefToError returns non-nil -- empty traces", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "ErrorWithTracesRefToError returns non-nil -- with traces", ExpectedInput: args.Map{"nonNil": true}},
}

// ── ConcatMessageWithErr ──

var concatMessageTestCases = []coretestcases.CaseV1{
	{Title: "ConcatMessageWithErr returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "ConcatMessageWithErr returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "ConcatMessageWithErrWithStackTrace returns nil -- nil error", ExpectedInput: args.Map{"isNil": true}},
	{Title: "ConcatMessageWithErrWithStackTrace returns non-nil -- with error", ExpectedInput: args.Map{"nonNil": true}},
}

// ── ErrorToSplitLines ──

var errorToSplitLinesTestCases = []coretestcases.CaseV1{
	{Title: "ErrorToSplitLines returns empty -- nil error", ExpectedInput: args.Map{"len": 0}},
	{Title: "ErrorToSplitLines returns 2 -- multiline error", ExpectedInput: args.Map{"len": 2}},
	{Title: "ErrorToSplitNonEmptyLines returns >=2 -- multiline error", ExpectedInput: args.Map{"ge2": true}},
	{Title: "ErrorToSplitNonEmptyLines returns slice -- nil error", ExpectedInput: args.Map{"noErr": true}},
}

// ── Handlers ──

var handleErrTestCases = []coretestcases.CaseV1{
	{Title: "HandleErr does not panic -- nil error", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "HandleErr panics -- with error", ExpectedInput: args.Map{"panics": true}},
	{Title: "HandleErrMessage does not panic -- empty msg", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "HandleErrMessage panics -- with msg", ExpectedInput: args.Map{"panics": true}},
	{Title: "SimpleHandleErr does not panic -- nil error", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "SimpleHandleErr panics -- with error", ExpectedInput: args.Map{"panics": true}},
	{Title: "SimpleHandleErrMany does not panic -- nil errors", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "SimpleHandleErrMany panics -- with error", ExpectedInput: args.Map{"panics": true}},
	{Title: "MustBeEmpty does not panic -- nil error", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "MustBeEmpty panics -- with error", ExpectedInput: args.Map{"panics": true}},
}

var handleGetterTestCases = []coretestcases.CaseV1{
	{Title: "HandleCompiledErrorGetter does not panic -- nil", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "HandleCompiledErrorWithTracesGetter does not panic -- nil", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "HandleErrorGetter does not panic -- nil", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "HandleFullStringsWithTracesGetter does not panic -- nil", ExpectedInput: args.Map{"noPanic": true}},
}

var printErrorTestCases = []coretestcases.CaseV1{
	{Title: "PrintError does not panic -- nil", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "PrintError does not panic -- with error", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "PrintErrorWithTestIndex does not panic -- nil", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "PrintErrorWithTestIndex does not panic -- with error", ExpectedInput: args.Map{"noPanic": true}},
}

var panicOnIndexTestCases = []coretestcases.CaseV1{
	{Title: "PanicOnIndexOutOfRange does not panic -- valid index", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "PanicOnIndexOutOfRange panics -- out of range", ExpectedInput: args.Map{"panics": true}},
}
