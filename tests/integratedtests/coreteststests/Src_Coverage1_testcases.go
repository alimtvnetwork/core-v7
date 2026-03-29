package coreteststests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var srcAnyToBytesTestCases = []coretestcases.CaseV1{
	{
		Title: "AnyToBytes returns bytes pass-through -- bytes input",
		ArrangeInput: args.Map{
			"input": []byte("hello"),
			"type":  "bytes",
		},
		ExpectedInput: args.Map{
			"result": "hello",
		},
	},
	{
		Title: "AnyToBytes returns nil -- nil bytes input",
		ArrangeInput: args.Map{
			"input": nil,
			"type":  "nilBytes",
		},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
	{
		Title: "AnyToBytes returns converted bytes -- string input",
		ArrangeInput: args.Map{
			"input": "world",
			"type":  "string",
		},
		ExpectedInput: args.Map{
			"result": "world",
		},
	},
	{
		Title: "AnyToBytes returns json bytes -- map input",
		ArrangeInput: args.Map{
			"input": map[string]int{"a": 1},
			"type":  "other",
		},
		ExpectedInput: args.Map{
			"nonEmpty": true,
		},
	},
}

var srcDraftTypePtrOrNonPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "PtrOrNonPtr returns ptr -- asPtr true",
		ArrangeInput: args.Map{
			"string1": "a",
			"integer": 1,
			"asPtr":   true,
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
	{
		Title: "PtrOrNonPtr returns non-ptr -- asPtr false",
		ArrangeInput: args.Map{
			"string1": "a",
			"integer": 1,
			"asPtr":   false,
		},
		ExpectedInput: args.Map{
			"isDraftType": true,
		},
	},
}

var srcDraftTypeClonePtrNilTestCase = coretestcases.CaseV1{
	Title: "ClonePtr returns nil -- nil receiver",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var srcDraftTypeIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEqual returns true -- identical drafts",
		ArrangeInput: args.Map{
			"scenario": "equal",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsEqual returns false -- different SampleString2",
		ArrangeInput: args.Map{
			"scenario": "diffString2",
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "IsEqual returns false -- different SampleInteger",
		ArrangeInput: args.Map{
			"scenario": "diffInteger",
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "IsEqual returns false -- different RawBytes",
		ArrangeInput: args.Map{
			"scenario": "diffRawBytes",
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "IsEqual returns false -- different Lines",
		ArrangeInput: args.Map{
			"scenario": "diffLines",
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "IsEqual returns true -- both nil",
		ArrangeInput: args.Map{
			"scenario": "bothNil",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsEqual returns false -- nil vs non-nil",
		ArrangeInput: args.Map{
			"scenario": "nilVsNonNil",
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "IsEqual returns true -- same pointer",
		ArrangeInput: args.Map{
			"scenario": "samePtr",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
}

var srcSimpleTestCaseTitlesTestCase = coretestcases.CaseV1{
	Title: "SimpleTestCase CaseTitle returns correct title -- title set",
	ArrangeInput: args.Map{
		"title": "test-title",
	},
	ExpectedInput: args.Map{
		"caseTitle":       "test-title",
		"formTitleNotEmpty": true,
		"customTitleNotEmpty": true,
	},
}
