# ─────────────────────────────────────────────────────────────────────────────
# PreCommitCheck.psm1 — Pre-commit API mismatch checker for Coverage* files
#
# Runs regression guard, safeTest boundary lint, Go auto-fixer, bracecheck,
# and then compile-checks all test packages that contain Coverage* files
# to detect API mismatches before committing.
#
# Usage:
#   Import-Module ./scripts/PreCommitCheck.psm1 -Force
#
# Dependencies:
#   - Utilities.psm1     (Write-Header, Write-Success, Write-Fail,
#                          Merge-UniqueOutputLines, ParseCompileErrors)
#   - DashboardUI.psm1   (Register-Phase, Reset-Phases, Write-PhaseSummaryBox)
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-PreCommitCheck {
    <#
    .SYNOPSIS
        Run the full pre-commit validation pipeline.
    .DESCRIPTION
        Phases:
          1. Regression guard (CaseV1 fields, corejson.Result.Err usage)
          2. SafeTest boundary + empty-if lint
          3. Go auto-fixer (optional, skip with --no-autofix)
          4. Bracecheck syntax validation (skip with --skip-bracecheck)
          5. Compile-check all Coverage* packages (parallel or --sync)

        Produces a JSON report at data/precommit/api-check.json and
        renders a dashboard phase summary.
    .PARAMETER singlePkg
        Optional package name to check only one package.
    .EXAMPLE
        Invoke-PreCommitCheck
        Invoke-PreCommitCheck "corejsontests"
    #>
    [CmdletBinding()]
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

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Invoke-PreCommitCheck'
)
