package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── ValidValue ──

// Branch: ValidValue factory methods
var srcC04ValidValueFactoriesTestCase = coretestcases.CaseV1{
	Title: "ValidValue NewValidValue returns valid state -- string input",
	ExpectedInput: args.Map{
		"value":   "hello",
		"isValid": true,
	},
}

var srcC04ValidValueEmptyFactoryTestCase = coretestcases.CaseV1{
	Title: "ValidValue NewValidValueEmpty returns valid empty -- no input",
	ExpectedInput: args.Map{
		"value":   "",
		"isValid": true,
	},
}

var srcC04ValidValueInvalidTestCase = coretestcases.CaseV1{
	Title: "ValidValue InvalidValidValue returns invalid -- with message",
	ExpectedInput: args.Map{
		"isValid":           false,
		"invalidNoMsgValid": false,
	},
}

// Branch: ValidValue methods
var srcC04ValidValueMethodsTestCase = coretestcases.CaseV1{
	Title: "ValidValue methods return correct state -- hello input",
	ExpectedInput: args.Map{
		"isEmpty":               false,
		"hasValidNonEmpty":      true,
		"hasSafeNonEmpty":       true,
		"isWhitespace":          false,
		"hasValidNonWhitespace": true,
		"trim":                  "hello",
		"isHello":               true,
		"isWorld":               false,
		"isAnyOfHelloWorld":     true,
		"isAnyOfEmpty":          true,
		"isAnyOfX":              false,
		"isContainsHel":         true,
		"isAnyContainsHel":      true,
		"isAnyContainsEmpty":    true,
		"isAnyContainsXyz":      false,
		"isEqualNonSensitive":   true,
	},
}

// Branch: ValidValue conversions
var srcC04ValidValueConversionsTestCase = coretestcases.CaseV1{
	Title: "ValidValue ValueInt ValueByte ValueFloat64 return correct -- 42",
	ExpectedInput: args.Map{
		"valueInt":     42,
		"valueDefInt":  42,
		"valueByte":    42,
		"valueDefByte": 42,
		"trueBool":     true,
	},
}

var srcC04ValidValueBadConversionsTestCase = coretestcases.CaseV1{
	Title: "ValidValue bad conversions return defaults -- abc input",
	ExpectedInput: args.Map{
		"valueInt":    99,
		"valueDefInt": 0,
		"valueFloat":  1.0,
	},
}

// Branch: ValidValue BytesOnce
var srcC04ValidValueBytesOnceTestCase = coretestcases.CaseV1{
	Title: "ValidValue ValueBytesOnce returns cached bytes -- hello input",
	ExpectedInput: args.Map{
		"len1": 5,
		"len2": 5,
	},
}

// Branch: ValidValue Regex
var srcC04ValidValueRegexTestCase = coretestcases.CaseV1{
	Title: "ValidValue Regex methods return safe defaults -- nil regexp",
	ExpectedInput: args.Map{
		"isMatch":    false,
		"findStr":    "",
		"hasAll":     false,
		"allLen":     0,
		"findAllLen": 0,
	},
}

// Branch: ValidValue Split
var srcC04ValidValueSplitTestCase = coretestcases.CaseV1{
	Title: "ValidValue Split returns 3 -- comma separated",
	ExpectedInput: args.Map{
		"splitLen": 3,
	},
}

// Branch: ValidValue Clone
var srcC04ValidValueCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "ValidValue Clone returns copy -- valid input",
		ArrangeInput: args.Map{
			"isNil": false,
		},
		ExpectedInput: args.Map{
			"value": "hello",
		},
	},
	{
		Title: "ValidValue Clone returns nil -- nil receiver",
		ArrangeInput: args.Map{
			"isNil": true,
		},
		ExpectedInput: args.Map{
			"isNilResult": true,
		},
	},
}

// Branch: ValidValue String/FullString
var srcC04ValidValueStringTestCase = coretestcases.CaseV1{
	Title: "ValidValue String FullString return values -- hello and nil",
	ExpectedInput: args.Map{
		"string":          "hello",
		"fullStringEmpty": false,
		"nilString":       "",
		"nilFullString":   "",
	},
}

// Branch: ValidValue Clear/Dispose
var srcC04ValidValueClearDisposeTestCase = coretestcases.CaseV1{
	Title: "ValidValue Clear Dispose execute safely -- including nil",
	ExpectedInput: args.Map{
		"clearedValue": "",
		"noPanic":      true,
	},
}

// Branch: ValidValue Json/Serialize
var srcC04ValidValueJsonTestCase = coretestcases.CaseV1{
	Title: "ValidValue Json JsonPtr Serialize execute without panic -- valid input",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── ValidValues ──

var srcC04ValidValuesBasicTestCase = coretestcases.CaseV1{
	Title: "ValidValues basic operations return correct state -- add two items",
	ExpectedInput: args.Map{
		"emptyIsEmpty": true,
		"emptyHasAny":  false,
		"emptyLength":  0,
		"addedLength":  2,
	},
}

var srcC04ValidValuesNilTestCase = coretestcases.CaseV1{
	Title: "ValidValues nil receiver returns 0 length -- nil pointer",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var srcC04ValidValuesFactoriesTestCase = coretestcases.CaseV1{
	Title: "ValidValues factories return correct lengths -- UsingValues and Cap",
	ExpectedInput: args.Map{
		"usingValuesLen": 1,
		"emptyLen":       0,
		"capLen":         0,
	},
}

var srcC04ValidValuesSafeValuesTestCase = coretestcases.CaseV1{
	Title: "ValidValues SafeValueAt returns correct values -- two items",
	ExpectedInput: args.Map{
		"safeAt0":          "a",
		"safeAt99":         "",
		"safeValidAt0":     "a",
		"safeValidAt99":    "",
		"safeValuesLen":    2,
		"safeValidValsLen": 1,
	},
}

var srcC04ValidValuesStringsTestCase = coretestcases.CaseV1{
	Title: "ValidValues Strings FullStrings String return non-empty -- one item",
	ExpectedInput: args.Map{
		"stringsLen":     1,
		"fullStringsLen": 1,
		"stringNonEmpty": true,
	},
}

// ── ValueStatus ──

var srcC04ValueStatusTestCase = coretestcases.CaseV1{
	Title: "ValueStatus factories and clone return correct state -- invalid with message",
	ExpectedInput: args.Map{
		"invalidNoMsgValid": false,
		"cloneMessage":      "msg",
	},
}

// ── TextWithLineNumber ──

var srcC04TextWithLineNumberTestCase = coretestcases.CaseV1{
	Title: "TextWithLineNumber returns correct state -- line 1 hello",
	ExpectedInput: args.Map{
		"hasLineNumber":    true,
		"isInvalidLine":    false,
		"length":           5,
		"isEmpty":          false,
		"isEmptyText":      false,
		"isEmptyBoth":      false,
		"nilHasLine":       false,
		"nilIsInvalidLine": true,
		"nilLength":        0,
		"nilIsEmpty":       true,
		"nilIsEmptyText":   true,
	},
}

// ── KeyValuePair ──

var srcC04KeyValuePairTestCase = coretestcases.CaseV1{
	Title: "KeyValuePair methods return correct state -- k v pair",
	ExpectedInput: args.Map{
		"keyName":          "k",
		"variableName":     "k",
		"valueString":      "v",
		"isVarNameEqual":   true,
		"isValueEqual":     true,
		"compileNonEmpty":  true,
		"stringNonEmpty":   true,
		"isKeyEmpty":       false,
		"isValueEmpty":     false,
		"isKeyValueEmpty":  false,
		"hasKey":           true,
		"hasValue":         true,
		"trimKey":          "k",
		"trimValue":        "v",
		"isKV":             true,
		"isKey":            true,
		"isVal":            true,
		"isKeyValAnyEmpty": false,
	},
}

var srcC04KeyValuePairConversionsTestCase = coretestcases.CaseV1{
	Title: "KeyValuePair value conversions return correct -- 42",
	ExpectedInput: args.Map{
		"valueInt":    42,
		"valueDefInt": 42,
	},
}

// ── LeftRight ──

var srcC04LeftRightFactoriesTestCase = coretestcases.CaseV1{
	Title: "LeftRight factories return correct state -- a b pair",
	ExpectedInput: args.Map{
		"left":            "a",
		"right":           "b",
		"isValid":         true,
		"invalidIsValid":  false,
		"invalid2IsValid": false,
	},
}

var srcC04LeftRightMethodsTestCase = coretestcases.CaseV1{
	Title: "LeftRight methods return correct state -- hello world pair",
	ExpectedInput: args.Map{
		"leftBytes":             "hello",
		"rightBytes":            "world",
		"leftTrim":              "hello",
		"rightTrim":             "world",
		"isLeftEmpty":           false,
		"isRightEmpty":          false,
		"isLeftWhitespace":      false,
		"isRightWhitespace":     false,
		"hasValidNonEmptyLeft":  true,
		"hasValidNonEmptyRight": true,
		"hasValidNonWsLeft":     true,
		"hasValidNonWsRight":    true,
		"hasSafeNonEmpty":       true,
		"isLeft":                true,
		"isRight":               true,
		"is":                    true,
		"isLeftRegexNil":        false,
		"isRightRegexNil":       false,
		"cloneLeft":             "hello",
		"isEqual":               true,
	},
}

var srcC04LeftRightFromSliceTestCase = coretestcases.CaseV1{
	Title: "LeftRight FromSlice returns correct state -- various inputs",
	ExpectedInput: args.Map{
		"twoLeft":    "a",
		"twoRight":   "b",
		"oneLeft":    "a",
		"oneRight":   "",
		"nilIsValid": false,
		"trimLeft":   "a",
		"trimRight":  "b",
	},
}

var srcC04LeftRightFromSplitTestCase = coretestcases.CaseV1{
	Title: "LeftRight FromSplit returns correct state -- key=val",
	ExpectedInput: args.Map{
		"splitLeft":  "key",
		"splitRight": "val",
		"trimLeft":   "key",
		"trimRight":  "val",
		"fullLeft":   "a",
		"fullRight":  "b:c",
	},
}

// ── LeftMiddleRight ──

var srcC04LeftMiddleRightTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRight methods return correct state -- a b c",
	ExpectedInput: args.Map{
		"left":            "a",
		"middle":          "b",
		"right":           "c",
		"isLeftEmpty":     false,
		"isMiddleEmpty":   false,
		"isRightEmpty":    false,
		"hasSafeNonEmpty": true,
		"isAll":           true,
		"is":              true,
		"cloneLeft":       "a",
		"toLrLeft":        "a",
		"toLrRight":       "c",
		"invalidIsValid":  false,
	},
}

var srcC04LeftMiddleRightFromSplitTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRight FromSplit returns correct -- a.b.c",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c",
		"nLeft":   "a",
		"nMiddle": "b",
		"nRight":  "c:d",
	},
}

// ── KeyAnyValuePair ──

var srcC04KeyAnyValuePairTestCase = coretestcases.CaseV1{
	Title: "KeyAnyValuePair methods return correct state -- k 42",
	ExpectedInput: args.Map{
		"keyName":        "k",
		"variableName":   "k",
		"valueAny":       42,
		"isVarNameEqual": true,
		"isValueNull":    false,
		"hasNonNull":     true,
		"hasValue":       true,
		"isEmptyStr":     false,
		"isWhitespace":   false,
		"compileEmpty":   false,
		"stringEmpty":    false,
	},
}

// ── HashmapDiff ──

var srcC04HashmapDiffTestCase = coretestcases.CaseV1{
	Title: "HashmapDiff methods return correct state -- one item",
	ExpectedInput: args.Map{
		"isEmpty":    false,
		"hasAny":     true,
		"length":     1,
		"lastIndex":  0,
		"hasChanges": true,
		"isRawEqual": false,
		"nilLength":  0,
	},
}
