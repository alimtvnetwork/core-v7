# ─────────────────────────────────────────────────────────────────────────────
# CoverageCompileCheck.psm1 — Pre-coverage compile checks & split recovery
#
# Usage:
#   Import-Module ./scripts/CoverageCompileCheck.psm1 -Force
#
# Dependencies: Utilities.psm1 (Merge-UniqueOutputLines, Add-BuildErrorsForPackage,
#               Add-RuntimeFailuresForPackage, Extract-BuildErrorLines,
#               Extract-ExecutionFailureLines, Write-Header, Write-Success, Write-Fail)
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-CoverageCompileCheck {
    <#
    .SYNOPSIS
        Build-check each test package individually before coverage run.
    .DESCRIPTION
        Returns a hashtable with TestPkgs, BlockedPkgs, BlockedErrors,
        BuildErrorsByPackage, RuntimeFailuresByPackage.
    .PARAMETER AllTestPkgs
        Array of all test package paths.
    .PARAMETER CovPkgList
        Comma-separated list of source packages for -coverpkg.
    .PARAMETER IsSyncMode
        If $true, run sequentially; otherwise use ForEach-Object -Parallel.
    #>
    [CmdletBinding()]
    param(
        [string[]]$AllTestPkgs,
        [string]$CovPkgList,
        [bool]$IsSyncMode
    )

    $blockedPkgs = [System.Collections.Generic.List[string]]::new()
    $blockedErrors = [System.Collections.Generic.Dictionary[string, string]]::new()
    $testPkgs = [System.Collections.Generic.List[string]]::new()
    $buildErrorsByPackage = @{}
    $runtimeFailuresByPackage = @{}

    $modeLabel = if ($IsSyncMode) { "sync" } else { "parallel" }
    Write-Host ""
    Write-Header "Pre-coverage compile check ($($AllTestPkgs.Count) packages, $modeLabel mode)"

    if ($IsSyncMode) {
        foreach ($testPkg in $AllTestPkgs) {
            $shortName = $testPkg -replace '.*integratedtests/?', ''
            if (-not $shortName) { $shortName = "(root)" }

            $prevPref = $ErrorActionPreference
            $ErrorActionPreference = "Continue"
            $compileOut = & go test -count=1 -run '^$' -gcflags=all=-e "-coverpkg=$CovPkgList" "$testPkg" 2>&1 | ForEach-Object { $_.ToString() }
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
        $throttle = [Math]::Min($AllTestPkgs.Count, [Environment]::ProcessorCount * 2)
        Write-Host "  Launching $($AllTestPkgs.Count) compile checks ($throttle parallel)..." -ForegroundColor Gray

        $compileResults = $AllTestPkgs | ForEach-Object -ThrottleLimit $throttle -Parallel {
            $pkg = $_
            $covPkgs = $using:CovPkgList
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
                    if ($seen.Add($normalized)) { $merged.Add($normalized) | Out-Null }
                }
                $out = $merged.ToArray()
            }

            [pscustomobject]@{ Pkg = $pkg; ExitCode = $ec; Output = $out }
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

    return @{
        TestPkgs                  = $testPkgs
        BlockedPkgs               = $blockedPkgs
        BlockedErrors             = $blockedErrors
        BuildErrorsByPackage      = $buildErrorsByPackage
        RuntimeFailuresByPackage  = $runtimeFailuresByPackage
    }
}

function Invoke-CoverageSplitRecovery {
    <#
    .SYNOPSIS
        Split blocked packages into per-file subfolders and recheck each.
    .DESCRIPTION
        When a package fails to compile, splits each *_test.go into its own
        subfolder, rechecks independently, and promotes passing subfolders.
    .PARAMETER CompileResult
        Hashtable from Invoke-CoverageCompileCheck.
    .PARAMETER AllTestPkgs
        Array of all test package paths.
    .PARAMETER CovPkgList
        Comma-separated coverpkg list.
    .PARAMETER IsSyncMode
        Sequential or parallel mode.
    #>
    [CmdletBinding()]
    param(
        [hashtable]$CompileResult,
        [string[]]$AllTestPkgs,
        [string]$CovPkgList,
        [bool]$IsSyncMode
    )

    $blockedPkgs = $CompileResult.BlockedPkgs
    $blockedErrors = $CompileResult.BlockedErrors
    $testPkgs = $CompileResult.TestPkgs
    $buildErrorsByPackage = $CompileResult.BuildErrorsByPackage

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

            if ($bpTestFiles.Count -lt 2) { continue }

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

            $subDirs = Get-ChildItem -LiteralPath $pkgDir -Directory | Sort-Object Name

            if ($IsSyncMode) {
                foreach ($sd in $subDirs) {
                    $subPkg = "./tests/integratedtests/$bp/$($sd.Name)/"
                    $prevPref = $ErrorActionPreference
                    $ErrorActionPreference = "Continue"
                    $subOut = & go test -count=1 -run '^$' -gcflags=all=-e "-coverpkg=$CovPkgList" "$subPkg" 2>&1 | ForEach-Object { $_.ToString() }
                    $subExit = $LASTEXITCODE
                    $ErrorActionPreference = $prevPref
                    $subfolderResults.Add([pscustomobject]@{ Name = $sd.Name; Pkg = $subPkg; ExitCode = $subExit; Output = $subOut })
                }
            } else {
                $throttle = [Math]::Min($subDirs.Count, [Environment]::ProcessorCount * 2)
                $parallelResults = $subDirs | ForEach-Object -ThrottleLimit $throttle -Parallel {
                    $sd = $_
                    $bpName = $using:bp
                    $covPkgs = $using:CovPkgList
                    $subPkg = "./tests/integratedtests/$bpName/$($sd.Name)/"
                    $ErrorActionPreference = "Continue"
                    $subOut = & go test -count=1 -run '^$' -gcflags=all=-e "-coverpkg=$covPkgs" "$subPkg" 2>&1 | ForEach-Object { $_.ToString() }
                    [pscustomobject]@{ Name = $sd.Name; Pkg = $subPkg; ExitCode = $LASTEXITCODE; Output = $subOut }
                }
                foreach ($pr in ($parallelResults | Sort-Object Name)) { $subfolderResults.Add($pr) }
            }

            $subPass = @($subfolderResults | Where-Object { $_.ExitCode -eq 0 })
            $subFail = @($subfolderResults | Where-Object { $_.ExitCode -ne 0 })

            Write-Host "    ✓ $($subPass.Count) subfolders compile OK" -ForegroundColor Green
            if ($subFail.Count -gt 0) {
                Write-Host "    ✗ $($subFail.Count) subfolders failed:" -ForegroundColor Red
                foreach ($sf in $subFail) {
                    Write-Host "      ✗ $($sf.Name)" -ForegroundColor Red
                    $splitBlockedFiles.Add("$bp/$($sf.Name)")
                }
            }

            foreach ($sp in $subPass) {
                $fullSubPkg = $AllTestPkgs | Where-Object { $_ -match "integratedtests/$bp$" } | Select-Object -First 1
                if ($fullSubPkg) { $subFullPkg = $fullSubPkg + "/" + $sp.Name }
                else { $subFullPkg = $sp.Pkg }
                $testPkgs.Add($subFullPkg)
                $splitRecoveredCount++
            }

            $blockedPkgs.Remove($bp)
            $blockedErrors.Remove($bp)

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

    return @{
        SplitRecoveredCount = $splitRecoveredCount
        SplitBlockedFiles   = $splitBlockedFiles
        SplitCleanupDirs    = $splitCleanupDirs
    }
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Invoke-CoverageCompileCheck',
    'Invoke-CoverageSplitRecovery'
)
