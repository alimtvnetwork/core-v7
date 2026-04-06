package issettertests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/issetter"
)

// =============================================================================
// IsOnLogically test cases
// =============================================================================

var isOnLogicallyTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsOnLogically - Uninitialized returns false",
		ArrangeInput:  args.Map{"value": issetter.Uninitialized},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsOnLogically - True returns true",
		ArrangeInput:  args.Map{"value": issetter.True},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsOnLogically - False returns false",
		ArrangeInput:  args.Map{"value": issetter.False},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsOnLogically - Unset returns false",
		ArrangeInput:  args.Map{"value": issetter.Unset},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsOnLogically - Set returns true",
		ArrangeInput:  args.Map{"value": issetter.Set},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsOnLogically - Wildcard returns false",
		ArrangeInput:  args.Map{"value": issetter.Wildcard},
		ExpectedInput: args.Map{"result": false},
	},
}

// =============================================================================
// IsOffLogically test cases
// =============================================================================

var isOffLogicallyTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsOffLogically - Uninitialized returns false",
		ArrangeInput:  args.Map{"value": issetter.Uninitialized},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsOffLogically - True returns false",
		ArrangeInput:  args.Map{"value": issetter.True},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsOffLogically - False returns true",
		ArrangeInput:  args.Map{"value": issetter.False},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsOffLogically - Unset returns true",
		ArrangeInput:  args.Map{"value": issetter.Unset},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsOffLogically - Set returns false",
		ArrangeInput:  args.Map{"value": issetter.Set},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsOffLogically - Wildcard returns false",
		ArrangeInput:  args.Map{"value": issetter.Wildcard},
		ExpectedInput: args.Map{"result": false},
	},
}

// =============================================================================
// WildcardApply test cases
// =============================================================================

var wildcardApplyTestCases = []coretestcases.CaseV1{
	{
		Title:         "WildcardApply - Wildcard passes through true",
		ArrangeInput:  args.Map{
			"value": issetter.Wildcard,
			"input": true,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "WildcardApply - Wildcard passes through false",
		ArrangeInput:  args.Map{
			"value": issetter.Wildcard,
			"input": false,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "WildcardApply - Uninitialized passes through true",
		ArrangeInput:  args.Map{
			"value": issetter.Uninitialized,
			"input": true,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "WildcardApply - Unset passes through false",
		ArrangeInput:  args.Map{
			"value": issetter.Unset,
			"input": false,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "WildcardApply - True ignores input returns true",
		ArrangeInput:  args.Map{
			"value": issetter.True,
			"input": false,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "WildcardApply - False ignores input returns false",
		ArrangeInput:  args.Map{
			"value": issetter.False,
			"input": true,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "WildcardApply - Set ignores input returns false",
		ArrangeInput:  args.Map{
			"value": issetter.Set,
			"input": true,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

// =============================================================================
// IsWildcardOrBool test cases
// =============================================================================

var isWildcardOrBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsWildcardOrBool - Wildcard always true",
		ArrangeInput:  args.Map{
			"value": issetter.Wildcard,
			"input": false,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsWildcardOrBool - True with true",
		ArrangeInput:  args.Map{
			"value": issetter.True,
			"input": true,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsWildcardOrBool - False with false",
		ArrangeInput:  args.Map{
			"value": issetter.False,
			"input": false,
		},
		ExpectedInput: args.Map{"result": true},
	},
}

// =============================================================================
// ToByteCondition test cases
// =============================================================================

var toByteConditionTestCases = []coretestcases.CaseV1{
	{
		Title:         "ToByteCondition - True returns trueVal",
		ArrangeInput:  args.Map{
			"value": issetter.True,
			"trueVal": byte(10),
			"falseVal": byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: args.Map{"result": 10},
	},
	{
		Title:         "ToByteCondition - False returns falseVal",
		ArrangeInput:  args.Map{
			"value": issetter.False,
			"trueVal": byte(10),
			"falseVal": byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: args.Map{"result": 20},
	},
	{
		Title:         "ToByteCondition - Uninitialized returns invalid",
		ArrangeInput:  args.Map{
			"value": issetter.Uninitialized,
			"trueVal": byte(10),
			"falseVal": byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: args.Map{"result": 255},
	},
	{
		Title:         "ToByteCondition - Set returns invalid",
		ArrangeInput:  args.Map{
			"value": issetter.Set,
			"trueVal": byte(10),
			"falseVal": byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: args.Map{"result": 255},
	},
	{
		Title:         "ToByteCondition - Wildcard returns invalid",
		ArrangeInput:  args.Map{
			"value": issetter.Wildcard,
			"trueVal": byte(10),
			"falseVal": byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: args.Map{"result": 255},
	},
}

// =============================================================================
// ToByteConditionWithWildcard test cases
// =============================================================================

var toByteConditionWithWildcardTestCases = []coretestcases.CaseV1{
	{
		Title:         "ToByteConditionWithWildcard - Wildcard returns wildcard byte",
		ArrangeInput:  args.Map{
			"value": issetter.Wildcard,
			"wildcardVal": byte(99),
			"trueVal": byte(10),
			"falseVal": byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: args.Map{"result": 99},
	},
	{
		Title:         "ToByteConditionWithWildcard - True returns trueVal",
		ArrangeInput:  args.Map{
			"value": issetter.True,
			"wildcardVal": byte(99),
			"trueVal": byte(10),
			"falseVal": byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: args.Map{"result": 10},
	},
	{
		Title:         "ToByteConditionWithWildcard - False returns falseVal",
		ArrangeInput:  args.Map{
			"value": issetter.False,
			"wildcardVal": byte(99),
			"trueVal": byte(10),
			"falseVal": byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: args.Map{"result": 20},
	},
	{
		Title:         "ToByteConditionWithWildcard - Uninitialized returns invalid",
		ArrangeInput:  args.Map{
			"value": issetter.Uninitialized,
			"wildcardVal": byte(99),
			"trueVal": byte(10),
			"falseVal": byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: args.Map{"result": 255},
	},
}

// =============================================================================
// IsDefinedLogically test cases
// =============================================================================

var isDefinedLogicallyTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsDefinedLogically - Uninitialized false",
		ArrangeInput:  args.Map{"value": issetter.Uninitialized},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsDefinedLogically - True true",
		ArrangeInput:  args.Map{"value": issetter.True},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsDefinedLogically - False true",
		ArrangeInput:  args.Map{"value": issetter.False},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsDefinedLogically - Unset true",
		ArrangeInput:  args.Map{"value": issetter.Unset},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsDefinedLogically - Set true",
		ArrangeInput:  args.Map{"value": issetter.Set},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsDefinedLogically - Wildcard false",
		ArrangeInput:  args.Map{"value": issetter.Wildcard},
		ExpectedInput: args.Map{"result": false},
	},
}

// =============================================================================
// IsUndefinedLogically test cases
// =============================================================================

var isUndefinedLogicallyTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsUndefinedLogically - Uninitialized true",
		ArrangeInput:  args.Map{"value": issetter.Uninitialized},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsUndefinedLogically - True false",
		ArrangeInput:  args.Map{"value": issetter.True},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsUndefinedLogically - False false",
		ArrangeInput:  args.Map{"value": issetter.False},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsUndefinedLogically - Unset false",
		ArrangeInput:  args.Map{"value": issetter.Unset},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsUndefinedLogically - Set false",
		ArrangeInput:  args.Map{"value": issetter.Set},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsUndefinedLogically - Wildcard true",
		ArrangeInput:  args.Map{"value": issetter.Wildcard},
		ExpectedInput: args.Map{"result": true},
	},
}

// =============================================================================
// IsPositive test cases
// =============================================================================

var isPositiveTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsPositive - Uninitialized false",
		ArrangeInput:  args.Map{"value": issetter.Uninitialized},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsPositive - True true",
		ArrangeInput:  args.Map{"value": issetter.True},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsPositive - False false",
		ArrangeInput:  args.Map{"value": issetter.False},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsPositive - Unset false",
		ArrangeInput:  args.Map{"value": issetter.Unset},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsPositive - Set true",
		ArrangeInput:  args.Map{"value": issetter.Set},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsPositive - Wildcard false",
		ArrangeInput:  args.Map{"value": issetter.Wildcard},
		ExpectedInput: args.Map{"result": false},
	},
}

// =============================================================================
// IsNegative test cases
// =============================================================================

var isNegativeTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsNegative - Uninitialized true",
		ArrangeInput:  args.Map{"value": issetter.Uninitialized},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsNegative - True false",
		ArrangeInput:  args.Map{"value": issetter.True},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsNegative - False true",
		ArrangeInput:  args.Map{"value": issetter.False},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsNegative - Unset true",
		ArrangeInput:  args.Map{"value": issetter.Unset},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsNegative - Set false",
		ArrangeInput:  args.Map{"value": issetter.Set},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsNegative - Wildcard false",
		ArrangeInput:  args.Map{"value": issetter.Wildcard},
		ExpectedInput: args.Map{"result": false},
	},
}

// =============================================================================
// GetSetBoolOnInvalid test cases
// =============================================================================

var getSetBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetSetBoolOnInvalid - already True returns true ignores setter",
		ArrangeInput:  args.Map{
			"initial": issetter.True,
			"setter": false,
		},
		ExpectedInput: args.Map{
			"result": true,
			"isTrueOrFalse": true,
		},
	},
	{
		Title:         "GetSetBoolOnInvalid - already False returns false ignores setter",
		ArrangeInput:  args.Map{
			"initial": issetter.False,
			"setter": true,
		},
		ExpectedInput: args.Map{
			"result": false,
			"isTrueOrFalse": true,
		},
	},
	{
		Title:         "GetSetBoolOnInvalid - Uninitialized with true sets True",
		ArrangeInput:  args.Map{
			"initial": issetter.Uninitialized,
			"setter": true,
		},
		ExpectedInput: args.Map{
			"result": true,
			"isTrueOrFalse": true,
		},
	},
	{
		Title:         "GetSetBoolOnInvalid - Uninitialized with false sets False",
		ArrangeInput:  args.Map{
			"initial": issetter.Uninitialized,
			"setter": false,
		},
		ExpectedInput: args.Map{
			"result": false,
			"isTrueOrFalse": true,
		},
	},
	{
		Title:         "GetSetBoolOnInvalid - Set triggers setter with true",
		ArrangeInput:  args.Map{
			"initial": issetter.Set,
			"setter": true,
		},
		ExpectedInput: args.Map{
			"result": true,
			"isTrueOrFalse": true,
		},
	},
}

// =============================================================================
// LazyEvaluateBool test cases
// =============================================================================

var lazyEvaluateBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "LazyEvaluateBool - Uninitialized calls func sets True",
		ArrangeInput:  args.Map{"initial": issetter.Uninitialized},
		ExpectedInput: args.Map{
			"called": true,
			"returnedTrue": true,
			"isTrue": true,
		},
	},
	{
		Title:         "LazyEvaluateBool - already True skips func",
		ArrangeInput:  args.Map{"initial": issetter.True},
		ExpectedInput: args.Map{
			"called": false,
			"returnedTrue": false,
			"isTrue": true,
		},
	},
	{
		Title:         "LazyEvaluateBool - already False skips func",
		ArrangeInput:  args.Map{"initial": issetter.False},
		ExpectedInput: args.Map{
			"called": false,
			"returnedTrue": false,
			"isTrue": false,
		},
	},
}

// =============================================================================
// LazyEvaluateSet test cases
// =============================================================================

var lazyEvaluateSetTestCases = []coretestcases.CaseV1{
	{
		Title:         "LazyEvaluateSet - Uninitialized calls func sets Set",
		ArrangeInput:  args.Map{"initial": issetter.Uninitialized},
		ExpectedInput: args.Map{
			"called": true,
			"returnedTrue": true,
			"isSet": true,
		},
	},
	{
		Title:         "LazyEvaluateSet - already Set skips func",
		ArrangeInput:  args.Map{"initial": issetter.Set},
		ExpectedInput: args.Map{
			"called": false,
			"returnedTrue": false,
			"isSet": true,
		},
	},
	{
		Title:         "LazyEvaluateSet - already Unset skips func",
		ArrangeInput:  args.Map{"initial": issetter.Unset},
		ExpectedInput: args.Map{
			"called": false,
			"returnedTrue": false,
			"isSet": false,
		},
	},
}
