# Reliability & Failure-Chance Report

## Date: 2026-03-29 (Refreshed)
## Scope: Full spec set for `github.com/alimtvnetwork/core`

---

## Executive Summary

This Go utility framework (`core`) has completed 10+ major phases of modernization, coverage expansion, and quality work. The spec set is **mature and extensive** — 25+ architecture docs, 41 bug audit files, 8 testing guidelines, 18 failing-test analyses, and a full dead-code registry. All major phases (1–8, A–E) are ✅ Done.

**Remaining work** is concentrated in: deprecated API cleanup (S-009), performance benchmarks (S-010), pointer receiver audit (S-012), mutex optimization (S-013), and continued coverage push for low-coverage packages (S-014).

---

## 1. Success Probability Estimates

### By Module Complexity Tier

| Tier | Modules | Success Probability | Assumptions |
|------|---------|:-------------------:|-------------|
| **Simple** (mechanical, well-scoped) | Package README creation, constant renaming, deprecation notices, single-file fixes | **95%** | Clear inputs, minimal cross-file impact, low risk of API mismatch. |
| **Medium** (multi-file, API-aware) | Deprecated API removal (S-009), pointer receiver audit (S-012), benchmark creation (S-010), mutex audit (S-013) | **70–80%** | Requires reading source to verify method signatures and callers. Deprecated removal needs grep across all consumers. Benchmarks need real measurement. |
| **Complex / Agentic** (coverage push, reflection-heavy) | Coverage for `corestr` (3.3%), `coredynamic` (0.9%), `corejson` (45%), `corepayload` (56%) | **40–50%** | **Documented root cause of repeated failure**: AI hallucinates Go API signatures (see `.lovable/memory/workflow/03-api-hallucination-root-cause.md`). Reflection-heavy packages resist test generation. |
| **End-to-End** (full verification pipeline) | Write tests → `./run.ps1 PC` → fix → `./run.ps1 TC` → confirm % | **35–45%** | AI cannot compile or run Go tests in sandbox. Every cycle needs user-side verification. Latency amplifies error accumulation. |

### Key Global Assumptions

1. AI **cannot compile or run Go tests** — all Go work is write-only until user verifies with `./run.ps1 PC` / `TC`.
2. AI **must read source before every test edit** — naming-pattern inference is the #1 root cause of failure.
3. Spec set is extensive but **some cross-references may be stale** after many iterations.
4. The `.release` folder must never be modified by AI.
5. Any code change must bump at least the minor version (excluding `.release` folder).

---

## 2. Failure Map

### 2.1 Where Failures Are Likely

| Module / Workflow | Likelihood | Why | Symptoms |
|---|:---:|---|---|
| **Coverage: `coredynamic`** (57 files, 0.9%) | **VERY HIGH** | 53 files at 0%. Reflection-heavy, complex generics, Dynamic typing. Largest uncovered package. | Massive API mismatch. Tests won't compile. Build cascade blocks other packages. |
| **Coverage: `corestr`** (52 files, 3.3%) | **VERY HIGH** | 42 files at 0%. Collection/Hashmap/Hashset/LinkedList with many data structure methods. | Wrong method signatures, missing type fixtures. |
| **Coverage: `corejson`** (18 files, 45%) | **HIGH** | Serialization/deserialization with generics and reflection. | Type assertion failures, wrong generic parameters. |
| **Coverage: `corepayload`** (23 files, 56%) | **HIGH** | Typed generics, complex collection methods, JSON interop. | Wrong factory function signatures, paging logic errors. |
| **Deprecated API removal** (S-009, 110 functions) | **MEDIUM** | Need to confirm no external consumers. Batch removal risks breaking integrated tests. | Compile errors in test packages that use deprecated functions. |
| **Pointer receiver audit** (S-012, 5224 receivers) | **MEDIUM** | Interface satisfaction may break. Value-receiver migration on types with caching fields causes silent bugs. | Tests pass but runtime behavior regresses (e.g., caching stops working). |
| **Benchmark creation** (S-010) | **LOW** | Straightforward API calls, but AI may hallucinate method signatures for unfamiliar packages. | Compile errors in benchmark files. |

### 2.2 Root Causes of AI Failure (from postmortem)

| Root Cause | Frequency | Mitigation |
|---|:---:|---|
| **API hallucination** — AI invents plausible method names/signatures | Very High | Read source → copy exact signatures → write test |
| **Bulk generation** — one wrong assumption cascades across dozens of call sites | High | One-package-at-a-time gate |
| **Stale spec references** — completed phases re-attempted | Medium | Check plan.md and suggestions tracker before starting |
| **No compile feedback** — errors only discovered when user runs PC/TC | High | Minimize batch size, maximize user verification frequency |
| **Assumed naming patterns** — "similar package = similar API" | Medium | Never infer; always read source file |

### 2.3 How Failures Manifest

1. **Silent compilation failures** — Test files created, user gets wall of errors from `./run.ps1 PC`.
2. **Coverage regression** — Blocked test packages make coverage numbers drop across the board.
3. **API mismatch cascade** — One wrong method signature blocks the entire integrated test package.
4. **Duplicate work** — Stale specs cause AI to re-implement completed phases.
5. **Value-receiver cache bug** — Migrating pointer→value receiver on caching types silently breaks runtime behavior.

---

## 3. Corrective Actions (Prioritized)

| # | Fix | Where | Expected Reliability Gain |
|---|-----|-------|:---:|
| 1 | **Run `./run.ps1 PC` and `TC` before any new work** — Establish current baseline | User action | +15% — Real truth of what compiles and what coverage is |
| 2 | **One-package-at-a-time gate** — Never bulk-generate coverage for HIGH RISK packages | Process rule (enforced in memory) | +15% — Prevents cascade failures |
| 3 | **Method signature inventory** — Before writing tests for any HIGH RISK package, create signature snapshot | New file per package | +10% — Prevents API hallucination |
| 4 | **Version bump discipline** — Any code change bumps minor version; `.release` folder untouched | Process rule | +3% — Prevents version confusion |
| 5 | **Stale spec cleanup** — Mark historical code review reports as archived | `spec/01-app/15-*.md`, `16-*.md`, `17-*.md` | +2% — Reduces confusion for new AI |
| 6 | **External consumer audit** — Grep across auk-go repos before deprecated API removal (S-009) | User action | +5% — Prevents breaking downstream |

---

## 4. Readiness Decision

### Verdict: **READY — with process enforcement** ✅

**Strengths:**
- ✅ 25+ architecture spec files covering every major subsystem
- ✅ 41 bug audit files documenting every found issue with root cause and fix
- ✅ Comprehensive testing guidelines (8 files) with naming conventions, AAA patterns, branch coverage strategy
- ✅ Full dead-code registry (11 packages, all ✅ Closed) justifying every coverage gap
- ✅ Detailed postmortem documenting why coverage work failed and how to prevent it
- ✅ API hallucination root cause documented with 8+ concrete examples
- ✅ Well-structured suggestions tracker with completion archives
- ✅ All major phases (1–8, A–E) completed
- ✅ CI pipeline with lint, test, coverage gate, govulncheck

**Before starting implementation:**
1. **MUST**: Run `./run.ps1 PC` and `./run.ps1 TC` to establish current baseline
2. **MUST**: Follow one-package-at-a-time gate for all coverage work
3. **MUST**: Read source before every test edit — never infer APIs
4. **MUST**: Bump at least minor version on any code change (not `.release`)
5. **SHOULD**: Create method signature inventory for HIGH RISK packages
6. **SHOULD**: Accept that `coredynamic` and `corestr` coverage will take 5-8 sessions each

### Overall Success Rate

| Scenario | Estimate |
|:---|:---:|
| Handed to another AI as-is (no process enforcement) | **55–60%** |
| With process enforcement (read-first, one-at-a-time gate) | **72–78%** |
| With process enforcement + user verification loop (PC/TC after each batch) | **82–88%** |

The gap is: API hallucination (~10%), no compile feedback (~8%), stale references (~3%), bulk generation (~4%).

---

## 5. Spec Coverage Analysis

### What's Well-Specified ✅

| Area | Spec Files | Quality |
|---|---|---|
| Repository architecture | `00-repo-overview.md`, `01-folder-map.md` | Excellent — complete folder tree with descriptions |
| Package-level docs | 12 folders specs, 13+ package READMEs | Good — most packages documented |
| Bug audits | 41 files in `spec/13-app-issues/golang/` | Excellent — root cause, fix, impact for each |
| Testing patterns | 8 files in `spec/testing-guidelines/` | Excellent — naming, assertions, branch coverage |
| Improvement roadmap | `20-improvement-plan.md` (246 lines) | Good — all phases documented |
| Coverage workflow | 11 workflow files, 2 completed | Good — detailed session logs |
| Dead code justification | `dead-code-registry.md` (285 lines, 11 packages) | Excellent — every gap justified |

### What's Weak or Missing ⚠️

| Gap | Risk | Recommendation |
|---|---|---|
| No spec for `coreinstruction/` package beyond source | Low | Create folder spec if modification planned |
| No spec for `mutexbykey/` concurrency model | Medium | Document lock semantics before S-013 |
| No formal API contract spec (method signatures) | High for coverage work | Create signature inventories for HIGH RISK packages |
| `spec/01-app/15-17` code review reports are historical | Low | Mark as archived to prevent confusion |
| No benchmark baselines | Low | Will be addressed by S-010 |

---

## Related Documents

- [Improvement Plan](../../spec/01-app/20-improvement-plan.md)
- [Coverage Remediation Root Cause](../workflow/completed/02-coverage-remediation-root-cause.md)
- [API Hallucination Root Cause](../workflow/03-api-hallucination-root-cause.md)
- [Dead Code Registry](../testing/dead-code-registry.md)
- [Suggestions Tracker](../suggestions/01-suggestions-tracker.md)
- [Plan.md](../../plan.md)
