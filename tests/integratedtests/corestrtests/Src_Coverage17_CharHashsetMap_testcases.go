package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── CharHashsetMap creators (C17) ──

var srcC17CreatorsTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap creators Cap CapItems Strings Empty return correct -- various",
	ExpectedInput: args.Map{
		"capNN":       true,
		"capEmpty":    true,
		"capItemsLen": 2,
		"stringsLen":  2,
		"stringsNilNN": true,
		"emptyNN":     true,
		"emptyEmpty":  true,
	},
}

var srcC17GetCharsGroupsTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap GetCharsGroups return correct -- items and empty",
	ExpectedInput: args.Map{
		"groupsLen":    2,
		"emptySameRef": true,
	},
}

var srcC17AddTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap Add AddStrings AddLock AddStringsLock AddSameStartingCharItems return correct -- various",
	ExpectedInput: args.Map{
		"addLen":        2,
		"addSum":        3,
		"addStrLenX":    2,
		"addStrEmptyE":  true,
		"addLockLen":    1,
		"addStrLockLen": 1,
		"addStrLockEE":  true,
		"sameLen":       2,
		"sameMoreLen":   3,
		"sameEmptyE":    true,
	},
}

var srcC17AddCollTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap AddCollectionItems AddCharCollectionMapItems AddCollectionItemsAsyncLock return correct -- various",
	ExpectedInput: args.Map{
		"colLen":    2,
		"colNilE":   true,
		"ccmLen":    2,
		"ccmNilE":   true,
		"asyncDone": true,
	},
}

var srcC17HasTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap Has HasWithHashset HasWithHashsetLock return correct -- various",
	ExpectedInput: args.Map{
		"hasFoo":        true,
		"hasBaz":        false,
		"hasZzz":        false,
		"hasEmptyX":     false,
		"hwHasFoo":      true,
		"hwHsNonE":      true,
		"hwMissHas":     false,
		"hwEmptyHas":    false,
		"hwEmptyHsNN":   true,
		"hwlHasFoo":     true,
		"hwlHsNN":       true,
		"hwlMissHas":    false,
		"hwlEmptyHas":   false,
	},
}

var srcC17LengthTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap LengthOf LengthOfLock AllLengthsSum AllLengthsSumLock LengthOfHashsetFromFirstChar return correct -- various",
	ExpectedInput: args.Map{
		"lenOfA":       2,
		"lenOfZ":       0,
		"lenOfEmptyA":  0,
		"lockA":        1,
		"lockZ":        0,
		"lockEmptyA":   0,
		"allSum":       3,
		"allSumEmpty":  0,
		"allSumLock":   2,
		"allSumLockE":  0,
	},
}

var srcC17StateTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap IsEmpty HasItems IsEmptyLock return correct -- various",
	ExpectedInput: args.Map{
		"emptyIsEmpty": true,
		"emptyHasIt":   false,
		"addedIsEmpty": false,
		"addedHasIt":   true,
		"emptyLock":    true,
	},
}

var srcC17EqualsTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap IsEquals IsEqualsLock return correct -- various",
	ExpectedInput: args.Map{
		"equalSame":      true,
		"equalNil":       false,
		"equalSameRef":   true,
		"equalBothEmpty": true,
		"equalOneEmpty":  false,
		"equalDiffLen":   false,
		"equalDiffCont":  false,
		"equalMissKey":   false,
		"equalLock":      true,
	},
}

var srcC17GetHashsetTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap GetHashset GetHashsetLock GetHashsetByChar HashsetByChar HashsetByCharLock HashsetByStringFirstChar return correct -- various",
	ExpectedInput: args.Map{
		"getNN":           true,
		"getMissNil":      true,
		"getCreateNN":     true,
		"getLockNN":       true,
		"getByCharNN":     true,
		"hsByCharNN":      true,
		"hsByCharLockNN":  true,
		"hsByCharLockZE":  true,
		"hsByStrNN":       true,
		"hsByStrLockNN":   true,
	},
}

var srcC17AddSameCharsTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap AddSameCharsCollection AddSameCharsHashset AddSameCharsCollectionLock AddHashsetLock return correct -- various",
	ExpectedInput: args.Map{
		"colLen":         3,
		"colMoreLen":     3,
		"colNilNN":       true,
		"colExNilNN":     true,
		"hsLen":          2,
		"hsNilNN":        true,
		"hsExNilNN":      true,
		"hsAddExLen":     2,
		"colLockNN":      true,
		"colLockNilNN":   true,
		"colLockExNilNN": true,
		"colLockExAddNN": true,
		"hsLockNN":       true,
		"hsLockNilNN":    true,
		"hsLockExNilNN":  true,
		"hsLockExAddNN":  true,
	},
}

var srcC17AddHashsetItemsTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap AddHashsetItems AddHashsetItemsAsyncLock return correct -- various",
	ExpectedInput: args.Map{
		"itemsSum":  2,
		"itemsEmE":  true,
		"asyncNilOk": true,
	},
}

var srcC17HashsetsCollTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap HashsetsCollection HashsetsCollectionByChars HashsetsCollectionByStringsFirstChar return correct -- various",
	ExpectedInput: args.Map{
		"hscNN":       true,
		"hscEmptyNN":  true,
		"hscCharsNN":  true,
		"hscCharsENN": true,
		"hscStrNN":    true,
		"hscStrENN":   true,
	},
}

var srcC17ListSortTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap List SortedListAsc SortedListDsc return correct -- various",
	ExpectedInput: args.Map{
		"listLen":     2,
		"ascFirst":    "apple",
		"ascLen":      3,
		"dscFirst":    "cherry",
		"dscLen":      3,
	},
}

var srcC17MapStringPrintTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap GetMap GetCopyMapLock String SummaryString Print return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC17JsonTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap JSON marshal unmarshal Json JsonPtr ParseInjectUsingJson ParseInjectUsingJsonMust JsonParseSelfInject As* return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC17ClearDataModelTestCase = coretestcases.CaseV1{
	Title: "CharHashsetMap RemoveAll Clear DataModel return correct -- various",
	ExpectedInput: args.Map{
		"removeAllE":   true,
		"clearE":       true,
		"clearEmptyOk": true,
		"dataModelNN":  true,
	},
}
