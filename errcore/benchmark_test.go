package errcore

import (
	"errors"
	"testing"
)

func BenchmarkMergeErrors(b *testing.B) {
	errs := []error{
		errors.New("err1"),
		errors.New("err2"),
		errors.New("err3"),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeErrors(errs...)
	}
}

func BenchmarkExpecting(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Expecting("title", "expected-val", "actual-val")
	}
}

func BenchmarkExpectingSimple(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ExpectingSimple("title", "expected", "actual")
	}
}

func BenchmarkConcatMessageWithErr(b *testing.B) {
	err := errors.New("base error")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatMessageWithErr("context message", err)
	}
}

func BenchmarkExpectingFuture(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ExpectingFuture("expecting length at least", 5)
	}
}
