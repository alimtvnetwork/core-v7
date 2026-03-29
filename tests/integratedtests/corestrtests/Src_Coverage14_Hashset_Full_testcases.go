package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── Hashset Full (C14) ──

var srcC14HashsetCapResizeTestCase = coretestcases.CaseV1{
	Title: "Hashset AddCapacitiesLock AddCapacities Resize ResizeLock return correct -- various",
	ExpectedInput: args.Map{
		"capLockHas":  true,
		"capHas":      true,
		"resizeHas":   true,
		"resizeLHas":  true,
	},
}

var srcC14HashsetCollMiscTestCase = coretestcases.CaseV1{
	Title: "Hashset Collection IsEmptyLock ConcatNewHashsets ConcatNewStrings return correct -- various",
	ExpectedInput: args.Map{
		"colNonE":       true,
		"emptyLock":     true,
		"concatHsGe2":   true,
		"concatHsEGe1":  true,
		"concatStrGe2":  true,
		"concatStrEGe1": true,
	},
}

var srcC14HashsetAddVariantsTestCase = coretestcases.CaseV1{
	Title: "Hashset Add variants return correct -- various",
	ExpectedInput: args.Map{
		"addPtrHas":        true,
		"addPtrLockHas":    true,
		"addBoolFirst":     false,
		"addBoolSecond":    true,
		"nonEmptyLen":      1,
		"nonEmptyWSLen":    1,
		"addIfLen":         1,
		"addIfManyLen":     2,
		"addFuncLen":       1,
		"addStringsLen":    2,
		"addStringsLkLen":  1,
		"addCollLen":       1,
		"addCollsLen":      2,
		"addHsItemsHas":   true,
		"addItemsMapHas":  true,
		"addLockHas":       true,
		"addsLen":          2,
	},
}

var srcC14HashsetWgLockTestCase = coretestcases.CaseV1{
	Title: "Hashset AddWithWgLock return correct -- sync",
	ExpectedInput: args.Map{
		"has": true,
	},
}

var srcC14HashsetQueryTestCase = coretestcases.CaseV1{
	Title: "Hashset HasAnyItem IsMissing IsMissingLock Contains IsEqual SortedList Filter return correct -- various",
	ExpectedInput: args.Map{
		"hasAnyItem":     true,
		"isMissingA":     false,
		"isMissingB":     true,
		"isMissingLockA": false,
		"containsA":      true,
		"isEqual":        true,
		"sortedFirst":    "a",
		"filterLen":      1,
	},
}

var srcC14HashsetHasVariantsTestCase = coretestcases.CaseV1{
	Title: "Hashset HasLock HasAllStrings HasAllCollectionItems HasAll IsAllMissing HasAny HasWithLock return correct -- various",
	ExpectedInput: args.Map{
		"hasLock":        true,
		"hasAllStr":      true,
		"hasAllColl":     true,
		"hasAll":         true,
		"allMissingT":    true,
		"allMissingF":    false,
		"hasAny":         true,
		"hasAnyMiss":     false,
		"hasWithLock":    true,
	},
}

var srcC14HashsetListTestCase = coretestcases.CaseV1{
	Title: "Hashset OrderedList SafeStrings Lines SimpleSlice return correct -- various",
	ExpectedInput: args.Map{
		"orderedFirst":  "a",
		"safeStrLen":    1,
		"linesLen":      1,
		"simpleSliceNN": true,
	},
}

var srcC14HashsetFilterExceptTestCase = coretestcases.CaseV1{
	Title: "Hashset GetFilteredItems GetFilteredCollection GetAllExcept variants return correct -- various",
	ExpectedInput: args.Map{
		"filteredLen":   2,
		"filteredColNE": true,
		"exceptHsLen":   2,
		"exceptLen":     1,
		"exceptSpLen":   1,
		"exceptColLen":  1,
	},
}

var srcC14HashsetItemsTestCase = coretestcases.CaseV1{
	Title: "Hashset Items List MapStringAny return correct -- various",
	ExpectedInput: args.Map{
		"itemsLen":     1,
		"listLen":      1,
		"mapStrAnyLen": 1,
		"noPanic":      true,
	},
}

var srcC14HashsetSortJoinTestCase = coretestcases.CaseV1{
	Title: "Hashset JoinSorted ListPtrSortedAsc ListPtrSortedDsc return correct -- various",
	ExpectedInput: args.Map{
		"joinSortedNE": true,
		"ascFirst":     "a",
		"dscFirst":     "c",
	},
}

var srcC14HashsetClearTestCase = coretestcases.CaseV1{
	Title: "Hashset Clear ListCopyLock ToLowerSet LengthLock return correct -- various",
	ExpectedInput: args.Map{
		"clearLen":    0,
		"listCopyLen": 1,
		"lowerHas":    true,
		"lengthLock":  1,
	},
}

var srcC14HashsetEqualRemoveTestCase = coretestcases.CaseV1{
	Title: "Hashset IsEquals IsEqualsLock Remove SafeRemove RemoveWithLock return correct -- various",
	ExpectedInput: args.Map{
		"isEquals":     true,
		"isEqualsLock": true,
		"removeOk":     true,
		"safeRemoveOk": true,
		"removeLockOk": true,
	},
}

var srcC14HashsetStringTestCase = coretestcases.CaseV1{
	Title: "Hashset String StringLock return correct -- items",
	ExpectedInput: args.Map{
		"strNonE":     true,
		"strLockNonE": true,
	},
}

var srcC14HashsetUnmarshalTestCase = coretestcases.CaseV1{
	Title: "Hashset UnmarshalJSON return correct -- valid json",
	ExpectedInput: args.Map{
		"unmarshalLen": 1,
	},
}

var srcC14HashsetJsonInterfacesTestCase = coretestcases.CaseV1{
	Title: "Hashset AsJson interfaces return correct -- noPanic",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC14HashsetDiffTestCase = coretestcases.CaseV1{
	Title: "Hashset DistinctDiffLinesRaw DistinctDiffHashset DistinctDiffLines return correct -- various",
	ExpectedInput: args.Map{
		"diffRawLen":    2,
		"diffRawEmpty":  0,
		"diffHsLen":     2,
		"diffLinesLen":  2,
	},
}

var srcC14HashsetSerializeTestCase = coretestcases.CaseV1{
	Title: "Hashset Serialize Deserialize return correct -- items",
	ExpectedInput: args.Map{
		"serializeOk":   true,
		"deserializeOk": true,
	},
}

var srcC14HashsetWrapTestCase = coretestcases.CaseV1{
	Title: "Hashset WrapDoubleQuote WrapSingleQuote WrapIfMissing Transpile JoinLine return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC14HashsetFilterAddTestCase = coretestcases.CaseV1{
	Title: "Hashset AddsUsingFilter AddsAnyUsingFilter AddsAnyUsingFilterLock AddFuncErr AddStringsPtrWgLock AddHashsetWgLock AddSimpleSlice ListPtr return correct -- various",
	ExpectedInput: args.Map{
		"addsFilterLen":     2,
		"addsAnyFilterLen":  1,
		"addsAnyFLockLen":   1,
		"funcErrLen":        1,
		"strPtrWgHas":       true,
		"hsWgHas":           true,
		"simpleSliceLen":    2,
		"listPtrLen":        1,
	},
}
