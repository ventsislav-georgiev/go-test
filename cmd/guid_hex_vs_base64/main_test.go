package main

import "testing"

func BenchmarkGuidHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		guidHex()
	}
}

func BenchmarkGuidBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		guidBase64()
	}
}

func BenchmarkGuidBase62(b *testing.B) {
	for i := 0; i < b.N; i++ {
		guidBase62()
	}
}
