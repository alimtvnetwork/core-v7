# S-012 Pointer Receiver Revert — Why Value Receivers Were Wrong

## What Went Wrong

The S-012 audit migrated 46 methods from pointer receivers (`*Type`) to value receivers (`Type`) across three files:

- `coredata/corestr/LeftRight.go` — 18 methods
- `coredata/corestr/LeftMiddleRight.go` — 22 methods
- `errcore/ExpectingRecord.go` — 6 methods

The reasoning was: "these methods are read-only and don't modify state, so value receivers are fine."

**This reasoning was incorrect.**

## Why Pointer Receivers Are Required

### 1. Nil-Safety Is a Feature, Not an Optimization

Pointer receivers allow **nil-guard checks** at the top of every method:

```go
func (it *LeftRight) IsLeftEmpty() bool {
    return it == nil || it.Left == ""
}
```

With a value receiver, calling `IsLeftEmpty()` on a nil pointer **panics** because Go dereferences the pointer to copy the value before the method even runs. With a pointer receiver, the method receives nil safely and can return a sensible default.

### 2. The "Object Is Null → Field Is Null" Pattern

The project follows a defensive pattern: if the object itself is nil, then all its fields are considered empty/null by convention. This enables safe chaining:

```go
var lr *LeftRight // nil
lr.IsLeftEmpty()           // returns true (safe)
lr.HasSafeNonEmpty()       // returns false (safe)
lr.LeftBytes()             // returns nil (safe)
```

With value receivers, every one of these would panic on a nil pointer.

### 3. Consistency Across the Codebase

Types like `Attributes`, `PagingInfo`, `PayloadWrapper`, `RawErrCollection`, and every mutex-embedding collection all use pointer receivers with nil guards. The three migrated types were the **only** ones that broke this convention after S-012.

### 4. Read-Only ≠ Safe for Value Receiver

The original audit criteria ("doesn't modify state") missed the critical point: **pointer receivers aren't just about mutation — they're about nil-safety**. A method that reads fields can still panic if the receiver is nil and uses a value receiver.

## What Was Fixed

All 46 methods were reverted back to pointer receivers with proper nil-guard checks added:

| File | Methods Reverted | Nil Guard Pattern |
|------|-----------------|-------------------|
| `coredata/corestr/LeftRight.go` | 18 | `it == nil \|\| ...` for emptiness, `it != nil && ...` for positive checks, `if it == nil { return zero }` for value returns |
| `coredata/corestr/LeftMiddleRight.go` | 22 | Same pattern |
| `errcore/ExpectingRecord.go` | 6 | `if it == nil { return "" }` / `return nil` |

### Nil Guard Patterns Used

**Boolean emptiness checks** — return `true` if nil (nil is empty):
```go
func (it *LeftRight) IsLeftEmpty() bool {
    return it == nil || it.Left == ""
}
```

**Boolean positive checks** — return `false` if nil (nil has nothing):
```go
func (it *LeftRight) HasValidNonEmptyLeft() bool {
    return it != nil && it.IsValid && !it.IsLeftEmpty()
}
```

**Value returns** — return zero value if nil:
```go
func (it *LeftRight) LeftBytes() []byte {
    if it == nil {
        return nil
    }
    return []byte(it.Left)
}
```

**Error returns** — return nil error if nil:
```go
func (it *ExpectingRecord) Error(actual any) error {
    if it == nil {
        return nil
    }
    return errors.New(it.Message(actual))
}
```

## Updated Rule

**All methods on pointer-returned types MUST use pointer receivers**, regardless of whether they modify state. The only exception is `NonPtr()` which by definition returns a copy, but even that now has a nil guard.

This rule supersedes the previous S-012 criteria. The receiver-selection-criteria memory has been updated accordingly.
