package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── Collection extended ──

var srcC10CollectionTakeSkipTestCase = coretestcases.CaseV1{
	Title: "Collection Take Skip Limit return correct lengths -- 5 items",
	ExpectedInput: args.Map{
		"takeLen":     2,
		"skipLen":     3,
		"limitLen":    2,
		"limitAllLen": 5,
	},
}

var srcC10CollectionAddNonEmptyTestCase = coretestcases.CaseV1{
	Title: "Collection AddNonEmptyStrings AddNonEmptyStringsSlice return correct -- filter empty",
	ExpectedInput: args.Map{
		"addNonEmptyLen":      2,
		"addNonEmptySliceLen": 2,
	},
}

var srcC10CollectionNonEmptyListTestCase = coretestcases.CaseV1{
	Title: "Collection NonEmptyList NonEmptyListPtr return correct -- a empty b",
	ExpectedInput: args.Map{
		"listLen":    2,
		"listPtrLen": 2,
	},
}

var srcC10CollectionMethodsNoPanicTestCase = coretestcases.CaseV1{
	Title: "Collection various methods execute without panic -- items list copy filter hashset",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC10CollectionHasTestCase = coretestcases.CaseV1{
	Title: "Collection Has HasLock HasAll HasUsingSensitivity return correct -- a b c",
	ExpectedInput: args.Map{
		"has":           true,
		"hasMissing":    false,
		"hasLock":       true,
		"hasAll":        true,
		"hasAllMissing": false,
		"hasSensLower":  true,
		"hasSensExact":  true,
	},
}

var srcC10CollectionSortedTestCase = coretestcases.CaseV1{
	Title: "Collection Sorted methods return correct order -- c a b",
	ExpectedInput: args.Map{
		"ascFirst":  "a",
		"ascCFirst": "a",
		"ascLFirst": "a",
		"dscFirst":  "c",
	},
}

var srcC10CollectionFilterTestCase = coretestcases.CaseV1{
	Title: "Collection Filter FilterPtr FilteredCollection return correct -- len > 1",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC10CollectionContainsTestCase = coretestcases.CaseV1{
	Title: "Collection IsContainsPtr IsContainsAll IsContainsAllSlice return correct -- a b c",
	ExpectedInput: args.Map{
		"containsPtr":  true,
		"containsAll":  true,
		"containsLock": true,
		"containsSlc":  true,
		"hashsetAll":   true,
	},
}

var srcC10CollectionNewExpandMergeTestCase = coretestcases.CaseV1{
	Title: "Collection New ExpandSlicePlusAdd MergeSlicesOfSlice return correct -- various",
	ExpectedInput: args.Map{
		"newLen":    2,
		"expandLen": 3,
		"mergeLen":  3,
	},
}

var srcC10CollectionExceptTestCase = coretestcases.CaseV1{
	Title: "Collection GetAllExceptCollection GetAllExcept return correct -- exclude b",
	ExpectedInput: args.Map{
		"exceptColLen": 2,
		"exceptLen":    2,
	},
}

var srcC10CollectionJoinsTestCase = coretestcases.CaseV1{
	Title: "Collection Joins NonEmptyJoins NonWhitespaceJoins return non-empty -- various",
	ExpectedInput: args.Map{
		"joins":         true,
		"nonEmptyJoins": true,
		"nonWSJoins":    true,
	},
}

var srcC10CollectionStringTestCase = coretestcases.CaseV1{
	Title: "Collection String StringLock SummaryString return non-empty -- one item",
	ExpectedInput: args.Map{
		"string":     true,
		"stringLock": true,
		"summary":    true,
		"summaryH":   true,
		"noPanic":    true,
	},
}

var srcC10CollectionJsonTestCase = coretestcases.CaseV1{
	Title: "Collection Json Serialize Deserialize execute without panic -- one item",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC10CollectionClearDisposeTestCase = coretestcases.CaseV1{
	Title: "Collection Clear Dispose return correct -- a b",
	ExpectedInput: args.Map{
		"clearLen": 0,
		"noPanic":  true,
	},
}

var srcC10CollectionAddFuncTestCase = coretestcases.CaseV1{
	Title: "Collection AddFuncResult AddStringsByFuncChecking return correct -- various",
	ExpectedInput: args.Map{
		"funcLen":   1,
		"filterLen": 2,
	},
}

var srcC10CollectionParseInjectTestCase = coretestcases.CaseV1{
	Title: "Collection ParseInjectUsingJson JsonParseSelfInject execute without panic -- round trip",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── Hashmap extended ──

var srcC10HashmapBasicTestCase = coretestcases.CaseV1{
	Title: "Hashmap basic methods return correct -- add a=1",
	ExpectedInput: args.Map{
		"hasItems":    true,
		"colNotNil":   true,
		"isEmptyLock": true,
		"addLockLen":  1,
	},
}

var srcC10HashmapContainsTestCase = coretestcases.CaseV1{
	Title: "Hashmap Contains Has IsKeyMissing return correct -- key a",
	ExpectedInput: args.Map{
		"contains":       true,
		"containsLock":   true,
		"notMissing":     false,
		"notMissingLock": false,
		"hasLock":        true,
		"hasAllStr":      true,
		"hasAllStrFail":  false,
		"hasAll":         true,
		"hasAny":         true,
		"hasWithLock":    true,
	},
}

var srcC10HashmapKeysValsTestCase = coretestcases.CaseV1{
	Title: "Hashmap Keys Values Pairs return correct lengths -- one item",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC10HashmapRemoveTestCase = coretestcases.CaseV1{
	Title: "Hashmap Remove RemoveWithLock return correct -- remove a",
	ExpectedInput: args.Map{
		"removeLen":     0,
		"removeLockLen": 0,
	},
}

var srcC10HashmapStringTestCase = coretestcases.CaseV1{
	Title: "Hashmap String StringLock return non-empty -- one item",
	ExpectedInput: args.Map{
		"string":     true,
		"stringLock": true,
	},
}

var srcC10HashmapItemsCopyTestCase = coretestcases.CaseV1{
	Title: "Hashmap ItemsCopyLock SafeItems ValuesListCopyLock execute without panic -- one item",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC10HashmapMutateTestCase = coretestcases.CaseV1{
	Title: "Hashmap ValuesToLower KeysToLower execute without panic -- mutation",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC10HashmapEqualCloneTestCase = coretestcases.CaseV1{
	Title: "Hashmap IsEqual Clone return correct -- same map",
	ExpectedInput: args.Map{
		"isEqual":     true,
		"isEqualLock": true,
		"cloneLen":    1,
	},
}

var srcC10HashmapGetValueTestCase = coretestcases.CaseV1{
	Title: "Hashmap GetValue Join JoinKeys return correct -- a=1",
	ExpectedInput: args.Map{
		"value":        "1",
		"found":        true,
		"joinNonEmpty": true,
		"keysNonEmpty": true,
	},
}

var srcC10HashmapDisposeTestCase = coretestcases.CaseV1{
	Title: "Hashmap Dispose ToError ToDefaultError execute without panic -- one item",
	ExpectedInput: args.Map{
		"noPanic":  true,
		"hasError": true,
		"hasDefErr": true,
	},
}

var srcC10HashmapJsonTestCase = coretestcases.CaseV1{
	Title: "Hashmap Json Serialize Deserialize ParseInject execute without panic -- round trip",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC10HashmapAddVariantsTestCase = coretestcases.CaseV1{
	Title: "Hashmap Add variants execute without panic -- hashmap map collection kvp wg",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC10HashmapSetTestCase = coretestcases.CaseV1{
	Title: "Hashmap Set SetTrim SetBySplitter return correct -- various",
	ExpectedInput: args.Map{
		"setLen":      1,
		"trimVal":     "1",
		"splitVal":    "value",
	},
}

var srcC10HashmapDiffTestCase = coretestcases.CaseV1{
	Title: "Hashmap DiffRaw Diff execute without panic -- compare maps",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC10HashmapFilterExceptTestCase = coretestcases.CaseV1{
	Title: "Hashmap HasAllCollectionItems Filter GetAllExcept return correct -- a b",
	ExpectedInput: args.Map{
		"hasAllCol":   true,
		"filterLen":   1,
		"filterColNN": true,
		"exceptVals":  1,
		"exceptKeys":  1,
		"exceptCol":   true,
	},
}

var srcC10HashmapCompilerConcatTestCase = coretestcases.CaseV1{
	Title: "Hashmap ToStringsUsingCompiler ConcatNew ConcatNewUsingMaps return correct -- merge",
	ExpectedInput: args.Map{
		"compilerLen":  1,
		"concatLen":    2,
		"concatMapLen": 2,
	},
}

var srcC10HashmapFilterFuncTestCase = coretestcases.CaseV1{
	Title: "Hashmap AddsOrUpdatesUsingFilter variants execute without panic -- filter funcs",
	ExpectedInput: args.Map{
		"filterLen": 2,
		"noPanic":   true,
	},
}

var srcC10HashmapAddTypedTestCase = coretestcases.CaseV1{
	Title: "Hashmap AddOrUpdate typed variants execute without panic -- int float any kvp",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC10HashmapKeyValLinesTestCase = coretestcases.CaseV1{
	Title: "Hashmap KeyValStringLines returns correct length -- one item",
	ExpectedInput: args.Map{
		"linesLen": 1,
	},
}

// ── Hashset extended ──

var srcC10HashsetExtendedTestCase = coretestcases.CaseV1{
	Title: "Hashset extended methods execute without panic -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── newCreator paths ──

var srcC10NewCreatorTestCase = coretestcases.CaseV1{
	Title: "New creator methods return correct -- various factory methods",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── emptyCreator paths ──

var srcC10EmptyCreatorTestCase = coretestcases.CaseV1{
	Title: "Empty creator methods return non-nil -- all types",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── Standalone functions ──

var srcC10AnyToStringTestCase = coretestcases.CaseV1{
	Title: "AnyToString returns correct -- int string nil",
	ExpectedInput: args.Map{
		"intNonEmpty":    true,
		"stringVal":      "hello",
		"noPanic":        true,
	},
}

var srcC10AllIndividualStringsLenTestCase = coretestcases.CaseV1{
	Title: "AllIndividualStringsOfStringsLength returns correct -- 2 slices",
	ExpectedInput: args.Map{
		"length": 3,
	},
}

var srcC10AllIndividualsSimpleSlicesTestCase = coretestcases.CaseV1{
	Title: "AllIndividualsLengthOfSimpleSlices returns correct -- 2 slices",
	ExpectedInput: args.Map{
		"length": 3,
	},
}
