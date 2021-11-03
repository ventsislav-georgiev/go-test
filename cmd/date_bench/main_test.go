package main

import "testing"

func BenchmarkDateWithSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dateWithSlice()
	}
}
