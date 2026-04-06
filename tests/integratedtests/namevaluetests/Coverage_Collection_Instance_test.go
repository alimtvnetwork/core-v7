package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/namevalue"
)

func Test_Cov_Instance_Dispose(t *testing.T) {
	inst := namevalue.Instance[string, string]{Name: "n", Value: "v"}
	inst.Dispose()
	if inst.Name != "" || inst.Value != "" {
		t.Error("expected zeroed")
	}
}

func Test_Cov_Instance_IsNull(t *testing.T) {
	inst := namevalue.Instance[string, int]{Name: "n", Value: 1}
	if inst.IsNull() {
		t.Error("expected not null")
	}
}

func Test_Cov_Collection_PrependIf_False(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.PrependIf(false, namevalue.Instance[string, string]{Name: "b", Value: "2"})
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov_Collection_PrependUsingFuncIf(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.PrependUsingFuncIf(true, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "b", Value: "2"}}
	})
	if c.Length() != 2 || c.Items[0].Name != "b" {
		t.Error("expected b first")
	}
	c.PrependUsingFuncIf(false, nil)
	if c.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_Collection_AppendUsingFuncIf(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendUsingFuncIf(true, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	})
	if c.Length() != 1 {
		t.Error("expected 1")
	}
	c.AppendUsingFuncIf(false, nil)
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov_Collection_AppendPrependIf(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "m", Value: "1"})
	pre := []namevalue.Instance[string, string]{{Name: "p", Value: "0"}}
	post := []namevalue.Instance[string, string]{{Name: "a", Value: "2"}}
	c.AppendPrependIf(true, pre, post)
	if c.Length() != 3 {
		t.Error("expected 3")
	}
	c.AppendPrependIf(false, pre, post)
	if c.Length() != 3 {
		t.Error("expected still 3")
	}
}

func Test_Cov_Collection_ConcatNewPtr(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	item := namevalue.Instance[string, string]{Name: "b", Value: "2"}
	n := c.ConcatNewPtr(&item)
	if n.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_Collection_AddsIf(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsIf(false, namevalue.Instance[string, string]{Name: "a"})
	if c.Length() != 0 {
		t.Error("expected 0")
	}
	c.AddsIf(true, namevalue.Instance[string, string]{Name: "a"})
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov_Collection_ErrorUsingMessage(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	if c.ErrorUsingMessage("msg") != nil {
		t.Error("expected nil for empty")
	}
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.ErrorUsingMessage("msg") == nil {
		t.Error("expected error")
	}
}

func Test_Cov_Collection_ClonePtr_Nil(t *testing.T) {
	var c *namevalue.Collection[string, string]
	if c.ClonePtr() != nil {
		t.Error("expected nil")
	}
}

func Test_Cov_Collection_Clear_Nil(t *testing.T) {
	var c *namevalue.Collection[string, string]
	if c.Clear() != nil {
		t.Error("expected nil")
	}
}

func Test_Cov_Collection_JoinCsv(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.JoinCsv() == "" {
		t.Error("expected non-empty")
	}
}

func Test_Cov_Collection_JoinCsvLine(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.JoinCsvLine() == "" {
		t.Error("expected non-empty")
	}
}

func Test_Cov_Collection_IsEqualByString(t *testing.T) {
	a := namevalue.NewGenericCollectionDefault[string, string]()
	a.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	b := namevalue.NewGenericCollectionDefault[string, string]()
	b.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if !a.IsEqualByString(b) {
		t.Error("expected equal")
	}
	// diff
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "2"})
	if a.IsEqualByString(c) {
		t.Error("expected not equal")
	}
	// nil
	var d *namevalue.Collection[string, string]
	if !d.IsEqualByString(nil) {
		t.Error("both nil equal")
	}
	if d.IsEqualByString(a) {
		t.Error("nil vs non-nil")
	}
	if a.IsEqualByString(d) {
		t.Error("non-nil vs nil")
	}
	// diff length
	e := namevalue.NewGenericCollectionDefault[string, string]()
	if a.IsEqualByString(e) {
		t.Error("diff length")
	}
}

func Test_Cov_Collection_HasCompiledString(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	_ = c.CompiledLazyString()
	if !c.HasCompiledString() {
		t.Error("expected compiled")
	}
	// second call uses cached
	_ = c.CompiledLazyString()
}

func Test_Cov_Collection_AddsPtr(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	item := namevalue.Instance[string, string]{Name: "a", Value: "1"}
	c.AddsPtr(&item, nil)
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov_Collection_JsonString(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.JsonString() == "" {
		t.Error("expected non-empty")
	}
	e := namevalue.EmptyGenericCollection[string, string]()
	if e.JsonString() != "" {
		t.Error("expected empty for empty coll")
	}
}

func Test_Cov_Collection_JoinJsonStrings(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.JoinJsonStrings(",") == "" {
		t.Error("expected non-empty")
	}
}
