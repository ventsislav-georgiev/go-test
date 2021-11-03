package main

import "testing"

func BenchmarkNoException(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noException()
	}
}

func BenchmarkWithException(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withException()
	}
}
