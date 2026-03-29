# Issue: Coverage23 Generic Type Parameter Mismatch

## Summary
`Coverage23_TypedFuncs_test.go` defined `testUserCov23` but called all generic functions with `[testUser, ...]` type params (from `TypedCollection_testcases.go`). Go generics require exact type match — named types are not interchangeable.

## Root Cause
AI-generated tests used assumed type `testUser` in generic instantiations while the collection contained `testUserCov23`. Go's type system treats these as distinct types.

## Fix
- Replaced all `[testUser,` → `[testUserCov23,` in generic function calls
- Replaced all `func(u testUser)` → `func(u testUserCov23)` in lambda signatures
- Replaced `[]testUser{` → `[]testUserCov23{` in slice literals

## Also Fixed
- `Coverage24_TypedWrapperMethods_test.go`: `makeCollection()` → `makeCollectionCov23()` (undefined function)

## Impact
Blocked entire `corepayloadtests` package from compilation (50+ errors).
