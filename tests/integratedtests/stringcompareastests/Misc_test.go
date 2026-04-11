package stringcompareastests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ==========================================================================
// VerifyMessage — ignore case branches
// ==========================================================================

var covVerifyMessageTestCases = []coretestcases.CaseV1{
	{
		Title: "VerifyMessage match returns empty -- case sensitive",
		ArrangeInput: args.Map{
			"variant": stringcompareas.Equal, "ignoreCase": false,
			"content": "hello", "search": "hello",
		},
		ExpectedInput: args.Map{
			"isEmpty": true,
			"isNegativeMsg": false,
		},
	},
	{
		Title: "VerifyMessage positive mismatch -- case sensitive",
		ArrangeInput: args.Map{
			"variant": stringcompareas.Equal, "ignoreCase": false,
			"content": "hello", "search": "world",
		},
		ExpectedInput: args.Map{
			"isEmpty": false,
			"isNegativeMsg": false,
		},
	},
	{
		Title: "VerifyMessage positive mismatch -- case ignored",
		ArrangeInput: args.Map{
			"variant": stringcompareas.Equal, "ignoreCase": true,
			"content": "Hello", "search": "world",
		},
		ExpectedInput: args.Map{
			"isEmpty": false,
			"isNegativeMsg": false,
		},
	},
	{
		Title: "VerifyMessage negative mismatch -- case sensitive",
		ArrangeInput: args.Map{
			"variant": stringcompareas.NotEqual, "ignoreCase": false,
			"content": "hello", "search": "hello",
		},
		ExpectedInput: args.Map{
			"isEmpty": false,
			"isNegativeMsg": true,
		},
	},
	{
		Title: "VerifyMessage negative mismatch -- case ignored",
		ArrangeInput: args.Map{
			"variant": stringcompareas.NotEqual, "ignoreCase": true,
			"content": "hello", "search": "hello",
		},
		ExpectedInput: args.Map{
			"isEmpty": false,
			"isNegativeMsg": true,
		},
	},
}

func Test_VerifyMessage(t *testing.T) {
	for caseIndex, testCase := range covVerifyMessageTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		variant := input["variant"].(stringcompareas.Variant)
		ignoreCase := input["ignoreCase"].(bool)
		content := input["content"].(string)
		search := input["search"].(string)

		// Act
		msg := variant.VerifyMessage(ignoreCase, content, search)

		// Assert
		actual := args.Map{
			"isEmpty":      msg == "",
			"isNegativeMsg": len(msg) > 0 && variant.IsNegativeCondition(),
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// VerifyError
// ==========================================================================

var covVerifyErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "VerifyError match returns nil",
		ArrangeInput: args.Map{
			"variant": stringcompareas.StartsWith, "ignoreCase": false,
			"content": "hello world", "search": "hello",
		},
		ExpectedInput: args.Map{"isNil": true},
	},
	{
		Title: "VerifyError mismatch returns error",
		ArrangeInput: args.Map{
			"variant": stringcompareas.StartsWith, "ignoreCase": false,
			"content": "hello world", "search": "world",
		},
		ExpectedInput: args.Map{"isNil": false},
	},
}

func Test_VerifyError(t *testing.T) {
	for caseIndex, testCase := range covVerifyErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		variant := input["variant"].(stringcompareas.Variant)
		ignoreCase := input["ignoreCase"].(bool)
		content := input["content"].(string)
		search := input["search"].(string)

		err := variant.VerifyError(ignoreCase, content, search)

		// Act
		actual := args.Map{"isNil": err == nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// OnlySupportedMsgErr
// ==========================================================================

var covOnlySupportedMsgErrTestCases = []coretestcases.CaseV1{
	{
		Title:         "OnlySupportedMsgErr with partial names returns error",
		ArrangeInput:  args.Map{"names": []string{"Equal"}},
		ExpectedInput: args.Map{"hasError": true},
	},
}

func Test_OnlySupportedMsgErr(t *testing.T) {
	for caseIndex, testCase := range covOnlySupportedMsgErrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		names := input["names"].([]string)

		err := stringcompareas.Equal.OnlySupportedMsgErr("test message", names...)

		// Act
		actual := args.Map{"hasError": err != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// Compare funcs — isIgnoreCase=true branches
// ==========================================================================

var covIgnoreCaseCompareFuncTestCases = []coretestcases.CaseV1{
	{
		Title: "Anywhere ignore case matches",
		ArrangeInput: args.Map{
			"variant": stringcompareas.Anywhere, "content": "Hello World", "search": "hello",
		},
		ExpectedInput: args.Map{"match": true},
	},
	{
		Title: "Contains ignore case matches",
		ArrangeInput: args.Map{
			"variant": stringcompareas.Contains, "content": "Hello World", "search": "hello",
		},
		ExpectedInput: args.Map{"match": true},
	},
	{
		Title: "AnyChars ignore case matches",
		ArrangeInput: args.Map{
			"variant": stringcompareas.AnyChars, "content": "Hello", "search": "HLO",
		},
		ExpectedInput: args.Map{"match": true},
	},
	{
		Title: "NotEqual ignore case same returns false",
		ArrangeInput: args.Map{
			"variant": stringcompareas.NotEqual, "content": "Hello", "search": "hello",
		},
		ExpectedInput: args.Map{"match": false},
	},
	{
		Title: "NotEqual ignore case different returns true",
		ArrangeInput: args.Map{
			"variant": stringcompareas.NotEqual, "content": "Hello", "search": "world",
		},
		ExpectedInput: args.Map{"match": true},
	},
	{
		Title: "NotContains ignore case returns false when found",
		ArrangeInput: args.Map{
			"variant": stringcompareas.NotContains, "content": "Hello World", "search": "hello",
		},
		ExpectedInput: args.Map{"match": false},
	},
	{
		Title: "NotContains ignore case returns true when not found",
		ArrangeInput: args.Map{
			"variant": stringcompareas.NotContains, "content": "Hello World", "search": "xyz",
		},
		ExpectedInput: args.Map{"match": true},
	},
	{
		Title: "NotAnyChars ignore case returns false when chars found",
		ArrangeInput: args.Map{
			"variant": stringcompareas.NotAnyChars, "content": "Hello", "search": "HLO",
		},
		ExpectedInput: args.Map{"match": false},
	},
	{
		Title: "NotAnyChars ignore case returns true when chars not found",
		ArrangeInput: args.Map{
			"variant": stringcompareas.NotAnyChars, "content": "Hello", "search": "XYZ",
		},
		ExpectedInput: args.Map{"match": true},
	},
	{
		Title: "StartsWith ignore case matches",
		ArrangeInput: args.Map{
			"variant": stringcompareas.StartsWith, "content": "Hello World", "search": "hello",
		},
		ExpectedInput: args.Map{"match": true},
	},
	{
		Title: "EndsWith ignore case matches",
		ArrangeInput: args.Map{
			"variant": stringcompareas.EndsWith, "content": "Hello World", "search": "WORLD",
		},
		ExpectedInput: args.Map{"match": true},
	},
	{
		Title: "NotStartsWith ignore case matches when not prefix",
		ArrangeInput: args.Map{
			"variant": stringcompareas.NotStartsWith, "content": "Hello World", "search": "WORLD",
		},
		ExpectedInput: args.Map{"match": true},
	},
	{
		Title: "NotEndsWith ignore case matches when not suffix",
		ArrangeInput: args.Map{
			"variant": stringcompareas.NotEndsWith, "content": "Hello World", "search": "HELLO",
		},
		ExpectedInput: args.Map{"match": true},
	},
	{
		Title: "NotMatchRegex matches when regex does not match",
		ArrangeInput: args.Map{
			"variant": stringcompareas.NotMatchRegex, "content": "hello", "search": `^\d+$`,
		},
		ExpectedInput: args.Map{"match": true},
	},
	{
		Title: "NotMatchRegex returns false when regex matches",
		ArrangeInput: args.Map{
			"variant": stringcompareas.NotMatchRegex, "content": "123", "search": `^\d+$`,
		},
		ExpectedInput: args.Map{"match": false},
	},
}

func Test_CompareFuncs_IgnoreCase(t *testing.T) {
	for caseIndex, testCase := range covIgnoreCaseCompareFuncTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		variant := input["variant"].(stringcompareas.Variant)
		content := input["content"].(string)
		search := input["search"].(string)

		match := variant.IsCompareSuccess(true, content, search)

		// Act
		actual := args.Map{"match": match}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// IsAnyEnumsEqual — no match path
// ==========================================================================

var covIsAnyEnumsEqualTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsAnyEnumsEqual returns false when no match",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"match": false},
	},
}

func Test_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	for caseIndex, testCase := range covIsAnyEnumsEqualTestCases {
		// Arrange
		a := stringcompareas.Equal
		b := stringcompareas.StartsWith
		c := stringcompareas.EndsWith

		// Act
		actual := args.Map{"match": a.IsAnyEnumsEqual(&b, &c)}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// MinInt
// ==========================================================================

var covMinIntTestCases = []coretestcases.CaseV1{
	{
		Title:         "MinInt returns 0",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"value": fmt.Sprintf("%d", 0)},
	},
}

func Test_MinInt(t *testing.T) {
	for caseIndex, testCase := range covMinIntTestCases {
		// Act
		actual := args.Map{"value": fmt.Sprintf("%d", stringcompareas.Equal.MinInt())}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// ValueByte
// ==========================================================================

var covValueByteTestCases = []coretestcases.CaseV1{
	{
		Title:         "ValueByte returns byte value",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"value": byte(0)},
	},
}

func Test_ValueByte(t *testing.T) {
	for caseIndex, testCase := range covValueByteTestCases {
		// Arrange
		v := stringcompareas.Equal

		// Act
		actual := args.Map{"value": v.ValueByte()}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// UnmarshalJSON error path
// ==========================================================================

var covUnmarshalJSONErrTestCases = []coretestcases.CaseV1{
	{
		Title:         "UnmarshalJSON invalid data returns error",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"hasError": true},
	},
}

func Test_UnmarshalJSON_Error(t *testing.T) {
	for caseIndex, testCase := range covUnmarshalJSONErrTestCases {
		// Arrange
		v := stringcompareas.Equal
		err := v.UnmarshalJSON([]byte("invalid-not-json"))

		// Act
		actual := args.Map{"hasError": err != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
