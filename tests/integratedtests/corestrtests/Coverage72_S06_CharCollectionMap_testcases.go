package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ============================================================
// AllIndividualStringsOfStringsLength — all branches
// ============================================================

// Branch: nil pointer input → returns 0
var covS06AllIndStrOfStrLenNilTestCase = coretestcases.CaseV1{
	Title: "AllIndividualStringsOfStringsLength returns 0 -- nil pointer input",
	ArrangeInput: args.Map{
		"useNil": true,
	},
	ExpectedInput: args.Map{
		"length": 0,
	},
}

// Branch: non-nil with items → returns sum of inner lengths
var covS06AllIndStrOfStrLenValidTestCase = coretestcases.CaseV1{
	Title: "AllIndividualStringsOfStringsLength returns total count -- valid slices",
	ArrangeInput: args.Map{
		"useNil": false,
		"items": [][]string{
			{"a", "b"},
			{"c"},
			{"d", "e", "f"},
		},
	},
	ExpectedInput: args.Map{
		"length": 6,
	},
}

// Branch: non-nil but empty outer slice → returns 0
var covS06AllIndStrOfStrLenEmptyTestCase = coretestcases.CaseV1{
	Title: "AllIndividualStringsOfStringsLength returns 0 -- empty outer slice",
	ArrangeInput: args.Map{
		"useNil": false,
		"items":  [][]string{},
	},
	ExpectedInput: args.Map{
		"length": 0,
	},
}

// ============================================================
// AllIndividualsLengthOfSimpleSlices — all branches
// ============================================================

// Branch: nil variadic → returns 0
var covS06AllIndLenSimpleSlicesNilTestCase = coretestcases.CaseV1{
	Title: "AllIndividualsLengthOfSimpleSlices returns 0 -- nil input",
	ArrangeInput: args.Map{
		"useNil": true,
	},
	ExpectedInput: args.Map{
		"length": 0,
	},
}

// Branch: valid slices → returns sum
var covS06AllIndLenSimpleSlicesValidTestCase = coretestcases.CaseV1{
	Title: "AllIndividualsLengthOfSimpleSlices returns total -- valid slices",
	ArrangeInput: args.Map{
		"useNil": false,
		"slice1": []string{"a", "b"},
		"slice2": []string{"c", "d", "e"},
	},
	ExpectedInput: args.Map{
		"length": 5,
	},
}

// ============================================================
// AnyToString — all branches
// ============================================================

var covS06AnyToStringTestCases = []coretestcases.CaseV1{
	{
		Title: "AnyToString returns empty -- empty string input",
		ArrangeInput: args.Map{
			"input":              "",
			"isIncludeFieldName": false,
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "AnyToString returns formatted value -- string without field name",
		ArrangeInput: args.Map{
			"input":              "hello",
			"isIncludeFieldName": false,
		},
		ExpectedInput: args.Map{
			"hasResult": true,
		},
	},
	{
		Title: "AnyToString returns formatted with field -- string with field name",
		ArrangeInput: args.Map{
			"input":              "hello",
			"isIncludeFieldName": true,
		},
		ExpectedInput: args.Map{
			"hasResult": true,
		},
	},
	{
		Title: "AnyToString returns formatted -- pointer input dereferenced",
		ArrangeInput: args.Map{
			"usePointer":         true,
			"isIncludeFieldName": false,
		},
		ExpectedInput: args.Map{
			"hasResult": true,
		},
	},
}

// ============================================================
// CharCollectionDataModel — both constructor directions
// ============================================================

var covS06CharCollDataModelTestCases = []coretestcases.CaseV1{
	{
		Title: "NewCharCollectionMapUsingDataModel returns valid map -- from data model",
		ArrangeInput: args.Map{
			"capacity": 5,
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
	{
		Title: "NewCharCollectionMapDataModelUsing returns data model -- from char map",
		ArrangeInput: args.Map{
			"items": []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
}

// ============================================================
// CloneSlice — branches
// ============================================================

var covS06CloneSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "CloneSlice returns empty -- nil input",
		ArrangeInput: args.Map{
			"useNil": true,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "CloneSlice returns copy -- valid input",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
}

// ============================================================
// CloneSliceIf — branches
// ============================================================

var covS06CloneSliceIfTestCases = []coretestcases.CaseV1{
	{
		Title: "CloneSliceIf returns empty -- empty input",
		ArrangeInput: args.Map{
			"isClone": true,
			"items":   []string{},
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "CloneSliceIf returns same ref -- isClone false",
		ArrangeInput: args.Map{
			"isClone": false,
			"items":   []string{"x", "y"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
	{
		Title: "CloneSliceIf returns cloned -- isClone true",
		ArrangeInput: args.Map{
			"isClone": true,
			"items":   []string{"x", "y"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

// ============================================================
// reflectInterfaceVal — branches
// ============================================================

var covS06ReflectInterfaceValTestCases = []coretestcases.CaseV1{
	{
		Title: "reflectInterfaceVal returns nil -- nil input",
		ArrangeInput: args.Map{
			"useNil": true,
		},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
	{
		Title: "reflectInterfaceVal returns value -- non-pointer input",
		ArrangeInput: args.Map{
			"input": "hello",
		},
		ExpectedInput: args.Map{
			"isNil": false,
		},
	},
}

// ============================================================
// utils — WrapDoubleIfMissing, WrapSingleIfMissing, etc.
// ============================================================

var covS06UtilsWrapTestCases = []coretestcases.CaseV1{
	{
		Title: "WrapDoubleIfMissing returns wrapped -- bare string",
		ArrangeInput: args.Map{
			"method": "WrapDoubleIfMissing",
			"input":  "hello",
		},
		ExpectedInput: args.Map{
			"result": "\"hello\"",
		},
	},
	{
		Title: "WrapDoubleIfMissing returns same -- already wrapped",
		ArrangeInput: args.Map{
			"method": "WrapDoubleIfMissing",
			"input":  "\"hello\"",
		},
		ExpectedInput: args.Map{
			"result": "\"hello\"",
		},
	},
	{
		Title: "WrapDoubleIfMissing returns empty quotes -- empty string",
		ArrangeInput: args.Map{
			"method": "WrapDoubleIfMissing",
			"input":  "",
		},
		ExpectedInput: args.Map{
			"result": "\"\"",
		},
	},
	{
		Title: "WrapSingleIfMissing returns wrapped -- bare string",
		ArrangeInput: args.Map{
			"method": "WrapSingleIfMissing",
			"input":  "hello",
		},
		ExpectedInput: args.Map{
			"result": "'hello'",
		},
	},
	{
		Title: "WrapSingleIfMissing returns same -- already wrapped",
		ArrangeInput: args.Map{
			"method": "WrapSingleIfMissing",
			"input":  "'hello'",
		},
		ExpectedInput: args.Map{
			"result": "'hello'",
		},
	},
	{
		Title: "WrapSingleIfMissing returns empty quotes -- empty string",
		ArrangeInput: args.Map{
			"method": "WrapSingleIfMissing",
			"input":  "",
		},
		ExpectedInput: args.Map{
			"result": "''",
		},
	},
	{
		Title: "WrapDouble returns wrapped -- any string",
		ArrangeInput: args.Map{
			"method": "WrapDouble",
			"input":  "x",
		},
		ExpectedInput: args.Map{
			"result": "\"x\"",
		},
	},
	{
		Title: "WrapSingle returns wrapped -- any string",
		ArrangeInput: args.Map{
			"method": "WrapSingle",
			"input":  "x",
		},
		ExpectedInput: args.Map{
			"result": "'x'",
		},
	},
	{
		Title: "WrapTilda returns wrapped -- any string",
		ArrangeInput: args.Map{
			"method": "WrapTilda",
			"input":  "x",
		},
		ExpectedInput: args.Map{
			"result": "`x`",
		},
	},
}

// ============================================================
// CharCollectionMap — Core Methods (first ~200 stmts)
// ============================================================

// Branch: GetChar with empty string vs non-empty
var covS06CharCollMapGetCharTestCases = []coretestcases.CaseV1{
	{
		Title: "GetChar returns first byte -- non-empty string",
		ArrangeInput: args.Map{
			"input": "hello",
		},
		ExpectedInput: args.Map{
			"char": "h",
		},
	},
	{
		Title: "GetChar returns emptyChar -- empty string",
		ArrangeInput: args.Map{
			"input": "",
		},
		ExpectedInput: args.Map{
			"char": "\x00",
		},
	},
}

// Branch: GetCharsGroups empty vs non-empty
var covS06CharCollMapGetCharsGroupsTestCases = []coretestcases.CaseV1{
	{
		Title: "GetCharsGroups returns self -- empty items",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: args.Map{
			"isNil": false,
		},
	},
	{
		Title: "GetCharsGroups returns grouped map -- multiple items",
		ArrangeInput: args.Map{
			"items": []string{"apple", "avocado", "banana"},
		},
		ExpectedInput: args.Map{
			"hasItems": true,
		},
	},
}

// Branch: IsEmpty, HasItems, Length, AllLengthsSum
var covS06CharCollMapBasicTestCases = []coretestcases.CaseV1{
	{
		Title: "CharCollectionMap IsEmpty returns true -- empty map",
		ArrangeInput: args.Map{
			"useEmpty": true,
		},
		ExpectedInput: args.Map{
			"isEmpty":       true,
			"hasItems":      false,
			"length":        0,
			"allLengthsSum": 0,
		},
	},
	{
		Title: "CharCollectionMap IsEmpty returns false -- populated map",
		ArrangeInput: args.Map{
			"items": []string{"apple", "banana"},
		},
		ExpectedInput: args.Map{
			"isEmpty":       false,
			"hasItems":      true,
			"length":        2,
			"allLengthsSum": 2,
		},
	},
}

// Branch: Add, AddStrings, Has
var covS06CharCollMapAddHasTestCases = []coretestcases.CaseV1{
	{
		Title: "CharCollectionMap Add returns map with item -- single add",
		ArrangeInput: args.Map{
			"addItem": "hello",
		},
		ExpectedInput: args.Map{
			"has":    true,
			"length": 1,
		},
	},
	{
		Title: "CharCollectionMap Add adds to existing group -- same first char",
		ArrangeInput: args.Map{
			"addItems": []string{"hello", "hi", "hey"},
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 3,
			"charGroups":    1,
		},
	},
	{
		Title: "CharCollectionMap Has returns false -- item not present",
		ArrangeInput: args.Map{
			"addItem":   "hello",
			"checkItem": "world",
		},
		ExpectedInput: args.Map{
			"has": false,
		},
	},
	{
		Title: "CharCollectionMap Has returns false -- empty map",
		ArrangeInput: args.Map{
			"useEmpty":  true,
			"checkItem": "hello",
		},
		ExpectedInput: args.Map{
			"has": false,
		},
	},
}

// Branch: LengthOf, LengthOfCollectionFromFirstChar
var covS06CharCollMapLengthOfTestCases = []coretestcases.CaseV1{
	{
		Title: "LengthOf returns count -- char with items",
		ArrangeInput: args.Map{
			"items": []string{"apple", "avocado"},
			"char":  "a",
		},
		ExpectedInput: args.Map{
			"lengthOf": 2,
		},
	},
	{
		Title: "LengthOf returns 0 -- char without items",
		ArrangeInput: args.Map{
			"items": []string{"apple"},
			"char":  "z",
		},
		ExpectedInput: args.Map{
			"lengthOf": 0,
		},
	},
	{
		Title: "LengthOfCollectionFromFirstChar returns count -- matching string",
		ArrangeInput: args.Map{
			"items":    []string{"apple", "avocado"},
			"checkStr": "abc",
		},
		ExpectedInput: args.Map{
			"lengthFromChar": 2,
		},
	},
}

// Branch: IsEquals, IsEqualsCaseSensitive
var covS06CharCollMapEqualsTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEquals returns true -- same items",
		ArrangeInput: args.Map{
			"items1": []string{"apple", "banana"},
			"items2": []string{"apple", "banana"},
		},
		ExpectedInput: args.Map{
			"isEquals": true,
		},
	},
	{
		Title: "IsEquals returns false -- different items",
		ArrangeInput: args.Map{
			"items1": []string{"apple"},
			"items2": []string{"banana"},
		},
		ExpectedInput: args.Map{
			"isEquals": false,
		},
	},
	{
		Title: "IsEquals returns false -- nil other",
		ArrangeInput: args.Map{
			"items1":      []string{"apple"},
			"useNilOther": true,
		},
		ExpectedInput: args.Map{
			"isEquals": false,
		},
	},
	{
		Title: "IsEquals returns true -- both empty",
		ArrangeInput: args.Map{
			"useEmpty1": true,
			"useEmpty2": true,
		},
		ExpectedInput: args.Map{
			"isEquals": true,
		},
	},
	{
		Title: "IsEquals returns true -- same pointer",
		ArrangeInput: args.Map{
			"useSelf": true,
			"items1":  []string{"apple"},
		},
		ExpectedInput: args.Map{
			"isEquals": true,
		},
	},
}

// Branch: HasWithCollection
var covS06HasWithCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "HasWithCollection returns true with collection -- item exists",
		ArrangeInput: args.Map{
			"items":    []string{"apple", "avocado"},
			"checkStr": "apple",
		},
		ExpectedInput: args.Map{
			"found":            true,
			"collectionNotNil": true,
		},
	},
	{
		Title: "HasWithCollection returns false -- item missing",
		ArrangeInput: args.Map{
			"items":    []string{"apple"},
			"checkStr": "banana",
		},
		ExpectedInput: args.Map{
			"found":            false,
			"collectionNotNil": true,
		},
	},
	{
		Title: "HasWithCollection returns false -- empty map",
		ArrangeInput: args.Map{
			"useEmpty": true,
			"checkStr": "apple",
		},
		ExpectedInput: args.Map{
			"found":            false,
			"collectionNotNil": true,
		},
	},
}

// Branch: GetCollection, GetCollectionByChar
var covS06GetCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "GetCollection returns collection -- char exists",
		ArrangeInput: args.Map{
			"items":           []string{"apple"},
			"checkStr":        "a-test",
			"isAddNewOnEmpty": false,
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
	{
		Title: "GetCollection returns nil -- char missing no-add",
		ArrangeInput: args.Map{
			"items":           []string{"apple"},
			"checkStr":        "z-test",
			"isAddNewOnEmpty": false,
		},
		ExpectedInput: args.Map{
			"isNotNil": false,
		},
	},
	{
		Title: "GetCollection creates new -- char missing with add",
		ArrangeInput: args.Map{
			"items":           []string{"apple"},
			"checkStr":        "z-test",
			"isAddNewOnEmpty": true,
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
}

// Branch: SortedListAsc, List
var covS06ListTestCases = []coretestcases.CaseV1{
	{
		Title: "List returns all items -- populated map",
		ArrangeInput: args.Map{
			"items": []string{"banana", "apple"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
	{
		Title: "SortedListAsc returns sorted -- populated map",
		ArrangeInput: args.Map{
			"items":  []string{"banana", "apple", "avocado"},
			"sorted": true,
		},
		ExpectedInput: args.Map{
			"first": "apple",
		},
	},
	{
		Title: "SortedListAsc returns empty -- empty map",
		ArrangeInput: args.Map{
			"useEmpty": true,
			"sorted":   true,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}

// Branch: JSON serialization/deserialization
var covS06CharCollMapJsonTestCases = []coretestcases.CaseV1{
	{
		Title: "CharCollectionMap Json returns valid result -- populated map",
		ArrangeInput: args.Map{
			"items": []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"hasBytes": true,
			"hasError": false,
		},
	},
	{
		Title: "CharCollectionMap ParseInjectUsingJson round-trips -- valid json",
		ArrangeInput: args.Map{
			"items": []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"roundTrip": true,
		},
	},
}

// Branch: AddSameStartingCharItems
var covS06AddSameCharItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "AddSameStartingCharItems adds to existing -- char exists",
		ArrangeInput: args.Map{
			"initialItems": []string{"apple"},
			"newItems":     []string{"avocado", "apricot"},
			"char":         "a",
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 3,
		},
	},
	{
		Title: "AddSameStartingCharItems creates new -- char missing",
		ArrangeInput: args.Map{
			"initialItems": []string{"apple"},
			"newItems":     []string{"banana"},
			"char":         "b",
			"isCloneAdd":   true,
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 2,
		},
	},
	{
		Title: "AddSameStartingCharItems returns self -- empty items",
		ArrangeInput: args.Map{
			"initialItems": []string{"apple"},
			"newItems":     []string{},
			"char":         "z",
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 1,
		},
	},
}

// Branch: Resize, AddLength
var covS06ResizeTestCases = []coretestcases.CaseV1{
	{
		Title: "Resize returns self -- newLength <= current",
		ArrangeInput: args.Map{
			"items":     []string{"apple", "banana"},
			"newLength": 1,
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
	{
		Title: "Resize expands -- newLength > current",
		ArrangeInput: args.Map{
			"items":     []string{"apple"},
			"newLength": 50,
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
	{
		Title: "AddLength expands capacity -- additional lengths",
		ArrangeInput: args.Map{
			"items":   []string{"apple"},
			"lengths": []int{10, 20},
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
}

// Branch: Clear, Dispose
var covS06ClearDisposeTestCases = []coretestcases.CaseV1{
	{
		Title: "Clear removes all items -- populated map",
		ArrangeInput: args.Map{
			"items": []string{"apple", "banana"},
		},
		ExpectedInput: args.Map{
			"isEmpty": true,
		},
	},
	{
		Title: "Dispose sets items nil -- populated map",
		ArrangeInput: args.Map{
			"items": []string{"apple"},
		},
		ExpectedInput: args.Map{
			"isDisposed": true,
		},
	},
}

// Branch: HashsetByChar, HashsetByStringFirstChar
var covS06HashsetByCharTestCases = []coretestcases.CaseV1{
	{
		Title: "HashsetByChar returns hashset -- char exists",
		ArrangeInput: args.Map{
			"items": []string{"apple", "avocado"},
			"char":  "a",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
	{
		Title: "HashsetByChar returns nil -- char missing",
		ArrangeInput: args.Map{
			"items": []string{"apple"},
			"char":  "z",
		},
		ExpectedInput: args.Map{
			"isNotNil": false,
		},
	},
}

// Branch: HashsetsCollection, HashsetsCollectionByChars
var covS06HashsetsCollTestCases = []coretestcases.CaseV1{
	{
		Title: "HashsetsCollection returns collection -- populated map",
		ArrangeInput: args.Map{
			"items": []string{"apple", "banana"},
		},
		ExpectedInput: args.Map{
			"hasItems": true,
		},
	},
	{
		Title: "HashsetsCollection returns empty -- empty map",
		ArrangeInput: args.Map{
			"useEmpty": true,
		},
		ExpectedInput: args.Map{
			"hasItems": false,
		},
	},
}

// Branch: AddHashmapsValues, AddHashmapsKeysValuesBoth
var covS06AddHashmapsTestCases = []coretestcases.CaseV1{
	{
		Title: "AddHashmapsValues adds values -- from hashmap",
		ArrangeInput: args.Map{
			"keys":   []string{"k1", "k2"},
			"values": []string{"v1", "v2"},
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 2,
		},
	},
	{
		Title: "AddHashmapsKeysValuesBoth adds keys and values -- from hashmap",
		ArrangeInput: args.Map{
			"addBoth": true,
			"keys":    []string{"k1"},
			"values":  []string{"v1"},
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 2,
		},
	},
}

// Branch: AddCollectionItems
var covS06AddCollectionItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "AddCollectionItems adds from collection -- non-empty",
		ArrangeInput: args.Map{
			"items": []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 2,
		},
	},
	{
		Title: "AddCollectionItems returns self -- nil collection",
		ArrangeInput: args.Map{
			"useNil": true,
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 0,
		},
	},
}

// Branch: Print (isPrint=false skips), String, SummaryString
var covS06StringOutputTestCases = []coretestcases.CaseV1{
	{
		Title: "String returns formatted output -- populated map",
		ArrangeInput: args.Map{
			"items": []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"hasOutput": true,
		},
	},
	{
		Title: "SummaryString returns formatted summary -- populated map",
		ArrangeInput: args.Map{
			"items":   []string{"alpha", "bravo"},
			"summary": true,
		},
		ExpectedInput: args.Map{
			"hasOutput": true,
		},
	},
}

// Branch: JsonModel, AsJsonContractsBinder, AsJsoner, etc.
var covS06InterfaceAdaptersTestCases = []coretestcases.CaseV1{
	{
		Title: "AsJsonContractsBinder returns non-nil -- valid map",
		ArrangeInput: args.Map{
			"items": []string{"alpha"},
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
}

// Branch: AddCharHashsetMap
var covS06AddCharHashsetMapTestCases = []coretestcases.CaseV1{
	{
		Title: "AddCharHashsetMap adds items -- from non-empty hashset map",
		ArrangeInput: args.Map{
			"hashsetItems": []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 2,
		},
	},
	{
		Title: "AddCharHashsetMap returns self -- empty hashset map",
		ArrangeInput: args.Map{
			"useEmptyHashset": true,
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 0,
		},
	},
}

// Branch: AddSameCharsCollection
var covS06AddSameCharsCollTestCases = []coretestcases.CaseV1{
	{
		Title: "AddSameCharsCollection adds to existing -- char exists with collection",
		ArrangeInput: args.Map{
			"initialItems": []string{"apple"},
			"addItems":     []string{"avocado", "apricot"},
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 3,
		},
	},
	{
		Title: "AddSameCharsCollection returns existing -- char exists empty collection",
		ArrangeInput: args.Map{
			"initialItems":    []string{"apple"},
			"useNilCollToAdd": true,
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 1,
		},
	},
	{
		Title: "AddSameCharsCollection creates new empty -- char missing nil collection",
		ArrangeInput: args.Map{
			"initialItems":    []string{"banana"},
			"useNilCollToAdd": true,
			"checkStr":        "apple",
		},
		ExpectedInput: args.Map{
			"collectionNotNil": true,
		},
	},
	{
		Title: "AddSameCharsCollection assigns collection -- char missing with data",
		ArrangeInput: args.Map{
			"initialItems": []string{"banana"},
			"addItems":     []string{"apple", "avocado"},
			"checkStr":     "apple",
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 3,
		},
	},
}

// Branch: newCharCollectionMapCreator methods
var covS06NewCreatorTestCases = []coretestcases.CaseV1{
	{
		Title: "CapSelfCap returns map with capacity -- large caps",
		ArrangeInput: args.Map{
			"cap":     20,
			"selfCap": 15,
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"isEmpty":  true,
		},
	},
	{
		Title: "CapSelfCap applies minimum -- small caps",
		ArrangeInput: args.Map{
			"cap":     1,
			"selfCap": 1,
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
	{
		Title: "Empty returns empty map -- default",
		ArrangeInput: args.Map{
			"useEmpty": true,
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"isEmpty":  true,
		},
	},
	{
		Title: "Items returns populated map -- with items",
		ArrangeInput: args.Map{
			"useItems": true,
			"items":    []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"isEmpty":  false,
		},
	},
	{
		Title: "Items returns empty map -- empty items",
		ArrangeInput: args.Map{
			"useItems": true,
			"items":    []string{},
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"isEmpty":  true,
		},
	},
	{
		Title: "ItemsPtrWithCap returns map -- with items and cap",
		ArrangeInput: args.Map{
			"useItemsPtr":   true,
			"items":         []string{"alpha"},
			"additionalCap": 5,
			"eachCap":       3,
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"isEmpty":  false,
		},
	},
	{
		Title: "ItemsPtrWithCap returns empty -- no items",
		ArrangeInput: args.Map{
			"useItemsPtr":   true,
			"items":         []string{},
			"additionalCap": 5,
			"eachCap":       3,
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"isEmpty":  true,
		},
	},
}
