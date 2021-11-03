package main

import (
	"fmt"
	"unsafe"
)

type calc struct {
	x int
	y int
}

func (c *calc) Sum() int {
	return c.x + c.y
}

type summer interface {
	Sum() int
}

type interfaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

func main() {
	c := calc{x: 10, y: 5}
	var s summer = &c
	var i interface{} = s
	sw := (*interfaceWords)(unsafe.Pointer(&s))
	iw := (*interfaceWords)(unsafe.Pointer(&i))

	fmt.Printf("summer      typ: %v, data: %v\n", sw.typ, sw.data)
	fmt.Printf("interface{} typ: %v, data: %v", iw.typ, iw.data)
}
