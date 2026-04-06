# Coverage & Testing Master Plan

## Status: ✅ FULLY COMPLETE — 100% reachable coverage achieved for all 21 non-internal packages
## Last Updated: 2026-04-06

## Critical Root Cause Checkpoint
Coverage work has been repeatedly invalidated by assumed APIs, broad unverified coverage-file generation, and skipping the compile-first gate. Do **not** treat newly written coverage files as successful until `./run.ps1 PC` and then `./run.ps1 TC` confirm the result.

See finalized postmortem memory: `.lovable/memory/workflow/completed/02-coverage-remediation-root-cause.md`
See API hallucination root cause: `.lovable/memory/workflow/03-api-hallucination-root-cause.md`
See issue record: `issues/repeated-coverage-remediation-failure-root-cause.md`

---

## Final Coverage Results (2026-04-06)

### All 21 Packages at 100% Reachable Coverage

`corecmp`, `codestack`, `corepayload`, `corejson`, `coretests/results`, `reflectmodel`, `coretests`, `corevalidator`, `chmodhelper`, `coredynamic`, `enumimpl`, `errcore`, `corestr`, `coretests/args`, `coretests/coretestcases`, `namevalue`, `stringslice`, `corerange`, `stringutil`, `coreversion`, `coreonce`

### Accepted Unreachable Gaps

| Package | File:Line | Reason |
|---------|-----------|--------|
| `stringutil` | `IsEndsWith.go:37` | `remainingLength < 0` unreachable after prior length check |
| `coreversion` | `hasDeductUsingNilNess.go:20` | Exhaustive nil checks above make this unreachable |
| `coreonce` | `JsonStringMust` error branches | `json.Marshal` cannot fail on simple maps/slices |

### Test Restructuring (Split Recovery)

Four packages were restructured into per-method granular test files to enable parallel failure diagnosis:
- `chmodhelpertests`, `coredynamictests`, `corestrtests`, `corepayloadtests`

---

## Remaining Coverage Work

None. Protocol is in maintenance mode for regression handling only.

---

## Completed Tasks

### 1–7: ✅ All batches, prompt generator, compile baseline, audit, TC run — all done.

See `plan.md` for full completion history.

---

## Process Rules (From Postmortem)

1. **List first, then fix one-by-one.** Regenerate blocked packages before new work.
2. **Read source before every test edit.** Never infer signatures from naming patterns.
3. **Use a package gate.** Fix one package → compile verify → move on.
4. **Do not trust coverage percentages while blockers exist.**
5. **Do not report success from edits alone.** Only `./run.ps1 PC` / `TC` are evidence.
6. **Do not bulk-create coverage suites for unfamiliar packages.**
7. **Honor project behavior standards.** Vacuous truth, nil-handling, byte-slice clone.
8. **Honor naming standards.** `Test_Cov[N]_{Method}_{Context}` format.
