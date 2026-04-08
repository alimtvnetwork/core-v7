# Coverage Stabilization Plan — Detailed File-Level Coverage Map

## Generated: 2026-03-16 from `./run.ps1 TC` results
## Overall: 68 packages, 1210 files, 755 at 100%, 455 below 100%
## Packages at 100%: 21/68

---

## Priority Tiers

Coverage work is organized into tiers by **effort-to-impact ratio**. Tier 1 packages are close to 100% and need only a few test cases. Tier 4 packages are massive (0% avg) and require full test suites.

---

## Tier 1 — Near 100% (95-99%) — Quick Wins

These packages need only a handful of test cases for specific uncovered branches.

### 1.1 `coredata/coreonce` (avg 95.7%, 7 files < 100%)

| File | Coverage | Gap |
|------|----------|-----|
| `StringOnce.go` | 97.0% | ~1-2 branches |
| `MapStringStringOnce.go` | 96.9% | ~1-2 branches |
| `ErrorOnce.go` | 95.8% | ~2 branches |
| `IntegersOnce.go` | 95.2% | ~2-3 branches |
| `StringsOnce.go` | 93.3% | ~3-4 branches |
| `AnyErrorOnce.go` | 86.6% | ~5-6 branches |
| `AnyOnce.go` | 83.6% | ~6-8 branches |

**Estimated effort**: 1 coverage file, ~30 test cases
**Strategy**: Read each file, identify uncovered nil/edge paths, write targeted tests.

### 1.2 `coreimpl/enumimpl` (avg 95.9%, 16 files < 100%)

| File | Coverage | Gap |
|------|----------|-----|
| `newBasicInt16Creator.go` | 98.7% | tiny |
| `newBasicInt8Creator.go` | 98.7% | tiny |
| `BasicInt16.go` | 98.1% | 1 branch |
| `BasicInt32.go` | 98.1% | 1 branch |
| `BasicInt8.go` | 98.1% | 1 branch |
| `BasicUInt16.go` | 98.1% | 1 branch |
| `numberEnumBase.go` | 97.0% | 2 branches |
| `DiffLeftRight.go` | 96.7% | 1-2 branches |
| `BasicString.go` | 95.2% | 2-3 branches |
| `BasicByte.go` | 94.4% | 3 branches |
| `DynamicMap.go` | 92.9% | 2-3 branches |
| `toStringPrintableDynamicMap.go` | 88.9% | 3-4 branches |
| `toHashset.go` | 83.3% | 3-5 branches |
| `newBasicByteCreator.go` | 80.5% | 5-7 branches |
| `newBasicStringCreator.go` | 69.7% | 10+ branches |
| `ConvAnyValToInteger.go` | 60.0% | 10+ branches |

**Estimated effort**: 1-2 coverage files, ~40 test cases

### 1.3 `keymk` (avg 95.6%, 5 files < 100%)

| File | Coverage | Gap |
|------|----------|-----|
| `Key.go` | 96.2% | 1-2 branches |
| `KeyJson.go` | 93.3% | 2-3 branches |
| `KeyCompiler.go` | 91.8% | 3-4 branches |
| `appendAnyItemsWithBaseStrings.go` | 90.0% | 3-4 branches |
| `appendStringsWithBaseAnyItems.go` | 71.4% | 5-7 branches |

**Estimated effort**: 1 coverage file, ~15-20 test cases

### 1.4 `coredata/corerange` (avg 94.3%, 11 files < 100%)

| File | Coverage | Gap |
|------|----------|-----|
| `within.go` | 98.9% | tiny |
| `MinMaxInt16.go` | 98.6% | tiny |
| `MinMaxInt.go` | 98.5% | 1 branch |
| `MinMaxInt8.go` | 95.7% | 2 branches |
| `MinMaxInt64.go` | 95.7% | 2 branches |
| `StartEndInt.go` | 94.3% | 2-3 branches |
| `StartEndSimpleString.go` | 92.0% | 3-4 branches |
| `MinMaxByte.go` | 89.7% | 3-4 branches |
| `RangeByte.go` | 84.0% | 5-6 branches |
| `RangeInt8.go` | 82.8% | 5-7 branches |
| `RangeInt16.go` | 79.3% | 6-8 branches |

**Estimated effort**: 1 coverage file, ~25-30 test cases

### 1.5 `corevalidator` (avg 91.2%, 10 files < 100%)

| File | Coverage | Gap |
|------|----------|-----|
| `SliceValidator.go` | 98.5% | tiny |
| `TextValidator.go` | 97.4% | 1 branch |
| `LinesValidators.go` | 97.0% | 1-2 branches |
| `SliceValidatorVerify.go` | 95.3% | 2 branches |
| `TextValidators.go` | 93.4% | 3 branches |
| `LineValidator.go` | 89.7% | 3-4 branches |
| `SliceValidatorMessages.go` | 88.2% | 4 branches |
| `SliceValidatorAssertions.go` | 80.0% | 5-6 branches |
| `HeaderSliceValidators.go` | 61.5% | 10+ branches |
| `SliceValidators.go` | 41.3% | 15+ branches |

**Estimated effort**: 1-2 coverage files, ~35 test cases

### 1.6 `coredata/stringslice` (avg 90.6%, 24 files < 100%)

| File | Coverage | Gap |
|------|----------|-----|
| `ExpandBySplits.go` | 92.3% | 2 branches |
| `SafeIndexes.go` | 90.9% | 3 branches |
| `AppendAnyItemsWithStrings.go` | 90.9% | 3 branches |
| `LinesSimpleProcessNoEmpty.go` | 88.9% | 3 branches |
| `LinesSimpleProcess.go` | 87.5% | 3-4 branches |
| `MergeSlicesOfSlices.go` | 85.7% | 4 branches |
| `SafeRangeItems.go` | 84.6% | 4-5 branches |
| `AnyItemsCloneUsingCap.go` | 80.0% | 4-5 branches |
| `SafeIndexRanges.go` | 78.6% | 5-6 branches |
| `SplitTrimmedNonEmptyAll.go` | 75.0% | 5-6 branches |
| `SplitTrimmedNonEmpty.go` | 75.0% | 5-6 branches |
| `RegexTrimmedSplitNonEmptyAll.go` | 75.0% | 5-6 branches |
| `CloneSimpleSliceToPointers.go` | 75.0% | 5-6 branches |
| `TrimmedEachWordsPtr.go` | 66.7% | 5-7 branches |
| `SafeRangeItemsPtr.go` | 66.7% | 5-7 branches |
| `NonWhitespacePtr.go` | 66.7% | 5-7 branches |
| `NonWhitespaceJoinPtr.go` | 66.7% | 5-7 branches |
| `NonEmptyJoinPtr.go` | 66.7% | 5-7 branches |
| `SafeIndexAtUsingLastIndex.go` | 66.7% | 5-7 branches |
| `LastSafeIndexPtr.go` | 66.7% | 5-7 branches |
| `ProcessOptionAsync.go` | 0.0% | full file |
| `ProcessAsync.go` | 0.0% | full file |
| `LinesAsyncProcess.go` | 0.0% | full file |
| `AnyLinesProcessAsyncUsingProcessor.go` | 0.0% | full file |

**Note**: 4 async files at 0% may require goroutine testing patterns.
**Estimated effort**: 2 coverage files, ~50 test cases

---

## Tier 2 — Moderate Gap (50-90%) — Focused Effort

### 2.1 `errcore` (avg 90.2%, 15 files < 100%)

Already has Coverage2-9. Need targeted additions for remaining uncovered branches.
**Estimated effort**: 1 coverage file (Coverage10), ~20 test cases

### 2.2 `coredata/corejson` (avg 45.0%, 18 files < 100%)

Major package with serialization/deserialization logic. Many internal files at 0%.
**Estimated effort**: 2-3 coverage files, ~60 test cases
**Risk**: HIGH — complex generics, reflection-based logic

### 2.3 `coredata/corepayload` (avg 56.4%, 23 files < 100%)

Large package with JSON, paging, and collection logic.
**Estimated effort**: 3-4 coverage files, ~80 test cases
**Risk**: HIGH — typed generics, complex collection methods

### 2.4 `reflectcore/reflectmodel` (avg 72.6%, 3 files < 100%)

| File | Coverage |
|------|----------|
| `isNull.go` | 0.0% |
| `MethodProcessor.go` | 81.5% |
| `utils.go` | 81.6% |

**Estimated effort**: 1 coverage file, ~15 test cases

### 2.5 `internal/reflectinternal` (avg 80.4%, 11 files < 100%)

Reflection utilities. Some files at 12-55%.
**Estimated effort**: 2 coverage files, ~40 test cases
**Risk**: MEDIUM — reflection-heavy, may need specific type fixtures

---

## Tier 3 — Major Gap (0-50%) — Large Effort

### 3.1 `coredata/corestr` (avg 3.3%, 52 files < 100%)

Massive package, 42 files at 0%. Includes Collection, Hashmap, Hashset, LinkedList, etc.
**Estimated effort**: 5-8 coverage files, ~150+ test cases
**Risk**: HIGH — largest uncovered package, many data structure methods

### 3.2 `coredata/coredynamic` (avg 0.9%, 57 files < 100%)

53 files at 0%. Dynamic typing, reflection, collection operations.
**Estimated effort**: 5-8 coverage files, ~150+ test cases
**Risk**: VERY HIGH — reflection-heavy, complex generics

### 3.3 `codestack` (avg 0.0%, 11 files at 0%)

Stack trace utilities. All files uncovered.
**Estimated effort**: 1-2 coverage files, ~25 test cases
**Risk**: MEDIUM — runtime-dependent stack traces may need careful testing

### 3.4 `corecmp` (avg 10.8%, 22 files < 100%)

19 files at 0%, mostly type-specific comparison functions.
**Estimated effort**: 2-3 coverage files, ~40 test cases

---

## Tier 4 — Supporting Packages (Not in original 12 target list)

These packages also need coverage but are lower priority:

| Package | Avg | Uncovered Files | Priority |
|---------|-----|-----------------|----------|
| `chmodhelper` | 64.8% | 48 | LOW (filesystem) |
| `coreinstruction` | 63.4% | 14 | MEDIUM |
| `coretests/args` | 34.6% | 30 | LOW (test infra) |
| `coretests/coretestcases` | 82.4% | 6 | LOW (test infra) |
| `coretests` | 70.3% | 11 | LOW (test infra) |

---

## Recommended Execution Order

**Phase 1 — Quick Wins (Tier 1, ~6 batches)**
1. `coreonce` — 7 files, highest avg (95.7%)
2. `keymk` — 5 files
3. `corerange` — 11 files
4. `enumimpl` — 16 files
5. `corevalidator` — 10 files
6. `stringslice` — 24 files (skip async files initially)

**Phase 2 — Moderate Effort (Tier 2, ~5 batches)**
7. `errcore` — 15 files
8. `reflectmodel` — 3 files
9. `reflectinternal` — 11 files
10. `corejson` — 18 files ⚠️ HIGH RISK
11. `corepayload` — 23 files ⚠️ HIGH RISK

**Phase 3 — Heavy Lift (Tier 3, ~8 batches)**
12. `codestack` — 11 files
13. `corecmp` — 22 files
14. `corestr` — 52 files ⚠️ VERY HIGH RISK
15. `coredynamic` — 57 files ⚠️ VERY HIGH RISK

**Mandatory Process Per Batch:**
1. Read ALL source files before writing tests
2. Write coverage file
3. Run `./run.ps1 PC` to verify compilation
4. Fix any mismatches
5. Run `./run.ps1 TC` to measure coverage delta

---

## Summary Statistics

| Tier | Packages | Files < 100% | Est. Test Cases | Est. Batches |
|------|----------|--------------|-----------------|--------------|
| 1 | 6 | 73 | ~195 | 6-8 |
| 2 | 5 | 50 | ~215 | 5-7 |
| 3 | 4 | 142 | ~365+ | 8-12 |
| 4 | 5 | 109 | ~200+ | 5-8 |
| **Total** | **20** | **374** | **~975+** | **24-35** |
