package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// Branch: Hashmap basic empty state
var srcC06HashmapBasicTestCase = coretestcases.CaseV1{
	Title: "Hashmap Empty returns correct state -- new empty hashmap",
	ExpectedInput: args.Map{
		"isEmpty":       true,
		"hasItems":      false,
		"hasAny":        false,
		"length":        0,
		"nilLength":     0,
		"nilSafeItems":  true,
	},
}

// Branch: AddOrUpdate and Get
var srcC06HashmapAddGetTestCase = coretestcases.CaseV1{
	Title: "Hashmap AddOrUpdate Get return correct state -- add then update",
	ExpectedInput: args.Map{
		"isNewFirst":  true,
		"isNewSecond": false,
		"getValue":    "v2",
		"found":       true,
	},
}

// Branch: Set variants
var srcC06HashmapSetTestCase = coretestcases.CaseV1{
	Title: "Hashmap Set SetTrim SetBySplitter return correct length -- 4 items",
	ExpectedInput: args.Map{
		"afterSetTrim": 2,
		"afterSplit":   4,
	},
}

// Branch: Has/Contains methods
var srcC06HashmapHasTestCase = coretestcases.CaseV1{
	Title: "Hashmap Has Contains HasAll HasAny return correct -- a b map",
	ExpectedInput: args.Map{
		"has":          true,
		"contains":     true,
		"notMissing":   false,
		"zMissing":     true,
		"hasAll":       true,
		"hasAllStr":    true,
		"hasAnyAZ":     true,
		"hasAnyXZ":     false,
	},
}

// Branch: HasLock methods
var srcC06HashmapHasLockTestCase = coretestcases.CaseV1{
	Title: "Hashmap lock methods return correct -- a key",
	ExpectedInput: args.Map{
		"hasLock":      true,
		"hasWithLock":  true,
		"containsLock": true,
		"notMissing":   false,
	},
}

// Branch: Add variants (typed keys)
var srcC06HashmapAddVariantsTestCase = coretestcases.CaseV1{
	Title: "Hashmap AddOrUpdate typed variants return correct length -- 7 items",
	ExpectedInput: args.Map{
		"length": 7,
	},
}

// Branch: AddOrUpdateHashmap/Map
var srcC06HashmapMergeTestCase = coretestcases.CaseV1{
	Title: "Hashmap AddOrUpdateHashmap Map return correct lengths -- merge two",
	ExpectedInput: args.Map{
		"afterHashmap": 2,
		"afterMap":     3,
	},
}

// Branch: AddsOrUpdates
var srcC06HashmapAddsTestCase = coretestcases.CaseV1{
	Title: "Hashmap AddsOrUpdates methods return correct length -- 3 items",
	ExpectedInput: args.Map{
		"length": 3,
	},
}

// Branch: AddOrUpdateCollection
var srcC06HashmapCollectionAddTestCase = coretestcases.CaseV1{
	Title: "Hashmap AddOrUpdateCollection returns correct length -- keys vals",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

// Branch: WgLock
var srcC06HashmapWgLockTestCase = coretestcases.CaseV1{
	Title: "Hashmap AddOrUpdateWithWgLock returns correct length -- 1 item",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

// Branch: Filter methods
var srcC06HashmapFilterTestCase = coretestcases.CaseV1{
	Title: "Hashmap GetKeysFilteredItems Collection return correct -- abc filter",
	ExpectedInput: args.Map{
		"filteredItemsLen": 1,
		"filteredColLen":   1,
		"emptyLen":         0,
	},
}

// Branch: Filter with break
var srcC06HashmapFilterBreakTestCase = coretestcases.CaseV1{
	Title: "Hashmap filter with break returns 1 -- break on first",
	ExpectedInput: args.Map{
		"itemsLen": 1,
		"colLen":   1,
	},
}

// Branch: Keys and Values
var srcC06HashmapKeysTestCase = coretestcases.CaseV1{
	Title: "Hashmap Keys AllKeys KeysCollection return correct -- 2 items",
	ExpectedInput: args.Map{
		"allKeysLen":     2,
		"keysLen":        2,
		"keysColLen":     2,
	},
}

var srcC06HashmapValuesTestCase = coretestcases.CaseV1{
	Title: "Hashmap ValuesList ValuesCollection return correct -- 1 item",
	ExpectedInput: args.Map{
		"valuesListLen":   1,
		"kvColKeysLen":    1,
		"kvColValsLen":    1,
		"kvListKeysLen":   1,
		"kvListValsLen":   1,
	},
}

// Branch: KeyValuePairs
var srcC06HashmapPairsTestCase = coretestcases.CaseV1{
	Title: "Hashmap KeysValuePairs return correct -- 1 item",
	ExpectedInput: args.Map{
		"pairsLen":    1,
		"pairsColLen": 1,
	},
}

// Branch: Remove
var srcC06HashmapRemoveTestCase = coretestcases.CaseV1{
	Title: "Hashmap Remove RemoveWithLock return correct lengths -- 2 items",
	ExpectedInput: args.Map{
		"afterRemove":     1,
		"afterRemoveLock": 0,
	},
}

// Branch: Diff
var srcC06HashmapDiffTestCase = coretestcases.CaseV1{
	Title: "Hashmap Diff DiffRaw execute without panic -- a key",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// Branch: IsEqual
var srcC06HashmapEqualTestCase = coretestcases.CaseV1{
	Title: "Hashmap IsEqualPtr returns correct -- same and different",
	ExpectedInput: args.Map{
		"equalSame":      true,
		"equalSameLock":  true,
		"equalDifferent": false,
	},
}

// Branch: ConcatNew
var srcC06HashmapConcatTestCase = coretestcases.CaseV1{
	Title: "Hashmap ConcatNew returns correct length -- merge two",
	ExpectedInput: args.Map{
		"concatLen": true,
		"noPanic":   true,
	},
}

// Branch: String and JSON
var srcC06HashmapStringJsonTestCase = coretestcases.CaseV1{
	Title: "Hashmap String Json methods return non-empty -- 1 item",
	ExpectedInput: args.Map{
		"stringNonEmpty":     true,
		"stringLockNonEmpty": true,
		"noPanic":            true,
	},
}

// Branch: KeysToLower
var srcC06HashmapKeysToLowerTestCase = coretestcases.CaseV1{
	Title: "Hashmap KeysToLower returns lowercase key -- ABC",
	ExpectedInput: args.Map{
		"hasLowercase": true,
	},
}

// Branch: GetExcept methods
var srcC06HashmapExceptTestCase = coretestcases.CaseV1{
	Title: "Hashmap GetExcept methods return correct lengths -- exclude a",
	ExpectedInput: args.Map{
		"hashsetExcLen":    1,
		"keysExcLen":       1,
		"collectionExcLen": 1,
	},
}

// Branch: HasAllCollectionItems
var srcC06HashmapHasAllColTestCase = coretestcases.CaseV1{
	Title: "Hashmap HasAllCollectionItems returns correct -- a b keys",
	ExpectedInput: args.Map{
		"hasAll":    true,
		"hasAllNil": false,
	},
}

// Branch: Clone
var srcC06HashmapCloneTestCase = coretestcases.CaseV1{
	Title: "Hashmap Clone ClonePtr return correct -- 1 item",
	ExpectedInput: args.Map{
		"cloneLen":    1,
		"clonePtrLen": 1,
		"nilClone":    true,
	},
}

// Branch: ToStringsUsingCompiler
var srcC06HashmapCompilerTestCase = coretestcases.CaseV1{
	Title: "Hashmap ToStringsUsingCompiler returns correct -- 1 item",
	ExpectedInput: args.Map{
		"compiledLen": 1,
		"emptyLen":    0,
	},
}

// Branch: AddOrUpdateStringsPtrWgLock
var srcC06HashmapStringsPtrWgLockTestCase = coretestcases.CaseV1{
	Title: "Hashmap AddOrUpdateStringsPtrWgLock returns correct length -- 1 item",
	ExpectedInput: args.Map{
		"length": 1,
	},
}
