package corecmptests

// NOTE: The 9 uncovered statements in corecmp are all unreachable dead code.
//
// Functions: Byte, Integer, Integer8, Integer16, Integer32, Integer64, Time,
//            VersionSliceByte, VersionSliceInteger
//
// Each has a final `return corecomparator.NotEqual` after exhaustive
// if/else-if chains covering equal, less-than, and greater-than.
// Since these three conditions are logically exhaustive for comparable types,
// the fallthrough return is dead code — it exists only as a safety net.
//
// These 9 statements CANNOT be covered by any test input.
// Filed as known dead code in spec/13-app-issues/golang/41-unreachable-branches-corejson-coreonce.md.
