# ─────────────────────────────────────────────────────────────────────────────
# ErrorExtractor.psm1 — Go build/runtime error line extraction
#
# Dependencies: None (standalone)
# ─────────────────────────────────────────────────────────────────────────────

function Filter-BlockedCompileLines {
    <# .SYNOPSIS Remove noisy/irrelevant lines from Go compile output. #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)
    $filtered = [System.Collections.Generic.List[string]]::new()
    foreach ($raw in $lines) {
        if ($null -eq $raw) { continue }; $line = $raw.ToString().TrimEnd("`r"); $trimmed = $line.Trim()
        if (-not $trimmed) { continue }
        if ($trimmed -match '^\s*warning:\s*no packages being tested depend on matches for pattern') { continue }
        if ($trimmed -match '^#\s+\S+' -and $trimmed -notmatch '\.go:\d+') { continue }
        if ($trimmed -match '^(github\.com|gitlab\.com)/\S+(\s+\[[^\]]+\])?$' -and $trimmed -notmatch '\.go:\d+') { continue }
        $filtered.Add($line) | Out-Null
    }
    return $filtered.ToArray()
}

function Extract-BuildErrorLines {
    <# .SYNOPSIS Extract compile-time error lines from Go build output. #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)
    $candidates = Filter-BlockedCompileLines $lines
    $errors = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
    foreach ($raw in $candidates) {
        if ($null -eq $raw) { continue }; $line = $raw.ToString().TrimEnd("`r"); $trimmed = $line.Trim()
        if (-not $trimmed) { continue }
        if ($trimmed -match '\.go:\d+(?::\d+)?:' -or $trimmed -match '^#\s+\S+' -or
            $trimmed -match '\[build failed\]' -or $trimmed -match '(?i)\bbuild failed\b' -or
            $trimmed -match '\[setup failed\]') {
            if ($seen.Add($line)) { $errors.Add($line) | Out-Null }
        }
    }
    return $errors.ToArray()
}

function Extract-ExecutionFailureLines {
    <# .SYNOPSIS Extract execution failure lines (compile + runtime) from Go output. #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)
    $candidates = Filter-BlockedCompileLines $lines
    $errors = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
    foreach ($raw in $candidates) {
        if ($null -eq $raw) { continue }; $line = $raw.ToString().TrimEnd("`r"); $trimmed = $line.Trim()
        if (-not $trimmed) { continue }
        if ($trimmed -match '\.go:\d+(?::\d+)?:' -or $trimmed -match '^#\s+\S+' -or
            $trimmed -match '\[build failed\]' -or $trimmed -match '(?i)\bbuild failed\b' -or
            $trimmed -match '\[setup failed\]' -or
            $trimmed -match '^(?i)panic:' -or $trimmed -match '^(?i)fatal error:' -or
            $trimmed -match '^--- FAIL:\s+' -or $trimmed -match '^\s*FAIL\s+\S+' -or
            $trimmed -match '^\s*exit status \d+\s*$') {
            if ($seen.Add($line)) { $errors.Add($line) | Out-Null }
        }
    }
    return $errors.ToArray()
}

function Extract-RuntimeFailureLines {
    <# .SYNOPSIS Extract ONLY runtime failure lines from Go output (no compile errors). #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)
    $candidates = Filter-BlockedCompileLines $lines
    $errors = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
    foreach ($raw in $candidates) {
        if ($null -eq $raw) { continue }; $line = $raw.ToString().TrimEnd("`r"); $trimmed = $line.Trim()
        if (-not $trimmed) { continue }
        if ($trimmed -match '^(?i)panic:' -or $trimmed -match '^(?i)fatal error:' -or
            $trimmed -match '^(?i)goroutine \d+' -or $trimmed -match '^--- FAIL:\s+' -or
            $trimmed -match '^\s*FAIL\s+\S+' -or $trimmed -match '^\s*exit status \d+\s*$' -or
            $trimmed -match '(?i)signal:\s+' -or $trimmed -match '(?i)runtime error:' -or
            $trimmed -match '\[setup failed\]') {
            if ($seen.Add($line)) { $errors.Add($line) | Out-Null }
        }
    }
    return $errors.ToArray()
}

function Extract-SetupFailedContext {
    <#
    .SYNOPSIS
        Capture preceding context lines before a [setup failed] or [build failed] FAIL line.
    .DESCRIPTION
        Go outputs plain-text error messages before the final "FAIL pkg [setup failed]" line.
        Standard extractors miss these because they don't match .go:line: or panic: patterns.
        This function walks backward from each FAIL marker and captures up to N preceding
        non-empty lines as diagnostic context.
    .PARAMETER lines
        Raw Go test output lines.
    .PARAMETER ContextLineCount
        Max number of preceding lines to capture per FAIL marker (default 10).
    .EXAMPLE
        $context = Extract-SetupFailedContext $rawOutput
        # Returns context lines + the FAIL line for each [setup failed] occurrence
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param(
        [string[]]$lines,
        [int]$ContextLineCount = 10
    )

    if (-not $lines -or $lines.Count -eq 0) { return @() }

    $result = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)

    for ($i = 0; $i -lt $lines.Count; $i++) {
        $raw = $lines[$i]
        if ($null -eq $raw) { continue }
        $trimmed = $raw.ToString().TrimEnd("`r").Trim()
        if ($trimmed -match '\[setup failed\]' -or ($trimmed -match '\[build failed\]' -and $trimmed -match '^\s*FAIL\s+')) {
            # Walk backward to capture context
            $startIdx = [Math]::Max(0, $i - $ContextLineCount)
            for ($j = $startIdx; $j -le $i; $j++) {
                $ctxRaw = $lines[$j]
                if ($null -eq $ctxRaw) { continue }
                $ctxLine = $ctxRaw.ToString().TrimEnd("`r")
                if (-not $ctxLine.Trim()) { continue }
                # Skip noise lines
                if ($ctxLine.Trim() -match '^\s*warning:\s*no packages being tested') { continue }
                if ($seen.Add($ctxLine)) { $result.Add($ctxLine) | Out-Null }
            }
        }
    }

    return $result.ToArray()
}

function Get-RawFallbackLines {
    <#
    .SYNOPSIS
        Return all non-empty filtered lines as fallback when extractors find nothing actionable.
    .DESCRIPTION
        Used when Extract-BuildErrorLines and Extract-ExecutionFailureLines both return empty.
        Returns Filter-BlockedCompileLines output so plain-text error messages are preserved.
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)
    $candidates = Filter-BlockedCompileLines $lines
    $result = [System.Collections.Generic.List[string]]::new()
    foreach ($raw in $candidates) {
        if ($null -eq $raw) { continue }; $line = $raw.ToString().TrimEnd("`r")
        if ($line.Trim()) { $result.Add($line) | Out-Null }
    }
    return $result.ToArray()
}

Export-ModuleMember -Function @(
    'Filter-BlockedCompileLines', 'Extract-BuildErrorLines',
    'Extract-ExecutionFailureLines', 'Extract-RuntimeFailureLines',
    'Extract-SetupFailedContext', 'Get-RawFallbackLines'
)
