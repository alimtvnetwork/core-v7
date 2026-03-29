package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── SimpleStringOnce ──

var srcC09SimpleStringOnceCoreTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce Core returns correct state -- init hello",
	ExpectedInput: args.Map{
		"value":             "hello",
		"isInitialized":     true,
		"isDefined":         true,
		"isUninitialized":   false,
		"isInvalid":         false,
		"safeValue":         "hello",
		"isEmpty":           false,
		"isWhitespace":      false,
		"trim":              "hello",
		"hasValidNonEmpty":  true,
		"hasValidNonWS":     true,
		"hasSafeNonEmpty":   true,
		"isHello":           true,
		"isWorld":           false,
		"isAnyOfHello":      true,
		"isAnyOfX":          false,
		"isContainsHel":     true,
		"isAnyContainsHel":  true,
		"isEqualNonSens":    true,
	},
}

var srcC09SimpleStringOnceSetTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce SetOnUninitialized returns correct -- set once",
	ExpectedInput: args.Map{
		"firstErr":  true,
		"secondErr": false,
	},
}

var srcC09SimpleStringOnceGetSetOnceTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce GetSetOnce returns first value -- set twice",
	ExpectedInput: args.Map{
		"first":  "first",
		"second": "first",
	},
}

var srcC09SimpleStringOnceGetOnceTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce GetOnce returns empty -- uninitialized",
	ExpectedInput: args.Map{
		"value": "",
	},
}

var srcC09SimpleStringOnceGetOnceFuncTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce GetOnceFunc returns computed then cached -- two calls",
	ExpectedInput: args.Map{
		"first":  "computed",
		"second": "computed",
	},
}

var srcC09SimpleStringOnceSetOnceIfTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce SetOnceIfUninitialized returns correct -- set twice",
	ExpectedInput: args.Map{
		"first":  true,
		"second": false,
	},
}

var srcC09SimpleStringOnceInvalidateTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce Invalidate Reset clear state -- init then invalidate",
	ExpectedInput: args.Map{
		"afterInvalidate": false,
		"afterReset":      false,
	},
}

var srcC09SimpleStringOnceConversionsTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce Conversions return correct -- various types",
	ExpectedInput: args.Map{
		"int":      42,
		"valInt":   42,
		"defInt":   42,
		"byte":     byte(42),
		"valByte":  byte(42),
		"defByte":  byte(42),
		"float64":  true,
		"defFloat": true,
		"boolTrue": true,
		"boolDef":  true,
		"isValBool": true,
		"boolYes":  true,
		"intAbc":   0,
		"noPanic":  true,
	},
}

var srcC09SimpleStringOnceWithinRangeTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce WithinRange returns correct -- 50 in 0-100",
	ExpectedInput: args.Map{
		"value":   50,
		"ok":      true,
		"defVal":  50,
		"defOk":   true,
		"noPanic": true,
	},
}

var srcC09SimpleStringOnceConcatTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce ConcatNew returns concatenated -- hello world",
	ExpectedInput: args.Map{
		"value":   "hello world",
		"noPanic": true,
	},
}

var srcC09SimpleStringOnceSplitTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce Split methods execute without panic -- a,b,c",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC09SimpleStringOnceVariousTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce Various methods execute without panic -- hello",
	ExpectedInput: args.Map{
		"noPanic":        true,
		"nilString":      "",
		"nilStringPtr":   true,
		"nilClonePtr":    true,
	},
}

var srcC09SimpleStringOnceJsonTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce Json methods execute without panic -- hello",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── KeyValueCollection ──

var srcC09KeyValueCollectionCoreTestCase = coretestcases.CaseV1{
	Title: "KeyValueCollection Core returns correct state -- add k1 k2",
	ExpectedInput: args.Map{
		"length":    2,
		"count":     2,
		"lastIndex": 1,
		"hasIdx0":   true,
		"hasIdx5":   false,
		"hasKey":    true,
		"contains":  true,
		"getValue":  "v1",
		"found":     true,
		"noPanic":   true,
	},
}

var srcC09KeyValueCollectionAddTestCase = coretestcases.CaseV1{
	Title: "KeyValueCollection Add variants execute without panic -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC09KeyValueCollectionFindTestCase = coretestcases.CaseV1{
	Title: "KeyValueCollection Find returns correct -- find key a",
	ExpectedInput: args.Map{
		"resultLen": 1,
	},
}

var srcC09KeyValueCollectionSafeTestCase = coretestcases.CaseV1{
	Title: "KeyValueCollection Safe methods execute without panic -- one item",
	ExpectedInput: args.Map{
		"safeVal0":  "1",
		"safeVal99": "",
		"noPanic":   true,
	},
}

var srcC09KeyValueCollectionJsonTestCase = coretestcases.CaseV1{
	Title: "KeyValueCollection Json methods execute without panic -- one item",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC09KeyValueCollectionClearTestCase = coretestcases.CaseV1{
	Title: "KeyValueCollection Clear Dispose execute without panic -- nil safe",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── NonChainedLinkedListNodes ──

var srcC09NonChainedLLNodesTestCase = coretestcases.CaseV1{
	Title: "NonChainedLinkedListNodes returns correct state -- 2 nodes",
	ExpectedInput: args.Map{
		"length":          2,
		"firstElement":    "a",
		"lastElement":     "b",
		"chainingBefore":  false,
		"chainingAfter":   true,
		"noPanic":         true,
	},
}

// ── NonChainedLinkedCollectionNodes ──

var srcC09NonChainedLCNodesTestCase = coretestcases.CaseV1{
	Title: "NonChainedLinkedCollectionNodes returns correct state -- 2 nodes",
	ExpectedInput: args.Map{
		"length":        2,
		"chainingAfter": true,
		"noPanic":       true,
	},
}

// ── CollectionsOfCollection ──

var srcC09CollectionsOfCollectionTestCase = coretestcases.CaseV1{
	Title: "CollectionsOfCollection methods return correct -- 2 collections",
	ExpectedInput: args.Map{
		"length":     2,
		"allItemLen": 3,
		"listLen":    3,
		"noPanic":    true,
	},
}

// ── HashsetsCollection ──

var srcC09HashsetsCollectionTestCase = coretestcases.CaseV1{
	Title: "HashsetsCollection methods execute without panic -- multiple adds",
	ExpectedInput: args.Map{
		"noPanic":   true,
		"lengthGe3": true,
	},
}

var srcC09HashsetsCollectionHasAllTestCase = coretestcases.CaseV1{
	Title: "HashsetsCollection HasAll returns correct -- a b present",
	ExpectedInput: args.Map{
		"hasAB":    true,
		"hasZ":     false,
		"emptyHas": false,
	},
}

var srcC09HashsetsCollectionConcatTestCase = coretestcases.CaseV1{
	Title: "HashsetsCollection ConcatNew AddHashsetsCollection execute without panic -- merge",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── CharCollectionMap ──

var srcC09CharCollectionMapMethodsTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap methods execute without panic -- abc adef bcd",
	ExpectedInput: args.Map{
		"isEmpty":  false,
		"noPanic":  true,
		"getCharE": byte(0),
		"getCharA": byte('a'),
	},
}

var srcC09CharCollectionMapAddTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap Add variants execute without panic -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC09CharCollectionMapClearTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap Clear Dispose execute without panic -- nil safe",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── CharHashsetMap ──

var srcC09CharHashsetMapMethodsTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap methods execute without panic -- abc adef bcd",
	ExpectedInput: args.Map{
		"isEmpty": false,
		"noPanic": true,
	},
}

var srcC09CharHashsetMapAddTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap Add variants execute without panic -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC09CharHashsetMapClearTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap Clear RemoveAll execute without panic -- one item",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── LinkedCollections ──

var srcC09LinkedCollectionsBasicTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections Basic methods execute without panic -- add collection",
	ExpectedInput: args.Map{
		"length":  1,
		"noPanic": true,
	},
}

var srcC09LinkedCollectionsAddTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections Add variants execute without panic -- multiple adds",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC09LinkedCollectionsLoopTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections Loop iterates items -- 2 string collections",
	ExpectedInput: args.Map{
		"countGt0": true,
	},
}

var srcC09LinkedCollectionsEqualsTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections IsEqualsPtr returns true -- same content",
	ExpectedInput: args.Map{
		"isEquals": true,
	},
}

// ── Collection remaining ──

var srcC09CollectionRemainingTestCase = coretestcases.CaseV1{
	Title: "Collection remaining methods execute without panic -- a b c",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC09CollectionFilterTestCase = coretestcases.CaseV1{
	Title: "Collection Filter returns correct count -- filter len > 1",
	ExpectedInput: args.Map{
		"filterLen": 2,
		"noPanic":   true,
	},
}

var srcC09CollectionAppendAnysTestCase = coretestcases.CaseV1{
	Title: "Collection AppendAnys AddsNonEmpty execute without panic -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── SimpleSlice remaining ──

var srcC09SimpleSliceRemainingTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice remaining methods execute without panic -- a b c",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC09SimpleSliceIsEqualTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice IsEqual variants return true -- same content",
	ExpectedInput: args.Map{
		"isEqual":               true,
		"isEqualLines":          true,
		"isEqualUnordered":      true,
		"isEqualUnorderedClone": true,
		"isDistinctEqual":       true,
		"isDistinctEqualRaw":    true,
		"isUnorderedEqual":      true,
		"isUnorderedEqualRaw":   true,
	},
}

var srcC09SimpleSliceDistinctDiffTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice DistinctDiff AddedRemovedLinesDiff execute without panic -- a b diff b c",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC09SimpleSliceRemoveIndexesTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice RemoveIndexes returns correct -- remove index 1",
	ExpectedInput: args.Map{
		"newLen":   2,
		"noErr":    true,
		"emptyErr": true,
	},
}

var srcC09SimpleSliceCloneTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice Clone methods execute without panic -- a b",
	ExpectedInput: args.Map{
		"noPanic":      true,
		"nilClonePtr":  true,
	},
}

var srcC09SimpleSliceClearDisposeTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice Clear Dispose execute without panic -- nil safe",
	ExpectedInput: args.Map{
		"noPanic":  true,
		"nilClear": true,
	},
}
