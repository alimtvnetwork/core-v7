# Coverage Iteration Plan — 100% for All Non-Internal Packages

## Status: ✅ Complete — All 20 iterations executed
## Last Updated: 2026-03-20

## Strategy
- 2 packages per iteration, lowest coverage first
- **Skip `internal/*` packages — NEVER write coverage tests for internal packages** (see `.lovable/memory/testing/internal-packages-no-coverage.md`)
- After each iteration, list remaining items
- **CRITICAL**: Run `./run.ps1 PC` then `./run.ps1 TC` after fixes to validate

---

## Blocked Package Fixes (Just Completed)
11 packages were blocked due to compile errors from previous test additions.
All 11 have been fixed:
- bytetypetests: `Variant{}` → `new(bytetype.Variant)`
- chmodhelpertests: `IsExitOnInvalid` field → `IsSkipOnInvalid: false`; `hm.Add` → `hm.AddOrUpdate`
- converterstests: `converters.AnyItem` → `converters.AnyTo`
- corejsontests: func→bytesSerializer interface; nil→empty variadic; error→fmt.Stringer
- coremathtests: duplicate name + `coremath.IntegerOutOfRange` → `coremath.IsOutOfRange.Integer`
- corestrtests: duplicate `errForTest`; wrong signatures for 5 methods
- issettertests: `issetter.Names` → `OnlySupportedErr` path
- keymktests: `NewKey("x")` → `NewKey.Default("x")`; duplicate names; `NewOption` → `&Option{}`
- reqtypetests: removed non-existent `IsDynamicAction` method
- stringcompareastests: `nc.Compare` → `nc.IsCompareSuccess`
- versionindexestests: value→pointer receiver `&minor`

---

## run.ps1 Fix
Removed all root-level file writes. All logs now go to `data/coverage/` and `data/test-logs/` only.

---

## Iteration Plan (2 packages at a time, lowest coverage first)

### Iteration 1: reflectcore/reflectmodel (0.8%) + reqtype (1.3%)
- reflectmodel: 251 uncovered stmts — MethodProcessor invoke paths, nil receivers, type checks
- reqtype: 227 uncovered stmts — Is* methods, enum operations, JSON, marshal/unmarshal

### Iteration 2: enums/versionindexes (1.4%) + coredata/corestr (3.8%)
- versionindexes: 69 uncovered stmts — Index methods, enum marshal, ranges
- corestr: 5540 uncovered stmts — Collection, Hashmap, Hashset, SimpleSlice, LinkedList, SSO, ValidValue

### Iteration 3: issetter (7.2%) + converters (8.1%)
- issetter: 244 uncovered stmts — Value methods, name maps, IsSetter, JSON
- converters: 376 uncovered stmts — StringsToMapConverter, anyItemConverter, stringTo, stringsTo, bytesTo

### Iteration 4: coredata/corejson (9.2%) + enums/stringcompareas (27.9%)
- corejson: 1941 uncovered stmts — Result, ResultsCollection, MapResults, serializer, deserializer, anyTo
- stringcompareas: 93 uncovered stmts — Variant methods, compare functions, IsLineCompareFunc

### Iteration 5: coredata/corepayload (86.4%) + regexnew (87.0%)
- corepayload: 224 uncovered stmts
- regexnew: 29 uncovered stmts

### Iteration 6: coretests (88.9%) + coretests/args (89.0%)
- coretests: 41 uncovered stmts
- args: 189 uncovered stmts

### Iteration 7: coredata/coredynamic (90.3%) + errcore (93.3%)
- coredynamic: 216 uncovered stmts
- errcore: 56 uncovered stmts

### Iteration 8: corecmp (95.1%) + codestack (95.2%)
- corecmp: 9 uncovered stmts
- codestack: 24 uncovered stmts

### Iteration 9: coretests/results (95.9%) + coreimpl/enumimpl (95.9%)
- results: 6 uncovered stmts
- enumimpl: 60 uncovered stmts

### Iteration 10: coretests/coretestcases (95.9%) + coreinstruction (95.9%)
- coretestcases: 11 uncovered stmts
- coreinstruction: 16 uncovered stmts

### Iteration 11: corevalidator (96.1%) + codegen/coreproperty (96.2%)
- corevalidator: 28 uncovered stmts
- coreproperty: 2 uncovered stmts

### Iteration 12: coredata/coreonce (96.8%) + coreversion (97.1%)
- coreonce: 22 uncovered stmts
- coreversion: 11 uncovered stmts

### Iteration 13: coreutils/stringutil (97.3%) + iserror (97.4%)
- stringutil: 12 uncovered stmts
- iserror: 1 uncovered stmt

### Iteration 14: ostype (97.6%) + namevalue (97.9%)
- ostype: 4 uncovered stmts
- namevalue: 4 uncovered stmts

### Iteration 15: simplewrap (98.1%) + isany (98.7%)
- simplewrap: 2 uncovered stmts
- isany: 2 uncovered stmts

### Iteration 16: coretaskinfo (98.8%) + corecomparator (99.1%)
- coretaskinfo: 3 uncovered stmts
- corecomparator: 1 uncovered stmt

### Iteration 17: coredata/stringslice (99.2%) + coredata/coregeneric (99.5%)
- stringslice: 4 uncovered stmts
- coregeneric: 5 uncovered stmts

### Iteration 18: coredata/corerange (99.7%) + coreappend (100% ✓)
- corerange: 2 uncovered stmts

---

## Already at 100%
anycmp, internal/csvinternal, internal/trydo, osconsts, pagingutil,
chmodhelper/chmodclasstype, chmodhelper/chmodins, defaultcapacity,
corefuncs, corecsv, coresort/intsort, coresort/strsort,
coreunique/intunique, internal/fsinternal, mutexbykey,
coredata/coreapi, coreindexes, internal/msgcreator, typesconv,
coreimpl/enumimpl/enumtype, coreappend, converters/coreconverted,
coredata, conditional, corefuncs
