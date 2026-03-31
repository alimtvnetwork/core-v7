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
            $trimmed -match '\[build failed\]' -or $trimmed -match '(?i)\bbuild failed\b') {
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
            $trimmed -match '(?i)signal:\s+' -or $trimmed -match '(?i)runtime error:') {
            if ($seen.Add($line)) { $errors.Add($line) | Out-Null }
        }
    }
    return $errors.ToArray()
}

Export-ModuleMember -Function @(
    'Filter-BlockedCompileLines', 'Extract-BuildErrorLines',
    'Extract-ExecutionFailureLines', 'Extract-RuntimeFailureLines'
)
