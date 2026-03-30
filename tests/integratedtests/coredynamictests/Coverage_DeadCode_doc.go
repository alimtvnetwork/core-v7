package coredynamictests

// ══════════════════════════════════════════════════════════════════════════════
// Coverage — coredata/coredynamic remaining gaps (47 uncovered lines)
//
// Most gaps fall into these categories:
// 1. Nil receiver guards (CollectionLock.LengthLock, etc.) — dead code,
//    Lock() would panic before the nil check
// 2. json.Marshal error branches (AnyCollection.JsonString, DynamicJson) — dead code
// 3. ReflectSetFromTo error branches — requires specific reflect failures
// 4. MapAnyItems nil/empty guards — defensive dead code
// 5. CastTo type assertion branches — requires specific uncastable types
//
// These require either internal tests (unexported methods) or represent
// defensive dead code. Documented for future internal test coverage.
// ══════════════════════════════════════════════════════════════════════════════
