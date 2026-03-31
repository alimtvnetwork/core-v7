# ─────────────────────────────────────────────────────────────────────────────
# ErrorParser.psm1 — Go build/runtime error extraction and classification
#
# Usage:
#   Import-Module ./scripts/ErrorParser.psm1 -Force
#
# Dependencies: None (standalone)
# ─────────────────────────────────────────────────────────────────────────────

function Filter-BlockedCompileLines {
    <#
    .SYNOPSIS
        Remove noisy/irrelevant lines from Go compile output.
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
        if ($trimmed -match '^\s*warning:\s*no packages being tested depend on matches for pattern') { continue }
        if ($trimmed -match '^#\s+\S+' -and $trimmed -notmatch '\.go:\d+') { continue }
        if ($trimmed -match '^(github\.com|gitlab\.com)/\S+(\s+\[[^\]]+\])?$' -and $trimmed -notmatch '\.go:\d+') { continue }
        $filtered.Add($line) | Out-Null
    }

    return $filtered.ToArray()
}

function Extract-BuildErrorLines {
    <#
    .SYNOPSIS
        Extract compile-time error lines from Go build output.
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
        if ($trimmed -match '\.go:\d+(?::\d+)?:' -or $trimmed -match '^#\s+\S+' -or
            $trimmed -match '\[build failed\]' -or $trimmed -match '(?i)\bbuild failed\b') {
            if ($seen.Add($line)) { $errors.Add($line) | Out-Null }
        }
    }

    return $errors.ToArray()
}

function Extract-ExecutionFailureLines {
    <#
    .SYNOPSIS
        Extract execution failure lines (compile + runtime) from Go output.
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
    <#
    .SYNOPSIS
        Extract ONLY runtime failure lines from Go output (no compile errors).
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
        if ($trimmed -match '^(?i)panic:' -or $trimmed -match '^(?i)fatal error:' -or
            $trimmed -match '^(?i)goroutine \d+' -or $trimmed -match '^--- FAIL:\s+' -or
            $trimmed -match '^\s*FAIL\s+\S+' -or $trimmed -match '^\s*exit status \d+\s*$' -or
            $trimmed -match '(?i)signal:\s+' -or $trimmed -match '(?i)runtime error:') {
            if ($seen.Add($line)) { $errors.Add($line) | Out-Null }
        }
    }

    return $errors.ToArray()
}

function Add-BuildErrorsForPackage {
    <#
    .SYNOPSIS
        Accumulate build errors into a per-package hashtable.
    #>
    [CmdletBinding()]
    param([hashtable]$BuildErrorMap, [string]$PackageName, [string[]]$Lines)

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
    #>
    [CmdletBinding()]
    param([hashtable]$FailureMap, [string]$PackageName, [string[]]$Lines)

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

function ParseCompileErrors {
    <#
    .SYNOPSIS
        Parse Go compile error lines into structured objects.
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
            $category = "other"
            if ($msg -match 'too many arguments|not enough arguments') { $category = "arg-count" }
            elseif ($msg -match 'undefined:') { $category = "undefined" }
            elseif ($msg -match 'cannot use .* as') { $category = "type-mismatch" }
            elseif ($msg -match 'has no field or method') { $category = "missing-member" }
            elseif ($msg -match 'cannot call non-function') { $category = "field-vs-method" }
            $errors.Add(@{ file = $file; line = $lineNum; message = $msg; category = $category; raw = $line })
        }
    }
    return $errors.ToArray()
}

Export-ModuleMember -Function @(
    'Filter-BlockedCompileLines',
    'Extract-BuildErrorLines', 'Extract-ExecutionFailureLines', 'Extract-RuntimeFailureLines',
    'Add-BuildErrorsForPackage', 'Add-RuntimeFailuresForPackage',
    'ParseCompileErrors'
)
