package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

type testUserCov23 struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Age   int    `json:"Age"`
}

func makeTypedWrapperCov23(name, id string, data testUserCov23) *corepayload.TypedPayloadWrapper[testUserCov23] {
	tw, err := corepayload.NewTypedPayloadWrapperFrom[testUserCov23](name, id, "testUser", data)
	if err != nil {
		panic(err)
	}
	return tw
}

func makeCollectionCov23() *corepayload.TypedPayloadCollection[testUserCov23] {
	col := corepayload.NewTypedPayloadCollection[testUserCov23](3)
	col.Add(makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice", Email: "a@a.com", Age: 30}))
	col.Add(makeTypedWrapperCov23("user", "2", testUserCov23{Name: "Bob", Email: "b@b.com", Age: 25}))
	col.Add(makeTypedWrapperCov23("user", "3", testUserCov23{Name: "Carol", Email: "c@c.com", Age: 35}))
	return col
}

// ── MapTypedPayloads ──

func Test_Cov23_MapTypedPayloads(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	names := corepayload.MapTypedPayloads[testUserCov23, string](col,
		func(item *corepayload.TypedPayloadWrapper[testUserCov23]) string {
			return item.Data().Name
		},
	)

	// Assert
	actual := args.Map{"length": len(names), "first": names[0]}
	expected := args.Map{"length": 3, "first": "Alice"}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloads returns names -- 3 items", actual)
}

func Test_Cov23_MapTypedPayloads_Empty(t *testing.T) {
	// Arrange
	col := corepayload.EmptyTypedPayloadCollection[testUserCov23]()

	// Act
	result := corepayload.MapTypedPayloads[testUserCov23, string](col,
		func(item *corepayload.TypedPayloadWrapper[testUserCov23]) string {
			return item.Data().Name
		},
	)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloads returns empty -- empty source", actual)
}

// ── MapTypedPayloadData ──

func Test_Cov23_MapTypedPayloadData(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	emails := corepayload.MapTypedPayloadData[testUserCov23, string](col,
		func(u testUserCov23) string { return u.Email },
	)

	// Assert
	actual := args.Map{"length": len(emails), "first": emails[0]}
	expected := args.Map{"length": 3, "first": "a@a.com"}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloadData returns emails -- 3 items", actual)
}

// ── FlatMapTypedPayloads ──

func Test_Cov23_FlatMapTypedPayloads(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := corepayload.FlatMapTypedPayloads[testUserCov23, string](col,
		func(item *corepayload.TypedPayloadWrapper[testUserCov23]) []string {
			return []string{item.Data().Name, item.Data().Email}
		},
	)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 6}
	expected.ShouldBeEqual(t, 0, "FlatMapTypedPayloads returns flattened -- 3x2", actual)
}

func Test_Cov23_FlatMapTypedPayloads_Empty(t *testing.T) {
	// Arrange
	col := corepayload.EmptyTypedPayloadCollection[testUserCov23]()

	// Act
	result := corepayload.FlatMapTypedPayloads[testUserCov23, string](col,
		func(item *corepayload.TypedPayloadWrapper[testUserCov23]) []string {
			return []string{item.Data().Name}
		},
	)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "FlatMapTypedPayloads returns empty -- empty source", actual)
}

// ── FlatMapTypedPayloadData ──

func Test_Cov23_FlatMapTypedPayloadData(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := corepayload.FlatMapTypedPayloadData[testUserCov23, string](col,
		func(u testUserCov23) []string { return []string{u.Name} },
	)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "FlatMapTypedPayloadData returns names -- 3 items", actual)
}

// ── ReduceTypedPayloads ──

func Test_Cov23_ReduceTypedPayloads(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	totalAge := corepayload.ReduceTypedPayloads[testUserCov23, int](col, 0,
		func(acc int, item *corepayload.TypedPayloadWrapper[testUserCov23]) int {
			return acc + item.Data().Age
		},
	)

	// Assert
	actual := args.Map{"totalAge": totalAge}
	expected := args.Map{"totalAge": 90}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloads returns sum -- 30+25+35", actual)
}

func Test_Cov23_ReduceTypedPayloads_Empty(t *testing.T) {
	// Arrange
	col := corepayload.EmptyTypedPayloadCollection[testUserCov23]()

	// Act
	result := corepayload.ReduceTypedPayloads[testUserCov23, int](col, 42,
		func(acc int, item *corepayload.TypedPayloadWrapper[testUserCov23]) int {
			return acc + 1
		},
	)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": 42}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloads returns initial -- empty", actual)
}

// ── ReduceTypedPayloadData ──

func Test_Cov23_ReduceTypedPayloadData(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := corepayload.ReduceTypedPayloadData[testUserCov23, int](col, 0,
		func(acc int, u testUserCov23) int { return acc + u.Age },
	)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": 90}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloadData returns sum -- 90", actual)
}

// ── GroupTypedPayloads ──

func Test_Cov23_GroupTypedPayloads(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	groups := corepayload.GroupTypedPayloads[testUserCov23, string](col,
		func(item *corepayload.TypedPayloadWrapper[testUserCov23]) string {
			if item.Data().Age >= 30 {
				return "senior"
			}
			return "junior"
		},
	)

	// Assert
	actual := args.Map{
		"groupCount":   len(groups),
		"seniorCount":  groups["senior"].Length(),
		"juniorCount":  groups["junior"].Length(),
	}
	expected := args.Map{
		"groupCount":   2,
		"seniorCount":  2,
		"juniorCount":  1,
	}
	expected.ShouldBeEqual(t, 0, "GroupTypedPayloads returns grouped -- 2 groups", actual)
}

func Test_Cov23_GroupTypedPayloads_Empty(t *testing.T) {
	// Arrange
	col := corepayload.EmptyTypedPayloadCollection[testUserCov23]()

	// Act
	groups := corepayload.GroupTypedPayloads[testUserCov23, string](col,
		func(item *corepayload.TypedPayloadWrapper[testUserCov23]) string { return "x" },
	)

	// Assert
	actual := args.Map{"count": len(groups)}
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "GroupTypedPayloads returns empty -- empty source", actual)
}

// ── GroupTypedPayloadData ──

func Test_Cov23_GroupTypedPayloadData(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	groups := corepayload.GroupTypedPayloadData[testUserCov23, string](col,
		func(u testUserCov23) string {
			if u.Age >= 30 {
				return "old"
			}
			return "young"
		},
	)

	// Assert
	actual := args.Map{"groupCount": len(groups)}
	expected := args.Map{"groupCount": 2}
	expected.ShouldBeEqual(t, 0, "GroupTypedPayloadData returns grouped -- 2", actual)
}

// ── PartitionTypedPayloads ──

func Test_Cov23_PartitionTypedPayloads(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	matching, notMatching := corepayload.PartitionTypedPayloads[testUserCov23](col,
		func(item *corepayload.TypedPayloadWrapper[testUserCov23]) bool {
			return item.Data().Age >= 30
		},
	)

	// Assert
	actual := args.Map{
		"matchingLen":    matching.Length(),
		"notMatchingLen": notMatching.Length(),
	}
	expected := args.Map{
		"matchingLen":    2,
		"notMatchingLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "PartitionTypedPayloads returns split -- 2+1", actual)
}

func Test_Cov23_PartitionTypedPayloads_Empty(t *testing.T) {
	// Arrange
	col := corepayload.EmptyTypedPayloadCollection[testUserCov23]()

	// Act
	matching, notMatching := corepayload.PartitionTypedPayloads[testUserCov23](col,
		func(item *corepayload.TypedPayloadWrapper[testUserCov23]) bool { return true },
	)

	// Assert
	actual := args.Map{
		"matchingLen":    matching.Length(),
		"notMatchingLen": notMatching.Length(),
	}
	expected := args.Map{
		"matchingLen":    0,
		"notMatchingLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "PartitionTypedPayloads returns empty -- empty source", actual)
}

// ── AnyTypedPayload / AllTypedPayloads ──

func Test_Cov23_AnyTypedPayload_True(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := corepayload.AnyTypedPayload[testUserCov23](col,
		func(item *corepayload.TypedPayloadWrapper[testUserCov23]) bool {
			return item.Data().Name == "Bob"
		},
	)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyTypedPayload returns true -- Bob exists", actual)
}

func Test_Cov23_AnyTypedPayload_False(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := corepayload.AnyTypedPayload[testUserCov23](col,
		func(item *corepayload.TypedPayloadWrapper[testUserCov23]) bool {
			return item.Data().Name == "Nobody"
		},
	)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyTypedPayload returns false -- Nobody missing", actual)
}

func Test_Cov23_AllTypedPayloads_True(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := corepayload.AllTypedPayloads[testUserCov23](col,
		func(item *corepayload.TypedPayloadWrapper[testUserCov23]) bool {
			return item.Data().Age > 0
		},
	)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllTypedPayloads returns true -- all age>0", actual)
}

func Test_Cov23_AllTypedPayloads_False(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := corepayload.AllTypedPayloads[testUserCov23](col,
		func(item *corepayload.TypedPayloadWrapper[testUserCov23]) bool {
			return item.Data().Age >= 30
		},
	)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllTypedPayloads returns false -- Bob is 25", actual)
}

// ── ConvertTypedPayloads ──

func Test_Cov23_ConvertTypedPayloads_Valid(t *testing.T) {
	// Arrange
	type simpleUser struct {
		Name string `json:"Name"`
	}
	col := makeCollectionCov23()

	// Act
	converted, err := corepayload.ConvertTypedPayloads[testUserCov23, simpleUser](col)

	// Assert
	actual := args.Map{
		"length":  converted.Length(),
		"noError": err == nil,
	}
	expected := args.Map{
		"length":  3,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "ConvertTypedPayloads returns converted -- 3 items", actual)
}

func Test_Cov23_ConvertTypedPayloads_Empty(t *testing.T) {
	// Arrange
	type other struct{ X int }
	col := corepayload.EmptyTypedPayloadCollection[testUserCov23]()

	// Act
	converted, err := corepayload.ConvertTypedPayloads[testUserCov23, other](col)

	// Assert
	actual := args.Map{
		"isEmpty": converted.IsEmpty(),
		"noError": err == nil,
	}
	expected := args.Map{
		"isEmpty": true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "ConvertTypedPayloads returns empty -- empty source", actual)
}

// ── TypedPayloadCollection methods ──

func Test_Cov23_TypedPayloadCollection_ForEachData(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()
	var names []string

	// Act
	col.ForEachData(func(index int, data testUserCov23) {
		names = append(names, data.Name)
	})

	// Assert
	actual := args.Map{"length": len(names)}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "ForEachData iterates all -- 3 items", actual)
}

func Test_Cov23_TypedPayloadCollection_ForEachBreak(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()
	count := 0

	// Act
	col.ForEachBreak(func(index int, item *corepayload.TypedPayloadWrapper[testUserCov23]) bool {
		count++
		return index >= 1
	})

	// Assert
	actual := args.Map{"count": count}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "ForEachBreak stops early -- 2 iterations", actual)
}

func Test_Cov23_TypedPayloadCollection_FilterByData(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := col.FilterByData(func(u testUserCov23) bool {
		return u.Age >= 30
	})

	// Assert
	actual := args.Map{"length": result.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "FilterByData returns filtered -- 2 seniors", actual)
}

func Test_Cov23_TypedPayloadCollection_FirstByData(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := col.FirstByData(func(u testUserCov23) bool {
		return u.Name == "Bob"
	})

	// Assert
	actual := args.Map{
		"notNil": result != nil,
		"name":   result.Data().Name,
	}
	expected := args.Map{
		"notNil": true,
		"name":   "Bob",
	}
	expected.ShouldBeEqual(t, 0, "FirstByData returns Bob -- found", actual)
}

func Test_Cov23_TypedPayloadCollection_FirstByName(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := col.FirstByName("user")

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FirstByName returns item -- name user", actual)
}

func Test_Cov23_TypedPayloadCollection_FirstById(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := col.FirstById("2")

	// Assert
	actual := args.Map{
		"notNil": result != nil,
		"name":   result.Data().Name,
	}
	expected := args.Map{
		"notNil": true,
		"name":   "Bob",
	}
	expected.ShouldBeEqual(t, 0, "FirstById returns Bob -- id 2", actual)
}

func Test_Cov23_TypedPayloadCollection_CountFunc(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := col.CountFunc(func(item *corepayload.TypedPayloadWrapper[testUserCov23]) bool {
		return item.Data().Age >= 30
	})

	// Assert
	actual := args.Map{"count": result}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "CountFunc returns 2 -- seniors", actual)
}

func Test_Cov23_TypedPayloadCollection_SkipTake(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	skipped := col.Skip(1)
	taken := col.Take(2)

	// Assert
	actual := args.Map{
		"skippedLen": len(skipped),
		"takenLen":   len(taken),
	}
	expected := args.Map{
		"skippedLen": 2,
		"takenLen":   2,
	}
	expected.ShouldBeEqual(t, 0, "Skip/Take return correct -- slice ops", actual)
}

func Test_Cov23_TypedPayloadCollection_AllData(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	data := col.AllData()

	// Assert
	actual := args.Map{
		"length": len(data),
		"first":  data[0].Name,
	}
	expected := args.Map{
		"length": 3,
		"first":  "Alice",
	}
	expected.ShouldBeEqual(t, 0, "AllData returns typed slice -- 3 items", actual)
}

func Test_Cov23_TypedPayloadCollection_AllNames(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	names := col.AllNames()

	// Assert
	actual := args.Map{"length": len(names)}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "AllNames returns names -- 3 items", actual)
}

func Test_Cov23_TypedPayloadCollection_AllIdentifiers(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	ids := col.AllIdentifiers()

	// Assert
	actual := args.Map{"length": len(ids), "first": ids[0]}
	expected := args.Map{"length": 3, "first": "1"}
	expected.ShouldBeEqual(t, 0, "AllIdentifiers returns ids -- 3 items", actual)
}

func Test_Cov23_TypedPayloadCollection_ToPayloadsCollection(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	pc := col.ToPayloadsCollection()

	// Assert
	actual := args.Map{"length": pc.Length()}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "ToPayloadsCollection returns PC -- 3 items", actual)
}

func Test_Cov23_TypedPayloadCollection_Clone(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	cloned, err := col.Clone()

	// Assert
	actual := args.Map{
		"length":  cloned.Length(),
		"noError": err == nil,
	}
	expected := args.Map{
		"length":  3,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "Clone returns copy -- 3 items", actual)
}

func Test_Cov23_TypedPayloadCollection_CloneMust(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	cloned := col.CloneMust()

	// Assert
	actual := args.Map{"length": cloned.Length()}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "CloneMust returns copy -- 3 items", actual)
}

func Test_Cov23_TypedPayloadCollection_ConcatNew(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()
	extra := makeTypedWrapperCov23("user", "4", testUserCov23{Name: "Dave", Age: 40})

	// Act
	result, err := col.ConcatNew(extra)

	// Assert
	actual := args.Map{
		"length":  result.Length(),
		"noError": err == nil,
	}
	expected := args.Map{
		"length":  4,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "ConcatNew returns new -- 4 items", actual)
}

func Test_Cov23_TypedPayloadCollection_RemoveAt(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	ok := col.RemoveAt(1)

	// Assert
	actual := args.Map{
		"ok":     ok,
		"length": col.Length(),
	}
	expected := args.Map{
		"ok":     true,
		"length": 2,
	}
	expected.ShouldBeEqual(t, 0, "RemoveAt removes item -- index 1", actual)
}

func Test_Cov23_TypedPayloadCollection_RemoveAt_OutOfBounds(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	ok := col.RemoveAt(99)

	// Assert
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "RemoveAt returns false -- out of bounds", actual)
}

// ── Paging ──

func Test_Cov23_TypedPayloadCollection_GetPagesSize(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	pages := col.GetPagesSize(2)

	// Assert
	actual := args.Map{"pages": pages}
	expected := args.Map{"pages": 2}
	expected.ShouldBeEqual(t, 0, "GetPagesSize returns 2 -- 3 items / 2", actual)
}

func Test_Cov23_TypedPayloadCollection_GetPagedCollection(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	pages := col.GetPagedCollection(2)

	// Assert
	actual := args.Map{
		"pageCount":   len(pages),
		"page1Length": pages[0].Length(),
		"page2Length": pages[1].Length(),
	}
	expected := args.Map{
		"pageCount":   2,
		"page1Length": 2,
		"page2Length": 1,
	}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection returns pages -- 2+1", actual)
}

func Test_Cov23_TypedPayloadCollection_GetPagedCollectionWithInfo(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	pages := col.GetPagedCollectionWithInfo(2)

	// Assert
	actual := args.Map{
		"pageCount":        len(pages),
		"page1CurrentPage": pages[0].Paging.CurrentPageIndex,
		"page1TotalPages":  pages[0].Paging.TotalPages,
		"page1TotalItems":  pages[0].Paging.TotalItems,
	}
	expected := args.Map{
		"pageCount":        2,
		"page1CurrentPage": 1,
		"page1TotalPages":  2,
		"page1TotalItems":  3,
	}
	expected.ShouldBeEqual(t, 0, "GetPagedCollectionWithInfo returns paging -- 2 pages", actual)
}

// ── IsValid / HasErrors / Errors / MergedError ──

func Test_Cov23_TypedPayloadCollection_IsValid(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := col.IsValid()

	// Assert
	actual := args.Map{"isValid": result}
	expected := args.Map{"isValid": true}
	expected.ShouldBeEqual(t, 0, "IsValid returns true -- all parsed", actual)
}

func Test_Cov23_TypedPayloadCollection_HasErrors(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := col.HasErrors()

	// Assert
	actual := args.Map{"hasErrors": result}
	expected := args.Map{"hasErrors": false}
	expected.ShouldBeEqual(t, 0, "HasErrors returns false -- no errors", actual)
}

func Test_Cov23_TypedPayloadCollection_Errors(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	errs := col.Errors()

	// Assert
	actual := args.Map{"length": len(errs)}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "Errors returns empty -- no errors", actual)
}

func Test_Cov23_TypedPayloadCollection_FirstError(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := col.FirstError()

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FirstError returns nil -- no errors", actual)
}

func Test_Cov23_TypedPayloadCollection_MergedError(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := col.MergedError()

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergedError returns nil -- no errors", actual)
}

// ── NewTypedPayloadCollectionSingle / FromData ──

func Test_Cov23_NewTypedPayloadCollectionSingle(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})

	// Act
	col := corepayload.NewTypedPayloadCollectionSingle[testUserCov23](tw)

	// Assert
	actual := args.Map{"length": col.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "NewTypedPayloadCollectionSingle returns 1 -- single item", actual)
}

func Test_Cov23_NewTypedPayloadCollectionSingle_Nil(t *testing.T) {
	// Arrange & Act
	col := corepayload.NewTypedPayloadCollectionSingle[testUserCov23](nil)

	// Assert
	actual := args.Map{"isEmpty": col.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewTypedPayloadCollectionSingle returns empty -- nil", actual)
}

func Test_Cov23_NewTypedPayloadCollectionFromData(t *testing.T) {
	// Arrange
	data := []testUserCov23{
		{Name: "Alice"},
		{Name: "Bob"},
	}

	// Act
	col, err := corepayload.NewTypedPayloadCollectionFromData[testUserCov23]("user", data)

	// Assert
	actual := args.Map{
		"length":  col.Length(),
		"noError": err == nil,
	}
	expected := args.Map{
		"length":  2,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "NewTypedPayloadCollectionFromData returns 2 -- from data", actual)
}

func Test_Cov23_NewTypedPayloadCollectionFromData_Empty(t *testing.T) {
	// Arrange & Act
	col, err := corepayload.NewTypedPayloadCollectionFromData[testUserCov23]("user", []testUserCov23{})

	// Assert
	actual := args.Map{
		"isEmpty": col.IsEmpty(),
		"noError": err == nil,
	}
	expected := args.Map{
		"isEmpty": true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "NewTypedPayloadCollectionFromData returns empty -- no data", actual)
}

func Test_Cov23_NewTypedPayloadCollectionFromDataMust(t *testing.T) {
	// Arrange
	data := []testUserCov23{{Name: "Alice"}}

	// Act
	col := corepayload.NewTypedPayloadCollectionFromDataMust[testUserCov23]("user", data)

	// Assert
	actual := args.Map{"length": col.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "NewTypedPayloadCollectionFromDataMust returns 1 -- from data", actual)
}

// ── TypedPayloadWrapperCreator funcs ──

func Test_Cov23_TypedPayloadWrapperRecord(t *testing.T) {
	// Arrange & Act
	tw, err := corepayload.TypedPayloadWrapperRecord[testUserCov23](
		"user-create", "usr-1", "task", "cat",
		testUserCov23{Name: "Alice"},
	)

	// Assert
	actual := args.Map{
		"name":    tw.Data().Name,
		"noError": err == nil,
	}
	expected := args.Map{
		"name":    "Alice",
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperRecord returns typed -- valid", actual)
}

func Test_Cov23_TypedPayloadWrapperRecords(t *testing.T) {
	// Arrange & Act
	tw, err := corepayload.TypedPayloadWrapperRecords[[]testUserCov23](
		"users", "batch-1", "task", "cat",
		[]testUserCov23{{Name: "A"}, {Name: "B"}},
	)

	// Assert
	actual := args.Map{
		"dataLen": len(tw.Data()),
		"noError": err == nil,
	}
	expected := args.Map{
		"dataLen": 2,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperRecords returns slice -- 2 items", actual)
}

func Test_Cov23_TypedPayloadWrapperNameIdRecord(t *testing.T) {
	// Arrange & Act
	tw, err := corepayload.TypedPayloadWrapperNameIdRecord[testUserCov23](
		"user", "1", testUserCov23{Name: "Alice"},
	)

	// Assert
	actual := args.Map{
		"name":    tw.Data().Name,
		"noError": err == nil,
	}
	expected := args.Map{
		"name":    "Alice",
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperNameIdRecord returns typed -- valid", actual)
}

func Test_Cov23_TypedPayloadWrapperNameIdCategory(t *testing.T) {
	// Arrange & Act
	tw, err := corepayload.TypedPayloadWrapperNameIdCategory[testUserCov23](
		"user", "1", "cat", testUserCov23{Name: "Alice"},
	)

	// Assert
	actual := args.Map{
		"noError": err == nil,
		"catName": tw.CategoryName(),
	}
	expected := args.Map{
		"noError": true,
		"catName": "cat",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperNameIdCategory returns typed -- valid", actual)
}

func Test_Cov23_TypedPayloadWrapperAll(t *testing.T) {
	// Arrange & Act
	tw, err := corepayload.TypedPayloadWrapperAll[testUserCov23](
		"name", "id", "task", "entity", "cat", false,
		testUserCov23{Name: "Alice"}, nil,
	)

	// Assert
	actual := args.Map{
		"noError": err == nil,
		"name":    tw.Data().Name,
	}
	expected := args.Map{
		"noError": true,
		"name":    "Alice",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperAll returns typed -- valid", actual)
}

func Test_Cov23_TypedPayloadWrapperDeserialize(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})
	jsonBytes := tw.SerializeMust()

	// Act
	result, err := corepayload.TypedPayloadWrapperDeserialize[testUserCov23](jsonBytes)

	// Assert
	actual := args.Map{
		"noError": err == nil,
		"name":    result.Data().Name,
	}
	expected := args.Map{
		"noError": true,
		"name":    "Alice",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperDeserialize returns typed -- valid", actual)
}

func Test_Cov23_TypedPayloadWrapperDeserializeUsingJsonResult(t *testing.T) {
	// Arrange
	tw := makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"})
	jsonResult := tw.JsonPtr()

	// Act
	result, err := corepayload.TypedPayloadWrapperDeserializeUsingJsonResult[testUserCov23](jsonResult)

	// Assert
	actual := args.Map{
		"noError": err == nil,
		"name":    result.Data().Name,
	}
	expected := args.Map{
		"noError": true,
		"name":    "Alice",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperDeserializeUsingJsonResult returns typed -- valid", actual)
}

func Test_Cov23_TypedPayloadCollectionDeserialize(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()
	pc := col.ToPayloadsCollection()
	jsonBytes, _ := corejson.Serialize.Raw(pc.Items)

	// Act
	result, err := corepayload.TypedPayloadCollectionDeserialize[testUserCov23](jsonBytes)

	// Assert
	actual := args.Map{
		"length":  result.Length(),
		"noError": err == nil,
	}
	expected := args.Map{
		"length":  3,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollectionDeserialize returns 3 -- from bytes", actual)
}
