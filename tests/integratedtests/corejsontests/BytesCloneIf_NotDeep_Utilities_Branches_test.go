package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// BytesCloneIf
// =============================================================================

func Test_BytesCloneIf_NotDeep(t *testing.T) {
	// Arrange
	b := []byte("hello")
	c := corejson.BytesCloneIf(false, b)

	// Act
	actual := args.Map{"len": len(c)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf not deep", actual)
}

func Test_BytesCloneIf_DeepEmpty(t *testing.T) {
	// Arrange
	c := corejson.BytesCloneIf(true, []byte{})

	// Act
	actual := args.Map{"len": len(c)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf deep empty", actual)
}

func Test_BytesCloneIf_DeepValid(t *testing.T) {
	// Arrange
	b := []byte("hello")
	c := corejson.BytesCloneIf(true, b)

	// Act
	actual := args.Map{
		"len": len(c),
		"same": &c[0] != &b[0],
	}

	// Assert
	expected := args.Map{
		"len": 5,
		"same": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf deep valid", actual)
}

// =============================================================================
// BytesDeepClone
// =============================================================================

func Test_BytesDeepClone_Empty(t *testing.T) {
	// Arrange
	c := corejson.BytesDeepClone([]byte{})

	// Act
	actual := args.Map{"len": len(c)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone empty", actual)
}

func Test_BytesDeepClone_Valid(t *testing.T) {
	// Arrange
	b := []byte("hello")
	c := corejson.BytesDeepClone(b)

	// Act
	actual := args.Map{"len": len(c)}

	// Assert
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone valid", actual)
}

// =============================================================================
// BytesToString / BytesToPrettyString
// =============================================================================

func Test_BytesToString_Empty(t *testing.T) {
	// Act
	actual := args.Map{"r": corejson.BytesToString([]byte{})}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "BytesToString empty", actual)
}

func Test_BytesToString_Valid(t *testing.T) {
	// Act
	actual := args.Map{"r": corejson.BytesToString([]byte("hello"))}

	// Assert
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesToString valid", actual)
}

func Test_BytesToPrettyString_Empty(t *testing.T) {
	// Act
	actual := args.Map{"r": corejson.BytesToPrettyString([]byte{})}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString empty", actual)
}

func Test_BytesToPrettyString_Valid(t *testing.T) {
	// Arrange
	b := []byte(`{"a":1}`)
	s := corejson.BytesToPrettyString(b)

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString valid", actual)
}

// =============================================================================
// JsonStringOrErrMsg
// =============================================================================

func Test_JsonStringOrErrMsg_Valid(t *testing.T) {
	// Arrange
	s := corejson.JsonStringOrErrMsg("hello")

	// Act
	actual := args.Map{"r": s}

	// Assert
	expected := args.Map{"r": `"hello"`}
	expected.ShouldBeEqual(t, 0, "JsonStringOrErrMsg valid", actual)
}

func Test_JsonStringOrErrMsg_Unmarshalable(t *testing.T) {
	// Arrange
	s := corejson.JsonStringOrErrMsg(make(chan int))

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "JsonStringOrErrMsg unmarshalable", actual)
}

// =============================================================================
// New / NewPtr
// =============================================================================

func Test_New_Valid(t *testing.T) {
	// Arrange
	r := corejson.New("hello")

	// Act
	actual := args.Map{
		"noErr": !r.HasError(),
		"hasBytes": r.Length() > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "New valid", actual)
}

func Test_New_Unmarshalable(t *testing.T) {
	// Arrange
	r := corejson.New(make(chan int))

	// Act
	actual := args.Map{"hasErr": r.HasError()}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "New unmarshalable", actual)
}

func Test_NewPtr_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"noErr": !r.HasError(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "NewPtr valid", actual)
}

func Test_NewPtr_Unmarshalable(t *testing.T) {
	// Arrange
	r := corejson.NewPtr(make(chan int))

	// Act
	actual := args.Map{"hasErr": r.HasError()}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewPtr unmarshalable", actual)
}

// =============================================================================
// emptyCreator
// =============================================================================

func Test_Empty_Result(t *testing.T) {
	// Arrange
	r := corejson.Empty.Result()

	// Act
	actual := args.Map{"empty": r.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Result", actual)
}

func Test_Empty_ResultPtr(t *testing.T) {
	// Arrange
	r := corejson.Empty.ResultPtr()

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"empty": r.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "Empty.ResultPtr", actual)
}

func Test_Empty_ResultWithErr(t *testing.T) {
	// Arrange
	r := corejson.Empty.ResultWithErr("int", nil)

	// Act
	actual := args.Map{"typeName": r.TypeName}

	// Assert
	expected := args.Map{"typeName": "int"}
	expected.ShouldBeEqual(t, 0, "Empty.ResultWithErr", actual)
}

func Test_Empty_ResultPtrWithErr(t *testing.T) {
	// Arrange
	r := corejson.Empty.ResultPtrWithErr("int", nil)

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"typeName": r.TypeName,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"typeName": "int",
	}
	expected.ShouldBeEqual(t, 0, "Empty.ResultPtrWithErr", actual)
}

func Test_Empty_BytesCollection(t *testing.T) {
	// Arrange
	bc := corejson.Empty.BytesCollection()

	// Act
	actual := args.Map{"len": bc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty.BytesCollection", actual)
}

func Test_Empty_BytesCollectionPtr(t *testing.T) {
	// Arrange
	bc := corejson.Empty.BytesCollectionPtr()

	// Act
	actual := args.Map{"notNil": bc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.BytesCollectionPtr", actual)
}

func Test_Empty_ResultsCollection(t *testing.T) {
	// Arrange
	rc := corejson.Empty.ResultsCollection()

	// Act
	actual := args.Map{
		"notNil": rc != nil,
		"len": rc.Length(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "Empty.ResultsCollection", actual)
}

func Test_Empty_ResultsPtrCollection(t *testing.T) {
	// Arrange
	rc := corejson.Empty.ResultsPtrCollection()

	// Act
	actual := args.Map{"notNil": rc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultsPtrCollection", actual)
}

func Test_Empty_MapResults(t *testing.T) {
	// Arrange
	mr := corejson.Empty.MapResults()

	// Act
	actual := args.Map{
		"notNil": mr != nil,
		"len": mr.Length(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "Empty.MapResults", actual)
}

// =============================================================================
// KeyAny / KeyWithResult
// =============================================================================

func Test_KeyAny(t *testing.T) {
	// Arrange
	ka := corejson.KeyAny{Key: "x", AnyInf: 42}

	// Act
	actual := args.Map{
		"key": ka.Key,
		"val": ka.AnyInf,
	}

	// Assert
	expected := args.Map{
		"key": "x",
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "KeyAny", actual)
}

func Test_KeyWithResult(t *testing.T) {
	// Arrange
	kwr := corejson.KeyWithResult{Key: "x", Result: corejson.New("hello")}

	// Act
	actual := args.Map{
		"key": kwr.Key,
		"noErr": !kwr.Result.HasError(),
	}

	// Assert
	expected := args.Map{
		"key": "x",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyWithResult", actual)
}
