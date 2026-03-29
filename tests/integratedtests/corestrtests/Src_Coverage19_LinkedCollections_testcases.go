package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── LinkedCollections creators (C19) ──

var srcC19CreatorsTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections creators Create Empty Strings PointerStringsPtr UsingCollections return correct -- various",
	ExpectedInput: args.Map{
		"createNN":    true,
		"createEmpty": true,
		"emptyLen":    0,
		"stringsLen":  1,
		"stringsEE":   true,
		"ptrLen":      1,
		"ptrNilE":     true,
		"usingLen":    2,
		"usingNilNN":  true,
		"emptyLCe":    true,
	},
}

var srcC19HeadTailTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections Head Tail First Last Single FirstOrDefault LastOrDefault return correct -- various",
	ExpectedInput: args.Map{
		"headNN":     true,
		"tailNN":     true,
		"firstLen":   1,
		"lastLen":    1,
		"singleLen":  1,
		"fodEmptyNN": true,
		"lodEmptyNN": true,
		"fodHasLen":  1,
	},
}

var srcC19LengthTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections LengthLock AllIndividualItemsLength return correct -- various",
	ExpectedInput: args.Map{
		"lengthLock": 1,
		"allItems":   3,
	},
}

var srcC19StateTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections IsEmpty HasItems IsEmptyLock return correct -- various",
	ExpectedInput: args.Map{
		"emptyIsEmpty": true,
		"emptyHasIt":   false,
		"emptyLock":    true,
	},
}

var srcC19AddTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections Add AddLock AddStrings AddStringsLock return correct -- various",
	ExpectedInput: args.Map{
		"addLen1":       1,
		"addLen2":       2,
		"addLockLen":    1,
		"addStrLen":     1,
		"addStrEE":      true,
		"addStrLockLen": 1,
		"addStrLockEE":  true,
	},
}

var srcC19AddFrontTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections AddFront AddFrontLock PushFront PushBack PushBackLock Push return correct -- various",
	ExpectedInput: args.Map{
		"frontFirst":    "a",
		"frontEmptyLen": 1,
		"frontLockLen":  2,
		"pushFrontLen":  2,
		"pushBackLen":   1,
		"pushBackLkLen": 1,
		"pushLen":       1,
	},
}

var srcC19AppendNodeTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections AppendNode AddBackNode return correct -- various",
	ExpectedInput: args.Map{
		"appendLen2":  2,
		"addBackLen":  1,
	},
}

var srcC19InsertAtTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections InsertAt return correct -- front and middle",
	ExpectedInput: args.Map{
		"frontFirst":  "a",
		"middleLen":   3,
	},
}

var srcC19AttachTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections AttachWithNode return correct -- nil, non-nil next, success",
	ExpectedInput: args.Map{
		"nilCurrErr":    true,
		"nonNilNextErr": true,
		"successOk":     true,
	},
}

var srcC19AddAnotherCollTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections AddAnother AddCollection AddCollectionsPtr AddCollections AppendCollections AppendCollectionsPointers return correct -- various",
	ExpectedInput: args.Map{
		"anotherLen":    2,
		"anotherNilLen": 1,
		"colLen":        1,
		"colNilE":       true,
		"colsPtrLen":    1,
		"colsLen":       1,
		"colsEmptyE":    true,
		"appendLen":     1,
		"appendNilE":    true,
		"appendPtrLen":  1,
		"appendPtrLkLen": 1,
	},
}

var srcC19LoopFilterTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections Loop Filter FilterAsCollection FilterAsCollections return correct -- various",
	ExpectedInput: args.Map{
		"loopCount":     2,
		"loopBreak":     1,
		"loopEmptyOk":   true,
		"filterLen":     2,
		"filterEmptyLen": 0,
		"filterColLen":  3,
		"filterColEE":   true,
		"filterColsLen": 2,
	},
}

var srcC19RemoveTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections GetNextNodes GetAllLinkedNodes RemoveNodeByIndex RemoveNodeByIndexes RemoveNode AddAfterNode return correct -- various",
	ExpectedInput: args.Map{
		"nextNodesLen":   2,
		"allNodesLen":    2,
		"rmIdxFirstLen":  1,
		"rmIdxLastLen":   1,
		"rmIdxMidLen":    2,
		"rmIdxesLen":     1,
		"rmIdxesEmpLen":  1,
		"rmNodeFirstLen": 1,
		"rmNodeMidLen":   2,
		"afterLen":       2,
	},
}

var srcC19ConcatIndexTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections ConcatNew IndexAt SafeIndexAt SafePointerIndexAt return correct -- various",
	ExpectedInput: args.Map{
		"concatLen":      2,
		"concatCloneLen": 1,
		"concatSameRef":  true,
		"idxAt1NN":       true,
		"idxAt0NN":       true,
		"idxNegNil":      true,
		"safeAt1NN":      true,
		"safeOorNil":     true,
		"safeNegNil":     true,
		"ptrAt0NN":       true,
		"ptrOorNil":      true,
	},
}

var srcC19ToCollStrTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections ToCollection ToCollectionSimple ToStrings ToStringsPtr ToCollectionsOfCollection ItemsOfItems SimpleSlice return correct -- various",
	ExpectedInput: args.Map{
		"toColLen":      3,
		"toColEmptyE":   true,
		"toColSimLen":   1,
		"toStrLen":      2,
		"toStrPtrLen":   1,
		"toCocNN":       true,
		"toCocEmptyNN":  true,
		"ioiLen":        2,
		"ioiEmptyLen":   0,
		"ioicLen":       1,
		"ssNN":          true,
	},
}

var srcC19AddStrOfStrAsyncTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections AddStringsOfStrings AddAsyncFuncItems AddAsyncFuncItemsPointer return correct -- various",
	ExpectedInput: args.Map{
		"sosLen":         2,
		"sosEmptyE":      true,
		"asyncLen":       1,
		"asyncEmptyE":    true,
		"asyncNilE":      true,
		"asyncPtrLen":    1,
		"asyncPtrNilE":   true,
	},
}

var srcC19StringJoinListTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections String StringLock Join Joins List ListPtr return correct -- various",
	ExpectedInput: args.Map{
		"strNonE":       true,
		"strEmptyNonE":  true,
		"strLockNonE":   true,
		"strLockENonE":  true,
		"join":          "a,b",
		"joins":         "a,b",
		"joinsNil":      "a",
		"listLen":       2,
		"listEmptyLen":  0,
		"listPtrLen":    1,
	},
}

var srcC19EqualsCompareTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections IsEqualsPtr GetCompareSummary return correct -- various",
	ExpectedInput: args.Map{
		"equalSame":     true,
		"equalNil":      false,
		"equalSameRef":  true,
		"equalBothE":    true,
		"equalOneE":     false,
		"equalDiffLen":  false,
		"summaryNonE":   true,
	},
}

var srcC19JsonTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections JSON marshal unmarshal Json JsonPtr ParseInjectUsingJson ParseInjectUsingJsonMust JsonParseSelfInject As* return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC19ClearTestCase = coretestcases.CaseV1{
	Title: "LinkedCollections RemoveAll Clear return correct -- various",
	ExpectedInput: args.Map{
		"removeAllE":  true,
		"clearE":      true,
		"clearEmptyOk": true,
	},
}

var srcC19NodeExportedTestCase = coretestcases.CaseV1{
	Title: "LinkedCollectionNode IsEmpty HasElement ListPtr Join String StringList IsEqual IsEqualValue AddNext AddNextNode AddStringsToNode AddCollectionToNode return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC19NonChainedEmptyTestCase = coretestcases.CaseV1{
	Title: "NonChainedLinkedCollectionNodes FirstOrDefault LastOrDefault return nil -- empty",
	ExpectedInput: args.Map{
		"fodNil": true,
		"lodNil": true,
	},
}
