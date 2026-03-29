# Issue: ResultsPtrCollection.Clone Always Returns Empty

## Summary
`ResultsPtrCollection.Clone()` returns an empty collection regardless of source content.

## Root Cause
Line 779 in `ResultsPtrCollection.go`:
```go
if newResults.Length() == 0 {  // ← always true after UsingCap
    return newResults
}
```
`UsingCap(n)` creates `make([]*Result, 0, n)` — length 0, capacity n. So `newResults.Length()` is always 0, causing immediate return before copying items.

## Fix
Changed to `if it.Length() == 0 {` — checks the SOURCE collection, not the empty target.

## Affected
- `Test_Cov50_ResultsPtrColl_Clone_DeepClone`
- Any production code calling `ResultsPtrCollection.Clone()`
