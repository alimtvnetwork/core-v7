package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice — Segment 4a: Add, Insert, Accessors, Contains, Index, Length
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg4_SS_Add(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Add", func() {
		s := corestr.SimpleSlice{}
		s.Add("a").Add("b")
		actual := args.Map{"len": s.Length(), "first": s.First()}
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "Add -- 2 items", actual)
	})
}

func Test_Seg4_SS_AddSplit(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddSplit", func() {
		s := corestr.SimpleSlice{}
		s.AddSplit("a,b,c", ",")
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AddSplit -- 3 items from CSV", actual)
	})
}

func Test_Seg4_SS_AddIf(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddIf", func() {
		s := corestr.SimpleSlice{}
		s.AddIf(true, "yes").AddIf(false, "no")
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddIf -- only true added", actual)
	})
}

func Test_Seg4_SS_Adds(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Adds", func() {
		s := corestr.SimpleSlice{}
		s.Adds("a", "b", "c")
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Adds -- 3 items", actual)
	})
}

func Test_Seg4_SS_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Adds_Empty", func() {
		s := corestr.SimpleSlice{}
		s.Adds()
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds empty -- no change", actual)
	})
}

func Test_Seg4_SS_Append(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Append", func() {
		s := corestr.SimpleSlice{}
		s.Append("a", "b")
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Append -- 2 items", actual)
	})
}

func Test_Seg4_SS_Append_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Append_Empty", func() {
		s := corestr.SimpleSlice{}
		s.Append()
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Append empty -- no change", actual)
	})
}

func Test_Seg4_SS_AppendFmt(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AppendFmt", func() {
		s := corestr.SimpleSlice{}
		s.AppendFmt("hello %s", "world")
		actual := args.Map{"val": s.First()}
		expected := args.Map{"val": "hello world"}
		expected.ShouldBeEqual(t, 0, "AppendFmt -- formatted", actual)
	})
}

func Test_Seg4_SS_AppendFmt_EmptySkip(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AppendFmt_EmptySkip", func() {
		s := corestr.SimpleSlice{}
		s.AppendFmt("")
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendFmt empty -- skipped", actual)
	})
}

func Test_Seg4_SS_AppendFmtIf(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AppendFmtIf", func() {
		s := corestr.SimpleSlice{}
		s.AppendFmtIf(true, "v=%d", 42)
		s.AppendFmtIf(false, "skip=%d", 99)
		actual := args.Map{"len": s.Length(), "val": s.First()}
		expected := args.Map{"len": 1, "val": "v=42"}
		expected.ShouldBeEqual(t, 0, "AppendFmtIf -- only true", actual)
	})
}

func Test_Seg4_SS_AppendFmtIf_EmptySkip(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AppendFmtIf_EmptySkip", func() {
		s := corestr.SimpleSlice{}
		s.AppendFmtIf(true, "")
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendFmtIf empty format -- skipped", actual)
	})
}

func Test_Seg4_SS_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddAsTitleValue", func() {
		s := corestr.SimpleSlice{}
		s.AddAsTitleValue("Name", "Alice")
		actual := args.Map{"nonEmpty": s.First() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AddAsTitleValue -- formatted", actual)
	})
}

func Test_Seg4_SS_AddAsTitleValueIf(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddAsTitleValueIf", func() {
		s := corestr.SimpleSlice{}
		s.AddAsTitleValueIf(true, "Name", "Alice")
		s.AddAsTitleValueIf(false, "Skip", "Bob")
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddAsTitleValueIf -- only true", actual)
	})
}

func Test_Seg4_SS_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddAsCurlyTitleWrap", func() {
		s := corestr.SimpleSlice{}
		s.AddAsCurlyTitleWrap("Key", "Val")
		actual := args.Map{"nonEmpty": s.First() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AddAsCurlyTitleWrap -- formatted", actual)
	})
}

func Test_Seg4_SS_AddAsCurlyTitleWrapIf(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddAsCurlyTitleWrapIf", func() {
		s := corestr.SimpleSlice{}
		s.AddAsCurlyTitleWrapIf(true, "K", "V")
		s.AddAsCurlyTitleWrapIf(false, "S", "X")
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddAsCurlyTitleWrapIf -- only true", actual)
	})
}

func Test_Seg4_SS_InsertAt(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_InsertAt", func() {
		s := corestr.SimpleSlice{"a", "c"}
		s.InsertAt(1, "b")
		actual := args.Map{"len": s.Length(), "mid": s[1]}
		expected := args.Map{"len": 3, "mid": "b"}
		expected.ShouldBeEqual(t, 0, "InsertAt -- inserted in middle", actual)
	})
}

func Test_Seg4_SS_InsertAt_OutOfBounds(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_InsertAt_OutOfBounds", func() {
		s := corestr.SimpleSlice{"a"}
		s.InsertAt(-1, "x")
		s.InsertAt(99, "y")
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "InsertAt out of bounds -- no change", actual)
	})
}

func Test_Seg4_SS_AddStruct(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddStruct", func() {
		s := corestr.SimpleSlice{}
		s.AddStruct(false, struct{ Name string }{"Alice"})
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStruct -- 1 item", actual)
	})
}

func Test_Seg4_SS_AddStruct_Nil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddStruct_Nil", func() {
		s := corestr.SimpleSlice{}
		s.AddStruct(false, nil)
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStruct nil -- no change", actual)
	})
}

func Test_Seg4_SS_AddPointer_Nil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddPointer_Nil", func() {
		s := corestr.SimpleSlice{}
		s.AddPointer(false, nil)
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddPointer nil -- no change", actual)
	})
}

func Test_Seg4_SS_AddsIf(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddsIf", func() {
		s := corestr.SimpleSlice{}
		s.AddsIf(true, "a", "b").AddsIf(false, "c")
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsIf -- only true batch", actual)
	})
}

func Test_Seg4_SS_AddError(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddError", func() {
		s := corestr.SimpleSlice{}
		s.AddError(errors.New("err")).AddError(nil)
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddError -- only non-nil", actual)
	})
}

// ── Accessors ───────────────────────────────────────────────────────────────

func Test_Seg4_SS_First_Last(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_First_Last", func() {
		s := corestr.SimpleSlice{"a", "b", "c"}
		actual := args.Map{"first": s.First(), "last": s.Last(), "firstDyn": s.FirstDynamic(), "lastDyn": s.LastDynamic()}
		expected := args.Map{"first": "a", "last": "c", "firstDyn": "a", "lastDyn": "c"}
		expected.ShouldBeEqual(t, 0, "First/Last/Dynamic -- correct", actual)
	})
}

func Test_Seg4_SS_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_FirstOrDefault_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"val": s.FirstOrDefault(), "dyn": s.FirstOrDefaultDynamic()}
		expected := args.Map{"val": "", "dyn": ""}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault empty -- empty string", actual)
	})
}

func Test_Seg4_SS_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_LastOrDefault_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"val": s.LastOrDefault(), "dyn": s.LastOrDefaultDynamic()}
		expected := args.Map{"val": "", "dyn": ""}
		expected.ShouldBeEqual(t, 0, "LastOrDefault empty -- empty string", actual)
	})
}

func Test_Seg4_SS_AsError(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AsError", func() {
		s := corestr.SimpleSlice{"e1", "e2"}
		err := s.AsError("; ")
		actual := args.Map{"notNil": err != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AsError -- non-nil", actual)
	})
}

func Test_Seg4_SS_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AsError_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"nil": s.AsError("; ") == nil, "defNil": s.AsDefaultError() == nil}
		expected := args.Map{"nil": true, "defNil": true}
		expected.ShouldBeEqual(t, 0, "AsError empty -- nil", actual)
	})
}

func Test_Seg4_SS_AsError_Nil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AsError_Nil", func() {
		var s *corestr.SimpleSlice
		actual := args.Map{"nil": s.AsError("; ") == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "AsError nil -- nil", actual)
	})
}

// ── Skip / Take / Limit ────────────────────────────────────────────────────

func Test_Seg4_SS_Skip(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Skip", func() {
		s := corestr.SimpleSlice{"a", "b", "c"}
		actual := args.Map{"len": len(s.Skip(1)), "dynLen": len(s.SkipDynamic(1).([]string))}
		expected := args.Map{"len": 2, "dynLen": 2}
		expected.ShouldBeEqual(t, 0, "Skip -- 2 remaining", actual)
	})
}

func Test_Seg4_SS_Skip_MoreThanLen(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Skip_MoreThanLen", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{"len": len(s.Skip(10))}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Skip more than len -- empty", actual)
	})
}

func Test_Seg4_SS_Take(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Take", func() {
		s := corestr.SimpleSlice{"a", "b", "c"}
		actual := args.Map{"len": len(s.Take(2))}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Take 2 -- 2 items", actual)
	})
}

func Test_Seg4_SS_Take_MoreThanLen(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Take_MoreThanLen", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{"len": len(s.Take(10))}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Take more than len -- all items", actual)
	})
}

func Test_Seg4_SS_Limit(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Limit", func() {
		s := corestr.SimpleSlice{"a", "b", "c"}
		actual := args.Map{"len": len(s.Limit(2))}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Limit -- delegates to Take", actual)
	})
}

func Test_Seg4_SS_LimitDynamic(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_LimitDynamic", func() {
		s := corestr.SimpleSlice{"a", "b", "c"}
		actual := args.Map{"len": len(s.LimitDynamic(2).([]string))}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LimitDynamic -- delegates to Take", actual)
	})
}

func Test_Seg4_SS_TakeDynamic(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_TakeDynamic", func() {
		s := corestr.SimpleSlice{"a", "b"}
		// Fix: TakeDynamic returns []string, not SimpleSlice
		// See issues/corestrtests-takedynamic-type-assertion.md
		actual := args.Map{"len": len(s.TakeDynamic(10).([]string))}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "TakeDynamic more -- all items", actual)
	})
}

// ── Length / Count / CountFunc / IsEmpty / HasAnyItem ────────────────────────

func Test_Seg4_SS_Length_Nil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Length_Nil", func() {
		var s *corestr.SimpleSlice
		actual := args.Map{"len": s.Length(), "empty": s.IsEmpty()}
		expected := args.Map{"len": 0, "empty": true}
		expected.ShouldBeEqual(t, 0, "Length nil -- 0", actual)
	})
}

func Test_Seg4_SS_Count(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Count", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"count": s.Count()}
		expected := args.Map{"count": 2}
		expected.ShouldBeEqual(t, 0, "Count -- 2", actual)
	})
}

func Test_Seg4_SS_CountFunc(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_CountFunc", func() {
		s := corestr.SimpleSlice{"a", "bb", "ccc"}
		actual := args.Map{"count": s.CountFunc(func(i int, item string) bool { return len(item) > 1 })}
		expected := args.Map{"count": 2}
		expected.ShouldBeEqual(t, 0, "CountFunc -- 2 match", actual)
	})
}

func Test_Seg4_SS_CountFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_CountFunc_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"count": s.CountFunc(func(i int, item string) bool { return true })}
		expected := args.Map{"count": 0}
		expected.ShouldBeEqual(t, 0, "CountFunc empty -- 0", actual)
	})
}

func Test_Seg4_SS_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_HasAnyItem", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{"has": s.HasAnyItem()}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasAnyItem -- true", actual)
	})
}

func Test_Seg4_SS_LastIndex(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_LastIndex", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"idx": s.LastIndex()}
		expected := args.Map{"idx": 1}
		expected.ShouldBeEqual(t, 0, "LastIndex -- 1", actual)
	})
}

func Test_Seg4_SS_HasIndex(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_HasIndex", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"has0": s.HasIndex(0), "has5": s.HasIndex(5), "hasNeg": s.HasIndex(-1)}
		expected := args.Map{"has0": true, "has5": false, "hasNeg": false}
		expected.ShouldBeEqual(t, 0, "HasIndex -- valid and invalid", actual)
	})
}

// ── IsContains / IsContainsFunc / IndexOf / IndexOfFunc ─────────────────────

func Test_Seg4_SS_IsContains(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsContains", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"has": s.IsContains("a"), "miss": s.IsContains("z")}
		expected := args.Map{"has": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "IsContains -- found and missing", actual)
	})
}

func Test_Seg4_SS_IsContains_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsContains_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"has": s.IsContains("a")}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "IsContains empty -- false", actual)
	})
}

func Test_Seg4_SS_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsContainsFunc", func() {
		s := corestr.SimpleSlice{"hello", "world"}
		actual := args.Map{"has": s.IsContainsFunc("hello", func(item, search string) bool { return item == search })}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "IsContainsFunc -- found", actual)
	})
}

func Test_Seg4_SS_IsContainsFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsContainsFunc_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"has": s.IsContainsFunc("x", func(item, search string) bool { return true })}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "IsContainsFunc empty -- false", actual)
	})
}

func Test_Seg4_SS_IndexOf(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IndexOf", func() {
		s := corestr.SimpleSlice{"a", "b", "c"}
		actual := args.Map{"idx": s.IndexOf("b"), "miss": s.IndexOf("z")}
		expected := args.Map{"idx": 1, "miss": -1}
		expected.ShouldBeEqual(t, 0, "IndexOf -- found and missing", actual)
	})
}

func Test_Seg4_SS_IndexOf_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IndexOf_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"idx": s.IndexOf("a")}
		expected := args.Map{"idx": -1}
		expected.ShouldBeEqual(t, 0, "IndexOf empty -- -1", actual)
	})
}

func Test_Seg4_SS_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IndexOfFunc", func() {
		s := corestr.SimpleSlice{"aa", "bb"}
		actual := args.Map{"idx": s.IndexOfFunc("bb", func(item, search string) bool { return item == search })}
		expected := args.Map{"idx": 1}
		expected.ShouldBeEqual(t, 0, "IndexOfFunc -- found", actual)
	})
}

func Test_Seg4_SS_IndexOfFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IndexOfFunc_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"idx": s.IndexOfFunc("x", func(item, search string) bool { return true })}
		expected := args.Map{"idx": -1}
		expected.ShouldBeEqual(t, 0, "IndexOfFunc empty -- -1", actual)
	})
}

// ── Strings / List / Hashset ────────────────────────────────────────────────

func Test_Seg4_SS_StringsList(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_StringsList", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"strLen": len(s.Strings()), "listLen": len(s.List())}
		expected := args.Map{"strLen": 2, "listLen": 2}
		expected.ShouldBeEqual(t, 0, "Strings/List -- same", actual)
	})
}

func Test_Seg4_SS_Hashset(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Hashset", func() {
		s := corestr.SimpleSlice{"a", "b", "a"}
		actual := args.Map{"len": s.Hashset().Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset -- 2 unique", actual)
	})
}

// ── Wrap methods ────────────────────────────────────────────────────────────

func Test_Seg4_SS_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_WrapDoubleQuote", func() {
		s := corestr.SimpleSlice{"a", "b"}
		w := s.WrapDoubleQuote()
		actual := args.Map{"len": w.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "WrapDoubleQuote -- 2 items", actual)
	})
}

func Test_Seg4_SS_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_WrapSingleQuote", func() {
		s := corestr.SimpleSlice{"a"}
		w := s.WrapSingleQuote()
		actual := args.Map{"len": w.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapSingleQuote -- 1 item", actual)
	})
}

func Test_Seg4_SS_WrapTildaQuote(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_WrapTildaQuote", func() {
		s := corestr.SimpleSlice{"a"}
		w := s.WrapTildaQuote()
		actual := args.Map{"len": w.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapTildaQuote -- 1 item", actual)
	})
}

func Test_Seg4_SS_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_WrapDoubleQuoteIfMissing", func() {
		s := corestr.SimpleSlice{"a"}
		w := s.WrapDoubleQuoteIfMissing()
		actual := args.Map{"len": w.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapDoubleQuoteIfMissing -- 1 item", actual)
	})
}

func Test_Seg4_SS_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_WrapSingleQuoteIfMissing", func() {
		s := corestr.SimpleSlice{"a"}
		w := s.WrapSingleQuoteIfMissing()
		actual := args.Map{"len": w.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapSingleQuoteIfMissing -- 1 item", actual)
	})
}

// ── Transpile / TranspileJoin ───────────────────────────────────────────────

func Test_Seg4_SS_Transpile(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Transpile", func() {
		s := corestr.SimpleSlice{"a", "b"}
		result := s.Transpile(func(x string) string { return x + "!" })
		actual := args.Map{"first": (*result)[0]}
		expected := args.Map{"first": "a!"}
		expected.ShouldBeEqual(t, 0, "Transpile -- modified", actual)
	})
}

func Test_Seg4_SS_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Transpile_Empty", func() {
		s := corestr.SimpleSlice{}
		result := s.Transpile(func(x string) string { return x })
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Transpile empty -- empty", actual)
	})
}

func Test_Seg4_SS_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_TranspileJoin", func() {
		s := corestr.SimpleSlice{"a", "b"}
		result := s.TranspileJoin(func(x string) string { return x + "!" }, ",")
		actual := args.Map{"val": result}
		expected := args.Map{"val": "a!,b!"}
		expected.ShouldBeEqual(t, 0, "TranspileJoin -- joined", actual)
	})
}

// ── Join variants ───────────────────────────────────────────────────────────

func Test_Seg4_SS_Join(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Join", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"val": s.Join(",")}
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Join -- comma", actual)
	})
}

func Test_Seg4_SS_Join_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Join_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"val": s.Join(",")}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Join empty -- empty", actual)
	})
}

func Test_Seg4_SS_JoinLine(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JoinLine", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"nonEmpty": s.JoinLine() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinLine -- non-empty", actual)
	})
}

func Test_Seg4_SS_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JoinLine_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"val": s.JoinLine()}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "JoinLine empty -- empty", actual)
	})
}

func Test_Seg4_SS_JoinLineEofLine(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JoinLineEofLine", func() {
		s := corestr.SimpleSlice{"a", "b"}
		result := s.JoinLineEofLine()
		actual := args.Map{"nonEmpty": result != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinLineEofLine -- non-empty", actual)
	})
}

func Test_Seg4_SS_JoinLineEofLine_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JoinLineEofLine_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"val": s.JoinLineEofLine()}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "JoinLineEofLine empty -- empty", actual)
	})
}

func Test_Seg4_SS_JoinSpace(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JoinSpace", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"val": s.JoinSpace()}
		expected := args.Map{"val": "a b"}
		expected.ShouldBeEqual(t, 0, "JoinSpace -- space separated", actual)
	})
}

func Test_Seg4_SS_JoinComma(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JoinComma", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"val": s.JoinComma()}
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "JoinComma -- comma separated", actual)
	})
}

func Test_Seg4_SS_JoinWith(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JoinWith", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"nonEmpty": s.JoinWith(",") != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinWith -- non-empty", actual)
	})
}

func Test_Seg4_SS_JoinWith_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JoinWith_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"val": s.JoinWith(",")}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "JoinWith empty -- empty", actual)
	})
}

func Test_Seg4_SS_JoinCsv(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JoinCsv", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"nonEmpty": s.JoinCsv() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinCsv -- non-empty", actual)
	})
}

func Test_Seg4_SS_JoinCsvLine(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JoinCsvLine", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"nonEmpty": s.JoinCsvLine() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinCsvLine -- non-empty", actual)
	})
}

func Test_Seg4_SS_JoinCsvString(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JoinCsvString", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"nonEmpty": s.JoinCsvString(",") != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinCsvString -- non-empty", actual)
	})
}

func Test_Seg4_SS_JoinCsvString_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JoinCsvString_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"val": s.JoinCsvString(",")}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "JoinCsvString empty -- empty", actual)
	})
}

// ── EachItemSplitBy / PrependJoin / AppendJoin / PrependAppend ──────────────

func Test_Seg4_SS_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_EachItemSplitBy", func() {
		s := corestr.SimpleSlice{"a,b", "c,d"}
		result := s.EachItemSplitBy(",")
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "EachItemSplitBy -- 4 items", actual)
	})
}

func Test_Seg4_SS_PrependJoin(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_PrependJoin", func() {
		s := corestr.SimpleSlice{"c", "d"}
		result := s.PrependJoin(",", "a", "b")
		actual := args.Map{"nonEmpty": result != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "PrependJoin -- non-empty", actual)
	})
}

func Test_Seg4_SS_AppendJoin(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AppendJoin", func() {
		s := corestr.SimpleSlice{"a", "b"}
		result := s.AppendJoin(",", "c", "d")
		actual := args.Map{"nonEmpty": result != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AppendJoin -- non-empty", actual)
	})
}

func Test_Seg4_SS_PrependAppend(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_PrependAppend", func() {
		s := corestr.SimpleSlice{"b"}
		s.PrependAppend([]string{"a"}, []string{"c"})
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "PrependAppend -- 3 items", actual)
	})
}

func Test_Seg4_SS_PrependAppend_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_PrependAppend_Empty", func() {
		s := corestr.SimpleSlice{"b"}
		s.PrependAppend(nil, nil)
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "PrependAppend empty -- no change", actual)
	})
}

// ── IsEqual / IsEqualLines / IsEqualUnorderedLines ──────────────────────────

func Test_Seg4_SS_IsEqual(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqual", func() {
		s1 := corestr.SimpleSlice{"a", "b"}
		s2 := corestr.SimpleSlice{"a", "b"}
		s3 := corestr.SimpleSlice{"x"}
		actual := args.Map{
			"eq":       s1.IsEqual(&s2),
			"neq":      s1.IsEqual(&s3),
			"nilBoth":  (*corestr.SimpleSlice)(nil).IsEqual(nil),
			"nilOne":   s1.IsEqual(nil),
		}
		expected := args.Map{
			"eq":       true,
			"neq":      false,
			"nilBoth":  true,
			"nilOne":   false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqual -- various cases", actual)
	})
}

func Test_Seg4_SS_IsEqual_EmptyBoth(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqual_EmptyBoth", func() {
		s1 := corestr.SimpleSlice{}
		s2 := corestr.SimpleSlice{}
		actual := args.Map{"eq": s1.IsEqual(&s2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqual empty both -- true", actual)
	})
}

func Test_Seg4_SS_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualLines", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{
			"eq":       s.IsEqualLines([]string{"a", "b"}),
			"neq":      s.IsEqualLines([]string{"a", "c"}),
			"diffLen":  s.IsEqualLines([]string{"a"}),
			"nilBoth":  (*corestr.SimpleSlice)(nil).IsEqualLines(nil),
			"nilLeft":  (*corestr.SimpleSlice)(nil).IsEqualLines([]string{"a"}),
		}
		expected := args.Map{
			"eq":       true,
			"neq":      false,
			"diffLen":  false,
			"nilBoth":  true,
			"nilLeft":  false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualLines -- various", actual)
	})
}

func Test_Seg4_SS_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualUnorderedLines", func() {
		s := corestr.SimpleSlice{"b", "a"}
		actual := args.Map{
			"eq":      s.IsEqualUnorderedLines([]string{"a", "b"}),
			"diffLen": s.IsEqualUnorderedLines([]string{"a"}),
		}
		expected := args.Map{
			"eq":      true,
			"diffLen": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLines -- sorted comparison", actual)
	})
}

func Test_Seg4_SS_IsEqualUnorderedLines_NilBoth(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualUnorderedLines_NilBoth", func() {
		actual := args.Map{"eq": (*corestr.SimpleSlice)(nil).IsEqualUnorderedLines(nil)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLines nil -- true", actual)
	})
}

func Test_Seg4_SS_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualUnorderedLinesClone", func() {
		s := corestr.SimpleSlice{"b", "a"}
		actual := args.Map{"eq": s.IsEqualUnorderedLinesClone([]string{"a", "b"})}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLinesClone -- sorted clone comparison", actual)
	})
}

func Test_Seg4_SS_IsEqualUnorderedLinesClone_NilBoth(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualUnorderedLinesClone_NilBoth", func() {
		actual := args.Map{"eq": (*corestr.SimpleSlice)(nil).IsEqualUnorderedLinesClone(nil)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLinesClone nil -- true", actual)
	})
}

func Test_Seg4_SS_IsEqualUnorderedLinesClone_DiffLen(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualUnorderedLinesClone_DiffLen", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"eq": s.IsEqualUnorderedLinesClone([]string{"a"})}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLinesClone diff len -- false", actual)
	})
}

func Test_Seg4_SS_IsEqualUnorderedLinesClone_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualUnorderedLinesClone_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"eq": s.IsEqualUnorderedLinesClone([]string{})}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLinesClone empty -- true", actual)
	})
}

// ── IsUnorderedEqual / IsUnorderedEqualRaw ──────────────────────────────────

func Test_Seg4_SS_IsUnorderedEqualRaw_Clone(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsUnorderedEqualRaw_Clone", func() {
		s := corestr.SimpleSlice{"b", "a"}
		actual := args.Map{"eq": s.IsUnorderedEqualRaw(true, "a", "b")}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqualRaw clone -- equal", actual)
	})
}

func Test_Seg4_SS_IsUnorderedEqualRaw_NoClone(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsUnorderedEqualRaw_NoClone", func() {
		s := corestr.SimpleSlice{"b", "a"}
		actual := args.Map{"eq": s.IsUnorderedEqualRaw(false, "a", "b")}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqualRaw no clone -- equal", actual)
	})
}

func Test_Seg4_SS_IsUnorderedEqualRaw_DiffLen(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsUnorderedEqualRaw_DiffLen", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{"eq": s.IsUnorderedEqualRaw(true, "a", "b")}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqualRaw diff len -- false", actual)
	})
}

func Test_Seg4_SS_IsUnorderedEqualRaw_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsUnorderedEqualRaw_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"eq": s.IsUnorderedEqualRaw(true)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqualRaw empty -- true", actual)
	})
}

func Test_Seg4_SS_IsUnorderedEqual(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsUnorderedEqual", func() {
		s := corestr.SimpleSlice{"b", "a"}
		r := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"eq": s.IsUnorderedEqual(true, &r)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqual -- equal", actual)
	})
}

func Test_Seg4_SS_IsUnorderedEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsUnorderedEqual_BothEmpty", func() {
		s := corestr.SimpleSlice{}
		r := corestr.SimpleSlice{}
		actual := args.Map{"eq": s.IsUnorderedEqual(true, &r)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqual both empty -- true", actual)
	})
}

func Test_Seg4_SS_IsUnorderedEqual_NilRight(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsUnorderedEqual_NilRight", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{"eq": s.IsUnorderedEqual(true, nil)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqual nil right -- false", actual)
	})
}

// ── IsEqualByFunc / IsEqualByFuncLinesSplit ─────────────────────────────────

func Test_Seg4_SS_IsEqualByFunc(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualByFunc", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{
			"eq":  s.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b"),
			"neq": s.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "c"),
		}
		expected := args.Map{"eq": true, "neq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualByFunc -- match and mismatch", actual)
	})
}

func Test_Seg4_SS_IsEqualByFunc_DiffLen(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualByFunc_DiffLen", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{"eq": s.IsEqualByFunc(func(i int, l, r string) bool { return true }, "a", "b")}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualByFunc diff len -- false", actual)
	})
}

func Test_Seg4_SS_IsEqualByFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualByFunc_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"eq": s.IsEqualByFunc(func(i int, l, r string) bool { return true })}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualByFunc empty -- true", actual)
	})
}

func Test_Seg4_SS_IsEqualByFuncLinesSplit(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualByFuncLinesSplit", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{
			"eq": s.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return l == r }),
		}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualByFuncLinesSplit -- match", actual)
	})
}

func Test_Seg4_SS_IsEqualByFuncLinesSplit_Trim(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualByFuncLinesSplit_Trim", func() {
		s := corestr.SimpleSlice{" a ", " b "}
		actual := args.Map{
			"eq": s.IsEqualByFuncLinesSplit(true, ",", "a,b", func(i int, l, r string) bool { return l == r }),
		}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualByFuncLinesSplit trim -- match", actual)
	})
}

func Test_Seg4_SS_IsEqualByFuncLinesSplit_DiffLen(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualByFuncLinesSplit_DiffLen", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{
			"eq": s.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return true }),
		}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualByFuncLinesSplit diff len -- false", actual)
	})
}

func Test_Seg4_SS_IsEqualByFuncLinesSplit_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsEqualByFuncLinesSplit_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{
			"eq": s.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool { return true }),
		}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualByFuncLinesSplit empty -- mismatch on split", actual)
	})
}

// ── IsDistinctEqual / IsDistinctEqualRaw ────────────────────────────────────

func Test_Seg4_SS_IsDistinctEqualRaw(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsDistinctEqualRaw", func() {
		s := corestr.SimpleSlice{"a", "b", "a"}
		actual := args.Map{"eq": s.IsDistinctEqualRaw("a", "b")}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsDistinctEqualRaw -- distinct equal", actual)
	})
}

func Test_Seg4_SS_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_IsDistinctEqual", func() {
		s := corestr.SimpleSlice{"a", "b"}
		r := corestr.SimpleSlice{"b", "a"}
		actual := args.Map{"eq": s.IsDistinctEqual(&r)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsDistinctEqual -- distinct equal", actual)
	})
}

// ── Collection / ToCollection ───────────────────────────────────────────────

func Test_Seg4_SS_Collection(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Collection", func() {
		s := corestr.SimpleSlice{"a", "b"}
		c := s.Collection(true)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection -- 2 items", actual)
	})
}

func Test_Seg4_SS_ToCollection(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_ToCollection", func() {
		s := corestr.SimpleSlice{"a"}
		c := s.ToCollection(false)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ToCollection -- 1 item", actual)
	})
}

// ── NonPtr / Ptr / ToPtr / ToNonPtr ─────────────────────────────────────────

func Test_Seg4_SS_PtrNonPtr(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_PtrNonPtr", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{
			"nonPtrLen": len(s.NonPtr()),
			"ptrLen":    s.Ptr().Length(),
			"toPtrLen":  s.ToPtr().Length(),
			"toNonLen":  len(s.ToNonPtr()),
		}
		expected := args.Map{
			"nonPtrLen": 1,
			"ptrLen":    1,
			"toPtrLen":  1,
			"toNonLen":  1,
		}
		expected.ShouldBeEqual(t, 0, "Ptr/NonPtr -- all same", actual)
	})
}

// ── String ──────────────────────────────────────────────────────────────────

func Test_Seg4_SS_String(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_String", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"nonEmpty": s.String() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg4_SS_String_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_String_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"val": s.String()}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "String empty -- empty", actual)
	})
}

// ── ConcatNew / ConcatNewStrings / ConcatNewSimpleSlices ────────────────────

func Test_Seg4_SS_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_ConcatNew", func() {
		s := corestr.SimpleSlice{"a"}
		result := s.ConcatNew("b", "c")
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ConcatNew -- 3 items", actual)
	})
}

func Test_Seg4_SS_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_ConcatNewStrings", func() {
		s := corestr.SimpleSlice{"a"}
		result := s.ConcatNewStrings("b")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNewStrings -- 2 items", actual)
	})
}

func Test_Seg4_SS_ConcatNewStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_ConcatNewStrings_Nil", func() {
		var s *corestr.SimpleSlice
		result := s.ConcatNewStrings("a")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ConcatNewStrings nil -- clone of items", actual)
	})
}

func Test_Seg4_SS_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_ConcatNewSimpleSlices", func() {
		s := corestr.SimpleSlice{"a"}
		s2 := corestr.SimpleSlice{"b", "c"}
		result := s.ConcatNewSimpleSlices(&s2)
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ConcatNewSimpleSlices -- 3 items", actual)
	})
}

// ── CsvStrings ──────────────────────────────────────────────────────────────

func Test_Seg4_SS_CsvStrings(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_CsvStrings", func() {
		s := corestr.SimpleSlice{"a", "b"}
		actual := args.Map{"len": len(s.CsvStrings())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CsvStrings -- 2 quoted", actual)
	})
}

func Test_Seg4_SS_CsvStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_CsvStrings_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"len": len(s.CsvStrings())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CsvStrings empty -- empty", actual)
	})
}

// ── Sort / Reverse ──────────────────────────────────────────────────────────

func Test_Seg4_SS_Sort(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Sort", func() {
		s := corestr.SimpleSlice{"c", "a", "b"}
		s.Sort()
		actual := args.Map{"first": s.First()}
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Sort -- ascending", actual)
	})
}

func Test_Seg4_SS_Reverse(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Reverse", func() {
		s := corestr.SimpleSlice{"a", "b", "c", "d"}
		s.Reverse()
		actual := args.Map{"first": s.First(), "last": s.Last()}
		expected := args.Map{"first": "d", "last": "a"}
		expected.ShouldBeEqual(t, 0, "Reverse -- reversed", actual)
	})
}

func Test_Seg4_SS_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Reverse_Two", func() {
		s := corestr.SimpleSlice{"a", "b"}
		s.Reverse()
		actual := args.Map{"first": s.First()}
		expected := args.Map{"first": "b"}
		expected.ShouldBeEqual(t, 0, "Reverse 2 -- swapped", actual)
	})
}

func Test_Seg4_SS_Reverse_Single(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Reverse_Single", func() {
		s := corestr.SimpleSlice{"a"}
		s.Reverse()
		actual := args.Map{"first": s.First()}
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Reverse 1 -- unchanged", actual)
	})
}

// ── JSON / Serialize / Deserialize ──────────────────────────────────────────

func Test_Seg4_SS_Json(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Json", func() {
		s := corestr.SimpleSlice{"a"}
		j := s.Json()
		actual := args.Map{"noErr": !j.HasError()}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg4_SS_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_MarshalJSON", func() {
		s := corestr.SimpleSlice{"a"}
		b, err := s.MarshalJSON()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg4_SS_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_UnmarshalJSON", func() {
		s := corestr.SimpleSlice{}
		err := s.UnmarshalJSON([]byte(`["a","b"]`))
		actual := args.Map{"noErr": err == nil, "len": s.Length()}
		expected := args.Map{"noErr": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Seg4_SS_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_UnmarshalJSON_Invalid", func() {
		s := corestr.SimpleSlice{}
		err := s.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Seg4_SS_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Serialize", func() {
		s := corestr.SimpleSlice{"a"}
		b, err := s.Serialize()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_Seg4_SS_Deserialize(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Deserialize", func() {
		s := corestr.SimpleSlice{"a", "b"}
		var dest []string
		err := s.Deserialize(&dest)
		actual := args.Map{"noErr": err == nil, "len": len(dest)}
		expected := args.Map{"noErr": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

func Test_Seg4_SS_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_ParseInjectUsingJson", func() {
		s := corestr.SimpleSlice{"a"}
		jr := s.JsonPtr()
		s2 := corestr.SimpleSlice{}
		result, err := s2.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil, "len": result.Length()}
		expected := args.Map{"noErr": true, "len": 1}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- round trip", actual)
	})
}

func Test_Seg4_SS_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_ParseInjectUsingJsonMust", func() {
		s := corestr.SimpleSlice{"a"}
		jr := s.JsonPtr()
		s2 := corestr.SimpleSlice{}
		result := s2.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg4_SS_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_ParseInjectUsingJsonMust_Panic", func() {
		defer func() { recover() }()
		s := corestr.SimpleSlice{}
		_ = s.ParseInjectUsingJsonMust(&corejson.Result{})
	})
}

func Test_Seg4_SS_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JsonParseSelfInject", func() {
		s := corestr.SimpleSlice{"a"}
		jr := s.JsonPtr()
		s2 := corestr.SimpleSlice{}
		err := s2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

func Test_Seg4_SS_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JsonModel", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{"len": len(s.JsonModel())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel -- returns items", actual)
	})
}

func Test_Seg4_SS_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_JsonModelAny", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{"notNil": s.JsonModelAny() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

// ── Interface casts ─────────────────────────────────────────────────────────

func Test_Seg4_SS_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_InterfaceCasts", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{
			"binder":   s.AsJsonContractsBinder() != nil,
			"jsoner":   s.AsJsoner() != nil,
			"injector": s.AsJsonParseSelfInjector() != nil,
			"marsh":    s.AsJsonMarshaller() != nil,
		}
		expected := args.Map{
			"binder":   true,
			"jsoner":   true,
			"injector": true,
			"marsh":    true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

// ── Clone / DeepClone / ShadowClone ─────────────────────────────────────────

func Test_Seg4_SS_Clone(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Clone", func() {
		s := corestr.SimpleSlice{"a", "b"}
		c := s.Clone(true)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Clone deep -- same items", actual)
	})
}
	safeTest(t, "Test_Seg4_SS_DeepClone", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{"len": s.DeepClone().Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DeepClone -- 1 item", actual)
	})
}

func Test_Seg4_SS_ShadowClone(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_ShadowClone", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{"len": s.ShadowClone().Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ShadowClone -- 1 item", actual)
	})
}

// ── Clear / Dispose ─────────────────────────────────────────────────────────

func Test_Seg4_SS_Clear(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Clear", func() {
		s := corestr.SimpleSlice{"a", "b"}
		s.Clear()
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg4_SS_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Clear_Nil", func() {
		var s *corestr.SimpleSlice
		actual := args.Map{"nil": s.Clear() == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clear nil -- returns nil", actual)
	})
}

func Test_Seg4_SS_Dispose(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Dispose", func() {
		s := corestr.SimpleSlice{"a"}
		s.Dispose()
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleared", actual)
	})
}

func Test_Seg4_SS_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_Dispose_Nil", func() {
		var s *corestr.SimpleSlice
		s.Dispose() // should not panic
	})
}

// ── DistinctDiff / DistinctDiffRaw ──────────────────────────────────────────

func Test_Seg4_SS_DistinctDiffRaw(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_DistinctDiffRaw", func() {
		s := corestr.SimpleSlice{"a", "b", "c"}
		diff := s.DistinctDiffRaw("a", "d")
		actual := args.Map{"nonEmpty": len(diff) > 0}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "DistinctDiffRaw -- has diff", actual)
	})
}

func Test_Seg4_SS_DistinctDiffRaw_BothNil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_DistinctDiffRaw_BothNil", func() {
		var s *corestr.SimpleSlice
		diff := s.DistinctDiffRaw(nil...)
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "DistinctDiffRaw both nil -- empty", actual)
	})
}

func Test_Seg4_SS_DistinctDiffRaw_LeftNil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_DistinctDiffRaw_LeftNil", func() {
		var s *corestr.SimpleSlice
		diff := s.DistinctDiffRaw("a")
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffRaw left nil -- returns right", actual)
	})
}

func Test_Seg4_SS_DistinctDiffRaw_RightNil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_DistinctDiffRaw_RightNil", func() {
		s := corestr.SimpleSlice{"a"}
		diff := s.DistinctDiffRaw(nil...)
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffRaw right nil -- returns left", actual)
	})
}

func Test_Seg4_SS_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_DistinctDiff", func() {
		s := corestr.SimpleSlice{"a", "b"}
		r := corestr.SimpleSlice{"a", "c"}
		diff := s.DistinctDiff(&r)
		actual := args.Map{"nonEmpty": len(diff) > 0}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "DistinctDiff -- has diff", actual)
	})
}

func Test_Seg4_SS_DistinctDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_DistinctDiff_BothNil", func() {
		var s *corestr.SimpleSlice
		diff := s.DistinctDiff(nil)
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "DistinctDiff both nil -- empty", actual)
	})
}

func Test_Seg4_SS_DistinctDiff_LeftNil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_DistinctDiff_LeftNil", func() {
		var s *corestr.SimpleSlice
		r := corestr.SimpleSlice{"a"}
		diff := s.DistinctDiff(&r)
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiff left nil -- right", actual)
	})
}

func Test_Seg4_SS_DistinctDiff_RightNil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_DistinctDiff_RightNil", func() {
		s := corestr.SimpleSlice{"a"}
		diff := s.DistinctDiff(nil)
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiff right nil -- left", actual)
	})
}

// ── AddedRemovedLinesDiff ───────────────────────────────────────────────────

func Test_Seg4_SS_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddedRemovedLinesDiff", func() {
		s := corestr.SimpleSlice{"a", "b"}
		added, removed := s.AddedRemovedLinesDiff("b", "c")
		actual := args.Map{"addedLen": len(added), "removedLen": len(removed)}
		expected := args.Map{"addedLen": 1, "removedLen": 1}
		expected.ShouldBeEqual(t, 0, "AddedRemovedLinesDiff -- 1 added 1 removed", actual)
	})
}

func Test_Seg4_SS_AddedRemovedLinesDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_AddedRemovedLinesDiff_BothNil", func() {
		var s *corestr.SimpleSlice
		added, removed := s.AddedRemovedLinesDiff(nil...)
		actual := args.Map{"addedNil": added == nil, "removedNil": removed == nil}
		expected := args.Map{"addedNil": true, "removedNil": true}
		expected.ShouldBeEqual(t, 0, "AddedRemovedLinesDiff both nil -- nil", actual)
	})
}

// ── RemoveIndexes ───────────────────────────────────────────────────────────

func Test_Seg4_SS_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_RemoveIndexes", func() {
		s := corestr.SimpleSlice{"a", "b", "c", "d"}
		result, err := s.RemoveIndexes(1, 3)
		actual := args.Map{"len": result.Length(), "noErr": err == nil}
		expected := args.Map{"len": 2, "noErr": true}
		expected.ShouldBeEqual(t, 0, "RemoveIndexes -- 2 remaining", actual)
	})
}

func Test_Seg4_SS_RemoveIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_RemoveIndexes_Empty", func() {
		s := corestr.SimpleSlice{}
		_, err := s.RemoveIndexes(0)
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "RemoveIndexes empty -- error", actual)
	})
}

func Test_Seg4_SS_RemoveIndexes_OutOfBounds(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_RemoveIndexes_OutOfBounds", func() {
		s := corestr.SimpleSlice{"a", "b"}
		result, err := s.RemoveIndexes(5)
		actual := args.Map{"hasErr": err != nil, "len": result.Length()}
		expected := args.Map{"hasErr": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveIndexes out of bounds -- error with all items", actual)
	})
}

// ── SafeStrings ─────────────────────────────────────────────────────────────

func Test_Seg4_SS_SafeStrings(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_SafeStrings", func() {
		s := corestr.SimpleSlice{"a"}
		actual := args.Map{"len": len(s.SafeStrings())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "SafeStrings -- 1 item", actual)
	})
}

func Test_Seg4_SS_SafeStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg4_SS_SafeStrings_Empty", func() {
		s := corestr.SimpleSlice{}
		actual := args.Map{"len": len(s.SafeStrings())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SafeStrings empty -- empty", actual)
	})
}
