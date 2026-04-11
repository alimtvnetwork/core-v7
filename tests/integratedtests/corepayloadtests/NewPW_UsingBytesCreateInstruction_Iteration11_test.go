package corepayloadtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

type stringer16 struct{ val string }

func (s stringer16) String() string { return s.val }

// ==========================================================================
// newPayloadWrapperCreator — UsingBytesCreateInstruction
// ==========================================================================

func Test_NewPW_UsingBytesCreateInstruction(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw("hello")
	pw := corepayload.New.PayloadWrapper.UsingBytesCreateInstruction(
		&corepayload.BytesCreateInstruction{
			Name: "n", Identifier: "id", TaskTypeName: "task",
			EntityType: "entity", CategoryName: "cat",
			HasManyRecords: false, Payloads: b,
		})

	// Act
	actual := args.Map{
		"name": pw.Name,
		"id": pw.Identifier,
		"entity": pw.EntityType,
	}

	// Assert
	expected := args.Map{
		"name": "n",
		"id": "id",
		"entity": "entity",
	}
	expected.ShouldBeEqual(t, 0, "UsingBytesCreateInstruction returns correct value -- with args", actual)
}

func Test_NewPW_UsingBytesCreateInstructionTypeStringer(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw("hello")
	pw := corepayload.New.PayloadWrapper.UsingBytesCreateInstructionTypeStringer(
		&corepayload.BytesCreateInstructionStringer{
			Name: "n", Identifier: "id",
			TaskTypeName: stringer16{"task"},
			CategoryName: stringer16{"cat"},
			EntityType:   "entity", Payloads: b,
		})

	// Act
	actual := args.Map{
		"name": pw.Name,
		"task": pw.TaskTypeName,
		"cat": pw.CategoryName,
	}

	// Assert
	expected := args.Map{
		"name": "n",
		"task": "task",
		"cat": "cat",
	}
	expected.ShouldBeEqual(t, 0, "UsingBytesCreateInstructionTypeStringer returns correct value -- with args", actual)
}

// ==========================================================================
// UsingCreateInstructionTypeStringer
// ==========================================================================

func Test_NewPW_UsingCreateInstructionTypeStringer(t *testing.T) {
	// Arrange
	pw, err := corepayload.New.PayloadWrapper.UsingCreateInstructionTypeStringer(
		&corepayload.PayloadCreateInstructionTypeStringer{
			Name: "n", Identifier: "id",
			TaskTypeNameStringer: stringer16{"task"},
			CategoryNameStringer: stringer16{"cat"},
			Payloads:             testUser{Name: "Alice"},
		})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": pw.Name,
		"task": pw.TaskTypeName,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "n",
		"task": "task",
	}
	expected.ShouldBeEqual(t, 0, "UsingCreateInstructionTypeStringer returns correct value -- with args", actual)
}

// ==========================================================================
// UsingCreateInstruction — string payload branch
// ==========================================================================

func Test_NewPW_UsingCreateInstruction_StringPayload(t *testing.T) {
	// Arrange
	pw, err := corepayload.New.PayloadWrapper.UsingCreateInstruction(
		&corepayload.PayloadCreateInstruction{
			Name: "n", Identifier: "id", TaskTypeName: "task",
			EntityType: "entity", CategoryName: "cat",
			Payloads: `{"Name":"Bob"}`,
		})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"entity": pw.EntityType,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"entity": "entity",
	}
	expected.ShouldBeEqual(t, 0, "UsingCreateInstruction returns correct value -- string payload", actual)
}

// ==========================================================================
// CreateUsingTypeStringer
// ==========================================================================

func Test_NewPW_CreateUsingTypeStringer(t *testing.T) {
	// Arrange
	pw, err := corepayload.New.PayloadWrapper.CreateUsingTypeStringer(
		"n", "id", stringer16{"task"}, stringer16{"cat"},
		testUser{Name: "X"})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"task": pw.TaskTypeName,
		"cat": pw.CategoryName,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"task": "task",
		"cat": "cat",
	}
	expected.ShouldBeEqual(t, 0, "CreateUsingTypeStringer returns correct value -- with args", actual)
}

// ==========================================================================
// NameIdCategoryStringer
// ==========================================================================

func Test_NewPW_NameIdCategoryStringer(t *testing.T) {
	// Arrange
	pw, err := corepayload.New.PayloadWrapper.NameIdCategoryStringer(
		"n", "id", stringer16{"cat"}, testUser{Name: "Y"})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"cat": pw.CategoryName,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"cat": "cat",
	}
	expected.ShouldBeEqual(t, 0, "NameIdCategoryStringer returns correct value -- with args", actual)
}

// ==========================================================================
// RecordsTypeStringer, RecordTypeStringer
// ==========================================================================

func Test_NewPW_RecordsTypeStringer(t *testing.T) {
	// Arrange
	pw, err := corepayload.New.PayloadWrapper.RecordsTypeStringer(
		"n", "id", stringer16{"task"}, stringer16{"cat"},
		[]testUser{{Name: "A"}, {Name: "B"}})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"many": pw.HasManyRecords,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"many": true,
	}
	expected.ShouldBeEqual(t, 0, "RecordsTypeStringer returns correct value -- with args", actual)
}

func Test_NewPW_RecordTypeStringer(t *testing.T) {
	// Arrange
	pw, err := corepayload.New.PayloadWrapper.RecordTypeStringer(
		"n", "id", stringer16{"task"}, stringer16{"cat"},
		testUser{Name: "Z"})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"task": pw.TaskTypeName,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"task": "task",
	}
	expected.ShouldBeEqual(t, 0, "RecordTypeStringer returns correct value -- with args", actual)
}

// ==========================================================================
// NameIdTaskStringerRecord, NameTaskNameRecord
// ==========================================================================

func Test_NewPW_NameIdTaskStringerRecord(t *testing.T) {
	// Arrange
	pw, err := corepayload.New.PayloadWrapper.NameIdTaskStringerRecord(
		"n", "id", stringer16{"task"}, testUser{Name: "Q"})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"task": pw.TaskTypeName,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"task": "task",
	}
	expected.ShouldBeEqual(t, 0, "NameIdTaskStringerRecord returns error -- with args", actual)
}

func Test_NewPW_NameTaskNameRecord(t *testing.T) {
	// Arrange
	pw, err := corepayload.New.PayloadWrapper.NameTaskNameRecord(
		"id", "task", testUser{Name: "R"})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"task": pw.TaskTypeName,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"task": "task",
	}
	expected.ShouldBeEqual(t, 0, "NameTaskNameRecord returns correct value -- with args", actual)
}

// ==========================================================================
// ManyRecords
// ==========================================================================

func Test_NewPW_ManyRecords(t *testing.T) {
	// Arrange
	pw, err := corepayload.New.PayloadWrapper.ManyRecords(
		"n", "id", "task", "cat",
		[]testUser{{Name: "A"}})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": pw.Name,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "n",
	}
	expected.ShouldBeEqual(t, 0, "ManyRecords returns correct value -- with args", actual)
}

// ==========================================================================
// PayloadsCollection — DeserializeMust, DeserializeToMany, DeserializeUsingJsonResult
// ==========================================================================

func Test_NewPC_DeserializeMust(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	b, _ := corejson.Serialize.Raw(pc)
	result := corepayload.New.PayloadsCollection.DeserializeMust(b)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewPC.DeserializeMust returns correct value -- with args", actual)
}

func Test_NewPC_DeserializeToMany_Valid(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	b, _ := corejson.Serialize.Raw([]*corepayload.PayloadsCollection{pc})
	result, err := corepayload.New.PayloadsCollection.DeserializeToMany(b)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(result),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "NewPC.DeserializeToMany returns non-empty -- valid", actual)
}

func Test_NewPC_DeserializeUsingJsonResult_Valid(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	b, _ := corejson.Serialize.Raw(pc)
	jr := corejson.NewResult.UsingTypeBytesPtr("test", b)
	result, err := corepayload.New.PayloadsCollection.DeserializeUsingJsonResult(jr)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "NewPC.DeserializeUsingJsonResult returns non-empty -- valid", actual)
}

// ==========================================================================
// PayloadsCollectionGetters — Dynamic accessors
// ==========================================================================

func makePC3(t *testing.T) *corepayload.PayloadsCollection {
	t.Helper()
	pw1, _ := corepayload.New.PayloadWrapper.NameIdCategory("n1", "1", "cat", "a")
	pw2, _ := corepayload.New.PayloadWrapper.NameIdCategory("n2", "2", "cat", "b")
	pw3, _ := corepayload.New.PayloadWrapper.NameIdCategory("n3", "3", "cat", "c")
	return corepayload.New.PayloadsCollection.UsingWrappers(pw1, pw2, pw3)
}

func Test_PC_FirstDynamic(t *testing.T) {
	// Arrange
	pc := makePC3(t)

	// Act
	actual := args.Map{"notNil": pc.FirstDynamic() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FirstDynamic returns correct value -- with args", actual)
}

func Test_PC_FirstDynamic_Nil(t *testing.T) {
	// Arrange
	var pc *corepayload.PayloadsCollection

	// Act
	actual := args.Map{"nil": pc.FirstDynamic() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "FirstDynamic returns nil -- nil", actual)
}

func Test_PC_LastDynamic(t *testing.T) {
	// Arrange
	pc := makePC3(t)

	// Act
	actual := args.Map{"notNil": pc.LastDynamic() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LastDynamic returns correct value -- with args", actual)
}

func Test_PC_LastDynamic_Nil(t *testing.T) {
	// Arrange
	var pc *corepayload.PayloadsCollection

	// Act
	actual := args.Map{"nil": pc.LastDynamic() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "LastDynamic returns nil -- nil", actual)
}

func Test_PC_FirstOrDefaultDynamic(t *testing.T) {
	// Arrange
	pc := makePC3(t)

	// Act
	actual := args.Map{"notNil": pc.FirstOrDefaultDynamic() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultDynamic returns correct value -- with args", actual)
}

func Test_PC_LastOrDefaultDynamic(t *testing.T) {
	// Arrange
	pc := makePC3(t)

	// Act
	actual := args.Map{"notNil": pc.LastOrDefaultDynamic() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LastOrDefaultDynamic returns correct value -- with args", actual)
}

// ==========================================================================
// PayloadsCollectionGetters — Slice operations
// ==========================================================================

func Test_PC_SkipDynamic(t *testing.T) {
	// Arrange
	pc := makePC3(t)
	result := pc.SkipDynamic(1)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SkipDynamic returns correct value -- with args", actual)
}

func Test_PC_SkipCollection(t *testing.T) {
	// Arrange
	pc := makePC3(t)
	result := pc.SkipCollection(1)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SkipCollection returns correct value -- with args", actual)
}

func Test_PC_TakeDynamic(t *testing.T) {
	// Arrange
	pc := makePC3(t)
	result := pc.TakeDynamic(2)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TakeDynamic returns correct value -- with args", actual)
}

func Test_PC_TakeCollection(t *testing.T) {
	// Arrange
	pc := makePC3(t)
	result := pc.TakeCollection(2)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TakeCollection returns correct value -- with args", actual)
}

func Test_PC_LimitCollection(t *testing.T) {
	// Arrange
	pc := makePC3(t)
	result := pc.LimitCollection(1)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LimitCollection returns correct value -- with args", actual)
}

func Test_PC_LimitDynamic(t *testing.T) {
	// Arrange
	pc := makePC3(t)
	result := pc.LimitDynamic(2)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LimitDynamic returns correct value -- with args", actual)
}

func Test_PC_Limit(t *testing.T) {
	// Arrange
	pc := makePC3(t)
	result := pc.Limit(2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Limit returns correct value -- with args", actual)
}

// ==========================================================================
// PayloadsCollectionGetters — IsEqualItems
// ==========================================================================

func Test_PC_IsEqualItems_Same(t *testing.T) {
	// Arrange
	pc := makePC3(t)

	// Act
	actual := args.Map{"val": pc.IsEqualItems(pc.Items[0], pc.Items[1], pc.Items[2])}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEqualItems returns correct value -- same", actual)
}

func Test_PC_IsEqualItems_DiffLen(t *testing.T) {
	// Arrange
	pc := makePC3(t)

	// Act
	actual := args.Map{"val": pc.IsEqualItems(pc.Items[0])}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsEqualItems returns correct value -- diff len", actual)
}

func Test_PC_IsEqualItems_NilPC(t *testing.T) {
	// Arrange
	var pc *corepayload.PayloadsCollection

	// Act
	actual := args.Map{"val": pc.IsEqualItems(nil)}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsEqualItems returns false -- nil pc variadic nil", actual)
}

// ==========================================================================
// PayloadsCollectionPaging — GetPagedCollection, GetSinglePageCollection
// ==========================================================================

func Test_PC_GetPagedCollection(t *testing.T) {
	// Arrange
	pc := makePC3(t)
	pages := pc.GetPagedCollection(2)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 2}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection returns correct value -- with args", actual)
}

func Test_PC_GetPagedCollection_SmallEnough(t *testing.T) {
	// Arrange
	pc := makePC3(t)
	pages := pc.GetPagedCollection(10)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection returns correct value -- small", actual)
}

func Test_PC_GetSinglePageCollection(t *testing.T) {
	// Arrange
	pc := makePC3(t)
	page := pc.GetSinglePageCollection(2, 2)

	// Act
	actual := args.Map{"len": page.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection returns correct value -- page 2", actual)
}

func Test_PC_GetSinglePageCollection_SmallEnough(t *testing.T) {
	// Arrange
	pc := makePC3(t)
	page := pc.GetSinglePageCollection(10, 1)

	// Act
	actual := args.Map{"len": page.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection returns correct value -- small", actual)
}

// ==========================================================================
// PayloadCreateInstructionTypeStringer — PayloadCreateInstruction()
// ==========================================================================

func Test_PayloadCreateInstructionTypeStringer_FromNewPWUsingBytesCreat(t *testing.T) {
	// Arrange
	pci := corepayload.PayloadCreateInstructionTypeStringer{
		Name: "n", Identifier: "id",
		TaskTypeNameStringer: stringer16{"task"},
		CategoryNameStringer: stringer16{"cat"},
		Payloads:             "data",
	}
	pi := pci.PayloadCreateInstruction()

	// Act
	actual := args.Map{
		"name": pi.Name,
		"task": pi.TaskTypeName,
		"cat": pi.CategoryName,
	}

	// Assert
	expected := args.Map{
		"name": "n",
		"task": "task",
		"cat": "cat",
	}
	expected.ShouldBeEqual(t, 0, "PayloadCreateInstructionTypeStringer returns correct value -- with args", actual)
}

// ==========================================================================
// PayloadTypeExpander — struct coverage
// ==========================================================================

func Test_PayloadTypeExpander(t *testing.T) {
	// Arrange
	pte := corepayload.PayloadTypeExpander{
		CategoryStringer: stringer16{"cat"},
		TaskTypeStringer: stringer16{"task"},
	}

	// Act
	actual := args.Map{
		"cat":  pte.CategoryStringer.String(),
		"task": pte.TaskTypeStringer.String(),
	}

	// Assert
	expected := args.Map{
		"cat": "cat",
		"task": "task",
	}
	expected.ShouldBeEqual(t, 0, "PayloadTypeExpander returns correct value -- with args", actual)
}

// ==========================================================================
// BytesCreateInstructionStringer — struct fields
// ==========================================================================

func Test_BytesCreateInstructionStringer(t *testing.T) {
	// Arrange
	bci := corepayload.BytesCreateInstructionStringer{
		Name: "n", Identifier: "id",
		TaskTypeName: stringer16{"task"},
		EntityType:   "entity",
		CategoryName: stringer16{"cat"},
		Payloads:     []byte("data"),
	}

	// Act
	actual := args.Map{
		"name": bci.Name,
		"task": bci.TaskTypeName.String(),
		"cat": fmt.Sprintf("%v", bci.CategoryName),
	}

	// Assert
	expected := args.Map{
		"name": "n",
		"task": "task",
		"cat": "cat",
	}
	expected.ShouldBeEqual(t, 0, "BytesCreateInstructionStringer returns correct value -- with args", actual)
}

// ==========================================================================
// CastOrDeserializeFrom — valid path
// ==========================================================================

func Test_NewPW_CastOrDeserializeFrom_Valid(t *testing.T) {
	// Arrange
	pw, _ := corepayload.New.PayloadWrapper.NameIdCategory("n", "id", "cat", "data")
	pw2, err := corepayload.New.PayloadWrapper.CastOrDeserializeFrom(pw)
	// CastOrDeserializeFrom uses corejson.CastAny.FromToDefault (JSON round-trip).
	// MarshalJSON preserves Name in payloadWrapperModel, UnmarshalJSON reads it back.
	// Name should survive the round-trip since both use the same model.

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": pw2 != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "CastOrDeserializeFrom returns non-nil -- valid", actual)
}

// ==========================================================================
// DeserializeToMany — valid path for PayloadWrapper
// ==========================================================================

func Test_NewPW_DeserializeToMany_Valid(t *testing.T) {
	// Arrange
	pw, _ := corepayload.New.PayloadWrapper.NameIdCategory("n", "id", "cat", "data")
	b, _ := corejson.Serialize.Raw([]*corepayload.PayloadWrapper{pw})
	result, err := corepayload.New.PayloadWrapper.DeserializeToMany(b)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(result),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "NewPW.DeserializeToMany returns non-empty -- valid", actual)
}

// ==========================================================================
// DeserializeToCollection
// ==========================================================================

func Test_NewPW_DeserializeToCollection(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Empty()
	b, _ := corejson.Serialize.Raw(pc)
	result, err := corepayload.New.PayloadWrapper.DeserializeToCollection(b)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeToCollection returns correct value -- with args", actual)
}

// ==========================================================================
// emptyCreator — remaining methods
// ==========================================================================

func Test_Empty_PayloadWrapper_FromNewPWUsingBytesCreat(t *testing.T) {
	// Arrange
	pw := corepayload.Empty.PayloadWrapper()

	// Act
	actual := args.Map{"notNil": pw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.PayloadWrapper returns empty -- with args", actual)
}

func Test_Empty_PayloadsCollection_FromNewPWUsingBytesCreat(t *testing.T) {
	// Arrange
	pc := corepayload.Empty.PayloadsCollection()

	// Act
	actual := args.Map{"len": pc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty.PayloadsCollection returns empty -- with args", actual)
}
