package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── LinkedList creators (C18) ──

var srcC18CreatorsTestCase = coretestcases.CaseV1{
	Title: "LinkedList creators Create Empty Strings SpreadStrings UsingMap PointerStringsPtr return correct -- various",
	ExpectedInput: args.Map{
		"createNN":    true,
		"createEmpty": true,
		"emptyLen":    0,
		"stringsLen":  3,
		"stringsEE":   true,
		"spreadLen":   2,
		"spreadEE":    true,
		"mapLen":      2,
		"mapNilE":     true,
		"ptrLen":      2,
		"ptrNilE":     true,
		"emptyLLe":    true,
	},
}

var srcC18HeadTailTestCase = coretestcases.CaseV1{
	Title: "LinkedList Head Tail LengthLock return correct -- spread strings",
	ExpectedInput: args.Map{
		"head":      "a",
		"tail":      "c",
		"lengthLock": 2,
	},
}

var srcC18StateTestCase = coretestcases.CaseV1{
	Title: "LinkedList IsEmpty HasItems IsEmptyLock return correct -- various",
	ExpectedInput: args.Map{
		"emptyIsEmpty": true,
		"emptyHasIt":   false,
		"addedIsEmpty": false,
		"addedHasIt":   true,
		"emptyLock":    true,
	},
}

var srcC18AddTestCase = coretestcases.CaseV1{
	Title: "LinkedList Add AddLock AddNonEmpty AddNonEmptyWhitespace AddIf AddsIf AddFunc AddFuncErr Push Adds AddStrings AddsLock AddItemsMap return correct -- various",
	ExpectedInput: args.Map{
		"addLen":       2,
		"addLockLen":   1,
		"nonEmptyLen":  1,
		"noWhiteLen":   1,
		"addIfLen":     1,
		"addIfHead":    "yes",
		"addsIfLen":    2,
		"funcHead":     "computed",
		"funcErrLen":   1,
		"funcErrErr":   true,
		"pushLen":      2,
		"addsLen":      3,
		"addsEmptyE":   true,
		"stringsLen":   2,
		"addsLockLen":  2,
		"mapLen":       1,
		"mapEmptyE":    true,
	},
}

var srcC18AddFrontTestCase = coretestcases.CaseV1{
	Title: "LinkedList AddFront PushFront return correct -- various",
	ExpectedInput: args.Map{
		"frontHead":    "a",
		"frontLen":     3,
		"frontEmHead":  "only",
		"frontEmLen":   1,
		"pushHead":     "a",
	},
}

var srcC18AppendNodeTestCase = coretestcases.CaseV1{
	Title: "LinkedList AppendNode AddBackNode return correct -- various",
	ExpectedInput: args.Map{
		"appendLen1":  1,
		"appendHead":  "x",
		"appendLen2":  2,
		"addBackLen":  1,
	},
}

var srcC18InsertAtTestCase = coretestcases.CaseV1{
	Title: "LinkedList InsertAt AddAfterNode return correct -- front and middle",
	ExpectedInput: args.Map{
		"frontHead":   "a",
		"middleAt1":   "b",
		"afterAt1":    "b",
	},
}

var srcC18AttachTestCase = coretestcases.CaseV1{
	Title: "LinkedList AttachWithNode return correct -- nil current, non-nil next, success",
	ExpectedInput: args.Map{
		"nilCurrErr":    true,
		"nonNilNextErr": true,
		"successLen":    2,
	},
}

var srcC18AddCollPtrTestCase = coretestcases.CaseV1{
	Title: "LinkedList AddCollectionToNode AddPointerStringsPtr AddCollection return correct -- various",
	ExpectedInput: args.Map{
		"colToNodeGe2": true,
		"ptrLen":       1,
		"colLen":       2,
		"colNilE":      true,
	},
}

var srcC18LoopTestCase = coretestcases.CaseV1{
	Title: "LinkedList Loop return correct -- full, break, empty, breakSecond",
	ExpectedInput: args.Map{
		"fullCount":    3,
		"breakCount":   1,
		"emptyOk":      true,
		"breakSecond":  2,
	},
}

var srcC18FilterTestCase = coretestcases.CaseV1{
	Title: "LinkedList Filter return correct -- keep, empty, breakFirst, breakSecond",
	ExpectedInput: args.Map{
		"keepLen":       2,
		"emptyLen":      0,
		"breakFirstLen": 1,
		"breakSecLen":   2,
	},
}

var srcC18RemoveTestCase = coretestcases.CaseV1{
	Title: "LinkedList RemoveNodeByElementValue RemoveNodeByIndex RemoveNodeByIndexes RemoveNode return correct -- various",
	ExpectedInput: args.Map{
		"rmValFirstLen":   2,
		"rmValFirstHead":  "b",
		"rmValMiddleLen":  2,
		"rmValCILen":      1,
		"rmIdxFirstHead":  "b",
		"rmIdxLastLen":    2,
		"rmIdxMiddleLen":  2,
		"rmIdxesLen":      2,
		"rmIdxesEmptyLen": 1,
		"rmNodeLen":       2,
		"rmNodeFirstHead": "b",
		"rmNodeNilLen":    1,
	},
}

var srcC18IndexAtTestCase = coretestcases.CaseV1{
	Title: "LinkedList IndexAt SafeIndexAt SafeIndexAtLock SafePointerIndexAt SafePointerIndexAtUsingDefault return correct -- various",
	ExpectedInput: args.Map{
		"idxAt0":       "a",
		"idxAt2":       "c",
		"idxNegNil":    true,
		"safeAt1":      "b",
		"safeOorNil":   true,
		"safeNegNil":   true,
		"safeLock0":    "a",
		"ptrAt0":       "a",
		"ptrOorNil":    true,
		"ptrDef0":      "a",
		"ptrDefOor":    "def",
		"ptrDefLock0":  "a",
	},
}

var srcC18NextNodesTestCase = coretestcases.CaseV1{
	Title: "LinkedList GetNextNodes GetAllLinkedNodes return correct -- various",
	ExpectedInput: args.Map{
		"nextLen": 2,
		"allLen":  2,
	},
}

var srcC18ToCollListTestCase = coretestcases.CaseV1{
	Title: "LinkedList ToCollection List ListPtr ListLock ListPtrLock return correct -- various",
	ExpectedInput: args.Map{
		"toColLen":    2,
		"toColEmptyE": true,
		"listLen":     2,
		"listFirst":   "a",
		"listEmptyLen": 0,
		"listPtrLen":  1,
		"listLockLen": 1,
		"listPtrLkLen": 1,
	},
}

var srcC18StringJoinTestCase = coretestcases.CaseV1{
	Title: "LinkedList String StringLock Join JoinLock Joins return correct -- various",
	ExpectedInput: args.Map{
		"strNonE":      true,
		"strEmptyNonE": true,
		"strLockNonE":  true,
		"strLockENonE": true,
		"join":         "a,b",
		"joinLock":     "a,b",
		"joins":        "a,b,c",
		"joinsNil":     "a",
	},
}

var srcC18CompareEqualsTestCase = coretestcases.CaseV1{
	Title: "LinkedList GetCompareSummary IsEquals IsEqualsWithSensitive return correct -- various",
	ExpectedInput: args.Map{
		"summaryNonE":  true,
		"equalSame":    true,
		"equalDiffLen": false,
		"equalSameRef": true,
		"equalBothE":   true,
		"equalOneE":    false,
	},
}

var srcC18AddStrToNodeTestCase = coretestcases.CaseV1{
	Title: "LinkedList AddStringsToNode AddStringsPtrToNode return correct -- various",
	ExpectedInput: args.Map{
		"strToNodeGe3":  true,
		"strSingleAt1":  "b",
		"strEmptyLen":   1,
		"strNilNodeLen": 1,
		"ptrToNodeGe2":  true,
		"ptrNilLen":     1,
	},
}

var srcC18JsonTestCase = coretestcases.CaseV1{
	Title: "LinkedList JsonModel JsonModelAny MarshalJSON UnmarshalJSON Json JsonPtr ParseInjectUsingJson ParseInjectUsingJsonMust JsonParseSelfInject AsJsonMarshaller return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC18ClearTestCase = coretestcases.CaseV1{
	Title: "LinkedList RemoveAll Clear return correct -- various",
	ExpectedInput: args.Map{
		"removeAllE":  true,
		"clearE":      true,
		"clearEmptyOk": true,
	},
}

// ── LinkedListNode (exported-only) ──

var srcC18NodeExportedTestCase = coretestcases.CaseV1{
	Title: "LinkedListNode IsEqual IsEqualSensitive IsEqualValue IsEqualValueSensitive AddNext AddNextNode AddStringsToNode AddStringsPtrToNode AddCollectionToNode return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// ── NonChainedLinkedListNodes ──

var srcC18NonChainedTestCase = coretestcases.CaseV1{
	Title: "NonChainedLinkedListNodes basic ApplyChaining ToChainedNodes FirstOrDefault LastOrDefault return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}
