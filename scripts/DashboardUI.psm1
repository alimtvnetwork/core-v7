# ─────────────────────────────────────────────────────────────────────────────
# DashboardUI.psm1 — Reusable PowerShell Dashboard Rendering Module
#
# Implements the spec at: spec/02-tooling/powershell-dashboard-ui.md
#
# Usage:
#   Import-Module ./scripts/DashboardUI.psm1 -Force
#   Initialize-DashboardUI              # auto-detects theme
#   Initialize-DashboardUI -Theme light # explicit theme
#   Write-Dashboard $data               # render full dashboard
#
# Requirements: PowerShell 7+, UTF-8 terminal, ANSI color support
# ─────────────────────────────────────────────────────────────────────────────

# ═══════════════════════════════════════════════════════════════════════════════
# §1  Environment Setup
# ═══════════════════════════════════════════════════════════════════════════════

$script:ESC    = [char]27
$script:cReset = "$script:ESC[0m"
$script:cBold  = "$script:ESC[1m"
$script:cDim   = "$script:ESC[2m"

# Default box width (internal content area, excluding ║ walls)
$script:BoxWidth = 48

# ═══════════════════════════════════════════════════════════════════════════════
# §13  Theme Detection & Color Initialization
# ═══════════════════════════════════════════════════════════════════════════════

function Get-TerminalTheme {
    [CmdletBinding()]
    [OutputType([string])]
    param()

    # Priority 1: Explicit override via environment variable
    if ($env:DASHBOARD_THEME) {
        return $env:DASHBOARD_THEME.ToLower()
    }

    # Priority 2: Windows Terminal settings JSON
    if ($env:WT_SESSION -and $env:LOCALAPPDATA) {
        $wtSettings = Join-Path $env:LOCALAPPDATA "Packages\Microsoft.WindowsTerminal_8wekyb3d8bbwe\LocalState\settings.json"
        if (Test-Path $wtSettings) {
            try {
                $json = Get-Content $wtSettings -Raw | ConvertFrom-Json
                $schemeName = $json.profiles.defaults.colorScheme
                if (-not $schemeName -and $json.defaultProfile) {
                    $schemeName = $json.profiles.list |
                        Where-Object { $_.guid -eq $json.defaultProfile } |
                        Select-Object -ExpandProperty colorScheme -ErrorAction SilentlyContinue
                }
                if ($schemeName) {
                    $scheme = $json.schemes | Where-Object { $_.name -eq $schemeName }
                    if ($scheme -and $scheme.background) {
                        $bg = $scheme.background -replace '^#', ''
                        $r = [convert]::ToInt32($bg.Substring(0,2), 16)
                        $g = [convert]::ToInt32($bg.Substring(2,2), 16)
                        $b = [convert]::ToInt32($bg.Substring(4,2), 16)
                        $luminance = (0.2126 * $r + 0.7152 * $g + 0.0722 * $b) / 255
                        return $(if ($luminance -lt 0.5) { "dark" } else { "light" })
                    }
                }
            } catch { }
        }
    }

    # Priority 3: macOS/Linux — query terminal via OSC 11
    if ($IsLinux -or $IsMacOS) {
        try {
            $sttyOld = & stty -g 2>/dev/null
            & stty raw -echo min 0 time 1 2>/dev/null
            [Console]::Write("$([char]27)]11;?$([char]27)\")
            Start-Sleep -Milliseconds 100
            $response = ""
            while ([Console]::KeyAvailable) {
                $response += [char][Console]::Read()
            }
            & stty $sttyOld 2>/dev/null
            if ($response -match 'rgb:([0-9a-f]{2,4})/([0-9a-f]{2,4})/([0-9a-f]{2,4})') {
                $r = [convert]::ToInt32($Matches[1].Substring(0,2), 16)
                $g = [convert]::ToInt32($Matches[2].Substring(0,2), 16)
                $b = [convert]::ToInt32($Matches[3].Substring(0,2), 16)
                $luminance = (0.2126 * $r + 0.7152 * $g + 0.0722 * $b) / 255
                return $(if ($luminance -lt 0.5) { "dark" } else { "light" })
            }
        } catch { }
    }

    # Priority 4: PowerShell host background heuristic
    try {
        $bg = $Host.UI.RawUI.BackgroundColor
        $lightBgs = @("White", "Gray", "Yellow", "Cyan")
        if ($bg -in $lightBgs) { return "light" }
    } catch { }

    # Default: dark
    return "dark"
}

function Set-ThemeColors {
    [CmdletBinding()]
    param([string]$Theme = "dark")

    $e = $script:ESC
    if ($Theme -eq "light") {
        $script:cLime   = "$e[38;2;22;163;74m"
        $script:cRed    = "$e[38;2;185;28;28m"
        $script:cPurple = "$e[38;2;109;40;217m"
        $script:cCyan   = "$e[38;2;14;116;144m"
        $script:cYellow = "$e[38;2;161;98;7m"
        $script:cMuted  = "$e[38;2;107;114;128m"
        $script:cWhite  = "$e[38;2;15;23;42m"
        $script:cBarE   = "$e[38;2;209;213;219m"
        $script:cBorder = "$e[38;2;156;163;175m"
    } else {
        $script:cLime   = "$e[38;2;163;230;53m"
        $script:cRed    = "$e[38;2;244;63;94m"
        $script:cPurple = "$e[38;2;168;85;247m"
        $script:cCyan   = "$e[38;2;6;182;212m"
        $script:cYellow = "$e[38;2;250;204;21m"
        $script:cMuted  = "$e[38;2;156;163;175m"
        $script:cWhite  = "$e[38;2;255;255;255m"
        $script:cBarE   = "$e[38;2;100;100;100m"
        $script:cBorder = "$e[38;2;156;163;175m"
    }
}

function Initialize-DashboardUI {
    <#
    .SYNOPSIS
        Initialize the dashboard module: UTF-8 encoding + theme colors.
    .PARAMETER Theme
        Force "dark" or "light". Omit to auto-detect.
    #>
    [CmdletBinding()]
    param([string]$Theme)

    [console]::OutputEncoding = [System.Text.Encoding]::UTF8

    if (-not $Theme) {
        $Theme = Get-TerminalTheme
    }
    $script:CurrentTheme = $Theme
    Set-ThemeColors $Theme
}

# ═══════════════════════════════════════════════════════════════════════════════
# §4  Progress Bar
# ═══════════════════════════════════════════════════════════════════════════════

function Get-ProgressBar {
    <#
    .SYNOPSIS
        Returns a colored progress bar string using ANSI + Unicode block chars.
    .PARAMETER Score
        Current value (0–MaxScore).
    .PARAMETER MaxScore
        Maximum value (default 100).
    .PARAMETER BarWidth
        Character width of the bar (default 15).
    #>
    [CmdletBinding()]
    [OutputType([string])]
    param(
        [Parameter(Mandatory)]
        [int]$Score,
        [int]$MaxScore = 100,
        [int]$BarWidth = 15
    )

    if ($MaxScore -le 0) { $MaxScore = 100 }
    $percentage  = [math]::Min(1.0, [math]::Max(0.0, $Score / $MaxScore))
    $filledCount = [math]::Round($percentage * $BarWidth)
    $emptyCount  = $BarWidth - $filledCount

    $filled = if ($filledCount -gt 0) { [string]::new([char]0x2588, $filledCount) } else { "" }
    $empty  = if ($emptyCount  -gt 0) { [string]::new([char]0x2592, $emptyCount)  } else { "" }

    return "$($script:cLime)$filled$($script:cBarE)$empty$($script:cReset)"
}

# ═══════════════════════════════════════════════════════════════════════════════
# §5  Box-Drawing Layout Helpers
# ═══════════════════════════════════════════════════════════════════════════════

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
    <#
    .SYNOPSIS
        Write content inside box walls. Caller ensures visual width fits.
        Pads with spaces to fill the box width for clean right wall.
    .PARAMETER Content
        ANSI-colored content string.
    .PARAMETER VisualLength
        The visible character count of Content (excluding ANSI codes).
        Used to calculate right-side padding. Default: no right wall.
    #>
    [CmdletBinding()]
    param(
        [string]$Content,
        [int]$Width = $script:BoxWidth,
        [int]$VisualLength = -1
    )

    if ($VisualLength -ge 0) {
        $rightPad = [math]::Max(0, $Width - $VisualLength - 1)
        Write-Host "$($script:cBorder)║$($script:cReset) $Content$(" " * $rightPad)$($script:cBorder)║$($script:cReset)"
    } else {
        Write-Host "$($script:cBorder)║$($script:cReset) $Content"
    }
}

function Write-BoxLineCenter {
    <#
    .SYNOPSIS
        Write centered text inside box walls.
    #>
    [CmdletBinding()]
    param(
        [string]$Text,
        [int]$Width = $script:BoxWidth,
        [string]$Color = ""
    )

    if (-not $Color) { $Color = $script:cWhite }
    $textLen = $Text.Length
    $leftPad = [math]::Max(0, [math]::Floor(($Width - $textLen) / 2))
    $rightPad = [math]::Max(0, $Width - $textLen - $leftPad)
    $line = (" " * $leftPad) + $Text + (" " * $rightPad)
    Write-Host "$($script:cBorder)║$($script:cReset)$Color$($script:cBold)$line$($script:cReset)$($script:cBorder)║$($script:cReset)"
}

# ═══════════════════════════════════════════════════════════════════════════════
# §6.1  Header Banner
# ═══════════════════════════════════════════════════════════════════════════════

function Write-DashboardHeader {
    <#
    .SYNOPSIS
        Render the product header with lightning bolt and horizontal rule.
    .PARAMETER Data
        Hashtable with ProductName and Version keys.
    #>
    [CmdletBinding()]
    param([hashtable]$Data)

    $name    = if ($Data.ProductName) { $Data.ProductName } else { "Dashboard" }
    $version = if ($Data.Version)     { $Data.Version }     else { "" }

    Write-Host "  $($script:cLime)⚡$($script:cReset)  $($script:cWhite)$($script:cBold)$name $version$($script:cReset)"
    Write-Host "  $($script:cMuted)$("─" * ($script:BoxWidth - 2))$($script:cReset)"
}

# ═══════════════════════════════════════════════════════════════════════════════
# §6.2  Scan Summary Block
# ═══════════════════════════════════════════════════════════════════════════════

function Write-ScanSummary {
    <#
    .SYNOPSIS
        Render the scan/fix/agents summary (no box).
    .PARAMETER Data
        Hashtable with IssuesFound, IssuesFixed, AgentCount, Agents.
    #>
    [CmdletBinding()]
    param([hashtable]$Data)

    $issuesFound = if ($null -ne $Data.IssuesFound) { $Data.IssuesFound } else { 0 }
    $issuesFixed = if ($null -ne $Data.IssuesFixed) { $Data.IssuesFixed } else { 0 }
    $agentCount  = if ($null -ne $Data.AgentCount)  { $Data.AgentCount }  else { 0 }
    $agents      = if ($Data.Agents) { $Data.Agents } else { @() }

    $labelWidth = 22
    $scanLabel   = "Scanning...".PadRight($labelWidth)
    $fixLabel    = "Auto-fixing...".PadRight($labelWidth)
    $agentLabel  = "$agentCount agents running".PadRight($labelWidth)
    $agentList   = ($agents -join " $($script:cMuted)·$($script:cReset) $($script:cMuted)")

    Write-Host "  $($script:cCyan)▶$($script:cReset) $($script:cCyan)$scanLabel$($script:cReset)$($script:cRed)$issuesFound issues found$($script:cReset)"
    Write-Host "  $($script:cCyan)▶$($script:cReset) $($script:cCyan)$fixLabel$($script:cReset)$($script:cLime)$issuesFixed resolved ✓$($script:cReset)"

    if ($agentCount -gt 0) {
        Write-Host "  $($script:cCyan)▶$($script:cReset) $($script:cCyan)$agentLabel$($script:cReset)$($script:cMuted)$agentList$($script:cReset)"
    }
}

# ═══════════════════════════════════════════════════════════════════════════════
# §6.3  Score Dashboard Box
# ═══════════════════════════════════════════════════════════════════════════════

function Write-ScoreBox {
    <#
    .SYNOPSIS
        Render the bordered score dashboard with progress bars.
    .PARAMETER Data
        Hashtable with Scores (ordered dict), OverallScore, Status, StatusReady.
    .PARAMETER Title
        The spaced-out title text. Default: "S C O R E".
    #>
    [CmdletBinding()]
    param(
        [hashtable]$Data,
        [string]$Title = ""
    )

    $w = $script:BoxWidth

    # Determine title
    if (-not $Title -and $Data.ProductName) {
        $spaced = ($Data.ProductName.ToUpper().ToCharArray() -join " ")
        $Title = "$spaced   S C O R E"
    } elseif (-not $Title) {
        $Title = "S C O R E"
    }

    # Top + title
    Write-BoxTop -Width $w
    Write-BoxLineCenter -Text $Title -Width $w
    Write-BoxDivider -Width $w
    Write-BoxEmptyLine -Width $w

    # Score rows
    $labelCol  = 16
    $scoreCol  = 7
    $barWidth  = 15

    if ($Data.Scores) {
        foreach ($key in $Data.Scores.Keys) {
            $val   = $Data.Scores[$key]
            $label = $key.PadRight($labelCol)

            if ($val -is [int] -or $val -is [double] -or $val -is [decimal]) {
                $intVal    = [int]$val
                $scoreText = "$intVal/100".PadLeft($scoreCol)
                $bar       = Get-ProgressBar -Score $intVal -BarWidth $barWidth
                $visLen    = 1 + $labelCol + 1 + $scoreCol + 2 + $barWidth  # "║ label score  bar"
                Write-BoxLine -Content "$($script:cWhite)$label $scoreText  $bar" -Width $w -VisualLength $visLen
            } else {
                # String value — PASS/FAIL label
                $valStr = "$val"
                if ($valStr -eq "PASS") {
                    $colored = "$($script:cLime)$($script:cBold)$valStr$($script:cReset)"
                } elseif ($valStr -eq "FAIL") {
                    $colored = "$($script:cRed)$($script:cBold)$valStr$($script:cReset)"
                } else {
                    $colored = "$($script:cWhite)$valStr$($script:cReset)"
                }
                $visLen = 1 + $labelCol + 1 + $valStr.Length
                Write-BoxLine -Content "$($script:cWhite)$label $colored" -Width $w -VisualLength $visLen
            }
        }
    }

    Write-BoxEmptyLine -Width $w
    Write-BoxDivider -Width $w
    Write-BoxEmptyLine -Width $w

    # Overall score
    $overallLabel = "OVERALL".PadRight($labelCol)
    $overallVal   = if ($null -ne $Data.OverallScore) { "$($Data.OverallScore)/100" } else { "—" }
    $overallVal   = $overallVal.PadLeft($scoreCol)
    $visLen = 1 + $labelCol + 1 + $scoreCol
    Write-BoxLine -Content "$($script:cWhite)$($script:cBold)$overallLabel $overallVal$($script:cReset)" -Width $w -VisualLength $visLen

    # Status
    $statusLabel = "STATUS".PadRight($labelCol)
    $statusText  = if ($Data.Status) { $Data.Status } else { "UNKNOWN" }
    $statusReady = if ($null -ne $Data.StatusReady) { $Data.StatusReady } else { $false }
    $statusIcon  = "$($script:cYellow)[?]$($script:cReset) "
    if ($statusReady) {
        $statusColor = "$($script:cLime)"
    } else {
        $statusColor = "$($script:cRed)"
    }
    $visLen = 1 + $labelCol + 1 + 4 + $statusText.Length  # "[?] " = 4
    Write-BoxLine -Content "$($script:cWhite)$statusLabel $statusIcon$statusColor$statusText$($script:cReset)" -Width $w -VisualLength $visLen

    Write-BoxEmptyLine -Width $w
    Write-BoxBottom -Width $w
}

# ═══════════════════════════════════════════════════════════════════════════════
# §6.4  Resolution Summary
# ═══════════════════════════════════════════════════════════════════════════════

function Write-ResolutionSummary {
    <#
    .SYNOPSIS
        Render the fixed/todo summary lines (no box).
    .PARAMETER Data
        Hashtable with IssuesFixed, ManualTodos.
    #>
    [CmdletBinding()]
    param([hashtable]$Data)

    $fixed = if ($null -ne $Data.IssuesFixed) { $Data.IssuesFixed } else { 0 }
    $todos = if ($null -ne $Data.ManualTodos) { $Data.ManualTodos } else { 0 }

    Write-Host "  $($script:cLime)✓$($script:cReset) $($script:cLime)Fixed:$($script:cReset)  $($script:cWhite)$fixed$($script:cReset) $($script:cMuted)issues auto-resolved$($script:cReset)"

    if ($todos -gt 0) {
        Write-Host "  $($script:cYellow)●$($script:cReset) $($script:cYellow)Todo:$($script:cReset)   $($script:cWhite)$todos$($script:cReset) $($script:cMuted)manual items remaining$($script:cReset)"
    }
}

# ═══════════════════════════════════════════════════════════════════════════════
# §6.5  Footer Tagline
# ═══════════════════════════════════════════════════════════════════════════════

function Write-FooterTagline {
    <#
    .SYNOPSIS
        Render the footer tagline in lime bold.
    .PARAMETER Text
        Custom tagline text. Default: "Ship it. One command. Production-ready."
    #>
    [CmdletBinding()]
    param([string]$Text = "Ship it. One command. Production-ready.")

    Write-Host "  $($script:cLime)$($script:cBold)$Text$($script:cReset)"
}

# ═══════════════════════════════════════════════════════════════════════════════
# §12  Phase Registry & Rendering
# ═══════════════════════════════════════════════════════════════════════════════

# Module-level phase store
$script:Phases = [ordered]@{}

function Register-Phase {
    <#
    .SYNOPSIS
        Record a phase result for the final dashboard summary.
    .PARAMETER Name
        Phase display name (e.g., "Git Pull", "Compile Check").
    .PARAMETER Status
        One of: pass, fail, skip, warn.
    .PARAMETER Detail
        One-line detail string (e.g., "90/90 passed").
    #>
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][string]$Name,
        [Parameter(Mandatory)][ValidateSet("pass","fail","skip","warn")][string]$Status,
        [string]$Detail = ""
    )
    $script:Phases[$Name] = @{ Status = $Status; Detail = $Detail }
}

function Reset-Phases {
    <# .SYNOPSIS Clear all registered phases. #>
    [CmdletBinding()]
    param()
    $script:Phases = [ordered]@{}
}

function Get-PhaseIcon {
    [CmdletBinding()]
    [OutputType([string])]
    param([string]$Status)

    switch ($Status) {
        "pass" { return "$($script:cLime)✓$($script:cReset)" }
        "fail" { return "$($script:cRed)✗$($script:cReset)" }
        "skip" { return "$($script:cMuted)⊘$($script:cReset)" }
        "warn" { return "$($script:cYellow)⚠$($script:cReset)" }
        default { return "$($script:cMuted)?$($script:cReset)" }
    }
}

function Write-PhaseStart {
    <#
    .SYNOPSIS
        Print a single-line phase-started indicator during live execution.
    #>
    [CmdletBinding()]
    param([Parameter(Mandatory)][string]$Name)

    Write-Host "  $($script:cCyan)▶$($script:cReset) $($script:cWhite)$Name$($script:cReset)$($script:cMuted)...$($script:cReset)"
}

function Write-PhaseEnd {
    <#
    .SYNOPSIS
        Print a single-line phase-completed indicator during live execution.
    #>
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][string]$Name,
        [Parameter(Mandatory)][ValidateSet("pass","fail","skip","warn")][string]$Status,
        [string]$Detail = ""
    )

    $icon = Get-PhaseIcon $Status
    $detailColor = switch ($Status) {
        "fail" { $script:cRed }
        "warn" { $script:cYellow }
        default { $script:cMuted }
    }
    $detailStr = if ($Detail) { "  $detailColor$Detail$($script:cReset)" } else { "" }
    Write-Host "  $icon $($script:cWhite)$Name$($script:cReset)$detailStr"
}

function Write-PhaseSummaryBox {
    <#
    .SYNOPSIS
        Render the bordered phase summary box from registered phases.
    .PARAMETER Phases
        Ordered dictionary of phase results. If omitted, uses module-level $script:Phases.
    #>
    [CmdletBinding()]
    param([System.Collections.Specialized.OrderedDictionary]$Phases)

    if (-not $Phases) { $Phases = $script:Phases }
    if ($Phases.Count -eq 0) { return }

    $w = $script:BoxWidth
    $phaseLabelWidth = 20

    # Title
    Write-BoxTop -Width $w
    Write-BoxLineCenter -Text "P H A S E   S U M M A R Y" -Width $w
    Write-BoxDivider -Width $w
    Write-BoxEmptyLine -Width $w

    # Phase rows
    $passCount = 0
    $failCount = 0
    $warnCount = 0

    foreach ($key in $Phases.Keys) {
        $phase  = $Phases[$key]
        $status = $phase.Status
        $detail = if ($phase.Detail) { $phase.Detail } else { "" }
        $icon   = Get-PhaseIcon $status

        switch ($status) {
            "pass" { $passCount++ }
            "fail" { $failCount++ }
            "warn" { $warnCount++ }
        }

        $label = $key.PadRight($phaseLabelWidth)
        $visLen = 1 + 2 + $phaseLabelWidth + $detail.Length  # "║ ✓ label detail"
        Write-BoxLine -Content "$icon $($script:cWhite)$label$($script:cReset)$($script:cMuted)$detail$($script:cReset)" -Width $w -VisualLength $visLen
    }

    Write-BoxEmptyLine -Width $w
    Write-BoxDivider -Width $w
    Write-BoxEmptyLine -Width $w

    # Summary: PHASES x/y passed
    $total = $Phases.Count
    $phasesLabel = "PHASES".PadRight($phaseLabelWidth - 6)
    $phasesVal   = "$passCount/$total passed"
    $visLen = 1 + $phasesLabel.Length + $phasesVal.Length
    Write-BoxLine -Content "$($script:cWhite)$($script:cBold)$phasesLabel$($script:cReset) $($script:cWhite)$phasesVal$($script:cReset)" -Width $w -VisualLength $visLen

    # STATUS line
    $statusLabel = "STATUS".PadRight($phasesLabel.Length)
    if ($failCount -gt 0) {
        $statusIcon = "$($script:cRed)✗$($script:cReset)"
        $statusText = "$($script:cRed)BLOCKED$($script:cReset)"
        $statusVisText = "✗ BLOCKED"
    } elseif ($warnCount -gt 0) {
        $statusIcon = "$($script:cYellow)⚠$($script:cReset)"
        $statusText = "$($script:cYellow)REVIEW$($script:cReset)"
        $statusVisText = "⚠ REVIEW"
    } else {
        $statusIcon = "$($script:cLime)✓$($script:cReset)"
        $statusText = "$($script:cLime)READY TO COMMIT$($script:cReset)"
        $statusVisText = "✓ READY TO COMMIT"
    }
    $visLen = 1 + $statusLabel.Length + 1 + $statusVisText.Length
    Write-BoxLine -Content "$($script:cWhite)$($script:cBold)$statusLabel$($script:cReset) $statusIcon $statusText" -Width $w -VisualLength $visLen

    Write-BoxEmptyLine -Width $w
    Write-BoxBottom -Width $w
}

# ═══════════════════════════════════════════════════════════════════════════════
# §12.8  Error Detail Section
# ═══════════════════════════════════════════════════════════════════════════════

function Write-BlockedDetails {
    <#
    .SYNOPSIS
        Render blocked package error details below the dashboard.
    .PARAMETER BlockedDetails
        Array of @{ Package = "name"; Errors = @("error line 1", ...) }.
    #>
    [CmdletBinding()]
    param([array]$BlockedDetails)

    if (-not $BlockedDetails -or $BlockedDetails.Count -eq 0) { return }

    $dividerWidth = $script:BoxWidth
    Write-Host ""
    Write-Host "  $($script:cMuted)── Blocked Packages $("─" * ($dividerWidth - 22))$($script:cReset)"
    Write-Host ""

    foreach ($block in $BlockedDetails) {
        $pkg = if ($block.Package) { $block.Package } else { "unknown" }
        Write-Host "  $($script:cRed)$($script:cBold)✗ $pkg$($script:cReset)"

        if ($block.Errors) {
            foreach ($errLine in $block.Errors) {
                Write-Host "      $($script:cYellow)$errLine$($script:cReset)"
            }
        }
        Write-Host ""
    }

    Write-Host "  $($script:cMuted)$("─" * $dividerWidth)$($script:cReset)"
}

# ═══════════════════════════════════════════════════════════════════════════════
# §9 + §12  Composite Renderers
# ═══════════════════════════════════════════════════════════════════════════════

function Write-Dashboard {
    <#
    .SYNOPSIS
        Render the full dashboard from a data hashtable.

        Sections rendered (in order):
          1. Header banner          (§6.1)
          2. Scan summary           (§6.2)
          3. Score box              (§6.3)  — if Scores present
          4. Phase summary box      (§12.5) — if Phases present
          5. Blocked details        (§12.8) — if BlockedDetails present
          6. Resolution summary     (§6.4)
          7. Footer tagline         (§6.5)

    .PARAMETER Data
        Hashtable following the data contract (spec §8 / §12.7).
    #>
    [CmdletBinding()]
    param([Parameter(Mandatory)][hashtable]$Data)

    Write-Host ""

    # §6.1 Header
    Write-DashboardHeader -Data $Data
    Write-Host ""

    # §6.2 Scan summary
    if ($null -ne $Data.IssuesFound) {
        Write-ScanSummary -Data $Data
        Write-Host ""
    }

    # §6.3 Score box
    if ($Data.Scores -and $Data.Scores.Count -gt 0) {
        Write-ScoreBox -Data $Data
        Write-Host ""
    }

    # §12.5 Phase summary
    $phases = if ($Data.Phases) { $Data.Phases } else { $script:Phases }
    if ($phases -and $phases.Count -gt 0) {
        Write-PhaseSummaryBox -Phases $phases
        Write-Host ""
    }

    # §14 Per-package coverage table
    if ($Data.CoverageData -and $Data.CoverageData.Count -gt 0) {
        Write-CoverageTable -CoverageData $Data.CoverageData
        Write-Host ""
    }

    # §12.8 Blocked details
    if ($Data.BlockedDetails -and $Data.BlockedDetails.Count -gt 0) {
        Write-BlockedDetails -BlockedDetails $Data.BlockedDetails
        Write-Host ""
    }

    # §6.4 Resolution summary
    if ($null -ne $Data.IssuesFixed -or $null -ne $Data.ManualTodos) {
        Write-ResolutionSummary -Data $Data
        Write-Host ""
    }

    # §6.5 Footer
    $tagline = if ($Data.Tagline) { $Data.Tagline } else { "Ship it. One command. Production-ready." }
    Write-FooterTagline -Text $tagline
    Write-Host ""
}

# ═══════════════════════════════════════════════════════════════════════════════
# §13.6  Theme Test
# ═══════════════════════════════════════════════════════════════════════════════

function Test-DashboardTheme {
    <#
    .SYNOPSIS
        Render a color swatch for both themes to visually verify contrast.
    #>
    [CmdletBinding()]
    param()

    foreach ($theme in @("dark", "light")) {
        Set-ThemeColors $theme
        Write-Host ""
        Write-Host "$($script:cBold)=== Theme: $theme ===$($script:cReset)"
        Write-Host "  $($script:cLime)✓ Success / Lime$($script:cReset)"
        Write-Host "  $($script:cRed)✗ Error / Red$($script:cReset)"
        Write-Host "  $($script:cPurple)● Purple / Todo$($script:cReset)"
        Write-Host "  $($script:cCyan)▶ Cyan / Info$($script:cReset)"
        Write-Host "  $($script:cYellow)⚠ Yellow / Warning$($script:cReset)"
        Write-Host "  $($script:cMuted)Muted text$($script:cReset)"
        Write-Host "  $($script:cWhite)Primary text$($script:cReset)"
        Write-Host "  Bar: $(Get-ProgressBar -Score 73)"
        Write-Host ""
    }
    # Restore
    Set-ThemeColors $script:CurrentTheme
}

# ═══════════════════════════════════════════════════════════════════════════════
# §14  Per-Package Coverage Results Table
# ═══════════════════════════════════════════════════════════════════════════════

function Write-CoverageTable {
    <#
    .SYNOPSIS
        Render a bordered per-package coverage table with progress bars.
        Packages are sorted by coverage % ascending (lowest first).

    .PARAMETER CoverageData
        Array of hashtables, each with:
          - Package  [string]  short package name (e.g. "corestr")
          - Coverage [double]  coverage percentage (0-100)
          - Tests    [int]     number of tests (optional, for display)

    .PARAMETER Title
        Box title. Default: "P A C K A G E   C O V E R A G E"

    .PARAMETER ShowTarget
        If $true, shows a target line at 100%. Default: $true.

    .PARAMETER BarWidth
        Width of progress bars. Default: 12.
    #>
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)]
        [array]$CoverageData,

        [string]$Title = "P A C K A G E   C O V E R A G E",
        [bool]$ShowTarget = $true,
        [int]$BarWidth = 12
    )

    if (-not $CoverageData -or $CoverageData.Count -eq 0) { return }

    # Sort ascending by coverage (worst first)
    $sorted = $CoverageData | Sort-Object { [double]$_.Coverage }

    # Column widths
    $pkgCol   = 24   # package name
    $pctCol   = 7    # "100.0%"
    $testCol  = 5    # test count

    # Calculate box width to fit: "║ pkg  pct  bar  tests ║"
    # 1 + pkgCol + 1 + pctCol + 2 + barWidth + 2 + testCol + 1 = content
    $contentWidth = 1 + $pkgCol + 1 + $pctCol + 2 + $BarWidth + 2 + $testCol + 1
    $w = [math]::Max($script:BoxWidth, $contentWidth)

    # Header
    Write-BoxTop -Width $w
    Write-BoxLineCenter -Text $Title -Width $w
    Write-BoxDivider -Width $w
    Write-BoxEmptyLine -Width $w

    # Column headers
    $hdrPkg   = "Package".PadRight($pkgCol)
    $hdrPct   = "Cov %".PadLeft($pctCol)
    $hdrBar   = "".PadRight($BarWidth)
    $hdrTests = "Tests".PadLeft($testCol)
    $hdrLine  = "$($script:cMuted)$hdrPkg $hdrPct  $hdrBar  $hdrTests$($script:cReset)"
    $hdrVisLen = 1 + $pkgCol + 1 + $pctCol + 2 + $BarWidth + 2 + $testCol
    Write-BoxLine -Content $hdrLine -Width $w -VisualLength $hdrVisLen

    # Separator under headers
    $sepLine = "$($script:cMuted)$("─" * $pkgCol) $("─" * $pctCol)  $("─" * $BarWidth)  $("─" * $testCol)$($script:cReset)"
    Write-BoxLine -Content $sepLine -Width $w -VisualLength $hdrVisLen

    # Data rows
    $totalCoverage = 0.0
    $at100Count    = 0
    $below100Count = 0

    foreach ($entry in $sorted) {
        $pkg = "$($entry.Package)"
        $cov = [double]$entry.Coverage
        $tests = if ($null -ne $entry.Tests) { [int]$entry.Tests } else { 0 }

        $totalCoverage += $cov
        if ($cov -ge 100.0) { $at100Count++ } else { $below100Count++ }

        # Truncate long package names
        if ($pkg.Length -gt $pkgCol) {
            $pkg = $pkg.Substring(0, $pkgCol - 2) + ".."
        }
        $pkgStr = $pkg.PadRight($pkgCol)

        # Format percentage
        $pctStr = ("{0:F1}%" -f $cov).PadLeft($pctCol)

        # Color based on coverage level
        $rowColor = if ($cov -ge 100.0) {
            $script:cLime
        } elseif ($cov -ge 98.0) {
            $script:cWhite
        } elseif ($cov -ge 95.0) {
            $script:cYellow
        } else {
            $script:cRed
        }

        # Progress bar
        $bar = Get-ProgressBar -Score ([int][math]::Round($cov)) -BarWidth $BarWidth

        # Test count
        $testStr = "$tests".PadLeft($testCol)

        $rowContent = "$rowColor$pkgStr$($script:cReset) $rowColor$pctStr$($script:cReset)  $bar  $($script:cMuted)$testStr$($script:cReset)"
        $rowVisLen = 1 + $pkgCol + 1 + $pctCol + 2 + $BarWidth + 2 + $testCol
        Write-BoxLine -Content $rowContent -Width $w -VisualLength $rowVisLen
    }

    Write-BoxEmptyLine -Width $w
    Write-BoxDivider -Width $w
    Write-BoxEmptyLine -Width $w

    # Summary row
    $avgCoverage = if ($sorted.Count -gt 0) { $totalCoverage / $sorted.Count } else { 0 }
    $summaryLabel = "AVERAGE".PadRight($pkgCol)
    $summaryPct   = ("{0:F1}%" -f $avgCoverage).PadLeft($pctCol)
    $summaryBar   = Get-ProgressBar -Score ([int][math]::Round($avgCoverage)) -BarWidth $BarWidth
    $summaryTests = "$($sorted.Count)".PadLeft($testCol)
    $summaryVisLen = 1 + $pkgCol + 1 + $pctCol + 2 + $BarWidth + 2 + $testCol
    Write-BoxLine -Content "$($script:cWhite)$($script:cBold)$summaryLabel$($script:cReset) $($script:cWhite)$($script:cBold)$summaryPct$($script:cReset)  $summaryBar  $($script:cMuted)$summaryTests$($script:cReset)" -Width $w -VisualLength $summaryVisLen

    # 100% vs below counts
    $countLabel = "".PadRight($pkgCol)
    $countText  = "$($script:cLime)$at100Count$($script:cReset)$($script:cMuted) at 100%$($script:cReset)  $($script:cYellow)$below100Count$($script:cReset)$($script:cMuted) below$($script:cReset)"
    $countVisLen = 1 + $pkgCol + 10 + 10  # approximate
    Write-BoxLine -Content "$countLabel$countText" -Width $w -VisualLength $countVisLen

    if ($ShowTarget) {
        $targetLabel = "TARGET".PadRight($pkgCol)
        $targetText  = "$($script:cLime)$($script:cBold)100.0%$($script:cReset)$($script:cMuted) (non-internal packages)$($script:cReset)"
        Write-BoxLine -Content "$($script:cWhite)$targetLabel$($script:cReset)$targetText" -Width $w -VisualLength ($pkgCol + 30)
    }

    Write-BoxEmptyLine -Width $w
    Write-BoxBottom -Width $w
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Exports
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    # Init
    'Initialize-DashboardUI'
    'Get-TerminalTheme'
    'Set-ThemeColors'

    # Progress bar
    'Get-ProgressBar'

    # Box primitives
    'Write-BoxTop'
    'Write-BoxBottom'
    'Write-BoxDivider'
    'Write-BoxEmptyLine'
    'Write-BoxLine'
    'Write-BoxLineCenter'

    # Section renderers
    'Write-DashboardHeader'
    'Write-ScanSummary'
    'Write-ScoreBox'
    'Write-ResolutionSummary'
    'Write-FooterTagline'
    'Write-BlockedDetails'

    # Phase system
    'Register-Phase'
    'Reset-Phases'
    'Write-PhaseStart'
    'Write-PhaseEnd'
    'Write-PhaseSummaryBox'

    # Composite
    'Write-Dashboard'

    # Testing
    'Test-DashboardTheme'
)
