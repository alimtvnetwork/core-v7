package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── Collection Deep (C12) ──

var srcC12CollectionAddIfFuncErrTestCase = coretestcases.CaseV1{
	Title: "Collection AddIf AddFuncErr AddError return correct -- various",
	ExpectedInput: args.Map{
		"addIfTrue":      1,
		"addIfFalse":     0,
		"funcErrOk":      1,
		"funcErrFail":    0,
		"errHandled":     true,
		"addErrVal":      "test-err",
	},
}

var srcC12CollectionHashmapMethodsTestCase = coretestcases.CaseV1{
	Title: "Collection AddHashmaps methods return correct -- keys values filter",
	ExpectedInput: args.Map{
		"hmValsLen":     1,
		"hmValsNil":     0,
		"hmValsNilNil":  0,
		"hmKeysLen":     1,
		"hmKeysNil":     0,
		"hmKVLen":       2,
		"hmKVNil":       0,
		"hmFilterLen":   1,
		"hmFilterNil":   0,
		"hmFilterBreak": 1,
	},
}

var srcC12CollectionWgLockTestCase = coretestcases.CaseV1{
	Title: "Collection AddWithWgLock returns correct -- sync add",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var srcC12CollectionIndexTestCase = coretestcases.CaseV1{
	Title: "Collection IndexAt SafeIndexAtUsingLength First Last Single return correct -- various",
	ExpectedInput: args.Map{
		"indexAt":       "b",
		"safeInRange":   "a",
		"safeOutRange":  "def",
		"first":         "x",
		"last":          "y",
		"lastOrDefE":    "",
		"lastOrDef":     "a",
		"firstOrDefE":   "",
		"firstOrDef":    "z",
		"single":        "only",
		"singlePanics":  true,
	},
}

var srcC12CollectionTakeSkipTestCase = coretestcases.CaseV1{
	Title: "Collection Take Skip return correct -- various edge cases",
	ExpectedInput: args.Map{
		"take2":       2,
		"takeMore":    1,
		"takeZero":    0,
		"skip1":       2,
		"skipZero":    1,
		"skipPanics":  true,
	},
}

var srcC12CollectionReverseTestCase = coretestcases.CaseV1{
	Title: "Collection Reverse returns correct -- 3 2 1 items",
	ExpectedInput: args.Map{
		"rev3First": "c",
		"rev3Last":  "a",
		"rev2First": "b",
		"rev1First": "a",
	},
}

var srcC12CollectionPagingTestCase = coretestcases.CaseV1{
	Title: "Collection GetPagesSize GetPagedCollection GetSinglePageCollection return correct -- various",
	ExpectedInput: args.Map{
		"pagesSize":     5,
		"pagesZero":     0,
		"pagesNeg":      0,
		"pagedLen":      4,
		"pagedSmall":    1,
		"singlePageLen": 3,
		"lastPageLen":   1,
		"smallPageLen":  1,
	},
}

var srcC12CollectionInsertRemoveTestCase = coretestcases.CaseV1{
	Title: "Collection InsertAt ChainRemoveAt RemoveItemsIndexes return correct -- various",
	ExpectedInput: args.Map{
		"insertLen":       1,
		"chainRemoveLen":  2,
		"removeIdxLen":    2,
		"removeIdxNoop":   1,
	},
}

var srcC12CollectionAppendTestCase = coretestcases.CaseV1{
	Title: "Collection AppendCollectionPtr AppendCollections return correct -- various",
	ExpectedInput: args.Map{
		"appendPtrLen":  2,
		"appendColsLen": 2,
		"appendEmpty":   0,
	},
}

var srcC12CollectionAppendAnysTestCase = coretestcases.CaseV1{
	Title: "Collection AppendAnys variants return correct -- various",
	ExpectedInput: args.Map{
		"anysLen":          2,
		"anysEmpty":        0,
		"anysLock":         1,
		"anysLockEmpty":    0,
		"filterLen":        2,
		"filterSkip":       0,
		"filterBreak":      1,
		"filterEmpty":      0,
		"filterNil":        0,
		"filterLock":       1,
		"filterLockNil":    0,
		"filterLockBreak":  1,
		"filterLockSkip":   0,
		"filterLockNilI":   0,
		"nonEmptyLen":      1,
		"nonEmptyNil":      0,
		"addsNonEmpty":     2,
		"addsNonEmptyNil":  0,
	},
}

var srcC12CollectionPtrLockTestCase = coretestcases.CaseV1{
	Title: "Collection AddsNonEmptyPtrLock return correct -- filter nil and empty",
	ExpectedInput: args.Map{
		"ptrLockLen": 1,
		"ptrLockNil": 0,
	},
}

var srcC12CollectionUniqueTestCase = coretestcases.CaseV1{
	Title: "Collection UniqueBoolMap UniqueList return correct -- deduplicate",
	ExpectedInput: args.Map{
		"boolMapLen":     2,
		"boolMapLockLen": 1,
		"uniqueLen":      2,
		"uniqueLockLen":  2,
	},
}

var srcC12CollectionFilterDeepTestCase = coretestcases.CaseV1{
	Title: "Collection Filter FilterPtr deep edge cases return correct -- break empty",
	ExpectedInput: args.Map{
		"filterLen":       2,
		"filterEmpty":     0,
		"filterBreak":     1,
		"filterLock":      2,
		"filteredCol":     2,
		"filteredColLock": 2,
		"filterPtr":       2,
		"filterPtrEmpty":  0,
		"filterPtrBreak":  1,
		"filterPtrLock":   2,
		"ptrLockEmpty":    0,
		"ptrLockBreak":    1,
	},
}

var srcC12CollectionNonEmptyDeepTestCase = coretestcases.CaseV1{
	Title: "Collection NonEmpty Items Hashset return correct -- filter blanks",
	ExpectedInput: args.Map{
		"nonEmptyList":     2,
		"nonEmptyListE":    0,
		"nonEmptyListPtr":  1,
		"hashsetAsIs":      true,
		"hashsetDouble":    true,
		"hashsetLock":      true,
		"nonEmptyItems":    2,
		"nonEmptyItemsPtr": 1,
		"nonEmptyWS":       2,
		"nonEmptyWSPtr":    1,
	},
}

var srcC12CollectionHasDeepTestCase = coretestcases.CaseV1{
	Title: "Collection Has HasPtr HasAll HasUsingSensitivity return correct -- deep",
	ExpectedInput: args.Map{
		"has":             true,
		"hasMiss":         false,
		"hasEmpty":        false,
		"hasPtrA":         true,
		"hasPtrNil":       false,
		"hasPtrEmpty":     false,
		"hasAll":          true,
		"hasAllMiss":      false,
		"hasAllEmpty":     false,
		"hasLock":         true,
		"sensCaseTrue":    true,
		"sensCaseFalse":   false,
		"sensInsensitive": true,
	},
}

var srcC12CollectionContainsExceptTestCase = coretestcases.CaseV1{
	Title: "Collection IsContains GetAllExcept GetHashsetPlusHasAll return correct -- deep",
	ExpectedInput: args.Map{
		"containsPtr":       true,
		"containsPtrNil":    false,
		"containsAllSlice":  true,
		"containsAllSliceF": false,
		"containsAllSliceE": false,
		"containsAll":       true,
		"containsAllLock":   true,
		"hashsetHasAll":     true,
		"hashsetHasAllNil":  false,
		"exceptColLen":      2,
		"exceptColNil":      1,
		"exceptLen":         2,
		"exceptNil":         1,
	},
}

var srcC12CollectionNewAddNonEmptyTestCase = coretestcases.CaseV1{
	Title: "Collection New AddNonEmptyStrings AddFuncResult AddStringsByFunc return correct -- various",
	ExpectedInput: args.Map{
		"newLen":           2,
		"newEmpty":         0,
		"addNonEmptyStr":   2,
		"addNonEmptyStrE":  0,
		"funcResultLen":    2,
		"funcResultNil":    0,
		"funcCheckLen":     2,
	},
}

var srcC12CollectionExpandMergeCharTestCase = coretestcases.CaseV1{
	Title: "Collection ExpandSlicePlusAdd MergeSlicesOfSlice CharCollectionMap return correct -- various",
	ExpectedInput: args.Map{
		"expandLen":  2,
		"mergeLen":   2,
		"charMapNN":  true,
	},
}

var srcC12CollectionStringDeepTestCase = coretestcases.CaseV1{
	Title: "Collection String Summary Csv Join return correct -- deep edge cases",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC12CollectionJsonDeepTestCase = coretestcases.CaseV1{
	Title: "Collection Json Marshal Unmarshal Serialize Deserialize return correct -- deep",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC12CollectionClearDisposeDeepTestCase = coretestcases.CaseV1{
	Title: "Collection Clear Dispose return correct -- nil safe",
	ExpectedInput: args.Map{
		"clearLen":   0,
		"nilClear":   true,
		"nilDispose": true,
		"noPanic":    true,
	},
}

var srcC12CollectionMiscTestCase = coretestcases.CaseV1{
	Title: "Collection AddPointerCollectionsLock ListCopyPtrLock Items ListPtr return correct -- misc",
	ExpectedInput: args.Map{
		"ptrColLock":    1,
		"listCopy":      1,
		"listCopyEmpty": 0,
		"itemsLen":      1,
		"listPtrLen":    1,
	},
}

var srcC12CollectionSortedDeepTestCase = coretestcases.CaseV1{
	Title: "Collection Sorted methods return correct -- deep edge cases",
	ExpectedInput: args.Map{
		"ascFirst":    "a",
		"ascEmptyLen": 0,
		"noPanic":     true,
	},
}

var srcC12CollectionCapResizeTestCase = coretestcases.CaseV1{
	Title: "Collection AddCapacity Resize return correct -- deep edge cases",
	ExpectedInput: args.Map{
		"capGe10":  true,
		"noPanic":  true,
		"resGe100": true,
	},
}

var srcC12CollectionJoinsDeepTestCase = coretestcases.CaseV1{
	Title: "Collection Joins NonEmptyJoins NonWhitespaceJoins Join JoinLine return correct -- deep",
	ExpectedInput: args.Map{
		"joins":         true,
		"joinsExtra":    true,
		"nonEmptyJoins": true,
		"nonWSJoins":    true,
		"joinAB":        "a,b",
		"joinEmpty":     "",
		"joinLine":      true,
		"joinLineEmpty": "",
	},
}
