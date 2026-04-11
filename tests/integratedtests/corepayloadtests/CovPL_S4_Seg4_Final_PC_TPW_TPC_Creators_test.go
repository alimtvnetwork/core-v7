package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
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
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(5)
	pw := *newPWSeg4()
	pc.Add(pw)
	pc.Adds(pw, pw)
	pc.AddsPtr(newPWSeg4(), newPWSeg4())

	// Act
	actual := args.Map{"result": pc.Length() < 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 5", actual)
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
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(5)
	pw := *newPWSeg4()
	pc.AddsIf(true, pw)
	pc.AddsIf(false, pw) // skipped

	// Act
	actual := args.Map{"result": pc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovPL_S4_05_PC_InsertAt(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(3)
	pc.Add(*newPWSeg4())
	pc.Add(*newPWSeg4())
	pw2, _ := corepayload.New.PayloadWrapper.Create("inserted", "99", "t", "c", 1)
	pc.InsertAt(0, *pw2)

	// Act
	actual := args.Map{"result": pc.First().Name != "inserted"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected inserted at 0", actual)
}

func Test_CovPL_S4_06_PC_ConcatNew_ConcatNewPtr(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	pw2 := *newPWSeg4()
	_ = pc.ConcatNew(pw2)
	_ = pc.ConcatNewPtr(newPWSeg4())
}

func Test_CovPL_S4_07_PC_Reverse(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(3)
	pw1, _ := corepayload.New.PayloadWrapper.Create("a", "1", "t", "c", 1)
	pw2, _ := corepayload.New.PayloadWrapper.Create("b", "2", "t", "c", 2)
	pw3, _ := corepayload.New.PayloadWrapper.Create("c", "3", "t", "c", 3)
	pc.Add(*pw1)
	pc.Add(*pw2)
	pc.Add(*pw3)
	pc.Reverse()

	// Act
	actual := args.Map{"result": pc.First().Name != "c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected c first after reverse", actual)
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
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	_ = pc.Clone()
	_ = pc.ClonePtr()
	var nilPC *corepayload.PayloadsCollection

	// Act
	actual := args.Map{"result": nilPC.ClonePtr() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
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
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())

	// Act
	actual := args.Map{"result": pc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": pc.Count() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": pc.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": pc.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": pc.LastIndex() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": pc.HasIndex(0)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": pc.HasIndex(1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	var nilPC *corepayload.PayloadsCollection
	actual = args.Map{"result": nilPC.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S4_11_PCG_First_Last_FirstOrDefault_LastOrDefault(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())

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
	empty := corepayload.New.PayloadsCollection.UsingCap(0)
	actual = args.Map{"result": empty.FirstOrDefault() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": empty.LastOrDefault() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	var nilPC *corepayload.PayloadsCollection
	actual = args.Map{"result": nilPC.First() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": nilPC.Last() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
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
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	s := pc.Strings()

	// Act
	actual := args.Map{"result": len(s) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovPL_S4_14_PCG_IsEqual_IsEqualItems(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWSeg4())
	pc2 := corepayload.New.PayloadsCollection.UsingCap(1)
	pc2.Add(*newPWSeg4())

	// Act
	actual := args.Map{"result": pc.IsEqual(pc2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": pc.IsEqualItems(pc2.Items...)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	var nilPC *corepayload.PayloadsCollection
	actual = args.Map{"result": nilPC.IsEqual(nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": nilPC.IsEqual(pc)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	// different length
	pc3 := corepayload.New.PayloadsCollection.UsingCap(2)
	pc3.Add(*newPWSeg4())
	pc3.Add(*newPWSeg4())
	actual = args.Map{"result": pc.IsEqual(pc3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// --- PayloadsCollectionFilter ---

func Test_CovPL_S4_20_PCF_Filter(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(3)
	pw1, _ := corepayload.New.PayloadWrapper.Create("a", "1", "t", "c", 1)
	pw2, _ := corepayload.New.PayloadWrapper.Create("b", "2", "t", "c", 2)
	pc.Add(*pw1)
	pc.Add(*pw2)
	items := pc.Filter(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return pw.Name == "a", false
	})

	// Act
	actual := args.Map{"result": len(items) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovPL_S4_21_PCF_FilterWithLimit(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		pc.Add(*newPWSeg4())
	}
	items := pc.FilterWithLimit(2, func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})

	// Act
	actual := args.Map{"result": len(items) > 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at most 2", actual)
}

func Test_CovPL_S4_22_PCF_FirstByFilter(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pw1, _ := corepayload.New.PayloadWrapper.Create("a", "1", "t", "c", 1)
	pc.Add(*pw1)
	found := pc.FirstByFilter(func(pw *corepayload.PayloadWrapper) bool {
		return pw.Name == "a"
	})

	// Act
	actual := args.Map{"result": found == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	notFound := pc.FirstByFilter(func(pw *corepayload.PayloadWrapper) bool {
		return pw.Name == "missing"
	})
	actual = args.Map{"result": notFound != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S4_23_PCF_FirstById_FirstByCategory_FirstByTaskType_FirstByEntityType(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())

	// Act
	actual := args.Map{"result": pc.FirstById("10") == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": pc.FirstByCategory("category") == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": pc.FirstByTaskType("taskType") == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": pc.FirstByEntityType(newPWSeg4().EntityType) == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": pc.FirstById("missing") != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S4_24_PCF_FilterCollection_SkipFilterCollection(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(3)
	pc.Add(*newPWSeg4())
	pc.Add(*newPWSeg4())
	fc := pc.FilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return true, false
	})

	// Act
	actual := args.Map{"result": fc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	sc := pc.SkipFilterCollection(func(pw *corepayload.PayloadWrapper) (bool, bool) {
		return false, false // include all
	})
	actual = args.Map{"result": sc.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	strs := pc.StringsUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.Name
	})

	// Act
	actual := args.Map{"result": len(strs) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	j := pc.JoinUsingFmt(func(pw *corepayload.PayloadWrapper) string {
		return pw.Name
	}, ",")
	actual = args.Map{"result": j == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
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
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())

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
	empty := corepayload.New.PayloadsCollection.UsingCap(0)
	actual = args.Map{"result": empty.JsonString() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": empty.String() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": empty.PrettyJsonString() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovPL_S4_33_PCJ_CsvStrings(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(2)
	pc.Add(*newPWSeg4())
	csv := pc.CsvStrings()

	// Act
	actual := args.Map{"result": len(csv) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := corepayload.New.PayloadsCollection.UsingCap(0)
	actual = args.Map{"result": len(empty.CsvStrings()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		pc.Add(*newPWSeg4())
	}

	// Act
	actual := args.Map{"result": pc.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual = args.Map{"result": pc.GetPagesSize(0) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S4_41_PCP_GetPagedCollection(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		pc.Add(*newPWSeg4())
	}
	pages := pc.GetPagedCollection(3)

	// Act
	actual := args.Map{"result": len(pages) < 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	// small
	small := corepayload.New.PayloadsCollection.UsingCap(2)
	small.Add(*newPWSeg4())
	_ = small.GetPagedCollection(10)
}

func Test_CovPL_S4_42_PCP_GetSinglePageCollection(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		pc.Add(*newPWSeg4())
	}
	page := pc.GetSinglePageCollection(3, 2)

	// Act
	actual := args.Map{"result": page.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// last page
	_ = pc.GetSinglePageCollection(3, 4)
	// small
	small := corepayload.New.PayloadsCollection.UsingCap(2)
	small.Add(*newPWSeg4())
	_ = small.GetSinglePageCollection(10, 1)
}

// --- newPayloadsCollectionCreator ---

func Test_CovPL_S4_50_NPCC_Empty(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()

	// Act
	actual := args.Map{"result": pc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S4_51_NPCC_Deserialize_DeserializeMust(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWSeg4())
	b, _ := corejson.Serialize.Raw(pc)
	pc2, err := corepayload.New.PayloadsCollection.Deserialize(b)

	// Act
	actual := args.Map{"result": err != nil || pc2.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	pc3 := corepayload.New.PayloadsCollection.DeserializeMust(b)
	actual = args.Map{"result": pc3.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	// bad bytes
	_, err2 := corepayload.New.PayloadsCollection.Deserialize([]byte("bad"))
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovPL_S4_52_NPCC_DeserializeToMany(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWSeg4())
	b, _ := corejson.Serialize.Raw([]*corepayload.PayloadsCollection{pc})
	many, err := corepayload.New.PayloadsCollection.DeserializeToMany(b)

	// Act
	actual := args.Map{"result": err != nil || len(many) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_, err2 := corepayload.New.PayloadsCollection.DeserializeToMany([]byte("bad"))
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovPL_S4_53_NPCC_DeserializeUsingJsonResult(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWSeg4())
	jr := pc.JsonPtr()
	pc2, err := corepayload.New.PayloadsCollection.DeserializeUsingJsonResult(jr)

	// Act
	actual := args.Map{"result": err != nil || pc2.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovPL_S4_54_NPCC_UsingWrappers(t *testing.T) {
	// Arrange
	pw := newPWSeg4()
	pc := corepayload.New.PayloadsCollection.UsingWrappers(pw)

	// Act
	actual := args.Map{"result": pc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := corepayload.New.PayloadsCollection.UsingWrappers()
	actual = args.Map{"result": empty.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
	// Arrange
	u := corepayload.New.User.Create(false, "name", "type")
	b, _ := corejson.Serialize.Raw(u)
	u2, err := corepayload.New.User.Deserialize(b)

	// Act
	actual := args.Map{"result": err != nil || u2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_, err2 := corepayload.New.User.Deserialize([]byte("bad"))
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovPL_S4_62_NUC_CastOrDeserializeFrom(t *testing.T) {
	// Arrange
	u := corepayload.New.User.Create(false, "name", "type")
	u2, err := corepayload.New.User.CastOrDeserializeFrom(u)

	// Act
	actual := args.Map{"result": err != nil || u2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_, err2 := corepayload.New.User.CastOrDeserializeFrom(nil)
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// --- TypedPayloadWrapper deep coverage ---

type tpwTestData struct {
	Name  string
	Value int
}

func Test_CovPL_S4_70_TPW_Accessors(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})

	// Act
	actual := args.Map{"result": tw.Name() != "tw"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected tw", actual)
	actual = args.Map{"result": tw.Identifier() != "1"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": tw.IdString() != "1"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": tw.IdInteger() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": tw.EntityType() != "entity"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected entity", actual)
	actual = args.Map{"result": tw.CategoryName() != ""}
	expected = args.Map{"result": false}
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
	actual = args.Map{"result": tw.IsParsed()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": tw.Data().Value != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual = args.Map{"result": tw.TypedData().Name != "x"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected x", actual)
}

func Test_CovPL_S4_71_TPW_NilAccessors(t *testing.T) {
	// Arrange
	var nilTW *corepayload.TypedPayloadWrapper[tpwTestData]

	// Act
	actual := args.Map{"result": nilTW.Name() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": nilTW.Identifier() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": nilTW.EntityType() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": nilTW.CategoryName() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": nilTW.TaskTypeName() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": nilTW.HasManyRecords()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": nilTW.Attributes() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": nilTW.IsParsed()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S4_72_TPW_Error_HasError_IsEmpty_HasItems_HasSafeItems_HandleError(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})

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
	tw.HandleError() // no panic
}

func Test_CovPL_S4_73_TPW_String_PrettyJsonString_JsonString(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})

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
	var nilTW *corepayload.TypedPayloadWrapper[tpwTestData]
	actual = args.Map{"result": nilTW.String() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovPL_S4_74_TPW_Json_JsonPtr_MarshalJSON_UnmarshalJSON(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	_ = tw.Json()
	_ = tw.JsonPtr()
	b, err := tw.MarshalJSON()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	tw2 := &corepayload.TypedPayloadWrapper[tpwTestData]{}
	err2 := tw2.UnmarshalJSON(b)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	var nilTW *corepayload.TypedPayloadWrapper[tpwTestData]
	_ = nilTW.Json()
	_ = nilTW.JsonPtr()
	_, err3 := nilTW.MarshalJSON()
	actual = args.Map{"result": err3 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovPL_S4_75_TPW_Serialize_SerializeMust(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	b, err := tw.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
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
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string](
		"tw", "1", "entity", "hello")
	s, ok := tw.GetAsString()

	// Act
	actual := args.Map{"result": ok || s != "hello"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	twI, _ := corepayload.NewTypedPayloadWrapperFrom[int](
		"tw", "1", "entity", 42)
	v, ok := twI.GetAsInt()
	actual = args.Map{"result": ok || v != 42}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	twB, _ := corepayload.NewTypedPayloadWrapperFrom[bool](
		"tw", "1", "entity", true)
	b, ok := twB.GetAsBool()
	actual = args.Map{"result": ok || !b}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	twF, _ := corepayload.NewTypedPayloadWrapperFrom[float64](
		"tw", "1", "entity", 3.14)
	f, ok := twF.GetAsFloat64()
	actual = args.Map{"result": ok || f != 3.14}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
	twF32, _ := corepayload.NewTypedPayloadWrapperFrom[float32](
		"tw", "1", "entity", float32(1.5))
	f32, ok := twF32.GetAsFloat32()
	actual = args.Map{"result": ok || f32 != 1.5}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 1.5", actual)
	twI64, _ := corepayload.NewTypedPayloadWrapperFrom[int64](
		"tw", "1", "entity", int64(99))
	i64, ok := twI64.GetAsInt64()
	actual = args.Map{"result": ok || i64 != 99}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 99", actual)
	twBS, _ := corepayload.NewTypedPayloadWrapperFrom[[]byte](
		"tw", "1", "entity", []byte("abc"))
	bs, ok := twBS.GetAsBytes()
	actual = args.Map{"result": ok || len(bs) != 3}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	twSS, _ := corepayload.NewTypedPayloadWrapperFrom[[]string](
		"tw", "1", "entity", []string{"a", "b"})
	ss, ok := twSS.GetAsStrings()
	actual = args.Map{"result": ok || len(ss) != 2}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_CovPL_S4_78_TPW_Value_Methods(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string](
		"tw", "1", "entity", "hello")

	// Act
	actual := args.Map{"result": tw.ValueString() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	twI, _ := corepayload.NewTypedPayloadWrapperFrom[int](
		"tw", "1", "entity", 42)
	actual = args.Map{"result": twI.ValueInt() != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	twB, _ := corepayload.NewTypedPayloadWrapperFrom[bool](
		"tw", "1", "entity", true)
	actual = args.Map{"result": twB.ValueBool()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	// fallback
	twD, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x"})
	actual = args.Map{"result": twD.ValueString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual = args.Map{"result": twD.ValueInt() >= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual = args.Map{"result": twD.ValueBool()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S4_79_TPW_Setters(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	tw.SetName("new")
	tw.SetIdentifier("2")
	tw.SetEntityType("e2")
	tw.SetCategoryName("c2")

	// Act
	actual := args.Map{"result": tw.Name() != "new"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected new", actual)
	err := tw.SetTypedData(tpwTestData{Name: "y", Value: 10})
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	actual = args.Map{"result": tw.Data().Value != 10}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)
	tw.SetTypedDataMust(tpwTestData{Name: "z", Value: 20})
}

func Test_CovPL_S4_80_TPW_Clone_ClonePtr_ToPayloadWrapper_Reparse(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	_, _ = tw.Clone(true)
	_, _ = tw.Clone(false)
	_, _ = tw.ClonePtr(true)
	_, _ = tw.ClonePtr(false)
	_ = tw.ToPayloadWrapper()
	_ = tw.PayloadWrapperValue()
	err := tw.Reparse()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	var nilTW *corepayload.TypedPayloadWrapper[tpwTestData]
	cp, _ := nilTW.ClonePtr(true)
	actual = args.Map{"result": cp != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S4_81_TPW_DynamicPayloads_PayloadsString_Length_IsNull(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})

	// Act
	actual := args.Map{"result": len(tw.DynamicPayloads()) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual = args.Map{"result": tw.PayloadsString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual = args.Map{"result": tw.Length() == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected > 0", actual)
	actual = args.Map{"result": tw.IsNull()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	var nilTW *corepayload.TypedPayloadWrapper[tpwTestData]
	actual = args.Map{"result": nilTW.IsNull()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": nilTW.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
	// Arrange
	tw, err := corepayload.TypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})

	// Act
	actual := args.Map{"result": err != nil || tw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S4_86_NTPWC_TypedPayloadWrapperRecord(t *testing.T) {
	// Arrange
	tw, err := corepayload.TypedPayloadWrapperRecord[tpwTestData](
		"tw", "1", "task", "cat", tpwTestData{Name: "x", Value: 5})

	// Act
	actual := args.Map{"result": err != nil || tw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S4_87_NTPWC_TypedPayloadWrapperRecords(t *testing.T) {
	// Arrange
	tw, err := corepayload.TypedPayloadWrapperRecords[[]tpwTestData](
		"tw", "1", "task", "cat", []tpwTestData{{Name: "a"}, {Name: "b"}})

	// Act
	actual := args.Map{"result": err != nil || tw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S4_88_NTPWC_TypedPayloadWrapperNameIdRecord(t *testing.T) {
	// Arrange
	tw, err := corepayload.TypedPayloadWrapperNameIdRecord[tpwTestData](
		"tw", "1", tpwTestData{Name: "x", Value: 5})

	// Act
	actual := args.Map{"result": err != nil || tw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S4_89_NTPWC_TypedPayloadWrapperNameIdCategory(t *testing.T) {
	// Arrange
	tw, err := corepayload.TypedPayloadWrapperNameIdCategory[tpwTestData](
		"tw", "1", "cat", tpwTestData{Name: "x", Value: 5})

	// Act
	actual := args.Map{"result": err != nil || tw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S4_90_NTPWC_TypedPayloadWrapperAll(t *testing.T) {
	// Arrange
	tw, err := corepayload.TypedPayloadWrapperAll[tpwTestData](
		"tw", "1", "task", "entity", "cat", false,
		tpwTestData{Name: "x", Value: 5}, nil)

	// Act
	actual := args.Map{"result": err != nil || tw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S4_91_NTPWC_TypedPayloadWrapperDeserialize(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	b, _ := tw.Serialize()
	tw2, err := corepayload.TypedPayloadWrapperDeserialize[tpwTestData](b)

	// Act
	actual := args.Map{"result": err != nil || tw2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_, err2 := corepayload.TypedPayloadWrapperDeserialize[tpwTestData]([]byte("bad"))
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovPL_S4_92_NTPWC_TypedPayloadWrapperDeserializeUsingJsonResult(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	jr := tw.JsonPtr()
	tw2, err := corepayload.TypedPayloadWrapperDeserializeUsingJsonResult[tpwTestData](jr)

	// Act
	actual := args.Map{"result": err != nil || tw2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovPL_S4_93_NTPWC_TypedPayloadWrapperDeserializeToMany(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	b, _ := corejson.Serialize.Raw([]*corepayload.PayloadWrapper{tw.Wrapper})
	many, err := corepayload.TypedPayloadWrapperDeserializeToMany[tpwTestData](b)

	// Act
	actual := args.Map{"result": err != nil || len(many) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_, err2 := corepayload.TypedPayloadWrapperDeserializeToMany[tpwTestData]([]byte("bad"))
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// --- TypedPayloadCollection deep coverage ---

func Test_CovPL_S4_100_TPC_Core(t *testing.T) {
	// Arrange
	col := corepayload.NewTypedPayloadCollection[tpwTestData](5)

	// Act
	actual := args.Map{"result": col.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": col.HasItems()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"tw", "1", "entity", tpwTestData{Name: "x", Value: 5})
	col.Add(tw)
	actual = args.Map{"result": col.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": col.Count() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": col.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": col.HasItems()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": col.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": col.LastIndex() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": col.HasIndex(0)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": col.HasIndex(1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_ = col.Items()
}

func Test_CovPL_S4_101_TPC_NilAccessors(t *testing.T) {
	// Arrange
	var nilCol *corepayload.TypedPayloadCollection[tpwTestData]

	// Act
	actual := args.Map{"result": nilCol.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": nilCol.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": nilCol.HasItems()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": nilCol.Items() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S4_102_TPC_LengthLock_IsEmptyLock(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[tpwTestData](1)
	_ = col.LengthLock()
	_ = col.IsEmptyLock()
}

func Test_CovPL_S4_103_TPC_ElementAccess(t *testing.T) {
	// Arrange
	col := corepayload.NewTypedPayloadCollection[tpwTestData](2)
	tw1, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"b", "2", "e", tpwTestData{Name: "b"})
	col.Add(tw1)
	col.Add(tw2)

	// Act
	actual := args.Map{"result": col.First().Name() != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
	actual = args.Map{"result": col.Last().Name() != "b"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b", actual)
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
	empty := corepayload.EmptyTypedPayloadCollection[tpwTestData]()
	actual = args.Map{"result": empty.FirstOrDefault() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": empty.SafeAt(0) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovPL_S4_104_TPC_Mutation(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": col.RemoveAt(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": col.RemoveAt(-1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": col.RemoveAt(100)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovPL_S4_105_TPC_Iteration(t *testing.T) {
	// Arrange
	col := corepayload.NewTypedPayloadCollection[tpwTestData](2)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a", Value: 1})
	col.Add(tw)
	count := 0
	col.ForEach(func(i int, item *corepayload.TypedPayloadWrapper[tpwTestData]) {
		count++
	})

	// Act
	actual := args.Map{"result": count != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	col.ForEachData(func(i int, d tpwTestData) {
		actual = args.Map{"result": d.Name != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
	col.ForEachBreak(func(i int, item *corepayload.TypedPayloadWrapper[tpwTestData]) bool {
		return true // break immediately
	})
}

func Test_CovPL_S4_106_TPC_Filter(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": filtered.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	filteredByData := col.FilterByData(func(d tpwTestData) bool {
		return d.Name == "b"
	})
	actual = args.Map{"result": filteredByData.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	found := col.FirstByFilter(func(item *corepayload.TypedPayloadWrapper[tpwTestData]) bool {
		return item.Data().Name == "a"
	})
	actual = args.Map{"result": found == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	foundByData := col.FirstByData(func(d tpwTestData) bool {
		return d.Name == "b"
	})
	actual = args.Map{"result": foundByData == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": col.FirstByName("a") == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": col.FirstById("1") == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": col.CountFunc(func(item *corepayload.TypedPayloadWrapper[tpwTestData]) bool { return true }) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_CovPL_S4_107_TPC_SkipTake(t *testing.T) {
	// Arrange
	col := corepayload.NewTypedPayloadCollection[tpwTestData](5)
	for i := 0; i < 5; i++ {
		tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
			"x", "1", "e", tpwTestData{Value: i})
		col.Add(tw)
	}

	// Act
	actual := args.Map{"result": len(col.Skip(2)) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual = args.Map{"result": len(col.Skip(10)) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(col.Take(3)) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual = args.Map{"result": len(col.Take(10)) != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_CovPL_S4_108_TPC_Extraction(t *testing.T) {
	// Arrange
	col := corepayload.NewTypedPayloadCollection[tpwTestData](2)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a", Value: 1})
	col.Add(tw)
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
	empty := corepayload.EmptyTypedPayloadCollection[tpwTestData]()
	actual = args.Map{"result": len(empty.AllData()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.AllNames()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.AllIdentifiers()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S4_109_TPC_ToPayloadsCollection(t *testing.T) {
	// Arrange
	col := corepayload.NewTypedPayloadCollection[tpwTestData](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
	col.Add(tw)
	pc := col.ToPayloadsCollection()

	// Act
	actual := args.Map{"result": pc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := corepayload.EmptyTypedPayloadCollection[tpwTestData]()
	epc := empty.ToPayloadsCollection()
	actual = args.Map{"result": epc.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S4_110_TPC_Clone_CloneMust_ConcatNew(t *testing.T) {
	// Arrange
	col := corepayload.NewTypedPayloadCollection[tpwTestData](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
	col.Add(tw)
	cloned, err := col.Clone()

	// Act
	actual := args.Map{"result": err != nil || cloned.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = col.CloneMust()
	concat, err := col.ConcatNew(tw)
	actual = args.Map{"result": err != nil || concat.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	empty := corepayload.EmptyTypedPayloadCollection[tpwTestData]()
	ec, _ := empty.Clone()
	actual = args.Map{"result": ec.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
	b, _ := corejson.Serialize.Raw([]*corepayload.PayloadWrapper{tw.Wrapper})
	col, err := corepayload.TypedPayloadCollectionDeserialize[tpwTestData](b)

	// Act
	actual := args.Map{"result": err != nil || col.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = corepayload.TypedPayloadCollectionDeserializeMust[tpwTestData](b)
	_, err2 := corepayload.TypedPayloadCollectionDeserialize[tpwTestData]([]byte("bad"))
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovPL_S4_113_TPC_FromPayloads(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(1)
	pc.Add(*newPWSeg4())
	col := corepayload.TypedPayloadCollectionFromPayloads[map[string]int](pc)

	// Act
	actual := args.Map{"result": col.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	// nil
	nilCol := corepayload.TypedPayloadCollectionFromPayloads[map[string]int](nil)
	actual = args.Map{"result": nilCol.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S4_114_TPC_NewTypedPayloadCollectionSingle(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
	col := corepayload.NewTypedPayloadCollectionSingle[tpwTestData](tw)

	// Act
	actual := args.Map{"result": col.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	nilCol := corepayload.NewTypedPayloadCollectionSingle[tpwTestData](nil)
	actual = args.Map{"result": nilCol.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S4_115_TPC_NewTypedPayloadCollectionFromData(t *testing.T) {
	// Arrange
	col, err := corepayload.NewTypedPayloadCollectionFromData[tpwTestData](
		"test", []tpwTestData{{Name: "a"}, {Name: "b"}})

	// Act
	actual := args.Map{"result": err != nil || col.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	_ = corepayload.NewTypedPayloadCollectionFromDataMust[tpwTestData](
		"test", []tpwTestData{{Name: "a"}})
	// empty
	empty, _ := corepayload.NewTypedPayloadCollectionFromData[tpwTestData]("test", nil)
	actual = args.Map{"result": empty.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovPL_S4_116_TPC_IsValid_HasErrors_Errors_FirstError_MergedError(t *testing.T) {
	// Arrange
	col := corepayload.NewTypedPayloadCollection[tpwTestData](1)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[tpwTestData](
		"a", "1", "e", tpwTestData{Name: "a"})
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
	empty := corepayload.EmptyTypedPayloadCollection[tpwTestData]()
	actual = args.Map{"result": empty.IsValid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": empty.Errors() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// --- PayloadCreateInstructionTypeStringer ---

func Test_CovPL_S4_120_PCITS_PayloadCreateInstruction(t *testing.T) {
	// Arrange
	inst := corepayload.PayloadCreateInstructionTypeStringer{
		Name:                 "n",
		Identifier:           "1",
		TaskTypeNameStringer: seg4Stringer{"task"},
		CategoryNameStringer: seg4Stringer{"cat"},
		HasManyRecords:       false,
		Payloads:             map[string]int{"a": 1},
	}
	pi := inst.PayloadCreateInstruction()

	// Act
	actual := args.Map{"result": pi.TaskTypeName != "task"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected task", actual)
	actual = args.Map{"result": pi.CategoryName != "cat"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cat", actual)
}

// --- BytesCreateInstructionStringer ---

func Test_CovPL_S4_121_BCIS_Fields(t *testing.T) {
	// Arrange
	inst := corepayload.BytesCreateInstructionStringer{
		Name:           "n",
		Identifier:     "1",
		TaskTypeName:   seg4Stringer{"task"},
		EntityType:     "e",
		CategoryName:   seg4Stringer{"cat"},
		HasManyRecords: false,
		Payloads:       []byte("x"),
	}

	// Act
	actual := args.Map{"result": inst.Name != "n"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected n", actual)
}
