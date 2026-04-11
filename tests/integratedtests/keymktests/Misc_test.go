package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/keymk"
	"github.com/alimtvnetwork/core/coretests/args"
)

// TestKey_Default verifies default key creation.
func TestKey_Default(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "child1", "child2")

	// Act
	compiled := key.Compile()

	// Assert
	actual := args.Map{"result": compiled != "root-child1-child2"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-child1-child2', got ''", actual)
}

// TestKey_DefaultMain verifies main-only key.
func TestKey_DefaultMain(t *testing.T) {
	key := keymk.NewKey.DefaultMain("main")
	actual := args.Map{"result": key.Compile() != "main"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'main', got ''", actual)
}

// TestKey_DefaultStrings verifies string-based construction.
func TestKey_DefaultStrings(t *testing.T) {
	key := keymk.NewKey.DefaultStrings("root", "a", "b")
	actual := args.Map{"result": key.Compile() != "root-a-b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-a-b', got ''", actual)
}

// TestKey_CurlyBraces verifies curly brace option.
func TestKey_CurlyBraces(t *testing.T) {
	key := keymk.NewKey.Curly("root", "a")
	compiled := key.Compile()
	actual := args.Map{"result": compiled != "{root}-{a}"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '{root}-{a}', got ''", actual)
}

// TestKey_SquareBrackets verifies square bracket option.
func TestKey_SquareBrackets(t *testing.T) {
	key := keymk.NewKey.SquareBrackets("root", "a")
	compiled := key.Compile()
	actual := args.Map{"result": compiled != "[root]-[a]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '[root]-[a]', got ''", actual)
}

// TestKey_Parenthesis verifies parenthesis option.
func TestKey_Parenthesis(t *testing.T) {
	key := keymk.NewKey.Parenthesis("root", "a")
	compiled := key.Compile()
	actual := args.Map{"result": compiled != "(root)-(a)"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '(root)-(a)', got ''", actual)
}

// TestKey_AppendChain verifies append chain.
func TestKey_AppendChain(t *testing.T) {
	key := keymk.NewKey.Default("root")
	key.AppendChain("a", "b")
	actual := args.Map{"result": key.Compile() != "root-a-b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-a-b', got ''", actual)
}

// TestKey_AppendChainStrings verifies string append.
func TestKey_AppendChainStrings(t *testing.T) {
	key := keymk.NewKey.Default("root")
	key.AppendChainStrings("x", "y")
	actual := args.Map{"result": key.Compile() != "root-x-y"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-x-y', got ''", actual)
}

// TestKey_SkipEmpty verifies empty entry skipping.
func TestKey_SkipEmpty(t *testing.T) {
	key := keymk.NewKey.Default("root", "", "b")
	compiled := key.Compile()
	actual := args.Map{"result": compiled != "root-b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-b', got ''", actual)
}

// TestKey_Length verifies chain length.
func TestKey_Length(t *testing.T) {
	key := keymk.NewKey.Default("root", "a", "b")
	actual := args.Map{"result": key.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// TestKey_NilLength verifies nil key length.
func TestKey_NilLength(t *testing.T) {
	var key *keymk.Key
	actual := args.Map{"result": key.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil length should be 0", actual)
}

// TestKey_IsEmpty verifies empty check.
func TestKey_IsEmpty(t *testing.T) {
	key := keymk.NewKey.Default("")
	actual := args.Map{"result": key.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty main with no chains should be empty", actual)
}

// TestKey_MainName verifies main name getter.
func TestKey_MainName(t *testing.T) {
	key := keymk.NewKey.Default("myroot")
	actual := args.Map{"result": key.MainName() != "myroot"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'myroot'", actual)
}

// TestKey_KeyChains verifies chain getter.
func TestKey_KeyChains(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	chains := key.KeyChains()
	actual := args.Map{"result": len(chains) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 chain", actual)
}

// TestKey_NilKeyChains verifies nil chains.
func TestKey_NilKeyChains(t *testing.T) {
	var key *keymk.Key
	actual := args.Map{"result": key.KeyChains() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil key should return nil chains", actual)
}

// TestKey_AllRawItems verifies all raw items.
func TestKey_AllRawItems(t *testing.T) {
	key := keymk.NewKey.Default("root", "a", "b")
	items := key.AllRawItems()
	actual := args.Map{"result": len(items) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

// TestKey_NilAllRawItems verifies nil raw items.
func TestKey_NilAllRawItems(t *testing.T) {
	var key *keymk.Key
	actual := args.Map{"result": key.AllRawItems() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

// TestKey_HasInChains verifies chain search.
func TestKey_HasInChains(t *testing.T) {
	key := keymk.NewKey.Default("root", "a", "b")
	actual := args.Map{"result": key.HasInChains("a")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have 'a'", actual)
	actual := args.Map{"result": key.HasInChains("c")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have 'c'", actual)
}

// TestKey_NilHasInChains verifies nil chain search.
func TestKey_NilHasInChains(t *testing.T) {
	var key *keymk.Key
	actual := args.Map{"result": key.HasInChains("x")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

// TestKey_Finalized verifies finalization.
func TestKey_Finalized(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	key.Finalized("b")
	actual := args.Map{"result": key.IsComplete()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be complete", actual)
	actual := args.Map{"result": key.CompiledChain() != "root-a-b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-a-b', got ''", actual)
}

// TestKey_CompiledChain_NotComplete verifies incomplete returns empty.
func TestKey_CompiledChain_NotComplete(t *testing.T) {
	key := keymk.NewKey.Default("root")
	actual := args.Map{"result": key.CompiledChain() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "incomplete should return empty", actual)
}

// TestKey_ClonePtr verifies clone.
func TestKey_ClonePtr(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	cloned := key.ClonePtr()
	actual := args.Map{"result": cloned.Compile() != key.Compile()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone should match original", actual)
}

// TestKey_NilClonePtr verifies nil clone.
func TestKey_NilClonePtr(t *testing.T) {
	var key *keymk.Key
	actual := args.Map{"result": key.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should be nil", actual)
}

// TestKey_String verifies String method.
func TestKey_String(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	actual := args.Map{"result": key.String() != "root-a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-a', got ''", actual)
}

// TestKey_Name verifies Name method.
func TestKey_Name(t *testing.T) {
	key := keymk.NewKey.Default("root")
	actual := args.Map{"result": key.Name() != "root"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root'", actual)
}

// TestKey_KeyCompiled verifies KeyCompiled.
func TestKey_KeyCompiled(t *testing.T) {
	key := keymk.NewKey.Default("r", "a")
	actual := args.Map{"result": key.KeyCompiled() != "r-a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'r-a', got ''", actual)
}

// TestKey_CompileStrings verifies CompileStrings.
func TestKey_CompileStrings(t *testing.T) {
	key := keymk.NewKey.Default("root")
	r := key.CompileStrings("x", "y")
	actual := args.Map{"result": r != "root-x-y"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-x-y', got ''", actual)
}

// TestKey_CompileDefault verifies CompileDefault.
func TestKey_CompileDefault(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	actual := args.Map{"result": key.CompileDefault() != "root-a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-a', got ''", actual)
}

// TestKey_IntRange verifies IntRange.
func TestKey_IntRange(t *testing.T) {
	key := keymk.NewKey.Default("item")
	r := key.IntRange(0, 2)
	actual := args.Map{"result": len(r) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 items", actual)
	actual := args.Map{"result": r[0] != "item-0"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'item-0', got ''", actual)
}

// TestKey_IntRangeEnding verifies IntRangeEnding.
func TestKey_IntRangeEnding(t *testing.T) {
	key := keymk.NewKey.Default("item")
	r := key.IntRangeEnding(1)
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
}

// TestKey_AppendChainKeys verifies keys append.
func TestKey_AppendChainKeys(t *testing.T) {
	key1 := keymk.NewKey.Default("root")
	key2 := keymk.NewKey.Default("sub", "a")
	key1.AppendChainKeys(key2)
	actual := args.Map{"result": key1.Compile() != "root-sub-a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-sub-a', got ''", actual)
}

// TestKey_AppendChainKeys_Nil verifies nil key skipping.
func TestKey_AppendChainKeys_Nil(t *testing.T) {
	key := keymk.NewKey.Default("root")
	key.AppendChainKeys(nil)
	actual := args.Map{"result": key.Compile() != "root"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root', got ''", actual)
}

// TestKey_ConcatNewUsingKeys verifies concat.
func TestKey_ConcatNewUsingKeys(t *testing.T) {
	key1 := keymk.NewKey.Default("root")
	key2 := keymk.NewKey.Default("sub")
	result := key1.ConcatNewUsingKeys(key2)
	actual := args.Map{"result": result.Compile() != "root-sub"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-sub', got ''", actual)
}

// TestKey_CompileKeys verifies compile with extra keys.
func TestKey_CompileKeys(t *testing.T) {
	key := keymk.NewKey.Default("root")
	sub := keymk.NewKey.Default("sub", "x")
	r := key.CompileKeys(sub)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestKey_CompileKeys_Empty verifies empty compile keys.
func TestKey_CompileKeys_Empty(t *testing.T) {
	key := keymk.NewKey.Default("root")
	r := key.CompileKeys()
	actual := args.Map{"result": r != "root"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root', got ''", actual)
}

// TestKey_JoinUsingJoiner verifies custom joiner.
func TestKey_JoinUsingJoiner(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	r := key.JoinUsingJoiner("/", "b")
	actual := args.Map{"result": r != "root/a/b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root/a/b', got ''", actual)
}

// TestKey_CompileReplaceCurlyKeyMap verifies curly replace.
func TestKey_CompileReplaceCurlyKeyMap(t *testing.T) {
	key := keymk.NewKey.Curly("root", "name")
	compiled := key.CompileReplaceCurlyKeyMap(map[string]string{
		"root": "myroot",
		"name": "myname",
	})
	actual := args.Map{"result": compiled == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestOption_ClonePtr verifies option clone.
func TestOption_ClonePtr(t *testing.T) {
	opt := keymk.JoinerOption.ClonePtr()
	actual := args.Map{"result": opt == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone should not be nil", actual)
	actual := args.Map{"result": opt.Joiner != keymk.JoinerOption.Joiner}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "joiner should match", actual)
}

// TestOption_NilClonePtr verifies nil option clone.
func TestOption_NilClonePtr(t *testing.T) {
	var opt *keymk.Option
	actual := args.Map{"result": opt.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should be nil", actual)
}

// TestOption_Clone verifies value clone.
func TestOption_Clone(t *testing.T) {
	opt := keymk.JoinerOption.Clone()
	actual := args.Map{"result": opt.Joiner != keymk.JoinerOption.Joiner}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "joiner should match", actual)
}

// TestOption_IsAddEntryRegardlessOfEmptiness verifies option method.
func TestOption_IsAddEntryRegardlessOfEmptiness(t *testing.T) {
	opt := &keymk.Option{IsSkipEmptyEntry: false}
	actual := args.Map{"result": opt.IsAddEntryRegardlessOfEmptiness()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be true when not skipping", actual)
	opt2 := &keymk.Option{IsSkipEmptyEntry: true}
	actual := args.Map{"result": opt2.IsAddEntryRegardlessOfEmptiness()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be false when skipping", actual)
	var nilOpt *keymk.Option
	actual := args.Map{"result": nilOpt.IsAddEntryRegardlessOfEmptiness()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

// TestKey_PathTemplate verifies path template.
func TestKey_PathTemplate(t *testing.T) {
	key := keymk.NewKey.PathTemplate("root", "sub")
	r := key.Compile()
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestKey_PathTemplateDefault verifies default path template.
func TestKey_PathTemplateDefault(t *testing.T) {
	key := keymk.NewKey.PathTemplateDefault("sub")
	actual := args.Map{"result": key.Compile() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestKey_Strings verifies Strings method.
func TestKey_Strings(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	s := key.Strings()
	actual := args.Map{"result": len(s) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}
