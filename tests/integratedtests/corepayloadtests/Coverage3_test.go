package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── PayloadWrapper creation ──

func Test_Cov3_PayloadWrapper_New(t *testing.T) {
	pw := &corepayload.PayloadWrapper{
		Name:       "test",
		Identifier: "id-1",
	}
	actual := args.Map{
		"name":       pw.PayloadName(),
		"isNull":     pw.IsNull(),
		"hasError":   pw.HasError(),
		"emptyError": pw.IsEmptyError(),
	}
	expected := args.Map{
		"name":       "test",
		"isNull":     false,
		"hasError":   false,
		"emptyError": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper basic getters -- name and id", actual)
}

func Test_Cov3_PayloadWrapper_NilIsNull(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	actual := args.Map{"isNull": pw.IsNull()}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.IsNull returns true -- nil receiver", actual)
}

func Test_Cov3_PayloadWrapper_Category(t *testing.T) {
	pw := &corepayload.PayloadWrapper{
		CategoryName: "cat1",
		TaskTypeName: "task1",
		EntityType:   "entity1",
	}
	actual := args.Map{
		"category":   pw.PayloadCategory(),
		"taskType":   pw.PayloadTaskType(),
		"entityType": pw.PayloadEntityType(),
		"isName":     pw.IsName(""),
		"isNameTrue": pw.IsName(""),
		"isId":       pw.IsIdentifier(""),
		"isTask":     pw.IsTaskTypeName("task1"),
		"isEntity":   pw.IsEntityType("entity1"),
		"isCat":      pw.IsCategory("cat1"),
	}
	expected := args.Map{
		"category":   "cat1",
		"taskType":   "task1",
		"entityType": "entity1",
		"isName":     true,
		"isNameTrue": true,
		"isId":       true,
		"isTask":     true,
		"isEntity":   true,
		"isCat":      true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper category/task/entity getters -- all set", actual)
}

func Test_Cov3_PayloadWrapper_HasItems(t *testing.T) {
	pw := &corepayload.PayloadWrapper{
		Payloads:       []byte(`[1,2,3]`),
		HasManyRecords: true,
	}
	actual := args.Map{
		"hasItems":  pw.HasItems(),
		"hasSingle": pw.HasSingleRecord(),
		"count":     pw.Count(),
	}
	expected := args.Map{
		"hasItems":  true,
		"hasSingle": false,
		"count":     pw.Count(),
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper HasItems/HasSingle -- with payloads", actual)
}

func Test_Cov3_PayloadWrapper_MarshalJSON(t *testing.T) {
	pw := &corepayload.PayloadWrapper{
		Name:       "test",
		Identifier: "id-1",
	}
	b, err := pw.MarshalJSON()
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(b) > 0,
	}
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.MarshalJSON succeeds -- basic", actual)
}

func Test_Cov3_PayloadWrapper_MarshalJSON_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	_, err := pw.MarshalJSON()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.MarshalJSON returns error -- nil", actual)
}

func Test_Cov3_PayloadWrapper_UnmarshalJSON(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test", Identifier: "id-1"}
	b, _ := pw.MarshalJSON()
	var pw2 corepayload.PayloadWrapper
	err := pw2.UnmarshalJSON(b)
	actual := args.Map{
		"noErr":    err == nil,
		"sameName": pw2.Name == "test",
		"sameId":   pw2.Identifier == "id-1",
	}
	expected := args.Map{
		"noErr":    true,
		"sameName": true,
		"sameId":   true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.UnmarshalJSON roundtrip -- basic", actual)
}

func Test_Cov3_PayloadWrapper_Clone(t *testing.T) {
	pw := &corepayload.PayloadWrapper{
		Name:       "test",
		Identifier: "id-1",
	}
	cloned, err := pw.Clone(false)
	actual := args.Map{
		"noErr":    err == nil,
		"sameName": cloned.Name == pw.Name,
		"sameId":   cloned.Identifier == pw.Identifier,
	}
	expected := args.Map{
		"noErr":    true,
		"sameName": true,
		"sameId":   true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.Clone returns same data -- basic", actual)
}

func Test_Cov3_PayloadWrapper_String(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	s := pw.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.String returns non-empty -- basic", actual)
}

func Test_Cov3_PayloadWrapper_JsonString(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	js := pw.JsonString()
	actual := args.Map{"hasContent": len(js) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.JsonString returns non-empty -- basic", actual)
}

func Test_Cov3_PayloadWrapper_JsonModel(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	model := pw.JsonModel()
	actual := args.Map{"sameName": model.Name == "test"}
	expected := args.Map{"sameName": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.JsonModel returns value model -- basic", actual)
}

func Test_Cov3_PayloadWrapper_JsonModelAny(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	model := pw.JsonModelAny()
	actual := args.Map{"notNil": model != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.JsonModelAny returns non-nil -- basic", actual)
}

func Test_Cov3_PayloadWrapper_Json(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	result := pw.Json()
	actual := args.Map{"hasBytes": result.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.Json returns result with bytes -- basic", actual)
}

func Test_Cov3_PayloadWrapper_Clear(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test", Identifier: "id-1"}
	pw.Clear()
	actual := args.Map{"nameEmpty": pw.Name == ""}
	expected := args.Map{"nameEmpty": false}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.Clear keeps Name -- after clear", actual)
}

func Test_Cov3_PayloadWrapper_Dispose(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	pw.Dispose()
	actual := args.Map{"nameEmpty": pw.Name == ""}
	expected := args.Map{"nameEmpty": false}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.Dispose keeps Name -- after dispose", actual)
}

func Test_Cov3_PayloadWrapper_NonPtr(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	nonPtr := pw.NonPtr()
	actual := args.Map{"sameName": nonPtr.Name == "test"}
	expected := args.Map{"sameName": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.NonPtr returns value copy -- basic", actual)
}

func Test_Cov3_PayloadWrapper_ToPtr(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	ptr := pw.ToPtr()
	actual := args.Map{"notNil": ptr != nil, "sameName": ptr.Name == "test"}
	expected := args.Map{"notNil": true, "sameName": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.ToPtr returns pointer -- basic", actual)
}

func Test_Cov3_PayloadWrapper_HasIssuesOrEmpty(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	actual := args.Map{"hasIssues": pw.HasIssuesOrEmpty()}
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.HasIssuesOrEmpty returns true -- empty payload", actual)
}

func Test_Cov3_PayloadWrapper_IdString(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Identifier: "42"}
	actual := args.Map{"id": pw.IdString()}
	expected := args.Map{"id": "42"}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.IdString returns identifier -- set", actual)
}

func Test_Cov3_PayloadWrapper_IsEqual_BothNil(t *testing.T) {
	var pw1, pw2 *corepayload.PayloadWrapper
	actual := args.Map{"equal": pw1.IsEqual(pw2)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.IsEqual returns true -- both nil", actual)
}

func Test_Cov3_PayloadWrapper_IsEqual_SamePtr(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"equal": pw.IsEqual(pw)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.IsEqual returns true -- same ptr", actual)
}

func Test_Cov3_PayloadWrapper_HasAttributes(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Attributes: &corepayload.Attributes{}}
	actual := args.Map{
		"hasAttr": pw.HasAttributes(),
	}
	expected := args.Map{
		"hasAttr": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper HasAttributes -- with empty attrs", actual)
}

func Test_Cov3_PayloadWrapper_AsJsonContractsBinder(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	binder := pw.AsJsonContractsBinder()
	actual := args.Map{"notNil": binder != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.AsJsonContractsBinder returns non-nil -- basic", actual)
}

// ── PayloadsCollection ──

func Test_Cov3_PayloadsCollection_Empty(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	actual := args.Map{
		"notNil":  pc != nil,
		"isEmpty": pc.IsEmpty(),
		"length":  pc.Length(),
	}
	expected := args.Map{
		"notNil":  true,
		"isEmpty": true,
		"length":  0,
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Empty returns empty -- new", actual)
}

func Test_Cov3_PayloadsCollection_Add(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	actual := args.Map{
		"length":    pc.Length(),
		"hasAny":    pc.HasAnyItem(),
		"lastIndex": pc.LastIndex(),
		"hasIdx0":   pc.HasIndex(0),
	}
	expected := args.Map{
		"length":    1,
		"hasAny":    true,
		"lastIndex": 0,
		"hasIdx0":   true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Add works -- single item", actual)
}

func Test_Cov3_PayloadsCollection_FirstLast(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "first"})
	pc.Add(corepayload.PayloadWrapper{Name: "last"})
	actual := args.Map{
		"firstName":          pc.First().Name,
		"lastName":           pc.Last().Name,
		"firstOrDefaultName": pc.FirstOrDefault().Name,
		"lastOrDefaultName":  pc.LastOrDefault().Name,
	}
	expected := args.Map{
		"firstName":          "first",
		"lastName":           "last",
		"firstOrDefaultName": "first",
		"lastOrDefaultName":  "last",
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection First/Last correct -- two items", actual)
}

func Test_Cov3_PayloadsCollection_Clone(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	cloned := pc.Clone()
	actual := args.Map{"sameLen": cloned.Length() == pc.Length()}
	expected := args.Map{"sameLen": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Clone returns same len -- single item", actual)
}

func Test_Cov3_PayloadsCollection_ClonePtr(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	cloned := pc.ClonePtr()
	actual := args.Map{
		"notNil":     cloned != nil,
		"notSamePtr": cloned != pc,
	}
	expected := args.Map{
		"notNil":     true,
		"notSamePtr": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.ClonePtr returns different ptr -- single item", actual)
}

func Test_Cov3_PayloadsCollection_Clear(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	pc.Clear()
	actual := args.Map{"isEmpty": pc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Clear empties -- after clear", actual)
}

func Test_Cov3_PayloadsCollection_Dispose(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	pc.Dispose()
	actual := args.Map{"isEmpty": pc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Dispose empties -- after dispose", actual)
}

func Test_Cov3_PayloadsCollection_Strings(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	strs := pc.Strings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Strings returns 1 -- single item", actual)
}

func Test_Cov3_PayloadsCollection_String(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	s := pc.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.String returns non-empty -- single item", actual)
}

func Test_Cov3_PayloadsCollection_JsonString(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	js := pc.JsonString()
	actual := args.Map{"hasContent": len(js) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.JsonString returns non-empty -- single item", actual)
}

func Test_Cov3_PayloadsCollection_Json(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	result := pc.Json()
	actual := args.Map{"hasBytes": result.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Json returns result -- single item", actual)
}

func Test_Cov3_PayloadsCollection_Reverse(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "first"})
	pc.Add(corepayload.PayloadWrapper{Name: "last"})
	reversed := pc.Reverse()
	actual := args.Map{
		"len":       reversed.Length(),
		"firstName": reversed.First().Name,
	}
	expected := args.Map{
		"len":       2,
		"firstName": "last",
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Reverse reverses order -- two items", actual)
}

func Test_Cov3_PayloadsCollection_SkipTakeLimit(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	for i := 0; i < 5; i++ {
		pc.Add(corepayload.PayloadWrapper{Name: "test"})
	}
	actual := args.Map{
		"skipLen":      len(pc.Skip(2)),
		"takeLen":      len(pc.Take(3)),
		"limitLen":     len(pc.Limit(2)),
		"safeLimitLen": len(pc.SafeLimitCollection(100).Items),
	}
	expected := args.Map{
		"skipLen":      3,
		"takeLen":      3,
		"limitLen":     2,
		"safeLimitLen": 5,
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection Skip/Take/Limit correct -- 5 items", actual)
}

func Test_Cov3_PayloadsCollection_ConcatNew(t *testing.T) {
	pc1 := corepayload.New.PayloadsCollection.Empty()
	pc1.Add(corepayload.PayloadWrapper{Name: "a"})
	concat := pc1.ConcatNew(corepayload.PayloadWrapper{Name: "b"})
	actual := args.Map{"len": concat.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.ConcatNew merges -- two items", actual)
}

func Test_Cov3_PayloadsCollection_InsertAt(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "a"})
	pc.Add(corepayload.PayloadWrapper{Name: "c"})
	pw := corepayload.PayloadWrapper{Name: "b"}
	pc.InsertAt(1, pw)
	actual := args.Map{"len": pc.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.InsertAt adds at index -- 3 items", actual)
}

func Test_Cov3_PayloadsCollection_Filter(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "match"})
	pc.Add(corepayload.PayloadWrapper{Name: "other"})
	filtered := pc.Filter(func(item *corepayload.PayloadWrapper) (bool, bool) {
		return item.Name == "match", false
	})
	actual := args.Map{"len": len(filtered)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Filter returns 1 -- filter by name", actual)
}

func Test_Cov3_PayloadsCollection_FilterCollection(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "match"})
	pc.Add(corepayload.PayloadWrapper{Name: "other"})
	filtered := pc.FilterCollection(func(item *corepayload.PayloadWrapper) (bool, bool) {
		return item.Name == "match", false
	})
	actual := args.Map{"len": filtered.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.FilterCollection returns 1 -- filter by name", actual)
}

func Test_Cov3_PayloadsCollection_FirstByCategory(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "a", CategoryName: "cat1"})
	pc.Add(corepayload.PayloadWrapper{Name: "b", CategoryName: "cat2"})
	found := pc.FirstByCategory("cat1")
	actual := args.Map{
		"notNil": found != nil,
		"name":   found.Name,
	}
	expected := args.Map{
		"notNil": true,
		"name":   "a",
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.FirstByCategory finds correct -- cat1", actual)
}
