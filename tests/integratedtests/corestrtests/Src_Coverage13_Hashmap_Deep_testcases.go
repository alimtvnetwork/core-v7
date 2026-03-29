package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── Hashmap Deep (C13) ──

var srcC13HashmapDiffTestCase = coretestcases.CaseV1{
	Title: "Hashmap DiffRaw Diff return non-empty -- changed value",
	ExpectedInput: args.Map{
		"diffRawGt0": true,
		"diffNonE":   true,
	},
}

var srcC13HashmapHasAllCollTestCase = coretestcases.CaseV1{
	Title: "Hashmap HasAllCollectionItems returns correct -- items nil empty",
	ExpectedInput: args.Map{
		"hasAll":    true,
		"hasNil":    false,
		"hasEmpty":  false,
	},
}

var srcC13HashmapHasVariantsTestCase = coretestcases.CaseV1{
	Title: "Hashmap HasAll HasAnyItem HasAny HasWithLock return correct -- various",
	ExpectedInput: args.Map{
		"hasAll":       true,
		"hasAllMiss":   false,
		"hasAnyItem":   true,
		"emptyAnyItem": false,
		"hasAny":       true,
		"hasAnyMiss":   false,
		"hasWithLock":  true,
		"hasWLMiss":    false,
	},
}

var srcC13HashmapFilterTestCase = coretestcases.CaseV1{
	Title: "Hashmap GetKeysFilteredItems GetKeysFilteredCollection return correct -- various",
	ExpectedInput: args.Map{
		"filterLen":     2,
		"filterEmpty":   0,
		"filterBreak":   1,
		"filterSkip":    0,
		"colNonEmpty":   true,
		"colEmpty":      true,
		"colBreakLen":   1,
	},
}

var srcC13HashmapItemsTestCase = coretestcases.CaseV1{
	Title: "Hashmap Items SafeItems ItemsCopyLock return correct -- various",
	ExpectedInput: args.Map{
		"itemsLen":    1,
		"safeLen":     1,
		"safeNilNil":  true,
		"copyLen":     1,
	},
}

var srcC13HashmapValuesTestCase = coretestcases.CaseV1{
	Title: "Hashmap ValuesCollection ValuesHashset ValuesList ValuesCollectionLock ValuesHashsetLock return correct -- items",
	ExpectedInput: args.Map{
		"colNonE":      true,
		"hsNonE":       true,
		"colLockNonE":  true,
		"hsLockNonE":   true,
		"valsLen":      1,
	},
}

var srcC13HashmapKeysValuesTestCase = coretestcases.CaseV1{
	Title: "Hashmap KeysValuesCollection KeysValuesList KeysValuePairs KeysValuePairsCollection KeysValuesListLock return correct -- items",
	ExpectedInput: args.Map{
		"kvColNonE":       true,
		"kvListLen":       1,
		"pairsLen":        1,
		"pairsColNonNil":  true,
		"kvListLockLen":   1,
	},
}

var srcC13HashmapKeysTestCase = coretestcases.CaseV1{
	Title: "Hashmap AllKeys Keys KeysCollection KeysLock ValuesListCopyLock return correct -- various",
	ExpectedInput: args.Map{
		"allKeysLen":      1,
		"allKeysEmpty":    0,
		"keysLen":         1,
		"keysColNonE":     true,
		"keysLockLen":     1,
		"keysLockEmpty":   0,
		"valsCopyLen":     1,
	},
}

var srcC13HashmapLowerTestCase = coretestcases.CaseV1{
	Title: "Hashmap KeysToLower ValuesToLower return correct -- ABC",
	ExpectedInput: args.Map{
		"keysLower":  true,
		"valsLower":  true,
	},
}

var srcC13HashmapLengthTestCase = coretestcases.CaseV1{
	Title: "Hashmap Length LengthLock return correct -- items nil",
	ExpectedInput: args.Map{
		"length":     1,
		"nilLength":  0,
		"lockLength": 1,
	},
}

var srcC13HashmapEqualTestCase = coretestcases.CaseV1{
	Title: "Hashmap IsEqual IsEqualPtr IsEqualPtrLock return correct -- various",
	ExpectedInput: args.Map{
		"equal":         true,
		"ptrEqual":      true,
		"bothNil":       true,
		"oneNil":        false,
		"samePtr":       true,
		"bothEmpty":     true,
		"diffLen":       false,
		"diffVal":       false,
		"ptrLockEqual":  true,
	},
}

var srcC13HashmapRemoveTestCase = coretestcases.CaseV1{
	Title: "Hashmap Remove RemoveWithLock return correct -- key a",
	ExpectedInput: args.Map{
		"removed":     true,
		"removedLock": true,
	},
}

var srcC13HashmapStringTestCase = coretestcases.CaseV1{
	Title: "Hashmap String StringLock return non-empty -- items and empty",
	ExpectedInput: args.Map{
		"strNonE":      true,
		"strEmptyNonE": true,
		"lockNonE":     true,
		"lockEmptyNE":  true,
	},
}

var srcC13HashmapExceptTestCase = coretestcases.CaseV1{
	Title: "Hashmap GetValuesExceptKeysInHashset GetValuesKeysExcept GetAllExceptCollection return correct -- various",
	ExpectedInput: args.Map{
		"exceptHsLen":    1,
		"exceptHsNil":    1,
		"exceptKeysLen":  1,
		"exceptKeysNil":  1,
		"exceptColLen":   1,
		"exceptColNil":   1,
	},
}

var srcC13HashmapJoinJsonTestCase = coretestcases.CaseV1{
	Title: "Hashmap Join JoinKeys JsonModel JsonModelAny MarshalJSON UnmarshalJSON Json JsonPtr return correct -- various",
	ExpectedInput: args.Map{
		"joinNonE":     true,
		"joinKeysNonE": true,
		"jsonModelLen": 1,
		"marshalOk":    true,
		"unmarshalLen": 1,
		"unmarshalErr": true,
		"jsonNoErr":    true,
		"jsonPtrNoErr": true,
		"noPanic":      true,
	},
}

var srcC13HashmapErrorTestCase = coretestcases.CaseV1{
	Title: "Hashmap ToError ToDefaultError return correct -- items",
	ExpectedInput: args.Map{
		"toErr":     true,
		"toDefErr":  true,
	},
}

var srcC13HashmapMiscTestCase = coretestcases.CaseV1{
	Title: "Hashmap KeyValStringLines Clear Serialize Deserialize AsJson return correct -- various",
	ExpectedInput: args.Map{
		"kvLinesLen":    1,
		"clearLen":      0,
		"clearNilOk":    true,
		"serializeOk":   true,
		"deserializeOk": true,
		"noPanic":       true,
	},
}

var srcC13HashmapCompilerTestCase = coretestcases.CaseV1{
	Title: "Hashmap ToStringsUsingCompiler return correct -- items and empty",
	ExpectedInput: args.Map{
		"compLen":    1,
		"compEmpty":  0,
	},
}

var srcC13HashmapCloneTestCase = coretestcases.CaseV1{
	Title: "Hashmap ClonePtr Clone return correct -- items nil empty",
	ExpectedInput: args.Map{
		"cloneLen":     2,
		"cloneNilNil":  true,
		"cloneIndep":   true,
		"cloneEmpty":   0,
	},
}

var srcC13HashmapGetTestCase = coretestcases.CaseV1{
	Title: "Hashmap Get GetValue return correct -- key a",
	ExpectedInput: args.Map{
		"getVal":      "1",
		"getFound":    true,
		"getValVal":   "1",
		"getValFound": true,
	},
}

var srcC13HashmapFilterVariantsTestCase = coretestcases.CaseV1{
	Title: "Hashmap AddsOrUpdatesAnyUsingFilter AddsOrUpdatesUsingFilter variants return correct -- various",
	ExpectedInput: args.Map{
		"anyFilter":       1,
		"anyFilterNil":    0,
		"anyFilterBreak":  1,
		"anyFilterSkip":   0,
		"anyFilterLock":   1,
		"anyFilterLNil":   0,
		"anyFilterLBreak": 1,
		"filter":          1,
		"filterNil":       0,
		"filterBreak":     1,
		"filterSkip":      0,
	},
}

var srcC13HashmapWgLockTestCase = coretestcases.CaseV1{
	Title: "Hashmap AddOrUpdateWithWgLock AddOrUpdateStringsPtrWgLock return correct -- various",
	ExpectedInput: args.Map{
		"wgLockHas":   true,
		"strPtrHas":   true,
		"strPtrEmpty": 0,
		"panicOnMismatch": true,
	},
}

var srcC13HashmapDiffTypeTestCase = coretestcases.CaseV1{
	Title: "HashmapDiff methods return correct -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}
