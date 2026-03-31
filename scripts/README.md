# scripts/ — PowerShell Module Architecture

## Overview

The `run.ps1` dispatcher (≤200 lines) imports all modules from this directory and routes commands to the appropriate function. Each `.psm1` module is self-contained with PowerShell help documentation.

## Module Reference

| Module | Purpose | Key Functions |
|---|---|---|
| `DashboardUI.psm1` | ANSI box-drawing dashboard, phase tracking, coverage diff | `Initialize-DashboardUI`, `Register-Phase`, `Write-PhaseSummaryBox`, `Write-CoverageDiffBox` |
| `Utilities.psm1` | Console output helpers, error extraction, line filtering | `Write-Header`, `Write-Success`, `Write-Fail`, `ParseCompileErrors` |
| `TestLogWriter.psm1` | Parse Go test output → structured log files | `Write-TestLogs` |
| `TestRunner.psm1` | Go test execution, build checks, git operations | `Invoke-AllTests`, `Invoke-PackageTests`, `Invoke-BuildCheck` |
| `CoverageRunner.psm1` | Full and per-package coverage pipelines | `Invoke-TestCoverage` (TC), `Invoke-PackageTestCoverage` (TCP) |
| `BuildTools.psm1` | Build, run, format, vet, tidy, clean | `Invoke-Build`, `Invoke-Format`, `Invoke-Vet` |
| `GoConvey.psm1` | Launch browser-based GoConvey test runner | `Invoke-GoConvey` |
| `PreCommitCheck.psm1` | Pre-commit API mismatch checker | `Invoke-PreCommitCheck` |
| `Help.psm1` | Help display, fail log viewer, integrated tests | `Show-Help`, `Invoke-ShowFailLog`, `Invoke-IntegratedTests` |

## Dependency Graph

```
DashboardUI          (standalone — ANSI/box-drawing)
Utilities            (standalone — basic helpers)
TestLogWriter        (→ Utilities)
TestRunner           (→ Utilities, TestLogWriter)
CoverageRunner       (→ Utilities, TestLogWriter, TestRunner, DashboardUI)
BuildTools           (→ Utilities)
GoConvey             (standalone)
PreCommitCheck       (→ Utilities, DashboardUI)
Help                 (→ Utilities, TestLogWriter, TestRunner)
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
