# Bracecheck Syntax Errors — 1142 Errors Across 11 Files

**Created:** 2026-03-29  
**Source:** `run.ps1 TC` bracecheck output  
**Location:** `tests/integratedtests/corestrtests/`  
**Status:** Open  

---

## Summary

After autofix runs (no fixable issues found), bracecheck reports **1142 errors in 11 files**. These are **not independent errors** — most are cascading failures from 4 root patterns. Fixing the root cause in each file will eliminate hundreds of downstream errors.

---

## Pattern 1: Inline Closure Returns Parsed as Argument Lists

**Error messages:** `missing ',' in argument list`, `missing ',' before newline in argument list`, `expected operand, found '}'`  
**Root cause:** Go's parser sees a single-line closure body containing `return x, y, z` and interprets the commas as argument separators of the outer function call — not as multiple return values.

**Example (broken):**
```go
filter := func(str string, i int) (string, bool, bool) { return str, true, false }
```

**Fix:** Split the closure body onto multiple lines:
```go
filter := func(str string, i int) (string, bool, bool) {
    return str, true, false
}
```

**Affected files:**
| File | Lines | Cascading errors |
|------|-------|-----------------|
| `Coverage42_Iteration38_test.go` | 301, 764, 780+ | ~17 |
| `Coverage56_Iteration52_test.go` | 195, 353, 372+ | ~23 |
| `Coverage57_Iteration53_test.go` | 744, 2146 | ~8 |
| `Src_Coverage12_Collection_Deep_test.go` | 265–267, 292–298+ | ~34 |
| `Src_Coverage18_LinkedList_test.go` | 292, 295, 298+ | ~46 |

---

## Pattern 2: Broken `safeTest` / For-Loop Closure Boundaries (Cascading 400+ Errors)

**Error messages:** `expected ';', found ','`, `expected operand, found '}'`, `expected operand, found 'range'`, `missing ',' in argument list` (on `var`, `if`, `for` keywords)  
**Root cause:** The `safeTest` closure or a test-case loop has a structural brace mismatch — typically an extra `},` or missing `}` before the loop. Once the parser loses brace context, every subsequent `for`, `var`, `if`, `:=` statement inside the closure is interpreted as part of an argument list, generating hundreds of cascading errors.

**Example (broken — Coverage72 line 127):**
```go
                tc.ShouldBeEqualMap(t, caseIndex, actual)
                },        // ← extra trailing comma on struct/closure close
                },        // ← second close — parser confused
        }
// next function starts, parser thinks it's still inside an argument list
```

**Fix:** Audit the brace structure around the test-case loop. The correct pattern is:
```go
func Test_CovS06_Verification(t *testing.T) {
    safeTest(t, "Test_CovS06_Verification", func() {
        for caseIndex, tc := range testCases {
            // Arrange / Act / Assert
            tc.ShouldBeEqualMap(t, caseIndex, actual)
        }
    })
}
```
Each `{` must have exactly one matching `}`. Remove stray `,` after closing braces that aren't inside composite literals.

**Affected files:**
| File | First error line | Cascading errors |
|------|-----------------|-----------------|
| `Coverage72_S06_CharCollectionMap_test.go` | 127 | **445** |
| `Coverage73_S07_CharHashsetMap_test.go` | 74 | **430** |

---

## Pattern 3: Statements Using Commas Instead of Newlines/Semicolons

**Error messages:** `missing ',' in argument list`, `expected '==', found '='`  
**Root cause:** Inside a `callPanics` closure, statements are separated by commas instead of newlines. Go interprets the comma-separated lines as a single expression list, so `:=` and `=` are invalid operators in that context.

**Example (broken — Src_Coverage19 line 553):**
```go
noPanic := !callPanicsSrcC19(func() {
    var nNil *corestr.LinkedCollectionNode,    // ← comma instead of newline
    _ = nNil.IsEmpty(),                        // ← comma
    n1 := &corestr.LinkedCollectionNode{...},  // ← comma
    _ = n1.IsEmpty(),                          // ← comma
})
```

**Fix:** Replace trailing commas with nothing (normal statement termination):
```go
noPanic := !callPanicsSrcC19(func() {
    var nNil *corestr.LinkedCollectionNode
    _ = nNil.IsEmpty()
    n1 := &corestr.LinkedCollectionNode{Element: nil}
    _ = n1.IsEmpty()
})
```

**Affected files:**
| File | Lines | Cascading errors |
|------|-------|-----------------|
| `Src_Coverage02_Collection_test.go` | 37–42, 194–195+ | ~18 |
| `Src_Coverage10_Full_test.go` | 1015–1021 (`callPanicsSrcC10` body) | ~4 |
| `Src_Coverage19_LinkedCollections_test.go` | 553–559+ | ~43 |

---

## Pattern 4: EOF Brace Mismatches / Unclosed Blocks

**Error messages:** `expected '}', found 'EOF'`, `expected ')', found 'EOF'`  
**Root cause:** A `{` block (function body or `safeTest` closure) was never closed. The parser reaches end-of-file while still expecting closing braces. Usually caused by one of the above patterns removing a `}` or `})` that was structurally necessary.

**Example (Src_Coverage01 line 532):**
```
expected '}', found 'EOF'
```

**Fix:** Count opening and closing braces for the entire file. Add the missing `}` or `})` at the correct nesting level. Often this is a missing final `}` for the last function, or a missing `})` to close a `safeTest` call.

**Affected files:**
| File | EOF line | Cascading errors |
|------|----------|-----------------|
| `Src_Coverage01_Helpers_test.go` | 532 | ~10 |
| `Coverage57_Iteration53_test.go` | 2662 | ~8 (combined with Pattern 1) |
| `Src_Coverage10_Full_test.go` | 1022 | ~4 (combined with Pattern 3) |

---

## File-Level Summary

| File | Primary Pattern(s) | Reported Errors | Estimated Root Fixes |
|------|-------------------|----------------|---------------------|
| `Coverage42_Iteration38_test.go` | 1 | 25 | ~8 inline closures to split |
| `Coverage56_Iteration52_test.go` | 1 | 31 | ~10 inline closures to split |
| `Coverage57_Iteration53_test.go` | 1, 4 | 16 | ~3 closures + EOF brace |
| `Coverage72_S06_CharCollectionMap_test.go` | 2 | 445 | 1 brace boundary fix |
| `Coverage73_S07_CharHashsetMap_test.go` | 2 | 430 | 1 brace boundary fix |
| `Src_Coverage01_Helpers_test.go` | 4 | 10 | 1 missing closing brace |
| `Src_Coverage02_Collection_test.go` | 3 | 26 | ~5 comma→newline fixes |
| `Src_Coverage10_Full_test.go` | 2, 3, 4 | 12 | brace fix + comma→newline |
| `Src_Coverage12_Collection_Deep_test.go` | 1, 2 | 42 | ~6 closures + 1 brace |
| `Src_Coverage18_LinkedList_test.go` | 1 | 54 | ~15 inline closures to split |
| `Src_Coverage19_LinkedCollections_test.go` | 2, 3 | 51 | brace fix + comma→newline |
| **Total** | | **1142** | **~50 manual edits** |

---

## Autofix Enhancement Opportunities

The current autofix (Rules A–F) does not handle these patterns because:

1. **Pattern 1** requires understanding that `{ return x, y, z }` on one line is a closure body, not an argument list. Autofix would need to detect single-line func literals with `return` and split them.
2. **Pattern 2** requires structural brace-depth analysis across the entire `safeTest` closure to find the misplaced `,` or missing `}`.
3. **Pattern 3** requires detecting that lines ending with `,` inside a `func() { ... }` block are statements, not composite literal fields.
4. **Pattern 4** is a consequence of 1–3 and resolves automatically once the root cause is fixed.

**Recommendation:** Patterns 1 and 3 are amenable to new autofix rules. Pattern 2 requires manual inspection of the specific brace structure in Coverage72/73.
