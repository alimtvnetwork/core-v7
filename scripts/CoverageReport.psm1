# ─────────────────────────────────────────────────────────────────────────────
# CoverageReport.psm1 — Coverage report generation (TXT, JSON, HTML, AI)
#
# Usage:
#   Import-Module ./scripts/CoverageReport.psm1 -Force
#
# Dependencies: Utilities.psm1 (Write-Header, Write-Success, Extract-BuildErrorLines,
#               Extract-ExecutionFailureLines)
#               DashboardUI.psm1 (Write-CoverageComparison, Load/Save-CoverageSnapshot)
# ─────────────────────────────────────────────────────────────────────────────

function Build-SourcePackageCoverage {
    <#
    .SYNOPSIS
        Parse merged coverage lines into per-source-package coverage stats.
    .RETURNS
        Ordered hashtable: shortPkgName → @{ Stmts; Covered }
    #>
    [CmdletBinding()]
    param([string[]]$MergedLines)

    $srcPkgStmts = [ordered]@{}
    foreach ($covLine in $MergedLines) {
        if ($covLine -match "^mode:") { continue }
        if ($covLine -match "^(\S+?):(\d+)\.(\d+),(\d+)\.(\d+)\s+(\d+)\s+(\d+)") {
            $filePath = $Matches[1]
            $stmts = [int]$Matches[6]
            $count = [int]$Matches[7]
            $shortSrc = $filePath -replace '.*alimtvnetwork/core/?', ''
            $shortSrc = $shortSrc -replace '/[^/]+$', ''
            if (-not $shortSrc) { $shortSrc = "(root)" }
            if (-not $srcPkgStmts.Contains($shortSrc)) {
                $srcPkgStmts[$shortSrc] = @{ Stmts = 0; Covered = 0 }
            }
            $srcPkgStmts[$shortSrc].Stmts += $stmts
            if ($count -gt 0) { $srcPkgStmts[$shortSrc].Covered += $stmts }
        }
    }
    return $srcPkgStmts
}

function Write-CoverageSummaryReport {
    <#
    .SYNOPSIS
        Generate coverage-summary.txt file.
    #>
    [CmdletBinding()]
    param(
        [string]$CoverSummary,
        [string]$CoverProfile,
        [string]$CoverHtml,
        [string[]]$FuncOutput,
        [hashtable]$SrcPkgStmts
    )

    $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    $summaryLines = [System.Collections.Generic.List[string]]::new()
    $summaryLines.Add("# Coverage Summary — $timestamp")
    $summaryLines.Add("")

    $totalLine = $FuncOutput | Where-Object { $_ -match "^total:" } | Select-Object -Last 1
    if ($totalLine) {
        $summaryLines.Add("## Total Coverage")
        $summaryLines.Add("  $totalLine")
        $summaryLines.Add("")
    }

    if ($SrcPkgStmts.Count -gt 0) {
        $summaryLines.Add("## Per-Package Coverage (Source)")
        $sorted = $SrcPkgStmts.GetEnumerator() | ForEach-Object {
            $pct = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
            [pscustomobject]@{ Name = $_.Key; Pct = $pct }
        } | Sort-Object Pct -Descending
        foreach ($entry in $sorted) { $summaryLines.Add("  $($entry.Pct)%`t$($entry.Name)") }
        $summaryLines.Add("")
    }

    $lowCovFuncs = Get-LowCoverageFunctions $FuncOutput
    if ($lowCovFuncs.Count -gt 0) {
        $summaryLines.Add("## Low Coverage Functions (< 50%)")
        $summaryLines.Add("  Count: $($lowCovFuncs.Count)")
        $summaryLines.Add("")
        foreach ($f in $lowCovFuncs) { $summaryLines.Add($f) }
        $summaryLines.Add("")
    }

    $summaryLines.Add("## Reports")
    $summaryLines.Add("  Profile:  $CoverProfile")
    $summaryLines.Add("  HTML:     $CoverHtml")
    $summaryLines.Add("  Summary:  $CoverSummary")

    Set-Content -Path $CoverSummary -Value ($summaryLines -join "`n") -Encoding UTF8
}

function Get-LowCoverageFunctions {
    <#
    .SYNOPSIS
        Extract functions with < 50% coverage from func output.
    #>
    [CmdletBinding()]
    param([string[]]$FuncOutput)

    $lowCovFuncs = [System.Collections.Generic.List[string]]::new()
    foreach ($line in $FuncOutput) {
        if ($line -match "(\d+\.\d+)%\s*$" -and $line -notmatch "^total:") {
            $pct = [double]$Matches[1]
            if ($pct -lt 50.0) { $lowCovFuncs.Add("  $line") }
        }
    }
    return $lowCovFuncs
}

function Write-CoverageJsonReport {
    <#
    .SYNOPSIS
        Generate coverage-summary.json with per-package and low-coverage data.
    #>
    [CmdletBinding()]
    param(
        [string]$CoverJsonFile,
        [string]$CoverProfile,
        [string]$CoverHtml,
        [string]$CoverSummary,
        [string[]]$FuncOutput,
        [hashtable]$SrcPkgStmts,
        [string]$CoverDir
    )

    $totalLine = $FuncOutput | Where-Object { $_ -match "^total:" } | Select-Object -Last 1
    $totalPct = 0.0
    if ($totalLine -match "(\d+\.\d+)%") { $totalPct = [double]$Matches[1] }

    $pkgJsonItems = [System.Collections.Generic.List[object]]::new()
    if ($SrcPkgStmts.Count -gt 0) {
        $sorted = $SrcPkgStmts.GetEnumerator() | ForEach-Object {
            $pct = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
            [pscustomobject]@{ Name = $_.Key; Pct = $pct; Stmts = $_.Value.Stmts; Covered = $_.Value.Covered }
        } | Sort-Object Pct
        foreach ($e in $sorted) {
            $pkgJsonItems.Add(@{
                package = $e.Name; coverage = $e.Pct; statements = $e.Stmts
                covered = $e.Covered; uncovered = $e.Stmts - $e.Covered
            })
        }
    }

    $lowCovJsonItems = [System.Collections.Generic.List[object]]::new()
    foreach ($line in $FuncOutput) {
        if ($line -match "(\d+\.\d+)%\s*$" -and $line -notmatch "^total:") {
            $pctF = [double]$Matches[1]
            if ($pctF -lt 50.0) {
                $funcName = ""; $funcFile = ""
                if ($line -match "^(\S+):\s+(\S+)\s+(\d+\.\d+)%") { $funcFile = $Matches[1]; $funcName = $Matches[2] }
                $lowCovJsonItems.Add(@{ file = $funcFile; function = $funcName; coverage = $pctF })
            }
        }
    }

    $blockedRef = @()
    $blockedJsonPath = Join-Path $CoverDir "blocked-packages.json"
    if (Test-Path $blockedJsonPath) {
        $blockedRef = @((Get-Content $blockedJsonPath -Raw | ConvertFrom-Json).blockedPackages | ForEach-Object { $_.package })
    }

    $coverJsonObj = @{
        timestamp            = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ")
        totalCoverage        = $totalPct
        packageCount         = $pkgJsonItems.Count
        packages             = $pkgJsonItems.ToArray()
        lowCoverageFuncCount = $lowCovJsonItems.Count
        lowCoverageFunctions = $lowCovJsonItems.ToArray()
        blockedPackages      = $blockedRef
        reports              = @{ profile = $CoverProfile; html = $CoverHtml; summary = $CoverSummary; json = $CoverJsonFile }
    }
    $coverJsonObj | ConvertTo-Json -Depth 4 | Set-Content -Path $CoverJsonFile -Encoding UTF8
    return $totalPct
}

function Write-PerPackageCoverageReport {
    <#
    .SYNOPSIS
        Generate per-package-coverage.txt and per-package-coverage.json.
    #>
    [CmdletBinding()]
    param(
        [string]$CoverDir,
        [hashtable]$SrcPkgStmts,
        [double]$TotalPct,
        [string]$TotalLine
    )

    $perPkgTxtFile = Join-Path $CoverDir "per-package-coverage.txt"
    $perPkgJsonFile = Join-Path $CoverDir "per-package-coverage.json"

    $txtLines = [System.Collections.Generic.List[string]]::new()
    $txtLines.Add("# Per-Package Coverage Report — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')")
    $txtLines.Add("# Total: $TotalLine")
    $txtLines.Add("")
    $txtLines.Add(("Package".PadRight(50)) + " " + ("Stmts".PadLeft(8)) + " " + ("Covered".PadLeft(8)) + " " + ("Uncovered".PadLeft(10)) + " " + ("Cov%".PadLeft(8)))
    $txtLines.Add(("─" * 50) + " " + ("─" * 8) + " " + ("─" * 8) + " " + ("─" * 10) + " " + ("─" * 8))

    $jsonItems = [System.Collections.Generic.List[object]]::new()

    if ($SrcPkgStmts.Count -gt 0) {
        $sorted = $SrcPkgStmts.GetEnumerator() | ForEach-Object {
            $pct = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
            [pscustomobject]@{ Name = $_.Key; Pct = $pct; Stmts = $_.Value.Stmts; Covered = $_.Value.Covered; Uncovered = $_.Value.Stmts - $_.Value.Covered }
        } | Sort-Object Pct

        foreach ($pp in $sorted) {
            $statusMark = if ($pp.Pct -ge 100) { "✓" } elseif ($pp.Pct -ge 80) { "○" } else { "✗" }
            $row = ("$statusMark $($pp.Name)").PadRight(50)
            $row += " " + $pp.Stmts.ToString().PadLeft(8)
            $row += " " + $pp.Covered.ToString().PadLeft(8)
            $row += " " + $pp.Uncovered.ToString().PadLeft(10)
            $row += " " + (([string]::Format([System.Globalization.CultureInfo]::InvariantCulture, "{0:0.0}", $pp.Pct)) + "%").PadLeft(8)
            $txtLines.Add($row)

            $jsonItems.Add(@{
                package = $pp.Name; coverage = $pp.Pct; statements = $pp.Stmts
                covered = $pp.Covered; uncovered = $pp.Uncovered
                status = if ($pp.Pct -ge 100) { "full" } elseif ($pp.Pct -ge 80) { "good" } else { "low" }
            })
        }

        $totalStmts = ($sorted | Measure-Object -Property Stmts -Sum).Sum
        $totalCovered = ($sorted | Measure-Object -Property Covered -Sum).Sum
        $fullCount = ($sorted | Where-Object { $_.Pct -ge 100 }).Count
        $lowCount = ($sorted | Where-Object { $_.Pct -lt 80 }).Count

        $txtLines.Add("")
        $txtLines.Add("# Summary")
        $txtLines.Add("#   Packages:  $($sorted.Count)")
        $txtLines.Add("#   100%:      $fullCount")
        $txtLines.Add("#   < 80%:     $lowCount")
        $txtLines.Add("#   Total stmts: $totalStmts  covered: $totalCovered  uncovered: $($totalStmts - $totalCovered)")
    }

    Set-Content -Path $perPkgTxtFile -Value ($txtLines -join "`n") -Encoding UTF8
    @{ timestamp = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"); totalCoverage = $TotalPct; packageCount = $jsonItems.Count; packages = $jsonItems.ToArray() } |
        ConvertTo-Json -Depth 4 | Set-Content -Path $perPkgJsonFile -Encoding UTF8
}

function Write-CoverageHtmlWithAiButton {
    <#
    .SYNOPSIS
        Inject a "Copy for AI" button into the Go HTML coverage report.
    #>
    [CmdletBinding()]
    param(
        [string]$CoverHtml,
        [string]$CoverProfile,
        [string[]]$FuncOutput,
        [hashtable]$SrcPkgStmts
    )

    $htmlArgs = @("-html=$CoverProfile", "-o=$CoverHtml")
    $htmlErr = & go tool cover $htmlArgs 2>&1
    $htmlExitCode = $LASTEXITCODE

    if ($htmlExitCode -ne 0 -or -not (Test-Path $CoverHtml)) {
        Write-Host "  ⚠ Failed to generate HTML report via 'go tool cover -html' (exit: $htmlExitCode)" -ForegroundColor Red
        if ($htmlErr) { Write-Host "  Error: $htmlErr" -ForegroundColor Red }
        $fallbackHtml = @"
<!DOCTYPE html><html><head><meta charset="utf-8"><title>Coverage Report</title>
<style>body{font-family:monospace;padding:20px;background:#1e1e2e;color:#cdd6f4}
pre{white-space:pre-wrap}</style></head><body>
<h1>Coverage Report</h1><pre>$($FuncOutput -join "`n")</pre></body></html>
"@
        Set-Content -Path $CoverHtml -Value $fallbackHtml -Encoding UTF8
        Write-Host "  Generated fallback HTML report" -ForegroundColor Yellow
    }

    # Build AI-friendly text
    $aiTextLines = [System.Collections.Generic.List[string]]::new()
    $aiTextLines.Add("# Coverage Report — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')")
    $aiTextLines.Add("")
    $aiTextLines.Add("## Goal: Improve test coverage for the packages listed below.")
    $aiTextLines.Add("Please write tests for uncovered functions, following the project's AAA pattern.")
    $aiTextLines.Add("")

    $totalLine = $FuncOutput | Where-Object { $_ -match "^total:" } | Select-Object -Last 1
    if ($totalLine) { $aiTextLines.Add("## Total Coverage"); $aiTextLines.Add($totalLine); $aiTextLines.Add("") }

    if ($SrcPkgStmts.Count -gt 0) {
        $aiTextLines.Add("## Per-Source-Package Coverage")
        $computed = $SrcPkgStmts.GetEnumerator() | ForEach-Object {
            $pct = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
            [pscustomobject]@{ Name = $_.Key; Pct = $pct; Stmts = $_.Value.Stmts; Covered = $_.Value.Covered }
        } | Sort-Object Pct
        foreach ($e in $computed) { $aiTextLines.Add("  $($e.Pct)%  $($e.Name)  ($($e.Covered)/$($e.Stmts) stmts)") }
        $aiTextLines.Add("")
    }

    $lowCovFuncs = Get-LowCoverageFunctions $FuncOutput
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

    if (Test-Path $CoverHtml) {
        $htmlContent = Get-Content -Path $CoverHtml -Raw

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
        Set-Content -Path $CoverHtml -Value $htmlContent -Encoding UTF8
        Write-Host "  ✓ Injected 'Copy for AI' button into HTML report" -ForegroundColor Green
    }
}

function Write-BuildErrorsReport {
    <#
    .SYNOPSIS
        Generate build-errors.txt and build-errors.json reports.
    #>
    [CmdletBinding()]
    param(
        [string]$CoverDir,
        [hashtable]$BuildErrorsByPackage,
        [System.Collections.Generic.List[string]]$BlockedPkgs,
        [System.Collections.Generic.Dictionary[string, string]]$BlockedErrors
    )

    # Merge blocked-package compile errors
    if ($BlockedPkgs.Count -gt 0) {
        foreach ($bp in ($BlockedPkgs | Sort-Object)) {
            if ($BlockedErrors.ContainsKey($bp)) {
                $rawErrLines = $BlockedErrors[$bp] -split "`n"
                $filteredErrLines = Extract-BuildErrorLines $rawErrLines
                if ($filteredErrLines.Count -eq 0) { $filteredErrLines = Extract-ExecutionFailureLines $rawErrLines }
                if ($filteredErrLines.Count -gt 0) {
                    if (-not $BuildErrorsByPackage.ContainsKey($bp)) {
                        $BuildErrorsByPackage[$bp] = [System.Collections.Generic.List[string]]::new()
                    }
                    foreach ($errLine in $filteredErrLines) {
                        if (-not $BuildErrorsByPackage[$bp].Contains($errLine)) {
                            $BuildErrorsByPackage[$bp].Add($errLine) | Out-Null
                        }
                    }
                }
            }
        }
    }

    $buildErrorsFile = Join-Path $CoverDir "build-errors.txt"
    $buildErrorsJsonFile = Join-Path $CoverDir "build-errors.json"
    $buildErrorPkgs = @($BuildErrorsByPackage.Keys | Sort-Object)

    $lines = [System.Collections.Generic.List[string]]::new()
    $lines.Add("# Build Errors — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')")
    $lines.Add("# Count: $($buildErrorPkgs.Count)")
    $lines.Add("")

    $jsonItems = [System.Collections.Generic.List[object]]::new()

    if ($buildErrorPkgs.Count -eq 0) { $lines.Add("No build errors captured.") }
    else {
        foreach ($pkgName in $buildErrorPkgs) {
            $pkgLines = @($BuildErrorsByPackage[$pkgName])
            $isBlocked = $BlockedPkgs.Contains($pkgName)
            $label = if ($isBlocked) { "## $pkgName [BLOCKED — compile failure]" } else { "## $pkgName [coverage-run error]" }
            $lines.Add($label)
            if ($pkgLines.Count -gt 0) { $lines.AddRange([string[]]$pkgLines) } else { $lines.Add("(no actionable compile errors captured)") }
            $lines.Add("")
            $jsonItems.Add(@{ package = $pkgName; errorCount = $pkgLines.Count; errors = $pkgLines; source = if ($isBlocked) { "compile-check" } else { "coverage-run" } }) | Out-Null
        }
    }

    Set-Content -Path $buildErrorsFile -Value ($lines -join "`n") -Encoding UTF8
    @{ timestamp = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"); packageCount = $buildErrorPkgs.Count; blockedCount = $BlockedPkgs.Count; packages = $jsonItems.ToArray() } |
        ConvertTo-Json -Depth 5 | Set-Content -Path $buildErrorsJsonFile -Encoding UTF8
}

function Write-RuntimeFailuresReport {
    <#
    .SYNOPSIS
        Generate runtime-failures.txt and runtime-failures.json reports.
    #>
    [CmdletBinding()]
    param(
        [string]$CoverDir,
        [hashtable]$RuntimeFailuresByPackage,
        [System.Collections.Generic.List[string]]$MissingProfiles
    )

    if ($MissingProfiles.Count -gt 0) {
        foreach ($mp in $MissingProfiles) {
            if (-not $RuntimeFailuresByPackage.ContainsKey($mp)) {
                $RuntimeFailuresByPackage[$mp] = [System.Collections.Generic.List[string]]::new()
            }
            $crashMsg = "coverage profile missing — test binary likely crashed (panic/os.Exit)"
            if (-not $RuntimeFailuresByPackage[$mp].Contains($crashMsg)) {
                $RuntimeFailuresByPackage[$mp].Add($crashMsg) | Out-Null
            }
        }
    }

    $runtimeFailuresFile = Join-Path $CoverDir "runtime-failures.txt"
    $runtimeFailuresJsonFile = Join-Path $CoverDir "runtime-failures.json"
    $runtimeFailurePkgs = @($RuntimeFailuresByPackage.Keys | Sort-Object)

    $rtLines = [System.Collections.Generic.List[string]]::new()
    $rtLines.Add("# Runtime Failures — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')")
    $rtLines.Add("# Panics, os.Exit, test binary crashes, fatal errors")
    $rtLines.Add("# Count: $($runtimeFailurePkgs.Count) package(s)")
    $rtLines.Add("")

    $rtJsonItems = [System.Collections.Generic.List[object]]::new()

    if ($runtimeFailurePkgs.Count -eq 0) { $rtLines.Add("No runtime failures captured.") }
    else {
        foreach ($pkgName in $runtimeFailurePkgs) {
            $pkgLines = @($RuntimeFailuresByPackage[$pkgName])
            $rtLines.Add("## $pkgName")
            if ($pkgLines.Count -gt 0) { $rtLines.AddRange([string[]]$pkgLines) } else { $rtLines.Add("(no failure details captured)") }
            $rtLines.Add("")
            $rtJsonItems.Add(@{ package = $pkgName; failureCount = $pkgLines.Count; failures = $pkgLines }) | Out-Null
        }
    }

    Set-Content -Path $runtimeFailuresFile -Value ($rtLines -join "`n") -Encoding UTF8
    @{ timestamp = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"); packageCount = $runtimeFailurePkgs.Count; packages = $rtJsonItems.ToArray() } |
        ConvertTo-Json -Depth 5 | Set-Content -Path $runtimeFailuresJsonFile -Encoding UTF8

    # Console warning
    if ($runtimeFailurePkgs.Count -gt 0) {
        Write-Host ""
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Magenta
        Write-Host "  │ RUNTIME FAILURES ($($runtimeFailurePkgs.Count) package(s))" -ForegroundColor Magenta
        Write-Host "  │" -ForegroundColor Magenta
        foreach ($rp in $runtimeFailurePkgs) { Write-Host "  │   ⚠ $rp" -ForegroundColor Yellow }
        Write-Host "  │" -ForegroundColor Magenta
        Write-Host "  │ See data/coverage/runtime-failures.txt for details." -ForegroundColor Yellow
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Magenta
    }
}

function Write-CoverageConsoleSummary {
    <#
    .SYNOPSIS
        Print coverage summary and comparison to console.
    #>
    [CmdletBinding()]
    param(
        [hashtable]$SrcPkgStmts,
        [string[]]$FuncOutput
    )

    $totalLine = $FuncOutput | Where-Object { $_ -match "^total:" } | Select-Object -Last 1
    $lowCovFuncs = Get-LowCoverageFunctions $FuncOutput

    if ($SrcPkgStmts.Count -gt 0) {
        Write-Host ""
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Cyan
        Write-Host "  │ COVERAGE SUMMARY" -ForegroundColor Cyan
        Write-Host "  │" -ForegroundColor Cyan
        $sorted = $SrcPkgStmts.GetEnumerator() | ForEach-Object {
            $pct = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
            [pscustomobject]@{ Name = $_.Key; Pct = $pct }
        } | Sort-Object Pct -Descending
        foreach ($entry in $sorted) {
            $color = if ($entry.Pct -ge 100) { "Green" } elseif ($entry.Pct -ge 80) { "Yellow" } else { "Red" }
            Write-Host "  │  $($entry.Pct)%`t$($entry.Name)" -ForegroundColor $color
        }
        Write-Host "  │" -ForegroundColor Cyan
        if ($totalLine -and $totalLine -match "(\d+\.\d+)%") {
            Write-Host "  │  total: (statements)  $($Matches[1])%" -ForegroundColor Cyan
        }
        if ($lowCovFuncs.Count -gt 0) {
            Write-Host "  │  ⚠ $($lowCovFuncs.Count) function(s) below 50% coverage" -ForegroundColor Yellow
        }
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Cyan
    }

    # Coverage comparison
    if (Get-Command Write-CoverageComparison -ErrorAction SilentlyContinue) {
        $currentCovData = @()
        if ($SrcPkgStmts.Count -gt 0) {
            $currentCovData = @($SrcPkgStmts.GetEnumerator() | ForEach-Object {
                $pct = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
                @{ Package = $_.Key; Coverage = $pct }
            })
        }
        if ($currentCovData.Count -gt 0) {
            $previousCovData = $null
            if (Get-Command Load-CoverageSnapshot -ErrorAction SilentlyContinue) { $previousCovData = Load-CoverageSnapshot }
            Write-Host ""
            Write-CoverageComparison -Current $currentCovData -Previous $previousCovData
            if (Get-Command Save-CoverageSnapshot -ErrorAction SilentlyContinue) { Save-CoverageSnapshot -CoverageData $currentCovData }
        }
    }
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Build-SourcePackageCoverage',
    'Write-CoverageSummaryReport',
    'Get-LowCoverageFunctions',
    'Write-CoverageJsonReport',
    'Write-PerPackageCoverageReport',
    'Write-CoverageHtmlWithAiButton',
    'Write-BuildErrorsReport',
    'Write-RuntimeFailuresReport',
    'Write-CoverageConsoleSummary'
)
