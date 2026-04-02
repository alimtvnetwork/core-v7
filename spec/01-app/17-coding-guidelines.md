# Coding Guidelines

## Receiver Types: Prefer Value Receivers (Future Direction)

> **Status**: Guideline for new code. Existing pointer receivers will be migrated incrementally.

### Rationale

- Value receivers are simpler, safer (no nil panics without guards), and communicate immutability.
- Pointer receivers should only be used when **mutating** the struct or when the struct is **large** (>~5 fields with complex types).
- All nil-safety guards (`if it == nil`) become unnecessary with value receivers.

### When to Use Pointer Receivers

- The method **modifies** the receiver (setter, initializer).
- The struct is **large** and copying would be expensive.
- The method must satisfy an interface that requires pointer receiver.
- The method implements `json.Marshaler` / `json.Unmarshaler`.

### When to Use Value Receivers

- The method is a **getter** (read-only).
- The method returns a **computed value** or **formatted string**.
- The struct is small (≤5 simple fields).
- `Json()`, `String()`, `Clone()`, `ToPtr()` — always value receivers.

### Example

```go
// ✅ Good: Value receiver for read-only methods
func (it Info) Name() string { return it.RootName }
func (it Info) Json() corejson.Result { return corejson.New(it) }
func (it Info) Clone() Info { return Info{...} }

// ✅ Good: Pointer receiver for mutation
func (it *Info) SetSecure() *Info { it.ExcludeOptions = ...; return it }
```

### Migration Plan

1. New code: Follow this guideline immediately.
2. Existing code: Migrate during refactoring passes — do NOT change receiver type in isolation (may break interface satisfaction).

---

## `interface{}` → `any` Migration

All new code must use `any` instead of `interface{}`. This is a Go 1.18+ alias and is semantically identical. See [Go Modernization Plan](/spec/01-app/11-go-modernization.md).

---

## One File Per Function

Each public function or method group gets its own `.go` file, named after the function/struct. This keeps files small and focused.

---

## Struct-as-Namespace Pattern

Group related operations on unexported struct types, exposed via package-level `var`:

```go
// unexported struct
type newCreator struct{}

// package-level var
var New newCreator

// usage: corepayload.New.PayloadWrapper.Empty()
```

---

## Zero-Nil Safety

- Return empty slices/maps instead of `nil`.
- All pointer-receiver methods must have nil guards if the receiver could be nil.
- Use `IsNull()` / `IsEmpty()` / `IsDefined()` consistently.

---

## Interface Naming

Follow Go's `-er` suffix convention:

| Pattern | Example |
|---------|---------|
| `*Getter` | `NameGetter`, `ValueGetter` |
| `*Checker` | `HasErrorChecker` |
| `*Binder` | `ContractsBinder`, `AttributesBinder` |
| `*er` | `Serializer`, `Csver`, `Stringer` |

---

## The `newCreator` Convention (Hierarchical Factory Pattern)

This is the **most important architectural pattern** in the codebase. Instead of flat `NewX()` functions, we decompose object creation into a tree of small factory structs exposed via a package-level `var New`:

```go
// vars.go
var New = newCreator{}

// newCreator.go — root aggregator
type newCreator struct {
    Widget newWidgetCreator
    Config newConfigCreator
}

// newWidgetCreator.go — one file per sub-creator
type newWidgetCreator struct{}

func (it newWidgetCreator) Empty() *Widget {
    return &Widget{Items: []string{}}
}

func (it newWidgetCreator) Create(name string) *Widget {
    return &Widget{Name: name, Items: []string{}}
}
```

**Usage**: `mypkg.New.Widget.Empty()` — IDE autocomplete guides users through the tree.

See the full guide: **[newCreator Convention](/spec/01-app/18-new-creator-convention.md)**

---

## Conditional Formatting & Readability

### Prefer Positive Conditions

Use **positive** boolean variables (`isInvalid`, `isEmpty`, `hasError`) rather than negating a variable inline (`!isValid`, `!isEmpty`). This improves readability and makes intent explicit.

```go
// ✅ Good: Positive condition via renamed variable
items, isValid := input.GetAsStrings("items")
isInvalid := !isValid

if isInvalid {
    errcore.HandleErrMessage("GetAsStrings 'items' failed")
}

// ❌ Bad: Negation inline
items, isValid := input.GetAsStrings("items")
if !isValid {
    errcore.HandleErrMessage("GetAsStrings 'items' failed")
}
```

**Exception**: When the variable is used only once and the meaning is obvious (e.g. `if !ok {`), inline negation is acceptable.

### Blank Line Before `return`

Always insert a blank line before a `return` statement when it is preceded by a line of code. This visually separates the function's exit point from its logic.

```go
// ✅ Good: Blank line before return
result := compute(input)

return result

// ✅ Good: Early return guard (no blank line needed after opening `{`)
func (it Info) Name() string {
    if it.IsEmpty() {
        return ""
    }

    return it.RootName
}

// ❌ Bad: No blank line before return
result := compute(input)
return result
```

**Exception**: Single-line function bodies do not need a blank line before `return`:

```go
func (it Info) Name() string { return it.RootName }
```

### Blank Line Rules for Control Flow Blocks

These rules apply uniformly to **all** control flow statements: `if`, `for`, `switch`, `select`, and `range`.

1. **Before the statement**: Always insert a blank line before a control flow statement when preceded by a line of code or a closing `}` (unless that `}` immediately closes an outer block).
2. **After `}`**: Insert a blank line after `}` only if the next line is **not** another `}` closing a parent block.
3. **Consecutive control flow**: When two control flow blocks appear back-to-back with no intervening code, a single blank line separates them.

```go
// ✅ Good: Spacing around if
items, isValid := input.GetAsStrings("items")
isInvalid := !isValid

if isInvalid {
    errcore.HandleErrMessage("GetAsStrings 'items' failed")
}

search, hasSearch := input.GetAsString("search")
isSearchMissing := !hasSearch

if isSearchMissing {
    errcore.HandleErrMessage("GetAsString 'search' failed")
}

// ✅ Good: Spacing around for
col := coredynamic.New.Collection.String.From(items)

for i := 0; i < col.Length(); i++ {
    process(col.SafeAt(i))
}

result := col.First()

// ✅ Good: Spacing around switch
kind := reflect.TypeOf(value).Kind()

switch kind {
case reflect.String:
    handleString(value)
case reflect.Int:
    handleInt(value)
default:
    handleOther(value)
}

// ✅ Good: Spacing around select
timeout := time.After(5 * time.Second)

select {
case msg := <-ch:
    process(msg)
case <-timeout:
    return ErrTimeout
}

// ✅ Good: Spacing around range
names := []string{"a", "b", "c"}

for _, name := range names {
    fmt.Println(name)
}

// ✅ Good: No blank line before closing parent }
for _, item := range items {
    if item == "" {
        continue
    }
}

// ❌ Bad: No breathing room
items, isValid := input.GetAsStrings("items")
isInvalid := !isValid
if isInvalid {
    errcore.HandleErrMessage("GetAsStrings 'items' failed")
}
search, hasSearch := input.GetAsString("search")
if !hasSearch {
    errcore.HandleErrMessage("GetAsString 'search' failed")
}
for i := 0; i < len(items); i++ {
    process(items[i])
}
```

---

## Function Call Argument Formatting

When a function call has **multiple arguments**, each argument must be placed on its own line — including the first argument. The closing parenthesis sits on its own line, aligned with the function call indentation.

```go
// ✅ Good: Each argument on its own line, first argument on the next line
verifyDefaultErr(
    t,
    0,
    "NilResult error is not nil",
    defaulterr.NilResult,
)

errcore.AssertDiffOnMismatch(
    t,
    caseIndex,
    tc.Title,
    actLines,
    expectedLines,
)

req := coreapi.NewTypedSimpleGenericRequest[string](
    attr,
    simpleReq,
)

// ✅ Good: Single argument can stay on the same line
Write-Success "All tests passed"
fmt.Println(value)

// ❌ Bad: Multiple arguments on the same line as function name
verifyDefaultErr(t, 0, "NilResult error is not nil", defaulterr.NilResult)

// ❌ Bad: First argument on the same line as function name
verifyDefaultErr(t,
    0,
    "NilResult error is not nil",
    defaulterr.NilResult,
)
```

**Exception**: Single-argument calls or very short two-argument calls where both fit comfortably on one line (e.g., `fmt.Sprintf("%v", value)`).

---

## First-Item Assertion Convenience Methods

When a test uses a **named single test case** (not a loop with `caseIndex`), use the `*First` assertion variants instead of passing a literal `0` for `caseIndex`. This eliminates magic numbers and improves readability.

```go
// ✅ Good: Named First variant — no magic 0
tc.ShouldBeEqualArgsFirst(
    t,
    emptyBefore,
    lenBefore,
    emptyAfter,
    lenAfter,
)

tc.ShouldMatchExpectedFirst(
    t,
    result,
)

tc.ShouldBeEqualUsingExpectedFirst(
    t,
    actLines,
)

// ❌ Bad: Magic 0 for non-loop test
tc.ShouldBeEqualArgs(
    t,
    0,
    emptyBefore,
    lenBefore,
)

// ✅ Good: Explicit caseIndex in a loop — use the indexed variant
for caseIndex, tc := range testCases {
    tc.ShouldBeEqualArgs(
        t,
        caseIndex,
        result,
    )
}
```

Available `*First` methods on `GenericGherkins`:

| Method | Wraps |
|--------|-------|
| `ShouldBeEqualFirst(t, actLines, expectedLines)` | `ShouldBeEqual(t, 0, ...)` |
| `ShouldBeEqualArgsFirst(t, actLines...)` | `ShouldBeEqualArgs(t, 0, ...)` |
| `ShouldBeEqualUsingExpectedFirst(t, actLines)` | `ShouldBeEqualUsingExpected(t, 0, ...)` |
| `ShouldMatchExpectedFirst(t, result)` | `ShouldMatchExpected(t, 0, ...)` |

## Variable Naming Conventions

### Avoid Numbered Suffixes

Do **not** use numbered variable names like `val1`, `val2`, `var1`, `var2`. Use descriptive names that convey meaning.

```go
// ✅ Good: Descriptive parameter names
func VarTwo(
    isIncludeType bool,
    firstName string,
    firstValue any,
    secondName string,
    secondValue any,
) string { ... }

// ❌ Bad: Numbered suffixes
func VarTwo(
    isIncludeType bool,
    var1 string,
    val1 any,
    var2 string,
    val2 any,
) string { ... }
```

### Naming Guidelines

| Pattern | Good | Bad |
|---------|------|-----|
| Loop variables | `item`, `name`, `key` | `v`, `x`, `tmp` |
| Boolean flags | `isValid`, `hasError` | `ok2`, `flag` |
| Positional params | `firstName`, `secondValue` | `val1`, `val2` |
| Iterators | `index`, `offset` | `i2`, `j2` |

**Exception**: Single-letter variables are acceptable in very short scopes (e.g., `i` in a `for` loop, `k`/`v` in a map range).

---

## Method Writing: Split Boolean-Flag Methods into Expressive Pairs

When a method's behavior changes based on a boolean parameter, **do not write one method with a `bool` flag**. Instead, create **two separate methods** with names that express each behavior. The caller's code then reads like documentation — no need to check what `true` or `false` means.

### The Rule

> **If a `bool` parameter selects between two behaviors, create two functions — one per behavior.**
>
> The `bool`-flag version may still exist as a **dispatcher** (for internal or generic use), but callers should use the named variants.

### Pattern 1: Lock vs No-Lock (Thread Safety)

The most common case in this codebase. Every mutable method has a non-locking version (for use when the caller already holds the lock or in single-threaded contexts) and a `*Lock` version (thread-safe).

```go
// ✅ Good: Two expressive methods — caller picks based on context

// Add appends a string (no locking — caller must manage concurrency).
func (it *Collection) Add(str string) *Collection {
    it.items = append(it.items, str)

    return it
}

// AddLock appends a string with mutex protection (thread-safe).
func (it *Collection) AddLock(str string) *Collection {
    it.Lock()
    defer it.Unlock()

    it.items = append(it.items, str)

    return it
}

// ❌ Bad: Boolean flag — caller must guess what `true` means
func (it *Collection) Add(str string, useLock bool) *Collection {
    if useLock {
        it.Lock()
        defer it.Unlock()
    }

    it.items = append(it.items, str)

    return it
}
```

**Naming convention**: `MethodName` (no lock) + `MethodNameLock` (with lock).

### Pattern 2: Conditional Execution (`*If` suffix)

When an action should only execute under a condition, create the unconditional version plus an `*If` variant. The `*If` method takes the condition as the **first parameter**.

```go
// ✅ Good: Two methods — unconditional + conditional

// FmtDebug always logs a debug message.
func FmtDebug(
    format string,
    items ...any,
) {
    slog.Debug(fmt.Sprintf(format, items...))
}

// FmtDebugIf logs a debug message only when isDebug is true.
func FmtDebugIf(
    isDebug bool,
    format string,
    items ...any,
) {
    if !isDebug {
        return
    }

    slog.Debug(fmt.Sprintf(format, items...))
}
```

**Naming convention**: `MethodName` (always executes) + `MethodNameIf` (conditional).

### Pattern 3: Behavioral Variants (Separate Named Methods)

When a boolean selects between two **different behaviors** (not just "do it" vs "skip it"), create two methods whose names describe the behavior.

```go
// ✅ Good: Behavior expressed in the name

// MsgHeader formats items with a header wrapper.
func MsgHeader(items ...any) string {
    return fmt.Sprintf(msgformats.MsgHeaderFormat, items...)
}

// MsgHeaderIf formats with header when isHeader=true,
// otherwise returns plain fmt.Sprint.
func MsgHeaderIf(
    isHeader bool,
    items ...any,
) string {
    if isHeader {
        return MsgHeader(items...)
    }

    return fmt.Sprint(items...)
}

// ✅ Good: IsValid / IsInvalid instead of IsValid(bool negate)
func (it Status) IsValid() bool   { return it != Invalid }
func (it Status) IsInvalid() bool { return it == Invalid }

// ❌ Bad: Single method with negation flag
func (it Status) IsValid(negate bool) bool {
    if negate {
        return it == Invalid
    }
    return it != Invalid
}
```

### Pattern 4: Lock + Conditional Combined (`*LockIf`)

When both locking and conditional behavior are needed, the `*LockIf` variant takes `isLock bool` as the first parameter and delegates to the appropriate path.

```go
// ✅ Good: Three tiers — no lock, lock, conditional lock

// CreateOrExisting creates a new lazy regex or returns the existing one.
func (it *lazyRegexMap) CreateOrExisting(
    patternName string,
) (*LazyRegex, bool) { ... }

// CreateOrExistingLock same as above, with mutex protection.
func (it *lazyRegexMap) CreateOrExistingLock(
    patternName string,
) (*LazyRegex, bool) {
    it.Lock()
    defer it.Unlock()

    return it.CreateOrExisting(patternName)
}

// CreateOrExistingLockIf conditionally applies mutex based on isLock.
func (it *lazyRegexMap) CreateOrExistingLockIf(
    isLock bool,
    patternName string,
) (*LazyRegex, bool) {
    if isLock {
        return it.CreateOrExistingLock(patternName)
    }

    return it.CreateOrExisting(patternName)
}
```

### Pattern 5: Collection Operations (`AddIf`, `AddsIf`, `PrependIf`)

Collection types provide conditional add/prepend/append variants. The condition is always the **first parameter** named with an `is*` prefix.

```go
// ✅ Good: Conditional collection operations

func (it *Collection[K, V]) AppendIf(
    isAppend bool,
    items ...Instance[K, V],
) *Collection[K, V] {
    isSkip := !isAppend

    if isSkip {
        return it
    }

    return it.Append(items...)
}

func (it *Hashset[T]) AddIf(isAdd bool, key T) *Hashset[T] {
    isSkip := !isAdd

    if isSkip {
        return it
    }

    return it.Add(key)
}
```

### Summary Table

| Suffix | When to Use | Example |
|--------|-------------|---------|
| `*Lock` | Thread-safe variant of a non-locking method | `Add` → `AddLock` |
| `*If` | Executes only when a condition is true | `FmtDebug` → `FmtDebugIf` |
| `*LockIf` | Conditionally applies locking | `Create` → `CreateLockIf` |
| No suffix (pair) | Two methods expressing opposite states | `IsValid` + `IsInvalid` |
| `*NonEmpty` | Variant that skips empty/nil inputs | `Add` → `AddNonEmpty` |

### Rules

1. **Name expresses behavior** — the caller should never need to look up what a `bool` parameter means.
2. **Condition parameter is always first** — `isAdd`, `isLock`, `isDebug`, `isHeader`.
3. **Use `is*` prefix** for all boolean parameters — never `flag`, `option`, `mode`.
4. **The `*If` variant calls the unconditional one** — don't duplicate logic.
5. **Each variant lives in its own file** — `Add.go`, `AddLock.go`, `AddIf.go`.

---

## Related Docs

- [Design Philosophy](/spec/01-app/00-repo-overview.md)
- [Interface Conventions](/spec/01-app/14-core-interface-conventions.md)
- [Go Modernization Plan](/spec/01-app/11-go-modernization.md)
- [newCreator Convention](/spec/01-app/18-new-creator-convention.md)
- [Testing Guidelines](/spec/01-app/16-testing-guidelines.md)
