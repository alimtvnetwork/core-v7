package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var boolOnceExtendedTestCases = []coretestcases.CaseV1{
	{
		Title:         "BoolOnce NewBoolOnce returns true",
		ArrangeInput:  args.Map{"value": true},
		ExpectedInput: args.Map{
			"execute": true,
			"unmarshalOk": true,
			"serializeOk": true,
		},
	},
	{
		Title:         "BoolOnce NewBoolOnce returns false",
		ArrangeInput:  args.Map{"value": false},
		ExpectedInput: args.Map{
			"execute": false,
			"unmarshalOk": true,
			"serializeOk": true,
		},
	},
}

var byteOnceExtendedTestCases = []coretestcases.CaseV1{
	{
		Title:         "ByteOnce Execute and UnmarshalJSON",
		ArrangeInput:  args.Map{"value": byte(42)},
		ExpectedInput: args.Map{
			"execute": byte(42),
			"unmarshalOk": true,
		},
	},
}

var integerOnceExtendedTestCases = []coretestcases.CaseV1{
	{
		Title:        "IntegerOnce positive value checks",
		ArrangeInput: args.Map{"value": 5},
		ExpectedInput: args.Map{
			"execute":          5,
			"isAboveEqualZero": true,
			"isLessThanEqZero": false,
			"isInvalidIndex":   false,
			"isValidIndex":     true,
			"unmarshalOk":      true,
			"serializeOk":      true,
		},
	},
	{
		Title:        "IntegerOnce negative value checks",
		ArrangeInput: args.Map{"value": -3},
		ExpectedInput: args.Map{
			"execute":          -3,
			"isAboveEqualZero": false,
			"isLessThanEqZero": true,
			"isInvalidIndex":   true,
			"isValidIndex":     false,
			"unmarshalOk":      true,
			"serializeOk":      true,
		},
	},
}

var errorOnceExtendedTestCases = []coretestcases.CaseV1{
	{
		Title:        "ErrorOnce with error -- unmarshal, execute, string, serialize, concat",
		ArrangeInput: args.Map{
			"hasError": true,
			"message": "test error",
		},
		ExpectedInput: args.Map{
			"hasError":    true,
			"isEmptyErr":  false,
			"executeOk":   true,
			"stringOk":    true,
			"serializeOk": true,
			"concatOk":    true,
			"unmarshalOk": true,
		},
	},
	{
		Title:        "ErrorOnce nil error -- handle does not panic",
		ArrangeInput: args.Map{"hasError": false},
		ExpectedInput: args.Map{
			"hasError":    false,
			"isEmptyErr":  true,
			"executeOk":   true,
			"stringOk":    true,
			"serializeOk": true,
			"concatOk":    true,
			"unmarshalOk": true,
		},
	},
}

var anyOnceExtendedTestCases = []coretestcases.CaseV1{
	{
		Title:        "AnyOnce with string value",
		ArrangeInput: args.Map{"value": "hello"},
		ExpectedInput: args.Map{
			"valueStringOk":   true,
			"valueStringMust": true,
			"safeStringOk":    true,
			"castStringOk":    true,
			"isInitialized":   true,
			"isStringEmpty":   false,
			"isStringEmptyWs": false,
			"deserializeOk":   true,
			"serializeOk":     true,
			"serializeMustOk": true,
			"valueOnlyOk":     true,
		},
	},
	{
		Title:        "AnyOnce with nil value",
		ArrangeInput: args.Map{
			"value": nil,
			"isNil": true,
		},
		ExpectedInput: args.Map{
			"isNull":          true,
			"isStringEmpty":   true,
			"isStringEmptyWs": true,
		},
	},
	{
		Title:        "AnyOnce with map value -- cast tests",
		ArrangeInput: args.Map{"valueType": "map"},
		ExpectedInput: args.Map{
			"castMapOk": true,
		},
	},
	{
		Title:        "AnyOnce with []string value -- cast tests",
		ArrangeInput: args.Map{"valueType": "strings"},
		ExpectedInput: args.Map{
			"castStringsOk": true,
		},
	},
	{
		Title:        "AnyOnce with []byte value -- cast tests",
		ArrangeInput: args.Map{"valueType": "bytes"},
		ExpectedInput: args.Map{
			"castBytesOk": true,
		},
	},
	{
		Title:        "AnyOnce with map[string]string -- hashmap cast",
		ArrangeInput: args.Map{"valueType": "hashmapMap"},
		ExpectedInput: args.Map{
			"castHashmapOk": true,
		},
	},
}

var anyErrorOnceExtendedTestCases = []coretestcases.CaseV1{
	{
		Title:        "AnyErrorOnce with value no error",
		ArrangeInput: args.Map{
			"value": "world",
			"hasError": false,
		},
		ExpectedInput: args.Map{
			"valueWithErrorOk": true,
			"executeMustOk":    true,
			"valueStringMust":  true,
			"isInitialized":    true,
			"isStringEmpty":    false,
			"isStringEmptyWs":  false,
			"castStringOk":     true,
			"serializeOk":      true,
			"serializeMustOk":  true,
			"serializeSkipOk":  true,
			"deserializeOk":    true,
		},
	},
	{
		Title:        "AnyErrorOnce with nil value and error -- cast tests",
		ArrangeInput: args.Map{"hasError": true},
		ExpectedInput: args.Map{
			"isNull":        true,
			"hasError":      true,
			"isStringEmpty": true,
		},
	},
	{
		Title:        "AnyErrorOnce cast map",
		ArrangeInput: args.Map{
			"valueType": "map",
			"hasError": false,
		},
		ExpectedInput: args.Map{
			"castMapOk": true,
		},
	},
	{
		Title:        "AnyErrorOnce cast strings",
		ArrangeInput: args.Map{
			"valueType": "strings",
			"hasError": false,
		},
		ExpectedInput: args.Map{
			"castStringsOk": true,
		},
	},
	{
		Title:        "AnyErrorOnce cast bytes",
		ArrangeInput: args.Map{
			"valueType": "bytes",
			"hasError": false,
		},
		ExpectedInput: args.Map{
			"castBytesOk": true,
		},
	},
	{
		Title:        "AnyErrorOnce cast hashmap",
		ArrangeInput: args.Map{
			"valueType": "hashmapMap",
			"hasError": false,
		},
		ExpectedInput: args.Map{
			"castHashmapOk": true,
		},
	},
}

var bytesOnceSerializeTestCases = []coretestcases.CaseV1{
	{
		Title:         "BytesOnce Serialize returns bytes",
		ArrangeInput:  args.Map{"value": "test-data"},
		ExpectedInput: args.Map{"serializeOk": true},
	},
}
