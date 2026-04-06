package corecomparatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var baseIsCaseSensitiveTestCases = []coretestcases.CaseV1{
	{
		Title:         "Base.IsCaseSensitive returns true -- IsIgnoreCase false",
		ArrangeInput:  args.Map{"isCaseSensitive": true},
		ExpectedInput: args.Map{
			"isIgnoreCase": false,
			"cloneMatch": true,
		},
	},
	{
		Title:         "Base.IsCaseSensitive returns false -- IsIgnoreCase true",
		ArrangeInput:  args.Map{"isCaseSensitive": false},
		ExpectedInput: args.Map{
			"isIgnoreCase": true,
			"cloneMatch": true,
		},
	},
}

var baseIsIgnoreCaseTestCases = []coretestcases.CaseV1{
	{
		Title:         "Base.IsIgnoreCase returns true -- IsCaseSensitive false",
		ArrangeInput:  args.Map{"isIgnoreCase": true},
		ExpectedInput: args.Map{
			"isCaseSensitive": false,
			"cloneMatch": true,
		},
	},
	{
		Title:         "Base.IsIgnoreCase returns false -- IsCaseSensitive true",
		ArrangeInput:  args.Map{"isIgnoreCase": false},
		ExpectedInput: args.Map{
			"isCaseSensitive": true,
			"cloneMatch": true,
		},
	},
}

var compareIsMethodTestCases = []coretestcases.CaseV1{
	{
		Title:        "Compare.Is returns true -- Equal vs Equal",
		ArrangeInput: args.Map{
			"value": 0,
			"other": 0,
		},
		ExpectedInput: args.Map{
			"is":                  true,
			"isInvalid":           false,
			"isValueEqual":        true,
			"isLeftGreater":       false,
			"isLeftGreaterEqual":  false,
			"isLeftLessEqual":     false,
			"isLeftLessOrLeOrEq":  true,
			"isDefinedPlus":       true,
			"isNotInconclusive":   true,
			"rangeNamesCsvNotEmpty": true,
			"sqlOpNotEmpty":       true,
			"stringValueNotEmpty": true,
			"valueInt8":           int8(0),
			"valueInt16":          int16(0),
			"valueInt32":          int32(0),
			"valueString":        "0",
			"formatPanic":        true,
		},
	},
	{
		Title:        "Compare.Is returns false -- LeftGreater vs Equal",
		ArrangeInput: args.Map{
			"value": 1,
			"other": 0,
		},
		ExpectedInput: args.Map{
			"is":                  false,
			"isInvalid":           false,
			"isValueEqual":        false,
			"isLeftGreater":       true,
			"isLeftGreaterEqual":  false,
			"isLeftLessEqual":     false,
			"isLeftLessOrLeOrEq":  false,
			"isDefinedPlus":       false,
			"isNotInconclusive":   true,
			"rangeNamesCsvNotEmpty": true,
			"sqlOpNotEmpty":       true,
			"stringValueNotEmpty": true,
			"valueInt8":           int8(1),
			"valueInt16":          int16(1),
			"valueInt32":          int32(1),
			"valueString":        "1",
			"formatPanic":        true,
		},
	},
	{
		Title:        "Compare.Is returns true -- Inconclusive vs Inconclusive",
		ArrangeInput: args.Map{
			"value": 6,
			"other": 6,
		},
		ExpectedInput: args.Map{
			"is":                  true,
			"isInvalid":           true,
			"isValueEqual":        true,
			"isLeftGreater":       false,
			"isLeftGreaterEqual":  false,
			"isLeftLessEqual":     false,
			"isLeftLessOrLeOrEq":  false,
			"isDefinedPlus":       false,
			"isNotInconclusive":   false,
			"rangeNamesCsvNotEmpty": true,
			"sqlOpNotEmpty":       true,
			"stringValueNotEmpty": true,
			"valueInt8":           int8(6),
			"valueInt16":          int16(6),
			"valueInt32":          int32(6),
			"valueString":        "6",
			"formatPanic":        true,
		},
	},
}

var baseIsCaseSensitiveNilTestCases = []coretestcases.CaseV1{
	{
		Title:         "Base.ClonePtr returns nil -- nil receiver",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"isNil": true},
	},
}

var baseIsIgnoreCaseNilTestCases = []coretestcases.CaseV1{
	{
		Title:         "Base.ClonePtr returns nil -- nil receiver",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"isNil": true},
	},
}

var compareUnmarshalJsonTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.UnmarshalJSON returns no error -- valid name 'Equal'",
		ArrangeInput:  args.Map{"data": "Equal"},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title:         "Compare.UnmarshalJSON returns error -- invalid name 'InvalidXyz'",
		ArrangeInput:  args.Map{"data": "InvalidXyz"},
		ExpectedInput: args.Map{"hasError": true},
	},
	{
		Title:         "Compare.UnmarshalJSON returns error -- nil data",
		ArrangeInput:  args.Map{"isNilData": true},
		ExpectedInput: args.Map{"hasError": true},
	},
}
