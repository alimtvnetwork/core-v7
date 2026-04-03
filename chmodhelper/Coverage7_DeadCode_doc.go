package chmodhelper

// ══════════════════════════════════════════════════════════════════════════════
// Coverage7 — Dead Code & Platform-Conditional Gap Documentation
//
// This file documents accepted coverage gaps that cannot be tested:
//
// 1. PLATFORM-CONDITIONAL (Linux-only code, tests run on Windows)
// 2. DEAD CODE (logically unreachable defensive returns)
// 3. UNEXPORTED (dirCreator, SingleRwx construction — only in-package access)
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
//   - ApplyChmod (lines 227-255): os.Chmod on Unix, invalidPathErr
//   - ApplyChmodIf (lines 299-307): isApplyOnMismatch + IsChmod + ApplyChmod
//   - LinuxApplyRecursive (lines 328-345): Linux chmod -R via cmd
//   - ApplyRecursive (lines 368-442): filesystem.Walk + os.Chmod on each path
//   - applyLinuxRecursiveChmodUsingCmd (lines 445-472): exec.Command("chmod -R ...")
//   - getLinuxRecursiveCmdForChmod (lines 475-489): builds exec.Cmd
//   - applyLinuxChmodOnManyNonRecursive (lines 544-568): iterates locations
//   - ApplyLinuxChmodOnMany (lines 579-588): dispatches recursive/non-recursive
//   - applyLinuxChmodOnManyRecursive (lines 594-618): recursive variant
//   - applyLinuxChmodRecursiveManyContinueOnError (lines 624-642): continue-on-error
//   - applyLinuxChmodNonRecursiveManyContinueOnError (lines 648-662): continue-on-error
//
// RwxVariableWrapper.go:
//   - ApplyRwxOnLocations (lines 186-218): ApplyChmod calls skipped on Windows
//   - IsEqualUsingLocation (lines 295-296): os.Stat + chmod comparison
//   - IsEqualUsingFileInfo (lines 309-310): chmod comparison
//
// fileWriter.go:
//   - All (lines 126-136): Unix chmod mismatch/apply branch (osconstsinternal.IsWindows guard)
//   - cleanUpErr (line 97): RemoveIf error propagation
//
// chmodApplier.go:
//   - RwxStringApplyChmod error (line 310): propagates from Linux chmod
//   - RwxOwnerGroupOtherApplyChmod error (line 342): propagates from Linux chmod
//
// chmodVerifier.go:
//   - IsEqualRwxFullSkipInvalid (lines 74-76): tested but Unix chmod needed
//   - IsEqualSkipInvalid (lines 99-101): tested but Unix chmod needed
//   - PathIf (lines 220-222): delegates to RwxFull (Unix)
//   - RwxFull mismatch (lines 274-276): requires Unix chmod state
//   - PathsUsingPartialRwxOptions (lines 313-315): error from NewRwxVariableWrapper
//   - PathsUsingRwxFull return nil (line 405): no-error path on non-continue
//   - GetExistingRwxWrapper bypass (line 359-364): platform chmod
//
// fwChmodApplier.go:
//   - OnDiffFile (lines 60-62): IsEqual requires Unix chmod
//
// fwChmodVerifier.go:
//   - IsEqualFile (lines 13-17): delegates to ChmodVerify.IsEqual (Unix)
//
// tempDirGetter.go:
//   - TempPermanent (line 32): Linux branch returns "/var/tmp/"
//
// dirCreator.go:
//   - Chmod != default branch (lines 69-77): os.Chmod on Unix
//   - ByChecking MkdirAll error (lines 108-115): filesystem error
//   - ByChecking Chmod error (lines 122-128): os.Chmod on Unix
//
// CreateDirWithFiles.go:
//   - removeDirErr (line 24-26): removeDirIf error
//   - fileManipulateErr (lines 62-68): Close() error
//   - chmodErr (lines 75-81): os.Chmod on Unix
//
// SimpleFileReaderWriter.go:
//   - errorWrapFilePath (line 133): error wrapping on write failure
//
// ── DEAD CODE GAPS (logically unreachable) ───────────────────────────────────
//
// RwxWrapper.go:
//   - ToUint32Octal error (line 86): valid RwxWrapper always produces valid
//     octal string, so strconv.ParseUint never fails
//
// RwxInstructionExecutor.go:
//   - CompiledRwxWrapperUsingFixedRwxWrapper fallback (lines 93-97): varWrapper
//     is always fixed or var type
//   - ApplyOnPath compiledErr (lines 119-125): requires broken varWrapper state
//
// chmodVerifier.go:
//   - GetRwx9 empty return (line 166): os.FileMode.String() always > 9 chars
//
// SingleRwx.go:
//   - ToRwxWrapper error on New.RwxWrapper (lines 145-147): valid rwx always parses
//   - ApplyOnMany error (lines 163-165): requires in-package SingleRwx construction
//   - ToRwxOwnerGroupOther default panic (line 94-95): exhaustive enum switch
//
// ── UNEXPORTED GAPS ──────────────────────────────────────────────────────────
//
// dirCreator (unexported struct via internalDirCreator):
//   - ByChecking, Default, Direct error paths: only accessible in-package
//   - Covered by in-package tests (Coverage1_Iteration12_test.go, etc.)
//
// ══════════════════════════════════════════════════════════════════════════════
