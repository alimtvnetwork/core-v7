# PowerShell Refactor & Spec Update Plan
## Created: 2026-03-31
## Status: ✅ Tasks 1-10 Complete | Tasks 11-15 Remaining

---

## Problem Statement
- `run.ps1` is **2720 lines** — target is **≤200 lines** (dispatch + imports only)
- UI has visual bugs visible in Phase Summary (screenshot reference)
- Specs are outdated after many recent changes (coverage diff, dashboard UI, etc.)
- No README or per-function documentation for scripts/

---

## Task Breakdown (Execute one per "next")

### Task 1: Fix Phase Summary UI Bugs
- Review screenshot: "Compile Check" line has `△` warning icon misaligned
- The phase summary box-drawing may have alignment issues
- Fix in `scripts/DashboardUI.psm1`

### Task 2: Extract Utility Functions Module
Extract from `run.ps1` into `scripts/Utilities.psm1`:
- `Write-Header` (L92)
- `Write-Success` (L101)
- `Write-Fail` (L105)
- `Ensure-TestLogDir` (L112)
- `Filter-TestWarnings` (L117)
- `Merge-UniqueOutputLines` (L123)
- `Filter-BlockedCompileLines` (L139)
- `Extract-BuildErrorLines` (L165)
- `Extract-ExecutionFailureLines` (L190)
- `Extract-RuntimeFailureLines` (L220)
- `Add-BuildErrorsForPackage` (L251)
- `Add-RuntimeFailuresForPackage` (L269)
- `ParseCompileErrors` (L2588)

### Task 3: Extract Test Log Writer Module
Extract into `scripts/TestLogWriter.psm1`:
- `Write-TestLogs` (L286) — this is ~176 lines, the log formatting engine

### Task 4: Extract Test Runner Module
Extract into `scripts/TestRunner.psm1`:
- `Invoke-GoTestAndLog` (L462)
- `Invoke-AllTests` (L557)
- `Invoke-PackageTests` (L580)
- `Open-FailingTestsIfAny` (L545)

### Task 5: Extract Coverage Runner Module
Extract into `scripts/CoverageRunner.psm1`:
- `Invoke-TestCoverage` (L612) — ~1018 lines, the largest function
- `Invoke-PackageTestCoverage` (L1928) — ~254 lines

### Task 6: Extract Build & Tool Commands Module
Extract into `scripts/BuildTools.psm1`:
- `Invoke-GitPull` (L480)
- `Invoke-FetchLatest` (L490)
- `Invoke-BuildCheck` (L498)
- `Invoke-RunMain` (L2205)
- `Invoke-Build` (L2210)
- `Invoke-BuildRun` (L2219)
- `Invoke-Format` (L2227)
- `Invoke-Vet` (L2233)
- `Invoke-Tidy` (L2240)
- `Invoke-Clean` (L2616)

### Task 7: Extract GoConvey & Misc Module
Extract into `scripts/GoConvey.psm1`:
- `Invoke-GoConvey` (L2246)

Extract into `scripts/PreCommitCheck.psm1`:
- `Invoke-PreCommitCheck` (L2272) — ~316 lines

### Task 8: Extract Help & Integrated Tests
Extract into `scripts/Help.psm1`:
- `Show-Help` (L2647)
- `Invoke-ShowFailLog` (L2627)
- `Invoke-IntegratedTests` (L2182)

### Task 9: Extract copyForAI
Extract into `scripts/CopyForAI.psm1`:
- `copyForAI` (L1630) — ~298 lines

### Task 10: Rewrite run.ps1 as Thin Dispatcher
- Keep: param block, module imports, switch dispatch
- Import all modules from `scripts/`
- Target: ≤200 lines

### Task 11: Add Documentation Per Module
For each `.psm1` in `scripts/`:
- Add module-level comment block (`.SYNOPSIS`, `.DESCRIPTION`)
- Add per-function comment blocks (`.SYNOPSIS`, `.PARAMETER`, `.EXAMPLE`)

### Task 12: Create scripts/README.md
- Table of all modules with descriptions
- Dependency graph (which modules depend on which)
- How to add a new command
- How the dispatch works

### Task 13: Update spec/02-tooling/powershell-dashboard-ui.md
- Reflect new modular architecture
- Update §15 integration references to new module paths
- Add section on module loading pattern

### Task 14: Update spec/03-powershell-test-run/ Specs
- Update `01-overview.md` with modular architecture
- Update `08-generic-go-test-coverage-runner.md` with current coverage flow
- Update `09-ai-agent-complete-reference.md` with:
  - Module structure
  - Go syntax validation patterns (`go vet`, brace checking)
  - How to write and run Go tests
  - Coverage workflow end-to-end

### Task 15: Create/Update PowerShell Implementation Spec
- `spec/02-tooling/powershell-implementation.md`
- Document: module loading, error guarding, phase tracking
- Go syntax validation (bracecheck, auto-fixer)
- Go test patterns and coverage generation
- How AI agents should interact with the tooling

---

## Module Dependency Order
```
DashboardUI.psm1      (standalone — ANSI/box-drawing)
Utilities.psm1         (standalone — basic helpers)
TestLogWriter.psm1     (depends on: Utilities)
TestRunner.psm1        (depends on: Utilities, TestLogWriter)
CoverageRunner.psm1   (depends on: Utilities, TestLogWriter, DashboardUI)
BuildTools.psm1        (depends on: Utilities)
GoConvey.psm1          (standalone)
PreCommitCheck.psm1    (depends on: Utilities)
CopyForAI.psm1         (standalone)
Help.psm1              (standalone)
```

---

## Current Function Line Counts (approximate)
| Function | Lines | Target Module |
|---|---|---|
| Invoke-TestCoverage | ~1018 | CoverageRunner |
| Invoke-PreCommitCheck | ~316 | PreCommitCheck |
| copyForAI | ~298 | CopyForAI |
| Invoke-PackageTestCoverage | ~254 | CoverageRunner |
| Write-TestLogs | ~176 | TestLogWriter |
| All utilities combined | ~162 | Utilities |
| All build/tool commands | ~80 | BuildTools |
| Other | ~100 | Various |
