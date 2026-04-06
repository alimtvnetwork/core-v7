package corestr

import (
	"testing"
)

// =============================================================================
// Hashmap Benchmarks
// =============================================================================

func BenchmarkHashmap_AddOrUpdate(b *testing.B) {
	hm := New.Hashmap.Cap(64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm.AddOrUpdate("key", "val")
	}
}

func BenchmarkHashmap_GetValue(b *testing.B) {
	hm := New.Hashmap.Cap(64)
	hm.AddOrUpdate("key", "val")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm.GetValue("key")
	}
}

func BenchmarkHashmap_Has(b *testing.B) {
	hm := New.Hashmap.Cap(64)
	hm.AddOrUpdate("key", "val")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm.Has("key")
	}
}

func BenchmarkHashmap_Length(b *testing.B) {
	hm := New.Hashmap.Cap(64)
	for i := 0; i < 100; i++ {
		hm.AddOrUpdate("k"+string(rune(i)), "v")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm.Length()
	}
}

// =============================================================================
// Hashset Benchmarks
// =============================================================================

func BenchmarkHashset_Add(b *testing.B) {
	hs := New.Hashset.Cap(64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hs.Add("item")
	}
}

func BenchmarkHashset_Has(b *testing.B) {
	hs := New.Hashset.Cap(64)
	hs.Add("item")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hs.Has("item")
	}
}

func BenchmarkHashset_Length(b *testing.B) {
	hs := New.Hashset.Cap(64)
	for i := 0; i < 100; i++ {
		hs.Add("item" + string(rune(i)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hs.Length()
	}
}

// =============================================================================
// Collection Benchmarks
// =============================================================================

func BenchmarkCollection_Add(b *testing.B) {
	c := New.Collection.Cap(64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Add("item")
	}
}

func BenchmarkCollection_Length(b *testing.B) {
	c := New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Length()
	}
}

func BenchmarkCollection_IsEmpty(b *testing.B) {
	c := New.Collection.Strings([]string{"a"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.IsEmpty()
	}
}

func BenchmarkCollection_First(b *testing.B) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.First()
	}
}

func BenchmarkCollection_Last(b *testing.B) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Last()
	}
}
