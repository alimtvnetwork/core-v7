# LLM Integration Guide — `github.com/alimtvnetwork/core`

> **Purpose**: A single-file reference for any LLM or AI agent that needs to understand, use, or extend this Go utility framework. Read this before writing any code that imports `core`.

---

## Table of Contents

1. [Module Identity](#module-identity)
2. [Design Philosophy](#design-philosophy)
3. [Package Map](#package-map)
4. [Import Conventions](#import-conventions)
5. [Core Root Package](#core-root-package)
6. [constants — Shared Constants](#constants--shared-constants)
7. [conditional — Ternary & Nil-Safe Helpers](#conditional--ternary--nil-safe-helpers)
8. [errcore — Error Construction](#errcore--error-construction)
9. [coreinterface — Interface Contracts](#coreinterface--interface-contracts)
10. [Enum System (enuminf + enumimpl)](#enum-system-enuminf--enumimpl)
11. [coredata — Data Structures & JSON](#coredata--data-structures--json)
12. [converters — Type Conversions](#converters--type-conversions)
13. [Utility Packages](#utility-packages)
14. [Testing Patterns](#testing-patterns)
15. [Code Style Rules](#code-style-rules)
16. [Common Mistakes to Avoid](#common-mistakes-to-avoid)
17. [Quick-Start Recipes](#quick-start-recipes)

---

## Module Identity

```
module github.com/alimtvnetwork/core
go 1.25.0
```

- **Zero external runtime dependencies** — only `github.com/smarty/assertions` and `github.com/smartystreets/goconvey` for testing.
- Install: `go get github.com/alimtvnetwork/core`

---

## Design Philosophy

| Principle | Rule |
|-----------|------|
| **One file per function** | Each public function lives in its own `.go` file, named after the function. Files stay 50–200 lines. |
| **Struct-as-namespace** | Related operations group on unexported struct types exposed via package-level `var`. E.g., `corejson.Serialize.ToString()`. |
| **Interface-first** | Contracts in `coreinterface/` using Go's `-er` suffix (`NameGetter`, `Serializer`). Depend on interfaces, not concrete types. |
| **Zero-nil safety** | Return empty slices/maps instead of nil. Pointer-receiver methods include nil guards. |
| **Generics where clear** | Generic versions alongside backward-compatible typed wrappers. |
| **Value receivers** | Read-only methods use value receivers. Pointer receivers only for mutation, large structs, or interface satisfaction. |
| **`newCreator` pattern** | Factories exposed via `New` package variable: `enumimpl.New.BasicByte.UsingTypeSlice(...)` |

---

## Package Map

```
core/
├── core.go, generic.go              # Root — generic slice/map factories
├── conditional/                     # Ternary helpers: If[T], IfFunc[T], NilDef[T]
├── constants/                       # 400+ named constants (strings, bytes, runes, numbers)
├── converters/                      # String↔bytes, maps, JSON formatting
│
├── coredata/                        # Data structures umbrella
│   ├── coreapi/                     #   Typed API request/response models
│   ├── coredynamic/                 #   Reflection-based dynamic data
│   ├── coregeneric/                 #   Generic Collection[T], Hashset[T], Hashmap[K,V]
│   ├── corejson/                    #   JSON pipeline: Serialize.*, Deserialize.*
│   ├── coreonce/                    #   Compute-once lazy values
│   ├── corepayload/                 #   PayloadWrapper — structured data transport
│   ├── corerange/                   #   Range types (int, byte)
│   ├── corestr/                     #   String Collection, Hashset, Hashmap
│   └── stringslice/                 #   80+ pure []string manipulation functions
│
├── coreinterface/                   # 100+ canonical interface contracts
│   ├── enuminf/                     #   Enum interfaces (BasicEnumer, BaseEnumer, etc.)
│   ├── errcoreinf/                  #   Error wrapper interfaces
│   ├── loggerinf/                   #   Logger interfaces
│   ├── serializerinf/               #   Serialization contracts
│   ├── entityinf/                   #   Entity-level interfaces
│   └── payloadinf/                  #   Payload interfaces
│
├── coreimpl/
│   └── enumimpl/                    #   Enum implementation engine
│       └── enumtype/                #     Enum type metadata (Variant)
│
├── errcore/                         # Rich error construction + stack traces
├── corefuncs/                       # Function type wrappers (ErrFunc, InOutErrFuncWrapper)
├── corevalidator/                   # Line, slice, text, range validators
├── coremath/                        # Min/Max for all numeric types
├── coresort/                        # Quick sort (strings, integers)
├── corecmp/                         # Typed comparison helpers
├── coreappend/                      # Append/prepend with nil-skip
├── coreunique/                      # Uniqueness helpers
├── isany/                           # Type checking predicates (Null, Zero, DeepEqual)
├── issetter/                        # 6-valued boolean (Uninitialized/True/False/Unset/Set/Wildcard)
├── bytetype/                        # Byte type utilities
├── namevalue/                       # Name-value pair types
├── keymk/                           # Key compilation with legends/templates
├── regexnew/                        # Lazy-compiled regex with thread-safe caching
├── chmodhelper/                     # File permission parsing/verification
├── typesconv/                       # Pointer ↔ value conversions
├── reflectcore/                     # Reflection utilities
├── mutexbykey/                      # Per-key mutex locking
├── defaultcapacity/                 # Default capacity constants
├── defaulterr/                      # Default error types
├── ostype/                          # OS type detection
├── osconsts/                        # OS-specific constants
├── filemode/                        # File mode types
├── pagingutil/                      # Paging/pagination utilities
├── coreversion/                     # Semantic versioning
│
├── coretests/                       # Testing framework
│   ├── args/                        #   Test argument types (FuncWrap, Map)
│   └── coretestcases/               #   CaseV1 test case definitions
│
├── enums/                           # Domain enum packages
│   ├── stringcompareas/             #   String comparison area enum
│   └── versionindexes/              #   Version index enum
│
└── internal/                        # Not importable externally
    ├── convertinternal/
    ├── reflectinternal/
    └── strutilinternal/
```

---

## Import Conventions

```go
import (
    // Root package
    "github.com/alimtvnetwork/core"

    // Sub-packages — use full path
    "github.com/alimtvnetwork/core/conditional"
    "github.com/alimtvnetwork/core/constants"
    "github.com/alimtvnetwork/core/converters"
    "github.com/alimtvnetwork/core/errcore"
    "github.com/alimtvnetwork/core/coredata/corejson"
    "github.com/alimtvnetwork/core/coredata/corestr"
    "github.com/alimtvnetwork/core/coredata/coregeneric"
    "github.com/alimtvnetwork/core/coreinterface/enuminf"
    "github.com/alimtvnetwork/core/coreimpl/enumimpl"
    "github.com/alimtvnetwork/core/isany"
    "github.com/alimtvnetwork/core/issetter"
)
```

**Never import `internal/` packages from outside the module.**

---

## Core Root Package

Generic slice/map factories. Prefer the non-deprecated versions:

```go
// Create empty slice (non-nil)
ints := core.EmptySlice[int]()           // []int{}

// Create slice with specific length
strs := core.SliceByLength[string](10)   // []string with len=10

// Create slice with length and capacity
buf := core.SliceByCapacity[byte](0, 1024)

// Create empty map pointer
m := core.EmptyMapPtr[string, int]()     // *map[string]int{}
```

**Deprecated** (still works): `EmptySlicePtr`, `SlicePtrByLength`, `SlicePtrByCapacity` — use non-pointer versions.

---

## constants — Shared Constants

**400+ named constants** — never hardcode these values. Always use `constants.X` instead of the raw string/byte/rune.

### Commonly Used

| Constant | Value | Type |
|----------|-------|------|
| `EmptyString` | `""` | `string` |
| `Space` | `" "` | `string` |
| `Comma` | `","` | `string` |
| `CommaSpace` | `", "` | `string` |
| `Dot` | `"."` | `string` |
| `Hyphen` | `"-"` | `string` |
| `Underscore` | `"_"` | `string` |
| `Colon` | `":"` | `string` |
| `ForwardSlash` | `"/"` | `string` |
| `DefaultLine` | `"\n"` | `string` |
| `Tab` | `"\t"` | `string` |
| `InvalidValue` / `InvalidIndex` | `-1` | `int` |
| `Zero` | `0` | `int` |
| `One` | `1` | `int` |
| `N0`–`N32` | `0`–`32` | `int` |
| `N0String`–`N10String` | `"0"`–`"10"` | `string` |
| `SpaceByte` | `' '` | `byte` |
| `DotChar` | `'.'` | `byte` |
| `DotRune` | `'.'` | `rune` |
| `MaxUnit8` | `255` | `byte` |
| `MaxInt16` | `math.MaxInt16` | — |
| `CompareEqual` / `CompareLess` / `CompareGreater` | `0` / `-1` / `1` | `int` |

### OS-Aware Line Endings

Use `constants.DefaultLine` (always `"\n"`). For Windows-specific: `constants.NewLineWindows` (`"\r\n"`). Platform-specific `Line` variable is in `constants/line_*.go`.

### Naming Convention

- **Byte variants**: `SpaceByte`, `DotChar`, `CommaChar`
- **Rune variants**: `SpaceRune`, `DotRune`, `CommaRune`
- **Compound**: `CommaSpace`, `SpaceHyphenSpace`, `NewLineHyphenSpace`

---

## conditional — Ternary & Nil-Safe Helpers

### Generic Base Functions

```go
// Ternary
result := conditional.If[int](isReady, 200, 500)

// Lazy ternary — only evaluates the chosen branch
val := conditional.IfFunc[string](ok,
    func() string { return expensiveCall() },
    func() string { return "fallback" },
)

// Nil-safe default
val := conditional.NilDef[int](ptr, 42)         // dereference or 42
p := conditional.NilDefPtr[string](ptr, "x")     // pointer or &"x"

// Zero-value deref
active := conditional.ValueOrZero[bool](flagPtr) // false if nil
```

### Typed Wrappers (15 primitive types)

Each type gets 11 functions — no type parameter needed:

```go
conditional.IfInt(cond, 2, 7)
conditional.IfFuncString(ok, trueFunc, falseFunc)
conditional.NilDefFloat64(ptr, 3.14)
conditional.ValueOrZeroBool(flagPtr)
```

Supported types: `bool`, `byte`, `string`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`.

### Batch Execution

```go
// Error functions with aggregation
err := conditional.ErrorFunc(fn1, fn2, fn3)

// Typed error functions
results, err := conditional.TypedErrorFunctionsExecuteResults[string](fn1, fn2)
```

---

## errcore — Error Construction

### RawErrorType — Typed Error Categories

`RawErrorType` is a `string` type with 80+ predefined error categories:

```go
// Common types
errcore.InvalidValueType            // "Invalid : value cannot process it."
errcore.CannotBeNilOrEmptyType      // "Values or value cannot be nil or null or empty."
errcore.NotFound                    // "not found"
errcore.FailedToParseType           // "Failed : request failed to parse!"
errcore.ValidationFailedType        // "Validation failed!"
errcore.UnMarshallingFailedType     // "Failed to unmarshal or deserialize."
errcore.OutOfRangeType              // "Out of range : given value, cannot process it."
```

### Creating Errors from RawErrorType

```go
// With reference
err := errcore.InvalidValueType.Error("field name", someRef)

// Format string
err := errcore.FailedToParseType.Fmt("cannot parse %q as date", input)

// Conditional
err := errcore.ValidationFailedType.FmtIf(len(name) == 0, "name is required")

// No reference
err := errcore.NotFound.ErrorNoRefs("user with id 42")

// Merge with existing error
err := errcore.FailedToConvertType.MergeError(originalErr)
err := errcore.FailedToConvertType.MergeErrorWithMessage(originalErr, "while converting X")
```

### Struct-as-Namespace Entry Points

```go
// Assertion-style
msg := errcore.ShouldBe.StrEqMsg("actual", "expected")
err := errcore.ShouldBe.AnyEqErr(got, want)

// Expectation comparison (with type info)
err := errcore.Expected.But("config", "production", "staging")
err := errcore.Expected.ButUsingType("field", 42, "not a number")

// Stack trace enhancement
err := errcore.StackEnhance.Error(originalErr)
msg := errcore.StackEnhance.Msg("something went wrong")
```

### Variable Formatting

```go
// Two-variable context
msg := errcore.VarTwo("src", srcVal, "dst", dstVal)
// → "(src [t:string], dst[t:int]) = (hello, 42)"

// Without types
msg := errcore.VarTwoNoType("left", 5, "right", 10)
// → "(left, right) = (5, 10)"

// Message + variable map
msg := errcore.MessageVarMap("validation failed", map[string]any{"field": "email", "reason": "invalid"})
```

### Error Combining

```go
combined := errcore.MergeErrors(err1, err2, err3)
singleErr := errcore.ManyErrorToSingle(errorSlice)
errFromStrings := errcore.SliceToError([]string{"issue 1", "issue 2"})
```

### Function Types

```go
errcore.ErrFunc          // func() error
errcore.ErrBytesFunc     // func() ([]byte, error)
errcore.ErrStringsFunc   // func() ([]string, error)
errcore.ErrStringFunc    // func() (string, error)
errcore.ErrAnyFunc       // func() (any, error)
```

---

## coreinterface — Interface Contracts

100+ composable interfaces following `-er` suffix convention. **Key categories**:

| Pattern | Examples | Purpose |
|---------|----------|---------|
| `*Getter` | `NameGetter`, `ValueStringGetter` | Read a value |
| `*Checker` | `IsEmptyChecker`, `IsValidChecker` | Boolean predicate |
| `*Binder` | `ContractsBinder`, `JsonContractsBinder` | Compose interfaces |
| `*er` | `Serializer`, `Stringer`, `Disposer` | Action performer |

### Sub-packages

| Package | Key Interfaces |
|---------|----------------|
| `enuminf/` | `BaseEnumer`, `BasicEnumer`, `StandardEnumer`, `BasicByteEnumer`, `BasicInt8Enumer`, etc. |
| `errcoreinf/` | Error wrappers, should-be assertions |
| `loggerinf/` | Logger contracts |
| `serializerinf/` | Serialization/deserialization contracts |
| `entityinf/` | Entity identity and lifecycle |
| `payloadinf/` | Payload transport interfaces |

### Composition Pattern

```go
// Interfaces are small and composable
type IsSuccessValidator interface {
    IsValidChecker    // IsValid() bool
    IsSuccessChecker  // IsSuccess() bool
    IsFailedChecker   // IsFailed() bool
}
```

---

## Enum System (enuminf + enumimpl)

> **Full details**: See `spec/01-app/29-enum-authoring-guide.md`

### Architecture

```
enuminf (interfaces) → enumimpl (implementation engine) → your enum package
```

### Supported Backing Types

| Type | Creator | Use When |
|------|---------|----------|
| `byte` | `enumimpl.New.BasicByte` | ≤255 values (most common) |
| `int8` | `enumimpl.New.BasicInt8` | ≤127 signed values |
| `int16` | `enumimpl.New.BasicInt16` | Larger ordinal space |
| `int32` | `enumimpl.New.BasicInt32` | Large values, 32-bit interop |
| `uint16` | `enumimpl.New.BasicUInt16` | Unsigned 16-bit |
| `string` | `enumimpl.New.BasicString` | String-backed enums |

### Minimal Enum Recipe (byte)

**Step 1: Define constants** (`consts.go`)

```go
package status

type Status byte

const (
    Invalid Status = iota
    Pending
    Ready
    Failed
)
```

**Step 2: Create lookup data** (`vars.go`)

```go
package status

import (
    "github.com/alimtvnetwork/core/coreimpl/enumimpl"
    "github.com/alimtvnetwork/core/internal/reflectinternal"
)

var (
    Ranges = [...]string{
        Invalid: "Invalid",
        Pending: "Pending",
        Ready:   "Ready",
        Failed:  "Failed",
    }

    BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
        reflectinternal.TypeName(Invalid),
        Ranges[:],
    )
)
```

**Step 3: Implement methods** (`Status.go`) — all methods are required:

```go
package status

import "github.com/alimtvnetwork/core/coreinterface/enuminf"

// Value accessors (BasicEnumValuer) — ALL required
func (it Status) Value() byte         { return byte(it) }
func (it Status) ValueByte() byte     { return byte(it) }
func (it Status) ValueInt() int       { return int(it) }
func (it Status) ValueInt8() int8     { return int8(it) }
func (it Status) ValueInt16() int16   { return int16(it) }
func (it Status) ValueUInt16() uint16 { return uint16(it) }
func (it Status) ValueInt32() int32   { return int32(it) }
func (it Status) ValueString() string { return BasicEnumImpl.ToNumberString(it.Value()) }

// Naming
func (it Status) Name() string        { return BasicEnumImpl.ToEnumString(it.Value()) }
func (it Status) String() string      { return BasicEnumImpl.ToEnumString(it.Value()) }
func (it Status) TypeName() string    { return BasicEnumImpl.TypeName() }
func (it Status) NameValue() string   { return BasicEnumImpl.NameWithValue(it.Value()) }
func (it Status) ToNumberString() string { return BasicEnumImpl.ToNumberString(it.Value()) }

// Equality
func (it Status) IsNameEqual(name string) bool { return it.Name() == name }
func (it Status) IsAnyNamesOf(names ...string) bool {
    n := it.Name()
    for _, name := range names { if name == n { return true } }
    return false
}

// Valid/Invalid
func (it Status) IsValid() bool   { return it != Invalid }
func (it Status) IsInvalid() bool { return it == Invalid }

// Range info (BasicEnumer)
func (it Status) RangeNamesCsv() string              { return BasicEnumImpl.RangeNamesCsv() }
func (it Status) MinMaxAny() (min, max any)          { return BasicEnumImpl.MinMaxAny() }
func (it Status) MinValueString() string             { return BasicEnumImpl.MinValueString() }
func (it Status) MaxValueString() string             { return BasicEnumImpl.MaxValueString() }
func (it Status) MaxInt() int                        { return BasicEnumImpl.MaxInt() }
func (it Status) MinInt() int                        { return BasicEnumImpl.MinInt() }
func (it Status) RangesDynamicMap() map[string]any   { return BasicEnumImpl.RangesDynamicMap() }
func (it Status) AllNameValues() []string            { return BasicEnumImpl.AllNameValues() }
func (it Status) IntegerEnumRanges() []int           { return BasicEnumImpl.IntegerEnumRanges() }

// OnlySupportedNamesErrorer
func (it Status) OnlySupportedErr(names ...string) error {
    return BasicEnumImpl.OnlySupportedErr(names...)
}
func (it Status) OnlySupportedMsgErr(message string, names ...string) error {
    return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

// Format — keys: {type-name}, {name}, {value}
func (it Status) Format(format string) string {
    return BasicEnumImpl.Format(format, it.Value())
}

// Type-specific (BasicByteEnumer)
func (it Status) MaxByte() byte      { return BasicEnumImpl.Max() }
func (it Status) MinByte() byte      { return BasicEnumImpl.Min() }
func (it Status) RangesByte() []byte { return BasicEnumImpl.Ranges() }

// Range validation
func (it Status) IsValidRange() bool           { return BasicEnumImpl.IsValidRange(it.Value()) }
func (it Status) IsInvalidRange() bool         { return !it.IsValidRange() }
func (it Status) RangesInvalidMessage() string { return BasicEnumImpl.RangesInvalidMessage() }
func (it Status) RangesInvalidErr() error      { return BasicEnumImpl.RangesInvalidErr() }

// String ranges
func (it Status) StringRanges() []string    { return BasicEnumImpl.StringRanges() }
func (it Status) StringRangesPtr() []string { return BasicEnumImpl.StringRangesPtr() }

// JSON
func (it Status) MarshalJSON() ([]byte, error) {
    return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}
func (it *Status) UnmarshalJSON(data []byte) error {
    val, err := it.UnmarshallEnumToValue(data)
    if err == nil { *it = Status(val) }
    return err
}
func (it Status) UnmarshallEnumToValue(data []byte) (byte, error) {
    return BasicEnumImpl.UnmarshallToValue(true, data)
}

// EnumType
func (it Status) EnumType() enuminf.EnumTyper {
    return BasicEnumImpl.EnumType()
}

// Domain-specific checkers
func (it Status) IsPending() bool { return it == Pending }
func (it Status) IsReady() bool   { return it == Ready }
func (it Status) IsFailed() bool  { return it == Failed }
```

### Adapting for Other Backing Types

| Backing Type | `Value()` returns | Unmarshal method name | Type-specific methods |
|---|---|---|---|
| `byte` | `byte` | `UnmarshallEnumToValue` | `MaxByte`, `MinByte`, `RangesByte` |
| `int8` | `int8` | `UnmarshallEnumToValueInt8` | `MaxInt8`, `MinInt8`, `RangesInt8`, `ToEnumString(int8)` |
| `int16` | `int16` | `UnmarshallEnumToValueInt16` | `MaxInt16`, `MinInt16`, `RangesInt16`, `ToEnumString(int16)` |
| `int32` | `int32` | `UnmarshallEnumToValueInt32` | `MaxInt32`, `MinInt32`, `RangesInt32`, `ToEnumString(int32)` |

### Factory Method Reference

| Method | Description |
|--------|-------------|
| `UsingTypeSlice(typeName, names[])` | Contiguous iota from string slice |
| `Default(firstItem, names[])` | Same, infers typeName via reflection |
| `DefaultWithAliasMap(firstItem, names[], aliasMap)` | Contiguous + aliases |
| `CreateUsingMap(typeName, map[T]string)` | Non-contiguous explicit values |
| `CreateUsingMapPlusAliasMap(typeName, map[T]string, aliasMap)` | Explicit + aliases |

---

## coredata — Data Structures & JSON

### JSON Pipeline

```go
import "github.com/alimtvnetwork/core/coredata/corejson"

// Serialize
jsonStr, err := corejson.Serialize.ToString(myStruct)
jsonBytes, err := corejson.Serialize.Raw(myStruct)

// Deserialize
err := corejson.Deserialize.UsingBytes(jsonBytes, &target)
err := corejson.Deserialize.FromTo(source, &target)  // deep copy via JSON

// Pretty print
pretty := corejson.NewPtr(myStruct).PrettyJsonString()
```

### String Collections

```go
import "github.com/alimtvnetwork/core/coredata/corestr"

collection := corestr.NewCollectionPtrUsingStrings(&values, 0)
collection.AddsLock("new item")  // thread-safe add
fmt.Println(collection.Length())
```

### Generic Collections

```go
import "github.com/alimtvnetwork/core/coredata/coregeneric"

// Hashset, Hashmap, Collection[T], LinkedList[T], SimpleSlice[T]
```

### Compute-Once Values

```go
import "github.com/alimtvnetwork/core/coredata/coreonce"
// Lazy-evaluated cached values for all common types
```

### PayloadWrapper

```go
import "github.com/alimtvnetwork/core/coredata/corepayload"

payload := corepayload.New.PayloadWrapper.Empty()
payload = corepayload.New.PayloadWrapper.UsingInstruction(&corepayload.PayloadCreateInstruction{
    Name:       "user-create",
    Identifier: "usr-123",
    Payloads:   myStruct,
})
```

---

## converters — Type Conversions

```go
import "github.com/alimtvnetwork/core/converters"

// String → integer
val, err := converters.StringTo.Integer("42")
val, ok := converters.StringTo.IntegerWithDefault("abc", -1)

// String → float64
f, err := converters.StringTo.Float64("3.14")

// String → byte
b, err := converters.StringTo.Byte("255")

// Bytes → string
s := converters.BytesTo.String([]byte("hello"))

// Pretty JSON
prettyStr := converters.PrettyJson.String(jsonBytes)
```

---

## Utility Packages

### isany — Type Predicates

```go
import "github.com/alimtvnetwork/core/isany"

isany.Null(val)              // true if nil
isany.Defined(val)           // true if non-nil
isany.Zero(val)              // true if zero value
isany.DeepEqual(a, b)        // reflect.DeepEqual wrapper
isany.JsonEqual(a, b)        // compare via JSON serialization
```

### issetter — 6-Valued Boolean

```go
import "github.com/alimtvnetwork/core/issetter"

// 6 states: Uninitialized(0), True(1), False(2), Unset(3), Set(4), Wildcard(5)
status := issetter.True
status.IsOn()              // true (True or Set)
status.IsOff()             // false (False or Unset)
status.HasInitialized()    // true
```

### regexnew — Lazy Compiled Regex

```go
import "github.com/alimtvnetwork/core/regexnew"

// Package-level (no lock needed at init)
var digitRegex = regexnew.New.Lazy(`\d+`)

// Inside methods (thread-safe)
lazy := regexnew.New.LazyLock(`^[a-z]+\d+$`)
lazy.IsMatch(input)
```

### coremath — Min/Max

```go
import "github.com/alimtvnetwork/core/coremath"
// Min/Max for byte, int, int16, int32, int64, float32, float64
```

### corecmp — Typed Comparisons

```go
import "github.com/alimtvnetwork/core/corecmp"
// Byte, Integer, Integer8/16/32/64, String, Time comparisons + pointer variants
```

### coresort — Sorting

```go
import "github.com/alimtvnetwork/core/coresort/strsort"

fruits := []string{"banana", "mango", "apple"}
strsort.Quick(&fruits)    // [apple banana mango]
strsort.QuickDsc(&fruits) // [mango banana apple]
```

### corevalidator — Validators

Line, slice, text, and range validators with assertion capabilities.

### corefuncs — Function Wrappers

```go
import "github.com/alimtvnetwork/core/corefuncs"
// GetFuncName, GetFuncFullName — for debug/error reporting
// ActionReturnsErrorFuncWrapper, InOutErrFuncWrapper, etc.
```

### namevalue — Name-Value Pairs

```go
import "github.com/alimtvnetwork/core/namevalue"
// Instance (single pair), Collection (multiple pairs)
```

### keymk — Key Compilation

```go
import "github.com/alimtvnetwork/core/keymk"
// Template-based key builders with legends and placeholders
```

---

## Testing Patterns

### Test Structure

```
tests/integratedtests/
└── <package_name>/
    ├── somefunc_testcases.go    # Test data (CaseV1 slices)
    └── somefunc_test.go         # Test logic (AAA pattern)
```

### CaseV1 Pattern

```go
import (
    "github.com/alimtvnetwork/core/coretests/coretestcases"
    "github.com/alimtvnetwork/core/coretests/args"
    "github.com/alimtvnetwork/core/coretests"
)

// _testcases.go
var myTestCases = []coretestcases.CaseV1{
    {
        Title: "valid input returns expected output",
        ArrangeInput: args.Map{
            "actual": "hello",
            "expect": "HELLO",
        },
        ExpectedInput: []string{"HELLO"},
    },
}

// _test.go
func Test_MyFunction(t *testing.T) {
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

---

## Code Style Rules

| Rule | Details |
|------|---------|
| File naming | `FunctionName.go` — one public function per file |
| Receiver name | Always `it` |
| Constructor pattern | `newCreator` struct + `New` package variable |
| Error returns | Use `errcore.RawErrorType.Error(...)` — never raw `errors.New` for categorized errors |
| Nil returns | Return empty slice/map instead of nil |
| Constants | Always use `constants.X` — never hardcode `""`, `" "`, `","`, etc. |
| Generics | Prefer generic functions; add typed wrappers only for the 15 primitive types |
| Interfaces | Define in `coreinterface/` with `-er` suffix; keep small and composable |
| Package vars | Use `var` blocks in `vars.go` for singletons and factories |
| Split large files | By responsibility: `.naming.go`, `.json.go`, `.checkers.go`, `.values.go` |
| Bool-flag methods | Split into expressive pairs — never use a `bool` to switch behavior |

### Method Writing: Split Boolean-Flag Methods

**Critical rule**: When a method's behavior changes based on a `bool` parameter, create **two separate methods** with names that express each behavior. The caller's code reads like documentation — no need to check what `true` or `false` means.

#### Pattern 1: Lock vs No-Lock (`*Lock`)

```go
// ✅ Good: Named variants
func (it *Collection) Add(str string) *Collection { ... }      // no lock
func (it *Collection) AddLock(str string) *Collection {         // thread-safe
    it.Lock()
    defer it.Unlock()
    return it.Add(str)
}

// ❌ Bad: Boolean flag
func (it *Collection) Add(str string, useLock bool) *Collection { ... }
```

#### Pattern 2: Conditional Execution (`*If`)

```go
// ✅ Good: Always-execute + conditional variant
func FmtDebug(format string, items ...any) {
    slog.Debug(fmt.Sprintf(format, items...))
}

func FmtDebugIf(isDebug bool, format string, items ...any) {
    if !isDebug { return }
    FmtDebug(format, items...)
}
```

#### Pattern 3: Behavioral Pairs

```go
// ✅ Good: Opposite states as separate methods
func (it Status) IsValid() bool   { return it != Invalid }
func (it Status) IsInvalid() bool { return it == Invalid }

// ❌ Bad: Single method with negation flag
func (it Status) IsValid(negate bool) bool { ... }
```

#### Pattern 4: Conditional Locking (`*LockIf`)

```go
func (it *lazyRegexMap) CreateOrExisting(pattern string) (*LazyRegex, bool) { ... }
func (it *lazyRegexMap) CreateOrExistingLock(pattern string) (*LazyRegex, bool) { ... }
func (it *lazyRegexMap) CreateOrExistingLockIf(isLock bool, pattern string) (*LazyRegex, bool) {
    if isLock { return it.CreateOrExistingLock(pattern) }
    return it.CreateOrExisting(pattern)
}
```

#### Pattern 5: Collection Conditionals (`AddIf`, `AppendIf`)

```go
func (it *Hashset[T]) AddIf(isAdd bool, key T) *Hashset[T] {
    isSkip := !isAdd
    if isSkip { return it }
    return it.Add(key)
}
```

#### Pattern 6: Filtering Variants (`*NonEmpty`, `*NonEmptyWhitespace`)

String methods provide filtering variants that silently skip items failing a check. Strictness hierarchy:

| Variant | Rejects | Accepts |
|---------|---------|---------|
| `Add` | nothing | everything |
| `AddNonEmpty` | `""` | `" "`, `"a"` |
| `AddNonEmptyWhitespace` | `""`, `" "`, `"\n"` | `"a"` |

```go
// AddNonEmpty — skip empty strings only
func (it *Collection) AddNonEmpty(str string) *Collection {
    if str == "" { return it }
    return it.Add(str)
}

// AddNonEmptyWhitespace — skip empty + whitespace-only
func (it *Collection) AddNonEmptyWhitespace(str string) *Collection {
    if strutilinternal.IsEmptyOrWhitespace(str) { return it }
    return it.Add(str)
}

// Variadic: AddNonEmptyStrings filters each element
func (it *Collection) AddNonEmptyStrings(items ...string) *Collection { ... }

// Standalone slice functions (package stringslice):
stringslice.NonEmptyStrings(slice)    // removes ""
stringslice.NonWhitespace(slice)      // removes "" and whitespace
stringslice.TrimmedEachWords(slice)   // trims + removes empty

// Conditional dispatch:
stringslice.NonEmptyIf(isNonEmpty, slice)
stringslice.TrimmedEachWordsIf(isTrim, slice)

// Filter + join:
stringslice.NonEmptyJoin(slice, ", ")
stringslice.NonWhitespaceJoin(slice, "\n")
```

**Naming rules**: (1) `NonEmpty` = rejects `""` only. (2) `NonEmptyWhitespace`/`NonWhitespace` = rejects `""` + whitespace. (3) `Trimmed*` = trims then rejects empty. (4) `*Strings` suffix for variadic. (5) `*Join` for filter-then-join. (6) `*If` for conditional dispatch. (7) Each variant in its own file.

#### Summary

| Suffix | When | Example |
|--------|------|---------|
| `*Lock` | Thread-safe variant | `Add` → `AddLock` |
| `*If` | Conditional execution | `FmtDebug` → `FmtDebugIf` |
| `*LockIf` | Conditional locking | `Create` → `CreateLockIf` |
| (pair) | Opposite states | `IsValid` + `IsInvalid` |
| `*NonEmpty` | Skip empty strings | `Add` → `AddNonEmpty` |
| `*NonEmptyWhitespace` | Skip empty + whitespace | `Add` → `AddNonEmptyWhitespace` |
| `*NonWhitespace` | Same (standalone functions) | `NonWhitespace(slice)` |
| `*Trimmed*` | Trim then filter | `TrimmedEachWords` |
| `*Join` | Filter then join | `NonEmptyJoin` |

**Rules**: (1) Name expresses behavior. (2) Bool param always first, uses `is*` prefix. (3) `*If` calls the unconditional version — no duplicate logic. (4) Each variant in its own file. (5) Delegate upward — `AddNonEmpty` calls `Add`.

---

## Common Mistakes to Avoid

| ❌ Don't | ✅ Do |
|----------|-------|
| `errors.New("invalid value")` | `errcore.InvalidValueType.Error("fieldName", ref)` |
| `""` in code | `constants.EmptyString` |
| `" "` in code | `constants.Space` |
| `","` in code | `constants.Comma` |
| `"\n"` in code | `constants.DefaultLine` |
| `if x { return a } return b` | `conditional.If[T](x, a, b)` |
| Hardcoded enum string | Implement full enum with `enumimpl` |
| `*` pointer receiver on read methods | Value receiver (`func (it T)`) |
| `func New()` as bare constructor | `newCreator` struct with `New` var |
| Import `internal/` from outside | Use public API only |
| Model bitmask flags as enum | Build flags helper (see `chmodhelper/`) |
| Return `nil` slice | Return `make([]T, 0)` or `core.EmptySlice[T]()` |
| `func Do(flag bool)` for 2 behaviors | Two methods: `Do()` + `DoLock()` or `Do()` + `DoIf()` |

---

## Quick-Start Recipes

### Recipe: Create a New Enum

1. Choose backing type (byte for ≤255 values)
2. Create `consts.go` with `type MyEnum <backing_type>` and `iota` constants
3. Create `vars.go` with `Ranges` array and `BasicEnumImpl` via `enumimpl.New.Basic<Type>.UsingTypeSlice(...)`
4. Create `MyEnum.go` with ALL interface methods (see [Enum System](#enum-system-enuminf--enumimpl))
5. Add domain-specific `IsX()` checkers

### Recipe: Create a Structured Error

```go
err := errcore.InvalidValueType.Error("email", userInput)
err := errcore.ValidationFailedType.Fmt("field %q must be non-empty", fieldName)
err := errcore.FailedToParseType.MergeErrorWithMessage(parseErr, "config file")
```

### Recipe: JSON Round-Trip

```go
jsonBytes, err := corejson.Serialize.Raw(myStruct)
err = corejson.Deserialize.UsingBytes(jsonBytes, &target)
```

### Recipe: Conditional Value

```go
label := conditional.If[string](isAdmin, "Administrator", "User")
timeout := conditional.NilDef[int](configTimeout, 30)
```

### Recipe: Safe String Collection

```go
coll := corestr.NewCollectionPtrUsingStrings(&items, constants.Zero)
coll.AddsLock("new item")
```

---

## Further Reading

| Topic | Location |
|-------|----------|
| Enum authoring (full templates) | `spec/01-app/29-enum-authoring-guide.md` |
| Testing guidelines | `spec/01-app/16-testing-guidelines.md` |
| Coding guidelines | `spec/01-app/17-coding-guidelines.md` |
| AI agent test reference | `spec/03-powershell-test-run/09-ai-agent-complete-reference.md` |
| Package READMEs | Each package has its own `README.md` / `readme.md` |
| Root README | `README.md` |
