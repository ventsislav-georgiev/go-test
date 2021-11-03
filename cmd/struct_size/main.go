package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	f  bool  // bool is 1, +padding is 4
	f2 bool  // bool is 1, +padding is 4
	n  int32 // 4 as the biggest field in this struct
}

func main() {
	fmt.Println(unsafe.Sizeof(Foo{}))
}
