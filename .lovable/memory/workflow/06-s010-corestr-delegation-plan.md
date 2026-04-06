# S-010: corestr → coregeneric Delegation Plan

**Date**: 2026-04-06  
**Status**: Planning  
**Scope**: Refactor all 6 corestr types to embed coregeneric counterparts internally

---

## Overview

Move shared collection behavior from `corestr` into `coregeneric` by having each `corestr` type embed its generic counterpart. All exported `corestr` method signatures remain **unchanged** — only internal implementation delegates to `coregeneric`.

**Total lines affected**: ~9,014 across 6 files

---

## Critical Constraints

1. **NEVER remove any exported method** — they are public framework APIs
2. **NEVER change any method signature** — external consumers depend on them
3. **All existing tests must continue to pass**
4. **No circular imports** — `corestr` → `coregeneric` is safe (already verified)

---

## Type-by-Type Analysis

### 1. Collection (2,206 lines) — EMBEDDABLE ✅

**Current struct:**
```go
type Collection struct {
    items []string
    sync.Mutex
}
```

**Proposed struct:**
```go
type Collection struct {
    inner *coregeneric.Collection[string]
}
```

**Notes:**
- `coregeneric.Collection[string]` already embeds `sync.Mutex`, so `Lock()`/`Unlock()` are still promoted
- All `it.items` references (100+) must change to `it.inner.Items()` or use a helper `it.items()` method
- **Delegatable methods** (~30): Add, AddLock, Adds, AddsLock, AddSlice, AddIf, AddIfMany, AddFunc, AddCollection, AddCollections, RemoveAt, First, Last, FirstOrDefault, LastOrDefault, Length, LengthLock, Count, Capacity, IsEmpty, IsEmptyLock, HasItems, HasAnyItem, HasIndex, LastIndex, ForEach, ForEachBreak, Filter, Clone, Reverse, ConcatNew, CountFunc
- **String-specific methods** (~70): Join, NonEmptyList, NonEmptyJoin, JsonPtr, JsonString, ToSimpleSlice, Sort, IsEquals, IsEqualsWithSensitive, IsContains, IndexOf, SafeIndexAt, GetPagesSize, GetPagedCollection, InsertAt, EachItemSplitBy, etc.
- **corestr-specific methods** (~20): AddHashmapsKeys, AddHashmapsValues, AddCollection(corestr), ConcatNew(corestr), AddError, AsError, ToError, etc.

**Risk**: MEDIUM — large file but straightforward field swap. The `sync.Mutex` is embedded in `coregeneric.Collection` so `Lock()/Unlock()` still work via promotion.

**Challenge**: Methods returning `*Collection` (corestr) can't be auto-promoted from coregeneric. Every chaining method needs a wrapper.

---

### 2. SimpleSlice (1,321 lines) — NOT EMBEDDABLE ❌

**Current type:**
```go
type SimpleSlice []string
```

**Problem**: This is a **type alias for `[]string`**, not a struct. It cannot embed a field. Changing it to a struct would break:
- All `*it = append(*it, ...)` patterns
- All `[]string(*it)` conversions
- All `*SimpleSlice` pointer semantics
- The `MarshalJSON`/`UnmarshalJSON` implementations
- External code that does `SimpleSlice(someSlice)`

**Proposed approach**: Use **shared helper functions** instead of embedding:
```go
// In coregeneric, add package-level helper functions:
func SliceAdd[T any](s *[]T, item T)
func SliceFilter[T any](s []T, pred func(T) bool) []T
// etc.

// In corestr SimpleSlice, delegate:
func (it *SimpleSlice) Filter(...) *SimpleSlice {
    result := coregeneric.SliceFilter(...)
    ...
}
```

**Risk**: LOW — minimal changes, just adding delegation for common logic.

---

### 3. Hashmap (1,315 lines) — PARTIALLY EMBEDDABLE ⚠️

**Current struct:**
```go
type Hashmap struct {
    hasMapUpdated bool
    items         map[string]string
    cachedList    []string
    sync.Mutex
}
```

**Problem**: `coregeneric.Hashmap[string, string]` does NOT have:
- `hasMapUpdated bool` — cache invalidation flag
- `cachedList []string` — cached keys list

**Proposed struct:**
```go
type Hashmap struct {
    inner         *coregeneric.Hashmap[string, string]
    hasMapUpdated bool
    cachedList    []string
}
```

**Notes:**
- `sync.Mutex` is embedded in `coregeneric.Hashmap`, so Lock/Unlock still promoted
- All `it.items[key]` references must change to `it.inner.Map()[key]` or use a helper
- Cache logic (`hasMapUpdated`, `cachedList`) stays in corestr
- **Delegatable methods** (~15): IsEmpty, HasItems, Length, LengthLock, Has, Contains, ContainsLock, IsKeyMissing, Remove, RemoveLock, Keys, Values, Map, ForEach, ForEachBreak, Clone, ConcatNew, IsEquals
- **String-specific methods** (~30): Set with string conversion variants, AddOrUpdateKeyStrValInt, AddOrUpdateKeyStrValFloat, ValuesList, KeysList, SortedKeys, Join, etc.

**Risk**: MEDIUM — cache invalidation logic must be preserved.

---

### 4. Hashset (1,475 lines) — PARTIALLY EMBEDDABLE ⚠️

**Current struct:**
```go
type Hashset struct {
    hasMapUpdated bool
    items         map[string]bool
    cachedList    []string
    sync.Mutex
}
```

**Same problem as Hashmap** — extra cache fields.

**Proposed struct:**
```go
type Hashset struct {
    inner         *coregeneric.Hashset[string]
    hasMapUpdated bool
    cachedList    []string
}
```

**Notes:**
- **Delegatable methods** (~20): IsEmpty, HasItems, Length, LengthLock, Add, AddBool, AddLock, Adds, AddSlice, AddSliceLock, AddIf, AddIfMany, AddHashsetItems, AddItemsMap, Has, Contains, ContainsLock, HasAll, HasAny, Remove, RemoveLock, List, ListPtr, Map, Resize, IsEquals
- **String-specific methods** (~25): AddNonEmpty, AddNonEmptyWhitespace, AddPtr, AddPtrLock, IsContains, IsContainsFold, Collection (returns corestr.Collection), Sort, Join, etc.
- **Cache-aware methods**: List/ListPtr use cachedList when hasMapUpdated is false

**Risk**: MEDIUM — cache invalidation must be maintained.

---

### 5. LinkedList (1,146 lines) — EMBEDDABLE ✅

**Current struct:**
```go
type LinkedList struct {
    head, tail *LinkedListNode
    length     int
    sync.Mutex
}
```

**Problem**: `coregeneric.LinkedList[string]` uses `*coregeneric.LinkedListNode[string]`, but `corestr.LinkedList` uses `*corestr.LinkedListNode`. These are different types — cannot embed directly.

**Proposed approach**: Cannot embed due to node type mismatch. Use **shared helper functions** for common traversal/manipulation logic, keeping the current struct.

**Risk**: LOW if using helper approach.

---

### 6. LinkedCollections (1,551 lines) — NOT EMBEDDABLE ❌

**Current struct:**
```go
type LinkedCollections struct {
    head, tail *LinkedCollectionNode
    length     int
    sync.Mutex
}
```

**Problem**: Node type is `*LinkedCollectionNode` which holds `*Collection` (corestr). No generic equivalent exists. Would need `coregeneric.LinkedList[*corestr.Collection]` which creates a circular dependency.

**Proposed approach**: Keep as-is. No delegation possible without circular imports.

**Risk**: N/A

---

## Execution Plan

### Phase 1: Foundation (coregeneric helpers)
Add shared package-level helper functions to coregeneric that both embedded and non-embedded types can use:
- `SliceNonEmpty[T comparable](items []T, zero T) []T`
- `SliceJoin[T any](items []T, stringer func(T) string, sep string) string`
- Other common patterns

### Phase 2: Collection
1. Change struct to embed `*coregeneric.Collection[string]`
2. Add `items()` accessor method returning `it.inner.Items()`
3. Refactor all methods to delegate where possible
4. Run tests

### Phase 3: Hashmap + Hashset
1. Change structs to embed coregeneric counterparts
2. Preserve cache logic in corestr layer
3. Refactor delegatable methods
4. Run tests

### Phase 4: LinkedList + SimpleSlice
1. Add shared helper functions for LinkedList traversal
2. SimpleSlice: delegate via package-level helpers where possible
3. Run tests

### Phase 5: LinkedCollections
1. Keep as-is (no delegation possible)
2. Document the reasoning

---

## Summary Table

| Type | Lines | Approach | Delegatable Methods | Risk |
|------|-------|----------|-------------------|------|
| Collection | 2,206 | Embed `*coregeneric.Collection[string]` | ~30 | MEDIUM |
| SimpleSlice | 1,321 | Helper functions (not a struct) | ~10 | LOW |
| Hashmap | 1,315 | Embed with cache preservation | ~15 | MEDIUM |
| Hashset | 1,475 | Embed with cache preservation | ~20 | MEDIUM |
| LinkedList | 1,146 | Helper functions (node type mismatch) | ~5 | LOW |
| LinkedCollections | 1,551 | Keep as-is (circular dep risk) | 0 | N/A |

**Total delegatable method count**: ~80 methods that will use coregeneric internally.
