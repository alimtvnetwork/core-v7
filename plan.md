# Plan — Future Work Roadmap

## Last Updated: 2026-03-29T16:00:00+08:00

---

## Status Overview

| Phase | Status | Description |
|-------|--------|-------------|
| Phase 1 (Foundation) | ✅ Done | `interface{}` → `any`, Go 1.24, bug fixes |
| Phase 2 (Generics — Collections) | ✅ Done | Collection[T], Hashset[T], Hashmap[K,V], SimpleSlice[T], LinkedList[T] |
| Phase 3 (Generics — Payload/Dynamic) | ✅ Done | TypedPayloadWrapper[T], TypedDynamic[T], generic deserialize helpers |
| Phase 4 (Test Coverage Expansion) | ✅ Done | `conditional/`, `errcore/`, `converters/` expanded |
| Phase 5 (File Splitting) | ✅ Done | PayloadWrapper, Attributes, Info, Dynamic, BaseTestCase |
| Phase 6 (Value Receiver Migration) | ✅ Done | issetter, coreversion, corepayload; remaining audited |
| Phase 7 (Expert Code Review Fixes) | ✅ Done | 16 findings across 4 sub-phases |
| Phase 8 (Deep Quality Sweep) | ✅ Done | ~190 inline negation refactors, bug fixes, regression tests |
| Error Modernization | ✅ Done | errors.Join, errors.Is/As, fmt.Errorf with %w |
| Go Modernization (Phases 1-7) | ✅ Done | All 7 phases complete including slog, legacy removal |
| Test Title Audit (Batches 1-5) | ✅ Done | ~375+ titles renamed |
| Package READMEs | ✅ Done | All core packages documented |
| Phase A (Coverage Stabilization) | ✅ Done | 20-iteration plan executed |
| Phase B (Code Cleanup) | ✅ Done | Codegen removal, spec reconciliation |
| Phase C (Future Architecture) | ✅ Done | Generic interfaces, iter adoption, CI, module splitting decision |
| Phase D (Tooling & Runner) | ✅ Done | Test title audit, diagnostic regression tests |
| Phase E (Unit Coverage Fix) | ✅ Done | All blocked packages resolved, dead-code registry closed |

---

## Active Work — Prioritized Backlog

### Priority 1: Coverage Push — Remaining Packages (S-014)
- **Objective**: Achieve 100% test coverage for all non-internal packages
- **Dependencies**: User runs `./run.ps1 TC` for current baseline
- **Expected outputs**: New `Coverage*_test.go` files in integrated test packages
- **Acceptance criteria**: All packages at 100% (excluding dead-code registry entries). `./run.ps1 PC` and `TC` pass.
- **Phase**: Ongoing — one package at a time
- **Subpackages by risk**:

| Package | Current Coverage | Risk | Estimated Sessions |
|---------|:---:|:---:|:---:|
| `coreonce` | 95.7% | Low | 1 |
| `keymk` | 95.6% | Low | 1 |
| `corerange` | 94.3% | Low | 1 |
| `enumimpl` | 95.9% | Low | 1 |
| `corevalidator` | 91.2% | Medium | 1 |
| `stringslice` | 90.6% | Medium | 1 |
| `errcore` | 90.2% | Medium | 1 |
| `reflectmodel` | 72.6% | Medium | 1 |
| `reflectinternal` | 80.4% | Medium | 1-2 |
| `corejson` | 45.0% | ⚠️ HIGH | 2-3 |
| `corepayload` | 56.4% | ⚠️ HIGH | 3-4 |
| `codestack` | 0.0% | Medium | 1-2 |
| `corecmp` | 10.8% | Medium | 2 |
| `corestr` | 3.3% | ⚠️ VERY HIGH | 5-8 |
| `coredynamic` | 0.9% | ⚠️ VERY HIGH | 5-8 |

### Priority 2: Deprecated API Cleanup (S-009)
- **Objective**: Remove or sunset 110 deprecated functions/methods across 30+ files
- **Dependencies**: External consumer audit — user must `grep` across auk-go repos
- **Expected outputs**: Removed deprecated functions, updated call sites in tests
- **Acceptance criteria**: Zero `// Deprecated:` markers remain (except documented external consumers). `./run.ps1 PC` and `TC` pass.

### Priority 3: Performance Benchmarks (S-010)
- **Objective**: Add benchmark tests for hot-path operations
- **Dependencies**: None
- **Expected outputs**: `*_bench_test.go` files in 6+ packages, benchmark summary document
- **Acceptance criteria**: ≥30 benchmarks across 6+ packages. Results documented.
- **Subpackages**: `coredata/corestr/Collection`, `coredata/coredynamic`, `errcore`, `codestack`, `regexnew`, `mutexbykey`

### Priority 4: Pointer Receiver Audit (S-012)
- **Objective**: Migrate unnecessary pointer receivers to value receivers for idiomatic Go
- **Dependencies**: None
- **Expected outputs**: Updated method signatures, spec documenting decisions
- **Acceptance criteria**: Identified methods migrated without behavior changes. `./run.ps1 TC` passes.
- **⚠️ Risk**: Types with caching fields MUST keep pointer receivers.

### Priority 5: Sync.Mutex → sync.RWMutex Audit (S-013)
- **Objective**: Improve concurrent read performance for read-heavy collections
- **Dependencies**: S-010 (benchmarks needed to measure improvement)
- **Expected outputs**: Migrated mutex types, before/after benchmark comparison
- **Acceptance criteria**: Measurable improvement in read-heavy benchmark scenarios.

---

## Process Rules (Mandatory for Any AI)

1. **Read source before every test edit.** Never infer APIs from naming patterns.
2. **One package at a time.** Fix → compile verify → move on.
3. **Do not trust coverage percentages while blockers exist.** Fix blockers first.
4. **Do not report success from edits alone.** Only `./run.ps1 PC` and `./run.ps1 TC` are evidence.
5. **Do not bulk-create coverage suites.** Especially for `errcore`, `corejson`, `corepayload`, `coredynamic`, `corestr`.
6. **Honor naming standards.** Coverage tests: `Test_Cov[N]_{Method}_{Context}`. Titles: `"{Function} returns {Result} -- {Input Context}"`.
7. **Honor project behavior standards.** Vacuous truth (`All*` on empty = true, `Any*` on empty = false), nil-handling, byte-slice clone.
8. **Version bump.** Any code change bumps at least minor version. Never modify `.release` folder.
9. **Suggestions tracking.** Every new suggestion goes to `.lovable/memory/suggestions/01-suggestions-tracker.md`. Update status on completion.

---

## Next Task Selection

Pick one of these to implement next:

| # | Task | Risk | Time Estimate |
|---|------|:---:|:---:|
| 1 | Coverage: `coreonce` (95.7% → 100%) | Low | 1 session |
| 2 | Coverage: `keymk` (95.6% → 100%) | Low | 1 session |
| 3 | Coverage: `corerange` (94.3% → 100%) | Low | 1 session |
| 4 | Coverage: `enumimpl` (95.9% → 100%) | Low | 1 session |
| 5 | Coverage: `corevalidator` (91.2% → 100%) | Medium | 1 session |
| 6 | Coverage: `stringslice` (90.6% → 100%) | Medium | 1 session |
| 7 | Deprecated API Cleanup — Phase 1 audit | Medium | 1 session |
| 8 | Performance Benchmarks — first 6 packages | Low-Medium | 1-2 sessions |
| 9 | Coverage: `codestack` (0% → 100%) | Medium | 1-2 sessions |
| 10 | Coverage: `corecmp` (10.8% → 100%) | Medium | 2 sessions |

**Recommendation**: Start with tasks 1–4 (low-risk coverage quick wins) to build momentum, then proceed to medium-risk items.

---

## Memory & Reference Index

| Document | Location | Purpose |
|---|---|---|
| Reliability Risk Report | `.lovable/memory/reports/01-reliability-risk-report.md` | Failure analysis and success estimates |
| Suggestions Tracker | `.lovable/memory/suggestions/01-suggestions-tracker.md` | Active/completed suggestions |
| Coverage Plan | `.lovable/memory/workflow/01-coverage-and-testing-plan.md` | Coverage execution details |
| API Hallucination Root Cause | `.lovable/memory/workflow/03-api-hallucination-root-cause.md` | Why coverage tests fail |
| Dead Code Registry | `.lovable/memory/testing/dead-code-registry.md` | Justified coverage gaps |
| Improvement Plan | `spec/01-app/20-improvement-plan.md` | Full phase-by-phase history |
| Testing Guidelines | `spec/testing-guidelines/` | Naming, assertions, branch coverage |
| Bug Audits | `spec/13-app-issues/golang/` | 41 bug audit files |
