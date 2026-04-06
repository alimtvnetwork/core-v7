# Suggestions Tracker

## Last Updated: 2026-04-06

## Convention

- **Location**: `.lovable/memory/suggestions/` — this file for active tracking, `completed/` for archives.
- **File naming**: Single tracker file (`01-suggestions-tracker.md`). Individual completed suggestions archived in `completed/NN-slug.md`.
- **Statuses**: `open` → `inProgress` → `done`
- **Completion handling**: When done, update status here. Optionally create detail file in `completed/`.
- **Suggestion file content** (when creating individual files):
  - suggestionId, createdAt, source, affectedProject, description, rationale, proposed change, acceptance criteria, status, completion notes

---

## Active Suggestions

### S-009: Deprecated API Cleanup
- **suggestionId**: S-009
- **createdAt**: 2026-03-21
- **source**: Lovable (codebase audit)
- **affectedProject**: core
- **description**: Remove or sunset 110 deprecated functions/methods across 30+ files. Largest concentrations: `coreindexes/indexes.go` (21), `core.go` (13), `coredata/corestr/` (15+), `coredata/corejson/` (6+), `coredata/stringslice/` (5+).
- **rationale**: Deprecated functions add API surface confusion and maintenance cost. Generic replacements already exist.
- **proposed change**: Phase approach — (1) audit all 110 deprecated markers, (2) confirm generic replacements exist, (3) remove in batches with compile verification.
- **acceptance criteria**: Zero `// Deprecated:` markers remain (or only those with documented external consumers). `./run.ps1 PC` and `TC` pass.
- **status**: open
- **dependencies**: External consumer audit (user must grep across auk-go repos)
- **completion notes**: —

### S-010: Performance Benchmarks
- **suggestionId**: S-010
- **createdAt**: 2026-03-21
- **source**: Lovable (codebase audit)
- **affectedProject**: core
- **description**: Add `Benchmark*` tests for hot-path operations. Priority targets: `coredata/corestr/Collection` (Add, Get, Join), `coredata/coredynamic` (type casting), `errcore` (error construction with stack traces), `codestack` (trace capture), `regexnew` (lazy compile), `mutexbykey` (lock contention).
- **rationale**: No performance baseline exists. Regressions are invisible without benchmarks.
- **proposed change**: Create `benchmark_test.go` files in priority packages.
- **acceptance criteria**: ≥30 benchmarks across 6+ packages. Results documented in benchmark summary.
- **status**: **done** (completed 2026-04-06)
- **dependencies**: None
- **completion notes**: 38 benchmarks created across 6 packages: corestr (13), coredynamic (13), errcore (5), codestack (6), regexnew (7), mutexbykey (4). Files: `benchmark_test.go` in each package.

### S-012: Pointer Receiver Audit
- **suggestionId**: S-012
- **createdAt**: 2026-03-21
- **source**: Lovable (codebase audit)
- **affectedProject**: core
- **description**: 5,224 pointer receivers vs 2,836 value receivers. Many small readonly methods (getters, checkers, formatters) on immutable types likely use pointer receivers unnecessarily.
- **rationale**: Value receivers are idiomatic for small, read-only types. They enable better compiler optimizations and prevent nil-receiver panics.
- **proposed change**: Audit top packages (`coredata/corestr`, `errcore`, `coredata/corepayload`) for methods that could safely use value receivers.
- **acceptance criteria**: Identified methods migrated without behavior changes. `./run.ps1 TC` passes.
- **status**: **done** (completed 2026-04-06)
- **dependencies**: None
- **completion notes**: 46 methods migrated to value receivers. LeftRight (18), LeftMiddleRight (22), ExpectingRecord (6). corepayload skipped — all methods require nil-guard pointer receivers. Types embedding sync.Mutex correctly excluded.

### S-013: Sync.Mutex → sync.RWMutex Audit
- **suggestionId**: S-013
- **createdAt**: 2026-03-21
- **source**: Lovable (codebase audit)
- **affectedProject**: core
- **description**: 27 `sync.Mutex` usages found. Read-heavy collection types (Collection, Hashmap, Hashset) may benefit from `sync.RWMutex` for concurrent read performance.
- **rationale**: `RWMutex` allows multiple concurrent readers, improving throughput for read-heavy workloads.
- **proposed change**: Audit each mutex usage. Migrate to `RWMutex` where read methods (Get, Contains, Len, IsEmpty) dominate.
- **acceptance criteria**: Identified candidates migrated. Benchmark showing improvement for read-heavy scenarios.
- **status**: open (depends on S-010 for benchmarks)
- **dependencies**: S-010 (benchmarks needed to measure improvement)
- **completion notes**: —

### S-014: Coverage Push — Remaining Packages
- **suggestionId**: S-014
- **createdAt**: 2026-03-21
- **source**: Lovable (carried from prior coverage work)
- **affectedProject**: core
- **description**: ~~Continue coverage push for packages below 100%.~~ All 21 non-internal packages now at 100% reachable coverage.
- **rationale**: Coverage gaps hide bugs, especially in high-risk packages.
- **proposed change**: Completed. Maintenance mode only.
- **acceptance criteria**: All packages at 100% coverage (excluding dead-code registry entries). ✅ Met.
- **status**: **done** (completed 2026-04-06)
- **dependencies**: None remaining
- **completion notes**: 21 packages at 100%. 3 accepted unreachable gaps documented. 4 packages restructured via split recovery.

### S-015: Version Bump Discipline
- **suggestionId**: S-015
- **createdAt**: 2026-03-29
- **source**: User instruction
- **affectedProject**: core
- **description**: Any code change must bump at least the minor version everywhere except the `.release` folder which must never be modified.
- **rationale**: User requirement for version tracking discipline.
- **proposed change**: Enforce version bump check on every code modification session.
- **acceptance criteria**: Every code-changing session includes a version bump. `.release` folder never touched.
- **status**: open (permanent process rule)
- **dependencies**: None
- **completion notes**: —

---

## Completed Suggestions (Archive)

| # | Title | Completed | Notes |
|---|-------|-----------|-------|
| 1 | Diagnostic Formatting Improvements | 2026-03-11 | 4-space indent, separator headers, tab-indented entries |
| 2 | Test Title Audit (Batches 1-5) | 2026-03-16 | ~375+ titles renamed across all listed packages |
| 3 | Fix 21 Failing Tests | 2026-03-11 | All fixed |
| 4 | Coverage Push Batch 1 (11 packages) | 2026-03-14 | Packages 75-97% |
| 5 | Coverage Push Batch 2 (6 packages) | 2026-03-14 | Packages 0-57% |
| 6 | Coverage Push Batch 3 (7 packages) | 2026-03-15 | Generic/utility packages |
| 7 | Coverage Prompt Generator System | 2026-03-15 | PowerShell-based prompt generation |
| 8 | Deep Clone Production Bug Fix | 2026-03-15 | `corepayload` nil AnyMap |
| 9 | Nil Receiver Coverage Audit | 2026-03-15 | All types audited |
| 10 | Test Runner Hardening Review | 2026-03-15 | Verified |
| 11 | Diagnostic Output Regression Tests | 2026-03-15 | Snapshot tests |
| 12 | Coverage Push Batch 4 (6 packages) | 2026-03-16 | Verified |
| 13 | Value Receiver Migration (Phase 6) | 2026-03-16 | All convertible methods migrated |
| 14 | Remaining Package READMEs | 2026-03-16 | All packages have READMEs |
| 15 | High-Risk Coverage File Audit (6 files) | 2026-03-16 | Audited, 1 fix |
| S-001 | Compile Baseline | 2026-03-16 | Completed |
| S-002 | Verify Batch 4 | 2026-03-16 | Completed |
| S-006 | Codegen Removal | 2026-03-21 | Fully removed |
| S-007 | Spec Reconciliation | 2026-03-17 | 9 files fixed |
| S-008 | CI Pipeline Setup | 2026-03-18 | GitHub Actions |
| S-010 | Performance Benchmarks | 2026-04-06 | 38 benchmarks across 6 packages (corestr, coredynamic, errcore, codestack, regexnew, mutexbykey) |
| S-011 | Missing Package READMEs (10 packages) | 2026-03-21 | All 10 created |
| S-012 | Pointer Receiver Audit | 2026-04-06 | 46 methods migrated (LeftRight, LeftMiddleRight, ExpectingRecord). corepayload skipped (nil guards). |
| S-014 | Coverage Push — All Packages to 100% | 2026-04-06 | 21 packages at 100%, 3 accepted gaps, 4 split-recovery packages |

> Detail files in `completed/` subfolder.
