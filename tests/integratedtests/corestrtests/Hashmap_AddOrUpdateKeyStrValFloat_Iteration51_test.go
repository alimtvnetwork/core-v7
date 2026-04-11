package corestrtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/smartystreets/goconvey/convey"
)

// =============================================================================
// Hashmap — deeper coverage (AddOrUpdate variants, lock variants, keys/values)
// =============================================================================

func Test_Hashmap_AddOrUpdateKeyStrValFloat(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateKeyStrValFloat", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddOrUpdateKeyStrValFloat("pi", 3.14)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateKeyStrValFloat",
			ExpectedInput: args.Map{
				"HasKey": true,
				"IsEmpty": false,
			},
		}
		actual := args.Map{
			"HasKey": h.Has("pi"),
			"IsEmpty": h.IsEmpty(),
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateKeyStrValFloat64", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddOrUpdateKeyStrValFloat64("e", 2.71828)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateKeyStrValFloat64",
			ExpectedInput: args.Map{"HasKey": true},
		}
		actual := args.Map{"HasKey": h.Has("e")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValAny(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateKeyStrValAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddOrUpdateKeyStrValAny("val", 42)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateKeyStrValAny",
			ExpectedInput: args.Map{"HasKey": true},
		}
		actual := args.Map{"HasKey": h.Has("val")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyValueAny(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateKeyValueAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		pair := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		// Act
		h.AddOrUpdateKeyValueAny(pair)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateKeyValueAny",
			ExpectedInput: args.Map{
				"Has": true,
				"Length": 1,
			},
		}
		actual := args.Map{
			"Has": h.Has("k"),
			"Length": h.Length(),
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyVal(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateKeyVal", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}
		// Act
		isNew := h.AddOrUpdateKeyVal(kv)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateKeyVal new",
			ExpectedInput: args.Map{"IsNew": true},
		}
		actual := args.Map{"IsNew": isNew}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyVal_Existing(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateKeyVal_Existing", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "old")
		kv := corestr.KeyValuePair{Key: "a", Value: "new"}
		// Act
		isNew := h.AddOrUpdateKeyVal(kv)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateKeyVal existing",
			ExpectedInput: args.Map{"IsNew": false},
		}
		actual := args.Map{"IsNew": isNew}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyValues(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateKeyValues", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddOrUpdateKeyValues(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateKeyValues",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyValues_Empty(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateKeyValues_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddOrUpdateKeyValues()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateKeyValues empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyAnyValues(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateKeyAnyValues", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddOrUpdateKeyAnyValues(
			corestr.KeyAnyValuePair{Key: "x", Value: 10},
		)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateKeyAnyValues",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": h.Has("x")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyAnyValues_Empty(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateKeyAnyValues_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddOrUpdateKeyAnyValues()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateKeyAnyValues empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddOrUpdateLock("k", "v")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateLock",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": h.Has("k")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateWithWgLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		// Act
		h.AddOrUpdateWithWgLock("wg", "val", wg)
		wg.Wait()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateWithWgLock",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": h.Has("wg")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateStringsPtrWgLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		keys := []string{"a", "b"}
		vals := []string{"1", "2"}
		// Act
		h.AddOrUpdateStringsPtrWgLock(wg, keys, vals)
		wg.Wait()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateStringsPtrWgLock",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateStringsPtrWgLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateStringsPtrWgLock_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		// Act
		h.AddOrUpdateStringsPtrWgLock(wg, []string{}, []string{})
		wg.Wait()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateStringsPtrWgLock empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateHashmap", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		other := corestr.New.Hashmap.Empty()
		other.AddOrUpdate("x", "y")
		// Act
		h.AddOrUpdateHashmap(other)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateHashmap",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": h.Has("x")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateHashmap_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddOrUpdateHashmap(nil)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateHashmap nil",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateMap", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddOrUpdateMap(map[string]string{"k": "v"})
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateMap",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": h.Has("k")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateMap_Empty(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateMap_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddOrUpdateMap(map[string]string{})
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateMap empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		// Act
		h.AddOrUpdateCollection(keys, vals)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateCollection",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateCollection_NilKeys(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateCollection_NilKeys", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		vals := corestr.New.Collection.Strings([]string{"1"})
		// Act
		h.AddOrUpdateCollection(nil, vals)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateCollection nil keys",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddOrUpdateCollection_Mismatch(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddOrUpdateCollection_Mismatch", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1"})
		// Act
		h.AddOrUpdateCollection(keys, vals)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddOrUpdateCollection mismatch",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddsOrUpdates", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddsOrUpdates(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsOrUpdates",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddsOrUpdates_Nil(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddsOrUpdates_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		h.AddsOrUpdates()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsOrUpdates nil",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddsOrUpdatesAnyUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return fmt.Sprintf("%v", pair.Value), true, false
		}
		// Act
		h.AddsOrUpdatesAnyUsingFilter(filter, corestr.KeyAnyValuePair{Key: "k", Value: 1})
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsOrUpdatesAnyUsingFilter",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": h.Has("k")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddsOrUpdatesAnyUsingFilter_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return "", false, true
		}
		// Act
		h.AddsOrUpdatesAnyUsingFilter(filter, corestr.KeyAnyValuePair{Key: "k", Value: 1})
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsOrUpdatesAnyUsingFilter break",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": h.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddsOrUpdatesAnyUsingFilterLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return "val", true, false
		}
		// Act
		h.AddsOrUpdatesAnyUsingFilterLock(filter, corestr.KeyAnyValuePair{Key: "k", Value: 1})
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsOrUpdatesAnyUsingFilterLock",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": h.Has("k")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_AddsOrUpdatesUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Value, true, false
		}
		// Act
		h.AddsOrUpdatesUsingFilter(filter, corestr.KeyValuePair{Key: "k", Value: "v"})
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsOrUpdatesUsingFilter",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": h.Has("k")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ConcatNew_WithArgs(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ConcatNew_WithArgs", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Empty()
		other.AddOrUpdate("b", "2")
		// Act
		result := h.ConcatNew(false, other)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ConcatNew with args",
			ExpectedInput: args.Map{
				"HasA": true,
				"HasB": true,
			},
		}
		actual := args.Map{
			"HasA": result.Has("a"),
			"HasB": result.Has("b"),
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ConcatNew_NoArgs(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ConcatNew_NoArgs", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		result := h.ConcatNew(true)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ConcatNew no args",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result.Has("a")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ConcatNewUsingMaps", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		result := h.ConcatNewUsingMaps(false, map[string]string{"b": "2"})
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ConcatNewUsingMaps",
			ExpectedInput: args.Map{"HasB": true},
		}
		actual := args.Map{"HasB": result.Has("b")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ConcatNewUsingMaps_NoArgs(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ConcatNewUsingMaps_NoArgs", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		result := h.ConcatNewUsingMaps(true)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ConcatNewUsingMaps no args",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result.Has("a")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_HasAny(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_HasAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		found := h.HasAny("b", "a")
		notFound := h.HasAny("c", "d")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "HasAny",
			ExpectedInput: args.Map{
				"Found": true,
				"NotFound": false,
			},
		}
		actual := args.Map{
			"Found": found,
			"NotFound": notFound,
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_HasWithLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_HasWithLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		// Act
		result := h.HasWithLock("k")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "HasWithLock",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_HasAllCollectionItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		// Act
		result := h.HasAllCollectionItems(col)
		resultNil := h.HasAllCollectionItems(nil)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "HasAllCollectionItems",
			ExpectedInput: args.Map{
				"HasAll": true,
				"Nil": false,
			},
		}
		actual := args.Map{
			"HasAll": result,
			"Nil": resultNil,
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ContainsLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ContainsLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		// Act
		result := h.ContainsLock("k")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ContainsLock",
			ExpectedInput: args.Map{"Contains": true},
		}
		actual := args.Map{"Contains": result}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_IsKeyMissing(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_IsKeyMissing", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		// Act
		missing := h.IsKeyMissing("other")
		notMissing := h.IsKeyMissing("k")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "IsKeyMissing",
			ExpectedInput: args.Map{
				"Missing": true,
				"NotMissing": false,
			},
		}
		actual := args.Map{
			"Missing": missing,
			"NotMissing": notMissing,
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_IsKeyMissingLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_IsKeyMissingLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		// Act
		result := h.IsKeyMissingLock("other")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "IsKeyMissingLock",
			ExpectedInput: args.Map{"Missing": true},
		}
		actual := args.Map{"Missing": result}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_Diff(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_Diff", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Empty()
		other.AddOrUpdate("a", "2")
		// Act
		diff := h.Diff(other)
		// Assert
		convey.Convey("Diff returns non-nil", t, func() {
			convey.So(diff, convey.ShouldNotBeNil)
		})
	})
}

func Test_Hashmap_KeysValuesCollection(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_KeysValuesCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		keys, values := h.KeysValuesCollection()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "KeysValuesCollection",
			ExpectedInput: args.Map{
				"KeysLen": 1,
				"ValsLen": 1,
			},
		}
		actual := args.Map{
			"KeysLen": keys.Length(),
			"ValsLen": values.Length(),
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_KeysValuesList(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_KeysValuesList", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		keys, values := h.KeysValuesList()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "KeysValuesList",
			ExpectedInput: args.Map{
				"KeysLen": 1,
				"ValsLen": 1,
			},
		}
		actual := args.Map{
			"KeysLen": len(keys),
			"ValsLen": len(values),
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_KeysValuePairs(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_KeysValuePairs", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		pairs := h.KeysValuePairs()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "KeysValuePairs",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(pairs)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_KeysValuePairsCollection(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_KeysValuePairsCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		col := h.KeysValuePairsCollection()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "KeysValuePairsCollection",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": col.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_KeysValuesListLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_KeysValuesListLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		keys, values := h.KeysValuesListLock()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "KeysValuesListLock",
			ExpectedInput: args.Map{
				"KeysLen": 1,
				"ValsLen": 1,
			},
		}
		actual := args.Map{
			"KeysLen": len(keys),
			"ValsLen": len(values),
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_KeysLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_KeysLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		keys := h.KeysLock()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "KeysLock",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(keys)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ItemsCopyLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ItemsCopyLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		items := h.ItemsCopyLock()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ItemsCopyLock",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(*items)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ValuesCollection(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ValuesCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		col := h.ValuesCollection()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ValuesCollection",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": col.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ValuesHashset(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ValuesHashset", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		hs := h.ValuesHashset()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ValuesHashset",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ValuesCollectionLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ValuesCollectionLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		col := h.ValuesCollectionLock()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ValuesCollectionLock",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": col.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ValuesHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ValuesHashsetLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		hs := h.ValuesHashsetLock()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ValuesHashsetLock",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ValuesToLower(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ValuesToLower", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("KEY", "VAL")
		// Act
		result := h.ValuesToLower()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ValuesToLower (deprecated alias)",
			ExpectedInput: args.Map{"HasLower": true},
		}
		actual := args.Map{"HasLower": result.Has("key")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_StringLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		result := h.StringLock()
		// Assert
		convey.Convey("StringLock non-empty", t, func() {
			convey.So(result, convey.ShouldNotBeEmpty)
		})
	})
}

func Test_Hashmap_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_StringLock_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		result := h.StringLock()
		// Assert
		convey.Convey("StringLock empty", t, func() {
			convey.So(result, convey.ShouldContainSubstring, "No Element")
		})
	})
}

func Test_Hashmap_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetKeysFilteredItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("abc", "1")
		h.AddOrUpdate("def", "2")
		filter := func(s string, i int) (string, bool, bool) {
			return s, s == "abc", false
		}
		// Act
		result := h.GetKeysFilteredItems(filter)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetKeysFilteredItems",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_GetKeysFilteredItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetKeysFilteredItems_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, false
		}
		// Act
		result := h.GetKeysFilteredItems(filter)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetKeysFilteredItems empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_GetKeysFilteredItems_Break(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetKeysFilteredItems_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, true
		}
		// Act
		result := h.GetKeysFilteredItems(filter)
		// Assert
		convey.Convey("GetKeysFilteredItems break", t, func() {
			convey.So(len(result), convey.ShouldBeGreaterThan, 0)
		})
	})
}

func Test_Hashmap_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetKeysFilteredCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("abc", "1")
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, false
		}
		// Act
		result := h.GetKeysFilteredCollection(filter)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetKeysFilteredCollection",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": result.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_GetKeysFilteredCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetKeysFilteredCollection_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, false
		}
		// Act
		result := h.GetKeysFilteredCollection(filter)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetKeysFilteredCollection empty",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": result.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_GetKeysFilteredCollection_Break(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetKeysFilteredCollection_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, true
		}
		// Act
		result := h.GetKeysFilteredCollection(filter)
		// Assert
		convey.Convey("GetKeysFilteredCollection break", t, func() {
			convey.So(result.Length(), convey.ShouldBeGreaterThan, 0)
		})
	})
}

func Test_Hashmap_SafeItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_SafeItems_Nil", func() {
		// Arrange
		var h *corestr.Hashmap
		// Act
		result := h.SafeItems()
		// Assert
		convey.Convey("SafeItems nil", t, func() {
			convey.So(result, convey.ShouldBeNil)
		})
	})
}

func Test_Hashmap_GetValue(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetValue", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		// Act
		val, found := h.GetValue("k")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetValue",
			ExpectedInput: args.Map{
				"Val": "v",
				"Found": true,
			},
		}
		actual := args.Map{
			"Val": val,
			"Found": found,
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ToError(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ToError", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		// Act
		err := h.ToError(", ")
		// Assert
		convey.Convey("ToError", t, func() {
			convey.So(err, convey.ShouldNotBeNil)
		})
	})
}

func Test_Hashmap_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ToDefaultError", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		// Act
		err := h.ToDefaultError()
		// Assert
		convey.Convey("ToDefaultError", t, func() {
			convey.So(err, convey.ShouldNotBeNil)
		})
	})
}

func Test_Hashmap_KeyValStringLines(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_KeyValStringLines", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		// Act
		lines := h.KeyValStringLines()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "KeyValStringLines",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(lines)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ToStringsUsingCompiler_Empty(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ToStringsUsingCompiler_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		compiler := func(k, v string) string { return k + "=" + v }
		// Act
		result := h.ToStringsUsingCompiler(compiler)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ToStringsUsingCompiler empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_GetValuesExceptKeysInHashset(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetValuesExceptKeysInHashset", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		exclude := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := h.GetValuesExceptKeysInHashset(exclude)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetValuesExceptKeysInHashset",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_GetValuesExceptKeysInHashset_Nil(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetValuesExceptKeysInHashset_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		result := h.GetValuesExceptKeysInHashset(nil)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetValuesExceptKeysInHashset nil",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_GetValuesKeysExcept(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetValuesKeysExcept", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		// Act
		result := h.GetValuesKeysExcept([]string{"a"})
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetValuesKeysExcept",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_GetValuesKeysExcept_Nil(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetValuesKeysExcept_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		result := h.GetValuesKeysExcept(nil)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetValuesKeysExcept nil",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetAllExceptCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		col := corestr.New.Collection.Strings([]string{"a"})
		// Act
		result := h.GetAllExceptCollection(col)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetAllExceptCollection",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_GetAllExceptCollection_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		result := h.GetAllExceptCollection(nil)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetAllExceptCollection nil",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ClonePtr_Nil", func() {
		// Arrange
		var h *corestr.Hashmap
		// Act
		result := h.ClonePtr()
		// Assert
		convey.Convey("ClonePtr nil", t, func() {
			convey.So(result, convey.ShouldBeNil)
		})
	})
}

func Test_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_IsEqual", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Empty()
		other.AddOrUpdate("a", "1")
		// Act
		result := h.IsEqual(*other)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "IsEqual",
			ExpectedInput: args.Map{"IsEqual": true},
		}
		actual := args.Map{"IsEqual": result}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_IsEqualPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_IsEqualPtrLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		result := h.IsEqualPtrLock(h)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "IsEqualPtrLock same ptr",
			ExpectedInput: args.Map{"IsEqual": true},
		}
		actual := args.Map{"IsEqual": result}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_Remove(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_Remove", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		h.Remove("a")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Remove",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": h.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_RemoveWithLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		h.RemoveWithLock("a")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "RemoveWithLock",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": h.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

// =============================================================================
// emptyCreator — all factory methods
// =============================================================================

func Test_EmptyCreator_LinkedList(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_LinkedList", func() {
		// Arrange // Act
		ll := corestr.Empty.LinkedList()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Empty.LinkedList",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": ll.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_EmptyCreator_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_SimpleSlice", func() {
		// Arrange // Act
		ss := corestr.Empty.SimpleSlice()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Empty.SimpleSlice",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": ss.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_EmptyCreator_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_KeyAnyValuePair", func() {
		// Arrange // Act
		p := corestr.Empty.KeyAnyValuePair()
		// Assert
		convey.Convey("Empty.KeyAnyValuePair", t, func() {
			convey.So(p, convey.ShouldNotBeNil)
			convey.So(p.Key, convey.ShouldBeEmpty)
		})
	})
}

func Test_EmptyCreator_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_KeyValuePair", func() {
		// Arrange // Act
		p := corestr.Empty.KeyValuePair()
		// Assert
		convey.Convey("Empty.KeyValuePair", t, func() {
			convey.So(p, convey.ShouldNotBeNil)
			convey.So(p.Key, convey.ShouldBeEmpty)
		})
	})
}

func Test_EmptyCreator_KeyValueCollection(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_KeyValueCollection", func() {
		// Arrange // Act
		c := corestr.Empty.KeyValueCollection()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Empty.KeyValueCollection",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": c.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_EmptyCreator_LinkedCollections(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_LinkedCollections", func() {
		// Arrange // Act
		lc := corestr.Empty.LinkedCollections()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Empty.LinkedCollections",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": lc.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_EmptyCreator_LeftRight(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_LeftRight", func() {
		// Arrange // Act
		lr := corestr.Empty.LeftRight()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Empty.LeftRight",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": lr.IsLeftEmpty() && lr.IsRightEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_EmptyCreator_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_Hashset", func() {
		// Arrange // Act
		hs := corestr.Empty.Hashset()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Empty.Hashset",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": hs.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_EmptyCreator_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_HashsetsCollection", func() {
		// Arrange // Act
		hc := corestr.Empty.HashsetsCollection()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Empty.HashsetsCollection",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": hc.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_EmptyCreator_Hashmap(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_Hashmap", func() {
		// Arrange // Act
		hm := corestr.Empty.Hashmap()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Empty.Hashmap",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": hm.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_EmptyCreator_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_CharCollectionMap", func() {
		// Arrange // Act
		ccm := corestr.Empty.CharCollectionMap()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Empty.CharCollectionMap",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": ccm.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_EmptyCreator_KeyValuesCollection(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_KeyValuesCollection", func() {
		// Arrange // Act
		kvc := corestr.Empty.KeyValuesCollection()
		// Assert
		convey.Convey("Empty.KeyValuesCollection", t, func() {
			convey.So(kvc, convey.ShouldNotBeNil)
		})
	})
}

func Test_EmptyCreator_CollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_CollectionsOfCollection", func() {
		// Arrange // Act
		coc := corestr.Empty.CollectionsOfCollection()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Empty.CollectionsOfCollection",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": coc.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_EmptyCreator_CharHashsetMap(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_CharHashsetMap", func() {
		// Arrange // Act
		chm := corestr.Empty.CharHashsetMap()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Empty.CharHashsetMap",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": chm.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_EmptyCreator_SimpleStringOnce(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_SimpleStringOnce", func() {
		// Arrange // Act
		sso := corestr.Empty.SimpleStringOnce()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Empty.SimpleStringOnce",
			ExpectedInput: args.Map{"IsInit": false},
		}
		actual := args.Map{"IsInit": sso.IsInitialized()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_EmptyCreator_SimpleStringOncePtr(t *testing.T) {
	safeTest(t, "Test_Cov55_EmptyCreator_SimpleStringOncePtr", func() {
		// Arrange // Act
		sso := corestr.Empty.SimpleStringOncePtr()
		// Assert
		convey.Convey("Empty.SimpleStringOncePtr", t, func() {
			convey.So(sso, convey.ShouldNotBeNil)
			convey.So(sso.IsInitialized(), convey.ShouldBeFalse)
		})
	})
}

// =============================================================================
// StringUtils — WrapDouble, WrapSingle, WrapTilda, WrapDoubleIfMissing, WrapSingleIfMissing
// =============================================================================

func Test_StringUtils_WrapDouble(t *testing.T) {
	safeTest(t, "Test_Cov55_StringUtils_WrapDouble", func() {
		// Arrange // Act
		result := corestr.StringUtils.WrapDouble("hello")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapDouble",
			ExpectedInput: "\"hello\"",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_StringUtils_WrapSingle(t *testing.T) {
	safeTest(t, "Test_Cov55_StringUtils_WrapSingle", func() {
		// Arrange // Act
		result := corestr.StringUtils.WrapSingle("hello")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapSingle",
			ExpectedInput: "'hello'",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_StringUtils_WrapTilda(t *testing.T) {
	safeTest(t, "Test_Cov55_StringUtils_WrapTilda", func() {
		// Arrange // Act
		result := corestr.StringUtils.WrapTilda("hello")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapTilda",
			ExpectedInput: "`hello`",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_StringUtils_WrapDoubleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_Cov55_StringUtils_WrapDoubleIfMissing_Empty", func() {
		// Arrange // Act
		result := corestr.StringUtils.WrapDoubleIfMissing("")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapDoubleIfMissing empty",
			ExpectedInput: "\"\"",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_StringUtils_WrapDoubleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_Cov55_StringUtils_WrapDoubleIfMissing_AlreadyWrapped", func() {
		// Arrange // Act
		result := corestr.StringUtils.WrapDoubleIfMissing("\"hello\"")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapDoubleIfMissing already",
			ExpectedInput: "\"hello\"",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_StringUtils_WrapDoubleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_Cov55_StringUtils_WrapDoubleIfMissing_NotWrapped", func() {
		// Arrange // Act
		result := corestr.StringUtils.WrapDoubleIfMissing("hello")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapDoubleIfMissing not wrapped",
			ExpectedInput: "\"hello\"",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_StringUtils_WrapSingleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_Cov55_StringUtils_WrapSingleIfMissing_Empty", func() {
		// Arrange // Act
		result := corestr.StringUtils.WrapSingleIfMissing("")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapSingleIfMissing empty",
			ExpectedInput: "''",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_StringUtils_WrapSingleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_Cov55_StringUtils_WrapSingleIfMissing_AlreadyWrapped", func() {
		// Arrange // Act
		result := corestr.StringUtils.WrapSingleIfMissing("'hello'")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapSingleIfMissing already",
			ExpectedInput: "'hello'",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_StringUtils_WrapSingleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_Cov55_StringUtils_WrapSingleIfMissing_NotWrapped", func() {
		// Arrange // Act
		result := corestr.StringUtils.WrapSingleIfMissing("hello")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapSingleIfMissing not wrapped",
			ExpectedInput: "'hello'",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_StringUtils_WrapDoubleIfMissing_DoubleEmpty(t *testing.T) {
	safeTest(t, "Test_Cov55_StringUtils_WrapDoubleIfMissing_DoubleEmpty", func() {
		// Arrange // Act
		result := corestr.StringUtils.WrapDoubleIfMissing("\"\"")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapDoubleIfMissing double-empty",
			ExpectedInput: "\"\"",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_StringUtils_WrapSingleIfMissing_SingleEmpty(t *testing.T) {
	safeTest(t, "Test_Cov55_StringUtils_WrapSingleIfMissing_SingleEmpty", func() {
		// Arrange // Act
		result := corestr.StringUtils.WrapSingleIfMissing("''")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapSingleIfMissing single-empty",
			ExpectedInput: "''",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

// =============================================================================
// Hashmap — DiffRaw, Collection, SetTrim, SetBySplitter
// =============================================================================

func Test_Hashmap_DiffRaw(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_DiffRaw", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		result := h.DiffRaw(map[string]string{"a": "2"})
		// Assert
		convey.Convey("DiffRaw", t, func() {
			convey.So(result, convey.ShouldNotBeNil)
		})
	})
}

func Test_Hashmap_Collection(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_Collection", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		col := h.Collection()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Hashmap.Collection",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": col.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_SetTrim(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_SetTrim", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		isNew := h.SetTrim("  key  ", "  val  ")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "SetTrim",
			ExpectedInput: args.Map{
				"IsNew": true,
				"Has": true,
			},
		}
		actual := args.Map{
			"IsNew": isNew,
			"Has": h.Has("key"),
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_SetBySplitter", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		isNew := h.SetBySplitter("=", "key=value")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "SetBySplitter",
			ExpectedInput: args.Map{
				"IsNew": true,
				"Has": true,
			},
		}
		actual := args.Map{
			"IsNew": isNew,
			"Has": h.Has("key"),
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_SetBySplitter_NoSplit(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_SetBySplitter_NoSplit", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		isNew := h.SetBySplitter("=", "nosplit")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "SetBySplitter no split",
			ExpectedInput: args.Map{
				"IsNew": true,
				"Has": true,
			},
		}
		actual := args.Map{
			"IsNew": isNew,
			"Has": h.Has("nosplit"),
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_KeysCollection(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_KeysCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		// Act
		col := h.KeysCollection()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "KeysCollection",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": col.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_ParseInjectUsingJsonMust", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		jsonResult := h.JsonPtr()
		target := corestr.New.Hashmap.Empty()
		// Act
		result := target.ParseInjectUsingJsonMust(jsonResult)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ParseInjectUsingJsonMust",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result.Has("k")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Hashmap_Clone_Empty(t *testing.T) {
	safeTest(t, "Test_Cov55_Hashmap_Clone_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		// Act
		cloned := h.Clone()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Clone empty",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": cloned.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}
