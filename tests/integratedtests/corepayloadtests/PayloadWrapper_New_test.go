package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── PayloadWrapper creation ──

func Test_PayloadWrapper_New(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{
		Name:       "test",
		Identifier: "id-1",
	}

	// Act
	actual := args.Map{
		"name":       pw.PayloadName(),
		"isNull":     pw.IsNull(),
		"hasError":   pw.HasError(),
		"emptyError": pw.IsEmptyError(),
	}

	// Assert
	expected := args.Map{
		"name":       "test",
		"isNull":     false,
		"hasError":   false,
		"emptyError": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper basic getters -- name and id", actual)
}

func Test_PayloadWrapper_NilIsNull(t *testing.T) {
	// Arrange
	var pw *corepayload.PayloadWrapper

	// Act
	actual := args.Map{"isNull": pw.IsNull()}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.IsNull returns true -- nil receiver", actual)
}

func Test_PayloadWrapper_Category(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{
		CategoryName: "cat1",
		TaskTypeName: "task1",
		EntityType:   "entity1",
	}

	// Act
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

	// Assert
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

func Test_PayloadWrapper_HasItems(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{
		Payloads:       []byte(`[1,2,3]`),
		HasManyRecords: true,
	}

	// Act
	actual := args.Map{
		"hasItems":  pw.HasItems(),
		"hasSingle": pw.HasSingleRecord(),
		"count":     pw.Count(),
	}

	// Assert
	expected := args.Map{
		"hasItems":  true,
		"hasSingle": false,
		"count":     pw.Count(),
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper HasItems/HasSingle -- with payloads", actual)
}

func Test_PayloadWrapper_MarshalJSON(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{
		Name:       "test",
		Identifier: "id-1",
	}
	b, err := pw.MarshalJSON()

	// Act
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.MarshalJSON succeeds -- basic", actual)
}

func Test_PayloadWrapper_MarshalJSON_Nil(t *testing.T) {
	// Arrange
	var pw *corepayload.PayloadWrapper
	_, err := pw.MarshalJSON()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.MarshalJSON returns error -- nil", actual)
}

func Test_PayloadWrapper_UnmarshalJSON(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test", Identifier: "id-1"}
	b, _ := pw.MarshalJSON()
	var pw2 corepayload.PayloadWrapper
	err := pw2.UnmarshalJSON(b)

	// Act
	actual := args.Map{
		"noErr":    err == nil,
		"sameName": pw2.Name == "test",
		"sameId":   pw2.Identifier == "id-1",
	}

	// Assert
	expected := args.Map{
		"noErr":    true,
		"sameName": true,
		"sameId":   true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.UnmarshalJSON roundtrip -- basic", actual)
}

func Test_PayloadWrapper_Clone_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{
		Name:       "test",
		Identifier: "id-1",
	}
	cloned, err := pw.Clone(false)

	// Act
	actual := args.Map{
		"noErr":    err == nil,
		"sameName": cloned.Name == pw.Name,
		"sameId":   cloned.Identifier == pw.Identifier,
	}

	// Assert
	expected := args.Map{
		"noErr":    true,
		"sameName": true,
		"sameId":   true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.Clone returns same data -- basic", actual)
}

func Test_PayloadWrapper_String(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test"}
	s := pw.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.String returns non-empty -- basic", actual)
}

func Test_PayloadWrapper_JsonString(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test"}
	js := pw.JsonString()

	// Act
	actual := args.Map{"hasContent": len(js) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.JsonString returns non-empty -- basic", actual)
}

func Test_PayloadWrapper_JsonModel(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test"}
	model := pw.JsonModel()

	// Act
	actual := args.Map{"sameName": model.Name == "test"}

	// Assert
	expected := args.Map{"sameName": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.JsonModel returns value model -- basic", actual)
}

func Test_PayloadWrapper_JsonModelAny(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test"}
	model := pw.JsonModelAny()

	// Act
	actual := args.Map{"notNil": model != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.JsonModelAny returns non-nil -- basic", actual)
}

func Test_PayloadWrapper_Json(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test"}
	result := pw.Json()

	// Act
	actual := args.Map{"hasBytes": result.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.Json returns result with bytes -- basic", actual)
}

func Test_PayloadWrapper_Clear(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test", Identifier: "id-1"}
	pw.Clear()

	// Act
	actual := args.Map{"nameEmpty": pw.Name == ""}

	// Assert
	expected := args.Map{"nameEmpty": false}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.Clear keeps Name -- after clear", actual)
}

func Test_PayloadWrapper_Dispose_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test"}
	pw.Dispose()

	// Act
	actual := args.Map{"nameEmpty": pw.Name == ""}

	// Assert
	expected := args.Map{"nameEmpty": false}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.Dispose keeps Name -- after dispose", actual)
}

func Test_PayloadWrapper_NonPtr(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test"}
	nonPtr := pw.NonPtr()

	// Act
	actual := args.Map{"sameName": nonPtr.Name == "test"}

	// Assert
	expected := args.Map{"sameName": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.NonPtr returns value copy -- basic", actual)
}

func Test_PayloadWrapper_ToPtr(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test"}
	ptr := pw.ToPtr()

	// Act
	actual := args.Map{
		"notNil": ptr != nil,
		"sameName": ptr.Name == "test",
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"sameName": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.ToPtr returns pointer -- basic", actual)
}

func Test_PayloadWrapper_HasIssuesOrEmpty(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}

	// Act
	actual := args.Map{"hasIssues": pw.HasIssuesOrEmpty()}

	// Assert
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.HasIssuesOrEmpty returns true -- empty payload", actual)
}

func Test_PayloadWrapper_IdString(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Identifier: "42"}

	// Act
	actual := args.Map{"id": pw.IdString()}

	// Assert
	expected := args.Map{"id": "42"}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.IdString returns identifier -- set", actual)
}

func Test_PayloadWrapper_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var pw1, pw2 *corepayload.PayloadWrapper

	// Act
	actual := args.Map{"equal": pw1.IsEqual(pw2)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.IsEqual returns true -- both nil", actual)
}

func Test_PayloadWrapper_IsEqual_SamePtr(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test"}

	// Act
	actual := args.Map{"equal": pw.IsEqual(pw)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.IsEqual returns true -- same ptr", actual)
}

func Test_PayloadWrapper_HasAttributes_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Attributes: &corepayload.Attributes{}}

	// Act
	actual := args.Map{
		"hasAttr": pw.HasAttributes(),
	}

	// Assert
	expected := args.Map{
		"hasAttr": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper HasAttributes -- with empty attrs", actual)
}

func Test_PayloadWrapper_AsJsonContractsBinder_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test"}
	binder := pw.AsJsonContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper.AsJsonContractsBinder returns non-nil -- basic", actual)
}

// ── PayloadsCollection ──

func Test_PayloadsCollection_Empty_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()

	// Act
	actual := args.Map{
		"notNil":  pc != nil,
		"isEmpty": pc.IsEmpty(),
		"length":  pc.Length(),
	}

	// Assert
	expected := args.Map{
		"notNil":  true,
		"isEmpty": true,
		"length":  0,
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Empty returns empty -- new", actual)
}

func Test_PayloadsCollection_Add_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})

	// Act
	actual := args.Map{
		"length":    pc.Length(),
		"hasAny":    pc.HasAnyItem(),
		"lastIndex": pc.LastIndex(),
		"hasIdx0":   pc.HasIndex(0),
	}

	// Assert
	expected := args.Map{
		"length":    1,
		"hasAny":    true,
		"lastIndex": 0,
		"hasIdx0":   true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Add works -- single item", actual)
}

func Test_PayloadsCollection_FirstLast(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "first"})
	pc.Add(corepayload.PayloadWrapper{Name: "last"})

	// Act
	actual := args.Map{
		"firstName":          pc.First().Name,
		"lastName":           pc.Last().Name,
		"firstOrDefaultName": pc.FirstOrDefault().Name,
		"lastOrDefaultName":  pc.LastOrDefault().Name,
	}

	// Assert
	expected := args.Map{
		"firstName":          "first",
		"lastName":           "last",
		"firstOrDefaultName": "first",
		"lastOrDefaultName":  "last",
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection First/Last correct -- two items", actual)
}

func Test_PayloadsCollection_Clone(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	cloned := pc.Clone()

	// Act
	actual := args.Map{"sameLen": cloned.Length() == pc.Length()}

	// Assert
	expected := args.Map{"sameLen": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Clone returns same len -- single item", actual)
}

func Test_PayloadsCollection_ClonePtr(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	cloned := pc.ClonePtr()

	// Act
	actual := args.Map{
		"notNil":     cloned != nil,
		"notSamePtr": cloned != pc,
	}

	// Assert
	expected := args.Map{
		"notNil":     true,
		"notSamePtr": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.ClonePtr returns different ptr -- single item", actual)
}

func Test_PayloadsCollection_Clear(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	pc.Clear()

	// Act
	actual := args.Map{"isEmpty": pc.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Clear empties -- after clear", actual)
}

func Test_PayloadsCollection_Dispose_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	pc.Dispose()

	// Act
	actual := args.Map{"isEmpty": pc.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Dispose empties -- after dispose", actual)
}

func Test_PayloadsCollection_Strings_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	strs := pc.Strings()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Strings returns 1 -- single item", actual)
}

func Test_PayloadsCollection_String(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	s := pc.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.String returns non-empty -- single item", actual)
}

func Test_PayloadsCollection_JsonString(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	js := pc.JsonString()

	// Act
	actual := args.Map{"hasContent": len(js) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.JsonString returns non-empty -- single item", actual)
}

func Test_PayloadsCollection_Json(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "test"})
	result := pc.Json()

	// Act
	actual := args.Map{"hasBytes": result.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Json returns result -- single item", actual)
}

func Test_PayloadsCollection_Reverse_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "first"})
	pc.Add(corepayload.PayloadWrapper{Name: "last"})
	reversed := pc.Reverse()

	// Act
	actual := args.Map{
		"len":       reversed.Length(),
		"firstName": reversed.First().Name,
	}

	// Assert
	expected := args.Map{
		"len":       2,
		"firstName": "last",
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Reverse reverses order -- two items", actual)
}

func Test_PayloadsCollection_SkipTakeLimit(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	for i := 0; i < 5; i++ {
		pc.Add(corepayload.PayloadWrapper{Name: "test"})
	}

	// Act
	actual := args.Map{
		"skipLen":      len(pc.Skip(2)),
		"takeLen":      len(pc.Take(3)),
		"limitLen":     len(pc.Limit(2)),
		"safeLimitLen": len(pc.SafeLimitCollection(100).Items),
	}

	// Assert
	expected := args.Map{
		"skipLen":      3,
		"takeLen":      3,
		"limitLen":     2,
		"safeLimitLen": 5,
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection Skip/Take/Limit correct -- 5 items", actual)
}

func Test_PayloadsCollection_ConcatNew_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pc1 := corepayload.New.PayloadsCollection.Empty()
	pc1.Add(corepayload.PayloadWrapper{Name: "a"})
	concat := pc1.ConcatNew(corepayload.PayloadWrapper{Name: "b"})

	// Act
	actual := args.Map{"len": concat.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.ConcatNew merges -- two items", actual)
}

func Test_PayloadsCollection_InsertAt_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "a"})
	pc.Add(corepayload.PayloadWrapper{Name: "c"})
	pw := corepayload.PayloadWrapper{Name: "b"}
	pc.InsertAt(1, pw)

	// Act
	actual := args.Map{"len": pc.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.InsertAt adds at index -- 3 items", actual)
}

func Test_PayloadsCollection_Filter_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "match"})
	pc.Add(corepayload.PayloadWrapper{Name: "other"})
	filtered := pc.Filter(func(item *corepayload.PayloadWrapper) (bool, bool) {
		return item.Name == "match", false
	})

	// Act
	actual := args.Map{"len": len(filtered)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.Filter returns 1 -- filter by name", actual)
}

func Test_PayloadsCollection_FilterCollection_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "match"})
	pc.Add(corepayload.PayloadWrapper{Name: "other"})
	filtered := pc.FilterCollection(func(item *corepayload.PayloadWrapper) (bool, bool) {
		return item.Name == "match", false
	})

	// Act
	actual := args.Map{"len": filtered.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.FilterCollection returns 1 -- filter by name", actual)
}

func Test_PayloadsCollection_FirstByCategory_FromPayloadWrapperNew(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	pc.Add(corepayload.PayloadWrapper{Name: "a", CategoryName: "cat1"})
	pc.Add(corepayload.PayloadWrapper{Name: "b", CategoryName: "cat2"})
	found := pc.FirstByCategory("cat1")

	// Act
	actual := args.Map{
		"notNil": found != nil,
		"name":   found.Name,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"name":   "a",
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection.FirstByCategory finds correct -- cat1", actual)
}
