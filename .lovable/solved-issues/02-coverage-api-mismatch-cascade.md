# Solved: Coverage Test API Mismatch Cascade

## Resolved: 2026-04-06

## Root Cause
12 test packages blocked from compiling because coverage test files were generated with **assumed** API signatures instead of verified ones. Methods, parameters, and return types referenced in tests didn't exist in source packages.

### Categories of Mismatch
1. **Wrong parameter count/types** (e.g., `Clone()` vs `Clone(false)`, wrong param types)
2. **Non-existent methods/fields** (e.g., `fw.Name()` vs `fw.Name` field)
3. **Value vs pointer type mismatches** (e.g., nil-checking value types)
4. **Build artifact conflicts** (stale compiled binaries)

## Solution
Fixed each test file to use verified API signatures. Every API call checked against source before writing test code.

## Learnings
1. Never write test code with assumed signatures — always read source first.
2. Never submit coverage test files in bulk without compile verification.
3. Check return types before asserting — value types can't be nil-checked.
4. A non-compiling test blocks the entire package (worse than a failing test).

## Impact
- 12 packages blocked → 0% coverage for cross-referenced packages
- Recovery required individual file-by-file signature verification
