package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
)

// ══════════════════════════════════════════════════════════════════════════════
// corepayload Coverage — Segment 4 (Final): PayloadsCollection (all files),
//                         TypedPayloadWrapper deep, TypedPayloadCollection deep,
//                         newPayloadsCollectionCreator, newUserCreator,
//                         newTypedPayloadWrapperCreator, instruction stringers
// ══════════════════════════════════════════════════════════════════════════════

type seg4Stringer struct{ val string }

func (s seg4Stringer) String() string { return s.val }

func newPWSeg4() *corepayload.PayloadWrapper {
	pw, _ := corepayload.New.PayloadWrapper.Create(
		"seg4", "10", "taskType", "category",
		map[string]int{"a": 1},
	)
	return pw
}

// --- PayloadsCollection ---

func Test_CovPL_S4_01_PC_Add_Adds_AddsPtr(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(5)
	pw := *newPWSeg4()
	pc.Add(pw)
	pc.Adds(pw, pw)
	pc.AddsPtr(newPWSeg4(), newPWSeg4())
	if pc.Length() < 5 {
		t.Fatal("expected at least 5")
	}
}

func Test_CovPL_S4_02_PC_AddsPtrOptions(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(5)
	pw := newPWSeg4()
	pc.AddsPtrOptions(false, pw)
	pc.AddsPtrOptions(true, pw)
	pc.AddsPtrOptions(true) // empty
}

func Test_CovPL_S4_03_PC_AddsOptions(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(5)
	pw := *newPWSeg4()
	pc.AddsOptions(false, pw)
	pc.AddsOptions(true, pw)
	pc.AddsOptions(true) // empty
}

func Test_CovPL_S4_04_PC_AddsIf(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(5)
	pw := *newPWSeg4()
	pc.AddsIf(true, pw)
	pc.AddsIf(false, pw) // skipped
	if pc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovPL_S4_05_PC_InsertAt(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(3)
	pc.Add(*newPWSeg4())
	pc.Add(*newPWSeg4())
	pw2, _ := corepayload.New.PayloadWrapper.Create("inserted", "99", "t", "c", 1)
	pc.InsertAt(0, *pw2)
	if pc.First().Name != "inserted" {
		t.Fatal("expected inserted at 0")
	}
}

func Test_CovPL_S4_06_PC_ConcatNew_ConcatNewPtr(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	pw2 := *newPWSeg4()
	_ = pc.ConcatNew(pw2)
	_ = pc.ConcatNewPtr(newPWSeg4())
}

func Test_CovPL_S4_07_PC_Reverse(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(3)
	pw1, _ := corepayload.New.PayloadWrapper.Create("a", "1", "t", "c", 1)
	pw2, _ := corepayload.New.PayloadWrapper.Create("b", "2", "t", "c", 2)
	pw3, _ := corepayload.New.PayloadWrapper.Create("c", "3", "t", "c", 3)
	pc.Add(*pw1)
	pc.Add(*pw2)
	pc.Add(*pw3)
	pc.Reverse()
	if pc.First().Name != "c" {
		t.Fatal("expected c first after reverse")
	}
	// 2 items
	pc2 := corepayload.New.PayloadsCollection.UsingCap(2)
	pc2.Add(*pw1)
	pc2.Add(*pw2)
	pc2.Reverse()
	// 1 item
	pc3 := corepayload.New.PayloadsCollection.UsingCap(1)
	pc3.Add(*pw1)
	pc3.Reverse()
	// 0 items
	pc4 := corepayload.New.PayloadsCollection.UsingCap(0)
	pc4.Reverse()
}

func Test_CovPL_S4_08_PC_Clone_ClonePtr(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	_ = pc.Clone()
	_ = pc.ClonePtr()
	var nilPC *corepayload.PayloadsCollection
	if nilPC.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S4_09_PC_Clear_Dispose(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	pc.Clear()
	pc2 := corepayload.New.PayloadsCollection.UsingCap(2)
	pc2.Add(*newPWSeg4())
	pc2.Dispose()
	var nilPC *corepayload.PayloadsCollection
	nilPC.Clear()
	nilPC.Dispose()
}

// --- PayloadsCollectionGetters ---

func Test_CovPL_S4_10_PCG_Length_Count_IsEmpty_HasAnyItem_LastIndex_HasIndex(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	if pc.Length() != 1 {
		t.Fatal("expected 1")
	}
	if pc.Count() != 1 {
		t.Fatal("expected 1")
	}
	if pc.IsEmpty() {
		t.Fatal("expected false")
	}
	if !pc.HasAnyItem() {
		t.Fatal("expected true")
	}
	if pc.LastIndex() != 0 {
		t.Fatal("expected 0")
	}
	if !pc.HasIndex(0) {
		t.Fatal("expected true")
	}
	if pc.HasIndex(1) {
		t.Fatal("expected false")
	}
	var nilPC *corepayload.PayloadsCollection
	if nilPC.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S4_11_PCG_First_Last_FirstOrDefault_LastOrDefault(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
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
	empty := corepayload.New.PayloadsCollection.UsingCap(0)
	if empty.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	if empty.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
	var nilPC *corepayload.PayloadsCollection
	if nilPC.First() != nil {
		t.Fatal("expected nil")
	}
	if nilPC.Last() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S4_12_PCG_Skip_Take_Limit_Collection(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		pc.Add(*newPWSeg4())
	}
	_ = pc.Skip(2)
	_ = pc.SkipDynamic(2)
	_ = pc.SkipCollection(2)
	_ = pc.Take(3)
	_ = pc.TakeDynamic(3)
	_ = pc.TakeCollection(3)
	_ = pc.LimitCollection(3)
	_ = pc.SafeLimitCollection(3)
	_ = pc.SafeLimitCollection(100) // exceeds
	_ = pc.LimitDynamic(3)
	_ = pc.Limit(3)
}

func Test_CovPL_S4_13_PCG_Strings(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	s := pc.Strings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovPL_S4_14_PCG_IsEqual_IsEqualItems(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWSeg4())
	pc2 := corepayload.New.PayloadsCollection.UsingCap(1)
	pc2.Add(*newPWSeg4())
	if !pc.IsEqual(pc2) {
		t.Fatal("expected true")
	}
	if !pc.IsEqualItems(pc2.Items...) {
		t.Fatal("expected true")
	}
	var nilPC *corepayload.PayloadsCollection
	if !nilPC.IsEqual(nil) {
		t.Fatal("expected true")
	}
	if nilPC.IsEqual(pc) {
		t.Fatal("expected false")
	}
	// different length
	pc3 := corepayload.New.PayloadsCollection.UsingCap(2)
	pc3.Add(*newPWSeg4())
	pc3.Add(*newPWSeg4())
	if pc.IsEqual(pc3) {
		t.Fatal("expected false")
	}
}

// --- PayloadsCollectionFilter ---

func Test_CovPL_S4_20_PCF_Filter(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(3)
	pw1, _ := corepayload.New.PayloadWrapper.Create("a", "1", "t", "c", 1)
	pw2, _ := corepayload.New.PayloadWrapper.Create("b", "2", "t", "c", 2)
	pc.Add(*pw1)
	pc.Add(*pw2)
	items := pc.Filter(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return pw.Name == "a", false
	})
	if len(items) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovPL_S4_21_PCF_FilterWithLimit(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		pc.Add(*newPWSeg4())
	}
	items := pc.FilterWithLimit(2, func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})
	if len(items) > 2 {
		t.Fatal("expected at most 2")
	}
}

func Test_CovPL_S4_22_PCF_FirstByFilter(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pw1, _ := corepayload.New.PayloadWrapper.Create("a", "1", "t", "c", 1)
	pc.Add(*pw1)
	found := pc.FirstByFilter(func(pw *corepayload.PayloadWrapper) bool {
		return pw.Name == "a"
	})
	if found == nil {
		t.Fatal("expected non-nil")
	}
	notFound := pc.FirstByFilter(func(pw *corepayload.PayloadWrapper) bool {
		return pw.Name == "missing"
	})
	if notFound != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S4_23_PCF_FirstById_FirstByCategory_FirstByTaskType_FirstByEntityType(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	if pc.FirstById("10") == nil {
		t.Fatal("expected non-nil")
	}
	if pc.FirstByCategory("category") == nil {
		t.Fatal("expected non-nil")
	}
	if pc.FirstByTaskType("taskType") == nil {
		t.Fatal("expected non-nil")
	}
	if pc.FirstByEntityType(newPWSeg4().EntityType) == nil {
		t.Fatal("expected non-nil")
	}
	if pc.FirstById("missing") != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S4_24_PCF_FilterCollection_SkipFilterCollection(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(3)
	pc.Add(*newPWSeg4())
	pc.Add(*newPWSeg4())
	fc := pc.FilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})
	if fc.Length() != 2 {
		t.Fatal("expected 2")
	}
	sc := pc.SkipFilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return false, false // include all
	})
	if sc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CovPL_S4_25_PCF_FilterCollectionByIds_FilterNameCollection_FilterCategoryCollection_FilterEntityTypeCollection_FilterTaskTypeCollection(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(3)
	pc.Add(*newPWSeg4())
	_ = pc.FilterCollectionByIds("10")
	_ = pc.FilterNameCollection("seg4")
	_ = pc.FilterCategoryCollection("category")
	_ = pc.FilterEntityTypeCollection(newPWSeg4().EntityType)
	_ = pc.FilterTaskTypeCollection("taskType")
}

// --- PayloadsCollectionJson ---

func Test_CovPL_S4_30_PCJ_StringsUsingFmt_JoinUsingFmt(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	strs := pc.StringsUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.Name
	})
	if len(strs) != 1 {
		t.Fatal("expected 1")
	}
	j := pc.JoinUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.Name
	}, ",")
	if j == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovPL_S4_31_PCJ_JsonStrings_JoinJsonStrings_Join_JoinCsv_JoinCsvLine(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	_ = pc.JsonStrings()
	_ = pc.JoinJsonStrings(",")
	_ = pc.Join(",")
	_ = pc.JoinCsv()
	_ = pc.JoinCsvLine()
}

func Test_CovPL_S4_32_PCJ_JsonString_String_PrettyJsonString(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	if pc.JsonString() == "" {
		t.Fatal("expected non-empty")
	}
	if pc.String() == "" {
		t.Fatal("expected non-empty")
	}
	if pc.PrettyJsonString() == "" {
		t.Fatal("expected non-empty")
	}
	empty := corepayload.New.PayloadsCollection.UsingCap(0)
	if empty.JsonString() != "" {
		t.Fatal("expected empty")
	}
	if empty.String() != "" {
		t.Fatal("expected empty")
	}
	if empty.PrettyJsonString() != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovPL_S4_33_PCJ_CsvStrings(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	csv := pc.CsvStrings()
	if len(csv) != 1 {
		t.Fatal("expected 1")
	}
	empty := corepayload.New.PayloadsCollection.UsingCap(0)
	if len(empty.CsvStrings()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S4_34_PCJ_Json_JsonPtr_ParseInjectUsingJson_ParseInjectUsingJsonMust(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWSeg4())
	_ = pc.Json()
	jr := pc.JsonPtr()
	pc2 := corepayload.New.PayloadsCollection.UsingCap(0)
	_, _ = pc2.ParseInjectUsingJson(jr)
	pc3 := corepayload.New.PayloadsCollection.UsingCap(0)
	_ = pc3.ParseInjectUsingJsonMust(jr)
}

func Test_CovPL_S4_35_PCJ_AsInterfaces(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	_ = pc.AsJsonContractsBinder()
	_ = pc.AsJsoner()
	_ = pc.AsJsonParseSelfInjector()
	_ = pc.JsonParseSelfInject(pc.JsonPtr())
}

// --- PayloadsCollectionPaging ---

func Test_CovPL_S4_40_PCP_GetPagesSize(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		pc.Add(*newPWSeg4())
	}
	if pc.GetPagesSize(3) != 4 {
		t.Fatal("expected 4")
	}
	if pc.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S4_41_PCP_GetPagedCollection(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		pc.Add(*newPWSeg4())
	}
	pages := pc.GetPagedCollection(3)
	if len(pages) < 3 {
		t.Fatal("expected at least 3")
	}
	// small
	small := corepayload.New.PayloadsCollection.UsingCap(2)
	small.Add(*newPWSeg4())
	_ = small.GetPagedCollection(10)
}

func Test_CovPL_S4_42_PCP_GetSinglePageCollection(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		pc.Add(*newPWSeg4())
	}
	page := pc.GetSinglePageCollection(3, 2)
	if page.Length() == 0 {
		t.Fatal("expected non-empty")
	}
	// last page
	_ = pc.GetSinglePageCollection(3, 4)
	// small
	small := corepayload.New.PayloadsCollection.UsingCap(2)
	small.Add(*newPWSeg4())
	_ = small.GetSinglePageCollection(10, 1)
}

// --- newPayloadsCollectionCreator ---

func Test_CovPL_S4_50_NPCC_Empty(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	if pc.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S4_51_NPCC_Deserialize_DeserializeMust(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWSeg4())
	b, _ := corejson.Serialize.Raw(pc)
	pc2, err := corepayload.New.PayloadsCollection.Deserialize(b)
	if err != nil || pc2.Length() != 1 {
		t.Fatal("expected 1")
	}
	pc3 := corepayload.New.PayloadsCollection.DeserializeMust(b)
	if pc3.Length() != 1 {
		t.Fatal("expected 1")
	}
	// bad bytes
	_, err2 := corepayload.New.PayloadsCollection.Deserialize([]byte("bad"))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S4_52_NPCC_DeserializeToMany(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWSeg4())
	b, _ := corejson.Serialize.Raw([]*corepayload.PayloadsCollection{pc})
	many, err := corepayload.New.PayloadsCollection.DeserializeToMany(b)
	if err != nil || len(many) != 1 {
		t.Fatal("expected 1")
	}
	_, err2 := corepayload.New.PayloadsCollection.DeserializeToMany([]byte("bad"))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S4_53_NPCC_DeserializeUsingJsonResult(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWSeg4())
	jr := pc.JsonPtr()
	pc2, err := corepayload.New.PayloadsCollection.DeserializeUsingJsonResult(jr)
	if err != nil || pc2.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovPL_S4_54_NPCC_UsingWrappers(t *testing.T) {
	pw := newPWSeg4()
	pc := corepayload.New.PayloadsCollection.UsingWrappers(pw)
	if pc.Length() != 1 {
		t.Fatal("expected 1")
	}
	empty := corepayload.New.PayloadsCollection.UsingWrappers()
	if empty.Length() != 0 {
		t.Fatal("expected 0")
	}
}

// --- newUserCreator ---

func Test_CovPL_S4_60_NUC_Methods(t *testing.T) {
	_ = corepayload.New.User.Empty()
	_ = corepayload.New.User.Create(false, "name", "type")
	_ = corepayload.New.User.NonSysCreate("name", "type")
	_ = corepayload.New.User.NonSysCreateId("id", "name", "type")
	_ = corepayload.New.User.System("name", "type")
	_ = corepayload.New.User.SystemId("id", "name", "type")
	_ = corepayload.New.User.UsingNameTypeStringer("name", seg4Stringer{"type"})
	_ = corepayload.New.User.SysUsingNameTypeStringer("name", seg4Stringer{"type"})
	_ = corepayload.New.User.UsingName("name")
	_ = corepayload.New.User.All(false, "id", "name", "type", "token", "hash")
	_ = corepayload.New.User.AllTypeStringer(false, "id", "name", seg4Stringer{"type"}, "token", "hash")
	_ = corepayload.New.User.AllUsingStringer(false, "id", "name", seg4Stringer{"type"}, "token", "hash")
}

func Test_CovPL_S4_61_NUC_Deserialize(t *testing.T) {
	u := corepayload.New.User.Create(false, "name", "type")
	b, _ := corejson.Serialize.Raw(u)
	u2, err := corepayload.New.User.Deserialize(b)
	if err != nil || u2 == nil {
		t.Fatal("expected non-nil")
	}
	_, err2 := corepayload.New.User.Deserialize([]byte("bad"))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S4_62_NUC_CastOrDeserializeFrom(t *testing.T) {
	u := corepayload.New.User.Create(false, "name", "type")
	u2, err := corepayload.New.User.CastOrDeserializeFrom(u)
	if err != nil || u2 == nil {
		t.Fatal("expected non-nil")
	}
	_, err2 := corepayload.New.User.CastOrDeserializeFrom(nil)
	if err2 == nil {
		t.Fatal("expected error")
	}
}

// --- TypedPayloadWrapper deep coverage ---

type tpwTestData struct {
	Name  string
	Value int
}

func Test_CovPL_S4_70_TPW_Accessors(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	if tw.Name() != "tw" {
		t.Fatal("expected tw")
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
	if tw.EntityType() != "entity" {
		t.Fatal("expected entity")
	}
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
	if !tw.IsParsed() {
		t.Fatal("expected true")
	}
	if tw.Data().Value != 5 {
		t.Fatal("expected 5")
	}
	if tw.TypedData().Name != "x" {
		t.Fatal("expected x")
	}
}

func Test_CovPL_S4_71_TPW_NilAccessors(t *testing.T) {
	var nilTW *corepayload.TypedPayloadWrapper[tpwTestData]
	if nilTW.Name() != "" {
		t.Fatal("expected empty")
	}
	if nilTW.Identifier() != "" {
		t.Fatal("expected empty")
	}
	if nilTW.EntityType() != "" {
		t.Fatal("expected empty")
	}
	if nilTW.CategoryName() != "" {
		t.Fatal("expected empty")
	}
	if nilTW.TaskTypeName() != "" {
		t.Fatal("expected empty")
	}
	if nilTW.HasManyRecords() {
		t.Fatal("expected false")
	}
	if nilTW.Attributes() != nil {
		t.Fatal("expected nil")
	}
	if nilTW.IsParsed() {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S4_72_TPW_Error_HasError_IsEmpty_HasItems_HasSafeItems_HandleError(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
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
	tw.HandleError() // no panic
}

func Test_CovPL_S4_73_TPW_String_PrettyJsonString_JsonString(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	if tw.String() == "" {
		t.Fatal("expected non-empty")
	}
	if tw.PrettyJsonString() == "" {
		t.Fatal("expected non-empty")
	}
	if tw.JsonString() == "" {
		t.Fatal("expected non-empty")
	}
	var nilTW *corepayload.TypedPayloadWrapper[tpwTestData]
	if nilTW.String() != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovPL_S4_74_TPW_Json_JsonPtr_MarshalJSON_UnmarshalJSON(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	_ = tw.Json()
	_ = tw.JsonPtr()
	b, err := tw.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
	tw2 := &corepayload.TypedPayloadWrapper[tpwTestData]{}
	err2 := tw2.UnmarshalJSON(b)
	if err2 != nil {
		t.Fatal("expected no error")
	}
	var nilTW *corepayload.TypedPayloadWrapper[tpwTestData]
	_ = nilTW.Json()
	_ = nilTW.JsonPtr()
	_, err3 := nilTW.MarshalJSON()
	if err3 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S4_75_TPW_Serialize_SerializeMust(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	b, err := tw.Serialize()
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
	_ = tw.SerializeMust()
}

func Test_CovPL_S4_76_TPW_TypedDataJson_TypedDataJsonPtr_TypedDataJsonBytes(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	_ = tw.TypedDataJson()
	_ = tw.TypedDataJsonPtr()
	_, _ = tw.TypedDataJsonBytes()
}

func Test_CovPL_S4_77_TPW_GetAs_Methods(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string](
		"tw", "1", "entity", "hello")
	s, ok := tw.GetAsString()
	if !ok || s != "hello" {
		t.Fatal("expected hello")
	}
	twI, _ := corepayload.NewTypedPayloadWrapperFrom[int](
		"tw", "1", "entity", 42)
	v, ok := twI.GetAsInt()
	if !ok || v != 42 {
		t.Fatal("expected 42")
	}
	twB, _ := corepayload.NewTypedPayloadWrapperFrom[bool](
		"tw", "1", "entity", true)
	b, ok := twB.GetAsBool()
	if !ok || !b {
		t.Fatal("expected true")
	}
	twF, _ := corepayload.NewTypedPayloadWrapperFrom[float64](
		"tw", "1", "entity", 3.14)
	f, ok := twF.GetAsFloat64()
	if !ok || f != 3.14 {
		t.Fatal("expected 3.14")
	}
	twF32, _ := corepayload.NewTypedPayloadWrapperFrom[float32](
		"tw", "1", "entity", float32(1.5))
	f32, ok := twF32.GetAsFloat32()
	if !ok || f32 != 1.5 {
		t.Fatal("expected 1.5")
	}
	twI64, _ := corepayload.NewTypedPayloadWrapperFrom[int64](
		"tw", "1", "entity", int64(99))
	i64, ok := twI64.GetAsInt64()
	if !ok || i64 != 99 {
		t.Fatal("expected 99")
	}
	twBS, _ := corepayload.NewTypedPayloadWrapperFrom[[]byte](
		"tw", "1", "entity", []byte("abc"))
	bs, ok := twBS.GetAsBytes()
	if !ok || len(bs) != 3 {
		t.Fatal("expected 3")
	}
	twSS, _ := corepayload.NewTypedPayloadWrapperFrom[[]string](
		"tw", "1", "entity", []string{"a", "b"})
	ss, ok := twSS.GetAsStrings()
	if !ok || len(ss) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CovPL_S4_78_TPW_Value_Methods(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string](
		"tw", "1", "entity", "hello")
	if tw.ValueString() != "hello" {
		t.Fatal("expected hello")
	}
	twI, _ := corepayload.NewTypedPayloadWrapperFrom[int](
		"tw", "1", "entity", 42)
	if twI.ValueInt() != 42 {
		t.Fatal("expected 42")
	}
	twB, _ := corepayload.NewTypedPayloadWrapperFrom[bool](
		"tw", "1", "entity", true)
	if !twB.ValueBool() {
		t.Fatal("expected true")
	}
	// fallback
	twD, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x"})
	if twD.ValueString() == "" {
		t.Fatal("expected non-empty")
	}
	if twD.ValueInt() >= 0 {
		t.Fatal("expected invalid")
	}
	if twD.ValueBool() {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S4_79_TPW_Setters(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	tw.SetName("new")
	tw.SetIdentifier("2")
	tw.SetEntityType("e2")
	tw.SetCategoryName("c2")
	if tw.Name() != "new" {
		t.Fatal("expected new")
	}
	err := tw.SetTypedData(tpwTestData{Name: "y", Value: 10})
	if err != nil {
		t.Fatal("expected no error")
	}
	if tw.Data().Value != 10 {
		t.Fatal("expected 10")
	}
	tw.SetTypedDataMust(tpwTestData{Name: "z", Value: 20})
}

func Test_CovPL_S4_80_TPW_Clone_ClonePtr_ToPayloadWrapper_Reparse(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	_, _ = tw.Clone(true)
	_, _ = tw.Clone(false)
	_, _ = tw.ClonePtr(true)
	_, _ = tw.ClonePtr(false)
	_ = tw.ToPayloadWrapper()
	_ = tw.PayloadWrapperValue()
	err := tw.Reparse()
	if err != nil {
		t.Fatal("expected no error")
	}
	var nilTW *corepayload.TypedPayloadWrapper[tpwTestData]
	cp, _ := nilTW.ClonePtr(true)
	if cp != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S4_81_TPW_DynamicPayloads_PayloadsString_Length_IsNull(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	if len(tw.DynamicPayloads()) == 0 {
		t.Fatal("expected non-empty")
	}
	if tw.PayloadsString() == "" {
		t.Fatal("expected non-empty")
	}
	if tw.Length() == 0 {
		t.Fatal("expected > 0")
	}
	if tw.IsNull() {
		t.Fatal("expected false")
	}
	var nilTW *corepayload.TypedPayloadWrapper[tpwTestData]
	if !nilTW.IsNull() {
		t.Fatal("expected true")
	}
	if nilTW.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S4_82_TPW_Clear_Dispose(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	tw.Clear()
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	tw2.Dispose()
	var nilTW *corepayload.TypedPayloadWrapper[tpwTestData]
	nilTW.Clear()
	nilTW.Dispose()
}

func Test_CovPL_S4_83_TPW_NewTypedPayloadWrapperMust(t *testing.T) {
	pw := newPWSeg4()
	_ = corepayload.NewTypedPayloadWrapperMust[map[string]int](pw)
}

// --- newTypedPayloadWrapperCreator ---

func Test_CovPL_S4_85_NTPWC_TypedPayloadWrapperFrom(t *testing.T) {
	tw, err := corepayload.TypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	if err != nil || tw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S4_86_NTPWC_TypedPayloadWrapperRecord(t *testing.T) {
	tw, err := corepayload.TypedPayloadWrapperRecord[tpwTestData](
		"tw", "1", "task", "cat", tpwTestData{Name: "x", Value: 5})
	if err != nil || tw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S4_87_NTPWC_TypedPayloadWrapperRecords(t *testing.T) {
	tw, err := corepayload.TypedPayloadWrapperRecords[[]tpwTestData](
		"tw", "1", "task", "cat", []tpwTestData{{Name: "a"}, {Name: "b"}})
	if err != nil || tw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S4_88_NTPWC_TypedPayloadWrapperNameIdRecord(t *testing.T) {
	tw, err := corepayload.TypedPayloadWrapperNameIdRecord[tpwTestData](
		"tw", "1", tpwTestData{Name: "x", Value: 5})
	if err != nil || tw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S4_89_NTPWC_TypedPayloadWrapperNameIdCategory(t *testing.T) {
	tw, err := corepayload.TypedPayloadWrapperNameIdCategory[tpwTestData](
		"tw", "1", "cat", tpwTestData{Name: "x", Value: 5})
	if err != nil || tw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S4_90_NTPWC_TypedPayloadWrapperAll(t *testing.T) {
	tw, err := corepayload.TypedPayloadWrapperAll[tpwTestData](
		"tw", "1", "task", "entity", "cat", false,
		tpwTestData{Name: "x", Value: 5}, nil)
	if err != nil || tw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S4_91_NTPWC_TypedPayloadWrapperDeserialize(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	b, _ := tw.Serialize()
	tw2, err := corepayload.TypedPayloadWrapperDeserialize[tpwTestData](b)
	if err != nil || tw2 == nil {
		t.Fatal("expected non-nil")
	}
	_, err2 := corepayload.TypedPayloadWrapperDeserialize[tpwTestData]([]byte("bad"))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S4_92_NTPWC_TypedPayloadWrapperDeserializeUsingJsonResult(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	jr := tw.JsonPtr()
	tw2, err := corepayload.TypedPayloadWrapperDeserializeUsingJsonResult[tpwTestData](jr)
	if err != nil || tw2 == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovPL_S4_93_NTPWC_TypedPayloadWrapperDeserializeToMany(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	b, _ := corejson.Serialize.Raw([]*corepayload.PayloadWrapper{tw.Wrapper})
	many, err := corepayload.TypedPayloadWrapperDeserializeToMany[tpwTestData](b)
	if err != nil || len(many) != 1 {
		t.Fatal("expected 1")
	}
	_, err2 := corepayload.TypedPayloadWrapperDeserializeToMany[tpwTestData]([]byte("bad"))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

// --- TypedPayloadCollection deep coverage ---

func Test_CovPL_S4_100_TPC_Core(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](5)
	if !col.IsEmpty() {
		t.Fatal("expected true")
	}
	if col.HasItems() {
		t.Fatal("expected false")
	}
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	col.Add(tw)
	if col.Length() != 1 {
		t.Fatal("expected 1")
	}
	if col.Count() != 1 {
		t.Fatal("expected 1")
	}
	if col.IsEmpty() {
		t.Fatal("expected false")
	}
	if !col.HasItems() {
		t.Fatal("expected true")
	}
	if !col.HasAnyItem() {
		t.Fatal("expected true")
	}
	if col.LastIndex() != 0 {
		t.Fatal("expected 0")
	}
	if !col.HasIndex(0) {
		t.Fatal("expected true")
	}
	if col.HasIndex(1) {
		t.Fatal("expected false")
	}
	_ = col.Items()
}

func Test_CovPL_S4_101_TPC_NilAccessors(t *testing.T) {
	var nilCol *corepayload.TypedPayloadCollection[tpwTestData]
	if nilCol.Length() != 0 {
		t.Fatal("expected 0")
	}
	if !nilCol.IsEmpty() {
		t.Fatal("expected true")
	}
	if nilCol.HasItems() {
		t.Fatal("expected false")
	}
	if nilCol.Items() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S4_102_TPC_LengthLock_IsEmptyLock(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](1)
	_ = col.LengthLock()
	_ = col.IsEmptyLock()
}

func Test_CovPL_S4_103_TPC_ElementAccess(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](2)
	tw1, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"b", "2", "e", tpwTestData{Name: "b"})
	col.Add(tw1)
	col.Add(tw2)
	if col.First().Name() != "a" {
		t.Fatal("expected a")
	}
	if col.Last().Name() != "b" {
		t.Fatal("expected b")
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
	empty := corepayload.EmptyTypedPayloadCollection[tpwTestData]()
	if empty.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	if empty.SafeAt(0) != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovPL_S4_104_TPC_Mutation(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](5)
	tw1, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"b", "2", "e", tpwTestData{Name: "b"})
	col.Add(tw1)
	col.AddLock(tw2)
	col.Adds(tw1, tw2)
	col2 := corepayload.NewTypedPayloadCollection[tpwTestData](1)
	col2.Add(tw1)
	col.AddCollection(col2)
	col.AddCollection(corepayload.EmptyTypedPayloadCollection[tpwTestData]())
	if !col.RemoveAt(0) {
		t.Fatal("expected true")
	}
	if col.RemoveAt(-1) {
		t.Fatal("expected false")
	}
	if col.RemoveAt(100) {
		t.Fatal("expected false")
	}
}

func Test_CovPL_S4_105_TPC_Iteration(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](2)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a", Value: 1})
	col.Add(tw)
	count := 0
	col.ForEach(func(i int, item *corepayload.TypedPayloadWrapper[tpwTestData]) {
		count++
	})
	if count != 1 {
		t.Fatal("expected 1")
	}
	col.ForEachData(func(i int, d tpwTestData) {
		if d.Name != "a" {
			t.Fatal("expected a")
		}
	})
	col.ForEachBreak(func(i int, item *corepayload.TypedPayloadWrapper[tpwTestData]) bool {
		return true // break immediately
	})
}

func Test_CovPL_S4_106_TPC_Filter(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](3)
	tw1, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a", Value: 1})
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"b", "2", "e", tpwTestData{Name: "b", Value: 2})
	col.Add(tw1)
	col.Add(tw2)
	filtered := col.Filter(func(item *corepayload.TypedPayloadWrapper[tpwTestData]) bool {
		return item.Data().Value == 1
	})
	if filtered.Length() != 1 {
		t.Fatal("expected 1")
	}
	filteredByData := col.FilterByData(func(d tpwTestData) bool {
		return d.Name == "b"
	})
	if filteredByData.Length() != 1 {
		t.Fatal("expected 1")
	}
	found := col.FirstByFilter(func(item *corepayload.TypedPayloadWrapper[tpwTestData]) bool {
		return item.Data().Name == "a"
	})
	if found == nil {
		t.Fatal("expected non-nil")
	}
	foundByData := col.FirstByData(func(d tpwTestData) bool {
		return d.Name == "b"
	})
	if foundByData == nil {
		t.Fatal("expected non-nil")
	}
	if col.FirstByName("a") == nil {
		t.Fatal("expected non-nil")
	}
	if col.FirstById("1") == nil {
		t.Fatal("expected non-nil")
	}
	if col.CountFunc(func(item *corepayload.TypedPayloadWrapper[tpwTestData]) bool { return true }) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CovPL_S4_107_TPC_SkipTake(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](5)
	for i := 0; i < 5; i++ {
		tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
			"x", "1", "e", tpwTestData{Value: i})
		col.Add(tw)
	}
	if len(col.Skip(2)) != 3 {
		t.Fatal("expected 3")
	}
	if len(col.Skip(10)) != 0 {
		t.Fatal("expected 0")
	}
	if len(col.Take(3)) != 3 {
		t.Fatal("expected 3")
	}
	if len(col.Take(10)) != 5 {
		t.Fatal("expected 5")
	}
}

func Test_CovPL_S4_108_TPC_Extraction(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](2)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a", Value: 1})
	col.Add(tw)
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
	empty := corepayload.EmptyTypedPayloadCollection[tpwTestData]()
	if len(empty.AllData()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.AllNames()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.AllIdentifiers()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S4_109_TPC_ToPayloadsCollection(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
	col.Add(tw)
	pc := col.ToPayloadsCollection()
	if pc.Length() != 1 {
		t.Fatal("expected 1")
	}
	empty := corepayload.EmptyTypedPayloadCollection[tpwTestData]()
	epc := empty.ToPayloadsCollection()
	if epc.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S4_110_TPC_Clone_CloneMust_ConcatNew(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
	col.Add(tw)
	cloned, err := col.Clone()
	if err != nil || cloned.Length() != 1 {
		t.Fatal("expected 1")
	}
	_ = col.CloneMust()
	concat, err := col.ConcatNew(tw)
	if err != nil || concat.Length() != 2 {
		t.Fatal("expected 2")
	}
	empty := corepayload.EmptyTypedPayloadCollection[tpwTestData]()
	ec, _ := empty.Clone()
	if ec.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S4_111_TPC_Clear_Dispose(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
	col.Add(tw)
	col.Clear()
	col2 := corepayload.NewTypedPayloadCollection[tpwTestData](1)
	col2.Add(tw)
	col2.Dispose()
	var nilCol *corepayload.TypedPayloadCollection[tpwTestData]
	nilCol.Clear()
	nilCol.Dispose()
}

func Test_CovPL_S4_112_TPC_Deserialization(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
	b, _ := corejson.Serialize.Raw([]*corepayload.PayloadWrapper{tw.Wrapper})
	col, err := corepayload.TypedPayloadCollectionDeserialize[tpwTestData](b)
	if err != nil || col.Length() != 1 {
		t.Fatal("expected 1")
	}
	_ = corepayload.TypedPayloadCollectionDeserializeMust[tpwTestData](b)
	_, err2 := corepayload.TypedPayloadCollectionDeserialize[tpwTestData]([]byte("bad"))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovPL_S4_113_TPC_FromPayloads(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWSeg4())
	col := corepayload.TypedPayloadCollectionFromPayloads[map[string]int](pc)
	if col.Length() != 1 {
		t.Fatal("expected 1")
	}
	// nil
	nilCol := corepayload.TypedPayloadCollectionFromPayloads[map[string]int](nil)
	if nilCol.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S4_114_TPC_NewTypedPayloadCollectionSingle(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
	col := corepayload.NewTypedPayloadCollectionSingle[tpwTestData](tw)
	if col.Length() != 1 {
		t.Fatal("expected 1")
	}
	nilCol := corepayload.NewTypedPayloadCollectionSingle[tpwTestData](nil)
	if nilCol.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S4_115_TPC_NewTypedPayloadCollectionFromData(t *testing.T) {
	col, err := corepayload.NewTypedPayloadCollectionFromData[tpwTestData](
		"test", []tpwTestData{{Name: "a"}, {Name: "b"}})
	if err != nil || col.Length() != 2 {
		t.Fatal("expected 2")
	}
	_ = corepayload.NewTypedPayloadCollectionFromDataMust[tpwTestData](
		"test", []tpwTestData{{Name: "a"}})
	// empty
	empty, _ := corepayload.NewTypedPayloadCollectionFromData[tpwTestData]("test", nil)
	if empty.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovPL_S4_116_TPC_IsValid_HasErrors_Errors_FirstError_MergedError(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
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
	empty := corepayload.EmptyTypedPayloadCollection[tpwTestData]()
	if !empty.IsValid() {
		t.Fatal("expected true")
	}
	if empty.Errors() != nil {
		t.Fatal("expected nil")
	}
}

// --- PayloadCreateInstructionTypeStringer ---

func Test_CovPL_S4_120_PCITS_PayloadCreateInstruction(t *testing.T) {
	inst := corepayload.PayloadCreateInstructionTypeStringer{
		Name:                 "n",
		Identifier:           "1",
		TaskTypeNameStringer: seg4Stringer{"task"},
		CategoryNameStringer: seg4Stringer{"cat"},
		HasManyRecords:       false,
		Payloads:             map[string]int{"a": 1},
	}
	pi := inst.PayloadCreateInstruction()
	if pi.TaskTypeName != "task" {
		t.Fatal("expected task")
	}
	if pi.CategoryName != "cat" {
		t.Fatal("expected cat")
	}
}

// --- BytesCreateInstructionStringer ---

func Test_CovPL_S4_121_BCIS_Fields(t *testing.T) {
	inst := corepayload.BytesCreateInstructionStringer{
		Name:           "n",
		Identifier:     "1",
		TaskTypeName:   seg4Stringer{"task"},
		EntityType:     "e",
		CategoryName:   seg4Stringer{"cat"},
		HasManyRecords: false,
		Payloads:       []byte("x"),
	}
	if inst.Name != "n" {
		t.Fatal("expected n")
	}
}
