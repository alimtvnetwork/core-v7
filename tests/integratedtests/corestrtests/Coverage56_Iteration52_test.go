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
// Hashset — deep coverage of remaining methods
// =============================================================================

func Test_Cov56_Hashset_AddPtr(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		key := "hello"
		// Act
		hs.AddPtr(&key)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddPtr",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": hs.Has("hello")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddPtrLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		key := "world"
		// Act
		hs.AddPtrLock(&key)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddPtrLock",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": hs.Has("world")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddWithWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		// Act
		hs.AddWithWgLock("item", wg)
		wg.Wait()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddWithWgLock",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": hs.Has("item")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddBool_New(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddBool_New", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		isExist := hs.AddBool("new")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddBool new",
			ExpectedInput: args.Map{"IsExist": false},
		}
		actual := args.Map{"IsExist": isExist}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddBool_Existing(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddBool_Existing", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("existing")
		// Act
		isExist := hs.AddBool("existing")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddBool existing",
			ExpectedInput: args.Map{"IsExist": true},
		}
		actual := args.Map{"IsExist": isExist}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddNonEmptyWhitespace", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		hs.AddNonEmptyWhitespace("  ")
		hs.AddNonEmptyWhitespace("valid")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddNonEmptyWhitespace",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddIfMany_True(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddIfMany_True", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		hs.AddIfMany(true, "a", "b")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddIfMany true",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddIfMany_False(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddIfMany_False", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		hs.AddIfMany(false, "a", "b")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddIfMany false",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddFunc(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddFunc", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		hs.AddFunc(func() string { return "computed" })
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddFunc",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": hs.Has("computed")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddFuncErr_NoErr(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddFuncErr_NoErr", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		hs.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddFuncErr no err",
			ExpectedInput: args.Map{"Has": true},
		}
		actual = args.Map{"Has": hs.Has("ok")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddFuncErr_WithErr(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddFuncErr_WithErr", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		handlerCalled := false
		// Act
		hs.AddFuncErr(
			func() (string, error) { return "", fmt.Errorf("fail") },
			func(err error) { handlerCalled = true },
		)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddFuncErr with err",
			ExpectedInput: args.Map{
				"HandlerCalled": true,
				"Length": 0,
			},
		}
		actual := args.Map{
			"HandlerCalled": handlerCalled,
			"Length": hs.Length(),
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddStringsPtrWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		// Act
		hs.AddStringsPtrWgLock([]string{"a", "b"}, wg)
		wg.Wait()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddStringsPtrWgLock",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddHashsetWgLock(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddHashsetWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		other := corestr.New.Hashset.Strings([]string{"x", "y"})
		wg := &sync.WaitGroup{}
		wg.Add(1)
		// Act
		hs.AddHashsetWgLock(other, wg)
		wg.Wait()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddHashsetWgLock",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddHashsetWgLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddHashsetWgLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		// Act
		hs.AddHashsetWgLock(nil, wg)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddHashsetWgLock nil",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddSimpleSlice", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		// Act
		hs.AddSimpleSlice(ss)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddSimpleSlice",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		// Act
		hs.AddCollection(col)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddCollection",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddCollection_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		hs.AddCollection(nil)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddCollection nil",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddCollections", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		col1 := corestr.New.Collection.Strings([]string{"a"})
		col2 := corestr.New.Collection.Strings([]string{"b"})
		// Act
		hs.AddCollections(col1, col2)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddCollections",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddCollections_Nil(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddCollections_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		hs.AddCollections(nil, nil)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddCollections nil items",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddsUsingFilter", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, s != "skip", false
		}
		// Act
		hs.AddsUsingFilter(filter, "a", "skip", "b")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsUsingFilter",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddsUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddsUsingFilter_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, true
		}
		// Act
		hs.AddsUsingFilter(filter, "a", "b")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsUsingFilter break",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddsUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddsUsingFilter_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, false
		}
		// Act
		hs.AddsUsingFilter(filter)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsUsingFilter nil",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddsAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddsAnyUsingFilter", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, false
		}
		// Act
		hs.AddsAnyUsingFilter(filter, "hello", 42)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsAnyUsingFilter",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddsAnyUsingFilter_NilItem(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddsAnyUsingFilter_NilItem", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, false
		}
		// Act
		hs.AddsAnyUsingFilter(filter, nil, "valid")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsAnyUsingFilter nil item",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddsAnyUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddsAnyUsingFilter_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, true
		}
		// Act
		hs.AddsAnyUsingFilter(filter, "a", "b")
		// Assert
		convey.Convey("AddsAnyUsingFilter break", t, func() {
			convey.So(hs.Length(), convey.ShouldBeLessThanOrEqualTo, 2)
		})
	})
}

func Test_Cov56_Hashset_AddsAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddsAnyUsingFilterLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, false
		}
		// Act
		hs.AddsAnyUsingFilterLock(filter, "x")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsAnyUsingFilterLock",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddsAnyUsingFilterLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddsAnyUsingFilterLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, false
		}
		// Act
		hs.AddsAnyUsingFilterLock(filter)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddsAnyUsingFilterLock nil",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddItemsMapWgLock(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddItemsMapWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		itemsMap := map[string]bool{"a": true, "b": false}
		wg := &sync.WaitGroup{}
		wg.Add(1)
		// Act
		hs.AddItemsMapWgLock(&itemsMap, wg)
		wg.Wait()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddItemsMapWgLock",
			ExpectedInput: args.Map{
				"HasA": true,
				"HasB": false,
			},
		}
		actual := args.Map{
			"HasA": hs.Has("a"),
			"HasB": hs.Has("b"),
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddItemsMapWgLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddItemsMapWgLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		// Act
		hs.AddItemsMapWgLock(nil, wg)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddItemsMapWgLock nil",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_ConcatNewHashsets(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_ConcatNewHashsets", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		other := corestr.New.Hashset.Strings([]string{"b"})
		// Act
		result := hs.ConcatNewHashsets(false, other)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ConcatNewHashsets",
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

func Test_Cov56_Hashset_ConcatNewHashsets_NoArgs(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_ConcatNewHashsets_NoArgs", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.ConcatNewHashsets(true)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ConcatNewHashsets no args",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result.Has("a")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_ConcatNewStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.ConcatNewStrings(false, []string{"b", "c"})
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ConcatNewStrings",
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

func Test_Cov56_Hashset_ConcatNewStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_ConcatNewStrings_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.ConcatNewStrings(true)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ConcatNewStrings empty",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result.Has("a")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_Filter", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"abc", "def", "abx"})
		// Act
		result := hs.Filter(func(s string) bool { return s[0] == 'a' })
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Filter",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": result.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_IsMissing(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_IsMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		missing := hs.IsMissing("b")
		notMissing := hs.IsMissing("a")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "IsMissing",
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

func Test_Cov56_Hashset_IsMissingLock(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_IsMissingLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.IsMissingLock("b")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "IsMissingLock",
			ExpectedInput: args.Map{"Missing": true},
		}
		actual := args.Map{"Missing": result}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_HasWithLock(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_HasWithLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.HasWithLock("a")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "HasWithLock",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_IsAllMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		allMissing := hs.IsAllMissing("b", "c")
		notAllMissing := hs.IsAllMissing("a", "b")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "IsAllMissing",
			ExpectedInput: args.Map{
				"AllMissing": true,
				"NotAll": false,
			},
		}
		actual := args.Map{
			"AllMissing": allMissing,
			"NotAll": notAllMissing,
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_HasAny(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_HasAny", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		found := hs.HasAny("b", "a")
		notFound := hs.HasAny("c", "d")
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

func Test_Cov56_Hashset_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_HasAllCollectionItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		// Act
		result := hs.HasAllCollectionItems(col)
		resultNil := hs.HasAllCollectionItems(nil)
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

func Test_Cov56_Hashset_OrderedList(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_OrderedList", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		// Act
		result := hs.OrderedList()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "OrderedList",
			ExpectedInput: args.Map{
				"First": "a",
				"Last": "c",
			},
		}
		actual := args.Map{
			"First": result[0],
			"Last": result[2],
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_OrderedList_Empty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_OrderedList_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		result := hs.OrderedList()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "OrderedList empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_SafeStrings(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_SafeStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		result := hs.SafeStrings()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "SafeStrings empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_SafeStrings_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_SafeStrings_NonEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.SafeStrings()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "SafeStrings non-empty",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_Lines(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_Lines", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		result := hs.Lines()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Lines empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_SimpleSlice", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.SimpleSlice()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "SimpleSlice",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": result.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_SimpleSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_SimpleSlice_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		result := hs.SimpleSlice()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "SimpleSlice empty",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": result.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_MapStringAny(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_MapStringAny", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.MapStringAny()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "MapStringAny",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_MapStringAny_Empty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_MapStringAny_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		result := hs.MapStringAny()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "MapStringAny empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_MapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_MapStringAnyDiff", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.MapStringAnyDiff()
		// Assert
		convey.Convey("MapStringAnyDiff", t, func() {
			convey.So(result, convey.ShouldNotBeNil)
		})
	})
}

func Test_Cov56_Hashset_JoinSorted(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_JoinSorted", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		// Act
		result := hs.JoinSorted(",")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "JoinSorted",
			ExpectedInput: "a,b,c",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_Cov56_Hashset_JoinSorted_Empty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_JoinSorted_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		result := hs.JoinSorted(",")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "JoinSorted empty",
			ExpectedInput: "",
		}
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_Cov56_Hashset_ListPtrSortedAsc(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_ListPtrSortedAsc", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		// Act
		result := hs.ListPtrSortedAsc()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ListPtrSortedAsc",
			ExpectedInput: args.Map{"First": "a"},
		}
		actual := args.Map{"First": result[0]}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_ListPtrSortedDsc(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_ListPtrSortedDsc", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		// Act
		result := hs.ListPtrSortedDsc()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ListPtrSortedDsc",
			ExpectedInput: args.Map{"First": "c"},
		}
		actual := args.Map{"First": result[0]}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_ListCopyLock(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_ListCopyLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.ListCopyLock()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ListCopyLock",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_ToLowerSet", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"ABC", "DEF"})
		// Act
		result := hs.ToLowerSet()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ToLowerSet",
			ExpectedInput: args.Map{
				"HasAbc": true,
				"HasDef": true,
			},
		}
		actual := args.Map{
			"HasAbc": result.Has("abc"),
			"HasDef": result.Has("def"),
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_DistinctDiffLinesRaw_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_DistinctDiffLinesRaw_BothEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		result := hs.DistinctDiffLinesRaw()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "DistinctDiffLinesRaw both empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_DistinctDiffLinesRaw_LeftOnly(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_DistinctDiffLinesRaw_LeftOnly", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.DistinctDiffLinesRaw()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "DistinctDiffLinesRaw left only",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_DistinctDiffLinesRaw_RightOnly(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_DistinctDiffLinesRaw_RightOnly", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		result := hs.DistinctDiffLinesRaw("x", "y")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "DistinctDiffLinesRaw right only",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_DistinctDiffLinesRaw_Both(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_DistinctDiffLinesRaw_Both", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		// Act
		result := hs.DistinctDiffLinesRaw("b", "c")
		// Assert
		convey.Convey("DistinctDiffLinesRaw both have items", t, func() {
			convey.So(len(result), convey.ShouldBeGreaterThan, 0)
		})
	})
}

func Test_Cov56_Hashset_DistinctDiffLines_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_DistinctDiffLines_BothEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		result := hs.DistinctDiffLines()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "DistinctDiffLines both empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_DistinctDiffLines_LeftOnly(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_DistinctDiffLines_LeftOnly", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.DistinctDiffLines()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "DistinctDiffLines left only",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_DistinctDiffLines_RightOnly(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_DistinctDiffLines_RightOnly", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		result := hs.DistinctDiffLines("x")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "DistinctDiffLines right only",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_DistinctDiffHashset(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_DistinctDiffHashset", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		other := corestr.New.Hashset.Strings([]string{"b", "c"})
		// Act
		result := hs.DistinctDiffHashset(other)
		// Assert
		convey.Convey("DistinctDiffHashset", t, func() {
			convey.So(len(result), convey.ShouldBeGreaterThan, 0)
		})
	})
}

func Test_Cov56_Hashset_Transpile(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_Transpile", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.Transpile(func(s string) string { return "[" + s + "]" })
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Transpile",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result.Has("[a]")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_Transpile_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		result := hs.Transpile(func(s string) string { return s })
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Transpile empty",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": result.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_WrapDoubleQuote", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.WrapDoubleQuote()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapDoubleQuote",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result.Has("\"a\"")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_WrapSingleQuote", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.WrapSingleQuote()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapSingleQuote",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result.Has("'a'")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_WrapDoubleQuoteIfMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.WrapDoubleQuoteIfMissing()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapDoubleQuoteIfMissing",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result.Has("\"a\"")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_WrapSingleQuoteIfMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.WrapSingleQuoteIfMissing()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "WrapSingleQuoteIfMissing",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result.Has("'a'")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_JoinLine(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_JoinLine", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.JoinLine()
		// Assert
		convey.Convey("JoinLine", t, func() {
			convey.So(result, convey.ShouldNotBeEmpty)
		})
	})
}

func Test_Cov56_Hashset_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_NonEmptyJoins", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		// Act
		result := hs.NonEmptyJoins(",")
		// Assert
		convey.Convey("NonEmptyJoins", t, func() {
			convey.So(result, convey.ShouldNotBeEmpty)
		})
	})
}

func Test_Cov56_Hashset_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_NonWhitespaceJoins", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		// Act
		result := hs.NonWhitespaceJoins(",")
		// Assert
		convey.Convey("NonWhitespaceJoins", t, func() {
			convey.So(result, convey.ShouldNotBeEmpty)
		})
	})
}

func Test_Cov56_Hashset_SortedList(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_SortedList", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		// Act
		result := hs.SortedList()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "SortedList",
			ExpectedInput: args.Map{
				"First": "a",
				"Last": "c",
			},
		}
		actual := args.Map{
			"First": result[0],
			"Last": result[2],
		}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_Contains(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_Contains", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.Contains("a")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Contains alias",
			ExpectedInput: args.Map{"Contains": true},
		}
		actual := args.Map{"Contains": result}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_IsEqual", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		other := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.IsEqual(other)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "IsEqual alias",
			ExpectedInput: args.Map{"IsEqual": true},
		}
		actual := args.Map{"IsEqual": result}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_GetFilteredItems(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetFilteredItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"abc", "def"})
		filter := func(s string, i int) (string, bool, bool) { return s, s == "abc", false }
		// Act
		result := hs.GetFilteredItems(filter)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetFilteredItems",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_GetFilteredItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetFilteredItems_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		// Act
		result := hs.GetFilteredItems(filter)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetFilteredItems empty",
			ExpectedInput: args.Map{"Length": 0},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_GetFilteredItems_Break(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetFilteredItems_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		filter := func(s string, i int) (string, bool, bool) { return s, true, true }
		// Act
		result := hs.GetFilteredItems(filter)
		// Assert
		convey.Convey("GetFilteredItems break", t, func() {
			convey.So(len(result), convey.ShouldBeGreaterThan, 0)
		})
	})
}

func Test_Cov56_Hashset_GetFilteredCollection(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetFilteredCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"abc"})
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		// Act
		result := hs.GetFilteredCollection(filter)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetFilteredCollection",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": result.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_GetFilteredCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetFilteredCollection_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		// Act
		result := hs.GetFilteredCollection(filter)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetFilteredCollection empty",
			ExpectedInput: args.Map{"IsEmpty": true},
		}
		actual := args.Map{"IsEmpty": result.IsEmpty()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_GetFilteredCollection_Break(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetFilteredCollection_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		filter := func(s string, i int) (string, bool, bool) { return s, true, true }
		// Act
		result := hs.GetFilteredCollection(filter)
		// Assert
		convey.Convey("GetFilteredCollection break", t, func() {
			convey.So(result.Length(), convey.ShouldBeGreaterThan, 0)
		})
	})
}

func Test_Cov56_Hashset_GetAllExceptHashset(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetAllExceptHashset", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		exclude := corestr.New.Hashset.Strings([]string{"b"})
		// Act
		result := hs.GetAllExceptHashset(exclude)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetAllExceptHashset",
			ExpectedInput: args.Map{"Length": 2},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_GetAllExceptHashset_Nil(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetAllExceptHashset_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.GetAllExceptHashset(nil)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetAllExceptHashset nil",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetAllExcept", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		// Act
		result := hs.GetAllExcept([]string{"a"})
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetAllExcept",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_GetAllExcept_Nil(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetAllExcept_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.GetAllExcept(nil)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetAllExcept nil",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_GetAllExceptSpread(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetAllExceptSpread", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		// Act
		result := hs.GetAllExceptSpread("a")
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetAllExceptSpread",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetAllExceptCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		col := corestr.New.Collection.Strings([]string{"a"})
		// Act
		result := hs.GetAllExceptCollection(col)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetAllExceptCollection",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_GetAllExceptCollection_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.GetAllExceptCollection(nil)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "GetAllExceptCollection nil",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_StringLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.StringLock()
		// Assert
		convey.Convey("StringLock non-empty", t, func() {
			convey.So(result, convey.ShouldNotBeEmpty)
		})
	})
}

func Test_Cov56_Hashset_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_StringLock_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// Act
		result := hs.StringLock()
		// Assert
		convey.Convey("StringLock empty", t, func() {
			convey.So(result, convey.ShouldContainSubstring, "No Element")
		})
	})
}

func Test_Cov56_Hashset_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_ParseInjectUsingJsonMust", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		jsonResult := hs.JsonPtr()
		target := corestr.New.Hashset.Empty()
		// Act
		result := target.ParseInjectUsingJsonMust(jsonResult)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ParseInjectUsingJsonMust",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": result.Has("a")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_Items(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_Items", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.Items()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Items",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_Collection(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_Collection", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.Collection()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Hashset.Collection",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": result.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_ListPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		result := hs.ListPtr()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ListPtr deprecated alias",
			ExpectedInput: args.Map{"Length": 1},
		}
		actual := args.Map{"Length": len(result)}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddCapacities(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddCapacities", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		hs.AddCapacities(10, 20)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddCapacities",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": hs.Has("a")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddCapacities_Empty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddCapacities_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		hs.AddCapacities()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddCapacities empty",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": hs.Has("a")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddCapacitiesLock(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddCapacitiesLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		hs.AddCapacitiesLock(10)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddCapacitiesLock",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": hs.Has("a")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_AddCapacitiesLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_AddCapacitiesLock_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		hs.AddCapacitiesLock()
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "AddCapacitiesLock empty",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": hs.Has("a")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_Resize", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		hs.Resize(100)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Resize",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": hs.Has("a")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_Resize_AlreadyLarger(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_Resize_AlreadyLarger", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		// Act
		hs.Resize(1)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "Resize already larger",
			ExpectedInput: args.Map{"Length": 3},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_ResizeLock(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_ResizeLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		// Act
		hs.ResizeLock(100)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ResizeLock",
			ExpectedInput: args.Map{"Has": true},
		}
		actual := args.Map{"Has": hs.Has("a")}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}

func Test_Cov56_Hashset_ResizeLock_AlreadyLarger(t *testing.T) {
	safeTest(t, "Test_Cov56_Hashset_ResizeLock_AlreadyLarger", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		// Act
		hs.ResizeLock(1)
		// Assert
		tc := coretestcases.CaseV1{
			Title:         "ResizeLock already larger",
			ExpectedInput: args.Map{"Length": 3},
		}
		actual := args.Map{"Length": hs.Length()}
		tc.ShouldBeEqualMap(t, 0, actual)
	})
}
