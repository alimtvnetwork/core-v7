# Contributing

## Authors

- Md. Alim Ul Karim — [LinkedIn](https://www.linkedin.com/in/alimkarim)

---

## Branching Strategy

This project follows a **trunk-based** workflow with release branches.

| Branch | Purpose |
|---|---|
| `main` | Stable, production-ready code |
| `develop` | Integration branch for ongoing work (if used) |
| `feature/<name>` | New features — branched from `main` |
| `bugfix/<name>` | Bug fixes — branched from `main` |
| `hotfix/<name>` | Urgent production fixes — branched from `main`, merged back to `main` |
| `release/vX.Y.Z` | Release candidates — triggers the automated release pipeline |

### Branch Naming

Use lowercase kebab-case with a category prefix:

```
feature/add-retry-logic
bugfix/fix-nil-pointer
hotfix/patch-auth-bypass
release/v1.6.0
```

---

## PR Workflow

### 1. Create a Branch

```bash
git checkout main
git pull origin main
git checkout -b feature/my-feature
```

### 2. Make Changes

- Write clear, focused commits (one logical change per commit).
- Include or update tests for any new/changed behavior.
- Run linting and tests locally before pushing:

```bash
golangci-lint run
go test -race ./...
```

### 3. Open a Pull Request

- Target branch: `main` (or `release/*` for release prep).
- Fill in the PR template (if available).
- Add a clear title and description explaining **what** and **why**.
- Link related issues (e.g., `Closes #42`).

### 4. CI Checks

Every PR must pass **all** required status checks before merging:

| Check | What it does |
|---|---|
| **Lint** | `golangci-lint` static analysis |
| **Test & Coverage** | `go test -race` with ≥ 60% coverage threshold |
| **Security Scan** | `govulncheck` for known vulnerabilities |
| **Build** | Ensures the project compiles cleanly |

### 5. Code Review

- At least **1 approving review** is required.
- Address all review comments; re-request review after changes.
- Stale approvals are automatically dismissed when new commits are pushed.

### 6. Merge

- Use **squash merge** for feature/bugfix branches to keep history clean.
- Use **merge commit** for release branches to preserve the full history.
- Delete the source branch after merging.

---

## Release Process

1. Create a branch named `release/vX.Y.Z` from `main`.
2. Push the branch — CI runs all checks.
3. Once CI passes, the release pipeline automatically:
   - Creates a Git tag (`vX.Y.Z`)
   - Builds release binaries via GoReleaser
   - Publishes a GitHub Release
   - Sends a Slack/Discord notification

---

## Code Standards

- Follow idiomatic Go conventions (`gofmt`, `go vet`).
- All exported symbols must have doc comments.
- Minimum **60% test coverage** enforced by CI.
- No unresolved `govulncheck` findings in dependencies.

---

## Getting Help

Open an issue or start a discussion if you have questions about contributing.
