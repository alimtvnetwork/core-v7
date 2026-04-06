package coremathtests

import (
	"math"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var integerWithinToInt32TestCases = []coretestcases.CaseV1{
	{
		Title:         "ToInt32 within range -- true",
		ArrangeInput:  args.Map{
			"when": "value within int32",
			"value": 100,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "ToInt32 zero -- true",
		ArrangeInput:  args.Map{
			"when": "zero value",
			"value": 0,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "ToInt32 max int32 -- true",
		ArrangeInput:  args.Map{
			"when": "max int32",
			"value": math.MaxInt32,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "ToInt32 min int32 -- true",
		ArrangeInput:  args.Map{
			"when": "min int32",
			"value": math.MinInt32,
		},
		ExpectedInput: args.Map{"result": true},
	},
}

var integerWithinToUint16TestCases = []coretestcases.CaseV1{
	{
		Title:         "ToUnsignedInt16 within range -- true",
		ArrangeInput:  args.Map{
			"when": "255",
			"value": 255,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "ToUnsignedInt16 max -- true",
		ArrangeInput:  args.Map{
			"when": "max uint16",
			"value": math.MaxUint16,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "ToUnsignedInt16 negative -- false",
		ArrangeInput:  args.Map{
			"when": "negative",
			"value": -1,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "ToUnsignedInt16 too large -- false",
		ArrangeInput:  args.Map{
			"when": "too large",
			"value": math.MaxUint16 + 1,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var integerWithinToUint32TestCases = []coretestcases.CaseV1{
	{
		Title:         "ToUnsignedInt32 within range -- true",
		ArrangeInput:  args.Map{
			"when": "zero",
			"value": 0,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "ToUnsignedInt32 negative -- false",
		ArrangeInput:  args.Map{
			"when": "negative",
			"value": -1,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var integerWithinToUint64TestCases = []coretestcases.CaseV1{
	{
		Title:         "ToUnsignedInt64 zero -- true",
		ArrangeInput:  args.Map{
			"when": "zero",
			"value": 0,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "ToUnsignedInt64 positive -- true",
		ArrangeInput:  args.Map{
			"when": "positive",
			"value": 1000,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "ToUnsignedInt64 negative -- false",
		ArrangeInput:  args.Map{
			"when": "negative",
			"value": -1,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var integerOutOfRangeToInt8TestCases = []coretestcases.CaseV1{
	{
		Title:         "OutOfRange ToInt8 within -- false",
		ArrangeInput:  args.Map{
			"when": "within int8",
			"value": 100,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "OutOfRange ToInt8 too large -- true",
		ArrangeInput:  args.Map{
			"when": "too large",
			"value": 200,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "OutOfRange ToInt8 too small -- true",
		ArrangeInput:  args.Map{
			"when": "too small",
			"value": -200,
		},
		ExpectedInput: args.Map{"result": true},
	},
}

var integerOutOfRangeToInt16TestCases = []coretestcases.CaseV1{
	{
		Title:         "OutOfRange ToInt16 within -- false",
		ArrangeInput:  args.Map{
			"when": "within int16",
			"value": 1000,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "OutOfRange ToInt16 too large -- true",
		ArrangeInput:  args.Map{
			"when": "too large",
			"value": math.MaxInt16 + 1,
		},
		ExpectedInput: args.Map{"result": true},
	},
}

var integerOutOfRangeToInt32TestCases = []coretestcases.CaseV1{
	{
		Title:         "OutOfRange ToInt32 within -- false",
		ArrangeInput:  args.Map{
			"when": "within int32",
			"value": 1000,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var integerOutOfRangeToUint16TestCases = []coretestcases.CaseV1{
	{
		Title:         "OutOfRange ToUnsignedInt16 within -- false",
		ArrangeInput:  args.Map{
			"when": "within uint16",
			"value": 1000,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "OutOfRange ToUnsignedInt16 negative -- true",
		ArrangeInput:  args.Map{
			"when": "negative",
			"value": -1,
		},
		ExpectedInput: args.Map{"result": true},
	},
}

var integerOutOfRangeToUint64TestCases = []coretestcases.CaseV1{
	{
		Title:         "OutOfRange ToUnsignedInt64 positive -- false",
		ArrangeInput:  args.Map{
			"when": "positive",
			"value": 100,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "OutOfRange ToUnsignedInt64 negative -- true",
		ArrangeInput:  args.Map{
			"when": "negative",
			"value": -1,
		},
		ExpectedInput: args.Map{"result": true},
	},
}

var maxFloat32TestCases = []coretestcases.CaseV1{
	{
		Title:         "MaxFloat32 left bigger -- returns left",
		ArrangeInput:  args.Map{
			"when": "10.5 vs 5.5",
			"a": 10.5,
			"b": 5.5,
		},
		ExpectedInput: args.Map{"result": 10.5},
	},
	{
		Title:         "MaxFloat32 right bigger -- returns right",
		ArrangeInput:  args.Map{
			"when": "3.0 vs 7.0",
			"a": 3.0,
			"b": 7.0,
		},
		ExpectedInput: args.Map{"result": 7.0},
	},
	{
		Title:         "MaxFloat32 equal -- returns either",
		ArrangeInput:  args.Map{
			"when": "5.0 vs 5.0",
			"a": 5.0,
			"b": 5.0,
		},
		ExpectedInput: args.Map{"result": 5.0},
	},
}

var minFloat32TestCases = []coretestcases.CaseV1{
	{
		Title:         "MinFloat32 left smaller -- returns left",
		ArrangeInput:  args.Map{
			"when": "3.0 vs 7.0",
			"a": 3.0,
			"b": 7.0,
		},
		ExpectedInput: args.Map{"result": 3.0},
	},
	{
		Title:         "MinFloat32 right smaller -- returns right",
		ArrangeInput:  args.Map{
			"when": "10.0 vs 5.0",
			"a": 10.0,
			"b": 5.0,
		},
		ExpectedInput: args.Map{"result": 5.0},
	},
	{
		Title:         "MinFloat32 equal -- returns either",
		ArrangeInput:  args.Map{
			"when": "5.0 vs 5.0",
			"a": 5.0,
			"b": 5.0,
		},
		ExpectedInput: args.Map{"result": 5.0},
	},
}
