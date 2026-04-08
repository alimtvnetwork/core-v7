# AAA Format Compliance Audit Report

**Date**: 2026-04-06
**Standard**: Arrange → Act (args.Map) → Assert (expected.ShouldBeEqual)

---

## Summary

- **Total packages with issues**: 53
- **Total non-compliant assertion calls**: 33150
- **Total files with issues**: 393
- **Files missing args.Map entirely**: 326

## Package Priority (by violation count)

| # | Package | Violations | Files | No args.Map | Tests |
|---|---------|-----------|-------|------------|-------|
| 1 | `corestrtests` | 15353 | 101 | 86 | 8755 |
| 2 | `corejsontests` | 3138 | 31 | 30 | 2444 |
| 3 | `coredynamictests` | 2644 | 14 | 11 | 1414 |
| 4 | `corepayloadtests` | 1771 | 8 | 6 | 657 |
| 5 | `enumimpltests` | 1446 | 7 | 6 | 600 |
| 6 | `chmodhelpertests` | 1361 | 24 | 18 | 812 |
| 7 | `argstests` | 740 | 6 | 1 | 226 |
| 8 | `issettertests` | 618 | 6 | 6 | 294 |
| 9 | `coregenerictests` | 613 | 6 | 4 | 384 |
| 10 | `reflectmodeltests` | 531 | 8 | 5 | 311 |
| 11 | `converterstests` | 416 | 4 | 4 | 203 |
| 12 | `corevalidatortests` | 332 | 11 | 11 | 214 |
| 13 | `reqtypetests` | 324 | 6 | 6 | 167 |
| 14 | `corecmptests` | 311 | 4 | 4 | 72 |
| 15 | `coreversiontests` | 295 | 10 | 10 | 190 |
| 16 | `codestacktests` | 287 | 5 | 2 | 275 |
| 17 | `coreoncetests` | 282 | 7 | 5 | 186 |
| 18 | `coretaskinfotests` | 272 | 5 | 5 | 129 |
| 19 | `namevaluetests` | 259 | 7 | 5 | 198 |
| 20 | `errcoretests` | 252 | 10 | 6 | 285 |
| 21 | `keymktests` | 238 | 6 | 5 | 155 |
| 22 | `stringslicetests` | 211 | 9 | 7 | 240 |
| 23 | `isanytests` | 169 | 15 | 15 | 93 |
| 24 | `conditionaltests` | 162 | 4 | 3 | 95 |
| 25 | `coreinstructiontests` | 155 | 4 | 4 | 90 |
| 26 | `stringcompareastests` | 131 | 3 | 2 | 58 |
| 27 | `simplewraptests` | 101 | 14 | 14 | 55 |
| 28 | `coreapitests` | 90 | 2 | 2 | 62 |
| 29 | `corecomparatortests` | 89 | 4 | 3 | 47 |
| 30 | `coremathtests` | 84 | 4 | 4 | 15 |
| 31 | `coreutilstests` | 77 | 1 | 0 | 62 |
| 32 | `enumtypetests` | 75 | 1 | 1 | 22 |
| 33 | `bytetypetests` | 73 | 3 | 1 | 38 |
| 34 | `coreappendtests` | 49 | 2 | 2 | 21 |
| 35 | `coreindexestests` | 38 | 1 | 1 | 9 |
| 36 | `corerangetests` | 38 | 5 | 5 | 72 |
| 37 | `coredatatests` | 30 | 2 | 2 | 25 |
| 38 | `resultstests` | 25 | 2 | 1 | 33 |
| 39 | `regexnewtests` | 19 | 4 | 1 | 33 |
| 40 | `stringutiltests` | 13 | 2 | 2 | 9 |
| 41 | `casenilsafetests` | 11 | 1 | 0 | 7 |
| 42 | `ostypetests` | 11 | 2 | 2 | 12 |
| 43 | `iserrortests` | 5 | 2 | 2 | 5 |
| 44 | `coretestcasestests` | 4 | 1 | 0 | 32 |
| 45 | `versionindexestests` | 3 | 2 | 2 | 3 |
| 46 | `chmodinstests` | 2 | 1 | 0 | 12 |
| 47 | `coreteststests` | 1 | 2 | 1 | 113 |
| 48 | `trydotests` | 1 | 2 | 1 | 7 |
| 49 | `anycmptests` | 0 | 1 | 1 | 1 |
| 50 | `corecsvtests` | 0 | 5 | 5 | 13 |
| 51 | `coreflecttests` | 0 | 2 | 2 | 13 |
| 52 | `corerangestests` | 0 | 3 | 3 | 7 |
| 53 | `coretesttests` | 0 | 1 | 1 | 1 |

---

## Detailed File-Level Breakdown

### `corestrtests` (15353 violations, 101 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage13_test.go` | 355 | ❌ | t.Fatalf:7, t.Fatal(:211, if.*!=.*t.Error (manual assert):137 |
| `Coverage14_test.go` | 265 | ❌ | t.Fatalf:3, t.Fatal(:157, if.*!=.*t.Error (manual assert):105 |
| `Coverage21_Collection_Full_test.go` | 162 | ❌ | t.Fatalf:2, t.Fatal(:94, if.*!=.*t.Error (manual assert):66 |
| `Coverage22_Hashmap_Hashset_Full_test.go` | 88 | ❌ | t.Fatalf:2, t.Fatal(:57, if.*!=.*t.Error (manual assert):29 |
| `Coverage23_SimpleSlice_Types_test.go` | 95 | ❌ | t.Fatal(:56, if.*!=.*t.Error (manual assert):39 |
| `Coverage24_LinkedList_Full_test.go` | 284 | ❌ | t.Errorf:79, t.Error(:87, if.*!=.*t.Error (manual assert):118 |
| `Coverage25_LinkedCollections_Full_test.go` | 212 | ❌ | t.Errorf:49, t.Error(:85, if.*!=.*t.Error (manual assert):78 |
| `Coverage26_CharMaps_Full_test.go` | 232 | ❌ | t.Errorf:40, t.Error(:126, if.*!=.*t.Error (manual assert):66 |
| `Coverage27_FinalGaps_test.go` | 50 | ❌ | t.Errorf:7, t.Error(:30, if.*!=.*t.Error (manual assert):13 |
| `Coverage27_Iteration32_test.go` | 339 | ❌ | t.Errorf:2, t.Error(:209, if.*!=.*t.Error (manual assert):128 |
| `Coverage27_RemainingTypes_Full_test.go` | 319 | ❌ | t.Errorf:23, t.Error(:193, if.*!=.*t.Error (manual assert):103 |
| `Coverage28_Iteration33_test.go` | 327 | ❌ | t.Errorf:5, t.Error(:199, if.*!=.*t.Error (manual assert):123 |
| `Coverage29_Iteration34_test.go` | 330 | ❌ | t.Errorf:3, t.Error(:192, if.*!=.*t.Error (manual assert):135 |
| `Coverage30_Collection_test.go` | 104 | ❌ | t.Fatalf:1, t.Fatal(:70, if.*!=.*t.Error (manual assert):33 |
| `Coverage31_Hashset_Hashmap_test.go` | 17 | ❌ | t.Fatal(:15, if.*!=.*t.Error (manual assert):2 |
| `Coverage32_SimpleSlice_SSO_test.go` | 26 | ❌ | t.Fatal(:17, if.*!=.*t.Error (manual assert):9 |
| `Coverage33_Full_test.go` | 79 | ✅ | t.Fatalf:1, t.Fatal(:46, if.*!=.*t.Error (manual assert):32 |
| `Coverage33_LinkedList_LinkedColl_test.go` | 21 | ❌ | t.Fatal(:15, if.*!=.*t.Error (manual assert):6 |
| `Coverage34_Remaining_test.go` | 6 | ✅ | t.Fatal(:4, if.*!=.*t.Error (manual assert):2 |
| `Coverage35_Collection_Full_Coverage_test.go` | 410 | ❌ | t.Errorf:5, t.Error(:237, if.*!=.*t.Error (manual assert):168 |
| `Coverage35_Seg1_Utilities_test.go` | 1 | ✅ | t.Fatal(:1 |
| `Coverage36_Hashmap_Hashset_Full_Coverage_test.go` | 274 | ❌ | t.Error(:167, if.*!=.*t.Error (manual assert):107 |
| `Coverage36_Seg2_CollectionMid_test.go` | 3 | ✅ | t.Fatal(:3 |
| `Coverage37_Collection_Deep_Coverage_test.go` | 398 | ❌ | t.Fatal(:242, if.*!=.*t.Error (manual assert):156 |
| `Coverage37_Iteration33_test.go` | 376 | ❌ | t.Errorf:131, t.Error(:88, t.Fatal(:9, if.*!=.*t.Error (manual assert):148 |
| `Coverage38_Hashmap_Hashset_Coverage_test.go` | 407 | ❌ | t.Fatal(:251, if.*!=.*t.Error (manual assert):156 |
| `Coverage38_Iteration34_test.go` | 306 | ❌ | t.Errorf:87, t.Error(:86, t.Fatal(:6, if.*!=.*t.Error (manual assert):127 |
| `Coverage39_Iteration35_test.go` | 294 | ❌ | t.Errorf:89, t.Error(:80, t.Fatal(:5, if.*!=.*t.Error (manual assert):120 |
| `Coverage39_SimpleSlice_SSO_Coverage_test.go` | 413 | ❌ | t.Fatal(:263, if.*!=.*t.Error (manual assert):150 |
| `Coverage40_Iteration36_test.go` | 455 | ❌ | t.Errorf:80, t.Error(:195, if.*!=.*t.Error (manual assert):180 |
| `Coverage40_Types_Remaining_Coverage_test.go` | 323 | ❌ | t.Fatal(:221, if.*!=.*t.Error (manual assert):102 |
| `Coverage41_Iteration8_test.go` | 190 | ❌ | t.Fatalf:5, t.Fatal(:116, if.*!=.*t.Error (manual assert):69 |
| `Coverage42_Collection_I8_test.go` | 216 | ❌ | t.Fatalf:4, t.Fatal(:127, if.*!=.*t.Error (manual assert):85 |
| `Coverage42_Iteration38_test.go` | 0 | ❌ | — |
| `Coverage43_Hashset_I8_test.go` | 129 | ❌ | t.Fatalf:1, t.Fatal(:82, if.*!=.*t.Error (manual assert):46 |
| `Coverage43_Iteration39_test.go` | 0 | ❌ | — |
| `Coverage44_Hashmap_I8_test.go` | 96 | ❌ | t.Fatalf:1, t.Fatal(:67, if.*!=.*t.Error (manual assert):28 |
| `Coverage44_Iteration40_test.go` | 0 | ❌ | — |
| `Coverage45_Iteration41_test.go` | 0 | ❌ | — |
| `Coverage45_SimpleSlice_I8_test.go` | 110 | ❌ | t.Fatal(:58, if.*!=.*t.Error (manual assert):52 |
| `Coverage46_LinkedList_I8_test.go` | 285 | ❌ | t.Fatalf:29, t.Fatal(:140, if.*!=.*t.Error (manual assert):116 |
| `Coverage47_LinkedCollections_I8_test.go` | 244 | ❌ | t.Fatalf:22, t.Fatal(:121, if.*!=.*t.Error (manual assert):101 |
| `Coverage48_CharMaps_I8_test.go` | 231 | ❌ | t.Fatalf:14, t.Fatal(:141, if.*!=.*t.Error (manual assert):76 |
| `Coverage48_SimpleSlice_S03_test.go` | 36 | ✅ | t.Fatal(:29, if.*!=.*t.Error (manual assert):7 |
| `Coverage49_Iteration45_test.go` | 0 | ❌ | — |
| `Coverage49_RemainingTypes_I8_test.go` | 267 | ❌ | t.Fatalf:3, t.Fatal(:164, if.*!=.*t.Error (manual assert):100 |
| `Coverage56_Iteration52_test.go` | 1 | ✅ | t.Fatal(:1 |
| `Coverage58_CollSeg1_test.go` | 1 | ✅ | t.Fatal(:1 |
| `Coverage64_LinkedListSeg1_test.go` | 1 | ✅ | t.Fatal(:1 |
| `Coverage72_S06_CharCollectionMap_test.go` | 29 | ✅ | t.Errorf:8, t.Error(:15, if.*!=.*t.Error (manual assert):6 |
| `Coverage73_S07_CharHashsetMap_test.go` | 48 | ✅ | t.Errorf:8, t.Error(:31, if.*!=.*t.Error (manual assert):9 |
| `Coverage74_S14_KeyValue_CollOfColl_test.go` | 299 | ❌ | t.Errorf:67, t.Error(:116, t.Fatalf:4, if.*!=.*t.Error (manual assert):112 |
| `Coverage75_S15_SimpleSlice_Extended_test.go` | 181 | ❌ | t.Errorf:45, t.Error(:54, t.Fatalf:2, if.*!=.*t.Error (manual assert):80 |
| `Coverage76_S16_SimpleSlice_Search_Filter_test.go` | 216 | ❌ | t.Errorf:26, t.Error(:113, if.*!=.*t.Error (manual assert):77 |
| `Coverage77_S17_SSO_Core_test.go` | 166 | ❌ | t.Errorf:22, t.Error(:85, if.*!=.*t.Error (manual assert):59 |
| `Coverage78_S18_SSO_Extended_test.go` | 79 | ❌ | t.Errorf:16, t.Error(:28, t.Fatalf:2, if.*!=.*t.Error (manual assert):33 |
| `Coverage79_S19_ValidValue_Types_test.go` | 240 | ❌ | t.Errorf:32, t.Error(:114, if.*!=.*t.Error (manual assert):94 |
| `Coverage80_S20_NonChainedNodes_HashmapDiff_test.go` | 116 | ❌ | t.Errorf:19, t.Error(:53, t.Fatalf:1, t.Fatal(:3, if.*!=.*t.Error (manual assert):40 |
| `Coverage81_S21_CloneSlice_Empty_Reflect_test.go` | 50 | ❌ | t.Errorf:4, t.Error(:27, t.Fatalf:3, t.Fatal(:3, if.*!=.*t.Error (manual assert):13 |
| `Coverage82_S22_FromSplit_Creators_test.go` | 197 | ❌ | t.Errorf:61, t.Error(:54, t.Fatal(:2, if.*!=.*t.Error (manual assert):80 |
| `Coverage83_S23_HashsetsCollection_test.go` | 94 | ❌ | t.Errorf:61, t.Fatalf:1, if.*!=.*t.Error (manual assert):32 |
| `Coverage84_S24_DataModels_Helpers_test.go` | 39 | ❌ | t.Errorf:23, if.*!=.*t.Error (manual assert):16 |
| `Coverage85_S25_CollOfColl_KV_Creators_test.go` | 70 | ❌ | t.Errorf:42, if.*!=.*t.Error (manual assert):28 |
| `Coverage_CharCollectionMap_test.go` | 138 | ❌ | t.Fatal(:98, if.*!=.*t.Error (manual assert):40 |
| `Coverage_CharHashsetMap_test.go` | 140 | ❌ | t.Fatal(:103, if.*!=.*t.Error (manual assert):37 |
| `Coverage_Collection_Part1_test.go` | 128 | ❌ | t.Fatalf:2, t.Fatal(:77, if.*!=.*t.Error (manual assert):49 |
| `Coverage_Collection_Part2_test.go` | 232 | ❌ | t.Fatalf:5, t.Fatal(:131, if.*!=.*t.Error (manual assert):96 |
| `Coverage_Collection_S08_test.go` | 193 | ❌ | t.Fatalf:26, t.Fatal(:88, if.*!=.*t.Error (manual assert):79 |
| `Coverage_Collection_S08b_test.go` | 241 | ❌ | t.Fatalf:41, t.Fatal(:107, if.*!=.*t.Error (manual assert):93 |
| `Coverage_Creators_Utils_test.go` | 138 | ❌ | t.Fatalf:1, t.Fatal(:70, if.*!=.*t.Error (manual assert):67 |
| `Coverage_Hashmap_Part1_test.go` | 74 | ❌ | t.Fatalf:2, t.Fatal(:56, if.*!=.*t.Error (manual assert):16 |
| `Coverage_Hashmap_Part2_test.go` | 107 | ❌ | t.Fatal(:62, if.*!=.*t.Error (manual assert):45 |
| `Coverage_Hashmap_S09_test.go` | 256 | ❌ | t.Fatalf:12, t.Fatal(:152, if.*!=.*t.Error (manual assert):92 |
| `Coverage_Hashset_Part1_test.go` | 116 | ❌ | t.Fatalf:1, t.Fatal(:83, if.*!=.*t.Error (manual assert):32 |
| `Coverage_Hashset_Part2_test.go` | 130 | ❌ | t.Fatalf:7, t.Fatal(:64, if.*!=.*t.Error (manual assert):59 |
| `Coverage_Hashset_S10a_test.go` | 144 | ❌ | t.Fatalf:2, t.Fatal(:96, if.*!=.*t.Error (manual assert):46 |
| `Coverage_Hashset_S10b_test.go` | 151 | ❌ | t.Fatalf:15, t.Fatal(:74, if.*!=.*t.Error (manual assert):62 |
| `Coverage_HashsetsCollection_test.go` | 83 | ❌ | t.Fatal(:50, if.*!=.*t.Error (manual assert):33 |
| `Coverage_KeyValueCollection_test.go` | 164 | ❌ | t.Fatal(:97, if.*!=.*t.Error (manual assert):67 |
| `Coverage_LeftRight_CollOfColl_test.go` | 112 | ❌ | t.Fatalf:3, t.Fatal(:69, if.*!=.*t.Error (manual assert):40 |
| `Coverage_LinkedCollections_Part1_test.go` | 112 | ❌ | t.Fatalf:2, t.Fatal(:60, if.*!=.*t.Error (manual assert):50 |
| `Coverage_LinkedCollections_Part2_test.go` | 97 | ❌ | t.Fatalf:4, t.Fatal(:53, if.*!=.*t.Error (manual assert):40 |
| `Coverage_LinkedCollections_S13_test.go` | 191 | ❌ | t.Fatalf:5, t.Fatal(:111, if.*!=.*t.Error (manual assert):75 |
| `Coverage_LinkedList_Part1_test.go` | 164 | ❌ | t.Fatal(:89, if.*!=.*t.Error (manual assert):75 |
| `Coverage_LinkedList_Part2_test.go` | 108 | ❌ | t.Fatalf:2, t.Fatal(:71, if.*!=.*t.Error (manual assert):35 |
| `Coverage_LinkedList_S12_test.go` | 184 | ❌ | t.Fatalf:4, t.Fatal(:103, if.*!=.*t.Error (manual assert):77 |
| `Coverage_SimpleSlice_Part1_test.go` | 157 | ❌ | t.Fatalf:4, t.Fatal(:91, if.*!=.*t.Error (manual assert):62 |
| `Coverage_SimpleSlice_Part2_test.go` | 91 | ❌ | t.Fatalf:4, t.Fatal(:53, if.*!=.*t.Error (manual assert):34 |
| `Coverage_SimpleSlice_S11a_test.go` | 154 | ❌ | t.Fatalf:5, t.Fatal(:81, if.*!=.*t.Error (manual assert):68 |
| `Coverage_SimpleSlice_S11b_test.go` | 138 | ❌ | t.Fatalf:3, t.Fatal(:93, if.*!=.*t.Error (manual assert):42 |
| `Coverage_SimpleStringOnce_test.go` | 190 | ❌ | t.Fatalf:4, t.Fatal(:113, if.*!=.*t.Error (manual assert):73 |
| `Coverage_ValidValue_ValidValues_test.go` | 226 | ❌ | t.Fatalf:3, t.Fatal(:135, if.*!=.*t.Error (manual assert):88 |
| `Extended_test.go` | 178 | ✅ | t.Errorf:52, t.Error(:62, if.*!=.*t.Error (manual assert):64 |
| `Hashmap_test.go` | 96 | ❌ | t.Errorf:24, t.Error(:42, if.*!=.*t.Error (manual assert):30 |
| `Hashset_test.go` | 103 | ❌ | t.Errorf:31, t.Error(:38, if.*!=.*t.Error (manual assert):34 |
| `NilReceiver_test.go` | 0 | ❌ | — |
| `Src_Coverage02_Collection_test.go` | 1 | ✅ | t.Fatal(:1 |
| `Src_Coverage03_SimpleSlice_test.go` | 6 | ✅ | t.Fatal(:4, if.*!=.*t.Error (manual assert):2 |
| `Src_Coverage04_Types_test.go` | 3 | ✅ | t.Fatal(:3 |
| `Src_Coverage06_Hashmap_test.go` | 5 | ✅ | t.Fatal(:4, if.*!=.*t.Error (manual assert):1 |
| `trydo_safe_test_helper_test.go` | 0 | ❌ | — |

### `corejsontests` (3138 violations, 31 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage10_Result_Migrated_test.go` | 162 | ❌ | t.Fatal(:121, if.*!=.*t.Error (manual assert):41 |
| `Coverage11_Serializer_Migrated_test.go` | 35 | ❌ | t.Fatal(:27, if.*!=.*t.Error (manual assert):8 |
| `Coverage12_Deserializer_Migrated_test.go` | 116 | ❌ | t.Fatal(:63, if.*!=.*t.Error (manual assert):53 |
| `Coverage13_Creators_Migrated_test.go` | 22 | ❌ | t.Fatal(:14, if.*!=.*t.Error (manual assert):8 |
| `Coverage14_AnyTo_CastAny_Migrated_test.go` | 50 | ❌ | t.Fatal(:33, if.*!=.*t.Error (manual assert):17 |
| `Coverage15_Collections_Migrated_test.go` | 171 | ❌ | t.Fatal(:95, if.*!=.*t.Error (manual assert):76 |
| `Coverage16_Gaps_ResultDeser_test.go` | 45 | ❌ | t.Fatal(:31, if.*!=.*t.Error (manual assert):14 |
| `Coverage17_Gaps_Collections_test.go` | 54 | ❌ | t.Fatal(:30, if.*!=.*t.Error (manual assert):24 |
| `Coverage17_Gaps_test.go` | 0 | ❌ | — |
| `Coverage18_Result_Full_test.go` | 105 | ❌ | t.Fatal(:80, if.*!=.*t.Error (manual assert):25 |
| `Coverage19_Collections_Full_test.go` | 163 | ❌ | t.Fatalf:8, t.Fatal(:85, if.*!=.*t.Error (manual assert):70 |
| `Coverage20_MapResults_Funcs_test.go` | 102 | ❌ | t.Fatalf:1, t.Fatal(:59, if.*!=.*t.Error (manual assert):42 |
| `Coverage20_ResultAnyTo_test.go` | 187 | ❌ | t.Fatal(:136, if.*!=.*t.Error (manual assert):51 |
| `Coverage21_Result_Methods_test.go` | 78 | ❌ | t.Fatal(:58, if.*!=.*t.Error (manual assert):20 |
| `Coverage22_BytesCollection_test.go` | 133 | ❌ | t.Fatalf:2, t.Fatal(:71, if.*!=.*t.Error (manual assert):60 |
| `Coverage23_MapResults_test.go` | 78 | ❌ | t.Fatalf:2, t.Fatal(:41, if.*!=.*t.Error (manual assert):35 |
| `Coverage24_Creators_Deser_test.go` | 39 | ❌ | t.Fatal(:28, if.*!=.*t.Error (manual assert):11 |
| `Coverage25_AnyTo_Deser_test.go` | 91 | ❌ | t.Fatal(:53, if.*!=.*t.Error (manual assert):38 |
| `Coverage26_CollCreators_test.go` | 37 | ❌ | t.Fatal(:23, if.*!=.*t.Error (manual assert):14 |
| `Coverage27_Collections_Remaining_test.go` | 32 | ❌ | t.Fatal(:20, if.*!=.*t.Error (manual assert):12 |
| `Coverage28_Result_Full_Coverage_test.go` | 161 | ❌ | t.Fatal(:116, if.*!=.*t.Error (manual assert):45 |
| `Coverage29_Collections_Full_Coverage_test.go` | 112 | ❌ | t.Fatal(:69, if.*!=.*t.Error (manual assert):43 |
| `Coverage30_BytesColl_MapResults_Coverage_test.go` | 209 | ❌ | t.Fatal(:121, if.*!=.*t.Error (manual assert):88 |
| `Coverage31_PtrColl_Creators_Deser_Coverage_test.go` | 149 | ❌ | t.Fatal(:86, if.*!=.*t.Error (manual assert):63 |
| `Coverage33_Iteration20_test.go` | 275 | ❌ | t.Fatalf:3, t.Fatal(:189, if.*!=.*t.Error (manual assert):83 |
| `Coverage35_Seg2_Collections_Deser_test.go` | 264 | ❌ | t.Fatal(:149, if.*!=.*t.Error (manual assert):115 |
| `Coverage36_Seg3_AnyTo_Funcs_Utils_test.go` | 88 | ❌ | t.Fatal(:53, if.*!=.*t.Error (manual assert):35 |
| `Coverage37_Seg4_Result_RC_PtrRC_MR_test.go` | 120 | ❌ | t.Fatal(:86, if.*!=.*t.Error (manual assert):34 |
| `Coverage38_Seg5_Final_Interfaces_Edge_test.go` | 59 | ❌ | t.Fatal(:39, if.*!=.*t.Error (manual assert):20 |
| `Coverage51_Gaps_test.go` | 1 | ✅ | if.*!=.*t.Error (manual assert):1 |
| `Deserializer_Apply_test.go` | 0 | ❌ | — |

### `coredynamictests` (2644 violations, 14 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage17_Iteration7_test.go` | 1 | ✅ | if.*!=.*t.Error (manual assert):1 |
| `Coverage18_Full_test.go` | 211 | ✅ | t.Fatalf:5, t.Fatal(:135, if.*!=.*t.Error (manual assert):71 |
| `Coverage31_Iteration27_test.go` | 459 | ❌ | t.Errorf:282, if.*!=.*t.Error (manual assert):177 |
| `Coverage32_Iteration28_test.go` | 523 | ❌ | t.Errorf:321, if.*!=.*t.Error (manual assert):202 |
| `Coverage33_Iteration29_test.go` | 374 | ❌ | t.Errorf:73, t.Error(:154, if.*!=.*t.Error (manual assert):147 |
| `Coverage34_Iteration30_test.go` | 391 | ❌ | t.Errorf:26, t.Error(:214, if.*!=.*t.Error (manual assert):151 |
| `Coverage35_Iteration31_test.go` | 290 | ❌ | t.Errorf:15, t.Error(:171, if.*!=.*t.Error (manual assert):104 |
| `Coverage36_Iteration32_test.go` | 304 | ❌ | t.Errorf:12, t.Error(:182, if.*!=.*t.Error (manual assert):110 |
| `Coverage74_FinalGaps_test.go` | 29 | ❌ | t.Errorf:5, t.Error(:17, if.*!=.*t.Error (manual assert):7 |
| `Coverage7_test.go` | 2 | ✅ | t.Errorf:1, if.*!=.*t.Error (manual assert):1 |
| `Extended_test.go` | 60 | ❌ | t.Errorf:2, t.Error(:39, if.*!=.*t.Error (manual assert):19 |
| `NilReceiver_test.go` | 0 | ❌ | — |
| `ReflectSetFromTo_InvalidCases_test.go` | 0 | ❌ | — |
| `ReflectSetFromTo_ValidCases_test.go` | 0 | ❌ | — |

### `corepayloadtests` (1771 violations, 8 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage15_test.go` | 2 | ✅ | t.Fatalf:1, if.*!=.*t.Error (manual assert):1 |
| `Coverage17_Gaps_test.go` | 0 | ❌ | — |
| `Coverage18_PayloadWrapper_Attrs_test.go` | 200 | ❌ | t.Fatal(:143, if.*!=.*t.Error (manual assert):57 |
| `Coverage19_Seg2_Collections_TypedPW_test.go` | 320 | ❌ | t.Fatal(:220, if.*!=.*t.Error (manual assert):100 |
| `Coverage20_Seg3_PW_Attrs_GenericHelpers_test.go` | 129 | ❌ | t.Fatal(:91, if.*!=.*t.Error (manual assert):38 |
| `Coverage21_Seg4_Final_PC_TPW_TPC_Creators_test.go` | 330 | ❌ | t.Fatal(:206, if.*!=.*t.Error (manual assert):124 |
| `Coverage26_Iteration28_test.go` | 12 | ✅ | t.Fatalf:6, if.*!=.*t.Error (manual assert):6 |
| `Coverage9_test.go` | 778 | ❌ | t.Fatalf:4, t.Fatal(:510, if.*!=.*t.Error (manual assert):264 |

### `enumimpltests` (1446 violations, 7 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage12_test.go` | 431 | ❌ | t.Fatal(:314, if.*!=.*t.Error (manual assert):117 |
| `Coverage13_Iteration9_test.go` | 434 | ❌ | t.Fatalf:4, t.Fatal(:281, if.*!=.*t.Error (manual assert):149 |
| `Coverage16_DynamicMap_test.go` | 310 | ❌ | t.Fatal(:216, if.*!=.*t.Error (manual assert):94 |
| `Coverage16_Gaps_test.go` | 0 | ❌ | — |
| `Coverage18_FinalGaps_test.go` | 16 | ❌ | t.Errorf:1, t.Error(:13, if.*!=.*t.Error (manual assert):2 |
| `Coverage_test.go` | 65 | ❌ | t.Errorf:7, t.Error(:42, if.*!=.*t.Error (manual assert):16 |
| `Extended_test.go` | 190 | ✅ | t.Errorf:46, t.Error(:85, if.*!=.*t.Error (manual assert):59 |

### `chmodhelpertests` (1361 violations, 24 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `ApplyOnPath_test.go` | 0 | ❌ | — |
| `Attribute_test.go` | 22 | ✅ | t.Errorf:5, t.Error(:9, t.Fatalf:1, if.*!=.*t.Error (manual assert):7 |
| `Coverage10_Executors_test.go` | 48 | ❌ | t.Fatalf:1, t.Fatal(:25, if.*!=.*t.Error (manual assert):22 |
| `Coverage11_VarWrapper_test.go` | 33 | ❌ | t.Fatal(:26, if.*!=.*t.Error (manual assert):7 |
| `Coverage12_SimpleFileRW_test.go` | 30 | ❌ | t.Fatal(:21, if.*!=.*t.Error (manual assert):9 |
| `Coverage13_Remaining_test.go` | 91 | ❌ | t.Fatalf:2, t.Fatal(:64, if.*!=.*t.Error (manual assert):25 |
| `Coverage14_Iteration18_test.go` | 64 | ❌ | t.Fatalf:6, t.Fatal(:38, if.*!=.*t.Error (manual assert):20 |
| `Coverage15_Iteration10_test.go` | 202 | ❌ | t.Fatalf:5, t.Fatal(:123, if.*!=.*t.Error (manual assert):74 |
| `Coverage17_Gaps_test.go` | 2 | ❌ | t.Fatal(:1, if.*!=.*t.Error (manual assert):1 |
| `Coverage19_RwxInstruction_test.go` | 8 | ✅ | t.Fatalf:4, if.*!=.*t.Error (manual assert):4 |
| `Coverage19_RwxWrapper_ApplyMany_test.go` | 12 | ✅ | t.Fatalf:6, if.*!=.*t.Error (manual assert):6 |
| `Coverage7_test.go` | 74 | ❌ | t.Fatal(:47, if.*!=.*t.Error (manual assert):27 |
| `Coverage8_ErrorsAndPaths_test.go` | 54 | ❌ | t.Fatalf:1, t.Fatal(:38, if.*!=.*t.Error (manual assert):15 |
| `Coverage9_RwxWrapper_test.go` | 30 | ❌ | t.Fatalf:1, t.Fatal(:20, if.*!=.*t.Error (manual assert):9 |
| `DirFilesWithContent_test.go` | 0 | ❌ | — |
| `Extended2_test.go` | 483 | ✅ | t.Errorf:111, t.Error(:217, t.Fatalf:1, if.*!=.*t.Error (manual assert):154 |
| `Extended_test.go` | 87 | ✅ | t.Errorf:12, t.Error(:54, t.Fatalf:1, if.*!=.*t.Error (manual assert):20 |
| `LinuxApplyRecursiveOnPath_test.go` | 0 | ❌ | — |
| `RwxCompileValue_test.go` | 0 | ❌ | — |
| `RwxWrapperManyApplyValue_test.go` | 0 | ❌ | — |
| `RwxWrapper_test.go` | 121 | ✅ | t.Errorf:33, t.Error(:36, t.Fatalf:6, t.Fatal(:6, if.*!=.*t.Error (manual assert):40 |
| `SimpleFileWriter_CreateDir_test.go` | 0 | ❌ | — |
| `VerifyPartialRwxLocations_test.go` | 0 | ❌ | — |
| `VerifyRwxChmodUsingRwxInstructions_test.go` | 0 | ❌ | — |

### `argstests` (740 violations, 6 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage10_test.go` | 196 | ✅ | t.Fatalf:1, t.Fatal(:122, if.*!=.*t.Error (manual assert):73 |
| `Coverage8_Iteration18_test.go` | 218 | ✅ | t.Fatalf:1, t.Fatal(:148, if.*!=.*t.Error (manual assert):69 |
| `Coverage9_test.go` | 1 | ✅ | if.*!=.*t.Error (manual assert):1 |
| `Coverage_test.go` | 82 | ✅ | t.Errorf:8, t.Error(:42, if.*!=.*t.Error (manual assert):32 |
| `Extended_test.go` | 243 | ✅ | t.Errorf:23, t.Error(:132, if.*!=.*t.Error (manual assert):88 |
| `FuncWrap_Creation_test.go` | 0 | ❌ | — |

### `issettertests` (618 violations, 6 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage8_Iteration3_test.go` | 247 | ❌ | t.Fatalf:2, t.Fatal(:179, if.*!=.*t.Error (manual assert):66 |
| `Coverage_test.go` | 149 | ❌ | t.Errorf:3, t.Error(:116, if.*!=.*t.Error (manual assert):30 |
| `Coverage_toHashset_test.go` | 5 | ❌ | t.Errorf:2, t.Error(:1, if.*!=.*t.Error (manual assert):2 |
| `Extended2_test.go` | 12 | ❌ | t.Error(:8, t.Fatalf:1, if.*!=.*t.Error (manual assert):3 |
| `Extended_test.go` | 205 | ❌ | t.Errorf:36, t.Error(:89, t.Fatalf:5, if.*!=.*t.Error (manual assert):75 |
| `QuickWin_test.go` | 0 | ❌ | — |

### `coregenerictests` (613 violations, 6 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage_Collection_LinkedList_Numeric_test.go` | 86 | ❌ | t.Errorf:2, t.Error(:45, if.*!=.*t.Error (manual assert):39 |
| `Coverage_test.go` | 483 | ❌ | t.Errorf:2, t.Error(:291, if.*!=.*t.Error (manual assert):190 |
| `NilReceiver_test.go` | 0 | ❌ | — |
| `PairTripleExtended_test.go` | 20 | ✅ | t.Errorf:10, if.*!=.*t.Error (manual assert):10 |
| `PairTriple_test.go` | 14 | ✅ | t.Errorf:7, if.*!=.*t.Error (manual assert):7 |
| `QuickWin_test.go` | 10 | ❌ | t.Fatal(:5, if.*!=.*t.Error (manual assert):5 |

### `reflectmodeltests` (531 violations, 8 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage10_Iteration13_test.go` | 1 | ✅ | t.Fatal(:1 |
| `Coverage6_test.go` | 111 | ❌ | t.Fatalf:11, t.Fatal(:64, if.*!=.*t.Error (manual assert):36 |
| `Coverage7_Iteration8_test.go` | 145 | ✅ | t.Fatalf:5, t.Fatal(:92, if.*!=.*t.Error (manual assert):48 |
| `Coverage8_Iteration8_test.go` | 118 | ❌ | t.Fatalf:9, t.Fatal(:74, if.*!=.*t.Error (manual assert):35 |
| `Coverage9_Full_test.go` | 71 | ✅ | t.Fatalf:21, t.Fatal(:22, if.*!=.*t.Error (manual assert):28 |
| `FieldProcessor_test.go` | 18 | ❌ | t.Errorf:3, t.Error(:6, t.Fatal(:6, if.*!=.*t.Error (manual assert):3 |
| `MethodProcessor_test.go` | 57 | ❌ | t.Errorf:18, t.Error(:17, t.Fatalf:1, t.Fatal(:1, if.*!=.*t.Error (manual assert):20 |
| `ReflectValueKind_test.go` | 10 | ❌ | t.Errorf:3, t.Error(:4, if.*!=.*t.Error (manual assert):3 |

### `converterstests` (416 violations, 4 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage5_Iteration3_test.go` | 237 | ❌ | t.Fatalf:12, t.Fatal(:128, if.*!=.*t.Error (manual assert):97 |
| `Coverage_AnyItem_StringTo_StringsTo_test.go` | 15 | ❌ | t.Errorf:4, t.Error(:6, if.*!=.*t.Error (manual assert):5 |
| `Extended_test.go` | 162 | ❌ | t.Errorf:37, t.Error(:54, t.Fatalf:1, if.*!=.*t.Error (manual assert):70 |
| `QuickWin_test.go` | 2 | ❌ | t.Fatal(:1, if.*!=.*t.Error (manual assert):1 |

### `corevalidatortests` (332 violations, 11 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `BaseLinesValidators_test.go` | 27 | ❌ | t.Errorf:4, t.Error(:18, if.*!=.*t.Error (manual assert):5 |
| `Coverage17_Gaps_test.go` | 0 | ❌ | — |
| `Coverage_test.go` | 124 | ❌ | t.Error(:84, if.*!=.*t.Error (manual assert):40 |
| `LinesValidators_test.go` | 33 | ❌ | t.Errorf:9, t.Error(:13, t.Fatal(:2, if.*!=.*t.Error (manual assert):9 |
| `NilReceiver_test.go` | 0 | ❌ | — |
| `SliceValidatorDiff_test.go` | 53 | ❌ | t.Errorf:18, t.Error(:9, t.Fatalf:2, t.Fatal(:8, if.*!=.*t.Error (manual assert):16 |
| `SliceValidatorExtra_test.go` | 60 | ❌ | t.Errorf:13, t.Error(:30, if.*!=.*t.Error (manual assert):17 |
| `SliceValidator_test.go` | 0 | ❌ | — |
| `SliceValidators_test.go` | 24 | ❌ | t.Errorf:4, t.Error(:11, if.*!=.*t.Error (manual assert):9 |
| `TestValidators_test.go` | 0 | ❌ | — |
| `ValidatorDiffDiagnostics_test.go` | 11 | ❌ | t.Errorf:4, t.Error(:2, t.Fatal(:5 |

### `reqtypetests` (324 violations, 6 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage4_test.go` | 110 | ❌ | t.Errorf:1, t.Fatalf:1, t.Fatal(:93, if.*!=.*t.Error (manual assert):15 |
| `Coverage5_Iteration2_test.go` | 0 | ❌ | — |
| `Coverage_Request_test.go` | 6 | ❌ | t.Error(:6 |
| `Coverage_test.go` | 130 | ❌ | t.Errorf:3, t.Error(:112, if.*!=.*t.Error (manual assert):15 |
| `Extended_test.go` | 74 | ❌ | t.Errorf:5, t.Error(:59, if.*!=.*t.Error (manual assert):10 |
| `QuickWin_test.go` | 4 | ❌ | t.Fatal(:2, if.*!=.*t.Error (manual assert):2 |

### `corecmptests` (311 violations, 4 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage17_deadcode_test.go` | 0 | ❌ | — |
| `Coverage18_deadcode_gaps_test.go` | 66 | ❌ | t.Fatal(:33, if.*!=.*t.Error (manual assert):33 |
| `Coverage_test.go` | 152 | ❌ | t.Error(:86, if.*!=.*t.Error (manual assert):66 |
| `Extended3_test.go` | 93 | ❌ | t.Error(:53, if.*!=.*t.Error (manual assert):40 |

### `coreversiontests` (295 violations, 10 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Comparison_test.go` | 0 | ❌ | — |
| `Coverage2_test.go` | 90 | ❌ | t.Errorf:3, t.Error(:71, if.*!=.*t.Error (manual assert):16 |
| `Coverage5_Iteration2_test.go` | 3 | ❌ | t.Fatalf:2, if.*!=.*t.Error (manual assert):1 |
| `Coverage6_DeadCode_test.go` | 0 | ❌ | — |
| `Coverage_test.go` | 101 | ❌ | t.Errorf:10, t.Error(:71, if.*!=.*t.Error (manual assert):20 |
| `Creation_test.go` | 0 | ❌ | — |
| `Extended_test.go` | 99 | ❌ | t.Errorf:3, t.Error(:73, if.*!=.*t.Error (manual assert):23 |
| `Json_test.go` | 0 | ❌ | — |
| `Methods_test.go` | 0 | ❌ | — |
| `QuickWin_test.go` | 2 | ❌ | t.Fatal(:1, if.*!=.*t.Error (manual assert):1 |

### `codestacktests` (287 violations, 5 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage12_Gaps_test.go` | 0 | ❌ | — |
| `Coverage13_FinalGaps_test.go` | 0 | ❌ | — |
| `Coverage_test.go` | 125 | ✅ | t.Errorf:6, t.Error(:97, if.*!=.*t.Error (manual assert):22 |
| `Extended_test.go` | 138 | ✅ | t.Errorf:36, t.Error(:62, if.*!=.*t.Error (manual assert):40 |
| `Trace_test.go` | 24 | ✅ | t.Errorf:7, t.Error(:9, if.*!=.*t.Error (manual assert):8 |

### `coreoncetests` (282 violations, 7 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `BytesErrorOnce_NilReceiver_test.go` | 0 | ❌ | — |
| `Coverage10_Iteration19_test.go` | 17 | ❌ | t.Fatalf:2, t.Fatal(:11, if.*!=.*t.Error (manual assert):4 |
| `Coverage11_Iteration20_test.go` | 253 | ❌ | t.Fatalf:7, t.Fatal(:179, if.*!=.*t.Error (manual assert):67 |
| `Coverage12_Iteration1_test.go` | 8 | ❌ | t.Fatalf:6, if.*!=.*t.Error (manual assert):2 |
| `Coverage16_Gaps_test.go` | 0 | ❌ | — |
| `Coverage_test.go` | 1 | ✅ | if.*!=.*t.Error (manual assert):1 |
| `Extended_test.go` | 3 | ✅ | t.Errorf:1, t.Error(:1, if.*!=.*t.Error (manual assert):1 |

### `coretaskinfotests` (272 violations, 5 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage3_Iteration2_test.go` | 1 | ❌ | t.Fatalf:1 |
| `Coverage_ExcludingOptions_InfoMap_test.go` | 19 | ❌ | t.Error(:19 |
| `Extended_test.go` | 248 | ❌ | t.Errorf:6, t.Error(:181, if.*!=.*t.Error (manual assert):61 |
| `NilReceiver_test.go` | 0 | ❌ | — |
| `QuickWin_test.go` | 4 | ❌ | t.Fatal(:3, if.*!=.*t.Error (manual assert):1 |

### `namevaluetests` (259 violations, 7 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage10_FinalGaps_test.go` | 0 | ❌ | — |
| `Coverage2_test.go` | 104 | ❌ | t.Errorf:1, t.Error(:62, if.*!=.*t.Error (manual assert):41 |
| `Coverage6_Full_test.go` | 94 | ✅ | t.Fatalf:4, t.Fatal(:50, if.*!=.*t.Error (manual assert):40 |
| `Coverage9_NilAndDeadCode_test.go` | 0 | ❌ | — |
| `Coverage_Collection_Instance_test.go` | 45 | ❌ | t.Error(:29, if.*!=.*t.Error (manual assert):16 |
| `Extended_test.go` | 11 | ✅ | t.Errorf:3, t.Error(:3, if.*!=.*t.Error (manual assert):5 |
| `QuickWin_test.go` | 5 | ❌ | t.Fatal(:3, if.*!=.*t.Error (manual assert):2 |

### `errcoretests` (252 violations, 10 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage14_UtilityFunctions_test.go` | 161 | ❌ | t.Fatal(:118, if.*!=.*t.Error (manual assert):43 |
| `Coverage15_Gaps_test.go` | 0 | ❌ | — |
| `Coverage15_helpers_test.go` | 0 | ❌ | — |
| `Coverage16_FinalGaps_test.go` | 10 | ❌ | t.Errorf:2, t.Error(:5, if.*!=.*t.Error (manual assert):3 |
| `ErrType_test.go` | 1 | ✅ | t.Error(:1 |
| `Extended_test.go` | 74 | ❌ | t.Errorf:2, t.Error(:51, if.*!=.*t.Error (manual assert):21 |
| `NilReceiver_test.go` | 0 | ❌ | — |
| `Src_ErrorHandling_test.go` | 1 | ✅ | if.*!=.*t.Error (manual assert):1 |
| `Src_LineDiffExpecting_test.go` | 3 | ✅ | t.Fatal(:3 |
| `Src_RawErrCollection_test.go` | 2 | ✅ | t.Fatal(:2 |

### `keymktests` (238 violations, 6 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage6_SkipEmpty_test.go` | 0 | ❌ | — |
| `Coverage_KeyCompiler_KeyJson_test.go` | 27 | ❌ | t.Errorf:2, t.Error(:19, if.*!=.*t.Error (manual assert):6 |
| `Extended2_test.go` | 122 | ✅ | t.Errorf:19, t.Error(:76, t.Fatalf:2, if.*!=.*t.Error (manual assert):25 |
| `Extended_test.go` | 87 | ❌ | t.Errorf:26, t.Error(:24, if.*!=.*t.Error (manual assert):37 |
| `QuickWin_test.go` | 2 | ❌ | t.Fatal(:2 |
| `helpers_test.go` | 0 | ❌ | — |

### `stringslicetests` (211 violations, 9 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage12_Iteration1_test.go` | 8 | ❌ | t.Fatalf:4, if.*!=.*t.Error (manual assert):4 |
| `Coverage13_Gaps_test.go` | 0 | ❌ | — |
| `Coverage14_FinalGaps_test.go` | 0 | ❌ | — |
| `Coverage3_test.go` | 1 | ✅ | t.Error(:1 |
| `Coverage_MergeSlices_Split_Regex_test.go` | 12 | ❌ | t.Errorf:5, t.Error(:1, if.*!=.*t.Error (manual assert):6 |
| `Extended_test.go` | 79 | ❌ | t.Errorf:30, t.Error(:12, if.*!=.*t.Error (manual assert):37 |
| `QuickWin_test.go` | 2 | ❌ | t.Fatal(:1, if.*!=.*t.Error (manual assert):1 |
| `Src_StateChecks_test.go` | 20 | ✅ | t.Fatalf:10, t.Fatal(:1, if.*!=.*t.Error (manual assert):9 |
| `StringSlice_test.go` | 89 | ❌ | t.Errorf:38, t.Error(:10, if.*!=.*t.Error (manual assert):41 |

### `isanytests` (169 violations, 15 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `AllNull_test.go` | 0 | ❌ | — |
| `AnyNull_test.go` | 0 | ❌ | — |
| `BothDefined_test.go` | 0 | ❌ | — |
| `Conclusive_test.go` | 0 | ❌ | — |
| `Coverage_Conclusive_JsonEqual_test.go` | 7 | ❌ | t.Error(:7 |
| `Coverage_test.go` | 87 | ❌ | t.Errorf:1, t.Error(:85, if.*!=.*t.Error (manual assert):1 |
| `DefinedAllOf_test.go` | 0 | ❌ | — |
| `DefinedAnyOf_test.go` | 0 | ❌ | — |
| `Defined_test.go` | 0 | ❌ | — |
| `Extended2_test.go` | 72 | ❌ | t.Errorf:1, t.Error(:70, if.*!=.*t.Error (manual assert):1 |
| `JsonEqual_test.go` | 0 | ❌ | — |
| `NullBoth_test.go` | 0 | ❌ | — |
| `Null_test.go` | 0 | ❌ | — |
| `QuickWin_test.go` | 3 | ❌ | t.Fatal(:3 |
| `ReflectionTypesVerify_test.go` | 0 | ❌ | — |

### `conditionaltests` (162 violations, 4 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage_AnyFunctions_TypedError_test.go` | 9 | ❌ | t.Errorf:2, t.Error(:3, if.*!=.*t.Error (manual assert):4 |
| `Coverage_test.go` | 80 | ❌ | t.Errorf:35, t.Error(:7, if.*!=.*t.Error (manual assert):38 |
| `Extended2_test.go` | 33 | ❌ | t.Error(:18, if.*!=.*t.Error (manual assert):15 |
| `Extended3_test.go` | 40 | ✅ | t.Errorf:20, if.*!=.*t.Error (manual assert):20 |

### `coreinstructiontests` (155 violations, 4 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage_test.go` | 91 | ❌ | t.Errorf:5, t.Error(:57, if.*!=.*t.Error (manual assert):29 |
| `Extended_test.go` | 36 | ❌ | t.Errorf:1, t.Error(:29, if.*!=.*t.Error (manual assert):6 |
| `NilReceiver_test.go` | 0 | ❌ | — |
| `Src_Coverage01_Iteration15_test.go` | 28 | ❌ | t.Fatal(:24, if.*!=.*t.Error (manual assert):4 |

### `stringcompareastests` (131 violations, 3 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage_IsNotContains_test.go` | 2 | ❌ | t.Error(:2 |
| `Extended_test.go` | 121 | ❌ | t.Errorf:5, t.Error(:94, if.*!=.*t.Error (manual assert):22 |
| `Glob_test.go` | 8 | ✅ | t.Errorf:2, t.Error(:4, if.*!=.*t.Error (manual assert):2 |

### `simplewraptests` (101 violations, 14 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage8_Iteration2_test.go` | 4 | ❌ | t.Fatalf:2, if.*!=.*t.Error (manual assert):2 |
| `Coverage_DoubleQuoteWrap_test.go` | 10 | ❌ | t.Error(:5, if.*!=.*t.Error (manual assert):5 |
| `Coverage_test.go` | 75 | ❌ | t.Errorf:27, t.Error(:14, if.*!=.*t.Error (manual assert):34 |
| `CurlyWrapOptions_test.go` | 0 | ❌ | — |
| `DoubleQuoteWrapElementsWithIndexes_test.go` | 0 | ❌ | — |
| `DoubleQuoteWrapElements_test.go` | 0 | ❌ | — |
| `Extended2_test.go` | 8 | ❌ | t.Errorf:1, t.Error(:6, if.*!=.*t.Error (manual assert):1 |
| `MsgCsvItems_test.go` | 0 | ❌ | — |
| `MsgWrapMsg_test.go` | 0 | ❌ | — |
| `ParenthesisWrap_test.go` | 0 | ❌ | — |
| `QuickWin_test.go` | 4 | ❌ | t.Fatal(:2, if.*!=.*t.Error (manual assert):2 |
| `SquareWrap_test.go` | 0 | ❌ | — |
| `TitleCurlyMeta_test.go` | 0 | ❌ | — |
| `WithBrackets_test.go` | 0 | ❌ | — |

### `coreapitests` (90 violations, 2 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage_test.go` | 90 | ❌ | t.Error(:61, if.*!=.*t.Error (manual assert):29 |
| `NilReceiver_test.go` | 0 | ❌ | — |

### `corecomparatortests` (89 violations, 4 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage4_test.go` | 84 | ❌ | t.Errorf:1, t.Error(:65, if.*!=.*t.Error (manual assert):18 |
| `Coverage_Compare_test.go` | 2 | ❌ | t.Error(:2 |
| `Extended3_test.go` | 2 | ✅ | t.Error(:1, if.*!=.*t.Error (manual assert):1 |
| `QuickWin_test.go` | 1 | ❌ | t.Fatal(:1 |

### `coremathtests` (84 violations, 4 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage4_DeadCode_test.go` | 0 | ❌ | — |
| `Coverage_IntegerOutOfRange_test.go` | 1 | ❌ | t.Error(:1 |
| `Extended2_test.go` | 81 | ❌ | t.Errorf:2, t.Error(:73, if.*!=.*t.Error (manual assert):6 |
| `QuickWin_test.go` | 2 | ❌ | t.Fatal(:2 |

### `coreutilstests` (77 violations, 1 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Extended_test.go` | 77 | ✅ | t.Errorf:35, t.Error(:8, if.*!=.*t.Error (manual assert):34 |

### `enumtypetests` (75 violations, 1 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage_test.go` | 75 | ❌ | t.Errorf:3, t.Error(:55, if.*!=.*t.Error (manual assert):17 |

### `bytetypetests` (73 violations, 3 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage_Variant_test.go` | 1 | ❌ | t.Error(:1 |
| `Extended2_test.go` | 35 | ✅ | t.Errorf:6, t.Error(:23, if.*!=.*t.Error (manual assert):6 |
| `Extended_test.go` | 37 | ✅ | t.Errorf:10, t.Error(:18, if.*!=.*t.Error (manual assert):9 |

### `coreappendtests` (49 violations, 2 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage_PrependAppendFunc_test.go` | 2 | ❌ | t.Errorf:2 |
| `MapAppend_test.go` | 47 | ❌ | t.Errorf:23, t.Error(:1, if.*!=.*t.Error (manual assert):23 |

### `coreindexestests` (38 violations, 1 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Extended_test.go` | 38 | ❌ | t.Errorf:3, t.Error(:20, if.*!=.*t.Error (manual assert):15 |

### `corerangetests` (38 violations, 5 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage7_Gaps_test.go` | 0 | ❌ | — |
| `Coverage7_Iteration2_test.go` | 0 | ❌ | — |
| `Coverage8_FinalGaps_test.go` | 0 | ❌ | — |
| `Coverage_MinMaxByte_Within_test.go` | 38 | ❌ | t.Errorf:1, t.Error(:20, if.*!=.*t.Error (manual assert):17 |
| `QuickWin_test.go` | 0 | ❌ | — |

### `coredatatests` (30 violations, 2 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `BytesError_test.go` | 30 | ❌ | t.Errorf:6, t.Error(:17, if.*!=.*t.Error (manual assert):7 |
| `FuncWrap_Creation_test.go` | 0 | ❌ | — |

### `resultstests` (25 violations, 2 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage2_Iteration9_test.go` | 14 | ✅ | t.Fatal(:11, if.*!=.*t.Error (manual assert):3 |
| `Coverage3_Iteration1_test.go` | 11 | ❌ | t.Fatalf:7, if.*!=.*t.Error (manual assert):4 |

### `regexnewtests` (19 violations, 4 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `CreateMust_test.go` | 3 | ✅ | t.Error(:3 |
| `Enhancements_test.go` | 4 | ✅ | t.Errorf:4 |
| `LazyRegex_Concurrency_test.go` | 8 | ❌ | t.Error(:5, if.*!=.*t.Error (manual assert):3 |
| `LazyRegex_EdgeCases_test.go` | 4 | ✅ | t.Error(:3, if.*!=.*t.Error (manual assert):1 |

### `stringutiltests` (13 violations, 2 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage7_Gaps_test.go` | 0 | ❌ | — |
| `QuickWin_test.go` | 13 | ❌ | t.Fatal(:8, if.*!=.*t.Error (manual assert):5 |

### `casenilsafetests` (11 violations, 1 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `CaseNilSafe_test.go` | 11 | ✅ | t.Errorf:5, t.Error(:1, if.*!=.*t.Error (manual assert):5 |

### `ostypetests` (11 violations, 2 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage_Group_Variation_test.go` | 9 | ❌ | t.Error(:9 |
| `QuickWin_test.go` | 2 | ❌ | t.Fatal(:2 |

### `iserrortests` (5 violations, 2 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage_Equal_test.go` | 2 | ❌ | t.Error(:2 |
| `QuickWin_test.go` | 3 | ❌ | t.Fatal(:3 |

### `coretestcasestests` (4 violations, 1 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage7_test.go` | 4 | ✅ | t.Fatalf:2, if.*!=.*t.Error (manual assert):2 |

### `versionindexestests` (3 violations, 2 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage_Index_test.go` | 1 | ❌ | t.Error(:1 |
| `QuickWin_test.go` | 2 | ❌ | t.Fatal(:2 |

### `chmodinstests` (2 violations, 1 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage_Parse_test.go` | 2 | ✅ | t.Fatal(:2 |

### `coreteststests` (1 violations, 2 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage3_Iteration7_test.go` | 1 | ✅ | t.Fatal(:1 |
| `Coverage_Iteration6_test.go` | 0 | ❌ | — |

### `trydotests` (1 violations, 2 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `NilReceiver_test.go` | 0 | ❌ | — |
| `WrappedErr_test.go` | 1 | ✅ | if.*!=.*t.Error (manual assert):1 |

### `anycmptests` (0 violations, 1 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Cmp_test.go` | 0 | ❌ | — |

### `corecsvtests` (0 violations, 5 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `AnyItemsToCsvString_test.go` | 0 | ❌ | — |
| `AnyToTypesCsvStrings_test.go` | 0 | ❌ | — |
| `RangeNamesWithValuesIndexes_test.go` | 0 | ❌ | — |
| `StringersToStringDefault_test.go` | 0 | ❌ | — |
| `StringsToCsvString_test.go` | 0 | ❌ | — |

### `coreflecttests` (0 violations, 2 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `FuncWrap_Creation_test.go` | 0 | ❌ | — |
| `ReflectcoreVars_test.go` | 0 | ❌ | — |

### `corerangestests` (0 violations, 3 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `range_test.go` | 0 | ❌ | — |
| `start_end_test.go` | 0 | ❌ | — |
| `within_range_test.go` | 0 | ❌ | — |

### `coretesttests` (0 violations, 1 files)

| File | Violations | args.Map | Issues |
|------|-----------|----------|--------|
| `Coverage3_Gaps_test.go` | 0 | ❌ | — |
