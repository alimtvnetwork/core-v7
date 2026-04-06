package codestack

import (
	"testing"
)

func BenchmarkTrace_Create(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New.Create(0)
	}
}

func BenchmarkTrace_Default(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New.Default()
	}
}

func BenchmarkStackTrace_Default(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New.StackTrace.Default(0, 5)
	}
}

func BenchmarkStackTrace_SkipOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New.StackTrace.SkipOne()
	}
}

func BenchmarkTrace_String(b *testing.B) {
	trace := New.Default()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trace.String()
	}
}

func BenchmarkNameOf_All(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NameOf.All("github.com/alimtvnetwork/core/codestack.BenchmarkNameOf_All")
	}
}
