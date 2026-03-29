package coregenerictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList.Values() — multi-element iteration (covers LinkedListIter.go L36)
// ══════════════════════════════════════════════════════════════════════════════

var cov5ValuesIterMultiElementTestCase = coretestcases.CaseV1{
	Title: "Values returns all elements -- multi-element list",
	ExpectedInput: args.Map{
		"count":  3,
		"first":  10,
		"second": 20,
		"third":  30,
	},
}

var cov5ValuesIterBreakEarlyTestCase = coretestcases.CaseV1{
	Title: "Values returns partial -- break after first",
	ExpectedInput: args.Map{
		"count": 1,
		"first": 10,
	},
}
