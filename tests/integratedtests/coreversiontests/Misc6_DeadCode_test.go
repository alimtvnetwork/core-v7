package coreversiontests

// ══════════════════════════════════════════════════════════════════════════════
// Coverage6 — hasDeductUsingNilNess dead code gap
//
// Target: coreversion/hasDeductUsingNilNess.go:20-22
//
//	if left == nil || right == nil {
//	    return corecomparator.NotEqual, true
//	}
//
// This branch is LOGICALLY UNREACHABLE:
//   - Line 8:  both nil   → return Equal
//   - Line 12: left!=nil, right==nil → return LeftGreater
//   - Line 16: left==nil, right!=nil → return LeftLess
//
// After these three checks, both left and right are guaranteed non-nil.
// The guard at line 20 can never be true. This is a defensive dead-code guard.
//
// No test can reach this line without modifying production code.
// Documented as accepted dead code — 0 lines reachable but uncovered.
// ══════════════════════════════════════════════════════════════════════════════
