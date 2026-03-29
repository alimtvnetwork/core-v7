package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ══════════════════════════════════════════════════════════════════════════════
// MapStringStringOnce.JsonStringMust() — success path
// Covers MapStringStringOnce.go L306-317 (non-error path already covered,
// error panic path L309-314 is unreachable with valid map data)
// ══════════════════════════════════════════════════════════════════════════════

var cov13MapStringStringOnceJsonStringMustTestCase = coretestcases.CaseV1{
	Title: "JsonStringMust returns valid JSON -- non-empty map",
	ExpectedInput: args.Map{
		"nonEmpty": true,
		"noPanic":  true,
	},
}

// ══════════════════════════════════════════════════════════════════════════════
// StringsOnce.JsonStringMust() — success path
// Covers StringsOnce.go L248-258 (non-error path already covered,
// error panic path L251-256 is unreachable with valid string slice data)
// ══════════════════════════════════════════════════════════════════════════════

var cov13StringsOnceJsonStringMustTestCase = coretestcases.CaseV1{
	Title: "JsonStringMust returns valid JSON -- non-empty strings",
	ExpectedInput: args.Map{
		"nonEmpty": true,
		"noPanic":  true,
	},
}
