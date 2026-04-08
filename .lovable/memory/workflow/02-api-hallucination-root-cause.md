# Root Cause: AI Hallucination of Go API Signatures

## Status: 🟡 Reference — Coverage complete, guardrails remain mandatory for future test work

## The One Root Cause

**The AI hallucinates Go API signatures when writing coverage tests.** Every blocked-package incident traces back to this single failure mode. The AI generates code calling methods that don't exist, using wrong parameter counts, assuming wrong return types, or inventing enum constants — even when it has read the source files moments before.

## Why It Keeps Repeating

1. The AI's language model produces "plausible-looking" Go code that matches naming conventions but not actual APIs.
2. Reading a file doesn't guarantee accurate recall — the AI's generated code drifts from what it read.
3. Bulk file generation amplifies the problem: one wrong assumption propagates to dozens of call sites.
4. The "write tests, move on" pattern means errors are only caught when the user runs `./run.ps1 PC`.

## Concrete Examples (recurring patterns)

| Hallucinated Call | Actual API | Category |
|---|---|---|
| `corestr.New.Hashmap.StringsKv("k","v")` | `corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key:"k",Value:"v"})` | Non-existent method |
| `coredynamic.NewMapAnyItems()` | `coredynamic.NewMapAnyItems(0)` (requires capacity int) | Wrong param count |
| `stringcompareas.EqualMatch` | `stringcompareas.Equal` | Non-existent constant |
| `bc.AddBytes([]byte("x"))` | `bc.Add([]byte("x"))` | Non-existent method |
| `corejson.JsonString(x)` single return | `corejson.JsonString(x)` returns `(string, error)` | Wrong return count |
| `mr.Has("key")` | No `Has` method on `MapResults`; use `HasAnyItem()` | Non-existent method |
| `enumimpl.New.BasicByte.Create(1,"Name")` | `Create(typeName, ranges, names, min, max)` — 5 params | Wrong param count + conceptual misunderstanding |
| `bb.Value()`, `bb.Name()` | BasicByte has `Max()`, `Min()`, `GetStringValue(byte)` — no Value/Name | Non-existent methods |

## Mandatory Guardrails

### Before writing ANY test call:
1. **Search the source file for the exact method name** using `code--search_files`
2. **Read the full function signature** including parameter types and return types
3. **Never assume a method exists** based on naming patterns from similar packages
4. **Never assume parameter counts** — always verify

### Before submitting any test file:
1. **Verify every unique API call** has been checked against source
2. **Check enum constants exist** by reading the consts block in the enum package
3. **Check constructor signatures** — many take capacity/options params
4. **Check return types** — value vs pointer, single vs multi-return

### Process:
1. Read source → List methods → Write test → Verify each call → Submit
2. Never write more than ONE test file before verification
3. If a method name "feels right" but wasn't explicitly found in source, DO NOT USE IT

## What NOT To Do
- ❌ Write tests against assumed APIs
- ❌ Generate bulk test files without per-file compilation check
- ❌ Assume `MethodName()` exists because `MethodNamePtr()` exists
- ❌ Assume `New.Type.Create(simpleArgs)` when Create often takes complex params
- ❌ Assume single return value when Go commonly returns `(value, error)`
- ❌ Invent enum constants like `EqualMatch` when the actual constant is `Equal`
