# Deprecated API Cleanup Guidelines

## S-009 Incident Summary

The S-009 cleanup (Phase 1) was **fully reverted** because it deleted public framework methods that external consumers depend on. The mistake: treating "zero internal callers" as "safe to delete."

## Rules for Any Future Cleanup

1. **NEVER delete exported functions** — they are public API, period
2. **NEVER delete files containing only exported functions** — same reason
3. **NEVER delete tests for exported functions** — they verify the public contract
4. **Internal-only helpers** (unexported/lowercase) may be cleaned up if truly unused
5. **Deprecated APIs** should be kept forever with `// Deprecated:` comments
6. The correct approach is to **refactor internals** (e.g., delegate to `coregeneric`) while keeping the exported signature intact

## What IS Allowed

- Removing truly internal (unexported) dead code
- Consolidating internal implementation details
- Adding deprecation comments
- Refactoring a deprecated method to delegate to its replacement internally
