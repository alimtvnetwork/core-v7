package coredynamic

import (
	"testing"
)

// =============================================================================
// Dynamic Benchmarks
// =============================================================================

func BenchmarkNewDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDynamic("test-data", true)
	}
}

func BenchmarkNewDynamicPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDynamicPtr("test-data", true)
	}
}

func BenchmarkDynamic_Value(b *testing.B) {
	d := NewDynamic("hello", true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Value()
	}
}

func BenchmarkDynamic_IsValid(b *testing.B) {
	d := NewDynamic("hello", true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.IsValid()
	}
}

func BenchmarkDynamic_ReflectKind(b *testing.B) {
	d := NewDynamic("hello", true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.ReflectKind()
	}
}

func BenchmarkDynamic_ReflectTypeName(b *testing.B) {
	d := NewDynamic("hello", true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.ReflectTypeName()
	}
}

func BenchmarkDynamic_Length(b *testing.B) {
	d := NewDynamic([]string{"a", "b", "c"}, true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Length()
	}
}

func BenchmarkDynamic_Clone(b *testing.B) {
	d := NewDynamic("hello", true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Clone()
	}
}

// =============================================================================
// Collection[T] Benchmarks
// =============================================================================

func BenchmarkCollection_Add(b *testing.B) {
	c := NewCollection[string](64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Add("item")
	}
}

func BenchmarkCollection_At(b *testing.B) {
	c := NewCollection[string](10)
	c.Add("a")
	c.Add("b")
	c.Add("c")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.At(1)
	}
}

func BenchmarkCollection_Length(b *testing.B) {
	c := NewCollection[string](10)
	c.Add("a")
	c.Add("b")
	c.Add("c")
	c.Add("d")
	c.Add("e")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Length()
	}
}

func BenchmarkCollection_Contains(b *testing.B) {
	c := NewCollection[string](101)
	for i := 0; i < 100; i++ {
		c.Add("item")
	}
	c.Add("target")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(c, "target")
	}
}

func BenchmarkCollection_IndexOf(b *testing.B) {
	c := NewCollection[int](100)
	for i := 0; i < 100; i++ {
		c.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IndexOf(c, 99)
	}
}

// =============================================================================
// CollectionLock Benchmarks
// =============================================================================

func BenchmarkCollectionLock_AddLock(b *testing.B) {
	c := NewCollection[string](64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.AddLock("item")
	}
}

func BenchmarkCollectionLock_LengthLock(b *testing.B) {
	c := NewCollection[string](10)
	c.Add("a")
	c.Add("b")
	c.Add("c")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.LengthLock()
	}
}

func BenchmarkCollectionLock_ItemsLock(b *testing.B) {
	c := NewCollection[string](10)
	c.Add("a")
	c.Add("b")
	c.Add("c")
	c.Add("d")
	c.Add("e")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ItemsLock()
	}
}
