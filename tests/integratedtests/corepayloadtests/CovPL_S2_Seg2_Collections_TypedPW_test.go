package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// corepayload Coverage — Segment 2: PayloadsCollection, TypedPayloadWrapper,
//                         TypedPayloadCollection, SessionInfo, AuthInfo,
//                         PagingInfo, User, UserInfo, Creators
// ══════════════════════════════════════════════════════════════════════════════

func newTestPWForSeg2() *corepayload.PayloadWrapper {
	pw, _ := corepayload.New.PayloadWrapper.Create(
		"seg2", "1", "taskType", "category",
		map[string]int{"a": 1},
	)
	return pw
}

func newTestPC() *corepayload.PayloadsCollection {
	pc := corepayload.New.PayloadsCollection.UsingCap(4)
	pw1 := newTestPWForSeg2()
	pw2, _ := corepayload.New.PayloadWrapper.Create(
		"seg2b", "2", "taskType2", "category2",
		map[string]int{"b": 2},
	)
	pc.AddsPtr(pw1, pw2)
	return pc
}

// --- PayloadsCollection Getters ---

func Test_CovPL_S2_01_Length_Count_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	pc := newTestPC()

	// Act
	actual := args.Map{"result": pc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": pc.Count() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": pc.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": pc.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	// nil
	var nilPC *corepayload.PayloadsCollection
	actual = args.Map{"result": nilPC.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S2_02_LastIndex_HasIndex(t *testing.T) {
	// Arrange
	pc := newTestPC()

	// Act
	actual := args.Map{"result": pc.LastIndex() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": pc.HasIndex(0)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": pc.HasIndex(5)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S2_03_First_Last_FirstOrDefault_LastOrDefault(t *testing.T) {
	// Arrange
	pc := newTestPC()

	// Act
	actual := args.Map{"result": pc.First() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": pc.Last() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": pc.FirstOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": pc.LastOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// dynamic
	actual = args.Map{"result": pc.FirstDynamic() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": pc.LastDynamic() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": pc.FirstOrDefaultDynamic() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": pc.LastOrDefaultDynamic() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// empty
	empty := corepayload.New.PayloadsCollection.Empty()
	actual = args.Map{"result": empty.FirstOrDefault() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": empty.LastOrDefault() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S2_04_Skip_Take_Limit_SafeLimit(t *testing.T) {
	pc := newTestPC()
	_ = pc.Skip(1)
	_ = pc.SkipDynamic(1)
	_ = pc.SkipCollection(1)
	_ = pc.Take(1)
	_ = pc.TakeDynamic(1)
	_ = pc.TakeCollection(1)
	_ = pc.LimitCollection(1)
	_ = pc.SafeLimitCollection(1)
	_ = pc.LimitDynamic(1)
	_ = pc.Limit(1)
}

func Test_CovPL_S2_05_Strings_IsEqual_IsEqualItems(t *testing.T) {
	// Arrange
	pc := newTestPC()
	ss := pc.Strings()

	// Act
	actual := args.Map{"result": len(ss) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": pc.IsEqual(pc)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": pc.IsEqualItems(pc.Items...)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	// nil
	var nilPC *corepayload.PayloadsCollection
	actual = args.Map{"result": nilPC.IsEqual(nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": nilPC.IsEqual(pc)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// --- PayloadsCollection Mutation ---

func Test_CovPL_S2_06_Add_Adds_AddsPtr(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pw := newTestPWForSeg2()
	pc.Add(*pw)
	pc.Adds(*pw)
	pc.AddsPtr(pw)
}

func Test_CovPL_S2_07_AddsPtrOptions_AddsOptions(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pw := newTestPWForSeg2()
	pc.AddsPtrOptions(false, pw)
	pc.AddsPtrOptions(true, pw)
	pc.AddsOptions(false, *pw)
	pc.AddsOptions(true, *pw)
}

func Test_CovPL_S2_08_AddsIf_InsertAt(t *testing.T) {
	pc := newTestPC()
	pw := newTestPWForSeg2()
	pc.AddsIf(false, *pw)
	pc.AddsIf(true, *pw)
	pc.InsertAt(0, *pw)
}

func Test_CovPL_S2_09_ConcatNew_ConcatNewPtr(t *testing.T) {
	// Arrange
	pc := newTestPC()
	pw := newTestPWForSeg2()
	c := pc.ConcatNew(*pw)

	// Act
	actual := args.Map{"result": c.Length() < 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 3", actual)
	c2 := pc.ConcatNewPtr(pw)
	actual = args.Map{"result": c2.Length() < 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 3", actual)
}

func Test_CovPL_S2_10_Reverse(t *testing.T) {
	pc := newTestPC()
	pc.Reverse()
	// single
	single := corepayload.New.PayloadsCollection.UsingCap(1)
	single.Add(*newTestPWForSeg2())
	single.Reverse()
	// 3 items
	triple := corepayload.New.PayloadsCollection.UsingCap(3)
	triple.Add(*newTestPWForSeg2())
	triple.Add(*newTestPWForSeg2())
	triple.Add(*newTestPWForSeg2())
	triple.Reverse()
}

func Test_CovPL_S2_11_Clone_ClonePtr(t *testing.T) {
	// Arrange
	pc := newTestPC()
	c := pc.Clone()

	// Act
	actual := args.Map{"result": c.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	cp := pc.ClonePtr()
	actual = args.Map{"result": cp == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	var nilPC *corepayload.PayloadsCollection
	actual = args.Map{"result": nilPC.ClonePtr() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S2_12_Clear_Dispose(t *testing.T) {
	// Arrange
	pc := newTestPC()
	pc.Clear()

	// Act
	actual := args.Map{"result": pc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	pc2 := newTestPC()
	pc2.Dispose()
	var nilPC *corepayload.PayloadsCollection
	nilPC.Clear()
	nilPC.Dispose()
}

// --- PayloadsCollection Filter ---

func Test_CovPL_S2_13_Filter_FilterWithLimit(t *testing.T) {
	// Arrange
	pc := newTestPC()
	items := pc.Filter(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})

	// Act
	actual := args.Map{"result": len(items) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	items2 := pc.FilterWithLimit(1, func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})
	actual = args.Map{"result": len(items2) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovPL_S2_14_FirstByFilter_FirstById_FirstByCategory_FirstByTaskType_FirstByEntityType(t *testing.T) {
	// Arrange
	pc := newTestPC()
	f := pc.FirstByFilter(func(pw *corepayload.PayloadWrapper) bool {
		return pw.IsIdentifier("1")
	})

	// Act
	actual := args.Map{"result": f == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_ = pc.FirstById("1")
	_ = pc.FirstByCategory("category")
	_ = pc.FirstByTaskType("taskType")
	_ = pc.FirstByEntityType("unknown")
}

func Test_CovPL_S2_15_FilterCollection_SkipFilterCollection(t *testing.T) {
	// Arrange
	pc := newTestPC()
	fc := pc.FilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})

	// Act
	actual := args.Map{"result": fc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	sc := pc.SkipFilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return false, false
	})
	actual = args.Map{"result": sc.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_CovPL_S2_16_FilterCollectionByIds_FilterNameCollection_FilterCategory_FilterEntityType_FilterTaskType(t *testing.T) {
	pc := newTestPC()
	_ = pc.FilterCollectionByIds("1")
	_ = pc.FilterNameCollection("seg2")
	_ = pc.FilterCategoryCollection("category")
	_ = pc.FilterEntityTypeCollection("unknown")
	_ = pc.FilterTaskTypeCollection("taskType")
}

// --- PayloadsCollection Paging ---

func Test_CovPL_S2_17_GetPagesSize_GetPagedCollection_GetSinglePageCollection(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		pc.Add(*newTestPWForSeg2())
	}

	// Act
	actual := args.Map{"result": pc.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual = args.Map{"result": pc.GetPagesSize(0) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	pages := pc.GetPagedCollection(3)
	actual = args.Map{"result": len(pages) < 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	_ = pc.GetSinglePageCollection(3, 2)
	// small collection
	small := newTestPC()
	pages2 := small.GetPagedCollection(10)
	actual = args.Map{"result": len(pages2) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = small.GetSinglePageCollection(10, 1)
}

// --- PayloadsCollection JSON ---

func Test_CovPL_S2_18_StringsUsingFmt_JoinUsingFmt(t *testing.T) {
	// Arrange
	pc := newTestPC()
	ss := pc.StringsUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.PayloadName()
	})

	// Act
	actual := args.Map{"result": len(ss) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	j := pc.JoinUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.PayloadName()
	}, ",")
	actual = args.Map{"result": j == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovPL_S2_19_JsonStrings_JoinJsonStrings_Join_JoinCsv_JoinCsvLine(t *testing.T) {
	pc := newTestPC()
	_ = pc.JsonStrings()
	_ = pc.JoinJsonStrings(",")
	_ = pc.Join(",")
	_ = pc.JoinCsv()
	_ = pc.JoinCsvLine()
}

func Test_CovPL_S2_20_JsonString_String_PrettyJsonString_CsvStrings(t *testing.T) {
	// Arrange
	pc := newTestPC()

	// Act
	actual := args.Map{"result": pc.JsonString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual = args.Map{"result": pc.String() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual = args.Map{"result": pc.PrettyJsonString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	_ = pc.CsvStrings()
	// empty
	empty := corepayload.New.PayloadsCollection.Empty()
	actual = args.Map{"result": empty.JsonString() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovPL_S2_21_Json_JsonPtr_ParseInject_AsJsoner(t *testing.T) {
	// Arrange
	pc := newTestPC()
	_ = pc.Json()
	jp := pc.JsonPtr()

	// Act
	actual := args.Map{"result": jp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	pc2 := corepayload.New.PayloadsCollection.Empty()
	_, err := pc2.ParseInjectUsingJson(jp)
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = pc.AsJsonContractsBinder()
	_ = pc.AsJsoner()
	_ = pc.AsJsonParseSelfInjector()
	_ = pc.JsonParseSelfInject(jp)
}

func Test_CovPL_S2_22_ParseInjectUsingJsonMust(t *testing.T) {
	pc := newTestPC()
	jp := pc.JsonPtr()
	pc2 := corepayload.New.PayloadsCollection.Empty()
	_ = pc2.ParseInjectUsingJsonMust(jp)
}

// --- PayloadsCollection Creator ---

func Test_CovPL_S2_23_NewPC_Empty_UsingCap_UsingWrappers(t *testing.T) {
	_ = corepayload.New.PayloadsCollection.Empty()
	_ = corepayload.New.PayloadsCollection.UsingCap(5)
	pw := newTestPWForSeg2()
	_ = corepayload.New.PayloadsCollection.UsingWrappers(pw)
	_ = corepayload.New.PayloadsCollection.UsingWrappers()
}

func Test_CovPL_S2_24_NewPC_Deserialize(t *testing.T) {
	// Arrange
	pc := newTestPC()
	b, _ := corejson.Serialize.Raw(pc)
	pc2, err := corepayload.New.PayloadsCollection.Deserialize(b)

	// Act
	actual := args.Map{"result": err != nil || pc2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S2_25_NewPC_DeserializeToMany(t *testing.T) {
	// Arrange
	pc := newTestPC()
	pcs := []*corepayload.PayloadsCollection{pc}
	b, _ := corejson.Serialize.Raw(pcs)
	many, err := corepayload.New.PayloadsCollection.DeserializeToMany(b)

	// Act
	actual := args.Map{"result": err != nil || len(many) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovPL_S2_26_NewPC_DeserializeUsingJsonResult(t *testing.T) {
	// Arrange
	pc := newTestPC()
	jr := pc.JsonPtr()
	pc2, err := corepayload.New.PayloadsCollection.DeserializeUsingJsonResult(jr)

	// Act
	actual := args.Map{"result": err != nil || pc2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// --- TypedPayloadWrapper ---

func Test_CovPL_S2_30_TypedPW_Create_TypedData(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, err := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})

	// Act
	actual := args.Map{"result": err != nil || tw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": tw.TypedData().A != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected A=1", actual)
	actual = args.Map{"result": tw.Data().A != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected A=1", actual)
	actual = args.Map{"result": tw.IsParsed()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovPL_S2_31_TypedPW_Accessors(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})

	// Act
	actual := args.Map{"result": tw.Name() != "n"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected n", actual)
	actual = args.Map{"result": tw.Identifier() != "1"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": tw.IdString() != "1"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": tw.IdInteger() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": tw.EntityType() != "e"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected e", actual)
	// nil
	var nilTW *corepayload.TypedPayloadWrapper[D]
	actual = args.Map{"result": nilTW.Name() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": nilTW.IsParsed()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S2_32_TypedPW_ErrorHandling(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})

	// Act
	actual := args.Map{"result": tw.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": tw.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": tw.HasItems()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": tw.HasSafeItems()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": tw.Error() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S2_33_TypedPW_StringRepresentation(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})

	// Act
	actual := args.Map{"result": tw.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual = args.Map{"result": tw.PrettyJsonString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual = args.Map{"result": tw.JsonString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// nil
	var nilTW *corepayload.TypedPayloadWrapper[D]
	actual = args.Map{"result": nilTW.String() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovPL_S2_34_TypedPW_JSON(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	_ = tw.Json()
	_ = tw.JsonPtr()
	b, err := tw.MarshalJSON()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[D]("x", "2", "e", D{})
	err2 := tw2.UnmarshalJSON(b)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_, _ = tw.Serialize()
	_ = tw.SerializeMust()
	_ = tw.TypedDataJson()
	_ = tw.TypedDataJsonPtr()
	_, _ = tw.TypedDataJsonBytes()
}

func Test_CovPL_S2_35_TypedPW_GetAs(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "1", "e", "hello")
	s, ok := tw.GetAsString()

	// Act
	actual := args.Map{"result": ok || s != "hello"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	_ = tw.ValueString()

	twi, _ := corepayload.NewTypedPayloadWrapperFrom[int]("n", "1", "e", 42)
	i, ok2 := twi.GetAsInt()
	actual = args.Map{"result": ok2 || i != 42}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	_ = twi.ValueInt()

	twb, _ := corepayload.NewTypedPayloadWrapperFrom[bool]("n", "1", "e", true)
	b, ok3 := twb.GetAsBool()
	actual = args.Map{"result": ok3 || !b}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	_ = twb.ValueBool()

	// non-matching
	_, ok4 := tw.GetAsInt()
	actual = args.Map{"result": ok4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, ok5 := tw.GetAsInt64()
	actual = args.Map{"result": ok5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, ok6 := tw.GetAsFloat64()
	actual = args.Map{"result": ok6}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, ok7 := tw.GetAsFloat32()
	actual = args.Map{"result": ok7}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, ok8 := tw.GetAsBytes()
	actual = args.Map{"result": ok8}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, ok9 := tw.GetAsStrings()
	actual = args.Map{"result": ok9}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S2_36_TypedPW_Setters(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	tw.SetName("new")
	tw.SetIdentifier("2")
	tw.SetEntityType("new_e")
	tw.SetCategoryName("cat")
	err := tw.SetTypedData(D{A: 5})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	actual = args.Map{"result": tw.TypedData().A != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_CovPL_S2_37_TypedPW_Clone_ToPayloadWrapper_Reparse(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	cp, err := tw.ClonePtr(true)

	// Act
	actual := args.Map{"result": err != nil || cp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_, _ = tw.Clone(true)
	_ = tw.ToPayloadWrapper()
	_ = tw.PayloadWrapperValue()
	_ = tw.DynamicPayloads()
	_ = tw.PayloadsString()
	_ = tw.Length()
	actual = args.Map{"result": tw.IsNull()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	err2 := tw.Reparse()
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	// nil
	var nilTW *corepayload.TypedPayloadWrapper[D]
	c, _ := nilTW.ClonePtr(true)
	actual = args.Map{"result": c != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S2_38_TypedPW_Clear_Dispose(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	tw.Clear()
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	tw2.Dispose()
	var nilTW *corepayload.TypedPayloadWrapper[D]
	nilTW.Clear()
	nilTW.Dispose()
}

func Test_CovPL_S2_39_TypedPW_OtherAccessors(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})

	// Act
	actual := args.Map{"result": tw.CategoryName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": tw.TaskTypeName() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": tw.HasManyRecords()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": tw.HasSingleRecord()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": tw.Attributes() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	_ = tw.InitializeAttributesOnNull()
}

func Test_CovPL_S2_40_TypedPW_NewFromWrapper(t *testing.T) {
	// Arrange
	type D struct{ A int }
	pw := newTestPWForSeg2()
	tw, err := corepayload.NewTypedPayloadWrapper[D](pw)

	// Act
	actual := args.Map{"result": err != nil || tw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// nil wrapper
	tw2, err2 := corepayload.NewTypedPayloadWrapper[D](nil)
	actual = args.Map{"result": err2 == nil || tw2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// --- TypedPayloadCollection ---

func Test_CovPL_S2_50_TPC_Create_Length_IsEmpty_HasItems(t *testing.T) {
	// Arrange
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](5)

	// Act
	actual := args.Map{"result": col.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": col.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col.Add(tw)
	actual = args.Map{"result": col.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": col.HasItems()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": col.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": col.Count() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovPL_S2_51_TPC_Empty_From(t *testing.T) {
	// Arrange
	type D struct{ A int }
	empty := corepayload.EmptyTypedPayloadCollection[D]()

	// Act
	actual := args.Map{"result": empty.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	from := corepayload.TypedPayloadCollectionFrom[D]([]*corepayload.TypedPayloadWrapper[D]{tw})
	actual = args.Map{"result": from.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovPL_S2_52_TPC_FromPayloads(t *testing.T) {
	// Arrange
	type D struct{ A int }
	pc := newTestPC()
	col := corepayload.TypedPayloadCollectionFromPayloads[D](pc)

	// Act
	actual := args.Map{"result": col.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	// nil
	nilCol := corepayload.TypedPayloadCollectionFromPayloads[D](nil)
	actual = args.Map{"result": nilCol.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S2_53_TPC_ElementAccess(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw1, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "2", "e", D{A: 2})
	col := corepayload.NewTypedPayloadCollection[D](2)
	col.Add(tw1)
	col.Add(tw2)

	// Act
	actual := args.Map{"result": col.First().Data().A != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": col.Last().Data().A != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": col.FirstOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": col.LastOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": col.SafeAt(0) == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": col.SafeAt(10) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	// empty
	empty := corepayload.EmptyTypedPayloadCollection[D]()
	actual = args.Map{"result": empty.FirstOrDefault() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S2_54_TPC_Add_AddLock_Adds_AddCollection_RemoveAt(t *testing.T) {
	// Arrange
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](5)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col.AddLock(tw)
	col.Adds(tw)
	col2 := corepayload.NewTypedPayloadCollection[D](2)
	col2.Add(tw)
	col.AddCollection(col2)
	ok := col.RemoveAt(0)

	// Act
	actual := args.Map{"result": ok}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	ok2 := col.RemoveAt(-1)
	actual = args.Map{"result": ok2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S2_55_TPC_Iteration_Filter(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col := corepayload.NewTypedPayloadCollection[D](2)
	col.Add(tw)
	col.ForEach(func(i int, item *corepayload.TypedPayloadWrapper[D]) {})
	col.ForEachData(func(i int, data D) {})
	col.ForEachBreak(func(i int, item *corepayload.TypedPayloadWrapper[D]) bool { return false })
	fc := col.Filter(func(item *corepayload.TypedPayloadWrapper[D]) bool { return true })

	// Act
	actual := args.Map{"result": fc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	fd := col.FilterByData(func(d D) bool { return d.A == 1 })
	actual = args.Map{"result": fd.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = col.FirstByFilter(func(item *corepayload.TypedPayloadWrapper[D]) bool { return true })
	_ = col.FirstByData(func(d D) bool { return d.A == 1 })
	_ = col.FirstByName("n")
	_ = col.FirstById("1")
	_ = col.CountFunc(func(item *corepayload.TypedPayloadWrapper[D]) bool { return true })
}

func Test_CovPL_S2_56_TPC_Skip_Take_AllData_AllNames_AllIdentifiers(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col := corepayload.NewTypedPayloadCollection[D](2)
	col.Add(tw)
	_ = col.Skip(0)
	_ = col.Take(1)
	data := col.AllData()

	// Act
	actual := args.Map{"result": len(data) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	names := col.AllNames()
	actual = args.Map{"result": len(names) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	ids := col.AllIdentifiers()
	actual = args.Map{"result": len(ids) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	// empty
	empty := corepayload.EmptyTypedPayloadCollection[D]()
	actual = args.Map{"result": len(empty.AllData()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S2_57_TPC_ToPayloadsCollection_Clone(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col := corepayload.NewTypedPayloadCollection[D](2)
	col.Add(tw)
	pc := col.ToPayloadsCollection()

	// Act
	actual := args.Map{"result": pc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	c, err := col.Clone()
	actual = args.Map{"result": err != nil || c.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = col.CloneMust()
	_, _ = col.ConcatNew(tw)
	// empty
	empty := corepayload.EmptyTypedPayloadCollection[D]()
	actual = args.Map{"result": empty.ToPayloadsCollection().Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S2_58_TPC_Clear_Dispose(t *testing.T) {
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](2)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col.Add(tw)
	col.Clear()
	col.Dispose()
	var nilCol *corepayload.TypedPayloadCollection[D]
	nilCol.Clear()
	nilCol.Dispose()
}

func Test_CovPL_S2_59_TPC_LengthLock_IsEmptyLock_HasIndex_LastIndex(t *testing.T) {
	// Arrange
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](2)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col.Add(tw)

	// Act
	actual := args.Map{"result": col.LengthLock() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": col.IsEmptyLock()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": col.LastIndex() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": col.HasIndex(0)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovPL_S2_60_TPC_IsValid_HasErrors_Errors_FirstError_MergedError(t *testing.T) {
	// Arrange
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](2)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col.Add(tw)

	// Act
	actual := args.Map{"result": col.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": col.HasErrors()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": len(col.Errors()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": col.FirstError() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": col.MergedError() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S2_61_TPC_Deserialization(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col := corepayload.NewTypedPayloadCollection[D](1)
	col.Add(tw)
	pc := col.ToPayloadsCollection()
	// TypedPayloadCollectionDeserialize calls DeserializeToMany which expects
	// a JSON array [{},...], not {"Items":[...]} — serialize Items directly
	b, _ := corejson.Serialize.Raw(pc.Items)
	col2, err := corepayload.TypedPayloadCollectionDeserialize[D](b)

	// Act
	actual := args.Map{"result": err != nil || col2.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = corepayload.TypedPayloadCollectionDeserializeMust[D](b)
}

func Test_CovPL_S2_62_TPC_Single_FromData(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	single := corepayload.NewTypedPayloadCollectionSingle[D](tw)

	// Act
	actual := args.Map{"result": single.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	nilSingle := corepayload.NewTypedPayloadCollectionSingle[D](nil)
	actual = args.Map{"result": nilSingle.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	fromData, err := corepayload.NewTypedPayloadCollectionFromData[D]("n", []D{{A: 1}, {A: 2}})
	actual = args.Map{"result": err != nil || fromData.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	_ = corepayload.NewTypedPayloadCollectionFromDataMust[D]("n", []D{{A: 1}})
	emptyData, err2 := corepayload.NewTypedPayloadCollectionFromData[D]("n", []D{})
	actual = args.Map{"result": err2 != nil || emptyData.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// --- TypedPayloadWrapper Creator functions ---

func Test_CovPL_S2_65_TypedPW_Creators(t *testing.T) {
	// Arrange
	type D struct{ A int }
	_, err := corepayload.TypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_, err2 := corepayload.TypedPayloadWrapperRecord[D]("n", "1", "t", "c", D{A: 1})
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	// TypedPayloadWrapperRecords calls SafeTypeNameOfSliceOrSingle(false, data)
	// which calls SliceFirstItemTypeName → rt.Elem() — data MUST be a slice
	_, err3 := corepayload.TypedPayloadWrapperRecords[[]D]("n", "1", "t", "c", []D{{A: 1}})
	actual = args.Map{"result": err3 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_, err4 := corepayload.TypedPayloadWrapperNameIdRecord[D]("n", "1", D{A: 1})
	actual = args.Map{"result": err4 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_, err5 := corepayload.TypedPayloadWrapperNameIdCategory[D]("n", "1", "c", D{A: 1})
	actual = args.Map{"result": err5 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_, err6 := corepayload.TypedPayloadWrapperAll[D]("n", "1", "t", "e", "c", false, D{A: 1}, nil)
	actual = args.Map{"result": err6 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovPL_S2_66_TypedPW_Deserialize(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	b, _ := tw.Serialize()
	tw2, err := corepayload.TypedPayloadWrapperDeserialize[D](b)

	// Act
	actual := args.Map{"result": err != nil || tw2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	jr := tw.JsonPtr()
	tw3, err2 := corepayload.TypedPayloadWrapperDeserializeUsingJsonResult[D](jr)
	actual = args.Map{"result": err2 != nil || tw3 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S2_67_TypedPW_DeserializeToMany(t *testing.T) {
	// Arrange
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	pws := []*corepayload.PayloadWrapper{tw.ToPayloadWrapper()}
	b, _ := corejson.Serialize.Raw(pws)
	many, err := corepayload.TypedPayloadWrapperDeserializeToMany[D](b)

	// Act
	actual := args.Map{"result": err != nil || len(many) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// --- SessionInfo ---

func Test_CovPL_S2_70_SessionInfo(t *testing.T) {
	// Arrange
	si := corepayload.SessionInfo{Id: "5", User: &corepayload.User{Name: "u"}, SessionPath: "/p"}

	// Act
	actual := args.Map{"result": si.IdentifierInteger() != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual = args.Map{"result": si.IdentifierUnsignedInteger() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual = args.Map{"result": si.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": si.IsValid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": si.IsUserNameEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": si.IsUserEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": si.HasUser()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": si.IsUsernameEqual("u")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	c := si.Clone()
	_ = c
	cp := si.ClonePtr()
	actual = args.Map{"result": cp == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_ = si.Ptr()
	// empty
	empty := corepayload.SessionInfo{}
	actual = args.Map{"result": empty.IdentifierInteger() >= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual = args.Map{"result": empty.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	// nil
	var nilSI *corepayload.SessionInfo
	actual = args.Map{"result": nilSI.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": nilSI.ClonePtr() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// --- AuthInfo ---

func Test_CovPL_S2_71_AuthInfo(t *testing.T) {
	// Arrange
	ai := corepayload.AuthInfo{
		Identifier:   "5",
		ActionType:   "act",
		ResourceName: "res",
		SessionInfo:  &corepayload.SessionInfo{Id: "1"},
		UserInfo:     &corepayload.UserInfo{User: &corepayload.User{Name: "u"}},
	}

	// Act
	actual := args.Map{"result": ai.IdentifierInteger() != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual = args.Map{"result": ai.IdentifierUnsignedInteger() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual = args.Map{"result": ai.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": ai.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": ai.IsValid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": ai.HasActionType()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": ai.HasResourceName()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": ai.HasUserInfo()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": ai.HasSessionInfo()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": ai.IsActionTypeEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": ai.IsResourceNameEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_ = ai.String()
	_ = ai.PrettyJsonString()
	_ = ai.Json()
	_ = ai.JsonPtr()
	c := ai.Clone()
	_ = c
	cp := ai.ClonePtr()
	actual = args.Map{"result": cp == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_ = ai.Ptr()
	// nil
	var nilAI *corepayload.AuthInfo
	actual = args.Map{"result": nilAI.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": nilAI.ClonePtr() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	// empty
	empty := corepayload.AuthInfo{}
	actual = args.Map{"result": empty.IdentifierInteger() >= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_CovPL_S2_72_AuthInfo_Setters(t *testing.T) {
	// Arrange
	ai := &corepayload.AuthInfo{}
	ai.SetActionType("act")
	ai.SetResourceName("res")
	ai.SetIdentifier("5")
	ai.SetSessionInfo(&corepayload.SessionInfo{Id: "1"})
	u := &corepayload.User{Name: "u"}
	ai.SetUserInfo(&corepayload.UserInfo{User: u})
	ai.SetUser(u)
	ai.SetSystemUser(u)
	ai.SetUserSystemUser(u, u)
	// nil setters
	var nilAI *corepayload.AuthInfo
	r := nilAI.SetActionType("act")

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_ = nilAI.SetResourceName("res")
	_ = nilAI.SetIdentifier("5")
	_ = nilAI.SetSessionInfo(nil)
	_ = nilAI.SetUserInfo(nil)
	_ = nilAI.SetUser(u)
	_ = nilAI.SetSystemUser(u)
	_ = nilAI.SetUserSystemUser(u, u)
}

// --- PagingInfo ---

func Test_CovPL_S2_73_PagingInfo(t *testing.T) {
	// Arrange
	pi := corepayload.PagingInfo{
		CurrentPageIndex: 1,
		TotalPages:       5,
		PerPageItems:     10,
		TotalItems:       50,
	}

	// Act
	actual := args.Map{"result": pi.IsEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": pi.HasTotalPages()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": pi.HasCurrentPageIndex()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": pi.HasPerPageItems()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": pi.HasTotalItems()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": pi.IsInvalidTotalPages()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": pi.IsInvalidCurrentPageIndex()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": pi.IsInvalidPerPageItems()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": pi.IsInvalidTotalItems()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": pi.IsEqual(&pi)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	c := pi.Clone()
	_ = c
	cp := pi.ClonePtr()
	actual = args.Map{"result": cp == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// nil
	var nilPI *corepayload.PagingInfo
	actual = args.Map{"result": nilPI.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": nilPI.ClonePtr() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": nilPI.IsEqual(nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": nilPI.IsEqual(&pi)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// --- User ---

func Test_CovPL_S2_74_User(t *testing.T) {
	// Arrange
	u := corepayload.User{
		Identifier:   "5",
		Name:         "u",
		Type:         "admin",
		AuthToken:    "tok",
		PasswordHash: "hash",
		IsSystemUser: false,
	}

	// Act
	actual := args.Map{"result": u.IdentifierInteger() != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual = args.Map{"result": u.IdentifierUnsignedInteger() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual = args.Map{"result": u.HasAuthToken()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": u.HasPasswordHash()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": u.IsPasswordHashEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": u.IsAuthTokenEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": u.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": u.IsValidUser()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": u.IsNameEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": u.IsNameEqual("u")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": u.IsNotSystemUser()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": u.IsVirtualUser()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": u.HasType()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": u.IsTypeEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_ = u.String()
	_ = u.PrettyJsonString()
	_ = u.Json()
	_ = u.JsonPtr()
	_, _ = u.Serialize()
	_ = u.Deserialize([]byte(`{"Name":"x"}`))
	c := u.Clone()
	_ = c
	cp := u.ClonePtr()
	actual = args.Map{"result": cp == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_ = u.Ptr()
	// nil
	var nilU *corepayload.User
	actual = args.Map{"result": nilU.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": nilU.ClonePtr() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	// empty
	empty := corepayload.User{}
	actual = args.Map{"result": empty.IdentifierInteger() >= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

// --- UserInfo ---

func Test_CovPL_S2_75_UserInfo(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "u"}
	su := &corepayload.User{Name: "sys", IsSystemUser: true}
	ui := corepayload.UserInfo{User: u, SystemUser: su}

	// Act
	actual := args.Map{"result": ui.HasUser()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": ui.HasSystemUser()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": ui.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": ui.IsUserEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": ui.IsSystemUserEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	c := ui.Clone()
	_ = c
	cp := ui.ClonePtr()
	actual = args.Map{"result": cp == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_ = ui.Ptr()
	_ = ui.ToNonPtr()
	// setters
	ui2 := &corepayload.UserInfo{}
	ui2.SetUser(u)
	ui2.SetSystemUser(su)
	ui2.SetUserSystemUser(u, su)
	// nil setters
	var nilUI *corepayload.UserInfo
	r := nilUI.SetUser(u)
	actual = args.Map{"result": r == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_ = nilUI.SetSystemUser(su)
	_ = nilUI.SetUserSystemUser(u, su)
	actual = args.Map{"result": nilUI.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": nilUI.ClonePtr() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	nonPtr := nilUI.ToNonPtr()
	actual = args.Map{"result": nonPtr.HasUser()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// --- User Creator ---

func Test_CovPL_S2_76_NewUser_Creator(t *testing.T) {
	_ = corepayload.New.User.Empty()
	_ = corepayload.New.User.Create(false, "u", "t")
	_ = corepayload.New.User.NonSysCreate("u", "t")
	_ = corepayload.New.User.NonSysCreateId("1", "u", "t")
	_ = corepayload.New.User.System("u", "t")
	_ = corepayload.New.User.SystemId("1", "u", "t")
	_ = corepayload.New.User.UsingName("u")
	_ = corepayload.New.User.All(false, "1", "u", "t", "tok", "hash")
}

func Test_CovPL_S2_77_NewUser_Deserialize_CastOrDeserializeFrom(t *testing.T) {
	// Arrange
	u := corepayload.New.User.Create(false, "u", "t")
	b, _ := u.Serialize()
	u2, err := corepayload.New.User.Deserialize(b)

	// Act
	actual := args.Map{"result": err != nil || u2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	u3, err2 := corepayload.New.User.CastOrDeserializeFrom(u)
	actual = args.Map{"result": err2 != nil || u3 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// nil
	_, err3 := corepayload.New.User.CastOrDeserializeFrom(nil)
	actual = args.Map{"result": err3 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// --- PayloadCreateInstructionTypeStringer ---

func Test_CovPL_S2_80_PayloadCreateInstructionTypeStringer(t *testing.T) {
	// Arrange
	type stringer struct{ v string }
	s := stringer{v: "task"}
	// can't use stringer directly, use a real Stringer
	// Use a concrete type implementing Stringer
	inst := corepayload.PayloadCreateInstructionTypeStringer{
		Name:                 "n",
		Identifier:           "1",
		TaskTypeNameStringer: stringerImpl{"task"},
		CategoryNameStringer: stringerImpl{"cat"},
	}
	pci := inst.PayloadCreateInstruction()

	// Act
	actual := args.Map{"result": pci == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_ = s
}

type stringerImpl struct{ v string }

func (s stringerImpl) String() string { return s.v }
