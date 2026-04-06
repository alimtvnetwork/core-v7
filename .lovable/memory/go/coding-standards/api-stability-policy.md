# API Stability Policy

**CRITICAL RULE — NEVER VIOLATE**

## Rule: Never Remove Exported/Public Functions

All exported (uppercase) functions, methods, types, and constants in this project are **public framework APIs**. External consumers depend on them.

**Even if a function has ZERO internal callers, it MUST NOT be removed.**

Zero internal callers does NOT mean zero external callers. This is a library/framework — external packages import and use these APIs.

## What "Deprecated" Means Here

A `// Deprecated:` comment is a hint to external users to migrate, but the function itself **must remain in the codebase indefinitely** until a major version bump with an explicit migration guide.

## Allowed Actions on Deprecated APIs

- ✅ Add `// Deprecated:` comments pointing to replacements
- ✅ Refactor internals to delegate to newer implementations
- ✅ Add wrapper logic that calls the new API internally
- ❌ **NEVER delete the function**
- ❌ **NEVER remove the exported symbol**
- ❌ **NEVER remove tests that cover deprecated APIs**

## S-009 Incident (2026-04-06)

Phase 1 of S-009 incorrectly deleted 85 exported functions, 34 files, and ~250 tests because they had "zero internal callers." This was wrong — they are framework APIs used by external consumers. The entire phase was reverted.

## Pointer Methods (`*Ptr` suffix)

`*Ptr` methods exist for null-safety checking patterns (e.g., check if object is nil, then check if field is nil). These are intentional API design choices and must be preserved.
