# Solved: Coverage Regression from Build-Failure Cascade

## Resolved: 2026-03-25

## Root Cause
Coverage appeared to degrade because multiple integrated test packages failed to compile, causing broad measurement gaps. Parallel output ordering also made failure patterns harder to read.

### Sub-causes
1. **Build-failure cascade**: Blocked test packages prevent coverage measurement for all cross-referenced source packages.
2. **API signature drift**: Coverage test files written with assumed (incorrect) API signatures caused compile blockers.
3. **Parallel output perception gap**: Async completion without ordered display created impression of randomness.

## Solution
1. Deterministic parallel-sync reporting (plan-indexed, sorted output)
2. Compile-blocker cleanup (verify API signatures from source before writing tests)
3. Package-derived sanitized filenames for artifacts

## Learnings
1. Coverage metrics are meaningless when compile blockers exist.
2. Assumed test signatures are a high-cost failure mode.
3. Parallel execution must use deterministic sync for operator trust.

## What Not To Repeat
- Never write tests against assumed signatures — verify from source first.
- Never stream parallel output directly — always sync-sort before display.
- Never use collision-prone temp names for parallel jobs.
- Always treat blocked-packages count as a quality gate.
