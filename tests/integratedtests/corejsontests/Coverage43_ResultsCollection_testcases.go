package corejsontests

import (
	"errors"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// ResultsCollection — Length / IsEmpty / HasAnyItem / nil receiver
// =============================================================================

var resultsCollectionNilLengthTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection Length returns 0 -- nil receiver",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var resultsCollectionIsEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection IsEmpty returns true -- empty collection",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var resultsCollectionHasAnyItemTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection HasAnyItem returns true -- one item added",
	ExpectedInput: args.Map{
		"hasAny": true,
	},
}

// =============================================================================
// ResultsCollection — FirstOrDefault / LastOrDefault
// =============================================================================

var resultsCollectionFirstOrDefaultEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection FirstOrDefault returns nil -- empty collection",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var resultsCollectionFirstOrDefaultHasItemTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection FirstOrDefault returns non-nil -- one item",
	ExpectedInput: args.Map{
		"isNil": false,
	},
}

var resultsCollectionLastOrDefaultEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection LastOrDefault returns nil -- empty collection",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var resultsCollectionLastOrDefaultHasItemTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection LastOrDefault returns non-nil -- one item",
	ExpectedInput: args.Map{
		"isNil": false,
	},
}

// =============================================================================
// ResultsCollection — Take / Limit / Skip
// =============================================================================

var resultsCollectionTakeEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection Take returns empty -- empty collection",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var resultsCollectionTakeValidTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection Take returns 2 -- take first 2 of 3",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

var resultsCollectionLimitTakeAllTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection Limit returns all -- limit -1 means take all",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

var resultsCollectionSkipEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection Skip returns empty -- empty collection",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var resultsCollectionSkipValidTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection Skip returns 2 -- skip 1 of 3",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

// =============================================================================
// ResultsCollection — Add variants
// =============================================================================

var resultsCollectionAddSkipOnNilNilTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection AddSkipOnNil returns 0 -- nil result skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var resultsCollectionAddSkipOnNilValidTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection AddSkipOnNil returns 1 -- valid result added",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var resultsCollectionAddNonNilNonErrorNilTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection AddNonNilNonError returns 0 -- nil result skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var resultsCollectionAddNonNilNonErrorErrTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection AddNonNilNonError returns 0 -- error result skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var resultsCollectionAddNonNilNonErrorValidTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection AddNonNilNonError returns 1 -- valid result added",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

// =============================================================================
// ResultsCollection — HasError / AllErrors
// =============================================================================

var resultsCollectionHasErrorFalseTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection HasError returns false -- no error items",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var resultsCollectionHasErrorTrueTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection HasError returns true -- one error item",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var resultsCollectionAllErrorsEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection AllErrors returns empty and false -- empty collection",
	ExpectedInput: args.Map{
		"errorCount": 0,
		"hasAnyErr":  false,
	},
}

var resultsCollectionAllErrorsWithErrTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection AllErrors returns 1 error and true -- one error item",
	ExpectedInput: args.Map{
		"errorCount": 1,
		"hasAnyErr":  true,
	},
}

// =============================================================================
// ResultsCollection — GetAtSafe / UnmarshalAt / Paging
// =============================================================================

var resultsCollectionGetAtSafeOutOfRangeTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection GetAtSafe returns nil -- index out of range",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var resultsCollectionUnmarshalAtTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection UnmarshalAt returns hello -- index 0 valid result",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

var resultsCollectionGetPagesSizeZeroTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection GetPagesSize returns 0 -- page size 0",
	ExpectedInput: args.Map{
		"pages": 0,
	},
}

var resultsCollectionGetPagesSizeValidTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection GetPagesSize returns 2 -- 3 items page size 2",
	ExpectedInput: args.Map{
		"pages": 2,
	},
}

// =============================================================================
// ResultsCollection — Clear / Dispose / Clone / Json
// =============================================================================

var resultsCollectionClearTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection Clear returns 0 -- items cleared",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var resultsCollectionClearNilTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection Clear returns nil -- nil receiver",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var resultsCollectionDisposeTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection Dispose sets items nil -- items disposed",
	ExpectedInput: args.Map{
		"nilItems": true,
	},
}

var resultsCollectionCloneTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection Clone returns 1 -- Clone correctly copies items",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var resultsCollectionClonePtrNilTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection ClonePtr returns nil -- nil receiver",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var resultsCollectionJsonTestCase = coretestcases.CaseV1{
	Title: "ResultsCollection Json returns no error -- valid collection",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// =============================================================================
// newResultsCollectionCreator
// =============================================================================

var newResultsCollectionDeserializeInvalidTestCase = coretestcases.CaseV1{
	Title: "NewResultsCollection DeserializeUsingBytes returns error -- invalid bytes",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var newResultsCollectionDeserializeResultErrorTestCase = coretestcases.CaseV1{
	Title: "NewResultsCollection DeserializeUsingResult returns error -- result with error",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

// suppress unused import
var _ = errors.New
