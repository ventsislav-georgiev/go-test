package main

import "testing"

func BenchmarkDefaultJsonSerializer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defaultJsonSerializer()
	}
}

func BenchmarkEasyJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		easyJson()
	}
}
