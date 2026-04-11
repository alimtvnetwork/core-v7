package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// testUserCov23, makeTypedWrapperCov23, makeCollectionCov23 are declared in Coverage23_TypedFuncs_test.go

// ── TypedPayloadWrapper setters ──

func Test_TypedPayloadWrapper_SetName(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	tw.SetName("renamed")

	// Assert
	actual := args.Map{"name": tw.Name()}
	expected := args.Map{"name": "renamed"}
	expected.ShouldBeEqual(t, 0, "SetName updates name -- renamed", actual)
}

func Test_TypedPayloadWrapper_SetIdentifier(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	tw.SetIdentifier("new-id")

	// Assert
	actual := args.Map{"id": tw.Identifier()}
	expected := args.Map{"id": "new-id"}
	expected.ShouldBeEqual(t, 0, "SetIdentifier updates id -- new-id", actual)
}

func Test_TypedPayloadWrapper_SetEntityType(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	tw.SetEntityType("NewEntity")

	// Assert
	actual := args.Map{"entity": tw.EntityType()}
	expected := args.Map{"entity": "NewEntity"}
	expected.ShouldBeEqual(t, 0, "SetEntityType updates entity -- NewEntity", actual)
}

func Test_TypedPayloadWrapper_SetCategoryName(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	tw.SetCategoryName("NewCat")

	// Assert
	actual := args.Map{"cat": tw.CategoryName()}
	expected := args.Map{"cat": "NewCat"}
	expected.ShouldBeEqual(t, 0, "SetCategoryName updates category -- NewCat", actual)
}

func Test_TypedPayloadWrapper_SetTypedData_FromTypedPayloadWrapperS(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	err := tw.SetTypedData(testUserCov23{Name: "Bob", Age: 30})

	// Assert
	actual := args.Map{
		"noError": err == nil,
		"name":    tw.Data().Name,
		"age":     tw.Data().Age,
	}
	expected := args.Map{
		"noError": true,
		"name":    "Bob",
		"age":     30,
	}
	expected.ShouldBeEqual(t, 0, "SetTypedData updates data -- Bob/30", actual)
}

func Test_TypedPayloadWrapper_SetTypedDataMust(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})
	panicked := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		tw.SetTypedDataMust(testUserCov23{Name: "Charlie"})
	}()

	// Assert
	actual := args.Map{
		"panicked": panicked,
		"name":     tw.Data().Name,
	}
	expected := args.Map{
		"panicked": false,
		"name":     "Charlie",
	}
	expected.ShouldBeEqual(t, 0, "SetTypedDataMust sets data -- Charlie", actual)
}

// ── TypedPayloadWrapper accessors ──

func Test_TypedPayloadWrapper_GetAsString(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("test", "1", "string", "hello")

	// Act
	val, ok := tw.GetAsString()

	// Assert
	actual := args.Map{
		"val": val,
		"ok": ok,
	}
	expected := args.Map{
		"val": "hello",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "GetAsString returns value -- string type", actual)
}

func Test_TypedPayloadWrapper_GetAsInt(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[int]("test", "1", "int", 42)

	// Act
	val, ok := tw.GetAsInt()

	// Assert
	actual := args.Map{
		"val": val,
		"ok": ok,
	}
	expected := args.Map{
		"val": 42,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "GetAsInt returns value -- int type", actual)
}

func Test_TypedPayloadWrapper_GetAsBool(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[bool]("test", "1", "bool", true)

	// Act
	val, ok := tw.GetAsBool()

	// Assert
	actual := args.Map{
		"val": val,
		"ok": ok,
	}
	expected := args.Map{
		"val": true,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "GetAsBool returns value -- bool type", actual)
}

func Test_TypedPayloadWrapper_ValueString(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	result := tw.ValueString()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "ValueString returns non-empty -- struct type", actual)
}

func Test_TypedPayloadWrapper_ValueInt(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[int]("test", "1", "int", 42)

	// Act
	result := tw.ValueInt()

	// Assert
	actual := args.Map{"val": result}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ValueInt returns 42 -- int type", actual)
}

func Test_TypedPayloadWrapper_ValueBool(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[bool]("test", "1", "bool", true)

	// Act
	result := tw.ValueBool()

	// Assert
	actual := args.Map{"val": result}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "ValueBool returns true -- bool type", actual)
}

// ── Reparse / Clone / Clear / Dispose ──

func Test_TypedPayloadWrapper_Reparse(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	err := tw.Reparse()

	// Assert
	actual := args.Map{
		"noError": err == nil,
		"name":    tw.Data().Name,
	}
	expected := args.Map{
		"noError": true,
		"name":    "Alice",
	}
	expected.ShouldBeEqual(t, 0, "Reparse succeeds -- valid payload", actual)
}

func Test_TypedPayloadWrapper_ClonePtr_FromTypedPayloadWrapperS(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	cloned, err := tw.ClonePtr(true)

	// Assert
	actual := args.Map{
		"noError": err == nil,
		"name":    cloned.Data().Name,
	}
	expected := args.Map{
		"noError": true,
		"name":    "Alice",
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns deep clone -- valid", actual)
}

func Test_TypedPayloadWrapper_Clone_FromTypedPayloadWrapperS(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	cloned, err := tw.Clone(false)

	// Assert
	actual := args.Map{
		"noError": err == nil,
		"name":    cloned.Data().Name,
	}
	expected := args.Map{
		"noError": true,
		"name":    "Alice",
	}
	expected.ShouldBeEqual(t, 0, "Clone returns shallow clone -- valid", actual)
}

func Test_TypedPayloadWrapper_Clear(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	tw.Clear()

	// Assert
	actual := args.Map{
		"parsed":  tw.IsParsed(),
		"isEmpty": tw.IsEmpty(),
	}
	expected := args.Map{
		"parsed":  false,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Clear resets typed wrapper -- empty", actual)
}

func Test_TypedPayloadWrapper_Dispose(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	tw.Dispose()

	// Assert
	actual := args.Map{"isNull": tw.IsNull()}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "Dispose nils wrapper -- null", actual)
}

// ── Nil receiver tests ──

func Test_TypedPayloadWrapper_NilReceiver(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUserCov23]

	// Act & Assert
	actual := args.Map{
		"name":           tw.Name(),
		"id":             tw.Identifier(),
		"entity":         tw.EntityType(),
		"category":       tw.CategoryName(),
		"taskType":       tw.TaskTypeName(),
		"hasMany":        tw.HasManyRecords(),
		"hasError":       tw.HasError(),
		"isEmpty":        tw.IsEmpty(),
		"isNull":         tw.IsNull(),
		"length":         tw.Length(),
		"str":            tw.String(),
		"prettyJson":     tw.PrettyJsonString(),
		"jsonStr":        tw.JsonString(),
		"payloadsStr":    tw.PayloadsString(),
		"dynamicLen":     len(tw.DynamicPayloads()),
		"isParsed":       tw.IsParsed(),
		"attrsNil":       tw.Attributes() == nil,
		"errorNil":       tw.Error() == nil,
		"toWrapperNil":   tw.ToPayloadWrapper() == nil,
	}
	expected := args.Map{
		"name":           "",
		"id":             "",
		"entity":         "",
		"category":       "",
		"taskType":       "",
		"hasMany":        false,
		"hasError":       false,
		"isEmpty":        true,
		"isNull":         true,
		"length":         0,
		"str":            "",
		"prettyJson":     "",
		"jsonStr":        "",
		"payloadsStr":    "",
		"dynamicLen":     0,
		"isParsed":       false,
		"attrsNil":       true,
		"errorNil":       true,
		"toWrapperNil":   true,
	}
	expected.ShouldBeEqual(t, 0, "Nil receiver returns defaults -- all methods", actual)
}

// ── TypedPayloadCollection Lock methods ──

func Test_TypedPayloadCollection_LengthLock(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := col.LengthLock()

	// Assert
	actual := args.Map{"length": result}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "LengthLock returns 3 -- locked", actual)
}

func Test_TypedPayloadCollection_IsEmptyLock(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := col.IsEmptyLock()

	// Assert
	actual := args.Map{"isEmpty": result}
	expected := args.Map{"isEmpty": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyLock returns false -- has items", actual)
}

func Test_TypedPayloadCollection_AddLock(t *testing.T) {
	// Arrange
	col := corepayload.EmptyTypedPayloadCollection[testUserCov23]()
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	col.AddLock(tw)

	// Assert
	actual := args.Map{"length": col.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "AddLock adds item -- 1 item", actual)
}

func Test_TypedPayloadCollection_AddCollection(t *testing.T) {
	// Arrange
	col1 := makeCollectionCov23()
	col2 := corepayload.NewTypedPayloadCollection[testUserCov23](1)
	col2.Add(makeTypedWrapperCov23("user", "4", testUserCov23{Name: "Dave"}))

	// Act
	col1.AddCollection(col2)

	// Assert
	actual := args.Map{"length": col1.Length()}
	expected := args.Map{"length": 4}
	expected.ShouldBeEqual(t, 0, "AddCollection merges -- 4 items", actual)
}

func Test_TypedPayloadCollection_SafeAt(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	valid := col.SafeAt(1)
	invalid := col.SafeAt(99)

	// Assert
	actual := args.Map{
		"validNotNil": valid != nil,
		"invalidNil":  invalid == nil,
	}
	expected := args.Map{
		"validNotNil": true,
		"invalidNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "SafeAt returns correct -- bounds check", actual)
}

func Test_TypedPayloadCollection_ClearDispose(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	col.Clear()

	// Assert
	actual := args.Map{"isEmpty": col.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Clear empties collection -- empty", actual)
}

func Test_TypedPayloadCollection_Dispose(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	col.Dispose()

	// Assert
	actual := args.Map{"isEmpty": col.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dispose nils collection -- empty", actual)
}

// ── TypedPayloadCollectionFromPayloads ──

func Test_TypedPayloadCollectionFromPayloads(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()
	pc := col.ToPayloadsCollection()

	// Act
	result := corepayload.TypedPayloadCollectionFromPayloads[testUserCov23](pc)

	// Assert
	actual := args.Map{"length": result.Length()}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollectionFromPayloads returns 3 -- from PC", actual)
}

func Test_TypedPayloadCollectionFromPayloads_Nil(t *testing.T) {
	// Arrange & Act
	result := corepayload.TypedPayloadCollectionFromPayloads[testUserCov23](nil)

	// Assert
	actual := args.Map{"isEmpty": result.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollectionFromPayloads returns empty -- nil", actual)
}
