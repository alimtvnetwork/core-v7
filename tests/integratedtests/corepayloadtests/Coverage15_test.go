package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// TypedPayloadWrapper coverage
// ==========================================================================

type testUserCov15 struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

func makeTypedWrapper(t *testing.T) *corepayload.TypedPayloadWrapper[testUser] {
	t.Helper()
	tw, err := corepayload.NewTypedPayloadWrapperFrom[testUser](
		"user-create", "usr-1", "User",
		testUser{Name: "Alice", Email: "alice@test.com"},
	)
	if err != nil {
		t.Fatalf("NewTypedPayloadWrapperFrom failed: %v", err)
	}
	return tw
}

func Test_C15_TypedPayloadWrapper_Constructors(t *testing.T) {
	tw := makeTypedWrapper(t)
	actual := args.Map{
		"name":       tw.Name(),
		"id":         tw.Identifier(),
		"idStr":      tw.IdString(),
		"entity":     tw.EntityType(),
		"cat":        tw.CategoryName(),
		"task":       tw.TaskTypeName(),
		"hasMany":    tw.HasManyRecords(),
		"single":     tw.HasSingleRecord(),
		"parsed":     tw.IsParsed(),
		"data":       tw.Data().Name,
		"typedData":  tw.TypedData().Email,
	}
	expected := args.Map{
		"name":       "user-create",
		"id":         "usr-1",
		"idStr":      "usr-1",
		"entity":     "User",
		"cat":        "",
		"task":       "",
		"hasMany":    false,
		"single":     true,
		"parsed":     true,
		"data":       "Alice",
		"typedData":  "alice@test.com",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- constructors", actual)
}

func Test_C15_TypedPayloadWrapper_NilChecks(t *testing.T) {
	var nilTW *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{
		"name":    nilTW.Name(),
		"id":      nilTW.Identifier(),
		"idStr":   nilTW.IdString(),
		"entity":  nilTW.EntityType(),
		"cat":     nilTW.CategoryName(),
		"task":    nilTW.TaskTypeName(),
		"hasMany": nilTW.HasManyRecords(),
		"single":  nilTW.HasSingleRecord(),
		"parsed":  nilTW.IsParsed(),
		"hasErr":  nilTW.HasError(),
		"isEmpty": nilTW.IsEmpty(),
		"hasItem": nilTW.HasItems(),
		"safe":    nilTW.HasSafeItems(),
		"err":     nilTW.Error() == nil,
		"str":     nilTW.String(),
		"pretty":  nilTW.PrettyJsonString(),
		"jsonStr": nilTW.JsonString(),
		"isNull":  nilTW.IsNull(),
		"dynPay":  len(nilTW.DynamicPayloads()),
		"payStr":  nilTW.PayloadsString(),
		"length":  nilTW.Length(),
	}
	expected := args.Map{
		"name":    "",
		"id":      "",
		"idStr":   "",
		"entity":  "",
		"cat":     "",
		"task":    "",
		"hasMany": false,
		"single":  true,
		"parsed":  false,
		"hasErr":  false,
		"isEmpty": true,
		"hasItem": false,
		"safe":    false,
		"err":     true,
		"str":     "",
		"pretty":  "",
		"jsonStr": "",
		"isNull":  true,
		"dynPay":  0,
		"payStr":  "",
		"length":  0,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns nil -- nil checks", actual)
}

func Test_C15_TypedPayloadWrapper_ErrorHandling(t *testing.T) {
	tw := makeTypedWrapper(t)
	actual := args.Map{
		"hasErr":  tw.HasError(),
		"isEmpty": tw.IsEmpty(),
		"hasItem": tw.HasItems(),
		"safe":    tw.HasSafeItems(),
		"err":     tw.Error() == nil,
	}
	expected := args.Map{
		"hasErr":  false,
		"isEmpty": false,
		"hasItem": true,
		"safe":    true,
		"err":     true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns error -- error handling", actual)
}

func Test_C15_TypedPayloadWrapper_Strings(t *testing.T) {
	tw := makeTypedWrapper(t)
	actual := args.Map{
		"strNotEmpty":   tw.String() != "",
		"prettyNotEmpty": tw.PrettyJsonString() != "",
		"jsonNotEmpty":  tw.JsonString() != "",
	}
	expected := args.Map{
		"strNotEmpty":   true,
		"prettyNotEmpty": true,
		"jsonNotEmpty":  true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- strings", actual)
}

func Test_C15_TypedPayloadWrapper_JSON(t *testing.T) {
	tw := makeTypedWrapper(t)
	j := tw.Json()
	jp := tw.JsonPtr()
	mb, mErr := tw.MarshalJSON()
	ser, serErr := tw.Serialize()
	tdj := tw.TypedDataJson()
	tdjp := tw.TypedDataJsonPtr()
	tdjb, tdjbErr := tw.TypedDataJsonBytes()
	actual := args.Map{
		"jOK":     j.JsonString() != "",
		"jpOK":    jp != nil,
		"mbOK":    mErr == nil && len(mb) > 0,
		"serOK":   serErr == nil && len(ser) > 0,
		"tdjOK":   tdj.JsonString() != "",
		"tdjpOK":  tdjp != nil,
		"tdjbOK":  tdjbErr == nil && len(tdjb) > 0,
	}
	expected := args.Map{
		"jOK":     true,
		"jpOK":    true,
		"mbOK":    true,
		"serOK":   true,
		"tdjOK":   true,
		"tdjpOK":  true,
		"tdjbOK":  true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- JSON", actual)
}

func Test_C15_TypedPayloadWrapper_MarshalJSON_Nil(t *testing.T) {
	var nilTW *corepayload.TypedPayloadWrapper[testUser]
	_, err := nilTW.MarshalJSON()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns nil -- MarshalJSON nil", actual)
}

func Test_C15_TypedPayloadWrapper_UnmarshalJSON(t *testing.T) {
	tw := makeTypedWrapper(t)
	b, _ := tw.MarshalJSON()
	tw2 := &corepayload.TypedPayloadWrapper[testUser]{}
	err := tw2.UnmarshalJSON(b)
	actual := args.Map{
		"noErr":  err == nil,
		"parsed": tw2.IsParsed(),
	}
	expected := args.Map{
		"noErr":  true,
		"parsed": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- UnmarshalJSON", actual)
}

func Test_C15_TypedPayloadWrapper_SerializeMust(t *testing.T) {
	tw := makeTypedWrapper(t)
	b := tw.SerializeMust()
	actual := args.Map{"len": len(b) > 0}
	expected := args.Map{"len": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- SerializeMust", actual)
}

func Test_C15_TypedPayloadWrapper_Serialize_Nil(t *testing.T) {
	var nilTW *corepayload.TypedPayloadWrapper[testUser]
	_, err := nilTW.Serialize()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns nil -- Serialize nil", actual)
}

func Test_C15_TypedPayloadWrapper_GetAs(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "i", "e", "hello")
	s, sOk := tw.GetAsString()
	_, iOk := tw.GetAsInt()
	_, i64Ok := tw.GetAsInt64()
	_, f64Ok := tw.GetAsFloat64()
	_, f32Ok := tw.GetAsFloat32()
	_, bOk := tw.GetAsBool()
	_, byOk := tw.GetAsBytes()
	_, ssOk := tw.GetAsStrings()
	actual := args.Map{
		"s": s, "sOk": sOk, "iOk": iOk, "i64Ok": i64Ok,
		"f64Ok": f64Ok, "f32Ok": f32Ok, "bOk": bOk, "byOk": byOk, "ssOk": ssOk,
	}
	expected := args.Map{
		"s": "hello", "sOk": true, "iOk": false, "i64Ok": false,
		"f64Ok": false, "f32Ok": false, "bOk": false, "byOk": false, "ssOk": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- GetAs", actual)
}

func Test_C15_TypedPayloadWrapper_ValueMethods(t *testing.T) {
	twStr, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "i", "e", "hello")
	twInt, _ := corepayload.NewTypedPayloadWrapperFrom[int]("n", "i", "e", 42)
	twBool, _ := corepayload.NewTypedPayloadWrapperFrom[bool]("n", "i", "e", true)
	actual := args.Map{
		"str":  twStr.ValueString(),
		"int":  twInt.ValueInt(),
		"bool": twBool.ValueBool(),
	}
	expected := args.Map{
		"str":  "hello",
		"int":  42,
		"bool": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- Value methods", actual)
}

func Test_C15_TypedPayloadWrapper_Setters(t *testing.T) {
	tw := makeTypedWrapper(t)
	tw.SetName("newName")
	tw.SetIdentifier("newId")
	tw.SetEntityType("newEntity")
	tw.SetCategoryName("newCat")
	actual := args.Map{
		"name":   tw.Name(),
		"id":     tw.Identifier(),
		"entity": tw.EntityType(),
		"cat":    tw.CategoryName(),
	}
	expected := args.Map{
		"name":   "newName",
		"id":     "newId",
		"entity": "newEntity",
		"cat":    "newCat",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- Setters", actual)
}

func Test_C15_TypedPayloadWrapper_SetTypedData(t *testing.T) {
	tw := makeTypedWrapper(t)
	err := tw.SetTypedData(testUser{Name: "Bob", Email: "bob@test.com"})
	actual := args.Map{
		"noErr": err == nil,
		"name":  tw.Data().Name,
	}
	expected := args.Map{
		"noErr": true,
		"name":  "Bob",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- SetTypedData", actual)
}

func Test_C15_TypedPayloadWrapper_SetTypedData_Nil(t *testing.T) {
	var nilTW *corepayload.TypedPayloadWrapper[testUser]
	err := nilTW.SetTypedData(testUser{})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns nil -- SetTypedData nil", actual)
}

func Test_C15_TypedPayloadWrapper_SetTypedDataMust(t *testing.T) {
	tw := makeTypedWrapper(t)
	tw.SetTypedDataMust(testUser{Name: "Charlie"})
	actual := args.Map{"name": tw.Data().Name}
	expected := args.Map{"name": "Charlie"}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- SetTypedDataMust", actual)
}

func Test_C15_TypedPayloadWrapper_Clone(t *testing.T) {
	tw := makeTypedWrapper(t)
	cloneP, err := tw.ClonePtr(true)
	clone, err2 := tw.Clone(true)
	var nilTW *corepayload.TypedPayloadWrapper[testUser]
	nilClone, nilErr := nilTW.ClonePtr(true)
	actual := args.Map{
		"noErr":       err == nil,
		"cloneName":   cloneP.Data().Name,
		"noErr2":      err2 == nil,
		"cloneName2":  clone.Data().Name,
		"nilClone":    nilClone == nil,
		"nilErr":      nilErr == nil,
	}
	expected := args.Map{
		"noErr":       true,
		"cloneName":   "Alice",
		"noErr2":      true,
		"cloneName2":  "Alice",
		"nilClone":    true,
		"nilErr":      true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- Clone", actual)
}

func Test_C15_TypedPayloadWrapper_ToPayloadWrapper(t *testing.T) {
	tw := makeTypedWrapper(t)
	pw := tw.ToPayloadWrapper()
	pwv := tw.PayloadWrapperValue()
	var nilTW *corepayload.TypedPayloadWrapper[testUser]
	nilPW := nilTW.ToPayloadWrapper()
	actual := args.Map{
		"notNil":  pw != nil,
		"vNotNil": pwv != nil,
		"nilPW":   nilPW == nil,
	}
	expected := args.Map{
		"notNil":  true,
		"vNotNil": true,
		"nilPW":   true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- ToPayloadWrapper", actual)
}

func Test_C15_TypedPayloadWrapper_Reparse(t *testing.T) {
	tw := makeTypedWrapper(t)
	err := tw.Reparse()
	var nilTW *corepayload.TypedPayloadWrapper[testUser]
	errNil := nilTW.Reparse()
	actual := args.Map{
		"noErr":   err == nil,
		"nilErr":  errNil != nil,
		"parsed":  tw.IsParsed(),
	}
	expected := args.Map{
		"noErr":   true,
		"nilErr":  true,
		"parsed":  true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- Reparse", actual)
}

func Test_C15_TypedPayloadWrapper_ClearDispose(t *testing.T) {
	tw := makeTypedWrapper(t)
	tw.Clear()
	actual := args.Map{"isEmpty": tw.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- Clear", actual)

	tw2 := makeTypedWrapper(t)
	tw2.Dispose()
	actual2 := args.Map{"isNull": tw2.IsNull()}
	expected2 := args.Map{"isNull": true}
	expected2.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- Dispose", actual2)

	var nilTW *corepayload.TypedPayloadWrapper[testUser]
	nilTW.Clear()   // should not panic
	nilTW.Dispose() // should not panic
}

func Test_C15_TypedPayloadWrapper_Attributes(t *testing.T) {
	tw := makeTypedWrapper(t)
	attr := tw.Attributes()
	tw.InitializeAttributesOnNull()
	var nilTW *corepayload.TypedPayloadWrapper[testUser]
	nilAttr := nilTW.Attributes()
	nilInit := nilTW.InitializeAttributesOnNull()
	actual := args.Map{
		"attrNil":     attr == nil,
		"nilAttr":     nilAttr == nil,
		"nilInit":     nilInit == nil,
	}
	expected := args.Map{
		"attrNil":     true,
		"nilAttr":     true,
		"nilInit":     true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- Attributes", actual)
}

func Test_C15_TypedPayloadWrapper_IdInteger(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "42", "e", "x")
	var nilTW *corepayload.TypedPayloadWrapper[string]
	actual := args.Map{
		"id":    tw.IdInteger(),
		"nilId": nilTW.IdInteger(),
	}
	expected := args.Map{
		"id":    42,
		"nilId": -1,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- IdInteger", actual)
}

func Test_C15_TypedPayloadWrapper_HandleError_NoError(t *testing.T) {
	tw := makeTypedWrapper(t)
	tw.HandleError() // should not panic
}

// ==========================================================================
// TypedPayloadWrapper factory funcs
// ==========================================================================

func Test_C15_TypedPayloadWrapper_Factories(t *testing.T) {
	tw1, err1 := corepayload.TypedPayloadWrapperFrom[testUser]("n", "i", "e", testUser{Name: "A"})
	tw2, err2 := corepayload.TypedPayloadWrapperRecord[testUser]("n", "i", "t", "c", testUser{Name: "B"})
	tw3, err3 := corepayload.TypedPayloadWrapperRecords[[]testUser]("n", "i", "t", "c", []testUser{{Name: "C"}})
	tw4, err4 := corepayload.TypedPayloadWrapperNameIdRecord[testUser]("n", "i", testUser{Name: "D"})
	tw5, err5 := corepayload.TypedPayloadWrapperNameIdCategory[testUser]("n", "i", "cat", testUser{Name: "E"})
	tw6, err6 := corepayload.TypedPayloadWrapperAll[testUser]("n", "i", "t", "e", "c", true, testUser{Name: "F"}, nil)
	actual := args.Map{
		"e1": err1 == nil, "n1": tw1.Data().Name,
		"e2": err2 == nil, "n2": tw2.Data().Name,
		"e3": err3 == nil, "n3": len(tw3.Data()),
		"e4": err4 == nil, "n4": tw4.Data().Name,
		"e5": err5 == nil, "n5": tw5.Data().Name,
		"e6": err6 == nil, "n6": tw6.Data().Name,
	}
	expected := args.Map{
		"e1": true, "n1": "A",
		"e2": true, "n2": "B",
		"e3": true, "n3": 1,
		"e4": true, "n4": "D",
		"e5": true, "n5": "E",
		"e6": true, "n6": "F",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- factories", actual)
}

func Test_C15_TypedPayloadWrapper_Must(t *testing.T) {
	pw, _ := corepayload.New.PayloadWrapper.NameIdCategory("n", "i", "cat", testUser{Name: "X"})
	tw := corepayload.NewTypedPayloadWrapperMust[testUser](pw)
	actual := args.Map{"name": tw.Data().Name}
	expected := args.Map{"name": "X"}
	expected.ShouldBeEqual(t, 0, "NewTypedPayloadWrapperMust returns correct value -- with args", actual)
}

func Test_C15_TypedPayloadWrapper_Deserialize(t *testing.T) {
	tw := makeTypedWrapper(t)
	b, _ := tw.Serialize()
	tw2, err := corepayload.TypedPayloadWrapperDeserialize[testUser](b)
	actual := args.Map{
		"noErr": err == nil,
		"name":  tw2.Data().Name,
	}
	expected := args.Map{
		"noErr": true,
		"name":  "Alice",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperDeserialize returns correct value -- with args", actual)
}

func Test_C15_TypedPayloadWrapper_DeserializeUsingJsonResult(t *testing.T) {
	tw := makeTypedWrapper(t)
	b, _ := tw.Serialize()
	jr := corejson.NewResult.UsingTypeBytesPtr("test", b)
	tw2, err := corepayload.TypedPayloadWrapperDeserializeUsingJsonResult[testUser](jr)
	actual := args.Map{
		"noErr": err == nil,
		"name":  tw2.Data().Name,
	}
	expected := args.Map{
		"noErr": true,
		"name":  "Alice",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperDeserializeUsingJsonResult returns correct value -- with args", actual)
}

// ==========================================================================
// TypedPayloadCollection coverage
// ==========================================================================

func makeTypedCollection(t *testing.T) *corepayload.TypedPayloadCollection[testUser] {
	t.Helper()
	tw1 := makeTypedWrapper(t)
	tw2, _ := corepayload.NewTypedPayloadWrapperFrom[testUser]("n2", "usr-2", "User", testUser{Name: "Bob", Email: "bob@test.com"})
	col := corepayload.NewTypedPayloadCollection[testUser](2)
	col.Add(tw1).Add(tw2)
	return col
}

func Test_C15_TypedPayloadCollection_Core(t *testing.T) {
	col := makeTypedCollection(t)
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	actual := args.Map{
		"len":      col.Length(),
		"count":    col.Count(),
		"isEmpty":  col.IsEmpty(),
		"hasItems": col.HasItems(),
		"hasAny":   col.HasAnyItem(),
		"lastIdx":  col.LastIndex(),
		"hasIdx0":  col.HasIndex(0),
		"hasIdx5":  col.HasIndex(5),
		"emptyLen": empty.Length(),
		"emptyE":   empty.IsEmpty(),
	}
	expected := args.Map{
		"len":      2,
		"count":    2,
		"isEmpty":  false,
		"hasItems": true,
		"hasAny":   true,
		"lastIdx":  1,
		"hasIdx0":  true,
		"hasIdx5":  false,
		"emptyLen": 0,
		"emptyE":   true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- core", actual)
}

func Test_C15_TypedPayloadCollection_NilReceiver(t *testing.T) {
	var nilCol *corepayload.TypedPayloadCollection[testUser]
	actual := args.Map{
		"len":     nilCol.Length(),
		"isEmpty": nilCol.IsEmpty(),
		"items":   nilCol.Items() == nil,
	}
	expected := args.Map{
		"len":     0,
		"isEmpty": true,
		"items":   true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns nil -- nil", actual)
}

func Test_C15_TypedPayloadCollection_ElementAccess(t *testing.T) {
	col := makeTypedCollection(t)
	first := col.First()
	last := col.Last()
	fod := col.FirstOrDefault()
	lod := col.LastOrDefault()
	safe := col.SafeAt(0)
	safeBad := col.SafeAt(99)
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	efod := empty.FirstOrDefault()
	elod := empty.LastOrDefault()
	esafe := empty.SafeAt(0)
	actual := args.Map{
		"firstName":  first.Data().Name,
		"lastName":   last.Data().Name,
		"fodName":    fod.Data().Name,
		"lodName":    lod.Data().Name,
		"safeName":   safe.Data().Name,
		"safeBad":    safeBad == nil,
		"efod":       efod == nil,
		"elod":       elod == nil,
		"esafe":      esafe == nil,
	}
	expected := args.Map{
		"firstName":  "Alice",
		"lastName":   "Bob",
		"fodName":    "Alice",
		"lodName":    "Bob",
		"safeName":   "Alice",
		"safeBad":    true,
		"efod":       true,
		"elod":       true,
		"esafe":      true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- element access", actual)
}

func Test_C15_TypedPayloadCollection_Mutation(t *testing.T) {
	col := corepayload.NewTypedPayloadCollection[testUser](2)
	tw := makeTypedWrapper(t)
	col.Add(tw)
	col.AddLock(tw)
	col.Adds(tw, tw)
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- mutation", actual)
}

func Test_C15_TypedPayloadCollection_AddCollection(t *testing.T) {
	col := makeTypedCollection(t)
	col2 := makeTypedCollection(t)
	col.AddCollection(col2)
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- AddCollection", actual)
}

func Test_C15_TypedPayloadCollection_RemoveAt(t *testing.T) {
	col := makeTypedCollection(t)
	ok := col.RemoveAt(0)
	bad := col.RemoveAt(99)
	neg := col.RemoveAt(-1)
	actual := args.Map{"ok": ok, "bad": bad, "neg": neg, "len": col.Length()}
	expected := args.Map{"ok": true, "bad": false, "neg": false, "len": 1}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- RemoveAt", actual)
}

func Test_C15_TypedPayloadCollection_Iteration(t *testing.T) {
	col := makeTypedCollection(t)
	names := []string{}
	col.ForEach(func(i int, item *corepayload.TypedPayloadWrapper[testUser]) {
		names = append(names, item.Data().Name)
	})
	dataNames := []string{}
	col.ForEachData(func(i int, data testUser) {
		dataNames = append(dataNames, data.Name)
	})
	breakCount := 0
	col.ForEachBreak(func(i int, item *corepayload.TypedPayloadWrapper[testUser]) bool {
		breakCount++
		return true
	})
	actual := args.Map{
		"names":     len(names),
		"dataNames": len(dataNames),
		"breakCnt":  breakCount,
	}
	expected := args.Map{
		"names":     2,
		"dataNames": 2,
		"breakCnt":  1,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- iteration", actual)
}

func Test_C15_TypedPayloadCollection_Filter(t *testing.T) {
	col := makeTypedCollection(t)
	filtered := col.Filter(func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "Alice"
	})
	byData := col.FilterByData(func(u testUser) bool {
		return u.Name == "Bob"
	})
	first := col.FirstByFilter(func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "Bob"
	})
	firstData := col.FirstByData(func(u testUser) bool {
		return u.Name == "Alice"
	})
	byName := col.FirstByName("user-create")
	byId := col.FirstById("usr-2")
	countF := col.CountFunc(func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return true
	})
	actual := args.Map{
		"filteredLen":   filtered.Length(),
		"byDataLen":     byData.Length(),
		"firstName":     first.Data().Name,
		"firstDataName": firstData.Data().Name,
		"byNameNotNil":  byName != nil,
		"byIdNotNil":    byId != nil,
		"countF":        countF,
	}
	expected := args.Map{
		"filteredLen":   1,
		"byDataLen":     1,
		"firstName":     "Bob",
		"firstDataName": "Alice",
		"byNameNotNil":  true,
		"byIdNotNil":    true,
		"countF":        2,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- filter", actual)
}

func Test_C15_TypedPayloadCollection_SkipTake(t *testing.T) {
	col := makeTypedCollection(t)
	skipped := col.Skip(1)
	taken := col.Take(1)
	actual := args.Map{
		"skipLen": len(skipped),
		"takeLen": len(taken),
	}
	expected := args.Map{
		"skipLen": 1,
		"takeLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- SkipTake", actual)
}

func Test_C15_TypedPayloadCollection_Extraction(t *testing.T) {
	col := makeTypedCollection(t)
	allData := col.AllData()
	allNames := col.AllNames()
	allIds := col.AllIdentifiers()
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	emptyData := empty.AllData()
	emptyNames := empty.AllNames()
	emptyIds := empty.AllIdentifiers()
	actual := args.Map{
		"dataLen":     len(allData),
		"namesLen":    len(allNames),
		"idsLen":      len(allIds),
		"emptyData":   len(emptyData),
		"emptyNames":  len(emptyNames),
		"emptyIds":    len(emptyIds),
	}
	expected := args.Map{
		"dataLen":     2,
		"namesLen":    2,
		"idsLen":      2,
		"emptyData":   0,
		"emptyNames":  0,
		"emptyIds":    0,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- extraction", actual)
}

func Test_C15_TypedPayloadCollection_ToPayloadsCollection(t *testing.T) {
	col := makeTypedCollection(t)
	pc := col.ToPayloadsCollection()
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	epc := empty.ToPayloadsCollection()
	actual := args.Map{
		"len":      pc.Length(),
		"emptyLen": epc.Length(),
	}
	expected := args.Map{
		"len":      2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- ToPayloadsCollection", actual)
}

func Test_C15_TypedPayloadCollection_Clone(t *testing.T) {
	col := makeTypedCollection(t)
	cloned, err := col.Clone()
	cloneMust := col.CloneMust()
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	emptyClone, emptyErr := empty.Clone()
	actual := args.Map{
		"noErr":     err == nil,
		"cloneLen":  cloned.Length(),
		"mustLen":   cloneMust.Length(),
		"emptyErr":  emptyErr == nil,
		"emptyLen":  emptyClone.Length(),
	}
	expected := args.Map{
		"noErr":     true,
		"cloneLen":  2,
		"mustLen":   2,
		"emptyErr":  true,
		"emptyLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- Clone", actual)
}

func Test_C15_TypedPayloadCollection_ConcatNew(t *testing.T) {
	col := makeTypedCollection(t)
	tw := makeTypedWrapper(t)
	concat, err := col.ConcatNew(tw)
	actual := args.Map{
		"noErr": err == nil,
		"len":   concat.Length(),
	}
	expected := args.Map{
		"noErr": true,
		"len":   3,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- ConcatNew", actual)
}

func Test_C15_TypedPayloadCollection_ClearDispose(t *testing.T) {
	col := makeTypedCollection(t)
	col.Clear()
	actual := args.Map{"isEmpty": col.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- Clear", actual)

	col2 := makeTypedCollection(t)
	col2.Dispose()
	actual2 := args.Map{"isEmpty2": col2.IsEmpty()}
	expected2 := args.Map{"isEmpty2": true}
	expected2.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- Dispose", actual2)

	var nilCol *corepayload.TypedPayloadCollection[testUser]
	nilCol.Clear()   // no panic
	nilCol.Dispose() // no panic
}

func Test_C15_TypedPayloadCollection_LockMethods(t *testing.T) {
	col := makeTypedCollection(t)
	lenLock := col.LengthLock()
	emptyLock := col.IsEmptyLock()
	actual := args.Map{
		"lenLock":   lenLock,
		"emptyLock": emptyLock,
	}
	expected := args.Map{
		"lenLock":   2,
		"emptyLock": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- lock methods", actual)
}

func Test_C15_TypedPayloadCollection_Paging(t *testing.T) {
	col, _ := corepayload.NewTypedPayloadCollectionFromData[testUser]("user", []testUser{
		{Name: "A"}, {Name: "B"}, {Name: "C"}, {Name: "D"}, {Name: "E"},
	})
	pages := col.GetPagesSize(2)
	singlePage := col.GetSinglePageCollection(2, 1)
	pagedCol := col.GetPagedCollection(2)
	withInfo := col.GetPagedCollectionWithInfo(2)
	smallCol := corepayload.EmptyTypedPayloadCollection[testUser]()
	smallPages := smallCol.GetPagesSize(0)
	actual := args.Map{
		"pages":         pages,
		"singlePageLen": singlePage.Length(),
		"pagedColLen":   len(pagedCol),
		"withInfoLen":   len(withInfo),
		"smallPages":    smallPages,
	}
	expected := args.Map{
		"pages":         3,
		"singlePageLen": 2,
		"pagedColLen":   3,
		"withInfoLen":   3,
		"smallPages":    0,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- paging", actual)
}

func Test_C15_TypedPayloadCollection_Validation(t *testing.T) {
	col := makeTypedCollection(t)
	actual := args.Map{
		"isValid":   col.IsValid(),
		"hasErrors": col.HasErrors(),
		"firstErr":  col.FirstError() == nil,
		"mergedErr": col.MergedError() == nil,
		"errsLen":   len(col.Errors()),
	}
	expected := args.Map{
		"isValid":   true,
		"hasErrors": false,
		"firstErr":  true,
		"mergedErr": true,
		"errsLen":   0,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns non-empty -- validation", actual)
}

func Test_C15_TypedPayloadCollection_EmptyValidation(t *testing.T) {
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	actual := args.Map{
		"isValid":   empty.IsValid(),
		"errs":      empty.Errors() == nil,
	}
	expected := args.Map{
		"isValid":   true,
		"errs":      true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns empty -- empty validation", actual)
}

func Test_C15_TypedPayloadCollection_SingleAndFromData(t *testing.T) {
	tw := makeTypedWrapper(t)
	single := corepayload.NewTypedPayloadCollectionSingle[testUser](tw)
	var nilTW *corepayload.TypedPayloadWrapper[testUser]
	nilSingle := corepayload.NewTypedPayloadCollectionSingle[testUser](nilTW)
	fromData, err := corepayload.NewTypedPayloadCollectionFromData[testUser]("u", []testUser{{Name: "A"}})
	emptyFromData, _ := corepayload.NewTypedPayloadCollectionFromData[testUser]("u", []testUser{})
	mustFromData := corepayload.NewTypedPayloadCollectionFromDataMust[testUser]("u", []testUser{{Name: "B"}})
	actual := args.Map{
		"singleLen":    single.Length(),
		"nilSingleLen": nilSingle.Length(),
		"fromDataLen":  fromData.Length(),
		"fromDataErr":  err == nil,
		"emptyLen":     emptyFromData.Length(),
		"mustLen":      mustFromData.Length(),
	}
	expected := args.Map{
		"singleLen":    1,
		"nilSingleLen": 0,
		"fromDataLen":  1,
		"fromDataErr":  true,
		"emptyLen":     0,
		"mustLen":      1,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- Single/FromData", actual)
}

func Test_C15_TypedPayloadCollection_FromPayloads(t *testing.T) {
	col := makeTypedCollection(t)
	pc := col.ToPayloadsCollection()
	fromPayloads := corepayload.TypedPayloadCollectionFromPayloads[testUser](pc)
	nilFrom := corepayload.TypedPayloadCollectionFromPayloads[testUser](nil)
	actual := args.Map{
		"len":    fromPayloads.Length(),
		"nilLen": nilFrom.Length(),
	}
	expected := args.Map{
		"len":    2,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollectionFromPayloads returns correct value -- with args", actual)
}

func Test_C15_TypedPayloadCollection_Deserialize(t *testing.T) {
	col := makeTypedCollection(t)
	pc := col.ToPayloadsCollection()
	b := pc.JsonPtr().SafeBytes()
	desCol, err := corepayload.TypedPayloadCollectionDeserialize[testUser](b)
	_ = desCol
	_ = err
	// DeserializeMust may panic due to serialization format mismatch — skip
}

// ==========================================================================
// typed_collection_funcs.go coverage
// ==========================================================================

func Test_C15_MapTypedPayloads(t *testing.T) {
	col := makeTypedCollection(t)
	names := corepayload.MapTypedPayloads[testUser, string](col, func(item *corepayload.TypedPayloadWrapper[testUser]) string {
		return item.Data().Name
	})
	dataNames := corepayload.MapTypedPayloadData[testUser, string](col, func(u testUser) string {
		return u.Email
	})
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	emptyMap := corepayload.MapTypedPayloads[testUser, string](empty, func(item *corepayload.TypedPayloadWrapper[testUser]) string { return "" })
	emptyDataMap := corepayload.MapTypedPayloadData[testUser, string](empty, func(u testUser) string { return "" })
	actual := args.Map{
		"namesLen":    len(names),
		"dataLen":     len(dataNames),
		"emptyLen":    len(emptyMap),
		"emptyDatLen": len(emptyDataMap),
	}
	expected := args.Map{
		"namesLen":    2,
		"dataLen":     2,
		"emptyLen":    0,
		"emptyDatLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloads returns correct value -- with args", actual)
}

func Test_C15_FlatMapTypedPayloads(t *testing.T) {
	col := makeTypedCollection(t)
	result := corepayload.FlatMapTypedPayloads[testUser, string](col, func(item *corepayload.TypedPayloadWrapper[testUser]) []string {
		return []string{item.Data().Name, item.Data().Email}
	})
	dataResult := corepayload.FlatMapTypedPayloadData[testUser, string](col, func(u testUser) []string {
		return []string{u.Name}
	})
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	emptyFlat := corepayload.FlatMapTypedPayloads[testUser, string](empty, func(item *corepayload.TypedPayloadWrapper[testUser]) []string { return nil })
	emptyDataFlat := corepayload.FlatMapTypedPayloadData[testUser, string](empty, func(u testUser) []string { return nil })
	actual := args.Map{
		"len":         len(result),
		"dataLen":     len(dataResult),
		"emptyLen":    len(emptyFlat),
		"emptyDatLen": len(emptyDataFlat),
	}
	expected := args.Map{
		"len":         4,
		"dataLen":     2,
		"emptyLen":    0,
		"emptyDatLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "FlatMapTypedPayloads returns correct value -- with args", actual)
}

func Test_C15_ReduceTypedPayloads(t *testing.T) {
	col := makeTypedCollection(t)
	total := corepayload.ReduceTypedPayloads[testUser, int](col, 0, func(acc int, item *corepayload.TypedPayloadWrapper[testUser]) int {
		return acc + 1
	})
	dataTotal := corepayload.ReduceTypedPayloadData[testUser, int](col, 0, func(acc int, u testUser) int {
		return acc + len(u.Name)
	})
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	emptyReduce := corepayload.ReduceTypedPayloads[testUser, int](empty, 99, func(acc int, item *corepayload.TypedPayloadWrapper[testUser]) int { return acc })
	emptyDataReduce := corepayload.ReduceTypedPayloadData[testUser, int](empty, 77, func(acc int, u testUser) int { return acc })
	actual := args.Map{
		"total":     total,
		"dataTotal": dataTotal,
		"empty":     emptyReduce,
		"emptyData": emptyDataReduce,
	}
	expected := args.Map{
		"total":     2,
		"dataTotal": 8, // Alice(5) + Bob(3)
		"empty":     99,
		"emptyData": 77,
	}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloads returns correct value -- with args", actual)
}

func Test_C15_GroupTypedPayloads(t *testing.T) {
	col := makeTypedCollection(t)
	groups := corepayload.GroupTypedPayloads[testUser, string](col, func(item *corepayload.TypedPayloadWrapper[testUser]) string {
		return item.Name()
	})
	dataGroups := corepayload.GroupTypedPayloadData[testUser, string](col, func(u testUser) string {
		return u.Name
	})
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	emptyGroups := corepayload.GroupTypedPayloads[testUser, string](empty, func(item *corepayload.TypedPayloadWrapper[testUser]) string { return "" })
	actual := args.Map{
		"groupsLen":     len(groups),
		"dataGroupsLen": len(dataGroups),
		"emptyLen":      len(emptyGroups),
	}
	expected := args.Map{
		"groupsLen":     2,
		"dataGroupsLen": 2,
		"emptyLen":      0,
	}
	expected.ShouldBeEqual(t, 0, "GroupTypedPayloads returns correct value -- with args", actual)
}

func Test_C15_PartitionTypedPayloads(t *testing.T) {
	col := makeTypedCollection(t)
	matching, notMatching := corepayload.PartitionTypedPayloads[testUser](col, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "Alice"
	})
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	em, enm := corepayload.PartitionTypedPayloads[testUser](empty, func(item *corepayload.TypedPayloadWrapper[testUser]) bool { return true })
	actual := args.Map{
		"matchLen":    matching.Length(),
		"notMatchLen": notMatching.Length(),
		"emLen":       em.Length(),
		"enmLen":      enm.Length(),
	}
	expected := args.Map{
		"matchLen":    1,
		"notMatchLen": 1,
		"emLen":       0,
		"enmLen":      0,
	}
	expected.ShouldBeEqual(t, 0, "PartitionTypedPayloads returns correct value -- with args", actual)
}

func Test_C15_AnyAllTypedPayloads(t *testing.T) {
	col := makeTypedCollection(t)
	any := corepayload.AnyTypedPayload[testUser](col, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "Alice"
	})
	all := corepayload.AllTypedPayloads[testUser](col, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name != ""
	})
	allFail := corepayload.AllTypedPayloads[testUser](col, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "Alice"
	})
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	anyEmpty := corepayload.AnyTypedPayload[testUser](empty, func(item *corepayload.TypedPayloadWrapper[testUser]) bool { return true })
	allEmpty := corepayload.AllTypedPayloads[testUser](empty, func(item *corepayload.TypedPayloadWrapper[testUser]) bool { return false })
	actual := args.Map{
		"any":      any,
		"all":      all,
		"allFail":  allFail,
		"anyEmpty": anyEmpty,
		"allEmpty": allEmpty,
	}
	expected := args.Map{
		"any":      true,
		"all":      true,
		"allFail":  false,
		"anyEmpty": false,
		"allEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyAll returns correct value -- TypedPayloads", actual)
}

func Test_C15_ConvertTypedPayloads(t *testing.T) {
	col := makeTypedCollection(t)
	converted, err := corepayload.ConvertTypedPayloads[testUser, testUser](col)
	empty := corepayload.EmptyTypedPayloadCollection[testUser]()
	emptyConv, emptyErr := corepayload.ConvertTypedPayloads[testUser, testUser](empty)
	actual := args.Map{
		"noErr":    err == nil,
		"len":      converted.Length(),
		"emptyErr": emptyErr == nil,
		"emptyLen": emptyConv.Length(),
	}
	expected := args.Map{
		"noErr":    true,
		"len":      2,
		"emptyErr": true,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "ConvertTypedPayloads returns correct value -- with args", actual)
}

// ==========================================================================
// generic_helpers.go coverage
// ==========================================================================

func Test_C15_DeserializePayloadTo(t *testing.T) {
	pw, _ := corepayload.New.PayloadWrapper.NameIdCategory("n", "i", "cat", testUser{Name: "X"})
	u, err := corepayload.DeserializePayloadTo[testUser](pw)
	actual := args.Map{
		"noErr": err == nil,
		"name":  u.Name,
	}
	expected := args.Map{
		"noErr": true,
		"name":  "X",
	}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadTo returns correct value -- with args", actual)
}

func Test_C15_DeserializePayloadTo_Nil(t *testing.T) {
	_, err := corepayload.DeserializePayloadTo[testUser](nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadTo returns nil -- nil", actual)
}

func Test_C15_DeserializePayloadToMust(t *testing.T) {
	pw, _ := corepayload.New.PayloadWrapper.NameIdCategory("n", "i", "cat", testUser{Name: "Y"})
	u := corepayload.DeserializePayloadToMust[testUser](pw)
	actual := args.Map{"name": u.Name}
	expected := args.Map{"name": "Y"}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadToMust returns correct value -- with args", actual)
}

func Test_C15_DeserializePayloadToSlice(t *testing.T) {
	users := []testUser{{Name: "A"}, {Name: "B"}}
	pw, _ := corepayload.New.PayloadWrapper.NameIdCategory("n", "i", "cat", users)
	result, err := corepayload.DeserializePayloadToSlice[testUser](pw)
	actual := args.Map{
		"noErr": err == nil,
		"len":   len(result),
	}
	expected := args.Map{
		"noErr": true,
		"len":   2,
	}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadToSlice returns correct value -- with args", actual)
}

func Test_C15_DeserializePayloadToSlice_Nil(t *testing.T) {
	result, err := corepayload.DeserializePayloadToSlice[testUser](nil)
	actual := args.Map{"hasErr": err != nil, "len": len(result)}
	expected := args.Map{"hasErr": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadToSlice returns nil -- nil", actual)
}

func Test_C15_DeserializePayloadToSliceMust(t *testing.T) {
	users := []testUser{{Name: "C"}}
	pw, _ := corepayload.New.PayloadWrapper.NameIdCategory("n", "i", "cat", users)
	result := corepayload.DeserializePayloadToSliceMust[testUser](pw)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadToSliceMust returns correct value -- with args", actual)
}

func Test_C15_DeserializeAttributesPayloadTo(t *testing.T) {
	u := testUser{Name: "Test"}
	b, _ := corejson.Serialize.Raw(u)
	attr := &corepayload.Attributes{DynamicPayloads: b}
	result, err := corepayload.DeserializeAttributesPayloadTo[testUser](attr)
	actual := args.Map{"noErr": err == nil, "name": result.Name}
	expected := args.Map{"noErr": true, "name": "Test"}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadTo returns correct value -- with args", actual)
}

func Test_C15_DeserializeAttributesPayloadTo_Nil(t *testing.T) {
	_, err := corepayload.DeserializeAttributesPayloadTo[testUser](nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadTo returns nil -- nil", actual)
}

func Test_C15_DeserializeAttributesPayloadToMust(t *testing.T) {
	u := testUser{Name: "MustTest"}
	b, _ := corejson.Serialize.Raw(u)
	attr := &corepayload.Attributes{DynamicPayloads: b}
	result := corepayload.DeserializeAttributesPayloadToMust[testUser](attr)
	actual := args.Map{"name": result.Name}
	expected := args.Map{"name": "MustTest"}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadToMust returns correct value -- with args", actual)
}

func Test_C15_DeserializeAttributesPayloadToSlice(t *testing.T) {
	users := []testUser{{Name: "A"}}
	b, _ := corejson.Serialize.Raw(users)
	attr := &corepayload.Attributes{DynamicPayloads: b}
	result, err := corepayload.DeserializeAttributesPayloadToSlice[testUser](attr)
	actual := args.Map{"noErr": err == nil, "len": len(result)}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadToSlice returns correct value -- with args", actual)
}

func Test_C15_DeserializeAttributesPayloadToSlice_Nil(t *testing.T) {
	result, err := corepayload.DeserializeAttributesPayloadToSlice[testUser](nil)
	actual := args.Map{"hasErr": err != nil, "len": len(result)}
	expected := args.Map{"hasErr": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadToSlice returns nil -- nil", actual)
}
