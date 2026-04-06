package conditionaltests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// Setter
// ==========================================================================

var extSetterTestCases = []coretestcases.CaseV1{
	{
		Title:         "Setter true -- returns trueValue",
		ArrangeInput:  args.Map{"isTrue": true},
		ExpectedInput: args.Map{"isTrue": true},
	},
	{
		Title:         "Setter false -- returns falseValue",
		ArrangeInput:  args.Map{"isTrue": false},
		ExpectedInput: args.Map{"isTrue": false},
	},
}

// ==========================================================================
// SetterDefault
// ==========================================================================

var extSetterDefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "SetterDefault unset -- returns default",
		ArrangeInput:  args.Map{"useUnset": true},
		ExpectedInput: args.Map{"isTrue": true},
	},
	{
		Title:         "SetterDefault set -- returns current",
		ArrangeInput:  args.Map{"useUnset": false},
		ExpectedInput: args.Map{"isTrue": false},
	},
}

// ==========================================================================
// BoolFunctionsByOrder
// ==========================================================================

var extBoolFuncsByOrderTestCases = []coretestcases.CaseV1{
	{
		Title:         "BoolFunctionsByOrder all false -- returns false",
		ArrangeInput:  args.Map{"values": []bool{false, false, false}},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "BoolFunctionsByOrder second true -- returns true",
		ArrangeInput:  args.Map{"values": []bool{false, true, false}},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "BoolFunctionsByOrder empty -- returns false",
		ArrangeInput:  args.Map{"values": []bool{}},
		ExpectedInput: args.Map{"result": false},
	},
}

// ==========================================================================
// VoidFunctions
// ==========================================================================

var extVoidFunctionsTestCases = []coretestcases.CaseV1{
	{
		Title:         "VoidFunctions true -- executes true funcs",
		ArrangeInput:  args.Map{"isTrue": true},
		ExpectedInput: args.Map{
			"trueCount": 2,
			"falseCount": 1,
		},
	},
	{
		Title:         "VoidFunctions false -- executes false funcs only",
		ArrangeInput:  args.Map{"isTrue": false},
		ExpectedInput: args.Map{
			"trueCount": 0,
			"falseCount": 1,
		},
	},
}

// ==========================================================================
// ErrorFunc
// ==========================================================================

var extErrorFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "ErrorFunc true -- returns true func",
		ArrangeInput:  args.Map{"isTrue": true},
		ExpectedInput: args.Map{"result": "true-err"},
	},
	{
		Title:         "ErrorFunc false -- returns false func",
		ArrangeInput:  args.Map{"isTrue": false},
		ExpectedInput: args.Map{"result": "false-err"},
	},
}

// ==========================================================================
// ErrorFunctionResult
// ==========================================================================

var extErrorFunctionResultTestCases = []coretestcases.CaseV1{
	{
		Title:         "ErrorFunctionResult true -- executes true func",
		ArrangeInput:  args.Map{"isTrue": true},
		ExpectedInput: args.Map{"result": "true-err"},
	},
	{
		Title:         "ErrorFunctionResult false -- executes false func",
		ArrangeInput:  args.Map{"isTrue": false},
		ExpectedInput: args.Map{"result": "false-err"},
	},
}

// ==========================================================================
// ErrorFunctionsExecuteResults
// ==========================================================================

var extErrorFunctionsExecTestCases = []coretestcases.CaseV1{
	{
		Title:         "ErrorFunctionsExecuteResults true no errors -- nil",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"hasError": false,
		},
		ExpectedInput: args.Map{"isNil": true},
	},
	{
		Title:         "ErrorFunctionsExecuteResults true with error -- has error",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"hasError": true,
		},
		ExpectedInput: args.Map{"isNil": false},
	},
	{
		Title:         "ErrorFunctionsExecuteResults empty funcs -- nil",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"empty": true,
		},
		ExpectedInput: args.Map{"isNil": true},
	},
}

// ==========================================================================
// AnyFunctions
// ==========================================================================

var extAnyFunctionsTestCases = []coretestcases.CaseV1{
	{
		Title:         "AnyFunctions true -- returns true funcs",
		ArrangeInput:  args.Map{"isTrue": true},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "AnyFunctions false -- returns false funcs",
		ArrangeInput:  args.Map{"isTrue": false},
		ExpectedInput: args.Map{"length": 1},
	},
}

// ==========================================================================
// AnyFunctionsExecuteResults
// ==========================================================================

var extAnyFunctionsExecTestCases = []coretestcases.CaseV1{
	{
		Title:         "AnyFunctionsExecuteResults true take all -- collects results",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"scenario": "take-all",
		},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "AnyFunctionsExecuteResults with break -- stops early",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"scenario": "break-early",
		},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title:         "AnyFunctionsExecuteResults with nil func -- skips nil",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"scenario": "with-nil",
		},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title:         "AnyFunctionsExecuteResults empty -- returns nil",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"scenario": "empty",
		},
		ExpectedInput: args.Map{"length": 0},
	},
}

// ==========================================================================
// FunctionsExecuteResults (generic)
// ==========================================================================

var extFunctionsExecTestCases = []coretestcases.CaseV1{
	{
		Title:         "FunctionsExecuteResults true take all -- collects typed results",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"scenario": "take-all",
		},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "FunctionsExecuteResults with break -- stops early",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"scenario": "break-early",
		},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title:         "FunctionsExecuteResults false -- uses false funcs",
		ArrangeInput:  args.Map{
			"isTrue": false,
			"scenario": "take-all",
		},
		ExpectedInput: args.Map{"length": 1},
	},
}

// ==========================================================================
// Functions (generic selector)
// ==========================================================================

var extFunctionsSelectorTestCases = []coretestcases.CaseV1{
	{
		Title:         "Functions true -- returns true funcs slice",
		ArrangeInput:  args.Map{"isTrue": true},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "Functions false -- returns false funcs slice",
		ArrangeInput:  args.Map{"isTrue": false},
		ExpectedInput: args.Map{"length": 1},
	},
}

// ==========================================================================
// TypedErrorFunctionsExecuteResults
// ==========================================================================

var extTypedErrorFunctionsTestCases = []coretestcases.CaseV1{
	{
		Title:         "TypedErrorFunctionsExecuteResults true no errors -- collects results",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"scenario": "success",
		},
		ExpectedInput: args.Map{
			"resultLen": 2,
			"hasError": false,
		},
	},
	{
		Title:         "TypedErrorFunctionsExecuteResults true with errors -- partial results",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"scenario": "mixed",
		},
		ExpectedInput: args.Map{
			"resultLen": 1,
			"hasError": true,
		},
	},
	{
		Title:         "TypedErrorFunctionsExecuteResults empty -- nil results",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"scenario": "empty",
		},
		ExpectedInput: args.Map{
			"resultLen": 0,
			"hasError": false,
		},
	},
	{
		Title:         "TypedErrorFunctionsExecuteResults with nil func -- skips nil",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"scenario": "with-nil",
		},
		ExpectedInput: args.Map{
			"resultLen": 1,
			"hasError": false,
		},
	},
}

// ==========================================================================
// Typed convenience wrappers: NilDef, NilDefPtr, ValueOrZero, PtrOrZero, NilVal, NilValPtr
// for: String, Bool, Int, Byte
// ==========================================================================

var extNilDefStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilDefString nil -- returns default",
		ArrangeInput:  args.Map{
			"isNil": true,
			"defVal": "fallback",
		},
		ExpectedInput: "fallback",
	},
	{
		Title:         "NilDefString non-nil -- returns value",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": "actual",
			"defVal": "fallback",
		},
		ExpectedInput: "actual",
	},
}

var extNilDefPtrStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilDefPtrString nil -- returns ptr to default",
		ArrangeInput:  args.Map{
			"isNil": true,
			"defVal": "fb",
		},
		ExpectedInput: "fb",
	},
	{
		Title:         "NilDefPtrString non-nil -- returns original ptr",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": "orig",
			"defVal": "fb",
		},
		ExpectedInput: "orig",
	},
}

var extValueOrZeroStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "ValueOrZeroString nil -- returns empty",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: []string{""},
	},
	{
		Title:         "ValueOrZeroString non-nil -- returns value",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": "hi",
		},
		ExpectedInput: "hi",
	},
}

var extPtrOrZeroStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "PtrOrZeroString nil -- returns ptr to empty",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: []string{""},
	},
	{
		Title:         "PtrOrZeroString non-nil -- returns value",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": "world",
		},
		ExpectedInput: "world",
	},
}

var extNilValStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilValString nil -- returns onNil",
		ArrangeInput:  args.Map{
			"isNil": true,
			"onNil": "n",
			"onNonNil": "s",
		},
		ExpectedInput: "n",
	},
	{
		Title:         "NilValString non-nil -- returns onNonNil",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": "x",
			"onNil": "n",
			"onNonNil": "s",
		},
		ExpectedInput: "s",
	},
}

var extNilValPtrStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilValPtrString nil -- returns ptr to onNil",
		ArrangeInput:  args.Map{
			"isNil": true,
			"onNil": "n",
			"onNonNil": "s",
		},
		ExpectedInput: "n",
	},
	{
		Title:         "NilValPtrString non-nil -- returns ptr to onNonNil",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": "x",
			"onNil": "n",
			"onNonNil": "s",
		},
		ExpectedInput: "s",
	},
}

var extNilDefBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilDefBool nil -- returns default",
		ArrangeInput:  args.Map{
			"isNil": true,
			"defVal": true,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "NilDefBool non-nil -- returns value",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": false,
			"defVal": true,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var extNilDefIntTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilDefInt nil -- returns default",
		ArrangeInput:  args.Map{
			"isNil": true,
			"defVal": 42,
		},
		ExpectedInput: args.Map{"result": 42},
	},
	{
		Title:         "NilDefInt non-nil -- returns value",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": 7,
			"defVal": 42,
		},
		ExpectedInput: args.Map{"result": 7},
	},
}

var extNilDefByteTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilDefByte nil -- returns default",
		ArrangeInput:  args.Map{
			"isNil": true,
			"defVal": byte(0xFF),
		},
		ExpectedInput: args.Map{"result": byte(0xFF)},
	},
	{
		Title:         "NilDefByte non-nil -- returns value",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": byte(0x0A),
			"defVal": byte(0xFF),
		},
		ExpectedInput: args.Map{"result": byte(0x0A)},
	},
}

var extValueOrZeroBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "ValueOrZeroBool nil -- returns false",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "ValueOrZeroBool non-nil true -- returns true",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": true,
		},
		ExpectedInput: args.Map{"result": true},
	},
}

var extValueOrZeroIntTestCases = []coretestcases.CaseV1{
	{
		Title:         "ValueOrZeroInt nil -- returns 0",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"result": 0},
	},
	{
		Title:         "ValueOrZeroInt non-nil -- returns value",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": 99,
		},
		ExpectedInput: args.Map{"result": 99},
	},
}

var extValueOrZeroByteTestCases = []coretestcases.CaseV1{
	{
		Title:         "ValueOrZeroByte nil -- returns 0",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"result": byte(0)},
	},
	{
		Title:         "ValueOrZeroByte non-nil -- returns value",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": byte(42),
		},
		ExpectedInput: args.Map{"result": byte(42)},
	},
}

var extNilValBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilValBool nil -- returns onNil",
		ArrangeInput:  args.Map{
			"isNil": true,
			"onNil": true,
			"onNonNil": false,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "NilValBool non-nil -- returns onNonNil",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": true,
			"onNil": true,
			"onNonNil": false,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var extNilValIntTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilValInt nil -- returns onNil",
		ArrangeInput:  args.Map{
			"isNil": true,
			"onNil": -1,
			"onNonNil": 1,
		},
		ExpectedInput: args.Map{"result": -1},
	},
	{
		Title:         "NilValInt non-nil -- returns onNonNil",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": 50,
			"onNil": -1,
			"onNonNil": 1,
		},
		ExpectedInput: args.Map{"result": 1},
	},
}

var extNilValByteTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilValByte nil -- returns onNil",
		ArrangeInput:  args.Map{
			"isNil": true,
			"onNil": byte(0),
			"onNonNil": byte(1),
		},
		ExpectedInput: args.Map{"result": byte(0)},
	},
	{
		Title:         "NilValByte non-nil -- returns onNonNil",
		ArrangeInput:  args.Map{
			"isNil": false,
			"value": byte(5),
			"onNil": byte(0),
			"onNonNil": byte(1),
		},
		ExpectedInput: args.Map{"result": byte(1)},
	},
}

// IfString typed wrapper
var extIfStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "IfString true -- returns trueValue",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"trueValue": "yes",
			"falseValue": "no",
		},
		ExpectedInput: "yes",
	},
	{
		Title:         "IfString false -- returns falseValue",
		ArrangeInput:  args.Map{
			"isTrue": false,
			"trueValue": "yes",
			"falseValue": "no",
		},
		ExpectedInput: "no",
	},
}

// IfTrueFuncInt
var extIfTrueFuncIntTestCases = []coretestcases.CaseV1{
	{
		Title:         "IfTrueFuncInt true -- returns trueFunc value",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"trueValue": 42,
		},
		ExpectedInput: args.Map{"result": 42},
	},
	{
		Title:         "IfTrueFuncInt false -- returns zero",
		ArrangeInput:  args.Map{
			"isTrue": false,
			"trueValue": 42,
		},
		ExpectedInput: args.Map{"result": 0},
	},
}

// IfFuncByte
var extIfFuncByteTestCases = []coretestcases.CaseV1{
	{
		Title:         "IfFuncByte true -- returns trueFunc value",
		ArrangeInput:  args.Map{
			"isTrue": true,
			"trueValue": byte(1),
			"falseValue": byte(0),
		},
		ExpectedInput: args.Map{"result": byte(1)},
	},
	{
		Title:         "IfFuncByte false -- returns falseFunc value",
		ArrangeInput:  args.Map{
			"isTrue": false,
			"trueValue": byte(1),
			"falseValue": byte(0),
		},
		ExpectedInput: args.Map{"result": byte(0)},
	},
}
