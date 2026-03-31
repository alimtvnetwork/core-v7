# ─────────────────────────────────────────────────────────────────────────────
# Utilities.psm1 — Common helper functions for the project runner toolchain
#
# Provides: console output wrappers, test-log directory management,
#           line filtering/merging, build/runtime error extraction,
#           and compile error parsing.
#
# Usage:
#   Import-Module ./scripts/Utilities.psm1 -Force
#
# Dependencies: DashboardUI.psm1 (optional — gracefully degrades)
# ─────────────────────────────────────────────────────────────────────────────

# ═══════════════════════════════════════════════════════════════════════════════
# §1  ANSI Escape Sequences (fallback if DashboardUI not loaded)
# ═══════════════════════════════════════════════════════════════════════════════

$script:ESC    = [char]27
$script:cReset = "$script:ESC[0m"

# If DashboardUI hasn't set these, provide safe fallbacks
if (-not $script:cLime)  { $script:cLime  = "$script:ESC[38;2;163;230;53m" }
if (-not $script:cRed)   { $script:cRed   = "$script:ESC[38;2;244;63;94m" }
if (-not $script:cCyan)  { $script:cCyan  = "$script:ESC[38;2;6;182;212m" }
if (-not $script:cWhite) { $script:cWhite = "$script:ESC[38;2;255;255;255m" }

# ═══════════════════════════════════════════════════════════════════════════════
# §2  Console Output Wrappers
# ═══════════════════════════════════════════════════════════════════════════════

function Write-Header {
    <#
    .SYNOPSIS
        Print a phase-start header line. Uses DashboardUI if available,
        falls back to a simple "=== msg ===" banner.
    .PARAMETER msg
        The header text to display.
    .EXAMPLE
        Write-Header "Running tests"
    #>
    [CmdletBinding()]
    param([string]$msg)

    if (Get-Command Write-DashboardHeader -ErrorAction SilentlyContinue) {
        Write-Host ""
        Write-PhaseStart -Name $msg
    } else {
        Write-Host "`n=== $msg ===" -ForegroundColor Cyan
    }
}

function Write-Success {
    <#
    .SYNOPSIS
        Print a green success line with a ✓ icon.
    .PARAMETER msg
        The success message.
    .EXAMPLE
        Write-Success "All tests passed"
    #>
    [CmdletBinding()]
    param([string]$msg)

    Write-Host "  $($script:cLime)✓$($script:cReset) $($script:cLime)$msg$($script:cReset)"
}

function Write-Fail {
    <#
    .SYNOPSIS
        Print a red failure line with a ✗ icon.
    .PARAMETER msg
        The failure message.
    .EXAMPLE
        Write-Fail "3 tests failed"
    #>
    [CmdletBinding()]
    param([string]$msg)

    Write-Host "  $($script:cRed)✗$($script:cReset) $($script:cRed)$msg$($script:cReset)"
}

# ═══════════════════════════════════════════════════════════════════════════════
# §3  Test Log Directory
# ═══════════════════════════════════════════════════════════════════════════════

function Ensure-TestLogDir {
    <#
    .SYNOPSIS
        Create the test-log output directory if it doesn't exist.
    .DESCRIPTION
        Uses the $TestLogDir variable from the caller's scope (set in run.ps1).
        Default path: data/test-logs/
    .EXAMPLE
        Ensure-TestLogDir
    #>
    [CmdletBinding()]
    param()

    if (-not (Test-Path $TestLogDir)) {
        New-Item -ItemType Directory -Path $TestLogDir -Force | Out-Null
    }
}

# ═══════════════════════════════════════════════════════════════════════════════
# §4  Line Filtering & Merging
# ═══════════════════════════════════════════════════════════════════════════════

function Filter-TestWarnings {
    <#
    .SYNOPSIS
        Remove "no packages being tested" warnings from Go test output.
    .PARAMETER lines
        Array of raw output lines.
    .EXAMPLE
        $clean = Filter-TestWarnings $rawOutput
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)

    return $lines | Where-Object {
        $_ -notmatch '^\s*warning: no packages being tested depend on matches for pattern'
    }
}

function Merge-UniqueOutputLines {
    <#
    .SYNOPSIS
        Merge two string arrays, deduplicating by exact content.
    .PARAMETER primary
        First array (higher priority).
    .PARAMETER secondary
        Second array to merge in.
    .EXAMPLE
        $merged = Merge-UniqueOutputLines $buildOutput $testOutput
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param(
        [string[]]$primary,
        [string[]]$secondary
    )

    $merged = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)

    foreach ($line in @($primary + $secondary)) {
        if ($null -eq $line) { continue }
        $normalized = $line.ToString().TrimEnd("`r")
        if (-not $normalized) { continue }
        if ($seen.Add($normalized)) {
            $merged.Add($normalized) | Out-Null
        }
    }

    return $merged.ToArray()
}

function Filter-BlockedCompileLines {
    <#
    .SYNOPSIS
        Remove noisy/irrelevant lines from Go compile output, keeping only
        actionable error lines.
    .DESCRIPTION
        Strips: "warning: no packages being tested..." lines, bare package
        headers (# github.com/...) without file:line references, and bare
        package path lines.
    .PARAMETER lines
        Array of raw compiler output lines.
    .EXAMPLE
        $actionable = Filter-BlockedCompileLines $compileOutput
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)

    $filtered = [System.Collections.Generic.List[string]]::new()

    foreach ($raw in $lines) {
        if ($null -eq $raw) { continue }

        $line = $raw.ToString().TrimEnd("`r")
        $trimmed = $line.Trim()

        if (-not $trimmed) { continue }

        # Strip "warning: no packages being tested..." noise
        if ($trimmed -match '^\s*warning:\s*no packages being tested depend on matches for pattern') { continue }

        # Strip bare package headers like "# github.com/org/repo [...]" without file:line
        if ($trimmed -match '^#\s+\S+' -and $trimmed -notmatch '\.go:\d+') { continue }

        # Strip bare package path lines (github/gitlab) without file:line
        if ($trimmed -match '^(github\.com|gitlab\.com)/\S+(\s+\[[^\]]+\])?$' -and $trimmed -notmatch '\.go:\d+') { continue }

        $filtered.Add($line) | Out-Null
    }

    return $filtered.ToArray()
}

# ═══════════════════════════════════════════════════════════════════════════════
# §5  Error Extraction
# ═══════════════════════════════════════════════════════════════════════════════

function Extract-BuildErrorLines {
    <#
    .SYNOPSIS
        Extract compile-time error lines from Go build output.
    .DESCRIPTION
        Filters through blocked-compile lines and returns only those matching
        Go error patterns: file.go:line:, package headers, [build failed].
    .PARAMETER lines
        Array of raw build output lines.
    .EXAMPLE
        $errors = Extract-BuildErrorLines $buildOutput
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)

    $candidates = Filter-BlockedCompileLines $lines
    $errors = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)

    foreach ($raw in $candidates) {
        if ($null -eq $raw) { continue }

        $line = $raw.ToString().TrimEnd("`r")
        $trimmed = $line.Trim()
        if (-not $trimmed) { continue }

        if ($trimmed -match '\.go:\d+(?::\d+)?:' -or
            $trimmed -match '^#\s+\S+' -or
            $trimmed -match '\[build failed\]' -or
            $trimmed -match '(?i)\bbuild failed\b') {
            if ($seen.Add($line)) {
                $errors.Add($line) | Out-Null
            }
        }
    }

    return $errors.ToArray()
}

function Extract-ExecutionFailureLines {
    <#
    .SYNOPSIS
        Extract execution failure lines (compile + runtime) from Go output.
    .DESCRIPTION
        Returns lines matching compile errors, panics, fatal errors,
        test failures (--- FAIL), FAIL lines, and exit status codes.
    .PARAMETER lines
        Array of raw output lines.
    .EXAMPLE
        $failures = Extract-ExecutionFailureLines $testOutput
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)

    $candidates = Filter-BlockedCompileLines $lines
    $errors = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)

    foreach ($raw in $candidates) {
        if ($null -eq $raw) { continue }

        $line = $raw.ToString().TrimEnd("`r")
        $trimmed = $line.Trim()
        if (-not $trimmed) { continue }

        if ($trimmed -match '\.go:\d+(?::\d+)?:' -or
            $trimmed -match '^#\s+\S+' -or
            $trimmed -match '\[build failed\]' -or
            $trimmed -match '(?i)\bbuild failed\b' -or
            $trimmed -match '^(?i)panic:' -or
            $trimmed -match '^(?i)fatal error:' -or
            $trimmed -match '^--- FAIL:\s+' -or
            $trimmed -match '^\s*FAIL\s+\S+' -or
            $trimmed -match '^\s*exit status \d+\s*$') {
            if ($seen.Add($line)) {
                $errors.Add($line) | Out-Null
            }
        }
    }

    return $errors.ToArray()
}

function Extract-RuntimeFailureLines {
    <#
    .SYNOPSIS
        Extract ONLY runtime failure lines from Go output (no compile errors).
    .DESCRIPTION
        Captures: panics, fatal errors, goroutine dumps, test crashes,
        FAIL lines, exit status, signals, and runtime errors.
        Does NOT include compile errors (.go:line: syntax) or [build failed].
    .PARAMETER lines
        Array of raw output lines.
    .EXAMPLE
        $runtime = Extract-RuntimeFailureLines $testOutput
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)

    $candidates = Filter-BlockedCompileLines $lines
    $errors = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)

    foreach ($raw in $candidates) {
        if ($null -eq $raw) { continue }

        $line = $raw.ToString().TrimEnd("`r")
        $trimmed = $line.Trim()
        if (-not $trimmed) { continue }

        if ($trimmed -match '^(?i)panic:' -or
            $trimmed -match '^(?i)fatal error:' -or
            $trimmed -match '^(?i)goroutine \d+' -or
            $trimmed -match '^--- FAIL:\s+' -or
            $trimmed -match '^\s*FAIL\s+\S+' -or
            $trimmed -match '^\s*exit status \d+\s*$' -or
            $trimmed -match '(?i)signal:\s+' -or
            $trimmed -match '(?i)runtime error:') {
            if ($seen.Add($line)) {
                $errors.Add($line) | Out-Null
            }
        }
    }

    return $errors.ToArray()
}

# ═══════════════════════════════════════════════════════════════════════════════
# §6  Package Error Aggregation
# ═══════════════════════════════════════════════════════════════════════════════

function Add-BuildErrorsForPackage {
    <#
    .SYNOPSIS
        Accumulate build errors into a per-package hashtable.
    .PARAMETER BuildErrorMap
        Hashtable mapping package names to lists of error lines.
    .PARAMETER PackageName
        The Go package name (e.g., "regexnewtests").
    .PARAMETER Lines
        Raw build output lines to extract errors from.
    .EXAMPLE
        Add-BuildErrorsForPackage $errorMap "mypackage" $output
    #>
    [CmdletBinding()]
    param(
        [hashtable]$BuildErrorMap,
        [string]$PackageName,
        [string[]]$Lines
    )

    if (-not $BuildErrorMap -or -not $PackageName) { return }

    $buildLines = Extract-BuildErrorLines $Lines
    if (-not $buildLines -or $buildLines.Count -eq 0) { return }

    if (-not $BuildErrorMap.ContainsKey($PackageName)) {
        $BuildErrorMap[$PackageName] = [System.Collections.Generic.List[string]]::new()
    }

    foreach ($line in $buildLines) {
        if (-not $BuildErrorMap[$PackageName].Contains($line)) {
            $BuildErrorMap[$PackageName].Add($line) | Out-Null
        }
    }
}

function Add-RuntimeFailuresForPackage {
    <#
    .SYNOPSIS
        Accumulate runtime failures into a per-package hashtable.
    .PARAMETER FailureMap
        Hashtable mapping package names to lists of failure lines.
    .PARAMETER PackageName
        The Go package name.
    .PARAMETER Lines
        Raw test output lines to extract runtime failures from.
    .EXAMPLE
        Add-RuntimeFailuresForPackage $failMap "mypackage" $output
    #>
    [CmdletBinding()]
    param(
        [hashtable]$FailureMap,
        [string]$PackageName,
        [string[]]$Lines
    )

    if (-not $FailureMap -or -not $PackageName) { return }

    $runtimeLines = Extract-RuntimeFailureLines $Lines
    if (-not $runtimeLines -or $runtimeLines.Count -eq 0) { return }

    if (-not $FailureMap.ContainsKey($PackageName)) {
        $FailureMap[$PackageName] = [System.Collections.Generic.List[string]]::new()
    }

    foreach ($line in $runtimeLines) {
        if (-not $FailureMap[$PackageName].Contains($line)) {
            $FailureMap[$PackageName].Add($line) | Out-Null
        }
    }
}

# ═══════════════════════════════════════════════════════════════════════════════
# §7  Compile Error Parsing
# ═══════════════════════════════════════════════════════════════════════════════

function ParseCompileErrors {
    <#
    .SYNOPSIS
        Parse Go compile error lines into structured objects.
    .DESCRIPTION
        Extracts file name, line number, message, and error category
        from standard Go compiler output (file.go:line: message).
        Categories: arg-count, undefined, type-mismatch, missing-member,
        field-vs-method, other.
    .PARAMETER output
        Array of raw compiler output lines.
    .OUTPUTS
        Array of hashtables with keys: file, line, message, category, raw.
    .EXAMPLE
        $parsed = ParseCompileErrors $buildOutput
        $parsed | Where-Object { $_.category -eq "undefined" }
    #>
    [CmdletBinding()]
    [OutputType([hashtable[]])]
    param([string[]]$output)

    $errors = [System.Collections.Generic.List[object]]::new()
    foreach ($line in $output) {
        if ($line -match '^(.+?\.go):(\d+)(?::\d+)?:\s*(.+)$') {
            $file = Split-Path $Matches[1] -Leaf
            $lineNum = [int]$Matches[2]
            $msg = $Matches[3].Trim()

            # Classify error
            $category = "other"
            if ($msg -match 'too many arguments|not enough arguments') { $category = "arg-count" }
            elseif ($msg -match 'undefined:') { $category = "undefined" }
            elseif ($msg -match 'cannot use .* as') { $category = "type-mismatch" }
            elseif ($msg -match 'has no field or method') { $category = "missing-member" }
            elseif ($msg -match 'cannot call non-function') { $category = "field-vs-method" }

            $errors.Add(@{
                file     = $file
                line     = $lineNum
                message  = $msg
                category = $category
                raw      = $line
            })
        }
    }
    return $errors.ToArray()
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Write-Header',
    'Write-Success',
    'Write-Fail',
    'Ensure-TestLogDir',
    'Filter-TestWarnings',
    'Merge-UniqueOutputLines',
    'Filter-BlockedCompileLines',
    'Extract-BuildErrorLines',
    'Extract-ExecutionFailureLines',
    'Extract-RuntimeFailureLines',
    'Add-BuildErrorsForPackage',
    'Add-RuntimeFailuresForPackage',
    'ParseCompileErrors'
)
