package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── ValidValue (C20) ──

var srcC20ValidValueConstructorsTestCase = coretestcases.CaseV1{
	Title: "ValidValue constructors NewValidValueUsingAny NewValidValueUsingAnyAutoValid NewValidValueEmpty InvalidValidValueNoMessage InvalidValidValue return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC20ValidValueBytesOnceTestCase = coretestcases.CaseV1{
	Title: "ValidValue ValueBytesOnce ValueBytesOncePtr return correct -- test",
	ExpectedInput: args.Map{
		"b1":    "test",
		"b2":    "test",
		"bPtr":  "test",
	},
}

var srcC20ValidValueStringChecksTestCase = coretestcases.CaseV1{
	Title: "ValidValue IsEmpty IsWhitespace Trim HasValidNonEmpty HasValidNonWhitespace HasSafeNonEmpty return correct -- hello",
	ExpectedInput: args.Map{
		"isEmpty":     false,
		"isWhite":     false,
		"trim":        "hello",
		"validNonE":   true,
		"validNonW":   true,
		"safeNonE":    true,
	},
}

var srcC20ValidValueConversionsTestCase = coretestcases.CaseV1{
	Title: "ValidValue ValueBool ValueInt ValueDefInt ValueByte ValueDefByte ValueFloat64 ValueDefFloat64 return correct -- various",
	ExpectedInput: args.Map{
		"boolTrue":      true,
		"int42":         42,
		"defInt42":      42,
		"byte42":        42,
		"defByte42":     42,
		"float314":      3.14,
		"defFloat314":   3.14,
		"boolInvalid":   false,
		"intDefault":    99,
		"boolEmpty":     false,
	},
}

var srcC20ValidValueComparisonsTestCase = coretestcases.CaseV1{
	Title: "ValidValue Is IsAnyOf IsContains IsAnyContains IsEqualNonSensitive return correct -- hello",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC20ValidValueCloneDisposeTestCase = coretestcases.CaseV1{
	Title: "ValidValue Clone Clear Dispose String FullString return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC20ValidValueJsonTestCase = coretestcases.CaseV1{
	Title: "ValidValue Json JsonPtr Serialize ParseInjectUsingJson return correct -- test",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC20ValidValueSplitTestCase = coretestcases.CaseV1{
	Title: "ValidValue Split SplitNonEmpty SplitTrimNonWhitespace return correct -- a,b,c",
	ExpectedInput: args.Map{
		"splitLen":    3,
		"nonEmptyGe1": true,
		"trimGe1":     true,
	},
}

// ── ValidValues (C20) ──

var srcC20ValidValuesTestCase = coretestcases.CaseV1{
	Title: "ValidValues basics UsingValues Strings Find SafeValidValueAt SafeIndexes ConcatNew Hashmap AddHashset return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── ValueStatus (C20) ──

var srcC20ValueStatusTestCase = coretestcases.CaseV1{
	Title: "ValueStatus constructors Clone return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── TextWithLineNumber (C20) ──

var srcC20TextWithLineNumberTestCase = coretestcases.CaseV1{
	Title: "TextWithLineNumber HasLineNumber IsInvalidLineNumber Length IsEmpty IsEmptyText return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── LeftRight (C20) ──

var srcC20LeftRightTestCase = coretestcases.CaseV1{
	Title: "LeftRight constructors UsingSlice UsingSlicePtr TrimmedUsingSlice Methods IsEqual Clone Dispose return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── LeftMiddleRight (C20) ──

var srcC20LeftMiddleRightTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRight constructors Methods Clone Dispose return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}
