package coregenerictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// MapCollection
// ==========================================================================

var mapCollectionIntToStringTestCase = coretestcases.CaseV1{
	Title:         "MapCollection int to string",
	ExpectedInput: args.Map{
		"length": 3,
		"first": "v1",
		"last": "v3",
	},
}

var mapCollectionNilSourceTestCase = coretestcases.CaseV1{
	Title:         "MapCollection nil source",
	ExpectedInput: args.Map{"isEmpty": true},
}

var mapCollectionEmptySourceTestCase = coretestcases.CaseV1{
	Title:         "MapCollection empty source",
	ExpectedInput: args.Map{"isEmpty": true},
}

// ==========================================================================
// FlatMapCollection
// ==========================================================================

var flatMapCollectionFlattensTestCase = coretestcases.CaseV1{
	Title:         "FlatMapCollection flattens",
	ExpectedInput: args.Map{"length": 6},
}

var flatMapCollectionNilTestCase = coretestcases.CaseV1{
	Title:         "FlatMapCollection nil",
	ExpectedInput: args.Map{"isEmpty": true},
}

// ==========================================================================
// ReduceCollection
// ==========================================================================

var reduceCollectionSumTestCase = coretestcases.CaseV1{
	Title:         "ReduceCollection sum",
	ExpectedInput: args.Map{"result": 10},
}

var reduceCollectionNilTestCase = coretestcases.CaseV1{
	Title:         "ReduceCollection nil returns initial",
	ExpectedInput: args.Map{"result": 99},
}

var reduceCollectionConcatTestCase = coretestcases.CaseV1{
	Title:         "ReduceCollection string concat",
	ExpectedInput: args.Map{"result": "abc"},
}

// ==========================================================================
// GroupByCollection
// ==========================================================================

var groupByCollectionGroupsTestCase = coretestcases.CaseV1{
	Title:         "GroupByCollection groups",
	ExpectedInput: args.Map{
		"groupCount": 2,
		"evenCount": 3,
		"oddCount": 3,
	},
}

var groupByCollectionNilTestCase = coretestcases.CaseV1{
	Title:         "GroupByCollection nil",
	ExpectedInput: args.Map{"groupCount": 0},
}

// ==========================================================================
// ContainsFunc
// ==========================================================================

var containsFuncFoundTestCase = coretestcases.CaseV1{
	Title:         "ContainsFunc found",
	ExpectedInput: args.Map{"result": true},
}

var containsFuncNotFoundTestCase = coretestcases.CaseV1{
	Title:         "ContainsFunc not found",
	ExpectedInput: args.Map{"result": false},
}

var containsFuncNilTestCase = coretestcases.CaseV1{
	Title:         "ContainsFunc nil",
	ExpectedInput: args.Map{"result": false},
}

// ==========================================================================
// ContainsItem
// ==========================================================================

var containsItemFoundTestCase = coretestcases.CaseV1{
	Title:         "ContainsItem found",
	ExpectedInput: args.Map{"result": true},
}

var containsItemNotFoundTestCase = coretestcases.CaseV1{
	Title:         "ContainsItem not found",
	ExpectedInput: args.Map{"result": false},
}

var containsItemNilTestCase = coretestcases.CaseV1{
	Title:         "ContainsItem nil",
	ExpectedInput: args.Map{"result": false},
}

// ==========================================================================
// IndexOfFunc
// ==========================================================================

var indexOfFuncFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOfFunc found",
	ExpectedInput: args.Map{"index": 1},
}

var indexOfFuncNotFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOfFunc not found",
	ExpectedInput: args.Map{"index": -1},
}

var indexOfFuncNilTestCase = coretestcases.CaseV1{
	Title:         "IndexOfFunc nil",
	ExpectedInput: args.Map{"index": -1},
}

// ==========================================================================
// IndexOfItem
// ==========================================================================

var indexOfItemFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOfItem found",
	ExpectedInput: args.Map{"index": 2},
}

var indexOfItemNotFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOfItem not found",
	ExpectedInput: args.Map{"index": -1},
}

// ==========================================================================
// Distinct
// ==========================================================================

var distinctRemovesDuplicatesTestCase = coretestcases.CaseV1{
	Title:         "Distinct removes duplicates",
	ExpectedInput: args.Map{"length": 3},
}

var distinctNilTestCase = coretestcases.CaseV1{
	Title:         "Distinct nil",
	ExpectedInput: args.Map{"isEmpty": true},
}

var distinctNoDuplicatesTestCase = coretestcases.CaseV1{
	Title:         "Distinct no duplicates",
	ExpectedInput: args.Map{"length": 3},
}

// ==========================================================================
// MapSimpleSlice
// ==========================================================================

var mapSimpleSliceTransformsTestCase = coretestcases.CaseV1{
	Title:         "MapSimpleSlice transforms",
	ExpectedInput: args.Map{"length": 3},
}

var mapSimpleSliceNilTestCase = coretestcases.CaseV1{
	Title:         "MapSimpleSlice nil",
	ExpectedInput: args.Map{"isEmpty": true},
}
