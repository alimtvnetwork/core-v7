# Plan — Future Work Roadmap

## Last Updated: 2026-03-27T01:45:00+08:00

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
| Phase A (Coverage Stabilization) | ✅ Done | 20-iteration plan executed; all non-internal packages covered |
| Phase B.2 (Spec Reconciliation) | ✅ Done | 9 spec files cleaned |

---

## Phase B: Code Cleanup

### B.1 — Codegen Removal ✅ Done
- **Objective**: Remove deprecated `codegen/` entirely
- **Dependencies**: User runs external audit (`grep` across auk-go repos)
- **Expected outputs**: Deleted `codegen/`, `cmd/main/unitTestGenerator.go`, `tests/integratedtests/codegentests/`, `tests/integratedtests/corepropertytests/`; updated specs
- **Acceptance criteria**: All exit criteria in `spec/01-app/10-codegen-deprecation-plan.md` met
- **Completed**: Removed all codegen packages, consumers, and tests. Updated README, repo overview, folder map, and deprecation plan. Confirmed zero remaining Go imports. Shared types in `coretests/` unaffected.

---

## Phase C: Future Architecture (Low Priority)

### C.1 — Generic Interfaces in `coreinterface/` ✅ Done
- **Objective**: Evaluate `ValueGetter[T]` generic interfaces
- **Dependencies**: None
- **Expected outputs**: Architecture decision doc
- **Acceptance criteria**: Decision documented with rationale
- **Spec reference**: `spec/01-app/20-generic-interfaces-decision.md`
- **Completed**: Added `TypedValueGetter[T]`, `TypedValuesGetter[T]`, `TypedKeyValueGetter[K,V]`. Additive adoption — existing interfaces retained.

### C.2 — `iter` Package Adoption (Go 1.23+) ✅ Done
- **Objective**: Use `iter.Seq` for collection iteration patterns
- **Dependencies**: None
- **Expected outputs**: Prototype in `coregeneric/`
- **Acceptance criteria**: Working iterator pattern with tests
- **Spec reference**: `spec/01-app/11-go-modernization.md`
- **Completed**: Added `All()`, `Values()`, `Backward()` iter.Seq/Seq2 methods to Collection, Hashset, Hashmap, SimpleSlice, LinkedList via dedicated `*Iter.go` files.

### C.3 — CI Pipeline ✅ Done
- **Objective**: Add `golangci-lint`, test coverage, and security scanning
- **Dependencies**: None
- **Expected outputs**: CI config file, lint config
- **Acceptance criteria**: CI runs on push, blocks on failures
- **Completed**: GitHub Actions workflow (`.github/workflows/ci.yml`) with 4 jobs: lint, test+coverage gate (60%), govulncheck, build. `.golangci.yml` updated with gocritic, nilerr, durationcheck, prealloc, gosimple.

### C.4 — Module Splitting ✅ Done
- **Objective**: Evaluate splitting monorepo into focused Go modules
- **Dependencies**: All coverage work complete ✅
- **Expected outputs**: Architecture decision doc
- **Acceptance criteria**: Decision documented with migration path
- **Spec reference**: `spec/01-app/26-module-splitting-decision.md`
- **Completed**: Decision — **keep single module**. High internal coupling (especially `internal/*`, `constants`, `errcore`) makes clean splits impractical. Documented triggers for re-evaluation.

---

## Phase D: Tooling & Runner Improvements

### D.1 — Test Title Audit — All Packages ✅ Done
- **Objective**: Audit all packages for test title consistency
- **Dependencies**: None
- **Acceptance criteria**: All test titles follow `"{Function} returns {Result} -- {Input Context}"` format
- **Completed**: All 15,161 titles across 65 packages now conform. Auto-conformer script applied across all integrated test packages.

### D.2 — Diagnostic Output Regression Tests ✅ Done
- **Objective**: Create snapshot tests for diagnostic output formatting
- **Dependencies**: None
- **Acceptance criteria**: Snapshot tests pass for all formatter outputs
- **Completed**: Created `errcore/Coverage11_DiagnosticSnapshots_test.go` with 25 snapshot tests covering MapMismatchError, LineDiffToString, LineDiff, SliceDiffSummary, GherkinsString, ExpectingRecord, ExpectationMessageDef, Expecting, ExpectingSimple, ExpectingSimpleNoType.

---

## Phase E: Unit Coverage Fix & Test Migration ✅ Done

- **Objective**: Fix all build issues, blocked packages, failing tests; migrate tests to `/tests/integratedtests/`; achieve 100% coverage across all non-internal packages.
- **Registry**: All justified coverage gaps (unreachable/dead code) documented in [`.lovable/memory/testing/dead-code-registry.md`](.lovable/memory/testing/dead-code-registry.md) — 11 packages, ~30-45 lines total, all ✅ Closed.
- **Completed**: All blocked packages resolved, all failing tests fixed, all reachable code paths covered. Remaining gaps are architecturally unreachable (platform-specific, `os.Exit`, exhaustive type switches, compiler-required fallbacks).

---

## Next Task Selection

All roadmap tasks are complete. 🎉

| Phase | Status |
|-------|--------|
| Phases 1–8 (Foundation → Deep Quality) | ✅ Done |
| Error Modernization, Go Modernization | ✅ Done |
| Coverage Stabilization (Phase A) | ✅ Done |
| Code Cleanup (Phase B) | ✅ Done |
| Future Architecture (Phase C) | ✅ Done |
| Tooling & Runner (Phase D) | ✅ Done |
| Unit Coverage Fix & Migration (Phase E) | ✅ Done |
| Dead Code Registry | ✅ Closed (11 packages) |

**Potential future work** (user-initiated):
- Performance benchmarking and optimization
- Additional generic collection types
- Extended CI/CD (release automation, changelog generation)

---

## Process Rules (Mandatory for Any AI)

1. **Read source before every test edit.** Never infer APIs from naming patterns.
2. **One package at a time.** Fix → compile verify → move on.
3. **Do not trust coverage percentages while blockers exist.** Fix blockers first.
4. **Do not report success from edits alone.** Only `./run.ps1 PC` and `./run.ps1 TC` are evidence.
5. **Do not bulk-create coverage suites.** Especially for `errcore`, `corejson`, `corepayload`, `coredynamic`, `corestr`.
6. **Honor naming standards.** Coverage tests: `Test_Cov[N]_{Method}_{Context}`. Titles: `"{Function} returns {Result} -- {Input Context}"`.
7. **Honor project behavior standards.** Vacuous truth (`All*` on empty = true, `Any*` on empty = false), nil-handling, byte-slice clone.
