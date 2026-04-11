package reqtypetests

// NOTE: The 2 uncovered statements in reqtype are in unexported functions
// `start()` and `end()` (start.go:8-10, end.go:6-8).
//
// Both contain `if len(reqs) == 0 { return nil }` guards.
// However, the only caller `RangesNotMeet` already checks `len(reqs) == 0`
// and returns early before calling start/end.
//
// These 2 statements are therefore unreachable dead code.
// Documented in spec/13-app-issues/golang/41-unreachable-branches-corejson-coreonce.md.
