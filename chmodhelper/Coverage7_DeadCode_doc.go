package chmodhelper

// ══════════════════════════════════════════════════════════════════════════════
// Coverage7 — Dead Code & Platform-Conditional Gap Documentation
//
// This file documents accepted coverage gaps that cannot be tested:
//
// 1. PLATFORM-CONDITIONAL (Linux-only code, tests run on Windows)
// 2. DEAD CODE (logically unreachable defensive returns)
//
// Total: ~100 statements across these categories.
// ══════════════════════════════════════════════════════════════════════════════
//
// ── PLATFORM-CONDITIONAL GAPS (Linux-only) ───────────────────────────────────
//
// These functions contain code guarded by osconsts.IsWindows / osconsts.IsLinux.
// Since tests run on Windows, the Linux-specific branches are never reached.
//
// RwxWrapper.go:
//   - LinuxApplyRecursive (lines 328-345): Linux chmod -R via cmd
//   - ApplyRecursive (lines 368-442): filesystem.Walk + os.Chmod on each path
//   - applyLinuxRecursiveChmodUsingCmd (lines 445-472): exec.Command("chmod -R ...")
//   - getLinuxRecursiveCmdForChmod (lines 475-489): builds exec.Cmd
//   - applyLinuxChmodOnManyNonRecursive (lines 544-568): iterates locations
//   - ApplyLinuxChmodOnMany (lines 579-588): dispatches recursive/non-recursive
//   - applyLinuxChmodOnManyRecursive (lines 594-618): recursive variant
//   - applyLinuxChmodRecursiveManyContinueOnError (lines 624-642): continue-on-error
//   - applyLinuxChmodNonRecursiveManyContinueOnError (lines 648-662): continue-on-error
//   - IsRwxEqualLocation null fileInfo (line 700): os.Stat failure branch
//
// fileWriter.go:
//   - All (lines 126-136): Unix chmod mismatch/apply branch (osconstsinternal.IsWindows guard)
//
// chmodApplier.go:
//   - RwxPartial error (line 275): propagates from Linux chmod
//   - RwxStringApplyChmod error (line 310): propagates from Linux chmod
//   - RwxOwnerGroupOtherApplyChmod error (line 342): propagates from Linux chmod
//
// tempDirGetter.go:
//   - TempPermanent (line 32): Linux branch returns "/var/tmp/"
//
// ── DEAD CODE GAPS (logically unreachable) ───────────────────────────────────
//
// RwxWrapper.go:
//   - ToUint32Octal error (line 86): valid RwxWrapper always produces valid
//     octal string, so strconv.ParseUint never fails
//   - VerifyPaths (line 53): tested indirectly through other paths
//
// RwxInstructionExecutor.go:
//   - CompiledWrapper fallback (line 75): varWrapper is always fixed or var
//   - CompiledRwxWrapperUsingFixedRwxWrapper fallback (lines 93-97): same reason
//   - verifyChmodLocationsContinueOnError CompiledWrapper error (lines 204-208):
//     requires broken varWrapper state (unreachable via public API)
//   - verifyChmodLocationsNoContinue CompiledWrapper error (lines 235-241): same
//
// chmodVerifier.go:
//   - GetRwx9 empty return (line 166): os.FileMode.String() always > 9 chars
//
// SingleRwx.go:
//   - ToRwxOwnerGroupOther default panic (line 94-95): exhaustive enum switch
//
// ══════════════════════════════════════════════════════════════════════════════
