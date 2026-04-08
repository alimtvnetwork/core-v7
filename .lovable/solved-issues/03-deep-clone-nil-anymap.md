# Solved: Deep Clone Production Bug — Nil AnyMap

## Resolved: 2026-03-15

## Root Cause
`Attributes.deepClonePtr()` unconditionally called `it.AnyKeyValuePairs.ClonePtr()`. When `AnyKeyValuePairs` is nil, `MapAnyItems.ClonePtr()` returns `defaulterr.NilResult` — a non-nil error that propagated up, causing `Clone(deep=true)` to fail.

## Fix
Added nil guard: only call `AnyKeyValuePairs.ClonePtr()` if non-nil; otherwise pass nil through.

## File Changed
- `coredata/corepayload/Attributes.go` — `deepClonePtr()` method

## Learning
Always guard nil receivers before delegating clone operations, especially when the delegated method returns a sentinel error for nil input.
