# S-009: Deprecated API Cleanup — Detailed Audit Plan

**Total deprecated items**: 110 across 64 files  
**Packages affected**: 16 packages  
**Strategy**: Verify replacement exists → check for internal callers → remove if safe, migrate callers if not

---

## Category 1: `core.go` — Type-Specific Empty Slice/Map Creators (11 functions)

**Risk**: LOW — generic replacements exist in `generic.go`  
**Internal callers**: None found (only constants/arrayvarsptr.go references similarly-named constants, not these functions)

| # | Deprecated Function | Replacement | Action |
|---|---|---|---|
| 1 | `EmptyAnysPtr() *[]any` | `EmptySlicePtr[any]()` | DELETE |
| 2 | `EmptyFloat32Ptr() *[]float32` | `EmptySlicePtr[float32]()` | DELETE |
| 3 | `EmptyFloat64Ptr() *[]float64` | `EmptySlicePtr[float64]()` | DELETE |
| 4 | `EmptyBoolsPtr() *[]bool` | `EmptySlicePtr[bool]()` | DELETE |
| 5 | `EmptyIntsPtr() *[]int` | `EmptySlicePtr[int]()` | DELETE |
| 6 | `EmptyBytePtr() []byte` | `[]byte{}` (built-in) | DELETE |
| 7 | `EmptyStringsMapPtr() *map[string]string` | `EmptyMapPtr[string, string]()` | DELETE |
| 8 | `EmptyStringToIntMapPtr() *map[string]int` | `EmptyMapPtr[string, int]()` | DELETE |
| 9 | `EmptyStringsPtr() *[]string` | `EmptySlicePtr[string]()` | DELETE |
| 10 | `EmptyPointerStringsPtr() *[]*string` | `EmptySlicePtr[*string]()` | DELETE |
| 11 | `StringsPtrByLength(length int) *[]string` | `SlicePtrByLength[string](length)` | DELETE |

**Before (core.go)**:
```go
// Deprecated: Use EmptySlicePtr[any]() instead.
func EmptyAnysPtr() *[]any {
    s := make([]any, 0); return &s
}
// ... 10 more similar functions
```

**After (core.go)**: All 11 functions removed. Only `generic.go` remains with generic versions.

---

## Category 2: `generic.go` — Type-Specific Generic Slice Creators (3 functions)

**Risk**: LOW — generic replacements exist in same file  
**Internal callers**: None found

| # | Deprecated Function | Replacement | Action |
|---|---|---|---|
| 12 | `EmptySlicePtrAny() *[]any` | `EmptySlice[any]()` | DELETE |
| 13 | `SlicePtrByLengthAny(length) *[]any` | `SliceByLength[any](length)` | DELETE |
| 14 | `SlicePtrByCapacityAny(length, cap) *[]any` | `SliceByCapacity[any](length, cap)` | DELETE |

**Before (generic.go)**:
```go
// Deprecated: Use EmptySlice instead.
func EmptySlicePtrAny() *[]any { ... }
```

**After**: These 3 functions removed from `generic.go`.

---

## Category 3: `coredata/stringslice/` — Ptr-Suffix Functions (22 files to DELETE)

**Risk**: LOW for most, MEDIUM for 6 that have internal callers  
**Internal callers**: 6 functions are called from `corestr/Collection.go` and `corestr/Hashset.go`

### 3a: No Internal Callers — Safe to Delete Entire Files (16 files)

| # | File | Deprecated Function | Replacement |
|---|---|---|---|
| 15 | `ClonePtr.go` | `ClonePtr(slice)` | `Clone(slice)` |
| 16 | `EmptyPtr.go` | `EmptyPtr()` | `Empty()` |
| 17 | `FirstLastDefaultPtr.go` | `FirstLastDefaultPtr(slice)` | `FirstLastDefault(slice)` |
| 18 | `FirstLastDefaultStatusPtr.go` | `FirstLastDefaultStatusPtr(slice)` | `FirstLastDefaultStatus(slice)` |
| 19 | `FirstOrDefaultPtr.go` | `FirstOrDefaultPtr(slice)` | `FirstOrDefault(slice)` |
| 20 | `FirstPtr.go` | `FirstPtr(slice)` | `First(slice)` |
| 21 | `HasAnyItemPtr.go` | `HasAnyItemPtr(slice)` | `HasAnyItem(slice)` |
| 22 | `IsEmptyPtr.go` | `IsEmptyPtr(slice)` | `IsEmpty(slice)` |
| 23 | `LastIndexPtr.go` | `LastIndexPtr(slice)` | `LastIndex(slice)` |
| 24 | `LastOrDefaultPtr.go` | `LastOrDefaultPtr(slice)` | `LastOrDefault(slice)` |
| 25 | `LastPtr.go` | `LastPtr(slice)` | `Last(slice)` |
| 26 | `LastSafeIndexPtr.go` | `LastSafeIndexPtr(slice)` | `LastSafeIndex(slice)` |
| 27 | `LengthOfPointer.go` | `LengthOfPointer(slice)` | `len(slice)` |
| 28 | `MakeDefaultPtr.go` | `MakeDefaultPtr(cap)` | `MakeDefault(cap)` |
| 29 | `MakeLenPtr.go` | `MakeLenPtr(len)` | `MakeLen(len)` |
| 30 | `MakePtr.go` | `MakePtr(len, cap)` | `Make(len, cap)` |
| 31 | `SlicePtr.go` | `SlicePtr(args)` | No longer needed |
| 32 | `TrimmedEachWordsPtr.go` | `TrimmedEachWordsPtr(slice)` | `TrimmedEachWords(slice)` |

### 3b: Has Internal Callers — Must Migrate First (6 files)

| # | File | Deprecated Function | Internal Callers | Action |
|---|---|---|---|---|
| 33 | `NonEmptyJoinPtr.go` | `NonEmptyJoinPtr(slice, joiner)` | `Collection.go:2071`, `Hashset.go:1218` | Migrate callers → `NonEmptyJoin` → DELETE |
| 34 | `NonEmptySlicePtr.go` | `NonEmptySlicePtr(slice)` | `Collection.go:1552` | Migrate caller → `NonEmpty` → DELETE |
| 35 | `NonWhitespaceJoinPtr.go` | `NonWhitespaceJoinPtr(slice, joiner)` | `Collection.go:2080`, `Hashset.go:1227` | Migrate callers → `NonWhitespaceJoin` → DELETE |
| 36 | `NonWhitespacePtr.go` | `NonWhitespacePtr(slice)` | `Collection.go:1562` | Migrate caller → `NonWhitespace` → DELETE |
| 37 | `SafeIndexAtUsingLastIndexPtr.go` | `SafeIndexAtUsingLastIndexPtr(...)` | None found in non-test | DELETE |
| 38 | `SafeIndexAtWithPtr.go` | `SafeIndexAtWithPtr(...)` | None found in non-test | DELETE |
| 39 | `SafeIndexesDefaultPtr.go` | `SafeIndexesDefaultPtr(...)` | None found in non-test | DELETE |
| 40 | `SafeRangeItemsPtr.go` | `SafeRangeItemsPtr(...)` | None found in non-test | DELETE |

**Before (Collection.go:2071)**:
```go
return stringslice.NonEmptyJoinPtr(it.items, joiner)
```
**After**:
```go
return stringslice.NonEmptyJoin(it.items, joiner)
```

---

## Category 4: `internal/strutilinternal/` — Ptr-Suffix Functions (4 files to DELETE)

**Risk**: LOW — no internal callers found  

| # | File | Deprecated Function | Replacement |
|---|---|---|---|
| 41 | `NonEmptySlicePtr.go` | `NonEmptySlicePtr(slice)` | `NonEmptySlice(slice)` |
| 42 | `NonWhitespaceJoinPtr.go` | `NonWhitespaceJoinPtr(slice, joiner)` | `NonWhitespaceJoin(slice, joiner)` |
| 43 | `NonWhitespaceSlicePtr.go` | `NonWhitespaceSlicePtr(slice)` | `NonWhitespaceSlice(slice)` |
| 44 | `NonWhitespaceTrimSlicePtr.go` | `NonWhitespaceTrimSlicePtr(slice)` | `NonWhitespaceTrimSlice(slice)` |

---

## Category 5: `coredata/corestr/` — Ptr-Suffix Methods (14 methods across 8 files)

**Risk**: MEDIUM — some have internal callers and delegation chains  
**Internal callers**: `StringsOptions` heavily used (15+ callers), others minimal

| # | File | Deprecated Method | Replacement | Internal Callers |
|---|---|---|---|---|
| 45 | `Collection.go` | `NonEmptyListPtr()` | `NonEmptyList()` | None |
| 46 | `Collection.go` | `NonEmptyItemsPtr()` | `NonEmptyItems()` | None |
| 47 | `Collection.go` | `NonEmptyItemsOrNonWhitespacePtr()` | `NonEmptyItemsOrNonWhitespace()` | None |
| 48 | `Collection.go` | `ListPtr()` | `List()` or `Items()` | None |
| 49 | `Hashset.go` | `ListPtr()` | `List()` | None |
| 50 | `Hashmap.go` | `ValuesToLower()` | `KeysToLower()` | None |
| 51 | `LinkedCollectionNode.go` | `ListPtr()` | `List()` | None |
| 52 | `LinkedCollections.go` | `ToStringsPtr()` | `ToStrings()` | None |
| 53 | `LinkedCollections.go` | `ListPtr()` | `List()` | None |
| 54 | `LinkedList.go` | `AddStringsPtrToNode(...)` | `AddStringsToNode(...)` | None |
| 55 | `LinkedList.go` | `ListPtr()` | `List()` | None |
| 56 | `LinkedList.go` | `ListPtrLock()` | `ListLock()` | None |
| 57 | `LinkedListNode.go` | `AddStringsPtrToNode(...)` | `AddStringsToNode(...)` | None |
| 58 | `LinkedListNode.go` | `ListPtr()` | `List()` | None |
| 59 | `ValidValue.go` | `ValueBytesOncePtr()` | `ValueBytesOnce()` | None |
| 60 | `newSimpleSliceCreator.go` | `StringsPtr(...)` | `Strings(...)` | None |
| 61 | `newSimpleSliceCreator.go` | `StringsOptions(...)` | `Strings()` or `StringsClone()` | **15+ callers** ⚠️ |
| 62 | `LeftRight.go` | `LeftRightUsingSlicePtr(slice)` | `LeftRightUsingSlice(slice)` | None |

**Before (Collection.go)**:
```go
// Deprecated: Use NonEmptyList instead.
func (it *Collection) NonEmptyListPtr() *[]string {
    ...
}
```
**After**: Method removed entirely.

### ⚠️ HIGH-RISK: `StringsOptions` (item #61)

This deprecated method has **15+ internal callers** across:
- `Collection.go` (5 calls)
- `CharCollectionMap.go` (1 call)
- `CollectionsOfCollection.go` (1 call)
- `Hashmap.go` (5 calls)
- `Hashset.go` (3 calls)
- `LinkedCollections.go` (2 calls)

**Before**:
```go
func (it *newSimpleSliceCreator) StringsOptions(
    isCloneAdd bool,
    stringsItems []string,
) *Collection { ... }
```

**Migration**: All 15+ callers must be migrated to `Strings()` or `StringsClone()` based on the `isCloneAdd` boolean:
- `StringsOptions(false, items)` → `Strings(items)`
- `StringsOptions(true, items)` → `StringsClone(items)`

---

## Category 6: `coredata/coredynamic/` — Ptr-Suffix Methods (2 methods)

**Risk**: MEDIUM — both have internal delegation callers

| # | File | Deprecated Method | Replacement | Internal Callers |
|---|---|---|---|---|
| 63 | `AnyCollection.go` | `ListStringsPtr(isInclude)` | `ListStrings(isInclude)` | `AnyCollection.go:223` (self-delegation) |
| 64 | `DynamicCollection.go` | `ListStringsPtr()` | `ListStrings()` | `DynamicCollection.go:201` (self-delegation) |

**Before (AnyCollection.go:220-223)**:
```go
func (it *AnyCollection) ListStrings(isIncludeFieldName bool) []string {
    return it.ListStringsPtr(isIncludeFieldName)
}
```
**After**: Inline the logic from `ListStringsPtr` into `ListStrings`, then delete `ListStringsPtr`.

---

## Category 7: `coredata/corejson/` — Ptr-Suffix Methods (6 methods across 4 files)

**Risk**: LOW — no internal callers

| # | File | Deprecated Method | Replacement |
|---|---|---|---|
| 65 | `BytesCollection.go` | `GetAtSafePtr(...)` | `GetAtSafe(...)` |
| 66 | `BytesCollection.go` | `StringsPtr()` | `Strings()` |
| 67 | `MapResults.go` | `GetErrorsStringsPtr()` | `GetErrorsStrings()` |
| 68 | `MapResults.go` | `GetStringsPtr()` | `GetStrings()` |
| 69 | `ResultsPtrCollection.go` | `GetErrorsStringsPtr()` | `GetErrorsStrings()` |
| 70 | `ResultsPtrCollection.go` | `GetStringsPtr()` | `GetStrings()` |

---

## Category 8: `coredata/corejson/Result.go` — SafeValuesPtr (1 method)

| # | Deprecated Method | Replacement |
|---|---|---|
| 71 | `SafeValuesPtr()` | `SafeValues()` |

---

## Category 9: `coredata/coreonce/StringsOnce.go` — ValuesPtr (1 method)

| # | Deprecated Method | Replacement |
|---|---|---|
| 72 | `ValuesPtr()` | `Values()` |

---

## Category 10: `coredata/coreapi/` — Type Aliases (2 type aliases)

**Risk**: LOW — type aliases, so removing them only breaks external consumers who use the old name  

| # | File | Deprecated Type | Replacement |
|---|---|---|---|
| 73 | `GenericResponseResult.go` | `GenericResponseResult` (alias) | `TypedResponseResult[any]` |
| 74 | `SimpleGenericRequest.go` | `SimpleGenericRequest` (alias) | `TypedSimpleGenericRequest[any]` |

**Before**: 
```go
// Deprecated: Use TypedResponseResult[any] instead.
type GenericResponseResult = TypedResponseResult[any]
```
**After**: Entire file deleted.

---

## Category 11: `coreindexes/indexes.go` — Old Index Constants (21 constants)

**Risk**: LOW — simple constant renames  
**Internal callers**: None found

| # | Deprecated Constant | Replacement |
|---|---|---|
| 75-95 | `Zero` through `Twenty` | `Index0` through `Index20` |

**Before**:
```go
// Deprecated: Use Index0 instead.
Zero = Index0
// Deprecated: Use Index1 instead.
One = Index1
// ... through Twenty = Index20
```
**After**: All 21 deprecated constant aliases removed.

---

## Category 12: `coremath/` — Built-in Replacements (4 files to DELETE)

**Risk**: LOW — Go 1.21+ built-in `min()`/`max()` replaces these  
**Internal callers**: None found

| # | File | Deprecated Function | Replacement |
|---|---|---|---|
| 96 | `MaxByte.go` | `MaxByte(a, b)` | `max(a, b)` |
| 97 | `MaxFloat32.go` | `MaxFloat32(a, b)` | `max(a, b)` |
| 98 | `MinByte.go` | `MinByte(a, b)` | `min(a, b)` |
| 99 | `MinFloat32.go` | `MinFloat32(a, b)` | `min(a, b)` |

---

## Category 13: Miscellaneous (11 items)

| # | File | Deprecated Item | Replacement | Internal Callers | Action |
|---|---|---|---|---|---|
| 100 | `chmodhelper/RwxInstructionExecutors.go` | `ApplyOnPathsPtr(locations)` | `ApplyOnPaths(locations)` | **5 callers** in chmodhelper | Migrate callers → DELETE |
| 101 | `conditional/NilCheck.go` | `NilCheck(val, fallback)` | `NilVal[T](val, fallback)` | None | DELETE file |
| 102 | `constants/arrayvars.go` | `EmptyIntToPtrIntsMap` constant | Non-pointer map types | Check `arrayvarsptr.go` ref | DELETE constant |
| 103 | `converters/stringTo.go` | `Float64Conditional(...)` | `Float64Default(...)` | None | DELETE method |
| 104 | `converters/unsafeBytesTo.go` | `UnsafeBytesToStrings(...)` | None (broken function) | None | DELETE file |
| 105 | `coreinstruction/BaseTags.go` | `NewBaseTags(...)` | `NewTags(...)` | None | DELETE function |
| 106 | `coretests/AnyToBytesPtr.go` | `AnyToBytesPtr(...)` | `AnyToBytes(...)` | None | DELETE file |
| 107 | `errcore/RawErrorType.go` | `ValidationFailedConst` constant | `ValidationFailedType` | None | DELETE constant |
| 108 | `errcore/funcs.go` | `ErrFuncHandler` type | `ErrFunc` type | None | DELETE type |
| 109 | `core.go` | `StringsPtrByCapacity(len, cap)` | `SlicePtrByCapacity[string](len, cap)` | None | DELETE |
| 110 | `core.go` | `PointerStringsPtrByCapacity(len, cap)` | `SlicePtrByCapacity[*string](len, cap)` | None | DELETE |

---

## Execution Order

### Phase 1: Zero-Caller Deletions (85 items, ~40 files)
No internal callers — safe to delete immediately.
- Categories 1, 2, 3a, 4, 7, 8, 9, 10, 11, 12
- Misc items 101, 103, 104, 105, 106, 107, 108
- corestr items 45-50, 51-59, 60, 62

### Phase 2: Caller Migrations Then Deletions (25 items)
Must update internal callers first, then delete.
- **StringsOptions** (#61): 15+ callers → migrate to `Strings()`/`StringsClone()`
- **stringslice Ptr funcs** (#33-36): 6 callers in Collection.go/Hashset.go
- **ApplyOnPathsPtr** (#100): 5 callers in chmodhelper
- **ListStringsPtr** (#63-64): 2 self-delegation callers in coredynamic
- **EmptyIntToPtrIntsMap** (#102): 1 ref in arrayvarsptr.go

### Phase 3: Verify & Clean
- Run full test suite
- Remove any orphaned imports
- Version bump

---

## Summary

| Category | Items | Files Affected | Risk | Has Callers? |
|---|---|---|---|---|
| core.go empty creators | 13 | 2 | LOW | No |
| stringslice Ptr files | 22 | 22 (delete files) | LOW-MED | 6 have callers |
| strutilinternal Ptr files | 4 | 4 (delete files) | LOW | No |
| corestr Ptr methods | 18 | 8 | MED | `StringsOptions` has 15+ |
| coredynamic Ptr methods | 2 | 2 | MED | Self-delegation |
| corejson Ptr methods | 7 | 4 | LOW | No |
| coreapi type aliases | 2 | 2 (delete files) | LOW | No |
| coreindexes constants | 21 | 1 | LOW | No |
| coremath builtins | 4 | 4 (delete files) | LOW | No |
| chmodhelper | 1 | 1 | MED | 5 callers |
| Misc | 6 | 6 | LOW | No |
| **TOTAL** | **110** | **~56 files** | | |

### Files to Delete Entirely: ~34 files
### Files to Edit (remove methods/constants): ~22 files  
### Internal Callers to Migrate: ~28 call sites
