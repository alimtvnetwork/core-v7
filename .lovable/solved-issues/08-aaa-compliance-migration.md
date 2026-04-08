# Solved: AAA Compliance Migration (S-016)

## Status: Closed — Complete (2026-04-08)

## Root Cause
33,150 test assertions used raw `t.Fatal`/`t.Errorf` instead of the project's `args.Map` + `ShouldBeEqual` AAA pattern.

## Solution
- Bulk migration using `scripts/aaa_transform.py`, `aaa_comments.py`, and `args_map_multiline.py`
- 31,936 violations resolved (96.3%)
- 80 packages fully AAA-compliant

## Remaining (Intentional Exceptions)
1,214 violations in 12 source packages (`chmodhelper`, `reflectinternal`, `args`, `regexnew`, `reflectmodel`, `codestack`, `corestr`, `errcore`, `enumimpl`, `stringutil`, `coregeneric`, `mapdiffinternal`).

**Why they cannot be migrated**: These are in-package tests accessing unexported symbols. The `check-inpkg-imports.ps1` linter forbids importing `coretests/args`, `goconvey`, or `testify` in source packages — doing so causes `[setup failed]` errors during instrumented coverage runs (`-coverpkg=./...`).

## Learning
- In-package tests that access unexported symbols are architecturally constrained to stdlib assertions.
- 96.3% is the maximum achievable migration rate given the project's coverage pipeline design.

## What Not to Repeat
- Don't attempt to migrate in-package tests to framework assertions — it will break the coverage pipeline.
