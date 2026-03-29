package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Basic ──

func Test_Src02_Collection_Basic(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_Basic", func() {
		// Arrange
		tc := collectionBasicTestCases[0]

		// Act
		c := corestr.New.Collection.Empty()

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"isEmpty":   c.IsEmpty(),
			"hasItems":  c.HasItems(),
			"len":       c.Length(),
			"count":     c.Count(),
			"lastIndex": c.LastIndex(),
			"cap":       c.Capacity(),
			"hasIndex0": c.HasIndex(0),
		})
	})
}

func Test_Src02_Collection_NilReceiver(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_NilReceiver", func() {
		// Arrange
		tc := collectionNilReceiverTestCases[0]
		var c *corestr.Collection

		// Act
		length := c.Length()
		isEmpty := c.IsEmpty()

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len":     length,
			"isEmpty": isEmpty,
		})
	})
}

// ── Add methods ──

func Test_Src02_Collection_Add(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_Add", func() {
		// Arrange
		tc := collectionAddTestCases[0]

		// Act
		c := corestr.New.Collection.Cap(5)
		c.Add("a").Add("b")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len":    c.Length(),
			"capGe2": c.Capacity() >= 2,
		})
	})
}

func Test_Src02_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddNonEmpty", func() {
		// Arrange
		tc := collectionAddTestCases[1]

		// Act
		c := corestr.New.Collection.Empty()
		c.AddNonEmpty("")
		c.AddNonEmpty("a")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

func Test_Src02_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddNonEmptyWhitespace", func() {
		// Arrange
		tc := collectionAddTestCases[2]

		// Act
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyWhitespace("   ")
		c.AddNonEmptyWhitespace("a")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

func Test_Src02_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddError", func() {
		// Arrange
		tc := collectionAddTestCases[3]

		// Act
		c := corestr.New.Collection.Empty()
		c.AddError(nil)
		c.AddError(errors.New("e"))

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

func Test_Src02_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddIf", func() {
		// Arrange
		tc := collectionAddTestCases[4]

		// Act
		c := corestr.New.Collection.Empty()
		c.AddIf(false, "skip")
		c.AddIf(true, "add")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

func Test_Src02_Collection_AddIfMany(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddIfMany", func() {
		// Arrange
		tc := collectionAddTestCases[5]

		// Act
		c := corestr.New.Collection.Empty()
		c.AddIfMany(false, "a", "b")
		c.AddIfMany(true, "c", "d")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

func Test_Src02_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_Adds", func() {
		// Arrange
		tc := collectionAddTestCases[6]

		// Act
		c := corestr.New.Collection.Empty()
		c.Adds("a", "b", "c")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

func Test_Src02_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddStrings", func() {
		// Arrange
		tc := collectionAddTestCases[7]

		// Act
		c := corestr.New.Collection.Empty()
		c.AddStrings([]string{"x", "y"})

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

func Test_Src02_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddFunc", func() {
		// Arrange
		tc := collectionAddTestCases[8]

		// Act
		c := corestr.New.Collection.Empty()
		c.AddFunc(func() string { return "hello" })

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

func Test_Src02_Collection_AddFuncErr_NoErr(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddFuncErr_NoErr", func() {
		// Arrange
		tc := collectionAddTestCases[9]

		// Act
		c := corestr.New.Collection.Empty()
		c.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { t.Fatal("should not be called") },
		)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

func Test_Src02_Collection_AddFuncErr_WithErr(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddFuncErr_WithErr", func() {
		// Arrange
		tc := collectionAddTestCases[10]
		called := false

		// Act
		c := corestr.New.Collection.Empty()
		c.AddFuncErr(
			func() (string, error) { return "", errors.New("e") },
			func(err error) { called = true },
		)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len":    c.Length(),
			"called": called,
		})
	})
}

func Test_Src02_Collection_AddLock(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddLock", func() {
		// Arrange
		tc := collectionAddTestCases[11]

		// Act
		c := corestr.New.Collection.Empty()
		c.AddLock("a")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

func Test_Src02_Collection_AddsLock(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddsLock", func() {
		// Arrange
		tc := collectionAddTestCases[12]

		// Act
		c := corestr.New.Collection.Empty()
		c.AddsLock("a", "b")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

// ── Merge ──

func Test_Src02_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddCollection", func() {
		// Arrange
		tc := collectionMergeTestCases[0]

		// Act
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c1.AddCollection(c2)
		lenAfterMerge := c1.Length()
		c1.AddCollection(corestr.New.Collection.Empty())

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len":         lenAfterMerge,
			"lenAfterAdd": c1.Length(),
		})
	})
}

func Test_Src02_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddCollections", func() {
		// Arrange
		tc := collectionMergeTestCases[1]

		// Act
		c := corestr.New.Collection.Empty()
		c.AddCollections(
			corestr.New.Collection.Strings([]string{"a"}),
			corestr.New.Collection.Empty(),
			corestr.New.Collection.Strings([]string{"b"}),
		)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

// ── Remove ──

func Test_Src02_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_RemoveAt", func() {
		// Arrange
		tc := collectionRemoveTestCases[0]

		// Act
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		success := c.RemoveAt(1)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"success":  success,
			"len":      c.Length(),
			"failNeg":  c.RemoveAt(-1),
			"failHigh": c.RemoveAt(100),
		})
	})
}

// ── Query ──

func Test_Src02_Collection_ListStrings(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_ListStrings", func() {
		// Arrange
		tc := collectionQueryTestCases[0]

		// Act
		c := corestr.New.Collection.Strings([]string{"a"})

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"lenList":    len(c.ListStrings()),
			"lenListPtr": len(c.ListStringsPtr()),
		})
	})
}

func Test_Src02_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_LengthLock", func() {
		// Arrange
		tc := collectionQueryTestCases[1]

		// Act
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.LengthLock(),
		})
	})
}

func Test_Src02_Collection_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_IsEmptyLock", func() {
		// Arrange
		tc := collectionQueryTestCases[2]

		// Act
		c := corestr.New.Collection.Empty()

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"isEmpty": c.IsEmptyLock(),
		})
	})
}

func Test_Src02_Collection_HasIndex(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_HasIndex", func() {
		// Arrange
		tc := collectionQueryTestCases[3]

		// Act
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"has0":    c.HasIndex(0),
			"has1":    c.HasIndex(1),
			"has2":    c.HasIndex(2),
			"hasNeg1": c.HasIndex(-1),
		})
	})
}

// ── Error conversion ──

func Test_Src02_Collection_AsError(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AsError", func() {
		// Arrange
		tc := collectionErrorTestCases[0]

		// Act
		c := corestr.New.Collection.Empty()
		defaultNilEmpty := c.AsDefaultError() == nil
		asErrorNilEmpty := c.AsError(",") == nil
		c.Add("e1")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"defaultNilEmpty": defaultNilEmpty,
			"asErrorNilEmpty": asErrorNilEmpty,
			"defaultNonNil":   c.AsDefaultError() != nil,
		})
	})
}

func Test_Src02_Collection_ToError(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_ToError", func() {
		// Arrange
		tc := collectionErrorTestCases[1]

		// Act
		c := corestr.New.Collection.Empty()
		toErrorNilEmpty := c.ToError(",") == nil
		toDefaultNilEmpty := c.ToDefaultError() == nil
		c.Add("e")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"toErrorNilEmpty":   toErrorNilEmpty,
			"toDefaultNilEmpty": toDefaultNilEmpty,
			"toErrorNonNil":     c.ToError(",") != nil,
		})
	})
}

// ── Misc ──

func Test_Src02_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_EachItemSplitBy", func() {
		// Arrange
		tc := collectionMiscTestCases[0]

		// Act
		c := corestr.New.Collection.Strings([]string{"a,b", "c"})
		result := c.EachItemSplitBy(",")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": len(result),
		})
	})
}

func Test_Src02_Collection_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_ConcatNew_Empty", func() {
		// Arrange
		tc := collectionMiscTestCases[1]

		// Act
		c := corestr.New.Collection.Strings([]string{"a"})
		newC := c.ConcatNew(0)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": newC.Length(),
		})
	})
}

func Test_Src02_Collection_ConcatNew_WithItems(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_ConcatNew_WithItems", func() {
		// Arrange
		tc := collectionMiscTestCases[2]

		// Act
		c := corestr.New.Collection.Strings([]string{"a"})
		newC := c.ConcatNew(0, "b", "c")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": newC.Length(),
		})
	})
}

func Test_Src02_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_IsEquals", func() {
		// Arrange
		tc := collectionMiscTestCases[3]

		// Act
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"equal": a.IsEquals(b),
		})
	})
}

func Test_Src02_Collection_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_IsEqualsWithSensitive", func() {
		// Arrange
		tc := collectionMiscTestCases[4]

		// Act
		a := corestr.New.Collection.Strings([]string{"A"})
		b := corestr.New.Collection.Strings([]string{"a"})

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"insensitiveEqual": a.IsEqualsWithSensitive(false, b),
			"sensitiveEqual":   a.IsEqualsWithSensitive(true, b),
		})
	})
}

func Test_Src02_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_JsonString", func() {
		// Arrange
		tc := collectionMiscTestCases[5]

		// Act
		c := corestr.New.Collection.Strings([]string{"a"})

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"jsonNonEmpty":     c.JsonString() != "",
			"jsonMustNonEmpty": c.JsonStringMust() != "",
			"stringJSON":       c.StringJSON() != "",
		})
	})
}

// ── Hashmap integration ──

func Test_Src02_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddHashmapsValues", func() {
		// Arrange
		tc := collectionHashmapTestCases[0]

		// Act
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsValues(hm)
		c.AddHashmapsValues(nil)
		c.AddHashmapsValues(corestr.New.Hashmap.Empty())

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

func Test_Src02_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddHashmapsKeys", func() {
		// Arrange
		tc := collectionHashmapTestCases[1]

		// Act
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeys(hm)
		c.AddHashmapsKeys(nil)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}

func Test_Src02_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_Src02_Collection_AddPointerCollectionsLock", func() {
		// Arrange
		tc := collectionHashmapTestCases[2]

		// Act
		c := corestr.New.Collection.Empty()
		c2 := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(c2)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": c.Length(),
		})
	})
}
