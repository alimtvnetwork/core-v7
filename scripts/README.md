# scripts/ — PowerShell Module Architecture

## Overview

The `run.ps1` dispatcher (≤200 lines) imports all modules from this directory and routes commands to the appropriate function. Each `.psm1` module is self-contained with PowerShell help documentation.

## Module Reference

| Module | Lines | Purpose | Key Functions |
|---|---|---|---|
| `DashboardUI.psm1` | ~1234 | ANSI box-drawing dashboard, phase tracking, coverage diff | `Initialize-DashboardUI`, `Register-Phase`, `Write-PhaseSummaryBox` |
| `Utilities.psm1` | ~482 | Console output helpers, error extraction, line filtering | `Write-Header`, `Write-Success`, `Write-Fail`, `ParseCompileErrors` |
| `TestLogWriter.psm1` | ~214 | Parse Go test output → structured log files | `Write-TestLogs` |
| `TestRunner.psm1` | ~276 | Go test execution, build checks, git operations | `Invoke-AllTests`, `Invoke-PackageTests`, `Invoke-BuildCheck` |
| `CoveragePreChecks.psm1` | ~128 | Pre-coverage validation (safetest, autofix, bracecheck) | `Invoke-CoveragePreChecks` |
| `CoverageCompileCheck.psm1` | ~273 | Compile checks & per-file split recovery | `Invoke-CoverageCompileCheck`, `Invoke-CoverageSplitRecovery` |
| `CoverageProfileMerger.psm1` | ~120 | Profile merging & missing profile detection | `Merge-CoverageProfiles`, `Find-MissingCoverageProfiles` |
| `CoverageReport.psm1` | ~554 | Report generation (TXT, JSON, HTML, AI button) | `Write-CoverageSummaryReport`, `Write-CoverageJsonReport`, etc. |
| `CoverageRunner.psm1` | ~298 | TC orchestrator — calls all coverage sub-modules | `Invoke-TestCoverage` |
| `PackageCoverage.psm1` | ~171 | Single-package coverage command (TCP) | `Invoke-PackageTestCoverage` |
| `BuildTools.psm1` | ~137 | Build, run, format, vet, tidy, clean | `Invoke-Build`, `Invoke-Format`, `Invoke-Vet` |
| `GoConvey.psm1` | ~57 | Launch browser-based GoConvey test runner | `Invoke-GoConvey` |
| `PreCommitCheck.psm1` | ~357 | Pre-commit API mismatch checker | `Invoke-PreCommitCheck` |
| `Help.psm1` | ~140 | Help display, fail log viewer, integrated tests | `Show-Help`, `Invoke-ShowFailLog` |

## Dependency Graph

```
DashboardUI              (standalone — ANSI/box-drawing)
Utilities                (standalone — basic helpers)
TestLogWriter            (→ Utilities)
TestRunner               (→ Utilities, TestLogWriter)
CoveragePreChecks        (→ Utilities, DashboardUI)
CoverageCompileCheck     (→ Utilities)
CoverageProfileMerger    (standalone)
CoverageReport           (→ Utilities, DashboardUI)
CoverageRunner           (→ all Coverage*, Utilities, TestLogWriter, TestRunner, DashboardUI)
PackageCoverage          (→ CoveragePreChecks, CoverageReport, TestRunner, Utilities)
BuildTools               (→ Utilities)
GoConvey                 (standalone)
PreCommitCheck           (→ Utilities, DashboardUI)
Help                     (→ Utilities, TestLogWriter, TestRunner)
```

## How the Dispatch Works

1. `run.ps1` defines `param($Command, [string[]]$ExtraArgs)`
2. Imports all `.psm1` modules via `Import-Module -Force -DisableNameChecking`
3. A `switch ($Command.ToLower())` routes to the matching function
4. Each command alias maps to exactly one function (e.g., `T`/`-t`/`test` → `Invoke-AllTests`)

## Adding a New Command

1. **Create or extend a module** in `scripts/` with the new function
2. **Export it** via `Export-ModuleMember -Function @('Your-Function')`
3. **Add a switch case** in `run.ps1`:
   ```powershell
   { $_ -in "cmd", "-cmd", "command-name" } { Your-Function $firstExtraArg }
   ```
4. **Update `Show-Help`** in `Help.psm1` with the new command
5. **Add documentation**: `.SYNOPSIS`, `.PARAMETER`, `.EXAMPLE` blocks

## Module Loading

All modules are loaded with `Import-Module -Force -DisableNameChecking`. The `-Force` flag ensures re-imports pick up changes during development. Cross-module calls (e.g., `CoverageRunner` calling `Write-Header` from `Utilities`) work because all modules share the same PowerShell session scope after import.

DashboardUI is always guarded with `Get-Command ... -ErrorAction SilentlyContinue` so `run.ps1` degrades gracefully if the module is missing.
