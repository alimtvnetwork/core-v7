package anycmptests

import (
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var sharedPtr = &coretests.DraftType{}

// Branch: left == right (same pointer) → Equal
var cmpSamePointerTestCases = []coretestcases.CaseV1{
	{
		Title:        "Cmp returns Equal -- same pointer reference",
		ArrangeInput: args.Map{
			"when": "same pointer",
			"pair": args.TwoAny{First: sharedPtr, Second: sharedPtr},
		},
		ExpectedInput: args.Map{
			"name": "Equal",
		},
	},
	{
		Title:        "Cmp returns Equal -- same int value (== match)",
		ArrangeInput: args.Map{
			"when": "same int",
			"pair": args.TwoAny{First: 42, Second: 42},
		},
		ExpectedInput: args.Map{
			"name": "Equal",
		},
	},
	{
		Title:        "Cmp returns Equal -- same string value (== match)",
		ArrangeInput: args.Map{
			"when": "same string",
			"pair": args.TwoAny{First: "abc", Second: "abc"},
		},
		ExpectedInput: args.Map{
			"name": "Equal",
		},
	},
}

// Branch: left == nil && right == nil → Equal
var cmpBothNilTestCases = []coretestcases.CaseV1{
	{
		Title:        "Cmp returns Equal -- both untyped nil",
		ArrangeInput: args.Map{"when": "both nil"},
		ExpectedInput: args.Map{
			"name": "Equal",
		},
	},
}

// Branch: left == nil || right == nil → NotEqual
var cmpOneNilTestCases = []coretestcases.CaseV1{
	{
		Title:        "Cmp returns NotEqual -- left nil right int",
		ArrangeInput: args.Map{
			"when": "left nil",
			"pair": args.TwoAny{First: nil, Second: 1},
		},
		ExpectedInput: args.Map{
			"name": "NotEqual",
		},
	},
	{
		Title:        "Cmp returns NotEqual -- left int right nil",
		ArrangeInput: args.Map{
			"when": "right nil",
			"pair": args.TwoAny{First: 1, Second: nil},
		},
		ExpectedInput: args.Map{
			"name": "NotEqual",
		},
	},
	{
		Title:        "Cmp returns NotEqual -- left string right nil",
		ArrangeInput: args.Map{
			"when": "string vs nil",
			"pair": args.TwoAny{First: "hello", Second: nil},
		},
		ExpectedInput: args.Map{
			"name": "NotEqual",
		},
	},
}

// Branch: isLeftNull && isBothEqual (both typed-nil) → Equal
var cmpTypedNilBothNullTestCases = []coretestcases.CaseV1{
	{
		Title:        "Cmp returns Equal -- both typed nil pointers different types",
		ArrangeInput: args.Map{
			"when": "two nil ptrs diff type",
			"pair": args.TwoAny{First: (*int)(nil), Second: (*string)(nil)},
		},
		ExpectedInput: args.Map{
			"name": "Equal",
		},
	},
	{
		Title:        "Cmp returns Equal -- both typed nil channels",
		ArrangeInput: args.Map{
			"when": "two nil chans",
			"pair": args.TwoAny{First: (chan int)(nil), Second: (chan int)(nil)},
		},
		ExpectedInput: args.Map{
			"name": "Equal",
		},
	},
	{
		Title:        "Cmp returns Equal -- both typed nil pointers same type",
		ArrangeInput: args.Map{
			"when": "two nil ptrs",
			"pair": args.TwoAny{First: (*int)(nil), Second: (*int)(nil)},
		},
		ExpectedInput: args.Map{
			"name": "Equal",
		},
	},
}

// Branch: isBothDifferent && (isLeftNull || isRightNull) → NotEqual
var cmpTypedNilOneSideTestCases = []coretestcases.CaseV1{
	{
		Title:        "Cmp returns NotEqual -- typed nil ptr vs non-nil ptr",
		ArrangeInput: args.Map{
			"when": "nil ptr vs non-nil ptr",
			"pair": args.TwoAny{First: (*int)(nil), Second: new(int)},
		},
		ExpectedInput: args.Map{
			"name": "NotEqual",
		},
	},
	{
		Title:        "Cmp returns NotEqual -- non-nil ptr vs typed nil ptr",
		ArrangeInput: args.Map{
			"when": "non-nil ptr vs nil ptr",
			"pair": args.TwoAny{First: new(string), Second: (*string)(nil)},
		},
		ExpectedInput: args.Map{
			"name": "NotEqual",
		},
	},
	{
		Title:        "Cmp returns NotEqual -- non-nil chan vs typed nil chan",
		ArrangeInput: args.Map{
			"when": "non-nil chan vs nil chan",
			"pair": args.TwoAny{First: make(chan int), Second: (chan int)(nil)},
		},
		ExpectedInput: args.Map{
			"name": "NotEqual",
		},
	},
}

// Branch: fallthrough → Inconclusive
var cmpBothNonNilTestCases = []coretestcases.CaseV1{
	{
		Title:        "Cmp returns Inconclusive -- two different ints",
		ArrangeInput: args.Map{
			"when": "different ints",
			"pair": args.TwoAny{First: 1, Second: 2},
		},
		ExpectedInput: args.Map{
			"name": "Inconclusive",
		},
	},
	{
		Title:        "Cmp returns Inconclusive -- two different pointers",
		ArrangeInput: args.Map{
			"when": "different ptrs",
			"pair": args.TwoAny{First: &coretests.DraftType{}, Second: &coretests.DraftType{}},
		},
		ExpectedInput: args.Map{
			"name": "Inconclusive",
		},
	},
	{
		Title:        "Cmp returns Inconclusive -- two different strings",
		ArrangeInput: args.Map{
			"when": "different strings",
			"pair": args.TwoAny{First: "hello", Second: "world"},
		},
		ExpectedInput: args.Map{
			"name": "Inconclusive",
		},
	},
	{
		Title:        "Cmp returns Inconclusive -- two different non-nil channels",
		ArrangeInput: args.Map{
			"when": "different chans",
			"pair": args.TwoAny{First: make(chan int), Second: make(chan int)},
		},
		ExpectedInput: args.Map{
			"name": "Inconclusive",
		},
	},
}

