# `core` — Go Utility Framework

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go&logoColor=white)](https://go.dev/doc/install)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Pipeline Status](https://github.com/alimtvnetwork/core/badges/main/pipeline.svg)](https://github.com/alimtvnetwork/core/-/pipelines)
[![Coverage](https://github.com/alimtvnetwork/core/badges/main/coverage.svg)](https://github.com/alimtvnetwork/core/-/pipelines)
[![Go Report Card](https://goreportcard.com/badge/github.com/alimtvnetwork/core)](https://goreportcard.com/report/github.com/alimtvnetwork/core)

![Core logo](assets/core-250.png)

> **The power of Java/.NET — the simplicity of Go.**

## Why This Exists

Go is beautifully simple — but that simplicity comes with trade-offs. **Verbosity** is the first: operations that take one line in C# or Java often take five in Go. The second is **ecosystem fragmentation**: Go's culture of small, single-purpose packages means real-world projects end up stitching together dozens of third-party libraries, each written with a different philosophy, different naming conventions, different error-handling styles, and different testing approaches. The result? Your codebase looks like a patchwork of workarounds — not a cohesive system.

**`core` solves this.** It is the foundational package of the [**auk-go**](https://gitlab.com/auk-go) ecosystem — a series of packages designed to work together as a unified platform. When you build on `core`, every package in your project shares the same structure, the same conventions, and the same developer experience:

- **No more verbosity** — generic ternary helpers ([`conditional/`](/conditional/README.md)), nil-safe pointer conversions ([`typesconv/`](/typesconv/readme.md)), one-line JSON pipelines ([`corejson/`](/coredata/corejson/README.md)), and compute-once caching ([`coreonce/`](/coredata/coreonce/README.md)) eliminate boilerplate without hiding complexity.
- **No more magic strings** — named constants ([`constants/`](/constants/README.md), [`coreindexes/`](/coreindexes/readme.md)), strongly-typed enums ([`issetter/`](/issetter/README.md)), and semantic error types ([`errcore/`](/errcore/README.md)) replace hardcoded values everywhere.
- **No more inconsistent tests** — the built-in testing framework ([`coretests/`](/coretests/README.md)) provides `CaseV1` test cases, `args.Map` for semantic test inputs, `ShouldBeEqual`/`ShouldBeEqualMap` assertions, and the AAA (Arrange-Act-Assert) pattern out of the box. Separate your test data (`_testcases.go`) from test logic (`_test.go`) — every test reads the same way across the entire ecosystem.
- **No more interface guesswork** — 100+ canonical interface contracts ([`coreinterface/`](/coreinterface/README.md)) following Go's `-er` suffix convention (`NameGetter`, `Serializer`, `IsEmptyChecker`) ensure packages depend on behaviors, not concrete types.
- **No more unsafe nil panics** — zero-nil safety is a first-class design principle. Functions return empty slices/maps instead of nil. Pointer-receiver methods include nil guards. Use `IsNull()` / `IsEmpty()` / `IsDefined()` for explicit state checking.

The goal is simple: **your code should look the same, read the same, and test the same — whether it's written by you, your teammate, or someone contributing from the other side of the world.** `core` makes that possible by providing the shared foundation that every package in the ecosystem builds upon.

### Built for Real-World Scale

This library is not an academic exercise. It encapsulates **20 years of professional software engineering experience** by [**Md. Alim Ul Karim**](https://www.linkedin.com/in/alimkarim) — a system architect recognized as one of the top software architects globally. Alim has architected large-scale systems across enterprise, fintech, and distributed platforms, and brings that hard-won knowledge directly into every design decision in this framework. The patterns here — struct-as-namespace, one-file-per-function, interface-first design, `newCreator` factories — aren't arbitrary choices. They're battle-tested conventions refined over two decades of building production systems.

### We Welcome Your Feedback

If something doesn't feel right — if a pattern seems wrong, an API is confusing, or you think there's a better approach — **please tell us.** Open an [issue](https://github.com/alimtvnetwork/core/-/issues), start a discussion, or submit a merge request. We take all feedback seriously, treat criticism as a gift, and are committed to continuously improving this library for the community.

## Quick Start

### Prerequisites

| Tool | Version                   |
| ---- | ------------------------- |
| Go   | **1.24+** (latest stable) |
| Git  | ≥ 2.29                    |

### Install

```bash
go get github.com/alimtvnetwork/core
```

### Clone

```bash
git clone https://github.com/alimtvnetwork/core.git
```

### Build & Test

```bash
make                  # build and run default CLI
make build            # compile binary to build/cli
make run-tests        # run integration tests
make run-server       # start server entrypoint
make run-client       # start client entrypoint
make run-sample       # run sample/demo
```

## What This Framework Provides

| Category             | Packages                | What You Get                                                                     |
| -------------------- | ----------------------- | -------------------------------------------------------------------------------- |
| **Ternary helpers**  | `conditional/`          | Generic `If[T]`, `IfFunc[T]`, `IfSlice[T]` — replaces missing ternary operator   |
| **Data structures**  | `coredata/corestr/`     | `Collection`, `Hashmap`, `Hashset`, `LinkedList`, `SimpleSlice`                  |
| **JSON**             | `coredata/corejson/`    | `Serialize.*`, `Deserialize.*` — full JSON pipeline                              |
| **Payload system**   | `coredata/corepayload/` | `PayloadWrapper`, `Attributes`, `PayloadsCollection` — structured data transport |
| **Task info**        | `coretaskinfo/`         | `Info` — metadata container with name, URLs, examples, secure/plain text modes   |
| **Error building**   | `errcore/`              | Stack-traced errors, merge, expectations, Gherkins-style messages                |
| **File permissions** | `chmodhelper/`          | Parse, verify, and apply chmod on files and directories                          |
| **Interfaces**       | `coreinterface/`        | 100+ canonical interface contracts (`*Getter`, `*Checker`, `*Binder`)            |
| **Converters**       | `converters/`           | Type conversions: strings ↔ bytes, maps, pointers                                |
| **Testing**          | `coretests/`            | Assertion helpers, `FuncWrap`, `CaseV1` for AAA-pattern tests                    |
| **Regex**            | `regexnew/`             | `LazyRegex` — compiles only on first use, with optional locking                  |
| **Validators**       | `corevalidator/`        | Line, slice, text, and range validators                                          |
| **Sorting**          | `coresort/`             | Quick sort for strings and integers                                              |
| **Math**             | `coremath/`             | Min/Max for all numeric types                                                    |
| **Versioning**       | `coreversion/`          | Semantic version data type (major.minor.patch)                                   |
| **Constants**        | `constants/`            | OS line separators, empty values, capacity defaults                              |
| **Generics**         | `core.go`, `generic.go` | `EmptySlicePtr[T]`, `SlicePtrByCapacity[T]`, `EmptyMapPtr[K,V]`                  |

## Design Philosophy

1. **One file per function** — each public function lives in its own `.go` file, named after the function. This keeps files small (~50-200 lines) and makes navigation instant.

2. **Struct-as-namespace** — related operations are grouped on unexported struct types exposed via package-level `var`. This gives you IDE autocompletion trees like `corejson.Serialize.ToString()` or `New.PayloadWrapper.Empty()`.

3. **Interface-first** — all contracts are defined in `coreinterface/` using Go's `-er` suffix convention (e.g., `NameGetter`, `Csver`, `Serializer`). Packages depend on interfaces, not concrete types.

4. **Zero-nil safety** — functions return empty slices/maps instead of nil wherever possible. Pointer-receiver methods include nil guards. Use `IsNull()` / `IsEmpty()` / `IsDefined()` for checking.

5. **Generics where clear** — generic versions (`If[T]`, `EmptySlicePtr[T]`, `TypedErrorFunctionsExecuteResults[T]`) exist alongside backward-compatible type-specific wrappers.

6. **Prefer value receivers** (new code) — read-only methods use value receivers for simplicity. Pointer receivers are reserved for mutation, large structs, or interface satisfaction. See [Coding Guidelines](/spec/01-app/17-coding-guidelines.md).

---

## Examples

### Conditional (Ternary) Helpers

```go
import "github.com/alimtvnetwork/core/conditional"

// Generic (Go 1.22+)
result := conditional.If[int](true, 2, 7)          // 2
name := conditional.If[string](len(s) > 0, s, "default")

// With lazy evaluation — only the chosen branch executes
val := conditional.IfFunc[string](expensive, func() string {
    return computeValue()
}, func() string {
    return "fallback"
})

// Slice ternary
items := conditional.IfSlice[int](hasItems, filled, empty)

// Legacy type-specific (still works, deprecated)
result := conditional.Int(true, 2, 7)   // 2
```

### Generic Slice/Map Factories

```go
import "github.com/alimtvnetwork/core"

ints := core.EmptySlicePtr[int]()            // *[]int (empty, non-nil)
strs := core.SlicePtrByLength[string](10)    // *[]string with len=10
m := core.EmptyMapPtr[string, int]()          // *map[string]int (empty, non-nil)
capped := core.SlicePtrByCapacity[int](100)  // *[]int with cap=100
```

### Payload System (PayloadWrapper & Attributes)

The **payload system** (`coredata/corepayload/`) is the primary data transport mechanism. Use `PayloadWrapper` to carry structured data between components:

```go
import "github.com/alimtvnetwork/core/coredata/corepayload"

// Create an empty payload
payload := corepayload.New.PayloadWrapper.Empty()

// Create with instruction
payload = corepayload.New.PayloadWrapper.UsingInstruction(
    &corepayload.PayloadCreateInstruction{
        Name:       "user-create",
        Identifier: "usr-123",
        EntityType: "User",
        Payloads:   myStruct,  // auto-serialized to JSON bytes
    },
)

// Access data
fmt.Println(payload.PayloadName())       // "user-create"
fmt.Println(payload.IdString())          // "usr-123"
fmt.Println(payload.PayloadEntityType()) // "User"

// Deserialize dynamic payloads
var user User
err := payload.Deserialize(&user)

// Check for errors
if payload.HasError() {
    log.Fatal(payload.Error())
}

// Attributes (key-value pairs, auth info, paging)
attrs := payload.InitializeAttributesOnNull()
attrs.AddOrUpdateString("role", "admin")
value, found := attrs.GetStringKeyValue("role") // "admin", true

// Set authentication info
payload.SetUser(&corepayload.User{Name: "alice"})
fmt.Println(payload.Username()) // "alice"

// Serialize/Deserialize the whole wrapper
jsonBytes, err := payload.Serialize()
restored, err := corepayload.New.PayloadWrapper.Deserialize(jsonBytes)
```

### Task Info (coretaskinfo)

`coretaskinfo.Info` holds metadata about tasks, errors, or operations — name, description, URLs, examples, and security flags:

```go
import "github.com/alimtvnetwork/core/coretaskinfo"

// Create with factory
info := coretaskinfo.New.Info.Default(
    "user-validation",
    "Validates user input fields",
    "https://docs.example.com/validation",
)

// Create with examples
info = coretaskinfo.New.Info.Examples(
    "email-format",
    "RFC 5322 email validation",
    "https://tools.ietf.org/html/rfc5322",
    `validate("user@example.com")`,
    `validate("bad@")`, // fails
)

// Secure vs plain text (controls payload logging)
secureInfo := coretaskinfo.New.Info.Secure.Default(
    "password-reset", "Handles password reset", "",
)
fmt.Println(secureInfo.IsSecure())    // true
fmt.Println(secureInfo.IsPlainText()) // false

// Nil-safe access
var nilInfo *coretaskinfo.Info
fmt.Println(nilInfo.SafeName())        // "" (no panic)
fmt.Println(nilInfo.IsPlainText())     // true (default)

// Serialize to JSON
jsonStr := info.PrettyJsonString()

// Clone
copied := info.Clone()
```

### JSON Serialize / Deserialize

```go
import "github.com/alimtvnetwork/core/coredata/corejson"

type Example struct {
    A       string
    B       int
    SomeMap map[string]string
}

from := &Example{A: "hello", B: 42, SomeMap: map[string]string{}}
to := &Example{}

// Deep copy via JSON
err := corejson.Deserialize.FromTo(from, to)

// Serialize to string
jsonStr, err := corejson.Serialize.ToString(from)
// jsonStr = `{"A":"hello","B":42,"SomeMap":{}}`

// Serialize to bytes
jsonBytes, err := corejson.Serialize.Raw(from)

// Deserialize from bytes
err = corejson.Deserialize.UsingBytes(jsonBytes, to)

// Pretty print
pretty := corejson.NewPtr(from).PrettyJsonString()
```

### String Collections

```go
import (
    "github.com/alimtvnetwork/core/coredata/corestr"
    "github.com/alimtvnetwork/core/constants"
)

values := []string{"hello", "world", "something"}
collection := corestr.NewCollectionPtrUsingStrings(&values, constants.Zero)

fmt.Println(collection.Length())  // 3
fmt.Println(collection.IsEmpty()) // false

collection.AddsLock("else")
fmt.Println(collection.Length())  // 4
```

### Hashset

```go
import "github.com/alimtvnetwork/core/coredata/corestr"

hs := corestr.NewHashset(2)
hs.Add("alpha")
hs.Add("beta")
fmt.Println(hs.Length()) // 2
fmt.Println(hs.Has("alpha")) // true
```

### Error Construction

```go
import "github.com/alimtvnetwork/core/errcore"

// Rich error with stack trace
err := errcore.Expected.Error("config file", "/etc/app.conf")

// Merge multiple errors
combined := errcore.MergeErrors(err1, err2)

// Slice of error strings to single error
sliceErr := errcore.SliceToError([]string{"issue 1", "issue 2"})

// Type-based errors
err = errcore.CannotBeNilOrEmptyType.ErrorNoRefs("user input")
```

### Regex (Lazy Compiled) — `regexnew/`

The `regexnew` package provides thread-safe, lazy-compiled regular expressions with caching. Patterns compile only on first use and are stored in a global map for reuse.

#### Creating Regex — Lock vs No-Lock

```go
import "github.com/alimtvnetwork/core/regexnew"

// --- Package-level vars (init time, no goroutine contention) ---
// Use New.Lazy — no mutex lock needed at init
var digitRegex = regexnew.New.Lazy(`\d+`)
var emailRegex = regexnew.New.Lazy(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

// --- Inside methods/goroutines (concurrent access) ---
// Use New.LazyLock — mutex-protected for thread safety
func validateInput(input string) bool {
    lazy := regexnew.New.LazyLock(`^[a-z]+\d+$`)
    return lazy.IsMatch(input)
}
```

#### Matching & Validation

```go
lazy := regexnew.New.LazyLock(`\d+`)

// Simple match
lazy.IsMatch("abc123")          // true
lazy.IsMatch("no digits")       // false
lazy.IsFailedMatch("abc123")    // false (inverse of IsMatch)

// Byte matching
lazy.IsMatchBytes([]byte("abc123"))        // true
lazy.IsFailedMatchBytes([]byte("letters")) // true

// First submatch extraction
lazy2 := regexnew.New.LazyLock(`(\d+)`)
match, invalid := lazy2.FirstMatchLine("abc 123 def 456")
// match = "123", invalid = false

// Match with error return (for validation pipelines)
err := lazy.MatchError("abc123") // nil (matched)
err = lazy.MatchError("letters") // error with descriptive message
```

#### Compile Control

```go
lazy := regexnew.New.LazyLock(`\d+`)

// Explicit compile (returns cached after first call)
regex, err := lazy.Compile()

// Panic on compile failure
regex = lazy.CompileMust()

// State inspection
lazy.IsCompiled()   // true after first Compile/Match
lazy.IsApplicable() // true if pattern is valid and compiled
lazy.IsDefined()    // true if pattern is non-empty with compiler
lazy.HasError()     // true if compile failed
lazy.HasAnyIssues() // true if nil, undefined, or compile error
lazy.Pattern()      // returns the raw pattern string
lazy.String()       // same as Pattern()
```

#### Package-Level Functions (Lock = thread-safe)

```go
// Direct create (cached in global map)
regex, err := regexnew.Create(`\d+`)          // no lock, use in var init
regex, err = regexnew.CreateLock(`\d+`)        // with lock, use in methods
regex = regexnew.CreateMust(`\d+`)             // panics on error

// Conditional locking
regex, err = regexnew.CreateLockIf(true, `\d+`)
regex = regexnew.CreateMustLockIf(true, `\d+`)

// Create with applicability check
regex, err, isApplicable := regexnew.CreateApplicableLock(`\d+`)

// Quick match (creates regex + matches in one call)
regexnew.IsMatchLock(`\d+`, "abc123")   // true
regexnew.IsMatchFailed(`\d+`, "abc123") // false

// Match with error (for validation)
err = regexnew.MatchError(`^hello$`, "hello")     // nil
err = regexnew.MatchErrorLock(`^hello$`, "world")  // error

// Custom match function
err = regexnew.MatchUsingFuncErrorLock(
    `\d+`, "abc123",
    func(re *regexp.Regexp, term string) bool {
        return re.MatchString(term)
    },
)

// Custom match + custom error
err = regexnew.MatchUsingCustomizeErrorFuncLock(
    `\d+`, "abc123",
    func(re *regexp.Regexp, term string) bool {
        return re.MatchString(term)
    },
    func(pattern, term string, err error, re *regexp.Regexp) error {
        return fmt.Errorf("custom: %s failed on %s", pattern, term)
    },
)
```

#### Multi-Pattern Creation

```go
// Create two patterns atomically (single lock)
first, second := regexnew.New.LazyRegex.TwoLock(`\d+`, `[a-z]+`)
first.IsMatch("42")    // true
second.IsMatch("abc")  // true

// Create many patterns atomically
patterns := regexnew.New.LazyRegex.ManyUsingLock(`\d+`, `[a-z]+`, `\w+`)
patterns[`\d+`].IsMatch("123") // true

// Get all cached patterns
allPatterns := regexnew.New.LazyRegex.AllPatternsMap()

// Conditional lock
lazy := regexnew.New.LazyRegex.NewLockIf(true, `\d+`)
```

#### Pre-compiled Regex Constants

```go
import "github.com/alimtvnetwork/core/regexnew"

// Ready-to-use lazy regex for common patterns
regexnew.WhitespaceFinderRegex.IsMatch("hello world")          // true
regexnew.DollarIdentifierRegex.IsMatch("$MY_VAR")              // true
regexnew.PrettyNameRegex.IsMatch("John Doe")                   // true
regexnew.FirstNumberAnyWhereCheckerRegex.IsMatch("version 42") // true
```

### Sorting

```go
import "github.com/alimtvnetwork/core/coresort/strsort"

fruits := []string{"banana", "mango", "apple"}
strsort.Quick(&fruits)    // [apple banana mango]
strsort.QuickDsc(&fruits) // [mango banana apple]
```

### 6-Valued Boolean (`issetter`)

The [`issetter`](/issetter/) package provides a byte-backed boolean type with **6 states** for lazy evaluation, deferred decisions, and wildcard matching — going beyond Go's native `true`/`false`.

| Constant        | Byte Value | Meaning             | Use Case                     |
| --------------- | ---------- | ------------------- | ---------------------------- |
| `Uninitialized` | `0`        | Not yet evaluated   | Lazy fields, deferred config |
| `True`          | `1`        | Positive / yes / on | Standard boolean true        |
| `False`         | `2`        | Negative / no / off | Standard boolean false       |
| `Unset`         | `3`        | Explicitly cleared  | User removed a setting       |
| `Set`           | `4`        | Explicitly assigned | User confirmed a setting     |
| `Wildcard`      | `5`        | Match anything      | Filters, search patterns     |

**Logical grouping:**

- **Positive** (`IsOn`/`IsAccept`/`IsSuccess`): `True`, `Set`
- **Negative** (`IsOff`/`IsReject`/`IsFailed`): `False`, `Unset`
- **Indeterminate** (`IsAsk`/`IsSkip`): `Uninitialized`, `Wildcard`

```go
import "github.com/alimtvnetwork/core/issetter"

// Basic state checks
status := issetter.True
fmt.Println(status.HasInitialized()) // true
fmt.Println(status.IsOn())           // true  (True or Set)
fmt.Println(status.IsOff())          // false

// Uninitialized = "not yet decided"
pending := issetter.Uninitialized
fmt.Println(pending.HasInitialized())    // false
fmt.Println(pending.IsUndefinedLogically()) // true

// Wildcard = "accept anything"
filter := issetter.Wildcard
fmt.Println(filter.WildcardApply(false)) // false (passes through input)
fmt.Println(filter.IsWildcardOrBool(true)) // true (wildcard always true)

// Lazy evaluation — only runs func on first call
var lazyFlag issetter.Value // starts as Uninitialized (0)
called := lazyFlag.LazyEvaluateBool(func() {
    fmt.Println("computed!")  // prints only once
})
fmt.Println(called)          // true
fmt.Println(lazyFlag.IsTrue()) // true

// Convert between boolean ↔ set/unset semantics
v := issetter.True
fmt.Println(v.ToSetUnsetValue()) // Set
s := issetter.Set
fmt.Println(s.ToBooleanValue())  // True

// Ternary byte mapping
result := issetter.True.ToByteCondition(1, 0, 255)
fmt.Println(result) // 1
```

See [`issetter/Value.go`](/issetter/Value.go) for all methods and [`issetter/README.md`](/issetter/README.md) for the full API reference.

### File Permissions (chmodhelper)

```go
import "github.com/alimtvnetwork/core/chmodhelper"

// Parse rwx string
rwx := chmodhelper.ExpandCharRwx("rwxr-xr--")
// Verify file permissions
isValid := chmodhelper.ChmodVerify.IsFileHasRwx(path, expectedRwx)
```

### CSV Formatting

```go
import "github.com/alimtvnetwork/core/corecsv"

// Any type implementing Csver interface gets CSV output
line := corecsv.DefaultCsv(myStruct) // "field1,field2,field3"
```

### Generic Typed Functions Execution

```go
import "github.com/alimtvnetwork/core/conditional"

// Execute a set of functions and collect results
results, err := conditional.TypedErrorFunctionsExecuteResults[string](
    true,
    []func() (string, error){
        func() (string, error) { return "hello", nil },
        func() (string, error) { return "world", nil },
    },
    nil, // false-branch functions (not used)
)
// results = ["hello", "world"], err = nil
```

---

## Unit Test Pattern

This project follows the **Arrange-Act-Assert (AAA)** pattern with `coretestcases.CaseV1`, GoConvey assertions, and `errcore.AssertDiffOnMismatch`.

> 📖 **[Full Testing Guidelines](/spec/01-app/16-testing-guidelines.md)** — comprehensive reference covering all assertion methods, `args.Map` usage, named test case variables, `SliceValidator`, comparison modes (Equal/Contains/Regex/Sorted), custom test wrappers, panic testing, concurrency testing, anti-patterns, and a CaseV2 proposal.

### Quick Example

```go
import (
    "testing"
    "github.com/alimtvnetwork/core/coretests/coretestcases"
    "github.com/alimtvnetwork/core/coretests/args"
    "github.com/alimtvnetwork/core/coretests"
)

// In _testcases.go file:
var myTestCases = []coretestcases.CaseV1{
    {
        Title: "valid input returns expected output",
        ArrangeInput: args.Map{
            "when":   "given valid input",
            "actual": "hello",
            "expect": "HELLO",
        },
        ExpectedInput: []string{"HELLO"},
    },
}

// In _test.go file:
func Test_MyFunction_Verification(t *testing.T) {
    for caseIndex, testCase := range myTestCases {
        // Arrange
        input := testCase.ArrangeInput.(args.Map)

        // Act
        result := strings.ToUpper(input.Actual().(string))

        // Assert
        actualLines := coretests.GetAssert.ToStrings(result)
        testCase.ShouldBeEqual(t, caseIndex, actualLines...)
    }
}
```

### Key Principles

1. **Separate test data** — `_testcases.go` files keep data out of test logic.
2. **AAA comments** — always label `// Arrange`, `// Act`, `// Assert` sections.
3. **Index tracking** — always pass `caseIndex` for precise failure identification.
4. **String-line comparison** — convert results to `[]string` for readable diffs.

---

## Method Writing: Split Boolean-Flag Methods

When a method's behavior changes based on a `bool` parameter, **create two separate methods** with expressive names instead of one method with a flag. The caller's code reads like documentation.

### Lock vs No-Lock

```go
// Add appends without locking — caller manages concurrency.
func (it *Collection) Add(str string) *Collection { ... }

// AddLock appends with mutex protection (thread-safe).
func (it *Collection) AddLock(str string) *Collection {
    it.Lock()
    defer it.Unlock()
    return it.Add(str)
}
```

### Conditional Execution (`*If`)

```go
// FmtDebug always logs.
func FmtDebug(format string, items ...any) { slog.Debug(fmt.Sprintf(format, items...)) }

// FmtDebugIf logs only when isDebug is true.
func FmtDebugIf(isDebug bool, format string, items ...any) {
    if !isDebug { return }
    FmtDebug(format, items...)
}
```

### Behavioral Pairs

```go
func (it Status) IsValid() bool   { return it != Invalid }
func (it Status) IsInvalid() bool { return it == Invalid }
```

| Suffix | Pattern | Example |
|--------|---------|---------|
| `*Lock` | Thread-safe variant | `Add` → `AddLock` |
| `*If` | Conditional execution | `FmtDebug` → `FmtDebugIf` |
| `*LockIf` | Conditional locking | `Create` → `CreateLockIf` |
| (pair) | Opposite states | `IsValid` + `IsInvalid` |
| `*NonEmpty` | Skip empty strings | `Add` → `AddNonEmpty` |
| `*NonEmptyWhitespace` | Skip empty + whitespace | `Add` → `AddNonEmptyWhitespace` |
| `*Join` | Filter then join | `NonEmptyJoin`, `NonWhitespaceJoin` |

### Filtering Variants (`*NonEmpty`, `*NonEmptyWhitespace`)

String methods provide **filtering variants** that silently skip invalid items:

```go
// AddNonEmpty skips "" only
c.AddNonEmpty("")    // skipped
c.AddNonEmpty("a")   // added

// AddNonEmptyWhitespace skips "" and whitespace-only
c.AddNonEmptyWhitespace("   ")  // skipped
c.AddNonEmptyWhitespace("a")    // added

// Standalone slice functions
filtered := stringslice.NonEmptyStrings(input)   // removes ""
filtered := stringslice.NonWhitespace(input)      // removes "" and whitespace
filtered := stringslice.TrimmedEachWords(input)    // trims + removes empty

// Conditional dispatch
result := stringslice.NonEmptyIf(shouldFilter, input)

// Filter + join
joined := stringslice.NonEmptyJoin(input, ", ")
joined := stringslice.NonWhitespaceJoin(input, "\n")
```

See **[Coding Guidelines — Method Writing](/spec/01-app/17-coding-guidelines.md#method-writing-split-boolean-flag-methods-into-expressive-pairs)** for full details and all patterns.

### Pointer Variants (`*Ptr` Suffix)

Methods provide `*Ptr` variants for pointer return types and nil-safe pointer acceptance:

```go
// Return pointer variant
func (it Version) Json() corejson.Result    { return corejson.New(it) }
func (it Version) JsonPtr() *corejson.Result { return corejson.NewPtr(it) }

// Nil-safe checker variant
func IsEmpty(str string) bool        { return str == "" }
func IsEmptyPtr(str *string) bool    { return str == nil || *str == "" }

// Identity conversion
func (it Variant) ToPtr() *Variant   { return &it }
func (it Version) NonPtr() Version   { return it }
```

**Rules**: (1) `*Ptr` checkers treat `nil` as empty/absent. (2) Pointer-receiver `*Ptr` methods must guard `nil`. (3) Each variant in its own file (`IsEmpty.go` / `IsEmptyPtr.go`).

See **[Coding Guidelines — Pointer Variants](/spec/01-app/17-coding-guidelines.md#method-writing-pointer-variants-ptr-suffix)** for all five patterns.

---

## Interface Conventions

All interfaces in `coreinterface/` follow Go's `-er` suffix convention:

| Pattern    | Example                               | Purpose                     |
| ---------- | ------------------------------------- | --------------------------- |
| `*Getter`  | `NameGetter`, `ValueGetter`           | Read a single value         |
| `*Checker` | `HasErrorChecker`, `IsEmptyChecker`   | Boolean state check         |
| `*Binder`  | `ContractsBinder`, `AttributesBinder` | Compose multiple interfaces |
| `*er`      | `Csver`, `Serializer`, `Stringer`     | Action performer            |

---

## Troubleshooting

| Problem                                 | Solution                                                                                                                     |
| --------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `go get` fails with auth error          | Add SSH key to GitLab or use access token: `git config url."https://oauth2:TOKEN@gitlab.com".insteadOf "https://gitlab.com"` |
| `go mod tidy` reports version conflicts | Ensure `go.mod` specifies `go 1.24` and run `go mod tidy`                                                                    |
| Tests fail after clone                  | Run `make run-tests` — some tests require the full module graph                                                              |
| Import path has typo                    | Known: `convertinteranl` → `convertinternal`, `refeflectcore` → `reflectcore` (being fixed)                                  |

## Project Structure

```
core/
├── core.go, generic.go          # Root package — generic slice/map factories
├── makefile                     # Build, test, run targets
├── go.mod / go.sum              # Go module definition (Go 1.24+)
│
├── conditional/                 # Generic ternary helpers (If[T], IfFunc[T], IfSlice[T])
├── constants/                   # OS line separators, empty values, capacity defaults
├── converters/                  # Type conversions: strings ↔ bytes, maps, pointers
│
├── coredata/                    # Data structures & serialization
│   ├── coreapi/                 #   API request/response types (Generic + Typed[T])
│   ├── coredynamic/             #   Dynamic type wrappers, Collection[T], SimpleRequest
│   ├── coregeneric/             #   Generic Collection, Hashset, Hashmap, LinkedList
│   ├── corejson/                #   JSON serialize/deserialize pipeline
│   ├── corepayload/             #   PayloadWrapper — structured data transport
│   ├── corestr/                 #   String Collection, Hashset, Hashmap, ValidValue
│   ├── coreonce/                #   Lazy-evaluated cached values (StringOnce, IntegerOnce)
│   ├── corerange/               #   Range types (int, byte)
│   └── stringslice/             #   Slice utilities for []string
│
├── coreinterface/               # 100+ canonical interface contracts
│   ├── enuminf/                 #   Enum interfaces
│   ├── errcoreinf/              #   Error wrapper interfaces
│   ├── serializerinf/           #   Serializer/deserializer interfaces
│   └── baseactioninf/           #   Action/execution interfaces
│
├── corefuncs/                   # Function type definitions (generic + legacy)
├── coretaskinfo/                # Task metadata (Info, ExcludingOptions)
├── corevalidator/               # Line, slice, text, range validators
├── coreversion/                 # Semantic versioning (major.minor.patch)
├── coremath/                    # Min/Max for all numeric types
├── coresort/                    # Quick sort for strings and integers
├── corecsv/                     # CSV formatting utilities
│
├── errcore/                     # Rich error construction with stack traces
├── chmodhelper/                  # File permission parsing and verification
├── regexnew/                    # Lazy-compiled regex with thread-safe caching
├── issetter/                    # 6-valued boolean (Uninitialized/True/False/Unset/Set/Wildcard)
│
├── coretests/                   # Testing helpers, FuncWrap, assertion wrappers
│   ├── args/                    #   FuncWrap argument types (OneFunc, TwoFunc, etc.)
│   └── coretestcases/           #   CaseV1 test case definitions
│
├── tests/integratedtests/       # All unit/integration tests (per-package subdirs)
│
├── internal/                    # Internal packages (not importable externally)
│   ├── convertinternal/         #   Low-level type conversion
│   ├── reflectinternal/         #   Reflection helpers
│   └── strutilinternal/         #   String utility internals
│
├── # codegen/ removed (v1.6.0) — was test boilerplate generation
├── cmd/                         # CLI entrypoints (main, server, client, sample)
│
├── spec/                        # Architecture docs, coding guidelines, issue tracking
│   ├── 01-app/                  #   Core specs and conventions
│   └── 13-app-issues/           #   Known issues and improvement backlog
│
└── assets/                      # Logo and static assets
```

### Package READMEs

Each major package has its own README with detailed type hierarchies, usage examples, and method references:

#### Data Structures & Serialization (`coredata/`)

| Package                 | README                                                                                                      |
| ----------------------- | ----------------------------------------------------------------------------------------------------------- |
| `coredata/`             | [`README.md`](/coredata/README.md) — Umbrella index for all data sub-packages                               |
| `coredata/coregeneric/` | [`README.md`](/coredata/coregeneric/README.md) — Generic Collection, Hashset, Hashmap, LinkedList           |
| `coredata/corestr/`     | [`README.md`](/coredata/corestr/README.md) — String Collection, Hashmap, Hashset, ValidValue                |
| `coredata/corejson/`    | [`README.md`](/coredata/corejson/README.md) — JSON serialization pipeline                                   |
| `coredata/coredynamic/` | [`README.md`](/coredata/coredynamic/README.md) — Dynamic wrappers, Collection[T], TypedSimpleRequest/Result |
| `coredata/corepayload/` | [`README.md`](/coredata/corepayload/README.md) — PayloadWrapper, TypedPayloadWrapper[T], Attributes         |
| `coredata/coreapi/`     | [`README.md`](/coredata/coreapi/README.md) — Typed API request/response                                     |
| `coredata/coreonce/`    | [`README.md`](/coredata/coreonce/README.md) — Lazy-evaluated cached values                                  |
| `coredata/corerange/`   | [`README.md`](/coredata/corerange/README.md) — Range types (int, byte)                                      |
| `coredata/stringslice/` | [`README.md`](/coredata/stringslice/README.md) — Slice utilities for []string                               |

#### Interfaces & Contracts (`coreinterface/`)

| Package                        | README                                                                                 |
| ------------------------------ | -------------------------------------------------------------------------------------- |
| `coreinterface/`               | [`README.md`](/coreinterface/README.md) — Shared interface contracts                   |
| `coreinterface/enuminf/`       | [`README.md`](/coreinterface/enuminf/README.md) — Enum interface contracts             |
| `coreinterface/errcoreinf/`    | [`README.md`](/coreinterface/errcoreinf/README.md) — Error core interface contracts    |
| `coreinterface/loggerinf/`     | [`README.md`](/coreinterface/loggerinf/README.md) — Logger interface contracts         |
| `coreinterface/serializerinf/` | [`README.md`](/coreinterface/serializerinf/README.md) — Serializer interface contracts |

#### Implementations (`coreimpl/`)

| Package     | README                                                               |
| ----------- | -------------------------------------------------------------------- |
| `coreimpl/` | [`README.md`](/coreimpl/README.md) — Core implementations (enumimpl) |

#### Utilities & Helpers

| Package          | README                                                                     |
| ---------------- | -------------------------------------------------------------------------- |
| `conditional/`   | [`README.md`](/conditional/README.md) — Generic ternary & nil-safe helpers |
| `constants/`     | [`README.md`](/constants/README.md) — Shared constants & capacity values   |
| `converters/`    | [`README.md`](/converters/README.md) — Type conversion utilities           |
| `corefuncs/`     | [`README.md`](/corefuncs/README.md) — Function type definitions & wrappers |
| `coretaskinfo/`  | [`README.md`](/coretaskinfo/README.md) — Task metadata container           |
| `corevalidator/` | [`README.md`](/corevalidator/README.md) — Text, line & slice validators    |
| `coremath/`      | [`README.md`](/coremath/README.md) — Min/Max for all numeric types         |
| `coresort/`      | [`README.md`](/coresort/README.md) — Quick sort for strings and integers   |

#### Error Handling & Comparison

| Package    | README                                                                  |
| ---------- | ----------------------------------------------------------------------- |
| `errcore/` | [`README.md`](/errcore/README.md) — Error construction & formatting     |
| `anycmp/`  | [`README.md`](/anycmp/README.md) — Any-type quick comparison            |
| `isany/`   | [`README.md`](/isany/README.md) — Reflection-based type & null checking |

#### System & I/O

| Package        | README                                                             |
| -------------- | ------------------------------------------------------------------ |
| `chmodhelper/` | [`README.md`](/chmodhelper/README.md) — File permission management |
| `regexnew/`    | [`README.md`](/regexnew/README.md) — Lazy-compiled regex           |
| `issetter/`    | [`README.md`](/issetter/README.md) — Multi-valued boolean enum     |

#### Testing

| Package      | README                                                            |
| ------------ | ----------------------------------------------------------------- |
| `coretests/` | [`README.md`](/coretests/README.md) — Test utilities & assertions |

#### CLI

| Package | README                                          |
| ------- | ----------------------------------------------- |
| `cmd/`  | [`README.md`](/cmd/README.md) — CLI entrypoints |

For the complete folder-by-folder breakdown, see the [Folder Map](/spec/01-app/01-folder-map.md).

---

## Core Funcs — Function Type Definitions & Wrappers

The `corefuncs/` package defines reusable function signatures and structural wrappers — generic `[T]` types first, with `any`-based legacy types for backward compatibility. See the **[full corefuncs README](/corefuncs/README.md)** for complete documentation.

```go
import "github.com/alimtvnetwork/core/corefuncs"

// Generic function types (type-safe, recommended)
var transform corefuncs.InOutFuncOf[string, int] = func(s string) int {
    return len(s)
}

// Generic wrappers — named error reporting + ActionFunc conversion
wrapper := corefuncs.NewInOutErrWrapper[string, int](
    "parse-age",
    func(s string) (int, error) { return strconv.Atoi(s) },
)

age, err := wrapper.Exec("25")           // strongly typed
errFn := wrapper.AsActionReturnsErrorFunc("25") // convert to ActionReturnsErrorFunc
legacy := wrapper.ToLegacy()              // backward compatible InOutErrFuncWrapper

// ResultDelegatingFuncWrapperOf[T] — typed unmarshal/reflect targets
unmarshalWrapper := corefuncs.NewResultDelegatingWrapper[*User](
    "json-unmarshal",
    func(target *User) error { return json.Unmarshal(data, target) },
)

// Legacy wrappers via New creator
actionWrapper := corefuncs.New.ActionErr("cleanup", cleanupFunc)
successWrapper := corefuncs.New.IsSuccess("healthcheck", pingFunc)
```

---

## Core API — Typed Request/Response

The `coredata/coreapi/` package provides both dynamic (`any`-based) and strongly-typed (`[T]`) API types:

```go
import "github.com/alimtvnetwork/core/coredata/coreapi"

// --- Typed (Generic) API — compile-time type safety ---

// Strongly typed request
type UserCreateInput struct {
    Name  string
    Email string
}

req := coreapi.NewTypedRequestIn[UserCreateInput](
    &coreapi.RequestAttribute{
        Url:          "/api/users",
        ResourceName: "User",
        ActionName:   "Create",
        IsValid:      true,
    },
    UserCreateInput{Name: "Alice", Email: "alice@example.com"},
)

// Access is fully typed — no assertions needed
fmt.Println(req.Request.Name)  // "Alice"
fmt.Println(req.Request.Email) // "alice@example.com"

// Strongly typed response
type UserOutput struct {
    ID   int
    Name string
}

resp := coreapi.NewTypedResponse[UserOutput](
    &coreapi.ResponseAttribute{IsValid: true, HttpCode: 200},
    UserOutput{ID: 1, Name: "Alice"},
)
fmt.Println(resp.Response.ID) // 1

// Clone (deep copy)
clone := req.Clone()

// Convert to legacy dynamic type for backward compatibility
legacyReq := req.ToGenericRequestIn()

// --- Legacy Dynamic API (still supported) ---
dynamicReq := &coreapi.GenericRequestIn{
    Attribute: &coreapi.RequestAttribute{IsValid: true},
    Request:   map[string]string{"key": "value"}, // any type
}
```

---

## JSON — Comprehensive Examples

```go
import "github.com/alimtvnetwork/core/coredata/corejson"

// --- Serialization ---
type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}

user := User{Name: "Alice", Age: 30}

// To JSON string
jsonStr, err := corejson.Serialize.ToString(user)
// `{"name":"Alice","age":30}`

// To JSON bytes
jsonBytes, err := corejson.Serialize.Raw(user)

// Pretty print
result := corejson.NewPtr(user)
pretty := result.PrettyJsonString()

// --- Deserialization ---
var restored User
err = corejson.Deserialize.UsingBytes(jsonBytes, &restored)

// Deep copy via JSON round-trip
source := User{Name: "Bob", Age: 25}
target := User{}
err = corejson.Deserialize.FromTo(source, &target)

// --- Result type (wraps bytes + error) ---
result = corejson.NewPtr(user)
fmt.Println(result.HasError())         // false
fmt.Println(result.HasIssuesOrEmpty()) // false
bytes := result.SafeValues()           // []byte — safe, never nil

// Error handling
invalidResult := corejson.New(make(chan int)) // can't serialize channels
fmt.Println(invalidResult.HasError())         // true
fmt.Println(invalidResult.ErrorString())      // marshaling error message
```

---

## Testing Library — coretests

The `coretests/` package provides assertion helpers and test-case structures for the **AAA pattern**:

```go
import (
    "testing"
    "github.com/alimtvnetwork/core/coretests"
    "github.com/alimtvnetwork/core/coretests/coretestcases"
    "github.com/alimtvnetwork/core/coretests/args"
)

// === Test Cases (in _testcases.go) ===
var uppercaseTestCases = []coretestcases.CaseV1{
    {
        Title: "converts lowercase to uppercase",
        ArrangeInput: args.Map{
            "actual": "hello",
            "expect": "HELLO",
        },
        ExpectedInput: []string{"HELLO"},
    },
    {
        Title: "handles empty string",
        ArrangeInput: args.Map{
            "actual": "",
            "expect": "",
        },
        ExpectedInput: []string{""},
    },
}

// === Test Runner (in _test.go) ===
func Test_ToUpper(t *testing.T) {
    for caseIndex, tc := range uppercaseTestCases {
        // Arrange
        input := tc.ArrangeInput.(args.Map)
        actual := input["actual"].(string)

        // Act
        result := strings.ToUpper(actual)

        // Assert
        lines := coretests.GetAssert.ToStrings(result)
        tc.ShouldBeEqual(t, caseIndex, lines...)
    }
}

// === FuncWrap — reflection-based test wrappers ===
// Wraps a function for automatic input/output assertion
wrap := args.NewOneFunc(
    myFunc,                        // function under test
    "expected output",             // expected result
)
fmt.Println(wrap.WorkFunc)         // the function reference
fmt.Println(wrap.Expect)           // "expected output"

// === GetAs* assertion helpers ===
assert := coretests.GetAssert
lines := assert.ToStrings(result)       // any → []string for comparison
str := assert.ToString(result)          // any → string
```

---

## Specification Docs

Detailed architecture and conventions documentation for AI agents and contributors:

| Document                     | Path                                                                                                                       |
| ---------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| Repository Overview          | [`/spec/01-app/00-repo-overview.md`](/spec/01-app/00-repo-overview.md)                                                     |
| Folder Map                   | [`/spec/01-app/01-folder-map.md`](/spec/01-app/01-folder-map.md)                                                           |
| Per-Folder Specs             | [`/spec/01-app/folders/`](/spec/01-app/folders/)                                                                           |
| Module Splitting Decision    | [`/spec/01-app/26-module-splitting-decision.md`](/spec/01-app/26-module-splitting-decision.md)                             |
| Go Modernization Plan        | [`/spec/01-app/11-go-modernization.md`](/spec/01-app/11-go-modernization.md)                                               |
| CMD Entrypoints              | [`/spec/01-app/12-cmd-entrypoints.md`](/spec/01-app/12-cmd-entrypoints.md)                                                 |
| Testing Patterns             | [`/spec/01-app/13-testing-patterns.md`](/spec/01-app/13-testing-patterns.md)                                               |
| **newCreator Convention**    | [`/spec/01-app/18-new-creator-convention.md`](/spec/01-app/18-new-creator-convention.md)                                   |
| Interface Conventions        | [`/spec/01-app/14-core-interface-conventions.md`](/spec/01-app/14-core-interface-conventions.md)                           |
| Code Review Report           | [`/spec/01-app/15-code-review-report.md`](/spec/01-app/15-code-review-report.md)                                           |
| **Testing Guidelines**       | [`/spec/01-app/16-testing-guidelines.md`](/spec/01-app/16-testing-guidelines.md)                                           |
| **Coding Guidelines**        | [`/spec/01-app/17-coding-guidelines.md`](/spec/01-app/17-coding-guidelines.md)                                             |
| **Code Strengths Review**    | [`/spec/01-app/19-code-strengths.md`](/spec/01-app/19-code-strengths.md)                                                   |
| **Improvement Plan**         | [`/spec/01-app/20-improvement-plan.md`](/spec/01-app/20-improvement-plan.md)                                               |
| **newCreator Pattern**       | [`/spec/01-app/21-new-creator-pattern.md`](/spec/01-app/21-new-creator-pattern.md)                                         |
| **coregeneric Architecture** | [`/spec/01-app/22-coregeneric-architecture.md`](/spec/01-app/22-coregeneric-architecture.md)                               |
| Known Issues                 | [`/spec/13-app-issues/`](/spec/13-app-issues/)                                                                             |
| **Edge-Case Coverage Audit** | [`/spec/13-app-issues/testing/02-edge-case-coverage-audit.md`](/spec/13-app-issues/testing/02-edge-case-coverage-audit.md) |
| **Deep Coverage Scan**       | [`/spec/13-app-issues/testing/03-deep-coverage-scan.md`](/spec/13-app-issues/testing/03-deep-coverage-scan.md)             |
| **GoConvey Migration Plan**  | [`/spec/13-app-issues/testing/04-goconvey-migration-plan.md`](/spec/13-app-issues/testing/04-goconvey-migration-plan.md)   |

## Acknowledgement

External packages used:

- [`github.com/smartystreets/goconvey`](https://github.com/smartystreets/goconvey) — BDD-style testing
- [`github.com/smarty/assertions`](https://github.com/smarty/assertions) — assertion library
- [`golang.org/x/tools`](https://pkg.go.dev/golang.org/x/tools) — Go tooling support

## Reference Links

- [Go Slice Tricks Cheat Sheet](https://ueokande.github.io/go-slice-tricks/)
- [SliceTricks · golang/go Wiki](https://github.com/golang/go/wiki/SliceTricks)
- [Calling a method on a nil struct pointer](https://t.ly/aTp0)
- [Array of pointers to JSON](https://stackoverflow.com/questions/28101966/array-of-pointers-to-json)

## Issues

- [Create an issue](https://github.com/alimtvnetwork/core/-/issues)

## Contributors

- [**Md. Alim Ul Karim**](https://www.linkedin.com/in/alimkarim) — Creator & Lead Architect. System architect with 20+ years of professional software engineering experience across enterprise, fintech, and distributed systems. Recognized as one of the top software architects globally. Alim's architectural philosophy — consistency over cleverness, convention over configuration — is the driving force behind every design decision in this framework.
  - [Google Profile](https://www.google.com/search?q=Alim+Ul+Karim)
- [Riseup Asia LLC (Top Leading Software Company in WY)](https://riseup-asia.com) (2026)
  - [Facebook](https://www.facebook.com/riseupasia.talent/)
  - [LinkedIn](https://www.linkedin.com/company/105304484/)
  - [YouTube](https://www.youtube.com/@riseup-asia)

## License

MIT License — Copyright (c) 2020–2026. See [LICENSE](LICENSE)...
