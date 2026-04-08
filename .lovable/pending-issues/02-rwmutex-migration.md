# Pending: sync.Mutex → sync.RWMutex Migration (S-013)

## Status: Open — Ready to start (benchmarks available)

## Description
27 `sync.Mutex` usages found. Read-heavy collection types (Collection, Hashmap, Hashset) may benefit from `sync.RWMutex` for concurrent read performance.

## Prerequisites
- ✅ S-010 benchmarks complete (baseline available)

## Next Steps
1. Audit each of the 27 mutex usages
2. Identify read-heavy vs write-heavy patterns
3. Migrate candidates to `RWMutex` (use `RLock`/`RUnlock` for read methods)
4. Benchmark before/after to confirm improvement
5. `./run.ps1 TC` verification
