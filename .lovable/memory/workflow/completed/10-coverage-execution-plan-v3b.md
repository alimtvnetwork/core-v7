# Coverage Execution Plan v3

## Status: 🔄 Active
## Created: 2026-03-22
## Baseline: 53.8% total (30,411 stmts, 16,368 covered, 14,043 uncovered)

---

## Phase 0: Fix Blockers & Failing Tests (FIRST PRIORITY)

### 0A. Fix 5 Blocked Packages (compilation errors)

| Step | Package | Error Summary | File(s) to Fix |
|------|---------|---------------|----------------|
| 0A.1 | `codestacktests` | `codestack.New.StackTrace.Default` missing args (needs `int, int`) | `Coverage11_test.go` |
| 0A.2 | `coredynamictests` | `EachPageSize` undefined on `PagingInfo`; unused import; non-boolean `if` | `Coverage33_test.go`, `Coverage34_test.go`, `Coverage36_test.go`, `Coverage24_test.go` |
| 0A.3 | `corejsontests` | `corejson.Result` missing methods: `ToSafeBytes`, `JsonStringMust`, `IsValid`, `IsInvalid`, `HandleErr`, `AsJsonStringBinder`, `AsSimpleJsonBinder` | `Coverage38_Seg5_Final_Interfaces_Edge_test.go` |
| 0A.4 | `corepayloadtests` | `CreateUsingBytes` undefined on `newPayloadWrapperCreator` | `Coverage20_Seg3_PW_Attrs_GenericHelpers_test.go` |
| 0A.5 | `corestrtests` | Redeclared funcs; wrong arg types for `LinkedCollection.Strings`, `AddAsyncFuncItems`, `AddCollections`, `SpreadStrings`, undefined methods | `Coverage34_Remaining_test.go` |

### 0B. Fix 18 Failing Tests (across 5 test packages)

| Step | Package | Test | Root Cause |
|------|---------|------|------------|
| 0B.1 | `argstests` | `Test_Cov9_FuncWrap_InvokeError` | Panic: nil interface→error cast in `InvokeError` |
| 0B.2 | `coretestcasestests` | `Test_Cov10_ShouldBeSortedEqual` | `ShouldBeSortedEqual` doesn't sort before compare — expects sorted input |
| 0B.3 | `coretestcasestests` | `Test_Cov10_VerifyError_WithTypeVerify` | `noErr` actual=false, expected=true — wrong expectation |
| 0B.4 | `coretestcasestests` | `Test_Cov8_GenericGherkins_ShouldBeEqualMap_NotMap` (+sub) | Deliberate mismatch test expects `nil` error but gets diff |
| 0B.5 | `coretestcasestests` | `Test_Cov9_CaseV1_ShouldBeSortedEqual` | Same as 0B.2 — unsorted input |
| 0B.6 | `coretestcasestests` | `Test_Cov9_CaseV1_ShouldBeSortedEqualFirst` | Same as 0B.2 — unsorted input |
| 0B.7 | `coreteststests` | `Test_Cov2_SimpleTestCase_ShouldHaveNoError` | `ShouldHaveNoError` given 1 arg but expects 0 |
| 0B.8 | `coreteststests` | `Test_Cov2_SimpleTestCase_ShouldContains` | Container `[]string{}` doesn't contain "world" |
| 0B.9 | `coreteststests` | `Test_Cov3_BaseTestCase_ShouldBeExplicit_Mismatch` | Args must be `string` not `[]string` |
| 0B.10 | `coreteststests` | `Test_Cov3_BaseTestCase_TypeShouldMatch_WithMismatch` | Expected `nil` but got type mismatch error (correct behavior — fix assertion) |
| 0B.11 | `coreteststests` | `Test_Cov3_BaseTestCase_TypesValidationMustPasses_WithError` | Type validation panics as expected — test not catching it |
| 0B.12 | `coreteststests` | `Test_Cov3_DraftType_IsEqual_InnerF1StringCoverage` | `IsEqual` returns false, test expects true |
| 0B.13 | `errcoretests` | `Test_Cov13_RawErrCollection_AddCompiledErrorGetters` | Nil pointer deref — test passes nil `cov13CompiledErrGetter` |
| 0B.14 | `reflectmodeltests` | `Test_I13_Invoke_ReturnNilFunc` | `nil` check returns false for nil func — reflection issue |
| 0B.15 | `reflectmodeltests` | `Test_I13_Invoke_ReturnNilChan` | Same — nil chan not detected as nil |
| 0B.16 | `reflectmodeltests` | `Test_I13_InvokeFirstAndError_Success` | Panic on `reflect.Value.Interface` on zero Value |

---

## Phase 1: Large Packages (>1000 stmts) — Segmented

### 1. `coredata/corestr` — 5,612 uncovered stmts (~29 segments)

| Segment | Lines | Task |
|---------|-------|------|
| Seg-01 | 1–200 | Coverage tests for first 200 uncovered stmts |
| Seg-02 | 201–400 | |
| Seg-03 | 401–600 | |
| Seg-04 | 601–800 | |
| Seg-05 | 801–1000 | |
| Seg-06 | 1001–1200 | |
| Seg-07 | 1201–1400 | |
| Seg-08 | 1401–1600 | |
| Seg-09 | 1601–1800 | |
| Seg-10 | 1801–2000 | |
| Seg-11 | 2001–2200 | |
| Seg-12 | 2201–2400 | |
| Seg-13 | 2401–2600 | |
| Seg-14 | 2601–2800 | |
| Seg-15 | 2801–3000 | |
| Seg-16 | 3001–3200 | |
| Seg-17 | 3201–3400 | |
| Seg-18 | 3401–3600 | |
| Seg-19 | 3601–3800 | |
| Seg-20 | 3801–4000 | |
| Seg-21 | 4001–4200 | |
| Seg-22 | 4201–4400 | |
| Seg-23 | 4401–4600 | |
| Seg-24 | 4601–4800 | |
| Seg-25 | 4801–5000 | |
| Seg-26 | 5001–5200 | |
| Seg-27 | 5201–5400 | |
| Seg-28 | 5401–5612 | |

**Total: 28 segments**

### 2. `coredata/coredynamic` — 2,256 uncovered stmts (~12 segments)

| Segment | Stmts |
|---------|-------|
| Seg-01 | 1–200 |
| Seg-02 | 201–400 |
| Seg-03 | 401–600 |
| Seg-04 | 601–800 |
| Seg-05 | 801–1000 |
| Seg-06 | 1001–1200 |
| Seg-07 | 1201–1400 |
| Seg-08 | 1401–1600 |
| Seg-09 | 1601–1800 |
| Seg-10 | 1801–2000 |
| Seg-11 | 2001–2200 |
| Seg-12 | 2201–2256 |

**Total: 12 segments**

### 3. `coredata/corejson` — 2,038 uncovered stmts (~11 segments)

| Segment | Stmts |
|---------|-------|
| Seg-01 | 1–200 |
| Seg-02 | 201–400 |
| Seg-03 | 401–600 |
| Seg-04 | 601–800 |
| Seg-05 | 801–1000 |
| Seg-06 | 1001–1200 |
| Seg-07 | 1201–1400 |
| Seg-08 | 1401–1600 |
| Seg-09 | 1601–1800 |
| Seg-10 | 1801–2038 |

**Total: 10 segments**

### 4. `coredata/corepayload` — 1,654 uncovered stmts (~9 segments)

| Segment | Stmts |
|---------|-------|
| Seg-01 | 1–200 |
| Seg-02 | 201–400 |
| Seg-03 | 401–600 |
| Seg-04 | 601–800 |
| Seg-05 | 801–1000 |
| Seg-06 | 1001–1200 |
| Seg-07 | 1201–1400 |
| Seg-08 | 1401–1654 |

**Total: 8 segments**

### 5. `coretests/args` — 668 uncovered stmts (~4 segments)

| Segment | Stmts |
|---------|-------|
| Seg-01 | 1–200 |
| Seg-02 | 201–400 |
| Seg-03 | 401–600 |
| Seg-04 | 601–668 |

**Total: 4 segments**

### 6. `coreimpl/enumimpl` — 37 uncovered stmts (1 segment)

### 7. `coredata/coregeneric` — 2 uncovered stmts (1 segment)

---

## Phase 2: Medium Packages (300–1000 stmts uncovered)

| Step | Package | Uncovered | Segments |
|------|---------|-----------|----------|
| 2.1 | `errcore` | 588 | 3 |
| 2.2 | `codestack` | 501 | 3 |
| 2.3 | `chmodhelper` | 158 | 1 |

---

## Phase 3: Small Packages (<300 stmts uncovered)

| Step | Package | Uncovered | Task |
|------|---------|-----------|------|
| 3.1 | `reflectcore/reflectmodel` | 251 | 2 segments |
| 3.2 | `regexnew` | 28 | 1 task |
| 3.3 | `corecmp` | 9 | 1 task |
| 3.4 | `coretests` | 15 | 1 task |
| 3.5 | `corevalidator` | 28 | 1 task |
| 3.6 | `coretests/results` | 4 | 1 task |
| 3.7 | `iserror` | 1 | 1 task |
| 3.8 | `coreutils/stringutil` | 9 | 1 task |
| 3.9 | `namevalue` | 3 | 1 task |
| 3.10 | `coremath` | 1 | 1 task |
| 3.11 | `keymk` | 6 | 1 task |
| 3.12 | `reqtype` | 2 | 1 task |
| 3.13 | `coretests/coretestcases` | 2 | 1 task |
| 3.14 | `isany` | 1 | 1 task |
| 3.15 | `issetter` | 1 | 1 task |
| 3.16 | `coredata/stringslice` | 2 | 1 task |
| 3.17 | `coretaskinfo` | 1 | 1 task |
| 3.18 | `coredata/corerange` | 2 | 1 task |
| 3.19 | `coredata/coreonce` | 2 | 1 task |
| 3.20 | `coreversion` | 1 | 1 task |

---

## Execution Order (Sequential via "next")

| Task# | Description | Phase |
|-------|-------------|-------|
| 1 | Fix 5 blocked packages (0A.1–0A.5) | 0A |
| 2 | Fix 18 failing tests (0B.1–0B.16) | 0B |
| 3 | `coredata/corestr` Seg-01 | 1.1 |
| 4 | `coredata/corestr` Seg-02 | 1.1 |
| 5 | `coredata/corestr` Seg-03 | 1.1 |
| ... | ... (Seg-04 through Seg-28) | 1.1 |
| 30 | `coredata/corestr` Seg-28 | 1.1 |
| 31 | `coredata/coredynamic` Seg-01 | 1.2 |
| ... | ... (through Seg-12) | 1.2 |
| 42 | `coredata/coredynamic` Seg-12 | 1.2 |
| 43 | `coredata/corejson` Seg-01 | 1.3 |
| ... | ... (through Seg-10) | 1.3 |
| 52 | `coredata/corejson` Seg-10 | 1.3 |
| 53 | `coredata/corepayload` Seg-01 | 1.4 |
| ... | ... (through Seg-08) | 1.4 |
| 60 | `coredata/corepayload` Seg-08 | 1.4 |
| 61 | `coretests/args` Seg-01 | 1.5 |
| ... | ... (through Seg-04) | 1.5 |
| 64 | `coretests/args` Seg-04 | 1.5 |
| 65 | `enumimpl` remaining | 1.6 |
| 66 | `coregeneric` remaining | 1.7 |
| 67 | `errcore` Seg-01 | 2.1 |
| 68 | `errcore` Seg-02 | 2.1 |
| 69 | `errcore` Seg-03 | 2.1 |
| 70 | `codestack` Seg-01 | 2.2 |
| 71 | `codestack` Seg-02 | 2.2 |
| 72 | `codestack` Seg-03 | 2.2 |
| 73 | `chmodhelper` remaining | 2.3 |
| 74 | `reflectmodel` Seg-01 | 3.1 |
| 75 | `reflectmodel` Seg-02 | 3.1 |
| 76–95 | Small packages (3.2–3.20) | 3 |

**Total: ~95 tasks**

---

## Summary Statistics

| Category | Packages | Segments/Tasks |
|----------|----------|----------------|
| Blocked packages | 5 | 1 task (fix all) |
| Failing tests | 18 tests in 5 pkgs | 1 task (fix all) |
| Large (>1000 uncovered) | 5 | 62 segments |
| Medium (300–1000) | 3 | 7 segments |
| Small (<300) | 20 | 22 tasks |
| **Grand Total** | **33 packages** | **~93 tasks** |

---

## Rules

1. Internal packages (`internal/*`) — **DO NOT TOUCH**
2. Each "next" = 1 segment or task
3. Read source before writing tests
4. AAA pattern, `Test_Cov{N}_{Method}_{Context}` naming
5. Title format: `"{Function} returns {Result} -- {Input Context}"`
6. No flaky tests, deterministic outcomes
7. Do not modify production code unless required for blocker fixes
