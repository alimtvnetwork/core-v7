# Suggestions Tracker

## Last Updated: 2026-04-08 (Session 2)

## Convention

- **Location**: `.lovable/memory/suggestions/` — this file for active tracking, `completed/` for archives.
- **File naming**: Single tracker file (`01-suggestions-tracker.md`). Individual completed suggestions archived in `completed/NN-slug.md`.
- **Statuses**: `open` → `inProgress` → `done`
- **Completion handling**: When done, update status here and move detail to `completed/`.

---

## Active Suggestions (Pending)

### S-009: Deprecated API Cleanup
- **suggestionId**: S-009
- **createdAt**: 2026-03-21
- **source**: Lovable (codebase audit)
- **affectedProject**: core
- **description**: Remove or sunset 110 deprecated functions/methods across 30+ files. Largest concentrations: `coreindexes/indexes.go` (21), `core.go` (13), `coredata/corestr/` (15+), `coredata/corejson/` (6+), `coredata/stringslice/` (5+).
- **rationale**: Deprecated functions add API surface confusion and maintenance cost. Generic replacements already exist.
- **proposed change**: Phase approach — (1) audit all 110 deprecated markers, (2) confirm generic replacements exist, (3) remove in batches with compile verification. Detailed plan in `.lovable/memory/workflow/04-s009-deprecated-audit-plan.md`.
- **acceptance criteria**: Zero `// Deprecated:` markers remain (or only those with documented external consumers). `./run.ps1 PC` and `TC` pass.
- **status**: open
- **dependencies**: External consumer audit (user must grep across auk-go repos)
- **completion notes**: Detailed 349-line audit plan created but execution not started.

### S-013: Sync.Mutex → sync.RWMutex Audit
- **suggestionId**: S-013
- **createdAt**: 2026-03-21
- **source**: Lovable (codebase audit)
- **affectedProject**: core
- **description**: 27 `sync.Mutex` usages found. Read-heavy collection types (Collection, Hashmap, Hashset) may benefit from `sync.RWMutex` for concurrent read performance.
- **rationale**: `RWMutex` allows multiple concurrent readers, improving throughput for read-heavy workloads.
- **proposed change**: Audit each mutex usage. Migrate to `RWMutex` where read methods (Get, Contains, Len, IsEmpty) dominate.
- **acceptance criteria**: Identified candidates migrated. Benchmark showing improvement for read-heavy scenarios.
- **status**: open (depends on S-010 benchmarks — now complete)
- **dependencies**: S-010 (✅ done — benchmarks available as baseline)
- **completion notes**: —

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

### S-016: AAA Compliance Migration
- **suggestionId**: S-016
- **createdAt**: 2026-04-06
- **source**: Lovable (AAA audit)
- **affectedProject**: core
- **description**: 33,150 non-compliant assertion calls across 393 files in 53 test packages. Tests should use `args.Map` + `ShouldBeEqual` pattern per project standard.
- **rationale**: Consistent assertion format enables better diagnostics, machine-parseable output, and snapshot testing.
- **proposed change**: Migrate packages batch-by-batch, prioritizing high-violation packages. Full audit in `.lovable/memory/workflow/07-aaa-compliance-audit.md`.
- **acceptance criteria**: All test packages use `args.Map` + `ShouldBeEqual`. `./run.ps1 TC` passes.
- **status**: open
- **dependencies**: None
- **completion notes**: Audit report generated (733 lines). Migration not started.

### S-017: PR Template
- **suggestionId**: S-017
- **createdAt**: 2026-04-08
- **source**: Lovable (CI/CD improvement)
- **affectedProject**: core
- **description**: Add `.github/PULL_REQUEST_TEMPLATE.md` with sections for description, type of change, checklist, and linked issues.
- **rationale**: Standardizes PR submissions and ensures reviewers have context.
- **proposed change**: Create template file with standard sections.
- **acceptance criteria**: Template appears when opening PRs on GitHub.
- **status**: open
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
| S-008 | CI Pipeline Setup | 2026-04-06 | GitHub Actions — lint, test, security, build, release |
| S-010 | Performance Benchmarks | 2026-04-06 | 38 benchmarks across 6 packages |
| S-011 | Missing Package READMEs (10 packages) | 2026-03-21 | All 10 created |
| S-012 | Pointer Receiver Audit | 2026-04-06 | 46 methods migrated |
| S-014 | Coverage Push — All Packages to 100% | 2026-04-06 | 21 packages at 100%, 3 accepted gaps |
| S-008a | CI: Test Summary PR Comment | 2026-04-08 | Auto-comments test results on PRs |
| S-008b | CI: Issue Templates | 2026-04-08 | Bug report + feature request YAML forms |
| S-008c | CI: Release-assets cleanup | 2026-04-08 | Folder deleted, user to add to .gitignore |
| S-008d | CI: Workflow verification | 2026-04-08 | Full pipeline reviewed — test parallelism, PR comments, job gating confirmed |

> Detail files for completed suggestions in `completed/` subfolder.
