# Enum Authoring Guide — Reusing `enumimpl` & `enuminf`

## Goal

This spec explains how to create a new enum package that matches the existing core style, reuses the shared enum building blocks, and is easy for another AI or engineer to extend safely.

Use this guide when you want a package like `reqtype`, `ostype`, or `enums/versionindexes`.

---

## Table of Contents

1. [Architecture Overview](#architecture-overview)
2. [Available Backing Types](#available-backing-types)
3. [Interface Hierarchy](#interface-hierarchy)
4. [Byte Enum — Full Pattern](#byte-enum--full-pattern)
5. [Int8 Enum — Full Pattern](#int8-enum--full-pattern)
6. [Int16 Enum — Full Pattern](#int16-enum--full-pattern)
7. [Int32 Enum — Full Pattern](#int32-enum--full-pattern)
8. [Alias-Aware Enum Pattern](#alias-aware-enum-pattern)
9. [Explicit Non-Contiguous Values Pattern](#explicit-non-contiguous-values-pattern)
10. [Case-Insensitive Parsing](#case-insensitive-parsing)
11. [Formula Rule — Safe vs Unsafe](#formula-rule--safe-vs-unsafe)
12. [AI Authoring Checklist](#ai-authoring-checklist)

---

## Architecture Overview

```
enumimpl.New                          ← factory entry point (singleton)
  ├── .BasicByte   → *BasicByte      ← byte-backed enum impl
  ├── .BasicInt8   → *BasicInt8      ← int8-backed enum impl
  ├── .BasicInt16  → *BasicInt16     ← int16-backed enum impl
  ├── .BasicInt32  → *BasicInt32     ← int32-backed enum impl
  ├── .BasicUInt16 → *BasicUInt16    ← uint16-backed enum impl
  └── .BasicString → *BasicString    ← string-backed enum impl

All number variants embed numberEnumBase which provides:
  TypeName(), RangeNamesCsv(), MinMaxAny(), MinInt(), MaxInt(),
  MinValueString(), MaxValueString(), AllNameValues(),
  IntegerEnumRanges(), RangesDynamicMap(), StringRanges(),
  StringRangesPtr(), Format(), OnlySupportedErr(), OnlySupportedMsgErr(),
  RangesInvalidMessage(), RangesInvalidErr()

Each typed variant adds:
  ToEnumString(T), ToEnumJsonBytes(T), UnmarshallToValue(bool,[]byte)(T,error),
  Min(), Max(), Ranges(), Hashmap(), IsValidRange(T), EnumType()
```

---

## Available Backing Types

| Type    | Go Type  | Creator                    | Impl Struct      | `EnumType()` returns    |
|---------|----------|----------------------------|------------------|-------------------------|
| Byte    | `byte`   | `enumimpl.New.BasicByte`   | `*BasicByte`     | `enumtype.Byte`         |
| Int8    | `int8`   | `enumimpl.New.BasicInt8`   | `*BasicInt8`     | `enumtype.Integer8`     |
| Int16   | `int16`  | `enumimpl.New.BasicInt16`  | `*BasicInt16`    | `enumtype.Integer16`    |
| Int32   | `int32`  | `enumimpl.New.BasicInt32`  | `*BasicInt32`    | `enumtype.Integer32`    |
| UInt16  | `uint16` | `enumimpl.New.BasicUInt16` | `*BasicUInt16`   | `enumtype.UnsignedInteger16` |
| String  | `string` | `enumimpl.New.BasicString` | `*BasicString`   | `enumtype.String`       |

Choose the **smallest type** that fits your value range.

---

## Interface Hierarchy

Understanding which interfaces your enum must satisfy:

### `enuminf.BaseEnumer` (required for all enums)

```go
type BaseEnumer interface {
    // from enumNameStinger (unexported)
    String() string                          // human-readable, delegates to ToEnumString

    // SimpleEnumer
    Name() string                            // enum member name ("Ready")
    TypeName() string                        // full type name ("status.Status")
    ValueByte() byte                         // raw byte value
    IsValid() bool
    IsInvalid() bool

    // NameValuer
    NameValue() string                       // "Ready (2)" format

    // IsNameEqualer
    IsNameEqual(name string) bool

    // IsAnyNameOfChecker
    IsAnyNamesOf(names ...string) bool

    // ToNumberStringer
    ToNumberString() string                  // value as number string: "2"

    // IsValidInvalidChecker
    IsValid() bool
    IsInvalid() bool

    // BasicEnumValuer — ALL these must be implemented
    ValueByte() byte
    ValueInt() int
    ValueInt8() int8
    ValueInt16() int16
    ValueUInt16() uint16
    ValueInt32() int32
    ValueString() string                     // number as string, NOT name

    // RangeNamesCsvGetter
    RangeNamesCsv() string

    // corejson.JsonMarshaller
    MarshalJSON() ([]byte, error)
    UnmarshalJSON(data []byte) error
}
```

### `enuminf.BasicEnumer` (extends BaseEnumer)

```go
type BasicEnumer interface {
    BaseEnumer
    EnumFormatter                            // Format(format string) string
    MinMaxAny() (min, max any)
    MinValueString() string
    MaxValueString() string
    MaxInt() int
    MinInt() int
    RangesDynamicMapGetter                   // RangesDynamicMap() map[string]any
    AllNameValues() []string                 // ["Invalid (0)", "Ready (1)", ...]
    OnlySupportedNamesErrorer                // OnlySupportedErr(...string) error
    IntegerEnumRangesGetter                  // IntegerEnumRanges() []int
    EnumType() EnumTyper
}
```

### `enuminf.StandardEnumer` (extends BasicEnumer)

```go
type StandardEnumer interface {
    BasicEnumer
    StringRangesGetter                       // StringRanges(), StringRangesPtr()
    RangeValidateChecker                     // IsValidRange(), IsInvalidRange(), RangesInvalidMessage(), RangesInvalidErr()
    corejson.JsonContractsBinder
}
```

### Type-Specific Interfaces (`enuminf.BasicByteEnumer`, etc.)

Each backing type has a dedicated interface with typed min/max/ranges/unmarshal:

```go
// enuminf.BasicByteEnumer
type BasicByteEnumer interface {
    UnmarshallEnumToValueByter               // UnmarshallEnumToValue([]byte) (byte, error)
    MaxByte() byte
    MinByte() byte
    ValueByte() byte
    RangesByte() []byte
}

// enuminf.BasicInt8Enumer
type BasicInt8Enumer interface {
    UnmarshallEnumToValueInt8([]byte) (int8, error)
    MaxInt8() int8
    MinInt8() int8
    ValueInt8() int8
    RangesInt8() []int8
    ToEnumString(input int8) string
}

// enuminf.BasicInt16Enumer
type BasicInt16Enumer interface {
    UnmarshallEnumToValueInt16([]byte) (int16, error)
    MaxInt16() int16
    MinInt16() int16
    ValueInt16() int16
    RangesInt16() []int16
    ToEnumString(input int16) string
}

// enuminf.BasicInt32Enumer
type BasicInt32Enumer interface {
    UnmarshallEnumToValueInt32([]byte) (int32, error)
    MaxInt32() int32
    MinInt32() int32
    ValueInt32() int32
    RangesInt32() []int32
    ToEnumString(input int32) string
}
```

---

## Package Shape

```text
mypackage/
├── MyEnum.go         # enum type + methods
├── vars.go           # Ranges, RangesMap, BasicEnumImpl
├── consts.go         # package constants if needed
└── readme.md         # package overview
```

For large enums, split by responsibility:
- `MyEnum.naming.go` — Name, String, Format, NameValue
- `MyEnum.json.go` — MarshalJSON, UnmarshalJSON, UnmarshallEnumToValue
- `MyEnum.checkers.go` — IsValid, IsFailed, domain checkers
- `MyEnum.values.go` — Value*, Min*, Max*, Ranges*

---

## Byte Enum — Full Pattern

### `consts.go`

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

### `vars.go`

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

    RangesMap = map[string]Status{
        "Invalid": Invalid,
        "Pending": Pending,
        "Ready":   Ready,
        "Failed":  Failed,
    }

    BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
        reflectinternal.TypeName(Invalid),
        Ranges[:],
    )
)
```

### `Status.go` — Complete Method Set

```go
package status

import "github.com/alimtvnetwork/core/coreinterface/enuminf"

// ── Value accessors (BasicEnumValuer) ──────────────────────────

func (it Status) Value() byte       { return byte(it) }
func (it Status) ValueByte() byte   { return byte(it) }
func (it Status) ValueInt() int     { return int(it) }
func (it Status) ValueInt8() int8   { return int8(it) }
func (it Status) ValueInt16() int16 { return int16(it) }
func (it Status) ValueUInt16() uint16 { return uint16(it) }
func (it Status) ValueInt32() int32 { return int32(it) }
func (it Status) ValueString() string { return BasicEnumImpl.ToNumberString(it.Value()) }

// ── Naming (enumNameStinger, SimpleEnumer) ─────────────────────

func (it Status) Name() string     { return BasicEnumImpl.ToEnumString(it.Value()) }
func (it Status) String() string   { return BasicEnumImpl.ToEnumString(it.Value()) }
func (it Status) TypeName() string { return BasicEnumImpl.TypeName() }
func (it Status) NameValue() string {
    return BasicEnumImpl.NameWithValue(it.Value())
}
func (it Status) ToNumberString() string {
    return BasicEnumImpl.ToNumberString(it.Value())
}

// ── Equality & matching ────────────────────────────────────────

func (it Status) IsNameEqual(name string) bool {
    return it.Name() == name
}

func (it Status) IsAnyNamesOf(names ...string) bool {
    currentName := it.Name()
    for _, n := range names {
        if n == currentName {
            return true
        }
    }
    return false
}

func (it Status) IsByteValueEqual(value byte) bool {
    return byte(it) == value
}

func (it Status) IsAnyValuesEqual(anyByteValues ...byte) bool {
    v := byte(it)
    for _, b := range anyByteValues {
        if v == b {
            return true
        }
    }
    return false
}

func (it Status) IsValueEqual(value byte) bool {
    return byte(it) == value
}

// ── Valid/Invalid ──────────────────────────────────────────────

func (it Status) IsValid() bool   { return it != Invalid }
func (it Status) IsInvalid() bool { return it == Invalid }

// ── Range info (BasicEnumer) ───────────────────────────────────

func (it Status) RangeNamesCsv() string              { return BasicEnumImpl.RangeNamesCsv() }
func (it Status) MinMaxAny() (min, max any)          { return BasicEnumImpl.MinMaxAny() }
func (it Status) MinValueString() string             { return BasicEnumImpl.MinValueString() }
func (it Status) MaxValueString() string             { return BasicEnumImpl.MaxValueString() }
func (it Status) MaxInt() int                        { return BasicEnumImpl.MaxInt() }
func (it Status) MinInt() int                        { return BasicEnumImpl.MinInt() }
func (it Status) RangesDynamicMap() map[string]any   { return BasicEnumImpl.RangesDynamicMap() }
func (it Status) AllNameValues() []string            { return BasicEnumImpl.AllNameValues() }
func (it Status) IntegerEnumRanges() []int           { return BasicEnumImpl.IntegerEnumRanges() }

// ── OnlySupportedNamesErrorer ──────────────────────────────────

func (it Status) OnlySupportedErr(names ...string) error {
    return BasicEnumImpl.OnlySupportedErr(names...)
}
func (it Status) OnlySupportedMsgErr(message string, names ...string) error {
    return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

// ── Format (EnumFormatter) ─────────────────────────────────────
// Format string keys: {type-name}, {name}, {value}
// Example: "Enum of {type-name} - {name} - {value}"
//       →  "Enum of status.Status - Ready - 2"

func (it Status) Format(format string) string {
    return BasicEnumImpl.Format(format, it.Value())
}

// ── Type-specific: BasicByteEnumer ─────────────────────────────

func (it Status) MaxByte() byte    { return BasicEnumImpl.Max() }
func (it Status) MinByte() byte    { return BasicEnumImpl.Min() }
func (it Status) RangesByte() []byte { return BasicEnumImpl.Ranges() }

// ── Range validation (StandardEnumer) ──────────────────────────

func (it Status) IsValidRange() bool          { return BasicEnumImpl.IsValidRange(it.Value()) }
func (it Status) IsInvalidRange() bool        { return !it.IsValidRange() }
func (it Status) RangesInvalidMessage() string { return BasicEnumImpl.RangesInvalidMessage() }
func (it Status) RangesInvalidErr() error     { return BasicEnumImpl.RangesInvalidErr() }

// ── String ranges (StandardEnumer) ─────────────────────────────

func (it Status) StringRanges() []string    { return BasicEnumImpl.StringRanges() }
func (it Status) StringRangesPtr() []string { return BasicEnumImpl.StringRangesPtr() }

// ── JSON marshalling ───────────────────────────────────────────

func (it Status) MarshalJSON() ([]byte, error) {
    return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}

func (it *Status) UnmarshalJSON(data []byte) error {
    val, err := it.UnmarshallEnumToValue(data)
    if err == nil {
        *it = Status(val)
    }
    return err
}

func (it Status) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
    return BasicEnumImpl.UnmarshallToValue(true, jsonUnmarshallingValue)
}

// ── EnumType ───────────────────────────────────────────────────

func (it Status) EnumType() enuminf.EnumTyper {
    return BasicEnumImpl.EnumType()
}

// ── Domain-specific checkers (custom per enum) ─────────────────

func (it Status) IsPending() bool { return it == Pending }
func (it Status) IsReady() bool   { return it == Ready }
func (it Status) IsFailed() bool  { return it == Failed }
```

---

## Int8 Enum — Full Pattern

Use `int8` when you need more than 255 values or want signed range semantics with a small footprint.

### `consts.go`

```go
package severity

type Severity int8

const (
    Unknown  Severity = iota
    Low
    Medium
    High
    Critical
)
```

### `vars.go`

```go
package severity

import (
    "github.com/alimtvnetwork/core/coreimpl/enumimpl"
    "github.com/alimtvnetwork/core/internal/reflectinternal"
)

var (
    Ranges = [...]string{
        Unknown:  "Unknown",
        Low:      "Low",
        Medium:   "Medium",
        High:     "High",
        Critical: "Critical",
    }

    BasicEnumImpl = enumimpl.New.BasicInt8.UsingTypeSlice(
        reflectinternal.TypeName(Unknown),
        Ranges[:],
    )
)
```

### `Severity.go` — Complete Method Set

```go
package severity

import "github.com/alimtvnetwork/core/coreinterface/enuminf"

// ── Value accessors (BasicEnumValuer) ──────────────────────────

func (it Severity) Value() int8       { return int8(it) }
func (it Severity) ValueByte() byte   { return byte(it) }
func (it Severity) ValueInt() int     { return int(it) }
func (it Severity) ValueInt8() int8   { return int8(it) }
func (it Severity) ValueInt16() int16 { return int16(it) }
func (it Severity) ValueUInt16() uint16 { return uint16(it) }
func (it Severity) ValueInt32() int32 { return int32(it) }
func (it Severity) ValueString() string { return BasicEnumImpl.ToNumberString(it.Value()) }

// ── Naming ─────────────────────────────────────────────────────

func (it Severity) Name() string     { return BasicEnumImpl.ToEnumString(it.Value()) }
func (it Severity) String() string   { return BasicEnumImpl.ToEnumString(it.Value()) }
func (it Severity) TypeName() string { return BasicEnumImpl.TypeName() }
func (it Severity) NameValue() string { return BasicEnumImpl.NameWithValue(it.Value()) }
func (it Severity) ToNumberString() string { return BasicEnumImpl.ToNumberString(it.Value()) }

// ── Equality & matching ────────────────────────────────────────

func (it Severity) IsNameEqual(name string) bool { return it.Name() == name }
func (it Severity) IsAnyNamesOf(names ...string) bool {
    n := it.Name()
    for _, name := range names { if name == n { return true } }
    return false
}
func (it Severity) IsValueEqual(value int8) bool { return int8(it) == value }
func (it Severity) IsAnyValuesEqual(anyValues ...int8) bool {
    v := int8(it)
    for _, val := range anyValues { if v == val { return true } }
    return false
}

// ── Valid/Invalid ──────────────────────────────────────────────

func (it Severity) IsValid() bool   { return it != Unknown }
func (it Severity) IsInvalid() bool { return it == Unknown }

// ── Range info (BasicEnumer) — all delegate to BasicEnumImpl ───

func (it Severity) RangeNamesCsv() string            { return BasicEnumImpl.RangeNamesCsv() }
func (it Severity) MinMaxAny() (min, max any)        { return BasicEnumImpl.MinMaxAny() }
func (it Severity) MinValueString() string           { return BasicEnumImpl.MinValueString() }
func (it Severity) MaxValueString() string           { return BasicEnumImpl.MaxValueString() }
func (it Severity) MaxInt() int                      { return BasicEnumImpl.MaxInt() }
func (it Severity) MinInt() int                      { return BasicEnumImpl.MinInt() }
func (it Severity) RangesDynamicMap() map[string]any { return BasicEnumImpl.RangesDynamicMap() }
func (it Severity) AllNameValues() []string          { return BasicEnumImpl.AllNameValues() }
func (it Severity) IntegerEnumRanges() []int         { return BasicEnumImpl.IntegerEnumRanges() }

func (it Severity) OnlySupportedErr(names ...string) error {
    return BasicEnumImpl.OnlySupportedErr(names...)
}
func (it Severity) OnlySupportedMsgErr(message string, names ...string) error {
    return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

// ── Format ─────────────────────────────────────────────────────

func (it Severity) Format(format string) string {
    return BasicEnumImpl.Format(format, it.Value())
}

// ── Type-specific: BasicInt8Enumer ─────────────────────────────

func (it Severity) MaxInt8() int8      { return BasicEnumImpl.Max() }
func (it Severity) MinInt8() int8      { return BasicEnumImpl.Min() }
func (it Severity) RangesInt8() []int8 { return BasicEnumImpl.Ranges() }
func (it Severity) ToEnumString(input int8) string { return BasicEnumImpl.ToEnumString(input) }

// ── Range validation ───────────────────────────────────────────

func (it Severity) IsValidRange() bool          { return BasicEnumImpl.IsValidRange(it.Value()) }
func (it Severity) IsInvalidRange() bool        { return !it.IsValidRange() }
func (it Severity) RangesInvalidMessage() string { return BasicEnumImpl.RangesInvalidMessage() }
func (it Severity) RangesInvalidErr() error     { return BasicEnumImpl.RangesInvalidErr() }

func (it Severity) StringRanges() []string    { return BasicEnumImpl.StringRanges() }
func (it Severity) StringRangesPtr() []string { return BasicEnumImpl.StringRangesPtr() }

// ── JSON ───────────────────────────────────────────────────────

func (it Severity) MarshalJSON() ([]byte, error) {
    return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}

func (it *Severity) UnmarshalJSON(data []byte) error {
    val, err := it.UnmarshallEnumToValueInt8(data)
    if err == nil { *it = Severity(val) }
    return err
}

func (it Severity) UnmarshallEnumToValueInt8(jsonUnmarshallingValue []byte) (int8, error) {
    return BasicEnumImpl.UnmarshallToValue(true, jsonUnmarshallingValue)
}

// ── EnumType ───────────────────────────────────────────────────

func (it Severity) EnumType() enuminf.EnumTyper {
    return BasicEnumImpl.EnumType()
}
```

---

## Int16 Enum — Full Pattern

Use `int16` when values exceed `int8` range (-128..127) or you need a larger ordinal space.

### `vars.go`

```go
package region

import (
    "github.com/alimtvnetwork/core/coreimpl/enumimpl"
    "github.com/alimtvnetwork/core/internal/reflectinternal"
)

type Region int16

const (
    Unknown Region = iota
    USEast
    USWest
    Europe
    AsiaPacific
)

var (
    Ranges = [...]string{
        Unknown:     "Unknown",
        USEast:      "USEast",
        USWest:      "USWest",
        Europe:      "Europe",
        AsiaPacific: "AsiaPacific",
    }

    BasicEnumImpl = enumimpl.New.BasicInt16.UsingTypeSlice(
        reflectinternal.TypeName(Unknown),
        Ranges[:],
    )
)
```

### Key Differences from Int8

The method set is identical to Int8 except:

```go
// Value accessor returns int16
func (it Region) Value() int16 { return int16(it) }

// Type-specific interface: BasicInt16Enumer
func (it Region) MaxInt16() int16      { return BasicEnumImpl.Max() }
func (it Region) MinInt16() int16      { return BasicEnumImpl.Min() }
func (it Region) RangesInt16() []int16 { return BasicEnumImpl.Ranges() }
func (it Region) ToEnumString(input int16) string { return BasicEnumImpl.ToEnumString(input) }

// Typed equality
func (it Region) IsValueEqual(value int16) bool { return int16(it) == value }
func (it Region) IsAnyValuesEqual(anyValues ...int16) bool {
    v := int16(it)
    for _, val := range anyValues { if v == val { return true } }
    return false
}

// JSON unmarshal
func (it Region) UnmarshallEnumToValueInt16(data []byte) (int16, error) {
    return BasicEnumImpl.UnmarshallToValue(true, data)
}
```

All other methods (Name, String, TypeName, MarshalJSON, Format, Range*, etc.) are identical — just delegate to `BasicEnumImpl` the same way.

---

## Int32 Enum — Full Pattern

Use `int32` for enums with large value ranges or when interoperating with systems that use 32-bit identifiers.

### `vars.go`

```go
package errorcode

import (
    "github.com/alimtvnetwork/core/coreimpl/enumimpl"
    "github.com/alimtvnetwork/core/internal/reflectinternal"
)

type ErrorCode int32

const (
    None          ErrorCode = iota
    NotFound
    Unauthorized
    ServerError
    RateLimited
)

var (
    Ranges = [...]string{
        None:         "None",
        NotFound:     "NotFound",
        Unauthorized: "Unauthorized",
        ServerError:  "ServerError",
        RateLimited:  "RateLimited",
    }

    BasicEnumImpl = enumimpl.New.BasicInt32.UsingTypeSlice(
        reflectinternal.TypeName(None),
        Ranges[:],
    )
)
```

### Key Differences from Int8/Int16

```go
// Value accessor returns int32
func (it ErrorCode) Value() int32 { return int32(it) }

// Type-specific interface: BasicInt32Enumer
func (it ErrorCode) MaxInt32() int32      { return BasicEnumImpl.Max() }
func (it ErrorCode) MinInt32() int32      { return BasicEnumImpl.Min() }
func (it ErrorCode) RangesInt32() []int32 { return BasicEnumImpl.Ranges() }
func (it ErrorCode) ToEnumString(input int32) string { return BasicEnumImpl.ToEnumString(input) }

// Typed equality
func (it ErrorCode) IsValueEqual(value int32) bool { return int32(it) == value }
func (it ErrorCode) IsAnyValuesEqual(anyValues ...int32) bool {
    v := int32(it)
    for _, val := range anyValues { if v == val { return true } }
    return false
}

// JSON unmarshal
func (it ErrorCode) UnmarshallEnumToValueInt32(data []byte) (int32, error) {
    return BasicEnumImpl.UnmarshallToValue(true, data)
}
```

---

## Alias-Aware Enum Pattern

Aliases let JSON/user input accept multiple names for the same value.

### Byte with Aliases

```go
var BasicEnumImpl = enumimpl.New.BasicByte.CreateUsingSlicePlusAliasMapOptions(
    true,        // include uppercase/lowercase
    Invalid,
    Ranges[:],
    map[string]byte{
        "ok":    byte(Ready),
        "error": byte(Failed),
    },
)
```

### Int8 with Aliases

```go
var BasicEnumImpl = enumimpl.New.BasicInt8.DefaultWithAliasMap(
    Unknown,
    Ranges[:],
    map[string]int8{
        "warn": int8(Medium),
        "crit": int8(Critical),
    },
)
```

### Int32 with Aliases

```go
var BasicEnumImpl = enumimpl.New.BasicInt32.DefaultWithAliasMap(
    None,
    Ranges[:],
    map[string]int32{
        "404": int32(NotFound),
        "401": int32(Unauthorized),
    },
)
```

---

## Case-Insensitive Parsing

For enums that must parse `"ready"`, `"READY"`, and `"Ready"` identically, use the `AllCases` factory variants (available on int8):

```go
// Int8 — case insensitive
var BasicEnumImpl = enumimpl.New.BasicInt8.DefaultAllCases(
    Unknown,
    Ranges[:],
)

// Int8 — case insensitive with aliases
var BasicEnumImpl = enumimpl.New.BasicInt8.DefaultWithAliasMapAllCases(
    Unknown,
    Ranges[:],
    map[string]int8{"warn": int8(Medium)},
)
```

For byte enums, use `CreateUsingSlicePlusAliasMapOptions` with `isIncludeUppercaseLowercase = true`.

---

## Explicit Non-Contiguous Values Pattern

When values must be assigned explicitly (not iota):

### Byte

```go
const (
    Low    Priority = 1
    Medium Priority = 2
    High   Priority = 3
)

var BasicEnumImpl = enumimpl.New.BasicByte.CreateUsingMapPlusAliasMapOptions(
    false,
    Low,
    map[byte]string{
        byte(Low):    "Low",
        byte(Medium): "Medium",
        byte(High):   "High",
    },
    nil,
)
```

### Int8

```go
var BasicEnumImpl = enumimpl.New.BasicInt8.CreateUsingMap(
    reflectinternal.TypeName(Unknown),
    map[int8]string{
        0:  "Unknown",
        10: "Low",
        20: "Medium",
        30: "High",
    },
)
```

### Int32

```go
var BasicEnumImpl = enumimpl.New.BasicInt32.CreateUsingMap(
    reflectinternal.TypeName(None),
    map[int32]string{
        0:   "None",
        404: "NotFound",
        401: "Unauthorized",
        500: "ServerError",
    },
)
```

---

## Formula Rule — Safe vs Unsafe

### Safe for `BasicByte` / `BasicInt8` / `BasicInt16` / `BasicInt32`

- `0, 1, 2, 3, ...` (iota)
- `1, 2, 3, ...`
- Any values representing **one selected member**
- Non-contiguous values via `CreateUsingMap`

### NOT safe as a normal enum

- `1 << 0`, `1 << 1`, `1 << 2` — combinable bitmasks
- Permission formulas like `4`, `2`, `1`, `7`
- Bitwise flag sets

For flags, build a **flags helper** instead. See `chmodhelper/newAttributeCreator.go` for a real example.

---

## Creator Factory Methods Reference

All creators share the same method names. Substitute the type:

| Method | Description |
|--------|-------------|
| `UsingTypeSlice(typeName, names[])` | Contiguous iota enum from string slice |
| `Default(firstItem, names[])` | Same but infers typeName via reflection |
| `DefaultWithAliasMap(firstItem, names[], aliasMap)` | Contiguous + aliases |
| `DefaultAllCases(firstItem, names[])` | Contiguous + upper/lower parsing *(int8 only)* |
| `DefaultWithAliasMapAllCases(firstItem, names[], aliasMap)` | All cases + aliases *(int8 only)* |
| `CreateUsingMap(typeName, map[T]string)` | Non-contiguous explicit values |
| `CreateUsingMapPlusAliasMap(typeName, map[T]string, aliasMap)` | Explicit + aliases |
| `CreateUsingAliasMap(typeName, values[], names[], aliasMap, min, max)` | Full manual control |
| `UsingFirstItemSliceAliasMap(firstItem, names[], aliasMap)` | Infer type + aliases |

---

## Methods Delegation Quick Reference

This table shows which `BasicEnumImpl` method each enum method delegates to:

| Enum Method | Delegates To |
|-------------|-------------|
| `Name()` | `BasicEnumImpl.ToEnumString(it.Value())` |
| `String()` | `BasicEnumImpl.ToEnumString(it.Value())` |
| `TypeName()` | `BasicEnumImpl.TypeName()` |
| `NameValue()` | `BasicEnumImpl.NameWithValue(it.Value())` |
| `ToNumberString()` | `BasicEnumImpl.ToNumberString(it.Value())` |
| `ValueString()` | `BasicEnumImpl.ToNumberString(it.Value())` |
| `RangeNamesCsv()` | `BasicEnumImpl.RangeNamesCsv()` |
| `MinMaxAny()` | `BasicEnumImpl.MinMaxAny()` |
| `MinValueString()` | `BasicEnumImpl.MinValueString()` |
| `MaxValueString()` | `BasicEnumImpl.MaxValueString()` |
| `MaxInt()` | `BasicEnumImpl.MaxInt()` |
| `MinInt()` | `BasicEnumImpl.MinInt()` |
| `RangesDynamicMap()` | `BasicEnumImpl.RangesDynamicMap()` |
| `AllNameValues()` | `BasicEnumImpl.AllNameValues()` |
| `IntegerEnumRanges()` | `BasicEnumImpl.IntegerEnumRanges()` |
| `Format(fmt)` | `BasicEnumImpl.Format(fmt, it.Value())` |
| `OnlySupportedErr(...)` | `BasicEnumImpl.OnlySupportedErr(...)` |
| `MarshalJSON()` | `BasicEnumImpl.ToEnumJsonBytes(it.Value())` |
| `UnmarshalJSON()` | via `UnmarshallEnumToValue*` → `BasicEnumImpl.UnmarshallToValue(true, data)` |
| `EnumType()` | `BasicEnumImpl.EnumType()` |
| `Max*()` / `Min*()` | `BasicEnumImpl.Max()` / `BasicEnumImpl.Min()` |
| `Ranges*()` | `BasicEnumImpl.Ranges()` |
| `IsValidRange()` | `BasicEnumImpl.IsValidRange(it.Value())` |
| `RangesInvalidMessage()` | `BasicEnumImpl.RangesInvalidMessage()` |
| `RangesInvalidErr()` | `BasicEnumImpl.RangesInvalidErr()` |
| `StringRanges()` | `BasicEnumImpl.StringRanges()` |

---

## AI Authoring Checklist

When an AI creates a new enum package:

1. **Choose backing type** — byte (≤255), int8 (≤127), int16, int32
2. **Prefer contiguous constants** with `iota`; use `CreateUsingMap` only for non-contiguous
3. **Put lookup data** in `vars.go` — Ranges array, optional RangesMap, BasicEnumImpl
4. **Build BasicEnumImpl** with `enumimpl.New.Basic<Type>.*`
5. **Implement ALL `BasicEnumValuer` methods** — `ValueByte`, `ValueInt`, `ValueInt8`, `ValueInt16`, `ValueUInt16`, `ValueInt32`, `ValueString`
6. **Implement ALL `BaseEnumer` methods** — `Name`, `String`, `TypeName`, `NameValue`, `ToNumberString`, `IsNameEqual`, `IsAnyNamesOf`
7. **Implement ALL `BasicEnumer` methods** — `Format`, `MinMaxAny`, `Min/MaxValueString`, `Min/MaxInt`, `RangesDynamicMap`, `AllNameValues`, `IntegerEnumRanges`, `OnlySupportedErr`, `OnlySupportedMsgErr`, `EnumType`
8. **Implement type-specific interface** — `BasicByteEnumer`, `BasicInt8Enumer`, etc.
9. **Implement JSON** — `MarshalJSON`, `UnmarshalJSON`, `UnmarshallEnumToValue*`
10. **Implement range validation** — `IsValidRange`, `IsInvalidRange`, `RangesInvalidMessage`, `RangesInvalidErr`
11. **Implement string ranges** — `StringRanges`, `StringRangesPtr`
12. **Add domain `IsX()` methods** only for business logic
13. **Do NOT model bitmask flags** as a plain enum
14. **Split large files** by responsibility

---

## `UnmarshallEnumToValue` — Deep Dive

The JSON unmarshalling chain works in three layers:

```
UnmarshalJSON(data)           ← Go's json package calls this (pointer receiver)
  └→ UnmarshallEnumToValue*(data)   ← enum method, type-specific name
       └→ BasicEnumImpl.UnmarshallToValue(true, data)  ← shared impl in enumimpl
```

### How `BasicEnumImpl.UnmarshallToValue` Works Internally

```
1. If data is nil and isMappedToFirstIfEmpty=false → return error
2. If data is nil and isMappedToFirstIfEmpty=true  → return minVal (first enum value)
3. If data is "" or `""` and isMappedToFirstIfEmpty=true → return minVal
4. Otherwise → call GetValueByName(string(data))
   a. Try exact match in jsonDoubleQuoteNameToValueHashMap
   b. Try wrapping in double quotes and retry
   c. If not found → return error with type name + valid ranges CSV
```

The hashmap is pre-populated with: `"Name"`, `"\"Name\""`, `"0"`, `"\"0\""` for each member, plus any alias entries.

### Implementation Per Backing Type

Each type has a differently-named method to avoid Go interface conflicts:

#### Byte

```go
// The type-specific unmarshaller — name matches enuminf.UnmarshallEnumToValueByter
func (it Status) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
    return BasicEnumImpl.UnmarshallToValue(true, jsonUnmarshallingValue)
}

// Go's JSON hook — MUST be pointer receiver
func (it *Status) UnmarshalJSON(data []byte) error {
    val, err := it.UnmarshallEnumToValue(data)
    if err == nil {
        *it = Status(val)
    }
    return err
}
```

#### Int8

```go
func (it Severity) UnmarshallEnumToValueInt8(jsonUnmarshallingValue []byte) (int8, error) {
    return BasicEnumImpl.UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it *Severity) UnmarshalJSON(data []byte) error {
    val, err := it.UnmarshallEnumToValueInt8(data)
    if err == nil {
        *it = Severity(val)
    }
    return err
}
```

#### Int16

```go
func (it Region) UnmarshallEnumToValueInt16(jsonUnmarshallingValue []byte) (int16, error) {
    return BasicEnumImpl.UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it *Region) UnmarshalJSON(data []byte) error {
    val, err := it.UnmarshallEnumToValueInt16(data)
    if err == nil {
        *it = Region(val)
    }
    return err
}
```

#### Int32

```go
func (it ErrorCode) UnmarshallEnumToValueInt32(jsonUnmarshallingValue []byte) (int32, error) {
    return BasicEnumImpl.UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it *ErrorCode) UnmarshalJSON(data []byte) error {
    val, err := it.UnmarshallEnumToValueInt32(data)
    if err == nil {
        *it = ErrorCode(val)
    }
    return err
}
```

### Method Naming Convention

| Backing Type | Unmarshall Method Name | Interface |
|---|---|---|
| `byte` | `UnmarshallEnumToValue([]byte) (byte, error)` | `enuminf.UnmarshallEnumToValueByter` |
| `int8` | `UnmarshallEnumToValueInt8([]byte) (int8, error)` | `enuminf.BasicInt8Enumer` |
| `int16` | `UnmarshallEnumToValueInt16([]byte) (int16, error)` | `enuminf.BasicInt16Enumer` |
| `int32` | `UnmarshallEnumToValueInt32([]byte) (int32, error)` | `enuminf.BasicInt32Enumer` |

### `isMappedToFirstIfEmpty` Parameter

- **`true`** (default for most enums) — nil/empty JSON → first enum value (usually `Invalid`/`Unknown`). Use this when you want graceful degradation.
- **`false`** — nil/empty JSON → returns error. Use this when you want strict parsing and missing values should be rejected.

---

## `Format()` — Deep Dive

The `Format` method outputs a human-readable string by replacing placeholders in a format template.

### Placeholder Keys

| Placeholder | Replaced With | Example Value |
|---|---|---|
| `{type-name}` | `TypeName()` — the full Go type name | `status.Status` |
| `{name}` | `Name()` — the enum member name | `Ready` |
| `{value}` | `ValueString()` — the numeric value as string | `2` |

### Implementation

The enum method delegates to `numberEnumBase.Format()`:

```go
// On the enum type
func (it Status) Format(format string) string {
    return BasicEnumImpl.Format(format, it.Value())
}

// Inside numberEnumBase (what actually runs)
func (it numberEnumBase) Format(format string, value any) string {
    return Format(
        it.TypeName(),      // replaces {type-name}
        it.ToName(value),   // replaces {name}
        it.ValueString(value), // replaces {value}
        format,
    )
}
```

### Usage Examples

```go
s := status.Ready

// Basic info string
s.Format("Enum: {name} ({value})")
// → "Enum: Ready (2)"

// Full diagnostic
s.Format("Enum of {type-name} - {name} - {value}")
// → "Enum of status.Status - Ready - 2"

// Log line
s.Format("[{type-name}] current={name}")
// → "[status.Status] current=Ready"

// Error message
s.Format("invalid state: expected Ready, got {name} (raw={value})")
// → "invalid state: expected Ready, got Failed (raw=3)"

// Just the name
s.Format("{name}")
// → "Ready"

// Works the same on any backing type
severity := severity.Critical
severity.Format("{type-name}::{name}={value}")
// → "severity.Severity::Critical=4"
```

### When to Use Format vs Name/String

| Need | Use |
|---|---|
| Just the member name | `Name()` or `String()` |
| Name with value in parens | `NameValue()` → `"Ready (2)"` |
| Just the numeric value | `ToNumberString()` or `ValueString()` |
| Custom template with type info | `Format("{type-name} - {name} - {value}")` |
| Logging / diagnostics | `Format(...)` with your log template |

---

## Related Docs

- [coreimpl/enumimpl/readme.md](/coreimpl/enumimpl/readme.md)
- [coreinterface/enuminf/README.md](/coreinterface/enuminf/README.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [newCreator Convention](/spec/01-app/18-new-creator-convention.md)
