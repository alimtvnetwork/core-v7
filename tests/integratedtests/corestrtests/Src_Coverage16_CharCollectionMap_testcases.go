package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── CharCollectionMap (C16) ──

var srcC16CcmCreatorsTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap creators Empty CapSelfCap Items ItemsPtrWithCap return correct -- various",
	ExpectedInput: args.Map{
		"emptyIsEmpty":    true,
		"capNonNil":       true,
		"itemsLen":        2,
		"itemsEmptyEmpty": true,
		"ptrCapLen":       2,
		"ptrCapEmptyE":    true,
	},
}

var srcC16CcmGetCharsGroupsTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap GetCharsGroups return correct -- items and empty",
	ExpectedInput: args.Map{
		"groupsLen":   2,
		"groupsEmpty": true,
	},
}

var srcC16CcmAddTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap Add AddStrings AddLock AddSameStartingCharItems return correct -- various",
	ExpectedInput: args.Map{
		"addLen":       2,
		"addSum":       3,
		"strLenX":      2,
		"strEmptyE":    true,
		"lockLen":      1,
		"sameLen":      2,
		"sameLenMore":  3,
		"sameEmptyE":   true,
	},
}

var srcC16CcmHasTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap Has HasWithCollection HasWithCollectionLock return correct -- various",
	ExpectedInput: args.Map{
		"hasFoo":           true,
		"hasBaz":           false,
		"hasZzz":           false,
		"hasEmptyAnything": false,
		"hwcHas":           true,
		"hwcColNonE":       true,
		"hwcMissHas":       false,
		"hwcEmptyHas":      false,
		"hwcLockHas":       true,
		"hwcLockMiss":      false,
		"hwcLockEmpty":     false,
	},
}

var srcC16CcmLengthTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap LengthOf LengthOfLock LengthOfCollectionFromFirstChar AllLengthsSum AllLengthsSumLock return correct -- various",
	ExpectedInput: args.Map{
		"lenOfA":        2,
		"lenOfZ":        0,
		"lenOfEmpty":    0,
		"lenOfLockA":    1,
		"lenOfLockZ":    0,
		"lenOfLockE":    0,
		"colFromA":      2,
		"colFromZ":      0,
		"allSum":        3,
		"allSumLock":    2,
	},
}

var srcC16CcmStateTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap IsEmpty HasItems IsEmptyLock return correct -- various",
	ExpectedInput: args.Map{
		"emptyIsEmpty":  true,
		"emptyHasItems": false,
		"addedIsEmpty":  false,
		"addedHasItems": true,
		"emptyLock":     true,
	},
}

var srcC16CcmEqualsTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap IsEquals IsEqualsLock IsEqualsCaseSensitive IsEqualsCaseSensitiveLock return correct -- various",
	ExpectedInput: args.Map{
		"equalSame":       true,
		"equalNil":        false,
		"equalSameRef":    true,
		"equalBothEmpty":  true,
		"equalOneEmpty":   false,
		"equalDiffLen":    false,
		"equalDiffCont":   false,
		"equalLock":       true,
		"equalCSLock":     true,
		"equalMissingKey": false,
	},
}

var srcC16CcmGetCollTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap GetCollection GetCollectionLock GetCollectionByChar return correct -- various",
	ExpectedInput: args.Map{
		"getLen":        2,
		"getMissNil":    true,
		"getCreateNN":   true,
		"getLockNN":     true,
		"getByCharNN":   true,
	},
}

var srcC16CcmAddSameCharsColTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap AddSameCharsCollection AddSameCharsCollectionLock return correct -- various",
	ExpectedInput: args.Map{
		"addLen":          3,
		"addMoreLen":      3,
		"addNilNN":        true,
		"addExistNilNN":   true,
		"lockNN":          true,
		"lockNilNN":       true,
		"lockExistNilNN":  true,
		"lockAddExistNN":  true,
	},
}

var srcC16CcmAddColItemsHmTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap AddCollectionItems AddHashmapsValues AddHashmapsKeysValuesBoth AddHashmapsKeysOrValuesBothUsingFilter AddCharHashsetMap return correct -- various",
	ExpectedInput: args.Map{
		"colItemsLen":     2,
		"colItemsNilE":    true,
		"hmValsHasA":      true,
		"hmValsNilE":      true,
		"hmBothHasK":      true,
		"hmBothNilE":      true,
		"hmFilterHasK":    true,
		"hmFilterNilE":    true,
		"hmFilterBreak":   1,
		"charHsLen":       2,
	},
}

var srcC16CcmResizeListTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap Resize AddLength List ListLock SortedListAsc return correct -- various",
	ExpectedInput: args.Map{
		"resizeHas":     true,
		"resizeNoShrink": 2,
		"addLenHas":     true,
		"addLenEmptyOk": true,
		"listLen":       2,
		"listEmptyLen":  0,
		"listLockLen":   1,
		"sortedFirst":   "apple",
		"sortedEmptyLen": 0,
	},
}

var srcC16CcmMapStringTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap GetMap GetCopyMapLock String SummaryString StringLock SummaryStringLock Print return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC16CcmHashsetTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap HashsetByChar HashsetByCharLock HashsetByStringFirstChar HashsetByStringFirstCharLock HashsetsCollection HashsetsCollectionByChars HashsetsCollectionByStringFirstChar return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC16CcmJsonTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap JsonModel JsonModelAny MarshalJSON UnmarshalJSON Json JsonPtr ParseInjectUsingJson JsonParseSelfInject AsJson interfaces return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC16CcmClearDataModelTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap Clear DataModel return correct -- various",
	ExpectedInput: args.Map{
		"clearHasItems":  false,
		"clearEmptyOk":   true,
		"dataModelNN":    true,
	},
}
