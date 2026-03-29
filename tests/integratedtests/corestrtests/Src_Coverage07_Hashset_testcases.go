package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// Branch: Hashset basic empty state
var srcC07HashsetBasicTestCase = coretestcases.CaseV1{
	Title: "Hashset Empty returns correct state -- new empty hashset",
	ExpectedInput: args.Map{
		"isEmpty":   true,
		"hasItems":  false,
		"hasAny":    false,
		"length":    0,
		"nilLength": 0,
	},
}

// Branch: Add methods
var srcC07HashsetAddTestCase = coretestcases.CaseV1{
	Title: "Hashset Add variants execute without panic -- multiple add types",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// Branch: AddBool
var srcC07HashsetAddBoolTestCase = coretestcases.CaseV1{
	Title: "Hashset AddBool returns correct existed flag -- add twice",
	ExpectedInput: args.Map{
		"firstExisted":  false,
		"secondExisted": true,
	},
}

// Branch: AddPtr
var srcC07HashsetAddPtrTestCase = coretestcases.CaseV1{
	Title: "Hashset AddPtr AddPtrLock deduplicate -- same string",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

// Branch: Adds/AddStrings/AddCollection
var srcC07HashsetAddsTestCase = coretestcases.CaseV1{
	Title: "Hashset Adds AddStrings AddCollection execute without panic -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// Branch: Has/Contains
var srcC07HashsetHasTestCase = coretestcases.CaseV1{
	Title: "Hashset Has Contains HasAll HasAny return correct -- a b c",
	ExpectedInput: args.Map{
		"has":          true,
		"contains":     true,
		"hasLock":      true,
		"hasWithLock":  true,
		"notMissing":   false,
		"zMissing":     true,
		"missingLock":  false,
		"hasAll":       true,
		"hasAllStr":    true,
		"hasAnyAZ":     true,
		"hasAnyXZ":     false,
		"isAllMissXZ":  true,
		"isAllMissA":   false,
		"hasAllCol":    true,
		"hasAllColNil": false,
	},
}

// Branch: IsEquals
var srcC07HashsetEqualsTestCase = coretestcases.CaseV1{
	Title: "Hashset IsEquals IsEqual IsEqualsLock return true -- same sets",
	ExpectedInput: args.Map{
		"isEquals":     true,
		"isEqual":      true,
		"isEqualsLock": true,
	},
}

// Branch: Remove
var srcC07HashsetRemoveTestCase = coretestcases.CaseV1{
	Title: "Hashset Remove SafeRemove RemoveWithLock execute without panic -- a b",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// Branch: List methods
var srcC07HashsetListTestCase = coretestcases.CaseV1{
	Title: "Hashset List methods execute without panic -- one item",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// Branch: Filter
var srcC07HashsetFilterTestCase = coretestcases.CaseV1{
	Title: "Hashset Filter GetFilteredItems GetFilteredCollection return correct -- abc",
	ExpectedInput: args.Map{
		"filterLen":      1,
		"filteredItems":  1,
		"filteredColLen": 1,
	},
}

// Branch: GetAllExcept
var srcC07HashsetExceptTestCase = coretestcases.CaseV1{
	Title: "Hashset GetAllExcept methods return correct length -- exclude a",
	ExpectedInput: args.Map{
		"hashsetExcLen":    2,
		"keysExcLen":       2,
		"spreadExcLen":     2,
		"collectionExcLen": 2,
	},
}

// Branch: Resize / AddCapacities
var srcC07HashsetResizeTestCase = coretestcases.CaseV1{
	Title: "Hashset Resize AddCapacities execute without panic -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// Branch: ConcatNew
var srcC07HashsetConcatTestCase = coretestcases.CaseV1{
	Title: "Hashset ConcatNewHashsets ConcatNewStrings return correct -- merge",
	ExpectedInput: args.Map{
		"concatLen": true,
		"noPanic":   true,
	},
}

// Branch: String/Json
var srcC07HashsetStringJsonTestCase = coretestcases.CaseV1{
	Title: "Hashset String Json methods return non-empty -- one item",
	ExpectedInput: args.Map{
		"stringNonEmpty":     true,
		"stringLockNonEmpty": true,
		"noPanic":            true,
	},
}

// Branch: ToLowerSet
var srcC07HashsetToLowerTestCase = coretestcases.CaseV1{
	Title: "Hashset ToLowerSet returns lowercase -- ABC",
	ExpectedInput: args.Map{
		"hasLowercase": true,
	},
}

// Branch: DistinctDiff
var srcC07HashsetDistinctDiffTestCase = coretestcases.CaseV1{
	Title: "Hashset DistinctDiffLinesRaw returns correct -- a b diff b c",
	ExpectedInput: args.Map{
		"diffLen": 2,
		"noPanic": true,
	},
}

// Branch: WgLock
var srcC07HashsetWgLockTestCase = coretestcases.CaseV1{
	Title: "Hashset AddWithWgLock AddStringsPtrWgLock execute without panic -- sync",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// Branch: AddItemsMap
var srcC07HashsetAddItemsMapTestCase = coretestcases.CaseV1{
	Title: "Hashset AddItemsMap returns correct -- true values only",
	ExpectedInput: args.Map{
		"length":  2, // "a" from AddItemsMap + "c" from AddItemsMapWgLock
		"noPanic": true,
	},
}

// Branch: AddHashset
var srcC07HashsetAddHashsetTestCase = coretestcases.CaseV1{
	Title: "Hashset AddHashsetItems AddHashsetWgLock execute without panic -- merge",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// Branch: AddsUsingFilter
var srcC07HashsetAddsFilterTestCase = coretestcases.CaseV1{
	Title: "Hashset AddsUsingFilter variants execute without panic -- filter",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// Branch: Empty string
var srcC07HashsetEmptyStringTestCase = coretestcases.CaseV1{
	Title: "Hashset empty String returns NoElements -- empty set",
	ExpectedInput: args.Map{
		"stringNonEmpty":     true,
		"stringLockNonEmpty": true,
	},
}
