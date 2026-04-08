# Coverage Execution Plan v3

## Status: 🔄 Active
## Created: 2026-03-22
## Data Sources: coverage-12.out, blocked-packages-14.json, failing-tests-20.txt

---

## Phase 0: Fix Blocked Packages & Failing Tests (PRIORITY 1)

### 0A. Blocked Packages (3 packages)

| # | Test Package | Root Cause | Fix |
|---|---|---|---|
| 1 | `coredynamictests` | `Coverage50:419` — `ts.IsEqual()` needs `*TypeStatus` arg | Add `*coredynamic.TypeStatus` argument |
| 2 | `corejsontests` | `Coverage41:97` — `testStringer` redeclared (also in Coverage32:1544) | Rename to `cov41TestStringer` |
| 3 | `corestrtests` | `Coverage41:142-344` — 8 API signature mismatches | Fix arg order/count for 8 methods |

### 0B. Failing Tests (19 tests across 6 packages)

| # | Test | Package | Root Cause | Fix |
|---|---|---|---|---|
| 1 | `Test_I11_PC_IsEqualItems_NilPC` | corepayloadtests | `pc.IsEqualItems(nil)` passes `[]*PW{nil}` not nil — variadic wraps nil | Change expected to `false` |
| 2 | `Test_I11_NewPW_CastOrDeserializeFrom_Valid` | corepayloadtests | `CastOrDeserializeFrom` re-serializes via JSON; `NameIdCategory` doesn't set Name directly — Name is in inner fields | Check actual Name field path |
| 3 | `Test_CovPL_S1_05_*` | corepayloadtests | `HasAttributes()` returns false — `Create()` may not set Attributes | Check PayloadWrapper.Create() |
| 4 | `Test_CovPL_S1_35_*` | corepayloadtests | Attributes.IsValid/IsInvalid logic mismatch | Check source API |
| 5 | `Test_CovPL_S1_54_*` | corepayloadtests | DeserializeToCollection returns nil | Check serialization/deserialization |
| 6 | `Test_CovPL_S2_61_*` | corepayloadtests | col2.Length() != 1 | Check TypedPayloadCollectionDeserialize |
| 7 | `Test_CovPL_S2_65_*` | corepayloadtests | reflect panic: Elem of invalid type — `TypedPayloadWrapperRecords[D]` | Struct D defined inside func, not exported |
| 8 | `Test_Cov10_VerifyError_WithTypeVerify` | coretestcasestests | VerifyTypeOf.ExpectedInput=[]string but CaseV1.ExpectedInput=string | Fix VerifyTypeOf to match types |
| 9 | `Test_Cov10_GetSinglePage..._NegativePagePanic` | corepayloadtests | Same output contamination from IsEqualItems | Same as #1 |
| 10 | `Test_Cov8_GenericGherkins_ShouldBeEqualMap_NotMap` | coretestcasestests | Sub-test t.Fatalf propagates failure | Use `testing.T{}` instead of `t.Run` |
| 11 | `Test_Cov2_SimpleTestCase_ShouldHaveNoError` | coreteststests | `ShouldBeNil` gets extra comparison value from `it.Expected()` | Remove ExpectedInput or fix test setup |
| 12 | `Test_Cov2_SimpleTestCase_ShouldContains` | coreteststests | ShouldContain arg mismatch | Fix expected to match actual container |
| 13 | `Test_Cov3_BaseTestCase_TypeShouldMatch_WithMismatch` | coreteststests | Sub-test failure propagates | Expected behavior — use separate T |
| 14 | `Test_Cov3_TypesValidationMustPasses_WithError` | coreteststests | Sub-test t.Error propagates | Expected behavior — document intent |
| 15 | `Test_CovEnum_BB11_ExpectingEnumValueError` | enumimpltests | `ExpectingEnumValueError("Invalid", byte(0))` — "Invalid" not a valid enum name | Fix test to use correct enum name |
| 16 | `Test_I13_InvokeError_NilError` | reflectmodeltests | `ReflectValueToAnyValue` panics on zero Value for nil error return | Production bug: needs IsValid() guard |

---

## Phase 1: Coverage Gaps — Accurate Package Inventory

### Non-internal packages below 100% (ordered by uncovered statements):

| # | Package | Total | Covered | Uncov | Pct | Segments |
|---|---|---|---|---|---|---|
| 1 | `coredata/corestr` | 5,761 | 158 | 5,603 | 2.7% | 29 |
| 2 | `coredata/coredynamic` | 2,275 | 19 | 2,256 | 0.8% | 12 |
| 3 | `coredata/corejson` | 2,137 | 100 | 2,037 | 4.7% | 11 |
| 4 | `coredata/corepayload` | 1,654 | 0 | 1,654 | 0.0% | 9 |
| 5 | `reflectcore/reflectmodel` | 253 | 2 | 251 | 0.8% | 2 |
| 6 | `chmodhelper` | 1,638 | 1,480 | 158 | 90.4% | 1 |
| 7 | `coretests/args` | 1,723 | 1,609 | 114 | 93.4% | 1 |
| 8 | `coreimpl/enumimpl` | 1,475 | 1,438 | 37 | 97.5% | 1 |
| 9 | `corevalidator` | 719 | 691 | 28 | 96.1% | 1 |
| 10 | `regexnew` | 223 | 195 | 28 | 87.4% | 1 |
| 11 | `errcore` | 834 | 813 | 21 | 97.5% | 1 |
| 12 | `coretests` | 368 | 351 | 17 | 95.4% | 1 |
| 13 | `codestack` | 501 | 491 | 10 | 98.0% | 1 |
| 14 | `corecmp` | 184 | 175 | 9 | 95.1% | 1 |
| 15 | `coreutils/stringutil` | 445 | 436 | 9 | 98.0% | 1 |
| 16 | `keymk` | 397 | 391 | 6 | 98.5% | 1 |
| 17 | `coretests/results` | 147 | 143 | 4 | 97.3% | 1 |
| 18 | `namevalue` | 188 | 185 | 3 | 98.4% | 1 |
| 19-30 | 12 packages (2 uncov each) | varies | ~99.5% | 1-2 | ~99.7% | 1 each |

**Total: 30 packages, 12,263 uncovered statements, 88 segments**

---

## Phase 2: Segmented Execution Plan

### Execution Order (by "next" command):

#### Step 1: Fix 3 blocked packages ✅ (this iteration)
#### Step 2: Fix 19 failing tests (this iteration + next)

#### Steps 3-31: `coredata/corestr` (29 segments × ~200 stmts)
#### Steps 32-43: `coredata/coredynamic` (12 segments × ~200 stmts)
#### Steps 44-54: `coredata/corejson` (11 segments × ~200 stmts)
#### Steps 55-63: `coredata/corepayload` (9 segments × ~200 stmts)
#### Step 64-65: `reflectcore/reflectmodel` (2 segments × ~125 stmts)

#### Steps 66-76: Medium packages (1 segment each)
- chmodhelper, coretests/args, coreimpl/enumimpl, corevalidator, regexnew
- errcore, coretests, codestack, corecmp, stringutil, keymk

#### Step 77: Quick-win batch (12 packages with ≤4 uncovered stmts)
- coretests/results, namevalue, coregeneric, coreonce, corerange
- stringslice, coretestcases, reqtype, coremath, coretaskinfo
- coreversion, isany, iserror, issetter

---

## Summary

| Category | Count | Total Segments |
|---|---|---|
| Blocked package fixes | 3 | Step 1 |
| Failing test fixes | 19 | Step 2 |
| Very Large (>1000 uncov) | 4 packages | 61 segments |
| Medium (10-250 uncov) | 7 packages | 8 segments |
| Small (<10 uncov) | 19 packages | 19 segments |
| **TOTAL** | **~90 steps** | |

---

## Rules

1. Each "next" = 1-2 segments or fix batch
2. Read source before writing tests — never infer APIs
3. Follow AAA pattern, separate `_testcases.go` and `_test.go`
4. Title: `"{Function} returns {Result} -- {Input Context}"`
5. Verify buildability through reasoning
6. Do not modify production code unless required for blocker fixes
7. Do not touch internal packages
