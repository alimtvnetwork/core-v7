package mutexbykey

import (
	"strconv"
	"testing"
)

func BenchmarkGet_ExistingKey(b *testing.B) {
	Get("bench-existing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Get("bench-existing")
	}
	b.StopTimer()
	Delete("bench-existing")
}

func BenchmarkGet_NewKey(b *testing.B) {
	keys := make([]string, b.N)
	for i := range keys {
		keys[i] = "bench-new-" + strconv.Itoa(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Get(keys[i])
	}
	b.StopTimer()
	for _, k := range keys {
		Delete(k)
	}
}

func BenchmarkDelete(b *testing.B) {
	keys := make([]string, b.N)
	for i := range keys {
		keys[i] = "bench-del-" + strconv.Itoa(i)
		Get(keys[i])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Delete(keys[i])
	}
}

func BenchmarkGet_Contention(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			Get("contention-key-" + strconv.Itoa(i%100))
			i++
		}
	})
	b.StopTimer()
	// cleanup
	for i := 0; i < 100; i++ {
		Delete("contention-key-" + strconv.Itoa(i))
	}
}
