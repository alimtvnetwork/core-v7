# ─────────────────────────────────────────────────────────────────────────────
# TestRunner.psm1 — Go test execution, build checks, and git operations
#
# Provides: test invocation with log writing, build pre-checks,
#           git pull, dependency fetch, and failing-test file opener.
#
# Usage:
#   Import-Module ./scripts/TestRunner.psm1 -Force
#
# Dependencies:
#   - Utilities.psm1    (Write-Header, Write-Success, Write-Fail,
#                         Ensure-TestLogDir, Extract-BuildErrorLines,
#                         Filter-TestWarnings)
#   - TestLogWriter.psm1 (Write-TestLogs)
# ─────────────────────────────────────────────────────────────────────────────

# ═══════════════════════════════════════════════════════════════════════════════
# §1  Git & Dependency Operations
# ═══════════════════════════════════════════════════════════════════════════════

function Invoke-GitPull {
    <#
    .SYNOPSIS
        Pull the latest changes from the remote git repository.
    .DESCRIPTION
        Runs `git pull` and reports success/failure. Continues on failure
        so that tests can still run against the local state.
    .EXAMPLE
        Invoke-GitPull
    #>
    [CmdletBinding()]
    param()

    Write-Header "Pulling latest from remote"
    $prevPref = $ErrorActionPreference
    $ErrorActionPreference = "Continue"
    git pull 2>&1 | ForEach-Object { Write-Host "  $_" -ForegroundColor Gray }
    if ($LASTEXITCODE -eq 0) { Write-Success "Git pull complete" }
    else { Write-Fail "git pull failed (continuing anyway)" }
    $ErrorActionPreference = $prevPref
}

function Invoke-FetchLatest {
    <#
    .SYNOPSIS
        Pull git changes and run `go mod tidy` to sync dependencies.
    .EXAMPLE
        Invoke-FetchLatest
    #>
    [CmdletBinding()]
    param()

    Invoke-GitPull
    Write-Header "Fetching latest dependencies"
    go mod tidy
    if ($LASTEXITCODE -eq 0) { Write-Success "Dependencies up to date" }
    else { Write-Fail "go mod tidy failed" }
}

# ═══════════════════════════════════════════════════════════════════════════════
# §2  Build Check
# ═══════════════════════════════════════════════════════════════════════════════

function Invoke-BuildCheck {
    <#
    .SYNOPSIS
        Compile-check a Go package path before running tests.
    .DESCRIPTION
        Runs `go build` on the given path. On failure, writes build errors
        to the failing-tests log and opens it. Returns $true/$false.
    .PARAMETER buildPath
        The Go package path to build (e.g., "./..." or "./integratedtests/pkg/...").
    .EXAMPLE
        if (-not (Invoke-BuildCheck "./...")) { return }
    #>
    [CmdletBinding()]
    param([string]$buildPath)

    Write-Header "Build check: $buildPath"
    $prevPref = $ErrorActionPreference
    $ErrorActionPreference = "Continue"
    $output = & go build $buildPath 2>&1 | ForEach-Object { $_.ToString() }
    $exitCode = $LASTEXITCODE
    $ErrorActionPreference = $prevPref

    if ($exitCode -ne 0) {
        Write-Fail "Build failed — skipping tests"

        Ensure-TestLogDir
        $failingFile = Join-Path $TestLogDir "failing-tests.txt"
        $rawFile     = Join-Path $TestLogDir "raw-output.txt"
        $timestamp   = Get-Date -Format "yyyy-MM-dd HH:mm:ss"

        $buildErrors = Extract-BuildErrorLines $output
        $errorCount = if ($buildErrors.Count -gt 0) { $buildErrors.Count } else { 1 }

        $failingContent = @(
            "# Failing Tests — $timestamp",
            "# Count: $errorCount",
            "",
            "# Build Failed — tests were NOT run",
            "",
            "# ── Build Errors ──",
            ""
        )
        if ($buildErrors.Count -gt 0) {
            $failingContent += $buildErrors
        }
        else {
            $failingContent += $output
        }

        Set-Content -Path $failingFile -Value ($failingContent -join "`n") -Encoding UTF8
        Set-Content -Path $rawFile -Value ($output -join "`n") -Encoding UTF8

        $output | ForEach-Object { Write-Host "  $_" -ForegroundColor Red }
        Open-FailingTestsIfAny
        return $false
    }

    Write-Success "Build OK"
    return $true
}

# ═══════════════════════════════════════════════════════════════════════════════
# §3  Test Execution
# ═══════════════════════════════════════════════════════════════════════════════

function Invoke-GoTestAndLog {
    <#
    .SYNOPSIS
        Run `go test` with given args, print output, and write test logs.
    .DESCRIPTION
        Executes `go test -v -count=1` with the provided arguments,
        prints filtered output to the console, and invokes Write-TestLogs
        to create structured log files.
    .PARAMETER testArgs
        Arguments to pass to `go test` (e.g., "./..." or a specific package path).
    .OUTPUTS
        The exit code from `go test`.
    .EXAMPLE
        $exitCode = Invoke-GoTestAndLog "./..."
    #>
    [CmdletBinding()]
    param([string]$testArgs)

    $prevPref = $ErrorActionPreference
    $ErrorActionPreference = "Continue"
    $output = & go test -v -count=1 $testArgs 2>&1 | ForEach-Object { $_.ToString() }
    $exitCode = $LASTEXITCODE
    $ErrorActionPreference = $prevPref

    # Print to console
    Filter-TestWarnings $output | ForEach-Object { Write-Host $_ }

    # Write logs
    Write-TestLogs $output

    return $exitCode
}

function Open-FailingTestsIfAny {
    <#
    .SYNOPSIS
        Open the failing-tests log file if it contains failures.
    .DESCRIPTION
        Checks if data/test-logs/failing-tests.txt exists and has a
        non-zero count, then opens it in the default editor.
    .EXAMPLE
        Open-FailingTestsIfAny
    #>
    [CmdletBinding()]
    param()

    $failingFile = Join-Path $TestLogDir "failing-tests.txt"
    if ((Test-Path $failingFile)) {
        $content = Get-Content $failingFile -Raw
        if ($content -and $content -notmatch '# Count: 0') {
            Write-Host ""
            Write-Host "  Opening failing tests log..." -ForegroundColor Yellow
            Invoke-Item $failingFile
        }
    }
}

function Invoke-AllTests {
    <#
    .SYNOPSIS
        Run all Go tests with verbose output, build-check first.
    .DESCRIPTION
        Fetches latest code/deps, performs a build check on ./...,
        then runs all tests and writes logs. Opens failing log on failure.
    .EXAMPLE
        Invoke-AllTests
    #>
    [CmdletBinding()]
    param()

    Write-Header "Running all tests"
    Invoke-FetchLatest
    Push-Location tests
    try {
        if (-not (Invoke-BuildCheck "./...")) { return }

        $prevPref = $ErrorActionPreference
        $ErrorActionPreference = "Continue"
        $output = & go test -v -count=1 ./... 2>&1 | ForEach-Object { $_.ToString() }
        $exitCode = $LASTEXITCODE
        $ErrorActionPreference = $prevPref

        Filter-TestWarnings $output | ForEach-Object { Write-Host $_ }
        Write-TestLogs $output

        if ($exitCode -eq 0) { Write-Success "All tests passed" }
        else { Write-Fail "Some tests failed (exit code: $exitCode)" }
    }
    finally { Pop-Location }
    Open-FailingTestsIfAny
}

function Invoke-PackageTests {
    <#
    .SYNOPSIS
        Run Go tests for a single package under tests/integratedtests/.
    .PARAMETER pkg
        The package directory name (e.g., "regexnewtests").
    .EXAMPLE
        Invoke-PackageTests "regexnewtests"
    #>
    [CmdletBinding()]
    param([string]$pkg)

    if (-not $pkg) {
        Write-Fail "Package name required. Usage: ./run.ps1 TP <package>"
        Write-Host "  Available packages:" -ForegroundColor Yellow
        Get-ChildItem -Path tests/integratedtests -Directory | ForEach-Object {
            Write-Host "    - $($_.Name)" -ForegroundColor Gray
        }
        return
    }

    Write-Header "Running tests for package: $pkg"
    Invoke-FetchLatest
    Push-Location tests
    try {
        if (-not (Invoke-BuildCheck "./integratedtests/$pkg/...")) { return }

        $prevPref = $ErrorActionPreference
        $ErrorActionPreference = "Continue"
        $output = & go test -v -count=1 "./integratedtests/$pkg/..." 2>&1 | ForEach-Object { $_.ToString() }
        $exitCode = $LASTEXITCODE
        $ErrorActionPreference = $prevPref

        Filter-TestWarnings $output | ForEach-Object { Write-Host $_ }
        Write-TestLogs $output

        if ($exitCode -eq 0) { Write-Success "Package tests passed" }
        else { Write-Fail "Package tests failed (exit code: $exitCode)" }
    }
    finally { Pop-Location }
    Open-FailingTestsIfAny
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Invoke-GitPull',
    'Invoke-FetchLatest',
    'Invoke-BuildCheck',
    'Invoke-GoTestAndLog',
    'Open-FailingTestsIfAny',
    'Invoke-AllTests',
    'Invoke-PackageTests'
)
