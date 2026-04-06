package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ── Basic ──

func Test_Src04_RawErrCollection_Basic(t *testing.T) {
	// Arrange
	tc := rawErrCollBasicTestCases[0]
	// Act
	c := &errcore.RawErrCollection{}
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"isEmpty":          c.IsEmpty(),
		"hasError":         c.HasError(),
		"hasAnyError":      c.HasAnyError(),
		"len":              c.Length(),
		"hasAnyIssues":     c.HasAnyIssues(),
		"isValid":          c.IsValid(),
		"isSuccess":        c.IsSuccess(),
		"isFailed":         c.IsFailed(),
		"isInvalid":        c.IsInvalid(),
		"isCollectionType": c.IsCollectionType(),
	})
}

// ── Add methods ──

func Test_Src04_RawErrCollection_Add(t *testing.T) {
	tc := rawErrCollAddTestCases[0]
	c := &errcore.RawErrCollection{}
	c.Add(nil)
	c.Add(errors.New("e1"))
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddError(t *testing.T) {
	tc := rawErrCollAddTestCases[1]
	c := &errcore.RawErrCollection{}
	c.AddError(nil)
	c.AddError(errors.New("e"))
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_Adds(t *testing.T) {
	tc := rawErrCollAddTestCases[2]
	c := &errcore.RawErrCollection{}
	c.Adds()
	c.Adds(errors.New("a"), nil, errors.New("b"))
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddErrors(t *testing.T) {
	tc := rawErrCollAddTestCases[3]
	c := &errcore.RawErrCollection{}
	c.AddErrors(errors.New("a"))
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddString(t *testing.T) {
	tc := rawErrCollAddTestCases[4]
	c := &errcore.RawErrCollection{}
	c.AddString("")
	c.AddString("hello")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddMsg(t *testing.T) {
	tc := rawErrCollAddTestCases[5]
	c := &errcore.RawErrCollection{}
	c.AddMsg("hello")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddIf(t *testing.T) {
	tc := rawErrCollAddTestCases[6]
	c := &errcore.RawErrCollection{}
	c.AddIf(false, "skip")
	c.AddIf(true, "add")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddFunc(t *testing.T) {
	tc := rawErrCollAddTestCases[7]
	c := &errcore.RawErrCollection{}
	c.AddFunc(nil)
	c.AddFunc(func() error { return nil })
	c.AddFunc(func() error { return errors.New("e") })
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddFuncIf(t *testing.T) {
	tc := rawErrCollAddTestCases[8]
	c := &errcore.RawErrCollection{}
	c.AddFuncIf(false, func() error { return errors.New("e") })
	c.AddFuncIf(true, nil)
	c.AddFuncIf(true, func() error { return errors.New("e") })
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_ConditionalAddError(t *testing.T) {
	tc := rawErrCollAddTestCases[9]
	c := &errcore.RawErrCollection{}
	c.ConditionalAddError(false, errors.New("e"))
	c.ConditionalAddError(true, errors.New("e"))
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddMsgStackTrace(t *testing.T) {
	tc := rawErrCollAddTestCases[10]
	c := &errcore.RawErrCollection{}
	c.AddMsgStackTrace("")
	c.AddMsgStackTrace("msg")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddStackTrace(t *testing.T) {
	tc := rawErrCollAddTestCases[11]
	c := &errcore.RawErrCollection{}
	c.AddStackTrace(nil)
	c.AddStackTrace(errors.New("e"))
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddMsgErrStackTrace(t *testing.T) {
	tc := rawErrCollAddTestCases[12]
	c := &errcore.RawErrCollection{}
	c.AddMsgErrStackTrace("msg", nil)
	c.AddMsgErrStackTrace("msg", errors.New("e"))
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddMethodName(t *testing.T) {
	tc := rawErrCollAddTestCases[13]
	c := &errcore.RawErrCollection{}
	c.AddMethodName("")
	c.AddMethodName("msg")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddMessages(t *testing.T) {
	tc := rawErrCollAddTestCases[14]
	c := &errcore.RawErrCollection{}
	c.AddMessages()
	c.AddMessages("a", "b")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddErrorWithMessage(t *testing.T) {
	tc := rawErrCollAddTestCases[15]
	c := &errcore.RawErrCollection{}
	c.AddErrorWithMessage(nil, "msg")
	c.AddErrorWithMessage(errors.New("e"), "msg")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddErrorWithMessageRef(t *testing.T) {
	tc := rawErrCollAddTestCases[16]
	c := &errcore.RawErrCollection{}
	c.AddErrorWithMessageRef(nil, "msg", "ref")
	c.AddErrorWithMessageRef(errors.New("e"), "msg", nil)
	c.AddErrorWithMessageRef(errors.New("e"), "msg", "ref")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddFmt(t *testing.T) {
	tc := rawErrCollAddTestCases[17]
	c := &errcore.RawErrCollection{}
	c.AddFmt(nil, "fmt %s", "v")
	c.AddFmt(errors.New("e"), "fmt %s", "v")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_Fmt(t *testing.T) {
	tc := rawErrCollAddTestCases[18]
	c := &errcore.RawErrCollection{}
	c.Fmt("")
	c.Fmt("hello %s", "world")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_FmtIf(t *testing.T) {
	tc := rawErrCollAddTestCases[19]
	c := &errcore.RawErrCollection{}
	c.FmtIf(false, "skip")
	c.FmtIf(true, "add %s", "v")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_References(t *testing.T) {
	tc := rawErrCollAddTestCases[20]
	c := &errcore.RawErrCollection{}
	c.References("msg", "ref1")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddWithRef(t *testing.T) {
	tc := rawErrCollAddTestCases[21]
	c := &errcore.RawErrCollection{}
	c.AddWithRef(nil, "ref")
	c.AddWithRef(errors.New("e"), "ref")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddWithTraceRef(t *testing.T) {
	tc := rawErrCollAddTestCases[22]
	c := &errcore.RawErrCollection{}
	c.AddWithTraceRef(nil, []string{"t"}, "r")
	c.AddWithTraceRef(errors.New("e"), []string{"t"}, "r")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddWithCompiledTraceRef(t *testing.T) {
	tc := rawErrCollAddTestCases[23]
	c := &errcore.RawErrCollection{}
	c.AddWithCompiledTraceRef(nil, "t", "r")
	c.AddWithCompiledTraceRef(errors.New("e"), "t", "r")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddStringSliceAsErr(t *testing.T) {
	tc := rawErrCollAddTestCases[24]
	c := &errcore.RawErrCollection{}
	c.AddStringSliceAsErr()
	c.AddStringSliceAsErr("a", "", "b")
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})
}

func Test_Src04_RawErrCollection_AddErrorGetters(t *testing.T) {
	tc := rawErrCollAddTestCases[25]
	c := &errcore.RawErrCollection{}
	noPanic := !callPanicsErrcore(func() { c.AddErrorGetters() })
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src04_RawErrCollection_AddCompiledErrorGetters(t *testing.T) {
	tc := rawErrCollAddTestCases[26]
	c := &errcore.RawErrCollection{}
	noPanic := !callPanicsErrcore(func() { c.AddCompiledErrorGetters() })
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

// ── String/Error methods ──

func Test_Src04_RawErrCollection_Strings(t *testing.T) {
	tc := rawErrCollStringTestCases[0]
	c := &errcore.RawErrCollection{}
	emptyLen := len(c.Strings())
	c.Add(errors.New("a"))
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"emptyLen": emptyLen,
		"oneLen":   len(c.Strings()),
	})
}

func Test_Src04_RawErrCollection_String(t *testing.T) {
	tc := rawErrCollStringTestCases[1]
	c := &errcore.RawErrCollection{}
	emptyStr := c.String()
	c.Add(errors.New("a"))
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"emptyStr": emptyStr,
		"nonEmpty": c.String() != "",
	})
}

func Test_Src04_RawErrCollection_StringUsingJoiner(t *testing.T) {
	tc := rawErrCollStringTestCases[2]
	c := &errcore.RawErrCollection{}
	emptyStr := c.StringUsingJoiner(",")
	c.Add(errors.New("a"))
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"emptyStr": emptyStr,
		"nonEmpty": c.StringUsingJoiner(",") != "",
	})
}

func Test_Src04_RawErrCollection_StringUsingJoinerAdditional(t *testing.T) {
	tc := rawErrCollStringTestCases[3]
	c := &errcore.RawErrCollection{}
	emptyStr := c.StringUsingJoinerAdditional(",", "!")
	c.Add(errors.New("a"))
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"emptyStr": emptyStr,
		"nonEmpty": c.StringUsingJoinerAdditional(",", "!") != "",
	})
}

func Test_Src04_RawErrCollection_StringWithAdditionalMessage(t *testing.T) {
	tc := rawErrCollStringTestCases[4]
	c := &errcore.RawErrCollection{}
	emptyStr := c.StringWithAdditionalMessage("!")
	c.Add(errors.New("a"))
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"emptyStr": emptyStr,
		"nonEmpty": c.StringWithAdditionalMessage("!") != "",
	})
}

// ── Compiled methods ──

func Test_Src04_RawErrCollection_CompiledError(t *testing.T) {
	tc := rawErrCollCompiledTestCases[0]
	c := errcore.RawErrCollection{}
	emptyNil := c.CompiledError() == nil
	c.Items = append(c.Items, errors.New("a"))
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"emptyNil": emptyNil,
		"nonNil":   c.CompiledError() != nil,
	})
}

func Test_Src04_RawErrCollection_CompiledErrorUsingJoiner(t *testing.T) {
	tc := rawErrCollCompiledTestCases[1]
	c := &errcore.RawErrCollection{}
	emptyNil := c.CompiledErrorUsingJoiner(",") == nil
	c.Add(errors.New("a"))
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"emptyNil": emptyNil,
		"nonNil":   c.CompiledErrorUsingJoiner(",") != nil,
	})
}

func Test_Src04_RawErrCollection_CompiledErrorUsingJoinerAdditional(t *testing.T) {
	tc := rawErrCollCompiledTestCases[2]
	c := &errcore.RawErrCollection{}
	emptyNil := c.CompiledErrorUsingJoinerAdditionalMessage(",", "!") == nil
	c.Add(errors.New("a"))
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"emptyNil": emptyNil,
		"nonNil":   c.CompiledErrorUsingJoinerAdditionalMessage(",", "!") != nil,
	})
}

func Test_Src04_RawErrCollection_CompiledErrorWithStackTraces(t *testing.T) {
	tc := rawErrCollCompiledTestCases[3]
	c := &errcore.RawErrCollection{}
	emptyNil := c.CompiledErrorWithStackTraces() == nil
	c.Add(errors.New("a"))
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"emptyNil": emptyNil,
		"nonNil":   c.CompiledErrorWithStackTraces() != nil,
	})
}

// ── Misc methods ──

func Test_Src04_RawErrCollection_CompiledStackTracesString(t *testing.T) {
	c := &errcore.RawErrCollection{}
	empty := c.CompiledStackTracesString() == ""
	c.Add(errors.New("a"))
	nonEmpty := c.CompiledStackTracesString() != ""
	actual := args.Map{"result": empty || !nonEmpty}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Src04_RawErrCollection_CompiledErrorUsingStackTraces(t *testing.T) {
	c := &errcore.RawErrCollection{}
	emptyNil := c.CompiledErrorUsingStackTraces(",", []string{"t"}) == nil
	c.Add(errors.New("a"))
	nonNil := c.CompiledErrorUsingStackTraces(",", []string{"t"}) != nil
	actual := args.Map{"result": emptyNil || !nonNil}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Src04_RawErrCollection_CompiledJsonErrorWithStackTraces(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.Add(errors.New("e"))
	_ = c.CompiledJsonErrorWithStackTraces()
}

func Test_Src04_RawErrCollection_CompiledJsonStringWithStackTraces(t *testing.T) {
	c := &errcore.RawErrCollection{}
	_ = c.CompiledJsonStringWithStackTraces()
}

func Test_Src04_RawErrCollection_ErrorString(t *testing.T) {
	c := &errcore.RawErrCollection{}
	_ = c.ErrorString()
}

func Test_Src04_RawErrCollection_Compile(t *testing.T) {
	c := &errcore.RawErrCollection{}
	_ = c.Compile()
}

func Test_Src04_RawErrCollection_FullString(t *testing.T) {
	c := &errcore.RawErrCollection{}
	_ = c.FullString()
}

func Test_Src04_RawErrCollection_FullStringWithTraces(t *testing.T) {
	c := &errcore.RawErrCollection{}
	_ = c.FullStringWithTraces()
}

func Test_Src04_RawErrCollection_FullStringWithTracesIf(t *testing.T) {
	c := &errcore.RawErrCollection{}
	_ = c.FullStringWithTracesIf(true)
	_ = c.FullStringWithTracesIf(false)
}

func Test_Src04_RawErrCollection_FullStringSplitByNewLine(t *testing.T) {
	c := &errcore.RawErrCollection{}
	_ = c.FullStringSplitByNewLine()
}

func Test_Src04_RawErrCollection_FullStringWithoutReferences(t *testing.T) {
	c := &errcore.RawErrCollection{}
	_ = c.FullStringWithoutReferences()
}

func Test_Src04_RawErrCollection_ReferencesCompiledString(t *testing.T) {
	c := &errcore.RawErrCollection{}
	_ = c.ReferencesCompiledString()
}

func Test_Src04_RawErrCollection_Serialize(t *testing.T) {
	tc := rawErrCollMiscTestCases[0]
	c := &errcore.RawErrCollection{}
	b, err := c.Serialize()
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"bytesNil": b == nil,
		"errNil":   err == nil,
	})
}

func Test_Src04_RawErrCollection_SerializeWithoutTraces(t *testing.T) {
	tc := rawErrCollMiscTestCases[1]
	c := &errcore.RawErrCollection{}
	b, err := c.SerializeWithoutTraces()
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"bytesNil": b == nil,
		"errNil":   err == nil,
	})
}

func Test_Src04_RawErrCollection_MarshalJSON(t *testing.T) {
	tc := rawErrCollMiscTestCases[2]
	c := &errcore.RawErrCollection{}
	b, err := c.MarshalJSON()
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"bytesNil": b == nil,
		"errNil":   err == nil,
	})
}

func Test_Src04_RawErrCollection_UnmarshalJSON(t *testing.T) {
	tc := rawErrCollMiscTestCases[3]
	c := &errcore.RawErrCollection{}
	noPanic := !callPanicsErrcore(func() { _ = c.UnmarshalJSON([]byte("[]")) })
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src04_RawErrCollection_Value(t *testing.T) {
	tc := rawErrCollMiscTestCases[4]
	c := &errcore.RawErrCollection{}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": c.Value() == nil})
}

func Test_Src04_RawErrCollection_Log(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.Log()
	c.Add(errors.New("e"))
	c.Log()
}

func Test_Src04_RawErrCollection_LogWithTraces(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.LogWithTraces()
	c.Add(errors.New("e"))
	c.LogWithTraces()
}

func Test_Src04_RawErrCollection_LogIf(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.LogIf(false)
}

func Test_Src04_RawErrCollection_IsNull(t *testing.T) {
	c := &errcore.RawErrCollection{}
	_ = c.IsNull()
	_ = c.IsAnyNull()
}

func Test_Src04_RawErrCollection_ClearDispose(t *testing.T) {
	tc := rawErrCollMiscTestCases[5]
	c := &errcore.RawErrCollection{}
	c.Clear()
	c.Add(errors.New("e"))
	c.Clear()
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": c.Length()})

	c2 := &errcore.RawErrCollection{}
	c2.Dispose()
	c2.Add(errors.New("e"))
	c2.Dispose()
}

func Test_Src04_RawErrCollection_IsErrorsCollected(t *testing.T) {
	tc := rawErrCollMiscTestCases[6]
	c := &errcore.RawErrCollection{}
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"nilFalse": !c.IsErrorsCollected(nil),
		"errTrue":  c.IsErrorsCollected(errors.New("e")),
	})
}

func Test_Src04_RawErrCollection_CountStateChangeTracker(t *testing.T) {
	tc := rawErrCollMiscTestCases[7]
	c := errcore.RawErrCollection{}
	tracker := c.CountStateChangeTracker()
	tc.ShouldBeEqualMapFirst(t, args.Map{"isSame": tracker.IsSameState()})
}

func Test_Src04_RawErrCollection_ToRawErrCollection(t *testing.T) {
	tc := rawErrCollMiscTestCases[8]
	c := errcore.RawErrCollection{}
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": c.ToRawErrCollection() != nil})
}

func Test_Src04_RawErrCollection_ReflectSetTo(t *testing.T) {
	tc := rawErrCollMiscTestCases[9]
	c := &errcore.RawErrCollection{}
	valueErr := c.ReflectSetTo(errcore.RawErrCollection{}) != nil
	var nilP *errcore.RawErrCollection
	nilPtrErr := c.ReflectSetTo(nilP) != nil
	validNoErr := c.ReflectSetTo(&errcore.RawErrCollection{}) == nil
	otherErr := c.ReflectSetTo("other") != nil
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"valueErr":   valueErr,
		"nilPtrErr":  nilPtrErr,
		"validNoErr": validNoErr,
		"otherErr":   otherErr,
	})
}
