# coregeneric — Generic Data Structures

## Table of Contents

1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Entry Points](#entry-points)
4. [Generic Types](#generic-types)
5. [Typed Creator Pattern](#typed-creator-pattern)
6. [Generic Functions](#generic-functions)
   - [funcs.go — any Constraint](#funcsgo--any-constraint)
   - [comparablefuncs.go — comparable Constraint](#comparablefuncsgo--comparable-constraint)
   - [orderedfuncs.go — cmp.Ordered Constraint](#orderedfuncsgo--cmpordered-constraint)
7. [Type Aliases](#type-aliases)
8. [Supported Primitive Types](#supported-primitive-types)
9. [Key Patterns](#key-patterns)
10. [How to Extend Safely](#how-to-extend-safely)
11. [Related Docs](#related-docs)

## Overview

The `coregeneric` package provides **type-parameterized** versions of all core collection data structures. It serves as the foundational layer that type-specific packages like `corestr` build upon, offering code reuse, consistency, compile-time type safety, and IDE discoverability via the New Creator pattern.

## Architecture

```
coredata/coregeneric/
├── vars.go                      # New = &newCreator{}
├── newCreator.go                # Root aggregator struct
├── newCollectionCreator.go      # Collection sub-creator (16 typed fields)
├── newHashsetCreator.go         # Hashset sub-creator
├── newHashmapCreator.go         # Hashmap sub-creator
├── newSimpleSliceCreator.go     # SimpleSlice sub-creator
├── newLinkedListCreator.go      # LinkedList sub-creator
├── typedCollectionCreator.go    # Generic typed creator for Collection[T]
├── typedHashsetCreator.go       # Generic typed creator for Hashset[T]
├── typedHashmapCreator.go       # Generic typed creator for Hashmap[K,V]
├── typedSimpleSliceCreator.go   # Generic typed creator for SimpleSlice[T]
├── typedLinkedListCreator.go    # Generic typed creator for LinkedList[T]
├── Collection.go                # Collection[T] type + methods
├── Hashset.go                   # Hashset[T] type + methods
├── Hashmap.go                   # Hashmap[K,V] type + methods
├── SimpleSlice.go               # SimpleSlice[T] type + methods
├── LinkedList.go                # LinkedList[T] singly-linked list
├── LinkedListNode.go            # LinkedListNode[T] node type
├── types.go                     # Type aliases for all primitives
├── funcs.go                     # Generic functions (any constraint)
├── comparablefuncs.go           # Functions requiring comparable
├── orderedfuncs.go              # Functions requiring cmp.Ordered
└── README.md
```

## Entry Points

| Variable | Type | Description |
|----------|------|-------------|
| `New` | `*newCreator` | Root aggregator for the New Creator pattern |

## Generic Types

| Type | Constraint | Description |
|------|-----------|-------------|
| `Collection[T]` | `any` | Thread-safe slice-backed list with Add/Remove/Filter/Sort/Clone and `*Lock` variants |
| `Hashset[T]` | `comparable` | Thread-safe set backed by `map[T]bool` with Add/Has/Remove/Resize |
| `Hashmap[K, V]` | `K: comparable, V: any` | Thread-safe map wrapper with Set/Get/Has/Remove/Keys/Values |
| `SimpleSlice[T]` | `any` | Lightweight typed slice wrapper with Add/Filter/Skip/Take |
| `LinkedList[T]` | `any` | Singly-linked list with head/tail pointers and embedded mutex |

## Typed Creator Pattern

Instead of writing 16 separate creator structs, a single generic struct handles all primitive types:

```go
// One generic definition covers ALL types
type typedCollectionCreator[T any] struct{}
func (it typedCollectionCreator[T]) Empty() *Collection[T] { ... }
func (it typedCollectionCreator[T]) Cap(capacity int) *Collection[T] { ... }
func (it typedCollectionCreator[T]) From(items []T) *Collection[T] { ... }
func (it typedCollectionCreator[T]) Items(items ...T) *Collection[T] { ... }

// Instantiated per-type via struct fields
type newCollectionCreator struct {
    String  typedCollectionCreator[string]
    Int     typedCollectionCreator[int]
    Float64 typedCollectionCreator[float64]
    // ... all 16 primitive types
}
```

### Usage Examples

```go
col := coregeneric.New.Collection.String.Cap(10)
set := coregeneric.New.Hashset.Int.Items(1, 2, 3)
hm  := coregeneric.New.Hashmap.StringString.Cap(20)
ss  := coregeneric.New.SimpleSlice.Float64.Items(1.0, 2.5)
ll  := coregeneric.New.LinkedList.String.Empty()
```

## Generic Functions

Go does not allow additional type parameters on methods of generic types. Cross-type transformations are package-level functions organized by constraint level.

### `funcs.go` — `any` Constraint

| Function | Signature | Description |
|----------|-----------|-------------|
| `MapCollection` | `[T, U any]` | Transform `Collection[T]` → `Collection[U]` |
| `FlatMapCollection` | `[T, U any]` | Map + flatten `Collection[T]` → `Collection[U]` |
| `ReduceCollection` | `[T, U any]` | Reduce `Collection[T]` → `U` |
| `GroupByCollection` | `[T any, K comparable]` | Group into `map[K]*Collection[T]` |
| `ContainsFunc` | `[T any]` | Check if any item matches predicate |
| `IndexOfFunc` | `[T any]` | Index of first item matching predicate, or `-1` |
| `ContainsItem` | `[T comparable]` | Check if item exists by equality |
| `IndexOfItem` | `[T comparable]` | Index of first occurrence, or `-1` |
| `Distinct` | `[T comparable]` | Deduplicate collection |
| `MapSimpleSlice` | `[T, U any]` | Transform `SimpleSlice[T]` → `SimpleSlice[U]` |

### `comparablefuncs.go` — `comparable` Constraint

| Function | Target | Description |
|----------|--------|-------------|
| `ContainsAll` | `Collection[T]` | True if collection contains all given items |
| `ContainsAny` | `Collection[T]` | True if collection contains any given item |
| `RemoveItem` | `Collection[T]` | Remove first occurrence; returns bool |
| `RemoveAllItems` | `Collection[T]` | Remove all occurrences; returns count |
| `ToHashset` | `Collection[T]` | Convert collection to `Hashset[T]` |
| `DistinctSimpleSlice` | `SimpleSlice[T]` | Deduplicate simple slice |
| `ContainsSimpleSliceItem` | `SimpleSlice[T]` | Check if item exists in simple slice |

### `orderedfuncs.go` — `cmp.Ordered` Constraint

#### Collection Functions

| Function | Description |
|----------|-------------|
| `SortCollection` | Sort ascending (in-place) |
| `SortCollectionDesc` | Sort descending (in-place) |
| `MinCollection` / `MaxCollection` | Min/max element (panics if empty) |
| `MinCollectionOrDefault` / `MaxCollectionOrDefault` | Min/max with fallback default |
| `IsSortedCollection` | Check if sorted ascending |
| `SumCollection` | Sum all elements |
| `ClampCollection` | Clamp all values to `[min, max]` |

#### SimpleSlice Functions

| Function | Description |
|----------|-------------|
| `SortSimpleSlice` | Sort ascending (in-place) |
| `SortSimpleSliceDesc` | Sort descending (in-place) |
| `MinSimpleSlice` / `MaxSimpleSlice` | Min/max element (panics if empty) |
| `SumSimpleSlice` | Sum all elements |

#### Hashset Functions

| Function | Description |
|----------|-------------|
| `SortedListHashset` | Sorted ascending slice of items |
| `SortedListDescHashset` | Sorted descending slice of items |
| `MinHashset` / `MaxHashset` | Min/max element (panics if empty) |
| `MinHashsetOrDefault` / `MaxHashsetOrDefault` | Min/max with fallback default |
| `SortedCollectionHashset` | Convert to sorted `Collection[T]` |

#### Hashmap Functions

| Function | Description |
|----------|-------------|
| `SortedKeysHashmap` | Keys sorted ascending |
| `SortedKeysDescHashmap` | Keys sorted descending |
| `MinKeyHashmap` / `MaxKeyHashmap` | Min/max key (panics if empty) |
| `MinKeyHashmapOrDefault` / `MaxKeyHashmapOrDefault` | Min/max key with default |
| `SortedValuesHashmap` | Values sorted ascending (requires `V: cmp.Ordered`) |
| `MinValueHashmap` / `MaxValueHashmap` | Min/max value (requires `V: cmp.Ordered`) |
| `MinValueHashmapOrDefault` / `MaxValueHashmapOrDefault` | Min/max value with default |

## Type Aliases

Pre-defined aliases for all common primitives are in `types.go`:

| Data Structure | Aliases |
|----------------|---------|
| `Collection[T]` | `StringCollection`, `IntCollection`, `Int8Collection`, `Int16Collection`, `Int32Collection`, `Int64Collection`, `UintCollection`, `Uint8Collection`, `Uint16Collection`, `Uint32Collection`, `Uint64Collection`, `Float32Collection`, `Float64Collection`, `ByteCollection`, `BoolCollection`, `AnyCollection` |
| `Hashset[T]` | `StringHashset`, `IntHashset`, `Float32Hashset`, `BoolHashset`, ... (15 types) |
| `Hashmap[K,V]` | `StringStringHashmap`, `StringIntHashmap`, `StringAnyHashmap`, `IntStringHashmap`, `Int64StringHashmap`, ... (12 combinations) |
| `SimpleSlice[T]` | `StringSimpleSlice`, `IntSimpleSlice`, `AnySimpleSlice`, ... (16 types) |
| `LinkedList[T]` | `StringLinkedList`, `IntLinkedList`, `AnyLinkedList`, ... (16 types) |

## Supported Primitive Types

| Category | Types |
|----------|-------|
| Signed integers | `int`, `int8`, `int16`, `int32`, `int64` |
| Unsigned integers | `uint`, `uint8`, `uint16`, `uint32`, `uint64` |
| Floats | `float32`, `float64` |
| Other | `byte`, `bool`, `string`, `any` |

## Key Patterns

- All types use embedded `sync.Mutex` for thread-safe `*Lock` variants.
- Zero-nil safety: nil receivers return safe defaults (0, empty, false).
- Constraint hierarchy: `cmp.Ordered ⊂ comparable ⊂ any`.

## How to Extend Safely

- **New primitive type**: Add a field to each `newXCreator` struct and a type alias in `types.go`.
- **New data structure**: Create the generic type, a `typedXCreator[T]`, and a `newXCreator` aggregator. Add the field to `newCreator`.
- **New generic function**: Add to `funcs.go` / `comparablefuncs.go` / `orderedfuncs.go` based on constraint.

## Related Docs

- [Repo Overview](../../spec/01-app/00-repo-overview.md)
