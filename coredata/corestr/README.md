# corestr — String-Centric Collections

## Table of Contents

1. [Overview](#overview)
2. [Entry Points](#entry-points)
3. [Core Types](#core-types)
4. [New Creator Pattern](#new-creator-pattern)
5. [Key Capabilities](#key-capabilities)
   - [String-Specific Methods](#string-specific-methods-not-in-coregeneric)
   - [Data Models](#data-models)
6. [File Organization](#file-organization)
7. [How to Extend Safely](#how-to-extend-safely)
8. [Contributors](#contributors)
9. [Issues for Future Reference](#issues-for-future-reference)

---

## Overview

The `corestr` package provides string-specialized collection types with rich string manipulation methods (Join, EqualFold, Trim, Split, etc.). It is the primary data structure package for string-based workflows and sits atop `coregeneric` in the type hierarchy.

## Entry Points

| Variable | Type | Description |
|----------|------|-------------|
| `New` | `*newCreator` | Builder-pattern factory for all string types |
| `Empty` | `*emptyCreator` | Quick empty-instance factory |
| `StringUtils` | `utils` | Standalone string utility functions |

## Core Types

| Type | Description |
|------|-------------|
| `Collection` | Thread-safe string slice with Join, EqualFold, Trim, Filter, Sort, and `*Lock` variants |
| `Hashset` | String set backed by `map[string]bool` |
| `Hashmap` | String-keyed map wrapper with diff support (`HashmapDiff`) |
| `SimpleSlice` | Lightweight `[]string` wrapper with Add/Filter/Skip/Take |
| `LinkedList` | Singly-linked string list with head/tail pointers |
| `LinkedCollections` | Linked list of `Collection` nodes |
| `LinkedCollectionNode` | Node within a `LinkedCollections` list |
| `HashsetsCollection` | Collection of `Hashset` instances |
| `CollectionsOfCollection` | Nested collection of collections |
| `CharCollectionMap` | Map of `rune` → `Collection` |
| `CharHashsetMap` | Map of `rune` → `Hashset` |
| `SimpleStringOnce` | Compute-once string holder with lazy initialization |
| `KeyValuePair` / `KeyAnyValuePair` | Typed key-value pairs |
| `KeyValueCollection` | Collection of key-value pairs |
| `LeftRight` / `LeftMiddleRight` | Positional string containers |
| `ValidValue` / `ValidValues` | Validated string wrappers |
| `ValueStatus` | String with status flag |
| `TextWithLineNumber` | String line with line number metadata |

## New Creator Pattern

```go
// Collection
col := corestr.New.Collection.Cap(10)
col := corestr.New.Collection.Items("a", "b", "c")

// Hashset
set := corestr.New.Hashset.Items("x", "y")

// Hashmap
hm := corestr.New.Hashmap.Cap(20)

// SimpleSlice
ss := corestr.New.SimpleSlice.Cap(5)

// LinkedList
ll := corestr.New.LinkedList.Empty()

// Empty pattern
empty := corestr.Empty.SimpleSlice()
```

## Key Capabilities

### String-Specific Methods (not in coregeneric)

- **Join** / **JoinLock** — Join elements with separator
- **EqualFold** — Case-insensitive comparison
- **Trim** / **TrimSpace** — Whitespace handling
- **Split** / **SplitBy** — String splitting
- **Contains** / **HasPrefix** / **HasSuffix** — String matching
- **ToUpper** / **ToLower** — Case conversion

### Data Models

- `HashmapDataModel` — Serializable hashmap representation
- `HashsetDataModel` — Serializable hashset representation
- `CharCollectionDataModel` — Serializable char-collection map
- `CharHashsetDataModel` — Serializable char-hashset map
- `CollectionsOfCollectionModel` — Nested collection model
- `HashsetsCollectionDataModel` — Collection of hashsets model
- `SimpleStringOnceModel` — Serializable once-computed string

## File Organization

| File Pattern | Responsibility |
|-------------|---------------|
| `Collection.go` | Main string collection type |
| `Hashset.go` / `Hashmap.go` | Set and map types |
| `LinkedList.go` / `LinkedCollections.go` | Linked list types |
| `SimpleSlice.go` | Lightweight slice wrapper |
| `*DataModel.go` | Serializable data models |
| `new*Creator.go` | Factory structs per type |
| `emptyCreator.go` | Empty-instance factories |
| `funcs.go` | Package-level utility functions |
| `utils.go` | String utility methods |
| `vars.go` | Package-level variables |

## How to Extend Safely

- New string collection types go here with a `newXCreator` and an entry in `newCreator.go`.
- String-agnostic operations should go in `coregeneric` instead.
- Always provide both pointer and value receiver variants for thread safety.

## Contributors

## Issues for Future Reference
