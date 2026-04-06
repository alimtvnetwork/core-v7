package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/namevalue"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Cov_Instance_Dispose(t *testing.T) {
	inst := namevalue.Instance[string, string]{Name: "n", Value: "v"}
	inst.Dispose()
	actual := args.Map{"result": inst.Name != "" || inst.Value != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected zeroed", actual)
}

func Test_Cov_Instance_IsNull(t *testing.T) {
	inst := namevalue.Instance[string, int]{Name: "n", Value: 1}
	actual := args.Map{"result": inst.IsNull()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not null", actual)
}

func Test_Cov_Collection_PrependIf_False(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.PrependIf(false, namevalue.Instance[string, string]{Name: "b", Value: "2"})
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov_Collection_PrependUsingFuncIf(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.PrependUsingFuncIf(true, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "b", Value: "2"}}
	})
	actual := args.Map{"result": c.Length() != 2 || c.Items[0].Name != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b first", actual)
	c.PrependUsingFuncIf(false, nil)
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_Collection_AppendUsingFuncIf(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendUsingFuncIf(true, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	})
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	c.AppendUsingFuncIf(false, nil)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov_Collection_AppendPrependIf(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "m", Value: "1"})
	pre := []namevalue.Instance[string, string]{{Name: "p", Value: "0"}}
	post := []namevalue.Instance[string, string]{{Name: "a", Value: "2"}}
	c.AppendPrependIf(true, pre, post)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	c.AppendPrependIf(false, pre, post)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected still 3", actual)
}

func Test_Cov_Collection_ConcatNewPtr(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	item := namevalue.Instance[string, string]{Name: "b", Value: "2"}
	n := c.ConcatNewPtr(&item)
	actual := args.Map{"result": n.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_Collection_AddsIf(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsIf(false, namevalue.Instance[string, string]{Name: "a"})
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	c.AddsIf(true, namevalue.Instance[string, string]{Name: "a"})
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov_Collection_ErrorUsingMessage(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	actual := args.Map{"result": c.ErrorUsingMessage("msg") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.ErrorUsingMessage("msg") == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Cov_Collection_ClonePtr_Nil(t *testing.T) {
	var c *namevalue.Collection[string, string]
	actual := args.Map{"result": c.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Cov_Collection_Clear_Nil(t *testing.T) {
	var c *namevalue.Collection[string, string]
	actual := args.Map{"result": c.Clear() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Cov_Collection_JoinCsv(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.JoinCsv() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Cov_Collection_JoinCsvLine(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.JoinCsvLine() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Cov_Collection_IsEqualByString(t *testing.T) {
	a := namevalue.NewGenericCollectionDefault[string, string]()
	a.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	b := namevalue.NewGenericCollectionDefault[string, string]()
	b.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": a.IsEqualByString(b)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	// diff
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "2"})
	actual := args.Map{"result": a.IsEqualByString(c)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	// nil
	var d *namevalue.Collection[string, string]
	actual := args.Map{"result": d.IsEqualByString(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil equal", actual)
	actual := args.Map{"result": d.IsEqualByString(a)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil", actual)
	actual := args.Map{"result": a.IsEqualByString(d)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil vs nil", actual)
	// diff length
	e := namevalue.NewGenericCollectionDefault[string, string]()
	actual := args.Map{"result": a.IsEqualByString(e)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "diff length", actual)
}

func Test_Cov_Collection_HasCompiledString(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	_ = c.CompiledLazyString()
	actual := args.Map{"result": c.HasCompiledString()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected compiled", actual)
	// second call uses cached
	_ = c.CompiledLazyString()
}

func Test_Cov_Collection_AddsPtr(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	item := namevalue.Instance[string, string]{Name: "a", Value: "1"}
	c.AddsPtr(&item, nil)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov_Collection_JsonString(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.JsonString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	e := namevalue.EmptyGenericCollection[string, string]()
	actual := args.Map{"result": e.JsonString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for empty coll", actual)
}

func Test_Cov_Collection_JoinJsonStrings(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.JoinJsonStrings(",") == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}
