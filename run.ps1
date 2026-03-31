#!/usr/bin/env pwsh
<#
.SYNOPSIS
    Project runner script with shorthands for common operations.

.DESCRIPTION
    Usage: ./run.ps1 <command> [options]

    Commands (uppercase shorthands OR hyphen-lowercase):
        T   | -t   | test          Run all tests (verbose)
        TP  | -tp  | test-pkg      Run tests for a specific package: ./run.ps1 TP regexnewtests
        TC  | -tc  | test-cover    Run tests with coverage (parallel by default)
        TCP | -tcp | test-cover-pkg Run coverage for a specific package: ./run.ps1 TCP regexnewtests
        TI  | -ti  | test-int      Run integrated tests only
        TF  | -tf  | test-fail     Show last failing tests log
        GC  | -gc  | goconvey      Launch GoConvey (browser test runner)
        R   | -r   | run           Run the main application
        B   | -b   | build         Build the binary
        BR  | -br  | build-run     Build then run
        F   | -f   | fmt           Format all Go files
        L   | -l   | lint          Run go vet on all packages
        V   | -v   | vet           Run go vet
        TY  | -ty  | tidy          Run go mod tidy
        PC  | -pc  | pre-commit    Check Coverage* files for API mismatches
        C   | -c   | clean         Clean build artifacts
        H   | -h   | help          Show this help

    Mode options (for TC/TCP):
        --sync             Run precompile + tests sequentially (default: parallel)
        --open             Open HTML coverage report in browser (default: don't open)
        --skip-bracecheck  Skip the Go syntax pre-check for faster runs
        --no-autofix       Skip the Go auto-fixer before bracecheck
        --dry-run          Run auto-fixer in preview mode (show fixes without applying)

.EXAMPLE
    ./run.ps1 T
    ./run.ps1 -t
    ./run.ps1 TP regexnewtests
    ./run.ps1 -tp regexnewtests
    ./run.ps1 TC --sync
    ./run.ps1 -gc
#>

param(
    [Parameter(Position = 0)]
    [string]$Command = "help",

    [Parameter(Position = 1, ValueFromRemainingArguments)]
    [string[]]$ExtraArgs
)

# Normalize: if $Command was swallowed by PowerShell as a switch
# (e.g. -gc parsed away), $Command will be "help" — detect via $PSBoundParameters.
if (-not $PSBoundParameters.ContainsKey('Command')) {
    # Check $MyInvocation.Line for the actual argument
    $rawLine = $MyInvocation.Line
    $match = [regex]::Match($rawLine, '(?i)run\.ps1\s+(-?\w[\w-]*)\s*(.*)')
    if ($match.Success) {
        $Command = $match.Groups[1].Value
        # Capture remaining args that PowerShell swallowed
        $trailing = $match.Groups[2].Value.Trim()
        if ($trailing -and (-not $ExtraArgs -or $ExtraArgs.Count -eq 0)) {
            $ExtraArgs = @($trailing -split '\s+')
        }
    }
}

# -- Encoding (fix garbled Unicode on Windows terminals) --
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8
$OutputEncoding            = [System.Text.Encoding]::UTF8

$ErrorActionPreference = "Stop"

# -- Dashboard UI Module --
$dashboardModule = Join-Path $PSScriptRoot "scripts" "DashboardUI.psm1"
if (Test-Path $dashboardModule) {
    Import-Module $dashboardModule -Force -DisableNameChecking
    # Auto-detect theme, or honor --light / --dark flags
    $themeOverride = $null
    if ($ExtraArgs -contains '--light') { $themeOverride = "light" }
    if ($ExtraArgs -contains '--dark')  { $themeOverride = "dark" }
    if ($themeOverride) {
        Initialize-DashboardUI -Theme $themeOverride
    } else {
        Initialize-DashboardUI
    }
}

# -- Shared Variables --
$TestLogDir = Join-Path $PSScriptRoot "data" "test-logs"

# -- Utilities Module --
$utilitiesModule = Join-Path $PSScriptRoot "scripts" "Utilities.psm1"
if (Test-Path $utilitiesModule) {
    Import-Module $utilitiesModule -Force -DisableNameChecking
}

# -- Test Log Writer Module --
$testLogWriterModule = Join-Path $PSScriptRoot "scripts" "TestLogWriter.psm1"
if (Test-Path $testLogWriterModule) {
    Import-Module $testLogWriterModule -Force -DisableNameChecking
}



# -- Test Runner Module --
$testRunnerModule = Join-Path $PSScriptRoot "scripts" "TestRunner.psm1"
if (Test-Path $testRunnerModule) {
    Import-Module $testRunnerModule -Force -DisableNameChecking
}

# -- Coverage Runner Module --
$coverageRunnerModule = Join-Path $PSScriptRoot "scripts" "CoverageRunner.psm1"
if (Test-Path $coverageRunnerModule) {
    Import-Module $coverageRunnerModule -Force -DisableNameChecking
}

function Invoke-IntegratedTests {
    Write-Header "Running integrated tests only"
    Invoke-FetchLatest
    Push-Location tests
    try {
        if (-not (Invoke-BuildCheck "./integratedtests/...")) { return }

        $prevPref = $ErrorActionPreference
        $ErrorActionPreference = "Continue"
        $output = & go test -v -count=1 ./integratedtests/... 2>&1 | ForEach-Object { $_.ToString() }
        $exitCode = $LASTEXITCODE
        $ErrorActionPreference = $prevPref

        Filter-TestWarnings $output | ForEach-Object { Write-Host $_ }
        Write-TestLogs $output

        if ($exitCode -eq 0) { Write-Success "Integrated tests passed" }
        else { Write-Fail "Integrated tests failed (exit code: $exitCode)" }
    }
    finally { Pop-Location }
    Open-FailingTestsIfAny
}

function Invoke-RunMain {
    Write-Header "Running main application"
    go run ./cmd/main/*.go
}

function Invoke-Build {
    Write-Header "Building binary"
    $buildDir = "build"
    if (-not (Test-Path $buildDir)) { New-Item -ItemType Directory -Path $buildDir | Out-Null }
    go build -o "$buildDir/cli" ./cmd/main/
    if ($LASTEXITCODE -eq 0) { Write-Success "Build complete: $buildDir/cli" }
    else { Write-Fail "Build failed" }
}

function Invoke-BuildRun {
    Invoke-Build
    if ($LASTEXITCODE -eq 0) {
        Write-Header "Running built binary"
        & ./build/cli
    }
}

function Invoke-Format {
    Write-Header "Formatting Go files"
    gofmt -w -s .
    Write-Success "Formatting complete"
}

function Invoke-Vet {
    Write-Header "Running go vet"
    go vet ./...
    if ($LASTEXITCODE -eq 0) { Write-Success "No issues found" }
    else { Write-Fail "Issues found" }
}

function Invoke-Tidy {
    Write-Header "Running go mod tidy"
    go mod tidy
    Write-Success "Tidy complete"
}

function Invoke-GoConvey {
    Write-Header "Launching GoConvey"

    # Check if goconvey is installed
    $gcPath = Get-Command goconvey -ErrorAction SilentlyContinue
    if (-not $gcPath) {
        Write-Host "  GoConvey not found. Installing..." -ForegroundColor Yellow
        go install github.com/smartystreets/goconvey@latest
        if ($LASTEXITCODE -ne 0) {
            Write-Fail "Failed to install GoConvey"
            return
        }
        Write-Success "GoConvey installed"
    }

    $port = if ($ExtraArgs -and $ExtraArgs[0]) { $ExtraArgs[0] } else { "8080" }
    Write-Host "  Starting GoConvey on http://localhost:$port" -ForegroundColor Yellow
    Write-Host "  Press Ctrl+C to stop" -ForegroundColor Gray

    Push-Location tests
    try {
        goconvey -port $port
    }
    finally { Pop-Location }
}

function Invoke-PreCommitCheck {
    param([string]$singlePkg)

    # Reset phase tracker for this run
    if (Get-Command Reset-Phases -ErrorAction SilentlyContinue) { Reset-Phases }

    Write-Header "Pre-commit API mismatch checker"

    $isSyncMode = $false
    if ($ExtraArgs) {
        foreach ($ea in $ExtraArgs) {
            if ($ea -eq "--sync") { $isSyncMode = $true }
        }
    }

    # Fast regression guard (legacy CaseV1 fields + invalid corejson.Result.Err usage)
    $regressionScript = Join-Path $PSScriptRoot "scripts" "check-integrated-regressions.ps1"
    if (-not (Test-Path $regressionScript)) {
        Write-Fail "Regression guard script not found: $regressionScript"
        exit 1
    }

    Write-Host "  Running regression guard scan..." -ForegroundColor Yellow
    if ($singlePkg) {
        & $regressionScript -SinglePackage $singlePkg
    }
    else {
        & $regressionScript
    }

    if ($LASTEXITCODE -ne 0) {
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Regression Guard" "fail" "regressions detected" }
        Write-Fail "Regression guard failed. Fix reported issues before PC."
        exit 1
    }
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Regression Guard" "pass" "no regressions" }

    # safeTest boundary + empty-if lint check
    $boundaryScript = Join-Path $PSScriptRoot "scripts" "check-safetest-boundaries.ps1"
    if (Test-Path $boundaryScript) {
        Write-Host "  Running safeTest boundary + empty-if lint check..." -ForegroundColor Yellow
        & $boundaryScript
        if ($LASTEXITCODE -ne 0) {
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "SafeTest Lint" "fail" "boundary check failed" }
            Write-Fail "safeTest boundary check failed. Fix reported issues before PC."
            exit 1
        }
    }
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "SafeTest Lint" "pass" "all clean" }

    # ── Go auto-fixer ─────────────────────────────────────────────────
    $skipAutofix = $ExtraArgs -and ($ExtraArgs -contains '--no-autofix')
    $skipBrace = $ExtraArgs -and ($ExtraArgs -contains '--skip-bracecheck')
    if ($skipBrace) {
        Write-Host "  Skipping Go auto-fixer and syntax pre-check (--skip-bracecheck)" -ForegroundColor DarkYellow
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
            Register-Phase "Auto-Fixer" "skip" "skipped (--skip-bracecheck)"
        }
    } elseif ($skipAutofix) {
        Write-Host "  Skipping Go auto-fixer (--no-autofix)" -ForegroundColor DarkYellow
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
            Register-Phase "Auto-Fixer" "skip" "skipped (--no-autofix)"
        }
    } else {
        $dryRunFlag = if ($ExtraArgs -and ($ExtraArgs -contains '--dry-run')) { '--dry-run' } else { $null }
        $dryLabel = if ($dryRunFlag) { " (dry-run)" } else { "" }
        Write-Host "  Running Go auto-fixer$dryLabel..." -ForegroundColor Yellow
        $fixArgs = @('./scripts/autofix/')
        if ($dryRunFlag) { $fixArgs += '--dry-run' }
        $fixOut = & go run @fixArgs 2>&1
        if ($LASTEXITCODE -ne 0) {
            Write-Host ($fixOut | Out-String) -ForegroundColor Red
            Write-Fail "Go auto-fixer encountered errors."
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Auto-Fixer" "warn" "errors encountered" }
        } else {
            $fixStr = ($fixOut | Out-String).Trim()
            if ($fixStr) { Write-Success $fixStr }
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Auto-Fixer" "pass" "no fixable issues" }
        }
    }

    # ── Go syntax pre-check (bracecheck) ──────────────────────────────
    if ($skipBrace) {
        # already logged above
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Syntax Check" "skip" "skipped" }
    } else {
        Write-Host "  Running Go syntax pre-check (bracecheck)..." -ForegroundColor Yellow
        $braceOut = & go run ./scripts/bracecheck/ 2>&1
        if ($LASTEXITCODE -ne 0) {
            Write-Host ($braceOut | Out-String) -ForegroundColor Red
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Syntax Check" "fail" "bracecheck failed" }
            Write-Fail "Go syntax check failed. Fix reported issues before PC."
            exit 1
        } else {
            $braceStr3 = ($braceOut | Out-String).Trim()
            Write-Success $braceStr3
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Syntax Check" "pass" $braceStr3 }
        }

        # ── Write syntax-issues.txt report ────────────────────────────
        $syntaxReportDir = Join-Path $PSScriptRoot "data" "coverage"
        New-Item -ItemType Directory -Path $syntaxReportDir -Force | Out-Null
        $syntaxReportFile = Join-Path $syntaxReportDir "syntax-issues.txt"
        $braceStr = ($braceOut | Out-String).Trim()
        $syntaxTs = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
        $syntaxContent = @(
            "================================================================================"
            "  Syntax Issues Report — $syntaxTs"
            "  Generated by: autofix + bracecheck pipeline"
            "================================================================================"
            ""
        )
        if (Test-Path $syntaxReportFile) {
            $existing = Get-Content $syntaxReportFile -Raw
            $syntaxContent = @($existing.TrimEnd(), "", "")
        }
        $syntaxContent += @(
            "────────────────────────────────────────────────────────────────────────────────"
            " BRACECHECK RESULTS"
            "────────────────────────────────────────────────────────────────────────────────"
            ""
            "  $braceStr"
            ""
            "================================================================================"
        )
        Set-Content -Path $syntaxReportFile -Value ($syntaxContent -join "`n") -Encoding UTF8
    }

    # Discover test packages
    $testBaseDir = Join-Path $PSScriptRoot "tests" "integratedtests"
    if ($singlePkg) {
        $targetDirs = @(Join-Path $testBaseDir $singlePkg)
        if (-not (Test-Path $targetDirs[0])) {
            Write-Fail "Package not found: $singlePkg"
            return
        }
    } else {
        $targetDirs = @(Get-ChildItem -Path $testBaseDir -Directory | ForEach-Object { $_.FullName })
    }

    # Filter to only dirs containing Coverage* files
    $pkgsWithCoverage = [System.Collections.Generic.List[string]]::new()
    foreach ($dir in $targetDirs) {
        $coverFiles = Get-ChildItem -Path $dir -Filter "Coverage*" -File -ErrorAction SilentlyContinue
        if ($coverFiles -and $coverFiles.Count -gt 0) {
            $pkgsWithCoverage.Add($dir)
        }
    }

    if ($pkgsWithCoverage.Count -eq 0) {
        Write-Success "No Coverage* files found to check"
        return
    }

    $modeLabel = if ($isSyncMode) { "sync" } else { "parallel" }
    Write-Host "  Checking $($pkgsWithCoverage.Count) packages with Coverage* files ($modeLabel)..." -ForegroundColor Yellow
    Write-Host ""

    # Convert dirs to Go package paths
    $goTestPkgs = [System.Collections.Generic.List[string]]::new()
    foreach ($dir in $pkgsWithCoverage) {
        $relPath = $dir -replace [regex]::Escape($PSScriptRoot), '' -replace '^[\\/]', '' -replace '\\', '/'
        $goTestPkgs.Add("github.com/alimtvnetwork/core/$relPath")
    }

    # Compile check
    $compileTemp = Join-Path $PSScriptRoot "data" "precommit"
    if (Test-Path $compileTemp) { Remove-Item -Recurse -Force $compileTemp }
    New-Item -ItemType Directory -Path $compileTemp -Force | Out-Null

    $failures = [System.Collections.Generic.List[object]]::new()
    $passedCount = 0

    if ($isSyncMode) {
        foreach ($pkg in $goTestPkgs) {
            $shortName = $pkg -replace '.*integratedtests/?', ''
            $safeName = $pkg -replace '[^a-zA-Z0-9\.-]', '_'
            $outFile = Join-Path $compileTemp "check-$safeName.test"

            $prevPref = $ErrorActionPreference
            $ErrorActionPreference = "Continue"
            $compOut = & go test -c -gcflags=all=-e -o $outFile "$pkg" 2>&1 | ForEach-Object { $_.ToString() }
            $ec = $LASTEXITCODE
            $ErrorActionPreference = $prevPref

            if ($ec -eq 0) {
                $passedCount++
            } else {
                $prevPref = $ErrorActionPreference
                $ErrorActionPreference = "Continue"
                $diagOut = & go test -count=1 -run '^$' -gcflags=all=-e "$pkg" 2>&1 | ForEach-Object { $_.ToString() }
                $ErrorActionPreference = $prevPref
                $compOut = Merge-UniqueOutputLines $compOut $diagOut

                $parsedErrors = ParseCompileErrors $compOut
                $failures.Add(@{
                    package    = $shortName
                    errorCount = $parsedErrors.Count
                    errors     = $parsedErrors
                })
            }
        }
    } else {
        $throttle = [Math]::Min($goTestPkgs.Count, [Environment]::ProcessorCount * 2)
        $results = $goTestPkgs | ForEach-Object -ThrottleLimit $throttle -Parallel {
            $pkg = $_
            $tempDir = $using:compileTemp
            $safeName = $pkg -replace '[^a-zA-Z0-9\.-]', '_'
            $outFile = Join-Path $tempDir "check-$safeName.test"
            $ErrorActionPreference = "Continue"
            $rawOut = & go test -c -gcflags=all=-e -o $outFile "$pkg" 2>&1
            $ec = $LASTEXITCODE
            $out = @($rawOut | ForEach-Object { $_.ToString() })

            if ($ec -ne 0) {
                $diagRaw = & go test -count=1 -run '^$' -gcflags=all=-e "$pkg" 2>&1
                $diagOut = @($diagRaw | ForEach-Object { $_.ToString() })

                $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
                $merged = [System.Collections.Generic.List[string]]::new()

                foreach ($line in @($out + $diagOut)) {
                    if ($null -eq $line) { continue }
                    $normalized = $line.ToString().TrimEnd("`r")
                    if (-not $normalized) { continue }
                    if ($seen.Add($normalized)) {
                        $merged.Add($normalized) | Out-Null
                    }
                }

                $out = $merged.ToArray()
            }

            [pscustomobject]@{ Pkg = $pkg; ExitCode = $ec; Output = $out }
        }

        foreach ($r in ($results | Sort-Object Pkg)) {
            $shortName = $r.Pkg -replace '.*integratedtests/?', ''
            if ($r.ExitCode -eq 0) {
                $passedCount++
            } else {
                
                $parsedErrors = ParseCompileErrors $r.Output
                $failures.Add(@{
                    package    = $shortName
                    errorCount = $parsedErrors.Count
                    errors     = $parsedErrors
                })
            }
        }
    }

    # Clean up compile artifacts
    Get-ChildItem -Path $compileTemp -Filter "*.test" -ErrorAction SilentlyContinue | Remove-Item -Force

    # Summary
    $allPassed = $failures.Count -eq 0
    Write-Host ""
    if ($allPassed) {
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Green
        Write-Host "  │ ✓ ALL $passedCount PACKAGES PASSED API CHECK" -ForegroundColor Green
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Green
    } else {
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Red
        Write-Host "  │ ✗ $($failures.Count) PACKAGE(S) HAVE API MISMATCHES" -ForegroundColor Red
        Write-Host "  │" -ForegroundColor Red
        foreach ($f in $failures) {
            Write-Host "  │   ✗ $($f.package) ($($f.errorCount) error(s))" -ForegroundColor Red
        }
        Write-Host "  │" -ForegroundColor Red
        Write-Host "  │ Fix these before committing Coverage* files." -ForegroundColor Yellow
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Red

        # Print error details
        Write-Host ""
        foreach ($f in $failures) {
            Write-Host "  ── $($f.package) ──" -ForegroundColor Red
            foreach ($e in $f.errors) {
                $cat = $e.category
                Write-Host "    $($e.file):$($e.line) [$cat] $($e.message)" -ForegroundColor Yellow
            }
            Write-Host ""
        }
    }

    # Write JSON report
    $jsonReport = @{
        timestamp    = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ")
        passed       = $allPassed
        checkedCount = $goTestPkgs.Count
        passedCount  = $passedCount
        failedCount  = $failures.Count
        failures     = $failures.ToArray()
    }
    $jsonPath = Join-Path $compileTemp "api-check.json"
    $jsonReport | ConvertTo-Json -Depth 5 | Set-Content -Path $jsonPath -Encoding UTF8
    Write-Host "  Report → $jsonPath" -ForegroundColor Gray

    # Register compile check phase
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
        if ($allPassed) {
            Register-Phase "API Compile Check" "pass" "$passedCount/$($goTestPkgs.Count) passed"
        } else {
            Register-Phase "API Compile Check" "fail" "$($failures.Count) failed, $passedCount passed"
        }
    }

    # ── Dashboard Phase Summary ──
    if (Get-Command Write-PhaseSummaryBox -ErrorAction SilentlyContinue) {
        Write-Host ""
        Write-PhaseSummaryBox
    }

    if (-not $allPassed) { exit 1 }
}


# ParseCompileErrors — moved to scripts/Utilities.psm1

function Invoke-Clean {
    Write-Header "Cleaning build artifacts"
    if (Test-Path build) { Remove-Item -Recurse -Force build }
    if (Test-Path tests/coverage.out) { Remove-Item tests/coverage.out }
    $coverDir = Join-Path $PSScriptRoot "data" "coverage"
    if (Test-Path $coverDir) { Remove-Item -Recurse -Force $coverDir; Write-Success "Removed coverage reports" }
    $precommitDir = Join-Path $PSScriptRoot "data" "precommit"
    if (Test-Path $precommitDir) { Remove-Item -Recurse -Force $precommitDir; Write-Success "Removed precommit reports" }
    Write-Success "Clean complete"
}

function Invoke-ShowFailLog {
    $failingFile = Join-Path $TestLogDir "failing-tests.txt"
    if (-not (Test-Path $failingFile)) {
        Write-Header "No failing tests log found"
        Write-Host "  Run tests first: ./run.ps1 T" -ForegroundColor Yellow
        return
    }

    Write-Header "Last Failing Tests"
    $content = Get-Content $failingFile -Raw
    if ($content -match '# Count: 0') {
        Write-Success "No failing tests in last run"
    }
    else {
        Write-Host $content
    }
    Write-Host ""
    Write-Host "  Log file: $failingFile" -ForegroundColor Gray
}

function Show-Help {
    Write-Host ""
    Write-Host "  Project Runner — ./run.ps1 <command> [options]" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "  Testing:" -ForegroundColor Yellow
    Write-Host "    T   | -t   | test          Run all tests (verbose)"
    Write-Host "    TP  | -tp  | test-pkg      Run tests for a specific package"
    Write-Host "    TC  | -tc  | test-cover    Run tests with coverage (HTML + summary)"
    Write-Host "    TCP | -tcp | test-cover-pkg Run coverage for a specific package"
    Write-Host "    TI  | -ti  | test-int      Run integrated tests only"
    Write-Host "    TF  | -tf  | test-fail     Show last failing tests log"
    Write-Host "    GC  | -gc  | goconvey      Launch GoConvey (browser test runner)"
    Write-Host ""
    Write-Host "  Build & Run:" -ForegroundColor Yellow
    Write-Host "    R   | -r   | run           Run the main application"
    Write-Host "    B   | -b   | build         Build the binary"
    Write-Host "    BR  | -br  | build-run     Build then run"
    Write-Host ""
    Write-Host "  Code Quality:" -ForegroundColor Yellow
    Write-Host "    F   | -f   | fmt           Format all Go files"
    Write-Host "    L   | -l   | lint          Run go vet"
    Write-Host "    V   | -v   | vet           Run go vet"
    Write-Host "    TY  | -ty  | tidy          Run go mod tidy"
    Write-Host "    PC  | -pc  | pre-commit    Check Coverage* files for API mismatches"
    Write-Host ""
    Write-Host "  Other:" -ForegroundColor Yellow
    Write-Host "    C   | -c   | clean         Clean build artifacts"
    Write-Host "    H   | -h   | help          Show this help"
    Write-Host ""
    Write-Host "  Mode Options (for TC/TCP/PC):" -ForegroundColor Yellow
    Write-Host "    --sync      Run precompile + tests sequentially (default: parallel)"
    Write-Host "    --open      Open HTML coverage report in browser after TC/TCP"
    Write-Host ""
    Write-Host "  Examples:" -ForegroundColor Gray
    Write-Host "    ./run.ps1 T"
    Write-Host "    ./run.ps1 -t"
    Write-Host "    ./run.ps1 TP regexnewtests"
    Write-Host "    ./run.ps1 -tp regexnewtests"
    Write-Host "    ./run.ps1 TCP regexnewtests  (package coverage)"
    Write-Host "    ./run.ps1 TC                 (parallel by default)"
    Write-Host "    ./run.ps1 TC --sync          (sequential mode)"
    Write-Host "    ./run.ps1 TC --sync --no-open"
    Write-Host "    ./run.ps1 PC                 (pre-commit check)"
    Write-Host "    ./run.ps1 PC corejsontests   (check single package)"
    Write-Host "    ./run.ps1 -gc"
    Write-Host "    ./run.ps1 -gc 9090          (custom port)"
    Write-Host ""
}

# -- Dispatch --
$firstExtraArg = if ($ExtraArgs -and $ExtraArgs.Count -gt 0) { $ExtraArgs[0] } else { $null }

switch ($Command.ToLower()) {
    { $_ -in "t", "-t", "test" }              { Invoke-AllTests }
    { $_ -in "tp", "-tp", "test-pkg" }        { Invoke-PackageTests $firstExtraArg }
    { $_ -in "tc", "-tc", "test-cover" }      { Invoke-TestCoverage }
    { $_ -in "tcp", "-tcp", "test-cover-pkg" } { Invoke-PackageTestCoverage $firstExtraArg }
    { $_ -in "ti", "-ti", "test-int" }        { Invoke-IntegratedTests }
    { $_ -in "tf", "-tf", "test-fail" }       { Invoke-ShowFailLog }
    { $_ -in "gc", "-gc", "goconvey" }        { Invoke-GoConvey }
    { $_ -in "r", "-r", "run" }               { Invoke-RunMain }
    { $_ -in "b", "-b", "build" }             { Invoke-Build }
    { $_ -in "br", "-br", "build-run" }       { Invoke-BuildRun }
    { $_ -in "f", "-f", "fmt" }               { Invoke-Format }
    { $_ -in "l", "-l", "lint", "v", "-v", "vet" } { Invoke-Vet }
    { $_ -in "ty", "-ty", "tidy" }            { Invoke-Tidy }
    { $_ -in "pc", "-pc", "pre-commit" }      { Invoke-PreCommitCheck $firstExtraArg }
    { $_ -in "c", "-c", "clean" }             { Invoke-Clean }
    { $_ -in "h", "-h", "help", "" }          { Show-Help }
    default {
        Write-Fail "Unknown command: '$Command'"
        Show-Help
    }
}
