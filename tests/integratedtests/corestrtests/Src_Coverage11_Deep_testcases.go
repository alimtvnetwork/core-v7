package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── Collection extended (C11) ──

var srcC11CollectionJsonStringTestCase = coretestcases.CaseV1{
	Title: "Collection JsonString JsonStringMust return non-empty -- a b",
	ExpectedInput: args.Map{
		"jsonString":     true,
		"jsonStringMust": true,
	},
}

var srcC11CollectionStateTestCase = coretestcases.CaseV1{
	Title: "Collection HasAnyItem LastIndex HasIndex Count LengthLock return correct -- a b",
	ExpectedInput: args.Map{
		"hasAnyItem": true,
		"lastIndex":  1,
		"hasIdx1":    true,
		"hasIdx5":    false,
		"hasIdxNeg":  false,
		"count":      1,
		"lengthLock": 1,
	},
}

var srcC11CollectionListMethodsTestCase = coretestcases.CaseV1{
	Title: "Collection ListStringsPtr ListStrings StringJSON return correct -- a",
	ExpectedInput: args.Map{
		"listStrPtrLen": 1,
		"listStrLen":    1,
		"stringJSON":    true,
	},
}

var srcC11CollectionRemoveAtTestCase = coretestcases.CaseV1{
	Title: "Collection RemoveAt returns correct -- remove index 1 of 3",
	ExpectedInput: args.Map{
		"ok":        true,
		"newLength": 2,
		"fail100":   false,
		"failNeg":   false,
	},
}

var srcC11CollectionCapacityTestCase = coretestcases.CaseV1{
	Title: "Collection Capacity returns correct -- Cap(10) and empty",
	ExpectedInput: args.Map{
		"capGt0":  true,
		"emptyCap": 0,
	},
}

var srcC11CollectionEqualsTestCase = coretestcases.CaseV1{
	Title: "Collection IsEquals IsEqualsWithSensitive return correct -- various",
	ExpectedInput: args.Map{
		"equalSame":      true,
		"equalDiff":      false,
		"equalBothEmpty": true,
		"equalDiffLen":   false,
		"insensitive":    true,
		"isEmptyLock":    true,
	},
}

var srcC11CollectionHasItemsTestCase = coretestcases.CaseV1{
	Title: "Collection HasItems returns correct -- items and nil",
	ExpectedInput: args.Map{
		"hasItems":    true,
		"nilHasItems": false,
	},
}

var srcC11CollectionAddMethodsTestCase = coretestcases.CaseV1{
	Title: "Collection Add variants return correct lengths -- various",
	ExpectedInput: args.Map{
		"addLockLen":     1,
		"addNonEmpty":    1,
		"addNonEmptyWS": 1,
		"addErrorNil":   0,
		"addIfMany":     2,
		"addFunc":       1,
		"addsLock":      2,
		"addStrings":    2,
		"addCol":        1,
		"addColEmpty":   0,
		"addCols":       2,
	},
}

var srcC11CollectionErrorTestCase = coretestcases.CaseV1{
	Title: "Collection AsDefaultError AsError ToError return correct -- err items",
	ExpectedInput: args.Map{
		"asDefaultErr": true,
		"asErrorNil":   true,
		"noPanic":      true,
	},
}

var srcC11CollectionEachSplitConcatTestCase = coretestcases.CaseV1{
	Title: "Collection EachItemSplitBy ConcatNew return correct -- various",
	ExpectedInput: args.Map{
		"splitLen":      4,
		"concatLen":     3,
		"concatEmptyLen": 1,
	},
}

// ── Hashmap extended (C11) ──

var srcC11HashmapMethodsTestCase = coretestcases.CaseV1{
	Title: "Hashmap extended methods return correct -- KeyValues factory",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC11HashmapSetTestCase = coretestcases.CaseV1{
	Title: "Hashmap Set SetTrim SetBySplitter return correct -- various",
	ExpectedInput: args.Map{
		"setHas":       true,
		"setTrimHas":   true,
		"splitVal":     "value",
		"noSplitVal":   "",
	},
}

var srcC11HashmapAddVariantsTestCase = coretestcases.CaseV1{
	Title: "Hashmap Add variants return correct -- map kvp collection hashmap",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC11HashmapConcatTestCase = coretestcases.CaseV1{
	Title: "Hashmap ConcatNew ConcatNewUsingMaps return correct -- merge",
	ExpectedInput: args.Map{
		"concatGe2":    true,
		"concatEGe1":   true,
		"concatMapGe2": true,
		"concatMEGe1":  true,
	},
}

// ── Hashset (C11) ──

var srcC11HashsetHasItemsTestCase = coretestcases.CaseV1{
	Title: "Hashset HasItems returns correct -- items and nil",
	ExpectedInput: args.Map{
		"hasItems":    true,
		"nilHasItems": false,
	},
}

// ── SimpleSlice (C11) ──

var srcC11SimpleSliceHasItemsTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice HasAnyItem returns true -- one item",
	ExpectedInput: args.Map{
		"hasAnyItem": true,
	},
}

// ── Types (C11) ──

var srcC11ValidValueTestCase = coretestcases.CaseV1{
	Title: "ValidValue NewValidValue NewInvalidValue return correct -- hello",
	ExpectedInput: args.Map{
		"validIsValid":   true,
		"validValue":     "hello",
		"invalidIsValid": false,
	},
}

var srcC11LeftRightTestCase = coretestcases.CaseV1{
	Title: "LeftRight LeftMiddleRight FromSplit return correct -- key=val a:b:c",
	ExpectedInput: args.Map{
		"lrLeft":     "key",
		"lrRight":    "val",
		"noSplitL":   "nosplit",
		"lmrLeft":    "a",
		"lmrMiddle":  "b",
		"lmrRight":   "c",
	},
}

// ── CollectionsOfCollection (C11) ──

var srcC11CocTestCase = coretestcases.CaseV1{
	Title: "CollectionsOfCollection methods return correct -- various factory and add",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── CloneSlice (C11) ──

var srcC11CloneSliceTestCase = coretestcases.CaseV1{
	Title: "CloneSlice CloneSliceIf return correct -- deep clone and conditional",
	ExpectedInput: args.Map{
		"deepClone":     "a",
		"nilLen":        0,
		"cloneIfTrue":   2,
		"cloneIfFalse":  1,
	},
}

// ── KeyAnyValuePair (C11) ──

var srcC11KeyAnyValuePairTestCase = coretestcases.CaseV1{
	Title: "KeyAnyValuePair ValueString returns non-empty -- int 42",
	ExpectedInput: args.Map{
		"nonEmpty": true,
	},
}

// ── AllIndividualStringsOfStringsLength (C11) ──

var srcC11AllIndStrLenTestCase = coretestcases.CaseV1{
	Title: "AllIndividualStringsOfStringsLength returns correct -- 2 slices and nil",
	ExpectedInput: args.Map{
		"length":  3,
		"nilLen":  0,
	},
}
