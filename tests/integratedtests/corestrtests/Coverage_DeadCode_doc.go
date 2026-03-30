package corestrtests

// ══════════════════════════════════════════════════════════════════════════════
// Coverage — coredata/corestr remaining gaps (90 uncovered lines)
//
// Most gaps are nil-receiver guards and defensive dead code across:
// - CharCollectionMap: nil map guards, json marshal errors
// - CharHashsetMap: nil hashset guards, defensive branches
// - Collection: nil receiver, json errors
// - CollectionsOfCollection: empty return guards
// - LinkedCollections/LinkedList/LinkedListNode: nil node guards
// - SimpleSlice: nil receiver, boundary guards
// - SimpleStringOnce: lazy init fallback
// - ValidValue: dead fallback return
//
// These require either internal tests or represent defensive dead code.
// Documented for future internal test coverage.
// ══════════════════════════════════════════════════════════════════════════════
