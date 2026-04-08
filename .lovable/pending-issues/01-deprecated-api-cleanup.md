# Pending: Deprecated API Cleanup (S-009)

## Status: Open — Blocked on external consumer audit

## Description
110 deprecated functions/methods across 30+ files need removal or sunset. Largest concentrations: `coreindexes/indexes.go` (21), `core.go` (13), `coredata/corestr/` (15+).

## Detailed Plan
See `.lovable/memory/workflow/04-s009-deprecated-audit-plan.md` (349 lines, categorized by package and risk level).

## Blocker
User must grep across all external auk-go repos to confirm no external consumers before removal.

## Next Steps
1. User runs external consumer audit
2. Remove internal-only deprecated items first
3. Refactor `corestr` to use `coregeneric` delegation
4. Batch removal with compile verification
