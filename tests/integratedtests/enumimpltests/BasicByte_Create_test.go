package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BasicByte via newBasicByteCreator ──

func Test_BasicByte_Create_Default(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active", "Inactive"},
	)

	// Act
	actual := args.Map{
		"min":          bb.Min(),
		"max":          bb.Max(),
		"isValidRange": bb.IsValidRange(1),
		"outOfRange":   bb.IsValidRange(5),
		"toString":     bb.ToEnumString(1),
		"typeName":     bb.TypeName(),
		"length":       bb.Length(),
		"count":        bb.Count(),
	}

	// Assert
	expected := args.Map{
		"min":          byte(0),
		"max":          byte(2),
		"isValidRange": true,
		"outOfRange":   false,
		"toString":     "Active",
		"typeName":     "enumimpltests.myEnum",
		"length":       3,
		"count":        3,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_Create_Default returns correct value -- with args", actual)
}

func Test_BasicByte_IsAnyOf(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active", "Inactive"},
	)

	// Act
	actual := args.Map{
		"isAnyEmpty":   bb.IsAnyOf(1),
		"isAnyMatch":  bb.IsAnyOf(1, 0, 1, 2),
		"isAnyNoMatch": bb.IsAnyOf(1, 0, 2),
	}

	// Assert
	expected := args.Map{
		"isAnyEmpty":   true,
		"isAnyMatch":  true,
		"isAnyNoMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_IsAnyOf returns correct value -- with args", actual)
}

func Test_BasicByte_IsAnyNamesOf(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active", "Inactive"},
	)

	// Act
	actual := args.Map{
		"matchName":   bb.IsAnyNamesOf(1, "Active"),
		"noMatchName": bb.IsAnyNamesOf(1, "Invalid"),
	}

	// Assert
	expected := args.Map{
		"matchName":   true,
		"noMatchName": false,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_IsAnyNamesOf returns correct value -- with args", actual)
}

func Test_BasicByte_GetValueByString(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active", "Inactive"},
	)

	// Act
	actual := args.Map{
		"byName": bb.GetValueByString("Active"),
	}

	// Assert
	expected := args.Map{
		"byName": byte(1),
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_GetValueByString returns correct value -- with args", actual)
}

func Test_BasicByte_GetValueByName(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	val, err := bb.GetValueByName("Active")
	_, errNotFound := bb.GetValueByName("NotExist")

	// Act
	actual := args.Map{
		"val":      val,
		"noErr":    err == nil,
		"hasError": errNotFound != nil,
	}

	// Assert
	expected := args.Map{
		"val":      byte(1),
		"noErr":    true,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_GetValueByName returns correct value -- with args", actual)
}

func Test_BasicByte_GetStringValue(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"val": bb.GetStringValue(0),
	}

	// Assert
	expected := args.Map{
		"val": "Invalid",
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_GetStringValue returns correct value -- with args", actual)
}

func Test_BasicByte_Ranges(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"rangesLen": len(bb.Ranges()),
	}

	// Assert
	expected := args.Map{
		"rangesLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_Ranges returns correct value -- with args", actual)
}

func Test_BasicByte_Hashmap(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	hm := bb.Hashmap()
	hmPtr := bb.HashmapPtr()

	// Act
	actual := args.Map{
		"hmHasItems":    len(hm) > 0,
		"hmPtrNotNil":   hmPtr != nil,
	}

	// Assert
	expected := args.Map{
		"hmHasItems":    true,
		"hmPtrNotNil":   true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_Hashmap returns correct value -- with args", actual)
}

func Test_BasicByte_ToEnumJsonBytes(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	jsonBytes, err := bb.ToEnumJsonBytes(0)
	_, errNotFound := bb.ToEnumJsonBytes(99)

	// Act
	actual := args.Map{
		"hasBytes":  len(jsonBytes) > 0,
		"noErr":     err == nil,
		"notFound":  errNotFound != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes":  true,
		"noErr":     true,
		"notFound":  true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_ToEnumJsonBytes returns correct value -- with args", actual)
}

func Test_BasicByte_AppendPrependJoinValue(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	result := bb.AppendPrependJoinValue(".", 1, 0)

	// Act
	actual := args.Map{
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_AppendPrependJoinValue returns correct value -- with args", actual)
}

func Test_BasicByte_ToNumberString(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"notEmpty": bb.ToNumberString(1) != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_ToNumberString returns correct value -- with args", actual)
}

func Test_BasicByte_UnmarshallToValue(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	// nil with mapping to first
	val1, err1 := bb.UnmarshallToValue(true, nil)
	// nil without mapping
	_, err2 := bb.UnmarshallToValue(false, nil)
	// empty string with mapping
	val3, err3 := bb.UnmarshallToValue(true, []byte(""))
	// valid name
	val4, err4 := bb.UnmarshallToValue(false, []byte("Active"))

	// Act
	actual := args.Map{
		"nilMapped":    val1,
		"nilMappedErr": err1 == nil,
		"nilNoMap":     err2 != nil,
		"emptyMapped":  val3,
		"emptyNoErr":   err3 == nil,
		"validVal":     val4,
		"validNoErr":   err4 == nil,
	}

	// Assert
	expected := args.Map{
		"nilMapped":    byte(0),
		"nilMappedErr": true,
		"nilNoMap":     true,
		"emptyMapped":  byte(0),
		"emptyNoErr":   true,
		"validVal":     byte(1),
		"validNoErr":   true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_UnmarshallToValue returns correct value -- with args", actual)
}

func Test_BasicByte_EnumType(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"enumType": bb.EnumType().String(),
	}

	// Assert
	expected := args.Map{
		"enumType": "Byte",
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_EnumType returns correct value -- with args", actual)
}

func Test_BasicByte_AsBasicByter(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	byter := bb.AsBasicByter()

	// Act
	actual := args.Map{
		"notNil": byter != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_AsBasicByter returns correct value -- with args", actual)
}

// ── BasicByte with alias map ──

func Test_BasicByte_WithAliasMap(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.DefaultWithAliasMap(
		myEnum(0),
		[]string{"Invalid", "Active"},
		map[string]byte{"on": 1},
	)

	val, err := bb.GetValueByName("on")

	// Act
	actual := args.Map{
		"aliasVal": val,
		"noErr":    err == nil,
	}

	// Assert
	expected := args.Map{
		"aliasVal": byte(1),
		"noErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_WithAliasMap returns non-empty -- with args", actual)
}

func Test_BasicByte_AllCases(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.DefaultAllCases(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	val, err := bb.GetValueByName("active")

	// Act
	actual := args.Map{
		"lowerVal": val,
		"noErr":    err == nil,
	}

	// Assert
	expected := args.Map{
		"lowerVal": byte(1),
		"noErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_AllCases returns correct value -- with args", actual)
}

// ── numberEnumBase methods via BasicByte ──

func Test_NumberEnumBase_MinMaxAny(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	min, max := bb.MinMaxAny()

	// Act
	actual := args.Map{
		"min": min,
		"max": max,
	}

	// Assert
	expected := args.Map{
		"min": byte(0),
		"max": byte(1),
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_MinMaxAny returns correct value -- with args", actual)
}

func Test_NumberEnumBase_ValueStrings(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"minValStr":     bb.MinValueString() != "",
		"maxValStr":     bb.MaxValueString() != "",
		"minInt":        bb.MinInt(),
		"maxInt":        bb.MaxInt(),
		"rangesCsv":     bb.RangeNamesCsv() != "",
		"rangesInvalid": bb.RangesInvalidMessage() != "",
		"rangesErr":     bb.RangesInvalidErr() != nil,
	}

	// Assert
	expected := args.Map{
		"minValStr":     true,
		"maxValStr":     true,
		"minInt":        0,
		"maxInt":        1,
		"rangesCsv":     true,
		"rangesInvalid": true,
		"rangesErr":     true,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_ValueStrings returns non-empty -- with args", actual)
}

func Test_NumberEnumBase_StringRanges(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"strRangesLen":    len(bb.StringRanges()),
		"strRangesPtrLen": len(bb.StringRangesPtr()),
		"namesHashLen":    len(bb.NamesHashset()),
	}

	// Assert
	expected := args.Map{
		"strRangesLen":    2,
		"strRangesPtrLen": 2,
		"namesHashLen":    2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_StringRanges returns correct value -- with args", actual)
}

func Test_NumberEnumBase_DynamicMap(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	dm := bb.RangesDynamicMap()
	dm2 := bb.DynamicMap()

	// Act
	actual := args.Map{
		"dmLen":  len(dm),
		"dm2Len": len(dm2),
	}

	// Assert
	expected := args.Map{
		"dmLen":  2,
		"dm2Len": 2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_DynamicMap returns correct value -- with args", actual)
}

func Test_NumberEnumBase_IntegerEnumRanges(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	ranges := bb.IntegerEnumRanges()

	// Act
	actual := args.Map{
		"len": len(ranges),
	}

	// Assert
	expected := args.Map{
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_IntegerEnumRanges returns correct value -- with args", actual)
}

func Test_NumberEnumBase_AllNameValues(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	anv := bb.AllNameValues()

	// Act
	actual := args.Map{
		"len": len(anv),
	}

	// Assert
	expected := args.Map{
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_AllNameValues returns non-empty -- with args", actual)
}

func Test_NumberEnumBase_KeyAnyValues(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	kav := bb.KeyAnyValues()
	kvi := bb.KeyValIntegers()

	// Act
	actual := args.Map{
		"kavLen": len(kav),
		"kviLen": len(kvi),
	}

	// Assert
	expected := args.Map{
		"kavLen": 2,
		"kviLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_KeyAnyValues returns non-empty -- with args", actual)
}

func Test_NumberEnumBase_Loop(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	count := 0
	bb.Loop(func(index int, name string, anyVal any) (isBreak bool) {
		count++
		return false
	})

	intCount := 0
	bb.LoopInteger(func(index int, name string, anyVal int) (isBreak bool) {
		intCount++
		return false
	})

	// Act
	actual := args.Map{
		"loopCount":    count,
		"intLoopCount": intCount,
	}

	// Assert
	expected := args.Map{
		"loopCount":    2,
		"intLoopCount": 2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_Loop returns correct value -- with args", actual)
}

func Test_NumberEnumBase_LoopBreak(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active", "Inactive"},
	)

	count := 0
	bb.Loop(func(index int, name string, anyVal any) (isBreak bool) {
		count++
		return true // break after first
	})

	// Act
	actual := args.Map{
		"breakCount": count,
	}

	// Assert
	expected := args.Map{
		"breakCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_LoopBreak returns correct value -- with args", actual)
}

func Test_NumberEnumBase_Format(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	result := bb.Format("Enum of {type-name} - {name} - {value}", byte(1))

	// Act
	actual := args.Map{
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_Format returns correct value -- with args", actual)
}

func Test_NumberEnumBase_NameWithValue(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	nv := bb.NameWithValue(byte(1))
	nvOpt := bb.NameWithValueOption(byte(1), true)
	nvOptNo := bb.NameWithValueOption(byte(1), false)

	// Act
	actual := args.Map{
		"nv":      nv != "",
		"nvOpt":   nvOpt != "",
		"nvOptNo": nvOptNo != "",
	}

	// Assert
	expected := args.Map{
		"nv":      true,
		"nvOpt":   true,
		"nvOptNo": true,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_NameWithValue returns non-empty -- with args", actual)
}

func Test_NumberEnumBase_RangesMap(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	rm := bb.RangesMap()
	rism := bb.RangesIntegerStringMap()

	// Act
	actual := args.Map{
		"rmLen":   len(rm),
		"rismLen": len(rism),
	}

	// Assert
	expected := args.Map{
		"rmLen":   2,
		"rismLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_RangesMap returns correct value -- with args", actual)
}

func Test_NumberEnumBase_OnlySupportedErr(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	noErr := bb.OnlySupportedErr("Invalid", "Active")
	hasErr := bb.OnlySupportedErr("Invalid")
	msgErr := bb.OnlySupportedMsgErr("context", "Invalid")

	// Act
	actual := args.Map{
		"noErr":  noErr == nil,
		"hasErr": hasErr != nil,
		"msgErr": msgErr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr":  true,
		"hasErr": true,
		"msgErr": true,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_OnlySupportedErr returns error -- with args", actual)
}

func Test_NumberEnumBase_JsonString(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	js := bb.JsonString(byte(1))
	es := bb.ToEnumString(byte(1))
	tn := bb.ToName(byte(1))

	// Act
	actual := args.Map{
		"js": js != "",
		"es": es != "",
		"tn": tn != "",
	}

	// Assert
	expected := args.Map{
		"js": true,
		"es": true,
		"tn": true,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_JsonString returns correct value -- with args", actual)
}

// ── BasicString ──

func Test_BasicString_Create(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active", "Inactive"},
	)

	// Act
	actual := args.Map{
		"min":           bs.Min(),
		"max":           bs.Max(),
		"length":        bs.Length(),
		"hasAnyItem":    bs.HasAnyItem(),
		"maxIndex":      bs.MaxIndex(),
		"isValidActive": bs.IsValidRange("Active"),
		"isValidBad":    bs.IsValidRange("NotExist"),
	}

	// Assert
	expected := args.Map{
		"min":           bs.Min(),
		"max":           bs.Max(),
		"length":        3,
		"hasAnyItem":    true,
		"maxIndex":      2,
		"isValidActive": true,
		"isValidBad":    false,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_Create returns correct value -- with args", actual)
}

func Test_BasicString_IsAnyOf(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"empty":   bs.IsAnyOf("x"),
		"match":   bs.IsAnyOf("Active", "Active", "Invalid"),
		"noMatch": bs.IsAnyOf("Active", "Invalid"),
	}

	// Assert
	expected := args.Map{
		"empty":   true,
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_IsAnyOf returns correct value -- with args", actual)
}

func Test_BasicString_IsAnyNamesOf(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"match":   bs.IsAnyNamesOf("Active", "Active"),
		"noMatch": bs.IsAnyNamesOf("Active", "Invalid"),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_IsAnyNamesOf returns correct value -- with args", actual)
}

func Test_BasicString_GetNameByIndex(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active", "Inactive"},
	)

	// Act
	actual := args.Map{
		"valid":   bs.GetNameByIndex(1),
		"invalid": bs.GetNameByIndex(99),
	}

	// Assert
	expected := args.Map{
		"valid":   "Active",
		"invalid": "",
	}
	expected.ShouldBeEqual(t, 0, "BasicString_GetNameByIndex returns correct value -- with args", actual)
}

func Test_BasicString_GetIndexByName(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"found":   bs.GetIndexByName("Active"),
		"empty":   bs.GetIndexByName(""),
		"missing": bs.GetIndexByName("NotExist"),
	}

	// Assert
	expected := args.Map{
		"found":   bs.GetIndexByName("Active"),
		"empty":   -1,
		"missing": -1,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_GetIndexByName returns correct value -- with args", actual)
}

func Test_BasicString_Ranges(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"rangesLen":   len(bs.Ranges()),
		"hashsetLen":  len(bs.Hashset()),
		"hashsetPtr":  bs.HashsetPtr() != nil,
		"integersLen": len(bs.RangesIntegers()),
		"nameIdxLen":  len(bs.NameWithIndexMap()),
	}

	// Assert
	expected := args.Map{
		"rangesLen":   2,
		"hashsetLen":  len(bs.Hashset()),
		"hashsetPtr":  true,
		"integersLen": 2,
		"nameIdxLen":  len(bs.NameWithIndexMap()),
	}
	expected.ShouldBeEqual(t, 0, "BasicString_Ranges returns correct value -- with args", actual)
}

func Test_BasicString_GetValueByName(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	val, err := bs.GetValueByName("Active")
	_, errNotFound := bs.GetValueByName("NotExist")

	// Act
	actual := args.Map{
		"val":      val,
		"noErr":    err == nil,
		"hasError": errNotFound != nil,
	}

	// Assert
	expected := args.Map{
		"val":      "Active",
		"noErr":    true,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_GetValueByName returns correct value -- with args", actual)
}

func Test_BasicString_ToEnumJsonBytes(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	jsonBytes, err := bs.ToEnumJsonBytes("Active")
	_, errBad := bs.ToEnumJsonBytes("NotExist")

	// Act
	actual := args.Map{
		"hasBytes":  len(jsonBytes) > 0,
		"noErr":     err == nil,
		"errOnBad":  errBad != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes":  true,
		"noErr":     true,
		"errOnBad":  true,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_ToEnumJsonBytes returns correct value -- with args", actual)
}

func Test_BasicString_UnmarshallToValue(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	val1, err1 := bs.UnmarshallToValue(true, nil)
	_, err2 := bs.UnmarshallToValue(false, nil)
	val3, err3 := bs.UnmarshallToValue(true, []byte(""))
	val4, err4 := bs.UnmarshallToValue(false, []byte("Active"))

	// Act
	actual := args.Map{
		"nilMapped":   val1,
		"nilMapErr":   err1 == nil,
		"nilNoMapErr": err2 != nil,
		"emptyVal":    val3,
		"emptyErr":    err3 == nil,
		"validVal":    val4,
		"validErr":    err4 == nil,
	}

	// Assert
	expected := args.Map{
		"nilMapped":   val1,
		"nilMapErr":   true,
		"nilNoMapErr": true,
		"emptyVal":    val3,
		"emptyErr":    true,
		"validVal":    "Active",
		"validErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_UnmarshallToValue returns correct value -- with args", actual)
}

func Test_BasicString_EnumType(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"enumType": bs.EnumType().String(),
	}

	// Assert
	expected := args.Map{
		"enumType": "String",
	}
	expected.ShouldBeEqual(t, 0, "BasicString_EnumType returns correct value -- with args", actual)
}

func Test_BasicString_AppendPrependJoinValue(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	result := bs.AppendPrependJoinValue(".", "Active", "Invalid")

	// Act
	actual := args.Map{
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_AppendPrependJoinValue returns correct value -- with args", actual)
}

func Test_BasicString_OnlySupportedErr(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	noErr := bs.OnlySupportedErr("Invalid", "Active")
	hasErr := bs.OnlySupportedErr("Invalid")
	msgErr := bs.OnlySupportedMsgErr("context", "Invalid")

	// Act
	actual := args.Map{
		"noErr":  noErr == nil,
		"hasErr": hasErr != nil,
		"msgErr": msgErr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr":  true,
		"hasErr": true,
		"msgErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_OnlySupportedErr returns error -- with args", actual)
}

// ── differCheckerImpl ──

func Test_DifferChecker_GetSingleDiffResult(t *testing.T) {
	// Arrange
	leftResult := enumimpl.DefaultDiffCheckerImpl.GetSingleDiffResult(true, "L", "R")
	rightResult := enumimpl.DefaultDiffCheckerImpl.GetSingleDiffResult(false, "L", "R")

	// Act
	actual := args.Map{
		"left":  leftResult,
		"right": rightResult,
	}

	// Assert
	expected := args.Map{
		"left":  "L",
		"right": "R",
	}
	expected.ShouldBeEqual(t, 0, "DifferChecker_GetSingleDiffResult returns correct value -- with args", actual)
}

func Test_DifferChecker_GetResultOnKeyMissing(t *testing.T) {
	// Arrange
	result := enumimpl.DefaultDiffCheckerImpl.GetResultOnKeyMissingInRightExistInLeft("key", "val")

	// Act
	actual := args.Map{
		"result": result,
	}

	// Assert
	expected := args.Map{
		"result": "val",
	}
	expected.ShouldBeEqual(t, 0, "DifferChecker_GetResultOnKeyMissing returns correct value -- with args", actual)
}

func Test_DifferChecker_IsEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"regardlessSame": enumimpl.DefaultDiffCheckerImpl.IsEqual(true, 1, 1),
		"regardlessDiff": enumimpl.DefaultDiffCheckerImpl.IsEqual(true, 1, 2),
		"strictSame":     enumimpl.DefaultDiffCheckerImpl.IsEqual(false, "a", "a"),
		"strictDiff":     enumimpl.DefaultDiffCheckerImpl.IsEqual(false, "a", "b"),
	}

	// Assert
	expected := args.Map{
		"regardlessSame": true,
		"regardlessDiff": false,
		"strictSame":     true,
		"strictDiff":     false,
	}
	expected.ShouldBeEqual(t, 0, "DifferChecker_IsEqual returns correct value -- with args", actual)
}

func Test_DifferChecker_AsDifferChecker(t *testing.T) {
	// Arrange
	checker := enumimpl.DefaultDiffCheckerImpl.AsDifferChecker()

	// Act
	actual := args.Map{
		"notNil": checker != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "DifferChecker_AsDifferChecker returns correct value -- with args", actual)
}

// ── leftRightDiffCheckerImpl ──

func Test_LeftRightDiffChecker_GetSingleDiffResult(t *testing.T) {
	// Arrange
	result := enumimpl.LeftRightDiffCheckerImpl.GetSingleDiffResult(true, "L", "R")

	// Act
	actual := args.Map{
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRightDiffChecker_GetSingleDiffResult returns correct value -- with args", actual)
}

func Test_LeftRightDiffChecker_GetResultOnKeyMissing(t *testing.T) {
	// Arrange
	result := enumimpl.LeftRightDiffCheckerImpl.GetResultOnKeyMissingInRightExistInLeft("key", "val")

	// Act
	actual := args.Map{
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRightDiffChecker_GetResultOnKeyMissing returns correct value -- with args", actual)
}

func Test_LeftRightDiffChecker_IsEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"same": enumimpl.LeftRightDiffCheckerImpl.IsEqual(false, "a", "a"),
		"diff": enumimpl.LeftRightDiffCheckerImpl.IsEqual(false, "a", "b"),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"diff": false,
	}
	expected.ShouldBeEqual(t, 0, "LeftRightDiffChecker_IsEqual returns correct value -- with args", actual)
}

func Test_LeftRightDiffChecker_AsChecker(t *testing.T) {
	// Arrange
	checker := enumimpl.LeftRightDiffCheckerImpl.AsChecker()

	// Act
	actual := args.Map{
		"notNil": checker != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRightDiffChecker_AsChecker returns correct value -- with args", actual)
}

// ── FormatUsingFmt ──

func Test_FormatUsingFmt(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	result := bb.Format("{type-name}.{name}={value}", byte(1))

	// Act
	actual := args.Map{
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "FormatUsingFmt returns correct value -- with args", actual)
}

// ── ConvEnumAnyValToInteger additional branches ──

func Test_ConvEnumAnyValToInteger_Byte(t *testing.T) {
	// Arrange
	type myByte byte
	result := enumimpl.ConvEnumAnyValToInteger(myByte(5))

	// Act
	actual := args.Map{
		"isPositive": result >= 0,
	}

	// Assert
	expected := args.Map{
		"isPositive": true,
	}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger_Byte returns correct value -- with args", actual)
}

// ── BasicByte via CreateUsingMap ──

func Test_BasicByte_CreateUsingMap(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.CreateUsingMap(
		"testEnum",
		map[byte]string{0: "Invalid", 1: "Active"},
	)

	// Act
	actual := args.Map{
		"typeName": bb.TypeName(),
		"length":   bb.Length(),
	}

	// Assert
	expected := args.Map{
		"typeName": "testEnum",
		"length":   2,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_CreateUsingMap returns correct value -- with args", actual)
}

// ── BasicByte ExpectingEnumValueError ──

func Test_BasicByte_ExpectingEnumValueError(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	noErr := bb.ExpectingEnumValueError("Active", byte(1))
	hasErr := bb.ExpectingEnumValueError("Invalid", byte(1))
	parseErr := bb.ExpectingEnumValueError("NotExist", byte(1))

	// Act
	actual := args.Map{
		"matchNoErr":  noErr == nil,
		"mismatchErr": hasErr != nil,
		"parseErr":    parseErr != nil,
	}

	// Assert
	expected := args.Map{
		"matchNoErr":  true,
		"mismatchErr": true,
		"parseErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_ExpectingEnumValueError returns error -- with args", actual)
}

// ── NamesHashset empty ──

func Test_NumberEnumBase_NamesHashsetEmpty(t *testing.T) {
	// Arrange
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{},
	)

	// Act
	actual := args.Map{
		"len": len(bb.NamesHashset()),
	}

	// Assert
	expected := args.Map{
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_NamesHashsetEmpty returns empty -- with args", actual)
}
