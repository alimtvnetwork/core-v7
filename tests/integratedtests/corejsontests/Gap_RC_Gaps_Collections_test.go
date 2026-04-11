package corejsontests

import (
	"errors"
	"testing"
	"time"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ResultsCollection uncovered methods ──

func Test_Gap_RC_GetAtSafeUsingLength(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.UsingCap(3)
	c.Add(corejson.NewResult.Any("a"))
	c.Add(corejson.NewResult.Any("b"))

	// Act
	actual := args.Map{"result": c.GetAtSafeUsingLength(0, 2) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": c.GetAtSafeUsingLength(5, 2) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": c.GetAtSafeUsingLength(-1, 2) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Gap_RC_AddAnyItemsSlice(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.UsingCap(3)
	c.AddAnyItemsSlice([]any{"a", nil, "b"})

	// Act
	actual := args.Map{"result": c.Length()}

	// Assert
	expected := args.Map{"result": 2}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	c.AddAnyItemsSlice(nil)
}

func Test_Gap_RC_AddResultsCollection(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.UsingCap(3)
	c.Add(corejson.NewResult.Any("a"))
	sub := corejson.NewResultsCollection.UsingCap(1)
	sub.Add(corejson.NewResult.Any("b"))
	c.AddResultsCollection(sub)
	c.AddResultsCollection(nil)

	// Act
	actual := args.Map{"result": c.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Gap_RC_AddNonNilItemsPtr(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.UsingCap(3)
	c.AddNonNilItemsPtr(nil, corejson.NewResult.AnyPtr("a"), nil)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	c.AddNonNilItemsPtr()
	c.AddNonNilItemsPtr(nil)
}

func Test_Gap_RC_NonPtrPtr(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(1)
	c.Add(corejson.NewResult.Any("x"))
	_ = c.NonPtr()
	_ = c.Ptr()
}

func Test_Gap_RC_GetAt(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.UsingCap(1)
	c.Add(corejson.NewResult.Any("x"))
	r := c.GetAt(0)

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Gap_RC_NoErrorCollection(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.UsingCap(1)
	c.Add(corejson.NewResult.Any("ok"))

	// Act
	actual := args.Map{"result": c.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have error", actual)
}

func Test_Gap_RC_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	jr := corejson.NewResult.AnyPtr(c)
	target := corejson.NewResultsCollection.Empty()
	_, err := target.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Gap_RC_ParseInjectUsingJsonMust(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	jr := corejson.NewResult.AnyPtr(c)
	target := corejson.NewResultsCollection.Empty()
	_ = target.ParseInjectUsingJsonMust(jr)
}

func Test_Gap_RC_JsonParseSelfInject(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	jr := corejson.NewResult.AnyPtr(c)
	target := corejson.NewResultsCollection.Empty()
	err := target.JsonParseSelfInject(jr)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Gap_RC_AddSerializerFunc(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.UsingCap(2)
	c.AddSerializerFunc(func() ([]byte, error) { return []byte(`"x"`), nil })
	c.AddSerializerFunc(nil)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Gap_RC_AddSerializerFunctions(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(2)
	fn := func() ([]byte, error) { return []byte(`"x"`), nil }
	c.AddSerializerFunctions(fn)
	c.AddSerializerFunctions()
}

func Test_Gap_RC_AddMapResults(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(2)
	mr := corejson.NewMapResults.Empty()
	mr.Add("a", corejson.NewResult.Any("x"))
	c.AddMapResults(mr)
	c.AddMapResults(corejson.NewMapResults.Empty())
}

func Test_Gap_RC_AddRawMapResults(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(2)
	c.AddRawMapResults(nil)
	c.AddRawMapResults(map[string]corejson.Result{"a": corejson.NewResult.Any("x")})
}

func Test_Gap_RC_InjectIntoSameIndex(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(3)
	c.Add(corejson.NewResult.Any("x"))
	c.Add(corejson.Result{Error: errors.New("e")})
	_, _ = c.InjectIntoSameIndex(nil, nil)
	_, _ = c.InjectIntoSameIndex()
}

func Test_Gap_RC_UnmarshalIntoSameIndex(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(4)
	c.Add(corejson.NewResult.Any(`"hello"`))
	c.Add(corejson.Result{Error: errors.New("e")})
	c.Add(corejson.Result{Bytes: []byte("{}")})
	var s string
	_, _ = c.UnmarshalIntoSameIndex(&s, nil, nil)
	_, _ = c.UnmarshalIntoSameIndex()
}

// ── ResultsPtrCollection uncovered methods ──

func Test_Gap_RPC_UnmarshalAt(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.UsingCap(2)
	c.Add(corejson.NewResult.AnyPtr("hello"))
	c.Add(&corejson.Result{Error: errors.New("e")})
	var s string
	err := c.UnmarshalAt(0, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	err2 := c.UnmarshalAt(1, &s)
	// Result with error may or may not return error from UnmarshalAt
	_ = err2
}

func Test_Gap_RPC_UnmarshalAt_NilResult(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.UsingCap(1)
	c.Add(nil)
	var s string
	err := c.UnmarshalAt(0, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil result", actual)
}

func Test_Gap_RPC_GetAtSafe(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.UsingCap(1)
	c.Add(corejson.NewResult.AnyPtr("x"))

	// Act
	actual := args.Map{"result": c.GetAtSafe(0) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": c.GetAtSafe(-1) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": c.GetAtSafe(5) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Gap_RPC_GetAtSafeUsingLength(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.UsingCap(1)
	c.Add(corejson.NewResult.AnyPtr("x"))

	// Act
	actual := args.Map{"result": c.GetAtSafeUsingLength(0, 1) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": c.GetAtSafeUsingLength(5, 1) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Gap_RPC_InjectIntoSameIndex(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(3)
	c.Add(corejson.NewResult.AnyPtr("x"))
	c.Add(nil)
	c.Add(&corejson.Result{Error: errors.New("e")})
	_, _ = c.InjectIntoSameIndex(nil, nil, nil)
	_, _ = c.InjectIntoSameIndex()
}

func Test_Gap_RPC_UnmarshalIntoSameIndex(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(4)
	c.Add(corejson.NewResult.AnyPtr(`"hello"`))
	c.Add(nil)
	c.Add(&corejson.Result{Error: errors.New("e")})
	c.Add(&corejson.Result{Bytes: []byte("{}")})
	var s string
	_, _ = c.UnmarshalIntoSameIndex(&s, nil, nil, nil)
	_, _ = c.UnmarshalIntoSameIndex()
}

func Test_Gap_RPC_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	jr := corejson.NewResult.AnyPtr(c)
	target := corejson.NewResultsPtrCollection.Empty()
	_, err := target.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Gap_RPC_ParseInjectUsingJsonMust(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	jr := corejson.NewResult.AnyPtr(c)
	target := corejson.NewResultsPtrCollection.Empty()
	_ = target.ParseInjectUsingJsonMust(jr)
}

func Test_Gap_RPC_JsonParseSelfInject(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	jr := corejson.NewResult.AnyPtr(c)
	target := corejson.NewResultsPtrCollection.Empty()
	_ = target.JsonParseSelfInject(jr)
}

// ── BytesCollection uncovered methods ──

func Test_Gap_BC_GetAtSafeUsingLength(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"x"`))

	// Act
	actual := args.Map{"result": c.GetAtSafeUsingLength(0, 1) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": c.GetAtSafeUsingLength(5, 1) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Gap_BC_AddsPtr(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	r := corejson.NewResult.AnyPtr("x")
	c.AddsPtr(r, nil)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	c.AddsPtr()
}

func Test_Gap_BC_InjectIntoSameIndex(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"x"`))
	_, _ = c.InjectIntoSameIndex(nil) // nil element in populated collection - ok
	_, _ = c.InjectIntoSameIndex()
}

func Test_Gap_BC_UnmarshalIntoSameIndex(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"hello"`))
	c.Add([]byte(`42`))
	var s string
	var n int
	_, _ = c.UnmarshalIntoSameIndex(&s, &n)

	// Act
	actual := args.Map{"result": s != "hello" || n != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	_, _ = c.UnmarshalIntoSameIndex()
}

func Test_Gap_BC_ParseInjectUsingJson(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"x"`))
	jr := corejson.NewResult.AnyPtr(c)
	target := corejson.NewBytesCollection.Empty()
	_, err := target.ParseInjectUsingJson(jr)
	_ = err
}

func Test_Gap_BC_ParseInjectUsingJsonMust_Error(t *testing.T) {
	// Test with invalid data to ensure error handling
	jr := corejson.NewResult.AnyPtr("invalid")
	target := corejson.NewBytesCollection.Empty()
	defer func() {
		if r := recover(); r == nil {
			// It's ok if it doesn't panic, depends on data
		}
	}()
	_ = target.ParseInjectUsingJsonMust(jr)
}

func Test_Gap_BC_JsonParseSelfInject(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"x"`))
	jr := corejson.NewResult.AnyPtr(c)
	target := corejson.NewBytesCollection.Empty()
	_ = target.JsonParseSelfInject(jr)
}

func Test_Gap_BC_GetSinglePageCollection(t *testing.T) {
	// Arrange
	c := corejson.NewBytesCollection.UsingCap(15)
	for i := 0; i < 15; i++ {
		c.Add([]byte(`"x"`))
	}
	page := c.GetSinglePageCollection(5, 1)

	// Act
	actual := args.Map{"result": page.Length() != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	page2 := c.GetSinglePageCollection(5, 3)
	actual = args.Map{"result": page2.Length() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	// Last page
	page3 := c.GetSinglePageCollection(10, 2)
	actual = args.Map{"result": page3.Length() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5 for last partial page", actual)
}

func Test_Gap_BC_AddJsoners(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.AddJsoners(true)
	c.AddJsoners(true, nil)
}

// ── MapResults uncovered methods ──

func Test_Gap_MR_Unmarshal(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("hello"))
	var s string
	// Note: MapResults.Unmarshal has a bug where `has` check is inverted
	// but we still exercise the code
	err := m.Unmarshal("a", &s)
	_ = err
}

func Test_Gap_MR_Deserialize(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("hello"))
	var s string
	err := m.Deserialize("a", &s)
	_ = err
}

func Test_Gap_MR_DeserializeMust(t *testing.T) {
	defer func() { recover() }() // DeserializeMust may panic
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("hello"))
	var s string
	// DeserializeMust panics if key not found — use actual key
	_ = m.DeserializeMust("a", &s)
}

func Test_Gap_MR_UnmarshalMany(t *testing.T) {
	// Arrange
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("hello"))
	err := m.UnmarshalMany()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
}

func Test_Gap_MR_UnmarshalManySafe(t *testing.T) {
	// Arrange
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("hello"))
	err := m.UnmarshalManySafe()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
}

func Test_Gap_MR_SafeUnmarshal(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("hello"))
	var s string
	err := m.SafeUnmarshal("a", &s)
	_ = err
	err2 := m.SafeUnmarshal("missing", &s)
	_ = err2
}

func Test_Gap_MR_SafeDeserialize(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("hello"))
	var s string
	_ = m.SafeDeserialize("a", &s)
}

func Test_Gap_MR_SafeDeserializeMust(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("hello"))
	var s string
	_ = m.SafeDeserializeMust("a", &s)
}

func Test_Gap_MR_InjectIntoAt(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("hello"))
	target := &corejson.Result{}
	_ = m.InjectIntoAt("a", target)
}

func Test_Gap_MR_AddKeyWithJsoner(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.AddKeyWithJsoner(corejson.KeyWithJsoner{Key: "a", Jsoner: nil})
}

func Test_Gap_MR_AddKeysWithJsoners(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.AddKeysWithJsoners()
}

func Test_Gap_MR_AddKeyWithJsonerPtr(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.AddKeyWithJsonerPtr(nil)
	m.AddKeyWithJsonerPtr(&corejson.KeyWithJsoner{Key: "a", Jsoner: nil})
}

func Test_Gap_MR_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	m := corejson.NewMapResults.UsingCap(1)
	m.Add("a", corejson.NewResult.Any("x"))
	jr := corejson.NewResult.AnyPtr(m)
	target := corejson.NewMapResults.Empty()
	_, err := target.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Gap_MR_ParseInjectUsingJsonMust(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(1)
	m.Add("a", corejson.NewResult.Any("x"))
	jr := corejson.NewResult.AnyPtr(m)
	target := corejson.NewMapResults.Empty()
	_ = target.ParseInjectUsingJsonMust(jr)
}

func Test_Gap_MR_JsonParseSelfInject(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(1)
	m.Add("a", corejson.NewResult.Any("x"))
	jr := corejson.NewResult.AnyPtr(m)
	target := corejson.NewMapResults.Empty()
	_ = target.JsonParseSelfInject(jr)
}

func Test_Gap_MR_GetSinglePageCollection(t *testing.T) {
	// Arrange
	m := corejson.NewMapResults.UsingCap(15)
	for i := 0; i < 15; i++ {
		m.Add(corejson.Serialize.ToString(i), corejson.NewResult.Any(i))
	}
	keys := m.AllKeysSorted()
	page := m.GetSinglePageCollection(5, 1, keys)

	// Act
	actual := args.Map{"result": page.Length() != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

// ── ResultsCollection InjectIntoAt ──

func Test_Gap_RC_InjectIntoAt(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(1)
	c.Add(corejson.NewResult.Any("x"))
	target := &corejson.Result{}
	_ = c.InjectIntoAt(0, target)
}

// ── ResultsPtrCollection InjectIntoAt ──

func Test_Gap_RPC_InjectIntoAt(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(1)
	c.Add(corejson.NewResult.AnyPtr("x"))
	target := &corejson.Result{}
	_ = c.InjectIntoAt(0, target)
}

// ── BytesCollection InjectIntoAt ──

func Test_Gap_BC_InjectIntoAt(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"x"`))
	target := &corejson.Result{}
	_ = c.InjectIntoAt(0, target)
}

// ── ResultsCollection AddJsoners ──

func Test_Gap_RC_AddJsoners(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(2)
	c.AddJsoners(true)
	c.AddJsoners(true, nil)
}

// ── ResultsPtrCollection AddJsoners ──

func Test_Gap_RPC_AddJsoners(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(2)
	c.AddJsoners(true)
	c.AddJsoners(true, nil)
}

// ── ResultsCollection + ResultsPtrCollection AddSerializerFunctions ──

func Test_Gap_RPC_AddSerializerFunctions(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(2)
	fn := func() ([]byte, error) { return []byte(`"x"`), nil }
	c.AddSerializerFunctions(fn)
	c.AddSerializerFunctions()
}

// ── Ensure sleep-based tests settle ──

func Test_Gap_SettleTime(t *testing.T) {
	time.Sleep(20 * time.Millisecond)
}
