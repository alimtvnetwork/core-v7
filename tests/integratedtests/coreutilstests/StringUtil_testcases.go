package coreutilstests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var isEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmpty returns true for empty string",
		ArrangeInput: args.Map{
			"when":  "given empty string",
			"input": "",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEmpty returns false for non-empty string",
		ArrangeInput: args.Map{
			"when":  "given non-empty string",
			"input": "hello",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsEmpty returns false for whitespace-only string",
		ArrangeInput: args.Map{
			"when":  "given whitespace string",
			"input": "   ",
		},
		ExpectedInput: "false",
	},
}

var isBlankTestCases = []coretestcases.CaseV1{
	{
		Title: "IsBlank returns true for empty string",
		ArrangeInput: args.Map{
			"when":  "given empty string",
			"input": "",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsBlank returns true for whitespace only",
		ArrangeInput: args.Map{
			"when":  "given whitespace string",
			"input": "   ",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsBlank returns false for non-blank string",
		ArrangeInput: args.Map{
			"when":  "given non-blank string",
			"input": "hello",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsBlank returns true for tab and newline only",
		ArrangeInput: args.Map{
			"when":  "given tab and newline",
			"input": "\t\n",
		},
		ExpectedInput: "true",
	},
}

var isEmptyOrWhitespaceTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmptyOrWhitespace true for empty string",
		ArrangeInput: args.Map{
			"when":  "given empty",
			"input": "",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEmptyOrWhitespace true for single space",
		ArrangeInput: args.Map{
			"when":  "given single space",
			"input": " ",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEmptyOrWhitespace true for newline",
		ArrangeInput: args.Map{
			"when":  "given newline",
			"input": "\n",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEmptyOrWhitespace true for mixed whitespace",
		ArrangeInput: args.Map{
			"when":  "given tabs and spaces",
			"input": "  \t  ",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEmptyOrWhitespace false for content",
		ArrangeInput: args.Map{
			"when":  "given actual content",
			"input": " hello ",
		},
		ExpectedInput: "false",
	},
}

var safeSubstringTestCases = []coretestcases.CaseV1{
	{
		Title: "SafeSubstring returns full content for -1,-1",
		ArrangeInput: args.Map{
			"when":    "given -1,-1 indices",
			"content": "hello world",
			"start":   -1,
			"end":     -1,
		},
		ExpectedInput: "hello world",
	},
	{
		Title: "SafeSubstring returns substring for valid range",
		ArrangeInput: args.Map{
			"when":    "given 0,5 indices",
			"content": "hello world",
			"start":   0,
			"end":     5,
		},
		ExpectedInput: "hello",
	},
	{
		Title: "SafeSubstring returns empty for empty input",
		ArrangeInput: args.Map{
			"when":    "given empty content",
			"content": "",
			"start":   0,
			"end":     5,
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "SafeSubstring returns from start with -1 end",
		ArrangeInput: args.Map{
			"when":    "given start=2, end=-1",
			"content": "hello",
			"start":   2,
			"end":     -1,
		},
		ExpectedInput: "llo",
	},
	{
		Title: "SafeSubstring returns to end with -1 start",
		ArrangeInput: args.Map{
			"when":    "given start=-1, end=3",
			"content": "hello",
			"start":   -1,
			"end":     3,
		},
		ExpectedInput: "hel",
	},
	{
		Title: "SafeSubstring returns empty for out-of-range",
		ArrangeInput: args.Map{
			"when":    "given out-of-range indices",
			"content": "hi",
			"start":   5,
			"end":     10,
		},
		ExpectedInput: []string{""},
	},
}

var splitLeftRightTestCases = []coretestcases.CaseV1{
	{
		Title: "SplitLeftRight splits on separator",
		ArrangeInput: args.Map{
			"when":      "given key=value format",
			"input":     "name=John",
			"separator": "=",
		},
		ExpectedInput: args.Map{
			"left": "name",
			"right": "John",
		},
	},
	{
		Title: "SplitLeftRight returns empty right when no separator",
		ArrangeInput: args.Map{
			"when":      "given no separator present",
			"input":     "noseparator",
			"separator": "=",
		},
		ExpectedInput: args.Map{
			"left": "noseparator",
			"right": "",
		},
	},
	{
		Title: "SplitLeftRight splits only first occurrence",
		ArrangeInput: args.Map{
			"when":      "given multiple separators",
			"input":     "a=b=c",
			"separator": "=",
		},
		ExpectedInput: args.Map{
			"left": "a",
			"right": "b=c",
		},
	},
}

var isStartsWithTestCases = []coretestcases.CaseV1{
	{
		Title: "IsStartsWith matches prefix case-sensitive",
		ArrangeInput: args.Map{
			"when":         "given matching prefix",
			"content":      "Hello World",
			"startsWith":   "Hello",
			"isIgnoreCase": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsStartsWith fails on case mismatch when case-sensitive",
		ArrangeInput: args.Map{
			"when":         "given wrong case prefix",
			"content":      "Hello World",
			"startsWith":   "hello",
			"isIgnoreCase": false,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsStartsWith matches case-insensitive",
		ArrangeInput: args.Map{
			"when":         "given wrong case but ignore case",
			"content":      "Hello World",
			"startsWith":   "hello",
			"isIgnoreCase": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsStartsWith returns true for empty prefix",
		ArrangeInput: args.Map{
			"when":         "given empty prefix",
			"content":      "anything",
			"startsWith":   "",
			"isIgnoreCase": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsStartsWith returns false when prefix longer than content",
		ArrangeInput: args.Map{
			"when":         "given prefix longer than content",
			"content":      "Hi",
			"startsWith":   "Hello World",
			"isIgnoreCase": false,
		},
		ExpectedInput: "false",
	},
}

var isEndsWithTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEndsWith matches suffix case-sensitive",
		ArrangeInput: args.Map{
			"when":         "given matching suffix",
			"content":      "Hello World",
			"endsWith":     "World",
			"isIgnoreCase": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEndsWith fails on case mismatch when case-sensitive",
		ArrangeInput: args.Map{
			"when":         "given wrong case suffix",
			"content":      "Hello World",
			"endsWith":     "world",
			"isIgnoreCase": false,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsEndsWith matches case-insensitive",
		ArrangeInput: args.Map{
			"when":         "given wrong case but ignore case",
			"content":      "Hello World",
			"endsWith":     "world",
			"isIgnoreCase": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEndsWith returns true for empty suffix",
		ArrangeInput: args.Map{
			"when":         "given empty suffix",
			"content":      "anything",
			"endsWith":     "",
			"isIgnoreCase": false,
		},
		ExpectedInput: "true",
	},
}

var removeManyTestCases = []coretestcases.CaseV1{
	{
		Title: "RemoveMany removes single occurrence",
		ArrangeInput: args.Map{
			"when":    "given single removal",
			"content": "hello world",
			"removes": []string{"world"},
		},
		ExpectedInput: "hello ",
	},
	{
		Title: "RemoveMany removes multiple patterns",
		ArrangeInput: args.Map{
			"when":    "given multiple removals",
			"content": "hello beautiful world",
			"removes": []string{"beautiful ", "world"},
		},
		ExpectedInput: "hello ",
	},
	{
		Title: "RemoveMany returns same for empty content",
		ArrangeInput: args.Map{
			"when":    "given empty content",
			"content": "",
			"removes": []string{"anything"},
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "RemoveMany no match leaves content unchanged",
		ArrangeInput: args.Map{
			"when":    "given no matching patterns",
			"content": "hello",
			"removes": []string{"xyz"},
		},
		ExpectedInput: "hello",
	},
}

var replaceWhiteSpacesToSingleTestCases = []coretestcases.CaseV1{
	{
		Title: "ReplaceWhiteSpacesToSingle collapses multiple spaces",
		ArrangeInput: args.Map{
			"when":  "given multiple spaces",
			"input": "  some  nothing    ",
		},
		ExpectedInput: "some nothing",
	},
	{
		Title: "ReplaceWhiteSpacesToSingle removes tabs",
		ArrangeInput: args.Map{
			"when":  "given tabs in content",
			"input": "hello\tworld",
		},
		ExpectedInput: "helloworld",
	},
	{
		Title: "ReplaceWhiteSpacesToSingle returns empty for whitespace only",
		ArrangeInput: args.Map{
			"when":  "given whitespace only",
			"input": "   \t\n  ",
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "ReplaceWhiteSpacesToSingle preserves single-spaced text",
		ArrangeInput: args.Map{
			"when":  "given already single-spaced",
			"input": "hello world",
		},
		ExpectedInput: "hello world",
	},
}
