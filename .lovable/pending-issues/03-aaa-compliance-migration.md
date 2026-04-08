# Pending: AAA Compliance Migration (S-016)

## Status: Open — Large effort, low priority

## Description
33,150 non-compliant assertion calls across 393 files in 53 test packages need migration to `args.Map` + `ShouldBeEqual` pattern.

## Full Audit
See `.lovable/memory/workflow/07-aaa-compliance-audit.md` (733 lines).

## Next Steps
1. Prioritize by violation count (highest-violation packages first)
2. Migrate one package at a time
3. Compile + test verify after each package
4. Update audit report as packages are completed
