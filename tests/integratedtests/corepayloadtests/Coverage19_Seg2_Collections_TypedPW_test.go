package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
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
	pc := newTestPC()
	if pc.Length() != 2 {
		t.Fatal("expected 2")
	}
	if pc.Count() != 2 {
		t.Fatal("expected 2")
	}
	if pc.IsEmpty() {
		t.Fatal("expected false")
	}
	if !pc.HasAnyItem() {
		t.Fatal("expected true")
	}
	// nil
	var nilPC *corepayload.PayloadsCollection
	if nilPC.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S2_02_LastIndex_HasIndex(t *testing.T) {
	pc := newTestPC()
	if pc.LastIndex() != 1 {
		t.Fatal("expected 1")
	}
	if !pc.HasIndex(0) {
		t.Fatal("expected true")
	}
	if pc.HasIndex(5) {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S2_03_First_Last_FirstOrDefault_LastOrDefault(t *testing.T) {
	pc := newTestPC()
	if pc.First() == nil {
		t.Fatal("expected non-nil")
	}
	if pc.Last() == nil {
		t.Fatal("expected non-nil")
	}
	if pc.FirstOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	if pc.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	// dynamic
	if pc.FirstDynamic() == nil {
		t.Fatal("expected non-nil")
	}
	if pc.LastDynamic() == nil {
		t.Fatal("expected non-nil")
	}
	if pc.FirstOrDefaultDynamic() == nil {
		t.Fatal("expected non-nil")
	}
	if pc.LastOrDefaultDynamic() == nil {
		t.Fatal("expected non-nil")
	}
	// empty
	empty := corepayload.New.PayloadsCollection.Empty()
	if empty.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	if empty.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
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
	pc := newTestPC()
	ss := pc.Strings()
	if len(ss) != 2 {
		t.Fatal("expected 2")
	}
	if !pc.IsEqual(pc) {
		t.Fatal("expected true")
	}
	if !pc.IsEqualItems(pc.Items...) {
		t.Fatal("expected true")
	}
	// nil
	var nilPC *corepayload.PayloadsCollection
	if !nilPC.IsEqual(nil) {
		t.Fatal("expected true")
	}
	if nilPC.IsEqual(pc) {
		t.Fatal("expected false")
	}
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
	pc := newTestPC()
	pw := newTestPWForSeg2()
	c := pc.ConcatNew(*pw)
	if c.Length() < 3 {
		t.Fatal("expected >= 3")
	}
	c2 := pc.ConcatNewPtr(pw)
	if c2.Length() < 3 {
		t.Fatal("expected >= 3")
	}
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
	pc := newTestPC()
	c := pc.Clone()
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
	cp := pc.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}
	var nilPC *corepayload.PayloadsCollection
	if nilPC.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S2_12_Clear_Dispose(t *testing.T) {
	pc := newTestPC()
	pc.Clear()
	if pc.Length() != 0 {
		t.Fatal("expected 0")
	}
	pc2 := newTestPC()
	pc2.Dispose()
	var nilPC *corepayload.PayloadsCollection
	nilPC.Clear()
	nilPC.Dispose()
}

// --- PayloadsCollection Filter ---

func Test_CovPL_S2_13_Filter_FilterWithLimit(t *testing.T) {
	pc := newTestPC()
	items := pc.Filter(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})
	if len(items) != 2 {
		t.Fatal("expected 2")
	}
	items2 := pc.FilterWithLimit(1, func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})
	if len(items2) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovPL_S2_14_FirstByFilter_FirstById_FirstByCategory_FirstByTaskType_FirstByEntityType(t *testing.T) {
	pc := newTestPC()
	f := pc.FirstByFilter(func(pw *corepayload.PayloadWrapper) bool {
		return pw.IsIdentifier("1")
	})
	if f == nil {
		t.Fatal("expected non-nil")
	}
	_ = pc.FirstById("1")
	_ = pc.FirstByCategory("category")
	_ = pc.FirstByTaskType("taskType")
	_ = pc.FirstByEntityType("unknown")
}

func Test_CovPL_S2_15_FilterCollection_SkipFilterCollection(t *testing.T) {
	pc := newTestPC()
	fc := pc.FilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})
	if fc.Length() != 2 {
		t.Fatal("expected 2")
	}
	sc := pc.SkipFilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return false, false
	})
	if sc.Length() != 2 {
		t.Fatal("expected 2")
	}
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
	pc := corepayload.New.PayloadsCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		pc.Add(*newTestPWForSeg2())
	}
	if pc.GetPagesSize(3) != 4 {
		t.Fatal("expected 4")
	}
	if pc.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	pages := pc.GetPagedCollection(3)
	if len(pages) < 3 {
		t.Fatal("expected at least 3")
	}
	_ = pc.GetSinglePageCollection(3, 2)
	// small collection
	small := newTestPC()
	pages2 := small.GetPagedCollection(10)
	if len(pages2) != 1 {
		t.Fatal("expected 1")
	}
	_ = small.GetSinglePageCollection(10, 1)
}

// --- PayloadsCollection JSON ---

func Test_CovPL_S2_18_StringsUsingFmt_JoinUsingFmt(t *testing.T) {
	pc := newTestPC()
	ss := pc.StringsUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.PayloadName()
	})
	if len(ss) != 2 {
		t.Fatal("expected 2")
	}
	j := pc.JoinUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.PayloadName()
	}, ",")
	if j == "" {
		t.Fatal("expected non-empty")
	}
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
	pc := newTestPC()
	if pc.JsonString() == "" {
		t.Fatal("expected non-empty")
	}
	if pc.String() == "" {
		t.Fatal("expected non-empty")
	}
	if pc.PrettyJsonString() == "" {
		t.Fatal("expected non-empty")
	}
	_ = pc.CsvStrings()
	// empty
	empty := corepayload.New.PayloadsCollection.Empty()
	if empty.JsonString() != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovPL_S2_21_Json_JsonPtr_ParseInject_AsJsoner(t *testing.T) {
	pc := newTestPC()
	_ = pc.Json()
	jp := pc.JsonPtr()
	if jp == nil {
		t.Fatal("expected non-nil")
	}
	pc2 := corepayload.New.PayloadsCollection.Empty()
	_, err := pc2.ParseInjectUsingJson(jp)
	if err != nil {
		t.Fatal("expected no error")
	}
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
	pc := newTestPC()
	b, _ := corejson.Serialize.Raw(pc)
	pc2, err := corepayload.New.PayloadsCollection.Deserialize(b)
	if err != nil || pc2 == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S2_25_NewPC_DeserializeToMany(t *testing.T) {
	pc := newTestPC()
	pcs := []*corepayload.PayloadsCollection{pc}
	b, _ := corejson.Serialize.Raw(pcs)
	many, err := corepayload.New.PayloadsCollection.DeserializeToMany(b)
	if err != nil || len(many) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovPL_S2_26_NewPC_DeserializeUsingJsonResult(t *testing.T) {
	pc := newTestPC()
	jr := pc.JsonPtr()
	pc2, err := corepayload.New.PayloadsCollection.DeserializeUsingJsonResult(jr)
	if err != nil || pc2 == nil {
		t.Fatal("expected non-nil")
	}
}

// --- TypedPayloadWrapper ---

func Test_CovPL_S2_30_TypedPW_Create_TypedData(t *testing.T) {
	type D struct{ A int }
	tw, err := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	if err != nil || tw == nil {
		t.Fatal("expected non-nil")
	}
	if tw.TypedData().A != 1 {
		t.Fatal("expected A=1")
	}
	if tw.Data().A != 1 {
		t.Fatal("expected A=1")
	}
	if !tw.IsParsed() {
		t.Fatal("expected true")
	}
}

func Test_CovPL_S2_31_TypedPW_Accessors(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	if tw.Name() != "n" {
		t.Fatal("expected n")
	}
	if tw.Identifier() != "1" {
		t.Fatal("expected 1")
	}
	if tw.IdString() != "1" {
		t.Fatal("expected 1")
	}
	if tw.IdInteger() != 1 {
		t.Fatal("expected 1")
	}
	if tw.EntityType() != "e" {
		t.Fatal("expected e")
	}
	// nil
	var nilTW *corepayload.TypedPayloadWrapper[D]
	if nilTW.Name() != "" {
		t.Fatal("expected empty")
	}
	if nilTW.IsParsed() {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S2_32_TypedPW_ErrorHandling(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	if tw.HasError() {
		t.Fatal("expected false")
	}
	if tw.IsEmpty() {
		t.Fatal("expected false")
	}
	if !tw.HasItems() {
		t.Fatal("expected true")
	}
	if !tw.HasSafeItems() {
		t.Fatal("expected true")
	}
	if tw.Error() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S2_33_TypedPW_StringRepresentation(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	if tw.String() == "" {
		t.Fatal("expected non-empty")
	}
	if tw.PrettyJsonString() == "" {
		t.Fatal("expected non-empty")
	}
	if tw.JsonString() == "" {
		t.Fatal("expected non-empty")
	}
	// nil
	var nilTW *corepayload.TypedPayloadWrapper[D]
	if nilTW.String() != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovPL_S2_34_TypedPW_JSON(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	_ = tw.Json()
	_ = tw.JsonPtr()
	b, err := tw.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[D]("x", "2", "e", D{})
	err2 := tw2.UnmarshalJSON(b)
	if err2 != nil {
		t.Fatal("expected no error")
	}
	_, _ = tw.Serialize()
	_ = tw.SerializeMust()
	_ = tw.TypedDataJson()
	_ = tw.TypedDataJsonPtr()
	_, _ = tw.TypedDataJsonBytes()
}

func Test_CovPL_S2_35_TypedPW_GetAs(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "1", "e", "hello")
	s, ok := tw.GetAsString()
	if !ok || s != "hello" {
		t.Fatal("expected hello")
	}
	_ = tw.ValueString()

	twi, _ := corepayload.NewTypedPayloadWrapperFrom[int]("n", "1", "e", 42)
	i, ok2 := twi.GetAsInt()
	if !ok2 || i != 42 {
		t.Fatal("expected 42")
	}
	_ = twi.ValueInt()

	twb, _ := corepayload.NewTypedPayloadWrapperFrom[bool]("n", "1", "e", true)
	b, ok3 := twb.GetAsBool()
	if !ok3 || !b {
		t.Fatal("expected true")
	}
	_ = twb.ValueBool()

	// non-matching
	_, ok4 := tw.GetAsInt()
	if ok4 {
		t.Fatal("expected false")
	}
	_, ok5 := tw.GetAsInt64()
	if ok5 {
		t.Fatal("expected false")
	}
	_, ok6 := tw.GetAsFloat64()
	if ok6 {
		t.Fatal("expected false")
	}
	_, ok7 := tw.GetAsFloat32()
	if ok7 {
		t.Fatal("expected false")
	}
	_, ok8 := tw.GetAsBytes()
	if ok8 {
		t.Fatal("expected false")
	}
	_, ok9 := tw.GetAsStrings()
	if ok9 {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S2_36_TypedPW_Setters(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	tw.SetName("new")
	tw.SetIdentifier("2")
	tw.SetEntityType("new_e")
	tw.SetCategoryName("cat")
	err := tw.SetTypedData(D{A: 5})
	if err != nil {
		t.Fatal("expected no error")
	}
	if tw.TypedData().A != 5 {
		t.Fatal("expected 5")
	}
}

func Test_CovPL_S2_37_TypedPW_Clone_ToPayloadWrapper_Reparse(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	cp, err := tw.ClonePtr(true)
	if err != nil || cp == nil {
		t.Fatal("expected non-nil")
	}
	_, _ = tw.Clone(true)
	_ = tw.ToPayloadWrapper()
	_ = tw.PayloadWrapperValue()
	_ = tw.DynamicPayloads()
	_ = tw.PayloadsString()
	_ = tw.Length()
	if tw.IsNull() {
		t.Fatal("expected false")
	}
	err2 := tw.Reparse()
	if err2 != nil {
		t.Fatal("expected no error")
	}
	// nil
	var nilTW *corepayload.TypedPayloadWrapper[D]
	c, _ := nilTW.ClonePtr(true)
	if c != nil {
		t.Fatal("expected nil")
	}
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
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	if tw.CategoryName() != "" {
		t.Fatal("expected empty")
	}
	if tw.TaskTypeName() != "" {
		t.Fatal("expected empty")
	}
	if tw.HasManyRecords() {
		t.Fatal("expected false")
	}
	if !tw.HasSingleRecord() {
		t.Fatal("expected true")
	}
	if tw.Attributes() != nil {
		t.Fatal("expected nil")
	}
	_ = tw.InitializeAttributesOnNull()
}

func Test_CovPL_S2_40_TypedPW_NewFromWrapper(t *testing.T) {
	type D struct{ A int }
	pw := newTestPWForSeg2()
	tw, err := corepayload.NewTypedPayloadWrapper[D](pw)
	if err != nil || tw == nil {
		t.Fatal("expected non-nil")
	}
	// nil wrapper
	tw2, err2 := corepayload.NewTypedPayloadWrapper[D](nil)
	if err2 == nil || tw2 != nil {
		t.Fatal("expected error")
	}
}

// --- TypedPayloadCollection ---

func Test_CovPL_S2_50_TPC_Create_Length_IsEmpty_HasItems(t *testing.T) {
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](5)
	if col.Length() != 0 {
		t.Fatal("expected 0")
	}
	if !col.IsEmpty() {
		t.Fatal("expected true")
	}
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col.Add(tw)
	if col.Length() != 1 {
		t.Fatal("expected 1")
	}
	if !col.HasItems() {
		t.Fatal("expected true")
	}
	if !col.HasAnyItem() {
		t.Fatal("expected true")
	}
	if col.Count() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovPL_S2_51_TPC_Empty_From(t *testing.T) {
	type D struct{ A int }
	empty := corepayload.EmptyTypedPayloadCollection[D]()
	if empty.Length() != 0 {
		t.Fatal("expected 0")
	}
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	from := corepayload.TypedPayloadCollectionFrom[D]([]*corepayload.TypedPayloadWrapper[D]{tw})
	if from.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovPL_S2_52_TPC_FromPayloads(t *testing.T) {
	type D struct{ A int }
	pc := newTestPC()
	col := corepayload.TypedPayloadCollectionFromPayloads[D](pc)
	if col.Length() != 2 {
		t.Fatal("expected 2")
	}
	// nil
	nilCol := corepayload.TypedPayloadCollectionFromPayloads[D](nil)
	if nilCol.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S2_53_TPC_ElementAccess(t *testing.T) {
	type D struct{ A int }
	tw1, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "2", "e", D{A: 2})
	col := corepayload.NewTypedPayloadCollection[D](2)
	col.Add(tw1)
	col.Add(tw2)
	if col.First().Data().A != 1 {
		t.Fatal("expected 1")
	}
	if col.Last().Data().A != 2 {
		t.Fatal("expected 2")
	}
	if col.FirstOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	if col.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	if col.SafeAt(0) == nil {
		t.Fatal("expected non-nil")
	}
	if col.SafeAt(10) != nil {
		t.Fatal("expected nil")
	}
	// empty
	empty := corepayload.EmptyTypedPayloadCollection[D]()
	if empty.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S2_54_TPC_Add_AddLock_Adds_AddCollection_RemoveAt(t *testing.T) {
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](5)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col.AddLock(tw)
	col.Adds(tw)
	col2 := corepayload.NewTypedPayloadCollection[D](2)
	col2.Add(tw)
	col.AddCollection(col2)
	ok := col.RemoveAt(0)
	if !ok {
		t.Fatal("expected true")
	}
	ok2 := col.RemoveAt(-1)
	if ok2 {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S2_55_TPC_Iteration_Filter(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col := corepayload.NewTypedPayloadCollection[D](2)
	col.Add(tw)
	col.ForEach(func(i int, item *corepayload.TypedPayloadWrapper[D]) {})
	col.ForEachData(func(i int, data D) {})
	col.ForEachBreak(func(i int, item *corepayload.TypedPayloadWrapper[D]) bool { return false })
	fc := col.Filter(func(item *corepayload.TypedPayloadWrapper[D]) bool { return true })
	if fc.Length() != 1 {
		t.Fatal("expected 1")
	}
	fd := col.FilterByData(func(d D) bool { return d.A == 1 })
	if fd.Length() != 1 {
		t.Fatal("expected 1")
	}
	_ = col.FirstByFilter(func(item *corepayload.TypedPayloadWrapper[D]) bool { return true })
	_ = col.FirstByData(func(d D) bool { return d.A == 1 })
	_ = col.FirstByName("n")
	_ = col.FirstById("1")
	_ = col.CountFunc(func(item *corepayload.TypedPayloadWrapper[D]) bool { return true })
}

func Test_CovPL_S2_56_TPC_Skip_Take_AllData_AllNames_AllIdentifiers(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col := corepayload.NewTypedPayloadCollection[D](2)
	col.Add(tw)
	_ = col.Skip(0)
	_ = col.Take(1)
	data := col.AllData()
	if len(data) != 1 {
		t.Fatal("expected 1")
	}
	names := col.AllNames()
	if len(names) != 1 {
		t.Fatal("expected 1")
	}
	ids := col.AllIdentifiers()
	if len(ids) != 1 {
		t.Fatal("expected 1")
	}
	// empty
	empty := corepayload.EmptyTypedPayloadCollection[D]()
	if len(empty.AllData()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S2_57_TPC_ToPayloadsCollection_Clone(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col := corepayload.NewTypedPayloadCollection[D](2)
	col.Add(tw)
	pc := col.ToPayloadsCollection()
	if pc.Length() != 1 {
		t.Fatal("expected 1")
	}
	c, err := col.Clone()
	if err != nil || c.Length() != 1 {
		t.Fatal("expected 1")
	}
	_ = col.CloneMust()
	_, _ = col.ConcatNew(tw)
	// empty
	empty := corepayload.EmptyTypedPayloadCollection[D]()
	if empty.ToPayloadsCollection().Length() != 0 {
		t.Fatal("expected 0")
	}
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
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](2)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col.Add(tw)
	if col.LengthLock() != 1 {
		t.Fatal("expected 1")
	}
	if col.IsEmptyLock() {
		t.Fatal("expected false")
	}
	if col.LastIndex() != 0 {
		t.Fatal("expected 0")
	}
	if !col.HasIndex(0) {
		t.Fatal("expected true")
	}
}

func Test_CovPL_S2_60_TPC_IsValid_HasErrors_Errors_FirstError_MergedError(t *testing.T) {
	type D struct{ A int }
	col := corepayload.NewTypedPayloadCollection[D](2)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col.Add(tw)
	if !col.IsValid() {
		t.Fatal("expected true")
	}
	if col.HasErrors() {
		t.Fatal("expected false")
	}
	if len(col.Errors()) != 0 {
		t.Fatal("expected 0")
	}
	if col.FirstError() != nil {
		t.Fatal("expected nil")
	}
	if col.MergedError() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S2_61_TPC_Deserialization(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	col := corepayload.NewTypedPayloadCollection[D](1)
	col.Add(tw)
	pc := col.ToPayloadsCollection()
	// TypedPayloadCollectionDeserialize calls DeserializeToMany which expects
	// a JSON array [{},...], not {"Items":[...]} — serialize Items directly
	b, _ := corejson.Serialize.Raw(pc.Items)
	col2, err := corepayload.TypedPayloadCollectionDeserialize[D](b)
	if err != nil || col2.Length() != 1 {
		t.Fatal("expected 1")
	}
	_ = corepayload.TypedPayloadCollectionDeserializeMust[D](b)
}

func Test_CovPL_S2_62_TPC_Single_FromData(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	single := corepayload.NewTypedPayloadCollectionSingle[D](tw)
	if single.Length() != 1 {
		t.Fatal("expected 1")
	}
	nilSingle := corepayload.NewTypedPayloadCollectionSingle[D](nil)
	if nilSingle.Length() != 0 {
		t.Fatal("expected 0")
	}
	fromData, err := corepayload.NewTypedPayloadCollectionFromData[D]("n", []D{{A: 1}, {A: 2}})
	if err != nil || fromData.Length() != 2 {
		t.Fatal("expected 2")
	}
	_ = corepayload.NewTypedPayloadCollectionFromDataMust[D]("n", []D{{A: 1}})
	emptyData, err2 := corepayload.NewTypedPayloadCollectionFromData[D]("n", []D{})
	if err2 != nil || emptyData.Length() != 0 {
		t.Fatal("expected 0")
	}
}

// --- TypedPayloadWrapper Creator functions ---

func Test_CovPL_S2_65_TypedPW_Creators(t *testing.T) {
	type D struct{ A int }
	_, err := corepayload.TypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	if err != nil {
		t.Fatal("expected no error")
	}
	_, err2 := corepayload.TypedPayloadWrapperRecord[D]("n", "1", "t", "c", D{A: 1})
	if err2 != nil {
		t.Fatal("expected no error")
	}
	// TypedPayloadWrapperRecords calls SafeTypeNameOfSliceOrSingle(false, data)
	// which calls SliceFirstItemTypeName → rt.Elem() — data MUST be a slice
	_, err3 := corepayload.TypedPayloadWrapperRecords[[]D]("n", "1", "t", "c", []D{{A: 1}})
	if err3 != nil {
		t.Fatal("expected no error")
	}
	_, err4 := corepayload.TypedPayloadWrapperNameIdRecord[D]("n", "1", D{A: 1})
	if err4 != nil {
		t.Fatal("expected no error")
	}
	_, err5 := corepayload.TypedPayloadWrapperNameIdCategory[D]("n", "1", "c", D{A: 1})
	if err5 != nil {
		t.Fatal("expected no error")
	}
	_, err6 := corepayload.TypedPayloadWrapperAll[D]("n", "1", "t", "e", "c", false, D{A: 1}, nil)
	if err6 != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovPL_S2_66_TypedPW_Deserialize(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	b, _ := tw.Serialize()
	tw2, err := corepayload.TypedPayloadWrapperDeserialize[D](b)
	if err != nil || tw2 == nil {
		t.Fatal("expected non-nil")
	}
	jr := tw.JsonPtr()
	tw3, err2 := corepayload.TypedPayloadWrapperDeserializeUsingJsonResult[D](jr)
	if err2 != nil || tw3 == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S2_67_TypedPW_DeserializeToMany(t *testing.T) {
	type D struct{ A int }
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[D]("n", "1", "e", D{A: 1})
	pws := []*corepayload.PayloadWrapper{tw.ToPayloadWrapper()}
	b, _ := corejson.Serialize.Raw(pws)
	many, err := corepayload.TypedPayloadWrapperDeserializeToMany[D](b)
	if err != nil || len(many) != 1 {
		t.Fatal("expected 1")
	}
}

// --- SessionInfo ---

func Test_CovPL_S2_70_SessionInfo(t *testing.T) {
	si := corepayload.SessionInfo{Id: "5", User: &corepayload.User{Name: "u"}, SessionPath: "/p"}
	if si.IdentifierInteger() != 5 {
		t.Fatal("expected 5")
	}
	if si.IdentifierUnsignedInteger() != 5 {
		t.Fatal("expected 5")
	}
	if si.IsEmpty() {
		t.Fatal("expected false")
	}
	if !si.IsValid() {
		t.Fatal("expected true")
	}
	if si.IsUserNameEmpty() {
		t.Fatal("expected false")
	}
	if si.IsUserEmpty() {
		t.Fatal("expected false")
	}
	if !si.HasUser() {
		t.Fatal("expected true")
	}
	if !si.IsUsernameEqual("u") {
		t.Fatal("expected true")
	}
	c := si.Clone()
	_ = c
	cp := si.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}
	_ = si.Ptr()
	// empty
	empty := corepayload.SessionInfo{}
	if empty.IdentifierInteger() >= 0 {
		t.Fatal("expected invalid")
	}
	if !empty.IsEmpty() {
		t.Fatal("expected true")
	}
	// nil
	var nilSI *corepayload.SessionInfo
	if !nilSI.IsEmpty() {
		t.Fatal("expected true")
	}
	if nilSI.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

// --- AuthInfo ---

func Test_CovPL_S2_71_AuthInfo(t *testing.T) {
	ai := corepayload.AuthInfo{
		Identifier:   "5",
		ActionType:   "act",
		ResourceName: "res",
		SessionInfo:  &corepayload.SessionInfo{Id: "1"},
		UserInfo:     &corepayload.UserInfo{User: &corepayload.User{Name: "u"}},
	}
	if ai.IdentifierInteger() != 5 {
		t.Fatal("expected 5")
	}
	if ai.IdentifierUnsignedInteger() != 5 {
		t.Fatal("expected 5")
	}
	if ai.IsEmpty() {
		t.Fatal("expected false")
	}
	if !ai.HasAnyItem() {
		t.Fatal("expected true")
	}
	if !ai.IsValid() {
		t.Fatal("expected true")
	}
	if !ai.HasActionType() {
		t.Fatal("expected true")
	}
	if !ai.HasResourceName() {
		t.Fatal("expected true")
	}
	if !ai.HasUserInfo() {
		t.Fatal("expected true")
	}
	if !ai.HasSessionInfo() {
		t.Fatal("expected true")
	}
	if ai.IsActionTypeEmpty() {
		t.Fatal("expected false")
	}
	if ai.IsResourceNameEmpty() {
		t.Fatal("expected false")
	}
	_ = ai.String()
	_ = ai.PrettyJsonString()
	_ = ai.Json()
	_ = ai.JsonPtr()
	c := ai.Clone()
	_ = c
	cp := ai.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}
	_ = ai.Ptr()
	// nil
	var nilAI *corepayload.AuthInfo
	if !nilAI.IsEmpty() {
		t.Fatal("expected true")
	}
	if nilAI.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
	// empty
	empty := corepayload.AuthInfo{}
	if empty.IdentifierInteger() >= 0 {
		t.Fatal("expected invalid")
	}
}

func Test_CovPL_S2_72_AuthInfo_Setters(t *testing.T) {
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
	if r == nil {
		t.Fatal("expected non-nil")
	}
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
	pi := corepayload.PagingInfo{
		CurrentPageIndex: 1,
		TotalPages:       5,
		PerPageItems:     10,
		TotalItems:       50,
	}
	if pi.IsEmpty() {
		t.Fatal("expected false")
	}
	if !pi.HasTotalPages() {
		t.Fatal("expected true")
	}
	if !pi.HasCurrentPageIndex() {
		t.Fatal("expected true")
	}
	if !pi.HasPerPageItems() {
		t.Fatal("expected true")
	}
	if !pi.HasTotalItems() {
		t.Fatal("expected true")
	}
	if pi.IsInvalidTotalPages() {
		t.Fatal("expected false")
	}
	if pi.IsInvalidCurrentPageIndex() {
		t.Fatal("expected false")
	}
	if pi.IsInvalidPerPageItems() {
		t.Fatal("expected false")
	}
	if pi.IsInvalidTotalItems() {
		t.Fatal("expected false")
	}
	if !pi.IsEqual(&pi) {
		t.Fatal("expected true")
	}
	c := pi.Clone()
	_ = c
	cp := pi.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}
	// nil
	var nilPI *corepayload.PagingInfo
	if !nilPI.IsEmpty() {
		t.Fatal("expected true")
	}
	if nilPI.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
	if !nilPI.IsEqual(nil) {
		t.Fatal("expected true")
	}
	if nilPI.IsEqual(&pi) {
		t.Fatal("expected false")
	}
}

// --- User ---

func Test_CovPL_S2_74_User(t *testing.T) {
	u := corepayload.User{
		Identifier:   "5",
		Name:         "u",
		Type:         "admin",
		AuthToken:    "tok",
		PasswordHash: "hash",
		IsSystemUser: false,
	}
	if u.IdentifierInteger() != 5 {
		t.Fatal("expected 5")
	}
	if u.IdentifierUnsignedInteger() != 5 {
		t.Fatal("expected 5")
	}
	if !u.HasAuthToken() {
		t.Fatal("expected true")
	}
	if !u.HasPasswordHash() {
		t.Fatal("expected true")
	}
	if u.IsPasswordHashEmpty() {
		t.Fatal("expected false")
	}
	if u.IsAuthTokenEmpty() {
		t.Fatal("expected false")
	}
	if u.IsEmpty() {
		t.Fatal("expected false")
	}
	if !u.IsValidUser() {
		t.Fatal("expected true")
	}
	if u.IsNameEmpty() {
		t.Fatal("expected false")
	}
	if !u.IsNameEqual("u") {
		t.Fatal("expected true")
	}
	if !u.IsNotSystemUser() {
		t.Fatal("expected true")
	}
	if !u.IsVirtualUser() {
		t.Fatal("expected true")
	}
	if !u.HasType() {
		t.Fatal("expected true")
	}
	if u.IsTypeEmpty() {
		t.Fatal("expected false")
	}
	_ = u.String()
	_ = u.PrettyJsonString()
	_ = u.Json()
	_ = u.JsonPtr()
	_, _ = u.Serialize()
	_ = u.Deserialize([]byte(`{"Name":"x"}`))
	c := u.Clone()
	_ = c
	cp := u.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}
	_ = u.Ptr()
	// nil
	var nilU *corepayload.User
	if !nilU.IsEmpty() {
		t.Fatal("expected true")
	}
	if nilU.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
	// empty
	empty := corepayload.User{}
	if empty.IdentifierInteger() >= 0 {
		t.Fatal("expected invalid")
	}
}

// --- UserInfo ---

func Test_CovPL_S2_75_UserInfo(t *testing.T) {
	u := &corepayload.User{Name: "u"}
	su := &corepayload.User{Name: "sys", IsSystemUser: true}
	ui := corepayload.UserInfo{User: u, SystemUser: su}
	if !ui.HasUser() {
		t.Fatal("expected true")
	}
	if !ui.HasSystemUser() {
		t.Fatal("expected true")
	}
	if ui.IsEmpty() {
		t.Fatal("expected false")
	}
	if ui.IsUserEmpty() {
		t.Fatal("expected false")
	}
	if ui.IsSystemUserEmpty() {
		t.Fatal("expected false")
	}
	c := ui.Clone()
	_ = c
	cp := ui.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}
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
	if r == nil {
		t.Fatal("expected non-nil")
	}
	_ = nilUI.SetSystemUser(su)
	_ = nilUI.SetUserSystemUser(u, su)
	if !nilUI.IsEmpty() {
		t.Fatal("expected true")
	}
	if nilUI.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
	nonPtr := nilUI.ToNonPtr()
	if nonPtr.HasUser() {
		t.Fatal("expected false")
	}
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
	u := corepayload.New.User.Create(false, "u", "t")
	b, _ := u.Serialize()
	u2, err := corepayload.New.User.Deserialize(b)
	if err != nil || u2 == nil {
		t.Fatal("expected non-nil")
	}
	u3, err2 := corepayload.New.User.CastOrDeserializeFrom(u)
	if err2 != nil || u3 == nil {
		t.Fatal("expected non-nil")
	}
	// nil
	_, err3 := corepayload.New.User.CastOrDeserializeFrom(nil)
	if err3 == nil {
		t.Fatal("expected error")
	}
}

// --- PayloadCreateInstructionTypeStringer ---

func Test_CovPL_S2_80_PayloadCreateInstructionTypeStringer(t *testing.T) {
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
	if pci == nil {
		t.Fatal("expected non-nil")
	}
	_ = s
}

type stringerImpl struct{ v string }

func (s stringerImpl) String() string { return s.v }
