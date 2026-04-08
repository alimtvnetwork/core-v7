# Master Project Plan

## Last Updated: 2026-04-08 (Session 3)

---

## Completed Work

### Phase 1: Test & Coverage Foundation (2026-03-11 — 2026-03-16)
- ✅ Fix 21 failing tests (2 production bugs, 3 test logic bugs, 16 expectation corrections)
- ✅ Diagnostic formatting improvements (MapMismatchError, LineDiff, separator headers)
- ✅ Test title audit (~375+ titles renamed to self-documenting convention)
- ✅ Coverage push batches 1-4 (24 packages to 100%)
- ✅ Coverage prompt generator system (PowerShell-based)
- ✅ Deep clone production bug fix (`corepayload` nil AnyMap)
- ✅ Nil receiver coverage audit (all 9 types verified)
- ✅ Test runner hardening review (3 patterns verified correct)
- ✅ Diagnostic output regression tests (16 snapshot tests)
- ✅ Compile baseline established
- ✅ Value receiver migration (Phase 6)
- ✅ All package READMEs created
- ✅ Codegen removal (S-006)
- ✅ Spec reconciliation (S-007, 9 files fixed)

### Phase 2: CI/CD & Quality (2026-03-18 — 2026-04-08)
- ✅ CI pipeline setup (GitHub Actions: lint, test, security, build, release)
- ✅ GoReleaser configuration (multi-platform builds)
- ✅ PR triggers for CI (opened, synchronize, reopened)
- ✅ Coverage PR comment (auto-updates on push)
- ✅ Test results summary PR comment (pass/fail/skip table)
- ✅ Slack/Discord notification step
- ✅ Branch protection documentation (`docs/BRANCH_PROTECTION.md`)
- ✅ Issue templates (bug report + feature request YAML forms)
- ✅ Release-assets folder cleanup (deleted, user to add to .gitignore)
- ✅ CI workflow full verification (pipeline structure, test parallelism, PR comment logic confirmed correct)

### Phase 3: Performance & Code Quality (2026-04-06)
- ✅ Performance benchmarks (S-010: 38 benchmarks across 6 packages)
- ✅ Pointer receiver audit (S-012: 46 methods migrated)
- ✅ 100% reachable coverage for all 21 non-internal packages (S-014)
- ✅ PowerShell refactor plan complete (15/15 tasks)

---

## Pending Work

### S-009: Deprecated API Cleanup (Open)
- 110 deprecated functions across 30+ files
- Detailed audit plan created (`.lovable/memory/workflow/04-s009-deprecated-audit-plan.md`)
- Blocked on: external consumer audit by user
- **Priority**: Medium

### S-013: sync.Mutex → sync.RWMutex Migration (Open)
- 27 mutex usages to audit
- S-010 benchmarks now available as baseline
- **Priority**: Medium

### S-015: Version Bump Discipline (Permanent Rule)
- Every code change must bump minor version
- `.release` folder must never be modified
- **Priority**: Ongoing

### S-016: AAA Compliance Migration (Open)
- 33,150 non-compliant assertions across 393 files
- Full audit in `.lovable/memory/workflow/07-aaa-compliance-audit.md`
- **Priority**: Low (large effort, no functional impact)

### S-017: PR Template (Open)
- Create `.github/PULL_REQUEST_TEMPLATE.md`
- **Priority**: Low

---

## Key Files & References

| File | Purpose |
|------|---------|
| `.lovable/memory/suggestions/01-suggestions-tracker.md` | Active/completed suggestions |
| `.lovable/memory/workflow/01-coverage-and-testing-plan.md` | Coverage completion record |
| `.lovable/memory/workflow/04-s009-deprecated-audit-plan.md` | Deprecated API audit detail |
| `.lovable/memory/workflow/07-aaa-compliance-audit.md` | AAA compliance audit report |
| `.lovable/memory/testing/dead-code-registry.md` | Accepted unreachable code |
| `.lovable/memory/testing/internal-packages-no-coverage.md` | Internal packages excluded from coverage |
| `.lovable/pending-issues/` | Unresolved issues |
| `.lovable/solved-issues/` | Resolved issues with root cause |
| `.github/workflows/ci.yml` | CI/CD pipeline |
| `.goreleaser.yml` | Release automation config |
| `docs/BRANCH_PROTECTION.md` | Branch protection rules |

---

## Process Rules

1. **Read source before writing tests.** Never assume API signatures.
2. **Compile-first gate.** `./run.ps1 PC` before `TC`.
3. **Version bump on every change.** (S-015)
4. **`.release` folder is off-limits.** Never modify.
5. **Test naming**: `Test_Cov[N]_{Method}_{Context}`
6. **Assertion style**: `args.Map` + `ShouldBeEqual`
