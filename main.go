package main

import (
	"fmt"
	"runtime"
)

type Foo struct {
	v []byte
}

func main() {
	foos := make([]Foo, 1_000)
	printAlloc()

	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()

	two := keepFirstTwoElementsOnly(foos)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}

func keepFirstTwoElementsOnly(foos []Foo) []Foo {
	return foos[:2]

	// Memory leaks can be avoided in the following two ways:
	// Memory leaks can be avoided in the following two ways:
	//
	// res := make([]Foo, 2)
	// copy(res, foos)
	// return res
	//
	//for i := 2; i < len(foos); i++ {
	//	foos[i].v = nil
	//}
	//return foos[:2]
}
