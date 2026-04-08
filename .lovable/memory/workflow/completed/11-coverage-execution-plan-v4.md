# Coverage Execution Plan v4

## Status: 🔄 Active
## Created: 2026-03-23
## Data Source: coverage-17.out

---

## Overview

- **30 packages** below 100% coverage (excluding internal)
- **6,241 uncovered statements** total
- **129 estimated tasks** (segments)

---

## Phase 0: Fix Blocked Packages & Failing Tests (PRIORITY 1)

> Carry forward any unresolved blockers from v3 Plan.
> Check failing-tests-41.txt and blocked-packages.json for current state.

---

## Phase 1: Large Packages (>1000 stmts) — 93 segments

### 1A. `coredata/corestr` — 4.0% → 100% (5,761 stmts, 5,530 uncov, 29 segments)

| Step | Segment | Lines (approx) | Focus |
|------|---------|-----------------|-------|
| S01 | Coverage49 | stmts 1–200 | AllIndividualStrings*, AllIndividualsLength* |
| S02 | Coverage50 | stmts 201–400 | CharCollectionMap core methods |
| S03 | Coverage51 | stmts 401–600 | CharCollectionMap extended |
| S04 | Coverage52 | stmts 601–800 | Collection core |
| S05 | Coverage53 | stmts 801–1000 | Collection extended |
| S06 | Coverage54 | stmts 1001–1200 | Collection iteration/search |
| S07 | Coverage55 | stmts 1201–1400 | Collection JSON/serialize |
| S08 | Coverage56 | stmts 1401–1600 | CollectionPtr methods |
| S09 | Coverage57 | stmts 1601–1800 | HashsetRaw* |
| S10 | Coverage58 | stmts 1801–2000 | Hashmap* |
| S11 | Coverage59 | stmts 2001–2200 | Hashmap extended |
| S12 | Coverage60 | stmts 2201–2400 | KeyValues* |
| S13 | Coverage61 | stmts 2401–2600 | New.* creators |
| S14 | Coverage62 | stmts 2601–2800 | Pair*, SimpleSlice core |
| S15 | Coverage63 | stmts 2801–3000 | SimpleSlice extended |
| S16 | Coverage64 | stmts 3001–3200 | SimpleSlice search/filter |
| S17 | Coverage65 | stmts 3201–3400 | SimpleSlice JSON |
| S18 | Coverage66 | stmts 3401–3600 | SimpleStringOnce*, StringsConcat* |
| S19 | Coverage67 | stmts 3601–3800 | StringsOf* |
| S20 | Coverage68 | stmts 3801–4000 | lazy evaluators |
| S21 | Coverage69 | stmts 4001–4200 | clone/dispose |
| S22 | Coverage70 | stmts 4201–4400 | validators |
| S23 | Coverage71 | stmts 4401–4600 | formatters |
| S24 | Coverage72 | stmts 4601–4800 | reflect helpers |
| S25 | Coverage73 | stmts 4801–5000 | iterators |
| S26 | Coverage74 | stmts 5001–5200 | comparators |
| S27 | Coverage75 | stmts 5201–5400 | edge methods |
| S28 | Coverage76 | stmts 5401–5600 | remaining uncovered |
| S29 | Coverage77 | stmts 5601–5761 | final sweep |

> **Note**: Segments S01–S05 already completed in v3 plan. Resume from S06.

### 1B. `coredata/coredynamic` — 97.7% → 100% (2,289 stmts, 52 uncov, 12 segments)

| Step | Segment | Focus |
|------|---------|-------|
| S01 | Coverage51 | Dynamic.go, DynamicGetters.go uncovered branches |
| S02 | Coverage52 | DynamicReflect.go, DynamicJson.go uncovered |
| S03 | Coverage53 | TypeStatus, LengthOfReflect edge cases |

> Only 52 uncovered stmts — likely 1–3 focused tasks needed.

### 1C. `coredata/corejson` — 95.6% → 100% (2,138 stmts, 95 uncov, 11 segments)

| Step | Segment | Focus |
|------|---------|-------|
| S01 | Coverage42 | BytesCollection uncovered methods |
| S02 | Coverage43 | MapResults, Result edge cases |
| S03 | Coverage44 | Serialize/Deserialize error paths |

> ~95 uncovered stmts — likely 2–3 focused tasks.

### 1D. `coretests/args` — 93.2% → 100% (1,723 stmts, 117 uncov, 9 segments)

| Step | Segment | Focus |
|------|---------|-------|
| S01 | CovArgs01 | Dynamic.go, DynamicFunc.go uncovered branches |
| S02 | CovArgs02 | FuncWrap*, Map edge cases |
| S03 | CovArgs03 | String, Holder, LeftRight methods |

> ~117 uncovered stmts — likely 2–3 focused tasks.

### 1E. `coredata/corepayload` — 95.5% → 100% (1,654 stmts, 75 uncov, 9 segments)

| Step | Segment | Focus |
|------|---------|-------|
| S01 | CovPL_S3 | Attributes nil/edge paths |
| S02 | CovPL_S4 | PayloadWrapper deep-clone, TypedPayload |
| S03 | CovPL_S5 | Collection edge cases |

> ~75 uncovered stmts — likely 2–3 focused tasks.

### 1F. `chmodhelper` — 90.4% → 100% (1,638 stmts, 158 uncov, 9 segments)

| Step | Segment | Focus |
|------|---------|-------|
| S01 | CovChmod01 | CreateDefaultPaths, CreateDirWithFiles |
| S02 | CovChmod02 | DirFilesWithContent, GetRecursivePaths |
| S03 | CovChmod03 | Permission helpers, remaining uncovered |

> ~158 uncovered stmts — likely 2–4 focused tasks. Note: filesystem-dependent tests may need mocking.

### 1G. `coreimpl/enumimpl` — 97.5% → 100% (1,475 stmts, 37 uncov, 8 segments)

| Step | Segment | Focus |
|------|---------|-------|
| S01 | CovEnum02 | Remaining 37 uncovered branches |

> Likely 1 focused task.

### 1H. `coredata/coregeneric` — 99.8% → 100% (1,071 stmts, 2 uncov, 6 segments)

| Step | Segment | Focus |
|------|---------|-------|
| S01 | Quick fix | 2 uncovered statements |

> 1 micro-task.

---

## Phase 2: Medium Packages (300–1000 stmts) — 24 segments

| # | Package | Pct | Stmts | Uncov | Est. Tasks |
|---|---------|-----|-------|-------|------------|
| 1 | `errcore` | 97.5% | 834 | 21 | 1 |
| 2 | `corevalidator` | 96.1% | 719 | 28 | 1–2 |
| 3 | `coredata/coreonce` | 99.7% | 676 | 2 | 1 |
| 4 | `coredata/corerange` | 99.7% | 664 | 2 | 1 |
| 5 | `codestack` | 98.0% | 501 | 10 | 1 |
| 6 | `coredata/stringslice` | 99.6% | 491 | 2 | 1 |
| 7 | `coreutils/stringutil` | 98.0% | 445 | 9 | 1 |
| 8 | `keymk` | 98.5% | 397 | 6 | 1 |
| 9 | `coreversion` | 99.7% | 373 | 1 | 1 |
| 10 | `coretests` | 95.4% | 368 | 17 | 1 |

---

## Phase 3: Small Packages (<300 stmts) — 12 tasks

| # | Package | Pct | Stmts | Uncov | Est. Tasks |
|---|---------|-----|-------|-------|------------|
| 1 | `reflectcore/reflectmodel` | 90.8% | 260 | 24 | 1 |
| 2 | `regexnew` | 87.4% | 223 | 28 | 1 |
| 3 | `corecmp` | 95.1% | 184 | 9 | 1 |
| 4 | `coretests/results` | 97.3% | 147 | 4 | 1 |
| 5 | `coretests/coretestcases` | 99.3% | 270 | 2 | 1 |
| 6 | `namevalue` | 98.4% | 188 | 3 | 1 |
| 7 | `reqtype` | 99.1% | 230 | 2 | 1 |
| 8 | `issetter` | 99.6% | 263 | 1 | 1 |
| 9 | `isany` | 99.4% | 156 | 1 | 1 |
| 10 | `coremath` | 98.5% | 65 | 1 | 1 |
| 11 | `iserror` | 97.4% | 39 | 1 | 1 |
| 12 | `coretaskinfo` | 99.6% | 254 | 1 | 1 |

---

## Execution Order (by "next" command)

| Step | Action | Package | Type |
|------|--------|---------|------|
| 1 | Fix blockers | All blocked/failing | Priority 0 |
| 2 | S06 | `coredata/corestr` | Large segment |
| 3 | S07 | `coredata/corestr` | Large segment |
| 4 | S08 | `coredata/corestr` | Large segment |
| ... | S09–S29 | `coredata/corestr` | Large segments |
| 26 | S01–S03 | `coredata/coredynamic` | Focused |
| 29 | S01–S03 | `coredata/corejson` | Focused |
| 32 | S01–S03 | `coretests/args` | Focused |
| 35 | S01–S03 | `coredata/corepayload` | Focused |
| 38 | S01–S03 | `chmodhelper` | Focused |
| 41 | S01 | `coreimpl/enumimpl` | Quick |
| 42 | Quick fix | `coredata/coregeneric` | Micro |
| 43–52 | Medium packages | 10 packages | 1 each |
| 53–64 | Small packages | 12 packages | 1 each |

**Estimated total steps: ~65 (many large pkg segments combined)**

---

## Rules

1. Each "next" = 1 segment or 1 package task
2. Read source before writing tests — never infer APIs
3. Follow AAA pattern, separate `_testcases.go` and `_test.go`
4. Title: `"{Function} returns {Result} -- {Input Context}"`
5. Verify buildability through reasoning
6. Do not modify production code unless required for blocker fixes
7. Do not touch internal packages
8. Always list remaining tasks after each completion
