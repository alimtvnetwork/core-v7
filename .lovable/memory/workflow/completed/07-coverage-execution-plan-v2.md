# Coverage Execution Plan v2

## Status: 🔄 Active
## Created: 2026-03-22
## Data Sources: coverage-summary-14.txt, coverage-10.out, blocked-packages-12.json, failing-tests-17.txt

---

## Phase 0: Fix Blocked Packages & Failing Tests (PRIORITY 1)

### 0A. Blocked Packages (8 packages)

| # | Test Package | Root Cause | Fix Required |
|---|---|---|---|
| 1 | `codestacktests` | `codestack.New.TraceCollection` undefined in Coverage11_test.go | Fix API calls — read source for correct creator |
| 2 | `coredynamictests` | Multiple: `NewSimpleResult` missing arg, `sync` undefined, `m.Add().Add` chaining wrong, `ValueString` pointer method, `j.IsNil` undefined | Fix 4 test files: Coverage29, 31, 32, 33 |
| 3 | `corejsontests` | `corejson.NewResult.ApplyMust` undefined in Coverage35_Seg2 | Fix API call in Coverage35 |
| 4 | `corepayloadtests` | `cloned == nil` type mismatch in Coverage17; `HasUser` pointer method in Coverage19 | Fix 2 test files |
| 5 | `corestrtests` | Multiple: `GetValueByKey`, `AddKeyValues`, `HashmapOptions`, `AsJsonMarshaller` undefined; wrong args to `UsingKeyValueStrings`, `LeftMiddleRightFromSplitN`; `ValidValue.JsonModel` undefined | Fix Coverage34_Remaining_test.go |
| 6 | `coretestcasestests` | `*string` not `*coretests.VerifyTypeOf`; `0` not `corevalidator.Condition` | Fix Coverage10_Iteration6_test.go |
| 7 | `coreteststests` | `ShouldHaveNoError` missing arg; `IsEnable` undefined on SimpleTestCase | Fix Coverage2_Iteration6, Coverage3_Iteration7 |
| 8 | `errcoretests` | `mockLengthGetter` redeclared; `g3` declared unused | Fix Coverage14 rename + Coverage13 use g3 |

### 0B. Failing Tests (7 tests)

| # | Test | Package | Root Cause |
|---|---|---|---|
| 1 | `Test_Cov9_FuncWrap_InvokeFirstAndError` | argstests | nil interface → error cast panic at FuncWrapInvoke.go:135 |
| 2 | `Test_I18_NewRwxWrapper_RwxFullString` | chmodhelpertests | Length assertion: expected 10 got different |
| 3 | `Test_Cov14_DynamicMap_ConvMapStringString_NotFound` | enumimpltests | Expected len:0, actual len:1 |
| 4 | `Test_I13_Invoke_ReturnNilSlice` | reflectmodeltests | nil:false expected nil:true — reflect value not nil |
| 5 | `Test_I13_Invoke_ReturnNilMap` | reflectmodeltests | nil:false expected nil:true — reflect value not nil |
| 6 | `Test_I13_Invoke_ReturnNilInterface` | reflectmodeltests | reflect.Value.Interface on zero Value panic |
| 7 | `Test_QW_Instance_JsonString_Nil` | namevaluetests | nil pointer dereference |

---

## Phase 1: Coverage Gaps — Package Inventory

### Non-internal packages below 100% (ordered by uncovered statements):

| # | Package | Coverage | Total Stmts | Uncovered Stmts | Category | Segments |
|---|---|---|---|---|---|---|
| 1 | `coredata/corestr` | 2.5% | 4,021 | 3,926 | **VERY LARGE** | 20 |
| 2 | `coredata/coredynamic` | 0.8% | 1,704 | 1,687 | **VERY LARGE** | 9 |
| 3 | `coredata/corejson` | 4.6% | 1,602 | 1,527 | **VERY LARGE** | 8 |
| 4 | `coredata/corepayload` | 0% | 1,332 | 1,332 | **VERY LARGE** | 7 |
| 5 | `errcore` | 28.1% | ~1,240 | ~892 | **LARGE** | 5 |
| 6 | `coretests/args` | 61.2% | 1,272 | 513 | **LARGE** | 3 |
| 7 | `coretests/coretestcases` | 30.7% | 167 | 121 | **SMALL** | 1 |
| 8 | `reflectcore/reflectmodel` | 0.8% | 185 | 183 | **SMALL** | 1 |
| 9 | `namevalue` | 0% | ~100 | ~100 | **SMALL** | 1 |
| 10 | `codestack` | 0% | ~200 | ~200 | **SMALL** | 1 |
| 11 | `coreimpl/enumimpl` | 97.5% | 990 | 34 | **MEDIUM** | 1 |
| 12 | `corecmp` | 95.1% | ~500 | ~25 | **SMALL** | 1 |
| 13 | `chmodhelper` | 90.4% | ~300 | ~30 | **SMALL** | 1 |
| 14 | `regexnew` | 87.4% | ~100 | ~13 | **SMALL** | 1 |
| 15 | `coretests` | 83.2% | ~300 | ~50 | **SMALL** | 1 |
| 16 | `coretests/results` | 97.3% | 105 | 4 | **SMALL** | 1 |
| 17 | `corevalidator` | 96.1% | ~200 | ~8 | **SMALL** | 1 |
| 18 | `iserror` | 97.4% | ~100 | ~3 | **SMALL** | 1 |
| 19 | `keymk` | 98.5% | ~100 | ~2 | **SMALL** | 1 |
| 20 | `coremath` | 98.5% | ~100 | ~2 | **SMALL** | 1 |
| 21 | `isany` | 99.4% | ~100 | ~1 | **SMALL** | 1 |
| 22 | `reqtype` | 99.1% | ~100 | ~1 | **SMALL** | 1 |
| 23 | `coreutils/stringutil` | 98% | 362 | 8 | **SMALL** | 1 |
| 24 | `issetter` | 99.6% | ~100 | ~1 | **SMALL** | 1 |
| 25 | `coredata/coregeneric` | 99.8% | 835 | 2 | **SMALL** | 1 |
| 26 | `coredata/corerange` | 99.7% | 499 | 2 | **SMALL** | 1 |
| 27 | `coredata/coreonce` | 99.7% | 506 | 2 | **SMALL** | 1 |
| 28 | `coredata/stringslice` | 99.6% | 378 | 2 | **SMALL** | 1 |
| 29 | `coreversion` | 99.7% | ~100 | ~1 | **SMALL** | 1 |
| 30 | `coretaskinfo` | 99.6% | ~100 | ~1 | **SMALL** | 1 |

**Note:** Internal packages (`reflectinternal`, `convertinternal`, `jsoninternal`, `pathinternal`, `strutilinternal`, `mapdiffinternal`) excluded per policy.

---

## Phase 2: Segmented Execution Plan

### Execution Order (by "next" command):

#### Step 1: Fix all 8 blocked packages
#### Step 2: Fix all 7 failing tests

#### Step 3-22: `coredata/corestr` (20 segments × ~200 stmts)
| Seg | Stmts | Focus |
|-----|-------|-------|
| S01 | 1-200 | First set of source files |
| S02 | 201-400 | ... |
| S03 | 401-600 | ... |
| S04 | 601-800 | ... |
| S05 | 801-1000 | ... |
| S06 | 1001-1200 | ... |
| S07 | 1201-1400 | ... |
| S08 | 1401-1600 | ... |
| S09 | 1601-1800 | ... |
| S10 | 1801-2000 | ... |
| S11 | 2001-2200 | ... |
| S12 | 2201-2400 | ... |
| S13 | 2401-2600 | ... |
| S14 | 2601-2800 | ... |
| S15 | 2801-3000 | ... |
| S16 | 3001-3200 | ... |
| S17 | 3201-3400 | ... |
| S18 | 3401-3600 | ... |
| S19 | 3601-3800 | ... |
| S20 | 3801-3926 | Final segment |

#### Step 23-31: `coredata/coredynamic` (9 segments × ~200 stmts)
| Seg | Stmts |
|-----|-------|
| S01-S09 | 200 stmts each, S09 = remainder |

#### Step 32-39: `coredata/corejson` (8 segments × ~200 stmts)
| Seg | Stmts |
|-----|-------|
| S01-S08 | 200 stmts each, S08 = remainder |

#### Step 40-46: `coredata/corepayload` (7 segments × ~200 stmts)
| Seg | Stmts |
|-----|-------|
| S01-S07 | 200 stmts each, S07 = remainder |

#### Step 47-51: `errcore` (5 segments × ~200 stmts)
| Seg | Stmts |
|-----|-------|
| S01-S05 | ~180 stmts each |

#### Step 52-54: `coretests/args` (3 segments × ~170 stmts)
| Seg | Stmts |
|-----|-------|
| S01-S03 | ~170 stmts each |

#### Step 55: `coretests/coretestcases` (1 segment, 121 uncovered)
#### Step 56: `reflectcore/reflectmodel` (1 segment, 183 uncovered)
#### Step 57: `namevalue` (1 segment, ~100 uncovered)
#### Step 58: `codestack` (1 segment, ~200 uncovered)
#### Step 59: `coreimpl/enumimpl` (1 segment, 34 uncovered)
#### Step 60: `corecmp` (1 segment, ~25 uncovered)
#### Step 61: `chmodhelper` (1 segment, ~30 uncovered)
#### Step 62: `regexnew` (1 segment, ~13 uncovered)
#### Step 63: `coretests` (1 segment, ~50 uncovered)

#### Step 64: Quick-win batch (all packages with <10 uncovered stmts)
- `coretests/results`, `corevalidator`, `iserror`, `keymk`, `coremath`
- `isany`, `reqtype`, `stringutil`, `issetter`
- `coregeneric`, `corerange`, `coreonce`, `stringslice`
- `coreversion`, `coretaskinfo`

---

## Summary

| Category | Count | Total Segments |
|---|---|---|
| Blocked package fixes | 8 | Step 1 |
| Failing test fixes | 7 | Step 2 |
| Very Large (>1000 stmts) | 4 packages | 44 segments |
| Large (500-1000 stmts) | 2 packages | 8 segments |
| Medium/Small (<500 stmts) | 24 packages | 10 segments |
| **TOTAL** | **~64 steps** | |

---

## Rules

1. Each "next" = 1 step (blocked fix, failing fix, or 1 segment)
2. Read source before writing tests — never infer APIs
3. Follow AAA pattern, `Test_Cov{N}_{Method}_{Context}` naming
4. Title: `"{Function} returns {Result} -- {Input Context}"`
5. Verify buildability through reasoning
6. Do not modify production code unless required for blocker fixes
