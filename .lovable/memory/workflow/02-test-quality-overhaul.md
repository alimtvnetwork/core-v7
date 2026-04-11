# Test Quality Overhaul — Phased Plan

**Status:** Planning complete — awaiting Phase 1 execution
**Created:** 2026-04-08

## Problem Summary

Across 72 packages, test files and functions use meaningless names like `Coverage5_test.go`, `Test_Cov4_...`, `Test_I12_...`. Additionally, ~1,308 raw `t.Fatal`/`t.Error` calls and ~29 `t.Log` calls should be replaced with proper framework assertions (`args.Map` + `ShouldBeEqual` / `CaseV1`).

## Scope

| Category | Count |
|----------|-------|
| Badly-named test files | 725 |
| Badly-named test functions | ~39,662 |
| Raw `t.Fatal`/`t.Error` usages | 1,308 |
| Raw `t.Log`/`t.Logf` usages | 29 |
| Affected packages (integrated tests) | 66 |
| Affected packages (in-package tests) | 6 |
| Affected packages (internal) | ~6 (skip per policy) |

---

## Phase 1: In-Package Tests — ✅ DONE

**Scope:** 6 packages, ~13 files
**Packages:** `chmodhelper`, `coreimpl/enumimpl`, `coreutils/stringutil`, `reflectcore/reflectmodel`, `regexnew`

Tasks:
- [ ] Rename files from `Coverage{N}_...` → `{FunctionOrType}_{Behavior}_test.go`
- [ ] Rename test functions: strip `Cov{N}_`, `I{N}_` prefixes, keep descriptive suffix
- [ ] Replace `t.Fatal`/`t.Error` with `ShouldBeEqual` / `args.Map` where architecturally possible
- [ ] Replace/remove `t.Log` calls
- [ ] Verify all tests pass

---

## Phase 2: Integrated Tests — File/Filesystem Packages — ⬜ TODO

**Scope:** ~12 packages, ~100+ files
**Packages:** `chmodhelpertests`, `chmodinstests`, `chmodclasstypetests`, `fsinternaltests`, `pathinternaltests`, `simplewraptests`, and related

Tasks:
- [ ] Rename files: `Coverage10_Executors_test.go` → `RwxInstructionExecutors_test.go`
- [ ] Rename functions: strip `Cov{N}_` / `Coverage{N}_` prefixes
- [ ] Migrate raw assertions to framework style
- [ ] Verify all tests pass

---

## Phase 3: Integrated Tests — Core Data Packages — ⬜ TODO

**Scope:** ~15 packages, ~150+ files
**Packages:** `corecmptests`, `coredynamictests`, `corestrtests`, `coregenerictests`, `corerangetests`, `coreindexestests`, `coreoncetests`, `corejsontests`, `corecsvtests`, `coreappendtests`, `coreconvertedtests`, `coremathtests`, `coreuniquetests`, `corefuncstests`, `coretaskinfotests`

Tasks:
- [ ] Rename files and functions
- [ ] Migrate assertions
- [ ] Verify all tests pass

---

## Phase 4: Integrated Tests — Testing Framework Packages — ⬜ TODO

**Scope:** ~12 packages, ~100+ files
**Packages:** `argstests`, `coreteststests`, `coretestsargstests`, `coretestsresultstests`, `coretestcasestests`, `resultstests`, `coretesttests`, `coreargstests`, `codestacktests`, `corevalidatortests`, `corecomparatortests`, `conditionaltests`

⚠️ Extra care needed — these packages test the test framework itself.

Tasks:
- [ ] Rename files and functions
- [ ] Migrate assertions
- [ ] Verify all tests pass

---

## Phase 5: Integrated Tests — Type/Enum/Reflect/Converter Packages — ⬜ TODO

**Scope:** ~15 packages, ~150+ files
**Packages:** `anycmptests`, `bytetypetests`, `enumimpltests`, `enumtypetests`, `errcoretests`, `isanytests`, `iserrortests`, `issettertests`, `converterstests`, `convertinternaltests`, `reflectmodeltests`, `reflectinternaltests`, `typesconvtests`, `namevaluetests`, `stringslicetests`

Tasks:
- [ ] Rename files and functions
- [ ] Migrate assertions
- [ ] Verify all tests pass

---

## Phase 6: Remaining Packages + Final Sweep — ⬜ TODO

**Scope:** ~17 packages, ~100+ files
**Packages:** `stringutiltests`, `stringcompareastests`, `strutilinternaltests`, `regexnewtests`, `keymktests`, `msgcreatortests`, `ostypetests`, `reqtypetests`, `mapdiffinternaltests`, `jsoninternaltests`, `csvinternaltests`, `trydotests`, `versionindexestests`, `coreapitests`, `coreinstructiontests`, `corepayloadtests`, `coreversiontests`

Tasks:
- [ ] Rename files and functions
- [ ] Migrate assertions
- [ ] Final sweep: verify zero `Coverage{N}` file/function names remain
- [ ] Verify all tests pass
- [ ] Update memory: mark plan complete

---

## Naming Convention (New Standard)

### File names
- **Pattern:** `{FunctionOrType}_{Behavior}_test.go`
- **Example:** `Coverage5_DeepCoverage_test.go` → `RwxWrapper_WildcardPaths_test.go`

### Function names
- **Pattern:** `Test_{TypeOrFunc}_{Scenario}_{ExpectedOutcome}`
- **Strip:** `Cov{N}_`, `Coverage{N}_`, `I{N}_` prefixes — keep the descriptive part
- **Example:** `Test_Cov3_RwxWrapper_LinuxApplyRecursive_ValidDir` → `Test_RwxWrapper_LinuxApplyRecursive_ValidDir`

### Assertions
- **Replace:** `t.Fatal("expected X")` → `expected.ShouldBeEqual(t, delta, "message", actual)`
- **Replace:** `t.Log("info")` → remove or convert to assertion
- **Use:** `CaseV1` for data-driven tests, `args.Map` for flexible assertions

## Execution Rules
- Each phase: rename files → rename functions → migrate assertions → run tests → verify
- No behavior changes — only cosmetic/structural improvements
- Internal packages skipped per project policy
- As each phase completes, update this file: ⬜ TODO → ✅ DONE