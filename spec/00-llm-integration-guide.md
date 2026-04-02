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
   - [Commonly Used](#commonly-used) · [OS-Aware Line Endings](#os-aware-line-endings) · [Naming Convention](#naming-convention)
7. [conditional — Ternary & Nil-Safe Helpers](#conditional--ternary--nil-safe-helpers)
   - [Generic Base Functions](#generic-base-functions) · [Typed Wrappers](#typed-wrappers-15-primitive-types) · [Batch Execution](#batch-execution)
8. [errcore — Error Construction](#errcore--error-construction)
   - [RawErrorType](#rawerrortype--typed-error-categories) · [Creating Errors](#creating-errors-from-rawerrortype) · [Struct-as-Namespace Entry Points](#struct-as-namespace-entry-points) · [Variable Formatting](#variable-formatting) · [Error Combining](#error-combining) · [Function Types](#function-types)
9. [coreinterface — Interface Contracts](#coreinterface--interface-contracts)
   - [Sub-packages](#sub-packages) · [Composition Pattern](#composition-pattern)
10. [Enum System (enuminf + enumimpl)](#enum-system-enuminf--enumimpl)
    - [Architecture](#architecture) · [Supported Backing Types](#supported-backing-types) · [Minimal Enum Recipe](#minimal-enum-recipe-byte) · [Adapting for Other Backing Types](#adapting-for-other-backing-types) · [Factory Method Reference](#factory-method-reference)
11. [coredata — Data Structures & JSON](#coredata--data-structures--json)
    - [JSON Pipeline](#json-pipeline) · [String Collections](#string-collections) · [Generic Collections](#generic-collections) · [Compute-Once Values](#compute-once-values) · [PayloadWrapper](#payloadwrapper)
12. [converters — Type Conversions](#converters--type-conversions)
13. [Utility Packages](#utility-packages)
    - [isany](#isany--type-predicates) · [issetter](#issetter--6-valued-boolean) · [regexnew](#regexnew--lazy-compiled-regex) · [coremath](#coremath--minmax) · [corecmp](#corecmp--typed-comparisons) · [coresort](#coresort--sorting) · [corevalidator](#corevalidator--validators) · [corefuncs](#corefuncs--function-wrappers) · [namevalue](#namevalue--name-value-pairs) · [keymk](#keymk--key-compilation)
14. [Testing Patterns](#testing-patterns)
    - [Test Folder Structure](#test-folder-structure) · [CaseV1](#casev1--the-primary-workhorse) · [CaseNilSafe](#casenilsafe--nil-receiver-safety) · [MapGherkins](#mapgherkins--bdd-style-tests) · [args.Map Quick Reference](#argsmap-quick-reference) · [Decision Matrix](#decision-matrix) · [Assertion Methods](#assertion-methods) · [Testing Anti-Patterns](#testing-anti-patterns)
15. [Code Style Rules](#code-style-rules)
    - [Split Boolean-Flag Methods](#method-writing-split-boolean-flag-methods) · [Pointer Variants (`*Ptr`)](#method-writing-pointer-variants-ptr-suffix) · [`*Must` Suffix](#method-writing-must-suffix-panic-on-error) · [`*Slice` vs Variadic](#method-writing-slice-vs-variadic) · [`*Or*` Fallback Pattern](#method-writing-or-fallback-pattern) · [Deprecation Convention](#deprecation-convention)
16. [Common Mistakes to Avoid](#common-mistakes-to-avoid)
17. [Quick-Start Recipes](#quick-start-recipes)
    - [New Enum](#recipe-create-a-new-enum) · [Structured Error](#recipe-create-a-structured-error) · [JSON Round-Trip](#recipe-json-round-trip) · [Conditional Value](#recipe-conditional-value) · [Safe String Collection](#recipe-safe-string-collection) · [PayloadWrapper](#recipe-payloadwrapper-creation--deserialization) · [PayloadsCollection](#recipe-payloadscollection-usage)
18. [coregeneric — Generic Data Structures API Reference](#coregeneric--generic-data-structures-api-reference)
    - [Collection[T]](#collectiont-any) · [Hashset[T]](#hashsett-comparable) · [Hashmap[K,V]](#hashmapk-comparable-v-any) · [SimpleSlice[T]](#simpleslicet-any) · [LinkedList[T]](#linkedlistt-any) · [Pair[L,R]](#pairl-any-r-any) · [Triple[A,B,C]](#triplea-any-b-any-c-any) · [Package-Level Helpers](#package-level-functional-helpers-funcsgo--orderedfuncsgo)
19. [Further Reading](#further-reading)

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

> 📖 **Full reference**: [`/spec/testing-guidelines/`](/spec/testing-guidelines/) — 8 detailed guides.

### Test Folder Structure

```
tests/integratedtests/
├── mypkgtests/                            # One directory per source package
│   ├── params.go                          # Reusable key constants for args.Map
│   ├── MyFunc_testcases.go                # Test data — expectations only (NO import "testing")
│   ├── MyFunc_test.go                     # Test runner — assertions only (NO hardcoded values)
│   ├── MyStruct_NilReceiver_testcases.go  # Nil-safety test data
│   ├── NilReceiver_test.go                # Nil-safety test runner (one per package)
│   └── helpers.go                         # Shared test-only structs/utilities
└── anotherpkgtests/
    └── ...
```

**Critical separation**: Test data (`_testcases.go`) is strictly separated from test logic (`_test.go`). Never put expected values in `_test.go` or test functions in `_testcases.go`.

### CaseV1 — The Primary Workhorse

```go
type CaseV1 struct {
    Title         string  // "FuncName returns X -- given Y input"
    ArrangeInput  any     // Input data (args.Map, args.One, etc.)
    ActualInput   any     // Set dynamically after Act phase
    ExpectedInput any     // Expected output (args.Map, string, []string, bool)
}
```

#### CaseV1 with args.Map (Preferred — Multi-Field)

```go
// _testcases.go
var validateEmailTestCases = []coretestcases.CaseV1{
    {
        Title: "ValidateEmail returns valid -- given well-formed email",
        ArrangeInput: args.Map{
            "email": "user@example.com",
        },
        ExpectedInput: args.Map{
            "isValid":    true,
            "errorCount": 0,
        },
    },
    {
        Title: "ValidateEmail returns invalid -- given empty email",
        ArrangeInput: args.Map{
            "email": "",
        },
        ExpectedInput: args.Map{
            "isValid":    false,
            "errorCount": 1,
        },
    },
}

// _test.go
func Test_ValidateEmail_Verification(t *testing.T) {
    for caseIndex, tc := range validateEmailTestCases {
        // Arrange
        input := tc.ArrangeInput.(args.Map)
        email, _ := input.GetAsString("email")

        // Act
        result := validator.ValidateEmail(email)
        actual := args.Map{
            "isValid":    result.IsValid,
            "errorCount": len(result.Errors),
        }

        // Assert
        tc.ShouldBeEqualMap(t, caseIndex, actual)
    }
}
```

#### CaseV1 — Single Test Case (Named Variable)

For one-off tests, use a named standalone variable and `*First` methods:

```go
// _testcases.go
var concatMessageNilTestCase = coretestcases.CaseV1{
    Title: "ConcatMessageWithErr nil error returns nil",
    ArrangeInput: args.Map{"message": "should not appear"},
    ExpectedInput: args.Map{"isNil": true},
}

// _test.go
func Test_ConcatMessageWithErr_NilPassthrough(t *testing.T) {
    tc := concatMessageNilTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    msg, _ := input.GetAsString("message")

    // Act
    result := errcore.ConcatMessageWithErr(msg, nil)
    actual := args.Map{"isNil": result == nil}

    // Assert
    tc.ShouldBeEqualMapFirst(t, actual)
}
```

#### CaseV1 with String Assertions

For simple single-value comparisons:

```go
// _testcases.go
var toUpperTestCases = []coretestcases.CaseV1{
    {
        Title: "ToUpper converts lowercase -- given hello",
        ArrangeInput: args.Map{"input": "hello"},
        ExpectedInput: []string{"HELLO"},
    },
}

// _test.go
func Test_ToUpper_Verification(t *testing.T) {
    for caseIndex, tc := range toUpperTestCases {
        // Arrange
        input := tc.ArrangeInput.(args.Map)
        str, _ := input.GetAsString("input")

        // Act
        result := strings.ToUpper(str)

        // Assert
        tc.ShouldBeEqual(t, caseIndex, result)
    }
}
```

### CaseNilSafe — Nil-Receiver Safety

**Use ONLY for pointer-receiver methods**, never for package-level functions.

```go
// _NilReceiver_testcases.go
var myStructNilSafeTestCases = []coretestcases.CaseNilSafe{
    {
        Title: "IsValid on nil returns false",
        Func:  (*MyStruct).IsValid,   // Method expression (zero-arg)
        Expected: results.ResultAny{
            Value:    "false",
            Panicked: false,
        },
    },
    {
        Title: "HasKey on nil returns false",
        Func: func(m *MyStruct) bool {  // Wrap methods with args
            return m.HasKey("anything")
        },
        Expected: results.ResultAny{
            Value:    "false",
            Panicked: false,
        },
    },
    {
        Title: "Clear on nil does not panic",
        Func:  (*MyStruct).Clear,
        Expected: results.ResultAny{Panicked: false},
        CompareFields: []string{"panicked", "returnCount"},  // Void — explicit fields
    },
}

// NilReceiver_test.go
func Test_MyStruct_NilReceiver(t *testing.T) {
    for caseIndex, tc := range myStructNilSafeTestCases {
        tc.ShouldBeSafe(t, caseIndex)
    }
}
```

### MapGherkins — BDD-Style Tests

Alias for `GenericGherkins[args.Map, args.Map]`. Use for multi-field I/O with BDD fields.

```go
// params.go — MANDATORY for args.Map key constants
var params = struct {
    pattern      string
    compareInput string
    isDefined    string
    isApplicable string
    isMatch      string
}{
    pattern:      "pattern",
    compareInput: "compareInput",
    isDefined:    "isDefined",
    isApplicable: "isApplicable",
    isMatch:      "isMatch",
}

// _testcases.go
var lazyRegexTestCases = []coretestcases.MapGherkins{
    {
        Title: "New.Lazy matches word pattern",
        When:  "given a simple word pattern",
        Input: args.Map{
            params.pattern:      "hello",
            params.compareInput: "hello world",
        },
        Expected: args.Map{
            params.isDefined:    true,
            params.isApplicable: true,
            params.isMatch:      true,
        },
    },
}

// _test.go
func Test_LazyRegex_New_Verification(t *testing.T) {
    for caseIndex, tc := range lazyRegexTestCases {
        // Arrange
        pattern, _ := tc.Input.GetAsString(params.pattern)
        compareInput, _ := tc.Input.GetAsString(params.compareInput)

        // Act
        lazy := regexnew.New.Lazy(pattern)
        actual := args.Map{
            params.isDefined:    lazy.IsDefined(),
            params.isApplicable: lazy.IsApplicable(),
            params.isMatch:      lazy.IsMatch(compareInput),
        }

        // Assert
        tc.ShouldBeEqualMap(t, caseIndex, actual)
    }
}
```

### args.Map Quick Reference

```go
type Map map[string]any

// Typed getters (return value, ok):
email, ok := input.GetAsString("email")
count, ok := input.GetAsInt("count")
flag, ok  := input.GetAsBool("isEnabled")
items, ok := input.GetAsStrings("tags")

// Default getters (no ok):
count := input.GetAsIntDefault("count", 0)
flag  := input.GetAsBoolDefault("isEnabled", false)

// Presence checks:
input.Has("key")           // exists (may be nil)
input.HasDefined("key")    // exists AND non-nil
input.IsKeyMissing("key")  // does not exist
```

### Decision Matrix

| Condition | Use |
|-----------|-----|
| Standard input → string output | `CaseV1` + `ShouldBeEqual` |
| Standard input → multi-field output | `CaseV1` + `args.Map` + `ShouldBeEqualMap` |
| Nil-receiver safety | `CaseNilSafe` + `ShouldBeSafe` |
| BDD with typed input/output | `GenericGherkins[T1, T2]` |
| Multi-field I/O with semantic keys | `MapGherkins` |
| Domain-specific data extraction | Custom wrapper embedding `CaseV1` |

### Assertion Methods

| Method | Use When |
|--------|----------|
| `ShouldBeEqual(t, idx, actual...)` | Loop, exact string match |
| `ShouldBeEqualFirst(t, actual...)` | Single case (idx=0) |
| `ShouldBeEqualMap(t, idx, actualMap)` | Loop, map comparison |
| `ShouldBeEqualMapFirst(t, actualMap)` | Single case, map comparison |
| `ShouldContains(t, idx, actual...)` | Substring match |
| `ShouldStartsWith(t, idx, actual...)` | Prefix match |
| `ShouldEndsWith(t, idx, actual...)` | Suffix match |
| `ShouldBeNotEqual(t, idx, actual...)` | Inverse match |
| `ShouldBeTrimEqual(t, idx, actual...)` | Trimmed comparison |
| `ShouldBeSortedEqual(t, idx, actual...)` | Sorted + trimmed |
| `ShouldBeRegex(t, idx, actual...)` | Regex match |
| `ShouldBeSafe(t, idx)` | Nil-receiver safety |

### Testing Anti-Patterns

| ❌ Don't | ✅ Do |
|----------|-------|
| Raw `t.Error`/`t.Errorf` | Use `ShouldBeEqual`/`ShouldBeEqualMap` |
| Expected values in `_test.go` | Move to `_testcases.go` |
| `CaseNilSafe` for package functions | Use `CaseV1` |
| Stringified booleans `"true"` | Use native `true` |
| Inline `args.Map` with 2+ entries | Multi-line, one key-value per line |
| Raw string keys `"isDefined"` | Use `params.isDefined` constants |
| Missing AAA comments | Always include `// Arrange`, `// Act`, `// Assert` |
| Vague title `"returns false"` | `"FuncName returns false -- given empty input"` |

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

#### Pattern 7: Combined Suffixes & Ordering

When multiple suffixes combine, they follow a **fixed order**: **Base + Filter + Type + Lock + If + Must**.

| Position | Suffix | Example |
|----------|--------|---------|
| 1 | Base name | `Add`, `Adds`, `Create` |
| 2 | Filter | `NonEmpty`, `NonEmptyWhitespace` |
| 3 | Type modifier | `Strings`, `Slice`, `Ptr` |
| 4 | Lock | `Lock` |
| 5 | If | `If` |
| 6 | Must | `Must` |

Real examples: `Add` → `AddNonEmpty` → `AddNonEmptyStrings` → `AddsNonEmptyPtrLock` → `CreateOrExistingLockIf`.

```go
// Combined Lock + Filter — delegates to simpler variant
func (it *Collection) AddNonEmptyLock(str string) *Collection {
    it.Lock()
    defer it.Unlock()
    return it.AddNonEmpty(str)
}

// Full chain example from codebase:
// AddsNonEmptyPtrLock = base(Adds) + filter(NonEmpty) + type(Ptr) + lock(Lock)
func (it *Collection) AddsNonEmptyPtrLock(itemsPtr ...*string) *Collection {
    it.Lock()
    defer it.Unlock()
    for _, ptr := range itemsPtr {
        if ptr == nil { continue }
        s := *ptr
        if s == "" { continue }
        it.items = append(it.items, s)
    }
    return it
}
```

#### Master Suffix Reference Table

**Behavioral:**

| Suffix | Purpose | Example |
|--------|---------|---------|
| `*Lock` | Wraps with mutex | `Add` → `AddLock` |
| `*If` | Conditional on `is*` bool | `FmtDebug` → `FmtDebugIf` |
| `*LockIf` | Conditional locking | `Create` → `CreateLockIf` |
| `*Must` | Panics on error | `Deserialize` → `DeserializeMust` |
| (pair) | Opposite states | `IsValid` + `IsInvalid` |

**Filtering (string-specific):**

| Suffix | Purpose | Example |
|--------|---------|---------|
| `*NonEmpty` | Skips `""` | `Add` → `AddNonEmpty` |
| `*NonEmptyWhitespace` | Skips `""` + whitespace | `Add` → `AddNonEmptyWhitespace` |
| `*Trimmed*` | Trims then filters empty | `TrimmedEachWords` |
| `*Join` | Filter then join | `NonEmptyJoin` |

**Type modifiers:**

| Suffix | Purpose | Example |
|--------|---------|---------|
| `*Ptr` | Returns/accepts pointer with nil-safety | `Json` → `JsonPtr` |
| `ToPtr` | Value receiver → pointer | `(it Value) ToPtr() *Value` |
| `*Strings` | Variadic `...string` | `AddNonEmptyStrings` |
| `*Slice` | Accepts `[]T` | `AddNonEmptyStringsSlice` |
| `*Collection(s)` | Accepts `*Collection` | `AddCollection`, `AddCollections` |

**Result modifiers:**

| Suffix | Purpose | Example |
|--------|---------|---------|
| `*OrDefault` | Returns zero value if missing | `First` → `FirstOrDefault` |
| `*OrDefaultWith` | Custom fallback value | `FirstOrDefaultWith(slice, "N/A")` |
| `*New` | Returns new slice (no mutation) | `AppendLineNew`, `MergeNew` |

**Constructors:**

| Pattern | Purpose | Example |
|---------|---------|---------|
| `New*` | Factory | `NewCollection(cap)` |
| `*Using*` | From specific input | `UsingCap(n)` |
| `*From*` | Conversion | `FromSlice(parts)` |
| `ParseInjectUsingJson*` | JSON → existing struct | `ParseInjectUsingJsonMust` |
| `Serialize*` / `Deserialize*` | JSON round-trip | `SerializeMust()` |

**Suffix order** (mandatory): **Base + Filter + Type + Lock + If + Must**

**Rules**: (1) Name expresses behavior. (2) Bool param always first, uses `is*` prefix. (3) `*If` calls the unconditional version. (4) Each variant in its own file. (5) Delegate upward. (6) Suffix order is fixed — never rearrange. (7) `*Must` always panics. (8) `*OrDefault` returns zero; `*OrDefaultWith` accepts custom fallback. (9) `ToPtr` is value-receiver only.

### Method Writing: Pointer Variants (`*Ptr` Suffix)

When a method returns `T`, provide a `*Ptr` variant returning `*T`. When a checker accepts `T`, provide a `*Ptr` variant accepting `*T` with nil-safety.

```go
// Pattern 1: Return pointer — *Ptr calls value version, returns &result
func (it Version) Json() corejson.Result     { return corejson.New(it) }
func (it Version) JsonPtr() *corejson.Result  { return corejson.NewPtr(it) }

func (it *Version) ClonePtr() *Version {
    if it == nil { return nil }  // nil-safe
    clone := it.Clone()
    return &clone
}

// Pattern 2: Checker pointer — nil treated as empty/absent
func IsEmpty(str string) bool      { return str == "" }
func IsEmptyPtr(str *string) bool  { return str == nil || *str == "" }

func IsBlank(str string) bool      { ... }
func IsBlankPtr(s *string) bool    { return s == nil || IsBlank(*s) }

func IsDefined(str string) bool       { return !(str == "" || strings.TrimSpace(str) == "") }
func IsDefinedPtr(str *string) bool   { return !(str == nil || IsEmptyOrWhitespace(*str)) }

// Pattern 3: Identity conversion — ToPtr / NonPtr
func (it Variant) ToPtr() *Variant   { return &it }
func (it Version) NonPtr() Version   { return it }

// Pattern 4: Collection pointer — ListPtr, ValuePtr
func (it *Hashset[T]) List() []T      { ... }
func (it *Hashset[T]) ListPtr() *[]T  { list := it.List(); return &list }
```

| Suffix | When | Nil Handling |
|--------|------|--------------|
| `*Ptr` (return) | Caller needs `*T` | Pointer-receiver: check `it == nil` |
| `*Ptr` (accept) | Caller has `*T` | `nil` = empty/absent |
| `ToPtr` | Value → pointer | N/A (value receiver) |
| `NonPtr` | Pointer → value | Identity |
| `ClonePtr` | Deep copy as `*T` | `nil` → `nil` |

**File naming**: Each variant in its own file — `IsEmpty.go` / `IsEmptyPtr.go`, `Clone.go` / `ClonePtr.go`.

### Method Writing: `*Must` Suffix (Panic-on-Error)

```go
// Deserialize returns (*T, error). DeserializeMust panics on error.
func (it creator) Deserialize(bytes []byte) (*T, error) { ... }
func (it creator) DeserializeMust(bytes []byte) *T {
    result, err := it.Deserialize(bytes)
    errcore.HandleErr(err)
    return result
}

// SerializeMust — nil-guard then delegate to Serialize
func (it *TypedPayloadWrapper[T]) SerializeMust() []byte {
    if it == nil || it.Wrapper == nil {
        panic(defaulterr.NilResult)
    }
    bytes, err := it.Serialize()
    errcore.HandleErr(err)
    return bytes
}
```

**Rules**: (1) `*Must` always panics — never log, return a default, or swallow the error. (2) Use `errcore.HandleErr(err)` — not bare `panic(err)`. (3) The `*Must` variant calls the error-returning version — never duplicate logic. (4) File naming: `Deserialize.go` / `DeserializeMust.go`. (5) Combine with `*Ptr`: `ResultMust` returns `T`, `ResultPtrMust` returns `*T`.

### Method Writing: `*Slice` vs Variadic

```go
// Variadic is primary form
func (it *Collection) AddNonEmptyStrings(items ...string) *Collection { ... }

// *Slice companion accepts []string — delegates via spread
func (it *Collection) AddNonEmptyStringsSlice(items []string) *Collection {
    return it.AddNonEmptyStrings(items...)
}
```

**When**: Provide `*Slice` when method only has variadic `...T`. Skip if first param is already `[]T` (caller can spread).

### Method Writing: `*Or*` Fallback Pattern

**Fallback hierarchy:**

| Suffix | Fallback | Example |
|--------|----------|---------|
| `First` | Panics if empty | `it.items[0]` |
| `FirstOrDefault` | Returns zero value of `T` | `var zero T; return zero` |
| `FirstOrDefaultWith(fallback)` | Returns caller-provided fallback | `return fallback` |
| `GetOrDefault(key, fallback)` | Map fallback | `return defaultVal` |
| `CreateOrExisting(name)` | Create-or-retrieve; returns `(*T, bool)` | Cache/registry pattern |

```go
// First() panics. FirstOrDefault() returns zero. FirstOrDefaultWith() returns custom.
func (it *Collection[T]) First() T             { return it.items[0] }
func (it *Collection[T]) FirstOrDefault() T    { if it.IsEmpty() { var z T; return z }; return it.items[0] }
func FirstOrDefaultWith(s []string, d string) string { if len(s)==0 { return d }; return s[0] }

// GetOrDefault — map fallback
func (it *Hashmap[K, V]) GetOrDefault(key K, defaultVal V) V {
    val, found := it.items[key]
    if !found { return defaultVal }
    return val
}
```

#### `*OrExisting` Pattern

Used when creating-or-retrieving from a cache/registry:

```go
// CreateOrExisting returns an existing LazyRegex or creates a new one.
func (it *lazyRegexMap) CreateOrExisting(
    patternName string,
) (lazyRegex *LazyRegex, isExisting bool) {
    existing, found := it.items[patternName]
    if found {
        return existing, true
    }
    created := NewLazyRegex(patternName)
    it.items[patternName] = created
    return created, false
}

// CreateOrExistingLock — thread-safe variant (delegates to base).
func (it *lazyRegexMap) CreateOrExistingLock(
    patternName string,
) (*LazyRegex, bool) {
    it.Lock()
    defer it.Unlock()
    return it.CreateOrExisting(patternName)
}

// CreateOrExistingLockIf — combined: base + lock + if
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

**File naming**: `CreateOrExisting.go` / `CreateOrExistingLock.go` / `CreateOrExistingLockIf.go`.

```
First.go                  # First() T — panics
FirstOrDefault.go         # FirstOrDefault() T — zero value
FirstOrDefaultWith.go     # FirstOrDefaultWith(slice, default) T
GetOrDefault.go           # GetOrDefault(key, default) V
CreateOrExisting.go       # CreateOrExisting(name) (*T, bool)
CreateOrExistingLock.go   # CreateOrExistingLock(name) (*T, bool)
CreateOrExistingLockIf.go # CreateOrExistingLockIf(isLock, name) (*T, bool)
```

> **`*OrDefault` returns zero value. `*OrDefaultWith` accepts custom fallback. `*OrExisting` returns existing instance instead of creating new.**

#### Compound `*Or*` Naming in Filter Chains

When a method tries multiple filter strategies in sequence, name it using `Or` to chain the filter names.

> **`AOrB` means: try filter A first, fall back to filter B if A yields nothing.**

```go
// NonEmptyItemsOrNonWhitespace tries NonEmpty first;
// if all items are empty strings, falls back to NonWhitespace filtering.
func (it *Collection[T]) NonEmptyItemsOrNonWhitespace() []T {
    result := it.NonEmptyItems()
    if len(result) > 0 {
        return result
    }
    return it.NonWhitespaceItems()
}

// GetOrKeyOrDefault tries the primary key, then a fallback key,
// then returns the default value.
func (it *Hashmap[string, string]) GetOrKeyOrDefault(
    primaryKey string,
    fallbackKey string,
    defaultVal string,
) string {
    if val, ok := it.Get(primaryKey); ok { return val }
    if val, ok := it.Get(fallbackKey); ok { return val }
    return defaultVal
}
```

| Pattern | Meaning |
|---------|---------|
| `NonEmptyOrNonWhitespace` | Try non-empty filter, fall back to non-whitespace |
| `FirstNonEmptyOrDefault` | First non-empty item, or zero value |
| `GetOrKeyOrDefault` | Primary key → fallback key → default value |
| `TrimmedOrNonEmpty` | Try trimmed filter, fall back to non-empty |

**Rules**: (1) Each segment must be a real filter/fallback name. (2) Evaluation order matches reading order. (3) Delegate internally to standalone methods. (4) File uses full compound name: `NonEmptyItemsOrNonWhitespace.go`. (5) Max two `Or` segments.

### Deprecation Convention

```go
// Deprecated: Use EmptySlicePtr[any]() instead.
func EmptyAnysPtr() *[]any { return EmptySlicePtr[any]() }

// Deprecated: Use the built-in max() function (Go 1.21+).
func MaxByte(a, b byte) byte { ... }
```

**Rules**: (1) Always `// Deprecated: Use X instead.` (Go tooling recognizes this). (2) Name exact replacement. (3) Delegate to replacement — don't duplicate logic. (4) Keep the function — don't delete.

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

### Recipe: PayloadWrapper Creation & Deserialization

```go
// Empty wrapper
pw := corepayload.New.PayloadWrapper.Empty()

// From a single record (serializes to JSON internally)
pw, err := corepayload.New.PayloadWrapper.Create(
    "user-update", "42", "UpdateTask", "users", myStruct,
)

// From raw JSON bytes
pw, err := corepayload.New.PayloadWrapper.Deserialize(jsonBytes)

// Using BytesCreateInstruction
pw := corepayload.New.PayloadWrapper.UsingBytesCreateInstruction(&corepayload.BytesCreateInstruction{
    Name: "order", Identifier: "99", TaskTypeName: "Process",
    EntityType: "Order", Payloads: rawBytes,
})

// Deserialize payloads into a target struct
var order Order
err := pw.Deserialize(&order)       // returns error
pw.DeserializeMust(&order)          // panics on error

// Clone
cloned, err := pw.ClonePtr(true)    // deep clone
cloned, err := pw.ClonePtr(false)   // shallow clone
```

### Recipe: PayloadsCollection Usage

```go
// Create
col := corepayload.New.PayloadsCollection.Empty()
col := corepayload.New.PayloadsCollection.UsingCap(10)
col, err := corepayload.New.PayloadsCollection.Deserialize(jsonBytes)

// Mutate (fluent)
col.Add(wrapper).Adds(w1, w2).AddsPtr(ptrW1, ptrW2)
col.InsertAt(0, wrapper)
col.Reverse()

// Query
col.Length()                        // int
col.IsEmpty()                       // bool
col.First()                         // *PayloadWrapper (nil if empty)
col.Last()                          // *PayloadWrapper
col.FirstOrDefault()                // nil-safe
col.Skip(5)                         // []*PayloadWrapper
col.Take(3)                         // []*PayloadWrapper

// Filter
col.FirstById("42")                 // *PayloadWrapper
col.FirstByCategory("orders")       // *PayloadWrapper
col.FilterCategoryCollection("orders") // *PayloadsCollection
col.FilterEntityTypeCollection("Order") // *PayloadsCollection
col.Filter(func(pw *PayloadWrapper) (isTake, isBreak bool) { ... })

// Clone & Concat
cloned := col.Clone()               // value copy
col.ConcatNew(w1, w2)               // new collection
```

---

## coregeneric — Generic Data Structures API Reference

### Collection[T any]

Slice-backed collection with embedded `sync.Mutex`. Constraint: `T any`.

```go
// Construction via New Creator
col := coregeneric.New.Collection.String.Cap(10)
col := coregeneric.New.Collection.Int.Items(1, 2, 3)
col := coregeneric.New.Collection.Float64.Empty()

// Or via package-level functions
col := coregeneric.EmptyCollection[string]()
col := coregeneric.NewCollection[MyStruct](20)
col := coregeneric.CollectionFrom(existingSlice)   // no copy
col := coregeneric.CollectionClone(existingSlice)   // copies
```

**Mutation (fluent, returns `*Collection[T]`)**:

| Method | Description |
|--------|-------------|
| `Add(item)` | Append one item |
| `AddLock(item)` | Append with mutex |
| `Adds(items...)` | Append variadic |
| `AddsLock(items...)` | Variadic with mutex |
| `AddSlice([]T)` | Append from slice |
| `AddIf(bool, item)` | Conditional append |
| `AddIfMany(bool, items...)` | Conditional variadic |
| `AddFunc(func() T)` | Append function result |
| `AddCollection(*Collection[T])` | Merge another collection |
| `AddCollections(...*Collection[T])` | Merge multiple |
| `RemoveAt(index) bool` | Remove by index |
| `SortFunc(less)` | In-place sort |
| `Reverse()` | In-place reverse |
| `ConcatNew(items...)` | New collection = this + items |
| `Clone()` | Deep copy |

**Query**:

| Method | Returns | Description |
|--------|---------|-------------|
| `Length()` / `Count()` | `int` | Number of items |
| `LengthLock()` | `int` | Thread-safe length |
| `IsEmpty()` / `IsEmptyLock()` | `bool` | Empty check |
| `HasItems()` / `HasAnyItem()` | `bool` | Non-empty check |
| `HasIndex(i)` | `bool` | Bounds check |
| `Items()` | `[]T` | Underlying slice |
| `First()` / `Last()` | `T` | Panics if empty |
| `FirstOrDefault()` / `LastOrDefault()` | `T` | Zero-value if empty |
| `SafeAt(i)` | `T` | Zero-value if OOB |
| `Skip(n)` / `Take(n)` | `[]T` | Slice operations |
| `Filter(pred)` | `*Collection[T]` | New filtered collection |
| `CountFunc(pred)` | `int` | Count matching |
| `ForEach(fn)` | — | Iterate with index |
| `ForEachBreak(fn)` | — | Iterate with early exit |

**Iterators** (Go 1.23+ `iter` package):

```go
for i, item := range col.All() { ... }     // iter.Seq2[int, T]
for item := range col.Values() { ... }     // iter.Seq[T]
```

**Package-level generic functions** (`funcs.go`, `comparablefuncs.go`):

| Function | Constraint | Description |
|----------|-----------|-------------|
| `MapCollection(src, fn)` | `T→U` | Transform Collection[T] → Collection[U] |
| `FlatMapCollection(src, fn)` | `T→[]U` | Flatten-transform |
| `ReduceCollection(src, init, fn)` | `T→U` | Fold to single value |
| `GroupByCollection(src, keyFn)` | `K comparable` | Group into map[K]*Collection[T] |
| `ContainsFunc(src, pred)` | `T any` | Predicate search |
| `ContainsItem(src, item)` | `T comparable` | Direct equality search |
| `IndexOfFunc(src, pred)` / `IndexOfItem(src, item)` | — | Find index |
| `Distinct(src)` | `T comparable` | New deduped collection |
| `ContainsAll(src, items...)` | `T comparable` | All items present? |
| `ContainsAny(src, items...)` | `T comparable` | Any item present? |
| `RemoveItem(src, item)` | `T comparable` | Remove first occurrence |
| `RemoveAllItems(src, item)` | `T comparable` | Remove all occurrences |
| `ToHashset(src)` | `T comparable` | Convert to Hashset[T] |

### Hashset[T comparable]

Set backed by `map[T]bool` with embedded `sync.Mutex`.

```go
hs := coregeneric.New.Hashset.String.Empty()
hs := coregeneric.New.Hashset.Int.Cap(100)
hs := coregeneric.HashsetFrom([]string{"a", "b"})
hs := coregeneric.HashsetFromMap(existingMap)
```

| Method | Returns | Description |
|--------|---------|-------------|
| `Add(key)` | `*Hashset[T]` | Add item (fluent) |
| `AddBool(key)` | `bool` | Returns true if already existed |
| `AddLock(key)` | `*Hashset[T]` | Thread-safe add |
| `Adds(keys...)` | `*Hashset[T]` | Variadic add |
| `AddSlice([]T)` / `AddSliceLock([]T)` | `*Hashset[T]` | Add from slice |
| `AddIf(bool, key)` / `AddIfMany(bool, keys...)` | `*Hashset[T]` | Conditional |
| `AddHashsetItems(other)` | `*Hashset[T]` | Merge sets |
| `Has(key)` / `Contains(key)` / `ContainsLock(key)` | `bool` | Membership |
| `HasAll(keys...)` / `HasAny(keys...)` | `bool` | Bulk membership |
| `Remove(key)` / `RemoveLock(key)` | `bool` | Remove (returns existed) |
| `Length()` / `LengthLock()` | `int` | Size |
| `IsEmpty()` / `IsEmptyLock()` | `bool` | Empty check |
| `List()` | `[]T` | All keys as slice |
| `Map()` | `map[T]bool` | Underlying map |
| `Collection()` | `*Collection[T]` | Convert to Collection |
| `Resize(cap)` | `*Hashset[T]` | Grow internal map |
| `IsEquals(other)` | `bool` | Set equality |

**Iterators**:

```go
for item, _ := range hs.All() { ... }   // iter.Seq2[T, bool]
for item := range hs.Values() { ... }   // iter.Seq[T]
```

### Hashmap[K comparable, V any]

Map wrapper with embedded `sync.Mutex`.

```go
hm := coregeneric.New.Hashmap.StringString.Cap(10)
hm := coregeneric.New.Hashmap.StringAny.Empty()
hm := coregeneric.HashmapFrom(existingMap)     // no copy
hm := coregeneric.HashmapClone(existingMap)    // copies
```

| Method | Returns | Description |
|--------|---------|-------------|
| `Set(key, val)` | `bool` (isNew) | Add or update |
| `SetLock(key, val)` | `*Hashmap` | Thread-safe set |
| `Get(key)` | `(V, bool)` | Retrieve |
| `GetOrDefault(key, default)` | `V` | Retrieve with fallback |
| `GetLock(key)` | `(V, bool)` | Thread-safe get |
| `Has(key)` / `Contains(key)` / `ContainsLock(key)` | `bool` | Key check |
| `IsKeyMissing(key)` | `bool` | Negated Has |
| `Remove(key)` / `RemoveLock(key)` | `bool` | Delete (returns existed) |
| `AddOrUpdateMap(map[K]V)` | `*Hashmap` | Merge from raw map |
| `AddOrUpdateHashmap(other)` | `*Hashmap` | Merge from Hashmap |
| `Keys()` | `[]K` | All keys |
| `Values()` | `[]V` | All values |
| `Map()` | `map[K]V` | Underlying map |
| `ForEach(fn)` / `ForEachBreak(fn)` | — | Iterate |
| `ConcatNew(others...)` | `*Hashmap` | New merged hashmap |
| `Clone()` | `*Hashmap` | Copy |
| `Length()` / `LengthLock()` | `int` | Size |
| `IsEmpty()` / `IsEmptyLock()` | `bool` | Empty check |
| `IsEquals(other)` | `bool` | Key equality check |

**Iterators**:

```go
for k, v := range hm.All() { ... }      // iter.Seq2[K, V]
for k := range hm.Keys() { ... }        // iter.Seq[K]  (via HashmapIter.go)
```

### SimpleSlice[T any]

Thin generic slice wrapper. Uses `[T any]` constraint; for ordered operations (Sort, Min, Max), use package-level `SortSimpleSlice[T cmp.Ordered]()` functions.

```go
ss := coregeneric.New.SimpleSlice.String.Empty()
ss := coregeneric.New.SimpleSlice.Int.Cap(100)
ss := coregeneric.New.SimpleSlice.Float64.Items(1.0, 2.5, 3.7)
ss := coregeneric.New.SimpleSlice.String.From(existingSlice)   // no copy
ss := coregeneric.New.SimpleSlice.String.Clone(existingSlice)  // copies
```

**Package-level constructors** (for custom types):

```go
ss := coregeneric.EmptySimpleSlice[MyType]()
ss := coregeneric.NewSimpleSlice[MyType](capacity)
ss := coregeneric.SimpleSliceFrom[MyType](items)    // no copy
ss := coregeneric.SimpleSliceClone[MyType](items)   // copies
```

**API quick reference:**

| Method | Returns | Description |
|--------|---------|-------------|
| `Add(item)` | `*SimpleSlice` | Append single item |
| `Adds(items...)` | `*SimpleSlice` | Append variadic items |
| `AddSlice([]T)` | `*SimpleSlice` | Append from slice |
| `AddIf(bool, item)` | `*SimpleSlice` | Conditional append |
| `AddsIf(bool, items...)` | `*SimpleSlice` | Conditional variadic append |
| `AddFunc(func() T)` | `*SimpleSlice` | Append function result |
| `InsertAt(index, item)` | `*SimpleSlice` | Insert at position |
| `First()` / `Last()` | `T` | First/last element (panics if empty) |
| `FirstOrDefault()` / `LastOrDefault()` | `T` | First/last or zero value |
| `Skip(n)` / `Take(n)` | `[]T` | Subsequence helpers |
| `Length()` / `Count()` | `int` | Number of elements |
| `IsEmpty()` / `HasAnyItem()` / `HasItems()` | `bool` | Emptiness checks |
| `HasIndex(i)` | `bool` | Valid index check |
| `LastIndex()` | `int` | Index of last element |
| `Items()` | `[]T` | Underlying slice |
| `Filter(predicate)` | `*SimpleSlice` | New filtered slice |
| `CountFunc(predicate)` | `int` | Count matching items |
| `ForEach(fn)` | — | Iterate with index |
| `Clone()` | `*SimpleSlice` | Deep copy |
| `String()` | `string` | String representation |

**Iterators (Go 1.23+):**

```go
for i, item := range ss.All() { ... }    // iter.Seq2[int, T]
for item := range ss.Values() { ... }    // iter.Seq[T]
```

---

### LinkedList[T any]

Generic singly-linked list with head/tail pointers and embedded `sync.Mutex`.

```go
ll := coregeneric.New.LinkedList.String.Empty()
ll := coregeneric.New.LinkedList.Int.From([]int{1, 2, 3})
ll := coregeneric.New.LinkedList.Float64.Items(1.0, 2.5, 3.7)
```

**Package-level constructors:**

```go
ll := coregeneric.EmptyLinkedList[MyType]()
ll := coregeneric.LinkedListFrom[MyType](items)
```

**API quick reference:**

| Method | Returns | Description |
|--------|---------|-------------|
| `Add(item)` / `Push(item)` / `PushBack(item)` | `*LinkedList` | Append to back |
| `AddLock(item)` | `*LinkedList` | Thread-safe append |
| `Adds(items...)` | `*LinkedList` | Append variadic |
| `AddSlice([]T)` | `*LinkedList` | Append from slice |
| `AddIf(bool, item)` | `*LinkedList` | Conditional append |
| `AddsIf(bool, items...)` | `*LinkedList` | Conditional variadic append |
| `AddFunc(func() T)` | `*LinkedList` | Append function result |
| `AddFront(item)` / `PushFront(item)` | `*LinkedList` | Prepend to front |
| `AppendNode(node)` | `*LinkedList` | Append existing node |
| `AppendChainOfNodes(head)` | `*LinkedList` | Append a chain of nodes |
| `First()` / `Last()` | `T` | First/last element (panics if empty) |
| `FirstOrDefault()` / `LastOrDefault()` | `T` | First/last or zero value |
| `Head()` / `Tail()` | `*LinkedListNode` | Head/tail node access |
| `IndexAt(i)` | `*LinkedListNode` | Node at index — O(n) |
| `Length()` / `LengthLock()` | `int` | Size (with/without mutex) |
| `IsEmpty()` / `IsEmptyLock()` | `bool` | Empty check |
| `HasItems()` | `bool` | Has at least one item |
| `Items()` | `[]T` | Collect all to slice |
| `Collection()` | `*Collection[T]` | Convert to Collection |
| `ForEach(fn)` | — | Iterate with index |
| `ForEachBreak(fn)` | — | Iterate; stop if fn returns true |
| `String()` | `string` | String representation |

**LinkedListNode[T] methods:**

| Method | Returns | Description |
|--------|---------|-------------|
| `Element` | `T` | The stored value (public field) |
| `HasNext()` | `bool` | Has a next node |
| `Next()` | `*LinkedListNode` | Next node |
| `EndOfChain()` | `(*LinkedListNode, int)` | Last node + chain length |
| `Clone()` | `*LinkedListNode` | Copy (without chain) |
| `ListPtr()` | `*[]T` | Collect chain to slice |

**Iterators (Go 1.23+):**

```go
for i, value := range ll.All() { ... }    // iter.Seq2[int, T]
for value := range ll.Values() { ... }    // iter.Seq[T]
```

---

### Pair[L any, R any]

Generic two-value container with validity tracking. Generalizes `corestr.LeftRight`.

```go
// Direct construction
pair := coregeneric.NewPair("key", 42)              // Pair[string, int]
pair := coregeneric.NewPairOf(10, 20)                // Pair[int, int] (same-type shortcut)

// Via New Creator
pair := coregeneric.New.Pair.StringString("a", "b")
pair := coregeneric.New.Pair.StringInt("name", 42)
pair := coregeneric.New.Pair.StringAny("key", value)
pair := coregeneric.New.Pair.Any("left", "right")

// Invalid pairs
pair := coregeneric.InvalidPair[string, int]("reason")
pair := coregeneric.InvalidPairNoMessage[string, int]()

// String splitting
pair := coregeneric.New.Pair.Split("key=value", "=")           // Left="key", Right="value"
pair := coregeneric.New.Pair.SplitTrimmed(" key = value ", "=") // trimmed
pair := coregeneric.New.Pair.SplitFull("a:b:c:d", ":")          // Left="a", Right="b:c:d"
pair := coregeneric.New.Pair.FromSlice([]string{"a", "b"})

// Number division
pair := coregeneric.New.Pair.DivideInt(11)                     // {5, 6}
pair := coregeneric.New.Pair.DivideIntWeighted(100, 0.3)       // {30, 70}
```

**API quick reference:**

| Method | Returns | Description |
|--------|---------|-------------|
| `Left` / `Right` | `L`, `R` | Public fields |
| `IsValid` / `Message` | `bool`, `string` | Validity tracking fields |
| `Values()` | `(L, R)` | Both values |
| `IsInvalid()` | `bool` | Nil-safe invalid check |
| `HasMessage()` | `bool` | Has non-empty message |
| `Clone()` | `*Pair` | Shallow copy |
| `IsEqual(other)` | `bool` | Compare via `fmt.Sprintf` |
| `String()` | `string` | Formatted representation |
| `Clear()` / `Dispose()` | — | Reset to zero values |

---

### Triple[A any, B any, C any]

Generic three-value container with validity tracking. Generalizes `corestr.LeftMiddleRight`.

```go
// Direct construction
triple := coregeneric.NewTriple("left", 42, true)        // Triple[string, int, bool]
triple := coregeneric.NewTripleOf(1, 2, 3)               // Triple[int, int, int] (same-type)

// Via New Creator
triple := coregeneric.New.Triple.StringStringString("a", "b", "c")
triple := coregeneric.New.Triple.Any("a", 42, true)

// Invalid triples
triple := coregeneric.InvalidTriple[string, int, bool]("reason")
triple := coregeneric.InvalidTripleNoMessage[string, int, bool]()

// String splitting
triple := coregeneric.New.Triple.Split("a.b.c", ".")            // Left="a", Middle="b", Right="c"
triple := coregeneric.New.Triple.SplitN("a:b:c:d:e", ":")       // Left="a", Middle="b", Right="c:d:e"
triple := coregeneric.New.Triple.SplitTrimmed(" a . b . c ", ".")
triple := coregeneric.New.Triple.FromSlice([]string{"a", "b", "c"})

// Number division
triple := coregeneric.New.Triple.DivideInt(10)                  // {3, 3, 4}
triple := coregeneric.New.Triple.DivideIntWeighted(100, 0.2, 0.3) // {20, 30, 50}
```

**API quick reference:**

| Method | Returns | Description |
|--------|---------|-------------|
| `Left` / `Middle` / `Right` | `A`, `B`, `C` | Public fields |
| `IsValid` / `Message` | `bool`, `string` | Validity tracking fields |
| `Values()` | `(A, B, C)` | All three values |
| `IsInvalid()` | `bool` | Nil-safe invalid check |
| `HasMessage()` | `bool` | Has non-empty message |
| `Clone()` | `*Triple` | Shallow copy |
| `IsEqual(other)` | `bool` | Compare via `fmt.Sprintf` |
| `String()` | `string` | Formatted representation |
| `Clear()` / `Dispose()` | — | Reset to zero values |

---

### Package-Level Functional Helpers (`funcs.go` + `orderedfuncs.go`)

Go does not allow generic methods with additional type parameters on a generic receiver.
These **package-level functions** bridge that gap, enabling cross-type transformations, aggregations, and ordered operations on `Collection[T]`, `SimpleSlice[T]`, `Hashset[T]`, and `Hashmap[K,V]`.

#### Cross-Type Transformations (`funcs.go`)

| Function | Signature | Description |
|----------|-----------|-------------|
| `MapCollection` | `[T,U any](source, mapper) → *Collection[U]` | Transform each item to a different type |
| `FlatMapCollection` | `[T,U any](source, mapper) → *Collection[U]` | Map then flatten slices |
| `ReduceCollection` | `[T,U any](source, initial, reducer) → U` | Fold into a single value |
| `GroupByCollection` | `[T any, K comparable](source, keyFunc) → map[K]*Collection[T]` | Group items by key |
| `ContainsFunc` | `[T any](source, predicate) → bool` | Predicate search (non-comparable T) |
| `IndexOfFunc` | `[T any](source, predicate) → int` | Index of first match, or -1 |
| `ContainsItem` | `[T comparable](source, item) → bool` | Equality check for comparable T |
| `IndexOfItem` | `[T comparable](source, item) → int` | Index of first occurrence, or -1 |
| `Distinct` | `[T comparable](source) → *Collection[T]` | Remove duplicates, preserve order |
| `MapSimpleSlice` | `[T,U any](source, mapper) → *SimpleSlice[U]` | Transform SimpleSlice items |

```go
// MapCollection — transform users to names
names := coregeneric.MapCollection(users, func(u User) string { return u.Name })

// FlatMapCollection — flatten tags from posts
allTags := coregeneric.FlatMapCollection(posts, func(p Post) []string { return p.Tags })

// ReduceCollection — sum prices
total := coregeneric.ReduceCollection(prices, 0, func(acc int, p Price) int { return acc + p.Amount })

// GroupByCollection — group by department
groups := coregeneric.GroupByCollection(employees, func(e Employee) string { return e.Dept })

// Distinct — deduplicate
unique := coregeneric.Distinct(ids)

// ContainsFunc / IndexOfFunc — non-comparable search
found := coregeneric.ContainsFunc(items, func(it Item) bool { return it.ID == targetID })
idx   := coregeneric.IndexOfFunc(items, func(it Item) bool { return it.Name == "x" })

// ContainsItem / IndexOfItem — comparable search
has := coregeneric.ContainsItem(names, "Alice")
pos := coregeneric.IndexOfItem(nums, 42)

// MapSimpleSlice — transform SimpleSlice
lengths := coregeneric.MapSimpleSlice(words, func(w string) int { return len(w) })
```

#### Ordered Operations (`orderedfuncs.go`)

Require `T` (or `K`/`V`) to satisfy `cmp.Ordered`. Use `slices.Sort` / `slices.Min` / `slices.Max` internally.

**Collection[T] ordered helpers:**

| Function | Signature | Description |
|----------|-----------|-------------|
| `SortCollection` | `[T cmp.Ordered](source) → *Collection[T]` | In-place ascending sort |
| `SortCollectionDesc` | `[T cmp.Ordered](source) → *Collection[T]` | In-place descending sort |
| `MinCollection` / `MaxCollection` | `[T cmp.Ordered](source) → T` | Min/max (panics if empty) |
| `MinCollectionOrDefault` / `MaxCollectionOrDefault` | `[T cmp.Ordered](source, defVal) → T` | Min/max with fallback |
| `IsSortedCollection` | `[T cmp.Ordered](source) → bool` | Check ascending order |
| `SumCollection` | `[T cmp.Ordered](source) → T` | Sum all elements |
| `ClampCollection` | `[T cmp.Ordered](source, min, max) → *Collection[T]` | Clamp values in-place |

**SimpleSlice[T] ordered helpers:**

| Function | Signature | Description |
|----------|-----------|-------------|
| `SortSimpleSlice` | `[T cmp.Ordered](source) → *SimpleSlice[T]` | In-place ascending sort |
| `SortSimpleSliceDesc` | `[T cmp.Ordered](source) → *SimpleSlice[T]` | In-place descending sort |
| `MinSimpleSlice` / `MaxSimpleSlice` | `[T cmp.Ordered](source) → T` | Min/max (panics if empty) |
| `SumSimpleSlice` | `[T cmp.Ordered](source) → T` | Sum all elements |

**Hashset[T] ordered helpers:**

| Function | Signature | Description |
|----------|-----------|-------------|
| `SortedListHashset` / `SortedListDescHashset` | `[T cmp.Ordered](source) → []T` | Sorted slice from set |
| `SortedCollectionHashset` | `[T cmp.Ordered](source) → *Collection[T]` | Sorted Collection from set |
| `MinHashset` / `MaxHashset` | `[T cmp.Ordered](source) → T` | Min/max (panics if empty) |
| `MinHashsetOrDefault` / `MaxHashsetOrDefault` | `[T cmp.Ordered](source, defVal) → T` | Min/max with fallback |

**Hashmap[K,V] ordered helpers:**

| Function | Signature | Description |
|----------|-----------|-------------|
| `SortedKeysHashmap` / `SortedKeysDescHashmap` | `[K cmp.Ordered, V any](source) → []K` | Sorted keys |
| `MinKeyHashmap` / `MaxKeyHashmap` | `[K cmp.Ordered, V any](source) → K` | Min/max key |
| `MinKeyHashmapOrDefault` / `MaxKeyHashmapOrDefault` | `[K cmp.Ordered, V any](source, defVal) → K` | Min/max key with fallback |
| `SortedValuesHashmap` | `[K comparable, V cmp.Ordered](source) → []V` | Sorted values |
| `MinValueHashmap` / `MaxValueHashmap` | `[K comparable, V cmp.Ordered](source) → V` | Min/max value |
| `MinValueHashmapOrDefault` / `MaxValueHashmapOrDefault` | `[K comparable, V cmp.Ordered](source, defVal) → V` | Min/max value with fallback |

```go
// Sort + min/max on Collection
coregeneric.SortCollection(scores)           // in-place ascending
best := coregeneric.MaxCollection(scores)    // panics if empty
safe := coregeneric.MinCollectionOrDefault(scores, 0)

// Sort SimpleSlice
coregeneric.SortSimpleSlice(names)           // in-place ascending
coregeneric.SortSimpleSliceDesc(names)       // in-place descending

// Sum
total := coregeneric.SumCollection(amounts)
ssTotal := coregeneric.SumSimpleSlice(values)

// Hashset → sorted slice
sorted := coregeneric.SortedListHashset(tagSet)

// Hashmap sorted keys/values
keys := coregeneric.SortedKeysHashmap(config)
minKey := coregeneric.MinKeyHashmapOrDefault(config, "")
```

> **Why package-level?** Go prohibits `func (c *Collection[T]) Map[U any](...)` — methods cannot introduce new type parameters. Package-level functions like `MapCollection[T, U]` work around this limitation.

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
