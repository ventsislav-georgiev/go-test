package main

type myint int

func (i *myint) inc() {
	*i = *i + 1
}

type Incrementer interface {
	inc()
}

func typeMethod(v *myint, n int) {
	for i := 0; i < n; i++ {
		v.inc()
	}
}

func interfaceMethod(v Incrementer, n int) {
	for i := 0; i < n; i++ {
		v.inc()
	}
}

func typeSwitch(unknown interface{}, n int) {
	for i := 0; i < n; i++ {
		switch v := unknown.(type) {
		case *myint:
			v.inc()
		}
	}
}

func typeAssertion(unknown interface{}, n int) {
	for i := 0; i < n; i++ {
		if newint, ok := unknown.(*myint); ok {
			newint.inc()
		}
	}
}

type myincWithKind struct {
	typeName string
	value    int
}

func (i *myincWithKind) inc() {
	i.value = i.value + 1
}

func (i *myincWithKind) kind() string {
	return i.typeName
}

type IncWrapper interface {
	kind() string
	inc()
}

func kindCompare(incWrapper IncWrapper, n int) {
	for i := 0; i < n; i++ {
		if incWrapper.kind() == "myint" {
			incWrapper.inc()
		}
	}
}
