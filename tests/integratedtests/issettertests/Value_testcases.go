package issettertests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/issetter"
)

type getSetInput struct {
	condition bool
	trueVal   issetter.Value
	falseVal  issetter.Value
}

var valueNewTestCases = []coretestcases.CaseV1{
	{
		Title:         "Value.New returns True -- 'true' string input",
		ArrangeInput:  "true",
		ExpectedInput: args.Map{
			"hasError": false,
			"name": "True",
		},
	},
	{
		Title:         "Value.New returns False -- 'false' string input",
		ArrangeInput:  "false",
		ExpectedInput: args.Map{
			"hasError": false,
			"name": "False",
		},
	},
	{
		Title:         "Value.New returns True -- 'yes' string input",
		ArrangeInput:  "yes",
		ExpectedInput: args.Map{
			"hasError": false,
			"name": "True",
		},
	},
	{
		Title:         "Value.New returns False -- 'no' string input",
		ArrangeInput:  "no",
		ExpectedInput: args.Map{
			"hasError": false,
			"name": "False",
		},
	},
	{
		Title:         "Value.New returns Set -- 'Set' string input",
		ArrangeInput:  "Set",
		ExpectedInput: args.Map{
			"hasError": false,
			"name": "Set",
		},
	},
	{
		Title:         "Value.New returns Unset -- 'Unset' string input",
		ArrangeInput:  "Unset",
		ExpectedInput: args.Map{
			"hasError": false,
			"name": "Unset",
		},
	},
	{
		Title:         "Value.New returns Wildcard -- '*' string input",
		ArrangeInput:  "*",
		ExpectedInput: args.Map{
			"hasError": false,
			"name": "Wildcard",
		},
	},
	{
		Title:         "Value.New returns Uninitialized -- empty string input",
		ArrangeInput:  "",
		ExpectedInput: args.Map{
			"hasError": false,
			"name": "Uninitialized",
		},
	},
	{
		Title:         "Value.New returns error -- invalid string input",
		ArrangeInput:  "invalid_value_xyz",
		ExpectedInput: args.Map{
			"hasError": true,
			"name": "Uninitialized",
		},
	},
}

var getBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetBool returns True -- true input",
		ArrangeInput:  true,
		ExpectedInput: "True",
	},
	{
		Title:         "GetBool returns False -- false input",
		ArrangeInput:  false,
		ExpectedInput: "False",
	},
}

var newBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "NewBool returns True -- true input",
		ArrangeInput:  true,
		ExpectedInput: "True",
	},
	{
		Title:         "NewBool returns False -- false input",
		ArrangeInput:  false,
		ExpectedInput: "False",
	},
}

// booleanLogicTestCases
// Expected keys: isOn, isOff, isTrue, isFalse, isSet, isUnset, isValid, isWildcard
var booleanLogicTestCases = []coretestcases.CaseV1{
	{
		Title:        "Value returns all false -- Uninitialized variant",
		ArrangeInput: issetter.Uninitialized,
		ExpectedInput: args.Map{
			"isOn": false, "isOff": false, "isTrue": false, "isFalse": false,
			"isSet": false, "isUnset": false, "isValid": false, "isWildcard": false,
		},
	},
	{
		Title:        "Value returns isOn true and isTrue true -- True variant",
		ArrangeInput: issetter.True,
		ExpectedInput: args.Map{
			"isOn": true, "isOff": false, "isTrue": true, "isFalse": false,
			"isSet": false, "isUnset": false, "isValid": true, "isWildcard": false,
		},
	},
	{
		Title:        "Value returns isOff true and isFalse true -- False variant",
		ArrangeInput: issetter.False,
		ExpectedInput: args.Map{
			"isOn": false, "isOff": true, "isTrue": false, "isFalse": true,
			"isSet": false, "isUnset": false, "isValid": true, "isWildcard": false,
		},
	},
	{
		Title:        "Value returns isOn true and isSet true -- Set variant",
		ArrangeInput: issetter.Set,
		ExpectedInput: args.Map{
			"isOn": true, "isOff": false, "isTrue": false, "isFalse": false,
			"isSet": true, "isUnset": false, "isValid": true, "isWildcard": false,
		},
	},
	{
		Title:        "Value returns isOff true and isUnset true -- Unset variant",
		ArrangeInput: issetter.Unset,
		ExpectedInput: args.Map{
			"isOn": false, "isOff": true, "isTrue": false, "isFalse": false,
			"isSet": false, "isUnset": true, "isValid": true, "isWildcard": false,
		},
	},
	{
		Title:        "Value returns isWildcard true and isValid true -- Wildcard variant",
		ArrangeInput: issetter.Wildcard,
		ExpectedInput: args.Map{
			"isOn": false, "isOff": false, "isTrue": false, "isFalse": false,
			"isSet": false, "isUnset": false, "isValid": true, "isWildcard": true,
		},
	},
}

var combinedBooleansTestCases = []coretestcases.CaseV1{
	{
		Title:         "AllTrue returns True -- all true input",
		ArrangeInput:  []bool{true, true, true},
		ExpectedInput: "True",
	},
	{
		Title:         "AllTrue returns False -- one false in input",
		ArrangeInput:  []bool{true, false, true},
		ExpectedInput: "False",
	},
	{
		Title:         "AllTrue returns True -- empty slice input",
		ArrangeInput:  []bool{},
		ExpectedInput: "True",
	},
	{
		Title:         "AllTrue returns True -- single true input",
		ArrangeInput:  []bool{true},
		ExpectedInput: "True",
	},
	{
		Title:         "AllTrue returns False -- single false input",
		ArrangeInput:  []bool{false},
		ExpectedInput: "False",
	},
}

// conversionTestCases
// Expected keys: toBooleanValue, toSetUnsetValue
var conversionTestCases = []coretestcases.CaseV1{
	{
		Title:         "ToBooleanValue returns True and ToSetUnsetValue returns Set -- True variant",
		ArrangeInput:  issetter.True,
		ExpectedInput: args.Map{
			"toBooleanValue": "True",
			"toSetUnsetValue": "Set",
		},
	},
	{
		Title:         "ToBooleanValue returns False and ToSetUnsetValue returns Unset -- False variant",
		ArrangeInput:  issetter.False,
		ExpectedInput: args.Map{
			"toBooleanValue": "False",
			"toSetUnsetValue": "Unset",
		},
	},
	{
		Title:         "ToBooleanValue returns True and ToSetUnsetValue returns Set -- Set variant",
		ArrangeInput:  issetter.Set,
		ExpectedInput: args.Map{
			"toBooleanValue": "True",
			"toSetUnsetValue": "Set",
		},
	},
	{
		Title:         "ToBooleanValue returns False and ToSetUnsetValue returns Unset -- Unset variant",
		ArrangeInput:  issetter.Unset,
		ExpectedInput: args.Map{
			"toBooleanValue": "False",
			"toSetUnsetValue": "Unset",
		},
	},
	{
		Title:         "ToBooleanValue returns Wildcard and ToSetUnsetValue returns Wildcard -- Wildcard variant",
		ArrangeInput:  issetter.Wildcard,
		ExpectedInput: args.Map{
			"toBooleanValue": "Wildcard",
			"toSetUnsetValue": "Wildcard",
		},
	},
	{
		Title:         "ToBooleanValue returns Uninitialized -- Uninitialized variant",
		ArrangeInput:  issetter.Uninitialized,
		ExpectedInput: args.Map{
			"toBooleanValue": "Uninitialized",
			"toSetUnsetValue": "Uninitialized",
		},
	},
}

var getSetTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetSet returns True -- condition true with True/False pair",
		ArrangeInput:  getSetInput{condition: true, trueVal: issetter.True, falseVal: issetter.False},
		ExpectedInput: "True",
	},
	{
		Title:         "GetSet returns False -- condition false with True/False pair",
		ArrangeInput:  getSetInput{condition: false, trueVal: issetter.True, falseVal: issetter.False},
		ExpectedInput: "False",
	},
	{
		Title:         "GetSet returns Set -- condition true with Set/Unset pair",
		ArrangeInput:  getSetInput{condition: true, trueVal: issetter.Set, falseVal: issetter.Unset},
		ExpectedInput: "Set",
	},
}

var isOutOfRangeTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsOutOfRange returns false -- byte 0 (Uninitialized, in range)",
		ArrangeInput:  byte(0),
		ExpectedInput: "false",
	},
	{
		Title:         "IsOutOfRange returns true -- byte 5 (beyond max valid value)",
		ArrangeInput:  byte(5),
		ExpectedInput: "true",
	},
	{
		Title:         "IsOutOfRange returns true -- byte 6",
		ArrangeInput:  byte(6),
		ExpectedInput: "true",
	},
	{
		Title:         "IsOutOfRange returns true -- byte 255",
		ArrangeInput:  byte(255),
		ExpectedInput: "true",
	},
}
