package coreappendtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreappend"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── AppendAnyItemsToStringSkipOnNil ──

func Test_AppendAnyItems_Basic(t *testing.T) {
	// Arrange
	result := coreappend.AppendAnyItemsToStringSkipOnNil(
		", ", "suffix", "a", "b",
	)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a, b, suffix"}
	expected.ShouldBeEqual(t, 0, "AppendAnyItems basic -- joined", actual)
}

func Test_AppendAnyItems_NilItems(t *testing.T) {
	// Arrange
	result := coreappend.AppendAnyItemsToStringSkipOnNil(
		", ", "end", nil, "a", nil,
	)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a, end"}
	expected.ShouldBeEqual(t, 0, "AppendAnyItems nil items -- skipped", actual)
}

func Test_AppendAnyItems_NilAppend(t *testing.T) {
	// Arrange
	result := coreappend.AppendAnyItemsToStringSkipOnNil(
		", ", nil, "a", "b",
	)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a, b"}
	expected.ShouldBeEqual(t, 0, "AppendAnyItems nil append -- no suffix", actual)
}

// ── PrependAnyItemsToStringSkipOnNil ──

func Test_PrependAnyItems_Basic(t *testing.T) {
	// Arrange
	result := coreappend.PrependAnyItemsToStringSkipOnNil(
		", ", "prefix", "a", "b",
	)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "prefix, a, b"}
	expected.ShouldBeEqual(t, 0, "PrependAnyItems basic -- joined", actual)
}

func Test_PrependAnyItems_NilPrepend(t *testing.T) {
	// Arrange
	result := coreappend.PrependAnyItemsToStringSkipOnNil(
		", ", nil, "a", "b",
	)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a, b"}
	expected.ShouldBeEqual(t, 0, "PrependAnyItems nil prepend -- no prefix", actual)
}

// ── PrependAppendAnyItemsToStringSkipOnNil ──

func Test_PrependAppendAnyItems_Both(t *testing.T) {
	// Arrange
	result := coreappend.PrependAppendAnyItemsToStringSkipOnNil(
		", ", "pre", "post", "a",
	)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "pre, a, post"}
	expected.ShouldBeEqual(t, 0, "PrependAppend both -- joined", actual)
}

func Test_PrependAppendAnyItems_BothNil(t *testing.T) {
	// Arrange
	result := coreappend.PrependAppendAnyItemsToStringSkipOnNil(
		", ", nil, nil, "a",
	)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "PrependAppend both nil -- items only", actual)
}

// ── PrependAppendAnyItemsToStringsUsingFunc ──

func Test_PrependAppendUsingFunc_Basic(t *testing.T) {
	// Arrange
	fn := func(item any) string {
		if item == nil {
			return ""
		}
		return item.(string)
	}
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		true, fn, "pre", "post", "a", nil, "b",
	)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 4} // pre, a, b, post
	expected.ShouldBeEqual(t, 0, "UsingFunc basic -- skips empty", actual)
}

func Test_PrependAppendUsingFunc_NoSkipEmpty(t *testing.T) {
	// Arrange
	fn := func(item any) string {
		if item == nil {
			return ""
		}
		return item.(string)
	}
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		false, fn, "pre", "post", "a", nil,
	)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3} // pre, a, post (nil item skipped at line 20)
	expected.ShouldBeEqual(t, 0, "UsingFunc no skip -- includes empty", actual)
}

// ── MapStringStringAppendMapStringToAnyItems ──

func Test_MapStringStringAppend_Basic(t *testing.T) {
	// Arrange
	mainMap := map[string]string{"a": "1"}
	appendMap := map[string]any{"b": 2, "c": "three"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(
		false, mainMap, appendMap,
	)

	// Act
	actual := args.Map{
		"hasA": result["a"] == "1",
		"hasB": result["b"] != "",
		"hasC": result["c"] == "three",
	}

	// Assert
	expected := args.Map{
		"hasA": true,
		"hasB": true,
		"hasC": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAppend basic -- merged", actual)
}

func Test_MapStringStringAppend_EmptyAppend(t *testing.T) {
	// Arrange
	mainMap := map[string]string{"a": "1"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(
		false, mainMap, nil,
	)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAppend empty append -- unchanged", actual)
}

func Test_MapStringStringAppend_SkipEmpty(t *testing.T) {
	// Arrange
	mainMap := map[string]string{}
	appendMap := map[string]any{"a": "", "b": "val"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(
		true, mainMap, appendMap,
	)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1} // "a" skipped because empty
	expected.ShouldBeEqual(t, 0, "MapAppend skip empty -- only b", actual)
}

// ── PrependAppendAnyItemsToStringsSkipOnNil (direct slice) ──

func Test_PrependAppendStrings_Empty(t *testing.T) {
	// Arrange
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		nil, nil,
	)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PrependAppendStrings empty -- zero items", actual)
}
