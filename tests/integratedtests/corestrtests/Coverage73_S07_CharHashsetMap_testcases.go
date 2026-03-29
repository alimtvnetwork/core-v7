package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ============================================================
// CharHashsetDataModel — constructor directions
// ============================================================

var covS07CharHashsetDataModelTestCases = []coretestcases.CaseV1{
	{
		Title: "NewCharHashsetMapUsingDataModel returns valid map -- from data model",
		ArrangeInput: args.Map{
			"direction": "toMap",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
	{
		Title: "NewCharHashsetMapDataModelUsing returns data model -- from hashset map",
		ArrangeInput: args.Map{
			"direction": "toModel",
			"items":     []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
}

// ============================================================
// CharHashsetMap — GetChar, GetCharOf
// ============================================================

var covS07GetCharTestCases = []coretestcases.CaseV1{
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
	{
		Title: "GetCharOf returns first byte -- non-empty string",
		ArrangeInput: args.Map{
			"input":    "world",
			"useGetOf": true,
		},
		ExpectedInput: args.Map{
			"char": "w",
		},
	},
	{
		Title: "GetCharOf returns emptyChar -- empty string",
		ArrangeInput: args.Map{
			"input":    "",
			"useGetOf": true,
		},
		ExpectedInput: args.Map{
			"char": "\x00",
		},
	},
}

// ============================================================
// CharHashsetMap — IsEmpty, HasItems, Length, AllLengthsSum
// ============================================================

var covS07BasicTestCases = []coretestcases.CaseV1{
	{
		Title: "CharHashsetMap IsEmpty returns true -- empty map",
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
		Title: "CharHashsetMap HasItems returns true -- populated map",
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

// ============================================================
// CharHashsetMap — Has, HasWithHashset
// ============================================================

var covS07HasTestCases = []coretestcases.CaseV1{
	{
		Title: "Has returns true -- item exists",
		ArrangeInput: args.Map{
			"items":    []string{"apple", "banana"},
			"checkStr": "apple",
		},
		ExpectedInput: args.Map{
			"has": true,
		},
	},
	{
		Title: "Has returns false -- item missing",
		ArrangeInput: args.Map{
			"items":    []string{"apple"},
			"checkStr": "banana",
		},
		ExpectedInput: args.Map{
			"has": false,
		},
	},
	{
		Title: "Has returns false -- empty map",
		ArrangeInput: args.Map{
			"useEmpty": true,
			"checkStr": "apple",
		},
		ExpectedInput: args.Map{
			"has": false,
		},
	},
	{
		Title: "Has returns false -- char group exists but item not in it",
		ArrangeInput: args.Map{
			"items":    []string{"alpha"},
			"checkStr": "avocado",
		},
		ExpectedInput: args.Map{
			"has": false,
		},
	},
}

// Branch: HasWithHashset
var covS07HasWithHashsetTestCases = []coretestcases.CaseV1{
	{
		Title: "HasWithHashset returns true with hashset -- item exists",
		ArrangeInput: args.Map{
			"items":    []string{"apple", "avocado"},
			"checkStr": "apple",
		},
		ExpectedInput: args.Map{
			"found":    true,
			"hsNotNil": true,
		},
	},
	{
		Title: "HasWithHashset returns false -- char missing",
		ArrangeInput: args.Map{
			"items":    []string{"apple"},
			"checkStr": "banana",
		},
		ExpectedInput: args.Map{
			"found":    false,
			"hsNotNil": true,
		},
	},
	{
		Title: "HasWithHashset returns false -- empty map",
		ArrangeInput: args.Map{
			"useEmpty": true,
			"checkStr": "apple",
		},
		ExpectedInput: args.Map{
			"found":    false,
			"hsNotNil": true,
		},
	},
}

// ============================================================
// CharHashsetMap — LengthOf, LengthOfHashsetFromFirstChar
// ============================================================

var covS07LengthOfTestCases = []coretestcases.CaseV1{
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
		Title: "LengthOf returns 0 -- char missing",
		ArrangeInput: args.Map{
			"items": []string{"apple"},
			"char":  "z",
		},
		ExpectedInput: args.Map{
			"lengthOf": 0,
		},
	},
	{
		Title: "LengthOf returns 0 -- empty map",
		ArrangeInput: args.Map{
			"useEmpty": true,
			"char":     "a",
		},
		ExpectedInput: args.Map{
			"lengthOf": 0,
		},
	},
	{
		Title: "LengthOfHashsetFromFirstChar returns count -- matching string",
		ArrangeInput: args.Map{
			"items":    []string{"apple", "avocado"},
			"checkStr": "abc",
		},
		ExpectedInput: args.Map{
			"lengthFromChar": 2,
		},
	},
	{
		Title: "LengthOfHashsetFromFirstChar returns 0 -- no match",
		ArrangeInput: args.Map{
			"items":    []string{"apple"},
			"checkStr": "zoo",
		},
		ExpectedInput: args.Map{
			"lengthFromChar": 0,
		},
	},
}

// ============================================================
// CharHashsetMap — Add, AddStrings
// ============================================================

var covS07AddTestCases = []coretestcases.CaseV1{
	{
		Title: "Add inserts item -- new char group",
		ArrangeInput: args.Map{
			"addItem": "hello",
		},
		ExpectedInput: args.Map{
			"has":    true,
			"length": 1,
		},
	},
	{
		Title: "Add adds to existing group -- same first char",
		ArrangeInput: args.Map{
			"addItems": []string{"hello", "hi", "hey"},
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 3,
			"charGroups":    1,
		},
	},
	{
		Title: "AddStrings adds multiple items -- mixed chars",
		ArrangeInput: args.Map{
			"addItems": []string{"apple", "banana", "avocado"},
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 3,
			"charGroups":    2,
		},
	},
}

// ============================================================
// CharHashsetMap — AddSameStartingCharItems
// ============================================================

var covS07AddSameCharTestCases = []coretestcases.CaseV1{
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

// ============================================================
// CharHashsetMap — IsEquals
// ============================================================

var covS07IsEqualsTestCases = []coretestcases.CaseV1{
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
	{
		Title: "IsEquals returns false -- one empty one not",
		ArrangeInput: args.Map{
			"items1":    []string{"apple"},
			"useEmpty2": true,
		},
		ExpectedInput: args.Map{
			"isEquals": false,
		},
	},
	{
		Title: "IsEquals returns false -- different lengths",
		ArrangeInput: args.Map{
			"items1": []string{"apple", "banana"},
			"items2": []string{"apple"},
		},
		ExpectedInput: args.Map{
			"isEquals": false,
		},
	},
	{
		Title: "IsEquals returns false -- same length different keys",
		ArrangeInput: args.Map{
			"items1": []string{"alpha"},
			"items2": []string{"bravo"},
		},
		ExpectedInput: args.Map{
			"isEquals": false,
		},
	},
}

// ============================================================
// CharHashsetMap — GetHashset
// ============================================================

var covS07GetHashsetTestCases = []coretestcases.CaseV1{
	{
		Title: "GetHashset returns hashset -- char exists",
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
		Title: "GetHashset returns nil -- char missing no-add",
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
		Title: "GetHashset creates new -- char missing with add",
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

// ============================================================
// CharHashsetMap — GetCharsGroups
// ============================================================

var covS07GetCharsGroupsTestCases = []coretestcases.CaseV1{
	{
		Title: "GetCharsGroups returns self -- empty items",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
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

// ============================================================
// CharHashsetMap — List, SortedListAsc, SortedListDsc
// ============================================================

var covS07ListTestCases = []coretestcases.CaseV1{
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
		Title: "SortedListAsc returns sorted ascending -- populated map",
		ArrangeInput: args.Map{
			"items":  []string{"banana", "apple", "avocado"},
			"sorted": "asc",
		},
		ExpectedInput: args.Map{
			"first": "apple",
		},
	},
	{
		Title: "SortedListDsc returns sorted descending -- populated map",
		ArrangeInput: args.Map{
			"items":  []string{"banana", "apple", "avocado"},
			"sorted": "dsc",
		},
		ExpectedInput: args.Map{
			"first": "banana",
		},
	},
}

// ============================================================
// CharHashsetMap — JSON
// ============================================================

var covS07JsonTestCases = []coretestcases.CaseV1{
	{
		Title: "CharHashsetMap Json returns valid bytes -- populated map",
		ArrangeInput: args.Map{
			"items": []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"hasBytes": true,
			"hasError": false,
		},
	},
	{
		Title: "CharHashsetMap ParseInjectUsingJson round-trips -- valid json",
		ArrangeInput: args.Map{
			"items": []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"roundTrip": true,
		},
	},
}

// ============================================================
// CharHashsetMap — AddSameCharsCollection
// ============================================================

var covS07AddSameCharsCollTestCases = []coretestcases.CaseV1{
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
		Title: "AddSameCharsCollection returns existing -- char exists nil collection",
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
			"hsNotNil": true,
		},
	},
	{
		Title: "AddSameCharsCollection assigns from collection -- char missing with data",
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

// ============================================================
// CharHashsetMap — AddSameCharsHashset
// ============================================================

var covS07AddSameCharsHashsetTestCases = []coretestcases.CaseV1{
	{
		Title: "AddSameCharsHashset adds to existing -- char exists",
		ArrangeInput: args.Map{
			"initialItems": []string{"apple"},
			"addItems":     []string{"avocado"},
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 2,
		},
	},
	{
		Title: "AddSameCharsHashset returns existing -- char exists nil hashset",
		ArrangeInput: args.Map{
			"initialItems":  []string{"apple"},
			"useNilHashset": true,
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 1,
		},
	},
	{
		Title: "AddSameCharsHashset creates new -- char missing nil hashset",
		ArrangeInput: args.Map{
			"initialItems":  []string{"banana"},
			"useNilHashset": true,
			"checkStr":      "apple",
		},
		ExpectedInput: args.Map{
			"hsNotNil": true,
		},
	},
	{
		Title: "AddSameCharsHashset assigns hashset -- char missing with data",
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

// ============================================================
// CharHashsetMap — AddCollectionItems, AddCharCollectionMapItems, AddHashsetItems
// ============================================================

var covS07AddFromSourceTestCases = []coretestcases.CaseV1{
	{
		Title: "AddCollectionItems adds from collection -- non-empty",
		ArrangeInput: args.Map{
			"source": "collection",
			"items":  []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 2,
		},
	},
	{
		Title: "AddCollectionItems returns self -- nil collection",
		ArrangeInput: args.Map{
			"source": "collectionNil",
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 0,
		},
	},
	{
		Title: "AddCharCollectionMapItems adds from char coll map -- non-empty",
		ArrangeInput: args.Map{
			"source": "charCollMap",
			"items":  []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 2,
		},
	},
	{
		Title: "AddCharCollectionMapItems returns self -- nil map",
		ArrangeInput: args.Map{
			"source": "charCollMapNil",
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 0,
		},
	},
	{
		Title: "AddHashsetItems adds from hashset -- non-empty",
		ArrangeInput: args.Map{
			"source": "hashset",
			"items":  []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 2,
		},
	},
	{
		Title: "AddHashsetItems returns self -- empty hashset",
		ArrangeInput: args.Map{
			"source": "hashsetEmpty",
		},
		ExpectedInput: args.Map{
			"allLengthsSum": 0,
		},
	},
}

// ============================================================
// CharHashsetMap — HashsetsCollection, ByChars, ByStringsFirstChar
// ============================================================

var covS07HashsetsCollTestCases = []coretestcases.CaseV1{
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

// ============================================================
// CharHashsetMap — Clear, RemoveAll
// ============================================================

var covS07ClearTestCases = []coretestcases.CaseV1{
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
		Title: "RemoveAll clears all items -- populated map",
		ArrangeInput: args.Map{
			"items":     []string{"apple"},
			"useRemove": true,
		},
		ExpectedInput: args.Map{
			"isEmpty": true,
		},
	},
	{
		Title: "Clear returns self -- empty map",
		ArrangeInput: args.Map{
			"useEmpty": true,
		},
		ExpectedInput: args.Map{
			"isEmpty": true,
		},
	},
}

// ============================================================
// CharHashsetMap — String output
// ============================================================

var covS07StringOutputTestCases = []coretestcases.CaseV1{
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
		Title: "SummaryString returns summary -- populated map",
		ArrangeInput: args.Map{
			"items":   []string{"alpha", "bravo"},
			"summary": true,
		},
		ExpectedInput: args.Map{
			"hasOutput": true,
		},
	},
}

// ============================================================
// CharHashsetMap — newCreator methods
// ============================================================

var covS07NewCreatorTestCases = []coretestcases.CaseV1{
	{
		Title: "Cap returns map with capacity -- valid caps",
		ArrangeInput: args.Map{
			"method":  "cap",
			"cap":     20,
			"selfCap": 15,
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"isEmpty":  true,
		},
	},
	{
		Title: "Cap applies minimum -- small caps",
		ArrangeInput: args.Map{
			"method":  "cap",
			"cap":     1,
			"selfCap": 1,
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
	{
		Title: "CapItems returns populated map -- with items",
		ArrangeInput: args.Map{
			"method":  "capItems",
			"cap":     10,
			"selfCap": 5,
			"items":   []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"isEmpty":  false,
		},
	},
	{
		Title: "Strings returns populated map -- with items",
		ArrangeInput: args.Map{
			"method":  "strings",
			"selfCap": 5,
			"items":   []string{"alpha", "bravo"},
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"isEmpty":  false,
		},
	},
	{
		Title: "Strings returns empty map -- nil items",
		ArrangeInput: args.Map{
			"method":  "stringsNil",
			"selfCap": 5,
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"isEmpty":  true,
		},
	},
}
