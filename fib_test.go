package main

import "testing"

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(30) // run fib(30) b.N times
	}
}
func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(10)
	}
}

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(20)
	}
}
