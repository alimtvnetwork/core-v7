package conditionaltests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var stringDefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "StringDefault true -- returns value",
		ArrangeInput:  args.Map{
			"when": "true",
			"isTrue": true,
			"trueValue": "hello",
		},
		ExpectedInput: "hello",
	},
	{
		Title:         "StringDefault false -- returns empty",
		ArrangeInput:  args.Map{
			"when": "false",
			"isTrue": false,
			"trueValue": "hello",
		},
		ExpectedInput: "",
	},
}

var nilOrEmptyStrTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilOrEmptyStr nil -- returns onNilOrEmpty",
		ArrangeInput:  args.Map{
			"when": "nil ptr",
			"isNil": true,
			"onNilOrEmpty": "default",
			"onNonNil": "present",
		},
		ExpectedInput: "default",
	},
	{
		Title:         "NilOrEmptyStr empty string -- returns onNilOrEmpty",
		ArrangeInput:  args.Map{
			"when": "empty str",
			"isNil": false,
			"value": "",
			"onNilOrEmpty": "default",
			"onNonNil": "present",
		},
		ExpectedInput: "default",
	},
	{
		Title:         "NilOrEmptyStr non-empty -- returns onNonNil",
		ArrangeInput:  args.Map{
			"when": "has value",
			"isNil": false,
			"value": "hello",
			"onNilOrEmpty": "default",
			"onNonNil": "present",
		},
		ExpectedInput: "present",
	},
}

var nilOrEmptyStrPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilOrEmptyStrPtr nil -- returns ptr to onNilOrEmpty",
		ArrangeInput:  args.Map{
			"when": "nil ptr",
			"isNil": true,
			"onNilOrEmpty": "default",
			"onNonNil": "present",
		},
		ExpectedInput: "default",
	},
	{
		Title:         "NilOrEmptyStrPtr non-empty -- returns ptr to onNonNil",
		ArrangeInput:  args.Map{
			"when": "has value",
			"isNil": false,
			"value": "hello",
			"onNilOrEmpty": "default",
			"onNonNil": "present",
		},
		ExpectedInput: "present",
	},
}

var nilDefPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilDefPtr nil -- returns ptr to defVal",
		ArrangeInput:  args.Map{
			"when": "nil ptr",
			"isNil": true,
			"defVal": "fallback",
		},
		ExpectedInput: "fallback",
	},
	{
		Title:         "NilDefPtr non-nil -- returns original ptr value",
		ArrangeInput:  args.Map{
			"when": "has value",
			"isNil": false,
			"value": "original",
			"defVal": "fallback",
		},
		ExpectedInput: "original",
	},
}

var nilValTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilVal nil -- returns onNil",
		ArrangeInput:  args.Map{
			"when": "nil ptr",
			"isNil": true,
			"onNil": "nilval",
			"onNonNil": "present",
		},
		ExpectedInput: "nilval",
	},
	{
		Title:         "NilVal non-nil -- returns onNonNil",
		ArrangeInput:  args.Map{
			"when": "non-nil ptr",
			"isNil": false,
			"onNil": "nilval",
			"onNonNil": "present",
		},
		ExpectedInput: "present",
	},
}

var nilValPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "NilValPtr nil -- returns ptr to onNil",
		ArrangeInput:  args.Map{
			"when": "nil ptr",
			"isNil": true,
			"onNil": "nilval",
			"onNonNil": "present",
		},
		ExpectedInput: "nilval",
	},
	{
		Title:         "NilValPtr non-nil -- returns ptr to onNonNil",
		ArrangeInput:  args.Map{
			"when": "non-nil",
			"isNil": false,
			"onNil": "nilval",
			"onNonNil": "present",
		},
		ExpectedInput: "present",
	},
}

var valueOrZeroTestCases = []coretestcases.CaseV1{
	{
		Title:         "ValueOrZero nil -- returns zero value",
		ArrangeInput:  args.Map{
			"when": "nil string ptr",
			"isNil": true,
		},
		ExpectedInput: "",
	},
	{
		Title:         "ValueOrZero non-nil -- returns value",
		ArrangeInput:  args.Map{
			"when": "has value",
			"isNil": false,
			"value": "hello",
		},
		ExpectedInput: "hello",
	},
}

var ptrOrZeroTestCases = []coretestcases.CaseV1{
	{
		Title:         "PtrOrZero nil -- returns ptr to zero",
		ArrangeInput:  args.Map{
			"when": "nil int ptr",
			"isNil": true,
		},
		ExpectedInput: args.Map{
			"isNil": false,
			"value": 0,
		},
	},
	{
		Title:         "PtrOrZero non-nil -- returns original ptr",
		ArrangeInput:  args.Map{
			"when": "has value",
			"isNil": false,
			"value": 42,
		},
		ExpectedInput: args.Map{
			"isNil": false,
			"value": 42,
		},
	},
}

var ifPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "IfPtr true -- returns true ptr",
		ArrangeInput:  args.Map{
			"when": "true",
			"isTrue": true,
			"trueValue": "yes",
			"falseValue": "no",
		},
		ExpectedInput: "yes",
	},
	{
		Title:         "IfPtr false -- returns false ptr",
		ArrangeInput:  args.Map{
			"when": "false",
			"isTrue": false,
			"trueValue": "yes",
			"falseValue": "no",
		},
		ExpectedInput: "no",
	},
}

var funcTestCases = []coretestcases.CaseV1{
	{
		Title:         "Func true -- returns true func",
		ArrangeInput:  args.Map{
			"when": "true condition",
			"isTrue": true,
		},
		ExpectedInput: "true-result",
	},
	{
		Title:         "Func false -- returns false func",
		ArrangeInput:  args.Map{
			"when": "false condition",
			"isTrue": false,
		},
		ExpectedInput: "false-result",
	},
}
