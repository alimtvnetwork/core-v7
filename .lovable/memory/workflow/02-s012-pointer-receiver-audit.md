# S-012: Pointer Receiver Audit — Decision Log

## Date: 2026-04-06

## Summary

Audited pointer receivers across `corestr`, `errcore`, and `corepayload`. Migrated 46 methods to value receivers. Several issues in the audit process are documented below for future reference.

---

## Issues Found in the Audit Process

### 1. Initial Over-Estimation of Candidates

**Problem**: The first scan reported **79 candidates** across all three packages, including complex methods like `PayloadWrapper.Deserialize`, `PayloadWrapper.Serialize`, and `PayloadWrapper.Clone`. These are **not** safe migrations — they involve mutation, deserialization side-effects, or return pointers tied to internal state.

**Root Cause**: The Python candidate-finder script used a simple heuristic (no `it.field =` assignment detected in the first 300 chars of the method body). This missed:
- Methods that mutate via nested calls (e.g., `it.initializeAuthOnDemand()` which sets `it.Attributes.AuthInfo`)
- Methods that return `*T` tied to internal state (e.g., `Clone()` returning a new pointer)
- Methods longer than 300 chars where the mutation occurs later in the body

**Lesson**: Automated candidate scanning needs a **conservative exclusion list** — any method that calls another method on `it` which is itself a pointer receiver should be flagged for manual review, not auto-approved.

### 2. corepayload: All Methods Require Pointer Receivers

**Problem**: Initially identified 19 `Attributes` getter candidates and 32 `PayloadWrapper` candidates. After manual review, **zero** were safe to migrate.

**Root Cause**: 
- `Attributes` getters call `it.IsEmpty()` which contains `it == nil` check → requires pointer receiver
- `PagingInfo` methods all have `it == nil` guards → requires pointer receiver  
- `PayloadWrapper` methods either mutate state or call through to methods that do

**Lesson**: When a type's foundational check methods (`IsEmpty`, `IsNull`, `IsValid`) use nil-guards, **all methods that call them transitively must also be pointer receivers**. This creates a "nil-guard contagion" effect.

### 3. corestr: Major Types All Embed sync.Mutex

**Problem**: Initial scan found hundreds of pointer receiver methods in `corestr`. Investigation revealed all major types (`Collection`, `Hashmap`, `Hashset`, `CharCollectionMap`, `CharHashsetMap`, `LinkedList`, `LinkedCollections`) embed `sync.Mutex`.

**Root Cause**: Types with embedded `sync.Mutex` **must** use pointer receivers for all methods — copying a mutex is a race condition. This eliminated the vast majority of `corestr` candidates.

**Lesson**: Always check for `sync.Mutex` embedding **first** before scanning individual methods. It's a package-level disqualifier for most types.

### 4. Value Receiver Method Chains

**Problem**: `HasValidNonEmptyLeft()` calls `it.IsLeftEmpty()`. If we migrate `HasValidNonEmptyLeft` to a value receiver but leave `IsLeftEmpty` as a pointer receiver, the call `it.IsLeftEmpty()` from a value receiver would take the address of the value copy — which works in Go but creates an unnecessary allocation.

**Resolution**: Migrated **all related methods together** in each type to avoid mixed receiver sets. This is the correct Go idiom — a type should have a consistent receiver convention, with pointer receivers only for methods that genuinely need them (mutation, nil checks, interface satisfaction).

---

## What Was Actually Migrated

| Type | File | Methods → Value | Methods Kept Pointer | Reason Kept |
|------|------|----------------|---------------------|-------------|
| `LeftRight` | `coredata/corestr/LeftRight.go` | 18 | 4 | `Ptr` (returns self), `IsEqual` (nil check), `Clone`/`Clear`/`Dispose` (mutation) |
| `LeftMiddleRight` | `coredata/corestr/LeftMiddleRight.go` | 22 | 3 | `Clone` (returns new ptr), `Clear`/`Dispose` (mutation + nil check) |
| `ExpectingRecord` | `errcore/ExpectingRecord.go` | 6 | 0 | All methods are pure formatters, no nil checks needed |

## What Was NOT Migrated (and Why)

| Package/Type | Reason |
|-------------|--------|
| `corestr/Collection` | Embeds `sync.Mutex` |
| `corestr/Hashmap` | Embeds `sync.Mutex` |
| `corestr/Hashset` | Embeds `sync.Mutex` |
| `corestr/CharCollectionMap` | Embeds `sync.Mutex` |
| `corestr/CharHashsetMap` | Embeds `sync.Mutex` |
| `corestr/LinkedList` | Embeds `sync.Mutex` |
| `corestr/LinkedCollections` | Embeds `sync.Mutex` |
| `corepayload/Attributes` | Nil-guard contagion via `IsEmpty()` |
| `corepayload/PagingInfo` | All methods use `it == nil` guards |
| `corepayload/PayloadWrapper` | Mutation via nested calls, returns internal pointers |
| `corepayload/PayloadsCollection` | Has Items slice, mutation methods |
| `errcore/RawErrCollection` | Has Items slice, mutation methods |

---

## Migration Criteria (for future audits)

A method is safe to migrate to a value receiver when **ALL** of these hold:

1. ✅ The struct does **not** embed `sync.Mutex` or `sync.RWMutex`
2. ✅ The method does **not** check `it == nil` or `it != nil`
3. ✅ The method does **not** assign to any field (`it.field = ...`)
4. ✅ The method does **not** call `append(it.slice, ...)` or `delete(it.map, ...)`
5. ✅ The method does **not** call `it.Lock()` / `it.Unlock()`
6. ✅ The method does **not** satisfy an interface that requires pointer receivers
7. ✅ All methods called on `it` within the body are also value receivers (or being migrated together)
