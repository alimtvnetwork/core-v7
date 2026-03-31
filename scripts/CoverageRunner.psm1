# ─────────────────────────────────────────────────────────────────────────────
# CoverageRunner.psm1 — Go test coverage execution and reporting
#
# Contains the two main coverage commands:
#   - Invoke-TestCoverage        (TC)  — Full coverage run across all packages
#   - Invoke-PackageTestCoverage (TCP) — Coverage for a single package
#
# These are the largest functions in the toolchain (~1500 lines combined).
# They handle: pre-checks (safetest, autofix, bracecheck), parallel/sync
# compile checks, per-file split recovery, coverage profile merging,
# per-package coverage computation, HTML/JSON/TXT report generation,
# coverage comparison (diff), and AI prompt generation.
#
# Usage:
#   Import-Module ./scripts/CoverageRunner.psm1 -Force
#
# Dependencies:
#   - Utilities.psm1     (Write-Header, Write-Success, Write-Fail,
#                          Ensure-TestLogDir, Filter-TestWarnings,
#                          Extract-BuildErrorLines, Extract-ExecutionFailureLines,
#                          Merge-UniqueOutputLines, Add-BuildErrorsForPackage,
#                          Add-RuntimeFailuresForPackage)
#   - TestLogWriter.psm1 (Write-TestLogs)
#   - TestRunner.psm1    (Invoke-FetchLatest, Invoke-BuildCheck,
#                          Open-FailingTestsIfAny)
#   - DashboardUI.psm1   (Register-Phase, Reset-Phases, Write-PhaseSummaryBox,
#                          Write-CoverageComparison, Load/Save-CoverageSnapshot)
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-TestCoverage {
    <#
    .SYNOPSIS
        Run full test coverage across all packages (TC command).
    .DESCRIPTION
        Executes the complete coverage pipeline: pre-checks (safetest, autofix,
        bracecheck), parallel/sync compile checks, test execution with coverage
        profiling, report generation (HTML/JSON/TXT), coverage diff comparison,
        and AI prompt generation.
    #>
    Write-Header "Running tests with coverage"

    # Reset phase tracker for this run
    if (Get-Command Reset-Phases -ErrorAction SilentlyContinue) { Reset-Phases }

    Invoke-FetchLatest
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
        Register-Phase "Git Pull" "pass" "pulled from remote"
        Register-Phase "Dependencies" "pass" "up to date"
    }

    # Clean data folder before running tests
    $dataDir = Join-Path $PSScriptRoot "data"
    if (Test-Path $dataDir) {
        Remove-Item -Recurse -Force $dataDir
        Write-Host "  Cleaned data/ folder" -ForegroundColor Yellow
    }
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
        Register-Phase "Data Cleanup" "pass" "cleaned"
    }

    $coverDir = Join-Path $PSScriptRoot "data" "coverage"
    $partialDir = Join-Path $coverDir "partial"
    New-Item -ItemType Directory -Path $partialDir -Force | Out-Null

    $coverProfile = Join-Path $coverDir "coverage.out"
    $coverHtml    = Join-Path $coverDir "coverage.html"
    $coverSummary = Join-Path $coverDir "coverage-summary.txt"
    $repoBuildErrorsFile = Join-Path $coverDir "repo-build-errors.txt"
    $repoBuildErrorsJsonFile = Join-Path $coverDir "repo-build-errors.json"

    $repoBuildErrorsScript = Join-Path $PSScriptRoot "scripts" "coverage" "Export-RepoBuildErrors.ps1"
    if (Test-Path $repoBuildErrorsScript) {
        & $repoBuildErrorsScript -OutputTxt $repoBuildErrorsFile -OutputJson $repoBuildErrorsJsonFile
    }

    # Build coverpkg list: all source packages EXCLUDING tests/
    $allPkgs = go list ./... 2>&1 | ForEach-Object { $_.ToString() }
    $srcPkgs = $allPkgs | Where-Object { $_ -notmatch '/tests/' }
    $covPkgList = $srcPkgs -join ","

    # Get all test packages to run individually (deterministic order)
    # Include both integration tests AND source packages with in-package _test.go files
    $integrationTestPkgs = go list ./tests/integratedtests/... 2>&1 |
        ForEach-Object { $_.ToString() } |
        Where-Object { $_ -and $_ -notmatch '^warning:' }

    # Find source packages that have _test.go files (for unexported symbol coverage)
    $inPkgTestPkgs = @()
    foreach ($srcPkg in $srcPkgs) {
        $relPath = $srcPkg -replace '^github\.com/alimtvnetwork/core/', ''
        if ($relPath -and (Test-Path $relPath) -and (Get-ChildItem -Path $relPath -Filter '*_test.go' -File -ErrorAction SilentlyContinue)) {
            $inPkgTestPkgs += $srcPkg
        }
    }

    $allTestPkgs = @($integrationTestPkgs) + @($inPkgTestPkgs) | Sort-Object -Unique

    # ── safeTest boundary + empty-if lint check ────────────────────
    $boundaryScript = Join-Path $PSScriptRoot "scripts" "check-safetest-boundaries.ps1"
    if (Test-Path $boundaryScript) {
        Write-Host ""
        Write-Host "  Running safeTest boundary + empty-if lint check..." -ForegroundColor Yellow
        & $boundaryScript
        if ($LASTEXITCODE -ne 0) {
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "SafeTest Lint" "fail" "boundary check failed" }
            Write-Fail "safeTest boundary check failed. Fix reported issues before TC."
            exit 1
        }
    }
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "SafeTest Lint" "pass" "all clean" }

    # ── Go auto-fixer ─────────────────────────────────────────────────
    $skipAutofix = $ExtraArgs -and ($ExtraArgs -contains '--no-autofix')
    $skipBrace = $ExtraArgs -and ($ExtraArgs -contains '--skip-bracecheck')
    if ($skipBrace) {
        Write-Host "  Skipping Go auto-fixer and syntax pre-check (--skip-bracecheck)" -ForegroundColor DarkYellow
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Auto-Fixer" "skip" "skipped (--skip-bracecheck)" }
    } elseif ($skipAutofix) {
        Write-Host "  Skipping Go auto-fixer (--no-autofix)" -ForegroundColor DarkYellow
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Auto-Fixer" "skip" "skipped (--no-autofix)" }
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
            $fixStr = ($fixOut | Out-String).Trim() -replace '^\s*✓\s*', ''
            if ($fixStr) { Write-Success $fixStr }
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Auto-Fixer" "pass" "no fixable issues" }
        }
    }

    # ── Go syntax pre-check (bracecheck) ──────────────────────────────
    if ($skipBrace) {
        # already logged above
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Syntax Check" "skip" "skipped (--skip-bracecheck)" }
    } else {
        Write-Host "  Running Go syntax pre-check (bracecheck)..." -ForegroundColor Yellow
        $braceOut = & go run ./scripts/bracecheck/ 2>&1
        if ($LASTEXITCODE -ne 0) {
            Write-Host ($braceOut | Out-String) -ForegroundColor Red
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Syntax Check" "fail" "bracecheck failed" }
            Write-Fail "Go syntax check failed. Fix reported issues before TC."
            exit 1
        } else {
            $braceStr2 = ($braceOut | Out-String).Trim() -replace '^\s*✓\s*', ''
            Write-Success $braceStr2
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Syntax Check" "pass" $braceStr2 }
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
        # Append autofix section if report already has content from autofix
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

    # ── Pre-coverage compile check ──────────────────────────────────
    # Build-check each test package individually (go test -run '^$') to detect
    # build failures BEFORE running coverage. Packages that fail to
    # compile are excluded from the coverage run so they don't produce
    # misleading 0% profiles or cascade failures.

    # Determine sync vs parallel mode from ExtraArgs (--sync flag)
    $isSyncMode = $false
    if ($ExtraArgs) {
        foreach ($ea in $ExtraArgs) {
            if ($ea -eq "--sync") { $isSyncMode = $true }
        }
    }

    $modeLabel = if ($isSyncMode) { "sync" } else { "parallel" }
    Write-Host ""
    Write-Header "Pre-coverage compile check ($($allTestPkgs.Count) packages, $modeLabel mode)"

    $blockedPkgs = [System.Collections.Generic.List[string]]::new()
    $blockedErrors = [System.Collections.Generic.Dictionary[string, string]]::new()
    $testPkgs = [System.Collections.Generic.List[string]]::new()
    $buildErrorsByPackage = @{}
    $runtimeFailuresByPackage = @{}

    if ($isSyncMode) {
        # ── Sequential compile check ──
        foreach ($testPkg in $allTestPkgs) {
            $shortName = $testPkg -replace '.*integratedtests/?', ''
            if (-not $shortName) { $shortName = "(root)" }

            $prevPref = $ErrorActionPreference
            $ErrorActionPreference = "Continue"
            $compileOut = & go test -count=1 -run '^$' -gcflags=all=-e "-coverpkg=$covPkgList" "$testPkg" 2>&1 | ForEach-Object { $_.ToString() }
            $compileExit = $LASTEXITCODE
            $ErrorActionPreference = $prevPref

            if ($compileExit -eq 0) {
                $testPkgs.Add($testPkg)
            } else {
                $prevPref = $ErrorActionPreference
                $ErrorActionPreference = "Continue"
                $diagOut = & go test -count=1 -run '^$' -gcflags=all=-e "$testPkg" 2>&1 | ForEach-Object { $_.ToString() }
                $ErrorActionPreference = $prevPref
                $combinedOut = Merge-UniqueOutputLines $compileOut $diagOut

                $blockedPkgs.Add($shortName)
                $blockedErrors[$shortName] = ($combinedOut -join "`n")
                Add-BuildErrorsForPackage $buildErrorsByPackage $shortName $combinedOut
                Add-RuntimeFailuresForPackage $runtimeFailuresByPackage $shortName $combinedOut
            }
        }
    } else {
        # ── Parallel compile check (ForEach-Object -Parallel, runspace-based) ──
        $throttle = [Math]::Min($allTestPkgs.Count, [Environment]::ProcessorCount * 2)
        Write-Host "  Launching $($allTestPkgs.Count) compile checks ($throttle parallel)..." -ForegroundColor Gray

        $compileResults = $allTestPkgs | ForEach-Object -ThrottleLimit $throttle -Parallel {
            $pkg = $_
            $covPkgs = $using:covPkgList
            $ErrorActionPreference = "Continue"
            $rawOut = & go test -count=1 -run '^$' -gcflags=all=-e "-coverpkg=$covPkgs" "$pkg" 2>&1
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

            [pscustomobject]@{
                Pkg      = $pkg
                ExitCode = $ec
                Output   = $out
            }
        }

        foreach ($result in ($compileResults | Sort-Object Pkg)) {
            $shortName = $result.Pkg -replace '.*integratedtests/?', ''
            if (-not $shortName) { $shortName = "(root)" }

            if ($result.ExitCode -eq 0) {
                $testPkgs.Add($result.Pkg)
            } else {
                $blockedPkgs.Add($shortName)
                $blockedErrors[$shortName] = ($result.Output -join "`n")
                Add-BuildErrorsForPackage $buildErrorsByPackage $shortName $result.Output
                Add-RuntimeFailuresForPackage $runtimeFailuresByPackage $shortName $result.Output
            }
        }
    }

    # ── Per-file split for blocked packages ─────────────────────────
    # When a package fails to compile, split each *_test.go into its own
    # subfolder (with shared support files), recheck each independently,
    # and promote passing subfolders into the coverage run. This prevents
    # one broken file from blocking the entire package (e.g. corestrtests
    # with 178 files / 14K tests).

    $splitRecoveredCount = 0
    $splitBlockedFiles = [System.Collections.Generic.List[string]]::new()
    $splitCleanupDirs = [System.Collections.Generic.List[string]]::new()

    if ($blockedPkgs.Count -gt 0) {
        $blockedSnapshot = @($blockedPkgs)
        foreach ($bp in $blockedSnapshot) {
            $pkgDir = Join-Path "tests" "integratedtests" $bp
            if (-not (Test-Path $pkgDir)) { continue }

            $bpTestFiles = Get-ChildItem -LiteralPath $pkgDir -Filter "*_test.go" -File |
                Where-Object { $_.Name -notlike "*helper*" } | Sort-Object Name
            $bpHelperTestFiles = Get-ChildItem -LiteralPath $pkgDir -Filter "*helper*_test.go" -File | Sort-Object Name
            $bpSupportFiles = Get-ChildItem -LiteralPath $pkgDir -Filter "*.go" -File |
                Where-Object { $_.Name -notlike "*_test.go" } | Sort-Object Name

            if ($bpTestFiles.Count -lt 2) { continue }  # no benefit in splitting a single file

            Write-Host ""
            Write-Host "  Splitting $bp ($($bpTestFiles.Count) test files) for per-file recheck..." -ForegroundColor Yellow

            $subfolderResults = [System.Collections.Generic.List[pscustomobject]]::new()

            foreach ($tf in $bpTestFiles) {
                $folderName = $tf.BaseName -replace '_test$', ''
                $dest = Join-Path $pkgDir $folderName
                if (-not (Test-Path $dest)) { New-Item -ItemType Directory -Path $dest -Force | Out-Null }
                Copy-Item -LiteralPath $tf.FullName -Destination (Join-Path $dest $tf.Name) -Force
                foreach ($sf in $bpSupportFiles) {
                    Copy-Item -LiteralPath $sf.FullName -Destination (Join-Path $dest $sf.Name) -Force
                }
                foreach ($hf in $bpHelperTestFiles) {
                    Copy-Item -LiteralPath $hf.FullName -Destination (Join-Path $dest $hf.Name) -Force
                }
                $splitCleanupDirs.Add($dest)
            }

            # Compile-check each subfolder
            $subDirs = Get-ChildItem -LiteralPath $pkgDir -Directory | Sort-Object Name

            if ($isSyncMode) {
                foreach ($sd in $subDirs) {
                    $subPkg = "./tests/integratedtests/$bp/$($sd.Name)/"
                    $prevPref = $ErrorActionPreference
                    $ErrorActionPreference = "Continue"
                    $subOut = & go test -count=1 -run '^$' -gcflags=all=-e "-coverpkg=$covPkgList" "$subPkg" 2>&1 | ForEach-Object { $_.ToString() }
                    $subExit = $LASTEXITCODE
                    $ErrorActionPreference = $prevPref
                    $subfolderResults.Add([pscustomobject]@{
                        Name     = $sd.Name
                        Pkg      = $subPkg
                        ExitCode = $subExit
                        Output   = $subOut
                    })
                }
            } else {
                $throttle = [Math]::Min($subDirs.Count, [Environment]::ProcessorCount * 2)
                $parallelResults = $subDirs | ForEach-Object -ThrottleLimit $throttle -Parallel {
                    $sd = $_
                    $bpName = $using:bp
                    $covPkgs = $using:covPkgList
                    $subPkg = "./tests/integratedtests/$bpName/$($sd.Name)/"
                    $ErrorActionPreference = "Continue"
                    $subOut = & go test -count=1 -run '^$' -gcflags=all=-e "-coverpkg=$covPkgs" "$subPkg" 2>&1 | ForEach-Object { $_.ToString() }
                    [pscustomobject]@{
                        Name     = $sd.Name
                        Pkg      = $subPkg
                        ExitCode = $LASTEXITCODE
                        Output   = $subOut
                    }
                }
                foreach ($pr in ($parallelResults | Sort-Object Name)) {
                    $subfolderResults.Add($pr)
                }
            }

            # Tally results
            $subPass = @()
            $subFail = @()
            foreach ($sr in $subfolderResults) {
                if ($sr.ExitCode -eq 0) {
                    $subPass += $sr
                } else {
                    $subFail += $sr
                }
            }

            Write-Host "    ✓ $($subPass.Count) subfolders compile OK" -ForegroundColor Green
            if ($subFail.Count -gt 0) {
                Write-Host "    ✗ $($subFail.Count) subfolders failed:" -ForegroundColor Red
                foreach ($sf in $subFail) {
                    Write-Host "      ✗ $($sf.Name)" -ForegroundColor Red
                    $splitBlockedFiles.Add("$bp/$($sf.Name)")
                }
            }

            # Promote passing subfolders into testPkgs for coverage
            foreach ($sp in $subPass) {
                # Resolve full Go package path for the subfolder
                $fullSubPkg = $allTestPkgs | Where-Object { $_ -match "integratedtests/$bp$" } | Select-Object -First 1
                if ($fullSubPkg) {
                    $subFullPkg = $fullSubPkg + "/" + $sp.Name
                } else {
                    # Fallback: construct from relative path
                    $subFullPkg = $sp.Pkg
                }
                $testPkgs.Add($subFullPkg)
                $splitRecoveredCount++
            }

            # Remove the original blocked package entry (it's now replaced by subfolders)
            $blockedPkgs.Remove($bp)
            $blockedErrors.Remove($bp)

            # Add individual file failures to blocked tracking
            foreach ($sf in $subFail) {
                $failName = "$bp/$($sf.Name)"
                $blockedPkgs.Add($failName)
                $blockedErrors[$failName] = ($sf.Output -join "`n")
                Add-BuildErrorsForPackage $buildErrorsByPackage $failName $sf.Output
            }
        }

        if ($splitRecoveredCount -gt 0) {
            Write-Host ""
            Write-Success "Recovered $splitRecoveredCount subfolders from blocked packages via per-file split"
        }
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
            Register-Phase "Split Recovery" "pass" "$splitRecoveredCount subfolders recovered"
        }
    } else {
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
            Register-Phase "Split Recovery" "skip" "not needed"
        }
    }

    # Print blocked summary
    if ($blockedPkgs.Count -gt 0) {
        Write-Host ""
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Red
        Write-Host "  │ BLOCKED PACKAGES ($($blockedPkgs.Count) failed to compile)" -ForegroundColor Red
        Write-Host "  │" -ForegroundColor Red
        foreach ($bp in ($blockedPkgs | Sort-Object)) {
            Write-Host "  │   ✗ $bp" -ForegroundColor Red
        }
        Write-Host "  │" -ForegroundColor Red
        Write-Host "  │ These packages will be SKIPPED in coverage." -ForegroundColor Yellow
        Write-Host "  │ Fix their build errors to include them." -ForegroundColor Yellow
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Red
        Write-Host ""

        # Write blocked details to file for AI/human review
        $blockedFile = Join-Path $coverDir "blocked-packages.txt"
        $sortedBlocked = $blockedPkgs | Sort-Object
        $blockedContent = @(
            "# Blocked Packages — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')",
            "# Count: $($blockedPkgs.Count)",
            "",
            "# ── CLI Summary ──",
            "# ┌─────────────────────────────────────────────────",
            "# │ BLOCKED PACKAGES ($($blockedPkgs.Count) failed to compile)"
        )
        foreach ($bp in $sortedBlocked) {
            $blockedContent += "# │   ✗ $bp"
        }
        $blockedContent += @(
            "# │",
            "# │ These packages will be SKIPPED in coverage.",
            "# │ Fix their build errors to include them.",
            "# └─────────────────────────────────────────────────",
            "",
            "# ── Package List ──"
        )
        foreach ($bp in $sortedBlocked) {
            $blockedContent += "# - $bp"
        }
        $blockedContent += ""

        # (All output goes to $coverDir only — no root writes)

        foreach ($bp in $sortedBlocked) {
            $blockedContent += "## $bp"
            if ($blockedErrors.ContainsKey($bp)) {
                $rawErrLines = $blockedErrors[$bp] -split "`n"
                $filteredErrLines = Extract-BuildErrorLines $rawErrLines
                if ($filteredErrLines.Count -eq 0) {
                    $filteredErrLines = Extract-ExecutionFailureLines $rawErrLines
                }
                if ($filteredErrLines.Count -gt 0) {
                    $blockedContent += ($filteredErrLines -join "`n")
                } else {
                    $blockedContent += "(no actionable compile errors captured)"
                }
            }
            $blockedContent += ""
        }
        $fileContent = $blockedContent -join "`n"
        Set-Content -Path $blockedFile -Value $fileContent -Encoding UTF8

        # ── JSON export for blocked packages ──
        $blockedJsonFile = Join-Path $coverDir "blocked-packages.json"
        
        $blockedJsonItems = [System.Collections.Generic.List[object]]::new()
        foreach ($bp in $sortedBlocked) {
            $errText = ""
            if ($blockedErrors.ContainsKey($bp)) { $errText = $blockedErrors[$bp] }
            $errLines = @()
            if ($errText) {
                $errLines = Extract-BuildErrorLines ($errText -split "`n")
                if ($errLines.Count -eq 0) {
                    $errLines = Extract-ExecutionFailureLines ($errText -split "`n")
                }
            }
            $blockedJsonItems.Add(@{
                package    = $bp
                errorCount = $errLines.Count
                errors     = $errLines
            })
        }
        $blockedJsonObj = @{
            timestamp        = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ")
            blockedCount     = $blockedPkgs.Count
            compiledCount    = $testPkgs.Count
            totalCount       = $allTestPkgs.Count
            blockedPackages  = $blockedJsonItems.ToArray()
            missingProfiles  = @()
        }
        $blockedJson = $blockedJsonObj | ConvertTo-Json -Depth 4
        Set-Content -Path $blockedJsonFile -Value $blockedJson -Encoding UTF8
    } else {
        Write-Host ""
        Write-Success "All $($testPkgs.Count) packages compiled successfully"
    }
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
        if ($blockedPkgs.Count -gt 0) {
            Register-Phase "Compile Check" "warn" "$($testPkgs.Count)/$($allTestPkgs.Count) passed, $($blockedPkgs.Count) blocked"
        } else {
            Register-Phase "Compile Check" "pass" "$($testPkgs.Count)/$($allTestPkgs.Count) passed"
        }
    }

    if ($testPkgs.Count -eq 0) {
        Write-Fail "No packages compiled — aborting coverage run"
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Coverage Run" "fail" "no packages to run" }
        return
    }

    # ── Coverage run (only compilable packages) ─────────────────────
    $allOutput = [System.Collections.Generic.List[string]]::new()
    $pkgCoverMap = [ordered]@{}
    $overallExit = 0
    $pkgIndex = 0

    Write-Host ""
    Write-Host "  Running $($testPkgs.Count) test packages ($modeLabel)..." -ForegroundColor Yellow

    if ($isSyncMode) {
        # ── Sequential coverage run ──
        foreach ($testPkg in $testPkgs) {
            $pkgIndex++
            $shortName = $testPkg -replace '.*integratedtests/?', ''
            if (-not $shortName) { $shortName = "(root)" }
            $srcTarget = $shortName -replace 'tests$', '' -replace 'tests/', '/'
            if (-not $srcTarget) { $srcTarget = $shortName }

            $partialProfile = Join-Path $partialDir "cover-$pkgIndex.out"

            $prevPref = $ErrorActionPreference
            $ErrorActionPreference = "Continue"
            $output = & go test -count=1 "-coverprofile=$partialProfile" "-coverpkg=$covPkgList" "$testPkg" 2>&1 | ForEach-Object { $_.ToString() }
            $pkgExit = $LASTEXITCODE
            $ErrorActionPreference = $prevPref

            if ($pkgExit -ne 0) {
                $overallExit = $pkgExit
                Add-BuildErrorsForPackage $buildErrorsByPackage $shortName $output
                Add-RuntimeFailuresForPackage $runtimeFailuresByPackage $shortName $output
            }

            if ($output) { foreach ($line in $output) { $allOutput.Add([string]$line) } }
        }
    } else {
        # ── Parallel coverage run (ForEach-Object -Parallel) ──
        $throttle = [Math]::Min($testPkgs.Count, [Environment]::ProcessorCount * 2)

        $coverResults = $testPkgs | ForEach-Object -ThrottleLimit $throttle -Parallel {
            $pkg = $_
            $covPkgs = $using:covPkgList
            $pDir = $using:partialDir
            $safePkgName = $pkg -replace '[^a-zA-Z0-9\.-]', '_'
            $profile = Join-Path $pDir "cover-$safePkgName.out"
            $ErrorActionPreference = "Continue"
            $out = & go test -count=1 "-coverprofile=$profile" "-coverpkg=$covPkgs" "$pkg" 2>&1 | ForEach-Object { $_.ToString() }
            [pscustomobject]@{
                Pkg      = $pkg
                Profile  = $profile
                ExitCode = $LASTEXITCODE
                Output   = $out
            }
        }

        foreach ($result in ($coverResults | Sort-Object Pkg)) {
            $shortName = $result.Pkg -replace '.*integratedtests/?', ''
            if (-not $shortName) { $shortName = "(root)" }

            if ($result.ExitCode -ne 0) {
                $overallExit = $result.ExitCode
                Add-BuildErrorsForPackage $buildErrorsByPackage $shortName $result.Output
                Add-RuntimeFailuresForPackage $runtimeFailuresByPackage $shortName $result.Output
            }
            if ($result.Output) { foreach ($line in $result.Output) { $allOutput.Add([string]$line) } }
        }
        $pkgIndex = $testPkgs.Count
    }

    # ── Safety Guard: detect missing coverage profiles (binary crash) ──
    # When a test binary panics/crashes, Go never writes the coverage profile.
    # This silently drops that package's coverage, producing misleadingly low %.
    $missingProfiles = [System.Collections.Generic.List[string]]::new()
    foreach ($testPkg in $testPkgs) {
        # Determine which profile path was used
        if ($isSyncMode) {
            $idx = [array]::IndexOf($testPkgs, $testPkg) + 1
            $expectedProfile = Join-Path $partialDir "cover-$idx.out"
        } else {
            $safeName = $testPkg -replace '[^a-zA-Z0-9\.-]', '_'
            $expectedProfile = Join-Path $partialDir "cover-$safeName.out"
        }
        if (-not (Test-Path $expectedProfile)) {
            $shortPkg = $testPkg -replace '.*integratedtests/?', ''
            if (-not $shortPkg) { $shortPkg = $testPkg }
            $missingProfiles.Add($shortPkg)
        }
    }
    if ($missingProfiles.Count -gt 0) {
        Write-Host ""
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Magenta
        Write-Host "  │ ⚠ MISSING COVERAGE PROFILES ($($missingProfiles.Count) package(s))" -ForegroundColor Magenta
        Write-Host "  │" -ForegroundColor Magenta
        Write-Host "  │ These test binaries likely crashed (panic/os.Exit)" -ForegroundColor Magenta
        Write-Host "  │ before Go could write their coverage profile." -ForegroundColor Magenta
        Write-Host "  │ Their coverage is NOT included in the report." -ForegroundColor Magenta
        Write-Host "  │" -ForegroundColor Magenta
        foreach ($mp in $missingProfiles) {
            Write-Host "  │   ⚠ $mp" -ForegroundColor Yellow
        }
        Write-Host "  │" -ForegroundColor Magenta
        Write-Host "  │ Fix: ensure tests use recover() for expected panics" -ForegroundColor Magenta
        Write-Host "  │ and never call os.Exit() in test code." -ForegroundColor Magenta
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Magenta
    }

    # ── Backfill missing-profiles into blocked-packages JSON ──
    $blockedJsonFile = Join-Path $coverDir "blocked-packages.json"
    if (Test-Path $blockedJsonFile) {
        $existingJson = Get-Content $blockedJsonFile -Raw | ConvertFrom-Json
        $mpArray = @($missingProfiles | ForEach-Object { $_ })
        $existingJson | Add-Member -NotePropertyName "missingProfileCount" -NotePropertyValue $missingProfiles.Count -Force
        $existingJson | Add-Member -NotePropertyName "missingProfiles" -NotePropertyValue $mpArray -Force
        $existingJson | ConvertTo-Json -Depth 4 | Set-Content -Path $blockedJsonFile -Encoding UTF8
    } elseif ($missingProfiles.Count -gt 0) {
        # No blocked packages but we have missing profiles — create the file
        $mpOnly = @{
            timestamp           = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ")
            blockedCount        = 0
            compiledCount       = $testPkgs.Count
            totalCount          = $allTestPkgs.Count
            blockedPackages     = @()
            missingProfileCount = $missingProfiles.Count
            missingProfiles     = @($missingProfiles | ForEach-Object { $_ })
        }
        $mpOnly | ConvertTo-Json -Depth 4 | Set-Content -Path $blockedJsonFile -Encoding UTF8
    }

    # Write test logs to files (no raw dump to console)
    Write-TestLogs $allOutput.ToArray()

    # ── Failing Test Summary (console) ──
    $failedTestNames = [System.Collections.Generic.HashSet[string]]::new()
    foreach ($line in $allOutput) {
        if ($line -match "--- FAIL:\s+(.+?)\s+\(") {
            $failedTestNames.Add($Matches[1].Trim()) | Out-Null
        }
    }
    if ($failedTestNames.Count -gt 0) {
        Write-Host ""
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Red
        Write-Host "  │ FAILING TESTS ($($failedTestNames.Count) failed)" -ForegroundColor Red
        Write-Host "  │" -ForegroundColor Red
        foreach ($ft in ($failedTestNames | Sort-Object)) {
            Write-Host "  │   ✗ $ft" -ForegroundColor Red
        }
        Write-Host "  │" -ForegroundColor Red
        Write-Host "  │ See data/test-logs/failing-tests.txt for details." -ForegroundColor Yellow
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Red
    }

    # Merge all partial profiles into one, using MAX count per unique line.
    # -coverpkg instruments ALL source packages in every test run, so each line
    # appears N times. Without dedup, the last count=0 overwrites covered entries.

    $partialFiles = Get-ChildItem -Path $partialDir -Filter "cover-*.out" | Sort-Object Name
    $coverMap = [System.Collections.Generic.Dictionary[string, int]]::new()

    foreach ($pf in $partialFiles) {
        $lines = Get-Content $pf.FullName
        foreach ($line in $lines) {
            if (-not $line -or $line -match "^mode:") { continue }
            # Coverage line format: "pkg/file.go:startLine.startCol,endLine.endCol numStatements count"
            # Require full format with colon before line numbers to reject malformed lines
            if ($line -match "^(\S+\.go:\d+\.\d+,\d+\.\d+\s+\d+)\s+(\d+)\s*$") {
                $key = $Matches[1]
                $count = [int]$Matches[2]
                if ($coverMap.ContainsKey($key)) {
                    if ($count -gt $coverMap[$key]) {
                        $coverMap[$key] = $count
                    }
                } else {
                    $coverMap[$key] = $count
                }
            }
        }
    }

    $mergedLines = [System.Collections.Generic.List[string]]::new()
    $mergedLines.Add("mode: set")
    foreach ($entry in $coverMap.GetEnumerator()) {
        $mergedLines.Add("$($entry.Key) $($entry.Value)")
    }

    Set-Content -Path $coverProfile -Value ($mergedLines -join "`n") -Encoding UTF8
    # (file write messages deferred to written files summary)

    if (Test-Path $coverProfile) {
        # Generate func-level summary
        # Generate func-level summary (no debug output)

        $funcOutput = & go tool cover "-func=$coverProfile" 2>&1 | ForEach-Object { $_.ToString() }

        $uncoveredJsonScript = Join-Path $PSScriptRoot "scripts" "coverage" "Export-UncoveredMethodsJson.ps1"
        if (Test-Path $uncoveredJsonScript) {
            $uncoveredJsonFile = Join-Path $coverDir "uncovered-method-lines.json"
            & $uncoveredJsonScript -CoverProfile $coverProfile -FuncOutput $funcOutput -OutputFile $uncoveredJsonFile -ProjectRoot $PSScriptRoot
        }

        # Generate HTML report — use explicit argument list to avoid variable interpolation issues
        $htmlArgs = @("-html=$coverProfile", "-o=$coverHtml")
        
        $htmlErr = & go tool cover $htmlArgs 2>&1
        $htmlExitCode = $LASTEXITCODE

        if ($htmlExitCode -ne 0 -or -not (Test-Path $coverHtml)) {
            Write-Host "  ⚠ Failed to generate HTML report via 'go tool cover -html' (exit: $htmlExitCode)" -ForegroundColor Red
            if ($htmlErr) { Write-Host "  Error: $htmlErr" -ForegroundColor Red }
            
            # Fallback: generate a basic HTML from the func output
            $fallbackHtml = @"
<!DOCTYPE html><html><head><meta charset="utf-8"><title>Coverage Report</title>
<style>body{font-family:monospace;padding:20px;background:#1e1e2e;color:#cdd6f4}
pre{white-space:pre-wrap}</style></head><body>
<h1>Coverage Report</h1><pre>$($funcOutput -join "`n")</pre></body></html>
"@
            Set-Content -Path $coverHtml -Value $fallbackHtml -Encoding UTF8
            Write-Host "  Generated fallback HTML report" -ForegroundColor Yellow
        }

        # Build AI-friendly coverage text for the copy button
        $aiTextLines = [System.Collections.Generic.List[string]]::new()
        $aiTextLines.Add("# Coverage Report — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')")
        $aiTextLines.Add("")

        # Build summary report
        $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
        $summaryLines = [System.Collections.Generic.List[string]]::new()
        $summaryLines.Add("# Coverage Summary — $timestamp")
        $summaryLines.Add("")

        # Extract total line
        $totalLine = $funcOutput | Where-Object { $_ -match "^total:" } | Select-Object -Last 1
        if ($totalLine) {
            $summaryLines.Add("## Total Coverage")
            $summaryLines.Add("  $totalLine")
            $summaryLines.Add("")
        }

        # Per-SOURCE-package coverage from merged profile
        $srcPkgCovMap = [ordered]@{}
        foreach ($line in $funcOutput) {
            if ($line -match "^(\S+):" -and $line -notmatch "^total:") {
                $srcPkg = $Matches[1]
                # Extract short name from full import path
                $shortSrc = $srcPkg -replace '.*alimtvnetwork/core/?', ''
                if (-not $shortSrc) { $shortSrc = "(root)" }
                if (-not $srcPkgCovMap.Contains($shortSrc)) {
                    $srcPkgCovMap[$shortSrc] = @{ Stmts = 0; Covered = 0 }
                }
            }
        }
        # Parse coverage.out lines to compute per-source-package %
        $srcPkgStmts = [ordered]@{}
        foreach ($covLine2 in $mergedLines) {
            if ($covLine2 -match "^mode:") { continue }
            # Format: pkg/file.go:startLine.col,endLine.col numStmts count
            if ($covLine2 -match "^(\S+?):(\d+)\.(\d+),(\d+)\.(\d+)\s+(\d+)\s+(\d+)") {
                $filePath2 = $Matches[1]
                $stmts = [int]$Matches[6]
                $count = [int]$Matches[7]
                # Extract package from file path
                $shortSrc2 = $filePath2 -replace '.*alimtvnetwork/core/?', ''
                $shortSrc2 = $shortSrc2 -replace '/[^/]+$', ''  # remove filename
                if (-not $shortSrc2) { $shortSrc2 = "(root)" }
                if (-not $srcPkgStmts.Contains($shortSrc2)) {
                    $srcPkgStmts[$shortSrc2] = @{ Stmts = 0; Covered = 0 }
                }
                $srcPkgStmts[$shortSrc2].Stmts += $stmts
                if ($count -gt 0) { $srcPkgStmts[$shortSrc2].Covered += $stmts }
            }
        }
        if ($srcPkgStmts.Count -gt 0) {
            $summaryLines.Add("## Per-Package Coverage (Source)")
            $sortedSrcPkgs = $srcPkgStmts.GetEnumerator() | ForEach-Object {
                $pctVal = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
                [pscustomobject]@{ Name = $_.Key; Pct = $pctVal }
            } | Sort-Object Pct -Descending
            foreach ($entry in $sortedSrcPkgs) {
                $summaryLines.Add("  $($entry.Pct)%`t$($entry.Name)")
            }
            $summaryLines.Add("")
        }

        # Extract low-coverage functions (< 50%)
        $lowCovFuncs = [System.Collections.Generic.List[string]]::new()
        foreach ($line in $funcOutput) {
            if ($line -match "(\d+\.\d+)%\s*$" -and $line -notmatch "^total:") {
                $pct = [double]$Matches[1]
                if ($pct -lt 50.0) {
                    $lowCovFuncs.Add("  $line")
                }
            }
        }

        if ($lowCovFuncs.Count -gt 0) {
            $summaryLines.Add("## Low Coverage Functions (< 50%)")
            $summaryLines.Add("  Count: $($lowCovFuncs.Count)")
            $summaryLines.Add("")
            foreach ($f in $lowCovFuncs) { $summaryLines.Add($f) }
            $summaryLines.Add("")
        }

        # File paths
        $summaryLines.Add("## Reports")
        $summaryLines.Add("  Profile:  $coverProfile")
        $summaryLines.Add("  HTML:     $coverHtml")
        $summaryLines.Add("  Summary:  $coverSummary")

        Set-Content -Path $coverSummary -Value ($summaryLines -join "`n") -Encoding UTF8

        # ── JSON export for coverage summary ──
        $coverJsonFile = Join-Path $coverDir "coverage-summary.json"
        

        # Parse total coverage
        $totalPct = 0.0
        if ($totalLine -match "(\d+\.\d+)%") { $totalPct = [double]$Matches[1] }

        # Build per-package array sorted by coverage ascending (lowest first for AI prioritization)
        $pkgJsonItems = [System.Collections.Generic.List[object]]::new()
        if ($srcPkgStmts.Count -gt 0) {
            $sortedForJson = $srcPkgStmts.GetEnumerator() | ForEach-Object {
                $pctJ = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
                [pscustomobject]@{ Name = $_.Key; Pct = $pctJ; Stmts = $_.Value.Stmts; Covered = $_.Value.Covered }
            } | Sort-Object Pct
            foreach ($e in $sortedForJson) {
                $pkgJsonItems.Add(@{
                    package    = $e.Name
                    coverage   = $e.Pct
                    statements = $e.Stmts
                    covered    = $e.Covered
                    uncovered  = $e.Stmts - $e.Covered
                })
            }
        }

        # Build low-coverage functions array
        $lowCovJsonItems = [System.Collections.Generic.List[object]]::new()
        foreach ($line in $funcOutput) {
            if ($line -match "(\d+\.\d+)%\s*$" -and $line -notmatch "^total:") {
                $pctF = [double]$Matches[1]
                if ($pctF -lt 50.0) {
                    # Parse: pkg/file.go:line: FuncName    pct%
                    $funcName = ""
                    $funcFile = ""
                    if ($line -match "^(\S+):\s+(\S+)\s+(\d+\.\d+)%") {
                        $funcFile = $Matches[1]
                        $funcName = $Matches[2]
                    }
                    $lowCovJsonItems.Add(@{
                        file     = $funcFile
                        function = $funcName
                        coverage = $pctF
                    })
                }
            }
        }

        # Check for blocked packages JSON
        $blockedRef = @()
        $blockedJsonPath = Join-Path $coverDir "blocked-packages.json"
        if (Test-Path $blockedJsonPath) {
            $blockedRef = @((Get-Content $blockedJsonPath -Raw | ConvertFrom-Json).blockedPackages | ForEach-Object { $_.package })
        }

        $coverJsonObj = @{
            timestamp           = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ")
            totalCoverage       = $totalPct
            packageCount        = $pkgJsonItems.Count
            packages            = $pkgJsonItems.ToArray()
            lowCoverageFuncCount = $lowCovJsonItems.Count
            lowCoverageFunctions = $lowCovJsonItems.ToArray()
            blockedPackages     = $blockedRef
            reports             = @{
                profile = $coverProfile
                html    = $coverHtml
                summary = $coverSummary
                json    = $coverJsonFile
            }
        }
        $coverJson = $coverJsonObj | ConvertTo-Json -Depth 4
        Set-Content -Path $coverJsonFile -Value $coverJson -Encoding UTF8
        

        # ── Per-Package Coverage report (TXT + JSON) ──
        $perPkgTxtFile = Join-Path $coverDir "per-package-coverage.txt"
        $perPkgJsonFile = Join-Path $coverDir "per-package-coverage.json"

        $perPkgTxtLines = [System.Collections.Generic.List[string]]::new()
        $perPkgTxtLines.Add("# Per-Package Coverage Report — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')")
        $perPkgTxtLines.Add("# Total: $totalLine")
        $perPkgTxtLines.Add("")
        $perPkgTxtLines.Add(("Package".PadRight(50)) + " " + ("Stmts".PadLeft(8)) + " " + ("Covered".PadLeft(8)) + " " + ("Uncovered".PadLeft(10)) + " " + ("Cov%".PadLeft(8)))
        $perPkgTxtLines.Add(("─" * 50) + " " + ("─" * 8) + " " + ("─" * 8) + " " + ("─" * 10) + " " + ("─" * 8))

        $perPkgJsonItems = [System.Collections.Generic.List[object]]::new()

        if ($srcPkgStmts.Count -gt 0) {
            $sortedPerPkg = $srcPkgStmts.GetEnumerator() | ForEach-Object {
                $pctP = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
                [pscustomobject]@{
                    Name      = $_.Key
                    Pct       = $pctP
                    Stmts     = $_.Value.Stmts
                    Covered   = $_.Value.Covered
                    Uncovered = $_.Value.Stmts - $_.Value.Covered
                }
            } | Sort-Object Pct

            foreach ($pp in $sortedPerPkg) {
                $statusMark = if ($pp.Pct -ge 100) { "✓" } elseif ($pp.Pct -ge 80) { "○" } else { "✗" }
                $rowPackage = ("$statusMark $($pp.Name)").PadRight(50)
                $rowStmts = $pp.Stmts.ToString().PadLeft(8)
                $rowCovered = $pp.Covered.ToString().PadLeft(8)
                $rowUncovered = $pp.Uncovered.ToString().PadLeft(10)
                $rowPct = (([string]::Format([System.Globalization.CultureInfo]::InvariantCulture, "{0:0.0}", $pp.Pct)) + "%").PadLeft(8)
                $perPkgTxtLines.Add("$rowPackage $rowStmts $rowCovered $rowUncovered $rowPct")

                $perPkgJsonItems.Add(@{
                    package    = $pp.Name
                    coverage   = $pp.Pct
                    statements = $pp.Stmts
                    covered    = $pp.Covered
                    uncovered  = $pp.Uncovered
                    status     = if ($pp.Pct -ge 100) { "full" } elseif ($pp.Pct -ge 80) { "good" } else { "low" }
                })
            }

            # Summary footer
            $totalStmts = ($sortedPerPkg | Measure-Object -Property Stmts -Sum).Sum
            $totalCovered = ($sortedPerPkg | Measure-Object -Property Covered -Sum).Sum
            $totalUncovered = $totalStmts - $totalCovered
            $fullCount = ($sortedPerPkg | Where-Object { $_.Pct -ge 100 }).Count
            $lowCount = ($sortedPerPkg | Where-Object { $_.Pct -lt 80 }).Count

            $perPkgTxtLines.Add("")
            $perPkgTxtLines.Add("# Summary")
            $perPkgTxtLines.Add("#   Packages:  $($sortedPerPkg.Count)")
            $perPkgTxtLines.Add("#   100%:      $fullCount")
            $perPkgTxtLines.Add("#   < 80%:     $lowCount")
            $perPkgTxtLines.Add("#   Total stmts: $totalStmts  covered: $totalCovered  uncovered: $totalUncovered")
        }

        $perPkgJsonObj = @{
            timestamp     = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ")
            totalCoverage = $totalPct
            packageCount  = $perPkgJsonItems.Count
            packages      = $perPkgJsonItems.ToArray()
        }

        Set-Content -Path $perPkgTxtFile -Value ($perPkgTxtLines -join "`n") -Encoding UTF8
        $perPkgJson = $perPkgJsonObj | ConvertTo-Json -Depth 4
        Set-Content -Path $perPkgJsonFile -Value $perPkgJson -Encoding UTF8

        # Build AI-friendly text for copy button
        $aiTextLines.Add("## Goal: Improve test coverage for the packages listed below.")
        $aiTextLines.Add("Please write tests for uncovered functions, following the project's AAA pattern.")
        $aiTextLines.Add("")
        if ($totalLine) {
            $aiTextLines.Add("## Total Coverage")
            $aiTextLines.Add($totalLine)
            $aiTextLines.Add("")
        }
        if ($srcPkgStmts.Count -gt 0) {
            $aiTextLines.Add("## Per-Source-Package Coverage")
            $computedSrcPkgs = $srcPkgStmts.GetEnumerator() | ForEach-Object {
                $pctVal3 = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
                [pscustomobject]@{ Name = $_.Key; Pct = $pctVal3; Stmts = $_.Value.Stmts; Covered = $_.Value.Covered }
            } | Sort-Object Pct
            foreach ($e in $computedSrcPkgs) {
                $aiTextLines.Add("  $($e.Pct)%  $($e.Name)  ($($e.Covered)/$($e.Stmts) stmts)")
            }
            $aiTextLines.Add("")
        }
        if ($lowCovFuncs.Count -gt 0) {
            $aiTextLines.Add("## Uncovered/Low-Coverage Functions (< 50%)")
            $aiTextLines.Add("Count: $($lowCovFuncs.Count)")
            $aiTextLines.Add("")
            foreach ($f in $lowCovFuncs) { $aiTextLines.Add($f.TrimStart()) }
            $aiTextLines.Add("")
        }
        $aiTextLines.Add("## Instructions")
        $aiTextLines.Add("- Tests go in tests/integratedtests/{pkg}tests/")
        $aiTextLines.Add("- Use CaseV1 table-driven pattern with AAA comments")
        $aiTextLines.Add("- Focus on the lowest coverage packages first")

        $aiTextEscaped = ($aiTextLines -join "`n") -replace '\\', '\\\\' -replace "'", "\\\'" -replace "`n", '\n' -replace "`r", '' -replace '"', '\"'

        # Inject "Copy for AI" button into the Go HTML report
        if (Test-Path $coverHtml) {
            $htmlContent = Get-Content -Path $coverHtml -Raw

            $buttonHtml = @'
<div id="ai-copy-panel" style="position:fixed;top:12px;right:12px;z-index:9999;font-family:system-ui,sans-serif;">
<button onclick="copyForAI()" style="
  background:linear-gradient(135deg,#6366f1,#8b5cf6);color:#fff;border:none;
  padding:10px 20px;border-radius:8px;font-size:14px;font-weight:600;
  cursor:pointer;box-shadow:0 4px 12px rgba(99,102,241,0.4);
  display:flex;align-items:center;gap:6px;transition:all 0.2s;
" onmouseover="this.style.transform='scale(1.05)'" onmouseout="this.style.transform='scale(1)'">
  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="9" y="9" width="13" height="13" rx="2"/><path d="M5 15H4a2 2 0 01-2-2V4a2 2 0 012-2h9a2 2 0 012 2v1"/></svg>
  Copy for AI
</button>
<span id="ai-copy-status" style="display:none;color:#22c55e;font-size:13px;margin-top:4px;text-align:center;">Copied!</span>
</div>
<script>
var __aiCoverageText =
'@
            # Insert the escaped text between the two halves
            $scriptEnd = @'
';
function copyForAI(){
  try {
    var ta = document.createElement("textarea");
    ta.value = __aiCoverageText;
    ta.style.position = "fixed";
    ta.style.left = "-9999px";
    document.body.appendChild(ta);
    ta.select();
    document.execCommand("copy");
    document.body.removeChild(ta);
    var s = document.getElementById("ai-copy-status");
    s.style.display = "block";
    setTimeout(function(){ s.style.display = "none"; }, 2000);
  } catch(e) {
    alert("Copy failed: " + e.message);
  }
}
</script>
'@
            $injectedHtml = $buttonHtml + $aiTextEscaped + $scriptEnd
            $htmlContent = $htmlContent -replace '</body>', ($injectedHtml + "`n</body>")
            Set-Content -Path $coverHtml -Value $htmlContent -Encoding UTF8
            Write-Host "  ✓ Injected 'Copy for AI' button into HTML report" -ForegroundColor Green
        }

        # ── Coverage Summary (console) ──
        if ($srcPkgStmts.Count -gt 0) {
            Write-Host ""
            Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Cyan
            Write-Host "  │ COVERAGE SUMMARY" -ForegroundColor Cyan
            Write-Host "  │" -ForegroundColor Cyan
            $sortedSrcPkgs2 = $srcPkgStmts.GetEnumerator() | ForEach-Object {
                $pctVal2 = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
                [pscustomobject]@{ Name = $_.Key; Pct = $pctVal2 }
            } | Sort-Object Pct -Descending
            foreach ($entry2 in $sortedSrcPkgs2) {
                $color = if ($entry2.Pct -ge 100) { "Green" } elseif ($entry2.Pct -ge 80) { "Yellow" } else { "Red" }
                Write-Host "  │  $($entry2.Pct)%`t$($entry2.Name)" -ForegroundColor $color
            }
            Write-Host "  │" -ForegroundColor Cyan
            if ($totalLine -and $totalLine -match "(\d+\.\d+)%") {
                $totalPctDisplay = $Matches[1]
                Write-Host "  │  total: (statements)  $totalPctDisplay%" -ForegroundColor Cyan
            }
            if ($lowCovFuncs.Count -gt 0) {
                Write-Host "  │  ⚠ $($lowCovFuncs.Count) function(s) below 50% coverage" -ForegroundColor Yellow
            }
            Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Cyan
        }

        # ── Coverage Comparison (dashboard diff) ──
        if (Get-Command Write-CoverageComparison -ErrorAction SilentlyContinue) {
            # Build current coverage array in the format expected by DashboardUI
            $currentCovData = @()
            if ($srcPkgStmts.Count -gt 0) {
                $currentCovData = @($srcPkgStmts.GetEnumerator() | ForEach-Object {
                    $pctCmp = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
                    @{ Package = $_.Key; Coverage = $pctCmp }
                })
            }

            if ($currentCovData.Count -gt 0) {
                # Load previous snapshot
                $previousCovData = $null
                if (Get-Command Load-CoverageSnapshot -ErrorAction SilentlyContinue) {
                    $previousCovData = Load-CoverageSnapshot
                }

                Write-Host ""
                Write-CoverageComparison -Current $currentCovData -Previous $previousCovData

                # Save current as the new snapshot for next run
                if (Get-Command Save-CoverageSnapshot -ErrorAction SilentlyContinue) {
                    Save-CoverageSnapshot -CoverageData $currentCovData
                }
            }
        }

        # ── Written Files Summary (console) ──
        # Consolidate ALL build errors: blocked-package compile errors + coverage-run errors
        $buildErrorsFile = Join-Path $coverDir "build-errors.txt"
        $buildErrorsJsonFile = Join-Path $coverDir "build-errors.json"

        # Merge blocked-package compile errors into buildErrorsByPackage
        if ($blockedPkgs.Count -gt 0) {
            foreach ($bp in ($blockedPkgs | Sort-Object)) {
                if ($blockedErrors.ContainsKey($bp)) {
                    $rawErrLines = $blockedErrors[$bp] -split "`n"
                    $filteredErrLines = Extract-BuildErrorLines $rawErrLines
                    if ($filteredErrLines.Count -eq 0) {
                        $filteredErrLines = Extract-ExecutionFailureLines $rawErrLines
                    }
                    if ($filteredErrLines.Count -gt 0) {
                        if (-not $buildErrorsByPackage.ContainsKey($bp)) {
                            $buildErrorsByPackage[$bp] = [System.Collections.Generic.List[string]]::new()
                        }
                        foreach ($errLine in $filteredErrLines) {
                            if (-not $buildErrorsByPackage[$bp].Contains($errLine)) {
                                $buildErrorsByPackage[$bp].Add($errLine) | Out-Null
                            }
                        }
                    }
                }
            }
        }

        $buildErrorPkgs = @($buildErrorsByPackage.Keys | Sort-Object)

        $buildErrorLines = [System.Collections.Generic.List[string]]::new()
        $buildErrorLines.Add("# Build Errors — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')")
        $buildErrorLines.Add("# Count: $($buildErrorPkgs.Count)")
        $buildErrorLines.Add("")

        $buildErrorJsonItems = [System.Collections.Generic.List[object]]::new()

        if ($buildErrorPkgs.Count -eq 0) {
            $buildErrorLines.Add("No build errors captured.")
        }
        else {
            foreach ($pkgName in $buildErrorPkgs) {
                $pkgLines = @($buildErrorsByPackage[$pkgName])
                $isBlocked = $blockedPkgs.Contains($pkgName)
                $sectionLabel = if ($isBlocked) { "## $pkgName [BLOCKED — compile failure]" } else { "## $pkgName [coverage-run error]" }
                $buildErrorLines.Add($sectionLabel)
                if ($pkgLines.Count -gt 0) {
                    $buildErrorLines.AddRange([string[]]$pkgLines)
                }
                else {
                    $buildErrorLines.Add("(no actionable compile errors captured)")
                }
                $buildErrorLines.Add("")

                $buildErrorJsonItems.Add(@{
                    package    = $pkgName
                    errorCount = $pkgLines.Count
                    errors     = $pkgLines
                    source     = if ($isBlocked) { "compile-check" } else { "coverage-run" }
                }) | Out-Null
            }
        }

        Set-Content -Path $buildErrorsFile -Value ($buildErrorLines -join "`n") -Encoding UTF8
        $buildErrorJsonObj = @{
            timestamp    = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ")
            packageCount = $buildErrorPkgs.Count
            blockedCount = $blockedPkgs.Count
            packages     = $buildErrorJsonItems.ToArray()
        }
        $buildErrorJsonObj | ConvertTo-Json -Depth 5 | Set-Content -Path $buildErrorsJsonFile -Encoding UTF8

        # ── Runtime Failures report (panic/os.Exit/crashes) ──
        $runtimeFailuresFile = Join-Path $coverDir "runtime-failures.txt"
        $runtimeFailuresJsonFile = Join-Path $coverDir "runtime-failures.json"
        $runtimeFailurePkgs = @($runtimeFailuresByPackage.Keys | Sort-Object)

        # Include missing coverage profiles as runtime crashes
        if ($missingProfiles.Count -gt 0) {
            foreach ($mp in $missingProfiles) {
                if (-not $runtimeFailuresByPackage.ContainsKey($mp)) {
                    $runtimeFailuresByPackage[$mp] = [System.Collections.Generic.List[string]]::new()
                }
                $crashMsg = "coverage profile missing — test binary likely crashed (panic/os.Exit)"
                if (-not $runtimeFailuresByPackage[$mp].Contains($crashMsg)) {
                    $runtimeFailuresByPackage[$mp].Add($crashMsg) | Out-Null
                }
            }
            $runtimeFailurePkgs = @($runtimeFailuresByPackage.Keys | Sort-Object)
        }

        $rtLines = [System.Collections.Generic.List[string]]::new()
        $rtLines.Add("# Runtime Failures — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')")
        $rtLines.Add("# Panics, os.Exit, test binary crashes, fatal errors")
        $rtLines.Add("# Count: $($runtimeFailurePkgs.Count) package(s)")
        $rtLines.Add("")

        $rtJsonItems = [System.Collections.Generic.List[object]]::new()

        if ($runtimeFailurePkgs.Count -eq 0) {
            $rtLines.Add("No runtime failures captured.")
        }
        else {
            foreach ($pkgName in $runtimeFailurePkgs) {
                $pkgLines = @($runtimeFailuresByPackage[$pkgName])
                $rtLines.Add("## $pkgName")
                if ($pkgLines.Count -gt 0) {
                    $rtLines.AddRange([string[]]$pkgLines)
                }
                else {
                    $rtLines.Add("(no failure details captured)")
                }
                $rtLines.Add("")

                $rtJsonItems.Add(@{
                    package      = $pkgName
                    failureCount = $pkgLines.Count
                    failures     = $pkgLines
                }) | Out-Null
            }
        }

        Set-Content -Path $runtimeFailuresFile -Value ($rtLines -join "`n") -Encoding UTF8
        $rtJsonObj = @{
            timestamp    = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ")
            packageCount = $runtimeFailurePkgs.Count
            packages     = $rtJsonItems.ToArray()
        }
        $rtJsonObj | ConvertTo-Json -Depth 5 | Set-Content -Path $runtimeFailuresJsonFile -Encoding UTF8

        # ── Runtime Failures console warning ──
        if ($runtimeFailurePkgs.Count -gt 0) {
            Write-Host ""
            Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Magenta
            Write-Host "  │ RUNTIME FAILURES ($($runtimeFailurePkgs.Count) package(s))" -ForegroundColor Magenta
            Write-Host "  │" -ForegroundColor Magenta
            foreach ($rp in $runtimeFailurePkgs) {
                Write-Host "  │   ⚠ $rp" -ForegroundColor Yellow
            }
            Write-Host "  │" -ForegroundColor Magenta
            Write-Host "  │ See data/coverage/runtime-failures.txt for details." -ForegroundColor Yellow
            Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Magenta
        }

        Write-Host ""
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Gray
        Write-Host "  │ WRITTEN FILES" -ForegroundColor Gray
        Write-Host "  │  $coverProfile" -ForegroundColor Gray
        Write-Host "  │  $coverHtml" -ForegroundColor Gray
        Write-Host "  │  $coverSummary" -ForegroundColor Gray
        Write-Host "  │  $coverJsonFile" -ForegroundColor Gray
        Write-Host "  │  $perPkgTxtFile" -ForegroundColor Gray
        Write-Host "  │  $perPkgJsonFile" -ForegroundColor Gray
        Write-Host "  │  $buildErrorsFile" -ForegroundColor Gray
        Write-Host "  │  $buildErrorsJsonFile" -ForegroundColor Gray
        Write-Host "  │  $runtimeFailuresFile" -ForegroundColor Gray
        Write-Host "  │  $runtimeFailuresJsonFile" -ForegroundColor Gray
        if (Test-Path $repoBuildErrorsFile) {
            Write-Host "  │  $repoBuildErrorsFile" -ForegroundColor Gray
        }
        if (Test-Path $repoBuildErrorsJsonFile) {
            Write-Host "  │  $repoBuildErrorsJsonFile" -ForegroundColor Gray
        }
        if ($blockedPkgs.Count -gt 0) {
            $bFile = Join-Path $coverDir "blocked-packages.txt"
            $bJsonFile = Join-Path $coverDir "blocked-packages.json"
            Write-Host "  │  $bFile" -ForegroundColor Gray
            Write-Host "  │  $bJsonFile" -ForegroundColor Gray
        }
        $syntaxIssuesFile = Join-Path $coverDir "syntax-issues.txt"
        if (Test-Path $syntaxIssuesFile) {
            Write-Host "  │  $syntaxIssuesFile" -ForegroundColor Gray
        }
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Gray

        # ── Generate AI coverage prompts ──────────────────────────────
        $promptScript = Join-Path $PSScriptRoot "scripts" "coverage" "Generate-CoveragePrompts.ps1"
        if (Test-Path $promptScript) {
            Write-Host ""
            Write-Header "Generating coverage improvement prompts"
            $promptsDir = Join-Path $PSScriptRoot "data" "prompts"
            & $promptScript -CoverProfile $coverProfile -FuncOutput $funcOutput -OutputDir $promptsDir -BatchSize 500 -ProjectRoot $PSScriptRoot
        }

        # HTML auto-open disabled — use --open flag to open manually
        $openHtml = $false
        if ($ExtraArgs -and $ExtraArgs[0] -eq "--open") { $openHtml = $true }
        if ($openHtml -and (Test-Path $coverHtml)) {
            Write-Host ""
            Write-Host "  Opening HTML coverage report in browser..." -ForegroundColor Yellow
            Start-Process $coverHtml
        }
    }
    Open-FailingTestsIfAny

    # Register coverage phases
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
        Register-Phase "Coverage Run" "pass" "$($testPkgs.Count) packages"
        Register-Phase "Coverage Report" "pass" "generated"
    }

    # ── Cleanup split subfolders ──
    if ($splitCleanupDirs.Count -gt 0) {
        Write-Host ""
        Write-Host "  Cleaning up $($splitCleanupDirs.Count) split subfolders..." -ForegroundColor Gray
        foreach ($cleanDir in $splitCleanupDirs) {
            if (Test-Path $cleanDir) {
                Remove-Item -LiteralPath $cleanDir -Recurse -Force
            }
        }
        Write-Host "  ✓ Split subfolders removed" -ForegroundColor Green
    }

    # ── Dashboard Phase Summary ──
    if (Get-Command Write-PhaseSummaryBox -ErrorAction SilentlyContinue) {
        Write-Host ""
        Write-PhaseSummaryBox
    }
}

function Invoke-PackageTestCoverage {
    <#
    .SYNOPSIS
        Run coverage for a single test package (TCP command).
    .PARAMETER pkg
        The test package directory name under tests/integratedtests/.
    .EXAMPLE
        Invoke-PackageTestCoverage "regexnewtests"
    #>
    param([string]$pkg)

    if (-not $pkg) {
        Write-Fail "Usage: ./run.ps1 TCP <package-name>"
        Write-Host "  Example: ./run.ps1 TCP regexnewtests" -ForegroundColor Gray
        return
    }

    Write-Header "Running coverage for package: $pkg"
    Invoke-FetchLatest

    # Clean data folder before running tests
    $dataDir = Join-Path $PSScriptRoot "data"
    if (Test-Path $dataDir) {
        Remove-Item -Recurse -Force $dataDir
        Write-Host "  Cleaned data/ folder" -ForegroundColor Yellow
    }

    # ── safeTest boundary + empty-if lint check ────────────────────
    $boundaryScript = Join-Path $PSScriptRoot "scripts" "check-safetest-boundaries.ps1"
    if (Test-Path $boundaryScript) {
        Write-Host ""
        Write-Host "  Running safeTest boundary + empty-if lint check..." -ForegroundColor Yellow
        & $boundaryScript
        if ($LASTEXITCODE -ne 0) {
            Write-Fail "safeTest boundary check failed. Fix reported issues before TCP."
            exit 1
        }
    }

    # ── Go auto-fixer ─────────────────────────────────────────────────
    $skipAutofix = $ExtraArgs -and ($ExtraArgs -contains '--no-autofix')
    $skipBrace = $ExtraArgs -and ($ExtraArgs -contains '--skip-bracecheck')
    if ($skipBrace) {
        Write-Host "  Skipping Go auto-fixer and syntax pre-check (--skip-bracecheck)" -ForegroundColor DarkYellow
    } elseif ($skipAutofix) {
        Write-Host "  Skipping Go auto-fixer (--no-autofix)" -ForegroundColor DarkYellow
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
        } else {
            $fixStr = ($fixOut | Out-String).Trim() -replace '^\s*✓\s*', ''
            if ($fixStr) { Write-Success $fixStr }
        }
    }

    # ── Go syntax pre-check (bracecheck) ──────────────────────────────
    if ($skipBrace) {
        # already logged above
    } else {
        Write-Host "  Running Go syntax pre-check (bracecheck)..." -ForegroundColor Yellow
        $braceOut = & go run ./scripts/bracecheck/ 2>&1
        if ($LASTEXITCODE -ne 0) {
            Write-Host ($braceOut | Out-String) -ForegroundColor Red
            Write-Fail "Go syntax check failed. Fix reported issues before TCP."
            exit 1
        } else {
            $braceStr2b = ($braceOut | Out-String).Trim() -replace '^\s*✓\s*', ''
            Write-Success $braceStr2b
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

    # Build check from tests/ dir
    Push-Location tests
    try { if (-not (Invoke-BuildCheck "./integratedtests/$pkg/...")) { return } }
    finally { Pop-Location }

    $coverDir = Join-Path $PSScriptRoot "data" "coverage"
    New-Item -ItemType Directory -Path $coverDir -Force | Out-Null

    $coverProfile = Join-Path $coverDir "coverage-$pkg.out"
    $coverHtml    = Join-Path $coverDir "coverage-$pkg.html"
    $coverSummary = Join-Path $coverDir "coverage-$pkg-summary.txt"

    # Build coverpkg list: all source packages EXCLUDING tests/
    $allPkgs = go list ./... 2>&1 | ForEach-Object { $_.ToString() }
    $srcPkgs = $allPkgs | Where-Object { $_ -notmatch '/tests/' }
    $covPkgList = $srcPkgs -join ","

    # Run from project ROOT so -coverpkg can instrument all source packages
    $prevPref = $ErrorActionPreference
    $ErrorActionPreference = "Continue"
    $output = & go test -v -count=1 "-coverprofile=$coverProfile" "-coverpkg=$covPkgList" "./tests/integratedtests/$pkg/..." 2>&1 | ForEach-Object { $_.ToString() }
    $exitCode = $LASTEXITCODE
    $ErrorActionPreference = $prevPref

    Filter-TestWarnings $output | ForEach-Object { Write-Host $_ }
    Write-TestLogs $output

    if (Test-Path $coverProfile) {
        $funcOutput = & go tool cover "-func=$coverProfile" 2>&1 | ForEach-Object { $_.ToString() }
        $htmlArgs = @("-html=$coverProfile", "-o=$coverHtml")
        & go tool cover $htmlArgs 2>&1 | Out-Null

        $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
        $summaryLines = [System.Collections.Generic.List[string]]::new()
        $summaryLines.Add("# Coverage Summary ($pkg) — $timestamp")
        $summaryLines.Add("")

        $totalLine = $funcOutput | Where-Object { $_ -match "^total:" } | Select-Object -Last 1
        if ($totalLine) {
            $summaryLines.Add("## Total Coverage")
            $summaryLines.Add("  $totalLine")
            $summaryLines.Add("")
        }

        $lowCovFuncs = [System.Collections.Generic.List[string]]::new()
        foreach ($line in $funcOutput) {
            if ($line -match "(\d+\.\d+)%\s*$" -and $line -notmatch "^total:") {
                $pct = [double]$Matches[1]
                if ($pct -lt 50.0) {
                    $lowCovFuncs.Add("  $line")
                }
            }
        }

        if ($lowCovFuncs.Count -gt 0) {
            $summaryLines.Add("## Low Coverage Functions (< 50%)")
            $summaryLines.Add("  Count: $($lowCovFuncs.Count)")
            $summaryLines.Add("")
            foreach ($f in $lowCovFuncs) { $summaryLines.Add($f) }
            $summaryLines.Add("")
        }

        $summaryLines.Add("## Reports")
        $summaryLines.Add("  Profile:  $coverProfile")
        $summaryLines.Add("  HTML:     $coverHtml")
        $summaryLines.Add("  Summary:  $coverSummary")

        Set-Content -Path $coverSummary -Value ($summaryLines -join "`n") -Encoding UTF8

        Write-Host ""
        if ($totalLine -and $totalLine -match "(\d+\.\d+)%") {
            $totalPctDisplay2 = $Matches[1]
            Write-Host "  total: (statements)  $totalPctDisplay2%" -ForegroundColor Cyan
        }
        Write-Host ""
        Write-Success "Coverage profile:  $coverProfile"
        Write-Success "HTML report:       $coverHtml"
        Write-Success "Summary:           $coverSummary"

        if ($lowCovFuncs.Count -gt 0) {
            Write-Host ""
            Write-Host "  ⚠ $($lowCovFuncs.Count) function(s) below 50% coverage" -ForegroundColor Yellow
        }

        # ── Coverage Comparison (dashboard diff) ──
        if (Get-Command Write-CoverageComparison -ErrorAction SilentlyContinue) {
            # Build current coverage array from funcOutput lines
            $currentCovData = @()
            foreach ($fLine in $funcOutput) {
                if ($fLine -match "^(.+?):\s+\S+\s+(\d+\.\d+)%\s*$" -and $fLine -notmatch "^total:") {
                    # Extract source package name from path
                    $fPath = $Matches[1]
                    $fPct = [double]$Matches[2]
                    if ($fPath -match '/([^/]+)/[^/]+$') {
                        $srcPkg = $Matches[1]
                    } else {
                        $srcPkg = $fPath
                    }
                    # We'll aggregate below
                }
            }
            # Simpler: use totalLine for single-package, build from funcOutput aggregated by source pkg
            $srcPkgMap = @{}
            foreach ($fLine in $funcOutput) {
                if ($fLine -match "^(\S+):\s+(\S+)\s+(\d+\.\d+)%\s*$" -and $fLine -notmatch "^total:") {
                    $filePath = $Matches[1]
                    $fPct = [double]$Matches[3]
                    # Extract package from file path (e.g. github.com/org/repo/pkg/file.go → pkg)
                    $pathParts = $filePath -split '/'
                    $srcPkg = $pathParts[-2]  # directory containing the file
                    if (-not $srcPkgMap.ContainsKey($srcPkg)) {
                        $srcPkgMap[$srcPkg] = [System.Collections.Generic.List[double]]::new()
                    }
                    $srcPkgMap[$srcPkg].Add($fPct)
                }
            }
            $currentCovData = @($srcPkgMap.GetEnumerator() | ForEach-Object {
                $avg = ($_.Value | Measure-Object -Average).Average
                @{ Package = $_.Key; Coverage = [math]::Round($avg, 1) }
            })

            if ($currentCovData.Count -gt 0) {
                $previousCovData = $null
                if (Get-Command Load-CoverageSnapshot -ErrorAction SilentlyContinue) {
                    $previousCovData = Load-CoverageSnapshot
                }

                Write-Host ""
                Write-CoverageComparison -Current $currentCovData -Previous $previousCovData

                if (Get-Command Save-CoverageSnapshot -ErrorAction SilentlyContinue) {
                    Save-CoverageSnapshot -CoverageData $currentCovData
                }
            }
        }

        # ── Generate AI coverage prompts (per-package) ──────────────
        $promptScript = Join-Path $PSScriptRoot "scripts" "coverage" "Generate-CoveragePrompts.ps1"
        if (Test-Path $promptScript) {
            Write-Host ""
            Write-Header "Generating coverage improvement prompts"
            $promptsDir = Join-Path $PSScriptRoot "data" "prompts"
            & $promptScript -CoverProfile $coverProfile -FuncOutput $funcOutput -OutputDir $promptsDir -BatchSize 500 -ProjectRoot $PSScriptRoot
        }

        # HTML auto-open disabled — use --open flag to open manually
        $openHtml = $false
        if ($ExtraArgs -and $ExtraArgs[-1] -eq "--open") { $openHtml = $true }
        if ($openHtml -and (Test-Path $coverHtml)) {
            Write-Host ""
            Write-Host "  Opening HTML coverage report..." -ForegroundColor Yellow
            Start-Process $coverHtml
        }
    }
    Open-FailingTestsIfAny
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Invoke-TestCoverage',
    'Invoke-PackageTestCoverage'
)
