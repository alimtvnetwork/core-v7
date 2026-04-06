package coreappendtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coreappend"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// PrependAppendAnyItemsToStringsSkipOnNil
// ==========================================

func Test_PrependAppendToStrings_AllNonNil(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		"PRE", "POST", "a", "b",
	)
	actual := args.Map{"result": len(result) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": result[0] != "PRE" || result[len(result)-1] != "POST"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected PRE...POST", actual)
}

func Test_PrependAppendToStrings_NilPrepend(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		nil, "POST", "a",
	)
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 (skip nil prepend)", actual)
	actual := args.Map{"result": result[0] != "a" || result[1] != "POST"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected result:", actual)
}

func Test_PrependAppendToStrings_NilAppend(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		"PRE", nil, "a",
	)
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 (skip nil append)", actual)
	actual := args.Map{"result": result[0] != "PRE" || result[1] != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected result:", actual)
}

func Test_PrependAppendToStrings_BothNil(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		nil, nil, "a",
	)
	actual := args.Map{"result": len(result) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_PrependAppendToStrings_NilInMiddle(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		"PRE", "POST", "a", nil, "b",
	)
	actual := args.Map{"result": len(result) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 (skip nil middle)", actual)
}

func Test_PrependAppendToStrings_NoItems(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		"PRE", "POST",
	)
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 (just pre+post)", actual)
}

func Test_PrependAppendToStrings_AllNil(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		nil, nil,
	)
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// AppendAnyItemsToStringSkipOnNil
// ==========================================

func Test_AppendToString_Basic(t *testing.T) {
	result := coreappend.AppendAnyItemsToStringSkipOnNil(
		",", "SUFFIX", "a", "b",
	)
	actual := args.Map{"result": result != "a,b,SUFFIX"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'a,b,SUFFIX', got ''", actual)
}

func Test_AppendToString_NilAppend(t *testing.T) {
	result := coreappend.AppendAnyItemsToStringSkipOnNil(
		",", nil, "a",
	)
	actual := args.Map{"result": result != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'a', got ''", actual)
}

// ==========================================
// PrependAnyItemsToStringSkipOnNil
// ==========================================

func Test_PrependToString_Basic(t *testing.T) {
	result := coreappend.PrependAnyItemsToStringSkipOnNil(
		",", "PREFIX", "a", "b",
	)
	actual := args.Map{"result": result != "PREFIX,a,b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'PREFIX,a,b', got ''", actual)
}

func Test_PrependToString_NilPrepend(t *testing.T) {
	result := coreappend.PrependAnyItemsToStringSkipOnNil(
		",", nil, "a",
	)
	actual := args.Map{"result": result != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'a', got ''", actual)
}

// ==========================================
// PrependAppendAnyItemsToStringSkipOnNil (joined)
// ==========================================

func Test_PrependAppendToString_Joined(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringSkipOnNil(
		"-", "PRE", "POST", "mid",
	)
	actual := args.Map{"result": result != "PRE-mid-POST"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'PRE-mid-POST', got ''", actual)
}

// ==========================================
// PrependAppendAnyItemsToStringsUsingFunc
// ==========================================

func Test_PrependAppendUsingFunc_Basic(t *testing.T) {
	compiler := func(item any) string {
		return fmt.Sprintf("[%v]", item)
	}
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		false, compiler, "pre", "post", "a", "b",
	)
	actual := args.Map{"result": len(result) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": result[0] != "[pre]" || result[3] != "[post]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_PrependAppendUsingFunc_SkipEmpty(t *testing.T) {
	compiler := func(item any) string {
		if item == nil {
			return ""
		}
		return fmt.Sprintf("%v", item)
	}
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		true, compiler, nil, nil, "a", nil, "b",
	)
	// prepend=nil→"" skipped, append=nil→"" skipped, nil middle skipped
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 (skip empties), got:", actual)
}

func Test_PrependAppendUsingFunc_NoSkipEmpty(t *testing.T) {
	compiler := func(item any) string {
		if item == nil {
			return ""
		}
		return fmt.Sprintf("%v", item)
	}
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		false, compiler, nil, nil, "a",
	)
	// prepend="" included, append="" included, nil middle skipped
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 (include empties), got:", actual)
}

// ==========================================
// MapStringStringAppendMapStringToAnyItems
// ==========================================

func Test_MapAppend_Basic(t *testing.T) {
	mainMap := map[string]string{"a": "1"}
	appendMap := map[string]any{"b": 2, "c": "three"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(false, mainMap, appendMap)
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_MapAppend_EmptyAppend(t *testing.T) {
	mainMap := map[string]string{"a": "1"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(false, mainMap, map[string]any{})
	actual := args.Map{"result": len(result) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_MapAppend_SkipEmpty(t *testing.T) {
	mainMap := map[string]string{}
	appendMap := map[string]any{"a": "", "b": "val"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(true, mainMap, appendMap)
	// "a" has value "" which after Sprintf becomes "" → skipped
	actual := args.Map{"result": _, has := result["a"]; has}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SkipEmpty should skip empty string values", actual)
	actual := args.Map{"result": result["b"] != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'val', got ''", actual)
}

func Test_MapAppend_OverwriteExisting(t *testing.T) {
	mainMap := map[string]string{"k": "old"}
	appendMap := map[string]any{"k": "new"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(false, mainMap, appendMap)
	actual := args.Map{"result": result["k"] != "new"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected overwrite to 'new', got ''", actual)
}
