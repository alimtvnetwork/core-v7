# ─────────────────────────────────────────────────────────────────────────────
# DashboardBoxPrimitives.psm1 — Progress bar + box-drawing primitives
#
# Dependencies: DashboardTheme.psm1 (script-scope color variables)
# ─────────────────────────────────────────────────────────────────────────────

function Get-AnsiVisualLength {
    <# .SYNOPSIS Calculate visual column width of a string, stripping ANSI escape codes and accounting for wide Unicode chars (e.g. ⚠ = 2 columns). #>
    [CmdletBinding()]
    [OutputType([int])]
    param([Parameter(Mandatory)][string]$Text)
    # Strip ANSI escape sequences: ESC[...m and ESC[...;...m etc.
    $stripped = $Text -replace "$([char]27)\[[0-9;]*m", ''
    $len = 0
    foreach ($ch in $stripped.GetEnumerator()) {
        $cp = [int]$ch
        # Wide characters: CJK, fullwidth, and specific symbols like ⚠ (U+26A0)
        if ($cp -eq 0x26A0 -or
            ($cp -ge 0x2600 -and $cp -le 0x26FF) -or   # Misc Symbols
            ($cp -ge 0x2700 -and $cp -le 0x27BF) -or   # Dingbats
            ($cp -ge 0xFE00 -and $cp -le 0xFE0F) -or   # Variation Selectors
            ($cp -ge 0x1F300 -and $cp -le 0x1F9FF) -or  # Emoji
            ($cp -ge 0x3000 -and $cp -le 0x9FFF) -or    # CJK
            ($cp -ge 0xF900 -and $cp -le 0xFAFF) -or    # CJK Compat
            ($cp -ge 0xFF01 -and $cp -le 0xFF60)) {     # Fullwidth
            $len += 2
        } else {
            $len += 1
        }
    }
    return $len
}

function Get-ProgressBar {
    <# .SYNOPSIS Returns a colored progress bar string using ANSI + Unicode block chars. #>
    [CmdletBinding()]
    [OutputType([string])]
    param([Parameter(Mandatory)][int]$Score, [int]$MaxScore = 100, [int]$BarWidth = 15)

    if ($MaxScore -le 0) { $MaxScore = 100 }
    $percentage  = [math]::Min(1.0, [math]::Max(0.0, $Score / $MaxScore))
    $filledCount = [math]::Round($percentage * $BarWidth)
    $emptyCount  = $BarWidth - $filledCount
    $filled = if ($filledCount -gt 0) { [string]::new([char]0x2588, $filledCount) } else { "" }
    $empty  = if ($emptyCount  -gt 0) { [string]::new([char]0x2592, $emptyCount)  } else { "" }
    return "$($script:cLime)$filled$($script:cBarE)$empty$($script:cReset)"
}

function Write-BoxTop {
    [CmdletBinding()]
    param([int]$Width = $script:BoxWidth)
    Write-Host "$($script:cBorder)╔$("═" * $Width)╗$($script:cReset)"
}

function Write-BoxBottom {
    [CmdletBinding()]
    param([int]$Width = $script:BoxWidth)
    Write-Host "$($script:cBorder)╚$("═" * $Width)╝$($script:cReset)"
}

function Write-BoxDivider {
    [CmdletBinding()]
    param([int]$Width = $script:BoxWidth)
    Write-Host "$($script:cBorder)╠$("═" * $Width)╣$($script:cReset)"
}

function Write-BoxEmptyLine {
    [CmdletBinding()]
    param([int]$Width = $script:BoxWidth)
    Write-Host "$($script:cBorder)║$($script:cReset)$(" " * $Width)$($script:cBorder)║$($script:cReset)"
}

function Write-BoxLine {
    [CmdletBinding()]
    param([string]$Content, [int]$Width = $script:BoxWidth, [int]$VisualLength = -1)
    if ($VisualLength -lt 0) { $VisualLength = Get-AnsiVisualLength $Content }
    $rightPad = [math]::Max(0, $Width - $VisualLength - 1)
    Write-Host "$($script:cBorder)║$($script:cReset) $Content$(" " * $rightPad)$($script:cBorder)║$($script:cReset)"
}

function Write-BoxLineCenter {
    [CmdletBinding()]
    param([string]$Text, [int]$Width = $script:BoxWidth, [string]$Color = "")
    if (-not $Color) { $Color = $script:cWhite }
    $textLen = $Text.Length
    $leftPad = [math]::Max(0, [math]::Floor(($Width - $textLen) / 2))
    $rightPad = [math]::Max(0, $Width - $textLen - $leftPad)
    $line = (" " * $leftPad) + $Text + (" " * $rightPad)
    Write-Host "$($script:cBorder)║$($script:cReset)$Color$($script:cBold)$line$($script:cReset)$($script:cBorder)║$($script:cReset)"
}

Export-ModuleMember -Function @(
    'Get-AnsiVisualLength',
    'Get-ProgressBar', 'Write-BoxTop', 'Write-BoxBottom', 'Write-BoxDivider',
    'Write-BoxEmptyLine', 'Write-BoxLine', 'Write-BoxLineCenter'
)
