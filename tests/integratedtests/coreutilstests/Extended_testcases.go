package coreutilstests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// IsNotEmpty
// ==========================================

var extIsNotEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsNotEmpty returns false for empty string",
		ArrangeInput: args.Map{
			"input": "",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsNotEmpty returns true for non-empty string",
		ArrangeInput: args.Map{
			"input": "hello",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsNotEmpty returns true for whitespace",
		ArrangeInput: args.Map{
			"input": "  ",
		},
		ExpectedInput: "true",
	},
}

// ==========================================
// IsDefined
// ==========================================

var extIsDefinedTestCases = []coretestcases.CaseV1{
	{
		Title: "IsDefined returns false for empty",
		ArrangeInput: args.Map{
			"input": "",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsDefined returns false for single space",
		ArrangeInput: args.Map{
			"input": " ",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsDefined returns false for newline",
		ArrangeInput: args.Map{
			"input": "\n",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsDefined returns false for tabs and spaces",
		ArrangeInput: args.Map{
			"input": " \t ",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsDefined returns true for actual content",
		ArrangeInput: args.Map{
			"input": "hello",
		},
		ExpectedInput: "true",
	},
}

// ==========================================
// IsStarts (case sensitive wrapper)
// ==========================================

var extIsStartsTestCases = []coretestcases.CaseV1{
	{
		Title: "IsStarts matches prefix case-sensitive",
		ArrangeInput: args.Map{
			"content":    "HelloWorld",
			"startsWith": "Hello",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsStarts fails on case mismatch",
		ArrangeInput: args.Map{
			"content":    "HelloWorld",
			"startsWith": "hello",
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// IsEnds (case sensitive wrapper)
// ==========================================

var extIsEndsTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEnds matches suffix case-sensitive",
		ArrangeInput: args.Map{
			"content":  "HelloWorld",
			"endsWith": "World",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEnds fails on case mismatch",
		ArrangeInput: args.Map{
			"content":  "HelloWorld",
			"endsWith": "world",
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// IsStartsChar
// ==========================================

var extIsStartsCharTestCases = []coretestcases.CaseV1{
	{
		Title: "IsStartsChar true for matching first char",
		ArrangeInput: args.Map{
			"content": "Hello",
			"char":    byte('H'),
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsStartsChar false for non-matching",
		ArrangeInput: args.Map{
			"content": "Hello",
			"char":    byte('h'),
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsStartsChar false for empty",
		ArrangeInput: args.Map{
			"content": "",
			"char":    byte('H'),
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// IsEndsChar
// ==========================================

var extIsEndsCharTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEndsChar true for matching last char",
		ArrangeInput: args.Map{
			"content": "Hello",
			"char":    byte('o'),
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEndsChar false for non-matching",
		ArrangeInput: args.Map{
			"content": "Hello",
			"char":    byte('O'),
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsEndsChar false for empty content",
		ArrangeInput: args.Map{
			"content": "",
			"char":    byte('x'),
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// IsStartsRune
// ==========================================

var extIsStartsRuneTestCases = []coretestcases.CaseV1{
	{
		Title: "IsStartsRune true for matching rune",
		ArrangeInput: args.Map{
			"content": "Hello",
			"rune":    'H',
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsStartsRune false for non-matching",
		ArrangeInput: args.Map{
			"content": "Hello",
			"rune":    'h',
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsStartsRune false for empty",
		ArrangeInput: args.Map{
			"content": "",
			"rune":    'H',
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// IsEndsRune
// ==========================================

var extIsEndsRuneTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEndsRune true for matching last rune",
		ArrangeInput: args.Map{
			"content": "Hello",
			"rune":    'o',
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEndsRune false for non-matching",
		ArrangeInput: args.Map{
			"content": "Hello",
			"rune":    'O',
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// IsStartsAndEndsChar
// ==========================================

var extIsStartsAndEndsCharTestCases = []coretestcases.CaseV1{
	{
		Title: "IsStartsAndEndsChar true for matching both",
		ArrangeInput: args.Map{
			"content":   "{hello}",
			"startChar": byte('{'),
			"endChar":   byte('}'),
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsStartsAndEndsChar false when start doesnt match",
		ArrangeInput: args.Map{
			"content":   "hello}",
			"startChar": byte('{'),
			"endChar":   byte('}'),
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsStartsAndEndsChar false for empty",
		ArrangeInput: args.Map{
			"content":   "",
			"startChar": byte('{'),
			"endChar":   byte('}'),
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// IsStartsAndEndsWith
// ==========================================

var extIsStartsAndEndsWithTestCases = []coretestcases.CaseV1{
	{
		Title: "IsStartsAndEndsWith true for matching both case-sensitive",
		ArrangeInput: args.Map{
			"content":      "Hello World",
			"startsWith":   "Hello",
			"endsWith":     "World",
			"isIgnoreCase": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsStartsAndEndsWith false when end doesnt match case-sensitive",
		ArrangeInput: args.Map{
			"content":      "Hello World",
			"startsWith":   "Hello",
			"endsWith":     "world",
			"isIgnoreCase": false,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsStartsAndEndsWith true case-insensitive",
		ArrangeInput: args.Map{
			"content":      "Hello World",
			"startsWith":   "hello",
			"endsWith":     "world",
			"isIgnoreCase": true,
		},
		ExpectedInput: "true",
	},
}

// ==========================================
// IsStartsAndEnds (case-sensitive wrapper)
// ==========================================

var extIsStartsAndEndsTestCases = []coretestcases.CaseV1{
	{
		Title: "IsStartsAndEnds true for matching both",
		ArrangeInput: args.Map{
			"content":    "Hello World",
			"startsWith": "Hello",
			"endsWith":   "World",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsStartsAndEnds false for case mismatch",
		ArrangeInput: args.Map{
			"content":    "Hello World",
			"startsWith": "hello",
			"endsWith":   "World",
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// IsAnyStartsWith
// ==========================================

var extIsAnyStartsWithTestCases = []coretestcases.CaseV1{
	{
		Title: "IsAnyStartsWith true when one matches",
		ArrangeInput: args.Map{
			"content":      "Hello",
			"isIgnoreCase": false,
			"terms":        []string{"Hi", "Hello"},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsAnyStartsWith false when none match",
		ArrangeInput: args.Map{
			"content":      "Hello",
			"isIgnoreCase": false,
			"terms":        []string{"Bye", "World"},
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsAnyStartsWith true case-insensitive",
		ArrangeInput: args.Map{
			"content":      "Hello",
			"isIgnoreCase": true,
			"terms":        []string{"hello"},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsAnyStartsWith true both empty",
		ArrangeInput: args.Map{
			"content":      "",
			"isIgnoreCase": false,
			"terms":        []string{},
		},
		ExpectedInput: "true",
	},
}

// ==========================================
// IsAnyEndsWith
// ==========================================

var extIsAnyEndsWithTestCases = []coretestcases.CaseV1{
	{
		Title: "IsAnyEndsWith true when one matches",
		ArrangeInput: args.Map{
			"content":      "Hello",
			"isIgnoreCase": false,
			"terms":        []string{"xyz", "llo"},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsAnyEndsWith false when none match",
		ArrangeInput: args.Map{
			"content":      "Hello",
			"isIgnoreCase": false,
			"terms":        []string{"abc", "def"},
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// FirstChar / FirstCharOrDefault / LastChar / LastCharOrDefault
// ==========================================

var extFirstCharTestCases = []coretestcases.CaseV1{
	{
		Title: "FirstCharOrDefault returns first byte",
		ArrangeInput: args.Map{
			"input": "Hello",
		},
		ExpectedInput: "72", // 'H' == 72
	},
	{
		Title: "FirstCharOrDefault returns 0 for empty",
		ArrangeInput: args.Map{
			"input": "",
		},
		ExpectedInput: "0",
	},
}

var extLastCharTestCases = []coretestcases.CaseV1{
	{
		Title: "LastCharOrDefault returns last byte",
		ArrangeInput: args.Map{
			"input": "Hello",
		},
		ExpectedInput: "111", // 'o' == 111
	},
	{
		Title: "LastCharOrDefault returns 0 for empty",
		ArrangeInput: args.Map{
			"input": "",
		},
		ExpectedInput: "0",
	},
}

// ==========================================
// ClonePtr / SafeClonePtr
// ==========================================

var extClonePtrTestCases = []coretestcases.CaseV1{
	{
		Title: "ClonePtr returns nil for nil input",
		ArrangeInput: args.Map{
			"isNil": true,
		},
		ExpectedInput: "true", // result is nil
	},
	{
		Title: "ClonePtr clones value",
		ArrangeInput: args.Map{
			"isNil": false,
			"value": "hello",
		},
		ExpectedInput: "hello",
	},
}

var extSafeClonePtrTestCases = []coretestcases.CaseV1{
	{
		Title: "SafeClonePtr returns empty string ptr for nil",
		ArrangeInput: args.Map{
			"isNil": true,
		},
		ExpectedInput: "",
	},
	{
		Title: "SafeClonePtr clones value",
		ArrangeInput: args.Map{
			"isNil": false,
			"value": "world",
		},
		ExpectedInput: "world",
	},
}

// ==========================================
// Ptr functions: IsEmptyPtr, IsBlankPtr, IsEmptyOrWhitespacePtr, IsNullOrEmptyPtr, IsDefinedPtr
// ==========================================

var extPtrFunctionsTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmptyPtr true for nil",
		ArrangeInput: args.Map{
			"func":  "IsEmptyPtr",
			"isNil": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEmptyPtr true for empty string ptr",
		ArrangeInput: args.Map{
			"func":  "IsEmptyPtr",
			"isNil": false,
			"value": "",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEmptyPtr false for non-empty",
		ArrangeInput: args.Map{
			"func":  "IsEmptyPtr",
			"isNil": false,
			"value": "abc",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsBlankPtr true for nil",
		ArrangeInput: args.Map{
			"func":  "IsBlankPtr",
			"isNil": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsBlankPtr true for whitespace ptr",
		ArrangeInput: args.Map{
			"func":  "IsBlankPtr",
			"isNil": false,
			"value": "  \t ",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsBlankPtr false for non-blank",
		ArrangeInput: args.Map{
			"func":  "IsBlankPtr",
			"isNil": false,
			"value": "x",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsEmptyOrWhitespacePtr true for nil",
		ArrangeInput: args.Map{
			"func":  "IsEmptyOrWhitespacePtr",
			"isNil": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEmptyOrWhitespacePtr true for whitespace",
		ArrangeInput: args.Map{
			"func":  "IsEmptyOrWhitespacePtr",
			"isNil": false,
			"value": "  ",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsNullOrEmptyPtr true for nil",
		ArrangeInput: args.Map{
			"func":  "IsNullOrEmptyPtr",
			"isNil": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsNullOrEmptyPtr true for empty",
		ArrangeInput: args.Map{
			"func":  "IsNullOrEmptyPtr",
			"isNil": false,
			"value": "",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsNullOrEmptyPtr false for non-empty",
		ArrangeInput: args.Map{
			"func":  "IsNullOrEmptyPtr",
			"isNil": false,
			"value": "abc",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsDefinedPtr false for nil",
		ArrangeInput: args.Map{
			"func":  "IsDefinedPtr",
			"isNil": true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsDefinedPtr true for non-whitespace content",
		ArrangeInput: args.Map{
			"func":  "IsDefinedPtr",
			"isNil": false,
			"value": "abc",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsDefinedPtr false for whitespace",
		ArrangeInput: args.Map{
			"func":  "IsDefinedPtr",
			"isNil": false,
			"value": "  ",
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// ToBool
// ==========================================

var extToBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "ToBool empty returns false",
		ArrangeInput:  args.Map{"input": ""},
		ExpectedInput: "false",
	},
	{
		Title:         "ToBool yes returns true",
		ArrangeInput:  args.Map{"input": "yes"},
		ExpectedInput: "true",
	},
	{
		Title:         "ToBool YES returns true",
		ArrangeInput:  args.Map{"input": "YES"},
		ExpectedInput: "true",
	},
	{
		Title:         "ToBool 1 returns true",
		ArrangeInput:  args.Map{"input": "1"},
		ExpectedInput: "true",
	},
	{
		Title:         "ToBool no returns false",
		ArrangeInput:  args.Map{"input": "no"},
		ExpectedInput: "false",
	},
	{
		Title:         "ToBool 0 returns false",
		ArrangeInput:  args.Map{"input": "0"},
		ExpectedInput: "false",
	},
	{
		Title:         "ToBool true returns true",
		ArrangeInput:  args.Map{"input": "true"},
		ExpectedInput: "true",
	},
	{
		Title:         "ToBool false returns false",
		ArrangeInput:  args.Map{"input": "false"},
		ExpectedInput: "false",
	},
	{
		Title:         "ToBool invalid returns false",
		ArrangeInput:  args.Map{"input": "invalid"},
		ExpectedInput: "false",
	},
}

// ==========================================
// ToByte / ToByteDefault
// ==========================================

var extToByteTestCases = []coretestcases.CaseV1{
	{
		Title:         "ToByte valid returns byte",
		ArrangeInput:  args.Map{
			"input": "200",
			"def": byte(0),
		},
		ExpectedInput: "200",
	},
	{
		Title:         "ToByte invalid returns default",
		ArrangeInput:  args.Map{
			"input": "abc",
			"def": byte(99),
		},
		ExpectedInput: "99",
	},
	{
		Title:         "ToByte overflow returns default",
		ArrangeInput:  args.Map{
			"input": "300",
			"def": byte(10),
		},
		ExpectedInput: "10",
	},
	{
		Title:         "ToByte negative returns default",
		ArrangeInput:  args.Map{
			"input": "-1",
			"def": byte(5),
		},
		ExpectedInput: "5",
	},
}

var extToByteDefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "ToByteDefault valid",
		ArrangeInput:  args.Map{"input": "100"},
		ExpectedInput: "100",
	},
	{
		Title:         "ToByteDefault invalid returns 0",
		ArrangeInput:  args.Map{"input": "abc"},
		ExpectedInput: "0",
	},
	{
		Title:         "ToByteDefault overflow returns 0",
		ArrangeInput:  args.Map{"input": "256"},
		ExpectedInput: "0",
	},
}

// ==========================================
// ToInt / ToIntDef / ToIntDefault
// ==========================================

var extToIntTestCases = []coretestcases.CaseV1{
	{
		Title:         "ToInt valid int",
		ArrangeInput:  args.Map{
			"input": "42",
			"def": 0,
		},
		ExpectedInput: "42",
	},
	{
		Title:         "ToInt invalid returns default",
		ArrangeInput:  args.Map{
			"input": "abc",
			"def": -1,
		},
		ExpectedInput: "-1",
	},
}

// ==========================================
// ToInt8 / ToInt8Def
// ==========================================

var extToInt8TestCases = []coretestcases.CaseV1{
	{
		Title:         "ToInt8 valid",
		ArrangeInput:  args.Map{
			"input": "50",
			"def": int8(0),
		},
		ExpectedInput: "50",
	},
	{
		Title:         "ToInt8 invalid returns default",
		ArrangeInput:  args.Map{
			"input": "abc",
			"def": int8(-1),
		},
		ExpectedInput: "-1",
	},
}

// ==========================================
// ToInt16 / ToInt16Default
// ==========================================

var extToInt16TestCases = []coretestcases.CaseV1{
	{
		Title:         "ToInt16 valid",
		ArrangeInput:  args.Map{
			"input": "1000",
			"def": int16(0),
		},
		ExpectedInput: "1000",
	},
	{
		Title:         "ToInt16 invalid returns default",
		ArrangeInput:  args.Map{
			"input": "abc",
			"def": int16(-1),
		},
		ExpectedInput: "-1",
	},
}

// ==========================================
// ToInt32 / ToInt32Def
// ==========================================

var extToInt32TestCases = []coretestcases.CaseV1{
	{
		Title:         "ToInt32 valid",
		ArrangeInput:  args.Map{
			"input": "65536",
			"def": int32(0),
		},
		ExpectedInput: "65536",
	},
	{
		Title:         "ToInt32 invalid returns default",
		ArrangeInput:  args.Map{
			"input": "abc",
			"def": int32(-1),
		},
		ExpectedInput: "-1",
	},
}

// ==========================================
// ToUint16Default / ToUint32Default
// ==========================================

var extToUint16DefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "ToUint16Default valid",
		ArrangeInput:  args.Map{"input": "5000"},
		ExpectedInput: "5000",
	},
	{
		Title:         "ToUint16Default invalid returns 0",
		ArrangeInput:  args.Map{"input": "abc"},
		ExpectedInput: "0",
	},
	{
		Title:         "ToUint16Default overflow returns 0",
		ArrangeInput:  args.Map{"input": "70000"},
		ExpectedInput: "0",
	},
	{
		Title:         "ToUint16Default negative returns 0",
		ArrangeInput:  args.Map{"input": "-1"},
		ExpectedInput: "0",
	},
}

var extToUint32DefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "ToUint32Default valid",
		ArrangeInput:  args.Map{"input": "100000"},
		ExpectedInput: "100000",
	},
	{
		Title:         "ToUint32Default invalid returns 0",
		ArrangeInput:  args.Map{"input": "abc"},
		ExpectedInput: "0",
	},
	{
		Title:         "ToUint32Default negative returns 0",
		ArrangeInput:  args.Map{"input": "-5"},
		ExpectedInput: "0",
	},
}

// ==========================================
// AnyToString / AnyToStringNameField / AnyToTypeString
// ==========================================

var extAnyToStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "AnyToString nil returns empty",
		ArrangeInput:  args.Map{"input": nil},
		ExpectedInput: "",
	},
	{
		Title:         "AnyToString int returns string",
		ArrangeInput:  args.Map{"input": 42},
		ExpectedInput: "42",
	},
	{
		Title:         "AnyToString string returns string",
		ArrangeInput:  args.Map{"input": "hello"},
		ExpectedInput: "hello",
	},
}

// ==========================================
// MaskLine / MaskTrimLine / MaskLines / MaskTrimLines
// ==========================================

var extMaskLineTestCases = []coretestcases.CaseV1{
	{
		Title: "MaskLine returns mask for empty line",
		ArrangeInput: args.Map{
			"mask": "----------",
			"line": "",
		},
		ExpectedInput: "----------",
	},
	{
		Title: "MaskLine pads short line with mask",
		ArrangeInput: args.Map{
			"mask": "----------",
			"line": "hi",
		},
		ExpectedInput: "hi--------",
	},
	{
		Title: "MaskLine returns line if longer than mask",
		ArrangeInput: args.Map{
			"mask": "---",
			"line": "hello world",
		},
		ExpectedInput: "hello world",
	},
	{
		Title: "MaskLine returns line if mask empty",
		ArrangeInput: args.Map{
			"mask": "",
			"line": "hello",
		},
		ExpectedInput: "hello",
	},
}

// ==========================================
// IsContains
// ==========================================

var extIsContainsTestCases = []coretestcases.CaseV1{
	{
		Title: "IsContains true case-sensitive",
		ArrangeInput: args.Map{
			"lines":       []string{"a", "b", "c"},
			"find":        "b",
			"start":       0,
			"caseSensitive": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsContains false case-sensitive",
		ArrangeInput: args.Map{
			"lines":       []string{"a", "b", "c"},
			"find":        "B",
			"start":       0,
			"caseSensitive": true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsContains true case-insensitive",
		ArrangeInput: args.Map{
			"lines":       []string{"Hello", "World"},
			"find":        "hello",
			"start":       0,
			"caseSensitive": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsContains nil lines returns false",
		ArrangeInput: args.Map{
			"lines":       nil,
			"find":        "x",
			"start":       0,
			"caseSensitive": true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsContains empty lines returns false",
		ArrangeInput: args.Map{
			"lines":       []string{},
			"find":        "x",
			"start":       0,
			"caseSensitive": true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsContains starts at index skips earlier",
		ArrangeInput: args.Map{
			"lines":       []string{"a", "b", "c"},
			"find":        "a",
			"start":       1,
			"caseSensitive": true,
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// SplitLeftRightTrimmed
// ==========================================

var extSplitLeftRightTrimmedTestCases = []coretestcases.CaseV1{
	{
		Title: "SplitLeftRightTrimmed trims whitespace",
		ArrangeInput: args.Map{
			"input":     " key = value ",
			"separator": "=",
		},
		ExpectedInput: args.Map{
			"left": "key",
			"right": "value",
		},
	},
	{
		Title: "SplitLeftRightTrimmed no separator returns trimmed left",
		ArrangeInput: args.Map{
			"input":     " nosep ",
			"separator": "=",
		},
		ExpectedInput: args.Map{
			"left": "nosep",
			"right": "",
		},
	},
}

// ==========================================
// SplitFirstLast
// ==========================================

var extSplitFirstLastTestCases = []coretestcases.CaseV1{
	{
		Title: "SplitFirstLast single separator",
		ArrangeInput: args.Map{
			"input":     "a/b/c",
			"separator": "/",
		},
		ExpectedInput: args.Map{
			"first": "a",
			"last": "c",
		},
	},
	{
		Title: "SplitFirstLast no separator",
		ArrangeInput: args.Map{
			"input":     "abc",
			"separator": "/",
		},
		ExpectedInput: args.Map{
			"first": "abc",
			"last": "",
		},
	},
}

// ==========================================
// SafeSubstringStarts / SafeSubstringEnds
// ==========================================

var extSafeSubstringStartsTestCases = []coretestcases.CaseV1{
	{
		Title:         "SafeSubstringStarts normal",
		ArrangeInput:  args.Map{
			"content": "hello",
			"start": 2,
		},
		ExpectedInput: "llo",
	},
	{
		Title:         "SafeSubstringStarts -1 returns full",
		ArrangeInput:  args.Map{
			"content": "hello",
			"start": -1,
		},
		ExpectedInput: "hello",
	},
	{
		Title:         "SafeSubstringStarts empty returns empty",
		ArrangeInput:  args.Map{
			"content": "",
			"start": 0,
		},
		ExpectedInput: "",
	},
	{
		Title:         "SafeSubstringStarts out of range returns empty",
		ArrangeInput:  args.Map{
			"content": "hi",
			"start": 5,
		},
		ExpectedInput: "",
	},
}

var extSafeSubstringEndsTestCases = []coretestcases.CaseV1{
	{
		Title:         "SafeSubstringEnds normal",
		ArrangeInput:  args.Map{
			"content": "hello",
			"end": 3,
		},
		ExpectedInput: "hel",
	},
	{
		Title:         "SafeSubstringEnds -1 returns full",
		ArrangeInput:  args.Map{
			"content": "hello",
			"end": -1,
		},
		ExpectedInput: "hello",
	},
	{
		Title:         "SafeSubstringEnds empty returns empty",
		ArrangeInput:  args.Map{
			"content": "",
			"end": 3,
		},
		ExpectedInput: "",
	},
	{
		Title:         "SafeSubstringEnds over length returns full",
		ArrangeInput:  args.Map{
			"content": "hi",
			"end": 10,
		},
		ExpectedInput: "hi",
	},
}

// ==========================================
// RemoveManyBySplitting
// ==========================================

var extRemoveManyBySplittingTestCases = []coretestcases.CaseV1{
	{
		Title: "RemoveManyBySplitting removes then splits",
		ArrangeInput: args.Map{
			"content":  "hello-world-test",
			"splitsBy": "-",
			"removes":  []string{"world-"},
		},
		ExpectedInput: "hello,test",
	},
}

// ==========================================
// ReplaceTemplate methods
// ==========================================

var extReplaceTemplateCurlyOneTestCases = []coretestcases.CaseV1{
	{
		Title: "CurlyOne replaces single key",
		ArrangeInput: args.Map{
			"format": "Hello {name}!",
			"key":    "name",
			"value":  "World",
		},
		ExpectedInput: "Hello World!",
	},
	{
		Title: "CurlyOne empty format returns empty",
		ArrangeInput: args.Map{
			"format": "",
			"key":    "name",
			"value":  "World",
		},
		ExpectedInput: "",
	},
}

var extReplaceTemplateCurlyTwoTestCases = []coretestcases.CaseV1{
	{
		Title: "CurlyTwo replaces two keys",
		ArrangeInput: args.Map{
			"format": "{greeting} {name}!",
			"key1":   "greeting",
			"val1":   "Hi",
			"key2":   "name",
			"val2":   "Alice",
		},
		ExpectedInput: "Hi Alice!",
	},
}

var extReplaceWhiteSpacesTestCases = []coretestcases.CaseV1{
	{
		Title: "ReplaceWhiteSpaces removes all whitespace",
		ArrangeInput: args.Map{
			"input": "  some  nothing    \t",
		},
		ExpectedInput: "somenothing",
	},
	{
		Title: "ReplaceWhiteSpaces empty returns empty",
		ArrangeInput: args.Map{
			"input": "   ",
		},
		ExpectedInput: "",
	},
}

// ==========================================
// MaskLines / MaskTrimLines
// ==========================================

var extMaskLinesTestCases = []coretestcases.CaseV1{
	{
		Title: "MaskLines pads all lines",
		ArrangeInput: args.Map{
			"mask":  "----------",
			"lines": []string{"hi", "hey", ""},
		},
		ExpectedInput: "hi--------,hey-------,----------",
	},
	{
		Title: "MaskLines empty mask returns lines",
		ArrangeInput: args.Map{
			"mask":  "",
			"lines": []string{"hello"},
		},
		ExpectedInput: "hello",
	},
}

var extMaskTrimLinesTestCases = []coretestcases.CaseV1{
	{
		Title: "MaskTrimLines trims then pads",
		ArrangeInput: args.Map{
			"mask":  "----------",
			"lines": []string{"  hi  ", "  hey  "},
		},
		ExpectedInput: "hi--------,hey-------",
	},
}
