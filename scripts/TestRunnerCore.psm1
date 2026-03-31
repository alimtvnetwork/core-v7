# ─────────────────────────────────────────────────────────────────────────────
# TestRunnerCore.psm1 — Git ops, build check, test invocation primitives
#
# Dependencies: Utilities.psm1, ErrorParser.psm1, TestLogWriter.psm1
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-GitPull {
    <# .SYNOPSIS Pull the latest changes from the remote git repository. #>
    [CmdletBinding()]
    param()
    Write-Header "Pulling latest from remote"
    $prevPref = $ErrorActionPreference; $ErrorActionPreference = "Continue"
    git pull 2>&1 | ForEach-Object { Write-Host "  $_" -ForegroundColor Gray }
    if ($LASTEXITCODE -eq 0) { Write-Success "Git pull complete" } else { Write-Fail "git pull failed (continuing anyway)" }
    $ErrorActionPreference = $prevPref
}

function Invoke-FetchLatest {
    <# .SYNOPSIS Pull git changes and run `go mod tidy` to sync dependencies. #>
    [CmdletBinding()]
    param()
    Invoke-GitPull
    Write-Header "Fetching latest dependencies"
    go mod tidy
    if ($LASTEXITCODE -eq 0) { Write-Success "Dependencies up to date" } else { Write-Fail "go mod tidy failed" }
}

function Invoke-BuildCheck {
    <# .SYNOPSIS Compile-check a Go package path before running tests. Returns $true/$false. #>
    [CmdletBinding()]
    param([string]$buildPath)
    Write-Header "Build check: $buildPath"
    $prevPref = $ErrorActionPreference; $ErrorActionPreference = "Continue"
    $output = & go build $buildPath 2>&1 | ForEach-Object { $_.ToString() }
    $exitCode = $LASTEXITCODE; $ErrorActionPreference = $prevPref

    if ($exitCode -ne 0) {
        Write-Fail "Build failed — skipping tests"
        Ensure-TestLogDir
        $failingFile = Join-Path $TestLogDir "failing-tests.txt"
        $rawFile     = Join-Path $TestLogDir "raw-output.txt"
        $timestamp   = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
        $buildErrors = Extract-BuildErrorLines $output
        $errorCount = if ($buildErrors.Count -gt 0) { $buildErrors.Count } else { 1 }
        $failingContent = @("# Failing Tests — $timestamp", "# Count: $errorCount", "", "# Build Failed — tests were NOT run", "", "# ── Build Errors ──", "")
        if ($buildErrors.Count -gt 0) { $failingContent += $buildErrors } else { $failingContent += $output }
        Set-Content -Path $failingFile -Value ($failingContent -join "`n") -Encoding UTF8
        Set-Content -Path $rawFile -Value ($output -join "`n") -Encoding UTF8
        $output | ForEach-Object { Write-Host "  $_" -ForegroundColor Red }
        Open-FailingTestsIfAny
        return $false
    }
    Write-Success "Build OK"; return $true
}

function Invoke-GoTestAndLog {
    <# .SYNOPSIS Run `go test` with given args, print output, and write test logs. Returns exit code. #>
    [CmdletBinding()]
    param([string]$testArgs)
    $prevPref = $ErrorActionPreference; $ErrorActionPreference = "Continue"
    $output = & go test -v -count=1 $testArgs 2>&1 | ForEach-Object { $_.ToString() }
    $exitCode = $LASTEXITCODE; $ErrorActionPreference = $prevPref
    Filter-TestWarnings $output | ForEach-Object { Write-Host $_ }
    Write-TestLogs $output
    return $exitCode
}

function Open-FailingTestsIfAny {
    <# .SYNOPSIS Open the failing-tests log file if it contains failures. #>
    [CmdletBinding()]
    param()
    $failingFile = Join-Path $TestLogDir "failing-tests.txt"
    if ((Test-Path $failingFile)) {
        $content = Get-Content $failingFile -Raw
        if ($content -and $content -notmatch '# Count: 0') {
            Write-Host ""; Write-Host "  Opening failing tests log..." -ForegroundColor Yellow
            Invoke-Item $failingFile
        }
    }
}

Export-ModuleMember -Function @(
    'Invoke-GitPull', 'Invoke-FetchLatest', 'Invoke-BuildCheck',
    'Invoke-GoTestAndLog', 'Open-FailingTestsIfAny'
)
