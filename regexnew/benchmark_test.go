package regexnew

import (
	"testing"
)

func BenchmarkCreateLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateLock(`^\d+$`)
	}
}

func BenchmarkLazyRegex_Compile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lr := &LazyRegex{
			pattern:  `^\d+$`,
			compiler: CreateLock,
		}
		lr.Compile()
	}
}

func BenchmarkLazyRegex_IsMatch_Compiled(b *testing.B) {
	lr := &LazyRegex{
		pattern:  `^\d+$`,
		compiler: CreateLock,
	}
	lr.Compile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lr.IsMatch("12345")
	}
}

func BenchmarkLazyRegex_IsMatch_Miss(b *testing.B) {
	lr := &LazyRegex{
		pattern:  `^\d+$`,
		compiler: CreateLock,
	}
	lr.Compile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lr.IsMatch("abc")
	}
}

func BenchmarkLazyRegexMap_CreateOrExistingLock_Hit(b *testing.B) {
	// prime the cache
	lazyRegexOnceMap.CreateOrExistingLock(`^bench-hit$`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lazyRegexOnceMap.CreateOrExistingLock(`^bench-hit$`)
	}
}

func BenchmarkLazyRegexMap_CreateOrExistingLock_Miss(b *testing.B) {
	patterns := make([]string, b.N)
	for i := range patterns {
		patterns[i] = `^miss-` + string(rune(i%26+'a')) + `$`
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lazyRegexOnceMap.CreateOrExistingLock(patterns[i])
	}
}

func BenchmarkIsMatchLock(b *testing.B) {
	// prime
	IsMatchLock(`^\d+$`, "123")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsMatchLock(`^\d+$`, "123")
	}
}
