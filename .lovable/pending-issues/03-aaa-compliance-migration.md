# Pending: AAA Compliance Migration (S-016)

## Status: Open — Near-complete; residual violations are intentional exceptions

## Re-Audit: 2026-04-08 (updated)

### Progress
- **Previous (2026-04-06)**: 33,150 violations across 393 files in 53 packages
- **Current (2026-04-08)**: 1,214 violations across 34 files in 12 packages
- **Resolved**: 31,936 violations (96.3% complete)

### ⚠️ Key Finding: Remaining Violations Are Intentional

All 1,214 remaining violations are in **source packages** (not `tests/integratedtests/`). These are in-package tests that access **unexported symbols** (unexported structs, fields, functions).

**These files CANNOT be migrated** because the project's `check-inpkg-imports.ps1` linter **forbids** importing `coretests/args`, `goconvey`, or `testify` in source packages. Importing heavy test frameworks causes `[setup failed]` errors during instrumented coverage runs (`-coverpkg=./...`).

Using `t.Fatal`/`t.Errorf` is the **correct and only available** assertion pattern for in-package tests.

### Remaining Violations (1,214 — all intentional exceptions)

| # | Package | Violations | Files | Tests | Reason |
|---|---------|-----------|-------|-------|--------|
| 1 | `chmodhelper` | 290 | 6 | 221 | Unexported symbols |
| 2 | `reflectinternal` | 271 | 6 | 157 | Internal package |
| 3 | `args` | 215 | 3 | 130 | Core test framework itself |
| 4 | `regexnew` | 137 | 3 | 102 | Unexported symbols |
| 5 | `reflectmodel` | 114 | 4 | 87 | Unexported symbols |
| 6 | `codestack` | 112 | 2 | 106 | Unexported symbols |
| 7 | `corestr` | 56 | 4 | 56 | Unexported symbols |
| 8 | `errcore` | 9 | 2 | 9 | Unexported symbols |
| 9 | `enumimpl` | 6 | 1 | 6 | Unexported symbols |
| 10 | `stringutil` | 2 | 1 | 2 | Unexported symbols |
| 11 | `coregeneric` | 1 | 1 | 1 | Unexported symbols |
| 12 | `mapdiffinternal` | 1 | 1 | 1 | Internal package |

### Recommendation

**Close S-016 as complete.** The 96.3% migration is the maximum achievable. The remaining 1,214 violations are architecturally constrained — changing them would break the coverage pipeline.

### Clean Stats
- **80 packages** fully AAA-compliant (0 violations)
- **1,071 total test files** scanned
- **43,895 total test functions**

## Full Audit Reference
Previous detailed audit: `.lovable/memory/workflow/04-aaa-compliance-audit.md` (733 lines)
