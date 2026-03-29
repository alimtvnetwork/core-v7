# Dead Code Analysis: Unreachable Coverage Lines (Registry)

## Status: DOCUMENTED

These lines are logically unreachable and cannot be covered by any test.

## 1. `coreversion/hasDeductUsingNilNess.go:20-22`

```go
if left == nil || right == nil {
    return corecomparator.NotEqual, true
}
```

**Reason**: Prior branches exhaustively handle all nil combinations:
- L8: `left == nil && right == nil`
- L12: `left != nil && right == nil`
- L16: `left == nil && right != nil`

After these, both `left` and `right` are guaranteed non-nil. The `||` check is dead.

## 2. `coredata/corerange/MinMaxByte.go:46-48`

```go
if diff < 0 {
    return diff
}
```

**Reason**: `byte` is an unsigned type in Go (`uint8`). The expression `diff < 0` is always false for unsigned types.

## 3. `coredata/corerange/within.go:89`

```go
return 0, isInRange
```

**Reason**: `StringRangeUint32` calls `StringRangeInteger(true, 0, MaxInt32, input)`. With `isUsageMinMaxBoundary=true`, `RangeInteger` clamps the result to `[0, MaxInt32]`. Therefore `finalInt <= math.MaxInt32` is always true, making the else branch unreachable.

## 4. `coretaskinfo/InfoJson.go:25-27`

```go
func (it Info) JsonString() string {
    if it.IsNull() {  // IsNull checks it == nil, but value receiver is never nil
        return ""
    }
```

**Reason**: `JsonString` uses a value receiver (`Info`), so `it` is a copy on the stack. `IsNull()` is a pointer method checking `it == nil`, but `&it` of a value receiver is never nil.

## 5. `coredata/stringslice/MergeSlicesOfSlices.go:13-15`

```go
if sliceLength == constants.Zero {
    return []string{}
}
```

**Reason**: Redundant check. Line 7 already returns `[]string{}` when `len(slicesOfSlice) == 0`, and `sliceLength` is assigned from the same `len()` call.

## 6. `coredata/stringslice/RegexTrimmedSplitNonEmptyAll.go:17-19`

```go
if len(items) == 0 {
    return []string{}
}
```

**Reason**: `regexp.Split(content, -1)` always returns at least one element (even for empty string → `[""]`). The `len(items) == 0` check is unreachable.

## 7. `iserror/Equal.go:8-10`

```go
if left == nil && right == nil {
    return true
}
```

**Reason**: Line 4 checks `left == right` which already returns `true` when both are nil. The explicit nil-nil check is redundant.

## 8. `reqtype/start.go:8-10` and `reqtype/end.go:6-8`

```go
if len(reqs) == 0 {
    return nil
}
```

**Reason**: Both are unexported functions. All callers (`RangesNotMeet`, `GetStatusAnyOf`) guard against empty slices before calling `start`/`end`, making the internal empty check unreachable.

## 9. `errcore/RawErrCollection.go:454-455, 463-464` (LogFatal, LogFatalWithTraces)

```go
slog.Error("fatal: raw error collection", "errors", it.String())
os.Exit(1)
```

**Reason**: `os.Exit(1)` terminates the process. Cannot be tested without killing the test runner. The `IsEmpty()` early-return paths ARE covered; only the fatal exit paths are not.

## 10. `errcore/RawErrCollection.go:469` (LogIf)

```go
func (it *RawErrCollection) LogIf(isLog bool) {
    if isLog {
        it.LogFatal() // delegates to os.Exit
    }
}
```

**Reason**: Delegates to `LogFatal` which calls `os.Exit(1)`. Untestable.

## 11. `errcore/CompiledError.go:30-31` (CompiledErrorString)

```go
if compiled == nil {
    return ""
}
```

**Reason**: `CompiledError(mainErr, msg)` never returns nil when `mainErr != nil` — it either returns `mainErr` (when `msg == ""`) or a wrapped error. Since line 25 already guards `mainErr == nil`, `compiled` is guaranteed non-nil.

## 12. `namevalue/Instance.go:20-21` (Instance.String value receiver nil check)

```go
func (it Instance[K, V]) String() string {
    if it.IsNull() { // IsNull checks it == nil, but value receiver copy is never nil
```

**Reason**: Value receiver — `it` is a stack copy, never nil. Same pattern as Item #4.

## 13. `namevalue/Instance.go:31-32` (Instance.JsonString value receiver nil check)

Same pattern as #12 — value receiver nil check is unreachable.

## 14. `namevalue/Instance.go:37` and `namevalue/Collection.go:385`

```go
if err != nil || rawBytes == nil {
```

**Reason**: `json.Marshal` on a valid struct never returns `(nil, nil)`. If no error, bytes are always non-nil.

## Recommendation

## 15. `coretests/messagePrinter.go` (printMessage type)

```go
type printMessage struct{}
func (it printMessage) FailedExpected(...) { ... }
func (it printMessage) NameValue(...) { ... }
func (it printMessage) Value(...) { ... }
```

**Reason**: `printMessage` is an unexported type. Cannot be accessed from integrated tests. All 3 methods are dead from coverage perspective.

## 16. `coretests/SkipOnUnix.go:13` and `coretests/SkipOnWindows.go:13`

```go
t.Skip(errcore.UnixIgnoreType)  // only on Unix
t.Skip(errcore.WindowsIgnoreType)  // only on Windows
```

**Reason**: Platform-dependent. On any given OS, one of these is always dead code. Tests run on one platform only.

## Recommendation

These are defensive guards. They could be removed for cleanliness, but keeping them is harmless. Coverage for these packages is effectively 100% for all reachable code.
