# Pending: AAA Compliance Migration (S-016)

## Status: Open — Mostly complete, residual in-package tests remain

## Re-Audit: 2026-04-08

### Progress
- **Previous (2026-04-06)**: 33,150 violations across 393 files in 53 packages
- **Current (2026-04-08)**: 1,214 violations across 34 files in 12 packages
- **Resolved**: 31,936 violations (96.3% complete)

### Remaining Violations (1,214 total)

All remaining violations are in **source packages** (not `integratedtests/`). These are in-package tests that use unexported symbols and were intentionally skipped during the bulk migration.

| # | Package | Violations | Files | Tests | Notes |
|---|---------|-----------|-------|-------|-------|
| 1 | `chmodhelper` | 290 | 6 | 221 | Uses unexported symbols |
| 2 | `reflectinternal` | 271 | 6 | 157 | Internal package |
| 3 | `args` | 215 | 3 | 130 | Core test framework itself |
| 4 | `regexnew` | 137 | 3 | 102 | Uses unexported symbols |
| 5 | `reflectmodel` | 114 | 4 | 87 | Uses unexported symbols |
| 6 | `codestack` | 112 | 2 | 106 | Uses unexported symbols |
| 7 | `corestr` | 56 | 4 | 56 | Uses unexported symbols |
| 8 | `errcore` | 9 | 2 | 9 | Uses unexported symbols |
| 9 | `enumimpl` | 6 | 1 | 6 | — |
| 10 | `stringutil` | 2 | 1 | 2 | — |
| 11 | `coregeneric` | 1 | 1 | 1 | — |
| 12 | `mapdiffinternal` | 1 | 1 | 1 | Internal package |

### Violation Types
- **t.Fatal/t.Fatalf**: 1,114 (91.8%)
- **t.Error/t.Errorf**: 100 (8.2%)

### Clean Stats
- **80 packages** fully AAA-compliant (0 violations)
- **1,071 total test files** scanned
- **43,895 total test functions**

## Why These Remain

These tests live inside source packages (not `tests/integratedtests/`) because they access unexported symbols. Migrating them to `args.Map` + `ShouldBeEqual` is possible without moving files — only the assertion style needs changing.

## Next Steps
1. Migrate small packages first: `mapdiffinternal` (1), `coregeneric` (1), `stringutil` (2)
2. Then medium: `enumimpl` (6), `errcore` (9), `corestr` (56)
3. Then large: `codestack` (112), `reflectmodel` (114), `regexnew` (137), `args` (215)
4. Finally largest: `reflectinternal` (271), `chmodhelper` (290)
5. Compile + test after each package

## Full Audit Reference
Previous detailed audit: `.lovable/memory/workflow/04-aaa-compliance-audit.md` (733 lines)
