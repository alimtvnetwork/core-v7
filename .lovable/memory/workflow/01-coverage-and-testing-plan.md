# Coverage & Testing Master Plan

## Status: ✅ Complete — All iterations done, baseline established
## Last Updated: 2026-03-29T16:00:00+08:00

## Critical Root Cause Checkpoint
Coverage work has been repeatedly invalidated by assumed APIs, broad unverified coverage-file generation, and skipping the compile-first gate. Do **not** treat newly written coverage files as successful until `./run.ps1 PC` and then `./run.ps1 TC` confirm the result.

See finalized postmortem memory: `.lovable/memory/workflow/completed/02-coverage-remediation-root-cause.md`
See API hallucination root cause: `.lovable/memory/workflow/03-api-hallucination-root-cause.md`
See issue record: `issues/repeated-coverage-remediation-failure-root-cause.md`

---

## TC Run Results (2026-03-16)

- **68 packages**, **1210 files**, **755 at 100%**, **455 below 100%**
- **21 packages at 100%**
- **0 blocked packages** (all compile)
- Detailed file-level plan: `.lovable/memory/workflow/03-coverage-file-level-plan.md`

---

## Remaining Coverage Work

Moved to `plan.md` Priority 1 section. See `plan.md` → "Active Work — Prioritized Backlog" → "Priority 1: Coverage Push".

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
