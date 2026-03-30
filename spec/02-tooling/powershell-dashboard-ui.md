# PowerShell Dashboard UI Spec

## Overview

Transform standard sequential `Write-Host` console output into a structured, bordered, color-coded dashboard UI using ANSI escape sequences and Unicode box-drawing characters. This spec is self-contained — any AI or developer can implement it for any PowerShell script.

---

## 1. Environment Setup

Before any output, configure the console for UTF-8 and ANSI support:

```powershell
[console]::OutputEncoding = [System.Text.Encoding]::UTF8
$ESC = [char]27
```

All color output uses ANSI 24-bit RGB sequences, **not** `Write-Host -ForegroundColor`.

---

## 2. Design Tokens

### 2.1 Color Palette

| Token Name       | Purpose                          | RGB             | Hex       | ANSI Foreground Code              |
|------------------|----------------------------------|-----------------|-----------|-----------------------------------|
| `$cLime`         | Success, checkmarks, bars, titles | `163, 230, 53`  | `#a3e635` | `$ESC[38;2;163;230;53m`           |
| `$cRed`          | Errors, failure counts           | `244, 63, 94`   | `#f43f5e` | `$ESC[38;2;244;63;94m`            |
| `$cPurple`       | Action items, todos              | `168, 85, 247`  | `#a855f7` | `$ESC[38;2;168;85;247m`           |
| `$cCyan`         | Sub-items, info labels           | `6, 182, 212`   | `#06b6d4` | `$ESC[38;2;6;182;212m`            |
| `$cYellow`       | Warnings, phase headers          | `250, 204, 21`  | `#facc15` | `$ESC[38;2;250;204;21m`           |
| `$cMuted`        | Borders, dim text                | `156, 163, 175` | `#9ca3af` | `$ESC[38;2;156;163;175m`          |
| `$cWhite`        | Headers, scores, emphasis        | `255, 255, 255` | `#ffffff` | `$ESC[38;2;255;255;255m`          |
| `$cBarEmpty`     | Empty portion of progress bars   | `100, 100, 100` | `#646464` | `$ESC[38;2;100;100;100m`          |
| `$cReset`        | Reset all formatting             | —               | —         | `$ESC[0m`                         |

### 2.2 Text Formatting

| Style     | ANSI Code       |
|-----------|-----------------|
| Bold      | `$ESC[1m`       |
| Dim       | `$ESC[2m`       |
| Italic    | `$ESC[3m`       |
| Reset     | `$ESC[0m`       |

### 2.3 Variable Definitions (Copy-Paste Block)

```powershell
$ESC    = [char]27
$cLime  = "$ESC[38;2;163;230;53m"
$cRed   = "$ESC[38;2;244;63;94m"
$cPurple= "$ESC[38;2;168;85;247m"
$cCyan  = "$ESC[38;2;6;182;212m"
$cYellow= "$ESC[38;2;250;204;21m"
$cMuted = "$ESC[38;2;156;163;175m"
$cWhite = "$ESC[38;2;255;255;255m"
$cBarE  = "$ESC[38;2;100;100;100m"
$cReset = "$ESC[0m"
$cBold  = "$ESC[1m"
$cDim   = "$ESC[2m"
```

---

## 3. Unicode Characters

### 3.1 Icons

| Symbol | Name              | Unicode  | Usage                    |
|--------|-------------------|----------|--------------------------|
| `⚡`   | Lightning Bolt    | `U+26A1` | Product name prefix      |
| `▶`    | Triangular Bullet | `U+25B6` | Action/step prefix       |
| `✓`    | Checkmark         | `U+2713` | Success indicator        |
| `●`    | Solid Dot         | `U+25CF` | Todo/pending indicator   |
| `✗`    | Ballot X          | `U+2717` | Failure indicator        |

### 3.2 Box-Drawing Characters (Double-Line Style)

| Symbol | Name              | Unicode  |
|--------|-------------------|----------|
| `╔`    | Top-Left Corner   | `U+2554` |
| `╗`    | Top-Right Corner  | `U+2557` |
| `╚`    | Bottom-Left Corner| `U+255A` |
| `╝`    | Bottom-Right Cnr  | `U+255D` |
| `║`    | Vertical Wall     | `U+2551` |
| `═`    | Horizontal Wall   | `U+2550` |
| `╠`    | Left T-Junction   | `U+2560` |
| `╣`    | Right T-Junction  | `U+2563` |

### 3.3 Progress Bar Characters

| Symbol | Name          | Unicode  | Usage              |
|--------|---------------|----------|--------------------|
| `█`    | Full Block    | `U+2588` | Filled portion     |
| `▒`    | Medium Shade  | `U+2592` | Empty portion      |

---

## 4. Progress Bar Function

A reusable function that returns a colored progress bar string.

```powershell
function Get-ProgressBar {
    param (
        [int]$Score,
        [int]$MaxScore = 100,
        [int]$BarWidth = 15
    )

    $percentage   = $Score / $MaxScore
    $filledCount  = [math]::Round($percentage * $BarWidth)
    $emptyCount   = $BarWidth - $filledCount

    $filled = if ($filledCount -gt 0) { "█" * $filledCount } else { "" }
    $empty  = if ($emptyCount  -gt 0) { "▒" * $emptyCount  } else { "" }

    return "${cLime}${filled}${cBarE}${empty}${cReset}"
}
```

**Rules:**
- Bar width is always fixed (default 15 chars).
- Filled portion uses `$cLime`, empty uses `$cBarEmpty`.
- For "PASS"/"FAIL" labels (e.g., Browser Test), skip the bar and print the label in `$cLime` or `$cRed`.

---

## 5. Layout Engine

### 5.1 Box Width

All boxes use a fixed internal content width. Recommended: **48 characters** (50 including `║ ` and ` ║`).

```powershell
$boxWidth = 48
```

### 5.2 Helper Functions

```powershell
function Write-BoxTop {
    param([int]$Width = 48)
    Write-Host "${cMuted}╔$("═" * $Width)╗${cReset}"
}

function Write-BoxBottom {
    param([int]$Width = 48)
    Write-Host "${cMuted}╚$("═" * $Width)╝${cReset}"
}

function Write-BoxDivider {
    param([int]$Width = 48)
    Write-Host "${cMuted}╠$("═" * $Width)╣${cReset}"
}

function Write-BoxLine {
    param(
        [string]$Content,
        [int]$Width = 48
    )
    # $Content may contain ANSI codes, so visible length != string length.
    # Caller must ensure visual content fits within $Width.
    Write-Host "${cMuted}║${cReset} ${Content}"
}

function Write-BoxLineCenter {
    param(
        [string]$Text,
        [int]$Width = 48,
        [string]$Color = $cWhite
    )
    $pad = [math]::Max(0, [math]::Floor(($Width - $Text.Length) / 2))
    $line = (" " * $pad) + $Text
    Write-Host "${cMuted}║${cReset}${Color}${cBold}${line}${cReset}"
}
```

### 5.3 Column Alignment

Use `.PadRight()` and `.PadLeft()` for strict column alignment:

```powershell
# Score row example
$label     = "SEO/GEO/AEO".PadRight(16)
$scoreText = "92/100".PadLeft(7)
$bar       = Get-ProgressBar -Score 92

Write-BoxLine "$cWhite$label $scoreText  $bar"
```

**Column layout for score grid:**

| Column        | Width   | Alignment |
|---------------|---------|-----------|
| Label         | 16 char | Left      |
| Score (N/100) | 7 char  | Right     |
| Gap           | 2 char  | —         |
| Progress Bar  | 15 char | Left      |

---

## 6. Output Sections (Top-to-Bottom)

The dashboard is rendered in sequential sections. Each section is a self-contained block.

### 6.1 Header Banner

```
  ⚡  PRODUCT_NAME v1.2.0
  ────────────────────────────────
```

- Lightning bolt in `$cLime`, product name in `$cWhite` + `$cBold`.
- Horizontal rule using `─` (`U+2500`) in `$cMuted`, width matches box width.

### 6.2 Scan Summary Block (No Box)

```
  ▶ Scanning...        47 issues found
  ▶ Auto-fixing...     12 resolved ✓
  ▶ 5 agents running   SEO · Perf · Security · Quality · Browser
```

- `▶` in `$cCyan`.
- Labels in `$cCyan`.
- Issue count in `$cRed`.
- Resolved count in `$cLime` with `✓`.
- Agent names in `$cMuted`, separated by ` · `.
- Use `.PadRight()` on labels to align the right column.

### 6.3 Score Dashboard Box

```
╔══════════════════════════════════════════════════╗
║       U L T R A S H I P   S C O R E             ║
╠══════════════════════════════════════════════════╣
║                                                  ║
║  SEO/GEO/AEO    92/100  ███████████████▒▒        ║
║  Performance    87/100  █████████████▒▒▒▒        ║
║  Security       95/100  ██████████████▒▒         ║
║  Code Quality   88/100  █████████████▒▒▒         ║
║  Browser Test   PASS    ███████████████           ║
║                                                  ║
╠══════════════════════════════════════════════════╣
║                                                  ║
║  OVERALL        90/100                           ║
║  STATUS         [?] READY TO SHIP                ║
║                                                  ║
╚══════════════════════════════════════════════════╝
```

**Rendering rules:**
- Title: spaced-out letters (`S P A C E D`), centered, in `$cWhite` + `$cBold`.
- Score rows: label in `$cWhite`, score in `$cWhite`, bar from `Get-ProgressBar`.
- "PASS" keyword: rendered in `$cLime` + `$cBold`, no bar.
- "FAIL" keyword: rendered in `$cRed` + `$cBold`, no bar.
- OVERALL score: `$cWhite` + `$cBold`.
- STATUS text: `$cLime` if passing, `$cRed` if failing.
- `[?]` prefix on status in `$cYellow`.

### 6.4 Resolution Summary Block (No Box)

```
  ✓ Fixed:  12 issues auto-resolved
  ● Todo:    2 manual items remaining
```

- `✓` in `$cLime`, "Fixed:" label in `$cLime`.
- `●` in `$cYellow`, "Todo:" label in `$cYellow`.
- Counts in `$cWhite`.
- Description text in `$cMuted`.

### 6.5 Footer Tagline

```
  Ship it. One command. Production-ready.
```

- Entire line in `$cLime` + `$cBold`.

---

## 7. Indentation Rules

- All content outside boxes: **2-space indent** from left edge.
- Content inside boxes: **1 space** after `║` and before closing `║`.
- Blank lines between sections: exactly **1 blank line**.

---

## 8. Data Contract

The rendering functions should accept a data object, not hardcoded values:

```powershell
$dashboardData = @{
    ProductName = "ULTRASHIP"
    Version     = "v1.2.0"
    IssuesFound = 47
    IssuesFixed = 12
    AgentCount  = 5
    Agents      = @("SEO", "Perf", "Security", "Quality", "Browser")
    Scores      = [ordered]@{
        "SEO/GEO/AEO"   = 92
        "Performance"    = 87
        "Security"       = 95
        "Code Quality"   = 88
        "Browser Test"   = "PASS"   # string = label, int = score
    }
    OverallScore = 90
    Status       = "READY TO SHIP"
    StatusReady  = $true
    ManualTodos  = 2
}
```

**Type rules:**
- If a score value is `[int]` → render numeric score + progress bar.
- If a score value is `[string]` → render as label (PASS/FAIL) with appropriate color, no bar.

---

## 9. Composability

The spec is designed so each section is a standalone function:

```powershell
function Write-DashboardHeader  ($data) { ... }
function Write-ScanSummary      ($data) { ... }
function Write-ScoreBox         ($data) { ... }
function Write-ResolutionSummary($data) { ... }
function Write-FooterTagline    ($data) { ... }

# Main render
function Write-Dashboard ($data) {
    Write-Host ""
    Write-DashboardHeader   $data
    Write-Host ""
    Write-ScanSummary       $data
    Write-Host ""
    Write-ScoreBox          $data
    Write-Host ""
    Write-ResolutionSummary $data
    Write-Host ""
    Write-FooterTagline     $data
    Write-Host ""
}
```

---

## 10. Terminal Compatibility Notes

| Requirement          | Minimum                          |
|----------------------|----------------------------------|
| PowerShell version   | 7.0+ (pwsh) recommended          |
| Windows Terminal     | Any version (ANSI native)        |
| Legacy `conhost.exe` | May not render ANSI; add VT check|
| Encoding             | UTF-8 required                   |

**Optional VT fallback check:**

```powershell
$vtSupported = $null -ne $env:WT_SESSION -or $PSVersionTable.PSVersion.Major -ge 7
if (-not $vtSupported) {
    Write-Warning "Terminal may not support ANSI colors. Use Windows Terminal or pwsh 7+."
}
```

---

## 11. Complete Rendering Example

For reference, a full rendering call:

```powershell
# 1. Setup
[console]::OutputEncoding = [System.Text.Encoding]::UTF8
# ... define color vars, functions ...

# 2. Collect results from your script logic
$data = @{
    ProductName  = "ULTRASHIP"
    Version      = "v1.2.0"
    IssuesFound  = 47
    IssuesFixed  = 12
    AgentCount   = 5
    Agents       = @("SEO", "Perf", "Security", "Quality", "Browser")
    Scores       = [ordered]@{
        "SEO/GEO/AEO" = 92
        "Performance"  = 87
        "Security"     = 95
        "Code Quality" = 88
        "Browser Test" = "PASS"
    }
    OverallScore = 90
    Status       = "READY TO SHIP"
    StatusReady  = $true
    ManualTodos  = 2
}

# 3. Render
Write-Dashboard $data
```

This produces the exact visual layout shown in the reference image.

---

## 12. Adapting to `run.ps1` Phases

`run.ps1` has two primary dashboard-producing commands: **TC** (test-cover) and **PC** (pre-commit). Each runs through a pipeline of phases. This section maps each phase to the dashboard UI components defined above.

### 12.1 Phase Registry

Each phase is tracked in a `$phases` ordered dictionary. As each phase completes, it records its status so the final dashboard can render them all.

```powershell
$phases = [ordered]@{}

function Register-Phase {
    param(
        [string]$Name,
        [string]$Status,   # "pass", "fail", "skip", "warn"
        [string]$Detail    # optional one-line summary
    )
    $phases[$Name] = @{ Status = $Status; Detail = $Detail }
}
```

### 12.2 Phase Definitions — TC (Test Coverage) Command

| #  | Phase Name              | Source Function / Code Block              | Success Condition                         | Dashboard Label          |
|----|-------------------------|-------------------------------------------|-------------------------------------------|--------------------------|
| 1  | Git Pull                | `Invoke-GitPull`                          | No merge conflicts                        | `Git Pull`               |
| 2  | Dependency Fetch        | `Invoke-FetchLatest` → `go mod tidy`      | `$LASTEXITCODE -eq 0`                     | `Dependencies`           |
| 3  | Data Cleanup            | `Remove-Item data/` + `Cleaned data/`     | Directory removed                         | `Data Cleanup`           |
| 4  | SafeTest Boundary Lint  | `check-safetest-boundaries.ps1`           | `$LASTEXITCODE -eq 0`                     | `SafeTest Lint`          |
| 5  | Go Auto-Fixer           | `go run ./scripts/autofix/`               | `$LASTEXITCODE -eq 0`                     | `Auto-Fixer`             |
| 6  | Syntax Pre-Check        | `go run ./scripts/bracecheck/`            | `$LASTEXITCODE -eq 0`, file count logged  | `Syntax Check`           |
| 7  | Pre-Coverage Compile    | Parallel `go test -run '^$'` per package  | 0 blocked packages                        | `Compile Check`          |
| 8  | Per-File Split Recovery | Split blocked pkgs, recheck per-file      | Recovered file count                      | `Split Recovery`         |
| 9  | Coverage Run            | `go test -coverprofile` per package       | All packages produce profiles             | `Coverage Run`           |
| 10 | Coverage Merge & Report | Merge profiles, generate HTML             | Report files generated                    | `Coverage Report`        |

### 12.3 Phase Definitions — PC (Pre-Commit) Command

| #  | Phase Name              | Source Function / Code Block              | Success Condition                         | Dashboard Label          |
|----|-------------------------|-------------------------------------------|-------------------------------------------|--------------------------|
| 1  | Regression Guard        | `check-integrated-regressions.ps1`        | `$LASTEXITCODE -eq 0`                     | `Regression Guard`       |
| 2  | SafeTest Boundary Lint  | `check-safetest-boundaries.ps1`           | `$LASTEXITCODE -eq 0`                     | `SafeTest Lint`          |
| 3  | Go Auto-Fixer           | `go run ./scripts/autofix/`               | `$LASTEXITCODE -eq 0`                     | `Auto-Fixer`             |
| 4  | Syntax Pre-Check        | `go run ./scripts/bracecheck/`            | `$LASTEXITCODE -eq 0`, file count logged  | `Syntax Check`           |
| 5  | API Compile Check       | `go test -c` per Coverage* package        | 0 failures                                | `API Compile Check`      |

### 12.4 Phase Status Mapping to UI

Each phase status maps to specific colors and icons:

| Status   | Icon | Color      | Token      |
|----------|------|------------|------------|
| `pass`   | `✓`  | Lime Green | `$cLime`   |
| `fail`   | `✗`  | Red        | `$cRed`    |
| `skip`   | `⊘`  | Muted Gray | `$cMuted`  |
| `warn`   | `⚠`  | Yellow     | `$cYellow` |

### 12.5 Phase Summary Box

After all phases complete, render a summary box using the score dashboard pattern:

```
╔══════════════════════════════════════════════════╗
║         P H A S E   S U M M A R Y               ║
╠══════════════════════════════════════════════════╣
║                                                  ║
║  ✓ Git Pull            pulled 6 files            ║
║  ✓ Dependencies        up to date                ║
║  ✓ Data Cleanup        cleaned                   ║
║  ✓ SafeTest Lint       all clean                 ║
║  ✓ Auto-Fixer          no fixable issues         ║
║  ✓ Syntax Check        209 files parsed OK       ║
║  ✓ Compile Check       90/90 passed              ║
║  ⊘ Split Recovery      not needed                ║
║  ✓ Coverage Run        88 packages               ║
║  ✓ Coverage Report     generated                 ║
║                                                  ║
╠══════════════════════════════════════════════════╣
║                                                  ║
║  PHASES      10/10 passed                        ║
║  STATUS      ✓ READY TO COMMIT                   ║
║                                                  ║
╚══════════════════════════════════════════════════╝
```

**Rendering rules:**
- Status icon and label use the color from §12.4.
- Detail text uses `$cMuted`.
- Label column: `.PadRight(20)`.
- If any phase is `fail`, STATUS becomes `✗ BLOCKED` in `$cRed`.
- If any phase is `warn` but none `fail`, STATUS becomes `⚠ REVIEW` in `$cYellow`.

### 12.6 Live Phase Progress (Inline)

During execution, each phase prints a single-line status as it starts and completes:

```powershell
# On phase start:
Write-Host "  ${cCyan}▶${cReset} ${cWhite}Syntax Check${cReset}${cMuted}...${cReset}"

# On phase complete (overwrite or append):
Write-Host "  ${cLime}✓${cReset} ${cWhite}Syntax Check${cReset}  ${cMuted}209 files parsed OK${cReset}"

# On phase fail:
Write-Host "  ${cRed}✗${cReset} ${cWhite}Compile Check${cReset}  ${cRed}3 packages blocked${cReset}"
```

### 12.7 Data Contract for `run.ps1`

Extend the generic data contract (§8) with phase-specific fields:

```powershell
$dashboardData = @{
    ProductName  = "run.ps1"
    Version      = "TC"                  # or "PC"
    Command      = "test-cover"          # or "pre-commit"
    
    # Phase results (ordered)
    Phases       = [ordered]@{
        "Git Pull"         = @{ Status = "pass"; Detail = "pulled 6 files" }
        "Dependencies"     = @{ Status = "pass"; Detail = "up to date" }
        "Data Cleanup"     = @{ Status = "pass"; Detail = "cleaned" }
        "SafeTest Lint"    = @{ Status = "pass"; Detail = "all clean" }
        "Auto-Fixer"       = @{ Status = "pass"; Detail = "no fixable issues" }
        "Syntax Check"     = @{ Status = "pass"; Detail = "209 files parsed OK" }
        "Compile Check"    = @{ Status = "pass"; Detail = "90/90 passed" }
        "Split Recovery"   = @{ Status = "skip"; Detail = "not needed" }
        "Coverage Run"     = @{ Status = "pass"; Detail = "88 packages" }
        "Coverage Report"  = @{ Status = "pass"; Detail = "generated" }
    }
    
    # Score metrics (TC only — from coverage results)
    Scores       = [ordered]@{
        "Overall Coverage" = 97    # percentage from merged profile
        "Package Pass"     = 88    # packages that passed / total
        "Compile Pass"     = 90    # packages that compiled / total
        "Lint Check"       = "PASS"
        "Syntax Check"     = "PASS"
    }
    
    # Numeric rollups
    TotalPackages   = 90
    PassedPackages  = 88
    BlockedPackages = 2
    OverallCoverage = 97.3
    
    # Status
    Status       = "READY TO COMMIT"
    StatusReady  = $true
    
    # Issue tracking
    IssuesFound  = 5
    IssuesFixed  = 3
    ManualTodos  = 2
    
    # Blocked package details (for error section)
    BlockedDetails = @(
        @{ Package = "mapdiffinternal"; Errors = @("undefined: someSym") }
        @{ Package = "corepayloadtests"; Errors = @("type mismatch in Coverage20") }
    )
}
```

### 12.8 Error Detail Section

When packages are blocked, render an error detail section below the dashboard box:

```
  ── Blocked Packages ──────────────────────────────

  ✗ mapdiffinternal
      Coverage12_Gaps_test.go:45 [undefined] undefined: someSym
      Coverage12_Gaps_test.go:52 [type-mismatch] cannot use x as y

  ✗ corepayloadtests
      Coverage20_test.go:18 [undefined] undefined: stringerImpl

  ─────────────────────────────────────────────────
```

**Rendering rules:**
- Package name in `$cRed` + `$cBold`.
- Error lines: file:line in `$cYellow`, category in `$cMuted` brackets, message in `$cWhite`.
- Uses single-line box drawing `─` (`U+2500`) for dividers.
- Error categories come from `ParseCompileErrors`: `arg-count`, `undefined`, `type-mismatch`, `missing-member`, `field-vs-method`, `other`.

### 12.9 Integration Points

To integrate into `run.ps1`, add `Register-Phase` calls at each phase boundary:

```powershell
# Example: in Invoke-FetchLatest
function Invoke-FetchLatest {
    Invoke-GitPull
    Register-Phase "Git Pull" "pass" "pulled from remote"
    
    Write-Header "Fetching latest dependencies"
    go mod tidy
    if ($LASTEXITCODE -eq 0) {
        Register-Phase "Dependencies" "pass" "up to date"
    } else {
        Register-Phase "Dependencies" "fail" "go mod tidy failed"
    }
    # ... rest of function
}

# At end of TC or PC, render the dashboard:
Write-PhaseSummaryBox $phases
Write-Dashboard $dashboardData
```

### 12.10 Shared vs Command-Specific Phases

| Phase              | TC | PC | Notes                                    |
|--------------------|----|----|------------------------------------------|
| Git Pull           | ✓  | ✗  | TC pulls; PC assumes already pulled      |
| Dependencies       | ✓  | ✗  | TC runs `go mod tidy`                    |
| Data Cleanup       | ✓  | ✗  | TC cleans `data/` dir                    |
| Regression Guard   | ✗  | ✓  | PC-only CaseV1/corejson regression scan  |
| SafeTest Lint      | ✓  | ✓  | Both commands run boundary check         |
| Auto-Fixer         | ✓  | ✓  | Both, skippable via `--no-autofix`       |
| Syntax Check       | ✓  | ✓  | Both, skippable via `--skip-bracecheck`  |
| Compile Check      | ✓  | ✓  | TC: all pkgs; PC: Coverage* pkgs only    |
| Split Recovery     | ✓  | ✗  | TC-only per-file split for blocked pkgs  |
| Coverage Run       | ✓  | ✗  | TC-only actual test execution            |
| Coverage Report    | ✓  | ✗  | TC-only merge + HTML generation          |

