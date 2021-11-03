package main

import (
	"testing"
)

var i myint
var iWithKind = myincWithKind{typeName: "myint"}

func BenchmarkTypeMethod(b *testing.B) {
	typeMethod(&i, b.N)
}

func BenchmarkInterfaceMethod(b *testing.B) {
	interfaceMethod(&i, b.N)
}

func BenchmarkTypeSwitch(b *testing.B) {
	typeSwitch(&i, b.N)
}

func BenchmarkTypeAssertion(b *testing.B) {
	typeAssertion(&i, b.N)
}

func BenchmarkKindCompare(b *testing.B) {
	kindCompare(&iWithKind, b.N)
}
