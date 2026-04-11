package coredynamictests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// New.Collection creators (newCreator → newCollectionCreator → generic/typed)
// =============================================================================

func Test_New_Collection_String_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Empty", actual)
}

func Test_New_Collection_String_Cap(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Cap(10)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Cap", actual)
}

func Test_New_Collection_String_From(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.From", actual)
}

func Test_New_Collection_String_Clone(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Clone([]string{"a"})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Clone", actual)
}

func Test_New_Collection_String_Items(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Items("a", "b", "c")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Items", actual)
}

func Test_New_Collection_String_Create(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Create([]string{"x"})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Create", actual)
}

func Test_New_Collection_String_LenCap(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.LenCap(3, 10)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.LenCap", actual)
}

func Test_New_Collection_Int_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.Empty", actual)
}

func Test_New_Collection_Int_LenCap(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int.LenCap(5, 10)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.LenCap", actual)
}

func Test_New_Collection_Int64_LenCap(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int64.LenCap(2, 8)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int64.LenCap", actual)
}

func Test_New_Collection_Byte_LenCap(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Byte.LenCap(4, 16)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "New.Collection.Byte.LenCap", actual)
}

func Test_New_Collection_Any_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Any.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Any.Empty", actual)
}

func Test_New_Collection_Any_Cap(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Any.Cap(5)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Any.Cap", actual)
}

func Test_New_Collection_Bool_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Bool.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Bool.Empty", actual)
}

func Test_New_Collection_Float32_From(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Float32.From([]float32{1.5, 2.5})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Float32.From", actual)
}

func Test_New_Collection_Float64_Items(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Float64.Items(1.0, 2.0)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Float64.Items", actual)
}

func Test_New_Collection_ByteSlice_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.ByteSlice.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.ByteSlice.Empty", actual)
}

func Test_New_Collection_AnyMap_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.AnyMap.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.AnyMap.Empty", actual)
}

func Test_New_Collection_StringMap_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.StringMap.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.StringMap.Empty", actual)
}

func Test_New_Collection_IntMap_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.IntMap.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.IntMap.Empty", actual)
}

// =============================================================================
// BytesConverter — remaining branches
// =============================================================================

func Test_BytesConverter_ToBool_Valid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("true"))
	r, err := bc.ToBool()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"r": r,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"r": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBool valid", actual)
}

func Test_BytesConverter_ToBoolMust_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("false"))

	// Act
	actual := args.Map{"r": bc.ToBoolMust()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBoolMust", actual)
}

func Test_BytesConverter_SafeCastString_Empty_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte{})

	// Act
	actual := args.Map{"r": bc.SafeCastString()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "BytesConverter SafeCastString empty", actual)
}

func Test_BytesConverter_SafeCastString_Valid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("hello"))

	// Act
	actual := args.Map{"r": bc.SafeCastString()}

	// Assert
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesConverter SafeCastString valid", actual)
}

func Test_BytesConverter_CastString_Empty_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte{})
	_, err := bc.CastString()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter CastString empty", actual)
}

func Test_BytesConverter_CastString_Valid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	s, err := bc.CastString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"r": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"r": "hello",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter CastString valid", actual)
}

func Test_BytesConverter_ToString_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	s, err := bc.ToString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"r": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"r": "hello",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToString", actual)
}

func Test_BytesConverter_ToStringMust_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"world"`))

	// Act
	actual := args.Map{"r": bc.ToStringMust()}

	// Assert
	expected := args.Map{"r": "world"}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStringMust", actual)
}

func Test_BytesConverter_ToStrings_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	ss, err := bc.ToStrings()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(ss),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStrings", actual)
}

func Test_BytesConverter_ToStringsMust_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`["x"]`))
	ss := bc.ToStringsMust()

	// Act
	actual := args.Map{"len": len(ss)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStringsMust", actual)
}

func Test_BytesConverter_ToInt64_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("42"))
	v, err := bc.ToInt64()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"v": v,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"v": int64(42),
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToInt64", actual)
}

func Test_BytesConverter_ToInt64Must_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("99"))

	// Act
	actual := args.Map{"v": bc.ToInt64Must()}

	// Assert
	expected := args.Map{"v": int64(99)}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToInt64Must", actual)
}

func Test_BytesConverter_ToHashmap_Valid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`{"key":"val"}`))
	h, err := bc.ToHashmap()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": h != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashmap valid", actual)
}

func Test_BytesConverter_ToHashmap_Invalid_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToHashmap()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashmap invalid", actual)
}

func Test_BytesConverter_ToHashset_Valid(t *testing.T) {
	// Arrange
	b, _ := json.Marshal(map[string]bool{"a": true, "b": true})
	bc := coredynamic.NewBytesConverter(b)
	h, err := bc.ToHashset()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": h != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashset valid", actual)
}

func Test_BytesConverter_ToHashset_Invalid_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToHashset()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashset invalid", actual)
}

func Test_BytesConverter_ToCollection_Valid(t *testing.T) {
	// Arrange
	b, _ := json.Marshal([]string{"a"})
	bc := coredynamic.NewBytesConverter(b)
	c, err := bc.ToCollection()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": c != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToCollection valid", actual)
}

func Test_BytesConverter_ToCollection_Invalid_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToCollection invalid", actual)
}

func Test_BytesConverter_ToSimpleSlice_Valid(t *testing.T) {
	// Arrange
	b, _ := json.Marshal([]string{"a"})
	bc := coredynamic.NewBytesConverter(b)
	s, err := bc.ToSimpleSlice()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": s != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToSimpleSlice valid", actual)
}

func Test_BytesConverter_ToSimpleSlice_Invalid_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToSimpleSlice()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToSimpleSlice invalid", actual)
}

func Test_BytesConverter_ToKeyValCollection_Valid(t *testing.T) {
	// Arrange
	b, _ := json.Marshal(map[string]any{"Items": []map[string]any{{"Key": "a", "Value": 1}}})
	bc := coredynamic.NewBytesConverter(b)
	c, err := bc.ToKeyValCollection()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": c != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToKeyValCollection valid", actual)
}

func Test_BytesConverter_ToKeyValCollection_Invalid_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToKeyValCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToKeyValCollection invalid", actual)
}

func Test_BytesConverter_ToAnyCollection_Invalid_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToAnyCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToAnyCollection invalid", actual)
}

func Test_BytesConverter_ToMapAnyItems_Invalid_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToMapAnyItems()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToMapAnyItems invalid", actual)
}

func Test_BytesConverter_ToDynamicCollection_Invalid_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToDynamicCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToDynamicCollection invalid", actual)
}

func Test_BytesConverter_ToJsonResultCollection_Invalid_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToJsonResultCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToJsonResultCollection invalid", actual)
}

func Test_BytesConverter_ToJsonMapResults_Invalid_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToJsonMapResults()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToJsonMapResults invalid", actual)
}

func Test_BytesConverter_ToBytesCollection_Invalid_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToBytesCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBytesCollection invalid", actual)
}

func Test_BytesConverter_Deserialize_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	var s string
	err := bc.Deserialize(&s)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"r": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"r": "hello",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter Deserialize", actual)
}

func Test_BytesConverter_DeserializeMust_FromNewCollectionNewCrea(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"world"`))
	var s string
	bc.DeserializeMust(&s)

	// Act
	actual := args.Map{"r": s}

	// Assert
	expected := args.Map{"r": "world"}
	expected.ShouldBeEqual(t, 0, "BytesConverter DeserializeMust", actual)
}

func Test_NewBytesConverterUsingJsonResult_Valid(t *testing.T) {
	// Arrange
	r := corejson.New("test")
	bc, err := coredynamic.NewBytesConverterUsingJsonResult(r.Ptr())

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": bc != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "NewBytesConverterUsingJsonResult valid", actual)
}

func Test_NewBytesConverterUsingJsonResult_Invalid(t *testing.T) {
	// Arrange
	r := corejson.Result{}
	_, err := coredynamic.NewBytesConverterUsingJsonResult(&r)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewBytesConverterUsingJsonResult invalid", actual)
}

// =============================================================================
// DynamicCollectionModel
// =============================================================================

func Test_DynamicCollectionModel(t *testing.T) {
	// Arrange
	dcm := coredynamic.DynamicCollectionModel{
		Items: []coredynamic.Dynamic{
			*coredynamic.NewDynamicPtr("a", true),
		},
	}

	// Act
	actual := args.Map{"len": len(dcm.Items)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollectionModel", actual)
}
