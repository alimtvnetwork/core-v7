# Unit Coverage Fix & Test Migration Protocol

## Trigger: "fix unit test" or "Unit Coverage Fix"

## 1. Objectives
1. Fix build issues, runtime failures, blocked packages, and failing tests first
2. Move all tests from inside packages to `/tests/integratedtests/<package-name>tests/`
3. Fix all assertion, formatting, and structural violations
4. Achieve 100% code coverage across all packages
5. Enhance testing guidelines where gaps exist

## 2. Prerequisites
- The `.out` file and related JSON coverage files must be provided before starting. If not given, **ask for them**.
- If build issues, blocked packages, or failing tests exist, **ask for related files** before attempting fixes.
- Read the testing guideline spec before writing any tests.
- Read CMP test examples as the reference standard.
- You do not need to instruct the user to run `-TC`. The user handles that.

## 3. Execution Priority
**Always fix in this order:**
1. **Build issues** — fix compilation errors first
2. **Blocked packages** — resolve dependencies, runtime errors
3. **Failing tests** — fix existing broken tests
4. Then proceed to migration, refactoring, and coverage

For each, if relevant files are not provided, **ask the user for them**.

## 4. Target Folder Structure
```
/tests/
  integrationtests/
    <package-name>tests/
      <function>_test.go
```
No `*_test.go` files inside package directories. All tests under `/tests/integratedtests/`.

## 5. Skip Rules
- **Internal packages**: Skip entirely.
- **Private methods**: Discuss with user whether to skip or test indirectly. Do not assume.

## 6. Mandatory Test Rules

### 6.1 Assertion Style
- **NEVER** use `t.Error`, `t.Fail`, or `t.Fatalf`
- **ALWAYS** use GoConvey / `Should` style assertions

### 6.2 AAA Pattern (Strictly Enforced)
```go
// Arrange
// ... setup inputs, mocks, expected values

// Act
// ... call the function under test

// Assert
// ... So(result, ShouldEqual, expected)
```

### 6.3 Map Formatting
Each key-value pair on a separate line. Never inline.

### 6.4 Coverage Requirements
Tests must cover: normal paths, edge cases, error handling, boundary conditions, all branches.
Tests must be deterministic, non-flaky, with full branch coverage.

## 7. Execution Phases
1. **Build & Blocker Fix** — compilation errors, blocked packages, failing tests
2. **Audit & Discovery** — identify violations, create mapping, document in plan
3. **Test Migration** — move tests to integration folder, update imports
4. **Test Refactoring** — replace t.Error, apply AAA, fix map formatting
5. **Fix Existing Integration Tests** — same fixes for existing tests
6. **Coverage Gap Closure** — exactly two packages per iteration
7. **Guideline Enhancement** — update testing spec with missing rules

## 8. Iteration Flow (triggered by "next")
Select next two packages → migration → refactoring → coverage → validation → report completed/remaining.

## 9. Blocked Package Resolution
1. Identify: failing tests, missing deps, runtime errors
2. Root cause: stack traces, divide-and-conquer
3. Fix: resolve conditions, mock/stub deps
4. Do NOT modify production logic unless strictly required

## 10. Version Rule
Any code change bumps at least minor version. Never modify `.release` folder.

## 11. Acceptance Criteria
1. Zero test files inside package directories
2. All tests in `/tests/integratedtests/<package-name>tests/`
3. All GoConvey/Should assertions (zero t.Error)
4. All AAA pattern with explicit comments
5. All maps line-by-line
6. All packages 100% coverage
7. No failing/flaky tests
8. Blocked packages resolved with documented root cause
9. Testing guidelines updated
10. Each iteration processes exactly two packages
11. Remaining packages listed after each cycle
12. Internal packages skipped
13. Private method coverage discussed with user
