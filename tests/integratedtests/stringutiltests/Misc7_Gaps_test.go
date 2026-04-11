package stringutiltests

// ══════════════════════════════════════════════════════════════════════════════
// Coverage7 — stringutil remaining gaps
//
// Target 1: IsEndsWith.go:37-39 — remainingLength < 0
//   Dead code: line 25-27 already returns false when endsWithLength > basePathLength.
//
// Target 2: replaceTemplate.go:316-341 — UsingNamerMapOptions both branches
//   The map parameter uses unexported interface `namer`, so this function
//   cannot be called from external tests. Needs internal test.
//
// Both targets are either dead code or require internal tests.
// Documented as accepted gaps for external test coverage.
// ══════════════════════════════════════════════════════════════════════════════
