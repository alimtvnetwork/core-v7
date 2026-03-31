# latest.branch.ps1
# Finds the most recently updated remote branch (by tip commit date),
# then resolves the exact branch name(s) using the SHA (points-at).

[CmdletBinding()]
param(
  [string]$Remote = "origin",
  [switch]$AllRemotes,
  [switch]$ContainsFallback,
  [int]$Top = 0
)

$ErrorActionPreference = "Stop"

# Ensure we are inside a git repo
git rev-parse --is-inside-work-tree 1>$null 2>$null
if ($LASTEXITCODE -ne 0) { throw "Not inside a Git repository." }

# Update remote refs
git fetch --all --prune | Out-Null

# Get remote-tracking branches (avoid parsing refs/remotes/* patterns that can misbehave on some setups)
$refs = @(git branch -r 2>$null |
  ForEach-Object { $_.Trim() } |
  Where-Object { $_ -and ($_ -notmatch '\s->\s') })  # remove lines like: origin/HEAD -> origin/main

if (-not $AllRemotes) {
  $refs = $refs | Where-Object { $_ -like "$Remote/*" }
}

if (-not $refs -or $refs.Count -eq 0) {
  if ($AllRemotes) { throw "No remote-tracking branches found on any remote." }
  throw "No remote-tracking branches found for remote '$Remote'. Try: .\latest.branch.ps1 -AllRemotes"
}

# Build list of (date, sha, subject) for each remote branch tip
$items = foreach ($r in $refs) {
  $line = git log -1 --format="%cI|%H|%s" $r 2>$null
  if (-not $line) { continue }
  $p = $line -split '\|', 3
  if ($p.Count -ne 3) { continue }

  [pscustomobject]@{
    RemoteRef  = $r
    CommitDate = $p[0]
    Sha        = $p[1]
    Subject    = $p[2]
  }
}

if (-not $items -or $items.Count -eq 0) {
  throw "Could not read commit info for remote branches."
}

# Pick most recent by commit date
$latest = $items |
  Sort-Object @{ Expression = { [DateTimeOffset]::Parse($_.CommitDate) }; Descending = $true } |
  Select-Object -First 1

$selectedRemote = ($latest.RemoteRef -split '/', 2)[0]   # e.g. origin from origin/feature/x
$sha = $latest.Sha

# Resolve branch name(s) using the SHA (branches that point exactly at that SHA)
$pointsAt = @(git for-each-ref --points-at=$sha "refs/remotes/$selectedRemote" --format="%(refname:short)" 2>$null |
  Where-Object { $_ -and ($_ -notlike "$selectedRemote/HEAD") })

# Optional fallback: branches that contain the SHA (not necessarily tip)
if (($pointsAt.Count -eq 0) -and $ContainsFallback) {
  $pointsAt = @(git branch -r --contains $sha 2>$null |
    ForEach-Object { $_.Trim() } |
    Where-Object { $_ -like "$selectedRemote/*" -and $_ -notmatch '\s->\s' -and $_ -ne "$selectedRemote/HEAD" })
}

$branchNames = $pointsAt |
  ForEach-Object { $_ -replace "^$([regex]::Escape($selectedRemote))/", "" } |
  Sort-Object -Unique

if (-not $branchNames -or $branchNames.Count -eq 0) { $branchNames = @("<unknown>") }

"Latest branch name(s): $($branchNames -join ', ')"     # e.g. feature/v1.5.1
"Remote: $selectedRemote"
"Latest SHA: $sha"
"Commit date: $($latest.CommitDate)"
"Subject: $($latest.Subject)"
"Chosen from tip of: $($latest.RemoteRef)"

if ($Top -gt 0) {
  ""
  "Top $Top most recently updated remote branches on ${selectedRemote}:"
  $items |
    Sort-Object @{ Expression = { [DateTimeOffset]::Parse($_.CommitDate) }; Descending = $true } |
    Select-Object -First $Top |
    ForEach-Object {
      [pscustomobject]@{
        CommitDate = $_.CommitDate
        RemoteRef  = $_.RemoteRef
        Sha        = $_.Sha.Substring(0, 7)
        Subject    = $_.Subject
      }
    } |
    Format-Table -AutoSize
}