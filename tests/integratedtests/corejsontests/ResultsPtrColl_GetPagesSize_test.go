package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ResultsPtrCollection: Paging ──

func Test_ResultsPtrColl_GetPagesSize(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a", "b", "c")

	// Act
	pages := coll.GetPagesSize(2)

	// Assert
	actual := args.Map{"pages": pages}
	expected := args.Map{"pages": 2}
	expected.ShouldBeEqual(t, 0, "GetPagesSize returns 2 -- 3 items / 2", actual)
}

func Test_ResultsPtrColl_GetPagesSize_ZeroPageSize(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a")

	// Act
	pages := coll.GetPagesSize(0)

	// Assert
	actual := args.Map{"pages": pages}
	expected := args.Map{"pages": 0}
	expected.ShouldBeEqual(t, 0, "GetPagesSize returns 0 -- zero page size", actual)
}

func Test_ResultsPtrColl_GetPagedCollection(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a", "b", "c", "d", "e")

	// Act
	pages := coll.GetPagedCollection(2)

	// Assert
	actual := args.Map{
		"pageCount":  len(pages),
		"page1Items": pages[0].Length(),
		"page3Items": pages[2].Length(),
	}
	expected := args.Map{
		"pageCount":  3,
		"page1Items": 2,
		"page3Items": 1,
	}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection returns 3 pages -- 5/2", actual)
}

func Test_ResultsPtrColl_GetPagedCollection_SmallCollection(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a")

	// Act
	pages := coll.GetPagedCollection(10)

	// Assert
	actual := args.Map{"pageCount": len(pages)}
	expected := args.Map{"pageCount": 1}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection returns 1 page -- smaller than page size", actual)
}

// ── ResultsPtrCollection: Clone ──

func Test_ResultsPtrColl_Clone_DeepClone(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a", "b")

	// Act
	cloned := coll.Clone(true)

	// Assert
	actual := args.Map{"length": cloned.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Clone returns deep clone -- 2 items", actual)
}

func Test_ResultsPtrColl_Clone_Nil(t *testing.T) {
	// Arrange
	var coll *corejson.ResultsPtrCollection

	// Act
	cloned := coll.Clone(true)

	// Assert
	actual := args.Map{"isNil": cloned == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Clone returns nil -- nil receiver", actual)
}

// ── ResultsPtrCollection: AddNonNilNonError ──

func Test_ResultsPtrColl_AddNonNilNonError(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.Empty()
	valid := corejson.Serialize.Apply("valid")
	errResult := corejson.Empty.ResultPtrWithErr("err", nil)

	// Act
	coll.AddNonNilNonError(valid)
	coll.AddNonNilNonError(nil)
	coll.AddNonNilNonError(errResult)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "AddNonNilNonError adds valid -- skips nil", actual)
}

// ── ResultsPtrCollection: AddResult (non-ptr) ──

func Test_ResultsPtrColl_AddResult(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New("hello")

	// Act
	coll.AddResult(r)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "AddResult adds non-ptr -- 1 item", actual)
}

// ── ResultsPtrCollection: Adds ──

func Test_ResultsPtrColl_Adds(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.Empty()
	r1 := corejson.Serialize.Apply("a")
	r2 := corejson.Serialize.Apply("b")

	// Act
	coll.Adds(r1, r2)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Adds appends multiple -- 2 items", actual)
}

// ── ResultsPtrCollection: Clear / Dispose ──

func Test_ResultsPtrColl_Clear(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a", "b")

	// Act
	coll.Clear()

	// Assert
	actual := args.Map{"isEmpty": coll.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Clear empties collection -- empty", actual)
}

func Test_ResultsPtrColl_Dispose(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a")

	// Act
	coll.Dispose()

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "Dispose nils items -- 0 length", actual)
}

// ── ResultsPtrCollection: GetStrings ──

func Test_ResultsPtrColl_GetStrings(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("hello", 42)

	// Act
	result := coll.GetStrings()

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "GetStrings returns string list -- 2 items", actual)
}

// ── ResultsPtrCollection: ParseInjectUsingJson ──

func Test_ResultsPtrColl_ParseInjectUsingJson_Valid(t *testing.T) {
	// Arrange
	original := corejson.NewResultsPtrCollection.AnyItems("a", "b")
	jsonResult := original.JsonPtr()
	target := corejson.NewResultsPtrCollection.Empty()

	// Act
	result, err := target.ParseInjectUsingJson(jsonResult)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"hasItems": result.Length() > 0,
	}
	expected := args.Map{
		"noErr": true,
		"hasItems": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson parses -- valid", actual)
}

// ── ResultsPtrCollection: GetErrorsAsSingle ──

func Test_ResultsPtrColl_GetErrorsAsSingle(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a")

	// Act
	result := coll.GetErrorsAsSingle()

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetErrorsAsSingle returns error -- always", actual)
}

// ── ResultsPtrCollection: Json / JsonPtr / AsJsoner ──

func Test_ResultsPtrColl_Json(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a")

	// Act
	result := coll.Json()

	// Assert
	actual := args.Map{"hasBytes": result.Length() > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Json returns result -- serialized", actual)
}

// ── ResultsPtrCollection: NonPtr / Ptr ──

func Test_ResultsPtrColl_NonPtrPtr(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a")

	// Act
	nonPtr := coll.NonPtr()
	ptr := coll.Ptr()

	// Assert
	actual := args.Map{
		"nonPtrLen": nonPtr.Length(),
		"ptrLen": ptr.Length(),
	}
	expected := args.Map{
		"nonPtrLen": 1,
		"ptrLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "NonPtr/Ptr return same -- 1 item", actual)
}

// ── ResultsPtrCollection: AsJsonContractsBinder / AsJsonParseSelfInjector ──

func Test_ResultsPtrColl_Interfaces(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a")

	// Act
	binder := coll.AsJsonContractsBinder()
	jsoner := coll.AsJsoner()
	injector := coll.AsJsonParseSelfInjector()

	// Assert
	actual := args.Map{
		"binderNotNil":   binder != nil,
		"jsonerNotNil":   jsoner != nil,
		"injectorNotNil": injector != nil,
	}
	expected := args.Map{
		"binderNotNil":   true,
		"jsonerNotNil":   true,
		"injectorNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Interface accessors return non-nil -- all", actual)
}

// ── newResultsPtrCollectionCreator ──

func Test_NewResultsPtrColl_Creators(t *testing.T) {
	// Arrange & Act
	empty := corejson.NewResultsPtrCollection.Empty()
	def := corejson.NewResultsPtrCollection.Default()
	withCap := corejson.NewResultsPtrCollection.UsingCap(5)
	anyItems := corejson.NewResultsPtrCollection.AnyItemsPlusCap(2, "a")
	emptyAnyItems := corejson.NewResultsPtrCollection.AnyItemsPlusCap(2)

	// Assert
	actual := args.Map{
		"emptyLen":     empty.Length(),
		"defLen":       def.Length(),
		"withCapLen":   withCap.Length(),
		"anyItemsLen":  anyItems.Length(),
		"emptyAnyLen":  emptyAnyItems.Length(),
	}
	expected := args.Map{
		"emptyLen":     0,
		"defLen":       0,
		"withCapLen":   0,
		"anyItemsLen":  1,
		"emptyAnyLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "Creators return correct -- various", actual)
}

func Test_NewResultsPtrColl_UsingResults(t *testing.T) {
	// Arrange
	r1 := corejson.Serialize.Apply("a")
	r2 := corejson.Serialize.Apply("b")

	// Act
	coll := corejson.NewResultsPtrCollection.UsingResults(r1, r2)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "UsingResults returns 2 -- from results", actual)
}

func Test_NewResultsPtrColl_DeserializeUsingResult(t *testing.T) {
	// Arrange
	original := corejson.NewResultsPtrCollection.AnyItems("a", "b")
	jsonResult := original.JsonPtr()

	// Act
	result, err := corejson.NewResultsPtrCollection.DeserializeUsingResult(jsonResult)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"hasItems": result != nil && result.Length() > 0,
	}
	expected := args.Map{
		"noErr": true,
		"hasItems": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeUsingResult returns coll -- valid", actual)
}
